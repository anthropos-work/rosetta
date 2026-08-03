**Type:** tik — under `TOK-04` (*pin the target, or stop calling it a measurement*), executing its
`Next-tik direction` for iter-55.

# iter-55 — re-establish the ref baseline

## P3 fired before the first measurement was taken

TOK-04 was written against platform `ef32d4c` and named it the target ref. It was stale within the day.

    origin/main at iter open   0dab54d  chore(compose): run without the standalone storage; rename graphql -> core
    our clone at iter open     ef32d4c  (1 behind)
    action                     git merge --ff-only  ->  0dab54d, tree clean, zero platform edits

This is the **second consecutive iteration** in which the detecting iteration has had to re-point inside
itself. It is not drift noise: `0dab54d` is the **v9.0 `support-in-app`** step that
`platform-alignment.md` §1 already listed as IN FLIGHT (PRs #1096/#1098/#1103). The program is running to
its published schedule; what has no schedule is our tracking of it.

## What `0dab54d` actually changed, and what it broke

| platform change | what it did to our tooling |
|---|---|
| `profiles: [graphql, …]` → `[core, …]` on `backend` + `gotenberg` | **five rext sites** named the profile as a literal |
| `storage` → `profiles: [storage-legacy]`, out of the default set | the default stack loses a container |
| `messenger` out of `all` | no demo impact; recorded |
| `context: ${APP_BUILD_CONTEXT:-../app}` | a parameterised build context our parsers had never seen |

**The profile literal is the dangerous one, and not because it errors.** `docker compose --profile graphql up`
against `0dab54d` is a *successful* command that starts **zero containers**. No error text, no non-zero
exit, nothing for a log-reader to catch. The bring-up would have reported success over an empty Docker host.
Compare the storage removal, which merely deletes a container — loud by comparison.

## Phase A — three hand-maintained tuples, replaced by one derivation

The re-point was not "change `graphql` to `core` in five places." That would have set the clock for the
next rename, which the platform's own roadmap says is coming. The survey found that the bring-up path
carries **three** hand-maintained statements of platform topology, all of the shape
`platform-alignment.md` §2 dissects for `migrate-demo.sh`'s 4-tuple:

| tuple | where | how wrong it was |
|---|---|---|
| the profile name | 5 sites — `up-injected.sh`, `rosetta-demo`, `gen_injected_override.py` (×3) | dead as of `0dab54d` |
| `verify_svcs` | `up-injected.sh` | **5 of its 10 services no longer exist** — `graphql` (deleted `2adcf71`), `cms`/`jobsimulation`/`roadrunner` (`ef32d4c`), `storage` (`0dab54d`) |
| `INJECT_SVCS="app cms jobsimulation"` | `up-injected.sh` | 2 of 3 are deleted services; it still cloned, patched and built injected images for them |

`INJECT_SVCS` is the one worth pausing on, because it *demonstrates* the §2 time bomb rather than
describing it. **It has not failed.** It works on this box only because the pre-prune clones are still on
disk — `stack-demo/` still carries `cms/`, `jobsimulation/`, `roadrunner/` and `graphql-wundergraph/`. On a
clean host `ensure-clones.sh` follows `repos.yml`, which at `0dab54d` names **six** repos and neither of
those two, so the loop reaches for directories that were never created. The next clean checkout arms it.

**`stack-injection/platform_topology.py`** (new) derives all three from the platform's own compose:

- **profile** ← the `backend` service's own `profiles:` row, minus its own name and `all`;
- **service set** ← every service carrying that profile, plus the always-on services that declare no
  `profiles:` key at all;
- **build (repo) set** ← the `build.context: ../X` of those services.

### The property the literals never had, and how it is tested

    same code, ef32d4c fixture  ->  profile "graphql", service set includes storage
    same code, 0dab54d fixture  ->  profile "core",    service set excludes storage

Ref-independence is the first thing `tests/test_platform_topology.py` asserts, because it is the only
reason to prefer a derivation over a corrected literal. A corrected literal is right today and wrong at the
next rename; this is right at both refs with no edit between them.

The parse is **block-scoped** per §8 rule 6, not a whole-file substring match, and the fixture carries a
`decoy-core` service declaring `profiles: [core, all]` **immediately before `backend`** — so a scan for
`profiles: [core` finds the decoy first and answers the wrong name with total confidence. Every fail-loud
branch (no `backend` service, no profiles, two candidates, missing file) has a fixture that reaches it, per
§8 rule 2: a fixture that cannot fail proves nothing.

**One correction the first cut needed, recorded because it is the §5 rule 4 failure verbatim.** The first
version read only `docker-compose.yml` and returned `sentinel backend gotenberg` — *a stack with no
database in it*. `postgresql` and `redis` are defined in `common.yml`, reached through a column-0
`include:`. Enumerating the search set is not optional; the derivation now follows the include.

### Two defects this iter introduced and caught, rather than shipped

Recorded in full, because the previous iteration's lesson was that repair text is the highest-risk text
there is.

1. **82 green tests turned red.** The first cut derived `INJECT_SVCS` at source time — where the lib-only
   test seam sources the script before `log`/`die` exist and with no platform clone present. Result:
   `log: command not found`, demo-stack **89F** against a 7F baseline. Moved into a named function
   immediately before the build loop; back to exactly **7F/1030**.

2. **A live fence started skipping, and a skip reads exactly like a pass** (§5 rule 8). That same change
   made `INJECT_SVCS=` an assignment of a *variable reference*, and `test_studio_acquisition_m257`'s
   live-clone fence reads that line with a regex. It captured the literal string `"$INJECT_CANDIDATES"`,
   matched no clone, and `skipTest`'d. The suite went from `skipped=2` to `skipped=3`; everything else
   looked identical. **The only reason it was caught is that the skip count was read** — which is the
   rule's entire point. The fence now reads `INJECT_CANDIDATES` and **refuses a non-literal value** rather
   than skipping on one.

**Suites after Phase A:** stack-injection **316 OK** (299 baseline + 17 new) · demo-stack **7F/1030**, the
exact prior baseline, 2 skips as before · dev-stack **138 OK**. `shellcheck` is **present on this host and
returns clean** on both edited scripts — stated explicitly because §5 rule 8 applies to shellcheck's own
absence, which is how iter-04 lost a finding.

Committed as rext `dfb0929`, tagged `fast-build-m257x-iter-55`, **verified on origin** with
`git ls-remote --tags`, and the `stack-demo` consumption clone re-pinned to it.

## Measurement 1 — clause 2, as a CONTROL (not a restoration)

Taken first because it was cheap and a live stack was already up. **It does not restore clause 2.**

```
refs:
  platform:  28c5f0dd1453336f8f935425a7f5dd1a87dd5645   (stack-demo/clones.pin.json — the ref this stack
                                                         was BUILT at, 44 h before this reading, and
                                                         THREE folds behind origin HEAD)
  rext:      fast-build-m257x-iter-37                   (stack-demo consumption clone at read time)
  rosetta:   c2de4ffd75aa3a2f5faaef68fae5482ad44db55a
  taken:     2026-08-03 15:22 local
  command:   stack-demo/rosetta-extensions/playthroughs/e2e/run-playthroughs.sh 1 --reset
```

**Result: `passing=30 failing=0 unimplemented=1 unimplementable=0`** (`total: 31`, coverage 0.968). The one
unimplemented entry is the declared in-manifest TODO carrying an explicit `will-not-build` verdict — not a
failure and not an error.

That is the clause-2 *shape*. **It is not clause 2.** The stack it ran on still had `cms`,
`jobsimulation`, `roadrunner` and `storage` containers running — four services the platform has since
deleted or defaulted out. A pass there says the *old* topology works, which is the one thing nobody
doubted. Publishing it as a restoration would be a claim about a ref the reading never touched: **its ref
would exist only by adjacency**, and the adjacency here is 44 hours and three folds wide. That is precisely
what happened to the iter-37 reading TOK-04 objected to.

Its value is as a control. It establishes that the suite, the seed and the harness are healthy going into
the cold cycles, so a red in the binding reading is attributable to the re-point rather than to rot.

## The cold cycle refuted the pre-registration — and found a teardown that removed nothing

TOK-04 pre-registered **both clauses green**. Clause 1's first cycle at the new ref did not reach a
verdict at all, because **the teardown that precedes it silently did nothing**.

`rosetta-demo down 1 --purge` printed a clean teardown and exited **0** with **eleven containers still
running**. The only signal was one line inside compose's own output:

    service "roadrunner" has neither an image nor a build context specified: invalid compose project

The stack's `docker-compose.injected.yml` was generated 45 hours earlier, when the platform still declared
`roadrunner`, `storage`, `cms` and `jobsimulation`. Overlaid on the *new* base compose, a stale override
contributes only **overrides for services that no longer have a base definition** — no image, no build — so
the merged project is invalid and compose refuses to act on the **entire project**, including the eleven
services that were perfectly fine.

What followed is worse than the failure, and is why this is a finding rather than a nuisance. The call site
was `compose down … || true`:

- the refusal was discarded;
- `purge_data_dir` then deleted the data directory **out from under eleven running containers**;
- `data purged` and `removing this stack's images` both printed;
- the registry slot was released, and `cmd_down` exited **0**;
- the bring-up started on top of all of it.

**Every previous "cold" claim on this box rests on this command.** It is the same shape as M218's F-9 (*the
purge that never purged*) one layer up: F-9 was a data dir that survived, this is the containers that
survived. `platform-alignment.md` §5 rule 7 names the class — a probe that satisfies itself. `compose down`
was treated as its own evidence of teardown.

**The fix asks Docker, not the compose file** — `sweep_project_containers` reads
`docker ps -aq --filter label=com.docker.compose.project=demo-N`, names what it finds, removes it,
**re-reads**, and fails the purge if anything survives. There is no compose-file-shaped fix available, because
the compose file is the thing that went stale.

### The fix caught a survivor on its first live run, by a second mechanism

```
$ rosetta-demo down 1 --purge          # at rext fast-build-m257x-iter-55b
⚠ demo-1: 'compose down' left containers behind — removing them by project label: demo-1-storage-1
==> demo-1: label sweep removed the leftovers; the project has no containers
```

This run's compose project was *valid* — the override had been regenerated — and `--remove-orphans` was in
the command as it always is. `demo-1-storage-1` survived anyway, because at `0dab54d` `storage` is still
**declared** in the base compose under `profiles: [storage-legacy]`, so it is **not an orphan**, while not
being in the `core` profile, so it is **not selected**. **A container can be simultaneously
not-an-orphan and not-selected, and fall through both.** Expect one of these per fold, at the moment a
service moves to a rollback profile instead of being deleted.

Landed as rext `fb94d85`, tagged `fast-build-m257x-iter-55b`, verified on origin, consumption clone re-pinned.
8 new tests, `docker` stubbed; four of them assert the **call site**, because a sweep that is not called
fixes nothing.

## Clauses 3 and 4, re-read at `0dab54d` — and clause 3 does NOT hold

Both were re-run rather than inherited, since P3 makes a reading at the previous ref inadmissible.

```
refs:
  platform:  0dab54dfac6beacdef54a671e2500d3940fd7329   (origin/main; clone fast-forwarded at iter open)
  rext:      fast-build-m257x-iter-55b                  (authoring copy fb94d85; tag on origin)
  rosetta:   c2de4ffd75aa3a2f5faaef68fae5482ad44db55a   (+ this iter's uncommitted work)
  taken:     2026-08-03, this iteration
```

**Clause 4 — MET.** `test_write_target_schema_fence` + `test_migration_derivation_fence`, **35 tests OK**
run with `PLATFORM_REPOS_YML` pointed at the `0dab54d` file. Ref-independent by construction, and it held
across the fold again, unaided.

**Clause 3 — the fence is GREEN and the map is FALSE.** `platform_alignment_guard` reports
*"platform-migration-status.md and repos.yml agree in both directions"* — correctly, because `0dab54d`
changed **no repo membership**. The guard's own header says it fences *membership*, not prose. The prose it
does not fence is now wrong:

| map site | claim | at `0dab54d` |
|---|---|---|
| `:75` storage, local column | `live-standalone` | **false** — `profiles: [storage-legacy]`, not started by the default profile |
| `:75` storage, citation | `docker-compose.yml:90` | **resolves to an unrelated line** — inside `backend`'s `volumes:` block (`- $HOME/.aws/credentials…`), which does not mention storage |
| `:76` messenger | *"not started by the default `graphql` profile"* | profile renamed `core`; messenger also removed from `all` |
| `:76` messenger, citation | `docker-compose.yml:178` | **resolves to `REDIS_STREAMS_INDEX=4`** |
| `:180` | *"v9.0 `storage` + `messenger`, PR #1103 **open**"* | its compose half has **landed** |

So **the gate reads 1 of 5 at origin HEAD on the clauses examined this iteration**, not the 2 of 5 TOK-04
booked — clause 4 holds, clause 3 does not, clauses 1 and 2 are unrestored. This is a downgrade **produced by
looking**, which is the whole point of P3; nothing regressed, and the same three commits' worth of prose drift
was already routed.

**The map is NOT repaired here, deliberately.** Repairing five prose sites on one seat's reading is precisely
what `CHECK-M257x-iter52-second-ai-manager` forbids and what iter-54 did thirty minutes before catching
itself. It routes to iter-56, which TOK-04 already scoped as the derived-and-fenced prose sweep, with
TOK-03 move 4's two blind pre-commit readers — the instrument this repair would need and this iteration
cannot convene. **What is delivered instead is the honest gate reading**, which is worth more than a
five-line edit taken on my own authority.

**And the sharper lesson, measured twice in one iteration:** the derived and the fenced artefacts survived
the fold untouched (clause 4's derivation, the membership guard), while every hand-written statement of the
same facts — the profile literal, `verify_svcs`, `INJECT_SVCS`, and now five map rows — was falsified by a
single commit. That is TOK-04's P4 ordering reproduced on a second, independent event.

## Measurement 2 — clause 1, cycle A at origin HEAD: **RED**, and the cause is not ours

The first cold cycle that ran to completion at `0dab54d`, after the teardown fix.

```
refs:
  platform:  0dab54dfac6beacdef54a671e2500d3940fd7329   (origin/main; re-checked at close, still HEAD)
  app clone: v1.363.2  (5ba17044, 2026-07-31 — the demo's pinned build ref)
  rext:      fast-build-m257x-iter-55b                  (consumption clone; tag on origin)
  rosetta:   c2de4ffd75aa3a2f5faaef68fae5482ad44db55a   (+ this iter's work)
  taken:     2026-08-03, cold `down --purge` -> `up-injected.sh 1`
  verdict:   autoverify — 3 check(s) FAILED
```

**TOK-04 pre-registered clause 1 green. It is refuted.** Three failed checks, all one root cause:

    ⚠ backend /api/health did NOT answer 200 on :18082
    ✗ backend          HTTP 000000 (unexpected)
    ⚠ container demo-1-backend-1 is NOT RUNNING (status=exited exit=0 restarts=0 — a CLEAN exit(0))

**Everything else went green**, which is what makes the reading informative rather than a wall of noise:
`sentinel.casbin_rules = 1251` · `directus.directus_collections = 21` · Directus per-stack-local (not prod)
· demo-patches all applied, none refused, none skipped · frontend builds are this run's · `public.skills =
42790` · presenter cockpit answering · Clerkenstein fake-FAPI answering · hiring org set-dressed (5 shared
positions, 42 candidate sessions) · academy catalog rendering · studio-desk AI key present. **The re-point
worked**: the bring-up ran under `-p demo-1, core profile — derived from the platform compose`, injected
`app` only (`derived from the platform compose's build set: sentinel app`, skipping the two folded
services), and the whole set-dress + seed pipeline completed.

### Root cause, cited

`backend` exits **0**, silently — the shape autoverify already names as *"the M256 self-termination shape
(a service shutting itself down on its own DB-health monitor)"*. Its last words are two `DB health monitor
shutting down` lines and nothing else. Measured, not inferred:

| probe | result |
|---|---|
| `docker inspect … .State.ExitCode` | `0`, no `.State.Error`, not OOM-killed |
| backend container env | **`STORAGE_RPC_ADDR` is ABSENT** (`0dab54d` deleted it; `STORAGE_S3_BUCKET`/`_PUBLIC_BUCKET`/`AWS_REGION` are present in its place) |
| the pinned app source | `v1.363.2` **reads it three times** — `main.go:446`, `main.go:516`, `main.go:983` (`storage.NewClient(os.Getenv("STORAGE_RPC_ADDR"), …)`) |
| app tags on origin | `v1.365.0`, `v1.364.1`, `v1.364.0` all newer than our pin |

**The compose half of v9.0 `support-in-app` has landed; the app half is not in the release the demo pins.**
`0dab54d` removes the variable *because* `app` now serves storage in-process — true of the app on the
cutover branch, not of `v1.363.2`. So this is **Trap D** (`platform-alignment.md` §4 — *the platform ships
coordinated multi-repo changes*) with the two halves landing on different clocks, and it is the **exact
failure §9's watch bullet predicted in advance**: *"because their env flags are being deleted rather than
defaulted, anything reading them will break rather than degrade."* It broke. It also did so **silently, with
exit 0**, which is worse than the predicted break.

**This is not a tooling defect and there is no rext fix for it.** The remedy is to advance the demo's `app`
pin to `v1.365.0` — which §7 rule 4 says must be done **deliberately, recording what the advance contains**,
and which is the same move that broke the seeders at v2.1 and again at v2.7. It is a decision, not a
reflex, and it is routed rather than taken here.

### §5 rule 15 — the path this cycle took

Recorded because the Directus bootstrap race is nondeterministic and a battery of greens can certify a
branch that was never entered. **Cycle A took the fresh-bootstrap path**: `CREATE SCHEMA directus` on an
empty cluster (the purge really did purge this time) → `node cli.js bootstrap` → structure auto-provision →
replay → boot, ending in `[directus] 'demo-1-directus-1' serving /server/health (verified before
autoverify)`. **No race was observed and no retry fired.** One cycle, one path — which is exactly why one
cycle is not clause 1.

## Clause 2 — the binding reading was NOT taken

It requires the Playthrough suite against the cold stack, and that stack's `backend` is down. Running it
would produce 30 failures attributable to a known-dead container, which is not a measurement of anything.
**Clause 2 remains unrestored**, and the control at the top of this document is not a substitute.

## Close — 2026-08-03

**Outcome:** the ref baseline is re-established and three hand-maintained topology tuples are replaced by one
ref-independent derivation; the first cold cycle at origin HEAD found a teardown that removed nothing (fixed,
proven live) and then a **RED clause 1** whose root cause is a platform-side version skew, cited to
`main.go:446/516/983`. TOK-04's pre-registration of *both green* is **refuted, on both clauses**.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: y — (5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-4

**Gate reading at close, against platform `0dab54d` (re-checked at close; origin has not moved):**

| clause | reading | basis |
|---|---|---|
| 1 — 3 cold cycles green | **NOT MET** | cycle A RED, root-caused; cycles B/C not attempted |
| 2 — full Playthrough suite | **NOT MET** | control green at a 3-fold-stale ref; binding reading impossible with backend down |
| 3 — the migration-status map | **NOT MET** | membership fence GREEN; **5 map claims falsified by `0dab54d`**, two citations resolving to unrelated lines |
| 4 — zero writes to a dropped schema | **MET** | 35 tests OK, run against the `0dab54d` `repos.yml` |
| 5 — KB-fidelity | NOT MET | untouched this iteration; not re-cut |

**1 of 5.** The booked figure was 2 of 5. **Nothing regressed** — the difference is that clause 3 was
*looked at* rather than inherited, which is what P3 exists to force.

**Decisions:** D-M257x-55-1 … D-M257x-55-7 (`iter-55/decisions.md`)

**Side-deliverables:**
- `platform-alignment.md` §8 gains *"A TEARDOWN is a write path too"* — three rules, plus the
  not-an-orphan-and-not-selected corollary. Per the protocol-evolution rule, in the iter's own commit.
- `test_studio_acquisition_m257`'s live fence repaired (it began skipping, silently, on my own change).

**Routes carried forward:**
- `FIX-M257x-iter55-app-pin-lag` → **iter-56, or a dedicated cycle.** Advance the demo's `app` pin
  `v1.363.2` → `v1.365.0` (the v9.0 app half) and re-run cycle A. §7 rule 4: record what the advance
  contains. **This blocks clauses 1 and 2** — neither is measurable until it lands.
- `FIX-M257x-iter55-map-storage-messenger` → **iter-56**, folded into the derived-and-fenced prose class.
  5 sites in `platform-migration-status.md` (rows `:75`, `:76`, `:180`), two of them citations that now
  resolve to unrelated lines. Not repaired here on one seat's reading.
- `FIX-M257x-iter55-stranded-demopatch-revert` → 3 manifests whose whole-file shas are stale against the
  demo clone. Non-blocking (measured: `G4 idempotent no-op` on the very next run).
- `CHECK-M257x-iter55-map-prose-unfenced` → the membership guard says in its own header that it fences
  membership, not prose, and clause 3's wording asks for **cited claims**. Either the guard grows a
  citation-resolution clause (the `anchor_construct_guard` already does this for the corpus and would have
  caught both dead citations) or clause 3 is knowingly half-fenced.
- **Cycles B and C of clause 1**, once the pin lag is resolved.

**Lessons:**
1. **A teardown is a write path, and nobody had ever checked one.** Every rule in the protocol pointed at
   the way in. The way out was where a stale topology statement did the most damage — it deleted a live
   database out from under eleven running containers and exited 0.
2. **Derivation beat prose again, on a second independent event.** The derived (clause 4) and the fenced
   (membership) survived the fold with zero human action. Every hand-written statement of the same facts —
   profile literal, `verify_svcs`, `INJECT_SVCS`, five map rows — was falsified by one commit. TOK-04's P4
   ordering is now measured twice, on two different platform commits.
3. **A refuted pre-registration was the most valuable output available**, exactly as TOK-04 said it would
   be. Two green readings would have told us nothing we did not already believe; the red ones produced a
   fixed teardown, a named platform skew, and an honest 1-of-5.
4. **The watch bullet worked and it was not enough.** §9 predicted this failure *by name* — "their env
   flags are being deleted rather than defaulted, anything reading them will break rather than degrade" —
   and it still cost a full cold cycle to discover, because nothing checks a prediction against the pinned
   build. A watch item with no fence is a note.
