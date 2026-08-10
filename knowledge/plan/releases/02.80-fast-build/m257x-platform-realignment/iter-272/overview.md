---
iter: 272
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
active_strategy: TOK-08
route: FIX-M257x-267-capture-the-succession-RESPONSE
---

# iter-272 — capture the succession RESPONSE: read the wire, not the tables

**Type:** tik, under `TOK-08` (*census the mechanical classes; stop sampling them*).

## Step 0 — re-survey (mandatory, before targeting)

Gate **clause 2** is the last functional clause: the Playthrough suite must pass **30 live / 0 failing / 0
error**. One Playthrough fails — `pt-workforce-succession` — with correct chrome and empty projection
tables, while three siblings pass on the same login and the same seed.

iter-267 refuted **both** candidate causes by direct measurement and named the next one:

> The next measurement is not more SQL — it is the `GetSuccession` HTTP response for Org A, captured
> from a logged-in manager session, which distinguishes *the backend returned rows and the page dropped
> them* from *the backend returned an empty projection from populated inputs*.

Two preconditions that blocked it are now gone, both cleared by this run:

| precondition | state |
|---|---|
| the frozen rext pin (`D-M257x-258-1`) | **spent** at iter-270; pin is `fast-build-m257x-iter-270` |
| a green stack to measure against | **up** — iter-271, 4 green cycles, the 4th at 20:31:00Z |

**The target is still meaningful, and one thing about it must be re-measured before anything else:**
iter-267's failure was observed on a `demo-2` built by the *old* tooling. This stack has been rebuilt four
times since, at a pin 206 commits newer. **Whether the Playthrough still fails is not an assumption this
iter is allowed to make** — it is the iter's first measurement, and a pass would be a finding in its own
right.

## Cluster / target identified

`GetSuccession` (`app/internal/workforce/succession.go:215`) fans out six queries; iter-267 ran all six
verbatim against this org and **every input is populated** (28 members / 280 role requirements / 266
declared skills / 33 verified / 89 session activity / 12 interview signals), no query errors, no
decommissioned schema on the stack. The fault is therefore **above the data layer** — scoring/threshold
arithmetic, the response caps, the API/route, or the frontend — and SQL cannot separate those.

The wire can. **One response body partitions the residual**: rows on the wire ⇒ the fault is at or below
the renderer; no rows on the wire from populated inputs ⇒ the fault is in the projection arithmetic or its
caps.

## Hypothesis

The measurement is decisive regardless of which way it lands. **No cause is predicted** — iter-267's
lesson 1 was that a disjunction offered without its residual is a framing, and this iter is not going to
repeat it one layer up.

## Expected lift

- A captured `GetSuccession` response body for the manager's own org, from a real logged-in session.
- The residual **partitioned** — the fault placed above or below the renderer, with the response as
  evidence.
- No gate-metric movement is promised. Clause 2 needs a fix; this iter buys the fix a target.

## Phase plan (declared multi-step — the tripwire counts UNPLANNED lines only)

1. Seal these pre-registrations (first commit).
2. Re-measure: run the single Playthrough against the freshly-rebuilt stack, reset-to-seed.
3. If it still fails, extract the succession response from the retained Playwright **trace** — the harness
   already sets `trace: 'retain-on-failure'`, so the capture needs **zero** code and zero edits to the
   pinned clone.
4. Read the body; partition the residual; record.

## Out of this iter's planned scope (declared, so the tripwire is clean)

- **Fixing** whatever the response reveals. If the fix is in rext it needs a tag; if it is in the platform
  it is forbidden outright (v2.8: 0 platform edits). Either way the fix is the *next* iter's work, and
  saying so now is what keeps this iter one line of investigation.
- Gate clause 5 (the reading), and every route inherited from iter-270.

## Escalation conditions

- **No edits to the pinned `stack-demo/rosetta-extensions` clone.** Tooling is authored in the
  `.agentspace` copy and tagged. This iter is designed to need neither.
- `demo-1` is not ours. `--reset` targets `demo-2` only, and `stackseed`'s own N=0 guard is the floor.
- If the response cannot be captured from the trace, say so and name what would capture it — do **not**
  reimplement the login flow to get one, which would be a second, unmeasured instrument.

## Acceptable close-no-lift outcomes

- The Playthrough **passes** on the rebuilt stack — clause 2 moves and the failure was environmental. A
  real result, and it must be verified as a *repeatable* pass, not a single draw.
- The response is captured and the residual is partitioned but no fix lands. That is the deliverable.

## Pre-registrations (sealed in this iter's FIRST commit, before any measurement)

- **PR-1 — it still fails.** `pt-workforce-succession` fails again on the rebuilt stack after
  `--reset`. *Refuted by:* a pass.
- **PR-2 — the failure is unchanged in shape.** It fails on a projection/row assertion (the succession
  signal, the at-risk section, the role→candidate rows, the key-role card, or the hero row) and **not** on
  login, URL or the page heading. *Refuted by:* a bounce to `/login`, a URL mismatch, or a missing heading.
- **PR-3 — the trace carries the answer.** The retained failure trace contains the succession GraphQL
  response body, so no new instrument is needed. *Refuted by:* no trace, or a trace with no response body
  for that operation.
- **PR-4 — the wire is EMPTY.** The captured response returns a succession projection with **zero** rows
  (empty arrays / nulls) despite the six populated inputs — i.e. the fault is in the projection arithmetic
  or its caps, not in the renderer. *Refuted by:* rows present on the wire.

PR-4 is the one real prediction here, and it is deliberately the half that would make the fix **ours to
find in arithmetic** rather than a rendering bug. It is stated so it can be wrong.
