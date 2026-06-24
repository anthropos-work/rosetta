# M42e Spec Notes

Technical notes accumulate here during the iteration loop. The authoritative spec is
[`.agentspace/profile_gaps.md`](../../../../.agentspace/profile_gaps.md) (live-demo review, 2026-06-24;
root-cause workflow w7t4wq2z4). The iteration protocol this milestone authors is
[`corpus/ops/demo/coverage-protocol.md`](../../../../corpus/ops/demo/coverage-protocol.md).

## The Playwright coverage harness (a new rext section — `stack-coverage` or under `stack-verify`)
TODO: the harness home + dir layout; demo-only; the first non-Go rext dev/test dependency (Playwright pinned
under its own dir + lockfile) — how it's wired alongside the Go rext tooling.

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
