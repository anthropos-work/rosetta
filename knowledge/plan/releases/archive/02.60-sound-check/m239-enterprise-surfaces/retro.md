# M239 — Retro

## Summary
M239 was the **second post-barrier fix** in v2.6 "sound check". The headline: **talk-to-data went FULL** — a real AWS
**Bedrock credential class** for `app` (5 genes: `AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY` **required·standard**,
`AWS_REGION`/`AWS_SESSION_TOKEN`/`CLAUDE_CODE_USE_BEDROCK` optional; **deliberately NOT critical**, R3 — a creds-less
demo still boots with Talk to Data merely inert) provisioned **values-blind** via `/stack-secrets`, then **bridged**
`app/.env → platform/.env` (the demo backend reads `env_file: .env`; the M217 override drops the `~/.aws` mount, so env
vars are the only vehicle to the container). **Proven end-to-end live** on demo-1, cold reset-to-seed: a manager asked
*"how many members?"* → a real Bedrock round-trip → *"Cervato Systems has 51 members"* (backend agentic loop
`tool_use`→`query_postgres`→`end_turn`, ~7 s). **#4 library** and **#1 menu** were **no-defect verdicts, not manufactured
fixes** — the `:5050` carry is already resolved (offset endpoint baked) and the grouped manager nav renders — both
recorded honestly per the three-fate guard and live-GREEN. Section milestone, closed-complete; **0 platform-repo edits**,
values-blind throughout.

## Incidents This Cycle
- **P2 — the F1 disk pre-flight fix shipped a real regression, caught + fixed at close (D10).** The harden Pass-3 F1 fix
  (`053db23`) correctly re-pointed the disk probe at the Docker VM disk via `docker run --rm busybox df`, but the
  pipeline had no `|| true`. Under the script's `set -euo pipefail`, a **daemon-present-but-unreachable/wedged** run (or
  an offline box that can't pull busybox) made the pipeline exit nonzero, and the caller's bare
  `avail_kb="$(_vm_disk_avail_kb)"` then tripped errexit and **aborted the whole bring-up before the host-`/` fallback
  ever ran** — defeating the probe's own non-fatal contract and re-introducing the ISSUE-7 errexit class the file warns
  about. Confirmed real on bash 3.2, fixed (`|| true`) with a **non-vacuous, mutation-verified** regression test (the
  test invokes the pre-flight as a BARE command under `set -e`, because an `if (...)` subshell is an "-e is ignored"
  context that would make it pass even unfixed). rext `cf89365`.
- **P3 — the bridge lacked a trailing-newline guard (D11).** `bridge_bedrock_creds` appended the first Bedrock key with
  no guarantee the base env ended in a newline; a newline-less base would concatenate `GH_PAT=…AWS_ACCESS_KEY_ID=…` onto
  one line — corrupting BOTH keys AND breaking the `^KEY=` idempotency match (re-append every re-up). Fixed with a
  once-before-first-append leading-`\n` guard (mirroring provision's own `io.go` guard), values-blind. rext `cf89365`.
- **A 9th demo-stack test failure surfaced (host-state, not a regression).** The close full sweep reproduced the
  documented 8 standing failures **plus** `test_reap…test_a_RACED_listener_exits_silently` — **root-caused to a
  test-isolation collision** (the test hardcodes port 17700 via the `_reap_with_stubs` default; reap.sh correctly
  establishes real `/dev/tcp` occupancy, so a live demo-1 cockpit on 17700 on this box makes it refuse). reap.sh is
  correct; a clean box passes it. → Fate-3 → M244 with a fix recipe (D13).
- **No product regressions.** The 9 full-sweep failures are all host-state debt (0 M239 regressions); touched suites
  green (Python 106/106, Go secretdna PASS); flake gate 5/5.

## What Went Well
- **The load-bearing uncertainty was proven, not assumed.** The milestone's genuinely-uncertain part — "do the creds /
  region / IAM permission actually work?" — was answered YES at TWO layers: a raw `converse` → `pong`
  (`eu.anthropic.claude-sonnet-4-6`, `eu-west-1`), then the running backend → a real data-grounded UI answer. The
  headline is a live fact, not a unit assertion.
- **Two verdicts, honestly recorded instead of manufactured.** #4 library and #1 menu were investigated and found to be
  already-resolved; per the three-fate guard they were recorded as verdicts (with the exact `up-injected.sh:744/1023`
  evidence and a live GREEN), NOT dressed up as fixes.
- **The close did real work.** The code-quality + adversarial reviews found two genuine defects in the milestone's own
  new code — one of them a bring-up-aborting regression the harden pass had just introduced — and both were fixed with
  mutation-verified regression tests. The reviews were not a rubber stamp.
- **Values-blind held end-to-end.** A real cloud credential was provisioned, bridged, live-round-tripped, and even
  stripped-and-restored (a harden mutation) without any value ever reaching a log, a test, a commit, or reasoning — the
  `SENTINEL` fence + counts-only logging proved it.
- **Zero platform edits held** — the whole feature is env-wiring + a secret class + a bridge + docs.

## What Didn't
- **A harden fix introduced a regression the harden pass didn't catch.** The F1 disk-probe fix omitted the `|| true` the
  same file's own GH_PAT extraction documents as the ISSUE-7 remedy, and the harden test only exercised the empty-exit-0
  path (never the nonzero-exit path), so the abort was invisible. The lesson: a new `docker run` in a `set -euo pipefail`
  script needs its errexit path tested, not just its happy path — the close added exactly that stub arm.
- **The full demo-stack suite is host-sensitive and wasn't run during harden.** The harden flake gate ran a scoped
  3-suite set; the full sweep (required at close) is what surfaced both the 9th reap collision and gave the errexit
  regression its context. Running the full suite earlier would have surfaced these sooner.

## Carried Forward
- **DEF-M239-01 → M244** (D12): the second F1 candidate "fail the BUILD loudly on ENOSPC" (build/compose error-handling,
  higher-risk) — optional build-robustness hardening for the terminal reliability closer.
- **D13 reap `test_a_RACED…` 17700 test-isolation collision → M244**: fix recipe recorded (give it a `_free_port()` like
  its siblings). Not in the durable standing-8 (only manifests with a live cockpit up).
- **Standing-8 demo-stack test debt → M244** (Fate-2, M238-D5, confirmed): identical set, 0 M239 regressions.
- **The 2 live-proof e2e specs → M244** live billion sweep (they require a live demo; executed on demo-1 at close).

## Metrics Delta
- **Tests:** Python touched 106/106 (bridge 12 [+2 close], frontend-build 94 [+1 close]); full demo-stack 794 passed / 9
  host-state fails (0 M239 regressions). Go secretdna PASS (70 funcs; +2 Bedrock measure). 2 live-proof e2e specs GREEN
  on demo-1. **Flake: 0** (5/5 sequential).
- **rext commits:** build `443a365`/`a18fac3`/`bd7e8db` · harden `0a59673`/`21444bb`/`053db23` · close `cf89365`. All 3
  consumption tags re-pinned to `cf89365`.
- **Platform-repo edits:** 0. **Supply chain:** 0 net-new deps (a Bedrock *secret class*, not a dep).
