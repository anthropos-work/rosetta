# M201 Progress

This is an **`iterative`**, **USER-GUIDED** milestone — progress accrued per top-down pass toward the **exit
gate** (not a fixed section checklist). The gate + the iteration protocol live in [`overview.md`](overview.md).
Worked **conversationally** (the user directed each pass). The gate:

> **The manifest corpus is comprehensively outlined, validated, and written as prose-intent YAML — covering the
> platform's products × their must-work user journeys, each use case carrying goal + actor + flow +
> intermediate/final expectations, structurally valid (ids unique + both-way) — and the USER signs off the corpus
> as the complete-enough v2.0 coverage contract.**

**Status:** `done` — **closed-on-gate 2026-06-29** (gate MET: corpus authored + structurally valid + user
sign-off). See the Gate Outcome Ledger below.

## Running ledger

The corpus was curated top-down, one product per pass, conversationally — each pass: outline the product's
stories → use cases → validate against the **real platform surface** (next-web clones + corpus + the
path-migration spec) → write prose-intent YAML → the user OK'd/steered the next slice. Deliverable:
[`manifest-draft.yaml`](manifest-draft.yaml) (single draft file — the spec §5.3 "one file per product" split is an
M202 task per `overview.md`).

| pass | product / slice | validated against | precondition → M202 | user steer / sign-off |
|------|-----------------|-------------------|---------------------|-----------------------|
| 1 | **product spine** (10 → 9 active) + Onboarding (4 flavors) | live next-web onboarding branch (OnboardingIndividual/User, managerImport, hiring) | day-0 actors ×4; hiring org + assigned sim | spine locked; hiring = "land + assigned hiring sim" |
| 2 | Skill Paths (legacy / academy) | path-migration spec R2 + skillpath engine | legacy path w/ assessment sim; academy login wiring | OK |
| 3 | AI Simulations (chat / code / interview) | jobsimulation engine; Roadrunner→Judge0 | code-exec path; per-flavor sim in catalog | OK (non-voice; voice deferred) |
| 4 | Profile & Skills (import / self-eval / verified / growth) | /profile, Skill Spotlight, my-growth surfaces | resume fixture; ≥2 verified datapoints | OK |
| 5 | Workforce Intelligence (funnel / roster / talent-pool / ai-readiness-mon) | /enterprise/workforce(+ (new)/*) routes | AI-readiness org; varied AI-skill coverage | OK |
| 6 | Org Admin & Settings (roles / members / tags / feature-config) | /enterprise/{roles,members,tags,settings} | manager-admin actor; two-backend reset | OK |
| 7 | Assignment & Monitoring (assign + track) | /enterprise/{assignments,activity-dashboard} | assigned content w/ mixed progress | "in" |
| 8 | Studio (guided / advanced) | studio-desk mode chooser + studio-room | studio-desk cross-port login | OK (both modes) |
| 9 | **AI Labs → DEFERRED** | backend lab.v1.LabSessionService (nil client) | — | "defer to next release" (no runnable surface) |
| 10 | Talk to Data | /enterprise/talk-to-data + internal/askengine | PostHog FF + admin gate; Bedrock creds | OK |
| 11 | **adversarial fidelity verify** (11 agents, wf `wvpnpvozh`) | the live stack-demo clones | all negative verdicts (re-ground post-sync) | **SIGN OFF** 2026-06-29 |

## Gate Outcome Ledger (iterative close)

- **Gate status: `closed-on-gate` — MET.**
  - *Structural half:* 9 products · 26 stories · 28 use-cases; every UC carries goal + actor + flow +
    intermediate/final expectations; ids unique + both-way (manual check — the §5.3 **validator** is an **M202**
    deliverable, so structural validity is asserted by inspection here, to be machine-confirmed when M202's
    validator runs over this corpus).
  - *User-sign-off half:* the user **signed off** the corpus as the complete-enough v2.0 coverage contract
    (2026-06-29).
- **Deliverable:** [`manifest-draft.yaml`](manifest-draft.yaml) — the prose-intent Product → Story → Use Case
  corpus (single draft file; one-file-per-product split + landing in the rext `playthroughs` section is M202).
- **Beyond-gate enhancement (not required by the gate, done for quality):** an 11-agent **adversarial fidelity
  verification** (workflow `wvpnpvozh`) re-grounded every UC against the live clones → **15/27 runnable · 11
  partial · 1 not-runnable**. It resolved all 4 open flags with code evidence, caught a **2nd AI-Labs-class trap**
  (academy not runnable in the demo), and **discovered the stale-clone problem** (next-web @ v2.33.2, **115+
  commits behind prod**) — which turned the member-AI-readiness "not-runnable" into a proven **false negative**.
- **Carry-forward / re-grounding (handed to the v1.10 backfill, NOT this milestone):**
  - Every **negative** verify verdict must be **re-grounded against fresh clones** before it's trusted (the
    staleness caveat in the manifest header).
  - 2 PENDING items left **as-authored** per the user: `skill-paths.academy.UC1` ("leave here" — re-decide in the
    backfill) and `onboarding.enterprise-workforce-ai-readiness.UC1` (stale false negative — KEEP, re-verify).
  - The **M202 seed/wiring corrections** + the **coverage holes** (WI Growth/Trends/Activity-Log tabs; document
    sim; assignment edit/unassign; roster⊂members dedup) are catalogued at the bottom of the manifest for M202.

## Pivot note

After sign-off, the user **paused v2.0 "opening night"** to interpose a **dedicated v1.10 backfill** (re-sync the
corpus + stack clones to current prod, then re-validate) — the staleness this milestone surfaced is the backfill's
thesis. This M201 corpus is **preserved as the v2.0 spec**, resumable after the backfill. The user drives the
backfill kickoff; this close merges the milestone work to `main` so the backfill starts from it.
