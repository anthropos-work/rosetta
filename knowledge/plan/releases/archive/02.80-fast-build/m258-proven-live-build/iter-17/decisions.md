# iter-17 — decisions

## D82 — The fix is proven END-TO-END. 15 reds → 0, cold, on a fresh stack.

`PROVE-M258-iter17-policy-fix-on-demo-and-dev` is **closed on the demo half**. A cold `demo-4`,
built from the newest platform mains with `rosetta-extensions` at `fast-build-m258-iter-16`:

| artifact | reading |
|---|---|
| `batch-gate.json` | `verdict: green` · **`red_count: 0`** · `red_set: []` · `runner_exit: 0` · `total: 31` |
| `last-report.json` | `passing: 30` · **`failing: 0`** · `unimplemented: 1` · coverage **96.77 %** |
| playwright | **`215 passed (2.1m)`** |
| `autoverify.json` | **`green: true, warnings: 0`** |
| bring-up | `BATCH GATE: GREEN — red set EMPTY on demo-4.` → *"UP, and every journey verified."* |

The single non-passing entry is the **declared in-manifest TODO**, which the corpus already documents
as the standing shape (*30 live Playthroughs + 1 verdicted TODO*). It is not a red.

**All three fix sites fired in one run** — `up-injected.sh` after set-dress (log `:404`),
`run-playthroughs.sh` after reset-to-seed (`:516`), `restore-presenter-world.sh` after the restore
(`:868`), each printing `✓ policy invalidated via redis pub/sub (in-process enforcer reloaded)`. On
`demo-3` the same three points printed the *"non-fatal — a non-AI-sim run is unaffected"* warning.

**The negative controls are the sharpest confirmation.** Both that failed on `demo-3` now pass —
*"another ORG's seeded workforce data is ABSENT from a different tenant's dashboards"* (21.5 s) and
*"a hiring org's shared-position board is ABSENT from a Workforce tenant's activity grid"*. They
assert a manager sees her **own** tenant and not another's; with every org-scoped read refused, the
own-tenant half could not pass. They are the tests that distinguish *"correctly isolated"* from
*"uniformly blind"*, and they were blind.

`batch_seconds` fell **629 → 129** and `restore_seconds` **29 → 7**, because a refused query costs a
20–60 s timeout and a granted one costs ~1 s. The old batch was slow *because* it was broken.

## D83 — The prediction was written before the run, and one quarter of it was wrong.

`overview.md` committed to four falsifiable predictions before the stack existed. Scored honestly:

| # | prediction | outcome |
|---|---|---|
| 1 | the 15 org-scoped reds go green | ✅ |
| 2 | both negative controls go green | ✅ |
| 3 | the 15 user-scoped passes stay green | ✅ |
| 4 | `pt-hiring-recruiter-compare` **stays red** on the independent under-set-dress defect | ❌ **it went green** |

**#4 is the useful one.** `FIX-M258-iter15-hiring-under-set-dressed` did **not reproduce**:
`demo-4` seeded `hiring-funnel rows=50` against `demo-3`'s **38**, cleared the ≥40 floor, and
`autoverify` recorded **0 warnings** where `demo-3` recorded 1. Same seeder, same preset, same
platform ref — so **38 was a condition of that stack, not a defect in the code**, which is exactly
what the warning's own guidance offered as its first cause (*a starved `SIMULATION_TYPE_HIRING`
pool / cold snapshot cache*). The item is re-scoped from *fix* to **not-reproducible**; it must not
be "fixed" on the strength of a single starved observation.

## D84 — `--public-host` is default-ON and it turns the batch gate OFF on its own host.

The first `demo-4` bring-up (bare `up-injected.sh 4`, i.e. every default) reported:

```
BATCH GATE: SKIPPED — demo-4 is published on marcos-mac-mini.taildc510.ts.net (--public-host).
```

The skip is **correct** — a `--public-host` demo bakes the MagicDNS origin into the frontend build,
and `docker-proxy` binds `0.0.0.0`, so a connection from the demo host to its own tailscale IP
bypasses `tailscale serve` and every GraphQL call dies `ERR_SSL_PROTOCOL_ERROR` (M255 spike (e)).
Running the suite anyway would produce 31 reds that mean nothing.

But the consequence deserves naming, because it is the milestone's own gate clause: **`/demo-up N`
with no flags does not drive the batch on this host.** `--public-host` is default-on (`D-DESIGN-3`,
v2.3 M220) and the batch needs the localhost path, so *"one cold command brings the stack up AND
drives the full Playthrough batch"* requires **`--no-public-host`** here, or a tailnet peer. The
milestone's `exit_gate` anticipated exactly this (*"may need a peer or `--no-public-host`"*) — this
is that prediction coming true, measured. Recorded, not fixed: the default is a deliberate v2.3
design decision and changing it is a user call, not a sub-agent's.

## D85 — The dev audit found one real gap in six pieces, and it was the one nobody would have guessed.

Each piece of this release's new functionality, asked all three questions rather than given one
verdict for the set:

| piece | applies to dev? | wired? | correct? |
|---|---|---|---|
| **batch gate** | **yes** — `dev-N --inject` gives the Clerkenstein seat-switch the actor needs | no | **correct by DEFAULT, gap as an OPTION** — see below |
| restore-presenter-world leg | only if the batch runs | no | **correct** — it exists solely to undo the batch's reset |
| studio-desk multi-stage image | yes (dev builds it under `PROFILE=all`) | no | **correct by constraint** — the Dockerfile is rext-owned *because* rext may not edit the platform's; consequence below |
| L1 multi-stage next-web / hiring | same | no | **correct by constraint**, same reason |
| **`down -v` volume fix** | **yes — identical image, identical compose** | **NO** | ❌ **GAP — closed this iter** |
| TOK-02 space classes | host-level | n/a | operator procedure, not path-specific; the durable half is the producer fix above |

**The gap.** `dev-stack`'s teardown ran `docker compose -p "dev-$n" down` with no `-v`, while the dev
path reaches the **same** `bitnamilegacy/postgresql` through the **same** platform compose
(`docker-compose.yml` `include:`s `common.yml`, whose `postgresql` binds one of the image's three
`VOLUME`s). So dev orphaned two anonymous volumes on **every** teardown — the exact producer iter-11
measured at **178 dangling volumes / 5.297 GB**. iter-14 fixed the demo half and stopped, because the
demo path is where the measuring happened. Fixed, with the safety argument re-derived against *this*
compose (**no top-level `volumes:` key in either file** → no named volumes to lose; the DB is a host
bind mount `./data/postgresql`), fenced four ways including a twin-fence so a future fix cannot land
on one half again. rext `fast-build-m258-iter-17`, **on origin**.

**The batch-gate nuance, stated properly.** It is *not* impossible on dev — `--inject` exists. It is
correct that it does not run by **default**, because the batch opens with `stackseed --reset`, a
`TRUNCATE … CASCADE` of the whole world: fine for a disposable demo, destructive for a developer's
working stack. What is missing is an **opt-in** path, and there is also no `hiring-app` on the dev
compose (the platform declares four services; the demo adds it), so 1 of the 30 Playthroughs could
not run there today. Routed, not built.

**The reportable consequence of the two "correct by constraint" rows:** a dev stack's UI images stay
**pre-L1 fat**. The release's largest image win (next-web **4.04 GB → 417 MB**) reaches the demo path
only. Closing that would require editing the platform's Dockerfiles, which is the one line this
release never crosses — so it is a **finding for the user**, not a defect to fix.

## D86 — `down -v` proven live, twice, on the demo side.

Two full `--purge` teardowns of a 10-container stack this iter, dangling volumes measured either
side: **9 → 9 → 9**. Zero orphans across 20 container removals. Before the fix the same pair of
teardowns would have left four. Measured, `macmini`, 11:42Z and 11:53Z, 179 GiB free.

## D87 — Converged to one stack, and the survivor is the user's — with a caveat that must be said.

`END-M258-one-stack` is re-established: `demo-4` torn down and purged, **`demo-3` untouched
throughout** (10 containers, never restarted, never reseeded, never reset — verified at every step).

**The caveat, because it is the user's decision and not mine:** `demo-3` was built by tooling that
carries this bug. Its presenter world answers correctly *now* — the enforcer holds the policy
iter-16's live probe reloaded at 11:26:44Z — but the stack has no automatic invalidation on any
future re-seed, so **anything that re-seeds `demo-3` will silently return it to the all-`forbidden`
state**. `demo-4` demonstrated that a stack built with the fixed tooling comes up green end-to-end in
one command. Re-creating `demo-3` from the fixed tooling is that one command; it was **not** done
here because tearing down the user's stack is explicitly out of bounds.
