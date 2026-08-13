# seat-10 report

**Files owned:** `corpus/services/sentinel.md` · `corpus/architecture/security_compliance.md` ·
`corpus/architecture/external_services.md` · `corpus/services/README.md`.
**Anchors booked:** 6. **Sites found:** 8. **Sites repaired:** 6 (+1 induced-defect prevention). No commit.

**Ground truth re-derived at this iter's refs** — platform `0c91421`, `app` `ad9f3c49` (== `origin/main`),
`sentinel` `f2c46190`, rext `stack-demo/rosetta-extensions` `09d06070` (the **pinned per-stack** tree, per
§5 rule 45). Nothing was inherited from iter-99/iter-101 sheets. No `git fetch` was run.

**Line-count neutrality — deliberate, and it is the headline safety property of this seat.**
All four files are byte-for-byte the same length as at `HEAD`: `sentinel.md` (unchanged count),
`security_compliance.md` 281 → **281**, `external_services.md` 883 → **883**, `services/README.md` 88 → **88**.
Every edit was reflowed *inside* its existing line span. This was forced, not stylistic: five corpus files
pin line numbers **below** my edit sites — `platform-alignment.md:1513` → `security_compliance.md:227` under
`:231-256`; `CLAUDE.md:271`/`:285`, `security_compliance.md:7`/`:198`/`:201`/`:218`,
`architecture_overview.md:283`/`:285`, `ai_architecture.md:51`/`:60`/`:87`, `shared_libraries.md:130` →
`external_services.md:567` / `:579` / `:602-607` / `:619-629`. Adding a single line would have silently
re-pointed all of them — the exact iter-100 induction class. **Proved, not asserted:**
`diff` of `security_compliance.md` `:106-281` and `external_services.md` `:212-883` against `HEAD` returns
IDENTICAL.

---

## Ledger rows

| # | the false claim | anchor | what is true | sites repaired |
|---|---|---|---|---|
| 1 | "It is also the only *service* address compose sets at all: there are **zero** `*_RPC_ADDR` variables left, so `backend → sentinel` is the one cross-process edge a local stack has." | `corpus/services/sentinel.md:85` | **CANON-1 applied verbatim-in-substance.** The `*_RPC_ADDR`-is-zero half is TRUE and survives: `git grep '_RPC_ADDR' 0c91421 -- '*.yml' '*.yaml' '.env_example'` → **zero** (the only hits in the whole repo are prose in platform's own `CLAUDE.md:34-35`). The generalisation is FALSE: `backend`'s `environment:` block at `0c91421` also carries `GOTENBERG_URL=http://gotenberg:3200` (`docker-compose.yml:57`), `JUDGE0_BASE_URL` (`:59`), `REDIS_ADDR` (`:66`), `SUPABASE_DB_CONN`/`COPILOT_DB_CONN` (`:93-94`). `gotenberg` is declared at `:170` with `profiles: [core, backend, all]` at `:183` — the **default** profile — and is reached over **plain HTTP, not Connect-RPC**: `app/internal/converter/gotenberg.go:31` @ `ad9f3c49` is `req, err := http.NewRequestWithContext(ctx, "POST", gotenbergURL+"/forms/libreoffice/convert", &body)`. Model wording verified live at `architecture_overview.md:321`. | 1 |
| 2 | "Measured at `app` HEAD: of **135** Ent schemas (139 `.go` files, 4 of which declare no schema), only **30** use `OrganizationMixin{}` — the one that carries the privacy `Policy()` (`mixin.go:126`)." | `corpus/architecture/security_compliance.md:67-68` | **30 MENTION it; 29 USE it**, and the sentence's predicate is *use*. At `ad9f3c49`: `git grep -l 'OrganizationMixin{}' -- 'internal/data/ent/schema/*.go'` → **30** files; excluding the one commented line → **29**. The 30th is `internal/data/ent/schema/user_resource.go:22` = `// OrganizationMixin{},  // We need to work on this`. `mixin.go:126` = `func (OrganizationMixin) Policy() ent.Policy {` ✓. 135/139/7 all reproduce exactly. The same blockquote already stated 29 at `:76`, `:84-88` and `:135` — the fence's opening sentence was the only site still saying 30. Also replaced the moving label *"`app` HEAD"* with the pin `ad9f3c49` + date. | 1 |
| 3 | "The genuine remainder — global reference data with no `organization_id` at all — is what carries no org column by design." | `corpus/architecture/security_compliance.md:104-105` | The remainder (135 − 31 policed − 23 unpoliced-with-`organization_id` = **81**) is **not** uniformly global reference data. At `ad9f3c49`, **four** members are per-TENANT: `lab.go:63`, `academy_chapter.go:84`, `academy_skill_path.go:78` each `field.String("tenant_eid")`, and `skill_path_session.go:43` `field.UUID("tenant_id", uuid.UUID{})`. All four carry **no** `organization_id` and **neither** org mixin (measured per-file); the first three declare **no `Policy()` at all** (only 4 files in the directory declare any: `mixin.go`, `org_membership.go`, `organization.go`, `user.go`), and the fourth carries `UserMixin{}` (`skill_path_session.go:26`) whose `Policy()` (`mixin.go:98`) is an **owner** filter — by *user*, never by its `tenant_id`. **Net-new beyond the adjudicator's finding:** a fifth, `academy_feedback.go`, falls in the same remainder and **does** carry an `organization_id` (`:129`), so *"with no `organization_id` at all"* fails on its own terms too. | 1 |
| 4 | "so `backend` is the only consumer left and per-service re-point tooling has ONE target, not two — see the ⚠️ under *Architecture* below." | `corpus/architecture/external_services.md:136` | The *"only live consumer"* half is TRUE and survives. The target count is FALSE: `DIRECTUS_DATA_CONSUMERS = ("cms", "backend")` — **two members** — at `stack-injection/gen_injected_override.py:86` **and** identically in the dev twin `stack-core/gen_override.py:58`, both @ the demo's pinned rext `09d06070`. The source itself explains the shape at `:77-81`: `cms` is an **inert key**, the test *"never matches it on a current clone"*, *"kept only so a ROLLBACK/older platform clone that still DEFINES the container gets re-pointed too."* **The cross-reference was the wrong-construct half:** `:136` routed the reader to the ⚠️ under *Architecture* as corroboration, and `:206` **refutes** it in bold — *"The `--local-content` re-point targets BOTH `cms` and `backend`."*, with the tuple restated at `:209`. Repaired to say the block **qualifies** rather than corroborates. Cited line re-opened and verified after the edit. | 1 |
| 5 | "`rosetta-extensions/stack-injection/gen_injected_override.py:669-670` re-points every service in `DIRECTUS_DATA_CONSUMERS`, which is **`("cms", "backend")`** (`:84`)." | `corpus/architecture/external_services.md:208-209` | **Not booked — repaired as induced-defect prevention, same predicate, my file.** These anchors were byte-exact at the *prior* rext pin `ab81527a` (which is why both readings' bookings of `:208-211` were REJECTED). The demo's pin has since advanced to `09d06070` (2026-08-06, a descendant of `ab81527a`), moving them to `:698-699` and `:86`. Had I left them, my row-4 repair 70 lines above would have cited `:86` while this block cited `:84` for the *same constant* — a self-contradiction I would have manufactured. Not TRAP A: the construct was **not deleted**, it moved inside the same file, and the repair names the settling ref rather than re-anchoring at a different file. | 1 |
| 6 | "**plus** the folded skiller (taxonomy, matching, embeddings), skillpath, jobsimulation, cms and roadrunner domains" | `corpus/services/README.md:37` | There is **no `roadrunner` domain in `app`**. At `ad9f3c49`, `git ls-tree --name-only ad9f3c49 internal/ \| grep -i road` → **exit 1**, and `git ls-tree -r --name-only ad9f3c49 \| grep -i roadrunner` → **no path in the entire tree** (cross-checked with a working-tree `find`, and `git log -S 'internal/roadrunner'` returns nothing). Judge0 execution lives in the **jobsim** domain as `app/internal/jobsimulation/runner/`, constructed at `app/internal/jobsimwiring/wiring.go:123` (`jsrunner.NewRunnerManager(getenv("JUDGE0_API_KEY"), getenv("JUDGE0_BASE_URL"))`). The row contradicted its **own file** at `:20-23` (*"`roadrunner` is the eighth, and it is different: **orphaned, not merged-and-undeployed**"*) and `:42`, and the fenced `platform-migration-status.md:87` (*"**`app/internal/roadrunner/` does not exist**"*). **Second half of the same sentence, fixed in the same edit:** the row listed **5** folded domains where the file's own banner at `:11-12` names **seven** — `storage`, `messenger`, `customerio-sync` were missing. All seven `app/internal/{skiller,skillpath,jobsimulation,cms,storage,messenger,customeriosync}/` verified present at `ad9f3c49`. | 1 |

---

## Reach — graded, not asserted

| predicate | anchors BOOKED | sites FOUND | sites REPAIRED | how I searched |
|---|---|---|---|---|
| `sentinel-only-cross-process-edge` | 1 (`sentinel.md:85`) | **1** in my files | **1** | `git grep -inE 'cross-process\|only service address\|single service address\|exactly one\|AUTHORIZATION_ADDRESS\|GOTENBERG\|_RPC_ADDR\|one cross'` over all 4 files, then corpus-wide. `sentinel.md:5` and `:89` also say *"exactly one block"* / *"the only compose block given its address"* — both scoped to `AUTHORIZATION_ADDRESS` specifically, both **TRUE**, deliberately left (rule 5). |
| `org-mixin-user-count` | 1 (`security_compliance.md:67-68`) | **2** (`:67-68` false · `:135` already correct) | **1** | `git grep -inE 'OrganizationMixin\|30 (schemas )?use'` over my files + corpus-wide. `:135` already reads *"the **29** live `OrganizationMixin{}` users (a 30th is commented out at `user_resource.go:22`)"* — verified, not rewritten. |
| `unpoliced-remainder-is-reference-data` | 1 (`security_compliance.md:104-105`) | **1** in my files | **1** | `git grep -inE 'global reference data\|no org column by design\|reference data\|tenant_eid'` over my files + corpus-wide. |
| `directus-repoint-target-count` | 1 (`external_services.md:136`) | **2** (`:136` false · `:208-209` anchors rotted by the rext re-pin) | **2** | `git grep -inE 'ONE target\|only consumer\|DIRECTUS_DATA_CONSUMERS\|re-point'` over my files; ground truth from `git grep 'DIRECTUS_DATA_CONSUMERS' 09d06070` in the **pinned** rext clone (which also surfaced the dev twin). |
| `readme-folded-roadrunner-domain` | 1 (`services/README.md:37`) | **1** in my files | **1** | `git grep -i roadrunner` over all 4 files; then corpus-wide `git grep -inE 'roadrunner (domain\|domains)\|folded .{0,60}roadrunner\|roadrunner.{0,40}(folded\|merged into\|in-process\|in-app)\|app/internal/roadrunner'`. `external_services.md:173` and `sentinel.md:5`/`:85`/`:89` mention `roadrunner` only as a **deleted compose service** — correct, left alone. |

**Booked 6 → found 8 → repaired 6 + 1 prevention = 7 edits.** The one found-but-not-repaired site is
`security_compliance.md:135`, which was already correct. Honest residual: **the predicate `readme-folded-roadrunner-domain` is far wider than my partition** — 8 further sites live outside my files (below).

---

## Twins outside my files (REPORT, do not edit)

| predicate | file:line | why it is the same claim |
|---|---|---|
| `readme-folded-roadrunner-domain` | `CLAUDE.md:198` | *"**Plus the cms, jobsimulation and roadrunner domains**"* — asserts a `roadrunner` **domain** inside `app`. |
| `readme-folded-roadrunner-domain` | `CLAUDE.md:201` | *"- **roadrunner domain**: Judge0 code execution, called directly via `JUDGE0_BASE_URL`"* — names the non-existent domain outright. |
| `readme-folded-roadrunner-domain` | `corpus/architecture/architecture_overview.md:323` | *"→ cms / jobsimulation / roadrunner domains in-process"* — same predicate, inside the local-stack diagram that `:321` is the CANON-1 model for. |
| `readme-folded-roadrunner-domain` | `corpus/ops/platform_repo.md:68-70` | *"**Eight services were folded into `backend`.** `skiller`, `skillpath`, `roadrunner`, … all run in-process inside `app`"* — the strongest form: it puts roadrunner in the fold set **and** asserts in-process residency. |
| `readme-folded-roadrunner-domain` | `corpus/README.md:7` · `corpus/ops/README.md:7` · `corpus/ops/update_guide.md:16` | The same *"**Eight** services — … `roadrunner` …"* fold sentence, three more times. |
| `readme-folded-roadrunner-domain` | `corpus/services/backend.md:11` | Fold-table row `\| [roadrunner](./roadrunner.md) \| with jobsim-in-app \|` — lists roadrunner as a folded domain of `app`. |
| `sentinel-only-cross-process-edge` | `corpus/ops/platform_repo.md:73-75` | *"The only service address `docker-compose.yml` still sets is `AUTHORIZATION_ADDRESS=…`"* — the exact generalisation CANON-1 retracts. Its **second** clause (*"the single cross-process RPC edge"*) is correctly scoped; only the first is false. **Outside the audited `services/**` + `architecture/**` partition**, like `CLAUDE.md` was. |
| `unpoliced-remainder-is-reference-data` | `corpus/architecture/architecture_overview.md:360` | *"(**not** on every table — the taxonomy and other global reference data carry none by design)"* — the same characterisation of the no-`organization_id` set as reference data. Its sibling at `:364` is already correct on the mixin count, so this file is half-repaired on my two predicates. |

Note: `corpus/services/gotenberg.md:50`, `dependency_map.md:103`, `platform-migration-status.md:105`,
`service_taxonomy.md:425`, `backend.md:47-49`/`:294`, `jobsimulation.md:145-146`,
`platform-migration-status.md:93` and `CLAUDE.md:282` all already carry the CANON-1 form (other seats /
the orchestrator landed them during this pass). My `sentinel.md:85` is consistent with all of them.

---

## Noticed, not repaired

1. **The rext pin moved and it silently rotted anchors.** `stack-demo/rosetta-extensions` advanced
   `ab81527a` → **`09d06070`** (2026-08-06, a descendant). This is the class §5 rule 45 was written for,
   arriving from the other direction: iter-99/#21 and #22 both booked `external_services.md:208-211` and
   both were **rejected** because the anchors were byte-exact at `ab81527a` — and one re-pin later they are
   not. **Any unpinned rext `file:line` in the corpus is now suspect**, not just this one. I pinned the two
   in my file to `09d06070` and recorded the prior values; I did not sweep other files for the class.
2. **`platform-alignment.md:1513`'s pin survives only because I forced line-neutrality.** It is a
   `corpus/ops/**` → `corpus/architecture/**` line pin into the middle of a file three seats could touch.
   Nothing mechanical protects it. Worth a fence.
3. **`security_compliance.md`'s tenancy fence is now internally redundant.** `:67-68` (repaired) and
   `:82-90` (the historical ⚠️) now both explain the 30-vs-29 split. They agree; it is verbosity, not a
   defect, and I left `:82-90` alone because it carries the *derivation history* the fence exists to teach.
4. **`sentinel.md:5` and `:85` name different platform refs** for the same measurement (`0dab54d` vs
   `0c91421`). Both claims are true at their own named ref, so neither is a defect — but a reader diffing
   the two paragraphs will notice. Not in my anchor set; not touched.

---

## The two things I was told not to touch — verified, and reported

**1. `security_compliance.md:185` — REJECTED at iter-101 as a mis-read. Not a defect. UNTOUCHED.**
It is byte-identical to `HEAD` (proved by the `:106-281` diff above), and its line **number** is also
unchanged, because the file's length did not move.

**2. The EU AI Act block at `:226-256` — finished work, left silent. VERIFIED, and it still governs.**
The corpus asserts **no** classification, and I added none, no placeholder, and no "improvement".

- `:226` `### EU AI Act`; the three stated bullets are at **`:227`, `:228`, `:229`** — unchanged.
- The retraction blockquote runs **`:231-256`**, contiguous, every line `>`-prefixed — unchanged.
- The fence sentence is at **`:252`**: *"**Both bullets above are what is STATED, not what this corpus
  asserts** — including the consequence bullet."* It sits **inside** `:231-256`, and the only content
  between it and the bullets is the rest of that same blockquote. **It governs `:227-229`.**
  `:248-249` reinforces it from the other end — *"**Do not cite this section as evidence of a Limited-Risk
  classification** — re-derive it."*
- The whole region is **byte-identical to `HEAD`**, and `markdown_structure_guard` confirms the blockquote
  is structurally intact. `platform-alignment.md:1513`'s citation (*"`security_compliance.md:227` under
  `:231-256`"*) therefore still resolves exactly as written.

---

## Guard runs

| guard | invocation | result |
|---|---|---|
| `corpus_index_guard` | `python3 corpus_index_guard.py /Users/marco/workspace/anthropos/rosetta/corpus` | **GREEN — exit 0.** *"OK — all 84 doc(s) across 6 index-bearing directory/ies … have their directory-README index row."* |
| `markdown_structure_guard` | `python3 markdown_structure_guard.py --repo-root …/rosetta` | **GREEN — exit 0.** *"scanned 112 published file(s) … no structural damage."* Run because my `services/README.md:37` edit is inside a **markdown table cell** and my `security_compliance.md` edits are inside a **blockquote** — the two structures an in-place reflow can break. |

> ⚠️ **`corpus_index_guard` takes `CORPUS_ROOT` positionally, not `--repo-root`.** The brief's invocation
> (`--repo-root <repo>`) exits **2** with *"CORPUS_ROOT PosixPath('--repo-root') is not a directory"* — a
> **fail-loud non-zero**, so it cannot be mistaken for a pass, but it is not the GREEN the brief expects.
> The working form is `corpus_index_guard.py <repo>/corpus`. Worth fixing in the brief; the two guards in
> this table genuinely take different flags.

---

## What I could not settle, and why

**Nothing material was left unsettled.** Every claim I wrote resolves at a named ref, and I re-opened each
cited line after editing.

One judgement call is worth recording because a re-reader may land differently. `external_services.md:136`'s
*"ONE target, not two"* is false about the **declared** tuple (2 members) and true about what **matches on a
current clone** (1). I did not pick a side — the repair states both, which is what removes the in-file
contradiction with `:206` without weakening either true clause. A reader who reads *"target"* as *"service
actually re-pointed"* would have called `:136` merely imprecise; the adjudicator upheld it under the
self-contradiction rule, and the repair honours that reading while keeping the operational fact.

Two scoping limits, stated rather than smoothed:

- I did **not** verify the seven fold **wiring call sites** (`app/main.go:690`, `:721`, `:751`, `:1153`,
  `:524`, `:1471`, `:395`) that `platform-migration-status.md:87` pins at `2035f9a`. My `services/README.md`
  repair asserts only that the seven **packages exist** at `ad9f3c49`, which I measured directly. Seat 2
  owns that row.
- `sentinel` moved 2 commits / 3 files (`CHANGELOG.md`, `go.mod`, `go.sum`) to `f2c46190`, so I re-derived
  rather than inherited: `go.mod:3` = `go 1.26.0` ✓ and `Dockerfile:2` / `Dockerfile.dev:2` =
  `FROM golang:1.26-bookworm AS build` ✓ both still hold — `sentinel.md:12` is unaffected. My actual anchor
  (`:85`) is a **platform-compose** claim and is settled by `platform` `0c91421`, not by the sentinel repo
  at all; the dependency bump touches nothing it cites.
