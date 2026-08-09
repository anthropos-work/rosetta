**Type:** tik — under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

# iter-191 — the guard says "all 119 scanned doc(s) agree" and scanned 164 files

## Phase A — the re-survey moved the target one line

`SURVEY-M257x-iter188-the-other-walks-are-unmeasured` named three prune rules; this is the third and last.
Re-surveyed, **the prune list is fine** — component-matched, reasoned, carrying the story of the bug it
memorialises. The finding is one line down, in `main()` (`D-M257x-191-1`):

| | files |
|---|---:|
| printed — *"all N scanned doc(s) agree"*, and the **refusal gate**, from `rglob("*.md")` | **119** |
| actually walked by `find_violations` over `_SCANNED_SUFFIXES` (`.md` 119 · `.yaml` 33 · `.sh` 9 · `.yml` 3) | **164** |

**The published number described 72.6 % of what the guard checked**, and the same expression gates the
refusal, so the error runs **both ways** (`D-M257x-191-2`): a scope with no markdown but 45 shell/YAML
sites exits 2 with *"0 markdown file(s) in scope. Nothing was checked; this is not GREEN"* — a **false
CANNOT-RUN** over a scope with 45 files to check. Refusing looks conservative, which is why it is the
harder direction to notice.

**And the non-markdown half is not empty of the claim class**, measured before repairing: **25 `.sh` +
99 `.yaml`** lines match the guard's own context pattern. The markdown-only denominator was not a
harmless label on an all-markdown population.

**Second line, latent, and it is the other half of a bug this module already tells** (`D-M257x-191-4`).
`_EXCLUDED_DIRS`' comment memorialises a real RED — the list held `.agentspace`, matched as a substring
of the **absolute** path, and silently skipped both of rext's own sites. The repair fixed the *matching
mode* (*"match components; never substrings"*) and left the *scope*: `_excluded` still tested the
absolute path, so a checkout under any directory named `knowledge` or `node_modules` excludes everything
beneath it. Both halves were in the same sentence of the original bug report. Sized: **0 difference**
here, and the refusal fails closed.

## Phase B — one derivation, printed

`in_scope(roots) -> {suffix: count}` and `excluded_reach(roots)`; `main()` reads both, the refusal keys on
the true total, and the scope is printed ahead of the verdict:

```
story-org-count-guard: scope — 119 .md, 9 .sh, 33 .yaml, 3 .yml = 164 file(s); 0 excluded by 4 pruned dir name(s)
story-org-count-guard: OK — stories.seed.yaml ships 4 orgs, stories-maya.seed.yaml ships 1, and all 164 scanned file(s) agree
```

The unit word moved with the number — *"doc(s)"* → *"file(s)"* — because `.sh` and `.yaml` are not docs,
and the old word is what made a markdown-only count read as correct. The reach is printed **as zero**
(`D-M257x-191-6`): *"4 names excluded nothing here"* is the fact a reader needs in order to size the list.

## Phase C/D — the fence, and an existing arm re-based off a spelling

`test_the_real_tree_is_still_GREEN_and_states_its_cardinality` asserted the regex `all \d+ scanned
doc\(s\) agree` — a **spelling**, so it stayed green while the number described 72.6 % of the population,
and would have stayed green after any future narrowing (`§5` r70/71). It now parses the cardinality and
asserts it **equals** `sum(in_scope(scan_roots(root)).values())` (`D-M257x-191-3`).

New class `TheDENOMINATORIsDerivedNotRestated`, **5 arms**: the denominator covers **every** scanned
suffix · it is **strictly larger** than the markdown half (the `§9` control — every other arm here is
satisfied by a derivation that only ever counts markdown) · `main()` reads the same derivation and no
longer contains `rglob("*.md")` · the exclusion is **scoped to the scan root** · the prune list's reach is
derived and can return non-zero.

**5/5 mutants RED:**

```
RED ✔ M1 denominator back to markdown only (2 arms)   RED ✔ M4 exclusion back to absolute scope
RED ✔ M2 non-markdown suffixes contribute nothing     RED ✔ M5 excluded_reach can only return zero
RED ✔ M3 main() stops reading the derivation
```

## Runs — runner and scope named (`§5` r60/75/76)

| scope | runner | result |
|---|---|---|
| `test_story_org_count_guard.py` (25 → **30** arms) | unittest 3.14.6 / pytest 8.4.2 (3.9.6) | **30 / 30 passed**, both |
| + `test_guard_family` + `test_claim_census_skip_registry` + `test_dual_reader_parity` | pytest | **94 passed · 0 failed** (16.8 s) |
| the guard itself, live | — | `OK`, **verdict unchanged**, number 119 → **164**, reach `0` |

**Not covered, stated:** the guard's scan roots are 4 of the repos' trees, unchanged by this iter; 264 Go
+ 75 TS still UNMEASURED; the dual-reader enumeration still covers `stack-core` only.

## Close — 2026-08-09

**Outcome:** the last member of iter-188's routed set turned out to be sound, and the defect was one line
below it: the guard **published 119 as the size of a population of 164** — 72.6 % — and keyed its
*refusal* on the same markdown-only count, so it would have reported *"nothing was checked"* over 45
shell/YAML files with 124 candidate lines between them. Denominator derived, scope and reach printed, the
unit word corrected, the exclusion scoped to the scan root, and an existing arm re-based from a spelling
onto the derivation. Verdict unchanged; 5 new arms, 5/5 mutants RED.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (twenty-third consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: **y** (fifth tik of this invocation) — (6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **exit-5**
**Decisions:** `D-M257x-191-1` … `D-M257x-191-6` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none.

**Routes carried forward:**
- `SURVEY-M257x-iter188-the-other-walks-are-unmeasured` — **CLOSED.** All three named members handled:
  `claim_census_guard._SKIP_DIRS` (iter-188), the `platform_predicate_guard` pair (iters 189–190),
  `story_org_count_guard._EXCLUDED_DIRS` (here — sound, with the defect one line below it).
- `SURVEY-M257x-iter191-published-denominators-are-unenumerated` — **NEW.** This defect's shape is *a
  guard printing a cardinality derived differently from the population it graded*, and it is the second
  instance in five iters (iter-186's suite census was the first). **No census exists for it**, and it has
  a mechanical selector: *a printed count whose derivation does not appear in the function that produced
  the verdict.*
- `SURVEY-M257x-iter190-the-dual-reader-census-covers-one-section-of-eleven` ·
  `SURVEY-M257x-iter190-one-construct-two-regexes-is-unenumerated` ·
  `SURVEY-M257x-iter187-the-grain-question-is-unasked-elsewhere` ·
  `SURVEY-M257x-iter186-264-go-tests-have-never-been-read` ·
  `SURVEY-M257x-iter185-other-declared-populations-unaudited` · `D-M257x-145-3` (the user's to rule) ·
  `SURVEY-M257x-iter185-reach-percentages-predate-the-citation-widening` ·
  `SURVEY-M257x-iter184-the-standing-queue-wildcard-is-unbounded` ·
  `SURVEY-M257x-iter183-segment-grammar-refuses-28-multi-id-segments` ·
  `SURVEY-M257x-iter179-thirty-battery-tests-unrun` ·
  `SURVEY-M257x-iter180-relation-grammar-supports-only-equality` · `FIX-M257x-iter173-ledger-denominator` ·
  `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` ·
  `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` (observed half) ·
  `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites` · `FIX-M257x-h36-labeled-prover-denominator`
  — unchanged; open. Standing queue unchanged.

**Lessons:** **a false CANNOT-RUN is as wrong as a false green, and much harder to see** — refusing looks
conservative, so a refusal gate keyed on the wrong unit can sit for four releases while reading as
caution. And the re-survey lesson, which is why the phase exists: **a routed subject can be sound while
the defect sits one line below it** — taking `_EXCLUDED_DIRS` at face value would have produced a
cosmetic iter. Written into `platform-alignment.md` §8 in this iter's commit.
