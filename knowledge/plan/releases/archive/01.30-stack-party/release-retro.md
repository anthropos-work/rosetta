# Release Retro: v1.3 "stack party"

**Shipped:** 2026-06-07 · tag `v1.3` · **Milestones:** M12 → M13 → M14 → M15 (all `section`)
**Theme:** the dev/demo convergence — dev stacks become first-class peers (unified first-available-N registry; a local per-stack Directus + auto-snapshot + a light `dev-min` seed on build), one generic `stack-*` skill set (hard-renamed, no aliases), and a single code-cited safety & security doc. Two stack kinds, one operating model.

## What shipped
- **M12** — the unified dev+demo stack registry (`stack-core/stack_registry.py`) + first-available-N allocator (fcntl.flock + atomic write + docker-ps adds-only reconcile). One shared N-pool; `dev-N`/`demo-N` can never collide on ports. The first non-Go milestone surface (Python + bash).
- **M13** — dev as a full-fidelity peer: a `dev-stack up` gets the demo treatment by default — a per-stack Directus (via the new executable `stack-snapshot/cmd/provision-plan` runner that makes M10's library-only contract a CLI), a cache-first `stacksnap replay` of the public taxonomy+directus, and the new `dev-min` light seed. Capture never runs on dev (replay-only); media refs-only; the n=0-dev guard doubled.
- **M14** — the operation skills converged onto both stack kinds via a **hard rename, no aliases**: `dev-up`/`dev-down` + `stack-list`/`stack-seed`/`stack-snapshot`/`stack-update`; the 6 old dirs removed; full reference sweep; `TestDocSourceSkillRename_M14` pins the rename so it can't rot.
- **M15** — the net-new code-cited `corpus/ops/safety.md` (read-side: never reads private data; write-side: never touches prod) + a dual-repo KB refresh; 7 fail-closed docs↔code drift guards pin every load-bearing safety claim to source.

## Incidents (P1+) across the release
- **None shipped.** Zero P0/P1, zero regressions, zero flakes across all 4 milestones.
- **3 robustness bugs caught + fixed in M13 harden** (all in M13's own new code, each mutation-pinned): a whitespace-only Directus env slipping the prod-safe firewall (fixed fail-closed); the set-dress re-run hint always appending `--no-snapshot`; a `set -u` trailing-flag error leaking a raw shell message. None reached a merge.
- **1 real correctness bug caught in M13 build review**: `build_cli` resolved the cmd package one dir too high (hermetic stub tests couldn't catch it) — fixed + regression-tested with a real-build test.

## Cross-milestone patterns
- **The "new deliverable shipped without its index row" miss recurred — now at the corpus level.** M10 and M13 each closed with a per-unit handbook-index miss (a new package/runner shipped without its README Packages-table row). At v1.3 close-release the same class surfaced one level up: the M15 headline `safety.md` shipped absent from both corpus directory index tables (`corpus/ops/README.md` + `corpus/README.md`). **Process nudge for v1.4:** extend the "index-row in the same commit as the deliverable" check from per-unit extension READMEs to **corpus directory READMEs** too.
- **Deferral-record discipline drifted on a forward-moved item.** DEF-M10-01's destination moved v1.2(→v1.3)→v1.4, but it was never re-formalized as a `RELEASE-SCOPE-DEFER` escape-hatch when it moved (it had been recorded as Fate-2), and the per-milestone audits called it "no repeat" when by the audit skill's own definition a destination-updated-forward item is a `DRIFT_DEFER`. Caught + corrected at close-release (formal signed decision + classification + date fix). **Nudge:** when a deferral's destination moves forward, re-run the escape-hatch formalization, don't carry the old fate.
- **Docs/rename milestones reward contract guards over test counts.** M14 (+1 test) and M15 (+7 drift guards) added little raw coverage but high assurance — the value was the reference-integrity + docs↔code drift guards. The pattern (turn a silent-rot failure mode into a test failure) is now the house style for documentation-shaped milestones.

## Carry-forward → v1.4 (in `roadmap-vision.md § v1.4 seeds`)
- **DEF-M10-01** — S3 media blob **bytes** + the cloud `SnapshotStore` backend. The sole release escape-hatch; formally signed → v1.4 at close-release (gated on eu-west-1 S3-read access the project doesn't have; v1.3 ships media refs-only by design). `RELEASE-SCOPE-DEFER` recorded in `m15-safety-doc/decisions.md`.
- The other v1.4 seeds (AI-generated rich content, external stack shareability, more mirror engines, the deployment/injection CI gate) were staged to v1.4 when the user scoped v1.3 = "stack party" (2026-06-07).

## Metrics delta (vs v1.2)
- **Go test funcs 693 → 713** (+20, test-only matcher; stack-snapshot +15, stack-seeding +5; alignment + clerkenstein unchanged). **No regression.**
- **Python +52 net new** (M12's unified-registry suite — `test_stack_registry.py` 0→50 + the ERR-trap-frees paths); 174 Python total.
- Coverage flat/up on continuing logic packages (`stack_registry.py` 98%, `cmd/provision-plan` 93%, `directus`/`blueprint` 100%); the M15 drift guards mutation-proven fail-closed.
- Flake **0**; triple-clean **3/3** (4 Go modules `-race -shuffle` + 4 Python suites); `-race`/gofmt/vet/shellcheck/py_compile all clean.
- Supply-chain GREEN: **0 called third-party CVEs**; the 12 called vulns are Go stdlib @go1.25.3 → go1.25.11+ toolchain (same as v1.2); all-permissive licenses.
- Clerkenstein alignment gates held **100%/100%** on all 4 surfaces (v1.3 touched no clerkenstein).

## Stats delta
- Release-close snapshot: `knowledge/journal/stats/2026-06-07.json`. Milestones **22/22 done** (M0–M15 across v1.0–v1.3). Full per-release detail: `releases/01.30-stack-party/metrics.json`; release history: `knowledge/plan/metrics-history.md`.
