# M253 — Progress

Iterative milestone (perf, measure→patch→re-measure). Primary metric: first-meaningful-paint < 1000 ms + no blank > 1 s,
p95 over 5 consecutive cold loads on a cold demo (state the environment — laptop vs tailnet), gated on a fresh-green
`autoverify.json`.

## Running ledger

- iter-01 (tok/bootstrap): authored TOK-01 (shell-before-awaits + no-thirdparty demopatches + FCP runner); baseline skeleton-visible 4669 ms (demo-2, laptop); dominant await = canAccess (~3.9 s), not clerk.load — see iter-01/progress.md
- iter-02 (tik): shipped both demopatches + FCP runner + extended M249 ladder (rext @ b8969c0, tag pushed); rebuilt studio image on demo-2 → skeleton-visible p95 **817 ms** (was 4669 ms), 5/5 cold loads, 0 login bounce — numerical gate MET; fresh-green cold confirmation → M254 — see iter-02/progress.md
- iter-03 (tik/cleanup): landed the 3 docs Delivers (latency-budget.md studio budget · demopatch-spec.md 2 patches, count 21→23 · studio-desk.md MPA boot model). Gate MET (M253 local-bootstrap charter); fresh-green COLD confirmation → M254 — see iter-03/progress.md

## M253: Final Review (close-milestone)

🔍 **M253 review found 2 findings:** 0 scope · 0 code-quality (no code in this tree — the demopatches + FCP runner are code-of-record in rext at the tag) · 1 docs · 0 tests · 1 decision-triage. Addressed all fully (no partial fixes) — no sign-off needed (gate-met iterative close).

### Scope
- [x] Gate-distance + iter-ledger audit — 3 iters accounted (1 tok bootstrap + 2 tiks), all closed-fixed; numerical gate MET; 1 Fate-2 carry-forward (CARRY-M253-01 → M254).

### Code Quality
- [x] No code in this tree — the 2 demopatches + the `run-studio-fcp.sh` runner are code-of-record in rext @ `july-jitter-m253-studio-first-paint` (b8969c0, on origin). `bash -n` clean + `demopatch check` PASS were verified in iter-02.

### Documentation
- [x] [should-fix] `demopatch-spec.md` §2.1 stale illustrative count — "R1 sweep iterates … (all 21 today)" left behind by M253's canonical 21→23 inventory reconcile → fixed to "all 23 today". Historical counts (14 / "the other 11" / "swept 14 manifest(s)") correctly unchanged. Verified the 3 doc Delivers accurate + the latency-budget ↔ demopatch-spec ↔ studio-desk cross-ref triangle resolves (0 broken links).

### Tests & Benchmarks
- [x] No rosetta test suite (docs + plan only, per M248/M251). The rext FCP runner + `TestPatchInventory` re-pin (23-total / 5-studio-desk) live at the code-of-record tag. Benchmark PASS: skeleton-visible p95 817 ms < 1000 ms (5 cold loads, demo-2 laptop) — recorded in metrics.json.

### Decision Triage
- [x] iter-01 D1 (inline KB-fidelity) · D2 (dominant-await = canAccess) · iter-02 D3 (chained-manifest sha) · D4 (lib-only rebuild vehicle) · D5 (green-gate non-achievable on warm demo-2 → M254) · TOK-01 (bootstrap strategy) → **archive** (maintainer-only rext mechanism + bootstrap strategy). Their load-bearing platform facts (per-leg baseline, the reorder fix, the chained patch pair, the MPA boot model) were **already blended** into the 3 corpus doc Delivers during iter-03 — verified accurate, no duplication.

## Gate Outcome Ledger (Phase 9-iter)

**Close status:** `closed-on-gate`

### Gate
- **Target:** cold demo (state environment — laptop vs tailnet), first-meaningful-paint < 1000 ms (the `.page-skeleton` header+sidemenu shell visible) AND no blank > 1 s, p95 over 5 consecutive cold loads; never gate on `networkidle`; always gate on a fresh-green `autoverify.json`.
- **Achieved:** skeleton-visible **p95 817 ms** (p50 743, max 817; samples 817/795/480/539/743 ms) over 5 consecutive cold loads on **demo-2 (LOCAL LAPTOP)**, 5/5 reached the shell, 0 login bounces. Baseline 4669 ms → **~5.7× faster**.
- **Distance:** numerical gate **MET** (p95 817 ms < 1000 ms, max ≤ 1000 ms) on the local-bootstrap charter. The fully-green COLD-p95 confirmation on billion (fresh-green `autoverify.json`) is chartered to M254 by **coordination rule 9** (two live-measured iteratives can't share billion RAM).
- **Status:** **`closed-on-gate`** — the gate fired on M253's local-bootstrap charter.

### Iter ledger summary
- **iter-01** (tok/bootstrap): authored TOK-01 (shell-before-awaits + no-thirdparty demopatches + FCP runner on the M249 ladder); baseline skeleton-visible 4669 ms; dominant await = `userService.canAccess()` ~3.9 s (NOT clerk.load — 140 ms). closed-fixed.
- **iter-02** (tik): shipped both demopatches + the extended `build_frontend_studio_desk` ladder (5-manifest fingerprint) + net-new `run-studio-fcp.sh`; rext @ b8969c0 tagged + pushed; rebuilt studio image on demo-2 → skeleton-visible p95 **817 ms**, 5/5 cold, 0 bounce. Numerical gate MET. closed-fixed.
- **iter-03** (tik/cleanup): landed the 3 docs Delivers (latency-budget.md studio budget · demopatch-spec.md 2 patches 21→23 · studio-desk.md MPA boot model); cross-refs wired both ways. Gate-met exit-1. closed-fixed.
- All 3 iters closed; every commit maps to an iter (one-commit-per-iter); no orphan iters/commits.

### Routes carried forward (Fate 2/3/escape-hatch)
- **CARRY-M253-01 → Fate 2 (confirmed-covered by M254).** The fresh-green COLD-p95 confirmation on billion (re-measure the studio FCP gate on a freshly brought-up, fully set-dressed cold demo with a green `autoverify.json`). **M254 exit gate part (f) "studio first-paint < 1 s cold p95 (← M253)" already owns it verbatim** — no M254 `overview.md` edit needed. This is the milestone's deliberate coordination-rule-9 split, NOT a gate miss. Durable record: `carry-forward.md`.

### Dropped
- None.

### Protocol evolution
- None. The measure→patch→re-measure loop of `latency-budget.md` applied verbatim; the milestone confirmed the "clerk.load 10 s timeout" hypothesis was a red herring (140 ms actual) and the blank is a `userService.canAccess()` GraphQL-404 retry ladder — resolved to a pure paint-ordering fix independent of the 404 (out of scope).
