---
iter: 246
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
---

# iter-246 — census the `CLAUDE.md` fences, the milestone's longest-standing structural lead

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

## Step 0 — re-survey

`ROUTE-M257x-238-claude-md-fences-are-unmaintained` is **six-for-six**: six consecutive iters each found a
`CLAUDE.md` **code fence** contradicting `CLAUDE.md` **prose** written later. The stated cause is
structural — *the prose is swept by `/update-knowledge`; the fences are not.*

**Six iters found six defects, one each. That is a SAMPLE, and `TOK-08` says stop sampling.** The
population is small enough to enumerate completely, which is exactly the condition the strategy asks for:

| | |
|---|---|
| fenced blocks in `CLAUDE.md` | **11** |
| runnable command lines | **55** |
| comment lines inside fences (they carry port numbers and `file:line` claims) | **12** |
| blank | 6 |

Every prior finding came from a *different* iter noticing one line in passing. **Nobody has ever read all
55.**

## Hypothesis

The fence body is a runnable-input surface that no sweep and no guard treats as one. `skill_invocation_guard`
(iter-239) reads fenced lines that begin with `/`; nothing reads the rest. If the route's causal story is
right, a complete census finds more than the incidental rate of one-per-iter — and the residue tells us
which mechanical rule is worth a fence.

## Pre-registered numeric claims — SEALED IN THIS COMMIT

Graded against the clone set `stack-demo/` (13 repos) and the rext tree, both named.

| id | claim | prediction |
|---|---|---|
| **P-246-1** | runnable command lines (of 55) that fail against the clone set | **≤ 6** |
| **P-246-2** | ≥ 1 **comment-borne** claim inside a fence is false (a port, a `file:line`, a count) | **YES** |
| **P-246-3** | no existing guard takes a `CLAUDE.md` fence body as its subject, except `skill_invocation_guard` for `/`-lines | **CONFIRMED** |
| **P-246-4** | ≥ 1 `cd` target named in a fence does not exist under the parent the same document documents | **YES** |
| **P-246-5** | **control** — `make init # Clone the 4 repos in repos.yml: app, sentinel, next-web-app, studio-desk` is CORRECT | **HELD** (repos.yml declares exactly those four) |

`P-246-5` is a **positive control**, and it is here because an instrument that flags everything also
"finds" every defect. A census that cannot report a correct line as correct is not a census.

**Falsification:** if the full census of 55 lines finds **≤ 1** defect, the six-for-six pattern was
**coincidence, not structure** — six iters each happening to touch the same document. The route then closes
as a mis-diagnosis, which is a real and reportable outcome.

## Phase plan

A — enumerate the 55 + 12 lines and their checkable tokens. B — grade each against the clone set.
C — repair. D — fence whatever mechanical rule the residue supports (or record why none is supportable).
E — re-derive last.

## Escalation

A fence line that is correct for a **stack that no longer exists** is a documentation-scope question, not
a typo — disposition it, do not silently modernise it.
