# M217 — Decisions

_Implementation choices with rationale. One entry per decision; cite the code/doc it binds._

---

## D1 — The demo-patch gate becomes SELF-HEALING: the anchor is the contract, the sha is a baseline

**Decided by the user, 2026-07-13 (BD-3).** **Binds:** `stack-injection/apply_patch.py`,
`corpus/ops/demo/demopatch-spec.md` § "The freshness gate".

**The forcing fact.** `pre_sha256` hashes the **whole file**, but the demo builds its scratch clone at *"the newest
`v*` tag **on this box**"*. So any unrelated edit anywhere in that file, in any `app` release, breaks the pin. And
it is **not fixable by re-pinning**:

> `internal/workforce/ai_readiness.go` is **not byte-identical** between app **v1.334.1** (local) and **v1.337.0**
> (`billion`).

**No single committed whole-file pin can be correct on both boxes. The manifest schema cannot express the truth.**
Meanwhile the **anchor** — the exact code being replaced — survives every tag tested, occurring exactly once.

**The decision.** Anchor present exactly once + sha drifted → **recompute for this box, report loudly, apply**.
Anchor 0× or 2+× → a **real semantic break** → **ABORT (exit 3)**. G7 (the apply post-condition) still holds; it is
verified against the *recomputed* post-sha, so a bad swap still cannot be written.

**Rejected — keep the hard sha gate + a `--repin` verb.** It would abort a bring-up on **every** app release, and —
because the boxes are on different tags — a pin committed from one box would abort the other. It was protecting
against "something else in the file changed", which for a **perf-only, read-path, data-identical** shortcut in a
**demo** is a proxy, not a real protection.

**Rejected — a per-tag pin table.** Explicit and safe, but a manifest schema change that *still rots* (every new app
tag needs a new entry).

**Validated live:** on `billion` the freshly re-pinned manifest drifted **immediately** (`b3216968… → dc9e167e…`)
and the gate self-healed and applied. **A static pin would have failed on the very first run after being committed.**

---

## D2 — A REFUSED patch stays NON-FATAL. A BROKEN ANCHOR is FATAL.

**Binds:** `demo-stack/up-injected.sh` (both `app` applier call sites).

The M18/M19 contract — *a failed patch warns loudly but never aborts a good bring-up* — is **correct and kept**.
**Silence was the bug, not non-fatality.**

The one exception: **exit 3 (anchor broken)** aborts. An anchor that no longer resolves means the app team
refactored the code out from under the patch; the demo would then silently ship a **76-second members grid** (or a
manager page that never loads). **A demo that silently ships the slow path is worse than one that refuses to
build** — and unlike sha drift, this is not something the tooling can reason its way through. Both `DEMO_NO_*`
escapes bypass it, so a deliberate no-patch run is never blocked.

---

## D3 — Reap by PORT, but NEVER kill a process we don't own

**Binds:** `demo-stack/reap.sh`.

Reaping by port is the only correct answer (the port is what actually blocks the next bind; the pidfile is written
*before* the bind succeeds and is *overwritten* by the next run). But a naive `kill $(lsof -t -i:PORT)` is a
**footgun**: on a shared box some unrelated service may legitimately own `17700`.

So `reap_port` requires an **identity regex** and kills only matching listeners. A foreign process is **reported
loudly and left alive** — a port conflict with someone else's server surfaces as a diagnosable message instead of
either a silent hang or us murdering their process. Fenced by a test that asserts the foreign listener **survives**.

---

## D4 — The rext pin guard: WARN → FAIL

**Binds:** `demo-stack/ensure-clones.sh`. **Escape:** `DEMO_ALLOW_UNPINNED_REXT=1`.

The guard warned for **six releases** while nobody read the log, and the clones drifted anyway — the local one 5
tags stale, the remote one not even *on* a tag (it warned about **itself** on every run). **A guard that only warns
is not a guard.** A stack running tooling that is not the one you think it is attributes its results to the wrong
code — precisely how the perf-patch rot went unnoticed.

The deliberate-authoring-work case the original non-fatality protected is **real**, so it keeps an escape hatch. It
just no longer gets to be the **silent default**.

> **The guard alone would not have been enough.** `/stack-secrets` hardcoded a **v1.6-era tag** and checked it out
> in **the same clone `/demo-up` pins** — and `/demo-up` **invokes `/stack-secrets`**. Every bring-up silently
> dragged the clone back. **Re-pinning without killing that injector is a no-op within a single run.**

---

## D5 — Fix `jobsimulation` in the GENERATED OVERRIDE, not with a demo-patch

**Binds:** `stack-injection/gen_injected_override.py`, `stack-core/gen_override.py`.

The compose file that starts `jobsimulation` lives in `stack-demo/platform` — a **platform repo**, which we may
never edit. But rext already generates an **injected override** layered on top, and it already emits
`!override`/`!reset` tags for `ports`, `volumes` (postgres), and **`build: !reset null` for jobsimulation itself**.
So `volumes: !reset null` is proven-legal on both hosts. **No demo-patch, no sha pin, no escalation.**

**Verified against the compose binary rather than assumed — and the assumption would have been wrong:**

| override | result |
|----------|--------|
| `volumes: []` | ❌ **the bind SURVIVES** — compose *merges* volume sequences |
| `volumes: !reset null` | ✅ removed |
| `volumes: !override` (no items) | ✅ removed — what the dev emitter's `to_yaml` produces |

My first dev-path fix relied on the bare empty list and was **wrong**; the comment explaining it was wrong too.
Both corrected, and the tests now prove the removal by running the emitters' **real output** through
`docker compose config` — because emitting the right-looking line is not the same as the bind actually being gone.

**Also fixed on the dev path (Fate-1, same class):** `/dev-up` on any Linux box had the identical dead service.

---

## D6 — Ship a warm snapshot cache to the remote box (no prod credential on the target)

**Binds:** `corpus/ops/snapshot-cold-start.md` § Option 3.

Every existing cold-start option requires a **prod DSN on the target host**. `billion` has none — no `~/.pgpass`, no
staging dump — **and should not have one**. A demo VM holding prod credentials is a worse posture than one that
cannot reach prod at all.

Shipping the cache needs **no prod credential on the target**, a strict *improvement* on Options 1 and 2. It is safe
for the same reasons the cache already is: the payloads are the **public-only, firewalled** data the tooling already
replays into every stack (`public_only: true` + a predicate; `AssertPublicOnly` hard-fails a capture on a single
customer-scoped row), they carry **no secrets and no host paths** (grepped), and **replay verifies every payload's
SHA-256 before any write** — so a truncated transfer **fails loud** rather than poisoning a stack.

---

## KB items (Phase 0b fidelity gate — verdict **RED → cleared**)

- **KB-1** — the milestone's own `overview.md` carried **three false claims**, all corrected in S0. The worst
  (jobsimulation) would have sent the implementation down a path that **breaks the service it was meant to fix**.
  *This is the strongest possible argument for the pre-flight gate: the errors were in the plan, not the code.*
- **KB-2** — `demopatch-spec.md` was a **blind area** (the contract lived in a Python docstring). Authored in S1,
  **before any code**, because S3/S5 build on it.
- **KB-3** — `cockpit-spec.md` claimed teardown reaps the cockpit. True of the *intent*, false of the *behaviour*.
- **KB-4** — `verification.md` documented a "N check(s) FAILED" contract that **could not name the failing probe**.
- **KB-5** — `rosetta_demo.md` carried a **prose copy** of the rext pin that was **six releases stale**. *A prose
  copy of a pin is the drift class it was meant to retire.* Deleted; the file is the SoT.
- **KB-6** — `tailscale-serve.md`'s fresh-VM runbook **omitted `git fetch --tags`**. That omission is **literally
  how `billion` landed on an untagged commit**.

---

## Pattern worth naming — the tooling was implicitly Docker-Desktop-shaped

Two independent bugs in this one milestone are **the same failure**:

| bug | mechanism |
|-----|-----------|
| `jobsimulation` exits(1) | Docker **auto-creates** the missing `$HOME/.aws/credentials` host path as an empty **directory**; the AWS SDK opens it, fails `EISDIR`, and takes the service down |
| Directus never provisions | **`host.docker.internal` does not resolve on Linux Docker Engine** — a Docker Desktop convenience. `CREATE SCHEMA` failed, the whole local-content provision was skipped, and the demo silently read content **live from prod over the WAN** |

Both are *"works on a Mac workstation, dead on a fresh Linux VM"*, and both were invisible because **nobody had run
a demo end-to-end on Linux until v2.2 put one on `billion`**. Expect more of this class. The mitigation is
structural: **the remote VM is now a first-class proving ground** (M221 makes it the release's acceptance gate).

---

## Exit gate — MET (billion, cold reset-to-seed, 2026-07-13)

```
✓ demo-patches: none refused
✓ taxonomy replayed: public.skills = 42790
✓ presenter cockpit answering on :17700
✓ clerkenstein fake-FAPI answering on :15400 (hero login is possible)
▶ autoverify demo-1: OK — verified-working.

set-dressed (content:local-content, snapshot:taxonomy=replayed directus=replayed
             directus=local-served sim-embeddings=replayed, stories seeded)
autoverify.json: {"project":"demo-1","offset":10000,"warnings":0,"green":true}
```

| Gate | Before M217 | After |
|------|-------------|-------|
| verify warnings | 1+ (**unnamed**) | **0 — green** |
| snapshot replays | 2 cache-miss + 1 rc=4 | **3/3 exit 0** |
| `app` perf patches | **0/2** (silently refused, 4 releases) | **2/2 applied** (one **self-healed**) |
| `jobsimulation` | exits(1) in **every** demo | **serving** |
| content plane | **live from prod over the WAN** | **local** |
| cockpit | logged "serving" unconditionally | **healthz-gated**, serving all 5 heroes |
| rext pin | warned; both clones drifted | **FATAL**; both at `cue-to-cue-m217` |

**This is the first time this box has ever come up green.** M218 may now measure — and the `autoverify.json` signal
is what it gates on, so it can never again measure a broken stack.
