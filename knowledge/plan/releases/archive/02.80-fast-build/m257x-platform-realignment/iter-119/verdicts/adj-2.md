# adj-2 — seats C and G (readings #31 and #32)

## Trees read, and at which refs

Every verdict below was **re-derived by opening the file myself**; no seat's quoted evidence was used as
proof. `export DEVELOPER_DIR=/Library/Developer/CommandLineTools`; cwd `/Users/marco/workspace/anthropos/rosetta`.
Read-only, **no fetch**, no git state change.

| tree | path | ref I read at |
|---|---|---|
| corpus (rosetta) | `corpus/**` | working tree @ `c18d56b`. **`git diff --stat 4d4530d..HEAD -- corpus/` is EMPTY**, so the corpus the seats graded (`4d4530d` / `5b23559`) is byte-identical to what I read — no line number moved under either reading |
| app | `stack-demo/app` | `ad9f3c49` (checkout = origin/main); also read `b948604f` and the fold commit `1e457fa70` where a claim named them |
| app/studio (nested, own checkout) | `stack-demo/app/studio` | `aeec036a` — grepped at **its own** ref, never at the host `app` HEAD |
| platform | `stack-demo/platform` | `0c91421d` |
| next-web-app | `stack-demo/next-web-app` | `8297c684` |
| studio-desk | `stack-demo/studio-desk` | `41ee3575` |
| messenger | `stack-demo/messenger` | **both** `fa47850d` (checkout) **and** `e9421c68` (origin/main) — the C31-B3 verdict turns on which |
| rosetta-extensions | `stack-demo/rosetta-extensions` (**pinned per-stack consumption clone**) | `09d06070` — the tree that settles "what the tooling does on a stack" (G32-B1) |

I did **not** need the authoring copy `.agentspace/rosetta-extensions`: the only rext booking in my set
(G32-B1) is a claim about what a stack's bring-up **does**, which the pinned clone settles. Naming the tree
explicitly per the `DEF-M257x-iter101-briefing-rext-tree` note.

**`git status --porcelain` at my open: EMPTY** (verified as the first command of the session).
**At my close:** the single line `?? knowledge/plan/releases/.../iter-119/verdicts/` — the untracked
directory holding this file and my peers'. No tracked file changed; I created/edited nothing else.

Hard bar honoured: under `knowledge/plan/**` I read only the adjudicator brief, my four assigned seat
reports, and this output file. No prior verdict, ledger, or answer key was opened, greped, or listed.

---

## Verdicts

### Seat C — reading #31

```
C31 B1 | corpus/architecture/ai_architecture.md:92, :94, :294-295 | UPHELD | IN-SCOPE | PREDICATE: ai_architecture.md's three bare self-citations (`:86`, `:87`, `:95`) name the constructs they point at.
   evidence: I opened corpus/architecture/ai_architecture.md at HEAD (corpus unchanged since 4d4530d) and read :80-120 and :255-320 directly.
     :86 = "The **Default client** column is the client the vendor const resolves to — *not* a position in a" (prose). The OpenAI table ROW is :91. Cited at :92 as "its OpenAI sibling at `:86`".
     :87 = "fallback order (there is none; see above)." (the prose sentence's tail). The Anthropic table ROW is :92. Cited at :294-295 as "the table row at `:87`, which is where the 4.6/Opus-4.8 omission lived".
     :95 = "| **Transcription** | GPT-4o Transcribe | Azure EU (US via `flag_use_azure_us`) |" — a table row saying nothing about the `ai` module. The fold proposition is :100 ("it is no longer a shared private module for any service a stack builds"), heading :98, go.mod evidence :106-109. Cited at :94 as "(see `:95` below)".
     All three are off by exactly +5 in the same direction; the fourth bare self-citation in the file (:114 → ":15-17", the ⚠️ blockquote) sits ABOVE the insertion point and resolves correctly — which is the control that this is drift, not my miscount. Two of the three sentences exist solely to route a reader to a construct, so the anchor is the entire payload.
     Not historical (rule 7): all three are live pointers in present tense, not records of where something once was.
```

```
C31 B2 | corpus/architecture/ai_architecture.md:40 | UPHELD | IN-SCOPE | PREDICATE: The `mistral-ocr-latest` literal is at `studio/tools/pdf2md.py:24`.
   evidence: The bullet names its own ref three lines below (":42 — `git -C app/studio grep -i mistral aeec036a`"), so it grades at stack-demo/app/studio @ aeec036a, its own checkout. There, `git show aeec036a:tools/pdf2md.py | awk` line 24 = `from mistralai import Mistral` — an import. `git grep -n "mistral-ocr-latest" aeec036a` returns EXACTLY ONE hit in the whole nested repo: `tools/pdf2md.py:127`.
     I resolved my hesitation upward for two reasons I re-derived myself: (a) the parenthetical shape `<file>:<line> (<construct>)` is used identically by the Go half of the very same bullet (":34 `markdownManager.go:30` — the constructor body `return &MarkdownManager{ocr: mistralocr.New(aiKey)}, nil`"), so the parenthetical names the construct AT the line; (b) that same bullet at :36-37 records M257x iter-115 repairing this exact class on the Go side (`:19` → `:30`, "`:19` is a **doc-comment** line … not code"). By the document's own just-applied standard an import line is not the OCR-model line.
     NB seat C's own second reading (#32) demoted this to MINOR; the demotion does not survive the file's own precedent, and the frozen instrument grades a wrong-construct anchor as a blocker.
```

```
C31 B3 | corpus/architecture/platform-migration-status.md:93 | REJECTED | — | PREDICATE (asserted, found TRUE): messenger's prod ECS service is scaled to zero at `terraform/main.tf:29` in an otherwise-intact 121-line module.
   evidence: I read messenger/terraform/main.tf at BOTH refs the brief's ground-truth table names.
     @ origin/main `e9421c68`: the file is **121 lines**; `:19-25` is verbatim the v9.0 comment the cell quotes ("app … has taken over this service's Redis consumer group … `terraform apply` runs UNTARGETED …"); `:25` is the cms-precedent line; `:27-28` is "The image and task definition stay declared: this is the rollback path."; `:29` is `service_desired_count          = 0`. **Every anchor in the cell is exact, word for word.**
     @ the local checkout `fa47850d`: 111 lines, `:19` = `service_desired_count = 1`, `:29` = `container_definitions = <<EOF`.
   class: ref-discipline — the booking survives only by adopting a STALE LOCAL CHECKOUT as the grading ref. The claim is a statement about the platform's production state; the platform's messenger repo is at `e9421c68`, which the brief's own ground-truth table names precisely because the clone is behind. Grading a correct, current description of prod against a mirror that has not been fetched is the mirror image of the ref-discipline class (older evidence used to refute, rather than newer), and yields the perverse result that an exact citation is "false". The seat itself booked at confidence LOW and wrote that "either verdict repairs the same way"; seat C's own second reading (#32) opened the same two refs and CLEARED it. Nothing in the map's promise ("Every claim is cited to a sha or a `file:line`") is violated — it is cited to a file:line, and the file:line resolves.
```

### Seat C — reading #32

```
C32 B1 | corpus/architecture/ai_architecture.md:92, :94, :295 | UPHELD | IN-SCOPE | PREDICATE: ai_architecture.md's three bare self-citations (`:86`, `:87`, `:95`) name the constructs they point at.
   evidence: same three anchors, same file, re-derived above under C31-B1 (I opened :80-120 and :255-320 and read every cited line). Collapses onto P1 — one predicate, two bookings, three anchors.
```

```
C32 B2 | corpus/architecture/ai_architecture.md:211 vs :220 | UPHELD | IN-SCOPE | PREDICATE: The voice engine is selectable per SIMULATION in CMS.
   evidence: I read ai_architecture.md:195-231 — :211 asserts "**Configuration**: Voice engine is selectable per simulation in CMS (`livekitgptrealtime`)" and :220 asserts, in bold, "**Engine choice is per SEQUENCE, from the CMS `voice_engine` field**". Both are live assertions; neither retracts the other, so rule 5's retraction carve-out does not apply.
     Ground truth picks the second: at app `ad9f3c49` I extracted internal/cms/directus/collections/jobsimulation.go and read it. `VoiceEngine *SimulationVoiceEngine \`json:"voice_engine,omitempty"\`` is at :911, and the last `^type … struct` at or above :911 is `:834 type Sequence struct {` (closing brace :913) — the field is on **Sequence**, alongside `AIVendor` :905 and `AIModel` :906, which the corpus itself correctly calls per-sequence. It is read per sequence inside the sequence loop: `:1350 VoiceEngine: voiceEngineFromDirectus(seq.VoiceEngine),` on the same `seq` appended at :1353. The enum is a 4-member `SimulationVoiceEngine` at :1081-1086.
     So :211 is a stale pre-correction line standing nine lines ABOVE the bolded correction — a reader scanning the *Active Engine* section meets the wrong one first.
```

```
C32 B3 | corpus/architecture/ai_architecture.md:264 | UPHELD | IN-SCOPE | PREDICATE: The Directus `AIModel` enum is `jobsimulation.go:983-990`.
   evidence: the sentence names no ref, so it grades at the checkout. I read internal/cms/directus/collections/jobsimulation.go @ app `ad9f3c49`: `type AIModel string` :977, const block `:979-991`, **11 members at :980-990**. The cited range `:983-990` opens four lines inside the block and drops the first three members — `Anthropic35SonnetAws :980`, `Anthropic37SonnetAws :981`, `Anthropic4SonnetAws :982`.
     I verified the SET before the arithmetic (rule 4): the enum's cardinality is 11, the cited range covers 8. The dropped members are load-bearing elsewhere in the same document — :287-293 makes `Anthropic37SonnetAWS…` and `Anthropic35Sonnet…` the two Anthropic fallback arms.
     Not line drift: at the older `b948604f` the block is `:979-991` with the same 11 members at :980-990, so `:983-990` was never the enum at any ref in the clone set.
```

### Seat G — reading #31

```
G31 B1 | corpus/architecture/shared_libraries.md:128-130, :137, :162-164 | REJECTED | — | PREDICATE (asserted, found TRUE at its pin): the `ai` module `github.com/anthropos-work/ai` @ v1.40.2 exposes a 9-method `ai.AI` and Anthropic/Mistral `panic` on the unimplemented methods.
   evidence: the block sits inside the `## ai` section whose property table, two lines above, states **Module = `github.com/anthropos-work/ai`** (:123) and **Version pin = `v1.40.2`** (:125), and whose `Imported by` row (:126) explicitly separates the module from `app`'s in-tree copy. Its subject is therefore the module at v1.40.2, not the fold.
     `app/internal/ai/module_import_guard_test.go:14-18` @ `ad9f3c49` states the package "was folded into app at tag **v1.40.2**", which makes the fold commit's tree the module's v1.40.2 content. I read it: `git show 1e457fa70:internal/ai/ai.go` declares a **9-method** `AI` interface INCLUDING `ChatCompletionStream` at :10 — exactly the nine names the corpus lists, in order. `git grep -n 'panic(' 1e457fa70 -- internal/ai/` returns `anthropic/completion.go:180` and `:189` (`panic("not implemented")` — the ChatCompletionStream/CreateSpeech pair) and `mistral/completion.go:40,44,48,52` (four = chat/embeddings/speech). `internal/ai/mistral/completion.go:25` is `func NewMistral(...)`; `internal/ai/openai/completion.go:25` is `func New(...)` and `:29` `func NewOpenAI(...)`. **Every enumerated claim holds verbatim at the pin.**
     The seat's contradiction argument does not hold either: :150-151's "since the fold `app/internal/ai/ai.go` is **21 lines**" is a measurement of a DIFFERENT artifact (I counted it: `ad9f3c49:internal/ai/ai.go` is exactly 21 lines / 8 methods; `1e457fa70`'s is 22 / 9), and it enumerates no methods, so no two passages assert incompatible things about the same artifact. The corpus records the divergence elsewhere in its own words — `ai_architecture.md:94` says the two TTS consts "belong to the **standalone** … module … they were dropped in the fold".
   class: ref-discipline — a pinned claim (`v1.40.2`, stated in the section's own property table) booked because newer evidence (`app`'s post-fold in-tree copy at `ad9f3c49`, diverged at `9048ce1b4`) contradicts it. Rule 1: a pin is a date, not an excuse; true at its named ref ⇒ TRUE, however stale. (The seat's own hesitation paragraph reaches the same finding and books anyway.)
```

```
G31 B2 | corpus/services/hiring.md:47-49 | UPHELD | IN-SCOPE | PREDICATE: `service_taxonomy.md:62` is the Tier-1 **Database** characteristic bullet.
   evidence: I opened both corpus files. `corpus/architecture/service_taxonomy.md:62` = `### Tier 1: Core Backend Services (Dockerized Go Microservices)` — a section heading containing none of the quoted text. :64 `**Characteristics**:`, :65 Language, :66 Deployment, :67 Communication, and the **Database** bullet carrying BOTH quoted fragments ("one schema, `public`, owned by `app`, which is the only repo with migrations" and "the `cms`, `jobsimulation` and `skillpath` schemas are legacy husks") is **:68**.
     The quoted text is accurate; the anchor names the wrong construct. Not historical (rule 7): the parenthetical `(`:62` — "…")` is a live pointer. The genuinely historical clause is the NEXT sentence ("This cited `service_taxonomy.md:52` until M257x iter-102 …"), which is correct prose — and which establishes that the file has now mis-pointed this one citation twice running.
```

```
G31 B3 | corpus/services/clerk-integration.md:40 | UPHELD | IN-SCOPE | PREDICATE: Clerk sign-in tokens are minted ONLY for app-native admin impersonation.
   evidence: I grepped each clone at the ref the ground-truth table names and then OPENED each site.
     1. `app` @ `ad9f3c49` — `internal/admin/impersonation/manager.go:101` `m.signInTokenCl.Create(ctx, &clerkSignInToken.CreateParams{…})`. The documented one; the "chosen over Enterprise-tier Actor Tokens" half is correct.
     2. `next-web-app` @ `8297c684` — `apps/web/src/app/api/dev/login-as/route.ts:79` `await client.signInTokens.createSignInToken({ userId: user.id, expiresInSeconds: 600 })`, redirecting to `/dev/accept?ticket=…` (:85-88).
     3. `studio-desk` @ `41ee3575` — `src/routes/dev.ts:83` `await clerkClient.signInTokens.createSignInToken({…})`, redirecting to `/dev-accept.html?ticket=…` (:89-92).
     Sites 2 and 3 are dev-gated (route.ts:33-35 hard-404s when `!DEV_LOGIN_ENABLED`), but they are real, shipped, present-tense consumers of the same Clerk Backend-API surface, and the corpus documents the harness elsewhere (`studio-desk.md`). The defect is the bolded universal: "we use sign-in tokens for X" would be an incomplete list (a MINOR); "**only** for X" is a false exclusivity claim on the page that bills itself as the single source of truth for the platform's Clerk surface — which is exactly what a Clerkenstein author reads to decide which Clerk surfaces the mock must reproduce, and the ticket→session exchange is the flow the demo cockpit depends on.
```

### Seat G — reading #32

```
G32 B1 | corpus/services/hiring.md:434-437 | UPHELD | IN-SCOPE | PREDICATE: The demo's hiring image bakes FOUR demo-patches.
   evidence: settled at the **pinned per-stack consumption clone** `stack-demo/rosetta-extensions` `09d06070` — the code a stack actually runs. I extracted `demo-stack/up-injected.sh` and read `build_frontend_hiring()` (declared :1082) end to end.
     SEVEN manifests are declared: `:1096` next-hiring-role-remap · `:1102` next-hiring-members-pagination · `:1111` next-web-studio-url · `:1112` next-web-public-website-url · `:1117` next-web-interview-flag-container · `:1118` next-web-interview-flag-result · `:1122` next-web-back-to-cockpit. All seven are folded into the cache key at `:1124-1126` (`next_web_patchset_fp` with seven arguments), all seven are REVERTED by the RETURN trap at `:1202`, and all seven are APPLIED — I read the apply arms: studio `:1220`, pubweb `:1223`, rolemap `:1242`, pagination `:1259`, interview-container `:1273`, interview-result `:1281`, back-to-cockpit `:1290`.
     The corpus sentence — "**Four** demo-patches on the hiring image make it land — 2 net-new … + the 2 chained shared `urls.ts` patches" — is short by three, and the two missing pairs are dated in the script's own comments (`[M232]` :1113, `[M249]` :1119).
     The cross-reference is also dead: the sentence sends the reader to `demopatch-spec.md` "§ the four hiring-image patches"; that file has **no** heading matching hiring at all (heading census run) and 0 occurrences of "four hiring"/"hiring-image patches".
     Corroborating, and re-derived rather than taken from the seat: the twin file was half-repaired and now contradicts itself — `demopatch-spec.md` says hiring's fingerprint is over **7** and tabulates all seven by name under an explicit "Corrected at v2.8 M255" blockquote, while a later paragraph still reads "**bakes FOUR patches**, not two … fenced by a **4-manifest patch-set fingerprint union**". `hiring.md` was not repaired at all. (That twin is `corpus/ops/**` — out of scope, and not booked here; it is evidence, not an anchor.)
     Materiality is stated in the code itself at `:1087-1094`: "a manifest added there and forgotten here is invisible to the cache key, which is the exact bug the fingerprint exists to kill." A doc that fixes the set at four is that stale-list defect one layer out.
```

```
G32 B2 | corpus/services/hiring.md:47-49 | UPHELD | IN-SCOPE | PREDICATE: `service_taxonomy.md:62` is the Tier-1 **Database** characteristic bullet.
   evidence: identical anchor to G31-B2, re-derived above (service_taxonomy.md :62 = the Tier-1 section heading; the quoted bullet is :68). Collapses onto P5 — one predicate, two bookings, one anchor.
```

---

## PREDICATE ROLL-UP

```
P1 | ai_architecture.md's three bare self-citations (`:86`, `:87`, `:95`) name the constructs they point at | anchors: C31 B1 @ corpus/architecture/ai_architecture.md:92, :94, :294-295; C32 B1 @ corpus/architecture/ai_architecture.md:92, :94, :295
P2 | The `mistral-ocr-latest` literal is at `studio/tools/pdf2md.py:24` | anchors: C31 B2 @ corpus/architecture/ai_architecture.md:40
P3 | The voice engine is selectable per SIMULATION in CMS | anchors: C32 B2 @ corpus/architecture/ai_architecture.md:211 (against :220)
P4 | The Directus `AIModel` enum is `jobsimulation.go:983-990` | anchors: C32 B3 @ corpus/architecture/ai_architecture.md:264
P5 | `service_taxonomy.md:62` is the Tier-1 Database characteristic bullet | anchors: G31 B2 @ corpus/services/hiring.md:47-49; G32 B2 @ corpus/services/hiring.md:47-49
P6 | Clerk sign-in tokens are minted only for app-native admin impersonation | anchors: G31 B3 @ corpus/services/clerk-integration.md:40
P7 | The demo's hiring image bakes four demo-patches | anchors: G32 B1 @ corpus/services/hiring.md:434-437
```

Dedup notes: P1 collapses two bookings across seat C's two readings onto ONE predicate carrying THREE
anchors (the seat's own note on the booking unit is correct and I have preserved all three anchor sites, so
a repairer cannot fix one and leave the others). P5 collapses two bookings across seat G's two readings onto
ONE predicate at ONE anchor. **No cross-SEAT collapse was made**: P1–P4 (seat C, `ai_architecture.md`) and
P5–P7 (seat G, `hiring.md` / `clerk-integration.md`) share no falsehood. In particular P1, P2, P4 and P5 are
all "an anchor names the wrong construct" in *shape*, but they are four distinct false propositions about
four distinct constructs in three distinct files, and the brief forbids collapsing anchors that merely look
similar.

Both rejections are recorded as `ref-discipline`, the class the brief predicts and instructs me to filter:
C31-B3 grades a current, exact citation against a stale local checkout; G31-B1 grades a `v1.40.2`-pinned
module description against `app`'s diverged post-fold in-tree copy. Neither is a `wrong-tree` rejection —
both seats named their trees correctly and reasoned from the right files.

---

BOOKED=11 UPHELD=9 REJECTED=2 IN-SCOPE-UPHELD-BLOCKERS=9 DISTINCT-IN-SCOPE-PREDICATES=7
