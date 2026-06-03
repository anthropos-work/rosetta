**Type:** tok (bootstrap)

# iter-01 — bootstrap tok

Authored the milestone's initial strategy from the M1 overview + spec-notes + the alignment protocol.

## Close — 2026-06-03

**Outcome:** TOK-01 authored — *mirror-by-score, easy-side-first*: author the Clerk DNA → build the offline-capturable authn twin to 100% critical → then the live-SaaS orgclient twin → inject via `go.mod replace`. Identified the platform-consumed Clerk Go surface (authn provider claims + 10 orgclient methods) and surfaced the milestone's central open decision (D1): the orgclient golden source for the live Clerk API.
**Type:** tok (bootstrap)
**Status:** closed-fixed (the bootstrap tok's deliverable — the initial strategy — landed)
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this is a *bootstrap* tok, not triggered) — (3) re-scope: n — (4) user-blocker: **y** (iter-02's planned first step — stand up the `clerkenstein` workspace + author the DNA + run the first measurement — cannot land until D1 is decided: the orgclient golden source + whether creating `anthropos-demo/clerkenstein` + cloning the pinned Clerk SDK is available/in-scope in this environment) — (5) cap-reached: n — (6) protocol-stop: n — **Outcome: exit-4 (user-blocker, surfaces at the iter-02 boundary)**
**Decisions:** TOK-01 (milestone-root `decisions.md`); D1 open decision (same file).
**Side-deliverables:** none.
**Routes carried forward:** iter-02 (tik) — stand up `anthropos-demo/clerkenstein` + author the Clerk Alignment DNA via `/align-dna`; blocked pending D1.
**Lessons:** the iteration gate is an alignment score, which requires source goldens; for a live-SaaS source (Clerk's orgclient API) the golden-source approach is a real upfront decision, not an implementation detail — surface it before the first tik rather than discovering it mid-build.
