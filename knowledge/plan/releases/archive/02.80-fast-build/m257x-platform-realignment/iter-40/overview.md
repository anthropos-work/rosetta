---
iter: 40
milestone: M257x
iteration_type: tik
iter_shape: cleanup
status: closed-fixed
opened: 2026-08-02
---

# iter-40 — `FIX-M257x-iter39-claim-scoped-repair`

**Active strategy reference:** `TOK-01: instrument first, then follow` (milestone-root `decisions.md`).

## Step 0 — re-survey (run before this plan was committed to)

The routed target still holds, and the re-survey **changed its shape** — which is the point of the step.

Run 21's hand-off framed the claim-scoped repair as a fix for cross-file drift *inside* the clause-5 file
set (iter-39's C1–C8, five of eight being a claim corrected in one file and left standing in a twin). A
whole-tree grep for every claim iter-38/39 adjudicated says otherwise:

- **Inside the 40 in-scope files the corpus is UNIFORM on all of them.** Taxonomy figures, `--template`,
  the subgraph count, `organization_id on every table`, the "five internal Go libraries" — each is either
  corrected everywhere or fenced everywhere. iter-39's adversarial half did land its fixes corpus-wide
  *within scope*.
- **Every surviving instance is past the scope boundary** — `corpus/ops/**`, `CLAUDE.md`, `corpus/README.md`,
  `.claude/skills/**`.

That is exactly what a **file-partition-scoped** repair predicts. The partition *was* the 40 files; a claim
does not respect it, so the drift leaked to the edge and stopped. The repair is therefore not "finish
iter-39's sweep" but "extend it to the claim's real extent."

## Cluster / target identified

Two deliverables, both routed:

1. `FIX-M257x-iter39-claim-scoped-repair` — codify candidate **§5 rule 19** (`D-M257x-39-2`) in the protocol
   doc, and execute it against the surviving out-of-scope instances.
2. `DOC-M257x-iter38-ops-collateral` — four `corpus/ops/**` docs asserting what iter-38 retracted. Open for
   two runs. The same class, so it belongs in the same pass.

## Hypothesis

Not about the corpus's content. The procedural claim under test: **a repair scoped by CLAIM reaches sites a
repair scoped by FILE structurally cannot**, and the surviving instances will cluster at the boundary of the
previous repair's file partition rather than being scattered.

The re-survey above already confirms the second half. What this iter adds is the fix.

## Planned scope — the claim list, declared up front

This is a **cleanup-shaped** iter: each claim is a *planned* line of work, so the scope-creep tripwire counts
against this list and not against single-target tik shape. A claim not on this list is routed, not swept.

| # | claim | adjudicated | surviving sites |
|---|---|---|---|
| A | "60K skills / 18K roles" | iter-39 #35 — 18K **REFUTED** (≥22,470 public), 60K **UNSUPPORTED, not refuted** (42,790 floor). **Two verdicts; do not collapse them** | 12: `.claude/skills/stack-snapshot`, `CLAUDE.md`×2, 5 `corpus/ops/**` docs |
| B | `gen.py --template` | iter-39 #33 — the flag does not exist and `parse_known_args` **silently absorbs** it | 3: `CLAUDE.md:432`, `run_guide.md:347`, `setup_guide.md:595` |
| C | cms-in-app took the supergraph **2 → 1** | iter-39 #15 — it was **3 → 1**; the "2→1" comes from `915da06`'s own (wrong) commit subject | 2: `CLAUDE.md:183,212` |
| D | "the **five** internal Go libraries", authn among them | iter-39 #29 — `authn` is imported by nothing; **four** modules, authn ships inside colony | 2: `CLAUDE.md:362`, `corpus/README.md:26` |
| E | "`organization_id` on **every** table" | iter-33 (retracted in `security_compliance.md`), re-found iter-33 adversarial | 2: `CLAUDE.md:247`, `corpus/ops/safety.md:1025` |
| F | `:5050` is a live local router port | iter-38/39 #20/#34 (prod half) + platform `2adcf71` deleted the service | 7 across `.claude/skills/dev-up/**` + `dev-for-dummies/reference.md` |
| G | "simulation scoring is deterministic, NOT AI-scored" | iter-38 (retracted) → iter-39 #37 (the honest claim is **most**, not all) | `playthroughs.md:521-523` — **changes test-design behaviour** |
| H | `SkillPathSessionService` in the RPC mux; singular `academy_*` tables | iter-38 | `platform_repo.md:111`, `content-stories-{routes,spec}.md` |

Plus: **§5 rule 19** authored into `corpus/ops/platform-alignment.md` (the protocol-evolution rule requires a
generalising lesson to land in the protocol doc in the same commit).

## Repair discipline — binding on this iter

1. **Propagate the adjudicated verdict verbatim; derive nothing new.** Every edit either restates a
   correction already established and fenced in the clause-5 corpus, or links to its canonical anchor. This
   iter must not author a *new* derived claim — that is what makes repaired text the highest-risk text in
   the corpus (§5 rule 18), and pass six reads whatever this iter writes.
2. **Where a figure has two different verdicts, carry both.** Collapsing "REFUTED" and "UNSUPPORTED" into
   one word is the specific error the hand-off warns about.
3. **In-scope files are touched only for uniformity**, never re-derived. (In the event: none needed it.)

## Phase plan

- **Phase A** — the whole-tree claim survey (done at Step 0; recorded above).
- **Phase B** — author §5 rule 19.
- **Phase C** — execute claims A–H as anchored edits.
- **Phase D** — re-measure: 5 corpus guards + the orphaned-continuation grep + a re-run of the whole-tree
  survey showing each claim now resolves uniformly.

## Escalation conditions

- A platform commit landing mid-iter → **re-scope trigger occurrence 2** → STOP and escalate.
- A surviving site whose correct form is **not already established** in the clause-5 corpus → do NOT derive
  it here; route it. This iter propagates, it does not adjudicate.
- Anything needing a platform-repo edit → route (v2.8 constraint is binding).

## Acceptable close-no-lift outcomes

This iter cannot move the gate — clause 5 is met by a *reading*, and this is a repair. `closed-fixed` is
earned by the claim list landing with guards green. A finding that the surviving sites are fewer or more
than surveyed is a result, not a failure.

## Expected lift

**Zero gate movement, by construction.** The lift is on pass six's *comparability*: a corpus that
contradicts itself across the scope boundary makes the next auditor spend budget adjudicating rather than
measuring (`D-M257x-39-2`).
