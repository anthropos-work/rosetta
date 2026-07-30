---
iteration_type: tik
status: in-progress
opened: 2026-07-30
---

# iter-31 — the weak assertion inside the new work, and the verdict that is not a TODO

**Type:** tik · **Active strategy:** `TOK-01` move 4 — *"close the honesty items last, deliberately, not as
leftovers"*.

## Step 0 — re-survey (mandatory)

Ran before targeting, per Phase 1 Step 0:

| fact | reading |
|---|---|
| `ptvalidate --manifest-dir ./manifest` | **VALID** — 10 products, 31 use cases, **29 live**, **2 TODO** |
| the 2 TODOs | `onboarding.enterprise-workforce-standard.UC1` · `onboarding.individual.UC1` |
| `demo-2` | **16 containers Up, 0 exited** (checked `docker ps -a`, per iter-15 D76) |
| suite baseline | `197 passed` × 3 cold, rc 0 each, 0 flake (iter-29) |

Both TOK-01-directed targets are still live and still meaningful; no substitution needed.

## Cluster / target identified

Two planned lines, declared here so the scope-creep tripwire counts against a **2-step planned shape** and
not a single-target tik:

**Line 1 — `ONBOARD-M256-prepared-persistence` (routed by iter-29).** iter-29's own standing mutant S1c
(iter-27's Q1, generalised) **PASSED** against `pt-onboarding-org-prepared`: the persistence half asserted
*"the prepared summary is gone"* on a fresh `/onboarding`, and `toHaveCount(0)` immediately after a
navigation is satisfied by a page that has **not hydrated yet**. The assertion was removed rather than
weakened, and the repair was routed with the one measurement it needs: **a POSITIVE locator on the screen a
reload actually lands on** — the ROLE step, because the component opens on `lastStep || Import` and
`lastStep` is now `role`. Nothing has driven that screen.

This is the milestone's signature defect **sitting inside its own newest work**, so it must not survive
closure.

**Line 2 — `standard.UC1`'s verdict, written properly (`D104`).** The coordinator recorded `D104`: clause 3's
onboarding half is MET at **4 landed + 1 written verdict**, because the self-import journey's only advancing
path scrapes a live public LinkedIn profile on a site that blocks automation (a real person's profile would
become a permanent fixture, and its RED would read as a product regression), while the deterministic CV route
is blocked by a **measured product defect** (upload POSTs 200 for a valid PDF *and* a docx while the forward
control never enables, 100 s). iter-18's refusal stands; this is the twelfth iter it has held.

Today the manifest carries that reasoning **in prose in a story `note`** and the use case's machine-readable
field is a bare `playthrough: TODO` — which `ptreport` renders with the generic detail *"declared use case, no
Playthrough yet (build-reference gap)"*. **That sentence is false for this UC.** It is not a gap awaiting
effort; it is a measured decision. The four-state map therefore states the wrong position about the one use
case the gate is being closed around.

## Hypotheses

- **H1 (line 1).** A fresh `/onboarding` after the role step is confirmed lands on the **ROLE** screen, which
  carries at least one positive, hero-specific locator (her confirmed role) that a *day-0* seat and a
  *not-yet-confirmed* seat both read 0 for. If so the read-back becomes liveness-THEN-absence and S1c goes RED.
- **H1-alt.** If the reload does *not* land on a distinguishable Role screen, the honest outcome is a
  **measured non-fact** recorded at the locator (iter-27's shape), not a re-worded absence.
- **H2 (line 2).** A `verdict` block on a TODO use case — closed-enum disposition + non-vacuous rationale,
  fenced **bidirectionally** (a TODO without one fails; a landed UC carrying one fails) — makes "zero silent
  gaps" a machine property instead of a prose claim, and is the mechanical form of iter-30's **D117** (*a
  routed blocker must carry the measurement that produced it*).

## Expected lift

- Clause 2's control **quality** (not its count): one assertion that currently cannot fail becomes one that
  can, proven by watching S1c go RED.
- Clause 3's onboarding half stated where the tooling reads it, not only where a human reads it.
- No gate **count** moves on either line. That is expected and is not under-delivery: iter-31's deliverable is
  the honesty of the two artifacts the gate is about to be graded on.

## Phase plan

- **Phase A — measure first.** Drive `pt-onboard-prepared` through the confirm, then reload `/onboarding` and
  dump the screen (URL, headings, buttons, her role text) plus the same dump for a *day-0* contrast seat. No
  assertion is written before this returns.
- **Phase B — implement.** Line 1: the page-object locator + the read-back. Line 2: the schema field, the
  bidirectional validator check, the `ptreport` detail wiring, and the two verdict blocks.
- **Phase C — mutants.** S1c re-run (must now be RED) + S1/S2 re-confirmed RED; and for line 2, verdict-removal
  / bad-enum / vacuous-rationale / stale-verdict-on-a-landed-UC each watched RED.
- **Phase D — the gate.** `run-playthroughs.sh 2 --reset` × 3 cold, rc captured per run into a variable,
  `ptvalidate` + `ptreport` + `--policy-check`, gofmt, the four Python suites, and the DRIFTED cockpit fixture
  restored **sha-verified**.

## Escalation conditions

- If Phase A shows the reload lands on a screen with **no** positive discriminating locator, line 1 closes as
  a recorded non-fact (H1-alt) and the iter does **not** invent one — no absence-only assertion is restored
  under any circumstance (iter-29 `D115`).
- If the verdict fence would require touching a landed UC's semantics, drop the bidirectional half rather than
  widen the change.
- `standard.UC1` is **not** to be landed and no route to landing it is to be sought (`D104`; iter-18's refusal).

## Acceptable close-no-lift outcomes

A measured refutation of H1 (the reload is not distinguishable) with the non-fact recorded at the locator and
the routed handler updated to say so, plus line 2 landed, is a complete iter.
