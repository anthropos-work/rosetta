# M258 iter-05 — progress

**Type:** tik · **Active strategy:** `TOK-01` step 1 — *measure the composition before engineering it.*

## Phase B — the campaign

`buildbench run 1 --reps 3 --profile macmini --no-public-host --label m258-iter05-gateable`,
launched 05:58:04Z at `load1` 3.49, foreground-polled to completion at 06:12:54Z. `CAMP_RC=1`
(**the documented "report is RED" code, not a failure to run** — `0 = ok · 1 = RED · 2 = could not run`).

| rep | total | `peak_load1` | headroom | green | ISOLATION | identity | usable |
|---|---|---|---|---|---|---|---|
| 1 | 344.82 s | 21.77 | **FAIL** | ✅ | ✅ | match | ✗ |
| 2 | 249.13 s | 10.62 | **FAIL** *(floor is 10)* | ✅ | ✅ | match | ✗ |
| 3 | **247.79 s** | 6.53 | **OK** | ✅ | ✅ | match | **✓** |

`gateable: False` · `total_s` p50 **249.13** [min 247.79 – max 344.82] · `phases_complete: true` ·
`host_identity: match ×3` · `isolation_ok: [True, True, True]`.

**Why RED, precisely:** *"rep(s) [1, 2] are not usable measurements"* — both on **headroom**, because the
user's load was still draining when the campaign started (`peak_load1` 21.77, then 10.62 against a floor
of 10 — rep 2 missed by **0.62**). Nothing about the platform was red: **every rep was green, isolated,
identity-matched and phase-complete.**

**`bringup_argv` on all three reps ends `1 --no-public-host`** — read straight out of the rep ledgers.
That field did not exist before iter-03; it is the side-deliverable proving its own point, and it is why
this iter can state its mode from the record instead of grepping a progress line.

## Phase C — THE 395-vs-287 QUESTION IS ANSWERED, and the answer is a cold cache

The inherited priority was: *explain the 395.31 vs 286.99 delta before treating either as "the bring-up
half."* The per-sub-phase table settles it.

| sub-phase | iter-02 (n=1) | iter-05 p50 (n=3) | note |
|---|---|---|---|
| **`ui_studio_desk`** | **115.35 s** | **7.12 s** [7.11–8.05] | **the whole delta** |
| `set_dress` | 80.50 | 81.23 [78.92–117.51] | unchanged — still the largest phase |
| `ui_next_web` | 60.05 | 49.01 [46.60–55.44] | |
| `ui_hiring` | 44.85 | 45.32 [44.51–93.12] | |
| `compose_up` | 32.56 | 43.87 [43.17–44.69] | |
| `host_preflight` | 35.40 | 9.33 [9.22–9.71] | |

```
395.31 − 286.99 = 108.32 s      ← the delta to be explained
115.35 −   7.12 = 108.23 s      ← ui_studio_desk, cold vs warm
```

**The two agree to 0.09 s.** So:

- **The mode was NOT the explanation.** iter-02 and M257's 286.99 s were **both** `--public-host`; the
  record simply mis-stated iter-02's. Comparing them was never a mode comparison.
- **`CHECK-M258-iter02-studio-desk-is-the-untouched-leg` is CONFIRMED — with the nuance that changes what
  to do about it.** studio-desk is **not** a standing 115 s cost. It is ~7 s warm and ~115 s **cold**,
  and iter-02 paid the cold price because that bring-up followed a re-pin. It is a **cache** finding, not
  a lever finding — so it is **not** the reserve `LEVER-M257-L5-setdress` was hoping for, and L5 remains
  the real target: `set_dress` is *still* the largest single phase at **81.23 s**, exactly as M257 left it.
- **Neither 395.31 nor 286.99 is "the bring-up half" for this milestone.** Both are public-host. The
  single-box bring-up half is **247.79 s** (the one gateable rep), with rep 2 independently at 249.13 s —
  the two agree within **1.34 s**.

## The composed arithmetic — the first honest one

| half | value | provenance |
|---|---|---|
| bring-up (single-box) | **247.79 s** | rep 3, **headroom OK**, green, isolated, identity match |
| batch | **129 s** | iter-04, n=1, `load1` 30.4 — **contended**, so likely an over-estimate |
| **composed** | **≈ 376.8 s** | against the **480 s** ceiling → **inside, by ~103 s** |

⚠️ **This is not the gate, and must not be reported as it.** The gate is a **p50 over 3 consecutive cold
cycles** of the *composed* thing, with the batch wired into the bring-up (`TOK-01` step 2, unbuilt) and
the world restored (step 3, unbuilt). What this is: **n=1 + n=1 from two separate runs**, which is the
first evidence that **480 s is reachable** rather than a ceiling-sum hope. `overview.md` § *Budget
honesty* said 480 s *"is reachable only if M257 lands nearer its ~240–300 s estimate"* — M257 landed
286.99 s public-host, and single-box lands **247.79 s**, i.e. **at the bottom of that window**.

**`C2` honoured:** the spread is published beside the p50 — 247.79 / 249.13 / 344.82, and the outlier is
attributed (`peak_load1` 21.77) rather than dropped. buildbench's own note fires too: *"rep 1 ran 38 %
above p50 with no reclaim recorded before it — investigate; this one is NOT explained by eviction."*
It is explained: **load**, not eviction.

## Close — 2026-08-12

**Outcome:** The first **gateable** single-box bring-up half — **247.79 s** (rep 3: headroom OK, green,
ISOLATION ok, identity match, phases complete), corroborated by rep 2 at 249.13 s. The inherited
**395-vs-287 question is answered exactly**: the 108.32 s delta is a **cold `ui_studio_desk` build**
(115.35 s cold vs 7.12 s warm = 108.23 s), both figures were public-host, and neither was the
single-box half. Composed with iter-04's batch, **≈ 376.8 s against a 480 s ceiling** — the first
evidence the ceiling is reachable. Campaign RED on headroom for reps 1–2 only (the user's load draining);
no platform red in any rep.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n *(the gate is a composed p50 over 3 cold cycles; steps 2–4 of
`TOK-01` are unbuilt)* — (2) triggered-tok: n — (3) re-scope: n *(≈377 s is far below the 600 s valve,
and it is not a p50 anyway)* — (4) user-blocker: n — (5) cap-reached: n *(3 tiks)* — (6) protocol-stop: n
— (7) budget-exhausted: **y** — Outcome: **exit-7**
**Decisions:** D13 (this iter's `decisions.md`)
**Side-deliverables:** none — `bringup_argv` landed in iter-03 and merely paid out here.
**Routes carried forward:**

- **`TOK-01` step 2 — wire the batch-gate at the tail hook** (`up-injected.sh:2810`) under `D-v28-3`
  semantics: runs to completion, never halts at first red, never retries, ONE consolidated red set at
  batch end, stack left UP regardless, bring-up exits non-zero and loudly on a non-empty set. **This is
  the next iter's job** and it is now the only thing between the milestone and a composed measurement.
- **`RESTORE-M258-world-contract`** (`TOK-01` step 3) — still owed. ⚠️ **`demo-1` is currently a
  `pt-world` stack**: iter-04's batch `--reset` truncated the demo world, and this campaign's three
  bring-ups re-seeded the demo world each time, then **nothing re-ran the batch**, so demo-1 is presently
  in its **post-bring-up demo state** (not pt-world). The restore leg is still unbuilt and still required
  the moment the batch is wired into the bring-up.
- **`CHECK-M258-iter02-studio-desk-is-the-untouched-leg`** → **CLOSED with a finding**: confirmed as the
  delta's cause, but it is a **cold-cache** cost (7.12 s warm), not a standing lever. Reserve hopes stay
  with **`LEVER-M257-L5-setdress`** — `set_dress` is still the largest phase at **81.23 s**.
- Unchanged: `FIX-M258-iter03-guard-scans-its-own-scratch` ·
  `ROUTE-M258-iter02-isolation-names-two-causes-not-three` ·
  `ROUTE-M258-iter02-headroom-defaults-to-billion` ·
  `ROUTE-M258-iter02-purge-did-not-clear-the-stack-dir` (observed live at iter-04).

**Lessons:**

- **Two numbers for "the same thing" meant the definitions differed — and neither was the thing.**
  395.31 and 286.99 were both public-host; the argument about which was "the bring-up half" was an
  argument between two figures that were both the *wrong mode*. The sub-phase table answered in one line
  what ceiling-arithmetic could not.
- **A cache effect wears a lever's clothing.** A 115 s phase that is 7 s warm looks like the biggest
  optimisation target in the table and is worth nothing to optimise. **Ask cold-or-warm before ranking.**
- **A RED campaign can still contain the measurement you came for.** Two reps failed headroom by margins
  as small as 0.62; the third is clean and is a real number. The instrument reporting RED is not the
  instrument reporting *nothing*.
