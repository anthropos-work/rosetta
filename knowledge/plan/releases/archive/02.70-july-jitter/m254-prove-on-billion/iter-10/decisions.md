# M254 · iter-10 — decisions

## D1 — full Playthrough suite driven to completion on billion (the (e)+(h) proof)
Re-reset pt-world on billion (`--reset-only`, devops driver on-host: full FK-ordered TRUNCATE + fresh
pt-world seed + 30-identity roster re-export + fake-service restart + fake-FAPI 200), then drove the browser
suite from THIS tailnet peer (`PT_HOST=billion.taildc510.ts.net PT_APP_SCHEME=https ./run-playthroughs.sh 1`).
The coordinator's **16 live Playthroughs** (6 employee + 4 manager + 4 AI-readiness + hiring-recruiter +
assignment-assign) **ALL GREEN** on the first full run → gate **(h)-Playthroughs** half MET.

## D2 — gate (e) studio builders: stale-Playthrough drift, re-authored to studio-desk v0.152.1 (FIXED)
The 2 studio-builder Playthroughs (`pt-studio-advanced-generate` + `pt-studio-guided-generate`) FAILED —
the ONLY 2 fails (16/18). Root cause (decisive DOM evidence): the studio builders **render + function live**
(Morgan/pt-manager logged in, "Go to Anthropos"→:13000, both surfaces hydrated) but the **M252 Playthrough
tested STALE routes** (`/sim-advanced-builder` + an immediate `Generate` button). studio-desk on billion is
**v0.152.1 (2026-07-03)** — a REDESIGN that **predates M252** (designed 2026-07-23): both builders now land on
a unified `/simulation-builder` entry. The M252 page object was authored-blind ("matchers chosen without a
live render in hand; the orchestrator live-tunes any locator that misses the real DOM") and **never
live-tuned** — M254 (prove-on-billion) is where that tuning belongs.
Re-authored (live-tuned against billion): ADVANCED = `/simulation-builder` → describe scenario → "Design it
with AI — Advanced mode" → the AI **drafts a full scenario** (characters Markus Vogel + Sofía Alvarez, mission
tasks) and the advanced designer renders it at `/sim-advanced-builder` (the TRUE generation completion
boundary — a *stronger* gate-e proof than the old button-click). GUIDED = `/sim-guided-builder` 5-part
interview live + interactive (Part-1 goal question landmark; Generate is at Part 5 behind a live-LLM
interview, P6-out). Page object gained new locators; old unit-tested matchers kept (5 unit specs stay green);
manifest prose updated (ids stable). BOTH green.

## D3 — gate (h) skillpath-legacy networkidle flake → anti-deadlock fix (FIXED)
Full run #2 (post-studio-fix): 17/18 — the 2 studio GREEN, but `pt-skillpath-legacy` hit a **120 s
`networkidle` timeout** on the `/home` login (flaky: PASSED run #1 @1.7 m, tipped >120 s run #2 on tailnet
jitter; UNRELATED to the studio-only edits). Root cause: **4 `/home`-landing logins default to `networkidle`**
— but `/home` hosts the member AI-readiness POLLING surface that never idles (the exact "NEVER gate on
networkidle" hazard the protocol + `PageObject.goto` doctrine forbid). Applied `waitUntil:'domcontentloaded'`
to all 4 (`skillpath-legacy` + `aireadiness-member-done`/`-progress` + `aisim-chat-launch`) — the assertions
auto-retry, so it is strictly safer. Also collapsed the suite wall-clock (~13 m → 3.8 m) by removing the
~100 s networkidle stalls.

## D4 — definitive full suite 18/18 GREEN + rung-zero (gate (e)+(h) MET)
Run #3 (cold reset-to-seed): **117 passed, DONE_rc=0 — Playthroughs 18/18 passing (100 %), 0 failing,
0 unimplementable**, `ptreport --gate no-regressions` PASSED. rext committed + tagged
`july-jitter-m254-studio-pt-retune` (4f1409e) + **`git push --tags` verified ON ORIGIN** (rung-zero). **No
billion re-pin** — the Playthrough suite runs from the local rext clone (which carries the fix) against
billion's *unchanged* served app; the fix is test-only (no demopatch/seeder change), so billion's demo build
at `dfdd9bc` is unaffected. **0 platform-repo edits.**

## D5 — the 3 coordinator-approved dispositions recorded (see milestone-root decisions.md)
Recorded (f)-FCP-p95 = ACCEPTED-environmental, (c)-academy-durability = Fate-3 academy-durable follow-up,
(g)-testhealth = carry-forward `FIX-M254-g-testhealth`, into the milestone-root `decisions.md` +
`carry-forward.md`. They get formally fated at close-milestone's deferral audit.
