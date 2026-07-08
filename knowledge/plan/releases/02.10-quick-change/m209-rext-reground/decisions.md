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

## Adversarial review (close Phase 2c)
No NEW adversarial scenario surfaced at close — the harden phase (rext `42ad600`/`72a5259`/`2f06e78`) already
exercised the milestone's non-obvious failure modes and recorded them in `progress.md § M209: Hardening`:
- **Schema-const revert** — the pre-existing identity check `s.Schema != Schema` was tautological (both sides
  the same const); harden added a literal-value guard (`Schema == "public"` + every `TableSpec.Schema`/PublicVia
  is `public.`-qualified) so a flip back to `"skiller"` is caught. The real regression guard on THE change.
- **MinRows off-by-one / empty-schema** — floor is one-sided (`rows>=MinRows`), aborts in-loop before any store
  write AND before the post-loop `AssertCaptured` leak gate; the 0-row wrong-schema capture trips it.
- **Digest scope leak** — a STRUCTURE-bearing surface (directus) returns `nil` → whole-schema digest (a new
  dynamic collection still invalidates); a row-only surface (taxonomy) returns its 10 tables → stable vs the
  merged `public` monolith's app-migration churn. Both sides of the cache-key comparison use the same helper.
- **Seeder query-shape** — a recorder now asserts `taxonomy_snapshot`'s countConn counts
  `public.skills WHERE organization_id IS NULL` (never `skiller.`). 0 bugs, 0 flakes across all three passes.

## D-close-2 — rext stack-seeding/README test-count drift (pre-existing) → routed, not fixed in-place
Close Phase-4 handbook reconciliation surfaced `stack-seeding/README.md:106` = "496 test funcs across 8
packages" vs authoritative `go test -list` = **788 across 13 packages**. **Pre-existing, cross-release drift:**
the README count was last reconciled at **M41** (commit `0346113`, v1.10 "method acting"); the gap accumulated
across v1.10b (M47–53) + v2.0 (M201–204) + v2.1 seeder additions. **M209 did NOT touch this README** (git-verified)
and its own stack-seeding test delta was small (renamed matchers + the taxonomy_snapshot harden test). The file
lives in the **rext repo, mandate-frozen at HEAD `2f06e78`** for this close (tag `quick-change-m209`→`2f06e78`,
clean tree required per the orchestrator + top-of-prompt dirty-tree ban) — an in-place fix would require a rext
commit that moves HEAD off `2f06e78` (violating the explicit tag mandate) or a dirty rext tree (blocker). So it
is **recorded + routed** to the next legitimate rext advance — the **v2.1 rext roll at `/developer-kit:close-release`**
(the rext repo's chartered v2.1 re-tag) or an earlier **M211 rext re-tag** — whichever advances rext past
`2f06e78` first. Nice-to-have doc hygiene; not load-bearing; not an M209 deliverable; does not gate the close.
Tracked as a standing rext-doc-hygiene item (see state.md Standing backlog).
