---
iter: 4
milestone: M236
iteration_type: tik
iter_shape: multi-step (calibrate → author → prove)
status: closed-fixed
date: 2026-07-20
metric_pre: "0/31"
metric_post: "16/31"
---

# iter-04 — tik: Phase H, calibrate the content-stories result render and author the harness

## Step 0 — re-survey (mandatory)

Re-ran the measurement precondition. State changed materially in iter-03 (a live stack now exists), so the
re-survey both confirms the target and **shrinks it**:

- Stack live at `playbill-m235-hardened`; cockpit `/content-manifest.json` HTTP 200.
- All demo endpoints are reachable **from this workstation over the tailnet**: web `:13000` → 307
  (unauth redirect, as expected), cockpit `:17700` → 200, fake-FAPI `:15400` → 404 on `/` (an API root).
  So the harness can be driven locally against the remote stack — no need to install a toolchain on the host.
- **The harness is smaller than the audit's B3 framing implied.** `cockpit-login.ts:76 loginAs()` already
  accepts a **`landingPath`**, so a content seat can be landed **directly** on its exact result URL. The
  content-stories sweep therefore needs **no crawl at all** — it is an exact-path visit per (seat × path),
  not a BFS. The `VantageManifest`/`seedPaths`/BFS machinery — the thing whose singular `identityKey` forced
  "13 seats → 13 manifests" — is simply **not the right tool** and is not needed.
- Seats resolved from the live manifest: **13 player seats** `content-player-23 … content-player-35`,
  manager seat **`dan-manager`**. Products carry `id` (not `key`): `simulation`, `skill-path-legacy`,
  `ai-labs`, `skill-path-new` (Academy, `app_base: academy`).

TOK-01's Phase H is the right target. **Substitution within it:** build an exact-path content-stories
runner rather than 13 `VantageManifest`s. Same strategy, cheaper and more correct shape — recorded per
Phase 1 Step 0.

## Active strategy reference

**TOK-01 "publish-then-prove", Phase H** (harness-after-first-live-render).

## Cluster / target identified

The stack renders; nothing measures it. Every remaining pair is blocked on the same missing capability:
*log in as a content seat, visit an exact result URL, assert the result is real and non-empty.*

## Hypothesis

Calibrating against **one** live seeded render will yield a correct section descriptor; an exact-path
runner built on that descriptor will then measure all 31 pairs, moving the metric off 0 for the first time.

## Expected lift

**≥1 pair proven** (the calibration target). Realistically the simulation arm's 26 follow immediately if
the descriptor generalizes — but the iter is graded on **a working, calibrated harness + a real reading**,
not on hitting 31.

## Phase plan (declared multi-step — see the tripwire carve-out)

1. **Calibrate.** Log in as one content-player seat, land on its exact result path, and capture what the
   page *actually* renders (headings, text volume, key locators). No authoring before this.
2. **Author** the result page-object + the exact-path content-stories runner from that evidence.
3. **Prove** — run it across the manifest and take a real reading of the primary metric.
4. **Reverse the `skipPaths` `/result/` exclusion** and **amend `coverage-protocol.md`** in the same
   change (B4), including *why* the original rule was right when written and what refuted it.

Steps 1–4 are the **planned** shape; an unplanned 3rd line still fires the tripwire.

## The B4 amendment, stated precisely

`coverage-protocol.md:421-431` excludes result deep-links via `skipPaths`
(`tests/coverage.spec.ts:39 RESULT_DEEP_LINK_SKIP`). The rule was **correct when written** (M42e
iter-04/05): its premise is that a result page is a **runtime-computed artifact** — "an AI evaluation the
jobsimulation pipeline computes server-side … never written by a seed" — so a seeded demo could not fill it
without a platform action.

**M231 refuted the premise.** `/sim/<slug>/result/<sessionId>` is a **PERSISTED READ**:
`jobsimulation/internal/graph/queries.resolvers.go:70` does plain Ent SELECTs of
`validation_attempt_results`, with no engine/LLM recompute on render. M232 then *writes* that fan-out.
iter-03 verified it live: **13/13** sessions carry attempt-result rows.

So the exclusion must be reversed **for seeded content-story sessions** — and the doc amended to record
that the rule's justification, not merely its scope, changed. Blanket-crawl exclusion of *arbitrary*
historical result links stays sound; these paths are enumerated from a manifest, not discovered by a crawl.

## Escalation conditions

- **The result page renders empty despite present rows** → a real content defect; triage here (in-scope).
- **The seat cannot log in** → cockpit/roster defect; triage here (in-scope, it is the harness's premise).
- A fix that would need a **platform-repo edit** → hard stop.
- Manager-vantage paths behaving differently from player paths → note and route to the manager arm rather
  than expanding this iter.

## Acceptable close-no-lift outcomes

If calibration shows the render is blocked for a characterized reason (e.g. an entitlement gate), that
falsification is a complete iter outcome — it converts an unknown into a named defect with a handler.
