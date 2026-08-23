---
milestone: M270
title: "Skill-paths first paint"
milestone_shape: section
status: planned
release: "02.100-content-consolidation"
depends_on: "none"
parallel_with: "M266, M267, M268"
complexity: medium
last_updated: "2026-08-23"
---

# M270: Skill-paths first paint

**Goal:** `/library/skill-paths` shows something honest on first paint — a loading affordance, then
content — instead of rendering empty and popping in.

Serves annotation request **B3**: *"the very first time i load or click to reach skill path … the library
appears empty/not rendered .. then it does.. not sure if it is the first load has an issue or is very slow
and no animation tells me what is happening. can you both speed up the loading and fix the first load
issues (depending on which of the 2 is the issue)?"*

**BOTH of the reviewer's hypotheses are TRUE and they are SEPARATE defects.** The milestone does not have to
choose between them; it has to keep them apart, because they have different fixes and different vehicles.

## Scope

**In:**

  - **(a) THE EMPTY FLASH — a disabled query reads as "not loading".** The chain, measured end to end:
      - `apps/web/src/app/(authenticated)/(verified)/library/skill-paths/page.tsx` is `'use client'` and
        pulls `organizationId` from `useGetClerkOrganization()` — but does **NOT** destructure or pass
        `isLoadingOrg`, which that hook explicitly returns
        (`apps/web/src/hooks/useGetClerkOrganization.ts`:
        `isLoadingOrg: !(signInLoaded && userLoaded && orgLoaded)`).
      - While clerk-js hydrates, `organizationId` is `undefined`, so in
        `packages/graphql/src/hooks/skillpath/useGetPrivateSkillPaths.tsx` the TanStack query is
        `enabled: Boolean(organizationId && enable)` → **FALSE**. In TanStack Query v5 a DISABLED query is
        `status:'pending'`, `fetchStatus:'idle'` → **`isLoading === FALSE` with `data === undefined`**.
      - `packages/ui/src/Library/LibrarySkillPathsContainer.tsx:664-666` then renders the "For {org}"
        section only when `(!loadingPrivateSkillPaths && filteredSkillPaths.private.length > 0)` — so the
        section is **not rendered AND NO SPINNER TAKES ITS PLACE**. When Clerk finishes, the query enables
        and it pops in. That is exactly *"appears empty, then it does"*.

  - **(b) THE SLOWNESS — the public query is un-paginated and deeply nested.**
      - `useGetPublicSkillPaths` requests `GET_PUBLIC_SKILL_PATHS` with
        `{ sortField: DateCreated, asc: false, offset: 0 }` and **NO LIMIT**; the selection set
        (`packages/graphql/src/query/skill-path.ts:634-700+`) pulls every path with
        `libraryCategories{...macroCategory}`, `skills`, `curation`, and
        `chapters{ jobSimulations{id}, steps{ resource{ ...on YoutubeVideo/WebResource/UdemyCourse/Podcast } } }`.
        The container then reduces all of it **client-side just to compute a `contentType` badge count**.
      - **IN-REPO PRECEDENT FOR THE FIX:** `apps/integration/src/app/library/skill-paths/page.tsx` is
        **ALREADY** a server component calling `getPublicSkillPathsServerOrThrow` from
        `@anthropos/graphql/src/server/skillpath` and passing `initialSkillPaths` — a prop
        `LibrarySkillPathsContainer` already accepts and `apps/web` never uses. **The SSR-prefetch path
        exists; `apps/web` simply does not take it.**

  - **DECIDE THE VEHICLE EXPLICITLY (D-1).** ⚠️ **SHAPE FINDING: every line above is in `next-web-app`, a
    platform repo.** Under the release-wide zero-platform-edit rule the fix must ship as a **sha-pinned
    demopatch** ([`demopatch-spec.md`](../../../../../corpus/ops/demo/demopatch-spec.md)) — or, if the change
    is big enough that a patch is the wrong vehicle, be **raised with the platform team** and this milestone
    **re-scoped to the diagnosis + the patch for the loading affordance only**. This is a decision to be
    made and written down in `decisions.md`; **do not let the vehicle be chosen by accident.**

  - **Record the demopatch maintenance cost, in the milestone, at authoring time.** A demopatch manifest is
    **sha-pinned** and WILL drift when the platform moves. The freshness model is *the anchor is the
    contract; the whole-file sha is only a baseline* (`demopatch-spec.md` §6, G2), so a moved file
    self-heals but a **moved or duplicated anchor REFUSES** — zero occurrences and two-or-more occurrences
    both refuse. Adding a manifest also means updating **§5 the patch inventory** and the
    `TestPatchInventory` fence, which pins **the exact inventory total + the per-repo breakdown**
    (`demopatch-spec.md:98`, `:294`, `:616`). That fence has shipped RED before, at v2.7 M253, because the
    table was updated and the constants were not.

  - **Grade the "speed up" half at the leg, not at the suite.** Any (b) change carries a DIRECT
    before/after measurement of the leg it targets, on the same stack, with **the environment stated
    alongside every number** — the standing release rule, and
    [`latency-budget.md`](../../../../../corpus/ops/demo/latency-budget.md) is the definition of what "fast"
    means and how it is graded. First paint is not currently a metered leg in that budget; see Open
    questions.

**Out:**

  - the skill-path builder
  - any change to the skill-path data model

## Depends on

none

## Parallel with

M266, M267, M268 — **a different repo entirely** (`next-web-app` vs the cockpit / entitlement / seeding
surfaces), so there is no shared-file contention to sequence around.

## Open questions

  - **Is (b) inside a demopatch's reach at all?** (a) is a small, anchorable edit — destructure
    `isLoadingOrg`, thread it, render an affordance. (b) as described (adopt the SSR-prefetch path in
    `apps/web`) is a **page-shape change** — client component → server component — and that is a
    different size of diff to hold in a sha-pinned anchor across platform moves. **This is D-1 and it is
    genuinely open; the two halves may take two different verdicts.**
  - **What does "honest" mean for the affordance?** A skeleton in the section's place, a spinner, or the
    section rendered with a loading state — the container already computes
    `loadingPublicSkillPaths || loadingPrivateSkillPaths` in at least three other places
    (`LibrarySkillPathsContainer.tsx:523`, `:545`, `:766`), so a house pattern probably already exists and
    should be reused rather than invented. **Not yet chosen.**
  - **Is first paint a gradeable leg?** `latency-budget.md` defines **ACCESS** (authenticated shell
    rendered + interactive) and gates the click→login path at p95 < 5 s. A *within-app navigation to a
    library route* is not obviously one of its legs. Either this milestone reuses an existing leg, adds
    one, or measures out-of-budget and says so — **undecided, and it changes what "graded" means here.**
  - **Does fixing (a) mask (b)?** Once a spinner occupies the space, the page stops *looking* empty while
    still being slow. The diagnosis must establish the (b) timing **before** (a) lands, or the evidence for
    (b) is gone.
  - **How much of the reviewer's report is the `--public-host` tailnet path?** B3 was observed on
    `https://demo1.anthropos.work:13000`. `latency-budget.md` records that the same defect cost ~6 s on a
    laptop and ~112 s on the tailnet VM. **A local-only measurement may not reproduce what was reported.**

## KB dependencies

- [`demopatch-spec.md`](../../../../../corpus/ops/demo/demopatch-spec.md) — the patch vehicle + its 7 guards
  (G1 path-assert · G2 anchor gate · G3 never-commit · G4 idempotent · G5 journalled self-revert · G6
  demo-only · G7 apply post-condition), the 10-key manifest schema, the three apply vehicles, and §7
  *Adding a new patch*
- [`next-web-app.md`](../../../../../corpus/services/next-web-app.md) — the monorepo the whole defect lives in
- [`latency-budget.md`](../../../../../corpus/ops/demo/latency-budget.md) — what "fast" means and how it is
  graded (ACCESS, the p95 < 5 s gate, the per-leg attribution model, *state the environment with every
  number*)

## Delivers →

[`corpus/ops/demo/demopatch-spec.md`](../../../../../corpus/ops/demo/demopatch-spec.md) — a new manifest joins
the inventory. **Conditional on D-1:** if the vehicle decision routes (b) (or both halves) to the platform
team instead, this milestone delivers a **diagnosis + escalation record** and only the loading-affordance
manifest, and the spec's §5 table + `TestPatchInventory` constants move by that smaller amount.

## Constraints (release-wide, non-negotiable)

- **Zero platform-repo edits.** "Update the platform repos" in this release means **pull them fresh**. A need
  that can only be met by a platform edit **escalates**; it does not edit. For this milestone that constraint
  is the whole shape of the work — **every line of the defect is in `next-web-app`.**
- All stack tooling lives in `rosetta-extensions`, built + tested in the authoring copy, tagged, **pushed**,
  then consumed per-stack at a pinned tag.
- Secrets handled values-blind.
