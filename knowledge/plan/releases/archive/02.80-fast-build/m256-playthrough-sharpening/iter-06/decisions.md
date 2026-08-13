# M256 · iter-06 — decisions

## D21 — the two obvious read-back surfaces were REFUTED by measurement, and the refutations are the finding

The iter opened on two hypotheses that both read as near-certain from the source. Both were wrong, and being
wrong twice in one Phase A is now the third time in this milestone that **reasoning about a surface produced a
worse answer than probing it** (iter-02's studio false green, iter-05's `force: true`).

**H1 — "Start creates a session, so the path CTA should flip to Continue." REFUTED.** The write is real: a
`public.skill_path_sessions` row appears the instant `Start` is clicked (verified in the DB by `created_at`
inside the probe window). But it lands `status='pending', progress=0, started_at=NULL`, and next-web's CTA
label needs `progress` or `started_at`:

```
SkillPathHeader.tsx:216-219  /  SkillPathContent.tsx:618-629
  !progress && !startedAt     -> "Start"
  startedAt && progress == 0  -> "Continue"
  progress && progress < 100  -> "Continue (N%)"
```

Nothing on the create path sets `started_at`. So **a started-but-unadvanced path reads "Start" forever** — a
learner who starts a path and leaves is shown "Start" again on return, and the platform cannot distinguish
"never opened" from "opened, no step done". That is a **real product data-integrity wart**, reported here and
**not worked around** (zero platform edits). It is also why the read-back had to be `progress`, not the
enrolment.

**H2 — "`/sim/<slug>/session-list` will list the launched sim session." REFUTED, twice over.**
1. The launch writes **nothing**: `select count(*) from jobsimulation.sessions where created_at > <probe start>`
   returned **0** after a click that reached `/sim/<slug>/start` and rendered the launch confirmation.
2. Worse for the surface itself: it rendered "No sessions found" with 0 rows for a hero who has **6 seeded
   sessions in the DB**. `session-list/page.tsx:52-54` gates its query on `!user?.externalId`, and a
   Clerkenstein identity carries no `externalId`, so the query is **skipped** and the surface is permanently
   empty on a demo. A negative control built on it could never go green — it would be a **false RED**, which is
   exactly as dishonest as a false green.

## D22 — a mutating Playthrough's PRE-STATE read IS its negative control (H3, ACCEPTED — and it was already there)

Clause 2 asks that every Playthrough be *demonstrably RED when its outcome is absent*. The instinct is a second
stack, a mock, or a DOM ablation. For a Playthrough that WRITES, none of that is needed: read the target state
**before** the action and make the final assertion a **strict inequality or strict negation** against that
reading. Such a predicate is **false by construction at the pre-state**, so the run itself demonstrates the
assertion discriminates the outcome rather than matching chrome — against *real* product state, in the same run.

The retro-audit the plan called for then produced the better half of the finding: **all three pre-existing
mutating Playthroughs already had this property and nobody had named it.**

| Playthrough | its pre-state control | shape |
|---|---|---|
| `pt-orgadmin-tag-create` | `toHaveCount(0)` on the tag name before the write (`:45`) | explicit absence |
| `pt-orgadmin-setting-toggle` | `before = isOn()`, post-reload asserts `.toBe(!before)` | strict negation |
| `pt-assignment-assign` | `before = assignableCount()`, post asserts `.toBe(before - 1)` | strict delta |

So clause 2's negative-control sub-target was **partly already met and unrecorded** — 3 of the 5, before this
iter added 2 more. It was not visible because nothing named or counted it. Hence D23.

**Sub-decision: prefer the DELTA form over the absolute.** "The label is not yet `Continue (N%)`" would
false-RED on a re-run against a world that was not reset. A false red is as dishonest as a false green, so the
delta form is the one that ships.

## D23 — the mutation class is a MEASUREMENT, so it gets a fence (TOK-01 move 2's residual, discharged)

TOK-01 move 2 required "a machine-checked per-spec `MUTATES` / `READ-ONLY` / `UNKNOWN` tag (greppable, fenced by
a test) that the lane consumes instead of an assumed 17." iter-03 shipped the other half of that move; this half
did not land, so through iter-05 the count clause 2 GATES ON was a **prose claim** — three specs, agreed by
reading. A gate whose metric is a narrative is not a gate.

Shipped: `@pt-mutation: MUTATES|READ-ONLY|UNKNOWN` on every Playthrough (+ `@pt-negative-control:` required iff
`MUTATES`), fenced by `e2e/tests/mutation-class-fence.unit.spec.ts`. **Computed, not narrated:**

```
@pt-mutation registry (20 spec files, 21 Playthroughs): MUTATES=5  READ-ONLY=14  UNKNOWN=2
```

Three design points that are load-bearing rather than cosmetic:

1. **One class per `@pt:` id, not per file.** `studio-builder.spec.ts` holds TWO Playthroughs; a single tag on
   it would leave one unclassified while the fence read green — the same shape of blind spot the `networkidle`
   fence was widened to close at iter-03.
2. **The grammar must be DISJOINT from `@pt:`.** The first draft used `@pt:mutation`, and `ptvalidate` refused
   the tree with two ORPHAN errors (`discover.go:20` scans `@pt:([a-z0-9][a-z0-9._-]*)` and treats an unmatched
   hit as a dangling test). The fence pins the disjointness against its own copy of that regex, so a future
   widening of either grammar fails in the harness rather than in the Go validator.
3. **`UNKNOWN` is not "probably fine".** The two studio Playthroughs fire a real LLM generation that plausibly
   persists a draft, but neither is proven to read the write back — and `pt-studio-advanced-generate` is a
   **known false green** (iter-02). Classifying them `MUTATES` would have inflated clause 2's count with the one
   Playthrough already known not to prove its own outcome. `MUTATES` requires the read-back, full stop.

Mutation-verified: removing one tag turns 3 of the fence's 4 tests RED.

## D24 — the fifth write came from the writes the suite ALREADY made, not from a new surface

iter-05 D20 offered three candidates for the fifth mutating Playthrough (`Remove Tags` bulk action · profile
self-evaluation · an onboarding completion) — all *new* surfaces, and one of them (`profile.self-evaluation`) is
M206-reserved. The re-survey found a cheaper and more honest answer D20 had not considered: **two Playthroughs
already performed a real server-side write and simply never re-read it.** After H1/H2 removed the aisim half,
the shape still held for the skill path, and the second write came from the same surface's other control.

- **#4 `pt-skillpath-legacy`** — extended by one click. It previously stopped at *"the step-completion control
  is present"*, which the milestone's own problem statement named as cause #3 (*journeys stop at boundaries*),
  **while its manifest use case already promised "advance a step … progress tracks"**. So this was a
  spec-vs-manifest fidelity gap, not new scope: the spec now matches the promise, and reads the advance back on
  a fresh load (`Start` → `Continue (14%)`, measured).
- **#5 `pt-skillpath-bookmark`** — net-new, and chosen because it was **measured before it was declared**: the
  toggle held its new state across **three consecutive full re-navigations**, and the row is visible in
  `public.user_bookmarks`. It writes *and deletes*, reading both back, which makes it **self-cleaning** and the
  only mutating Playthrough in the suite that is re-runnable without a reset. The delete half doubles as the
  negative control in the second direction: the assertion that just passed is shown to fail once the outcome is
  removed.

**Honesty note:** `skill-paths.save-for-later` is **not** one of the M201 curated corpus's 28 use cases. It grows
the **manifest** denominator, not the curated one, so clause 3's coverage arithmetic (12 of 28 curated) is
unaffected. Recorded in the manifest header and the spec docblock so the two denominators cannot be conflated
later.

## D25 — a retry loop whose first attempt can eat the test budget is decoration (side-deliverable)

The Phase-C gate run went red on `pt-assignment-assign` — **245 s timeout in the suite, 6.0 s green on an
immediate solo re-run**. Under `retries: 0` ("a flaky Playthrough is a defect, not something to paper over") that
is a defect to fix, not a run to repeat.

Root cause: `assignments-page.ts` wrapped the antd-Select interaction in a **3-attempt retry loop whose first
`combobox.click()` had no `timeout`**. A Playwright action without one inherits the **test** budget (240 s here).
antd's Modal plays an open animation and its Form re-mounts the Select's inner `<input>` when the async option set
arrives, so the click saw *"element is not stable"* → *"element was detached from the DOM, retrying"* and
Playwright retried **silently for the full 240 s**. `attempt` never reached 1: **the retry loop was unreachable.**

This is M244's own lesson one level down — that fix gated on an OPTION being painted but left the click that
opens the list unbounded. Fixed in two parts, both necessary: wait for the **form** to have mounted (the submit
button attaching is the cheapest semantic signal the dialog body is rendered rather than mid-animation), and give
**every** interaction an explicit `timeout` so a stuck attempt yields to the next.

This is an unrelated-but-correct side-fix surfaced by the iter's own gate run — recorded as a side-deliverable,
committed separately, and it does **not** upgrade the iter's close status.

## D26 — clause 2 is now cleanly halved, and the remainder is named

| Clause 2 sub-target | Before iter-06 | After | Verdict |
|---|---|---|---|
| `>= 5` MUTATING Playthroughs | 3 (narrated) | **5 (machine-counted)** | **MET** |
| every Playthrough passes a negative control | 0 | **5 of 21** | NOT met |
| `>= 1` `blocked` outcome | 0 | 0 | NOT met |

The negative-control remainder is **16 Playthroughs that do not write** (14 `READ-ONLY` + 2 `UNKNOWN`). D22's
pre-state pattern does **not** reach them — there is no mutation whose absence to demonstrate. They need a
genuinely different mechanism, and iter-06 deliberately did not invent one under time pressure. The candidate
with the strongest evidence behind it is **outcome ablation**: block the surface's own data query with
`page.route` so the outcome is *genuinely* absent (not simulated), then assert the final locator does **not**
match. That is precisely what would have caught the studio false green, which asserts a header that renders with
no data at all — so the mechanism and its first proof target are the same piece of work
(`FIX-M256-studio-false-green`). Routed to iter-07 as `NEGCTL-M256-ablation-harness`.
