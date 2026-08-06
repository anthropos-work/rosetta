# iter-109 — decisions

## `D-M257x-109-1` — a seat-commit subject named one seat and carried two. Recorded, not rewritten.

`57d34f6`'s subject reads *"seat C committed VERBATIM on landing"*. It carries **`r27-C.md` AND
`r27-D.md`** — seat D landed in the window between the notification and the `git add`, and a
directory-scoped `git add` swept both.

**The substance of the discipline held and was verified**: neither file was read, edited or graded before
it was committed, and `git show --name-only` on all five seat commits reconstructs exactly which seat
landed in which commit — `ed50c78` F, `bd7b088` B, `9ca8270` A, **`57d34f6` C+D**. The *record* is
recoverable; the *subject line* is wrong.

**Not amended.** Rewriting the commit would erase the evidence of the slip, which is the opposite of what
this milestone is for. A commit subject is a claim like any other, and the correction belongs beside it
rather than on top of it.

**Lesson, and it generalises past this iter:** `git add <dir>` under a concurrent producer stages whatever
has appeared, not what you believe appeared. When the commit *subject* is an assertion about scope, stage
the **named paths**, not the directory — otherwise the message and the content are free to disagree
silently. Same shape as `D-M257x-108-4`'s reach-denominator lesson: state what you measured over.
