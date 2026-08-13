**Type:** tik · shape: standard (single target: `org-admin.members.UC1`)

# iter-17 — the wall was real; the route through it was a key press

## Phase A — probe the three routes

iter-05 left this use case with two named untried candidates, both pointer routes. A third, unlisted one is
stronger on first principles — **a pointer-interception overlay cannot block a KEY event** — so all three were
measured on the live modal before anything was written.

| route | result |
|---|---|
| click the `<label for=…>` (iter-05's candidate) | **REFUTED** — 14 labels, **`labelsWithFor: 0`**, inputs with no `id`; antd wraps the input *inside* the label, and the label click is intercepted exactly like the box |
| plain `click()` on the checkbox | **times out** — iter-05 finding 1 re-confirmed |
| `focus()` + `Space` | **WORKS** — `checked: 0 → 1`, submit `"Assign Tags[DISABLED]" → "Assign Tags (1)[enabled]"` |

The dropdown stays open throughout (`openMenuitems: 12`, even on a passing run): the route stops needing it to
close rather than closing it. iter-05's other candidate — a dropdown-free route — was never needed.

**The observation that proves it worked is the submit button's own tally**, not `isChecked()`. `Assign Tags (1)`
+ `enabled` come from React state; `isChecked()` is only the DOM's opinion.

## Phase B — implement

`MembersBulkActionsPage` gained `tagCheckbox(name)` (scoped by the antd wrapper's visible text, since there is
no `for` and no `id` to bind to), `assignTagsSubmit()`, and `assignTag(name)` — focus, Space, wait for the
submit to report a tally, then commit **by keyboard too** (the button is under the same overlay). The parked
draft moved from `e2e/drafts/` into `tests/` with its `@pt-mutation: MUTATES` +
`@pt-mutation-evidence` + `@pt-negative-control` tags, and the manifest's `playthrough: TODO` became
`pt-orgadmin-member-tag`. The 13-iter-old "MODAL IS CURRENTLY UNREACHABLE" note was rewritten rather than
deleted: **findings 1–3 still hold and are what make the keyboard route legible.**

## Phase C — mutants, and one that PASSED when it should have gone RED

| mutant | expected | result |
|---|---|---|
| M1 — remove the assignment step entirely | RED | **RED** (the tally stays 0; the final is falsifiable) |
| M2 — swap the keyboard tick for `check({force:true})` | RED (per iter-05 D19) | **PASSED** |

M2 was the interesting one. Rather than adjust the mutant, the original iter-05 experiment was **repeated in
isolation** — and it does not reproduce (full retraction in [`decisions.md`](decisions.md) **D84**):

```
before: {"checked":0,"submit":"Assign Tags|DISABLED"}
after : {"checked":1,"submit":"Assign Tags (1)|enabled"}     ← check({force:true}) ALONE
```

So `force: true` **does** drive antd's state on this surface today. Findings 1–3 still hold (a plain `click()`
still times out, re-measured), so only the `force`-specific conclusion is withdrawn.

**Two corrections were made because of it.** The comment this iter had just written claimed `force: true`
"could never produce the tally" — that was **overstated and is fixed in place**, in the page object and in the
manifest. And keyboard still ships, on the argument that survives: **`force: true` exists to SKIP actionability
checks, so it is the one interaction that CAN manufacture a state the app never learns about** (iter-07's rule).
A route needing no `force` cannot have that failure mode. That is a decision about what a green run is evidence
*of*, not a claim about what `force` can do.

*A mutant that passes when you expected RED is data.* The first instinct was to fix the mutant.

## Phase D — re-measure

**Gate: 3 consecutive cold reset-to-seed runs, `173 passed`, rc 0, 0 flake.**

| run | result | `ptreport` | wall |
|---|---|---|---|
| 1 | **173 passed**, rc 0 | **25** passing / 0 failing / **6** TODO / 0 unimplementable | 1.9 m |
| 2 | **173 passed**, rc 0 | same | 1.8 m |
| 3 | **173 passed**, rc 0 | same | 1.3 m |

- `@pt-negative-control` registry, computed: **22 of 25** (9 self-declared + 13 via the control spec) — the new
  Playthrough arrives already covered (D85), so numerator and denominator both moved by one.
- `ptvalidate`: **VALID** — 10 products, 31 use cases, **25** live Playthroughs, **6** TODO (was 24 / 7).
- `playthroughs` + `clerkenstein` Go modules rc 0, 0 FAIL. `gofmt -l` clean. `tsc --noEmit` clean.
- Drifted `demo-2` cockpit fixture backed up before the first `--reset`, restored after the last: **`99e2f315`**.

## Close — 2026-07-29

**Outcome:** **`pt-orgadmin-member-tag` is LIVE — org-admin 2 of 4 → 3 of 4**, through a route iter-05's four
pointer-based attempts had not tried: the keyboard. And **iter-05 D19's `force: true` finding is retracted**,
found by a mutant that passed when it should have failed and then confirmed by repeating the original
experiment.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — clause 2 mutating **7** (≥5 MET), `blocked` **1/1 MET**, negative controls **22 of 25**;
clause 3 verdict half **COMPLETE**, landed half **org-admin 3 of 4** (was 2), onboarding **1 of 5**; clause 1
leg half **N/A** (no speed mechanism), flake half **MET**; **D-v28-5 root-caused + half-fixed** (iter-16).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (landed its planned target) — (3) re-scope: n (0 `unimplementable`; the escape valve was NOT needed) — (4) user-blocker: **y — an INSTRUCTED PAUSE, not a technical block** (the user is commuting; the coordinator directed a stop at this iter boundary with budget remaining) — (5) cap-reached: n (4th tik of this invocation, cap is 5) — (6) protocol-stop: n — Outcome: exit-4
**Decisions:** D83 (**a pointer-interception overlay cannot block a KEY event** — the route through a
13-iter-old wall, with the submit button's own tally as the proof the *application* registered the tick),
D84 (**RETRACTION of iter-05 D19** — `check({force:true})` does drive antd's state on this surface today;
findings 1–3 stand; keyboard still ships because `force` is the one interaction that can manufacture unlearned
state), D85 (the negative control came with the journey: the Playthrough creates the tag it assigns, so the
tally starts at a known 0 and the final asserts 0 → ≥ 1 across a surface boundary).
**Side-deliverables:** none.
**Routes carried forward:**
- `PT-M256-orgadmin-role-create` → **the LAST org-admin TODO (3 of 4 landed).** Unchanged shape from iter-05
  D18: "Suggest skills" transforms the create-role dialog and the primary button becomes **`Generate`** (a
  live-LLM leg → budget it with the studio lane, not the median). Separately: **report the product defect** —
  `Save` is enabled with the form incomplete and fails silently with an EMPTY `alert` region. **New for the
  next attempt:** try the KEYBOARD route (D83) before concluding anything is unreachable, and do NOT rely on
  iter-05 D19's `force` claim (retracted, D84).
- `ONBOARD-M256-import-path` → onboarding still **1 of 5**; the four remaining each have their specific missing
  piece written into `manifest/onboarding.yaml` (résumé fixture + async LLM import · the org-prepared trigger
  condition · an org-less actor · an Org C stage-0 seat · a day-0 hiring-org seat). Longest pole in the gate.
- `NEGCTL-M256-cross-vantage` → **22 of 25.** `pt-hiring-recruiter-compare` still needs a same-vantage control
  whose *absence* half is unmeasured (priced at iter-15); 2 studio blocked behind `FIX-M256-studio-false-green`.
- `D-v28-5-cockpit-logout` → half done (iter-16); D81's handshake rule + one joint live proof on a `fake-fapi`
  rebuild.
- Everything else from iter-16's list stands unchanged.
**Lessons:**
1. **When a wall has been measured four times, check whether every attempt shared an assumption.** All four of
   iter-05's were pointer interactions, and so were both of the candidates it left behind. The route through
   was a different *input modality* — cheap to try, and untried for thirteen iters.
2. **A mutant that passes when you expected RED is data, not a broken mutant.** The first instinct was to fix
   the mutant. Repeating the original experiment instead retracted a milestone-level finding that had been
   relayed onward as fact.
3. **Correct your own prose in the same breath.** This iter wrote "`force: true` could never produce the tally"
   and measured the opposite twenty minutes later. Leaving it would have shipped a fresh overstatement inside
   the very comment that documents a retraction.
