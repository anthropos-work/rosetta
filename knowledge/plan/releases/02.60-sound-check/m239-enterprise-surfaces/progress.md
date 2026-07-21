# M239 — Progress

Section milestone. Checklist stub from the roadmap In-list.

## Sections

- [x] **talk-to-data (a) flag enablement** — `NEXT_PUBLIC_DEMO_FLAGS_ALL=true` baked into web+hiring `.env.local` (env var the platform already reads; not a patch) + folded into `next_web_patchset_fp`. rext `443a365` / tag `sound-check-m239-enterprise-surfaces`. Tested (`test_bedrock_bridge_m239.py`).
- [x] **talk-to-data (b) real AWS Bedrock creds** — Bedrock cred class in the secret-DNA (2 req·standard + 3 optional; NOT critical per R3) + values-blind `bridge_bedrock_creds()` (app/.env→platform/.env) + provisioned via the assembled source. Tested (`secret_dna_json_test.go` + bridge tests). **Live Bedrock round-trip PROVEN** (`converse` → `pong`, `eu.anthropic.claude-sonnet-4-6`, eu-west-1). Full-UI live confirm on the demo bring-up.
- [x] **#4 library empty-first-load** — VERDICT: no remaining defect. The `:5050` carry is already resolved (`up-injected.sh:744/1023` bake the offset endpoint for both web+hiring, overriding the Dockerfile ARG default). No client-fetch race remains (M237: grid populates 7→29). Recorded verdict, not a manufactured fix. Live-confirm on the bring-up.
- [x] **#1 hierarchical manager menu** — M237-resolved (grouped Organization nav for managers). Manager hero = org admin (`users.go:141`). Live presence-verify + coverage-sweep assertion (calibrated against the live manager render).
- [x] **Delivers** — `corpus/ops/secrets-spec.md` (the Bedrock cred class subsection + updated 61-gene map) + `corpus/ops/safety.md` §2.10 (secrets-posture shift: demo `app` holds real Bedrock creds; blast-radius + operator-scope caveat).

_All sections implemented + unit-tested + committed (rext) + documented (rosetta), and **LIVE-VERIFIED GREEN on demo-1** (localhost, cold reset-to-seed, 2026-07-21):_
- _§1 talk-to-data: **ANSWERS live** — manager asked member count → "Cervato Systems has 51 members" (backend agentic loop tool_use→SQL→end_turn, ~7 s Bedrock round-trip); backend container holds the creds; flag gate held. `talk-to-data-m239.spec.ts` GREEN._
- _§2 #4 library: populates first-load (offset endpoint baked). §3 #1 menu: grouped Organization nav renders. `enterprise-surfaces-m239.spec.ts` GREEN._
- _Live-proof specs at rext tag `sound-check-m239-live-proof`. Infra finding (demo disk-exhaustion → cryptic redis-exit + recovery) recorded in decisions.md F1._

## M239: Hardening

### Pass 1 — 2026-07-21 — mutation-verify the live-proof specs (on the warm demo-1)
Proved the 3 talk-to-data break-modes actually turn the specs RED (a check that can't fail is worse than none):
- **flag/gate OFF** — logged in as the non-admin employee seat `maya-thriving` (the same `!isEnabled` client
  redirect a flag-off admin hits): the page bounced to `/home` and the spec went RED. **Surfaced a real
  weakness:** the old flag-gate `toHaveURL(/talk-to-data/)` PASSED on the bounce — at `waitUntil:'commit'` it
  matched the transient landing URL BEFORE the client `router.replace('/home')` fired, so the assertion the
  comment credited as "the flag proof" was near-vacuous. **Fixed inline** (rext `0a59673`): race the
  composer-visible vs bounced-to-`/home`, then assert not-`/home` — re-verified RED at the flag gate + GREEN on
  the happy path.
- **creds MISSING** — physically stripped the AWS Bedrock keys from the demo backend's `platform/.env`
  (values-blind, by key-name) + recreated `demo-1-backend-1`: the Bedrock answer never streamed → RED ("chat
  did not grow"). Restored + recreated → GREEN. (Backend masked-`printenv` confirmed `<absent>`→`<set>`; no
  secret value ever read/echoed; no stray `.env` left.)
- **stubbed/empty answer** — characterized the answer matcher deterministically: it rejects echo / empty /
  refusal / vague / error-surface strings (all → RED); its one blind spot is a HALLUCINATED `<n> members`
  count without a `query_postgres` call — accepted + noted in the spec, since the real protection is the
  no-answer→RED path, not the digit match.
- **vacuity** — `enterprise-surfaces` #1 (manager nav) went RED for an employee (non-vacuous); #4 (library) is
  a shared surface (GREEN for an employee — not a mutation; endpoint-broken would need a rebuild).

### Pass 2 — 2026-07-21 — Bedrock DNA measure-layer + bridge coverage (values-blind)
- Added `stack-secrets/secretdna/bedrock_measure_test.go` (rext `21444bb`): `secret_dna_json_test.go` pinned the
  class STRUCTURE; these prove the EFFECT through `Measure`/`MeasureForStack` on the REAL committed
  `secret-dna.json` — a missing Bedrock cred is **caught** (named short in the app rollup + the per-gene score
  flips) yet **never moves the Critical gate** (R3: `Critical(withAWS)==Critical(withoutAWS)`, a credless demo
  still boots), and the **demo overlay does NOT satisfy** the operator-provided creds (an absent AWS cred still
  reports missing on a demo, with the Clerkenstein-minted-passes contrast on the same empty source).
- Confirmed the bridge suite (10 tests) incl. the `a18fac3` idempotent-truth regression
  (`test_already_present_logs_wired_not_inert`) + the flag baked/folded tests all green.
- **Priority-3 note:** M239 introduced **no demopatch** (per D1 the flag is `NEXT_PUBLIC_DEMO_FLAGS_ALL`
  env-var wiring, folded into `next_web_patchset_fp`) — there is no apply/revert surface to strand; covered by
  `TestTalkToDataFlag`.

### Pass 3 — 2026-07-21 — F1 disk pre-flight LANDED (Fate-1)
- **F1's premise ("the demo has no build-phase disk pre-flight") was imprecise** — it HAS one (M49 #6), called
  before the frontend build. The real defect: it measured host `/` via `df -Pk /`, which on Docker Desktop is a
  **different, usually-huge filesystem that does NOT reflect the VM's own virtual disk** (the fs a cold build
  ENOSPCs). So it read ~200 GB "free" on the host while the VM disk filled — staying GREEN through the exact
  failure, then surfacing as the cryptic `redis exited (1)`.
- **Fix** (rext `053db23`): probe the VM's internal disk via a throwaway `busybox df`, fall back to host `/`
  only when Docker/df is unreachable, and **name the redis mis-attribution** in the warn. Still non-fatal.
  `DEMO_DISK_AVAIL_KB` seam still short-circuits. **+4 tests** (VM-measured / VM-OK / host-fallback / seam-wins)
  via a busybox-df branch in the docker stub. **Live-proof: 25 GiB VM-disk free vs 212 GiB host-`/` free on the
  same box** — the exact blind spot the old proxy had. Doc corrected in `frontend-tier.md`.

**Bugs fixed inline:** (1) the talk-to-data flag-gate assertion was near-vacuous (transient-URL timing) —
strengthened + mutation-re-verified (rext `0a59673`). (2) the disk pre-flight measured the wrong filesystem
(host `/` ≠ Docker VM disk) — corrected to the VM disk + named the mis-attribution (rext `053db23`).

**Flakes stabilized:** none observed — flake gate GREEN 3/3 (go bedrock-measure + py disk-preflight[7] +
bedrock-bridge[10], sequential); both live-proof specs GREEN in hardened state.

### Stop condition
3 passes. All four harden priorities addressed (mutation-verify · DNA/bridge measure-layer · demopatch check
[N/A] · F1 disk pre-flight LANDED); the Step-2b scan surfaced nothing new worth adding; flake gate clean.
harden commits live in **rosetta-extensions** (`0a59673`, `21444bb`, `053db23`) — close-milestone should re-pin
the consumed tags to the hardened HEAD (the M237/M238 precedent).
