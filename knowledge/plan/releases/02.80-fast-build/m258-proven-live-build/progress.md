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

- iter-04 (tik, `closed-fixed`): **THE BATCH HALF EXISTS** — `TOK-01` step 1's outstanding deliverable,
  discharged. **129 s**, `Running 215 tests using 1 worker` → **215 passed**, ptreport
  **`passing=30 failing=0 unimplemented=1`** (the known `will-not-build` TODO), **red set EMPTY**,
  `BATCH_RC=0`. ⚠️ **n=1, not a p50** — `C2`'s 2.04× spread stands, and it is **not** comparable to
  M256's 56.6 s (18 specs vs 215). The headroom gate **refused a `buildbench` cycle outright** (`load1`
  16 → 62, third-party), so the cold cycle was driven as an **operator** (`up-injected.sh`'s pre-flights
  are advisory by design) for the halves contention cannot corrupt. **All four iter-03 fixes confirmed
  live**: single-box mode engaged (`up-injected demo-1 (localhost)`), the minted-host line **visible** in
  the log, `.env.demo-1` **24 → 1**, ISOLATION **ok** on fresh images with `own_pk` decoding to
  `127.0.0.1:15400`. Composed **910 s** vs the 480 s ceiling — **contended, and NOT a gate reading; it
  did not fire the re-scope trigger** (which reads a *p50 after 3 tiks*). The shape it does support: the
  batch half is **small** (~14 %), so M257's 286.99 s + ~129 s ≈ **416 s** would sit inside 480 s.
  — see `iter-04/progress.md`

- iter-05 (tik, `closed-fixed`): **the first GATEABLE single-box bring-up half — 247.79 s** (rep 3:
  headroom OK, green, ISOLATION ok, identity match, phases complete), corroborated by rep 2 at 249.13 s
  (missed headroom by **0.62**) with a contended outlier at 344.82 s (`peak_load1` 21.77). Campaign RED
  **on headroom only** — every rep was green + isolated + identity-matched. **The inherited 395-vs-287
  question is ANSWERED**: `ui_studio_desk` is **115.35 s cold vs 7.12 s warm = 108.23 s**, against a
  delta of **108.32 s** — they agree to 0.09 s. **Both prior figures were `--public-host`**, so neither
  was ever the single-box half, and studio-desk is a **cache** effect, not a lever. `set_dress` is still
  the largest phase (**81.23 s**, vs M257's 82.04) — `LEVER-M257-L5-setdress` remains the real reserve.
  **Composed ≈ 376.8 s vs the 480 s ceiling** (247.79 + 129, two separate n=1 runs — *not* the gate) —
  the first evidence the ceiling is reachable. — see `iter-05/progress.md`

- iter-06 (tik, `closed-fixed`): **THE BATCH GATE IS WIRED — the gate is measurable for the first
  time.** `TOK-01` steps **2 + 3** landed together (`D16`: a batch without a restore leg is a
  *regression*, not a partial delivery — it would make every bring-up end in the dead-CTA world the
  overview warns about). New in rext: `batch-gate.sh` (the `D-v28-3` contract), `restore-presenter-world.sh`
  (world contract (b)), `check-cockpit-roster.py`, `stack-paths.sh`; hooked at `up-injected.sh:2839`
  **after** the `UP.` line. **Proven live end-to-end**: 215 passed, 31 use cases, `passing=30
  unimplemented=1`, **red set EMPTY**, restore **7 s** (vs the 20–45 s assumed), phase table complete with
  **`batch_gate` 166.36 s beside `autoverify` 2.41 s** — without the new anchor those two were one
  168.77 s phase that still *summed* (`D17`). **The live run found a real defect in this iter's own
  restore leg** (`D19`): on a box with **two** rext clones it resolved the stack dir from its own
  location, so the roster went to the live path while the cockpit/content manifests went to a **stale**
  one — a 35-identity stories roster beside an 11-seat pt-world menu, while the gate printed "restored"
  and exited 0. Fixed by asking **docker** where the stack's files are, plus a **post-condition** (`D20`),
  because the defect passed every step-level check. Composed **545.22 s** — n=1, **cold-cache**, rep
  `green:false` and correctly **not gate-usable**; against iter-05's gateable bring-up the projection is
  **414.15 s** vs 480. Knob fence went RED in the direction it exists for (`DEMO_NO_BATCH`
  undiscoverable) and is now OK both ways; **8 pre-existing stale anchors** + a live
  `FIX-M257-anchor-guard-content-drift` instance repaired. — see `iter-06/progress.md`

## Next-iter routing

- ✅ **iter-03 discharged `FIX-M258-iter02-inject-appends-and-swallows`** — in substance, with its stated
  cause **retracted** (see the ledger entry above and `iter-03/decisions.md` D8).
- ✅ **iter-04 discharged `MEASURE-M258-batch-half`** — 129 s, red set empty (n=1, contended).
- ✅ **iter-05 discharged `MEASURE-M258-gateable-composition`** (bring-up side) **and closed
  `CHECK-M258-iter02-studio-desk-is-the-untouched-leg`** with a finding: it is a **cold-cache** cost, not
  a lever.
- ✅ **iter-06 discharged `TOK-01` steps 2 AND 3** — the batch gate is wired at `up-injected.sh:2839` and
  the world-contract restore leg ships with it, both proven live. **`RESTORE-M258-world-contract` is
  CLOSED**: the restore is now a mechanism every bring-up carries (7 s), not a one-off repair, and
  `demo-1` is verified presenter-usable (4 story orgs / 591 users / 12 cockpit seats all resolving).
- **iter-07 (tik, under `TOK-01`) — step 4: the composed 3× cold campaign.** Unblocked; both halves are
  wired into one command. ⚠️ **Run it from the CONSUMPTION clone at a PUSHED tag** — iter-06 measured that
  an authoring-copy bring-up has no sibling `platform/`, so `autoverify`'s `postgres-schemas` probe
  refuses to assert and **every rep grades `green:false` and unusable regardless of its timings**. Publish
  `fast-build-m258-iter-06` to origin, re-pin `stack-demo/rosetta-extensions`, then campaign. Publish the
  **spread beside the p50** (`C2`), on both halves — the batch is now 129 s (iter-04) and 160 s (iter-06)
  with no p50 yet.

<details><summary>superseded routing (iter-06's plan, kept for the audit trail)</summary>

- **iter-06 (tik, under `TOK-01`) — `TOK-01` step 2: wire the batch-gate at the tail hook**
  (`up-injected.sh:2810`, beside the `autoverify.sh` invocation) under `D-v28-3` semantics: runs to
  completion, never halts at first red, **never retries**, ONE consolidated red set at batch end, the
  stack left **UP** regardless, and the bring-up exiting **non-zero and loudly** on a non-empty set.
  **This is now the only thing between the milestone and a composed measurement** — both halves are
  measured separately and the arithmetic (≈377 s) fits, but the gate is a p50 over 3 cold cycles of the
  *composed* thing, which cannot be taken until the batch runs inside the bring-up.
  Then step 3 (world-contract restore) and step 4 (the 3× cold campaign).
- **`RESTORE-M258-world-contract`** (`TOK-01` step 3) — **now owed in FACT**: iter-04's batch `--reset`
  TRUNCATEd the demo world and re-seeded **pt-world**, so `demo-1` is currently a Playthrough world
  behind a cockpit projected from the stories preset — verbatim the state `overview.md` § *The world
  contract* warns about, and the reason resolution **(b) restore after** was chosen.
- **`FIX-M258-iter03-guard-scans-its-own-scratch`** (net-new) — `test_decommissioned_instruction_guard`
  walks `demo-stack/stacks/**`, which `demo-stack/.gitignore:8` ignores, and reports the *platform's*
  source inside a demo's ephemeral clone as a rext named-consumer list. **Proven pre-existing** by
  running it in the pristine clone at `fast-build-m257-close`: identical 2 failures. Fires on any box
  that has ever run a demo. Plausibly the same root as M257's *"the stack-core sweep did not complete."*
- Then, unchanged: wire the batch-gate at `up-injected.sh:2810` under `D-v28-3` semantics → land the
  world-contract restore leg (b) → the composed 3× cold campaign, **spread published beside the p50**.

</details>

- **`FIX-M258-iter03-guard-scans-its-own-scratch`** — **still open**, re-confirmed pre-existing at
  iter-06 (identical 2 failures). iter-06 adds a **second member of the same family**:
  `test_fence_provenance::test_the_escape_accepts_and_records`, whose two RED guards (`dev_flag_guard`,
  `demo_knob_guard`) were run against a **pristine `HEAD` extract** and are RED there too. Both tests
  fire only on a box that has ever run a demo — which is why they surface here and skip in a clean
  checkout, and why they keep being mistaken for regressions.
- ⚠️ `demo-2` (11 containers) and the 5-container dev stack are the **user's**: do not tear down,
  re-seed, restart or reset either. `demo-1` is left UP, **restored and presenter-usable** (verified at
  iter-06: 4 story orgs, 591 users, 12 cockpit seats all resolving in a 35-identity roster).
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
