# iter-156 — progress

**Type:** tik · **Strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07)

## Phase A — the census, and its denominators

`SURVEY-M257x-iter152-other-guards-may-read-prose-as-data` asked for the population, derived from disk,
graded on the PROPERTY (*the marker string does not occur in the scanned tree outside its structural
position*) and explicitly **not** on anchoring. Three readings were taken, each with its denominator.

### A1 — the marker population, derived by introspection (not by grep)

Every `stack-core/*.py` was imported and its module-level `re.Pattern` attributes enumerated. This is
derived rather than listed, so a marker added tomorrow is in the population without anyone editing a list.

| reading | number |
|---|---|
| guard modules on disk | **32** |
| modules that read a prose tree (`corpus/**` or `knowledge/plan/**`) | **23** |
| module-level compiled patterns in those 23 | **171** (180 across all 32) |

⚠️ **The first run of this instrument reported 13 of 32 modules as `IMPORT-FAIL` — and that was MY loader,
not a defect in any guard.** `@dataclass` resolves its own module through `sys.modules`, which a
`module_from_spec` + `exec_module` pair does not populate. Registering the module first cleared all 13.
Published as a finding it would have been thirteen false defects; it is recorded here because the near-miss
is the same shape as iter-150's 30-to-1 and iter-138's withdrawn 127 — **the instrument is a claim too.**

### A2 — the structural-position reading over the whole marker set

Each of the 171 patterns was run over its own guard's declared prose trees, and every match classified by
position (fenced code block · blockquote · inline-code span · prose).

**Raw signal: 105 of 171 patterns flagged.** It is almost entirely noise, and the reason is the
denominator: the population contains *helpers* (`_WS`, `_DROP`, `SENTENCE_SPLIT`, `_CAVEAT_CUT`) that are
applied to already-segmented text, and patterns whose subject is a compose file or a `go.mod` and which
were graded against markdown only because this instrument binds patterns to trees per-GUARD, not per-use.
**A naive publication of "105 guards read prose as data" would have been this milestone's third false
number.** It is not published as a defect count; it is reported as the raw signal it is.

### A3 — the four markers that could be bound by hand, graded exactly

For the four markers whose use-site could be read directly, the property was graded per match:

| marker | tree | files | matches outside prose |
|---|---|---|---|
| `blocking_state_guard._GRADING_HEAD` | milestone dir | 794 | **0** (142 matches, all prose) |
| `derived_value_guard._DOC_GO` | `corpus/services` | 29 | **0** (7 matches) |
| `derived_value_guard._DOC_SIZE` | `corpus/services` | 29 | **0** (1 match) |
| `evidence_visibility_guard._CITATION_RE` | `knowledge/plan` | 2285 | **0** (15 matches) |

And the first/last-match ambiguity that iter-152's fix resolves: **0 of 140** iter `progress.md` files carry
more than one `^**Phase 5 grading:**`. The instrument is proven against the same corpus **unanchored**,
where it finds the two files iter-152 named (`iter-150`, `iter-152`) — so the zero is a measurement, not a
silence (§9: *a census that returns ZERO must prove its instrument*).

**So the corpus-reading guards are clean at this reading.** iter-152's warning was right that anchoring is
not the property — and grading the property found the same answer for a better reason.

## Phase B — where the live instance actually was: the REPORTING layer

The class is real. It is one layer up from the corpus, in the mechanism that reports every guard's verdict.

`guard_family.py` ran each member with `capture_output`, **merged stdout and stderr** (`:257`), and took the
**last line of the merged stream** as a GREEN member's verdict (`:286`). On this tree `claim_census_guard`
emits a `DeprecationWarning` — `maxsplit` passed positionally, py3.13+ — and CPython echoes the offending
**source line** to stderr *after* the guard has finished speaking. The family therefore printed, as that
guard's verdict:

    GREEN  rc=0  claim_census_guard  [tree]    first = re.split(r"[\s,:;(]", s, 1)[0].strip("*_`.,:;").lower()

while the guard's own `claim-census: OK — ratchet holds over 41 files (1138 unevidenced assertions,
baseline 1164)` was **invisible in the one view that claims to summarise the family**.

**Census, with its denominator: 1 of the 17 runnable `[tree]` members** emits output it did not author.
(7 members are NOT-RUN for want of `--platform`/`--range`; they are named in the family's own tail and are
declared here rather than counted as passes.)

### `D-M257x-156-1` — the defect is in the RUNNER, and fixing the warning would have hidden that

The one-line fix to `claim_census_guard` removes today's instance and leaves the mechanism: any library
that writes to stderr re-creates it, on any member, silently. The runner cannot tell a guard's voice from
everything else on the wire, and that is the property to repair.

### `D-M257x-156-2` — the RED path was worse than the GREEN one

A warning echo is **indented**, which is exactly the shape `headline()` selects findings on. A member that
went RED while anything wrote to stderr could have had a line of Python reported as *its first finding* —
the same failure iter-87 wrote `headline()` to fix, re-entering through a door iter-87 could not see.

## Phase C — the repair

**Derived, never listed.** Guards already print their summary flush-left, prefixed with their own module
name and a colon — a convention `headline()`'s docstring has asserted since iter-87 and which **nothing
checked**. `speaks_for(line, name)` makes it checkable, on a folded token so all three live spellings pass:
`anchor-construct-guard:` (hyphens), `union_apply_guard:` (underscores), and `claim-census:` — a **prefix**
of its module name, which is why this is a prefix test and not equality. A "does it look like a summary"
heuristic was rejected: it would re-create the defect one remove up, since a warning's echo looks like
whatever the source line looks like.

- `verdict_line()` — three rungs (`own-summary` → `own-last` → `merged-last`), and **the rung is returned
  and printed** when it is not the first. iter-153's `D-M257x-153-5`: without it, *"the guard summarised
  itself"* and *"we could not find a summary, so here is the last thing on the wire"* print identically.
- `headline()` takes the guard's own stdout and prefers findings from it; **the merged stream stays the
  fallback**, so §5 rule 8 (a guard that fails without itemising is still reported) and harden pass-20's
  *"a guard that exits 1 while echoing a corpus line"* case are both untouched. The repair narrows the
  SOURCE of findings, never the fallback.
- **Noise is DECLARED, not dropped** — `⚠ NOISE <member>: N line(s) on stderr that the guard did not
  author`, plus a count in the family's summary line. Dropping it would be the same swallow in the other
  direction, and is how a warning went unseen for four releases *while being printed as a verdict*.
- **Noise does not turn a member RED.** A warning is not a finding, and grading it as one would be the
  runner inventing a verdict — the thing it exists to stop. Two contracts, `D-M255-1`'s precedent: the run
  **discloses**, the fence **gates**.

The `CANNOT RUN` / `Nothing was checked` sniff at `:277` still reads the merged stream. That is the same
class and it is **routed, not silently changed** — narrowing it needs its own evidence, and harden pass-20
wrote that line for a reason worth re-reading before touching it.

Source fix: `claim_census_guard.py:441` → `maxsplit=1` as a keyword. Behaviour identical; what changed is
that the source line stops being echoed. Census of the same construct across the monorepo: **3 remaining
occurrences, all three in prose** (this iter's own comment, docstring and test fixture) — **0 in live code**.

## Phase D — the fence, and what proves it

`stack-core/tests/test_guard_family_verdict_line_m257x.py`, **22 tests**, every class **above** the
`__main__` guard with collection parity asserted in-file rather than assumed (iter-153's lesson).

Arms: the three live name-spellings · a warning's **both** lines rejected (the location line *names the
guard* and carries a colon, so a "does it mention the guard" test would have accepted it) · a neighbouring
guard's summary rejected · **last**-own-summary-wins (anchoring alone is defeated by an earlier flush-left
self-named line — the exact defect iter-152 repaired in `blocking_state_guard`, which this must not
re-introduce) · every fallback rung **disclosed** · `headline` prefers a stdout finding and its inferred
count excludes the echo · the merged fallback **preserved** · the live census (0 noise, 0 fallback rungs,
the not-run members named) · and the specific regression pinned to the **structural** property — a new
ratchet baseline must not turn it red (§5 rule 71).

Controls: **anti-vacuity on the fixtures** (the echo must be indented and finding-shaped, the location line
must name the guard — otherwise every arm above passes while testing nothing); **anti-vacuity on the live
run** (≥10 members must have actually run, else the census is near-vacuous); one **mutation control**
reproducing the pre-fix selection over the same fixture and asserting it produces the defect.

⚠️ **`D-M257x-156-3` — the fence is interpreter-dependent, and saying so is half the deliverable.** The
family is launched with `sys.executable`. The motivating warning exists on **py3.13+ and not on py3.9** —
and **the only interpreter on this host carrying pytest is 3.9**, where the real defect is invisible: the
first full run was **18/18 green while the defect was live**. A fence that passes because of an unstated
precondition is iter-145's measurement-preconditions rule again. Two responses, both landed: the mechanism
is proven **interpreter-independently** by `TestRunOneWiring`, which drives `run_one` against a synthetic
member that writes to stderr on purpose; and the live arm **states the interpreter it measured with**.

**RED-PROOF against the real pre-fix tree, under py3.14: `Ran 22 tests … FAILED (failures=1)`** — exactly
the census arm, naming `claim_census_guard`. The other 21 are regression pins, **not** gap proofs, and are
booked as such. Post-fix: **22/22 OK under py3.14**, and **63 passed** under py3.9 together with the
pre-existing `test_guard_family.py`.

## Gates

- `test_guard_family_verdict_line_m257x` — **22/22** (py3.14) · **22 + 41 = 63 passed** with
  `test_guard_family.py` (py3.9)
- live `guard_family` — **17 GREEN · 0 RED · 0 could-not-check · 7 not-run · 0 NOISE · 0 fallback rungs**
- blast radius (10 test modules importing `guard_family` or `headline`) — see the close section
- **Not re-run, and saying so** (§5 rule 60): the full `stack-core` section, `stack-verify`, `dev-stack`,
  `demo-stack`, `stack-injection`. This iter touched `guard_family.py`, `claim_census_guard.py` and one new
  test file, all inside `stack-core`; the scoped run covers every module that imports either.

**No `N` reading taken, so no `N` movement is claimed.** Gate **4 of 5**, unchanged.

## Close — 2026-08-08

**Outcome:** the census over the corpus-reading guards came back **clean with its instrument proven**
(0 of 142 grading heads, 0 of 140 files ambiguous, 0 non-prose matches on the four hand-bound markers) —
and the live instance of the class was **one layer up, in the reporting layer**: `guard_family` printed a
Python `DeprecationWarning`'s source echo as `claim_census_guard`'s verdict. **1 of 17** runnable members.
Runner repaired (derived speaker test + disclosed rung + declared noise), source warning silenced, 22-test
fence RED-proofed 1-of-22 against the real pre-fix tree.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**no `N` reading taken, so the metric is
UNMEASURED not unmoved (§9); and a successor strategy remains FORBIDDEN by `TOK-08`'s sealed rule**) —
(3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted:
n — Outcome: continue
**Decisions:** `D-M257x-156-1` … `D-M257x-156-7` (see [`decisions.md`](decisions.md))
**Side-deliverables:** none — the `claim_census_guard` one-line fix is the instance the planned census
found, not an unrelated side discovery, and is graded as planned scope.
**Routes carried forward:**
- `FIX-M257x-iter156-cannot-run-sniff-reads-merged-stream` — **NEW.** `guard_family:277` sniffs
  `"CANNOT RUN"` / `"Nothing was checked"` over merged stdout+stderr. Same class; harden pass-20 wrote
  that line deliberately, so narrowing it needs its own evidence (`D-M257x-156-6`).
- `SURVEY-M257x-iter156-other-reporting-layers` — **NEW, and the generalisation.** The census was pointed
  at the layer that READS and the defect was in the layer that REPORTS. Other runners that summarise a
  subprocess — `autoverify`, the bring-up verify tails, `/test-platform`'s report generator — have never
  been graded on whether they can tell their subject's voice from the wire.
- `SURVEY-M257x-iter152-other-guards-may-read-prose-as-data` — **CLOSED by this iter**, on the property it
  asked for, with the denominators stated and the instrument proven (§9). The corpus-reading guards are
  clean at this reading; the class is real and its live instance is recorded above.
- Unchanged and still queued: `SURVEY-M257x-iter152-half-up-services-are-ungradeable` ·
  `SURVEY-M257x-iter150-partition-completeness-elsewhere` ·
  `FIX-M257x-iter155-add-injected-rows-vs-the-pinned-liveness-only-decision` (**deliberately not taken —
  `D-M257x-156-7`**) · `SURVEY-M257x-iter154-other-fences-may-be-pinned-to-spellings` ·
  `FIX-M257x-iter145-sha-baseline-drift` · `-iter145-migrate-race-needs-a-host-postgres` ·
  `-iter145-green-but-stale-graphql-mentions` · `-iter143-wrong-head-is-unfenced` ·
  `-iter143-scope-derivation-by-grep` · `-iter143-appending-to-the-protocol-doc-rots-the-ledger` ·
  `-iter144-correction-vs-retraction-unfenced` · `SURVEY-M257x-iter144-orphan-arm-is-the-residual` ·
  `FIX-M257x-iter142-path-arm-window` · `-iter142-value-change-articles` · `-iter142-tier-b-underflag` ·
  `FIX-M257x-iter135-adjudicated-live-defects` · `-iter140-receipts-not-checkable-here` ·
  `-iter140-receipt-fence` · `-iter138-anchor-rot-fence` ·
  `-iter134-fence-family-has-no-shared-predicate-layer` · `-iter133-two-fives-need-a-fence` ·
  `-iter131-predicate-sets-not-enumerated`
**Lessons:** `§5` gains **rule 72** with five corollaries — the RED path bites hardest; declare the noise
but do not grade it; a fence whose verdict depends on the interpreter must say which one; and **the
instrument is a claim too** (this iter's own enumerator produced 13 false `IMPORT-FAIL`s and a 105-of-171
raw signal, neither published as a defect count). The census discipline that mattered most: **when a
mechanical census over the obvious layer returns zero, ask which layer you pointed it at** — the answer
here was the one that reports.
