---
milestone: M217
slug: clean-stage
version: v2.3 "cue to cue"
milestone_shape: section
status: planned
created: 2026-07-13
last_updated: 2026-07-13
complexity: medium
depends_on: none
delivers: a /demo-up that comes up GREEN — so that every number measured afterwards is real. Plus corpus/ops/demo/demopatch-spec.md (the sanctioned escape hatch this release depends on has no corpus doc)
issues: "the last real run on billion: cockpit CRASHED on a leaked port (stale manifest → dead clerk-ids); all 3 snapshot replays SKIPPED (cold cache → structural-only catalog); autoverify ended FAILING; jobsimulation exits(1); both app perf demo-patches REFUSE on sha-drift with the reason piped to /dev/null"
---

# M217 — Clean stage

## Goal
A `/demo-up` that comes up **green**. This milestone builds no feature — it removes the confounds that make every
downstream measurement a lie.

## Why this is a HARD BARRIER (not just a nice-to-have first step)

The user's reported defect (1–2 min cockpit login) was measured on a stack where:

- **The cockpit had CRASHED.** `OSError: [Errno 98] Address already in use` at `cockpit.py:567` — port `7700+off`
  was still held by a **leaked cockpit from the prior run**, and the new one died unhandled (the bind is outside any
  `try`). The bring-up **logged "presenter cockpit serving on …" anyway** (`up-injected.sh:1295` is unconditional —
  no healthz probe, no exit-code check), so the operator was driving a **stale predecessor**. The run *before* that
  aborted entirely on a leaked `0.0.0.0:18082`. **Two of the last three runs on `billion` were broken by leaked
  ports.**
  > **CORRECTED 2026-07-13 (KB-fidelity audit, F2.3).** The design draft said the stale cockpit served **"dead
  > `__clerk_identity` keys"**. **That mechanism does not exist.** `stack-seeding/seeders/cockpit.go:145-168` —
  > `CockpitHero` carries **no clerk id**; the seat `key` is the stable `stories.yaml` id (`maya-thriving`),
  > resolved against the **fresh roster bind-mounted every bring-up** (`gen_injected_override.py:406-419`). The
  > **real** stale-cockpit hazards are baked scheme/host drift, preset drift, a stale seed-manifest download, and
  > **a dead new cockpit falsely logged as serving**. M218 must **re-measure on a green stack**, not inherit a
  > phantom clerk-id hypothesis.
- **Both `app` perf demo-patches silently REFUSED** (whole-file sha-drift: pinned @ v1.295.0/v1.315.0, box runs
  **v1.337.0**), leaving the un-patched 76 s members grid and the 180 s AI-readiness read live — **and the refusal
  reason was piped to `/dev/null`.**
- **TWO of the three snapshot replays were cache misses** (`taxonomy`, `sim-embeddings` — **rc=5**) → a
  structural-only catalog.
  > **CORRECTED 2026-07-13 (F5.1).** The design draft said **all three** SKIPPED on a cold cache. The **directus**
  > one is **rc=4 (`exitUnprovisioned`)**, *not* a cache miss: the run was `--no-local-content`, so **no `directus`
  > schema existed at all** and cms read content **live from prod over the WAN**. **Priming the cache alone will
  > NOT fix directus** — the billion run must ALSO be **local-content ON**, from a **purged** stack (auto-provision
  > is gated on a virgin schema; the DDL is not idempotent).
- **Autoverify ended FAILING** ("1 check(s) FAILED — the stack is UP but may be non-functional") — **and the
  failing probe's identity was discarded** (`autoverify.sh:138` swallows `verify.sh`'s stderr, collapsing N failing
  probes into exactly 1 nameless warning). By elimination it is the **`jobsimulation` liveness probe** (see item 4).

**No number taken before this milestone lands is trustworthy.** M218 must not measure anything until M217 is green.

## Why section
The deliverables are enumerable up front — every one is a known, file:line-mapped defect with a known fix surface.
No exploratory path.

## Scope

### In
1. **Reap the leaked cockpit port** (`7700+offset`) on bring-up, and make `demo-down` reap the **whole offset port
   range** — not just the containers. (Root cause of 2 of the last 3 broken runs.)
2. **Un-swallow the demo-patch REFUSE reason.** `up-injected.sh:701,717` pipes the applier's stderr to `/dev/null`
   while the applier prints the exact sha mismatch (`apply-app-authz-skip.sh:60-61`). A refusal must be **visible**.
3. **Re-pin the two `app` perf demo-patches** + add a **patch-freshness preflight that FAILS LOUD**. The anchors
   were mechanically verified to still occur **exactly once** in the current files — so this is a **re-pin, not a
   re-authoring**. (**DEF-M215-01 / F5**, Fate-1.)
   > **A perf patch that silently degrades a demo from 5 s to 120 s is worse than a patch that refuses to apply.**
4. **Fix `jobsimulation` exits(1)** — **AI-Simulations is dead in every demo today**, and it is the near-certain
   cause of the failing autoverify probe. (**DEF-M215-04 / F13**, Fate-1.)
   > **⚠️ ROOT CAUSE CORRECTED 2026-07-13 (F6.1) — the design draft's diagnosis was WRONG and its proposed fix
   > would have BROKEN the service.** The draft said *"it prints CLI help → investigate a compose-command fix"*.
   > **There is no missing subcommand.** `jobsimulation/cmd/root.go:59-62` — the cobra **root command's `RunE` IS
   > the server**; the Dockerfile is `ENTRYPOINT ["./application"]` with no CMD; compose passes **no `command:`**.
   > It is *supposed* to run with zero args. **Adding `command: serve` would produce a real
   > `unknown command "serve"` → exit 1.** (`stack-demo/jobsimulation/CLAUDE.md` documents `go run . serve` — that
   > command does not exist. It is a platform repo: do not trust it, do not edit it.)
   >
   > **The REAL cause:** `docker-compose.yml:171-172` binds `$HOME/.aws/credentials:/root/.aws/credentials:ro` —
   > the only AWS bind in the file. When the host path **does not exist, Docker auto-creates it as an empty
   > DIRECTORY** (verified live on `billion`: `drwxr-xr-x root root … credentials`). The container then sees a
   > **directory** where a file belongs; `aws-sdk-go-v2`'s `config.LoadDefaultConfig` `os.Open`s it successfully,
   > then `io.ReadAll` fails **EISDIR** → `ai.NewAIManager` errors (`internal/ai/ai.go:85-91`) → the root `RunE`
   > returns an error → cobra (no `SilenceUsage`) prints `Error: …` **followed by the usage block** — *that* is the
   > "prints CLI help" everyone saw — → `exit(1)`. **With the path simply ABSENT, `LoadDefaultConfig` returns
   > `nil`.** Read the FIRST log line, never the help.
   >
   > **The fix — rext-only, zero platform edits, no demo-patch, no escalation:** drop the bind in the *generated
   > override* (`stack-injection/gen_injected_override.py::build_lines`): `if name == "jobsimulation": body.append("    volumes: !reset null")`.
   > Proven-legal — the generator **already** emits `ports: !override`, `volumes: !override` (postgres), and
   > **`build: !reset null` for jobsimulation itself**. **A demo carries NO AWS credentials at all** (verified:
   > zero `AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY` anywhere in `platform/.env`), so the mount could only ever be
   > the broken empty directory.
   >
   > **Mirror (Fate-1, same section):** the identical bug exists on the **dev** path
   > (`stack-core/gen_override.py:117-128` rewrites `volumes` only for `postgresql`) ⇒ `/dev-up` on a Linux box has
   > the same dead jobsimulation. Same 3-line fix.
5. **Prime the snapshot cache on `billion`** so replay stops skipping. (**DEF-M215-02 / F9**, Fate-1. Also feeds
   M218's **C-3** hypothesis — a cold federation tier is the leading both-hero latency suspect.)
   > **The cache is a DROP-IN — no prod credential is needed on the target box.** The three digests `billion` asked
   > for are all present in the local store: `taxonomy/5afc0bcc…` (1.4 GB), `sim-embeddings/032c99ea…` (5.2 MB),
   > `directus/ea2e187a…` (26 MB). Ship them by `rsync` to `~/panorama/.agentspace/snapshots/` (the store root the
   > box's own walk-up resolves to). Safe: the payloads are the **public-only, firewalled** data the tooling
   > already replays into every demo, they carry **no host paths and no secrets** (grepped), and **replay verifies
   > every payload's SHA-256 before any write** — so a truncated transfer fails loud rather than corrupting.
   > **Capture-on-box is impossible anyway** (no `~/.pgpass`, no staging dump — every option in
   > `snapshot-cold-start.md` is dead on that box), which is itself a doc gap this section fixes.
   >
   > **⚠️ THE SECOND HALF, easily missed:** priming alone leaves **directus at rc=4**. The billion run must ALSO be
   > **local-content ON** (do *not* set `DEMO_NO_LOCAL_CONTENT=1`), started from a **PURGED** stack — the directus
   > auto-provision is gated on a **virgin** schema and its DDL is **not idempotent**, so a half-provisioned schema
   > silently falls through to rc=5 and never self-repairs. Keep `directus:11.6.1` pinned (the digest is
   > whole-schema; another image drifts it off `ea2e187a…` with no cure but a real recapture).
6. **Re-pin the drifted rext consumption clones** — SoT `.agentspace/rext.tag` = `v2.2`; local `stack-demo` =
   `quick-change-m211` (5 tags stale); remote = `panorama-m214-3-g41a28aa` (not even on a tag; the box warns about
   itself every run).
   > **DATA-LOSS: PROVEN SAFE** (audit §5-6, re-verify immediately before the destructive step). The remote clone is
   > dirty in 2 files but has **no local commits** (`git log origin/main..HEAD` empty), no stashes, no untracked
   > files; `41a28aa` **is an ancestor of `v2.2`**; `migrate-demo.sh` is **byte-identical** to the v2.2 blob; and
   > `up-injected.sh` is a **strict subset** (`diff remote v2.2 | grep '^<'` is empty — v2.2 only *adds* the F12
   > block). The M215 in-place edits were already round-tripped upstream. Belt: `git stash push -u` first.
   >
   > **⚠️ ORDERING (not optional):** the cockpit/academy **pidfiles and `tailscale-serve.sh` live INSIDE this
   > clone**. **Reap the ports BEFORE the re-pin**, or every leaked host-native listener becomes permanently
   > unreapable.
   >
   > **⚠️ AND FIX THE DRIFT INJECTOR, or the re-pin is a no-op within one run:**
   > `.claude/skills/stack-secrets/SKILL.md:75` **hardcodes `stage-door-m30`** (a v1.6-era tag) and checks it out in
   > **the same clone `/demo-up` pins** — and `/demo-up` **invokes `/stack-secrets`**. Also: `git fetch --tags` is
   > **mandatory** (neither clone has `v2.2` locally), and its omission from `tailscale-serve.md:129-139` is
   > **exactly how the remote landed on a bare sha**. Promote the `ensure-clones.sh` pin guard from **WARN to FAIL**
   > (+ a `DEMO_ALLOW_UNPINNED_REXT=1` escape) — a guard that warns while nobody reads the log is not a guard.

### Out
- **Any latency fix.** That is M218 — and M218 must **measure before it guesses**. Do not scaffold a fix here.
- Any change to `/demo-up`'s defaults (M220) or the AI-readiness render path (M219).

## Delivers → knowledge/corpus

**`corpus/ops/demo/demopatch-spec.md`** — **BLIND AREA (promoted → this is SECTION 1, authored before any code).**
The demo-patch mechanism is the **sanctioned zero-platform-edit escape hatch this entire release depends on**, and it
has **no corpus doc**: the contract lives in a Python module docstring plus a rext-side README. The corpus's only
mentions are a 16-line blockquote in `frontend-tier.md` and a routing-table *cell* in `coverage-protocol.md`, and
**neither documents the manifest schema, the sha-gate semantics, the exit codes, the chain rule, or the two
non-`demopatch` apply vehicles.**

Must document (ground truth in `kb-fidelity-audit.md` §5-0): **G1–G7** (there is an unnamed **7th** guard — the
apply post-condition); the **10 mandatory manifest keys**; the sha gate (**the WHOLE FILE is hashed, not the anchor
region — that is the rot source**); the anchor mechanism (byte-exact substring, **exactly-once required**); the
**THREE apply vehicles** (the `app` patches target the build-scratch clone *outside* the workspace, so `demopatch`'s
own G1/G6 **correctly refuse** them — two shell helpers re-implement the guard ladder against the same canonical
manifest); the **CHAIN RULE** (`next-web-public-website-url`'s `pre_sha256` **IS** `next-web-studio-url`'s
`post_sha256` — it reads "DRIFTED" against a pristine file **by design; do not "fix" it**); why the `app` patches are
never reverted (the scratch clone is force-reset each run); the three `DEMO_NO_*` opt-outs; the patch inventory; the
**BD-3 decision**; the **freshness preflight**; and the **re-pin runbook**.

## Open questions

**BD-3 — the demo-patch gate.** **The audit made this decisive, and it is worse than the design draft assumed.**

`up-injected.sh:671` builds the scratch clone at *"the newest `v*` tag **on this box**"*. Therefore:
- `internal/roles/roles.go` is **byte-identical** @ v1.334.1 and v1.337.0 → **one pin works on both boxes.**
- `internal/workforce/ai_readiness.go` is **NOT** (`b3216968…` @ v1.334.1 vs `dc9e167e…` @ v1.337.0).

⇒ **Any static whole-file sha pin for `ai_readiness` is WRONG on one of the two boxes the moment it is committed.
The manifest schema literally cannot express "correct on both boxes." A one-shot re-pin CANNOT close this milestone.**
Meanwhile **the anchor survives every tag tested** (occurs exactly 1×) — the *semantic* target is stable; only the
whole-file proxy rots.

- **Recommendation (unchanged, now proven):** keep the sha gate for its drift guarantee + the G7 post-condition, and
  make it **self-maintaining** — a **`--repin` verb** + a **LOUD freshness preflight** that computes
  `sha256(target @ the tag THIS box will build)`, checks the anchor still occurs exactly once, and emits the
  **paste-ready** corrected pin lines.
- **STILL NEEDS A USER CALL:** (1) does the preflight **ABORT** the bring-up or **warn loudly and continue**?
  (2) for a per-box-divergent file, do we accept a committed pin that is right on one box and wrong on the other
  (the preflight catching it everywhere), or change the manifest schema to a **per-tag pin table**?
- **Note:** ref-pinning the rext clone (scope item 6) does **NOT** stop this rot — items 3 and 6 are **independent**.
  The platform build-scratch clone is force-checked-out at the newest `v*` tag on **every** bring-up
  (`up-injected.sh:671,681`).

## KB dependencies
- **`kb-fidelity-audit.md`** (this milestone) — **the ground truth. Build from §5, not from the corpus docs.**
- `corpus/ops/rosetta_demo.md` (the demo lifecycle) · `corpus/ops/verification.md` (the autoverify contract)
- `corpus/ops/snapshot-spec.md` + `corpus/ops/snapshot-cold-start.md` (the cache-prime path)
- `corpus/ops/demo/coverage-protocol.md` (the fix-surface routing table)
- rext: `demo-stack/patches/demopatch` (the guards, in the module docstring — the thing being documented)

> **KB-fidelity gate (2026-07-13): RED → cleared by this correction pass.** The audit found **14 load-bearing stale
> claims**, three of them **in this overview** — including a jobsimulation diagnosis whose prescribed fix would have
> **actively broken the service**. All three are corrected above. Six corpus docs remain stale; each is backfilled
> by the section that touches it (see `progress.md`).
