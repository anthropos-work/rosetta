# `/demo-update` — Known Findings surfaced during M-B live-run

> **Status:** Draft · spec-draft · 2026-07-10 · sibling of [`spec.md`](spec.md) + [`spec-progress.md`](spec-progress.md)
>
> Two orthogonal defects surfaced by M-B's closure gate on the demo-1 live-run. Neither was
> introduced by M-B — the gate did its job (P1 win). Recorded here so they are not lost when
> M-C ships. Both are **out of the M-B lane**; the real fixes belong to `stack-snapshot` and
> `stack-seeding` respectively.

## Finding A — `stacksnap` taxonomy capture is stale vs post-seed FK indexes

**Symptom.** `stacksnap replay --surface taxonomy --stack demo-1` returned `rc=5` (cache miss)
on the M-B live-run — the schema digest the surface expected was
`5afc0bccf1df7ef538b643321fc6362f`, but the cached capture carries
`c75ce94d6a8021cad2915ddb4fb3dd4d`.

**Root cause.** `/demo-up 1`'s post-seed pass adds foreign-key indexes on
`membership_skills` and `membership_tags` after the initial replay. Those indexes shift the
schema fingerprint that `stacksnap` uses to key its capture cache. The cached capture
predates them, so any replay attempt *after* a demo has been fully brought up will fail with
a cache miss unless the cache is refreshed against the post-seed schema state.

**Scope.** `rosetta-extensions/stack-snapshot`. The capture side, not the replay side.
Fix is a `stacksnap capture --surface taxonomy` run against a fully brought-up stack,
committed to the capture manifest. Alternatively, decouple the fingerprint from
post-seed FK indexes so the pre- and post-seed schema states hash the same for the
taxonomy surface.

**Impact on `/demo-update`.** Phase 6 (snapshot replay) will always fail with `rc=5` on a
successful demo-N until the capture is refreshed. M-B currently treats this as a hard
failure. Once the capture is refreshed, Phase 6 will go GREEN with no code change on the
`/demo-update` side.

**Not user-visible.** This is a tooling defect, not a demo-content defect. Customer demos
do not see this.

## Finding B — `stories.seed.yaml` gen-batch mints fabricated `K-*` verified-skill node-ids

**Symptom.** `datadna measure-closure --stack demo-1` (Phase 7b of M-B) returned `rc=1`:
600 seeded verified-skill node-ids absent from the replayed taxonomy (sample:
`"K-15FIVE-6A8E"`).

**Root cause.** The `stories.seed.yaml` preset's gen-batch expansion mints
fabricated `K-<slug>-<hex>` node-ids for verified skills instead of resolving each name
through the `TaxonomyRefs` resolver against the replayed public taxonomy. The `K-` prefix
is diagnostic — real taxonomy node-ids never carry that shape.

**Scope.** `rosetta-extensions/stack-seeding` — the `stories.seed.yaml` preset and the
`GeneratedBatchSeeder` / gen-batch mother-prompt path that expands it. The
CODE-owns-structure / AI-owns-content boundary from
[`ai-generation-spec.md`](../../../corpus/ops/demo/ai-generation-spec.md) says non-resolving
names must **drop**, not be minted. That contract is being bypassed.

**Not caused by M-B.** These fabricated refs already exist on demo-1 *before* `/demo-update`
runs. The M-B additive seed contributes zero net-new fabricated refs; it added 539 users
using existing patterns and the closure gate re-audit found the same pre-existing 600.

### User-visible impact — YES, live customer demos are affected

Measured on demo-1 (2026-07-10) directly against the `postgres` DB:

| Table              | Rows with `skill_id LIKE 'K-%'` | Distinct fabricated IDs | Notes |
|--------------------|--------------------------------:|------------------------:|-------|
| `user_skills`      | **3 460**                       | —                       | of these, **168 rows carry `is_verified = true`** — the Skill Spotlight surface Stefano's Stories & Heroes narrative rests on |
| `membership_skills`| **3 522**                       | **113**                 | rendered on org-workforce and member listings |

Sampled fabricated IDs carry plausible-looking `skill_name` + `category_name`:

- `K-ABITOI-3F06` "Ability to Implement Machine Learning Algorithms" (Retail) — **262 members**
- `K-ABSTHI-9F5D` "Abstract Thinking" (Education and Training) — 14 members
- `K-ACCPLA-8708` "Account Planning and Management" (Sales) — 12 members
- `K-ACCCON-6555` "Access Control Rules Definition" (ICT) — 11 members
- `K-3DCOMG-7743` "3D Computer Graphics" (Design) — 11 members

**What this means in a customer demo.** Any surface that renders `user_skills` or
`membership_skills` — the hero profile's Skill Spotlight, the org-workforce
mapped/verified rollup, the Skills tab on `/enterprise/members` — is showing
**fabricated verified-skill labels** that do not exist in the real public taxonomy.
A customer clicking through to the taxonomy view (or comparing against a
production tenant) will hit dangling references.

The 168 `is_verified = true` rows are the most consequential: they populate the
verified-skill chain that the M-A/M-B live-run gate exists to protect (see
[`stories-spec.md`](../../../corpus/ops/demo/stories-spec.md) §G14). Every one of those
168 verified-skill records is a claim we cannot substantiate against the real taxonomy.

**Recommendation.** Prioritise the `stories.seed.yaml` / gen-batch fix independently
of `/demo-update`. Track as a separate stack-seeding item so demo-N substrate becomes
closure-clean regardless of whether `/demo-update` ever runs.

## Finding C — demo-1 has an `Exited(1)` directus container that never came up

**Symptom.** M-C's T2 (`stack-verify/live/verify.sh`) fires RED on demo-1 with
`postgres-schemas fail: missing schemas: directus`. `probe_postgres_schemas` (in
`stack-verify/lib/readiness.sh`) conditionally expects the `directus` schema when
the `directus` container is inspectable — and on demo-1, `docker inspect
demo-1-directus-1` succeeds because the container **exists**, but it's in
`Exited (1)` since the original `/demo-up 1` bring-up. The directus schema was
therefore never created, and every verify pass will surface this until the
crashed container is either removed or restarted-and-migrated.

**Root cause.** The demo-1 stack was brought up with `--local-content` (Directus
requested) but the Directus container failed its first-start and stayed exited.
The bring-up either did not surface the crash fatally, or a subsequent event
knocked directus over. On this substrate T2 will always be RED until directus is
either (a) restarted + its migrations re-run, or (b) removed from the demo-1
compose project so `probe_postgres_schemas` no longer expects the schema.

**Scope.** Two candidate owners:
- `rosetta-extensions/demo-stack` (bring-up side) — `/demo-up` should either
  fatally guard on the first-start of `directus` when `--local-content`, or expose
  a repair path (e.g. `--repair-directus`) so an operator can bring the crashed
  container back cleanly without a teardown.
- Operator-side box remediation — for demo-1 specifically, `docker rm
  demo-1-directus-1 && /demo-up 1` or a targeted `docker compose -p demo-1 up -d
  directus` + a directus migration re-run.

**Not user-visible in the *content* sense.** demo-1 without Directus falls back to
prod-read content (the documented `--local-content`-absent path), and customer
demos still render. Directus-authored CONTENT is served from prod. This finding
is a **tooling / verify-scope** RED, not a content RED.

**Impact on `/demo-update`.** M-C's T2 will always fail on demo-1 until the
crashed directus container is resolved. This is the SAME pattern as Findings A
and B: the gate fires on real pre-existing substrate debt, not on anything M-C
introduced. Once the container is either restarted+migrated or removed, T2 will
return GREEN with no `/demo-update` change.

**M-C hardening that DID land in-lane (commit `6bc207f`).** Two live-run refinements
required by the M-C gate contract itself (not workarounds for A/B/C):

1. **T1 readiness wait** — Phase 5 `migrate-demo.sh` restarts sentinel + backend
   at its tail; T1's `/api/health` probe fired inside the readiness race. Added a
   bounded 20-try × 1s retry so T1's fatal semantic tolerates the migrate-restart
   window without loosening it (probe still exits 1 after 20s of failure).

2. **T2 STACK_SERVICES scope pass** — `stack-verify/live/verify.sh` already had
   an M18 scope filter (skip probes for services not in `STACK_SERVICES`), but
   T2 wasn't feeding it. T2 now enumerates the running compose services via
   `docker ps --filter label=com.docker.compose.project=demo-N` and passes them
   through so a stack brought up without `--local-content` skips directus probes
   cleanly. (This is what surfaced Finding C — with the scope filter honoured,
   the remaining T2 RED on demo-1 is *genuine* stale-substrate signal.)

## Cross-cutting note

All three findings validate the `/demo-update` verify gate design (spec P1 §3.7 +
§2.3 T1–T3). None is in the `/demo-update` lane; all are recorded here so:

1. The `stack-snapshot` maintainer can schedule the capture refresh for Finding A.
2. The `stack-seeding` maintainer can prioritise the `stories.seed.yaml` gen-batch
   TaxonomyRefs-resolution fix for Finding B (with the user-visible impact above as
   the business case).
3. The `demo-stack` maintainer (or a demo-1 box operator) can resolve Finding C
   with a targeted directus restart-and-migrate or a controlled removal.
4. When all three are resolved, the M-C verify gate on `/demo-update` will return
   GREEN on demo-N substrates end-to-end — no `/demo-update` change required.
