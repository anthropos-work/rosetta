---
milestone: M257x
iter: 06
---

# iter-06 — progress

**Type:** tik, under `TOK-01` step 2→3. Declared **3-step planned shape** (see overview.md), so the
scope-creep tripwire counts against that, not a single-target tik.

## Step 0 — the re-survey found something the pre-compute had missed

The pre-computed mapping (committed at `34cb120`, deliberately flagged "verify, do not trust") listed
**11** tables. The live write surface is **12**: `validation_check_results` (`content_stories.go:204`)
was absent. Found by measuring the *schema string literal* the `CopyRows` call actually carries
(`rg '"jobsimulation"' -g '!*_test.go'`) rather than the SQL text — 7 files, and they are exactly the 7
seeders that failed on iter-04's bring-up.

**Column drift: refuted, and it was measured, not eyeballed.** All 12 targets exist in `public`, and
every one of the 12 COPY column sets is a subset of its target's `information_schema.columns` — proven
not by reading but by running the seed and watching all 12 land (row counts below). `jobsimulation`
holds **0** tables, and `public.sessions` does not exist, so nothing here could have failed silently.

## Step 1 — the re-point, and the PAIR that had to be removed instead

| what | how |
|---|---|
| 11 tables | `jobsimulation.<t>` → `public.<t>`, same name |
| `sessions` | → `public.job_simulation_sessions` — the one rename (`D-M257x-1`) |

**Three of the five session writers were writing a PAIR** — the same rows, the same ids, into
`jobsimulation.sessions` *and* `public.job_simulation_sessions` (`persona_write.go`,
`content_stories.go`, `hiring_funnel.go`). M257 added the app-side half beside the legacy one rather
than replacing it. Re-pointing the legacy half would have written one table twice, so it was **removed**
— and `platform-alignment.md` §7 rule 2 ("never re-point to nothing without asserting the replacement")
is satisfied *by the very next step in the same slice*, which is the strongest form that assertion takes.
The other two (`jobsim_sessions.go`, `ai_readiness_funnel.go`) had no app-side half and were genuinely
re-pointed.

**A re-point turned a free ordering into a required one — twice.** Two writes in different schemas have
no FK between them; once both land in `public` the FK is real:

- `public.job_simulation_sessions` had to move to the **front** of both flush slices (every other row in
  the fan-out FKs it; it had been last).
- `interactions` + `actors` had to move **above** it in `resetTables` (they became true children of a
  list entry, where before their parent lived in another schema). The existing
  `TestResetTables_CoversAllSeededSurfacesInFKOrder` is what would have caught a miss here.

`dna/fidelity_probe.go`'s read of `validation_attempt_skill_results` was re-pointed too — it reads the
demo stack.

**Deliberately NOT re-pointed:** `cmd/content-capture/main.go` and `contentsession/sourcing.go`. They
read **production** read-only at authoring time, and prod still carries the legacy `jobsimulation` schema
pending platform M710 (`D-M257x-3`'s per-environment axis: the same name is live in one environment and
absent in the other). Its test assertion was rewritten by the sweep and reverted by hand.

## Step 2 — the fence: `stack-core/tests/test_write_target_schema_fence.py`

`platform-alignment.md` §8 asked for a static schema fence "generalised off its hardcoded relation
tuple". This one goes further: **it names no dead schema at all.** It reads the legal set at run time
from `repos_yml_schemas_to_create` — the same `stack-core/lib/repos_yml.sh` the migrator (iter-02) and
the verifier's schema probe (iter-05) use — and asserts

    every schema a seeder WRITES  ⊆  the schemas the migrate step CREATES

so when the next fold drops a schema from `repos.yml`, the legal set shrinks by itself and every stale
write goes RED on the next test run, with no edit to the fence and nobody needing to notice.

**Its first cut was wrong in an instructive way, and that is now a protocol rule.** Recognising
`{"<schema>", "<table>",` and `"<schema>.<table>",` *anywhere in a file* flagged 40-odd casbin grants
(`{"default", "admin", "org:feature:insights"}`) and the string `"clerk.com"`. That is the same mistake
as asserting on `file.read()` — the regex WAS the construct. Scoping each to its enclosing block
(`var resetTables = []string{`, `[]struct{ schema, table string … }{`) removed every false positive
without weakening the true one. The tempting alternative, an allow-list, is Trap A in miniature. Landed
as §8 rule 4.

**Mutation-verified RED twice, then reverted:** a seeder re-pointed back at `jobsimulation` (named
`activity.go:97`), and a `resetTables` entry under `skillpath` (named `main.go:52`). Its fixture is
itself mutation-verified (`test_fixture_can_fail`) — iter-02 shipped one that could not fail.

## Step 3 — the debt shrank, and the fence made it a deliberate act

`REXT_TRANSITIONAL_SCHEMAS` went `"cms jobsimulation"` → `"cms"`. The no-growth fence fired exactly as
designed:

    AssertionError: Debt paid down (['jobsimulation'] re-pointed) — update _EXPECTED_TRANSITIONAL in
    this fence to lock the win in. This failure is GOOD NEWS and the fix is a one-line edit.

That branch is the reason the paydown is a claim in writing rather than an invisible one-word diff.

## Live re-measure — the headline

Re-ran the real set-dress seed (`stackseed --stack demo-1 --seed presets/stories.seed.yaml`) against the
running demo-1. **7 failing seeders → 0**, all seven writing real rows:

    content-stories 2014 · hiring-funnel 46 · jobsim-sessions 853 · personas 1232
    activity 2160 · succession 165 · ai-readiness-funnel 8959
    isolation: clean (66 audited write attempts, 62980 rows, prod=false)

and all 12 targets populated in `public` — `job_simulation_sessions` 1644 · `activity_events` 2160 ·
`interactions` 1574 · `validation_check_results` 1479 · `validation_attempt_skill_results` 1087 ·
`validation_criterion_results` 635 · `validation_attempt_results` 524 · `actors` 377 ·
`interview_extraction_results` 167 · `collaborative_assets` 6 · `code_submissions` 2 ·
`interview_aggregated_reports` 1.

**autoverify FAILED 2 → 1:**

    ✓ hiring org set-dressed: 5 shared positions + 42 candidate HIRING sessions   <- was 0 sessions
    ✓ verify live: all liveness + readiness probes passed
    ✓ container liveness: all 16 expected container(s) running

iter-05's read that the hiring-org warning was **downstream of these 42P01s and not a separate defect**
is confirmed. The single remaining ✗ is `FIX-M257x-academy-not-serving`, already routed.

## Side discovery — a guard RED since iter-02, in a suite nobody ran

`dev-stack/tests/test_dev_stack.py::test_migrate_dev_is_shellcheck_clean` has been RED since **iter-02**
added the `. repos_yml.sh` source line to `migrate-dev.sh`: without `-x` shellcheck reports SC1091
("not following") and exits 1. iter-04 installed shellcheck and closed two guards of exactly this class;
this one sat in a different suite. Fixed with `-x` (the demo twin's own precedent for a sourcing script,
`test_tooling.py` on `up-injected.sh`) plus the two `# shellcheck source=` directives `migrate-demo.sh`
already carries. `-x` is also the stronger check — it analyses the sourced lib instead of skipping it.

**Third consecutive iter in which the wider suite run was where the finding was.**

## Suite state

| suite | result |
|---|---|
| `stack-core/tests` | **354 passed** (346 + the 8 new fence tests) |
| `stack-seeding` (go) | all packages ok — 128 test assertion sites re-pointed |
| `stack-snapshot` / `stack-secrets` / `alignment` (go) | all ok |
| `stack-verify/tests` | **210 passed** (iter-05's 3 e2e-collection failures now pass on this host) |
| `demo-stack/tests` + `dev-stack/tests` | 1132 passed / 8 failed → shellcheck fixed → **7 failed** |
| the 7 | **reproduce identically on the pristine `stack-demo` control clone @ `fast-build-m257x-iter-05`** — pre-existing, routed |

## Close — 2026-07-31

**Outcome:** the milestone's headline route landed and was proven live. All **12** `jobsimulation.*`
write targets (the pre-compute said 11) re-pointed at the `public` tables jobsim-in-app created; the
three co-written session PAIRs collapsed to their canonical half; `REXT_TRANSITIONAL_SCHEMAS` shrank to
`cms`; and a new **derived** write-target schema fence makes the next fold fail offline instead of on a
bring-up. Measured on the live stack: **7 failing seeders → 0**, autoverify **2 → 1**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this was a tik) — (3) re-scope: n — (4)
user-blocker: n (the 7 demo-stack failures reproduce on the untouched control clone, so they are a
routed-forward pre-existing item, not a regression this iter introduced; both trees clean) — (5)
cap-reached: n (1 tik of 5) — (6) protocol-stop: n — Outcome: continue
**Metric:** clauses met **0/5 → 0/5**. Clause 4's jobsimulation half is **paid and fenced** (the debt
list shrank, deliberately, through a fence that fails on a shrink); the `cms` half remains, so the clause
is not yet claimable. Sub-progress: seeder failures **7 → 0**; autoverify FAILED **2 → 1**; 12 write
targets moved; transitional schemas **2 → 1**.
**Decisions:** D-M257x-6 (removal-over-re-point for a co-written pair), D-M257x-7 (prod-read paths stay
on the legacy schema until platform M710) — see this iter's `decisions.md`.
**Side-deliverables:** the `migrate-dev.sh` shellcheck guard, RED since iter-02 in an unrun suite.
**Routes carried forward:**
- `REPOINT-M257x-cms-similarity-writes` (**NEW**) — `stack-snapshot/simembeddings` replays
  `cms.similarities` + 3 children; all four moved to `public`, and the demo already logs
  `sim-embeddings replay skipped (rc=4) — the stack's "cms" schema is missing/empty`. It is the LAST
  transitional entry, so it is what stands between here and a claimable clause 4. → next tik.
- `CHECK-M257x-live-clone-suites-red` (**NEW**) — 7 `demo-stack/tests` live-clone/live-container tests
  RED on this host, reproduced on the pristine control clone. Two are demopatch apply/revert
  round-trips, which is the class that shipped a 76 s members grid for four releases, so this is not
  cosmetic. → later tik.
- `FIX-M257x-academy-not-serving` — now the **only** remaining autoverify ✗, so it has become
  clause-1-critical rather than incidental.
- `HOST-M257x-toolchain` — **residual shrank and grew**: the Playwright/npm e2e tests now pass (210/210
  in `stack-verify`), but Go tests in `stack-seeding` need `GOPRIVATE` + a `url.insteadOf` SSH rewrite
  to resolve the private `ai v1.40.1` module. Same Trap E class as iter-04's `apply-authn.sh`, in the
  developer's own Go env this time.
- All other routes unchanged.
**Lessons:**
- **A pre-computed mapping is a hypothesis with better provenance, not a finding.** The pre-compute was
  short by one table and correct about everything else. Its most valuable content was the two things it
  refused to conclude.
- **Removal can be the correct re-point.** §7 rule 2 forbids re-pointing to nothing — but when the
  canonical target is already written on the next line, deleting the legacy write IS the re-point, and
  the assertion the rule demands is already sitting in the diff.
- **A re-point can silently turn an arbitrary ordering into a required one.** Cross-schema writes have
  no FK; same-schema writes do.
- **Scope a fence's construct to its block.** A fence that cries wolf gets disabled, and a disabled
  fence is indistinguishable from never having written one.
- **Design debt lists to fail on a SHRINK.** Otherwise paying debt down is a one-word diff nobody sees,
  and a release later somebody asks "didn't we already fix that?"
