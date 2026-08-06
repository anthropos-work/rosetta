# adjudicator 2 — iter-103 verdicts (seats r25-C, r26-C, r25-D, r26-D)

**Trees I read.** Every clone ref re-verified at this adjudication's open with `git rev-parse`, **no fetch,
no pull, no state change**: `platform` `0c91421d` · `app` `ad9f3c49` · `next-web-app` `8297c684` ·
`sentinel` `f2c46190` · `studio-desk` `41ee3575` · `ant-academy` `22df69dd` · `cms` `ca50c817` ·
`jobsimulation` `462343b0` · `messenger` `fa47850d` · `storage` `4ce8ece5` · `roadrunner` `87d8d443` ·
`graphql-wundergraph` `60c229f3` · **`stack-demo/rosetta-extensions` `09d06070` — the PINNED per-stack
consumption clone, the tree I used for every tooling/rext claim** (the authoring copy `.agentspace/…`
was not needed to settle anything below). Corpus at rosetta `bd9d40d1`, `git status --porcelain corpus/`
empty. `find stack-demo -name .git -maxdepth 4` → **15** trees (13 top-level + the two nested `studio`
checkouts at `aeec036a`); every absence/uniqueness claim below was measured across all 15 at each tree's
own HEAD **and** with a raw `grep -ar` filesystem sweep, and the two instruments agreed.

---

r25-C B1 | corpus/services/jobsimulation.md:50 | UPHELD | IN-SCOPE | "one occurrence anywhere in the clone set" — the literal occurs 6 times in 2 repos
   evidence: per-tree `git grep -F 'http://backend.internal.anthropos:8081'` at each of the 15 trees' own HEADs returns `stack-demo/app/knowledge/service-dependencies.md:52` (1) **and 5 more in `stack-demo/rosetta-extensions` @ `09d06070`** — `stack-core/tests/fixtures/repair_leak/{pre,post}/corpus/services/cms.md:32` and `:157`, `stack-core/tests/test_platform_predicate_guard.py:435`. A raw `grep -ar` over `stack-demo` returns the same 4 files. The sentence's own denominator forbids excluding rext: I counted tracked `.tf` per tree — app 5 · jobsimulation 6 · storage 5 · sentinel 4 · messenger 4 · cms 4 · studio-desk 4 · graphql-wundergraph 4 · roadrunner 4 · next-web-app 2 · **rosetta-extensions 2** = **44**, and `ls -d stack-demo/*/` = **13**. Without rext the sum is 42; 44 is reachable only with it counted. Raw-filesystem `.tf` = 59, and 0 `.tf` hits — the first half of the sentence is correct and reproduces exactly; the uniqueness clause is not.

r25-C B2 | corpus/services/jobsimulation.md:146 | UPHELD | IN-SCOPE | citation for the "correctly-scoped model wording" lands on the production Cosmo-Router line
   evidence: `corpus/architecture/architecture_overview.md:321` is `User → Vercel (Next.js) → Clerk (JWT) → ALB → Cosmo Router (port 8080)` — the first line of the fence headed at `:318` *"**In production** (the router still exists there …)"*. The wording the citing sentence needs is `:331`, `→ Connect-RPC to sentinel   (the only cross-process RPC edge out of backend on a core stack)`, inside the *"**On a local stack** (platform `2adcf71` deleted the router …)"* fence at `:327-336`. Decisive corroboration from the corpus itself: `corpus/services/sentinel.md:85` makes the same `:321` citation **and quotes the text it expects** — *"the only cross-process RPC edge out of backend on a core stack"* — which is verbatim `:331`, not `:321`. (`corpus/services/backend.md:54` carries the same wrong anchor; outside this seat's file set, recorded for the repair.)

r26-C B1 | corpus/services/jobsimulation.md:50 | UPHELD | IN-SCOPE | same predicate as r25-C B1 — the uniqueness clause is false at 6 occurrences
   evidence: as above — 1 hit in `app` @ `ad9f3c49`, 5 in the **pinned** `stack-demo/rosetta-extensions` @ `09d06070`, 0 in the other 13 trees including both nested `studio` checkouts at `aeec036a`. The seat's 44/42 arithmetic reproduces exactly and settles the "clone set" scope question against the document.

r26-C B2 | corpus/services/jobsimulation.md:208 | UPHELD | IN-SCOPE | unpinned `app/main.go:670` names a SKILLER Azure assignment; the jobsim fatal is `:723`
   evidence: `stack-demo/app` @ `ad9f3c49` (checkout **and** `origin/main`): `main.go:669-670` is `if v := os.Getenv("SKILLER_AZURE_OPENAI_ENDPOINT_URL"); v != "" {` / `skillerAzureEndpointEu = &v`; `log.Fatalf("jobsim-in-app: engine wiring failed …")` is at **`:723`**. `:670` is correct only at `9d00a313` and `7177374`; `:614` @ `b948604` is correct and I verified it. The sentence's own parenthetical pins **only** the `:614` alternative — under §5 rule 33 a pin's scope is the claim's own wrapped sentence, and the neighbouring sentence's refs are explicitly attached to the `main.go:216` anchor. Not ref-discipline: the `:670` half names no ref, and the twin at `:237-238` **is** pinned (`9d00a313`) and grades TRUE — one twin pinned, one left bare.

r25-D B1 | corpus/services/clerkenstein.md:276 | UPHELD | IN-SCOPE | "sentinel and storage are still on v0.34.3" — sentinel is on colony v0.35.2
   evidence: `stack-demo/sentinel` @ `f2c46190` (**= `origin/main`**, so both candidate refs agree), `go.mod:8` = `github.com/anthropos-work/colony v0.35.2`. Bumped at `88036d7` *"chore(deps): update dependencies to latest versions"* (2026-08-03); it read `v0.34.3` only at the superseded `88bc5592`, which I confirmed by `git show`. All seven Go clones at their own refs: app v0.35.2 · sentinel **v0.35.2** · messenger v0.35.2 · cms v0.35.1 · jobsimulation v0.35.1 · storage v0.34.3 · roadrunner v0.34.3 — `storage` holds, `sentinel` does not, and `storage` has had no compose service and no `repos.yml` entry since `838d907`. The sentence names no ref of its own (the preceding sentence's *"At platform `2adcf71`"* is attached to `app/go.mod`, and a **platform** sha cannot date a **sentinel** `go.mod` claim in any case) and is present-tense *"still"*. The corpus's own twin is handled correctly and is thereby inconsistent with this one: `corpus/architecture/shared_libraries.md:57` pins the identical figure to `sentinel/go.mod:8` @ `88bc5592` and grades TRUE there.

r25-D B2 | corpus/architecture/service_taxonomy.md:509, corpus/architecture/service_taxonomy.md:166 | UPHELD | IN-SCOPE | `ai` counted among private modules "pulled at Docker build"; no built service requires it
   evidence: `git show HEAD:go.mod` per clone at its own ref — `anthropos-work/ai` is **ABSENT** from app `ad9f3c49`, sentinel `f2c46190`, storage, messenger, roadrunner; present only in `cms` `ca50c817` and `jobsimulation` `462343b0` (`v1.40.2`), which have no compose service and no `repos.yml` entry. Not transitive: `anthropos-work/ai` = **0** in `app/go.sum` (control: `anthropos-work/colony` = 2). `app` dropped it at `1e457fa70` (2026-08-04, *"refactor(ai): fold the ai library into app as internal/ai"*) and fenced the return — `app/internal/ai/module_import_guard_test.go` and `app/.github/workflows/ai-module-guard.yml` both exist and are tracked (r26-D's contrary sub-claim that `app` has no `internal/ai` is wrong — `git ls-tree HEAD internal/ai/` lists 17 entries — but it does not touch the booked predicate). `app/go.mod` @ `ad9f3c49` requires colony · proto · taxonomy **plus** `analytics-go v0.3.1` and `storage v0.15.2`, neither in the table, so "4 imported" is wrong in both directions. Cross-file drift confirmed: `corpus/architecture/shared_libraries.md:126` already carries the corrected form (*"**No repo a stack builds** … `app` **dropped** the module at `1e457fa70`"*), and `:166` links that file as the "Full reference".

r25-D B3 | corpus/architecture/frontend_architecture.md:59 | UPHELD | IN-SCOPE | `apps/web/package.json:46` quoted as `"^16.2.7"`; it reads `"~16.2.12"` at the checkout
   evidence: `stack-demo/next-web-app` @ `8297c684`, `apps/web/package.json:46` = `"next": "~16.2.12",`; same value in `apps/hiring:45`, `apps/integration:28`, `apps/maintenance:9`. `"^16.2.7"` is exactly `bb3313bc:apps/web/package.json:46`. The claim names no ref and is present tense (*"reads"*), so it grades at the checkout — and the same file at `:41` explicitly declares `8297c684` authoritative (*"the ref this stack builds … 41 commits / 192 files past the `bb3313bc` the old figures came from"*) eighteen lines above. Caret→tilde is a semantic change, not cosmetic drift.

r26-D B1 | corpus/architecture/frontend_architecture.md:59 | UPHELD | IN-SCOPE | same predicate as r25-D B3 — the `next` literal is false at the ref the file declares authoritative
   evidence: as above, re-derived at `8297c684` — `~16.2.12` in all four apps; `^16.2.7` survives only at `bb3313bc`.

r26-D B2 | corpus/services/clerkenstein.md:276 | UPHELD | IN-SCOPE | same predicate as r25-D B1 — sentinel is on colony v0.35.2, not v0.34.3
   evidence: as above — `sentinel/go.mod:8` @ `f2c46190` = `colony v0.35.2`; `88036d7` (2026-08-03) moved it; `v0.34.3` holds only for `storage` and `roadrunner`, both frozen legacy. (The artifact pin the sentence defends is itself correct: `stack-demo/rosetta-extensions/clerkenstein/go.mod` @ the **pinned** `09d06070` = `colony v0.34.3`.)

r26-D B3 | corpus/architecture/service_taxonomy.md:47 | UPHELD | IN-SCOPE | `.env.example:45` cited for VITE_GRAPHQL_ENDPOINT; :45 is VITE_ENVIRONMENT=production
   evidence: `stack-demo/studio-desk` @ `41ee3575` (**= `origin/main`**; the block names no ref), `.env.example:44` = `VITE_GRAPHQL_ENDPOINT=http://localhost:8082/graphql/query` and `:45` = `VITE_ENVIRONMENT=production`. It was `:45` at the superseded `14a5442a`, so this is drift introduced by the clone advancing — but the anchor as written resolves to a different construct at the only ref it can be graded at, and the passage it supports is a *correction* of a previously-backwards diagram. Value, variable and file are all correct; the line is not.

r26-D B4 | corpus/architecture/service_taxonomy.md:166, corpus/architecture/service_taxonomy.md:509 | UPHELD | IN-SCOPE | same predicate as r25-D B2 — `ai` is pulled at Docker build by nothing a stack builds
   evidence: as above — `anthropos-work/ai` absent from app/sentinel `go.mod` **and** `go.sum` (0 hits, control colony = 2), present only in the two frozen repos `make init` does not clone. The seat's supporting claim that `app` has no `internal/ai` package is itself false (the package exists, 17 tracked entries, with a one-way-door guard test); the booked predicate is unaffected.

r26-D B5 | corpus/services/clerkenstein.md:3 | UPHELD | IN-SCOPE | "Last updated: 2026-07-14" is false and self-contradicted by the file's own body
   evidence: `git log -1 --date=short -- corpus/services/clerkenstein.md` → **`328ece5 2026-08-05`** (prior `b925199` 2026-08-02, `05fbcde` 2026-08-02). The body documents work strictly later than the header's M218 terminus and its date: `:54` *"✅ RESOLVED at M219"*, `:61` the `M219 … 97.2% -> 100%` DNA quote, `:232` *"v2.4 'casting call' M224"*, `:161`/`:185` *"v2.8 M256"*, `:274` *"v2.8 M257x iter-23"*. Not a retraction — both the stamp and the post-M218 content are asserted live, which is the self-contradiction case rule 5 names. Metadata rather than mechanism, and I record that; the claim is nonetheless false against the repository's own history.

BOOKED=12 UPHELD=12 REJECTED=0 IN-SCOPE-UPHELD-BLOCKERS=12

---

## Note on the zero-rejection result

I looked hard for the two classes the brief warns about and found neither firing.
**Ref-discipline: zero occurrences.** Every booking above was checked for a pin in the claim's own
block before I graded it, and in each case there was none — `jobsimulation.md:208`'s `:670` pins only
its `:614` alternative; `clerkenstein.md:276` carries a *platform* sha in a **neighbouring** sentence,
about a different repo; `frontend_architecture.md:59`, `service_taxonomy.md:47`, `:166`, `:509` and
`clerkenstein.md:3` name no ref at all. Where a seat *could* have booked a pinned claim and correctly
declined, I checked the decline and agree: r25-C's `jobsimulation` M810 note (true at `6092c6d2`, false
at the checkout `462343b0`) and both C seats' `messenger/terraform/main.tf` observation (true at
`origin/main` `e9421c68`, false at the checkout `fa47850d`) were left in Minors / "recorded, not booked",
which is the pin working.
**Wrong-tree: zero occurrences.** All four seats stated they graded rext claims at the pinned per-stack
clone `09d06070`, and the one booking that turns on a rext tree (r25-C B1 / r26-C B1) reproduces
identically there; I re-measured it at `09d06070` myself.
One seat sub-derivation is wrong without changing its verdict: r26-D B4's *"`app` has no `internal/ai`
package"* (its `git ls-tree HEAD internal/ | grep '^ai'` matches against the mode field, not the name).
`app/internal/ai/` exists with 17 tracked entries. The booked predicate stands on the `go.mod`/`go.sum`
measurement, which is correct.

## Predicate groups

- **tf-literal-uniqueness**: r25-C B1 (corpus/services/jobsimulation.md:50) + r26-C B1 (corpus/services/jobsimulation.md:50) — an absolute uniqueness quantifier ("one occurrence anywhere in the clone set") over a set the same sentence defines to include `rosetta-extensions`, where the literal occurs 6 times in 4 files across 2 repos. **1 predicate, 1 anchor, 2 seats.**
- **archoverview-321-wrong-construct**: r25-C B2 (corpus/services/jobsimulation.md:146) — the intra-corpus citation for the "correctly-scoped model wording" points at `architecture_overview.md:321`, the **production** Cosmo-Router line, where the wording it names is at `:331` inside the local-stack fence. (Same wrong anchor recurs at `corpus/services/sentinel.md:85` and `corpus/services/backend.md:54`, outside these seats' file sets.)
- **mainGo-670-unpinned**: r26-C B2 (corpus/services/jobsimulation.md:208) — an unpinned `app/main.go:670` for the jobsim wiring fatal, which is `:723` at `ad9f3c49`; `:670` is a SKILLER Azure endpoint assignment there. Distinct from the group above: different anchor, different file-under-citation, different failure (stale offset vs. wrong target line in a live file).
- **sentinel-colony-pin**: r25-D B1 (corpus/services/clerkenstein.md:276) + r26-D B2 (corpus/services/clerkenstein.md:276) — *"`sentinel` and `storage` are still on `v0.34.3`"*; sentinel is on `v0.35.2` at `f2c46190`, so the clause that softens the artifact-drift warning now rests on two frozen legacy repos alone. **1 predicate, 1 anchor, 2 seats.**
- **ai-module-imported**: r25-D B2 (corpus/architecture/service_taxonomy.md:509 + :166) + r26-D B4 (corpus/architecture/service_taxonomy.md:166 + :509) — `ai` listed among the private Go modules "pulled at Docker build"; no repo a stack clones or builds requires it, directly or transitively. **1 predicate, 2 anchors, 2 seats.**
- **next-version-literal**: r25-D B3 (corpus/architecture/frontend_architecture.md:59) + r26-D B1 (corpus/architecture/frontend_architecture.md:59) — `apps/web/package.json:46` quoted as `"^16.2.7"` where it reads `"~16.2.12"` at `8297c684`, the ref the same file declares authoritative 18 lines above. **1 predicate, 1 anchor, 2 seats.**
- **studio-desk-envexample-line**: r26-D B3 (corpus/architecture/service_taxonomy.md:47) — `studio-desk/.env.example:45` cited for `VITE_GRAPHQL_ENDPOINT`; at `41ee3575` that variable is `:44` and `:45` is `VITE_ENVIRONMENT=production`. (r25-D measured the identical fact and graded it MINOR rather than booking it, so it contributes one seat, not two.)
- **clerkenstein-currency-header**: r26-D B5 (corpus/services/clerkenstein.md:3) — *"Last updated: 2026-07-14"* against a last commit of 2026-08-05 and a body citing M219/M220/M224/M256/M257x-iter-23. (r25-D and r25-C measured the same shape — at this anchor and at `corpus/architecture/alignment_testing.md:3` respectively — and graded both MINOR, so this contributes one seat.)

**8 distinct false predicates across 12 upheld bookings and 9 corpus anchors.**
