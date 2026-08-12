# M258 — progress

> ## ⚠️ STATUS — ACHIEVED BY USER RULING (2026-08-12), NOT `gate-met`
>
> The user ruled M258 has achieved its goal, accepting it on clauses **1, 2, 4, 5** plus the **~402 s
> clean projection**, having concluded the CPU contention on this box is not something he can remove.
> Record it as **achieved by user ruling, timing clause unmeasured under load** — the shape of M257x's
> `TOK-09`. **Clause 3 is NOT met and must never be recorded as met.** Specifically, and not to be
> softened: the **840.01 s** figure stays **instrument-rejected** (3/3 `headroom=FAIL`), **401.60 s is
> a PROJECTION** never measured as one cycle, and no clean p50 over 3 cold cycles has ever been taken.
> The auto-arm stays armed; clause 3 is now an opportunistic bonus. (`iter-10/decisions.md` `D52`.)
>
> **Remaining scope, in order (user ruling):** ~5 iters of build-time fruit → **then fire `TOK-02`, a
> new goal: SPACE optimisation** (`iter-11/decisions.md` `D57`) → **end state `END-M258-one-stack`:
> exactly ONE stack up, built with the new mechanism from the newest platform repos.** Currently three
> (`demo-1`, `demo-2`, dev). ⚠️ **Build-and-verify the new stack FIRST, then tear the others down** —
> never teardown-first; this is the one sanctioned exception to "never touch `demo-2`", and `demo-2` is
> *not* the one to keep (it is on **pre-L1** images). **Space must never be bought with time** (`D58`):
> no `-af`, `--filter until=24h` only, every cache policy argued on both axes with measurements.

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

- iter-07 (tik, `closed-fixed-partial`): **the campaign's last blocker discharged BY MEASUREMENT, not by
  a campaign** — and then the host refused. `D27`: iter-06's `green:false` is proven to be **clone
  topology alone**, from two path expressions (`readiness.sh:65,71` — `_ry` = `<rext>/../platform/
  repos.yml`, absent for the authoring clone, present for the consumption one) plus one query (demo-1
  has all four derived schemas `extensions/sentinel/public/directus`), so the probe will **assert and
  pass** and `green` flips `false → true` for topology reasons alone. Consumption clone re-pinned
  `iter-03 → iter-06`; the gap was **exactly the feature under test** (all four batch-gate files absent,
  no hook) — the M236 shape. **4 of the gate's 5 clauses proven** (1 batch-drives ✅ · 2 red set EMPTY ✅ ·
  4 **0 platform edits PASS**, measured across all six peer clones ✅ · 5 **12/12 cockpit seats resolve**
  in the 35-identity roster ✅); **clause 3 — the only TIMING clause — NOT TAKEN.** `load1` minimum
  **11.93** vs a limit of **10** across ~30 min of polling, trending **up to 62.88**, saturated by a
  **different user project's** parallel campaign. **Deliberately not run as an operator** (`D28`): the
  suite is `retries: 0` by contract, so a run at load 40–60 manufactures **false reds** that `D-v28-3`
  escalates to the user as if they described the product. `demo-1` **never torn down** — still
  presenter-usable. — see `iter-07/progress.md`
- iter-08 (tik, `closed-fixed-partial`): **THE COMPOSED 3× COLD CAMPAIGN RAN** — the wait was ended by
  **automating the trigger**, not by the box getting quieter. iter-07 hand-polled 30 min at ~2-minute
  granularity and correctly found no window; `autoarm-campaign.sh` (15 s sampling, fire on 3 consecutive
  `load1 ≤ 5.0`, safe-by-refusal so it re-arms) **armed 08:17:41Z and fired 08:19:12Z — 91 s later**,
  into a dip below 5.0 that lasted **75 s** (`D29`). *A hand-sampled trigger cannot catch a window
  shorter than its own interval.* Campaign 08:19:19Z → 09:02:12Z: **3/3 reps `up_rc=0 green=True
  isolation=OK phases=complete`, `red_count 0 · passing 30 · failing 0` each time — and 3/3
  `headroom=FAIL`** (`peak_load1` 40.09/74.77/51.80 vs 10), because anima8's next batch resumed **60 s
  after launch** (`D30`). So **clause 3 stays NOT MET** — total p50 **840.01 s [811.06–859.06]** is the
  p50 of reps buildbench itself calls *"not usable measurements"* (`D35`); the **spread is 1.06×**,
  against the 2.04× M256 escalated as undecidable. **Clauses 1, 2, 5 are now proven on THREE consecutive
  cold cycles** rather than one — clause 5 (**12/12 cockpit seats**) verified *after* the campaign, so the
  world-contract restore leg held 3× running. `D27`'s prediction **held**: `autoverify` `green:true
  warnings:0` on all three reps (`D31`). `D17` **passes in both directions** — `autoverify` 5.32 s vs
  `batch_gate` 153.81 s attribute *separately* **and** Σ sub-phases = `P4_BRINGUP` exactly, with
  `batch-gate.json`'s own `142+11=153` matching the anchor-derived 153.81 s (`D32`). Walked past a live
  stale-artifact trap: the authoring clone still held iter-06's `batch-gate.json` reading *identical*
  `verdict green · red_count 0` — only the `ts` betrays it (`D33`). **The iter's real deliverable was a
  scare and its settlement:** `set_dress` measured **283.53 s vs iter-05's 81.23 s** (3.49× vs a 2.05×
  cohort), which would have put the clean total at ~600 s and made L5 urgent — **refuted by two cheap
  checks** (`D38`): all 3 reps do identical work (same digest `b4cb55bcee08 → ea2e187a1605`), and
  `git diff --name-only iter-03..iter-06` filtered for `setdress|snapshot|stacksnap` returns **zero
  files**, so the code path is **byte-identical**. Excess is environmental; **projection 401.60 s vs 480
  holds**; L5 stays a reserve against a ~81 s phase. Re-scope graded **n** despite 840 > 600 (`D37`):
  reps instrument-rejected, *and* the remedy (split the suite) misses the diagnosis — `batch_gate` p50
  **179.37 s** is inside M256's 200 s budget *while contended*. — see `iter-08/progress.md`

- iter-09 (tik, `closed-fixed`): **`LEVER-M257-L5-setdress` HAS A TARGET** —
  `FIX-M258-iter08-set-dress-has-no-internal-attribution` discharged, and **settled retroactively on
  logs we already had, at zero host cost, on a box that never dropped below `load1` 19.72** (`D41`).
  The 252.73 s span was **one operation, not two**: `sd_replay_taxonomy` is **249.69 / 258.63 /
  266.66 s = 87–91 %** of `set_dress` across iter-08's three cold reps, every rep tiling the parent
  with **residual 0.0**, everything else in the phase under 16 s (`D40`). The lever's real target is
  a single `stacksnap replay --surface taxonomy` moving **~1.47 GB** (`skill_embeddings` 825 MB +
  `job_role_embeddings` 364 MB) and rebuilding **two** pgvector indexes. ⚠️ Those seconds are
  **contended**; the durable finding is the **share**, which prices the replay at ~70 s against
  iter-05's quiet phase. Built **NESTED, not four more flat anchors** (`D39`) — children are not
  siblings: that would double-count against `P4_BRINGUP` *and* redefine the `set_dress` series
  `D38` had just used to settle the release's biggest scare; verified both ways (Σ `sub_phases`
  still exact at 802.39/822.19/840.91, `set_dress` still 283.53). Level two landed too: `replay.Run`
  now attributes verify/clear/**copy**/**reindex**/sequences per table with an **explicit**
  unattributed residual (`D45`) — the copy-vs-reindex number L5 needs and no captured log can give.
  **The thesis was made mechanical** (`D42`): a mutant billing the reindex to the copy leaves the
  **sum test passing** and fails only the attribution test. A second mutant **SURVIVED** — a *test*
  gap, not a code gap (`D43`) — and was fenced. Net-new finding: **the literal ratchets are polluted
  by the demo stack dir** (DOCSTRING +10, TEST_MODULE +9 from `stacks/demo-1/clones/app/studio/**`),
  a **third** consumer of the `guard-scans-its-own-scratch` family, so M257's recorded 254/663 may
  itself be polluted (`D44`); own contribution paid to **+0 on all three**, never waived. Tagged
  `fast-build-m258-iter-09`, **verified on origin**, consumption clone re-pinned with the
  feature-present check. **Clause 3 armed, not awaited** (`D47`) against a **fresh** output dir.
  — see `iter-09/progress.md`

- iter-11 (tik, `closed-fixed`): **THE SPACE AXIS HAS ITS FIRST MEASUREMENT — and post-teardown was
  indeed the defect.** **178 of 184 volumes dangling, 5.297 GB, 100 % reclaimable**; reclaimed after
  proving **OVERLAP 0** against the six in-use volumes (which resolve to `demo-1`/`demo-2`/dev
  postgres — both of the user's stacks), all three stacks verified resident afterwards (`D54`).
  **Producer named:** the bitnami Postgres image declares three `VOLUME`s and compose binds only
  `/bitnami/postgresql`, so **every container start mints two anonymous volumes**
  (`/docker-entrypoint-{pre,}initdb.d`) that a non-`-v` teardown or a container recreate orphans —
  three stacks × five days = 178 (`D55`). **`--purge` is EXONERATED by measurement**, not assumed: it
  runs `down -v --remove-orphans` and today's 3-rep campaign produced **zero** orphans while the newest
  orphan predates it by five hours (`D56`). Two traps recorded so nobody re-derives them wrong:
  **`docker images` SIZE overstates reclaimable ~5×** — the four `m257-*:probe` leftovers read 8.88 GB
  but share 5 of 10 layers with an in-use image, and `system df` says **1.754 GB** (`D53`); and
  **host-side stack dirs are invisible to `docker system df`** — **4.2 GB** under
  `stack-demo/…/stacks/` plus an orphan `demo-4/`, which upgrades
  `ROUTE-M258-iter02-purge-did-not-clear-the-stack-dir` from a curiosity to a space finding. Build
  cache **deliberately NOT pruned** (19.28 GB reclaimable) — space must not be bought with time
  (`D58`). — see `iter-11/progress.md`

- iter-12 (**tok**, bootstrap flavor — **user-directed**, `closed-fixed`): **`TOK-02` authored — space
  partitioned by its COUPLING TO TIME**, which is what turns the user's constraint from a veto into a
  sort (`D60`): **Class A** zero-coupling (orphaned volumes, dead images, stack dirs, *a stack that is
  going away anyway*) — take all of it; **Class B** *favourable* coupling — **image size**, because
  export/unpack is size-proportional at a measured **5.73–8.05 s/GB** on this host class, so a smaller
  image is a **faster** build (L1 proved it: next-web **4.04 GB → 417 MB AND 114.7 s → 53.31 s**, one
  lever); **Class C** adverse — the **build cache**, 19.28 GB and the largest single reclaim on the box,
  **out of bounds by default** at a measured 173 s per 356.8 MB evicted. Two unmeasured axes opened.
  **The host axis (`D59`): `Docker.raw` is allocated 50.68 GB against a docker-logical 51.56 GB — 1.7 %
  agreement — so in-VM reclaim really does return SSD**; this had to be measured, because on a host whose
  sparse file does not TRIM every reclaim figure in this milestone would have been fiction. Total
  M258-attributable disk **≈ 63.5 GB** (50.68 Docker + 12.8 host tree). **studio-desk attributed to a
  layer:** single-stage Dockerfile, `npm ci` = **1.04 GB = 61 % of the image**, shipped to produce
  **63.2 MB** of output and run a bare `node dist/index.js` (30 prod deps vs **33 dev**). **The
  constraint changed the fix** (`D61`): prune-and-copy, **never re-install** — the naive runner's
  `npm ci --omit=dev` buys space with time, and L1 escaped that only because `next build` *emits*
  standalone. **And it trimmed the claim** (`D62`): studio-desk's time prize is **≈ 7–10 s**, not
  iter-11's 115 s — the other ~105 s is `npm ci` + `tsc` + `vite build`, which no runtime shape removes.
  Third trap sibling recorded: **`ls -l Docker.raw` reads 137.44 GB, 2.7× the truth.** `demo-4/` orphan
  corrected to **8 KB — a hygiene item, not a space finding** (`D63`). — see `iter-12/progress.md`

- iter-13 (tik, `closed-fixed`): **`TIK-A` — studio-desk stops being the demo's largest UI image:
  1.7 GB → 1.35 GB (350 MB/stack, 20.6 %)**, via an **rext-owned multi-stage prune-and-copy** Dockerfile
  in the shape L1 sanctioned — clone as build CONTEXT only, **zero platform-repo edits**, not even a
  demopatch — **verified booting and serving identically to the live `demo-1` control** (302/302/302 on
  `/`, `/home`, `/assets`; minted pk still greppable at `/app/dist`, so `buildbench`'s ISOLATION probe
  keeps its contract). **The headline is one-THIRD of what was predicted, and the refutation is the more
  useful result** (`D64`): **838 MB of the 1.04 GB dependency layer survives `--omit=dev`** because
  `@clerk/clerk-js` carries a crypto-wallet tree (`viem` 68.2 · `@solana` 20.6 · `ox` 9.2 · `@base-org`
  8.2 MB) plus React-Native/Hermes as **runtime** deps — **studio-desk's image is dominated by PRODUCTION
  dependencies, not by the toolchain**, which was only ~20 % of it. Three things this iter is worth
  reading for: **(1) the anti-silence assert fired on a prune that had WORKED** — `typescript` is a real
  production transitive (`@clerk/clerk-js → @base-org/account → ox → abitype`), so the *sentinel* was
  wrong, not the prune (`D65`; corrected to `vite`, which `npm ls --omit=dev` proves empty). **(2) The
  Dockerfile had to go INTO the cache key or the lever would have shipped NOTHING** — the reuse check
  keys on endpoint + patchset, **neither of which moves when the Dockerfile changes**, so every box
  holding a single-stage image would have reused it under a green log (`D66`; the "applied is not
  shipped" class in a new costume). **(3) The TIME half is WITHDRAWN, not claimed** (`D67`): `D62`'s
  7–10 s came from **5.73–8.05 s/GB measured on `billion` (x86_64/containerd)** applied to an
  **arm64/overlayfs** host — the exact cross-host error `build-budget.md`'s opening rule exists to
  prevent — and the available logs compare a **cold** export (33.2 s) against a **warm** one (9.6 s
  unpack), so they settle nothing. `TIK-C`'s cold bring-up yields it for free. Suite **9 failed / 1080
  passed**, and the nine were **PROVEN pre-existing** (`D68`) by re-running them from a `git archive HEAD`
  extract at the same directory depth: **the same nine by name**, 9 failed / 131 passed.
  — see `iter-13/progress.md`

- iter-14 (tik, `closed-fixed`): **`TIK-B` — the orphaned-volume leak stopped AT ITS PRODUCER, at zero
  time cost.** `--purge` was already innocent (`D56`); the **plain `down`** was the leak — it passed no
  `-v`, so the two undeclared bitnami anonymous volumes outlived every non-purge teardown and every
  container recreate (178 volumes / 5.297 GB over five days). Now `down -v`. **The one-flag fix still
  needed a measurement** (`D69`): `-v` also removes NAMED volumes, which would make a plain teardown
  destructive — so a live census of every container in the project was taken, and the **only volume-type
  mounts in an entire demo stack are those two anonymous ones**. Zero named volumes to lose, and the
  database is a **host bind mount** `-v` never touches. Fenced by
  `test_down_plain_removes_anonymous_volumes`, which asserts both branches pass `-v`, that no bare
  `down ||` survives, **and the rationale sentence** — so a future named volume re-opens the decision
  rather than silently making teardown destructive. Second half **priced and deliberately NOT taken**
  (`D70`): `purge_data_dir` is scoped to `data/` alone, so **≈276 MB per stack survives a full `--purge`**
  (`clones/` 220 MB · `bin/` 37 MB · fake-Clerk 18.5 MB) — widening a `rm -rf` whose safety rests on a
  G1 path-assert, minutes before the milestone's binding end state, is the wrong trade; `TIK-C` measures
  it for free. — see `iter-14/progress.md`

- iter-15 (tik, `closed-fixed`): **`END-M258-one-stack` ACHIEVED — exactly ONE stack up (`demo-3`),
  built with the new mechanism from the newest platform mains** (`platform` `766df6c` · `app`
  `c52dbc51e` **+76** · `next-web-app` `3379072e9` **+59** · `ant-academy` `7ae25e95`), in the
  **mandatory order**: build-and-verify first, teardown last, and the **user's own stack torn down LAST
  of all** — enforced *in code* (`teardown-others.sh` refuses unless demo-3 is up, and re-checks between
  every step), not by discipline. Survivor verified presenter-usable: **4 orgs / 591 users / 42,790
  skills / 12 of 12 cockpit seats** across 4 stories, cockpit 200 · web 307 · studio 302 · backend
  health 200, UI tier **2.15 GB vs demo-2's pre-L1 9.68 GB**. **SPACE: 11.54 GB of real SSD reclaimed
  at ZERO build-time cost** (Docker.raw **53.84 → 42.30 GB**, images 25.71 → 14.02, free **170 →
  182 GiB**) — **with the 21.03 GB reclaimable build cache deliberately UNTOUCHED**, which is `TOK-02`'s
  constraint honoured rather than quoted. **`D70` validated to within 1 MB** (predicted ≈276 MB survives
  `--purge`; measured **277 MB**), and **`D59` held again** (host file fell 11.54 GB against an in-VM
  11.69 GB — real SSD, not VM bookkeeping). **The batch gate returned `verdict: red`, `red_count: 15`
  — ESCALATED, not swept — and the reading that decided the milestone is `D74`: it grades `pt-world`,
  the DECOUPLED TEST SEED, not the presenter world** the restore leg rebuilt and that was measured
  healthy. Its causes are **left unresolved and labelled**: 4 plain timeouts at `load1` 26–33 with
  `retries: 0` (`D28`'s false-red condition) + 11 data-shape assertions agreeing with autoverify's
  under-set-dress warning, and **no `SQLSTATE 42P01` anywhere**, so the "newest platform moved a table"
  story is **unproven and must not be reported as diagnosed**. Three more findings: the **rext pin FATAL
  fired on a half-completed re-pin** and *both* remedies it offered were wrong — `rext.tag` is an intent,
  not a lock (`D71`); the **clause-3 waiter was disarmed** before the transition because a firing
  campaign tears `demo-1` down and could have left the user stackless (`D72`); and **`compose down`
  FAILED on both demos** (*"service sentinel has neither an image nor a build context"*) — **the label
  sweep is what actually removed all 22 containers**. ⚠️ **`SETTLE-M258-iter13-studio-desk-cold-time` did
  NOT settle** (`D75`): BuildKit reused iter-13's probe layers, so the "free" cold number was a cache hit
  (1.5 s export). The studio-desk TIME axis stays **UNMEASURED**; only the 350 MB space win is measured.
  — see `iter-15/progress.md`

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
- ✅ **iter-07 discharged every PRECONDITION of `TOK-01` step 4** — the tag is on origin, the consumption
  clone is re-pinned to `fast-build-m258-iter-06` and carries the gate, and the `postgres-schemas`
  refusal is **proven satisfiable there** (`D27`). It also graded **4 of the gate's 5 clauses green**.
- ✅ **iter-09 discharged `FIX-M258-iter08-set-dress-has-no-internal-attribution`** — and with it the
  precondition on `LEVER-M257-L5-setdress`, which now has a named target (the taxonomy replay, ~88 %
  of `set_dress`) instead of an opaque span. The lever itself is **still unspent and still not
  needed**: the 401.60 s projection sits inside 480.
- **iter-10 (tik, under `TOK-01`) — step 4, unchanged: a GATEABLE composed campaign.** The waiter is
  **armed** at the `fast-build-m258-iter-09` pin against `campaign-iter09/`, so its ledgers will
  carry the set-dress attribution live. ⚠️ **Use a FRESH output dir per campaign** — `build_report`
  globs `rep-*/ledger.json`, so re-using a previous campaign's dir silently aggregates its reps into
  the new p50 (`D47`). If a clean p50 lands **over** 480, L5's price is now known well enough to
  spend it; if it lands under, L5 stays unspent.
- **`SPLIT-M258-iter09-copy-vs-reindex`** (net-new) — the level-two instrument shipped but is
  **unmeasured**: no run has yet produced a `replay "taxonomy" phase costs:` line. The first campaign
  at this pin yields it for free. Until then, **do not assert whether the taxonomy replay is
  COPY-bound or REINDEX-bound** — the remedies differ and the guess is not evidence.
- **`ROUTE-M258-iter09-literal-ratchets-scan-the-demo-clone`** (net-new, `D44`) — the three
  `derivation_registry` literal censuses `rglob("*.py")` from the repo root and so census
  `demo-stack/stacks/demo-N/clones/app/studio/**`, the **platform's** source inside a demo's
  ephemeral clone. Third consumer of `FIX-M258-iter03-guard-scans-its-own-scratch`'s root cause;
  the shared fix is a root-selection change. **Consequence:** any ratchet figure measured on a box
  that has run a demo, without excluding `stacks/`, is not a measurement of this repo.
- <details><summary>superseded routing (iter-08's plan, kept for the audit trail)</summary>

- **iter-08 (tik, under `TOK-01`) — step 4: the composed 3× cold campaign. NOTHING IS LEFT TO BUILD.**
  Run `.agentspace/scratch/work-m258/launch-iter07-campaign.sh` — one command; it asserts the user's
  stacks resident and **refuses** otherwise, asserts headroom **before** the teardown, tears `demo-1`
  down from its **owning** (authoring) clone, then runs 3 reps from the consumption clone.
  ⚠️ **The ONLY remaining input is ~30 minutes of a host at `load1 < 10`.** iter-07 watched for ~30 min
  and the minimum was **11.93**, trending to **62.88**, because a *different user project*
  (`hyperspace/anima8`) is running a parallel campaign on this box. **Do not route around this by
  running the batch contended** — `retries: 0` turns browser timeouts into a red set that `D-v28-3`
  escalates to the user as a product verdict (`D28`). Publish the **spread beside the p50** (`C2`) on
  both halves — the batch has 129 s (iter-04) and 160 s (iter-06) and still no p50.

</details>

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
- **`ROUTE-M258-iter07-demopatch-G5-does-not-revert-the-native-clone`** (net-new, iter-07 `D26`) —
  `stack-demo/ant-academy` carries **4 modified tracked files**, all self-identifying demo-patches.
  Structural, not a per-patch bug: `demopatch`'s apply-then-revert assumes an **ephemeral build-scratch
  clone** the image build discards, but **ant-academy runs NATIVELY** and is never imaged, so its clone
  is the durable peer clone and G5 has nothing to throw away. Matters because **G2 is drift-refuse** — a
  later patch baselined against a pristine file can read DRIFTED against a clone an earlier run patched
  (the silent-refusal class that shipped a 76 s members grid for four releases). **Not repaired:**
  reverting tracked files is a forbidden op, *and* this clone is what the **user's demo-2** serves from.
- ⚠️ `demo-2` (11 containers) and the 5-container dev stack are the **user's**: do not tear down,
  re-seed, restart or reset either. **⚠️ NOTE (iter-07 `D23`): `demo-2` is owned by the CONSUMPTION
  clone** (`stack-demo/rosetta-extensions`) while **`demo-1` is owned by the AUTHORING clone** —
  resolved from the docker mount, not from any script's location. Re-pinning the consumption clone is
  therefore an operation *near* the user's stack; it is safe only because `demo-stack/.gitignore:8`
  ignores `stacks/` (verified with `git check-ignore -v`). `demo-1` is left UP, **restored and
  presenter-usable** (verified again at iter-07: 12 cockpit seats all resolving in a 35-identity
  roster; never torn down).
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

- iter-16 (tik): **the 15-red batch verdict attributed AND fixed — one dangling reload, not fifteen
  product failures.** Platform `766df6c` folded **sentinel into `app`** (v11.0, the 8th merge); our
  three post-seed reload sites still drove the deleted container's RPC and logged the miss as
  *"non-fatal — a non-AI-sim run is unaffected"*. A stale casbin enforcer refuses **every** org-scoped
  read and write with `forbidden` at HTTP 200 (the silent-403 class) — 15 of 31 Playthroughs **and both
  negative controls**. Contention, a partial seed, and "a table moved" are all **refuted**, from
  artifacts already captured. Fixed by publishing to `sentinel:policy:invalidate` (app's own channel),
  **proven live on demo-3 before the code was written**; pre-v11.0 RPC retained beneath. 0 platform
  edits, 0 fences edited. rext `fcdc651` / tag `fast-build-m258-iter-16` **on origin**. —
  see iter-16/progress.md
- iter-17 (tik): **the fix PROVEN END-TO-END — 15 reds → 0, cold, on a fresh `demo-4` from the newest
  platform mains.** `red_set: []` · `runner_exit: 0` · 30/31 passing (the 1 is the declared TODO) ·
  `215 passed` · `autoverify green: true, warnings: 0` · *"UP, and every journey verified."* All three
  invalidation sites fired; **both previously-failing negative controls pass** — the tests that separate
  *correctly isolated* from *uniformly blind*. `batch_seconds` **629 → 129** (the old batch was slow
  BECAUSE it was broken). The dev half asked six pieces three questions each and found **one real gap**:
  the anonymous-volume leak fixed for demo at iter-14 was **never carried to dev**, which runs the same
  image through the same compose — fixed, fenced four ways, 159 dev tests green. Recorded not changed:
  **`--public-host` is default-on and turns the batch gate OFF on its own host**, and
  **`FIX-M258-iter15-hiring-under-set-dressed` does NOT reproduce** (50 rows vs 38) → re-scoped to WATCH.
  Converged to one stack; **`demo-3` untouched throughout**. rext `fast-build-m258-iter-17` on origin. —
  see iter-17/progress.md
