---
iter: 232
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-10
---

# iter-232 — does the corpus cite files the platform has DELETED?

## Step 0 — Re-survey

iter-230 proved the corpus's **refs** exist. iter-231 proved its **current-ref claims** agree. Both are
about shas. The alignment question the milestone actually exists for is one level down and is still
unmeasured: **the corpus cites `path:line` into platform repos, and the platform has been deleting whole
services.** A citation into a file that no longer exists is not a rotted offset — it is a claim about a
thing that is gone, and it is **mechanically decidable**: `git cat-file -e <ref>:<path>`.

`anchor_construct_guard` resolves anchors and grades whether the cited LINE is a construct. It does not
report the **file-level** population: how many distinct paths the corpus cites per repo, and how many of
those the repo no longer has at its current `origin/main`. That is the cheapest decidable form of *"is the
corpus describing the platform as it is?"* and nothing has asked it.

**Active strategy reference:** `TOK-08` — census a mechanical class exhaustively.

## Hypothesis

The v8/v9 merge program deleted containers, workflows and whole service trees. A corpus that grew across
that program should still cite some paths that are gone — most likely in `app` (where files MOVED during
the folds) and in the archived repos (where whole trees were removed).

## Predictions — SEALED BEFORE MEASUREMENT

| id | prediction |
|----|-----------|
| `P-232-1` | ≥ 200 distinct (repo, path) pairs are cited across `corpus/**` + `CLAUDE.md` |
| `P-232-2` | ≥ 1 cited path does **not** exist at its repo's current `origin/main` |
| `P-232-3` | the missing paths concentrate in `app` — the repo the folds moved files *into* and *within* — rather than in the archived repos, whose trees were frozen, not rewritten |
| `P-232-4` | ≥ 1 missing path is cited in a **live, present-tense** claim rather than only in a dated/historical one |

## Expected lift

No `N`/`P` reading. Deliverable: the per-repo population of cited paths, the missing set with its
denominator, each missing path classified **gone** vs **moved** vs **never-existed**, and repair of any
live claim that names a deleted file.

## Escalation conditions

- A path in an **uncloned** repo is `UNMEASURED` (iter-230's partition rule), never missing.
- A path cited **at an explicit historical ref** ("`app/main.go:504` @ `2035f9a4`") is a claim about that
  ref and must be graded **there**, not at `origin/main` — grading it at HEAD is the iter-231 tense error
  wearing a different hat.
- Repairing prose is in scope. A standing fence is a second deliverable → tripwire → route forward.

## Acceptable close-no-lift outcomes

If every cited path still exists, that is a strong alignment result and the iter's deliverable — provided
the instrument is proved with a known-deleted control path.
