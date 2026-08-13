---
milestone: M256
iter: 09
iteration_type: tik
status: closed
active_strategy: TOK-01
created: 2026-07-28
---

# M256 · iter-09 — clause 3's verdict half, and the cheapest coverage it exposed

**Type:** `tik` · **Active strategy:** `TOK-01` move 4 (the honesty items). Handler:
`VERDICT-M256-remaining-uncovered`.

## Step 0 — re-survey

`ptvalidate --manifest-dir manifest --e2e-dir e2e/tests --seed-worlds …` VALID: 10 products, 29 use cases,
22 live, 7 TODO. Clause 1 at 0.5950×; clause 2 mutating 6/5 MET, negative controls 6 of 22, `blocked` 0;
clause 3 org-admin 2/4, onboarding 1 live + 5 verdicts, **verdicts still owed for 7 curated UCs**. Target current.

## Cluster / target identified

Clause 3's second half: *"every remaining uncovered curated UC carries a written verdict — zero silent gaps."*
Seven curated use cases still had none — 3 truly un-homed (`workforce.organization-feedback`,
`profile-skills.import`, `talk-to-data.query`) plus the **4 five-release-old M206/M207 reservations**, which the
gate requires be *verdicted*, not merely inherited. This needs no live stack and no new mechanism, so it is the
one remaining clause-3 requirement that can be **completed** rather than advanced.

## Hypothesis

**H1.** Writing a real verdict — one that names the **specific missing piece** rather than a category — will
show that **none** of the seven is `unimplementable-without-platform-edit`, which is the substantive claim the
re-scope trigger turns on.

**H2 (the one that paid).** Pricing each UC properly will expose at least one that is much cheaper than its
un-homed status implies, because the verdict work forces a look at whether its data is already seeded.

## Expected lift

- Clause 3's verdict requirement **COMPLETE**: 0 curated UCs without a written verdict.
- Any UC the pricing exposes as cheap, landed in the same iter.

## Phase plan

Read the curated definitions (M201 `manifest-draft.yaml`) → write the verdict register → land whatever the
pricing exposes as cheap → gate run ×3 → close.

## Escalation conditions

- If a verdict comes out as genuinely `unimplementable-without-platform-edit`, count it against the re-scope
  trigger and report the count rather than softening the verdict.

## Acceptable close-no-lift outcomes

- All seven price as blocked with nothing cheap enough to land; the register alone still completes the clause-3
  requirement.
