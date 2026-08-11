# iter-54 — intra-iter decisions

## D-M257x-54-1 — the platform clone is fast-forwarded, never edited

`stack-demo/platform` was brought `2adcf714 → ef32d4cd` by `git fetch` + fast-forward. **Zero platform-repo
edits**; nothing committed into it. The v2.8 constraint holds unchanged. Recorded because "we updated the
clone" and "we changed the platform" are one keystroke apart in a report and infinitely apart in fact.

## D-M257x-54-2 — the clause-3 fence was run BEFORE any corpus edit, deliberately

Running it after the map update would have proved nothing: a GREEN fence on a repaired tree is consistent
with a fence that does nothing. It was run first, went **RED with 3 direction-B findings naming
cms / jobsimulation / roadrunner**, and only then was the map updated (`EXIT=0`). This is the milestone's
own §8 discipline — *watch it go RED* — applied to a departure the fence was **not shown**, which is the
first non-staged catch it has made.

## D-M257x-54-3 — the over-claim is CORRECTED IN PLACE and the correction is left visible

Job 1 committed (at `6485151`) the claim *"the armed failure is now armed"*, citing
`demo-stack/migrate-demo.sh:81-85` / `:106` — line anchors into code that rext `54bccf7` (**this milestone's
own iter-02**) deleted. iter-01's finding was quoted forward without re-measuring against iter-02's repair.

Verified false by running the derivation live, not by reading it:

```
$ source stack-core/lib/repos_yml.sh; repos_yml_migration_pairs   <repos.yml @ ef32d4c>  → app:public
$ …                                   repos_yml_schemas_to_create <repos.yml @ ef32d4c>  → extensions sentinel public
$ echo "$REXT_TRANSITIONAL_SCHEMAS"                                                       → (empty)
```

**Decision: correct the claim, and keep a visible note of what it said.** Not erased. It is the cheapest
demonstration the milestone has of its own founding class — a stale claim with stale anchors, in the map
built to stop stale claims, that no fence covers — and deleting the evidence to make the document look
clean is the behaviour this milestone exists to end. Two files corrected: the map's §5 row 3 + header, and
`iter-54/platform-before-after.md` §4.

## D-M257x-54-4 — §5's "rows to watch" signal for `storage`/`messenger` was DEAD, and is replaced

It read *"when `repos.yml` flips either to `migrations: false` with a `legacy` comment, the fold has
landed."* Both have read `migrations: false` since long before the fold was announced (`repos.yml:18-23` @
`ef32d4c`) — so the stated signal was already true and could never fire. That is the map committing **Trap A
from its own §1**. Replaced with the signal that actually fires and actually did: **departure** from
`repos.yml` + compose-service deletion, i.e. direction B of the §4 fence.

## D-M257x-54-5 — the gate is reported at **2 of 5**, not the booked 4 of 5

Clauses 1 and 2 were met against `2adcf71`. The gate's own wording is *"Against platform @ **origin HEAD**
(never a pinned pre-drift commit)"*, and origin HEAD is `ef32d4c`. **Stale, not failed** — and reporting
them as still-met would be precisely the "pinned pre-drift commit" the clause forbids. Restoring both costs
~40 minutes of machine time; the expected result is green and is pre-registered in `reassessment.md` so it
can be refuted.

## D-M257x-54-6 — clause 2's gate-meeting run recorded no platform ref, and this is booked as evidence

`iter-37/progress.md` contains **no platform sha**; the only sha-shaped token in the file is `ad524614`. Its
ref is inferred from iter-36 (which re-fetched `2adcf71` at open and close and said so). Booked as
**occurrence 4** of TOK-04's class — an input free to move without appearing anywhere — and it is the
occurrence that sits **inside the gate itself**. It is the direct evidence for policy rule **P1**.

## D-M257x-54-7 — clause 5's residual is reported as a FLOW, and the number is negative

81 fresh drift sites in one working day against a repair rate of ~18/iter at ~50% induction gives a net of
**−72** for the cycle. Recorded as a decision rather than an observation because it changes what the
milestone is optimising, and because **the previous clause-5 metric could not go negative** — so it could
never have reported that we were losing. Clause 5 itself is **not** re-cut (user ruling, third time).
