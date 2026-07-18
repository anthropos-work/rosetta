**Type:** tok (bootstrap). Authors TOK-01 `reprove-corrected-hiring-on-billion` from the M228 overview + the M226
proven recipe + the billion recon. Protocol: verification.md + coverage-protocol.md + latency-budget.md.

# iter-01 — work log

1. **Phase 0b KB-fidelity: GREEN** (`../kb-fidelity-audit.md`). All 8 topics PAIRED; the 4 M227 corrections ALIGNED
   across corpus (`hiring.md`) + the `casting-call-m227-sections` rext tooling; M226 shared-infra fold-ins present.
2. **billion recon** (see `../spec-notes.md § billion recon`):
   - The M226 `demo-1` IS UP — `docker ps -a` (as root) shows 17 `demo-1-*` containers Up 5h at rext
     `casting-call-m226-c2-race-fix`. (`docker ps` as `marco` = 0 → marco has no docker access; **devops** is the
     operator: groups `docker sudo`.)
   - Workspace `/home/devops/panorama/stack-demo` (devops-owned); rext consumption clone there @ c2-race-fix.
   - Bring-up entry `up-injected.sh` via `demo-stack/rosetta-demo`; teardown `rosetta-demo down 1 --purge`.
   - **20 demo-1-* images cached** (M227 = pure seed/content → images stay valid). 7 stale serve fronts.
   - billion mem 7.3 GiB + 15 GiB swap, disk 38 G avail. ssh users marco/root/devops (tailscale SSH).
3. **Authored TOK-01** `reprove-corrected-hiring-on-billion` (`../decisions.md`): the M226 recipe RE-RUN with
   `casting-call-m227-sections` — teardown → rext-only cutover → default cold `up-injected.sh 1` → measure the retuned
   7-condition gate + the 4 M227-correction render checks from this Mac → attribute → fix via tooling → 2 clean cycles.

## Close — 2026-07-17

**Outcome:** Bootstrap strategy TOK-01 authored — the M226 prove-on-billion recipe re-run against the M227-corrected
data (`casting-call-m227-sections`). Phase 0b KB-fidelity GREEN. billion recon complete: the M226 demo is UP at the
pre-correction tag, devops-operated, images cached — a clean teardown → rext cutover → default cold up is the path.
**Type:** tok (bootstrap)
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (bootstrap, does NOT exit) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (0 tiks) — (6) protocol-stop: n — Outcome: continue (iter-02, a tik under TOK-01)
**Decisions:** iter-01 D1 (devops is the billion operator user); D2 (rext-only cutover — M227 changed no image/patch).
**Side-deliverables:** none.
**Routes carried forward:**
- iter-02 (tik): teardown the M226 demo + verify clean from this Mac + rext cutover → `casting-call-m227-sections` +
  default cold `up-injected.sh 1` synchronously + FIRST retuned-7-condition + 4-M227-render-check measurement from this
  Mac. Handler `PROVE-M228-iter02-first-corrected-cold-bringup`.
**Lessons:** The billion state deviated from the task's stated assumption in a way that mattered — `docker ps` as the
login user (marco) returned 0 containers, reading as "demo down," but `docker ps -a` as **root** showed the demo fully
UP; marco simply lacks docker-group access. Always run the docker/tooling recon as the OPERATOR user (devops/root), not
the login user. The M227 corrections being pure seed/content tooling makes the cutover cheap (rext-only, cached images).
