# seat-9 report

**Files owned:** `corpus/services/{storage,messenger,customerio-sync,skillpath,gotenberg}.md`.
**Anchors booked:** 8 (5 repair + 1 CANON-1 verify-only + the `messenger.md:22` positive-control re-derivation).
**Sites found:** 15. **Sites repaired:** 14. **Verified-correct, deliberately not rewritten:** 2.

**Ground truth used.** platform `0c91421d` (== `git ls-remote origin HEAD`) · `app` `ad9f3c49`
(**HEAD == `origin/main` == the demo build pin** — `stack-demo/clones.pin.json` now pins `ad9f3c49`, so the
`b948604f` / `2035f9a4` split that produced iter-101's readings has **collapsed to one tree**) ·
`next-web-app` clone pin `8297c684` (origin/main has moved on to `f97ba659`) · `messenger` clone
`fa47850d`, origin/main `e9421c68` · `storage` clone `4ce8ece5`, origin/main `9f8cb532`. No fetch was
issued; no clone was written.

---

## Ledger rows

| # | the false claim | anchor | what is true | sites repaired |
|---|---|---|---|---|
| 1 | "`ENVIRONMENT` \| (empty) \| Environment name" — under a column whose contract is *"what `docker-compose.yml` set on the `storage` service block"* | `corpus/services/storage.md:215` | The block **did** set it: `- ENVIRONMENT=development` at `docker-compose.yml:119` @ platform `0dab54d` and `:206` @ `2adcf71` (verified by `git show <ref>:docker-compose.yml`; the two blocks are byte-identical — `diff` of the `storage:` stanza at `0dab54d` vs `838d907^` is empty, and `0dab54d..838d907` is a **single** commit, so `0dab54d` **is** the last ref that had the block). Only this one row of eight was wrong; `PORT`/`RPC_PORT`/`STORAGE_S3_*`/`AWS_*`/`SERVICE_NAME`/`SENTRY_DSN` all re-derive correct. **The error was load-bearing, not cosmetic:** `development` is the exact value that makes `deployedEnvironment()` return false and *disarms* app's boot guards — the mechanism the same file's HAZARD block depends on — so recording it as never-set hid the causal link. | 2 (the row + the table banner, which now names its measuring ref and states that *(empty)* means "the block did not set it") |
| 2 | **[quoted form deliberately withdrawn — the defect is an ABSENCE, not a false string.]** The nothing-warns sentence naming the two boot guards at `main.go:518-523` / `:529-535` and `app/env_guards.go:37-44` was published **unpinned and present-tense**. All three anchors name the right constructs at `ad9f3c49` and `2035f9a4`; they fail only at the demo's *former* pin `b948604f`, where `env_guards.go` does not exist. Repaired by **pinning**, so the repaired text legitimately still contains the same anchor strings — a verbatim-quote fence cannot represent a missing-ref defect, and quoting the true headline around it would fence a true sentence. Disclosed into `claim_ledger`'s *"quoted no refuted form"* bucket rather than fenced wrongly. | `corpus/services/storage.md:73-75` (HAZARD block `:60-80`) | **All three anchors resolve, and name exactly the right constructs, at `app` `ad9f3c49`** — `:518-523` is the `log.Fatalf` empty-bucket guard, `:529-535` the `verifyBucketAccess` guard, `env_guards.go:37-44` is `func deployedEnvironment()`. Identical at `2035f9a4` (`git diff --stat 2035f9a4 ad9f3c49 -- main.go env_guards.go` is **empty**). They were false **only at the demo's former pin `b948604f`**, where `ls-tree b948604f -- env_guards.go` is empty and `:518-523`/`:529-535` are the public-storage-clients and academy-asset-uploader blocks (both re-derived here). **The defect is therefore the MISSING REF, not the line numbers** — repairing by moving the anchors would have been wrong. Ref pinned, the former-pin failure recorded, and the disarming value cited on the compose side (`docker-compose.yml:56` @ `0c91421`). | 1 |
| 3 | "`STORAGE_RPC_ADDR` is read by **nothing** *at `app` origin/main*" **and** "run the same grep ref-less, on the older `app` checkout a demo pins, and it returns **15 hits, 7 of them live env reads**" | `corpus/services/storage.md:29` (CANON-3) | Two separate expiries in one cell. (a) **CANON-3:** `2035f9a` is no longer `origin/main`; `ad9f3c49` is. The pin is kept, the moving label dropped. (b) **The warning's premise expired with the clone:** the demo no longer pins `b948604f` — it pins `ad9f3c49`, where the ref-less grep returns the same **3 comment hits** (measured), not 15. The 15/7 figure is correct **at `b948604f`** (re-derived: 15 lines, 7 of them `os.Getenv`/`getenv` reads) and is now stated historically. The cell's self-imposed "never name two refs" rule was kept honest by measuring that the anchors do **not** differ across them. | 2 (`:29` + the table header at `:23`) |
| 4 | "at the demo's pinned build ref `b948604` the same guard still reports a mid-fold with six read sites" | `corpus/services/storage.md:33-36` (paraphrase twin of #3, same file) | `b948604` is the demo's **former** build pin. Restated as a two-ref contrast with the numbers I measured (15 hits / 7 env reads at `b948604f`; 3 comment hits at `ad9f3c49`) rather than re-publishing the unverified "six read sites". | 1 |
| 5 | "(`app/env_guards.go:61`), which defaults to **off** on a developer machine." — unpinned, present-tense | `corpus/services/messenger.md:53` | `env_guards.go:61` = `envMessengerEnabled = "MESSENGER_ENABLED"` at `ad9f3c49` **and** `2035f9a4`; the file **does not exist** at `b948604f`. Ref pinned; the former-pin absence recorded. | 1 |
| 6 | "`app/main.go:295` `log.Fatalf`s if `MESSENGER_ENABLED` is on with an empty key." — unpinned, present-tense (twin of #5) | `corpus/services/messenger.md:149` | At `ad9f3c49` the guard is `main.go:295-300`: `:295` is the **condition** (`(messengerEnabled \|\| customerIOSyncEnabled) && os.Getenv("BREVO_KEY") == ""`) and `:296` the `log.Fatalf`. At `b948604f`, `:295` is `if err != nil {` in the Azure-OpenAI client init — a different construct. Ref pinned, range corrected, and the true-but-partial condition widened to name `CUSTOMERIO_SYNC_ENABLED` as well. | 1 |
| 7 | "(`app/main.go:15`, `:62`, `:63` @ `app` origin/main)" and "Every anchor in this row moved between `9d00a313` and origin/main `2035f9a`" | `corpus/services/messenger.md:43` (CANON-3) | **This row is the clearest specimen of the stale-currency-pin class, exactly as briefed.** iter-100 repaired it once — correctly diagnosing that the bare `9d00a313` pinned the row to the ref its anchors had moved *away from* — but repaired it **to a moving label** (`origin/main`) rather than to the sha that label then denoted. On 2026-08-06 `origin/main` advanced `2035f9a4 → ad9f3c49` (5 commits) and the same line went stale a **second** time, one iteration later. The anchors themselves never moved: `git diff --stat 2035f9a4 ad9f3c49 -- main.go` is empty and `:15`/`:62`/`:63`/`:1450`/`:1471`/`:1473`/`:1416-1421` all re-derive at both refs. Repaired to the sha, with the two-repairs-of-one-line history recorded in place. | 1 |
| 8 | "* [External Services](../architecture/external_services.md) — Customer.io as an integrated SaaS" | `corpus/services/customerio-sync.md:140` | False in **both** directions, re-derived at corpus HEAD. `external_services.md` has **no** Customer.io section and **no** Brevo section — its per-service `##` sections are Clerk, Directus, the Cosmo router, AI Providers, LiveKit, AWS Chime — and `git grep -ic brevo` over the file returns **0**. It also contradicted this file's own fossil-name banner 122 lines above (*"the destination has been **Brevo**, not Customer.io"*). **Not re-anchored** (TRAP A): the link target is a real, correct related doc, so the *gloss* was corrected and a second bullet added pointing at `messenger.md`, which genuinely documents the Brevo client (13 `brevo` lines). | 1 |
| 9 | "re-derived at `app` origin/main `2035f9a`" | `corpus/services/skillpath.md:35` (CANON-3) | `app/CLAUDE.md:109` and `app/knowledge/architecture.md:28` **both still list `SkillPathSessionService`** at `ad9f3c49` (re-derived; `CLAUDE.md` is one of the 5 files the 5 new commits touched, so this needed checking rather than assuming — the line did not shift). `SkillPathSessionService` has **0** occurrences in `*.go` at `ad9f3c49`, so the bullet's headline stands. Label replaced by the sha it denoted plus the current one. | 1 |
| 10 | "(verified against `next-web-app` `origin/main`)" — a **bare** moving label naming no sha at all | `corpus/services/skillpath.md:97` (paraphrase twin of #9) | Settled by the **demo build pin** `8297c684` (the claim's subject is what a stack renders), where every clause re-derives: `InsightsBySkillPathStudentSimulationsContainer.tsx:31-34` returns `null as unknown as MembershipEnriched`, the results `<Table` is commented out at `:140`, and `:138` renders `t('enterprise.insights.comingSoon')` → `configs/i18n/messages/en/enterprise.json:2101` = `"Coming soon"`. `next-web-app`'s `origin/main` has meanwhile advanced to `f97ba659`, which is precisely why a bare label was the defect. | 1 |
| 11 | (predicate-width find, **not booked**) seven unpinned present-tense `app` citations — `doc.go:4`, `env_guards.go:62`, `main.go:286`, `:395`, `:393-396`, `:295`, `:284` | `corpus/services/customerio-sync.md:9`,`:12`,`:21`,`:53`,`:114`,`:131` | Same predicate as #5/#6 (`messenger-unpinned-anchors`), in a file the union never booked. All re-derived at `ad9f3c49`: `env_guards.go:62` = `envCustomerIOSyncEnabled`, `main.go:286` = the `mustSubsystemSwitch` call, `:395` = `customeriosync.New(logger, copilotDB, os.Getenv("BREVO_KEY"))`, `:393-396` = the shared-`copilotDB` construction — all correct, all previously unpinned. **Two named the wrong construct:** `doc.go:4` (the quoted sentence wraps to `:5`) and `main.go:284`, which is only a **comment** — the deployed-unset-is-fatal mechanism is `env_guards.go:98-104` via `mustSubsystemSwitch`'s `log.Fatalf` at `:87`. File-level ref pin added; both construct errors corrected. | 6 |

---

## Reach — graded, not asserted

| predicate | anchors BOOKED | sites FOUND | sites REPAIRED | how I searched |
|---|---|---|---|---|
| `storage-env-compose-value` | 1 | 2 | 2 | `grep -n 'ENVIRONMENT'` over all 5 owned files (1 hit outside storage.md, at `customerio-sync.md:13`, and it is TRUE — it states the app-side consequence, not a compose value); then row-by-row re-derivation of **all eight** table rows against `git show 0dab54d:docker-compose.yml` and `2adcf71:docker-compose.yml` |
| `storage-boot-guard-anchors` | 1 | 1 | 1 | `git grep -n 'env_guards\.go'` and `'main\.go:518\|main\.go:529'` over `corpus/` + `CLAUDE.md`. The two `main.go` boot-guard anchors are cited **nowhere else in the corpus**; `env_guards.go:37-44` also appears at `service_taxonomy.md:106`, already pinned `@ ad9f3c49` by another seat (consistent — no action) |
| CANON-3 (currency pin) | 3 (`storage.md:29`, `messenger.md:43`, `skillpath.md:35`) | **6** | 6 | Per-file `git grep -n` at HEAD for `origin/main`, `2035f9a`, `b948604` across all 5 owned files — run **one file per invocation** after a multi-path `git grep ... -- $F` silently returned empty (§5 rule 44 in the wild: the multi-pathspec form under this shell produced a false negative on the first try). Found 3 booked + 3 unbooked: `storage.md:23` (table header), `storage.md:33-36` (stale demo-pin paraphrase), `skillpath.md:97` (a **bare** `origin/main` with no sha). **Booked width 3, live width 6 — 2×, consistent with the brief's "assume ~3× wider"** |
| `messenger-unpinned-anchors` | 2 (`messenger.md:53`, `:149`) | **8** | 8 | `git grep -n 'main\.go:295\|main\.go:284'` + `'env_guards\.go'` corpus-wide; the 6 extra sites are all in `customerio-sync.md`, a file this predicate was never booked against. **Booked width 2, live width 8 — 4×** |
| `customerio-external-services-section` | 1 | 1 | 1 | `git grep -rn 'Customer.io as an integrated SaaS'` and `'external_services.md).*[Cc]ustomer'` over `corpus/` + `CLAUDE.md` → exactly one site, confirmed by two independent patterns. Negative half cross-checked two ways: `git grep -n '^## '` enumerating the target's headings, and `git grep -cin 'brevo'` → **0** |
| CANON-1 @ `gotenberg.md:50` | 1 (**verify only**) | 1 | **0 — verified correct, deliberately not rewritten** | see below |

**Totals: 8 anchors booked → 15 sites found → 14 repaired + 1 verified-correct.**

### CANON-1 verification result — `gotenberg.md:50`

> `* **Env var**: `GOTENBERG_URL=http://gotenberg:3200` (injected via the backend's compose `environment:`)`

**CORRECT as written, at platform `0c91421`.** `docker-compose.yml:57` is `- GOTENBERG_URL=http://gotenberg:3200`,
inside the `backend` service's `environment:` block (which runs `:46`–`:94`). Not rewritten.

Two adjacent claims in the same file were verified while I was there and also hold, which is what makes the
line trustworthy rather than lucky:

- `gotenberg.md:14` — `profiles: [core, backend, all]` at `docker-compose.yml:183`: **exact**, and the
  file is **186 lines**, exactly as the bullet says.
- `app/internal/converter/gotenberg.go:31` (cited by the CANON-1 evidence table) is
  `http.NewRequestWithContext(ctx, "POST", gotenbergURL+"/forms/libreoffice/convert", &body)` at
  `ad9f3c49` — **plain HTTP, not Connect-RPC**, which is the distinction the whole CANON-1 repair turns on.

**The finding is that this file is right where three others are wrong, and it is right for a legible
reason:** `gotenberg.md` states the *mechanism* (a compose-injected env var naming a second container) and
never generalises it into a claim about the *set* of cross-process edges. The three defective sites all
made the same move — from "the only RPC address" to "the only service address / the one cross-process
edge" — in files whose subject is a different service, where the gotenberg row was out of view. **A claim
about a set belongs in the file that owns the set** (`architecture_overview.md:321`, `dependency_map.md`,
`platform-migration-status.md`), not in a per-service page. This is the model, not a defect.

---

## Twins outside my files (REPORT, do not edit)

| predicate | file:line | why it is the same claim |
|---|---|---|
| CANON-3 currency pin | `corpus/architecture/service_taxonomy.md:106` | Cites `app/env_guards.go:37-44` — the identical anchor as my ledger row 2. Already pinned `@ ad9f3c49` by the owning seat; recorded only to confirm the two repairs agree and no third wording exists. |
| `messenger-unpinned-anchors` | `corpus/architecture/dependency_map.md:58` | Same `env_guards.go` mechanism, and it carries an anchor error of its own (below). Seat 7's file. |
| `storage-boot-guard-anchors` | — | **No twin exists.** `main.go:518-523` / `:529-535` are cited in `storage.md` only, corpus-wide. Recorded because the absence is the finding: this predicate is genuinely width-1, which is rare in this milestone. |

---

## Noticed, not repaired

1. **Three anchor-out-of-range REDs appear to have been INDUCED by concurrent seats in this very pass**, all
   from the CANON-1 fan-out. `anchor_construct_guard` is RED with 4 anchors, **none in my files**:
   - `corpus/architecture/dependency_map.md:7` cites `app/internal/converter/gotenberg.go:59` and `:66` —
     the file has **53 lines** at `ad9f3c49` (measured).
   - `corpus/services/sentinel.md:85` cites the same non-existent `gotenberg.go:59`.
   - `corpus/architecture/dependency_map.md:58` cites `env_guards.go:1437` — that file has **201** lines;
     `1437` is a `main.go` line number attributed to the wrong file.

   This is the brief's rule-4 rate (~2 induced per cycle) reproducing live, and it clusters on exactly the
   predicate the seats were told to word identically. **The correct anchor is `gotenberg.go:31`**, as
   `canonical-repairs.md` CANON-1 itself states — three seats appear to have carried a line number from a
   different measurement. Worth a targeted re-check before the iter closes.

2. **`repair_postcondition` is RED with 10 sites; exactly one is mine — `corpus/services/storage.md:64` —
   and it is a BASELINE artifact, not a regression.** The `claim_twin_guard` pattern is derived from
   `union-iter101.md` row 9, whose *"the false claim"* cell quotes the HAZARD **headline**
   (*"on a stock stack NEITHER manager uses local FS, and the private one writes to production"*) as the
   frame around the three anchors that were actually refuted. That headline is **TRUE** — re-derived here:
   `docker-compose.yml:82`/`:83` @ `0c91421` hardcode both buckets to production names on the `backend`
   block — so brief rule 5 forbids weakening it, and the text at `:64` is **byte-unchanged by me**. The
   baseline (`repair_postcondition_baseline.json`, `rosetta_head: 29eb414`) predates the iter-102 evidence
   files that supply the pattern, so it could not contain this key. **Clearing it requires either a
   `claim_twin_waivers.json` entry or a narrower quote in the union row — both outside `corpus/`, hence
   outside my edit scope.** Flagging rather than silently leaving it: a repair seat reporting green while a
   fence is red on its own file would be the exact failure this milestone measures.

3. `skillpath.md:107` says the page *"renders the literal string **"Coming soon"**"*. Strictly it renders
   `t('enterprise.insights.comingSoon')`; the `en` bundle resolves that to `"Coming soon"`
   (`configs/i18n/messages/en/enterprise.json:2101` @ `8297c684`), so the sentence is true for an
   English-locale reader and false-ish for the other six locales. Below the milestone's literal-falsity
   bar; left alone rather than risk an induced defect on a line I was not asked to touch.

4. `storage.md` still carries `9d00a313` and `2adcf71` as bare pins in several places, and
   `messenger.md:168` cites `@ b948604 v1.366.0`. **These are pins, not labels — CANON-3 explicitly
   protects them** ("a pin is a pin") and I left every one of them untouched.

---

## What I could not settle, and why

1. **`storage.md:55`, `:154`, `:181` (the `DEF-M257x-iter80-storage-prod-bucket` hold).** Deliberately not
   touched, per the brief. I made exactly the one sanctioned touch: the disposition sentence in the HAZARD
   block now names the register entry. **The id in my brief (`D-M257x-102-1`) is not the id in the file** —
   the entry is filed under
   `PLATFORM-M257x-compose-points-local-backend-at-the-PRODUCTION-S3-buckets` in
   `knowledge/plan/platform-defect-register.md` (currently **staged but uncommitted**, which is why a
   working-tree `grep` for it comes back empty and `git grep --cached` finds it — another §5 rule 44
   instance). I cited the id that exists, kept the `DEF-` id alongside, and made **no** claim that the
   `s3-private` isolation registry has changed. `isolation.go:106` was not read or asserted about.

2. **`platform_alignment_guard` could not be run** — it refuses without an explicit `repos.yml` path by
   design (*"there is deliberately no default — a fidelity check against the wrong reference passes"*), and
   supplying one is an orchestrator-level choice about which tree is the reference. Not run rather than run
   against a guessed reference.

3. **The rext-internal figure behind the old "six read sites"** (`platform_predicate_guard.py`'s own
   docstring says *three* app read sites at `b948604`, and its G6 comment says *one* variable) could not be
   reconciled to the corpus's "six", and `rosetta-extensions` is not my file. I resolved this by **not
   republishing an unverified count**: the repaired passage now carries only numbers I measured myself
   (15 Go hits / 7 env reads at `b948604f`; 3 comment hits at `ad9f3c49`).

---

## Guard state after my edits

| guard | verdict | mine? |
|---|---|---|
| `markdown_structure_guard` | **OK** — 112 files, no structural damage | — |
| `corpus_index_guard` | **OK** — 84 docs across 6 index-bearing dirs | — |
| `unreadable_repo_claim_guard` | **OK** — all 7 `module.*_euwest1` mentions marked unmeasurable | — |
| `anchor_construct_guard` | RED ×4 | **none in my files** (see Noticed #1) |
| `claim_twin_guard` / `repair_postcondition` | RED | **1 of 10 in my files** — `storage.md:64`, baseline artifact, unchanged text (see Noticed #2) |
