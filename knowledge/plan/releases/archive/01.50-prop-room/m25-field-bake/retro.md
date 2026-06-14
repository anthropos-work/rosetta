# M25 — Retro

## Summary
M25 — the **final** milestone of v1.5 "prop room" — was a **live field-bake**: its done-bars were real stack runs
on the actual 16 GB box, not code sections. It proved the whole release by **observation** — fresh `/demo-up` and
`/dev-up 2 --local-content` bring the stack up and the **local Directus observably SERVES** the captured public
catalog (curl-proven: `/server/health` 200, `/items/simulations` real published rows, asset URL →
`content.anthropos.work`; data plane local + asset plane prod), N=0 stays prod-read, re-runs converge, the
cold-start capture is exercised, teardown reclaims the container + port. All **5 done-bars GREEN**. It earned its
keep loudly: the operator-sanctioned prod read filled the cache structure-bearing, and the live runs surfaced **4
real release bugs** that clean-room unit passes (synthetic fixtures, no live apply) structurally could not see —
all fixed Fate-1 in their owning ext module and re-proven by re-run. The rosetta-side surface is markdown/text
(the field-bake log + 2 curl-evidence files + 2 skill arg-hint fixes); the code fixes live in `rosetta-extensions`
(tag `prop-room-m25` @ `1a2fd91`).

## Incidents This Cycle
- **4 field bugs surfaced live + fixed Fate-1 — 0 regressions, 0 flakes.** This is the field-bake working as
  designed (pre-paying the field-fix tail prior releases shipped after the fact):
  - **P1 — the `directus_files` tenant-data LEAK (M25-D5), the headline.** The first real structure-bearing capture
    OVER-captured 158 tenant-referenced files; the **tenant-data firewall caught it FAIL-CLOSED — zero data
    written**. Root cause: `directus_files` has no tenancy column (tenancy is *inferred* from what references it),
    and the M23 closure's `resource`-clause pulled every resource file incl. tenant-shared ones + 150 dual-referenced
    files, while the leak probe used a false "uncaptured-remainder = leak" premise. Fixed **in the FILTER** (closure
    appends `AND NOT (TenantReferencedFilesFilter)`; the probe counts captured-AND-tenant rows; the firewall
    `AssertPlan` now REQUIRES referenced-subset tables to declare + embed their tenant-exclusion) — **the firewall
    was never weakened**. The real prod capture then PASSED (`public-only=true`, `directus_files` 1417→1257).
  - **P2 — `directus_collections.group` self-FK (M25-D6, apply-side SQLSTATE 23503).** The serve-row registration
    emitted each content collection's captured `group` (an admin-UI folder collection outside the served set) →
    dangling self-FK → structure apply rolled back, Directus served nothing. Fixed: the render NULLs `group`
    (admin-UI organization only, zero serve effect).
  - **P2 — `directus_files.folder`/`uploaded_by`/`modified_by` FKs (M25-D7, replay-side 23503).** FKs to uncaptured
    `directus_folders`/`directus_users`. Fixed via a general `TableSpec.NullColumns` mechanism. The **complete
    enumeration was proven** (the per-stack content tables ship PRIMARY KEYS only — 0 content-table FKs — so the
    only live FK constraints a replay can violate are the bootstrap-created `directus_files` ones), so no
    whack-a-mole.
  - **P2 — offline-build `GOTOOLCHAIN` regression (M25-D1).** M24's Go pin (`toolchain go1.25.11`) broke the
    offline clerkenstein cross-compile (`GOPROXY=off GOSUMDB=off` can't fetch+verify the toolchain module) →
    `/demo-up` aborted at the build gate. Fixed with `GOTOOLCHAIN=local`. Precisely the field-fix-tail class M25
    exists to catch — a clean-room M24 unit pass couldn't see it (the toolchain-switch only triggers in the offline
    cross-compile path the live bring-up exercises).
- **Close was clean: 3 findings, 0 fixable bugs.** 1 scope checkbox nit (DB-4 box `[ ]` while its text read
  EXERCISED+GREEN — flipped), 1 docs (the 16 GB-host VM ceiling field note), 1 decision-blend + 5 archive. The ext
  harden pass found 0 bugs (the build's fix code held under deepening).

## What Went Well
- **The firewall did exactly its job — fail-closed on a real leak.** The single most important outcome: a genuine
  tenant-data-leak hit the firewall and **zero tenant data was written**. The safety contract (read-only,
  public-only, fail-closed) held against real prod data, and the fix landed in the FILTER without ever weakening
  the firewall — it was *strengthened* (the new `AssertPlan` referenced-subset requirement means no future
  referenced-subset table can ship without a tenant-leak definition).
- **The field-bake earned its name.** 4 bugs that 4 prior milestones' green unit suites never surfaced — referential
  integrity (a tenant-leak via inferred-tenancy + two dangling FKs to admin-UI tables) and an offline-build
  toolchain regression — only a real prod capture + live replay exercises these. The release pre-paid its field-fix
  tail instead of shipping it.
- **The complete-enumeration discipline avoided whack-a-mole.** Rather than null FK columns one failure at a time,
  the prod read established the structural truth (content tables = PKs only, 0 FKs) that proves
  `NullColumns: [folder, uploaded_by, modified_by]` is the *complete* set — a closed argument, not a patch.
- **The observable-behavior gate gave honest evidence.** The done-bars were curl outputs against the exact surface
  a browser calls (cms + the per-stack Directus), captured to `db1/db2-serve-evidence.txt` — "it serves" is proven,
  not asserted.

## What Didn't
- **The box couldn't host the full UI tier — a planned-prereq that didn't fit reality (M25-D2).** The documented
  12 GB Docker VM *fails to boot* on this 16 GB host (~10 GB practical ceiling); the full UI render (next-web build
  spikes ~3.7 GB) can't be co-resident with a backend stack. DB-1/DB-2 ran `--no-ui`, proving the core behavior at
  the data-plane level. The full-UI Playwright render is dropped as a v1.5 deliverable (host-budget, not a tooling
  gap) and the corpus now warns about the 16 GB-host ceiling.
- **A blocker held mid-milestone until an operator decision (M25-D3).** The local Directus served 0 content because
  the box's cache was rows-only (pre-M21, no `_structure.sql`). Unblocking required an operator-sanctioned prod read
  — correctly NOT autonomous (a self-referential "capture from a locally-populated Directus" would have fabricated
  manifest provenance, exactly what `audit-kb-fidelity` forbids). The operator sanctioned the `marco_read`
  primary-read DSN; the read surfaced M25-D5 and the fix cleared the blocker.
- **A pre-existing dev migrate-ordering nuance surfaced on dev-2 (M25-D9).** The taxonomy replay returned `rc=4`
  (target schema empty) on the opt-in dev-2 stack — non-fatal, orthogonal to the content-serve path (the directus
  replay exits 0 and serves). Tracked as dev-stack tooling-debt follow-up, not a v1.5 content deliverable.

## Carried Forward
- **M25-DEF-01** (full-UI Playwright render) → **DROPPED as a v1.5 deliverable / environmental backlog** (M25-D8) —
  a host-budget constraint; the behavior is proven by the data-plane curl evidence. Needs a bigger box, not code.
- **M25-DEF-02** (dev-2 taxonomy `rc=4`) → tracked **dev migrate-ordering** follow-up (M25-D9) — non-fatal,
  orthogonal to the content-serve charter.
- **DEF-M21-01** (replayCmd conn-seam) → tracked tooling-debt follow-up (untouched by M25 — its ext fixes are
  firewall/directus/media/capture paths, not `replay.go`'s conn-seam).
- **DEF-M10-01** (S3 blob bytes + cloud store) → backlog (unscheduled); DB-1's served-cover-resolves-to-prod
  evidence re-confirms the deliberate refs-only posture works.
- **RESOLVED, not carried:** **DEF-M21-02** (serve-live-integration harness) landed Fate-1 in M25 (its destination
  milestone — the live serve-proof IS the integration it needed). Drops off the ledger.

## Metrics Delta
- **Rosetta test delta: none** — M25's rosetta branch is markdown/text (the field-bake log + 2 evidence files + 2
  skill arg-hint fixes + the close records). No rosetta code/test/lint/bench surface.
- **Ext (`rosetta-extensions`, tag `prop-room-m25` @ `1a2fd91`):** the 4 field-bug fixes + **+8 harden tests** —
  firewall `AssertPlan` **98.0→100.0%** (the 2 reject branches), directus closure 100% (behavioral composition +
  tenant-set symmetry), all `cmd/stacksnap` M25 fix functions 100% (+1 fuzz, ~960k execs 0 crashers). **0 bugs** in
  harden (the build's fix code held), **0 flakes** (3 clean sequential runs).
- **Field bugs:** 4 surfaced live + fixed Fate-1 (1 P1 tenant-leak firewall-caught + 2 P2 dangling-FK + 1 P2
  offline-build). 0 regressions.
- **Review:** 3 findings (1 scope checkbox nit / 0 code-quality / 1 docs / 0 tests / 1 blend + 5 archive).
  **Deferral audit:** GREEN — the strongest outcome (DEF-M21-02 RESOLVED in its destination milestone); 2 M25 env
  items fated fresh; **0 escape-hatch**, 0 repeat, 0 aged-out.
- **Release status:** v1.5 "prop room" is **feature-complete** — all 5 milestones M21→M25 closed + merged to
  `release/01.50-prop-room`. Next: `/developer-kit:close-release`.
- Full machine-readable record: [`metrics.json`](metrics.json).
