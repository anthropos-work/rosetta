# M26 — Decisions

_Implementation decisions with rationale, numbered `M26-D1`, `M26-D2`, … . Empty at scaffold; filled during build._

_Pre-decided at design (in `overview.md` as D-MAIN + D1–D6 — do NOT re-litigate):_
- _D-MAIN — PLAT (compose dir) MOVES to `stack-demo/platform` (full self-containment, not the keep-on-stack-dev variant)._
- _D1 — ensure-clones placement BELOW the `UP_INJECTED_LIB_ONLY` seam (~`up-injected.sh:172`), ABOVE the M28 pre-flight;
  bare call, no `|| true`._
- _D2 — add `DEMO="$DEMO_WS"` alias (keep both names)._
- _D3 — repoint all 4 demo scripts + port both conditional test files (no drops)._
- _D4 — ensure-clones phase-(b) `.env` copy is NON-FATAL/copy-if-present (the one refinement of the orphan; defers to M30)._
- _D5 — `reuse_dev_images` OFF by default; opt in via `DEMO_REUSE_DEV_IMAGES=1` → `--reuse-dev-images`._
- _D6 — pin the GUIDE test count to the live-recomputed GuideDocTruth-formula sum (~41), NOT the orphan's stale 40._

## M26-D4-impl — the D4 refinement CHANGES a ported test's assertion (not a verbatim port)
The orphan's `ensure-clones.sh` phase (b) **hard-`exit 1`'d** when BOTH `stack-demo/platform/.env` and
`stack-dev/platform/.env` were absent ("the single secrets source on this box is missing"). The orphan's
`TestEnsureClones.test_env_copy_is_copy_if_absent_and_fails_loud` asserted exactly that fail-loud behavior
(`assertRegex(... 'no stack-demo/platform/\.env and no stack-dev/platform/\.env')` + `exit 1`).

**D4 makes phase (b) non-fatal** (copy-if-present, skip-if-absent — a box with no `stack-dev` but with
`.agentspace/secrets` is fully supported; M30 provisions the real `.env`). So the ported test MUST assert the
NEW behavior: the `.env` cp is still guarded copy-if-absent (never clobber), but a missing `stack-dev/.env` is a
**non-fatal skip with a log note**, not an `exit 1`. The `exit 1` fail-loud guards that REMAIN in the script
(phase (a): the broken-partial-clone case + the platform clone-failure case) keep `TestEnsureClones`'s
"lost its fail-loud exits" assertion valid — the script still has `exit 1`s, just not for the absent-`.env` case.
Recorded so the divergence from the orphan's test is intentional, not a port miss.

## M26-D6-impl — the live-recomputed GUIDE count is 41 (computed at build)
`TestGuideDocTruth.test_advertised_test_count_matches_collection` sums `test_*` methods across a FIXED class list.
On current `main` that formula = **28** (GUIDE advertises 28). M26's deltas to that formula:
- `TestShellcheck` 2 → 3 (+`test_ensure_clones_is_shellcheck_clean`)
- `TestRenameDrift` 3 → 5 (replace `test_stack_dev_is_the_preferred_default` with
  `test_dev_stack_still_prefers_stack_dev` + `test_demo_scripts_resolve_stack_demo_as_the_build_source`; add
  `test_ensure_clones_reads_stack_dev_only_for_secrets`)
- `TestEnsureClones` (NEW) +6 · `TestSelfContainedSource` (NEW) +4 — both ADDED to the formula's class list
→ 28 + 1 + 2 + 6 + 4 = **41**. The orphan's stale literal was 40 (computed against a pre-v1.6/v1.7 ancestor).
GUIDE advertises 41; the GuideDocTruth test recomputes dynamically (trust the formula). Re-verify at the
test-section commit (the actual collected sum is authoritative if any count shifts during the port).

## M26-D2-impl — `DEMO` alias vs `DEMO_WS` (the two-subsystem reconciliation)
Current `main` (M30) introduced `DEMO_WS="$REPO_ROOT/stack-demo"` as the **provision-target** var (referenced at
`up-injected.sh:215/228/230`). The orphan used `DEMO` as the **build-source** var. D2 keeps BOTH: `DEMO_WS` stays
for M30's provision references (zero churn there), and `DEMO="$DEMO_WS"` is added as a one-line alias for the
build-SOURCE reads. The orphan's ported tests (`TestSelfContainedSource`) assert the literal `$DEMO` token
(`src="$DEMO/$svc"`, `local ctx="$DEMO/next-web-app"`), so the build-source reads use `$DEMO` and neither
subsystem's references churn.
