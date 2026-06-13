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
