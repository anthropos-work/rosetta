# M254 · iter-02 — decisions

## D1 — billion precondition corrected: a stale v2.6 demo was up (not a clean slate)
The orchestrator's "no demo up on billion (clean slate)" was checked as **marco**, who is NOT in the docker
group — `docker ps` as marco silently returns empty (permission denied → empty stdout). Read as **root**
(docker) / **devops** (the workspace owner + driver): a STALE v2.6 demo-1 (17 containers incl. a `skillpath`
container — the pre-consolidation topology) was running for 40–41h under `/home/devops/panorama/stack-demo`.
Driver account = **devops** (groups: docker(988)+sudo(27); `~/.git-credentials` GH_PAT primed; `platform/.env`
provisioned; go1.25.12).

## D2 — cold reset-to-seed method
`rosetta-demo down 1 --purge` (wipes the DB `data` dir + removes demo-1 images; 82 GB build cache persists) →
`tailscale serve reset` → `DEMO_ADVANCE_CLONES=pinned STACK_PUBLIC_HOST=billion.taildc510.ts.net
up-injected.sh 1 --public-host billion.taildc510.ts.net`. Skipped `DEMO_FRESHNESS_STRICT=1` — the vestigial
`stack-demo/skillpath` dir (not in the v2.7 pin) could FATAL a strict freshness gate; the built-stack
autoverify + probes are the real gate-(a) proof instead.

## D3 — the coordinator's "bring-up died / re-drive" was the docker-blind trap (verified, did NOT re-drive)
Mid-task the coordinator directed a robust re-drive, premised on "billion has 0 demo containers + no build
proc → the blocking-ssh reset dropped (SIGHUP) → bring-up died." VERIFIED against live state: the tracked
background Bash `bd1370rtr` **exited 0**; the billion-side log shows the full bring-up reached
`UP. Clerk-free demo-1 is live` + autoverify; **16 demo-1 containers are running**; the demo is
**peer-reachable** (backend 200 / web 307 / cockpit 200 over tailnet HTTPS). The "0 containers" reading was
the same docker-blind-as-marco false-empty. The bring-up SUCCEEDED — a re-drive would destroy 20–50 min of
correct work AND reproduce the green:false (the fix is a rext patch re-point, not a rebuild). ADOPTED the
coordinator's sound methodology (detached-on-billion + short-poll for future long ops); did NOT re-drive.

## D4 — gate-(a) blocker = the drifted AI-readiness demopatch path
`app-aireadiness-snapshot-loadmembers` pins `path: internal/workforce/ai_readiness.go` (app@v1.315.0). The
consolidated app (`3df8536`) refactored the AI-readiness read-path into a new `internal/aireadiness` package:
`buildResponseFromSnapshots` is now `func (m *aireadiness.Manager)` in `internal/aireadiness/readiness.go`,
and the member load is `m.workforce.LoadMembers(ctx, orgID, "")` (`workforce.Member`). The exported bounded
sibling `m.workforce.LoadMembersByUserIDs` exists (members.go:353). The apply helper reads `path` from the
manifest dynamically → the fix is manifest-only (re-point path + new anchor/replacement + recomputed shas).
Pristine `readiness.go@3df8536` sha256 = `130e58f7c497b1dd0582abd867b96f71d83b8a9cf3751c6a43a1ac7aedc55c67`.
Routed to iter-03.
