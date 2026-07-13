# M217 — Progress

_Section checklist, derived from `overview.md` § Scope.In + the `kb-fidelity-audit.md` §7b build plan._

> **Order is load-bearing.**
> - **S1 first** — S3/S5 build on the demopatch contract being written down.
> - **S2's port reap MUST precede S7's re-pin** — the cockpit/academy pidfiles and `tailscale-serve.sh` live
>   *inside* the clone S7 replaces. Re-pin first and every leaked listener becomes **permanently unreapable**.
> - **S4 before S8** — `jobsimulation` is the near-certain cause of the failing autoverify probe. Fix it, *then*
>   re-measure; do not hunt the verify failure separately first.

## Sections

- [x] **S0 — Clear the RED KB gate** (docs only, no code)
  - [x] Correct the 3 false claims in `overview.md`: the jobsimulation root cause (AWS bind mount, **not** a missing
        subcommand — the drafted fix would have broken the service); the stale-cockpit mechanism (**no** dead
        clerk-ids); **two**, not three, replay cache-misses (directus is rc=4 from `--no-local-content`)
  - [x] Kill the **live orphaned cockpit on `billion`** (pid 83214, `0.0.0.0:17700` — survived the `/demo-down`)
  - [x] Confirm a demo carries **no AWS credentials at all** (0 hits for `AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY`)
  - [x] Write `kb-fidelity-audit.md`
  - [ ] Propagate the corrections to `roadmap.md` + `state.md`

- [x] **S1 — `corpus/ops/demo/demopatch-spec.md`** (the blind area — FIRST, before any code)
  - [ ] Author the spec from the audit's §5-0 ground truth: G1–**G7**, the 10 mandatory manifest keys, the sha gate
        (**whole-file** — the rot source) + both exit-code spaces, the byte-exact exactly-once anchor, the **three
        apply vehicles**, the **chain rule**, why the `app` patches are never reverted, the three `DEMO_NO_*`
        opt-outs, the full patch inventory, the BD-3 decision, the freshness preflight, the re-pin runbook
  - [ ] Back-link from `frontend-tier.md`, `coverage-protocol.md`, `seeding-spec.md`, `stories-spec.md`,
        `ai-readiness.md`, `tailscale-serve.md`, `clerkenstein.md`
  - [ ] Index it from `corpus/ops/demo/README.md` + `CLAUDE.md`

- [x] **S2 — Port reap: bring-up + teardown** (scope 1)
  - [ ] `demo-stack/reap.sh`: `reap_port` + `reap_stack_ports` (argv-guarded; port-authoritative, not pid-only)
  - [ ] `rosetta-demo cmd_down`: reap the **whole offset range** after `compose down`; stop the pidfile lie
        (`:152-156` discards `kill`'s status, `rm -f`s regardless, and prints "stopped the presenter cockpit" even
        when nothing was killed)
  - [ ] `up-injected.sh`: **pre-bind reap** before the cockpit launch + a compose-range preflight before `up`
  - [ ] `cockpit.py`: wrap the bind (`:567`) in try/except → clean exit 2, not an unhandled traceback
  - [ ] Replace the **unconditional** "presenter cockpit serving on …" log with a real `/healthz` probe
  - [ ] Docs: `cockpit-spec.md` (teardown is **port**-authoritative), `demo/README.md` + `demo-down/SKILL.md` (the
        real teardown), `rosetta_demo.md` (the registry is **not** the port source), `idempotency.md` (re-up over a
        half-dead stack now self-heals)

- [x] **S3 — Un-swallow the REFUSE reason** (scope 2)
  - [ ] `up-injected.sh:701,717` — capture the helpers' stderr and re-emit it in the else-branch (**keep NON-FATAL**)
  - [ ] Same discipline at `autoverify.sh:138` — it swallows `verify.sh`'s stderr, collapsing N failing probes into
        exactly **1 nameless** warning. Propagate the real fail-count + name the probes; tee to `<stack>/autoverify.log`
  - [ ] Static fence in `test_frontend_build.py`: assert no applier invocation redirects stderr to `/dev/null`
  - [ ] Preserve the asymmetry: the three next-web `demopatch` calls already do **not** swallow

- [x] **S4 — `jobsimulation` exits(1)** (scope 4)
  - [ ] `gen_injected_override.py::build_lines` → `if name == "jobsimulation": body.append("    volumes: !reset null")`
  - [ ] Mirror the fix on the **dev** path (`stack-core/gen_override.py`) — same bug, same 3 lines (Fate-1)
  - [ ] Non-fatal warn if `[ -d "$HOME/.aws/credentials" ]` (a Docker-created directory is a smell worth surfacing)
  - [ ] `test_injection.py`: positive assert (`volumes: !reset null` on jobsimulation) **and negative** (no service
        ever gets a `command:` key — the drafted fix that would have broken it)
  - [ ] Doc: `corpus/services/jobsimulation.md` — a **Startup contract** section (root `RunE` **is** the server; no
        `serve`/`run` subcommand; any init error ⇒ `Error:` + usage + exit 1, so *"it printed help" means an INIT
        ERROR*)

- [x] **S5 — Re-pin the two `app` perf patches + the LOUD freshness preflight** (scope 3)
  - [ ] Re-pin `app-targetrole-authz-skip` (one pin works on both boxes)
  - [ ] Re-pin `app-aireadiness-snapshot-loadmembers` **at the tag the target box builds** (⚠ the two boxes diverge —
        see BD-3)
  - [ ] The freshness preflight (fails **LOUD**, emits paste-ready corrected pins) before the inject loop
  - [ ] A `--repin` verb
  - [ ] **Live-clone pin tests for BOTH `app` manifests** — closing the test gap that let the drift ship

- [x] **S6 — Prime the snapshot cache on `billion`** (scope 5)
  - [ ] `rsync` the 3 digest dirs (~1.45 GB) to `~/panorama/.agentspace/snapshots/`; verify with `stacksnap status`
  - [ ] Re-run **with local content ON**, from a **purged** stack, `directus:11.6.1` pinned
  - [ ] Docs: `snapshot-cold-start.md` — a **new Option 3: ship a warm cache to a remote box** (every existing
        option is **dead on that box**: no `~/.pgpass`, no staging dump); `snapshot-spec.md` — why a cache is
        transportable at all (row surfaces digest their own tables; directus digests the whole schema)

- [x] **S7 — Re-pin the drifted rext clones** (scope 6) — **AFTER S2's reap**
  - [ ] Re-verify the subset proof, then `git fetch --tags` (**mandatory** — neither clone has `v2.2`) + `checkout`
  - [ ] **Fix the drift injector**: `.claude/skills/stack-secrets/SKILL.md:75` hardcodes `stage-door-m30` and
        checks it out in the same clone `/demo-up` pins — and `/demo-up` invokes `/stack-secrets`. Without this the
        re-pin is a **no-op within one run**
  - [ ] Promote `ensure-clones.sh`'s pin guard from **WARN to FAIL** (+ `DEMO_ALLOW_UNPINNED_REXT=1`)
  - [ ] Docs: `rosetta_demo.md` (drop the stale `v1.10.1` prose copy of the pin — the prose copy *is* the drift
        class); `tailscale-serve.md` (resolve the `<panorama-tag>` placeholder **and add the missing
        `git fetch --tags`** — that omission is how the remote landed on a bare sha)

- [ ] **S8 — Green-stack proof + a machine-readable signal**
  - [ ] Cold reset-to-seed `/demo-up` on `billion` at the new rext tag
  - [ ] Emit `<stack>/autoverify.json` so **M218 can gate its measurements on "the stack came up green"** (today
        `/demo-up` exits 0 on a red verify and still prints UP)
  - [ ] Add the 4 missing cheap-wins: demo-patch **applied**, snapshot **replayed**, cockpit **up**, **fake-fapi up**
        (*a dead fake-fapi means nobody can log in — and verify stays green today*)
  - [ ] **Exit gate:** 0 verify warnings · 0 leaked ports · 3/3 replays exit 0 · 2/2 app patches applied ·
        jobsimulation serving · cockpit serving a fresh manifest · both hero vantages reachable

## Notes

**2026-07-13 — the KB gate came back RED and it was right.** 14 load-bearing stale claims, **three inside this
milestone's own `overview.md`**. The worst: the drafted `jobsimulation` fix (`command: serve`) would have produced a
real `unknown command "serve"` → exit 1. The root cause is a `$HOME/.aws/credentials` bind that Docker auto-creates
as an empty **directory**, which makes the AWS SDK hard-error inside `ai.NewAIManager`, which makes cobra print its
usage block — the "prints CLI help" everyone mis-read as a missing subcommand.

**Also surfaced:** the `/demo-down` I ran earlier today **left an orphaned cockpit alive on `billion`** —
`0.0.0.0:17700`, pid 83214, an unauthenticated hero-vending panel pointing at a deleted database. Killed manually.
That is S2's defect, live.


## S8 findings — a NEW Linux-only defect, found by the proof run

**`host.docker.internal` does not resolve on Linux Docker Engine.** The per-stack Directus provision reaches the
host's offset Postgres from *inside a container* by that name — a **Docker Desktop** convenience. On `billion`:

```
docker run --rm alpine getent hosts host.docker.internal                       -> (nothing)
docker run --rm --add-host=host.docker.internal:host-gateway alpine getent ... -> 172.17.0.1
```

So `CREATE SCHEMA directus` failed, the **entire local-content provision was skipped**, and the demo quietly fell
back to reading content **live from prod over the WAN** — while the directus replay exited **rc=4** (schema
missing), which *looks* like a snapshot-cache problem and is not. **Priming the cache could never have fixed it.**

**This is the same class as the `jobsimulation` AWS-mount bug: fine on a Mac workstation, dead on a fresh Linux
VM.** Both were invisible because nobody ran a demo end-to-end on Linux until v2.2 put one on `billion`. Two
independent instances of the same blind spot in one milestone is a pattern worth naming — **the demo tooling was
implicitly Docker-Desktop-shaped.**

Fixed with `--add-host=host.docker.internal:host-gateway` (Docker 20.10+, a no-op on Docker Desktop).

## Proof-run evidence (billion, cold reset-to-seed)

| What | Before M217 | After |
|------|-------------|-------|
| `app-targetrole-authz-skip` | REFUSED silently (4 releases) | **applied** |
| `app-aireadiness-snapshot-loadmembers` | REFUSED silently (4 releases) | **SELF-HEALED + applied** — on the v1.337.0 box where *no committed pin could ever be right* (`b3216968… → dc9e167e…`) |
| taxonomy replay | cache miss (rc=5) | **330,261 rows replayed** |
| sim-embeddings replay | cache miss (rc=5) | **1,490 rows replayed** |
| directus | rc=4, prod-read over the WAN | Linux `host-gateway` fix (S8) |
| `jobsimulation` | exits(1) in every demo | **no longer in the failing probe set** |
| autoverify | "1 check(s) FAILED" — *no name* | **names every failing probe** |
| cockpit | logged "serving" unconditionally | **healthz-gated**; a dead cockpit is reported |
| rext pin | warned; both clones drifted anyway | **FATAL**; both clones at `cue-to-cue-m217` |
