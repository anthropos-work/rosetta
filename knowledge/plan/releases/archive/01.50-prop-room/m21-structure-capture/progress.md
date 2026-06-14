# M21 — Progress

**Status:** `archived` (completed 2026-06-13 — closed-on-gate, gate MET at iter-08). **Shape:** iterative (exit gate in `overview.md`).
**Closed by:** `/developer-kit:close-milestone` → merged to `release/01.50-prop-room`.
**Active strategy:** TOK-01 (staged-pipeline) refined by **M21-D7 (option A)** + **M21-D9 (serve recipe)**.
**Furthest pipeline stage passing:** **6 of 6 — DEMONSTRATED end-to-end** (iter-05). **Code-ification: ALL stages
AUTOMATED.** iter-07 automated the SCHEMA (stages 3-4 — `stacksnap` captures structure + auto-provisions a
bootstrapped-gap stack); **iter-08 automated the SERVE rows** (directus_collections registration + public-policy
permissions, behind the new firewall structural-metadata admissibility class — M21-D13; rosetta-extensions). The
exit_gate "stacksnap applies the captured structure → serves a captured sim anonymously over HTTP" is now **MET by
tooling** — no hand SQL. Full capture→apply Go path live-validated; the serve rows are byte-equivalent to iter-05's
demonstrated `serve.sql`.

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

- iter-06 (tik, closed-fixed): **code-ification capture core shipped.** manifest additive `Structure` artifact + a
  dynamic, digest-aligned directus DDL+PK structure generator (`CaptureStructure`) + `pg.QueryRowString`, unit-tested
  + live-validated on prod (26 tables/8 seqs/26 PKs). Found+fixed the privilege-visibility alignment (M21-D10:
  capture must scope to the digest's information_schema view, not pg_class). rosetta-extensions `2c42ed5`.

- iter-07 (tik, closed-fixed): **apply side automated → stages 3-4 by tooling.** `stacksnap` captures the structure
  (`_structure.sql` + `manifest.Structure`) + auto-provisions a bootstrapped-gap stack before the row replay
  (`replayCmd` → `tryAutoProvision`, `pg.ExecScript`); redefined exit semantics (empty→4, gap+no-structure→5,
  diverged→5, gap+structure→0). Robustness: default-reference seq capture (M21-D11). Adversarially reviewed; fixed a
  gap-guard regression (M21-D12). Validated live (capture→bootstrap→auto-provision→10128 rows). rosetta-extensions
  `fe01d64`. See iter-07/progress.md.

- iter-08 (tik, closed-fixed): **GATE MET — serve half automated.** `stacksnap` now captures the
  `directus_collections` registration + public-policy `directus_permissions` read rows (faithful dynamic render) +
  the new firewall **structural-metadata admissibility class** (admit `directus_*` system tables as "structure, not
  tenant data" iff zero tenant-scope columns — extend, never loosen; M21-D13), applied on auto-provision via the
  existing path. A bootstrapped + stacksnap-provisioned stack serves the captured catalog anonymously with NO hand
  SQL. Full capture→apply Go path live-validated (render faithfulness + idempotent no-collision apply + the 6-stmt
  CaptureStructure→ExecScript chain). Serve rows byte-equivalent to iter-05's `serve.sql`. See iter-08/progress.md.

## Next-iter queue (Fate-3) — GATE MET; routed to harden + future milestones
- To `/developer-kit:harden-mstone-iters` (final pass): AP-1 (replayCmd-wiring hermetic test + conn seam), AP-2
  (multi-snapshot tie-break determinism), AP-3 (exit-4-boundary regression guard), user-defined-type/identity guards,
  firewall-ordering direct test, PLUS iter-08: a live-integration harness for the serve-row render SQL (hermetically
  unit-tested + hand-validated live, no automated integration test yet) + a direct assert-then-read ordering test.
- Carried (future milestones): `STRUCT-M21-iter03-artifact`; `directus_files` ref capture (wire the dead `media.go`);
  M23 referential closure of the 20 dangling relations.
