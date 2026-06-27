**Type:** tik (tooling; design-plan P7 — the semantic-coverage harness rebuild)

# iter-21 — P7 semantic-coverage harness rebuild

## What was built (rext `stack-verify/e2e/`)
- **`lib/empty-states.ts`** — the placeholder / empty-state / skeleton / error catalog the gate REJECTS as
  not-real-content (`EMPTY_STATE_PHRASES` + `SKELETON_SELECTORS` + `EMPTY_SELECTORS` + `stripEmptyPhrases`).
- **`lib/coverage-manifest.ts`** — per-page, per-section DESCRIPTORS { region selector, realContent (text/
  count/both), minCount floor, exception+reason }. **region-not-found = FAIL.** TWO namespaces: **employee**
  (Maya, 7 pages, fully CALIBRATED against demo-3) + **manager** (Dan, the M36 Workforce dashboard surfaces —
  verification funnel / teams / role gap+mobility / succession / mobility+feedback; authored, `calibrated:false`
  until the M42m sweep).
- **`lib/section-assert.ts`** — the per-section engine: resolve region (0→FAIL), reject error/skeleton/empty,
  assert real text (mustInclude + meaningful-len after stripping empty-state copy) + cardinality ≥ floor (or
  exception-relaxed). ONE bounded re-assert after an extra settle (slow-paint vs genuinely-empty). `warmStack`
  primes caches before the authoritative sweep.
- **`lib/persona-assert.ts`** — (1) role↔skills coherence: the allow-set is derived from the hero's OWN
  rendered role skill-panel at sweep time (platform-resolved `job_role_skills`) + a junk-pool denylist (the
  flat-pool head bug) — NO hand-list, NO pg dependency (supply-chain stays Playwright-only); (2) avatar
  menu==profile + is-a-real-photo (raster data-URI, not a silhouette SVG); (3) org name + logo.
- **`lib/crawl.ts`** DEMOTED to reachability + escape-classification ONLY (the content verdict moved to the
  manifest via an `onPageAssert` hook). Kept the prod-eject-vs-allowed-external policy (chapter citations +
  LinkedIn-import help = presenter-notes). Added **template-identical SAMPLE rules** (`/sim/<slug>` cap 20,
  `/skill-path/<slug>/chapter` cap 12) + `sampledOut` disclosure + a `frontierRemaining` that excludes
  sampled-out paths (honest `cappedAtFrontier`) — so the ≈300-sim library doesn't explode the frontier while
  the manifest + escape frontier still EXHAUSTS.
- **`lib/review-html.ts`** — emits `coverage-review.html` (per-section verdicts + screenshots +
  documentedExceptions[] + presenterNotes[] + sampledOut[]).
- **`tests/coverage.spec.ts`** rewritten to compose crawl + manifest section-asserts + persona into the gate
  (`0 failingSections + 0 personaFailures + 0 escapes + 0 notReached`, frontier-EXHAUSTED). `COVERAGE_NO_GATE=1`
  for a baseline/diagnostic run. `run-coverage.sh` updated for the new report schema + `COVERAGE_EXPECTED_ORG`.
- Diagnostic scaffolding kept: `tests/calibrate.spec.ts`, `tests/calibrate-avatar.spec.ts`,
  `tests/persona-check.spec.ts` (all env-gated; the re-calibration tools).

## Calibration (one sweep) + floor tuning
The calibration sweep (demo-3, `COVERAGE_NO_GATE=1`) read the real render: `/library/ai-simulations` 20 sim
cards, `/library/skill-paths` 22 path cards, `/home` 4 assignments + 11 saved links, `/profile` 6 skill tags,
etc. Floors set "substantial but achievable" (sim/path grid ≥6, assignments ≥2, saved ≥3, tags ≥4) — calibrated
to pass the seed's actual output, not new false-fails. `/settings` carries the documented terse-menu exception.
The first calibration sweep hit the 150 cap (the ≈300-sim library exploded the frontier) → added the
template-sample rules → the re-sweep **frontier-EXHAUSTED at 62 pages** (`cappedAtFrontier=false`,
`frontierRemaining=0`; sampledOut = 78 sims + 10 chapters disclosed).

## Self-test (demo-3, which has the P0-P6 fixes) — the gate DISCRIMINATES
**PASSES-FIXED** — all **11 manifest sections PASS** + role-skills-coherence PASS: the P6 library (20 sims +
real categories + `searchSimulations=obj`), the 22 skill-paths, the P1 coherent DevOps skills (no junk-pool),
the P3 activity (4 assignments / 10 saved), the P2 profile timeline/skills. The text-density gate would have
green-lit these AND the placeholder pages alike; the semantic gate passes the real ones on their actual content.

**FAILS-OLD** — **2 persona failures** the old gate NEVER caught: (1) avatar-consistency — the menu avatar is
`/svg/avatar-placeholder.svg` (a silhouette) ≠ the /profile avatar (a real jpeg data-URI); (2) org-identity —
the org name renders ("Cervato Systems") but no org logo image is in the header. **Both are a LIVE RE-APPLY
GAP, not a code gap:** demo-3's `roster.json` (exported 03:11, BEFORE iter-16/18) LACKS the `picture`/`org_logo`
fields, so the fake-fapi serves `image_url=""`. The Clerkenstein `resources.go` + the `roster.go` seeder code
(iter-16/18) correctly thread `Picture→userRes.image_url` + `OrgLogo→orgRes.image_url` — a FRESH demo-up
re-exports the roster with the current code → the menu avatar + org glyph thread clean. So this residual closes
on the P8 fresh-demo-up; NOT hand-patched into live demo-3 (against the reproducibility rule).

## EMPLOYEE residual under the NEW semantic gate
`reachable=62 (frontier EXHAUSTED), failingSections=0, escapes=0, notReached=0, personaFailures=2` — the 2
persona failures (avatar menu-thread + org-logo header-render) are a **roster re-apply gap** that the P8 fresh
demo-up closes. Sections + cardinality + escapes are GREEN; persona is the residual.

## Close — 2026-06-25

**Outcome:** the manifest-driven semantic coverage harness is BUILT (7 new/rewritten files) + CALIBRATED (one
sweep, frontier-exhausted 62 pages) + SELF-TESTED (passes-fixed: 11/11 sections + role coherence; fails-old: 2
persona failures the textLen gate missed). Employee residual = 2 persona failures (a live roster re-apply gap,
not a code gap → P8). coverage-protocol.md updated with the manifest model + exception table.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (P7 of P0-P8; the authoritative gate is P8's fresh-demo-up acceptance + the manager sweep —
this iter does NOT claim gate-met; the employee residual is honestly 2 persona re-apply failures)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (manifest namespaces: employee calibrated / manager authored), D2 (persona allow-set from the
rendered role panel, no pg dep), D3 (template-identical sample rules), D4 (the 2 persona fails are a live
re-apply gap, routed to P8 not hand-patched).
**Side-deliverables:** none.
**Routes carried forward:** P8 (fresh-demo-up zero-manual acceptance + manager sweep) → a later run — the
authoritative gate; it closes the 2 persona re-apply failures (the fresh roster export threads the avatar +
org logo). The manager manifest namespace (`calibrated:false`) is calibrated by the M42m manager sweep at P8.
**Lessons:** A snapshot/seed SURFACE existing in code ≠ it being reflected on a LIVE stack — the avatar/org-logo
threading is correct in `resources.go`+`roster.go` but demo-3's stale roster.json (pre-iter16/18) doesn't carry
the new fields; only a re-export (fresh demo-up, or a live roster re-export + fapi restart) applies it. The
semantic gate's persona checks are exactly what surfaces such a live-staleness gap (the old text gate couldn't).
Documented in coverage-protocol.md (the manifest model + the re-scoped gate).
