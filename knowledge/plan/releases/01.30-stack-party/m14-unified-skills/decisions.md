# M14 — Decisions

_Implementation decisions with rationale. ID scheme: M14-D1, M14-D2, …_

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| M14-D1 | Hard rename, no back-compat aliases | user 2026-06-07 — clean break; update every in-repo reference | 2026-06-07 (user) |
| M14-D2 | `dev-up` consolidates **operation**, not the build itself (Q1) | The heavy first-time machine build (tool install, org clone, ~15-25min) stays the `setup_guide.md` one-time path that `dev-up` *drives*; `dev-up N` (N≥1) spins additional isolated copies via `dev-stack up`. The `dev-stack` CLI already documented this boundary, so `dev-up` is a skill-level consolidation of setup+start, not a re-implementation. | 2026-06-07 |
| M14-D3 | Target-detection UX = **positional `dev-N\|demo-N`** at the skill level (Q2) | The skills take the stack as a positional (e.g. `/stack-seed dev-1`) and resolve it to the CLI's `--stack` flag + the per-stack-role clone (`stack-dev/` vs `stack-demo/`). Matches the existing `/demo-*` skill ergonomics; the raw CLIs keep `--stack`. | 2026-06-07 |
| M14-D4 | **No generic `stack-up`/`stack-down`** — lifecycle stays type-specific (Q3) | Bring-up differs by kind (demo = Clerkenstein-injected + disposable + clone-at-tag; dev = real-Clerk + set-dressed + tracks main). One generic up/down would hide that. The *ops* (list/seed/snapshot/update) are kind-agnostic and DO converge; the *lifecycle* (`dev-up/down` ∥ `demo-up/down`) stays paired-but-separate. | 2026-06-07 |
| M14-D5 | `--preset` is a **skill-level shorthand**, not a `stackseed` CLI flag | PR-review finding D1: the binary only knows `--seed <path>`. The skills keep the `--preset NAME` UX (established by the old `/demo-seed`) but the SKILL now explicitly maps it to `--seed presets/NAME.seed.yaml` so an automated invocation never passes a bogus flag. | 2026-06-07 |
| M14-D6 | CHANGELOG: new **Unreleased v1.3** entry; v1.1/v1.2 dated entries left immutable | Keep-a-Changelog convention — historical release entries record what shipped under the old names at that time; the rename is documented forward in the v1.3 Added/Changed/Removed, not by rewriting history. Same for `knowledge/plan/` archive + state/roadmap history. | 2026-06-07 |

## Open at design — all resolved during build
- M14-Q1 → **M14-D2** (consolidate operation, not the build).
- M14-Q2 → **M14-D3** (positional `dev-N\|demo-N` → CLI `--stack`).
- M14-Q3 → **M14-D4** (no generic up/down — lifecycle stays type-specific).

## Notes
- The `stack-snapshot` **skill** namespace is distinct from the `stack-snapshot` **extensions section** name
  (`rosetta-extensions/stack-snapshot/`) — the skill operates the `stacksnap` CLI that section builds. Called
  out in the skill body to avoid conceptual collision (per the milestone brief).
- Extensions-clone references updated on `main` (commit `b37e831`): dev-stack header + README, demo-stack
  GUIDE, stack-verify run.sh — doc/comment-only, no CLI/section-name change, shellcheck CLEAN.
