# iter-111 — decisions

## `D-M257x-111-1` — where provenance belongs when the payload has a grammar

**The routed question** (`FIX-M257x-harden23-json-polluted-by-provenance-stamp`) was framed as a
dilemma: keep iter-105's *printed-FIRST-on-stdout* property and leave `--json` unparseable, or default
`stamp()` to `sys.stderr` and retire a property documented in `fence_provenance.py`'s docstring plus 16
per-guard `__main__` comments.

**It is a false dilemma, and re-reading iter-105's own rationale is what shows it.** Both stated reasons
are about **order and shape**, not about the **stream**:

1. *"`guard_family.run_one` reports `lines[-1]` for a green member, so a trailing stamp would replace the
   guard's own summary"* — an ordering argument.
2. *"`headline()` counts finding-SHAPED lines; the stamp is flush-left so it cannot inflate a RED
   cardinality"* — a shape argument.

Neither requires stdout. And a flat move to stderr **would** have cost something real: a verdict archived
with `guard.py > verdict.txt` would no longer contain its tree, which is iter-103's condition again.

**Decision — the mode decides, and the mode is derived from the invocation:**

| mode | human line | where the tree is stated |
|---|---|---|
| **text** (no `--json`) | **stdout, FIRST** — unchanged, byte for byte | the printed line |
| **machine** (`--json`) | **stderr**, so a terminal still shows it | **inside the document**, `fence_tree` |

This is not a retreat from iter-105's doctrine. It is the doctrine applied to a payload that has a
grammar: **a preamble line does not make a JSON verdict weaker, it makes it not a verdict** — the
consumer dies at char 0 and gets nothing. And an archived `verdict.json` now states its own tree
**without the terminal that produced it**, which neither a stdout preamble captured by `>` nor a stderr
line ever did. The doctrine got stronger, not weaker.

**And the mode is read from `sys.argv`, never from the environment.** `wants_machine_output()` matches
`--json` and `--json=…` (so a later change to a valued flag cannot silently re-arm the defect). No new
knob was added, and specifically not an environment variable — the environment is what was holding the
suite green.

**Shipped:** `fence_provenance.wants_machine_output` / `stamp_main` / `json_payload` / `emit_json`; all
**19** guards' `__main__` moved from `stamp()` to `stamp_main()`; **11** JSON emission sites moved to
`emit_json`. `stamp()` itself is untouched and `guard_family` still forces its own line.

## `D-M257x-111-2` — the fix was never *"JSON is polluted"*; it was *"the suite's green had a hidden condition"*

Harden pass-23's finding had two halves and the second is the one that matters: **every test that parsed
a guard's `--json` set `FENCE_PROVENANCE_STAMPED=1` first.** Repairing the pollution and leaving the
workaround would have left the *mechanism* in place — an undocumented, load-bearing env var that makes
tests pass — ready to hide the next defect.

So the workaround is **removed at all five call sites** (four the harden pass named, plus one it missed
in `test_clone_drift_guard.py:141`, written as `G.__name__ and "FENCE_PROVENANCE_STAMPED"`), replaced by
a `_clean_env()` helper that **POPS** the variable — so these tests now exercise the path an operator
actually gets.

And it is **fenced going forward**: `TestTheWorkaroundIsGone` scans every `test_*.py` on disk for a
*setter* (the variable as a dict key or subscript with a value assigned), with the `fence_provenance`
suite itself the single allowed exception because suppression semantics are its subject. Its
anti-vacuity control writes a synthetic setter and asserts the matcher still recognises one — **and
asserts the matcher does NOT fire on the removal helper**, because an over-broad scan is a scan that
gets suppressed rather than fixed.

## `D-M257x-111-3` — a DECLARED flag that is read nowhere: the false promise, inverted

Found while deriving the `--json` surface rather than by looking for it. **`anchor_construct_guard.py`
declared `--json` and read `args.json` nowhere.** `--json` parsed, exited 0, and printed the human
report — so a consumer that trusted the flag got prose where it asked for a document, with no error
anywhere.

`demo_knob_guard`'s rule is that *a doc-promised flag with no parser entry is a **false promise***. This
is the same defect with the halves swapped — **a parser-promised flag the code does not keep** — and
nothing in the suite was looking for it, because no test ever asked this guard for JSON.

**Implemented rather than removed**, against the same values the text report prints so the two
renderings cannot disagree. And **fenced**: `test_a_declared_json_flag_is_actually_READ_by_its_guard`
walks every `--json`-declaring guard's AST and asserts the flag is read.

*Routed, not taken:* `buildbench.py` declares `--json` on four subcommands and reads it on three — the
`parse` subcommand emits JSON unconditionally, so its `--json` is a **no-op rather than a false
promise** (the output is a document either way). Milder, out of this iter's planned scope, and
`buildbench` does not stamp, so it carries none of the pollution.
→ `FIX-M257x-iter111-buildbench-parse-json-is-a-noop-flag`.

## `D-M257x-111-4` — ⚠ THE SUITE DOES NOT HANG. The observation was real; the inference was wrong.

`FIX-M257x-iter108-stackcore-suite-hangs` recorded that `pytest tests/` in `stack-core` **"BLOCKS
INDEFINITELY"** at `test_m220_mutation_battery.py::DevWiringMutationBattery::test_the_dev_fences_are_red_proven`
— *"blocked, not slow"* — on the evidence *"12.6 s CPU over 3 m 43 s, frozen at 442 results, reproduced
in two runs."* The corollary drawn from it was that the standing *"975 pass / 1 fail"* figure **cannot be
produced by a plain full-suite run on this host**.

**Measured this iter, three times, and it is not so.** The suite completes:

| run | invocation | wall | result |
|---|---|---|---|
| A (tree edited mid-run — a confound, recorded) | `python3 -m pytest tests/ -q --tb=no -p no:cacheprovider --no-header` | **414.14 s** | 17 failed / 982 passed / 4 skipped |
| B (clean tree, edits settled) | as above `--tb=line --durations=8` | **431.36 s** | 17 failed / 991 passed / 4 skipped |
| C (after this iter's repairs) | see `progress.md` — stated with its count | ~**7 min** | see `progress.md` |

**The mechanism is fully accounted for and it is duration, not blocking:**

- `test_the_dev_fences_are_red_proven` is the suite's slowest test at **132.72 s** (pytest `--durations`,
  run B) — and the module alone takes **142.38 s** end to end, measured separately.
- It emits **no output at all** while running: it is one `pytest -q` dot, and inside it the harness runs
  a nested suite 8 times (baseline + 7 mutants) at ~16 s each, `subprocess.run(..., timeout=900)`.
- Low parent CPU is therefore *expected*, not diagnostic — the work is in child processes and in waiting.
- Watched directly at 45-second intervals in run B: output was **frozen at 522 bytes from ~20:51:48 to
  ~20:54:03 — 2 m 15 s — and then advanced.** That interval **is** the reported "hang".

> **A silent test is not a blocked test, and CPU-idle is not evidence of a block when the work is in a
> child.** iter-108 observed a 3 m 43 s freeze and stopped there; the missing measurement was *"how long
> is this test supposed to take?"*, and it is ~2¼ minutes.

**What is RETRACTED:** *"blocks indefinitely"*, *"blocked, not slow"*, and *"the 975/1 figure cannot be
produced by a plain full-suite run on this host."* A plain full-suite run does produce a total.

**What SURVIVES, and it was the better half of the routing all along:** ***state the invocation with the
count.*** That rule is correct whether or not the suite completes — a scoped pass must never read as a
whole-suite pass — and it is kept, in §5 and in every count this iter quotes. The routed item is closed
by refutation-plus-retention, not by a fix.

## `D-M257x-111-5` — ⚠ 16 of the 17 failures were MINE, made and closed inside this iter — and the cause is worth more than the fix

The full-suite runs came back **17 failed**. One is the documented pre-existing
`test_claim_twin_guard_iter48_answer_key::test_02`. **The other 16 were caused by this iter**, and
saying so plainly is the point: the temptation was to book them as *"a hidden RED, exposed at last"*,
which is the flattering reading and is false.

**Mechanism.** Five mutation batteries stage a **subset** of `stack-core` into a temp tree and run the
fence's suite there. This iter added `import fence_provenance` at **module scope** to eleven guards — so
a staged tree without `fence_provenance.py` **cannot import the guard at all**, and the battery reports
a **RED BASELINE**, which reads as *"the fence is broken"* rather than *"you forgot a file."* One battery
also pinned the exact literal `print(json.dumps({"leaks": …}, indent=2))` as a mutation target; this iter
moved that line onto `emit_json`, so the mutant became a **no-op** — the very shape the harness's own
`MUTANT IS A NO-OP` guard exists to catch, arriving in a file the guard does not cover.

**Fixed:** `fence_provenance.py` declared in all five `_COPY_FILES` lists with the reason recorded
inline; the moved literal re-pinned. All five batteries green — **30 passed, 1 failed** (the documented
answer-key failure), 664 s, invocation stated in `progress.md`.

**The generalisation, now in §8:** ***a battery that stages a SUBSET carries a dependency contract, and
nothing derives it.*** The batteries do assert *"the dependency list names a file that does not exist"* —
the **presence** direction — and have no assertion in the **absence** direction, which is the one that
fires when a guard grows an import. Routed as
`FIX-M257x-iter111-staged-battery-dependency-is-underived` rather than fixed here: deriving a staged
tree's closure is a real piece of work (an import graph, not a grep), and this iter had two planned
lines already.

## `D-M257x-111-6` — `test_claim_twin_guard_iter48_answer_key::test_02` is RE-ATTESTED, and it still fails

Carried for two sessions as *"unchanged but not re-attested — recorded as not-re-run rather than as a
pass."* This iter **actually ran it**, four times (runs A, B, the battery re-run, run C). It **FAILS**,
exactly as documented, with the same signature (`corpus/04.md:1 <- …`). Recorded as *re-attested,
failing* — which is what it was always believed to be, now measured rather than assumed.

## `D-M257x-111-7` — the 3-pass harden cap fired without stabilization for the SECOND time. Recorded, not papered over.

Harden pass 22 hit the incremental 3-pass cap with the dimension scan still surfacing items; **pass 25
did it again** — the third pass surfaced two more (the answer key's own rationale quoting a total under
one part's name, and the `--json` pollution). Two occurrences is a pattern, not an accident.

**It is recorded here as a standing signal and deliberately NOT converted into a request for more
passes.** The honest reading is the one pass 25 gave: what a non-empty third pass means is that **the
scan's dimensions are not exhausted**, not that three passes are too few — each pass this session read a
*different* dimension (code · code again · what the tooling says about itself), and a fourth dimension
would find a fourth class. The cap is doing its job: it bounds a scan that has no natural end.

**What this iter adds to that record:** one of pass 25's two routed items (`--json`) turned out to be
**live, not theoretical** — and closing it surfaced `D-M257x-111-3` and `D-M257x-111-5`, neither of which
any harden pass had reached. *The cap is not costing coverage; it is deferring it.* Kept visible in the
hardening ledger and in this iter's close so the next `harden-mstone-iters` invocation inherits the
signal rather than re-deriving it.
