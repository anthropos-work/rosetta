# M236 — Spec notes

## Pre-flight audits — iter-01

**`/developer-kit:audit-kb-fidelity --milestone=M236`: INVOKED (delegated sub-agent), verdict NOT RETURNED
within the iter-01 window.**

Recorded honestly rather than assumed: the audit was dispatched at iter-01 open against the milestone's
load-bearing doc set (`tailscale-serve.md`, `verification.md`, `content-stories-routes.md`,
`session-clone-spec.md`, `content-stories-spec.md`, `cockpit-spec.md`, `coverage-protocol.md`,
`playthroughs.md`, `latency-budget.md`, `ant-academy.md`, `demo-up-defaults.md`, `demopatch-spec.md`,
`safety.md`) and had not reported a verdict when iter-01 closed. **No verdict was inferred.**

**Judgment recorded for audit:** iter-01 is a strategy tok that ships no code, and iter-02's planned scope
(publish the `rosetta-extensions` tags to origin, re-pin `billion`'s `.agentspace/rext.tag`, prune the
host build cache) is **provably insensitive to knowledge-doc fidelity** — it is a git/host operation whose
correctness is established by the tag graph and the pin guard, not by any corpus claim. So the milestone
proceeds into iter-02 with the gate open, and the verdict is consumed before the first
knowledge-doc-dependent iter (Phase L, which triages render failures against the route/spec docs).

If the audit returns RED, it surfaces as a critical decision at that point per Phase 0b.

### Standing finding fed INTO the audit (established live at iter-01, B2)

The v2.5 tooling is unpublished — `origin/main` is the M228 commit, 0 of 13 `playbill-*` tags are on
origin — while `billion` consumes tooling only via a pinned tag from origin. Flagged to the auditor as a
candidate blind area: no corpus doc names **publication** as a step in the authoring-copy → tagged-clone
consumption path. `CLAUDE.md` and `rosetta_demo.md` both describe the path as "built and tested in the
authoring copy and tagged, then consumed per-stack via a pinned-tag clone" — which reads as though tagging
alone makes a tag consumable by a remote host. It does not.

## Thresholds and denominators

- **Primary metric denominator = 31** landable (session × action) pairs. Derived per-session from
  `stack-seeding/presets/content-manifest.json`, not from prose. See `iter-01/decisions.md` D1 — note the
  `has_manager_view` per-session vs per-product trap (a product-level read returns `None` for all 4
  products and silently under-counts to 18).
- **p95 click→ACCESS < 5 s**, ACCESS defined per `corpus/ops/demo/latency-budget.md` (authenticated shell
  rendered + interactive with the hero's identity present). **State the environment with every number** —
  the same defect measured ~6 s on a laptop and ~112 s on the tailnet VM.
- **ai-labs is presence-only** (M231 verdict) — 2 rows must render, 0 landable result pages. Not a gap.

## Environment facts (billion)

- Workspace `/home/devops/panorama/`; stack workspace `stack-demo/`; rext pin SoT at
  `/home/devops/panorama/.agentspace/rext.tag`.
- 7.3 GiB RAM + 15 GiB swap; 193 G disk, 40 G free at iter-01. `docker system df`: **109 GB build cache,
  107.6 GB reclaimable** — prune before any cold UI-tier rebuild.
- The M217 rext pin guard is **FATAL** on mismatch (`demo-stack/ensure-clones.sh`); escape hatch is
  `DEMO_ALLOW_UNPINNED_REXT=1` for deliberate un-tagged authoring work.

## Publishing milestone tags is the established workflow (not an escalation)

origin carries **160 tags**, including the complete `casting-call-m225…m228` milestone series of the
previous release — and `casting-call-m228-hiring-scope-fix` (the tag `billion` currently runs) is itself on
origin. So pushing v2.5's milestone tags is routine release practice, not a novel action; M230–M235 simply
ran offline and never exercised the publish half. All 13 `playbill-*` tags are ancestors of local `main`,
`main` is a clean fast-forward over `origin/main` (`1d97861..60eff14`, 20 commits), and there are **0 tag
name collisions** on origin — so the publish is purely additive.

## Local toolchain gap (non-blocking)

`pytest` is not installed in the workstation's python3.14 — the `stack-core` / `demo-stack` Python suites
cannot be run locally as a pre-publish check. Go side was verified instead: `stack-seeding` builds clean
and **16 packages pass, 0 fail**; `stack-snapshot` / `stack-secrets` / `alignment` all build. The Python
surfaces are exercised on `billion` at bring-up. Routed as an observation, not an iter.
