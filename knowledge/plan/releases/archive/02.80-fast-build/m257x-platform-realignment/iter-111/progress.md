# iter-111 — TOK-07 step 0: the two hardening items with teeth

**Type:** tik · `iter_shape: tooling` (§9's tooling-iter: ship the instrument AND use it in the same iter).
**Active strategy:** `TOK-07`, step 0.

## Line 1 — `FIX-M257x-harden23-json-polluted-by-provenance-stamp`: DECIDED and closed

The routed question was *where does provenance belong*, and it was posed as a dilemma between iter-105's
*printed-FIRST-on-stdout* property and a parseable `--json`. **It is a false dilemma** — see
`D-M257x-111-1`. iter-105's two stated reasons are about **order and shape**; neither is about the
**stream**.

**Decision.** The payload's grammar decides, and the mode is derived from the invocation:

| mode | human line | where the tree is stated |
|---|---|---|
| text | **stdout, FIRST** — unchanged, byte for byte | the printed line |
| `--json` | **stderr** | **inside the document**, under `fence_tree` |

This is the doctrine *strengthened*: an archived `verdict.json` now states its own tree without the
terminal that produced it — which neither a stdout preamble captured by `>` nor a stderr line ever did.

**Shipped:**
- `fence_provenance.py` — `wants_machine_output()` (argv-derived, matches `--json` and `--json=…` so a
  later valued-flag change cannot silently re-arm it), `stamp_main()`, `json_payload()`, `emit_json()`.
  `stamp()` itself is untouched; `guard_family` still forces its own line. Module docstring amended so
  the *"prints on stdout"* claim carries its text-mode scope.
- **19 guards** moved from `stamp()` to `stamp_main()` in `__main__`, uniformly, with the iter-111 reason
  in each comment (the mechanical rewrite reported **19 rewritten, 0 misses** — no file hand-edited).
- **11 JSON emission sites** moved to `emit_json`.
- **Module-scope `import fence_provenance`** in the 11 affected guards, with the reason inline: a
  `__main__`-only import is green under `python guard.py` and `NameError` under `from guard import main`
  — *green exactly where nobody looks*.

**The workaround is gone, not just the defect** (`D-M257x-111-2`). All five call sites that set
`FENCE_PROVENANCE_STAMPED=1` now POP it via a `_clean_env()` helper — including the fifth the harden pass
did not name, written as `G.__name__ and "FENCE_PROVENANCE_STAMPED"`. Fenced by
`TestTheWorkaroundIsGone`, whose anti-vacuity control writes a synthetic setter, confirms the matcher
still sees one, **and confirms it does NOT fire on the removal helper** — an over-broad scan is a scan
that gets suppressed rather than fixed.

**Net-new, found by deriving rather than by looking** (`D-M257x-111-3`): **`anchor_construct_guard`
declared `--json` and read it nowhere** — the flag parsed, exited 0, and printed prose. The
`demo_knob_guard` false-promise class with its halves swapped. Implemented against the same values the
text report prints, and fenced by an AST walk over every `--json`-declaring guard.

**Controls, per `TOK-06`'s binding clause, and both demonstrably fire.** `TestMachineModeMutationControls`
loads *mutated copies* of `fence_provenance.py` and shows the assertions going red against them:
dropping the injection leaves a document with no tree; forcing the stamp back onto stdout makes the
document raise `JSONDecodeError`. Plus the anti-vacuity control on the discovery itself — if the AST walk
stops finding `--json` parsers, the loop would iterate an empty list and pass while measuring nothing, so
it asserts **≥ 10 found** and **≥ 5 documents actually produced**.

`tests/test_fence_provenance.py`: **34 passed**, `python3 -m pytest tests/test_fence_provenance.py -q
-p no:cacheprovider --no-header`, 81.82 s.

## Line 2 — `FIX-M257x-iter108-stackcore-suite-hangs`: REFUTED as stated, and the better half kept

**The suite does not hang.** (`D-M257x-111-4`.) It completes, measured three times.

| | invocation | wall | result |
|---|---|---|---|
| A | `pytest tests/ -q --tb=no -p no:cacheprovider --no-header` (tree edited mid-run — a confound, recorded) | 414.14 s | 17 failed · 982 passed · 4 skipped |
| B | `pytest tests/ -q --tb=line -p no:cacheprovider --no-header --durations=8` | 431.36 s | 17 failed · 991 passed · 4 skipped |
| **C** | `pytest tests/ -q --tb=line -p no:cacheprovider --no-header --durations=5`, after this iter's repairs | **1090.88 s** | **1 failed · 1011 passed** |

**The freeze is real and fully accounted for.** `test_the_dev_fences_are_red_proven` is 132–136 s by
pytest's own `--durations`; its module alone measures **142.38 s**; it prints nothing while running (one
`-q` dot) and inside it a harness runs a nested suite 8× at ~16 s each through `subprocess.run`. Low
parent CPU is therefore **expected, not diagnostic**. Watched at 45-second intervals in run B, output sat
at **522 bytes for 2 m 15 s** and then advanced. That interval *is* the reported hang.

**Retracted:** *"blocks indefinitely"*, *"blocked, not slow"*, and *"the standing total cannot be produced
by a plain full-suite run on this host."*
**Kept, and it was the better half of the routing all along: state the invocation with the count.**

**And run C carries its own finding: the suite got 2.5× SLOWER by being fixed** — 431 s → 1090 s — because
a battery that dies on its baseline never runs its mutants. The mechanical-fences battery alone went from
aborting to **428.52 s** of real mutant execution. *A fast suite is not evidence of a healthy one.*

## The 17 failures were 16 mine and 1 documented — said plainly

`D-M257x-111-5`. Runs A and B came back **17 failed**. The flattering reading — *"a hidden RED, exposed at
last"* — is false. **16 were caused by this iter**: five mutation batteries stage a hand-listed **subset**
of `stack-core`, this iter gave eleven guards a module-scope `import fence_provenance`, and a staged tree
without that file cannot import the guard — so the battery reports a **RED BASELINE**, i.e. *"the fence is
broken"*, for *"you forgot a file."* One battery had also pinned the exact literal
`print(json.dumps({"leaks": …}, indent=2))` as a mutation target; this iter moved that line, so the mutant
silently became a **no-op**.

Fixed inside the iter: `fence_provenance.py` declared in all five `_COPY_FILES` with the reason inline;
the moved literal re-pinned. Five batteries + the answer key: **30 passed · 1 failed** in 664.44 s,
invocation `pytest tests/test_m257x_claim_twin_mutation_battery.py …(5 modules)… tests/test_claim_twin_guard_iter48_answer_key.py -q --tb=line -p no:cacheprovider --no-header`.

**The 1 is `test_claim_twin_guard_iter48_answer_key::test_02` — now genuinely RE-ATTESTED**
(`D-M257x-111-6`). Carried for two sessions as *"unchanged but not re-run"*; it has now actually been run
**four times** this iter and fails exactly as documented, same signature.

**Routed, not fixed** (`FIX-M257x-iter111-staged-battery-dependency-is-underived`): the batteries assert
the *presence* direction — *"the list names a file that does not exist"* — and have no assertion in the
*absence* direction, which is the one that fires when a guard grows an import. Deriving a staged tree's
closure is an import graph, not a grep, and this iter had two planned lines already.

## The harden cap, recorded

`D-M257x-111-7`. The 3-pass incremental cap fired **without stabilization for the second time** (pass 22,
now pass 25). Recorded as a standing signal and deliberately **not** converted into a request for a fourth
pass: each pass this session read a *different dimension*, so a non-empty third pass means the dimensions
are not exhausted, not that three passes are too few. What this iter adds: one of pass 25's two routed
items was **live, not theoretical**, and closing it surfaced two findings no harden pass had reached.
**The cap is not costing coverage; it is deferring it.**

## Protocol updated in the same commit (§8)

Three sections, each generalising past this iter: *a verdict with a grammar states its provenance inside
the payload* (incl. the **both-directions** flag rule) · *a battery that stages a subset carries an
underived dependency contract* · *a silent test is not a blocked test*.

## Close — 2026-08-06

**Outcome:** `TOK-07` step 0 landed. `--json` is parseable and **self-describing** on every guard that
offers it, with the undocumented env var removed at all five sites and fenced; one guard's `--json` was a
**declared-but-never-read** flag and is now implemented and fenced; and the routed *"suite hangs"* is
**refuted by measurement** — it completes, `1 failed · 1011 passed` in 1090.88 s with the invocation
stated, the freeze being a 2¼-minute silent test whose duration is now measured rather than inferred.
**16 of the 17 REDs the runs surfaced were this iter's own** and are recorded as such.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** D-M257x-111-1 … D-M257x-111-7
**Side-deliverables:** `anchor_construct_guard --json` implemented (`D-M257x-111-3`) — a declared flag the
code did not keep; found while deriving the JSON surface, in-scope, separately recorded.
**Routes carried forward:**
- `FIX-M257x-iter111-staged-battery-dependency-is-underived` → a future iter; derive the staged closure or
  make the staged run's failure name its cause
- `FIX-M257x-iter111-buildbench-parse-json-is-a-noop-flag` → a future iter; `buildbench parse --json` is
  read nowhere (the subcommand always emits JSON), so it is a no-op rather than a false promise
- `TOK-07` step 1 — the per-predicate corpus-wide enumerator → **next iter**
- `FIX-M257x-iter107-drift-fence-satisfiable-by-prose` → stays open, de-ranked by `TOK-07`
- `DEF-M257x-iter101-briefing-rext-tree` → stays open
**Lessons:**
- **A dilemma between two properties usually means one of them was described in terms of its
  implementation.** *"Prints first on stdout"* bundled an ordering rule, a shape rule and a stream choice;
  only the first two had reasons, and separating them dissolved the trade-off entirely.
- **Fix the mechanism, not just the instance.** The finding was never *"JSON is polluted"* — it was *"the
  suite is green because of an undocumented variable."* Repairing the first and leaving the second would
  have kept the hiding mechanism in service for the next defect.
- **Before calling a stall a block, measure how long the thing is supposed to take.** A long test that
  captures its children's output emits nothing and burns no parent CPU: the exact signature of a hang,
  produced by working normally.
- **A fast suite is not evidence of a healthy one** — repairing five batteries made the suite 2.5×
  slower, because a battery that dies on its baseline never runs its mutants.
