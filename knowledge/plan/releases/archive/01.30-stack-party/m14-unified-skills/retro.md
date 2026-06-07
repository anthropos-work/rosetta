# M14 — Retro

**Milestone:** Unified `stack-*` skills + `dev-up`/`dev-down` (hard-renamed) · **Shape:** section · **Closed:** 2026-06-07
**Branch:** `m14/unified-skills` → merged `--no-ff` into `release/01.30-stack-party` · **Tag:** `stack-party-m14` @ `33fc525` (extensions)

## Summary

The **tooling-convergence milestone of v1.3**: M12 unified dev+demo for N-allocation, M13 for DATA — M14 unifies
them for the **operator interface**. One coherent stack-operations skill set now works on any stack: `dev-up`
(consolidating the former `setup-platform` + `start-platform` into one dev bring-up that drives the M13 set-dress
flow) + net-new `dev-down`, and four hard-renamed ops skills — `stack-list`/`stack-seed`/`stack-snapshot`/
`stack-update` (←`demo-status`/`demo-seed`/`demo-snapshot`/`update-platform`) — each taking a `dev-N|demo-N`
target. The rename is a **clean break, no aliases** (user 2026-06-07): the 6 old skill dirs were removed and
**every** in-repo reference swept (the CLAUDE.md 14-skill table rewritten, root+corpus READMEs, all `corpus/ops/`
guides + the `demo/` recipe family, CHANGELOG; `demo-up`/`demo-down` retained + aligned with the dev lifecycle).
This is a docs/reference-integrity milestone — the deliverable IS the converged skill set + the corpus, so the
quality bar is the **contract** (rename invariant · CLAUDE.md⇄dirs bijection · skill→CLI resolution), not new
test counts. Built in 5 sections, hardened in 1 pass (a reference-integrity guard), closed clean with 0 findings.

## Incidents This Cycle

**No defects shipped. No regressions. 0 P2 flakes.**

- **1 PR-review finding (build, fixed inline, not a defect):** the `stack-seed --preset NAME` UX (inherited from
  the old `/demo-seed`) is a **skill-level shorthand** — the `stackseed` binary only knows `--seed <path>`. The
  SKILL.md now explicitly maps `--preset NAME` → `--seed presets/NAME.seed.yaml` so an automated invocation never
  passes a bogus flag (M14-D5). Caught before merge, no shipped impact.
- **1 stale guard-comment (harden, fixed + guarded):** the M11-established stacksnap docs↔parser flag-drift guard
  named the **retired** `/demo-snapshot` skill (+ its deleted SKILL.md path) as the doc-side source of truth in 6
  comment lines. M14 hard-renamed that skill, so the comments pointed a future maintainer at a path that no longer
  exists. Fixed to the renamed skill + added **`TestDocSourceSkillRename_M14`** (asserts the renamed dir exists +
  the retired one stays gone; negative-tested by re-appearing the retired dir), turning the silent-comment-rot /
  dir-resurrection failure mode into a test failure.

## What Went Well

- **The hard-rename blast radius was fully contained.** The clean-break tradeoff (the user accepted it) was de-
  risked by an exhaustive reference sweep + a project-wide invariant scan: at close, **0 live stragglers** — every
  remaining mention of a retired name is an intentional "formerly /X → /Y" provenance marker.
- **Reference-integrity was made a *test*, not a convention.** The rename invariant now fails CI (in the authoring
  clone) if a retired skill dir reappears or the renamed one vanishes — the durable guard against future rot.
- **`dev-up` consolidation dropped nothing.** An element-by-element diff of the consolidated `dev-up` body against
  the former `setup-platform` + `start-platform` bodies confirmed every step preserved (STEP RUN discipline,
  confirmation policy, ops reports, the 12-container set, error recovery).
- **Clean straight-through close** — 0 findings, deferral re-audit GREEN, no fix queue.

## What Didn't

- **Nothing material.** The one process note: the previous milestone's close sub-agent had erroneously returned
  right after writing the close-journal header without performing the merge (a known harness bug). This close
  resumed from that clean build+harden state and ran to full completion (merge + branch-delete + tag + bookkeeping).

## Carried Forward

- **DEF-M10-01** (S3 media blob bytes + cloud `SnapshotStore` backend) → **v1.4**, inherited from v1.2, signed by
  the user at v1.3 design. M14 (a skills/docs milestone) touched no media or snapshot-store code path; the item is
  not aged out. Re-confirmed in this close's deferral audit (GREEN).
- **The safety & security doc + dual-repo KB refresh → M15** (the LAST v1.3 milestone). This is Fate-2
  confirmed-covered (M15's `In:` list owns `corpus/ops/safety.md`), **not** a deferral.
- **No new deferrals introduced by M14** — every scope item landed Fate 1.

## Metrics Delta

(Source: `metrics.json`.)

- **Go test funcs:** 720 → **721** (+1: `TestDocSourceSkillRename_M14`, the reference-integrity guard).
  alignment 46 · clerkenstein 218 · stack-seeding 233 · stack-snapshot 223→**224**.
- **Python test funcs:** **174** (unchanged — M14 touched no Python surface).
- **Quality gates:** all 4 Go modules `-race -count=1` green; gofmt + `go vet` clean; 3 CLIs shellcheck-clean;
  py_compile clean; **flake 0 (5/5)**.
- **Reference-integrity (the headline):** rename invariant 0 live stragglers · CLAUDE.md ⇄ 14 skill dirs perfect
  bijection · SKILL.md `name:`⇄dir 0 mismatches · skill→CLI contract resolves · all doc links resolve.
- **Findings:** **0** at close (1 build PR-review fix + 1 harden guard-comment fix, both pre-merge).
