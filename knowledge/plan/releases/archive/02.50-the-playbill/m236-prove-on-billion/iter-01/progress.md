# iter-01 — bootstrap tok

**Type:** tok (bootstrap)

Authors the first strategy for M236 from `overview.md` + M235's `carry-forward.md` + the protocol docs
(`corpus/ops/verification.md` + `corpus/ops/demo/tailscale-serve.md`). No prior iters exist; there is no
stalled strategy to review.

## Baseline measurement (taken this iter)

The protocol calls for a baseline in the bootstrap phase. Measured live, 2026-07-20.

### B1 — `billion` is reachable and already carries a demo

- `ssh root@billion` / `ssh marco@billion` both succeed (BatchMode, no prompt). Tailscale peer
  `100.110.136.3 billion tagged-devices linux active; direct 88.99.101.196:41641`.
- Workspace: **`/home/devops/panorama/`** (not a rosetta corpus clone — a bare stack workspace).
  `stack-demo/` holds the full platform clone set + `rosetta-extensions/`.
- A **`demo-1` stack is UP** — 17 containers, `Up 2 days` (left from the v2.4 M228 live proof).
  Cockpit `http://localhost:17700/` → **200**; next-web `:13000` → **307**.
- Host: 7.3 GiB RAM + 15 GiB swap, 193 G disk with **40 G free (80 % used)**.

### B2 — the load-bearing blocker: **the whole v2.5 tooling is unpublished**

`billion` consumes `rosetta-extensions` at a pinned tag read from `.agentspace/rext.tag`
(`demo-stack/ensure-clones.sh:57-100`; the pin guard is **FATAL** on mismatch since M217).

| | value |
|---|---|
| `billion:/home/devops/panorama/.agentspace/rext.tag` | `casting-call-m228-hiring-scope-fix` |
| `origin/main` (GitHub `anthropos-work/rosetta-extensions`) | `1d97861` — **the same M228 commit** |
| local authoring copy `HEAD` | `60eff14` = tag `playbill-m235-hardened` |
| commits local-ahead-of-origin | **20** |
| `playbill-*` tags on origin | **0 of 13** |

**Every v2.5 milestone's tooling — M230 (academy fs-published patch), M232 (`ContentStorySeeder` +
sourcing + modality substrate), M233 (`content-manifest.json` projection), M234 (the cockpit
Content-stories tab), M235 (the 13-session fixture matrix + all 3 non-sim sections) — exists ONLY in
the local authoring copy at `.agentspace/rosetta-extensions/`. None of it is pushed.** M230–M235 were
all offline/local milestones, so the tags were tagged but never published.

Confirmed directly against the live cockpit on `billion`:

```
curl http://localhost:17700/content-manifest.json   → 404
curl http://localhost:17700/cockpit-manifest.json   → 404
cockpit HTML (22 400 B): <title>Presenter Cockpit — demo-1</title>, no Content-stories tab markup
```

Push access is available — `git push --dry-run origin main` reports a clean fast-forward
`1d97861..60eff14` (non-mutating capability probe).

### B3 — the fixture ships in-tag (no out-of-band transfer needed)

`stack-seeding/contentsession/fixture/` is **git-tracked**: `content-sessions.yaml` + 13
`content/<key>.json` files (one per simulation session), plus `sourcing.go` / `content.go` /
`contentsession.go` and their tests. `git check-ignore` reports no hits. So a tag checkout on `billion`
carries the whole anonymized content corpus — nothing needs to be copied out of band, and no prod read
happens on `billion`.

### B4 — the gate denominator

Computed from the canonical `stack-seeding/presets/content-manifest.json` (the M233 honesty-gated
projection). `has_manager_view` is a **per-session** field, not per-product.

| product | sessions | landable actions |
|---|---:|---:|
| `simulation` | 13 | 26 (player + manager each) |
| `skill-path-legacy` | 2 | 4 (player + manager each) |
| `ai-labs` | 2 | 0 — presence-only by M231 verdict |
| `skill-path-new` (academy) | 1 | 1 (player only) |
| **TOTAL** | **18** | **31** |

Player seats are `content-player-23 … content-player-27`; the manager seat is `dan-manager`.

**Primary metric: `landed-NON-EMPTY / 31` (session × action) pairs, live on `billion`, on a cold
reset-to-seed.** Secondary gate components: the 2 ai-labs presence rows render; the ant-academy grid
renders real cards (Thread A); p95 click→ACCESS < 5 s per `latency-budget.md`; 0 platform-repo edits;
tailnet-scoped reach only.

**Baseline: 0 / 31.** Not "failing" — *unreachable*: the tooling that seeds and links those 31 actions
is not on the host.

### B5 — host capacity note (routed, not blocking)

`docker system df` on `billion`: **109 GB of build cache, 107.6 GB reclaimable**, against 40 G free.
A cold `/demo-up` rebuild of the UI tier needs headroom; a `docker builder prune` is a cheap
prerequisite for the first cold run. Recorded as a first-tik pre-step, not a separate iter.

## Close — 2026-07-20

**Outcome:** TOK-01 authored — "publish-then-prove": land the v2.5 tooling on `billion` at a real
published tag, then drive the 31 (session × action) pairs live, triaging by product arm.
**Type:** tok (bootstrap)
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (bootstrap, does NOT terminate) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (0 tiks) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (baseline denominator = 31), D2 (publish-first ordering), D3 (host-capacity prune folded into tik-1)
**Side-deliverables:** none
**Routes carried forward:** none new — M235's 3 carry-forward clusters are already in `overview.md`'s `In:` list and are sequenced by TOK-01's phase ordering.
**Lessons:** A milestone whose gate is "prove X live on host H" must, in its bootstrap baseline, verify that
H can *obtain* the artifact under test. Here the artifact was 20 unpushed commits away from the host, and
nothing in the milestone plan named that step — the plan assumed publication had happened as a side effect
of tagging. **Tagging is not publishing.** Recorded to the protocol doc (`corpus/ops/verification.md`) as a
pre-flight rung.
