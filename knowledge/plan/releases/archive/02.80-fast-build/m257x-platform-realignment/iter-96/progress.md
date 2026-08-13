# iter-96 — the REPAIR, by predicate. 13 anchors → 51 sites.

**Shape:** repair pass. No reading happens inside this iter; iter-97 takes the re-read against the
tree this iter produces. The separation is iter-95's binding condition and is what made `140 → 43`
mean anything.

## The number that matters

| quantity | value |
|---|---|
| anchors iter-95 booked | **13** |
| distinct predicates | **12** (+2 instrument-control defects) |
| **sites repaired** | **51** across 23 files |
| sites an anchor-wise repair would have left standing | **38** |
| adjudicators' estimate of unbooked twins | **≥8** |

**The estimate was low by more than 4×.** The largest single predicate — `mistralai` "imported
nowhere" — stood at **11** sites, of which iter-95 booked exactly **one**. Repairing the 13 anchors and
stopping would have left the corpus asserting the same twelve falsehoods in 38 other places, which is
precisely how P4 survived the repair it was repaired in.

Full per-predicate table with refuted forms and truth: [`blocker-ledger.md`](blocker-ledger.md).

## The repair is FENCED, not just written

`blocker-ledger.md` is deliberately written in the table shape `stack-core/claim_ledger.py` derives
from. The consequence is mechanical and is the whole point of TOK-05:

- `claim_twin_guard`'s adjudicated-claim set grew **101 → 114**, and it is **GREEN** — a machine-checked
  statement that **none of the 14 refuted forms is published anywhere in the tree**. That is a
  completeness proof of a predicate-wise repair, not a promise that one was done.
- `repair_postcondition` adopts the same set at the commit, so a future edit that re-publishes any of
  them cannot be committed through.
- `unreadable_repo_claim_guard`'s reach grew **5 → 7** `module.*_euwest1` mentions: the P5 repair spells
  the dotted construct, so two sites that had been paraphrasing *through* the unmeasurability boundary
  are now inside the fence instead of invisible to it. iter-93's own warning — *fencing a document does
  not fence its paraphrases* — was the defect, and this closes two instances of it.

## The bare-`grep` absence class, sized

The standing rule inherited from iter-95 was *"an absence is established only by `git grep` at a named
ref."* **Measured, that rule is necessary and not sufficient — and on the predicate that produced it,
the rule's own instrument scored WORSE than the one it replaced.**

Three mechanisms, measured across **15 git trees** (12 clones + corpus + rext + one nested repo — later
corrected to **two** nested repos):

| mechanism | population | which tool is blind |
|---|---|---|
| 1 — `.gitignore` hides **tracked** files | **12** non-empty text files | bare `grep` only |
| 2 — NUL-bearing source | **2** files (`useCoursebuilder.ts` 50,433 B / 1,178 NULs; `store.js`) | **both** tools |
| 3 — nested untracked repos | **2** (`app/studio`, `cms/studio` — same repo, 177 files, own HEAD `aeec036`) | **`git grep` at the host ref**, absolutely |

On the `mistralai` predicate the three instruments returned **1 / 0 / 22**. The **`0` is the ref-named
`git grep`** — `git -C app grep <anything> HEAD -- studio/` returns zero for *every* predicate, because
`app/.gitignore:79` is `studio/*` and the path is in `app`'s index at no sha. A guaranteed zero that
reads like evidence is worse than a noisy one.

**Claim side, sized:** 30 distinct instrument-derived absence-claims in `corpus/services/**` +
`corpus/architecture/**`; **13** exposed to one of the three mechanisms; re-derived with the correct
per-tree instrument, **1 flipped** (the `mistralai` one, already repaired), **11 held**, **1 is
undecidable in this workspace** (a claim ranging over 93 org repos of which 13 are cloned — flagged as
scope-unbounded, not repaired).

**So the class is small in claims and large in consequence**, and the honest disposition was to fix the
*rule*, not sweep the *claims*:

- `platform-alignment.md` §5 **rule 44** — the amended rule, with the three mechanisms, the measured
  populations, and a copy-pasteable enumerate-then-grep recipe.
- `anchor_construct_guard._clone_of` now **descends to the innermost git checkout**, so a citation into
  `app/studio/**` is read at `aeec036` rather than reported UNMEASURED at `app`'s ref. The guard caught
  this on my own repair before I did.

### The measurement demonstrated its own defect

My first census counted **4** gitignored-but-tracked text files and **1** nested repo. Both were wrong,
and wrong by the mechanism under study: `git check-ignore` without `--no-index` does not report tracked
paths, and the census did not descend into nested repos, so it could not see `cms/studio` at all. The
corrected figures are **12** and **2**. That is recorded in rule 44 rather than quietly fixed.

## Two guards caught defects this repair introduced

Worth naming, because it is the third layer working:

1. **`anchor_construct_guard` went CANNOT-RUN twice.** Once because my `app/studio/tools/pdf2md.py:24`
   citations were being graded at `app`'s ref (mechanism 3, in the instrument) — fixed in the guard.
   Once because my `storage.md:29` edit named a **second ref** inside a cell whose every `app` path
   resolves at `2035f9a` only; `app/internal/storage/service.go` does not exist at `b948604` at all.
   That is M257x run-53's one-ref-per-block rule, violated inside the repair for it — and rewritten to
   carry the contrast without the second sha.
2. **The line-anchor shift.** Growing `external_services.md` by 24 lines moved **13 citations** in 6
   other files (`:543 → :567`, `:555 → :579`, `:577-582 → :602-607`, `:594-604 → :619-629`). Six were
   caught as `anchor-on-blank-line`; the rest were found by diffing old→new line maps, because an anchor
   that lands on a *non-blank wrong* line is silent. **A prose repair is a line-number edit**, and the
   blank-line check only catches the lucky half.

## What is NOT in this iter, deliberately

- **No reading.** iter-97 measures.
- **No platform-side anything.** `DEF-M257x-iter80-storage-prod-bucket` is documented as a hazard, with
  the footgun named and cited on both sides, and its disposition left explicitly open. The corpus now
  says a stock stack writes private uploads to a production bucket; it prescribes no fix.
- **The archive-state class beyond jobsimulation.** The predicate is repaired where there is contrary
  evidence (jobsimulation, 6 sites). `skiller` / `skillpath` / `chronos` / `intelligence` /
  `graphql-wundergraph` carry the identical epistemic status with **no clone to measure and no contrary
  evidence** — routed as `FIX-M257x-iter96-archive-class` (15+ sites), not swept blind.
