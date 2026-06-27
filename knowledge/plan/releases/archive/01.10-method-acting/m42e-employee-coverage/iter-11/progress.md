# iter-11 progress

**Type:** tik (iter_shape: baseline — read-only, live demo-3). Active strategy: **TOK-10**.

P0 of the design-plan: re-confirm the before-state of the P1–P3 pages on live demo-3 (DB probes only, no fix).

## Phase A — measure (DB probes, demo-3)
- Found the hero: Maya Chen `d26e0467-31d9-5f73-96ad-525dfbb38d11`, `maya.chen1@cervato-systems.com`.
- P1: 20/30 verified skills are flat-pool junk (15Five/3D-Dental/3dcart/…); 10 are role-coherent. Same junk in
  the 60-skill claimed tail. → B1.
- P2: timeline rows EXIST (3 work + 1 edu) with real skill NAMES, but the `skills` json is a bare array
  `["Java",…]` (the LegacySkills shape mismatch → skeleton). → B1.
- P3: 0 personal_assignments, 0 completed skillpath sessions, 0 user_bookmarks for Maya. → B1.
- Specialized-role analysis (B2): every role caps at 10 skills; CHOSE **Backend Software Engineer** (coherent
  modern senior-backend stack, continuity with current skills). Surfaced for user confirmation.
- Curated claimed-pool (B3): 27 software skill names resolve in the real taxonomy (no-fabrication-safe).

## Close — 2026-06-25
**Outcome:** P0 baseline recorded (B1); specialized role chosen = Backend Software Engineer (B2); curated claimed
allow-list resolved to 27 real node-ids (B3). No code change (read-only).
**Type:** tik (baseline)
**Status:** closed-fixed
**Gate:** N/A (read-only baseline)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik #1) — (6) protocol-stop: n — Outcome: continue (into iter-12 P1)
**Decisions:** B1 (before-state), B2 (role choice), B3 (curated allow-list resolution) — see ./decisions.md.
**Routes carried forward:** P1 (iter-12 — skill draw + role specialize), P2 (iter-13 — json shape), P3 (iter-14 — activity).
**Lessons:** every public job_role = exactly 10 role-skills; specialization adds coherence + title, not depth. The depth is the curated allow-list top-up.
