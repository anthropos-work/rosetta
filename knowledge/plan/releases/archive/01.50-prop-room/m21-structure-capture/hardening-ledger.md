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

## Pass 2 — 2026-06-13 — final

**Iters hardened this pass:** all milestone-touched code (cumulative — re-scan of firewall/manifest/capture dimensions)
**Tiks covered since prior pass:** continuation of the single final sweep (pass 1 + pass 2 = one session)
**Coverage delta on touched files:**
- manifest/manifest.go: 95.2% -> 98.4% stmts (the M21 Structure-artifact Validate branches)
- firewall/firewall.go: 100.0% -> 100.0% (deepened MEANING, not the number — identity/ordering guards)
- capture/capture.go: 98.9% -> 98.9% (the CapturesStructure capability-mismatch branch now exercised;
  its 3 stmts round below the 0.1% display threshold but the branch is now covered)
- directus 100%, cmd/stacksnap 80.1%, pg 47.0% (unchanged — reachable surface saturated; residual is route-forward/live-DB)
**Tests added:**
- iter-06/07 (manifest Structure artifact) -> manifest/manifest_harden_test.go: 3 rejection cases
  (empty checksum, negative statements, + the empty-payload table case), zero-statements-OK edge,
  structure-bearing round-trip-revalidates
- iter-08 (firewall structural-metadata) -> firewall/structural_metadata_harden_test.go: identity-column
  guards (user_created/user_updated/owner/user each rejected + named), real serve-table column sets
  admitted (positive guard), extend-never-loosen ordering (the carve-out can't launder a row-side
  AssertPlan violation — predicate-agnostic), whole-column-not-substring match semantics
- iter-06 (capture orchestration) -> capture/capture_structure_harden_test.go: the CapturesStructure
  capability-mismatch branch (surface declares structure but capturer can't produce one -> fail loud)
**Bugs surfaced + fixed inline:** none. The firewall carve-out, the manifest validation, and the
capture capability dispatch all held. One DUPLICATE test I drafted (a structure-capture-error
re-assertion) was removed pre-commit — the branch was already pinned by
capture_structure_test.go:TestRun_StructureCaptureError_NothingWritten (no double-assertion landed).
**Flakes stabilized:** none. Flake gate: 3 consecutive clean runs of all 21 new tests — green.
**Cross-iter integration findings:**
- The structural-metadata carve-out (iter-08) is now PROVEN independent of the row-side gates
  (iter-02..05 AssertPlan/AssertCaptured): admitting the directus_* registry tables as structure does
  NOT widen or narrow the user-collection ROW firewall — the two gates compose without interference,
  so the M21 "extend, never loosen" claim is regression-pinned at the firewall level (not just the
  directus-adapter level it was pinned at before).
- The manifest Structure artifact (iter-06's additive field) survives Marshal -> Parse -> re-Validate
  intact, integrating the M21 field with the pre-M21 serialization + the Validate gate in one path.
**Knowledge backfill:** none required — the carve-out's extend-never-loosen + predicate-agnostic
properties are documented in firewall.go's package/type doc comments already; the new tests pin them
rather than revealing an undocumented truth.
**Stop condition:** stabilized — every inline-reachable statement across the 6 M21 packages is
covered (directus/firewall 100%, manifest 98.4% [only the unreachable json.MarshalIndent-error
defensive branch left], capture 98.9% [remaining branches are non-M21], cmd/stacksnap 80.1% + pg 47%
[residual is exclusively the live-DB connection wrappers + the replayCmd/replayAdapter real-conn
pass-throughs = the route-forward AP-1 conn-seam item]). The Phase-2 dimension scan found nothing
new fixable inline; the cross-pass delta on the reachable surface is < 2% (saturated). A pass-3 has
no inline-reachable target. Final harden CLOSED for M21.

## Routed forward (Fate 3) — exceed the harden inline-fix boundary
- **HARDEN-M21-AP1-replaycmd-conn-seam**: a hermetic `replayCmd` wiring test needs `replayCmd`
  refactored to accept an injectable connector (changes its signature + the `main` dispatch + touches
  the load-bearing replay path) — architectural, >50 lines. The same seam unlocks the `replayAdapter`
  (ClearForReplay/CopyIn/ReindexVector) + `adapters.go` real-conn pass-throughs (all 0% — live-DB only).
  Route to a follow-up iter (a `replayCmd`-seam build iter), NOT bundled in this harden pass.
- **HARDEN-M21-serve-live-integration**: an automated integration harness for the serve-row render
  SQL needs a live directus Postgres (stand the container up). The render is hermetically unit-tested
  (serve_test.go + structure_harden_test.go) + hand-validated live per iter-08. Route to the
  live-integration backlog.

These are recorded here (not in decisions.md as HARDEN-CAP-ACCEPTED) because the pass STABILIZED — the
route-forwards are scope-boundary routing of two integration/architectural items, not an accepted
coverage cap. close-milestone's deferral audit picks them up from this ledger.
