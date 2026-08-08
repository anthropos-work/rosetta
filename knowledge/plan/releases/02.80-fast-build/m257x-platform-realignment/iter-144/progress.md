**Type:** tik — under `TOK-08`. Closes `FIX-M257x-h30-crossline-repair`. Three planned lines (audit →
machine-control the audit → repair what survives), declared in this iter's `overview.md`.

# iter-144 — the wrapped retraction sites repaired, and the sub-class the fence cannot see

## 1. The re-survey corrected the route's own number

Harden pass 30 routed *"the **8** true sites across 6 files."* Measured at this iter's open, the
wrapped (SURVEY) arm holds **10 findings across 6 files**. The route's figure is not re-quoted.

**The pass routed the count without grading it** — reasonably, since it had just decided the arm must
be SURVEY rather than gating, and grading six documents of prose is an iter's work. But a routed count
is an **estimate of work**, and this iter's first job was to stop it becoming a measurement of defects.

## 2. The audit, and its machine control (`D-M257x-143-1`, one iter old)

All 10 were read in full context. The reading said **7 true / 3 false**. Then — per the rule iter-143
landed exactly one iter earlier — **every pin the reading called LIVE was resolved against the source
it claims**, by machine rather than by a second reading:

| site | the reading's call | what the machine says |
|---|---|---|
| `ai_architecture.md:303` `:98-99` | live, not retracted | `app` `ad9f3c49` → `default: aiModel = anthropic.Anthropic37SonnetAWS20250219` ✅ |
| `ai_architecture.md:303` `:110-111` | live, not retracted | same file → `default: aiModel = anthropic.Anthropic35Sonnet20241022` ✅ |
| `hiring.md:304` `:176-187` | live, not retracted | that range **is** the `job_position` bullet at HEAD ✅ |

**Three for three.** Unlike iter-143 — where the same control overturned **nine** of the reader's
calls — the control confirmed this reading. That is the point of running it: a control is not a
formality that only earns its place when it fires.

## 3. The finding — a retraction clause holds the CORRECTED pin as often as the retracted one

The three falses are one shape, and it is the shape of a **correction**, not a retraction:

* `ai_architecture.md:303` — *"Both anchors named the file nowhere in this bullet until iter-115 —
  they read as bare `:98-99` / `:110-111`"*. What is retracted is the **absence of a filename**. The
  numbers were never wrong and are not now.
* `hiring.md:304` — *"It cited an earlier range — the `job_position` bullet, now `:176-187`"*. The
  retracted value is *"an earlier range"* and is **deliberately not named**. `:176-187` is the
  **correction**.

> **The token a retraction must not reproduce is the OLD value. The token a correction must publish
> is the NEW one.** Same shape, same sentence, same markers — so a form-matching fence sees one class
> where there are two.

**Wrapped-arm precision: 7/10 = 70 %.** Pass 30 was right to keep the arm non-gating, and the remedy
is **not** to tighten it: a fence cannot read which half of a correction it holds. The remedy is to
grade before repairing. Landed as `§5` rule 67.

## 4. The repair — 7 sites, 4 files, all line-count FLAT

Rule 63(c′)'s remedy, with `D-M257x-142-4`'s constraint:

| file | what changed |
|---|---|
| `shared_libraries.md:38` | *"`ai` at `:14` and `messenger` at `:17`"* → *"`ai` and `messenger` were still direct requires"* |
| `ai-readiness.md:113` | *"This cited `:133-134` until iter-115"* → *"cited the two lines above it"* |
| `ai-readiness.md:487` | *"`:662-664` is the body of `audienceScope`"* → *"they landed in the body of `audienceScope`"* |
| `graphql-wundergraph.md:90` | two retracted pins (`:174-176`, `:193`) → *"once before iter-98 and again before iter-138, the second landing on the warning's opening line"* |
| `roadrunner.md:85` | *"the entry sat at `repos.yml:29-31` at `2adcf71`"* → *"`repos.yml` still carried a roadrunner entry at `2adcf71`"* |
| `roadrunner.md:228` | *"from `docker-compose.yml:296-302` at `2adcf71`"* → *"from its `environment:` block at `2adcf71`"* |

Every sentence keeps the number it was making a point with (the **9** entries, the **seven** modules,
the **twice**) and loses only the rot-prone token — *fence the token, not the digit*
(`D-M257x-142-3`). **Line-count flat on all four files**, verified with `git diff --numstat`: 1/1,
2/2, 3/3, 4/4.

**Two of the six are the ref-qualified class** (`D-M257x-142-2`): *"at `2adcf71`"* makes both
roadrunner sentences **true statements about an immutable ref** — and the fence still cannot read the
qualification, so the token is still the hazard. Repaired for the same reason iter-142 repaired its
own: the hazard is the token, not the truth of the claim wrapped around it.

## Test gates

| gate | result |
|---|---|
| `retracted_pin_guard --repo-root .` | gating arm **GREEN**; wrapped arm **10 → 3**, and all 3 survivors are the machine-confirmed live pins of § 2 |
| `test_retracted_pin_guard.py` | **51 passed / 0 failed** |
| Guard family, `--repo-root` + `--platform stack-demo/platform @ 0c91421` | **19 GREEN · 0 RED · 0 could-not-check · 4 not-run** — identical to iter-143's close, so the delta is attributable |
| Scope derivation (**rule 66, one iter old**) | this iter changed **corpus markdown only — zero code** — so the change-derived scope is the tree-scoped guard family plus the guard whose population moved. No call site changed, and that is a derivation rather than a recollection |
| Whole suite | **NOT re-run, and saying so (rule 60)** — iter-143's run (1 failed / 1,294 passed) stands: this iter touched **no** `rosetta-extensions` file, which is the exact condition under which `FIX-M257x-iter142-whole-suite-owed` says a prior run still covers the tree |

## Close — 2026-08-08

**Outcome:** the wrapped retraction-idiom population is **graded, repaired and explained** — **10
findings audited, 7 repaired across 4 files (all line-count flat), 3 survivors confirmed LIVE by
machine.** The route's *"8 true sites"* is corrected to **10 findings, 7 true**. The finding is the
sub-class: **a retraction clause holds the CORRECTED pin as often as the retracted one**, so a
form-matching fence sees one class where there are two — wrapped-arm precision **70 %**, and the arm
is right to stay SURVEY. `§5` gains **rule 67**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged. **No `N` reading was taken, so no `N` movement is claimed**
(`§9`'s guard-rail 1, in its required words).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**iters 135–144 took no reading, so the metric is UNMEASURED not unmoved — `§9`'s iter-type refinement; and `TOK-08`'s sealed refutation branch bars an agent-authored successor in any case**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (**2 tiks this session**) — (6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: **exit-7**
**Decisions:** `D-M257x-144-1` (a retraction clause holds the corrected pin as often as the retracted
one; 70 % is the right precision for a SURVEY arm and tightening is the wrong remedy) ·
`D-M257x-144-2` (grade a survey arm's findings before treating its count as a backlog — a routed
count is an estimate of work, not a measurement of defects) · `D-M257x-144-3` (the machine control
CONFIRMED this reading 3/3 where it overturned iter-143's 9/92 — a control earns its place by being
run, not by firing).
**Side-deliverables:** none.
**Routes carried forward:**
- **`FIX-M257x-iter144-correction-vs-retraction-unfenced` (NEW)** — the two-halves distinction is now
  named and measured but **nothing detects it**. A candidate discriminator exists and is deliberately
  NOT shipped un-audited: a correction's live pin is typically introduced by *"now"* / *"is"* while a
  retraction's is introduced by *"cited"* / *"read as"* / *"was"*. That is a predicate over a
  3-token window, on a **3-site denominator** — exactly the tuned-constant shape iter-143 declined
  twice. Derive it against a real population or leave the arm at 70 % and grade by hand.
- **`FIX-M257x-iter143-appending-to-the-protocol-doc-rots-the-ledger`** — ⚠️ **this iter made it worse
  and says so**: rule 67 added **~40 more lines** to `platform-alignment.md`, so the nine
  `knowledge/plan/**` pins below the insertion point are now off by ~112 rather than 72. Still out of
  `anchor_offset_guard`'s scope by design, still not repaired (rewriting frozen evidence would be
  worse), and now **twice-measured** rather than once.
- `SURVEY-M257x-iter144-orphan-arm-is-the-residual` (renumbered from iter-143's route) ·
  `FIX-M257x-iter143-wrong-head-is-unfenced` · `FIX-M257x-iter143-scope-derivation-by-grep` ·
  `FIX-M257x-h30-nonstackcore-suite` (**untouched — still the open scope call for the milestone**:
  "whole suite" in this ledger has always meant `stack-core` alone, one section of five) ·
  `FIX-M257x-iter142-value-change-articles` · `-iter142-path-arm-window` · `-iter142-tier-b-underflag` ·
  `FIX-M257x-iter135-adjudicated-live-defects` · `-iter140-receipts-not-checkable-here` ·
  `-iter140-receipt-fence` · `-iter138-anchor-rot-fence` (re-specified at iter-143) ·
  `-iter134-fence-family-has-no-shared-predicate-layer` · `-iter133-two-fives-need-a-fence` ·
  `-iter132-suite-walltime-is-not-a-measurement` · `-iter131-predicate-sets-not-enumerated`.
- **CLOSED by this iter:** `FIX-M257x-h30-crossline-repair`.
**Lessons:**
0. **A correction and a retraction are the same sentence shape carrying opposite obligations.** The
   old value must not be reproduced; the new one must be published. No form-matching fence can
   separate them, which is an argument for a survey arm — not against having one.
1. **Grade a survey arm before treating its count as a backlog.** *"8 true sites"* was an estimate of
   work that would have become a defect count the moment anyone quoted it. It was 10 findings, 7 true.
2. **Run the control even when you expect it to confirm.** It overturned 9 of 92 at iter-143 and
   confirmed 3 of 3 here. A control that is only run when it is expected to fire is not a control.
