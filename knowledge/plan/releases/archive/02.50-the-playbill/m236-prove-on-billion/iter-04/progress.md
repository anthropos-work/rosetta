# iter-04 — tik: Phase H (calibrate + author the content-stories harness)

**Type:** tik · **Shape:** declared multi-step (calibrate → author → prove → doc), per the tripwire carve-out.

## Step 1 — calibration (the step M235 refused to let us skip)

Logged in as real content seats on the live stack and captured what the pages **actually** render. Three
genuinely different shapes turned up — and the differences are exactly what a blind author would have
gotten wrong:

| Probe | Seat | Observed |
|---|---|---|
| `asmt-voice-pass` (interview) | `content-player-23` | **205 chars** — *"Interview completed … will be available for your organization's manager at Cervato Systems to review"*. **Correct and terse by design**: the player is not shown a report for an interview. |
| `asmt-code-pass` (scored) | `content-player-25` | **1888 chars** — *"Congrats, Clara! You did it!"*, **100/100**, a full LLM feedback paragraph (in Italian — the real prod session's language, faithfully cloned), and **Evaluated Skills: Prompt Engineering, Clear Messaging**. |
| manager view of the same | `dan-manager` | Header renders (*"…'s Results for Business Development Manager: craft a winning client proposal"*, *"2 skills measured"*) but the attempts table says **"No data"** and the player's name renders **"undefined undefined"**. |

**Had this been authored blind**, the natural descriptor — "a result page has ≥N chars and a score" — would
have reported the interview page (a *correct* render) as a failure, and would have graded the manager page
as a **pass** on its header alone, hiding the real defect. That is precisely the "INCORRECT, not merely
uncalibrated, load-bearing harness" USER-BLOCKER-M235-02 warned about.

## Step 2 — authoring, and a scope that SHRANK

The audit's B3 framing (13 seats → 13 manifests) assumed the `VantageManifest`/BFS-crawl machinery, whose
`identityKey` is singular. **It is the wrong tool.** `cockpit-login.loginAs()` already takes a
**`landingPath`**, so a content seat can be landed *directly* on its exact result URL — the sweep is an
**exact-path visit per (seat × path)**, with no crawl and no manifests at all.

Shipped in `rosetta-extensions` @ `playbill-m236-content-sweep`:

| File | Role |
|---|---|
| `stack-verify/e2e/lib/content-result-page.ts` | the result page-object + per-shape assertions (**authored from scratch** — `AISimulationResultContainer` is a next-web `.tsx` component, not a harness object) |
| `stack-verify/e2e/tests/content-stories.spec.ts` | the sweep; enumerates landable pairs from the served manifest |
| `stack-verify/e2e/run-content-stories.sh` | runner; local or `--host <magicdns>` for a remote stack |
| `stack-verify/e2e/aggregate-content.py` | turns the per-pair ledger into the reading |

**Two harness bugs found and fixed during the first runs** — both worth remembering because both produced
a *confidently wrong number*:

1. **`mode: 'serial'` aborted the sweep** at the first failure (29 tests "did not run"). Serial makes tests
   *dependent*; the question "how many of 31 land?" silently became "did pair #2 fail?". Fixed to
   independent tests + `--workers=1` (which gives the fake-FAPI seat-race safety **without** coupling
   outcomes).
2. **Playwright restarts its worker after every failure**, re-importing the spec module — so the in-memory
   results array reset and `afterAll` fired once per restart, printing **"LANDED 1 / 31" over and over**,
   each time from a fresh empty array. Fixed by appending one JSON line per pair to `pairs.jsonl` and
   aggregating in the **runner**.

**And a false PASS.** The first full reading graded skill-path pages with the scored-simulation shape. The
*in-progress* path failed ("no evaluated-skills section") — but the *completed* path **passed**, not because
it rendered a report, but because 11k chars of legitimate content happened to contain the word "feedback".
Calibrating the real shape (`player-skillpath`: chapters + a progress signal, selected **by route** rather
than by keyword) fixed both. **A gate that passes for the wrong reason is more dangerous than one that
fails**, and this one would have shipped a green skill-path arm that proved nothing.

## Step 3 — the reading

```
content-stories: 4 products | 31 landable pairs | 2 presence-only (ai-labs) | 0 dropped
content-stories: LANDED 16 / 31
  simulation:        13/26     ← all 13 PLAYER pairs land; 0 manager
  skill-path-legacy:  3/4      ← both players + 1 manager
  skill-path-new:     0/1      ← academy (unseeded, known)

  --- not landing (grouped by cause) ---
  x11  attempts table empty ("No data") + player name renders "undefined undefined"
  x2   no "<player>'s Results for <sim>" header      (the two interview sims)
  x1   page.goto timeout (a skill-path manager route)
  x1   route rendered a not-found                    (academy)
```

The harness independently re-derived the denominator as **31**, matching `metrics.json` and the shipped
manifest — three independent derivations now agree.

**All 13 simulation player pairs land.** The M232 session-clone → M233 projection → live render chain is
proven end-to-end: real scores, real cloned LLM feedback, real skill names, re-tenanted into the demo org.

## Step 4 — the documented-rule reversal (B4) + the protocol backfill (B5)

`coverage-protocol.md`'s `skipPaths` bullet justified excluding result pages by asserting they are
*"a runtime-computed AI evaluation … never written by a seed"*. **M231 refuted that premise** (a result page
is a **PERSISTED READ** — plain Ent SELECTs, no recompute) and **M236 has now proven it live** (13/13).

The amendment is deliberately **not** a blanket deletion of the rule:

- The **premise is corrected** in both the doc and `coverage.spec.ts`'s comment.
- The **rule is KEPT**, on its true and much narrower grounds — **crawl scope**: a hero's BFS should not
  dive into arbitrary historical session deep-links. Deleting it would prove nothing extra and would make
  the hero sweep re-walk pages the content-stories sweep already asserts far more precisely.
- The doc now names **where result pages ARE proven** (`tests/content-stories.spec.ts`).

Also backfilled, per B5: a full **"Content stories — the (session × action) LANDS sweep"** section in
`coverage-protocol.md` (the four shapes, the 33/31/18 denominator arithmetic, the two harness invariants,
how to run it) and a **"Content stories — where the (session × action) proof lives"** section in
`playthroughs.md` (why this is *not* a Playthrough, and what the two suites share).

## Metric

**0 / 31 → 16 / 31.** First numerator movement of the milestone.

## Close — 2026-07-20

**Outcome:** The content-stories proof surface exists and is calibrated. 16 of 31 pairs are proven landing
live on `billion`, including **every** simulation player pair. The residual is dominated by a **single
manager-vantage defect** (13 of the 15 non-landing pairs), now precisely characterized.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (3 tiks this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (exact-path over VantageManifest — the scope shrank), D2 (the false-PASS and why shape is route-selected), D3 (skipPaths KEPT, premise corrected — the B4 interpretation), D4 (manager defect characterized, routed)
**Side-deliverables:** the `coverage.spec.ts` comment correction (a doc-truth fix, separate from the harness).
**Routes carried forward:**
- **The manager vantage** — `No data` + `undefined undefined` on 11 pairs, no header on the 2 interview pairs. The mirror row IS present (13/13, iter-03) and the player user row IS correct in the DB (`Clara Romano`), so this is a **read-path** defect, not a seeding gap. → **iter-05**, handler `MANAGER-M236-iter05-scoreboard`.
- **Academy** — `/library/<slug>` renders not-found; academy tables empty. → the academy arm, handler `ACADEMY-M236-iterTBD-catalog-fill`.
- **The skill-path manager timeout** (1 pair) — may resolve with the manager fix; re-measure after iter-05. → handler `MANAGER-M236-iter05-scoreboard`.
**Lessons:**
- **Calibrate before authoring is not ceremony — it changes the descriptor.** Two of the four render shapes (the 205-char interview page, the progress-not-score skill-path page) are ones a competent author would have gotten *wrong* from the spec alone, in opposite directions.
- **Prefer route-derived classification over content-sniffing.** Selecting the assertion shape by URL prefix is stable; selecting it by keywords let 11k chars of unrelated real content trigger a false pass.
- **A harness that reports a number must survive its own failures.** Both harness bugs here (serial abort, worker-restart amnesia) produced *plausible-looking numbers* rather than errors. Any sweep that aggregates should persist per-item results outside the test process.
