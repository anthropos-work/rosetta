# M3 — progress (section checklist)

**Milestone:** M3 — Disposable multi-instance demo stacks · **Shape:** section · **Status:** done (2026-06-03)

## Build log

**2026-06-03 — S2 (the override/isolation engine) built + LIVE-PROVEN.** Tooling in the new gitignored
`anthropos-demo/rosetta-demo/` repo (M3-D4; commit `946c5ba`): `lib/gen_override.py` (offset ports +
repoint Postgres bind, via Compose `!override` so sequences are *replaced* not appended — the append would
re-bind the base port and collide with the dev stack) + the `rosetta-demo` lifecycle CLI (`up`/`down`/
`status`/`gen`, every op `-p demo-N`-scoped, `down` hard-refuses the dev project) + `registry.json`.
**Live proof on this 16 GB box, alongside the running 12-container dev stack:** `demo-1` (postgresql+redis)
came up on offset ports **5532/6479** with its own data dir → two independent live Postgres instances side
by side; `status` listed it; `down --purge` cleanly removed it — and the **dev `anthropos` stack stayed
fully intact** (12 containers, postgres healthy) the whole time. This satisfies the M3-D5 acceptance
("one demo stack alongside the dev stack, untouched").

**2026-06-03 — M3 sections built + hardened; S3 corrected (see below).** Tooling in `anthropos-demo/rosetta-demo/`
(commits `946c5ba` S2 · `cda2db3` S1 · `31bdcd8` S3 · `b626020` harden) + rosetta `/demo-*` skills + the
ops guide. 12 tooling unit tests green; shellcheck + python compile clean. **Genuinely proven:** S1
(clone-at-tag), S2 (the override/isolation engine — demo-1 live alongside the dev stack, M3-D5), the
publishable-key mint, S4 skills, S5 guide. **Corrected (S3):** the Clerkenstein injection is **wiring scaffold
only, NOT verified on live services** — a direct attempt to run the `app` exposed two hard blockers (app
needs the full graphql profile; no patched-colony module exists for the authn replace). Reframed below +
routed to M3-CF1.

## S1 — layout + per-demo clone-at-release-tag (M3-D1 + M3-D3) ✅
- [x] `anthropos-demo/rosetta-demo/stacks/demo-N/` workspace layout
- [x] per-demo clone step — `clone_repos.py` + `rosetta-demo clone`; each repo at its latest release tag (caller-overridable; bare + `v`-semver; main if untagged) — resolution proven on all 14 repos + real clones
- [x] the stack registry/ledger (`registry.json`: live N, ports, profile/services, **resolved ref per repo**)

## S2 — compose override + port-offset + per-stack project/env/data ✅ (live-proven)
- [x] generated `docker-compose.demo.yml` override (`!override` remaps ports; repoints the Postgres bind)
- [x] port-offset scheme (`demo-N → base + N·10000`) — _max-N bound documented as a tuning knob in the guide_
- [x] `.env.demo-N` template + generation (project name, offset, + Clerkenstein endpoint vars via S3)
- [x] per-stack Postgres data dir isolation

## S3 — Clerkenstein injection: authn + frontend BUILT + PROVEN; backend/webhook emitted; running-app deployment = M3-CF1
> **Two corrections, in order (2026-06-03).** First I over-claimed (checkmarks = "wiring emitted" ≠ "works").
> Then, pushed to actually do it, I found my "resource-gated" reason was **wrong**: the *whole dev stack* uses
> **~0.9 GB** (measured), not 10-12 — RAM was never the blocker. And the "patched colony doesn't exist" was
> "I hadn't built it." So I **built it and proved it**.
- [x] `authn` — **BUILT + PROVEN.** A disarmed colony clerk provider (`rosetta-demo/inject/colony-authn-disarmed/clerk.go`: same package/type/`NewProvider` signature, universal-key verify) vendored via `apply-authn.sh` (clone colony@pinned → swap clerk pkg → `replace => ./vendor-colony`). **Proven against colony v0.34.3 (app's pinned version):** the disarmed provider accepts a Clerkenstein token + extracts identity+org; rejects garbage/expired. App code unchanged. (`inject_proof_test.go`.)
- [x] `clerk-frontend` minted publishable key → env — **PROVEN** (byte-identical to clerkenstein's gated `MintPublishableKey`).
- [x] **running-app deployment — PROVEN end-to-end (M3-CF1 RESOLVED, 2026-06-04).** Built a demo `app` image with the vendored disarmed colony (`apply-authn.sh` + a Dockerfile fix), ran it live, and hit a protected route (`/api/workforce/members`) with three tokens: **none → 400, garbage → 401 (rejected), Clerkenstein → 403 (ACCEPTED — past authn, denied at authz)**. The 403-not-401 is the proof: a real running Anthropos `app` accepts a Clerkenstein universal-key token at its live HTTP auth middleware. Recipe + result: `rosetta-demo/inject/DEPLOYMENT-PROOF.md`.
- [~] `clerk-backend` api.clerk.com → fake-BAPI — `app` `extra_hosts` `!override` snippet emitted; the cert/redirect is the one remaining recipe to verify live (the authn path is fully proven).
- [~] `clerk-webhook` svix-signed POST — injector invocation emitted; run when the webhook flow is exercised.
- **M3-CF1 RESOLVED.** The headline ("a demo comes up Clerk-free, accepting Clerkenstein tokens") is now **demonstrated on a live app**, not just scaffolded. RAM was never a blocker (the whole dev stack uses ~0.9 GB — measured).

## S4 — lifecycle skills + teardown (M3-D2 manual only) ✅
- [x] `/demo-up [N]` skill (wraps `rosetta-demo clone → inject → up`, resource-aware)
- [x] `/demo-down [N]` skill (wraps `rosetta-demo down N --purge`; `-p`-scoped, dev stack untouched — proven)
- [x] `/demo-status` skill (wraps `rosetta-demo status`; registry + per-demo `ps` + resolved refs)

## S5 — the ops guide + the acceptance demo ✅
- [x] `corpus/ops/rosetta_demo.md` (collision problem + additive fix + `!override` + port-offset + clone-at-tag + Clerkenstein recipes + safety + resource budget + proven-vs-gated split); cross-linked from `corpus/ops/README.md`
- [x] **acceptance (M3-D5):** demo-1 (postgres+redis) ran isolated alongside the dev stack on offset ports with its own data; up→status→down; **dev stack untouched throughout**. (Two-concurrent-full-stack acceptance is resource-gated → bigger box.)

## Migrate step (2026-06-04) — sentinel healthy + /api/health 200; authorized 200s = M4 seed
`/demo-up` now runs `migrate-demo.sh`: creates the schemas (sentinel/cms/jobsimulation/skiller/skillpath +
extensions) + the pgvector/pg_trgm/pgcrypto extensions, atlas-migrates the 5 services against the demo DB,
restarts sentinel+backend. **Result:** sentinel stops crash-looping (it needed its `sentinel` schema for
casbin) — **0 restarts, healthy**; `/api/health → 200`; 6/6 schemas migrated. **Still 403 on *authorized*
endpoints** (e.g. /api/workforce/members) — that needs the **M4 seed** (casbin policies + the demo user/org
matching the Clerkenstein demo identity), not the migrate step. Found an M4 nuance: `init_policy.sql`
seeds `casbin_rules` (plural) but the gorm adapter auto-creates `casbin_rule` (singular).

## M3: Hardening

### Pass 1 — 2026-06-04 (the extended-work surface)
Test-deepening on the **post-close extended work** (the full injected stack + deployment/injection
surface), which shipped with thin or zero coverage. Driven as a 6-target × (deepen + adversarial-strengthen)
workflow; full-suite + flake gate + commits done in the main thread. **No production behaviour changed; the
deployment alignment gate held 100%/100% (7/7 genes) throughout.**

**Test counts (funcs/methods), before → after:**
- `clerkenstein/deploy/colony-authn` (clerk.go, the disarmed provider): **2 → 34** (100% stmt cov). Edge/error-class grid, exp boundary (strict `>`), a drift-equivalence battery vs `clerkenstein/shared`, header-tamper, base64-alphabet pin, `GetUserByID` via colony's real Manager fan-out, **2 fuzzers** (zero crashers).
- `clerkenstein/alignment/cmd/deployrun`: **2 → 31** (94.9% cov; remaining lines unreachable-by-design — the real provider never rejects a valid universal-key token). Both bad-sig flip arms, exact wire-key casing pin, unknown-variant divergence, empty-DNA path.
- `clerkenstein/cmd/fake-fapi`: **0 → 10**; `cmd/fake-bapi`: **0 → 12** (newServer 100%; only the `ListenAndServe` shell uncovered). httptest smokes incl. the **mint→backend-authn round-trip**, fresh-seed isolation, method-scoped 405.
- `rosetta-demo/tests/test_tooling.py`: **13 → 55** (added `TestGenInjectedOverride`: real-YAML-tree parse of the injected override — a mutation test proved the prior substring checks missed an alias mis-indent — + fuzz + the `resolved()` docker boundary stubbed; `gen_injected_override.py` 88% → **98%**).
- `rosetta-demo/tests/test_inject_scripts.py`: **0 → 23** (new). `apply-authn.sh` fully tested offline with a stubbed `git`; `up-injected.sh`/`migrate-demo.sh` get shellcheck + structural-wiring regression (re-arming couplings shellcheck can't see).

**Bugs fixed inline (harden-surfaced, with regression tests):**
- `inject/apply-authn.sh` (`5ab7b51`): the colony-version `grep` runs under `set -e`+`pipefail`, so a no-match aborted the script **before** the explicit guard — bare `exit 1`, no diagnostic. Added `|| true` to fall through to the actionable guard.
- `migrate-demo.sh` (`5ab7b51`): SC2015 `A && log ok || log warn` (warn could misfire if the ok-log failed) → `if/then/else`.

**Production refactors (testability only, no behaviour change):** extracted `newServer()` from `main()` in both fake servers; extracted a pure `build_lines()` from `gen_injected_override.py`'s `main()`.

**Finding routed to M4 — the demo identity is `user_clerkenstein`, not the runner fixture.** The harden exposed **two divergent demo identities** in clerkenstein: the real browser-login/demo seed `clerkfrontend.DefaultDemoUser()` = **`user_clerkenstein` / `demo@anthropos.test` / `org_clerkenstein` / admin**, vs. the alignment **runner fixtures** (`deployrun`/`expressrun`) = `user_2clerkenstein` / `demo@anthropos.work` / `org_clerkenstein`. They agree on org but diverge on user-sub + email. The gates are self-consistent (the fixture is just the contract's test vector — harmless to them), but **the M4 seed must seed `user_clerkenstein` / `demo@anthropos.test`** (what the browser flow actually produces), NOT the `user_2clerkenstein` fixture some earlier notes propagated. (Earlier S3 notes/`DEPLOYMENT-PROOF.md` cite the fixture identity + a non-existent `org_demo`; corrected in the corpus sweep.) **M4 action:** seed the real `DefaultDemoUser` identity; optionally reconcile the runner fixtures to it (would require re-capturing the deploy/express goldens — out of scope here).

**Hygiene swept (commits `6150198` clerkenstein · `5ab3818` rosetta-demo):** untracked a stray `fake-bapi` Mach-O build artifact + a `.coverage` file; added `/fake-fapi /fake-bapi /mintpk` to clerkenstein `.gitignore` and `.coverage`/`__pycache__/` to rosetta-demo.

**Knowledge backfill:** clerkenstein `knowledge/coverage-index.md` refreshed; the demo identity + deployment surface folded into the rosetta corpus (see the corpus sweep). 

**Verification:** clerkenstein **all 13 packages green under `-race`** (gofmt + vet clean); rosetta-demo **78 tests green** (shellcheck clean); deploy gate **100%/100%**; **flake gate 3/3 clean**.

### Stop condition
Stopped after one pass: the highest-value targets reached saturation (clerk.go 100%, gen_injected_override 98%, both fake servers at their testable ceiling), the adversarial pass found only depth gaps (now filled) not new bugs, and the two genuine production bugs were fixed + pinned. The shell orchestrators (`up-injected.sh`/`migrate-demo.sh`) are honestly documented as I/O-bound-uncoverable-offline with a three-fold compensating strategy (shellcheck + structural-wiring regression + the live `DEPLOYMENT-PROOF.md`).

## M3-extended: Close Review (2026-06-04)

`/developer-kit:close-milestone` over the extended work (full injected stack + deployment surface + harden +
corpus + rename). M3 was already `done` + on the release branch + pushed, so this close = the accountability
layer (review + verify + ledger + retro), **no branch merge / no archive flip** (M3 archives at close-release).

**Phase 1b deferral audit:** YELLOW (not blocking) — 7 single deferrals, 0 repeat/chronic/aged-out; Fate-3
routings annotated into M4 (login identity + casbin gotcha) and M5 (the two interactive-demo recipes). Report:
[audit-deferrals/deferral-audit-2026-06-04-m3-extended-close.md](audit-deferrals/deferral-audit-2026-06-04-m3-extended-close.md).

**Review found findings across 3 dimensions** (workflow: scope/decisions · code-quality · docs/tests/adversarial). Fixed:
- **[must-fix] ×100 port-offset collision** — the base `rosetta-demo up` path defaulted `OFFSET=100`, so demo-1's
  storage `8300+100=8400` collided with the dev stack's jobsimulation `8400` (a real bug the base path shipped
  with; `up-injected.sh` already used `N·10000`). Fixed the CLI default → **10000** (collision-free, matches
  `up-injected.sh`, identical ports); documented the `max-N ≈ 5` trade-off + offset-tuning guidance.
- **[must-fix] registry.json concurrent corruption** — the read-modify-write had no locking, so two concurrent
  `up`/`clone` could lose each other's entries. Added a **portable `fcntl` lock** (not `flock(1)` — absent on
  macOS) around all three registry writes; proven against 20 concurrent writers (no lost writes).
- **[should-fix] docs consistency** — swept every stale `N·100`/`default 100` → `N·10000` (ops guide, README,
  overview, progress, decisions); added the partial-bring-up recovery note + the `gen_override` offset-convention
  comment + an `injection.md` demo-identity callout.
- **[should-fix] stale state.md** — refreshed headline numbers (78 + 218), the M3 entry, the "release branch
  pushed" note. Decisions M3-D6 (rename) / M3-D7 (close routings) recorded; the "Open" items resolved.
- **Accepted as-is (nice-to-have, documented in the review, not bugs):** registry schema versioning, a `parse_pk`
  fuzz, a `clone_repos` warning-log, an `apply-authn` error-message hint, the `REUSE_DEV`/dead-`else` comments.

**Phase 9 — Completeness Ledger (section):**
- **Done (Fate 1):** all original S1–S5 + the extended full-injected-stack + migrate + the deployment/injection
  alignment surface + the harden (+125 funcs) + the corpus update + the rename + the two close must-fixes.
- **Annotated (Fate 3):** clerk-backend cert/redirect + browser-login → M5; clerk-webhook live POST + the
  login-identity (`user_clerkenstein`) + casbin plural/singular gotcha → M4.
- **Confirmed-covered (Fate 2):** express-gate CI → M5 (already in its `In:`).
- **Dropped / deliberate non-goal:** nightly auto-reaper (M3-D2). **Escape-hatch (cross-release):** none.

**Verification:** rosetta-demo **78 tests** green + shellcheck clean; clerkenstein 218 funcs green `-race`; deploy
gate **100%/100% (7/7)**; **flake gate 5/5**. dev stack untouched (12 containers).
