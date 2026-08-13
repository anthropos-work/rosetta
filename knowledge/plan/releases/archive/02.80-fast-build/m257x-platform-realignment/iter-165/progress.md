**Type:** tik — under [`TOK-08`](../decisions.md), on the axis iter-164 opened.

# iter-165 — the auditor fabricated 11 of its 11 findings

## Phase A — enumerate the accept clauses

Four checked-in waiver files, **20 entries**: `claim_twin` 17, `repair_reach` 1 (+`_README`),
`value_change` 1, `repair_leak` 1. Each is a standing *"this is fine"* — an accept clause with a
`path` and a quoted form, which looked decidable: either the form is still there or it is not.

**First run: 11 of 20 dead.** A milestone that re-runs these fences every iter, with more than half
its waivers silently rotted, would be a genuinely large finding.

## ⚠ Phase B — it was too large, and the instrument was wrong three separate ways

The number was implausible, so the **instrument** was re-checked before the **corpus** was. Every one
of the three defects produced confident output:

| assumption | reality | fabricated |
|---|---|---|
| normalising punctuation is enough | whitespace **runs** survived, so any waiver whose quote spans a wrapped markdown line missed | **10** |
| every waiver names a `path` | `repair_reach_waivers.json` holds **dispositions keyed `path:line`** — a different schema | **1** |
| `form_contains` matches raw prose | it matches the guard's **derived claim form** | **1** |

**Corrected: 0 of 20 provably dead.**

This is the **second** instance in this run. iter-163's first draft reported **346** findings that
were a pairing cross-product; this one reported **11** that were a normalisation bug. Different
mechanism, one shape: **an instrument's preprocessing is part of its predicate**, and a `re.sub`
written in five seconds is a predicate written in five seconds.

## Phase C — the falsification, and what it sharpens

Refusing to re-implement each guard's matching leaves one correct instrument: **ask the guard.** It is
not available.

- `repair_leak_guard` is **diff-scoped**: at HEAD it exits `CANNOT RUN — no candidate shingles`, and
  that is right — its subject is a repair diff and there is none. Its waivers cannot be honoured
  outside a repair, so they cannot be audited outside one either.
- `claim_twin_guard` runs against the live corpus and prints **nothing** about waiver usage. A waiver
  that never fires is indistinguishable from one that fires every run.

So the route sharpens from *"audit the accept side"* into something buildable: **a fence carrying a
waiver must report which waivers it honoured, per run** — the accept-side analogue of the reach
numbers the fire side already publishes. `§8` already says *a fence whose reach shrinks in silence is
the failure this milestone keeps finding*; **a waiver that never fires is a reach hole on the other
axis**, and nothing prints it.

## Gates

No source changed. The four guards were **executed** as part of the investigation
(`repair_leak_guard` → `CANNOT RUN` by design; `claim_twin_guard` → ran, reports no waiver usage);
no fence was modified, so no fence suite was re-run.

**NOT re-run, named in full (`§5` rule 60):** every suite — **this iter modified no code at all.**
The only tree changes are this iter's own records.

## Close — 2026-08-08

**Outcome:** the accept-side audit is **refuted as framed**. 20 waiver entries enumerated; the first
reading said 11 were dead; **all 11 were artifacts of the auditor** — a whitespace-normalisation bug,
a second waiver schema, and a form-vs-prose confusion. Corrected count: **0 provably dead**. The
falsification is the deliverable, and it sharpens the route into something buildable: the guards do
not report which waivers they honour, so the accept side cannot be measured from outside them.
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**this is 1 no-prog tik, not 3; and no
`N` reading was taken, so the metric is UNMEASURED not unmoved (`§9`) — a successor strategy remains
FORBIDDEN by `TOK-08`'s sealed rule**) — (3) re-scope: n (a single close-no-lift with documented
falsification is the protocol working) — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop:
n — (7) budget-exhausted: **y** — Outcome: **exit-7**
**Decisions:** `D-M257x-165-1` … `D-M257x-165-3` (see [`decisions.md`](decisions.md))
**Side-deliverables:** none. **No repair was landed on the withdrawn measurement's residue** — the one
`repair_leak` waiver whose quote is absent from raw prose matched a *derived form*, and deleting it
would make a fence louder for the wrong reason.
**Routes carried forward:**
- `SURVEY-M257x-iter164-acceptance-clauses-are-unaudited-for-over-reach` — **SHARPENED, not closed.**
  Now: `FIX-M257x-iter165-fences-do-not-report-which-waivers-they-honoured`. Buildable, small per
  guard, and it makes the accept side measurable from outside for the first time.
- `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer` — **new hard evidence.** Three
  different waiver schemas across four files is exactly the missing shared layer, and it cost this
  iter its whole first reading.
- Unchanged and still queued: `SURVEY-M257x-iter164-verification-662-claim-not-adjudicated` ·
  `FIX-M257x-iter163-block-ref-attaches-the-wrong-sha` ·
  `SURVEY-M257x-iter163-anchors-with-no-quoted-literal` ·
  `SURVEY-M257x-iter163-generic-literals-are-unadjudicable` ·
  `SURVEY-M257x-iter162-a-literal-has-a-ROLE-the-census-cannot-see` ·
  `SURVEY-M257x-iter162-small-derivations-are-coincidence-prone` ·
  `SURVEY-M257x-iter161-golden-vs-frozen-needs-input-provenance` ·
  `SWEEP-M257x-iter159-grade-the-961-haystack-candidates` ·
  `SURVEY-M257x-iter160-inexact-copies-are-invisible-to-an-equality` ·
  `FIX-M257x-iter160-b2-over-strict-direction-still-unfenced` ·
  `SURVEY-M257x-iter158-noise-classifier-is-narrow-by-choice` ·
  `SURVEY-M257x-iter156-other-reporting-layers` · `SURVEY-M257x-iter152-half-up-services-are-ungradeable` ·
  `FIX-M257x-iter155-add-injected-rows-vs-the-pinned-liveness-only-decision` ·
  `FIX-M257x-iter145-sha-baseline-drift` · `-iter145-migrate-race-needs-a-host-postgres` ·
  `-iter145-green-but-stale-graphql-mentions` · `-iter143-wrong-head-is-unfenced` ·
  `-iter143-scope-derivation-by-grep` · `-iter143-appending-to-the-protocol-doc-rots-the-ledger` ·
  `-iter144-correction-vs-retraction-unfenced` · `SURVEY-M257x-iter144-orphan-arm-is-the-residual` ·
  `FIX-M257x-iter142-path-arm-window` · `-iter142-value-change-articles` · `-iter142-tier-b-underflag` ·
  `FIX-M257x-iter135-adjudicated-live-defects` · `-iter140-receipts-not-checkable-here` ·
  `-iter140-receipt-fence` · `-iter133-two-fives-need-a-fence` · `-iter131-predicate-sets-not-enumerated`
**Lessons:** **check the instrument when the number is too good.** Both of this run's fabricated
readings — 346 and 11 — were caught by the same reflex, and neither by a test: the finding was larger
than the world it described, so the auditor was re-read before the corpus was. That reflex is the only
thing standing between a five-second `re.sub` and a batch of confident false repairs. And the
corollary this iter had to apply to itself: **a withdrawn measurement is not a licence to act on its
residue** — the one waiver still "failing" after the correction was failing against the wrong subject,
and deleting it would have made a fence louder for a reason that does not exist.
