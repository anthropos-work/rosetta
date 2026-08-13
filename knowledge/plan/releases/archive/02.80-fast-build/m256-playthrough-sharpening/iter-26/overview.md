---
iter: 26
iteration_type: tik
iter_shape: standard
status: closed-fixed
opened: 2026-07-30
---

# iter-26 — the day-0 readiness seat: a seeder capability, then the UC it unblocks

**Active strategy reference:** `TOK-01` move 3's successor clause — onboarding, *"its cost is a seeder +
capability + roster seat, not just specs."* iter-24 priced the seeder half; this iter builds it and spends it.

## Step 0 — re-survey (mandatory)

| checked | reading |
|---|---|
| gate clause 3 — onboarding | **1 of 5** — the milestone's only remaining gap |
| gate clause 3 — org-admin | **4 of 4 ✅** (iter-22) |
| gate clause 2 | controls **24 of 26**, mutating **8/5**, `blocked` **1/1** — MET with the D103 carve-out |
| gate clause 1 | flake half MET (187 × 3, 0 flake); leg half N/A |
| `D-v28-5` | discharged + proven live (iter-25) |
| `ONBOARD-M256-stage0-capability` | still owed; routing named **iter-25**, which spent itself on D-v28-5 |
| `aiReadinessStageFor` | re-read: hero branch is manager → 0, struggling → 1, **default → 3** — unchanged since iter-24 measured it |
| `seat_append_test.go` test 3 | still GREEN, i.e. the gap it holds is still open |

Target unchanged from the routing table. No substitution.

## Cluster / target identified

`ONBOARD-M256-stage0-capability` → and then the use case it is the prerequisite for,
`onboarding.enterprise-workforce-ai-readiness.UC1`.

The capability alone moves no gate clause, and a UC alone cannot be built — iter-24 proved the seat it needs
**cannot be declared**. So the honest unit of work is both, in one iter, in that order.

## Hypothesis

1. **A stage-0 end-user seat can be expressed without a Trajectory change.** Readiness stage is orthogonal to
   the life-arc, so an explicit per-hero field is both truer and ~5 seeders cheaper than a third trajectory.
2. **The stage-0 surface differs from stage 1 in a way a Playthrough can assert.** If it does not — if the
   day-0 and started members render the same thing — then the UC adds nothing over
   `pt-aireadiness-member-progress` and the honest outcome is a written verdict, not a Playthrough.
3. **The guided flow's first step is drivable and persists**, giving a mutating proof with a real read-back.

## Expected lift

- clause 3: onboarding **1 of 5 → 2 of 5** (the gate's last open clause);
- clause 2: controls **24/26 → 25/27**, mutating **8 → 9** (the UC lands WITH its control, so clause 2's
  ratio does not regress as clause 3 advances — the iter-17/22 rule);
- the `ONBOARD-M256-stage0-capability` handler closed, and `seat_append_test.go` test 3 discharged per its
  own failure message.

## Phase plan

- **A — probe LIVE first** (iter-21's rule). Manufacture a stage-0 vantage on `demo-2` by removing the
  started hero's one progress row, read the surface, restore. Answer hypotheses 2 and 3 before writing code.
- **B — the capability**, with mutants, in `stack-seeding`.
- **C — the seat**: append to `pt-world.seed.yaml` (append-only, per iter-24's fence) + the roster +
  the precondition; **reseed live and verify she reads stage 0**.
- **D — the Playthrough + its in-line negative control**, page-object layer measured on BOTH seats first.
- **E — the gate**: 3 × cold reset-to-seed, the Go module sweep, `gofmt`, docs.

## Escalation conditions

- If Phase A shows stage 0 and stage 1 render indistinguishably → **do not build the Playthrough.** Write the
  verdict instead and say why (a Playthrough that cannot fail is what this milestone removes).
- If the guided step cannot be completed on a demo → the UC is presence-only; re-price and route.

## Acceptable close-no-lift outcomes

A measured "the surfaces do not discriminate, so this UC is not landable and here is the verdict" is a
complete outcome. Building a Playthrough whose green does not depend on anyone onboarding is not.
