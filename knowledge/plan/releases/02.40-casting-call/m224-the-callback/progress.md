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

## Next iter

**iter-07 (tik A, under TOK-02) — the hiring-app UI container.** Build `apps/hiring` into the demo as a second UI
container from the **untouched clone** (same recipe as `apps/web` + `studio-desk`): a **rext-owned**
`hiring.Dockerfile` (the platform `Dockerfile.dev` hardcodes `--filter=@anthropos/web-app`/`start:web`/`EXPOSE 3000`,
so it can't be reused verbatim — filter the hiring package, its `start` cmd + `EXPOSE`) + a `FRONTENDS` row in
`demo-stack/…/gen_injected_override.py` (offset port, fake-Clerk pk + FAPI URL bake, `CORS_EXTRA_ORIGINS`,
`UI_BROWSER_FACING`). Bring up demo-1; **confirm the hiring app boots and reaches the same Cosmo backend + fake
FAPI**. Then iter-08 = cockpit re-point + probe re-point + first Hiring-app render reading; iter-09+ = drive to
green (≥40 × 5 over ≥3 cold runs). **Zero platform-repo edits** (build context only; at most the chained
`urls.ts HIRING_APP_URL` demo-patch, or bypass via the cockpit deep-link). Full plan + verified premise + risks:
**TOK-02** in `decisions.md`.

> **Note (bookkeeping):** iter-04/iter-05 were executed + committed (rext tags `…iter04`/`…iter05`; rosetta
> `ae4974e`) during a driven build-iter leg; their ledger rows are reconciled here at the TOK-02 boundary. iter-04's
> `overview.md` predates its close; its outcome is captured in this ledger + TOK-02's trigger section.
