**Type:** tik — under `TOK-01` (measure the composition before engineering it).

# iter-17 — prove it live: the policy fix, on a demo stack and on dev

## Phase A — re-pin, and complete the half iter-15 tripped on

`stack-demo/rosetta-extensions` re-pinned `fast-build-m258-iter-14` → **`fast-build-m258-iter-16`**,
and the **declaration** fixed with it: `.agentspace/rext.tag` still read `fast-build-m258-iter-09`,
which is precisely the mismatch that made the M217 FATAL guard fire at iter-15 (`D71`). Both halves
now agree. `demo-stack/stacks/` (demo-1 … demo-4, registry) survived the checkout untouched, as
`D23` predicted — it is gitignored.

Two incidental findings, recorded not chased: the stray
`stack-demo/rosetta-extensions/.agentspace/rext.tag` (pinning iter-09) is **not** the SoT — `rext_tag`
reads `$REPO_ROOT/.agentspace/rext.tag` — and `demo-stack/stacks/registry.json` is **`{}`** while
`demo-3` is up, so the unified registry `/stack-list` reads does not know about the running stack.

## Phase B — the first bring-up, and what the defaults did

Cold `up-injected.sh 4`, every default. It built, booted 10 containers, set-dressed, and then:

```
BATCH GATE: SKIPPED — demo-4 is published on marcos-mac-mini.taildc510.ts.net (--public-host).
```

`--public-host` is **default-on** and the batch needs the localhost path, so a bare `/demo-up N`
**does not drive the batch on its own host** — `D84`. The skip is right (the MagicDNS origin is baked
into the frontend build; browsing it from the host bypasses `tailscale serve` and every GraphQL call
dies on TLS), and `PT_HOST` cannot rescue it for the same reason. Re-run needed with
`--no-public-host`.

That teardown-and-retry paid for itself: **`down -v` proven live** (`D86`), dangling volumes **9 → 9**
across a full 10-container `--purge`.

## Phase C — the proof

Cold `up-injected.sh 4 --no-public-host`, started **11:43:03Z**, `macmini`, `load1` ~8.7 during the
build lanes, 179 GiB free, `demo-3` resident throughout.

**All three fix sites fired**, each printing `✓ policy invalidated via redis pub/sub (in-process
enforcer reloaded)` where `demo-3` printed the *"non-fatal"* warning: after set-dress (`:404`), after
reset-to-seed (`:516`), after the restore (`:868`).

**The verdict:**

```
BATCH GATE: GREEN — red set EMPTY on demo-4.
UP, and every journey verified.
```

`red_count: 0` · `red_set: []` · `runner_exit: 0` · `passing: 30 / failing: 0 / unimplemented: 1` ·
`215 passed (2.1m)` · `autoverify green: true, warnings: 0`. Full table + the negative-control
argument in [`decisions.md`](decisions.md) `D82`.

**15 → 0.** And `batch_seconds` **629 → 129**: the old batch was slow *because* it was broken — a
refused query costs a 20–60 s timeout, a granted one ~1 s.

The prediction written before the run scored **3 of 4** (`D83`). The miss is the informative one:
`pt-hiring-recruiter-compare` went **green** and `hiring-funnel` seeded **50 rows against demo-3's
38**, so `FIX-M258-iter15-hiring-under-set-dressed` **does not reproduce** and is re-scoped from
*fix* to *not-reproducible*.

## Phase D — the dev half

Six pieces × three questions (does it apply · is it wired · if not, is that correct) — the full table
is `D85`. **Five were correctly demo-only; one was a real gap.**

`dev-stack`'s teardown ran `compose down` **without `-v`** while reaching the *same* bitnami postgres
through the *same* platform compose — so every `dev-N` teardown orphaned two anonymous volumes, the
exact producer measured at **178 dangling volumes / 5.297 GB**. iter-14 fixed the demo half and
stopped there. Fixed, safety argument re-derived against the dev compose, fenced four ways including
a twin-fence. 159 dev-stack tests green.

The two "correct by constraint" rows carry a **reportable consequence**: a dev stack's UI images stay
**pre-L1 fat** (next-web's 4.04 GB → 417 MB win is demo-only), because closing it would mean editing
the platform's Dockerfiles.

## Phase E — converge

`demo-4` down + purged; **`demo-3` untouched throughout** — 10 containers, never restarted, reseeded
or reset, verified at every step. Second `down -v` data point: **9 → 9** dangling volumes.
`END-M258-one-stack` re-established. The caveat about `demo-3` having been built by the *buggy*
tooling is `D87` — it is the user's call, and it is in the report.

## Close — 2026-08-12

**Outcome:** **The iter-16 fix is PROVEN END-TO-END: 15 reds → 0, cold, on a fresh `demo-4` from the
newest platform mains — `red_set: []`, `runner_exit: 0`, 30/31 passing (the 1 is the declared TODO),
`215 passed`, `autoverify green: true, warnings: 0`, "UP, and every journey verified."** All three
invalidation sites fired in one run; both previously-failing negative controls — the tests that
distinguish *correctly isolated* from *uniformly blind* — now pass. The dev half found **one real gap
in six pieces**: the anonymous-volume leak fixed for demo at iter-14 was never carried to dev, which
runs the same image through the same compose; fixed and fenced. Two defaults-level findings recorded
rather than changed: **`--public-host` is default-on and turns the batch gate off on its own host**,
and **`FIX-M258-iter15-hiring-under-set-dressed` does not reproduce** (50 rows vs 38).
**Type:** tik
**Status:** closed-fixed
**Gate:** N/A — closed by user ruling (`D52`). **Clause 3 remains NOT MET and is not recorded as
met**: this iter's ~9-minute cycle ran on a **warm** cache with a second stack resident, which is not
a clean clause-3 measurement and is not offered as one. The clause-3 waiter stays disarmed (`D72`).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n
*(the `demo-3` caveat and the two defaults findings are report items, not mid-iter questions that
change what code lands)* — (5) cap-reached: n *(2 tiks)* — (6) protocol-stop: n — (7)
budget-exhausted: **y** *(between iters, tree clean — two bring-ups, a batch and a full dev audit
spent the session; the remaining routes each need fresh host time)* — Outcome: **exit-7**

**Decisions:** D82–D87

**Side-deliverables:** the `.agentspace/rext.tag` declaration corrected to the consumed tag (the
`D71` half-re-pin class, closed rather than re-tripped).

**Routes carried forward:**

- **`CORPUS-M258-iter16-sentinel-in-app`** (the big one, unchanged and now more urgent) — **sentinel
  is the 8th service merged into `app`**. `CLAUDE.md` still calls it Tier-1 always-on and still calls
  `backend → sentinel` *"the only cross-process Connect-RPC edge left in a local stack"*. Also
  undocumented: the Redis invalidation channel, and the compose refactor found this iter — the
  platform compose now `include:`s **`common.yml`**, where `postgresql` and `redis` actually live.
- **`REPORT-M258-iter17-public-host-default-skips-the-batch`** (net-new, `D84`) — a bare `/demo-up N`
  does not drive the batch on its own host. Deliberate v2.3 design; the *interaction* with the M258
  gate is what is new. **User decision, not a sub-agent's.**
- **`REPORT-M258-iter17-dev-ui-images-stay-pre-L1-fat`** (net-new, `D85`) — the release's largest
  image win is demo-only, and closing it would require platform-repo edits.
- **`ROUTE-M258-iter17-batch-gate-has-no-dev-opt-in`** (net-new, `D85`) — correct that it is not
  default-on for dev (the batch opens with a full-world `TRUNCATE`); a gap that there is no opt-in.
  Also: no `hiring-app` on the dev compose, so 1 of 30 Playthroughs could not run there today.
- **`ROUTE-M258-iter17-registry-is-empty-while-a-stack-is-up`** (net-new) — `stacks/registry.json` is
  `{}` with `demo-3` running, so `/stack-list` cannot see it.
- **Re-scoped:** `FIX-M258-iter15-hiring-under-set-dressed` → **`WATCH-M258-hiring-under-set-dress-not-reproducible`**.
- Unchanged, re-verified **open**: `FIX-M258-iter14-purge-leaves-276MB` ·
  `TARGET-M258-iter13-browser-only-deps` · `SETTLE-M258-iter13-studio-desk-cold-time` (`D75` — still
  needs a `--no-cache` A/B on a quiet box; **not** attempted here, and the original estimate stays
  refuted) · `ROUTE-M258-iter13-dockerfile-not-in-cache-key` ·
  `ROUTE-M258-iter15-compose-down-cannot-parse-an-older-stack`.

**Lessons:**

- **Write the prediction before the run, and score it afterwards.** Three of four held; the fourth
  failing is what turned a "known defect" into "not reproducible" instead of into a wasted fix.
- **A skipped gate is not a passed gate.** The first bring-up was green in every visible way and had
  quietly not run the thing the iter existed to measure. Read what the run *declined* to do.
- **Ask the sibling path the same question, always.** Five of six pieces were correctly demo-only,
  which is exactly why the sixth survived three iterations of nobody looking.
- **A slow suite can be a broken suite.** 629 s → 129 s came entirely from not waiting on refusals.
  Timing regressions and correctness regressions are not separate diagnoses.
- **Converge, and say what you left behind.** One stack is the end state; that the survivor was built
  by the buggy tooling is the user's to decide, and stays visible rather than tidy.
