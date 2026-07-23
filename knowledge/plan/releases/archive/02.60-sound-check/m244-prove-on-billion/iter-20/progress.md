# iter-20 — progress

**Type:** tik (run 8, under TOK-03 — gate (h) COMPLETION, final-push step 1)

Gate (h) = "every v2.6 fix proven live; p95 click→ACCESS < 5s hero vantages." 4/6 fixes were proven as
iter-18 byproducts (talk-to-data / library / cockpit UX / content-fidelity). This tik closed the residual:
the p95 gate + the last two fixes (academy course-start + M241 EN/IT language toggle) — all live on billion.

## What landed (all live on billion, from this tailnet peer)

### 1. Autoverify REFRESHED (the stale >12h verdict → fresh green)
Ran `autoverify.sh --project demo-1` on billion (as devops, STACK_DIR set). **0 warnings, green:true, ts
2026-07-23T00:34:22Z** (age ~9s when the latency gate read it). All probes: backend /api/health 200 ·
sentinel.casbin_rules=1250 · directus.directus_collections=21 (per-stack-local, not prod) · verify live all
pass · demo-patches all applied · frontend builds this-run · public.skills=42790 · cockpit :17700 ·
fake-FAPI :15400 · hiring org 5 positions + 42 sessions. scp'd locally → the latency green-gate.

### 2. p95 click→ACCESS < 5s — BOTH hero vantages PASS
`run-latency.sh 1 <vantage>` from this workstation, LATENCY_HOST=billion.taildc510.ts.net,
LATENCY_SCHEME=https, LATENCY_GATE_MS=5000, green-gate = the fresh scp'd autoverify.json (age ~9s; the M236
UTC-vs-local age-check bug is fixed in the local authoring copy 6aacc32 — `TZ=UTC` BSD branch present):
- **employee (maya-thriving): p95 1.46s** (p50 0.69s) — 5/5 reached ACCESS. gate < 5.0s ✅
- **manager (dan-manager): p95 1.31s** (p50 0.80s) — 5/5 reached ACCESS. gate < 5.0s ✅
(The `net::ERR_ABORTED …?_rsc=` lines are benign Next.js RSC-prefetch cancellations on navigation, NOT
ACCESS failures — every run reached ACCESS.)

### 3. academy course-start + academy language switch — FRESH browser-grade harness verdict
Fresh employee `run-coverage.sh 1 employee` (COVERAGE_HOST=billion.taildc510.ts.net, COVERAGE_APP_SCHEME=https)
against billion → **GATE MET** (reachable 62/150, failingSections 0 / personaFailures 0 / escapes 0 /
notReached 0 / crossPortFollowFails 0, frontier EXHAUSTED). The cross-port follow ran `probeAcademyChapter`
live: **"ant-academy home + chapter body (slug accessing-open-models) + ?lang=it OK
(billion.taildc510.ts.net:13077, HTTP 200)."**
- **academy course-start** (M238 `academy-fs-published-chapter-body`): the EN chapter body renders HTTP 200
  — NOT the "You wandered off the trail" 404. (Raw-curl cross-check: chapter `<title>How to Access Open
  Models — Literacy & Landscape — AI Academy</title>`, 386KB body; the raw-HTML grep hit on the 404 string
  is a false positive — that text ships in the Next.js error-boundary bundle regardless of render, which is
  exactly why the harness asserts the RENDERED `main,body` region with a real browser, not a grep.)
- **academy language switch** (`?lang=it`): the same chapter re-renders HTTP 200 under `?lang=it` (raw-curl:
  528KB, `lang="it"`). The M238 half-#2 shares the #3 backend-null path; both green.

### 4. language toggle (M241 EN/IT content-stories) — LIVE
billion's `content-manifest.json` (projected at bring-up) carries the M241 language axis: 28 content-story
sessions across 4 products, each with a `language` field (english/italian) + `lang_toggle`. **21 bilingual
tuples** (`lang_toggle=True`; EN 10 / IT 13). The presenter cockpit (:17700, curled from the peer vantage)
renders the **"Content stories" tab** + the **"EN/IT"** toggle label + **28 `data-lang` per-cell markers** +
English/Italian labels — the M241 per-cell EN/IT toggle rendered live.

## Re-measure
- **Pre-iter metric:** 5/8 gate parts (a,b,d,e,g).
- **Post-iter metric:** **6/8** — gate (h) now COMPLETE (all 6 v2.6 fixes proven live + p95<5s both heroes).
- **Delta:** +1 (met the expected lift exactly).
- Remaining: gate (f) 3 drift-carries → 7/8; gate (c) 16 Playthroughs LAST → 8/8 = GATE MET.

## Close — 2026-07-23

**Outcome:** gate (h) COMPLETE — metric 5/8 → **6/8**. Autoverify refreshed green; p95<5s both heroes
(employee 1.46s / manager 1.31s); academy course-start + `?lang=it` fresh browser-grade OK on billion; M241
EN/IT language toggle live (manifest 21 bilingual + cockpit EN/IT render). 0 platform edits; 0 code changes
(pure live verification of shipped fixes).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (6/8 — gate f + gate c remain)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik 1/5 this run) — (6) protocol-stop: n — Outcome: **continue**
**Decisions:** D1 (iter-20 decisions.md — the academy-curl false-positive + why the harness browser assert is authoritative).
**Side-deliverables:** none.
**Routes carried forward:** gate (f) 3 drift-carries (BURNIN-M221 needs a `/dev-up --public-host` remote dev burn-in; F-M220-4; PROBE-M218-c3) → iter-21; gate (c) 16 Playthroughs LAST on pt-world → after gate (f); DEF-M239-01 as budget.
**Lessons:** a raw-HTML grep of a Next.js SSR page is NOT a valid render assertion — the error-boundary ("wandered off the trail") text ships in the client bundle on EVERY page; only a browser assert on the rendered region (the harness `assertSection`) distinguishes a rendered chapter from the 404. The academy course-start proof MUST use the browser probe, never curl-grep.
