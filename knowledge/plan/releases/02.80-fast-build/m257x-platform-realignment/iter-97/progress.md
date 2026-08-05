# iter-97 — THE RE-READ. N = 20. Clause 5 is NOT met. Gate stays 4 of 5.

**Outcome: the reading was taken**, at platform `0c91421`, corpus `00be1ac` — the tree the iter-96
predicate-wise repair produced. 14 blind seats, two readings of an identical partition, all reports on
disk under `raw/`, all 58 booked blockers adjudicated by four independent graders re-deriving from the
clones.

## The number

**N = 20** distinct in-scope upheld BLOCKER anchors — **17** distinct predicates.
Reading #19 found **12**, reading #20 found **13**, they matched on **5**.

**The gate does NOT move. It stays 4 of 5.**

| quantity | iter-95 | **iter-97** |
|---|---|---|
| booked (14 seats) | 55 | **58** |
| upheld | 51 — **92.7 %** | **54 — 93.1 %** |
| rejected | 4, all ref-discipline | **4, all ref-discipline** |
| **graded N** (in-scope, BLOCKER, deduped) | 13 anchors / 12 predicates | **20 anchors / 17 predicates** |
| Chapman N̂ | ≈ 16.6 | **≈ 29.3** |
| per-pass recall | 60 % / 42 % | **41 % / 44 %** |
| union recall | ≈ 78 % | **≈ 68 %** |
| estimated still unfound | ≈ 4 | **≈ 9** |

## N went UP, and the honest reading of that is not "the repair failed"

51 sites were repaired. The count went **13 → 20**. Three things have to be said in order.

**1. Only 3 of the 20 are recurrences of a repaired predicate.** Seventeen are newly surfaced. So the
repair did not fail to hold — the reading went deeper. iter-95 recorded N̂ ≈ 16.6 with ~4 unfound and
said out loud that N was **a floor twice over**. It was: the true N at that tree was never 13.

**2. The class composition inverted, and that is the repair's signature.** iter-95: 7 of 13 traced to
the `0dab54d → 0c91421` platform move. iter-97: **1 of 20 (5 %)**. The platform-drift class is
essentially cleared. What dominates now is a class iter-95 barely graded — **intra-corpus citations
that resolve to the wrong construct** (3 blockers + the bulk of 60+ minors), **stale currency pins**
(`v1.363.2`), and **enumerations short a member** (12 Connect-RPC services, 20 handlers, 1 roadrunner
mention, 66 range-only citations). The corpus's falsehoods about the *platform* are largely gone; its
falsehoods about *itself* are what is left.

**3. Two of the 20 were caused by the repair, and one of those is the class the method exists to stop.**
`external_services.md:614` — the repair moved a construct 543 → 567, re-pointed the **three cross-file**
citations and missed the **in-file** one, so the corpus now cites two different lines for one construct.
And `dependency_map.md:59` — a cell the repair *wrote* names two `app` refs, which the same commit
forbids twice in its own prose. I stated `D-M257x-96-5` ("a prose repair is a line-number edit, and only
half of that is fenced") and then shipped exactly that defect in the half I had named.

## The three recurrences, which are not three of a kind

- **`service_taxonomy.md` (P4, archive state)** — `D-M257x-96-6` **routed** this class deliberately.
  The reading found the routed item. Routing was right; the scope was too narrow.
- **`backend.md:13` (P5, cms ECS)** — a **third** site, phrased *"still pending"*. The sweep enumerated
  the string `rollback path`. **A paraphrase is not reachable by a string sweep** — iter-93's own rule,
  and it cost a site.
- **`backend.md:33` (P10, skiller stream)** — a genuine miss. The repair fixed `backend.md:127` and left
  `backend.md:33` standing **in the same file**. That is the dominant defect class of this entire
  milestone, committed inside the pass designed to eliminate it.

`claim_twin_guard` is GREEN over all 14 refuted forms while all three of these are live. That is not the
fence lying — it matches **quoted verbatim forms**, and all three are paraphrases. **The fence is
necessary and demonstrably not sufficient**, and now there is a measurement of by how much: 3 of 51.

## The pre-registration — 3 of 7 held, and the 4 failures are the content

| held | failed |
|---|---|
| #4 upheld rate [86,96] → **93.1 %** | #1 per-reading [2,7] → **12 and 13** |
| #5 per-pass recall [30,62] → **41 / 44 %** | #2 union N [3,9] → **20** |
| #6 platform class < 50 % → **5 %** | #3 zero recurrence → **3** |
| #7 repair induces ≥ 2 → **exactly 2** | |

iter-95 graded **6 of 6** and booked it as a *warning, not a win*. The bands were narrowed until they
could fail. They failed on the three predictions that were guesses about magnitude, and held on the four
that were claims about *mechanism* — precision, recall, class composition, and induced-defect rate. That
is a more useful instrument than one that cannot be wrong: **the mechanism model is good and the
magnitude model was badly optimistic.** Predictions 1 and 2 were built on the assumption that repairing
51 sites reduces N. It does not, because N was never a measurement of how many defects exist — it is a
measurement of how many a 68 %-recall instrument surfaces from a corpus that has more.

## What this reading establishes, and what it does not

**Establishes:** the corpus carries **at least 20** blocking falsehoods inside clause 5's own scope at
`00be1ac`, of which ~9 more are estimated unfound; the instrument's precision is stable to within
1.0 points across four adjudications; the platform-drift class M257x was created to fix is down to
**1 of 20**; and predicate-wise repair reaches 51 sites where anchor-wise reaches 13, but leaves
paraphrases standing at a measured rate of 3 in 51.

**Does not establish:** that N is decreasing. Two readings on different trees with different dominant
classes are not a series, and this one is the first at a corpus the milestone itself repaired 51 sites
of. **Comparability: continuous in INSTRUMENT** (briefing byte-identical, sha `3858ec53…`, one commit
ever; 7 seats × 2 readings; same partition method; same grading rule and scope) **and continuous in
UPHELD RATE** (92.1 → 93.0 → 92.7 → **93.1**). **The COUNT is on the same basis as iter-95's** — that
re-baseline was declared there and is not re-declared here — but it is taken on a **different tree**
(`b7e6642` → `00be1ac`, 23 files, +230 −79), and the class composition moved underneath it.

## Routed, NOT repaired

**`FIX-M257x-iter97-read-union`** — the 20 anchors / 17 predicates, plus the named unbooked twins
(the `v1.363.2` pin at 2 more sites; the archive assertion at **10+** in-scope sites inside the very
table that retracts it; the `flow.go:70-95` range at `services/README.md:39`; `studio-room.md:371`), plus
the out-of-scope set: **rule 44's own "1,178 NULs" (the file has 1)**, its shell recipe returning 2 where
its worked example says 22, `platform-alignment.md:1345`, `CLAUDE.md:280`, and `safety.md:203/:207`
(whose Class column now disagrees with `isolation.go:106`, which still registers `s3-private` as
`PerStackIsolated` — the repair asserted a code change that did not happen).

Binding conditions inherited and extended:
1. **Repair by PREDICATE, and enumerate PARAPHRASES, not strings.** 3 of 51 escaped as paraphrases.
2. **Re-derive every inbound citation after any edit that changes a file's line count** — including
   **in-file** `:N` self-anchors, which iter-96 missed while fixing the cross-file ones.
3. **Do not write a measurement into the corpus without running the measuring command as printed.**
