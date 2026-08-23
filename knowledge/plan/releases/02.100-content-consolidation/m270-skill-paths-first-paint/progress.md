# M270 — Progress

**Status: PLANNED. Not started.**

One section per `In:` item in [`overview.md`](overview.md). A box is ticked when the thing is *measured
done*, not when it is *believed done*.

## (a) The empty flash — diagnose and confirm the chain

- [ ] Reproduce on a live demo stack — first paint of `/library/skill-paths`, cold, as a seeded hero
- [ ] Confirm `page.tsx` does not destructure or pass `isLoadingOrg`
- [ ] Confirm the hydration window leaves `organizationId` undefined → `enabled` FALSE
- [ ] Confirm TanStack v5 reports `isLoading === false` with `data === undefined` for the disabled query
- [ ] Confirm `LibrarySkillPathsContainer.tsx:664-666` renders neither the section nor a spinner
- [ ] Establish which seeded demo orgs have `showAcademyInOrg` true (it changes what reproduces — see
      `spec-notes.md`)

## (b) The slowness — diagnose and time it

- [ ] Confirm `GET_PUBLIC_SKILL_PATHS` is issued with no limit
- [ ] Size the response and the selection-set depth on a real demo catalog
- [ ] Time the leg **before** any (a) fix lands (once the spinner exists, the empty-page evidence is gone)
- [ ] Confirm the client-side reduce and what it actually feeds (the `contentType` badge count)
- [ ] Confirm `apps/integration` takes the SSR-prefetch path and `apps/web` does not
- [ ] Establish whether `LibrarySkillPathsContainer`'s `initialSkillPaths` prop is usable from `apps/web`
      unchanged

## Vehicle decision (D-1)

- [ ] Assess anchorability of the (a) fix
- [ ] Assess anchorability of the (b) fix (page-shape change — client → server component)
- [ ] **Decide and record D-1 in `decisions.md`** — demopatch, escalate, or split the two halves
- [ ] If escalating: raise with the platform team and record the escalation; re-scope this milestone to
      diagnosis + the loading-affordance patch only

## Ship the loading affordance

- [ ] Choose the affordance (reuse the container's existing loading pattern rather than invent one)
- [ ] Author the manifest (10-key schema, `scope: demo`, anchor occurring exactly once)
- [ ] All 7 guards pass (G1–G7)
- [ ] Verify **in a browser**, not by curl or bundle-grep
- [ ] Verify idempotent re-apply across a second `/demo-up` (G4)
- [ ] Verify revert leaves the clone git-clean (G5)

## Demopatch maintenance cost — recorded, not just incurred

- [ ] `demopatch-spec.md` §5 patch inventory updated
- [ ] `TestPatchInventory` constants updated (total + per-repo breakdown) and GREEN
- [ ] Reconciled against the `patches/` directory, not against the table
- [ ] Drift cost written down: what moves this manifest, and how it will be noticed

## Grade it

- [ ] Name the leg being measured; state whether it exists in `latency-budget.md` or is out-of-budget
- [ ] Before/after at that leg, same stack, n stated
- [ ] Environment stated with every number
- [ ] Reproduced (or explicitly not reproduced) on the `--public-host` path B3 was reported from

## Open questions closed

- [ ] Every Open question in `overview.md` is answered or explicitly carried forward
