# Bring-up Verification — the auto-verify safety net

**The authoritative statement of how a stack proves itself *working* — not just *started*.** Every
bring-up now ends with an automatic, scoped, **non-fatal** verification pass: a couple of decisive
cheap-win asserts followed by the full black-box probe set, targeted at the stack's **own offset ports**
and scoped to the **services actually brought up**. So when a bring-up says "UP", it means
*verified-working* — not merely *containers-started*.

> **Scope.** This doc covers the **v1.3b "dress rehearsal" / M18** verification net across the
> `rosetta-extensions` stack tooling: the offset/project/scope awareness added to `stack-verify`, the
> `$DEVDIR → $STACK_ROOT` bugfix, the cheap-win `/api/health` + `casbin_rules > 0` asserts, and the
> auto-wired scoped `verify live` at the tail of every bring-up (demo + dev). It is the
> *is-this-stack-actually-working* companion to [`rosetta_demo.md`](rosetta_demo.md) (the demo lifecycle +
> the unified registry that records the ports it targets), [`idempotency.md`](idempotency.md) (the
> *what-happens-on-a-re-run* contract), and the [`/test-platform`](../../.claude/skills/test-platform/SKILL.md)
> skill (the operator-driven, deeper verification surface this auto-run is a default-on, non-fatal subset of).
>
> All the code cited lives in the gitignored `rosetta-extensions` monorepo (authored + tagged in the
> authoring copy at `.agentspace/rosetta-extensions/`, consumed per-stack at a pinned tag) — **no platform
> repo is modified.**
>
> **In scope:** the backend `graphql`-profile services (what exists today). **Out of scope:** the
> frontend tier — the frontends don't exist in the stack yet; **M19** adds them and extends the verify
> service list. Deep behavioural / e2e flows remain the operator-driven `/test-platform` job; this
> auto-run is the always-on *smoke net*.

## For PMs — what "verified, not just started" means

Bringing up a demo or dev stack runs a dozen containers and a database migration. Before M18, the tooling
declared "UP … is live" the moment the containers started — even if the stack was actually broken. That
bit us for real: a stack once came up reporting success while its security-policy table had silently
failed to load, which would have made **every** logged-in page return "access denied". Nothing caught it
until someone hit the wall manually.

M18 closes that gap. Every bring-up now ends by **checking the stack actually works**: it confirms the
backend answers a health check and that the security policy loaded, then runs the full probe set across
the services. Crucially, this check is **non-fatal** — if it finds a problem it shouts loudly and tells
you how to dig in (`/test-platform N`), but it never *blocks* a stack that's genuinely fine. The result:
"UP" now carries a real promise, and the failure that bit us would have been caught at bring-up time, in
seconds, automatically.

## The contract, in one paragraph

At the tail of `demo-stack/up-injected.sh` (demo) and `dev-stack`'s `cmd_up` (dev), the bring-up invokes
`stack-verify/live/autoverify.sh --project <demo-N|dev-N> --offset <N×10000>`. That wrapper (1) runs the
**cheap-win asserts** on the stack's own offset ports, then (2) runs the full **offset/project/scope-aware
`verify live`**. It is **default-on** (opt out with `DEMO_NO_VERIFY=1` / `DEV_NO_VERIFY=1`) and **always
exits 0** — a failing check produces a loud `⚠` block and a "run `/test-platform N` to dig in" hint, never
an abort (#M18-D3). This mirrors the proven default-on + non-fatal pattern of `dev-setdress.sh`.

> **Host-native daemons outlive the bring-up task (FIX B).** The two host-native surfaces a demo brings up
> (ant-academy + the presenter cockpit) are now launched **session-detached** via
> `demo-stack/detach.sh::launch_detached` (`setsid` where present; a portable `python3 os.setsid`
> double-fork on macOS, which has no `setsid`), so they survive the launching session/task ending and a
> later visit still finds them alive. Previously a bare `nohup` left them in the launcher's process group,
> so a backgrounded `/demo-up` task's reaping took them down with it.

## The offset/scope model (why it targets the *right* ports)

A `demo-N`/`dev-N` stack publishes its host ports at **base + N×10000** (the offset engine in
`stack-core/gen_override.py`; see [`rosetta_demo.md`](rosetta_demo.md) § the unified registry). Before
M18, `stack-verify` hardcoded the **main dev stack's** `anthropos-*-1` containers at **base** ports, so:

- against an offset `demo-N`/`dev-N` it reported **everything `down`** (a wall of false negatives), and
- against a **reduced** bring-up (e.g. `--services "postgresql redis"`) every absent service was a
  **false `down`** too.

M18 teaches the probes three things, all in `stack-verify/lib/target.sh` (the new resolution helper sourced
by `lib/services.sh` + `lib/readiness.sh`):

| Concept | Env var | Resolution |
|---|---|---|
| **Project** | `STACK_PROJECT` | `demo-N` \| `dev-N` \| `anthropos` (default). Rewrites the container prefix: `anthropos-cms-1` → `demo-3-cms-1`. |
| **Offset** | `STACK_OFFSET` | Honoured if set; else **derived** from the project's `N` (`demo-3` → `30000`). Added to every host port: `8082` → `38082`. |
| **Scope** | `STACK_SERVICES` | Space-separated service names to probe (∩ the SERVICES array). Empty = all-in-profile. A service not in the set is **skipped**, not false-`down`ed. |

### The correctness mitigation (the load-bearing bit)

A **mis-derived offset** would report a healthy offset stack as `down` — the *exact* false-positive bug
M18 exists to fix. Two guards prevent that:

1. **Derive from what's known, cross-check against what's recorded** (#M18-D1). The bring-up passes
   `--project` + `--offset` explicitly (it allocated `N`, so it *knows* them — no drift). For
   operator-driven runs, `target_cross_check()` reads the unified registry's **recorded** host ports (M12
   records resolved ports per stack) and **warns** (non-fatal) if the derived offset doesn't land in the
   platform's base-port band — `(port − offset) ∈ [3000, 11000]`, which covers all 12 bases without a
   hardcoded table and without the broken `port // 10000 == n` decade assumption that would false-warn on
   roadrunner's high base (#M18-D5). The registry record — not a bare re-computed formula — is the source
   of truth.
2. **Non-fatal, always.** Because `autoverify.sh` always exits 0, even a verify/offset *bug* can never
   block a genuinely-good bring-up. The worst case of a wrong offset is a spurious warning, never a
   refused stack.

## The cheap-win asserts (the ISSUE-7 catcher)

Before the full probe set, `autoverify.sh` runs two decisive, dependency-free checks on the stack's own
offset ports — the precise failure class that shipped a "live" but blanket-403 stack:

1. **Backend health** — `curl -fsS http://localhost:$((8082+OFFSET))/api/health`. The web API actually
   answers. (Skipped if `backend` isn't in scope.)
2. **Authz policy loaded** — `SELECT count(*) FROM sentinel.casbin_rules` via `docker exec
   <project>-postgresql-1 psql …`, asserted `> 0`. An empty `casbin_rules` means the Sentinel policy
   never loaded → every authorized route 403s. This is the exact silent failure that bit us; the assert
   surfaces it at bring-up time. (Skipped if `sentinel` isn't in scope.)

A failed assert increments the warning count and prints a `⚠` line; it never aborts.

### The per-stack Directus cheap-wins (M22 — the same class)

On a stack brought up with **local content** (demo default; dev `--local-content`), `autoverify.sh` runs two
more cheap-wins — gated on the directus **container actually existing**, so a prod-read stack (no local
Directus) never false-warns even on an unscoped run:

3. **Directus REGISTERED the content model** — `SELECT count(*) FROM directus.directus_collections` via
   `docker exec <project>-postgresql-1 psql …`, asserted `> 0`. The silent-failure analog of the casbin
   assert: a Directus can be UP (`/server/health` 200) and have no content model at all. (Also runs as the
   `directus-collections` readiness probe.)

   > **⚠ This assert used to be described — here and in its own `✓` line — as *"Directus serves the
   > catalog"*, and it cannot know that.** `directus_collections` is a **registry** table written by the
   > *structure* replay; the content rows are loaded by a later step and the anon read grants by a third, so
   > the count is `> 0` on a Directus holding zero items and on one that 403s every read. On M257x clause
   > 1's three cold cycles it was exactly that: the content replay raised on `directus.sequences`, **0 of
   > 11986 rows** landed, every anon `/items` read returned 403 — and `autoverify` graded `green:true /
   > 0 warnings` three times over it, because nothing anywhere in the verify path asked the *running*
   > Directus for an item. Corrected in M257x harden pass 1; the rule it produced is
   > [`platform-alignment.md`](platform-alignment.md) §5 rule 14 — **REGISTERED is not SERVED.**

3b. **Directus actually SERVES an item** (`directus-serves-content` readiness probe, M257x harden pass 1) —
   the measurement the count above cannot make. It asks the **running** Directus over HTTP, on the stack's
   own offset port, **unauthenticated**, for one item out of a collection it did not choose: the target is
   **derived** from the stack's own catalog (the non-system `directus.*` table holding the most rows), so a
   re-modelled surface cannot make the probe stale and it cannot pick a target that guarantees its own
   success. The derivation is **fail-closed** — no user collection holding a single row IS the defect, not a
   skip. `403` (holds the content, serves it to nobody), `200` with an empty `data` array (serving, content
   absent) and no-response are each named distinctly, because they have three different repairs. Scoped to
   the `directus` service like the others, so a prod-read stack still skips cleanly.
4. **No prod read** — the per-stack Directus's `DB_CONNECTION_STRING` (read from the container's env) must
   resolve to the stack's **own** Postgres, never a prod host. The runtime mirror of the executed-provision
   firewall gate; warns (non-fatal) if a mis-wired override pointed the local Directus at prod.

> **The boot health-gate (FIX A) — why these probes no longer race.** The set-dress step that restarts the
> per-stack Directus (`dev-setdress.sh::boot_directus_step`) used to `docker restart` the container and
> return **immediately** — so the bring-up-tail autoverify fired while Directus was still ~30s into its
> re-introspect, and the directus liveness (`/server/health`) + `directus-collections` probes raced that
> window and **false-reported "down"** (a transient verify `⚠` on a stack that was actually fine).
> `boot_directus_step` now **waits** for the stack's own offset `/server/health` to answer `200` before
> returning (bounded by `DEV_SETDRESS_DIRECTUS_BOOT_TIMEOUT`, default `90s`; **non-fatal** on timeout or a
> missing `curl`), so autoverify can't run ahead of it. The probes themselves are **unchanged** — the fix
> lives at the restart, not at the probe. (The health-gate fix landed at `storytelling-postfix-1`; the demo
> consumes the tag recorded in `.agentspace/rext.tag` — the single source-of-truth pin, M49 #1.)

## What runs, and on which ports

`verify live` runs two phases over the **selected, offset-resolved** services
(`stack-verify/live/verify.sh`):

- **Liveness** — per service (`lib/services.sh::service_rows` + `probe_service`): docker-health /
  TCP-connect / HTTP-code, at `base+offset`, against `<project>-<svc>-1`.
- **Readiness** — deeper, correctness probes (`lib/readiness.sh`): postgres schemas present, redis
  `PING`, GraphQL introspection (`:8082+offset` at `/graphql/query` — re-pointed at M257x iter-13 when the
  `:5050` router was deleted; a wrong *host* refuses loudly but a wrong *path* connects and 404s, so both
  halves moved), gotenberg version (`:3200+offset`), sentinel
  Connect-RPC handler mounted (`:8087+offset`), storage RPC reachable (`:8301+offset`), and — on a
  local-content stack — the per-stack **Directus** liveness (`/server/health` at `:8055+offset`) plus its
  `directus-collections` registration check **and the `directus-serves-content` probe that reads an actual
  item back over HTTP** (§"cheap-wins" 3/3b above — the registration count alone graded three cold cycles
  green over a Directus that served nothing) — each resolving the offset port + project container via the same
  `target.sh` helpers. **Both** phases honour the `STACK_SERVICES` scope filter: the readiness phase skips a
  deep probe whose backing service isn't in scope (the same `target_service_selected` gate as liveness), so a
  reduced bring-up never produces a wall of false `down`s in *either* phase. (The directus row is scoped in
  only on a `--local-content` bring-up and gated on the container existing — a prod-read stack stays clean.)

The full base-port table (the offset-0 source of truth the offset is applied to) lives in
`stack-verify/lib/services.sh`; the `/test-platform` skill drives the same scripts for the deeper,
operator-initiated `repos` + `census` scopes.

## The `$DEVDIR → $STACK_ROOT` fix

A latent bug rode alongside the verify gap: `stack-verify/repos/run.sh` and `census/inventory.sh` resolved
each platform repo as `$DEVDIR/$repo` — an **undefined** variable (only `$STACK_ROOT` is defined). Every
repo collapsed to `/$repo` and was reported `not cloned`, breaking the `repos` + `census` scopes on the
first run. M18 fixes both to `$STACK_ROOT/$repo` (repos are siblings of `platform/` under the stack root).

## When verify warns — reading the output

A clean run ends `▶ autoverify <project>: OK — verified-working.` A run with problems ends with a loud
block:

```
⚠⚠ autoverify demo-3: 2 check(s) FAILED — the stack is UP but may be non-functional.
   Dig in with:  /test-platform 3 live   (or: STACK_PROJECT=demo-3 …/verify.sh)
   (non-fatal: the bring-up is NOT aborted — investigate, then re-run the failing step.)
```

The stack is still up; the warning points you straight at `/test-platform` for the deeper, scoped probe
and (for the casbin case) the fix is re-running the migrate step (see [`idempotency.md`](idempotency.md) —
migrate is re-run-safe).

### M217 — the warning now NAMES the failing probe

> **It didn't used to.** `autoverify` invoked `verify.sh` as `>/dev/null 2>&1`, which threw away **every**
> `✗ <service> <detail>` line and collapsed N failing probes into exactly **one nameless** warning. The live
> output on `billion` was:
>
> ```
> ⚠ verify live reported failing probe(s) — some service is not serving correctly.
> ⚠⚠ autoverify demo-1: 1 check(s) FAILED — the stack is UP but may be non-functional.
> ```
>
> …with **nothing anywhere naming the service.** (It was `jobsimulation`, dead in a crash loop in *every* demo.)
> **A safety net that fires without saying what it caught is barely a safety net.**

The failing probes are now listed inline, the full transcript is persisted to **`<stack>/autoverify.log`**, and a
non-zero exit with *no* `✗` line is reported as *"the verifier itself may be broken"* rather than blamed on a
service.

### The four cheap-wins verify could not see (M217)

Liveness probes cannot observe a stack that is **up and wrong**. These four are the same shape as the ISSUE-7
casbin assert — seconds to run, decisive:

| Check | Why it exists |
|-------|---------------|
| **a demo-patch was REFUSED** | **This is literally how the perf rot survived four releases.** Both `app` perf patches refused on every run, the reason was piped to `/dev/null`, and *nothing downstream noticed*. The stack was green, looked fine, and shipped a **76-second members grid**. |
| **snapshot replay SKIPPED** (`public.skills = 0`) | A cold cache exits `rc=5` **non-fatally**, so the catalog is empty and the bring-up still prints **UP**. Every skill surface renders blank. |
| **the cockpit isn't answering** | The presenter has **no way in** — and until M217 the bring-up logged *"presenter cockpit serving on …"* **unconditionally**, even when it had just died on a leaked port. |
| **the fake-FAPI isn't answering** | **NOBODY CAN LOG IN** — and verify stayed **green**, because no probe covered it. A demo nobody can log into is not a demo. |

### The hiring set-dress cheap-win (v2.4 "casting call" M225 — the same class)

A fifth demo-only cheap-win, added when the hiring org became a first-class demo surface. The recruiter's
candidate-comparison scoreboard is populated **entirely from the auto-set-dress**: the `HiringConfigSeeder` resolves
**5 real captured `SIMULATION_TYPE_HIRING` sims** (`readHiringSimPool`, off the replayed `directus.simulations` — the
HIRING sims ride along in the **standard** directus content-surface capture; there is **no** separate 5-sim capture and
**no `directus.job_position` replay** — 0 rows, unread by the scoreboard, M222 BA-6 / M223 D4) as the org's
`organization_sim_invitation_links` **positions**, and the `HiringFunnelSeeder` writes each candidate's scored
`local_jobsimulation_sessions` **mirror** row (the score the scoreboard reads).

The silent failure it catches: a **cold snapshot cache** (or a starved HIRING-typed pool) leaves `readHiringSimPool`
**empty** → the seeders **honestly degrade** to 0 positions / 0 sessions (never fabricate) → the recruiter comparison
renders with **no columns and no rows**, while the stack still prints **UP**. M223's adversarial review named the
downstream M224/M226 render gate as the loud catch; this cheap-win brings that catch **forward to the default
`/demo-up` tail** so "the hiring org comes up real with no manual steps" is a **checked** property, not an assumed one.

- **Gated on a hiring org existing** (`SELECT count(*) FROM public.organizations WHERE is_hiring = true` > 0) — so a
  hiring-less demo (a non-stories preset, or `DEMO_STORIES=0`) **skips cleanly**, never false-warns (the same
  discipline as the directus container-presence gate).
- **Floors:** `≥ 5` positions (the shared-positions contract, `reservedHiringSimRefs`; `< 5` is the cold-cache /
  starved-pool "dangerous middle") **and** `≥ 40` candidate `local_jobsimulation_sessions` for the hiring org (a full
  comparable cohort — a robust weak lower bound; the ~200 the funnel seeds is the healthy number, and the **strong**
  per-sim `≥ 40` floor is the M224/M226 render probe, not this cheap bring-up assert).

### `autoverify.json` — the machine-readable signal

```json
{"project":"demo-1","offset":10000,"warnings":0,"green":true,"ts":"2026-07-13T22:41:07Z"}
```

Its path is **`rosetta-extensions/demo-stack/stacks/<project>/autoverify.json`** — *not* `<stack>/autoverify.json`,
as this doc claimed until M218 (**iter-01 F-2/D3**; the tooling that gates on it, `stack-verify/e2e/run-latency.sh`,
reads the real path).

**Its lifecycle is the guarantee** (M218 harden, **F-10**). The file is **unlinked on teardown**
(`rosetta-demo`'s `clear_stack_verdict`, called *first* in `cmd_down`, before any teardown step that could
itself fail) **and again at the start of every bring-up** (`up-injected.sh`). So a verdict on disk can only
have been written by *the run under test*. **Absence is the safe state**: a grader with no verdict refuses to
measure. `ts` is defence-in-depth, not the guard — if some future path ever misses an unlink, the staleness
is at least *visible*.

Until M217, `/demo-up` **exited 0 on a red verify and still printed UP**, so nothing downstream could gate on
*"did this stack actually come up green?"*. **M218 measures login latency and must not measure a broken stack** —
this file is what it gates on.

> #### A green verdict is only as fresh as the image it graded (M218 iter-03 **D9**)
>
> The file records a **stack-at-an-instant**, not a stack-forever. M218 iter-03 found a `demo-1` whose
> `autoverify.json` said `{"green":true}` — written **nine hours** before an **out-of-band `docker build`** swapped
> the `next-web` image underneath it. The stack was, at that moment, **Clerkenstein-dewired** (see below) and
> **nobody could log in** — while the green file sat there asserting otherwise.
>
> ⇒ **Anything that gates on this file must ensure the verdict was written against the image actually under test.**
> A fresh bring-up writes a fresh verdict; trusting the file's mere *existence* re-opens the exact hazard M217 shut.

### Validate a baked constant by reading the artifact it was baked into (M218 iter-03 **D8/F-6**)

A cache-validator can only see the constants it can **read** — and `docker image inspect` sees only image **ENV**.

`up-injected.sh`'s `next-web` image cache-validator (M211) compared the baked `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`,
which *is* an image ENV. But the **minted publishable key** — the constant that **wires Clerkenstein** — is **not**:
it is written to a gitignored `apps/web/.env.local` overlay, consumed by `next build`, and **inlined into the
bundle**. The validator was therefore **structurally blind** to it, and an image carrying the **right offset** with
the **wrong, real-Clerk key** *passed the guard and got silently reused*. Consequences, both real:

- **login is broken** — the browser's clerk-js talks to the **real Clerk app**, finds no session, and loops to
  `/login`; and
- **the demo phones production auth**, which [`safety.md`](safety.md) forbids outright.

The fix asserts the minted pk is **present in the built bundle** (`docker run --rm … grep -rqs "$PK_DEMO"
/app/apps/web/.next/static`), **fail-safe toward rebuild**: anything unverifiable rebuilds, because a needless
~3 min build is strictly cheaper than shipping a real-Clerk-wired demo. All four overlay-borne constants come from
that one file, so a single probe covers the class.

> **The general rule:** *if a constant is inlined into an artifact, validate it by reading the artifact.* Build-time
> inlining (`NEXT_PUBLIC_*` and friends) is invisible to every ENV-level guard — and it was **M218's antagonist
> twice over**: once as C-1 (a build-inlined URL a consumer *cannot override*) and once as F-6 (a build-inlined key
> a rebuild path *failed to supply*).

## THE STALE-VERDICT HAZARD — a status artifact that outlives the thing it describes

> **This is a first-class, named hazard, not an anecdote.** M218 hit the *same failure class* **five separate
> times** — in the tooling, in the demo lifecycle, and in the probes themselves. It had already survived one
> full hardening pass (M217). Fixing the sixth instance in isolation would miss the point: **the class is the
> bug.**

**The shape.** Some artifact records a *verdict* — green, done, absent, empty, purged. Its subject then
changes or dies. The artifact does not. Something later reads it as **current evidence**, and it lies.

**The five instances, all in one milestone:**

| # | The artifact | What it asserted | The truth |
|---|---|---|---|
| **F-6** | `autoverify.json`, 9 h old | `{"green":true}` | An out-of-band rebuild had swapped the image; the stack was wired to **production Clerk** and nobody could log in. |
| **F-9** | `/demo-down --purge`'s **exit code** | *(a bare `rc=1`, unread)* | The purge **deleted nothing**. Postgres's UID-1001/0700 cluster dir defeated `rm -rf`; `set -euo pipefail` then aborted teardown. Every "cold" bring-up for **days** reused a 2-day-old database — including the one that first claimed the exit gate. |
| **F-10** | `autoverify.json`, not unlinked on teardown | `{"green":true,"warnings":0}` | The stack had **zero containers**. The bring-up had *failed*. It still graded green to every reader. |
| — | a `[ -e ]` file test | "absent" | **Permission denied.** Unreadable was read as not-there. |
| — | an `assertNotIn` probe | "the bad string isn't present" | The command had **failed and produced no output**. Absence of evidence scored as evidence of absence. |

**The two invariants that kill it:**

1. **A verdict must not outlive its subject.** Destroy it on teardown *and* at the **start** of every
   bring-up — and destroy it **first**, before any step that could itself fail and abort the sequence
   (exactly how F-9 leaked). Then the artifact's *presence* means "the run under test wrote this," which is
   the only thing a grader can safely conclude from it.
2. **Absence must be the safe state.** A grader with **no** verdict must **refuse to measure** — never
   default to pass. Nearly every instance above is the same mistake in different clothing: *treating
   "nothing here" as "nothing wrong."* An empty result is not evidence of success.

**And the corollary for the checks themselves:** a probe that can pass **without ever executing its
assertion** is a stale verdict in test form. `assertNotIn` on a failed command's empty output, a `[ -e ]`
that cannot read the path, a coverage check nobody wired up — all report green having measured **nothing**.
Assert that the check *ran*: count what you inspected, and fail if the count is zero. (M218's own new
regression tests do this — see `test_verdict_lifecycle.py`'s `checked > 0` guards.)

> **The sibling hazard, same family:** a **safeguard that exists only in prose**. `alignment_testing.md`
> named a "capability-coverage check" as *the* mitigation against a hollow alignment score. There was no such
> check — and the endpoint it would have caught was being called on every authenticated render while the
> mirror scored **100%**. A documented guarantee with no enforcement behind it is a stale verdict about your
> own tooling. See [`alignment_testing.md`](../architecture/alignment_testing.md#the-capability-coverage-check--what-it-actually-guarantees).

## PRE-FLIGHT RUNG ZERO — can the host even OBTAIN the thing under test? (v2.5 M236)

**Before verifying that a stack *works*, verify the stack can *get* the code you think you are testing.**
This rung sits upstream of every check on this page: all of them measure a running stack, and none of them
can tell you the stack is running **different tooling than you believe**.

**The shape.** A milestone whose gate is *"prove feature X live on host H"* implicitly assumes H can obtain
X. When X is delivered as a `rosetta-extensions` tag, that assumption has a step in it that is easy to skip
and invisible when skipped: **tagging is not publishing.**

M236 opened with the entire v2.5 tooling — M230, M232, M233, M234, M235; 20 commits and **13
`playbill-*` tags** — present only in the `.agentspace/rosetta-extensions/` **authoring copy**. `origin/main`
was still the M228 commit, and **zero** of the 13 tags existed on origin. `billion` consumes tooling only at
the tag named in `.agentspace/rext.tag`, so the host could not have obtained the feature under test by any
route. Nothing in the milestone plan named publication as a step, because `CLAUDE.md` and
[`rosetta_demo.md`](rosetta_demo.md) both describe the path as *"built and tested in the authoring copy and
tagged, then consumed per-stack via a pinned-tag clone"* — which reads as though tagging alone makes a tag
consumable. It does not. M230–M235 all ran offline, so the publish half was simply never exercised.

**Why the existing guards do not catch it.** The M217 rext-pin guard
(`demo-stack/ensure-clones.sh`) is FATAL and does exactly its job — it compares the *consumption clone's
checkout* against `.agentspace/rext.tag`. Both can agree perfectly on a **stale** tag. The guard answers
*"is this clone at the tag we pinned?"*, never *"is the tag we pinned the one carrying the work?"*

**The rung.** Before the first bring-up of any *prove-it-live* milestone, assert all four — each is one
command, and any one failing makes every downstream measurement a measurement of the wrong code:

1. The work is committed and **tagged** in the authoring copy.
2. The tag is **on origin** — `git ls-remote --tags origin | grep <tag>`. Local-only is the default failure.
3. `.agentspace/rext.tag` on the **target host** names that tag.
4. The host's consumption clone is **checked out** at it (what the M217 guard then re-asserts at bring-up).

**The generalization, and it is the point:** *a verdict must not outlive its subject* (the hazard above) has
a twin — **a precondition must not be assumed satisfied because it was satisfied locally.** The stale-verdict
family is about artifacts that lie about the past; this is about artifacts that lie about **reach**. Both
produce results confidently attributed to the wrong code, which is precisely how the perf-patch rot went
unnoticed for four releases.

### The image must be compiled from the PINNED ref, not the highest fetched tag (v2.6 M244 iter-25)

The rung above (and the M217 clone guard) prove the consumption *clone* is checked out at the pinned tag. A
**third** copy of the code slips past both: the demo *image build*. `up-injected.sh` builds each service
`:injected` image from a **build-scratch checkout** it materializes from the fetched refs — and M244 found it
resolving the **highest fetched v-tag** (`v1.351.0`) instead of the source clone's **pinned checkout**
(`v1.341.0`). The two schemas differed by one column (`ai_readiness_cycles.launched_by`): the binary `SELECT`ed
a column the migrated schema never created, so the cycles endpoint 500'd and every ai-readiness surface rendered
the **zero-state** — a failure that reads as a *seed gap*, not a *build skew*. The clone was at the right tag,
the guard was green, and the image was still wrong. Fixed durably (build-scratch resolves the pinned ref + an
M217-style preflight; rext `c755370`, +3 regression tests). The invariant: **clone-at-tag is necessary but not
sufficient — the image the container runs must be compiled from that same ref.** (#M244 iter-25)

### Drive every remote bring-up through a LOGIN shell (v2.5 M236 iter-03)

**`ssh host '<cmd>'` is not the same shell your operator gets.** A non-interactive `ssh host 'cmd'` sources
**no login profile**, so anything the profile puts on `PATH` is simply absent — and the host pre-flight then
fails in a way that **perfectly mimics a missing prerequisite**.

**Always:**

```bash
ssh <host> 'bash -lc "<the bring-up command>"'     # -l = login shell; sources the profile
```

**The concrete trap, because the shape is what makes it costly.** M236's `up-injected.sh` host pre-flight
reported *"Go NOT on PATH … install Go 1.25.x"* on `billion`. Go was installed and was the exact pinned
`go1.25.12` — at `/usr/local/go/bin/go`, a directory the **login profile** adds to `PATH`. What makes this a
trap rather than a footnote:

- **The two prereqs behave differently under the same invocation.** `atlas` lives in `/usr/local/bin`, already
  on the default non-login `PATH`, so it **passed**. One prereq green and one red reads as *prereq-specific*
  ("Go is missing") rather than *shell-specific* ("nothing from the profile is on `PATH`").
- **The remedy text reinforces the wrong reading.** The pre-flight recommends installing Go — so the operator
  installs a second Go, or edits `PATH` in a config file, and the real cause is never seen. A pre-flight that
  names a *remedy* rather than a *symptom* narrows the operator's hypothesis space, sometimes wrongly.
- **A partial-green pre-flight is weaker evidence than a fully-red one.** All-red would have been read as
  environmental in seconds.

**The cheap disproof is one command** — run it before believing any "prereq missing" verdict from a remote
pre-flight:

```bash
ssh <host> 'bash -lc "go version"'                 # green here + red in pre-flight ⇒ PATH, not prereqs
```

> **The general rule:** *when a check reports a missing dependency on a remote host, first prove the check and
> the operator are running in the same environment.* A tool's absence and a tool's **invisibility** produce
> identical output, and only one of them is fixed by installing anything.

## A SUITE NOBODY RUNS IS A SUITE THAT IS RED (v2.8 M256 harden pass)

Two findings from the same pass, both about the tooling's own tests rather than the stack's, and both worth
stating here because this is the *is-it-actually-working* doc.

**1. Enumerated test rosters go stale silently — and the roster is where the rot hides.** M255's close recorded
*"Python 1505 pass / 2 skip / 0 fail across **stack-core + demo-stack + stack-injection**"*. `stack-verify/tests`
is not in that list. Run at the M256 harden pass, it had **5 failures**, all from one cause: `autoverify.sh` grew
check (f) — the ant-academy `/library/` catalog probe — at M245, and of the four fixtures in `test_verify.py` that
stub `curl`, exactly **one** was taught about it. The others returned nothing, the probe warned, autoverify
reported *"N check(s) FAILED"*, and five tests went red — **none of them about the academy**. They also each paid
the probe's 3×3 s retry, so the suite ran 313 s instead of 163 s.

This is the *same* blind spot M255's own close found for `stack-snapshot`'s Go suite, one directory over: a suite
outside the roster had been red since v2.6 and two consecutive release closes missed it. **The lesson is not "add
stack-verify to the list."** It is that *an enumerated roster is a claim that needs its own fence* — the safe form
is "every section that has tests", discovered, with the count asserted, not a hand-maintained list of the sections
someone remembered. Until that exists, a close's Python verdict should be read as *"the sections we named are
green"*, which is a weaker statement than it looks.

**2. A test file can hide tests from the runner it invites you to use.** In six rext modules, a
`unittest.TestCase` subclass was defined **below** the module's `if __name__ == "__main__": unittest.main()`
guard. Python executes top to bottom, so the class does not exist when the guard runs: it is never registered,
never collected, and the run still prints **OK**. Measured — **15 classes / 76 tests**:

```
python3 demo-stack/tests/test_roster_invariant.py         ->  Ran 22 tests ... OK
python3 -m pytest demo-stack/tests/test_roster_invariant.py  ->  27 passed
```

The five silent tests included that file's own RED-proof for the live 12-dead-buttons cockpit defect — the class
its docstring calls *"the primary user-visible surface, and the one the defect is actually about."* Worst case:
`stack-injection/tests/test_apply_patch_selfheal.py`, **11 of 27**.

CI was never wrong — every runner uses pytest, which collects by inspection and ignores statement order. What was
broken is the loop a human or an agent actually uses: `python3 <the one file I am editing>` is what the file's own
`__main__` block invites, and it skips the part you just wrote. Appending a new hardening class to the bottom of a
file is the path of least resistance and the guard is already sitting there, which is why six files drifted the
same way across releases. Guards relocated to end-of-file; fenced repo-wide by
`stack-core/tests/test_test_collection_fence.py`.

> **The general rule behind both:** *a green verdict is only as wide as the set of things the runner actually
> looked at.* Both defects produced a confident PASS over an unexamined set — one because a roster did not name a
> directory, one because a file did not define a class yet. Neither is detectable from the verdict; both are
> detectable by asserting the **denominator**, which is why every fence this pass landed is fail-closed on how
> much it scanned.

### The third instance, and the biggest: `verify.sh` said "all live probes passed" over ZERO probes (v2.8 M256 harden pass 2)

The rule above was written about the tooling's *own tests*. The second harden pass found it in **this doc's
subject** — the live probe runner whose verdict everything downstream quotes.

`live/verify.sh`'s success condition was `fail_count -eq 0`. That is equally true of *"every probe passed"* and
*"no probe ran"*. Reproduced in one command:

```
STACK_SERVICES="skillpath" ./live/verify.sh
  ▶ liveness (is each service reachable?)     # ...no rows
  ▶ readiness (does each service answer correctly?)
  ✓ all live probes passed                    # rc 0
```

`skillpath` was a real service name until **M247** decommissioned it into `app` and deleted its row from
`lib/services.sh`'s `SERVICES` table. `target_service_selected()` returns 1 for an *unknown* name exactly as it
does for a *deselected* one, and — the aggravating half — `lib/target.sh`'s own contract has promised since M18
that *"A name not in the array is ignored **with a warning** (so a typo is visible, not silent)"*. **The warning
was never implemented.** So a stale, misspelled, or decommissioned name in `STACK_SERVICES` selects nothing,
silently, and the runner certifies it.

**Why this one propagates further than any other false green in the tree.** `verify.sh`'s rc becomes
`autoverify.sh`'s *"verify live"* verdict → `warnings=0` → `autoverify.json` `"green": true` → and that file is
the gate input for `run-latency.sh`, `run-coverage.sh`, `run-studio-fcp.sh` and `buildbench.py`, and is rendered
as `✓ pass` by `reports/generate.sh`. One unexamined set, five downstream consumers.

Fixed by the same discipline as the two above — **assert the denominator, and state it**:

- zero liveness probes is a hard refusal (`✗ NOTHING WAS PROBED`), naming `STACK_SERVICES`;
- the success line now reads `✓ all live probes passed (13 liveness + 7 readiness; 0 readiness probe(s) out of
  scope)`, so a reader sees the set without re-deriving it;
- the promised unknown-name warning exists, non-fatally (a bring-up's compose scope may legitimately name
  services this probe table does not model — the **zero** case is what fails).

Live-verified on a local `demo-2`: 13 + 7 at full scope, 2 + 2 under a legitimate `"postgresql redis"` filter,
refusal on the unknown name.

### Absence of evidence is not evidence — the `-s` trap (v2.8 M256 harden pass 2)

Two of `autoverify.sh`'s bring-up asserts keyed on `[ -s "$STACK_DIR/<log>" ]` and put the success message in the
`else`. **`-s` is false for a file that does not exist just as it is for one that exists and is empty**, so a
stack where the patch phase or the frontend-build phase never ran printed:

```
  ✓ demo-patches: all applied (none refused, none skipped)
  ✓ frontend builds: ok (the running images are this run's)
```

The second claim is a **gate input** by `autoverify.sh`'s own comment (`run-latency.sh` refuses to measure without
it), and it is exactly the claim an absent log cannot support. The distinction is decidable and free:
`up-injected.sh` **truncates** each log as it enters the phase (`: > "$STACK/demopatch.log"`), so **the file's
existence is the writer's receipt.** No receipt, no claim — absent is now its own branch and it warns.

> **The generalization worth carrying:** when a check reads an artifact to decide, enumerate **three** states, not
> two — *present-and-clean*, *present-and-dirty*, and *not-there-at-all*. Collapsing the third into the first is
> how a phase that never ran becomes a phase that passed. The same shape appeared four times in one pass: `-s` on
> two logs, a missing `generatedAt` in `run-coverage.sh`, and `EXPECTED_PAIRS` set-but-empty in
> `aggregate-content.py`.

### A guard that cannot find its subject must not exit 0 (v2.8 M256 harden pass 2)

`rc 0` is what an `&&` chain and a CI step read; a `SKIP` line on stdout is not a signal. `stack-core`'s
`dev_flag_guard.py` had **two** `SKIP → return 0` paths, so renaming `dev-stack/dev-stack` made the flag
cross-check silently stop existing — in a file that already reasoned the principle out four lines lower (*"An
empty result is a FINDING, not a pass"*) for its *empty*-result case but not its *missing-input* case.

The resolution is not "fail on every skip" — it is that **the two skips are different things**:

| situation | verdict | why |
|---|---|---|
| no rosetta root found | `rc 0`, and the message says **NOTHING WAS CHECKED** | rext can legitimately be cloned without rosetta beside it; failing a build for that is wrong |
| `dev-stack/dev-stack` absent | `rc 2` | that file ships **inside rext** — its absence is a broken tree or a moved path, never a rosetta-less checkout |

**`rc 2`, not `rc 1`:** *"could not run"* is a different diagnosis from *"found a problem"*, and collapsing them
sends the reader in the wrong direction — the same split `ptvalidate` already makes for `datadna`'s exit 3.
`corpus_index_guard.py` took the same treatment: it reported *"every doc has its directory-README index row"* from
an empty violation list, which is equally true when **no** directory was index-bearing (rename every `README.md`
and it sweeps zero docs and passes), so it now counts what it inspected and says so — *"all 82 doc(s) across 6
index-bearing directory/ies"*.

### A DELETED subject cannot fail a check that iterates the subjects (v2.8 M256 harden-final)

This is the same shape as the casbin-list fence (pass 1) and the `safety_doc_drift` host loop (pass 2), and
harden-final found it **four more times in one afternoon**. It is worth stating as a rule because it keeps
arriving in new costumes:

> **A check that RANGES OVER the thing it is checking cannot see a deletion.** The deleted item is absent from
> the list it would have been compared against, so both sides stay balanced and the check passes over less.

The four instances, and what each one measured:

| Where | What it iterated | What a deletion did |
|---|---|---|
| `playthroughs/manifest` | `m.UseCases()` — in every check | Delete a use case **and its spec**: `ptvalidate` reports `manifest VALID`, 30 use cases ↔ 29 ids **balanced**, the Go suite green. A whole milestone deliverable vanishes silently. |
| `seed-facts-fence` | `SEEDED_HEROES`, the module under test | 11 heroes seeded, **6 enrolled**. The 5 unenrolled were asserted by literal in the specs that play them; renaming one in the seed left the fence green and reddened the Playthrough instead, naming a product regression that had not happened. |
| `orgless_writers_fence` | files matching `membershipUUID(prefix, i)` | The pattern's scope was a **naming convention**. Rename the call's identifiers and the file is not selected at all — so the fence reports green over a writer with no guard. |
| `report.AllGreen` / `NoRegressions` | the `Summary` counters alone | Equally true of *"everything passed"* and *"nothing ran"* — and `report_edge_test.go` **asserted the vacuity as correct**, so writing the honest fence turned that test red. |

**The fix is the same each time, and it is not "iterate harder":** add an **EXTERNAL** assertion that names what
must exist, with a **denominator**. A presence pin that names ids (`pt-onboarding-individual`, …). A reverse
direction (seed → facts, not only facts → seed) with an explicit, *justified* exemption set. A `Total > 0`
clause. A floor set to the **measured** count, not comfortably below it — the org-less fence's floor of 6
against a true count of 8 left room for exactly the two renames it existed to detect, and **slack in a floor
permits the defect the floor is for**.

**And the exemption must be a property, not a roster.** The citation fence added here excuses a comment block
that *declares* a file absent (it is documenting an absence, not citing evidence) rather than keeping a list of
excused filenames — because an enumerated roster is the process bug this milestone has hit repeatedly.

### A citation is a claim about the repository, so it is checkable (v2.8 M256 harden-final)

Two shipped comments in `stack-seeding/seeders` cited **`orgless_footprint_test.go`** as the compensating
control for the half of the org-less capability no static fence can cover. **That file has never existed** —
not on disk, not anywhere in git history. The underlying measurement was real (a live uuid sweep on a reseeded
`demo-2`), but it lived in prose and the comments claimed a committed test.

This is the `demo_knob_guard` `Read at` rot (**28 of 29** citations wrong, guard green) in source comments
instead of a document — and that guard's fix was scoped to one document's citation format, so it generalised to
nothing. **A wrong citation is worse than no citation:** it sends the reader somewhere confident and empty, and
it makes an *unverified* property read as a verified one. Fence it where it lives.

### A gate whose exit code is discarded is not a gate (v2.8 M256 harden-final)

`run-playthroughs.sh` ran `ptreport --gate no-regressions` and threw the result away with `|| echo`, then exited
with playwright's code. `report.go`'s `!ok` branch is the **only** mechanism that notices a declared Playthrough
which never ran, and **playwright exits 0 when a spec file is simply absent** — so deleting a spec produced a
fully green run. Measured: 203 passed, rc 0.

The non-fatality was right for a `--grep` run (every un-grepped id correctly reports *"did not run"*), and the
comment above it had always said so — *"the whole-suite gate is meaningful on a full (un-grepped) run"*. The
**code collapsed both cases onto the permissive one**. Splitting them is the fix: binding on a full run,
advisory with a stated reason on a scoped one.

> **Then the fix itself taught the sharper lesson.** The first version went red live — but via `set -e`, which
> aborts a bare failing command group before `PTREPORT_EXIT=$?` can run. The full-run case produced the right
> exit code **by accident**, and the `--grep` case would have aborted with no explanation, silently destroying
> the advisory path the change existed to preserve. **A right answer reached by the wrong mechanism is not a
> right answer — it just fails somewhere else.** Put the invocation in an `if` condition (exempt from `set -e`)
> so the code decides, and re-prove it live. It was caught only because the *diagnostic did not print*.

### A CONTAINER-LIVENESS assert is the cheapest cheap-win there is, and the stack has no restart policy (v2.8 M256 iter-15)

**A clean `Exited (0)` is not a healthy container, and nothing in the bring-up notices.** Measured on `demo-2`:
Postgres restarted un-cleanly (*"database system was not properly shut down; automatic recovery in progress"*),
and `jobsimulation` + `cms` **self-terminated by design** on their DB-health monitors — *"DB too many ping
failures, shutting down"*, exit status **0**, a graceful shutdown working exactly as written. **Nothing restarts
them.** The stack then sat at **14 of 16 containers "Up"** for an hour.

What that looks like from the outside is the part worth internalising:

- **The application surfaces no error at all.** Every `jobsimulation`-backed surface renders **20 table rows
  with empty cells** — no empty state, no spinner, no toast, no console error. Held for 40 s of watching.
- **A `rows > 0` assertion passes on it.** The Playthrough covering that surface failed only three steps
  later, on a row-link wait, reporting a timeout that blamed a locator — three layers from the cause. An hour
  went into disbelieving a correct new assertion before anyone read `docker ps -a`.
- **It is NOT the ENOSPC trap.** Disk was 5 % used, 227 GiB free. [`demo/build-budget.md`](demo/build-budget.md)'s
  M239-F1 records that a mid-campaign ENOSPC *presents* as `redis exited (1)`, and this looks like its cousin
  and is not — so check disk, then stop, rather than assuming.
- **Recovery is `docker start <container>`** — no build, no compose, no teardown. The same class of action
  `run-playthroughs.sh --reset` already performs on the fake services every run.

**The gap this names — corrected at M257, because the first reading of it was wrong.** The cheap-win set does
not fire on this, but **`autoverify` as a whole is not blind to it**: `stack-verify/lib/services.sh:43-44`
carries `jobsimulation` and `cms` rows, both inside the demo `--services` scope (`up-injected.sh:2494`), and an
`Exited (0)` container un-publishes its port → `code=000` → `status=down` (`services.sh:130-133`) → `verify.sh`
rc≠0 → `autoverify.sh` warns → `green:false`. **A re-run would have gone red.** The stack stayed green through
the **stale-verdict class** — an `autoverify.json` written once at the bring-up tail and read later as current,
the same F-6/F-10 hazard recorded at `:252-260` and `:293-303` — **not** a missing check. The genuine coverage
gap is narrower and was undocumented: the containers with **no `services.sh` row at all** — `fake-fapi`,
`fake-bapi` and `hiring-app` (`gen_injected_override.py:353,560,581`). **That** is the 14-of-16 arithmetic:
13 rows in the demo scope against 16 containers. A **container-liveness assert** — *the stack's declared
container set is all running* — is one line, costs nothing, and would have named this instead of an hour of
Playthrough archaeology. Routed as `FIX-M256-demo2-service-self-termination` → M257/M258; landed at M257 as
`autoverify.sh` check **(h)**, which derives its expected set from `services.sh` and adds the injected trio.

**The general rule: check container liveness BEFORE diagnosing a test.** It is the cheapest measurement
available and it is the one that was taken fourth.

## What this doc does NOT verify — reach (v2.5 M236, user-authorized)

**Restricting *who can reach* a demo is the VM's and the VPN's job, not the demo stack's.** The stack's only
obligation on this axis is to **permit** VPN access; it does not enforce, narrow, or attest reach, and **no
gate here measures it.** v2.5 M236 opened with an exit-gate clause requiring the demo to be *"reachable only
over the tailnet"* and **dropped it** on that basis, with no off-tailnet probe deliverable.

**Why this is a scoping stance and not a safety claim.** It says which *layer* owns the control, not that the
control is unnecessary — and **not** that there is nothing to protect. That disclosure — including the fact
that **every demo container publishes on `0.0.0.0`, flag or no flag** — stands **as-is** and needed no
amendment for this decision. Read `safety.md` §3 Part 3 for the exposure picture; read this line only as
*the demo stack is not the layer that verifies it.*

> 🔴 **The rationale is demo-shape-dependent — do not generalize it.** For a **synthetic** demo the stack can
> decline the job cheaply, because [`safety.md`](safety.md) §3.3 argument 1 holds: no customer data **can** be
> in it, so reach is a network-perimeter concern rather than a data-exposure one. **For a content-story demo
> (§3.8) that premise is false** — it carries best-effort-scrubbed **real production session content**, and
> §3.3 marks argument 1 as not holding for that shape. Reach there **is** a data-exposure concern.
>
> **The layering decision is unchanged** (it is a statement about *ownership* of the control), but the weight
> it carries is not. Per [`safety.md` §3.3.1](safety.md), the VPN/tailnet scope is promoted from a supporting
> comfort to **the** control for that shape — the one the data-controller acceptance was explicitly
> conditioned on. So *"no gate here measures it"* is a **materially heavier gap for a content-story demo**:
> the control the acceptance rests on is **operator-maintained and unattested**, as strong as the network the
> box is on, and nothing in this document will tell you if it is weaker. `safety.md` §3.4 residual #2 records
> the same consequence from the other side.

## Cross-references

- [`rosetta_demo.md`](rosetta_demo.md) — the demo lifecycle + the unified registry whose **recorded ports**
  the offset cross-check reads.
- [`idempotency.md`](idempotency.md) — what happens when you re-run the step a failed verify points you at.
- [`safety.md`](safety.md) — why even the `docker exec` reads here only ever touch a **per-stack-isolated**
  store, never prod.
- [`/test-platform`](../../.claude/skills/test-platform/SKILL.md) — the operator-driven, deeper
  verification surface this auto-run is a default-on, non-fatal smoke subset of.
- [`demo/coverage-protocol.md`](demo/coverage-protocol.md) — the two Playwright sweeps that verify what a
  bring-up actually RENDERS, where this doc verifies that it came UP: the v1.10 M42e hero-vantage coverage
  sweep, and (v2.5 M236) the **content-stories `(session × action)` LANDS sweep**. The relationship matters
  operationally: `run-content-stories.sh` **gates on this doc's `autoverify.json` being fresh AND green**
  before it will trust its own reading, so a stale or red verify invalidates the render proof rather than
  being worked around.
- [`demo/content-stories-spec.md`](demo/content-stories-spec.md) — the `content-manifest.json` those sweeps
  read. A cockpit brought up without `--content-manifest` serves **404** and the sweep fails closed; the
  bring-up wires it via `--content-export` (non-fatal), logging to `$STACK/content-export.log`.
