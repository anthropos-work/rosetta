---
iter: 15
milestone: M211
iteration_type: tik
status: planned
created: 2026-07-08
---

# iter-15 — tik: M42m manager coverage GREEN (demopatch hash re-pin + next-web rebuild)

## Active strategy reference
**TOK-01** move (4) "iterate warm-first, prove cold" — close the **manager half** of sub-condition (e)
(M42 coverage) on the live demo-1 substrate.

## Step 0 — Re-survey
M42e employee coverage is GREEN (iter-14). Manager coverage (M42m, `dan-manager @ Cervato Systems` →
`MANAGER_MANIFEST_BASE`, org-independent of the Northwind AI-readiness showcase) has **never run on the
merged demo-1**. Run-3 evidence: the two next-web URL demopatches (`next-web-studio-url`,
`next-web-public-website-url`) **G2-REFUSED** (sha drift: re-synced urls.ts ≠ pinned `pre_sha256`), so
demo-1's current next-web image lacks the STUDIO_URL + PUBLIC_WEBSITE_URL demo-local overrides → the manager
vantage will show ≥2 prod-eject escapes. `dan.rossi3@cervato-systems.com` (admin) confirmed seeded.

## Cluster / target identified
Manager coverage GREEN needs: (1) re-pin the 2 drifted demopatch hashes to the re-synced v2.1 next-web
source (pure hash re-pin — the STUDIO_URL + PUBLIC_WEBSITE_URL anchor hunks are byte-identical to v1.10;
only sibling exports drifted → file-level sha moved, the same re-anchor M49 #8 did), (2) rebuild the demo-1
next-web image WITH the 2 patches applied, (3) any other manager-vantage fixes the baseline sweep surfaces.
Baseline-first: run the manager sweep on current demo-1 to triage the FULL manager picture before fixing.

## Hypothesis
Re-pinning the 4 chained hashes (studio pre/post + public-website pre/post) to the pristine re-synced
urls.ts (`0d4c3790…`) lets both demopatches G2-PASS → next-web bakes the demo-local STUDIO_URL/
PUBLIC_WEBSITE_URL → the manager nav Studio link + the ai-simulations drill-down stay demo-local → escapes
→ 0. If the baseline sweep surfaces other manager blockers (empty sections / perf wall / persona), fix
those too under the same tik (the manager-gate is the single target).

## Expected lift
M42m manager coverage: **GATE MET** (escapes 0, failingSections 0, personaFailures 0, crossPortFollowFails
0). Closes the manager half of sub-condition (e).

## Phase plan
1. Baseline manager sweep (demo-1, dan-manager @ Cervato) → triage all failures.
2. Re-pin the 4 demopatch hashes in the rext authoring copy (`demo-stack/patches/next-web-{studio,public-website}-url/`);
   validate by simulating apply (G2 pass + post_sha256 match); commit + move tag `quick-change-m211`.
3. Rebuild demo-1 next-web with the 2 patches applied (reap-safe detached) + recreate the container.
4. Address any other baseline-surfaced manager blockers (route or fix in-tik).
5. Re-run the manager sweep → expect GATE MET.

## Escalation conditions
- If a URL host has NO per-URL override AND is not demopatch-able (platform-bound) → re-scope trigger
  (needs a platform-source edit). (Not expected — both patches already exist + are behavior-identical.)
- If the anchor hunk itself changed (not just sibling exports) → the patch needs a new anchor, not just a
  hash re-pin → still Fate-1 (author the new anchor) unless it needs a platform edit.
- If a manager blocker requires a platform-repo edit → `unimplementable-without-platform-edit` escalation.

## Acceptable close-no-lift outcomes
If the demopatch re-pin lands + rebuild completes but a distinct NEW manager blocker (e.g. a data gap like
sim-embeddings) needs a separate fix, close-fixed-partial with the URL escapes cleared + the residual routed.
