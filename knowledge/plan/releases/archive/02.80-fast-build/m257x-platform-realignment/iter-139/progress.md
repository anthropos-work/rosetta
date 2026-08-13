**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*). iter-138 produced a
census; **this iter measured the instrument that produced it, before anything was repaired from it.**

# iter-139 — the census was wrong, and the audit that found it was pre-registered

## What this iter set out to do, and why

iter-138 published **`127 rotted / 222 decidable = 57.2 %`** and routed a fence to be built against that
baseline. **The probe behind it had never been audited.** Three specific reasons to distrust it were
written into this iter's `overview.md` *before* any case was opened:

1. deltas of `+1557`, `+820`, `+295` are implausible as same-file rot;
2. unique-**text** matching is not unique-**construct** matching;
3. `git blame` on the citing line dates the line's last edit, not the citation's authorship.

**Strata and selection rule sealed in advance:** 4 cases from `|Δ| ≥ 100`, 4 from `10 ≤ |Δ| < 100`, 4 from
`|Δ| < 10`, each taken **in census order from the top of its stratum** — reproducible, not cherry-picked.

## Result — 12 of 12 are FALSE POSITIVES

**Precision 0/12 = 0.0 %. Wilson 95 % [0.0, 24.3].**

| # | case | what the citing line actually is |
|---|---|---|
| 1 | `platform-alignment.md:2595` → `:788` | *"two of iter-45's five defects are relationships between line numbers — `:788` citing `:447`"* — a **historical example quoted** |
| 2 | `platform-alignment.md:1108` → `:1305` | *"(`:1305` was the first, iter-84)"* — a claim about a **past** position |
| 3 | `build-budget.md:394` → `:319` | `` (`:279`, `:319`) `` — **continuation pins into another file** |
| 4 | `security_compliance.md:386` → `:129` | `` (`isThrottlingError`, `:129` / `:166` / `:325`) `` — continuation pins |
| 5 | `platform-alignment.md:51` → `:106` | continuation pin into `migrate-demo.sh` |
| 6 | `ai_architecture.md:325` → `:75` | *"(`v3/validator/skills.go:53-64`) then … `:75` computes"* — **explicitly** a continuation pin |
| 7 | `shared_libraries.md:198` → `:259` | `` `AIManager.getClient` (`:259` and `:332`) `` |
| 8 | `platform-migration-status.md:106` → `:183` | `` `docker-compose.yml:170-171` … (`:183`) `` |
| 9 | `messenger.md:53` → `:29` | *"there is **no** `:29` declaration at all"* — a **negated** pin |
| 10–11 | `messenger.md:60` → `:63`, `:62` | `` `app/main.go:15`, `:62`, `:63` `` |
| 12 | `platform-alignment.md:54` → `:108` | continuation pin into `migrate-demo.sh` |

**In this corpus a bare `` `:NN` `` is overwhelmingly a cross-file CONTINUATION pin** — a second, third or
fourth line reference into a file named earlier in the same sentence — or a quoted / historical / negated
pin. **It is very rarely a same-file self-citation**, which is the only thing the probe could have been
measuring.

## The retraction, and where it was placed

**`127 rotted / 57.2 %` is WITHDRAWN** — not re-qualified, not narrowed, not restated with a caveat
(`D-M257x-122-3`'s class). Corrected **in place at all three sites that published it**, per rule 54:

| site | done |
|---|---|
| `iter-138/progress.md` | a RETRACTED banner **above** the numbers, which are retained for the record |
| the milestone ledger's iter-138 entry | corrected inline |
| `§5` **rule 63** in the protocol doc | the figures replaced by the audit, and **rules (a) and (b) rewritten** |

## The finding worth more than the number

iter-138 **disclosed** an `out-of-range-then` bucket of 241 and **named its cause exactly right** —
*"largely cross-file continuation pins, which the probe reads as same-file."* That honest disclosure is
what made the remaining 222 look clean.

**It was not.** A continuation pin lands in `out-of-range` only when the cited number exceeds the
**citing file's** length. In a 3,100-line protocol doc almost none does — so **the same failure mode
passed straight into the "decidable" set, undisclosed, and dominated it.**

> **`D-M257x-139-2`: a disclosed limitation is quarantined only if you show the boundary holds.**
> Naming a floor is not bounding it — and the disclosure made the number **more** persuasive, not less.
> **Sample the clean bucket for the disease you just disclosed.**

And the corollary for iter-138's own rule: `D-M257x-138-1` (*an exclusion is only as narrow as the
predicate that justified it*) **survives as a rule and is withdrawn as an application.** Both predicates
— content and rot — are blocked by the **same** unresolved thing: the pin's **head**. iter-138 assumed
the untested predicate was free of the blocker; it was not.

## What stands from iter-138 (so the retraction is not over-read)

- **All 9 citation repairs stand** — each came from `adj-E`/`adj-D`'s hand-verified list and each was
  **re-derived by opening it**. None came from the probe.
- `D-M257x-138-3` (name the construct, never re-pin) — **strengthened**: a head a purpose-built probe
  cannot resolve is a head a reader cannot resolve.
- `D-M257x-138-5` (choose suites by what you changed) — unaffected.
- **`FIX-M257x-iter138-anchor-rot-fence` is RE-SPECIFIED, not cancelled** — first deliverable is **head
  resolution**; it has **no baseline** until that exists.
- **`FIX-M257x-iter138-127-rotted-pins` is WITHDRAWN** — there is no such work list.

## Test gates

| gate | result |
|---|---|
| **Guard family** (`--repo-root` + `--platform stack-demo/platform @ 0c91421`) | **18 GREEN · 0 RED · 4 not-run** |
| **Scoped fence suites** — chosen by what this iter CHANGED (**prose + the protocol doc**, no anchors rewritten) | **102 passed / 0 failed** (the citation/anchor set kept from iter-138 deliberately, since the protocol-doc edit moves lines under existing pins) |
| **Whole suite** | **NOT re-run — §5 rule 60 requires saying so.** Zero `rosetta-extensions` files changed; iter-132's clean run stands on the same rext tree (`223e4a6`). Stated as a gap, not characterised as covered |
| **Suite wall-time** | not quoted — `FIX-M257x-iter132-suite-walltime-is-not-a-measurement` stands |

## Close — 2026-08-08

**Outcome:** iter-138's `127 rotted / 57.2 %` is **retracted** on a pre-registered stratified audit that
returned **0/12, Wilson95 [0.0, 24.3]**. The instrument was measuring a form that barely exists here: in
this corpus a bare `` `:NN` `` is overwhelmingly a **cross-file continuation pin**, not a same-file
self-citation. **Nothing downstream had consumed the wrong number** — no repair was driven by it and the
fence was not yet built — so the cost of the error was one iter, and the cost of skipping the audit would
have been a fence with a fabricated baseline and 127 unnecessary edits.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged; no reading taken.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**iters 134–139 took no reading, so the metric is UNMEASURED not unmoved — §9's iter-type refinement; `TOK-08`'s sealed refutation branch bars an agent-authored successor**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (**3 tiks this session**) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-139-1` (the retraction, with its interval) · `D-M257x-139-2` (a disclosed floor is
quarantined only if you show the boundary holds) · `D-M257x-139-3` (the fence's exclusion was better
founded than iter-138 credited — same blocker, both predicates) · `D-M257x-139-4` (what stands) ·
`D-M257x-139-5` (the loop caught it in one iter, before anything consumed it).
**Side-deliverables:** none.
**Routes carried forward:**
- **`FIX-M257x-iter138-anchor-rot-fence` — RE-SPECIFIED.** First deliverable: **head resolution for bare
  `:NN` pins** (which file does a continuation pin continue?). Only then is a rot baseline meaningful.
  **`anchor_construct_guard` already resolves heads for qualified citations** (868/1469, with a documented
  `bare-continuation` strategy at 235 resolutions) — **that resolver is the thing to reuse**, and reusing
  it is `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer` arriving with a concrete consumer.
- `FIX-M257x-iter135-adjudicated-live-defects` — remainder: `clerk-integration.md:126` ·
  `backend.md:13`'s *UNEVEN* cross-ref (**iter-137 made *UNEVEN* itself questionable — M810's ECS half has
  landed for both cms and jobsimulation; only the schema drops pend**) · `sentinel.md:5` ·
  `ai-readiness.md:18-20` · `org-repos.md:227`,`:370`,`:43` · `ai_architecture.md:111`,`:224` ·
  `next-web-app.md:17`,`:186` · `external_services.md:368`.
- `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer` (**now has a named consumer**) ·
  `FIX-M257x-iter133-two-fives-need-a-fence` · `FIX-M257x-iter132-suite-walltime-is-not-a-measurement` ·
  `FIX-M257x-iter131-predicate-sets-not-enumerated`.
- **WITHDRAWN this iter:** `FIX-M257x-iter138-127-rotted-pins`.
**Lessons:**
1. **An instrument is not a measurement until the instrument is measured.** Ninth time on this milestone.
   The audit cost one iter; skipping it would have cost a fence with a fabricated baseline.
2. **A disclosed floor is not a bounded one.** iter-138 named its failure mode and then reported a number
   over the population it had not excluded it from — and the disclosure made the number *more* credible.
   **Sample the clean bucket for the disease you just disclosed.**
3. **Pre-register the strata.** 0/12 is only believable because the strata and the take-in-census-order
   rule were written down before the first case was opened. An audit chosen after the fact proves nothing.
4. **Retract in place, at every site that published the number** — the iter record, the ledger, and the
   protocol rule. A retraction that only lives in the newest iter is invisible to everyone downstream.
