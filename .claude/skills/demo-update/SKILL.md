---
name: demo-update
description: Refresh a running demo-N to prod-latest IN PLACE — pull new SHAs on each injected repo, re-inject Clerkenstein, rolling-rebuild, migrate, replay snapshot, additive seed, then fatal 3-tier verify (with section→seeder routing that fills newly-present-but-empty sections). Data survives; presenter cockpit + login stay live; --rollback re-checks out the pre-update SHAs. Use when asked to update / refresh / re-ground a demo without teardown-reseed.
argument-hint: N [--ref svc=ref ...] [--dry-run] [--yes] [--no-migrate] [--no-snapshot] [--no-seed] [--force-reset] [--seed PATH] [--no-verify] [--no-routing] [--with-playthroughs] [--with-ui] [--refresh-rext-tag] [--rollback [op_id]] [--snapshot-db-first]
---

# Demo Update — refresh a running demo-N in place (no teardown-reseed)

Refreshes `demo-N` to **prod-latest** IN PLACE — pulls new SHAs on each injected repo (default: highest
semver `v*` tag per repo — same rule `up-injected.sh` uses; per-repo `--ref svc=ref` overrides supported),
re-applies Clerkenstein, rolling-rebuilds, migrates, replays the taxonomy snapshot, additive-seeds, then runs
the **fatal 3-tier verify gate** (T1 autoverify + T2 test-platform live + T3 coverage sweep with
**section→seeder routing** for newly-present-but-empty sections). Data survives every update path; the
demo stays up on the same offset ports; a presenter's browser session stays logged in. Source of truth:
[`corpus/ops/demo-update.md`](../../../corpus/ops/demo-update.md) (the operator guide) +
[`corpus/ops/rosetta_demo.md`](../../../corpus/ops/rosetta_demo.md) (the demo lifecycle).

> **Why a new verb (not extending `/stack-update`)** — `/stack-update` targets the **dev** side (pull `main`,
> in-place rebuild, no verify gate, no reseed contract). `/demo-update` targets a **demo** stack: source is
> prod-latest tags (not `main`), verify is **fatal**, data preservation is a **contract** (not a side-effect),
> and adaptation to new prod sections needs the section→seeder routing that only `/demo-update` carries.
> Different target, different source, different bar — a different verb.

## Mission

1. **Read the guides** — [`corpus/ops/demo-update.md`](../../../corpus/ops/demo-update.md) (the phase pipeline
   0..8 + the routing table + rollback surface) + [`corpus/ops/rosetta_demo.md`](../../../corpus/ops/rosetta_demo.md)
   (the lifecycle + registry + port-offset context this operates within).

2. **Bring it up** via the demo tooling — consumed at the pin in `.agentspace/rext.tag`, exactly like `/demo-up`.
   ```bash
   DEMO=stack-demo/rosetta-extensions/demo-stack
   # Refresh a running demo-N to prod-latest (auto-detect refs, migrate, replay, additive seed, fatal verify).
   "$DEMO/rosetta-demo" update N
   # Dry-run first (prints CURRENT → NEXT plan; no state change).
   "$DEMO/rosetta-demo" update N --dry-run
   # Pin one service ref (e.g. debug a specific tag on app, take latest for the rest).
   "$DEMO/rosetta-demo" update N --ref app=v2.1.3
   # Rollback: re-checkout the pre-update SHAs of the most-recent forward update on demo-N.
   "$DEMO/rosetta-demo" update N --rollback
   # Rollback to a specific op + take a pg_dump BEFORE touching code (presenter's insurance).
   "$DEMO/rosetta-demo" update N --rollback op_20260710-153012 --snapshot-db-first
   ```

3. **The fatal 3-tier verify gate** — after Phase 7b, the update runs:
   - **T1 autoverify** — cheap-win backend `/api/health` 200 + `sentinel.casbin_rules > 0` (silent-403 catcher). **Always fatal.**
   - **T2 test-platform live** — `verify.sh` scoped to the RUNNING compose services on the offset ports. **Always fatal.**
   - **T3 coverage sweep** — Playwright coverage runner on the seeded heroes; if it reports newly-present-but-empty sections, `/demo-update` consults `demo-stack/update-routing.yaml` (the section→seeder routing) and applies **ONE remediation pass** (additive `stackseed --seed <preset>` and/or `stacksnap replay --surface <surface>`, dedup'd), then re-sweeps. Pass-2 verdict is fatal. **`--no-verify` skips T3 only.**

   On any fatal failure the **demo stays UP** — the operator can inspect. `--rollback` re-applies the pre-update
   SHAs; the schema is **not** rolled back (documented limitation printed at every rollback invocation —
   `--snapshot-db-first` is the presenter's escape valve).

4. **Data preservation contract** — the default update path is **always** additive: seed is `additive=true`
   (`stackseed --stack demo-N --seed <preset>` — NEVER `--reset` in the default flow); snapshot replay is
   TRUNCATE-then-reload on **catalog surfaces only** (public taxonomy + Directus content), never tenant data;
   T3 remediation is additive. `--force-reset` is **double-gated** (`--force-reset` + `--force-reset-confirm=demo-N`)
   and **N=0 hard-refused** (would blow away the dev stack).

## Safety

Every op is `-p demo-N`-scoped; the tooling hard-refuses the dev project. The presenter cockpit + FontAwesome
CDN + already-authenticated browser tabs stay live (rolling restart, not full down/up). Data survives every
update path unless the operator explicitly double-gates `--force-reset`. The section→seeder routing is honest
(unroutable sections are surfaced in the T3 error output, never silently skipped). Rollback re-checks pre-update
SHAs from the recorded op record — no fuzzy resolution. See [`corpus/ops/safety.md`](../../../corpus/ops/safety.md)
for the family-wide contract.

## After update

The demo is on the **same offset ports** (no re-log-in). Log in → verify the seeded stories still play through
end-to-end; open the presenter cockpit (`:7700+offset`) to confirm the seat-switch still works. The op record
is at `stacks/demo-N/updates/<op_id>.json` (full detail) and one row is appended to
`stacks/demo-N/updates.log.json` (the chronological index — cheap to scan; used by `--rollback` when resolving
`<latest>`). If T3 remediation ran, the record's `verify.t3_remediation` block lists which seed presets +
snapshot surfaces were applied.

## Related skills

| Skill | Use when |
|-------|----------|
| `/demo-up` · `/demo-down` | Bring a demo up / tear it down (the disposable lifecycle — `/demo-update` refreshes an existing one) |
| `/stack-update` | Update the **dev** side (targets `main`; in-place; no fatal verify — a different bar) |
| `/stack-snapshot` · `/stack-seed` | Individual set-dress / seed passes (`/demo-update` composes both under a verify gate) |
| `/test-platform` | Run the same probes T1/T2 use, ad-hoc, against a running stack |
| `/stack-list` | See live stacks + their offset ports |
