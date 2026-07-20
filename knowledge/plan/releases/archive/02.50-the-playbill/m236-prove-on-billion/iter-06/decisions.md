# iter-06 — Decisions

## D1 — the interview manager view is a distinct surface (a false fail, not a defect)

`/activity-dashboard/interviews/<simId>/<membershipId>` renders a breadcrumb
(`AI Interviews / AI Readiness Interview / Nadia Ferrari`) over an attempts table with a *View Report*
action. It carries **no** `<player>'s Results for <sim>` header — the exact string the `manager-dashboard`
shape asserts on — so a fully working page was graded broken.

Fixed by adding `manager-interview` (shape 5), selected **by route**.

**The milestone-level pattern this completes.** Of the four distinct "manager/skill-path failures" M236
investigated:

| Symptom | Verdict |
|---|---|
| skill-path player graded as a scored sim | **wrong assertion** (and a false PASS) |
| manager scoreboard "No data" + "undefined undefined" | **real bug**, one clean cause (D1 iter-05) |
| interview manager "no header" | **wrong assertion** |
| interview *player* page suspiciously short (205 chars) | **correct page**, terse by design |

Three of four were the harness, not the product — and the single real defect had exactly one cause and one
small fix. **When a gate is new, disbelieve the gate first**: probing the page costs one run; mis-triaging
into the product costs an iter and can produce a "fix" for a non-problem.

## D2 — the skill-path manager hang is characterized, not fixed

`/enterprise/activity-dashboard/skill-paths/df9d2142…/50e1bb5e…` exceeds a **180 s** navigation timeout.
Three facts constrain it:

1. **Independent of the membership-id fix** — it failed identically before and after, and now carries the
   corrected membership id.
2. **Its sibling passes** — `sp-genai-in-progress` renders fine on the same route family, so the surface
   and the seat both work.
3. **The two differ by weight** — the hanging one is the **completed 13-chapter** path; the passing one is
   a **3-chapter** path at 45 %.

One heavy instance hanging while a light sibling passes is the **per-item fan-out** signature
`latency-budget.md` documents (a cost that scales with item count, not a broken route). That doc's guidance
— *name the arithmetic signature before reading code* — points at a per-chapter query rather than a
failure.

Not fixed here: this iter's declared scope was residual triage, and the 5-tik cap fires on its close.
Routed to **iter-07** (`SKILLPATH-M236-iter07-manager-hang`) with the contrast above as the starting
evidence, so the next iter opens on a hypothesis rather than a blank page.

**Not claimed:** that this is a demo-only issue or a platform defect. It has not been measured against a
non-demo stack, and no such claim belongs in the record until it is.
