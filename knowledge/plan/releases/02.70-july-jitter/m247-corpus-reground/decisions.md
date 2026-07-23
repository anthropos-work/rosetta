# M247 — Decisions

_(Implementation decisions with rationale, D-numbered, recorded during build.)_

## Pre-flight KB-fidelity findings (Phase 0b — verdict YELLOW)

Report: `kb-fidelity-audit.md`. YELLOW — no blind area unpromoted, no stale claim read as truth (the
implementation reads current `stack-demo/app` source, not stale corpus prose).

- **KB-1** — skillpath-is-live claim STALE across ~30 files (M246 D-01). Ground truth: 0 skillpath in
  repos.yml/compose, 3 subgraphs, runtime state in `public.skill_path_sessions`. Milestone deliverable (§1/§2).
- **KB-2** — literal "4 subgraphs" in 7 files; true count 3 (verified `supergraph-config-{compose,prod}.yaml`).
  Milestone deliverable (§2).
- **KB-3** — `skiller.md` redirect pattern ALIGNED; clean exemplar for the skillpath redirect.
- **KB-4** — `TEMPLATE.md` fact-sheet shape ALIGNED; the 4 new sheets follow it.
- **KB-5** — 4 net-new domains BLIND-AREA but already promoted to `overview.md §Delivers`; source confirmed present.
- **KB-6** — roadrunner "ORPHANED" flag STALE-toward-alive: roadrunner is IN repos.yml (10 repos) + compose
  (`roadrunner:` block, profile graphql) → alive. Section 8 resolves to the negative (confirm live, retire
  ORPHANED framing).
- **KB-7** — `ai-readiness.md` PAIRED; refresh reconciles the aireadiness-package refactor.

## D0 — rext-file drift is OUT of M247's doc-only scope (Fate-2/Fate-3 routing)

The M246 drift ledger proposes some rext-file edits to M247, but **M247 is DOC-ONLY (no rext, per the milestone
charter)**. These stay out and route to already-planned / better-fit siblings:

- **D-02** `gen_injected_override.py` dormant `skillpath` key ("4 injected") — Fate-3 → rext-hygiene / M251.
- **D-03** `test_injection.py` skillpath-injected pins (+ residual skiller `_cfg`) — Fate-2 → M251 (test-health)
  or rext-hygiene (behavioural-test redesign; not needed for green).
- **D-04** `exposure_claim_guard.py:124` skillpath:8095 fixture — Fate-2 → M251 (update with D-03).
- **D-06** `up-injected.sh:458` historical audit-prose comment ("…skiller, skillpath…") — cosmetic rext prose;
  the ledger proposed M247, but it is a rext file → **cannot** land in doc-only M247. Fate-3 → M251/rext-hygiene.

Recorded here (not a new deferral ledger row) — the work has a real home in the release plan. Flagged as a
**Fate-2 release-close reconcile item** for the ops/demo spec docs the CODE milestones own (see below).

## D-fate-2 — ops/demo spec-doc reconcile deferred to release close

Per the M247 charter, the `corpus/ops/demo/` spec docs owned by the CODE milestones
(`content-stories-spec.md`, `content-stories-routes.md`, `demopatch-spec.md`, `cockpit-spec.md`,
`latency-budget.md`, `secrets-spec.md`, and the studio-desk parts of `frontend-tier.md`/`studio-desk.md`) are
**left untouched by M247** — their not-yet-written deltas belong to M248/M249/M250/M252/M253. The
cross-milestone consistency pass over those docs is a **Fate-2 release-close item** (M247-reconcile → final
release-close pass).
