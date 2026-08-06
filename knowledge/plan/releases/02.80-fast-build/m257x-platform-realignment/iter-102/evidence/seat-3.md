# seat-3 report

**File owned:** `corpus/architecture/service_taxonomy.md` (only file edited).
**Anchors booked:** 6 (`:37`, `:266`→ live `:268`, `:405`→ live `:407`, `:101-102`, `:130-133`, `:109-111`).
**Sites found:** 6. **Sites repaired:** 6. Plus **2 induced-drift repairs inside my own file** (below) and
**1 induced-drift breakage in a file I do not own**, reported not edited.

> **Anchor-frame note.** The two iter-99 anchors `:266` and `:405` were measured against the pre-iter-100
> corpus. iter-100's two-line parenthetical (the very defect this seat repairs at `:130-133`) shifted them
> **+2** to live `:268` and `:407`. Same +2 that broke the archive note. I repaired the live sites.

## Ledger rows

| # | the false claim | anchor | what is true | sites repaired |
|---|---|---|---|---|
| 1 | `"    Academy -->\|GraphQL - academy subgraph\| Backend"` | `service_taxonomy.md:37` | There is no "academy" subgraph. **All three** `graphql-wundergraph` supergraph configs at `60c229f3` (`supergraph-config-prod.yaml`, `-dev`, `-compose`) declare a single `- name: backend` entry, and `schemas/` holds one file, `backend.graphqls`. Academy's types are **one SDL file inside the `backend` subgraph** — `internal/web/backend/graphql/graph/schemas/academy.graphqls`, 1 of **43** files in that dir at `app` `ad9f3c49`. Locally there is no router at all (platform `2adcf71`). Corpus already carried the correct form at `academy-backend.md:83` (*"There is no separate 'academy subgraph'"*) and this file self-contradicted at its own `Subgraphs` row (*"`backend` alone (1)"*, now `:407`). | 1 |
| 2 | `"**A GraphQL client of the platform \`app\` academy subgraph at runtime**"` | `service_taxonomy.md:268` (booked `:266`, pre-shift) | Same predicate, same evidence as row 1. Repaired to *"a GraphQL client of the platform `app` (`backend`)"* + an explicit **there-is-no-separate-academy-subgraph** clause carrying the measurement, the `academy-backend.md` twin, and the in-file `backend` alone (1) cross-check. | 1 |
| 3 | `"a second Brevo contact pusher alongside \`backend\`'s own"` **(the false half only — narrowed at consolidation).** The full sentence began *"so `make up-all` started …"*, and that first half is **TRUE and was kept**, so quoting it would have fenced a true clause. | `service_taxonomy.md:101-102` | The **first half is TRUE and kept**: `customerio-sync` was in the `all` profile until the deletion — `0dab54d:docker-compose.yml:154` `profiles: [customerio-sync, all]`. The **second half is false**: `backend`'s own in-process pusher was never on locally. `0dab54d:docker-compose.yml:56` sets `ENVIRONMENT=development` on `backend` (still `:56` at `0c91421`) → `deployedEnvironment()` returns **false** (`app/env_guards.go:37-44` @ `ad9f3c49`) → unset `CUSTOMERIO_SYNC_ENABLED` resolves `(false, nil)` not an error (`resolveSubsystemSwitch`, `env_guards.go:92-111`) → `main.go:394`'s `if customerIOSyncEnabled` never fires. **Independently re-derived for the pre-switch window**: the fold commit `app` `3e5bc33ef` (2026-08-04) gated the manager at `main.go:387` on `deployedEnvironment() && os.Getenv("BREVO_KEY") != ""` — also false locally; `3df469da8` (the switch) is its **immediate child**, so there is **no ref in the overlap window** where the in-app pusher was on with `ENVIRONMENT=development`. Exactly **one** Brevo pusher, the container. Also removed the self-contradiction with `:98-100` three lines up. | 1 |
| 4 | `"This note exists because rows \`service_taxonomy.md:137\`/\`:138\` published the flat form two rows above \`:139\`, a cell retracting exactly that predicate."` | `service_taxonomy.md:130-133` | **Repair-induced by iter-100.** Correct at `a229f8d^`; iter-100's own two-line parenthetical pushed the table +2 and left the numbers unmoved, so the note sent readers to **Chronos** (`:137`, no archive assertion) and **Intelligence** (`:138`, none) and called **Skiller**'s flat `ARCHIVED 2026-07-01` (`:139`) the retraction of itself. Repaired to name the rows **by service** as well as by number, at the post-edit indices: **Skiller `:157`**, **Skillpath `:158`** (both flat form), immediately above **Jobsimulation `:159`** (the *"report both, assert neither"* cell). | 1 |
| 5 | `"**Base services (no profile, always on with any \`make up\`)**:"` followed by **two** bullets (PostgreSQL, Redis) | `service_taxonomy.md:109-111` | The floor is **three**: `{postgresql, redis, sentinel}`. Measured at platform `0c91421` — `docker-compose.yml` declares five services and exactly **four** carry a `profiles:` key (`backend` `:110`, `studio-desk` `:141`, `next-web-app` `:168`, `gotenberg` `:183`); `sentinel` (`:5`) carries none, and `common.yml`'s `postgresql` (`:2`) / `redis` (`:24`) carry none. Under the two-bullet form the file's own arithmetic did not close (backend + gotenberg + 2 = four vs the **five** it states in three other places). Every other floor statement in the file — the *Services* paragraph, the Sentinel row's *"(always on — declares no `profiles:` key)"*, the *Profiles* table, the Summary Table — already said three; only this list was wrong. Sentinel added with its port + compose anchor and an explicit note that it is a Tier-1 service **and** a floor member. | 1 |
| 6 | `"there is **exactly one** cross-process edge left, \`backend → sentinel\`. Compose sets a single service address, \`AUTHORIZATION_ADDRESS=http://sentinel:8087\` (\`docker-compose.yml:48\`), and **zero \`*_RPC_ADDR\` variables**."` | `service_taxonomy.md:407` (booked `:405`, pre-shift) | **CANON-1 applied verbatim in substance** (`canonical-repairs.md` §CANON-1), adapted only in grammar. Re-verified every fact myself at `0c91421` / `ad9f3c49`: `AUTHORIZATION_ADDRESS` `:48` ✓, `GOTENBERG_URL=http://gotenberg:3200` `:57` ✓, `JUDGE0_BASE_URL` `:59` ✓, `gotenberg` `profiles: [core, backend, all]` `:183` ✓, plain-HTTP consumption at `app/internal/converter/gotenberg.go:31` (`http.NewRequestWithContext(ctx, "POST", gotenbergURL+"/forms/libreoffice/convert", &body)`) ✓, **zero `*_RPC_ADDR` in compose** ✓ — that clause is TRUE and was preserved unweakened. The generalisation to *"only cross-process edge / single service address"* is what was repaired, with the model form cited. | 1 |

## Reach — graded, not asserted

| predicate | anchors BOOKED | sites FOUND | sites REPAIRED | how I searched |
|---|---|---|---|---|
| `academy-subgraph-exists` | 2 (`:37`, `:266`) | 2 (`:37`, `:268`) | 2 | `grep -n -i "academy subgraph\|subgraph.*academy"` over my file; then `git grep -n -i "academy subgraph"` over `corpus/**` + `CLAUDE.md` for the tree-wide twin set (12 sites outside my files, listed below). Ground truth by `git show 60c229f3:supergraph-config-{prod,dev,compose}.yaml` + `git ls-tree` of the schemas dir at `ad9f3c49`. |
| `two-brevo-pushers-up-all` | 1 (`:101-102`) | 1 | 1 | `grep -n -i "brevo\|pusher\|up-all"` over my file → 2 hits, one of them the literal `make up-all` shell line (not a claim). Ground truth by `git show 0dab54d:docker-compose.yml` (profiles + `ENVIRONMENT`), `git show ad9f3c49:env_guards.go`, `git grep -n CUSTOMERIO_SYNC_ENABLED ad9f3c49 -- '*.go'`, plus a history walk (`log --diff-filter=A -- internal/customeriosync/`, `merge-base --is-ancestor`) to close the pre-switch window. |
| `archive-note-row-anchors` | 1 (`:130-133`) | 1 | 1 | Direct. Then the mandatory post-edit row-resolution proof (below). |
| `base-services-floor-cardinality` | 1 (`:109-111`) | 1 | 1 | `grep -n -i "floor\|base service\|always on\|no profiles"` over my file → 11 hits; **10 were already on the correct side (three)**, only `:109-111` said two. Ground truth by `git show 0c91421:docker-compose.yml \| grep -n '^  [a-z-]*:\|profiles:'` + `common.yml`. |
| `single-service-address-gotenberg` | 1 (`:405`) | 1 (`:407`) | 1 | `grep -n -i "only.*edge\|one.*hop\|only.*address\|one cross\|AUTHORIZATION_ADDRESS\|RPC_ADDR"` over my file → 3 hits; the other two (`:142` CMS row, `:171` content-vs-runtime note) correctly say the RPC hops are **gone**, not that sentinel is the only edge. |

**No predicate in my set turned out to be wider than its booking inside my file.** That is the honest
result and it is worth stating plainly against the brief's 3× expectation: `service_taxonomy.md` was
already uniformly on the correct side for the floor cardinality (10 of 11 sites) and the RPC-edge
scoping (2 of 3 sites). The width for `academy-subgraph-exists` is real but lives **outside** my file —
12 sites across 6 other documents, enumerated below.

## Twins outside my files (REPORT, do not edit)

| predicate | file:line | why it is the same claim |
|---|---|---|
| `academy-subgraph-exists` | `CLAUDE.md:265` | *"read from the platform academy subgraph over GraphQL"* — verbatim the same construct. Orchestrator-owned. |
| `academy-subgraph-exists` | `corpus/architecture/architecture_overview.md:40` | *"read from the platform academy subgraph over GraphQL"* |
| `academy-subgraph-exists` | `corpus/services/ant-academy.md:63` | *"a backend-authoritative read/WRITE GraphQL client of the platform `app` academy subgraph"* |
| `academy-subgraph-exists` | `corpus/services/ant-academy.md:95` | mermaid edge `Academy -->\|catalog: GraphQL academy subgraph\| App` — the exact twin of my `:37` |
| `academy-subgraph-exists` | `corpus/services/ant-academy.md:106` | *"**academy subgraph** as a GraphQL *client* … "no subgraph" ≠ "no GraphQL""* — **partially self-aware**; still names an academy subgraph |
| `academy-subgraph-exists` | `corpus/services/ant-academy.md:118` | *"queried from the **platform academy subgraph**"* |
| `academy-subgraph-exists` | `corpus/services/ant-academy.md:433` | *"**academy subgraph** (`app internal/academy`) over GraphQL"* |
| `academy-subgraph-exists` | `corpus/ops/run_guide.md:232` | *"reads its course catalog from the platform academy subgraph over GraphQL"* |
| `academy-subgraph-exists` | `corpus/ops/demo/content-stories-routes.md:384`, `:400`, `:457` | three sites, all *"the academy subgraph"* |
| `academy-subgraph-exists` | `corpus/ops/demo/frontend-tier.md:448`, `:477` | two sites; `:477` says *"compose the academy subgraph into the demo router"* |
| — (the MODEL, do not edit) | `corpus/services/academy-backend.md:83` | already correct: *"**There is no separate "academy subgraph"** — these types live in the `app`/`backend` federation subgraph — **the only subgraph left**."* This is the wording every site above should be brought to. |
| — (the MODEL, do not edit) | `corpus/architecture/architecture_overview.md:321` | CANON-1's model form, verified present and unchanged. |

## `service_taxonomy.md:52` — the cross-seat interaction

**`:52` did NOT move.** Proven two ways:

1. `git diff -U0` hunk headers: the first hunk that changes a line **count** is `@@ -101,2 +101,12 @@`.
   Everything at lines 1–100 is at an identical line number before and after. My only edit above line 101
   is `@@ -37 +37 @@` — a **one-line-for-one-line** mermaid label swap, zero delta.
2. `diff <(git show HEAD:…| sed -n '52p') <(sed -n '52p' …)` → **IDENTICAL**. Same for `:62`.

Seat 4 is safe. I also observed that **seat 4 has already landed its repair**: `hiring.md:38-46` now cites
`service_taxonomy.md`'s Tier-1 **Database** bullet at **`:62`** (not `:52`) and names it by description as
well as by number. `:62` is also unmoved and its content still matches seat 4's quote verbatim
(*"one schema, `public`, owned by `app`, which is the only repo with migrations"* / *"legacy husks"*).

## Post-edit row-resolution proof (rule 4 — the induction check)

Every row reference in the repaired archive note resolves to the row it names. Re-read the **entire**
table plus the note after the last edit:

| index | resolves to | archive assertion present? | note's claim about it |
|---|---|---|---|
| `:155` | **Chronos** | **none** (`grep -ci archiv` → 0) | not referenced (this is the row the broken note pointed at) |
| `:156` | **Intelligence** | **none** (0) | not referenced (ditto) |
| `:157` | **Skiller** | `ARCHIVED 2026-07-01` — **flat form** | named as a flat-form row ✔ |
| `:158` | **Skillpath** | `ARCHIVED 2026-07-31` — **flat form** | named as a flat-form row ✔ |
| `:159` | **Jobsimulation** | *"report both, assert neither"* — **the retraction** | named as the retracting cell, immediately below ✔ |
| `:160` | **CMS** | (own report-both clause; not referenced) | — |
| `:161` | **Roadrunner** | none; not referenced | — |

Wording changed from *"two rows above"* to *"immediately above"* because 157/158 is a **pair** and only
the first is literally two rows up — the original phrasing would have been imprecise even with correct
numbers. Rows are now named **by service** as well as numbered, so a future shift degrades to a stale
number beside a resolvable name rather than to a confident pointer at the wrong row.

The note block is **exactly 4 lines**, as before, deliberately: any delta there would have re-shifted the
very table it indexes.

## Induced and corrected

Three drift defects were created by my own edits' `+18` line shift. Two were **inside my file and I fixed
them**; one is **outside my file and is reported, not edited**. Plus one inherited from the canonical sheet.

| # | site | what broke | disposition |
|---|---|---|---|
| I1 | `service_taxonomy.md:394` (was `:376`) | *"see the archive-state note at `:124`"* — after my +18, `:124` landed inside the new Base-services bullet list. | **FIXED in-file**: now *"the archive-state note above the *Archived / merged* table, `:142`"* — named **and** numbered, same hardening as the note itself. |
| I2 | `service_taxonomy.md:397` (was `:379`) | *"Consistent with `:67-68` above"* — checked, **not broken**: `:67-68` is above my first count-changing hunk and still reads *"There is no `graphql` profile"*. | verified, no change needed |
| I3 | **`corpus/architecture/external_services.md:144`** — **NOT MY FILE** | cites `` [`service_taxonomy.md:332-339`] ``. My +18 moved that content to **`:350-357`**; `:332-339` now resolves to the Clerk *Environment Variables* / *Used By* block — a different subject entirely. Verified by diffing `git show HEAD:…\|sed -n '332,339p'` against the current `350,357p`: **byte-identical**, offset exactly +18. | **REPORTED, not edited.** Whoever owns `external_services.md` should change `:332-339` → `:350-357` (or, better, name the passage). |
| I4 | `service_taxonomy.md:425` | **Inherited from `canonical-repairs.md` §CANON-1**, not derived by me: the canonical replacement ended *"Judge0 directly via `JUDGE0_BASE_URL` (`:59`)"*, and the bare `:59` resolves against the most recently named file — `app/internal/converter/gotenberg.go`, which is **53 lines**. Out of range. Flagged mid-task by the orchestrator. | **FIXED**: `` (`docker-compose.yml:59`) ``, and `` (`:183`) `` made explicit in the same clause for the same reason. `anchor_construct_guard.py` re-run: **my site is gone from RED**. |

I4 is worth its own line in the induction-rate measurement, and I state it as the orchestrator asked:
**centralising a wording centralises its defects.** The whole point of §CANON-1 was to stop five seats
producing five differently-worded fixes — but it also removed the five independent re-derivations that
would have caught a bad anchor, and one error propagated to four sites at once. My seat re-derived every
*fact* in CANON-1 against the clones (and confirmed all of them), but I applied its *citation form*
verbatim, which is exactly the gap. **A canonical sheet should carry the facts centrally and still be
anchor-checked per site.**

## Noticed, not repaired

1. **`anchor_construct_guard.py` still RED on one site — `corpus/services/sentinel.md:85`**, same I4 class,
   same canonical sheet: `cites app/internal/converter/gotenberg.go:59 … file has 53 line(s)`. That is
   **seat 10's** file, not seat 1's — the orchestrator's message named `backend.md:48`/`:282` (seat 1,
   now clear). Flagging so it is not lost.
2. My file's `:147` cites `platform-migration-status.md:89` for the Jobsimulation archive retraction.
   I verified it **currently resolves** (the `jobsimulation` row at `:89` carries *"Repo archive state —
   REPORT BOTH, ASSERT NEITHER"* + the four post-dated 2026-08-04 commits). But
   `platform-migration-status.md` is **seat 2's file and is being edited concurrently**, so this citation
   is live drift risk in the same way `external_services.md:144` was for me. Not repointing it — that
   would be guessing at another seat's final line numbers.
3. `corpus/services/ant-academy.md:106` already half-knows the academy-subgraph problem (*"no subgraph"
   ≠ "no GraphQL"*) while still naming an academy subgraph three lines of prose later. Whoever repairs
   that file should note the passage is arguing *against itself*, not just mis-stating a fact.

## What I could not settle, and why

Nothing in my assignment was left unsettled — all 6 booked anchors were measured at named refs and
repaired. Two boundaries I deliberately did not cross:

- **The 12 out-of-file `academy subgraph` twins.** Real, enumerated, and outside my partition. I did not
  edit them (brief §"YOUR FILES") and I did not grade them as new findings (rule 6).
- **Archive state itself** (`gh api … --jq .archived`) remains unmeasurable from this host — `gh` absent,
  repos private. I preserved the file's existing, correct fence on this rather than trying to strengthen
  it; the whole point of the `:130-133` note is that the dates are *asserted on*, never *is*.
