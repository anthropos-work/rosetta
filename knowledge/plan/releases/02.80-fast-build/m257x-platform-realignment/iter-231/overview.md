---
iter: 231
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
---

# iter-231 — the corpus's CURRENT-ref claims, graded against each other and against the clones

## Step 0 — Re-survey

iter-230 proved every measurable cited sha **exists**. Existing is the weak claim. The strong one is
positional: *"`X` is `<repo>`'s `origin/main`"* — a statement that decays the moment the platform ships,
and one the corpus makes in prose, repeatedly, in more than one document per repo.

iter-228 found exactly one instance by hand: `CLAUDE.md` named `ad9f3c49` as `app`'s `origin/main` while
**two other corpus docs already carried the newer ref**, and closed with
`ROUTE-M257x-228-corpus-disagrees-with-itself-about-refs` — *"nothing compares `CLAUDE.md`'s ref claims to
the corpus's. Decidable."* This iter censuses that class instead of sampling it, per `TOK-08`, and it is
squarely the redirect's (a): **the corpus's claims about the platform.**

**Active strategy reference:** `TOK-08`.

## Hypothesis

Current-ref claims are the fastest-rotting sentence type in the corpus, and the population is small enough
to enumerate exhaustively. Internal disagreement is **cheaper to detect than external drift** — both
operands are in this repo — and nothing looks for it.

## Predictions — SEALED BEFORE MEASUREMENT

| id | prediction |
|----|-----------|
| `P-231-1` | ≥ 3 distinct repos carry a corpus claim of the form "`<sha>` is/was `origin/main`" or "HEAD `<sha>`" |
| `P-231-2` | ≥ 1 repo has **two different** shas claimed as its current ref in **different** documents |
| `P-231-3` | ≥ 1 claimed-current sha disagrees with that clone's **actual** `origin/main` |
| `P-231-4` | for ≥ 1 repo the corpus is **split** — at least one document right where another is wrong |

## Expected lift

No `N`/`P` reading. Deliverable: the enumerated population of current-ref claims with each graded
`agrees` / `stale` / `unmeasurable`, and repair of anything found false.

## Escalation conditions

- A claim about an **uncloned** repo is `UNMEASURED` (iter-230's partition rule), never stale.
- A claim explicitly written as **historical** ("was `origin/main` on <date>") is **not** a current-ref
  claim and must not be graded as one — grading a dated past-tense sentence as stale is the
  `§5` retraction-vocabulary error in reverse.
- Building a standing fence is a second deliverable → tripwire → route forward.

## Acceptable close-no-lift outcomes

If every current-ref claim already agrees, the census + its instrument proof is the deliverable — provided
a known-bad control demonstrates the extractor can see a disagreement at all.
