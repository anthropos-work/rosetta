**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*). Closes
`SURVEY-M257x-iter143-bare-orphan-bucket`, the route iter-142 recorded at its close (`4edad03`) as
this iter's starting point. Three planned lines (census → derive+audit → ship what the audit
supports), declared in this iter's `overview.md`.

# iter-143 — the `(bare)` bucket censused, head inference REFUSED, and the audit that was itself wrong

## The target, and why it was next

`TOK-08`'s pre-registered class list has two classes. Class 1 (intra-corpus citation resolution) is
fenced and GREEN. **Class 2 (platform-source citation resolution) is fenced at 59.1 %** —
`anchor_construct_guard` resolved **861 of 1,456** — and the largest single unresolved head was
**`(bare)` at 384**, one opaque number with a source comment explaining it:

> *"that orphan bucket is where `:5050`, `:8082`, `:7700` and every other port lives … resolving it
> would be the 134-findings first draft returning."*

`4edad03` had already measured that comment as **only marginally true** (38 of 621 bare matches carry
a known port). This iter censused the bucket, built the inference the comment says must not be built,
and measured whether the comment is right for the right reason.

## 1. The census — stated with its denominator (iter-114's rule), and stated TWICE

| refusal reason | at iter open | at iter close |
|---|---|---|
| `orphan-no-referent` — no citation and no self-marker anywhere before it in scope | **337** | 333 |
| `ambiguous-address` — a second `:NNN`-shaped referent sits between anchor and citation | 28 | 26 |
| `ambiguous-file-mention` — the prose switched files without re-citing one | 15 | 17 |
| `superseded-quote` — the sentence quotes the anchor in order to retract it | 4 | 4 |
| **total (= the `(bare)` head the same run prints)** | **384** | **380** |

**Both readings are published because this iter moved its own denominator** — the `_CODE_SUFFIX`
widening below made citations visible that anchors then inherited from, and the protocol-doc edit
added one. Quoting only the closing figure would present an intervention as a measurement. **The
shape is what carries: ~88 % of the bucket is `orphan-no-referent`** — not ambiguity the guard is
right to refuse, but referents it cannot see. That is the only arm any head inference could address.

## 2. The inference — built, hand-read at 100 %, and NOT shipped

`_FILE_MENTION` is a regex the guard **already has**, used only to BREAK a chain. It also describes
exactly how the corpus NAMES a file without citing it — *"all anchors in `handler.go`: `session` …
(`:1458`)"*. Letting it also START a chain places **129** of the 337 orphans, **92** of which resolve
to a real file on disk. Every one of the 92 was read in full line context before anything shipped
(`D-M257x-142-1`, in its own words).

| rule | precision | recall |
|---|---|---|
| mention referent, as drafted | **57.6 %** | 100.0 % (53 of 92) |
| + parenthesised-aside only | 74.5 % | 71.7 % |
| + a mention may START a chain, never HIJACK one | 51.8 % | 54.7 % |
| **+ both** | **77.3 %** | **32.1 %** |

**77.3 % is not fence quality**, and this milestone has already retracted one mechanical publication
over this exact population (iter-138, audited **0-for-12** at iter-139). **The inference is not
shipped.** That is a real answer to the standing `FIX-M257x-iter138-anchor-rot-fence` re-specification
— which asked for *head resolution first*, and is entitled to the answer **"not mechanically decidable
at fence quality on this construct, and here are the numbers."**

A numeric cut (`n < 3000`) separates the population perfectly at the open and was **declined**: it is
a tuned constant on a 92-site denominator, and **two such constants are already routed open**
(`-iter142-path-arm-window`, `-iter142-tier-b-underflag`). A third would be a pattern, not a fix.

## 3. ⚠️ The headline: the AUDIT was wrong, and no second reading could have told me

The iter did what iter-142 taught — audited before publishing. The reader's verdict was **62 true /
30 false**, and the paren-aside predicate scored **90.2 %**. Then the 62 "true" sites were pushed
through the guard's own `classify()`, and **nine came back `anchor-out-of-range` against files whose
line count makes the citation impossible.**

All nine are **one** mechanism, and it is one a reader is structurally poor at seeing: a bare file
**mention** sitting *nearer* to the anchor than the **qualified citation** actually governing it.

* `cms.md:125` — *"`studio/gen.py` at `studioManager.go:119` and `studio/postgen.py` at `:1045`"*.
  `:1045` is a **`studioManager.go`** line. `postgen.py` is simply the nearer noun.
* `demopatch-spec.md:165/234/235` — eight anchors that belong to **`up-injected.sh`** (cited two lines
  up as `up-injected.sh:1124-1126`), hijacked by the phrase *"the shared `urls.ts` pair"*.

Corrected, the population is **53 true / 39 false** and the same predicate scores **74.5 %**. **The
audit had inflated it by 15.7 points**, and in the direction that ships.

**And the mechanism split matters more than the rate.** Of the 39 false admits:

| mechanism | n | visibility when wrong |
|---|---|---|
| **port** — a network port in a bare anchor's exact token shape | **21** | **loud** — resolves out-of-range, shows up as a RED |
| **WRONG-HEAD** — a real line anchor booked against a file the sentence never named | **16** | **silent** — can land on a real construct and PASS |
| teaching illustration — a quoted idiom in `platform-alignment.md`, not a citation | 2 | n/a |

The guard's source comment named **only the loud half**, for five iters. Refusing the bucket is still
correct; the comment is now retracted **as an explanation** and replaced with the measured one.

## 4. What DID ship — the two gains that involve no inference at all

**(a) `_CODE_SUFFIX`, measured for the first time since iter-73.** Over the whole corpus there are
**32** backticked `name.EXT:NNN` citations with **no slash** whose suffix the list does not carry — so
`_QUALIFIED` cannot see them and they are **invisible** to the guard, not merely unresolved. Seven
suffixes added by measured count (`jsx` x8, `js` x8, `graphqls` x4, `ini` x2, `txt` x2, `hcl` x1,
`sum` x1). **Three were measured and DECLINED**, and the third is the counter-example that prices the
whole list:

| declined | why |
|---|---|
| `dev` (`Dockerfile.dev:18` x3) | `.dev` is also a TLD — the suffix would admit hostnames |
| `example` (`.env.example:120` x2) | admits any `word.example`; `example.com` is the canonical one |
| **`de`** (`u422950.your-storagebox.de:23`) | a **HOST:PORT that only looks like a citation** — admitting it books a storage-box hostname as a file, at a line number that is really a port |

Result: **reach 861 → 875 anchors**, and every one of the 14 newly-graded citations came back clean.
This is safe where the head inference is not, for one reason: **a qualified citation carries its own
path, so admitting it decides nothing the prose had not already said.**

**(b) The census is now REPORTED**, in the text run and in `--json`, with `BARE_REFUSAL_REASONS` as a
single constant so the three spelling sites cannot drift apart.

## 5. A live demonstration, unplanned and worth recording

This iter's own protocol-doc edit **quotes** the `studioManager.go` / `postgen.py` sentence as an
example. That quotation put a bare `` `:1045` `` into a scanned document — and the guard's existing
`ambiguous-file-mention` refusal caught it: the count moved 16 → 17 and the anchor was **not graded**.
The very ambiguity rule this iter measured as *costing recall* is what stopped the iter from inducing
a finding while writing about inducing findings.

## 6. ⚠️ The owed whole-suite run went RED — on THIS iter's own change

`FIX-M257x-iter142-whole-suite-owed` said a whole-suite run was owed before the next rext-touching
iter quotes coverage. It was run, and it returned **31 failed / 1,261 passed / 1 skipped**.

**30 of the 31 were mine.** The census had shipped as an **eighth member of `run()`'s return tuple**,
and `test_iter45_mechanical_fences` unpacks that tuple **positionally at six call sites** — every one
raised `ValueError: too many values to unpack (expected 7)` (27 tests), and the mutation battery that
runs that module cascaded 3 more, reporting *"the suite is RED before any mutation."* The 31st is the
milestone's **standing, documented RED** (`test_claim_twin_guard_iter48_answer_key::test_02`),
re-attested by a full run rather than carried.

**The scoped, change-derived suites this iter picked did not include that module.** I selected three
(`repair_postcondition`, `fence_provenance`, `guard_family`) on the reasoning *"these consume the
changed return value"* — 106 passed, and that green was real and irrelevant. `test_iter45_mechanical_fences`
consumes it too, and is the module that unpacks it positionally.

> **`run()`'s ARITY is a published interface.** Adding a member to a returned tuple is a breaking
> change, not an addition. The census now rides on a **module-level accumulator** (`BARE_REFUSALS`),
> which is this file's **own existing idiom** — `RESOLVE_ROUTES` and `NOT_CITATIONS` are both exactly
> that, both cleared by `run()` at entry. The right design was already in the file; the tuple was the
> lazy reach.

Two named regression tests ship: the arity is **pinned** (`len(run(root)) == 7`), and the accumulator
is asserted **per-run, never cumulative** — because an accumulator that is not cleared at entry doubles
on the second run, and the arithmetic control would then pass on a wrong number.

**This is what the owed run was for, and it is the second time in two iters that the thing which
caught a real miss was a check on a different axis from the one being changed** — iter-142's was
`repair_leak_guard` over the commit; this one is the whole suite over the tree.

## Test gates

| gate | result |
|---|---|
| `test_anchor_construct_denominator.py` (**+13 net-new**: 4 reason mutation controls · 2 anti-vacuity · 1 arithmetic control on the real corpus · 4 suffix admit/refuse · **2 arity/accumulator regressions**) | **26 passed / 0 failed** |
| `test_iter45_mechanical_fences.py` — the 27 the tuple change broke | **83 passed / 0 failed** after the fix |
| Consumers picked as the change-derived scope — `test_repair_postcondition`, `test_fence_provenance`, `test_guard_family` | **106 passed / 0 failed** — green, real, and **insufficient**; see § 6 |
| `anchor_construct_guard --repo-root .` | **GREEN**, exit 0, before and after every edit |
| **Whole rext `stack-core` suite** — the run `FIX-M257x-iter142-whole-suite-owed` demanded | **first run 31 failed / 1,261 passed / 1 skipped** (30 mine, 1 standing) → **re-run after the fix: 1 failed / 1,294 passed**, the 1 being the standing `test_claim_twin_guard_iter48_answer_key::test_02`, re-attested by a full run rather than carried. Collection reconciles: 1,293 → 1,295, the +2 being this iter's arity/accumulator regressions. Quoted, not assumed |
| Suite wall-time | **not quoted as a measurement**, and this run is the standing evidence for why: 761.15 s then 1,340.91 s for the *same suite on the same host within the same hour*, because an unrelated 100 %-CPU `pytest` from another project was running through the second. `FIX-M257x-iter132-suite-walltime-is-not-a-measurement` stands, and **counts are what this ledger quotes** |

## Close — 2026-08-08

**Outcome:** the `(bare)` orphan bucket is **censused** (384 → 380 refusals, ~88 % `orphan-no-referent`),
the head inference it invites was **built, hand-read at 100 %, and REFUSED at 77.3 % precision /
32.1 % recall**, and two **no-inference** gains shipped instead — a `_CODE_SUFFIX` widening that took
reach **861 → 875** and the census decomposition itself. **The headline is not the census: the hand
audit was wrong on 9 of 92**, all one mechanism (a bare file *mention* nearer the anchor than the
*qualified citation* governing it), and what caught it was `classify()` — the machine that was going
to consume the verdict — not a second reading. **The owed whole-suite run then caught a second one on
the same axis**: the census had shipped as an eighth member of `run()`'s return tuple and broke 27
tests in a module that unpacks it positionally. Guard **GREEN**; `§5` gains **rules 65 and 66**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged. **No `N` reading was taken, so no `N` movement is claimed**
(`§9`'s guard-rail 1, in its required words); `TOK-08` sequences the read after the mechanical sweep,
and this is one class of that sweep rather than its completion.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**iters 135–143 took no reading, so the metric is UNMEASURED not unmoved — `§9`'s iter-type refinement; and `TOK-08`'s sealed refutation branch bars an agent-authored successor in any case**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (**1 tik this session**) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-143-1` (**the headline** — an audit is a predicate too, and needs a control
that is not another reading) · `D-M257x-143-2` (head inference REFUSED, with its number — an answer to
the standing route, not a failure to attempt it) · `D-M257x-143-3` (count the LOUD and the SILENT
failure modes separately; decline a widening for the hazard that cannot be seen when it is wrong) ·
`D-M257x-143-4` (a reach gain requiring no inference is a different KIND of change) ·
`D-M257x-143-5` (publish a census that moved with BOTH readings) · `D-M257x-143-6` (a returned tuple's
ARITY is a published interface; derive a change-scope from CALL SITES, not from imports).
**Side-deliverables:** none.
**Routes carried forward:**
- **`SURVEY-M257x-iter144-orphan-arm-is-the-residual` (NEW)** — `orphan-no-referent` is **333 of 380**
  and is now the whole of class 2's remaining gap. This iter proved it is **not** closable by
  mention-inference at fence quality. What has NOT been tried: a **sourcing** approach — resolving the
  head from the document's own service/section context rather than from the sentence. ⚠️ Carry
  `D-M257x-143-1` with it: any candidate needs a control that is not a reading.
- **`FIX-M257x-iter143-wrong-head-is-unfenced` (NEW)** — the **16 wrong-head** false admits are a
  class **nothing currently detects**, and by construction they are the silent half (a real line
  anchor booked against a file the sentence never named can land on a real construct and PASS). They
  were found only because this iter built an inference and audited it. **The class is now named and
  measured but unfenced.**
- **`FIX-M257x-iter143-scope-derivation-by-grep` (NEW)** — rule 66 states the rule; nothing enforces
  it. A cheap fence would be a pre-commit check that any changed public function's **call sites** are
  in the iter's declared test scope.
- `FIX-M257x-iter142-value-change-articles` · `FIX-M257x-iter142-path-arm-window` ·
  `FIX-M257x-iter142-tier-b-underflag` (the two tuned constants — and iter-143 **declined to add a
  third**, `n < 3000`, which would have "worked").
- **`FIX-M257x-iter143-appending-to-the-protocol-doc-rots-the-ledger` (NEW)** — measured at this
  iter's close, and it is an exposure nothing watches. Rules 65+66 added **72 lines** to
  `platform-alignment.md` (3,360 → 3,432) at ~`:2220`. **Nine pins into that file sit below the
  insertion point** and are now off by 72: `iter-129/{readset,slice-F}.txt` (`:2324`, `:2356`,
  `:2774`, `:2925`, `:2967`), `iter-139/rot-census.txt` + `decisions.md` + `progress.md` (`:2315`,
  `:2595`, `:2596`) and `iter-132/progress.md` (`:2618`). **All nine are in `knowledge/plan/**`**,
  which `anchor_offset_guard`'s scope deliberately excludes — and correctly reported *"5 changed
  files, 0 of them cited"* over the published tree, where the count really is **zero**. They are
  frozen evidence artifacts of past iters, dated and provenanced, so **repairing them would be
  rewriting past evidence and is NOT proposed.** But the mechanism is general and recurring: **every
  iter that appends to the protocol doc rots every pin into it held by every prior iter's records,
  and nothing measures it.** Either the ledger is in scope for offset-grading (with the historical
  arm exempted by path, as `retracted_pin_guard` exempts teaching) or the exposure is disclosed
  standing. Right now it is neither.
- `FIX-M257x-iter135-adjudicated-live-defects` (non-citation remainder) ·
  `FIX-M257x-iter140-receipts-not-checkable-here` · `FIX-M257x-iter140-receipt-fence` ·
  `FIX-M257x-iter138-anchor-rot-fence` (**re-specified by this iter**: head resolution over the bare
  bucket is not mechanically decidable at fence quality; the route now needs a different mechanism,
  not another attempt at this one) ·
  `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer` ·
  `FIX-M257x-iter133-two-fives-need-a-fence` · `FIX-M257x-iter132-suite-walltime-is-not-a-measurement`
  (**re-attested here**: 761 s then 1,341 s for the same suite on the same host within the hour) ·
  `FIX-M257x-iter131-predicate-sets-not-enumerated`.
- **CLOSED by this iter:** `SURVEY-M257x-iter143-bare-orphan-bucket` · `FIX-M257x-iter142-whole-suite-owed`.
**Lessons:**
0. **An audit is a predicate too, and it needs a control that is not another reading.** iter-142's
   *audit before you publish* is right and this iter obeyed it — and was still wrong on 9 of 92, in
   one direction, on one mechanism. The cheapest control was already in the tree: push the reader's
   own TRUE set through the machine that will consume the verdict. It cost one script and moved the
   published precision by **15.7 points**.
1. **Grade a false positive by whether you could SEE it.** 21 ports (loud) and 16 wrong-head (silent)
   are not 37 errors of one kind. A justification that covers only the loud half invites exactly the
   re-litigation that produced this iter.
2. **A refusal with a number is a deliverable.** *"Not mechanically decidable at fence quality, here
   is the precision, here is the recall, here is what was tried"* answers a standing route. *"We did
   not get to it"* does not.
3. **Arity is a published interface.** Adding a member to a returned tuple is a breaking change
   wearing the costume of an addition — and the file already had the right idiom (a module-level
   accumulator) ten lines above the edit.
4. **The check that catches you is never the one you designed while making the change.** Three
   instances in two iters: a guard on another axis (142), `classify()` over the reader's own verdict
   (143), the whole suite over the scoped set (143). Everything *inside* the thing being built was
   green in all three.
