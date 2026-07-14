# M219 — Progress

_Section checklist. Populated from `overview.md` § Scope.In at build time; closed by `/developer-kit:close-milestone`._

## Sections

- [ ] **S1 — The surface split: current vs legacy, and every pointer repointed.** *(overview items 9 + 4)*
  - [ ] Establish + **document** the current/legacy split for BOTH vantages in `corpus/services/ai-readiness.md`
        (routes named; the legacy orphan named; the employee surface's no-route-of-its-own fact named).
  - [ ] Correct KB-1 (the misattributed `?cycle=` caveat), KB-3, and the missing `ai_readiness_recommendations` table.
  - [ ] Repoint the cockpit deep-link catalog (`stack-seeding/seeders/cockpit.go`) → `/ai-readiness`; add the
        **missing** `end-user`-vantage readiness entry.
  - [ ] Repoint Dana's `jump_to` → `/ai-readiness`; point Aria's + Ben's → the employee readiness surface (`/home`).
  - [ ] Repoint the coverage manifest's `AI_READINESS_PAGE` → `/ai-readiness`; re-derive its section asserts against
        the CURRENT page's DOM (the legacy asserts do not hold on it).
  - [ ] Fix the stale ACTIVE-vs-CLOSED comment (`stories.seed.yaml:112-117`) — item 4.
  - [ ] Regression tests: a pointer that resolves to a legacy route must FAIL LOUD.

- [ ] **S2 — Every element and sub-section FILLED, on both vantages.** *(overview items 1, 2, 3, 8)*
  - [ ] Seed the **ACTIVE cycle** (F-6) → Ben's `progress` hero + Aria's full `done` hero + the manager's
        `interview` / `diagnosis` / `sources` sections (F-4, F-5).
  - [ ] Keep the closed cycle (cycle history in the CyclePill; the 199 frozen rows stay live).
  - [ ] Close the residual zero cell: `howWeMeasure.cycleTotals.interviewQuestions = 0` (F-4).
  - [ ] Per-section manifest for BOTH readiness surfaces; assert every section non-empty + persona-consistent.
  - [ ] Fix the stale `app-aireadiness-snapshot-loadmembers` manifest header comment (F-7).
  - [ ] Re-measure the manager load and **report** it (D-DESIGN-1: reported, not gated).

- [ ] **S3 — `FIX-M219-bapi-org-eid` (F-11).** *(overview item 5 — inherited from M218)*
  - [ ] `SeedOrgIdentity` / `LookupOrgEid` on the BAPI store; wire the roster's real `org_eid` through
        `seedRosterMemberships` → `organizationWithEid`; keep the demo-org + stub fallbacks.
  - [ ] The `MembershipOrgIdentity/real-org-eid` gene goes **GREEN** (Go surface 97.2% → 100%).
  - [ ] A fresh **5-cycle cold battery** (it is on the demo runtime path).

- [ ] **S4 — The two absence-read-as-success gates.** *(overview items 6 + 7 — inherited from M218)*
  - [ ] `TEST-M219-expressrun-dep-gate` — a missing dep must **fail loud**, never present as absence-of-a-score.
  - [ ] `TEST-M219-freshness-gate-skips` — the silent skip becomes explicit and loud.
  - [ ] Both proven RED against the pre-fix commit, GREEN after.

- [ ] **S5 — The `ai-readiness` playthrough.** *(overview § Delivers — a BLIND AREA)*
  - [ ] `pt-world` gains an `ai-readiness` org + the completed/started/manager seats; `seed-worlds.yaml` gains the
        capabilities (else `checkPreconditionCoverage` hard-fails).
  - [ ] The `ai-readiness` product manifest + page objects + specs; `ptvalidate` green.
  - [ ] Its section in `corpus/ops/demo/playthroughs.md`.

## Notes

- **Phase 0b — KB-fidelity: YELLOW** (satisfied by the census; see `spec-notes.md`). KB-1..KB-3 fixed in S1.
- **Two of the overview's premises were REFUTED by measurement** (F-2, F-7) and the planned **new demo-patch is
  withdrawn** — the non-patch fix (point the demo at the *current* surface) is available, which is the correct
  order of preference per `demopatch-spec.md §1`. **Zero platform-repo edits; zero new demo-patches.**
- The user's kickoff report is **confirmed in code**: every demo pointer targets the **legacy** manager page (F-1).
