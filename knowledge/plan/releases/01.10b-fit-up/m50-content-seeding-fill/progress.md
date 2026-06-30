# M50 — progress

## Running ledger
_Appended after each iter (tik/tok). Each entry: what was attempted, what moved, gate distance._

- iter-01 (tok·bootstrap): authored TOK-01 + FRESH-demo-1 re-diagnosis (4 genuine seed-gap clusters; has-data surfaces flagged for sweep) — see iter-01/progress.md
- iter-02 (tik, closed-fixed-partial): EMPLOYEE gate MET on baseline (valid, frontier-exhausted, 59 pages, all 0) → employee half needs no fix; member-field fill (memberships join-date/location/last-activity) authored+tested+committed to rext (`fix(M50/02)`); MANAGER baseline (prematurely) read as frontier-capped → iter-03 — see iter-02/progress.md
- iter-03 (tik, closed-fixed): re-survey CORRECTED iter-02 — the manager sweep EXHAUSTS, not caps (cap=300 → reachable=68, cappedAtFrontier=false, gate-VALID); tooling-iter CANCELLED. Manager verdict: failingSections=0 personaFailures=0 crossPortFollowFails=0 escapes=1 → the SOLE blocker is a prod-eject to `anthropos.work` on the activity-dashboard sim drill-down (then-attributed to hardcoded `PUBLIC_WEBSITE_URL`). Content fully populated. — see iter-03/progress.md
- iter-04 (tik, closed-fixed-partial): built+tested+verified the `next-web-public-website-url` demopatch on demo-1 (cleared the JS-constant-built ejects — 6 AD drill-downs clean) BUT the re-sweep shows escapes still 1: the residual is a DIFFERENT class — replayed Directus content `public_landing_page_url`/`read_more_link` carrying a prod `anthropos.work` URL (28/14 sims). Routes to iter-05 (stack-snapshot content rewrite). — see iter-04/progress.md
- iter-05 (tik, closed-fixed, **GATE: MET**): post-replay Directus content-URL rewrite (simulations + skill_paths, regex over any anthropos.work subdomain → demo host) + cms cache clear killed the residual escape. **Manager re-sweep FINAL: reachable=69 failingSections=0 escapes=0 personaFailures=0 crossPortFailures=0 frontier=EXHAUSTED gateMet=True.** With employee (iter-02) → **M42 gate GREEN BOTH vantages on warm demo-1.** — see iter-05/progress.md
- iter-06 (tik, closed-fixed, **GATE: MET — run 2**): the D4/F1 reconciliation. Filled the two annotation gaps the run-1 gate passed BLIND: NEW `MemberLanguagesSeeder` (world_languages ISO catalog + per-member user_languages → membership_languages via the DB trigger) + extended `CertificatesSeeder` to ~45% role-coherent member coverage (cert rows 9→236). STRENGTHENED the manager manifest (NEW `preAssert` tab-click + `textMatch` OR-assert harness primitives) to ASSERT members-Location + Talent-tab languages/certs. **Manager re-sweep on the STRENGTHENED manifest: reachable=69 failingSections=0 escapes=0 personaFailures=0 frontier=EXHAUSTED gateMet=True** — the new sections all PASS real-content. M17-idempotent (2nd seed=0 rows). M42 gate now GREEN both vantages on the manifest that PROVES the gaps. Side-fix: run-coverage.sh arg-forwarding footgun. — see iter-06/progress.md

## GATE STATUS
**M42 coverage gate MET on BOTH vantages on the WARM demo-1, on the STRENGTHENED manifest** (employee iter-02 + manager iter-06; both frontier-exhausted, (failingSections,escapes)=(0,0), 0 persona, 0 cross-port). **Run 2 closed the D4/F1 reconciliation:** the run-1 gate passed BLIND to two M50-own annotation gaps (languages 0 rows, certs hero-only 9/340); iter-06 FILLED them (`MemberLanguagesSeeder` + the `CertificatesSeeder` member-coverage extension → 747 user_languages across all members + certs 9→236) and STRENGTHENED the manager manifest (new `preAssert` tab-click + `textMatch` primitives) to ASSERT members-Location + the Talent-tab languages/certs charts — the gate is now MET on the manifest that PROVES the gaps, not the blind one. All M50 seeders + fixes reproduce from the bring-up tooling. **The explicit milestone exit_gate ("on a COLD reset-to-seed demo") is the remaining acceptance step** — reserved for the heavy COLD pass (close/harden + the v1.10b M53 cold-rebuild milestone).

## Next-iter queue (routes carried forward — to close/harden + M53, NOT blocking the warm metric)
- **COLD reset-to-seed acceptance** (the explicit exit_gate): fresh `/demo-up` (all M50 seeders + fixes reproduce from tooling) + both-vantage sweeps on the strengthened manifest → confirm (0,0) both vantages on COLD. The v1.10b M53 "cold-rebuild acceptance" milestone owns this; close-milestone runs it or surfaces to the orchestrator.
- AI-provider-keys policy (F7) + academy menu-link/non-anonymous-session (F6): decision deliverables → secrets-spec.md; academy AI chat documented-as-absent (not a gate blocker) — for close/M51.
- Re-pin the consumption clone (`stack-demo/rosetta-extensions`) to the `fit-up-m50` tag at close (it carries the iter-04/05/06 fixes synced for live verification).
- (RESOLVED by iter-06) ~~manager manifest-strengthening (D4/F1)~~ — DONE; ~~languages seeder + cert coverage~~ — DONE + gate-PROVEN.

## M50: Final Review
_Close review (2026-06-30). Phases 1–5 ran as parallel scans (deferral audit blocking-gate GREEN; code-quality + test-coverage in the rext authoring copy; doc + adversarial review in the main thread). Iterative shape → Phase 9-iter (Gate Outcome Ledger). Default fix-everything; no escape-hatch deferrals surfaced._

### Scope
- [x] Gate-distance: M42 coverage gate MET both vantages (employee iter-02 + manager iter-06) on the STRENGTHENED manifest, frontier-exhausted, (failingSections,escapes)=(0,0). COLD acceptance → M53 (Fate-2, user-decided). No scope gap.
- [x] Iter-ledger: 6 iters (1 tok + 5 tiks), all closed with `**Gate:**` fields; one-commit-per-iter (d1b8d79…380aa0a) + 2 harden commits (c671633, 257ba8e). No orphan iters/commits.
- [x] No code TODO/FIXME/HACK in the M50-touched rext files (verified by the code-quality scan).

### Code Quality
- [x] [must-fix] `gofmt -l` flags `member_languages.go` + `users.go` (trailing-comment misalignment in `[]any{}` row literals — would fail a CI fmt gate). → `gofmt -w` both, re-test.
- [x] [nice-to-have] `member_languages.go` `nativeLanguageByCity`↔`userprofile.go` `locations` coupled by a hand-maintained "KEEP IN SYNC" comment → add a unit test pinning the invariant (every `locations` entry is mapped or deliberately English-fallback).
- [x] [nice-to-have] `member_languages.go` two `user_languages` unique keys (PK `id` + `UNIQUE(user, world_language)`) must stay co-derived; merge only conflicts on `id` → one-line comment noting the co-derivation invariant.

### Documentation
- [x] AI-keys policy (F7) → `secrets-spec.md`: convert the M50-placeholder note into the DECIDED policy (documented-as-absent; values-blind; inert-by-design surfaces; waived/optional class). [done in review]
- [x] NEW escape class `coverage-protocol.md` fix-surface routing table: the **replayed-CONTENT URL-field escape** (`public_landing_page_url`/`read_more_link` carrying a prod host) — the iter-05 post-replay content-URL rewrite. Distinct from the existing serve-grant "Federation/content error" + the JS-constant "Platform-bound escape" rows. Note the `PUBLIC_WEBSITE_URL` demopatch as a Platform-bound-escape instance.
- [x] `seeding-spec.md` version-history roster: add the v1.10b M50 line (`MemberLanguagesSeeder` + the `CertificatesSeeder` member-coverage extension + the member-field backfill) next to the M44 density line.

### Tests & Benchmarks
- [x] Coverage solid (seeders 97.4%, every hardening-ledger number verified accurate; all success/error/edge/idempotency + cross-iter uuid-space paths covered). No new tests required for correctness. Benchmarks not warranted (deterministic COPY seeders).
- [x] [nice-to-have] `test_frontend_build.py` benign `ResourceWarning` (unclosed file handle — `open(...).read()` without `with`). Closed the 2 the coverage scan named (402, 425). The remaining 5 (441/604/774/848/896) are a PRE-EXISTING file idiom OUTSIDE the M50 footprint (M50 added 4 content-URL-rewrite tests in the ~776 hunk, none with a bare `open()`) → not chasing pre-existing noise (three-fate scope = M50's footprint; not gold-plating). Suite stays 108 OK.

### Decision Triage
- [x] AI-keys policy (DECIDED: documented-as-absent) → blend into `secrets-spec.md` (done) + record D-AI-KEYS in `decisions.md`.
- [x] COLD reset-to-seed acceptance → M53 (Fate-2, user-decided) → record in `decisions.md`.
- [x] Academy course-content + menu-link/non-anonymous-session (F6) → M51 (Fate-3) → annotate M51 `overview.md` + record in `decisions.md`.
- [x] iter-05 D1/D2 content-URL-rewrite mechanism → blended into `coverage-protocol.md` routing table (above); the per-iter routing/guard rationale stays archive in the iter `decisions.md`.

## M50: Gate Outcome Ledger (Phase 9-iter)

**Gate**
- **Target:** M42 semantic coverage gate `(failingSections, escapes) = (0,0)` BOTH vantages, frontier-exhausted, 0 prod-eject escapes — on a COLD reset-to-seed demo.
- **Achieved (warm):** employee `(0,0)` reachable=59 frontier-EXHAUSTED (iter-02); manager `(0,0)` reachable=69 frontier-EXHAUSTED, 0 persona, 0 cross-port — on the STRENGTHENED manifest that ASSERTS the M50 fills (iter-06).
- **Distance:** the warm metric is MET on both vantages on the manifest that PROVES the gaps. The remaining clause — "on a COLD reset-to-seed demo" — is a documented Fate-2 carry-forward to M53 (below).
- **Status:** **`closed-on-gate`** (warm both vantages, strengthened manifest). The COLD-environment proof is M53's defining work — a user-decided Fate-2 carry-forward, NOT closed-incomplete, NOT an escape-hatch.

**Iter ledger summary** — 6 iters (1 tok + 5 tiks), all closed with `**Gate:**` fields, one-commit-per-iter:
- iter-01 (tok·bootstrap): TOK-01 strategy + FRESH-demo-1 re-diagnosis (4 genuine seed-gap clusters).
- iter-02 (tik): employee gate MET on baseline; member-field fill (joined_at/location/last_activity) landed.
- iter-03 (tik): corrected the manager baseline — the sweep EXHAUSTS (not frontier-capped), gate-VALID; sole blocker = 1 prod-eject.
- iter-04 (tik): `next-web-public-website-url` demopatch (killed the JS-constant ejects); residual = a DIFFERENT class (replayed content).
- iter-05 (tik): post-replay Directus content-URL rewrite killed the residual → **M42 GATE MET both vantages (warm)**.
- iter-06 (tik): the D4/F1 reconciliation — NEW `MemberLanguagesSeeder` + `CertificatesSeeder` member-coverage + STRENGTHENED manager manifest (preAssert + textMatch) → gate MET on the manifest that PROVES the gaps.

**Routes carried forward** (the three-fate rule — none escape-hatch):
- **COLD reset-to-seed acceptance → M53 (Fate-2, user-decided, D-CLOSE-2).** All M50 seeders + fixes reproduce from the bring-up tooling on a fresh `/demo-up`; M53 owns the single from-cold acceptance truth. No fresh sign-off (user already decided); no plan edit (M53 already lists both-vantage coverage on a from-cold rebuild).
- **Academy content + menu-link/non-anonymous-session (F6) → M51 (Fate-3, D-CLOSE-3).** Not on any M50 gate path; annotated to M51's candidate scope. The academy AI chat is documented-as-absent per the keys policy.
- **Consumption-clone re-pin → push-gated KEEP (release-level).** Target tag advances to `fit-up-m50`; authoritative box-level bump at M53 (same class as M47/M49). Not a repeat-defer.

**Resolved at close** (was inherited DEF-M49-01): **AI-provider-keys policy = documented-as-absent (Fate-1, D-CLOSE-1)** — values-blind, no key provisioned; AI surfaces inert-by-design; documented in `secrets-spec.md`.

**Dropped:** none.

**Protocol evolution:** iter-03 corrected an iter-02 misread (manager sweep EXHAUSTS, not caps — a tooling-iter was CANCELLED). iter-06 added two harness primitives (`preAssert` tab-click + `textMatch` OR-assert) so the manager manifest can ASSERT tab-gated/paginated content — additive (new branches gated on new fields), employee vantage provably unaffected. Lesson pinned: a green `(0,0)` gate is only as honest as its manifest ASSERTS (the run-1 blind-pass → run-2 strengthened-pass reconciliation).
