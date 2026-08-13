# iter-88 — decisions

## `D-M257x-88-1` — the volumes-reset set is DERIVED from the compose text, not a service name

`gen_injected_override.py` cleared `$HOME`-rooted binds with `if name == "jobsimulation":`. The service
was deleted at `d11a403`; `838d907` moved the identical `$HOME/.aws/credentials` bind onto `backend`.
The mitigation therefore guarded nothing while the hazard sat on the stack's most important container.

Replaced with `platform_topology.services_with_only_home_binds(platform_dir)`.

**Three properties, each chosen against a specific alternative:**

| choice | rejected alternative | why |
|---|---|---|
| parse the **raw** compose text | read the resolved `docker compose config` | after expansion, `$HOME/.aws/credentials` and `…/stack-demo/platform/data/postgresql` are both absolute paths under the user's home. The intent exists only before expansion |
| **ALL** volumes are home binds | **ANY** | `!reset null` clears the whole list. A service mixing a home bind with a volume it needs must be excluded, not blanket-cleared — the original comment asked for this and nothing enforced it |
| **fail closed** on an empty derivation | pass when nothing matches | "nothing to check" and "nothing is there" are the same observation from the check's side and opposite verdicts from the operator's (§5 rule 14) |

Non-fatal when `platform_dir` is absent or unreadable — a synthetic dir is passed by the guard and several
unit tests, and failing the generator over an unreadable clone would trade a broken mount for a broken
bring-up.

**The tripwire was rewritten to the same predicate**, so the generator and its test cannot disagree; it
additionally pins that `backend` **is** in the set and `jobsimulation` is **not**, so a fold reversal is
loud rather than silent.

## `D-M257x-88-2` — the demopatch target path is DERIVED from the manifest, and the skip is fail-closed

`TestRealManifestsAgainstTheRealClone._check` took `(manifest, target_rel)` — the path restated beside the
manifest that already declares it. M254 re-pointed the aireadiness manifest into
`internal/aireadiness/readiness.go`; the test's copy stayed at `internal/workforce/ai_readiness.go`;
`os.path.isfile` failed; the test skipped, and kept skipping.

`_check(manifest)` now reads the manifest's own `path:` (a line scan — `path:` is a top-level scalar and
this section carries no YAML dependency, consistent with `repos_yml.sh`'s machine-readable-fields-only
approach). The skip is split:

- **clone directory absent** → skip. A real reason; not every box has the clone.
- **clone present, manifest path missing** → **FAIL**. That is a dead patch, and it is the drift the test
  exists to catch.

Re-run at `0c91421`: the aireadiness anchor **resolves**, so M254's re-point was correct and had simply
been unverified for four releases. `stack-injection` went 2 skips → **0**.

A third assertion pins the derivation itself, so a future move has to change the manifest — which the
appliers read — rather than a literal in a test, which nothing reads.

## `D-M257x-88-3` — `test_emits_volumes_reset_for_jobsimulation` was pinning the defect in place

The test asserted the generator's **source** contained `if name == "jobsimulation":`. Had anyone tried to
remove the dead literal, this test would have failed and argued for keeping it. A test pinned to an
implementation's service literal is a second copy of the thing that goes stale — and this one had teeth.

Re-pointed at the **property** (the reset is emitted, keyed on a derived set) with an explicit
`assertNotIn` on the old literal, so the stale key cannot come back quietly. The behavioural checks moved
to assertions against the real compose file.

## `D-M257x-88-4` — the un-run sections are named as a hole, not implied

iter-87 ran only `stack-core`. This iter ran `stack-injection` fully and `stack-verify` partly;
`dev-stack` and `demo-stack` were still running at close. **Routed forward explicitly rather than left
unstated** — the iter's own finding is that unexercised checks go quiet, so an unfinished sweep recorded
vaguely would be the same defect one level up. `stack-verify`'s single ERROR is pre-existing, needs a live
demo-3 stack, and has no import path from this iter's diff (verified, not assumed).
