# Adjudication E — seats `iter-131/raw/r33-E.md`, `iter-131/raw/r34-E.md`

## Scope line

Independent re-adjudication of **16 booked BLOCKERS** across the two Seat-E readings (8 + 8), covering
`corpus/services/{messenger,academy-backend,graphql-wundergraph,ant-academy}.md`. Every verdict below rests
on evidence **I opened myself** at the ref the claim names (ground-truth table in the brief, plus the
authoring rext copy for the one fence-configuration reading). I read the brief and these two seat reports
and **nothing else** under `knowledge/plan/**`. Read-only apart from this file.

**Corpus currency.** `git diff --stat 90cbd3e HEAD -- corpus/` shows **none of my four target files changed
since the seats were dealt** (the only post-131 corpus edits are in `README`s, `architecture/*`,
`platform-alignment.md`, `platform_repo.md`, and `services/{askengine,backend,clerk-integration,cms,
jobsimulation,skiller,storage}.md`). So **no verdict here is `UPHELD (since-repaired)`** — every quoted
defect is still live at HEAD `0dd19e5`, and I graded the same bytes the seats did.

## Counts table

| | count |
|---|---|
| **claimed (blockers booked)** | **16** (r33-E 8, r34-E 8) |
| **UPHELD** | **16** |
| REJECTED — `wrong-tree` | 0 |
| REJECTED — `misread` | 0 |
| REJECTED — `true-at-its-ref` | 0 |
| REJECTED — `retraction-not-contradiction` | 0 |
| REJECTED — `minor-not-blocker` | 0 |
| REJECTED — `not-in-scope` | 0 |
| **CANNOT-SETTLE** | **0** |
| **wrong-tree errors (total)** | **0** |
| **upheld rate** | **16/16 = 100 %** |
| **DISTINCT-PREDICATES-IN-MY-SET** | **9** (see the mechanism caveat below — 5 of the 9 collapse to one cause) |

> **Do not read 100 % as a validation of the instrument.** This assignment is unusually citation-heavy —
> **12 of 16** booked blockers are `file:line` / path citations, the single most mechanically verifiable
> class in the corpus. A seat cannot easily be wrong about whether line 193 says a thing. The rate would
> not survive a set weighted toward interpretive platform claims.

## Verdict table

| seat | B# | anchor | verdict | class (rejection n/a) | predicate | finding class | multi-pin | repair-induced (sha) |
|---|---|---|---|---|---|---|---|---|
| r33-E | B1 | `corpus/services/messenger.md:121` | **UPHELD** | — | P-E1 | self-contradiction (+platform-drift) | no | no (`ebf8097`) |
| r33-E | B2 | `corpus/services/messenger.md:96` | **UPHELD** | — | P-E1 | self-contradiction (+platform-drift) | no | no (`ebf8097`) |
| r33-E | B3 | `corpus/services/academy-backend.md:66` | **UPHELD** | — | P-E2 | platform-drift (+self-contradiction) | **yes** | **yes — `2ae1052`** (iter-122) |
| r33-E | B4 | `corpus/services/academy-backend.md:15` | **UPHELD** | — | P-E3 | intra-corpus-citation | **yes** | **yes — `434caa8`** (iter-124) |
| r33-E | B5 | `corpus/services/academy-backend.md:136` | **UPHELD** | — | P-E4 | intra-corpus-citation | **yes** | no (`f8be5a1`, iter-108) |
| r33-E | B6 | `corpus/services/graphql-wundergraph.md:136` | **UPHELD** | — | P-E5 | intra-corpus-citation | **yes** | no (`cd16967`, iter-102) |
| r33-E | B7 | `corpus/services/graphql-wundergraph.md:88` | **UPHELD** | — | P-E7 | intra-corpus-citation | **yes** | no (`e858fd4`, iter-98) |
| r33-E | B8 | `corpus/services/ant-academy.md:63` | **UPHELD** | — | P-E8 | platform-drift (+self-contradiction) | no | no (`af66289`) |
| r34-E | B1 | `corpus/services/messenger.md:121` | **UPHELD** | — | P-E1 (dup of r33-E B1) | self-contradiction | no | no (`ebf8097`) |
| r34-E | B2 | `corpus/services/messenger.md:96` | **UPHELD** | — | P-E1 (dup of r33-E B2) | self-contradiction | no | no (`ebf8097`) |
| r34-E | B3 | `corpus/services/academy-backend.md:66` | **UPHELD** | — | P-E2 (dup) | platform-drift | **yes** | **yes — `2ae1052`** |
| r34-E | B4 | `corpus/services/academy-backend.md:136` | **UPHELD** | — | P-E4 (dup of r33-E B5) | intra-corpus-citation | **yes** | no (`f8be5a1`) |
| r34-E | B5 | `corpus/services/academy-backend.md:15` | **UPHELD** | — | P-E3 (dup of r33-E B4) | intra-corpus-citation | **yes** | **yes — `434caa8`** |
| r34-E | B6 | `corpus/services/graphql-wundergraph.md:136` | **UPHELD** | — | **P-E5 + P-E6** (two anchors in one sentence) | intra-corpus-citation | **yes** | no (`cd16967`) |
| r34-E | B7 | `corpus/services/graphql-wundergraph.md:88` | **UPHELD** | — | P-E7 (dup of r33-E B7) | intra-corpus-citation | **yes** | no (`e858fd4`) |
| r34-E | B8 | `corpus/services/ant-academy.md:45` | **UPHELD** | — | P-E9 | self-contradiction (+platform-drift) | no | **yes — `e75906b`** (iter-128) |

Repair-induced determined mechanically per the brief: `git log -L<n>,<n>:<file> --oneline | head -3`, most
recent touching commit. **Three anchors qualify** (`2ae1052`, `434caa8`, `e75906b`). See the causal caveat
in *Disagreements* §3 — the mechanical rule and the causal truth diverge on two of the three.

## Upheld predicates, deduplicated within my assignment

```
P-E1 | messenger's own code Liquid-renders the transactional email body before the Brevo send
     | (rather than Brevo rendering it from a Brevo-hosted template)
     | anchors: corpus/services/messenger.md:96, corpus/services/messenger.md:121
     | class: self-contradiction

P-E2 | the `app` repo contains `env/callsites_test.go`, and the STORAGE_RPC_ADDR comment callsite lives there
     | anchors: corpus/services/academy-backend.md:66
     | class: platform-drift

P-E3 | `academy-backend.md:85-89` corroborates the NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT / app-subgraph statement
     | anchors: corpus/services/academy-backend.md:15
     | class: intra-corpus-citation

P-E4 | `ant-academy.md:82-88` states in bold that there is no FS-as-published catalog fallback
     | anchors: corpus/services/academy-backend.md:136
     | class: intra-corpus-citation

P-E5 | `graphql-wundergraph.md:116-117` is the struck-through "`make up` rebuilds `graphql`" bullet,
     | and `:114-117` are the "Build-time, static composition" bullets
     | anchors: corpus/services/graphql-wundergraph.md:136
     | class: intra-corpus-citation

P-E6 | `graphql-wundergraph.md:84` is the *Ports* bullet
     | anchors: corpus/services/graphql-wundergraph.md:136
     | class: intra-corpus-citation

P-E7 | `graphql-wundergraph.md:193` states that `http://localhost:5050` refuses the connection
     | anchors: corpus/services/graphql-wundergraph.md:88
     | class: intra-corpus-citation

P-E8 | `ant-academy/code/tools/` holds the offline-parity CLI
     | anchors: corpus/services/ant-academy.md:63
     | class: platform-drift

P-E9 | Ant Academy chapters are delivered offline on the Expo mobile bundle
     | anchors: corpus/services/ant-academy.md:45
     | class: self-contradiction
```

**DISTINCT-PREDICATES-IN-MY-SET = 9** (per the brief's literal rule: a distinct proposition = a distinct
predicate). **Mechanism-level, the honest number is 5** — see *Disagreements* §2.

## The evidence I opened, per predicate

### P-E1 — messenger Liquid rendering (`messenger.md:96`, `:121`)

`messenger.md:121` reads verbatim: *"Messages carry user info, template ID, and template params; the body is
rendered through Liquid against those params before the Brevo send."* `messenger.md:96` reads
`` message/                 Message types + Liquid rendering ``.

Measured at `messenger` `fa47850d` (the brief's ref; identical conclusion at `origin/main` `e9421c68`):

* `git show fa47850d:internal/messenger/brevo/brevo.go` — `func (s *brevoSender) send(...)` at `:288`
  builds `brevo.SendSmtpEmail{… TemplateId: templateId, Params: props}` and posts it to
  `s.brevo.TransactionalEmailsApi.SendTransacEmail` at `:310`. **Brevo renders.** I read the function body,
  not a grep count.
* Complete non-test Liquid surface, repo-wide (`git grep -in liquid fa47850d -- '*.go'`, stderr read, 20
  hits enumerated by hand): `internal/flow/assignments.go:18,:489`, `internal/flow/whitelabel.go:8,:17`,
  `internal/messenger/console/console.go:16,:26,:33,:71`. The rest are `_test.go`. **None is on the Brevo
  send path.**
* `internal/messenger/message/` at `fa47850d` = `errors.go`, `message.go`, `message_test.go`,
  `validator.go`. I read `message.go` **in full**: it defines `New`, `NewWithDefaultSender`,
  `DefaultSenderUser` and `ConvertProps` (a JSON marshal/unmarshal into `map[string]any`). No rendering.
  `git grep -in "template\|render\|liquid"` over that path → exit 1; **positive control** `func` over the
  same path → 3/3/2 hits, so the pipeline was live.
* The nuance the seats did not state, and it strengthens rather than weakens the finding: Liquid **is**
  applied before *some* Brevo sends — `brevo.go:36-46` and `:64-67` show a whitelabel template (259) and a
  CMS-custom template (249) that carry an **inline-composed `custom_subject`/`custom_body`** rendered in
  `flow/`. So the true statement is *"the body is Brevo-rendered from a hosted template, except for
  whitelabel/CMS-custom categories whose body is composed in `flow/` first"* — which is **exactly what
  `:15-21` says** and exactly what `:121` and `:96` contradict.

`:15-21` is a **retraction**, not a contradiction, and I did not uphold anything against it. What I upheld
is that the retraction **did not reach two sites**, one of them 100 lines below it and one 25 lines below
it. That is the corpus's own §5 rule-54 class.

### P-E2 — `env/callsites_test.go` (`academy-backend.md:66`)

At `app` `ad9f3c49`, `git grep -n "STORAGE_RPC_ADDR" ad9f3c49 -- '*.go'` returns **exactly three** lines —
`internal/jobsimwiring/wiring.go:101`, `internal/storagens/callsites_test.go:189`, `main.go:504` — all
comments. **The count of 3 and the "all comments" predicate are both right; only the path is wrong.**

Three things push this past a typo:

1. `git ls-tree -r --name-only ad9f3c49 | grep -E '(^|/)env/'` → empty. There is no `env/` directory.
2. `git log --all --oneline --name-only --diff-filter=A -- '*callsites_test.go'` over `app`'s **whole
   history, all refs** → the only path ever added is `internal/storagens/callsites_test.go`.
   **`env/callsites_test.go` has never existed.** This is not drift; it is a fabricated path.
3. `grep -rn "callsites_test" corpus/` → **four** sites. Three (`architecture_overview.md:363`,
   `platform-migration-status.md:93`, `storage.md:29`) carry the **correct** path. Only
   `academy-backend.md:66` carries `env/`. So the corpus contradicts itself 3-to-1, and the odd one out is
   the one introduced most recently (`2ae1052`, iter-122 — repair-induced).

### P-E3 / P-E4 / P-E5 / P-E6 / P-E7 — the five self-citations

I opened each pin and each named target in the live tree.

| pin | what the sentence says it is | what it actually is | delta |
|---|---|---|---|
| `academy-backend.md:15` → `:85-89` | corroborates the WunderGraph-endpoint / subgraph claim | `:85` tail of the `academy_path_embeddings` Ent bullet; `:86-88` **Certificate minting**; `:89` blank. Target is `:93-96` | +8, across a `## Interface Discovery` boundary |
| `academy-backend.md:136` → `ant-academy.md:82-88` | *"states the opposite in bold"* (no FS-as-published fallback) | the **store.js / beacon write-path blockquote**. Target is `ant-academy.md:96-102` | +14, cross-file |
| `graphql-wundergraph.md:136` → `:116-117` | the struck-through *"`make up` rebuilds `graphql`"* bullet | the *"…requires re-running `wgc compose` … no hot reload"* bullet. Target is `:118-119` | +2 |
| `graphql-wundergraph.md:136` → `:84` | *"the *Ports* bullet … about `8080`/`5050`"* | `* **Federation**: Apollo Federation v2, federation_version: =2.3.2 (pinned)`. Ports is `:86-90` | +2 |
| `graphql-wundergraph.md:88` → `:193` | *"already said `localhost:5050` refuses the connection"* | the *"⚠️ At `0dab54d` the `graphql` token appears in no `profiles:` key"* warning. Target is `:196` | +3 |

`grep -n "refuses the connection" corpus/services/graphql-wundergraph.md` → `:89` (the citing sentence) and
`:196` only. `awk` line-addressed reads for every other row; nothing inferred.

**The finding that reframes all five (see *Disagreements* §1): each was CORRECT when written.**

* `git show cd16967:corpus/services/graphql-wundergraph.md` — at iter-102, `:84` **was** the Ports bullet
  and `:116-117` **was** the struck-through bullet.
* `git show e858fd4:corpus/services/graphql-wundergraph.md` — at iter-98, `:193` **was**
  `` `http://localhost:5050` refuses the connection. ``
* `git show cd16967:corpus/services/academy-backend.md` — at iter-102, `:85-89` **was** the
  *"Primary — GraphQL, on the `app` (backend) subgraph … NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT"* bullet.
* `git show f8be5a1:corpus/services/ant-academy.md` — at iter-108, `:82-88` **was**, verbatim and in bold,
  *"⚠️ It does *not* 'fall back to the committed FS catalog' — there is **no FS-as-published fallback** in
  the app"*.

So this is **pure line-pin rot**, not a repair writing a wrong anchor. Each pin was invalidated by a LATER,
unrelated edit that inserted lines above the target. Verdicts are unchanged — none of these five sentences
carries a ref, so the only tree they can be graded at is the live document, where every one of them is
false — but the causal story the seats told is wrong (§1 below).

**Materiality, stated because I nearly split this cluster.** My rule, applied uniformly: a citation defect
is a BLOCKER when the sentence **names a construct** and the pinned range **does not contain it**; it is
`minor-not-blocker` only when no belief a reader forms can change. All five name a construct
(*"the Ports bullet"*, *"the struck-through … bullet"*, *"states the opposite in bold"*, *"Consistent
with"*, *"already said `localhost:5050` refuses the connection"*) and none of the five ranges contains it.
I record that P-E7 (+3, same paragraph) and P-E6 (+2) are the two most recoverable, and that a
materiality-graded instrument would plausibly demote them.

### P-E8 — `code/tools/` = "offline-parity CLI" (`ant-academy.md:63`)

`git ls-tree -r --name-only 22df69dd8 -- code/tools/` → **one file**, `code/tools/apply-v3-metadata.mjs`,
whose own header (read) says *"Adds sidebarCategory, sidebarSubcategory, and audiences fields to every skill
path in catalog.js."* A metadata migration.

The offline-parity CLIs are **regression-fenced as absent** in the repo under test:
`code/tests/unit/next-scaffold.test.js:111` — `it('v0.5 M1: offline-parity scripts + serwist/esbuild deps
removed', …)` — plus the deleted-file table at `:161-163` enumerating `tools/offline-parity/lib.mjs`,
`capture-baseline.mjs`, `check-parity.mjs` with `expect(existsSync(...)).toBe(false)`. I read those lines.
The same corpus file states the v0.5 M1 removal at `:48` and `:316`.

### P-E9 — "offline on the Expo mobile bundle" (`ant-academy.md:45`)

`git show 22df69dd8:mobile/scripts/bundle-content.ts` — `:27` `const REPO_ROOT = join(SCRIPT_DIR,'..','..')`,
`:28` `const CONTENT_DIR = join(REPO_ROOT,'content')`; `findJsonFiles` at `:38-50` does a bare
`readdirSync(dir)` with **no existence check** and is called with `CONTENT_DIR` at `:53`.
`git ls-tree -r --name-only 22df69dd8 -- content/` → **0**; `-- code/public/content/` → **3,406**. The
bundler ENOENTs and ships nothing.

What decided this one for me is the sentence's **own register**: `:45` takes the trouble to correct the WEB
half in-line (*"but **no longer offline on the web** — the Serwist service worker was removed at v0.5 M1"*)
and then asserts the mobile half in the unhedged present tense — while the summary bullet 26 lines above
(`:19-23`) carries the hedge *"**intended to** bundle"* **and** the measurement, and the deep-dive at
`:451-459` states it a third time. A sentence that is doing correction bookkeeping in its own parenthesis is
not in a product-aspiration register.

## Rejections, with the evidence I opened

**None.** Every booked blocker in both seats stands. I list here the four candidates I actively tried to
reject and why each survived:

1. **`graphql-wundergraph.md:88` / `:193` (P-E7)** — 3 lines short, inside the paragraph that contains the
   target, so a reader recovers. Survived because the *same sentence* mis-describes a **second** anchor:
   its parenthetical calls `:174-176` *"the compose line-number caveat"*, and `:174-176` is the head of the
   **Upstream consumers** bullet — the caveat is `:179-181`. Both seats logged that half (r34-E as a minor);
   I opened it and it holds. A correction note that mis-describes two anchors in one sentence is not
   cosmetic.
2. **`ant-academy.md:63` (P-E8)** — a one-line annotation in an ASCII layout diagram. Survived because the
   named artifact provably does not exist and its absence is asserted by a live test in the repo the corpus
   is describing.
3. **`academy-backend.md:15` (P-E3)** — a self-citation, not a platform claim. Survived because it is a
   checkable statement about this document's content, it is false, and it has now been re-pointed **three
   times** (`:74-76` → `:80-83` at iter-97 → `:85-89` at iter-102) and is wrong again.
4. **r33-E B2 as a separate blocker from B1** — I considered `minor-not-blocker` for the code-block
   annotation. Survived as a blocker (a key-directories map is how a reader navigates a repo) but **not as a
   separate predicate** — see §2.

**No `wrong-tree` errors.** Both seats stated their refs up front, both read `ant-academy` via
`git show 22df69dd8:<path>` rather than the dirty tree, and both cross-checked the pinned rext clone
`09d06070` against the authoring copy `f2ea567b` and reported them identical for every rext anchor. I
re-checked the one rext read I needed (`stack-core/corpus_citation_guard.py`): `git diff f2ea567b HEAD` on
that file is **empty**, so my reading holds at the brief's ref even though the authoring copy has advanced
6 commits (now `223e4a6e`) — worth flagging for the next brief.

## Cannot-settle

**None among the 16 blockers.** Neither seat booked a blocker that depends on `infrastructure`, so the
`infrastructure`-not-in-the-clone-set gap both seats flagged in their *"could not settle"* sections
(`graphql-wundergraph.md:7`, `:27`, `messenger.md:59` @ `13c248e6`) produces **no CANNOT-SETTLE in my
assignment**. Both seats correctly declined to book it either way. I confirmed the substrate gap is real
(`find stack-demo -maxdepth 4 -name .git` enumerates no `infrastructure` tree) and did nothing further with
it, since it grades no claim I was assigned.

## Disagreements with how the seats framed their predicates

### 1. The causal framing of the citation cluster is wrong in r33-E, and right in r34-E

r33-E B6: *"this is not ambiguity — it is a **mis-anchor introduced by the very repair (iter-102)** that the
sentence advertises."* r33-E B7: *"This is the **second re-pointing** of the same citation … and it is still
off."* Both readings say the repair pass **wrote a wrong number**.

**It did not.** I checked out the document at the authoring commit in every case:

| pin | authored at | correct there? |
|---|---|---|
| `gql-wg.md:136` → `:116-117`, `:84` | `cd16967` (iter-102) | **yes, both** |
| `gql-wg.md:88` → `:193` | `e858fd4` (iter-98) | **yes** |
| `academy-backend.md:15` → `:85-89` | `cd16967` (iter-102) | **yes** |
| `academy-backend.md:136` → `ant-academy.md:82-88` | `f8be5a1` (iter-108) | **yes, verbatim and in bold** |

Every one was accurate when written and was invalidated by a **later, unrelated** edit inserting lines above
the target (+2 / +3 / +8 / +14). r34-E framed this correctly — *"§5 rule 34 … this is the correction note
itself having **rotted**"* — and r33-E did not. The distinction is not cosmetic: it decides whether the
remedy is *"repair harder"* (r33-E's implied reading — the repair author was careless) or *"fence the
citation form"* (the true reading — no author could have prevented it).

**And the corpus already knows.** `rext stack-core/corpus_citation_guard.py` (FENCE-M257x-iter117,
byte-identical at the brief's `f2ea567b`) declares **exactly this blind spot** in its own docstring:
*"**Bare `:NN` pins** (`` `:525` ``) are NOT mechanically decidable and are **excluded outright** … the
honest statement is that the construct half of the class is not reachable by a machine at scale — only its
resolution half is."* Its C3 arm asserts only that a pinned line **exists**, never that it **says the
thing**. **All five drift blockers land inside the declared exclusion.** That is the actionable finding, and
neither seat reached it.

### 2. r33-E's B1/B2 split is over-booking; they are ONE predicate

r33-E B2: *"**Booked separately from B1** because it is a distinct assertion about a distinct package with
its own evidence."* The same report, two sentences earlier: *"**Same measurement as B1**"* and *"This is the
**second surviving site of the claim the file retracts at `:15-21`**."* Those cannot both be true.

The corpus retracted **one phrase** — *"Liquid templating for the bodies"* — at `:15-21`; it survives at
`:96` and at `:121`. Under the brief's rule (*"two seats booking the same proposition at two different
anchors share ONE predicate"*) this is **one predicate, two anchors**. I have counted it as one. A milestone
that counts it as two is counting **reach failure** — how many sites a repair missed — and calling it
**predicate count**, which are different quantities and should not be summed into the same headline.

### 3. "Repair-induced" as the brief defines it mislabels two of my three positives

The brief's rule is mechanical: *most recent touching commit is in iters 120–130 → yes.* Applied, it flags
`academy-backend.md:66` (`2ae1052`), `academy-backend.md:15` (`434caa8`) and `ant-academy.md:45`
(`e75906b`). But causally:

* **`academy-backend.md:66` — genuinely repair-induced.** `git log -S "env/callsites_test.go" -- corpus/`
  returns exactly one commit, `2ae1052` (iter-122). iter-122 wrote a path that has never existed in `app`,
  while three sibling corpus sites already carried the right one. **True positive.**
* **`academy-backend.md:15` — the flag is right for the wrong reason.** `434caa8` (iter-124) touched line 15
  to fix the *router-in-production* clause and carried the stale `(Consistent with :85-89 below.)` through
  untouched. The pin was broken **earlier**, by `2ae1052`'s own expansion of the AssetUploader bullet at
  `:60-71`, which inserted ~8 lines **above** `:85` and pushed the target to `:93-96`. So the
  repair-that-broke-it and the repair-that-carried-it-forward are **different commits**, and `git log -L` on
  the citing line names only the latter.
* **`ant-academy.md:45` — a reach failure, not an induction.** `e75906b` (iter-128) rewrote line 45 to
  correct the *storefront* predicate and left *"and offline on the Expo mobile bundle"* standing; iter-129
  then repaired the **summary bullet** at `:19-23` for that same predicate and did not come back to `:45`.

The general point: **`git log -L` on the CITING line cannot find the commit that broke a self-citation,
because that commit touched lines ABOVE it and never touched the citation at all.** Any repair-induction
rate computed that way under-attributes the drift class and mis-attributes it to whichever later commit
happened to brush the line. If iter-135 is going to publish a repair-induced fraction, this is the
measurement bug to fix first.

### 4. Predicate arithmetic: 9 by the letter, 5 by mechanism — publish both

By the brief's literal rule (a distinct proposition = a distinct predicate) my set is **9**. But **five of
the nine** (P-E3…P-E7) are the same mechanism — an unpinned intra-corpus line reference that was true when
authored and rotted when the target grew — and would be closed by **one** fence covering **one** citation
form. If the milestone's headline P is meant to answer *"how many distinct things is the corpus wrong
about"*, this cluster inflates it **5×** for a single cause. My recommendation: report **P = 9 (5 distinct
mechanisms)** rather than either number alone.

## Counts

```
UPHELD=16 REJECTED=0 (of which wrong-tree=0) CANNOT-SETTLE=0
DISTINCT-PREDICATES-IN-MY-SET=9   [5 distinct mechanisms]
```
