---
iter: 236
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
---

# iter-236 — can you `cd` where the corpus tells you to?

## Step 0 — Re-survey

iter-235 closed the **target** half of a runnable instruction (`make <x>`) — aligned, 0 dead of 394 — and
found its real defect in the **directory** half, which nothing had ever asked. It asked that half only for
the eleven *archived-service* names, found 13 fenced copy-pasteable sites and repaired 2, and routed the
general form: `ROUTE-M257x-235-runnable-block-has-two-halves`.

This iter takes that route. The full class is **every fenced `cd <path>` in the corpus**, graded against
what a stack actually has — and it is the most direct mechanical statement of the user's redirect
(*"be able to build a working stack with the new platform repos, only the remaining ones"*) that the corpus
admits. A guide whose `cd` fails is a guide that cannot be followed, and unlike a prose claim it needs no
interpretation.

**A stack's directory set is enumerable, and from the platform's own file.** `repos.yml` @ `0c91421`
declares four clones — `app`, `sentinel`, `next-web-app`, `studio-desk` — beside `platform` itself, plus
the documented manual clone `ant-academy` (CLAUDE.md: not in `repos.yml` by design, v1.10b M49 #5). That is
the denominator. Everything else a fenced `cd` names is either a subdirectory of one of those, a stack
workspace, a `.agentspace` path, or a defect.

**Active strategy reference:** `TOK-08` — census a mechanical class exhaustively.

## Hypothesis

The workspace convention moved twice (bare repo dirs → `stack-dev/` → the `stack-*/` family) and the clone
set shrank from ~13 repos to 4. Fenced `cd` lines written across those changes should disagree with each
other about the **prefix** even where they agree about the repo — and prefix drift is invisible to every
existing guard, because each individual path is well-formed.

## Predictions — SEALED BEFORE MEASUREMENT

| id | prediction |
|----|-----------|
| `P-236-1` | ≥ 100 fenced `cd <path>` sites across `corpus/**` + `CLAUDE.md` + `.claude/skills/**` |
| `P-236-2` | ≥ 20 distinct directory targets |
| `P-236-3` | ≥ 1 fenced `cd` names a directory a fresh stack does not have **and** carries no historical/disclosed caveat nearby |
| `P-236-4` | the same repo is reached under **≥ 2 different prefixes** across the corpus (e.g. bare vs `stack-dev/`) — prefix drift is real |
| `P-236-5` | ≥ 1 fenced `cd` names a path that exists in **no** form — neither a clone, nor a subdirectory, nor a workspace |

## Expected lift

No `N`/`P` reading. Deliverable: the fenced-`cd` population with its denominator, each target classified
(clone-root / subdirectory-of-a-clone / workspace / agentspace / archived-disclosed /
archived-undisclosed / nonexistent), the prefix-drift measurement, and repair of any undisclosed or
nonexistent site.

## Phase plan

1. Derive the stack's real directory denominator from `repos.yml` @ platform HEAD + the documented manual
   clone, not from memory.
2. Enumerate fenced `cd <path>` sites; record the enclosing document and whether a disclosure sits nearby.
3. Classify each target against the denominator, resolving subdirectories against the actual clone trees.
4. Measure prefix drift per repo.
5. Prove the instrument fires on both arms; repair only undisclosed defects.

## Escalation conditions

- 0 fenced sites → the fence-detection is the finding (`§9`).
- A defect needs a platform edit → route, never edit the platform.

## Acceptable close-no-lift outcomes

0 undisclosed defects with a proven-non-vacuous instrument closes the class and refutes `P-236-3`/`P-236-5`
on the seal — as at iter-235, where the refutation was the useful part.
