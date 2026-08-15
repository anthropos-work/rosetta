# M263 — Progress

**Status: in progress** (2026-08-15). The taxonomy page needs a stack built on the NEW pins, and
getting one exposed a chain of platform-realignment defects that had nothing to do with taxonomy.

## Why a new stack at all

`/taxonomy` is a server-rendered route that calls `taxonomyCategories` over GraphQL. Three things must
line up: the **app** must serve the query (it does — `internal/web/backend/graphql/graph/schemas/
taxonomy.graphqls`, net-new), **next-web** must carry the routes (v2.144.x), and the **database** must
hold the canon. `demo-4` has none of the three: its app pin predates the v2 migrations, and it holds
42,790 old skills. It is also flagged do-not-reset, so a new stack (`demo-5`) it is.

## Three defects found bringing it up — all "the platform moved, the tooling didn't"

### 1. Clerkenstein's disarm had nothing to disarm

`app/go.mod` at `4bccda085` has **ZERO `anthropos-work/` requires**. Every first-party module was
folded in, **colony included**. `apply-authn`'s entire mechanism — clone colony at the pinned version,
swap its clerk provider for the disarmed twin, drop it in as `vendor-colony/`, add a `go.mod` replace
— has nothing to operate on. It failed with *"couldn't find the colony version in go.mod"*, and the
bring-up correctly called that **FATAL**: without the disarm the image builds against real Clerk and
**every demo login 401s**.

The authn code is now at `app/internal/authn`, provider at `internal/authn/provider/clerk/` —
mirroring the module layout. So the disarm becomes a **file swap in the ephemeral build-scratch
clone**, which is *simpler* than what it replaces: no private-repo clone (no `GH_PAT`), no `go.mod`
edit, no vendor dir. Detection is on the **absence** of the colony require, because that is exactly
the condition making the module path impossible.

### 2. The Dockerfile still asked for the artefact of the path not taken

`COPY vendor-colony` was injected unconditionally, so the build died with `"/vendor-colony": not
found` — *after* the disarm had already succeeded. Now gated on the **directory the applier actually
produced**, the one signal that cannot disagree with what is on disk.

### 3. The twin replaces a PACKAGE, not a file

In the module era the twin was dropped over colony's `authn/provider/clerk` wholesale, so declaring
`Clerk`, `User` and `Organization` in one file was right. In-tree that package is **split** across
`clerk.go` + `clerk_org.go` + `clerk_user.go`, so writing the twin over `clerk.go` alone left the
siblings:

```
internal/authn/provider/clerk/clerk_org.go:5:6:  Organization redeclared in this block
internal/authn/provider/clerk/clerk_user.go:41:6: User redeclared in this block
```

The package is now cleared before the twin is written, with a post-condition asserting exactly one
`.go` file remains — because a leftover sibling IS the bug, and it surfaces 17 seconds into a Docker
build rather than at the point of the mistake.

Shipped as `v2.9.2-rext` and `v2.9.3-rext`.

## The page is LIVE on our stack, and the full walk runs

Verified in a browser against `demo-5` (presenter world, `dan-manager`):

- **`/taxonomy` renders** — title *"Anthropos | Skills taxonomy"*.
- **The nav entry is there**, between *AI Academy* and *Organization*, exactly where `taxonomyMenuItem`
  sits in the Library section. It is **not** flag-gated (`restricted: ''`), so no PostHog bootstrap was
  needed — unlike the assign-content item beside it.
- **HOP 1 works**: clicking a category lands on `/taxonomy/category/agriculture` with the heading
  *"Agriculture"* — a real canon route with real canon content.

**⚠️ The section that stood here — filing the empty specialization hop as a platform question —
is RETRACTED. It was our bug.** Kept in outline because the misdiagnosis is the lesson.

### What I got wrong, and why it looked convincing

Every category page rendered `SPECIALIZATIONS | 0`, and the header said `CATEGORIES 32` while the
database held 25. I read the mismatch as the surface enumerating the wrong table and filed it
platform-side. Reading the resolver settles it: `taxonomyapi.Manager.Categories` deliberately appends
`categorieDiSoliRuoli` — *"the domains that exist in the role tree and not in the skill one: food
service, personal care, research, manufacturing, the armed forces"*. **32 is correct product
behaviour**, and role-only categories correctly have no specializations.

The real cause was one query away and I did not run it: **not one row in the replayed taxonomy had a
slug.** The resolver filters `category.SlugNEQ("")`, so all 25 real categories were excluded and the
only ones left to render were the role-only ones — which is exactly the symptom. My earlier
"specializations with a slug: 283" was itself wrong: it counted `slug IS NOT NULL`, and `''` is not
NULL.

### The actual defect: the capture was dropping 26 columns

The surface's column lists came from the PRE-taxonomy-v2 schema. M261 added four new TABLES and never
re-derived the existing tables' COLUMNS. Diffed against the live schema:

| table | dropped |
|---|---|
| `categories` | page, provenance, credibility, **slug** |
| `specializations` | page, provenance, credibility, **slug** |
| `skills` | kind, implements_node_id, page, provenance, credibility, **slug** |
| `job_roles` | family, esco_uri, esco_code, isco, onet_code, model_fill, page, provenance, credibility, **slug**, esco_title, source_job_role_uuid |
| `skill_translations` · `job_role_translations` | page |

That is the whole taxonomy-v2 "governed entity" model plus the canon's structural fields. **A row
count could never have caught it** — 3,562 skills replayed either way. It is the hand-maintained-list
rot M260 fenced against, walked into one milestone later.

### Two ordering defects fell out, both only reachable once the parents moved to DELETE

- **DELETEs must run child-first.** A surface declares parent-first because that is LOAD order;
  clearing is the mirror image. Deleting `categories` while `specializations` held rows violated the
  FK outright.
- **TRUNCATE must run BEFORE the deletes.** The truncatable tables are the leaves and they reference
  the DELETE-cleared parents. The old delete-first order was written for the directus case, where
  nothing inside the surface pointed at the DELETE table — free, not correct.

### Proven live

25 / 283 / 3,562 slugs and 3,562 pages replayed into `demo-5`, and the walk runs end to end:

```
/taxonomy/category/ai-skills  →  15 specializations
/taxonomy/specialization/ai-augmented-work  →  7 skills
/taxonomy/skill/ai-adoption-change-management   ("AI Adoption & Change Management")
```

**M263's gate — navigable and real — is met.** Shipped `v2.9.6-rext`.
