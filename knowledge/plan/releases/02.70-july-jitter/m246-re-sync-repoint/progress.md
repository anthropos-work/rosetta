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

## Completeness Ledger

### Deferred

### Dropped
