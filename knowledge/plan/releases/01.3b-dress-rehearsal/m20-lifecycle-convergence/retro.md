# M20 — Retro

_Lifecycle convergence — demo-up auto set-dress + cold-start capture. The 5th + **FINAL** milestone of v1.3b "dress rehearsal", closed 2026-06-09._

## Summary

M20 converged the **lifecycle**: `/demo-up` now auto-set-dresses by default — a cache-first snapshot **replay**
(real catalog + content) → a `small-200` light seed — at the bring-up tail, exactly as `/dev-up` has since M13.
The defining choice was to **reuse the same `dev-setdress.sh` engine via `--stack-type demo`** rather than fork a
demo copy: ONE engine, two lifecycles, so every read- and write-side safety guarantee carries over by
construction (the demo chain can't drift from the dev one). The net-new `corpus/ops/snapshot-cold-start.md`
runbook closes ISSUE-13 — the fresh-box capture workflow + the spike result that the wired `postgres` MCP is NOT
a capture source (it returns JSON rows, not COPY-format bytes), resolved **document-only** (#M20-D4). The close
was a clean shape — 5 findings (1 docs DOC-1 + 4 decision-triage; 0 scope/code/adversarial-new/tests). With this
close, all 5 v1.3b milestones (M16→M20) are done; the dev↔demo convergence is complete.

## Incidents This Cycle

- **None at close.** No P2 flakes (flake gate 5/5 deterministic on both touched suites), no regressions (all 4
  Go modules unchanged + green, the M15 `safety.md` drift guards re-ran GREEN after the §2.7 edit). The harden
  pass (1 deepening + 1 confirmation scan) fixed **0** production bugs — the build's logic held under every probe
  (capture-never, offset-DSN, n=0-across-types, atomicity-floor, firewall-abort). The close found 1 docs
  consistency gap (DOC-1) and otherwise blended decisions — no behavioral defect anywhere.

## What Went Well

- **Reuse-not-reinvent paid off.** Making the proven M13 engine `--stack-type`-aware (a prefix + a default-preset
  parameterization) was a small, surgical diff that delivered the entire convergence — and meant the prod-safety
  story needed no re-litigation: §2.7 documents that the guarantees carry over *by construction*.
- **The prod-safety boundary was pinned, not asserted.** Capture-never-on-a-bring-up is enforced on BOTH the
  happy path and the cache-miss-degraded branch, mutation-verified (a capture fallback in the degraded branch
  fails the test); the offset-DSN isolation is behaviourally pinned (a base-`5432` regression fails while the
  static fence still passes). The behavioural tests catch what string-match fences can't.
- **The spike was decided on evidence, not vibes.** ISSUE-13's MCP-adapter was settled by reading
  `Capturer.CopyPublic` (raw COPY bytes vs the MCP's JSON rows) — document-only, the M9b-D9 reasoning, zero
  capability gain. A resolved spike that ships the sanctioned path in full, not a deferral.

## What Didn't

- **The CLAUDE.md skill-table row lagged the SKILL.md.** DOC-1: the milestone updated the `/demo-up` SKILL
  description, the demo README, and the convergence narrative, but the root CLAUDE.md skill-table one-liner still
  read "Clerkenstein-wired, offset ports" with no mention of auto-set-dress — caught at close-review. A reminder
  that the *index* row needs the same edit as the doc it points to. Fixed Fate-1.

## Carried Forward

- **Nothing from M20.** All 4 In-scope items landed Fate-1; 0 routed/dropped/escape-hatch-deferred.
- **DEF-M10-01** (S3 media blob bytes + cloud `SnapshotStore` → v1.4, signed) — inherited from v1.2, **untouched
  across all of v1.3b** (a file-level scan of the M16→M20 extensions history found zero `SnapshotStore`/`store.go`/
  S3/blob touches), all 4 aging triggers negative. The release-wide M16→M20 deferral sweep (this close) confirmed
  it stays the one clean signed v1.4 carrier; `/developer-kit:close-release` will carry it forward.

## Metrics Delta

(Source: this milestone's `metrics.json`.)

- **Go test funcs:** **736** (unchanged — M20 touched no Go). All 4 modules build + pass `-count=1`; M15 drift guards GREEN.
- **Python collected:** 338 → **360** (+22 — the set-dress convergence suites: `DevSetdress` stack-type/atomicity + the `TestSetdressChainContract` static fences + behavioural tests).
- **Flake:** **0** (5/5 sequential on dev-stack 50 + demo-stack 84).
- **Lint:** shellcheck CLEAN (2 scripts) · py_compile CLEAN.
- **Extensions:** tag `dress-rehearsal-m20` @ `51a07cb` (reconciled from the build HEAD `e4d2f9b`); `stack-demo` re-consumed.
- **Close findings:** 5 (1 docs + 4 decision-triage; 0 scope/code/adversarial-new/tests). Deferral re-audit GREEN (release-wide sweep: 0 new/repeat/chronic/aged-out).
