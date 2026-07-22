---
iter: 07
milestone: M244
iteration_type: tik
status: planned
active_strategy: TOK-01
created: 2026-07-22
---

# iter-07 — gate (b): ground-truth sweep + voice cells → presence-only (#2)

**Type:** tik under **TOK-01**. Not bootstrap; not triggered-tok (iter-05/06 both made measurable progress).

## Active strategy reference
TOK-01 — discharge gate parts in dependency order. This iter targets gate **(b)** (content-stories sweep)
evidence + the decided **priority #2** (voice cells → presence-only).

## Cluster / target identified
iter-03 left gate (b) at 46/49 with "3 deterministic residuals (voice/interview render)". iter-06 fixed the
interview alignment (gate g). Remaining per the resume plan: the 2 voice cells (`hire-voice-fail` +
`asmt-voice-pass-en`) whose result surfaces cannot land on billion (**Bunny recording keys ABSENT** — user
decision: presence-only, denom 49→47). BUT there is NO existing per-session presence-only field (presence-only
is per-product, ai-labs). So this iter is EVIDENCE-FIRST: run the sweep against billion to see EXACTLY which
pairs fail + their render shape, then disposition.

## Hypothesis
The sweep confirms `hire-voice-fail` + `asmt-voice-pass-en` are the deterministic non-landing voice pairs;
implementing a per-session presence-only exclusion (denom 49→47 + manifest projection drops their landable
pairs, fail-closed) turns gate (b) toward its final green.

## Planned scope (deliverables)
1. Run `run-content-stories.sh 1 --host billion.taildc510.ts.net` (foreground) → ground-truth the failing pairs.
2. Disposition: implement the per-session presence-only exclusion for the 2 Bunny-absent voice cells —
   `content-denominator.json` 49→47 (+ arithmetic comment + recorded reason), manifest projection drops their
   landable pairs (fail-closed, no fabricated CTA), honesty gates (content-route-contract / content-pairs /
   canonical manifest) regenerated + green.
3. `go test` green (stack-seeding) + the affected unit specs green.
4. Tag + push rext.

## Expected lift
Gate (b) denominator corrected to the honest landable set (47); the 2 Bunny-absent voice residuals removed as
presence-only. Sets up gate (b) final green at the gate-(h) cold reset-to-seed. (Gate (b) is COUNTED only when
the live 47/47 lands on the m244-seeded billion — folded into gate h.)

## Phase plan
Phase A: sweep (evidence). Phase B: presence-only mechanism + denom + manifest. Phase C: go test + unit specs.
Phase D: tag+push.

## Escalation conditions
- If the sweep shows DIFFERENT failing pairs than the 2 expected → re-survey the disposition (don't blind-drop).
- Any platform-repo edit demanded → ESCALATE (blocker); route to a demopatch.

## Acceptable close-no-lift outcomes
If the sweep evidence contradicts the presence-only premise (the 2 cells actually land, or different cells
fail), close with the evidence recorded + the disposition re-planned (a complete falsification).
