---
iter: 01
milestone: M21
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
created: 2026-06-11
---

# M21 iter-01 — bootstrap tok

The milestone's unconditional iter-01 bootstrap tok. Job: author the **first** strategy (TOK-01) for closing the
M10 collection-schema gap, from the milestone `overview.md` + `spec-notes.md` + the protocol doc + a static
baseline. No prior iters; no stalled strategy to revise.

## Inputs
- `overview.md` — the exit gate (replay exits 0 + a booted Directus serves a captured sim anonymously over HTTP),
  In/Out scope, the structure-capture deliverables.
- `spec-notes.md` — the Phase 0b KB-fidelity verdict (YELLOW) + the 6-stage baseline (furthest stage today = 2) +
  the structure-source question.
- Protocol doc `corpus/architecture/alignment_testing.md` — the fidelity-measurement discipline; its
  **snapshot-fidelity dimension** (M9a, `datadna measure-snapshot` with the `snapshot-*` operators) is the relevant
  faithfulness frame once replay works. But the protocol is a *measurement* discipline, not a per-iter build
  protocol with Phase A–E shapes — so the bootstrap tok authors the concrete per-iter shape below.

## Initial strategy (→ TOK-01)
**Staged-pipeline build toward the binary serve-anonymously gate**, validated live against Docker each tik (Directus
empiricism only breaks live — the whole reason this milestone is iterative). The gate decomposes into the 6-stage
pipeline in `spec-notes.md`; the primary metric is **furthest stage passing (0–6)**, which turns a binary gate into
a measurable per-tik signal. Build the **capture-side structure artifact** (the 9 user-collection DDL + the
`directus_collections`/`fields`/`relations` registry rows + the `directus_files` ref capture) as a *new artifact
keyed by the source schema digest*, applied **before** the existing row replay so the target digest converges out of
the exit-4/exit-5 trap. Reuse the generic capture/replay/manifest machinery; add the structure artifact additively
(the `Predicate`-field precedent) rather than a parallel subsystem (the maintainability constraint).

## Strategy class
`new-direction` (bootstrap — no prior strategy to compare against).

## Distance-to-gate context
Furthest pipeline stage passing today (static): **2 of 6**. The pipeline dies at stage 3 (structure-apply,
`provision.go:108` placeholder); stages 4–6 are blocked behind it. The row cache (stage-4 input) is real + complete;
only the structure half is missing.

## Next-tik direction (iter-02, the first tik under TOK-01)
Stand up the **live baseline harness** — a throwaway pgvector/Postgres + a bootstrapped `directus/directus:11.6.1`
on the `directus` schema — and (1) confirm the live baseline (replay exits 4 against the empty schema), (2) resolve
the **structure-source question** (lean: option (c) — build a self-contained reference Directus, capture its schema
via Directus's `schema snapshot` YAML, since it needs no prod access and doubles as the test fixture), and (3)
produce the **first structure artifact** for the 9 collections. Target: advance furthest-passing-stage 2 → 3.

## Escalation / re-scope note
Per `overview.md` Re-scope trigger: if 5 consecutive toks fail to find a path to anonymous serving (e.g. Directus's
permission model needs a running-instance API call no pre-staged artifact can satisfy), escalate to user-strategic-
replan — the interim options (auto-heal / full-taxonomy capture) may become the deliverable instead of the full close.
