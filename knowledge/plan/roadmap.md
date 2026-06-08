# Roadmap

Active development plan for **Project Rosetta** (the Anthropos documentation corpus + environment-
builder skills).

> **Designed 2026-06-02** from the Demo Environment + Clerkenstein brief, **refined 2026-06-02** to
> promote alignment measurement into a first-class discipline (new **M0**). 3 research agents over the
> Clerk integration, the staging/dev-env tooling, and the data/seeding surface — all verified against
> the cloned platform in `stack-dev/`. Gap analysis:
> [`.agentspace/scratch/roadmap-research-2026-06-02.md`](../../.agentspace/scratch/roadmap-research-2026-06-02.md).
>
> **v1.0 "body double" — SHIPPED 2026-06-03** (merged to `main`, tagged `v1.0`; full detail in `## Done` below).
> **v1.1 "show floor" — SHIPPED 2026-06-05** (merged to `main`, tagged `v1.1`; full detail in `## Done — v1.1`
> below). 8 milestones M3→M8: the 2-repo consolidation + demo/dev stacks + the production-safe seeding stack
> (framework + data-DNA + fleet) + the corpus product layer.
>
> **v1.2 "set dressing" — SHIPPED 2026-06-07** (tag `v1.2`, merged to `main`; designed 2026-06-05, **refined
> 2026-06-06** against live prod). 4 milestones M9a→M9b→M10→M11: a **dedicated `stack-snapshot` extension** that
> lifts M7c's two `waived` surfaces (`taxonomy` + `content`) to **100% data-DNA coverage** — capture the real
> *public* skill taxonomy + content library once from a **safe, low-impact source** (default a prod `pg_dump`), replay per-stack,
> *measured-faithful* via a new snapshot-fidelity dimension, with a tested **tenant-data firewall** (never customer
> data) + a **`.agentspace` manifest cache** (snapshots never land in any git repo). Detail in `## Done — v1.2` below.
>
> **v1.3 "stack party" — SHIPPED 2026-06-07** (tag `v1.3`, merged `--no-ff` → `main`; designed + built + closed
> 2026-06-07). 4 milestones M12→M13→M14→M15: the **dev/demo convergence** — dev stacks became first-class peers
> (local per-stack Directus + auto-snapshot + a light default seed on build), a **unified stack registry** that
> allocates the first-available N across dev+demo (no port collisions), one **generic `stack-*` skill set**
> (`stack-list`/`seed`/`snapshot`/`update` + `dev-up`/`dev-down`), and a code-cited **safety & security doc** (how
> the tooling never reads private data + never touches prod). Full detail in `## Done — v1.3` below. (The former
> v1.3 seeds — cloud store, S3 blobs, AI content, shareability — moved to **v1.4**.)
>
> **v1.3b "dress rehearsal" — IN DEVELOPMENT** (designed 2026-06-08; branch `release/01.3b-dress-rehearsal`). A
> **field-hardening release** between shipped v1.3 and future v1.4: v1.3 converged the dev/demo *model*, but the
> first real run of `/demo-up` surfaced 14 issues — a demo stack comes up **backend-only, unseeded, unverified**, and
> announces "UP" even when authz silently failed. v1.3b makes `/demo-up` produce a **full, populated, verified,
> demoable** stack and makes the repo honest about it: 5 milestones M16→M20 — land the applied fixes + restore doc
> truth, add re-run safety (idempotency + the first-run race), a post-bring-up verification net (`stack-verify` made
> offset-/scope-aware + auto-wired non-fatal), the **frontend tier** (next-web + studio-desk + ant-academy, tooling-only),
> and **lifecycle convergence** (demo-up auto set-dresses like dev-up + a cold-start capture path). All **tooling +
> docs only — zero platform-repo edits** (verified constraint). v1.4 (cloud store / S3 / AI content / shareability)
> is untouched. Brief: [`.agentspace/demo-up-issue.md`](../../.agentspace/demo-up-issue.md) +
> [`.agentspace/demo-up-frontend-plan.md`](../../.agentspace/demo-up-frontend-plan.md).

## Version plan

| Version | Codename | Theme | Milestones | Status |
|---------|----------|-------|------------|--------|
| **v1.0** | **body double** | A *measured* stand-in the platform can't tell from the real thing | M0 → M1 → { M1b ∥ M2 } → M2b → M2c | ✅ **SHIPPED 2026-06-03** (tag `v1.0`) |
| **v1.1** | **show floor** | The platform-operations extension framework (demo + dev, in 2 repos) | M3 ✅ → M4 ✅ → M5 ✅ → M6 ✅ → M7a ✅ → M7b ✅ → M7c ✅ → M8 ✅ | ✅ **SHIPPED 2026-06-05** (tag `v1.1`) |
| **v1.2** | **set dressing** | Richer demo worlds — the real *public* taxonomy + content library, measured-faithful, to 100% data-DNA coverage | M9a ✅ → M9b ✅ → M10 ✅ → M11 ✅ | ✅ **SHIPPED 2026-06-07** (tag `v1.2`) |
| **v1.3** | **stack party** | dev + demo stacks as first-class peers — local Directus, auto-snapshot + light seed, smart shared ports, one unified `stack-*` skill set | M12 ✅ → M13 ✅ → M14 ✅ → M15 ✅ | ✅ **SHIPPED 2026-06-07** (tag `v1.3`) |
| **v1.3b** | **dress rehearsal** | Field-hardening — make `/demo-up` produce a full, populated, verified, demoable stack (the gaps the first real run surfaced) | M16 ✅ → M17 → M18 → M19 → M20 | 🚧 **IN DEVELOPMENT** (designed 2026-06-08; M16 closed 2026-06-08) |

The whole initiative layers a **second corpus + skill set on top of** the existing dev-environment
tooling, to build disposable demo environments. Hard constraints: **no modification to any platform
repo** (current or future) and **no disruption to the dev environment**. Each local stack lives in its
own gitignored **`stack-*/`** workspace spanning one full stack — its platform service repos *plus* its
own clone of `rosetta-extensions`: `stack-dev` (dev), `stack-demo` (demo), `stack-dev-2` (secondary
dev), and future `stack-stage` / `stack-tests`. **Policy:** all code/scripts that operate the
corpus/platform on a spawned stack live in `rosetta-extensions` — never scattered in the rosetta corpus,
never authored ad-hoc inside a stack dir. New tooling is built + tested in the authoring copy at
`.agentspace/rosetta-extensions/`, tagged, then consumed per-stack as `stack-<role>/rosetta-extensions @ <tag>`
(rosetta = read-only doc corpus + dev-env skills; `rosetta-extensions` = the executable stack tooling).
Full brief: [`.agentspace/demo-environment-draft.md`](../../.agentspace/demo-environment-draft.md).

## In Development — v1.3b "dress rehearsal" (designed 2026-06-08)

**Theme:** v1.3 "stack party" converged the dev/demo **model**; running `/demo-up` for real revealed it didn't yet
converge the **experience**. A demo stack comes up **backend-only** (no UI), **unseeded** (no org/users → every
authorized route 403s), and **unverified** — `up-injected.sh` prints "UP. Clerk-free demo-N is live" with **zero
automated checks**, and in this very session it announced "UP" while the Sentinel authz policy had silently failed
to load (`casbin_rules` = 0). v1.3b is the **dress rehearsal**: the full run-through that makes `/demo-up` produce a
**full, populated, verified, demoable** stack — and makes the tooling honest about what it does. It addresses the 14
issues logged in [`.agentspace/demo-up-issue.md`](../../.agentspace/demo-up-issue.md).

> **Designed 2026-06-08** from the user's `/demo-up` field log (14 issues, several already analyzed/fixed by the
> prior session). **Phase 0a deferral audit GREEN** — the only open deferral (DEF-M10-01, S3 blob bytes + cloud
> store) is orthogonal and stays → **v1.4**; v1.3b adds zero new deferrals at design. **Phase 0b KB blind-areas:**
> the **frontend tier** and **post-bring-up verification** have no corpus anchor (→ M19/M18 each `Deliver →` one);
> idempotency/cold-start/auto-set-dress are anchored but need a contract section (→ M17/M20). Research +
> verification (3 agents, every issue claim re-checked against live code in `.agentspace/rosetta-extensions/`):
> [`.agentspace/scratch/roadmap-research-2026-06-08.md`](../../.agentspace/scratch/roadmap-research-2026-06-08.md).
>
> **Scope decided (user, 2026-06-08):** the **5-milestone spine** (M16→M20, one per issue cluster); the heavy
> features (frontend UI in M19, auto snapshot+seed in M20) are **default-on + skippable** (full parity with
> `dev-up`: one command = a real demo; `--no-ui` / `--no-setdress` to opt out). Codename **"dress rehearsal"** —
> a demo *is* a show; this is the run-through that makes it actually perform.

> **The two-repo split (the load-bearing distinction for this release).** Every milestone names which repo owns
> what, because v1.3b touches **both**:
> - **`rosetta-extensions`** (the executable tooling) holds all **scripts / Go / Python / per-section README+GUIDE /
>   its own `knowledge/` KB**. Built + tested in the **authoring copy** `.agentspace/rosetta-extensions/`, **tagged**
>   (`dress-rehearsal-mNN`), then **consumed per-stack** at the pinned tag — never authored ad-hoc inside a `stack-*/`.
> - **`rosetta`** (this corpus) holds the **`.claude/skills/*`, `corpus/*`, `CLAUDE.md`, READMEs** — the docs and the
>   skills that *drive* the CLIs.
>
> So a typical milestone = a code change in `rosetta-extensions` (authoring → tag → consume) **+** the skill/corpus
> doc in `rosetta` that describes it. The `Delivers →` lines below split accordingly.

### M16: Land the field fixes + restore doc truth
**Status:** `done` (completed 2026-06-08) · **Shape:** `section` · **Complexity:** small–medium · **Dir:** [m16-land-fixes/](releases/01.3b-dress-rehearsal/m16-land-fixes/)
**Closed 2026-06-08** (build → 2 harden passes → close review → merged to `release/01.3b-dress-rehearsal`). **The first milestone of v1.3b "dress rehearsal"** — it lands the honest baseline (publishes the two stranded field fixes + restores doc truth) that every other milestone builds on. A **docs/publish/rename milestone** (the M14 shape): the two functional changes (the ISSUE-1 devpath rename-resolution + the ISSUE-7 migrate-race `|| echo 0` under `set -e`) were pre-applied in a prior session; M16 made them **durable + public + fenced**. Delivered, in the SEPARATE nested `.agentspace/rosetta-extensions` repo (gitignored from rosetta, pushed to `origin`): the **publish** (extensions `main` `a31d70b..e6161b0` to `origin`; the local-only `stack-party-devpath-fix` tag superseded — not deleted, never pushed); the **stack-core rename migration** making `stack-dev` the documented default across all 5 workspace-resolvers (`up-injected.sh`, `migrate-demo.sh`, `rosetta-demo` ×2, `dev-stack`) + `gen_override.py`/`clone_repos.py`, demoting `anthropos-dev` to the **single intentional legacy-alias fallback** (`[ -d ] || …anthropos-dev`, each line legacy-marked — M16-D2); the **prose sweep** (`demo-stack/README.md`/`GUIDE.md`, `dev-stack/README.md`, `gen_override.py` docstring → `stack-dev/platform`); the **GUIDE header truth-up** (remote-exists / 21 tests / `/stack-list` / v1.3 — was no-remote / 78 / `/demo-status` / v1.1·M3); the **pytest doc fix** (`pytest tests/ -v` + the 3.11/3.12 note); the **extensions `knowledge/` KB** (the consumption version-jump + the per-milestone tag scheme). In **rosetta**: a consolidated **`corpus/ops/rosetta_demo.md` stack-dev layout + back-compat note** (code snippet + CLAUDE.md cross-link + don't-reintroduce-bare-`anthropos-dev` guidance); `corpus/` had **0** stray `anthropos-dev` (sweep = verify no-op). **Harden** (2 passes) made the contract a *test*: 8 new demo-stack guards (`TestRenameDrift` 3 — `stack-dev` must lead `anthropos-dev` in every resolver + no unmarked `anthropos-dev`; `TestGuideDocTruth` 2 — the advertised test count + the documented `pytest` entrypoint pinned to live; `TestMigrateRaceGuard` 3 — the ISSUE-7 fence, negative-tested). Close found **1 finding** (the docs review surfaced `clerkenstein/knowledge/glossary.md` still naming the old `anthropos-demo/` workspace — the **last** non-fallback stale workspace name in the entire repo; landed Fate-1 as M16-D8, reversing build-time M16-D5's boundary deferral; ext `e6161b0`); the close **reconciled** the `dress-rehearsal-m16` tag from the build HEAD `44edc09` to the final HEAD `e6161b0` (the tag trailed the 2 harden commits + the 1 close commit — a sanctioned forced tag re-point, force-pushed; cf. v1.3 M9a) and **re-consumed** `stack-demo/rosetta-extensions` to it (three-way agreement: origin tag = authoring HEAD = consumed HEAD = `e6161b0`; the only per-stack clone — `stack-dev/` has none, M16-D6). 2 adversarial scenarios recorded (both-roots-exist resolves deterministically to `stack-dev`; the race is fenced). Deferral re-audit **GREEN** (1 in-release single — the live docker-harness migrate-race *behavior* test → **Fate-2 → M17**, already owned by M17's `In:` scope, no overview edit, M16-D7; 1 inherited DEF-M10-01 → v1.4 signed/not-aged; 0 repeat/chronic/aged-out — M16 is the release's first milestone, nothing inherited within v1.3b). Decisions M16-D1…D8. Extensions: tag **`dress-rehearsal-m16`** @ `e6161b0` (per-milestone `dress-rehearsal-mNN` naming, M16-D1; a final `v1.3.1` lands at close-release). Python test funcs 174→**182** (demo-stack 13→21, +8 guards); Go **713** unchanged (M16 touched no Go); flake **0** (5/5); shellcheck + py_compile CLEAN on all touched scripts. Retro: [m16-land-fixes/retro.md](releases/01.3b-dress-rehearsal/m16-land-fixes/retro.md). **Next:** M17 (bring-up re-run safety — idempotency + first-run race).
**Goal:** Make the two already-applied fixes durable and public, finish the `anthropos-dev → stack-dev` rename as
the *documented default*, and clear the stale tooling docs — so the repo tells the truth before more work lands on it.
**Scope:**
  - In:
    - **`rosetta-extensions`:** **push** the local fixes (`547de17` devpath + `ed72e94` migrate-race — currently 2
      commits ahead of `origin`, on a local-only `stack-party-devpath-fix` tag) to `origin` under a proper
      `dress-rehearsal-m16` tag, then re-consume per-stack (ISSUE-1①/ISSUE-7 push); the **stack-core rename
      migration** making `stack-dev` the documented default + demoting `anthropos-dev` to a single intentional
      "legacy alias" mention (ISSUE-1②); the **prose sweep** (`demo-stack/README.md:12`, `GUIDE.md:17`,
      `dev-stack/README.md:73`, `stack-core/gen_override.py:4` docstring — ISSUE-2); **GUIDE.md header truth** (remote
      exists; **13** tests not 78; `/stack-list` not `/demo-status`; **v1.3** not v1.1/M3 — ISSUE-3); the **pytest
      doc fix** (`pytest tests/` + a 3.11/3.12 note — ISSUE-4); refresh the repo's own `knowledge/` KB where it
      repeats any of these.
    - **`rosetta`:** sweep any residual `anthropos-dev` in `corpus/`; note the expected consumption version-jump (ISSUE-5).
  - Out: the race/idempotency *behavior* work (M17); anything frontend/verify/set-dress.
**Depends on:** none. **Parallel with:** none (the honest baseline every other milestone builds on).
**Estimated complexity:** small–medium
**Open questions:** the re-tag/version scheme (lean: per-milestone `dress-rehearsal-mNN` tags + a final `v1.3.1`
release tag at close, matching the established convention); whether to keep `anthropos-dev` back-compat in the
scripts forever or sunset it (lean: keep — it's a one-line fallback that costs nothing).
**KB dependencies:** `corpus/ops/rosetta_demo.md`, the `rosetta-extensions` GUIDE/README set + its `knowledge/` KB.
**Delivers → updates the `rosetta-extensions/knowledge/` KB + GUIDE/README truth-up** (rosetta-extensions) **+ a
consolidated `corpus/ops/` note on the `stack-dev` layout + the back-compat fallback** (rosetta).

### M17: Bring-up re-run safety — idempotency + first-run race
**Status:** `planned`
**Shape:** `section`
**Goal:** A re-run of migrate / snapshot-replay / seed is either safe-and-idempotent or fails loudly with a guard —
never silently doubles data or aborts mid-surface.
**Scope:**
  - In (all **`rosetta-extensions`** code; one **`rosetta`** doc):
    - the **`set -e` first-run-race audit** across the bring-up scripts (the same class as the fixed ISSUE-7
      migrate race — sweep `up-injected.sh`/`rosetta-demo`/`dev-stack`/`dev-setdress.sh`) + an optional
      **"wait-for-sentinel-ready"** defense-in-depth (ISSUE-7 residual);
    - **`stacksnap replay`** re-run protection — a per-stack-isolated `TRUNCATE`/skip/`ON CONFLICT` before the bare
      `COPY` so a second replay doesn't duplicate-key-abort mid-surface or silently double (`stack-snapshot/pg/pg.go`,
      `replay/replay.go` — ISSUE-11);
    - **`stackseed`** re-run protection — `ON CONFLICT` for the casbin g2 grant + the deterministic-UUID rows; **fix
      the stale `--reset` truncate list** to include the session/activity tables it currently skips
      (`stack-seeding/seeders/identity.go`, `cmd/stackseed/main.go` — ISSUE-11);
    - an explicit, **tested** idempotency contract for each component.
  - Out: the verify net (M18); the auto-chaining of snapshot/seed (M20 — M17 only makes the *primitives* re-run-safe).
**Depends on:** M16 (clean pushed baseline). **Parallel with:** **M18 (yes-with-caveats** — different surfaces;
both touch `up-injected.sh` only in different regions).
**Estimated complexity:** medium
**Open questions:** `TRUNCATE`-and-reload vs `ON CONFLICT DO NOTHING` for replay (lean: TRUNCATE the per-stack
target surface then reload — simplest correct semantics, and the target is always per-stack-isolated); whether to
make re-run safety automatic or behind an explicit `--idempotent`/`--force` (lean: safe-by-default with a loud guard).
**KB dependencies:** `corpus/ops/snapshot-spec.md`, `corpus/ops/seeding-spec.md`, the `stack-seeding`/`stack-snapshot` sections.
**Delivers → `corpus/ops/idempotency.md`** (rosetta — net-new: the per-component re-run verdicts + the
teardown-then-redo model + the new guards) **+ the new guards in `stack-seeding`/`stack-snapshot`** (rosetta-extensions).
**Risk (data-safety):** a `TRUNCATE` must *only* ever hit a per-stack-isolated offset target — the n=0 + prod-isolation
guards stay inviolate; tests pin the target class.

### M18: The verification safety net
**Status:** `planned`
**Shape:** `section`
**Goal:** Teach `stack-verify` to target an *offset* stack and scope to the services actually brought up, then
auto-run it (non-fatal) at the tail of every bring-up — so "UP" means **verified-working**, not just *containers-started*.
**Scope:**
  - In (**`rosetta-extensions`** code; one **`rosetta`** doc):
    - **project/offset awareness** — derive the `demo-N`/`dev-N` prefix + the N×10000 port offset from
      `STACK_ROOT`/the unified registry (today `lib/services.sh:25-39` hardcodes 12 `anthropos-*-1` names at **base**
      ports, so it reports an offset stack entirely `down` — ISSUE-12b);
    - a **service/profile filter** intersecting the checked set with what was requested (so a reduced bring-up isn't
      a wall of false `down`s);
    - **fix the undefined `$DEVDIR` → `$STACK_ROOT` bug** (`repos/run.sh:108`, `census/inventory.sh:75` — ISSUE-12);
    - the **cheap-win asserts available today** (`curl -fsS .../api/health` + `SELECT count(*) FROM
      sentinel.casbin_rules > 0` on the stack's offset ports at the bring-up tail — the exact ISSUE-7 silent-failure
      catcher — ISSUE-14);
    - the **auto-wired scoped `verify live`** at the bring-up tail, **default-on + non-fatal** (mirroring
      `dev-setdress`'s proven pattern — ISSUE-12c/ISSUE-14).
  - Out: verifying the *frontend* ports (added in M19, where the frontends first exist); deep behavioural/e2e probes
    (that's the `/test-platform` skill's job — M18 is the always-on smoke net).
**Depends on:** M16. **Parallel with:** M17 (yes-with-caveats).
**Estimated complexity:** medium–large
**Open questions:** derive the offset from `STACK_ROOT` parsing vs reading the registry's recorded host ports (lean:
read the registry — it already records resolved ports per M12); how loud "non-fatal but failed" should be (lean: a
clear ⚠️ block + a one-line "run `/test-platform N` to dig in").
**KB dependencies:** `corpus/ops/run_guide.md`, `corpus/ops/rosetta_demo.md`, the `stack-verify` section, the `/test-platform` skill.
**Delivers → `corpus/ops/verification.md`** (rosetta — net-new: the auto-verify contract + the offset/scope model)
**+ the offset-/scope-aware `stack-verify` + bring-up wiring** (rosetta-extensions).
**Risk (correctness):** a mis-derived offset would false-positive "down" — the very bug it fixes. Mitigate: derive
from the registry's recorded ports; non-fatal so a verify bug never blocks a genuinely good stack.

### M19: The demo-up frontend tier
**Status:** `planned`
**Shape:** `section`
**Goal:** `/demo-up` brings up the full UI — next-web + studio-desk at offset ports (per-demo **cached** image from
the **unmodified** platform Dockerfile) + ant-academy natively — so a demo is actually demoable, on a 16 GB Mac.
**Scope:**
  - In:
    - **`rosetta-extensions`:** extend `stack-injection/gen_injected_override.py` to **emit `next-web-app` +
      `studio-desk`** (offset ports via `ports:!override`, `image: demo-N-*` + `mem_limit:1g`, additive override —
      today it emits backend-only, ISSUE-8); `up-injected.sh` **builds the two frontends serially, before compose
      up**, from the unmodified Dockerfiles with **offset-URL build-args + the minted Clerk pk** via a gitignored
      `.env.local`/BuildKit overlay, **tag-guarded for cache reuse**; ship a sibling `.dockerignore` (5.6 GB context
      → <100 MB); **launch (or document) ant-academy natively** (port 3077, its own `.env`,
      `REQUIRE_ORGANIZATION_MEMBERSHIP=0`); a **12 GB Docker-VM pre-flight assert** (ISSUE-6/ISSUE-9); **register the
      frontend ports** so M18's verify net covers them. **Default-on + skippable** (`--no-ui`).
    - **`rosetta`:** update the `demo-up` SKILL.md (the UI is now in scope) + author the frontend-tier corpus doc.
  - Out: the **optional upstream platform PR** for *true* zero-rebrebuild (runtime-rewrites + `__env.js` +
    `output:standalone`) — it edits platform repos → **forbidden / user-owned**, documented as a follow-up (like
    v1.4's deploy-CI item), not built here.
**Depends on:** **M18** (so the verify net can cover the new frontend services). **Parallel with:** none.
**Estimated complexity:** large (the meatiest milestone of v1.3b)
**Open questions:** ant-academy native-launch *by* the tool vs documented-manual (lean: launch it, fall back to a
documented step if the native run proves fiddly); whether to pre-warm the frontend image cache as part of `dev-up`
too (defer — demo-first).
**KB dependencies:** [`.agentspace/demo-up-frontend-plan.md`](../../.agentspace/demo-up-frontend-plan.md) (the
verified tooling-only plan), `corpus/services/next-web-app.md`, `corpus/services/ant-academy.md`,
`corpus/services/studio-desk.md` *(if present)*, `corpus/ops/rosetta_demo.md`.
**Delivers → `corpus/ops/demo/frontend-tier.md`** (rosetta — net-new: ports, per-demo build, Clerk-pk baking, the
12 GB VM prereq, the honest "one ~3-min cached build per new demo-N" residual) **+ updated `demo-up` skill** (rosetta)
**+ the frontend-emitting override generator + per-demo build in `up-injected.sh`** (rosetta-extensions).
**Risk (scope+resource):** the ~3.7 GB / ~3 min per-frontend build swap-thrashes on an undersized VM (the original
"hour"). Mitigate: 12 GB preflight assert, serial cached builds, `mem_limit`, the sibling `.dockerignore`. **Hard
line: zero platform-repo edits — the repo is a build *context* only, the Dockerfile is unmodified (verified achievable).**

### M20: Lifecycle convergence — demo-up auto set-dress + cold-start capture
**Status:** `planned`
**Shape:** `section`
**Goal:** Close the dev↔demo asymmetry — `/demo-up` auto set-dresses (snapshot → seed, default-on + non-fatal) like
`dev-up` already does — and unblock the *real* catalog on a fresh box that has no safe `--dsn`.
**Scope:**
  - In:
    - **`rosetta-extensions`:** **chain** `stacksnap replay` → `stackseed` into `up-injected.sh` after migrate
      (reuse M13's proven `dev-setdress` pass; **default-on + non-fatal**, `--no-setdress` escape — ISSUE-10);
      enforce the **atomicity contract** (a partial snapshot with no seed = 403s, so it's both-or-neither, and the
      M17 re-run guards make a retry safe); the **cold-start capture** path (ISSUE-13) — a documented, prod-safe
      **DSN-export / restore-a-`pg_dump`-then-`--dsn`** workflow (the sanctioned route), **plus a spike** on whether
      a thin MCP-paging capture adapter (read via the wired `postgres` MCP) is cheap enough to build vs document.
    - **`rosetta`:** update the `demo-up`/`demo-down` skills + the `corpus/ops/demo/` recipe family for auto set-dress.
  - Out: **S3 media blob bytes + the cloud `SnapshotStore` backend** (DEF-M10-01 → **v1.4**, signed — untouched);
    AI-generated content (v1.4).
**Depends on:** **M18** (the post-set-dress verify) **+ M19** (the full experience) **+ M17** (re-run-safe primitives
make auto-chaining safe to retry). **Parallel with:** none (the closing milestone).
**Estimated complexity:** medium–large
**Open questions:** ISSUE-13 — build the MCP-paging capture adapter vs document the DSN-export step only (lean:
**document the sanctioned DSN-export/restore path now**, spike the MCP adapter only if cheap — the MCP is a query
tool, not a `--dsn` `stacksnap` can `COPY` through); whether auto-set-dress on `demo-up` should default to the
`dev-min`-style light seed or a demo preset (lean: a demo preset like `small-200`, since demos want a fuller world
than dev — confirm at build).
**KB dependencies:** `corpus/ops/snapshot-spec.md`, `corpus/ops/seeding-spec.md`, `corpus/ops/db-access.md`, the
`dev-setdress` mechanism (M13), `corpus/ops/demo/README.md` + the recipe family.
**Delivers → `corpus/ops/snapshot-cold-start.md`** (rosetta — net-new: the fresh-box capture workflow + the MCP
limitation + the sanctioned paths) **+ updated demo recipes/skills for auto set-dress** (rosetta) **+ the set-dress
chaining in `up-injected.sh` + the cold-start capture path** (rosetta-extensions).
**Risk (prod-safety):** auto set-dress + cold-start capture must preserve M13/M15's guarantees — **read-only bounded
capture, the tenant firewall, per-stack isolation, and a confirmation before any prod-touching read**. Mitigate:
reuse M13's proven set-dress pass verbatim; capture stays behind M9a's capture-source policy + `AssertPublicOnly`.

### Execution graph (v1.3b)
```
v1.3b "dress rehearsal" — make /demo-up produce a full, populated, verified, demoable stack
  M16 (land fixes + doc truth)
   └─→ M17 (re-run safety) ──┐
       M18 (verify net) ─────┴─→ M19 (frontend tier) ─→ M20 (auto set-dress + cold-start)
            (M17 ∥ M18 feasible — different surfaces / different up-injected.sh regions; lean sequential per the spine discipline)
```
**Mostly sequential.** M16 lands the honest baseline (pushes the stranded fixes). M17 + M18 harden the *primitives*
(re-run safety) and add the *net* (verification) — the one feasible parallel pair (different surfaces; both touch
`up-injected.sh` only in different regions), but the spine discipline leans sequential. M19 builds the **frontend
tier** (the meatiest; depends on M18 so the verify net covers the new services). M20 **converges the lifecycle**
(auto set-dress + cold-start) on top of the full experience — the closing milestone. **No B-milestones — the whole
release *is* the tooling layer** (the v1.3/M14 precedent).

### Parallelism matrix (v1.3b)
| Pair | Can parallelize | Shared surface | Merge risk | Strategy |
|------|-----------------|----------------|-----------|----------|
| **M17 ∥ M18** | yes-with-caveats | `up-injected.sh` (M17: migrate region · M18: tail) | low | land M17 first (smaller), then M18's tail-append; lean sequential per spine discipline |
| M16 ∥ * | no | the honest baseline | — | M16 goes first |
| M19 ∥ M18 | no | M19 needs M18's verify to cover its new frontend services | — | sequential |
| M20 ∥ M19 | no | M20 = full experience on top of M19 | — | M20 closes |

### Risks (v1.3b)
- **(M19, scope+resource — blocks-quality)** the per-frontend build (~3.7 GB / ~3 min) swap-thrashes on an
  undersized VM. Mitigate: 12 GB Docker-VM preflight assert, serial cached builds, `mem_limit`, the sibling `.dockerignore`.
- **(M18, correctness — degrades-quality)** a mis-derived offset false-positives "down" — the bug it fixes.
  Mitigate: derive from the registry's recorded ports; non-fatal so it never blocks a good stack.
- **(M17, data-safety — blocks-prod-safety)** a `TRUNCATE` on the wrong target = data loss. Mitigate: TRUNCATE only
  per-stack-isolated offset targets; the n=0 + isolation guards hold; tests pin the target class.
- **(M20, prod-safety — blocks-prod-safety)** auto set-dress + cold-start capture must keep M13/M15's read-only
  bounded capture + tenant firewall + per-stack isolation + the confirm-before-prod-read. Mitigate: reuse M13's pass; capture stays behind the M9a policy.
- **(M16, blast-radius — nice-to-resolve)** the **push** is v1.3b's first outward-facing action (the local fixes go
  public). Mitigate: re-tag cleanly (retire the local `stack-party-devpath-fix`), push, re-consume per-stack.
- **(cross-cut)** **zero platform-repo edits** — the verified hard line. Every frontend/build change uses the repo
  as a build *context* only; the *true zero-rebuild* path is an explicitly-OUT, user-owned upstream PR.

### Open decisions (resolve during build)
Re-tag/version scheme — M16 (lean `dress-rehearsal-mNN` + final `v1.3.1`); replay re-run = TRUNCATE-reload vs
ON CONFLICT — M17 (lean TRUNCATE per-stack target); offset from STACK_ROOT-parse vs registry-recorded-ports — M18
(lean registry); ant-academy native-launch by-tool vs documented — M19 (lean launch); ISSUE-13 MCP-adapter vs
documented-DSN-export — M20 (lean document now, spike adapter if cheap); demo auto-set-dress preset — M20 (lean a
demo preset over `dev-min`).

## Done — v1.3 "stack party" (SHIPPED 2026-06-07 · tag `v1.3`)

**Theme:** v1.0 made the platform run Clerk-free; v1.1 built the demo/dev stack framework; v1.2 set-dressed *demo*
stacks with the real public catalog. v1.3 **throws the party** — it makes **dev and demo stacks first-class peers**
that work the same way. A dev stack gets the demo treatment (its own **local Directus**, **auto-snapshot** of the
real reference data on build, a **light default seed** so it's never empty); a **unified stack registry** allocates
the first-available N across both kinds so they never collide on ports; and one **generic `stack-*` skill set**
operates any stack. The safety story (read-side private-data avoidance + write-side prod-protection) is consolidated
into a dedicated doc, and **both** the rosetta corpus and the `rosetta-extensions` knowledge base are brought to the
converged model.

> **Designed 2026-06-07** from 6 user requirements (the dev/demo convergence). Phase 0a deferral audit **GREEN**
> (v1.2 close — 1 open item, DEF-M10-01 S3 blobs + the cloud-store escape-hatch). **Scope decided (user, 2026-06-07):**
> the former v1.3 seeds (cloud store, S3 blobs, AI content, shareability, more mirrors) **move to v1.4**; v1.3 is the
> tight convergence release. M13 stays one milestone; the operation-skill renames are a **hard rename, no aliases**.
> Phase 0b KB blind-area: the **safety & security doc** is the one blind area → M15 `Delivers →` it. Research:
> [`.agentspace/scratch/roadmap-research-2026-06-07.md`](../../.agentspace/scratch/roadmap-research-2026-06-07.md).

### M12: Unified stack registry + first-available-N allocation
**Status:** `done` (completed 2026-06-07) · **Shape:** `section` · **Complexity:** medium · **Dir:** [m12-stack-registry/](releases/archive/01.30-stack-party/m12-stack-registry/)
**Closed 2026-06-07** (build → 3 harden passes → close review → merged to `release/01.30-stack-party`). **The first milestone of v1.3** and the only non-Go surface to date — the unified registry is **Python** (`stack-core/stack_registry.py`) + shell-CLI wiring. Delivered: a **single shared module** `stack_registry.py` (M12-D1 — one shared N-pool across both kinds, so `dev-1`/`demo-1` can never coexist; runtime file `stack-core/.stacks/registry.json`, gitignored) implementing **first-available-N allocation** — `allocate(type, n=None)` auto-allocates the lowest free N across dev+demo or validates an explicit N, the whole **read-reconcile-pick-write under one `fcntl.flock(LOCK_EX)`** on a sidecar `.lock` (portable macOS+Linux, M12-D3) with **atomic temp+`os.replace` writes** + corrupt-JSON recovery; **registry-of-record + `docker ps` reconcile that only ADDS used-N, never subtracts** (M12-D2 — adopts unmanaged live stacks; `release()` is the only free path = the race guard). Both up-paths wired (`dev-stack` + `rosetta-demo`): `up [N]` (explicit or `auto`/omit), an **ERR-trap that frees the slot on a failed bring-up** (installed *before* the post-allocation re-guard — the §2 review fix), `set-ports` records resolved host ports, `down` releases; `status` leads with the unified dev+demo view. KB-1 resolved (the GUIDE.md "registry assigns N" claim converged to truth). The guarantee — `dev, demo, dev, demo, demo` → `dev-1, demo-2, dev-3, demo-4, demo-5`, any interleaving, either CLI — verified at the **OS-process level** (12-process auto-allocation + a close-added 6-process explicit-N collision, not just threads). Extensions: tag `stack-party-m12` @ `be1c979` (v1.3 release-and-milestone tag naming). **89 Python test funcs** across the 3 M12-touched suites (stack-core 54 / dev-stack 22 / demo-stack 13; +52 net new), `stack_registry.py` 98% (2 structurally-unreachable misses); flake **0** (5/5); shellcheck + py_compile clean; Go total 708 unchanged (M12 touched no Go). Close: scope/code/docs/tests **GREEN**, **3 findings** (1 test added — the concurrent explicit-N collision; 2 decision-triage → archive); 4 adversarial scenarios recorded; deferral re-audit **GREEN** (1 inherited DEF-M10-01 → v1.4 signed, 0 repeat/aged/new — M12 added zero deferrals). Decisions M12-D1…D3 + KB-1 + Q1/Q2/Q3 resolved. Retro: [m12-stack-registry/retro.md](releases/archive/01.30-stack-party/m12-stack-registry/retro.md). **Next:** M13 (dev peers — local Directus + auto-snapshot + light seed).
**Goal:** One shared registry across **dev + demo** that tracks live stacks and allocates the **lowest free N**, so
bring-ups never collide on ports (build `dev, demo, dev, demo, demo` → `dev-1, demo-2, dev-3, demo-4, demo-5`).
**Scope:** In: a **unified stack registry** (extend the demo `registry.json` to span both kinds — type + N + ports +
status) in `stack-core`; a **first-available-N allocator** (reconcile the registry against live `docker ps`, return
the lowest free N) — race-safe (locked write); the up-paths **accept explicit-N or auto-allocate**; teardown frees
the slot. Out: the skill renames that consume it (M14); dev local-Directus/snapshot/seed (M13).
**Depends on:** v1.1's `stack-core` (the port-offset engine) + the demo-stack `registry.json`. **Parallel with:** M13 (feasible — different surfaces; lean sequential so M13's dev bring-up consumes the registry).
**Open questions:** registry as a lockfile vs `docker ps`-derived (lean: registry-of-record + `docker ps` reconcile); how a manually-started stack (outside the skills) is reconciled.
**KB dependencies:** `corpus/ops/rosetta_demo.md` (the demo registry/ports), the `stack-core`/`demo-stack`/`dev-stack` extension sections.
**Delivers → updates `corpus/ops/rosetta_demo.md`** (the unified registry + first-available-N model).

### M13: Dev stacks as full-fidelity peers — local Directus + auto-snapshot + light seed
**Status:** `done` (completed 2026-06-07) · **Shape:** `section` · **Complexity:** large · **Dir:** [m13-dev-peers/](releases/archive/01.30-stack-party/m13-dev-peers/)
**Closed 2026-06-07** (build → 1 harden pass → close review → merged to `release/01.30-stack-party`). **The meatiest milestone of v1.3** — it converges dev with demo for **DATA** (M12 converged them for N-allocation). Delivered: a `dev-stack up` bring-up now gets the **demo treatment by default** via a new **set-dressing pass** (`dev-stack/dev-setdress.sh`, default-on + non-fatal): (1) a **per-stack Directus** — the M10 store-fork recipe + the prod-Directus firewall, both made **executable** by a NEW small Go runner `stack-snapshot/cmd/provision-plan` that turns the library-only M10 `directus.ProvisionPlan`/`EnvContract`/`Validate` contract into a CLI (M13-D2 — one source of truth for the recipe AND the firewall, shared by both kinds; `--check-env` hard-aborts the pass before any replay if the per-stack Directus env ever resolves to prod); (2) a **cache-first `stacksnap replay`** of the public `taxonomy`+`directus` surfaces into `dev-N` (replay never captures; a stale/missing cache is a warning, the seed still runs — M13-D3 default-on-but-non-fatal, resolving the bring-up-heaviness Q1 + the default-on Q2); (3) the new **`dev-min` seed preset** (`stack-seeding/presets/dev-min.seed.yaml` — ~1 org + ~10 users + 1mo, fixed admin `dev@anthropos.test` → login 200; the role-mix floor, M13-D1 resolving Q3) so the stack is never empty. The **n=0-dev guard is doubled** — the set-dress pass refuses N=0 without `--force`, a second layer above `stackseed --reset`'s own refusal. **Escapes:** `--no-snapshot` (seed only) / `--no-setdress` (bare bring-up). **Prod-safety held unchanged:** capture is never run on dev (replay only — a per-stack WRITE to the isolated offset Postgres), media stays **refs-only** (blob bytes = v1.4), `AssertPublicOnly` + the isolation guard both hold. Corpus: `snapshot-spec.md` (the "Dev as a full-fidelity peer" section + dev as a replay target) + `seeding-spec.md` (the shipped-presets table + the dev-min/dev-auto-seed subsection + the two-layer n=0 guard); the stale "cloud/S3 store = v1.3" claim fixed to **v1.4**. Harden fixed **3 robustness bugs** inline (whitespace-env firewall fail-closed · the `--no-snapshot` re-run-hint leak · a `set -u` trailing-flag error), each mutation-pinned, +10 regression/edge tests. Close found **6 findings** (1 doc must-fix: the stack-snapshot README Packages table was missing the new `cmd/provision-plan` row — a per-unit handbook-index miss, fixed; 5 GREEN/no-action incl. 4 adversarial scenarios + 3 decision-triage blends ref-tagged); deferral re-audit **GREEN** (1 inherited DEF-M10-01 → v1.4 signed/unchanged, 0 repeat/aged/new — M13 added **zero** deferrals, all Fate-1). Decisions M13-D1…D4 + Q1/Q2/Q3 resolved. Extensions: tag **`stack-party-m13`** @ `cca4464`. Go test funcs 708→**720** (stack-snapshot 212→223 the provision-plan runner; stack-seeding 232→233 the dev-min pin); dev-stack pytest 33→**38**; flake **0** (5/5 both languages); both Go modules `-race` + gofmt + `go vet` clean; both CLIs shellcheck-clean. Retro: [m13-dev-peers/retro.md](releases/archive/01.30-stack-party/m13-dev-peers/retro.md). **Next:** M14 (unified `stack-*` skills + `dev-up`/`dev-down`, hard-renamed).
**Goal:** A freshly-built dev stack gets the demo treatment — its **own local per-stack Directus** (no longer
pointing at shared prod), an **auto-snapshot** on build that fills the real public reference data (taxonomy + the
now-local content), and a **light default seed** so it's never empty.
**Scope:** In: wire the dev bring-up to **spawn a per-stack Directus** (reuse M10's `stack-snapshot/directus/provision.go`
per-stack-Directus mechanism) + repoint the dev CMS at it; **auto-run `stacksnap replay`** (taxonomy + directus) as
part of dev build (cache-first, fast); a **`dev-min` seed preset** (smaller than demo's `small-200` — ~1 org + ~10
users + minimal activity) applied on build; keep the n=0-dev-reset guard. Out: the generic skills (M14); blob bytes
(v1.4 — refs-only, as for demo).
**Depends on:** **M12** (consumes the registry for dev-N) + v1.2's M10 (the per-stack Directus + the content snapshot) + v1.1's M6 (dev-stack) + M7 (seeding). **Parallel with:** none (gates M14).
**Open questions:** does spawning Directus + snapshot + seed make dev bring-up too heavy? (mitigate: cache-first snapshot, minimal seed, reuse M10's provision); should the auto-snapshot be default-on or opt-in for dev (lean: default-on, `--no-snapshot` to skip).
**KB dependencies:** `corpus/ops/snapshot-spec.md` (capture/replay + the per-stack Directus store), `corpus/ops/seeding-spec.md` (presets + the dev/n=0 guard), `corpus/services/cms.md` (Directus), the `dev-stack` extension section.
**Delivers → updates `corpus/ops/seeding-spec.md`** (the `dev-min` preset + dev auto-seed) **+ `corpus/ops/snapshot-spec.md`** (dev as a replay target + the local-Directus-on-dev path).

### M14: Unified `stack-*` skills + `dev-up`/`dev-down`
**Status:** `done` (completed 2026-06-07) · **Shape:** `section` · **Complexity:** large · **Dir:** [m14-unified-skills/](releases/archive/01.30-stack-party/m14-unified-skills/)
**Closed 2026-06-07** (build → 1 harden pass → close review → merged to `release/01.30-stack-party`). **The tooling-convergence milestone of v1.3** — it unifies the operation skills onto both stack kinds via a **hard rename, no aliases** (user 2026-06-07; the clean-break blast radius contained by sweeping every in-repo reference). Delivered: **`dev-up`** (consolidates the former `setup-platform` + `start-platform` into one dev bring-up — element-by-element diff against both former bodies confirmed **no dropped step**: STEP RUN discipline, confirmation policy, ops reports, the 12-container set, error recovery; `dev-up N` for N≥1 spins additional isolated dev-N; drives the M13 set-dress flow, M14-D2) + net-new **`dev-down`**; the 4 ops skills hard-renamed each taking a `dev-N|demo-N` target — **`stack-list`** (←`demo-status`), **`stack-seed`** (←`demo-seed`), **`stack-snapshot`** (←`demo-snapshot`), **`stack-update`** (←`update-platform`); the **6 old skill dirs removed** (no shims); a **full reference sweep** (CLAUDE.md skill table rewritten to the 14-skill set, root+corpus READMEs, all `corpus/ops/` guides + the `demo/` recipe family, CHANGELOG — new Unreleased v1.3 Added/Changed/Removed, v1.1/v1.2 dated entries left immutable, M14-D6); `demo-up`/`demo-down` retained + aligned with the dev lifecycle. The `--preset NAME` UX is kept as a **skill-level shorthand** mapped to `--seed presets/NAME.seed.yaml` (M14-D5, the PR-review fix — the `stackseed` binary only knows `--seed <path>`). Extensions: a doc-only reference sweep on `main` (`b37e831`: dev-stack header/README, demo-stack GUIDE, stack-verify run.sh) + the harden reference-integrity guard (`33fc525`: re-pointed the stacksnap drift-guard comments `/demo-snapshot`→`/stack-snapshot` + added **`TestDocSourceSkillRename_M14`** — asserts the renamed skill dir exists + the retired one stays gone, turning silent comment-rot/dir-resurrection into a test failure; negative-tested). The skill namespace `stack-snapshot` is deliberately distinct from the extensions **section** `stack-snapshot/` (the skill drives the `stacksnap` CLI that section builds — called out in the body). Tag **`stack-party-m14`** @ `33fc525` (v1.3 release-and-milestone naming, matching M12 `stack-party-m12` @ `be1c979` + M13 `stack-party-m13` @ `cca4464`). Close: scope/code/docs/tests **GREEN** with **0 findings** — a clean straight-through close; the close re-verified the 4 reference-integrity dimensions **independently** (rename invariant **0 live stragglers** project-wide, CLAUDE.md ⇄ 14 skill dirs perfect bijection, SKILL.md `name:`⇄dir 0 mismatches, skill→CLI contract resolves) + all `../`-relative doc links resolve + 0 broken links to deleted dirs. Deferral re-audit **GREEN** (1 inherited DEF-M10-01 → v1.4 signed/unchanged/not-aged, 0 repeat/aged/new — M14 added **zero** deferrals, all Fate-1). Decisions M14-D1…D6 (D2/D3/D4 resolve Q1/Q2/Q3; D5 = the `--preset` fix; D6 = CHANGELOG convention). Go test funcs 720→**721** (stack-snapshot 223→224: the rename guard; alignment/clerkenstein/stack-seeding unchanged); all 4 Go modules `-race -count=1` + gofmt + `go vet` clean; 174 Python green; all 3 CLIs shellcheck-clean; py_compile clean; flake **0** (5/5). Retro: [m14-unified-skills/retro.md](releases/archive/01.30-stack-party/m14-unified-skills/retro.md). **Next:** M15 (safety & security doc + dual-repo KB refresh — the LAST v1.3 milestone).
**Goal:** One coherent stack-operations skill set that works on any stack (`dev-N | demo-N`), with the dev lifecycle
mirroring demo's.
**Scope:** In: **`dev-up`** (consolidate `setup-platform` + `start-platform` into one dev bring-up that drives the M13
flow) + **`dev-down`**; **hard-rename** (no aliases, user 2026-06-07) the operation skills to generic, stack-target
forms — **`stack-list`** (←`demo-status`), **`stack-seed`** (←`demo-seed`), **`stack-snapshot`** (←`demo-snapshot`),
**`stack-update`** (←`update-platform`) — each detecting/accepting a `dev-N|demo-N` target; **remove** the old skill
dirs (`setup-platform`, `start-platform`, `update-platform`, `demo-status`, `demo-seed`, `demo-snapshot`); update
**every reference** (CLAUDE.md skill table, READMEs, the corpus skill docs + recipes). `demo-up`/`demo-down` stay as
the demo lifecycle (aligned with `dev-up`/`dev-down`). Out: the safety doc (M15).
**Depends on:** **M12 + M13** (the generic skills drive the registry + the dev peer capabilities). **Parallel with:** none.
**Open questions:** how much of `setup-platform`'s first-time machine setup (tool install, org clone) folds into `dev-up` vs stays a one-time prerequisite; the exact target-detection UX (`stack-seed dev-1` vs `stack-seed --stack dev-1`).
**KB dependencies:** the existing skill SKILL.md files, `corpus/ops/*` guides (setup/run/update/demo), `CLAUDE.md` (the skill table).
**Delivers → the unified `stack-*` + `dev-up`/`dev-down` skills + a rewritten `CLAUDE.md` skill table + refreshed `corpus/ops/` guides** (the converged stack model, hard-renamed).

### M15: Safety & security doc + dual-repo knowledge consolidation
**Status:** `done` (completed 2026-06-07) · **Shape:** `section` · **Complexity:** medium · **Dir:** [m15-safety-doc/](releases/archive/01.30-stack-party/m15-safety-doc/)
**Closed 2026-06-07** (build → 4 harden passes → close review → merged to `release/01.30-stack-party`). **The closing milestone of v1.3 — the LAST of the four; with it done the release is ready for `/developer-kit:close-release`.** A documentation/consolidation milestone (the M8/M11-analog). Delivered the net-new **`corpus/ops/safety.md`** (248 lines, code-cited, dual-level): the two inviolable guarantees of the `rosetta-extensions` tooling — **never reads private/customer data** (read-side: the tenant firewall `AssertPlan`/`AssertCaptured` [the conceptual `AssertPublicOnly` named to its two real Go gates, M15-D2], the per-surface public predicates byte-matched to `firewall.go` [`organization_id IS NULL` / `private = false AND tenant_id IS NULL AND status = 'published'`], the public-only data-DNA gene, bounded read-only capture [`SET TRANSACTION READ ONLY` + timeouts, **no offline file reader** — M15-D3 keeps the M9b-D9 drop true]) and **never touches prod data or services** (write-side: the 3-layer isolation guard `CheckWrite`/`PreflightEnv`/`AssertClean`, never-write shared Directus/prod-S3, the capture-source policy, the **doubled** n=0-dev guards [scoped precisely — `stacksnap` replay has none, correctly, M15-D4], the audit-proven zero-pollution assertion). Cross-linked from all 4 siblings (`db-access`/`snapshot-spec`/`seeding-spec`/`security_compliance`, both directions). Plus a **dual-repo KB refresh**: `rosetta-extensions/knowledge/` to the v1.3 converged dev≡demo model + safety contract (0 stale pre-M14 skill names); root READMEs + `CLAUDE.md` + the `demo/` recipe family for the unified `stack-*` skills + dev-as-peer + safety.md discoverability. **Harden** (4 passes) made doc accuracy a *test*: 7 fail-closed docs↔code drift guards pin every load-bearing literal/symbol/SQL-block safety.md quotes to the real code (read-side predicates + gate names + the bounded-read SQL `source.DefaultBounds().SetupSQL()`; write-side the complete `realClerkHosts` + `directusTokenKeys` rejection lists + the forced `STORAGE_S3_PUBLIC_BUCKET=""` override + the `CheckWrite`/`PreflightEnv`/`AssertClean` symbols), each `t.Skip`-ping gracefully in pinned-tag consumption clones; harden also landed the **M15-D4** n=0 over-claim fix (corrected in both `dev-setdress.sh` + the sibling test comment to match the shipped safety.md §2.5). Close found **1 finding** (a self-referential docs-accuracy drift: `decisions.md` M15-D4's text lagged the harden fix — corrected; 0 code/test/scope) and proved the guards fail-closed at close (Phase 2c — mutating the read-side predicate + the write-side bucket override each tripped its guard; safety.md restored byte-identical). Deferral re-audit **GREEN**, run release-wide M12→M15 as the terminal-milestone pre-release sweep (1 inherited DEF-M10-01 → v1.4 signed/unchanged/**not aged out**; 0 repeat/aged/new — M15 added **zero** deferrals, all Fate-1). Decisions M15-D1…D4 + Q1 resolved. Extensions: tag **`stack-party-m15`** @ `51ca18b`. Go test funcs **+7** (the 7 drift guards: stack-snapshot +3, stack-seeding +4); 174 Python (unchanged); all 4 Go modules `-race -count=1` + gofmt + `go vet` clean; `dev-setdress.sh` shellcheck-clean; py_compile clean; **flake 0** (5/5 both touched packages). Retro: [m15-safety-doc/retro.md](releases/archive/01.30-stack-party/m15-safety-doc/retro.md). **Next:** **`/developer-kit:close-release`** (merge `release/01.30-stack-party` → `main`, tag `v1.3`) — the user's separate step.
**Goal:** A single authoritative doc on **how the rosetta tooling stays safe** — it never reads private/customer data,
and it never touches production data or services — plus a refresh of **both** knowledge bases to the v1.3 model.
**Scope:** In: a new **`corpus/ops/safety.md`** consolidating (a) the **read-side** (private-data avoidance: the
tenant firewall `AssertPublicOnly`, the `organization_id IS NULL` / `private = false AND tenant_id IS NULL` public
predicates, the public-only data-DNA gene) and (b) the **write-side** (prod-protection: the 3-layer isolation guard,
read-only capture, never-write-shared-Directus / prod-S3, the capture-source policy, the n=0-dev guards); cross-link
from `snapshot-spec`/`seeding-spec`/`db-access`/`security_compliance`; **update the `rosetta-extensions/knowledge/`**
(the repo's own KB) for the converged stack model + the safety contract; refresh the root READMEs + the `demo/`
recipe family for the unified `stack-*` skills + dev-as-peer. Out: nothing (closing milestone).
**Depends on:** **M12 + M13 + M14** (documents their converged result). **Parallel with:** none (the closing milestone before `/developer-kit:close-release`).
**Open questions:** doc home — `corpus/ops/safety.md` vs `corpus/architecture/tooling-safety.md` (lean: `corpus/ops/safety.md`, ops-adjacent to seeding/snapshot/db-access).
**KB dependencies:** `corpus/ops/{snapshot-spec,seeding-spec,db-access}.md`, `corpus/architecture/security_compliance.md`, the `rosetta-extensions/knowledge/` base.
**Delivers → `corpus/ops/safety.md`** (net-new — the tooling's read-side + write-side safety contract) **+ updates the `rosetta-extensions/knowledge/` base** to the v1.3 converged model.

### Execution graph (v1.3)
```
v1.3 "stack party" — dev + demo stacks become first-class peers, one unified skill set
   M12 (unified registry + first-free-N) ─→ M13 (dev peers: local Directus + auto-snapshot + light seed) ─→ M14 (unified stack-* skills) ─→ M15 (safety doc + dual-repo KB)
```
**Sequential.** M12 lays the shared registry/port foundation; M13 makes dev a full peer (the meatiest — reuses M10's
per-stack Directus + the snapshot/seed framework); M14 converges the skills onto both stack kinds (hard rename); M15
consolidates the safety story + both knowledge bases. M12∥M13 is feasible (stack-core vs dev bring-up) but
serializing keeps M13's bring-up consuming the new registry cleanly. No B-milestones — M14 *is* the tooling layer.

### Risks (v1.3)
- **(M13, scope)** spawning a per-stack Directus + auto-snapshot + seed on **every dev build** could make dev
  bring-up slow/heavy. Mitigate: reuse M10's proven `provision.go`; cache-first snapshot (replay, not capture);
  minimal `dev-min` seed; `--no-snapshot` escape.
- **(M12, correctness)** the first-available-N allocator must be **race-safe** (two concurrent `up`s) and reconcile
  with reality. Mitigate: locked registry write + `docker ps` as the used-N source of truth.
- **(M14, blast radius)** the **hard rename** (no aliases) breaks any external reference to `setup-platform` /
  `demo-seed` / etc. Mitigate: update **every** in-repo reference in M14; the user accepted the clean-break tradeoff.
- **(cross-cut)** dev was historically the *protected* environment (real Clerk, the n=0 guard). Making it a
  snapshot/seed target must preserve those guards. Mitigate: M15's safety doc + the kept n=0-reset guard + read-only capture.

### Open decisions (resolve during build)
Registry-of-record vs `docker ps`-derived — M12 (lean registry + reconcile); dev auto-snapshot default-on vs opt-in
— M13 (lean default-on, `--no-snapshot`); how much first-time setup folds into `dev-up` — M14; the safety doc home
(`corpus/ops/safety.md`) — M15.

## Done — v1.2 "set dressing" (SHIPPED 2026-06-07 · tag `v1.2`)

**Theme:** v1.1 made demo stacks *structurally* populated (orgs, users, backdated activity) but consciously
**waived two surfaces** — `taxonomy` (the 60K-skill / 18K-role node hierarchy + embeddings) and `content` (the
shared Directus content library) — because they can't be *fabricated* structurally: taxonomy is reference data,
content lives in a shared store the isolation guard blocks. v1.2 lifts both via a **snapshot mechanism** —
capture the *real* surface once from a source, replay it per-stack — taking data-DNA coverage from "8 reachable
surfaces" to **100% of the full catalog**. The elevation: snapshot replay is itself an **alignment problem**
(does the replayed surface faithfully reproduce the captured source?), so v1.2 extends the M0/M7b
alignment+data-DNA discipline with a **snapshot-fidelity dimension** rather than shipping plain plumbing. Demo
worlds become *set-dressed*: the stage (v1.1 "show floor") gets its believable props.

> **Designed 2026-06-05** from the v1.1 carry-forward (the user-confirmed M7c waiver → roadmap-vision seed).
> Phase 0a deferral audit **GREEN** (6 carry-forwards, all single, clear v1.2/vision homes — no repeats/chronic;
> [`.agentspace/scratch/deferral-audit-2026-06-05.md`](../../.agentspace/scratch/deferral-audit-2026-06-05.md)).
> Phase 0b KB blind-area check: snapshot/AI-content/deploy-CI **YELLOW** (anchored, need a spec doc as a
> `Delivers →`); shareability **RED** (blind area — deferred to v1.3, not in v1.2). 3 research agents over the
> seeding stack, the skiller taxonomy + Directus content surfaces, and milestone history — verified against the
> clones in `stack-dev/` (platform) + `stack-demo/rosetta-extensions/` (the seeding stack). Gap analysis:
> [`.agentspace/scratch/roadmap-research-2026-06-05.md`](../../.agentspace/scratch/roadmap-research-2026-06-05.md).
>
> **Scope decided (user, 2026-06-05):** **snapshot spine only** — `taxonomy` + `content` to 100% coverage.
> **AI-generated rich content** (transcripts / AI-scored narratives / fresh embeddings) and **external
> shareability** (Tailscale vs ingress) are confirmed **v1.3** (kept in `roadmap-vision.md`), so v1.2 stays the
> tight, well-grounded release the snapshot work warrants. Codename **"set dressing"** continues the stage
> metaphor (body double → show floor → set dressing).
>
> **Refined 2026-06-06** (user — 5 notes on M9+) with **live production read access** (the wired
> `mcp__postgres__query` tool, `marco_read` over Tailscale; catalog-only queries — no GB scans). Changes: (1)
> snapshotting becomes a **dedicated reusable `rosetta-extensions/stack-snapshot/` section** (capture-read
> decoupled from seeding-write); (2) a **production-safe capture-source policy** — source-pluggable, default
> **ingest an existing prod `pg_dump`** → fallback **safe throttled primary read** (MVCC = read-only never blocks
> writers), with **restore-from-snapshot / read replica** as zero-impact upgrades; (3) a tested **tenant-data
> firewall** (`AssertPublicOnly`) — capture *only* `organization_id IS NULL`
> reference data, never customer rows; (4) snapshots live in a **`.agentspace` manifest-cached, pluggable store**
> (no GB blobs in git; cloud/S3 = v1.3); (5) the **`/db-query` skill is ported** into rosetta as the prod-read
> foundation. Former **M9 split → M9a (framework) + M9b (taxonomy surface)** (the M7a→M7c precedent). Prod
> findings (skiller ≈ 2.1 GB, taxonomy ~98% public, app-Postgres `cms.studio_*` = 100% customer → excluded, the
> public content lives in a *separate* Directus store): [`.agentspace/scratch/roadmap-research-2026-06-06.md`](../../.agentspace/scratch/roadmap-research-2026-06-06.md).

### The snapshot surfaces (prod-grounded 2026-06-06 — public-only, and why each needs a different mechanism)
- **`taxonomy`** — lives in the **per-stack** Postgres `skiller` schema but ships *empty* (normally loaded via the
  `importskills`/`importjobroles` cobras). Prod-measured: **~2.1 GB**, and **~98% public** — `skills` 42,763
  public / 794 private, `job_roles` 22,315 / 2,381. So the **tenant firewall** captures `organization_id IS NULL`
  (keeping the full public catalog, dropping the customer tail automatically); embeddings + translations carry no
  org column → scoped via the **public parent**. The snapshot is **Postgres→Postgres**: capture-once from a safe
  source, **bulk-`COPY` replay** per-stack (the M7a perf path). One refinement: `skill_embeddings` is 692 MB but
  the heap is 3.3 MB — ~689 MB is the **pgvector index**, so capture vectors verbatim and **rebuild the index on
  replay**. The *cleaner* surface — it proves the framework (M9b). (`data-dna.json`: `taxonomy` status `waived-m7c`.)
- **`content`** — the **public** simulation / skill-path **template library**. Prod correction: it is **not** the
  app-Postgres `cms` schema — `cms.studio_documents` + `cms.studio_tasks` are **100% org-scoped customer data (0
  public rows)** → **excluded** by the firewall. The public library lives in the **shared self-hosted Directus
  store** (`content.anthropos.work`) — *(M10 build later found this is the `directus` schema of the **same** prod
  Postgres, read-only via `marco_read`, not separate infra — see the M10 close note)*; the isolation guard hard-blocks writes to shared Directus,
  so replay needs a **per-stack content store**. The defining M10 decision (resolve early): per-stack Directus
  container fed from the captured snapshot vs replay straight into the per-stack Directus Postgres DB (Directus's
  own backing store is Postgres → stays in the per-stack-isolated class). The *highest-risk* surface in v1.2.
  (`data-dna.json`: `content` status `waived-m7c`.)

### M9a: Snapshot extension — capture-safe, public-only, manifest-cached framework + `/db-query` port
**Status:** `done` (2026-06-06) · **Shape:** `section` · **Complexity:** large · **Dir:** [m9a-snapshot-framework/](releases/archive/01.20-set-dressing/m9a-snapshot-framework/)
**Closed 2026-06-06** (build → harden 2 passes → close review → merged to `release/01.20-set-dressing`). Delivered the dedicated **`rosetta-extensions/stack-snapshot/`** section (9 Go pkgs + the `stacksnap` CLI, tagged `stack-snapshot-m9a` @ `1cc4dd2`): the capture/serialize/replay contract + portable `manifest.json` format; the **production-safe capture-source policy** (M9a-D3: cache-hit → dump-ingest [default] → safe primary read [MVCC] → restore/replica [AWS-gated upgrades]) with a bounded read-only session + catalog-first dry-run; the **tenant-data firewall** `AssertPublicOnly` (plan-time + post-capture, hard-fail on any tenant row, in-memory stash so nothing persists on a leak); the **`.agentspace` manifest-cached pluggable `SnapshotStore`** (localfs now, cloud/S3 = v1.3); the **data-DNA snapshot dimension** (`stack-seeding/dna/snapshot.go` — `snapshot-seeded` status that counts toward coverage + 5 two-sided fidelity operators) + the `/db-query` port. Proven end-to-end on the hermetic `reference-toy` surface. **556 Go test funcs (+147; stack-snapshot 128 new, stack-seeding 145→164)**, flake **0** (5× shuffled, `-race`), gofmt+vet clean. Close: scope/code/docs/tests **GREEN**; 1 finding fixed (the tag trailed HEAD by the harden commits → re-pointed + force-pushed); 5 adversarial scenarios recorded; deferral audit GREEN. Decisions M9a-D1…D7 + Q2/Q3/Q4. Retro: [m9a-snapshot-framework/retro.md](releases/archive/01.20-set-dressing/m9a-snapshot-framework/retro.md). **Next:** M9b (the real ~2.1 GB taxonomy surface).
**Goal:** A **dedicated, reusable `rosetta-extensions/stack-snapshot/` section** that captures a *public* reference
surface once from a **safe, low-impact source** (default a prod `pg_dump`), serializes it to a `.agentspace` manifest-cached store, and
replays it per-stack — with a tested **tenant-data firewall** (never customer data) + an **alignment extension that
measures replay fidelity**. Proven on a tiny reference surface (M0 toy-mirror discipline); ports the **`/db-query`**
skill as the prod-read foundation.
**Scope:**
  - In: **(note #1)** the dedicated **`stack-snapshot/`** section (capture + serialize + store + replay + the
    `stacksnap` CLI) — capture (a privileged prod **read**) decoupled from seeding (per-stack **writes**); the
    **snapshot contract + portable format** (per-table `COPY` payloads + `manifest.json`); **(note #2)** the
    **production-safe capture-source policy** (source-pluggable) — cache-hit (no prod read) → **(1) prod-`pg_dump`
    ingest** [default, zero new load] → **(2) safe throttled primary read** [MVCC = no write blocking] → **(3)
    restore-from-snapshot / (4) read replica** [zero-impact upgrades once eu-west-1 AWS/infra is wired], with
    bounded read-only sessions + a catalog-first dry-run (the **read half** the M7a guard lacks); **(note #3)** the **tenant-data firewall** `AssertPublicOnly`
    + a **public-only/provenance gene** (hard-fail on any captured tenant row); **(note #4)** the **`.agentspace`
    manifest-cached, pluggable `SnapshotStore`** (localfs now, cloud/S3 = v1.3; no GB blobs in git); the **data-DNA
    extension** — `snapshot-seeded` status + a **snapshot-fidelity gene class** (row-count / structural conformance
    / referential integrity / **embedding-dimension integrity**); **(note #5)** the **ported `/db-query` skill** +
    `corpus/ops/db-access.md` (the MCP-tool **and** pgpass/psql paths); a tiny reference surface proving
    capture→store→replay→fidelity-gate end-to-end.
  - Out: the taxonomy surface (M9b); the Directus content snapshot (M10); recipes/presets (M11); the cloud store
    backend + AI-generated content + shareability (v1.3).
**Depends on:** v1.1's M7a (isolation guard + perf path) + M7b (the data-DNA harness it extends). **Parallel with:** none (gates M9b + M10 + M11).
**Open questions:** capture source resolved (M9a-D3, investigated 2026-06-06 — no replica + no local AWS creds →
default **prod-`pg_dump` ingest**, fallback **safe primary read**; restore-from-snapshot/replica = later upgrades);
the manifest schema + cache-staleness rule; embedding capture (vectors verbatim, **rebuild the ~689 MB pgvector
index on replay**); the `SnapshotStore` interface so the v1.3 cloud swap is a backend change.
**KB dependencies:** `corpus/ops/seeding-spec.md` (the framework + isolation boundary + the DAG node), `corpus/architecture/alignment_testing.md` (the data dimension to extend), `corpus/ops/staging_from_dump.md` (the full-dump **anti-pattern** to contrast), the source `db-query` skill (ported).
**Delivers → `corpus/ops/snapshot-spec.md`** (net-new — the extension + capture/replay contract + capture-source policy + tenant firewall + the `.agentspace` manifest store) **+ `corpus/ops/db-access.md`** (net-new) **+ the `/db-query` skill** **+ extends `corpus/architecture/alignment_testing.md`** (the snapshot-fidelity + public-only genes).

### M9b: Taxonomy snapshot (the first real surface)
**Status:** `done` (completed 2026-06-06) · **Shape:** `section` · **Complexity:** large · **Dir:** [m9b-taxonomy-snapshot/](releases/archive/01.20-set-dressing/m9b-taxonomy-snapshot/)
**Closed 2026-06-06** (build → 3 harden passes → close attempt 1 BLOCKED at a RED deferral audit → user fate decision → close attempt 2 clean → merged to `release/01.20-set-dressing`). Delivered: the **10-table public taxonomy surface** in FK replay order (`categories`, `job_role_categories` → `specializations` → `skills`, `job_roles` → embeddings/translations/`job_role_skills`), `org_id IS NULL` filters + **parent-scoped predicates** for column-less tables (`TableSpec.PublicViaFK` + `firewall.ParentScopeFilter` — closing M9a's empty-filter gap, M9b-D2; `job_role_skills` both-endpoints, M9b-D3), **pgvector index rebuilt via `REINDEX` on replay** (vectors verbatim, dim 1536 in the manifest, M9b-D5), the **manifest→`CapturedTable` fidelity bridge** + `datadna measure-snapshot` CLI (M9b-D6), and the **`TaxonomySnapshotSeeder` DAG node** (`… → taxonomy → activity`, M9b-D7). Coverage `waived-m7c → snapshot-seeded-m9b` (all 5 fidelity operators). PR-review fix: the parent-leak probe must AND the capture filter (was scanning the whole table → false abort). **Deferral fate (M9b-D9):** the **offline pg_dump-FILE reader was DROPPED by the user** (an M9a→M9b repeat/aged-out deferral) — restore-then-`--dsn` + the safe primary read cover the need; the dead `--dump` flag was removed and docs+code corrected (dump-ingest = restore-the-dump-into-Postgres-then-`--dsn`, no offline reader, not seeded to v1.3). Decisions M9b-D1…D9. Extensions: tag `stack-snapshot-m9b` @ `55ee0e6`. Go test funcs 556→**635** (stack-snapshot 128→167, stack-seeding 164→204); flake 0; both modules `-race` green. Retro: [m9b-taxonomy-snapshot/retro.md](releases/archive/01.20-set-dressing/m9b-taxonomy-snapshot/retro.md).
**Goal:** Prove the M9a framework on the **real ~2.1 GB taxonomy surface** — capture the *public* skiller catalog
from a safe source, bulk-`COPY` replay per-stack, **rebuild the pgvector index on replay**, fidelity- + public-only
gated — driving data-DNA coverage from `waived` to its first `snapshot-seeded` surface.
**Scope:**
  - In: **public taxonomy capture** — `skiller.{categories,specializations,skills,job_roles}` filtered
    `organization_id IS NULL` (full public catalog, customer tail dropped); `{skill,job_role}_embeddings` (vectors
    only) + `{skill,job_role}_translations` + `job_role_skills` via the **public-parent** join; **bulk-`COPY`
    replay** per-stack (M7a perf path, per-stack-isolated only); **pgvector index rebuild on replay** (carry
    vectors verbatim, don't transport the ~689 MB index); the **taxonomy fidelity + public-only genes**;
    wiring into the `stack-seeding` DAG node. Coverage `waived → taxonomy-seeded`.
  - Out: the Directus content surface (M10); recipes/presets (M11); recompute of embeddings (v1.3).
**Depends on:** **M9a** (the `stack-snapshot` extension + capture-source policy + firewall + store + fidelity genes). **Parallel with:** none (gates M10 + M11).
**Open questions:** keyset-chunked vs single streamed `COPY` for skills/embeddings (size via the dry-run); the
pgvector index params to rebuild (match prod, record in the manifest); `job_role_skills` referential integrity vs
the public-skill set.
**KB dependencies:** `corpus/ops/snapshot-spec.md` (M9a's contract), `corpus/services/skiller.md` (the taxonomy schema + embeddings + translations), `corpus/ops/seeding-spec.md` (the perf path + the DAG node), `corpus/architecture/alignment_testing.md` (the genes).
**Delivers → extends `corpus/ops/snapshot-spec.md`** (the taxonomy capture/replay path) **+ updates `corpus/ops/seeding-spec.md`** (taxonomy promoted `waived` → `snapshot-seeded`).

### M10: Directus content snapshot-replay
**Status:** `done` (completed 2026-06-06) · **Shape:** `section` · **Complexity:** large (highest-risk) · **Dir:** [m10-content-snapshot/](releases/archive/01.20-set-dressing/m10-content-snapshot/)
**Closed 2026-06-06** (spike → build → 5 harden passes → close clean → merged to `release/01.20-set-dressing`). Delivered the **second real surface**: the **9-table public Directus content** library (`simulations` + `skill_paths` scope-bearing roots → `resource` pure-reference → `roles`/`sim_tasks`/`sequences` parent-scoped → `task_checks`/`task_sub_checks` **multi-level chains** → `sequences_roles` via BOTH). Generalized the firewall to a **per-surface `PublicPredicate`** (the spike-flagged org-only gap, M10-D1; taxonomy byte-for-byte unchanged via the zero-value→`DefaultPredicate` fallback); the directus predicate `private=false AND tenant_id IS NULL AND status='published'` (M10-D3, prod-verified); **multi-level parent chains** via `ParentScope.ParentFilter` chasing column-less intermediates to the scope-bearing root in one subquery (M10-D4); the **per-stack Directus store fork** = bootstrap→replay→boot on the per-stack Directus-backing Postgres `directus` schema (`PerStackIsolated`, M10-D2), `EnvContract.Validate` hard-rejecting any prod-Directus target; media **REFS** landed (ref columns + `ReferencedFilesFilter` + the 1,311 public `directus_files` rows + a local-storage/placeholder adapter) with blob **BYTES** S3-gated → **Fate 2 → v1.3** (M10-D5, confirmed-covered by the roadmap-vision cloud-store seed); the **content fidelity + public-only genes** (`ReplayedNonPublicRows`); the **`ContentSnapshotSeeder`** DAG node + the **`sim_id`/`skill_path_id`/`resource_id` linkage resolver** (v1.1 free-value refs now resolve against real replayed public templates; free-value fallback when no snapshot). Capture source **self-resolved**: the `directus` schema lives in the SAME app Postgres, read-only via `marco_read` (the spike's "separate store" premise disproven, KB-1 corrected). Coverage `content` **`waived-m7c → snapshot-seeded-m10`** → **NOTHING left waived → 100% over the full catalog (the v1.2 set-dressing thesis complete)**. Close found **2 findings** (both fixed): the stack-snapshot README package index was missing the `taxonomy`+`directus` rows (per-unit handbook contract — a latent M9b miss caught here), and the M10-D5 decision record. Decisions M10-D1…D5. Extensions: tag `stack-snapshot-m10` @ `df410f5`. Go test funcs 635→**701** (stack-snapshot 167→207, stack-seeding 204→230); flake 0 (5/5 both modules); both modules `-race` green; gofmt+vet clean. Retro: [m10-content-snapshot/retro.md](releases/archive/01.20-set-dressing/m10-content-snapshot/retro.md).
**Goal:** Capture the shared-Directus content library and replay it into a **per-stack content store** — never
touching shared Directus — taking data-DNA coverage to **100% of the full catalog** (the last `waived` surface
promoted to `snapshot-seeded` + fidelity-gated).
**Source correction (prod, 2026-06-06):** the content source is **not** app-Postgres `cms` — `cms.studio_documents`
+ `cms.studio_tasks` are **100% org-scoped customer data (0 public rows)** → **excluded** by the firewall. The
public template library lives in the **separate self-hosted Directus store**; M10 captures **only its public/global
templates**.
**Scope:**
  - In: the **per-stack content-store decision** resolved + built (per-stack Directus container vs direct
    per-stack Directus-Postgres replay — the defining fork, resolve in the first iter/spike); the **public content
    capture** (export the public/global Directus templates + media references from the separate Directus store — a
    privileged read via M9a's capture-source policy + tenant firewall, isolation-clean); the **content replay
    seeder** wired into M9a's snapshot framework + the seeder DAG (M7a), respecting the guard; the **content
    fidelity + public-only genes** in the data-DNA; the **`sim_id`/`skill_path_id`/`resource_id` linkage** so the
    v1.1 session/assignment seeders' content refs resolve against the real **public** templates (closing the
    "free-value refs" gap).
  - Out: app-Postgres `cms.studio_*` customer content (excluded — tenant data); AI-generated/authored content
    (v1.3 — this replays *real* captured public content, it does not generate); recipes/presets (M11); shareability (v1.3).
**Depends on:** **M9a + M9b** (the snapshot framework + fidelity DNA + the `stacksnap` CLI + the proven taxonomy surface). **Parallel with:** none (M11 curates its output).
**Open questions:** the content-store fork (above) — the load-bearing decision; where the public Directus templates
physically live + how the public/global subset is identified (confirm against the Directus store); whether
media/blobs are in-scope or refs-only for the demo MVP (S3-private is per-stack-isolated, so blobs *can* be
replayed — confirm at build); how much of the collection set the demo needs (the believable subset, per the M7c "reachable" discipline).
**KB dependencies:** `corpus/ops/snapshot-spec.md` (M9a's contract), `corpus/services/cms.md` (Directus) + `corpus/ops/db-access.md` (the Directus store connection), `corpus/ops/seeding-spec.md` (the isolation guard + the session/assignment content refs), `corpus/services/{jobsimulation,skillpath}.md` (the consumers of `sim_id`/`skill_path_id`).
**Delivers → extends `corpus/ops/snapshot-spec.md`** (the public-Directus content path + the store decision) **+ updates `corpus/ops/seeding-spec.md`** (content surface promoted from `waived` to `snapshot-seeded`).

### M11: Richer-world recipes, presets + corpus polish
**Status:** `done` (completed 2026-06-06) · **Shape:** `section` · **Complexity:** medium · **Dir:** [m11-richer-recipes/](releases/archive/01.20-set-dressing/m11-richer-recipes/)
**Closed 2026-06-06** (build → 1 harden pass → close review → merged to `release/01.20-set-dressing`). **The LAST milestone of v1.2** — the product/discoverability layer (the M8-analog) that curates the 100%-coverage full-fidelity world (M9a/M9b/M10) into usable recipes/presets, with **zero new production code path**. Delivered: the 3 **seed presets** (`small-200`/`mid-500`/`large-1k`) gained a documented FULL-FIDELITY PREREQUISITE comment-header (replay `taxonomy` + `directus` BEFORE seeding; graceful structural-only degradation without) — **presets stay purely structural** (snapshots are stack-global reference data, not org-scoped → no schema field, M11-D3); the **set-dressed `corpus/ops/demo/` recipe family** (both org-onboarding + skill-progression recipes gained a `/demo-snapshot` replay step + call out the real catalog/templates; the FALSE "waived/future-v1.2" note in `recipe-skill-progression.md` rewritten to the shipped 100%-coverage reality) + a **new `recipe-snapshot-world.md`** (the capture→replay→set-dressed-world curator walk-through); the **`/demo-snapshot` skill** (M11-D1, resolves M11-Q2 — a NEW skill, not a `/demo-seed` extension: capture = privileged prod READ vs seeding = per-stack WRITE; `replay` headline / `capture` rare-maintenance / `status`); corpus cross-links (CLAUDE.md skill-table row + key-docs entry; `snapshot-spec` ↔ `seeding-spec` ↔ the demo family bidirectional; data-DNA reads 100% throughout); the **§5 release-close hygiene carry** — the stale `stacksnap --help` text fixed (M11-D4: framework tag M9a/M9b → M9a/M9b/M10 + the registered-but-unlisted `directus` surface now listed). Decisions M11-D1…D4 + Q1/Q2 resolved. Extensions: tag `stack-snapshot-m11` @ `1e18df6`. Go test funcs 701→**708** (stack-snapshot 207→212: the `--help` contract pins + docs↔parser flag-drift guard; stack-seeding 230→232: shipped-preset strict-parse/validate + size-order); flake **0** (5/5 both M11-touched packages); both modules `-race` + gofmt + `go vet` clean. Close: deferral re-audit **GREEN** (1 inherited DEF-M10-01 unchanged, 0 repeat/aged/new — M11 added zero deferrals, all Fate-1); scope/code/docs/tests **GREEN** with **0 findings**; 4 adversarial scenarios recorded (each mutation-pinned). Retro: [m11-richer-recipes/retro.md](releases/archive/01.20-set-dressing/m11-richer-recipes/retro.md). **v1.2 is now complete → next: `/developer-kit:close-release`** (merge `release/01.20-set-dressing` → `main`, tag `v1.2`).
**Goal:** The product/discoverability layer that closes v1.2 (the M8-analog): make the full-fidelity worlds
*usable + discoverable* — refresh presets + recipes so a demo curator gets a real-taxonomy, real-content world
out of the box, and update the corpus to reflect 100% coverage.
**Scope:**
  - In: refreshed **seed presets** (small/mid/large) that now include the taxonomy + content snapshots; an updated
    **`corpus/ops/demo/` recipe family** (the end-to-end recipes now showcase a *set-dressed* world — real skills
    in the catalog, real simulations/skill-paths behind the seeded sessions); a **`/demo-snapshot` (or extended
    `/demo-seed`) skill** driving the `stacksnap` CLI; cross-linking + corpus updates (the data-DNA now reads 100%);
    the **release-close hygiene** carry (any small items surfaced in M9a/M9b/M10).
  - Out: new snapshot surfaces (M9a/M9b/M10 own them); AI-content + shareability (v1.3).
**Depends on:** **M9a + M9b + M10** (curates their output). **Parallel with:** none (the closing milestone before `/developer-kit:close-release`).
**Open questions:** whether snapshot capture is a curator step or a manifest-cached refresh (decide with M9a's capture-source policy); `/demo-seed` extension vs a new `/demo-snapshot` skill.
**KB dependencies:** `corpus/ops/demo/README.md` + the recipes, `corpus/ops/seeding-spec.md`, `corpus/ops/snapshot-spec.md`, the `/demo-seed` skill.
**Delivers → refreshes `corpus/ops/demo/`** (recipes + presets to full-fidelity) **+ the `/demo-snapshot` skill + the CLAUDE.md skill table.**

### Execution graph (v1.2)
```
v1.2 "set dressing" — richer demo worlds: the real *public* taxonomy + content, measured-faithful, to 100% coverage
   M9a (stack-snapshot framework: capture-safety + tenant firewall + .agentspace store + /db-query + fidelity-DNA)
        └─→ M9b (taxonomy surface: public skiller + embeddings, rebuild index) ─→ M10 (public Directus content) ─→ M11 (recipes + presets + corpus)
```
**Sequential.** M9a lands the **dedicated `stack-snapshot` extension** + the capture-source policy + the tenant
firewall + the `.agentspace` manifest store + the fidelity-DNA + the `/db-query` port, proven on a toy surface.
M9b proves it on the cleaner ~2.1 GB taxonomy (coverage waived→taxonomy-seeded). M10 takes the harder public
Directus content surface to **100% coverage**. M11 curates the full-fidelity worlds into usable recipes/presets +
closes the release. No parallel tracks — one extension + one data-DNA; serializing keeps the merge surface clean
(the v1.1 spine discipline).

### Risks (v1.2)
- **(M9a, blocks-prod-safety)** a capture that **reads the hot primary** under load, or **leaks a tenant row**.
  Mitigate: the capture-source policy (cache-hit → prod-`pg_dump` ingest → safe throttled primary read [MVCC = no
  write blocking] → restore-from-snapshot/replica upgrades) + bounded read-only sessions; the **tenant firewall**
  `AssertPublicOnly` + public-only gene — tested gates, not conventions.
- **(M10, blocks-100%-coverage)** the **content-store fork** + locating the public Directus template subset (a
  *separate* store). Mitigate: resolve the store decision in an M10 spike *first* (Directus's backing store *is*
  Postgres → a per-stack Directus-Postgres replay stays in the isolated class); fall back to refs-only believability
  if a full Directus stand-up proves too heavy for the demo MVP.
- **(M9b, scope)** **embedding fidelity** — carrying pgvector embeddings verbatim (dimension + value integrity) and
  **rebuilding the ~689 MB index on replay** rather than transporting it. Mitigate: capture vectors verbatim
  (offline + deterministic), gate the embedding-dimension gene; never recompute (that's AI-content, v1.3).
- **(note #4 → v1.3)** the local `.agentspace` cache doesn't share across machines / scale. Mitigate: the
  `SnapshotStore` interface keeps the **cloud/S3 swap** a v1.3 backend change (the manifest already addresses by location).
- **(cross-cut)** **the isolation guard's missing read half** — capture reads a reference source. Mitigate: capture
  is read-only + audited (the capture-source policy); replay writes only to per-stack-isolated stores, the existing
  3-layer guard asserts clean (extend `AssertClean` to cover snapshot replay).

### Open decisions (resolve during build)
The **capture source** — **resolved** M9a-D3 (user 2026-06-06): default **prod-`pg_dump` ingest** → fallback **safe
throttled primary read**; restore-from-snapshot/replica = zero-impact upgrades once eu-west-1 AWS/infra is wired;
the manifest schema + cache-staleness rule — M9a; embedding capture
(verbatim + rebuild-index-on-replay) — M9b (lean verbatim); the **per-stack content-store fork** (per-stack
Directus container vs direct Directus-Postgres replay) — M10, the defining decision; identifying the public
Directus template subset — M10; media/blobs in-scope vs refs-only — M10; `/demo-seed` extension vs a new
`/demo-snapshot` skill — M11.

## Done — v1.1 "show floor" (SHIPPED 2026-06-05 · tag `v1.1`)

**Theme (broadened 2026-06-04):** v1.0 made the platform run *without* Clerk; v1.1 started as "disposable
demo stacks" (M3 ✅) and now becomes **the platform-operations extension framework** — consolidate the repo
constellation into **two repos** (`rosetta` = the platform corpus + dev-env skills; `rosetta-extensions` = a
monorepo of operations sections), then deliver the seeded-demo capability *and* generalize the pattern to dev.
Everything stays **additive — zero change to any read-only platform repo**.

**Refactored 2026-06-04** (after M3 shipped, to keep the constellation from exploding): the standalone
`clerkenstein` + `rosetta-demo` repos collapse into `rosetta-extensions/{clerkenstein,demo-stack,…}`; the
former M4 (seeding) → **M7**, former M5 (recipes) → **M8**; new structural milestones M4–M6 inserted. Decisions:
**git subtree, history-preserving** (M4-D1) · **delete the old repos, not archive** (M4-D2, user) · **the
alignment framework stays in rosetta** (M4-D3) · per-demo clones (M3-D1) · clone-at-release-tag (M3-D3).

**Seeding redesigned 2026-06-04** (M3–M6 all shipped): the user asked to make seeding robust/resilient/drift-proof/
fast/**production-safe**, so the single `section` M7 splits into **M7a → M7b → M7c** (a section + section +
iterative "mix"). 3 research agents over the platform grounded it: the prod-pollution boundary is *small + fixed*
(Directus, S3-public, live Clerk/external SaaS — everything in the per-stack Postgres is isolated); the M0 alignment
pattern *extends to data* (new structural operators + schema-as-source); the perf bottleneck is *DB-IO, not CPU*
(Go-link-ent + `COPY` + fan-out; Rust buys nothing). Decisions: **3-way split, all in v1.1** (M7a-D1, user chose
keep-in-v1.1 over a v1.2 spin-out) · **the isolation guard is the load-bearing deliverable** (M7a-D2) · **extend
M0 to a data dimension, don't fork it** (M7b-D1) · **the data-DNA is the catalog that drives the fleet** (M7b-D2)
· **the fleet is iterative, gated on data-DNA coverage** (M7c-D1).

### M3: Disposable multi-instance demo stacks ✅ DONE (2026-06-03; extended close 2026-06-04)
**Status:** `done` · **Shape:** `section` · **Dir:** [m3-demo-stacks/](releases/archive/01.10-show-floor/m3-demo-stacks/)
Spun up `demo-N` as isolated, Clerkenstein-wired full stacks; the full Clerk-free injected stack + migrate are
LIVE-PROVEN; the deployment/injection alignment surface (`clerk-deploy-1`, 7/7) landed. 78 demo-stack tests, 218
clerkenstein funcs. **Delivered** `corpus/ops/rosetta_demo.md` + `/demo-*` skills.

### M4: Consolidate into the `rosetta-extensions` monorepo ✅ DONE (2026-06-04)
**Status:** `done` · **Shape:** `section` · **Dir:** [m4-consolidate-extensions/](releases/archive/01.10-show-floor/m4-consolidate-extensions/)
Created the **`rosetta-extensions`** monorepo (private, 73 commits); `git subtree`-imported `clerkenstein` +
`rosetta-demo`(→`demo-stack`) **with full history preserved**; the `knowledge/` nav; thinned rosetta to pointers;
fixed a +1-depth path break the verify gate caught (M4-D4); verified under the new paths (78 demo-stack tests +
deploy gate 7/7); pushed; **removed the old `clerkenstein` + `rosetta-demo` repos** (local + org, 404). Decisions
M4-D1 (subtree) / D2 (delete-not-archive) / D3 (alignment framework stays in rosetta) / D4 (path-depth fix).

### M5: Extract the reusable `stack-injection` layer ✅ DONE (2026-06-04)
**Status:** `done` · **Shape:** `section` · **Dir:** [m5-stack-injection/](releases/archive/01.10-show-floor/m5-stack-injection/)
Extracted the generic injection (`inject.py`, `gen_injected_override.py`, `apply-authn.sh`) into
`rosetta-extensions/stack-injection/`, consumable by any stack with a **demo-ON / dev-OFF** toggle; the mock stayed
in clerkenstein (dependency runs stack-injection→clerkenstein, M5-D1); the port-offset engine stayed in demo-stack
(M5-D2, settles the M4 open question — moves to shared in M6). Split the tests, repointed the consumers; **78
preserved**, flake 3/3, deploy gate 100%/100%.

### M6: `dev-stack` — tooled local dev environment ✅ DONE (2026-06-04)
**Status:** `done` · **Shape:** `section` · **Dir:** [m6-dev-stack/](releases/archive/01.10-show-floor/m6-dev-stack/)
Extracted the shared port-offset engine into a new **`stack-core/`** section (settles the M5-routed question —
demo + dev share it, M6-D1) and added a focused **`dev-stack/`**: isolated dev stacks (`dev-N`, offset ports,
guarded `-p dev-N`), **real Clerk by default**, Clerkenstein injection **optional** (reuses stack-injection).
Scoped to the proven value (M6-D2 — not speculative multi-dev). **87 tests** (+9), flake 3/3, deploy gate 100%/100%.

### M7a: Seeding framework + production-isolation safety ✅ DONE (2026-06-04)
**Status:** `done` · **Shape:** `section` · **Complexity:** large · **Dir:** [m7a-seeding-framework/](releases/archive/01.10-show-floor/m7a-seeding-framework/)
Built `rosetta-extensions/stack-seeding/` — a host Go module that seeds a stack by talking **directly to its
Postgres** (offset port, `COPY`; *not* ent-linking — `app/internal/bootstrap` is internal, unimportable, M7a-D3)
behind a **3-layer production-isolation guard** (CheckWrite · PreflightEnv · AssertClean). **LIVE-PROVEN**: a
fresh injected `demo-1` → `migrate-demo.sh` (now bootstraps the global Sentinel policy) → `stackseed` (org + 1000
users + the real `user_clerkenstein` identity + the casbin `g2` grant, isolation audit clean) → authenticated
login returns **HTTP 200** (`membershipsCount: 1001`). The proof caught + fixed **2 real bugs** (the g2 arg-order;
the missing global-policy bootstrap — M7a-D4). **68 tests**, all gates green. Delivered `corpus/ops/seeding-spec.md`.

### M7b: The data-alignment dimension ("data DNA") ✅ DONE (2026-06-04)
**Status:** `done` · **Shape:** `section` · **Complexity:** medium · **Dir:** [m7b-data-dna/](releases/archive/01.10-show-floor/m7b-data-dna/)
Extended the **M0 alignment framework** to a **data** dimension — the `datadna` harness (`rosetta-extensions/
stack-seeding/dna/`) that (a) enumerates the seedable surfaces (**4 seeded + 6 planned** — the M7c checklist) and
(b) measures a seeder's output conforms to the platform's **current schema** via **structural operators**
(type-match / constraint-satisfied [NOT-NULL + UNIQUE] / fk-valid / row-count) with **schema-as-source via
introspection**. A separate harness, not an alignctl runner (M7b-D3). **PROVEN live** on the M7a-seeded `demo-1`:
`measure` **100% / Critical 100%** across the 4 seeded surfaces; `diff` flags an injected column (exit 1) and
reads clean on revert. Caught + fixed the planned-surface introspection bug; hardened the UNIQUE leg (M7b-D4).
**dna 49 + cmd/datadna 10 + pg 17 tests.** Delivered the data dimension into `corpus/architecture/alignment_testing.md`.

### M7c: The seeder fleet, to a coverage gate ✅ DONE (2026-06-05, gate-met-over-reachable + waiver)
**Status:** `done` · **Shape:** `iterative` · **Complexity:** large · **Dir:** [m7c-seeder-fleet/](releases/archive/01.10-show-floor/m7c-seeder-fleet/)
Built the fleet across 5 iters (TOK-01 strategy → jobsim-sessions → skillpath-sessions → assignments → activity),
each a deterministic **backdated-activity** seeder (time-distributed, pass/fail per `pass_rate`, content refs as
free values — the believability core is reachable **without** the shared Directus). Drove data-DNA coverage
**40%→80%**, promoting each surface planned→seeded + conformance-gated. **Gate: 3 of 4 met outright** — (a)
login→**200** · (c) full 8-seeder seed **0.69s** (~8500 rows, <2min) · (d) isolation **clean**; (b) coverage is
**100% over the 8 reachable surfaces / critical 100%**, with **taxonomy + content waived** (the hard line —
skiller snapshot + shared Directus; Re-scope trigger, user-confirmed → ~v1.2). Caught + fixed 2 live bugs (the
skillpath UNIQUE constraint; the introspect-load harness bug). **20 seeder / 145 module tests.** Delivered
`rosetta-extensions/stack-seeding/seeders/` + the `waived` data-DNA status.

### M8: Corpus + use-case recipes + polish ✅ DONE (2026-06-05) — LAST v1.1 milestone
**Status:** `done` · **Shape:** `section` · **Complexity:** medium · **Dir:** [m8-corpus-recipes/](releases/archive/01.10-show-floor/m8-corpus-recipes/)
The consolidation/discoverability layer: a **`corpus/ops/demo/` family** (index + 3 end-to-end recipes —
enterprise-onboarding, skill-progression, browser-login [which lands the 2 M3-deferred injection recipes: the
`api.clerk.com` cert-redirect + the browser-login walk-through]); **3 seed presets** (small/mid/large, mid-500 +
large-1k seed-proven end-to-end); the **`/demo-seed` skill** + the CLAUDE.md skill table; the v1.0
**express-gate CI carry-forward** wired into clerkenstein `alignment.yml` (**validated 9/9** locally); and
cross-linking from corpus/README + root README + CLAUDE.md (all doc links resolve). **Next:** `/developer-kit:close-release`.

### Execution graph (v1.1)
```
v1.1 "show floor" — the platform-operations extension framework (demo + dev, in 2 repos)
   M3 ✅ ─→ M4 ✅ (consolidate) ─→ M5 ✅ (stack-injection) ─→ M6 ✅ (dev-stack)
                                            └──→ M7a (framework+safety) ─→ M7b (data-DNA) ─→ M7c (seeder fleet) ─→ M8 (recipes)
```
**Sequential.** M4–M6 shipped (the extension framework + demo/dev stacks). M7a lands the framework + the
isolation guard (a usable, safe demo); M7b builds the data-DNA catalog that lists + gates the seeders; M7c drives
the fleet to the coverage gate; M8 curates the output.

### Risks (v1.1)
- **(M7a, blocks-prod-safety)** a single un-guarded **shared-write reaching prod** (Directus / S3-public bucket) —
  mitigate with the hard isolation guard + the clean-audit assertion as a tested acceptance gate, not a convention.
- **(M7a, scope)** linking the platform's `app/internal/bootstrap`/ent client into a `rosetta-extensions/` Go
  module without a platform edit — confirm the import path early (fallback: `go run` CLIs, slower).
- **(M7b)** trustworthy schema-as-source — get ent introspection / `atlas inspect` golden right or the drift diff lies.
- **(M7c, scope)** the heaviest build: ~8–10 seeders + 1k-scale `COPY` perf + backdating fidelity, each gated on
  conformance — the believable-demo *subset* of surfaces is the real target (waive unreachable genes, don't chase 100%).

### Open decisions (resolve during build)
Directus snapshot-replay vs hard-block-and-skip for the demo MVP (M7a); ent-introspection vs `atlas inspect`
golden for schema-as-source (M7b); whether seed presets ship in M7c or M8; external shareability (Tailscale vs
ingress); the AI-content STRETCH trigger (now firmly v1.2, not M8).

## Done — v1.0 "body double" (SHIPPED 2026-06-03 · tag `v1.0`)

> **Shipped 2026-06-03.** All six milestones closed-on-gate / completeness-complete and merged to `main`;
> `release/01.00-body-double` deleted. Release records archived under
> [`releases/archive/01.00-body-double/`](releases/archive/01.00-body-double/) (review · retro · metrics ·
> lockfile · stats). Headline: a *measured* drop-in Clerk mock at **100%/100% on all three surfaces**
> (Go · JS/FAPI · `@clerk/express`), built by a first-class alignment framework, zero platform-code change.
> Close-release caught + fixed 1 blocker (an `@clerk/express` gate regression from the M2c close) — see the
> release retro.

**Theme:** Clerk authentication is the friction that blocks fast, throwaway demos. v1.0 delivers
**Clerkenstein** — a drop-in mock that mirrors the exact Clerk interface the platform uses, with
security/sync disarmed, injected via build-time `go.mod replace` + skip-worktree so **every platform
repo keeps "thinking" it uses Clerk with zero source changes**. The novelty: Clerkenstein isn't a
hand-built mock — it's the **first mirror produced by a reusable, measurable alignment process**
(M0). We don't just claim the stand-in is faithful; we **score** it (0–100%) against the real Clerk
and CI-gate that score against drift. This also removes Clerk's API rate limit as the blocker for
scale data-seeding in v1.1.

**Decided at design (2026-06-02):** two-version split (Clerkenstein first); **alignment is a
first-class test class** with its own framework (M0); **M1 is iterative** (its exit gate is an
alignment score); M2 frontend = attempt the fake Clerk FAPI server, **fall back to the real dev Clerk
app for the browser session** (backend stays fully mocked) if base-URL override proves too fragile.

### Alignment vocabulary (the M0 model, referenced by M1/M1b)
- **Target** — an engine exposing a surface. **Source target** = the canonical engine, version-pinned (Clerk `clerk-sdk-go/v2 @ v2.6.0`). **Mirror target** = our reimplementation (Clerkenstein).
- **Capability** — one endpoint/function of the source surface *(axis 1)*. **Variant** — one input/scenario class for a capability (standard + corner + error) *(axis 2)*.
- **Alignment test** — one **(capability × variant)** pair; feeds identical input to both targets and asserts behavioral equivalence. A **third test class** alongside unit & integration; **tagged** so it's parseable/countable/runnable as its own suite.
- **Alignment DNA** — the officially-enumerated complete set of (capability × variant) **genes** for a source target at a version; the machine-readable manifest that *defines* faithfulness and is the score's denominator.
- **Alignment score** — `aligned genes ÷ total genes × 100`, with a per-capability rollup. 100% = behaviorally indistinguishable across the whole DNA.

### M0: Alignment measurement framework
**Status:** `done` (2026-06-02)
**Shape:** `section`
**Goal:** A reusable, engine-agnostic process — two skills + a test class + a manifest format — that measures how faithfully any mirror reproduces any source engine, producing a 0–100% alignment score. This is the foundation M1 builds on and M1b reuses.

**Closed 2026-06-02** (build S1–S5 → harden 2 passes → close review → merged to `release/01.00-body-double`). Delivered: `test/alignment/` — `alignctl` (stdlib-only Go, builds/runs offline) with `run`/`capture`/`dna list|diff|validate`, engine-agnostic via a pluggable `--runner`; the 4 equivalence operators + weighted score (overall + separate critical gate); record/replay goldens; `internal/canon` precision-safe canonicalization; the `//go:build alignment` test class; and a toy reference proving end-to-end detection (**86.7% / 100% critical**, catches `Greet/padded-name`). Plus `/align-dna` + `/align-run` skills and `corpus/architecture/alignment_testing.md`. Open questions resolved: DNA format = **JSON** (M0-D1); capabilities enumerated from the consumed surface; goldens live per-mirror-repo. Close-review adversarial pass found + fixed a path-traversal must-fix + score-overflow (M0-D7); 45 test funcs (3 fuzz), 5/5 flake gate, core coverage 83–98%. Decisions M0-D1…D7. Retro: [m0-alignment-framework/retro.md](releases/archive/01.00-body-double/m0-alignment-framework/retro.md). Resolved repo split: framework in rosetta; the Clerk DNA/tests/mirror land in the `clerkenstein` repo (M1).
**Scope:**
  - In:
    - **`/align-dna` skill** (build & update alignment targets): given a source framework + version, pull the pinned source into `.agentspace/`; enumerate the **consumed** capabilities (scoped to what the platform calls, not the whole SDK); enumerate standard + corner-case **variants** per capability; emit/update the **Alignment DNA** manifest (each gene: input fixture, expected-shape descriptor, equivalence operator, criticality weight); **diff DNA across source versions** (added/removed/changed genes); **scaffold alignment-test stubs from the DNA** so tests never drift from the manifest.
    - **`/align-run` skill** (measure alignment of 2 targets): given a DNA + source version + mirror, pull the source, run every gene against **both** targets, assert equivalence per the gene's operator, compose the **0–100% score** + a per-capability divergence report.
    - the **alignment test-class convention** (tagging/marking so tests are discoverable + countable, distinct from unit/integration), the **DNA file format**, the **equivalence operators** (exact / same-shape / normalized / same-error-class), and **record/replay (golden capture)** support so a live-SaaS source can be measured reproducibly offline.
    - a **tiny toy reference mirror** (≈2 capabilities) proving the framework runs + scores end-to-end, independent of Clerk.
  - Out: the Clerk DNA + the real Clerkenstein mirror (M1); drift CI wiring (M1b); the JS surface (M2).
**Depends on:** none.
**Parallel with:** none (gates M1, M1b).
**Estimated complexity:** large
**Open questions:** DNA manifest format (YAML vs Go structs); how capabilities are enumerated (parse source surface vs curated list); where golden captures live + how they're refreshed.
**KB dependencies:** none new (greenfield — alignment is a documentation blind area).
**Delivers → `corpus/architecture/alignment_testing.md`:** the alignment test class, the DNA format, the two skills, equivalence + record/replay — the canonical reference (net-new doc).

### M1: Clerkenstein backend mirror (Go)
**Status:** `done` (2026-06-03, **closed-on-gate**)
**Shape:** `iterative`
**Goal:** The first real mirror — a drop-in Go stand-in for `colony/authn`'s provider + the Clerk `orgclient`, built *by* the M0 process and injected via `go.mod replace` (zero platform-repo edits), so backend services authenticate with one universal credential and locally-minted JWTs.

**Closed 2026-06-03** (5 iters: bootstrap tok TOK-01 → DNA → authn twin → critical orgclient → standard orgclient → **gate** → final harden → close). The **Clerkenstein backend mirror** (in the gitignored `anthropos-demo/clerkenstein` repo, its own git) scores **100% alignment / 100% critical** against the `clerk@2.6.0` DNA (22 genes), built offline. authn implements the real `colony/authn.Provider` (HS256, one universal key); orgclient is a disarmed in-memory twin. Score arc: 0 → 21.1 → 68.4 → **100%**. Final harden: authn + orgclient **0 → 100%** unit coverage (+1 fuzz, 0 bugs). Decisions: D1 hybrid goldens; iter-01-D1 authn injects via `go.mod replace` whole-colony; **M1-D2 orgclient injects via a fake-Clerk-API-server → routed to M2** (shared HTTP-interception with the JS side). Delivered `corpus/services/clerkenstein.md`. Retro: [m1-clerkenstein-backend/retro.md](releases/archive/01.00-body-double/m1-clerkenstein-backend/retro.md). The gate (alignment fidelity) is met; live injection into a running platform is rosetta-demo work (v1.1) / M2 (orgclient).
**Exit gate:** `/align-run` reports **100% alignment on the platform-consumed Clerk Go surface (critical capabilities) and ≥95% overall**, with any waived genes documented + justified in the divergence report.
**Iteration protocol:** `corpus/architecture/alignment_testing.md` (the M0-delivered alignment-measurement process) — the measure → fix-diverging-genes → re-measure loop.
**Why iterative (not section):** the deliverables are writable, but *which genes diverge and how costly each is to close* only emerges from measuring against the real Clerk — a fixed up-front checklist would be speculative. The score is the commitment; the path to it is open.
**Depends on:** M0 (its skills + DNA format + test class).
**Parallel with:** none (gates M1b and M2).
**Estimated complexity:** large
**Re-scope trigger:** if consecutive strategy iters (toks) can't close a diverging gene (e.g. a capability that's fundamentally unmockable offline), waive it with justification or escalate to the user — don't chase an unreachable 100%.
**Open questions:** which capabilities need live-Clerk record/replay vs pure local mint; the precise critical-capability set; stub just `authn`+`orgclient` or `replace` all of `colony` (`authn` is a package inside `colony`) — fallback is vendoring whole `colony`, as staging already does.
**KB dependencies:** `corpus/architecture/alignment_testing.md` (the iteration protocol), `corpus/services/clerk-integration.md`, `corpus/architecture/shared_libraries.md` (§ authn/colony), `corpus/ops/staging-clerk.md`, `corpus/ops/webhook_setup.md`.
**Delivers → `corpus/services/clerkenstein.md`** (the mirror design + injection mechanism — net-new) **+ the Clerk Alignment DNA** (`clerk@2.6.0` genome, authored via `/align-dna`).

### M1b: Clerk drift detection
**Status:** `done` (2026-06-03)
**Closes the gap after:** M1 (Clerkenstein is aligned at v2.6.0 — but must *stay* aligned as the platform bumps `clerk-sdk-go` / `@clerk/*`).

**Closed 2026-06-03** (2 sections + 1 harden pass). Automation/config over M0 — no new measurement machinery. In the clerkenstein repo: `scripts/gate.sh` (alignment gate, built-binary so exit 0 met / 2 regressed) + `scripts/drift-check.sh` (DNA-diff + gate, exit-code contract **0** none / **1** DNA moved / **2** gate regressed / **3** usage) + `.github/workflows/alignment.yml` (push + **weekly** CI) + `scripts/drift-test.sh` (9-assertion regression harness pinning the contract + the 2 build-phase fixes). Delivered the "Drift detection (M1b)" runbook in `corpus/services/clerkenstein.md`. Verified across all exit paths against a simulated `clerk@2.7.0` bump; shellcheck clean, flake 5/5. Close review: 0 findings.
**Goal:** Reuse M0 wholesale to make Clerk drift a flagged, mechanical event: on a version bump, `/align-dna` diffs the DNA (what changed) and `/align-run` re-scores the existing mirror against the new source (score drop = broken genes), CI-gated on "alignment ≥ threshold."
**Scope:**
  - In: the "bump pinned Clerk version → DNA-diff → re-score → report" workflow; the CI gate on alignment score; golden-capture refresh on bump.
  - Out: building the framework (M0); authoring the original mirror/DNA (M1); the JS surface (M2 owns its own genes).
**Depends on:** M1 (needs a built, aligned mirror + the Clerk DNA). Reuses M0's skills — **now automation/config over M0, not new machinery** (the right size for a B-milestone).
**Parallel with:** M2 (CI/automation vs JS code — disjoint surfaces).
**Acceleration effect:** every future Clerk bump becomes a flagged, scored update instead of a silent break — the brief's "follow platform updates within minutes" requirement, mechanized.

### M2: Clerkenstein — browser session + webhook coherence (JS)
**Status:** `done` (2026-06-03)
**Shape:** `section`
**Goal:** The frontend logs in with no real Clerk, and created/seeded users/orgs reach the DB without real Clerk webhooks.

**Closed 2026-06-03** (5 sections S1–S5 → 4 harden passes → close review → merged to `release/01.00-body-double`). Closes the last two Clerk seams so a demo stack is **Clerk-free end to end**. Delivered (in the gitignored `anthropos-demo/clerkenstein` repo): the **fake FAPI server** (`fapi/`) + the publishable-key codec — the browser logs in via a *minted publishable key* that encodes the fake FAPI host, **config-only, no SDK fork** (M2-D1 spike resolved the milestone's defining risk in the strong direction; the real-dev-Clerk fallback is documented but un-exercised); the **fake BAPI server** (`bapi/`) that disarms the platform's networked `orgclient` via an `api.clerk.com` DNS/base-URL redirect (the **M1-D2 Fate-3 pickup**), backed by the M1 orgclient twin made **concurrency-safe** (M2-D2); the **svix-signed webhook injector** (`webhook/`) for the 12 consumed event types → `POST /api/webhook/clerk`; and a **second Alignment DNA** (`clerk-js-5`, 9 genes, runner `cmd/jsfapirun`) scored at **100%/100%** like the Go side — proving the M0 framework is **surface-generic**. Both gates 100%/100% (Go 22/22 + JS 9/9); 112 Go test/fuzz funcs; flake 5/5; gofmt/vet/shellcheck clean. **Close review** found + fixed an `orgclient.ChangeRole` nil-map panic + phantom-membership divergence the alignment gate missed (reachable via the `bapi/` server) — M2-D4, with regression tests; plus a gofmt fix + the repo README refresh; 0 scope gaps, 0 deferrals (deferral audit GREEN). Decisions M2-D1…D4. Retro: [m2-browser-webhook-coherence/retro.md](releases/archive/01.00-body-double/m2-browser-webhook-coherence/retro.md). **This was the last *feature* milestone of v1.0**; a cleanup B-milestone **M2b (repo consolidation)** was inserted after it (2026-06-03) to tidy the `clerkenstein` repo before `/developer-kit:close-release`.
**Scope:**
  - In: a fake Clerk FAPI path for `@clerk/nextjs ^6.39.2` (next-web-app, ant-academy) and `@clerk/clerk-js ^5.52.3` (studio-desk) via publishable-key + base-URL/DNS override — **with the decided fallback**: keep the real dev Clerk app for the browser session while the backend stays fully mocked; a **webhook injector** feeding the existing `app/internal/clerk/events/` sync pipeline directly; **the JS surface's fidelity expressed as alignment genes via M0** where applicable (same score treatment as the Go side).
  - In (**routed from M1 close — M1-D2, Fate 3**): the **fake-Clerk-API-server** (HTTP interception of `api.clerk.com`) ALSO serves M1's **orgclient** injection — the Go `app/internal/clerk/orgclient` is app-internal + networked, so it can't `go.mod replace` like authn; it disarms via the same fake-API-server this milestone builds for the JS side. The Clerkenstein orgclient mirror behavior already exists + scores 100% (M1); M2 wires the HTTP redirect that makes the platform's real orgclient hit it.
  - Out: multi-instance stacks (M3); data seeding (M4).
**Depends on:** M1 (consumes the mock contract + minted-token shape). **Parallel with:** M1b (yes).
**Estimated complexity:** large — **highest technical risk in v1.0** (SDKs hard-code Clerk FAPI; no documented base-URL override).
**Open questions:** can `@clerk/*` be pointed at a fake FAPI without a fork? (the fallback exists because this is uncertain) — spike the override early.
**KB dependencies:** `corpus/architecture/alignment_testing.md`, `corpus/services/clerk-integration.md`, `corpus/architecture/frontend_architecture.md`, `corpus/services/next-web-app.md`, `corpus/ops/webhook_setup.md`.
**Delivers → `corpus/services/clerkenstein.md`:** extends the M1 doc with the JS path + webhook injection + the fallback decision.

### M2b: Clerkenstein repo consolidation + knowledge base
**Status:** `done` (completed 2026-06-03)
**Shape:** `section`
**Dir:** [m2b-clerkenstein-consolidation/](releases/archive/01.00-body-double/m2b-clerkenstein-consolidation/)

**Closed 2026-06-03** (5 sections S1–S5 → 1 harden pass → close review → merged to `release/01.00-body-double`). A pure-cleanup B-milestone that reorganized the `clerkenstein` repo (gitignored `anthropos-demo/clerkenstein`, its own git on `main`) into a clean, self-documented **library-named** structure — **no behavior change**, both alignment gates (Go 22/22, JS 9/9) + the drift harness (9/9) stayed green throughout. Delivered: the **library-named dirs** (`authn/` mocks colony/authn · `clerk-backend/` mocks clerk-sdk-go/v2 = the bapi server + orgclient store **merged** · `clerk-frontend/` mocks @clerk/clerk-js+nextjs · `clerk-webhook/` mocks svix) + `shared/` (the universal-key HS256 JWT, extracted because `clerk-frontend` **mints** and `authn` **verifies** the same token — `parse`→`shared.Parse` exported, M2b-D4) + `alignment/` (the M0-consumption harness: `cmd/{clerkrun,jsfapirun}` + `dna/` + `golden{,-js}/` + `scripts/`) via **69 history-preserving `git-mv` renames**; a self-contained **`knowledge/` base** (kb-index + scope + architecture + alignment + injection + coverage-index) + 6 per-library READMEs + slim root README; an `.agentspace/` (gitignored contents, dir preserved) + `.gitignore` baseline + asset hygiene; and `CLAUDE.md` + `singularity-manifest.md` (authored TO the `/singularity-kit:repo-consolidate` standard — the formal `repo-consolidate code` run is a **USER finalize**, M2b-D3/D8, since the skill is `disable-model-invocation`). Rosetta-side: slimmed `corpus/services/clerkenstein.md` 197→62 lines to a pointer at the repo's KB + fixed 2 stale refs in `alignment_testing.md`. **Close review** found + fixed 1 should-fix code-quality (a fuzz-test comment naming pre-reorg packages) + 2 doc findings (coverage-index count drift 112→113 after the harden test, state.md Headline refresh) — clerkenstein fixes on its own `main` (`ad87545`); 0 scope gaps, 0 deferrals (deferral audit **GREEN** — 2 inherited singles owned by close-release/M3, 0 repeat). Decisions M2b-D1…D8; D1/D2/D4 blended into the repo's own KB. Retro: [m2b-clerkenstein-consolidation/retro.md](releases/archive/01.00-body-double/m2b-clerkenstein-consolidation/retro.md). **This was the LAST milestone of v1.0** → next is `/developer-kit:close-release`.

**Goal:** The `clerkenstein` repo grew organically across M1/M1b/M2 into flat package dirs (`authn bapi orgclient fapi webhook cmd dna golden golden-js scripts`) with a single README and no knowledge base. M2b reorganizes it into a clean, self-documented **library-named** structure — one dir per mocked dependency + a shared dir + an alignment harness dir + a `knowledge/` base — following `/singularity-kit:repo-consolidate`, so the repo is navigable + operable by agents *before* v1.0 ships.
**Context (B-milestone — cleanup after M2):** pure reorg / docs / hygiene over the M2-complete repo. **No behavior change** — both alignment gates (Go 22/22, JS 9/9) and the drift harness stay green throughout; the move repoints imports + DNAs/goldens/runners/scripts, it does not alter the mocks. Class of work like M1b (tooling/cleanup over a shipped surface).
**Scope:**
  - In (**1 — Restructure**): one dir per mocked library/framework + one shared dir, **library-named** (user-chosen scheme): `authn/` (mocks `colony/authn`), `clerk-backend/` (mocks `clerk-sdk-go/v2` — the `bapi` server + the `orgclient` store **merged into one dir**), `clerk-frontend/` (mocks `@clerk/clerk-js` + `@clerk/nextjs` — the FAPI), `clerk-webhook/` (mocks `svix`); `shared/` (universal-key HS256 JWT + claims + canonical helpers — extracted because `clerk-frontend` mints and `authn` verifies with the same key); `alignment/` (the M0-consumption harness: `cmd/{clerkrun,jsfapirun}` + `dna/` + `golden{,-js}/` + `scripts/`). **Tests stay co-located within each library dir.** Go package identifiers can't contain hyphens → each hyphenated dir declares a clean package (e.g. `clerk-backend/` → `package clerkbackend`) — M2b-D1, confirmed at build.
  - In (**2 — Knowledge base**): a self-contained `knowledge/` dir documenting Clerkenstein — scope/goal; how it's built (the 4 mocks + shared); how fidelity is **validated with alignment tests against a pinned Clerk version** (the M0 framework + the two DNAs + the gate); **per-library injection recipes** (`go.mod replace` for `authn`; `api.clerk.com` HTTP/DNS redirect for `clerk-backend`; config-only publishable-key override for `clerk-frontend`; direct svix-signed POST for `clerk-webhook`); a coverage index. Per-library `README.md`s + a top-level index. Solid, well-written, well-distributed.
  - In (**3 — Hygiene**): an `.agentspace/` dir with contents **gitignored**; `.gitignore` cleanup (the current comment is mismatched); built-binary + transient hygiene per `repo-consolidate`'s asset-hygiene checks.
  - In (**4 — Consolidate**): run `/singularity-kit:repo-consolidate code` to standardize the repo (emit `CLAUDE.md` + `singularity-manifest.md`, audit against the code-repo + asset-hygiene standards, apply fixes), then re-verify both gates + the drift harness. **Note:** `repo-consolidate` is `disable-model-invocation` (user-invoked) — the build authors the structure TO its standard so the run is a clean finalize; the **user types the skill** (pointed at the `clerkenstein` repo).
  - Out: new library support / new alignment genes (the `@clerk/express` coverage gap — **now picked up by M2c**); any live injection wiring into a running platform (still v1.1/M3); any change to rosetta's M0 framework or to the platform repos.
**Depends on:** M2 (consolidates the M2-complete repo). **Parallel with:** none (touches the whole repo). **Precedes:** `/developer-kit:close-release`.
**Estimated complexity:** medium — mechanical but wide (touches every package + the gate/drift scripts); the only real risk is import/script repointing, fully caught by the **green-gate invariant** (gates + drift re-run after each section).
**KB dependencies:** `corpus/services/clerkenstein.md`, `corpus/architecture/alignment_testing.md`; the `/singularity-kit:repo-consolidate` standards (base + code-repo + asset-hygiene).
**Delivers → the `clerkenstein` repo's own `knowledge/` base** (net-new, self-contained) **+ slims `corpus/services/clerkenstein.md`** (rosetta) to a pointer at the repo's `knowledge/` + the new structure.

### M2c: Clerkenstein — `@clerk/express` backend session verification (RS256/JWKS)
**Status:** `done` (2026-06-03, **closed-on-gate**)
**Shape:** `iterative` (alignment-score gate, like M1) — a **feature** milestone; the letter suffix marks *insertion after M2b*, not a B/tooling milestone.

**Closed 2026-06-03** (5 iters: bootstrap TOK-01 → DNA → RS256 foundation → **crux proof** → full runner → gate; 1 final harden pass). Brought the **last un-gated Clerk consumer — `@clerk/express`** (studio-desk's Node backend) under the alignment framework at **100%/100%** (3rd DNA `clerk-express-1.json`, 9 genes). The **RS256 wall fell to an additive path** (M2c-D1/D2): an RSA keypair + a real JWKS + RS256 minting that the *genuine* `@clerk/backend` accepts networkless via `jwtKey` — **no HS256 migration**, so M1 (22/22) + M2 (9/9) stayed green. `@clerk/express` is **verified, not reimplemented** (no mock dir — the svix discipline; M2c-D5); the `expressrun` runner mints tokens (Go) + drives the real SDK (embedded `verify.js`, Node). The `clerkClient` BAPI reads were already covered by `clerk-backend` (M2c-D4). Close: folded the surface into the knowledge base + corpus, fixed a gitignore gap + 1 adversarial flake (`tamperSig`); deferral audit GREEN; the express-gate CI-wiring (needs Node) routed to v1.1. 128 test/fuzz funcs / 8 packages; all four gates green. Retro: [m2c-clerk-express-alignment/retro.md](releases/archive/01.00-body-double/m2c-clerk-express-alignment/retro.md).
**Dir:** [m2c-clerk-express-alignment/](releases/archive/01.00-body-double/m2c-clerk-express-alignment/)
**Goal:** Bring the **last un-gated Clerk consumer — `@clerk/express`** (studio-desk's Node backend auth) under the alignment framework: a new **`clerk-express/`** seam + a **3rd Alignment DNA**, driven to a gate, so studio-desk's backend genuinely verifies Clerkenstein tokens (not via its `MOCK_CLERK=true` bypass). Completes v1.0's thesis — *no* Clerk seam left un-faithful before shipping.
**Why iterative + the defining unknown (the RS256 wall):** `@clerk/express` (via `@clerk/backend`) verifies **RS256 via JWKS only** and **hard-rejects HS256** (`assertHeaderAlgorithm` → `TokenInvalidAlgorithm`). Clerkenstein mints HS256 universal-key tokens + serves an **empty JWKS**, so an HS256 shim is a dead end. The milestone must add an **RS256 path** (RSA keypair + a real JWKS from the fake FAPI + RS256 minting + the real-`@clerk/express` verifier). **The central iteration question:** can RS256 be **additive/parallel**, or must the existing HS256 seams (`authn`/`clerk-frontend`/`shared`) **migrate to RS256** — re-capturing the Go DNA goldens + re-gating M1/M2? The gate-driven iterations resolve it.
**Scope:**
  - In: a new **`clerk-express/`** seam (library-named); an **RSA keypair + a real (non-empty) JWKS** served by the fake FAPI (`clerk-frontend`'s `/.well-known/jwks.json`); **RS256 token minting**; the `@clerk/express` **DNA** (`clerk-express-1.json`, source `@clerk/express ^1.3.47`); a runner that drives **the real `@clerk/express` SDK** against the mock (the svix-pattern — verify against the genuine library); the **alignment gate** as the exit criterion.
  - In (confirm, don't rebuild): `@clerk/express` also calls `clerkClient.{getOrganizationMembershipList, getOrganization}` — those are **BAPI**, already 100%-mocked by `clerk-backend/`; M2c adds *integration* genes confirming that path, not a new BAPI mock.
  - Out: changing studio-desk or any platform repo (the `MOCK_CLERK` bypass is the platform's own); a webhook (svix) DNA (separate future gap); live injection into a running studio-desk (rosetta-demo work, v1.1).
**Candidate genes (~8, `clerk-express-1.json`):** `ExpressAuth/{valid, expired, malformed, bad-signature, no-token}` (error_class) · `ExtractIdentity/universal-user` (exact: verified claims → `req.auth`) · `JWKS/non-empty-rsa` (shape) · `ClerkClientBAPI/{org-membership-list, get-organization}` (integration vs `clerk-backend`).
**Exit gate:** alignment **≥ 95% overall / 100% critical** on `clerk-express-1.json`, AND the load-bearing test passes (a **real `@clerk/express` instance accepts a Clerkenstein-minted token + extracts the right identity**).
**Depends on:** M2 (the FAPI + token machinery it extends) + M2b (the consolidated repo it adds a seam to). **Precedes:** `/developer-kit:close-release`.
**Estimated complexity:** large — **highest fidelity-risk in v1.0**: the RS256 path may force a token-algorithm migration of the existing 100%-gated seams.
**KB dependencies:** `corpus/architecture/alignment_testing.md`; the clerkenstein repo's own `knowledge/` (alignment / architecture / injection / sources); the `@clerk/express` + `@clerk/backend` source under `anthropos-dev/studio-desk/node_modules`.
**Delivered → the clerkenstein repo's `knowledge/`** (alignment/architecture/sources updates) **+ a 3rd DNA + the `expressrun` runner;** updated `corpus/services/clerkenstein.md`'s scorecard to a **3rd *measured surface*** (`@clerk/express`, **verified-not-mocked** — no new mock dir, per M2c-D5; the genuine SDK is *satisfied* via an additive RS256/JWKS path).

### Execution graph

```
v1.0 "body double"   — a stand-in the platform can't tell apart, and we can prove it

  M0 (alignment framework: /align-dna + /align-run, test class, DNA format, golden capture, toy ref)
    │
    ↓
  M1 (Clerkenstein backend mirror — ITERATIVE: author Clerk DNA → drive alignment score to gate)
    │
    ├──→ M1b (Clerk drift detection — DNA-diff + re-score, CI-gated across version bumps)   ∥ M2
    └──→ M2 (browser session + webhook; reuses the alignment class for the JS surface)
              │  (both closed — repo feature-complete)
              ↓
    M2b (repo consolidation — library-named dirs + self-contained knowledge base; gates stay green)
              │
              ↓
    M2c (ITERATIVE: @clerk/express RS256/JWKS — new clerk-express/ seam + 3rd DNA → alignment gate)
              │
              ↓
    /developer-kit:close-release → v1.0 ships to main
```

### Parallelism

- **M0 → M1 → {M1b, M2}** sequential at the core: M1 needs M0's framework; M1b + M2 need M1's mirror/contract.
- **M1b ∥ M2:** disjoint surfaces — M1b is CI/automation over M0; M2 is JS + the webhook injector. Merge risk **low**.
- **M3 ∥ M2 (cross-version, yes-with-caveats):** sequenced cleanly by the version boundary (M3 starts after v1.0 closes).

### Risks (v1.0)

| Risk / decision | Severity | Mitigation |
|---|---|---|
| **Source is a live SaaS** — Clerk's API capabilities can't be hit freely/offline/deterministically | blocks-release (reproducibility) | M0 **record/replay golden captures** is a core requirement, not an afterthought — capture once, replay forever |
| **DNA completeness gaming** — 100% on a thin DNA is hollow | degrades-quality | `/align-dna` capability-coverage check (every platform-consumed endpoint present) + M1b version-bump DNA-diff keeps it complete |
| **Defining "equivalent"** — timestamps, generated IDs, error formats differ even when behavior matches | degrades-quality | M0 ships **equivalence operators** (exact / same-shape / normalized / same-error-class) chosen per gene |
| **JS/FAPI fake server** — SDKs hard-code Clerk FAPI, no base-URL override | blocks-release (full no-Clerk browser) | **Decided fallback:** real dev Clerk app for the browser, backend fully mocked; spike override early in M2 |
| **`colony` replace granularity** — `authn` is a package inside `colony`, not its own module | degrades-quality (M1 effort) | M1 early iter resolves it; fallback = vendor whole `colony` (staging precedent) |
| **Repo layout** — where the framework vs the Clerk mirror live | nice-to-resolve | **Decided:** the M0 framework (skills + format + doc) lives in rosetta; the Clerk DNA + alignment tests + mirror live in the `clerkenstein` repo, cloned into gitignored `anthropos-demo/` |
| **"Zero platform-code changes" interpretation** — `replace` edits the *clone's* go.mod | nice-to-resolve | build-time injection in the gitignored clone + skip-worktree; upstream repo never modified (same as staging's `vendor-colony/`) |

### Branch model

`release/01.00-body-double` (cut from `feat/demo-environment` at M0). Milestone branches:
`m0/alignment-framework`, `m1/clerkenstein-backend`, `m1b/clerk-drift-detection`,
`m2/browser-webhook-coherence`, `m2b/clerkenstein-consolidation`, `m2c/clerk-express-alignment`.
**M1 + M2c are iterative** → built by `/developer-kit:build-mstone-iters` (close on a Gate Outcome Ledger).
M0/M1b/M2/M2b are section → `/developer-kit:build-milestone`. All → `/developer-kit:close-milestone` →
`/developer-kit:close-release`.
The `clerkenstein` repo's own code commits stack on its `main` (its own gitignored git, no branch model);
the rosetta-side milestone records + corpus pointer land on the `m{N}/…` branch.

### Out of scope (v1.0 — recorded for v1.1+)
- Multi-instance disposable stacks, data seeding, use-case recipes → all v1.1 "show floor".
- Mirroring engines other than Clerk with M0 (the framework is generic, but v1.0 only exercises it on Clerk).
- AI-generated demo content (transcripts/embeddings) → v1.1 stretch or deferred.

## Shipped releases

- **v1.0 "body double"** — shipped **2026-06-03**, tag `v1.0`. The alignment-testing framework + Clerkenstein
  (100%/100% on Go · JS/FAPI · `@clerk/express`). Detail in the `## Done` section above; records archived at
  [`releases/archive/01.00-body-double/`](releases/archive/01.00-body-double/).

## Notes

- Milestone numbering is **flat sequential** (M0, M1, M2, …); a letter suffix has two uses: (1) a milestone **inserted after** the fact — `b` for tooling/cleanup (M1b drift CI, M2b consolidation), and the letter-suffixed *feature* milestone M2c (iterative, "inserted after M2b"); and (2) a **split** of one planned milestone into a sequential mini-arc — **M7a → M7b → M7c** is the single former M7 "seeding" split into framework+safety / data-DNA / fleet (2026-06-04, M7a-D1). Both reuse the letter suffix; context disambiguates. See [`context.md`](context.md).
- v1.0 mixes shapes: M0/M1b/M2/M2b are **section**; **M1 + M2c are iterative** (alignment-score gates).
- v1.1 "show floor" mixes shapes too: M3–M6 + M7a/M7b + M8 are **section**; **M7c is iterative** (data-DNA coverage gate).
- v1.2 "set dressing" is **all `section`** (M9a/M9b/M10/M11) — the snapshot surfaces are decomposable up front (the
  framework + 2 known surfaces + the product layer); the fidelity gate is a per-surface acceptance check, not an
  emergent-path iterative gate. (AI-content, the iterative-shaped candidate, was held to v1.3.) The former M9 split
  into **M9a (framework) + M9b (taxonomy surface)** on the 2026-06-06 refinement — the M7a→M7c precedent.
