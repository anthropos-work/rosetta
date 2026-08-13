**Type:** tik — under `TOK-05`, step 1 (**fence**) applied to the predicate iters 68–70 measured.
Carries `FENCE-M257x-iter68-citation-resolution`, the largest fence route still open.

# iter-71 — one knob, and the corpus does not have one ref

## Phase A — the class, measured before the build

| | n |
|---|---|
| resolvable citations | **125** |
| block names **exactly one** ref that resolves in that citation's own clone | **31** |
| block names **more than one** resolvable ref | **12** |
| block names none that resolve | 82 |

**A quarter of the class was being graded against a file it never claimed.** iter-68 gave three
guards a ref and gave each of them *one* — a process-wide `CITE_REF`. `backend.md:39` pins to `app`
`b948604` v1.366.0; iter-69 re-pointed `shared_libraries.md:79` to `9d00a313` v1.367.0. **No single
knob can be right about both**, and iter-68 measured the consequence in its sharpest form: the same
corpus is **GREEN at origin HEAD and 4-findings RED at the pinned build ref.**

## Phase B — `block_ref()`, and what it refuses to do

| block names | behaviour |
|---|---|
| exactly one sha that `rev-parse`s **in that citation's own clone** | read at it — `block-pinned` |
| **more than one** | fall back **and count it** — `ambiguous` |
| none that resolve | the `CITE_REF` ladder, unchanged — `default` |

**The ambiguous case is a refusal, not an oversight.** A block legitimately names two refs when it
*contrasts* them — `platform-alignment.md` rule 32, as iter-69 rewrote it, names the pin **and** the
ref that deleted the line. Guessing which governs would be a rule fitted to a sentence (§4 Trap A);
falling back in silence would hide 12 citations inside `default` and let the fence over-report its
own reach, which is the failure this milestone has now found eight times.

Live, and the reach is now printed beside the provenance:

```
adjudicated at origin/main@9d00a31 x31, worktree(fallback) x30, b948604@b948604 x11,
               915da06@915da06 x6, 5ba17044@5ba1704 x4, 9d00a313@9d00a31 x3, …
ref chosen by  default x57, block-pinned x31, no-clone x30, ambiguous x12
```

`CITE_REF=worktree` **still overrides every block pin** — verified live: 1 finding under `worktree`,
GREEN by default. A per-block pin that defeated the escape hatch would remove the only way to ask
*what does the checkout say*, which is the mode iter-68 used to demonstrate the defect at all.

A pin at which the cited **file does not exist** is **UNMEASURED** — not clean, not a finding (§5
rule 7).

## Phase C — RED before trusted, and one mutant SURVIVED

| mutant | caught |
|---|---|
| M1 ignore the block entirely | 6F |
| M2 guess on ambiguity (take the first of two) | 4F |
| M3 **window narrowed block → line** | **SURVIVED — then 1F** |
| M4 accept a sha without verifying it resolves | 4F |
| M5 block pin beats `CITE_REF=worktree` | 1F |
| M6 an unresolvable pin silently re-reads at `auto` | 2F |
| M7 window widened block → whole document | 2F |
| **no-op control** | **SURVIVED (55 OK)** — as it must |

**M3 is the finding.** Every fixture I had written put the pin and the citation on the same line, so
narrowing the window to `lines[i]` passed the whole suite — **the test agreed with the
implementation instead of with the corpus.** `backend.md:39` carries its pin mid-sentence across a
wrap; rule 33 exists *because* pins and the claims they qualify are separated in prose. **Twice
before in this milestone a one-line window has been the bug** (iter-63's retired-token
discriminator; iter-68's negation window, recorded as *"the second window bug of this milestone
wearing a policy's name"*). This is the third — and the first time it was in the **test**.

Two corpus-shaped fixtures added: a pin **two lines above** its citation in one block (must be
found), and a pin **one blank line away** in a different block (must **not** be — rule 33: a pin
exempts a claim, never a neighbourhood). Both window mutants now die.

## Phase D — a broken caller the baseline caught

Adding one reach counter grew `run()`'s return from a 6-tuple to a **7-tuple** and broke **four
existing test call sites**. It was visible only because the battery's baseline came back `4 errors`
instead of `OK` — which is the whole reason a battery is run against a **stated** baseline rather
than a remembered one. Fixed at the call sites; the fragility is **recorded, not refactored**
(`D-M257x-71-3`), because converting to a dataclass is a third line of work in a two-line iter.

## Phase E — gates

| gate | result |
|---|---|
| five corpus guards | **all OK** — and `anchor_construct_guard` now prints `ref chosen by …` beside `adjudicated at …` |
| `CITE_REF=worktree` | **still RED** (1 finding) — the escape hatch survives the change |
| `tests/test_iter45_mechanical_fences.py` | **55** (was 46); new class 9/9 |
| mutation battery | **7 mutants, all caught** (one after it SURVIVED and the test was rewritten against the corpus); **no-op control survived** |
| `stack-core` suite | **762 tests, 1F** — `test_claim_twin_guard_iter48_answer_key::test_02…`, the perishable iter-48 fixture. **Baseline matched by IDENTITY**, not count (+9 from this iter). 938 s against iter-69's 528 s: the m220 flake gate logged **7 retry attempts**, which is the gate working, not a regression |
| `stack-injection` · `dev-stack` · `demo-stack` | untouched sections; iter-69's runs stand (332 OK · 151 OK solo · 1048/7F by identity) |

## Close — 2026-08-04

**Outcome:** `FENCE-M257x-iter68-citation-resolution` lands. iter-68 gave three guards a ref and gave
each of them **one**; the corpus does not have one ref — **31 of 125 resolvable citations sit in a
block naming their own**, and every one was being read at `origin/main` regardless. Each citation is
now graded at the ref its own block names, the **12 ambiguous blocks are counted rather than hidden
inside the default**, a sha that does not resolve in that citation's own clone is not a pin, an
unresolvable pin is **UNMEASURED not clean**, and `CITE_REF=worktree` still overrides everything.
**A mutant survived**: narrowing the window from the block to the line passed the entire suite,
because every fixture I wrote put the pin on the citation's own line and the corpus does not write
that way — **the third one-line-window bug of this milestone, and the first inside a test.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5, unchanged; clause 5 is still graded only by a reading that returns zero.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n
— (5) cap-reached: n (3 tiks of 5) — (6) protocol-stop: n — Outcome: **continue**.
**Decisions:** `D-M257x-71-1` (a citation is graded at the ref its own block names, not at a
process-wide knob; ambiguity is a counted refusal), `D-M257x-71-2` (a mutant survived — the window
bug for the third time, this time in the test), `D-M257x-71-3` (`run()`'s 7-positional tuple broke
four callers; recorded, not refactored).
**Side-deliverables:** none.
**Routes carried forward:**
- `RF-M257x-iter71-run-returns-a-tuple` — `anchor_construct_guard.run()` returns a positional
  7-tuple that has grown twice in four iterations. A dataclass would end the class; it touches
  `main()`, `postcondition_sites()` and every test.
- `CHECK-M257x-iter71-ambiguous-blocks` — the **12** citations whose block names two resolvable
  refs. They fall back by design, so they are graded at the default while their sentence names two
  refs. Whether that is a corpus-writing defect (a contrast should not sit in the same block as the
  claim it qualifies) or a fence limitation is **not settled here**.
- Unchanged: `FENCE-M257x-iter70-line-or-port` · `CHECK-M257x-iter70-studio-room-lines` ·
  `FIX-M257x-iter58-mainline-shift` · `FIX-M257x-iter53-union-set` (**PENDING USER DECISION**) ·
  `FIX-M257x-iter56-assignment-flake` (**NOT DECIDED**) ·
  `CHECK-M257x-iter38-ai-act-classification` (owner outside this milestone) ·
  `CHECK-M257x-iter57-anchor-guard-bare-class` · `FENCE-M257x-iter54-refs-block` ·
  `FIX-M257x-iter57-within-block-drift` · `CHECK-M257x-iter58-derive-preregistrations` ·
  `CHECK-M257x-iter52-second-ai-manager` · `-cold-daemon-registry` · `-grep-vs-failclosed` ·
  `-empty-stdout-class` · `-baseline-refs` · RF-2/3/7–13.
- **Closed here:** `FENCE-M257x-iter68-citation-resolution`.

**Lessons:**

1. **A guard's ref is per-CLAIM, not per-process.** iter-68's `CITE_REF` was the right first move and
   the wrong final shape: it made the guard say which file it read, but forced every claim to be
   read against the same one. The corpus states its refs *per block* — so the instrument has to.
2. **Run the battery against a STATED baseline.** Four callers broke on a one-line signature change
   and the only reason it surfaced is that the battery's first line said `4 errors` where I expected
   `OK`. A battery compared against a remembered baseline would have called all seven mutants caught.
3. **Your fixtures are the thing most likely to agree with your bug.** Third one-line-window bug of
   this milestone, first one inside a test. The rule is written down in two places already; what
   made it apply this time was a mutant, not the rule.
4. **A fence that cannot decide must say it cannot.** 12 ambiguous blocks fall back to the default —
   the same answer as before the change — but they now appear in the reach line under their own
   name. The verdict is unchanged; the honesty of the verdict is not.
