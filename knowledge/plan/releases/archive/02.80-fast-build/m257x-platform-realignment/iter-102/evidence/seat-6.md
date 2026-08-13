# seat-6 report

**Files owned:** `corpus/services/ai-readiness.md` · `ai-labs.md` · `askengine.md` · `coursebuilder.md`.
**Anchors booked: 7. Sites found: 8. Sites repaired: 8.** No file outside the four was edited; nothing committed;
no clone was fetched or written.

**Ground truth used** — `app` `ad9f3c49` (`== origin/main`, re-verified `git rev-parse`), `next-web-app`
`8297c684` (the ref the brief names; the clone's `origin/main` is `f97ba659` and was **not** read),
platform `0c91421`. Every line number below was re-derived at those refs with `git grep`/`git show` at a named
ref — none inherited from a previous sheet.

## Ledger rows

| # | the false claim | anchor | what is true | sites repaired |
|---|---|---|---|---|
| 1 | "**`/ai-readiness` is the only readiness route the navbar links** — `AI_READINESS_URL` (`packages/core-js/src/constants/urls.ts:52`), consumed by `packages/ui/src/NavBar/useNavbarSections.tsx` — imported at `:4`, built into `aiReadinessMenuItem` at `:398-400`, gated at `:547` (`showAIReadiness ? aiReadinessMenuItem : null`). A repo-wide grep finds the constant in exactly those two non-`node_modules` files." | `corpus/services/ai-readiness.md:305` | At `next-web-app` `8297c684`: `AI_READINESS_URL` is at **`urls.ts:51`**; `:52` is `TALK_TO_DATA_URL` (it was `ORGANIZATION_FEEDBACK_URL` at `bb3313bc`). **Independently re-derived, not inherited:** over *every* reachable commit touching that file the constant's line is **41, 50 or 51 only** — never 52. `aiReadinessMenuItem` is built at **`:401-408`** (`key: AI_READINESS_URL` at `:403`), gated at **`:569`**; the `:4` import still holds. And the grep is **3** non-`node_modules` files, not two — the third is the platform's own KB page `knowledge/ai-readiness/frontend-architecture.md:15`. | 1 |
| 2 | "This paragraph stated completed work as outstanding and contradicted `:459` of this same file, which is already in the past tense; corrected M257x iter-46." | `corpus/services/ai-readiness.md:46` | `:459` opened the `> **✅ CORRECTED M219 …**` blockquote — an unrelated claim, carrying nothing in the past tense about the D-07 re-anchor. The statement the sentence needs is the **⚠⚠ M51 iter-08/09** block's parenthetical *"(now `aireadiness/readiness.go`, formerly `workforce/ai_readiness.go:512`)"*, **now `:496`** (verified post-edit). The cite has never resolved: iter-46 wrote `:458` (then `**unbounded whole-org member hydration**`, verified at `301d61a`), and iter-100 mechanically +1'd it to `:459` (`a229f8d`). Repaired by **naming the construct** and pinning the number as a convenience. | 1 |
| 3 | "The field exists in the API and in the FE's TypeScript type, `apps/web/src/hooks/useAIReadiness.ts:326`, and is drawn by nothing" | `corpus/services/ai-readiness.md:595` | **The number is now right and the citation was still defective — it was unpinned.** At `bb3313bc` (iter-101's graded ref) `interviewQuestions: number;` was at `:274` and `:326` was `headers: {`; at `8297c684` it *is* at `:326`, inside `export interface AIReadinessCycleTotals` (`:323-330`). iter-100 changed `:274`→`:326` and was wrong at the ref of the day, right at the ref of this day. Repaired by pinning the ref and naming the interface — plus verifying the five sibling anchors in the same parenthetical all still hold at `8297c684` (`HowWeMeasureTab.tsx:1879`/`:1903`/`:1915`/`:1921`/`:1927`, 1,989 lines, `grep -c interviewQuestions` → 0). | 1 |
| 4 | "a failed Bedrock init **disables** Talk-to-Data but doesn't crash `app`" | `corpus/services/askengine.md:88-89` | At `app` `ad9f3c49`, `main.go:467-471` is `bedrockClient, err := askengine.NewBedrockClient(serverContext)` / `if err != nil { logger.Error("bedrock client unavailable; talk-to-data disabled", …); return }`. The `return` is one tab-level inside `func main()` (`:229`; **no `func` declaration falls between `:229` and `:467`** — enumerated), so it returns from `main`: the RPC mux (`:1295`), the meta HTTP server (`:1361`), the Echo router, the Asynq pools and the Redis subscribers (`:1438`+) are never constructed. Nothing is "disabled" — the corpus had taken the platform's own log string as a description of behaviour. **Narrowness disclosed rather than smoothed:** `NewBedrockClient` (`internal/askengine/bedrock.go:161`) has exactly one error return — `config.LoadDefaultConfig` failing — so *absent* creds do not normally trigger it. | 1 |
| 5 | "**SSE wire contract** (`event:` name / `data:` JSON): `session`, `stage`, `outline`, `progress`, `patch_applied`, `patch_skipped`, `preview_ready`, `draft_kept`, `translation_ready`, `rebuild_required`, `error`, `cost`, and an always-last `done`." | `corpus/services/coursebuilder.md:77-79` | **16 wire names at `ad9f3c49`, and `cost` is not one of them.** `session` is written directly at `handler.go:1458`; `renderEvent` (`:2604-2772`) returns the other 15 — `text` `:2607`, `score` `:2609`, `patch_applied` `:2614`, `patch_skipped` `:2620`, `stage` `:2637`/`:2639`, `outline` `:2654`, `progress` `:2661`, `preview_ready` `:2671`, `draft_kept` `:2678`, `error` `:2704`, `translation_ready` `:2724`, `rebuild_required` `:2732`, `steering_received` `:2744`, `steering_applied` `:2754`, `done` `:2768` (the `default` arm re-emits `text` at `:2770`). `case cb.EventCost:` (`:2709`) returns an **empty** name (`:2717`) under the in-code D2 ruling, and the loop skips `writeSSE` on an empty name (`:1478-1480`). The doc listed 13, included `cost`, and omitted `text`, `score`, `steering_received`, `steering_applied`. | 1 |
| 6 | "`app` is **`v1.369.0`** @ origin/main `2035f9a4` as of 2026-08-06, six releases on" | `corpus/services/coursebuilder.md:132` (**CANON-3**) | `app` `origin/main` is **`ad9f3c49`** (2026-08-06); `2035f9a4..ad9f3c49` = **5 commits**. The newest **tag** is still `v1.369.0` — `git describe --tags ad9f3c49` → `v1.369.0-7-gad9f3c498`, i.e. seven commits past the tag. CANON-3 **move (2)** applied: re-derived at `ad9f3c49`, re-stated with the new ref and a date, and the expired **label** separated from the still-valid **pin** (`2035f9a4` = `v1.369.0-2-g2035f9a40`). No standing "current" reintroduced (the iter-98 P6 class this exact file caused). | 1 |
| 7 | "those run in the `v1.3xx` range — **`v1.369.0`** @ origin/main `2035f9a4`, measured 2026-08-06" | `corpus/services/ai-labs.md:18` (**CANON-3**) | Same measurement, same CANON-3 move (2), same wording family as row 6. | 1 |
| 8 | "In current `origin/main` there is **no `checkout.session.completed` webhook, no labs↔credits linkage, and `/credits/purchase` was removed (Wave 13)**." | `corpus/services/ai-labs.md:16` (**paraphrase of CANON-3 — not booked, found by predicate**) | A bare moving label, which CANON-3 forbids outright. All three facts **re-verified true** at `ad9f3c49` before pinning — `checkout.session.completed` = 0 occurrences in Go source; `internal/labs/` imports `internal/credits` nowhere; `/credits/purchase` removed in Wave 13 (`internal/web/backend/credits/handler.go:12`). Truth unchanged, label converted to a pin. | 1 |

## Reach — graded, not asserted

| predicate | anchors BOOKED | sites FOUND | sites REPAIRED | how I searched |
|---|---|---|---|---|
| `ai-readiness-url-line` / `ai-readiness-urls-ts-52` | 1 | 1 | 1 | `git grep -n "AI_READINESS_URL\|urls\.ts"` over `corpus/**` + `CLAUDE.md` + `README.md`. 19 hits, all others about the demopatch `urls.ts` **chain**, a different subject. No twin. |
| `ai-readiness-self-cite-458` | 1 | 1 | 1 | `grep -n "this same file\|this file\|above at \`:\|below at \`:"` over all four files; plus `git log -p` on the anchor to recover what iter-46 and iter-100 each wrote. |
| `interview-questions-fe-line` | 1 | 1 | 1 | `git grep -n "useAIReadiness\.ts\|interviewQuestions"` over `corpus/**` + `CLAUDE.md`. Two other hits (`frontend_architecture.md:39`, `ai-readiness.md:582`) name the file/field but assert no line for the FE type. |
| `askengine-bedrock-disabled` | 1 | 1 | 1 | `git grep -ni "talk-to-data disabled\|bedrock init\|bedrock client unavailable\|failed Bedrock"` corpus-wide → exactly one site. Cross-checked with a degradation-vocabulary sweep (`crash\|disable\|degrad\|nil client\|no-op\|unmount`) over all four files. |
| `coursebuilder-sse-cost-event` | 1 | 1 | 1 | `git grep -n "preview_ready\|draft_kept\|translation_ready\|SSE wire"` corpus-wide → only the two lines of this bullet. |
| **CANON-3 currency pin** | **2** | **3** | **3** | `grep -n "2035f9a\|origin/main\|b948604\|5ba17044\|ad9f3c4\|v1\.3[0-9][0-9]"` over all four files. The third site (`ai-labs.md:16`) is a **paraphrase** — a bare `origin/main` with no sha — exactly the class rule 2 says escapes anchor-scoped repair. |
| **TOTAL** | **7** | **8** | **8** | |

**Booked-vs-live width: 7 → 8 (1.14×).** Lower than the milestone's ~3× average, and I believe that is real
rather than under-searching: six of my seven anchors are *citations to a specific `file:line`*, which by
construction occur once. The one generalising predicate (CANON-3) is the one that widened.

**Collateral corrections inside repaired sentences** (not separate predicates — text I rewrote and therefore had
to leave true, per rule 4): `aiReadinessMenuItem` `:398-400`→`:401-408`, the nav gate `:547`→`:569`, the
"exactly those two non-`node_modules` files" count (2→2 **source** files + 1 KB page), and the bare **"at HEAD"**
moving label in the same block → pinned at `8297c684`.

**Verified-and-left-alone** (rule 5 — true halves, and true anchors, not weakened): the `:4` import in
`useNavbarSections.tsx`; `WorkforceNewClient.tsx:125-151`; `useWorkforceAIReadiness.ts:23-27` and its 0 `cycle`
mentions; `HowWeMeasureTab.tsx` `:1879`/`:1903`/`:1915`/`:1921`/`:1927` + the 1,989-line count + the 0
`interviewQuestions` hits; `e2e/specs/web.ai-readiness.spec.ts`; `bedrock.go:161`; the `b948604` v1.366.0 pins at
`coursebuilder.md:48`/`:99`. All re-measured at the named ref, all hold.

**`ai-labs.md:75` (iter-101's wrong-convention REJECTION) was not touched.** My only edit to that file is
lines 16–20; the rejected site is now `:77`, byte-identical to `HEAD:corpus/services/ai-labs.md:75`.

## Twins outside my files (REPORT, do not edit)

| predicate | file:line | why it is the same claim |
|---|---|---|
| — | — | **None.** All six of my non-CANON-3 predicates are single-site; the CANON-3 sites outside my files are already partitioned to seats 1, 2, 5, 7, 9 and the orchestrator by `canonical-repairs.md`. |

**Inbound-citation drift my own edit caused — flagged because I am not allowed to fix it.**
`corpus/ops/demo/stories-spec.md:603` cites `services/ai-readiness.md:371,449-450`. My edits added 13 lines above
both: the cited constructs are now **`:384`** (`5. ai_readiness_cycles **× 2 …`) and **`:462-463`**
(`**The frozen path is reachable BOTH …**`). Two notes for whoever repairs it: (a) that citation was *already*
imprecise before I touched anything — neither construct is the 2.09 s retraction it is cited for, which sits at
`ai-readiness.md:458-460`; (b) `ai-readiness.md` total went **651 → 668**, so any other inbound line cite below
`:46` is shifted by +5, below `:302` by +12, below `:311` by +13, and below `:604` by +17.

## Noticed, not repaired

1. **`coursebuilder.md:71-72` — a count that is wrong at its own pinned ref.** It says *"the **two**
   `logger.Warn(… routes disabled)` arms at `:774` / `:778`"* @ `app` `b948604` v1.366.0. Measured at that exact
   ref: there are **three** — `:774` (author client), `:778` (grader client) and **`:806`** (*"coursebuilder
   service init failed; /coursebuilder routes disabled"*). The two named arms resolve and are correctly quoted;
   only the cardinality fails. Same shape at `ad9f3c49` (`:884`, `:888`, `:916`). A different predicate from any
   of my seven, so booked here rather than repaired.
2. **`ai-readiness.md:84`** — *"It gates on `orgEnabled` alone (`:133-134`, `const { orgEnabled } =
   useAiReadinessEnabled(true)`)"*. At `8297c684` that statement is `AIReadinessClient.tsx:135`. Same *class* as
   my rows 1 and 3 (an unpinned FE line that drifted) but a different construct and a different predicate.
3. **`ai-readiness.md` cites two different files by the same basename.** `:522` uses `useAIReadiness.ts:48-62`
   meaning `apps/web/src/components/ai-readiness/useAIReadiness.ts` (`deriveMode` at `:48` — **verified, it
   resolves**), while row 3's anchor means `apps/web/src/hooks/useAIReadiness.ts`. Both files exist at
   `8297c684`. Not false, but the short form is ambiguous and cost me a measurement to disentangle.
4. **`ai-readiness.md:460`** cites `ops/demo/stories-spec.md:599` for the 2.09 s retraction; `:599` is the
   *retracted claim*, the retraction blockquote opens at `:601`. Off by two, arguably the intended target.

## What I could not settle, and why

Nothing in my assignment was left unsettled — all 7 booked anchors and the 1 predicate-expanded site are
repaired against a named ref.

Two disclosures rather than gaps:

- **The `next-web-app` clone's `origin/main` is `f97ba659`, ahead of the `8297c684` the brief pins.** Per the
  brief I did not fetch and did not read `f97ba659`, so every frontend number here is a reading at `8297c684`
  and is written into the corpus **as such**. Row 3 is the cautionary case: the same citation was false at
  `bb3313bc` and true at `8297c684` without anyone editing it, which is why I pinned the ref rather than just
  correcting the digit.
- **Row 4 is the one place I deliberately did not repair as hard as the booking allowed.** The adjudicator's
  finding (`return` exits `main`) is exact and is now stated. But the natural next sentence — *"so Bedrock is a
  hard boot prerequisite"* — does **not** follow: `NewBedrockClient`'s only error return is
  `config.LoadDefaultConfig`, which does not fail on merely-absent AWS credentials. I measured that before
  writing, and wrote the narrower true claim instead of the wider satisfying one. It also means my repair does
  **not** contradict `coursebuilder.md`'s graceful-404 story (Course Builder builds its own clients, `main.go`
  `:882-916`), which I checked explicitly for the self-contradiction rule-4 requires.
