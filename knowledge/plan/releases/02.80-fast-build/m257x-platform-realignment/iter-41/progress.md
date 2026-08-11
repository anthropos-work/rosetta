**Type:** tik

# iter-41 — `MEASURE-M257x-iter41-clause5-sixth-pass`

## What this pass did differently — one variable, not two

Every prior clause-5 pass differed from its predecessor in **both** the corpus (a repair landed between
them) and the instrument (more auditors, better briefing). iter-39 proved the consequence: `25 → 13 → 11 →
17 → 37` **measured the instruments**. Nothing in that series licensed a claim about the corpus.

This pass held the instrument fixed on every knob — **7 auditors** (6 full-read A–F + 1 adversarial
diff-reader G), the same briefing, the same partition **method**, all 40 files read top-to-bottom with a
`wc -l` positive control per file. And iter-40's repair, by design, touched `corpus/ops/**`, `.claude/**`
and `CLAUDE.md` but **not one in-scope file** — verified by an empty
`git diff b925199..HEAD -- corpus/services/ corpus/architecture/`.

**So `37 → 18` is the first like-for-like measurement in the milestone.** It says the iter-39 repair
roughly **halved** the residual — and did not approach zero.

## The result

**18 unique in-scope blockers** (21 raw; 3 duplicates, each an independent double-find across the partition
boundary). **21 of 21 re-derived by this iteration before acceptance** — unlike iter-22, not one handed
correction was refuted. Full ledger: `blocker-ledger.md`; re-verification: `adjudication.md`.

**Both pre-registered predictions HELD — the first time in six passes.** Count predicted 8–20, actual
**18**; untouched-file blockers predicted <5, actual **3**; the named prediction that at least one blocker
would sit in text written to *explain* a correction hit four times. Four consecutive passes had refuted
their own predictions. **Holding the instrument fixed is what made the measurement predictable**, which is
itself the strongest confirmation that the earlier series measured instruments.

Location: **15 of 18 (83%) in the 20 files iter-39 edited; 3 in the 20 it never opened** — 0.75 vs 0.15 per
file, a **5×** density ratio (9× → 7.3× → 4.4× → 5×).

## The three that would cost a reader real time

- **The multi-tenancy fence has failed a FIFTH time, toward *"isolation is handled."*** `security_compliance.md:76`
  counts **16** schemas carrying `organization_id` with "no policy of any kind"; **7 more** use
  `OrganizationIDMixin{}`, which declares **0** `Policy()` — and the doc **names that very class as
  unpoliced seven lines earlier at `:69`**. Re-measured by this iteration: **only four files in the entire
  schema dir declare any `Policy()`.** **Found independently by two auditors from two different files** —
  exactly what a disjoint partition exists to produce.
- **A residency claim that is false at HEAD.** `security_compliance.md:175` — *"'Anthropic Direct' is not
  used at all"* — in the **EU Data Residency** section, while `coursebuilder/bedrock.go:108-112` routes
  every coursebuilder call to `api.anthropic.com` whenever `ANTHROPIC_API_KEY` is set. The same sweep added
  an "Anthropic Direct" provider row to `external_services.md:489`.
- **The retracted EU-first ladder is still published verbatim** at `architecture_overview.md:243` — *"Azure
  OpenAI EU → Azure OpenAI US → direct OpenAI"* — in the file most readers hit first, while
  `external_services.md:537` says *"There is **no** ordered EU-first fallback chain."*

## The finding that ends the loop — a 50/50 induced/genuine split

Classified per blocker: **9 of 18 were MANUFACTURED by iter-39's repair** (over-corrections in its own new
blockquotes, a false retraction it authored, a blockquote spliced into a bullet list orphaning the member
that states a legal consequence, half-applied edits that fixed one twin and left the other) and **9 were
genuine pre-existing claims** five passes had missed.

> **The residual is not converging to zero because each repair injects new defects at a rate comparable to
> what it removes.** `37 → 18` is real improvement. But a seventh pass would repair 18, induce ~9, and
> measure ~9–15. **The fixed point of this process is not zero.**

Two supporting observations: **for a fifth consecutive pass every `file:line` anchor a sweep introduced
resolved correctly** (G resolved ~110 across 91 hunks, zero failures) — **the failures are entirely in
prose**, the layer a machine fence could cover and a hand sweep demonstrably cannot. And in **two
consecutive iterations the author of a newly-written rule violated it while writing it** (iter-40's rule 19;
this iter's `D-M257x-41-2`, where iter-40's uniformity claim was verified for five of eight claims and
asserted for all eight — which is how blocker 12 survived).

## What five passes have established as CLEAN (the negative results are load-bearing)

`hiring.md` — **repaired twice, defective after both** — is now clean across ~40 exact anchors. The
*"no `manager` Casbin role"* fix holds (verified three ways incl. a live `p_type='g2'` query). *"Standalone
`authn` is imported by nothing"* holds with a positive control. *"`gen.py` registers exactly nine
arguments"* holds. The **5→4→3→1** subgraph ladder reproduced independently by four auditors. And the
**135-vs-112 denominator ambiguity is RESOLVED in the doc's favour** — 112 is a grep artifact — which
**refutes half of `D-M257x-39-3`'s stated reason** for refusing the tenancy edit.

## Close — 2026-08-02

**Outcome:** the sixth clause-5 pass returned **18** blockers on a byte-identical corpus with an instrument
held fixed — the milestone's first controlled comparison. `37 → 18`: the repair halved the residual and did
not approach zero. **Half the residual (9 of 18) was manufactured by the repair that preceded it.** Both
pre-registered predictions held for the first time in the series. **Nothing was repaired, by
pre-commitment.**
**Type:** tik
**Status:** closed-fixed *(the reading IS this iter's planned deliverable; `overview.md` committed to a
measurement and forbade repair. The 18 are enumerated, anchored and re-verified.)*
**Gate:** NOT MET — 4 of 5. Clause 5 requires 0 blockers; the reading returns 18.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (platform origin `2adcf71`
re-fetched at open, unchanged; occurrence stays 1 of 2) — (4) **user-blocker: y** — (5) cap-reached: n
(2 tiks) — (6) protocol-stop: n — Outcome: **exit-4**
**Decisions:** `D-M257x-41-1` … `D-M257x-41-6`
**Side-deliverables:** **zero words of in-scope corpus text** — deliberately, to preserve the property that
makes the number interpretable. The only corpus edit is `corpus/ops/platform-alignment.md` (outside
clause-5 scope, so it cannot affect the reading), where the protocol-evolution rule requires this iter's
lessons to land: **§5 rule 19 gains the list-derivation clause** (`D-M257x-41-2` — derive the claim list
from the prior pass's ledger, never by hand; state coverage as a fraction) and **§5 gains rule 20 —
*measure what the repair INDUCES, not only what it leaves*.**
**Routes carried forward:**
- `CHECK-M257x-iter41-tenancy-fence-fifth-failure` — supersedes `CHECK-M257x-iter39-tenancy-fence-off-by-one`.
  Denominator now settled (135); mechanism corroborated twice; base count still disputed 23 vs 24.
- `FIX-M257x-iter41-blocker-set` — the 18, fully anchored in `blocker-ledger.md`, ready to apply **if and
  when the user decides the loop should continue**.
- `DOC-M257x-iter41-ops-collateral` — G6, the academy FS-fallback claim at 4 `corpus/ops/demo/**` sites.
- `DOC-M257x-iter41-minors` — **~85** minors with exact anchors across the seven reports (A 11 · B 15 ·
  C 13 · D 16 · E 15 · F 8). Clause 5's *"YELLOW with 0 blockers"* admits them.
- `CHECK-M257x-iter38-ai-act-classification` — unchanged; still needs an owner outside this milestone.
**Lessons:**
1. **Change one variable.** Five passes varied corpus and instrument together and produced a series that
   meant nothing. One pass varied only the corpus and produced the milestone's only interpretable number —
   and the first predictions that held.
2. **Measure the repair, not just the residual.** The induced/genuine split is the decision-relevant
   quantity and no prior pass computed it. Without it, `37 → 18` reads as convergence; with it, it reads as
   a process with a non-zero fixed point.
3. **The anchors are fine; the prose is not.** Fifth consecutive pass with zero anchor failures and every
   defect in surrounding prose. This is now a strong enough regularity to design a fence around.
