**Type:** tik (demopatch hash re-pin + next-web rebuild → M42m manager coverage GREEN). Under TOK-01 move (4).

# iter-15 — tik progress

## Execution log
1. **Baseline manager sweep** (demo-1, Dan @ Cervato, patches NOT yet applied): **reachable=70/150,
   failingSections=0, personaFailures=0, crossPortFollowFails=0, notReached=0, escapes=68** — the SOLE
   blocker is prod-eject URL escapes (STUDIO_URL + PUBLIC_WEBSITE_URL rendered across every manager surface).
   Confirms every section has real content (no data gaps — the sim-embeddings fill also feeds the manager's
   Library links), persona coherent, no perf wall at Cervato's 221 members.
2. **Diagnosed the demopatch drift:** both `next-web-studio-url` + the chained `next-web-public-website-url`
   G2-REFUSED (run-3 evidence + iter-12 log): the skiller-in-app merge bumped next-web-app to **v2.106.1**,
   drifting sibling exports in `packages/core-js/src/constants/urls.ts` → the file-level `pre_sha256` no
   longer matched. The STUDIO_URL + PUBLIC_WEBSITE_URL anchor hunks are byte-identical (grep-confirmed).
3. **Computed the 4 re-pinned hashes** via the demopatch's OWN `manifest_loader` + `body.replace(anchor,
   replacement,1)` + `sha256_text` (the exact apply code path): studio pre `0d4c3790` / post `fe15aa71`;
   public-website pre `fe15aa71` (chained ✓) / post `d92fa701`. Anchors occur exactly once, post_markers
   present.
4. **Re-pinned the 2 demopatch yamls** in the rext authoring copy (+ version-reference comments); committed
   `84e15e9`, moved tag `quick-change-m211`, **re-synced the consumption clone** `stack-demo/rosetta-extensions`
   (local tag-fetch) → both clones carry the new hashes.
5. **Rebuilt demo-1 next-web** WITH the patches: `docker image rm demo-1-next-web` + `build_frontend_next_web`
   → the demopatches now **APPLY cleanly** ("next-web-studio-url applied", "next-web-public-website-url
   applied", "next-web-members-pagination applied" — no G2 refuse), image bakes the demo-local URLs into the
   JS bundle, demo clone urls.ts **reverted clean** (trap; sha back to `0d4c3790`, git clean). Recreated the
   next-web-app container `--no-deps --force-recreate` (postgres untouched — M46 warning). build-rc 0.
6. **Re-ran the manager sweep:** **GATE: MET ✅** — reachable=70/150, **escapes 68→0**, failingSections=0,
   personaFailures=0, crossPortFollowFails=0, notReached=0, frontier=EXHAUSTED. (The demo-local
   `/library/job-simulations/<slug>` 404 is expected + fine — it's demo-local, not a prod eject, per the
   PUBLIC_WEBSITE_URL patch design.)

## Re-measurement (gate sub-conditions)
| Sub-condition | Pre-iter | Post-iter |
|---|---|---|
| (e) M42 coverage — employee vantage | GREEN (iter-14) | GREEN |
| (e) M42 coverage — **manager vantage** | NOT MET (escapes 68) | **GREEN** (escapes 0, GATE MET) |
| (e) v2.0 Playthroughs | NOT MET | NOT MET (next: iter-16) |
**Metric:** M42m manager escapes **68 → 0**, manager GATE MET. **Both M42 vantages now GREEN** → the
**coverage half of sub-condition (e) is COMPLETE**. Overall composite gate still needs v2.0 Playthroughs +
cold /dev-up.

## Close — 2026-07-08

**Outcome:** Cleared the M42m manager gate's sole blocker (68 prod-eject escapes) by re-pinning the two drifted
next-web URL demopatch hashes to the v2.1 next-web source (v2.106.1 — pure re-anchor, hunks byte-identical,
mirrors M49 #8) + rebuilding next-web with the patches baked. Manager coverage **GATE MET** (escapes 68→0,
failingSections/persona/crossPort/notReached all 0). With iter-14's employee GREEN, the **full M42 coverage
sweep (both vantages) is GREEN**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (both M42 vantages GREEN, but the composite M211 gate still needs v2.0 Playthroughs + cold /dev-up)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (tik with +progress) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (the demopatch drift is a FILE-level sha re-anchor, not a hunk change — the STUDIO_URL/PUBLIC_WEBSITE_URL anchors are byte-identical to v1.10; compute the 4 chained hashes via the tool's own loader + replace path, the same maintenance M49 #8 did), D2 (the consumption clone `stack-demo/rosetta-extensions` — not the authoring copy — drives the BUILD, so re-pin must commit+tag the authoring copy THEN local-fetch the moved tag into the consumption clone), D3 (rebuild next-web via `build_frontend_next_web` from the consumption clone + recreate `--no-deps` — never touch postgres per the M46 warning)
**Side-deliverables:** none (the re-pin + rebuild are the planned deliverable).
**Routes carried forward (Fate-3 → next iters this session):** v2.0 Playthroughs suite (iter-16); cold `/dev-up` (final).
**Lessons:** A re-sync release that bumps a frontend repo drifts EVERY file-level-sha-gated demopatch on that repo's touched files — even when the patched HUNK is unchanged. The re-anchor is mechanical (recompute the chained hashes via the demopatch's own loader), but it MUST land on BOTH the authoring copy (commit + tag) AND the consumption clone the build reads (local tag-fetch) — an authoring-only edit is invisible to the build. This is the frontend-repo analog of iter-10's build-scratch re-sync for the Go services.
