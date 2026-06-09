# M18 — Spec Notes

_Technical notes accumulated during build — mechanisms, file paths (with line cites), gotchas, and the concrete shape of each change. Populated by `/developer-kit:build-milestone`. The verified code locations from the design-time research are in the milestone `overview.md` and `.agentspace/demo-up-issue.md`._

## Code-base map (verified against live code, authoring copy `dress-rehearsal-m17`==`0d36251`)

### stack-verify (the toolkit to make offset/scope-aware)
- `stack-verify/lib/services.sh` — the `SERVICES` array: 12 rows `name|container|host:port|kind|target`,
  all hardcoded `anthropos-*-1` container names + BASE host ports (5432/6379/8087/8082/8085/8100/8400/8090/8300/10400/5050/3200).
  `service_rows()` emits trimmed rows; `probe_service()` does docker/tcp/http probes.
- `stack-verify/lib/readiness.sh` — deep probes; hardcodes `anthropos-postgresql-1`, `anthropos-redis-1`,
  and BASE ports `5050` (graphql), `3200` (gotenberg), `8087` (sentinel), `8301` (storage rpc).
- `stack-verify/live/verify.sh` — drives liveness (over `service_rows`) then readiness (named fn calls).
  No project/offset/scope awareness today.
- `stack-verify/repos/run.sh:~108` — **`$DEVDIR` bug**: `repo_dir="$DEVDIR/$repo"` but only `STACK_ROOT` is defined.
- `stack-verify/census/inventory.sh:~75` — same **`$DEVDIR` bug**: `dir="$DEVDIR/$repo"`.
  (Both: repos are siblings of `platform/` under the stack root, so the fix is `$STACK_ROOT/$repo`.)

### The registry (source of truth for offset/ports — read this, don't recompute)
- `stack-core/stack_registry.py` — unified dev+demo registry. Runtime file `stack-core/.stacks/registry.json`
  (gitignored), keyed `"<type>-<N>"`, rec = `{type,n,ports:[host ints],status,created}`. `set_ports()` records
  resolved host ports. CLI: `stack_registry.py list` → JSON rows; `--registry` overrides path; `$STACK_REGISTRY` env in CLIs.
- Offset derivation: a stack's recorded `ports` are `base + N*10000`. We read the registry record for the
  project and take its `n` (authoritative) — `offset = n*10000`. The `ports` array cross-checks it.

### Bring-up tails (where the cheap-win asserts + auto-verify wire in)
- `demo-stack/up-injected.sh` — the FULL injected demo path (the one ISSUE-12/14 exercised). `OFFSET=$((N*10000))`
  inline. Tail: `migrate-demo.sh "$N" || log "..."` then final `log "UP. ..."`. **Does NOT call set-ports/reg_set**
  (bypasses the unified registry) → the auto-verify here must pass project/offset explicitly (it knows N).
- `dev-stack/dev-stack` `cmd_up` — DOES `reg_cli set-ports dev "$n" ...`. Set-dressing block is the
  **default-on + non-fatal pattern to mirror**: `if ! dev-setdress.sh ...; then echo "⚠ ..."; fi`.
- `demo-stack/rosetta-demo` `cmd_up` — DOES `reg_cli set-ports demo "$n" ...` (the non-injected demo path).

## Design decisions (resolving the overview's open questions)

- **M18-D1 — offset source = the N the bring-up already knows, cross-checked by the registry's recorded ports.**
  `stack-verify` gains a `STACK_PROJECT` (e.g. `demo-1`/`dev-3`) + `STACK_OFFSET` env contract. The bring-up
  passes these explicitly at the tail (it allocated N, so it knows them — no drift). For *operator-driven* runs
  (`/test-platform N`), stack-verify resolves the offset from the registry record's `n` (authoritative), with the
  recorded `ports` as a sanity cross-check. **Never** recompute from a fragile formula alone — the registry is the record.
  This is the load-bearing correctness mitigation: a mis-derived offset would false-`down` a healthy stack (the very bug M18 fixes).

- **M18-D2 — scope filter via `STACK_SERVICES` (space-separated) ∩ the SERVICES array.** The bring-up CLIs already
  parse `--services`/`--profile`; pass the resolved service set through. Empty/unset = all-in-profile (no filtering).
  A row not in the requested set is simply not probed (not a false `down`).

- **M18-D3 — auto-run is default-on + NON-FATAL.** Mirror `dev-stack cmd_up`'s set-dress pattern exactly:
  `if ! verify ...; then echo "⚠ ... run /test-platform N to dig in"; fi`. A verify bug never blocks a good bring-up.

- **M18-D4 — cheap-win asserts are a thin, dependency-free pre-check at the bring-up tail** (the exact ISSUE-7
  catcher): `curl -fsS localhost:$((8082+OFFSET))/api/health` + `casbin_rules > 0` on the offset Postgres via
  `docker exec <project>-postgresql-1`. They live inside the auto-verify path so there's one non-fatal warning surface.

## Pre-flight audits — Offset/project awareness (first section)
- **Phase 0b KB-fidelity: GREEN** (2026-06-09, sha `f305dba`). Report: `kb-fidelity-audit.md`. All dependency
  topics PAIRED + ALIGNED; the net-new `corpus/ops/verification.md` is a pre-promoted `Delivers →` deliverable
  (blind area that does not block). 0 stale load-bearing claims. SEVERITY=clear.
- **Audit reuse:** this verdict covers all M18 sections (one subsystem: stack-verify + the two bring-up tails +
  the one rosetta doc). Later sections reuse it unless the KB contracts change.
