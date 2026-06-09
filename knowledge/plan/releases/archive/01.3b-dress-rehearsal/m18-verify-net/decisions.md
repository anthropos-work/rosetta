# M18 — Decisions

_Implementation decisions with rationale. ID scheme: M18-D1, M18-D2, … Open questions from `overview.md` get resolved here during build._

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| M18-D1 | Offset source = the N the bring-up already knows (`--project`/`--offset`), cross-checked (non-fatal) against the registry's RECORDED ports — never a bare formula. | The bring-up allocated N so it knows the offset (zero drift); operator runs resolve from the registry record (M12 records resolved ports). Resolves the overview open-Q "STACK_ROOT-parse vs recorded ports": read the record. | 2026-06-09 |
| M18-D2 | Offset/scope contract = env vars `STACK_PROJECT` / `STACK_OFFSET` / `STACK_SERVICES`, resolved in a new shared `lib/target.sh` sourced by services.sh + readiness.sh. | One resolution helper, applied once centrally; the SERVICES table stays a single base-port source of truth. Empty filter = all-in-profile (back-compat default). | 2026-06-09 |
| M18-D3 | Auto-verify = a thin NON-FATAL wrapper `live/autoverify.sh` (always exits 0), wired default-on at both bring-up tails, opt-out via `DEMO_NO_VERIFY`/`DEV_NO_VERIFY`. | Mirrors dev-setdress's proven default-on + non-fatal pattern. The load-bearing correctness mitigation: a verify/offset bug can never block a genuinely-good stack. Resolves the overview open-Q "how loud": a clear ⚠ block + a "run /test-platform N" hint. | 2026-06-09 |
| M18-D4 | Cheap-win asserts live INSIDE autoverify (before the full probe set), gated by the same scope filter. | One non-fatal warning surface; the /api/health + casbin_rules>0 checks are the seconds-fast, decisive ISSUE-7 catcher; skipping out-of-scope services avoids false warnings on a reduced bring-up. | 2026-06-09 |
| M18-D5 | The cross-check uses a base-port BAND `(port - offset) ∈ [3000, 11000]`, not a `port//10000==n` decade lane. | PR-review A1: roadrunner's base 10400 → 20400 for n=1 sits in the n=2 decade; the lane test would false-warn a correct offset. The band covers all 12 bases (3200 gotenberg .. 10400 roadrunner) with no hardcoded table. | 2026-06-09 |
| M18-D6 | Tests added to a NEW `stack-verify/tests/test_verify.py` (stdlib unittest), not to demo-stack's `test_tooling.py`. | stack-verify had no Python tests before (only Playwright e2e); the wiring fences belong with the code they fence. Keeps the demo-stack GUIDE's pinned `test_tooling.py` count (27) untouched. | 2026-06-09 |
| M18-D7 | A non-numeric `STACK_OFFSET`/`--offset` is sanitized at the single resolution boundary (`target_resolve_offset` validates `^[0-9]+$`, else warns non-fatally + derives from the project's N). | close FINDING-A1: the un-validated value flowed verbatim into three `$(( base + offset ))` sites and, under `set -u`, crashed with "unbound variable", silently skipping the cheap-win asserts (non-fatal held, but confusing). Fixing once at the resolver — not at each arithmetic site — matches the M18-D2 "one resolution helper" design and keeps the SERVICES table the single source of truth. | 2026-06-09 |

## Adversarial review (close Phase 2c)

Scenarios considered against the load-bearing non-fatal invariant — _can any path abort or systematically false-`down` a genuinely good bring-up?_

1. **autoverify with no registry / no docker / no backend, reduced scope** — must exit 0. ✓ (`set -uo pipefail`, every probe wrapped, final `exit 0`; call sites also `|| true`.)
2. **autoverify with an empty `--services ""`** under `set -u` — must exit 0. ✓ (empty filter = all-in-profile; no unbound access.)
3. **No `--project` (degenerate)** — must skip + exit 0 non-fatally. ✓ (explicit early `exit 0` with a warning.)
4. **Non-numeric `--offset` ("abc")** — FOUND FINDING-A1: crashed three arithmetic sites with "unbound variable" under `set -u`, silently skipping the cheap-win asserts. The wrapper still exited 0 (invariant intact) but the asserts vanished with a confusing message. **Fixed** (M18-D7): sanitize at `target_resolve_offset`; warn + derive. Regression pinned at both the unit (`TestTargetHelperBoundaries`) and integration (`TestAutoVerifyEdges`) level.

Verdict: the non-fatal invariant held in every scenario (the bring-up was never abortable); A1 was a robustness/clarity gap on a non-production input path (operator/`/test-platform` typo), fixed inline as Fate-1.
