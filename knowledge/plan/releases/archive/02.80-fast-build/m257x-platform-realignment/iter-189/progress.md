**Type:** tik — under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

# iter-189 — one population, two readers, two different exclusion rules — and the repo already wrote down the one they violate

## Phase A — the target, and the two sizings taken before touching it

iter-188 routed its residual **with its number**: three name-based prune rules in `stack-core`, of which
`_SKIP_DIRS` was one. The largest of the other two is `platform_predicate_guard`'s pair — and it is not a
prune list, it is **the same exclusion written twice**:

| reader | rule | shape |
|---|---|---|
| `_reads_worktree` | `parts = set(path.parts)`; `"vendor" in parts` | component-exact, on an **absolute** path |
| `_reads_at_ref` | `"vendor/" in rel` | **substring**, on a repo-relative path |

`_reads_at_ref`'s own docstring opens *"**Same derivation**, read from the object store instead of the
checkout."* `"cloud-vendor/x.go"` contains `"vendor/"`.

This is the reader behind **G6** — the guard whose output is the corpus's evidence for *"zero
`*_RPC_ADDR` variables anywhere"* — and on a **consumer** side over-exclusion means *a read that exists
is reported as absent*.

**And the rule was already written down in this repo, by a sibling guard, and broken here:**
`story_org_count_guard.py:125` — *"an exclusion that can swallow the whole repo. **Match components;
never substrings.**"*

**Sized twice before any edit** (`D-M257x-189-1`):

- **structural** — directories whose name *ends* in `vendor`/`node_modules` without being exactly that,
  across `stack-demo/` + the tooling repo: **0**;
- **behavioural** — the comparison nothing had ever run for this pair: both readers against the real
  `app` clone at `origin/main@ad9f3c49` → **both `{}`. They agree today.**

So: **latent**, and recorded as latent. Neither pre-registered escalation fired (`D-M257x-189-5`).

**A third defect was in the same three lines and is named separately** (`D-M257x-189-4`):
`_reads_worktree` tested the **absolute** path, so the exclusion depended on **where the clone lives** —
a checkout under any directory named `vendor` excludes *every* file and reports a confident zero for the
entire consumer side. Sized: 0 such ancestors here. Fixed to repo-relative, which is also what makes the
two readers comparable at all.

## Phase B — one predicate, not two reconciled rules

`VENDORED_PATH_COMPONENTS` + `is_vendored_path()`; both readers call it; the substring form is gone
(`D-M257x-189-3`). Asserting that two rules agree would have fenced the symptom and left the cause — two
literals expressing one exclusion, iter-177's shape, one file away from where iter-188 had just removed
it.

## Phase C/D — the fence and its mutants

`tests/test_rpc_reader_parity_m257x.py`, **5 arms in two classes**:

| class | arms |
|---|---|
| `ThePredicateMatchesCOMPONENTS` | 8-path table (component ≠ substring) · **both** readers call the one predicate and neither has a substring test · the worktree reader is **repo-relative** |
| `TheTwoReadersAGREE` | the two readers return **the same files** over a synthetic git tree · they return the **right** set, and the fixture can **detect** a divergence (`§9`) |

**5/5 mutants RED**, plus the reading that justifies the second arm's existence:

```
RED ✔ M1 at-ref goes back to the SUBSTRING rule      (3 arms)
RED ✔ M2 worktree tests the ABSOLUTE path
RED ✔ M3 the shared predicate becomes substring-based
RED ✔ M4 the predicate excludes NOTHING              (2 arms)
RED ✔ M5 the predicate excludes EVERYTHING
(as designed) agreement-alone under M5: 0 failures — which is exactly why the §9 arm exists
```

M5 is the point: **the readers still agree when the predicate excludes everything.** A parity check
without an expected set passes on two identical zeros — and the reading it was about to confirm *was* a
zero.

## Runs — runner and scope named (`§5` r60/75/76)

| scope | runner | result |
|---|---|---|
| `test_rpc_reader_parity_m257x.py` | unittest 3.14.6 / pytest 8.4.2 (3.9.6) | **5 / 5 passed**, both |
| + `test_platform_predicate_guard.py` | pytest | **184 passed · 0 failed** (6.9 s) |
| + `test_claim_census_skip_registry_m257x` + `test_guard_family` | unittest | **238 passed · 0 failed** (16.5 s) |
| the guard itself, live, against `stack-demo/platform` | — | `G6 8 RPC var(s) graded {'unconfigured': 8}, 0 mid-fold; app consumer side measured @ origin/main@ad9f3c4` → **OK**, unchanged by the refactor |

**Not covered, stated:** the third prune rule iter-188 named — `story_org_count_guard._EXCLUDED_DIRS` —
is untouched; the 264 Go + 75 TS remain UNMEASURED.

## Close — 2026-08-09

**Outcome:** the guard behind the corpus's *"zero `*_RPC_ADDR`"* claim read its population two ways, and
the two disagreed by construction — component-exact in one reader, **substring** in the other, 60 lines
from a docstring calling them the same derivation, and both violating a *"match components, never
substrings"* rule this repo had already written down in a sibling guard. One predicate now; a third
defect (an absolute-path test that made the exclusion depend on where the clone lives) fixed alongside
and named separately; 5 arms, 5/5 mutants RED, and the parity reading given the expected set it needed
because the agreement it confirms is `{} == {}`.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (twenty-first consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-189-1` … `D-M257x-189-5` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none.

**Routes carried forward:**
- `SURVEY-M257x-iter189-the-parity-question-is-unasked-for-every-other-dual-reader` — **NEW.** This pair
  was found by a prune-rule grep, not by looking for dual readers. `platform_predicate_guard` alone
  carries several worktree-vs-ref pairs, and **no arm anywhere compares any of them**; iter-175's rule
  (*two derivations of ONE population must be COMPARED*) has been applied case-by-case and never
  enumerated. The selector is mechanical: *a function whose docstring or name claims to be the same
  derivation as another.*
- `SURVEY-M257x-iter188-the-other-walks-are-unmeasured` — **advanced, not closed.** Two of the three
  named members are now handled (`_SKIP_DIRS` at iter-188, this pair here);
  `story_org_count_guard._EXCLUDED_DIRS` is untouched and its reach is still unreported.
- `SURVEY-M257x-iter187-the-grain-question-is-unasked-elsewhere` ·
  `SURVEY-M257x-iter186-264-go-tests-have-never-been-read` (now 264 Go + 75 TS) ·
  `SURVEY-M257x-iter185-other-declared-populations-unaudited` ·
  `D-M257x-145-3` (the user's to rule) ·
  `SURVEY-M257x-iter185-reach-percentages-predate-the-citation-widening` ·
  `SURVEY-M257x-iter184-the-standing-queue-wildcard-is-unbounded` ·
  `SURVEY-M257x-iter183-segment-grammar-refuses-28-multi-id-segments` ·
  `SURVEY-M257x-iter179-thirty-battery-tests-unrun` ·
  `SURVEY-M257x-iter180-relation-grammar-supports-only-equality` · `FIX-M257x-iter173-ledger-denominator` ·
  `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` ·
  `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` (observed half) ·
  `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites` — unchanged; open. Standing queue unchanged.

**Lessons:** **a rule this repo has already written down can be broken elsewhere in it, and nothing looks.**
`story_org_count_guard` states *match components, never substrings*; `platform_predicate_guard` broke it,
in the pair whose own docstring claims they are the same derivation. A stated-but-unfenced rule is a
comment. And the parity corollary: **a parity check without an expected set passes on two identical
zeros** — the agreement this iter set out to confirm *was* `{} == {}`, and M5 shows the arm would have
held with the predicate excluding everything. Written into `platform-alignment.md` §8 in this commit.
