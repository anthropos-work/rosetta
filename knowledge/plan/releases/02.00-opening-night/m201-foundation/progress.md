# M201 Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in [`overview.md`](overview.md).

## Section checklist

- [ ] **(1) Manifest model + light validator** — the manifest format (Products → Stories → Use Cases →
  Playthroughs) + a validator enforcing **both-way id integrity** + **precondition-coverage** (every declared
  seed/preconditions resolves to a named seeded world), **datadna-gated**.
- [ ] **(2) Per-surface locator/landmark page-object layer** — the shared registry every Playthrough imports
  (semantic locators + a landmark registry for ambiguous surfaces), **1 surface to start**; re-pin O(surfaces).
- [ ] **(3) Dedicated decoupled seed preset** — test data ≠ demo data; **spans entitlement tiers +
  multi-org-private**; covered by the `datadna` gate.
- [ ] **(4) Reset-to-seed lifecycle + serial-default runner** — per-suite reset via the real `--reset` path
  (additive re-seed FORBIDDEN as a reset); `workers: 1`, `fullyParallel: false`.
- [ ] **(5) 4-state reporting map** — passing / failing / unimplemented / unimplementable-without-platform-edit
  (the last escalates, never edits the platform).
- [ ] **(6) One trivial proof Playthrough** — login → /profile → assert hero identity (the foundation smoke
  test).
- [ ] **Docs** — `corpus/ops/demo/playthroughs.md` **(NEW)** graduates the spec-draft (becomes the M202/M203
  `iteration_protocol_ref`); update the `demo/README.md` index + `CLAUDE.md` docs list.

**Status:** `planned` — not yet started. Next: `/developer-kit:build-milestone` (M201). Tooling + docs only —
zero platform-repo edits; the new `playthroughs` rext section is authored + tagged per the tooling policy.
