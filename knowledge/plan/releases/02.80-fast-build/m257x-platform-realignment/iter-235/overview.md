---
iter: 235
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-10
---

# iter-235 — does every `make` target the corpus tells you to run still exist?

## Step 0 — Re-survey

iter-234 closed the citation ladder by **refusing** its top rung: a corpus quotes source by paraphrase, so
containment cannot grade it. That refusal is specific, not general — it kills instruments that must match
**free text**. It says nothing about claims whose subject is an **enumerable set**.

The user's redirect names two halves: *"the corpus's claims about the platform"* and *"be able to build a
working stack."* There is one class that is **both at once and is a closed set**: the corpus's
**runnable instructions**. `CLAUDE.md`, `setup_guide.md`, `run_guide.md`, `update_guide.md` and the demo
family tell a human — or an agent driving `/dev-up` — to run `make <target>`. A Makefile's targets are
**enumerable from the Makefile**. A documented target either is in that set or it is not. No sentence has
to be interpreted, and no paraphrase can blur it: `make up` is not a paraphrase of anything, it is a
command that runs or fails.

This is the same lesson the corpus already learned once and wrote down for compose profiles — *"grade a
documented command on **does it still select anything**, never on **does it still parse**"* (fenced by
`platform_predicate_guard` G1/G3, after the retired `graphql` profile was found to exit 0 while starting
an application-less stack). **That lesson was never applied to `make`**, which is the layer *above*
profiles and the actual entry point every guide opens with.

**Active strategy reference:** `TOK-08` — census a mechanical class exhaustively.

## Hypothesis

`platform`'s Makefile has been rewritten across the v8/v9 merge program (`0dab54d` renamed the default
profile, `838d907` deleted three services and cut `repos.yml` down). Guides written before those commits
should carry at least one target that no longer exists — and a missing `make` target is a **harder** failure
than a stale profile token: it exits 2 with *"No rule to make target"* rather than silently doing nothing.

## Predictions — SEALED BEFORE MEASUREMENT

| id | prediction |
|----|-----------|
| `P-235-1` | ≥ 40 distinct `make <target>` invocations appear across `corpus/**` + `CLAUDE.md` |
| `P-235-2` | `stack-demo/platform/Makefile` declares ≥ 20 targets |
| `P-235-3` | ≥ 1 documented `make` target does **not** exist in the current Makefile |
| `P-235-4` | ≥ 1 non-existent target appears inside a **fenced, copy-pasteable** block, not only in prose |
| `P-235-5` | the misses concentrate in the **ops guides** (setup / run / update / demo) rather than in architecture docs |

## Expected lift

No `N`/`P` reading. Deliverable: the documented-target set with its denominator, the Makefile's declared
target set, the difference **in both directions**, each miss classified (renamed / deleted / never-existed /
belongs-to-another-repo's-Makefile), and repair of any miss that sits in a runnable block.

## Phase plan

1. Enumerate declared targets from `stack-demo/platform/Makefile` (and note any other Makefiles the corpus
   addresses, since `make` is repo-relative — a target can be real in the wrong repo).
2. Enumerate documented `make <target>` invocations, recording fenced-vs-prose context per site.
3. Diff both directions; a Makefile target no corpus document ever mentions is a **discoverability** gap,
   not a defect — report it separately and do not conflate the two.
4. Prove the instrument on a control it cannot have fitted.
5. Classify and repair only misses that would actually fail for a reader.

## Escalation conditions

- The census returns 0 documented invocations → the regex is the finding (`§9`).
- A miss turns out to require a platform edit → route, never edit the platform.

## Acceptable close-no-lift outcomes

A measured 0 misses with a proven-non-vacuous instrument closes the class and refutes `P-235-3`/`P-235-4`
on the seal — the mechanism working, exactly as at iter-234.
