# M201 Spec Notes

Technical notes accumulate here during build. The authoritative design lives in [`overview.md`](overview.md) +
the consolidated capability spec
[`knowledge/plan/spec-drafts/playthroughs/spec.md`](../../../spec-drafts/playthroughs/spec.md) (v0.3). M201 is
**tooling + docs only — zero platform-repo edits**; the platform stays read-only (an un-drivable surface
escalates via `unimplementable-without-platform-edit`, it never edits the platform).

## Scope (see overview.md for the authoritative cards)
- **(1) Manifest model + light validator** — both-way id integrity + precondition-coverage + datadna-gated.
  TODO (build): the manifest format (Products → Stories → Use Cases → Playthroughs) + the validator wiring.
- **(2) Per-surface locator/landmark page-object layer** — 1 surface to start. TODO (build): the shared registry
  every Playthrough imports (semantic locators + landmark anchors); re-pin is O(surfaces), not O(tests).
- **(3) Dedicated decoupled seed preset** — entitlement tiers + multi-org-private. TODO (build): the preset +
  its `datadna` coverage.
- **(4) Reset-to-seed lifecycle + serial-default runner** — `--reset` path (NOT additive re-seed); `workers: 1`,
  `fullyParallel: false`. TODO (build).
- **(5) 4-state reporting map** — passing / failing / unimplemented / unimplementable-without-platform-edit.
  TODO (build): the report reconciles the manifest against results.
- **(6) One trivial proof Playthrough** — login → /profile → assert hero identity. TODO (build).

## Reuse paths (the shared M42 e2e foundation)
Built on the foundation it shares with `stack-verify` (§5.6) — cite/wire these:
- `stack-demo/rosetta-extensions/stack-verify/e2e/lib/cockpit-login.ts` — hero login (the M37 handshake).
- `stack-demo/rosetta-extensions/stack-verify/e2e/lib/section-assert.ts` — per-section assertion helpers.
- `stack-demo/rosetta-extensions/stack-verify/e2e/lib/empty-states.ts` — placeholder/empty-state detection.
- `stack-demo/rosetta-extensions/stack-verify/e2e/lib/coverage-manifest.ts` — the manifest-driven section model.
- `stack-demo/rosetta-extensions/stack-seeding/` — the seeding machinery (`stackseed` / `--reset` / `datadna`).

## Tag / two-repo state
TODO (build): the rext authoring copy (`.agentspace/rosetta-extensions`) commit + tag for the new `playthroughs`
section; the consumption clone (`stack-demo/rosetta-extensions`) checkout; the corpus m201 branch (the
`Delivers → playthroughs.md` runbook + plan files).

## Open questions (carry into the build; record resolutions in decisions.md)
- Harness home: the `playthroughs` section's own dir vs nesting under `stack-verify/e2e/` — decide against how
  much it reuses the foundation's plumbing.
- Manifest format: YAML descriptor (mirroring `stack.stories.yaml`) vs TS — decide against the validator + the
  page-object import ergonomics.
- The 1 starting surface for the page-object layer — pick the one the trivial proof Playthrough (/profile)
  needs.
- The dedicated seed's relationship to the demo seed (starting point) — how the entitlement-tier +
  multi-org-private span is expressed in the preset.

## Delivers — `corpus/ops/demo/playthroughs.md` (NEW)
TODO (build): graduate the spec-draft into the corpus runbook (the capability, the manifest model, the
page-object layer, the dedicated-seed + reset-to-seed lifecycle, the serial-default runner, the 4-state reporting
map). This becomes the `iteration_protocol_ref` for M202/M203.
