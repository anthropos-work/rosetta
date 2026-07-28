# M256 · iter-06 — progress

**Type:** tik · **Active strategy:** `TOK-01` — move 3 exhausted, move 4 opened; this iter is
`PT-M256-clause2-fifth-write` (iter-05 D20), answered differently than D20 framed it.
Iter shape per `corpus/ops/demo/playthroughs.md` § The iteration protocol (steps 3 → 4 → 5 → 6).

## Phase 0b / 0d

**0b SKIPPED** — plain tik; the milestone's standing pre-flight verdict (YELLOW, iter-01) is inherited and this
tik redirects into no unaudited subsystem. **0d RUN and PASSED** — the iter wires new artifacts through the
manifest → `ptvalidate` → spec → `ptreport` pipeline, so the gate tools were dry-run against the existing tree
first: `9 product(s), 22 use case(s), 20 live, 2 TODO`, valid; `ptreport` 20/22 passing, 0 failing.

## Phase A — measure the surface before writing a line of spec. Both hypotheses REFUTED.

Three probe rounds on the live `demo-2`, each followed by a DB read. This is the third iter in a row where
probing beat reasoning, and it is now written into the protocol doc.

**H1 — the skill-path CTA should flip `Start` → `Continue` after Start. REFUTED.** The write is real (a
`public.skill_path_sessions` row appears the instant `Start` is clicked) but lands `progress=0,
started_at=NULL`, and next-web's CTA needs one of those. **A started-but-unadvanced path reads "Start"
forever** — a real product wart, reported, not worked around.

**H2 — `/sim/<slug>/session-list` should list the launched session. REFUTED twice.** The launch created **0**
`jobsimulation.sessions` rows; and the surface showed "No sessions found" for a hero with **6 seeded sessions**,
because it gates its query on a Clerkenstein-absent `externalId`. A negative control there could never go
green — a **false RED**, as dishonest as a false green.

**What the probes then found (round 2/3, measured):** step completion → `Continue (14%)` after a full
re-navigation (`Start` count 2 → 0); and the bookmark toggle holding its new state across **three consecutive
re-navigations**, with the row visible in `public.user_bookmarks`. Round 2 also caught a **read race** — a bare
`count()` immediately after a re-navigation reported the stale state — which is why the shipped locators are
read through auto-retrying `expect`, never a raw read (documented at the page-object contract).

## Phase B — what landed

**Clause 2's mutating floor, 3 → 5, from writes the suite ALREADY made** (D24):

- **`pt-skillpath-legacy` extended by one click** → mutating #4. It stopped at *"the step-completion control is
  present"* — the milestone's own cause #3 — **while its manifest use case already promised "advance a step …
  progress tracks"**. A spec-vs-manifest fidelity gap, now closed: it completes the step and re-reads the
  persisted progress on a fresh load.
- **`pt-skillpath-bookmark`** (net-new, `skill-paths.save-for-later.UC1`) → mutating #5. Writes **and deletes**,
  reading both back — **self-cleaning**, and the only mutating Playthrough re-runnable without a reset. Declared
  in the manifest header as **NOT** an M201 curated use case, so the two denominators cannot be conflated.

**The negative-control pattern, ratified rather than invented** (D22 — H3 accepted): a mutating Playthrough's
**pre-state read IS its negative control**, provided the final assertion is a strict inequality/negation against
it — false by construction at the pre-state, demonstrated against real product state inside the same run. The
retro-audit found **all three pre-existing mutating Playthroughs already had it, unnamed** (`toHaveCount(0)` /
`.toBe(!before)` / `.toBe(before - 1)`). Delta form preferred over absolute: the absolute would false-RED on an
un-reset re-run.

**The mutation-class registry + fence** (D23 — TOK-01 move 2's undelivered half): `@pt-mutation:` on every
Playthrough + `@pt-negative-control:` required iff `MUTATES`, fenced by `mutation-class-fence.unit.spec.ts`.
The count clause 2 gates on is now **computed**:

```
@pt-mutation registry (20 spec files, 21 Playthroughs): MUTATES=5  READ-ONLY=14  UNKNOWN=2
```

One class **per `@pt:` id** (one file holds two Playthroughs); grammar **disjoint** from `@pt:` (the first draft
used `@pt:mutation` and `ptvalidate` refused the tree with two ORPHAN errors — the fence pins the disjointness
against its own copy of `discover.go:20`'s regex); `UNKNOWN` for the two studio journeys, deliberately, because
`MUTATES` requires the read-back and one of them is a known false green. **Mutation-verified:** removing one tag
turns 3 of the fence's 4 tests RED.

**Side-deliverable — a real flake fixed, not re-run** (D25): the Phase-C gate run went red on
`pt-assignment-assign` (**245 s in-suite / 6.0 s solo**). Its 3-attempt retry loop was **unreachable**: the first
`combobox.click()` had no `timeout`, so it inherited the 240 s test budget while antd's Modal animation +
async-option re-mount kept the inner `<input>` unstable/detached. Fixed by waiting for the form to mount and
bounding every interaction. Committed separately.

## Phase C — the gate run

**3 consecutive `run-playthroughs.sh 2 --reset` runs: `141 passed` each. 0 flake, 0 red.**
`ptreport`: **21/23 passing, 0 failing, 2 `[TODO]`, 0 unimplementable.** Both new/extended Playthroughs green on
all three: `skill-paths.legacy.UC1`, `skill-paths.save-for-later.UC1`.

### Clause 1 re-verified on the grown denominator (19 non-studio, was 18)

| Figure | Value |
|---|---:|
| **Median per non-studio Playthrough — the GATED metric** | **2.077 s** |
| **Ratio vs the iter-02 baseline (3.326 s)** | **0.6245×** — gate `<= 0.79×` **MET** |
| Honesty cross-check, the ORIGINAL 16 only | 2.006 s = **0.6033×** |
| Studio lane (excluded, budgeted separately) | 1.85 s / 1.28 s |
| Suite wall-clock (REPORTED, not gated) | 53.5 / 53.4 / 59.0 s → median **53.5 s** |

**Stated honestly:** this is *slower* than iter-04's 0.5434×, in two parts and both expected. (1) By design —
`pt-skillpath-legacy` grew 3.16 → 4.14 s (two extra navigations + a click) and `pt-skillpath-bookmark` at 3.51 s
sits above the median: **proving a write costs more than proving a render**, and that is the trade clause 2 asks
for. (2) The original-16 cross-check also drifted (1.778 → 2.006 s) with no code change to those specs beyond
docblock comments — run-to-run variance on a 9.70 GiB Docker VM against the documented 12 GB floor, plus the
`assignment-assign` form-mount wait. The gate is relative to the **iter-02 baseline** and holds with margin
either way; no number here is comparable to billion's (D-v28-12).

**Environment:** `Kirality-Mac-Pro-6.local`, darwin 25.1.0, Docker VM **9.70 GiB**; `demo-2` offset 20000,
**localhost/http**, `--no-public-host`, `workers: 1`, `retries: 0`.

## Phase D — knowledge

`corpus/ops/demo/playthroughs.md`: the live count reconciled **18 → 21 live / 2 TODO** (iter-04's addition was
never written down, so the corpus was two iters stale before this one made it worse); the `outcome` vocabulary
note updated with the M256 pre-flight's `actor.entitlement`-is-declared-only finding; and three
protocol-evolution entries added to § The iteration protocol — the pre-state-as-negative-control pattern, the
mutation-class registry, and the bound-every-interaction-in-a-retry-loop lesson.

rext commit + tag **`fast-build-m256-clause2-writes`**, pushed to origin.

## Close — 2026-07-28

**Outcome:** clause 2's mutating floor **3 → 5, MET and machine-counted**, discharged from writes the suite
already made rather than new surfaces — after both of the iter's own hypotheses were refuted by measurement. The
negative-control pattern was ratified (and found to be already present, unnamed, in all three pre-existing
mutating Playthroughs), TOK-01 move 2's undelivered fence landed, and a real 245 s suite flake was diagnosed to
an unreachable retry loop and fixed.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET — clause 1 **met** (0.6245× on the grown 19-test denominator, 0 flake over 3 runs); clause 2
mutating floor **MET (5)** but negative controls **5 of 21** and `blocked` outcomes still **0**; clause 3
unchanged (2/4 org-admin, 0/5 onboarding, verdicts unwritten); **D-v28-5** unstarted.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this tik progressed: mutating 3→5, negative controls 0→5, so no no-prog streak) — (3) re-scope: n (0 surfaces claimed unimplementable this iter; the onboarding assessment is routed, not tripped) — (4) user-blocker: n (the batch ended GREEN — the one red found mid-iter was diagnosed and fixed inside the iter, so D-v28-3 has an empty consolidated red set) — (5) cap-reached: n (1st tik of run 2) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D21 (both read-back hypotheses refuted; the `started_at`-NULL product wart; the session-list
`externalId` false-RED trap), D22 (a mutating Playthrough's pre-state read IS its negative control — and all
three pre-existing ones already had it), D23 (the mutation class is a measurement → the `@pt-mutation` registry +
fence, discharging TOK-01 move 2), D24 (the fifth write came from writes the suite already made), D25 (an
unbounded interaction makes a retry loop unreachable), D26 (clause 2 halved; the remainder named).
**Side-deliverables:** the `pt-assignment-assign` flake fix (D25) — a real defect surfaced by this iter's gate
run, committed separately; it does not upgrade the close status.
**Routes carried forward:**
- `NEGCTL-M256-ablation-harness` → **iter-07.** The 16 non-writing Playthroughs cannot use D22's pattern (no
  mutation whose absence to demonstrate). Build **outcome ablation**: block the surface's own data query with
  `page.route` so the outcome is *genuinely* absent, then assert the final locator does not match. Its first
  proof target IS `FIX-M256-studio-false-green` — the same piece of work, since that Playthrough asserts a header
  that renders with no data at all.
- `FIX-M256-studio-false-green` + `DOC-M256-llm-lane-premise` → **iter-07**, folded into the above.
- `BLOCKED-M256-refusal-surface` → a later tik. Clause 2's `blocked` outcome. `actor.entitlement` is
  declared-only (iter-01 D4), so it needs a real refusal: the `orgMemberCannotStartModal` the AI-sim Playthrough
  already asserts ABSENT is the strongest candidate — seed a member whose org lacks the
  `FEATURE_JOB_SIMULATIONS` g3 grant and the deny modal becomes the outcome. The locator already exists.
- `PT-M256-orgadmin-role-create` / `PT-M256-orgadmin-member-tag` → unchanged from iter-05; both still parked as
  diagnosed drafts.
- `FENCE-M256-bounded-interaction` → a later tik. D25's defect class is general: a source-scan fence asserting no
  unbounded `click`/`press` inside a retry loop in the harness. Not built here (it would have been a 4th line).
- `ONBOARD-M256-assessment` → **iter-07, explicitly and early.** Clause 3's onboarding half is 5 of the 9 UCs;
  the pre-flight established no pre-onboarding state exists and none can be declared. Assess whether it is
  *unbuilt* or *impossible* before spending a tik on specs — if impossible, that is >3 un-homed UCs proving
  unimplementable, i.e. the milestone's re-scope trigger.
- All pre-existing routes in `../progress.md` § Next-iter routing still stand.
**Lessons:**
1. **Ask what the platform WRITES, not what the UI implies.** Both refutations came from one `psql` query each.
   H1 looked certain from the label logic and H2 from the page's own purpose; a DB read cost seconds and
   overturned both. The corollary is sharper: *a write that lands is not the same as a write that is visible* —
   the skill-path session was real and unobservable, which is a different failure from not writing at all.
2. **A surface that is always empty is a false RED waiting to happen.** The session-list page would have made a
   perfect-looking negative control that could never go green. Before building a control on "the outcome is
   absent", prove the surface can ever show the outcome PRESENT.
3. **Count what the gate gates on.** Clause 2 gates on a number that was, for five iters, a sentence. Once it was
   computed, it was wrong in both directions — one Playthrough assumed to write that writes nothing, and three
   negative controls that already existed and were not credited.
4. **A retry loop with an unbounded first attempt is decoration.** It reads as defensive and cannot fire. The
   giveaway signature is a test that times out at exactly the test budget in a suite and passes in seconds alone.
