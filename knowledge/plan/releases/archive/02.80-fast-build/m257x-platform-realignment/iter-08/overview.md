---
milestone: M257x
iter: 08
iteration_type: tik
status: archived
opened: 2026-07-31
---

# iter-08 — `FENCE-M257x-write-fence-scans-one-section-of-nine`

**Active strategy reference:** `TOK-01` — *instrument first, then follow*, step 3: **land the fences, each
watched going RED, before trusting any green.**

## Step 0 — re-survey, and it REFUTED most of the target

iter-07 routed this forward claiming the write-target fence's scope limit was undocumented and that the
fence "would not have caught the defect iter-07 just fixed." **Re-measurement refutes the first claim
outright and reduces the second to something much narrower.** Recorded here before any work, because
building on it would have produced a fix for a problem that does not exist.

**Refuted — "nothing says so".** `test_write_target_schema_fence.py` carries a **ten-line justification
comment directly above `SCORED_SECTIONS`**, and it is a good one:

> *"`stack-snapshot` is deliberately OUT of scope, and the reason is worth stating rather than leaving as
> an omission. Its replay surfaces name schemas that rext's migrate step does not create and should not:
> `directus` is created by Directus's own `node cli.js bootstrap`, and `ref` is rext-owned. Scoring them
> would make the fence permanently RED, and the natural fix — allow-listing them — is exactly protocol
> Trap A's warning…"*

iter-07 read line 92 and the module docstring, and did not read the ten lines between them. **That is the
milestone's own dominant defect, committed by this milestone's own iter** — a claim reported without being
measured — and §5's closing rule names it exactly: *verify a claim before escalating it, including a claim
made by an audit.*

**Refuted — "widen the tuple".** Running the fence's **own scanner** over every Go-bearing section:

    stack-seeding    92 write-construct hits   schemas=[public]   ILLEGAL=0
    stack-snapshot    0 hits
    stack-secrets     0 hits
    clerkenstein      0 hits
    alignment         0 hits
    playthroughs      0 hits

Widening `SCORED_SECTIONS` would score **nothing** and report GREEN — the precise trap iter-07's own
pre-compute predicted, now measured rather than predicted.

**Also refuted — the vacuity gap I expected to find.** `test_the_seeders_actually_write_something` already
asserts `found > 40`, so the fence cannot pass by scanning nothing. iter-06 had thought of it.

## What survives, and it is real

**1. The exemption's REASONS are now stale, and they will mislead the next reader exactly as they misled
iter-07.** The comment names `cms.similarities` as *"its one genuinely stale target"*, says the covering
signal is that it *"already fails LOUD at replay time … rc=4"*, and points at
`REPOINT-M257x-cms-similarity-writes` as the tracker. **iter-07 closed that route and removed that rc=4
signal** — the replay now resolves and succeeds. A comment whose stated safety argument has been deleted is
worse than no comment: it reads as a live justification.

**2. Nothing maps a section to which of §8's THREE layers covers it.** The protocol declares three layers
(map↔`repos.yml` · static write-target fence · live `information_schema` assert). The question *"is
`stack-snapshot` fenced?"* has a correct answer — **the live layer, because after `D-M257x-8` its write
target is DERIVED at run time and there is genuinely nothing static to fence** — and that answer is written
down nowhere. iter-07 got it wrong by guessing. Clause 4 currently rests on a reader making the same guess
correctly.

**3. The scope set is hand-maintained and unfenced — the milestone's own recurring shape.** `SCORED_SECTIONS`
is a tuple someone must remember to extend. A NEW rext section (v9.0 is already adding surface) would appear
uncovered and **silently**, because the fence only ever asserts about what it already scans. That is the
same defect class as the migrate tuple iter-02 removed, one level up: *the fence's own scope is a
hand-maintained list of the system's parts.*

## Hypothesis

Derive the fence's **coverage** the way iter-02 derived the migration set: enumerate the Go-bearing sections
from disk, require every one to carry a declared layer + reason, and go RED when a section appears that
nobody has classified. Then refresh the stale rationale, and clause 4 can be claimed against a written,
machine-checked statement of what each layer covers.

## Expected lift

- The fence's scope stops being hand-maintained-and-trusted (`§2`'s governing rule, applied to the fence).
- The three-layer coverage claim becomes explicit and testable, so **gate clause 4 is claimable honestly**.
- iter-07's incorrect finding is corrected in the record rather than inherited.

## Phase plan

1. **Phase A** — correct the iter-07 record and the milestone route.
2. **Phase B** — derive the section-coverage set + declare a layer & reason per section.
3. **Phase C** — fence it; mutation-verify (build the mutant first, per `§8 rule 5`).
4. **Phase D** — refresh the stale exemption rationale; update `platform-alignment.md` §8's layer table.

## Escalation conditions

If the coverage derivation shows a section that genuinely needs static scoring and has constructs the
scorer cannot see, that is a scorer-widening job with its own trap surface — route it forward rather than
half-widen.

## Acceptable close-no-lift outcomes

The refutation itself is a first-class outcome: if the coverage fence proves unnecessary, the corrected
record plus a written layer map is the deliverable.
