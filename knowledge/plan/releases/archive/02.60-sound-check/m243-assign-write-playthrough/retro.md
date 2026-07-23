# M243 — Retro

## Summary
M243 was the **sixth and last post-barrier fix** in v2.6 "sound check" — the **one net-new hero journey**, a
`section` milestone. It landed `pt-assignment-assign`, the **FIRST MUTATING Playthrough**: a manager
(Morgan / `pt-manager`) logs in → `/enterprise/assignments` → Skill Paths tab → opens the assign builder for a
member with no skill-path org-assignment → keyboard-picks a public-catalog skill path + accepts the pre-filled
deadline → **Assign** → and the write **LANDS** — a real `public.organization_assignments` row
(`app.createOrganizationAssignments`) **read back** through the real members surface as the assignable "Assign
Skill Path" affordance count dropping by **exactly one**. All code lives in `rext playthroughs/` (Go manifest +
TS e2e) at tag `sound-check-m243-assign-write-playthrough` @ `2ef5962`; the corpus side is docs-only. It closed
the **~10-routing `DEF-M235-03` / M204 assign-WRITE carry** — a deferral that had ridden **5 releases** — as
**Fate-1**, taking the corpus **15 → 16 live Playthroughs, 0 TODO**. **0 platform-repo edits.** A clean complete
section close: the review found **zero fix-required findings** and made **zero close code fixes**.

## Incidents This Cycle
- **None at close.** The build shipped as a single clean commit (no `fix`/`bug` commits); the 1 harden pass
  surfaced 1 gap (below), not a defect; the close found 0 fix-required findings across scope/code/docs/tests/
  decisions and made 0 code fixes.
- **P3 (harden, not a defect) — an orthogonal-dimension test-parity gap.** M204's final-harden block had covered
  its five manager landing predicates against cross-predicate false-matches, but M243's net-new `isOnAssignments`
  lacked that coverage. Since the assign-WRITE Playthrough uses `ASSIGNMENTS_URL` as its `/login`-bounce guard, a
  predicate that false-matched a sibling manager route would be a green-but-wrong trap. Closed by a dedicated +4
  unit-test block (mutual-exclusion both directions, sub-tab-no-leak, path-relative parity, adversarial fuzz
  grid); mutation-verified (host-anchoring `ASSIGNMENTS_URL` reddened exactly the new path-relative case). No
  production bug — the predicate was correct; the test just didn't yet cover the orthogonal dimension.
- **No regressions.** Go 2005 unchanged (manifest test modified in place, +0 func); TS combined unit-spec run
  73→77; `ptvalidate` 16 live / 0 TODO; flake 0 (77 5/5 sequential).

## What Went Well
- **The anti-toothlessness thesis got its sharpest test — and passed structurally.** This is the FIRST
  Playthrough whose action mutates real state, so a modal-closed pass would have proven nothing. The assertion is
  the **read-back FLIP** — `expect.poll(() => assignableCount()).toBe(before - 1)` — which can only reach
  `before - 1` if a real `organization_assignments` row landed AND is read back through the real members query; a
  silent write-failure leaves the count at `before` and the 20 s poll times out RED (fail-closed). The close's
  adversarial pass confirmed the delta is fail-closed under table pagination/re-sort [S1], refetch-transient
  [S2], and antd keyboard mis-pick [S3] — S1 empirically clean on the real 40-member Org A roster.
- **A 5-release-old carry discharged as Fate-1, not re-deferred.** `DEF-M235-03` was the exact class the v2.5
  close named — *a fate needs a MILESTONE, not "a phase" / "the next X"*. M238's reservation gave it a real
  milestone (M243); M243 landed it in full. A deferral **leaving** the ledger is the deferral audit's success
  story, the opposite of scope erosion.
- **No new seed CODE, and the precondition was declared + enforced.** Org A already provides ~34/40 unassigned
  targets, so the assign flow had a deterministic reset-to-seed target without touching the seeder. And rather
  than rely on that implicitly, UC1 names `seed.preconditions: [public-catalog, org-unassigned-member]` with
  `org-unassigned-member` added to `seed-worlds.yaml` in lockstep — so a future "assign-to-everyone" seed change
  trips `ptvalidate` at validate-time, not as a mystery live failure.
- **The Go honesty-teeth were proven to bite, empirically.** Renaming the `org-unassigned-member` capability
  reddens `TestRealCorpus_ValidatesAgainstSeedWorlds`; flipping UC1 to TODO reddens
  `TestRealCorpus_ManagerCoverageIsPresent` (the reversed M204-era pin). Verify-then-restore, no commit — the
  guards were shown to fail on the realistic silent-no-op vectors, not just asserted to.

## What Didn't
- **The build under-covered its own net-new predicate's orthogonal dimension.** The `isOnAssignments` cross-
  predicate parity was a real gap the build missed and harden had to close — the recurring "a new predicate
  needs the SAME mutual-exclusion + path-relative coverage the established ones have" lesson (the M204 block was
  right there as the template). Cheap to fix, caught before close, but a reminder that a net-new route predicate
  should be born with its full parity block, not get one retroactively.
- **The live re-break-the-write was not re-run at close.** A full live break-the-mutation-while-a-target-exists
  was left to the build's DB proof (demo-1, `organization_assignments` 6→7) + M244's cold re-drive — re-seeding a
  shared stack for a docs-only close was disproportionate. Correct call, but it means the close's confidence in
  the live path rests on the build's single live run + the static teeth, not a fresh live re-run.

## Carried Forward
- **None new from M243** — a clean complete section close, 0 new deferrals; the deferral audit is **YELLOW**.
- **(Resolved this milestone) `DEF-M235-03` / M204 assign-WRITE** — LANDED as `pt-assignment-assign` (Fate-1);
  leaves the standing backlog. The expiry condition ("land it or DROP") met by landing it.
- **(Follow-on proof) → M244:** the **cold live re-drive on `billion`** of `assignment-assign.spec.ts` (now the
  16th of the 40 live-browser specs M244's exit-gate (c) owns) — proven GREEN + DB-verified at build on demo-1;
  M244 re-drives it cold, reset-to-seed, alongside the rest.
- **(Inherited, confirmed) → M244:** the **9 standing demo-stack test failures** (Fate-2, M238-D5 / M239-D13
  reap-17700), **DEF-M240-01** (real-video exhibit, Fate-3, pre-approved), **DEF-M239-01** (ENOSPC loud-build-
  fail, Fate-3) — all already homed; M244 (terminal) is the named, now-adjacent expiry.

## Metrics Delta
- **Tests:** Go **2005** (unchanged — manifest test modified in place, +0 func). TS combined playthroughs
  unit-spec run **73 → 77** (`url-shapes.unit` 60 → 64: the `isOnAssignments` parity block). Live `@pt:`
  Playthrough specs **15 → 16** (+1 `assignment-assign.spec.ts`; cold re-drive is M244's — so v2.6's live-browser
  denominator moves **39 → 40** = 24 stack-verify + 16 Playthroughs). `ptvalidate` **7 products / 16 use cases /
  16 live / 0 TODO** (was 15 live / 1 TODO). Flake **0** (77 5/5). Platform-repo edits **0**.
- **rext code-of-record:** tag `sound-check-m243-assign-write-playthrough` @ **`2ef5962`** (annotated on origin,
  peels to `2ef5962` = origin/main = local HEAD; **unchanged by close** — 0 close code fixes).
- **Deferral audit:** **YELLOW** — 0 new; 1 standing carry RESOLVED Fate-1 (`DEF-M235-03`); 4 inherited → M244.
