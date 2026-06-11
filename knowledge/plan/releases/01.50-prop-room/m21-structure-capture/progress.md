# M21 — Progress

**Status:** in-progress (iter-05 closed). **Shape:** iterative (exit gate in `overview.md`).
**Build with:** `/developer-kit:build-mstone-iters`.
**Active strategy:** TOK-01 (staged-pipeline) refined by **M21-D7 (option A)**.
**Furthest pipeline stage passing:** **6 of 6 — DEMONSTRATED end-to-end** (iter-05): a booted per-stack Directus serves
a captured public simulation anonymously (`GET /items/simulations?limit=1` → 200, real published row). **Gate NOT yet
met-by-tooling:** the structure was hand-applied; the exit_gate's "stacksnap applies the captured structure" clause
needs the code-ification (`STRUCT-M21-codeify`) — the remaining critical-path deliverable.

## Running ledger
_Appended after each iter (tik = a standard iter toward the gate; tok = a strategy/retro iter)._

- iter-01 (tok/bootstrap): authored TOK-01 (staged-pipeline strategy) + the 6-stage metric + static baseline
  (stage 2/6); Phase 0b KB-fidelity YELLOW; infra confirmed runnable (Docker + cached directus image + complete
  row cache) — see iter-01/progress.md
- iter-02 (tik, closed-fixed-partial): live baseline established; **stage 2 secured** (fixed the `.local` admin-email
  bug Directus 11.6.1 rejects — M21-D1); baseline refined to **exit 5** not 4 (M21-D3); digest trap crystallized as
  full-schema-keyed (M21-D5); **structure-apply mechanism validated** = Directus `schema apply` of a snapshot YAML
  creates tables + registry (M21-D2); structure-source finding — pure option (c) can't provide prod types, real
  artifact + source decision routed to iter-03 (M21-D4). furthest-stage stays 2 (live-confirmed). Routes carried
  forward: STRUCT-M21-iter03-source, -iter03-artifact, -digest-keying, + directus_files wiring — see iter-02/progress.md

- iter-03 (tik, closed-fixed-partial): **structure-source blocker RESOLVED** — operator sanctioned a bounded read-only
  prod structural read via the wired `postgres` MCP (M21-D6). Captured the **real faithful structure** for all 9
  collections (exact `pg_catalog` DDL + registry inventory: 9 collections / 217 fields / 43 relations, 20 dangling →
  M23). **Decisive digest finding** (M21-D5 → option B): prod digest `6cd35278…` is over the full 53-table schema
  (27 system + 26 collections); surface captures 9 of 26 → whole-schema digest can never converge → re-key per-surface.
  furthest-stage stays 2 (structure not yet applied). EXIT_REASON user-blocker: the digest-keying fork (A vs B, touches
  shared taxonomy keying) surfaced to the operator. See iter-03/progress.md.

- iter-04 (tik, closed-fixed): **Option A validated end-to-end → stage 2 → 4.** prod's system digest = bootstrapped
  11.6.1 (no version skew); applied the **26-collection** structure (`iter-04/structure.sql`) → digest converged to
  `6cd35278…` (stage 3) → `stacksnap replay` exit 0, 10128 rows, simulations=304 (stage 4). The digest is over column
  structure not registry rows, so stage 4 decoupled from serve. Caveat: apply was hand-applied DDL; stacksnap
  code-ification pending (M21-D8). See iter-04/progress.md.

- iter-05 (tik, closed-fixed): **stages 5–6 DEMONSTRATED → furthest-stage 4 → 6.** Booted Directus + served a captured
  sim anonymously (200, real published row). Root-caused the milestone's live-only risk: the structure artifact needs
  **PRIMARY KEY constraints** (Directus ignores PK-less collections; digest doesn't see PKs). Serve recipe = struct(+PK)
  → register (directus_collections) → public read permission on Directus's hardcoded policy `abf8a154`; directus_fields
  NOT needed (M21-D9). Gate demonstrated, not yet automated. See iter-05/progress.md.

## Next-iter queue (Fate-3, → iter-06 under TOK-01)
- `STRUCT-M21-codeify` — **THE critical path to the gate:** make `stacksnap` capture the structure (26-collection DDL
  **including PKs** + the directus_collections registration rows + the public read permissions) over `--dsn` into the
  snapshot, and APPLY it before the row replay in provision. Flips the gate demonstrated → met. (Proven recipe:
  `iter-04/structure.sql` + `iter-05/pks.sql` + `iter-05/serve.sql`.)
- Carried: `directus_files` ref capture (wire the dead `media.go`) + M23 referential closure (not needed for the basic
  gate, confirmed).
- `STRUCT-M21-iter03-artifact` (carried) + `directus_files` ref capture (wire the dead `media.go`) + M23 referential
  closure of the 20 dangling relations.
