# M246 — Progress

Section milestone. Checklist from the roadmap In-list. **HARD go/no-go barrier — gates M247–M254.**

## Sections

- [x] **Seeder re-point** — rext `stack-seeding` writes `skillpath.skill_path_sessions → public.skill_path_sessions` in the **live** seeder code + tests (`cmd/stackseed/main.go:97`, `seeders/hero_activity.go:180`, `skillpath_sessions.go`, `content_nonsim.go`, `dna/data-dna.json` + the in-package test assertions). Surface **names** (`skillpath-sessions`) + the mirror `public.local_skill_path_sessions` left untouched. **DONE** — rext `97585f5`; 8 live sites + DNA + comments + ~16 test assertions; build/vet/gofmt clean, package tests green, zero missed sites (all-file-types sweep).
- [x] **Demo clone pins** — author `stack-demo/clones.pin.json` + wire the `DEMO_ADVANCE_CLONES=pinned` advance path + bump the **demo** clone pins to current `origin/main` (jobsimulation stays standalone — still live). **DONE** — rext `ee44b9a`. The advance path was ALREADY wired (M237); delivered the durable canonical pin (`demo-stack/clones.pin.json`, 12 repos @ current origin/main HEAD shas, **skillpath excluded** — not in current repos.yml) + a copy-if-absent seam that seeds it into the ephemeral workspace (never clobbers an operator pin). 2 new tests; module green.
- [x] **Injection-comment fix** — `stack-injection/gen_injected_override.py:16` skillpath comment → 3 subgraphs. **DONE + EXPANDED** — rext `88bcdb8`. Scope grew to "de-skillpath the LIVE bring-up path" once the compose check proved skillpath has no service: also dropped skillpath from `up-injected.sh` `INJECT_SVCS` (was building `demo-N-skillpath:injected`) + `verify_svcs` (was verifying a skillpath container) — both required for a green bring-up. The `INJECTED`-dict/`test_injection.py`/`exposure_claim_guard` skillpath hygiene is inert (no compose service) → M247 drift ledger.
- [x] **Cold `/demo-up` GREEN + drift ledger** — prove ONE cold `/demo-up` GREEN on the consolidated platform; transcribe observed drift into the M247 ledger. **DONE — GO/NO-GO PASS.** Cold demo-2 (pinned advance + strict freshness, local, billion untouched) came up green: build exit 0, 16 services, **561 rows in `public.skill_path_sessions`** (re-point proven), 3 subgraphs + 0 skillpath, health 200 + casbin 1250, all probes pass. 3 autoverify warnings all non-firing (D-07 AI-readiness demopatch→M250, D-08 fake-FAPI probe artifact, D-09 academy peripheral). Drift ledger `drift-ledger.md` D-01..D-09 emitted for M247. rext consumed at tag `july-jitter-m246-re-sync-repoint` (on origin).

## M246: Hardening

### Pass 1 — 2026-07-23
**Scope manifest (rext `stack-seeding` + `demo-stack`, the milestone's REAL code — rosetta branch carries no source):**
- Go re-point (`skillpath.*→public.skill_path_sessions`): 8 live write-sites (`cmd/stackseed/main.go`, `seeders/hero_activity.go`, `skillpath_sessions.go`, `content_nonsim.go`) + `dna/data-dna.json` + audit tags + comments. Existing tests already assert `public.skill_path_sessions` at EVERY site — the reset list + FK-order (`main_test.go`), each seeder's `findCopy(t,"public",…)` (activity/contentref/content_nonsim/hero_activity tests), and the copy-error `failTable:"public.skill_path_sessions"` paths. **Adequately covered + live-proven (561 rows) — no gap, no new test.** A blanket "no `skillpath` schema anywhere" grep-guard was considered and REJECTED as fragile (the surface NAME `skillpath-sessions`, the Go symbols, and `skillpath_sessions.go` legitimately keep the word).
- Pin/copy-if-absent seam (`demo-stack/ensure-clones.sh` + `clones.pin.json`): shipped 2 land-tests (seed-when-absent / never-clobber-operator). **2 seam branches were unasserted → deepened this pass.**

**Coverage delta (milestone-touched files):** the two shipped seam land-tests reached branches {canonical-present+no-op-pin→seed, operator-pin-present→skip}. The 2 new tests reach the remaining 2 branches: {no-pin-anywhere→no-op fallback} and {seed-runs-before-advance-gate ordering under `=main`}. Seam branch coverage 2/4 → 4/4 (Python `test_tooling.py` has no line-coverage tool wired; branch enumeration used as the finder — the go/no-go seam is small + fully enumerable).

**Tests added (rext repo, commit `c9fbf6b`, tag `july-jitter-m246-harden` @ `9b29f3a` on origin):**
- `demo-stack/tests/test_tooling.py`: 2 regression/edge tests —
  - `test_pinned_with_no_pin_anywhere_is_a_clean_no_op` — `DEMO_ADVANCE_CLONES=pinned` with NEITHER a canonical rext pin NOR an operator workspace pin: seeds nothing, checks nothing out, exits 0 with the "nothing to advance" no-op (the seam else-comment's promised fallback).
  - `test_canonical_pin_seeded_but_default_advance_does_not_check_out` — ORDERING invariant: the seam runs BEFORE the advance gate, so a canonical pin + `=main` advance seeds the workspace pin but must NOT turn into a surprise pinned checkout (protects the deliberate-staleness-is-legitimate contract; a refactor folding the seam into the `pinned` arm would break it).

**Bugs fixed inline:** none — the re-point + seam were correct; the deepening pinned two behavioral contracts the shipped tests didn't reach.

**Flakes stabilized:** none observed. Flake gate: 3/3 consecutive clean sequential runs of the 2 new tests.

**Knowledge backfill:** no KB-worthy findings — the two contracts are documented in-place (the seam's `ensure-clones.sh` comment + the test docstrings); no `knowledge/`/corpus doc needed a new behavioral fact (D-2 in decisions.md already records the copy-if-absent + never-clobber design).

### Verification
- Go `stack-seeding` package: all sub-packages `ok`.
- Python `test_tooling.py`: 168/168 green (full file). `py_compile` clean. Full `TestCloneFreshnessM237` class 26/26.
- rext worktree clean; new consumption tag `july-jitter-m246-harden` pushed to origin (rung-zero verified via `git ls-remote`).

### Stop condition
ONE pass — right-sized per the milestone's nature. The re-point is a MECHANICAL schema qualifier flip, already asserted at every write-site and LIVE-PROVEN (561 rows in `public.skill_path_sessions` on cold demo-2). The only genuinely under-tested surface — the 2 residual copy-if-absent seam branches — was closed. Full Step-2b scan surfaced nothing else worth adding; the loop terminates (scan clean).

## M246: Final Review

Cross-cutting close review (2026-07-23). Real code = rext `stack-seeding` + `demo-stack` +
`stack-injection` at tag `july-jitter-m246-harden` @ `9b29f3a` (on origin); rosetta branch carries plan
artifacts + the `update_guide.md` re-sync note.

### Scope
- [x] All 4 declared sections delivered (Fate 1); the §3 expansion (de-skillpath the LIVE bring-up path) is Fate-1, recorded D-3. No scope gaps.

### Code Quality
- [x] [clean] Seeder re-point uniform across all 8 write-sites (schema `skillpath`→`public` for `skill_path_sessions` only); comments follow code; mirror `public.local_skill_path_sessions` + surface name `skillpath-sessions` correctly untouched. gofmt/vet/build clean.
- [x] [clean] Copy-if-absent clone-pin seam defensive (never clobbers an operator pin; non-fatal on cp failure) with clear provenance. `clones.pin.json` = 12 repos @ current origin/main, skillpath correctly excluded.
- [x] [clean] `gen_injected_override.py` inert skillpath key carries an honest comment routing full removal to the M247 drift ledger.
- [x] [note] `up-injected.sh` header's trailing v2.1-M209 historical parenthetical ("4 services not 5") considered — a correctly-scoped history note, not a current-count claim; the authoritative "3 Go services" is stated above with the M246 annotation. Not a defect; not re-tagged (see decisions.md close note).

### Documentation
- [x] `corpus/ops/update_guide.md` re-sync note accurate (3 subgraphs; `public.skill_path_sessions`; migrate roster app/cms/jobsimulation) — matches drift-ledger ground truth. No new top-level unit ⇒ no handbook needed.
- [x] roadmap.md + state.md consistent with the built reality (updated in Phase 10).

### Tests & Benchmarks
- [x] All touched-stack suites green: Go `stack-seeding` (all pkgs ok), `test_tooling.py` 168/168, `test_frontend_build.py` 94/94, `stack-injection` 264 (9 skipped). Flake gate 5/5 on TestCloneFreshnessM237 (incl. the 2 new seam tests). Re-point asserted at every write-site + live-proven (561 rows). No gaps.
- [x] [note] pre-existing `ResourceWarning` in `TestMigrateRaceGuard` (test_tooling.py:1413) is not M246-touched; benign (unclosed file in an unrelated test) — carried, not a failure.

### Decision Triage
- [x] D-1 consolidation fact (skillpath→`public`, 3 subgraphs) → already blended into `corpus/ops/update_guide.md` (M246 reference tag present). No new blend.
- [x] D-2/D-3/D-4/D-5 + KB-1..KB-5 → archive (rext-internal / maintainer-only); no knowledge home warranted.

### Adversarial
- [x] One scenario recorded in decisions.md (pin/tag skew silently mis-schemas skill-path writes) — handled by the tag-pin contract + the harden seam tests; the dangerous direction fails loudly by design.

## Completeness Ledger

### Done (Fate 1 — delivered in M246)
- Seeder re-point `skillpath.*` → `public.skill_path_sessions` (8 live sites + DNA + comments + ~16 test assertions) — rext `97585f5`.
- Durable canonical `clones.pin.json` (12 repos @ current origin/main) + copy-if-absent seam — rext `ee44b9a`.
- De-skillpath the LIVE bring-up path (`gen_injected_override.py` comment + `up-injected.sh` INJECT_SVCS/verify_svcs) — rext `88bcdb8` (§3 Fate-1 expansion, D-3).
- Cold `/demo-up` GREEN go/no-go PASS (561 rows in `public.skill_path_sessions`, 3 subgraphs, 0 skillpath) + drift ledger emitted.
- `corpus/ops/update_guide.md` consolidation re-sync note.
- Harden: 2 copy-if-absent seam tests (branch coverage 2/4→4/4) — rext `c9fbf6b`.

### Confirmed-covered (Fate 2 — owned by a sibling milestone / the drift-ledger handoff)
- D-01 corpus skillpath→app reconciliation → M247 (explicit In-list). D-06 audit-prose → M247. D-07 AI-readiness perf-patch anchor → M250 (its domain). D-08 login re-prove → M254 gate (h) / probe-fix M251. D-02/D-03/D-04 inert rext hygiene → M247's triage of the durable drift ledger. All tracked in `drift-ledger.md` + `audit-deferrals/deferral-audit-2026-07-23-m246-close.md` (verdict GREEN).

### Annotated (Fate 3 — roadmap mutation)
- None. No sibling `overview.md` `In:` list was edited.

### Dropped
- None.

### Release-scope-breaking deferral (escape hatch)
- None.
