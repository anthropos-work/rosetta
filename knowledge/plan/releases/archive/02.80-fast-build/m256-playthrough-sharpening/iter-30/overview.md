---
iter: 30
iteration_type: tik
iter_shape: standard
status: closed-fixed
opened: 2026-07-30
---

# iter-30 — pricing the last onboarding UC, by measurement rather than by estimate

**Active strategy reference:** `TOK-01` move 3's onboarding clause, final target.

## Step 0 — re-survey (mandatory)

| checked | reading |
|---|---|
| onboarding | **4 of 5** — `individual.UC1` is the only gap in the whole exit gate |
| clauses 1 + 2 | MET (197×3, 0 flake; controls 27/29, MUTATES 11, blocked 1/1) |
| `individual.UC1`'s recorded blocker | *"needs a seeder change (a member-less user + a roster seat)"* — an ESTIMATE, never measured |

## Cluster / target identified

`onboarding.individual.UC1`. It is the last thing standing between this milestone and its exit gate, and its
blocker has been carried as prose since iter-08 without anyone testing either half of it.

## Hypothesis

Two things are unknown and only one of them is about cost:

1. **Can an org-less user reach the app at all?** If the platform cannot serve a member with no membership, the
   UC is `unimplementable-without-platform-edit` and the milestone's re-scope trigger comes into view. This is
   the load-bearing unknown and **nothing has ever tested it.**
2. **What does a member-less user actually cost to seed?** The recorded blocker says "a seeder change",
   singular.

## Expected lift

**No gate clause moves.** The deliverable is the UC priced from measurement instead of estimate — landable or
not, and if landable, exactly which seeders must change.

## Phase plan

- **A — measure (1)** the cheap way: delete a seeded hero's membership on `demo-2` (restored by `--reset`,
  never a seeder written first) and drive her.
- **B — measure (2)** from the schema: the FK fan-out from `memberships`, and the real dependent-row count for
  one hero.
- **C — write the verdict** and route the build with both numbers attached.

## Escalation conditions

- If (1) fails — the app cannot serve an org-less user — say so plainly; that is a re-scope conversation, not a
  seeding one.
- **Do NOT start the seeder change in this iter.** Its size is precisely what is being measured, and iter-28's
  stop showed what happens when a build is begun on an unmeasured premise late in a session.

## Acceptable close-no-lift outcomes

Either answer to (1) is a complete outcome. What would not be acceptable is carrying "needs a seeder change"
forward for a fifth iter without testing it.
