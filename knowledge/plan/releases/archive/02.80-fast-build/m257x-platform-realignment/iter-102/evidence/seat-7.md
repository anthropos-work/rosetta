# seat-7 report

**Files owned:** `corpus/architecture/dependency_map.md` · `corpus/architecture/shared_libraries.md` ·
`corpus/architecture/frontend_architecture.md`. Nothing outside them was edited. No commit, no
`git fetch`, no write to any clone.

**Refs settled against:** platform `0c91421` · `app` `ad9f3c49` (== `origin/main`, 2026-08-06) with
`2035f9a4` and `b948604f` read as historical · `next-web-app` `8297c684` (demo pin, 2026-08-05) with
`bb3313bc` and `origin/main` `f97ba659` read for comparison · `sentinel` `f2c46190` · `cms` `ca50c817` ·
`jobsimulation` `462343b0` · `messenger` `fa47850d` · `storage` `4ce8ece5` · `roadrunner` `87d8d443`.

**Line-number note:** `dependency_map.md` ends at **delta 0** — every anchor keeps its number
(`:58`, `:59`, `:103` still name the same cells). See *Induction caught and eliminated* below; this was
not free. `shared_libraries.md` +27 and `frontend_architecture.md` +11, both verified safe (no corpus
file cites either by a line number below the edit).

---

## Ledger rows

| # | the false claim | anchor | what is true | sites repaired |
|---|---|---|---|---|
| 1 | "That one is not opt-in: it is the stock `core` selection." | `corpus/architecture/dependency_map.md:58` | **It is opt-in and OFF by default.** The entire messenger-in-app subscriber-server block sits behind `if messengerEnabled {` at `app/main.go:1445` @ `ad9f3c49` (identically at `2035f9a4`), fed by `mustSubsystemSwitch(envMessengerEnabled)` at `:285`, where `env_guards.go:61` is `envMessengerEnabled = "MESSENGER_ENABLED"`. The code states it in-line at `main.go:1437`: *"ALL OF IT is behind MESSENGER_ENABLED."* Compose deliberately sets nothing and says why where the variables would have gone — `docker-compose.yml:84-92` @ `0c91421`: *"which default to OFF on a developer machine."* The cell contradicted `:21` of its own file. **Companion defect fixed in the same clause:** the cited `app/main.go:1442` is a comment line at both newer refs and **does not exist at all** at `b948604f` (`main.go` is 1361 lines there). | 3 |
| 2 | "`SKILLER_STREAM` has **6** Go occurrences across 4 files, not one" | `corpus/architecture/dependency_map.md:59` | **6 occurrences across 3 files.** `git grep -n SKILLER_STREAM ad9f3c49 -- '*.go'` → 6 lines; `git grep -l` → 3 files (`main.go`, `subscriber_merge_test.go`, `subscriber_wiring.go`); identical at `2035f9a4`. Off the `*.go` pathspec: **6 files**. **No scope yields 4.** The occurrence count 6 was always right — that half was preserved. Defect **induced by iter-98's own repair**. | 1 |
| 3 | "**not one of those line numbers resolves**" | `corpus/architecture/dependency_map.md:59` (**second, distinct** predicate in the same cell — kept separate, per adj-4) | **6 of 7 drift; the seventh HOLDS.** At `ad9f3c49` (and `2035f9a4`): `main.go:287` = `logger.Info("subsystem switches",`, `:637` = `)`, `:1039` = `serverContext,`, `:1276` = `apiKeyManager,`, `wiring.go:127` = an asynq-client construction, `:180` = blank. But `internal/roles/roles.go:791` = `func (r *RoleManager) SkillerSubscriber() *pubsub.Subscriber {` — **byte-identical at all three refs**. A bolded universal falsified by a member of its own enumerated set. | 1 |
| 4 | "This cell also named origin/main `2035f9a4`" — CANON-3, the currency pin | `corpus/architecture/dependency_map.md:59` | `app` `origin/main` is **`ad9f3c49`** as of 2026-08-06; `2035f9a4` is 5 commits behind. Applied CANON-3 move 1: the sha is kept (a pin is a pin — every claim at `2035f9a4` still holds there), the moving label is retired. The cell now reads `` at `app` `2035f9a4`, and identically at `ad9f3c49` (`origin/main` on 2026-08-06) ``. | 1 |
| 5 | "there is exactly **one** left, `AUTHORIZATION_ADDRESS`" — CANON-1, in its *"compose sets exactly one service address"* generalised form | `corpus/architecture/dependency_map.md:7` (**not in the booking — found by predicate expansion inside my own file**) | `backend`'s `environment:` block @ `0c91421` carries **six** service addresses: `AUTHORIZATION_ADDRESS` (`:48`), **`GOTENBERG_URL=http://gotenberg:3200` (`:57`)**, `JUDGE0_BASE_URL` (`:59`), `REDIS_ADDR` (`:66`), `SUPABASE_DB_CONN` (`:93`), `COPILOT_DB_CONN` (`:94`). `gotenberg` is on the **default** `core` profile (`:183`) and is reached over **plain HTTP**, not Connect-RPC (`app/internal/converter/gotenberg.go:31` @ `ad9f3c49`). Canonical form applied. **The `*_RPC_ADDR`-is-zero clause is TRUE — verified (0 hits in `docker-compose.yml` @ `0c91421`) and preserved verbatim.** | 1 |
| 6 | "that shared plumbing lives in **five small repos that the services pull in** like any third-party dependency" · "The Anthropos Go services share **five** internal libraries" | `corpus/architecture/shared_libraries.md:11` and `:3` | **Four of the five are pulled by at least one repo; only three by the repos a stack builds.** `go.mod` requires over the 7 Go repos on disk at their pinned refs: colony **7/7**, proto **7/7**, taxonomy **6/7** (not roadrunner), ai **2/7**, **authn 0/7**. `authn` ships inside colony as `colony/authn` (0 `go.mod` requires *and* 0 Go-source imports of the standalone module, all 7 repos). **New at this iter:** `app` **dropped** `github.com/anthropos-work/ai` at `1e457fa70` (2026-08-04, *"refactor(ai): fold the ai library into app as internal/ai"*) — `ad9f3c49:go.mod` has no such line and `go.sum` has zero — so `ai` is now pulled by **no repo a stack builds**, only by the frozen `cms`/`jobsimulation`. | 6 |
| 7 | "Vendor selection lives in each consumer's own `internal/ai/ai.go` wrapper: an EU Azure client by default, a US Azure client gated by the PostHog flag `flag_use_azure_us`, and an Azure→direct-OpenAI fallback on HTTP 429." | `corpus/architecture/shared_libraries.md:130-131` (now `:141-157`) | **No file called `internal/ai/ai.go` contains any of that, at any ref.** At `b948604f` `app/internal/ai/` does not exist. At `2035f9a4`/`ad9f3c49` `app/internal/ai/ai.go` is **21 lines** — `type AI interface` + `type TokenEncoder interface`, nothing else. The mechanics are in **two** files: `internal/jobsimulation/ai/ai.go` (`getClient` `:259`, flag `:267` and `:344`, `isThrottlingError` `:129` used at `:166`/`:325`) and `internal/skillerai/ai.go` (`getClient` `:332`, flag `:347`, `isThrottlingError` `:128` used at `:176`). The 429 path is a **retry target** (`vendor = Openai` on the next attempt), not a fallback rung — restated as such. **The `external_services.md:579` "no EU-first ladder" cross-reference is TRUE (verified: `:579` reads *"There is **no ordered EU-first fallback chain.**"*) and was preserved.** | 1 |
| 8 | "**but there are direct REST/SSE calls**, **29 of them across 21 non-test files**" (with the four core-js clients glossed as *"12 sites between them"*) | `corpus/architecture/frontend_architecture.md:39` | **Upheld — see the `:39` verdict below.** 29 and 12 are counts of `NEXT_PUBLIC_BACKEND_API_URL` **occurrences**, not of calls; the sentence's own noun is "calls". Re-derived at **`8297c684`**: 22 non-test source files, **31** env-var occurrences, **47** `fetch(` call sites; core-js four clients **12 env / 27 calls**. `new EventSource(` is **0** everywhere. | 1 |

**CANON-1 verify-only anchor — `corpus/architecture/dependency_map.md:103`: VERIFIED CORRECT, NOT
REWRITTEN.** It reads `` *   `GOTENBERG_URL=http://gotenberg:3200` is injected via the backend's compose
env. `` — exact against `docker-compose.yml:57` @ `0c91421`. Its neighbour `:102` (Gotenberg's
`/forms/libreoffice/convert` endpoint, `app/internal/converter/gotenberg.go`) is exact too, against
`ad9f3c49:internal/converter/gotenberg.go:31`. **The finding is the split within one file:**
`dependency_map.md` stated the gotenberg edge correctly in §6 while its own header line `:7` asserted
there was *"exactly one"* service address — 96 lines apart, in the same document. The repaired `:7` now
points a reader at §6 rather than contradicting it.

---

## Reach — graded, not asserted

| predicate | anchors BOOKED | sites FOUND | sites REPAIRED | how I searched |
|---|---|---|---|---|
| `messenger-not-opt-in` | 1 (`dependency_map.md:58`) | **3** | **3** | `git grep -nE 'opt-in\|MESSENGER_ENABLED'` at HEAD over my 3 files, then read the whole Event-Streams table. `:58` is the independent assertion; `:60` (jobsimulation) and `:61` (cms) said *"the same takeover as the `backend` row"* — dependent sites that would have inherited the retracted claim by reference. Both given the gate explicitly. |
| `skiller-stream-file-count` | 1 | 1 | 1 | `git grep -n/-l SKILLER_STREAM <ref> -- '*.go'` at all three refs, plus the no-pathspec control (6 files) to confirm no scope yields 4. |
| `no-line-number-resolves` | 1 (same cell) | 1 | 1 | Resolved all 7 enumerated anchors at `b948604f`, `2035f9a4` and `ad9f3c49` by `git show <ref>:<file> \| sed -n Np`. Kept separate from the row above, per adj-4. |
| CANON-3 currency pin | 1 | 1 | 1 | `git grep -n 'origin/main\|2035f9a'` over my 3 files. Exactly one labelled site; no twins. |
| CANON-1 `sentinel-only-cross-process-edge` | 1 (verify-only, `:103`) | **2** | **1 repaired + 1 verified** | Booked as verify-only. Verified `:103` correct. Then searched the file for the *fact* rather than the phrasing and found the generalisation restated at `:7` as *"exactly **one** left"* — **a paraphrase site the booking did not carry.** Applied the canonical wording. |
| `five-shared-library-repos` | 2 (`shared_libraries.md:3`, `:11`) | **6** | **6** | Booked 2; the predicate is *"how many of the five do services actually pull"*, so I re-derived it from every repo's `go.mod` and then swept both files for every place that answers it. Beyond `:3` and `:11`: `shared_libraries.md:38` (*"Each is pulled as a private Go module"* — the mechanism sentence, false for `ai` and `authn`), `shared_libraries.md:122` + `:123` (the `ai` section's **Version pin** and **Imported by** rows, both naming `app`), and `dependency_map.md:42` + the `ai`/`authn` rows at `:48`/`:49`. **Booked width 2, live width 6 — a 3× under-count, exactly the ratio the brief predicts.** |
| `vendor-selection-path` | 1 | 1 | 1 | `git grep -n "internal/ai/ai.go"` over my files (1 hit), then resolved the named file at all three refs and enumerated `flag_use_azure_us` / `isThrottlingError` across `app`'s Go source to locate the real sites. |
| `frontend-rest-call-count` | 1 | 1 | 1 | Reproduced **both** adjudicators' derivations at `bb3313bc` before touching anything (see verdict), then re-derived at `8297c684` and `f97ba659`. |
| **TOTAL** | **8 booked** (incl. 1 verify-only) | **16 found** | **15 repaired + 1 verified** | |

---

## The `frontend_architecture.md:39` verdict — **UPHELD; repaired**

The instruction was to re-derive against `8297c684` and to **decline with evidence** if the defect did
not survive the 41-commit / 192-file move. It survives. Reasoning, with the numbers:

**Step 1 — I reproduced both sides of the disclosed disagreement at `bb3313bc`, exactly.**

| derivation | result at `bb3313bc` | whose reading |
|---|---|---|
| non-test `.ts`/`.tsx` files naming `NEXT_PUBLIC_BACKEND_API_URL` (`e2e/` excluded) | **21** | both |
| **occurrences** of that env var in those files | **29** | reading #24's clearance — *exact* |
| `fetch(` + `new EventSource(` **call sites** in those same files | **43** | adj-1's refutation — *exact* |
| the four `packages/core-js` clients | **12** env / **25** calls | adj-1's *"12 sites … 25 outbound calls"* — *exact* |

Both adjudicators computed correctly. They measured **different quantities**, and the sentence names
only one of them: its grammatical subject is *"direct REST/SSE calls, **29 of them**"*. The number
attached to that noun is an env-var occurrence count. adj-1's grading — *"the arithmetic is right and
the predicate is wrong"* — is the correct disposition.

**Step 2 — does a 192-file move invalidate it?** It moves the numbers and leaves the defect untouched:

| quantity | `bb3313bc` | `8297c684` (pin) | `f97ba659` (origin/main) |
|---|---|---|---|
| files | 21 | **22** | 22 |
| env-var occurrences | 29 | **31** | 31 |
| `fetch(` call sites | 43 | **47** | 47 |
| core-js four clients | 12 env / 25 calls | **12 env / 27 calls** | — |

The move added one file (`apps/web/src/components/organisms/Workforce/MemberAnalyticsContainer.tsx`)
and two calls in `coursebuilder/api.ts`. The *conflation* is unchanged, and the gap between the two
quantities actually widened (43→47 calls against 29→31 mentions). **A count claim about frontend call
sites is exactly the kind a 192-file move invalidates — but this claim was never a count of call sites,
which is the defect.** So the repair is not "fixing a true sentence"; it is naming which quantity is
which and giving both.

**Repair shape:** the prose sentence no longer carries a bare number (it now says *"spread over 22
non-test source files that name `NEXT_PUBLIC_BACKEND_API_URL`"*, and its enumeration was extended to
all 22 files). The two quantities are separated into an explicit table with the retracted wording quoted
verbatim above it. `new EventSource(` = 0 is recorded, with `packages/core-js/src/talkToData/api.ts:214`
— *"uses POST so we can send a JSON body — that rules out EventSource"* — as the reason SSE rides on
`fetch`. The old parenthetical's warning (*"treat any figure here as pinned, never current"*) was TRUE
and is kept, restated as iter-98 P6's rule.

---

## Induction caught and eliminated — the one that would have shipped

My first pass added a 2-line blockquote at `dependency_map.md:44`. That pushed every line below it down
by two — and **`corpus/services/hiring.md:39` cites `dependency_map.md:78`**, which is the sentence
*"the legacy `jobsimulation` schema is non-authoritative."* After my edit that content sat at `:80`.
**This is iter-100's failure mode reproduced exactly** (a parenthetical shifting a table and leaving
citations pointing at the wrong rows), in a file another seat owns and cannot be told to chase.

I found it by grepping the whole corpus for `(dependency_map|shared_libraries|frontend_architecture)\.md:[0-9]+`
before finishing, and eliminated it rather than reporting it: the blockquote was folded into the existing
`:42` paragraph, restoring `dependency_map.md` to **delta 0**. `:78` verified to point at the original
sentence again.

The other two live cross-file citations were checked and are safe: `service_taxonomy.md:185` →
`frontend_architecture.md:11` (above my edit, byte-identical), and `shared_libraries.md#taxonomy-figures`
(an `<a id>` anchor at `:228`, line-independent). `dependency_map.md`'s own self-reference at `:90`
(*"Consistent with :9/:15/:31 above"*) re-verified — all three still point where they did.

One further self-check: I introduced a wrong number into the new `frontend_architecture.md` table
(`3 env` for core-js at `bb3313bc`, where the measurement is `12 env`) and corrected it before finishing.
Recording it because the induction rate is the number this milestone tracks, and a caught induction is
still an induction attempt.

---

## Twins outside my files (REPORT, do not edit)

| predicate | file:line | why it is the same claim |
|---|---|---|
| `five-shared-library-repos` | **`CLAUDE.md:248`** | *"**five repos, of which four are imported as private Go modules** (`ai`, `colony`, `proto`, `taxonomy`…)"* — verbatim the same cardinality claim. Measured, `ai` is required by **no** repo a stack builds since `1e457fa70`, so the built-set figure is **three**. The *"`authn` is imported by nothing"* half of that same line is **TRUE** and must survive any repair. Orchestrator's file. |
| `five-shared-library-repos` | **`CLAUDE.md:413`** | *"only **four are imported as private modules** (`ai`, `colony`, `proto`, `taxonomy`)"* — the same sentence in the doc index. Same disposition. |
| `five-shared-library-repos` (the `ai`-has-a-live-importer half) | `corpus/architecture/service_taxonomy.md:148` + `:154` | Its Shared-Libraries table lists `ai` under *"imported as private Go modules … pulled at Docker build via `GH_PAT`/`GOPRIVATE`"*. Nothing a stack builds pulls it any more. Seat 3's file. |
| `five-shared-library-repos` (`ai` module currency) | `corpus/ops/demo/ai-generation-spec.md:47` | Pins the seeding module to *"the shared `ai` library (`github.com/anthropos-work/ai`, pinned `v1.40.1`)"*. **This one is probably still correct and is worth protecting**: `app/internal/ai/module_import_guard_test.go:15-17` says the repo *"was deliberately left in place because at least one consumer outside this codebase (anthropos-work/rosetta-extensions/stack-seeding) pins it."* The rext consumer is the reason the module still exists. Out of `corpus/architecture`+`corpus/services` scope. |
| `vendor-selection-path` | `corpus/ops/demo/ai-generation-spec.md:49` | *"the `internal/ai/ai.go` pattern (`shared_libraries.md` §ai + `ai_architecture.md`)"* — it cites, by name, the exact phrasing I have just retracted at its source. Out of scope; will read as a dangling pattern-reference until touched. |
| `vendor-selection-path` (one file short) | `corpus/architecture/external_services.md:581` | *"The real mechanics, all in `app/internal/jobsimulation/ai/ai.go`"* — measured, the identical three levers (EU-Azure default, `flag_use_azure_us`, 429→OpenAI) are **also** in `app/internal/skillerai/ai.go` (`:332`, `:347`, `:128`/`:176`). *"all in"* is a completeness claim that is one file short. Note `external_services.md:579`, which I cite and preserve, is **correct** — the defect is two lines below it. Seat unknown; not mine. |

---

## Noticed, not repaired

1. **`app` dropped the `ai` module — a platform fact the corpus has not absorbed.** `1e457fa70`
   (2026-08-04) folded `github.com/anthropos-work/ai` into `app/internal/ai` and removed the `go.mod`
   requirement, behind a one-way guard (`internal/ai/module_import_guard_test.go`,
   `TestNoExternalAIModuleImports`). I repaired the sites in my own files; the twins above are the
   remainder. **This landed inside the `b948604f`→`ad9f3c49` window and is not in either union** — it
   is a change of ground truth, not a mis-read, and it is the kind the `platform_alignment_guard`
   fence cannot see because `ai` was never in `repos.yml`.
2. **`dependency_map.md:18`** (the Roadrunner row) ends *"Prod terraform still reads `= 1`"* — a
   present-tense prod-terraform assertion of the CANON-2 family, unpinned and un-refed. I did **not**
   touch it: CANON-2 does not list it, and it is a different repo (`roadrunner`, not `cms`) from the
   one the canonical wording was derived for. Flagging it as a likely CANON-2 sibling for whoever
   owns that predicate's next pass.
3. **`clones.pin.json` now pins `app` to `ad9f3c49`**, i.e. the demo pin and `origin/main` have
   converged. The corpus's `b948604f` citations therefore no longer name "the ref this stack builds"
   even where they are otherwise correct. That is a labelling question across many files, well beyond
   this seat.

---

## What I could not settle, and why

1. **Whether the retraction quotes will trip `claim_twin_guard`.** Each repair quotes the false
   sentence verbatim inside a *"this previously read …"* clause — which is what makes the repair
   checkable and matches the house style already in `dependency_map.md:59`. A guard that greps for the
   false string tree-wide will hit these. I judged the auditability worth it, but I did not run the
   guard and cannot say which way it grades.
2. **The exact prior of `ai 3/7` vs my `ai 2/7`.** adj-2's iter-99 derivation recorded `ai 3/7`; I
   measure **2/7** at the current pinned refs. Both are right at their own ref — adj-2 graded with
   `app` at `b948604f`, where `app` still required the module. I state 2/7 with the refs named rather
   than "correcting" adj-2, because the difference is `1e457fa70` landing, not an arithmetic error.
3. **Nothing about production.** Where a claim in my files depends on the `infrastructure` repo it was
   left alone (TRAP A); no citation was repointed to make it resolve.
