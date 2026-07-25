---
iter: 03
milestone: M254
iteration_type: tik
status: closed-fixed
created: 2026-07-24
---

# M254 · iter-03 — fix the drifted AI-readiness demopatch → gate (a) GREEN

**Type:** tik · **Active strategy:** TOK-01 (cluster 1 — the DRIVE, gate a green). Routed from iter-02.

## Target
Make the cold reset-to-seed autoverify GREEN by fixing the sole blocker — the
`app-aireadiness-snapshot-loadmembers` demopatch that skipped on the consolidated app (target path drifted).

## What landed
- **Re-authored the manifest** for the consolidated app (`3df8536`): `path`
  `internal/workforce/ai_readiness.go` → `internal/aireadiness/readiness.go`; new anchor
  (`members, _ := m.workforce.LoadMembers(ctx, orgID, "")` + `workforce.Member`); bounded via the exported
  `m.workforce.LoadMembersByUserIDs`; recomputed `pre_sha256 130e58f7` / `post_sha256 ab611ed5`.
- **Verified locally:** manifest_loader parses; `apply_patch.py --check` → "pinned sha matches; anchor 1x. OK";
  `--apply` → on-disk sha == post_sha256; re-apply → idempotent no-op; gofmt clean; **52/52 demopatch tests green**.
- **Shipped (rung-zero):** rext commit `997272b`, tag `july-jitter-m254-aireadiness-repoint` pushed to origin.
- **Re-verified on billion (robust: detached + poll):** re-pinned billion → new tag; cold reset-to-seed
  (`down --purge` → serve reset → advance-pinned bring-up) launched DETACHED (`setsid nohup`) + polled from a
  local bg Bash. Completed ~12 min. **autoverify `green:true, 0 warnings`** — demopatch.log EMPTY (the patch
  applied), 16 containers, 0 skillpath, peer-reachable (backend 200 / web 307 / cockpit 200).

## Gate (a): MET ✓
The re-grounded stack builds + comes up GREEN on the consolidated platform (3 subgraphs, skillpath-in-app),
fresh green `autoverify.json` (0 warnings), cold reset-to-seed, driven from a tailnet peer, 0 platform edits.

## Next
Gate (a) is the enabling precondition. iter-04+ work gates b–h per TOK-01 (read-only sweeps fan-out →
latency solo → mutating tail) against THIS green bring-up.
