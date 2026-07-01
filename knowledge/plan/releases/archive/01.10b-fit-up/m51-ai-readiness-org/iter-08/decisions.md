# iter-08 — decisions

## D1 — resolve the closed cycle id (DB ground truth) + confirm the FE reads `?cycle=` from the URL
**Context:** the run-5 deep-link strategy needs the actual closed cycle id + must confirm the FE actually
routes a URL `?cycle=` to the backend GET.
**Finding:**
- DB (demo-1 offset postgres): the Northwind closed cycle is `95d9fc3d-48c0-53ca-82ac-e10713100c97`, status
  `closed`, 199 frozen snapshots, 78.4% stage-3, Aria=stage3 / Ben=stage1 / Dana no snapshot. Matches the
  deterministic `deterministicUUID("ai-readiness-cycle:" + StoryOrgID("ai-readiness"))`.
- FE `AIReadinessClient.tsx`: `selectedCycle = sp.get('cycle')` → `effectiveCycleId = selectedCycle ?? …` →
  `useAIReadiness({ cycleId: effectiveCycleId, enabled: featureOn && cyclesQ.isFetched })`. So the FE DOES
  read `?cycle=` from the URL and would pass it as the data GET's `cycle` query param — **the deep-link is
  wired to the FE**. But the data query waits on `cyclesQ.isFetched` (the `/cycles` list), and the backend log
  showed 0 `/cycles` completions — so a page-navigation probe alone can't tell "frozen path slow" from
  "cycles-gate never clears."
**Choice:** disambiguate with a DIRECT authenticated dual-endpoint probe (D2), not a page navigation.

## D2 — the cheap dual-endpoint DIRECT probe → the deep-link premise is FALSIFIED
**Context:** the brief mandated verifying the frozen fast path BEFORE any cockpit/manifest edit or sweep.
**Choice:** authored `probe-aireadiness-deeplink.spec.ts` — log in as Dana, lift her bearer from a real
outbound request, then hit `/cycles` AND the frozen data GET `?cycle=<closed>&includePeople=true` DIRECTLY via
`page.request` (bypassing the FE React Query gate), with timing.
**Finding:**
- `/cycles` → **200 in 40 ms** (fast; returns the closed cycle — the FE gate is fine).
- the frozen data GET `?cycle=<closed>` → **NEVER COMPLETED** (180 s timeout) — identical to the live default.
- Root cause (source read, zero platform edit): `buildResponseFromSnapshots` (`ai_readiness.go:512`) reads the
  frozen scores fast, then calls `loadMembers(orgID, "")` (line 520) — a FULL UNBOUNDED org-member hydration
  (`hydrateMembers(nil, nil)` → whole-org tag/skill/sim aggregation) to attach CURRENT identity/tags to each
  snapshot. At 200 members that member-load IS the org-scale wall (the `ai_readiness_refresh` worker logs
  `context deadline exceeded`). It is NOT the demo-patchable targetRole authz RPC — `queryBaseMembers` reads
  `jobRole` from a SQL column, so `app-targetrole-authz-skip` (already applied) does nothing for it.
**Why this is a strictly STRONGER falsification than iter-07:** iter-07 concluded only "the default FE GET
omits `?cycle=`, so the frozen branch is never selected." iter-08 shows the frozen branch, even when selected
by a direct `?cycle=` request, is ITSELF org-scale-slow — so the deep-link (the user's chosen zero-edit fix)
cannot clear the wall even in principle. The cockpit/manifest deep-link edits were therefore NOT made (they'd
ship an inert `?cycle=` that still hangs), and the GATED sweep was NOT run (it would only reconfirm
`failingSections=5`).

## D3 — the USER-BLOCKER (the deep-link is falsified; the disclosed-allow needs sign-off)
**Context:** with the deep-link falsified, every path to `(0,0)` needs a user/architectural decision the
build-iter invariant forbids taking unilaterally.
**Finding / options (full text in progress.md's USER-BLOCKER block):**
  (a) **DISCLOSED residual** (the run-5 brief's stated fallback): data proven-correct in the DB; disclose the
      slow-but-correct sections as presenter-notes (green-with-disclosure). Per the coverage-protocol + iter-07
      this needs the user's EXPLICIT sign-off — it is NOT an auto-allow, and self-granting it would be
      loosening the gate to force a pass (the invariant's hard line). NOT self-granted.
  (b) **ESCALATE a platform edit** (Re-scope trigger): bound `loadMembers` in the snapshot path (pass the
      snapshot user-ids as `restrictUserIDs` instead of a whole-org hydration) OR add the deferred
      `frozen_tags jsonb` column (M314b) so the snapshot read needn't re-join live members. Forbidden here.
  (c) a NEW app read-path demo-patch bounding `loadMembers` in the snapshot path (the authz-skip precedent) — a
      substantial new tooling investment NOT in the run-5 brief's chosen scope.
**Why user-blocker (Phase 5 §4):** closing the gap changes what would land, and every option needs a decision
the invariant reserves for the user (a needs sign-off; b/c touch platform / new tooling). iter-08 surfaces the
decision. The closed-cycle seed + the diagnostic probe are KEPT + committed (correct + reusable regardless of
the resolution). `fit-up-m51` is NOT tagged.
