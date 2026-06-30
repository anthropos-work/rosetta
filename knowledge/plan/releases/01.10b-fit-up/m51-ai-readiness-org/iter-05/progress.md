**Type:** tik (tooling-iter — coverage-protocol.md "Iter type selection" refinement) — under TOK-01 (active-cycle signals-true).

# iter-05 — deepen the manager-grid WARM (content-presence) + land TOK-01 strand-4 (manifest AI-readiness assertion + cockpit jump_to + org-agnostic email fix)

## What was attempted
Per the overview Phase plan, against the LIVE demo-1 (UP @ fit-up-m50, AI-readiness showcase org seeded — no
re-up/re-seed; harness-only changes run against the live offset ports):

- **Phase C — the harness deepening (the tooling deliverable):** added `warmHeavyGrids` to
  `stack-verify/e2e/lib/section-assert.ts` — a **content-presence** warm for the manager's heavy org-scale
  `/enterprise/*` grids: navigate, then wait for REAL ROWS to paint (via the new `hasRealRows` helper — mirrors
  `isOnlySkeleton`'s real-row selector at page scope, requiring real rows to outnumber skeleton rows) up to a
  generous `HEAVY_GRID_WARM_CEILING_MS = 25_000` ceiling (> the measured ~11.6s cold query), polling every 1s,
  stopping early on the first hydrated grid. It REPLACES the prior plain-`warmStack` push of the manager grids
  (which bailed at the 4s networkidle ceiling — and next-web's long-poll connections never go networkidle — so
  React Query cancelled the in-flight cold query on unmount and the authoritative visit paid the cold cost again
  → a skeleton-frame false-fail the bounded re-assert couldn't recover). Wired into `tests/coverage.spec.ts`'s
  warm step (manager vantage only). **Best-effort + bounded:** a grid that never paints real rows within the
  ceiling is left as-is, so the authoritative assert still FAILs honestly — the warm never masks a
  genuinely-empty grid (a STRENGTHENING that cannot mask a real empty).

- **TOK-01 strand-4 — the manifest AI-readiness assertion (the milestone's headline gate proof):** added the
  `/enterprise/workforce/ai-readiness` page descriptor to `coverage-manifest.ts` (MANAGER_PAGES) with two
  section assertions — `ai-readiness-org-score` (the HeroCard: "AI Readiness" + "Overall org readiness" +
  "Members" → proves the org-aggregate rendered, not Empty/error) and `ai-readiness-funnel` (the 3-step
  completion funnel: "Stage breakdown" + Stage 1/2/3 + "Steps completion" → the milestone's core proof). Added
  the route to `MANAGER_MANIFEST.seedPaths` to PRIME the BFS to visit it (it's a `(new)` route group with NO
  nav-link — confirmed by exhaustive grep — so it must be seeded; a seed that 404s/redirects is dropped, so
  seeding can't false-inflate). **Without this the M51 seed renders on the live dashboard but the gate never
  ASSERTED it** (the iter-03/iter-04 "second learning") — now the gate PROVES the seeded funnel renders.

- **TOK-01 strand-4 — the cockpit jump_to re-point:** `cockpit.go` `DeepLinkCatalog()` gained the
  `ai-readiness` deep-link (`/enterprise/workforce/ai-readiness`, vantage manager); `stories.seed.yaml` Dana's
  `jump_to` re-pointed `/enterprise/workforce` → `/enterprise/workforce/ai-readiness` (the showcase manager
  hero now lands ON the AI-readiness dashboard).

- **The org-agnostic email fix (a gate CORRECTION, root-cause):** the SHARED manager manifest hardcoded
  `cervato-systems.com` (the M42m Cervato calibration) in `members-roster` + `assign-roster` `mustInclude` —
  which **false-failed the Northwind roster even though it rendered real rows** (a stale-org-domain false-fail,
  a root cause behind the perf-wall noise). Replaced with STRUCTURAL tokens only (`Members` + `Location`;
  `Assign AI Simulation`) and added a net-new `members-emails` section asserting an **org-agnostic
  corporate-email regex** (`[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}`, `minMatches: 1`) — so a blank/skeleton grid
  matches NONE → FAIL, a hydrated roster matches every visible member email, ORG-AGNOSTICALLY. The real-email
  proof is preserved (org-agnostic), the org-hardcoding removed.

## Phase D — re-measure (GATED manager sweep on demo-1)
GATED manager-vantage semantic sweep (`coverage-out/manager/coverage-report.json`, generatedAt 16:47Z):
**`failingSections = 5` (from 6), `escapes = 0`, `personaFailures = 0`, `reachable = 65` (from 49),
`frontierRemaining = 0` (frontier-exhausted), persona all GREEN** (role↔skills coherent · avatar real-photo
consistent menu==profile · org "Northwind Aviation" + real logo). **Delta: failingSections 6 → 5 (−1);
reachable 49 → 65 (+16).** The root cause cleared was the **stale-Sentinel-policy** (the Northwind casbin
grants made the org-scale `/enterprise/*` surface reachable — reachable jumped 49→65) — confirming the prior
6 were partly an authz-reachability artifact, not pure perf.

The residual **5** are base-Workforce org-scale sections still showing skeleton/below-floor at the authoritative
visit: `verification-funnel` + `talent-languages` (/enterprise/workforce), `members-roster` +
`assign-roster` (the `cervato-systems.com` false-fails — **already corrected by THIS iter's org-agnostic
manifest fix, which had not yet taken effect in the 16:47 report**), `activity-table` (177 < 200 floor). The
2 `cervato-systems.com` rows are eliminated by the committed manifest correction; the other 3 are the residual
org-scale cold-grid perf-wall → next iter (re-sweep with the corrected manifest + the deepened warm both in
effect).

## Close — 2026-06-30

**Outcome:** failingSections 6 → 5 (−1), reachable 49 → 65 (+16), escapes 0, persona GREEN. Cleared the
stale-Sentinel-policy reachability root cause; landed the content-presence `warmHeavyGrids`, the gate-proving
AI-readiness manifest descriptor (the milestone's headline assertion), the cockpit jump_to re-point, and the
org-agnostic email correction (fixes 2 of the residual 5 on the next sweep).
**Type:** tik (tooling-iter)
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (content-presence warm vs networkidle bail), D2 (org-agnostic email regex replacing the hardcoded org domain — a gate CORRECTION not a loosening), D3 (the AI-readiness manifest descriptor + seedPaths prime).
**Side-deliverables (if any):** none — all landed work is planned TOK-01 strand-4 + the tooling deliverable.
**Routes carried forward:**
  - The residual 3 genuine perf-wall sections (`verification-funnel`, `talent-languages`, `activity-table`) →
    iter-06 (re-sweep with the corrected manifest + deepened warm in effect; triage whether the content-presence
    warm cleared them or a further lever — widened per-section heavy-grid poll / disclosed-presenter-note for a
    genuinely-slow-but-correct heavy section — is needed). Handler: `PERF-M51-iter06-residual-workforce-grids`.
  - The 2 `cervato-systems.com` false-fails are NOT carried forward as bugs — the committed org-agnostic
    manifest correction resolves them; iter-06's re-sweep verifies.
  - **Bake Sentinel-reload-after-seed into the seeding flow** (so the casbin grants take effect without a manual
    reload — a real fix, surfaced by the stale-Sentinel-policy root cause cleared this iter) → iter-06.
    Handler: `FIX-M51-iter06-sentinel-reload-after-seed`.
**Lessons:**
  - A SHARED coverage manifest must assert ORG-AGNOSTIC structural/regex tokens, never a calibration-org's
    literal domain — a hardcoded org email false-fails every other vantage org even when it renders real rows.
    (Generalizes: applies to any multi-org coverage gate → also captured in the manifest comment.)
  - A networkidle warm is the wrong primitive for next-web's heavy org-scale grids (long-poll connections never
    go idle; the cold federated query outruns the 4s ceiling). A content-presence (real-rows) warm with a
    generous ceiling is the right primitive — and stays honest (best-effort, never masks a genuinely-empty grid).
  - The pre-fix 6 were a MIX (authz-reachability + stale-org-domain + cold-grid perf), not a single perf cluster
    — clearing the Sentinel grants alone moved reachability +16 and failingSections −1.
