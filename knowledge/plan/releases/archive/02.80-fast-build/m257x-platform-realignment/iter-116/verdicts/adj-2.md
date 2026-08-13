# adj-2 — verdicts for seats C (r29-C, r30-C) and G (r29-G, r30-G)

## Trees I read (stated per rule 45/50 and the §-warning about the frozen instrument's line 37)

| tree | path | ref I read |
|---|---|---|
| app | `stack-demo/app` | `ad9f3c49` (+ `1e457fa70`, `2b3a65cf0` for the fold history) |
| app/studio (nested, own checkout) | `stack-demo/app/studio` | `aeec036a` |
| next-web-app | `stack-demo/next-web-app` | `8297c684` |
| studio-desk | `stack-demo/studio-desk` | `41ee3575` |
| sentinel | `stack-demo/sentinel` | `f2c46190` (+ `88bc5592`, `88036d7`) |
| rosetta-extensions — **pinned per-stack consumption clone** | `stack-demo/rosetta-extensions` | `09d06070` — **the tree I settled every "what the tooling does on a stack" claim in** |
| rosetta-extensions — authoring copy | `.agentspace/rosetta-extensions` | `1dc1eb82` — read only to confirm the two trees agree on `persona_write.go` (byte-identical, `diff` empty) and on `build_frontend_hiring()` (7 in both) |

Corpus read at the working tree. No `git fetch`, no `git pull`, no git state changes; the only file I
wrote anywhere in the repo is this one.

Every verdict below was re-derived by opening the platform/corpus file myself. Where a seat's own line
numbers differ from mine (e.g. r30-G's `up-injected.sh:1096…1122` vs my `:1220…1289` in the pinned
clone) I report **my** measurement; the verdict does not depend on the seat's arithmetic.

---

## Verdicts

### r29-C

```
r29-C B1 | corpus/architecture/ai_architecture.md:92 | UPHELD | IN-SCOPE | PREDICATE: The OpenAI Model-Families table row of ai_architecture.md is at line 86.
   evidence: corpus/architecture/ai_architecture.md, read with awk line numbers: :86-87 is the wrapped
   prose sentence "The **Default client** column is the client the vendor const resolves to — *not* a
   position in a / fallback order (there is none; see above)."; :89-90 are the table header+separator;
   :91 is "| **OpenAI (Azure EU + Direct US)** | GPT-5.4, GPT-5.4-mini, …". The citation at :92 reads
   "its OpenAI sibling at `:86` is an exact enumeration of `internal/ai/openai/config.go:8-26`" — the
   sibling ROW is five lines below the line cited. The substantive half is true (11 consts / 11 models),
   which is why only the locator is booked. Not ref-discipline: an intra-corpus self-citation carries no
   pin and is settled at the corpus's current state.
```

```
r29-C B2 | corpus/architecture/ai_architecture.md:94 | UPHELD | IN-SCOPE | PREDICATE: The ai-module-fold passage of ai_architecture.md is at line 95.
   evidence: :94 says "a module **no repo a stack builds requires** since the fold at `1e457fa70`
   (see `:95` below)". Measured, :95 is the table row "| **Transcription** | GPT-4o Transcribe |
   Azure EU (US via `flag_use_azure_us`) |" — a speech-to-text row that says nothing about a module
   fold. The passage that actually substantiates it is :100-108 ("it is no longer a shared private
   module for any service a stack builds … `app` folded the library into its own tree at `1e457fa70`").
   I confirmed the underlying claim independently at `app` `ad9f3c49`: `internal/ai/` is in-tree and
   `git show ad9f3c49:go.mod | grep -c anthropos-work/ai` → 0. Same −5 offset as B1/B3.
```

```
r29-C B3 | corpus/architecture/ai_architecture.md:295 | UPHELD | IN-SCOPE | PREDICATE: The Anthropic Model-Families table row of ai_architecture.md is at line 87.
   evidence: :295 reads "**NB this bullet is not a claim about the platform's Anthropic model SET**
   (that is the table row at `:87`, which is where the 4.6/Opus-4.8 omission lived)". Measured, :87 is
   "fallback order (there is none; see above)." — the wrapped tail of the prose sentence introducing the
   table. The Anthropic row is :92 (it is the row that carries "**Claude 4.6 Sonnet** … five families
   over six consts … plus **Claude Opus 4.8**"). Present-tense locator ("that IS the table row at :87"),
   so rule 7's historical-anchor carve-out does not reach it.
```

```
r29-C B4 | corpus/architecture/ai_architecture.md:131 | UPHELD | IN-SCOPE | PREDICATE: `studio/configs/production_config.ini` is readable, and was measured, at `app` HEAD.
   evidence: `git -C stack-demo/app show ad9f3c49:studio/configs/production_config.ini` →
   "fatal: path 'studio/configs/production_config.ini' exists on disk, but not in 'ad9f3c49'";
   `git -C stack-demo/app ls-tree --name-only ad9f3c49 studio` → empty. The file is in the nested
   `anthropos-studio-room` checkout `stack-demo/app/studio` @ `aeec036a`, where its real in-repo path
   has no `studio/` prefix at all (`git ls-tree aeec036a` shows repo-root entries). It also contradicts
   an in-scope corpus statement asserted as live: `corpus/architecture/external_services.md:569-575` —
   "**`app/studio/**` is an IN-IMAGE path, and it is in no `app` commit.** … `git show <ref>:studio/…`
   against an `app` clone returns nothing at **every** ref … Resolve them against the studio-room repo."
   The VALUES are correct — I read `configs/production_config.ini:26-36` @ `aeec036a` and all ten
   `*_AI_*_MODEL` lines match the table at :137-141 exactly (FAST/STRICT `azure, gpt-5-mini, none`;
   EXECUTION `azure, gpt-5.4, none`; CREATIVE `…, low`; REASONING `…, medium`; stable == experimental).
   This is a provenance/ref defect only, and I say so; but "measured at `app` HEAD" against a path
   `app` HEAD does not contain is unsupportable, and the same file gets it right at :40-42 by grepping
   `app/studio` at `aeec036a`.
```

```
r29-C B5 | corpus/architecture/ai_architecture.md:326 | UPHELD | IN-SCOPE | PREDICATE: The `ai` library is still a shared module external to `app`.
   evidence: :326 — "fed by `Event_AiUsage` messages that the AI-consuming services publish over Redis
   Streams (**the shared `ai` library** itself only returns provider token counts)". The same file at
   :100-108 asserts the contrary AS LIVE: "it is **no longer a shared private module for any service a
   stack builds** … `app` **folded the library into its own tree** at `1e457fa70` … the library lives at
   `app/internal/ai/`" — and at :105-106 instructs "**When you repair this predicate, repair both sites
   or neither** (`TOK-07` rule 3)". Re-derived at `app` `ad9f3c49`: `internal/ai/` is an in-tree package
   (`git ls-tree ad9f3c49:internal/ai` → ai.go, anthropic/, openai/, …) and `go.mod` requires
   `anthropos-work/ai` 0 times. Rule 5's retraction carve-out does not apply: the retraction and the
   retracted descriptor are BOTH asserted live, 218 lines apart, in one file.
   NOTE — this is the weakest of my upholds: the mechanical clause around the adjective is true, and
   r30-C filed the same site as a MINOR. I uphold it because rule 5 makes an intra-file live
   self-contradiction a finding on its own and the file names this exact pair as a repair unit.
```

### r30-C

```
r30-C B1 | corpus/architecture/ai_architecture.md:92 | UPHELD | IN-SCOPE | PREDICATE: The OpenAI Model-Families table row of ai_architecture.md is at line 86.
   evidence: identical anchor and identical falsehood to r29-C B1 — see that entry. Collapses onto P1.
```

```
r30-C B2 | corpus/architecture/ai_architecture.md:295 | UPHELD | IN-SCOPE | PREDICATE: The Anthropic Model-Families table row of ai_architecture.md is at line 87.
   evidence: identical anchor and identical falsehood to r29-C B3 — see that entry. Collapses onto P3.
```

```
r30-C B3 | corpus/architecture/ai_architecture.md:110-111 | UPHELD | IN-SCOPE | PREDICATE: `app`'s folded `ai.AI` interface still spans Mistral, unchanged by the fold.
   evidence: the paragraph re-scopes itself to the folded in-tree library and pins itself
   ("at `app` `ad9f3c49` neither `app/go.mod` nor `sentinel/go.mod` requires … the library lives at
   `app/internal/ai/`", :107-108), then asserts currency at :110 ("What the interface provides is
   **unchanged**:") and lists Mistral at :111. Measured at that very ref:
   `git ls-tree --name-only ad9f3c49:internal/ai` → ai.go, **anthropic**, core_characterization_test.go,
   json.go, messages.go, metadata.go, module_import_guard_test.go, **openai**, options.go, prompt.go,
   response.go, retry.go, retry_characterization_test.go, speech.go, tokenizer.go, transcriptions.go,
   types.go — **there is no `mistral/`**. `internal/ai/ai.go:8-17` declares `type AI interface` with
   eight methods. The surviving Mistral code is `internal/cms/studio/mistralocr/mistralocr.go`, whose
   own package doc (:1-11) says "It used to be internal/ai/mistral, where it satisfied the nine-method
   ai.AI interface … No interface, no LLM methods, no tokenizer, and no panics." Datable:
   `git ls-tree 1e457fa70:internal/ai` DOES list `mistral`; `2b3a65cf0` ("refactor(cms): re-home Mistral
   OCR + fix the API-key bug pair", 2026-08-04) removed it and `git merge-base --is-ancestor 2b3a65cf0
   ad9f3c49` succeeds. So the fold-plus-re-home DID change what the interface provides — which is what
   ":110 unchanged" denies. The only `ai.AI` implementors of `OCRProcess` at `ad9f3c49` are
   `openai/ocr.go:24` and `anthropic/completion.go:195` (the latter returns an error).
```

### r29-G

```
r29-G B1 | corpus/services/clerk-integration.md:40 | UPHELD | IN-SCOPE | PREDICATE: Clerk sign-in tokens are minted only for app-native admin impersonation.
   evidence: I grepped each clone at its own ref. `createSignInToken` / `signInTokens` resolves to FOUR
   minting sites, three of which are not admin impersonation:
     · app `ad9f3c49` `internal/admin/impersonation/manager.go:101` — `m.signInTokenCl.Create(...)`,
       the admin-impersonation feature the bullet names (import at :29). This half is correct.
     · next-web-app `8297c684` `apps/web/src/app/api/dev/login-as/route.ts:79` —
       `client.signInTokens.createSignInToken({ userId: user.id, expiresInSeconds: 600 })`, reached
       after `client.users.getUserList({ emailAddress: [email] })` at :65-69. I read the whole handler:
       it signs you in as an ARBITRARY user found by email, with no admin scoping.
     · next-web-app `8297c684` `e2e/auth.setup.ts:72` — the Playwright auth bootstrap.
     · studio-desk `41ee3575` `src/routes/dev.ts:83` — the same call, ticket handed to
       `/dev-accept.html` at :89-91.
   Both dev routes are NODE_ENV-gated, but they are checked-in code in two repos and the corpus itself
   documents the surface (`studio-desk.md:34`, `:90` name `dev-accept.html`), so "only" is false as
   written. Caveat on the seat's corroboration, which I checked and do NOT rely on: the current
   `next-web-app.md:75` allowlist reads "/login, /sign-up, /checkout, /free-trial, /monitoring, /print,
   /api/bunny/thumbnail" — it does NOT list `/api/dev/login-as`. The uphold rests on the source grep.
```

```
r29-G B2 | corpus/services/studio-desk.md:319 | UPHELD | IN-SCOPE | PREDICATE: `MOCK_CLERK=true` is studio-desk's path for interactive local dev without real Clerk.
   evidence: the corpus offers it under the heading "**Local dev without real Clerk**". At studio-desk
   `41ee3575` the repo denies exactly that purpose in three places:
     · `src/index.ts:50-54` — "MOCK_CLERK bypasses Clerk for the headless automated test suite …
       **It is NOT the path for interactive / agentic dev** — for that, sign in as a REAL user via
       /api/dev/login-as (no Google, no 2FA)."; the boot warning at :56 says "test-harness mode".
     · `.env.example:23-27` — "MOCK_CLERK / VITE_MOCK_CLERK bypass Clerk for the AUTOMATED TEST HARNESS
       only … **They are NOT a path for interactive dev**: the synthetic user's fake org never matches
       real platform data, so **GraphQL-backed content won't load.** To work as a real signed-in user,
       use the DEV LOGIN above." (the DEV LOGIN block is `.env.example:15-21`).
     · `app/core/main.ts:133` — "VITE_MOCK_CLERK is enabled — synthetic test-harness auth (not for
       interactive dev; use /api/dev/login-as)."
   And the sanctioned alternative is mounted at `src/index.ts:149-155` yet appears NOWHERE in
   `studio-desk.md` (`grep -n 'login-as' corpus/services/studio-desk.md` → no hits; only `dev-accept.html`
   at :34 and :90). A documented procedure whose source says it does not do the documented job is
   unsupportable, and the failure mode is concrete (content silently empty → the reader debugs the
   platform). NOTE: the literal sub-clause "to bypass Clerk auth" is true; a stricter adjudicator could
   reject this as `already-true`. I resolve upward because the heading states the purpose.
```

```
r29-G B3 | corpus/services/hiring.md:243 | UPHELD | IN-SCOPE | PREDICATE: The PersonaSeeder's three validation-table writes are at persona_write.go:69-71,143-167.
   evidence: settled in the PINNED per-stack clone `stack-demo/rosetta-extensions` `09d06070` (this is a
   claim about what the tooling does on a stack); `diff` against the authoring copy `1dc1eb82` is empty,
   so the tree choice cannot change the answer. Measured in `stack-seeding/seeders/persona_write.go`:
     · :69 is a bare `//`; :70-71 are prose about the write that was **removed** — "The old
       `jobsimulation.sessions` step is REMOVED rather than re-pointed: it wrote the SAME / a.sessions
       rows under the SAME id as the `public.job_simulation_sessions` step below, so …".
     · the three-table mapping comment is :66-68.
     · the ACTUAL writes are the `steps` table at :91-95 — `{"public","validation_attempt_results",…}`
       :92, `{"public","validation_attempt_skill_results",…}` :93,
       `{"public","validation_criterion_results",…}` :94.
     · :143-167 is `}` (end of `flush`) :143, `// --- column lists …` :145, `sessionCols()` :152-159
       (the SESSION columns, not a validation table) and `attemptResultCols()` :161-167.
       `skillResultCols()` begins at :169 — outside the cited range.
   So neither cited range contains the construct the sentence attributes to it. The substantive claim
   (the PersonaSeeder does write those three tables) is TRUE; only the locator is false.
```

```
r29-G B4 | corpus/architecture/shared_libraries.md:57 | UPHELD | IN-SCOPE | PREDICATE: sentinel's `88036d7` is two commits past `88bc5592`.
   evidence: `git -C stack-demo/sentinel rev-list --count 88bc5592..88036d7` → **1**;
   `git log --oneline 88bc5592..88036d7` → a single line, "88036d7 chore(deps): update dependencies to
   latest versions", i.e. `88036d7` is the immediate child of `88bc5592`. (`88bc5592..f2c46190` → 2 —
   the CHECKOUT is two past, which is the likely origin of the number.) The load-bearing half is exact
   and I verified it: `git show 88bc5592:go.mod` → `colony v0.34.3` at :8; `git show 88036d7:go.mod` →
   `colony v0.35.2` at :8. Ambiguity considered and rejected: the appositive "two commits past …"
   grammatically modifies `88036d7`, the subject of "took it v0.34.3 → v0.35.2", not the row's own ref.
```

### r30-G

```
r30-G B1 | corpus/services/hiring.md:434-436 | UPHELD | IN-SCOPE | PREDICATE: Four demo-patches bake into the demo hiring image.
   evidence: settled in the PINNED clone `09d06070` (what a stack runs), and cross-checked in the
   authoring copy. `awk '/^build_frontend_hiring\(\)/,/^}/' demo-stack/up-injected.sh |
   grep -c '"\$demopatch" apply'` → **7** at `09d06070` AND **7** at `1dc1eb82`. The function opens at
   up-injected.sh:1082 and the seven apply calls are at :1220 (`next-web-studio-url`), :1223
   (`next-web-public-website-url`), :1242 (`next-hiring-role-remap`), :1259
   (`next-hiring-members-pagination`), :1272 (`next-web-interview-flag-container`), :1280
   (`next-web-interview-flag-result`), :1289 (`next-web-back-to-cockpit`); it declares exactly those
   seven `local *_manifest=` paths, and the RETURN trap reverts all seven. So the corpus's "Four
   demo-patches … 2 net-new … + the 2 chained shared `urls.ts` patches" understates by three, and the
   cross-reference it hands the reader ("§ the four hiring-image patches") is built on the same figure.
   (The sibling `corpus/ops/demo/demopatch-spec.md:293` "bakes FOUR patches" and :300 "4-manifest
   patch-set fingerprint union" carry the same stale count, but that file is OUT-OF-SCOPE; the booked
   anchor is in `corpus/services/`.)
```

```
r30-G B2 | corpus/services/hiring.md:47-49 | UPHELD | IN-SCOPE | PREDICATE: service_taxonomy.md's Tier-1 Database characteristic bullet is at line 62.
   evidence: `corpus/architecture/service_taxonomy.md:62` is `### Tier 1: Core Backend Services
   (Dockerized Go Microservices)` — a section heading. The quoted bullet is at **:68**:
   "- **Database**: PostgreSQL — **one schema, `public`, owned by `app`**, which is the only repo with
   migrations (`repos.yml:14-17`) … the `cms`, `jobsimulation` and `skillpath` schemas are legacy husks"
   (`grep -n 'one schema' corpus/architecture/service_taxonomy.md` → 68, single hit). The sibling
   citation in the same sentence resolves exactly: `dependency_map.md:78` is the jobsimulation
   session-state bullet ending "(the legacy `jobsimulation` schema is non-authoritative)". The substance
   of hiring.md's claim (the twin says X) is TRUE at :68; only the number is false. The passage hedges
   ("The bullet is named as well as numbered because `service_taxonomy.md` is edited concurrently and
   `:62` can move") and records an earlier repair from `:52` — but naming the construct does not
   un-assert the number, and rule 7 protects a passage that names a construct INSTEAD of a line, not one
   that asserts a wrong line alongside it.
```

```
r30-G B3 | corpus/services/hiring.md:243 | UPHELD | IN-SCOPE | PREDICATE: The PersonaSeeder's three validation-table writes are at persona_write.go:69-71,143-167.
   evidence: identical anchor and identical falsehood to r29-G B3 — see that entry for the full
   re-derivation in the pinned clone. Collapses onto P9.
```

---

## Rejections

**None.** I screened every booking for the ref-discipline class explicitly and none is a member: no
booked claim is pinned, dated or past-tense in a way that newer evidence contradicts. The five
`ai_architecture.md` anchor bookings are intra-corpus self-citations, which carry no pin and settle at
the corpus's current state; the two rext bookings settle at the PINNED per-stack clone (and agree with
the authoring copy, so neither is `wrong-tree`); `shared_libraries.md:57`'s distance claim is
ref-independent; `clerk-integration.md:40` and `studio-desk.md:319` are undated procedural claims
refuted by the clones at the refs the ground-truth table names.

Two upholds are close and I name them so they can be re-graded rather than inherited: **r29-C B5**
(one retracted adjective republished — the mechanical clause around it is true, and r30-C filed the
same site as a MINOR) and **r29-G B2** (the literal sub-clause "bypasses Clerk auth" is true; what is
false is the stated purpose).

---

## PREDICATE ROLL-UP

```
P1  | The OpenAI Model-Families table row of ai_architecture.md is at line 86.                        | anchors: r29-C B1 @ corpus/architecture/ai_architecture.md:92, r30-C B1 @ corpus/architecture/ai_architecture.md:92
P2  | The ai-module-fold passage of ai_architecture.md is at line 95.                                 | anchors: r29-C B2 @ corpus/architecture/ai_architecture.md:94
P3  | The Anthropic Model-Families table row of ai_architecture.md is at line 87.                     | anchors: r29-C B3 @ corpus/architecture/ai_architecture.md:295, r30-C B2 @ corpus/architecture/ai_architecture.md:295
P4  | `studio/configs/production_config.ini` is readable, and was measured, at `app` HEAD.            | anchors: r29-C B4 @ corpus/architecture/ai_architecture.md:131
P5  | The `ai` library is still a shared module external to `app`.                                    | anchors: r29-C B5 @ corpus/architecture/ai_architecture.md:326
P6  | `app`'s folded `ai.AI` interface still spans Mistral, unchanged by the fold.                    | anchors: r30-C B3 @ corpus/architecture/ai_architecture.md:111
P7  | Clerk sign-in tokens are minted only for app-native admin impersonation.                        | anchors: r29-G B1 @ corpus/services/clerk-integration.md:40
P8  | `MOCK_CLERK=true` is studio-desk's path for interactive local dev without real Clerk.           | anchors: r29-G B2 @ corpus/services/studio-desk.md:319
P9  | The PersonaSeeder's three validation-table writes are at persona_write.go:69-71,143-167.        | anchors: r29-G B3 @ corpus/services/hiring.md:243, r30-G B3 @ corpus/services/hiring.md:243
P10 | sentinel's `88036d7` is two commits past `88bc5592`.                                            | anchors: r29-G B4 @ corpus/architecture/shared_libraries.md:57
P11 | Four demo-patches bake into the demo hiring image.                                              | anchors: r30-G B1 @ corpus/services/hiring.md:434
P12 | service_taxonomy.md's Tier-1 Database characteristic bullet is at line 62.                      | anchors: r30-G B2 @ corpus/services/hiring.md:48
```

### Collapses I made, and the ones I deliberately did NOT make

- **Made:** r29-C B1 ≡ r30-C B1 (P1) and r29-C B3 ≡ r30-C B2 (P3) — the two seat readings booked the
  same anchor for the same falsehood. r29-G B3 ≡ r30-G B3 (P9) — likewise, `hiring.md:243`.
- **Did NOT make:** P1, P2 and P3 all arise from one editing event (each self-citation is exactly −5
  lines, the signature of an insertion above `:86`), but they are three DIFFERENT false propositions
  naming three different target constructs — repairing one leaves the other two false, and each sends a
  reader to a different wrong place. Under the brief's phrasing test they do not yield the same
  sentence. I record the shared cause here so a repairer fixes all three in one sweep.
- **Did NOT make:** P5 and P6 both live in the `ai`-library paragraph family, but "the library is still
  shared" and "the interface still spans Mistral" are independent propositions with independent
  refutations (`go.mod` vs `ls-tree internal/ai`).
- **Did NOT make:** P9 and P12 are both wrong-locator predicates in `hiring.md`, and P11 is a wrong
  count in the same file — three distinct falsehoods, not one "hiring.md is stale" predicate.

---

BOOKED=15 UPHELD=15 REJECTED=0 IN-SCOPE-UPHELD-BLOCKERS=15 DISTINCT-IN-SCOPE-PREDICATES=12
