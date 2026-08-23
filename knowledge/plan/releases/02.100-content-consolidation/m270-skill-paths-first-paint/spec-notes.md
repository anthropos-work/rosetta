# M270 — Spec notes

_Stub. Headers derived from the scope; filled during the milestone._

## Pre-flight: the citations, re-resolved on disk

_Measured 2026-08-23 at `stack-demo/next-web-app` @ `20a410d7d` (read-only, at scaffold time), so the
milestone starts from a known ref rather than from a remembered one. Three divergences from the brief's
wording are recorded here and NOT propagated into `overview.md`, whose citations are carried verbatim as
the milestone contract:_

1. **The hook file is `useGetClerkOrganization.tsx`, not `.ts`** at this ref. The contract's claim about
   what it returns is unaffected; the extension is not.
2. **`LibrarySkillPathsContainer.tsx:664-666` has a THIRD disjunct** the contract's wording does not
   mention — the condition reads
   `{(!loadingPrivateSkillPaths && filteredSkillPaths.private.length > 0) || showAcademyInOrg ? (`.
   `showAcademyInOrg` is computed at `:281`. **This matters for the fix**: an org with academy cards
   renders the section anyway, so the empty-flash may not reproduce on every seeded org. Establish which
   demo orgs have `showAcademyInOrg` true before concluding a fix worked.
3. **`useGetPrivateSkillPaths.tsx:58`** is the `enabled: Boolean(organizationId && enable)` line — the
   contract gives the file but no line.

_Everything else resolved exactly: the page is `'use client'` (`:1`) and destructures
`{ organizationId, organization }` at `:17` with no `isLoadingOrg`; `GET_PUBLIC_SKILL_PATHS` is declared at
`packages/graphql/src/query/skill-path.ts:634`; `useGetPublicSkillPaths.tsx` passes `sortField` (`:25`) and
`offset: 0` (`:27`) and no limit; `apps/integration/.../page.tsx` is a server component calling
`getPublicSkillPathsServerOrThrow` (`:2`, `:24`) and passing `initialSkillPaths` (`:39`)._

_The `patches/` directory in the authoring copy holds **26** manifests today; `demopatch-spec.md` §5 last
reconciled at **23** (v2.7 M253). Reconcile against the directory, not the table, when adding one._

## (a) The empty flash — a disabled query reads as "not loading"

_Not yet written._

### The Clerk-hydration window

### TanStack v5: disabled ⇒ `isLoading === false`

### The unrendered section at `LibrarySkillPathsContainer.tsx:664-666`

### What the affordance should be

## (b) The slowness — un-paginated, deeply nested public query

_Not yet written._

### The selection set and what it costs

### The client-side reduce that only feeds a badge count

### The in-repo SSR precedent (`apps/integration`) and whether `apps/web` can take it

## Vehicle: demopatch vs escalation

_Not yet written. Decided in `decisions.md` D-1._

### Anchorability of each half

### Manifest shape (10-key schema, `scope: demo`)

### Maintenance cost: sha-pinning drift, §5 inventory, `TestPatchInventory`

## Measurement

_Not yet written._

### The leg being graded, and whether it exists in `latency-budget.md`

### Before/after, with the environment stated

### Local vs `--public-host` reproduction
