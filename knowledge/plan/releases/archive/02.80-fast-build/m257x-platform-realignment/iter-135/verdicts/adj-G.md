# Adjudication G — seats r33-G, r34-G (M257x iter-131, re-adjudicated at iter-135)

**Scope line.** Two seat reports, both authored by Seat G (readings #33 and #34), adjudicated
independently against ground truth. **12 claimed BLOCKERs** (4 in `r33-G`, 8 in `r34-G`), collapsing to
**8 distinct propositions** — four propositions are booked twice, once per reading. I read only this
adjudicator's brief and the two assigned seat reports from `knowledge/plan/**`; no other iter dir, no
`progress.md`, no `decisions.md`, no prior adjudication, no iter-132/133/134 material. Every citation
below was opened by me at the ref named. Read-only throughout; this file is my only write.

**Repair status.** I diffed all four anchor files between the sealed pre-registration commit
`a532493` (the state the seats read) and HEAD. `ai_architecture.md`, `next-web-app.md` and
`secrets-spec.md` are **byte-identical**. `org-repos.md` changed in exactly **one** line — `:68`, the
`infrastructure` row — which no seat anchors on. **No claim in my set is affected by post-seat repair,
so no `UPHELD (since-repaired)` verdicts arise.** Line counts are unchanged (466 / 354 / 188 / 454), so
every seat anchor still resolves to the text it quoted.

---

## Counts

| metric | n |
|---|---|
| **CLAIMED (blockers)** | **12** |
| **UPHELD** | **10** |
| **REJECTED** | **2** |
| — of class `wrong-tree` | **0** |
| — of class `misread` | 0 |
| — of class `true-at-its-ref` | 0 |
| — of class `retraction-not-contradiction` | 0 |
| — of class `minor-not-blocker` | **2** |
| — of class `not-in-scope` | 0 |
| **CANNOT-SETTLE** | **0** |
| **DISTINCT-PREDICATES-IN-MY-SET** | **7** |
| **upheld rate** | **10/12 = 83.3 %** |

---

## Verdict table

| seat | B# | anchor | verdict | rejection class | predicate (if upheld) | class | multi-pin | repair-induced (sha) |
|---|---|---|---|---|---|---|---|---|
| r33-G | B1 | `corpus/architecture/org-repos.md:227` | **UPHELD** | — | P1 | platform-drift | yes | **yes — `3cd96f2`** (`iter(M257x/123)`) |
| r33-G | B2 | `corpus/architecture/org-repos.md:370` | **UPHELD** | — | P2 | intra-corpus-citation | no | **yes — `3cd96f2`** (`iter(M257x/123)`) |
| r33-G | B3 | `corpus/architecture/org-repos.md:43` | **UPHELD** | — | P3 | self-contradiction | yes | **yes — `3cd96f2`** (`iter(M257x/123)`) |
| r33-G | B4 | `corpus/architecture/ai_architecture.md:40` | **REJECTED** | `minor-not-blocker` | — | — | yes | no (`f075ba9`) |
| r34-G | B1 | `corpus/architecture/ai_architecture.md:110-111` | **UPHELD** | — | P4 | platform-drift | yes | no (`d791487`) |
| r34-G | B2 | `corpus/services/next-web-app.md:17` | **UPHELD** | — | P5 | self-contradiction | no | no (`cd16967`, iter-102) |
| r34-G | B3 | `corpus/services/next-web-app.md:186` | **UPHELD** | — | P6 | self-contradiction | no | no (`328ece5`, iter-81) |
| r34-G | B4 | `corpus/architecture/org-repos.md:43-45` | **UPHELD** | — | P3 *(dup of r33 B3)* | self-contradiction | yes | **yes — `3cd96f2`** |
| r34-G | B5 | `corpus/architecture/org-repos.md:370-371` | **UPHELD** | — | P2 *(dup of r33 B2)* | intra-corpus-citation | no | **yes — `3cd96f2`** |
| r34-G | B6 | `corpus/architecture/org-repos.md:226-228` | **UPHELD** | — | P1 *(dup of r33 B1)* | platform-drift | yes | **yes — `3cd96f2`** |
| r34-G | B7 | `corpus/architecture/ai_architecture.md:39-41` | **REJECTED** | `minor-not-blocker` | — | — | yes | no (`f075ba9`) |
| r34-G | B8 | `corpus/architecture/ai_architecture.md:224` | **UPHELD** | — | P7 | platform-drift | yes | no (`f075ba9`) |

---

## Upheld predicates, deduplicated within my assignment

| id | predicate (one line, anchor-independent) | anchors | class |
|---|---|---|---|
| **P1** | `JUDGE0_BASE_URL` is injected into the app task definition at `app/terraform/main.tf:638-639` | `corpus/architecture/org-repos.md:227` | platform-drift |
| **P2** | `secrets-spec.md:309` is the site at which the corpus borrows `hyper-studio`'s `.env.example` as the template for `app`'s five AWS Bedrock genes | `corpus/architecture/org-repos.md:370` (propagated to `corpus/architecture/platform-migration-status.md:161`) | intra-corpus-citation |
| **P3** | The corpus documents no observability tier at all — `git grep -i grafana -- corpus/ CLAUDE.md` returns 0 files | `corpus/architecture/org-repos.md:43` | self-contradiction |
| **P4** | Mistral is a provider spanned by the `ai.AI` interface, and the fold left what that interface provides unchanged | `corpus/architecture/ai_architecture.md:111` | platform-drift |
| **P5** | `apps/web` is the only frontend in platform compose | `corpus/services/next-web-app.md:17` | self-contradiction |
| **P6** | The Cosmo/WunderGraph federated router is prod-only — i.e. it still runs in production | `corpus/services/next-web-app.md:186` (**and the unflagged sibling `corpus/architecture/external_services.md:368`**) | self-contradiction |
| **P7** | The voice engine is selectable **per simulation** in the CMS | `corpus/architecture/ai_architecture.md:224` | platform-drift |

**DISTINCT-PREDICATES-IN-MY-SET = 7.**

---

## Upholds, with the evidence I opened

### P1 — `JUDGE0_BASE_URL` at `app/terraform/main.tf:638-639` (r33 B1 / r34 B6)

The corpus sentence (`org-repos.md:226-228`) names no ref for the `app` hop, so it grades at the
checkout `ad9f3c498`. At that ref `git grep -n JUDGE0 -- terraform/` returns **exactly two** hits:
`terraform/main.tf:531` `"name": "JUDGE0_BASE_URL",` (`:532` `"value": "${var.judge0_base_url}"`) and
`:693` `"name": "JUDGE0_API_KEY",`. Lines `:637-639` are the **`ELEVENLABS_WEBHOOK_SECRET`** block —
`:638` is `"valueFrom": "${aws_ssm_parameter.elevenlabs_webhook_secret.arn}"`, `:639` a closing `},`.

I ruled out `true-at-its-ref` exhaustively rather than by spot check. I swept **all 105 commits** that
touch `terraform/main.tf` and tested lines 638–639 at each: `JUDGE0_BASE_URL` appears there at **zero**
of them. The single hit at that line pair in the whole history is `25ce02f65`, where it is
`JUDGE0_API_KEY` — a different variable. Positive control in the same pass: **18 of the 105** commits
contain `JUDGE0_BASE_URL` *somewhere* in the file, so the zero is a real zero and not a broken pipeline.
At the four refs the seat named it resolves to `b948604f:329`, `2035f9a40:531`, `9d00a313:374`,
`ad9f3c498:531` — every one of the seat's reported line numbers is exact.

Materiality: `:638-639` is not merely a wrong line, it is a wrong **injection mechanism** — a
`valueFrom` SSM secret rather than a plain `"value"` env entry. A reader following the chain to learn
how the opaque tfvar reaches the app is led to conclude the Judge0 URL arrives as a secret. It does not.
The other end of the chain verifies exactly: `internal/jobsimwiring/wiring.go:123` is
`runnerManager := jsrunner.NewRunnerManager(getenv("JUDGE0_API_KEY"), getenv("JUDGE0_BASE_URL"))`.

### P2 — `secrets-spec.md:309` for the hyper-studio template (r33 B2 / r34 B5)

`corpus/ops/secrets-spec.md:309` is a table row reading
`| platform/ANTHROPIC_API_KEY · platform/OPENAI_API_KEY · platform/LIVEKIT_API_KEY ·
platform/LIVEKIT_API_SECRET | standard | required | key-present, nonempty |` — `platform`-repo secrets,
in a different table, in a different section. No hyper-studio, no AWS, no Bedrock. The claim's real site
is **`:344`**: *"**The 5 genes** (all on the `app` repo, target `app/.env`; the
`../hyper-studio/.env.example` template):"*, under the heading *"The Bedrock cred class for app"* at
`:333`. `grep -n hyper-studio corpus/ops/secrets-spec.md` returns exactly two lines, `:344` and `:351`.

This clears the blocker bar because the citation *is* the evidence: the sentence's rhetorical payload is
**"The corpus borrows a file from a repo it never mentions"**, and a reader who opens `:309` to check it
finds a table with no hyper-studio in it and concludes the assertion is unsupported. The pin also
propagated — `corpus/architecture/platform-migration-status.md:161` repeats *"`secrets-spec.md:309`
already borrows its `.env.example`"* — so it is the cross-file drift class, not an isolated typo.

### P3 — "the corpus documents no observability tier at all" (r33 B3 / r34 B4)

I ran the document's own printed command verbatim. `git grep -il grafana -- corpus/ CLAUDE.md` returns
**3 files** — `corpus/architecture/org-repos.md`, `corpus/ops/README.md`,
`corpus/ops/observability.md` — at HEAD **and** at the seal `a532493`. Positive control in the same
pass: `git grep -il directus -- corpus/ CLAUDE.md` → 55 files. `corpus/ops/observability.md` exists
(9,034 bytes) and is the observability tier's documentation — the file this very sentence links to two
clauses later.

Upheld, but see my framing disagreement below: I uphold this on **stale present tense plus a printed
derivation that does not reproduce**, not on the seats' stronger "self-refuting universal quantifier"
analogy, which I think overstates it. The decisive evidence is that the corpus states this identical
fact correctly in two sibling places — `corpus/ops/README.md` (*"documented nowhere **until
2026-08-07**"*) and `corpus/ops/observability.md` (*"**returned** 0 files"*) — both past-tensed and
date-scoped. `org-repos.md:42-43` is the un-tensed residue of the same repair, and it prints a live,
undated, reproducible-looking command that returns 3 where it claims 0.

### P4 — Mistral on the `ai.AI` interface (r34 B1)

`corpus/architecture/ai_architecture.md:110-111` says *"What the interface provides is **unchanged**: —
A single `ai.AI` interface across providers (OpenAI, Azure, Anthropic, Bedrock, **Mistral**)"*, with the
subject fixed by `:106-108` as the folded in-tree library at `app/internal/ai/`.

At `app` `ad9f3c498`, `git ls-tree -r internal/ai/` shows **exactly two** provider sub-packages —
`anthropic/` and `openai/`. There is no `internal/ai/mistral`. The constructors returning `ai.AI` are
exactly three: `openai/completion.go:24` `NewOpenAI`, `:53` `NewAzure`, `anthropic/completion.go:48`
`NewAnthropic(cfg *aws.Config, …)` (the Bedrock path). No Mistral constructor exists.

The platform states the removal in its own words, which is what makes *"unchanged"* false on its own
terms: `internal/cms/studio/mistralocr/mistralocr.go:1-11` reads *"It used to be `internal/ai/mistral`,
where it satisfied the nine-method `ai.AI` interface in order to expose exactly one working method… So
this is what it always was: upload a document, get a signed URL for it, run OCR on that URL… **No
interface**, no LLM methods, no tokenizer, and no panics."* It also contradicts this same file twice —
`:33` (*"Mistral is nowhere in this path… Every use of it in `app` is **OCR**, never generation"*) and
the *Mistral (EU)* row at `:93` (*"**OCR only** — the cms domain's Go client"*).

### P5 — "the only frontend in platform compose" (r34 B2)

At `platform` `0c91421df`, `docker-compose.yml` declares five services — `sentinel`(:5), `backend`(:28),
**`studio-desk`(:112)**, `next-web-app`(:143), `gotenberg`(:170). `studio-desk` publishes `"9000:9000"`
(:123) and **`"9100:9100"`** (:124), is built from `Dockerfile.dev` with `VITE_*` build args (:117-121),
and is wired at `:135` to `VITE_GRAPHQL_ENDPOINT=http://localhost:8082/graphql/query`. It is
unambiguously a second browser frontend in platform compose. Every line number the seat gave is exact.

The same document already says so, quoting the retracted phrase verbatim: `:121-124` — *"`apps/web` is
the only **`next-web-app`** app in platform compose — ⚠️ **not "the only frontend"**: `studio-desk` is a
second compose frontend with its own browser UI"*. This is **not** `retraction-not-contradiction`: the
retraction at `:121-124` is correct prose doing its job, and the defect is at `:17`, which still asserts
the retracted proposition in its own voice as a live claim. The repair reached one site and not the
other, and `:17` is the *Key Functions* summary bullet — the higher-propagation of the two.

### P6 — the router called "prod-only" (r34 B3)

`corpus/services/next-web-app.md:186` — *"the federated gateway (**prod-only** since `2adcf71`)"* — is
incompatible with `:80` of the same file: *"**in prod the router is destroyed — iter-124**"*. At least
one is wrong regardless of any external repo, so the self-contradiction is fully settled from text I
opened.

Which side the corpus holds is also settled from text I opened. `corpus/architecture/org-repos.md:144`
records `module.wundergraph_euwest1` as **deleted**, and `:391-397` states *"**RESOLVED AT M257x
iter-124, IN AKB'S FAVOUR: AKB was right and this corpus was wrong**"*. Two sibling sites have already
been repaired and each quotes this exact wording as the form they corrected:
`corpus/architecture/service_taxonomy.md:530` (*"corrected M257x iter-124, where this cell read 'Cosmo
Router (**prod only**)'"*) and `corpus/architecture/architecture_overview.md:73` (*"where this sentence
said 'it survives in production only'"*). `:186` is the unrepaired residue of a settled question.

I considered and rejected `true-at-its-ref`. *"prod-only **since** `2adcf71`"* is durative — it asserts
a state holding from that ref onward, i.e. now — not a state pinned at `2adcf71`. And the corpus itself
treats this wording as the retracted form rather than as a ref-pinned truth.

**Caveat, stated rather than laundered:** the underlying production fact rests on `infrastructure` @
`13c248e6`, which is in no clone set on this box, so I did not re-derive it. The blocker does not depend
on my doing so — it stands on the in-file contradiction plus the corpus's own adjudicated resolution.

**Additional in-scope site the seat did not flag:** `corpus/architecture/external_services.md:368`
carries the identical false clause — *"**prod-only** since platform `2adcf71` deleted it from local
dev"*. Same predicate P6, second anchor; it should be repaired in the same pass.

### P7 — voice engine "per simulation" (r34 B8)

The field is on the **sequence**, not the simulation. At `app` `ad9f3c498`,
`internal/cms/directus/collections/jobsimulation.go:911` declares
`VoiceEngine *SimulationVoiceEngine \`json:"voice_engine,omitempty"\``. I established which struct that
line sits in rather than taking the seat's word: the preceding `type … struct` declarations are
`JobSimulationCollection`(:25), `JobSimulation`(:724) and **`Sequence`(:834)**, the struct closes at
`:913`, and the next type declaration is `SequenceType`(:915) — so `:911` is unambiguously inside
`Sequence`, alongside `AIVendor`(:905) and `AIModel`(:906). `git grep -n VoiceEngine` over the file
returns no simulation-level field. The only read is `:1350`
`VoiceEngine: voiceEngineFromDirectus(seq.VoiceEngine)` — per `seq`, inside the per-sequence loop.

The same file contradicts `:224` nine lines later, in bold: `:233` — *"**Engine choice is per SEQUENCE,
from the CMS `voice_engine` field**"* — pinned in full to `:1594-1597` and described at `:238` as
*"Pinned in full, and to `:1597` rather than `:1600`, **deliberately** (M257x iter-115)"*. The enum it
describes as 4-member checks out exactly (`:1082-1085`). A corpus that bolds and deliberately pins this
distinction is treating it as load-bearing, and a seeder or content author acting on `:224` goes looking
for a `voice_engine` field on the simulation collection, where none exists.

---

## Rejections, with the evidence I opened

### r33 B4 and r34 B7 — `pdf2md.py:24` cited for `mistral-ocr-latest` → **REJECTED, `minor-not-blocker`**

**The factual finding is correct and I confirmed it independently.** At the nested repo's own ref
`aeec036a5` (`stack-demo/app/studio` is its own checkout — I read it there, not at the `app` host ref),
`tools/pdf2md.py:24` is `from mistralai import Mistral`, a bare SDK import.
`git grep -n mistral-ocr-latest aeec036a5` returns exactly one hit: `tools/pdf2md.py:127`
(`model="mistral-ocr-latest",`), inside `client.ocr.process(` at `:126`. Positive control: `grep -il
mistral` at that ref → 3 files, matching the corpus's own published derivation. So the pin is genuinely
103 lines off.

**I reject it as a BLOCKER on the seat's own stated distinguishing test**, which the seat articulated
and then declined to apply. r34 B7 writes: *"the cited line **is** Mistral-related, so a reader is not
sent to an unrelated subject the way B5/B6 send them."* That is exactly right, and it is the line
between P1/P2 and this one. P1 sends a reader into a different variable's block under a different
injection mechanism; P2 sends a reader into a different repo's secrets in a different table. Here the
file is correct, the subject (Mistral OCR in the in-image studio tree) is correct, the architectural
claim (*"a standalone CLI on neither the AI manager's path nor the generation pipeline's"*) is correct
and verified, and the literal is one grep away **in the file already named**. No false inference is
induced.

The seat's fallback argument is meta — that the same bullet, one clause earlier, retracts `:19` *"and
`:19` is a **doc-comment** line … not code"*, so the bullet violates its own standard. I opened that
retraction (`ai_architecture.md:36-38`) and it does not carry the weight placed on it: `:19`'s content
(*"It used to take aiKey and then IGNORE it"*) **contradicted** the assertion it was cited for, which is
why it warranted an in-text retraction. `:24`'s content contradicts nothing; it under-supports. Those
are different defect classes, and the corpus grades the second one as drift elsewhere.

For internal consistency: this seat booked pins 1–3 lines off as MINORs (`hiring.md:245`,
`ai_architecture.md:277`, `:325`). The distance here is larger, but distance is not the criterion the
seat itself proposed — subject displacement is, and there is none. Both readings booked this at
**medium** confidence, the lowest in the set, which is consistent with where I land.

---

## Cannot-settle

**None.** Every claimed blocker in my assignment resolved to a verdict on evidence I opened directly.

For the record, one upheld blocker (**P6**) has a component I could not re-derive — whether the
WunderGraph module is in fact destroyed in production, which lives in `infrastructure` @ `13c248e6`, a
repo in no clone set on this box and one I did not clone (this task is read-only). That component is
**not load-bearing for the verdict**: the blocker stands on the `:80` vs `:186` contradiction inside a
single corpus file plus the corpus's own recorded resolution and two already-repaired sibling sites.
Cloning `infrastructure` @ `13c248e6` and reading
`terraform/production/services.tf:509-517` would settle the production fact itself.

---

## Where I disagree with how the seats framed a predicate

**1. P3 (grafana / observability) — the analogy is overstated, and the predicate should be narrower.**
Both readings frame this as *"a universal quantifier falsified by its own sentence"* — the identical
defect the same file books at `:37`. I do not accept that framing. At `:37` the corpus books two
sentences that denied documentation existed *while themselves being the documentation*, with no
corrective pointer. Here the sentence **supplies the pointer in its own final clause** (*"See
`observability.md`"*), so a reader is not left holding the false belief. I also weighed a genuine
defense the seats raised and dismissed: the sentence sits inside a blockquote whose entire subject is
*"what the census got right / wrong"*, and the document globally date-stamps itself at `:22-23` to
2026-08-07 — the same day `observability.md` was created. That is a real mitigation.

I uphold anyway, on narrower and firmer ground: the `:22-23` date-stamp is explicitly scoped to *"every
per-repo fact below … re-derived from a clone"*, and this is not a per-repo clone fact — it is a claim
about **this repo**, checkable in place, in one second, in the present tense, with a printed derivation
offered as reproducible that returns 3 where it claims 0. The corpus states the same fact correctly and
past-tensed in two sibling files, which is what makes this the un-tensed residue of a half-completed
repair rather than a defensible historical note. **The predicate to book is the non-reproducing
present-tense census claim, not a self-refuting quantifier** — and a repairer who fixes it by adding
"until 2026-08-07" has fully closed it, which would not be true under the seats' framing.

**2. P7 (voice engine) — the seats under-state this one; I strengthen it.** Both readings hedge
(*"can be read loosely as 'configurable, in the CMS, for a simulation'"*, medium confidence). Having
established from the struct boundaries that `voice_engine` exists **only** on `Sequence` and that the
sole read is per-`seq`, I think the loose reading is unavailable: there is no simulation-level field to
be loosely describing. This is a clean platform-drift blocker, not a phrasing hedge.

**3. Predicate granularity for P1/P2 vs the pdf2md claim.** The seats treat all three as one family
("mis-anchored citation"). They are not one predicate and should not be scored as one. P1 and P2
displace the reader to a **different subject** (a different variable and injection mechanism; a
different repo's secrets in a different table) and each is the sole evidence for the sentence's payload.
The pdf2md pin under-supports a claim that is independently supported twice over in the same corpus
(`ai_architecture.md:42`'s own grep derivation and the `:93` table row). Collapsing them inflates the
count by one and, more importantly, hides that the corpus's citation discipline failed in two distinct
ways with two distinct repair costs.

**4. Provenance — the strongest signal in my set, which neither reading surfaces.** All **three**
`org-repos.md` blockers (P1, P2, P3) were last touched by the **same commit**: `3cd96f2`,
*"iter(M257x/123): cloning infrastructure settled four standing questions at once — and a service repo's
desired_count is not evidence"* (2026-08-07). That commit is inside the 120–130 repair window, so all
three are **repair-induced**. The four non-`org-repos.md` predicates (P4–P7) trace to much older commits
(`d791487`, `f075ba9`, `cd16967` = iter-102, `328ece5` = iter-81) and are **not** repair-induced. The
iteration the corpus credits with settling four standing questions introduced three new citation defects
in the same stroke — a 3:4 repair-induced-to-legacy split that is worth more to the milestone than any
individual verdict here.

---

## Seat-quality notes (not verdicts, but they bear on how much weight the reports carry)

- **The same seat quoted the same line two different ways.** For `secrets-spec.md:309`, `r33-G` quotes
  the `platform/ANTHROPIC_API_KEY · … LIVEKIT_API_SECRET` row — **correct**. `r34-G` quotes the
  `platform/ELEVENLABS_API_KEY · MISTRAL_API_KEY · …` row — that is line **`:311`**, not `:309`. Both
  readings reach the right verdict, but `r34-G` reaches it through a misquotation of its own key
  evidence. Verdict unaffected; confidence in `r34-G`'s quoting should be.
- **Both readings mis-cite the propagation site** as `platform-migration-status.md:160`. It is **`:161`**.
- **`r33-G`'s corpus-wide hyper-studio count is wrong**: it reports *"exactly 5 sites"*; I measure **7**
  (`secrets-spec.md:344,351`; `architecture/README.md:12`; `org-repos.md:79,356,370,371`). `r34-G`'s
  narrower claim — *"exactly two lines in `secrets-spec.md`, `:344` and `:351`"* — is exact.
- **Credit where due:** every one of the ~15 source line numbers I re-derived independently across both
  reports (terraform `:531`/`:532`/`:637-639`/`:693`; `wiring.go:123`; `pdf2md.py:24`/`:127`;
  `docker-compose.yml:112`/`:124`/`:135` and the five-service set; `jobsimulation.go:911`/`:1350`;
  `internal/ai/` constructor set; the four JUDGE0 candidate refs `:329`/`:531`/`:374`/`:531`) matched
  exactly. The seats' *measurements* are reliable; it is their *blocker-vs-minor thresholding* and one
  quotation that needed adjudication.
- **A pipeline trap I hit myself, worth recording.** My first history sweep used
  `git show "$r:terraform/main.tf"`, which zsh parsed as the `:t` (tail) modifier — producing
  `ad9f3c498erraform/main.tf`, a `fatal:` on stderr, and a **clean but meaningless zero**. It was caught
  only because I ran a positive control in the same pass and the control also returned empty. Re-run
  with `${r}:` the sweep gave the real answer over 105 commits with an 18/105 control. `r34-G` reports
  hitting the identical zsh modifier trap in its own pass and catching it the same way.

---

## Counts (brief format)

```
UPHELD=10 REJECTED=2 (of which wrong-tree=0; minor-not-blocker=2) CANNOT-SETTLE=0
DISTINCT-PREDICATES-IN-MY-SET=7
```
