---
iter: 11
iteration_type: tik
iter_shape: baseline
status: closed-fixed
---

# iter-11 — P0 baseline confirm (read-only, live demo-3)

**Type:** tik (iter_shape: baseline — read-only measurement). **Active strategy: TOK-10**
(persona-believability-by-root-cause). Per the design-plan's P0: re-confirm on live demo-3 the before-state of
the pages P1–P3 will fix, via authenticated DB probes (`docker exec demo-3-postgresql-1 psql`).

## Cluster / target identified
P0 is the baseline anchor for the P1–P3 root-cause fixes. TOK-10's next-tik direction named exactly this:
record the before-state of /home "Latest skills" + /profile/skills (P1), the /profile timeline skeleton (P2),
and /home Paths/pills + /profile/activities (P3) so each fix's before→after is provable on re-seed.

## Hypothesis / expected lift
No metric lift (read-only baseline). The deliverable is the recorded before-state + the resolved
specialized-role choice for P1 (the user-facing character decision the design-plan defers to this run).

## Phase plan
Phase A (measure) only — DB probes against demo-3. No fix, no re-seed.

## Close — 2026-06-25
**Outcome:** P0 baseline recorded; all 4 roots (P1 skill-draw, P2 json-skeleton, P3 activity) confirmed live on
demo-3; the specialized-role candidate set queried + the curated software claimed-pool resolved (27 real
node-ids). Maya hero = `d26e0467-31d9-5f73-96ad-525dfbb38d11`, `maya.chen1@cervato-systems.com`.
**Type:** tik (baseline)
**Status:** closed-fixed
**Gate:** N/A (read-only baseline; the semantic harness that grades the new gate is P7)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik #1 of session) — (6) protocol-stop: n — Outcome: continue (into iter-12 P1)
**Decisions:** see ./decisions.md (B1 baseline before-state; B2 specialized-role candidate analysis).
**Routes carried forward:** P1 (iter-12), P2 (iter-13), P3 (iter-14).
**Lessons:** every public job_role caps at exactly 10 role-skills (the taxonomy is 10-core-per-role) — so role
specialization does NOT add skill DEPTH; it adds COHERENCE + a believable senior title. The depth (~12 verified
+ ~18 claimed) comes from the curated allow-list top-up, NOT the role. Confirms the design-plan's note.
