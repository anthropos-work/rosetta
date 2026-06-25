# M42e Spec Notes

Technical notes accumulate here during the iteration loop. The authoritative spec is
[`.agentspace/profile_gaps.md`](../../../../.agentspace/profile_gaps.md) (live-demo review, 2026-06-24;
root-cause workflow w7t4wq2z4). The iteration protocol this milestone authors is
[`corpus/ops/demo/coverage-protocol.md`](../../../../corpus/ops/demo/coverage-protocol.md).

## Pre-flight audits — iter-01
- **`/developer-kit:audit-kb-fidelity --milestone=M42e`** → **GREEN** (2026-06-25). Report:
  [`kb-fidelity-audit.md`](kb-fidelity-audit.md). 5 fidelity findings all ALIGNED, 0 blind areas. Two
  DOC-ONLY items — `coverage-protocol.md` + the link-rewriting fix surface — are both declared M42e
  deliverables, not blind areas. The bootstrap tok authored its strategy against verified docs.

## The Playwright coverage harness — RESOLVED: extends the EXISTING `stack-verify/e2e/` harness
**Finding (iter-01):** Playwright is NOT a from-scratch new dependency — `stack-verify/e2e/` already exists
with `package.json` (`@playwright/test ^1.49.0`), `playwright.config.ts`, and `tests/smoke.spec.ts` (an
unauthenticated login-page smoke test). The M42e coverage harness is the **authenticated, multi-page,
semantic-content + escape-detection** evolution of this. **Harness home: under `stack-verify`** (reuses the
offset/project/scope plumbing in `lib/target.sh` + the e2e dir's Playwright pin). The coverage spec/config is
a sibling under `stack-verify/e2e/` (or a `coverage/` peer), keyed off `STACK_OFFSET`/`STACK_PROJECT` like the
Go probes. This answers overview open-question #1 (harness home) and #4 (Playwright wiring — already pinned).

## Cockpit-handshake login as an employee hero
TODO: how the harness logs in as a roster hero (Maya Chen) via the cockpit handshake (rides the Clerkenstein
injection + demo lifecycle from `rosetta_demo.md`); the employee/member vantage selection.

## In-app nav crawl (as the employee vantage)
TODO: the crawl strategy (nav-link discovery vs a route manifest) over the demo UI tier (next-web +
studio-desk + ant-academy) on offset ports; reachable-page enumeration.

## Per-page assertions — DOM non-emptiness + screenshot + demo-local link host
TODO: the non-empty semantic-content assertion (real text/rows per section, not a shell); the populated
screenshot capture; the every-link-host-is-demo-local assertion (offset port, not prod).

## The coverage report (pages, pass/fail, escapes)
TODO: the report schema the sweep emits (pages, pass/fail per page, escape list) — the gate input (0 failing
pages + 0 escapes).

## Fix surface: stack-seeding (empty section / missing seed)
TODO: seeding the data a page reads so its sections fill (the rext `stack-seeding` side).

## Fix surface: stack-snapshot serve-grants (federation / content error)
TODO: serve-grants so a page reads replayed content (the rext `stack-snapshot` side).

## Fix surface: the demo injection + env link-rewriting (out-of-demo link)
TODO: rewriting an escaping link host to the demo-local offset port in the demo injection/env (e.g. left-menu
"Studio" → local studio-desk, NOT prod `studio.anthropos.work`).

## Re-scope / escalation: platform-only blockers (zero-edit line)
TODO: recording a 100%-blocking gap that can ONLY be closed by a platform change — escalate, do NOT edit
next-web/app/cms/jobsimulation (read-only this release).

## corpus/ops/demo/coverage-protocol.md (NEW — authored by this milestone)
TODO: the iteration protocol itself — the measure → triage → fix → re-sweep loop + the fix-surface routing.
