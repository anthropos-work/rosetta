# Hardening Ledger — M257 first-light build

The milestone ran **9 iters (7 tiks + 2 toks)** and reached its exit gate without ever being hardened, so
the first pass below is cumulative over the whole milestone rather than incremental over a batch.

**Scope of `--final` here spans TWO trees.** M257's code lives in `rosetta-extensions` (9 commits,
`06761b5^..HEAD`); the `rosetta` side is the corpus + plan. Both are in scope, and both carry harden
commits.

---

## Pass 1 — 2026-08-12 — final

**Iters hardened this pass:** all milestone-touched code (iter-02 … iter-09, cumulative)
**Tiks covered since prior pass:** all iters in milestone — this is the milestone's first harden pass

**Scope manifest (cumulative), by subsystem:**

| subsystem | files | touched by |
|---|---|---|
| `stack-core` (the gate instrument) | `buildbench.py`, `anchor_construct_guard.py`, `derivation_registry.py`, `claim_census_guard.py`, `suite_census.py`, `hostprofiles/{macmini,billion,laptop}.json` | iters 04, 06, 07, 08, 09 |
| `demo-stack` (L1 + blocker B2) | `frontend/next-web.Dockerfile` (net-new), `frontend/hiring.Dockerfile`, `up-injected.sh`, `ensure-clones.sh`, `lib/studio.sh` | iters 03, 09 |
| `stack-seeding` (blocker B1, 34 sites / 20 files) | 20 seeder + test files | iter 03 |
| `stack-verify` | `live/autoverify.sh`, `tests/test_verify.py` | iters 02, 03 |
| `stack-injection`, `playthroughs` | `gen_injected_override.py`, manifest/seed yaml | iter 03 |
| `rosetta` corpus | 11 docs | iters 04–09 |

**Cross-iter integration findings (final mode's defining work):**

- **`buildbench.py` is the file three iters edited independently** (06 headroom units, 08 the baseline, 09
  ISOLATION), and the two falsifiable gate clauses it now carries meet only in `rep_is_ok`. Checked: both
  are read there, `headroom` unconditionally and `isolation` present-or-absent with the `absent ≠ failing`
  rule the identity rollup already uses. No seam.
- **`hostprofiles/*.json` were written by three iters against a loader iter-06 made stricter.** Verified the
  kind-keyed rule holds across all three shipped profiles: the two `docker-desktop-vm` profiles declare
  `host_logical_cores`, `billion` is `native-linux` where `cores` is the correct basis by construction, and
  `test_every_shipped_profile_can_actually_grade_clause_one` already ranges over the family, so a fourth
  profile cannot be added without one.
- **iter-09's ISOLATION assert is the net under a fix iter-09 itself deferred.** The `.env` finding below
  routes to a `.dockerignore` repair whose obvious form re-creates the M218 iter-03 incident (a bundle built
  without the minted-key overlay) — and that is precisely the `foreign_pk` arm. The lever and the assert that
  guards it were landed together, and this is the first evidence that the pairing pays.

**Coverage delta on touched files:**

Reported as controls-on-the-surface rather than as a line percentage, because the gap this pass found was
not un-run lines but an un-tested **boundary**: `isolation_assert`'s probes are injected so the logic tests
run without Docker, and the two live probes therefore had **zero** direct tests. Measured at `8956e69`, the
milestone's last iter commit: `_image_env` / `_image_bundle_pks` occur **1×** in the whole `stack-core/tests`
tree (a docstring mention). After: **13×**, across 10 dedicated controls. Both fail-opens lived exactly
there.

- `stack-core/buildbench.py`: **77 %** stmts under its own four test files (195 tests, 11 subtests)
- `stack-core/anchor_construct_guard.py`: 50 % under the same subset (its main denominator suite not in
  this run — it is a 54 s module and was run separately, green)
- `stack-core/suite_census.py`: 23 % under the same subset (likewise — it has consumers in ~10 other files)
- test count on touched files: **+29** (isolation 17 → 34, frontend-build 103 → 105, anchor-ladder 14 → 19,
  census-collection 17 → 22)

**Tests added:**

- iter-09 → `stack-core/tests/test_isolation_assert_m257.py`: **+17** — 3 fail-open controls on the
  composed assert (including the exact pre-fix conjunction that returned `ok=True` with zero failures),
  4 probe-level controls on `_image_env`, 6 on `_image_bundle_pks`'s grep capability, and
  `test_THE_REAL_SHELL_TEXT_fails_closed_against_a_broken_grep`, which runs the probe's **actual emitted
  shell text** against a grep stubbed to exit 2
- iter-09 → `demo-stack/tests/test_frontend_build.py`: **+2** — the standalone assert proven to *fail a
  build* rather than to *appear in one* (extracted `RUN` line, executed under `sh`, rc=1 required), and the
  `.dockerignore` pairing rule with `.env*` as the declared exception carrying its routed handler
- iter-08 → `stack-core/tests/test_anchor_ladder_alternates_m257.py`: **+5** — `TheContentDriftBlindSpot`,
  which pins the fence's largest known limit executably and goes RED when the handler lands
- (net-new surface) → `stack-core/tests/test_suite_census_collection.py`: **+5** — the ephemeral-clone
  scope rule, live and on a synthetic tree

**Bugs surfaced + fixed inline:**

- **the ISOLATION assert's THIRD fail-open** — `_image_env` returned `[]` for both "no env" and "the inspect
  did not run", and the assert *derives* "is this a UI image" from that env, so an empty env granted the
  "backend image, no bundle expected" excuse and **disarmed the fail-closed `unreadable_bundle` arm**.
  Measured pre-fix: a `demo-1-next-web` with both probes failing returned `ok=True`, zero failures. Both
  probes shell out through one `_run` that returns rc=127 on any exception, so the conjunction is what one
  daemon timeout does to every image at once (commit `7cf2768`)
- **the bundle probe could not tell "grep ran" from "grep found nothing"** — `__SCANNED__` is echoed *before*
  the greps, so it proves the build-output directory exists and nothing more. iter-09's busybox `--exclude`
  fix was the `None`-vs-`[]` return (real) plus a *discipline*. Re-measured with a stubbed grep and a real
  key planted in the scan root: clean verdict. Now a positive control planted in-container and grepped with
  the identical flag set (`7cf2768`)
- **five INVERTED range citations, with the fence already RED on the live tree** — `anchor_construct_guard`
  exited 1 and two of its tests were failing before this pass. `594-589`, `1135-1126` ×2, `2230-2219`,
  `1087-1085`: the tail of iter-09 D5, where the repair re-derived single-line anchors and left ranges with
  one number remapped and one not. Un-reversing them does not work — checked at the pinned ref, both orders
  are wrong — so each span was re-derived from what its own sentence claims (`b20f131b`, rosetta)
- **the suite census counted a demo's persistent platform clone as this repo's tests** — six RED tests, none
  about this repo. Every walk used a bare `rglob` into `demo-stack/stacks/**`; the Go arm was 100 %
  contaminated. `test_dropped_mirror_fence.py` had already declared and excluded that path for that reason,
  so the rule was in the codebase and applied in one walk and not the others. Now one shared predicate
  across all four (`ca9baff`)
- **a routed handler whose description sends its fixer to the wrong repo** —
  `FIX-M257-committed-env-ships-real-clerk-pk` calls `apps/web/.env` a *committed* file. It is untracked
  (`git ls-files 'apps/*/.env*'` returns only `.env.example`), so a fixer looks in the platform repo, finds
  nothing tracked, and concludes the finding was mistaken. Corrected with the real mechanism (`ce1a805`)
- **the anchor fence's largest blind spot was recorded only in one iter's `decisions.md`** — its module
  docstring carries a careful "what this does not catch" section naming #17, so an auditor saw a
  complete-looking list of limits with the newest and largest absent (`b364172`)

**Routed forward (not fixable inline):**

- **`FIX-M257-dockerignore-env-pattern-unpaired`** — `next-web.dockerignore` excludes `.env*`, Docker matches
  from the context root, and that rule covers `./.env` and nothing nested, while every sibling rule is paired
  with a `**/` twin. Consequence, measured on the real post-L1 image: `/app/apps/web/.env` at 19,087 bytes
  carrying the real-Clerk pk and `CLERK_SECRET_KEY`, no `.env.local` beside it, riding `.next/standalone`
  into the runner where `server.js` calls `loadEnvConfig` at boot. Masked for the Clerk vars by the injected
  override plus `@next/env`'s never-overwrite rule; the residual is the set difference. **Not fixed here
  because the tidy fix is a trap:** `**/.env*` also drops the minted-key overlay and bakes the real key
  (M218 iter-03). Needs a re-include and a real build to validate.
- **`FIX-M257-anchor-guard-content-drift`** (pre-existing, re-affirmed) — a new detection mode on a
  pre-commit fence grading 1,612 citations tree-wide, late in a milestone. Now pinned executably instead.

**Flakes stabilized:** none observed. No test in scope failed intermittently across the pass's re-runs.

**Knowledge backfill:**

- `corpus/ops/demo/frontend-tier.md` — what the L1 image actually contains (the `.env` finding with its
  mechanism, its masking, and the trap in its obvious fix), and the standalone assert now being proven
  behaviourally rather than by string presence
- `stack-core/anchor_construct_guard.py` module docstring — the content-drift measurement, why the class is
  deferred rather than dropped (it *is* mechanically detectable, unlike #17), and the sentence a reader
  needs: read this guard's green as *"every resolvable anchor points at something"*, never as *"the
  citations are correct"*
- `stack-core/buildbench.py` — the corrected handler description and the `.dockerignore` mechanism, at the
  probe whose scope decision the finding tests

**Stop condition:** continue-to-next-pass — the whole-tree sweep that grades the cumulative scope has not
returned yet, so the milestone's failure roster is unattributed and the dimension scan is incomplete.

---

## Pass 2 — 2026-08-12 — final

**Iters hardened this pass:** all milestone-touched code (cumulative, continued from Pass 1)
**Tiks covered since prior pass:** n/a — same cumulative scope, second sweep over it

**Coverage delta on touched files:** +6 tests (isolation 34 → 40). The delta is small by count and that is
the point: Pass 1 took the surface from *"the assert's logic is fenced, its edge with the world is not"* to
fenced-on-both-sides, and Pass 2 found the two remaining seams on that same edge rather than new territory.

**Tests added:**

- `stack-core/tests/test_isolation_assert_m257.py`: **+6** — five on `_stack_minted_pk` (the third live
  probe, and the one that decides what *foreign* means) and one static fence on the ledger emit

**Bugs surfaced + fixed inline:**

- **a rendered FINGERPRINT was adoptable as this stack's own key.** `pk_fingerprint` renders a key as
  `pk_test_MTI3…61fbfaf4` and `_PK_RE` matches the twelve-character prefix inside it; the discovery's
  fallback globs every `*.json` in the stack dir. The failure mode is the one the fail-closed arms do NOT
  cover: a MISSING key refuses to grade, a WRONG key reds every image against a truncation nobody can look
  up. **Wrong-and-loud is not fail-closed.** Fixed by skipping matches followed by the fingerprint
  separator, through one `_first_key` used by both arms (`45d8d1f`, sharpened at `a9869e7`)
- **the one tolerated `absent` in `rep_is_ok` was a promise, not a check.** A missing `isolation` block
  reads as consent — correctly, so iter-08's baseline is not retro-failed — and its safety rests entirely on
  the live path always emitting it, which was asserted in prose. Move the emit behind a condition and a rep
  that SKIPPED the assert scores like a rep that PASSED it, with nothing going red. Now a static fence
  requires the emit in the same dict-merge as `headroom` (`8d90b20`)
- **this pass's own prose breached both literal ratchets** — DOCSTRING 254 → 256, TEST_MODULE 663 → 667,
  because the findings were written up with their measured counts inside docstrings. The growth was
  **deleted**, not ceilinged: both are back at exactly iter-09's closing reading, with the counts moved to
  this ledger. The pre-existing breach (+14 / +10) is untouched (`474dff7`)

**Cross-iter / whole-tree verification:**

- **the gate is NOT retracted by any of this, and that was checked rather than assumed.** The three gate-met
  reps were re-graded with the post-harden code: `rep_is_ok` True ×3, campaign `ok=True`, identity `match`.
  Their ISOLATION blocks were real measurements, not fail-opens — 8 images checked per rep, a real
  `own_pk_fingerprint`, empty `foreign_pks`/`foreign_origins`. None of the four fail-opens closed this pass
  was active in them. The headline verifies against the raw ledgers exactly: p50 286.99 / min 280.99 /
  max 303.44, `green:true` + 0 warnings + HEADROOM OK + ISOLATION OK + phases complete + `rc_up=0` on all
  three
- **the `stack-core` roster is attributed** (`HARDEN-M257-2` in `decisions.md`) — 8 of iter-09's roster were
  fixed by Pass 1, and the largest remaining cluster is the BOX: `run_one`'s unittest column ImportErrors
  under the interpreter `pytest` resolves to here, a THIRD one that neither the census docstring (headed
  *"THERE ARE TWO INTERPRETERS"*) nor the README instructing its use names. Measured on a real module:
  3.9.6 OK, 3.12.13 `ModuleNotFoundError`, 3.14.6 OK
- **the other sections are green**: `stack-injection` 329 passed / 8 skipped; the five Go sections 44
  packages, 0 FAIL, run with `-count=1` after a first pass came back entirely cache hits (a cached green is
  not a reading); `demo-stack`'s milestone file 105 passed; B2's studio-acquisition fence 19 passed with its
  live arm actually running rather than skipping

**Flakes stabilized:** none needed. Flake gate: 3 consecutive clean runs of every test added this pass,
with pytest's random ordering plugin ENABLED (identical counts each time).

**Routed forward:** `FIX-M257-census-interpreter-namespace-import` — Fate 3. Each candidate repair changes
either what the census MEASURES or what pytest treats as a rootdir, and choosing between them is a decision
about which interpreter is canonical for this repo. It is also not M257's code — no iter of this milestone
touched `run_one`.

**Knowledge backfill:** `suite_census.py`'s two-interpreter heading, corrected in place with the third
interpreter, the measured import behaviour across all three, and why the repair is deferred.

**Stop condition:** continue-to-next-pass — Pass 2 still surfaced new material (the fingerprint hazard, the
emit seam, the interpreter cause), so the dimension scan has not yet come back empty. One confirmation pass.

---

## Pass 3 — 2026-08-12 — final

**Iters hardened this pass:** all milestone-touched code (cumulative) — the confirmation pass
**Tiks covered since prior pass:** n/a — same cumulative scope, third sweep over it

**Coverage delta on touched files:** **0 tests added.** That is the result, not an omission: the pass ran
the six dimensions over the cumulative scope again and found nothing new to write.

**Bugs surfaced + fixed inline:** none. Nothing new surfaced.

**Verification performed:**

- **flake gate, 3 consecutive clean runs** of every test this session added, with pytest's random-ordering
  plugin ENABLED rather than suppressed: 201 passed / 11 subtests, identical each run
- **the corpus fence family, green on the live tree**: `anchor_construct_guard` rc=0 (*"every resolvable
  anchor names a construct; 363 range citations lie inside their file"*), `corpus_index_guard` rc=0 over 86
  docs, `markdown_structure_guard` rc=0, `derived_value_guard` rc=0, `demo_knob_guard` rc=0 (*"the defaults
  table and the parsers agree, both directions"*). Three of these first returned **rc=2** under a wrong
  invocation of mine — they take positional arguments, not `--repo-root` — and exiting 2 (CANNOT RUN) rather
  than 0 is the fail-closed behaviour those guards are supposed to have. Re-run correctly, all green
- **`demo-stack`'s milestone files**: 124 passed / 7 subtests (`test_frontend_build` + the B2
  `test_studio_acquisition_m257`, whose live clone-walking arm RAN rather than skipping)
- **both literal ratchets unmoved** at 254 / 663, with COMMENT at 236 exact — the pass contributed nothing
  to any of the three
- **the user's environment untouched**, verified by container census at close: `demo-2`'s 11 containers and
  the 5-container dev stack (`anthropos-{backend,gotenberg,postgresql,redis,sentinel}`) all still up. No
  bring-up, teardown, seed or reset was run at any point; the only containers this session started were
  ephemeral `docker run --rm` reads against **`demo-1`** images, whose stack slot was already free

**One item did not return, and it is stated rather than waited out.** The whole-tree `stack-core` pytest
sweep launched at the start of Pass 1 was still running at close, 53+ minutes in against a 32–57 minute
measured band on this permanently-contended box — and its own runs of this session's tests were competing
with it for the box. It is **not load-bearing for the attribution**, which was answered on stronger
per-item evidence and is recorded as `HARDEN-M257-2`: iter-09's roster read off its own on-disk sweep log,
every previously-failing census file re-run individually, and the largest remaining cluster's cause measured
directly across all three interpreters on the box. An aggregate count from a run whose tree changed
underneath it mid-flight would be *less* interpretable than that, not more.

**Publication:** `rosetta-extensions` `main` pushed to origin at `7516fa8`, tagged
**`fast-build-m257-harden-1`**, and the tag verified **on origin** (`git ls-remote --tags`) rather than
merely created — *tagging is not publishing*.

**Stop condition:** stabilized — coverage delta 0 across the pass, and the dimension scan found nothing new.
