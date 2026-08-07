# iter-128 — the three diseases, treated separately

**Type:** tik · **Run 81.** The run's directive named three priorities in consequence order and this
document follows that order, not the order the work was easiest in.

---

## 1. `state.md` — back inside its cap, and a budget that has no owner to move to

**The breach was 21 bytes; the defect was 1,085.** `state.md` was 15,381 against a 15,360 cap. Fixing it
by trimming would have bought exactly one run — which is what
[`context.md` § state.md contract](../../../context.md) says in as many words, and why it authored
**per-field budgets** the day before. Measured against those budgets, one field was the whole problem:

| field | before | budget | after |
|---|---|---|---|
| `phase` | **1,985** | 900 | **861** |
| all five others | within | — | within |
| frontmatter | 2,791 | 2,600 | **1,667** |
| **file** | **15,381** | **15,360** | **14,826** |

*"`phase:` is the field that breaks first, every time"* — the contract's own words, confirmed on its first
enforcement. The repair was the contract's method: **move content to its owner.** iter-124's reframe went
to `iter-124/audit.md`; the standing rules to their `§5` and `D-N` owners.

### 1a. The thing the move found — a 38 KB deliverable with no owner

`corpus/architecture/org-repos.md` (38,842 bytes, all 93 org repos, iter-123's net-new deliverable) and
`corpus/ops/observability.md` were **recorded nowhere but `state.md`'s `phase:` field.** iter-123's own
`progress.md` never named either.

`phase:` is the field **every close overwrites**. A deliverable whose sole record is the rotating index is
one close away from being unattributable — and the compression I was about to perform is exactly the
event that would have erased it. `iter-123/progress.md` § 6 now owns both, plus the 8-site `db-backup`
repair. **This milestone's own *"a census is not a disclosure"* rule, pointed at its own plan record.**

### 1b. The finding: the BODY budget is not trimmable by the contract's own method

The body is **13,163 against a 12,000 budget** and it did not come down. Three probes, each asking *"who
else owns this?"*:

| probe | result |
|---|---|
| § Standing backlog | `roadmap-vision.md` (the contract's declared owner) mirrors **7 of 14** items. `PERF-M256-parallel-lane`, `PT-M257-self-evaluation`, `PT-M257-talk-to-data`, `platform-defect-register`, `DEF-M250-01`, `CAVEAT-1`, `PT-M256-resume-fixture-pair` appear **nowhere else** |
| M255 provenance clause | `roadmap.md` carries the *numbers* (4.84 GB → 379 MB, 146.8 → 2.9 s) but **not** the provenance (09:59–11:37Z, PRE-freeze, the 658/666/672 cluster, the three claims owed re-confirmation) |
| § Process flags | the `stackseed`-pin trap, the local-only tag state, the rext code-of-record — **sole owner** |

**In all three, `state.md` IS the owner.** *"Move it to its owner"* has no target. The 12,000-byte body
budget was set without measuring what the body uniquely owns, and it cannot be met by the method the same
contract prescribes — only by trimming prose, which the directive forbids and which would destroy the
sole record of seven backlog items.

**Not silently absorbed:** `state.md` carries a one-line pointer here, and the item is routed as
`FIX-M257x-iter128-body-budget-has-no-owner`. The honest options are to raise the body budget to a
measured floor or to give the un-mirrored content a real owner; **both are decisions, not edits**, so
neither was taken unilaterally.

### 1c. My own defect, caught by a reader and not by any fence

Building the new `phase:` through `.encode().decode('unicode_escape')` mangled a literal em dash into
`â\x80\x94`. Repaired — and repaired as a **repo-wide sweep** over all tracked files for nine
latin-1-round-trip damage sequences (0 other files affected) rather than as a one-character patch.

**No guard in this family checks encoding integrity, and no guard of the existing shape could.** The two
files that mention encoding at all (`corpus_index_guard.py:84`, `story_org_count_guard.py:220`) carry an
`except UnicodeDecodeError` — an *unreadable-file* handler. **Mojibake is VALID UTF-8**: `â\x80\x94`
decodes without error, it is simply the wrong three characters. A decode-error handler cannot see it by
construction.

It was caught by the previous run reading my output. Filed as what it is: a defect class with **no fence
and no possible fence of the current shape**, found by a human-shaped read.

---

## 2. The 820 — triaged, and the printed split corrected by its own audit

Run 80 named the complement of the sealed consequence class **"untriaged, not implied clean."** That was
the right call and this is the debt it created, paid.

**The partition is exact and asserted at runtime**, drawn by the same committed predicate:

```
|C1| 340  +  |complement| 820  ==  |tier-2 population| 1160
```

`iter-128/triage-complement.py` **imports** `iter-124/triage-predicate.py` and `iter-124/triage.py`
rather than copying them, so the two triages cannot drift and iter-124's published figures stay
byte-reproducible. **Anti-vacuity control:** mutating the predicate to match-anything collapses the
complement to 0 — the complement is genuinely predicate-derived, not hardcoded.

| fate | printed | **corrected** | |
|---|---|---|---|
| `cite` | 815 (99.4 %) | **≈ 738 (90.0 %)** | |
| `drop` | 5 (0.6 %) | **≈ 49 (6.0 %)** | colon lead-ins + imperative checklist steps |
| `hedge` | 0 (0.0 %) | **≈ 33 (4.0 %)** | almost entirely `chronos` + `db-backup` |
| `fix` | 0 | **a FLOOR of unknown height** | see the limitation below |

**The correction is measured on this population, not imported.** A seeded 30-of-329 **R4-only** hand
audit puts R4 at **76.7 %** here. All 7 disagreements run one way — `cite` → `drop`/`hedge` — the
direction R4 was declared generous in *before either sample was drawn*. Full table:
[`audit-complement.md`](audit-complement.md).

**Three limitations, and the first is load-bearing:**

1. **`fix = 0` over the complement means nothing was read, not nothing is false.** The triage cannot
   decide falsity. This run's reading was aimed at C1 by the consequence-ordering rule. **The
   complement's false-claim count is UNMEASURED.**
2. **R3 — 486 members, 59 % of the complement — was not re-audited**; iter-124's 100 % was carried
   forward. That is the largest untested block inside the corrected number.
3. iter-124 sampled the whole class, this run sampled R4-only. **66.7 % and 76.7 % are not a
   before/after pair** and must not be quoted as a trend.

**Two mechanical sub-classes named, because a future rule could catch them:** colon-terminated lead-ins
to a code block and imperative checklist steps (all 4 drops); and *"the subject is a repo no clone set
contains"* — `chronos`, `db-backup` (all 3 hedges), which `R2` never learned about because its
`UNCHECKABLE` list is keyed on vendor-internal phrases, not un-cloned first-party repos.

---

## 3. Priority 1 — the consequence-class false claims, found and repaired

**Method.** Six independent readings, one slice each, over the sealed consequence class
(**`|C1|` = 340**, the tier-2 members whose sentences touch security, data-handling, identity,
tenancy, residency, backup or access). **340 / 340 read.** Every clone was read **at a ref** via
`git show <ref>:<path>` / `git grep <ref>` — never a working tree, because **five of thirteen clones are
behind their own origin** (`storage` by 20 commits) and the standing rule is that a stale substrate
*manufactures evidence against a true claim*. Nothing in `stack-demo/` was written.

**Result: 13 distinct false claims → 30 repaired sites**, because four of the thirteen are **predicates**
whose twins had to be swept (§5 rule 54). 17 files, +137/−38 lines.

### 3a. The headline — the third security-surface UNDERSTATEMENT of this milestone

> `security_compliance.md` said the REST surface *"has no blanket authz middleware — two of its **six**
> groups opt into one (`cbGate`), the rest authorize per handler or not at all."*

`app` mounts **ELEVEN** non-test Echo groups on the **one** REST instance
(`internal/web/web.go:124-163` @ `ad9f3c498`). The table enumerated only the six declared inside
`backend.go`. Of the five it never saw, **three never touch the Clerk `authn` middleware**, and one has
**no authentication at all**:

| missing group | gate |
|---|---|
| **`/api/invitations`** | **`cors` ONLY.** `web.go:145-146`: *"Public invitation JSON endpoints (no auth required)"* |
| `/content/admin` | no Clerk authn — a bearer shared secret is the whole gate |
| `/v1/labs` | no Clerk authn — an org API key + `labs:write` scope |
| `/academy/embeddings` | `cors` + `authn` |
| `/api/workforce` | `cors` + `authn`, but grouped off the **root** `e`, so despite its `/api/` prefix it does **not** inherit the `/api` stack |

**This is the THIRD correction to that one paragraph.** iter-120 over-stated it (*"every Echo group … and
nothing else"*); iter-121 corrected the quantifier; run 81 found the **denominator** had been six all
along. **Both earlier repairs re-derived from `backend.go` — because that is where the previous sentence
pointed.** A count is only as wide as the search that produced it, and a repair that inherits its
predecessor's search inherits its blind spot.

### 3b. The four predicates, and the widest one

| predicate | sites | note |
|---|---|---|
| *"Ant Academy is an internal `@anthropos.work`-only portal"* | **14 in 9 files** | refuted at **iter-115 in `ant-academy.md`'s own Clerk bullet** — which then sat **469 lines below an opening sentence that still said it**. `CLAUDE.md` carried it too |
| *"the Clerk org role set is {`admin`, `basic_member`}"* | **5** | there are **three**; `content_creator` is set on Clerk memberships, read back from webhooks, **synced into Sentinel**, and gates studio-desk entirely. The file already named it in its SDK table — enumeration and counter-example in one document |
| the REST group count | 3 | above |
| *"OAuth / social — not used, mobile is email+password only"* | 1 | **both halves false** — Google **and** Microsoft Clerk OAuth, at `origin/main` `f97ba6599` **and** at `8297c684`, the ref the page pins, so not clone drift |

**My first reach regex found 4 of the 14.** Widening it found the rest. **The near-miss was the
enumeration, not the repair** — recorded because it is the same failure the headline describes, committed
by me, one hour after writing it down.

### 3c. The nine singletons, by consequence

- **`clerkenstein.md`** — remote reach called *"opt-in via `--public-host`"*; it has been **DEFAULT-ON
  (opt-out) on the demo path since M220** (`D-DESIGN-3`). **The exposure axis** — a reader believed a
  bare `/demo-up N` stayed local when it auto-publishes an unauthenticated, authz-weakened stack.
- **`cms.md`** — `StudioDocument` glossed *"(simulation blueprints)"*; it is a **customer-uploaded
  attachment and its extracted full text**. A **data-classification error in the very doc that declares
  those tables 100 % customer data**, and the schema notes it carries no Ent privacy policy.
- **`messenger.md`** — Redis credited with *"scheduled-message storage"*. Messenger stores nothing in
  Redis and `Schedule` is `CodeUnimplemented` — which this page already said 12 lines above. A
  retention/GDPR reader was sent to a store that holds no payloads.
- **`backend.md`** — `jobsimfeedback` called *"signals routed back to the skills domain"*. It is a
  **satisfaction survey**: no skill field, no skill edge, no route into the skills domain at all — and
  the mislabel hid that the table holds **user-authored free text**.
- **`ai-labs.md`** — grader/solution assets attributed to `s3_workspace_store.go`, which backs exactly
  one object kind. **The privacy claim it wrapped is TRUE**; it was bolted to the wrong file, sending an
  auditor of that asset's access control to a store that never holds it.
- **`ant-academy.md`** — the mobile bundler *"bundles `code/public/content/`"*. It reads
  **`<repo-root>/content/`**, deleted at `8199ea5d`: **0** entries there against **3,406** under
  `code/public/content/`. The script ENOENTs and bundles nothing.
- **`graphql-wundergraph.md`** — `/graphql` serves Apollo Sandbox **in development only**; unqualified,
  it over-stated production exposure.
- **`ant-academy.md`** — *"the platform's `/ant-*` skills in Rosetta"*: **there are none** (16 skills, 0
  matching).
- **`hiring.md`** — two in-file pins landed on the wrong construct because **iter-102 re-derived them by
  adding +23 and +16 to the old numbers instead of re-measuring**. *Arithmetic on a citation is not a
  citation.* Re-measured: `:196-209` and `:176-187`.

### 3d. Reach — with its denominator, and its limit

| statement | number | denominator |
|---|---|---|
| C1 read | **340 / 340 = 100 %** | the sealed consequence class at this tree |
| distinct false claims | **13** | of 340 read |
| sites repaired | **30**, 17 files | 13 claims × their predicate twins |
| census movement | **1,160 → 1,150** | the repairs added evidence; ratchet holds |
| **C1's share of the corpus's consequence surface** | **37.1 %** | 1,598 consequence-class sentences in the census's 41 files vs 2,714 in `corpus/ops/**` + `corpus/tools` + `CLAUDE.md`, which the census **never reads**. `CLAUDE.md` alone holds **108**, 0 enumerated |

**That last row is the honest limit of this reading, and it is the one to quote.** The census scope is
`corpus/services/*.md` + `corpus/architecture/*.md`. **`corpus/ops/**` and `CLAUDE.md` are outside it
entirely** — and `CLAUDE.md` is the file every agent loads. (Caveat, stated because the two numbers are
different measures: 1,598 counts *all* consequence-class sentences in scope, while `|C1|` = 340 counts
only the **unevidenced** subset. The 37.1 % is a statement about **which files the census reads**, not a
defect rate.) Predicate B was found *inside* C1 but **8 of its 14 sites were outside the census scope**
— the enumeration only reached them because the repair was driven by predicate rather than by the
census's own list.

### 3e. Against the prediction

iter-124 printed `fix = 4` over `|C1| = 344` and recall-corrected it to **≈ 11**. This run read the same
class exhaustively and found **13**. The estimate was **good** — but note what it was an estimate *of*:
iter-124's own `fix` sites were *repaired*, so they had left tier-2 before this run began (the triage
prints `fix = 0` now). **13 is a fresh count over a re-derived population, not 4 + 7**, and the two must
not be added. The agreement is evidence the recall correction was sound; it is not a measurement that
the class is now empty — **`fix` remains a floor**, and 8 of the 13 were found only because the reading
went outside the census's own flagged set.
