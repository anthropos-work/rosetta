---
milestone: M257x
iter: 18
iteration_type: tik
status: closed-fixed
opened: 2026-08-01
closed: 2026-08-01
---

# iter-18 — the Directus bootstrap race, and the restart it suppressed

**Type:** tik
**Active strategy reference:** `TOK-01: instrument first, then follow` (milestone-root `decisions.md`) —
step 2, *"fix the mechanism, not the symptom"*, and step 5, *"prove it cold"*.

## Step 0 — re-survey (mandatory, done before targeting)

- **Platform origin HEAD re-checked at open:** `2adcf714bd877a205e8948f59a23db49b884c054` (2026-07-31 15:58,
  *"Merge pull request #23 … chore/drop-wundergraph"*) — **unchanged** since iter-12. The re-scope trigger
  stays at **occurrence 1 of 2**.
- **The routed target is still live and still meaningful.** `dev-stack/dev-setdress.sh:258-267` still runs
  `node cli.js bootstrap >/dev/null 2>&1`; `demo-1` still answers `403` on the anon read.

### And the re-survey went one step further than the hand-off asked, which changed the target

iter-17 routed this forward as *"make the failure legible first, then fix what it reveals"*, and its
hand-off said the diagnosis could only arrive through a cold cycle *(~11 min)* because the failing state had
healed. **That is false, and measuring it cost four minutes.** The bootstrap's failing state is reproducible
in isolation, and the reproduction refutes the hand-off's central hypothesis.

**Measured, in this order:**

1. **The bootstrap command is not broken.** Run verbatim against a throwaway empty schema on the same
   Postgres (`directus_repro18`), the identical `docker run … node cli.js bootstrap` **exits 0** and applies
   all 87 migrations. So the failure is *context*, not command.
2. **The system schema on `demo-1` is fully bootstrapped** — `directus_collections`, `directus_roles`,
   `directus_policies`, `directus_permissions`, `directus_access`, `directus_users`, `directus_migrations`
   all present; 87 migration rows. iter-17's stated mechanism — *"the system schema whose grants serve it to
   an anonymous reader was never bootstrapped"* — is **REFUTED**.
3. **The public grants exist.** `directus_policies` holds `$t:public_label`; `directus_permissions` holds
   `read` on `task_sub_checks` bound to that policy; `directus_access` holds the `role IS NULL` row that
   makes it the public policy. The thing iter-17 said was missing was present the whole time.
4. **`docker restart demo-1-directus-1` turns the anon read `403` → `200`.** Three polls, ~9 s. The content
   and the grants were always there; the *running process* held a boot-time cache from before they landed.
5. **Who bootstrapped it, and why ours "failed":** `docker inspect directus/directus:11.6.1` →
   `CMD ["/bin/sh","-c",": && node cli.js bootstrap && pm2-runtime start ecosystem.config.cjs ;"]`.
   **The image bootstraps itself on every container start.** The container's own log (`docker logs -t`)
   crash-loops `schema "directus" does not exist` at `16:01:50, :52, :54, :56, :58, 16:02:01, :05, :13`
   (iter-05's `restart: on-failure` retrying), and then — the moment set-dress's `CREATE SCHEMA` lands —
   runs the migrations itself at `16:02:28.98 → 16:02:29.40` and finishes `INFO: Done`. The
   `directus_migrations` timestamps agree to the millisecond.

**So the mechanism is a RACE between two bootstrappers, and the 403 is its side effect:**

    compose directus container (image CMD)  ─┐
                                             ├─ both run `node cli.js bootstrap`
    dev-setdress provision_directus_step    ─┘

    container WINS  ->  our one-shot exits non-zero
                    ->  "⚠ bootstrap failed — skipping local content"   (output discarded: >/dev/null 2>&1)
                    ->  DIRECTUS_PROVISIONED=0
                    ->  the replay STILL runs and STILL succeeds (structure + serve rows + 11 986 rows)
                    ->  boot_directus_step is gated on DIRECTUS_PROVISIONED=1  ->  NEVER RUNS
                    ->  the container keeps its boot-time registry cache  ->  anon GET = 403
                    ->  verdict prints `content:prod-read` AND `directus=replayed` in one sentence, exit 0

This is **§5 rule 13's blast-radius note for the third time**: *when a step's success gates a side effect, a
failure costs both, and the second symptom looks like an unrelated bug.* Here the failing step and the
suppressed side effect are not even in the same phase.

## Cluster / target identified

`FIX-M257x-iter17-directus-bootstrap-blind` + `CHECK-M257x-iter17-setdress-verdict-contradiction`, both
routed to iter-18 by `D-M257x-17-3`. The re-survey does not substitute the target — it **narrows** it: the
legibility fix is still needed (it is why this cost two iters), but the repair is no longer unknown.

## Hypothesis

The 403 that falsified gate clause 1 is caused by a **suppressed post-replay restart**, not by a missing
bootstrap. Therefore:

1. Deriving the provision outcome from a **measured post-condition** (is the sentinel system table there?)
   instead of from **one racer's exit code** makes `DIRECTUS_PROVISIONED=1` on the very run that reported
   failure — which is the truth, because the schema *is* bootstrapped.
2. Gating `boot_directus_step` on **the replay having succeeded and a compose service existing** — rather
   than on a different step's exit code — makes the restart happen whether or not our one-shot won.
3. Capturing and classifying the bootstrap output makes the *next* occurrence legible in one read instead
   of two iters.

## Expected lift

`autoverify` on a cold cycle: `warnings:1 / green:false` → **`warnings:0 / green:true`**, with
`✓ directus-serves-content`. That is gate clause 1's first cycle. (Clause 1 needs three; this iter claims
**one measured cold cycle**, not the clause — the remaining two are a conjunction and belong to whichever
iter can run them without another source change in between.)

## Phase plan (a declared 2-step shape — the scope-creep tripwire counts against THIS)

- **Step 1 — legibility + honest verdict.** Capture / classify / report the bootstrap outcome (RF-1's shape,
  third file); sweep the sibling silenced calls in the same function (§5 rule 9); make the provision result
  a measured post-condition; forbid the `content:prod-read` + `directus=replayed` contradiction.
- **Step 2 — the causal repair.** Decouple `boot_directus_step` from `DIRECTUS_PROVISIONED`.
- Tests + a mutation battery with a **declared-GREEN no-op control** (§8 rule 5).
- Cold cycle on `demo-1` to prove it live.

Anything else surfaced is routed with a named handler.

## Escalation conditions

- A cold cycle that still 403s ⇒ the restart hypothesis is wrong; re-open at the serve layer, do not patch
  the probe.
- Any repair that would need a platform-repo edit ⇒ `demopatch` or escalate. (None expected: both files are
  rext-owned.)

## Acceptable close-no-lift outcomes

If the cold cycle refutes the restart hypothesis, the iter closes `closed-no-lift` with the falsification —
the legibility fix still lands, because its whole purpose is that the *next* diagnosis is one read.
