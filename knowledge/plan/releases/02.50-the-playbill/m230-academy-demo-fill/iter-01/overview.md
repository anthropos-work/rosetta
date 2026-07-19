---
iter: iter-01
milestone: M230
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
created: 2026-07-19
---

# iter-01 — bootstrap tok: choose the fill path (Option C vs B)

**Type:** tok (bootstrap) · authors TOK-01, the milestone's first strategy. Does NOT terminate the call.

## Inputs
- `overview.md` (exit gate; the Option C vs B ladder; In/Out lists)
- `spec-notes.md` (empty — freshly scaffolded; this iter seeds it)
- Protocol: `corpus/ops/verification.md` + `corpus/ops/demo/coverage-protocol.md`
- `roadmap.md` § v2.5 "the playbill" (the release design)
- Phase 0b KB-fidelity audit: **GREEN** (`kb-fidelity-audit.md`)

## The decision
**Option C** — restore an **FS-as-published fallback** on the demo's **own ephemeral ant-academy clone** via a
sha-pinned rext `demopatch`, so the home grid renders the committed catalog through the **real resolver + render
chain** with **NO "Draft" chip**. Chosen over **Option B** (a net-new firewalled academy-content snapshot surface:
prod capture → replay into the demo app DB + wire `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` + compose the academy subgraph
into the demo router).

Full rationale, distance-to-gate, and next-tik direction: **TOK-01** in the milestone-root `decisions.md`.

## Baseline framing
- **Gate metric:** rendered-card count on the academy home grid, employee vantage, via the coverage sweep's
  `ANT_ACADEMY` descriptor — **≥ floor, NO Draft chip, 0 prod-ejects, on a cold /demo-up**.
- **Baseline (a priori, F4 carry):** **0 real cards** — the demo neither sets `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`
  (confirmed: `ant-academy.sh` sets it 0 times) nor holds academy rows → `emptyCatalogView()`. Established by the
  F4 carry + iter-01's code+launcher verification; a separate cold-up baseline sweep of the UNPATCHED academy is
  not run (it would render the known 0 and pay the heaviest op twice). The first sweep measures the POST-fix grid.
- **Infra:** a cold /demo-up is FEASIBLE here — demo-1 injected images built 41h ago + `demo-stack/stacks/demo-1`
  artifacts present; docker up; 205Gi free.

## Escalation conditions (for the tiks that follow)
- A 100%-blocking failure closeable ONLY by a platform-source edit → **re-scope-trigger** (per coverage-protocol
  § Re-scope trigger). Option C is specifically chosen to avoid this.
- A cold /demo-up that cannot COMPLETE in this environment (docker wedge / ENOSPC / cold-cache + no DSN) →
  **user-blocker** with the specific obstacle (per the orchestrator's infra-reality guidance).

## Close
See `progress.md`.
