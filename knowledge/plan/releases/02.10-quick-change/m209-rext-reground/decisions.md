# M209 — Decisions

_Implementation choices with rationale, logged as they are made._

## Phase 0b — KB-fidelity audit: YELLOW (expected)
Report: `kb-fidelity-audit.md`. Pre-flip doc staleness only; no blind areas, no load-bearing stale claim for
M209's code re-ground (M209 grounds against M208's empirically-verified merged schema, not the doc bodies).

- **KB-1** — `corpus/ops/snapshot-spec.md` describes the taxonomy surface as `skiller.*` (19 mentions). STALE.
  Fate-2: **M210** (`m210-corpus-reground`) explicitly owns the snapshot-spec body-flip. No new deferral.
- **KB-2** — `corpus/ops/seeding-spec.md` prose ("public skiller catalog"). STALE-prose. Fate-2: **M210**.
- **KB-3** — `corpus/ops/safety.md` firewall evidence row `skiller.skills 42,763 public`. STALE (schema + count,
  now ~42,790). Fate-2: **M210** names it. Firewall CODE unchanged (schema-agnostic `organization_id IS NULL`).

## Risk 1 — digest narrowing: gate on `CapturesStructure`, not a hardcoded table set
The staleness digest (`pg.SchemaVersionSQL`) is narrowed to the surface's OWN enumerated tables
(`Surface.VersionTables`), threaded through `Capturer.SchemaVersion(ctx, schema, tables)` + all CLI adapters.
Discriminator: a STRUCTURE-bearing surface (directus) returns `nil` → whole-schema digest (its content model is
a DYNAMIC set of user collections; a new collection MUST invalidate its cache — verified against
`directus/structure.go`'s "every non-`directus_` table" capture). A row-only surface (taxonomy) returns its 10
tables → the digest is stable against the merged `public` monolith's unrelated app-migration churn (the Risk-1
cache thrash). Same helper used at capture (BuildPlan), the autoprovision re-probe, AND main.go's target-probe,
so capture-side and target-side digests always match.

## Risk 2 — capture column list: verified, NO code change needed
Verified every enumerated taxonomy column against merged-prod (`docker anthropos-postgresql-1`, schema only —
0 rows): all 10 tables present, every column name unchanged, `small_embedding3` still the vector column,
`ts_search` still correctly EXCLUDED. The `embedding → small_embedding3` rename the overview flagged did NOT
happen. The `extensions.`-qualified vector/opclass types never surface in the tooling: the capture SELECT is
names-only (`buildPublicSelect` → COPY renders vectors as text) and replay is `REINDEX TABLE <t>` (rebuilds
existing indexes by name). So the schema-flip alone is sufficient; the column lists are correct as-is.

## New post-capture MinRows floor (item 4's assertion half)
Added `TableSpec.MinRows` (a general capture-time sanity floor) + a check in `capture.Run` that aborts before
ANY store write if a table captured fewer rows. Set `MinRows: 40000` on `skills` (~42,790 public in prod; floor
below to tolerate churn). Catches empty/wrong-schema/under-capture. Over-capture is separately caught by the
firewall's `AssertCaptured` (0 tenant rows) — so the floor is one-sided by design. `AssertPublicOnly` kept.

## Recapture — Fate-3 → M211 (no local capture source)
The actual recapture from merged-prod is a DATA op, operationally gated: values-blind-confirmed NO `marco_read`/
prod-read Postgres DSN in `.agentspace/secrets` or `platform/.env`; the merged `stack-dev` Postgres has 0
taxonomy rows; the `postgres` MCP returns JSON not COPY bytes. Investigated Fate-1 (a real capture) — genuinely
infeasible here. Routed Fate-3 to **M211** (its exit gate already names the recapture — Fate-2 coverage — plus I
pinned a "Pre-surfaced recapture prerequisite" note in `m211/overview.md` so its first tik doesn't re-discover
the missing-source blocker). The CODE re-ground (tooling READY to capture/replay `public.*`) is M209's core and
is complete. Also corrected M211's stale ~42,763 → ~42,790 count in lockstep.

## Scope boundary — conceptual bare-word "skiller" comments left; done-bar met
The done-bar is "0 `skiller.<table>` queries in production" — VERIFIED 0 (dotted swap + grep). A handful of
CONCEPTUAL bare-word "skiller" mentions remain in seeder/dna code comments ("resolves in skiller", "the skiller
hierarchy", the `↔skiller` closure-probe label coupled to a production error string + its test assertion). These
are not schema-existence claims and not queries; rewording them is M210 doc-narrative territory. The factually-
WRONG "missing skiller schema" comments WERE fixed. `stack-verify/repos/run.sh` keeps a dead `skiller)` repo-test
case (unreachable — repos.yml has no skiller; harmless; out of chartered scope). `stack-secrets` synthetic
`skiller`/`BUNNY` waived-repo test fixtures left (overview: optional; self-contained; tests green).
