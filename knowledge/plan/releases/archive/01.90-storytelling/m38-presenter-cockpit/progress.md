# M38 — Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.

- [x] **The panel** — standalone served surface (rext `demo-stack`, offset port `7700+N·10000`), lists stories → hero trios, reads the cockpit manifest projected from the same `stack.stories.yaml` (`cockpit.py`, stdlib-only HTTP server)
- [x] **[Login as]** — wired to M37's active-user selection: a FAPI handshake redirect `?__clerk_identity=<key>` that switches the active seat to the chosen hero
- [x] **[Jump to section]** — the same handshake redirect with `redirect_url=<jump_to>`, so [Login as]+[Jump] land logged-in on the hero's deep-link in one move
- [x] **Deep-link catalog (O9)** — `DeepLinkCatalog()` enumerates next-web routes per vantage (profile/spotlight/growth/take-a-sim for end-users; the Workforce dashboard tabs + mobility/talent-pool for managers)
- [x] **Launch wiring** — `DEMO_STORIES=1` exports the roster → `FAKE_FAPI_ROSTER` (multi-identity fake-fapi), seeds the stories preset, serves the cockpit on the offset port; torn down with the stack (pidfile reaped by `rosetta-demo down`)
- [x] **Roster-export producer** (the M37 integration seam) — `stackseed --roster-export` derives the seeded heroes' exact clerk ids, single-sourced from the seeder's own derivation; consumed by Clerkenstein's registry
- [x] **Docs** — the cockpit section of `stories-spec.md` + the up→present flow in `demo/README.md`
- [x] **Tests** — demo-stack suite 157 green · stack-injection 115 green · stack-seeding green (`-race`); zero platform-repo edits; the 5 Clerkenstein alignment gates + stack-seeding suite stay green

_Last updated: 2026-06-23 (all sections shipped). rext tag `storytelling-m38` to be cut at close. Code:
`rosetta-extensions` @ `ce2b829` (stack-seeding roster/cockpit producers + demo-stack cockpit panel + launch
wiring). Docs on rosetta `m38/presenter-cockpit` @ `007378b`._

## M38: Hardening

### Pass 1 — 2026-06-23
**Scope manifest (milestone-touched code, rext `52c1be0..ce2b829`):**
- `stack-seeding/seeders/roster.go` (+ `roster_test.go`) — the roster-export producer
- `stack-seeding/seeders/cockpit.go` (+ `cockpit_test.go`) — the cockpit-manifest exporter + O9 catalog
- `stack-seeding/cmd/stackseed/main.go` (+ `main_test.go`) — the `--roster-export` / `--cockpit-export` CLI
- `demo-stack/cockpit.py` (+ `tests/test_cockpit.py`) — the stdlib served panel
- `demo-stack/up-injected.sh` (+ `tests/test_tooling.py::StorytellingCockpitWiring`) — launch wiring
- `stack-injection/gen_injected_override.py` (+ `tests/test_injection.py`) — the `--roster` fake-fapi mount

**Coverage delta (milestone-touched Go funcs):**
- `roster.go`: BuildRoster 92.3% → 100% · WriteRoster 83.3% → 100%
- `cockpit.go`: storyAnnotation 66.7% → 100% · WriteCockpitManifest 88.9% → 100% · BuildCockpitManifest 93.8% → 100% (Pass 2)
- `cmd/stackseed`: doRosterExport 75% → 95.8% · doCockpitExport 75% → 95.8% (Pass 2)
- `cockpit.py`: every genuine branch covered (the trace-tool "uncovered" handler body + `main` serve
  loop run in the `ThreadingHTTPServer` worker thread, which stdlib `trace` can't instrument — proven
  exercised end-to-end by `TestServedPanel` + the new `TestMainServes` over a real socket).

**Tests added:**
- `roster_test.go`: 2 (legacy Size<=0 story-skip; WriteRoster encode-error propagation)
- `cockpit_test.go`: 4 (unknown-jump_to generic label; storyAnnotation legacy + no-match; WriteCockpitManifest
  encode-error; manifest nameless-hero skip + roster-lockstep)
- `main_test.go`: 9 (unwritable --roster-out / --cockpit-out create-error; roster + cockpit n==0 warn;
  --stack override re-namespaces ids ×2; load-fail; validate-fail ×2)
- `test_cockpit.py`: 8 (STRUGGLING + empty-vantage _badges branches; struggling-hero render; main() real
  serve path bind→probe→shutdown→0; malformed-JSON manifest; redirect_url full percent-encoding invariant;
  non-ascii key escape)
- `test_injection.py`: 2 (D4 byte-identical default fake-fapi block; --roster main() flag threading)

**Bugs fixed inline:** none — the producers/cockpit/wiring were correct; every gap was a missing test, not a
missing fix. (No shallow tests: each probes real behavior — encode/create error propagation, escaping
correctness, the keys-match-roster invariant, the D4 byte-identical baseline.)

**Flakes stabilized:** none observed (3 clean sequential runs of the new tests; the `TestMainServes` thread
test joins on a clean shutdown, no port/temp leak).

**Knowledge backfill:** none warranted — every invariant the new tests pin (the roster single-source, the
keys-match-roster lockstep, the D4 byte-identical-off fallback, the fail-loud-roster vs non-fatal-cockpit
split, the O9 raw-path fallback) is already documented in `decisions.md` (D1–D7, O9) +
`corpus/ops/demo/stories-spec.md`. The harden pass deepened the *tests* of already-documented behavior, not
the behavior itself. (Question asked, answer recorded — no doc edit needed.)

**Decision recorded:** M38-D7 — the employee-hero `org_role=admin` finding routed to v1.9 close-review
(Fate 3): it's an M35 role-assignment-seam change (the claim is single-sourced from the seeded membership, so
a fix must move all three writes in lockstep), not M38-touched code; bounded but out-of-scope-for-harden.

### Stop condition
Loop stopped after 3 passes: the full 6-dimension scan found nothing new worth adding, the coverage deltas on
the milestone-touched code reached ~100% (only the effectively-unreachable CLI `WriteRoster`-returns-error line
+ the trace-thread-artifact handler body remain, neither a real gap), and no flaky tests remain (3 clean runs).
All regression suites green: stack-seeding `-race`, Clerkenstein 5 alignment gates `-race`, demo-stack
(163 tests), stack-injection (117 tests). Zero platform-repo edits.

## M38: Final Review

Consolidated from the close review (deferral re-audit + parallel code-quality / adversarial / docs / test scans).

### Scope
- [x] M38-D7 → **LAND-NOW (Fate 1)**: vantage-faithful hero `org_role` at the M35 seam (`roleForHero` single
  source + both call-sites in lockstep + regression test). Recorded as M38-D8. (deferral-audit GREEN)

### Code Quality
- [x] [must-fix] `roster.go:93` `BuildRoster` calls the OLD `roleForIndex` directly while `users.go` now uses
  `roleForHero` → the three-write lockstep is broken (a manager hero would export `org_role=member` but seed as
  `admin`). Switch to `roleForHero(idx, st.Size, st.RoleMix, &h)`. (the core of M38-D8)
- [x] [must-fix] `roster_test.go:74` asserts `wantRole := roleForIndex(...)` → masks the regression. Switch the
  reconstruction to `roleForHero(idx, st.Size, st.RoleMix, &h)`.
- [x] [should-fix] `gofmt` flags `seeders/cockpit_test.go` (multibyte `⇒` comment alignment). Run `gofmt -w`.
- [x] [nice-to-have] `cockpit.py` render: defensive skip of a hero with an empty `key` (the Go projection
  already guarantees non-empty keys, but a hand-tampered manifest shouldn't emit a no-op `__clerk_identity=`
  handshake link). Add the skip + a test.

### Documentation
- [x] `corpus/ops/demo/stories-spec.md` — document the FAPI port offset (`5400 + N·10000`) in the handshake-URL
  example so `<fapi-host>` is concrete, and cross-reference the `DEMO_NO_COCKPIT` escape hatch from the cockpit
  section. Note the vantage-faithful `org_role` (M38-D8) in the roster id-contract.
- [x] `corpus/ops/demo/README.md` — note the cockpit-serve is non-fatal (D6) + cross-ref `DEMO_NO_COCKPIT`.
- [x] `CLAUDE.md` — N/A re-check: CLAUDE.md doesn't enumerate the cockpit at a line that's now wrong (the M38
  description in roadmap/state owns it); no edit needed beyond the existing demo-family pointers. (verified)

### Tests & Benchmarks
- [x] Regression test: the three writes agree per hero (membership-row role == casbin-grant role == roster
  `org_role`) AND a manager hero → `admin`, an end-user hero → `member`. (new test in `roster_test.go` /
  `users` seam)
- [x] Full per-stack suites re-run post-fix (Go `-race`, demo-stack, stack-injection) + flake gate 5×.

### Decision Triage
- [x] M38-D8 (vantage-faithful `org_role` / `roleForHero` single source) → blend into
  `corpus/ops/demo/stories-spec.md` § the roster id-contract (a sentence; full record stays in `decisions.md`).
- [x] M38-D1..D6 → already blended into `stories-spec.md` during build (verified, tags present); archive the
  options-considered detail in `decisions.md`.
- [x] M38-D7 → archive (superseded by M38-D8's LAND-NOW; the route-to-close-review record stays for the trail).
