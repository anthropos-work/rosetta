---
iteration_type: tik
status: in-flight
active_strategy: TOK-08
---

# iter-209 — the citation fence reads 94 documents; CLAUDE.md's own maintenance contract binds 5 more that it cannot see

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07)
— *census the mechanical classes; stop sampling them.* Intra-corpus mis-citation is `TOK-08`'s named
first class, and `corpus_citation_guard` is the census that closed it. This iter turns the same question
on **that census's own source set**.

## Step 0 — re-survey before targeting

Two candidate targets were sized before choosing, and both readings are recorded because the rejected
one is a genuine zero:

1. **Location-keyed registries.** 23 module-level registries in rext hold **165** keys naming a file
   (`DECISIONS` 92, `BARE_FILE_ALLOW` 16, `ENV_GATED` 9, …). Graded for path resolution: **19
   unresolved, all 19 correct-by-design** — `BARE_FILE_ALLOW` is an allow-list of *basenames* and
   `NAME_SOURCE_FILES` holds platform-relative paths. **A true zero**, so `§9` applies and the class was
   not worth an iter without an instrument-proof; recorded as a survey and dropped.
2. **Markdown anchor fragments.** `corpus_citation_guard` already asserts this as **C2**, and it reports
   `C2 anchor resolution 193 … OK`. So the class is closed — *within the guard's source set.*

The second reading is where the iter is, and it is the same shape iter-208 just found one level up:
**the verdict is sound and the denominator is undeclared.**

## Cluster / target identified

`collect_sources()` = `corpus/**/*.md` + `EXTRA_SOURCES = ("README.md", "CLAUDE.md")` — **94
documents.** The tree holds **2,560** `.md` outside `stack-*/` and `.agentspace/`.

Most of that gap is correct and already argued: `knowledge/` is planning, `iter-NN/raw/` is captured
evidence, and the docstring records measured false-RED counts for several exclusions. **One slice of it
is not.** `CLAUDE.md` § *Interconnected Documentation* opens:

> **These files must be maintained together:**

and then names **eleven** files — six under `corpus/ops/` and **five under `.claude/skills/`**. The
fence reads `CLAUDE.md`, which names them, and does not read them.

Measured with the guard's **own** machinery (`resolve_link`, `anchors_of`, `heading_slugs`, `MD_LINK`,
`CORPUS_PATH`), so the only variable is the source set:

| | |
|---|---|
| skill documents | **20** (`.claude/skills/*/*.md`) |
| citations they carry | **135** — C1 133, C2 2 |
| C1 findings | **0 of 133** |
| C2 findings | **1** |

The one finding is real and is exactly the drift the contract exists to prevent:
`.claude/skills/stack-secrets/SKILL.md:142` links `corpus/ops/safety.md#…-v16-m27m28`, and that
heading now reads **M27–M30**. The milestone range grew, the heading followed, the anchor did not.

## Hypothesis

The fence's source set is a **registry** (`§5` iter-184, *a fence's POPULATION is a registry too*) and
has never been graded against the one contract in this repo that says which documents move together.
Grading it costs 20 documents and finds a live break with **zero** false REDs.

## Expected lift

1. `collect_sources()` covers the documents `CLAUDE.md`'s maintenance contract binds together, and an
   arm **derives that list from `CLAUDE.md`** rather than restating it — so a twelfth file added to the
   contract enrols itself or turns the arm RED.
2. The live C2 break is repaired.
3. The residual exclusion is **sized** rather than implied.

## Phase plan

Two planned lines:

1. Widen `collect_sources` to the skill docs; repair the C2 break; re-run the guard on the real tree.
2. Arms: contract-coverage derived from `CLAUDE.md`; a mutation control that drops a contract file from
   the source set and requires RED.

## Escalation conditions

- If widening produces false REDs, **narrow or revert** — this guard's docstring records that three of
  its first four runs existed only to kill a false-positive class, and `§8` rule 6 is unforgiving: a
  fence that cries wolf gets disabled.

## Acceptable close-no-lift outcomes

- The widening finds nothing beyond the one already-measured break → the source set still gains a
  derivation and the exclusion still gains a size, which is `§5`'s *sized, not argued*.
