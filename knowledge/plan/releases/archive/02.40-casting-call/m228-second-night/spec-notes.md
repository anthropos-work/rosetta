# M228 — Spec notes

_Iterative milestone: this file accumulates iteration-protocol-specific technical notes (live-run transcripts,
per-condition evidence, latency measurements, the M227-correction render checks). Per-iter detail lives in `iter-NN/`._

## The 7-condition live billion proof (RETUNED for M227) + the 4 M227-correction render checks
_Per-condition evidence: org present + is_hiring + exactly 5/45; ≥6 rows/each of 5 positions (each candidate on 1 sim,
~8/position); candidate profiles usable (external emails + matched avatars); reads-as-hiring + hiring-only content;
recruiter p95 click→ACCESS < 5 s; coexists with 3 workforce orgs; 0 platform edits._

## billion recon (iter-01, 2026-07-17)
- **The M226 demo IS UP** — `docker ps -a` (as root) shows 17 `demo-1-*` containers "Up 5 hours" at rext tag
  `casting-call-m226-c2-race-fix`. (`docker ps` as `marco` returns 0 — marco has no docker-group access; **devops** is
  the operator user: groups `docker sudo`.)
- Workspace: `/home/devops/panorama/stack-demo` (owned by `devops`). rext consumption clone
  `/home/devops/panorama/stack-demo/rosetta-extensions` @ `casting-call-m226-c2-race-fix` (`.agentspace/rext.tag`).
- Bring-up entry: `up-injected.sh` via `demo-stack/rosetta-demo`. Teardown: `rosetta-demo down 1 --purge`.
- **20 demo-1-* images cached** (next-web 4.64 G, hiring 4.56 G, app/cms/jobsim/skillpath injected, sentinel/storage/
  roadrunner/graphql/postgres/fapi/bapi/studio-desk). M227 = pure seed/content tooling (no new demo-patch, no image
  change) → the cached images stay valid; only the rext tooling changes + a re-seed at the new tag applies the fixes.
- **7 stale tailscale serve fronts** (13000/13001/13077/15050/17700-cockpit/18082/19000) — the live M226 demo's fronts.
- billion: mem 7.3 GiB + 15 GiB swap; disk 38 G avail (81% used). ssh: `marco@` + `root@` + `devops@` (tailscale SSH).
- Access model: run the demo tooling as **devops** (`ssh devops@billion.taildc510.ts.net`); assert from THIS Mac (peer)
  against billion's offset ports (`RENDER_HOST=billion.taildc510.ts.net RENDER_APP_SCHEME=https`, `LATENCY_HOST=…`).

## Pre-flight audits — iter-01
**KB-fidelity (Phase 0b, 2026-07-17): GREEN.** Report: `kb-fidelity-audit.md`. All 8 milestone-scope topics PAIRED;
the M227-correction delta (the only change vs the M226 GREEN 8 h earlier) is ALIGNED in BOTH corpus (`hiring.md`) and
the `casting-call-m227-sections` rext tooling:
- fix#1 hiring-only → `hiring_scope.go IsHiringOrg()` (#M227-D1); fix#2 external emails → `userprofile.go
  externalCandidateDomains` role-keyed (#M227-D2); fix#3 1-sim/candidate + floor `≥40→≥6` → `hiring_funnel.go`
  even round-robin + `RENDER_GATE_FLOOR ?? '6'`; fix#4 gender avatars → `gender.go` + `users.go:180`.
- M226 shared-infra fold-ins present: recruiter 3rd vantage + serve prereq (`latency-budget.md`), apps/hiring
  `3001+off` serve front (`tailscale-serve.md`). All `hiring.md` cross-refs resolve. No blind areas, no stale claims.
