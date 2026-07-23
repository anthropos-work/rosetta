---
iter: 25
milestone: M244
iteration_type: tik
status: closed-fixed-partial
created: 2026-07-23
---

# iter-25 — gate (c) final push: recover demo-1 + root-cause & fix the 4 ai-readiness Playthroughs

**Type:** tik (run 9, tik 1)
**Active strategy reference:** TOK-03 (HOLD TOK-01/02 + sequence the narrow final push). TOK-03's step (3) —
"gate (c) tick — the 16 Playthroughs LAST on pt-world → ticks 8/8 = GATE MET." The stack-verify half is done
(iter-18); 12/16 Playthroughs are green cold (iter-23); ONLY the 4 ai-readiness Playthroughs remain.

## Step 0 — re-survey (mandatory)
Metric is **7/8** (gate a/b/d/e/f/g/h all discharged; iter-24 closed gate f). The single remaining gate part is
**(c)**, and within it the single remaining item is the **4 ai-readiness Playthroughs** (2 member + 2 manager,
`ai-readiness.yaml`). Target confirmed still-valid and untouched. billion demo-1 is DOWN (0 containers, all 7
serve rules persist, 29 images cached, 5.8GB RAM free) — recovery is a precondition.

## Cluster / target identified
gate (c) 16/16. The 4 ai-readiness Playthroughs fail on a surface-render zero-state. iter-23 routed this
forward with a **correlation** (`ai_readiness_live_snapshots=0`) — but a read of the platform source
(`app/internal/workforce/ai_readiness.go` + `live_snapshots.go`) shows the manager dashboard live path
(`buildLiveResponse` → `computeOrgBreakdowns`) does **NOT read `ai_readiness_live_snapshots`** (that table is
read only by the Talk-to-Data askengine + the prune step). So the correlation is likely not the cause. The
true cause must be diagnosed LIVE.

## Hypothesis
The active-cycle live read renders a zero-state for one of: (a) `keepStartedMembers` filters out every seeded
member (none classified stage≥1 in the active cycle); (b) the active cycle isn't resolved by the dashboard
(cycle-picker / options); (c) `ai_readiness_diagnose_narrative` empty → manager UC2 interview-findings section
absent; (d) the member `/home` surface reads a per-user path that the pt-world seed doesn't populate. The fix
is SEED-side in `rext stack-seeding` / the pt-world seed (0 platform edits), analogous to the existing
`appendSnapshot` frozen-snapshot seed.

## Expected lift
gate (c) 16/16 → metric **7/8 → 8/8 = GATE MET**. (Binary-per-gate: ticks only when all 16 Playthroughs green.)

## Phase plan (playthroughs.md + verification.md diagnostic discipline)
1. **Recover** demo-1 on billion (`up-injected.sh 1 --public-host billion.taildc510.ts.net`, cached images) →
   fresh-green autoverify + web serving.
2. **Seed + diagnose**: pt-world reset-to-seed → run the 4 ai-readiness Playthroughs from the peer (networkidle
   fix already in, rext dddef18) → capture the true failure + live-DB probe which read path fires and what it
   returns (the diagnostic-probe step — do NOT inherit iter-23's correlation).
3. **Fix** the true cause SEED-side (rext stack-seeding / pt-world seed), commit + tag-move + push + re-pin
   billion's rext clone as needed, re-seed pt-world.
4. **Re-measure**: re-run the 4 → gate (c) 16/16.

## Escalation conditions
- If the true cause requires a PLATFORM edit → STOP, SEVERITY=blocker (do not edit the platform).
- If genuinely unseedable from a demo (M231 unseedable-surface class) → route honestly with a documented
  falsification, not a fabricated pass.

## Acceptable close-no-lift outcomes
A complete live diagnosis that characterizes the true root cause + a documented falsification (if the fix can't
land this iter) satisfies the protocol even if the metric doesn't tick — provided the diagnostic-probe step
actually ran live on billion.
