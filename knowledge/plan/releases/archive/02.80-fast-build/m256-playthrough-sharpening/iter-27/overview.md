---
iter: 27
iteration_type: tik
iter_shape: standard
status: closed-fixed
opened: 2026-07-30
---

# iter-27 — the hiring-org day-0 member, and the routing trap read at line level

**Active strategy reference:** `TOK-01` move 3's onboarding clause. iter-26 landed UC 2 of 5; this targets the
next one by *landability*, not by list order.

## Step 0 — re-survey (mandatory)

| checked | reading |
|---|---|
| onboarding | **2 of 5** (`completion.UC1`, `enterprise-workforce-ai-readiness.UC1`) |
| gate clauses 1 + 2 | unchanged and MET (195×3 0 flake; controls 25/27, MUTATES 9, blocked 1/1) |
| `standard.UC1` (self-import) | blocked behind a **measured product defect** (CV upload 200, forward control never enables, 100 s+, both formats) — iter-18's refusal STANDS |
| `standard.UC2` (org-prepared) | trigger still **unidentified** |
| `individual.UC1` (org-less) | needs a **member-less user** capability — every per-index seeder writes org-scoped rows for a hero's slot, so the blast radius is large |
| `enterprise-hiring.UC1` | needs a non-recruiter day-0 hiring seat + an assigned sim + a **discriminator against the cockpit's own routing** |

**Target chosen: `enterprise-hiring.UC1`** — the only one of the four whose blockers are all inside our own
seed and harness, with no product defect in front of it and no unknown trigger.

## Two source-level readings that re-priced it BEFORE any code (Phase 0 research)

1. **The routing trap is REAL and now confirmed at line level.**
   `apps/web/src/context/UserStatusContext.tsx:142-173` ejects to the hiring app on
   `userHasAllHiringOrgs` **alone** — there is no onboarding condition in it. So the UC's intermediate
   *"the member is routed into the hiring app"* would, if asserted in apps/web, prove **`UserStatusContext`**,
   not onboarding — exactly what the manifest warned about, and it is not a suspicion any more.
   **But the hiring app has its OWN onboarding route** (`apps/hiring/src/app/(authenticated)/(signedup)/onboarding`),
   whose `onClose` is `router.replace('/home')` — **inside the hiring app**. Driving that surface directly makes
   the destination onboarding-owned and never touches the eject.
2. **"Assigned and ready to start" needs NO new capability.** `heroHiringStage`
   (`seeders/hiring_funnel.go:184`) already pins a **STRUGGLING** candidate hero to `assignedOnly` —
   *"assigned-not-taken (a pending assignment, not yet on the scoreboard)"* — and an end-user hero in a hiring
   org becomes a **candidate** automatically (`endUserHeroRole`, M224). So the seat is one YAML append.

## Hypothesis

A day-0 candidate hero appended to Org D, declared `struggling` (→ assigned-only), can be driven through the
**hiring app's own** onboarding and land on the hiring app's `/home` with her pending assignment visible and
startable — a cross-app onboarding proof that does not lean on the `UserStatusContext` eject.

## Expected lift

clause 3: onboarding **2 of 5 → 3 of 5**; clause 2: controls 25/27 → 26/28, MUTATES 9 → 10.

## Phase plan

- **A — append the seat, reseed, probe LIVE** (hiring `/onboarding` served? completion destination? does
  `/home` show a startable pending assignment?).
- **B — the spec + its in-line control**, page objects measured first.
- **C — the gate** (3 × cold), module sweep, docs.

## Escalation conditions

- If the hiring app's onboarding is not reachable for her, or `/home` shows no startable assignment → the UC's
  final cannot be honestly asserted; write the **verdict** with the line-level evidence and close.

## Acceptable close-no-lift outcomes

The `UserStatusContext` finding alone upgrades this UC's verdict from *"a trap to settle first"* to a
line-level fact. If the hiring-app onboarding path does not carry the journey, that is a complete outcome.
