# Live engines — developing hyper-studio and studio-desk against a running dev stack

> **v2.10 M273 "hstudio-integration".** Two of a stack's engines can be served **natively from a git
> worktree** instead of from a container image, so an edit is live in seconds rather than a rebuild away.
> The two are **lodge** (`hyper-studio`) and **studio-desk**, and they are being developed together.

## Why this exists

Baked into images, a change to either engine costs an image build and a container replace. Served
natively they hot-reload: **lodge in ~3 s, studio-desk in ~2 s** (both measured, 2026-08-24, by editing
the worktree and polling the running stack until the change appeared). The stack around them stays
real — real backend, real Postgres, real content — because only the engine moves.

## The switch

```bash
rext=.agentspace/rosetta-extensions
$rext/dev-stack/engine-switch.sh 1 live      # serve both natively from the worktrees
$rext/dev-stack/engine-switch.sh 1 status    # which engine is serving each port, and from where
$rext/dev-stack/engine-switch.sh 1 baked     # back to the container image
$rext/dev-stack/engine-switch.sh 1 live --engines lodge     # one engine only
```

It operates on a **running** stack and takes ~3–5 s each way. `dev-stack up` remains the way to build a
stack; this only changes which engine serves it.

| | live | baked |
|---|---|---|
| **lodge** | native from the worktree, wire `8080 + N·10000`, panel `7787 + N·10000` | the `dev-N-lodge-1` container on the same ports |
| **studio-desk** | native from the worktree on `9200 + N·10000` | **nothing serves it** — a dev stack has no studio-desk container at all |

## The worktrees

Branch `feat/hstudio-integration` in both repos, checked out under `stack-dev/.worktrees/`:

```
stack-dev/.worktrees/hstudio-integration-hyper-studio    (pnpm install --frozen-lockfile)
stack-dev/.worktrees/hstudio-integration-studio-desk     (npm ci)
```

Override with `LODGE_WT` / `DESK_WT`, or rename the project with `HSTUDIO_PROJECT`.

## Five things that will cost you an afternoon if you don't know them

**1. `live` is a runtime overlay, not a different stack.** The lodge container is *stopped*, not removed:
it stays in the generated override, the port registry, the verify scope and the named-volume allowlist.
That is deliberate — keeping the override byte-identical is what makes the switch cheap and reversible.

**2. Lodge's job state does not travel.** The container keeps state on the named volume
`dev-N_lodge-data`; the native process keeps its own under `<worktree>/.agentspace/serve-lodge-dev-N/`.
Switching migrates neither, and leaves both intact — so switching back restores exactly what the
container had.

**3. The panel's auth differs between the two, and that is correct.** The container publishes through
docker, which can only forward to a `0.0.0.0` bind — and a non-loopback bind forces a ≥16-char secret
**and** a Host allowlist. Native lodge binds **loopback**, where none of that applies, so the live panel
takes **no credential at all**. Same panel, different exposure, different gate. Don't "fix" one to match.

**4. Studio-desk needs a CORS origin the backend does not ship.** `app/internal/cors/cors.go` allows the
**pre-migration** ports `9000`/`9100`. A native studio-desk on an offset port is refused at the
preflight, and the symptom is vicious: a signed-in Studio whose every data panel says "couldn't load",
over a backend whose logs look perfectly healthy because it answers `OPTIONS 204` and the POST never
arrives. `engine-switch.sh` adds the origin to `platform/.env` — but **the backend must be RE-CREATED,
not restarted**: compose reads `--env-file` at *create* time, so `docker restart` silently changes
nothing. The switch prints the exact recreate command.

**5. A fresh worktree has no `.env`, and the failure is AI-shaped.** studio-desk reads its own `.env`
from the source dir; a new worktree has none, and nothing provisions it automatically. The app still
boots, still answers health 200, still renders — and every AI call 500s with Azure's *"The API deployment
for this resource does not exist"*, because without `AI_PROVIDER_CHAIN` / `AI_OPENAI_API_KEY` /
`AI_ANTHROPIC_API_KEY` it falls through to the `AZURE_OPENAI_*` pair `platform/.env` does export, for
which no deployment name is configured anywhere. `engine-switch.sh` now seeds it from the canonical clone
and says so; if neither exists it warns and tells you to run `/stack-secrets`. Symptom to recognise:
`POST /api/ai/triage 500` from the advanced builder's copilot.

**6. `/api/health-check` answering 200 does not mean studio-desk works.** The route is public and
Clerk-blind by design, so it answers 200 while every gated page 500s on an empty publishable key. Judge
studio-desk by a page, not by its health route.

## The dev loop

1. Ask for a change to `hyper-studio`, `studio-desk`, or both.
2. The change is applied in the worktree on `feat/hstudio-integration`.
3. The running stack reflects it — no command, no rebuild.
   - lodge: `tsx watch` over `code/**`, ~3 s. **`LODGE_TYPE_CORPUS_HOME` points at the worktree's own
     `code/hyper-artifacts`**, so type-corpus edits are live too, not just engine code.
   - studio-desk: `next dev` (Turbopack), ~2 s.
4. Some changes are outside a watcher's reach (dependency changes, `next.config.ts`, env). Re-run
   `engine-switch.sh N live` — it reaps its own predecessor and restarts.

## Lifecycle

A native engine is a **host process**: `docker compose down` knows nothing about it, and a leftover
would hold the stack's offset ports and break the next stack allocated them — the same failure a
leftover `tailscale serve` causes. So `dev-stack down N` reads the `live.env` marker the switch writes
and calls `engine-switch.sh N stop` first. `stop` is teardown's verb and is deliberately **not** `baked`:
`baked` would start the lodge container, which is exactly wrong while the stack is being destroyed.

Ports are reclaimed with the identity-checked reaper the host-native services already use
(`reap_port <port> <identity-regex>`), which refuses an empty pattern and **reports rather than kills** a
process it cannot identify as ours.

> ⚠️ **`stop` HONOURS `--engines`, and did not until 2026-08-24.** Every other verb filtered and this
> one swept both, so `stop --engines lodge` killed studio-desk as its second act. That matters more
> than it sounds: stopping lodge alone is the obvious way to prove studio-desk survives a lodge
> outage, and the broken form made the result read as *the app crashes when lodge goes down* — the
> exact opposite of what the operator was testing. The two are indistinguishable from outside the box.
> Fixed, with `dev-stack/tests/test_engine_switch.py` asserting each engine's teardown sits under its
> own guard. `live.env` is now cleared only when the last engine goes.

## studio-desk needs lodge's address, and the switcher provides it

studio-desk fires a second generation at lodge whenever an author presses generate. It reads four
**server-only** variables, and `engine-switch.sh` writes them into the desk env at this stack's own
**offset** ports (`lodge_desk_env_apply`):

```
LODGE_ENABLED=1
LODGE_WIRE_URL=http://127.0.0.1:$((8080 + N*OFFSET))
LODGE_PANEL_URL=http://localhost:$((7787 + N*OFFSET))
LODGE_CUSTOMER=studio-desk
```

Two properties worth knowing, both of which have a failure mode that is silent:

- **They are REWRITTEN, never appended.** An append-only helper leaves two assignments of every key
  after two runs, and dotenv readers disagree about which wins.
- **The offset is the whole point.** Un-offset `8080`/`7787` submits this stack's designs to whatever
  lodge holds the default ports — on a two-stack box, the other stack's — and the jobs land in a panel
  nobody is watching. Nothing errors.

Both desk paths get them: the dev worktree's `.env`, and the demo's copied container environment (a
demo image built before the second engine existed carries none of these, so they are added at run time
rather than baked). The integration itself is
[`../services/lodge.md`](../services/lodge.md) § *studio-desk submits to it*.

## Related

- [`../services/lodge.md`](../services/lodge.md) — the engine, its two listeners, the boot preflight
- [`../services/studio-desk.md`](../services/studio-desk.md) — the Next.js 16 app and its ports
- [`rosetta_demo.md`](rosetta_demo.md) — the stack lifecycle these engines ride
