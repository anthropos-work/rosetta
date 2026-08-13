**Type:** tik — under `TOK-05`, consuming `FENCE-M257x-iter66-tier-membership` one iteration after
iter-66 opened it.

# iter-67 — G7: the service list beside a profile

## What was unfenced

iter-66 corrected a sentence placing `storage` in the default selection, and recorded that **no fence
could have caught it**. G1 asks whether a documented token is legal and selects *something*; G3 checks
the default's *count*. Nothing checked the **list** — so a row could name the right profile and the
wrong services indefinitely.

A profile-reference table states that list in a construct that *is* mechanically decidable, and the
legal set is already in hand: `compose.beyond_floor(tok)`.

## G7

The services column is found by its **header** (`| Profile | Services started … |`), exactly as
iter-63 taught the profile column to be; its cells are read **by shape** (comma / `+` / `and`
separated, emphasis and backticks stripped, kept only if it is a compose service name). Both
directions are reported — **MISSING** (the profile starts it, the doc omits it) and **NOT STARTED**
(the doc names it, the profile does not start it).

A prose cell (*"all backend services"*, *"the rollback path only"*) yields no service tokens and is
counted **UNREACHED**, never as an empty claim. Silence and zero are different findings.

**Reach on the live corpus: 22 membership rows, 12 checked, 10 UNREACHED — GREEN.** iter-62 had
already repaired every profile table by hand, so the assertion locks a correct state rather than
finding a defect; that is the right outcome for a fence built one iter after its class was repaired,
and it is *not* the evidence the rule works. The fixtures and mutants are.

## The reach bug the corpus's own spelling caused

The **most important row in the corpus** — `` | `core` *(default — `PROFILE ?= core`)* | … | `` — was
invisible to both G1 and G7. `_cell_profile_tokens` stripped a bare `(default)` and backticks, but this
qualifier is *emphasised* and contains its own backticks, so nothing matched and the cell yielded no
token. Fixed by stripping a trailing parenthetical (with any surrounding emphasis) before the
default-mark, which is derived from the format rather than from this one spelling. Rows checked went
**10 → 12**; profile sites **91 → 94**.

## Watched RED

| mutant | tests RED |
|---|---|
| compare against the whole selection, floor included | 3 |
| never report a mismatch (the no-op control) | 3 |
| revert the profile-cell qualifier strip | **0 → 1** |

The third is the finding worth keeping. Its first assertion was `membership_rows >= 1`, which passes
on the base fixture corpus alone and therefore **survived the mutant** — a test that cannot fail
proves nothing (§8 rule 2). Re-written to assert the **delta** (`base + 1`), it goes RED. **The
mutation battery caught a weak test, not a weak rule** — which is the only way that class of defect
gets found.

## Gates

| gate | result |
|---|---|
| `platform_predicate_guard` (now **seven** assertions) | **OK** — 94 profile sites / 8 tokens; 22 membership rows, 10 UNREACHED |
| `platform_alignment_guard` · `anchor_construct_guard` · `markdown_structure_guard` · `corpus_index_guard` | OK |
| `tests/test_platform_predicate_guard.py` | **67 tests** (was 60), all pass |
| `stack-core` suite | **682 tests, 1F** — the perishable iter-48 fixture |
| the second failure, and what it was | `test_m220_mutation_battery` went RED **because I left a previous suite running**; solo re-run **10/10 OK**. The same self-inflicted contention iter-63 recorded — and I caused it again in the same session |

## Close — 2026-08-04

**Outcome:** G7 lands the predicate iter-66 named one iteration earlier — *the services beside a
profile must be the services it starts* — with both directions reported and prose cells counted
UNREACHED rather than empty. Live GREEN at 12 checked rows; the value is the lock, not a catch. Along
the way the corpus's **most important profile row** turned out to be invisible to G1 as well, defeated
by its own `*(default — …)*` spelling, and the mutation battery caught a **weak test** of mine that a
`>= 1` assertion had made unfalsifiable.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5, unchanged; clause 5 is still graded only by a reading that returns zero.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: **y** — (6) protocol-stop: n — Outcome: exit-5
**Decisions:** `D-M257x-67-1` (G7: membership is decidable in the table construct; both directions,
prose cells UNREACHED), `D-M257x-67-2` (the `*(default — …)*` spelling hid the corpus's most important
profile row from G1 too), `D-M257x-67-3` (the mutation battery caught a weak TEST — a `>= 1` assertion
that could not fail).
**Side-deliverables:** none.
**Routes carried forward:** unchanged from iter-66, minus `FENCE-M257x-iter66-tier-membership`
(closed here). Still open: `FIX-M257x-iter63-app-citation-residual` (routed WHOLE) ·
`CHECK-M257x-iter63-quoting-a-retired-token` · `FIX-M257x-iter53-union-set` (**PENDING USER
DECISION**) · `FIX-M257x-iter56-assignment-flake` (**NOT DECIDED** — needs a failure *rate*) ·
`CHECK-M257x-iter38-ai-act-classification` (needs an owner outside this milestone) ·
`CHECK-M257x-iter57-anchor-guard-bare-class` · `FENCE-M257x-iter54-refs-block` ·
`FIX-M257x-iter57-within-block-drift` · `CHECK-M257x-iter58-derive-preregistrations` ·
`CHECK-M257x-iter52-second-ai-manager` · `-cold-daemon-registry` · `-grep-vs-failclosed` ·
`-empty-stdout-class` · `-baseline-refs` · RF-2/3/7–13.

**Lessons:**

1. **Build the fence in the iter after the one that names the class.** iter-66 found the defect in
   prose and sized the assertion; iter-67 built it. The gap was one iteration, and the sizing was
   still accurate — which is the argument for routing forward *with a design*, not with a title.
2. **A mutation battery's real job is auditing your tests.** Two of three mutants confirmed the rule.
   The third confirmed that one of my tests could not fail. That is the more valuable result.
3. **A fence built right after a hand-repair will be GREEN, and that is fine — say so.** The value is
   the lock on a correct state, not a catch. Reporting it as a catch would be the honesty failure
   this milestone keeps naming.
4. **I repeated iter-63's own contention mistake within the same session.** A second suite was still
   running when I started this one, and the m220 battery went RED for it. Written down twice now;
   the fix is to check `pgrep` before starting a suite, not to remember.
