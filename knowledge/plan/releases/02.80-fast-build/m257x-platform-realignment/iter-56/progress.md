**Type:** tik — under `TOK-04` (*pin the target, or stop calling it a measurement*), executing the
orchestrator's re-direction of its `Next-tik direction` (clause-5 sweep deferred; the ref baseline first,
because clauses 1 and 2 were unmeasurable).

# iter-56 — the pin advance was aimed at the wrong cause, and the cause was the host

## Platform ref, per P3

    at iter open    git fetch origin main -> 0dab54d ; clone already level (no re-point needed)
    at iter close   see the close section

First iteration in three that opened level: iters 54 and 55 each had to re-point inside themselves.

## The premise this iteration was handed, and what measuring it cost

The orchestrator resolved iter-55's `user-blocker` by directing an `app` pin advance
`v1.363.2 → v1.365.0`, on iter-55's root cause: *compose at `0dab54d` deleted `STORAGE_RPC_ADDR`, the
pinned app still reads it at `main.go:446/516/983`, so `backend` exits 0 in silence; the app half of the
storage fold is not in the pinned release.*

**Ten seconds of measurement, taken before touching anything, refuted the remedy:**

| probe | result |
|---|---|
| `git grep STORAGE_RPC_ADDR v1.365.0` | still read at `main.go:450`, `:520`, `:988`, `internal/jobsimwiring/wiring.go:115` — the same three sites, shifted 4–5 lines |
| `git rev-list --count v1.365.0..origin/main` | **0** — `v1.365.0` **is** app origin/main |
| the v9.0 storage-in-app app half | exists at no ref; the only v9.0 commits in range are `docs(plan):` design commits |

There is no app build, anywhere, in which that read is gone. The advance could not have restored the
variable, so it could not have been the fix.

### The real cause, by experiment rather than by inference

Three runs of the same image on the same network, using the dead container's own env:

| run | result |
|---|---|
| env as-is, **no mounts** | **starts** — 93 log lines, `RPC server started :8083`, `Web server started at :8082`, alive >2 min |
| env as-is + this host's `$HOME/.aws/credentials` mount | **exit 0 in 0 s, 2 log lines** — the exact stack signature |
| env as-is + a regular **empty file** at that path | **starts**, still up at 25 s |

`~/.aws/credentials` **does not exist on this Mac**. Docker does not fail on a missing bind source and
does not warn — it creates the source as an empty **DIRECTORY** and mounts that. (`~/.aws` and
`~/.aws/credentials` both carry a creation time of 16:12 today: compose made them.) The app's AWS config
load then dies with `read /root/.aws/credentials: is a directory`, and the process **exits 0** — which
reads as an orderly shutdown to every log-reader and every supervisor there is.

`platform-alignment.md` §5 **Trap E** — *the tooling's own host preconditions are invisible until a clean
host* — on the new Mac. Not a platform-version skew, and no platform inconsistency was involved.

The generalization is now §5 **rule 28**, *three true facts do not make a cause — join them with one
experiment*, with the corollary that cost this iteration nothing and iter-55 a full cold cycle: **check
that the proposed remedy contains the fix before taking the remedy.**

## Phase A — the host precondition becomes a derived, fenced pre-flight

P4 order: derivable, so derived. `platform_topology.py` (iter-55's module) gains `host_bind_mounts()` +
`check_host_mounts()`: host-absolute bind sources belonging to the **default profile's** services, read
from the platform's own compose. Scope comes from real properties, not a list — `./data/postgresql` is out
because it is workspace-relative and stack-owned; `storage`'s mount is out because `storage` is not in the
`core` profile at `0dab54d`. Both exclusions follow the platform with no human action.

**The check's shape is forced by the measurement, and this is the part worth keeping.** The obvious check
is *"does the bind source exist?"* — and it reports **GREEN on the exact host state that produced the
defect**, because Docker had already created the path. The residue of an auto-creation is a path that
**exists as an empty directory**. `test_this_fence_has_TEETH` runs that existence-only mutant and asserts
it **misses**, so the fence's weaker predecessor is pinned as a negative control.

Watched **RED on the real host and real clone**, naming `/Users/marco/.aws/credentials` and the exact
remedy; then GREEN after the repair. Block-scoped per §8 rule 6 — `ports:` uses the identical
`- "5432:5432"` spelling one key away, so the fixture carries a published port as a decoy.

**Suites:** stack-injection **326 OK** (316 baseline + 10 new) · demo-stack **7F/1038**, the exact recorded
baseline · the four platform-alignment guards **74 OK** · `shellcheck` **present on this host and clean**.

## The pin advance, kept for a better reason, and recorded per §7 rule 4

`v1.363.2` was three releases behind `app`'s own origin HEAD, and the gate says *against origin HEAD, never
a pinned pre-drift commit*. So the advance stands — on the gate's authority, not on a storage fix that does
not exist.

**What it contains** (37 commits):

| dimension | finding |
|---|---|
| migrations | **2**, both `ALTER TABLE … ADD COLUMN` with a default or NULL (`course_builder_sessions.brief`/`credits_spent`; `academy_chapter_progresses.completed_at` + an idempotent backfill) |
| destructive DDL | **0** `DROP TABLE` / `DROP COLUMN` / `RENAME` in the whole range |
| new hard-required config | **0** — not one `log.Fatalf` added in any non-test Go file |
| new env reads | `STRIPE_SECRET_KEY`, `BREVO_KEY` (ungated constructor args, `main.go:382`/`:437`), `WORKFORCE_TEST_DB` (test-only) |
| RPC addresses | unchanged |
| feature surface | member-analytics (new `internal/analytics` + 5 REST endpoints), course-builder credits/rename, academy `completed_at`, wundergraph drop, removal of a local-testing-only `WORKFORCE_FORCE_ORG_ID` override |

**Purely additive schema, no removed contract** — the safest shape an advance can have, and the reason PR-2
predicts rather than hopes. The class that broke the seeders at v2.1 and v2.7 was a *removed* table; there
is no removal here. The canonical `demo-stack/clones.pin.json` advances with it (`app` → `v1.365.0`,
`platform` → `0dab54d`) so the proven combination lives in a **committed file** (P2) instead of in whatever
a clone happened to be checked out at.

## Measurement 1 — clause 1, cycle A

```
refs:
  platform:  0dab54dfac6beacdef54a671e2500d3940fd7329   (origin/main; fetched at iter open)
  app:       v1.365.0 (bff61c91)                        (== app origin/main; advanced this iter)
  rext:      fast-build-m257x-iter-56                   (be657d3; tag verified on origin; clone re-pinned)
  rosetta:   5c9c099cdbdc576e431ce004f3b0e48197817fb8   (+ this iter's work)
  taken:     2026-08-03T14:38:20Z -> 14:47:13Z, cold `down --purge` -> `up-injected.sh 1`
  verdict:   {"project":"demo-1","offset":10000,"warnings":0,"green":true,"ts":"2026-08-03T14:47:13Z"}
```

**GREEN, 0 warnings, 8 m 35 s.** All 14 asserts pass, including the three that were red in iter-55:
`backend /api/health 200 on :18082`, `container liveness: all 11 expected container(s) running`. Plus
`sentinel.casbin_rules = 1251` · `public.skills = 42790` · directus per-stack-local · demo-patches all
applied, none refused, none skipped · frontend builds are this run's · hiring org set-dressed.

**The teardown, which iter-55 fixed, worked:** `down rc=0 survivors=0` — the label sweep confirming zero
containers under the project label.

**§5 rule 15 — the path cycle A took:** the **fresh-bootstrap** path.
`executing the per-stack provision (bootstrap → apply-structure → replay → boot)` → `node cli.js bootstrap`
→ structure auto-provision → replay (330,261 taxonomy rows) → boot. **No race observed, no retry fired.**

**PR-1 is confirmed.** The mount was the cause; the pin was irrelevant to it. It follows that iter-55's
cycle A would have gone green at `v1.363.2` with the mount repaired — the advance neither caused nor cured
the red.

## Measurement 2 — clause 2, and it took two readings

```
refs:
  platform:  0dab54dfac6beacdef54a671e2500d3940fd7329
  app:       v1.365.0 (bff61c91)
  rext:      fast-build-m257x-iter-56
  rosetta:   5c9c099cdbdc576e431ce004f3b0e48197817fb8   (+ this iter's work)
  taken:     2026-08-03, on the cycle-A stack
  command:   stack-demo/rosetta-extensions/playthroughs/e2e/run-playthroughs.sh 1 --reset
```

| reading | result |
|---|---|
| #1 | `passing=29 failing=1 unimplemented=1` — **rc=1**, `ptreport: GATE no-regressions FAILED` |
| #2 | `passing=30 failing=0 unimplemented=1 unimplementable=0` — **rc=0** |

**Clause 2's figure is met on reading #2, and reading #1 is reported rather than buried.** The single
failure was `pt-assignment-assign` (M243's assign-WRITE half), and its arithmetic names it as a **flake,
not a regression**:

    expect(count).toBe(before - 1)     Expected: 15   Received: 14

`before` was **16** and the grid settled at **14** — the assignable-affordance count fell by **two**, not
by one. A strict `toBe(before - 1)` over a baseline sampled while the grid is still settling fails in
exactly this shape, and the 20 s predicate timeout is spent waiting for a value the page has already gone
past. It passed on the immediately following reset-to-seed run, on the same stack, same refs.

**It is NOT attributable to the pin advance**, and that was checked rather than assumed: the advance
touches `internal/workforce/` in two files only (`types.go` +4/−1, `testdb_test.go`), `internal/assignments/`
not at all, and its one removal (`8fcbe09f`) deletes a `WORKFORCE_FORCE_ORG_ID` local-testing override that
a demo never sets. **PR-2 holds: the seeders survived the advance unchanged** — the class that broke at
v2.1 and v2.7 did not recur.

**PR-3 is refuted as stated.** It predicted `30 / 0 / 0`; the first binding reading returned `29 / 1 / 0`.
A suite that needs two runs is not a suite that passes, and `FIX-M257x-iter56-assignment-flake` is routed
with the arithmetic above. Recorded plainly because iter-55's lesson was that a refuted pre-registration is
the most valuable output available.
