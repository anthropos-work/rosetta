# `/demo-update` — Refresh a Running Demo Stack to Prod-Latest

> **Status:** Draft `v0.1` · spec-draft · 2026-07-10
> **Companions:** [`spec-progress.md`](spec-progress.md) (decision tracker) · [`next-release.md`](next-release.md) (out-of-scope / parking lot)
> **Target home:** `.claude/skills/demo-update/SKILL.md` + `rosetta-extensions/demo-stack/update-injected.sh` + a
> new `rosetta-demo update N` verb + `corpus/ops/demo-update.md`.
> **Peer of:** `/stack-update` (dev-only; kept as-is). This spec is why it's a *peer*, not an extension.

This spec defines a **new Rosetta capability**: `/demo-update <N>` refreshes a running `demo-N` stack to
**prod-latest** — pulling in new prod features (pages/sections, migrations, schema, RPC), reusing existing rosetta
synthetic-data machinery to populate anything the new features expose empty, and proving the result still works
via a MANDATORY fatal verify gate — **all without teardown-reseed**, so presenter state (heroes, seats,
generated batches) survives.

The three load-bearing requirements from the go-order:

1. **Feature parity** — when Anthropos launches a new section/page, `demo-update` pulls it in *automatically*.
2. **Synthetic data for new sections** — reuse existing seeders / `stacksnap` / `stackseed` / `gen-batch`; never
   leave a new section empty when tooling can populate it.
3. **Prove it still works** — end the update with a fatal verify gate (autoverify + `test-platform N live` +
   coverage sweep). Regression detection is the whole point.

---

## 1. Overview

### 1.1 North star

`/demo-update <N>` is the **update lifecycle verb** for `demo-N` — the missing sibling of `/demo-up` /
`/demo-down`. Today a stale `demo-N` has one path back to current: `demo-down` + `demo-up`, which discards
everything the demo was carrying. `/demo-update` is the **in-place, data-preserving** path.

**What "prod-latest" means concretely.** For each service repo cloned per-`N` under `stack-demo/demo-N/repos/`,
the target ref is the **highest semver `v*` tag** on that repo's default remote branch (the same
`git tag --list 'v*' --sort=-v:refname | head -1` rule `up-injected.sh` already uses at bring-up). Per-repo
overrides are allowed via `--ref <svc>=<ref>` (reusing `demo-stack/lib/clone_repos.py:resolve_ref()`).

**What survives.** The stack's Postgres schemas + data survive (subject to migrations). The generated-batch
cache under `.agentspace/.batchcache/` survives (invalidated only when the mother prompt or taxonomy digest
moves). The per-`N` env, offset ports, Clerkenstein PK, and registry entry are unchanged.

**What is intentionally excluded from `demo-N` prod-latest:** Ithaca-side presenter chrome (secrets in
`stack-demo/.env`, mkcert certs, ant-academy .env.local). Those are re-run only under `--reprovision`.

### 1.2 Relationship to `/stack-update` — extend or build new?

`/stack-update` today wraps `corpus/ops/update_guide.md` for `stack-dev/` — it runs `make pull` (default branch,
moving target), `make migrate`, `make up`. It does **none** of the three requirements above. Its skill doc
explicitly says *"Demo stacks aren't updated in place."*

**Decision: build new, keep `/stack-update` as-is.** Reasoning:

| Dimension | `/stack-update` (existing, dev-only) | `/demo-update` (this spec) |
|---|---|---|
| Target | `stack-dev/` (`dev-N`) | `stack-demo/` (`demo-N`) |
| Source model | `make pull` on default branch (moving) | Per-N clones @ pinned semver `v*` tag |
| Rebuild path | `make up --build` (dev docker-compose) | `up-injected.sh`'s inject-then-`docker compose -p demo-N ... up -d --build`, needs Clerkenstein re-inject after checkout |
| Migrations | `make migrate` (host-port Postgres) | `migrate-demo.sh N` (offset-port Postgres) |
| Data model | Dev is disposable-ish (`make reset-db` is an option) | Presenter state MUST survive — additive re-seed only; `--force-reset` is behind an operator gate |
| Verify | `curl -s localhost:8082/health` (a hint in the guide) | 3-tier MANDATORY fatal gate: autoverify → `test-platform live` → coverage sweep |
| Feature parity | Not a concern (dev tracks `main`) | Load-bearing: coverage-manifest as declarative catalog; sweep detects new-empty |
| Synthetic-data reuse | None | `stacksnap replay` + `stackseed --additive` + optional `gen-batch` + section→seeder routing |

Extending `/stack-update` would either bloat the dev skill with demo-only flags (`--demo`, `--force-reset`,
`--with-coverage`, `--with-playthroughs`, `--rollback`, `--reprovision`) or force operators to think in "same
skill, two very different contracts" mode — both are worse than a peer skill.

### 1.3 Scope

**In scope:** the CLI verb + `update-injected.sh`, the ref-resolution + rebuild + migrate flow, the
synthetic-data invocation glue, the 3-tier fatal verify gate, rollback, logs/provenance, the operator skill.

**Out of scope:** replacing `/stack-update`; touching `dev-N` at all; anything that requires a platform-repo
edit; re-provisioning secrets (that stays `/stack-secrets`'s job — an opt-in `--reprovision` calls it, doesn't
duplicate it); any change to `demo-up` semantics.

---

## 2. The three requirements — how each is served

### 2.1 R1 — Feature parity (new prod sections come in automatically)

**Mechanism, already in-tree:** `stack-verify/e2e/lib/coverage-manifest.ts` is the **declarative section catalog**
— the per-page list of expected UI sections. `coverage-manifest.unit.spec.ts` asserts every declared page has ≥1
section. `run-coverage.sh` drives a Playwright sweep (employee + manager vantage) that reports per-section
presence, persona consistency, and prod-eject escapes.

**How `demo-update` uses it:**

1. **Refresh the code.** For each per-N service repo, `git fetch --tags` + `git checkout <resolved-ref>`. Rebuild
   the injected + frontend images. This alone brings any new prod page/section into the running demo image.
2. **Refresh the rext pin** (opt-in `--bump-rext`). Because the coverage-manifest ships in
   `rosetta-extensions/`, updating `.agentspace/rext.tag` moves the section catalog forward. A new prod page's
   *declaration* and its per-seeder wiring land in the same rext release — so the two arrive together.
3. **Coverage sweep** (§4, R3) reports **present-but-empty** declared sections — the exact "new section came up
   empty" signal, distinguishing `unimplemented` vs `unimplementable-without-platform-edit` from `failing`.

**Honest boundary:** feature parity is not "invented at runtime by scraping prod." It's *"the manifest is the
source of truth, and it ships with the rext tag."* Any undeclared prod section that only exists in the
freshly-pulled Next.js will render but won't be gated — same fidelity contract as v2.0 Playthroughs and v1.10
M42 coverage.

### 2.2 R2 — Synthetic-data reuse for new sections

**Inventory of already-in-tree tooling `demo-update` will orchestrate:**

| Tool | Section | What it produces | When `demo-update` calls it |
|---|---|---|---|
| `stacksnap replay --surface taxonomy` | `stack-snapshot/` | Public taxonomy into `public.*` | After migrate — schema may have moved; replay is idempotent (TRUNCATE-then-reload) |
| `stacksnap replay --surface directus` | `stack-snapshot/` | Directus content into per-stack Directus | Same (only when the demo is `--local-content`) |
| `stackseed --preset <preset> --additive` | `stack-seeding/` | Structural fan-out via the 40+ seeders | After snapshot replay — **without** `--reset`, so existing state is preserved where the seeder is additive-safe |
| `stackseed --only <seeders>` | `stack-seeding/` | Subset invocation | When section→seeder routing (R2 loop) fires for a specific empty section |
| `gen-batch --max-cost-usd <cap>` | `stack-seeding/blueprint/` | AI-generated profile density (gpt-4o-mini) w/ `.batchcache/` hash-cache — $0 if unchanged | Only when a new preset lands or a batch descriptor changed |
| `datadna measure-closure` + `measure-snapshot` | `stack-seeding/cmd/datadna` | Closure gate + snapshot fidelity gate | Post-seed gate before R3 verify |

**Section→seeder routing** — the one net-new artifact. A checked-in
`rosetta-extensions/demo-stack/update-routing.yaml`:

```yaml
# When the R3 coverage sweep reports a declared section empty, invoke the matching seeder subset.
# Every section in coverage-manifest.ts must have a routing entry OR be waived (with reason).
sections:
  ai-readiness-funnel:
    seeders: [ai_readiness_config, ai_readiness_funnel]
    preset: gen-batch-org-fill        # optional
    notes: closed-cycle frozen-snapshot funnel; needs org_settings first
  members-directory:
    seeders: [org, roster, membership_skills, avatar, profile]
  skill-spotlight:
    seeders: [persona, taxonomyref, skillref_named, activity]
  # ...
waived:
  - section: presenter-cockpit-help
    reason: static UI copy; no data
```

A layout test (`TestSectionSeederRoutingCoverage`) enforces: every section in `coverage-manifest.ts` appears in
either `sections:` or `waived:` — keeps the table honest as new sections land.

**Data preservation is load-bearing here:** `demo-update` is **not** `demo-down + demo-up`. Presenter state
survives. `--reset` requires **both** `--force-reset` and `--force-reset-confirm=demo-N` (double-gate + the
existing seeder N=0 guard).

### 2.3 R3 — MANDATORY fatal verify gate (3 tiers)

Reuse `autoverify.sh`, `test-platform`, `run-coverage.sh`. Compose them as a **fatal** gate — inverting the
non-fatal contract that the bring-up tail uses. Rationale: the bring-up tail is non-fatal because a verify bug
must not block a good bring-up; an **update** is fatal because regression detection is its whole point.

| Tier | Probe | Fatal? | On failure |
|---|---|---|---|
| **T1 cheap-wins** | `autoverify.sh --project demo-N`: `curl -fsS :$((8082+OFFSET))/api/health` + `sentinel.casbin_rules > 0` + directus collections > 0 + no-prod-read DSN assert | Fatal | Abort, print exact re-run + diff-vs-pre-update state |
| **T2 scoped live** | `stack-verify/reports/generate.sh live` (offset/project/scope-aware liveness+readiness probe set) | Fatal | Same |
| **T3 semantic coverage** | `run-coverage.sh N employee` then `manager`. Parse `coverage-report.json`: `failingSections>0` OR `personaFailures>0` OR `escapes>0` → fail. `notReachedPages>0` → warn. Present-but-empty declared section → warn + emit routing suggestion (R2 loop) | Fatal (with escape) | Abort, print sections + suggested `stackseed --only <seeders>` |

**Escape valve:** `--no-verify` skips T3 only (T1 + T2 stay fatal always — an update that leaves the stack
dead is never acceptable). `--with-playthroughs` opts into a T4 (expensive; `rext playthroughs/` employee +
manager runs) — off by default.

**On failure**, the demo is *left running* so the operator can inspect. `--rollback` (see §3.9) is offered as
next-step advice.

---

## 3. Phases — the update pipeline

Named `Phase 0..9` matching the milestone breakdown in §5.

### 3.0 Pre-flight & lock

- Verify `demo-N` exists in the unified registry and is `up`.
- Acquire `.agentspace/updates/op_<ts>_demo-<N>.lock` (fcntl) — two concurrent updates cannot race.
- Snapshot pre-update state to `.agentspace/updates/op_<ts>_demo-<N>-pre.json`:
  - per-service current git SHA + tag (from each `stack-demo/demo-N/repos/<svc>/.git`),
  - `docker compose -p demo-N ps` output,
  - a hash of `coverage-manifest.ts` (to detect R1 catalog moves at Phase 8).
- Print the resolved plan; require `--yes` or `--dry-run`.

### 3.1 rext pin refresh (opt-in `--bump-rext[=<tag>]`)

Update `.agentspace/rext.tag` (default: highest semver `v*` in `rosetta-extensions`), re-pull the per-stack
consumption copy at `stack-demo/rosetta-extensions/`. Re-build any Go binaries the pipeline calls (`stacksnap`,
`stackseed`, `datadna`, `verify`). This is where new `coverage-manifest.ts` + new seeders + new routing entries
arrive together.

### 3.2 Per-service ref resolution

Reuse `demo-stack/lib/clone_repos.py:resolve_ref(repo, source, caller_refs)` which already implements the
`--ref app=v1.2.0,cms=main` syntax + `pick_latest(tags)` semver picker. Extend with `--resolve-only` for
dry-run print. **Pin by SHA, not tag name**, in the update log so a mid-update tag move can't drift the record.

### 3.3 Code refresh + Clerkenstein re-inject

For each service in the resolved plan:

```bash
cd stack-demo/demo-N/repos/<svc>
git fetch --tags origin
git checkout <resolved-sha>              # detached HEAD; pinned by SHA
# Re-apply Clerkenstein — checkout resets the working tree
$STK/rosetta-extensions/clerkenstein/apply-authn.sh <svc>
# Re-do the COPY vendor-colony perl edit up-injected.sh applies for injected svcs
```

The injected set is unchanged from `up-injected.sh`: `INJECT_SVCS="app cms jobsimulation skillpath"` (post-v2.1
skiller-in-app merge). Frontends (`next-web-app`, `studio-desk`) do not need injection but do need rebuild
with the same offset-baked build-args from `.env.demo-N`.

### 3.4 Rebuild + rolling restart

`docker compose -p demo-N -f $PLAT/docker-compose.yml -f $STK/rosetta-extensions/demo-stack/docker-compose.demo.yml
--env-file $PLAT/.env --env-file .env.demo-N --profile graphql up -d --build` — **service-by-service in
dependency order** from `corpus/architecture/dependency_map.md`. Wait for each service's healthcheck between
starts so the demo stays responsive throughout the update (the presenter demo doesn't fully go down).

Frontend rebuilds (`demo-N-next-web`, `demo-N-studio-desk`) are the biggest single time cost — ~3 min cached
per new tag. The rolling restart hides some of it but a visible restart-blip is expected on the next-web-app +
studio-desk cutover.

### 3.5 Migrate

`migrate-demo.sh N` — offset-port-aware Atlas migrations for the injected services.

**Warning printed BEFORE Phase 5:** *"Migrations may be non-reversible. If Phase 8 fails, a code rollback does
NOT roll back the schema. Consider `--snapshot-db-first` to `pg_dump` the offset-port Postgres to
`.agentspace/updates/pgdump-<ts>.sql.gz` first."*

### 3.6 Snapshot replay (idempotent)

```bash
stacksnap replay --surface taxonomy --stack demo-N
stacksnap replay --surface directus --stack demo-N    # only if the stack is --local-content
```

Both are idempotent (TRUNCATE-then-reload for taxonomy; the same directus-cache-first replay for content). Exit
codes 4/5 are handled by falling through to a targeted warning (missing/empty target schema after migration is
a bug; missing cached snapshot means the operator needs a fresh capture).

### 3.7 Additive seed

`stackseed --preset <recorded-preset> --additive --stack demo-N`. The preset is the one recorded in the demo's
own manifest (defaulting to the M52 consolidated `seed-generation-manifest.yaml`). `--additive` is the
newly-declared flag: **no TRUNCATE, no delete, no `--reset` semantics**; only INSERT-if-absent + UPSERT-safe
seeders run. Seeders that are not additive-safe (destructive fan-outs) skip themselves in `--additive` mode and
print a note.

`--force-reset` is available (destructive re-seed) but double-gated: requires `--force-reset` AND
`--force-reset-confirm=demo-N` on the CLI, and still refuses at N=0.

Optional `gen-batch --max-cost-usd <cap>` if a new preset landed or the mother prompt moved (cache-hits are $0;
the cap is re-asserted from the manifest — refuses to proceed without `--reconfirm-max-cost` if it wasn't set
in this invocation).

Post-seed: `datadna measure-closure --stack demo-N` gate (fail = abort update, roll forward to Phase 8 warning
but hold the gate).

### 3.8 Verify (fatal — see §2.3)

T1 → T2 → T3, each fatal, `--no-verify` skips T3 only. Optional `--with-playthroughs` runs T4.

### 3.9 Record + rollback surface

- Append to `stack-demo/updates.log.json`: from-SHA → to-SHA per service, coverage delta (pre vs post
  `coverage-report.json`: sections passing, failing, notReached), duration, verify outcome, operator, timestamp.
- Update `stack-demo/clones.lock.json` in place (SHA + tag per service, matching the ref-resolution semantics
  already recorded there by `ensure-clones.sh`).
- Emit a `.agentspace/updates/op_<ts>_demo-<N>-post.json` mirror of the pre-update snapshot.
- `rosetta-demo update N --rollback [op_<ts>]` re-checks out the pre-update SHAs, re-applies injection,
  rebuilds. **The schema is not rolled back** — an operator needing a full pre-update state must
  `--snapshot-db-first` and manually restore. This limitation is printed at every rollback invocation.

---

## 4. Principles

> The load-bearing contract. A `demo-update` change that violates a principle is wrong even if it "works."

- **P1 — Data preservation is the default.** The demo's Postgres data survives every update path. `--reset`
  semantics are behind a double-gate + N=0 guard. Any seeder invoked in the update loop must be additive-safe
  or must skip in `--additive` mode.
- **P2 — Prod-latest is the default; SHA-pinning is the record.** Ref *resolution* uses tags (semver); ref
  *recording* uses SHAs (immutable). A mid-update tag move never rewrites history.
- **P3 — Reuse, don't rewrite.** Every synthetic-data producer and every verify probe is already in-tree. The
  update pipeline is *glue*. The only net-new artifacts are `update-injected.sh`, `update-routing.yaml`, and
  the layout test that keeps the routing table honest.
- **P4 — Fatal-by-default verify.** T1 and T2 always fatal. T3 fatal, escape only via explicit `--no-verify`.
  The update's job is to detect regressions.
- **P5 — Zero platform-repo edits.** Same as the rest of rosetta-extensions. `up-injected.sh` establishes this
  boundary; `update-injected.sh` inherits it.
- **P6 — Values-blind for secrets.** `demo-update` does not read `.env` values, does not log them, does not
  re-provision them. `--reprovision` delegates to `/stack-secrets` — never duplicates its logic.
- **P7 — Presenter-visible cost.** Any tool that spends money (`gen-batch`) enforces `--max-cost-usd` from
  the manifest AND requires an in-invocation `--reconfirm-max-cost` before spending. Cache-hits at $0 do not
  need re-confirmation.
- **P8 — The pipeline is auditable.** Every update writes a `pre.json` + `post.json` + a `updates.log.json`
  row. Rollback is a first-class verb, not a "figure it out with git."

---

## 5. Milestone breakdown

Five milestones, each closing on a specific green signal. Numbered A → E for spec-time; will be renumbered to
the active release's `Mxyy` scheme when this spec is handed to `/developer-kit:design-roadmap`.

### M-A — Code refresh + rolling rebuild + migrate (Phases 0–5)

**Scope:** `update-injected.sh` end-to-end for Phases 0..5. Ref resolution + `--dry-run` + `--yes` gate.
Clerkenstein re-inject. Rolling rebuild. `migrate-demo.sh` call. **No data or verify yet.**

**Close-on-gate:** on a spare `demo-N` on Ithaca, run `demo-update N` from a deliberately-behind pin. After
Phase 5 completes, `test-platform N live` is GREEN. The stack answers on its offset ports at the new SHAs
(shown by `docker inspect` + `git rev-parse HEAD` in each per-N repo).

**Tests:**
- Unit: ref-resolver already covered in `demo-stack/lib/clone_repos_test.py`; add an update-plan builder unit
  test.
- Integration: a `demo-update-dry-run` integration test that spins the CLI with `--dry-run` on a fake
  registry entry + asserts the planned resolution matches the expected SHAs.
- Manual: the Ithaca run above (its output pasted into the PR).

### M-B — Data refresh (Phases 6–7)

**Scope:** `stacksnap replay` (both surfaces, guarded by `--local-content` for directus). `stackseed --additive`
including the new flag. `datadna measure-closure` gate. `--force-reset` double-gate + N=0 refusal test.

**Close-on-gate:** on the same demo-N, post-Phase-7 `datadna measure-closure --stack demo-N` is GREEN. Rows
count unchanged for additive-safe seeders; new seeder outputs (if the rext bump added any) are visible.

**Tests:**
- Unit: `--additive` flag parsing + per-seeder skip-logic dispatch (mock seeder registry).
- Integration: an integration test that runs the full data-refresh phase against a Postgres testcontainer
  and asserts pre-vs-post row counts for a fixed seeder set.

### M-C — Fatal verify gate (Phase 8)

**Scope:** T1 + T2 + T3 wiring, `coverage-report.json` parser, fatal-vs-warn classification, `--no-verify` +
`--with-playthroughs` flags. Failure paths leave the stack up and print the diagnostic hints.

**Close-on-gate:** on the same demo-N, `demo-update N` from a slightly-behind pin completes end-to-end with a
coverage sweep GREEN. Then, on a deliberately-broken pin (a known-bad prior tag), the same command **fails at
T3** and prints the failing-section list + the suggested `stackseed --only <seeders>` remediation.

**Tests:**
- Unit: coverage-report.json parser + fatal/warn thresholds (fixture-based; fixtures live under
  `demo-stack/tests/fixtures/coverage-*.json`).
- Integration: a Playwright-fixture-driven test that feeds a synthetic failing report through the verify gate
  and asserts exit code + printed remediation.

### M-D — Section→seeder routing (feature-parity plumbing)

**Scope:** `update-routing.yaml` schema + parser, `TestSectionSeederRoutingCoverage` layout test (every
`coverage-manifest.ts` section is routed or waived). When T3 emits a present-but-empty section, the update
optionally auto-invokes the matching seeder subset (opt-in `--auto-route`) and re-runs T3 once.

**Close-on-gate:** the layout test passes with the initial routing table covering every current
coverage-manifest section. A follow-up test: adding a new section to `coverage-manifest.ts` without adding a
routing entry FAILS the layout test.

**Tests:**
- Unit: routing YAML parser + auto-route dispatch (mock seeder invocation).
- Layout: `TestSectionSeederRoutingCoverage` (deterministic; runs in the fast tier).

### M-E — Rollback + logs + docs + `/demo-update` skill

**Scope:** `--rollback [op_id]` verb, `updates.log.json` writer, `clones.lock.json` update, optional
`--snapshot-db-first`, `corpus/ops/demo-update.md` (the operator guide), `.claude/skills/demo-update/SKILL.md`,
CLAUDE.md skill-table entry, `interconnected-documentation` list update.

**Close-on-gate:** after a passing update, `rosetta-demo update N --rollback` re-checks out the pre-update
SHAs, rebuilds, and the stack is `test-platform N live` GREEN at the old pin. `updates.log.json` has two rows
(forward + rollback). The skill invokes from `/demo-update N` and completes the same run.

**Tests:**
- Unit: log writer + rollback plan builder.
- Manual: the rollback e2e on Ithaca (pasted into the PR).

---

## 6. Non-goals

- Editing any platform repo.
- Modifying `/stack-update` (kept as-is for dev).
- Full schema rollback (documented limitation; `--snapshot-db-first` is the workaround).
- Continuous updates / a cron loop (out-of-scope; may be added later if `demo-update` proves stable).
- Multi-`N` batch updates (`demo-update N` operates on exactly one `N`).
- Update from `demo-down` state (pre-flight requires `up`; deliberate).
- Replacing `/demo-up` (a fresh `demo-N` still starts with `/demo-up`).

---

## 7. Honest boundaries + open questions

- **Non-reversible migrations.** Documented limitation. `--snapshot-db-first` is the presenter's insurance
  policy.
- **Frontend rebuild is ~3 min per new tag** (cached Dockerfile.dev). The rolling restart hides part of it.
- **Section→seeder routing** will drift as prod sections evolve. The layout test is the answer; whether it
  needs to run on a schedule vs only on rext bump is open (**Q1**).
- **Reference resolution edge case:** a service tagging a new `v*` between Phase 2 and Phase 3 would drift.
  P2 (SHA pin the record) solves the *record*; whether we snapshot-then-resolve or resolve-then-freeze at
  Phase 2 is open (**Q2**).
- **`gen-batch` cost enforcement across cache invalidation.** When a rext bump busts the batch cache, the
  `--max-cost-usd` re-check must not surprise the operator. Whether `--reconfirm-max-cost` should be
  auto-required by the mother-prompt digest change is open (**Q3**).
- **Multiple-in-flight lockfile granularity.** Per-N lock is clear; whether cross-N updates on the same box
  need a coarser lock (e.g. shared Docker network / port collisions during rolling restart) is open (**Q4**).
- **Rollback semantics when M-A failed mid-Phase-4.** A partial rolling restart leaves the stack in a mixed
  state. The Phase-4 loop must be checkpointed per-service so rollback is deterministic (**Q5**).

Open questions are logged in [`spec-progress.md`](spec-progress.md); parked items in
[`next-release.md`](next-release.md).

---

## 8. References

- `corpus/ops/update_guide.md` — `/stack-update`'s dev guide (peer, unchanged).
- `corpus/ops/rosetta_demo.md` — the demo-stack lifecycle.
- `corpus/ops/idempotency.md` — the re-run safety contract that `stacksnap replay` + `stackseed` already honor.
- `corpus/ops/verification.md` — the bring-up autoverify (non-fatal); `demo-update` inverts to fatal.
- `corpus/ops/demo/coverage-protocol.md` — the semantic coverage gate; `demo-update`'s T3 is this, run
  post-update.
- `corpus/ops/demo/playthroughs.md` — the optional T4.
- `.claude/skills/stack-update/SKILL.md` — the peer skill; `demo-update` shares the tone but not the contract.
- `stack-demo/rosetta-extensions/demo-stack/up-injected.sh` — the reference for `update-injected.sh`.
- `stack-demo/rosetta-extensions/demo-stack/lib/clone_repos.py` — `resolve_ref` + `pick_latest` reused verbatim.
- `stack-demo/rosetta-extensions/stack-verify/live/autoverify.sh` — T1 reused verbatim.
- `stack-demo/rosetta-extensions/stack-verify/e2e/run-coverage.sh` — T3 reused verbatim.
