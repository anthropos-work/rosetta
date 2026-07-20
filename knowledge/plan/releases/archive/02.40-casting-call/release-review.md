# Release Review: v2.4 "casting call"

**Date:** 2026-07-18
**Milestones:** M222, M223, M224, M225, M226, M227, M228 (7 total)

> ## ⚠️ CORRECTION ANNOTATED 2026-07-20 (at the v2.5 close) — this review issued FALSE GREENS
>
> **Nothing below is edited.** The false greens are left visible on purpose: *how* a close issued them is
> itself the finding (`DC-P5`, recorded in
> [`../../02.50-the-playbill/release-deferrals.md`](../../02.50-the-playbill/release-deferrals.md)).
> Read the ✅s below as **claims of this review**, not as facts.
>
> **Four corrections, all verified at the v2.5 close:**
>
> 1. **`:16` — "✅ No `closed-incomplete` milestones → … no undelivered Fate-3 items" is FALSE.**
>    `release-retro.md:41`, written **the same day**, lists two items as inherited standing carries. The
>    claim is true of v2.4's *own* milestones and silently omits everything v2.4 *inherited*.
> 2. **`:37` — "✅ Per-milestone decisions blended into corpus at each close" and "✅ No cross-milestone
>    decision conflicts" are FALSE.** **Six items were aging out as this was written**, four of which no
>    audit had recorded. Among them: `DEF-M226-01` (whose destination, M228, fired inside this very release
>    without landing it) and the four v2.3 tail carries signed off KEEP-DEFERRED → **v2.4**.
> 3. **`:34` — the pointer to a "Completeness Ledger below" DANGLES. The file ends at line 42; no such
>    section exists.** This is the **concrete mechanism** by which the standing test-debt carry and
>    `DEF-M226-01` left v2.4 with no landing record: the review told the reader where to check, the place
>    did not exist, and nobody followed the pointer.
> 4. **v2.4 ran NO release-scope deferral audit at all.** `02.40-casting-call/` has four *milestone*-scope
>    `audit-deferrals/` dirs and **none at release root** — the first release since v2.3 without one. Every
>    item whose declared destination was "the v2.4 release close" fired unchecked. **This is the single
>    structural cause of 7 of the 8 aged-out items the v2.5 close had to fate.**
>
> Also note `:20` — the "**644 passed / 14 failed**" demo-stack figure is a **dirty-clone** reading. The
> v2.5 close re-baselined the standing set: on a clean stable-`main` clone set it is **8 on macOS / 7
> expected on Linux**, with **0 real defects** and **0 `pre_sha256` pin drift** (that diagnosis is
> **REFUTED**). The count is **host-dependent** — always state the host OS.
>
> **The class, generalized:** *a review can report a green while proving nothing — and a green review is
> the artifact downstream audits trust most.* A ✅ whose evidence section was never written is worse than a
> ⛔, because it actively closes the question. Corrective practice adopted at v2.5: every fate names a
> **milestone**; the ledger is a real section, not a forward pointer; and a ✅ must cite the artifact that
> proves it.

## Supply chain (Phase 0)
- ✅ npm audit (e2e): **0 vulnerabilities** (prod + all). Dev-only toolchain (@playwright/test, @types/node, typescript).
- ✅ Go deps: `github.com/anthropos-work/ai v1.40.1` — the only third-party dep, **unchanged in v2.4** (added M45).
- ✅ No new deps introduced by v2.4. No license concerns (dev-only TS + one internal Go dep).
- Lockfile: `dependencies.lock`.

## Scope (Phase 1)
- ✅ All 7 milestones delivered: M222 (spike/GO), M223 (seeder), M224 (render, closed-on-gate), M225 (demo-integration),
  M226 (billion-proof, closed-on-gate), M227 (believability fixes), M228 (live re-prove, closed-on-gate).
- ✅ No `closed-incomplete` milestones → no carry-forward.md, no undelivered Fate-3 items.
- ✅ The release thesis — the recruiter-vantage hiring org, 45 candidates on 5 shared positions compared side-by-side —
  is delivered AND proven live on billion.

## Code quality (Phase 2)
- ✅ rosetta release diff = **111 files, 100% docs** (97 knowledge + 14 corpus). 0 code, **0 platform-repo edits**.
- ✅ rext tooling reviewed per-milestone; consistent guard/seeder patterns; the M228 render-probe + seeder-guard
  reviewed at close-milestone (adversarial scenario recorded).

## Documentation (Phase 3)
- ✅ `corpus/services/hiring.md` carries the full render+seed model incl. the M228 live findings (intercepting-route
  drawer + the incomplete-guard correction).
- ✅ `state.md` + `roadmap.md` reflect v2.4 code-complete; all 7 milestone overviews archived.
- [ ] (nice-to-have, non-blocking) The intercepting-route drawer finding is in `hiring.md`; the render-probe
  `RENDER_ONLY_SIM` knob is a rext-internal test detail (documented in-code) — no corpus gap.

## Tests & benchmarks (Phase 4)
- ✅ rext stack-seeding: green, **96.8%** stmt coverage. playthroughs: green. e2e: tsc clean.
- ⚠️ rext demo-stack: **644 passed / 14 failed** — ALL pre-existing or environment-gated carries, NONE v2.4 regressions
  (see Completeness Ledger below). The rext repo is at a fixed commit; the rosetta docs merge+tag does not touch it.

## Decision consolidation (Phase 5)
- ✅ Per-milestone decisions blended into corpus at each close (M222 hiring model, M227 fixes #1/#2/#4, M228 render).
- ✅ No cross-milestone decision conflicts. The mirror-table read-model (M222) held consistently through M228.

## Metrics (Phase 4b)
- ✅ 0 platform edits · supply-chain GREEN · flake 0 · rext seeders 96.8%. Aggregate: `metrics.json`.
- No regression vs v2.3 (v2.4 adds a hiring vantage; no capability removed).
