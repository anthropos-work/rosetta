---
milestone: M42e
slug: employee-coverage
version: v1.10 "method acting"
milestone_shape: iterative
exit_gate: "a Playwright sweep, logged in as an employee hero (e.g. Maya Chen via the cockpit), of EVERY reachable demo page asserts BOTH (a) non-empty semantic content in the DOM (real text/rows per section, not just a shell) AND (b) populated screenshots, for 100% of pages, with ZERO pages empty/error AND ZERO nav links escaping the demo platform (every in-app link/nav resolves to a demo-local surface on its offset port). Gate = the sweep reports 0 failing pages + 0 escapes."
iteration_protocol_ref: corpus/ops/demo/coverage-protocol.md
status: planned
created: 2026-06-24
last_updated: 2026-06-24
complexity: large
delivers: corpus/ops/demo/coverage-protocol.md (NEW — the Playwright demo-coverage sweep + triage + fix iteration protocol) + the Playwright coverage harness (a new rext section, demo-only — the first non-Go rext dev/test dependency) + updates across the touched specs (frontend-tier.md / verification.md / seeding-spec.md / snapshot-spec.md) as gaps close
depends_on: M39+M40+M41
spec_ref: .agentspace/profile_gaps.md (live-demo review, 2026-06-24; root-cause workflow w7t4wq2z4)
---

# M42e — Employee 100% demo coverage

## Goal
A hero of the **EMPLOYEE/member vantage** (e.g. Maya Chen) logged into the demo platform sees **100% of the
demo's pages populated** — no empty pages, no error pages, no out-of-demo escapes. Every page the employee can
reach renders real semantic content, and every in-app link/nav stays inside the demo platform on its offset
ports.

## Exit gate (objective, machine-verifiable)
A **Playwright sweep**, logged in as an employee hero (e.g. Maya Chen via the cockpit handshake), of **EVERY
reachable demo page** asserts BOTH:
- **(a)** non-empty **semantic content in the DOM** — real text/rows for each section, not just a shell; AND
- **(b)** populated **screenshots**,

for **100% of pages**, with **ZERO pages empty/error** AND **ZERO nav links escaping the demo platform** —
every in-app link/nav resolves to a **demo-local surface on its offset port** (e.g. left-menu "Studio" → the
local studio-desk, NOT prod `studio.anthropos.work`). An external link is **NOT valid filler** and must be
rewritten in the demo injection/env. **Gate = the sweep reports 0 failing pages + 0 escapes.**

## Why iterative (not section)
The page set and the failure modes — empty section / federation error / out-of-demo link / missing seed — are
**discovered BY the sweep**, not enumerable up front. Each iter measures (run the sweep) → triages the failures
→ fixes in `rosetta-extensions` (stack-seeding seeds / stack-snapshot serve-grants / the demo injection+env
link-rewriting / a corpus doc) → re-sweeps, until the gate is GREEN for an employee hero. A fixed `In:`
checklist would be speculative — the gate is the commitment, the path emerges from each tik's evidence. Build
with `/developer-kit:build-mstone-iters`.

## Iteration protocol
The loop is authored **by this milestone** as [`corpus/ops/demo/coverage-protocol.md`](../../../../corpus/ops/demo/coverage-protocol.md)
— the Playwright demo-coverage sweep + triage + fix loop. M42e **also delivers** the Playwright coverage
harness itself (a new rext section — `stack-coverage`, or under `stack-verify` — demo-only; runs against the
live demo on offset ports). Note: **Playwright is the first non-Go rext dev/test dependency** — sanctioned by
the coverage requirement.

## Approach sketch
Build the harness (a rext section) that:
1. **logs in** as a roster hero via the cockpit handshake,
2. **crawls** the in-app nav as that vantage,
3. **per page** asserts DOM non-emptiness + captures a screenshot + asserts every link host is demo-local,
4. **emits** a coverage report (pages, pass/fail, escapes).

Then iterate the fixes into `rosetta-extensions` until the gate is GREEN for an employee hero. The fix surfaces
per failure mode:
- **empty section / missing seed** → `stack-seeding` (seed the data the page reads),
- **federation / content error** → `stack-snapshot` serve-grants (the page reads replayed content),
- **out-of-demo link** → the demo injection + env **link-rewriting** (rewrite the host to the offset port),
- **a documented gap** → a corpus doc update.

**Re-scope trigger:** if a 100%-blocking gap can **ONLY** be closed by a platform change (the platform repos
are read-only this release), **escalate** (the zero-edit line) and **record it** — do not edit
next-web/app/cms/jobsimulation.

## Depends on / Parallel with
**Depends on:** M39 + M40 + M41 — the targeted fills (profile identity, Directus serve-grant, profile depth)
land first so the coverage loop only chases the **tail** of remaining gaps, not the known-and-already-scoped
ones. **Parallel with:** none.

## Open questions
- Harness home: a **new `stack-coverage`** rext section vs **under `stack-verify`** — decide in the bootstrap
  iter against how much it reuses verify's offset/port plumbing.
- Crawl strategy: pure in-app nav-link discovery vs a seeded route manifest — decide against escape-detection
  fidelity (a manifest can't catch a nav that escapes the demo).
- "Non-empty semantic content" assertion shape: per-section DOM selectors vs a generic text-density floor —
  decide per-iter against the false-pass/false-fail rate the sweep surfaces.
- Playwright wiring: how the first non-Go dev/test dependency is pinned + invoked alongside the Go rext tooling
  (its own dir + lockfile under the rext section).

## KB dependencies
Read as contract:
- [`corpus/ops/demo/frontend-tier.md`](../../../../corpus/ops/demo/frontend-tier.md) — the demo UI tier
  (next-web + studio-desk + ant-academy, offset ports, minted-pk + offset-URL baked) the sweep crawls.
- [`corpus/ops/verification.md`](../../../../corpus/ops/verification.md) — the bring-up auto-verify net (the
  scoped, non-fatal `verify live` on the stack's own offset ports) the coverage harness extends/sits beside.
- [`corpus/ops/rosetta_demo.md`](../../../../corpus/ops/rosetta_demo.md) — the demo-stack lifecycle + offset
  ports + Clerkenstein injection the cockpit login rides.
- [`corpus/ops/demo/stories-spec.md`](../../../../corpus/ops/demo/stories-spec.md) — the verified-skill chain +
  the roster hero (Maya) the sweep logs in as.

## Delivers →
- **NEW** [`corpus/ops/demo/coverage-protocol.md`](../../../../corpus/ops/demo/coverage-protocol.md) — the
  iteration protocol (the Playwright demo-coverage sweep + triage + fix loop).
- The **Playwright coverage harness** — a new rext section (demo-only; the first non-Go rext dev/test
  dependency).
- **Updates** across the touched specs — `corpus/ops/demo/frontend-tier.md`, `corpus/ops/verification.md`,
  `corpus/ops/seeding-spec.md`, `corpus/ops/snapshot-spec.md` — as gaps close.
