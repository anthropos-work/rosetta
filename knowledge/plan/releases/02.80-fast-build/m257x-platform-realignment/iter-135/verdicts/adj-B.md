# Adjudication B — seats `r33-B`, `r34-B` (M257x iter-131, re-adjudicated at iter-135)

**Scope line.** Independent adjudication of the **4 claimed BLOCKERS** in the two seat reports
`iter-131/raw/r33-B.md` (B1, B2) and `iter-131/raw/r34-B.md` (B1, B2). I read **only** the
adjudicator brief and those two seat reports from under `knowledge/plan/**`; no other iter dir, no
`progress.md` / `decisions.md` / `adjudication.md`, no other adjudicator's output, nothing from
iter-132/133/134. Every citation below was opened by me at the ref the claim names (brief §"A claim
is settled at the ref"), with a positive control in the same pass. Read-only apart from this file.
Corpus graded at `rosetta` HEAD `0dd19e5`; **none of the three anchors has been repaired** — all
three still carry the booked text live (histories below).

## Counts table

| metric | n |
|---|---|
| BLOCKERS claimed | **4** |
| UPHELD | **4** |
| REJECTED — `wrong-tree` | 0 |
| REJECTED — `misread` | 0 |
| REJECTED — `true-at-its-ref` | 0 |
| REJECTED — `retraction-not-contradiction` | 0 |
| REJECTED — `minor-not-blocker` | 0 |
| REJECTED — `not-in-scope` | 0 |
| CANNOT-SETTLE | 0 |
| **wrong-tree count (all classes)** | **0** |
| DISTINCT PREDICATES in my set | **3** |
| upheld rate | **4/4 = 100 %** |
| UPHELD *(since-repaired)* | 0 — all three anchors are live-unrepaired |

## Verdict table

| seat | B# | anchor | verdict | rejection class | predicate (if upheld) | class | multi-pin | repair-induced (sha) |
|---|---|---|---|---|---|---|---|---|
| r33-B | B1 | `corpus/services/sentinel.md:5` | **UPHELD** | — | **P-1** | arithmetic/count | yes | **no** — last touched `d18aee2` `fix(M257x/115)`; iter-115 is outside the 120–130 window |
| r34-B | B1 | `corpus/services/sentinel.md:5` | **UPHELD** | — | **P-1** (same anchor, same proposition — dedups with r33-B1) | arithmetic/count | yes | **no** — `d18aee2` (iter-115) |
| r33-B | B2 | `corpus/architecture/dependency_map.md:9` | **UPHELD** | — | **P-2** | self-contradiction (+ platform-drift) | no | **no** — last touched `904502c` `iter(M257x/87)` |
| r34-B | B2 | `corpus/services/ai-readiness.md:18-20` | **UPHELD** | — | **P-3** | platform-drift | yes (light: the M247 block pins `4c28365f` + a rename table of paths) | **no** — last touched `46cc66a` `build(M247) §7` |

## Upheld predicates, deduplicated within my assignment

**P-1** | *`messenger` @ `fa47850d` contains at least one `authorization` / `AUTHORIZATION_ADDRESS`
occurrence across `*.go` + `go.mod` — i.e. the corpus's published re-derivation returns **one**
unrelated hit* | anchors: `corpus/services/sentinel.md:5` | class: **arithmetic/count**
→ **FALSIFIED: it returns ZERO**, and the substring has never existed in that repo at any ref.

**P-2** | *`roadrunner` is one of the domains folded into `app` — an `app`-internal domain whose
edges collapsed into in-process calls* | anchors: `corpus/architecture/dependency_map.md:9` |
class: **self-contradiction** (against `corpus/services/README.md:20-24` + `:39` and
`corpus/architecture/platform-migration-status.md:88`), corroborated by **platform-drift**
→ **FALSIFIED: `app/internal/roadrunner/` exists at no ref and was never added, ever.**

**P-3** | *`WorkforceDirectory` is a **two-method member-directory** seam
(`LoadMembers`/`LoadMembersByUserIDs`) whose implementations are **entirely in**
`app/internal/workforce/members.go`, and the member directory is `aireadiness`'s **only** remaining
dependency on `workforce`* | anchors: `corpus/services/ai-readiness.md:18-20` | class:
**platform-drift** → **FALSIFIED: the interface declares FOUR methods; `LevelsCount` is the org's
skill-scale setting (not a member), and it is implemented at `internal/workforce/manager.go:90`.**

---

## The evidence I opened, per blocker

### P-1 — `sentinel.md:5`, the grep receipt (r33-B1 **and** r34-B1)

Corpus text, live at HEAD, unrepaired:

> *"What survives, and was re-derived: messenger's Go source imports no authorization client
> (`git grep "authorization\|AUTHORIZATION_ADDRESS" fa47850d -- '*.go' go.mod` returns **one
> unrelated hit**, against `colony` present as a positive control)"*

Run in `stack-demo/messenger` (`origin git@github.com:anthropos-work/messenger.git`, working tree
clean, `HEAD = fa47850d9c507d1928da7a38f7b37bac1bb8fabc`), literally as published:

```
git grep "authorization\|AUTHORIZATION_ADDRESS" fa47850d -- '*.go' go.mod   → exit 1, 0 lines
git grep -E "authorization|AUTHORIZATION_ADDRESS" fa47850d -- '*.go' go.mod → exit 1, 0 lines
git grep -in "authoriz" fa47850d          (whole tree, no pathspec)          → exit 1, 0 lines
```

**Positive control, identical invocation form** — `git grep -n "colony" fa47850d -- '*.go' go.mod`
→ hits in **9** files, first at `cmd/root.go:11`. The ref, the pathspec and the pipeline all
resolve. The nearest near-miss is `go.mod:67` `golang.org/x/oauth2 v0.36.0 // indirect`, which the
published pattern does not match.

**I added one check neither seat ran, and it is decisive.** This is not a stale measurement that
expired — the figure was **never** obtainable from this repo:

```
git log --all --oneline -S 'authoriz' -i   → 0 commits (positive control -S 'colony' → 35)
```

The substring `authoriz` has never appeared in `messenger` at **any** commit in its entire history.
So `true-at-its-ref` is unavailable as a rejection class in the strongest possible sense: there is no
ref at which the receipt was ever true.

**Why not `minor-not-blocker`** (the only class that had a real case, and both seats flagged the
same hesitation — see my framing note below). The carve-out is for cosmetic defects: line drift, an
omitted list member, a bare basename. This is none of those. It is a **false empirical assertion
about the platform**, published inside a clause explicitly framed *"was re-derived"*, and it is
**not self-correcting** — nothing in the passage tells the reader the true answer is zero, so a
reader who re-runs it lands in the "is my pipeline broken or is the corpus stale?" trap with no
in-corpus resolution. Contrast the same file's `PORT` default-column error (booked MINOR by r33-B),
which the adjacent description corrects inside the same cell.

**Dedup:** r33-B1 and r34-B1 are the **same anchor** and the **same proposition** → one predicate.
Their agreement corroborates the *fact*, not the *grading*; I graded the grading myself.

### P-2 — `dependency_map.md:9`, `roadrunner` as a folded domain (r33-B2)

Corpus text, live at HEAD, unrepaired (line 9 verified by `awk 'NR==9'`):

> *"Since the monolith merge most of this matrix collapsed: `skiller`, `skillpath`, **`roadrunner`**,
> `jobsimulation`, `cms`, `storage`, `messenger` and `customerio-sync` are all domains inside
> **Backend (`app`)**, so their edges are in-process calls, not dependencies."*

**Ground truth at `app` `ad9f3c498`** (the checkout — the banner names no ref):

- `git ls-tree -r --name-only ad9f3c49 | grep -i roadrunner` → **exit 1, no path in the whole tree**.
  Positive control `jobsimwiring` → **3** paths.
- `git log --all --diff-filter=A -- 'internal/roadrunner' 'internal/roadrunner/*'` → **0 commits,
  ever**. There has never been an `app/internal/roadrunner/`.
- The **seven** dirs the corpus's repaired position names all exist as trees under `internal/`:
  `cms`, `customeriosync`, `jobsimulation`, `messenger`, `skiller`, `skillpath`, `storage`
  (`git ls-tree ad9f3c49 internal/`, 73 entries).
- Judge0 lives in the **jobsim** domain: `internal/jobsimulation/runner` exists, and
  `internal/jobsimwiring/wiring.go:123` is byte-exact
  `runnerManager := jsrunner.NewRunnerManager(getenv("JUDGE0_API_KEY"), getenv("JUDGE0_BASE_URL"))`.

**The corpus contradicts itself, and its own repaired side is the correct one:**

- `corpus/services/README.md:39` — *"**`roadrunner` is NOT one of them**: `app/internal/roadrunner/`
  exists at no ref … This row listed a 'roadrunner domain' until M257x iter-102"* — i.e. this exact
  phrasing is booked, by name, as a defect already repaired elsewhere.
- `corpus/services/README.md:20-24` — *"`roadrunner` is **the eighth, and it is different**:
  orphaned, **not merged-and-undeployed**"*, because `roadrunner/terraform/main.tf:19` still reads
  `service_desired_count = 1` — which I confirmed byte-exact at `roadrunner` `87d8d443`. So
  roadrunner is not merely mis-listed; the list **flattens the one case that is materially
  different in production**.
- `corpus/architecture/platform-migration-status.md:88` (the `app` row) — *"Owns **seven** domains
  in-process … **`app/internal/roadrunner/` does not exist**"*.

Not `retraction-not-contradiction`: `:9` retracts nothing, it asserts the false membership plainly.
This is the half-repair class — corrected in `services/README.md` and the fenced map at iter-102,
left standing in the twin that no partition owned (`:9` last moved at `904502c`, iter-87).

### P-3 — `ai-readiness.md:18-20`, the `WorkforceDirectory` seam (r34-B2)

Corpus text, live at HEAD, unrepaired (lines 18/19/20 verified by `grep -n`):

> *"**The only remaining dependency on `workforce` is the member directory** (the
> `WorkforceDirectory` interface — `LoadMembers`/`LoadMembersByUserIDs`, whose implementations
> **stayed** in `app/internal/workforce/members.go`)."*

**Ground truth at `app` `ad9f3c498`** (no ref named in the block ⇒ graded at the checkout). Read
`internal/aireadiness/manager.go` in full at that ref:

- The interface declares **FOUR** methods — `LoadMembers` (`:43`), `LoadMembersByUserIDs` (`:45`),
  **`BaseMembers`** (`:48`), **`LevelsCount`** (`:50`), between `type WorkforceDirectory interface {`
  (`:40`) and `}` (`:51`).
- The **source's own type doc comment**, `manager.go:36-39`, is verbatim: *"WorkforceDirectory is
  the slice of the workforce manager this domain needs: the active-member directory … **and the
  org's skill-scale setting**."* The struct-field comment at `:121-122` says the same — *"workforce
  supplies the member directory **+ org skill scale**"*.
- `LevelsCount` is an **org setting**, not a member call: `readiness.go:770` is
  `maxLevel := int(m.workforce.LevelsCount(ctx, orgID))` and `:771` `hwm.MaxSkillLevel = maxLevel`.
- **The "stayed in `members.go`" clause is wrong for the fourth member.** `LoadMembers` /
  `LoadMembersByUserIDs` / `BaseMembers` are at `internal/workforce/members.go:349` / `:353` /
  `:357`; **`LevelsCount` is at `internal/workforce/manager.go:90`** (`git grep "func .*LevelsCount"`
  returns exactly three sites tree-wide: `manager.go:61` the unexported `getLevelsCount`,
  `manager.go:90` the exported one, and a test fake).
- Coupling calibration, re-derived: `internal/aireadiness` imports `internal/workforce` in **8
  non-test files** (`auto_assign.go`, `breakdown_enrichment.go`, `freeze_backfill.go`,
  `live_snapshots.go`, `manager.go`, `readiness.go`, `recommendation_engine.go`,
  `steps_completion.go`); `workforce.Member` occurs **60** times in that package. Both figures
  reproduce the seat's exactly.

Every one of the seat's three sub-findings reproduced byte-for-byte. **But see my framing
disagreement below — I found the corpus did not invent this universal, and that changes the
repair.**

---

## Rejections, with the evidence I opened

None. Four claimed, four upheld.

For the record, the classes I actively tested and ruled out rather than skipped:

- **`wrong-tree` (0).** Both seats stated their trees explicitly and read the **pinned** clone
  `09d06070` for tooling-on-a-stack claims and the authoring copy `f2ea567b` only for
  fence-configuration claims, and said which. None of the four blockers depends on either rext tree
  — they are settled in `messenger`, `app` and the corpus itself — so the two-rext-tree trap does
  not bite here at all. Neither seat touched the dirty `ant-academy` working tree for a blocker.
- **`true-at-its-ref` (0).** P-1 names its own ref (`fa47850d`) and is false there *and at every
  ref in that repo's history*. P-2 and P-3 name no ref and were graded at the brief's checkout.
- **`retraction-not-contradiction` (0).** None of the three anchors is doing retraction work. The
  `sentinel.md:5` **paragraph** is a retraction, but the falsified clause is the *fresh* evidence
  the retraction offers in place of the expired claim, not the retraction itself.
- **`not-in-scope` (0).** All three anchors are under `corpus/services/**` /
  `corpus/architecture/**`.

## Cannot-settle

None in my assigned blocker set. (Both seats' *own* "could not settle" lists — `infrastructure`,
`db-backup`, `chronos`, GitHub repo metadata, live-stack runtime figures — are correctly
non-booked and I did not convert any of them into a verdict.)

## Framing disagreements — where I part from how the seats stated the predicate

**1. r34-B2: the corpus inherited the false universal from the platform's own package doc; the seat
did not notice, and this changes the repair.** The seat's headline is *"`ai-readiness.md` states a
universal … that the interface it cites refutes in its own doc comment."* That is true of the
**type** doc comment (`manager.go:36-39`). But the corpus's sentence is a near-verbatim lift of the
**package** doc comment at `manager.go:7-9`:

> *"It was split out of internal/workforce: workforce keeps the org-analytics KPIs, aireadiness owns
> everything readiness-scoped. **The only dependency on workforce is the member directory
> (WorkforceDirectory)** — the population being scored is the same one the workforce dashboards
> list."*

So `app` contradicts **itself**: its package doc (`:7-9`) says "only … member directory", while its
type doc (`:36-39`), its struct-field comment (`:121`) and its own four-method declaration
(`:40-51`) all say "member directory **and** the org's skill-scale setting". The corpus adopted the
stale half. Three consequences the seat's framing misses: **(a)** the falsified proposition is best
stated as a property of the *code contract*, not of the corpus author's care; **(b)** the repair is
not "fix the corpus sentence" alone — a corpus fix leaves the platform's package doc stale and the
next reader re-imports it; **(c)** this is testimony-vs-evidence — the same rule r34-B itself
invoked correctly elsewhere in its own report (grading `d11a403` on the **diff**, not the commit
message) applied inconsistently here. I have stated **P-3** as a proposition about the interface
rather than about the sentence, for exactly this reason.

I also register that the seat's **sub-finding 1 alone** ("four methods, not two") is the weakest of
its three: read charitably, the em-dash could gloss rather than enumerate. **Sub-finding 3** (the
`members.go` pointer, false for `LevelsCount`) is the one that cannot be read charitably, and it
carries the blocker on its own.

**2. r33-B2 imports a count the corpus text does not state.** The seat's title is *"still lists
`roadrunner` as one of **EIGHT** domains"*. `dependency_map.md:9` lists eight names but **never
says "eight"** — the word appears in `services/README.md:20` (*"roadrunner is the eighth"*), which
is the *opposing* passage. The defect is a **membership** predicate, not an **arithmetic** one, and
I have booked P-2 that way. This matters for cross-seat dedup: a seat booking *"the `app` row's
'seven domains' figure is wrong"* would be a **different** predicate at a different anchor, and
must not be merged into P-2 on the strength of the shared word "eight".

**3. r33-B2's counter-citation has drifted one line.** The seat cites
`platform-migration-status.md:87` for the `app` row; at corpus HEAD `:87` is the table's separator
row and the `app` row is `:88`. This is the seat's own citation, not a corpus defect, so it changes
nothing — but the brief warns adjudicators not to accept a citation unopened, and I record that I
opened it and had to correct it.

**4. P-1 is the weakest of my three on materiality, and I say so rather than hiding it.** Both
seats independently reached BLOCKER and both independently flagged the same hesitation (r33-B: *"I
hesitated"*; r34-B: *"high [confidence] that the command returns 0; **medium** on the
BLOCKER/MINOR grading"*). That is agreement about the *fact* — which I confirmed and then
strengthened — but it is **not** independent corroboration of the *grading*, and it should not be
counted as such. A reasonable adjudicator could class P-1 `minor-not-blocker` on the ground that
the conclusion it supports is true and in fact stronger. I upheld it because (i) the brief's UPHELD
bar is "the corpus text really is false" and it plainly is; (ii) the carve-out is for
cosmetic/immaterial defects and a false, non-self-correcting empirical measurement is neither; and
(iii) the figure was **never true at any ref**, so it is an error rather than drift. If the
milestone wants a materiality-weighted headline as well as a truth-weighted one, P-1 is the
predicate to discount.

## Counts

```
UPHELD=4 REJECTED=0 (of which wrong-tree=0) CANNOT-SETTLE=0
DISTINCT-PREDICATES-IN-MY-SET=3
```
