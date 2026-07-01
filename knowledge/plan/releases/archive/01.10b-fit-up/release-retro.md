# Release Retro: v1.10b "fit-up"

**Shipped:** 2026-07-01 · **Tag:** `v1.10.1` (rext code-of-record @ `66a021e`) · **Verdict:** GREEN — cold-rebuild-accepted 6/6 + academy F6.
**Milestones (7):** M47 re-sync & recapture · M48 corpus re-ground · M49 bring-up hardening · M50 content & seeding fill · M51 AI-readiness showcase org · M52 single auditable seed+gen manifest · M53 cold-rebuild acceptance.
**Shape:** `M47 → { M48 ∥ M49 } → M50 → M51 → M52 → M53`. M47/M48/M49/M52/M53 `section`; M50/M51 `iterative` (closed-on-gate).
**Nature:** an interposed field-hardening BACKFILL (the v1.3b "dress rehearsal" lineage) — **tooling (rext) + docs (rosetta) ONLY, ZERO platform-repo edits**. NO new third-party deps (`ai v1.40.1` from v1.10 M45 carries forward unchanged).

Consolidates the 7 milestone retros (`m{47..53}-*/retro.md`). This is the durable release-level record; the blow-by-blow is in each milestone's retro + `roadmap.md` §§ M47–M53.

---

## 1. Incidents across the release (P0/P1/P2, and process)

**No P0. No shipped-to-user regression.** The load-bearing incidents:

### P1 — the AI-readiness read-path perf saga (M51 iters 06→09)
The single largest cost of the release. The M42 manager coverage gate coupled M51's *seeding* deliverable to a **platform** AI-readiness read-path wall. Three successive read-fast strategies were each **falsified by a cheap probe before any full sweep or shipped edit**:
1. active-cycle signals-true — never completes (per-skill translation N+1);
2. closed-cycle frozen snapshots — the default FE GET omits `?cycle=`, so the frozen branch is never selected;
3. deep-link the frozen branch — the frozen branch is *itself* org-scale-slow.

Root cause (iter-08): **"frozen SCORES ≠ frozen RESPONSE"** — `buildResponseFromSnapshots` re-joins CURRENT members via an unbounded whole-org `loadMembers`. The three-tik no-progress run correctly fired TOK-02; the user authored the pivot. iter-09 landed the `app-aireadiness-snapshot-loadmembers` demo-patch — a **pure, data-identical** bound of that hydration to the ~199 snapshot users (frozen `?cycle=` GET: 180s timeout → 19ms), and the gate fell 5→0. Zero platform-repo edit (a demo-only in-inject-loop source swap, the `app-targetrole-authz-skip` precedent). The **prod finding** (the platform's frozen-read still hydrates the whole org) is **M314b**, documented not fixed — a prod-team follow-up, out of tooling scope.

### P1 — AB4: an M51-owned cross-milestone gate regression caught from cold (M53)
On the first from-cold both-vantage sweep, the **M50 canonical manager gate** (`dan-manager` @ Cervato) was RED — `/enterprise/workforce/ai-readiness` rendered HTTP 200, 0 ejects, but showed "No AI readiness data yet." Root cause: M51 iter-05 added the ai-readiness page to `MANAGER_MANIFEST.seedPaths` **unconditionally**, but the 199 frozen snapshots seed only for Northwind (the showcase org); M51's gate ran ONLY Northwind and never re-ran the M50 Cervato sweep. Fix (user-approved gate exception, rext `117fe41`, +3 unit tests): org-condition the manager manifest — `manifestFor(vantage, expectedOrg)` returns the showcase manifest only for Northwind, else a new `MANAGER_MANIFEST_BASE` that omits the seedPath. **This is exactly the late cross-milestone regression M53 exists to catch** — the fix-on-live serialization never re-ran M50's gate after M51's manifest change; the from-cold both-vantage assertion is the first joint re-measure.

### Security — inherited HIGH CVE cleared at release-close (Phase 0/7)
`CVE-2026-39821 / GO-2026-5026` (HIGH, called): `golang.org/x/net@v0.53.0` `idna.ToASCII` hostname-validation bypass, path exists only in `stack-seeding` (`reloadStackSentinel → idna.ToASCII`), indirect via the `ai v1.40.1` tree. **Inherited** — in the tree since v1.10, CVE disclosed *after* the v1.10 close; NOT a v1.10b regression (v1.10b added 0 deps). Surfaced by the Phase-0 re-scan on 2026-07-01; cleared Phase-7 with `x/net v0.53.0→v0.55.0` (+ `x/text v0.36.0→v0.37.0`) in stack-seeding. `govulncheck` re-confirmed at Phase-8: **"No vulnerabilities found."**

### P2 (caught in review/harden, none shipped broken)
- **M48 — plausible-but-wrong seed contract.** The first `ai-readiness.md` draft offered a "snapshot-direct" strategy that a Phase-2c code read of `computeOrgBreakdowns` disproved (the active-cycle dashboard recomputes from signals; `ai_readiness_live_snapshots` is a Talk-to-Data cache). Rewritten as cycle-state-dependent before merge — had it shipped, M51 could have built against a wrong contract.
- **M49 — `rext_tag.sh` CRLF leak.** The awk reader captured a trailing `\r` from a CRLF-edited `.agentspace/rext.tag`, which would fail `git checkout <ref>` with a baffling "pathspec did not match." Fixed with a portable `gsub(/\r/,"",$1)`; negative-control regression test.
- **M50 — the run-1 gate passed BLIND.** iter-05 reached `(0,0)` both vantages but the manager manifest never ASSERTED languages/certs/member-fields — a green gate co-existing with two genuinely empty surfaces (languages 0 rows DB-wide, certs 9/340). iter-06 filled them AND strengthened the manifest. (The seed of the AB4 lesson: a green `(0,0)` is only as honest as the manifest asserts.)
- **M50/M51 — misreads corrected by protocol discipline.** M50 iter-02 misread the manager sweep as frontier-capped (iter-03 re-survey found it EXHAUSTS; the tooling-iter was cancelled). M51 iters 06→08 each falsified a plausible strategy cheaply.
- **M51 — pre-checked-but-absent fix-queue (resume-caught).** A close agent that died mid-Phase-7 had marked C4/T1/T2 `[x]` before authoring them; the resume verified ground truth and authored them for real. Lesson: check-off follows the landed artifact, never intent.
- **M47 — accidental background-task kill (P3, recovered).** A ~1.4 GB recapture was UI-killed mid-stream; the atomic `.tmp`→rename write left the prior cache intact and a clean re-launch completed it.

### PROCESS — sub-agent sweep-execution fragility (the orchestration-layer hazard)
Surfaced repeatedly, most acutely in M51 (and shaped the whole release's execution discipline). Not a corpus-doc subject — a durable **process** lesson. The hazards:
- **background-yield** — a long sweep spawned to run "in the background" that the driving turn then blocks on or loses track of;
- **the 13-min sweep vs the 10-min Bash cap** — the Playwright coverage sweep exceeds a single foreground Bash timeout, tempting a background spawn;
- **the concurrency fork** — a second agent touching the single shared demo / consumption-clone mid-run (M51 iter-04's dirty `stack-demo/rosetta-extensions` from a prior-concurrency partial M50 application, unblockable only via forbidden ops; the orchestrator reset it out-of-band).

**Mitigation that worked, keep it:** commit-first (land + tag before any long sweep, so a lost sweep never loses work) + bounded sweeps (per-vantage, frontier-exhausted, not open-ended) + foreground / no-double-spawn discipline (one agent owns the single demo machine at a time). The single-demo serialization (below) is the structural driver; disciplined commit-and-bound is the mitigation.

---

## 2. Cross-milestone patterns

- **The acceptance gate catching a real regression = the model working.** M53's whole reason to exist is "surface a regression the warm fix-on-live chain hid" — and it surfaced exactly one (AB4), narrow and test-covered. A from-cold, both-vantage, all-3-org acceptance is a real gate, not a rubber stamp. The org-conditional manifest split is the durable fix: showcase-only surfaces are now gated on the org, so a base-org gate can't be broken by a showcase-org prime.
- **The single-demo serialization shaped the entire release.** One live demo-1 was the shared substrate; M47→M52 fixed-on-live against it, then M53 destroyed + cold-rebuilt it as the single acceptance proof. This is why fix-on-live never re-ran M50's gate after M51's manifest change (AB4), and why the cold-rebuild is load-bearing rather than ceremonial. It also drove the sweep-fragility discipline above (a single machine = strict foreground serialization).
- **A green coverage gate is only as honest as its manifest ASSERTS.** The recurring structural risk of presence-only coverage: M50's run-1 blind-pass, then AB4. Cross-check the gate's assertions against the milestone's INTENT, and re-run EVERY gate a shared manifest can affect — not just the one being tuned.
- **Cheap-probe-first, then commit.** The highest-leverage protocol move of the release: a direct authed probe on the exact fast endpoint falsified the deep-link premise in 40ms+180s-timeout BEFORE any cockpit/manifest edit or a 13-min sweep. Measure the FAST branch end-to-end before committing to it.
- **A close-discovered staleness claim should be re-measured before it drives a release's shape.** M47's whole heavy framing (and part of v1.10b's "re-ground" thesis) rested on "clones 5 weeks behind" — which a `git fetch` + behind-count disproved instantly. The genuinely-stale surface was the corpus (M48), not the clones.
- **Verify-against-code, not table-names / prose totals.** M48 traced the read path in code (not inferred a write contract from table names); M49 pinned counts from the runner (not prose running-totals). Doc milestones that feed a downstream seeder must trace the read path.
- **The two-repo split held cleanly all release.** rext code froze per-milestone at `fit-up-m47..m52` and rolled to `v1.10.1`; rosetta carried only docs + records; close reviewed BOTH repos and merged only the rosetta doc half — exactly the tooling-policy shape.

---

## 3. Carry-forward items + destinations

| Item | Class | Destination |
|---|---|---|
| **Origin pushes** — `main` + the `v1.10` tag + v1.10 ext tags + `fit-up-m47..m52` + `v1.10.1` | push-gated KEEP (administrative, not a deferral) | **The user's manual step** (local closes deliberately do not push). The box-level re-pin (consumption clone + `.agentspace/rext.tag` → `v1.10.1`) is DONE. |
| **M314b** — prod frozen-read still hydrates the whole org (a `frozen_tags` column / bounded re-join) | prod finding, not tooling | **A prod-team follow-up** — documented in `coverage-protocol.md` + `services/ai-readiness.md`; the demo works around it, the platform still carries it. Out of tooling scope. |
| `member_languages` surface token underscore-vs-hyphen | nice-to-have, cosmetic (not in any DependsOn) | Open in `release-review.md` (Phase-2); consciously not landed — the one un-ticked review box, GREEN verdict stands. |
| Standing cross-release backlog: DEF-M10-01 (cloud SnapshotStore/S3), DEF-M21-01 (`replayCmd` hermetic test), M25-D9 (dev taxonomy `rc=4`); future v2 M205–M207 | unscheduled | `roadmap-vision.md` — none scheduled. |
| v2.0 "opening night" (Playthroughs) | PAUSED (M201 closed, M202∥M203∥M204 not started) | Resumes next, after this backfill ships. |

Every carry that pointed at a *milestone* landed in that milestone — the deferral audit closed **GREEN**, with the one historical repeat (academy F6, M50→M51→M53) **resolved by execution** at M53's cold rebuild.

---

## 4. Metrics delta + stats reference

From the aggregated `metrics.json` (baseline = v1.10 close):

| Metric | v1.10 close | v1.10b close | Delta |
|---|---|---|---|
| rext Go test funcs (5 modules) | 1551 | **1640** | **+89** (increase, no decrease) |
| — by module | seeding 706 · snapshot 363 · secrets 160 · align 52 · clerk 270 | seeding **791** · snapshot **364** · secrets **163** · align 52 · clerk 270 | seeding +85, snapshot +1, secrets +3; align/clerk untouched |
| demo-stack Python (whole-dir) | (touched-suite only) | **326** | every touched suite grew, none decreased |
| rext e2e TS unit (2 specs) | 17 | **42** | +25 (coverage-manifest 17→29 + NEW section-assert 13) |
| Flake | 0 | **0** | no regression (per-milestone 5/5; release-close 3×3 clean) |
| New third-party deps | — | **0** | `ai v1.40.1` carried forward unchanged; `go:embed` is stdlib |
| Coverage | — | no >2pp drop on any measured surface | seeders ~97.5→97.6%, NEW manifest pkg 100% stmt |
| Supply-chain | (clean at v1.10) | node audit 0 vulns; Go CVE-2026-39821 **cleared** (x/net→v0.55.0) | 1 inherited HIGH found + fixed at close |
| Alignment gates | 100%/100% (5 surfaces) | N/A-change (clerkenstein untouched) | carried forward |

**Apples-to-oranges caveat:** v1.10b is a docs+tooling backfill with ZERO platform edits; a strict like-for-like with v1.10 (a feature release) is partly apples-to-oranges (esp. the Python metric shape — v1.10 recorded touched-suite counts, v1.10b the whole-dir total). The load-bearing close-release check — **NO test regression + flake 0** — holds unambiguously: Go +89, Python/TS grew, flake 0, coverage stable, 0 new deps.

**Release-close verification (Phase 8/8b):** full suites re-run + a **3× shuffled flake gate** all clean — Go 1640 (5 modules, `go vet` clean), Python 326, TS unit 42; **3/3 rounds green per stack**; `govulncheck` on stack-seeding "No vulnerabilities found."

**Stats-delta reference:** the release-close project-stats snapshot is `knowledge/journal/stats/2026-07-01.json` (git: 635 commits / 77 merges / 551 co-authored / 193 in the last 7d / 160-day project age / 15 decisions). NOTE: the stats.sh source/docs auto-detection reports 0 code/doc lines for this repo — expected, and worth flagging: the rosetta corpus lives under `corpus/`+`knowledge/` and the **code-of-record lives in the gitignored `.agentspace/rosetta-extensions`** repo, neither of which the generic source-layout detector maps onto. The git/structure/decisions figures are the meaningful ones for a docs+tooling repo.

---

## 5. Release verdict

**GREEN from cold.** demo-1 destroyed + cold-rebuilt from `v1.10.1` by a single `/demo-up` on a `stack-demo`-only box; **6/6 acceptance criteria + academy F6 GREEN**. Tooling + docs only, zero platform-repo edits. The backfill did its job: the (M47-confirmed-current) clones re-synced + snapshot recaptured, the corpus re-grounded to current prod (the new `ai-readiness.md`), the from-cold `/demo-up` hardened across the known issues, the content-seeding + AI-readiness-org gaps filled to a `(0,0)` semantic coverage gate at both vantages, the seed intent consolidated into one auditable manifest, and cold-rebuild acceptance proven. Next: `/developer-kit:close-release` merge + tag (origin push is the user's step).
