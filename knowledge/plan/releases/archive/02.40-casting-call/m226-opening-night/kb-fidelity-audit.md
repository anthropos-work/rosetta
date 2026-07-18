---
title: "KB Fidelity Audit — M226 opening-night (the live billion proof)"
date: 2026-07-17
scope: milestone:M226
invoked-by: build-mstone-iters (Phase 0b, iter-01 bootstrap tok pre-flight)
---

## Verdict
**GREEN** — proceed to the bootstrap tok strategy authoring.

M226 is the v2.4 acceptance closer ("nothing net-new"). Every KB doc it depends on was authored/reviewed
during M222–M225 (each of which ran its own KB-fidelity gate + close-milestone doc review). This audit
re-verifies those docs are still aligned with the tooling reality the live proof will consume.

## Topic Inventory

| Topic | Knowledge doc | Code paths / tooling | Status |
|---|---|---|---|
| Remote reach (billion deploy path, default-on) | `corpus/ops/demo/tailscale-serve.md` | rext `demo-stack/up-injected.sh` `derive_public_host_vars()`; `--no-public-host` opt-out | PAIRED |
| Latency gate (click→ACCESS p95<5s) | `corpus/ops/demo/latency-budget.md` | rext `stack-verify/e2e/run-latency.sh` | PAIRED (+ recruiter 3rd vantage = declared deliverable) |
| Auto-verify safety net | `corpus/ops/verification.md` | rext `stack-verify` (`verify live`) | PAIRED |
| Safety contract (exposure side, Part 3) | `corpus/ops/safety.md` | the 3-layer isolation guard + demo-published-on-0.0.0.0 disclosure | PAIRED |
| Coverage protocol (hiring vantage) | `corpus/ops/demo/coverage-protocol.md` | rext `stack-verify/e2e/` Playwright sweep | PAIRED |
| Demo-patch mechanism + the 4 hiring patches | `corpus/ops/demo/demopatch-spec.md` §5 | rext `demo-stack/patches/{next-hiring-role-remap,next-hiring-members-pagination,next-web-studio-url,next-web-public-website-url}` | PAIRED |
| Hiring read-model (is_hiring / 5 mgr / 45 cand / comparison cohort) | `corpus/services/hiring.md` | seeders `HiringConfigSeeder` + `HiringFunnelSeeder`; `stories.seed.yaml` 4th story | PAIRED |

## Fidelity Findings

1. **The 4 hiring demo-patches — ALIGNED.** `demopatch-spec.md` §5 (line 147+) documents the HIRING image
   (`build_frontend_hiring`) baking FOUR patches: the 2 net-new `apps/hiring` patches
   (`next-hiring-role-remap` on `UserStatusContext.tsx`, `next-hiring-members-pagination` on
   `InsightsContext.tsx`) + the 2 chained shared-`urls.ts` patches (`next-web-studio-url` →
   `next-web-public-website-url`). All four exist in `.agentspace/rosetta-extensions/demo-stack/patches/`
   and are wired in `up-injected.sh` (the hiring apply ladder at lines 784–889, chain rule intact:
   pubweb applies ON TOP of studio, same file). Verdict: **ALIGNED**.

2. **tailscale-serve.md default-on state — ALIGNED.** The doc's top banner (lines 8–21) correctly states the
   v2.3 M220 S3 flip: a bare `/demo-up N` is DEFAULT-ON for remote reach (D-DESIGN-3, superseding v2.2's
   D-DESIGN-1); opt out with `--no-public-host`. This is exactly the billion path M226's gate exercises
   ("default `/demo-up N`, NO flags"). The 0.0.0.0 disclosure + `safety.md` §3.1 cross-ref present.
   Verdict: **ALIGNED**.

3. **latency-budget.md vantage coverage — ALIGNED (2 vantages) + declared extension.** The doc gates
   p95<5s for employee + manager over 5 cold reset-to-seed runs, with the harness contract
   (`run-latency.sh`, measure the response BODY not headers, refuse a non-green stack, never gate on
   networkidle). M226 condition 5 extends this to a **3rd measured path (recruiter)**. This is NOT a blind
   area — `overview.md` explicitly declares `Delivers → ... latency-budget.md (a hiring 3rd vantage)`. The
   doc extension is a planned milestone deliverable. Verdict: **ALIGNED**; extension tracked as deliverable.

4. **hiring.md read-model — ALIGNED with exit-gate conditions 1–4.** Documents the `is_hiring` DUAL-WRITE
   (backend `public.organizations.is_hiring` + Clerk `publicMetadata.isHiring`), the 5-admins + 45-candidates
   population (Meridian Talent, `narrative: hiring`), and the comparison cohort semantics (comparable =
   shared `jobsimulation_id` + `organization_id`; score is NOT `jobsimulation.sessions.score`). Matches the
   gate's conditions 1 (org present + is_hiring + 5/45), 2 (≥40 comparable rows/position), 3 (candidate
   profiles), 4 (reads as hiring). Verdict: **ALIGNED**.

5. **rext code-of-record tag — ALIGNED.** Authoring copy `.agentspace/rosetta-extensions` at tag
   `casting-call-m225-harden` (HEAD be431c3); `stack-demo/rosetta-extensions` consumption at
   `casting-call-m225-sections` (live tag). Matches the milestone context's stated code-of-record.

## Completeness Gaps
None load-bearing. The recruiter 3rd-vantage latency fold-in is the only doc-side content M226 will add, and
it is already declared as a milestone deliverable (finding 3).

## Applied Fixes
None required — all load-bearing claims aligned. No inline edits.

## Open Items (require user decision)
None.

**Incidental (non-blocking, noted for awareness):** `latency-budget.md` uses a doc-local `D-DESIGN-1`
(in-page data-completion REPORTED-never-gated) that shares the label with `safety.md`'s superseded
`D-DESIGN-1` (public reach never default-on). Both are internally consistent within their own docs; the
shared label is potentially confusing but not a code-fidelity bug. Not a finding that blocks — noted only.

## Gate Result
**GREEN — proceed.** Every milestone-scope topic is PAIRED and every load-bearing claim ALIGNED. No blind
areas; no stale load-bearing claims. The bootstrap tok may author its strategy against verified knowledge.
