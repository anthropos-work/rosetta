# M202 Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in [`overview.md`](overview.md).

## Section checklist

- [x] **(1) Manifest model + light validator** — the manifest format (Products → Stories → Use Cases →
  Playthroughs) + a validator enforcing **both-way id integrity** + **precondition-coverage** (every declared
  seed/preconditions resolves to a named seeded world), **datadna-gated**. rext `manifest/` + `cmd/ptvalidate`
  (tests -race green, `79df988`).
- [x] **(2) Per-surface locator/landmark page-object layer** — the shared registry every Playthrough imports
  (semantic locators + a landmark registry for ambiguous surfaces), **1 surface to start** (`/profile`); re-pin
  O(surfaces). rext `e2e/lib/{page-object,profile-page}.ts` + `hero-login.ts` (tsc green, `5353396`).
- [x] **(3) Dedicated decoupled seed preset** — test data ≠ demo data (`pt-world`, 2 private orgs); **spans
  entitlement tiers + multi-org-private**; covered by the `datadna` gate. rext `seed/pt-world.seed.yaml` +
  `seed/seed-worlds.yaml` (seeding-validator VALID, `de55b9b`).
- [x] **(4) Reset-to-seed lifecycle + serial-default runner** — per-suite reset via the real `--reset` path
  (additive re-seed FORBIDDEN as a reset), N=0-guarded; `workers: 1`, `fullyParallel: false`. rext
  `e2e/run-playthroughs.sh` + `playwright.config.ts` (`fcf45ad`).
- [x] **(5) 4-state reporting map** — passing / failing / unimplemented / unimplementable-without-platform-edit
  (the last escalates, never edits the platform). rext `report/` + `cmd/ptreport` + `unimplementable.yaml`,
  distinct glyphs (`ed0408a`).
- [x] **(6) One trivial proof Playthrough** — login → /profile → assert hero identity (the foundation smoke
  test). rext `e2e/tests/profile-identity.spec.ts` (`@pt:pt-profile-identity`) — **GREEN on demo-1** (`e77e176`;
  anchor-story layering fix M202-D4).
- [x] **Docs** — `corpus/ops/demo/playthroughs.md` **(NEW)** graduates the spec-draft (IS the M203/M204
  `iteration_protocol_ref`); cross-referenced from the `demo/README.md` index + `coverage-protocol.md` (function
  sibling) + `CLAUDE.md` docs list.

**Status:** `complete` — all 6 sections + the runbook deliverable landed; proof Playthrough GREEN on demo-1.
Tooling + docs only — zero platform-repo edits. rext authoring copy @ `e77e176` (§1–§6, `79df988..e77e176`),
tree clean; the `opening-night-m202` tag + the consumption-clone re-pin happen at CLOSE. Next:
`/developer-kit:harden-milestone` (optional) then `/developer-kit:close-milestone`.

## M202: Hardening

### Pass 1 — 2026-07-01

**Scope manifest (milestone-touched, the `playthroughs` rext section — authored in the
`.agentspace/rosetta-extensions/` authoring copy, NOT the rosetta corpus branch):**
- Go `manifest/`: `manifest.go` · `validator.go` · `seed_worlds.go` (each with a co-located `*_test.go`)
- Go `report/`: `report.go` · `playwright.go` (each with a co-located `*_test.go`)
- Go `cmd/ptvalidate/`: `main.go` · `discover.go` (`main_test.go`)
- Go `cmd/ptreport/`: `main.go` (`main_test.go`)
- TS `e2e/lib/`: `stack-env.ts` · `page-object.ts` · `profile-page.ts` · `hero-login.ts` (**NO unit tests** — the gap)
- shell `e2e/run-playthroughs.sh` (guard-covered by `manifest/runner_safety_test.go` — no untested arm)
- README present at section root (new-unit handbook contract satisfied)

**Coverage delta (milestone-touched Go packages):**
- report: 89.0% → 100.0%
- manifest: 89.2% → 99.4%
- ptvalidate: 77.8% → 97.5%
- ptreport: 65.5% → 94.8%
- (residual: the two trivial `main()` os.Exit wrappers at 0%, plus a JSON-encode-write-error
  arm + a TOCTOU ReadFile-mid-walk arm — fault-injection-only, left rather than box-ticked)

**Tests added (Pass 1, commit `482650f`):**
- report: 15 unit (JSON-parser error/malformed/empty-doc/same-tag-AND/timedOut; firstError
  status-fallback + no-result default; firstLine no-newline; stateGlyph unknown-state;
  truncate short/boundary/ellipsis; Reconcile empty-manifest zero-coverage + empty-id)
- manifest: 13 unit (Load/LoadDir/LoadSeedWorlds read+parse errors; LoadDir subdir/.txt skip +
  .yml accept + file-error propagation; ProvidesTier unknown-world fail-closed; PlaythroughIDs
  dedup/sort; validator empty-world + blank-precondition + Result.Error multi-line + sort
  comparator both arms + empty product id + isSeatKey seat-vs-descriptor)
- ptvalidate: 9 unit (bad-flag/load/discover/seed-worlds errors; single --manifest; datadna
  bin-not-found non-ExitError arm + PASS-with-DSN forwarding; discoverRegistry walk-error; countTODO)
- ptreport: 10 unit (bad-flag/manifest-load/results-parse/json-create/unimpl-load errors;
  no-regressions gate pass+fail; unknown-gate; loadUnimplementable missing-field/parse/valid)
- TS `e2e/lib/stack-env`: 7 unit (offset math; N=0 base ports; explicit-URL override verbatim;
  app-only override; negative + non-integer PT_STACK_N throw) — **the FIRST unit tests on the
  playthroughs `e2e/lib` layer**, mirroring stack-verify's `*.unit.spec.ts` (run under Playwright,
  never navigate → no live stack)

**Bugs fixed inline:** none — the production code was well-behaved on every error/edge/branch arm.

**Flakes stabilized:** none.

**Knowledge backfill:** no KB-worthy findings — every arm probed confirmed existing documented
behavior (the 4-state map, the fail-safe same-tag AND, the datadna subprocess boundary, the
serial-default runner, the offset math); the `playthroughs.md` runbook stays accurate (no fix
changed behavior). The question was asked; nothing new was surfaced to blend.

### Pass 2 — 2026-07-01

**Tests added (Pass 2, commit `39ef562`) — boundary/input fuzzing (spec dimension 5; matches
the sibling stack-seeding/stack-snapshot fuzz precedent):**
- report: `FuzzParsePlaywrightJSON` (arbitrary bytes → clean map or graceful error, never a
  panic; nil map on error; no empty @pt ids — 55K+ execs) + `FuzzExtractPTTag` (total + stable
  under re-extraction — 1.2M execs)
- manifest: `FuzzLoad` (YAML loader degrades loudly; PlaythroughIDs stays sorted+non-empty on any
  parsed manifest — 64K+ execs) + `FuzzValidateNeverPanics` (the validator is total over any
  parsed manifest, with + without Registry/SeedWorlds deps — a panic there would DoS the CI lint)

**Coverage delta:** negligible on statements (the fuzz seed corpora exercise already-covered
lines); the VALUE is the crash-free property, not a line bump. All 4 targets ran crash-free; no
corpus/crash files committed (cached in GOCACHE).

**Bugs fixed inline:** none. **Flakes stabilized:** none. **Knowledge backfill:** none warranted.

### Stop condition

Loop stopped after Pass 2 (well before the 5-pass cap): the Step 2b full scan found nothing new
worth adding (the only residual is trivial `main()` wrappers + fault-injection-only I/O error arms
that would be coverage-box-ticking, not real behavior probing — explicitly declined per the
no-shallow-tests rule); coverage deltas are now negligible (<2%); zero flakes (Go 3/3 shuffled
-race clean + TS 3/3 clean). Section total: **98.5% statements**, **94 Go+TS test/fuzz functions**
on the playthroughs section. No production-code change across the whole harden (zero bugs
surfaced). Tooling only — zero platform edits.
