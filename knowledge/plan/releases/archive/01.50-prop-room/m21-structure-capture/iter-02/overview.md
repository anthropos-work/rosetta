---
iter: 02
milestone: M21
iteration_type: tik
status: closed-fixed-partial
created: 2026-06-11
---

# M21 iter-02 — tik (first tik under TOK-01)

The first standard iteration under the bootstrap strategy. Job: stand up the **live baseline harness**, confirm the
live baseline empirically, and make the first real move on **stage 3 (structure-apply)** — the `provision.go:108` gap.

## Active strategy reference
**TOK-01** (staged-pipeline build toward the binary serve-anonymously gate; structure artifact keyed by source digest,
applied **before** row replay). This tik operates entirely under TOK-01's `Next-tik direction` for iter-02.

## Re-survey (Phase 1 Step 0)
Re-ran the static measurement: furthest-passing-stage = **2** (build + bootstrap implemented; `provision.go:102-108`
is still the literal `# NOT YET AUTOMATED` placeholder; `stacksnap replay --surface directus` exits 4). TOK-01's named
target (live baseline + structure-apply) is **untouched and still meaningful** — no sibling milestone moved it (M21 is
the first milestone; M22–M25 not started). Live infra confirmed available: Docker 28.5.1, `directus/directus:11.6.1`
(518 MB) cached, `pgvector/pgvector:pg16` + `postgres:16-alpine` present, the directus row cache complete
(`.agentspace/snapshots/directus/6cd35278…/`, 9 tables). Target stands; no substitution.

## Cluster / target identified
The pipeline dies at **stage 3 (structure-apply)**. Every downstream stage (replay-exit-0, boot, serve) is blocked
behind it. So stage 3 is the only meaningful target — advancing it is the whole game. iter-01 established the baseline
**statically** (reading code + the cache); this tik establishes it **live** (Docker) and attacks stage 3.

A crux the static analysis surfaced and this tik must pin empirically: **the digest trap.** `pg.SchemaVersionSQL()`
digests *every column of every table* in the `directus` schema (system `directus_*` tables + content tables). The
cache key `6cd35278…` therefore encodes the **whole prod directus schema**. A freshly bootstrapped Directus (27 system
tables, no content tables) digests differently → can never cache-hit (exit 5), and even after structure-apply the
target digest only converges if the *entire* schema (system + content + types) matches prod's. This tik measures the
real bootstrapped digest and characterizes exactly what convergence requires — load-bearing for the stage-3 design.

## Hypothesis
A bootstrapped Directus + a content-model structure applied via Directus's own `node cli.js schema apply <yaml>`
creates the 9 user-collection tables **and** their `directus_collections`/`fields`/`relations` registry rows in one
step (Directus owns both halves), advancing the pipeline past stage 3. The structure YAML is captured from a
self-contained reference Directus (TOK-01 option (c) — no prod access, doubles as the test fixture).

## Expected lift
furthest-passing-stage **2 → 3** (live structure-apply working: the 9 tables created + registered on a fresh
bootstrapped stack). Stretch: if convergence + replay also fall, 3 → 4.

## Phase plan (staged-pipeline tik)
1. **Live baseline harness** — throwaway Postgres (offset port, isolated) + bootstrap a real `directus/directus:11.6.1`
   on a `directus` schema; confirm the 27 system tables land (stage 2 **live**, not just static).
2. **Live baseline confirm** — `stacksnap replay --surface directus` against the bootstrapped-but-gap schema exits **4**
   (the live ErrEmptySchema baseline; also satisfies the Phase 0d pipeline pre-flight).
3. **Digest-trap characterization** — compute the bootstrapped directus-schema digest; compare to `6cd35278…`; record
   exactly what convergence requires (the stage-3/4 design input).
4. **Mechanism validation** — drive Directus `schema snapshot` / `schema apply` on the live container to prove the
   content-model round-trip mechanism (create-collection → snapshot YAML → apply on a fresh stack → tables+registry).
5. **First structure artifact** — produce the structure artifact for the 9 collections (column lists known from
   `directus.go Surface()`); apply it; re-measure the stage.

## Escalation conditions
- If the **reference-collection build** for all 9 collections proves heavier than one tik (≈250 fields), land the live
  baseline + mechanism-validation + digest characterization, and **route the full 9-collection artifact forward to
  iter-03** as a Fate-3 item under TOK-01 (not a re-scope — the strategy holds; only the per-tik target narrows).
- If Directus's anonymous-serve permission proves to need a **running-instance API call** no pre-staged artifact can
  satisfy (the `overview.md` Re-scope trigger's named risk), record the falsification and surface it — that is a
  strategy-level finding for a future tok, not a this-tik fix.

## Acceptable close-no-lift outcomes
- The live baseline harness + a **documented, empirically-pinned digest-trap characterization** + a **validated
  structure round-trip mechanism**, with the full 9-collection artifact routed to iter-03. The metric may stay at
  "stage 2 (now live-confirmed)" while the stage-3 *path* is de-risked — a first-class characterization deliverable
  under the protocol (the milestone is iterative precisely because this only breaks live).

## Test discipline
`go build ./...` + `go vet ./...` + `go test ./...` for any `stack-snapshot` code touched (the existing suite must stay
green); the live harness is throwaway (torn down at iter close — no persistent stack mutated).
