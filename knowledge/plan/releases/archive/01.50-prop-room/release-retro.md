# Release Retro: v1.5 "prop room"

**Shipped:** 2026-06-14, tag `v1.5`. **Milestones:** M21–M25. **Thesis:** turn "every stack reads content live from prod" into "every stack serves its own captured public catalog locally, with real images."

## Headline
v1.5 delivered the local-Directus capability end-to-end and **proved it observably on the real box** in the closing field-bake — a freshly bootstrapped, stacksnap-provisioned per-stack Directus serves the captured public catalog (data plane local, asset plane prod) with zero hand SQL. The defining moment of the release was the **tenant-data firewall holding fail-closed under live prod capture**: M25's first real capture tried to over-pull 158 tenant-referenced files; `AssertPublicOnly` refused to persist; **zero data leaked**; the fix went into the *filter*, never the firewall.

## Incidents this cycle
- **P3 (process) — M22 attempt-1 build crash + resume-in-place.** A network error crashed the build mid-§4; §1–3 were committed, attempt-2 reconciled + finished §4–6. No data loss; recovery discipline (no stash/reset, resume on committed state) worked.
- **P2 (env) — M25 Docker self-outage.** The field-bake agent programmatically resized the Docker VM (7.65→12 GB target) via `settings-store.json` and restarted Docker Desktop to apply it — the restart made Docker unreachable and self-interrupted the run. **Lesson:** never mutate Docker Desktop settings mid-run; the 12 GB prereq doesn't even boot on a 16 GB host (~10 GB is the practical ceiling — now a documented field note, M25-D2).
- **3 real release bugs the field-bake caught (none visible to clean-room unit tests) — all fixed Fate-1:**
  1. **Offline `GOTOOLCHAIN`** (M25-D1): M24's `go1.25.11` pin broke the offline (`GOPROXY=off`) clerkenstein build → `/demo-up` aborted. Fixed with `GOTOOLCHAIN=local`. *Its mirror test had the same drift — caught + fixed at close-release.*
  2. **`directus_files` tenant-file over-capture** (M25-D5): the referenced-subset closure followed the tenancy-column-less `resource` table → would have leaked 158 tenant files. Firewall fail-closed; fixed in the filter + defense-in-depth `AssertPlan`.
  3. **Dangling-FK class** (M25-D6/D7): captured public content referenced uncaptured admin/library UI tables (`group`/`folder`/`uploaded_by`/`modified_by`) → apply/replay FK violations. NULLed the cosmetic refs on capture.
- **close-release docs RED (1 must-fix):** the M24 corpus truth-up **missed `snapshot-cold-start.md`** entirely (it wasn't in any milestone diff) — it still shipped a false "M10 gap / nothing executed / exit 4" claim contradicting the release's headline. Caught by the release-level docs sweep; rewritten to the two-path truth.

## Cross-milestone patterns
- **The two-repo split held cleanly all release:** code → `rosetta-extensions` (tags `prop-room-mNN`), docs/plan → rosetta (the milestone branch). Every milestone navigated it without confusion; the only friction was a third repo (`ant-singularity`) for M24's `/project-stats` fix, transparently flagged.
- **Inherited-deferral resolution was strong, not just tracked:** M21's two Fate-3 items both *landed* in M23; DEF-M21-02 *resolved* in its M25 destination (the live serve IS the integration harness it wanted). No repeat-defer pattern across all 5 close-audits.
- **A recurring "doc didn't get swept" class:** M24 swept the corpus but missed one file not in any diff; the README-index-row guard (M24) catches *missing index rows* but not *stale-but-present* claims — a grep-for-retired-vocabulary check at close-release is the backstop that caught it.

## Carry-forward (destinations recorded)
- **DEF-M21-01** (replayCmd conn-seam hermetic test) → `roadmap-vision.md` Unscheduled backlog (added at close-release).
- **M25-D9** (dev-`N` taxonomy `rc=4` migrate-ordering) → `roadmap-vision.md` Unscheduled backlog (non-fatal, orthogonal to content-serve).
- **DEF-M10-01** (S3 blob bytes + cloud store) → backlog, re-signed; sting removed by asset-plane-on-prod.
- **K-AIFUNX-E658** — a real prod DATA-quality residual (a public sim references a customer-only skill) surfaced by M23's cross-surface gene; **operator-owned**, uncloseable by tooling. Documented, not a tooling deferral.
- **CI wiring** — triple-clean ran as the local 3× fallback (no CI wired for this corpus repo); wiring real CI is a standing nice-to-have.

## Metrics delta (vs v1.3b)
- Go test funcs **736 → 867** (+131; stack-snapshot carried most, +102, from the structure/capture/firewall work).
- Python **360 → 459 collected** (+99; every suite grew or held).
- Flake **0 → 0**; coverage no >2pp drop (firewall `AssertPlan` 98→100% at M25); supply-chain GREEN (the 12 stdlib advisories pinned out by go1.25.11).
- **Stats delta:** see the Phase 8c `/developer-kit:project-stats` release-close snapshot (compare against the v1.3b snapshot for code-size + velocity).
