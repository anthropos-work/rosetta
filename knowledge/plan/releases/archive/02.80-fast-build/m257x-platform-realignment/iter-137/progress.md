**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*), working the
adjudicated work list `FIX-M257x-iter135-adjudicated-live-defects` **by consequence**, per iter-136's
lesson 3.

# iter-137 — roadrunner was wrong in two directions at once

## What the corpus said, and what is true

Two adjudicators, blind to each other in iter-135, reported roadrunner defects pointing **opposite ways**.
Both were still live at HEAD (`6a872c0`), unrepaired by iters 132–136:

| | the corpus | measured at source |
|---|---|---|
| **Q1** | *"'There is no roadrunner service in production' overstates it — `roadrunner/terraform/main.tf:19` still reads `service_desired_count = 1`"*; *"the one row where prod and the platform's own declaration contradict each other — recorded, not resolved"*; *"treat retirement as pending, not done"* | **No roadrunner ECS service.** `infrastructure` @ `13c248e6`, `terraform/production/services.tf` declares **exactly ten** service modules — sentinel · directus · acm_media_certificate · storage-service · next-webapp · backend · jobsimulation · studio_desk · db-backup · metabase — and **`module "roadrunner"` is not among them** |
| **Q2** | *"**Eight** services are folded into `app`"*, roadrunner among them; *"**roadrunner domain**: Judge0 code execution"* | **`app/internal/roadrunner/` exists at no ref and was never added** — `git log --all --diff-filter=A -- internal/roadrunner` → **0 commits, ever** |

**A reader could come away believing roadrunner is running in production *and* that it is a package inside
the monolith.** Neither is true. `D-M257x-137-2` records why that matters more than either error alone:
**a subject wrong in two directions has no owner.**

## Ground truth, re-derived at source — not taken from the adjudicators

`D-M257x-136-1` (*a count claim is graded by re-enumerating, never by accepting the reporter's candidate*)
was applied to all four checks. Each carries a positive control.

| # | check | result | control |
|---|---|---|---|
| **G1** | `app/internal/roadrunner/` at HEAD `ad9f3c498` | **absent** | `jobsimwiring` → 3 paths |
| **G1b** | ever added, any ref | **0 commits** (`--all --diff-filter=A`) | full clone, `is-shallow=false`, **6,728 refs** |
| **G2** | where Judge0 is actually wired | `internal/jobsimwiring/wiring.go:123` → `jsrunner.NewRunnerManager(getenv("JUDGE0_API_KEY"), getenv("JUDGE0_BASE_URL"))`, against `internal/jobsimulation/runner` — **inside the jobsimulation domain** | the source's own comment: *"replaces the removed **roadrunner RPC edge**"* |
| **G3** | `roadrunner/terraform/main.tf` shape @ `87d8d443` | 95 lines; `:10-11` `module "roadrunner" { source = ".../base_internal_service" }`, inputs fed from **unbound `var.*`** — **a module awaiting a caller that does not exist** | `:19` `service_desired_count = 1` is one of those inputs |
| **G4** | `roadrunner` org-wide in `infrastructure` @ `13c248e6` | **7 hits, 0 of them terraform** | `^module "` in `services.tf` → **10** |

**G4's seven, itemised** — because "appears nowhere" and "appears only as a name" are different facts:
`.github/workflows/wf-terraform-deploy.yml:209-211` and `wf-terraform-plan-preview.yml:241-243` inject
`production_roadrunner_judge0_{api_key,base_url}` into `TF_VAR_judge0_*` under a `# Roadrunner` comment
(6 lines), plus `knowledge/service-dependencies.md:119`.

**And the read supplied a POSITIVE fact the corpus never had.** Those variables are consumed by
**`module "backend_euwest1"`** (`terraform/production/services.tf:384-385`). **Production wires Judge0
straight into `backend`, under roadrunner-named secret keys — that is the fold, visible at the config
layer**, and it is far better evidence than the count the corpus was reading. The platform says the same
thing in its own words at `infrastructure/knowledge/service-dependencies.md:119`: *"**Judge0** (code
execution — called directly now; `roadrunner` is off this path)."*

## Why both errors survived eleven iters of repair

`roadrunner/terraform/main.tf:19` is **exactly** the artifact `org-repos.md` § 3 exists to disarm — *a
service repo's own `service_desired_count` is not evidence of production state.* That rule was written at
iter-123, applied to `cms` at iter-124/127, to `messenger` and `graphql-wundergraph` in the same table —
and **never applied to roadrunner**, whose row sits in the same table. `adj-F` found the sharpest form of
it: `architecture_overview.md`'s **CMS row carries the full correction and the Roadrunner row one line
below does not.** A half-applied repair, inside a single table, one row apart.

## Repaired — 29 sites in 15 files

**Width measured first (§5 rule 57):** four independent searches, two per conjunct, before any edit.

| file | sites |
|---|---|
| `CLAUDE.md` | 4 — the eight-services banner, the *"plus the … roadrunner domains"* lead-in, the **"roadrunner domain"** bullet, the archived-list row |
| `corpus/services/roadrunner.md` | 5 — the MERGED banner, the *"overstates it"* precision block, the *"recorded, not resolved"* clause, the `:53` separator, the *"retirement as pending"* close |
| `corpus/architecture/architecture_overview.md` | 6 — the eight-microservices list, the **Domains-inside-Backend/App Roadrunner bullet**, the both-tables note, the Roadrunner row, the in-process ladder, the gRPC-hop note |
| `corpus/architecture/service_taxonomy.md` | 4 — the mermaid monolith node, the gone-from-compose list, the Roadrunner row, the archived/merged row |
| `corpus/architecture/dependency_map.md` | 3 — the matrix-collapsed banner, the Roadrunner row, the streams row |
| `corpus/services/README.md` | 2 — the *"eighth and different"* banner, the roadrunner index row |
| `corpus/README.md` · `README.md` · `corpus/ops/{README,platform_repo,update_guide,run_guide}.md` | 6 |
| `.claude/skills/{dev-up,dev-for-dummies×2,stack-update}` | 4 |
| `corpus/architecture/{org-repos,platform-migration-status}.md` | 2 — the positive production evidence added; the stale *"a repo this map has never read"* retired |

**The verification sweep found THREE the planning sweep did not**, and that is `D-M257x-137-4`:

| survivor | why the planning search missed it |
|---|---|
| `architecture_overview.md:357` — *"it was folded in with jobsim-in-app"* | sits in a paragraph about **gRPC hops**; no fold-vocabulary adjacent to "roadrunner" on the line |
| `services/README.md:57` — *"but prod terraform still reads `= 1`"* | an **index row**, phrased as a caveat rather than a claim |
| `architecture_overview.md:22-23` — a **Roadrunner bullet under a heading reading *"Domains inside Backend/App, not services"*** | the falsity is carried by the **heading**, not the bullet; no line-scoped grep can see a predicate asserted one level up. **Found by reading the file, not by any search** |

**Re-run the search with your own corrected wording excluded; what survives is what the planning search
could not see — and the third one says even that is a floor.** A predicate asserted by a section heading
and inherited by its bullets is invisible to every line-oriented instrument this milestone owns.

## The fences fired on my own repair, and were right

`anchor_construct_guard` and `repair_postcondition` both went **RED** after the edits. The cause is worth
the space: `roadrunner.md`'s async-tasks paragraph carried a bare line pin **as its own worked example of a
bad pin** — *"this said `:124` below, and at iter-120 `:124` was above this very line"* — and my repair
shifted the file until that quoted pin landed on a **blank line**.

**Fixed by deleting the pin rather than re-pinning it** (re-pinning restarts the same clock). This is
`adj-E`'s rotted-anchor class arriving in its purest form — *the citation that rotted was inside the
sentence warning about rotting citations* — and `D-M257x-137-3` states the general rule: **a fence matching
on form cannot distinguish asserting a pin from quoting a retracted one.** Same blindness iter-132 found on
hedge markers and iter-134 measured at 1-of-4 fences; here on the **anchor** axis. It is the guard's
declared floor working correctly.

## Side-deliverable — the 16th escape of the cms-M810 predicate, on the corpus front door

**Separate commit, and it does NOT grade this iter.** While editing `corpus/README.md` for Q2, line 18 read
*"M810 … is **uneven**: landed for jobsimulation, **not moved for cms**"* — the claim **retracted at
iter-124 and swept corpus-wide twice** (iter-127: 5 sites; iter-132: 15 sites). Width re-measured before
touching it: **1 live site**; every other match is a retraction quoting the old wording. Repaired with the
`13c248e6` citation.

Rule 55 in one line: **a reader who wants the migration state opens `corpus/README.md`**, and neither
sweep's search reached it.

## Test gates

| gate | result |
|---|---|
| **Guard family** (`--repo-root` + `--platform stack-demo/platform @ 0c91421`) | **18 GREEN · 0 RED · 4 not-run** — the 4 are commit/input-scoped (`anchor_offset`, `repair_leak`, `repair_reach`, `value_change`; need `--range`/`--ledger`). **The runner's own line says this is NOT a whole-family green, and neither does this iter.** |
| Guard family **before** the anchor fix | **16 GREEN · 2 RED** — both RED on my own edits; repaired by fixing the artifact, not the fence |
| **Scoped fence suites** (9 files: anchor-construct-denominator, corpus-citation, corpus-index, claim-twin, claim-census, platform-alignment, platform-predicate, repair-postcondition, m257x-mechanical-fences mutation battery) | **367 passed / 0 failed** in 623.05 s |
| **Whole suite** | **NOT re-run — §5 rule 60 requires saying so out loud.** **Zero `rosetta-extensions` files changed this iter**; iter-132's clean whole-suite run stands on the same rext tree (`223e4a6`), and iters 133–136 also changed none. The exposure is bounded to corpus content, which the 9 scoped suites + the 18-guard family cover. **Stated as a gap, not characterised as covered.** |
| **Suite wall-time** | **not quoted as a measurement** — `FIX-M257x-iter132-suite-walltime-is-not-a-measurement` stands; this host is not a stable timing substrate |

## Close — 2026-08-08

**Outcome:** roadrunner was asserted **both** to be running in production **and** to be a package inside
`app`. Neither is true, both were live at HEAD, and they contradict each other. **29 sites in 15 files
repaired** across both conjuncts, with ground truth re-derived at source under positive controls — and the
`infrastructure` read supplied the positive production evidence the corpus never had: **`module
"backend_euwest1"` receives the Judge0 credentials, under roadrunner-named secret keys.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged; no reading taken.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**the last three tiks took no reading, so the metric is UNMEASURED, not unmoved — §9's iter-type refinement; and `TOK-08`'s sealed refutation branch bars an agent-authored successor**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (**1 tik this session**) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-137-1` (*merged* vs *deleted* — only one predicts a package) · `D-M257x-137-2` (a
subject wrong in two directions has no owner; sweep the subject, not the sentence) · `D-M257x-137-3` (a
retraction that quotes the retracted pin re-publishes it) · `D-M257x-137-4` (the planning search and the
verification search must not be the same search).
**Side-deliverables:** `corpus/README.md:18` — the cms-M810 predicate's 16th escape, on the corpus front
door, after two corpus-wide sweeps. Separate commit; does not grade this iter.
**Routes carried forward:**
- `FIX-M257x-iter135-adjudicated-live-defects` — **the remainder, still the obvious next target**:
  `shared_libraries.md:77` (analytics-go wiring cited `main.go:507-508`, measured `:494-495`) ·
  `security_compliance.md:156` (→`clerk-integration.md:44`, not `:40`) · `clerk-integration.md:126` ·
  `backend.md:13`'s dangling *UNEVEN bullet* · `sentinel.md:5` · `ai-readiness.md:18-20` ·
  `org-repos.md:227`,`:370`,`:43` · `ai_architecture.md:111`,`:224` · `next-web-app.md:17`,`:186` ·
  `external_services.md:368` · `adj-E`'s five rotted anchors. **`dependency_map.md:9` and roadrunner's
  five anchors CLOSE this iter.**
- `FIX-M257x-iter135-bare-pin-blind-spot` — **strengthened by this iter's own RED**: `D-M257x-137-3` is a
  live instance, and `adj-E`'s five rotted anchors are the same class.
- `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer` ·
  `FIX-M257x-iter133-two-fives-need-a-fence` · `FIX-M257x-iter132-suite-walltime-is-not-a-measurement` ·
  `FIX-M257x-iter131-predicate-sets-not-enumerated`.
- **CLOSED this iter:** `FIX-M257x-iter132-infrastructure-is-cloneable-so-clone-it` — cloned, read, and
  the read settled roadrunner's production state *and* upheld `org-repos.md:143` byte-for-byte.
**Lessons:**
1. **A subject wrong in two opposite directions is a subject with no owner.** Repairing one conjunct makes
   the corpus look more consistent and no more true. Sweep the subject.
2. **Apply a rule to every row of the table it was written for, on the day you write it.** § 3's
   *"`service_desired_count` is not evidence"* reached cms, messenger and wundergraph at iter-123 and
   skipped roadrunner — one row away, in the same table, for four days.
3. **Verify with a different search than the one you planned with.** Four searches planned a 26-site
   repair; re-running them with the corrected vocabulary excluded found **two more**.
4. **Never quote a retracted line-pin.** Describe the artifact. A form-matching fence cannot tell the
   quotation from the assertion — and the pin will rot again, as this one did, twice.
