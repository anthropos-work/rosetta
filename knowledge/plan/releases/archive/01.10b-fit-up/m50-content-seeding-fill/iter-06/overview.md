---
iteration: 06
iteration_type: tik
status: closed-fixed
created: 2026-06-30
---

# iter-06 — languages + cert coverage + manager-manifest strengthening (the D4/F1 reconciliation)

**Type:** tik (under TOK-01).

## Re-survey (Phase 1 Step 0)
TOK-01's run-1 close reached `(0,0)` BOTH vantages on warm demo-1 — but the cross-check vs the iter-01
diagnosis showed the manager gate passed **BLIND** to two of M50's OWN annotation gaps (languages + certs),
because the coverage manifest doesn't ASSERT them. Re-measured the live demo-1: `world_languages`/
`user_languages`/`membership_languages` all exist but **0 rows**; `user_certifications` = **9** (hero-only) across
**340 members**. The targets are still genuine + untouched. NOT a re-scope — TOK-01's "fill the genuine empties,
sweep-driven, strengthen-the-manifest-to-prove-it" strategy still holds; this tik narrows it to the run-2 target.

## Active strategy reference
TOK-01 (seed-fill the genuine empties, sweep-driven, re-seed to iterate) — the manifest-strengthening half is
the D4/F1 reconciliation TOK-01's rationale + the iter-02/05 closes already routed forward.

## Cluster / target identified
The manager Talent-tab annotation complaint ("no language spoken + Certification are really low numbers") +
the `/enterprise/members` Location column the iter-02 fill renders but the manifest doesn't assert.

## Hypothesis
(1) A new `MemberLanguagesSeeder` (world_languages catalog + per-member `user_languages` → membership_languages
via the DB trigger) + extending `CertificatesSeeder` to ~45% role-coherent member coverage fills the Talent-tab
data. (2) Strengthening the manager manifest to ASSERT the members Location column + the Talent-tab languages/cert
content makes the gate PROVE the gaps are closed (no longer blind). Then the manager sweep is GREEN truly proving
the gaps.

## Expected lift
Manager `(failingSections, escapes)` stays `(0,0)` on a STRENGTHENED manifest (the new assertions PASS because the
data now renders) — the gate is MET only when the strengthened manifest passes, not the old blind one.

## Phase plan
A (calibrate the manager Talent-tab + members render) → C (seed code + re-seed demo-1 + manifest strengthening +
harness `preAssert` tab-click capability) → D (manager re-sweep with the strengthened manifest) → E (close).

## Escalation conditions
A surface that can't be asserted without a platform edit → re-scope (the zero-edit line). A Talent-tab that
renders languages/certs only as a chart with no assertable text/cardinality → author the assertion against the
chart's real DOM (recharts SVG), never weaken to a blind pass.

## Acceptable close-no-lift outcomes
If the calibration proves the Talent tab does NOT surface languages/certs at all (a platform gap, not a seed gap),
record the falsification + route the manifest assertion accordingly — but the seed fill still lands (the /profile +
data-density value), and the members-Location strengthening still lands.
