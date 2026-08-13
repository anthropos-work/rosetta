# iter-98 blocker ledger — `FIX-M257x-iter97-read-union`, repaired BY PREDICATE

The repair ledger for the **20 anchors / 17 predicates** iter-97's reading returned, plus the four
instrument defects it reported out-of-scope. Written in the **ledger table shape**
`rosetta-extensions/stack-core/claim_ledger.py` derives from — a table under a non-Minor heading with a
claim-shaped column and an anchor-shaped column — so every refuted form below is adopted **automatically**
by `claim_twin_guard` and `repair_postcondition` and becomes un-republishable tree-wide.

That is the point of writing it here rather than in prose: iter-97's binding condition was *repair by
PREDICATE and enumerate PARAPHRASES*, and a fence is the only thing that holds a predicate down after the
repairing agent is gone. It is also the only thing that makes the completeness claim checkable rather than
asserted.

## BLOCKERS — the refuted claims

| # | the false claim | anchor | what is true | sites repaired |
|---|---|---|---|---|
| P1 | "a standalone PDF→markdown utility nothing dispatches" | `external_services.md:565` | `tools/r3.py` **does** dispatch it, as step 2 of the offline chain — `r3.py:139` (`scripts = ["any2pdf.py", "pdf2md.py", "md2cleanMd.py"]`), `:190`, `:199-206` @ `app/studio aeec036`. What holds is the narrower claim: no **Go** caller and no `gen.py` path reaches it. | 2 |
| P2 | "no Go caller exists (Go execs only `studio/gen.py`, `studioManager.go:119`)" | `cms.md:95` | Go execs **two** studio scripts: `studio/gen.py` at `studioManager.go:119` **and `studio/postgen.py` at `:1045`** @ `app b948604f`. Neither is `pdf2md.py`, so the conclusion survives — the enumeration did not. | 2 |
| P3 | "prod teardown **M810 — still pending**" (cms) | `backend.md:13` | The corpus rules cms's M810 prod state **unmeasurable** — `infrastructure` has never been in a clone set — and every other site says *report both, assert neither* (`:36` of this same file, `service_taxonomy.md`, the fenced map). One flat assertion in a table cell, contradicted by its own document 23 lines later. | 1 |
| P4 | "`docker-compose.yml` citations — 23 of them" / "every other `file:line` … (66 citations)" | `platform-migration-status.md:205` | The guard prints its own reach: **92 citations — 22 subject-checked, 69 range-only, 1 outside any service block**. The coverage figure had drifted from the coverage. Take the three numbers from the run, never from the table. | 2 |
| P5 | "`messenger/internal/flow/flow.go:70-95` adds a subscriber on the `backend` stream with 20 handlers" | `messenger.md:17` | **21** handlers over **`:72-104`** @ `fa47850` — 22 `pubsub.EventHandler(…)` lines of which one (`OrgJobSimulationAssignmentPastDueHandler`) is commented *"not implemented"*. `AddSubscriber` is at `:72`, the closing `))` at `:104`. `dependency_map.md:58` already said 21. | 3 |
| P6 | "`app` is at **`v1.363.2`** @ `5ba17044`" stated as CURRENT | `coursebuilder.md:130-131` | A version is **a reading at a ref, never a standing "current"**. `app` is **v1.369.0** @ origin/main `2035f9a4` (2026-08-06) — six releases on. `academy-backend.md` published `v1.363.2` at `:20` while citing `v1.367.0` at `:57`, one document disagreeing with itself. | 3 |
| P7 | "the whole repo contains one `roadrunner` mention" / "zero hits outside CHANGELOG" / "no other platform repo references roadrunner at all" | `roadrunner.md:114`, `roadrunner.md:24-26` | Scoped to **Go source** it is exactly one, and a comment (`internal/runner/runner.go:3`). Repo-wide it is **14 lines across 8 files** (5 exact-case across 3) at `jobsimulation 462343b0`. "Outside CHANGELOG" is false (`knowledge/operational.md:68`), and every other clone is false too: app **25** files, jobsimulation 8, platform 3, studio-desk 3, next-web-app 1. The three *strings* (`ROADRUNNER_RPC_ADDR`/`RoadRunnerService`/`roadrunner:10401`) are genuinely **0 in Go** at both repos. | 2 |
| P8 | "repo … ARCHIVED `<date>`" asserted flat | `service_taxonomy.md:128`, `:129`, `:364` | Archive state lives in the **GitHub org API**, never in git objects — **no clone can measure it**, re-checked here: `gh` is not installed on this host and the repos are private, so even the anonymous REST path is closed. Every date is a **dated snapshot carrying an expiry**, which the `Jobsimulation` row proves by having been refuted by four post-dated commits. This table published the flat form two rows above the cell retracting it. | 2 |
| P9 | "DM Sans + Instrument Serif + JetBrains Mono (via `next/font/google`)" | `ant-academy.md:255` | `code/app/layout.jsx:1` imports exactly `{ Work_Sans, Instrument_Serif }` (`:41`, `:51`). **Neither DM Sans nor JetBrains Mono is loaded**; `code/academy.css:1` says display + mono usages *"fall back to system fonts."* Two of the three named fonts were wrong. | 1 |
| P10 | "the cockpit … **no longer sets the `e2e_persona` cookie**" | `ant-academy.md:322-326` | False in **both** directions, and the passage contradicted itself twice in nine lines. **The cockpit DOES still set it, at two live paths** — `demo-stack/cockpit.py:812` (`_ACADEMY_JS`, client-side) and `:1496` (`Set-Cookie` on the `/go` 302); `:327` names both. **And the BYPASS is gone from the academy launch env** (`demo-stack/ant-academy.sh:576-583`, fenced by two tests), because the demo academy now gets real Clerkenstein keys — so the cookie is **set and not honoured**. The whole documented auth model had inverted. | 4 |
| P11 | "Enumerated over every publisher constructor in `app` @ `b948604f` *and* origin/main `2035f9a4` … its one Go occurrence, `main.go:1276`, is an `AddSubscriber`" | `dependency_map.md:59` | **Not one of those line numbers resolves at the second ref** (`2035f9a4:main.go:1276` is `apiKeyManager,`) and `SKILLER_STREAM` has **6** Go occurrences across 4 files there, not one. A block naming two refs is `ambiguous` to the citation resolver besides. The consumer-only FINDING holds at both refs — the enumeration was pinned to one. | 1 |
| P12 | "not in the top-level `migrations/` dir (which holds only `atlas.sum`)" | `backend.md:299` | **There is no top-level `migrations/` dir.** `6a46e8445` (2026-06-18, *"chore(migrations): remove obsolete atlas.sum file"*) deleted it; `git ls-tree b948604f migrations/` is empty, as at origin/main `2035f9a4`. | 1 |
| P13 | "**Seven** services … are folded into `app` and **no longer deploy separately**" | `services/README.md:11-13` | "No longer deploy separately" is a claim about **prod**, and is not uniformly measurable: for `customerio-sync` the standalone's terraform lives in a repo that has never been in any clone set, so the fenced map asserts its prod half **from `app`'s side only** (`platform-migration-status.md:101`). The LOCAL claim — no compose service, no container — holds for all seven. | 1 |
| P14 | "`STORAGE_RPC_ADDR` … 3 hits **repo-wide** at `9d00a313`, every one a comment" | `architecture_overview.md:334-336` | **3 hits in GO SOURCE**, every one a comment (`main.go:451`, `internal/jobsimwiring/wiring.go:101`, `internal/storagens/callsites_test.go:189`). Repo-wide the same ref returns **29 lines across 18 files**. The Go scope is what carries the claim; "repo-wide" was a mis-transcription of it that made a true claim unverifiable. | 1 |
| P15 | a load-bearing intra-corpus `:N` self-citation offered AS the evidence, resolving to a **different construct** | `graphql-wundergraph.md:85-87`, `hiring.md:267-269`, `external_services.md:614`, `platform-alignment.md:1345` | `:174-176` is the compose line-number caveat, not the `localhost:5050` claim (**`:193`**). `:157-159` is the `job_position` bullet, not the mirror/sessions form (**`:170-175`**). `:543` is the *"When backend services add new GraphQL types"* heading — iter-96 moved that construct 543 → **567** and re-pointed only the three CROSS-FILE citations, leaving this doc's own in-file anchor behind. `backend.md:187` is now a directory listing; the messenger-four-addresses claim is at **`:241`**. | 4 |
| P16 | "**12** Connect-RPC services are defined" | `shared_libraries.md:77` | **At least 13** — the list omitted `StorageService`, which `storage.md:115` documents in full. And it is a **floor, not a count**: `proto` is a private Go module **in no clone set**, so the enumeration is hand-assembled from consumers and cannot be checked against its source of truth. | 2 |
| P17 | "**`app` owns the `skiller`, `skillpath`, `jobsimulation`, `cms` and `ai_usage` Redis Streams** — both producer and consumer are in-process" | `backend.md:33-34` | **Four of the five, not five.** At `b948604f` each of `SKILLPATH_STREAM` (`main.go:637`/`:1274`), `CMS_STREAM` (`:1039`/`:1303`), `JOBSIMULATION_STREAM` (`wiring.go:180`/`main.go:1285`) and `AI_USAGE_STREAM` (`wiring.go:127`/`main.go:1305`) has a publisher **and** a subscriber. `SKILLER_STREAM` has **only** `AddSubscriber` at `:1276` — no `NewPublisher` names it, at `b948604f` or at `2035f9a4`. The same file already said so at `:136` and `:264`. | 1 |
| C1 | "`next-web-app/apps/web/src/hooks/useCoursebuilder.ts` (50,433 bytes, **1,178 NULs**)" | `platform-alignment.md:1236` | The file has **1** NUL byte. **1,178 is its LINE count** — produced by `grep -c $'\x00'`, where `grep -c` counts matching *lines* and the zsh `$'\x00'` pattern degenerates to empty and matches every line. `store.js` likewise has **1**. **The rule about lying instruments was written with a lying instrument.** Count bytes: `tr -dc '\000' < FILE \| wc -c`. | 1 |
| C2 | the rule's own shell recipe — `printf '… hits=%-4s' "$(git -C "$d" grep -c "$TERM" HEAD \| wc -l)"` | `platform-alignment.md:1254-1260` | Wrong three ways, each flattering the result: it labels a **file** count as `hits=` (`grep -c` emits `file:count`, so `wc -l` counts files); it **drops `-i`**; and its last line sits **after `done`**, where `$d` is unbound. Against the rule's own worked example (`TERM=mistral`, `app/studio` @ `aeec036`) the printed form returns **2** where the prose publishes **22**. The corrected form returns **22 lines / 3 files**. | 1 |
| C3 | an `app/main.go:504` citation carrying **no ref**, inside a sentence that names `b948604f` for a different anchor (**deliberately not quoted** — the defect is a MISSING pin, not a false sentence, so there is no refuted form to fence and this row correctly derives none) | `CLAUDE.md:280` | `:504` resolves **only at origin/main `2035f9a4`**; at `b948604f` it is an unrelated jobsim-in-app comment. One sentence, two `app` refs, one of them implicit — the `ambiguous`-block defect the fenced map warns about, in the repo's own entry doc. | 1 |
| C4 | "**`S3-private` was in this row and has been REMOVED**" | `safety.md:207` | The **code still classes it `PerStackIsolated`** (`stack-seeding/isolation/isolation.go:106`). The row asserted a registry change that never happened. Withdrawn rather than made true: re-classing the store would resolve the user's open escalation `DEF-M257x-iter80-storage-prod-bucket`, which is not this iter's call. The doc and the registry now **state** their disagreement. (`safety.md:203`'s companion claim about `audit.go:146` re-derived and **correct** — the public bucket is the only forced override.) | 1 |

## The five INDUCED citation repairs, caught inside this iter

The defect iter-96 named (`D-M257x-96-5`) and then shipped: *a prose repair is a line-number edit, and only
half of that is fenced.* Every edit here that changed a file's line count triggered an immediate re-derivation
of that file's inbound and **in-file** citations, per binding condition 2. Five moved, and all five were
caught **before commit** rather than by the next reading:

| file | citation | was | now | how it moved |
|---|---|---|---|---|
| `academy-backend.md` | in-file, to the `v1.367.0` passage | `:55` | `:57` | my own P6 edit added 2 lines above it — written wrong, re-derived, fixed |
| `roadrunner.md` | in-file, to *"Upstream consumers"* | `:118` | `:124` | my own P7 edit added 7 lines above it |
| `graphql-wundergraph.md` | in-file, to the `localhost:5050` claim | `:192` | `:193` | **my own P15 fix moved its own target** by one line |
| `service_taxonomy.md` | in-file, to the archive rows | `:128`/`:129` | `:137`/`:138`/`:139` | my own P8 note added 9 lines above the table |
| `external_services.md:144` | **cross-file**, to `service_taxonomy.md` | `:321-330` | `:332-339` | the same P8 note; the only cross-file one |

Three of the five are the case iter-96 missed entirely (**in-file self-anchors**), and one is sharper still —
a citation fix that invalidated the number it was writing. That is not a near-miss to be glossed: it is the
strongest available evidence that **the re-derivation must run after the edit, never from the pre-edit
reading**, because the edit is part of the input.

## What "by predicate + paraphrase" bought, measured

| | count |
|---|---|
| anchors iter-97 booked | **20** |
| distinct predicates | **17** (+4 instrument defects, C1–C4) |
| **predicate sites repaired** | **37** |
| sites an anchor-wise repair would have left standing | **17** |
| induced-citation repairs, self-caught in-iter | **5** |
| total hunks (`git diff --unified=0 \| grep -c '^@@'`) | **42** = 37 + 5 ✓ |
| files touched | 22 (+167 / −80) |

**The multiplier fell, and that is a finding rather than a shortfall.** iter-96: 13 anchors → 51 sites
(**3.9×**). iter-98: 20 anchors → 37 sites (**1.85×**). The difference is not method — the method got
*stronger*, gaining the paraphrase axis. It is **composition**: iter-96's single largest predicate
(`mistralai`) had **11** sites, and iter-95's `storage.md` predicate had 10. **This iter's widest predicate
had 4.** The wide, highly-propagated predicates have been drained; what remains propagates narrowly.

## Fences this repair grew, rather than prose it rewrote

- **This ledger itself** — `claim_ledger.py` derives its claim set from ledger-shaped tables, so the **21
  refuted forms** above are adopted by `claim_twin_guard` tree-wide and by `repair_postcondition` at the
  commit. Completeness is fenced, not claimed.
- **`platform-migration-status.md` now prints the guard command that produces its own coverage numbers**,
  with both paths relative to the same root — the recipe it replaced named the guard relative to a stack
  workspace and its two arguments relative to the rosetta root, so pasting it anywhere resolved at most one
  of the three. Re-run verbatim; it reproduces.
- **Rule 44's recipe now reproduces rule 44's own worked example** (22 lines / 3 files). It did not before.
