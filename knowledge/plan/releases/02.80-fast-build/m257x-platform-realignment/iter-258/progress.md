# iter-258 — progress

**Type:** tik (protocol: [`corpus/ops/platform-alignment.md`](../../../../../corpus/ops/platform-alignment.md))
**Opened:** 2026-08-10T13:50:17Z

## Phase A — seal + pre-flight

Pre-registrations sealed in this iter's first commit before any bring-up work, per `TOK-08`'s standing
discipline.

### Environmental facts re-verified at open (13:44–13:49Z)

Measured, not inherited — each is something this iter depends on:

| fact | measured | note |
|---|---|---|
| `demo-1` live | **11 containers, `Up 4 days`** + a 4-day `cockpit.py` on `:17700` | never touched by this iter |
| canonical registry | `stack-core/.stacks/registry.json` = `demo-1` only | **slot 2 free** |
| disk | **206 GiB** free | ENOSPC is not this box's risk |
| docker driver | **`overlayfs`**, 8 CPU, 11.67 GiB VM | no image-unpack leg — `billion`'s attribution is not portable |
| load avg | 2.79 / 4.84 / 7.79 | third-party (`anima8` pytest battery); **timings are CONTENDED** |
| rext pin | `fast-build-m257x-iter-101` → `0011c10a`, **on origin** | rung zero passes; 157 iters stale — see `D-M257x-258-1` |
| rext authoring copy | **5 ahead / 0 behind** `origin/main`, untagged | invisible to any stack |
| clone advance | `app 3eaadae68` · `next-web-app 19423a1fb` · `ant-academy 249430c3` | == `clones.pin.json`, all six repos |
| `demo-1` build ref | **`demo-1/clones/app` = `ad9f3c498`** | **the pre-advance ref — demo-1's green does not cover the advance** |

### Three corrections the re-survey made before any work

1. **Two registries, and the obvious one is the wrong one.**
   `demo-stack/stacks/registry.json` lists **slot 2** and **omits the live `demo-1`** — it is the
   pre-M12 demo-only legacy file kept for provenance. `stack_registry.py:51` puts the allocator's
   registry at `stack-core/.stacks/registry.json`. The two disagree in both directions, and only the
   second decides anything.
2. **`stacks/demo-2/` and `stacks/demo-4/` already exist** as Jul-31 skeletons (3 files each) from a
   `rosetta-demo` invocation that generated an override and never brought up. `demo-2`'s
   `docker-compose.demo.yml` maps **only `app`'s port** — a 2-line file against `demo-1`'s full artifact
   set. Not a blocker (`up-injected.sh` consumes `$STACK/.env.demo-$N`, which the skeleton supplies with
   correct `COMPOSE_PROJECT_NAME=demo-2` / `DEMO_PORT_OFFSET=20000`), but it is a stale artifact in the
   path and is recorded as such.
3. **The decisive one:** `demo-1` is green at `ad9f3c498`. The advance is 28 commits past that. **The
   stack that is green is not the stack the corpus now describes** — which is exactly why this route
   was left open and why this iter is not redundant.

