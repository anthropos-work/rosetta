**Type:** tik (under TOK-01 `reprove-corrected-hiring-on-billion`). Root-cause + fix the 3 iter-02 findings (F1/F2/F3)
in the rext seed tooling, then a default cold re-bring-up + warm re-measure on billion. Protocol: verification.md +
coverage-protocol.md + latency-budget.md.

# iter-03 — work log

1. **Root-caused all 3 findings to ONE class** (see `overview.md § Root cause`): M227 fix#1 left two Workforce-
   Intelligence dashboard seeders unguarded, wrongly assuming they "write no jobsimulation/mirror rows":
   - **FeedbackSeeder (F2/F3)** — writes `public.local_jobsimulation_sessions` MIRROR rows on GENERIC `refs.sims` sims
     (since v1.10 M42m) → leaked training sims + a 2nd session per hiring candidate into the recruiter's list.
   - **SuccessionSeeder (F1)** — writes `jobsimulation.interview_extraction_results` FKing JobsimSessionsSeeder's
     (now-skipped-for-hiring) sessions → FK violation (SQLSTATE 23503) → the whole seed reports "failed".
2. **Fixed:** both now consult `skipGenericActivityForHiringOrg` (2 guard lines). The fix#1 regression test
   `TestGenericActivitySeeders_SkipHiringOrg` gains the feedback + succession cases — **RED-proven** (42 + 20 leaked
   rows without the guard); GREEN with it. `hiring_scope.go`'s rule comment corrected. `go vet` clean; 13 stack-seeding
   packages GREEN.
3. **Committed + tagged:** rext authoring `1d97861`, tag **`casting-call-m228-hiring-scope-fix`**, pushed origin
   (main + tag). rext.tag SoT updated. 0 runtime change beyond the 2 guard lines (harden commit 78a3cb2 confirmed test-only).
4. **Default cold re-bring-up at the fixed tag** — teardown --purge (RC=0) + rext cutover → `casting-call-m228-hiring-scope-fix`
   (1d97861) + `up-injected.sh 1` (NO FLAGS). **IN PROGRESS at wind-down** (Directus-provision/seed phase).

## Re-measurement — DB/seed-level VERIFIED; UI render re-measure PENDING (resume step)

**The default cold re-bring-up at `casting-call-m228-hiring-scope-fix` came up GREEN (UP_RC=0, autoverify OK).** All 3
findings are VERIFIED FIXED at the DB + seed-log level (measured from this Mac against billion's demo postgres):

| finding | iter-02 (broken) | iter-03 (fixed) | verdict |
|---|---|---|---|
| **F1** succession FK seed error | `succession rows=0 ERROR FK violation → "seed failed"` | **`succession rows=165 ok`** — clean seed, no "seed failed" | ✅ FIXED |
| **F2** hiring-only | 7 HIRING + **2 TRAINING** sims (leak) | **SIMULATION_TYPE_HIRING ONLY, 5 sims, 42 sessions** | ✅ FIXED |
| **F3** 1-sim/candidate | 26×1 + **17×2** sims | **1 sim for ALL 42 candidates**; per-position **8,8,8,9,9** (5 positions, all ≥6) | ✅ FIXED |
| **C1** counts | 5+45 | **5 admins + 45 candidates** (holds) | ✅ |

`hiring org set-dressed: 5 shared positions + 42 candidate HIRING sessions` (was 62 — the ~20 leaked feedback-mirror
sessions gone).

**PENDING (resume step):** the WARM UI render re-measure of C2 (recruiter list renders hiring-only ~8/position),
C3 (Cara/Cody/Rae usable), C4 (reads-as-hiring), C5 (recruiter p95 < 5 s) from this Mac — cold-tailnet-slow, launched
best-effort at wind-down (see `render-iter03.log`), warm-before-gate per M226 F5. **This iter is NOT yet closed** (no
close section until the UI render re-measure lands). Then iter-04 = a 2nd clean cold cycle for reproducibility → gate.
See the journal RESUME block.

## CLOSE — iter-03 closed-fixed: the WARM UI render re-measure LANDED (7/7 on the cold-seed billion instance)

The pending WARM UI re-measure completed from this Mac against the billion cold-reset-to-seed instance
(tag `casting-call-m228-hiring-scope-fix`), after a render-probe hardening pass (rext
`casting-call-m228-render-probe-fix`, commit `0670a47`). **All 7 conditions MET:**

| # | condition | result | how |
|---|---|---|---|
| 1 | hiring org, is_hiring, 5 mgr + 45 cand | ✅ | cockpit "Meridian Talent · 50 people"; DB verify (iter-03) |
| 2 | ≥6 non-junk rows/EACH of 5 positions, 1 sim/candidate | ✅ | **per-sim gated proof 5/5 PASS**: Talent-Mgr 8, BD-Lead 8, Inside-Sales 9, Project-Mgr 9, AWS-Security 8 (=42), each ≥ floor 6, junk=0 |
| 3 | 2 candidate heroes usable + external emails + gender-matched avatars | ✅ | m224-candidate-heroes 3/3 PASS (Cara assessed /home usable, Cody assigned-only usable, Rae recruiter scoreboard — none ejected); external emails + matched avatars visible in every drawer |
| 4 | reads as hiring, hiring-only content | ✅ | all 5 sims SIMULATION_TYPE_HIRING; recruiter scoreboard 100% HIRING-tagged; 0 training/assessment |
| 5 | recruiter p95 click→ACCESS < 5 s | ✅ | **p95 1.27 s** (p50 0.45 s), 5/5 reached ACCESS |
| 6 | coexists with 3 workforce orgs | ✅ | cockpit: Cervato Systems + Northwind Aviation + workforce story orgs + Meridian Talent |
| 7 | 0 platform-repo edits | ✅ | all changes in rext + corpus; the harness fix is a test-probe edit |

**The C2 render breakthrough:** the cold-tailnet all-5-drawer render kept false-zeroing sims 2-5. NOT a data
problem — the app's own list API returned all 5 HIRING sims at 8,8,9,9,8, and per-sim screenshots showed clean
external-email / 1-attempt drawers. Root cause: the recruiter comparison drawer is a **Next.js intercepting route**
— only the FIRST sim clicked per page-load reliably mounts as `.ant-drawer` + fires its client member POST + is
DOM-visible. Fixed the probe (RENDER_ONLY_SIM → prove each sim as "the first"; `:visible` scoping; DOM-authoritative
poll; env-tunable per-drawer budget) → clean automated 5/5.

**Reproducibility (the gate's "reproducibly" clause) → iter-04:** iter-03 is ONE cold cycle (GREEN bring-up +
DB-verify + 7/7 UI). iter-04 = a 2nd independent cold reset-to-seed for reproducibility, then the gate fires.
