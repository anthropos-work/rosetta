---
title: "KB Fidelity Audit — M217 clean-stage"
date: 2026-07-13
scope: milestone:M217
invoked-by: build-milestone
verdict: "RED (cleared by the S0 correction pass — see §8)"
---

# M217 "clean stage" — KB-Fidelity Audit

**Audited:** 2026-07-13 · 7 topics · docs = `corpus/**`, code = `.agentspace/rosetta-extensions/**` (authoring clone @ `v2.2`, HEAD `39e8013`), plus read-only reads of `stack-demo/`, `stack-dev/`, and `devops@billion:~/panorama/`.
**Rule of the audit:** docs lie; code does not. Every claim below carries `file:line` evidence.

---

## 1. Verdict

# 🔴 RED — remediable, but do not start coding from the current docs (or from the current milestone overview)

Two independent reasons, both under the verdict rules as given:

**(a) Load-bearing STALE claims exist — many.** Fourteen findings are marked `load_bearing_for_m217: true` **and** STALE/UNVERIFIABLE. Three of them are inside **`overview.md` itself**, and one would send the implementer down a wrong build path:

| # | Where | The lie | Consequence if built on |
|---|---|---|---|
| F-4.1 | `overview.md:53-55` + `roadmap.md:164` + `decisions.md:45-49` | "jobsimulation prints CLI help + exits — **no run/serve subcommand**; investigate a compose-command fix" | **The diagnosis is wrong.** `cmd/root.go:59-62` — `rootCmd` HAS a `RunE` that *is* the server; compose passes **no** `command:`; `Dockerfile.dev` is `ENTRYPOINT ["./application"]` with no CMD. Adding `command: serve` would **actively break it** (`unknown command "serve"` → real exit 1). The real cause is the `$HOME/.aws/credentials` bind mount (§5-4). |
| F-2.3 | `overview.md:25-29` | the leaked cockpit served "a stale `cockpit-manifest.json` (**dead `__clerk_identity` keys**)" | **The mechanism does not exist.** `stack-seeding/seeders/cockpit.go:145-168` — `CockpitHero` carries **no clerk id**; `key` is the stories.yaml id (`maya-thriving`), resolved against the FRESH roster bind-mounted each bring-up (`gen_injected_override.py:406-419`). Attributing the 1–2 min login to dead clerk-ids will send M218 hunting a phantom. |
| F-5.1 | `overview.md:32-33` | "**All three** snapshot replays SKIPPED (cold cache)" | Only **two** are cache misses (rc=5). The directus one is **rc=4** — the run was `DEMO_NO_LOCAL_CONTENT=1`, so no `directus` schema exists at all. **Priming the cache alone will NOT fix it.** |

Plus, outside the overview: `cockpit-spec.md:76-82` claims teardown reaps the cockpit (it does not — pid-only, and the pidfile is clobbered/deleted); `rosetta_demo.md:20-24` records the rext pin as `v1.10.1` (SoT is `v2.2`, **six releases stale**); `snapshot-cold-start.md:104-170` offers three cold-start options of which **zero are executable on `billion`**; `tailscale-serve.md:129-139` — the fresh-VM runbook that *produced* the drifted remote clone — has an unresolved `<panorama-tag>` placeholder and **omits `git fetch --tags`**.

**(b) The demopatch BLIND AREA is PROMOTED, not a blocker — verified.** `overview.md:66-73` carries an explicit `## Delivers → knowledge/corpus` section naming **`corpus/ops/demo/demopatch-spec.md`** as the deliverable, with a required-content list. Confirmed by `ls corpus/ops/demo/` (no such file today) and `grep -rn 'demopatch' corpus/` (~40 incidental hits across 8 files, **zero specs**). This is **not** a RED cause — **provided the doc is authored as SECTION 1 of the milestone**, because scope items 2 and 3 build directly on its contract.

**How to clear RED (cheap — ~1 hour, no code):**
1. Correct `overview.md` on the three points above (jobsimulation root cause; cockpit staleness mechanism; two-not-three replay skips). ← blocking
2. Land `corpus/ops/demo/demopatch-spec.md` as **Section 1** (already scoped).
3. Backfill the six stale corpus docs per §7 — these can land alongside their scope items, not before.

Everything the implementer needs is in **§5 GROUND TRUTH FOR THE BUILD**. Build from §5, not from the docs.

---

## 2. Topic Inventory

| # | Topic | Doc under audit | Code ground truth | Status |
|---|---|---|---|---|
| T1 | The demo-patch mechanism (the sanctioned zero-platform-edit escape hatch) | **none** → `corpus/ops/demo/demopatch-spec.md` (to author) | `demo-stack/patches/{demopatch,manifest_loader.py}` + 6 manifests + `stack-injection/apply-*.sh` (×3) + `up-injected.sh:379-433,657-727` + `ensure-clones.sh:180-238` + `ant-academy.sh:48-76,171-183` | 🕳️ **BLIND-AREA (PROMOTED — has a Delivers-> owner)** |
| T2 | Presenter cockpit process/port lifecycle | `corpus/ops/demo/cockpit-spec.md` | `demo-stack/{cockpit.py,up-injected.sh:1240-1298,detach.sh,rosetta-demo:139-197}` | 🟠 **PAIRED — 6 STALE / 1 UNVERIFIABLE** |
| T3 | Demo lifecycle + teardown contract (cmd_down, registry, offset formula) | `corpus/ops/rosetta_demo.md`, `demo/README.md`, `idempotency.md`, `.claude/skills/demo-down/SKILL.md` | `demo-stack/rosetta-demo`, `up-injected.sh`, `stack-core/stack_registry.py` | 🟠 **PAIRED — 5 STALE / 1 UNVERIFIABLE / 2 ALIGNED** |
| T4 | Bring-up auto-verify safety net | `corpus/ops/verification.md` | `stack-verify/live/{autoverify.sh,verify.sh}`, `stack-verify/lib/*`, the two bring-up tails | 🟠 **PAIRED — 5 STALE / 1 UNVERIFIABLE / 7 ALIGNED** |
| T5 | Snapshot cache + cold-start capture (priming `billion`) | `corpus/ops/snapshot-spec.md` + `snapshot-cold-start.md` | `stack-snapshot/cmd/stacksnap/*`, `store/`, `pg/`, `capture/`, `replay/`, `dev-setdress.sh` | 🟠 **PAIRED — 3 STALE / 1 UNVERIFIABLE / 5 ALIGNED** |
| T6 | `jobsimulation` exits(1) (DEF-M215-04 / F13) | `corpus/services/jobsimulation.md` (+ the roadmap's stated diagnosis) | `stack-injection/gen_injected_override.py:271-306`, `docker-compose.yml:113-176`, `jobsimulation/cmd/root.go`, `internal/ai/ai.go:85-91` | 🔴 **PAIRED — root cause in the roadmap is WRONG** |
| T7 | rext clone pinning + tag/consumption discipline | `corpus/ops/rosetta_demo.md:13-26` + `CLAUDE.md` §rosetta-extensions | `demo-stack/ensure-clones.sh:56-76`, `lib/rext_tag.sh:16-28`, `.claude/skills/{demo-up,stack-secrets}/SKILL.md` | 🟠 **PAIRED — 5 STALE (incl. an ACTIVE drift injector)** |

---

## 3. Fidelity Findings

*(Numbered `Fn.m` = topic n, finding m. `LB` = load-bearing for M217.)*

### F1 — The demo-patch mechanism (BLIND AREA, promoted)

**F1.1 · BLIND · LB=yes.** *Source:* `overview.md:68-73`. *Expected:* no corpus doc. *Actual:* **CONFIRMED, and the overview UNDERSTATES it.** The full contract lives in a Python module docstring (`demopatch:1-55`) plus a **rext-side** `patches/README.md` (not corpus). The closest corpus prose is a 16-line blockquote at `frontend-tier.md:171-186` and a routing-table *cell* at `coverage-protocol.md:201`. Both name the 6 guards only as a parenthetical 6-word list. **Neither documents:** the manifest schema, the sha-gate semantics, the exit codes, the **chaining rule**, the **two non-demopatch apply vehicles**, or any re-pin runbook. *Fix:* author `corpus/ops/demo/demopatch-spec.md` from §5-0; back-link from `frontend-tier.md:171-186`, `coverage-protocol.md:201-205`, `seeding-spec.md:373`, `stories-spec.md:471`, `ai-readiness.md:182-190`, `tailscale-serve.md:348-363`, `clerkenstein.md:150`.

**F1.2 · STALE · LB=yes.** *Claim:* the 6 guards are enforced (`frontend-tier.md:179`, `coverage-protocol.md:201`). *Actual:* **ALIGNED but SHALLOW, and one detail is wrong.** All six exist (`demopatch:144-181` G1+G6, `:204-224` G2, `:228-236` G3, `:266-268` G4, `:296-327` G5, `:153-159` G6). The corpus says G1/G6 read the demo type from *"the unified registry"* — that is **half true**: `demopatch:117-141` has a **second, STRUCTURAL** signal (`_is_demo_workspace`) which is the one that **actually fires at fresh-build patch time**, because the registry has no `demo-N` row yet (its own docstring at `:101-105` says so). There is also an **unnamed 7th guard** (the apply post-condition, `demopatch:277-283`).

**F1.3 · STALE · LB=yes.** *Claim:* `ai-readiness.md:182-190` / `stories-spec.md:471-476` / `seeding-spec.md:373` — "the demo bounds the AI-readiness frozen read with the `app-aireadiness-snapshot-loadmembers` demo-patch". *Actual:* **STALE IN EFFECT.** That patch has **not applied on any recent bring-up** — pinned `pre=8d509118…` (= `ai_readiness.go` @ v1.315.0) matches **neither** v1.334.1 (`b3216968…`) **nor** v1.337.0 (`dc9e167e…`). The helper REFUSEs with **exit 2**, `up-injected.sh:717`'s `>/dev/null 2>&1` eats the reason, and the demo runs the **unbounded whole-org hydration** (the 180 s wall). Same for `app-targetrole-authz-skip` (the 76 s members grid). *Fix:* correct after the re-pin lands; meanwhile add a "gated by the manifest pin — see demopatch-spec.md#freshness" caveat.

**F1.4 · STALE · LB=no.** *Claim:* `frontend-tier.md:183-186` — "re-anchored to the current source… the M47 re-sync moved next-web to **v2.89.0**… M49 recomputed both hashes from the v2.89.0 source." *Actual:* **stale text, correct outcome.** The clone is now `v2.106.1-1-g23bdbb5db`, re-pinned again at M211 iter-15. All three next-web pins verified **OK against the live clone today**. *Fix:* drop the version-number narrative; point at the new spec so the corpus stops carrying a hash-version it cannot keep current.

**F1.5 · ALIGNED · LB=no.** `tailscale-serve.md:390` (F5 row: "two app demopatches refused | sha-drift; non-fatal") is the **one corpus row that reports the true current state**. Cross-link it to the new spec's re-pin runbook.

**F1.6 · ALIGNED (invariant holds) · LB=no.** Wiring is demo-only: a full-tree grep for `demopatch|apply-app-authz-skip|apply-app-aireadiness|apply-ant-academy-dev-origins` hits **only** `demo-stack/{up-injected.sh,ensure-clones.sh,ant-academy.sh,patches/,tests/}` and `stack-injection/{README.md, the 3 apply-*.sh}`. **Zero** hits in `dev-stack/`, `stack-core/`, `stack-seeding/`, `stack-snapshot/`, `stack-verify/`, `playthroughs/`, `alignment/`.

**F1.7 · FALSE PREMISE (in the task brief, not the corpus) · LB=yes.** The `apply-*.sh` helpers do **not** live in `demo-stack/patches/` — they live in **`stack-injection/`**. The split is load-bearing: the helpers exist *precisely because* demopatch's G1/G6 path firewall **correctly REFUSES** their targets (an out-of-workspace build-scratch clone; a natively-run ant-academy). The spec must document **three** apply vehicles.

### F2 — Presenter cockpit lifecycle

**F2.1 · STALE · LB=yes.** *Claim:* `cockpit-spec.md:76-82` — the cockpit "is torn down with it (its PID is recorded in `<stack>/cockpit.pid`)… `down` leaves no orphan http.server". *Actual:* **FALSE three ways.** (1) `rosetta-demo:152-156` reaps **by PID only, never by port**; the `kill` result is discarded and the pidfile is `rm -f`'d **unconditionally**. (2) `detach.sh:31-32` writes the pidfile **unconditionally at launch, with no liveness check** → a re-up **clobbers** the predecessor's pid. (3) The pidfile lives **inside the rext consumption clone** — which M217 item 6 re-clones. **LIVE PROOF on `billion` (read-only):** `python3 …/cockpit.py --port 17700 --host 0.0.0.0` **pid 83214, started Sat Jul 11 18:01:55**, still `LISTEN 0.0.0.0:17700`, **while `find ~/panorama -name cockpit.pid` returns nothing** and every container is down. Unreapable.

**F2.2 · STALE · LB=yes.** *Claim:* `cockpit-spec.md:84-93` "D9 single source — the menu can never drift from the data". *Actual:* true of the **file**, false of the **process**. `cockpit.py:539` reads the manifest **once** into memory and `:565`/`:456` capture it in the handler closure. The file *is* regenerated every bring-up (`up-injected.sh:1254`), but a **leaked** cockpit never re-reads it. *Fix:* qualify to "single-sourced **at launch**".

**F2.3 · UNVERIFIABLE · LB=yes.** *Claim:* `overview.md:25-29` — the stale cockpit served "dead `__clerk_identity` keys". *Actual:* **unsupported by the code.** `stack-seeding/seeders/cockpit.go:145-168` — `CockpitHero` = `{key,name,role,vantage,vantage_label,trajectory,annotation,jump_to,jump_label}`; `key` is the stories.yaml id (`cockpit.go:22-26`), **not** a clerk id. The `?__clerk_identity=<key>` handshake (`cockpit.py:83-93`) is resolved by the fake-FAPI against the **fresh** roster, bind-mounted read-only per bring-up (`gen_injected_override.py:406-419`). With an unchanged preset the stale links **still resolve**. **The REAL stale-cockpit hazards** are: (a) baked `--app-base`/`--fapi-host`/`--academy-base` **scheme+host** from a different run mode (billion's orphan carries `--app-base https://billion.taildc510.ts.net:13000`) → CTAs point at an unreachable origin / mixed content; (b) a changed preset → key/`jump_to` drift; (c) a stale `[Download seed manifest]`; (d) **the new run's cockpit is DEAD while the log says it is serving**. *Fix:* correct `overview.md` and **re-measure the login latency on a green stack** before attributing it.

**F2.4 · STALE · LB=yes.** *Claim:* `cockpit-spec.md:80-82` "the serve is non-fatal". *Actual:* non-fatal ✅ **but also SILENT** — `up-injected.sh:1295` logs `presenter cockpit serving on http://$HOST:$cockpit_port` **unconditionally**. No `/healthz` probe (cockpit.py *serves* one; nobody calls it), no exit-code check, no log scrape. **The bring-up reports a crashed cockpit as serving.** Contrast `ant-academy.sh:188-196`, which *does* poll `kill -0`. *Fix:* "non-fatal **but verified**".

**F2.5 · STALE · LB=no.** *Claim:* `cockpit-spec.md:76-78,170-175` — reachable at `http://localhost:$((7700+N*10000))`. *Actual:* port formula ALIGNED (`up-injected.sh:1241`); **bind host undocumented** — `cockpit.py:534` defaults `127.0.0.1`, but `up-injected.sh:1284` passes `--host 0.0.0.0` whenever `--public-host` is set (`BIND_HOST`, `:66`). Matters for the reap: an orphan on `0.0.0.0` blocks **both** a loopback and a public re-bind.

**F2.6 · STALE (gap) · LB=yes.** *Claim:* none — the doc is silent on port-in-use. *Actual:* `cockpit.py` has **zero** port-in-use handling. `:567` `httpd = ThreadingHTTPServer((a.host, a.port), handler)` is **outside any try/except** → unhandled `OSError: [Errno 98]` traceback → exit 1. `allow_reuse_address=1` comes free from the stdlib base class and is **irrelevant** (SO_REUSEADDR does not let you bind over an **active** listener on Linux). `grep -rn 'lsof|fuser|pkill|ss -|netstat'` over `demo-stack/ stack-injection/ stack-core/` → **NO port-reap primitive exists anywhere in rext.**

**F2.7 · STALE · LB=yes.** *Claim:* `cockpit-spec.md:168-180` implies a re-`/demo-up N` is safe. *Actual:* `up-injected.sh` has **no pre-teardown and no port preflight** — its preflights (`:147` RAM, `:181` disk, `:209` host tools) cover none of this. It goes straight to `docker compose up -d` (`:944-946`) under `set -euo pipefail` → a leaked listener **aborts the whole bring-up** (the reported `0.0.0.0:18082`).

### F3 — Demo lifecycle + teardown

**F3.1 · STALE · LB=yes.** *Claim:* `demo/README.md:30,55` + `demo-down/SKILL.md` — "/demo-down N → tear it all down (AND reap the native cockpit process)". *Actual:* pid-only reap; see F2.1. **No port is ever reaped.**

**F3.2 · STALE · LB=yes.** *Claim:* `rosetta_demo.md:101-128` + `demo-up/SKILL.md:144-145` — "N is allocated from the unified dev+demo registry; teardown frees the slot; the registry is the source of truth for a stack's ports". *Actual:* **`up-injected.sh` NEVER calls `allocate` / `release` / `set-ports`.** The only `stack_registry` call in the entire bring-up is `set-host` at `:953`, and only under `STACK_PUBLIC_HOST`. `N` comes straight from argv (`:16`). A `/demo-up` stack enters the registry only via the `docker ps` **adoption** path (`stack_registry.py::_reconcile`), which writes `"ports": []`. Verified: the unified registry is `{}` **on both boxes**. ⇒ **DO NOT build the port-reap on the registry.**

**F3.3 · STALE · LB=yes.** *Claim:* `demo-down/SKILL.md` — teardown = containers + network (+ data with `--purge`) + free the slot. *Actual:* ALIGNED-but-incomplete: the real `cmd_down` (`rosetta-demo:139-197`) is an **8-step** contract (guard_n → academy stop → cockpit pid-kill → tailscale-serve reset → compose down → data rm → image rm → reg_del + ureg_release). SKILL.md mentions **none** of: the academy stop, the cockpit reap, the tailscale-serve reset, or the image removal.

**F3.4 · UNVERIFIABLE · LB=yes.** *Claim:* `idempotency.md:36-56` — "safe to retry". *Actual:* true for the **three DB steps only**. Nothing in the corpus covers **re-running the bring-up over a stack whose containers are gone but whose host-native listeners survive** — which is exactly the failure mode. *Fix:* add that section once the port-reap lands (it is the guard that *makes* the claim true).

**F3.5 · ALIGNED · LB=yes.** Offset formula `P + N·10000` (`up-injected.sh:56`, `rosetta-demo:20`, `gen_injected_override.py:293-295`). Host-native peers follow the same rule **outside** compose: cockpit `7700+OFFSET`, academy `3077+OFFSET`.

**F3.6 · ALIGNED (with gates) · LB=yes.** F12 tailscale-serve reset on teardown (`rosetta-demo:166-176`) — gated on `stacks/demo-N/tailscale-serve.sh` existing **and** `tailscale` on PATH. `up-injected.sh:986-1004` pre-resets on re-up but — **per its own ADV-1 comment at `:995-1000`** — runs **after** `docker compose up` (`:944`), so it **cannot** prevent the first bind conflict.

**F3.7 · STALE · LB=yes.** Both `app` perf pins are drifted (see F1.3); the refusal text is discarded at `up-injected.sh:701` and `:717`.

**F3.8 · STALE · LB=yes.** rext consumption clones drifted (see F7.1/F7.4).

### F4 — Auto-verify

**F4.1 · ALIGNED · LB=no.** Default-on at both tails; `DEMO_NO_VERIFY=1` / `DEV_NO_VERIFY=1` (`up-injected.sh:1327`, `dev-stack:153`).

**F4.2 · ALIGNED · LB=yes.** Non-fatal is REAL and belt-and-braces: `autoverify.sh:32` (`set -uo pipefail`, deliberately no `-e`), `:154` unconditional `exit 0`, plus `|| true` at both call sites. **M217 must NOT make it fatal — add signal, not aborts.**

**F4.3 · STALE · LB=yes.** *Claim:* `verification.md:161-174` — "the ⚠⚠ block tells you… and points you at /test-platform". *Actual:* **the failing probe names are SWALLOWED.** `autoverify.sh:138` `if "$HERE/verify.sh" >/dev/null 2>&1; then … else warn "verify live reported failing probe(s)…"`. `verify.sh` prints every `✗ <service> <detail>` (`:37`, `:68`) and the summary `✗ N probe(s) failed` (`:86`) — **all discarded**. This is the **same swallow class** as the demopatch REFUSE that scope item 2 exists to fix. Fix both with one discipline.

**F4.4 · STALE · LB=yes.** *Claim:* `verification.md:166-170` — "N check(s) FAILED" is the number of failing checks. *Actual:* **misleading.** `warnings` (`autoverify.sh:69-70`) counts *autoverify-level* checks; the **entire** `verify live` run collapses to **exactly 1** warning (`:141`). "1 check(s) FAILED" means "**≥1** probe failed".

**F4.5 · UNVERIFIABLE · LB=yes.** *Which* probe failed on billion? `coldrun2.log:348-355` **cannot say** (F4.3). Both cheap-wins **passed** (`✓ backend /api/health 200 on :18082`, `✓ sentinel.casbin_rules = 1150`) ⇒ the failure is inside `verify live`. By elimination + F13/DEF-M215-04: the **jobsimulation liveness probe** at `:18400` (`services.sh:42`) — container exits(1) → curl code `000` → down → `fail_count≥1` → `verify.sh` exit 1. **This is an inference, not a quoted probe line** (the stack is torn down; `docker ps -a` on billion is empty). Confirm after the next bring-up with `STACK_PROJECT=demo-1 STACK_OFFSET=10000 …/stack-verify/live/verify.sh`.

**F4.6 · STALE · LB=no.** `verification.md:22-25,133-152` still says "out of scope: the frontend tier — the frontends don't exist yet; M19 adds them". M19 **landed**: `services.sh:55-56` add `next-web-app:3000` + `studio-desk:9000`; `up-injected.sh:1329` scopes them in.

**F4.7 · STALE · LB=no.** `dev-stack/dev-stack:157` still passes **`skiller`** in `verify_svcs` — not a `SERVICES` row since the v2.1 merge; silently dropped. Related: `target.sh:20-22` **claims** "a name not in the array is ignored **with a warning**" — **no such warning exists**.

**F4.8 · STALE · LB=no.** `verify.sh:45-46` comment claims readiness runs "only if liveness passed". **False** — `run_readiness` (`:50-69`) gates only on scope (`:57`), so a dead service **double-counts** (liveness ✗ + readiness ✗).

**F4.9 · STALE · LB=yes.** `verification.md:51-56` (FIX B, session-detached daemons) — confirmed for the cockpit, but the launch is **fire-and-forget**; see F2.4.

**ALIGNED (no action):** cheap-win asserts (`verification.md:94-106` ↔ `autoverify.sh:77-99`); the M22 Directus cheap-wins (`:108-121` ↔ `autoverify.sh:105-134`); the offset/project/scope model (`:64-92` ↔ `lib/target.sh`); the `$DEVDIR → $STACK_ROOT` fix (`:155-159` ↔ `repos/run.sh:107`, `census/inventory.sh:75`); the Directus boot health-gate (`:122-131` ↔ `dev-setdress.sh:133-152,245`).

### F5 — Snapshot cache + cold start

**F5.1 · STALE · LB=yes.** *Claim:* `overview.md:32-33` "**All three** snapshot replays SKIPPED (cold cache)". *Actual:* **only two are cache misses.** `coldrun2.log:288` taxonomy → **rc=5**; `:302` sim-embeddings → **rc=5**; `:296` directus → `probe stack schema: pg: schema "directus": schema has no columns (empty digest)` → **rc=4**, *not a cache problem* — the run was `DEMO_NO_LOCAL_CONTENT=1` (`:115`, `:263`, `:346`), so `provision_directus_step` never ran. **M217 must fix TWO things:** prime the cache **and** run with local content on.

**F5.2 · STALE · LB=yes.** *Claim:* `snapshot-cold-start.md:104-170` gives three cold-start options (Option 1 dump-ingest, Option 2 primary-read via `~/.pgpass`, Option 2b the postgres-MCP DSN). *Actual:* **NONE is executable on `billion`.** `ls ~/.pgpass` → no such file; `ls ~/.claude.json` → no such file; no staging `pg_dump` on the box; no prod credential anywhere in `~/panorama`. **The only documented cold-start path is unavailable on the target box.**

**F5.3 · UNVERIFIABLE (blind) · LB=yes.** *Claim:* none — no corpus doc anywhere sanctions **shipping a warm cache to another box** (`grep -rn 'rsync|copy the cache|ship the cache|remote box'` over both snapshot docs → **0 hits**). The corpus sanctions only capture-over-DSN and the M211 cache-**migration** (`snapshot-spec.md:180-191`). Yet cache-transport is **the only mechanically available option on billion**. *Fix:* document it (§5-5); the bytes are provably safe to move.

**F5.4 · STALE · LB=no.** `snapshot-cold-start.md:130-131` (Option 1) captures only `taxonomy` + `directus` — the bring-up replays **three** (`dev-setdress.sh:311`), leaving `sim-embeddings` cold (→ `cms.similarities = 0` → `/library/ai-simulations` empty). Option 2b (`:161-163`) *does* list all three.

**ALIGNED (no action):** surface registry (`surfaces.go:34-39`, 4 surfaces); replay is cache-first and **never** captures (`main.go:297-414`, `replay.go:7-18`); cache layout + store-root resolution (`store.go:39,145-146`, `main.go:142-171`); the schema-digest key (`pg.go:228-236`, `capture.go:136-155`); the exit-code set 0/1/3/4/5 (`main.go:54-73`).

### F6 — jobsimulation exits(1)

**F6.1 · STALE (the roadmap's root cause is WRONG) · LB=yes.** *Claim:* `decisions.md:45-49` / `roadmap.md:164` / `overview.md:54-55` — "printed CLI help + exited (**no run/serve subcommand**)"; "investigate a **compose-command fix**". *Actual:* **FALSE on both counts.** (a) **No compose service passes a `command:`** — the jobsimulation block (`docker-compose.yml:113-176`) has none, and `Dockerfile.dev` ends `ENTRYPOINT ["./application"]` with **no CMD** → the binary is invoked with **zero args**. (b) **`rootCmd` HAS a `RunE` and it IS the whole server** — `cmd/root.go:59-62`. With no args, cobra runs it. The four subcommands (`aggregate`, `clone-session`, `test-command`, `validate`) are optional. `RunE` has been on `rootCmd` since 2024-05-07 (`a3f6b442`). **"It printed help" is cobra's default usage-on-RunE-error behavior** (`SilenceUsage`/`SilenceErrors` are never set): RunE returns err → cobra prints `Error: <msg>` **+ the full usage block** → `root.go:404-407` `os.Exit(1)`. ⚠️ **Adding `command: serve` would ACTIVELY BREAK IT** (cobra `legacyArgs` → `unknown command "serve"` → real exit 1).

**F6.2 · ROOT CAUSE (undocumented anywhere) · LB=yes.** `docker-compose.yml:171-172` mounts `$HOME/.aws/credentials:/root/.aws/credentials:ro` — the **only** aws bind in the whole compose file, on jobsimulation **only** (added 2024-12-16, `6daa67e`). When the host path is absent, **Docker auto-creates it as an empty DIRECTORY.** Proven live: `ssh devops@billion 'ls -la ~/.aws/'` → `drwxr-xr-x 2 root root … credentials` (a **root-owned directory**, ctime = the bring-up). The container then sees a **directory** where an ini file must be → `cmd/root.go:183 ai.NewAIManager(…)` → `internal/ai/ai.go:85-91 config.LoadDefaultConfig(…)` → `can't load AWS config` → `root.go:195-197 "can't init AI"` → cobra usage → **exit 1**. **Reproduced against the exact module version** (`aws-sdk-go-v2/config v1.32.25`): `failed to load shared config file, <path>, read all: read <path>: is a directory` (`shared_config.go:797-805` → `internal/ini@v1.8.6/ini.go:24-44` — `os.Open` on a dir **succeeds**, so it is *not* wrapped as `UnableToReadFile` and *not* skipped; `io.ReadAll` then EISDIRs). **Control:** with the path simply **absent**, `LoadDefaultConfig` returns `err = nil`.

**F6.3 · STALE · LB=yes.** *Claim:* `overview.md:12,54-55` / `roadmap.md:291` — "AI-Simulations is dead in **every** demo today". *Actual:* **overstated.** It dies only where `~/.aws/credentials` is **absent-or-a-directory**. This laptop has it as an existing **empty FILE** (`-rw------- 0 Nov 4 2025`) → parses as an empty ini → no error → jobsimulation boots (consistent with v2.0's 10 GREEN playthroughs, which include `pt-aisim-chat-launch`). **It is dead on every FRESH host** (billion, any Linux VM, any CI runner) — which is exactly the reproducible-demo claim M217 must make good on. Regression window: mount added 2024-12-16 (`6daa67e`); the fatal `LoadDefaultConfig` entered jobsimulation 2025-06-09 (`fce0b643`).

**F6.4 · STALE · LB=no.** `corpus/services/jobsimulation.md:53,69` lists **Roadrunner** as a direct RPC dependency and a Redis-Streams producer. *Actual:* `cmd/root.go:172-175` builds `runner.NewRunnerManager(JUDGE0_API_KEY, JUDGE0_BASE_URL)` — **Judge0 direct, in-process** (`internal/runner/`). There is **no roadrunner client**; `ROADRUNNER_RPC_ADDR` (`docker-compose.yml:148`) is a **dead env var**. Subscribers are exactly two (`root.go:284-285`): its own stream + `cms`.

**F6.5 · STALE (blind spot that CAUSED the misdiagnosis) · LB=yes.** `corpus/services/jobsimulation.md:22-33` says only "`cmd/` — Entrypoints". It never says the service is a **cobra CLI whose ROOT command IS the server**. Worse, the platform repo's own `stack-demo/jobsimulation/CLAUDE.md` ships `go run . serve  # Start service locally` — **a subcommand that does not exist**. That line is almost certainly the origin of the "no run/serve subcommand" theory. (Platform repo = READ-ONLY; do **not** fix it there.)

**F6.6 · ALIGNED (the fix vehicle exists) · LB=yes.** `gen_injected_override.py::build_lines` (L271) already emits `ports: !override` (`:293`), `volumes: !override` (`:296-299`), and **`build: !reset null` for jobsimulation itself** (`:302-306`, since `INJECTED` at `:17` contains it). ⇒ a jobsimulation-scoped `volumes:` reset is a **2-line addition** to a loop that already runs over that service. **No demo-patch, no sha-pin, no platform edit, no escalation.**

### F7 — rext clone pinning

**F7.1 · STALE · LB=yes.** *Claim:* `rosetta_demo.md:20-24` — "current v1.10b 'fit-up' pin: **`v1.10.1`**". *Actual:* `.agentspace/rext.tag` = **`v2.2`** (4 bytes `7632 2e32`, **no trailing newline**). The doc is **six releases stale**. A fresh box following it verbatim gets a pre-Playthroughs, pre-panorama pin.

**F7.2 · STALE — ACTIVE DRIFT INJECTOR · LB=yes.** *Claim:* `rosetta_demo.md:15-26` "The pin is a file, not prose… there is now exactly ONE read path" (M49 #1). *Actual:* **FALSE.** `.claude/skills/stack-secrets/SKILL.md:75` **hardcodes a stale tag and ignores the SoT**: `git -C "$SECDIR" fetch --tags --quiet && git -C "$SECDIR" checkout --quiet stage-door-m30` — where `$SECDIR = stack-demo/rosetta-extensions`, **the same clone `/demo-up` pins**. And `/demo-up` **invokes `/stack-secrets`** as an auto-provision step (`rosetta_demo.md:29-32`). ⇒ **A correctly re-pinned clone gets force-moved BACK to a v1.6-era tag on the very next bring-up.** Re-pinning without fixing this line is a no-op within one run.

**F7.3 · STALE · LB=yes.** *Claim:* the pin is "read by BOTH the `/demo-up` skill (it **checks the consumption clone out** at this tag) and `ensure-clones.sh`". *Actual:* the **WARN** half is real code (`ensure-clones.sh:64-76`, **explicitly non-fatal** per its own comment at `:61-63`). The **CHECKOUT** half exists **only as English prose** in `.claude/skills/demo-up/SKILL.md:34-38`, executed at an LLM agent's discretion. **NO code anywhere checks out `.agentspace/rext.tag`.** The discipline failed on **both** boxes; billion has been printing its own drift warning on every single bring-up and nothing stopped.

**F7.4 · ALIGNED (DATA-LOSS RISK: **PROVEN SAFE**) · LB=yes.** See §6 (a) — the remote dirty tree is a **strict subset** of `v2.2`. Zero content would be lost.

**F7.5 · STALE (cosmetic) · LB=no.** `overview.md:85` cites `up-injected.sh:669,679` for the scratch-clone checkout. Actual: **`:671`** (`for-each-ref … --count=1 'refs/tags/v*'`) and **`:681`** (`checkout --quiet -f`). *The overview's substantive point is CORRECT and important:* ref-pinning the rext clone does **not** stop demo-patch sha-rot. (The sibling citations `:701,717` **are** exact.)

**F7.6 · STALE · LB=no.** `CLAUDE.md` + `dev-up/SKILL.md:99-102` + `stack-seed/SKILL.md:27` assume `stack-dev/rosetta-extensions` exists. **It does not exist on this box.** Any pin-check preflight must **tolerate an absent dev clone**, not hard-fail.

**F7.7 · STALE · LB=yes.** `tailscale-serve.md:129-139` (the fresh-Linux-VM runbook) uses an unresolved `<panorama-tag>` placeholder **and omits `git fetch --tags`** before the checkout. **That omission is literally how the remote box ended up on a bare sha** (`41a28aa`) instead of a tag — its reflog shows `clone` → `checkout: moving from main to 41a28aa`.

---

## 4. Completeness Gaps (critical undocumented behaviors)

**G1 — THE PIN CANNOT BE CORRECT ON BOTH BOXES AT ONCE.** `up-injected.sh:671` checks the build-scratch clone out at *"the newest `v*` tag in whatever this box last fetched"*. Local `stack-demo/app` = **v1.334.1**; billion's = **v1.337.0**. `internal/roles/roles.go` is **byte-identical** across those tags (`88f25f06…`) ⇒ one pin works everywhere. `internal/workforce/ai_readiness.go` is **NOT** (`b3216968…` @1.334.1 vs `dc9e167e…` @1.337.0) ⇒ **any static file-sha pin for it is wrong on one of the two boxes the moment it is committed.** This is the mechanical proof behind BD-3. **A one-shot re-pin CANNOT close M217 for `ai_readiness`.**

**G2 — The two `app` patches have NO live-clone drift test — that is why the drift shipped.** `test_demopatch.py:456` and `:541` **do** assert the next-web pins against the real clone; `test_ant_academy.py:370-371` does the same. But the `app` manifests get only **manifest-internal** self-consistency tests (`:506`; and there is **no aireadiness manifest test at all**), and `:581` deliberately asserts the helper **REFUSES a synthetic file** (rc 2) — it never checks the pin against a real clone.

**G3 — THREE opt-out env vars, not one.** The corpus mentions only `DEMO_NO_PATCH=1` — which gates **only** the three next-web demopatches (`up-injected.sh:399`). The two app patches have **independent, undocumented** escapes: `DEMO_NO_AUTHZ_SKIP=1` (`:700`) and `DEMO_NO_AIREADINESS_LOADMEMBERS_BOUND=1` (`:716`).

**G4 — The two app patches are NEVER reverted.** No `revert` verb exists in either helper. Cleanliness is guaranteed **structurally** by `up-injected.sh:681`'s `git checkout -f "$tag"` on the gitignored throwaway scratch at the start of every build. ⇒ **The "G5 self-revert" contract does NOT apply to the app patches.** A reader will assume a revert exists.

**G5 — `ensure-clones.sh`'s R1 pristine-ing list is INCOMPLETE.** `:191-195` lists **3 of the 6** manifests; missing `next-web-public-website-url` and `app-aireadiness-snapshot-loadmembers`. Today it's masked by luck (`--force-pristine` = `git checkout -- urls.ts`, which restores the whole file; and the app scratch self-pristines). Latent hole. Also `app-targetrole-authz-skip` in that list is a **no-op** — its `repo: app` resolves to `stack-demo/app` (the SOURCE clone), which is never patched.

**G6 — A 7th, unnamed guard.** `demopatch:277-283` (+ every helper) recomputes `sha256(patched-in-memory)` and refuses to write unless it equals `post_sha256` **and** contains `post_marker`. Catches a manifest-authoring error **before a byte hits disk**.

**G7 — THE LEAK LOOP (nowhere documented).** Run A launches cockpit `P_A` → `cockpit.pid = P_A`. Run B launches: `detach.sh:31-32` **overwrites** `cockpit.pid` with `P_B` *before python even binds*; python dies at `cockpit.py:567` (EADDRINUSE). Net: `P_A` is orphaned **AND its pid record destroyed**. A later `down` kills the **dead** `P_B`, `rm -f`s the pidfile, prints *"stopped the presenter cockpit"*, and `P_A` serves forever.

**G8 — The pidfiles live INSIDE the rext consumption clone** (`demo-stack/stacks/demo-N/{cockpit,ant-academy}.pid`, `rosetta-demo:26`). **Re-pinning / re-cloning rext — which M217 item 6 does — DELETES the pidfiles of any running host-native listener.** Same for `tailscale-serve.sh`, on which the teardown's serve-reset is gated (`rosetta-demo:166`) ⇒ an rext re-pin **disarms the serve reset too**. **ORDER: reap FIRST, re-pin SECOND.**

**G9 — ZERO leaked-port detection exists in rext.** Repo-wide grep for `EADDRINUSE | "Address already in use" | lsof | fuser | ss | netstat | SO_REUSEADDR | allow_reuse` → **no hits in any script**. M217 is **adding** a capability, not fixing a broken one.

**G10 — THE VERIFY SET IS BLIND TO EVERY M217 DEFECT CLASS.** `grep -rniE 'demopatch|patch|snapshot|cockpit|taxonomy|academy'` over `stack-verify/{live,lib,repos,census}` → **exactly one hit**, an unrelated comment. So: **no** check that a demo-patch applied (this is *literally how it rotted for releases*); **no** check that a snapshot replay succeeded; **no** check that the cockpit is up; **no** probe on ant-academy or on Clerkenstein fake-fapi/fake-bapi (**a dead fake-fapi means nobody can log in — and verify stays green**).

**G11 — NO artifact is written by autoverify.** stderr only. No `<stack>/autoverify.log`, no JSON, no exit-code file. `/demo-up` **exits 0 on a red verify** and still prints `UP. Clerk-free demo-N is live.` Nothing downstream (M218, the skill, the registry) can machine-read whether the stack came up green.

**G12 — Cobra usage-on-error semantics.** jobsimulation is the only platform service whose `main` is a cobra root command. **Every future "it printed CLI help" report from this service is an INIT ERROR** — the real message is the **FIRST line, above the help**.

**G13 — Host-dependent failure class.** The demo is green on a dev mac (empty `~/.aws/credentials` FILE) and red on a fresh Linux VM (docker-created DIRECTORY). Same class as DEF-M215-02 (cold cache). *"Works on my box"* is **structurally load-bearing** for M217's green-demo gate.

**G14 — The same jobsimulation bug exists on the DEV path.** `stack-core/gen_override.py:117-128` rewrites `volumes` **only** for `postgresql`. A `/dev-up` on a Linux box without `~/.aws/credentials` has an identically dead jobsimulation. Out of M217's stated scope; **3-line mirror fix**. Flag it or it resurfaces.

**G15 — AUTO-PROVISION IS GATED ON A VIRGIN DIRECTUS SCHEMA.** `autoprovision.go::userCollectionCountSQL` counts non-`directus_%` base tables and provisions **only when that count is 0** (the DDL is intentionally **not** idempotent). A half-provisioned schema on a re-run falls through to **rc=5** and does **not** self-repair. **Any M217 re-run on billion must start from a purged stack.**

**G16 — The directus digest is WHOLE-SCHEMA** (`capture.go:145-155`: `CapturesStructure ⇒ VersionTables()=nil`), so it is sensitive to the **Directus image version's** bootstrap tables. The cached `ea2e187a…` was produced against `directus/directus:11.6.1` — exactly `dev-setdress.sh:173`'s default. **Do not bump that image on billion**, or the digest drifts with no cure but a real recapture.

---

## 5. GROUND TRUTH FOR THE BUILD

> The implementer works from this section. Everything is code-cited against the authoring clone `.agentspace/rosetta-extensions` @ `v2.2` (HEAD `39e8013`). **HARD RULE: no platform repo is edited by anything below.** All surfaces are `rosetta` (corpus + `.claude/skills/`) and `rosetta-extensions`.

---

### 5-0. SECTION 1 (FIRST) — Author `corpus/ops/demo/demopatch-spec.md`

Everything in scope items 2 and 3 depends on this contract being written down. Author it **before** touching code.

#### What it is
`demo-stack/patches/demopatch` — a **378-line stdlib-only Python3 CLI** (no PyYAML; the rext supply-chain rule) that applies a rext-owned, **content-anchored** source patch to a demo's own **gitignored clone** before an image build and reverts it after, so the **image** carries the fix and the clone is left git-clean. Canonical `anthropos-work` repos are never touched. Six guards make that mechanical. Born M42m/v1.10 because next-web's `STUDIO_URL` had no `NEXT_PUBLIC_*` override seam and the demo's Studio nav link ejected presenters to prod.

#### The guards — what each ACTUALLY enforces

| Guard | Site | Enforcement |
|---|---|---|
| **G1 — hard path-assert (demo-clone only)** | `demopatch:144-181` `assert_demo_clone_path` | Three sub-asserts, all must pass: **(a)** `:161-168` `realpath(target)` must be inside `realpath(workspace_root)` — realpath resolves `..` **and every symlink** (symlink-escape guard); **(b)** `:169-173` exact-path: `normpath(target) == normpath(join(ws, manifest['repo'], manifest['path']))` — no glob, no `..`; **(c)** `:174-180` the *manifest-derived* path is re-realpath'd and must also be inside the workspace (kills a `repo: ../stack-dev/...` manifest). Refusal prefix `G1 REFUSE:`. Upstream belt: `manifest_loader.validate:115-118` rejects `..`/absolute at **load** time (exit 2). |
| **G2 — pre-patch drift-refuse + single-occurrence anchor** | `demopatch:204-224`, classifier `:185-201` | `sha256(WHOLE FILE)` must be ∈ `{pre_sha256, post_sha256}`; any third hash → `drifted` → `G2 REFUSE`. `post_sha256` alone is **not enough** to count as `patched` — `_classify:194-200` **also** requires `post_marker in body` (a hash collision on a partial apply cannot masquerade). If `pristine`, the anchor must occur **EXACTLY ONCE** (`:217-223`): `n==0` → REFUSE ("content drift"); `n>1` → REFUSE ("occurs N× (ambiguous) — refusing to choose a hunk"). Missing file → `G2 REFUSE` (`:207-208`). |
| **G3 — never-commit / working-tree-only** | `demopatch:228-236` + `test_demopatch.py:93-119` | Post-write, `git diff --cached --quiet -- <rel>` must exit 0. Non-zero → `G3 REFUSE`, and `cmd_apply:288-291` **self-heals by reverting its own write** then re-raises. Structural proof: the test **greps the demopatch source** for any mutating git verb; the only `git checkout` is the `-- <path>` working-tree form, isolated to one function (`_git_checkout_path:324-327`) precisely so the grep can whitelist it. |
| **G4 — idempotent re-apply** | `demopatch:266-268` | The demo clone **persists** across `/demo-up`. Clean already-patched target (`==post_sha256` AND marker present) → `already applied (G4 idempotent no-op)` → **exit 0**, no write. |
| **G5 — content-anchored self-revert** | `demopatch:296-337` | Default: reverse swap `replacement → anchor` (once), then **re-assert** `sha256(restored)==pre_sha256`, else REFUSE (`:317-319`). Already-pristine → no-op (`:308-309`). Neither pre nor post → REFUSE, "manual drift; refusing to guess" (`:310-313`). `--force-pristine` → `git checkout -- <rel>` (working-tree restore, **not** a history op). `cmd_revert` runs G1+G6 **first** (`:333`). |
| **G6 — demo-only scope** | `demopatch:153-159` + `:99-141` | Two independent conditions: **(i)** `manifest['scope'] != 'demo'` → REFUSE (also hard-rejected at load, `manifest_loader.py:113-114`, exit 2). **(ii)** `_demo_present(ws)` = `_is_demo_workspace(ws) OR _registry_has_demo()`. `_is_demo_workspace` (`:117-134`) is the **STRUCTURAL** signal and **the one that actually fires at fresh-build patch time**: `realpath(ws)` must equal *this demopatch binary's own* clone-set workspace (`_HERE/../../..`), AND `<ws>/rosetta-extensions/demo-stack/stacks/` must exist. `_registry_has_demo` (`:99-114`) consults the unified registry — **its own docstring at `:101-105` says this is NOT reliable at fresh-build time** (no `demo-N` row yet). |
| **G7 (unnamed, real)** | `demopatch:277-283` | After the in-memory swap and **before** writing: `sha256(patched)==post_sha256` else REFUSE; `post_marker in patched` else REFUSE. |

`demopatch:37` states the invariant: **no write path bypasses G1+G2** (`cmd_apply:263-264` runs both before anything is written).

#### Manifest schema (`manifest_loader.py`)
A deliberately tiny **strict YAML subset** (`parse:50-103`): top-level `key: scalar` + `key: |` literal blocks (`|`, `|-`, `|+` all accepted, `:71`), dedented by the first non-empty body line's indent. Nested maps / flow collections / anchors → `ManifestError`.

**All TEN keys are MANDATORY — there are no optional keys** (`REQUIRED`, `:33-34`; present-but-empty also fails, `:107-109`):

| key | notes |
|---|---|
| `id` | by convention `patches/<id>/<id>.yaml` |
| `repo` | clone dir under the workspace root; `..`/absolute rejected (`:115-116`) |
| `path` | file inside that clone; same rejection (`:117-118`) |
| `pre_sha256` | sha256 of the **WHOLE PRISTINE FILE**. **64 lowercase hex** (`_HEX64`, `:32`, `:110-112`) — uppercase rejected |
| `post_sha256` | sha256 of the **WHOLE FILE AFTER** the single replacement. Same rule |
| `anchor` | **block scalar** — the exact pre-image hunk; must occur exactly once. Tabs survive the space-only dedent (the Go manifests rely on this) |
| `replacement` | **block scalar** — the post-image hunk |
| `post_marker` | substring that MUST be present in `replacement` (hard-rejected at load otherwise, `:119-120`). The G4 positive idempotency probe |
| `build_env` | a build-time env line the CALLER appends to the `.env.local` overlay, offset-templated (`$((9000+OFFSET))`). Stored **verbatim**; the caller expands. Source-only patches set it to a `# (no build-time env …)` comment (it is REQUIRED, so it cannot be omitted) |
| `scope` | must be literally `demo` |

**Design rule visible in every manifest:** the replacement is **behavior-identical when the env var is unset** (prepend `process.env.X ||`, keep the original as fallback). That is what lets a *dynamic* value (offset port, MagicDNS host) coexist with a *static* `post_sha256`.

#### The sha gate — precisely
- **WHAT is hashed: the WHOLE FILE**, not the anchor region (`_sha256_file:74-79`, `_sha256_text:82-83`). ⇒ **any unrelated edit anywhere in the file breaks the gate.** **This is the rot source.**
- The pin lives **only** in the manifest, checked into rext. No per-box override, no lockfile, no auto-refresh.
- On mismatch → `PatchError` → `main:370-373` writes `demopatch: G2 REFUSE: sha256(<path>) = <live> matches neither pre_sha256 nor a clean post_sha256+marker — the demo clone has DRIFTED …` to **stderr**, **exit 1**. demopatch's only failure codes: **1** (guard refuse) and **2** (manifest load / OSError, `main:354-356`).
- **NON-FATAL at the caller** by convention (M18/M19; the class docstring `demopatch:70-71` states it). `up-injected.sh` logs `⚠ … (non-fatal)` and continues (`:411,:417,:426,:704,:720`). **A refused patch = a silently-degraded demo, never a failed bring-up. That is exactly the M217 pathology.**
- **The helpers use a RICHER exit-code space** (separate scripts, `set -euo pipefail`): `0` applied-or-already-patched · `1` manifest/target missing · **`2` pre_sha256 drift** · `3` anchor count != 1 · `4` replacement was a no-op · `5` patched sha != post_sha256 · `6` post_marker absent · (`7`/`8` also in `apply-ant-academy-dev-origins.sh` for its revert path).

#### The anchor mechanism
Plain **substring** matching: `body.count(anchor)` / `body.replace(anchor, replacement, 1)` (`:220`, `:273`). No regex, no fuzz, no whitespace normalization. Byte-exact, **single-occurrence required**. Tabs matter (the Go manifests anchor on `"\t} else {\n\t\tuser := authn.UserFromContext(ctx)"`, asserted verbatim at `test_demopatch.py:513`).

#### Verbs
`demopatch <apply|revert [--force-pristine]|status|check> <workspace-root> --manifest <file.yaml>` (`:340-374`). `<workspace-root>` = the stack workspace dir (`…/stack-demo`); `repo/path` resolves relative to it (`main:359`). `status` (`:253-257`) prints `pristine|patched|drifted|absent`, **never refuses** (always exit 0). `check` (`:245-250`) = dry-run (G1+G6+G2, no write).

#### THREE apply VEHICLES (the single most under-documented fact)

| vehicle | patches | why |
|---|---|---|
| **`demopatch` (the tool)** | the 3 `next-web-app` patches | target lives **in** the demo workspace → G1/G6 pass |
| **`stack-injection/apply-app-{authz-skip,aireadiness-loadmembers}.sh`** | the 2 `app` patches | target is the **build-SCRATCH** clone `stacks/demo-N/clones/app`, **outside** the demo workspace → **demopatch's G1/G6 CORRECTLY REFUSE it**. The helpers re-implement the same guard ladder in an inline `python3` heredoc, reading **the same canonical manifest** via `manifest_loader` (manifest stays the single source of truth; only the vehicle differs). Headers: `apply-app-authz-skip.sh:5-19` |
| **`stack-injection/apply-ant-academy-dev-origins.sh apply\|revert`** | `ant-academy-dev-origins` | ant-academy runs **natively** (`next dev`), not baked into an image → the patch must **persist for the process lifetime** → apply-before-launch / revert-on-stop. Header `:6-13` |

#### next-web lifecycle (`up-injected.sh:379-433`, in `build_frontend_next_web`)
1. `:350-370` write the gitignored `apps/web/.env.local` overlay (`NEXT_PUBLIC_STUDIO_URL=$SCHEME://$HOST:$((9000+OFFSET))`, `NEXT_PUBLIC_ACADEMY_URL=…$((3077+OFFSET))`, `NEXT_PUBLIC_PUBLIC_WEBSITE_URL=…$((3000+OFFSET))`, Clerk pk) — these are the `build_env` values the patched source reads.
2. `:373-378` transient tooling `.dockerignore` (only if the repo has none).
3. **`:398` install the RETURN trap FIRST** — removes the overlays and runs `demopatch revert` for pubweb, studio, pagination (stderr → `/dev/null`). RETURN-scoped ⇒ fires on the failure/abort path too.
4. `:399-433` apply, gated on `DEMO_NO_PATCH != 1`: `studio` (`:403`) → **if it succeeded**, chain `pubweb` (`:408`); then `pagination` (`:423`, independent).
5. `:435` `docker build` (**unmodified Dockerfile; repo = build context only**).
6. return → trap reverts **LIFO (pubweb before studio)** → clone git-clean.

**THE CHAIN RULE** (`up-injected.sh:392-396`): `next-web-public-website-url` targets the **same** `urls.ts` as `next-web-studio-url`, so its `pre_sha256` (`fe15aa71…`) **IS** studio's `post_sha256`. It must be applied AFTER studio and reverted BEFORE studio, or its whole-file gate cannot match. It reads **"DRIFTED" against the pristine file BY DESIGN — do not "fix" this.** Fenced by `test_demopatch.py:541`.

#### app lifecycle (`up-injected.sh:657-727`, inject loop, `svc=app`)
`:671` resolve newest `v*` tag in the SOURCE clone → `:681` `git checkout --quiet -f "$tag"` the scratch (**discards the prior build's injections — this is why the app patches need no revert; see G4 gap above**) → `:688` `apply-authn.sh` (disarm colony for Clerkenstein) → `:690-692` Dockerfile vendor-colony fix → `:700-706` **authz-skip** (gated `DEMO_NO_AUTHZ_SKIP != 1`) → `:716-722` **aireadiness-loadmembers** (gated `DEMO_NO_AIREADINESS_LOADMEMBERS_BOUND != 1`) → `:724` `docker build -t demo-$N-app:injected`.
**Order is load-bearing: AFTER apply-authn, BEFORE the build.** A rebuild that skips the inject loop ships a backend that hits real `api.clerk.com` and rejects every Clerkenstein token.

#### Is it ever committed? NO.
G3 forbids every mutating git verb (grep-tested). Two structural belts in `ensure-clones.sh`: **R1** (`:190-208`) a pristine-ing pass at clone head (recovers from a build that crashed after apply but before the trap); **R1b** (`:214-224`) sweep a crash-left tooling `.dockerignore`; **R2** (`:226-236`) set every demo clone's **push** remote to `no-push://demo-clone-never-pushes` (a structural leak-block independent of G3; fetch untouched). ⚠ **R1's list at `:191-195` covers only 3 of 6 manifests** — see gap **G5**.

#### FULL PATCH INVENTORY + LIVE SHA VERIFICATION (every hash computed by hand, `shasum -a 256`)

| id | repo / path | what it does | pinned `pre_sha256` | LIVE state |
|---|---|---|---|---|
| **next-web-studio-url** (M42m) | `next-web-app` / `packages/core-js/src/constants/urls.ts` | prepend `process.env.NEXT_PUBLIC_STUDIO_URL \|\|` to the `STUDIO_URL` ternary → Studio nav stops ejecting to `studio.anthropos.work` | `0d4c3790…9180d73` | **OK** — matches live; anchor 1× |
| **next-web-public-website-url** (M50) | `next-web-app` / **same** `urls.ts` (**CHAINED**) | `PUBLIC_WEBSITE_URL` reads `NEXT_PUBLIC_PUBLIC_WEBSITE_URL` first → sim drill-down stays demo-local | `fe15aa71…bca5f31` | **OK BY CHAIN** — its `pre` == studio's `post`. Reads "DRIFTED" vs the pristine file **by design** |
| **next-web-members-pagination** (M46 T1) | `next-web-app` / `apps/web/src/context/InsightsContext.tsx` | `useGetOrganizationMembers` `limit: 1000 → 30` | `2771d5f1…7ca94fd` | **OK** — matches live; anchor 1× |
| **app-targetrole-authz-skip** (M46 B) | `app` / `internal/roles/roles.go` | `RoleManager.checkPermission` short-circuits `return true, nil` **before** the per-member `OrgCheckActionPermission` Sentinel RPC → members grid **76.7 s → 0.51 s**. READ-path only; mutations still enforce | `fd85159a…6ce6ce4f1` (= roles.go @ **v1.295.0 AND v1.315.0**) | **DRIFTED** — live = `88f25f06…20bcc97` at **BOTH** v1.334.1 and v1.337.0. Anchor still **1×** |
| **app-aireadiness-snapshot-loadmembers** (M51) | `app` / `internal/workforce/ai_readiness.go` | `buildResponseFromSnapshots`: unbounded `loadMembers(orgID,"")` → bounded `loadMembersByUserIDs(orgID,"",snapUserIDs)` (~199 users). **Data-identical**, pure perf (the 180 s wall) | `8d509118…c5809d25` (= @ **v1.315.0**) | **DRIFTED** — `b3216968…` @v1.334.1 but `dc9e167e…` @v1.337.0. **The file differs BETWEEN the two boxes.** Anchor still **1×** at both |
| **ant-academy-dev-origins** (M214/v2.2) | `ant-academy` / `code/next.config.js` | prepend `process.env.ANT_ACADEMY_ALLOWED_DEV_ORIGIN` to `allowedDevOrigins` → a `--public-host` demo's MagicDNS host is admitted to `next dev` | `6837cab9…7850a8e3` | **OK** — matches live; anchor 1× |

**Net: 4 OK, 2 DRIFTED (both `app`, both perf).** Anchors intact everywhere ⇒ **re-pin, not re-authoring.**

---

### 5-1. SCOPE ITEM 1 — Reap the leaked cockpit port + make `demo-down` reap the whole offset range

#### Current behavior

**Launch** (`up-injected.sh:1240-1298`):
- `:1241` `cockpit_port=$((7700+OFFSET))` (`OFFSET=N*10000`, `:56`)
- `:1254` `stackseed --cockpit-export …` **regenerates `$STACK/cockpit-manifest.json` every bring-up**
- `:1268` `app_base="$SCHEME://$HOST:$((3000+OFFSET))"` · `:1269` `fapi_host="$FAPI_HOST:$((5400+OFFSET))"` · `:1276` `--academy-base …$((3077+OFFSET))`
- `:1284` `--host "$BIND_HOST"` iff `STACK_PUBLIC_HOST` (`BIND_HOST=0.0.0.0`, set at `:66`); else `cockpit.py:534` default `127.0.0.1`
- `:1285-88` `( launch_detached "$STACK/cockpit.pid" -- python3 "$HERE/cockpit.py" … ) </dev/null > "$STACK/cockpit.log" 2>&1`
- **`:1295` logs `presenter cockpit serving on …` UNCONDITIONALLY** — no healthz probe, no exit-code check.

**`detach.sh:25-51`** `launch_detached PIDFILE -- CMD…`: Linux = `setsid "$@" & ; echo "$!" > "$pidfile"` (`:31-32`); macOS = python3 double-fork, intermediary writes the grandchild pid (`:46`). **The pidfile is written unconditionally, before/independent of a successful bind; no liveness check; it OVERWRITES any existing pidfile.**

**Bind + crash** (`cockpit.py`): imports are `argparse, html, json, sys, urllib.parse, http.server` — **no `socket`, no `os`, no `subprocess`**. `:538-542` manifest read **once** (try/except); `:551` seed-manifest slurped once; `:565` closure captures both; **`:567` `httpd = ThreadingHTTPServer((a.host, a.port), handler)` — OUTSIDE any try/except** → EADDRINUSE = unhandled `OSError: [Errno 98]` traceback → exit 1. `allow_reuse_address=1` is inherited from the stdlib base class and is **useless here** (SO_REUSEADDR does not permit binding over an **active** listener on Linux). **No SO_REUSEPORT, no port-in-use handling, no pre-bind reap.**

**Teardown** (`rosetta-demo:139-197`): `:143` `guard_n` (hard-refuses the dev project) → `:149` `ant-academy.sh N --stop` → **`:152-156` the cockpit reap, verbatim:**
```
  if [ -f "$stack/cockpit.pid" ]; then
    kill "$(cat "$stack/cockpit.pid")" 2>/dev/null || true
    rm -f "$stack/cockpit.pid"
    echo "==> demo-$n: stopped the presenter cockpit" >&2
  fi
```
**By pid only. Never by port. `kill` status DISCARDED. `rm -f` runs even when the kill failed. No `kill -0` re-check. No cmdline identity check (a RECYCLED pid gets killed blind). A missing/stale pidfile ⇒ a silent no-op that STILL prints "stopped the presenter cockpit".**
→ `:166-176` tailscale-serve reset (gated on `$stack/tailscale-serve.sh` + `tailscale` on PATH) → `:177-194` `docker compose -p demo-$n down` (`-v --remove-orphans` + image `rmi` under `--purge`) → `:195-196` `reg_del` + `ureg_release`.

**ALREADY-EXISTING, REUSABLE PRIMITIVE** — `rosetta-demo:56-57`:
```
ports_from_override() { [ -f "$1" ] || return 0; grep -oE '"[0-9]+:[0-9]+"' "$1" | tr -d '"' | cut -d: -f1 | sort -un | paste -sd, -; }
```
`gen_injected_override.py` emits quoted `- "13000:3000"` pairs, so this matches the injected override. **This is the right source for the compose half of the reap range. DO NOT use the registry** (see §5-6 / F3.2 — `/demo-up` never calls `set-ports`; the unified registry is `{}` on both boxes).

**`ant-academy.sh:67-73` is the BETTER existing pattern** (still pid-only, but liveness-checked — copy its shape):
```
  if [ -f "$PIDFILE" ] && kill -0 "$(cat "$PIDFILE")" 2>/dev/null; then
    kill "$(cat "$PIDFILE")" 2>/dev/null && log "stopped (pid …)."
    rm -f "$PIDFILE"
  else log "no running academy recorded for demo-$N (nothing to stop)."; fi
```

#### Every host-native (non-container) listener a demo owns
| listener | port | launcher | pidfile |
|---|---|---|---|
| presenter cockpit (`python3 http.server`) | **7700 + N·10000** | `up-injected.sh:1285-88` via `detach.sh` | `demo-stack/stacks/demo-N/cockpit.pid` |
| ant-academy (`next dev`) | **3077 + N·10000** | `ant-academy.sh:37` | `demo-stack/stacks/demo-N/ant-academy.pid` |

That is **all**. Everything else is a container. `tailscale serve` is a **node-level** listener (see below).

#### THE DEFINITIVE OFFSET PORT MAP for demo-N (OFFSET = N·10000)
**Compose-published** (base `docker-compose.yml` + `common.yml`, offset by `gen_injected_override.py:293-295`): postgresql **5432**, redis **6379**, graphql/cosmo **5050**, sentinel **8087**, backend **8081/8082/8083**, jobsimulation **8400/8401**, cms **8090/8091**, skillpath **8100/8101**, storage **8300/8301**, roadrunner **10400/10401**, gotenberg **3200**, studio-desk **9000/9100**, next-web-app **3000**. *(Not in the `graphql` profile ⇒ never started: customerio-sync 8080, messenger 8200/8201.)*
**Injected:** fake-fapi **5400** (`gen_injected_override.py:406`); **fake-bapi publishes NO host port** (PORT=443, network alias only, `:414-427`) ⇒ **not in the reap range**; directus **8055** (`stack-core/gen_override.py:152,166`) when local content is on.
**Host-native:** cockpit **7700**, ant-academy **3077**.
⇒ **demo-1 (off=10000):** `15432,16379,15050,18087,18081,18082,18083,18400,18401,18090,18091,18100,18101,18300,18301,20400,20401,13200,19000,19100,13000,15400,18055,17700,13077`.
**`18082` = the `0.0.0.0:18082` that aborted the earlier run** (backend, fronted by tailscale serve).

**RECOMMENDED REAP SET** = `ports_from_override "$STACK/docker-compose.injected.yml"` **∪** `{7700+off, 3077+off}`, with the literal table as the fallback when the override file is gone.

#### The THIRD leak class — `tailscale serve` (node-level, survives `compose down`)
`gen_tailscale_serve.py:38-45` fronts: backend **8082**, graphql **5050**, next-web **3000**, studio-desk **9000**, ant-academy **3077** — each `+OFFSET`. (FAPI 5400 **excluded** — self-TLS, `:47-48`.) Each is `tailscale serve --bg --https=<base+off> http://127.0.0.1:<base+off>` (`:68`) — a **real listener on the tailnet IP that persists past `compose down`** (`:24-29`). Reset at `rosetta-demo:166-176`, **gated on `$stack/tailscale-serve.sh` existing — which lives inside the rext clone, so an rext re-pin DISARMS it.** Manual unblock: `tailscale serve reset`.

#### LIVE EVIDENCE (billion, read-only, 2026-07-13)
```
$ ps -o pid,lstart,cmd -p 83214
83214  Sat Jul 11 18:01:55 2026  python3 /home/devops/panorama/stack-demo/rosetta-extensions/demo-stack/cockpit.py \
  --manifest …/stacks/demo-1/cockpit-manifest.json --fapi-host billion.taildc510.ts.net:15400 \
  --app-base https://billion.taildc510.ts.net:13000 --port 17700 --host 0.0.0.0 \
  --seed-manifest …/stacks/demo-1/seed-generation-manifest.yaml
$ ss -ltnp | grep 17700
LISTEN 0 5 0.0.0.0:17700 users:(("python3",pid=83214,fd=3))
$ find ~/panorama -name cockpit.pid
(nothing — only …/stacks/demo-1/ant-academy.pid)
$ docker ps -a      # EMPTY
```
The port has been held for **two days** and is **unreapable by pid**. This is the leak, live.

#### WHAT TO BUILD
**(a) NEW shared helper** — `demo-stack/reap.sh` (beside `detach.sh`), sourced by **both** `up-injected.sh` and `rosetta-demo`:
`reap_port PORT` — resolve the listener via `lsof -ti tcp:PORT -sTCP:LISTEN` (works on mac **and** linux) with an `ss -ltnpH "sport = :PORT"` fallback (linux, no lsof); TERM, short wait, then KILL; **verify with a re-probe**; log LOUD. No-op (exit 0) when the port is free. *(macOS has no `fuser -n tcp`; `lsof` is present on both.)*
**SAFETY:** refuse to kill any pid whose argv is not one of `{cockpit.py, next dev / ant-academy, a docker-proxy for -p demo-N}` — **never blind-kill a stranger's port.**
**(b)** `reap_stack_ports N` — the union set above. Call it from `rosetta-demo cmd_down` **AFTER `docker compose down`** (so compose releases its own ports first and only genuine orphans get killed), **non-fatally**.
**(c)** `rosetta-demo:152-156` — **stop lying**: `kill -0`-check first (the `ant-academy.sh:67-73` pattern), only `rm -f` the pidfile when the process is actually gone, and only print "stopped the presenter cockpit" when a kill actually happened.
**(d)** `up-injected.sh` — **PRE-BIND reap** of `7700+off` (defensibly, the whole set) **before** the `launch_detached` at `:1285`, so a re-`/demo-up` over a leaked stack self-heals. **AND a port preflight/reap for the compose range BEFORE `docker compose up` at `:944`** — the natural slot is between `precreate_linux_data_dirs "$STACK/data"` (`:938`) and `log "bringing up the full injected stack"` (`:939`). *(The existing tailscale pre-reset at `:986-1004` is AFTER compose-up and — per its own ADV-1 comment at `:995-1000` — cannot prevent the first bind conflict.)*
**(e)** `up-injected.sh` — **POST-launch** `GET http://127.0.0.1:$((7700+OFFSET))/healthz` probe (cockpit.py already serves `200 "ok"`); on failure log LOUD + tail `$STACK/cockpit.log`, **still non-fatal**. Replaces the unconditional "serving" line at `:1295`.
**(f)** `cockpit.py:567` — wrap the bind in `try/except OSError as e:` → `sys.stderr.write(f"cockpit: cannot bind {a.host}:{a.port}: {e}\n")` + `return 2` (a clean, greppable failure instead of a traceback). **Keep it stdlib-only; do NOT put the reap in cockpit.py** — the launcher owns killing a predecessor.

#### ⚠️ ORDERING HAZARD (interacts with scope item 6)
The pidfiles **and** `tailscale-serve.sh` live **inside the rext consumption clone** (`demo-stack/stacks/demo-N/`). **REAP + `tailscale serve reset` BEFORE re-cloning/re-pinning rext**, or every host-native listener and the serve config become permanently unreapable. Durable fix to consider: move `stacks/` **out** of the clone (or preserve it across re-pins).

---

### 5-2. SCOPE ITEM 2 — Un-swallow the demo-patch REFUSE reason

**Exact current lines:**
```
up-injected.sh:701   if "$HERE/../stack-injection/apply-app-authz-skip.sh" "$dst" >/dev/null 2>&1; then
up-injected.sh:717   if "$HERE/../stack-injection/apply-app-aireadiness-loadmembers.sh" "$dst" >/dev/null 2>&1; then
```
The appliers print the exact reason to **stderr** then exit 2 — `apply-app-authz-skip.sh:60-61`:
```
REFUSE — sha256(<target>)=<digest> != pre_sha256=<pin>
```
(`apply-app-aireadiness-loadmembers.sh:63-65` is the identical shape.) The else-branches (`:704`, `:720`) log only a generic `⚠ … failed/refused (non-fatal)` with **no reason**.

**PRESERVE THIS ASYMMETRY:** the **three next-web demopatch calls** (`:403`, `:408`, `:423`) do **NOT** swallow stderr — **only the two app HELPERS do.** (`apply-authn.sh:688` swallows the same way — out of scope, but note it.)

**Fix:**
```sh
if out=$("$HELPER" "$dst" 2>&1); then
  log "…applied"
else
  log "⚠ … refused (non-fatal) — rc=$?"
  printf '%s\n' "$out" | sed 's/^/    /' >&2
fi
```
Keep **NON-FATAL**. Also consider `2>>"$STACK/demopatch.log"` so the reason survives in an artifact.

**Add a static fence** in `demo-stack/tests/test_frontend_build.py`: assert the helper invocations do **not** contain a `2>&1` redirected to `/dev/null`.

**SAME SWALLOW CLASS — fix with the same discipline (see §5-7):** `autoverify.sh:138` `if "$HERE/verify.sh" >/dev/null 2>&1;`. Capture, don't discard.

---

### 5-3. SCOPE ITEM 3 — Re-pin the two `app` perf demo-patches + a LOUD patch-freshness preflight

#### The re-pin values (computed; ready to paste)

**`app-targetrole-authz-skip`** — **ONE pin works on BOTH boxes** (`roles.go` is byte-identical @ v1.334.1 and v1.337.0):
```
pre_sha256:  88f25f060bfedfb956803a4de4feee96549edd8fa88fbc029ae44a57e20bcc97
post_sha256: 4320955ddf60d574d3456a722ef6847697cc3848cad903554497fe761bc9b3d2   (marker OK)
```

**`app-aireadiness-snapshot-loadmembers`** — **NO single pin can satisfy both boxes:**
```
@ v1.334.1 (local stack-demo):  pre  = b32169682a28368b218e4133934d8b1c171393ce9030492e1f7d70d20ac4d29f
                                post = 2ae4906d7032383e3606d8d641628bb019eb02af9cf87a8a2611a44420802a14  (marker OK)
@ v1.337.0 (billion):           pre  = dc9e167eda1ad42e8c5eb1fe097b8602a037a4aa273c0f4c9ae767c5eb7a333b
                                post = 759590752af96cc515bdbf7b9a60a56a1975a5a1abf809ae1e8eff17abf97b94  (marker OK)
```

#### THIS IS THE BD-3 ANSWER, MECHANICALLY
The tag built is box-dependent — `up-injected.sh:671`:
```
tag=$(git -C "$src" for-each-ref --sort=-v:refname --format='%(refname:short)' --count=1 'refs/tags/v*')
```
*(Deliberately not `git tag --list | head -1` — M215: `head` SIGPIPEs git on app's ~337 v-tags → exit 141 under `set -o pipefail` → aborts the bring-up.)*
⇒ **A one-shot re-pin CANNOT close M217 for `ai_readiness`.** The whole-file sha gate over a file that churns for unrelated reasons **is the root defect**; the **anchor already survives** (1× at every tag tested).

**Therefore: KEEP the sha gate BUT make it self-maintaining.**

**(a) Freshness preflight (LOUD).** Early in `up-injected.sh`, **BEFORE the inject loop at `:655`** (i.e. before the clone/patch/build phase): for each `app` manifest — resolve `$tag` exactly as `:671` does, `git show $tag:<manifest.path>`, sha256 it, compare to `pre_sha256`. On mismatch print: the manifest **id**, the **tag**, the **expected/actual sha**, the **anchor-occurrence count**, **and the recomputed correct `pre`/`post` pair**, and **FAIL LOUD** (per the milestone's own decision: *"a perf patch that silently degrades a demo from 5 s to 120 s is worse than a patch that refuses"*). Because the anchor count is 1, the preflight can safely emit the exact two lines to paste.

**(b) A `--repin` verb.** Compute `pre` from `git show $tag:<path>`; assert anchor count == 1; apply; compute `post`; rewrite the two manifest lines in place. Flow becomes: preflight fails loud → `rext demopatch repin --all` → commit → tag.

**(c) DO NOT** ref-pin the source clones as the fix — `up-injected.sh:671,681` force-checks-out the **newest `v*` tag on every bring-up**, so the rot returns.

**(d) Keep the escapes bypassable.** `DEMO_NO_PATCH` / `DEMO_NO_AUTHZ_SKIP` / `DEMO_NO_AIREADINESS_LOADMEMBERS_BOUND` must skip the preflight too (a deliberate no-patch run must not be blocked by a loud gate).

#### Close the test gap that let the drift ship (G2)
Add **live-clone pin tests for BOTH `app` manifests**, computed at the `for-each-ref` tag — mirroring the existing `test_demopatch.py:456` (next-web studio) and `:541` (pubweb chain) and `test_ant_academy.py:370-371`. Today the `app` manifests get only manifest-internal self-consistency (`:506`), there is **no aireadiness manifest test at all**, and `:581` deliberately asserts the helper **refuses a synthetic file** — it never validates the pin against a real clone.

---

### 5-4. SCOPE ITEM 4 — Fix `jobsimulation` exits(1)

#### ⚠️ THE ROADMAP'S DIAGNOSIS IS WRONG. Discard it. Do NOT emit a `command:` key.

**The failure chain, every link code-cited:**
1. `stack-demo/platform/docker-compose.yml:171-172` — the jobsimulation block (L113-176) ends with:
   ```yaml
       volumes:
         - $HOME/.aws/credentials:/root/.aws/credentials:ro
   ```
   The **only** `aws/credentials` bind in the entire compose file (added 2024-12-16, `6daa67e` "fix: aws credentials"). No other service has it.
2. When the host path **does not exist, Docker auto-creates it as an empty DIRECTORY.** Live on billion: `ls -la ~/.aws/` → `drwxr-xr-x 2 root root 4096 Jul 11 18:01 credentials` (root-owned **directory**, created at bring-up).
3. The container has a **directory** at `/root/.aws/credentials`.
4. `jobsimulation/cmd/root.go:183` → `ai.NewAIManager(serverContext, …)` (again at `:198` for `aiResultManager`).
5. `internal/ai/ai.go:85-91`:
   ```go
   cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("eu-west-1"))
   if err != nil { return nil, fmt.Errorf("can't load AWS config: %w", err) }
   ```
6. `aws-sdk-go-v2/config v1.32.25` → `shared_config.go:793-805 loadIniFiles` → `internal/ini@v1.8.6/ini.go:24-44 OpenFile`: **`os.Open` on a directory SUCCEEDS**, so the error is *not* an `ini.UnableToReadFile` and is **not** skipped by the `continue` at `shared_config.go:802`; `io.ReadAll` then fails **EISDIR** → wrapped as `SharedConfigLoadError`. **Reproduced verbatim** with that exact module version:
   `failed to load shared config file, <path>, read all: read <path>: is a directory`
   **Control:** with the path simply **absent**, `LoadDefaultConfig` returns `err = nil` — no creds needed.
7. `cmd/root.go:195-197` → `return fmt.Errorf("can't init AI: %w", err)` **out of RunE**.
8. `cmd/root.go:59-62` — the root command sets neither `SilenceUsage` nor `SilenceErrors`, so cobra prints `Error: …` **followed by the full usage/help block** (this is the "prints CLI help" everyone saw), then `:402-408 Execute()` → `os.Exit(1)`.

**Falsifiable check for the implementer:** `docker logs demo-<N>-jobsimulation-1` on a fresh host must show, **immediately ABOVE the usage block**:
```
Error: can't init AI: can't load AWS config: failed to load shared config file, /root/.aws/credentials, read all: read /root/.aws/credentials: is a directory
```
If a *different* `can't init …` first line appears, the mount fix is still required but a second init failure exists. **Read the FIRST line, never the help.**

#### The CLI contract (so nobody re-derives the wrong fix)
- Image: `Dockerfile.dev` → `ENTRYPOINT ["./application"]`, **no CMD**. Compose: **no `command:` key**. Binary runs with **zero args**.
- `cmd/root.go:59-62`: `rootCmd = &cobra.Command{Use: "jobimulation", …, RunE: <the entire server>}` → **no args ⇒ server runs.**
- Optional subcommands: `aggregate` (`cmd/aggregate.go:56`), `clone-session` (`:42`), `test-command` (`cmd/test.go:45`), `validate` (`cmd/validate.go:41`). **There is no `serve` and no `run`.**
- ⚠️ **`command: serve` would make it WORSE** (cobra `legacyArgs` → `unknown command "serve"` → exit 1 for real).
- ⚠️ `stack-demo/jobsimulation/CLAUDE.md` (platform repo, READ-ONLY) documents `go run . serve` — **that command does not exist.** Do not trust it; **do not edit it.**

#### THE FIX — rext-only, zero platform edit, no demo-patch, no escalation
**Site:** `.agentspace/rosetta-extensions/stack-injection/gen_injected_override.py`, function **`build_lines()`** (defined L271), inside the `for name, svc in cfg.get("services", {}).items():` loop — immediately after the existing postgres branch at **L295-299**:

```python
        # postgres data dir
        if name == "postgresql":
            pgdir = os.path.abspath(os.path.join(data_root, "postgresql"))
            body.append("    volumes: !override")
            body.append(f'      - "{pgdir}:/bitnami/postgresql"')
        # ADD (M217/T6):
        # jobsimulation: DROP the base compose's `$HOME/.aws/credentials:/root/.aws/credentials:ro`
        # bind (docker-compose.yml:171-172 — the only aws bind in the file). On a host with no
        # ~/.aws/credentials FILE, Docker auto-creates an empty DIRECTORY at the source path; the
        # container then sees a DIRECTORY at /root/.aws/credentials and aws-sdk-go-v2's
        # config.LoadDefaultConfig() hard-errors ("read all: ... is a directory") inside
        # ai.NewAIManager (internal/ai/ai.go:85) -> cobra RunE error -> usage printed -> exit(1).
        # A demo needs no AWS creds (Bedrock/Chime/LiveKit-recording are not exercised), and with the
        # path simply ABSENT LoadDefaultConfig returns nil. Zero platform edits.
        if name == "jobsimulation":
            body.append("    volumes: !reset null")
```

**Why this is legal (verified, not assumed):** the generator's own docstring (L10) states it *"Uses Compose `!override`/`!reset` tags so inherited sequences/build are replaced, not merged"* — and it **already** emits `ports: !override` (`:293`), `volumes: !override` (`:296-299`), and **`build: !reset null` for jobsimulation itself** (`:302-306`, since `INJECTED` at `:17` contains `jobsimulation`). The tag is already proven against the compose binary on **both** hosts. *(Plain compose merge semantics would **append** the volume, never remove it — hence `!reset`/`!override`, not a bare `volumes:` key. `volumes: !override []` is an equivalent alternative; prefer `!reset null` for symmetry with the existing `build: !reset null`.)*

**Belt-and-braces (cheap):** in `up-injected.sh`, next to the patch-freshness preflight, warn **non-fatally** if `[ -d "$HOME/.aws/credentials" ]` — a docker-created directory is a smell worth surfacing on any host.

**Tests:** `stack-injection/tests/test_injection.py` — assert the emitted override contains a `jobsimulation:` service whose body includes `volumes: !reset null`; **and, negatively, that no service ever gets a `command:` key.**

**Mirror gap (G14, likely out of scope — flag it):** `stack-core/gen_override.py:117-128` rewrites `volumes` **only** for `postgresql` ⇒ `/dev-up` on a Linux box has the identically dead jobsimulation. Same 3-line fix.

#### Downstream while it is dead (measure the effect of the fix)
- **Autoverify FAILS deterministically.** `up-injected.sh:1328` puts `jobsimulation` in the verify scope; `services.sh:42` probes it on `8400+offset`. **A dead container ⇒ the exact "1 check(s) FAILED" symptom.** ⇒ **RE-MEASURE AUTOVERIFY AFTER THIS FIX BEFORE HUNTING THAT ITEM SEPARATELY.**
- **GraphQL federation:** jobsimulation is 1 of 4 subgraphs (`gen_injected_override.py:17`). Every jobsim-owned type errors at the Cosmo router.
- **Playthroughs:** `pt-aisim-chat-launch` (`playthroughs/manifest/ai-simulations.yaml:49`) cannot pass on a fresh host.
- **Events:** no session-completed events on the `jobsimulation` Redis stream (`root.go:281-287`) ⇒ Skillpath progression never sees completions.

---

### 5-5. SCOPE ITEM 5 — Prime the snapshot cache on `billion`

#### The mechanics (verified)
- **Surfaces (4):** `surfaces.go:34-39` — `reference-toy`, `taxonomy` (schema `public`), `directus` (schema `directus`, `CapturesStructure=true`), `sim-embeddings` (schema `cms`). **The bring-up replays exactly 3:** `dev-setdress.sh:311` `for s in taxonomy directus sim-embeddings`.
- **Replay is cache-first and provably never captures:** `main.go:297-414` opens exactly ONE connection — the **stack's** offset Postgres (`pg.DSNForOffset(baseDSN, n)`, `:322,:334`). No prod DSN, no `capture.Run`.
- **Cache layout:** root = `--store` > `$STACKSNAP_STORE` > `workspaceRootFrom(cwd)/.agentspace/snapshots` (`main.go:142-171`, `store.go:145-146`; the walk goes **UP** for the nearest `.agentspace`). Key = `<surface>/<schema_version>/` (`store.go:39`) + `manifest.json` + one `<schema>.<table>.copy` per table (+ `_structure.sql` for directus).
- **Key = schema digest:** `md5(string_agg(table||'.'||column||':'||data_type))` over `information_schema.columns` (`pg.go:228-236`). **Row surfaces narrow to their own tables; directus digests the WHOLE schema** (`capture.go:145-155`).
- **Replay verifies every payload's SHA-256 before any write** (`replay.go:99-102` → rc=1 `CORRUPT payload`) ⇒ **transport is self-verifying: a truncated rsync fails LOUD.**

#### Exit codes (exhaustive, `main.go:54-73`)
| rc | const | meaning | raised at | bring-up handling |
|---|---|---|---|---|
| 0 | exitOK | replayed | `main.go:413` | `replayed $s into $STACK` |
| 1 | exitError | connect / firewall / corrupt payload | `:224,244,337,367,405` | **non-fatal** warn (`dev-setdress.sh:336`) |
| 3 | exitUsage | bad args / unknown surface | `:99,190,310` | **non-fatal** |
| **4** | exitUnprovisioned | **TARGET stack schema absent/empty** (`pg.ErrEmptySchema`), probed **BEFORE the cache** | `:355-359` | **non-fatal** (`dev-setdress.sh:326-329`) |
| **5** | exitCacheMiss | no cached snapshot at the stack's digest (after the M21 auto-provision attempt) | `:395-399` | **non-fatal** (`:330-335`) |
*(There is no rc=2.)* The whole set-dress is itself non-fatal (`up-injected.sh:1084-1089`) ⇒ **a cold cache never FAILS a bring-up — it silently DEGRADES it.**

#### Why billion skipped (`~/panorama/coldrun2.log`)
- `:288` `cannot replay — cache miss: no snapshot for taxonomy/5afc0bccf1df7ef538b643321fc6362f.` → **rc=5**
- `:296` `probe stack schema: pg: schema "directus": schema has no columns (empty digest)` → **rc=4** — **NOT a cache problem.** The run was `--no-local-content` (`:115`, `:263`, `:346`) ⇒ `provision_directus_step` never ran.
- `:302` `cache miss: no snapshot for sim-embeddings/032c99ea47678187631c59c31b4ef059` → **rc=5**

#### capture-on-box is IMPOSSIBLE
`~/.pgpass` → absent. `~/.claude.json` → absent. No staging `pg_dump`. **Every option in `snapshot-cold-start.md:104-170` is dead on that box.**

#### THE DIGESTS MATCH — the local cache is a drop-in
| surface | digest billion asked for | present locally? |
|---|---|---|
| taxonomy | `5afc0bccf1df7ef538b643321fc6362f` | ✅ (330,261 rows / 10 tables / **1.4 GB**) |
| sim-embeddings | `032c99ea47678187631c59c31b4ef059` | ✅ (1,490 rows / 4 tables / **5.2 MB**) |
| directus | `ea2e187a16056d14749ad50dba31099b` | ✅ (11,986 rows / 14 tables + `_structure.sql` (425 stmts) / **26 MB**) — needed once `--local-content` is on |

#### THE PRIME RECIPE
```bash
# billion's store root is ~/panorama/.agentspace/snapshots — ALREADY EXISTS and is EMPTY.
# (Derived, not chosen: dev-setdress.sh:312 passes no --store, so main.go:142-171 walks up from
#  cwd; ~/panorama/stack-demo/rosetta-extensions has no .agentspace, so the walk lands on ~/panorama.)
SRC=/Users/kirality/Workspace/anthropos/rosetta/.agentspace/snapshots
rsync -avP --checksum "$SRC/taxonomy/5afc0bccf1df7ef538b643321fc6362f"       devops@billion:~/panorama/.agentspace/snapshots/taxonomy/
rsync -avP --checksum "$SRC/sim-embeddings/032c99ea47678187631c59c31b4ef059" devops@billion:~/panorama/.agentspace/snapshots/sim-embeddings/
rsync -avP --checksum "$SRC/directus/ea2e187a16056d14749ad50dba31099b"       devops@billion:~/panorama/.agentspace/snapshots/directus/
# ~1.45 GB total; billion has 129 GB free; /usr/bin/rsync present.
# SKIP the legacy taxonomy/c75ce94d… + sim-embeddings/10146f28… dirs (dead weight; hardlinked locally).

# verify on the box (Go is at /usr/local/go/bin — NOT on the default non-login PATH):
ssh devops@billion 'export PATH=$PATH:/usr/local/go/bin;
  cd ~/panorama/stack-demo/rosetta-extensions/stack-snapshot &&
  go run ./cmd/stacksnap status --store ~/panorama/.agentspace/snapshots'
# expect 3 lines: taxonomy schema=5afc0bcc… · directus schema=ea2e187a… · sim-embeddings schema=032c99ea…
```

**Safety argument to write into the doc:** manifests carry `public_only: true` + a `predicate` (org-null / directus-public-published / sim-embeddings-public-simulation); `payload` fields are **bare filenames** (no host paths, no host state); **grep for `password|secret|postgres://` across all 3 manifests and `_structure.sql` → 0 hits**; the bytes are the **same firewalled public-only data the tooling already replays into every demo**. Integrity is enforced by replay's per-payload SHA-256.

#### ⚠️ THE SECOND HALF M217 MUST NOT MISS
Priming alone leaves directus at **rc=4**. To get all three GREEN the billion run must **also be local-content ON** (do **not** export `DEMO_NO_LOCAL_CONTENT=1`; `up-injected.sh:116,152,1071` wire that flag into both the compose override and dev-setdress). Then the chain is:
`provision_directus_step` (`CREATE SCHEMA directus` → `node cli.js bootstrap` @ `directus/directus:11.6.1`, `dev-setdress.sh:173,184-215`) → `stacksnap replay --surface directus` probes the whole-schema digest → miss → `tryAutoProvision` (`autoprovision.go`) applies the cached `_structure.sql` (425 stmts) → re-probe converges to `ea2e187a…` → **cache HIT** → rows load → **exit 0** → `boot_directus_step` (`:319-322`).

**Two hard preconditions:**
1. **Auto-provision is gated on a VIRGIN directus schema** (`userCollectionCountSQL` — zero non-`directus_%` base tables; the DDL is **not idempotent**). A half-provisioned schema falls through to rc=5 and does **not** self-repair. **Start from a purged stack.**
2. **Keep `DEV_SETDRESS_DIRECTUS_IMAGE=directus/directus:11.6.1`** — the directus digest is whole-schema, so a different image drifts it off `ea2e187a…` with **no cure but a real recapture**.
*(RAM note: `up-injected.sh:150-154` budgets extra Docker-VM RAM when local content is on — probably why the operator disabled it. Check `preflight_vm_ram` on the box.)*

---

### 5-6. SCOPE ITEM 6 — Re-pin the drifted rext consumption clones

#### The SoT
`.agentspace/rext.tag` = **`v2.2`** — exactly 4 bytes `7632 2e32`, **no trailing newline**. Reader: `demo-stack/lib/rext_tag.sh:16-28`, fn `rext_tag [<repo_root>]`; strips `#` comments + CRLF; **always exits 0**; echoes `""` when absent. It is **sourced**, never executed. **`.agentspace/` is GITIGNORED ⇒ the pin is PER-BOX and does not travel to billion.**

#### Actual clone state (measured 2026-07-13)
| clone | HEAD | `describe` | dirty | has `v2.2`? |
|---|---|---|---|---|
| `.agentspace/rosetta-extensions` (AUTHORING) | `39e8013` | **`v2.2`** (exact) | clean | yes |
| `stack-demo/rosetta-extensions` (local consumption) | `b796857` | **`quick-change-m211`** | clean | **NO** |
| `stack-dev/rosetta-extensions` | — | — | — | **ABSENT (dir does not exist)** |
| `devops@billion:~/panorama/stack-demo/rosetta-extensions` | `41a28aa` | **`panorama-m214-3-g41a28aa`** (not a tag) | **DIRTY (2 files)** | **NO** |

#### DATA-LOSS VERDICT: **SAFE — ZERO CONTENT LOST** (proven, not assumed)
Remote dirty files: `demo-stack/migrate-demo.sh`, `demo-stack/up-injected.sh`. But:
- `git log origin/main..HEAD` → **EMPTY**. `git stash list` → empty. Untracked → none. Reflog → only `clone` + `checkout to 41a28aa`. **No local commits.**
- `git merge-base --is-ancestor 41a28aa v2.2` → **true**.
- `migrate-demo.sh` remote-WT sha256 = `c26449baba370bdbaffad1784ea9ee86bb6810ccae86e8406f22453943621789` == `git show v2.2:…` sha256. **BYTE-IDENTICAL.**
- `up-injected.sh`: remote-WT `d6e72e68…` (1315 ln) vs v2.2 `09982dbc…` (1335 ln). **`diff remote v2.2 | grep -c '^<'` == 0** ⇒ the remote is a **STRICT SUBSET**; v2.2 only ADDS the F12 tailscale-serve-reset block (~981-1000).
The in-place M215 edits were already round-tripped upstream (`38a4214`, `2952631`, `00ba6b6`). **v2.2 strictly supersedes the box.**

**RE-VERIFY IMMEDIATELY BEFORE THE DESTRUCTIVE STEP** (someone may have edited the box since):
```bash
ssh devops@billion 'D=~/panorama/stack-demo/rosetta-extensions; git -C $D log --oneline origin/main..HEAD; git -C $D stash list; git -C $D status --porcelain -uall'
ssh devops@billion 'cat ~/panorama/stack-demo/rosetta-extensions/demo-stack/up-injected.sh' > /tmp/r.sh
git -C .agentspace/rosetta-extensions show v2.2:demo-stack/up-injected.sh > /tmp/v.sh
diff /tmp/r.sh /tmp/v.sh | grep '^<'      # MUST be empty. If NOT empty -> STOP, salvage first.
```
Salvage path if it is ever non-empty: `git -C $D stash push -m m217-presalvage` (recoverable) — **never `reset --hard` blind.**

#### The re-pin procedure
**Local `stack-demo` (clean):**
```bash
cd /Users/kirality/Workspace/anthropos/rosetta
git -C stack-demo/rosetta-extensions fetch --tags origin        # MANDATORY: v2.2 is NOT in this clone
git -C stack-demo/rosetta-extensions checkout "$(cat .agentspace/rext.tag)"
git -C stack-demo/rosetta-extensions describe --tags --exact-match   # MUST print v2.2
```
**Remote `billion` (DIRTY — needs `-f`):**
```bash
ssh devops@billion 'D=~/panorama/stack-demo/rosetta-extensions; cd $D
  git fetch --tags origin                 # auth OK: credential.helper=store + url.https://github.com/.insteadOf git@github.com:
  git stash push -u -m m217-repin-safety  # belt+braces; proven-redundant but recoverable
  git checkout -f v2.2                    # -f REQUIRED: a plain checkout dies "Your local changes would be overwritten"
  git describe --tags --exact-match'      # MUST print v2.2
```
Also confirm the box's own `~/panorama/.agentspace/rext.tag` contains `v2.2` (it already warns against v2.2, so it does — re-confirm; it is gitignored and hand-maintained per-box).

**HAZARDS:**
1. **`git fetch --tags` is NOT optional.** Neither clone has `v2.2` locally (local's newest is `quick-change-m211`; remote's is `panorama-m214`). A bare `checkout v2.2` dies `pathspec 'v2.2' did not match`. **`tailscale-serve.md:129-139` OMITS the fetch — that is how the remote landed on a bare sha.**
2. **Detached HEAD is the CORRECT end state** — `ensure-clones.sh:66` keys on `describe --tags --exact-match`; a branch makes `_cur_ref` fall through to `main`/`HEAD` and trips the drift warning even when correct.
3. **⚠️ REAP FIRST (see §5-1(g)):** the pidfiles and `tailscale-serve.sh` live inside this clone. Re-pinning without reaping orphans every host-native listener **permanently**.

#### ⚠️ WHAT ELSE ITEM 6 MUST SHIP — else the drift RETURNS on the next run
1. **Fix the ACTIVE DRIFT INJECTOR:** `.claude/skills/stack-secrets/SKILL.md:75` →
   `git -C "$SECDIR" fetch --tags --quiet && git -C "$SECDIR" checkout --quiet "$(cat .agentspace/rext.tag)"`
   (it currently hardcodes `stage-door-m30`, a v1.6-era tag, against **the same clone `/demo-up` pins** — and `/demo-up` **invokes `/stack-secrets`**). **Re-pinning without this is a no-op within one run.**
2. **Promote the pin guard from WARN to FAIL** in `ensure-clones.sh:68-76`, with a `DEMO_ALLOW_UNPINNED_REXT=1` escape (the `DEMO_NO_*`/`DEMO_ALLOW_*` family convention: `DEMO_NO_HOST_PREFLIGHT`, `DEMO_NO_PATCH`, `DEMO_NO_UI`, `DEMO_NO_AUTHZ_SKIP`) preserving the deliberate-authoring-work case its `:61-63` comment protects. **A guard that warns for N runs while nobody reads the log is not a guard.**
3. **Doc fixes:** `rosetta_demo.md:22` `v1.10.1` → `v2.2` (better: delete the prose copy and point at the file — the prose copy **IS** the M49 drift class); `tailscale-serve.md:134` resolve `<panorama-tag>` → `v2.2` **and add the missing `git fetch --tags`**.
4. **Do not touch `stack-dev/rosetta-extensions`** — it does not exist on this box; any preflight must tolerate an absent dev clone.

#### DO NOT CONFLATE — the two clone kinds
| | **rext consumption clone** (item 6) | **platform build-scratch clone** (NOT this item) |
|---|---|---|
| path | `stack-demo/rosetta-extensions` | `$STACK/clones/{app,cms,jobsimulation,skillpath,sentinel}` |
| what | the TOOLING that runs the demo | the PLATFORM SOURCE the image is built from |
| pinned by | `.agentspace/rext.tag` (**warn-only**, `ensure-clones.sh:64-76`) | **nothing** — `up-injected.sh:671` recomputes the newest `v*` tag **every bring-up** |
| lifetime | persistent, hand-managed | gitignored throwaway; force-reset each run (`:681`) |

⇒ **BD-3's note is CORRECT: ref-pinning the rext clone does NOT stop demo-patch sha-rot.** Item 3 and item 6 are **independent fixes**. Fixing one does not fix the other.

---

### 5-7. CROSS-CUTTING — the auto-verify contract (context for items 2, 4, and M218)

**Non-fatality is real and load-bearing — DO NOT change it.** `autoverify.sh:32` `set -uo pipefail` (deliberately no `-e`), `:154` unconditional `exit 0`, plus `|| true` at both call sites (`up-injected.sh:1333`, `dev-stack:159`). **Add signal, never an abort** (#M18-D3).

**The swallow to fix (same class as item 2)** — `autoverify.sh:136-142`:
```
if "$HERE/verify.sh" >/dev/null 2>&1; then
  echo "  ✓ verify live: all liveness + readiness probes passed" >&2
else
  warn "verify live reported failing probe(s) — some service is not serving correctly."
fi
```
⇒ every `✗ <service> <detail>` line (`verify.sh:37,:68`) and the summary `✗ N probe(s) failed` (`:86`) is **thrown away**, and N failing probes collapse to **exactly 1** warning (`:141`). **Fix:** `vout=$("$HERE/verify.sh" 2>&1); rc=$?` and on `rc != 0` re-emit `grep -E '^\s*✗|fail'` into the ⚠ block; tee full output to `<stack>/autoverify.log`; propagate `verify.sh`'s real fail_count into the warning line.

**The probe set** (`verify.sh` exit 0 = all ok, **exit 1 = any probe failed**; rows in `services.sh:35-62`, container `<project>-<svc>-1`, port `base+offset`):
`postgresql 5432 (docker)` · `redis 6379 (docker)` · `sentinel 8087` · `backend 8082 /health` · `skillpath 8100` · **`jobsimulation 8400`** · `cms 8090` · `storage 8300` · `roadrunner 10400` · `graphql 5050 /health (http-200)` · `gotenberg 3200 /health (http-200)` · `next-web-app 3000` · `studio-desk 9000` · `directus 8055 /server/health (http-200, --local-content only)`.
Readiness (`lib/readiness.sh`) is gated **only on scope, NOT on liveness** (the `verify.sh:46` comment is wrong) ⇒ a dead service **double-counts**.

**The billion failure, verbatim (`coldrun2.log:348-355`):**
```
▶ autoverify demo-1 (offset 10000) — cheap-win asserts + scoped verify (non-fatal)
  ✓ backend /api/health 200 on :18082
  ✓ sentinel.casbin_rules = 1150 (authz policy loaded)
  ⚠ verify live reported failing probe(s) — some service is not serving correctly.

⚠⚠ autoverify demo-1: 1 check(s) FAILED — the stack is UP but may be non-functional.
```
Both cheap-wins PASSED ⇒ the failure is **inside `verify live`**, and the swallow means the log **cannot name it**. **By elimination + F13: the jobsimulation liveness probe at `:18400`.** ⇒ **Fix §5-4 first, then re-measure.** Confirm with:
```bash
STACK_PROJECT=demo-1 STACK_OFFSET=10000 ~/panorama/stack-demo/rosetta-extensions/stack-verify/live/verify.sh
```

**What verify does NOT check (G10) — candidate cheap-wins, in priority order:**
1. **A demo-patch REFUSED** — have the applier write `<stack>/demopatch-status.json` and have autoverify **warn on any `refused`**. *(This is literally how the drift rotted for releases.)*
2. **Snapshot replay skipped** — `SELECT count(*)` on the taxonomy tables in the stack's own Postgres, asserted `> 0` (mirrors the casbin assert exactly).
3. **Cockpit up** — `GET :7700+off/healthz`.
4. **fake-fapi up** (`5400+off`) — **a dead fake-fapi means NOBODY CAN LOG IN, and verify stays green.**

**For M218:** there is **no artifact today** (G11) — `/demo-up` exits **0** on a red verify and still prints `UP`. If M218 needs to gate its measurements on "the stack came up green", **M217 must emit one**: e.g. `<stack>/autoverify.json` `{"project":"demo-1","warnings":N,"failed":["liveness jobsimulation", …]}` written before the `exit 0`.

---

## 6. Open Items (require user decision)

### (a) DATA-LOSS RISK in re-pinning the remote rext clone — **RESOLVED: PROVEN SAFE, but re-verify**
**Verdict: ZERO content would be lost.** Full proof in §5-6: no local commits (`git log origin/main..HEAD` empty), no stashes, no untracked files; `41a28aa` is an **ancestor** of `v2.2`; `migrate-demo.sh` is **byte-identical** to the v2.2 blob; `up-injected.sh` is a **strict subset** (`diff | grep -c '^<'` == 0 — v2.2 only *adds* the F12 block). The M215 in-place edits were already round-tripped upstream.
**DECISION NEEDED:** confirm the implementer may run `git checkout -f v2.2` on billion **after** re-running the subset proof immediately beforehand (someone may edit the box in the interim). Recommended belt: `git stash push -u -m m217-repin-safety` first — proven redundant, but recoverable.
**⚠️ ORDERING CONSTRAINT (not optional):** the cockpit/academy **pidfiles** and `tailscale-serve.sh` live **inside this clone**. **REAP the ports and `tailscale serve reset` BEFORE the re-pin**, or the leaked listeners (incl. the live pid 83214 on `0.0.0.0:17700`) become **permanently unreapable**. Consider moving `stacks/` out of the clone as the durable fix — **that is a design decision, flag it.**

### (b) Can jobsimulation be fixed from the injected compose override, or does it need a demo-patch/escalation? — **RESOLVED: the override. No escalation.**
**Verdict: a 2-line addition to `gen_injected_override.py::build_lines` (§5-4).** It is **not** a compose-command problem (the roadmap's diagnosis is wrong — see F6.1); it is the `$HOME/.aws/credentials` bind auto-created as a **directory** (F6.2, proven live on billion + reproduced against `aws-sdk-go-v2/config v1.32.25`). The generator **already** emits `!override`/`!reset` tags for `ports`, `volumes` (postgres), and **`build: !reset null` for jobsimulation itself** — so `volumes: !reset null` is proven-legal on both hosts. **No demo-patch, no sha-pin, no platform edit, no escalation.**
**DECISIONS NEEDED:**
1. Confirm a demo **needs no AWS creds** (Bedrock / Chime / LiveKit-recording are not exercised in a demo). *If any demo surface DOES need Bedrock, the fix must instead ensure a real `~/.aws/credentials` FILE exists on the host — a different, host-provisioning fix.*
2. **Scope call on G14:** the identical bug exists on the **DEV** path (`stack-core/gen_override.py:117-128`). 3-line mirror fix. **In or out of M217?** (If out, it will resurface the first time anyone runs `/dev-up` on a Linux box.)
3. **Also correct `overview.md:53-55` + `decisions.md:45-49` + `roadmap.md:164`** — they instruct the implementer to "investigate a compose-command fix", which would **actively break the service**.

### (c) BD-3 — the sha-vs-anchor gate — **the mechanical evidence is now decisive; the decision is still yours**
**The audit's finding:** `up-injected.sh:671` builds the scratch clone at *"the newest `v*` tag **on this box**"*. `internal/roles/roles.go` is byte-identical @ v1.334.1 and v1.337.0 (one pin works). **`internal/workforce/ai_readiness.go` is NOT** (`b3216968…` @1.334.1 vs `dc9e167e…` @1.337.0). ⇒ **Any static whole-file sha pin for `ai_readiness` is WRONG on one of the two boxes the moment it is committed.** A one-shot re-pin **cannot** close M217 for that patch.
**Recommendation (agrees with the overview's, and now has the proof):** **KEEP the sha gate, add (b1) an auto-`--repin` verb + (b2) a LOUD freshness preflight** that computes `sha256(target @ the tag THIS box will build)` and fails loud with the exact re-pin lines to paste. The **anchor already survives** every tag tested (occurs exactly 1×), so the gate can be self-maintaining without weakening drift safety.
**DECISIONS NEEDED:**
1. **Confirm: keep sha + auto-repin + loud preflight** (vs. dropping to anchor-only single-occurrence matching, which survives bumps but loses the whole-file drift guarantee and the G7 post-condition).
2. **How loud is "fail loud"?** The milestone says *"a perf patch that silently degrades a demo from 5 s to 120 s is worse than a patch that refuses"*. Does the freshness preflight **ABORT the bring-up** (hard fail) or **warn LOUDLY and continue**? *(Recommendation: ABORT, with the three `DEMO_NO_*` escapes bypassing it — a green-demo gate that can be skipped by not reading a warning is the exact failure mode M217 exists to kill.)*
3. **Where does the repinned manifest live for a per-box-divergent file?** The auto-repin verb rewrites the checked-in manifest — which means **committing a pin that is correct on the box that ran it and wrong on the other.** Accept that (preflight catches it everywhere), or move to a per-tag pin table (`pins: {v1.334.1: {pre,post}, v1.337.0: {…}}`) — a **schema change to the manifest**, which is a bigger call. **Flagging: the current design cannot express "correct on both boxes".**

### (d) Additional decision surfaced by the audit — **the milestone's own premise for M218**
`overview.md:25-29` attributes the 1–2 min login to a stale cockpit serving "**dead `__clerk_identity` keys**". **That mechanism does not exist in the code** (F2.3 — `CockpitHero` carries no clerk id; the fake-FAPI resolves the stable stories.yaml key against a freshly bind-mounted roster). The **real** stale-cockpit hazards are baked scheme/host drift, preset drift, a stale download, and a **dead new cockpit falsely logged as serving**.
**DECISION NEEDED:** confirm M218 will **re-measure the login latency on a green stack** rather than inherit this (incorrect) mechanism as a hypothesis.

---

## 7. Backfill plan + proposed SECTION LIST for `progress.md`

### 7a. Backfill plan (doc deltas, mapped to their section)

| # | Doc | Delta | Lands in |
|---|---|---|---|
| B1 | **`corpus/ops/demo/demopatch-spec.md`** | **AUTHOR** (from §5-0): G1–G7, the manifest schema, the sha-gate semantics + exit codes, the anchor mechanism, the **three apply vehicles**, the chain rule, the app-patches-are-never-reverted note, the three opt-out env vars, the full patch inventory, the **BD-3 decision**, the **freshness preflight**, the **re-pin runbook** | **S1** |
| B2 | `frontend-tier.md:171-186` · `coverage-protocol.md:201-205` · `seeding-spec.md:373` · `stories-spec.md:471` · `ai-readiness.md:182-190` · `tailscale-serve.md:348-363,390` · `clerkenstein.md:150` | Back-link to the new spec; drop the stale `v2.89.0` narrative from `frontend-tier.md:183-186` | **S1** |
| B3 | `cockpit-spec.md:76-93` | Teardown is **port-authoritative**, not pid-authoritative; qualify D9 to "single-sourced **at launch**"; add the **bind-host** row (127.0.0.1 default / 0.0.0.0 under `--public-host`); "non-fatal **but verified**" | **S2** |
| B4 | `demo/README.md:30,55` · `.claude/skills/demo-down/SKILL.md` | Rewrite to the real **8-step** teardown + the new **offset-range port reap** | **S2** |
| B5 | `rosetta_demo.md:101-128` | The unified registry is **NOT** the port source for a `/demo-up` stack (`up-injected.sh` never calls `set-ports`); document the **two** registries (unified vs per-demo provenance) | **S2** |
| B6 | `idempotency.md` | New section: **"re-running the bring-up over a half-dead stack"** — now true because of the port-reap preflight | **S2** |
| B7 | `verification.md:161-174` | The failing-probe detail was **discarded**; "N check(s) FAILED" = **autoverify-level** checks, not probes; the verify-live aggregate is **one** check | **S3** |
| B8 | `verification.md:22-25,133-152` | Fold `next-web-app` + `studio-desk` into §"What runs"; drop the "frontends don't exist yet" caveat | **S3** |
| B9 | `corpus/services/jobsimulation.md` | **NEW "Startup contract"** subsection (cobra root RunE = the server; **no serve/run subcommand**; subcommands = aggregate/clone-session/test-command/validate; `ENTRYPOINT ./application`, no args; **any init error ⇒ `Error:` + usage + exit 1 — so "it printed help" means an INIT ERROR**). **Plus the `~/.aws/credentials` host requirement + the demo-override neutralization.** Drop Roadrunner from the RPC-dependency + Redis-consumer lists (code exec is **in-process to Judge0**, `internal/runner/`); note `ROADRUNNER_RPC_ADDR` is vestigial | **S4** |
| B10 | `snapshot-cold-start.md` | **NEW Option 3 — ship a warm cache to a remote box** (§5-5 rsync recipe + the byte-portability / no-secrets / self-verifying-SHA argument + "no prod credential needed on the target box"). Add the missing `--surface sim-embeddings` line to **Option 1** (`:130-131`) | **S5** |
| B11 | `snapshot-spec.md` §portable-format | One line: row surfaces digest **their own tables**; directus digests the **whole schema** (`capture.go:145-155`) — *this is why a cache can be transported at all* | **S5** |
| B12 | `rosetta_demo.md:20-24` | `v1.10.1` → **`v2.2`** — better: **delete the prose copy**, point at `.agentspace/rext.tag` (the prose copy IS the M49 drift class) | **S6** |
| B13 | `tailscale-serve.md:129-139` | Resolve `<panorama-tag>` → **`v2.2`** **and add the missing `git fetch --tags`** — *this omission is what produced the drifted remote clone* | **S6** |
| B14 | **`overview.md`** (this milestone) | **CORRECT 3 claims: (1)** jobsimulation root cause is the AWS bind mount, **not** a missing subcommand (`:53-55`); **(2)** the stale cockpit does **not** carry dead clerk-ids (`:25-29`); **(3)** **two**, not three, replays are cache misses — directus is rc=4 from `--no-local-content` (`:32-33`). Also `:85` line refs `669,679` → `671,681` | **S0 (pre-flight)** |

### 7b. Proposed SECTION LIST for `progress.md`

> **Order is load-bearing.** S1 first (everything in S3/S4 builds on the demopatch contract). **S2's reap MUST precede S7's re-pin** (the pidfiles live inside the clone being replaced). **S4 before S8** (jobsimulation is the prime suspect for the failing autoverify probe — fix it, then re-measure).

| § | Section | Deliverable | Gate / exit criterion |
|---|---|---|---|
| **S0** | **Clear RED — correct the milestone overview** *(no code)* | B14 — the 3 wrong claims in `overview.md` + the line refs | Overview no longer instructs a compose-command fix; the clerk-id mechanism is retracted; "two, not three" replay skips |
| **S1** | **`corpus/ops/demo/demopatch-spec.md`** — the blind area, FIRST | B1 + B2. The full §5-0 contract: G1–G7, manifest schema, sha gate + exit codes, anchor mechanism, **3 apply vehicles**, chain rule, the-app-patches-are-never-reverted, the 3 opt-outs, the patch inventory, BD-3, the freshness preflight, the re-pin runbook | The spec exists; all 7 back-links land; a reader can re-pin a patch without reading Go or Python |
| **S2** | **Port reap — bring-up + teardown** (scope 1) | `demo-stack/reap.sh` (`reap_port` + `reap_stack_ports`, argv-guarded); `rosetta-demo cmd_down` port reap **after** `compose down` + stop the pidfile lie (`:152-156`); `up-injected.sh` **pre-bind reap** before `:1285` **and a compose-range preflight before `:944`**; `cockpit.py:567` bind try/except → exit 2; post-launch `/healthz` probe replacing the unconditional `:1295` log. **+ B3, B4, B5, B6** | **Kill the live orphan on billion (pid 83214, `0.0.0.0:17700`).** A re-`/demo-up N` over a leaked stack self-heals. `demo-down` leaves **zero** listeners on the offset range. A crashed cockpit is **reported**, not claimed as serving |
| **S3** | **Un-swallow the REFUSE reason** (scope 2) | `up-injected.sh:701,717` capture stderr + re-emit in the else-branch (**keep NON-FATAL**); static fence in `test_frontend_build.py`. **Same discipline at `autoverify.sh:138`** (capture, don't discard; propagate `verify.sh`'s real fail_count). **+ B7, B8** | A refused patch prints its exact sha mismatch. A failing verify **names the probe** |
| **S4** | **jobsimulation exits(1)** (scope 4) | `gen_injected_override.py::build_lines` → `if name == "jobsimulation": body.append("    volumes: !reset null")`; the `[ -d "$HOME/.aws/credentials" ]` non-fatal warn; `test_injection.py` positive + negative asserts. **+ B9** | `docker logs demo-N-jobsimulation-1` on a **fresh** host shows a running server, not `Error: can't init AI` + usage. `pt-aisim-chat-launch` passes on billion |
| **S5** | **Re-pin the two `app` perf patches + the LOUD freshness preflight** (scope 3) | Re-pin `app-targetrole-authz-skip` (`pre=88f25f06… / post=4320955d…`); re-pin `app-aireadiness-snapshot-loadmembers` **at the tag the target box builds**; the freshness preflight before `up-injected.sh:655`; the `--repin` verb; **live-clone pin tests for BOTH app manifests** (close G2) | Both patches **APPLY** on billion. The preflight fails loud (with paste-ready lines) on the next `app` bump. **Members grid < 5 s; AI-readiness read < 5 s** |
| **S6** | **Prime the snapshot cache on billion** (scope 5) | rsync the 3 digest dirs to `~/panorama/.agentspace/snapshots/`; verify with `stacksnap status`; **run the demo WITH local content** (drop `DEMO_NO_LOCAL_CONTENT=1`) from a **purged** stack (G15) with the image pinned at `directus:11.6.1` (G16). **+ B10, B11** | All **three** replays exit **0**. Catalog is real; cms no longer reads content live from prod over the WAN |
| **S7** | **Re-pin the drifted rext consumption clones** (scope 6) — **AFTER S2's reap** | Re-verify the subset proof → `git fetch --tags` + `checkout [-f] v2.2` on **both** clones; **fix the drift injector** `stack-secrets/SKILL.md:75`; **promote `ensure-clones.sh:68-76` from WARN to FAIL** (+ `DEMO_ALLOW_UNPINNED_REXT=1`). **+ B12, B13** | Both clones `describe --tags --exact-match` → `v2.2`. A `/demo-up` → `/stack-secrets` cycle **no longer re-drifts** the clone. An unpinned clone **fails loud** |
| **S8** | **Green-stack proof + the machine-readable signal** | A **cold reset-to-seed `/demo-up 1 --public-host` on billion** at the new tag. Emit `<stack>/autoverify.json` (G11) so M218 can gate on green. Add the 4 missing cheap-wins (demopatch-applied, snapshot-replayed, cockpit-up, fake-fapi-up — G10) | **`/demo-up` comes up GREEN:** 0 verify warnings · 0 leaked ports · 3/3 replays exit 0 · 2/2 app patches applied · jobsimulation serving · cockpit serving a fresh manifest · both hero vantages reachable. **This is the M217 exit gate — and M218's precondition.** |

**Suggested commit/PR granularity:** S0+S1 (docs) → S2 (reap) → S3 (un-swallow) → S4 (jobsim) → S5 (re-pin+preflight) → S6 (cache) → S7 (rext re-pin) → S8 (proof + tag `v2.3`). Cut the rext tag **before** S7's remote re-pin so billion consumes the fixes.
---

## 8. Gate Result — RED → CLEARED (S0 correction pass, 2026-07-13)

**RED cause (a) — load-bearing STALE claims, 3 of them inside `overview.md`:** all three corrected in place.

| # | Claim | Correction |
|---|-------|------------|
| F6.1 | jobsimulation "prints CLI help — no `run`/`serve` subcommand → investigate a compose-command fix" | **The prescribed fix would have BROKEN the service** (`unknown command "serve"` → real exit 1). Root cause is the `$HOME/.aws/credentials` bind auto-created as a **directory** → `LoadDefaultConfig` EISDIR → `RunE` error → cobra prints usage → exit 1. Fix = `volumes: !reset null` in the generated override. **Independently confirmed: the demo carries ZERO AWS credentials** (0 hits for `AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY` in `platform/.env`), so the mount could only ever be the broken empty dir. |
| F2.3 | the stale cockpit served "dead `__clerk_identity` keys" | **That mechanism does not exist.** `CockpitHero` carries no clerk id; the seat key is the stable `stories.yaml` id resolved against a freshly bind-mounted roster. M218 must **re-measure on a green stack**, not inherit a phantom. |
| F5.1 | "**all three** snapshot replays SKIPPED (cold cache)" | **Two** are cache misses (rc=5). Directus is **rc=4** — the run was `--no-local-content`, so no `directus` schema existed. **Priming the cache alone will not fix it.** |

**RED cause (b) — the demopatch BLIND AREA:** **PROMOTED, not blocking.** `overview.md` carries an explicit
`Delivers → corpus/ops/demo/demopatch-spec.md`, and it is scheduled as **S1 — authored before any code lands**,
because S3 and S5 build directly on that contract.

**Verdict: YELLOW — proceed.** The six remaining stale corpus docs (B3–B13) are each backfilled by the section that
touches them; none is read as truth before its section runs. Tracked as `KB-1 … KB-6` in `decisions.md`.

### Live findings confirmed during the clear-pass
- **An orphaned cockpit was still alive on `billion`** (pid 83214, `0.0.0.0:17700`) — it **survived the
  `/demo-down`**, serving `demo-1`'s manifest against a database whose containers were removed. Killed manually.
  **This is S2's defect, caught live**: `cmd_down` reaps by PID only, discards `kill`'s status, `rm -f`s the pidfile
  regardless, and prints "stopped the presenter cockpit" even when it killed nothing.
- **The demo carries no AWS credentials at all** — confirming the S4 fix is unambiguously correct.
