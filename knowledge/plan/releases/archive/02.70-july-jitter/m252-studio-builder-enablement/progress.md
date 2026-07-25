# M252 — Progress

Section milestone. Checklist stub from the roadmap In-list.

## Sections

- [x] **AI-key demo-container wiring** — the provisioned studio-desk AI key reaches the demo container
  at runtime via an existence-guarded `env_file: ["<clone>/studio-desk/.env"]` in `gen_injected_override.py`
  `frontend_lines()` (`platform_dir` threaded, default `None` → `exposure_claim_guard` byte-identical);
  `/api/ai/completion` no longer 500s. **PROVEN live** (op1, demo-2): the container carries the AI keys +
  boots `ProviderHealth Initialized with chain: azure-openai->openai->anthropic`. (rext `4486cdd`)
- [x] **DNA hardening** — a demo-aware, non-fatal, values-blind autoverify `(g)` cheap-win asserts the
  studio-desk **container** carries a provider key (presence-of-NAME only; mirrors the directus
  `DB_CONNECTION_STRING` container check). 8 new studio autoverify tests green. (rext `4486cdd`; D2)
- [x] **Builder Playthrough** — `studio-builder-page.ts` page-object + `studioBaseUrl(9000+offset)` +
  studio Clerkenstein hero-login (`pt-manager`) + `manifest/studio-builders.yaml` (2 PTs:
  `pt-studio-advanced-generate` + `pt-studio-guided-generate`) + admin/content_creator precondition —
  studio-desk's first playthroughs-manifest entry. `ptvalidate` **18 live / 0 TODO**, unit-green.
  (rext `d80db9f`; D3) — the ~10-min async live RUN → **CARRY-M252-02** (Fate 2 → M254 (e)+(h)).
- [x] **Talk-to-data double-check** — M239 Bedrock path re-confirmed **COMPLETE, no gap** (audit trail; D4).
- [x] **Delivers** — `corpus/services/studio-desk.md` + `corpus/ops/secrets-spec.md` (demo-aware studio-desk
  AI note, incl. the KB-1 correction) + `corpus/ops/demo/frontend-tier.md` (the F8 env_file note) +
  `corpus/ops/demo/playthroughs.md` (the builder Playthrough + count 16→18). (rosetta `cedde09`)

## M252: Final Review

### Scope
- [x] Live end-to-end builder-generate drive → recorded as **CARRY-M252-02** (Fate 2 → M254 gate (e)+(h)); the
  BUILD deliverables (wiring PROVEN op1 + DNA assert + Playthrough artifacts) are complete.

### Code Quality
- [x] rext code-of-record reviewed (env_file existence-guarded + values-blind + byte-identical backcompat;
  autoverify (g) demo-aware + non-fatal + values-blind) — no issues. Shipped = docs (env_file-only). (D6)

### Documentation
- [x] 4 corpus docs accurate + internally consistent with shipped code; all relative cross-refs resolve.
- [x] 2 stale count-mirrors (`README.md:206`, `CLAUDE.md:321` "16 live") → **Fate 2** to the M247-reconcile
  tail / close-release Phase 3b (D5) — CLAUDE.md is M247-sole-owned (coordination rule #5); not fixed here.
- [x] Commit-message imprecision (`cedde09` "MOCK_CLERK/overlay") recorded as D6 (docs are correct; no rewrite).

### Tests & Benchmarks
- [x] rext suites recorded green (test_injection 156 OK, exposure_claim_guard 49 OK, 8 studio autoverify tests,
  ptvalidate 18 live/0 TODO, shellcheck clean). Rosetta is a docs corpus — no test suite.
- [x] 5 pre-existing stack-verify failures (M245 academy `/library/` cheap-win) = **CARRY-M252-01**
  (Fate 2 → M254 (g)); byte-identical pre/post M252 — M252 adds 0 failures.

### Decision Triage
- [x] D1 (env_file model) → already blended into studio-desk.md / frontend-tier.md / secrets-spec.md during build.
- [x] D2 (autoverify container assert) → already blended into secrets-spec.md + studio-desk.md.
- [x] D3 (Playthrough hero + completion-boundary) → already blended into playthroughs.md.
- [x] D4 (Bedrock complete) → archive (audit-trail; maintainer-only).
- [x] D5/D6 (close-time) → archive (maintainer/tracking; recorded in decisions.md).

## Completeness Ledger

### Done (Fate 1)
- AI-key demo-container wiring (env_file, PROVEN live op1).
- DNA hardening (autoverify (g) container-key assert + 8 tests).
- Builder Playthrough artifacts (studio-builders.yaml + page-object + studioBaseUrl + hero-login + precondition;
  ptvalidate 18 live/0 TODO).
- Talk-to-data double-check (M239 Bedrock CONFIRMED COMPLETE).
- All 4 doc Deliverables (studio-desk.md, secrets-spec.md incl. KB-1 correction, frontend-tier.md, playthroughs.md).

### Confirmed-covered (Fate 2)
- **CARRY-M252-01** — 5 pre-existing stack-verify failures (M245 academy cheap-win) → M254 exit-gate (g).
- **CARRY-M252-02** — the live end-to-end builder-generate Playthrough RUN → M254 exit-gate (e)+(h).
- **D5** — the 2 stale playthrough-count mirrors (README.md:206, CLAUDE.md:321) → M247-reconcile tail /
  close-release Phase 3b KB consolidation.

### Annotated (Fate 3)
- None.

### Dropped
- None.

### Release-scope-breaking deferral (escape hatch)
- None.
