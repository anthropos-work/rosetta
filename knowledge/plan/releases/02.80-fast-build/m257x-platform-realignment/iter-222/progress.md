**Type:** tik — under `TOK-08`, redirected by the user (`D-M257x-222-1`).

## Probe — sealed before any repair

See [`probe-evidence.md`](probe-evidence.md) and the pre-registered P1–P6 table in
[`overview.md`](overview.md). Sealed as this iter's first commit.

## What the census found

Every pre-registered claim held, and the fence reproduced P4 exactly — 5 findings, naming the same 5
repos, on its first run against the live tree.

| # | pre-registered | measured | verdict |
|---|---|---|---|
| P1 | pin names **11** repos | 11 | held |
| P2 | `repos.yml` @ `0c91421` names **4** | app · sentinel · next-web-app · studio-desk | held |
| P3 | **2** sanctioned extras → legitimate population **6** | platform · ant-academy | held |
| P4 | **5** phantom keys | cms · jobsimulation · storage · messenger · roadrunner | held — fence RED, 5 findings |
| P5 | **3 of 6** live pins behind `origin/main` | app 28 · next-web-app 12 · ant-academy 9 | held |
| P6 | **0** fences assert pin membership | 0 | held |

**The platform orchestrator repo has not moved.** `platform` `origin/main` is `0c91421` — the *same* commit
gate clauses 1+2 were proven at, 0 behind. What moved is underneath it: `app` **28**, `next-web-app` **12**,
`ant-academy` **9**. The run brief's premise that "the platform has moved" is true of the *clones* and false
of the *orchestrator*, and the distinction matters because `repos.yml` — the definition of *"the remaining
ones that are still part of it"* — lives in the repo that did not move.

## What landed

1. **`stack-core/clone_pin_guard.py`** (FENCE-M257x-iter222), 3 arms, all derived:
   - **A no phantom** — every pin key ∈ `repos.yml` ∪ the two sanctioned extras.
   - **B no hole** — every `repos.yml` repo has a pin entry, because `DEMO_ADVANCE_CLONES=pinned` leaves an
     unpinned repo **untouched**: a hole is a repo the barrier silently does not cover.
   - **C reproducible refs** — no moving branch. A barrier pinned to `main` names a different tree daily.

   It reuses `platform_alignment_guard.parse_repos_yml` (the registered derivation) rather than
   re-deriving the clone set — two derivations of one population makes the weaker one a silent census.
   Same no-default discipline for the reference path, for the same reason: a membership check against the
   wrong `repos.yml` **passes**.

2. **The repair.** The 5 phantom keys are gone; their shas are preserved in `decisions.md` so the
   provenance is not lost with them. Fence re-run: **exit 0**.

3. **`ensure-clones.sh`'s comment**, on two counts. *"jobsimulation stays standalone"* is **retracted** —
   it was already false when written. And *"the barrier's reproducible **current-origin/main** topology"*
   no longer asserts currency: a dated claim in a comment names a **run** and gets read as a **property**
   (`§8` iter-208). The comment now points at phase (e), which measures freshness per run, fetch-verified.

4. **`tests/test_clone_pin_guard.py`** — 16 tests, 0 skipped: a per-phantom mutation battery (one stage per
   dropped repo), an arrival control (the platform grows a repo → RED with no tooling edit), four
   anti-vacuity controls (empty pin / empty `repos.yml` / missing `repos.yml` / malformed pin are all
   **exit 2**, never green), and the sanctioned-extra **registry population** assert.

5. **Registered in `guard_family.INVOCATIONS`** — and the registry proved itself first: before the entry
   existed the family went **exit 2** naming `clone_pin_guard` as *"on disk but this runner has no
   invocation for it"*. That is the control, taken live.

6. **`corpus/ops/platform-alignment.md` §8** gains layer 7 + a new rule (below); `stack-core/README.md`
   gains its row.

## The finding this iter did not go looking for

**Nobody had fetched.** `anchor_construct_guard` and `repair_postcondition` resolve corpus anchors at the
app clone's `origin/main` — which pointed at `ad9f3c498`, *identical to the clone's own HEAD*. Both were
**green**. This iter ran a plain `git fetch origin` for the freshness survey, `origin/main` moved to
`3eaadae68` (28 commits), and the same two guards went **RED with 9 anchors** across four corpus documents
— **no corpus file changed, no guard edited.**

> A remote-tracking ref is a **cache**, not a remote. A guard resolving against `origin/main` on a box where
> nobody fetches is grading the corpus against *the corpus's own clone*.

The 9 are captured verbatim in `probe-evidence.md`. **They were going to be routed to iter-223 — and the
tooling refused.** The `repair_postcondition` pre-commit hook blocked the close outright: *"9 site(s)
restate an already-refuted claim and are NOT in the baseline. A repair may remove these; it may never add
one."* So the repair landed here, which is the correct answer and not the one this iter had chosen. **A
fence that only advises is a fence you can plan around.**

**Adjudicated first, per site: pinned or not.** `§5` rules 41/44 make a ref-scoped claim settleable at its
own ref, so a ref-pinned site read at `origin/main` would be the *guard's* error. **None of the nine is
pinned** — every cluster is bare (`app/main.go:15`, `` `:62` ``, …; `backend.go:289`; `readiness.go:710`),
so all nine are corpus rot against the platform's real HEAD, and the repair is on the corpus side.

**Repaired by re-derivation, never by arithmetic.** Every new line number was found by locating the *same
construct* at `3eaadae68` — `grep -n` for the identical source line — not by adding a diff offset. The
offsets are not uniform (`app/main.go` `+0 / +1 / +14 / +24`; `backend.go` a flat `+5`), which is exactly
why an offset would have been wrong somewhere.

| document | file | re-derived |
|---|---|---|
| `platform-migration-status.md:94` | `app/main.go` | `:62`→`63` · `:63`→`64` · `:1416-1421`→`1430-1435` · `:1445`→`1459` · `:1450`→`1464` · `:1471`→`1485` · `:1473`→`1487` · `:285`→`286` · `:1552`→`1576` (`:15` unchanged) |
| `observability.md:28` | `app/main.go` | `:277`→`278` (`colony.WithLoggingTracing(0.15, 0.15)`) |
| `security_compliance.md:213/214/215/262/263/264` | `internal/web/backend/backend.go` | `:289`→`294` · `:295`→`300` · `:301`→`306` (×2) · `:309`→`314` · `:315`→`320` (`:117` unchanged) |
| `ai-readiness.md:483` | `internal/aireadiness/readiness.go` | `:710`→`716` (`keepInCycleStep1`; call site `:388` unchanged) |

**The floor caught 9; the cluster was 15.** `anchor_construct_guard` states its own limit — it detects
*"resolves to nothing"*, never *"resolves to the WRONG construct"* — so `backend.go:309`, `app/main.go:285`,
`:1445`, `:1450`, `:1473`, `:1416-1421` had all rotted onto real-but-wrong lines and were invisible to it.
Repairing only the flagged nine would have left a **half-repaired citation cluster**, which reads as
verified. Each sibling was re-derived by the same method and repaired in the same edit.

**Grade the direction:** the guards did not regress. They stopped being blind.

## Suite state at close — stated, not implied

`stack-core`, pytest 8.4.2 / `/usr/bin/python3`, Python. The whole-section run taken mid-iter came back
**1,876 passed / 4 skipped** with failures in **five** modules. Split by cause, each re-run targeted:

- **`test_frozen_expectation_census_m257x` — 9 failures, ALL this iter's own induction, ALL repaired.**
  Two literal ratchets breached by exactly **+1** each (`DOCSTRING_LITERAL_CEILING` 209 → **210**,
  `TEST_MODULE_LITERAL_CEILING` 566 → **567**), re-pinned **with a recorded reason** naming the module,
  which is the mechanism's own instruction rather than a blind bump. And two genuinely good catches on
  the new code: `clone_pin_guard.py::check` *"became executable-here and was not graded"* → registered as
  `DECLINE:verdict` beside its `platform_alignment_guard` twin; and a test literal that **duplicated a
  value the tree derives** (`["app","sentinel","next-web-app","studio-desk"]` ≡ `parse_repos_yml`) →
  exempted, because it is a historical fixture pinning the iter-222 measurement and deriving it would
  make the regression track the live tree and silently stop testing its own case. Re-run: **115 passed.**
- **The other four modules — 6 failures, ALL the same 9 anchors**, and none of them mine:
  `test_iter45_mechanical_fences` (1), `test_repair_postcondition` (1),
  `test_repair_postcondition_audit_mode` (1), `test_m257x_mechanical_fences_mutation_battery` (3). Every
  one traces to `repair_postcondition` reporting `anchor_construct_guard`'s 9 sites.

**Not induced, disclosed.** That RED lived in `origin/main`'s position, which `git fetch` changed and which
is not a tracked file in either repo. It was latent on every box that fetches; this one just did.

**All six modules re-run together after the anchor repair: 256 passed, 0 failed** (pytest 8.4.2,
`/usr/bin/python3`, `stack-core`, Python). `anchor_construct_guard`: *"OK — every resolvable anchor names a
construct."* `demo_knob_guard`: *"OK — the defaults table and the parsers agree, both directions."* No
whole-section re-read is claimed post-repair; what is claimed is these six modules and the guard family's
own verdicts.

## Close — 2026-08-09

**Outcome:** the tooling's own declaration of the platform topology named 11 repos where the platform names
6; 5 phantom keys removed and the class fenced both ways — and a plain `git fetch` disclosed 9 rotted
corpus anchors that a stale remote-tracking ref had been hiding.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** D-M257x-222-1 (the user redirect, recorded so it survives the run), D-M257x-222-2 (freshness
is disclosed, not auto-advanced), D-M257x-222-3 (the fetch finding is a disclosure, not a regression).
**No `N` movement is claimed** — this iter took no graded reading.

**Side-deliverables:** two stale `ensure-clones.sh` anchors in `corpus/ops/demo/demo-up-defaults.md`
(`:199` → `:212`, `:454` → `:467`) re-pinned. These are **this iter's own induction** — the expanded
comment shifted the lines — caught by `demo_knob_guard` on the post-repair sweep, which is the sixth layer
working exactly as specified.

**Routes carried forward:**
- `ROUTE-M257x-222-other-clones-never-fetched` → a later tik. `app` was fetched because the pin survey
  needed it. **`next-web-app` (+12) and `ant-academy` (+9) were fetched too and their anchors have not
  been swept**; `sentinel` and `studio-desk` were at origin. The class this iter found is *"a guard's
  reference is only as current as the last fetch"* — one repo's worth of it is now repaired, and the
  sweep across the rest is a bounded, mechanical follow-on.
- `ROUTE-M257x-222-anchor-guard-floor-leaves-siblings` → a later tik. The floor caught **9** of a **15**
  -member cluster; the other 6 were found only because a human re-derived the whole citation block.
  *Resolves-to-the-wrong-construct* was measured and declined at iter-121 on cost grounds — but the
  measurement that declined it did not know the sibling rate. It is now one data point: **6 of 15**.
- `ROUTE-M257x-222-pin-advance-needs-a-reproof` → gate clause 1. Advancing the pin past app+28 /
  next-web+12 / academy+9 is a cold-bring-up decision with a re-proof attached, not a manifest edit.
- `ROUTE-M257x-222-guards-that-read-origin-must-say-when-they-fetched` → a later tik. `--verify-remote` is
  opt-in for a good reason (offline-runnable); what is missing is the verdict *saying* the reference's age.

**Lessons:**
1. **A guard that reads `origin/<branch>` measures whatever the last fetch left there.** Written into the
   protocol doc §8 as a standing rule, with its three consequences. This is the generalizing lesson.
2. **`git fetch` is a measurement, not a mutation** — it touches no tracked file. Running it before a fence
   sweep is the difference between measuring the platform and measuring a memory of it.
3. **The tooling makes the same claims as the corpus and gets fenced less.** Layer 1 has fenced the corpus's
   migration map since iter-20; the tooling's copy of that claim shipped unread for four releases. When a
   class is fenced on the documentation side, ask what *executes* the same claim.
