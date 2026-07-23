# M243 — Progress

Section milestone. Checklist stub from the roadmap In-list. **COMPLETE — live GREEN + DB-verified.**

## Sections

- [x] **`assign-and-track.UC1`** — filled `playthroughs/manifest/assignment-monitoring.yaml` UC1
  (`playthrough: pt-assignment-assign`; `preconditions: [public-catalog, org-unassigned-member]`).
- [x] **`/enterprise/assignments` page object** — `e2e/lib/assignments-page.ts` (`AssignmentsPage`),
  the first MUTATING page object; keyboard-picks the antd catalog Select; read-back = affordance-count delta.
- [x] **`pt-world` precondition** — `org-unassigned-member` added to `seed-worlds.yaml` in lockstep
  (no seed CODE change needed — Org A already provides 34/40 unassigned targets).
- [x] **Spec `e2e/tests/assignment-assign.spec.ts`** — tagged `@pt:pt-assignment-assign`; live GREEN on
  demo-1 (7.9s); DB-verified the `organization_assignments` row landed (Meridian Labs skill_path active 6→7).
- [x] **Delivers** — `corpus/ops/demo/playthroughs.md` (15 → **16 live Playthroughs, 0 TODO**) +
  `README.md` + `CLAUDE.md` count updates + the assign-WRITE subsection.

## Supporting
- [x] `url-shapes.ts` `ASSIGNMENTS_URL` + `isOnAssignments` + unit-test pin (73 TS unit tests green).
- [x] `manifest/corpus_test.go` M204 pin reversed (UC1 TODO → implemented `pt-assignment-assign`).
- [x] Gates green: ptvalidate static + both-way integrity + precondition-coverage + datadna closure PASS + Go tests.

## M243: Hardening

### Pass 1 — 2026-07-22
Section milestone; all milestone CODE lives in `rext` at tag `sound-check-m243-assign-write-playthrough`
(build `b327e78` → harden HEAD `2ef5962`, tag moved + re-pushed to origin; corpus side docs-only). Touched
files: `e2e/lib/assignments-page.ts` (live-Page page object), `e2e/tests/assignment-assign.spec.ts`
(the live e2e — the WRITE Playthrough), `e2e/lib/url-shapes.ts` + `tests/url-shapes.unit.spec.ts`
(`ASSIGNMENTS_URL`/`isOnAssignments`), `manifest/assignment-monitoring.yaml` (UC1),
`manifest/corpus_test.go` (M204 pin reversal), `seed/seed-worlds.yaml` (`org-unassigned-member`).

**Coverage delta (milestone-touched files):**
- Go `manifest` package: 100.0% → 100.0% statements (already maximal; not a line-bump target — the
  work was proving the two honesty-teeth actually BITE, see below).
- TS `url-shapes.ts` (`isOnAssignments`/`ASSIGNMENTS_URL`): fully exercised; `url-shapes.unit.spec.ts`
  73 → 77 tests (+4), closing the orthogonal-dimension gap M204's five manager predicates had but
  M243's net-new predicate lacked.

**Anti-toothlessness / mutation-verification (the release thesis, the crux for M243):**
- **The live assign-WRITE assertion is NOT toothless.** `assignment-assign.spec.ts` asserts the
  read-back FLIP — `expect.poll(() => assignableCount()).toBe(before - 1)` — NOT merely a closed
  dialog. Analytically airtight: a modal that submits but writes nothing leaves the members-query
  refetch returning the SAME count, so `before - 1` is never reached and the 20 s poll times out RED.
  `confirmAssign()` additionally fails loud if the modal never hides. The precondition (an unassigned
  target exists) is enforced at RUN-time by `expect(before).toBeGreaterThan(0)` + the
  `assignSkillPathButtons().first()` visibility gate — a seed that assigned every member fails the
  test explicitly, never a silent no-op. Live DB proof already on record at build (Meridian Labs
  `organization_assignments` skill_path/active 6 → 7, the exact assign).
- **EMPIRICALLY mutation-verified (verify-then-restore, no commit) the two Go teeth that keep the
  assign test honest across the realistic silent-no-op vector (seed / route drift):**
  1. Precondition-lockstep tooth — renamed the `org-unassigned-member` capability in
     `seed-worlds.yaml` → `TestRealCorpus_ValidatesAgainstSeedWorlds` went RED with the exact
     targeted error (`[precond] …UC1: precondition "org-unassigned-member" is not a capability of
     world "pt-world"`); restored → green. So a seed change that removes the assign target surfaces
     at validate-time in CI, not as a mysterious live failure.
  2. Corpus-pin tooth — flipped UC1 `playthrough:` to `TODO` → `TestRealCorpus_ManagerCoverageIsPresent`
     went RED with three targeted errors (regressed-to-TODO / no-longer-non-TODO-id / wrong-tag)
     while `…ValidatesAgainstSeedWorlds` correctly stayed GREEN (TODO is a legitimate build-reference
     gap); restored → green. So a silent deletion/regression of the one net-new v2.6 journey is a
     red test.
- **The live-write-path-breaks-while-a-target-still-exists vector** has no static guard by nature —
  it is caught by RUNNING the Playthrough (its whole purpose), on every `run-playthroughs.sh N
  --reset` / prove-on-billion cycle. Harden strengthens the static scaffolding around that live guard;
  it does not replace it. A full live re-break-the-write on demo-1 was NOT run: demo-1 is currently
  seeded with the showcase roster (8 orgs, not pt-world) and re-seeding it would destructively
  commandeer a shared stack (`stackseed` not on PATH; potential M244 use) — disproportionate for a
  harden pass given the build-time DB proof already exists.

**Tests added:**
- `e2e/tests/url-shapes.unit.spec.ts`: +4 unit tests — a dedicated `isOnAssignments — cross-predicate
  edge + interaction depth (M243 harden)` block: mutual-exclusion vs the other four manager landing
  routes (both directions), sub-tab-does-not-leak-into-a-sibling-by-tab-name, path-relative matching
  (`page.url()` vs `waitForURL` parity), and an adversarial fuzz grid (catastrophic-backtracking
  probe, embedded newline, wrong-case, percent-encoded slash). Parity with the M204 final-harden
  block, which had covered the five older manager predicates but not M243's `isOnAssignments`.

**Mutation-verification of the NEW tests (net-new protection proven):** host-anchored `ASSIGNMENTS_URL`
(`:\/\/[^/]+…`, the `AI_READINESS_URL` shape) — a regression NO pre-existing test caught — failed
EXACTLY ONE test, the new path-relative case (line 232), 63 others green; restored → 64 green, 3
consecutive clean runs (flake gate).

**Bugs fixed inline:** none — the milestone build was a single clean commit with no `fix`/`bug`
commits, and no defect surfaced during hardening. The assign assertion and the enforcement teeth were
already correctly designed; harden proved they bite and closed the predicate parity gap.

**Flakes stabilized:** none observed (unit specs deterministic; 3× clean).

**Knowledge backfill:** no KB-worthy findings this pass. The anti-toothlessness read-back contract is
already documented (decisions.md D2 + `corpus/ops/demo/playthroughs.md`'s Playthrough-vs-presence
framing); the antd-v6 rc-virtual-list Select keyboard-pick lesson is already in decisions.md D4; the
`isOnAssignments` predicate hardening is a test-parity deepening, not a new system invariant.

### Stop condition
One pass. The Step 2b scan found a single genuine gap (the `isOnAssignments` orthogonal-dimension
parity) — closed. Go coverage was already 100% (the substantive work was proving the honesty-teeth
bite, done empirically). A second pass would only add shallow line-bumps — stopped to avoid
gold-plating.

## M243: Final Review

Clean close — a small, well-hardened section milestone whose code lives entirely in `rext
playthroughs/` (Go manifest + TS e2e) with a docs-only corpus side. The review found **zero
fix-required findings**; no code fix commits were made at close. Full ledger:

### Scope
- [x] All 5 sections + supporting items checked; the sole in-manifest `TODO` M243 was chartered to
      fill is filled — `ptvalidate` reports **7 products / 16 use cases / 16 live / 0 TODO**; 16 `@pt:`
      spec files, 16 distinct tags. Counts consistent across `playthroughs.md` / `README.md` / `CLAUDE.md`
      (15→16, 1→0 TODO). 0 TODO/FIXME/HACK in the touched files. No scope gaps.

### Code Quality
- [x] No must/should/nice findings. `AssignmentsPage` follows the established `PageObject` discipline
      (find-only landmarks; `<main>`-scoped `table tbody tr`; role/visible-text anchors; no CSS/testid);
      `ASSIGNMENTS_URL` correctly segment-anchors so it excludes an `/enterprise/assignments-report`
      look-alike (the M203 `-`-boundary hazard). No dead code. `targetMemberName` is best-effort
      non-blocking (P2). `go vet` + `gofmt` + `tsc --noEmit` all clean.

### Adversarial (Phase 2c)
- [x] One scenario cluster recorded in `decisions.md` (count-delta read-back vs table pagination/re-sort
      [S1], refetch-transient [S2], antd keyboard mis-pick [S3]) — all verified **fail-closed**, S1
      empirically clean on the real 40-member Org A roster. No code change.

### Documentation
- [x] No gaps. D2 (anti-toothless read-back) / D3 (precondition lockstep) / D4 (antd-v6 Select lesson)
      already blended into `playthroughs.md` at build (verified accurate). No new top-level unit →
      per-unit handbook contract N/A. Doc house-style uses no `(#M-D)` reference tags.

### Tests & Benchmarks
- [x] Go `manifest` fresh-pass (`TestRealCorpus_*` green; 100% stmts, already maximal); TS unit **77**
      (73→77, the +4 `isOnAssignments` parity block); typecheck clean; `ptvalidate` static gate green.
      **Flake gate:** TS unit **77 passed 5/5** sequential (deterministic). The live-browser spec
      (`assignment-assign.spec.ts`) is proven GREEN + DB-verified at build; its cold re-drive on `billion`
      is **M244's declared scope** (roadmap `Out:` + M244 exit-gate (h)) — not a close gap.

### Decision Triage
- [x] D2 / D3 / D4 → already blended into `playthroughs.md` (verified). D1 (assign surface/flow) / D5
      (corpus-pin reversal) / D6 (dev-run roster refresh) / Adversarial subsection → archive
      (maintainer-only). No new blending.

### Deferrals (Phase 1b)
- [x] Audit **YELLOW** — 0 new; `DEF-M235-03` assign-WRITE **RESOLVED Fate-1** (the standing-backlog carry
      M243 was chartered to land); 4 inherited items already fated → M244 (now-adjacent terminal expiry).
      No blockers. Report: `audit-deferrals/deferral-audit-2026-07-22.md`.
