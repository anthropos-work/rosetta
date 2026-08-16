# M259 — Progress

**Status: COMPLETE — VERDICT `GO`.** Closed 2026-08-14.

- [x] pull EVERY platform repo fresh into BOTH clone sets — pull, never commit
- [x] read the canon, the redirect map, and the retired-id guard in `app`
- [x] measure the real new counts
- [x] establish whether the redirect map is TOTAL or PARTIAL, and what a dropped id should become
- [x] reconcile the three lineages
- [x] **Delivers →** `corpus/architecture/taxonomy-canon.md` (+ indexed in `corpus/architecture/README.md`)

## What was done

**Clone sets refreshed, both of them, fast-forward only.** `stack-dev/app` +175 → `4bccda085`,
`stack-dev/next-web-app` +86 → `20a410d7d`, `stack-dev/platform` +4, `stack-demo/app` +99,
`stack-demo/next-web-app` +27, `stack-demo/jobsimulation` +5. `stack-demo/rosetta-extensions` deliberately
NOT pulled — it is a pinned-tag consumption clone, not a main-tracking repo. Four clones showed dirty; all
four were untracked artifacts (`studio/` hand-clones, the git-ignored `.agentspace/rext.tag`), so nothing
was stashed and nothing was discarded.

**The canon is a checked-in artifact inside `app`** — `app/taxonomy-canon/` — so every figure below was
measured from files in the tree. **No production read was required**, which also means the barrier cost no
prod access and carries no tenant-data exposure.

## The measurements

| Quantity | Count |
|---|---:|
| canon skills | **3,562** |
| canonical roles | **706** |
| skills with a redirect | 12,835 |
| **skills dropped, no successor** | **26,518** |
| roles with a redirect | 11,182 |
| **roles dropped, no successor** | **10,689** |

Pre-consolidation, per the platform's own `redirect.go`: **43,584 skills / 22,511 job roles**.

## The four findings that resize downstream work

1. **The redirect map covers ~33 % of retirements.** 12,835 of 39,353 retired skills have a successor;
   **26,518 have none.** A remap-everything design for M262 is not viable — it needs a resolve-or-drop path
   for the other two thirds. The map itself is clean: 12,835 distinct old ids, no duplicates, no empty
   destinations, no chains.
2. **Five net-new tables sit OUTSIDE the capture surface** — `skill_redirect`, `job_role_redirect`,
   `category_translation`, `specialization_translation`, `taxonomy_canon_state`. Without them a replayed
   demo cannot resolve a retired id at all, loses two levels of the EN/IT axis, and has nothing behind the
   `/taxonomy` canon-state panel. **Net-new scope for M260/M261.**
3. **`taxonomyguard` closes the taxonomy to runtime minting.** Get-or-create no longer creates. Any tooling
   that fed a name and relied on minting now silently resolves to nothing. **Resolve or accept the miss.**
4. **`taxonomyredirect` is INERT by design** — the tables ship and populate while every read path behaves as
   before. So the mapping exists as DATA before it exists as BEHAVIOUR, which is exactly what lets this
   project consume it ahead of the platform.

## A correction this milestone made about its own work

An early **head-sample** of `skill_redirects.csv` read as semantically warped (generic skills mapping into
agriculture) and was nearly written up as a redirect-quality risk. The CSV is **ordered by target category**
and opens with Agriculture. Re-sampled **randomly**, both `review=true` and `review=false` classes read
sensibly, and the target-category spread is ordinary (IT 2,445 · Engineering 1,879 · …). **Redirect quality
is not a risk; coverage is.** The retraction is recorded in the delivered doc rather than quietly dropped,
because the biased-sample error is the reusable lesson.
