# M230 — Decisions

_(decisions recorded as they arise during build)_

## TOK-01: Option C — FS-as-published fallback via a sha-pinned demo-patch — 2026-07-19

**Tok type:** bootstrap (iter-01)

**Initial strategy:** Fill the empty demo ant-academy home grid with **Option C** — a sha-pinned rext
`demopatch` on the demo's **OWN ephemeral ant-academy clone** that **restores an FS-as-published fallback** in
the academy's server catalog resolver, so the grid renders the committed catalog through the **real resolver +
render chain** with **NO "Draft" chip**.

Concretely (verified against real code): `src/lib/serverTenant.js::getServerCatalogView()` (and its public twin
`getPublicCatalogView()`) is `const view = (await getBackendCatalogView(eids)) ?? emptyCatalogView(); return draftsEnabled() ? mergeDrafts(view, eids) : view`.
The M7 cutover deliberately removed a pre-existing FS-as-published fallback at exactly the `?? emptyCatalogView()`
expression. Option C's patch swaps that fallback for a **published (un-chipped) FS catalog view** — the same FS
tree `mergeDrafts` reads, but WITHOUT the `_draft:true` tag that produces the chip. Result: on a demo (empty
backend → null → fallback), REAL cards render through the unchanged RSC → `AcademyClient` → `SkillPathCard` chain,
production-faithful, no chip. Applied to the ephemeral clone before `next dev` launch and reverted on teardown
(the native-launch analog of the image-baked next-web patches), wired into `ant-academy.sh` — mirroring the
existing `patches/ant-academy-dev-origins` academy-patch precedent.

**Rationale (Option C over Option B):**
1. **Least infrastructure risk.** Option C needs NO prod DB read (removing the user-blocker risk the orchestrator
   flagged), NO new snapshot surface, NO academy-subgraph composition into the demo router, NO endpoint wiring.
   Option B needs all four, each a potential cold-up blocker. The gate must be PROVEN on a cold /demo-up here;
   Option C minimizes the surface that can fail.
2. **Zero platform edits (the release hard line).** Option C uses the sanctioned `demopatch` mechanism
   (`demopatch-spec.md`): patch the demo's own ephemeral clone, revert-clean; the canonical
   `anthropos-work/ant-academy` repo is never touched. Proven: `patches/ant-academy-dev-origins` already patches
   the ephemeral academy clone.
3. **The seam is verified + clean** (the `?? emptyCatalogView()` fallback — code-read this iter, not assumed).
4. **Gate-faithful.** The exit gate permits "the real DB-authoritative GraphQL path (OR a faithful equivalent)."
   Option C is the faithful equivalent: real resolver + render chain, sourced FS-as-published (behavior-identical
   to the pre-M7 fallback). Option A (the `ACADEMY_SHOW_DRAFTS` draft layer) is REJECTED because `mergeDrafts`
   stamps `_draft:true` → the visible chip the gate forbids.
5. **Reproducible on a cold /demo-up** — the academy runs natively (`next dev`, source read live); the patch
   applies at launch + reverts at teardown.

**Why NOT Option B (kept as the fallback):** higher fidelity (the true GraphQL path) but multiplies the
infrastructure that must ALL work on a cold up here — prod DB read (+ a public predicate for academy rows), a
net-new firewalled snapshot surface, confirmed academy-subgraph composition into the demo supergraph, and endpoint
wiring. Deferred unless Option C proves unfaithful/unpatchable (the seam is clean, so it does not).

**Strategy class:** new-direction (bootstrap tok — no prior strategy to compare against).

**Distance-to-gate context:** Gate metric = rendered-card count on the academy home grid, employee vantage, via
the coverage sweep `ANT_ACADEMY` descriptor — **≥ floor, NO Draft chip, 0 prod-ejects, on a cold /demo-up**.
Baseline = **0 real cards** (the F4 carry; `emptyCatalogView()`), a priori established + code+launcher-verified
this iter. Infra: a cold /demo-up is feasible here (demo-1 images built 41h ago).

**Next-tik direction (iter-02, first tik):** Author the Option C demo-patch in the rext authoring copy
(`.agentspace/rosetta-extensions/demo-stack/patches/ant-academy-fs-published-fallback/` + its `<name>.yaml`
manifest), content-anchored on the `?? emptyCatalogView()` fallback in `serverTenant.js` (BOTH
`getServerCatalogView` + `getPublicCatalogView`), producing a published (un-chipped) FS catalog view. Wire it
into `ant-academy.sh` (apply-before-launch + revert-on-teardown, mirroring `ant-academy-dev-origins`).
Unit-verify: the anchor matches the real source, and the produced view carries NO `_draft:true`. Tag the
authoring copy + re-pin the consumed tag. Then take the first proof: a cold /demo-up + the coverage sweep's
`ANT_ACADEMY` rendered-card count. If a cold /demo-up cannot complete in this environment, exit user-blocker with
the specific obstacle (do not fake the proof).
