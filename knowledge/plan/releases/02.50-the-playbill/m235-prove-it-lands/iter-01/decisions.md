# iter-01 — decisions (local)

_(The strategy record TOK-01 lives in the milestone-root `decisions.md`; these are iter-local notes.)_

- **KB-1 (from Phase 0b):** `stack-seeding/contentsession/fixture/content-sessions.yaml` header comment still
  describes the SUPERSEDED synthesize-first "provably PII-free" posture. The authoritative corpus doc
  (`session-clone-spec.md`) correctly documents copy-real+scrub. Fix routed to the first fixture-extension tik
  (which edits that exact file) rather than applied in the audit — a rext code change belongs in an iter commit.

- **Q1 resolution (overview open question):** the sim result page is a persisted read (M231), so a seeded
  fan-out renders it directly — the composed-outcome `/profile/activities` fallback the question floated is
  UNNECESSARY, and no demo-patch is authorized for the sim result page (only the two interview flag-gate
  demopatches, already built in M232).

- **Q2 resolution:** not-passed renders a meaningful result page (terminal `evaluation_status` + persisted
  sub-threshold score), not blocked/empty — so the "each type in passed AND not-passed" gate is renderable.
