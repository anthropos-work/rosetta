# M256 · iter-07 — decisions

## D27 — a COMMENT minted a phantom Playthrough id, and the harness could not see it (iter-06 defect)

iter-07's mandatory Step-0 re-survey ran `ptvalidate --manifest-dir manifest --e2e-dir e2e/tests
--seed-worlds …` and it **FAILED** on the tree iter-06 had just committed:

```
[both-way] mutation: e2e test tagged @pt:<the-rejected-tag> has no manifest use case declaring it — an ORPHAN test
```

iter-06's own fence file explained the grammar collision **by quoting the rejected tag verbatim in its header
prose**, and `discover.go` harvested the comment as a Playthrough id. The fence that exists to prevent phantom
ids created one.

**Why it survived a whole iter, which is the part worth keeping:** `run-playthroughs.sh` reconciles with
`ptreport`, and **`ptreport` does not scan `@pt:` tags**. Only `ptvalidate --e2e-dir` does, and that runs
separately. So the harness was green **three consecutive times** while the validator was red. Two gate tools,
two different views of the tree, and the loop only ran one of them.

Fixed (the mention is now spelled apart) **and fenced**: a new fence test applies the validator's own rule
inside the harness, over **every** file in `tests/` including the `*.unit.spec.ts` meta-tests — because that is
exactly where the phantom lived, and a scan that skipped them would have missed it. `ptvalidate` now reports
`21 live Playthrough(s)`, VALID.

## D28 — the re-scope risk is RETIRED: the audit's F5 conflated MEMBERSHIP with ONBOARDING

The milestone's re-scope trigger is *"> 3 un-homed curated UCs prove unimplementable"*, and onboarding is **5 of
the 9** UCs clause 3 must land — so F5's conclusion (*"no pre-onboarding state exists, and none can be
declared"*) was the single finding most able to break the milestone. **It was wrong, and wrong in a specific
way: it reasoned from `UsersSeeder` writing a membership unconditionally. Membership is not onboarding.**

Measured:

| Question | Answer |
|---|---|
| Where does onboarding completion live? | `public.user_params.onboarding`, a `jsonb` column (`app/internal/data/ent/userparam_update.go` §`SetOnboarding`; served by `onboarding(userId:)`, `queries.graphqls:46`). There is **no onboarding table**. |
| What is its seeded value? | **NULL for all 191 seeded users.** The pre-onboarding state is the **DEFAULT**, not something to seed. |
| Does the flow drive? | **Yes.** `/onboarding` probed live for `pt-employee` and `pt-manager`: renders the real first step with working `Upload` / `Skip` / `Next` controls, no redirect, no `/login` bounce. |

**Verdict: onboarding is UNBUILT, not impossible.** The re-scope trigger is **not** tripped and clause 3's scope
is **not** reduced. The build is routed to iter-08 as `ONBOARD-M256-build` — a 5-UC coverage cluster is not a
side errand to be crammed into a mechanism iter.

This also retires the second half of `overview.md`'s Open Question 1, which the audit had answered "no".

## D29 — GraphQL outcome-ablation FAILS the honest-control requirement (H1 REFUTED)

The mechanism routed as gate-critical for clause 2's remaining 16 negative controls was: intercept the app's
own data query so the surface renders with **no data**, then assert the final locator does not match.

The plan named the degenerate case in advance — *"if ablation blanks the app entirely, 'the locator does not
match' degenerates to 'a broken page shows nothing' and proves nothing"* — and measurement put it squarely in
that case. Fulfilling every `POST **/graphql**` with `{data: null}` (15 requests intercepted) on the profile
surface:

| | baseline | ablated |
|---|---:|---:|
| outcome locator (`identityRegion`) | 1 | **0** |
| `body.innerText` length | 2147 | **24** |
| `role=navigation` regions | — | **0** |
| buttons | — | **0** |

The outcome does go absent — but so does **the entire application**. `bodyLen 2147 → 24` with zero nav and zero
buttons is a **dead page, not an empty surface**. A control built on it would pass for every Playthroughist
regardless of what the Playthrough asserts, including a Playthrough asserting pure chrome. It cannot
discriminate, so it is not a negative control; it is a tautology with a green tick.

**A gentler ablation is not obviously available either:** `{data: null}` breaks the client, and returning
*valid-but-empty* payloads requires a per-operation response shape — **O(queries), not O(surfaces)** — which
violates the page-object layer's own scaling rule and would rot on every schema change.

**Recorded so the next iter does not re-try it verbatim.**

## D30 — the replacement mechanism: CROSS-VANTAGE discrimination (identified, not yet built)

D29 leaves the 13 presence Playthroughs without a mechanism. The candidate with real precedent in this corpus
is **cross-vantage discrimination**: run the Playthrough's own final locator against a hero (or org) for whom
the outcome legitimately does **not** exist, and assert it does not match.

Why it is the right shape:
- The absence is **real product state** — no mock, no ablation, no second stack, and the app stays alive, so
  the honest-control requirement is satisfied by construction.
- It proves the assertion discriminates **WHICH** data, not merely **THAT** data — which is a strictly stronger
  claim, and it is the M219 lesson (*"a surface that renders is not the same as the RIGHT surface"*) applied
  per-Playthrough instead of in the single place it currently lives (`LEGACY_AI_READINESS_URL`).
- The existing `LEGACY_AI_READINESS_URL` counter-assert is the same idea in the route dimension, so the
  suite already contains the pattern in miniature.

Cost, stated honestly: it is **O(tests), not O(surfaces)** — each Playthrough needs a chosen contrast vantage.
That is real authoring work and it is why it was not begun in the same iter that refuted its predecessor.
Routed as `NEGCTL-M256-cross-vantage`.

## D31 — iter-02's studio false-green DIAGNOSIS was itself wrong, and the real mechanism is worse

`FIX-M256-studio-false-green` was routed with iter-02's diagnosis: *"`advancedDesignerRendered()` matches the
route's own `Simulation Advanced Builder` header."* Driving the real journey and polling all three alternatives
for 5 minutes:

```
route-header:Simulation Advanced Builder  ->  NEVER (5 min)
draft:Scenario Characters                 ->  +2.1 s
draft:Mission Tasks                       ->  +2.1 s
```

**The blamed string never appears on the page at all.** The real mechanism is that the advanced designer opens
and paints its **empty section scaffolding** — the "Scenario Characters" / "Mission Tasks" section headings —
2.1 s after the click, well before the LLM draft populates them. The matcher fires on **section chrome that
renders whether or not anything was generated.**

This matters practically: **the obvious fix would not have worked.** Deleting the (never-matching) header
alternative from the regex changes nothing, and would have shipped as a fix while the Playthrough stayed a false
green. The real fix needs a **populated**-section landmark — a character card or a non-zero character count
(`designer.actors.counter.label` = "characters") — which is unbuilt. Evidence attached to the locator in
`studio-builder-page.ts` so the next attempt starts from measurement.

`DOC-M256-llm-lane-premise` is **not** discharged and deliberately not half-written: presence of a section
heading does not answer *"did the generation complete on this host?"*, so the doc correction still lacks its
load-bearing fact. Measuring section **content** answers the fix and the doc premise together — they stay one
piece of work.

## D32 — the tripwire fired, and what it cost

Planned scope was one line: the ablation harness plus its designated first proof target. Phase A **refuted the
harness** (D29) and then **overturned the proof target's diagnosis** (D31), which is two falsifications and a
third mechanism identified but unbuilt (D30). At that point the iter had no path to its planned deliverable
without opening a new mechanism from scratch, so the scope-creep tripwire was applied rather than pushed
through: land what is complete (D27's fix + fence), route the rest with the measured evidence attached, and
grade the iter on its planned scope — which did not land.

**Status: `closed-no-lift`.** No clause-2 count moved. Three routed handlers now carry measurements instead of
hypotheses, one of them replacing a diagnosis that would have produced a fix that did not fix anything.
