# iter-02 — decisions

## D1 — Member-field backfill targets `memberships.{joined_at,location,last_activity_date}` (not user_basic_info)
**Context:** the annotation flagged `/enterprise/members` showing no location / join-date / last-activity. The
re-diagnosis found these are SEPARATE NULL columns on `public.memberships` (0/221), distinct from the seeded
`user_basic_info.location` (the /profile header) and the membership `created_at` (the row's audit timestamp).
**Choice:** extend `UsersSeeder` (`users.go`) to write all three in the memberships COPY + an idempotent
re-seed backfill. `location` reuses `locationForUser(uid)` (so a member's /profile-header and /enterprise/members
location agree); `joined_at` = the same ramped instant as `created_at`; `last_activity_date` = a recent
deterministic instant after the join.
**Why:** the M44 §D `picture_url` lesson — `/enterprise/members` reads membership columns, not user/profile
columns; fill the column the surface actually reads.

## D2 — The member-field backfill uses a NULL-ONLY guard, not IS-DISTINCT-FROM (idempotency)
**Context:** the picture_url / user_basic_info backfills use an IS-DISTINCT-FROM guard (safe no-op on re-seed
because their values are deterministic). Two of my three fields (joined_at, last_activity_date) are computed
relative to `time.Now()`, so IS-DISTINCT-FROM would re-fire on every re-seed (the recomputed instant always
differs) — breaking the M17 re-run-idempotency contract (caught by `TestSeeders_ReSeedInsertsNothingNew`).
**Choice:** a NULL-ONLY guard (`WHERE joined_at IS NULL AND location IS NULL AND last_activity_date IS NULL`) —
heal a pre-M50 row exactly once, then no-op. A fresh seed fills the columns in the COPY itself, so the heal is
already a 0-row no-op on a fresh stack; it exists solely to fix a stack seeded before M50.
**Why:** preserves "2nd seed changes nothing" without forcing the timestamps to be wall-clock-deterministic
(which would diverge from how the rest of the fleet computes now-relative instants).

## D3 — iter-02 closes closed-fixed-partial; the MANAGER baseline is blocked by a frontier-cap tooling gap (→ iter-03 tooling-iter)
**Context:** the EMPLOYEE baseline sweep was gate-valid (frontier EXHAUSTED at 59 pages, cappedAtFrontier=false)
→ **gateMet=true**. But the MANAGER baseline sweep does NOT exhaust: the manager's reachable set includes the
per-member `/user/<uuid>` team-roster fan-out (221 Cervato members) AND per-sim/per-skill-path
`/enterprise/activity-dashboard/.../<uuid>` drill-downs, which inflate the BFS frontier far beyond the 150-page
`COVERAGE_MAX_PAGES` cap (observed q reaching 172 at page 39). A capped sweep is, per the protocol's measurement
convention, "FLOORS over a truncated slice, not the true residual" → **not gate-valid**.
**Options:** (a) raise `COVERAGE_MAX_PAGES` to whatever exhausts (likely 300+, a ~40-min sweep, and re-runs each
iter — costly + machine-unfriendly on the 9 GiB VM); (b) tighten the manager SAMPLE_RULES so the
template-identical fan-outs (`/user/<uuid>`, the activity-dashboard per-entity drill-downs) collapse to a
representative + boundary sample (the protocol's documented allowance for a "huge template-identical set"),
keeping the frontier where escapes/failures live exhausted.
**Choice (routed to iter-03 as a tooling-iter):** option (b) — tighten the SAMPLE_RULES (with a defensible
representative+boundary sample), so the manager gate becomes measurable + the sweep stays fast. This is the
protocol's tooling-iter pattern ("a tik closes blocked on a missing harness capability → ship the capability AND
use it within the same iter").
**Why iter-02 closes now:** the employee gate (valid) + the member-field fix (complete, unit-tested code) are
landed deliverables; the manager baseline + the member-field fix's gate-verification require the iter-03 tooling
fix first. Closing now keeps the cap-finding + the strategy revision reviewable; the capped sweep was stopped
(no value in an 18-min gate-invalid run + machine-aware: don't thrash the VM).

## D4 — F1 reconciliation (the coverage-protocol's "manager gate MET" claim)
The doc calls the manager gate MET as of M42m iter-04. The re-diagnosis shows the manager manifest's
`members-roster` section asserts only names+emails (`minMeaningfulLen 200`), NOT the location/join/last-activity
columns — so the member-field fix improves believability that the CURRENT manifest does not measure. M50 will
need to STRENGTHEN the manager manifest (assert the annotation columns) so the gate actually proves the gaps are
closed — routed to a future tik alongside the languages + cert-coverage fills. The "gate MET" claim was scoped
to a NARROWER manifest + the demo-3 calibration; M50 is filling/asserting sections the manager manifest never
covered, not regressing a green gate.
