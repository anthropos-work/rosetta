# seat-2 report

**File owned:** `corpus/architecture/platform-migration-status.md` (the only file edited).
**Guard:** `platform_alignment_guard.py` **GREEN (exit 0)** after the edits — run reproduced below.
**Ground truth used:** platform `0c91421` (== clone HEAD == `origin/main`), `app` `origin/main` = `ad9f3c49`
(5 commits past `2035f9a4`), `cms` `origin/main` = `f38c0c4a`. Nothing was fetched.

---

## Ledger rows

| # | the false claim | anchor | what is true | sites repaired |
|---|---|---|---|---|
| 1 | "**Do not generalise M810 from this row:** `cms` has not moved, and is still `cms/terraform/main.tf:39` `service_desired_count = 0`." | `corpus/architecture/platform-migration-status.md:89` | `cms` **has** moved, in the other direction, and the row 1 line above already said so. Both halves re-measured at `cms` `origin/main` `f38c0c4a`: `terraform/main.tf:39` is `  service_desired_count          = 0` (module still declared), **and** `6efa1d5` (merged `f38c0c4`, 2026-08-04) deletes `.github/workflows/build-production.yml` with the body *"M810 deletes `module \"cms_euwest1\"` … which destroys the ECS service and the production-cms ECR repository"* and *"would try to push an image into a registry that no longer exists"* — confirmed by `git ls-tree f38c0c4a -- .github/workflows/`, which no longer lists it. Two measured facts pointing opposite ways: **report both, assert neither**. The deciding declaration is in `infrastructure`, which is in no clone set — so the state is *unmeasurable*, not *unmoved*. | 1 |
| 2 | "while `cms` sits untouched at `service_desired_count = 0` and `storage`'s service block is gone by a different route entirely" | `corpus/architecture/platform-migration-status.md:270` | Same predicate, same evidence as row 1 — the §5 prose restatement. *"Sits untouched"* is the same flat assertion the `cms` row retracts. The `storage` half of the sentence is **TRUE and was preserved**. | 1 |
| 3 | "(same command on that file → 26 names, including the pre-history the clone set never knew: `nats`, `web-app`, `chromedp`, `simulator`, `realtime`)" | `corpus/architecture/platform-migration-status.md:78-81` | Re-derived over all **80** commits `git log --follow` reports for `docker-compose.yml` in `stack-demo/platform` (renames followed, per-commit path resolved): collecting only keys under the `services:` section gives **25**; a section-blind pass over every 2-space key gives **26**, and `comm`-style set difference yields exactly one extra token — **`app-network`**, the network declared under `networks:` (`docker-compose.yml:185-186` @ `0c91421`). The `repos.yml` half of the same passage (**14** names, all with rows) was **re-verified and preserved** (`git log -p --follow -- repos.yml` → 14). | 1 |
| 4 | "the only cross-process service address left in a local stack is `AUTHORIZATION_ADDRESS` on `backend`, pointing at sentinel (`docker-compose.yml:48`)" | `corpus/architecture/platform-migration-status.md:93` (messenger row, **RPC-edge clause only**) | CANON-1 applied verbatim in substance. Measured @ `0c91421`: `:48` `AUTHORIZATION_ADDRESS=http://sentinel:8087`, **`:57` `GOTENBERG_URL=http://gotenberg:3200`**, `:59` `JUDGE0_BASE_URL=…`, `:66` `REDIS_ADDR=redis:6379`; `gotenberg` is declared at `:170-171` with `profiles: [core, backend, all]` at `:183` — the **default** profile — and reached over **plain HTTP**, not Connect-RPC (`app/internal/converter/gotenberg.go:31` @ `ad9f3c49` = `http.NewRequestWithContext(ctx, "POST", gotenbergURL+"/forms/libreoffice/convert", &body)`). The `*_RPC_ADDR`-is-zero half is **TRUE and was preserved** — re-measured: `git grep '_RPC_ADDR' 0c91421 -- docker-compose.yml common.yml .env_example` → **exit 1, zero hits**. | 1 |
| 5 | "**Anchors re-resolved M257x iter-87 at `app` origin/main `2035f9a` (post-v1.369.0).**" | `corpus/architecture/platform-migration-status.md:87` | CANON-3. The **sha is a pin and still correct**; the **label expired**. `app` `origin/main` is `ad9f3c49`, **5 commits** on, touching `.claude/skills/publish/SKILL.md`, `CLAUDE.md`, `knowledge/deployment.md`, `terraform/main.tf`, `terraform/variables.tf` — **no Go source at all**; the one `terraform/main.tf` hunk rewrites a precondition `error_message` in place (`@@ -162,7 +162,7 @@`), and every anchor this cell carries still holds at `ad9f3c49` (checked: `terraform/main.tf:181` = `service_desired_count = 1`). Label dropped, sha kept, currency stated once here. | 1 |
| 6 | "(`jobsimwiring.Wire`, @ `app` origin/main `2035f9a`)" | `corpus/architecture/platform-migration-status.md:89` | CANON-3, same measurement. Moving label dropped; the sha kept as a pin, pointing at the banner for currency. | 1 |
| 7 | "(`app/internal/jobsimwiring/wiring.go:123` @ `app` origin/main `2035f9a` — the same line at `9d00a313`;" | `corpus/architecture/platform-migration-status.md:90` | CANON-3, same measurement. Label dropped; sha kept. | 1 |
| 8 | "a Go grep at `app` origin/main `2035f9a` returns **3 hits, every one of them a comment**" | `corpus/architecture/platform-migration-status.md:92` | CANON-3, same measurement. Label dropped; sha kept. **`ad9f3c49` was deliberately NOT named inside this cell** — the cell itself documents that naming a second ref makes every anchor in it ungradeable to the citation resolver (M257x run-53); the currency statement lives once, in the `app` row. | 1 |
| 9 | "(local clone and origin level; `app` @ `2035f9a`, post-v1.369.0)." | `corpus/architecture/platform-migration-status.md:16` | **PINS, does not label** — so per CANON-3 it is *not* a defect and the sha was not touched. The only exposure is that the preceding phrase *"local clone and origin level"* could be read as reaching across the semicolon onto `app`. Repaired to the extent of making the pin explicit (`app` **pinned** at `2035f9a` … a sha, not a moving label) — no sha changed, no line added. | 1 (clarification, not a booked defect) |
| — | *self-induced by row 4 and closed in the same pass:* "`docker-compose.yml` citations — **22** of them" / "(**69** citations, + **1** outside…)" / "# assertion F resolved 92 citation(s) — 22 subject-checked, 69 range-only" | `:204`, `:205`, `:220` | Row 4's repair adds 3 compose citations and 1 `app/…` citation, so the guard now reports **96 resolved / 25 subject-checked / 70 range-only / 1 outside**. These three sites were **true before my edit and false after it** — the exact drift §4 warns about (*"M257x iter-98 found the row above saying 23/66 while the guard said 22/69"*). Updated to match the run, with the move recorded in the prose. | 3 |

---

## Reach — graded, not asserted

| predicate | anchors BOOKED | sites FOUND | sites REPAIRED | how I searched |
|---|---|---|---|---|
| `cms-has-not-moved` | 2 in my file (`:89`, `:270`) + 1 cross-seat (`jobsimulation.md:12`) | 2 in my file; 1 outside (reported, not edited) | **2 / 2** in my file | `grep -n -iE 'has not moved\|sits untouched\|untouched at\|unmoved\|still stands at'` over the file, then a corpus-wide `git grep -n -iE 'cms.{0,40}has not moved\|cms.{0,40}sits untouched\|has not moved.{0,20}cms'`. Two mechanisms: the paraphrase grep and a `service_desired_count = 0` co-occurrence grep (which also surfaces the `messenger` and `storage` rows, correctly, as *different* claims) |
| `compose-service-census-26` | 1 (`:78-81`) | **1** — no twin, no paraphrase anywhere in the corpus | **1 / 1** | Re-derived the number from the platform clone (80 commits, rename-followed, section-aware vs section-blind), then swept corpus-wide with `git grep -n -iE '\b26\b.{0,40}(service\|compose\|name)\|(service\|compose\|name).{0,40}\b26\b'` minus date/line-number noise → **0 other sites**. Second mechanism: `git grep '26 names\|26 service names\|ever appeared in .docker-compose'` → only this file |
| `sentinel-only-cross-process-edge` (RPC-edge clause) | 1 in my file (`:93`) | 1 false + 1 already-correct (`:105`, verified) | **1 / 1** | `grep -n -iE 'only cross-process\|one cross-process\|single service address\|only service address\|exactly one service'` over the file → exactly one false site. Cross-checked the underlying fact directly in compose at `0c91421` rather than trusting the booking |
| `currency-pin-2035f9a` (CANON-3) | 4 (`:87`, `:89`, `:90`, `:92`) | **5** — the 4 labelled sites plus `:16`, which pins rather than labels | **4 / 4** labelled sites repaired; `:16` clarified without changing its sha | `git grep -n '2035f9a'` (5 hits) and `git grep -n 'origin/main'` (2 hits) over the file, reconciled against each other so a labelled site could not hide in either list alone. Post-edit re-grep: **0** remaining `origin/main` labels on `2035f9a` |
| *(induced, closed)* fence-coverage counts | 0 (not booked) | 3 (`:204`, `:205`, `:220`) | **3 / 3** | Ran the guard before and after; the doc states its own coverage numbers in three places and instructs the reader to take them from the run |

**Line-number stability was treated as part of the repair.** Two other corpus files cite rows of my file by
line — `corpus/architecture/service_taxonomy.md:129` → `platform-migration-status.md:89` and
`corpus/services/README.md:17` → `:101`. My first draft of the census repair grew §2's preamble by 5 lines
and would have slid **both** of those onto the wrong rows — the exact iter-100 induction class. I re-wrapped
the paragraph to keep it at **exactly 7 physical lines**, so every table row sits on the line it sat on
before: `app` 87 · `cms` 88 · `jobsimulation` 89 · `roadrunner` 90 · `sentinel` 91 · `storage` 92 ·
`messenger` 93 · `customerio-sync` 101 · `gotenberg` 105. Verified post-edit. The file's net +5 lines are all
in §5, below the table and below both cited anchors.

---

## Twins outside my files (REPORT, do not edit)

| predicate | file:line | why it is the same claim |
|---|---|---|
| `cms-has-not-moved` | `corpus/services/jobsimulation.md:12` (**seat 8**) | Verbatim the same predicate in the same words: *"to `cms`**, which has not moved (`cms/terraform/main.tf:39` `service_desired_count = 0`)."* Same citation, same false generalisation from jobsimulation's landed M810 onto cms. **My repair and seat 8's must be checked for consistency** — the canonical shape I used is *"two measured facts pointing opposite ways; report both, assert neither; the destruction lands in `infrastructure`, which is in no clone set, so the state is unmeasurable, not unmoved."* |
| `sentinel-only-cross-process-edge` | `corpus/services/sentinel.md:85` (seat 10) · `corpus/services/jobsimulation.md:145-146` (seat 8) · `corpus/architecture/service_taxonomy.md:405` (seat 3) · `CLAUDE.md:280` (orchestrator) | Per `canonical-repairs.md` §CANON-1 — one predicate, five writing seats. I applied the canonical form at `:93` only. |
| `sentinel-only-cross-process-edge` (**already correct — verify only**) | `corpus/architecture/architecture_overview.md:321` (model, do not edit) · `corpus/services/gotenberg.md:50` (seat 9) · `corpus/architecture/dependency_map.md:103` (seat 7) | The model wording and its two existing agreements. |
| `currency-pin-2035f9a` | `CLAUDE.md:223`, `:280` (orchestrator) · `dependency_map.md:59` (7) · `backend.md:70/138/254/299` (1) · `academy-backend.md:20` (5) · `ai-labs.md:18` (6) · `coursebuilder.md:132` (6) · `messenger.md:43` (9) · `skillpath.md:35` (9) · `storage.md:29` (9) | Per CANON-3's anchor table. Not touched. |

**CANON-1 verification of `:105` — requested, and it holds.** The `gotenberg` row reads
`| `gotenberg` | external | live-standalone | no | third-party image, `docker-compose.yml:170-171`
(`gotenberg/gotenberg:8`), default `core` profile (`:183` …) |`. Re-measured at `0c91421`: `:170-171` is the
`gotenberg:` service key and `image: gotenberg/gotenberg:8`; `:183` is `profiles: [core, backend, all]`.
**Correct as it stands; left untouched.** The finding stands as the adjudicator framed it: this row is
**12 rows below** the row that denied it (93 → 105), in the same table, and graded the container
`live-standalone` on a fresh local stack the whole time. My repair at `:93` now points at it explicitly, so
the two rows can no longer be read apart.

---

## Noticed, not repaired

1. **`corpus/architecture/service_taxonomy.md:129` and `corpus/services/README.md:17` cite my file by LINE.**
   Both currently resolve (`:89` = the `jobsimulation` row, `:101` = the `customerio-sync` row) and I
   preserved them. **Any seat that adds or removes a line above line 105 of `platform-migration-status.md`
   breaks both**, silently — they resolve to a neighbouring row rather than erroring. Worth a guard.
2. **The `app` row's *"Owns **seven** domains in-process"*** and its seven wiring call sites were not in my
   assignment and were not re-derived. They resolve (the guard checks range), but nothing checked the
   *construct* at each.
3. **`:78-81`'s audit instruction has a residual false-alarm surface I closed only partially.** Three of the
   25 compose service names (`backend`, `graphql`, `wundergraph`) have no row *under that name* — their rows
   are filed under the repo name (`app`, `graphql-wundergraph`). I added a one-clause caveat rather than a
   full mapping table, because a mapping table is a new claim surface and this was not the booked defect.
4. **`storage`'s row (`:92`) is the longest cell in the file and names five refs.** It documents its own
   ambiguity hazard and then keeps doing the thing it warns about. Out of scope here; noted.

## What I could not settle, and why

- **Whether the cms ECS service actually exists in production.** Both measured facts are in the `cms` repo
  and they point opposite ways; the deciding declaration is `module "cms_euwest1"` in **`infrastructure`**,
  which **has never been in any clone set**. TRAP A applies: I restated it as *not measurable from this repo*
  and did **not** re-anchor the citation at anything that would make it resolve. `unreadable_repo_claim_guard`
  agrees — it reports *"all 7 `module.*_euwest1` mention(s) are marked unmeasurable"* (GREEN, re-run below).
- **Cross-seat consistency of the `cms-has-not-moved` wording** with seat 8's repair at
  `jobsimulation.md:12`. That predicate is not in `canonical-repairs.md`, so no central wording exists; the
  two repairs need reconciling by the orchestrator. My form is quoted verbatim in the twins table above.

---

## Verification runs (post-edit)

```
$ python3 .agentspace/rosetta-extensions/stack-core/platform_alignment_guard.py \
      corpus/architecture/platform-migration-status.md stack-demo/platform/repos.yml
platform_alignment_guard: assertion F resolved 96 citation(s) — 25 subject-checked, 70 range-only,
  1 outside any service block; 0 unresolvable; 0 read from the WORKTREE (no ref resolved)
platform_alignment_guard: OK — platform-migration-status.md and repos.yml agree in both directions.
EXIT=0
```

Baseline before my edits was `92 resolved / 22 subject-checked / 69 range-only / 1 outside`, also exit 0 —
so the fence was GREEN before and after, and the delta is entirely the four citations row 4 added.

Two further guards, run unprompted because my edits touch their subject matter, both GREEN:

```
$ python3 .agentspace/rosetta-extensions/stack-core/platform_predicate_guard.py --platform stack-demo/platform
platform_predicate_guard: … G4 2 local RPC-address claim(s); … G6 8 RPC var(s) graded {'unconfigured': 8},
  0 mid-fold; app consumer side measured @ origin/main@ad9f3c4
platform_predicate_guard: OK — the corpus and the platform's configuration agree.            EXIT=0

$ python3 .agentspace/rosetta-extensions/stack-core/unreadable_repo_claim_guard.py
unreadable-repo-claim-guard: OK — all 7 `module.*_euwest1` mention(s) are marked unmeasurable
  (infrastructure is in no clone set).                                                        EXIT=0

$ python3 .agentspace/rosetta-extensions/stack-core/markdown_structure_guard.py
markdown-structure-guard: scanned 112 published file(s) — OK, no structural damage             EXIT=0
```

Nothing was committed. Nothing outside `corpus/architecture/platform-migration-status.md` was written except
this report.
