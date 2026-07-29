# Hardening Ledger — M256 playthrough sharpening

Mode: **incremental**. First harden pass of the milestone, so scope is **all 12 closed iters**
(iter-01 → iter-12; no prior ledger entry existed). Gate has NOT fired (iter-12: `**Gate:** NOT MET`),
so this is not `--final`.

**Theme this pass hunted, from the milestone's own record:** *checks that report success without having
checked.* M256 had already found, in its own code and in code it inherited, a fence shipped as
`_sha(p) != _sha(p)`, a citation guard green over 28 of 29 rotted anchors, a UTC-parsed-as-local age
check failing OPEN west of UTC, a flake gate reading `tail`'s exit code, and a cockpit advertising heroes
its identity server could not serve. So the operating question for every fence, guard, gate and assertion
in the 12-iter footprint was **"prove it can go RED"** — and where it could not, that is the finding.

Environment (stated with every number, per D-v28-13): local laptop, darwin 25.1.0, no live stack (this
pass is unit-test-level by constraint; nothing was run against a demo, and `billion` was not touched).

---

## Pass 1 — 2026-07-29 — incremental

**Iters hardened this pass:** iter-01 → iter-12 (all closed iters; first pass)
**Tiks covered since prior pass:** all 12 iters in the milestone (no prior harden pass)
**Scope:** rext `playthroughs/` · `demo-stack/` · `stack-seeding/` (the 57-file M256 footprint,
`6ca8764..HEAD`), plus `stack-core/` and `stack-verify/` where the pass's own findings led.

**Dimension emphasis:** this footprint is almost entirely *test and fence* code, so the productive
dimensions were **error paths / edge cases / regression** applied *to the fences themselves* (can each
one fail?) rather than to production runtime — plus **mutation testing**, which is the only dimension
that answers the milestone's actual question.

### Bugs surfaced + fixed inline

1. **The `networkidle` ban was two spellings, not a ban** (commit `aa237b0`). iter-03's fence encoded
   the doctrine as two tightly-anchored regexes. Four plausible shapes score **zero** hits against them,
   two of which are not hypothetical: `waitUntil: opts.waitUntil ?? 'networkidle'` is
   `stack-verify/e2e/lib/cockpit-login.ts:87` **verbatim** — the root cause of the entire class — and
   `waitForLoadState('networkidle', { timeout: 4_000 })` exists ~20 times one directory away, which
   `hero-login.ts` forwards into. Now a **token scan** with one enumerated allowance (`waitUntil?:`, a
   type position, provably non-executable) and the `stack-verify` scope exception written down.
   **RED-proven** by injecting the bounded-settle shape into a real page object → named
   `profile-page.ts:108`. The old regex pair saw nothing there.

2. **Negative-control coverage credited per FILE while the gate counts per PLAYTHROUGH** (`aa237b0`).
   Non-global regex → first hit only → every `@pt:` id in the file credited. iter-06 closed this exact
   arity hole for `@pt-mutation:`, one field over, and left it open on **the number clause 2 is scored
   by**. Latent only because `studio-builder.spec.ts` (the sole two-Playthrough file) declares no control
   yet — so the edit that closes those two controls is the edit that would have inflated the count by two.
   Now arity-checked, fail-closed (short arity credits **none**). **RED-proven**: injecting one tag there
   fails the new guard by name and the count holds at **13**, where the old parser read **15** — clearing
   the floor of 13 on two credits nobody declared.

3. **The casbin-reset fence pinned the list and never the statement** (commit `9ef4549`). `main.go`
   advertises the list as "the lockstep that stopped g3 from being forgotten twice". There was no
   lockstep: three tests pinned `resetCasbinPTypes` / `quotedList`; none read the DELETE. **Proven** by
   re-inlining the historical bug (`WHERE p_type IN ('g2')`) — **all three iter-11 tests PASS and the
   package reports `ok`**. The whole fence for the defect passes the defect. `resetCasbinDeleteSQL()`
   extracted; three tests now read the rendered SQL (every listed p_type present, no unlisted p_type
   present, scoped DELETE never TRUNCATE), with the checker driven by mutants.

4. **15 classes / 76 tests never ran under direct execution** (commit `52eb4f4`). A `TestCase` subclass
   defined *below* a module's `__main__` guard is not defined when the guard runs. `python3
   test_roster_invariant.py` → **Ran 22 ... OK**; pytest → **27 passed**. The five silent tests include
   that file's own RED-proof for the live 12-dead-buttons defect. Six modules affected; worst
   `test_apply_patch_selfheal.py` at **11 of 27**. Guards relocated; fenced repo-wide by
   `stack-core/tests/test_test_collection_fence.py` (7 RED-proofs + a named pin on the six repaired
   files), **RED-proven** by re-appending a class to a real file.

5. **Two retry loops whose deadlines could never be enforced** (commit `cfaa1a9`) —
   `FENCE-M256-bounded-interaction`, discharged. Playwright's action default is `0` (no timeout), so both
   loops written *after* iter-06 D25 reproduced D25's shape exactly: unbounded `click()` inside, unbounded
   `waitFor` guarding, only `dialog().waitFor` bounded. iter-12's routing undercounted its own class twice
   — it counted `waitFor` when D25's cause was `click()`, and two of its four sites *are* loop guards. Six
   sites bounded (clicks moved inside the `try`, per D25's yield-to-next remedy); fence landed with its
   **exception boundary enumerated** (28 straight-line sites deliberately out of scope, and why).
   **RED-proven** twice — against D25's historical snippet, and live by reverting the org-admin click.

6. **The blocked-outcome test could pass on an EMPTY org** (commit `19fabb3`). iter-11's g3-withholding
   test — the write half of clause 2's `>= 1 blocked` — asserted `g3 == 0` for the opt-out org, which is
   equally true when that org seeded no memberships; and asserted `> 0` for the granted org where its own
   comment says "one per membership". Now: withheld org must have `g2 > 0` (population witness) **and**
   `g3 == 0`; granted org must have `g3 == g2`. **RED-proven by mutating the seeder** — a partial grant
   (1 of 3) and an emptied org both now FAIL, and both PASS under the old assertions.

### Side-deliverables (unrelated-but-correct; pre-existing reds, Fate 1 rather than carried)

- **`test_buildbench.py` had a unit test whose verdict depended on the host's free disk** (`52eb4f4`).
  It asserted `rc == 0` while `run` performs the pre-rep headroom assert, which D-M255-1 says fails
  **loud** (rc 1) on a host without headroom — 14.7 GiB free here against a 7 + 18 GiB floor. A correct
  refusal read as a failure and took `test_m255_mutation_battery::test_00_the_unmutated_baseline_is_GREEN`
  with it. The test one line above already handled exactly this. Not weakened (rc 2 still fails); the
  **abort branch, previously untested altogether**, now pins D-M255-1's promise that a refusal is
  *recorded* (`aborted` + `headroom.ok == false`). stack-core **2 failed/282 → 284 passed**.
- **5 red tests in `stack-verify/tests`, a suite no release close has ever run** (`52eb4f4`).
  `autoverify.sh` grew the M245 ant-academy probe; 1 of 4 `curl` fixtures was taught about it. Now a
  shared `_CURL_ACADEMY_OK` constant at all five sites. **154 passed** (was 5 failed/149), and **313 s →
  163 s** because each red test paid the probe's 3×3 s retry.
- **`blueprint.go` + `stories.go` were gofmt-dirty** from the iters (`19fabb3`).
- **Both source-scan fences reported the wrong line number** (`cfaa1a9`): comment-stripping *deleted*
  block comments, shifting every offender by the length of the file's 70–90-line header. Right about
  *whether*, wrong about *where* — a live mutation at line 62 was reported as line 24. Now
  line-preserving; verified the mutant reports 62.

### Coverage delta on touched files

Measured as **tests added and mutants proven**, not as a line-coverage percentage: this footprint is
fences, where a line-coverage number is close to meaningless (a fence's lines execute on every run,
green or red — the question is whether it can *fail*, which coverage cannot see). That is itself the
milestone's thesis applied to its own metric.

| suite | before | after |
|---|---|---|
| playthroughs unit specs | 119 | **131** (+12; +1 new fence file) |
| `stack-seeding/cmd/stackseed` | 3 reset tests | **6** (+3 statement-level) |
| `stack-core` | 276 pass, **2 fail** | **284 pass, 0 fail** (+8 new fence tests) |
| `stack-verify` | 149 pass, **5 fail** | **154 pass, 0 fail** |
| demo-stack + stack-injection | 1260 pass | **1278 + 420** (rosters re-run whole) |
| **mutants proven RED** | — | **11** (across 6 fences, incl. 3 live source reverts) |

### Knowledge backfill

- `corpus/ops/demo/playthroughs.md` (the declared `iteration_protocol_ref`) — three additions: coverage
  is credited **per Playthrough, not per file** ("when a count feeds a gate, the unit the parser credits
  must be the unit the gate counts"); the bounded-interaction class **outlived its first fix**, with the
  enumerated exception boundary; and a source-scan fence must report the right **line**, not just the
  right verdict.
- `corpus/ops/demo/latency-budget.md` — the `networkidle` ban must be a **token scan, not a list of
  spellings**, with the four escaping shapes and the `stack-verify` scope exception named.
- `corpus/ops/verification.md` — new section **"A suite nobody runs is a suite that is red"**: the
  enumerated-roster staleness (and why the fix is a *discovered* roster with an asserted count, not a
  longer list) and the hidden-test-class defect, with the shared rule — *a green verdict is only as wide
  as the set of things the runner actually looked at.*

### Flakes stabilized

None found. Flake gate: **3 consecutive clean runs** of every touched suite, rc captured explicitly per
run (never off a pipe) — playwright 3×18 pass, pytest 3×117 pass, go 3× `ok`.

### Verification

6 Go modules `rc=0` · `gofmt -l` clean across all 7 sections · Python **1698 pass / 2 skip / 0 fail**
(demo-stack + stack-core + stack-verify + stack-injection) · playwright **131 pass** · `tsc --noEmit`
clean · `ptvalidate` **VALID** (24 Playthroughs, 31 use cases, no phantom id from the new fence file) ·
direct execution now matches pytest on all six repaired modules.

**Stop condition:** `cap reached without stabilization` — the 3-pass incremental cap fired, and pass 3
was still finding new instances of the target class (finding 6, the vacuous blocked-outcome test). The
class is not exhausted; what remains is named below rather than implied.

**What is still uncovered:**
- **The 9 structural negative-control finals** → routed (see below). Not a harden gap: iter-12 *measured*
  that no suppression mechanism can exist for them.
- **Nothing in this pass was verified against a live stack.** The six bounded-interaction sites,
  `negative-controls.spec.ts`'s liveness floor, and every Playthrough are proven only by source scan and
  unit test here. That is a constraint of the pass, not a claim about the harness.
- **The enumerated-roster problem is diagnosed but not fenced.** A discovered-with-asserted-count roster
  is the real fix and it belongs to close-release, not here.

### Routed forward (three-fate rule)

- **The 9 structural negative controls** (`pt-workforce-*` ×4, `pt-activity-drilldown`,
  `pt-profile-{verified,growth,timeline}`, `pt-hiring-recruiter-compare`) → **Fate 3 → a future iter of
  M256.** The routing anticipated that hardening would land a *mechanism* (a way to suppress the
  outcome). iter-12 measured that **no such mechanism can exist** for these nine: the M44
  profile-completeness seeder gives **every** member a career and skills, so there is no vantage for whom
  the outcome is legitimately absent (`pt-manager` reads 1 / 10 / 1). What remains is per-Playthrough
  assertion **sharpening** against named seeded data — O(tests) build work that needs a live stack to
  verify. There is no mechanism here to build, so this is not a harden deferral; it is build-iter scope.
- **`PT-M256-readiness-step-asserts`** → **Fate 3 → a future iter of M256.** iter-12's finding stands
  unchanged (`MANAGER_STEP_NAMES` matches the not-enabled upsell panel page-wide). Re-scoping the locator
  inside the method panel is a spec edit requiring live verification.
- **A discovered (not enumerated) test roster with an asserted section count** → **Fate 3 →
  `/developer-kit:close-release` for v2.8.** Two consecutive closes missed a red suite because the roster
  is hand-maintained; the fix is process-level and release-scoped, above a milestone harden pass.
- **Live re-verification of the six bounded-interaction sites + the fences** → **Fate 3 → the milestone's
  next live-prove opportunity.** No stack was permitted this pass.
