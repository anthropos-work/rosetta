# seat-8 report

**Files owned:** `corpus/services/jobsimulation.md` · `corpus/services/cms.md` ·
`corpus/services/graphql-wundergraph.md`. **Anchors booked: 8. Sites found: 8 defect sites (+3 graded
clean). Sites repaired: 8. Unrepaired defect sites in my files: 0.**

**Ground truth re-measured at this iter's open, not inherited:** platform `0c91421` · `app` `ad9f3c49`
(== `origin/main`) · `cms` `ca50c817` (checkout) / `f38c0c4a` (`origin/main`) · `jobsimulation`
`462343b0` (checkout) / `82cb66ec` (`origin/main`). No `git fetch` was run. No commit was made.

**Every edit is strictly line-count-neutral** (12 insertions, 12 deletions, delta 0 on all three files).
That was deliberate: all three files carry intra-file line self-citations that a shifted line would
silently break — `jobsimulation.md:95`→`:39-45`, `cms.md:211`→`:44-47`, `cms.md:235`→`:70-71`,
`graphql-wundergraph.md:86`→`:193`. All four were re-opened after editing and still resolve to the
constructs they name. This is the induced-defect class iter-100 shipped (a parenthetical that pushed a
table down two rows); line-neutrality removes it by construction rather than by care.

---

## Ledger rows

| # | the false claim | anchor | what is true | sites repaired |
|---|---|---|---|---|
| 1 | "`app/main.go:216` (@ `origin/main` `7177374`, identical at `9d00a313` v1.367.0; `:212` at the older `b948604` v1.366.0) is a plain `func main()`" | `corpus/services/jobsimulation.md:203-204` | The **pin is correct and was kept**: `git show 7177374:main.go \| sed -n 216p` → `func main() {`, and identically at `9d00a313`; `:212` at `b948604` → `func main() {`. Only the **`origin/main` LABEL expired** — `7177374` is **38 commits** behind `app` `origin/main` `ad9f3c49` (it was origin/main on 2026-08-04), where `func main()` is at **`:229`** and `:216` is `if strings.Contains(dsn, "://") {`. The surrounding claim survives re-derivation at the new ref: `git grep -l spf13/cobra ad9f3c49 -- '*.go'` → **exactly one file, `cmd/createTaxonomy/main.go`**. CANON-3 move 1 applied (keep the sha, drop the moving label) + the current reading added. | 1 |
| 2 | "**Do not generalise this to `cms`**, which has not moved (`cms/terraform/main.tf:39` `service_desired_count = 0`)." | `corpus/services/jobsimulation.md:12` | **Half true, and the false half is the headline.** The module-block half verifies at `cms` `origin/main` `f38c0c4a`: `terraform/main.tf:39` is still `service_desired_count = 0`, in a 191-line module byte-identical to `ca50c817` — **kept** (rule 5). But `cms` **has** moved: `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** `.github/workflows/build-production.yml` under the subject *"the cms ECR repository is decommissioned (M810)"*. Measured directly: `git ls-tree ca50c817 -- .github/workflows/` lists the file, `git ls-tree f38c0c4a` does not. Two measured facts point opposite ways and the deciding declaration is in `infrastructure`, in no clone set → restated as **report both, assert neither**, matching the corpus's standing fenced position. | 1 |
| 3 | "`sentinel`, which is the **only** cross-process hop a local stack has left and the only service address `backend`'s compose entry carries (`docker-compose.yml:48`)." | `corpus/services/jobsimulation.md:145-146` | CANON-1 applied verbatim in substance. `sentinel` is the only cross-process **Connect-RPC** edge; it is **not** the only cross-process edge and compose does **not** set exactly one service address. Re-measured at platform `0c91421` in `backend`'s own `environment:` block: `:48 AUTHORIZATION_ADDRESS=http://sentinel:8087`, **`:57 GOTENBERG_URL=http://gotenberg:3200`**, `:59 JUDGE0_BASE_URL=…`, `:66 REDIS_ADDR=redis:6379`. `gotenberg` is in the **default `core` profile** (`docker-compose.yml:183` `profiles: [core, backend, all]`) and is reached over **plain HTTP**, not Connect-RPC (`app/internal/converter/gotenberg.go:31` @ `ad9f3c49` — `http.NewRequestWithContext(ctx, "POST", gotenbergURL+"/forms/libreoffice/convert", &body)`). The **`*_RPC_ADDR`-is-zero clause is TRUE** (`git grep RPC_ADDR 0c91421 -- docker-compose.yml common.yml` → 0) and was preserved, not weakened. | 1 |
| 4 | "Production terraform still names `http://backend.internal.anthropos:8081`." | `corpus/services/jobsimulation.md:49-50` | CANON-2 (corrected form) applied — **the assertion is dropped, not softened**. Measured 2026-08-06: `git grep 'backend.internal.anthropos' <ref> -- '*.tf'` at each clone's own HEAD over all **44 tracked `.tf` files in the 13 `stack-demo` repos** → **0**; second mechanism, `find stack-demo -name '*.tf' \| xargs grep -l` over the **59** `.tf` files on disk → **0**. The literal's only occurrence anywhere in the clone set is a **markdown KB page** — `app/knowledge/service-dependencies.md:52` @ `ad9f3c49` — which is in the **past** tense under the heading *"**There are no external callers of app's RPC mux left.**"* And the deciding declaration is in `infrastructure`, in no clone set, so no *"still names"* claim can be made in **either** direction. | 1 |
| 5 | "Production terraform still names `http://backend.internal.anthropos:8081`." | `corpus/services/cms.md:196` | Same predicate, verbatim the same sentence, same evidence as row 4. Full form applied here (this is the booked blocker) plus the file's own **self-contradiction** made explicit: `cms.md:18` already says *"the deletion itself lands in `infrastructure`, **which has never been in any clone set we have**"* — a doc that states it cannot see the production terraform cannot report what it *"still names."* | 1 |
| 6 | "`http://backend.internal.anthropos:8081` in production." | `corpus/services/cms.md:55` | Same predicate in a compressed phrasing — a **paraphrase**, not a quoted twin, which is the axis iter-97 measured as the entire escape set. Short form of the CANON-2 replacement applied, pointing at the full form in the same file. Deliberately references the **RPC line under *Interface Discovery*** by construct name rather than by line number, so it cannot rot. | 1 |
| 7 | "re-derived at `app` origin/main `2035f9a`" | `corpus/services/cms.md:216-217` | **A pin AND a moving label in one phrase.** The pin verifies and was kept: at `2035f9a4`, `internal/cms/studio/markdownManager.go:10` is the `mistralocr` import and `:30` is `return &MarkdownManager{ocr: mistralocr.New(aiKey)}, nil` — **both byte-identical at `ad9f3c49`**, so this is CANON-3 move 1 (keep the sha, drop the moving label), with the still-current status recorded rather than the sha silently swapped. | 1 |
| 8 | "which is why a subgraph SDL change rebuilt the router, as `:84` describes." | `corpus/services/graphql-wundergraph.md:134` | `:84` is the **Ports** bullet (`* **Ports**: **8080 → 8080** (router `listen_addr 0.0.0.0:8080`…`), entirely about ports and the absent `5050`. The construct the sentence needs is the ***Build-time, static composition*** pair at **`:114-117`**, and specifically **`:116-117`** — the struck-through *"`make up` rebuilds `graphql`"* bullet whose text is *"It **used to** rebuild whenever any subgraph schema changed, because the build context is the parent dir (`..`) holding all sibling repos."* Re-pointed, and **`:84` is now named as the wrong anchor** so the error cannot re-enter silently. | 1 |

**Post-repair anchor verification (the thing iter-100's fix was built for, and iter-101 still found 4 of):**
I re-opened every line number I introduced. `graphql-wundergraph.md:114-117` and `:116-117` resolve to
the two *Build-time, static composition* bullets; `architecture_overview.md:321` (newly cited from
`jobsimulation.md:146` as the model wording) resolves to
`→ Connect-RPC to sentinel   (the only cross-process RPC edge out of backend on a core stack)`;
`docker-compose.yml:48/57/59/183` and `app/internal/converter/gotenberg.go:31` all resolve at their
named refs. `anchor_construct_guard` independently confirms: **RED with 4 sites, none of them mine.**

---

## Reach — graded, not asserted

| predicate | anchors BOOKED | sites FOUND (my files) | sites REPAIRED | how I searched |
|---|---|---|---|---|
| `prod-terraform-8081` (CANON-2) | 3 | **3** | **3** | `grep -n 'backend\.internal\.anthropos\|8081'` over all 3 files, then the paraphrase axis `grep -niE 'production terraform\|prod terraform\|terraform still\|8081. in production'`. The three other `8081` hits in my files (`jobsimulation.md:96`, `:290`, `cms.md:113`) are the **binary's own default RPC port** — a different, TRUE claim, left alone. Corpus-wide cross-check by `git grep` at HEAD + a second mechanism (`git ls-files \| xargs grep`), both returning 10. |
| `sentinel-only-cross-process-edge` (CANON-1) | 1 | **1** | **1** | `grep -niE 'only.*cross-process\|cross-process.*only\|only service address\|exactly one service address\|single service address\|the only .{0,20}hop'`. One site only in my files; `cms.md` states the same edge set without the false universal. |
| `cms-has-not-moved` | 1 | **3** (1 defect + **2 graded clean**) | **1** | `grep -niE 'has not moved\|sits untouched\|not moved\|unmoved\|desired_count'`. `cms.md:9` and `cms.md:72` also say *"has not moved"* — but both are **precisely scoped to the terraform module block** (true, re-measured at `f38c0c4a`) and both are immediately followed by the ⚠️ correction at `:11-19` / `:78-84`. Repairing them would have **weakened a true clause** (rule 5). Only `jobsimulation.md:12` published the flat unscoped form. |
| CANON-3 currency pin (`origin/main` as a label) | 2 | **2** (+1 graded clean) | **2** | `grep -n 'origin/main'` + `grep -n '2035f9a\|7177374\|9d00a313\|b948604'`. The third hit, `jobsimulation.md:25`, uses `origin/main` as a **branch**, not as a label for a sha, and names its four commits explicitly — re-verified: `6092c6d2`, `caf36c96`, `1e40d184`, `82cb66ec` all landed 2026-08-04 and `82cb66ec` is still the tip. Not a defect. |
| `gwg-self-cite-84` | 1 | **1** | **1** | Enumerated **every** intra-doc `:N` self-citation in all three files (`grep -nE '(this (same )?(doc\|file)\|above\|below\|bullet)[^.]{0,40}:[0-9]+'`) and opened each target. `graphql-wundergraph.md:86`→`:193`, `jobsimulation.md:95`→`:39-45`, `cms.md:211`→`:44-47`, `cms.md:235`→`:70-71` all resolve correctly; `:134`→`:84` was the only broken one. |

**Booked-vs-live width, for the milestone's own estimator:** 8 booked → **8** live defect sites in my
files. **The 3× under-count this milestone has been measuring did not reproduce in my partition** — and
I think that is a real signal rather than a shallow search. Three of my four predicates are
*cross-seat*, so their width had already been enumerated centrally in `canonical-repairs.md` before I
started; the booking was not under-counting, it was pre-expanded. The one predicate I expanded myself
(`cms-has-not-moved`) found 2 extra sites — and **both graded clean**, which is the outcome the
expansion rule exists to allow for.

**Search discipline (§5 rule 44).** No absence in this report rests on `grep -r`. Every negative was
established by `git grep` at a named ref, and every one that mattered was cross-checked with a second
mechanism: the `.tf` census by `git ls-tree`+`git grep` per clone **and** by `find … -exec grep`
(44 tracked / 59 on disk — the two numbers are different quantities, both 0 hits, and I say which is
which rather than picking the flattering one).

---

## Twins outside my files (REPORT, do not edit)

| predicate | file:line | why it is the same claim |
|---|---|---|
| `cms-has-not-moved` | `corpus/architecture/platform-migration-status.md:89` · `:270` | **SEAT 2.** These are the other two anchors of the predicate I repaired at `jobsimulation.md:12`. **Cross-seat consistency check requested:** my restatement says the module block has *not* moved (`f38c0c4a`, `terraform/main.tf:39`, 191-line module unchanged) **while** `6efa1d5` deleted the build-production workflow — *report both, assert neither*. Seat 2's two repairs should land on the same reading, and `:270`'s phrasing (*"while `cms` **sits untouched** at `service_desired_count = 0`"*) is a paraphrase, not a quoted twin, so `claim_twin_guard` will not catch a divergence there. |
| `prod-terraform-8081` | `corpus/services/backend.md:140`, `:153`, `:157`, `:282` | **SEAT 1**, already repaired in the working tree when I read it. I aligned my numbers to seat 1's independently — we both derived **44 tracked `.tf` / 13 repos** and **`service-dependencies.md:52` @ `ad9f3c49`**, from separate measurements. Consistent. |
| `prod-terraform-8081` | `corpus/architecture/dependency_map.md:27` | **SEAT 7.** Same predicate, different variable: *"(production terraform: `skiller_rpc_addr = http://backend.internal.anthropos:8081` …)"*. Present tense, no ref, and the address is not in any `.tf` in any clone. **Not in either union** — flagging it so the predicate-scoped repair does not leave it standing. |
| `prod-terraform-8081` | `corpus/services/skiller.md:19` | **SEAT 9 (?).** *"`skiller_rpc_addr = http://backend.internal.anthropos:8081` in production"* — the `cms.md:55` compressed phrasing, applied to skiller. iter-101 rejected a *different* booking at this same line; this clause is the CANON-2 predicate and is unaffected by that rejection. |
| `sentinel-only-cross-process-edge` | `corpus/services/sentinel.md:85` · `corpus/architecture/dependency_map.md:7`, `:60` | **SEATS 10 / 7.** `anchor_construct_guard` is RED on three of these for the **bare-`:N`** defect that `canonical-repairs.md`'s own ⚠️ correction documents: `app/internal/converter/gotenberg.go:59` and `:66` are cited where the file is **53 lines** long. My application of CANON-1 prefixes every anchor with `docker-compose.yml:`, so `jobsimulation.md:146` is clean — but the seats that took the canon's earlier bare-`:N` form are not. See `## Induced and corrected`. |

---

## Induced and corrected

**The first version of `canonical-repairs.md` §CANON-2 was under-strength, and I did not ship it.**
Its original wording said only *"not measurable from this repo"* — which reads as *"we didn't check"*
when in fact the adjudicators measured it and found **zero**. The canon was corrected at source
mid-seat by its author and I applied the corrected form, which carries **both** halves: the measured
zero **and** the unmeasurability of the one tree that could settle it. My three CANON-2 sites carry
both. Recorded here because it is this iter's own datum on how centralised wordings propagate: **the
correction reached me before I wrote, so the cost was zero — but only because the author pushed it.**

**Two numeric defects inside the corrected canon, which I did NOT propagate.** Applying a canonical
wording is not a licence to stop measuring, and both of these would have been wrong-construct
citations of exactly the class this milestone exists to kill:

1. **`app/knowledge/service-dependencies.md:46` is wrong at the ground-truth ref.** `:46` is correct
   only at the **old demo pin `b948604f`**. At `2035f9a4` and at today's `ad9f3c49` the literal is at
   **`:52`**, and `:46` names an unrelated construct (the `storage` row of a table). Publishing an
   unpinned `:46` would have manufactured a wrong-construct citation. **I published `:52` @
   `ad9f3c49`** — independently matching what seat 1 derived for `backend.md`.
2. **"12 clones / 59 `.tf` files" mixes two accounting bases.** Both are defensible and they are not
   the same quantity: **44** git-*tracked* `.tf` across the **13** `stack-demo` clone dirs, versus
   **59** `.tf` files a raw *filesystem* sweep finds there (the extra are untracked working-tree
   artifacts). "12 clones" counts the platform clones and excludes `rosetta-extensions`. I published
   **both numbers, each labelled with its basis**, and both are **0 hits**.

**Zero defects induced by my own edits, and the check is mechanical, not asserted:** delta 0 lines on
all three files; all four intra-file self-citations re-opened and still resolving; every new anchor
re-opened; `markdown_structure_guard` OK; `claim_twin_guard`, `anchor_construct_guard` and
`corpus_index_guard` all report **0 hits in my files**.

---

## Noticed, not repaired

1. **`repair_leak_guard` flags `jobsimulation.md:46` and `cms.md:48` — both are FALSE POSITIVES, and I
   verified rather than "fixed" them.** The guard matches the phrase *"on `app`'s single RPC mux, and
   nothing outside the process reaches it"* because seat 1 rewrote a near-identical string in
   `backend.md`. But the refuted claim there was a **universal over all eight folded services**
   (`eight-folded-rpc-mux`, union-iter101 row 5); my two sites make the **per-service** claim, and both
   are TRUE: `app` `main.go:1314` @ `ad9f3c49` is
   `mux.Handle(jobsimulationv1connect.NewJobSimulationServiceHandler(jobsimDj.RPCServer))` and `:1323`
   is `mux.Handle(cmsv1connect.NewCMSServiceHandler(cmsRPCServer))`. Rewriting them would have
   **weakened a true clause to satisfy a string matcher** — rule 5. Recorded so the next reader does
   not "fix" it either.
2. **`cms.md:9` calls cms *"the one M810 row whose terraform module block has not moved"*.** Outside my
   assignment and in neither union, but the word *one* is load-bearing and I could not fully settle it:
   `messenger`'s module block also still declares the service, at `service_desired_count = 0`
   (`messenger` `origin/main` `e9421c68`, `terraform/main.tf:29`) — though unlike cms's it **did**
   change (it reads `= 1` at the demo pin `fa47850d:terraform/main.tf:19`), so the exclusivity may hold
   on a "unchanged" reading and fail on a "still declared" one. **Routing to seat 2**, which owns the
   fenced map where the M810 row set is defined.
3. **`jobsimulation.md:9`'s `service_desired_count`-is-absent claim is TRUE only at `origin/main`.** At
   the demo checkout `462343b0` the `module "jobsimulation"` block is intact with
   `terraform/main.tf:40 service_desired_count = 0`; at `origin/main` `82cb66ec` the file is 56 lines
   and the block is replaced by the M810 decommission comment. The claim's **subject is production
   infrastructure**, so `origin/main` is the settling tree and it holds — but the divergence is large
   (343 → 56 lines) and worth knowing before anyone re-reads that block against the checkout.
4. **`jobsimulation.md:225-229`'s `docker-compose.yml:91` AWS-bind anchor** is pinned to `0dab54d` and
   is correct there; at `0c91421` the bind is at **`:100`**. Correctly pinned, so not a defect — noted
   only because the offset moved.

---

## What I could not settle, and why

1. **Whether `infrastructure/terraform/production/services.tf` still declares `module.cms_euwest1`** —
   and by design. The `infrastructure` repo has never been in any clone set, and this is **TRAP A** in
   its textbook shape: the fact was *deleted*, not moved. I **restated** at all three CANON-2 sites and
   **did not re-anchor** — specifically, I did **not** repoint any citation at
   `app/knowledge/service-dependencies.md:52`, which is a markdown KB page and not the production
   terraform the sentence claimed to be reporting. Citing it would have produced a *correctly-cited
   false statement*, which is worse than a stale one. `unreadable_repo_claim_guard` is **OK** (all 7
   `module.*_euwest1` mentions marked unmeasurable).
2. **cms's M810 production state, in either direction.** Two measured facts in the `cms` repo point
   opposite ways and the deciding file is unreadable from here. I held the corpus's standing position —
   *report both, assert neither* — and was careful that the CANON-2 repair at `cms.md:55`/`:196` did
   **not** drift into asserting cms's prod state, which was the specific caution in my brief.
3. **The `graphql-wundergraph` repo's GitHub archive state** (`graphql-wundergraph.md:8` claims
   ARCHIVED 2026-07-30). Not measurable from a clone — the same class as `jobsimulation.md:24`'s
   already-retracted archive assertion. Not booked, not in my assignment, not repaired; flagged as a
   live instance of a class this milestone has already retracted once elsewhere.

---

## Guard runs (verbatim results)

```
claim_twin_guard          RED — 8 sites   | hits in MY files: 0   (was 14 before my edits;
                                            cms.md:196 ×2 and jobsimulation.md:49 ×2 cleared)
anchor_construct_guard    RED — 4 sites   | hits in MY files: 0
markdown_structure_guard  OK — no structural damage
corpus_index_guard        (no findings)   | hits in MY files: 0
unreadable_repo_claim_guard OK — all 7 `module.*_euwest1` mentions marked unmeasurable
repair_leak_guard         RED            | 2 hits in my files, BOTH verified false positives (see
                                           Noticed #1) — the per-service mux claim is true at ad9f3c49
```

The 8 remaining `claim_twin_guard` sites and the 4 remaining `anchor_construct_guard` sites all belong
to other seats' files: `dependency_map.md`, `external_services.md`, `platform-migration-status.md`,
`shared_libraries.md`, `services/README.md`, `academy-backend.md`, `storage.md`, `sentinel.md`.

**Not committed.** Working tree only, per the seat brief.
