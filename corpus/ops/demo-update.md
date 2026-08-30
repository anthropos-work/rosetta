# `/demo-update` — refresh a running demo-N to prod-latest in place

> **Scope:** the operator guide for `/demo-update N` — the update lifecycle verb for a **demo** stack.
> **Not** a spec (see [`knowledge/plan/spec-drafts/demo-update/spec.md`](../../knowledge/plan/spec-drafts/demo-update/spec.md)),
> **not** a `/stack-update` extension (see the "why a new verb" note in the skill file). This doc is the
> HOW: phases, flags, verify semantics, routing, rollback, examples.

## What it does

`/demo-update N` refreshes a **running** `demo-N` stack to **prod-latest** IN PLACE — without teardown-reseed,
without losing the presenter's login, without touching a read-only platform repo. It's the missing sibling of
`/demo-up` and `/demo-down`: the update-in-place lifecycle verb.

Peer relationship:

| Verb | Target | Source of "latest" | Data | Verify | Verb |
|------|--------|--------------------|------|--------|------|
| `/dev-up` / `/dev-down` | `dev-N` | `main` (working set) | Preserved (dev) | Bring-up autoverify (non-fatal) | `dev-up` / `dev-down` |
| `/stack-update` | `dev-N` (in-place) | `main` | Preserved (dev) | Bring-up autoverify (non-fatal) | `stack-update` |
| `/demo-up` / `/demo-down` | `demo-N` (disposable) | Highest semver `v*` tag | Reseeded each time | Bring-up autoverify (non-fatal) | `demo-up` / `demo-down` |
| **`/demo-update`** | **`demo-N` (in-place)** | **Highest semver `v*` tag (per-repo `--ref` override)** | **Preserved (contract)** | **Fatal 3-tier gate** | **`demo-update`** |

## Invocation

Uses the same tooling pin as `/demo-up` (the `.agentspace/rext.tag` single-source pin):

```bash
DEMO=stack-demo/rosetta-extensions/demo-stack

# The default flow (auto-detect refs, migrate, replay, additive seed, fatal verify).
"$DEMO/rosetta-demo" update N

# Dry-run — prints the CURRENT → NEXT plan (per-repo describe/SHA + how="tag|ref|current"). No state change.
"$DEMO/rosetta-demo" update N --dry-run

# Per-service ref override (debug a suspect tag; unspecified services still take latest).
"$DEMO/rosetta-demo" update N --ref app=v2.1.3 --ref cms=v2.1.4

# Rollback the most-recent forward update on demo-N (re-checkout pre-update SHAs).
"$DEMO/rosetta-demo" update N --rollback

# Rollback a specific op + pg_dump BEFORE touching code (presenter's escape valve).
"$DEMO/rosetta-demo" update N --rollback op_20260710-153012 --snapshot-db-first
```

## The phase pipeline (10 phases)

| Phase | Name | What it does | Skippable? |
|------:|------|--------------|-----------|
| 0 | Pre-flight + lock + `op_id` | Refuse missing stack / stopped demo; acquire per-N + box-wide fcntl lock; mint `op_YYYYMMDD-HHMMSS`. | No. |
| 1 | rext-pin | Pins the tooling tag consumed by THIS op (records to the record). `--refresh-rext-tag` re-fetches. Skipped on rollback. | Rollback. |
| 2 | Pre-snapshot + resolve plan + gate | Snapshot the CURRENT per-repo SHAs (`pre-<op_id>.json`); resolve the NEW ref per repo (highest `v*` tag or `--ref` override); print the CURRENT → NEXT diff; on `--dry-run` stop here. Rollback synthesises the plan from a prior op's `pre` block. | Rollback (plan is synthesised, not resolved). |
| 3 | Code refresh + inject | Per repo: fetch + checkout the resolved ref; re-apply Clerkenstein injection (`apply-authn.sh` + vendor-colony perl edit + demopatches). | No. |
| 4 | Rebuild + rolling restart | Per compose service (in `INJECT_SVCS`): `docker compose -p demo-N build <svc>` + `up -d --no-deps --force-recreate <svc>`. Non-injected services stay up (postgres/redis/directus/next-web/studio-desk/ant-academy). | No. |
| 5 | Migrate | Run atlas `migrate apply` on demo-N's Postgres via the injected compose service. | `--no-migrate`. |
| 6 | Snapshot replay | `stacksnap replay --stack demo-N --surface catalog|directus` — TRUNCATE-then-reload catalog surfaces only. Idempotent. | `--no-snapshot`. **Skipped on rollback.** |
| 7 | Additive seed | `stackseed --stack demo-N --seed <preset>` (or `--seed PATH`). **Never `--reset` in the default flow.** `--force-reset` is double-gated + N=0 refused. | `--no-seed`. **Skipped on rollback.** |
| 7b | Closure gate | `datadna measure-closure --stack demo-N` — fails if any fabricated `K-*` refs slipped through. Skipped on rollback. | Bundled with 7. |
| 8 | Fatal 3-tier verify | T1 autoverify + T2 test-platform live + T3 coverage sweep (with section→seeder routing). See below. `--with-playthroughs` opts into T4. | T3 only, via `--no-verify`. |
| 9 | Record | Persist the full op to `stacks/demo-N/updates/<op_id>.json`; append a compact row to `stacks/demo-N/updates.log.json` (the chronological index). | No. |

## The fatal 3-tier verify gate (Phase 8)

The whole point of `/demo-update` — a demo is "up" but is it **actually working after the update**?

### T1 autoverify — `phase8_t1_autoverify`

Bounded 20×1s readiness retry (tolerates Phase 5 migrate-restart), then two cheap-win asserts:
- `curl :OFFSET+8082/api/health` returns 200 (backend liveness).
- `SELECT count(*) FROM casbin_rule` on demo-N Postgres — must be `> 0` (the silent-403 catcher: Sentinel is
  the wall between a stack that "looks up" and a stack that *serves 401 on everything*).

**Always fatal.** Failure aborts the op; the demo stays UP so the operator can inspect.

### T2 test-platform live — `phase8_t2_scoped_live`

Enumerates the RUNNING compose services via `docker ps --filter label=com.docker.compose.project=demo-N` +
passes `STACK_SERVICES` to `verify.sh` so its M18 scope filter is honoured (a service that isn't running is
never asked to answer). Runs on the demo's offset ports.

**Always fatal.** Same accept-RED-as-validation pattern as T1: the gate is doing its job; a persistent RED
here is a substrate defect the operator fixes elsewhere, not something `/demo-update` autocorrects.

### T3 coverage sweep + section→seeder routing — `phase8_t3_coverage_sweep`

The M-D layer — the one novel piece of the verify gate.

1. **Pass 1** — `_t3_sweep_once` runs `run-coverage.sh` against demo-N + parses
   `coverage-report.json` for `failingSections + personaFailures + escapes`. Green on all three → return 0
   (fatal semantic honoured, no work to do).

2. **Pass 1 red + `--no-routing`/`--no-seed`** — fail FATAL immediately (no remediation attempt).

3. **Otherwise** — consult `demo-stack/lib/update_routing.py plan --sections '<json>'`, which reads
   `demo-stack/update-routing.yaml` (the declarative section→seeder table) and returns:
   - `seed_presets` — the dedup'd preset list (currently `("stories",)` for every routed section).
   - `snapshot_surfaces` — the dedup'd surface list.
   - `unroutable` — sections with `fix: none` (the two known no-routes are `studio-desk-home` +
     `ant-academy-home`, both cross-port SPAs outside the seeder + snapshot lane).

4. **Apply ONE remediation pass** (additive, never `--reset` / `--force`):
   ```bash
   stackseed --stack demo-N --seed <preset>          # per preset
   stacksnap replay --stack demo-N --surface <surf>  # per surface
   ```

5. **Pass 2** — re-sweep. Pass 2's verdict is the fatal verdict; the record's
   `verify.t3_remediation.second_pass_ok` field carries the outcome. Unroutable sections that were failing in
   Pass 1 surface in the error output (honest disclosure — never silently skip).

**`--no-verify` skips T3 only** (T1 + T2 stay fatal). **`--no-routing`** keeps T3 but skips the remediation
pass (Pass 1 is the verdict).

### T4 playthroughs — `phase8_t4_playthroughs` (`--with-playthroughs`)

Opt-in the full functional-flow e2e runner (see [`demo/playthroughs.md`](demo/playthroughs.md)). Fatal when
requested.

## The routing table (`demo-stack/update-routing.yaml`)

One row per section in `stack-verify/e2e/lib/coverage-manifest.ts`. **Bidirectional layout gate** in
`demo-stack/tests/test_update_routing.py`:

- Every manifest section must have a route (missing → CI fail).
- Every route must have a manifest section (orphan → CI fail).
- `fix: seed|snapshot-replay|none` — with `preset:` required on `seed`, `surface:` on `snapshot-replay`,
  and `rationale:` required on `none` (honest disclosure).
- Trigger: **rext-bump only** — the layout gate fires on any drift in either direction and CI naturally runs
  on the coverage-manifest bump. A schedule would only detect the same drift later, never earlier.

## Rollback (`--rollback [op_id]`)

**Model.** Rollback is a REGULAR update whose "resolved plan" is synthesised from a prior op's `pre` block:
each service's ref = its pre-update SHA, `how="rollback"`. This reuses phase 3/4/5/8 unchanged.

**Skipped phases:** phase 1 (rext pin — stays on current), phase 2 pre-snapshot + resolve (plan is
synthesised), phases 6/7/7b (data-refresh is meaningless — going BACK, the data is already there).

**Op resolution:**
- `--rollback` (bare) — the most-recent successful `demo-update` record on demo-N (rollbacks themselves
  excluded by kind).
- `--rollback op_YYYYMMDD-HHMMSS` — the explicit op (must exist in `stacks/demo-N/updates/`).

**Schema is NOT rolled back.** Documented limitation (spec §3.9). Atlas migrations applied by the forward
update stay in place. Printed at **every** rollback invocation — no silent data assumption.

**`--snapshot-db-first`** — the presenter's escape valve. Takes a `pg_dumpall` of demo-N's Postgres BEFORE
touching code + records the path into the op record. `pg_dumpall` failure aborts the rollback BEFORE any
change (never leave the stack half-rolled-back). Manual restore is the operator's job.

## The op record + updates log

- **Full detail** — `stacks/demo-N/updates/<op_id>.json`: `kind` (`demo-update` | `demo-update-rollback`),
  `milestone: "M-E"`, per-repo `pre`/`planned`/`post`, `data_refresh`, `verify` (T1..T4 summaries +
  `t3_remediation` block: `applied`, `seed_presets`, `snapshot_surfaces`, `unroutable_sections`,
  `second_pass_ok`), `rollback_of` + `db_snapshot_first` on rollback ops.
- **Chronological index** — `stacks/demo-N/updates.log.json`: a JSON array, one compact row per op (fcntl-safe
  append: decode → append → encode). Used by `--rollback` when resolving `<latest>`, and cheap to scan for a
  presenter's "what did I do last?".

## Flag matrix

| Flag | Effect |
|------|--------|
| `--ref svc=ref` | Per-service override (repeatable). Unspecified services take highest `v*` tag. |
| `--dry-run` | Stop after Phase 2 gate (print plan, no state change). |
| `--yes` / `-y` | Skip interactive gate confirmations. |
| `--no-migrate` | Skip Phase 5. |
| `--no-snapshot` | Skip Phase 6. |
| `--no-seed` | Skip Phase 7 + Phase 7b (closure needs the seed). Implicitly disables T3 remediation. |
| `--force-reset` | With `--force-reset-confirm=demo-N` — data-destroying reset. **N=0 refused.** |
| `--seed PATH` | Alternate seed manifest for Phase 7. |
| `--no-verify` | Skip T3 (T1 + T2 stay fatal). |
| `--no-routing` | Keep T3 sweep but skip the remediation pass (Pass 1 is the verdict). |
| `--with-playthroughs` | Opt into T4. |
| `--with-ui` | Also rebuild the UI-tier containers (next-web / studio-desk). |
| `--refresh-rext-tag` | Re-read `.agentspace/rext.tag` + re-fetch. |
| `--rollback [op_id]` | Rollback mode (see above). |
| `--snapshot-db-first` | pg_dumpall demo-N's Postgres BEFORE change (rollback insurance). |

## Non-goals

- **Continuous / scheduled `/demo-update`** — deferred (next release). Operators drive updates.
- **Multi-N batch update** — deferred.
- **Update from a torn-down demo** — no; `/demo-up` a new one instead. `/demo-update` is for RUNNING demos.
- **Full schema rollback** — deferred (limitation; `--snapshot-db-first` is the workaround).

## Related

- [`knowledge/plan/spec-drafts/demo-update/spec.md`](../../knowledge/plan/spec-drafts/demo-update/spec.md) — the spec (why + contract).
- [`corpus/ops/rosetta_demo.md`](rosetta_demo.md) — the demo lifecycle context.
- [`corpus/ops/safety.md`](safety.md) — the family-wide safety contract.
- [`corpus/ops/verification.md`](verification.md) — the shared verify contract T1/T2 build on.
- [`corpus/ops/demo/coverage-protocol.md`](demo/coverage-protocol.md) — the coverage sweep T3 uses.
- [`corpus/ops/demo/playthroughs.md`](demo/playthroughs.md) — the T4 opt-in.
