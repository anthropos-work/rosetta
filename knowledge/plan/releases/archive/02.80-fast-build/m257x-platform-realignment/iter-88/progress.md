**Type:** tik, under `TOK-05`.

## What the sweep found

The hypothesis held, and the second instance is **a live defect in the demo bring-up path**, not a test
artifact.

### Finding 1 — the AWS-bind mitigation went dead when its service did (LIVE)

`docker-compose.yml` binds `$HOME/.aws/credentials` into a container. On a fresh Linux box that host path
does not exist, **Docker auto-creates it as an empty DIRECTORY**, `aws-sdk-go-v2`'s `LoadDefaultConfig`
opens it successfully (opening a directory succeeds) and then fails `EISDIR` — and because the cobra root
sets neither `SilenceUsage` nor `SilenceErrors`, the container prints the **full usage block** and exits 1.
That symptom was misread for an entire release cycle as a missing `serve` subcommand. M217 fixed it by
clearing the volume in the generated override — **zero platform edits**, correctly.

That mitigation was keyed on `if name == "jobsimulation":`.

- Platform **`d11a403`** deleted the `jobsimulation` compose service.
- Platform **`838d907`** put the identical `$HOME/.aws/credentials` bind on **`backend`**.

So **the hazard migrated to the stack's most important container and the mitigation did not follow it.**
The override became unreachable code guarding a service that no longer exists.

**And the tripwire built to catch exactly this said nothing.** `test_jobsimulation_has_ONLY_the_aws_bind`
looks the service up in compose, does not find it, and calls
`skipTest("jobsimulation not in the compose")` — a skip, which reads exactly like a pass (§5 rule 8). Its
sibling `test_no_OTHER_service_carries_a_HOME_bind` **passed**, asserting *"expected exactly 1 `$HOME`
bind (jobsimulation's AWS creds)"* — the count was right and the claim inside it was false, §5 rule 17
in one line.

**Repair — derived, not re-pinned.** `platform_topology.services_with_only_home_binds()` reads the raw
compose text and returns the services whose volumes are **all** `$HOME`-rooted binds; the generator keys
on that set. Measured at `0c91421` it returns exactly `{backend}`.

Three deliberate choices:
- **Raw text, not the resolved config.** After `docker compose config` expands them,
  `$HOME/.aws/credentials` and the stack's own `…/stack-demo/platform/data/postgresql` are both just
  absolute paths under the user's home; the intent survives only before expansion.
- **ALL, not ANY.** Clearing a service's volumes wholesale is safe only when every one is expendable. A
  service mixing a home bind with a volume it needs is **excluded** — the caution the original override's
  comment asked for, now enforced instead of described.
- **Fail closed on an empty derivation** (§5 rule 14): the tripwire now asserts that *some* service
  carries a home bind, so "the platform removed the last one" has to be a deliberate deletion rather than
  a silent pass.

### Finding 2 — the demopatch anchor check has been skipping since M254

`TestRealManifestsAgainstTheRealClone` exists, in its own words, because *"the app manifests had only
manifest-internal self-consistency checks and were NEVER validated against a real clone."* It took the
target path as a **second, hand-maintained argument** beside the manifest.

M254 re-pointed `app-aireadiness-snapshot-loadmembers` from `internal/workforce/ai_readiness.go` to
`internal/aireadiness/readiness.go` when `app` split the package. The manifest moved; the test's copy of
the path did not; `os.path.isfile` failed; the test **skipped**. It has skipped ever since — so the one
check that validates a demopatch anchor against a real clone silently stopped validating anything, which
is precisely the failure its own docstring says it exists to prevent.

**Repair:** the path is now **derived from the manifest's own `path:` field**, and the skip is fail-closed
— *"no clone on this box"* still skips; *"the clone is here and the manifest's path is not in it"* is now
a **failure**, because that is the drift. Re-run: the aireadiness anchor **resolves** (M254's re-point was
correct), and `stack-injection` went from **2 skips to 0**.

## Sections run at platform `0c91421`

| section | result |
|---|---|
| `stack-injection` | **7/7 OK, 0 skips** (was 7/7 with 2 skips — both holes, both closed above) |
| `stack-verify` · `dev-stack` · `demo-stack` | **41 files, 34 OK**, 5 FAILED, 1 ERROR, 2 skips |
| `stack-core` (iter-87) | 35 files, 34 OK, 1 FAILED — the documented perishable iter-48 fixture |

**The sweep's whole point was that unexercised checks go quiet, so the residual is enumerated rather than
summarised**, and none of it is repaired in this iter — see the routes below.

| file | failure | first read |
|---|---|---|
| `demo-stack/test_demopatch` | 2 × `pre_sha256` mismatch against the live clone | **the most interesting**, and it is *not* obviously a defect: `demopatch-spec.md`'s self-healing freshness gate says *"the anchor is the contract; the whole-file sha is only a baseline"*, so a drifted whole-file sha may be the design working. Needs the spec read against the assertion before anyone re-pins anything |
| `demo-stack/test_ant_academy` | 1 × sha mismatch on the real `next.config` | same class |
| `demo-stack/test_back_to_cockpit_m249` | `next-web-back-to-cockpit: revert failed` | a **G5 self-revert** failure — a different and more serious class than a sha baseline |
| `demo-stack/test_migrate_race_live` | 3 | name says `_live`; likely needs a live stack |
| `stack-verify/test_verify` | 1 ERROR, `test_demo_3_shifts_ports_and_project` | needs a live demo-3 stack |
| `demo-stack/test_interview_flag_patch_m232`, `test_purge` | 1 skip each | **unnamed holes** — this iter's own rule says name them |

**Do not read these as caused by the platform advance.** `next-web-app` sits 41 commits behind origin and
was **fetched** this run, and by §5 rule 41 a check that resolves at `origin/main` is armed by the fetch —
so the honest statement is that **the vantage changed and these have not been graded at it before**, which
is different from "they broke". Establishing which is the first job of the iter that takes them.

## Close — 2026-08-05

**Outcome:** the class iter-87 named is real and its second instance was a **live defect in the demo
bring-up path** — the `$HOME/.aws/credentials` mitigation had been keyed on a deleted service while the
hazard moved to `backend`, with its tripwire skipping. Both that and a demopatch anchor check that has
been skipping since M254 are now **derived from the platform artifact instead of restated beside it**, and
`stack-injection` runs with **zero skips**.
**Type:** tik
**Status:** closed-fixed — both declared lines landed; the sweep completed on the sections that read
platform artifacts
**Gate:** NOT MET — **4 of 5, unchanged.** No reading was taken; clause 5 is graded only by a reading that
returns zero. The reading series is untouched by this iter.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this was a tik) — (3) re-scope: n (platform
origin re-fetched at open and close: `0c91421`, unchanged) — (4) user-blocker: n — (5) cap-reached: n
(2 tiks this session) — (6) protocol-stop: n — **Outcome: continue**
**Decisions:** `D-M257x-88-1` (derive the home-bind reset set) · `D-M257x-88-2` (derive the demopatch
target path; fail closed on clone-present-path-absent)
**Side-deliverables:** none — both repairs were the planned scope.
**Routes carried forward** — all Fate 3, with named handlers, and deliberately **not** repaired here: the
scope-creep tripwire fired on the third line of investigation, and re-pinning a `pre_sha256` is exactly the
kind of edit that must not be made in the same breath as discovering it (`demopatch-spec.md`: *"Read it
before adding or re-pinning any patch — a silently-refused perf patch shipped a 76 s members grid for four
releases"*).
- `FIX-M257x-iter88-demopatch-sha-baselines` — the 3 sha-mismatch failures
  (`test_demopatch` ×2, `test_ant_academy` ×1). **Adjudicate before touching**: decide whether the
  self-healing freshness gate makes a drifted whole-file sha *correct*, and whether the assertion or the
  baseline is wrong. Do not re-pin first.
- `FIX-M257x-iter88-back-to-cockpit-revert` — `next-web-back-to-cockpit: revert failed`. A **G5
  self-revert** failure, the most serious of the set: a patch that does not revert leaves a dirty clone.
- `CHECK-M257x-iter88-live-stack-tests` — `test_migrate_race_live` (3) and
  `test_verify::test_demo_3_shifts_ports_and_project` (1). Establish whether these need a live stack; if
  so they belong to a clause-1/2 bring-up iter, not to a static sweep.
- `CHECK-M257x-iter88-unnamed-skips` — the 2 remaining skips
  (`test_interview_flag_patch_m232`, `test_purge`). This iter's own rule says a skip is a hole; these two
  are not yet named.
- The `guard_family` headline's site-path limitation (pinned in a test at iter-87) → **Fate 3**.
**Lessons:** generalised into the protocol doc as **§5 rule 43**. The short form: **a mitigation keyed on
a service NAME dies silently when the service does — and if its tripwire looks the service up first, the
tripwire dies with it, skipping rather than failing.** The pairing is what makes it dangerous: the fix and
the check that guards the fix share the same stale key, so they fail together and in the quietest
available direction. Derive the key from the property that made the service special (here: *carries a
`$HOME` bind*), never from its name — the property outlives the fold, and the name is exactly what the
platform has been deleting for three releases.
