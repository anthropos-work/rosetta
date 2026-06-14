---
iter: 06
milestone: M21
iteration_type: tik
status: closed-fixed
created: 2026-06-11
---

# M21 iter-06 — tik (fifth tik under TOK-01): code-ify the structure capture (capture-side core)

Under TOK-01. Begin `STRUCT-M21-codeify` — the critical path that flips the gate from *demonstrated* (iter-05) to
*met by tooling*. The full code-ification spans capture (this iter) → apply-via-stacksnap (iter-07) → serve
registry/permissions (iter-08); executing-at-bring-up is M22 (out of M21 scope).

## Active strategy reference
TOK-01 + M21-D7 (option A) + the M21-D8/D9 recipe (structure = the 26-collection DDL **incl. PRIMARY KEYs**).

## Re-survey (Phase 1 Step 0)
furthest-passing-stage = 6 (demonstrated, iter-05) but gate not met-by-tooling. Route `STRUCT-M21-codeify` is the
critical path. The proven artifact is `iter-04/structure.sql` + `iter-05/pks.sql`.

## Cluster / target identified
The capture-side CORE: a data model for the structure artifact in the snapshot + a generator that produces the
schema-structure SQL (CREATE SEQUENCE + CREATE TABLE + ADD PRIMARY KEY for every directus user collection) from a
source over a conn. Captured **dynamically** (all `directus.*` base tables NOT matching `directus_*`), so it is
version-robust and not a hardcoded 26-list.

## Hypothesis / deliverables
1. `manifest`: an **additive** `Structure *StructureArtifact` field (payload + sha256 + statement count) — no
   FormatVersion bump (the `Predicate` precedent), zero impact on taxonomy/reference manifests.
2. `directus/structure.go`: the catalog SQL (DDL / PK / sequences) as audited consts + `BuildStructureSQL(ctx, runner)`
   that runs them and assembles the ordered structure script (sequences → tables → PKs).
3. `pg`: a `QueryRowString` helper (the runner the generator needs).
4. Unit tests: the assembly/ordering against a fake runner + manifest additive-field validation + back-compat (a
   manifest without `structure` still parses).

## Expected lift
No furthest-stage change (the metric is already at 6-demonstrated; this iter builds the tooling that will let stage 3
be stacksnap-applied). This is **code-ification progress under the same gate** — graded on the planned deliverables
landing clean (closed-fixed), not a stage-ordinal move. The ordinal converts to "met (automated)" once iter-07/08 wire
apply+serve.

## Phase plan
1. Add the manifest additive field + validation + back-compat test.
2. Add `pg.QueryRowString`.
3. Write `directus/structure.go` (consts + builder) + `structure_test.go` (fake-runner assembly test).
4. `go build/vet/test ./...` green. (Live capture-wiring + apply = iter-07.)

## Escalation conditions
- If the additive manifest field can't stay back-compatible (older manifests fail to parse), stop + reconsider the
  artifact-storage design (sibling-file vs field) — that's the milestone's open question; record + route.

## Acceptable close-no-lift outcomes
- n/a — this is a build iter; closed-fixed when the capture-side core lands + tests green.

## Test discipline
`go build ./... && go vet ./... && go test ./...` across stack-snapshot (the suite is the gate for this iter).
