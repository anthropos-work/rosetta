# M7a — Retro

**Summary:** Built the **stack-seeding framework** — a self-contained Go module that backfills a demo/dev stack
by talking **directly to its Postgres** (offset port, `COPY`), gated by a **3-layer production-isolation guard**
— and **proved it live**: a fresh injected `demo-1` stack, seeded with one `stackseed` run, lets the real
`user_clerkenstein` identity log in and an authorized GraphQL query returns **HTTP 200** (`membershipsCount:
1001`). 68 tests, all gates green. The build draft was authored by a sub-agent against a tight spec; I verified
the load-bearing guard + seeders myself and drove the live proof.

## What went well
- **The spike paid for itself immediately (M7a-D3).** Confirming early that `app/internal/bootstrap` is an
  `internal/` package (unimportable) killed the research's "link the ent client" perf path before any code, and
  pointed at direct-Postgres `COPY` — which is also *faster* and needs no private deps. The whole module builds
  offline with just pgx + yaml.
- **The full live proof was worth the cost.** It caught **two real bugs that every unit test missed**: the
  casbin `g2` arg-order (the model matches `g2(org, sub, role)` — the row is `(org, user, role)`, not
  `(user, org, role)`) and the **missing global Sentinel policy** on demo stacks (`migrate-demo.sh` never ran
  `init_policy.sql`). Both now fixed + the g2 order pinned by a regression test.
- **The isolation guard is genuinely airtight** — 3 layers (CheckWrite / PreflightEnv / AssertClean), 97%
  covered, and the `AssertClean` "no shared/external writes landed" assertion ran clean on the live seed.
- **Delegation + verification split worked.** The sub-agent produced a coherent 19-file module; my job was the
  safety-critical review + the live proof, which is where the real bugs surfaced.

## What didn't / constraints
- **The build sub-agent hit a weekly limit** on its final report step (after writing all files); I picked up its
  output (which had landed) and finished the 4 cleanup issues (2 parse errors, a mutex-copy, the StackN logic)
  myself. No work lost.
- **The 403→200 proof was harder than expected** — the platform's authz is *feature-based* (`org:feature:X`),
  not just role-based, and two demo-build quirks got in the way: `/api/workforce/*` panics on a nil `*sql.DB`
  (a demo-build defect, not seeding), and the org queries needed the global policy + correct g2. Switching the
  proof vehicle to a GraphQL org query (the properly-wired ent path) gave the clean 200.
- **`pg`/`cmd` DB-touching code is live-proven, not unit-tested** (61.9% / 47.4%) — those paths need a real DB;
  the live login→200 is their coverage.

## Findings about the platform (not our bugs — noted)
- `/api/workforce/*` handlers panic on a nil database handle in the injected demo build (the workforce Manager's
  `*sql.DB` is unwired). A demo-build/platform concern, out of M7a scope — recorded so a future stack-build pass
  can investigate.

## Carried forward → M7b / M7c
- **M7b (data-DNA):** the direct-`COPY` perf path creates a hand-written-SQL drift risk; M7b's schema
  introspection is the gate. The reference seeders' column lists (`public.users/organizations/memberships`,
  `sentinel.casbin_rules`) are the first genes.
- **M7c (fleet):** Directus content via snapshot-replay (the guard blocks live Directus writes); the full seeder
  fleet (taxonomy, content, sessions, backdated activity, tier/feature grants beyond the minimum identity);
  feature-credit/tier grants so *every* feature-gated route (not just org-membership queries) authorizes.

## Metrics
See [metrics.json](metrics.json). 68 tests (blueprint 8 · isolation 18 · seeder 6 · dag 10 · pg 9 · seeders 11 ·
cmd 6); build/vet/`-race`/gofmt clean; shellcheck clean. Live proof: demo-1 (14 containers) → seed → login **200**,
isolation audit clean. 2 bugs caught+fixed. Monorepo now 6 sections (+stack-seeding).
