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

## Cross-cutting note

Both findings validate the `/demo-update` closure gate design (spec P1 §3.7). Neither is
in the `/demo-update` lane; both are recorded here so:

1. The `stack-snapshot` maintainer can schedule the capture refresh for Finding A.
2. The `stack-seeding` maintainer can prioritise the `stories.seed.yaml` gen-batch
   TaxonomyRefs-resolution fix for Finding B (with the user-visible impact above as
   the business case).
3. When both are fixed, the M-B closure gate on `/demo-update` will start returning
   GREEN on demo-N substrates end-to-end — no `/demo-update` change required.
