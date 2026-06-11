# M21 — Decisions

Implementation decisions with rationale, numbered `M21-D1`, `M21-D2`, … . Cross-iter decisions live here; per-iter
detail lives in each `iter-NN/decisions.md`. The strategy-of-record `TOK-NN` entries also live here (the milestone's
strategy-evolution chain).

## TOK-01: staged-pipeline build toward the binary serve-anonymously gate — 2026-06-11

**Tok type:** bootstrap (iter-01)
**Initial strategy:** Treat M21's gate as a **6-stage end-to-end pipeline** (build → bootstrap → structure-apply →
replay-exit-0 → boot → serve-anonymously; full table in `spec-notes.md`) and drive **furthest-passing-stage (0–6)**
as the primary per-tik metric — converting the binary gate into a measurable signal. Each tik validates **live
against Docker** (the directus/directus:11.6.1 image is cached; Directus bootstrap/permission empiricism only breaks
live — the reason this milestone is iterative, not section). The central build is the **capture-side structure
artifact**: the 9 user-collection table DDL + the `directus_collections`/`directus_fields`/`directus_relations`
registry rows (filtered to the public content model) + the `directus_files` ref capture — produced as a **new
artifact keyed by the source schema digest** and applied to a fresh bootstrapped stack **before** the existing row
replay, so the target schema digest converges out of the exit-4/exit-5 "digest trap." Reuse the generic
capture/replay/manifest machinery and add structure **additively** (the `Predicate`-field precedent — no parallel
subsystem), honoring the user's simple/maintainable constraint. All capture stays read-only / public-only / behind
`AssertPublicOnly` (extended to admit structural metadata, never loosened) — the never-touch-prod constraint.
**Rationale:** the gap is empirically hard (fix16 spent +479 lines on Directus provision quirks) and the failure
modes (anonymous serving, registry-row carve-out, digest convergence) only surface live — so a staged pipeline with
a per-stage metric + live validation each tik is the honest shape. The structure-as-additive-artifact choice keeps
the change inside the existing well-tested machinery rather than spawning a second capture path.
**Strategy class:** new-direction (bootstrap).
**Distance-to-gate context:** gate metric = furthest pipeline stage passing; gate = stage 6 (serve a captured sim
anonymously over HTTP). Static baseline today = **stage 2 of 6** (build + bootstrap implemented; structure-apply is
the `provision.go:108` placeholder; replay exits 4; boot/serve blocked behind). The row half of the cache is
complete; only the structure half is missing.
**Next-tik direction:** iter-02 (first tik) — stand up the live baseline harness (throwaway Postgres + bootstrapped
Directus), confirm the live baseline (replay exits 4), resolve the structure-source question (lean: option (c) — a
self-contained reference Directus whose schema is exported via Directus `schema snapshot`, no prod access, doubles
as the test fixture), and produce the first structure artifact for the 9 collections. Target: stage 2 → 3.
