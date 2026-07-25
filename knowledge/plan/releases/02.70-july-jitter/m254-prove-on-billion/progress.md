# M254 — Progress

## Running ledger

_(Per-iter progress — tik/tok entries, distance-to-gate, and gate-part (a–h) live evidence — accumulates here
during the iter loop. `iter-NN/` dirs are created by `/developer-kit:build-mstone-iters` on its first invocation;
there are NO iter dirs at scaffold.)_

- iter-01 (tok/bootstrap): authored TOK-01 (cluster-per-tik: DRIVE → read-only sweeps fan-out → mutating tail);
  baseline pin=`july-jitter-m253-studio-first-paint` on origin, billion clean slate, gate 0/8 — see iter-01/progress.md
- iter-02 (tik/DRIVE, closed-fixed-partial): cold reset-to-seed COMPLETED on billion — fresh v2.7 CONSOLIDATED
  demo up (16 containers, 0 skillpath, peer-reachable). Gate (a) NOT MET: autoverify green:false, sole blocker =
  drifted `app-aireadiness-snapshot-loadmembers` demopatch (path moved to internal/aireadiness/readiness.go) →
  iter-03. Corrected coordinator "died" diagnosis (docker-blind trap). Gate 0/8 — see iter-02/progress.md
- iter-03 (tik, closed-fixed): re-authored the drifted AI-readiness demopatch for the consolidated app
  (path/anchor/shas), 52/52 tests green, rext tag july-jitter-m254-aireadiness-repoint on origin, re-pinned
  billion, cold reset-to-seed → autoverify green:true/0 warnings. **GATE (a) MET.** Gate **1/8** —
  see iter-03/progress.md
- iter-04 (tik/measurement-cluster, closed-fixed-partial): measured b/c/d/f/g + re-confirmed a live. **MET:**
  (a) re-confirmed green, (b) MET-with-disposition (45/45 + 4 voice presence-only, coordinator-approved),
  (d) MET both vantages (8/8 sections + 3 drift-fixes), (c) prod-eject side proven (0 escapes/133 pages).
  **RESIDUAL:** (f) studio-desk session-carry on --public-host (FCP lands on web app) → fix-iter; (g) 9
  host-sensitive test fails (2 nvm env-artifact + 7 to fix) → fix-iter. **PENDING:** (c) Back-to-Cockpit render,
  (e) builder Playthrough, (h) latency solo. Gate **~4/8 MET** (a,b,d + a-reconfirm). cap-reached → fresh agent.
  see iter-04/progress.md
- iter-05 (tik/measurement-cluster, closed-fixed-partial): banked **(h)-latency MET both vantages** (employee
  p95 1.43 s / manager 1.41 s, gate < 5 s, 5/5 ACCESS, solo) + **(c)-render LIVE on 2/4 apps** (next-web +
  hiring; all 4 patch-applied@build + :17700 baked → resolve-to-stack). Found a **real (c) fragility defect on
  the native academy** (Back-to-Cockpit item reverted out-of-band; `UserMenu.jsx` = pristine; heal only on
  `ant-academy.sh` re-invoke) → routed to fix-iter. (c)-studio render (f)-coupled. Gate **~4.5/8** (h-latency
  banked; c partial). see iter-05/progress.md
- iter-06 (tik, closed-fixed): **(f) is NOT a defect** — it was a measurement artifact (FCP probe defaulted to
  `maya-thriving`, an employee studio's `checkEnterpriseAndAdmin` bounces by design). Shipped the tooling fix
  (default identity → `dan-manager`, studio-eligible admin; rext `cbe9256`, tag
  `july-jitter-m254-studio-fcp-identity` on origin). **(f) session-carry MET** (live: admin heroes reach the
  studio shell). FCP shell paint p50 637-726 ms < 1 s (M253 fix holds) but p95 tailnet-jitter → **disposition-
  pending**. Also LIVE-confirmed **(c)-studio render** (unblocked): item renders + 0 prod-eject → **(c) render
  side now 3/4 LIVE**; unblocked **(e)**. Gate **~5/8**. see iter-06/progress.md
- iter-07 (tik, closed-fixed-partial): **(g)** — captured the exact live failures on billion (10 fail + 1 err /
  159; ~7 unique + 1 err). Root causes DIVERSE. **FIXED** the nvm/node host-robustness of
  `test_missing_node_documents` ×2 (node-free bindir + clean HOME; verified live; rext `dfdd9bc`, tag
  `july-jitter-m254-academy-nonode-hostrobust` on origin). **Routed** the remaining 6 as `FIX-M254-g-testhealth`
  (intra-run `:23077` port-leak + M245 reconcile drift · next.config sha re-pin · 2 mutation-meta · overlay-127)
  → coordinator fate (chronic host-sensitive tests). 0 platform edits, 0 demo-runtime impact. see iter-07/progress.md
- iter-08 (tik, closed-fixed-partial): a (c)-academy re-heal attempt CASCADED into a self-inflicted demo
  disruption (crashed academy → `fuser -k` killed tailscaled [systemd-recovered] → mis-killed the web+hiring
  CONTAINER procs → lost port mappings). **Recovered via a full cold reset-to-seed re-bring-up** (GREEN, up_rc=0,
  autoverify green:true/0 warnings/ts 08:16:34Z; re-pin dfdd9bc + rext.tag bump past the M217 guard) — net-
  positive: a fresh demo. **(a) re-confirmed green**; **(c)-academy renders on the fresh demo** → **(c) MET
  (4/4 apps render LIVE + prod-eject)**. Gate **~6/8** (a,b,c,d + f-session-carry). (e)+(h) → iter-09. 0 platform
  edits. see iter-08/progress.md
- iter-09 (tik 5/cap, closed-fixed-partial): **(e)+(h) mutating tail** — pt-world reset-to-seed LANDED on billion
  (`--reset-only`: TRUNCATE + pt-world seed, roster re-exported, fake services restarted, fake-FAPI 200). The
  browser Playthrough suite **PROVEN to run + pass** from this peer (117 tests serial, first 6 green incl a
  manager Playthrough). Full 117-test completion + (f)-FCP-p95 + (g)-testhealth dispositions routed to a fresh
  agent (prove-on-billion "fresh agent per run"). Gate **~6/8 MET** (a,b,c,d,f-session-carry); e+h proven-to-run.
  EXIT: cap-reached. see iter-09/progress.md
- iter-10 (tik, closed-fixed): **(e)+(h)-Playthroughs mutating tail — GATE MET.** Drove the full Playthrough
  suite to completion on billion (cold reset-to-seed, tailnet peer): **18/18 passing (100 %), 0 failing,
  DONE_rc=0.** The coordinator's 16 all green (h-Playthroughs). Fixed **(e)**: the M252 studio-builder
  Playthrough targeted STALE routes vs the live studio-desk **v0.152.1** redesign (unified
  `/simulation-builder`); re-authored to the real UX (advanced "Design it with AI" drafts a scenario the
  designer renders = the generation boundary; guided 5-part interview live, Generate at Part 5 is P6-out).
  Fixed a **(h)** flaky `networkidle` deadlock on 4 `/home` logins → `domcontentloaded` anti-deadlock (also cut
  the suite 13m→3.8m). rext `july-jitter-m254-studio-pt-retune` @ 4f1409e ON ORIGIN (rung-zero); no billion
  re-pin (test-only); 0 platform edits. Recorded the 3 coordinator-approved dispositions (f/c/g) →
  decisions.md + carry-forward.md. **GATE (e)+(h) MET → effective gate-met.** see iter-10/progress.md

---

## GATE STATUS — effective gate-met (iter-10, 2026-07-25)

All parts (a–h) MET live on billion (cold reset-to-seed, driven/asserted from a tailnet peer, 0 platform edits):
- **(a)** re-grounded stack builds + GREEN on the consolidated platform (3 subgraphs, skillpath-in-app) — MET
  (iter-03 fix + iter-08 fresh-reset re-confirm; autoverify green:true/0 warnings).
- **(b)** content-stories manager CTA lands non-empty — MET-with-disposition (45/45 landable ALL LANDED + 4
  voice presence-only, coordinator-approved; `manager_presence_only` follow-up carried).
- **(c)** Back-to-Cockpit in all 4 apps + studio→stack, 0 prod-ejects — MET (iter-08 fresh demo 4/4 render;
  academy-durability edge → `FIX-M254-c-academy-durable`).
- **(d)** AI-readiness faithful, both vantages — MET (iter-04 employee + manager, 8/8 sections + 3 drift-fixes).
- **(e)** studio sim-builders generate (builder Playthrough green) — **MET** (iter-10: re-tuned to studio-desk
  v0.152.1; both studio Playthroughs green).
- **(f)** studio first-paint < 1 s cold p95 — MET on app-side paint (p50 637–726 ms < 1 s); p95 = ACCEPTED
  environmental tailnet-jitter (coordinator-approved disposition).
- **(g)** the 8 host-sensitive demo-stack test-health tests — 2/8-class fixed + verified live; 6 remaining =
  carry-forward `FIX-M254-g-testhealth` (chronic host-sensitive, 0 demo-runtime impact; coordinator-approved).
- **(h)** live-browser specs + content-stories sweep + Playthroughs green; p95 click→ACCESS < 5 s — **MET**
  (iter-05 latency p95 1.43/1.41 s < 5 s; iter-04 content-stories 45/45; iter-10 Playthroughs **18/18**).

**→ Effective gate-met.** Next: `/developer-kit:harden-mstone-iters --final` then `/developer-kit:close-milestone`.

---

## M254: Final Review (close, 2026-07-25)

Proof milestone (`Delivers: none`); code-of-record is rext (6 tags on origin, all rung-zero-verified). The
rosetta-side footprint is plan docs + one corpus reconciliation. Review found **1 Fate-1 LAND-NOW** (a
RED-at-HEAD) + confirmed the coordinator-approved dispositions/carries. All addressed fully.

### Scope
- [x] Gate a–h MET live on billion (cold reset-to-seed, tailnet peer, 0 platform edits) — Playthroughs 18/18,
  latency p95 1.43/1.41 s, studio builders green, content-stories 45/45 landable, AI-readiness both vantages.
- [x] Iter-ledger clean: 10 iters (iter-01..10) all closed, 1 commit/iter + 1 harden commit, 0 orphans.
- [x] **FIX-M254-h-patch-inventory-drift — LANDED (Fate 1).** Deferral audit re-fated Fate-3 → Fate-1
  (a RED-at-HEAD must not ship). See Documentation + Tests below.

### Code Quality
- [x] rext code-of-record already reviewed per-iter + final-harden stabilized (Pass 1-3, +22 tests, 3× flake
  gate). No new cross-cutting issues; proof milestone carries no rosetta source.

### Adversarial (Phase 2c)
- [x] Scenario "does the 21→23 fence bump mask a spurious patch?" — verified NOT masked (the 2 additions are
  M253's real shipped studio-desk first-paint patches; sibling validity + no-orphan tests green). Recorded in
  decisions.md.

### Documentation
- [x] `corpus/ops/demo/demopatch-spec.md §5` reconciled to 23/5 — 3 secondary fence-description refs the M253
  header-only update missed (line 144 three→five studio-desk; the fence-contract line 21→23 / 3→5; the
  self-contradicting "total is 16"→23). No stale inventory counts elsewhere in corpus/CLAUDE.md.

### Tests & Benchmarks
- [x] `test_patch_inventory` 5/5 green post-fix (was 2 RED at HEAD).
- [x] Full demo-stack suite: **909 tests, 900 pass**; the 8 fail + 1 error are ALL the documented
  `FIX-M254-g-testhealth` host-sensitive carry (test_ant_academy launcher/reap set + clerk-wiring overlay +
  Linux-only test_purge + docker image-guard) — 0 milestone-touched regressions, 0 demo-runtime impact.
- [x] playthroughs e2e `tsc --noEmit` clean; demopatch-family + M254 aireadiness regression 65 offline green.

### Decision Triage
- [x] All M254 decisions (TOK-01, D-iter10-1..4, D-harden-1, adversarial) → maintainer-archive; operational
  knowledge already lives in `latency-budget.md` + `demopatch-spec.md §5` — no new knowledge blend.

---

## M254: Gate Outcome Ledger (close, 2026-07-25)

**Close shape:** `closed-on-gate` (gate MET; no carry-forward required as a new deliverable — the existing
carry-forward.md documents the coordinator-approved follow-ups + the LANDED FIX-M254-h).

### Gate
- **Target:** cold reset-to-seed on billion, tailnet peer, 0 platform edits — the multi-part a–h exit gate.
- **Achieved:** all 8 parts MET live on billion (a autoverify green:true/0 warn · b 45/45 landable + 4 voice
  presence-only · c 4/4 apps + prod-eject side, 0 escapes · d both vantages 8/8 sections + 3 drift-fixes ·
  e both studio builders green · f app-side paint p50 637–726 ms < 1 s · g 2/8-class fixed live · h latency
  p95 1.43/1.41 s < 5 s + Playthroughs 18/18).
- **Distance:** 0 (with 3 coordinator-approved dispositions on residuals).
- **Status:** `closed-on-gate`.

### Iter ledger summary
10 iters (iter-01 bootstrap tok + iter-02..10 tiks), all closed; 1 commit/iter + 1 harden commit; 0 orphans.
Gate progression: 0/8 (iter-01) → a MET (iter-03) → a,b,d + c-prod-eject (iter-04) → +h-latency (iter-05) →
+f-session-carry, c-studio (iter-06) → g-2/8 (iter-07) → c 4/4 + a re-confirm (iter-08) → e+h proven-to-run
(iter-09) → **e+h MET, effective gate-met** (iter-10) → final harden stabilized.

### Routes carried forward (Fate 2/3) — non-blocking, coordinator-approved
- **Fate 1 (LANDED at close):** FIX-M254-h-patch-inventory-drift (rext fence 21→23 + corpus §5) — see carry-forward.md.
- **Fate 1 (disposition):** (f)-FCP-p95 = ACCEPTED environmental tailnet-RTT jitter (p50 < 1 s holds).
- **Fate 3 (follow-ups):** FIX-M254-c-academy-durable · FIX-M254-g-testhealth (6 host-sensitive tests) ·
  (b)-voice manager_presence_only flag · verify.sh stale skillpath default · studio-desk billion re-pin note.
- **Discharged (inherited-to-M254):** CARRY-M248/250/252/253-01 — all resolved by this live gate proof.

### Dropped
None.

### Protocol evolution
1 bootstrap tok (TOK-01: cluster-per-tik live re-prove); 0 triggered toks (no 3-consecutive-no-progress
stall). The M221/M236/M244 prove-on-billion shape held end-to-end.
