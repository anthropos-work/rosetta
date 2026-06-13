# Hardening Ledger — M21 Structure capture

The single source of truth for "what's been hardened" in M21. Final-mode passes do a
cumulative-scope sweep across ALL milestone-touched code (the `stack-snapshot` Go
module in the rosetta-extensions authoring copy: `capture`, `cmd/stacksnap`,
`directus`, `firewall`, `manifest`, `pg`). M21 ran 7 tiks + 1 bootstrap-tok with NO
prior harden, so the cumulative footprint is the whole serve+structure code-ification.

## Pass 1 — 2026-06-13 — final

**Iters hardened this pass:** all milestone-touched code (cumulative final sweep — iter-01..iter-08)
**Tiks covered since prior pass:** all iters in milestone (first harden)
**Coverage delta on touched files:**
- directus/structure.go: 94.2% -> 100.0% stmts (package directus 94.2% -> 100.0%)
- cmd/stacksnap (autoprovision.go, main.go exit classifier): 79.5% -> 80.1% stmts
- capture 98.9%, firewall 100.0%, manifest 95.2%, pg 47.0% (unchanged — see note)
**Tests added:**
- iter-08 (serve render) -> directus/structure_harden_test.go: 5 error-path (per-stage query
  errors: ddl/pk/collection-render/permission-render/admissibility-introspect), 3 edge-case
  (collection-only / permission-only serve blobs + serve-after-schema ordering for permission-only),
  1 assert-then-read (admissibility-fail skips render), 1 fuzz (served-collections literal escaping),
  1 provenance (countStatements ignores comment semicolons)
- iter-07 (auto-provision) -> cmd/stacksnap/autoprovision_harden_test.go: AP-2 multi-snapshot
  tie-break determinism (first-listed-structure wins, follows List() order, rows-only skipped),
  list-error + per-manifest-error skip, re-probe-error surface
- iter-07/08 (exit semantics) -> cmd/stacksnap/replay_exit_harden_test.go: AP-3 exit-code
  stable+distinct contract (4 vs 5 split) + replayProbeExit boundary (nil, doubly-wrapped
  ErrEmptySchema, sibling sentinel)
**Bugs surfaced + fixed inline:** none (the M21 code held under error-path + edge-case enumeration;
the firewall structural-metadata class, the assert-then-read ordering, and the auto-provision
gap-guard all behaved as designed under the deepened tests).
**Flakes stabilized:** none observed.
**Cross-iter integration findings:**
- AP-2 (multi-snapshot tie-break, iter-07): `findStructureSnapshot` selects the FIRST
  structure-bearing ref `store.List()` yields. The production `LocalFS.List()` sorts by
  `Ref.Key()` (surface/schemaVersion), so the choice is deterministic — but the `SnapshotStore`
  INTERFACE does not guarantee order, so the determinism is a property of the concrete backend,
  now pinned by a test against an order-controlled fake store. A future S3/cloud backend whose
  `List()` is unordered would change which structure is applied; the test makes that visible.
- The serve-row render (iter-08) + the schema capture (iter-06/07) compose through one
  `CaptureStructure` ordered-assembly path; the per-stage error enumeration now pins that a
  failure in ANY stage (seq/ddl/pk/serve-render/admissibility) surfaces with a stage-named,
  `errors.Is`-unwrappable error — no stage can silently swallow a catalog read failure.
**pg (47%) + cmd/stacksnap residual note:** the uncovered statements are the live-DB connection
wrappers (`pg.Connect/Close/CopyOut/CopyIn/Exec/ExecScript/QueryRow*`) and the `replayCmd` body +
`replayAdapter`/`adapters.go` real-`*pg.Conn` pass-throughs. These are integration-only by nature
(no conn seam on `replayCmd`/`replayAdapter`). Two items exceed the harden-pass inline-fix boundary
and are ROUTED FORWARD (Fate 3), NOT bundled here:
  - HARDEN-M21-AP1-replaycmd-conn-seam: hermetic `replayCmd` wiring test requires refactoring
    `replayCmd` to accept an injectable connector (changes its signature + the `main` dispatch +
    touches the load-bearing replay path) — architectural, >50 lines. Route to a follow-up iter.
  - HARDEN-M21-serve-live-integration: an automated integration harness for the serve-row render
    SQL needs a live directus Postgres (stand up the container); exceeds harden-pass scope. The
    render is hermetically unit-tested (structure_harden_test.go + serve_test.go) + hand-validated
    live per iter-08. Route to the live-integration backlog.
**Knowledge backfill:** none required this pass (the determinism contract is captured in the
cross-iter finding above + the AP-2 test's doc comment; no subsystem doc claimed otherwise).
**Stop condition:** continue-to-next-pass — directus hit 100%; cmd/stacksnap/pg residual is
live-DB-only (route-forward AP-1) — need a second pass to measure the cross-pass coverage delta
and re-scan the firewall/manifest/capture dimensions.
