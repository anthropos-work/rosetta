# M224 — Progress

_Iterative milestone: a **running ledger**, not a section checklist. `/developer-kit:build-mstone-iters` appends one
entry per iter and creates `iter-NN/` dirs as it goes. iter-01 is the BOOTSTRAP tok (authors the first strategy)._

## Running ledger

| iter | kind | what changed | gate metric | outcome |
|------|------|--------------|-------------|---------|
| iter-01 | tok (bootstrap) | KB-fidelity GREEN (hiring.md FAPI-pointer fix inline); authored TOK-01 (recruiter-render-first) | baseline UNMEASURED (presumed 0 rows) | closed-fixed — see iter-01/progress.md |
| iter-02 | tik | Clerkenstein org `publicMetadata.isHiring` wired end-to-end (seeder roster → FAPI); `/align-run` GREEN 100/100 ×2; rext tag `casting-call-m224-iter02` | UNCHANGED (fix-half/scaffold — no render yet) | closed-fixed — see iter-02/progress.md |
| iter-03 | tik | recruiter cockpit seat (Rae Ramirez, manager→admin, slot-1, funnel-skipped) + `curatedTalent` skill family + manifest regenerated; rext tag `casting-call-m224-iter03` | UNCHANGED (scaffold — no render yet) | closed-fixed — see iter-03/progress.md |
| iter-04 | tik | **FIRST GATE READING.** Baseline `min(rows/sim)=0` → attributed **SEED-GAP** (recruiter role `"Technical Recruiter"` unresolvable → hiring-seed cascade fail; M219 trap caught by measure-first). **Fixed** (role → `"Talent Acquisition Specialist"`). Cold reset → **Meridian 50 members; 5 sims × 43 candidates; scores 27–100** — DATA side MET. rext tag `casting-call-m224-iter04` | 0 (data now present; render still 0 — see iter-05) | closed-fixed — see iter-04/ |
| iter-05 | tik | **DIAGNOSIS.** Data present but recruiter still can't reach the scoreboard: `apps/web` **ejects** an all-hiring-orgs user to the Hiring sub-app (`UserStatusContext.tsx:141-173`), not in the demo. Attribution = **product-boundary eject**, NOT a render-gate → falsifies M222's apps/web premise. Cheap fix: render-probe timeout 150s→300s (kept fullPage fix). rext tag `casting-call-m224-iter05`; rosetta `ae4974e` | render 0 (blocked on the eject; data side MET) | closed-fixed (attribution) — see iter-05/ |
| iter-06 | tok (triggered) | **STRATEGY REVISION → TOK-02 "run-the-real-hiring-app".** iter-04/05 evidence + an adversarially-verified two-app feasibility workflow: the genuine comparison ships in `apps/hiring` and reads the SAME seeded `local_jobsimulation_sessions` via the SAME Cosmo backend. Pivot from "patch the eject + re-skin the workforce app" → **build the real Hiring app into the demo as a 2nd UI container.** More faithful; reuses existing data/backend; zero platform edits. | metric unchanged (render UNMEASURED in apps/hiring) | strategy-set — see **TOK-02** in decisions.md |
| iter-07 | tik A (TOK-02) | **THE TWO-APP DEMO IS LIVE.** Built the REAL `apps/hiring` into demo-1 as a 2nd UI container: rext-owned `hiring.Dockerfile` (filter `@anthropos/hiring-app`, port 3001 → offset `:13001`), `up-injected.sh build_frontend_hiring`, a self-sufficient `gen_injected_override.py hiring_lines` block, exposure-guard update; tests green (injection 149, frontend-build 86, exposure 15/15). `demo-1-hiring-app-1` **Up** (17 containers, the 16 healthy ones untouched); `:13001 GET /` → 307 to the **FAKE FAPI** `:15400/sign-in` (Clerkenstein, NOT real Clerk — the app talks to the demo's mock login, unlike apps/web) + reaches Cosmo `:15050`. rext tag `casting-call-m224-iter06`; ZERO platform edits (clone git-clean post-build). | render not yet measured (iter-08) | closed-fixed — rext `c24bc2b` |
| iter-08 | tik B (TOK-02) | **FIRST HIRING-APP RENDER READING: `min(rows/sim)=0` — but STRUCTURAL WINS.** Recruiter logs into the hiring app (`:13001`), **NOT ejected**, org **re-skins as HIRING** ("HIRING Meridian Talent"); the cockpit now routes her `[Log in as]` → the hiring Results page (Dan stays `:13000`). **Attribution = a platform-source role-remap ASYMMETRY** (not seed, not eject): `apps/hiring UserStatusContext:174` stores `role: userRole` raw (`org:admin`), lacking `apps/web`'s `remapUserRole` (`:77/198`, `org:admin→admin`) → she's treated non-admin → `EnterpriseWrapper` bounces her to candidate Home → 0 insights queries fire. DATA side still MET (43×5). rext tag `casting-call-m224-iter07` (373505c): cockpit `--hiring-base` + per-hero base + render-probe `--hiring` flag. ZERO platform edits. | render 0 (blocked on the role-remap wall; data MET) | closed-fixed (reading+attribution) — rext `373505c` |

## Next iter

**iter-09 (tik C, under TOK-02) — the role-remap demo-patch → rebuild → re-measure (drive toward green).** The one
remaining wall is a platform-source asymmetry: `apps/hiring` doesn't remap the Clerk org-role the way `apps/web`
does, so an admin recruiter reads as non-admin and is bounced off the enterprise routes. **Fix = ONE sha-pinned
demo-patch** on the demo's ephemeral `apps/hiring/src/context/UserStatusContext.tsx`: add `apps/web`'s
`remapUserRole` + change `:174` `role: userRole` → `role: remapUserRole(userRole)` (a single funnel that fixes
`EnterpriseWrapper` + template `isAdmin` + home). This makes `apps/hiring` **consistent with `apps/web`** (the
correct behavior); it does NOT touch Clerkenstein (which correctly emits `org:admin`, faithful to real Clerk — a
Clerkenstein "fix" to emit `admin` would REDUCE fidelity). Wire into `build_frontend_hiring`'s transient
apply/revert (like the pk overlay), sha-pin per `demopatch-spec.md` G1–G7 (standalone — NOT the `urls.ts` chain).
**Rebuild the hiring image (~7 min)**, then `run-hiring-render.sh 1 rae-recruiter --hiring` → expect enterprise nav
+ the insights query firing + ~43 rows/sim. Then drive the ≥40 × 5 gate over ≥3 cold reset-to-seed runs.
**Deferred (non-blocking):** 6 pre-existing `test_cockpit.py` failures (4 removed-academy-CTA, 2 v2.3.1 overlay-JS —
HEAD-identical, not this milestone's) → a future harden pass.

> **Note (bookkeeping):** iter-04/iter-05 were executed + committed (rext tags `…iter04`/`…iter05`; rosetta
> `ae4974e`) during a driven build-iter leg; their ledger rows are reconciled here at the TOK-02 boundary. iter-04's
> `overview.md` predates its close; its outcome is captured in this ledger + TOK-02's trigger section.
