# M42m — Spec notes

Iteration-protocol-specific technical notes (the manager-vantage coverage sweep, the harness reuse, the
manager-reachable page set, the Workforce-dashboard surfaces, the in-demo/escape boundary). Accumulates across
iters.

## Pre-flight audits — iter-01
`/developer-kit:audit-kb-fidelity` (manual, via sub-agent, 2026-06-25): **YELLOW** — docs accurate where they
make claims (coverage-protocol.md harness description matches the code; all 6 M36 seeders exist + write what
stories-spec claims, 3 spot-checked clean). YELLOW only because M42m's 3 core unknowns are undocumented (the
milestone's designed discovery work, not a fidelity failure). **Two load-bearing facts into iter-01:**
1. **The real manager route is `/enterprise/workforce?tab=skills-verification`** (confirmed in the seeded
   `stories.seed.yaml` cockpit `jump_to` + `test_cockpit.py`), NOT the manifest's guessed `/workforce/*`
   sub-routes. The `MANAGER_PAGES` paths are `calibrated:false` best-guesses → **reconcile the manifest route
   model FIRST** (tab-query, not sub-route). The notReached=5 is very likely a wrong-path manifest guess, not
   purely a nav/seed gap.
2. **`CORS_EXTRA_ORIGINS` on `/api/workforce/*`** is already wired by the demo injection
   (`gen_injected_override.py` per `frontend-tier.md`) — verify before assuming an empty dashboard is a seed
   gap.
Blind areas (expected, = discovery targets): (a) the next-web client page path for the workforce dashboard
(API `/api/workforce/*` documented; client path NOT) — likely `/enterprise/workforce?tab=…`; (b) what gates
the Workforce nav-link visibility (g3-for-sim-start is the closest analog precedent; manager JWT carries
`org_role=admin`); (c) where the baked `studio.anthropos.work` Studio left-nav link is defined + whether it's a
`NEXT_PUBLIC_*_URL` (rewritable) or hardcoded (re-scope trigger if hardcoded).

## Coverage harness + protocol (reused from M42e)
The M42e Playwright harness (`.agentspace/rosetta-extensions/stack-verify/e2e/`) is vantage-generic: the
`MANAGER_MANIFEST` (`coverage-manifest.ts`, seat `dan-manager`) + `run-coverage.sh <N> manager dan-manager`
drive it against the manager hero. The gate composition is identical to employee
(`0 failingSections + 0 personaFailures + 0 escapes + 0 notReached + frontier EXHAUSTED`). A "failing page" =
a manifest section below the content/cardinality bar; an "escape" = an off-demo prod-eject `<a href>` the
crawl's escape-classifier didn't clear via the allow-rule. Persona (role↔skills, avatar menu==profile,
org name+logo) already PASSES for dan-manager (the M42e identity machinery generalizes — iter-23 smoke-sweep).

## Manager-roster hero (demo-3)
TODO: name the canonical manager hero used for the sweep (Dan Rossi / Leah Donovan) + the demo-3 login
identity; note any per-hero gate scoping.

## M36 Workforce-Intelligence dashboard surfaces
TODO: enumerate the manager-only dashboard surfaces the gate covers — the mapped→verified funnel, teams, role
gap/mobility, succession, feedback, and the org-scale claimed-vs-verified gap — as the sweep flags each.

## Admin pages (manager vantage)
TODO: record which admin pages are reachable from the manager vantage and whether each is in-scope for populate
vs in-demo-only.

## rosetta-extensions fixes (manager vantage)
TODO: log the tooling-side fixes (seeding / set-dress / route handling) landed into `rosetta-extensions` to
close manager-vantage gaps — tooling-only, zero platform-repo edits.
