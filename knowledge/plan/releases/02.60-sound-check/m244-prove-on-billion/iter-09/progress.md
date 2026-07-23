**Type:** tik — gate (d) measurement, under TOK-01. Run 4, tik 4.

# iter-09 — progress

## Measurement (live, anon, billion academy :13077, Playwright render)
- **`/free`** → `/?tier=free`: **PASS** — 43 real course-card links rendered (mainLen 1281). NB **drafts=2**
  (2 "Draft" chips — content-stories-routes.md expects 0; a minor sub-finding, see routes).
- **`/library`** → `/library/`: **FAIL** — the grid renders **EMPTY** (0 course links, 0 cards; bodyLen 130 =
  header only: "The AI Academy library / Browse 140+ chapters across AI engineering, coding agents, and applied
  AI."). Stable empty (networkidle + 9s settle), not still-loading.

**⇒ Gate (d) is NOT met on billion: the twin does not both-render — `/free` shows cards, anon `/library` is an
empty grid.**

## Root cause (characterized)
`ant-academy/code/src/virtualLibrary/sources/publicSource.js` documents that the library grid "reads the
tenant-filtered catalog from **ServerCatalogContext / LocalizedCatalogContext**" (it deliberately does NOT
import the FS `ucourses/catalog.js`, to avoid leaking other orgs' tenant content to the browser). On a DEMO the
academy is **anonymous / Clerk-free with no tenant entitlement**, so that tenant-filtered catalog is EMPTY →
the `/library` grid renders no cards. `/free` renders because it reads a different (public/tier) source. This
is the "empty academy" class (v2.5 M229/M230) on the `/library` surface specifically — the demo academy serves
its committed FS catalog to `/free` but the tenant-gated `/library` grid has no reader for anon.

The m244 cold reset-to-seed will NOT fix this (it changes content-story data, not the academy catalog source /
tenant-gating). The fix is an ant-academy demopatch/config that serves the FS catalog to the `/library` grid
for the anon demo vantage — a real, non-trivial follow-up. 0 platform edits (route via demopatch, never edit
the ant-academy repo).

## Close — 2026-07-22

**Outcome:** gate (d) MEASURED live on billion — `/free` renders 43 real cards (✓), anon `/library` renders an
EMPTY grid (✗); the twin does not both-render, so gate (d) is NOT met. Root cause characterized (the `/library`
grid reads a tenant-filtered catalog that is empty for an anon demo user). No fix landed this iter (the
ant-academy demopatch/config is a substantial follow-up); the defect + root cause are documented + routed.
Metric stays **3/8**.
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET (gate (d) measured, not met; milestone 3/8)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (closed-no-lift is excluded from the no-prog streak; iter-07/08 made progress) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik 4 of run 4) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (measure gate d against current billion since it's re-seed-independent), D2 (root cause = tenant-filtered catalog empty for anon; fix is an ant-academy demopatch, routed not landed) — iter-09/decisions.md
**Side-deliverables:** none.
**Routes carried forward:**
- **(gate d fix, Fate-3) → a future M244 iter:** an ant-academy demopatch/config so the anon `/library` grid
  serves the FS catalog (like `/free`). Root cause documented above. Named handler: gate-(d) `/library` fix.
- **(sub-finding, Fate-3) → same fix iter:** `/free` shows **2 Draft chips** (expected 0) — check the FS
  catalog's draft filtering for the demo.
**Lessons:** (1) gate (d) is a genuine TWIN — `/free` and `/library` have DIFFERENT catalog sources; a passing
`/free` does not imply `/library`. (2) the academy `/library` grid is tenant-gated by design (no FS-catalog
import, to avoid cross-tenant leakage), so an anon demo needs an explicit demopatch/config to feed it — the
re-seed alone won't. (3) verifying a render gate against the CURRENT billion (no re-seed needed) is a cheap way
to advance a re-seed-independent gate part while the big cold reset-to-seed waits for a focused run.
