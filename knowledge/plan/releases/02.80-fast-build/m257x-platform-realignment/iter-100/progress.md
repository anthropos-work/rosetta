# iter-100 — the fence was lying about its own reach. Fixed, mutant-proven, 8 defects repaired.

**Type:** tik, under `TOK-05`. Discharges `CHECK-M257x-iter99-anchor-guard-blindspot` and answers
`CHECK-M257x-iter99-briefing-rext-ref`.

**Outcome:** `anchor_construct_guard` resolved **360** anchors and reported GREEN. It now resolves **528**
and the widened reach exposed **8** wrong-construct citations that were invisible — including the two
iter-99 booked, both adjudication panels upheld, and this guard passed. All 8 repaired; guard GREEN;
**26/26 mutants** match their declared verdict, 6 of them new and each a named kill.

**No reading is in this pass. `N` is unchanged at 28 and the gate does not move.**

## 1. What "resolvable" was hiding

The corpus writes anchors **anaphorically**: the file is named once and the anchors that follow carry it.

    hiring.md:80-81   `organization/manager.go:448` (… `candidate` … `member`) and `:485`
    ai-readiness.md:46  … contradicted `:458` of this same file

`_SELF_REF` admitted **one** spelling of that construct (`:N above|below|earlier`) and always resolved it to
the containing document — the right target for the second example, the wrong one for the first. Neither
reached `classify()`. `:485` is a closing brace; `:458` is a blank line. Both sat under a green fence.

**Measured before writing anything:** the corpus carries **311** complete-span bare `` `:N` `` anchors —
comparable to the 360 the guard resolved in total. Bucketed: **8** self-marked, **178** inheriting a file
from a preceding citation, **125 orphans**. The orphan bucket is where `:5050`, `:8082`, `:7700` live.

## 2. The narrowing is the deliverable, not the widening

The first widened draft resolved 511 and reported **23** findings. Triaged, several were the guard's own
documented failure mode returning in a new costume — *"134 findings, essentially all of them ports."* Four
narrowings, each derived from a construct rather than fitted to an answer key, and each measured:

| rule | mechanism | what it removed |
|---|---|---|
| complete backticked span only | `` `:8082/graphql/query` ``, `` `backend:8083` `` are not complete spans | the prefix class outright |
| an intervening **address** breaks the chain | a second candidate referent — `block_ref`'s *"more than one → ambiguous"* on a new axis | `external_services.md:367` `:5050`, `cms.md:113` `:8083` |
| an intervening **filename** breaks the chain, backticked **or bare** | the prose switched files without re-citing one | `cms.md:113` `:169`/`:173`; `backend.md:70` + `skillpath.md:36` (*"the CLAUDE.md line was `:80`"* — CLAUDE.md unbackticked) |
| a **superseded quote** is not graded | the corpus citing an anchor *in order to say it is wrong* | `security_compliance.md:220` (*"the anchor said `:489`"*), `cms.md:217` (*"it was `:11`/`:19`"*) |

The last one matters beyond its count: without it the fence **reddens on documented repairs**, which would
make it hostile to the exact discipline it exists to enforce.

Two reach routes were added on the same evidence:

- **suffix-unique resolution.** `organization/manager.go` is cited repo-relative inside `hiring.md`, whose
  stem names no repo, so every positional route missed it — which is *why* iter-99's finding was
  unreachable. Exactly one tracked file ends in that path. **A longer key is a stricter test than the bare
  basename route already shipped, not a looser one.** 45 sites resolve, 6 ambiguous and refused, 104 match
  nothing.
- **multi-ref grading.** A block naming two refs is the corpus recording a *move*
  (`gotenberg.md:14`: *"the anchor was `:268` at `0dab54d`; the file is now 186 lines"*). The ambiguity sent
  the guard to the default ladder and it graded a historical anchor against origin HEAD — reading a file the
  document never named for that anchor, which is the defect `read_target`'s own docstring was written
  against, arriving through the one door it left open. An anchor is now defective only if it names a
  non-construct at **every** ref its block offers. **6 findings dropped.**

Findings now carry **their own ref**. The run-level `adjudicated at` line names every ref the pass touched
and could not say which one graded a given anchor; this iter spent a full derivation deciding whether
`app/main.go:1450` was a `}` or a constructor call. **It is both — the ref is the whole answer.**

## 3. The 8 repaired, and one of them is a finding about the corpus's own conventions

| site | was | now |
|---|---|---|
| `CLAUDE.md:223`, `messenger.md:43` | `app/main.go:1450` graded at `9d00a313` — the only sha either block named, and *the ref the prose says the anchors moved AWAY from* | origin/main sha `2035f9a` stated, so the block offers both |
| `service_taxonomy.md:130` | `rows :137/:138` — meant **its own** rows, resolved against the doc cited one line earlier | `service_taxonomy.md:137` spelled out |
| `ai-readiness.md:46` | `:458` (blank) | `:459` — the `✅ CORRECTED M219` block |
| `ai-readiness.md:595` | `useAIReadiness.ts:274` (a `}`) | `:326` — `interviewQuestions` |
| `hiring.md:80` | `organization/manager.go:448` (blank) | `:450` — `switch org.IsHiring`, effect at `:453` |
| `hiring.md:81` | `:485` (a `}`) + `siminvitationlink.go:62` | `:537` — the `fmt.Errorf` + `:63` |

`service_taxonomy.md:130` is the one worth reading twice: a bare `rows :137/:138` in a document that had
just cited a *different* file reads as an anchor into that file. **The prose was ambiguous enough to fool
the resolver, so the repair is to name the file, not to teach the guard another guess.**

## 4. How many prior findings the fixed guard re-grades — **2 of 7, and the honest denominator is 2**

Of iter-99's seven upheld wrong-construct citations:

| # | site | now caught? | why |
|---|---|---|---|
| 1 | `ai-readiness.md:46` — self-citation onto a blank line | **YES** | self-marked route |
| 2 | `hiring.md:80-81` — `manager.go:485` is a closing brace | **YES** | continuation route + suffix-unique |
| 3 | `ai-readiness.md:305` — `urls.ts:52` names the wrong constant | no | the line has content; catching it needs the sentence's *claim* |
| 4 | `graphql-wundergraph.md:134` — `:84` is the Ports bullet | no | same, and it is an orphan (no source to inherit) |
| 5 | `hiring.md:38` — twin drifted to `:52` | no | same class as #3 |
| 6 | `messenger.md:53` — unpinned anchors at no named ref | no | ref discipline, not construct |
| 7 | `jobsimulation.md:203-204` — a ref *labelled* origin/main that is not | no | ref discipline |

**2 of 7 is the whole re-grade, and that is the correct number to report.** Five of the seven were never
this instrument's class: three require deciding what a sentence *claims* — the line iter-45 explicitly
declined to cross and re-declines here — and two are ref-discipline findings. The band was set at ≤1 to
detect a blind spot; it detected one, and the blind spot turns out to be **two mechanisms, not seven**.

The fence also found **6 wrong-construct citations no reading has ever named**, which is the more useful
number: `CLAUDE.md:223`, `messenger.md:43`, `service_taxonomy.md:130` (×2), `ai-readiness.md:595`,
`hiring.md:80`.

## 5. The briefing gap — it is in the FROZEN briefing, so it is reported, not fixed

`CHECK-M257x-iter99-briefing-rext-ref` supposed the briefing was *silent* on which rext clone grades an rext
claim. **It is not silent. It is wrong**, and that is a sharper finding.

`instrument/briefing-iter76-AS-RUN.md:37` — sha256 `3858ec53…`, one commit ever (`012edd2`) — reads:

    | rosetta-extensions (the tooling) | `.agentspace/rosetta-extensions` | authoring copy, `main` |

Both seats that booked `external_services.md:208-211` **followed the briefing correctly** and were rejected,
because the adjudicators graded against the pinned per-stack clone `ab81527a`. iter-99's ground-truth sheet
lists **both** clones and says which is which — it just never says which one *settles a claim*.

**This is not seat noise and not a gap; it is an instrument defect that produced two identical false
bookings by construction.** The briefing is frozen and may not be edited, so per the standing rule it is
**reported and read anyway**; the durable fix lands in the protocol doc (§5 rule 45) and the next reading's
ground-truth sheet carries it as a marked addendum.

## 6. Separating the three comparability mechanisms

iter-99 named three and resolved none. This iter separates **one and a half**:

- **Briefing gap — SEPARATED and CHARACTERISED.** Not ambiguity but a wrong instruction, accounting for
  **4 of the 10 rejections** (the ref-discipline class, now 17 occurrences over five readings). Its
  contribution to the precision drop is bounded and known, and it is fixable — though not before the next
  reading, since the fix cannot enter the frozen briefing.
- **Instrument degradation — PARTIALLY SEPARATED.** The wrong-construct class was an instrument blind spot,
  now closed for 2 of its 7 members. The remaining 5 are not this instrument's class at all, which means
  band #9's ~7× miss was **two mechanisms wide, not seven** — a materially smaller instrument defect than
  iter-99's write-up implies.
- **Residual hardening vs adjudicator variance — NOT separated.** One disagreement in 46 cannot be
  distinguished from noise at n=1, and nothing in this iter bears on it.

**The series is still not comparable and this iter does not make it so.** No claim is made that `N` rose,
fell, or converged.

## 7. What would bound the open-ended classes — answered, and the answer is not uniform

Asked directly: what would bound scoping errors, self-inflicted model drift, and intra-document
self-contradiction?

- **Scoping errors — BOUNDABLE, and the mechanism already exists.** A scoping error is a claim true at one
  ref/tree/profile and asserted unqualified. `platform_predicate_guard` and `demo_knob_guard` already fence
  exactly this by deriving a legal set and checking both directions. Cost: one sibling guard per predicate
  family, ~1 iter each. Bounded because the predicate families are enumerable from the platform's own config.
- **Self-inflicted model drift — BOUNDABLE by construction, and cheaply.** It is induced by *our* rext
  development, so it is bounded by the same `--range` machinery `repair_reach_guard` uses: any rext commit
  that changes a construct the corpus cites re-resolves the citations into it. This is iter-59's
  `D-M257x-59-3` (*a pin advance is not vetted until its CITATION delta is measured*) pointed at our own
  repo instead of the platform's. **Cost: one guard, and the whole set is small.**
- **Intra-document self-contradiction — NOT boundable by anything we can build, and this is the sentence
  that matters.** It is quadratic in a document's claim count and the relation is *semantic entailment*, not
  string identity. `frontend-tier.md` held both readings of the demo-academy auth model **nine lines
  apart**, in different vocabulary, each internally coherent. No guard in this family can see that: the
  whole family's line — stated in `anchor_construct_guard`'s own docstring about blocker #17 — is that it
  **does not decide what a sentence claims**, and self-contradiction detection is nothing but that decision,
  performed pairwise. A fence for it would be an LLM judging entailment over every claim pair in every
  document, which is neither derivable, nor cheap, nor stable enough to gate on.

  **So: nothing we can build bounds it, and clause 5's zero is therefore not reachable by fencing alone.**
  It is reachable only by *reading* — which is exactly what clause 5 already requires and why the union-of-
  two blind readings exists. The honest consequence: the fence family can drain the enumerable classes to
  zero and the residual will still be whatever two blind readings can find, with the recall those readings
  have (union ≈ 62 % and falling as the pool narrows). **This is stated plainly rather than hedged, because
  it is the strongest available claim about what the remaining distance to clause 5 actually consists of.**

## 8. Test evidence

- `TheContinuationAnchor` — **9 new behaviour tests**. Two prove RED on a synthetic wrong-construct citation
  of the missed class (closing brace, blank line); **seven prove SILENCE** on the over-match shapes (port
  beside a citation, unbackticked filename, superseded quote, table-cell boundary, non-unique suffix,
  anchor valid at one of its block's refs). The silent half is the larger half on purpose.
- `test_m257x_mechanical_fences_mutation_battery` — **26/26 declared verdicts met** (was 20). Six new
  mutants, each a **named kill**: span-completeness loosened · address-break removed · filename-break
  removed · superseded-test inverted · suffix-uniqueness loosened · multi-ref grading removed. The
  declared-GREEN control still survives and the signature-discrimination assertion still holds.
- `anchor_construct_guard` on the live corpus: **528 resolved / 0 findings**, from 360 resolved. Reach
  **+47 %**.

## Close — 2026-08-06

**Outcome:** the instrument's blind spot is closed and mutant-proven; 8 wrong-construct citations repaired,
6 of them never named by any reading; the briefing defect is localised to the frozen instrument and
reported; the open-ended-class question is answered, including a plain *no* for the class that decides
whether zero is reachable by fencing.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5, unchanged.** Clause 5 is met only by a reading that returns zero, and this
pass took no reading.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n (1 tik this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** `D-M257x-100-1` … `D-M257x-100-4` in [`decisions.md`](./decisions.md).
**Side-deliverables:** per-finding ref provenance in `anchor_construct_guard`'s output — unplanned, and
recorded separately because it was found by needing it, not by looking for it.
**Routes carried forward:**
- `FIX-M257x-iter100-read-union` — the reading itself. **Not taken this session** (budget), and it is now
  the next action: the instrument is repaired, so a reading taken under it is comparable in a way iter-99's
  was not.
- `CHECK-M257x-iter100-briefing-frozen-rext-ref` — the frozen briefing names the wrong rext tree. Cannot be
  edited; the next reading's ground-truth sheet must carry the grading rule as a marked addendum.
- `FIX-M257x-iter100-semantic-anchor-class` — the 3 of 7 that need the sentence's claim, re-declined here.
**Lessons:**
- **A fence's green is scoped by a word, and "resolvable" was doing the work.** The reach number is the
  fence's real claim; 360 of 555 was never stated as a coverage limit anywhere the reader would see it.
- **Widening a fence is the easy half; the narrowing is the deliverable.** 23 findings → 8, and every
  removal was a construct the corpus genuinely uses.
- **When prose fools the resolver, fix the prose.** `service_taxonomy.md`'s bare `rows :137` was ambiguous
  to a careful reader too; teaching the guard another guess would have been the worse repair.
- **A finding that does not name the ref it was graded at costs a derivation to interpret.**
