# M53 — progress

## Section checklist
_Checked off as each In-scope deliverable lands. Close when all boxes are ticked._

_Ordered per overview.md acceptance flow (a)→(f). Sections gate on prior sections._

- [x] **§1 — Academy F6 seeder/wiring** (the ONE new-code section; in the rext authoring copy) — rext `e91f004`
  - [x] Course content present (3250 static-JSON files ship with the clone — verify-only, not seeded)
  - [x] Hero academy menu-link: per-hero cockpit [Academy] deep-link (cockpit.go External entry + cockpit.py + up-injected --academy-base)
  - [x] Non-anonymous academy session: launcher sets BOTH e2e_persona bypass gates; cockpit link sets e2e_persona=member cookie
  - [x] AI chat documented-as-absent (Cosmo flag+key not provisioned; NO `/api/ai/chat` assertion) — D3 + frontend-tier.md
  - [x] Commit on rext authoring `main` (e91f004; tests +13, all green, shellcheck-clean)
- [x] **§2 — Roll v1.10.1 rext release tag** (rolls up `fit-up-m47..m52` + academy commit; **RE-ROLLED at the acceptance gate to include the AB4 fix**)
  - [x] Tag `v1.10.1` on rext authoring `main` (annotated; **re-rolled to point at the AB4 fix HEAD `117fe41`**; originally `e91f004`, 46-commit roll-up + academy F6 + AB4 fix)
  - [x] Bump `.agentspace/rext.tag` → `v1.10.1` + canonical pin in rosetta_demo.md
  - [x] Re-pin consumption clone `stack-demo/rosetta-extensions` → re-rolled `v1.10.1` (`117fe41`) (clean fetch + checkout; tree clean)
- [x] **§3 — DESTROY the live demo** (`/demo-down 1 --purge`) — all 17 containers + network removed, ALL demo-1 images purged (M49 #6 verified working); native academy/cockpit stopped
- [x] **§4 — COLD REBUILD** (single `/demo-up 1` at v1.10.1 pin, no manual steps) — EXIT 0, no #7 abort; 17 containers Up (0 Exited); autoverify GREEN; log: `cold-rebuild.log`
- [x] **§5 — ASSERT the acceptance bar** (6/6 + academy F6 PASS; **AB4 fixed at the gate — M51-owned regression, user-approved exception**)
  - [x] AB1 — all backends healthy: 17 containers Up, 0 Exited; health 200, casbin 1150, all probes passed
  - [x] AB2 — snapshot replay prompt-free from the filled cache (taxonomy/directus/sim-embeddings replayed, no prompt) — KB-1
  - [x] AB3 — set-dress + seed (3 orgs incl. Northwind AI-readiness) + verify + cockpit — EXIT 0, no #7 abort
  - [x] **AB4 — GREEN (fixed at gate): employee GREEN (0,0); manager `dan-manager`@Cervato GREEN after fix (reachable=69, failingSections 2→0, escapes=0, persona=0) — org-conditional manager manifest, rext `117fe41` (see decisions.md AB4-FIX). Re-verified both manager vantages.**
  - [x] AB5 — AI-readiness dashboard GREEN on 3rd org (dana-manager@Northwind), **re-verified post-fix**: reachable=70, both ai-readiness sections PASS (541 chars), 50/100, 199 members, 3-step funnel, renders fast — KB-2
  - [x] AB6 — cockpit [Download manifest] serves complete inlined `seed-generation-manifest.yaml` (7593B, 3 orgs + generation prompt + batch $0.3 ceiling + snapshot_sources)
  - [x] F6 — academy: content real (copilot/claude-code/ai-eng chapters), 9 cockpit [Academy] links→:13077, both e2e_persona gates set (authenticated member), Cosmo absent by design
- [x] **§6 — Acceptance record + rext.tag bump** (corpus acceptance note; feeds close-release) — acceptance-record.md updated to 6/6+F6 GREEN; AB4 fix + re-roll recorded

## M53: Hardening

### Pass 1 — 2026-07-01
**Scope (milestone-touched, from Phase 1):** rext authoring copy, two commits —
- F6 academy seeder/wiring (`e91f004`): `stack-seeding/seeders/cockpit.go` (+ `cockpit_test.go`),
  `demo-stack/cockpit.py` (+ `tests/test_cockpit.py`), `demo-stack/ant-academy.sh`,
  `demo-stack/up-injected.sh` (+ `tests/test_ant_academy.py`).
- AB4 manager-manifest org-conditional (`117fe41`): `stack-verify/e2e/lib/coverage-manifest.ts`
  (+ `tests/coverage-manifest.unit.spec.ts`), `tests/coverage.spec.ts`.

**Coverage delta (milestone-touched files):**
- Go `stack-seeding/seeders/cockpit.go`: 97.6% → 97.6% (unchanged — at achievable ceiling; the
  sole gap is `AcademyDeepLink`'s `return DeepLink{}, false` fall-through, UNREACHABLE via the
  constant `DeepLinkCatalog()` — a defensive branch, now documented by a test rather than contorted
  into false coverage).
- Python `demo-stack/cockpit.py`: 98% → 98% (unchanged — the 3 misses are the `KeyboardInterrupt`
  handler [576-577] + the `__main__` guard [584], untestable-by-design idioms).
- TS `coverage-manifest.ts`: covered by 29 passing unit tests (was 27). No line tool locally; the
  pure-logic manifest selection is exhaustively behavior-tested.

The deltas are 0% because both surfaces were already at their achievable statement ceiling at build
(F6 +13, AB4 +3). The hardening value here is BEHAVIORAL depth over new lines (finder-not-goal).

**Tests added (behavior, not lines):**
- `coverage-manifest.unit.spec.ts`: +2 — (1) `manifestFor` org-gating boundary edges: whitespace-only
  org, a PARTIAL showcase name (`"Northwind"` w/o `"Aviation"`) does NOT promote to showcase, the
  showcase name embedded in a LARGER label DOES, employee vantage is org-independent; (2) `manifestFor`
  referential stability (same singleton across calls — no per-call rebuild drift). These pin the exact
  org-substring gate that separates the regressed-then-fixed M50 base gate from the AB5 showcase gate.
- `cockpit_test.go`: +2 — `AcademyDeepLink()` returns the catalog's external entry VERBATIM
  (single-source with the manifest projection) + exactly-one-external invariant (documents the
  not-found return as intentional dead defensive code); call stability.
- `test_cockpit.py`: +5 (`TestAcademyCatalogEntryEdges`) — `_academy_catalog_entry` selects the academy
  entry among multiple externals / skips a non-external decoy / handles a missing catalog key; the
  renderer defaults path/persona/label on a tampered entry + HTML-escapes them.

**Bugs fixed inline:** none — no production code changed; no bug surfaced.

**Flakes stabilized:** none — flake gate 3/3 clean sequential runs (Go seeders subset + Python academy
subset + TS manifest suite).

**Knowledge backfill:** no KB-worthy findings — the two surfaces are already documented
(`snapshot`/`cockpit-spec.md` for F6, `coverage-protocol.md` + `decisions.md` AB4-FIX for the
org-conditional manifest); the harden added edge tests around already-documented behavior, surfacing no
new invariant.

### Stop condition
Stopped after Pass 1: the full Step 2b six-dimension scan found nothing further worth adding (the
remaining uncovered lines are untestable idioms + one unreachable defensive branch, now documented),
coverage delta < 2% on both measurable stacks, and 0 flakes across 3 consecutive sequential runs.

**rext authoring HEAD moved to `576dbcb`** (one harden commit past `v1.10.1` = `117fe41`) — `v1.10.1`
may need a re-roll at close to include the harden tests (the tag-at-close precedent).

## M53: Final Review

### Scope
- [x] All §1–§6 boxes checked; acceptance-only milestone; F6+AB4 sanctioned exceptions landed + tagged. No new deferral originates in M53.

### Code Quality
- [x] [confirm] F6 academy DeepLink is a single-source catalog entry; AB4 `manifestFor` org-gate is case-insensitive substring; shellcheck clean — 0 findings (build + harden already reviewed).

### Documentation
- [x] AB4 org-conditional manager manifest was UNDOCUMENTED in `coverage-protocol.md` → added the M53 AB4 org-conditional note (showcase `MANAGER_MANIFEST` vs `MANAGER_MANIFEST_BASE`).
- [x] KB-2 stale round number: `ai-readiness.md` carried `~80% (≈160 members)` / `≈160 "completed"` → reconciled to the shipped **78.4% (≈156 of 199)** (2 spots), consistent with `seeding-spec.md`.

### Tests & Benchmarks
- [x] Full unit suites green (no coverage sweep per acceptance-close directive): Go `stack-seeding` all ok (791 Test/Fuzz funcs) + other 4 Go modules green; Python `demo-stack` **326**; TS unit **29**. Flake gate below.

### Decision Triage
- [x] D1/D2/D3 (academy F6 mechanics) → already blended into `frontend-tier.md` + `ant-academy.md` (verified accurate + tagged).
- [x] AB4-FIX (org-conditional manifest) → blended into `coverage-protocol.md` (#M53 tag).
- [x] KB-1 (replay-from-filled-cache reading of AB2), KB-2 investigation, AB4-REGRESSION/ROUTING narrative → archive (maintainer-only; #M53 refs in `decisions.md`).

### Release-tag step
- [x] Re-roll `v1.10.1` to the final rext HEAD `576dbcb` (incl. the harden tests) + re-pin consumption clone.
