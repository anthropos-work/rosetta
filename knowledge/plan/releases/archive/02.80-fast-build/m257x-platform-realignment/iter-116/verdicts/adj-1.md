# adj-1 — verdicts for seats r29-A / r30-A / r29-B / r30-B (M257x iter-116)

## Ground truth I re-derived at my own open (no fetch, no git state change)

Corpus HEAD `6b0cb96d`, working tree clean. The only corpus commits since the reading's SEAL
(`85f6f1c`) are the fourteen seat-report landings under `knowledge/plan/**`; **no corpus file under
`corpus/**` moved between the seats' reads and mine**, so every same-file line pin I graded is graded
at the same content the seats saw.

Clone refs re-read at my open, all matching the brief's table:
`platform 0c91421d` · `app ad9f3c49` · `app/studio aeec036a` · `cms ca50c817` · `cms/studio aeec036a` ·
`next-web-app 8297c684` · `sentinel f2c46190` · `studio-desk 41ee3575` · `ant-academy 22df69dd` ·
`jobsimulation 462343b0` · `messenger fa47850d` · `storage 4ce8ece5` · `roadrunner 87d8d443` ·
`graphql-wundergraph 60c229f3` · `stack-demo/rosetta-extensions 09d06070`.

**Which rext tree I read (brief §"Two clone sets exist"):** every `rosetta-extensions` measurement below
was taken in the **pinned per-stack consumption clone `stack-demo/rosetta-extensions @ 09d06070`**, because
the only rext-touching booking (r30-A B1) is a claim about a *census over the stack's own tree set*, not
about a fence's verdict or configuration. I did not settle anything from `.agentspace/rosetta-extensions`.
No `wrong-tree` rejection arises.

Tree census re-derived independently before any tree-wide statement:
`find stack-demo -name .git -maxdepth 4` → **15** git trees (13 top-level + the two nested `studio`
checkouts). Nested repos were grepped at **their own ref `aeec036a`**, never through the host ref.

---

## Verdicts

### r29-A B1 | `corpus/services/askengine.md:81`, `:113` | UPHELD | IN-SCOPE | PREDICATE: The `ai.AI` interface app's ask-engine uses is a shared `ai` library (a private Go module).

   evidence: I opened `stack-demo/app` at `ad9f3c49` (the ref this document itself pins at `askengine.md:89`).
   `git show ad9f3c49:go.mod` requires exactly five `anthropos-work` modules at `:14-18` — `analytics-go`,
   `colony`, `proto`, `storage`, `taxonomy` — and **`github.com/anthropos-work/ai` is not among them**; the
   only tracked references to that module path anywhere in `app` are its own import-guard test
   (`internal/ai/module_import_guard_test.go:18` `const externalModule = "github.com/anthropos-work/ai"`),
   a CI guard workflow, and knowledge/CHANGELOG prose. I then opened the ask-engine's actual embedding path:
   `internal/web/backend/ask/embed.go` imports `"github.com/anthropos-work/app/internal/ai"` and calls
   `aiClient.CreateEmbeddings(ctx, text, ai.WithEmbeddingModel(ai.EmbeddingDefaultModel))` — an **in-tree
   package**, not a module. Corpus-side this is also a live self-contradiction: `external_services.md:554`
   states *"that interface is no longer a shared private module for any service a stack builds, and this
   sentence said 'the shared `ai` library' until M257x iter-115"*, and `jobsimulation.md:168` states
   *"one `ai.AI` interface that is **in-tree, not a shared module**"* — both asserted as live, while
   `askengine.md` asserts the retracted phrase at two sites. Two anchors, one predicate.

### r29-A B2 | `corpus/services/askengine.md:104` | UPHELD | IN-SCOPE | PREDICATE: app's Asynq worker pools and Redis subscriber servers are constructed at `main.go:1438` onward.

   evidence: I extracted `main.go` at `app` `ad9f3c49` (the ref pinned five lines above, at `askengine.md:89`)
   and read it. `:1436-1443` is a prose comment inside the `MESSENGER_ENABLED` rationale; **`:1438` is the
   fragment `// enough: merely constructing this server and calling Subscribe() attaches app to`** — no
   construct at all. The constructs the sentence names sit *before* it: `jobsimwiring.StartWorkers` `:729`,
   `appWorker := worker.NewServer(redisAddr)` `:767`, `skillerWorker` `:827`, `cbWorker` `:1009`,
   `cmsWorker` `:1209`, and app's own `subServer := pubsub.NewSubscriberServer(` at **`:1376`** (under the
   `// PubSub subscribers` label at `:1375`). The only `NewSubscriberServer` after `:1438` is the
   messenger-only one at `:1450`. So "the Asynq pools … (`:1438` onward)" is false for all five pools, and
   "the Redis subscribers (`:1438` onward)" excludes the only unconditional subscriber server. Positive
   control on the same derivation: the sibling anchors in the same sentence **do** resolve — `:1295` is
   `mux := http.NewServeMux()` under `// RPC`, and `:1361` is `// Meta HTTP Server (healthcheck, _meta, _asynq)`
   directly above `metaServer := meta.NewServer(...)` at `:1362`. Noted, and it does not change the verdict:
   the load-bearing proposition (nothing after the `return` at `:470` runs — I confirmed `:467-471` is the
   bedrock block and `func main()` opens at `:229` with no intervening `func`) is **true**. The false
   proposition is the locational one.

### r30-A B1 | `corpus/architecture/external_services.md:727` | UPHELD | IN-SCOPE | PREDICATE: `anthropos-agent-eu` returns 0 hits across all 15 clone trees and 0 on a gitignore-blind filesystem grep.

   evidence: I re-derived the census the sentence invokes — `find stack-demo -name .git -maxdepth 4` →
   **15** trees — then ran `git grep -c anthropos-agent-eu HEAD` in each tree **at its own ref**. Fourteen
   return 0; `stack-demo/rosetta-extensions @ 09d06070` returns **7**, in tracked non-empty text files:
   `stack-core/tests/fixtures/claim_twin_iter48/red/07.md:3`,
   `stack-core/tests/fixtures/mechanical/{green,red}/corpus/architecture/external_services.md:652`,
   `stack-core/tests/fixtures/repair_leak/{pre,post}/corpus/architecture/external_services.md:652`,
   `stack-core/tests/fixtures/repair_leak/{pre,post}/corpus/architecture/ai_architecture.md:{152,168}`.
   The second clause fails on the same evidence: `/usr/bin/grep -rl anthropos-agent-eu stack-demo` (no
   `.gitignore` filtering) returns those same **7 files**, not 0. Per brief rule 4 I state the SET first:
   the sentence's own denominator is the 15-tree census, and one member of it is non-zero — the zero holds
   only over 14. What survives, and I record it: the **substantive** claim is TRUE — every one of the seven
   hits is a checked-in copy of this corpus's own *retracted* prose parked as a guard fixture, so
   `anthropos-agent-eu` genuinely appears in no platform source (`app` @ `ad9f3c49` returns 0). The defect is
   the published derivation, which a re-deriving reader cannot reproduce.

### r30-A B2 | `corpus/services/jobsimulation.md:164` | UPHELD | IN-SCOPE | PREDICATE: The corpus spells `app/internal/cms/directus/collections/jobsimulation.go` correctly at eight OTHER sites.

   evidence: I enumerated the set myself with `/usr/bin/grep -rn` over `corpus/` (bypassing the
   `ugrep --ignore-files` shell wrapper, so nothing tracked is hidden), and I state the cardinality first —
   the exact path `app/internal/cms/directus/collections/jobsimulation.go` occurs on **7** lines:
   `ai_architecture.md:61,221,222,264,277`, `external_services.md:624`, and the citing sentence itself
   (`jobsimulation.md:164`) → **6 other**. Relaxing the predicate to allow the repo-relative spelling
   (`internal/cms/directus/collections/jobsimulation.go`, with or without the `app/` prefix) gives **8 total /
   7 other** (the eighth is `external_services.md:566`). Relaxing to the directory
   `app/internal/cms/directus` gives 15/14. **No reading yields eight *other* sites** — 8 is reachable only
   as a *total*, which is exactly what the word "other" forbids. Not date drift: at `cec0ddb`, the commit
   that wrote the sentence, the exact-path count was also 7 total (`ai_architecture.md` 5 +
   `external_services.md` 1 + `jobsimulation.md` 1) → 6 other. The adjacent clause *"this was the sole
   outlier"* does hold and is not part of this verdict.

### r29-B B1 | `corpus/services/cms.md:153-154` | UPHELD | IN-SCOPE | PREDICATE: The studio tree's only `.docx` mentions are two doc files, and nothing in it depends on python-docx.

   evidence: I grepped the nested studio checkouts **at their own ref**, never through the host:
   `git -C stack-demo/app/studio grep -cil docx aeec036a` returns **5 files**, not 2 —
   `agents/simulation/guidelines.py`, `knowledge/development/asset-examples/README.md` (the two the doc
   names), plus `.../scenario-simulation/internal_localization.json`, `.../simulation.json`, and
   **`tools/any2pdf.py`**. Identical set in `stack-demo/cms/studio @ aeec036a`. The strictest possible
   reading — the literal `.docx` with the dot — still refutes it: `tools/any2pdf.py:664`
   (`'.docx': convert_docx_to_pdf,`), `:742`, `:752` (`if ext in ['.docx', '.doc']:`). And the "no
   dependency" half is refuted by the same file, which I opened: `:38` `'docx': 'python-docx',  # For
   reading DOCX files` inside `check_dependencies()`'s `required_packages` dict, `:82` `import docx`, `:85`
   `subprocess.check_call([sys.executable, "-m", "pip", "install", "python-docx"])` on `ImportError`, `:95`
   `doc = docx.Document(input_path)`; its module docstring `:5` names DOCX first among the formats it
   converts. It is live code, not a fossil: `tools/r3.py:139` is
   `scripts = ["any2pdf.py", "pdf2md.py", "md2cleanMd.py"]` inside `check_dependencies()` — the same `r3.py`
   this document cites twelve lines later. The *surrounding* claim is correct and I confirmed it:
   `requirements.txt` is verbatim 9 packages (openai, anthropic, rich, pyyaml, requests, jinja2, mistralai,
   pytest, pytest-asyncio) and python-docx is absent. It is the parenthetical's generalisation from
   `requirements.txt` to "the tree" that is false — and it is stated in the corpus's own most-fenced idiom
   ("the only X in the tree").

### r29-B B2 | `corpus/services/ai-readiness.md:55-58` | UPHELD | IN-SCOPE | PREDICATE: ai-readiness.md's `✅ CORRECTED M219` blockquote is `:476-496` and the `⚠⚠ M51 iter-08/09` block opens at `:498`.

   evidence: I located the named constructs in the file myself. At corpus HEAD `6b0cb96d`
   (content-identical to the file's last commit `9d31cf1`, tree clean): the `> **✅ CORRECTED M219 …`
   blockquote **opens at `:512`** and its closing line (*"relying on either number; do not re-derive them
   from prose."*) is **`:536`**; the `**⚠⚠ M51 iter-08/09 …` block opens at **`:538`**; the quoted
   parenthetical *"(now `aireadiness/readiness.go`, formerly `workforce/ai_readiness.go:512`)"* is at
   **`:540`**. What actually sits at the cited lines: `:476` is the tail of the `user_skill_evidences`
   correction bullet, `:496` is prose inside the M51 strategy sentence, `:498` is *"⚠️ **That premise was
   refuted at M219: the recompute takes 2.09 s.**"*, `:500` is the `ops/demo/stories-spec.md:599` link.
   I then tested the brief's rule-7 historical-anchor defence and it **fails**: `git show
   9d31cf1^:corpus/services/ai-readiness.md` puts the blockquote at `:476-496`, the M51 block at `:498` and
   the parenthetical at `:500` — all exact — and `9d31cf1` (*"fix(M257x/115): repair ai-readiness … the
   third-generation pin is RETIRED, not re-derived"*, `47 insertions(+), 7 deletions(-)`, net +40, every
   inserted line above the anchors) is the commit that both moved them **and** published the pre-move
   numbers. So `:498` and `:500` — stated in the present tense, and offered as the value the retired pin
   *would* have taken (*"The pin is deleted rather than re-derived to `:498`"*) — were already wrong in the
   commit that wrote them, not merely rotted afterwards. That is not a record of where something once was.
   The *decision* the passage takes (retire the pin, name the construct) is correct and is not what I uphold.

### r29-B B3 | `corpus/services/cms.md:216` | UPHELD | IN-SCOPE | PREDICATE: The **Data** bullet of `cms.md`'s merge banner is at `:44-47`.

   evidence: I read the banner. At HEAD `:44` is `> Where everything went:`, `:45` is `>`, and `:46-47` are
   the **Domain** bullet (`app/internal/cms/` … *"wired from `app/internal/cms/wiring.go`"*) — which says
   nothing about a schema and cannot be what a sentence about *"the legacy `cms` schema is
   non-authoritative"* is asserting consistency with. The **Data** bullet (`similarities`,
   `similarity_categories`, `similarity_features`, `similarity_skills`, `studio_documents`, `studio_tasks`
   *"re-created in the `public` schema"* by `20260724132049_cms_data_model.sql`, *"The old `cms` DB schema is
   legacy — no longer authoritative"*) is **`:48-51`**. Cause re-derived, not taken from the seat: at
   `f8be5a1` the Data bullet **was** `:44-47`; `b4bdbfc` (*"fix(M257x/115): repair P10 (11 sites, 5 files)
   and P20"*) inserted 4 lines above it and left the cross-reference un-repointed. This is a live
   navigational pointer ("the **Data** bullet, :44-47 above"), not a record of a prior audit, so rule 7 does
   not shelter it.

### r29-B B4 | `corpus/services/cms.md:240` | UPHELD | IN-SCOPE | PREDICATE: The **Studio** bullet of `cms.md`'s merge banner is at `:70-71`.

   evidence: At HEAD `:70-71` are the first two lines of the **Events** bullet (*"`app` owns the
   `CMS_STREAM` subscriber. The folded similarity re-index + Studio handlers are merged onto app's existing
   CMS subscriber via `.AddHandler(...)`"*). The **Studio** bullet — *"the Python `anthropos-studio-room`
   project is now pulled into the **`app`** image via the CI `additional_repo` mechanism (app v1.360.1)"*,
   which is precisely the `additional_repo` evidence `:238-240` is routing the reader to — is **`:75-76`**.
   Same cause, independently re-derived: at `f8be5a1` the Studio bullet was `:70-71`; `b4bdbfc` shifted it
   by 4 and left the pointer. Distinct proposition from B3 (different bullet, different target), so it does
   **not** collapse onto it.

### r29-B B5 | `CLAUDE.md:275` | UPHELD | OUT-OF-SCOPE | PREDICATE: The `next-web-app` frontend monorepo is a Next.js **15** monorepo.

   evidence: `git -C stack-demo/next-web-app grep -n '"next":' 8297c684` → `"next": "~16.2.12"` in all four
   app packages plus the shared UI package (`apps/web:46`, `apps/hiring:45`, `apps/integration:28`,
   `apps/maintenance:9`, `packages/ui:65`). `CLAUDE.md:275` reads *"**Frontend Applications**: Next.js 15
   monorepo on Vercel"*, and `corpus/services/README.md:55` — in scope, and correct — reads *"The Next.js
   **16** monorepo on Vercel"*. The claim is false and the seat's disclosure is right, but the anchor is
   `CLAUDE.md`, outside `corpus/services/**` and `corpus/architecture/**`, so per brief rule 8 it does not
   enter `N` or `P`. Recorded so it is not lost.

### r30-B B1 | `corpus/services/ai-readiness.md:55-58` | UPHELD | IN-SCOPE | PREDICATE: ai-readiness.md's `✅ CORRECTED M219` blockquote is `:476-496` and the `⚠⚠ M51 iter-08/09` block opens at `:498`.

   evidence: same anchor and same proposition as r29-B B1's sibling booking above — collapses onto **P6**.
   Re-derived once, stated once above (`:512`/`:536`/`:538`/`:540` measured; `9d31cf1^` exact; `9d31cf1`
   both moved them and published the old numbers). This seat additionally names `:496`'s actual content
   and I confirmed it: `:496` is *"frozen `ai_readiness_snapshots`), after iters 03→06 falsified the
   active-signals path — on the premise that the"*, i.e. mid-sentence prose, not a blockquote terminator.
   One predicate, one anchor block, two seats.

### r30-B B2 | `corpus/services/cms.md:216` | UPHELD | IN-SCOPE | PREDICATE: The **Data** bullet of `cms.md`'s merge banner is at `:44-47`.

   evidence: identical anchor and proposition to r29-B B3 — collapses onto **P7**. Measured above:
   `:44` = `> Where everything went:`, `:46-47` = the **Domain** bullet, Data bullet = `:48-51`.

### r30-B B3 | `corpus/services/cms.md:240` | UPHELD | IN-SCOPE | PREDICATE: The **Studio** bullet of `cms.md`'s merge banner is at `:70-71`.

   evidence: identical anchor and proposition to r29-B B4 — collapses onto **P8**. Measured above:
   `:70-71` = the **Events** bullet's first two lines, Studio bullet = `:75-76`.

---

## Things I checked and did NOT convict (so a later pass does not re-litigate)

These are not verdicts — no seat booked them as blockers — but they are the places where a flat re-read
would have produced a false positive, and I want the reasoning on the record:

- **`askengine.md:104`'s load-bearing claim.** The `return` at `main.go:470` really does sit one level
  inside `func main()` (`:229`), with no intervening `func` declaration up to `:467` (I enumerated:
  `^func ` occurs at `:203`, `:211`, `:229` only). Everything named in that sentence *is* after it. Only
  the `:1438` anchor is wrong; the paragraph's conclusion is sound.
- **`external_services.md:727`'s substantive LiveKit claim.** `anthropos-agent-eu` genuinely appears in no
  platform source. I upheld the derivation, not the conclusion.
- **The `ai`-fold retractions at `external_services.md:554` and `jobsimulation.md:168`.** Both are correct
  prose and both are *retractions* (brief rule 5) — they are what makes `askengine.md` a live
  self-contradiction rather than merely stale.
- **`jobsimulation.md:164`'s *"this was the sole outlier"* clause.** Holds; `external_services.md:566`
  writes the path repo-relative but resolvable, which is a different convention, not the outlier class.
- **`cms.md:149-152`'s python-docx-not-in-requirements.txt claim.** True in both studio checkouts.
- **The rule-7 historical-anchor defence, applied to all three same-file pin bookings** (ai-readiness
  `:55-58`, cms `:216`, `:240`). It fails for each: none of the three is framed as *"what a prior audit
  found at line N"* — they are live navigational pointers and present-tense location statements — and in
  the ai-readiness case the numbers were already false in the commit that published them.

## PREDICATE ROLL-UP

```
P1 | The `ai.AI` interface app's ask-engine uses is a shared `ai` library (a private Go module).                          | anchors: r29-A B1 @ corpus/services/askengine.md:81, r29-A B1 @ corpus/services/askengine.md:113
P2 | app's Asynq worker pools and Redis subscriber servers are constructed at `main.go:1438` onward.                      | anchors: r29-A B2 @ corpus/services/askengine.md:104
P3 | `anthropos-agent-eu` returns 0 hits across all 15 clone trees and 0 on a gitignore-blind filesystem grep.            | anchors: r30-A B1 @ corpus/architecture/external_services.md:727
P4 | The corpus spells `app/internal/cms/directus/collections/jobsimulation.go` correctly at eight OTHER sites.           | anchors: r30-A B2 @ corpus/services/jobsimulation.md:164
P5 | The studio tree's only `.docx` mentions are two doc files, and nothing in it depends on python-docx.                 | anchors: r29-B B1 @ corpus/services/cms.md:153-154
P6 | ai-readiness.md's `✅ CORRECTED M219` blockquote is `:476-496` and the `⚠⚠ M51 iter-08/09` block opens at `:498`.    | anchors: r29-B B2 @ corpus/services/ai-readiness.md:55-58, r30-B B1 @ corpus/services/ai-readiness.md:55-58
P7 | The **Data** bullet of `cms.md`'s merge banner is at `:44-47`.                                                       | anchors: r29-B B3 @ corpus/services/cms.md:216, r30-B B2 @ corpus/services/cms.md:216
P8 | The **Studio** bullet of `cms.md`'s merge banner is at `:70-71`.                                                     | anchors: r29-B B4 @ corpus/services/cms.md:240, r30-B B3 @ corpus/services/cms.md:240
```

Out-of-scope upheld, excluded from `N` and `P` (recorded, not counted):
`X1 | The next-web-app frontend monorepo is a Next.js 15 monorepo. | anchor: r29-B B5 @ CLAUDE.md:275`

Collapse summary: 12 bookings → 11 in-scope upheld anchors → **8** distinct in-scope predicates.
Three collapses, all cross-seat and all at the *same* anchor (r30-B B1↔r29-B B2, r30-B B2↔r29-B B3,
r30-B B3↔r29-B B4 — seat B read the same two files twice). One predicate (P1) carries two anchors
inside a single booking. No two anchors were collapsed on resemblance alone: P6, P7 and P8 are all
"a same-file pointer names the wrong construct" and are deliberately kept **separate**, because they are
three different false propositions about three different targets.

BOOKED=12 UPHELD=12 REJECTED=0 IN-SCOPE-UPHELD-BLOCKERS=11 DISTINCT-IN-SCOPE-PREDICATES=8
