# ADJUDICATOR 1 — seats A and B, readings #31 and #32

## Trees read, and at which refs

Every verdict below was **re-derived by opening the file myself**, never from a seat's quoted evidence.

| tree | ref I read at | used for |
|---|---|---|
| `rosetta` (corpus) | HEAD `c18d56b` — `corpus/services/{askengine,ai-readiness,cms,jobsimulation}.md`, `corpus/architecture/{external_services,shared_libraries,ai_architecture,dependency_map,README}.md` all last touched at or before `9d31cf1`/`b4bdbfc`/`cec0ddb`/`cd16967`, i.e. **unchanged since both seat readings** (`5b23559`, `04cbcfc`) | all six corpus anchors |
| `stack-demo/app` | `ad9f3c49` (== brief) — plus `git show` at the historical corpus refs `9d31cf1`, `9d31cf1^`, `a9f8ed45`, `b4bdbfc^` for born-wrong-vs-rotted tests | P1, P2, P3 |
| `stack-demo/platform` | `0c91421d` (== brief) — `docker-compose.yml`, `repos.yml` | P2 |

`stack-demo/rosetta-extensions` (pinned `09d06070`) and `.agentspace/rosetta-extensions` (authoring
`43049308`) were **not needed**: none of the ten bookings assigned to me is a tooling claim, so no
`wrong-tree` risk applies to this seat group.

`git status --porcelain` was **empty at my open and empty at my close**. Read-only; no fetch; no git
state change; the only file I created is this one.

---

## Verdicts

### A#31 B1 | `corpus/services/askengine.md:81` (+ `:113`) | UPHELD | IN-SCOPE | PREDICATE: The Ask Engine's embeddings come from the shared `ai` library, a private Go module.

   evidence: `corpus/services/askengine.md:80-81` reads *"**Downstream**: … the shared **`ai`** library
   (`CreateEmbeddings`, text-embedding-3-small, for RAG golden-example retrieval)"*, and `:113-114`
   *"an embeddings provider for the shared `ai` client."* Opened `stack-demo/app` at `ad9f3c49`:
   `go.mod` lists exactly five `anthropos-work` requires — `analytics-go:14`, `colony:15`, `proto:16`,
   `storage:17`, `taxonomy:18` — and `git grep -n "anthropos-work/ai" ad9f3c49 -- go.mod go.sum` returns
   **rc 1**. The three files of the very package the sentence describes import the in-tree path:
   `internal/web/backend/ask/{embed.go:8,examples.go:15,handler.go:20}` = `github.com/anthropos-work/app/internal/ai`,
   and the cited call is `embed.go:23` `aiClient.CreateEmbeddings(...)` on that client. The fold commit
   is real (`1e457fa70` *"refactor(ai): fold the ai library into app as internal/ai"*). The sentence
   carries **no ref and no historical fence**, so it grades at the checkout. Independently, this is a
   live corpus self-contradiction: `corpus/architecture/external_services.md:554` states *"that interface
   is **no longer a shared private module for any service a stack builds**, and this sentence said 'the
   shared `ai` library' until M257x iter-115"*; `corpus/architecture/shared_libraries.md:126` gives
   *Imported by* = **"No repo a stack builds"**; `corpus/services/jobsimulation.md:168` and
   `corpus/architecture/README.md:21` carry the same correction. I verified all four sibling sites myself.

### A#32 B1 | `corpus/services/askengine.md:81` (+ `:113`) | UPHELD | IN-SCOPE | PREDICATE: The Ask Engine's embeddings come from the shared `ai` library, a private Go module.

   evidence: same anchors, same re-derivation as A#31 B1 above — one predicate, booked by two seats.
   A#32's added arithmetic ("67 `.go` files import the in-tree path", "22 importers") was not
   load-bearing and I did not grade it; the predicate stands on the `go.mod` absence + the three
   `ask/` imports.

### A#31 B2 | `corpus/architecture/external_services.md:725` and `:741` | UPHELD | IN-SCOPE | PREDICATE: LiveKit and AWS Chime integrate with the Jobsimulation **service**, which still exists.

   evidence: I opened `external_services.md:719-741`. Both property tables carry, present tense and with
   no historical marker, `| **Integration** | Jobsimulation service |` — `:725` (LiveKit) and `:741`
   (AWS Chime SDK). Against ground truth at `platform 0c91421d`: `docker-compose.yml` declares **five**
   services (`sentinel:5`, `backend:28`, `studio-desk:112`, `next-web-app:143`, `gotenberg:170`) and
   `repos.yml` **four** entries (app, sentinel, next-web-app, studio-desk) — no `jobsimulation` in either.
   Both integrations live inside `app`: `git ls-tree ad9f3c49` gives
   `internal/jobsimulation/calls/livekit.go` and `internal/jobsimulation/recording/chime.go`. And it is a
   live corpus self-contradiction I re-derived rather than took: **this same file** says at `:173`
   *"There is no `cms`, `jobsimulation` or `roadrunner` container to start"* and deliberately writes
   *"Jobsimulation **domain**"* in its own AI table at `:562` and `:564`; `corpus/architecture/dependency_map.md:76`
   says outright *"there is no jobsimulation service to reach"*. Corpus-wide, present-tense
   *"Jobsimulation service"* survives at exactly these two sites plus the negating `dependency_map.md:76`
   and one archived-service doc (`chronos.md:214`). Not ref-discipline: the rows name no ref and no date.

### A#32 B2 | `corpus/architecture/external_services.md:581` | UPHELD | IN-SCOPE | PREDICATE: All of `app`'s AI vendor-selection mechanics live in `internal/jobsimulation/ai/ai.go`.

   evidence: `external_services.md:580-581` — *"The real mechanics, **all in `app/internal/jobsimulation/ai/ai.go`**:"*
   — scopes the whole *Routing: what is actually implemented* section, which `:554` sends the reader to
   *"before relying on it for a residency argument"*, and whose `:602-607` enumeration (*"**Four** things
   can send a request outside the EU"*) grounds items 1 and 2 solely in that file's anchors. I opened
   `stack-demo/app` at `ad9f3c49` and found a **second, live, independent** manager:
   `internal/skillerai/ai.go` — package doc `:1-9` (*"the second, independent AI manager … Azure EU/US +
   OpenAI clients, retry/throttle->OpenAI fallback, PostHog-driven Azure-US selection … **Do not alias it
   onto backend's existing AI manager.**"*), vendor consts `:57-61`, the **same** `flag_use_azure_us`
   PostHog lookup `:345-349` with the error-keeps-EU arm `:350-353`, `isThrottlingError` (HTTP 429)
   `:128-140`, the `throttled ⇒ vendor = Openai` override `:163-165`, and its own
   `openai.NewOpenAI(*openaiKey)` at `:108`. It is wired, not dead: `main.go:675`
   `skillerai.NewAIManager(...)` fed by `SKILLER_AZURE_OPENAI_KEY` / `SKILLER_AZURE_OPENAI_ENDPOINT_URL` /
   `SKILLER_OPENAI_KEY` (`:666`, `:669`, `:672`), feeding `embeddings`/`translation`/`skilltaxonomy`/
   `jobrole` at `:680-689`. I enumerated the set rather than trusting either side:
   `git grep -n "NewOpenAI(" ad9f3c49 -- '*.go'` returns exactly three non-test sites — the constructor
   `internal/ai/openai/completion.go:24` and the **two** managers. (I re-derived the importer count as
   **17** `.go` files, not the seat's 22; immaterial to the predicate.) Two sibling corpus files already
   name **both** sites — `shared_libraries.md:152-159` (*"the two real sites are
   `internal/jobsimulation/ai/ai.go` **and** `internal/skillerai/ai.go`"*) and `ai_architecture.md:114`
   — so this is also a live self-contradiction. One nuance I checked and it does not save the claim:
   `main.go:675` passes `nil, nil` for `azureKeyUs/azureEndpointUs` (signature `:78-86`), so skillerai's
   US-Azure arm is unarmed on this wiring — but the 429→direct-OpenAI lever is not, and the falsified
   proposition is the word **"all"**, which fails on the existence of the second manager either way.

### B#31 B1 | `corpus/services/ai-readiness.md:54-58` | UPHELD | IN-SCOPE | PREDICATE: In ai-readiness.md the ✅CORRECTED-M219 blockquote spans :476-496 and the ⚠⚠M51 block opens at :498.

   evidence: I re-derived the construct positions myself at corpus HEAD `c18d56b`
   (`grep -n` + `awk 'NR==N'`): `> **✅ CORRECTED M219 …**` opens at **`:512`** and closes at **`:536`**;
   the `⚠⚠ M51 iter-08/09` block opens at **`:538`**; the quoted parenthetical
   *"(now `aireadiness/readiness.go`, formerly `workforce/ai_readiness.go:512`)"* is at **`:540`**. The
   four lines the passage names hold something else entirely — `:476` *"row in the **progress** table,
   never a `user_skill_evidences` row…"*, `:496` *"frozen `ai_readiness_snapshots`), after iters 03→06
   falsified…"*, `:498` *"per-object-RPC class). ⚠️ **That premise was refuted at M219…**"*, `:500` the
   `[ops/demo/stories-spec.md:599]` link line. **This is not ref-discipline and not a historical anchor.**
   The passage names its own ref — *"re-derived at iter-115"* — and iter-115 IS the commit that shipped
   it (`git blame -L 54,58` → `9d31cf1b`, *"fix(M257x/115): … the third-generation pin is RETIRED"*). I
   graded it there: at **`9d31cf1`** (712 lines) the blockquote is already at `:512` / M51 at `:538` /
   the parenthetical at `:540`; at **`9d31cf1^`** (672 lines) `:476` opens the blockquote, `:496` closes
   it, `:498` opens M51 and `:500` is the parenthetical. So the numbers were measured pre-edit and
   published post-edit **in the same commit** — false at its own named ref, in the one paragraph whose
   entire subject is a pin that keeps rotting. Verbs are present tense (*"`:496` **is**"*, *"the block
   **opens at** `:498`"*).

### B#32 B1 | `corpus/services/ai-readiness.md:54-58` | UPHELD | IN-SCOPE | PREDICATE: In ai-readiness.md the ✅CORRECTED-M219 blockquote spans :476-496 and the ⚠⚠M51 block opens at :498.

   evidence: same anchor, same re-derivation as B#31 B1 — one predicate, booked by two seats. B#32's
   *"related, same cluster"* note on `:50-51` (`:459` merely opens the blockquote) is **not** a separate
   booking and I do not grade it as one; on its own it would fall to the historical/ref-discipline class,
   which is why the seat correctly declined to book it separately.

### B#31 B2 | `corpus/services/cms.md:240` | UPHELD | IN-SCOPE | PREDICATE: cms.md's **Studio** bullet is at `:70-71`.

   evidence: I read `cms.md:44-80` and `:236-241` at HEAD. `:240` says *"see the **Studio** bullet,
   `:70-71` in the banner at the top of this doc."* `:70-71` is the **Events** bullet (*"`app` owns the
   `CMS_STREAM` subscriber. The folded similarity re-index + Studio handlers are merged onto app's
   **existing** CMS subscriber via `.AddHandler(...)`"*). The **Studio** bullet is at **`:75-76`**
   (`:74` is **Caching**). The trap is real: `:70-71` contains the word *"Studio"*, so an
   anchor-existence check passes it. It is a live same-file navigation pointer, not a record of a past
   measurement, and it names no ref — so rule 7 does not shield it. Born correct and rotted: `git blame`
   puts the line in `a9f8ed45` (2026-08-06) and at `b4bdbfc^` the Studio bullet **was** at `:70-71`;
   `b4bdbfc` inserted 5 lines above and did not re-point.

### B#32 B3 | `corpus/services/cms.md:240` | UPHELD | IN-SCOPE | PREDICATE: cms.md's **Studio** bullet is at `:70-71`.

   evidence: same anchor, same re-derivation as B#31 B2 — one predicate, booked by two seats.

### B#31 B3 | `corpus/services/cms.md:215-216` | UPHELD | IN-SCOPE | PREDICATE: cms.md's **Data** bullet is at `:44-47`.

   evidence: `cms.md:215-216` — *"the cms tables were re-created there at cms-in-app v8.0 … **Consistent
   with the *Data* bullet, `:44-47` above**"*. Measured at HEAD: `:44` `> Where everything went:`,
   `:45` `>`, `:46-47` the **Domain** bullet (*"`app/internal/cms/` (directus, similarity, studio,
   library, importer/exporter, aivideo, contentread, jobsimimport, rpcsrv, worker, …), wired from
   `app/internal/cms/wiring.go`"*). The **Data** bullet — the one that actually states the `public`-schema
   re-creation and names `20260724132049_cms_data_model.sql` — is **`:48-51`**. The cited range contains
   none of the claim it is offered as corroboration for. Distinct from B#31 B2: different anchor,
   different construct, different offset (+4 vs +5), so one re-point does not fix both. Same rot history
   (`a9f8ed45` correct → `b4bdbfc` shifted; verified at `b4bdbfc^`, where `:44-47` **is** the Data bullet).

### B#32 B2 | `corpus/services/cms.md:216` | UPHELD | IN-SCOPE | PREDICATE: cms.md's **Data** bullet is at `:44-47`.

   evidence: same anchor, same re-derivation as B#31 B3 — one predicate, booked by two seats.

---

## PREDICATE ROLL-UP

```
P1 | The Ask Engine's embeddings come from the shared `ai` library, a private Go module. | anchors: A#31 B1 @ corpus/services/askengine.md:81, A#31 B1 @ corpus/services/askengine.md:113, A#32 B1 @ corpus/services/askengine.md:81, A#32 B1 @ corpus/services/askengine.md:113
P2 | LiveKit and AWS Chime integrate with the Jobsimulation *service*, which still exists. | anchors: A#31 B2 @ corpus/architecture/external_services.md:725, A#31 B2 @ corpus/architecture/external_services.md:741
P3 | All of `app`'s AI vendor-selection mechanics live in `internal/jobsimulation/ai/ai.go`. | anchors: A#32 B2 @ corpus/architecture/external_services.md:581
P4 | In ai-readiness.md the ✅CORRECTED-M219 blockquote spans :476-496 and the ⚠⚠M51 block opens at :498. | anchors: B#31 B1 @ corpus/services/ai-readiness.md:55-58, B#32 B1 @ corpus/services/ai-readiness.md:54-58
P5 | cms.md's **Studio** bullet is at `:70-71`. | anchors: B#31 B2 @ corpus/services/cms.md:240, B#32 B3 @ corpus/services/cms.md:240
P6 | cms.md's **Data** bullet is at `:44-47`. | anchors: B#31 B3 @ corpus/services/cms.md:215-216, B#32 B2 @ corpus/services/cms.md:216
```

Collapses: A#31 B1 + A#32 B1 → **P1**. B#31 B1 + B#32 B1 → **P4**. B#31 B2 + B#32 B3 → **P5**.
B#31 B3 + B#32 B2 → **P6**. A#31 B2 (**P2**) and A#32 B2 (**P3**) stand alone and do **not** collapse
onto each other: one is a nonexistent-service noun in two integration rows, the other a
single-file scoping assertion contradicted by a second live AI manager.

Rejections: **none**. No booking in this seat group fell to ref-discipline (no booked claim is pinned,
past-tense or dated in a way that saves it — I checked each: P1/P2/P3/P5/P6 name no ref at all, and P4
names *iter-115*, at which commit I re-derived it and it is still false), and none fell to `wrong-tree`
(no tooling claim was booked).

---

BOOKED=10 UPHELD=10 REJECTED=0 IN-SCOPE-UPHELD-BLOCKERS=10 DISTINCT-IN-SCOPE-PREDICATES=6
