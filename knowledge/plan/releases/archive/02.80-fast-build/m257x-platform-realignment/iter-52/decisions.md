# iter-52 — decisions

## D-M257x-52-1 — the union, not the 14, and the fence agrees it was the right target

`claim_twin_guard --report` at open: **RED, 31 hits / 19 unique sites**. After the repair of the 18:
**RED, 6 hits / 4 unique sites**. The fence is derived from the ledgers by table *structure*, so it had
already absorbed reading #10's rows — which is independent corroboration that the union, not iter-49's 14,
was the correct repair target.

## D-M257x-52-2 — REFUTED: my own pre-registration of "net words added ≤ 0"

iter-52's `overview.md` pre-registered **net words added ≤ 0 across the pass** — the measurable form of
TOK-03 move 3. Measured on the repair diff:

| | words |
|---|---|
| added (`+` lines) | **1472** |
| removed (`-` lines) | **675** |
| **net** | **+797** |

**The prediction is REFUTED on its first outing.** Recorded here rather than softened, because a
pre-registration that is quietly restated after the measurement is worth nothing. Three honest observations:

1. **Deletion was available far less often than TOK-03 assumed.** Of the 18, only **two** — the
   `flag_use_realtime_openai` leak at `external_services.md:672` and the roadrunner "batch" clause — could
   be repaired by removing a clause and leaving a true sentence behind. The rest were **wrong values inside
   load-bearing sentences** (a count, a variable name, a function name, a column set). Deleting those
   removes the reader's answer along with the error.
2. **§5 rules 17/24 actively push the other way.** *"Where a list is load-bearing, ship the derivation
   command next to it so the next reader re-derives instead of trusting"* is a rule that ADDS words, and it
   is the rule that would have prevented the 31-vs-32 defect. Move 3 and rule 24 are in genuine tension and
   TOK-03 did not notice.
3. **The word budget was the wrong shape of metric.** What move 3 is actually trying to bound is *new
   assertions* — claims a future reading could book — not *characters*. A 40-word derivation command adds
   no assertion; a 6-word confident adjective adds one. **The next tok should re-cut this as a count of
   NEW ASSERTIONS, and iter-53's paired reading is what can measure it.**

## D-M257x-52-3 — the ratchet REFUSED the repair, and it is right to, for one of the four sites

`repair_postcondition.py --staged` (the installed pre-commit hook) returns **RED on 6 hits / 4 sites**.
The orchestrator's brief expected the normal path to work because this run is *audit → repair*; the mode
is not the problem. The mechanism is.

**What is actually happening.** `claim_twin_guard` matches a ledger's **quoted span** as ordered fragments
inside a normalized window. A repair that keeps the sentence and changes the wrong value therefore keeps
most of the quoted span — so **the fence fires on the CORRECTED text**. The four sites split two ways:

| site | class | judgement |
|---|---|---|
| `hiring.md:117` | **genuine retraction quote** — the repair quotes the withdrawn claim (*"said it 'cannot write the 5 positions in the first place' until iter-52"*) in order to retract it | exactly the case the waiver file exists for; same class as its 3 existing entries |
| `external_services.md:569` | **stem collision** — the matched fragments are the sentence's generic stem (*"things can send a request outside the EU, none of them a region-health failover"*); the distinguishing token (the count) is the part that changed | not a restatement |
| `ai-readiness.md:420` | **stem collision** — the matched sentence is now TRUE of the branch the repair re-scoped it to (`keepStartedMembers`, the `cyc == nil` branch); what was false was the scoping, which sits above it | not a restatement |
| `hiring.md:303` | **stem collision** — the matched span is the unchanged paragraph OPENING; the refuted clause sits further down and was repaired | not a restatement |

**The decision-relevant finding, and it is new:**

> **TOK-03 move 3 and this fence are in direct opposition. The more MINIMAL the repair, the more of the
> refuted span survives, and the more likely the fence is to fire on correct text.** A wholesale rewrite
> would have dodged it. **The instrument rewards the riskier repair shape** — the exact shape iter-49
> measured as the source of the induced defects.

This is recorded, not fixed by loosening anything. **No `--no-verify`.** The handling is the fence's own
designed second key: a waiver is honoured *only* while `_looks_retracted` finds a retraction marker within
320 characters of the quote, and *only* while the three answer-key suites still detect all their known
blockers. Both keys must hold or the site stays RED — which is the property that makes this safe to use and
would make it unsafe to widen.

Routed forward as **`FENCE-M257x-iter52-stem-collision`**: the matcher should require at least one fragment
that is *distinguishing* — present in the refuted form and absent from the adjudicated correction — before
it can fire. That is a fence change and needs its own RED-watch (§8 rule 7), so it is not done inside a
repair pass.

## D-M257x-52-4 — the diff was frozen while the two blind readers read it

§8 rule 7's recurrence corollary, and the one iter-49 violated against itself: **five of its own tests were
invalidated mid-session by its own repair.** TOK-03 move 4 puts two blind adversarial readers on the repair
diff *before* the commit; the diff they are given is a file on disk. Editing the corpus while they read it
would silently invalidate both reports and neither would say so.

So the ratchet's refusals were **diagnosed and written up here — inside `knowledge/plan/**`, which both
readers are barred from — and no corpus byte was touched until both reports were in.**
