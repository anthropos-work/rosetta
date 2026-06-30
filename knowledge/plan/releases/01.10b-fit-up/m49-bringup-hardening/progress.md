# M49 — progress

## Section checklist
_Checked off as each In-scope deliverable lands. Close when all boxes are ticked._
_Each section = one rext code fix (committed in `.agentspace/rosetta-extensions` @ `main`, rolled into tag `fit-up-m49`) + its corpus doc truth-up (committed in `rosetta` @ `m49/bringup-hardening`)._

- [x] **§1 rext-tag-sot** (#1) — `.agentspace/rext.tag` SoT + `lib/rext_tag.sh` reader + non-fatal pin guard in `ensure-clones.sh`; reconciled **4** conflicting prose pins (the overview's 3 + a stray in `verification.md`) to read from it. Pin = `fit-up-m49`. Docs: `rosetta_demo.md`, `demo-up/SKILL.md`, `frontend-tier.md`, `verification.md`. 96/96 tests green.
- [x] **§2 env-guard-order** (#3) — moved the `.env`-presence guard below the M30 provision block (now checks `BASE_ENV`); a stack-demo-only box provisions-then-checks instead of aborting. Regression test added. Doc: `rosetta_demo.md` bring-up-order step 2. 149/149 tests green.
- [x] **§3 invitation-hmac-secret** (#4) — added `INVITATION_HMAC_SECRET` as a critical/required platform/.env gene + `DemoGeneratedKeys` demo overlay (so demo pre-flight treats it satisfied) + up-injected auto-gen (openssl, values-blind, idempotent, non-fatal). Doc: `secrets-spec.md` (split now 40/8/8 + 13-crit, 56-gene; AI-keys policy noted as M50). 6 count-literals reconciled corpus-wide. Go + 98 demo-stack tests green.
- [x] **§4 ant-academy-clone** (#5) — ensure-clones.sh phase (d2) clones ant-academy EXPLICITLY (non-fatal), NOT via repos.yml (ephemeral platform clone = non-durable + platform edit). Static + 4 functional tests. Docs: `ant-academy.md` (3 spots), `CLAUDE.md` corrected (the M48 "M49 adds repos.yml entry" prediction → the actual explicit-clone fix). 103 demo-stack tests green.
- [x] **§5 disk-preflight-down-cleanup** (#6) — `preflight_disk_headroom()` (non-fatal, mirrors the RAM check, offers prune, `DEMO_DISK_MIN_GIB`/`DEMO_DISK_AVAIL_KB`) + `cmd_down --purge` removes the stack's `demo-N-*` images (scoped, non-fatal; plain `down` keeps them). 5 functional + 2 static tests. Docs: `frontend-tier.md` + `demo-up/SKILL.md`. 162 tests green.
- [x] **§6 nonfatal-frontend** (#7) — before `compose up`, scale any absent-image frontend to 0 (`--scale svc=0`, set-u-safe) so a failed UI build no longer aborts backend + set-dress + verify + cockpit. Static pins (image check, scale-to-0, NO_UI gate, ordering). Doc: `frontend-tier.md` (the "non-fatal" claim is now true). 163 tests green.
- [ ] **§7 demopatch-reanchor** (#8) — re-anchor `next-web-studio-url` demopatch `pre_sha256`/`post_sha256` to the current `stack-demo/next-web-app` source (v2.89.0). Doc: `corpus/ops/rosetta_demo.md` / `corpus/ops/demo/frontend-tier.md`.

### Close gate
- [ ] All 7 rext fixes committed in the authoring copy + tagged `fit-up-m49` (annotated).
- [ ] Both working trees clean.
- [ ] Static/code-level verification only (no live `/demo-up` — that's the orchestrator's live-verify gate + M53).
