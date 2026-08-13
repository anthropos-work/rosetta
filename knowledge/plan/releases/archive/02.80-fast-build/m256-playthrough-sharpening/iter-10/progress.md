# M256 · iter-10 — progress

**Type:** tik · **Active strategy:** `TOK-01` move 4. Handler: `D-v28-5-cockpit-logout`.

## Phase A — the defect is UNREPRODUCIBLE on this stack, and the reason is a DIFFERENT, worse defect

Drove the presenter's actual clicks against the real cockpit at `:27700`: log in as the first hero, return to the
cockpit, click a second hero once, then again — reading which hero the app reports each time.

**It never switched, on either click — and it never switched to the *first* hero either.** The app reported
`Morgan` throughout, which is neither of the heroes clicked. That is not the D-v28-5 double-click symptom; it is
something else, and measuring the click count on it would have been meaningless.

The cause, confirmed by comparing the two artifacts the demo actually resolves seats through:

| artifact | refreshed by | contents on this stack |
|---|---|---|
| `fake-fapi-roster.json` | **`run-playthroughs.sh --reset`** (the M211 iter-16 roster refresh) | **30** keys: `pt-employee`, `pt-manager`, `pt-free`, `pt-ai-completed`, … |
| `cockpit-manifest.json` | **nothing — baked at BRING-UP** | stories keys: `maya-thriving`, `tom-struggling`, `dan-manager`, … |

**They have drifted completely apart, and nothing tells anybody.** The reset path re-exports the roster (so hero
*login* works, which is why 23 Playthroughs are green) but **never re-exports the cockpit manifest**. So the
cockpit renders **35 `[Log in as]` buttons naming heroes that no longer exist in the world**, and each carries a
`__clerk_identity` key the roster cannot resolve.

**And the failure is SILENT by design.** `clerk-frontend/server.go:347-349` selects the seat best-effort and its
own comment says why: *"An unknown key is ignored (the active identity is unchanged) — a malformed deep-link must
not strand the demo signed-out; it just keeps the current seat."* That is a defensible choice against a
*malformed* link. Against a **systematically stale manifest** it becomes: **every** button silently logs you in
as whoever was last active. A presenter clicks "Maya Chen" and gets Morgan Reyes, with no error in the UI, no
warning in the log, and a perfectly successful-looking login.

**This is very likely the substrate of D-v28-5's reported symptom.** A presenter who clicks and does not get the
hero they asked for clicks again — which is exactly what *"logging out back to the cockpit requires two-or-more
clicks"* describes from the outside. **But that is a hypothesis, not a measurement**, and it is recorded as one:
the double-click cannot be isolated until the cockpit can select a hero that exists.

## Phase B — what landed

**Nothing in code, deliberately.** The fix has a clear shape — the reset path should re-export the cockpit
manifest (`stackseed --cockpit-export`) alongside the roster it already re-exports, so the two artifacts move
together — but it is a change to the reset lifecycle that must be verified end-to-end on a live bring-up, and
this iter did not have the budget to verify it. **Shipping an unverified lifecycle change to close a handler is
exactly the failure iter-07 D31 caught**: a fix that looks right, closes a ticket, and changes nothing (or breaks
the roster refresh that currently keeps 23 Playthroughs green).

Recorded as two routed handlers with the measurement attached, so the next attempt starts from evidence.

## Phase C — the suite

Untouched: no runtime code changed this iter. iter-09's figures stand — `146 passed` ×3, 0 flake, `ptreport`
23/30 passing, 0 failing; clause 1 **0.5652×**. **No re-measure claimed.**

## Close — 2026-07-28

**Outcome:** D-v28-5 could not be measured, and finding out why produced a **worse and previously unrecorded
defect**: on any Playthrough-reset demo the cockpit's every `[Log in as]` button names a hero that no longer
exists, and the seat selection fails **silently** — the presenter gets whoever was last active, with no error
anywhere. The reset path refreshes the fake-FAPI roster but not the cockpit manifest, and the handshake's
deliberate best-effort tolerance for a malformed key turns a systematically stale manifest into a
successful-looking wrong login.
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET — clause 1 **met** (0.5652×, unchanged); clause 2 mutating **6/5 MET**, negative controls
**6 of 23**, `blocked` **0**; clause 3 verdict half **COMPLETE**, landed half short (org-admin 2/4, onboarding
1 of 5); **D-v28-5 still unfixed** — now *diagnosed as blocked behind a prerequisite defect*, not merely unstarted.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (1 no-lift tik; iters 08 and 09 both progressed) — (3) re-scope: n (iter-09's register settles it: 0 of 28 curated UCs `unimplementable`) — (4) user-blocker: n (no red; nothing uncommitted; the batch is green) — (5) **cap-reached: y — this is the 5th tik of the invocation** — (6) protocol-stop: n — Outcome: exit-5
**Decisions:** D41 (the cockpit manifest and the fake-FAPI roster drift apart on every Playthrough reset, and the
handshake's best-effort unknown-key tolerance makes the resulting wrong login SILENT), D42 (D-v28-5 is not
measurable until D41 is fixed — the double-click hypothesis is recorded as a hypothesis, not a finding), D43 (no
unverified lifecycle fix shipped to close a handler — iter-07 D31's lesson applied to this iter's own temptation).
**Side-deliverables:** none.
**Routes carried forward:**
- `FIX-M256-cockpit-manifest-drift` → **next iter, and it now BLOCKS D-v28-5.** `run-playthroughs.sh --reset`
  re-exports `fake-fapi-roster.json` (M211 iter-16) but **not** `cockpit-manifest.json`, so the cockpit lists
  heroes that no longer exist and every seat selection silently falls back to the last-active seat
  (`clerk-frontend/server.go:347-349`). Shape of the fix: re-export the cockpit manifest alongside the roster in
  the reset path, so the two artifacts move together. **Verify on a live bring-up** — do not ship it unverified,
  and do not regress the roster refresh that keeps 23 Playthroughs green. Consider also making the unknown-key
  fallback *loud* on a demo (a visible error beats a successful-looking wrong login), which is a separate call.
- `D-v28-5-cockpit-logout` → **blocked on the above.** The double-click symptom is plausibly a *consequence* of
  D41 (a presenter who does not get the hero they clicked clicks again) — plausible, unmeasured, and it must be
  re-measured on a cockpit whose manifest matches its roster before any fix is designed.
- `NEGCTL-M256-cross-vantage` → clause 2's largest remaining gap (negative controls **6 of 23**).
- `PT-M256-resume-fixture-pair`, `ONBOARD-M256-import-path`, `BLOCKED-M256-refusal-surface`,
  `FIX-M256-studio-false-green` (re-aimed), `DOC-M256-llm-lane-premise`, `PT-M256-orgadmin-role-create`,
  `PT-M256-orgadmin-member-tag`, `FENCE-M256-bounded-interaction`, `PT-M257-self-evaluation`,
  `PT-M257-talk-to-data` — all stand.
**Lessons:**
1. **When a defect will not reproduce, the reason is the finding.** The probe was built to count clicks and
   instead exposed that the cockpit cannot select any current hero at all. Counting clicks on that would have
   produced a number, and the number would have meant nothing.
2. **Two artifacts that must agree, refreshed by different code paths, WILL drift — and something must fail loud
   when they do.** The reset path refreshes the roster and not the manifest; nothing compares them, so 23
   Playthroughs stayed green while the human-facing cockpit was entirely stale. This is the same shape as
   iter-07's phantom-id defect (two gate tools, two views of the tree), and it is the second time in this
   milestone that a gap between two views of the same state hid in plain sight.
3. **A deliberate fail-soft becomes a fail-silent when its premise changes.** Ignoring an unknown
   `__clerk_identity` is right for a malformed deep-link and wrong for a systematically stale manifest. A
   tolerance is only as good as the assumption it was chosen under — re-check the assumption, not just the code.
