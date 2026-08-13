**Type:** tik, under `TOK-05`. Answers `CHECK-M257x-iter91-guard-repo-root-scoping`.

# iter-94 — the family's own green: one member was reporting a pass over zero docs

## Why this mattered enough to spend a tik on

**The guard family's green is the evidence this milestone quotes.** A member that can report a pass it did
not earn contaminates every reading taken alongside it — including the clause-5 reading still owed.

iter-91's mutation run had shown the family, against an empty temp dir, reporting
`2 GREEN · 2 RED · 9 could-not-check`. Nine members correctly said *COULD NOT RUN — no corpus/*. Two said
**GREEN**.

## Two suspects, one defect — separating them was the work

| guard | subject | verdict |
|---|---|---|
| `union_apply_guard` | rext's own demopatch manifests | **correct by design.** The corpus is not its subject; making it honour `--repo-root` would break it in a rext-only checkout, which is how the family is consumed per-stack. Not changed. |
| `story_org_count_guard` | *"…**and every doc agrees**"* — those docs are the **corpus** | **real defect.** Fixed. |

A shared symptom is not a shared cause. *"It ignored `--repo-root`"* is only a defect if `--repo-root` names
the guard's subject.

## The defect: a control that could never fire

`scan_roots()` returns the corpus + skills **plus rext's own two directories** — and the guard lives in
rext, so those always exist. The existing `if not roots: return 2` **could never fire.** A run whose rosetta
half was missing scanned only rext, found nothing to contradict, and printed *"and every doc agrees"*.

Fixed with a positive control on **the corpus half specifically**, plus the cardinality:

| tree | before | after |
|---|---|---|
| empty | `rc=0 OK — … every doc agrees` | **`rc=2 CANNOT RUN — Nothing was checked; this is not GREEN`** |
| real | `rc=0 OK — … every doc agrees` | `rc=0 OK — all **116** scanned doc(s) agree` |

Both directions are tested, because the empty-tree test alone would pass for the wrong reason.

## The pattern is now the finding — three anti-vacuity defects in one session

- **iter-91** — `platform_alignment_guard` graded *total* resolver failure but not **partial** blindness.
- **iter-93** — the new guard's own live-corpus control **silently skipped** on a hardcoded `parents[3]`.
- **iter-94** — `story_org_count_guard`'s emptiness control **could never fire**.

Three guards, three authors' assumptions, one shape: **the check that would catch "I checked nothing" was
itself the weakest check in the guard** — each written by someone who had just read §5 rule 8.

**The sharper rule: an anti-vacuity control must be written against the SUBJECT of the guard, not against
its inputs.** All three are the same substitution — `roots` existed, a file was found, *some* citations
resolved — while the thing the guard exists to talk about was absent.

## Close — 2026-08-05

**Outcome:** the guard family can no longer report a green one of its members did not earn: the vacuous-pass
path in `story_org_count_guard` is closed with a subject-scoped positive control and a printed cardinality,
and `union_apply_guard` is adjudicated correct-by-design rather than "fixed" into breaking per-stack use.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5, unchanged.** No reading taken.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: **y (5 tiks)** — (6) protocol-stop: n — Outcome: exit-5
**Decisions:** D-M257x-94-1 … D-M257x-94-3
**Side-deliverables:** none.
**Routes carried forward:** all open CHECKs from iters 90–93; `CHECK-M257x-iter91-guard-repo-root-scoping` is CLOSED. **The clause-5 READING is the next iter's whole job** — the instrument has been under repair for four iters and must now be left alone.
**Lessons:**
- **Write the anti-vacuity control against the guard's SUBJECT, not its inputs.** Recorded in the protocol.
- **A shared symptom is not a shared cause** — half of a two-guard finding was correct by design, and
  "fixing" it would have broken per-stack consumption.
