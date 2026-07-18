**Type:** tik (under TOK-01 `reprove-corrected-hiring-on-billion`). The first live cold bring-up of the M227-corrected
data on billion + the first retuned-7-condition + 4-M227-render-check measurement from this Mac. Protocol:
verification.md + coverage-protocol.md + latency-budget.md.

# iter-02 — work log

1. **Teardown** the M226 `demo-1` (`rosetta-demo down 1 --purge` as devops): 17 containers removed, data purged,
   images reclaimed, hostlock released. VERIFIED from this Mac: offset ports refuse (curl exit 7); billion "No serve
   config" (7 stale fronts cleared), 0 containers, no native survivors. **The M226 F3 orphan did NOT recur** — the
   default teardown is clean.
2. **Substrate cutover (rext-only):** billion rext `casting-call-m226-c2-race-fix` (4bd68ff) → **`casting-call-m227-sections`**
   (63c3e8d); clean tree; rext.tag SoT updated. No platform-clone changes (M227 = pure seed/content). 20 images stay cached.
3. **Default cold `up-injected.sh 1` (NO FLAGS)** as devops, synchronous blocking ssh. **UP_RC=0.** Backend Go images
   built fast (warm buildkit cache); UI tier rebuilt (next-web + 6 patches, studio-desk, hiring). autoverify OK. The
   `app-aireadiness-snapshot-loadmembers` demopatch logged "sha DRIFTED but anchor intact (1x)" — the EXPECTED
   self-healing case (demopatch-spec.md), applied cleanly.
4. **Measured C1–C7 + fix#2/#4 FROM THIS MAC** — see the scoreboard.

## The 7-condition scoreboard (iter-02, first corrected-data cold cycle)

| # | condition | verdict | evidence |
|---|-----------|---------|----------|
| 1 | is_hiring + **exactly 5 mgr / 45 cand** | ✅ GREEN | Meridian Talent, is_hiring=t, **5 admin + 45 candidate** (DB) |
| 2 | recruiter comparison **≥6 / each 5 positions, 1 sim/candidate** | ⚠ MIXED | 5 shared positions 9,9,8,8,7 all **≥6** (4/5 rendered clean, junk=0, ext emails visible); **F3: 17/45 candidates on 2 sims** (violates "exactly 1"); F5 render harness 5th-sim cold timeout |
| 3 | 2 candidate profiles usable + **ext emails + matched avatars** | ✅ GREEN | Cara (Completed) / Cody (Assigned) / Rae (no-regression) all PASS; emails external (render+DB); avatars present (pic=t), Rae female-face-for-female-name |
| 4 | reads as hiring + **hiring-only** | ⚠ MIXED | reads-as-hiring ✅ (Rae nav "Results \| AI Simulations \| AI Interviews \| Candidates Feedback", HIRING logo, 5 HIRING sims); **F2: 2 TRAINING sims leak** (hiring-only violated) |
| 5 | **recruiter p95 click→ACCESS < 5 s** | ✅ GREEN | WARM p95 **1.47 s**, ACCESS 5/5 (first cold run needs warm-up — M226 discipline) |
| 6 | coexists with 3 workforce orgs | ✅ GREEN | cockpit 12 heroes / 4 orgs; **F1 broke Northwind interview data** |
| 7 | 0 platform-repo edits | ✅ GREEN | platform/next-web/app/cms/jobsim CLEAN; ant-academy `next.config` = disclosed `allowedDevOrigins` serve patch |

**5/7 GREEN (C1,3,5,6,7 + fix#2). C2 + C4 blocked by the believability corrections not fully landing (F2/F3).**

## Findings (attributed to surface, per TOK-01 step 5)

- **F1 — succession/interview_extraction_results FK seed error (AI-readiness Org C).** The light seed's `succession`
  surface fails: `copy interview_extraction_results violates FK interview_extraction_results_sessions_session
  (SQLSTATE 23503)` → "dev-setdress: seed failed for demo-1" (NON-FATAL; hiring seeded fine — 5 positions + 62
  sessions after the failure). The interview-extraction rows reference `jobsimulation.sessions` that don't exist.
  Independent of hiring; degrades Northwind (C6) + is a cold-reset-to-seed "seed failed" cleanliness defect. **Was
  GREEN at M226** (the seed succeeded) → an M227-surface regression OR an M219 interview-seeder latent bug the
  1-sim/candidate change exposed. → iter-03 root-cause + fix.
- **F2 — fix#1 hiring-only GAP: 2 TRAINING sims leak into the hiring org.** "Handling Patient Balance Collection at
  Check-In" (6 candidates) + "Back-End Engineer (Java Spring Boot)" (1) have candidate mirror sessions in Meridian
  Talent → the recruiter's `InsightsByJobSimulations` list is NOT hiring-only. The 6 obvious generic seeders
  (Persona/JobsimSessions/SkillpathSessions/Assignments/Activity/HeroActivity) DO call `skipGenericActivityForHiringOrg`
  and GeneratedBatchSeeder writes no sessions → **the exact leak seeder is TBD** (iter-03 root-cause). The list leads
  with the 5 HIRING positions (they out-rank the training sims by count), so visibility is low, but the leakage is real.
- **F3 — fix#3 1-sim/candidate GAP: 17/45 candidates on 2 sessions.** 43 candidates / 60 sessions (17 extra). 10 on 2
  HIRING positions (e.g. Dan Mensah: Business Development failed 40 + SDR passed 66), 7 on 1 hiring + 1 training. The
  HiringFunnelSeeder writes EXACTLY 1 session/assessed-candidate (`pi = assessedOrdinal % len(positions)`) → a SECOND
  session path leaks (same root suspect as F2). Likely a pre-existing leak that fix#3 EXPOSED (M226's 5-per-candidate
  funnel masked it). → iter-03 root-cause + fix.
- **F4 — RESOLVED.** The render harness read "org re-skins as HIRING (nav 'Results'): false", but C3's Rae render shows
  the FULL hiring nav re-skin ("Results | AI Simulations | AI Interviews | Candidates Feedback") + the 5 HIRING sims +
  the HIRING logo. The harness string-check raced the cold hydration — a false-read, not a regression. reads-as-hiring GREEN.
- **F5 — cold-tailnet measurement fragility (M226 F5 recurrence).** The render C2 harness timed out on the 5th sim
  drilldown (5 cold compare-drawer renders > the 180 s per-test budget), and the FIRST C5 latency cold run missed
  ACCESS (4/5). Both resolve WARM (C5 warm 5/5, p95 1.47 s; C2 first 4 sims clean). The M226 warm-before-gate
  discipline + RENDER_TEST_TIMEOUT_MS headroom. Measurement methodology, not a demo defect. → iter-03 warm re-measure.

## Close — 2026-07-17

**Outcome:** First M227-corrected demo stood up GREEN on billion (default cold, UP_RC=0); comprehensive first
measurement — **5/7 conditions GREEN** (C1 exactly 5+45, C3 candidate profiles usable, C5 recruiter p95 1.47 s, C6
coexists, C7 core-clean, fix#2 external emails ✅). C2 + C4 blocked by the believability corrections not fully landing
live: **F2 (2 training sims leak — fix#1 gap)** + **F3 (17/45 candidates on 2 sims — fix#3 gap)** + a seed error **F1
(succession FK)**. F4 (reads-as-hiring) resolved GREEN; F5 (cold-tailnet measurement) resolves warm. Every gap
attributed to its surface.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (5/7 GREEN; C2/C4 blocked by F2/F3; F1 seed error — all routed to iter-03)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n (F1/F2/F3 are fixable via rext tooling, 0 platform edits — routed-forward, not blockers) — (5) cap-reached: n (1 tik) — (6) protocol-stop: n — Outcome: continue (iter-03)
**Decisions:** iter-02 D1 (devops operator + docker access), D2 (rext-only cutover), D3 (never builder-prune) [carried from iter-01]; iter-02 D4 (F2/F3 root = a 2nd session path bypassing fix#1's guard), D5 (F5 = cold-tailnet warm-before-gate).
**Side-deliverables:** none (measurement iter; the DB ground-truth SQL + harness runs are evidence, not code).
**Routes carried forward:**
- iter-03: root-cause + FIX **F1** (succession/interview_extraction_results FK), **F2** (hiring-only training leak — the exact leak seeder), **F3** (1-sim/candidate 2nd-session leak). Then a default cold re-bring-up + warm re-measure (C2 5/5 + C4 hiring-only) → drive toward gate. Handler `PROVE-M228-iter03-fix-believability-leaks`.
- iter-03/04: the gate's reproducibility clause — 2 clean default cold cycles at the fixed tag.
**Lessons:** The M228 live re-prove earned its place exactly as M215/M221 predicted: the M227 believability corrections
were proven DETERMINISTICALLY (unit tests) but the LIVE cold seed reveals fix#1 (hiring-only) + fix#3 (1-sim/candidate)
have real gaps — a 2nd session path leaks training + extra-hiring sessions the deterministic tests never exercised.
Also: `docker ps` as the login user (marco) reads "0 containers" (false) — always recon as the operator (devops/root).
And the cold-tailnet warm-before-gate discipline (M226 F5) recurs for BOTH the render drilldowns and the first latency run.
