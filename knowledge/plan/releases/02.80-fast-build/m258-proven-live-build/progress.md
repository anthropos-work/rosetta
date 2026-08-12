# M258 — progress

## Running ledger

- iter-01 (**tok**, bootstrap): Phase 0b gate **YELLOW** — 0 blind areas, but **13 stale line anchors in
  M258's own `overview.md`** (all in-range, all landing on unrelated content; substance held in every
  case) repaired, and **two never-propagated measurements** recorded at the destination: the batch half
  of the composed budget has **no published wall-clock** (M256 was asked and did not report it), and
  M256 **escalated that this suite's timing is not decidable at n=3 on this host** (2.04× spread, no
  trend) — against a gate that is a p50 over n=3. `TOK-01` authored: *measure the composition before
  engineering it*. **World contract RESOLVED → (b) restore after**, refuting (a) on the gate's own text.
  Rung zero found the rext pin **one tag behind origin** (`R0`). `F1` re-verified against code and
  **survived** — it read as already-fixed and is open. — see `iter-01/progress.md`

- iter-02 (tik, `closed-fixed-partial`): **`R0` discharged** (rext re-pinned to `fast-build-m257-close`;
  the third pin copy proven **inert** by path arithmetic, not assumed). Bring-up half re-measured at the
  corrected pin: **395.31 s** (n=1, `load1 2.26`, contended + labelled) — `rc=0`, green, HEADROOM OK,
  identity MATCH, phases complete. **The batch half could not be measured, and the blocker is the
  finding: ISOLATION went RED on the first campaign** — all three UI images carry a **non-minted**
  publishable key. Both causes the assert names were **refuted** (fresh build, overlay present); the
  real mechanism is a third — `.env.demo-1` holds **24 appended Clerkenstein blocks** and this run's is
  the one carrying the foreign key, so **last-wins** wired the UI tier to a real Clerk app.
  `inject.py:89` appends instead of rewriting; `up-injected.sh:2036` runs it `2>/dev/null || true`.
  **Not caused by the re-pin** (its whole `up-injected.sh` diff is comments + `log` strings) and
  **`demo-2` is clean** (last key minted). — see `iter-02/progress.md`

- iter-03 (tik, `closed-fixed-partial`): **the routed diagnosis was refuted in all three claims** — the
  key is **Clerkenstein-minted** (`pk_test_bWFy…` = `marcos-mac-mini.taildc510.ts.net`), **no demo ever
  reached production auth**; `demo-1` was **public-host** (auto-discovered), not localhost-bound; and the
  `|| true` is a deliberate `set -e` guard, not a swallow. Real chain found and fixed (rext
  `fast-build-m258-iter-03`, **on origin**): `inject.py` appends → 24 blocks → **`_stack_minted_pk` reads
  first-wins while every consumer reads last-wins** → false RED; plus **`buildbench` could not express
  `--no-public-host`**, so campaigns silently ran the one mode in which the batch **cannot be driven from
  this host**. **Live-proven with no rebuild**: ISOLATION on the stack that reded went `FAIL (3×
  foreign_pk)` → **`ok: True`, 0 failures**; the real 128-line/24-block env replays to **37/1**,
  idempotent. **Batch half still unmeasured** — blocker discharged, but `load1` 39–46 vs 12 cores from
  **third-party** load (Spotlight + the user's own project) and the headroom gate correctly refuses.
  — see `iter-03/progress.md`

## Next-iter routing

- ✅ **iter-03 discharged `FIX-M258-iter02-inject-appends-and-swallows`** — in substance, with its stated
  cause **retracted** (see the ledger entry above and `iter-03/decisions.md` D8).
- **iter-04 (tik, under `TOK-01`)** — **`MEASURE-M258-batch-half`**, unchanged as `TOK-01` step 1's
  outstanding deliverable and still the milestone's primary unknown. The defect blocker is **gone and
  live-proven gone**; what it now needs is **headroom**. Pre-flight `assert-headroom --profile macmini`
  (name the host — never bare), then
  `buildbench run 1 --reps 1 --profile macmini --no-public-host --label m258-iter04` and the full
  Playthrough batch. **State `load1` and the environment with both halves**, and publish the batch's
  **spread** beside any p50 (`C2`, the 2.04× decidability caveat).
- **`FIX-M258-iter03-guard-scans-its-own-scratch`** (net-new) — `test_decommissioned_instruction_guard`
  walks `demo-stack/stacks/**`, which `demo-stack/.gitignore:8` ignores, and reports the *platform's*
  source inside a demo's ephemeral clone as a rext named-consumer list. **Proven pre-existing** by
  running it in the pristine clone at `fast-build-m257-close`: identical 2 failures. Fires on any box
  that has ever run a demo. Plausibly the same root as M257's *"the stack-core sweep did not complete."*
- Then, unchanged: wire the batch-gate at `up-injected.sh:2810` under `D-v28-3` semantics → land the
  world-contract restore leg (b) → the composed 3× cold campaign, **spread published beside the p50**.
- ⚠️ `demo-2` (11 containers) and the 5-container dev stack are the **user's**: do not tear down,
  re-seed, restart or reset either. `demo-1` is left UP as iter-04's reproduction.
  ✅ **The "must not be browsed — it talks to a real Clerk app" warning is WITHDRAWN** (`iter-03/
  decisions.md` D10): the premise was refuted and the live ISOLATION assert returns `ok: True` over all
  8 images. `demo-1` **is** tailnet-reachable (auto-discovered public host, `0.0.0.0`, real LE cert) —
  which `iter-02` D7 denied — but that is the **documented** demo posture (`safety.md` Part 3), not an
  exposure of production auth.

### Also routed from iter-02 (smaller, same tik or later)

- **`CHECK-M258-iter02-studio-desk-is-the-untouched-leg`** — `ui_studio_desk` **115.35 s** is the
  largest UI leg and the one L1 never touched (L1 multi-staged the two *Next* apps). Named suspect for
  the n=1 vs M257-n=3 delta and a candidate lever if the composed budget needs room. **Confirm against
  n≥3 before any claim.**
- **`ROUTE-M258-iter02-isolation-names-two-causes-not-three`** — the refusal text offers two
  explanations and both were refuted; a refusal naming the wrong cause sends the reader at the build.
- **`ROUTE-M258-iter02-headroom-defaults-to-billion`** — bare `assert-headroom` grades against
  `billion.json`; the host must be named every time (cluster 4's shape, second entry point).
- **`ROUTE-M258-iter02-purge-did-not-clear-the-stack-dir`** — 24 accumulated blocks means the stack dir
  survived repeated `--purge` cycles: `verification.md`'s **F-9** instance.

## Carried known-context

`TOK-01` § *Known-context* #1–#6 — `R0` (stale pin) · `C1` (batch half unmeasured) · `C2` (n=3
decidability, 2.04× spread) · `F1` (`FIX-M257-content-stories-pair-count`, verified open; gates the
**content-stories sweep**, not the batch) · `F2` (`ptvalidate` unwired) · the **SUSPECT-UNROUTED** rule
for every inherited M257x / M257 item. Not deferrals.
