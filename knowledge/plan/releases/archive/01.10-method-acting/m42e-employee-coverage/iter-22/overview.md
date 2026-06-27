---
iteration_type: tik
iter_shape: acceptance
status: planned
created: 2026-06-25
---

# iter-22 — P8 authoritative fresh-demo-up acceptance + employee semantic gate + manager smoke-sweep

**Active strategy reference:** TOK-10 (persona-believability-by-root-cause, re-scoped gate). P8 is the
authoritative acceptance: prove the whole M42e believability build reproduces from a zero-manual fresh
`demo-up`, then run the employee semantic gate (the M42e exit gate) + a manager smoke-sweep (M42m input).

**Cluster / target identified:** the fresh-demo-up reproduction + the two semantic sweeps. TOK-10 directed P8
as the closing acceptance after P0–P7 landed the seeders + the harness. Re-survey (the live fresh demo-up,
run this iter) confirmed P5 (Sentinel reload) + the directus/taxonomy replays + the stories seed reproduce
zero-manual, but surfaced one reproduction GAP: the sim-embeddings replay (P6) was skipped because
`dev-setdress.sh`'s `build_cli` skip-if-present reused a STALE `stacksnap` binary (pre-P6, no sim-embeddings
surface) that survives `demo-down --purge` in the stack's `bin/` dir.

**Hypothesis:** making `build_cli` ALWAYS rebuild every CLI from the consumed tag (tests opt out via
`DEV_SETDRESS_USE_STUB_BINS=1`) closes the gap — a fresh demo-up then replays sim-embeddings automatically,
`cms.similarities` fills with the 274 public sims, and `/library/ai-simulations` renders → the employee gate
can be measured zero-manual.

**Phase plan:** Phase A sweep (the fresh demo-up + diagnose) → Phase B triage (the stale-binary root cause) →
Phase C fix (build_cli ALWAYS-rebuild, in rext; re-apply set-dress) → Phase D re-measure (employee semantic
gate) → Phase E close. Plus the manager smoke-sweep (calibrate the M42m namespace; measure the manager
residual — do NOT fix manager gaps here).

**Escalation conditions:** a 100%-blocking employee gap that can ONLY be closed by a platform edit →
re-scope-trigger (the zero-edit line). A manager-vantage gap → recorded as M42m input, NOT fixed here.

**Acceptable close outcomes:** the fresh-demo-up reproduction is proven zero-manual (with the build_cli fix),
the employee gate verdict is measured + honestly reported (gateMet true → believability work reproducibly
proven; or the honest residual routed), and the manager residual is measured + reported as M42m input.
