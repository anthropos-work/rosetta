# Adjudication C — seats r33-C, r34-C (M257x iter-135, re-adjudicating iter-131)

**Scope.** Both assigned reports are Seat C reading the same six-file set
(`alignment_testing.md`, `backend.md`, `cms.md`, `clerk-integration.md`, `skillpath.md`,
`gotenberg.md`). r33-C books **7** blockers, r34-C books **5**; they overlap heavily — 5 of r34-C's
5 are re-bookings of r33-C blockers at the same anchors. **12 claimed blockers, 7 distinct anchors.**

I read only: the adjudicator brief, `iter-131/raw/r33-C.md`, `iter-131/raw/r34-C.md`. No other
`knowledge/plan/**` file was opened. I made no edit to any file except this one.

**Corpus ref graded.** The seats were dealt at corpus `a5324934c`
(`probe(M257x/131): the read is SEALED — pre-registration committed before any seat is dealt`),
which r34-C records in its own ground-truth table and which I confirmed is an ancestor of HEAD.
Every corpus quotation below was re-read at `a5324934c` via `git show`, then re-read at HEAD to
establish repair status. **The corpus has since been repaired** (`ceba938 fix/iter(M257x/132)`,
`1012679`, and iters 133/134); 10 of the 12 blockers are `UPHELD (since-repaired)` and **2 are still
live in the tree today.**

**Substrate note — I settled what both seats could not.** Both reports declare the entire
`infrastructure` cluster CANNOT-SETTLE and grade it as *self-contradiction only*. A genuine clone of
`anthropos-work/infrastructure` at exactly `13c248e6` exists on this box outside any stack workspace
(`…/claude-501/…/5de92915-…/scratchpad/infrastructure`; `origin git@github.com:anthropos-work/infrastructure.git`;
`13c248e6` = `ant-workflow[bot]`, *"Release backend v1.371.1 (#3262)"*, 2026-08-07). I read it
directly:

- `git grep 'module "cms' 13c248e6` → **zero hits org-wide**, with a positive control in the same
  pass: `git grep -c 'module "' 13c248e6 -- terraform/production/services.tf` → **12**. The absence
  is real, not a broken command.
- `terraform/production/services.tf:64-70` @ `13c248e6` is **verbatim** what the corpus quotes:
  *"M810: cms was removed (module block deleted here)… Deleting the module destroys its ECS service,
  task definition, ECR repository, IAM roles, security group, Cloud Map entry, log group, alarms and
  the ten /production/cms/* SSM parameters."*
- The declared module set is `acm_media_certificate_useast1`, `backend_euwest1`, `db-backup-euwest1`,
  `directus_euwest1`, `jobsimulation_euwest1`, `metabase_euwest1`, `next-webapp_euwest1`,
  `sentinel_euwest1`, `storage`, `storage-service_euwest1`, `studio_desk_euwest1` — **no `cms`, no
  `messenger`.**

So the cluster is not merely self-contradictory: **the resolved side is TRUE on the merits and the
"NOT MEASURABLE / assert neither" side is FALSE.** That is a stronger grade than either seat reached,
and it does not depend on which of two corpus passages carries more weight.

---

## Verdict table

| seat | B# | anchor (@ `a5324934c`) | verdict | rejection class | predicate (if upheld) | class | multi-pin | repair-induced (sha) |
|---|---|---|---|---|---|---|---|---|
| r33-C | B1 | `corpus/services/clerk-integration.md:115` | **UPHELD** *(still live)* | — | P5 | intra-corpus-citation | yes | **yes** — `e75906b` `iter(M257x/128)` |
| r33-C | B2 | `corpus/services/clerk-integration.md:106-108` | **UPHELD (since-repaired)** | — | P6 | self-contradiction | yes | **yes** — `0c20d8c` `iter(M257x/130)` |
| r33-C | B3 | `corpus/services/backend.md:13` | **UPHELD (since-repaired, in part)** | — | P1 + **P4** | platform-drift + intra-corpus-citation | no | no — `cd16967` `iter(M257x/102)` |
| r33-C | B4 | `corpus/services/backend.md:86-88` | **UPHELD (since-repaired)** | — | P1 | platform-drift | yes | no — `5eb67c7` `iter(M257x/92)` |
| r33-C | B5 | `corpus/services/cms.md:16-17` | **UPHELD (since-repaired)** | — | P1 + P2 | platform-drift + intra-corpus-citation | yes | no — `5eb67c7` `iter(M257x/92)` |
| r33-C | B6 | `corpus/services/cms.md:61` | **UPHELD (since-repaired)** | — | P1 | platform-drift | yes | no — `f8be5a1` `fix(M257x/108)` |
| r33-C | B7 | `corpus/services/cms.md:218` | **UPHELD (since-repaired)** | — | P1 + **P3** | platform-drift + intra-corpus-citation | yes | no — `b4bdbfc` `fix(M257x/115)` |
| r34-C | B1 | `corpus/services/backend.md:13` | **UPHELD (since-repaired, in part)** — dup of r33-B3 | — | P1 + **P4** | platform-drift + intra-corpus-citation | no | no — `cd16967` `iter(M257x/102)` |
| r34-C | B2 | `corpus/services/backend.md:86-90` | **UPHELD (since-repaired)** — dup of r33-B4, **one of three arguments rejected** | — | P1 + **P2** | platform-drift + intra-corpus-citation | yes | no — `5eb67c7` `iter(M257x/92)` |
| r34-C | B3 | `corpus/services/cms.md:14-17` | **UPHELD (since-repaired)** — dup of r33-B5 | — | P1 + P2 | platform-drift + intra-corpus-citation | yes | no — `5eb67c7` `iter(M257x/92)` |
| r34-C | B4 | `corpus/services/cms.md:218` | **UPHELD (since-repaired)** — dup of r33-B7 | — | P1 + P3 | platform-drift + intra-corpus-citation | yes | no — `b4bdbfc` `fix(M257x/115)` |
| r34-C | B5 | `corpus/services/clerk-integration.md:115` | **UPHELD** *(still live)* — dup of r33-B1 | — | P5 | intra-corpus-citation | yes | **yes** — `e75906b` `iter(M257x/128)` |

`git log -L<line>,<line>:<file> --oneline a5324934c` was run on every anchor; the sha in the last
column is the most recent commit touching that exact line at the sealed ref.

---

## Upheld predicates, deduplicated within my assignment

**P1** | *Because `infrastructure` is in no clone set, cms's production ECS teardown (M810) is
unmeasurable and the corpus must assert neither way.* | anchors: `backend.md:13`, `backend.md:86-88`,
`cms.md:16-17`, `cms.md:61`, `cms.md:218` | **platform-drift** *(the only substantive predicate in my
set — falsified directly by `infrastructure` `13c248e6`, read first-hand above)*

**P2** | *`platform-migration-status.md`'s `cms` row states the not-measurable limit and is
authoritative for it.* | anchors: `backend.md:89-90`, `cms.md:17` | **intra-corpus-citation**
*(the map's `cms` row at `:88` states the **resolution** — "RESOLVED at M257x iter-123 — the ECS
service is DESTROYED … was a clone-set limit, not a measurement limit, and the fix was to clone the
repo." Two sentences cite it for the reverse of what it says.) Entailed by P1 — one edit repairs
both; see the coupling note below.*

**P3** | *`cms.md:18` is the sentence stating that `infrastructure` has never been in any clone set.*
| anchor: `cms.md:218` | **intra-corpus-citation** *(`:18` is the merge-order list — "skiller,
skillpath and jobsimulation — **not the** last"; the sentence meant is at `:16`. A self-citation
landing on live prose, invisible to a line-existence check.)*

**P4** | *`backend.md` contains a bullet titled "M810 prod teardown is UNEVEN".* | anchor:
`backend.md:13` | **intra-corpus-citation** *(no such bullet exists. `UNEVEN` occurs exactly twice in
the file — at `:13` in this cross-reference, and at `:77` inside the retraction that **withdraws** the
word. **STILL LIVE AT HEAD** — see below.)*

**P5** | *`corpus/services/ant-academy.md:334` is the `DEV_LOGIN_ENABLED` public-route row.* | anchor:
`clerk-integration.md:115` | **intra-corpus-citation** *(`:334` is the **AI proxy** row,
`| AI proxy | /api/ai/chat | Does its own cookie-based auth() server-side |`. The `DEV_LOGIN_ENABLED`
row is at `:338` — `grep -n DEV_LOGIN_ENABLED` returns a single hit, `:338`. **STILL LIVE AT HEAD.**)*

**P6** | *All three `sign_in_tokens` sites in `corpus/ops/*.md` carry the literal
`curl -s -X POST https://api.clerk.com/v1/sign_in_tokens`.* | anchor: `clerk-integration.md:106-108`
| **self-contradiction** *(I enumerated all three from source at `a5324934c`: `staging-clerk.md:58`
and `staging_from_dump.md:421` carry the literal; `staging-bringup.md:528` carries only
`POST /v1/sign_in_tokens` — no `curl`, no host. The same sentence's preceding clause already carves
that site out by name.)*

**DISTINCT-PREDICATES-IN-MY-SET = 6** — with the coupling caveat that P2 is entailed by P1 and repairs
with it; a coordinator preferring an entailment-collapsed count should read **5**. I report 6 because
the brief's rule is "two different propositions at the same anchor are TWO predicates," and P2 is a
proposition about a *document's content*, not about cms's ECS state.

**Composition warning.** Only **P1** is a claim about the world. P2–P6 are anchor/citation-resolution
defects — five of six. A milestone headline that counts distinct predicates without separating those
two kinds will read this assignment as six factual errors when it contains one factual error and five
broken pointers.

---

## Rejections, with the evidence I opened

**None of the 12 blockers is rejected.** One *argument inside* an upheld blocker is rejected:

**r34-C B2, clause 1** — `misread`. The seat writes: *"'`infrastructure` has never been in one' — the
same bullet, at `:78-79`, states '`infrastructure` @ `13c248e6` declares no `module "cms"`' … **A repo
that was cloned and read at a named sha has been in a clone set.**" That inference is false. I
confirmed independently that `infrastructure` is in no clone set on this box: `find stack-demo -name
.git -maxdepth 4` enumerates **15** trees and none is `infrastructure`; `stack-demo/infrastructure`
and `.agentspace/infrastructure` do not exist. The clone I used is a loose scratchpad checkout from a
different session, not clone-set membership. The corpus itself holds both facts without contradiction
at `platform-migration-status.md:158` — *"the Terraform monorepo … **and it was never in a clone
set.** Reading it at iter-123 settled the `cms` row above."* **The blocker survives on its other two
clauses** (the bullet's own headline at `:76` does assert; the fenced-map citation at `:89-90`
misdescribes its target), so the verdict is UPHELD — but the seat's reasoning for clause 1 does not
stand, and the predicate it implies must not be booked. See Disagreement 1.

---

## Where I disagree with how the seats framed a predicate

**1. The falsified proposition is the INFERENCE, not the premise. (Major — affects the predicate text
itself.)**

Both seats frame this cluster as *"`infrastructure` has never been in a clone set" is false* —
r34-C B2 clause 1 explicitly, r33-C B4 as *"a claim it settles is by definition measurable"* preceded
by *"A repo that can be cited at a sha with a `file:line` was read."* The brief's own worked example
carries the same defect on its right-hand side:
`"cms's production ECS state is unmeasurable / infrastructure was never in a clone set"`.

**The premise is TRUE and must not be booked as refuted.** What is false is the inference chain
*in no clone set ⟹ not measurable ⟹ assert neither*. I verified both halves independently — the repo
is genuinely absent from every clone set here, and it is genuinely readable at `13c248e6`. The
corpus's own subsequent repair reaches exactly this conclusion, in these words
(HEAD `corpus/services/backend.md:91-92`):

> **This bullet said *"NOT MEASURABLE from any clone set we have — do not assert either way"*: the
> premise (`infrastructure` is not in the standing clone set) is TRUE and the inference was FALSE.**
> A repo we do not habitually clone is one `git clone` away, not beyond measurement.

and again at HEAD `cms.md:69`: *"the deciding declaration lives in `infrastructure` — **not in the
standing clone set, and read anyway**."*

**Why this matters beyond wording:** if the predicate is booked as *"infrastructure was never in a
clone set"*, then any future correctly-hedged site saying *"infrastructure is not in the standing
clone set"* scores as a recurrence of a refuted predicate when it is in fact true — the measurement
starts penalising the corpus for being right. The predicate must be written as the inference. I have
stated P1 that way.

**2. `backend.md:13` carries TWO predicates with DIFFERENT repair fates, and both seats fold them into
one. (Actionable — it hides a still-open defect.)**

r33-C B3 and r34-C B1 book the cell's *"NOT MEASURABLE … assert neither"* verdict and its dangling
cross-reference *"see the M810 prod teardown is UNEVEN bullet below"* as one finding. They are
separate propositions, and at HEAD they have diverged:

- the verdict half **was repaired** — HEAD `:13` now reads *"prod teardown **M810 — the ECS service is
  DESTROYED**, measured at `infrastructure` `13c248e6`"*;
- the cross-reference half **is still live**. HEAD `:13` still says *"see the *M810 prod teardown is
  UNEVEN* bullet below"*, and `UNEVEN` still occurs exactly twice in the file — at `:13`, and at
  `:77` inside the retraction that withdraws the word. **No bullet by that name exists.** The repair
  rewrote the verdict and left the pointer.

Folding these into one predicate books the anchor as fixed. It is half fixed. **P4 is open at HEAD.**

**3. `clerk-integration.md:115` is still wrong at HEAD, and neither seat could know that.**

Recorded because it is the other still-open defect in my set. At HEAD the citation has moved to
`:126` and still reads `` [`ant-academy.md:334`](./ant-academy.md) (the `DEV_LOGIN_ENABLED` public-route
pair) ``; `ant-academy.md:334` is still the AI-proxy row and `DEV_LOGIN_ENABLED` is still at `:338`.
**P5 is open at HEAD.** (Its sibling P6 at `:106-108` *was* repaired, and thoroughly — HEAD `:109-114`
now names the two literal sites, quotes the third's divergent form, and replaces the anchor
prescription with *"the robust re-derivation is the shared substring, and only that."* That the corpus
graded and repaired P6 while leaving P5 is itself evidence P6 was correctly booked as a blocker.)

**4. r33-C B7 over-reaches where r34-C B4 is honest, at the same anchor.**

At `cms.md:218` the falsified item is the *stated reason* (*"it lives in `infrastructure`, which has
never been in any clone set"*), not necessarily the *conclusion* (*"no 'still names' claim can be made
here in either direction"* about the production RPC **address**). iter-123 read the module
declaration, not an address. r34-C B4 flags this explicitly (*"the conclusion … may still be sound"*);
r33-C B7 does not, and books both. I uphold the blocker on the reason and on the `:18` self-citation
(P3), and I do **not** book the conclusion as falsified. The repair sided with neither — HEAD `:226`
drops the clause entirely rather than defending it.

**5. Two failure modes are mixed in this assignment and should be counted apart.**

The `git log -L` column separates them cleanly, and the split is not random:

- **Repairs that CREATED defects** — P5 (`e75906b`, `iter(M257x/128)`) and P6 (`0c20d8c`,
  `iter(M257x/130)`). Both repair-induced, both in `clerk-integration.md`, both fresh prose.
- **Repairs that MISSED twins** — the whole P1/P2/P3/P4 cluster. Its anchors were last touched at
  iters **92 / 108 / 115**, all pre-120, so none is repair-induced by the brief's test. But the
  contradiction was *created* at iter-127: `5e43cc5` (`iter(M257x/127)`) rewrote both retraction
  headlines — `backend.md:76` and `cms.md:8`, the same commit — and touched **none** of the four
  restatements at `backend.md:86-90`, `cms.md:16-17`, `cms.md:61`, `cms.md:218`. One commit, two
  twins corrected, four left standing.

A single "repair-induced: yes/no" column tests the anchor line and therefore reports **no** for the
larger and more damaging of the two modes. The M810 cluster is repair-*caused* without being
repair-*induced*; if the milestone is measuring whether repairs are safe, that distinction is the
finding.

---

## Cannot-settle

**None.** Every blocker in my assignment is settled.

The seats' own cannot-settle list (`infrastructure @ 13c248e6`, `services.tf:64-70`, `:85-86`,
`:88-94`) is **now settled** by the direct read documented in the scope section — no `module "cms"`
org-wide (positive control: 12 `module "` hits in the same file), and `services.tf:64-70` verbatim as
cited. Their remaining cannot-settles (colony internals, prod-DB row counts, historical alignment
scores, upstream gotenberg `/health`) sit under *Positively cleared* / *What I could not settle*, not
under any blocker, so they are out of my adjudication scope; I record that I agree those four are
genuinely unsettleable from this box and that neither seat laundered any of them into a blocker.

---

## Counts

```
CLAIMED=12
UPHELD=12   (of which: still-live-at-HEAD=2, since-repaired=10)
REJECTED=0  (wrong-tree=0, misread=0, true-at-its-ref=0,
             retraction-not-contradiction=0, minor-not-blocker=0, not-in-scope=0)
CANNOT-SETTLE=0
WRONG-TREE=0
DISTINCT-PREDICATES-IN-MY-SET=6   (5 if P2 is collapsed into P1 as entailed)
UPHELD-RATE=12/12 = 100%
DEDUP FACTOR: 12 claimed blockers -> 7 distinct anchors -> 6 distinct predicates -> 1 substantive
```

**On `retraction-not-contradiction`, the class the brief names as the most common over-booking.** I
tested every P1 site against it and it does not apply. The corpus does say *"X was wrong; the truth is
Y"* — at `backend.md:76` and `cms.md:8`. But the four sites the seats book are **not** the retraction:
they are unmarked restatements of X standing elsewhere in the same documents, in the present tense,
carrying bolded imperatives (*"Do not assert either way"*), and forwarding to two corroborating
locations that both say the opposite. A reader reaching `cms.md:17` has no signal that the sentence is
withdrawn. That is a half-applied edit, not prose doing its job — and the repair at iter-132 agrees,
because it went back and marked those four passages as retracted rather than defending them.
