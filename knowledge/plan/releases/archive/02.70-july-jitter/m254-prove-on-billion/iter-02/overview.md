---
iter: 02
milestone: M254
iteration_type: tik
iter_shape: drive
status: closed-fixed-partial
created: 2026-07-24
---

# M254 · iter-02 — the DRIVE (gate a): cold reset-to-seed on billion

**Type:** tik · **Active strategy:** TOK-01 (cluster 1 — the DRIVE).

## Target
Cold reset-to-seed `/demo-up` on billion at the v2.7 pin `july-jitter-m253-studio-first-paint`, on the
CONSOLIDATED platform (3 subgraphs, skillpath-in-app). Assert gate (a): builds + comes up GREEN, fresh
green `autoverify.json`.

## What landed
- **Precondition corrected:** billion was NOT a clean slate — a STALE v2.6 demo-1 (17 containers incl. a
  `skillpath` container) was up under `/home/devops/panorama/` (driver = **devops**: docker+sudo+PAT+.env).
- Re-pinned the rext clone → v2.7 + bumped `.agentspace/rext.tag`; wrote a cold reset-to-seed runner
  (`down --purge` → `tailscale serve reset` → `DEMO_ADVANCE_CLONES=pinned` bring-up).
- The bring-up **COMPLETED** (exit 0): `UP. Clerk-free demo-1 is live`. Fresh v2.7 consolidated demo up —
  **16 containers, NO skillpath container (consolidated ✓)**, 5 stale clones advanced to pin
  (platform/app/next-web/academy/graphql), health 200, casbin 1250, taxonomy 42790, cockpit + clerkenstein
  answering, hiring set-dressed, academy catalog rendering, studio-desk AI key present.
- **Peer-reachable from this workstation** (tailnet HTTPS): backend `:18082/api/health` → 200, apps/web
  `:13000` → 307, cockpit `:17700` → 200.

## Gate (a): NOT MET — one blocker
`autoverify.json` = `green:false, warnings:1` — **one FAILED check**: the `app-aireadiness-snapshot-loadmembers`
demopatch was SKIPPED ("target `internal/workforce/ai_readiness.go` not found in the app clone"). The
AI-readiness read-path was refactored in the consolidation → the file moved to
`internal/aireadiness/readiness.go` (new `aireadiness` package). The single drifted patch is the sole
green-blocker. Exactly the drift M254 exists to catch → route to rext (0 platform edits).

## Routes carried forward → iter-03
Re-author the `app-aireadiness-snapshot-loadmembers` manifest for the consolidated app (path
`internal/workforce/ai_readiness.go` → `internal/aireadiness/readiness.go`; new anchor
`members, _ := m.workforce.LoadMembers(ctx, orgID, "")` + `workforce.Member`; bound via the exported
`m.workforce.LoadMembersByUserIDs`; recompute pre/post sha256). Re-consume on billion + rebuild app image +
re-verify autoverify green → gate (a) MET.
