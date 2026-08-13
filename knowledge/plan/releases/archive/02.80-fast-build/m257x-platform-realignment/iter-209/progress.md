**Type:** tik — under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

# iter-209 — the citation fence reads 94 documents; CLAUDE.md's own contract binds 5 more it cannot see

## Two candidate classes were sized before one was chosen

`TOK-08` says work the mechanical classes in descending measured size. Both readings are recorded,
because the rejected one is a genuine zero and `§9` says a zero has to be reported with its instrument:

**Rejected — location-keyed registries.** 23 module-level registries in `rosetta-extensions` hold
**165** keys naming a file (`DECISIONS` 92, `BARE_FILE_ALLOW` 16, `ENV_GATED` 9, `N_OF_M_DISPOSITIONS`
8, `EXEMPT` 6, …, across 18 modules). Graded for path resolution against the rext root, the rosetta
root and `corpus/`: **19 unresolved, and all 19 are correct by design** — `BARE_FILE_ALLOW` is an
allow-list of *basenames* (`docker-compose.yml`, `main.go`) and `NAME_SOURCE_FILES` holds
platform-relative paths. **Zero dangling keys.** Routed as a survey, not worked.

**Chosen — the citation fence's own source set.** `corpus_citation_guard` already asserts anchor
resolution as **C2** and reports `C2 anchor resolution 193 … OK`. The class is closed *inside the
guard's source set* — and that set had never been graded. Same shape as iter-208 one level up: **the
verdict is sound and the denominator is undeclared.**

## The finding

`collect_sources()` was `corpus/**/*.md` + `EXTRA_SOURCES = ("README.md", "CLAUDE.md")` — **94
documents**, against **2,560** `.md` in the tree outside `stack-*/` and `.agentspace/`.

Nearly all of that gap is correct and already argued in the guard's docstring, which records measured
false-RED counts for each exclusion it adopted. **One slice was not.** `CLAUDE.md` §
*Interconnected Documentation* opens

> **These files must be maintained together:**

and names **eleven** documents — six under `corpus/ops/` and **five under `.claude/skills/`**. The fence
read `CLAUDE.md`, which names them, and did not read them. `§5` iter-184: **a fence's POPULATION is a
registry too.**

Measured *before* widening, using the guard's **own** `resolve_link` / `anchors_of` / `heading_slugs` /
`MD_LINK` / `CORPUS_PATH`, so the source set was the only variable:

| | |
|---|---|
| skill documents (`.claude/skills/*/*.md`) | **20** |
| citations they carry | **135** — C1 133, C2 2 |
| C1 findings | **0 of 133** |
| C2 findings | **1** |

**Zero false REDs**, which was the condition for widening at all: this guard's docstring records that
three of its first four runs existed only to kill a false-positive class, and `§8` rule 6 is
unforgiving — a fence that cries wolf gets disabled.

The one finding is real, and it is precisely the drift the contract exists to prevent:

> `.claude/skills/stack-secrets/SKILL.md:142` links
> `corpus/ops/safety.md#…-prod-write-path-v16-**m27m28**`.
> The heading is `### 2.9 … (v1.6 M27–M30)` → slug `…-v16-**m27m30**`.

**The milestone range grew from M27–M28 to M27–M30, the heading followed, and the anchor did not** — in
one of the five skill files `CLAUDE.md` names, invisible to the census built to catch exactly this.

## An instrument note worth more than the finding

The first draft of this iter's independent reading used a hand-written GitHub slugger and reported
**97 unresolved of 190**. Adjudicated: the slugger, not the corpus. It collapsed runs of spaces
(`—` between words leaves two) and stripped `_`. **The shipped `heading_slugs` is right and emits two
variants deliberately.** After adopting the guard's own slugger the same reading returned **6**.

**16× wrong, entirely in one direction.** That is iter-201's measured shape again — 18 false-RED /
0 false-GREEN — and the reason this iter reused the guard's machinery instead of re-deriving it: `§5`
iter-175, *two derivations of ONE population must be COMPARED, or the weaker one is a silent census*.
Here the weaker one was mine.

## What shipped

- `corpus_citation_guard.SKILL_SOURCE_GLOB = ".claude/skills/*/*.md"`, added to `collect_sources`, with
  the pre-widening measurement recorded beside it.
- The broken anchor repaired (`m27m28` → `m27m30`).
- **Four broken anchors of this run's own making**, repaired: iters 207 and 208 each linked
  `../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them` while the heading carries a
  trailing `— 2026-08-07`, so the real slug ends `…-them--2026-08-07`. Written twice per iter, in the
  four documents this run authored, and found by the census built in the same run.
- **A unit correction in iter-208's own headline**, left visible in place: it read *"the repo derives
  **five**"* where the clause under audit says *non-`stack-core`*, which is **four**. Five includes
  `stack-core`. Comparing ten against five is the apples-to-oranges that iter is about — committed in
  its headline while its body and `progress.md` both said four. `§5` r75, *name the unit*.
- Three arms in `tests/test_corpus_citation_guard.py`
  (`TheSourceSetHONOURSTheRepoOwnMaintenanceContract`): the contract list is **parsed out of
  `CLAUDE.md`**, never restated, so a twelfth file enrols itself or the arm goes RED; an anti-vacuity
  arm requiring the skill glob to actually contribute sources *and* the census to enumerate a non-zero
  C1 and C2 population on the real tree; and a staged mutation control with a contract file outside
  every glob, so the first arm is provably able to fire.

Guard on the real tree, after the widening: **1,801 citations over 114 source documents** (C1 1,601 ·
C2 195 · C3 5) — **OK, every enumerated citation resolves.** Before: 1,666 over 94.

## Close — 2026-08-09

**Outcome:** the intra-corpus citation census — `TOK-08`'s first and largest named class — was reading
**94 of the repo's documents** and had never stated that as a denominator. Five of the documents it
skipped are bound to the corpus by `CLAUDE.md`'s own maintenance contract; reading them cost **zero**
false REDs and found **one** live broken anchor, plus four more this run had just written. The source
set now derives its floor from that contract instead of from a hand-written tuple.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (forty-first consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — **counted: iters 207, 208, 209 = three tiks this run against a cap of five** —
(6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-209-1` … `D-M257x-209-3` (see [`decisions.md`](decisions.md))

**Audit:** `/usr/bin/python3 -m pytest` (**8.4.2** / CPython **3.9.6**), **Python**, `stack-core` only —
**22 passed** in `test_corpus_citation_guard.py` (the changed fence, +3 arms; the three new arms
verified to **run rather than skip**, `3 passed, 19 deselected`). `corpus_citation_guard` on the real
rosetta tree: **1,801 citations over 114 sources, 0 findings.**
**RED-proof battery, mtime-mitigated (`§5` r77):** the widening line was deleted from `collect_sources`
— **both live arms went RED and the staged mutation control stayed green**, which is correct: it does
not depend on the glob. Restore sha-verified against `2debebea…`.
*Scope, stated rather than implied (`§5` r60): `stack-core` only, Python only, changed-code reach. The
guard's residual exclusion — **2,446 `.md` outside its source set**, almost all of it `knowledge/`
planning and `iter-NN/raw/` captured evidence — is **now sized and still excluded**; this iter widened
it by the five contract files and nothing else. No Go, no TypeScript; the other four Python sections
were read at iter-208 and not re-read here.*

**Side-deliverables:** none this iter.

**Routes carried forward:**
- `SURVEY-M257x-iter209-location-keyed-registries-are-a-clean-zero` — **NEW.** 23 registries, 165
  location keys, 0 dangling; the 19 non-resolving keys are basenames and platform-relative paths, both
  by design. No fence enumerates them, and a dangling key would be silent. A clean zero today, which is
  exactly when `§9` says the instrument must be proved before the class is closed.
- `SURVEY-M257x-iter209-the-guard-reads-94-of-2560-markdown-documents` — **NEW.** Sized, not repaired.
  The `knowledge/` exclusion is deliberate and argued; what is new is that it now has a number. The
  16 unresolved anchors outside the source set are: 1 repaired here, 4 repaired here (this run's own),
  and 11 in `iter-122/raw/` capture files quoting corpus text verbatim.
- `SURVEY-M257x-iter208-env-gated-keys-are-not-nodeids` · `SURVEY-M257x-iter208-the-wrong-clause-is-
  still-in-the-hardening-ledger` · `SURVEY-M257x-iter208-a-language-triples-third-leg-was-missing-for-
  twelve-iters` — unchanged, all from iter-208.
- All of iter-207's routes, unchanged, plus the standing queue.

**Lessons:**
- **A fence's source set is a registry, and the repo may already state what belongs in it.** Here the
  statement was three screens above the guard's own entry in the same `CLAUDE.md` the guard reads.
- **Reuse the shipped derivation when you audit its denominator.** An independently written slugger was
  **16× wrong in one direction** on the first pass; the audit is of the *scope*, so the *machinery* must
  be held fixed.
- **A run that builds a census will be caught by it.** Four of the six anchors repaired this iter were
  written by iters 207 and 208, two hours earlier.
