**Type:** tik — Foundation (under TOK-01). Active strategy: TOK-01 (staged cold billion bring-up → gate-parts a–h).

# iter-02 — progress

## Step 0 re-survey (mandatory) — TOK-01 target SUBSTITUTED
TOK-01 named tik-1 as "from-scratch billion setup + scp secrets/cache." Re-survey found reality differs: **a 47-hour-old demo-1 was already running** under user `devops` at `/home/devops/panorama/` (17 containers up, full `tailscale serve`), but **M238-era and mis-pinned** (`rext.tag`=m237, clone checkout=m238/`5b13ec4`, autoverify 47 h stale) — it lacks all of M239–M243, so it cannot be the M244 proof. The devops workspace **already has** `.agentspace/secrets` (60K) + `.agentspace/snapshots` (1.5G) + the rext clone, so **no from-scratch setup and no 1.4G scp needed**. Substitution under the same TOK-01 strategy: **adopt the devops workspace → re-pin to m243 → teardown the stale stack (capturing serve-reap for gate e) → cold reset-to-seed bring-up.** (billion facts: marco lacks docker group [sudo only]; `devops` is the docker-group deploy user; `ssh devops@billion` works directly; toolchain present go1.25.12 / atlas v1.2.4 / tailscale 1.98.8; disk 98G free.)

## Work done
1. **Pre-flight rung zero — GREEN.** Set billion `.agentspace/rext.tag` = `sound-check-m243-assign-write-playthrough`; `git fetch --tags` + `checkout -f`; `git describe --exact-match` = the tag; HEAD = `2ef5962` (= origin = local authoring HEAD). Billion can now OBTAIN the tooling under test (the M217 pin guard will pass).
2. **Gate (a) ORG-CLEAN — PASS (live on billion).** Content-story fixture cleanliness + scrub suites (`stack-seeding/contentsession` + `scrub`) run `-count=1` on billion's m243 clone → both OK; **23 content-story fixtures inspected, 0 surviving source-org/PII tokens.** Read-only, before the bring-up, as the gate specifies. (Also PASS locally, byte-identical fixtures.)
3. **Gate (e) DEF-M226-01 serve-reap — TESTED + CONFIRMED (live on billion).** Captured `tailscale serve status` before/after `demo-stack/rosetta-demo down 1 --purge`: **7 offset serve ports (13000/13001/13077/15050/17700/18082/19000) → 0** after teardown; 17 containers → 0; data purged (the F9 UID-1001/0700 path exercised: "the host user cannot unlink it" handled); images reclaimed; hostlock released. **The serve-reap self-resolution claim holds** — gate (e) discharged as *actively tested + passing* (not dropped).
4. **Cold reset-to-seed bring-up — LAUNCHED.** `STACK_PUBLIC_HOST=billion.taildc510.ts.net bash demo-stack/up-injected.sh 1 --public-host billion.taildc510.ts.net` via a login shell as devops, teed to `~/m244-bringup.log`; supervised via a poll Monitor + the background task. Awaiting a fresh green `autoverify.json`.

## Cold bring-up RESULT — green
- **BRINGUP_EXIT=0** (start 12:37:15Z → done ~12:51Z, ~35 min full cold build+seed).
- **Fresh green `autoverify.json`**: `{"project":"demo-1","offset":10000,"warnings":0,"green":true,"ts":"2026-07-22T12:51:20Z"}` — written by THIS bring-up, 0 warnings (so the M22 cheap-wins all passed: backend health, casbin>0, directus_collections>0, no-prod-read, hiring set-dress ≥5/≥40).
- **Peer-vantage origins (from this workstation)**: web `:13000`→307, hiring `:13001`→307, cosmo `:15050`→200, backend `:18082`→404 (no root route; health is 200), academy `:13077`→200, **cockpit `:17700`→200**. All serving over the trusted MagicDNS cert.
- **Content populated**: `public.skills`=**42,790** (full taxonomy — snapshot replay worked), `directus.directus_collections`=**21**, `jobsimulation.sessions`=**1,644**.
- **content-manifest.json served** by the cockpit: **4 products / 28 sessions → forms exactly 49 pairs** (23 sim×2 + 2 skillpath-legacy-player + 1 skillpath-new; ai-labs excluded) = gate (b)'s external denominator. Precondition met.
- The `demo-1-directus-1 Exited(1)` seen mid-poll was **transient** (part of the bootstrap→apply→replay→boot provision sequence); directus ended **Up + healthy**, autoverify green.

## Poll-tooling lesson (recorded)
The first poll flagged a FALSE "failed": it curled the tailnet origins **FROM the VM** (the M219 loopback-TLS trap → 000) and tripped on the transient directus exit. Fixed poll: curl origins **from the peer** (this workstation), and treat **BRINGUP_EXIT** (from up-injected) as the authoritative completion signal, never a container-state guess.

## Close — 2026-07-22

**Outcome:** billion cold reset-to-seed GREEN at m243 (BRINGUP_EXIT=0, fresh green autoverify 12:51Z, all peer origins serving, content fully populated). Gate parts discharged this iter: **(a) ORG-CLEAN** + **(e) DEF-M226-01 serve-reap** — plus rung-zero GREEN + the enabling green stack. Metric 0/8 → **2/8**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (2/8 — needs all a–h)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik #1 of session) — (6) protocol-stop: n — Outcome: continue (iter-03)
**Decisions:** D1 (remote-foreground supervision), D2 (adopt devops workspace substitution) — iter-02/decisions.md
**Side-deliverables (if any):** none (0 platform edits; re-pin + teardown + bring-up are demo-side).
**Routes carried forward:** gates (b)(c)(d)(f)(g)(h) + p95 + inherited DEF-M239-01/reap-17700/DEF-M240-01 → iter-03+.
**Lessons:** (1) always curl tailnet origins from a PEER, never the VM (M219). (2) BRINGUP_EXIT is the authoritative bring-up verdict; container-state polls false-trip on transient provision exits. (3) a bare box may already carry a stale mis-pinned demo under a different deploy user — re-survey (Step 0) before assuming "bare."
