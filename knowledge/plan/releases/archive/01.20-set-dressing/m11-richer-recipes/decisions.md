# M11 — Decisions

_Implementation decisions with rationale. ID scheme: M11-D1, M11-D2, …_

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| M11-D1 | **`/demo-snapshot` is a NEW skill, not a `/demo-seed` extension** (resolves M11-Q2). | Capture is a privileged prod READ with its own safety contract (the read-side `AssertPublicOnly` firewall); seeding is per-stack WRITES (`AssertClean`). The spec already models them as **sibling extensions** (`stack-snapshot` vs `stack-seeding`), and capture runs once-per-release while replay runs per-stack — folding them into one skill would conflate two blast radii. A dedicated skill keeps the UX honest: `/demo-snapshot capture\|replay\|status` maps 1:1 to the `stacksnap` subcommands; `/demo-seed` stays the structural-data step. The full-fidelity curator flow chains them: up → snapshot replay → seed → login. | 2026-06-06 |
| M11-D2 | **Snapshot capture is a maintained golden (cache-first), not a per-curator step** (resolves M11-Q1, aligning with M9a-D3/M9a-Q1). | `store.Resolve` already implements cache-hit-by-default: a cached manifest whose `schema_version` matches → **zero prod read**, replay only. A curator therefore almost always **replays** an existing snapshot; **capture** is the rare maintenance op (first time per release, or on schema drift), and the skill surfaces it as such (replay is the headline verb; capture is the refresh path). This keeps the privileged prod read out of the common curator loop. | 2026-06-06 |
| M11-D3 | **Presets stay structural; the snapshot prerequisite is documented, not encoded in `stack.seed.yaml`.** | Snapshots are **stack-global reference data** (taxonomy/content), independent of the org/user spine — the `TaxonomySnapshotSeeder`/`ContentSnapshotSeeder` DAG nodes only **verify** a prior out-of-band replay, and the blueprint schema has no snapshot field. Adding one would be a false coupling (a preset describes an *org*, not the platform's reference library). So the refresh is a **documented prerequisite** (a header note in each preset + the README full-fidelity flow), not a schema change — and the seeder degrades gracefully (free content refs, empty catalog) when no snapshot is replayed. | 2026-06-06 |
| M11-D4 (hygiene) | **Fixed the stale `stacksnap --help` text** (Fate-1, §5). | The M11 KB-fidelity audit surfaced a help-text-vs-registry drift: the M10 `directus` content surface is registered in the CLI surface registry (`cmd/stacksnap/surfaces.go`) but the `--help` usage still said "framework (M9a/M9b)" and listed only `taxonomy \| reference-toy`. A doc-side false-omission (the surface works; the help under-advertised it). Landed now (last milestone, no sibling to defer into): bumped the tag to M9a/M9b/M10 + listed `directus`. Extensions change → folds into the `stack-snapshot-m11` tag at close. | 2026-06-06 |

## Deferrals — three-fate ledger
_(none)_ — all M11 work landed Fate-1. No items punted; no escape-hatch invoked. The only candidate (DEF-M10-01
S3 blob bytes) is already Fate-2-covered by the v1.3 cloud-store roadmap-vision seed (recorded in M10's decisions),
and M11's media-ref handling honors that boundary (refs are the floor; blob bytes are v1.3).
Re-confirmed GREEN at the M11 close deferral re-audit:
[`audit-deferrals/deferral-audit-2026-06-06-m11-close.md`](audit-deferrals/deferral-audit-2026-06-06-m11-close.md).

## Adversarial review (close Phase 2c)
The close ran an external-perspective pass over each non-trivial M11 surface. M11 ships **no new
production code path** — only a one-line help-text fix, comment-only preset headers, two regression
test files, and markdown. The scenarios therefore target the *contracts* the milestone pins and the
docs it ships, not new runtime behavior:

1. **A new snapshot surface is registered but the `--help` text is never updated** (the exact M10→M11
   drift class). → `TestHelp_NamesEveryRegisteredSurface` drives off `surfaceNames()`, so ANY
   future registered-but-unlisted surface fails, not just `directus`. Mutation-proven: dropping the
   directus line / reverting the tag → both `TestHelp_*` FAIL. **Handled.**
2. **The `/demo-snapshot` skill (sibling rosetta repo) documents a `stacksnap` flag the parser
   doesn't have, or a parser rename strands a documented flag.** → `TestDocsFlagsExistInParser`
   probes the live parser for every documented flag; `TestDroppedDumpFlagStaysGone` asserts the
   M9b-D9-removed `--dump` never returns. The cross-repo contract is mirrored as an explicit list
   with a lockstep comment. **Handled** (the asymmetry that bites — code drifting from the docs — is
   what's probed).
3. **A preset edit (the M11 §1 comment header, or any future change) silently breaks a shipped
   preset past the strict `KnownFields(true)` loader, surfacing only at a curator's runtime.** →
   `TestShippedPresets_ParseStrictAndValidate` loads every shipped file through the CLI's exact
   Load+Validate path. Mutation-proven: an unknown field → FAIL; a pure comment line → PASS (proving
   the comment-only header style is safe). `TestShippedPresets_SizesAreDistinctAndOrdered` guards the
   small<mid<large contract the README size table advertises. **Handled.**
4. **A documented `--source` value (`dump-ingest`/`primary-read`) drifts from the real
   `source.Kind` constants.** → `TestDocumentedSourceKindsAreReal` resolves each documented kind AND
   asserts the literal string matches the constant. **Handled.**

No scenario revealed an unhandled risk; all are pinned by the harden-pass regression tests, each
mutation-verified to bite. No escape-hatch / documented-risk acceptance was needed.

## Decision triage (close Phase 5)
- **M11-D1** (`/demo-snapshot` is a NEW skill) → already blended: the rationale (capture = privileged
  prod READ vs seeding = per-stack WRITE; sibling extensions; replay-headline/capture-rare) is woven
  into the `/demo-snapshot` SKILL.md, `recipe-snapshot-world.md`, and the README flow. **Stays
  archive** (the "why" is in the corpus; full options-considered list stays here).
- **M11-D2** (snapshot capture is a cache-first golden, not a per-curator step) → already blended into
  `recipe-snapshot-world.md` ("almost always a cache-hit", `store.Resolve` zero-prod-read) + the
  skill body. **Stays archive.**
- **M11-D3** (presets stay structural; snapshot prerequisite documented, not encoded) → already
  blended into each preset's FULL-FIDELITY PREREQUISITE header + the README preset paragraph.
  **Stays archive.**
- **M11-D4** (the `stacksnap --help` hygiene fix) → maintainer-only code-hygiene record; the fixed
  help text is self-documenting and pinned by `TestHelp_*`. **Stays archive.**

No decision needs a net-new knowledge edit — all the user-facing "why" already flowed into the corpus
during build. Nothing maintainer-only was promoted.
