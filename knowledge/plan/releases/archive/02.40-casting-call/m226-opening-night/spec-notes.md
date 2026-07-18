# M226 — Spec notes

_Iterative milestone: this file accumulates iteration-protocol-specific technical notes (live-run transcripts,
per-condition evidence, latency measurements). Per-iter detail lives in `iter-NN/`._

## The 7-condition live billion proof
_Per-condition evidence: org present + is_hiring=true + 5/45 counts; ≥40 rows per 5 positions; candidate profiles;
reads-as-hiring; recruiter p95 click→ACCESS; coexistence with 3 workforce orgs; 0 platform edits._

**iter-02 (first measurement, cold reset-to-seed, default no-flags, ~10.5 min GREEN bring-up):**
- C1 ⚠ Meridian Talent, is_hiring=true, 50 total, but **3 admin + 47 candidate** (≠ 5+45; Finding-2, hero-slot displacement).
- C2 ⛔ data GREEN (294 sessions / 5 sims) but render unreachable (:13001 not served; Finding-1).
- C3 ⛔ candidates land on :13001 (Finding-1).
- C4 ✅ is_hiring=true (DB) + cockpit "Hiring" label + roster org_is_hiring.
- C5 ⛔ recruiter lands on :13001 (Finding-1); harness recruiter vantage prepared.
- C6 ✅ cockpit shows 4 orgs (2 workforce + Northwind + Meridian-hiring).
- C7 ✅ 0 platform edits (cms `?? studio/` = disclosed M221 D-05h).
- **Finding-1** (:3001 not tailscale-served) fix committed+tagged `casting-call-m226-serve-hiring` (rext ee1bdf2) → apply+prove iter-03.
- **Finding-2** (3+47 vs 5+45) routed to iter-03.

## Latency — the recruiter 3rd vantage
_p95 click→ACCESS measurement over the tailnet origin; the latency-budget.md fold-in._

**iter-03 (from this Mac → billion, HTTPS, 5 cold runs, gated on fresh-green autoverify.json):**
- **recruiter (`rae-recruiter` → apps/hiring `/enterprise/activity-dashboard`): reached ACCESS 5/5, p50 0.44 s,
  p95 1.50 s < 5.0 s gate.** The 3rd measured access path (v2.3: employee 1.46 s / manager 1.40 s).
- Harness: `run-latency.sh 1 recruiter` — the new `recruiter` vantage (rext `casting-call-m226-serve-hiring`,
  ee1bdf2). measureLogin is vantage-agnostic; the recruiter's cockpit CTA lands on the hiring app (:3001+offset),
  and the same ACCESS predicate (loader-gone + hero identity in header/nav) is satisfied by the hiring shell.
- **Prerequisite (Finding-1):** the hiring app (:3001+offset) had NO `tailscale serve` HTTPS front → the recruiter
  was unreachable from the peer until `gen_tailscale_serve.py` was taught `("hiring", 3001)`. The latency-budget.md
  fold-in should note that the recruiter vantage requires the hiring port to be fronted (a demo-path prerequisite).

## Demo-patch re-prove-at-final-code
_Whatever M224 pinned, re-proven live (the M221 discipline); any live-only perf pin._

## Pre-flight audits — iter-01
**KB-fidelity (Phase 0b, 2026-07-17): GREEN.** Report: `kb-fidelity-audit.md`. Every milestone-scope topic
PAIRED + every load-bearing claim ALIGNED. Topic→doc→tooling triples verified:
- Remote reach → `tailscale-serve.md` → `up-injected.sh derive_public_host_vars()` (default-on, M220 D-DESIGN-3).
- Latency gate → `latency-budget.md` → `stack-verify/e2e/run-latency.sh` (2 vantages; recruiter 3rd = declared deliverable).
- Demo-patch/hiring → `demopatch-spec.md` §5 → the 4 patches (`next-hiring-role-remap`, `next-hiring-members-pagination`, chained `next-web-studio-url`→`next-web-public-website-url`), all present in `demo-stack/patches/` + wired in `up-injected.sh`.
- Hiring read-model → `hiring.md` → `HiringConfigSeeder`+`HiringFunnelSeeder`, `stories.seed.yaml` 4th story (is_hiring dual-write, 5 mgr/45 cand, comparison cohort).
- rext code-of-record: authoring `casting-call-m225-harden`; consumption `casting-call-m225-sections`.
- Incidental (non-blocking): `latency-budget.md`'s doc-local `D-DESIGN-1` shares its label with `safety.md`'s superseded `D-DESIGN-1` — noted, not a fidelity bug.
