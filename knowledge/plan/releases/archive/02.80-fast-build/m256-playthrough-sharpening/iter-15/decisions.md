# iter-15 — decisions

## D74 — a fence whose own example of a SAFE site produces the hang it said would move its boundary must actually move

`bounded-interaction-fence.unit.spec.ts` scoped itself to retry loops and enumerated what it left out:
*"straight-line interactions elsewhere in `lib/` and `tests/` — 28 sites at this pass … a blanket rule would
be 28 edits of which zero have evidence behind them … **If a straight-line site ever produces an opaque hang,
it becomes evidence and this boundary moves — with the measurement recorded, as D25's was.**"*

Its self-test (b), the one proving the fence was not trigger-happy, quoted two lines. One of them was
`ai-readiness-page.ts`'s `getByText(/How we measure/i).first().click()`.

**Measured this iter:** on an org that has not enabled AI-readiness that tab **does not exist** (0 elements —
`/ai-readiness` renders a marketing upsell instead), and a probe calling `openHowWeMeasure()` there hung for
**600 s** before reporting `locator.click: Test timeout exceeded`. The boundary's own trigger, on the boundary's
own example.

The old rationale — *"the test budget IS their intended ceiling and the failure names a real line"* — is true
and useless. The line was named, after ten minutes, for a fact decidable in one second.

**The widening, and why it is this one.** The harmful property is not straight-line-ness; it is that the
element **may legitimately not exist on some vantage**. That is not statically decidable in general. But a
method whose *name* is `open…` / `switchTo…` / `expand…` / `reveal…` / `drill…` exists to reveal a surface, so
"the control that reveals it is absent" is a legitimate outcome of calling it, and a bound is the only way to
report that outcome in bounded time. The author's chosen name is the declaration that makes the rule decidable.

**Measured before adopting, not argued:** the rule flags **7 sites across 5 files**. Seven evidence-backed
edits versus twenty-eight evidence-free ones is the whole difference between a boundary moving and a fence
rotting into noise and then being switched off.

Implementation notes that are part of the decision: the scanner is a sibling function (not a widened regex) so
the loop rule keeps its own precise report; it is **exported** so the self-tests drive the real function rather
than a re-implementation (iter-12's fence learned that one the hard way); the denominator of disclosure methods
is **asserted**, because a rule decided by a naming convention goes silently empty when the convention drifts,
and an empty scan is indistinguishable from a clean one; and self-test (b) is **kept**, correctly re-scoped, as
the honest record of where the boundary was and why it moved.

## D75 — a probe must use the predicate the CODE uses

Phase A pass 1 counted elements whose `textContent.trim()` **equalled** a step name and got **0** on the
not-enabled org — which read as a refutation of iter-12's finding that the upsell panel satisfies those
assertions. It was not a refutation. `stepMethod()` matches `new RegExp(name, 'i')`, a **substring**, and the
upsell renders `"STEP 1 / AI Skill Mapping with AI Framework"`. Re-measured with the accessor's own predicate:
1 match per step, exactly as iter-12 said.

An exact-text probe over a substring accessor is not a stricter measurement, it is a **different question** —
and answering the wrong question confidently is how a correct prior finding gets overturned by mistake. Sibling
of iter-14 D72 (*the DOM shape is a measurement, not an inference from a sibling page object*): here the thing
inferred rather than measured was the *matching semantics*.

## D76 — an assertion that cannot tell "no data" from "a service is down" is the could-not-fail class in a hat

`pt-activity-drilldown` asserted `contentRows().first()` visible, then `count() > 0`. Measured: when
`jobsimulation` is down the grid renders **20 `<tr>` with `textContent === ""`** — indefinitely, watched for
40 s, `mainLen` 469, **no empty state and no error anywhere in the UI**. Both assertions pass on that.

The Playthrough still failed, three steps later, on `drillIntoActiveContent()`'s row-link wait — a timeout
that blames a locator. The **sharpened** assertion failed immediately and named the actual condition: *the
rows carry no data*.

So the fix is not only "assert content because a skeleton exists during hydration" (the framing this iter
opened with); it is that a row-count assertion **cannot distinguish an empty result from a dead dependency**,
and a suite whose job is to detect breakage must not have assertions that report the wrong cause. Same class as
a check that cannot fail: it fires for the wrong reason instead of not firing at all.

## D77 — a clean `Exited (0)` is not a healthy container

`demo-2-postgresql-1` restarted un-cleanly at 14:38 (*"database system was not properly shut down; automatic
recovery in progress"*). `demo-2-jobsimulation-1` logged *"DB too many ping failures, **shutting down**"* and
exited — **status 0**, a graceful self-shutdown exactly as designed — taking `demo-2-cms-1` with it. Nothing
restarted either. Disk was fine (227 GiB free), so this is **not** the documented ENOSPC trap (`build-budget.md`
M239-F1), which is the first thing that came to mind and the wrong thing.

Recovery was `docker start` on the two containers: not a bring-up (no build, no compose, no teardown) and the
same class of action `run-playthroughs.sh --reset` already performs on the fake services every run. Both
Playthroughs green in 7.3 s afterwards.

The operational shape worth recording: **`docker ps` showed 14 of 16 containers "Up" and the application
surfaced no error at all** — the jobsimulation surfaces just rendered blank rows. Routed as
`FIX-M256-demo2-service-self-termination` → a container-liveness assert in `stack-verify`'s cheap-win class
would have named this in one line instead of costing an hour of Playthrough diagnosis.

## D78 — pick a control's liveness witness by measurement, not by which hero looks obvious

The drill-down control needs the contrast tenant's own equivalent asserted PRESENT through the same accessor
(iter-13 D64). The obvious choice was Org C's *thriving* hero — the analogue of the proof hero. **Measured: she
is not in that content's breakdown at all**; the STARTED hero is. Using her would have produced a control whose
liveness half was simply absent, i.e. a false RED, and the temptation to then "fix" it by relaxing the liveness
assert is exactly how a control quietly becomes vacuous.

## Escalation not taken, and why

The mid-iter RED met the letter of a Phase-5 § 4 user-blocker (the suite went red for a reason outside the
iter's planned scope). It was not escalated because five cheap measurements resolved it to a self-terminated
container and a non-destructive `docker start` recovered it inside the iter — no decision was required of
anyone. Recorded so the judgement is auditable rather than implicit: the test is whether the user must
*decide* something, not whether something went wrong.

## Scope: the third declared step was NOT closed

[`overview.md`](overview.md) declared three steps and labelled the third **investigate-or-verdict**:
`pt-hiring-recruiter-compare`. Two landed; the third is routed forward with its shape priced (the seed-pinned
cardinality final that is available, and the two candidate sources for the *absence* half that must be measured
before either is used). Per the iter's own escalation condition — *"if (1) and (2) land but (3) turns into a
build, route (3) forward rather than opening a fourth line"* — and because this iter had already spent an hour
on an unplanned stack diagnosis, opening it would have been the scope-creep the tripwire names. Hence
`closed-fixed-partial`, not `closed-fixed`.
