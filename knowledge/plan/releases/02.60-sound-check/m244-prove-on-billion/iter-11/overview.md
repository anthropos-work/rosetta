---
iter: 11
milestone: M244
iteration_type: tik
status: planned
created: 2026-07-22
---

# iter-11 — the gate-(h) cold reset-to-seed (the critical path)

## Type
tik (under TOK-01 — staged cold billion bring-up → gate-parts a–h one-cluster-per-tik).

## Step 0 — re-survey (mandatory)
- **Rung zero GREEN:** rext consumption tag `sound-check-m244-content-sweep-robustness` is on origin (annotated tag-object 7dbad4b, peels to **b38ad75**) — billion can OBTAIN the m244 tooling (M217 pin guard will pass).
- **Billion state:** demo-1 is UP + serving at the OLD **m243** pin (`sound-check-m243-assign-write-playthrough`) — 17 containers 6h, peer origins web :13000→307 / cockpit :17700→200 / hiring :13001→307. Workspace `/home/devops/panorama` (operate as devops via root@). rext clone `stack-demo/rosetta-extensions` checked out at m243; entrypoint `demo-stack/up-injected.sh`. Disk 86G free (no ENOSPC risk).
- **TOK-01 target still current:** the gate-(h) cold reset-to-seed is the next-queue critical path — unchanged, still untouched. No substitution.

## Active strategy reference
TOK-01 (bootstrap) — this tik executes the enabling re-seed that TOK-01 sequenced last (all remaining gate parts b/c/d/f/h gate on a fresh-green autoverify at the m244 pin).

## Cluster / target identified
Re-pin billion's rext to `sound-check-m244-content-sweep-robustness` (b38ad75, FROM ORIGIN) and cold reset-to-seed (teardown --purge → `up-injected.sh 1 --public-host billion.taildc510.ts.net`). This re-bakes the iter-08 interview-player-ack demopatch into the next-web/hiring images, re-seeds the iter-06 intv-voice-fail v1.4 re-pin + iter-07 voice player-presence-only + iter-05 simulations_extraction capture. Immediately after: verify fresh-green autoverify + re-verify the content acceptance on the fresh seed (gate (a) ORG-CLEAN, gate (b) content-stories 47/47, gate (g) interview report renders live + alignment assertion green).

## Hypothesis
A cold reset-to-seed at the m244 pin produces a fresh-green `autoverify.json` and — because it bakes/seeds all iter-05→08 fixes — makes gate (b) count **47/47 live** (the 3 residuals resolved in tooling now land on the fresh seed), moving the metric 3/8 → 4/8, with gate (a) + gate (g) re-confirmed live on the fresh seed.

## Expected lift
+1 gate part (gate b → green live, 3/8 → 4/8). Gates (a)/(g) re-confirmed (already counted). Foundation re-established at the m244 pin for iters 12+ (gates c/f/h).

## Phase plan (protocol: verification.md + tailscale-serve.md + coverage-protocol.md)
1. Re-pin: fetch --tags origin + checkout b38ad75 in the rext clone (`git describe --exact-match` MUST print the tag; peel MUST = b38ad75) + update `.agentspace/rext.tag`.
2. Teardown demo-1 `--purge` (serve-reap already gate-e; teardown is version-stable).
3. Cold bring-up `up-injected.sh 1 --public-host billion.taildc510.ts.net` under nohup on billion (survives SSH drop → never-kill-a-mid-build), teed to a log with a BRINGUP_EXIT sentinel; supervised from this peer with durable journal heartbeats every ~12 min. LONG (~20–50 min).
4. Verify BRINGUP_EXIT=0 + fresh-green `autoverify.json` (0 warnings) + all peer origins serving.
5. FOREGROUND live-verify on the fresh seed: gate (b) content-stories sweep → 47/47; gate (a) ORG-CLEAN re-verify (0 source-org tokens); gate (g) interview report renders live + alignment assertion green.

## Escalation conditions
- Any step needing a platform-repo edit → STOP, SEVERITY=blocker (route via sha-pinned demopatch or escalate).
- Bring-up BRINGUP_EXIT≠0 or autoverify not green → diagnose from the peer (never from the VM; M219 loopback-TLS trap); if it's an ENOSPC/host-minefield class, characterize + route.
- Gate (b) sweep < 47/47 with a deterministic (non-flake) residual → root-cause; route forward if it needs new tooling.

## Acceptable close-no-lift outcomes
- If the bring-up lands green but a live sweep surfaces a NEW deterministic defect that needs a tooling fix (not a re-run flake), close-fixed-partial with the defect root-caused + routed — the foundation re-establishment is still the iter's landed deliverable.

## Escalation / handler naming
Route-forward handlers use `GATE-M244-iter{NN}-{short}`.
