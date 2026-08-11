---
milestone: M257x
iter: 11
---

# iter-11 — progress

**Type:** tik

## Phase 1 — re-survey (the inherited hand-off, re-measured)

`platform-alignment.md` §5 has said since iter-08 that *the search can succeed and the conclusion still be
false*. iter-11 is the third time in this milestone that an inherited pre-compute did not survive contact
with a measurement, and this time the artifacts that refuted it had been sitting on disk for five hours.

**The two verdict files.**

    stack-demo/rosetta-extensions/demo-stack/stacks/demo-1/autoverify.json
      {"project":"demo-1","offset":10000,"warnings":1,"green":false,"ts":"2026-07-31T15:49:36Z"}   <- the BRING-UP's own

    stack-demo/autoverify.json
      {"project":"demo-1","offset":10000,"warnings":2,"green":false,"ts":"2026-07-31T20:37:55Z"}   <- a STANDALONE re-run

Same tool. Same stack. Different vantage, different answer, five hours apart, and nobody compared the
timestamps. iter-10's hand-off — *"2 FAILED, and both are this"* — described the second file and attributed
it to the first.

**Point by point.**

1. *"The path is wrong."* **Refuted.** `up-injected.sh:2550` calls autoverify with `STACK_DIR="$STACK"`,
   and `:75` sets `STACK="$HERE/stacks/demo-$N"`. The bring-up path is correct **by construction**. The
   2-warning reading came from a standalone run with `STACK_DIR` = the workspace root, which holds neither
   log.
2. *"The fix must distinguish absent / empty / populated — today all three read as 'the phase never ran'."*
   **Refuted.** autoverify has distinguished all three since **v2.8 M256 harden pass 2**
   (`[ ! -e ]` → absent-warn, `elif [ -s ]` → populated-warn, `else` → `✓`), and four tests in
   `TestEvidenceAbsenceIsNotEvidenceOfSuccess` pin it. **Empty is the HEALTHY state by design** — M217:
   the log is written *only* from the applier's failure branches, and truncated per run.
3. *"The message asserts a CAUSE from the absence."* **Partly refuted.** The message already named the
   alternative: `"(or STACK_DIR is not the bring-up's \$STACK)"` — exactly what had happened. The
   pre-compute quoted it **truncated at the em-dash**. §5 rule 10, third occurrence.

**What survived, and it is bigger than the pre-compute described.**

`STACK_DIR` was a **hand-supplied path with no derivation and no validation**, in a script that already
derives its offset from `--project` and cross-checks it against the registry. Every failure mode of a
hand-supplied path had already happened:

| caller | what it passes | consequence |
|---|---|---|
| `up-injected.sh:2550` | the correct dir | fine — one caller, right |
| `dev-stack:298` | **nothing at all** | every `dev-N` bring-up silently skipped the cheap-win block, the transcript, **and `autoverify.json`** — the machine-readable verdict every grader reads. A skip that reads exactly like a pass (§5 rule 8) |
| a standalone/operator run | whatever the operator types | two warnings that are indistinguishable from real defects, and `green:false`. **Cost this milestone one whole iter.** |

And the trap was **documented rather than removed** — `CLAUDE.md`'s latency-budget row already carries
*"`autoverify.sh` needs `STACK_DIR`"* as a thing to remember.

## Phase 2 — the fix

**`target_resolve_stack_dir <project>`** (`stack-verify/lib/target.sh`) derives the receipts dir from the
project, the same move `target_resolve_offset()` already makes for the ports (§2, *derive it at the point of
use*):

    demo-N -> <rext>/demo-stack/stacks/demo-N     (up-injected.sh:75)
    dev-N  -> <rext>/dev-stack/stacks/dev-N
    anything else -> NOTHING

The last row is load-bearing: `anthropos` (the main dev stack) has no bring-up that writes receipts, and a
fabricated path would turn *"we know nothing"* into a claim.

In `autoverify.sh` the derivation **wins** over an inherited `$STACK_DIR` whenever the project has a known
layout, and a mismatch is **named** rather than swallowed. Both sides are normalized through `cd && pwd`
first, so an equivalently-spelled path (trailing slash, symlinked root) cannot manufacture a mismatch on the
bring-up path — clause 1 needs that call to stay silent.

**`target_is_demo_project`** then gates the two receipt asserts on the project's **TYPE**, not on whether a
caller happened to set a variable. A dev stack has no patch phase and no UI tier, so those receipts are
correctly absent there; deriving the dir without the type gate would have manufactured two false warnings on
every dev bring-up. The transcript and the verdict json stay stack-type-neutral — so **dev gains both, which
it never had.**

### The first cut cried wolf, and only the wider suite said so

The first version **warned** when the derived dir did not exist. Ten targeted tests passed. The full
`stack-verify` suite then ran it into **18 pre-existing fixtures** — every autoverify test uses a synthetic
`--project demo-1` with no stack dir on disk. *A fence that cries wolf gets disabled, and a disabled fence is
indistinguishable from never having written one* (§8 rule 6).

Reasoned down rather than deleted. up-injected.sh `mkdir -p`s the stack dir at `:226` and truncates both logs
in the next eleven lines, so **a missing dir does not describe a real bring-up** — it describes running the
script from a clone that did not perform one. It is now a **named `·` skip** that prints the path it
consulted (which clone was asked is the entire diagnostic, §5 rule 12). Every case that *does* describe a
real bring-up — dir present, receipts absent or populated — stays a full warning. Nothing clause 1 depends
on was downgraded.

## Phase 3 — re-measure

**Offline.** `stack-verify` **224/224 green** (214 before + 10 new). 6 mutants, each parsing, each collecting
exactly 1 test, each **RED**, plus an **unmutated control that goes GREEN** — so the battery is measuring the
fence and not the harness:

| mutant | test that fired |
|---|---|
| M1 derivation does not win (`${STACK_DIR:-$derived}`) | wrong-supplied-dir |
| M2 no type gate | dev-stack receipts/verdict |
| M3 missing-dir hole unnamed | named-hole |
| M4 mismatch not named | wrong-supplied-dir |
| M5 verdict json gated on demo | dev-stack receipts/verdict |
| M6 *make the helper total* (**two-point**) | no-fabrication |
| C0 control: unmutated | — (**GREEN**, as required) |

**LIVE, on the running `demo-1`, from the stack's own pinned consumption clone** — the decisive evidence,
because iter-10's own first cut passed three unit tests and was caught only by a live negative control:

| vantage | before (iter-10) | after |
|---|---|---|
| standalone, **no** `STACK_DIR` | *(block skipped entirely)* | `✓ demo-patches` + `✓ frontend builds` → **`warnings:0 / green:true`** |
| standalone, **wrong** `STACK_DIR` (the exact iter-10 vantage) | `warnings:2 / green:false`, both "NO … evidence" | receipts still found (`✓✓`); the **only** warning is the named mismatch — about the caller, not a fabricated claim about the stack |
| **correct** `STACK_DIR` (the bring-up's own call) | — | **silent**, `warnings:0 / green:true` |

`demo-1`'s verdict file now reads `{"warnings":0,"green":true}` from a run that was given no path at all.

## Phase 4 — two harness lessons, promoted to the protocol doc

Both landed in `corpus/ops/platform-alignment.md` in this iter's commit:

- **§8 rule 5 addendum — read the COUNT with the exit code.** This iter's own mutation battery reported a
  clean `RED` for a mutant nothing had tested: it had invoked a `python3` without pytest, so every run
  returned non-zero and every mutant "fired". The same shape occurs with no tooling accident at all, because
  **pytest exits 5 when a `-k` filter matches nothing** — a renamed test makes its own battery report RED
  forever. Gate on `collected == 1` *and* an exit code that means failure rather than emptiness.
- **§8 rule 5 addendum — a mutant that changes nothing is not a survivor.** The battery's first M6 reported
  `GREEN (mutation SURVIVED)`; it had added an unreachable `case` arm below an early `return`. The invariant
  was enforced **twice** (early return AND closed case) — §8 rule 4 working. The honest response is the
  two-point mutant a future editor would actually write (*"make the helper total"*), which parses, collects,
  and goes red.
- **§5 rule 12 (new) — say which INVOCATION produced the number.** The whole iter-10 misreading in one rule,
  with its remedy: record the vantage and the artifact's timestamp with every measurement, then *remove the
  parameter* — a correct diagnostic that has to be read carefully is a weaker control than a parameter that
  no longer exists.

## Close — 2026-08-01

**Outcome:** the last standing autoverify warning class is closed at its real mechanism. `demo-1` standalone
autoverify **2 warnings / `green:false` → 0 warnings / `green:true`**, proven live from the pinned
consumption clone; `dev-N` bring-ups gain the transcript + verdict they never wrote. The inherited
pre-compute was refuted on 3 of 5 points before code was written.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — clause 4 remains claimable, clause 1's blocker is cleared but its three cold cycles are
unrun; clauses 2/3/5 untouched.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D-M257x-12, D-M257x-13, D-M257x-14 (see `decisions.md`)
**Side-deliverables:** none — the protocol-doc rules are part of this iter's own lesson, committed with it.
**Routes carried forward:**

| item | target | why |
|---|---|---|
| `FIX-M257x-vmram-gib-unit` | iter-12, **before** the cold cycles | it fires on every bring-up iter-12 is about to run, and it is a doc/code unit mismatch — this milestone's own subject matter. Note the cost it carries: editing `up-injected.sh` shifts the `demo-up-defaults.md` `file:line` citations (iter-05 hit exactly this; the corpus guard's `--fix` repairs it). |
| `FIX-M257x-dev-verify-svcs-stale` **NEW** | later tik | `dev-stack:298` still scopes the dev verify to `skiller skillpath jobsimulation cms roadrunner` — names the platform no longer ships. A hand-maintained list in the same file this iter touched, of exactly the class the milestone exists to end. Found in passing; **not measured** — whether it produces false `down`s or is absorbed by `target_warn_unknown_services` is unknown, and saying so is the point. |
| `CHECK-M257x-demopatch-pristine` | still open | unchanged by this iter; the `pristine-ing skipped/failed` collapse is the same *name the state you measured* shape. |

**Lessons:**

- **A verifier that takes "where to look" as a parameter can be aimed at a place the thing under test never
  wrote to — and it will then report absence of evidence in the tool's own voice.** That is
  indistinguishable from a real defect at a glance, and it is how a correct verifier produced a false
  hand-off. Derive the vantage; do not accept it.
- **The wider suite is where the finding is, for the sixth consecutive iter.** Ten targeted tests were green
  on a change that broke 18 existing ones. The targeted tests were not wrong — they were about the new
  behaviour, and the cost of the new behaviour was somewhere else entirely.
- **Reasoning a severity down is not weakening a check**, provided the reasoning is written and the cases
  that carry the real signal keep their severity. The alternative — a warning nobody can act on, on every
  synthetic run — ends as a disabled check.
