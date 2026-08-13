**Type:** tik (under `TOK-01: instrument first, then follow`, steps 3 + 5).

# iter-17 — clause 1 was never met, and now there is a measurement that says so

## Outcome in one line

The hypothesis was **falsified on the first cold cycle**, and the falsification is the deliverable: with an
instrument that can see served content, a cold `demo-down --purge` → `demo-up` reaches
**`green:false / warnings:1`**, failing on `anon GET /items/task_sub_checks -> 403 — the running Directus
holds the content but serves it to nobody`. Gate clause 1's three checked-in verdicts are **withdrawn**.

## The positive control, which was not a formality

Before spending three cycles, the probe was run against the standing (hand-repaired, therefore
known-serving) `demo-1`. This mattered: **a probe that no-ops produces three green cycles that mean exactly
as little as the last three**, and §5 rule 8 — *a check that SKIPS reads exactly like a check that PASSES* —
is the milestone's most-repeated finding.

The first look was alarming and wrong. `autoverify`'s output showed no `directus-serves-content` line at
all; the only textual match was the registry check's own cross-reference to it. Running `verify.sh` directly
resolved it — `autoverify` captures verify's output and prints only the `✗` rows, so a passing probe is
invisible by design. The probe **ran and passed**:

```
readiness directus-serves-content  ok  ok: anon GET /items/task_sub_checks served a real item
```

Worth keeping as a method note: *"I did not see it in the log"* and *"it did not run"* are different
findings, and the cheap way to tell them apart is to run the inner tool directly rather than reason about
the outer one.

## A wrong entry point, caught by arithmetic rather than by suspicion

The first teardown/bring-up used `rosetta-demo up 1`, which returned "up" in **30 seconds** against
iter-14's measured ~11 minutes. The log said `3 services, profile='base'` — `rosetta-demo up` is the bare
compose bring-up; `/demo-up` runs `up-injected.sh` (Clerkenstein injection, full UI tier, set-dress,
autoverify). Torn down and re-run through the real path.

Nothing about the bare stack would have *looked* wrong: it comes up, containers run, and an autoverify
scoped to it would have had less to fail on. `demo-up-defaults.md` already records that there are **two
entry points** and that the skill's own `argument-hint` had conflated them; this is that fact costing a
cycle. The tell was the **duration**, not the output — the same shape as `latency-budget.md`'s arithmetic
signatures.

## What the honest cycle found

The cold cycle is genuinely informative in both directions, and it landed on the informative side:

```
⚠ [directus] bootstrap failed — skipping local content (the stack stays on the prod-read path)
provisioned "directus" structure into demo-1 (schema b4cb55bcee08 → ea2e187a1605)
replayed "directus" into demo-1: 14 table(s) cleared, 11986 row(s) loaded, schema identity, 3 sequence(s) advanced
==> demo-1: set-dressed (content:prod-read, snapshot:taxonomy=replayed directus=replayed …, stories seeded).
✗ directus-serves-content  fail: anon GET /items/task_sub_checks -> 403 …
```

**First, the good news, and it is real:** iter-15's replay fix **works through the real bring-up path**. Its
own hand-off flagged this as unsettled — the fix was proven at its own layer and had never run inside
`demo-up`, because the auto-provision path fires on a bootstrapped-GAP schema, which is what a *purged*
stack has and what the standing demo-1 no longer had. `directus=replayed`, 11,986 rows, rc=0. That question
is now closed.

**Second, the defect that was hiding behind it.** The per-stack Directus **bootstrap** failed. The pass then
announced it was falling back to prod-read — and **replayed 11,986 content rows into the local Directus
anyway**, reporting `directus=replayed`. The closing verdict says, in one sentence,
`content:prod-read` *and* `directus=replayed`. Those cannot both be the operative truth, and the sentence
asserts both while exiting 0.

The consequence is precisely the 403: the content is in the tables, but the system schema those tables need
in order to be *served to an anonymous reader* was never bootstrapped, so the public-role grants were never
applied. The registry check sees 21 registered collections and passes. The serving probe asks for an item
and gets 403. Both are correct; only one of them was ever asked before today.

### Why the bootstrap failed is, right now, unknowable — and that is the third occurrence of one class

`dev-setdress.sh:258-265`:

```bash
if ! docker run --rm … "$img" node cli.js bootstrap >/dev/null 2>&1; then
  echo "    ⚠ [directus] bootstrap failed — skipping local content (the stack stays on the prod-read path)" >&2
```

`>/dev/null 2>&1`. **This is RF-1's exact shape in a third file** — the same *discard the output, then report
a failure the operator cannot act on* that iter-16 fixed in `migrate-dev.sh` hours earlier, and that M215 F8
fixed in `migrate-demo.sh` before that. Re-running the identical `docker run` by hand now returns:

```
INFO: Initializing bootstrap...
INFO: Database already initialized, skipping install
INFO: Running migrations...  INFO: Done
```

— i.e. **it succeeds now**, because the replay has since created the `directus_*` tables from the capture,
so the state that failed no longer exists. The diagnosis was available for exactly one moment and was
written to `/dev/null`. That is the whole argument for capturing it.

## Why cycles 2 and 3 were not run

The clause requires **three consecutive green** cycles. Cycle 1 is red, so the outcome is determined; two
more would take ~22 minutes and cannot change it. Running them anyway — to have a table with three rows in
it — would be the same instinct that produced the three withdrawn files. Recorded here so the absence reads
as a decision rather than as an interruption.

## Close — 2026-08-01

**Outcome:** Clause 1 falsified on the first cold cycle with the honest instrument, and its three
checked-in verdicts formally withdrawn (`evidence/av-cycle1.json.WITHDRAWN.md`). Root cause named and cited:
the per-stack Directus **bootstrap** fails, the set-dress pass reports `content:prod-read` while replaying
11,986 rows into that same Directus and calling it `directus=replayed`, and the anon read 403s because the
grants the bootstrap would have installed were never installed. iter-15's replay fix is **confirmed working
through the real bring-up path** — the question its hand-off left open.
**Type:** tik
**Status:** closed-no-lift (documented falsification — the protocol's first-class outcome; the planned
investigation completed and returned falsification evidence, and no fix attempt was landed)
**Gate:** NOT MET — and this iter **moves clause 1 from "met" to "not met"**. That is a correction, not a
regression: the clause was never met, and 2-of-5 was an over-count from iter-14 onward. The honest standing
is **1 of 5** (clause 4).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this is the 1st no-prog tik of a streak, not the 3rd) — (3) re-scope: n (platform origin HEAD re-checked at open and at close: `2adcf71`, unchanged — occurrence stays 1 of 2) — (4) user-blocker: n (the fix is routed with a named handler; no decision needed that changes what code lands) — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D-M257x-17-1 … D-M257x-17-3 (iter-local `decisions.md`).

**Metric:** clause 1 `green:true`×3 → **`green:false / warnings:1` on cycle 1**
(`evidence/av-iter17-cycle1.json`, `16:03:35Z`). Gate: **1 of 5** clauses (was recorded 2 of 5).

**Side-deliverables:** none.

**Routes carried forward:**
- `FIX-M257x-iter17-directus-bootstrap-blind` — capture + classify `node cli.js bootstrap`'s output at
  `dev-setdress.sh:258-265` exactly as iter-16 did for `migrate-dev.sh`, and decide what the pass should do
  when bootstrap fails: today it says "skipping local content" and then replays into it anyway. **The
  strongest candidate is that these are one bug, not two** — but that is a hypothesis, and the diagnosis
  must exist before it can be tested. → **iter-18, first thing.**
- `CHECK-M257x-iter17-setdress-verdict-contradiction` — `set-dressed (content:prod-read, … directus=replayed …)`
  asserts two incompatible states in one line and exits 0. iter-16 made the verdict a function of the
  *replay* outcomes; the **provision/bootstrap** outcome is not yet an input to it. → iter-18 (likely the
  same fix).
- `CHECK-M257x-iter17-two-entry-points` — 30 s vs ~11 min. `demo-up-defaults.md` documents the two entry
  points; nothing *warns* when the bare one is used where the injected one was meant. → later tik.

**Lessons:**
1. **A clause is only as true as the weakest probe behind it.** Clause 1 was declared met on three honest
   cycles of a blind instrument. Nothing was faked; the check simply did not ask the question. Before
   claiming a gate clause, write down what each probe would have to be blind to for the claim to be false —
   iter-14 could have asked *"what does green mean if the transcript says `directus=skipped(error)`?"* and
   the two were sitting in the same log.
2. **Un-measuring a clause is progress.** The milestone is worth less with a false 2-of-5 than with a true
   1-of-5, and every downstream decision made under the false count was made on bad information.
3. **`>/dev/null 2>&1` on a step whose failure you then report is a defect with its own signature** — third
   occurrence in this milestone across three files. The failing state is usually transient; by the time
   anyone investigates it has healed or moved, and the one moment the diagnosis existed is gone. Worth
   promoting to a grep-able rule (protocol addendum, iter-18).
4. **Duration is a measurement.** The wrong entry point produced a *plausible* stack and was caught by 30 s
   vs 11 min, not by anything in its output.
