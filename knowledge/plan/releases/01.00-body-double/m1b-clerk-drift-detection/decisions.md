# M1b — decisions

(Decisions recorded as they arise. M1b is automation/config over M0; the drift contract — exit codes
0/1/2 — is in `spec-notes.md`.)

## Adversarial review (Phase 2c — close)
The 2 shell scripts' failure modes are all exercised by `drift-test.sh` (9 assertions): empty/
not-found/unknown-arg usage errors (→3), reformatted-identical DNA (→0, canonical no-spurious-drift),
and the two build-phase regressions (gate-regression == exactly 2 via the built binary; relative
`--new` via abspath). 0 latent issues; `set -uo pipefail` + shellcheck-clean throughout.
