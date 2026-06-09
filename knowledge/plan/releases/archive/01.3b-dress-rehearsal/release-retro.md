# Release Retro: v1.3b "dress rehearsal"

**Shipped:** 2026-06-09 · **Milestones:** M16 → M17 → M18 → M19 → M20 (all section, strictly sequential) · **Branch:** `release/01.3b-dress-rehearsal`

## What this release was
A field-hardening release between shipped v1.3 and future v1.4. v1.3 converged the dev/demo *model*; the first real `/demo-up` run surfaced 14 issues showing it hadn't converged the *experience* (backend-only, unseeded, unverified, "UP" announced over a silently-failed authz). v1.3b makes `/demo-up` produce a **full, populated, verified, demoable** stack — all **tooling + docs, zero platform-repo edits**. All 14 issues delivered Fate-1.

## Incidents this cycle
No P0/P1, no production incidents (this is dev/demo tooling, not the platform). The notable in-cycle catches — each found AND fixed before its milestone closed:
- **M18 harden** found a *real bug build missed*: the readiness phase of `stack-verify` ignored `STACK_SERVICES`, so a reduced-profile bring-up still ran all 6 deep probes and false-`down`ed — the exact ISSUE-12b wall-of-false-downs the milestone exists to kill. Fixed inline + doc corrected.
- **M18 close** (adversarial Phase 2c) found A1: a non-numeric `--offset` crashed the verify under `set -u` and *silently skipped* the cheap-win asserts. Sanitized at the resolution boundary; the non-fatal invariant held throughout.
- **M16 close** found + fixed the *last* stale workspace name in the whole extensions repo (clerkenstein glossary `anthropos-demo`), reversing a build-time boundary defer (M16-D5 → Fate-1).
- **M19 build PR-review** caught a broken offset cross-check (`port//10000` decade lane) that would false-warn on roadrunner's high base port — fixed with 4 regression tests before close.

## Cross-milestone patterns
- **Recurring: "new deliverable shipped without its index row."** v1.3 flagged this in M14 + M15 retros; it **recurred** here — the 4 net-new corpus docs weren't all added to the *parent* `corpus/ops/README.md` index (M19 frontend-tier + M20 cold-start were missing). Caught at release-close (Phase 3) and fixed (2 rows added; the demo README also gained the missing `verification.md` entry). **Carry → process:** the per-milestone close index-row check must cover **corpus *directory* READMEs**, not just per-unit handbook READMEs. This is the third release it has appeared in — worth a lint/guard in a future release.
- **Harden earns its keep on code milestones.** Every code milestone's mandatory harden pass produced load-bearing value: M17 pinned the per-stack-only TRUNCATE target-class + injection sweep; M18 found the readiness-scope bug; M19 locked the zero-platform-edit invariant as a real-git-repo test (`TestZeroPlatformRepoEdit`); M20 extended capture-never to the cache-miss degraded path. The docs milestone (M16) used harden for drift/contract guards instead of coverage. Confirms the v1.3 lesson: rename/docs milestones reward *contract guards over raw test counts*.
- **Prod-safety held under field-hardening.** The release added destructive-ish primitives (replay TRUNCATE-reload, seed --reset, auto-set-dress chaining) yet the M15 safety contract held byte-for-byte — every milestone re-ran the `safety.md`↔code drift guards GREEN, and capture is test-pinned to *never* run on a bring-up (happy + degraded paths). One engine (dev-setdress.sh), two lifecycles, one safety boundary.

## Carry-forward
- **DEF-M10-01** — S3 media blob bytes + cloud `SnapshotStore` backend → **v1.4** (signed escape-hatch, recorded in `roadmap-vision.md`). Untouched across all of M16–M20; not aged. The clean signed carrier.
- **Go toolchain go1.25.11+** — accepted environment item (12 *stdlib* advisories @ go1.25.3; 0 called third-party CVEs). Identical disposition to v1.3. Clear by bumping the installed toolchain; no code change. Track into whatever release next touches the build environment.
- **True zero-rebuild frontend** (M19, OUT) — an optional **user-owned upstream platform PR** (runtime URL/pk rewrites + `__env.js` + `output:standalone`). Editing platform repos is forbidden to the tooling; documented in `frontend-tier.md §4`, not a deferral.
- **MCP-as-capture-source** (ISSUE-13, M20) — RESOLVED document-only (the wired `postgres` MCP returns JSON rows, not the COPY bytes the manifest checksum needs); not carried.

## Metrics delta (vs v1.3)
- **Go test funcs:** 713 → **736** (+23, all M17; stack-seeding 236→252, stack-snapshot 224→231; M18/M19/M20 touched no Go). All 4 modules `-race -count=1` clean.
- **Python:** 174 (v1.3 test-funcs) → **360 collected** — net-new `stack-verify` suite 0→87 (M18), demo-stack 13→84, dev-stack 38→50, stack-injection 69→85. (Matcher shifted funcs→collected mid-release; growth real either way, no section regressed.)
- **Tests executed at close:** 1,096; **flake 0**; **triple-clean 3/3** (shuffled, all 9 suites). Supply-chain **GREEN** (0 called third-party CVEs; all-permissive; Python stdlib-only). 4 net-new corpus docs. Data-DNA 100% held; Clerkenstein gates 100/100 (untouched).

## Stats delta reference
Snapshot `knowledge/journal/stats/2026-06-09.json` vs `2026-06-07.json` (v1.3 close): commits +39, source +64.5k lines, total +86k lines over the window. **Caveat:** the stats script scans the gitignored `stack-*/` platform clones (not just rosetta + extensions), so absolute file/line counts are inflated and the test-runner auto-detect reads "none" — treat the absolutes as noise; the commit/velocity deltas roughly track the corpus + extensions work. (Known limitation; a stats-scope fix would be a tooling nicety, not tracked.)

## Process notes for next time
- The work-milestone → build/harden/close pipeline ran 5 milestones cleanly with per-milestone extensions push + tag-reconcile-at-close; the two-repo (rosetta docs ∥ rosetta-extensions code) split held without confusion. The one orchestration glitch was a release-close triple-clean **shell-quoting bug** (zsh non-word-splitting an unquoted var) — caught immediately (it reported FAIL with a clean tree + a `cd: no such file` error, obviously not a test failure), re-run with literal loop lists, PASS. No impact.
