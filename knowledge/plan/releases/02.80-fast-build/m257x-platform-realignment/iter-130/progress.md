**Type:** tik — `TOK-08`'s repair-and-fence step. Per the protocol's *"Iter-type refinement — the
3-no-prog tok-trigger reads UNMEASURED as UNMEASURED"* (iter-108), **this iter took no reading and
no `N` movement is claimed**; the grading step is iter-131, declared in advance here and in the
overview.

---

## 1. Priority 1 — the routed residual, closed, with every finding re-verified at a ref

`iter-129/progress.md` § 2c and § 3c routed ~24 findings as `FIX-M257x-iter129-sweep-residual` rather
than half-repairing them. All are closed. **Nothing was repaired on the strength of the prior sweep's
prose** — each was re-opened at the ref, which is the only reason an *upheld* finding below counts as a
result rather than a miss.

### 1a. Graded by consequence, not by class

**The sharpest is a security-surface claim, and it is the third of this milestone.** `latency-budget.md`
credited the fake Clerk FAPI with **validating `redirect_url` against the public origin**. It does not:
`clerkenstein/clerk-frontend/server.go:414-423` @ rext `415240f` reads the query param, defaults it to
`/`, appends `__clerk_handshake` and issues a 303 — **no scheme, host, or allowlist check anywhere**
(`grep -n redirect_url clerk-frontend/*.go` → one non-test hit). The mock is an **open redirect that
mints and forwards a session with no credential**.

**That is the designed disarm, and the repair says so rather than raising an alarm.** The control was
never a validation step in the mock — it is `safety.md` §3 (*there is nothing behind the door*) plus the
tailnet scope. What was false is only the claim that a check stood in front of it, and a doc that
credits a component with a property it lacks is how the real control stops being maintained.

**The other consequence-class findings:**

| finding | measured at the ref |
|---|---|
| `staging-bringup.md` — *"the GraphQL ent privacy layer is unaffected"* | **refuted by 30 `OrganizationMixin` schemas** whose Query policy opens with `DenyIfNoOrganizationInContext` (`internal/data/ent/schema/mixin.go:137-141`; rule chain `rule/organization.go:15-25`, `:29-37`) @ `ad9f3c498` |
| **sentinel + Atlas**, 4 sites | `68272003` (2026-08-04) added a **second Atlas pipeline, owned by `app`** — `app/atlas.hcl:50-64` declares `env "sentinel"`, and `app/Makefile:59-60` says *"`atlas migrate apply --env sentinel` creates the schema itself, and that is what local/CI actually run"*. `repos.yml`'s `migrations: false` is still right **about the sentinel repo** (no `atlas.hcl` at `f2c461903`) — which is why the flag and the pipeline are both true |
| **`CORS_EXTRA_ORIGINS` documented as unlanded**, 6 sites | landed in `f664473` (2026-05-14) + `13410de` (2026-05-19), both `merge-base --is-ancestor` of `ad9f3c498`; `internal/cors/cors.go:24`, applied `:78-82` behind a `!IsProduction()` guard. **The skip-worktree patch it justified is now retirable** — keeping the mark hides real upstream changes |
| `staging_from_dump.md` — 3 of 4 notification examples | `welcome` / `recap` / `password` → **0 hits** in `app/internal/messenger` AND in the `messenger` clone. **Password resets are Clerk's** (`<SignIn>` at `next-web-app` `login/page.tsx:4`, `:40`), so blanking `BREVO_KEY` does not suppress them. Widened `password` across all of `app/internal/` → 16 hits, **none a sender** |
| `safety.md` — the **Bunny recording-key provisioning path** | **does not exist.** `BUNNY_RECORDING_*` → **0** in all of rext; widened to `bunny` case-insensitive → **39**, *the number moved and the verdict did not* — none is a recording-key site. The M239 bridge (`up-injected.sh:1358`) iterates a **fixed five-key Bedrock list** (`:1362`). The exhibit is blocked on **two** things, values *and* a path; only one was disclosed |
| `safety.md` — `AssertClean` *"every attempted write is recorded"* | true **only of the BLOCKED path**, which `stack-seeding/isolation/audit.go:82-96` states in as many words. Recording on the allowed path is voluntary per-seeder (`seeder/dag.go:197-207`). The gap is closed by a **separate** assert, `AssertRecorded` (`audit.go:97-115`) — a two-assert proof was documented as a one-assert one |
| `safety.md` — attacker-gain omits the **bridged AWS credentials** | upheld gap, closed as **§3.4 residual #4**. The `~/.aws` mount is cleared by both emitters, but M239's `bridge_bedrock_creds` writes `AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY` into `platform/.env`, which `docker-compose.yml:44-45` hands to `backend`. **The vehicle moved; the exposure did not close** |
| `secrets-spec.md` — the waived class | **as written it would have FALSE-FAILED its own gate.** `platform/OPENAI_KEY` and `platform/AZURE_OPENAI_KEY` are `critical·required` with `nonempty` — **2 of the 13 critical genes** — and the gate is `Critical == 1.0` with waived excluded from both denominators. Also *"the one secret"* is **two** genes (`secret-dna.json:191`, `:659`) |
| `demopatch-spec.md` — *"FOUR patches"* | **seven** (`up-injected.sh:1124-1126`). §4 of the same doc had **already** moved this 4→7 and left §5 behind — the mirrored-count failure the doc itself defines two boxes above. Apply/revert is **not LIFO** on either frontend lane; `studio-desk` alone is |
| `build-budget.md` — *"eleven sub-phases"* | **twelve** (`buildbench.py:114-130`). Widened `grep -rn BRINGUP_ANCHORS` over all rext → 2 hits, **count did not move**. The tables below the prose already said twelve |
| `recipe-browser-login.md` | names a **`cms` container** (there is none — 5 services at `0c91421df`) and **`extra_hosts`**; the live path is a docker network **alias** (`gen_injected_override.py:823-829`), and the file that writes `extra_hosts` is **never passed to compose** |
| `db-access.md` — the `274/733` split | measured on **`cms.similarities`** (`simembeddings.go:44` `const Schema = "cms"`), not `public.similarities` — a separate table the fold created |
| `platform-alignment.md` `:46-47`, `:308-314` | anchors on rewritten code → re-pinned **@ `38a4214`** as historical with live pointers; the `:308-314` mechanism refuted (both routes sit in the **same** `isPublic` matcher, `proxy.js:139-140`/`:170`) and replaced with the measured discriminator — the route group's layout (`(authed)/layout.jsx:24` `<ClerkProvider>` vs `(public)/layout.jsx:1-7`) |
| `setup_guide.md` — the `env_file`-already-carries-it premise | `env_file: .env` on **4 of 7** services; build values come from `args:`. The Studio-Desk symptom is a blank **value**, not the wiring |
| **UPHELD** — `demopatch-spec.md` §5's *"23 patches: 11 next-web · 2 app · 5 ant-academy · 5 studio-desk"* | **exactly correct.** `ls demo-stack/patches/*/*.yaml \| wc -l` → 23; `grep -h "^repo:" … \| sort \| uniq -c` reproduces the split. The routed "4 vs 7" was the *hiring-image* count, a different number in the same section |
| **UNCHECKABLE** — the second colony bug | colony is in no clone set; only `v0.34.3` is in the module cache. Bug 1 **is** fixed (`v0.34.4`, `b810b28`) and `app/go.mod:15` is now `v0.35.2` with **no colony `replace`** — the vendor recipe is dead and is marked so |

### 1b. The `@anthropos.work` predicate — and rule 57 is the finding, not a footnote

A claim refuted at iter-115, found at **14, then 15, then 16 sites** across successive enumerations. The
brief required measuring **this run's own enumeration width first**. Four independent regexes over
`corpus/**` + `CLAUDE.md` + `README.md` + `.claude/**`:

| enumerator | hits |
|---|---|
| E1 — the literal `@anthropos.work` (**what prior repairs used**) | 28 |
| E2 — internal-portal / employees-only vocabulary | 11 |
| E3 — any academy line mentioning employee/internal/domain/gate | 41 |
| E4 — `anthropos.work` with any leading char | 31 |
| **union** | **63 lines / 28 files** → triage → **10 genuinely false** |

> **Only 4 of the 10 false sites carry the literal `@anthropos.work`.** The other **6** say
> *"internal learning portal"* or *"Standalone Internal Apps"* with **no domain token at all** — and
> they are structurally invisible to the literal regex every prior repair used. **That is why the count
> kept growing 4 → 14 → 15 → 16: not because sites were being missed by carelessness, but because the
> predicate has a second surface form nobody had enumerated.**

Repaired: `services/README.md:58` (**the index `CLAUDE.md` tells every reader to start from**),
`external_services.md:55` (a **mermaid node label** — `Academy[… @anthropos.work only]`),
`service_taxonomy.md` ×2 (`:213-214` Purpose/Users, and the *"Standalone Internal Apps"* category),
`ant-academy.md:360` (**a runnable prerequisite** telling a developer their Clerk credentials must be on
the employee domain), `run_guide.md:230` (a **heading** whose own body one line below already retracted
it), `README.md` ×2, `CLAUDE.md:297`, `.claude/skills/stack-update/reference.md:55`.

**And the refutation is now mechanism-grade rather than an absence.** `ant-academy` `d5875e34`
*"replace @anthropos.work email gate with Clerk org-membership check"* **dropped the allowlist**; at
`22df69dd8` the gate is **any Clerk organization membership, active or not** (`code/proxy.js:2-5`,
`:298`), `/` is in the public matcher (`:112`), and `REQUIRE_ORGANIZATION_MEMBERSHIP=0` skips it
entirely (`:73`, `:91`). *"0 hits"* was true and weak; the commit that removed it is the citation.

**Note the corpus's excuse, because it is real:** `ant-academy`'s **own** `CLAUDE.md` still publishes the
dead gate at three sites (`:11`, `:89`, `:215`). The corpus inherited a claim the source repo's
documentation still makes — which is an argument for reading code at refs, not docs.

**The substrate was dirty and it was read at the ref anyway.** `stack-demo/ant-academy` has 3 modified
files in its working tree; every quotation above is `git show 22df69dd8:<path>`, never the tree.

### 1c. `anthropos-labs.md` — a DIFFERENT SUBJECT, deliberately not repaired

`anthropos-labs.md:12` (*"Access: `@anthropos.work` emails only"*) and `:87` (a *"Domain check"* flow
step) are about **`anthropos-work/experiments`**, a different application in a different repo. It is in
**no clone set** and **no `repos.yml`** (0 hits @ `0c91421df`), and the page carries no measurement date.

**So it is UNCLONABLE, not refuted** — the split rule 56 exists for. Both sites now say so, and the page
carries an explicit warning **not to transplant the Ant Academy correction onto it**. A correction moved
across subjects is a fabrication wearing a repair's clothes, and repairing this row by analogy would
have produced a confident sentence nobody measured.

### 1d. `chronos` — checked, and already closed

Verified rather than assumed: iter-129 repaired the one **consequential** block (`:122-124`, the
*Upstream Consumers* section asserting live cross-service consumers). The residual present tense at
`:42`, `:99`, `:137`, `:161-166` describes chronos's **internal** design under a top banner that says
*"preserved for historical context"*, and every cross-page mention is correctly historical. **No repair
was manufactured here.**

---

## 2. Priority 3 — the library rows had no fence, and now they have one that fires

### 2a. Why A–F could not have caught it

`platform_alignment_guard` fences the map against `repos.yml` in **both** directions, and it was GREEN
for four days over a false `ai` row. Correctly green: **`ai` is a module, not a clone**, so there is no
`repos.yml` entry for A or B to disagree with. *A fence is green over its reach, and the library rows
were outside this one's.*

### 2b. Assertion G — the module graph as the source of truth

Subject: the `go.mod` of every repo `repos.yml` declares, **read at its ref**. Three directions:

| | fires when |
|---|---|
| **G1** | a row says `library` and **no** declared repo's `go.mod` requires the module — *the `ai` shape* |
| **G2** | a declared repo requires a module the map has **no row for**, in either table — the arrival direction, mirroring A |
| **G3** | a row says `library-unimported` and a declared repo **does** require it — the departure direction, mirroring B, and what stops G1 being "fixed" by a blanket relabel |

**A ninth state token, `library-unimported`, was added for the same reason iter-64 added `mid-fold`: a
real state with no legal word for it.** §1 defines `library` as *"imported as a private Go module"*, so a
library repo nothing imports could only be written as `library` (asserting an import measurement
refutes) or `decommissioned` (an orchestration lifecycle a library was never in). **The `ai` row went
four days wrong partly because the only alternative to `library` was a worse word.**

### 2c. Its first run found three real rows — one of them iter-129's own half-repair

```
[G1 stale-library] ai:    prod says `library` … no declared repo's go.mod requires it
[G1 stale-library] authn: prod says `library` … (same)
[G1 stale-library] authn: fresh local stack says `library` … (same)
```

> **`ai`'s PROD cell was still wrong.** iter-129 repaired the `fresh local stack` cell and left the
> other — **a correction that reached one of the row's two cells, inside the row it cited as the rule-54
> exemplar.** G grades the two cells independently, and a regression test pins exactly that half-repair
> shape. Prod and local were always the same answer here: production runs the `app` image built from the
> **same `app/go.mod`**.

**`authn` is the same class, and nothing had ever caught it.** Neither `app/go.mod` @ `ad9f3c498` nor
`sentinel/go.mod` @ `f2c461903` requires it; both `go.sum`s have **0** hits. **Positive control: `colony`
returns 2 hits in each**, so the search is not vacuous. `rosetta-extensions` requires it in no `go.mod`
either. The corpus had this right in prose — `shared_libraries.md` and `service_taxonomy.md` both say it
— **and wrong in the fenced cell**, which is precisely why the cell needed a fence and not another sweep.

Both rows repaired; G green; **reach printed on every run** — `read 2 go.mod file(s), 5 org module
require(s) (analytics-go, colony, proto, storage, taxonomy), graded 4 library row(s)`.

### 2d. Controls, because eight vacuous fences have been caught here

- **Every mutant asserts it applied** before checking the guard fires (§5 rule 53).
- **Parser positive control:** go.mods present + **zero** org requires **raises**, rather than reporting
  the library rows aligned on a broken regex.
- **Zero-clone case prints `assertion G NOT RUN`** on every run instead of a silent green — the exact
  failure G exists to prevent, one level up.
- **Meta-mutation:** suppressing G's findings kills **4** tests; restoring returns **64/64**.

### 2e. And the map now DECLARES its unfenced classes

The brief's second option was taken **as well as** the first, because it is the durable half. §4 gained a
per-row-class reach table: **every row's MEMBERSHIP is fenced; only the library rows' STATE is**; merged/
decommissioned, external and `repos.yml`-declared rows remain **prose-under-review** and must be
re-derived at each ref bump. *A fenced artifact with a silently unfenced class in it is worse than an
unfenced one, because the green is read as a warrant over the whole file.*

---

## 3. Rule 54 sweeps — the corrections that had to reach more than one cell

Each agent reported out-of-scope sites rather than editing outside its set; all were closed here after
independent verification:

- **`platform_repo.md:111`** and **`services/sentinel.md:97`** — both said sentinel bypasses Atlas.
- **`staging-sync.md`** — the skip-worktree inventory still listed `cms`/`jobsimulation`/`storage`/
  `messenger` (**none in the clone set**) and the retirable `internal/cors/cors.go` mark; and
  **`vendor-colony/` at 2 sites**, dead since the fix landed in colony `v0.34.4`.
- **`media-substrate-spec.md` ×3 + `demo/README.md`** — the same false Bunny provisioning mechanism.
- **`snapshot-spec.md:190`** — `274 sim-embeddings` re-attributed to **`cms.similarities`**.
- **`.claude/skills/demo-up/SKILL.md:192`** — *"Clerk-only"* ant-academy; **`session-clone-spec.md:224`**
  — the *"LIFO revert trap"*.
- **rext, two source comments that contradicted the code beside them** (commit `f2ea567`):
  `isolation.go`'s package doc (*"records every attempted write"*, refuted by `audit.go` two files over)
  and `gen_injected_override.py`'s comment justifying the `~/.aws` removal on the premise that a demo
  carries **no** AWS credentials — false since M239 wired the env bridge. **Neither behaviour changed;
  both were the record a future reader would have trusted.**

---

## 4. Side-deliverable — the front door was false in all three rows

`README.md`'s tier table listed **CMS, Jobsimulation, Storage and Roadrunner as live Tier-1 Go
services**, Studio-Room as *"embedded in CMS"*, and **Wundergraph as the gateway**. Measured at
`0c91421df`: `docker-compose.yml` declares **five** services (`sentinel`, `backend`, `studio-desk`,
`next-web-app`, `gotenberg`; `postgresql`/`redis` via its `include:`) and `repos.yml` **four** repos.
This is the repository's highest-traffic page and no sweep had reached it. Recorded as a side-deliverable
— it is not planned scope and does not upgrade the close status.

---

## 5. Guards, and three REDs that were mine

**18 GREEN · 0 RED · 4 not-run** at open and at close. Invocation, ~32 s wall:

```
/usr/bin/python3 guard_family.py --repo-root <rosetta> --platform <rosetta>/stack-demo/platform
```

**Three fences went RED mid-iter on my own edits, and each was repaired by fixing the artifact:**

1. **`anchor_construct_guard` + `repair_postcondition`** — my `isolation.go` package-doc edit shifted the
   file **+12 lines**, so `safety.md:207` and `seeding-spec.md:102` cited `isolation.go:106`, which had
   become `},`. Re-pinned to **`:118`** and **verified at the construct** (`{Name: "s3-private", Class:
   PerStackIsolated…}`). A third, `clerk-integration.md:99`, had drifted onto a blank line because the
   staging repairs grew the file — re-pinned, **and the paragraph's own advice was applied to itself**:
   it already told readers to cite the stable quirk number, so it now carries the `grep` that re-derives
   all three pins. *This is the third recorded move of those pins; the line is a convenience, the grep is
   the contract.*
2. **`claim_census_guard`** — the map's ratchet rose **8 → 10** on my new §4 prose. **Evidenced, not
   loosened.** Census **1141 → 1140** overall.

**The fences caught the author again, and that remains the only reason their greens mean anything.**

**Suite: not re-run.** Load was checked FIRST (§5): the external `a8-test` process held ~1 of 12 cores
throughout. The rext changes carry their own targeted runs instead — **64 passed** on
`test_platform_alignment_guard.py` (up from 61; the +3 net is the new G battery minus the vocabulary
test's rename), `go vet ./isolation/` clean, `gen_injected_override.py` compiles.
`FIX-M257x-iter128-suite-timing-unattested` stays open.

---

## Close — 2026-08-07

**Outcome:** the routed residual set is closed with every finding re-verified at a ref; the
`@anthropos.work` predicate repaired at **10 sites after measuring the enumeration's own width** (union
63 → 10 false, **6 of them invisible to the literal regex every prior repair used**); the library rows
given assertion **G**, which fired on **3 real rows** including iter-129's own half-repaired `ai` cell,
and the map now declares which classes remain unfenced.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5, unchanged. **No reading was taken and no `N` movement is claimed.**
**Phase 5 grading:** (1) gate-met: n — 4 of 5, unchanged — (2) triggered-tok: n (**this iter took no
reading; per the iter-108 refinement an UNMEASURED metric does not count toward the 3-no-prog streak,
and `TOK-08` declared this step order in advance**) — (3) re-scope: n — (4) user-blocker: n (every
surfaced item was repairable in place or routed; the `experiments` repo is a substrate limit, disclosed,
not a decision needed from the user) — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n —
(7) budget-exhausted: n — Outcome: **continue**
**Decisions:** the `library-unimported` token (a vocabulary change, taken rather than forcing a worse
word); `anthropos-labs.md` disclosed-not-repaired; the README tier table taken as Fate 1 rather than
routed.
**Side-deliverables:** `README.md`'s three-row tier table (in `0c20d8c`); the two rext source comments
(`f2ea567`).
**Routes carried forward:**
- `FIX-M257x-iter130-secrets-manifest-reclassification` → **iter-132+**: `secrets-spec.md`'s waived
  class was repaired in the DOC; the underlying **manifest** classification of `platform/OPENAI_KEY` is
  the real decision and is a change to `secret-dna.json`, not prose.
- `FIX-M257x-iter130-bunny-provisioning-path` → **iter-132+**: the path does not exist; building it is a
  DNA gene pair or a bridge entry, out of this iter's scope.
- `FIX-M257x-iter128-suite-timing-unattested` → open, unchanged.
- `FIX-M257x-iter113-adjudication-is-judgement` → open; **G settles none of the 15 judged verdicts**,
  and this iter does not claim it does.
**Lessons:**
1. **A predicate can have a second surface form, and the count grows until somebody enumerates it
   rather than the token.** 6 of 10 live sites carried no domain token at all. Rule 57 was written about
   *width*; this is width in a second dimension — **vocabulary, not just regex reach**.
2. **A fence's green is read as a warrant over its whole artifact.** The remedy is not only more
   assertions but a **declared reach per row class**, so the unfenced part is visible to the reader
   rather than only to whoever wrote the guard.
3. **When the vocabulary has no honest word for a real state, the row will be wrong and the author will
   not be at fault.** Two rows sat wrong behind one missing token.
