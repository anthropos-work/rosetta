# iter-158 — progress

**Type:** tik · **Strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07)

## Phase A — the routed repair, measured and REFUTED

`FIX-M257x-iter156-cannot-run-sniff-reads-merged-stream` proposed narrowing the rc-0
`CANNOT RUN` / `Nothing was checked` sniff from the merged stream to stdout, on the reasoning that
iter-156's verdict-line defect came from reading stderr as the guard's datum.

**The measurement, one command: run the family against an empty repo-root.**

| reading | number |
|---|---|
| members forced into a could-not-run state | **14** |
| of those, reporting off the `merged-last` rung — i.e. **stdout empty, message on stderr** | **14 of 14** |

**Narrowing the sniff to stdout would have turned *"nothing was checked"* into GREEN for all fourteen.**
`D-M257x-158-1`: the routed repair is refuted, and the routing was right — *"narrowing it needs its own
evidence"* is the sentence that stopped iter-156 from shipping the defect its own repair suggested.

**And the disclosure iter-156 shipped one iter earlier was already firing falsely on the same run:**
`⚠ NOISE demo_knob_guard: … MISSING: the defaults contract corpus/ops/demo/demo-up-defaults.md does not
exist. An absent contract is a FINDING, not a skip.` — **that is the guard's own sentence.** iter-156's
classifier called a stderr line foreign whenever it did not `speaks_for` the guard, and a guard's
continuation prose is not name-prefixed. `D-M257x-158-2`: **a disclosure that mislabels its subject's own
words is the defect it was built to stop, pointed the other way.**

## Phase B — one classifier answers both

`interpreter_noise(err_lines)` recognises the **interpreter's own signature**: a warning header
(`<file>:<line>: …Warning:`) or a traceback header, plus the **indented** line a warning header claims —
that echo being exactly the shape that was printed as a guard's verdict for four releases. Everything else
is the guard speaking.

**The default is deliberately INVERTED from `speaks_for`'s.** For the verdict line, an unrecognised line
must not be accepted as a verdict (unsafe direction: accepting an echo). For authorship, an unrecognised
line must not be taken away from the guard (unsafe direction: mislabelling its words). Same stream, two
questions, opposite safe defaults — and the arm that actually matters, the verdict line, is protected by
`speaks_for` either way.

The sniff now reads **merged minus foreign**: stderr stays in scope (all 14 need it), and an interpreter
echo containing the phrase can no longer launder a green member into `CANNOT-CHECK`.

## Phase C — the fence

`test_guard_family_verdict_line_m257x.py` grows to **28 tests** (from 22). New arms: a warning header and
its indented echo are both claimed · a traceback is claimed · **a guard's own unprefixed diagnostic is NOT
claimed** (the `demo_knob_guard` regression, by its real text) · a guard resuming *unindented* after a
warning is not swallowed · a `CANNOT RUN` inside an interpreter echo does **not** grade the member · a
`CANNOT RUN` the guard wrote **on stderr** still does. Controls: anti-vacuity on the fixtures (the echo
indented, the guard prose flush-left and warning-free) and a mutation control showing iter-156's rule
**would** have flagged the guard's prose — so the arm reads the new classifier and not something that
passed anyway.

⚠️ **Two fences went RED on correct changes, and both were re-pointed rather than deleted — the sixth and
seventh instances of `§5` rule 71 in five iters.**

- **`test_fence_provenance::test_a_MEMBER_is_still_suppressed_by_the_family`** asserted
  `assertNotIn("force", inspect.getsource(run_one))` — **a whole-file substring check**, which §8 names
  explicitly as the anti-pattern (*assert against a parsed construct, never a whole-file substring*). It
  fired on this iter's **comment** containing the word *"forced"*: correct prose reading as a contract
  breach. Re-pointed to `ast.unparse(ast.parse(src))`, which drops comments and docstrings and leaves the
  construct the contract is about, with an anti-vacuity assert that the unparse did not drop the subject.
- **iter-156's own `test_a_member_that_says_nothing_on_stdout…`** used a fixture reading
  `Traceback (most recent call)` — **a string CPython never emits.** iter-156's loose classifier accepted
  it (anything unprefixed was noise), so **the arm passed on a fixture that was not its subject**;
  iter-158's stricter classifier caught it within one iter. `D-M257x-158-4`: **an inexact fixture under a
  loose classifier is a test that proves nothing and reports that it did** — and tightening the classifier
  is what audits the fixture.

## Gates

- `test_guard_family_verdict_line_m257x` + `test_guard_family` — **71 passed · 0 failed**
- `test_fence_provenance` — **38 passed · 0 failed** (1 failed before the re-point — this iter's own,
  graded not bypassed)
- `test_repair_postcondition` + `test_fence_registry_completeness_m257x` — **76 passed** in the same run
- live `guard_family` — **17 GREEN · 0 RED · 0 could-not-check · 7 not-run · 0 NOISE**; and on the empty
  root the false NOISE row is **gone** while all 14 could-not-check verdicts are **unchanged**
- **Not re-run, and saying so** (§5 rule 60): the full `stack-core` section, `stack-verify`, `dev-stack`,
  `demo-stack`, `stack-injection`.

**No `N` reading taken, so no `N` movement is claimed.** Gate **4 of 5**, unchanged.

## Close — 2026-08-08

**Outcome:** the routed narrowing is **REFUTED by measurement** (14 of 14 members write their
could-not-run message to stderr; narrowing to stdout would have graded all fourteen GREEN), and the real
repair was an **authorship** filter — which also fixes the false-positive noise disclosure iter-156
shipped one iter earlier. One classifier, both defects, 6 new arms.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**no `N` reading taken, so the metric is
UNMEASURED not unmoved (§9); a successor strategy remains FORBIDDEN by `TOK-08`'s sealed rule**) —
(3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted:
**y** — Outcome: **exit-7**
**Decisions:** `D-M257x-158-1` … `D-M257x-158-4` (see [`decisions.md`](decisions.md))
**Side-deliverables:** two fence re-points (`test_fence_provenance`'s substring assert → parsed construct;
iter-156's inexact traceback fixture → the real shape), both recorded as rule-71 instances.
**Routes carried forward:**
- `FIX-M257x-iter156-cannot-run-sniff-reads-merged-stream` — **CLOSED by this iter, as a refutation plus a
  different repair.**
- `SURVEY-M257x-iter158-noise-classifier-is-narrow-by-choice` — **NEW.** `interpreter_noise` recognises
  warning and traceback signatures only. Noise from a subprocess a guard shells out to, or a C-level
  library writing raw text, is not recognised and is attributed to the guard. That is the deliberate safe
  default (`D-M257x-158-3`), and it is a stated blind spot rather than a measured zero.
- `SURVEY-M257x-iter156-other-reporting-layers` · `SURVEY-M257x-iter152-half-up-services-are-ungradeable` ·
  `FIX-M257x-iter155-add-injected-rows-vs-the-pinned-liveness-only-decision` ·
  `SURVEY-M257x-iter154-other-fences-may-be-pinned-to-spellings` (**now seven instances in five iters**) ·
  `FIX-M257x-iter145-sha-baseline-drift` · `-iter145-migrate-race-needs-a-host-postgres` ·
  `-iter145-green-but-stale-graphql-mentions` · `-iter143-wrong-head-is-unfenced` ·
  `-iter143-scope-derivation-by-grep` · `-iter143-appending-to-the-protocol-doc-rots-the-ledger` ·
  `-iter144-correction-vs-retraction-unfenced` · `SURVEY-M257x-iter144-orphan-arm-is-the-residual` ·
  `FIX-M257x-iter142-path-arm-window` · `-iter142-value-change-articles` · `-iter142-tier-b-underflag` ·
  `FIX-M257x-iter135-adjudicated-live-defects` · `-iter140-receipts-not-checkable-here` ·
  `-iter140-receipt-fence` · `-iter138-anchor-rot-fence` ·
  `-iter134-fence-family-has-no-shared-predicate-layer` · `-iter133-two-fives-need-a-fence` ·
  `-iter131-predicate-sets-not-enumerated`
**Lessons:** the same stream can carry two authorship questions with **opposite safe defaults** — never
accept an unrecognised line as a verdict; never take an unrecognised line away from its author. And a
routed item's proposed repair is a hypothesis, not a plan: this one was refuted by a single command that
nobody had run in two iters.
