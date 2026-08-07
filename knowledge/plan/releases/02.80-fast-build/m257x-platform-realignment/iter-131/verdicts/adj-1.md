# Adjudication adj-1 — seats r33-A, r34-A

**Scope:** 20 booked blockers (8 from r33-A, 12 from r34-A) over
`corpus/architecture/external_services.md`, `corpus/architecture/platform-migration-status.md`,
`corpus/services/storage.md`, `corpus/services/askengine.md`, `corpus/services/db-backup.md`
(+ one r33-A anchor in `corpus/architecture/architecture_overview.md`).

**Everything below was opened by me.** No verdict rests on a seat's citation, a seat's confidence, or
a seat's arithmetic. Where I could not open the substrate I say so and do not launder it.

---

## The `infrastructure` question, settled first (it governs 6 of the 20 blockers)

The task asked me to settle, from the actual clone set and the actual cited evidence, whether
`cms`'s production ECS state is measurable and whether `infrastructure` "has ever been in a clone set".
Three independent measurements:

1. **`infrastructure` is in NO clone set on this box, now.** `find . -maxdepth 5 -name .git` returns
   **20** trees: the rosetta repo itself, `.agentspace/rosetta-extensions`, 15 under `stack-demo/`
   (including the two nested `studio` checkouts and rext), and 3 under `stack-dev/`
   (`studio-desk`, `studio-room`, one worktree). Neither `infrastructure` nor `db-backup` is among
   them. `ls stack-demo` and `ls stack-dev` confirm it directly.
2. **The fence agrees, out of its own mouth.** Running `platform_alignment_guard.py` from the
   **authoring** copy `f2ea567b3` (rule 50) prints
   `UNCLONABLE head 'infrastructure' x9` and `UNCLONABLE head 'db-backup' x1`, and returns
   **`OK OVER ITS REACH`** — explicitly declaring those 10 citations *not checked*.
3. **The corpus does not claim `infrastructure` is in a clone set — it claims it was READ once,
   transiently.** `platform-migration-status.md:158` says it in as many words: *"the Terraform
   monorepo … **and it was never in a clone set**. Reading it at iter-123 settled the `cms` row
   above."* `org-repos.md:13` says the same (*"hinge on `infrastructure`, which is not in any stack's
   clone set"*), and `org-repos.md` §3 (`:87-179`) carries a detailed derivation at
   `infrastructure` **`13c248e6`** with ~10 distinct line ranges, a per-repo table, and a
   name-collision warning at `services.tf:122-124`.

**So "in no clone set" and "has never been read" are different propositions, and the corpus
distinguishes them at `:158` while six other passages conflate them.** I cannot verify the *content*
of the iter-123 measurement — `infrastructure` is unreadable from here, and I say so under
Cannot-settle. But that is not what the blockers turn on. What they turn on is settleable and I
settle it: the same document says both *"we cloned it, read it, and the question is SETTLED"*
(`:88`, `:90`'s MEASURED-AT-LAST paragraph, `:158`) and *"a repo this map has never read" / "in no
clone set" / "unmeasurable is the state"* (`:89`, `:90`'s closing clause, `:310`, `:315-316`), with
two further files publishing the second reading (`external_services.md:175`,
`storage.md:175`). Those cannot both stand, and the weight is overwhelmingly on the read having
happened: iter-123 repaired three sibling rows on that measurement (`cms`, `messenger`,
`roadrunner`), iter-124 repaired a fourth (`graphql-wundergraph`), the census row carries a
2026-08-07 push date, and `external_services.md:3`'s banner quotes `services.tf:509-517` verbatim.
**`UNMEASURABLE` is the retracted reading; the sites still publishing it are stale.**

I verified the clone-visible halves of the cms evidence myself: `cms/terraform/main.tf:39` is
`service_desired_count = 0` at **both** `ca50c8170` and `f38c0c4a4`, and `6efa1d5` is
*"chore(ci): drop build-production — the cms ECR repository is decommissioned (M810)"*, 2026-08-04.
Both cited facts are real. It is the conclusion drawn from them that has been superseded.

---

## Verdict table

| seat | B# | anchor | verdict | rejection class | predicate (if upheld) | class | multi-pin | repair-induced (sha) |
|---|---|---|---|---|---|---|---|---|
| r33-A | B1 | `corpus/architecture/platform-migration-status.md:189` | **UPHELD** | — | P5 | self-contradiction | no | **yes** — `0c20d8c` `iter(M257x/130)` |
| r33-A | B2 | `corpus/architecture/platform-migration-status.md:89` (+`:310`) | **UPHELD** | — | P2 | self-contradiction | yes | no — `cd16967` `iter(M257x/102)` |
| r33-A | B3 | `corpus/architecture/external_services.md:175` | **UPHELD** | — | P2 | self-contradiction | yes | no — `4ea8c6c` `fix(M257x/92)` |
| r33-A | B4 | `corpus/services/storage.md:175` | **UPHELD** | — | P2 | self-contradiction | no | no — `4ea8c6c` `fix(M257x/92)` |
| r33-A | B5 | `corpus/architecture/external_services.md:368` (+`:14`, `:804-807`) | **UPHELD** | — | P1 | self-contradiction | yes | no — `a5126bc` `doc(M257x/21)` |
| r33-A | B6 | `corpus/services/askengine.md:81` (+`:113`) | **UPHELD** | — | P6 | platform-drift | no | no — `fb16e9a` `build(M247)` |
| r33-A | B7 | `corpus/architecture/external_services.md:727` | **REJECTED** | `minor-not-blocker` | — | — | — | — |
| r33-A | B8 | `corpus/architecture/architecture_overview.md:35` | **UPHELD** | — | P4 | self-contradiction | no | no — `ca8e381` `docs: sweep corpus drift` |
| r34-A | B1 | `corpus/architecture/platform-migration-status.md:96` | **UPHELD** | — | P1 | self-contradiction | yes | **yes** — `37d256f` `fix(M257x/126)` |
| r34-A | B2 | `corpus/architecture/platform-migration-status.md:90` (state cell) | **UPHELD** | — | P3 | self-contradiction | yes | **yes** — `3cd96f2` `iter(M257x/123)` |
| r34-A | B3 | `corpus/architecture/platform-migration-status.md:189` | **UPHELD** | — | P5 | self-contradiction | no | **yes** — `0c20d8c` `iter(M257x/130)` |
| r34-A | B4 | `corpus/architecture/platform-migration-status.md:310` | **UPHELD** | — | P2 | self-contradiction | yes | no — `cd16967` `iter(M257x/102)` |
| r34-A | B5 | `corpus/architecture/platform-migration-status.md:315-316` | **UPHELD** | — | P2 | self-contradiction | no | no — `904502c` `iter(M257x/87)` |
| r34-A | B6 | `corpus/architecture/platform-migration-status.md:90` ("never read" clause) | **UPHELD** | — | P2 | self-contradiction | yes | **yes** — `3cd96f2` `iter(M257x/123)` |
| r34-A | B7 | `corpus/architecture/platform-migration-status.md:89` | **UPHELD** | — | P2 | self-contradiction | yes | no — `cd16967` `iter(M257x/102)` |
| r34-A | B8 | `corpus/architecture/external_services.md:14` | **UPHELD** | — | P1 | self-contradiction | no | no — `a5126bc` `doc(M257x/21)` |
| r34-A | B9 | `corpus/architecture/external_services.md:368` | **UPHELD** | — | P1 | self-contradiction | yes | no — `a5126bc` `doc(M257x/21)` |
| r34-A | B10 | `corpus/architecture/external_services.md:175` | **UPHELD** | — | P2 | self-contradiction | yes | no — `4ea8c6c` `fix(M257x/92)` |
| r34-A | B11 | `corpus/services/storage.md:175` | **UPHELD** | — | P2 | self-contradiction | no | no — `4ea8c6c` `fix(M257x/92)` |
| r34-A | B12 | `corpus/architecture/platform-migration-status.md:102` | **UPHELD** | — | P4 | self-contradiction | yes | **yes** — `3cd96f2` `iter(M257x/123)` |

*(Repair-induced computed with `git log -L<line>,<line>:<file> --oneline | head -3` on each anchor.
Note for re-runners: **`git log -L$L,$L:path` silently mangles in zsh** — `$L:c` is eaten as a
parameter-expansion modifier and git fatals with `-L argument not 'start,end:file'`. It must be
`-L${L},${L}:path`. My first pass returned seven empty results from this; the stderr was the tell.)*

---

## Upheld predicates, deduplicated within my assignment

```
P1 | The Cosmo/WunderGraph router still runs in production ("prod-only" / prod state `live-standalone`)
   | anchors: corpus/architecture/platform-migration-status.md:96 ·
              corpus/architecture/external_services.md:14 ·
              corpus/architecture/external_services.md:368 ·
              corpus/architecture/external_services.md:804-807
   | class: self-contradiction

P2 | `infrastructure` is in no clone set / has never been read, so the folded services' production
     terraform disposition is UNMEASURABLE — the `cms` row "reports both and asserts neither", and the
     `roadrunner` disagreement cannot be settled
   | anchors: corpus/architecture/platform-migration-status.md:89 ·
              corpus/architecture/platform-migration-status.md:90 (the "a repo this map has never read" clause) ·
              corpus/architecture/platform-migration-status.md:310 ·
              corpus/architecture/platform-migration-status.md:315-316 ·
              corpus/architecture/external_services.md:175 ·
              corpus/services/storage.md:175
   | class: self-contradiction

P3 | `roadrunner` still runs as its own process on the production traffic path (prod state `live-standalone`)
   | anchors: corpus/architecture/platform-migration-status.md:90 (the prod STATE cell)
   | class: self-contradiction

P4 | `db-backup` is live in production — scheduled PostgreSQL backups fire
   | anchors: corpus/architecture/platform-migration-status.md:102 (the prod STATE cell) ·
              corpus/architecture/architecture_overview.md:35
   | class: self-contradiction

P5 | §1's state vocabulary has NINE members — i.e. `library-unimported` is a state defined in §1
   | anchors: corpus/architecture/platform-migration-status.md:189
              (the live undefined token is at :109)
   | class: self-contradiction (with an arithmetic/count surface)

P6 | The shared `ai` library is imported as a private Go module by something a stack builds —
     specifically, the Ask Engine depends on it downstream
   | anchors: corpus/services/askengine.md:81 · corpus/services/askengine.md:113
   | class: platform-drift
```

### How I drew the predicate boundaries, deliberately

- **P1 absorbs four booked blockers across two files and three constructs** (a fenced table's state
  cell, a PM summary bullet, a spec-table row, and a production-deployment instruction). All four
  assert one proposition — *the router is still deployed and serving in production*. The brief offers
  this predicate almost verbatim, so I did not split it by construct.
- **P2 absorbs six.** I considered splitting it into a cms half and a roadrunner half. I did not,
  because every one of the six sites states the *same literal proposition* — "in no clone set" /
  "never in a clone set" / "a repo this map has never read" / "cannot settle without reading
  **infrastructure**" — and every one of the six is repaired by the *same* fact (the iter-123 read).
  The brief's own worked example fuses the cms conclusion and the infrastructure premise into one
  predicate string, which is the reading I followed.
- **P2 and P3 are two predicates at ONE anchor** (`:90`). This is the brief's "two different
  propositions at the same anchor are TWO predicates" case, and it is a real one: the roadrunner cell
  contains a *positive* false claim (its prod state cell says the service is live) and, twenty lines
  later, a *negative* false claim (that infrastructure has never been read). Repairing either leaves
  the other standing. r34-A booked them separately and was right to.
- **P4 merges r33-A B8 and r34-A B12** even though one says "scheduled" and the other says
  `live-standalone`. §1 `:49` defines `live-standalone` as *"its own process, still on the traffic
  path"*, so both assert the same thing: db-backup is running. Nothing fires it.
- **P5 and P6 stand alone.** No other blocker in my set touches either.

---

## Rejections, with the evidence I opened

**r33-A B7 — `external_services.md:727`, the `anthropos-agent-eu` absence measurement
(*"0 hits across all 15 trees at their own refs, and 0 in a `.gitignore`-blind filesystem grep"*).
REJECTED, class `minor-not-blocker`.**

I re-ran both instruments rather than taking either side's word. `find stack-demo -maxdepth 4 -name
.git` returns exactly **15** trees, so the seat's denominator is right. Grepping each at its own
`HEAD` (`git grep -c -F 'anthropos-agent-eu' <ref>`) returns **0 in fourteen of them** and **7 files
in the fifteenth**, `stack-demo/rosetta-extensions` @ `09d06070f`. A `.gitignore`-blind filesystem
grep of `stack-demo` returns the same **7**. Positive control in the same pass: `anthropos-agent-us`
returns 11 files, and `anthropos-agent` fires in `app` and `jobsimulation`, so the pipeline is live.
**So the seat's arithmetic is exactly right and I reproduce it.**

I reject it anyway, because the predicate the parenthetical exists to support is TRUE and the seat
itself says so. I opened all seven hits:
`stack-core/tests/fixtures/claim_twin_iter48/red/07.md:3`,
`stack-core/tests/fixtures/mechanical/{green,red}/corpus/architecture/external_services.md:652`, and
`stack-core/tests/fixtures/repair_leak/{pre,post}/corpus/architecture/{ai_architecture.md,external_services.md}`.
Every one is a **verbatim archived copy of this corpus's own retracted prose**, checked into the
*tooling* monorepo as a regression fixture. `rosetta-extensions` is not the platform — the root
`CLAUDE.md` draws that line explicitly ("`rosetta` documents *how the platform works*;
`rosetta-extensions` is *the tooling*"), and the sentence under audit says *"the name appears nowhere
in **the platform**"*. It appears nowhere in the platform: 0 hits in all fourteen platform trees,
which I enumerated. The defect is a stated denominator that over-reaches by one non-platform tree,
and its entire failure mode is that a re-runner meets seven files that are literally this sentence's
own ancestor. That is a re-run annoyance worth repairing; it cannot mislead a reader doing real work
about what agent names the platform dispatches. The seat booked it at **medium** confidence and wrote
*"the substantive claim is TRUE"* in its own body; r34-A read the same file and cleared the same
claim. `minor-not-blocker`, and I would route it to the minors list rather than drop it.

**No `wrong-tree` rejections.** Both seats stated which rext tree settled what and both got it right.
I re-checked the one that matters: `ALLOWED_STATES` really does carry **nine** members including
`library-unimported` in the authoring copy `f2ea567b3` (`:89-99`) and **eight** in the pinned
per-stack clone `09d06070f`. Neither seat graded a tooling claim against the wrong tree.

---

## Notes on three upholds that deserved more than a nod

**P5 is stronger than the off-by-one it looks like, and I nearly minored it.** §1's vocabulary table
(`:47-56`) has exactly **eight** data rows (`sed -n '49,56p' | grep -c '^|'` → 8), and `:189` says
*"every state is one of the **nine** in §1"*. The token actually in use, `library-unimported`, occurs
in the file **once** — at `:109`, in both of the `authn` row's state cells — and is defined nowhere
in §1. I then established that the guard does not read §1 at all: `ALLOWED_STATES` is a hard-coded
literal consumed at `platform_alignment_guard.py:713`. So a green fence is not evidence for the
"nine". Then I ran both fences against the corpus:

- authoring `f2ea567b3` → **`OK OVER ITS REACH`**, exit 0.
- pinned/shipped `09d06070f` → **exit 1, two findings**:
  `[C vocabulary] authn: prod state 'library-unimported' is not one of [...eight...]` and the same for
  the local cell.

**Three live vocabularies — the doc's §1 (8), the shipped fence (8), the authoring fence (9) — and a
§4 row asserting a fourth thing about §1.** Any stack running the pinned guard against this corpus
goes RED today. That is not cosmetic.

**P4's `db-backup` half (r34-A B12) was the closest call in my set.** The seat flagged its own
hesitation honestly: unlike the router and roadrunner, db-backup genuinely *is* deployed — task
definition, ECR repo, IAM roles, log group and S3 bucket all survive — and §1 offers no clean token
for "deployed but unfired". I upheld it because §1 `:49` does not say "deployed"; it says *"its own
process, still on the traffic path"*, and nothing fires this one. The row's own evidence cell says
**"Deployed but not triggered"** four words later, `db-backup.md:3` opens *"the schedule has been OFF
for over a year"*, and `db-backup.md:53` reads **"Schedule | none, currently"**. A reader scanning
the state column of the file the corpus designates its fenced index of truth gets exactly the belief
seven corpus files were repaired to kill. It reads as a vocabulary gap *and* a false cell; the cell
is false either way.

**P4's `architecture_overview.md:35` half is in scope even though it is outside r33-A's assigned
file set.** The brief's `not-in-scope` class is defined by directory — outside
`corpus/services/**` + `corpus/architecture/**` — and `architecture_overview.md` is in
`corpus/architecture/**`. The seat flagged the out-of-assignment status itself and was right to book
it: `:35` reads *"Production-only: **db-backup** (scheduled PostgreSQL backups)"* with no qualifier,
while **the same file's `:206`** carries *"Trigger commented out since `7dd1b80`, 2025-05-29 …
deployed, unfired"* and **`:454-459`** carries a bolded retraction naming the exact prior wording.
One file, two answers, 170 lines apart.

---

## Cannot-settle

**None of the 20 blockers.** Every one was decidable on evidence I opened. But three things my
verdicts *rest beside* are not checkable from this box, and I record them so nothing here is read as
a clearance:

1. **The CONTENT of the iter-123 `infrastructure` measurement.** I cannot confirm that
   `module "cms"` is absent, that `services.tf:64-70` / `:85-86` / `:88-94` / `:509-517` / `:521` /
   `:571` / `:622` / `:664` read as quoted, or that `13c248e6` is what the corpus says it is.
   **P2 does not depend on this** — P2 books an internal contradiction that stands whichever way the
   repo reads. **Settled by:** `git clone anthropos-work/infrastructure`, checkout `13c248e6`, and
   run `grep -rn 'module "cms' terraform/ modules/` plus the eight cited ranges.
2. **Everything in `corpus/services/db-backup.md`** that cites the `db-backup` repo — the
   `6e1fb15b`/`v0.3.3`/HEAD identity, the 43-line Bash census, the all-157-objects Azure sweep, the
   `rate(12 hours)` value, `terraform/main.tf:10-27`/`:51`/`:215`, `storage.tf:24-38`, and the
   0-line `terraform/stage/main.tf`. The repo is in no clone set (the guard says
   `UNCLONABLE head 'db-backup' x1`). The page is internally coherent and its retraction has
   propagated; I found **no live contradiction to book inside it**, which is why it contributes no
   blocker of its own. **Settled by:** cloning `anthropos-work/db-backup` at `6e1fb15b`.
3. **Whether `:88`'s verdict (cms DESTROYED) is itself correct.** I uphold that the *stale* sites are
   wrong; I do not certify `:88`. If `:88` turned out to be the wrong half, P2's six sites would
   still be a defect — they would simply need repairing in the other direction, since a document
   cannot say both.

---

## Counts

```
UPHELD=19 REJECTED=1 (of which wrong-tree=0) CANNOT-SETTLE=0
DISTINCT-PREDICATES-IN-MY-SET=6
```

Per seat: r33-A **7 upheld / 1 rejected**; r34-A **12 upheld / 0 rejected**.
Overlap collapsed: 20 booked blockers → **6 distinct predicates** (a 3.3× dedup).
Repair-induced: **6 of the 19** upheld blockers sit on lines last touched by iters 120–130 —
`3cd96f2` (iter-123) ×3, `0c20d8c` (iter-130) ×2, `37d256f` (iter-126) ×1 — landing on **4 distinct
anchors**, all in `platform-migration-status.md`: `:90`, `:96`, `:102`, `:189`. Every one is the same
shape: a repair reached one construct of a row (or one row of a file) and left its sibling standing.
All 8 anchors outside that file are pre-120 and are drift the recent repair wave never touched.
