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

3. **Directus serves the catalog** — `SELECT count(*) FROM directus.directus_collections` via `docker exec
   <project>-postgresql-1 psql …`, asserted `> 0`. The silent-failure analog of the casbin assert: a Directus
   can be UP (`/server/health` 200) but serve **nothing** if the content-model never registered. (Also runs as
   the `directus-collections` readiness probe.)
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
  `PING`, GraphQL introspection (`:5050+offset`), gotenberg version (`:3200+offset`), sentinel
  Connect-RPC handler mounted (`:8087+offset`), storage RPC reachable (`:8301+offset`), and — on a
  local-content stack — the per-stack **Directus** liveness (`/server/health` at `:8055+offset`) plus its
  `directus-collections` serve-check — each resolving the offset port + project container via the same
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

### `autoverify.json` — the machine-readable signal

```json
{"project":"demo-1","offset":10000,"warnings":0,"green":true}
```

Its path is **`rosetta-extensions/demo-stack/stacks/<project>/autoverify.json`** — *not* `<stack>/autoverify.json`,
as this doc claimed until M218 (**iter-01 F-2/D3**; the tooling that gates on it, `stack-verify/e2e/run-latency.sh`,
reads the real path).

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

## Cross-references

- [`rosetta_demo.md`](rosetta_demo.md) — the demo lifecycle + the unified registry whose **recorded ports**
  the offset cross-check reads.
- [`idempotency.md`](idempotency.md) — what happens when you re-run the step a failed verify points you at.
- [`safety.md`](safety.md) — why even the `docker exec` reads here only ever touch a **per-stack-isolated**
  store, never prod.
- [`/test-platform`](../../.claude/skills/test-platform/SKILL.md) — the operator-driven, deeper
  verification surface this auto-run is a default-on, non-fatal smoke subset of.
