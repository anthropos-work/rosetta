# M202 Spec Notes

Technical notes accumulate here during build. The authoritative design lives in [`overview.md`](overview.md) +
the consolidated capability spec
[`knowledge/plan/spec-drafts/playthroughs/spec.md`](../../../spec-drafts/playthroughs/spec.md) (v0.3). M202 is
**tooling + docs only — zero platform-repo edits**; the platform stays read-only (an un-drivable surface
escalates via `unimplementable-without-platform-edit`, it never edits the platform). The manifest **content** this
milestone's validator consumes is the **M201 manifest corpus** (prose-only, authorable in parallel).

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
- **rext authoring copy** (`.agentspace/rosetta-extensions`, branch `main`): the new `playthroughs` section is
  committed at HEAD `e77e176` (§1–§6, `79df988..e77e176`), tree clean. The tag `opening-night-m202` is applied
  at **CLOSE**, not at build (per the tooling policy — a milestone tags on close).
- **consumption clone** (`stack-demo/rosetta-extensions`): re-pinned at close (not during build).
- **corpus m202 branch** (`m202/foundation`): the `Delivers → playthroughs.md` runbook + the cross-references +
  the milestone records land here (this commit set).
- The **M201 manifest corpus** YAML lands product-by-product in this `playthroughs/manifest/` dir in M203/M204;
  M202's `manifest/profile.yaml` carries only the single foundation proof use case.

## Open questions (carry into the build; record resolutions in decisions.md)
- Harness home: the `playthroughs` section's own dir vs nesting under `stack-verify/e2e/` — decide against how
  much it reuses the foundation's plumbing.
- Manifest format: YAML descriptor (mirroring `stack.stories.yaml`) vs TS — decide against the validator + the
  page-object import ergonomics.
- The 1 starting surface for the page-object layer — pick the one the trivial proof Playthrough (/profile)
  needs.
- The dedicated seed's relationship to the demo seed (starting point) — how the entitlement-tier +
  multi-org-private span is expressed in the preset.

## Delivers — `corpus/ops/demo/playthroughs.md` (NEW) — AUTHORED
DONE: graduated the spec-draft into the corpus runbook — dual-level (PM + engineer), citing the actual rext
`playthroughs/` files: the capability + P1–P8 principles, the manifest model (`manifest/manifest.go`) + light
validator (`manifest/validator.go` + `cmd/ptvalidate`), the per-surface page-object layer
(`e2e/lib/page-object.ts` + `profile-page.ts`, re-pin O(surfaces)), the hero login (`e2e/lib/hero-login.ts`
reusing the M37 cockpit seat-switch, never forked), the dedicated decoupled seed (`seed/pt-world.seed.yaml` +
`seed/seed-worlds.yaml`) + reset-to-seed + serial runner (`e2e/run-playthroughs.sh` + `playwright.config.ts`),
the 4-state map (`report/report.go` + `cmd/ptreport` + `report/unimplementable.yaml`), the proof Playthrough
(`e2e/tests/profile-identity.spec.ts`), and the layering finding (M202-D4). Includes an explicit **"The
iteration protocol (for M203/M204)"** section — the declare→extend-seed→page-object→run→triage→re-measure loop
+ the integration-boundary posture — so it IS the `iteration_protocol_ref` those coverage milestones follow.
Cross-referenced from CLAUDE.md doc-index + `demo/README.md` index + `coverage-protocol.md` (the function
sibling of the presence sweep).

## Pre-flight audits — Section 1 (Manifest model + validator)
KB-fidelity audit (2026-07-01): **GREEN** — report at
[`kb-fidelity-audit.md`](kb-fidelity-audit.md). Topic→doc→code triples (verified ALIGNED):
- reset/`--force` → `seeding-spec.md` + `idempotency.md` → `stack-seeding/cmd/stackseed/main.go::doReset`.
- isolation guard → `seeding-spec.md` §84-106 → `stack-seeding/isolation/isolation.go` (`Guard.CheckWrite`,
  `AuditLog.AssertClean`).
- cockpit-login → `coverage-protocol.md` §5.6 → `stack-verify/e2e/lib/cockpit-login.ts` (`selectSeat`+`loginAs`).
- datadna gate → `seeding-spec.md` §data-DNA → `stack-seeding/cmd/datadna` + `stack-seeding/dna/`.
The one BLIND-AREA (`corpus/ops/demo/playthroughs.md`) is the milestone's own declared `Delivers →` — authored
in the Docs section, not an unfilled blind area.
