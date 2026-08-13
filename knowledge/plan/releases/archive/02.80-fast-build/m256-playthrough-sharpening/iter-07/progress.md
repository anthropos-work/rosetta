# M256 · iter-07 — progress

**Type:** tik · **Active strategy:** `TOK-01` move 4 (the honesty items).
Iter shape per `corpus/ops/demo/playthroughs.md` § The iteration protocol (steps 3 → 4 → 5 → 6).

## Phase 0b / 0d

**0b SKIPPED** — plain tik, standing YELLOW inherited. **0d RUN and it FAILED on the tree iter-06 had just
committed** — see D27: iter-06's own fence file quoted the rejected tag verbatim in its header prose and
`ptvalidate --e2e-dir` harvested the comment as a **phantom Playthrough id**. The harness had been green three
consecutive times because `run-playthroughs.sh` reconciles with `ptreport`, which does **not** scan `@pt:` tags.
Fixed and fenced (the new fence test scans **every** file in `tests/`, unit specs included — where the phantom
lived). `ptvalidate` now VALID, `21 live Playthrough(s)`.

## The re-scope risk is RETIRED — the audit's F5 was wrong (D28)

Onboarding is **5 of the 9** UCs clause 3 must land, so F5 (*"no pre-onboarding state exists, and none can be
declared"*) was the finding most able to break this milestone. **It conflated org membership with onboarding
completion.**

- Onboarding completion lives in **`public.user_params.onboarding`** (`jsonb`; `SetOnboarding` in app's Ent,
  served by `onboarding(userId:)`). There is **no onboarding table**.
- It is **NULL for all 191 seeded users** — the pre-onboarding state is the **DEFAULT**, not something to seed.
- It **drives**: `/onboarding` renders the real first step with working `Upload` / `Skip` / `Next` for both
  `pt-employee` and `pt-manager`, no redirect, no `/login` bounce.

**Onboarding is UNBUILT, not impossible. The trigger is NOT tripped and clause 3's scope is NOT reduced.**
Build routed to iter-08 (`ONBOARD-M256-build`) rather than crammed into a mechanism iter.

## Phase A — the planned mechanism was REFUTED, and so was its proof target's diagnosis

**H1 (GraphQL outcome ablation) — REFUTED (D29).** The plan named the degenerate case in advance and
measurement landed squarely in it. Fulfilling every `POST **/graphql**` with `{data: null}` (15 intercepted):

| | baseline | ablated |
|---|---:|---:|
| outcome locator (`identityRegion`) | 1 | **0** |
| `body.innerText` length | 2147 | **24** |
| nav regions / buttons | — | **0 / 0** |

The outcome goes absent — **and so does the whole application.** A dead page, not an empty surface: the control
would pass for every Playthrough regardless of what it asserts, including one asserting pure chrome. It cannot
discriminate, so it is not a control. A gentler ablation needs per-operation response shapes — **O(queries), not
O(surfaces)** — which breaks the page-object layer's own scaling rule.

**H2 (the studio false green is its first proof target) — the TARGET's diagnosis was wrong (D31).** Driving the
real journey and polling all three matcher alternatives for 5 minutes:

```
route-header:Simulation Advanced Builder  ->  NEVER (5 min)
draft:Scenario Characters                 ->  +2.1 s
draft:Mission Tasks                       ->  +2.1 s
```

**The string iter-02 blamed never appears on the page at all.** The real mechanism is worse: the designer paints
its **empty section scaffolding** at +2.1 s, before the LLM draft populates it — the matcher fires on section
chrome that renders whether or not anything was generated. **The obvious fix would not have worked:** deleting
the never-matching header alternative changes nothing and would have shipped as a fix. The real fix needs a
**populated**-section landmark (a character card / a non-zero character count), which is unbuilt.

## Phase B — what landed, and what deliberately did not

Landed: **D27's phantom-id fix + the widened fence** (5 fence tests green; `ptvalidate --e2e-dir` VALID), and
the **measured evidence attached at the code** — `studio-builder-page.ts`'s locator now carries the timing table
and the explicit warning that the previously-routed fix is a no-op.

Deliberately NOT landed: a replacement negative-control mechanism. **Cross-vantage discrimination** is
identified with its rationale and its cost (D30 — run the Playthrough's own final locator against a hero for
whom the outcome legitimately does not exist; real absence, app stays alive, and it proves *which* data not
merely *that* data). It is **O(tests), not O(surfaces)**, so it is a build, not a coda to the iter that refuted
its predecessor. Routed as `NEGCTL-M256-cross-vantage`.

**The tripwire was applied, not pushed through (D32):** two falsifications and a third mechanism identified left
the iter with no path to its planned deliverable without starting a new mechanism from scratch.

## Phase C — the suite

All **120** unit/fence specs green (including the 5-test mutation-class fence). `ptvalidate --manifest-dir
manifest --e2e-dir e2e/tests --seed-worlds seed/seed-worlds.yaml`: **VALID — 9 products, 23 use cases, 21 live,
2 TODO.** No browser Playthrough changed behaviour this iter (the only runtime edit is a doc comment on a
locator), so iter-06's measured figures stand unchanged and **no re-measure is claimed**: clause 1 **0.6245×**,
`141 passed` ×3, 0 flake.

rext commit + tag **`fast-build-m256-negctl-falsified`**, pushed to origin.

## Close — 2026-07-28

**Outcome:** the routed negative-control mechanism was **refuted by measurement**, and the diagnosis of its
designated first proof target was **overturned** — the previously-planned studio fix would have shipped without
fixing anything. Against that, the milestone's largest risk was **retired**: onboarding is unbuilt, not
impossible, so clause 3 keeps its full scope. One iter-06 defect (a comment that minted a phantom Playthrough
id, invisible to the harness because the two gate tools see the tree differently) was found by this iter's own
re-survey, fixed, and fenced.
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET — clause 1 **met** (0.6245×, unchanged, no re-measure claimed); clause 2 mutating **5/5 MET**,
negative controls **5 of 21**, `blocked` **0**; clause 3 unchanged (2/4 org-admin, 0/5 onboarding — now known
BUILDABLE — verdicts unwritten); **D-v28-5** unstarted.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this is **1** no-lift tik; iter-06 progressed, so the streak is 1 of 3) — (3) re-scope: n — **and this is the iter that explicitly TESTED the trigger and found it NOT tripped** (D28: onboarding is buildable, so the 5 UCs are not `unimplementable`; 0 surfaces claimed unimplementable) — (4) user-blocker: n (suite green; the one red found was iter-06's own, fixed inside this iter; D-v28-3's consolidated red set is empty) — (5) cap-reached: n (2nd tik of run 2) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D27 (a comment minted a phantom Playthrough id; two gate tools, two views of the tree), D28 (the
re-scope risk retired — F5 conflated membership with onboarding), D29 (GraphQL ablation refuted — a dead page is
not an empty surface), D30 (cross-vantage discrimination identified as the replacement, with its O(tests) cost
stated), D31 (iter-02's studio diagnosis overturned — the blamed string never renders; the obvious fix is a
no-op), D32 (the tripwire fired and what it cost).
**Side-deliverables:** the phantom-id fix + the widened fence (D27) — a correction of an iter-06 defect, not
planned scope, so it does not upgrade the close status.
**Routes carried forward:**
- `NEGCTL-M256-cross-vantage` → **iter-08.** Replaces the refuted ablation route. Run each presence
  Playthrough's own final locator against a contrast vantage where the outcome legitimately does not exist.
  O(tests): budget it as a build across more than one tik.
- `FIX-M256-studio-false-green` → **re-aimed (D31).** Assert a **populated** section (character card /
  non-zero character count), NOT a section heading. The previously-routed fix is a **no-op** — do not ship it.
- `DOC-M256-llm-lane-premise` → still paired with the above, still **not** dischargeable until section CONTENT
  is measured (a heading's presence does not answer "did the generation complete on this host?").
- `ONBOARD-M256-build` → **iter-08.** Now known buildable (D28): `user_params.onboarding` is NULL for every
  seeded user and `/onboarding` drives. 5 of clause 3's 9 UCs.
- `BLOCKED-M256-refusal-surface`, `PT-M256-orgadmin-role-create`, `PT-M256-orgadmin-member-tag`,
  `FENCE-M256-bounded-interaction`, and all pre-existing routes in `../progress.md` still stand.
**Lessons:**
1. **Two gate tools that see the tree differently will hide a defect between them.** The harness was green three
   times over while the validator was red, because the loop runs `ptreport` and only `ptvalidate --e2e-dir`
   scans tags. When a pipeline has more than one gate, the iter loop must run **all** of them — or one of them
   must enforce the other's rule, which is what the new fence test now does.
2. **Name the degenerate case in the plan, then check for it FIRST.** The overview wrote down "if ablation
   blanks the app, the control proves nothing" before any code was written, so refuting H1 took one probe and
   two numbers instead of a day of building. Writing the failure mode down in advance is what made it cheap.
3. **Re-derive a routed diagnosis before implementing its fix.** `FIX-M256-studio-false-green` carried a precise,
   plausible, five-release-old-style diagnosis that was simply false. Implementing it would have produced a
   green commit, a closed handler, and an unchanged false green — the worst outcome available.
4. **A "no pre-X state exists" claim is usually a claim about the wrong column.** F5 reasoned from the seeder
   that writes memberships; onboarding lives in `user_params.onboarding`, is NULL by default, and was drivable
   all along. The milestone's biggest risk was a schema misreading.
