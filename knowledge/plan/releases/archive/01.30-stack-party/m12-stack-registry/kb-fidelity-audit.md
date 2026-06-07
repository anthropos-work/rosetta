---
title: "KB Fidelity Audit — M12 Unified stack registry + first-available-N allocation"
date: 2026-06-07
scope: milestone:M12
invoked-by: build-milestone
---

## Verdict
YELLOW

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Demo registry (per-demo record of project/offset/profile/clones) | `demo-stack/GUIDE.md`, `demo-stack/README.md` | `demo-stack/rosetta-demo` (`reg_set`/`reg_del`/`cmd_status`); runtime `demo-stack/stacks/registry.json` (gitignored) | PAIRED |
| Port-offset / `!override` engine | `dev-stack/README.md`, `knowledge/README.md` | `stack-core/gen_override.py` | PAIRED |
| Dev stack lifecycle (no registry) | `dev-stack/README.md` | `dev-stack/dev-stack` (`cmd_status` = `docker ps` only; no registry) | PAIRED |
| Unified cross-type registry (dev+demo) | — (rosetta `corpus/ops/rosetta_demo.md` is a pointer; no unified-registry model) | — (does not exist yet) | DOC-ONLY/BLIND — milestone deliverable |
| First-available-N allocator | — | — (does not exist yet) | BLIND — milestone deliverable |

## Fidelity Findings

1. **GUIDE.md overstates current N assignment** — `demo-stack/GUIDE.md` ~L39: "The registry assigns N and records the ports each demo owns."
   - **Expected (per doc):** the registry *assigns* N.
   - **Actual (code):** the caller passes an explicit N (`rosetta-demo up <N>`); `reg_set` only *records* the passed N. There is no allocator — N assignment is manual.
   - **Verdict:** STALE (aspirational). NOT load-bearing for M12's implementer — M12 *builds* the allocator that makes this claim true, so the implementer is not misled into assuming an allocator already exists. The claim resolves to truth at M12 close.
   - **Fix owner:** update code (M12 builds the allocator) → the doc claim becomes accurate. The rosetta-side delivery doc will state the unified-registry + first-available-N model.

2. **Independent dev/demo N-spaces are accurately reflected in code** — `dev-stack/README.md` L19 "DEV_OFFSET (default 10000 — collision-free, matches demo-stack)" describes only the *offset multiplier* matching, not a shared N-space.
   - **Actual (code):** `dev-N` and `demo-N` both compute `base + N*10000`, so `dev-1` and `demo-1` collide on every port. The docs do not claim they are collision-free against each other.
   - **Verdict:** ALIGNED (the docs do not over-promise cross-type isolation). This collision is precisely M12's target.

## Completeness Gaps

1. The demo registry record schema (`{project, offset, profile, services, override, clones}`) is implemented but only loosely described in prose. M12 extends it to a unified cross-type schema `{type, N, ports, status, created}`; the new schema + allocator contract will be documented as part of M12's Phase 5 (rosetta `corpus/ops/rosetta_demo.md` + extensions section docs). Not a blocker — it is the milestone's deliverable.

## Applied Fixes
- None applied pre-build. Finding 1 (GUIDE.md "assigns N") is resolved *by* M12's implementation, not by a pre-build doc edit — fixing the prose now (before the allocator exists) would make the doc claim true-but-unimplemented. The doc + code converge at M12 close (Phase 5).

## Open Items (require user decision)
- None.

## Gate Result
YELLOW — proceed with tracking. No blind area that blocks: the "unified registry" and "first-available-N allocator" topics are DOC-ONLY/BLIND *because they are the milestone's deliverables* (the milestone overview's `Delivers →` already names `corpus/ops/rosetta_demo.md`). The one STALE claim (GUIDE.md "registry assigns N") is non-load-bearing for the implementer and converges to truth at close. Track Finding 1 as KB-1 in `decisions.md`; address in Phase 5.
