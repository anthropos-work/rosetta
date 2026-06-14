# M25 — Decisions

_Implementation decisions with rationale, numbered `M25-D1`, `M25-D2`, … . Empty at scaffold; filled during build._

## M25-D1 — `GOTOOLCHAIN=local` on the offline clerkenstein builds (bug fixed inline)

**Surfaced:** DB-1 (`/demo-up`) aborted **deterministically** at the 7-build wait-gate, before compose-up
(0 containers), even though every service image built fine.

**Root cause:** the fake FAPI/BAPI host cross-compile (`up-injected.sh:191`) and the `cmd/jwtkey` `go run`
(`:232`) carry `GOPROXY=off GOSUMDB=off` (deliberate offline build). M24 pinned clerkenstein's `go.mod` to
`toolchain go1.25.11`. When the host Go must *switch* to that toolchain, it tries to fetch+verify the
toolchain **module** — which both offline flags forbid — so `go build` hard-fails with "verifying module:
checksum database disabled by `GOSUMDB=off`". Under the gate's `wait $pid || fail=1`, that one non-zero PID
aborts the whole bring-up. The `jwtkey` `go run` hit the same wall and, being `|| true`, silently returned
empty (which would have broken studio-desk's RS256 token verify on the UI path).

**Fix (Fate-1, in the owning module `rosetta-extensions/demo-stack`):** add `GOTOOLCHAIN=local` to both
offline Go invocations. The host toolchain already satisfies the 1.25.11 floor, so using it (and skipping the
fetch+verify) is correct and keeps the build offline. Committed authoring `35180c0`; tagged `prop-room-m25`;
re-pinned the `stack-demo` consumption clone. This is precisely the field-fix-tail class M25 exists to pre-pay
— a clean-room M24 unit pass couldn't see it because the toolchain-switch only triggers in the offline
cross-compile path the live bring-up exercises.

## M25-D5 — The directus_files closure must EXCLUDE the tenant-referenced file set (firewall fix; the field-bake's central safety find)

**Surfaced (the resume run):** with the operator-sanctioned prod-read DSN, the FIRST real
structure-bearing directus capture **FAILED the firewall** — correctly, fail-closed, **zero data
written**: `firewall: PUBLIC-ONLY VIOLATION — directus.directus_files captured 9065 tenant-scoped
rows — TENANT DATA LEAK`. This is exactly the field-fix tail M25 exists to pre-pay: a real safety
bug in the M23 directus_files capture that only the live capture against real prod data could
surface (the clean-room M23 unit tests used synthetic fixtures that never exercised the
resource→file / shared-file shape).

**Root cause (two compounding defects, both in the M23 `directus_files` referenced-subset capture):**
1. **The closure OVER-captured.** `directus_files` has **no tenancy column** — a file's tenancy is
   *inferred* from what references it. The M23 closure (`ReferencedFilesFilter`) was
   `(public sims/paths/roles/sequences refs) OR (ALL resource.image/file)`. But `resource` is a
   **pure-reference** table with **no tenancy column**, so its clause pulled in **every** resource
   file — including ones a **tenant** simulation also references (8 in prod). And **150** files are
   referenced by **both** a public published sim **and** a tenant sim; the public-only half kept all
   of them. Net: the closure captured **1417** files, **158** of them tenant-referenced.
2. **The leak probe used a false premise.** For a referenced-subset (strict-subset) table the M23
   post-capture probe counted rows **OUTSIDE** the closure (`WHERE NOT (closure)`) and called them
   leaks. But the uncaptured remainder of a 1.3k-of-10.5k subset is **not** a leak — that probe
   reported the **entire 9065-file remainder** as the "tenant leak" (the comment claimed "0 by
   construction", which only holds if the closure captures the whole table).

**Fix (Fate-1, in the owning module `rosetta-extensions/stack-snapshot`, tag `prop-room-m25` moved to
include it):** align the closure with the firewall's public-only *definition* — capture a file only
if public-referenced **AND NOT** tenant-referenced.
- `ReferencedFilesFilter` now appends `AND NOT (<TenantReferencedFilesFilter>)`. The new
  `TenantReferencedFilesFilter()` is the boolean tenant-reference predicate (negated public roots +
  tenant-parented roles/sequences; `resource` contributes nothing — no tenancy column, so a resource
  file is tenant *only* by being referenced by a tenant sim, caught via the sims/roles/sequences
  clauses). Captured `directus_files` drops **1417 → 1257** (the 158 excluded). **Safety-first:**
  under-capturing a shared public image is fine (its content still serves, sans that one asset);
  leaking a tenant-referenced file is not.
- The referenced-subset **leak probe** now counts **captured** rows that **are** tenant-referenced
  (`WHERE (closure) AND (tenantFilter)`) — 0 by construction, and a genuine defense-in-depth catch if
  the closure's exclusion ever stops constraining.
- The firewall **`AssertPlan`** gate now **requires** a referenced-subset table to declare its
  tenant-reference filter **and** requires the closure to embed the exclusion — so no future
  referenced-subset table can ship without a tenant-leak definition. The filter is **never weakened**;
  the fix is in the FILTER, exactly as the safety contract demands.

**Proof the fix is correct:** the **REAL prod capture now PASSES the firewall** — `captured "directus"
@ 6cd35278: 10 tables, 11480 rows, public-only=true, source=primary-read`, `directus_files=1257`, plus
the M21 `_structure.sql` artifact (62 statements). The rows-only cache is **upgraded in-place to
structure-bearing**, unblocking DB-1/DB-2/DB-4. Regression tests added (the resource→tenant
over-capture reproduced + proven excluded; the require-tenant-filter / closure-must-exclude firewall
gates; the new probe shape). `go test ./...` green. (This **supersedes the M25-D3 blocker** below: the
blocker was "needs a sanctioned prod read"; the operator sanctioned it, the read surfaced the real bug
M25-D5, and the fix + capture cleared it.)

## M25-D6 — NULL the `directus_collections` self-FK `group` in the serve-row registration (2nd field bug)

**Surfaced (the live DB-1 set-dress, only by the structure APPLY):** with the firewall-clean
structure-bearing snapshot in the cache, `/demo-up` got all the way to the per-stack Directus provision —
then the structure apply **failed**:
`apply structure (62 statements): ERROR: insert or update on table "directus_collections" violates foreign
key constraint "directus_collections_group_foreign" (SQLSTATE 23503)` → directus replay **rc=1**, the
local Directus came up serving **nothing** (0 collections, 0 content tables — the script is one
transaction, so it rolled back). Taxonomy replayed fine (329,859). This is a **second** real field bug the
clean-room M21 unit tests (fake `RowStringQuerier`, no live apply) couldn't see.

**Root cause (M21 serve-rows):** `directus_collections` carries a **self-referential FK**
(`directus_collections_group_foreign`): a collection's `group` references a **parent group collection** —
in Directus an **admin-UI data-model FOLDER** (`simulations`, `job_simulations`, `paths`,
`learning_resources`). Those group collections have **no table** and are **not** in `servedCollections`, so
the serve-row registration faithfully emitting each content collection's captured `group` value dangles the
self-FK.

**Fix (Fate-1, `stack-snapshot/directus/structure.go`):** the collection-registration render
(`serveCollectionsRowsSQL`) now **NULLs the `group` column**
(`CASE WHEN j.key = 'group' THEN 'NULL' ELSE quote_nullable(j.value) END`). The `group` is **admin-UI
organization only** — it nests collections in the data-model sidebar and has **zero** effect on serving
content over REST/GraphQL. Nulling it removes the dangling self-FK while the collections register + serve
identically. +regression test (`TestServeCollectionsRowsSQL_NullsGroupColumn`) + a guard that no served
collection is itself a group collection; render validated vs prod (20/20 columns intact, group → NULL).
Re-capture regenerates the structure artifact with `', NULL, 'open'` and still passes the firewall.

## M25-D7 — NULL `directus_files.folder` (dangling FK to uncaptured `directus_folders`) — 3rd field bug

**Surfaced (the live DB-1, ONLY after M25-D6 let the structure apply succeed and the directus ROW replay
run):** the `directus_files` COPY-load failed:
`copy into directus.directus_files: ERROR: insert or update on table "directus_files" violates foreign key
constraint "directus_files_folder_foreign" (SQLSTATE 23503)` → directus replay **rc=1**.

**Root cause:** `directus_files.folder` is an FK to `directus_folders` (the file-library admin-UI folder a
file lives in). `directus_folders` (11 rows, with its **own** self-FK `parent`) is **not** in the captured
surface, so the captured `folder` uuids (10,031 of 10,480 files carry one) **dangle** on replay. Same class
as the M25-D6 `directus_collections.group` self-FK: a dangling FK to an **admin-UI organizational table
outside the surface**.

**Fix (Fate-1, `stack-snapshot`):** a **general** `TableSpec.NullColumns` mechanism — columns rendered as
`NULL AS <col>` in the capture COPY SELECT instead of the source value (the column **stays** in position,
so the manifest/load shape is unchanged; only the captured VALUE is NULL). `directus_files` declares
`NullColumns: ["folder"]`. The folder is **admin-UI organization only** — zero effect on serving the asset
(which resolves by `id → storage/filename_disk`). Capturing `directus_folders` instead was rejected (it
carries a self-FK + would need its own structural-metadata admissibility; nulling loses nothing for
serving). Threaded through `CopyPublic` (signature gains `nullCols`; `buildPublicSelect` renders `NULL AS`
+ validates the null-col is in the column list); +tests. Re-capture: `folder` is **all NULL** (0 non-null
in the .copy), firewall still passes.

**The three FK fixes together (M25-D5/D6/D7)** are one theme the live runs taught: a real prod capture +
replay exercises **referential integrity** that synthetic unit fixtures never do — a tenant-leak via an
inferred-tenancy file, and two dangling FKs to admin-UI tables outside the surface. All three are fixed in
the FILTER / render (never by weakening the firewall), and the structure-bearing capture now **passes the
firewall and replays cleanly**.

## M25-D3 — BLOCKER: the local-Directus content serve needs a structure-bearing capture (operator prod read)

**This is the milestone's central field finding** — and a genuine **operator-decision blocker**.

**What the live runs showed (DB-1, reproduced on DB-3):** `/demo-up` (local-content default-on) brings the
stack fully up — taxonomy replayed real (329,859 rows), the per-stack Directus **container boots + is healthy
on the offset port (18055)**, cms's data plane is wired LOCAL (`DIRECTUS_BASE_ADDR=http://directus:8055`) and
the asset plane stays PROD (`DIRECTUS_PUBLIC_BASE_ADDR=https://content.anthropos.work`). **But the local
Directus serves NO content** (0 collections, 0 content tables): the directus **replay skips with rc=5
(cache-miss)** because the **cached directus snapshot is rows-only and predates the M21 structure extension**
(`stacksnap status`: digest `6cd35278edbc`, 9 tables, 10128 rows, captured 2026-06-10, **no Structure
artifact in the manifest**). With no captured structure to auto-provision onto the bootstrapped-gap schema,
the schema digest never converges (`b4cb55bcee08` ≠ cached) → no content tables → nothing to serve.

**Why it can't be unblocked autonomously:** the fix is a **structure-bearing directus re-capture** (the M21
`CaptureStructure` path), which reads the content-model DDL + serve-rows from a **real prod-origin source**
over `--dsn`. The capture-source policy admits **only** `dump-ingest` (a restored staging `pg_dump`) or
`primary-read` (a throttled prod read over Tailscale) — both **privileged, prod-touching, operator-confirmed**
reads. On this box: **no staging dump on disk**, **no `~/.pgpass` / Tailscale prod-read DSN configured**, and
policy (`snapshot-cold-start.md`, `safety.md` §1.4) **mandates operator confirmation** before any prod read. A
self-referential "capture from a locally-populated directus" is **rejected** — it would fabricate manifest
provenance (`source=primary-read`) for data that never came from prod, exactly the fabrication
`audit-kb-fidelity` forbids. So DB-1/DB-2 (content serve) and DB-4 (the capture itself) are **genuinely gated
on operator-provided prod access** — this is the milestone's prerequisite (overview DB-4 assumes "a restored
dump"), not a code defect.

**What IS proven without it:** the entire bring-up machinery + the data-plane/asset-plane WIRING + the
Directus boot path are sound (a manual restart booted the local directus healthy on 18055). The moment a
structure-bearing snapshot exists, the replay converges → restart fires → directus serves. Code is ready; the
**cache is the gap**.

**Resolution path (for the operator):** run, once, with a sanctioned source —
`stacksnap capture --surface directus --dsn <restored-staging-dump | prod-read-DSN>` (and `--surface taxonomy`
to refresh) — then re-run `/demo-up`; DB-1/DB-2/DB-4 should go green with no further code change. (Surfaced to
the user as the milestone's open blocker.)

## M25-D4 — N=0 stays prod-read (DB-2's testable half — VERIFIED in code)

DB-2's "confirm N=0 stays on the prod-read path untouched" is code-verified: `dev-setdress.sh` resolves
`LOCAL_CONTENT=0` for `dev` stack-type (opt-in via `--local-content`), AND **hard-refuses to set-dress N=0
without `--force`** (one of two independent N=0 guards; `stackseed --reset` is the other). So the main
`anthropos` dev stack structurally never gets a local Directus from the auto flow — it stays on the prod-read
default, untouched. GREEN by verification (the guard is the M15 safety-contract invariant).

## M25-D2 — Docker VM ceiling on a 16 GB host (resource finding)

The documented demo UI-tier prereq is a **12 GB Docker VM / 3 GB swap**. On this **16 GB host**, allocating
12 GB to the VM **failed to boot** (VM unreachable: `no route to host 192.168.65.7:2376`; `context deadline
exceeded`) — macOS + Docker Desktop overhead leaves no room. Backed off to **10 GB / 2 GB swap** (boots
reliably, ~9.7 GiB usable). Consequence for the field-bake: a **full UI-tier** demo (~10-12 GB runtime + a
~3.7 GB next-web build spike) does not fit co-resident; DB-1/DB-2 run **backend-only** (`DEMO_NO_UI=1` /
no `--local-content` UI), with the browser-observable behavior proven by HTTP/SQL probes against the stack's
own cms + per-stack Directus (the data plane) — exactly the surface a browser calls. The local-Directus data
plane + asset-plane-stays-prod + verify net are all exercised; only the heavyweight next-web/studio-desk
**rendering** tier is out of budget on this box. See journal for the doc/preflight follow-up.

## M25-D8 — Full-UI Playwright render: environmental backlog, NOT a v1.5 deliverable (close re-fate)

**Context (close Phase 1b deferral re-audit, `audit-deferrals/deferral-audit-2026-06-13-m25-close.md`):** the
field-bake proved the core observable behavior — the local Directus **serves** the catalog (data plane local,
asset plane prod) — by curl against cms + the per-stack Directus (the exact surface a browser calls), on
`--no-ui` stacks. The full-UI **rendered-page** Playwright screenshot didn't run: the ~10 GiB practical VM
ceiling on this 16 GB host (M25-D2) can't co-host the full UI tier (next-web build spikes ~3.7 GB) with a
backend stack.

**Fate (re-audit):** **DROP as a v1.5 tooling deliverable / KEEP as environmental backlog.** This is a
**host-budget** constraint, not a tooling defect — landing it requires a bigger box, not a v1.5 code change.
The done-bar "the browser shows content served by the local Directus" is satisfied at the **behavior** level by
the data-plane curl proof (DB-1/DB-2 evidence); the screenshot would add presentation confirmation, not
behavioral confirmation. No cross-release tooling work is owed. Single deferral (first appearance), not a
repeat, not aged. Recorded so it isn't silently lost; no roadmap mutation needed (no milestone owes UI-render
code).

## M25-D9 — dev-2 taxonomy replay `rc=4`: tracked dev migrate-ordering follow-up (close re-fate)

**Context (same re-audit):** on DB-2 (`/dev-up 2 --local-content`) the **taxonomy** replay returned `rc=4`
(target schema empty) — a pre-existing **dev-stack migrate-ordering** nuance on the opt-in dev-2 stack. It is
**non-fatal and unrelated to the content-serve path**: DB-2's directus content-serve (the core done-bar) is
GREEN — the directus replay exits 0 and serves real published rows on the dev offset port (`db2-serve-evidence.txt`).

**Fate (re-audit):** **KEEP (tracked tooling-debt follow-up — dev migrate-ordering).** Outside the field-bake's
content-serve charter; diagnosing the taxonomy-surface migrate-ordering on dev-N is dev-stack tooling debt, not
a v1.5 content-release deliverable, and the stack comes up and serves regardless. Single deferral (first
appearance), not a repeat, not aged. Surfaced so it isn't lost.
