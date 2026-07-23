**Type:** tik — gate (g) full green, under TOK-01. Run 4, tik 1.

# iter-06 — progress

## What landed (gate g: the plan-section-id alignment assertion, S-8/S-9)

**The assertion (rext `contentsession/`, committed + green + pushed — 0dec8ff, tag `sound-check-m244-content-sweep-robustness` → 70d75cb on origin):**
- `TestInterviewPlanSectionAlignment` — per interview content-story, asserts the report DATA section keys ⊆
  the captured PLAN section-ids of their render scope (`manager_report.results` ⊆ plan.manager;
  `user_report.results` ⊆ plan.player), plan non-null with ≥1 section, non-empty manager overlap. Reads
  section-id KEYS only, never a report value (PII-safe).
- `fixture/interview-plan-sections.json` — the golden PLAN section-ids per scope for the sole public interview
  sim (`ai-readiness-interview-d62`, plan **v1.4 +v1.4-mastery-axes**: 12 manager / 1 summary / 1 player),
  read read-only from prod (non-PII sim-def metadata).
- `TestOrphanKeysDetectsV13Drift` — teeth: the exact v1.3 drift keys (breadth/context_fit/frequency) MUST flag
  as orphans against the v1.4 plan.

**The real defect it CAUGHT + FIXED.** The assertion is not vacuous — it found a live misalignment:
`intv-voice-fail` was pinned to session `43a92fc0`, whose interview report was produced under an **older plan
version (v1.3: breadth/context_fit/frequency axes)** than the sim's current captured plan (v1.4). Its 3 orphan
keys were LOST on render; 5 v1.4 sections rendered empty. **FIX:** re-pinned `43a92fc0 → 05dae0f7` — same
sole public interview sim, a genuine FAILED interview (completion_status=failed, score 0), whose report is
**v1.4-clean** (11 in-plan manager keys, 0 orphans). Re-captured via `content-capture --only intv-voice-fail`
(read-only prod, scrubbed); cleanliness (surviving-token) green; both projection goldens
(`content-manifest.json` + `seed-generation-manifest.yaml`) regenerated + green.

## Live confirmation on billion (non-flaky, DB-level)
billion demo-1 (m243 seed) has the plan (`directus.simulations_extraction`: 14 sections, from iter-05's load)
+ 167 seeded `interview_extraction_results`. Direct DB alignment proof of the two seeded interview clones vs
the captured plan:
- **`97f3f681`** (intv-voice-**pass** clone / aligned session cba53b09): 12 manager keys, **0 orphans** →
  aligned live. Combined with iter-05's proven shell render (0→520), its manager sections render aligned.
- **`ace3b54e`** (intv-voice-**fail** clone / the m243-seeded OLD 43a92fc0): 10 manager keys, orphans =
  **{breadth, context_fit, frequency}** — the EXACT drift the assertion catches, live. The iter-06 re-pin
  fixes it (fixture + manifest now point at `05dae0f7`, 0 orphans); it lands on the next cold reset-to-seed
  (folded into gate h, which re-pins billion to the m244 tag).

## Full stack-seeding module: green (go test ./... exit 0), go vet clean. 0 platform edits.

## Close — 2026-07-22

**Outcome:** gate (g) DISCHARGED — the interview plan-section-id alignment assertion added + green + pushed;
it caught a real v1.3→v1.4 plan-version drift and the fix (intv-voice-fail re-pin to a v1.4-clean session)
landed; live-confirmed on billion (pass clone 0-orphans aligned; the assertion demonstrably catches the
m243-seeded fail clone's live drift, which the re-pin resolves on the cold reset-to-seed). Metric **2/8 → 3/8**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (milestone gate a–h not fully met; **3/8** — gate part (g) discharged this iter)
**Phase 5 grading:** (1) gate-met: n (3/8) — (2) triggered-tok: n (iter-03/05 made measurable progress) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik 1 of run 4) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (assertion at the contentsession layer + golden plan), D2 (re-pin as the real fix, not an allowlist), D3 (DB-level live alignment proof over a flaky auth click) — iter-06/decisions.md
**Side-deliverables:** none (all in gate-(g) planned scope; 0 platform edits).
**Routes carried forward:**
- (Fate-2, covered) the visual "Explore Key Moments" deep-section click-through render — re-driven by the
  gate-(b) content-stories LANDS sweep + gate-(h) cold reset-to-seed, which also RE-SEED the re-pinned
  intv-voice-fail (resolving ace3b54e's live drift). Named handler: gate-(b)/(h) live sweep.
- (Fate-2, covered) gate (a) ORG-CLEAN LIVE re-verify for the re-pinned `05dae0f7` fixture at the gate-(h)
  cold reset-to-seed (offline cleanliness already green in go test).
**Lessons:** (1) a "renders empty / renders partial" interview surface can be PLAN-VERSION DRIFT — a
source-pinned session whose report predates the sim's current plan — not just a null plan; check report-data
keys ⊆ captured-plan section-ids, per scope. (2) content-capture `--only <key>` is the clean re-pin
mechanism (read-only prod, scrub, cleanliness-gated); pair it with regenerating BOTH honesty-gated
projection goldens. (3) a DB-level alignment query on the live demo is a stronger, non-flaky substitute for
a flaky authenticated deep-render click.
