# close-release session log — v2.5 "the playbill"

**Paused:** 2026-07-20 (user paused the machine mid-Phase-8)
**Resume with:** `/developer-kit:close-release` — it reads this file at Before-You-Start step 2.
**Branch:** `release/02.50-the-playbill` @ `ddcfc17` · **rosetta-extensions:** `main` @ `29dec48` (pushed)
**Both trees clean. No stray git worktrees. Nothing in flight.**

---

## DO NOT RE-RUN — completed phases

| Phase | Status | Artifact |
|---|---|---|
| 0 supply chain | **GREEN** | `dependencies.lock` (`b855987`) — 0 CVEs, 0 GPL/AGPL |
| 1 release scope | **GREEN** | 11 of 12 in-release Fate-3 routings verified landed |
| 1b deferral audit | **was RED → discharged** | `audit-deferrals/deferral-audit-2026-07-20-release-close.md` + `release-deferrals.md` |
| 2 code quality | done | 4 must-fix + 15 should-fix + 10 nice-to-have — all landed |
| 3 documentation | done | 11 must-fix + 6 should-fix — all landed |
| 3b KB consolidation | **3 blockers landed** | clone-freshness anchor, R1 gap disclosure, `context.md` |
| 4 tests | done | counts below |
| 4b metrics regression | **was RED (`flake_count` 2)** | both flakes FIXED test-side in Phase 7; **metrics.json not yet updated — see TODO 3** |
| 5 decisions | done | 9 unblended + 7 conflicting — all landed |
| 6 review | done | `release-review.md` (`ddcfc17`) with per-item disposition + the 6 findings the review got WRONG |
| 7 fix everything | **done** | 6 parallel agents, ~32 commits (25 rosetta + 7 extensions) |

## Phase 8 — PARTIALLY DONE (resume here)

**Completed and trustworthy** (observed before the pause):

- **Full suites, both repos: 4152 tests · 8 failures · 0 errors.**
- The 8 failures are **exactly** the documented standing `demo-stack` macOS set (4 M53 academy-link,
  2 M218 overlay, 1 M215 browser-port list, 1 macOS-conditional purge). **8 = the clean-clone number, so
  the `stack-demo` clone set is still pristine.** (14 would have meant it went dirty — it did not.)
- One `gofmt` violation, **verified pre-existing from M225**, not introduced by Phase 7.

**NOT done — pick up here:**

1. **Cross-reference link check.** Six agents rewrote docs concurrently and `roadmap.md` was split
   2079 → 491 lines (`roadmap-archive-v2.0-v2.4.md` created), plus new `corpus/services/README.md` and
   `corpus/tools/README.md`. Verify every relative link and intra-doc `#anchor` resolves.
   **Known pre-existing, NOT regressions:** `roadmap-vision.md`'s two links into gitignored `.agentspace/`
   scratch; `coverage-protocol.md:333`'s literal `](#…)` placeholder from M46.
2. **Phase 8b triple-clean gate (BLOCKING).** No CI wiring — use the documented fallback: 3 independent
   full-suite runs in random order, in separate worktrees under the scratchpad, launched as parallel Bash
   calls. Green = the expected 8 standing failures and nothing else. Any failure restarts the counter.
   CI-wiring remains a carry-forward item.
3. **Update the flake count.** `releases/02.50-the-playbill/metrics.json` still records `flake_count: 2`
   from Phase 4b. Phase 7 A2 fixed **both** flakes test-side (neither assertion weakened) and measured
   5× clean; the stopped verification run saw 0 errors. **Confirm via 8b, then set it to 0** in
   `metrics.json` and the `metrics-history.md` row, with a note that they were fixed, not waived.
   **If 8b does not confirm 0, change nothing and report RED.**
4. **Cut the corrective tag** `playbill-v25-close` on rosetta-extensions `main` @ the verified state,
   annotated as the authoritative v2.5-close tooling state. **Why:** two Phase 7 agents committed
   concurrently and one used a broad `git add`, so `playbill-v25-close-go` and `playbill-v25-close-harness`
   have mixed attribution and do not describe their own contents. Documented, **not to be rewritten** —
   history is pushed.

## Then: Phases 8c → 12

- **8c** `/developer-kit:project-stats` snapshot (before merge, after CI cleans)
- **8d** `release-retro.md` — consolidate the 8 milestone retros
- **9** release completeness ledger + **explicit user sign-off** (see the two open decisions below)
- **10** roadmap rotation (`## Active` → `## Done`), archive `releases/02.50-the-playbill/` →
  `releases/archive/`, CHANGELOG entry, semver check
- **11** merge → `main`, tag `v2.5`, delete the release branch, delete this file, sweep
  `.agentspace/scratch/work-m*/journal.md`

---

## OPEN — needs the user, do not decide autonomously

### 1. The 13 exhibits cannot be proven org-clean offline (CQ-3 residual)

The org-scrub arm never fired across all 13 copied production sessions — **0 `<<ORG>>` placeholders against
840 `<<ACTOR_i>>`**. Phase 7 fixed the *mechanism* structurally: it now fails closed, an empty org is itself
an error, and "registered for scrubbing but not for leak-checking" was made **unexpressible** in the code.
**But the existing fixtures' status is unknown** — the source org name is never persisted by design, so
there is no in-repo artifact to check against. The asymmetry is equally consistent with "the company was
never mentioned" and "it was copied through verbatim". This is real customer session content.

**Settling it:** one read-only prod query resolving the 13 source org names, then `scrub.OrgTokens` +
`scrub.SurvivingToken` run offline over the already-committed fixtures. **No re-capture needed, and no
customer names need enter an agent transcript.** Clean → proven. Dirty → re-capture required (which will
now hard-fail at the leak gate rather than silently emit).

**User decision: tag v2.5 before or after this check.**

### 2. Already decided (2026-07-20) — carry forward, do not re-ask

- **Fix everything** (not a triaged subset) — done in Phase 7.
- **`29/29` tags as unit-proven, NOT live-re-proven**; the live re-prove is **v2.6's opening work**.
  Recorded as an escape hatch with a per-item why-Fate-1/2/3-failed rationale in `release-deferrals.md` §A,
  with an acceptance condition: the re-prove only counts if green *at the tag that actually shipped*, with
  the CQ-1 grader fix, CQ-2 runner wiring, and an externally-sourced `EXPECTED_PAIRS`. Anything less
  re-issues the same false green.
- Earlier M236 gate decisions (reach clause dropped; p95 hero-only) — landed in the corpus in Phase 7.

### 3. Phase 9 will surface these for sign-off

Escape-hatch items routed to **M237** (reserved): `CLOSE-D3` (29/29 not live-re-proven) · the 39
unexecuted live-browser specs · the anonymous academy catalog twin · `DEF-M226-01` (**with teeth** — M237
must *test* the never-tested "self-resolves" claim or DROP) · the three v2.3 `DRIFT_DEFER` carries · the
interview plan-section assertion. To **M238**: `DEF-M235-03` assign-WRITE (~10 routings across 5 releases)
with a drop-expiry.

**A release tagging with >0 escape-hatch items is unusual** — Phase 9 must ask whether v2.5 should tag as
`v2.5` or `v2.5-incomplete`. Given the headline metric ships unverified-live, that question is real.

---

## Context a fresh session will not otherwise have

- **The release's thesis:** *a check can report success while proving nothing* — 9 instances in M235–M236,
  and the class was found **alive inside the close that named it**. Treat any green with suspicion; ask
  what it prints when nothing happened.
- **Six findings this review got WRONG** are recorded at the end of `release-review.md`. Read that section
  before acting on any finding text — several "defects" were correct lines.
- **`/demo-up` rebuilds images from clones it never updates** (`app` was 249 behind `main`, `next-web-app`
  202). Anything verified against local `stack-dev` clones may be wrong; verify against platform
  `origin/main`. This was the user-reported stale left menu and is anchored in `rosetta_demo.md`
  § Clone freshness.
- **`state.md` is 14,043 / 15,360 bytes (91.4%).** The Phase 8 rewrite must not grow it past the cap.
