# M202 Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in [`overview.md`](overview.md).

## Section checklist

- [x] **(1) Manifest model + light validator** — the manifest format (Products → Stories → Use Cases →
  Playthroughs) + a validator enforcing **both-way id integrity** + **precondition-coverage** (every declared
  seed/preconditions resolves to a named seeded world), **datadna-gated**. rext `manifest/` + `cmd/ptvalidate`
  (tests -race green, `79df988`).
- [x] **(2) Per-surface locator/landmark page-object layer** — the shared registry every Playthrough imports
  (semantic locators + a landmark registry for ambiguous surfaces), **1 surface to start** (`/profile`); re-pin
  O(surfaces). rext `e2e/lib/{page-object,profile-page}.ts` + `hero-login.ts` (tsc green, `5353396`).
- [x] **(3) Dedicated decoupled seed preset** — test data ≠ demo data (`pt-world`, 2 private orgs); **spans
  entitlement tiers + multi-org-private**; covered by the `datadna` gate. rext `seed/pt-world.seed.yaml` +
  `seed/seed-worlds.yaml` (seeding-validator VALID, `de55b9b`).
- [x] **(4) Reset-to-seed lifecycle + serial-default runner** — per-suite reset via the real `--reset` path
  (additive re-seed FORBIDDEN as a reset), N=0-guarded; `workers: 1`, `fullyParallel: false`. rext
  `e2e/run-playthroughs.sh` + `playwright.config.ts` (`fcf45ad`).
- [x] **(5) 4-state reporting map** — passing / failing / unimplemented / unimplementable-without-platform-edit
  (the last escalates, never edits the platform). rext `report/` + `cmd/ptreport` + `unimplementable.yaml`,
  distinct glyphs (`ed0408a`).
- [x] **(6) One trivial proof Playthrough** — login → /profile → assert hero identity (the foundation smoke
  test). rext `e2e/tests/profile-identity.spec.ts` (`@pt:pt-profile-identity`) — **GREEN on demo-1** (`e77e176`;
  anchor-story layering fix M202-D4).
- [x] **Docs** — `corpus/ops/demo/playthroughs.md` **(NEW)** graduates the spec-draft (IS the M203/M204
  `iteration_protocol_ref`); cross-referenced from the `demo/README.md` index + `coverage-protocol.md` (function
  sibling) + `CLAUDE.md` docs list.

**Status:** `complete` — all 6 sections + the runbook deliverable landed; proof Playthrough GREEN on demo-1.
Tooling + docs only — zero platform-repo edits. rext authoring copy @ `e77e176` (§1–§6, `79df988..e77e176`),
tree clean; the `opening-night-m202` tag + the consumption-clone re-pin happen at CLOSE. Next:
`/developer-kit:harden-milestone` (optional) then `/developer-kit:close-milestone`.
