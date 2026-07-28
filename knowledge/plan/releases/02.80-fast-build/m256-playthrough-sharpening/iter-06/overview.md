---
milestone: M256
iter: 06
iteration_type: tik
status: in-progress
active_strategy: TOK-01
created: 2026-07-28
---

# M256 · iter-06 — clause 2's mutating floor, from the writes the suite *already makes*

**Type:** `tik` · **Active strategy:** `TOK-01` move 3/4 boundary — the org-admin cluster is exhausted at
2 of 4 (iter-05 refuted both remaining hypotheses), so the mutating floor must be discharged from
elsewhere. This is `PT-M256-clause2-fifth-write` (iter-05 D20), answered differently than D20 framed it.

## Step 0 — re-survey (mandatory)

- `ptvalidate --manifest-dir manifest` → **VALID**: 9 products, 22 use cases, 20 live, 2 TODO.
- `ptreport` → **20/22 passing, 0 failing, 2 unimplemented, 0 unimplementable.**
- Clause 1 unchanged (**0.5434×**, iter-04) — no harness change is planned that could move it, but it is
  re-verified in Phase C because the gate is measured on the post-coverage suite.
- **TOK-01's named next-target is stale in one respect and is substituted:** D20 offered three candidates
  for the fifth write (`Remove Tags` bulk action · profile self-evaluation · an onboarding completion).
  All three are *new* surfaces. The re-survey found a cheaper and more honest answer that D20 did not
  consider: **two Playthroughs already MUTATE and simply never READ BACK.** Same TOK strategy (move 3's
  discharge-per-unit-of-work rule); different target.

## Cluster / target identified

Clause 2 defines MUTATING as **mutates state AND reads it back**. The suite has **3** such Playthroughs.
But it has **two more that already perform a real server-side write** and stop at the launch boundary
without ever re-reading it:

| Playthrough | The write it already makes | Where the write is READABLE (not yet asserted) |
|---|---|---|
| `pt-skillpath-legacy` | `Start` creates a `SkillPathSession` (server-side create-on-read; the spec **self-declares** MUTATION at `:21-23`) | The path detail's CTA label is derived from the session: `SkillPathHeader.tsx:216-219` returns `t('start')` when `!progress && !startedAt`, `t('continue')` once `startedAt` exists. `SkillPathContent.tsx:618-629` computes the same flip independently. |
| `pt-aisim-chat-launch` | `Start Simulation` creates a `jobsimulation.sessions` row (`aisim-chat-launch.spec.ts:61`) | `/sim/<slug>/session-list` — an own-sessions table with an explicit empty state (*"No sessions found."*) and one row per session (id · Started At · Status *In Progress*). |

So the fifth write does not need a new surface, a new seeder, or a new modal. It needs the **second half of
a journey the suite already drives** — which is exactly what this milestone is named for.

## Hypothesis

**H1.** Re-navigating to the skill-path detail after `Start` renders `Continue` in place of `Start`, because
the label is computed from a **persisted** `SkillPathSession.startedAt` read back over GraphQL — so a full
re-navigation (not a client-side state read) proves the write landed.

**H2.** `/sim/<slug>/session-list` shows **zero rows and the "No sessions found" empty state** for
`pt-employee` on the sample chat sim before launch, and **exactly one more row** after — proving the
`jobsimulation.sessions` write landed.

**H3 (the one that matters most).** For a mutating Playthrough, **the pre-state assertion IS the negative
control** the gate asks for. Clause 2 wants each Playthrough "demonstrably RED when its outcome is absent."
A before/after flip demonstrates exactly that *within the same run, against real product state*: the
outcome-bearing locator is asserted **absent** before the action and **present** after. That is a stronger
demonstration than an ablation/mock, because the absence is genuine rather than simulated — and it costs
nothing extra, because the run is already there.

## Expected lift

- Clause 2 mutating count **3 → 5** — the floor **MET**.
- Clause 2 negative controls **0 → 2** demonstrated-RED-when-absent, under a pattern (H3) that generalises
  to every mutating Playthrough and is fenced so it cannot be dropped.
- A retro-audit of the 3 existing mutating Playthroughs against H3, so the count of Playthroughs carrying a
  negative control is stated honestly rather than assumed.

## Phase plan

Per `corpus/ops/demo/playthroughs.md` § The iteration protocol, steps 3 → 4 → 5 → 6:

- **Phase A** — probe both read-back surfaces LIVE on `demo-2` before writing a spec line. iter-02's
  studio false green and iter-05's `force: true` finding both came from asserting against a surface whose
  behaviour had been *reasoned about* rather than *measured*. Probe first.
- **Phase B** — extend the two page objects with the read-back locators; add the pre-state + post-reload
  assertions to the two specs; update the two manifest use cases' expectations. Fence H3.
- **Phase C** — `run-playthroughs.sh 2 --reset` + `ptreport`; re-verify clause 1 on the current suite.
- **Phase D** — close; commit; tag rext and **push the tag** (rung zero).

## Escalation conditions

- If either read-back surface does not behave as measured in Phase A → do **not** force it. Route the
  Playthrough forward with the measurement, and substitute the other of D20's candidates for the count.
- If the pre-state is **non-empty** (the seed already wrote sessions for this hero+sim), the negative
  control for that Playthrough is unavailable in this shape → assert the **delta** for the mutating half
  and record the negative control as routed forward, honestly, rather than claiming it.
- A platform defect discovered here escalates or gets a diagnosed draft — **never** a workaround
  (zero platform edits).

## Acceptable close-no-lift outcomes

- Both read-back surfaces measured and found **not** to reflect the write on re-navigation (i.e. the write
  is not persisted, or the surface reads client cache) — that would be a real finding about the platform
  and would refute H1/H2 with evidence.
- H3 refuted: the pre-state assertion turns out not to be a valid negative control for a reason the plan
  did not see. Recording *why* is worth more than a second unproven mutating count.
