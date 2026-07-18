---
iter: 04
milestone: M226
iteration_type: tik
status: closed-fixed
date: 2026-07-17
---

# iter-04 (tik-3) — fix Finding-2 (C1 count → 5+45) + reproducible DEFAULT cold re-bring-up (7/7)

**Active strategy reference:** TOK-01 `reprove-hiring-on-billion` (iter-03's close routed this).

**Step 0 — re-survey:** iter-03 left 6/7 GREEN; only C1 (3 admin + 47 candidate vs exactly 5+45) open, plus the
gate's reproducibility clause (prove the DEFAULT bring-up fronts :13001 + counts hold, no hand-holding). Both fresh.

## Cluster / target identified
Close C1 (Finding-2) with a seeder-preset fix, then prove the whole gate GREEN on a **default** cold bring-up.

## Root cause (confirmed in code)
`users.go roleForHero` is deliberately vantage-faithful: a hiring end-user hero → `candidate` (the
`HiringFunnelSeeder` skips non-candidates, so Cara/Cody MUST be candidates for C2/C3). Heroes ride the first
population slots (M38-D7). `roleForIndex` makes the first `int(size*admin)`=5 slots admin. So Rae(1)=admin,
Cara(2)+Cody(3)=candidate (overriding 2 admin-band slots) → **3 admin + 47 candidate**. The preset's own comment
claims "→ 5+45" — it doesn't achieve its stated intent.

## Fix (one-line preset, rext `stack-seeding`)
`stories.seed.yaml` hiring story `role_mix.admin: 0.1 → 0.14` (+ a documenting comment). adminCount = int(50×0.14)
= 7 admin-band slots; minus Cara+Cody overridden to candidate → **5 real admins (Rae + 4 fill) + 45 candidates**.
Safe: no hero moves (identities stable), no role changes (Cara/Cody stay candidates), no code, no invariant touched
(hiring_config/funnel use the FIRST admin = slot 1 = Rae, unaffected).

## Hypothesis
The preset fix yields exactly 5 admin + 45 candidate; a default cold re-bring-up at the fixed rext tag fronts
:13001 in the default path (Finding-1 fix) + seeds 5+45 (Finding-2 fix); all 7 conditions measure GREEN.

## Expected lift
6/7 → **7/7 GREEN**, reproducibly on a default cold reset-to-seed. Watch C2's margin: 47→45 candidates may drop
per-sim rows from 44 to ~42 (must stay ≥ 40).

## Phase plan
1. Preset fix (`role_mix.admin 0.1→0.14`) in the authoring copy + commit + tag (`casting-call-m226-count-5-45`).
2. Consume the tag on billion.
3. Full **DEFAULT** cold re-bring-up: `rosetta-demo down 1 --purge` → `up-injected.sh 1` (NO FLAGS) — proves
   Finding-1's serve fix + Finding-2's counts in the DEFAULT path.
4. Re-measure ALL 7 from this Mac → target 7/7 GREEN.

## Escalation conditions
- If the preset fix does NOT yield exactly 5+45 (e.g., float/rounding), attribute + adjust.
- If reducing candidates 47→45 drops any C2 sim below 40 → attribute (funnel take-rate) + adjust.
- An un-patchable surface → ESCALATE.

## Acceptable close-no-lift outcomes
If the re-bring-up surfaces a new gap characterized with falsification, that's a complete cycle → closed-no-lift.

## Close
See `progress.md`.
