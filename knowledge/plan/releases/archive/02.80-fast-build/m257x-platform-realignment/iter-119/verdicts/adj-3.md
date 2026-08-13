# Adjudicator 3 — seats D and F (readings #31 and #32)

## Trees read, and at which refs

`git status --porcelain` was **empty at my open** and **empty at my close**. Rosetta HEAD `c18d56bc`
(branch `m257x/platform-realignment`). No fetch, no pull, no checkout, no git state change; the only
file I created is this one.

Every ref below I re-derived myself with `git -C <path> rev-parse HEAD` before grading; all sixteen match
the brief's GROUND TRUTH table exactly.

| tree | path | ref I read |
|---|---|---|
| platform | `stack-demo/platform` | `0c91421dfdb08dc75f17f1aabfb61394070e770b` |
| app | `stack-demo/app` | `ad9f3c498e9c244187440562f83c11e5408d6554` |
| app/studio (nested, own checkout) | `stack-demo/app/studio` | `aeec036a51c8a4ae0c5b8f7d5d21cfa7086b658e` |
| cms/studio (nested, own checkout) | `stack-demo/cms/studio` | `aeec036a51c8a4ae0c5b8f7d5d21cfa7086b658e` |
| sentinel | `stack-demo/sentinel` | `f2c461903de022a6a506a3a10355dbf503515ce5` |
| messenger | `stack-demo/messenger` | `fa47850d9c507d1928da7a38f7b37bac1bb8fabc` |
| ant-academy | `stack-demo/ant-academy` | `22df69dd81f1d718ecc9c088bbf96b6ae681c3a2` |
| next-web-app | `stack-demo/next-web-app` | `8297c684caacefb84ae2bcdbf0135795268d6341` |
| studio-desk | `stack-demo/studio-desk` | `41ee3575ddd94930148706fff05e18aa805cc19a` |
| cms | `stack-demo/cms` | `ca50c8170fefe1122d680efe54f7e56798a79d82` |
| jobsimulation | `stack-demo/jobsimulation` | `462343b05c4f796513a43327d4d8d62d99128c4f` |
| storage | `stack-demo/storage` | `4ce8ece52adb7c095e792e235da4a8913214d190` |
| roadrunner | `stack-demo/roadrunner` | `87d8d44382ef07a9f165869530cbac9e5e0a4332` |
| graphql-wundergraph | `stack-demo/graphql-wundergraph` | `60c229f39adcbbe75c84cd58f0f45052b5423372` |
| rosetta-extensions (**pinned per-stack**) | `stack-demo/rosetta-extensions` | `09d06070fd99c742d7a671c468abf93074278575` |
| rosetta-extensions (**authoring**) | `.agentspace/rosetta-extensions` | `430493087c56199a79c24430f819ebbd46b10a58` |

**No booking in my group turned on a rext claim**, so no `wrong-tree` question arose; I nevertheless
verified both rext refs so that a rext-shaped rebuttal could be answered.

Corpus files I opened at rosetta HEAD (line counts match what both seats reported, so the subject did not
move under me): `service_taxonomy.md` 523 · `sentinel.md` 166 · `ant-academy.md` 503 ·
`architecture_overview.md` 435 · `roadrunner.md` 216 · `dependency_map.md` 103. Also opened for
cross-claim checks: `external_services.md`, `shared_libraries.md`, `architecture/README.md`.

Per the HARD BAR, the only things I read under `knowledge/plan/**` are the brief, my four assigned seat
reports, and this file.

---

## Verdicts

### D r31 B1

```
D r31 B1 | corpus/architecture/service_taxonomy.md:175 | UPHELD | IN-SCOPE | PREDICATE: The private Go modules a stack's Docker build pulls are exactly colony, proto, taxonomy.
```

   evidence: I opened `stack-demo/app` at the ref the surrounding section names (`ad9f3c49`) and read
   `go.mod` directly. The **direct** `require` block, lines 14–18, is
   `anthropos-work/analytics-go v0.3.1`, `colony v0.35.2`, `proto v1.210.0`, `storage v0.15.2`,
   `taxonomy v1.2.0` — **five** private `anthropos-work` modules, none marked `// indirect`, none listed
   in `platform 0c91421d:repos.yml` (which has exactly four entries: app, sentinel, next-web-app,
   studio-desk). I did not stop at `go.mod`: `git grep -l '"github.com/anthropos-work/storage'
   ad9f3c49 -- '*.go'` returns **32 files** (`internal/academy/asset.go`, `internal/cms/wiring.go`,
   `internal/jobsimulation/anticheat/anticheat.go`, …) and `analytics-go` is imported by
   `internal/payments/handler.go` and `internal/tracking/handler.go`. So both omitted modules are live
   compile-time dependencies of the one Go service every stack builds, pulled by exactly the
   `GH_PAT`/`GOPRIVATE` mechanism the sentence's own parenthetical defines
   (*"imported as private Go modules — not cloned by `make init`; pulled at Docker build via
   `GH_PAT`/`GOPRIVATE`"*). The sentence supplies a definition and then enumerates a set that fails it.
   The corpus already carries the correct five-member set at
   `corpus/architecture/external_services.md:554` — verbatim *"`app/go.mod:14-18` requires `analytics-go`,
   `colony`, `proto`, `storage`, `taxonomy` and nothing else"* — so this is falsity plus an intra-corpus
   conflict.
   I tested the charitable reading before upholding: the surrounding block heads a five-row
   *Shared Libraries* table, so *"of those five, three survive"* would be true, and that is exactly how
   the three sibling formulations are worded (`service_taxonomy.md:518` *"5 libraries, 3 imported by a
   service a stack builds"*; `architecture/README.md:21`; `shared_libraries.md:24-25`
   *"only three by the two repos a stack actually builds"*). `:175` is the one site that drops the
   scope and asserts *"the live private-module set a stack builds is colony, proto, taxonomy"* as an
   unscoped universal. Upheld on the words as written, and because the omission is load-bearing:
   `analytics-go` appears in the whole corpus only at
   `platform-migration-status.md:157` and in that `external_services.md` sentence, and the
   shared-libraries page the sentence points to as *"Full reference"* does not know either module exists.

### D r31 B2

```
D r31 B2 | corpus/services/sentinel.md:5 | UPHELD | IN-SCOPE | PREDICATE: That grep over messenger at fa47850d returns one hit; it returns zero.
```

   evidence: The brief's rule-1 multi-pin warning is the first thing to settle here, because this
   paragraph carries **five** refs — `d11a403`, `0dab54d`, `0c91421d`, `838d907` and `fa47850d`. Working
   out which one dates which proposition: `d11a403`/`0dab54d` date the *"`AUTHORIZATION_ADDRESS` is set in
   exactly one block"* clause; `0c91421d` dates the *"declares five services"* clause (which I confirmed
   independently — `git show 0c91421d:docker-compose.yml` declares exactly `sentinel:5`, `backend:28`,
   `studio-desk:112`, `next-web-app:143`, `gotenberg:170`, so that clause is **TRUE at the ref it names**
   and is *not* what is booked here); `838d907` dates the deletion sentence. The booked proposition is the
   last one in the paragraph and **names its own ref, `fa47850d`** — so per rule 1 the earlier pins do not
   date it, and no ref-discipline defence is available.
   Run verbatim, in `stack-demo/messenger`, at `fa47850d`:
   `git grep "authorization\|AUTHORIZATION_ADDRESS" fa47850d -- '*.go' go.mod` → **no output, exit 1**.
   I then ruled out all three instruments named in rule 3 before accepting the zero:
   (a) `git grep -in "authoriz" fa47850d` whole-tree → **0**; `git grep -ain` (binary-as-text) → **0**;
   (b) NUL census over **every** tracked blob at that ref, byte-counted with `tr -dc '\000' | wc -c`, not
   `grep -c` → **no NUL-bearing file exists** in this repo, so the mechanism-2 skip cannot apply;
   (c) a working-tree `grep -rIl -i authoriz .` (which sees untracked files too) → empty.
   Positive controls in the same pass: `git grep -c "colony" fa47850d -- go.mod` → `go.mod:7`
   (one line), and the `'*.go'` pathspec resolves (`git grep -l func … -- '*.go'` → 39 of the 42 `.go`
   files at that ref). So the pipeline is live and the zero is a measurement.
   I went one step further than either seat: `git log --all -S"authorization" -- '*.go' go.mod` and
   `git log --all -S"AUTHORIZATION_ADDRESS"` in messenger both return **nothing**, so the printed
   *"one unrelated hit"* was never true at **any** ref in that repo's history — it is not a stale
   transcription of an older sha.
   I considered and rejected the ambiguity defence the seat itself raised (that *"returns one unrelated
   hit"* might attach to the `colony` control, which does return exactly one line). The clause's subject
   is the quoted command; *"against `colony` present as a positive control"* is a separate adjunct.
   Upheld: the corpus publishes a re-derivation command **with its expected output**, in the one sentence
   in this file that advertises itself as *"re-derived"*, and a reader who runs it gets a result that
   disagrees with the document — the exact failure mode (an honest zero reading as a broken pipeline)
   that the instrument's own search-discipline rules exist to prevent, manufactured here by the corpus.
   The conclusion the parenthetical supports (*messenger imports no authorization client*) is TRUE and
   is not what is booked; the repair is *"zero hits"*, not a change of conclusion.

### D r32 B1

```
D r32 B1 | corpus/architecture/service_taxonomy.md:403 | UPHELD | IN-SCOPE | PREDICATE: service_taxonomy.md's archive-state note is at line 142.
```

   evidence: Read `service_taxonomy.md:136-162` myself. `:139-146` is a blockquote headed
   *"⚠️ Two different fates shared this table, and the second one has now closed."*, and `:142` is
   literally *"> so a bare `make up` started all three as unfederated husks — the
   `running_but_unfederated` state in"* — a sentence about compose fates, nothing about archive state.
   The **archive-state note** (*"Every `ARCHIVED <date>` in this table is a DATED SNAPSHOT, not a derived
   fact."*) is a separate blockquote at `:148-160`. The *Archived / merged* table header is at `:162`, so
   both blockquotes sit "above the table" and the citing text does not disambiguate by position — only
   the number does, and the number lands in the wrong one. The anchor resolves to a real, well-formed
   line, which is why it gives a reader no signal that it is not the corroboration promised.
   I checked rule 7 before upholding: this is **not** a historical anchor. The sentence is a live
   *"see …"* pointer in present tense, not a record of where something once was. It also is not the
   rule-7-approved construct-naming form: it names the construct **and** supplies a line, and the line is
   wrong. The irony is inside the target itself — `:156-160` records that this file's own row-number
   self-pins have already rotted **twice** and concludes *"A same-file line pin into a growing table is
   not worth its self-heal — search the row name."* `:142` is a third live instance of the pathology the
   file documents.

### D r32 B2

```
D r32 B2 | corpus/architecture/service_taxonomy.md:406-407 | UPHELD | IN-SCOPE | PREDICATE: The sentence "There is no `graphql` profile" is at service_taxonomy.md:67-68.
```

   evidence: Read `service_taxonomy.md:60-80`. `:67` is `- **Communication**: HTTP/RPC + Redis Streams`
   and `:68` is the `- **Database**: PostgreSQL — one schema, public, owned by app…` bullet. Neither
   contains the quoted string, and neither mentions profiles at all. The quoted sentence —
   *"**There is no `graphql` profile, and no cms / jobsimulation / roadrunner / storage / messenger /
   customerio-sync service of any kind.**"* — begins at the end of **`:74`** and completes on `:75`.
   The citing clause says *"Consistent with `:67-68` **above**"*, so the reference is unambiguously
   intra-file (and `:74` is indeed above `:406`, so the direction word is fine here — only the number is
   wrong). Wrong construct, off by seven lines, landing on two unrelated bullets.

### D r32 B3

```
D r32 B3 | corpus/architecture/service_taxonomy.md:175 | UPHELD | IN-SCOPE | PREDICATE: The private Go modules a stack's Docker build pulls are exactly colony, proto, taxonomy.
```

   evidence: Same anchor, same proposition, same re-derivation as **D r31 B1** above — `app@ad9f3c49`
   `go.mod:14-18` carries five private `anthropos-work` requires, `storage` imported by 32 `.go` files
   and `analytics-go` by 2. Collapses onto **P1**; it is one predicate at one anchor booked by the same
   seat in two readings, not a second anchor.

### D r32 B4

```
D r32 B4 | corpus/services/sentinel.md:5 | UPHELD | IN-SCOPE | PREDICATE: That grep over messenger at fa47850d returns one hit; it returns zero.
```

   evidence: Same anchor and same proposition as **D r31 B2**; re-derived once, exhaustively, as recorded
   there (verbatim run → exit 1; `-i`/`-a`/whole-tree → 0; NUL census clean; untracked sweep clean;
   two positive controls pass; `git log --all -S` shows the string never existed in this repo).
   Collapses onto **P2**.

### F r31 B1

```
F r31 B1 | corpus/services/ant-academy.md:31 (twins :5, :24) | UPHELD | IN-SCOPE | PREDICATE: Ant Academy is an internal-only portal gated to @anthropos.work employees.
```

   evidence: Rule 5 is the gate here — a retraction is not a self-contradiction unless the retraction and
   the thing it retracts are **both asserted as live**. They are. `:474` books the proposition as dead
   (*"⚠️ 'Domain-gated to `@anthropos.work` so external users cannot enter' is FALSE and was removed at
   M257x iter-115"*), while `:31` still asserts, in the document's own voice, present tense,
   *"**Primary Goal**: Internal-only learning portal that delivers AI-engineering chapters to
   `@anthropos.work` employees"*, `:5` asserts *"the **internal learning portal** for Anthropos employees
   … to anyone with an `@anthropos.work` email"*, and the service-overview table at `:24` asserts
   *"**Authentication** | Clerk (`@anthropos.work` **domain gate** + org-membership gate)"*. Live claim
   and live retraction, 443 lines apart.
   I did not adjudicate from the seat's quotes — I opened `stack-demo/ant-academy` at `22df69dd` myself:
   - **No domain gate exists.** `git grep "endsWith('@\|endsWith(\"@\|emailAddresses\|primaryEmail"` over
     `code/` returns only display/analytics uses, plus `@anthropos.work` **only** in `code/TESTING.md`
     persona docs and e2e fixtures (`_helpers.js` pairs `alice@anthropos.work` with `bob@gmail.com`).
     There is no allowlist, no `ALLOWED_DOMAIN`, no email-suffix predicate in any production path.
   - `code/proxy.js:112-186` — I read the whole `isPublic` matcher — opens `/`, `/latest(.*)`,
     `/chapters/(.*)`, `/courses`, `/courses/(.*)`, `/library`, `/library/(.*)`, `/free`, `/free/(.*)`,
     `/verify/(.*)`, `/api/verify/(.*)`, `/catalog.json` and more to **anonymous** traffic, with inline
     comments calling `/` *"M4 public catalog — the front door. Anonymous visitors browse…"*.
   - `code/src/lib/platformUrls.js:1-32` calls the Academy *"a **storefront** in front of the Anthropos
     platform"* and defines FLOW A (*"a new visitor registers"*) and FLOW B (*"Checkout contextually
     registers → signs in → subscribes for an **anonymous visitor**"*).
   - `code/src/lib/pricing.js:21` `export const STANDARD_YEARLY = { usd: 399, eur: 349 }` with a launch
     coupon; `code/src/lib/schema.js:3` `SITE_URL = 'https://aiacademy.anthropos.work'`.
   - `knowledge/user-types.md` enumerates exactly **Anonymous · Signed-in (free) · Subscriber ·
     Enterprise/Org member**, detected by `auth()` / `isInOrg()` / `billingPremium()` — **zero**
     `@anthropos.work` predicate anywhere in the matrix.
   The only real gate is org membership (`REQUIRE_ORGANIZATION_MEMBERSHIP`, `proxy.js:91`, `:310`,
   → `/no-organization`), which is the half `:474` says survives. A product that sells a $399/yr
   subscription to an anonymous visitor is not internal-only.

### F r31 B2

```
F r31 B2 | corpus/architecture/architecture_overview.md:40 (twin :260) | UPHELD | IN-SCOPE | PREDICATE: Ant Academy is an internal-only portal gated to @anthropos.work employees.
```

   evidence: Opened both lines at rosetta HEAD. `:40` — *"**Ant Academy** (`ant-academy`): **Internal**
   learning portal (Next.js 16 + Expo mobile) **for `@anthropos.work` employees**. Deployed on Vercel."*
   `:260` — the Frontend-applications table row, *"Internal learning portal for `@anthropos.work`
   employees (standalone, Vercel-deployed)"*. Neither names a ref, so both are graded at ground truth,
   and both are refuted by exactly the `ant-academy 22df69dd` evidence I re-derived under F r31 B1 — and
   by the corpus's own adjudication at `ant-academy.md:474`.
   Booked separately by the seat because a file-scoped repair would leave these standing; that reasoning
   is sound, but under the brief's dedup rule the **predicate is the same one**, so this contributes two
   further anchors to **P5** rather than a new predicate. (I confirmed `grep -n "anthropos\.work"` over
   `architecture_overview.md` returns exactly three lines — `:40`, `:178` (the unrelated
   `content.anthropos.work` Directus host) and `:260` — so the leak in this file is these two sites and
   no more.)

### F r31 B3

```
F r31 B3 | corpus/architecture/architecture_overview.md:80 | UPHELD | IN-SCOPE | PREDICATE: `ai` is one of the private Go modules a local stack's Docker build pulls.
```

   evidence: `:80` reads *"4. **Shared Libraries**: **four** imported private modules — colony, proto,
   **ai**, taxonomy (not deployed; pulled at Docker build)"*, and it sits as item 4 of a numbered list
   whose header at `:70` is *"**Service Tiers** (local development reality, default `core` profile)"* —
   so the predicate is scoped, by the document, to what a local `core` build pulls. No ref is named, so
   it is graded at ground truth.
   Re-derived: `git show ad9f3c49:go.mod | grep anthropos-work` → analytics-go, colony, proto, storage,
   taxonomy — **no `ai`**. `git show f2c46190:go.mod | grep anthropos-work` (sentinel) → colony, proto,
   taxonomy (`// indirect`) — **no `ai`**. `platform 0c91421d:repos.yml` has four entries, of which the
   Go ones are `app` and `sentinel`, so those two **are** the local Docker builds. `ai` survives as a
   requirement only in `cms/go.mod:9` and `jobsimulation/go.mod:11` (both `v1.40.2`), neither of which
   `make init` clones. I confirmed the mechanism: `app b948604f:go.mod:14` still carried
   `github.com/anthropos-work/ai v1.40.2`; `1e457fa70` (2026-08-04,
   *"refactor(ai): fold the ai library into app as internal/ai"*) removed it, and
   `internal/ai/module_import_guard_test.go` exists at `ad9f3c49` to keep it out one-way.
   This also contradicts four sibling corpus statements I opened:
   `dependency_map.md:42` (*"the two Go repos a stack clones and builds … pull in **THREE** of them —
   colony, proto, taxonomy … `app` **dropped** the `ai` module at `1e457fa70`"*),
   `shared_libraries.md:24-25`, `architecture/README.md:21` and `external_services.md:554` (*"neither
   `app/go.mod` nor `sentinel/go.mod` requires `github.com/anthropos-work/ai`"*).
   I weighed the charitable reading (*"four of the five are pulled by **at least one repo**"*, which is
   `shared_libraries.md`'s true framing) and rejected it: that framing ranges over the seven Go repos on
   disk, whereas `:80`'s own list header restricts it to *local development reality* and its own
   parenthetical says *pulled at Docker build*. In its own declared scope the claim is false.
   **Note this is the opposite error from P1**, not the same one: `service_taxonomy.md:175` says the set
   is three and is too small; `architecture_overview.md:80` says it is four and includes a module nothing
   local pulls. Two different false propositions about the same set — kept as separate predicates, and
   they contradict each other as well as ground truth.

### F r31 B4

```
F r31 B4 | corpus/services/roadrunner.md:130 | UPHELD | IN-SCOPE | PREDICATE: roadrunner.md's "Upstream consumers: none (orphaned)" line is at :124, below :130.
```

   evidence: Read `roadrunner.md:118-140` myself. `:124` is *"The repo contains an experimental WebSocket
   LSP proxy (`internal/lsp/lsp.go`) that is NOT wired into any running server — there is no reachable
   LSP endpoint today."* — a true, substantive claim about a **different subsystem**. The quoted string
   is at **`:134`**: *"* **Upstream consumers**: **none (orphaned — see the banner at the top).**"*
   `grep -n 'Upstream consumers' corpus/services/roadrunner.md` returns exactly two lines, `130` (the
   citing sentence's own quotation) and `134`, so `:124` cannot be a different-but-valid target.
   Doubly wrong: the citing clause says *"Consistent with `:124` **below**"*, and `:124` is six lines
   **above** `:130`, so no reading of "below" reaches it. The citing sentence's entire rhetorical job is
   to claim a second independent site in the same file agrees with it; the pointer it supplies does not
   go there, and lands somewhere plausible enough that a reader gets no signal.

### F r31 B5

```
F r31 B5 | corpus/services/ant-academy.md:304 | UPHELD | IN-SCOPE | PREDICATE: ant-academy has ~26 Playwright e2e spec files; it has 31.
```

   evidence: Set enumerated first, per rule 4, rather than re-deriving the document's arithmetic.
   `git show 22df69dd:code/playwright.config.js` sets `testDir: './tests/e2e'`, so the directory the cell
   names is the directory Playwright runs. `git ls-tree -r 22df69dd --name-only code/tests/e2e/` returns
   **40 paths**; I listed them all. Anchored on the basename ending, `*.spec.js` is **31**. The other nine
   are `_helpers.js`, `_academyBackendMock.js` and seven PNGs under `widgets.spec.js-snapshots/` — and
   note an unanchored `grep -c '\.spec\.js'` gives 38 because it also matches that snapshot **directory**
   name, so the anchoring is what makes 31 the right cardinality. Excluding the one underscore-prefixed
   spec still gives 30. There is no reading of the directory that yields 26.
   **This is the weakest predicate I uphold and I say so explicitly.** The figure carries an explicit `~`
   hedge, no reader doing real work is materially misled, and the same seat's other reading demoted it to
   a MINOR. I uphold because the brief's UPHELD test is falsity against ground truth, not materiality;
   because none of the enumerated rejection classes fits (it is not ref-discipline — the cell names no
   ref; not mis-read — I reproduce 31 independently; not already-true); and because a 19 % miss on a set
   you can `ls` exceeds what a tilde licenses in a document that elsewhere publishes exact cardinalities
   and fences them. A grader who reads `~` as covering 26→31 should strike this one predicate; nothing
   else in my group depends on it.

### F r32 B1

```
F r32 B1 | corpus/services/roadrunner.md:130 | UPHELD | IN-SCOPE | PREDICATE: roadrunner.md's "Upstream consumers: none (orphaned)" line is at :124, below :130.
```

   evidence: Same anchor and same proposition as **F r31 B4**; re-derived once as recorded there
   (`:124` = the WebSocket-LSP bullet, `:134` = the quoted line, only two `Upstream consumers`
   occurrences in the file, and `:124` is above not below `:130`). Collapses onto **P7**.

---

## Rejections

**None.** Zero of the twelve bookings failed re-derivation, and I record what I actively tested for so
that the zero is legible as a measurement rather than as agreement:

- **ref-discipline** (the class the brief says ran 17 occurrences across five readings for zero graded
  contribution): I checked every booking for a pin that would date the claim past the contradicting
  evidence. The only booking where a pin was in play is `sentinel.md:5`, and there the pin **points the
  other way** — the sentence names `fa47850d`, which is the ref at which the printed result is wrong, and
  the string it greps for has never existed in that repo at any ref. The four other refs in that
  paragraph date other propositions (I mapped each one; the `0c91421d`-dated five-service clause is
  **TRUE** and is not booked). No claim in my group is pinned, past-tense or dated in a way that excuses
  it.
- **wrong-tree**: no booking in my group turns on a `rosetta-extensions` claim, so the frozen
  instrument's line-37 defect could not bite. Both rext refs verified anyway.
- **historical-anchor**: tested against the three intra-corpus pointer bookings (P3, P4, P7). None is a
  record of *"what a prior audit found at line N"*; all three are live present-tense *"see :N"* pointers,
  and none is the rule-7-approved construct-name-instead-of-line form (`:403` names the construct **and**
  a wrong line; `:406` and `:130` quote the text **and** give a wrong line).
- **pointer-not-assertion** (rule 6): tested against P1/P6. Neither is a pointer at a
  derived-once-elsewhere value — `service_taxonomy.md:175` and `architecture_overview.md:80` each state
  their own enumeration, in their own voice, and disagree with the sites that derive it.
- **mis-read**: every cited line was opened and read with ±10 lines of context; I reproduced each seat's
  cardinality independently from source (`go.mod` requires, `.go` importer counts, the `tests/e2e/`
  tree, the compose service list) before accepting it.
- **already-true**: the closest call was P1's charitable *"of the five shared libraries"* scoping, which I
  worked through in full above and which the sentence's own definitional parenthetical defeats.

---

## PREDICATE ROLL-UP

```
P1 | The private Go modules a stack's Docker build pulls are exactly colony, proto, taxonomy. | anchors: D r31 B1 @ corpus/architecture/service_taxonomy.md:175, D r32 B3 @ corpus/architecture/service_taxonomy.md:175
P2 | That grep over messenger at fa47850d returns one hit; it returns zero. | anchors: D r31 B2 @ corpus/services/sentinel.md:5, D r32 B4 @ corpus/services/sentinel.md:5
P3 | service_taxonomy.md's archive-state note is at line 142. | anchors: D r32 B1 @ corpus/architecture/service_taxonomy.md:403
P4 | The sentence "There is no `graphql` profile" is at service_taxonomy.md:67-68. | anchors: D r32 B2 @ corpus/architecture/service_taxonomy.md:406-407
P5 | Ant Academy is an internal-only portal gated to @anthropos.work employees. | anchors: F r31 B1 @ corpus/services/ant-academy.md:31 (twins :5, :24), F r31 B2 @ corpus/architecture/architecture_overview.md:40 (twin :260)
P6 | `ai` is one of the private Go modules a local stack's Docker build pulls. | anchors: F r31 B3 @ corpus/architecture/architecture_overview.md:80
P7 | roadrunner.md's "Upstream consumers: none (orphaned)" line is at :124, below :130. | anchors: F r31 B4 @ corpus/services/roadrunner.md:130, F r32 B1 @ corpus/services/roadrunner.md:130
P8 | ant-academy has ~26 Playwright e2e spec files; it has 31. | anchors: F r31 B5 @ corpus/services/ant-academy.md:304
```

Dedup notes:
- **P1** and **P6** are deliberately **not** collapsed. They are opposite errors about the same set —
  `:175` omits `storage` and `analytics-go`; `:80` adds `ai`. Two seats booking these would not write the
  same sentence, and the two anchors contradict each other as well as ground truth.
- **P3**, **P4** and **P7** are deliberately **not** collapsed into one "wrong intra-corpus anchor"
  predicate. Each names a different construct at a different wrong line in a different place (two files);
  the false propositions are distinct, not the same falsehood in more places.
- **P5** is deliberately **collapsed** across two files. F booked them separately on repair-scope grounds,
  which is good bookkeeping, but it is one false proposition at three anchors in `ant-academy.md` and two
  in `architecture_overview.md`.
- **P1/P2/P7** each collapse a reading-#31 and a reading-#32 booking of the same seat at the same anchor.

Anchor count vs predicate count for my group: **8 distinct predicates** across **9 distinct corpus
anchors** (P5 carries two file-level anchors; every other predicate carries one), booked twelve times.

---

`git status --porcelain` at close: the **only** entry is the untracked
`knowledge/plan/releases/02.80-fast-build/m257x-platform-realignment/iter-119/verdicts/` directory into
which this panel's verdict files are written — i.e. no tracked file changed, and I created nothing
outside my one assigned output path. (The sibling `adj-*.md` files in that directory belong to the other
three adjudicators; I did not open them.) Read-only throughout; no fetch; no git state change.

BOOKED=12 UPHELD=12 REJECTED=0 IN-SCOPE-UPHELD-BLOCKERS=12 DISTINCT-IN-SCOPE-PREDICATES=8
