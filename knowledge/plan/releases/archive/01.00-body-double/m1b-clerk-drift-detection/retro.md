# M1b — Retro: Clerk drift detection

## Summary
A tight B-milestone: mechanized Clerk drift detection by reusing M0 wholesale — `gate.sh` +
`drift-check.sh` (exit-code contract 0/1/2/3) + a weekly CI alignment gate + a 9-assertion
`drift-test.sh`, all in the clerkenstein repo. Makes a Clerk version bump a flagged, mechanical event
("follow updates within minutes") instead of a silent break. 2 sections + 1 harden pass, 0 close findings.

## Incidents this cycle
- **None** (0 bugs, 0 flakes). Two build-phase gotchas, both caught + fixed during S1 and *pinned* as
  regressions in harden: `go run` squashes a non-zero exit to 1 (→ build the alignctl binary so the
  exact 0/2 propagates); a relative `--new` path broke under the script's `cd` (→ resolve to absolute).

## What went well
- A B-milestone really is *config over the framework* — no new measurement code, just two scripts + a
  CI YAML + a runbook. M0's `alignctl dna diff` + `--gate` did all the work.
- The harness's edge cases (reformatted-identical DNA → no spurious drift) double as an integration
  test of M0's canonical DNA-diff at the M1b layer.

## What didn't
- The two shell gotchas (go-run exit-squash, cd+relative-path) are easy to miss without the exit-code
  test — which is exactly why the harden harness exists now.

## Carried forward
- None. The CI workflow is honestly-unverified on a real GHA runner (no runner here) — it'll be
  exercised when clerkenstein is a pushed repo / the demo stacks run (v1.1).

## Metrics delta
9 shell assertions (flake 5/5), shellcheck clean, Go suite unchanged. Full figures: [metrics.json](metrics.json).
