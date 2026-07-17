# M226 — Spec notes

_Iterative milestone: this file accumulates iteration-protocol-specific technical notes (live-run transcripts,
per-condition evidence, latency measurements). Per-iter detail lives in `iter-NN/`._

## The 7-condition live billion proof
_Per-condition evidence: org present + is_hiring=true + 5/45 counts; ≥40 rows per 5 positions; candidate profiles;
reads-as-hiring; recruiter p95 click→ACCESS; coexistence with 3 workforce orgs; 0 platform edits._

## Latency — the recruiter 3rd vantage
_p95 click→ACCESS measurement over the tailnet origin; the latency-budget.md fold-in._

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
