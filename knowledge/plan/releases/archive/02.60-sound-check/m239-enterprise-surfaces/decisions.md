# M239 — Decisions

_(implementation choices with rationale accumulate here during build)_

- **User decision (2026-07-20): talk-to-data → FULL.** Real AWS Bedrock creds via `/stack-secrets` + secret-coverage DNA extension for `app` (reference `../hyper-studio/.env.example`), not just a flag. Recorded at design time; carried here for build traceability.

## §1 — talk-to-data build decisions

- **D1 — flag via env var, not a demopatch.** `NEXT_PUBLIC_DEMO_FLAGS_ALL=true` is baked into the web + hiring `.env.local` overlays (up-injected.sh). The frontend ALREADY reads it (`useTalkToDataAccess.ts:39` / `useCoursebuilderAccess.ts:40` — a demo escape hatch that forces the flag on while still requiring `isAdmin`), so this is demo-env wiring of an existing env var — cleaner than a demopatch (the M219 aireadiness patch was needed only because that gate reads no env var). Only those two admin-gated surfaces read it (verified by grep), so forcing it on unlocks nothing unintended. Folded into `next_web_patchset_fp` so a pre-M239 image rebuilds once (build-inlined value).

- **D2 — creds reach the container via `platform/.env`, bridged from `app/.env` (the key reconciliation).** The user's design said store in `.agentspace/secrets/app/`, DNA on `app`, "wire into the app compose service". Verified against actual code: the demo's **backend (`app`) container reads `env_file: .env` = the demo's `platform/.env`**, NOT `app/.env` (repo-local native-dev env, never mounted by the container). And the M217 override **drops the `~/.aws` mount** for a demo. So honoring the user's `app`-DNA framing correctly required a bridge: `bridge_bedrock_creds()` copies the Bedrock class `app/.env → platform/.env` right after provision (values-blind file→file, idempotent, non-fatal). This fully honors "wire into the app compose service" (the app/backend container now has the AWS env) while being correct for a containerized demo.

- **D3 — required-`standard`, deliberately NOT `critical` (R3).** The 2 real creds (`AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY`) are `required` (counted + flagged-if-missing) but `standard`, so their absence never fails the `Critical==100%` gate — a creds-less box still brings a demo up (Talk to Data merely inert, "no IMDS role"). Making them critical would break every creds-less demo (exactly R3). Region/session-token/`CLAUDE_CODE_USE_BEDROCK` are optional (config/STS-only/inert-for-app). NOT in `demoSatisfied` — operator-provided, not minted.

- **D4 — `CLAUDE_CODE_USE_BEDROCK` is inert for the app.** The audit + `bedrock.go` confirm askengine NEVER reads it (it always routes Bedrock unconditionally). It is a Claude Code CLI convention; provisioned for parity with the hyper-studio template, but does nothing for the app. Documented so a future reader doesn't chase it as load-bearing.

- **D5 — live verify: gate (b) Bedrock round-trip PROVEN.** The wired creds get a real answer from the exact app model/region — `aws bedrock-runtime converse --model-id eu.anthropic.claude-sonnet-4-6 --region eu-west-1` → `pong`, `stopReason end_turn`, 20 tokens (2026-07-21). This is the load-bearing, genuinely-uncertain part (the brief's "creds/region/permission" escape) — definitively answered YES. NOT a blocker. Full end-to-end UI click-path verification via a local demo bring-up follows.

## §2 — #4 library empty-first-load — VERDICT: no remaining defect (not a fix to force)

- **D6 — the `:5050` carry is already resolved in the current tooling.** apps/web's `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` is baked to the OFFSET origin by `up-injected.sh` (`build_frontend_next_web` line 744: `--build-arg NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=$SCHEME://$HOST:$((5050+OFFSET))/graphql`), which OVERRIDES the Dockerfile.dev ARG default `http://localhost:5050/graphql` (line 23) that `ENV` bakes. `build_frontend_hiring` (line 1023) does the same. The remaining `localhost:5050` references are the Dockerfile ARG *default* (overridden) + `packages/graphql/codegen.ts` (dev-authoring schema introspection, never in the runtime client). The M218/M220 image-reuse check rebuilds a stale-endpoint image. M237's re-triage already found the grid populates (7→29 cards, 0 errors). **No client-fetch race defect remains** — per the three-fate guard, this is recorded as a verified verdict, not a manufactured fix. Confirmed live on the demo bring-up.

## §3 — #1 hierarchical manager menu — presence-verify + coverage assertion

- **D7 — M237-resolved; confirmed by code + live.** M237 already RESOLVED #1 (the manager nav renders the grouped Organization structure on a fresh build). §3 is a presence-verify (live) + a coverage-sweep assertion. The coverage manifest notes manager descriptors need live calibration; the assertion is authored against the live manager render from this milestone's bring-up (avoids an uncalibrated false-RED).

## Live verification — ALL GREEN on demo-1 (localhost, cold reset-to-seed, 2026-07-21)

- **D8 — talk-to-data ANSWERS live, end-to-end (the milestone headline).** On demo-1 (17 containers up, 5 orgs / 191 users seeded), logged in as the manager admin hero `dan-manager` via the cockpit seat-switch, navigated to `/enterprise/talk-to-data`, asked "how many members does my organization have?" → **a real, data-grounded answer streamed back:** *"Your organization, Cervato Systems, has 51 members in total."* The backend `ask.stream` log shows the full agentic loop: `iter=0 stop_reason=tool_use` (the model wrote SQL) → `query_postgres` ran `SELECT COUNT(*) FROM memberships` → `iter=1 stop_reason=end_turn` (the answer), **~7 s**, a genuine Bedrock round-trip through the wired creds. Proven at BOTH layers: (1) raw creds → `converse` → `pong` (eu.anthropic.claude-sonnet-4-6, eu-west-1); (2) the running backend → the full UI answer. Playwright spec `talk-to-data-m239.spec.ts` GREEN (11.7 s).
  - **Verified live in the running demo:** backend container `demo-1-backend-1` carries all 4 Bedrock keys (`printenv` masked `<set>` — values-blind); the flag reference is baked into the served next-web bundle; and the flag gate held (the page loaded for the admin instead of `router.replace('/home')`).
  - **Test-authoring note:** the composer is a CONTROLLED textbox whose Send button stays disabled until React registers text — a Playwright `fill()` sets the DOM value without firing React's onChange, so Send stays disabled and Enter no-ops. The spec types real keystrokes (`pressSequentially`) + asserts Send enables + clicks it. (Two red runs traced to this before the green — a real gotcha for any future talk-to-data e2e.)
- **D9 — #4 library + #1 menu GREEN.** `enterprise-surfaces-m239.spec.ts`: `/library/skill-paths` populates on first load (real cards over the OFFSET GraphQL endpoint — confirms the `:5050` carry is resolved on the live image, `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=http://localhost:15050/graphql`); the manager nav renders the grouped Organization structure. GREEN (7.5 s). Both were verdicts, not fixes — the live run confirmed no defect remains.

## Live-verify finding — the demo-full-build disk exhaustion (a confusing-symptom class)

- **F1 — a full cold demo build can exhaust the Docker VM disk, and it surfaces as a CRYPTIC `redis exited (1)`, NOT a loud build failure.** The first demo-1 bring-up ran all 6 Go image builds + 3× ~3.7 GB frontend builds (next-web + studio-desk + hiring) — the frontend builds alone left **>16 GB of build cache** on top of ~13 GB of images. On a 59 GB Docker VM that pushed the disk to 100%, and the FIRST container that needed to write (redis, its `appendonlydir`) failed with `No space left on device` → `demo-1-redis-1 exited (1)` → the whole compose `up` aborted (every app/frontend container stuck `Created`, depending on a healthy redis). The symptom (`redis exited 1`) reads like a redis/port problem; the cause is disk. **Root-caused via `docker logs demo-1-redis-1` (the explicit `No space left on device` line), not guessed.** Recovery: `docker builder prune -af` reclaimed ~25 GB (the frontend build cache is pure overhead once the images exist), clean teardown + re-up reused the cached images (no rebuild) and completed. **Observation worth carrying:** the bring-up has a 12 GB VM-memory pre-flight but **no free-DISK headroom pre-flight for the full 3-frontend build**, and the failure it produces is mis-attributable. A candidate hardening (out of M239 scope; belongs to a demo-reliability milestone if pursued): pre-flight the build-phase disk headroom, or fail the BUILD loudly on ENOSPC rather than letting a downstream container be the messenger.

  - **F1 RESOLUTION (harden pass, LANDED — Fate-1, rext `053db23`).** The "pre-flight the build-phase disk
    headroom" candidate above was **imprecise**: a disk pre-flight ALREADY exists (`preflight_disk_headroom`,
    M49 #6) and runs BEFORE the frontend build. The genuine defect the harden pass found: it measured host `/`
    via `df -Pk /`, which on **Docker Desktop is a different, usually-huge filesystem that does NOT reflect the
    VM's own virtual disk** — the fs that actually ENOSPCs. So it read ~200 GB "free" on the host while the
    VM's disk filled, staying **GREEN through the exact failure**, which then surfaced as the cryptic `redis
    exited (1)`. **Landed fix:** probe the VM's internal disk via a throwaway `busybox df` (the container root
    == the VM overlay), fall back to host `/` only when Docker/df is unreachable, and **name the redis
    mis-attribution** in the warn (so a full VM reads as *disk*, not *redis*). Kept **non-fatal** (the
    thrice-stated pre-flight contract — never block a working bring-up on a soft heuristic). +4 unit tests via
    a busybox-df stub branch; the `DEMO_DISK_AVAIL_KB` seam still short-circuits. **Live-proof on this box: 25
    GiB VM-disk free vs 212 GiB host-`/` free** — the exact blind spot. `frontend-tier.md` corrected. The
    second candidate ("fail the BUILD loudly on ENOSPC") is NOT taken — it touches the build/compose error
    handling (higher risk, out of harden scope); the corrected pre-flight closes the misattribution at the
    front instead.

## Close review — decisions (M239 close, 2026-07-21)

The close code-quality + adversarial reviews of the M239 rext change set (`443a365^..053db23`) found **two real
defects in M239's own new code**, both landed Fate-1 (rext `cf89365`, main re-pushed + the 3 consumption tags
re-pinned to it). Values-blind held throughout; zero platform-repo edits.

- **D10 — [Fate-1, LANDED] the VM-disk pre-flight could ABORT the whole bring-up (the F1 fix had a regression).**
  `_vm_disk_avail_kb`'s `docker run --rm busybox df | awk` pipeline had no `|| true`. Under the script's
  `set -euo pipefail`, a **daemon-present-but-unreachable/wedged** run (or a box that can't pull busybox) makes
  the pipeline exit nonzero, and the caller's bare `avail_kb="$(_vm_disk_avail_kb)"` then trips errexit and
  **kills the bring-up before the host-`/` fallback ever runs** — defeating the probe's own non-fatal contract
  and re-introducing the ISSUE-7 errexit class the same file warns about. Confirmed REAL on bash 3.2 (`x=$(f)`
  under a bare `set -e` genuinely exits) and mutation-verified the fix's regression test is non-vacuous. Fix:
  `|| true` on the pipeline (empty-on-failure → caller falls back). Regression test invoked as a BARE command
  under `set -e` (an `if (...)` subshell is an "-e is ignored" context → would be vacuous) via a new
  `VM_PROBE_RC` stub seam.

- **D11 — [Fate-1, LANDED] `bridge_bedrock_creds` lacked a trailing-newline guard.** A base env whose last line
  had no trailing newline would concatenate the first bridged key onto it (`GH_PAT=…AWS_ACCESS_KEY_ID=…` on one
  line) — corrupting BOTH keys AND breaking the `^KEY=` idempotency match (re-append every re-up). Fix: a
  once-before-first-append leading-`\n` guard, mirroring provision's own `io.go` guard; values-blind (the last
  byte is tested, never emitted). +2 tests (concat-prevented + idempotent-on-newline-less-base).

### Adversarial review (Phase 2c) — scenarios considered

Recorded per the close contract (the scenario, not just the fix). Two became D10/D11 (fixed); three are
**accepted** with rationale:

- **AR-1 — credential ROTATION / `AWS_SESSION_TOKEN` expiry does not re-propagate on a re-up (copy-if-absent).**
  The bridge skips any key already in the base env, so a rotated value in `app/.env` never overwrites the stale
  line in `platform/.env`, and the log says "wired (idempotent)". **Accepted:** a demo is disposable +
  VPN-scoped, and the sanctioned hyper-studio template uses **permanent** IAM creds (no session token), so the
  expiry path is off the documented flow. Remediation if ever needed: delete the key from both `.env`s (or
  purge the stack-demo workspace) and re-up. Not a code change — a value-diff detector would fight values-blind
  for a case the intended creds don't hit.
- **AR-2 — the busybox probe can't run on the exact condition it targets (offline / VM-so-full-a-4MB-pull-ENOSPCs)
  → silent host-`/` fallback → GREEN through a full VM.** **Accepted by design:** you cannot measure the VM disk
  without running a container; falling back to a possibly-misleading host reading beats aborting or blocking a
  soft, non-fatal heuristic. The operator-visible signal is the `host /` vs `Docker VM disk` log label. (The
  ABORT half of this scenario WAS a real bug — fixed as D10.)
- **AR-3 — the two required cred genes are scored independently (no both-or-neither invariant), so a half-present
  pair reads ~50% coverage while 0% functional.** **Accepted:** a both-or-neither operator is a framework-wide
  scoring-engine change (affects every DNA class, higher risk), out of a docs+tooling fix milestone's altitude;
  the **Critical-gate neutrality** and **demo-overlay-doesn't-satisfy-AWS** properties ARE solid and tested
  (`bedrock_measure_test.go`). The half-present state is gate-neutral (nothing blocks); it only makes a coverage
  number optimistic — a soft, non-fatal signal.

### Scope re-fates (from the Phase 1b deferral audit — YELLOW)

- **D12 — DEF-M239-01: the second F1 candidate ("fail the BUILD loudly on ENOSPC") → Fate-3 → M244.** The F1
  RESOLUTION note above recorded it as "NOT taken" (a passive observation). Given a proper fate at close: it is a
  **distinct** build/compose error-handling change to `up-injected.sh` (higher-risk, outside this milestone's
  altitude), and the *actual* observed defect (measuring host `/` not the VM disk) already LANDED. Home = **M244
  prove-on-billion** (the terminal reliability closer that runs full cold builds on `billion`, where a
  build-phase ENOSPC recurs) as an OPTIONAL build-robustness hardening item; standing backlog is the acceptable
  alternative if M244's tooling surface shouldn't widen. Not blocking (single, first appearance).

- **D13 — the 9th demo-stack test failure (`test_reap.py::…::test_a_RACED_listener_exits_silently`) → Fate-3 →
  M244, with root cause + fix recipe.** The close full-suite sweep surfaced ONE failure beyond the documented
  standing-8: **root-caused to a test-isolation collision, NOT an M239 regression and NOT a reap.sh defect.**
  The RACED test calls `_reap_with_stubs(stubs)` with no port arg → the default `port=17700`; reap.sh
  (correctly, per its M217/M221 design) establishes real `/dev/tcp` occupancy BEFORE attribution, so on a box
  where a **live demo-1 cockpit is actually listening on 17700** (this box: PID 42277) it finds the port held,
  the stubbed `_listener_pids`/`_cmd_of` can't attribute it, and reap returns 1 ("cannot introspect / refusing
  to kill") — the correct behaviour, proven by its sibling `test_BUG_a_HELD_port_with_NO_attributable_pid…`. A
  clean box (no cockpit on 17700) passes it. Same host-state class as the standing-8 (0 product defect). **Fix
  recipe for M244:** give the RACED test (and any sibling that relies on the hardcoded 17700 default) a
  guaranteed-free `_free_port()` like the rest of the suite, so it is isolated from ambient infrastructure. Not
  blocking; not in the durable standing-8 (it only manifests with a live cockpit up).

- **STANDING-8 → M244 (confirm Fate-2, no new decision).** The close full sweep re-surfaced the identical
  documented 8 (`test_cockpit.py` ×6, `test_host_prereqs_m215.py` ×1, `test_purge.py` ×1) — 0 M239 regressions,
  already re-fated Fate-2 → M244 at the M238 close (M238-D5, one day old). Legitimate YELLOW repeat pattern,
  already fated; no new decision.
