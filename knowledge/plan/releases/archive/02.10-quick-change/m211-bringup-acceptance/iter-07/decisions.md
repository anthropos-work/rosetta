# iter-07 — decisions

## D1 — dev-up SKILL container count 12 → 11 (skiller container removed by the merge)
The v2.1 merge removed the `skiller` service/container; the merged 4-subgraph graphql profile brings up
**11** `anthropos-*` containers (verified against warm `docker ps`), not 12. The dev-up SKILL's build-step
and health-verify both asserted "12 containers" — a stale pre-merge count that would fail a cold `/dev-up`
health check on a phantom 12th container. Corrected both spots + annotated why (skiller merged into app's
`public` schema). Corpus-only, no platform/rext edit.

## D2 — reset-db extensions re-bootstrap: corpus, not rext (M25-D9 class)
M208 found `make reset-db` re-runs migrations against a wiped DB with no `extensions` schema → app/cms
`vector(1536)` + gin-trigram migrations fail cold. Fix surface = the DOC the `/dev-up` SKILL executes
(`setup_guide.md`), NOT a rext tooling change, because: (a) the platform Makefile that bundles reset-db's
auto-migrate is un-editable (zero platform edits), and (b) the main `/dev-up` (N=0) drives the documented
manual steps via `make`, not the rext `dev-stack` tooling (that's for `dev-N`, N≥1). Added a "⚠ Cold-reset
ordering" block documenting the re-create-schemas + reload-policy + re-migrate recovery. The first-time cold
build already dodges the race (creates schemas between `make up` and `make migrate`); only reset-db's bundled
auto-migrate hits it.
