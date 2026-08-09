**Type:** tik — under `TOK-08`, steered by the user redirect of 2026-08-09.

## Phase 1 — sealed

Predictions `P-224-1..4` sealed in this iter's `overview.md` before any `git fetch`. Population measured
pre-fetch: **150** citation occurrences into the six stale clones, **123** with an explicit `:NN` pin,
over 26 corpus files.

## Phase 2 — the fetch

Six clones, all six fetched cleanly. `HEAD..origin/main` per repo:

| repo | local HEAD | origin tip | commits behind |
|---|---|---|---|
| `storage` | `4ce8ece52` | `9f8cb53` | **20** |
| `messenger` | `fa47850d9` | `e9421c6` | **7** |
| `jobsimulation` | `462343b05` | `82cb66ec` | **4** |
| `cms` | `ca50c8170` | `f38c0c4` | **2** |
| `roadrunner` | `87d8d4438` | `87d8d44` | 0 |
| `graphql-wundergraph` | `60c229f39` | `60c229f` | 0 |

**`P-224-1` is REFUTED at first contact: 4 of 6 advanced, predicted ≤ 2.** And what landed in those four
is not incidental — it is the platform correcting the very claims this corpus carries:

- `messenger` `459b184` — *"correct the ECR claim — production-messenger was destroyed, not preserved"*
- `storage` `bef79bf` — *"correct the ECR claim — production-storage was destroyed, not orphaned"*
- `storage` `775196b` / `messenger` `22fb4ef` — *"freeze the repo — X is folded into app"*
- `storage` `e6fac97` — *"drop the ECS service, keep the assets, export them as outputs"*
- `jobsimulation` `6092c6d2`, `cms` `6efa1d5` — the M810 teardown commits the corpus already cites

## Phase 3 — the census

**The sealed population figure was my own overestimate, and the seal caught it.** The pre-fetch count
used `\b<repo>/` and matched **58 occurrences of `app/internal/cms/…`, `app/internal/jobsimulation/…`**
and friends — a repo name appearing as a *path segment inside `app`*. With a negative lookbehind the
population is **92 occurrences / 78 pinned**, not 150 / 123. Corrected here rather than carried.

Every one of the **78** pinned citations resolved at **both** local `HEAD` and `origin/main`:

| class | count |
|---|---|
| resolves at both | **76** |
| **rotted by the advance** | **0** |
| fixed by the advance | 0 |
| broken at both | 2 |

and both "broken at both" are **false positives of my own instrument** — `jobsimulation/ai/ai.go:267`
and `:129` (cited from `shared_libraries.md:200,202`) are *tails* of `app/internal/jobsimulation/ai/ai.go`,
which the surrounding prose names explicitly one line above. Likewise the one missing unpinned path,
`storage/storage.go` at `storage.md:116`, is a row in an indented **code-map tree** under `internal/`.
**True broken-citation count into these six repos: 0 of 92.** `P-224-2` holds, though on a denominator
that was itself wrong.

### And yet three load-bearing claims resolved ONLY at origin

Line-range resolution is the weak arm. The claims the corpus actually rests its M810 story on are about
**content**, and there the two substrates disagree completely:

| corpus claim | at local clone HEAD | at `origin/main` |
|---|---|---|
| `messenger/terraform/main.tf:29` = `service_desired_count = 0` | `container_definitions = <<EOF` | **`service_desired_count = 0`** ✓ |
| `jobsimulation/terraform/main.tf:15-22` = the M810 decommission comment | an atlas migration data block | **the M810 comment** ✓ |
| `storage/terraform/main.tf` is *"18 lines"* | **100 lines** | **18 lines** ✓ |

**The corpus is right and the substrate was wrong.** Every one of these claims is true of the platform as
it actually is, and false of the clone set the fence family reads. This is the mirror image of iter-222:
there, a stale clone let a **rotted** claim read green; here, a stale clone would have made a **correct**
claim read red. The direction of the error is not predictable from staleness alone.

Decisive corroboration: the corpus cites **both** the old HEADs and the new tips — `462343b0` (6 hits)
*and* `82cb66ec` (3); `4ce8ece` (7) *and* `9f8cb53` (9). **The corpus already knew the newer refs. Its
own graded substrate did not.**

### `P-224-4` — the guard is not merely blind to a stale clone; staleness SATISFIES it

`clone_drift_guard` returns **`OK`, exit 0** against the stale substrate and **`OK`, exit 0** against the
advanced one — the identical verdict for two materially different worlds, so the verdict carries no
information about freshness. The mechanism is in its D1 arm: *"at least one cited sha IS the clone's
current HEAD."* A clone parked on an **old** cited sha satisfies that exactly as well as one at the tip,
and the numbers above show the corpus cites both. **Not-fetching is a way to pass.**

This is distinct from the gap iter-222 already recorded: `clone_pin_guard`'s docstring disclaims
freshness for `clones.pin.json` (the 4 `repos.yml` repos + 2 sanctioned extras) and routed it as
`ROUTE-M257x-222-pin-advance-needs-a-reproof`. Neither that pin nor that route covers these six legacy
clones, and neither covers `clone_drift_guard`'s substrate.

## Phase 4 — repair

**Substrate.** The four stale clones advanced by `merge --ff-only` to their origin tips (`cms` `f38c0c4`,
`jobsimulation` `82cb66ec`, `messenger` `e9421c6`, `storage` `9f8cb53`), all four now `behind=0`. Same
treatment iter-222 gave `platform`, which this iter confirms is still at `behind=0`. No clone had tracked
modifications; `cms`'s single dirty entry is an untracked `studio/` directory, left untouched.

**Prose — one claim, two sites, and the repo contradicts itself.** `messenger/terraform/main.tf:27-28`
@ `e9421c6` still reads *"The image and task definition stay declared: this is the rollback path.
Restoring service is a one-line revert plus an apply, not a re-provision."* The corpus quoted it verbatim
at `corpus/services/messenger.md:59` and `corpus/architecture/platform-migration-status.md:94`. **The
same commit that IS `origin/main` retracts it in another file**: `459b184` rewrote `messenger/CLAUDE.md`
to record that the `production-messenger` ECR was **DESTROYED on 2026-08-05**, hand-deleted with
`production-storage`, `production-customerio-sync`, `production-skiller` and `production-wundergraph`
(infrastructure #3253), leaving *"exactly six repositories, one per live service"* in eu-west-1 and the
`removed { destroy = false }` block **inert** — so there is no image to revert to and restoring is *"a
re-provision, not a revert."* The terraform comment was simply never updated. Both corpus sites corrected;
both edits are inline within a single table row, so **no line numbers moved** (`1↔1` on both files) and no
citing site needed re-pinning.

Swept for the obvious sibling defect — no corpus site claims any of the five destroyed ECR repositories
still exists. The `production-storage…` hits are **S3 bucket names**, a different resource.

## Close — 2026-08-09

**Outcome:** the six clones nobody had fetched in four days were fetched; **4 of 6 had advanced** (refuting
the sealed prediction), the corpus's 92 citations into them were censused and **0 are broken**, three
load-bearing M810 claims were shown to resolve **only at origin** — and the retracted "rollback path" claim
the platform corrected on 2026-08-05 was corrected here too.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-224-1` (advancing a stale clone is substrate repair, not a finding),
`D-M257x-224-2` (the instrument's own false positives are reported, not silently filtered).
**No `N`/`P` movement is claimed** — this iter took no graded reading.

**Predictions, graded:**

| id | prediction | result |
|---|---|---|
| `P-224-1` | ≤ 2 of 6 advanced | **REFUTED — 4 of 6** |
| `P-224-2` | ≤ 8 pinned citations fail at origin | **HELD — 0** (on a corrected denominator, 78 not 123) |
| `P-224-3` | 0 fetch failures | **HELD — 6 of 6 fetched** |
| `P-224-4` | guard green before AND after | **HELD — `OK`/exit 0 at both substrates** |

**Side-deliverables:** the population-count correction (150/123 → 92/78) is a repair to this iter's own
sealed figure, disclosed rather than quietly restated.

**Suite state at close** — `guard_family`, repo-root `.`: **19 GREEN · 0 RED · 10 not-run** without
`--platform`; **24 GREEN · 0 RED · 5 not-run** with `--platform stack-demo/platform`. The 5 not-run are
commit/ledger-scoped members with no input supplied — that is **not** a whole-family green and the runner
says so itself (`EXIT 2`). No pytest section was run this iter; the `stack-core` baseline is untouched by
it (this iter modified no rext code).

**Routes carried forward:**
- `ROUTE-M257x-224-drift-guard-blind-to-stale-clone` → a later tik. `clone_drift_guard`'s D1 is satisfied
  BY staleness; it needs either an origin-freshness arm or an explicit UNMEASURED verdict when a clone is
  behind its own origin. Deliberately **not** landed here: the user's redirect puts instrument work below
  corpus + working-stack work, and it is a third line of investigation in this iter.
- `ROUTE-M257x-222-other-clones-never-fetched` → **CLOSED by this iter.** All 13 `stack-demo` clones are
  now at their origin tips.
- `ROUTE-M257x-222-pin-advance-needs-a-reproof` → still open, unchanged; still gated on clause 1's three
  cold cycles.
- `ROUTE-M257x-223-classify-the-ten-drifted-baselines` → still open, unchanged.

**Lessons:**
1. **A resolved anchor quoting a verbatim line is not the source's position.** A repo can retract a claim
   in one file and leave it standing in another **at the same ref**. `anchor_subject_census` grades this
   GREEN — correctly, by its own contract — because the cited line does say the thing. When a quote is
   load-bearing, read the repo's *retraction surface* (`CLAUDE.md`, `README.md`) and not only the anchor.
2. **A freshness gap is not a quiet gap; it is an inverted one.** A guard that grades the corpus against a
   clone and never against that clone's origin turns *"we never fetched"* into *"no drift"* — staleness
   makes the assertion easier to satisfy, not harder.
3. **A stale substrate can make a CORRECT claim look wrong.** iter-222 found the other direction. Neither
   direction is the default, so "the clone is behind" is never a safe thing to reason past.
4. **Widen a repo-name regex and you will match the monolith that absorbed it.** `\bcms/` matches
   `app/internal/cms/`. After a merge programme, every absorbed service's name is also a *path segment* in
   `app` — 58 of 150 apparent hits here. Pre-registration is what made this visible as an error rather
   than a published figure.
