**Type:** tik — `TOK-07` **step 2, second half** (the repair itself). Protocol:
[`corpus/ops/platform-alignment.md`](../../../../../corpus/ops/platform-alignment.md).

# iter-115 — the repair, over the enumerated set

## What was measured before anything was written

- `git diff --stat 461b547 HEAD -- corpus/` — **empty**. The corpus had not moved since the enumeration,
  so `iter-113/enumeration.json`'s line numbers were live coordinates rather than a stale snapshot.
- Pre-repair baseline, reproduced: `repair_reach_guard.py --enumeration iter-113/enumeration.json
  --range 461b547` → **`reach 0/71 = 0.0 %`**, `denominator: corpus-derived-per-predicate`.
- All **14** platform clones re-verified against iter-109's ground-truth table and matching byte for
  byte: `platform 0c91421d · app ad9f3c49 · next-web-app 8297c684 · sentinel f2c46190 · studio-desk
  41ee3575 · ant-academy 22df69dd · cms ca50c817 · jobsimulation 462343b0 · messenger fa47850d ·
  storage 4ce8ece5 · roadrunner 87d8d443 · graphql-wundergraph 60c229f3 · app/studio + cms/studio
  aeec036a`. **No fetch** (§5 rule 41a). Read-only throughout.

## The result

```
repair-reach: 71 enumerated site(s) over 24 predicate(s); 73 hunk(s) over 24 path(s); tolerance=3
repair-reach: denominator: corpus-derived-per-predicate (enumeration.json)
repair-reach: reach 71/71 = 100.0%
repair-reach: OK — every booked finding was reached or dispositioned.
```

exit 0. **And the aggregate is not the claim.** Re-derived independently, site by site from the same
enumeration and this iter's own diff hunks: **24 of 24 predicates closed at EVERY enumerated
instance** — P21 22/22 · P10 11/11 · P09 6/6 · P15 5/5 · P13 3/3 · P02 P03 P07 P12 P18 2/2 each ·
the remaining eleven 1/1. `TOK-07` rule 3's test is *"closed at every instance, or not closed"*, and
that is the number above, not the 100 %.

**Both promoted pair-halves landed with their twins, in one commit each** — which is the whole reason
they were promoted:

- **P10** `cms.md:171` (the mermaid `exec gen.py` line) with `cms.md:287` (the `bash -c` prose);
- **P12** `ai_architecture.md:212` (`:1594-1600`) with `jobsimulation.md:160` (`:1594-1597`) — the
  same-fact-different-pin pair. Both halves now carry **the same path and the same range**, which is
  what stops the repair from manufacturing the `external_services.md:554`/`:565` self-contradiction
  iter-108 created.

## The three queued findings, each resolved against source rather than against the ledger

**P08 — the pin was off by two, and the repair DELETES the pin rather than moving it.** Re-derived:
`:496` is the closing line of the `> **✅ CORRECTED M219 …**` blockquote (`:476-496`); the `⚠⚠ M51
iter-08/09` block opens at **`:498`** and the quoted parenthetical is at `:500`. Re-pinning to `:498`
was the obvious move and it is the wrong one: this is the **third generation** of one same-file anchor
(`:458` → `:459` → `:496`), each off by a handful of lines, in a passage published as a worked example
of a repaired anchor. Three generations is not three accidents. The number is gone; the construct is
named; the self-heal clause is retained for reader cost but no longer has a false number to heal.

**P13 — the superlative had a counter-example inside the corpus.** Re-derived over the whole history
of `platform`: **18 distinct git-URL repo contexts**, and building that way was the platform
**default** until `a2a3ee6`, with `realtime` holding one until `c17cc9a`. `customerio-sync` was the
**last**, never the only. `external_services.md`'s own *Build Context* note records the router was
also once built from a git+url context — one file away from the claim.

**P24 — one survivor against ten witnesses.** `838d907` deleted the messenger compose block;
`0c91421d` declares five services and `git grep -n messenger … docker-compose.yml common.yml
repos.yml` returns **only comments**. Ten corpus sites already recorded the deletion, **two of them in
`sentinel.md` itself**, neither framed as a retraction. The conclusion (*messenger is not a caller*)
was true and was re-derived; the present-tense **evidence clause** was what had expired.

## What the repair found that the enumeration had not

Three sites publish an enumerated predicate and were **not** in the enumerated set. All three were
repaired anyway — rule 3 does not care where a twin came from — and all three are recorded against
**`FIX-M257x-iter113-adjudication-is-judgement`**, which is exactly the evidence that item was routed
open to collect:

| site | predicate | why it matters |
|---|---|---|
| `jobsimulation.md` *AI providers* bullet | **P02** (*"via the shared `ai` library"*) | a **third** instance of a predicate enumerated at two |
| `architecture_overview.md:423` | **P22** (*"public subnets (ALB, Cosmo Router)"*) | adjudicator 4 named it as a second anchor; the enumerator **excluded** it as a subject-mention. It carries the claim verbatim |
| `cms.md` mermaid participant | **P10** | promoted by the enumerator, and correct — but its sibling `:171` was the twin |

The `architecture_overview.md:423` case is the sharp one: an adjudicator *booked* it, and the
ceiling adjudication then excluded it with a reason that reads plausibly and is wrong. **15 of 16
small-class verdicts rest on one agent's readings**, and this is one measured miss inside that set.

## Corrections re-derived, not copied — and one was materially different from the ledger

§5 rule 22's own lesson (*"verify the CORRECTION against platform source before you apply it, every
time"*) earned its keep at `external_services.md:672`: the ledger route would have carried the
existing text forward, and re-reading the file measured **six** env names at `gen.py:45-48`
(`AZURE_API_KEY`, `AZURE_ENDPOINT`, `OPENAI_API_KEY`, `OPENAI_ENDPOINT`, `ANTHROPIC_API_KEY`,
`ANTHROPIC_ENDPOINT`) where the corpus published **three** at `:45-47` — a range cut one line short of
the `ANTHROPIC_*` pair. Not booked by any seat; found only by opening the file. **And the tree matters
as much as the range**: `studio` is a *nested* checkout at `aeec036a`, so `git show ad9f3c49:studio/gen.py`
reads the host ref and is the wrong grep.

## The repair broke four of its own anchors, and the fence caught every one

`repair_postcondition` refused the commit **four times**, each time correctly, and each refusal is a
finding rather than an obstacle:

1. **`jobsimulation.md` → `ai_architecture.md:218`** — a cross-file line pin I *authored in the same
   commit that shifted the target*. Every intra-corpus reference this iter introduces is now a
   **construct** reference (section or bullet name), not a line.
2. **`service_taxonomy.md`'s table row numbers** — rotted a **second** time (iter-100 shifted them,
   iter-102 added row names as a self-heal, iter-115 shifted them again and the Jobsimulation number
   landed on the table **header**). Numbers deleted; the names were already there.
3. **`backend.md` → a FROZEN rext test fixture, cited by line.** `anchor_construct_guard` strips the
   `stack-core/tests/fixtures/repair_leak/{pre,post}/` prefix and resolves `corpus/services/cms.md:157`
   against the **live** corpus — so a corpus repair false-REDs a citation of a fixture that did not
   move. Routed: **`FIX-M257x-iter115-anchor-guard-resolves-fixture-paths-live`**.
4. **A bare `:335` in `platform-alignment.md`** that bound to the wrong document — it followed a
   `backend.md:54` citation, so the guard resolved it against `backend.md`. Named in full now.

And the sharpest one, in the protocol doc itself: **§5 rule 22's worked example rotted a FOURTH time,
moved by this repair — whose entire subject is that class.** Its line numbers are deleted; the
construct was already named beside them.

## Tests — with the invocation stated

`/usr/bin/python3 -m pytest stack-core/tests/{test_repair_reach_guard,test_predicate_enumerator,
test_claim_twin_guard,test_repair_postcondition,test_repair_postcondition_audit_mode,
test_anchor_offset_guard,test_m257x_repair_reach_mutation_battery,
test_m257x_repair_postcondition_mutation_battery,test_m257x_claim_twin_mutation_battery,
test_claim_twin_guard_iter47_answer_key,test_claim_twin_guard_iter48_answer_key}.py -q`
→ **1 failed, 183 passed in 118.47 s.**

The one failure is **`test_claim_twin_guard_iter48_answer_key.py::TestIter48AnswerKey::
test_02_the_green_twin_of_every_site_stays_SILENT`** — **the known pre-existing failure**, the same
test id recorded at iter-111's full-suite run (`1 failed, 1011 passed in 1090.88 s`), and its
assertion subject is the test's own synthetic fixture corpus (`corpus/04.md:1`, `corpus/05.md:1`
against `iter-49/raw/C.md:57`), which this repair does not touch. **Not introduced here, and not
papered over either.**

## What this iter did NOT do

**The read.** `TOK-07` step 3 is untouched: **no reading was taken**, so **no `N` movement is
claimed** and **`P` is UNMEASURED, not unmoved** (§9's iter-type refinement, in its mandated words).
The sequence was declared in advance by `TOK-07` — enumerate → repair → read — and steps 1 and 2 are
now complete, so step 3 is iter-116's entire content and is unblocked: the enumeration is checked in
and settled, the repair is committed at `71/71`, the instrument is untouched, and the trees are
frozen at the ground-truth shas.

Gate unchanged at **4 of 5**. Clause 5 is not re-cut, narrowed, reinterpreted or argued.

Zero platform-repo edits · `stack-demo/**` untouched · no clone fetched (§5 rule 41a) ·
`rosetta-extensions` on `main`, no tag cut · Chapman stays retired, no point estimate quoted.

---

## Close — 2026-08-07

**Outcome:** The repair `TOK-07` was authored to make possible has landed: **71/71 enumerated sites,
`denominator: corpus-derived-per-predicate`, exit 0** — and, the claim that actually matters, **24 of
24 predicates closed at every enumerated instance**, both promoted pair-halves with their twins, in
seven incremental commits so a death would have cost one predicate rather than the run. Three sites
that publish an enumerated predicate but were **not** enumerated were repaired anyway and booked as
measured evidence for `FIX-M257x-iter113-adjudication-is-judgement`. **No reading was taken; `P` is
UNMEASURED, not unmoved.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (the `P ≥ 15`
falsification condition cannot fire on an unmeasured `P`) — (4) user-blocker: n — (5) cap-reached: n
(1 tik this session) — (6) protocol-stop: n — (7) budget-exhausted: **y — between iters, tree clean,
both repos pushed** — Outcome: exit-7
**Decisions:** D-M257x-115-1, D-M257x-115-2, D-M257x-115-3
**Side-deliverables:** none — the four anchor repairs the fence forced are repairs of this iter's own
line shifts, i.e. planned-scope cleanup, not unrelated fixes.
**Routes carried forward:**
- **`TOK-07` step 3 — THE READ** → iter-116, and it is the whole of it. Unblocked; full discipline
  (instrument verbatim + sha re-checked, pre-registration sealed in its own commit before any seat is
  dealt, seats committed verbatim as they land, adjudicate before reporting, upheld rate reported
  twice, no repair inside the measuring pass, band on **both** `P` and `N`). **If it returns
  `P ≥ 15`, do NOT author TOK-08** — `TOK-07` pre-registered that as refuting repair-and-read, and
  the next move is a re-scope conversation with the user.
- **`FIX-M257x-iter115-anchor-guard-resolves-fixture-paths-live`** → open, **net-new**. A guard that
  resolves a frozen fixture's path against the live tree will false-RED the next corpus repair too.
- `FIX-M257x-iter113-adjudication-is-judgement` → open, **and now carrying three measured misses**
  (P02's third site, P22's second anchor, both named above) rather than only a design argument.
- `FIX-M257x-iter111-staged-battery-dependency-is-underived` → open
- `FIX-M257x-iter111-buildbench-parse-json-is-a-noop-flag` → open
- `FIX-M257x-iter107-drift-fence-satisfiable-by-prose` → open, de-ranked
- `DEF-M257x-iter101-briefing-rext-tree` → open
**Lessons:**
- **A repair of anchor rot induces anchor rot, and the fence is the only thing that notices.** Four
  refusals, one of them on §5 rule 22's own worked example — its **fourth** generation, produced by
  the repair whose subject is that class. **Never author an intra-corpus line pin in the commit that
  moves lines**; name the construct. This is now the standing rule for repair passes, not advice.
- **When an anchor has rotted three times, delete it — do not re-derive it a fourth.** A same-file
  line pin in a growing document has a measured recurrence rate here, and a self-heal clause lowers a
  reader's cost without making the number true.
- **Grade the predicate, never the percentage.** `71/71 = 100 %` and *"24 of 24 closed at every
  instance"* are different claims, and iter-108 is the proof: it was **arithmetically correct at
  100 %** over the wrong set. Report the second, site by site, or the first will be read as the second.
- **The enumeration is a floor on the repair, not a ceiling.** Three unenumerated instances turned up
  by simply reading the files being edited — one of them booked by an adjudicator and then excluded
  by the ceiling pass. Repair them and record them; renegotiating the denominator to include them
  would have destroyed the grade's meaning.
