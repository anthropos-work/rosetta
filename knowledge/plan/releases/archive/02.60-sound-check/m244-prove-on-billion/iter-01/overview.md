---
iter: 01
milestone: M244
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
created: 2026-07-22
---

# iter-01 — bootstrap tok (author TOK-01 strategy)

**Type:** tok (bootstrap) — iter-01 of the milestone; authors the FIRST strategy. Does NOT terminate the call.

## Inputs
- `overview.md` (the multi-part exit gate a–h), `spec-notes.md`, `roadmap.md` M244 detail.
- Protocol docs: `corpus/ops/verification.md` (pre-flight rung zero + autoverify), `corpus/ops/demo/tailscale-serve.md` (the billion `--public-host` runbook), `coverage-protocol.md` (the two Playwright sweeps), `playthroughs.md` (the functional-flow runner).
- Phase 0b KB-fidelity audit: **YELLOW** (see `kb-fidelity-audit.md`) — no blind areas; denominator is **49** (not 29); spec count is **40** (not 39).
- Pre-flight rung zero: **billion is BARE** (no stack/panorama/rext.tag/cache); the pin billion must consume is **`sound-check-m243-assign-write-playthrough`** (on origin = `2ef5962`); local `.agentspace/secrets` (24K) + `.agentspace/snapshots` (1.4G) exist to seed billion.

## Initial strategy (TOK-01)
**Stage a cold, green, reset-to-seed demo on `billion` at the m243 pin, then discharge the exit gate parts (a–h) in dependency order — one cluster per tik — driving/asserting from a tailnet PEER (this workstation), recording live evidence.** Primary metric: **gate parts discharged green live (0/8 → 8/8)** + the 2 inherited defer-items (DEF-M239-01 ENOSPC loud-fail, reap-17700 standing-9) + the conditional DEF-M240-01 video exhibit. 0 platform edits (walls route to sha-pinned demopatches or ESCALATE).

**Tik plan (dependency-ordered):**
- **tik-1 (iter-02) — Foundation + gate (a).** Pre-flight rung zero on billion (set `.agentspace/rext.tag`=m243, workspace layout, PAT-over-HTTPS, scp secrets + snapshot cache, clone+checkout rext at the pinned tag, `git describe --exact-match`). Run **gate (a) ORG-CLEAN read-only FIRST** (the content-story scrub-cleanliness check — 0 surviving source-org tokens). Then kick the **cold reset-to-seed** `up-injected.sh 1 --public-host billion.taildc510.ts.net` through a **LOGIN shell** (F2b). Drive to a **fresh green `autoverify.json`** (the gate every downstream sweep reads).
- **tik-2 (iter-03) — bring-up green + gate (h).** Confirm autoverify fresh+green from a peer; drive **p95 click→ACCESS < 5 s** both hero vantages (`latency.spec.ts`) + smoke the v2.6 fixes (academy course-start, talk-to-data live answer, library, language toggle, cockpit UX).
- **tik-3 (iter-04) — gate (b).** `run-content-stories.sh` green at **49/49** from a peer (CQ-1 grader + CQ-2 wiring + external denominator), gated on fresh green autoverify.
- **tik-4 (iter-05) — gate (c).** The **40 live-browser specs** (coverage + playthroughs incl. `assignment-assign`) green from a peer.
- **later tiks —** gate (d) academy /library+/free twin; gate (e) DEF-M226-01 serve-reap **test-or-DROP**; gate (f) the 3 v2.3 drift-carries burned-in; gate (g) the interview plan-section-id **alignment assertion** (build + green); the inherited DEF-M239-01 + reap-17700 + DEF-M240-01 (video exhibit IF Bunny keys reachable, else voice presence-only).

**Rationale:** the live-proof lineage (M215/M221/M226/M228/M236) shows the cold billion bring-up is the critical path + highest-risk step (F1–F12 host findings, the login-shell PATH trap, the M217 pin trap, the peer-vs-VM TLS trap). Front-load de-risking: verify the pin lands (rung zero), drive through a login shell, assert from a tailnet peer (never the VM — `tailscale serve` is bypassed on loopback). Everything downstream (sweeps, playthroughs, latency) **gates on a fresh green autoverify**, so a green cold bring-up is the enabling precondition for gates (b)(c)(h) — hence it is tik-1.

## Escalation conditions
- A platform-source wall with no config/env/demopatch seam → ESCALATE (0 platform edits is a hard gate).
- billion unreachable / un-provisionable (no usable secret source or snapshot cache) → user-blocker.
- A gate part that cannot be met without a design change → route forward / tok.

## Acceptable close-no-lift outcomes (for later tiks)
- DEF-M226-01 serve-reap: a documented DROP (the claim actively falsified) satisfies gate (e) as much as a passing test.
- DEF-M240-01 video: voice **presence-only** is the shipped deliverable if Bunny keys are not reachable on billion.

## Phase plan (this iter)
Bootstrap tok: author TOK-01 (done above) → record in milestone-root `decisions.md` → close. No metric move (tok).
