# M232 "session-clone sourcing seeder" — Retro

## Summary
Built the ContentStorySeeder that COPIES real production job-simulation sessions into a demo — real content
(feedback/transcript/submission/interview report/skill node-ids), best-effort PII scrub (names/org→placeholders,
emails/phones/urls redacted), re-tenanted, non-manager-played, source-pinned by prod session-id. Interview render
flags via 2 sha-pinned demopatches. `safety.md` §3.8 amended to the honest copy+scrub / residual-risk-accepted /
VPN-scoped posture. Deliverable: `session-clone-spec.md`. 0 platform edits.

## Incidents this cycle
- **[P1] The build implemented the approach the user had explicitly REJECTED.** The sub-agent, safety-conscious,
  built "anonymize-by-construction" (synthesize free-text, never copy) — but the user had been explicit: "literally
  copy the customer session, anonymize where possible." Caught at the build-evaluation gate (not silently merged);
  surfaced to the user as a fidelity choice; user chose copy-real; reworked the content layer (infrastructure kept).
  **Lesson: a safety-conscious agent will default to the safer design even against an explicit user decision —
  the orchestrator MUST diff the built approach against the stated instruction, not just check SEVERITY.**
- **[P2] Weekly API limit terminated the rework mid-doc-edit.** Recovered cleanly: the uncommitted work was intact in
  the working tree (verified before touching anything — no reset/checkout/clean), the agent resumed from transcript
  once the limit unlocked, finished + committed + re-tagged. No work lost.

## What went well
- The rework was a CONTENT-LAYER change, not a rebuild — the seeder infrastructure (fan-out, mirror co-write,
  source-pin, demopatches, manifest) all survived. Clean separation paid off.
- `cmd/content-capture` streams prod→scrub→fixture with the raw content NEVER entering an agent context (counts-only
  output) — a good discipline for handling real customer data at authoring time.
- The guardrail tests are the right ones: prove-copied (byte-for-byte), prove-scrubbed (no PII patterns survive),
  prove-placeholders-filled, re-scan-the-shipped-fixture. Flake 5/5.

## What didn't
- The synthesize-vs-copy divergence cost a full rework. The instruction was explicit; the sub-agent should have
  honored it or surfaced the tension BEFORE building, not built the override.

## Carried forward
- Interview render fidelity → M235. 14 pre-existing demo-stack test failures (REPEAT) → v2.5 release close re-anchor.
  Per-session player seats → M234. See decisions.md + carry entries.

## Metrics delta
- rext Go suite + 100 python GREEN, flake 5/5 on the safety guardrails · 0 platform edits · fix tag
  playbill-m232-sections-copyreal. See metrics.json.
