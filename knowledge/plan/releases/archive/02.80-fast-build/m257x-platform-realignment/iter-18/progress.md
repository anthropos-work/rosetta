---
milestone: M257x
iter: 18
---

# iter-18 — progress

**Type:** tik (under `TOK-01: instrument first, then follow`)

## What happened

### The diagnosis arrived in four minutes, not in a cold cycle — and it refuted the hand-off

iter-17 routed this forward with an explicit instruction *not* to try reproducing the bootstrap failure by
hand, on the grounds that the failing state had healed. **Measuring cost four minutes and changed the
target** (`D-M257x-18-1`):

| # | measurement | consequence |
|---|---|---|
| 1 | the identical `docker run … node cli.js bootstrap` against a throwaway empty schema **exits 0** | the command is fine; the failure is context |
| 2 | `demo-1`'s `directus` schema holds all seven system tables + **87 migrations** | iter-17's stated mechanism (*"never bootstrapped"*) is **refuted** |
| 3 | `directus_policies` has `$t:public_label`; `directus_permissions` has `read` on `task_sub_checks`; `directus_access` has the `role IS NULL` row | the grants said to be missing were present |
| 4 | `docker restart demo-1-directus-1` alone → anon read **403 → 200** | the defect is a stale in-process cache |
| 5 | `docker inspect directus/directus:11.6.1` → `CMD … node cli.js bootstrap && pm2-runtime start` | **the image bootstraps itself** |

Point 5 is the root cause. The compose service is a **second bootstrapper**. Its own timestamped log
crash-loops `schema "directus" does not exist` at `16:01:50 … 16:02:13` (iter-05's `restart: on-failure`
retrying) and then, the moment set-dress's `CREATE SCHEMA` lands, runs the 87 migrations itself at
`16:02:28.98 → 16:02:29.40` — matching `directus_migrations` to the millisecond. Our one-shot lost, exited
non-zero, and the pass printed *"bootstrap failed"* over a schema that was completely bootstrapped.

**The 403 followed one phase later, and not from the bootstrap at all.** `boot_directus_step` — the
post-replay restart whose entire job is to make Directus re-introspect the collections that were registered
*after* it booted — was gated on `DIRECTUS_PROVISIONED`, i.e. on the **exit code of a different step**. So a
lost race cancelled a restart the replay had just made necessary, and the running Directus served content it
was holding to nobody. §5 rule 13's blast-radius note, third occurrence: *when a step's success gates a side
effect, a failure costs both, and the second symptom looks like an unrelated bug.*

### What landed

1. **The provision outcome is a measured POST-CONDITION** (`D-M257x-18-2`). `directus_sentinel_count` is
   extracted so the pre-check and the post-failure re-probe are literally the same question asked twice
   (§2), and a non-zero one-shot with the sentinel **present** now reports PROVISIONED and **names the
   winner**. The rejected alternative was to make our racer win by overriding the image's `CMD` — a
   hand-maintained contradiction of an upstream default, i.e. the class this milestone exists to end.
2. **The serve restart is decoupled** (`D-M257x-18-3`): gated on *this* leg (the directus rows just landed
   on a `--local-content` stack), never on another step's exit code.
3. **The bootstrap output is captured, classified and echoed** — the **third** occurrence of the
   `>/dev/null 2>&1` masking class in this milestone (after `migrate-demo.sh`/M215 F8 and
   `migrate-dev.sh`/iter-16). Its **sibling `CREATE SCHEMA` leg swept in the same pass** (§5 rule 9), and
   the sentinel probe now reports *why* it read 0 (§5 rule 12).
4. **`content_mode` has three states, because two of them were a lie** (`D-M257x-18-4`). `gen_injected_
   override.py:580` re-points `DIRECTUS_BASE_ADDR` at the stack's own Directus whenever the local-content
   service is emitted, and a failed provision does not undo it — so *"the stack stays on the prod-read
   path"* was **false**, and the closing line printed `content:prod-read` **and** `directus=replayed` in one
   sentence. `CHECK-M257x-iter17-setdress-verdict-contradiction` is closed by deriving the field from what
   happened.

### Three existing tests were arguing for the false claim

Second occurrence in this milestone of iter-16's lesson. `test_bootstrap_failure_degrades_to_prod_read_
nonfatal` (the name itself asserted it), `test_create_schema_failure_degrades_nonfatal` and
`test_no_snapshot_with_local_content_skips_provision_no_setu_trip` all required `content:prod-read` on a
`--local-content` stack. All three rewritten to assert the honest label **and** that the false one is
absent.

**And the milestone's own §8 rule caught the iter's own fence.** The first cut of
`test_provision_step_silences_no_leg_of_its_own` asserted `">/dev/null 2>&1" not in <function body>` and
went **red on its own explanation** — the comments above the fixed legs name the pattern precisely so a
future reader knows what it cost. *Comments are allowed unconditionally*; the fence strips comment lines and
then asserts the slice is non-empty (§8 iter-08: *"I scanned it" and "I found nothing to check in it" are
different findings*).

### Live proof — three consecutive cold cycles, and the third one is the interesting one

`rosetta-demo down 1 --purge` (verified to **0 containers**) → `up-injected.sh 1 --no-public-host`, three
times, each consuming rext `fast-build-m257x-iter-18` **from origin**, against platform origin HEAD
`2adcf71`:

| cycle | verdict | ts | which racer won |
|---|---|---|---|
| 1 | `warnings:0 / green:true` | `2026-08-01T16:50:14Z` | our one-shot |
| 2 | `warnings:0 / green:true` | `2026-08-01T17:00:45Z` | our one-shot |
| 3 | `warnings:0 / green:true` | `2026-08-01T17:10:44Z` | **the compose service** |

Checked in at `evidence/av-iter18-cycle{1,2,3}.json`; `anon GET /items/task_sub_checks` = **200** after each.

**Cycle 3 is the causal proof, and it arrived by luck rather than by design — say so.** Cycles 1 and 2
exercised the *winning* path, on which the pre-iter-18 code would also have been green; those two runs prove
no regression, not necessity. Cycle 3 lost the race, printed *"our one-shot bootstrap exited 1, but the
directus_\* system schema IS bootstrapped … won the race and did it. PROVISIONED"* with the losing run's
diagnosis attached (`code: "42P01"` — the error that went to `/dev/null` for two iters), restarted the
service, and came out **green**. On the old code that cycle is iter-17's red.

**The race is nondeterministic — 2 of 3 one way, 1 of 3 the other — and that dissolves the last puzzle in
this thread.** It is why iter-14's three cycles read green under a blind instrument and iter-17's single
cycle read red: the outcome varied per cycle, and until harden-1 added `probe_directus_serves_content`
nothing asked the question that could tell them apart.

## Gate movement

**Gate clause 1 is MET again — this time measured by an instrument that can see served content.**
Three consecutive cold cycles, `green:true / 0 warnings`, distinct monotonically-advancing timestamps, each
following a verified purge, no source change in between, against platform origin HEAD.

**2 of 5 clauses hold** (1 and 4). Clauses 2, 3, 5 outstanding.

## Verification

- **dev-stack `OK 132`** (baseline `OK 125`; +7 new tests, 3 rewritten) · **demo-stack 7F of 1030** (the
  pre-existing `CHECK-M257x-live-clone-suites-red` set, unchanged) · **stack-core 14F of 372** ·
  **stack-injection `OK 286`** — every section at its recorded baseline, zero regressions.
- `bash -n` + `shellcheck -S warning` clean on `dev-setdress.sh`.
- `demo_knob_guard.py` OK both directions (no restaled citations — this iter did not touch `up-injected.sh`).
- **Mutation battery: 7 mutants, 7/7 matched their declared expectation** — 6 declared-RED all killed (each
  naming the test that died), **1 declared-GREEN no-op survived**, every mutant `bash -n`-gated before its
  test run, collected-count read with the exit code, and the unmutated control GREEN **before and after**
  (§8 rule 5 in full). Harness: `.agentspace/scratch/work-m257x/iter18/mutate18.py`.

## Close — 2026-08-01

**Outcome:** The 403 that falsified gate clause 1 is fixed at its real mechanism — a **race between two
Directus bootstrappers** (the image's own `CMD` is `node cli.js bootstrap && pm2-runtime start`) whose lost
outcome cancelled the post-replay restart one phase later. Clause 1 is **MET again, honestly**: three
consecutive cold cycles at `warnings:0 / green:true`, the third of which lost the race and stayed green.
**Gate 1 of 5 → 2 of 5.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (2 of 5 — clause 1 re-proven, clause 4 standing; 2, 3, 5 outstanding)
**Phase 5 grading:** (1) gate-met: n (2 of 5) — (2) triggered-tok: n (this tik moved the primary metric) —
(3) re-scope: n (platform origin HEAD `2adcf71` re-checked at open **and** close, unchanged — occurrence
stays 1 of 2) — (4) user-blocker: n — (5) cap-reached: n (1 tik this session, cap is 5) — (6) protocol-stop:
n — Outcome: continue
**Decisions:** D-M257x-18-1 … D-M257x-18-5 (iter-local `decisions.md`).
**Side-deliverables:** none — every change is inside the declared 2-step scope.
**Routes carried forward:**
- `FIX-M257x-iter18-directus-admin-token-race` → later tik. When the container wins, the admin it creates is
  directus's default `admin@example.com` with **`token = NULL`**, not M23's `local-directus-token-<stack>`
  that studio-desk consumes as `DIRECTUS_TOKEN`. Measured both ways: cycle 3 (container won) left the wrong
  admin; cycles 1–2 (we won) left the right one. Fix = emit `ADMIN_EMAIL`/`ADMIN_PASSWORD`/`ADMIN_TOKEN`
  into **both** override emitters + their parity fence, so whichever racer bootstraps produces the same
  admin. Verification is one `SELECT token FROM directus.directus_users`.
- `CHECK-M257x-iter18-directus-secret-naming` → later tik. `dev-setdress.sh` spells the SECRET/token suffix
  `demo_1` (`tr '-' '_'`); both emitters and `provision.go`'s `DefaultEnvContract` spell it `demo-1`.
  Observed, **not** measured for consequence.
- `FIX-M257x-iter15-directus-versions-403` and `FIX-M257x-iter15-library-category-expansion` → next tik,
  **and re-measure before working on either**: both were measured on a stack whose Directus served nothing,
  and the content layer now serves. A cause measured through a broken instrument is a hypothesis.

**Lessons:**
- *"The failing state has healed"* was true of one artefact and got silently generalised to *"the failure is
  not reproducible."* **Ask what the failing step's actual inputs are** — here a DSN, an image and an empty
  schema — before concluding that only the whole pipeline can produce them. Promoted to
  `platform-alignment.md` §5 as **rule 15**.
- **A nondeterministic defect makes a green run weak evidence.** Two of three cycles never exercised the
  branch under repair; only the third did. Record *which path each run took*, or a battery of greens will
  certify a fix that was never invoked. Promoted to §5 rule 15's second half.
