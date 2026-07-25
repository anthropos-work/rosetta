**Type:** tik (DRIVE) — TOK-01 cluster 1. Cold reset-to-seed on billion, gate (a).

# M254 · iter-02 — progress

## Work
- Investigated billion state as devops (marco is docker-blind → false-empty `docker ps`): a stale v2.6
  demo-1 was up (17 containers). Located the devops `panorama` workspace + rext clone.
- Re-pinned rext → `july-jitter-m253-studio-first-paint` (cumulative, on origin), bumped `.agentspace/rext.tag`.
- Confirmed the v2.7 `clones.pin.json` (12 repos, skillpath excluded, app bumped 3df8536); 5 clones were stale.
- Cold reset-to-seed: `rosetta-demo down 1 --purge` → serve reset → `DEMO_ADVANCE_CLONES=pinned` bring-up.
  Bring-up COMPLETED (exit 0), demo up on the consolidated platform.
- Verified peer reachability + the single autoverify blocker (drifted AI-readiness demopatch path).

## Close — 2026-07-24

**Outcome:** cold reset-to-seed COMPLETED; fresh v2.7 CONSOLIDATED demo up on billion (16 containers, 0
skillpath, peer-reachable). autoverify `green:false` — the SOLE failed check is the drifted
`app-aireadiness-snapshot-loadmembers` demopatch (target moved to `internal/aireadiness/readiness.go` in the
consolidation). Gate (a) NOT MET pending that one patch re-point.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET (a — autoverify green:false; 1 drifted demopatch; all other checks green)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n
(the coordinator's "re-drive, bring-up died" was premised on a docker-blind false-empty; verified the
bring-up actually SUCCEEDED — not a blocker, no re-drive) — (5) cap-reached: n (tik 1 of 5) —
(6) protocol-stop: n — Outcome: continue (iter-03 = the patch re-author)
**Decisions:** D1 (precondition correction), D2 (reset-to-seed method), D3 (coordinator-diagnosis correction —
docker-blind trap), D4 (gate-a blocker = drifted demopatch path).
**Side-deliverables:** none.
**Routes carried forward:** iter-03 — re-author `app-aireadiness-snapshot-loadmembers` for the consolidated app
(new path/anchor/replacement/shas) → commit + tag + push (rung-zero) → re-consume + rebuild app on billion →
re-verify autoverify green → gate (a) MET.
**Lessons:** (1) billion's REAL state must be read as the **devops** driver (or root for docker) — marco is
NOT in the docker group, so `docker ps` as marco silently returns empty (the docker-blind trap → a false
"clean slate / bring-up died" reading). (2) A completed detached bring-up + a fresh `autoverify.json` is the
ground truth, not a peer's container count. (3) The consolidation (skiller/skillpath→app) MOVED
`internal/workforce/ai_readiness.go` → `internal/aireadiness/readiness.go` — any demopatch pinned to the old
path silently skips; M254's live bring-up is what surfaces it.
