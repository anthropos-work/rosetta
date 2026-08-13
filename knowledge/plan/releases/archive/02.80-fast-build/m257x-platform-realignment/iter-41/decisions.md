# iter-41 — decisions

## D-M257x-41-1: hold the instrument FIXED, overruling the routed advice to re-partition — 2026-08-02

Run 21's hand-off advised re-partitioning again (§5 rule 18(b)). **Overruled deliberately, and the reason
is the whole design of this iter.** iter-39's central finding is that `25 → 13 → 11 → 17 → 37` measured
five *instruments*, not the corpus. A sixth number from a seventh instrument would extend that series and
settle nothing; clause 5 cannot be closed or refuted by a measurement that is not comparable to anything.

The tension with rule 18(b) is real and was resolved **by the method rather than by ignoring the rule**:
holding the *method* fixed does not hold the *partition* fixed, because iter-39's repairs moved 20 files'
line counts (+739/−250) and the same size-sort therefore deals a different hand. The properties rule 18(b)
exists to preserve were checked explicitly and all survive — `ai_architecture`/`security_compliance` split,
the three "what is `app`" docs split three ways, the merged/archived family in five different hands.

**It paid immediately.** Two of the three highest-value blockers are independent double-finds across the
partition boundary (the tenancy undercount, found from `security_compliance.md` **and** from
`architecture_overview.md`; the Anthropic-Direct residency claim, found by a full-read auditor **and** the
adversarial reader). And for the first time in the series, **both pre-registered predictions held** —
which is itself the strongest confirmation that the earlier numbers were measuring instruments.

## D-M257x-41-2: record that iter-40's uniformity claim was OVERSTATED — 2026-08-02

iter-40 reported that *"the 40 in-scope files are UNIFORM on all of them"* for the eight claims it swept.
**That was verified for five of the eight and asserted for all eight.** Two of this pass's blockers are the
consequence:

- Blocker 12 — `graphql-wundergraph.md:79` still asserts `:5050` as a live port, in present tense, inside
  clause-5 scope. iter-40 fixed `:5050` at eight sites in `corpus/ops/**` and `.claude/**` and **never ran
  the in-scope grep for it**.
- G6 — the academy FS-catalog-fallback claim survives at four `corpus/ops/demo/**` sites. It was not on
  iter-40's claim list at all, though iter-39 had retracted it in `ant-academy.md`.

**The lesson is exactly rule 19's own post-condition clause, under-applied by its author.** The rule says
*"re-run the same grep as a post-condition"*; iter-40 ran it per *claim it had listed*, never per *claim the
prior pass had adjudicated*. **A claim-scoped repair needs its claim list derived from the prior pass's
ledger, not assembled by hand.** Folded into `platform-alignment.md` §5 rule 19 as the list-derivation
clause.

This is the second consecutive iteration in which the author of a rule violated it while writing it. That
is not carelessness twice; it is evidence that **a hand-applied discipline does not survive contact with a
corpus this size** — which is the substance of the escalation below.

## D-M257x-41-3: do NOT repair these 18 — the iteration STOPS at the measurement — 2026-08-02

Pre-committed in `overview.md` before any auditor reported, and honoured:

> *Repairs are **out of scope for this iter**: fixing findings here would destroy the byte-identical
> property that makes the number mean anything.*

Three reasons, in order of weight:

1. **The measurement is the deliverable and it is not yet spent.** These 18 are the only blocker set in the
   milestone measured by an instrument identical to its predecessor's. Repairing them now converts the one
   controlled comparison the milestone has into another confounded one.
2. **The escalation is pre-committed.** `overview.md` states that a non-zero result means STOP, do not open
   a seventh pass, exit `user-blocker`. Repairing and re-reading *is* the seventh pass under another name.
3. **The 50/50 induced/genuine split says repairing is not obviously net-positive.** Nine of the 18 were
   manufactured by the previous repair. A tenth pass of repair-then-measure has an expected yield that the
   data no longer supports assuming is positive.

## D-M257x-41-4: the tenancy fence — measure, corroborate, and STILL do not change the number — 2026-08-02

`D-M257x-39-3` refused a one-site numeric edit on two grounds: the home derivation lived in a file the
repairer did not own, and the denominator was ambiguous (135 vs 112).

**One ground is now refuted and the other is satisfied**: auditor E resolved the denominator — **135 is
correct, and 112 is a grep artifact** (`grep '^\tent.Schema$'` misses 23 gofmt one-liners
`struct{ ent.Schema }`) — and both the home file and its twin were read this pass, by different auditors,
who **independently reached the same conclusion**.

**The number is still not changed here**, because this iter does not repair (`D-M257x-41-3`) and because
the two auditors' base counts disagree (24 vs 23) even while agreeing on the mechanism and the direction.
What *is* now established and must survive into any future repair:

- The 7 `OrganizationIDMixin{}` schemas carry `organization_id` with **no policy of any kind** —
  re-measured by this iteration, not taken on an auditor's word.
- **Only four files in the entire schema dir declare any `Policy()`.**
- `security_compliance.md` **contradicts itself seven lines apart** (`:69` names the class as unpoliced;
  `:76` excludes it from the unpoliced count).
- The error runs toward *"isolation is handled"* — the dangerous direction, for the **fifth** generation.

Routed as `CHECK-M257x-iter41-tenancy-fence-fifth-failure`, superseding
`CHECK-M257x-iter39-tenancy-fence-off-by-one`. **A fence wrong five times in both directions should stop
being hand-maintained** — which is the specific case for the derived-fact fence below.

## D-M257x-41-5: `D-M257x-39-4`'s one-way-door check is CORRECTED — 2026-08-02

iter-39 recorded: *"the adversarial pass confirmed neither `ai_architecture.md` nor `security_compliance.md`
now asserts a legal conclusion; both defer explicitly to counsel. That was the single most dangerous thing
this sweep could have done and it did not do it."*

**Half of that is wrong.** `security_compliance.md:7` does defer properly. But `:205` — a trailing bullet
orphaned when the retraction blockquote was spliced into the middle of a bullet list — states the operative
**legal consequence** as settled fact, three lines after `:202` says *"Do not cite this section as evidence
of a Limited-Risk classification."*

The corpus does not assert a classification wholesale, so iter-39's claim is over-broad rather than
inverted. Recorded precisely because **an over-broad all-clear on a compliance page is itself a compliance
risk**, and because it demonstrates the failure mode a fixed-instrument re-read exists to catch: iter-39's
adversarial pass looked for an *assertion* and found none, because the assertion was a **list member the
splice orphaned**, not a sentence anyone wrote.

The EU-AI-Act classification itself remains routed to an owner outside this milestone
(`CHECK-M257x-iter38-ai-act-classification`) and is **not** settled here.

## D-M257x-41-6: escalate — clause 5 asks for a property this process cannot deliver — 2026-08-02

**This is the user-blocker, and it is deliberately not a decision this iteration or the orchestrator takes.**

Clause 5 requires *"KB-fidelity audit GREEN, or YELLOW with **0 blockers**"*. Six passes have returned
**25, 13, 11, 17, 37, 18**. The sixth is the only one comparable to its predecessor, and it shows the
repair worked (37 → 18 on a fixed instrument) **and did not approach zero**.

The reason it will not approach zero is now measured rather than suspected: **9 of the 18 were created by
the repair that preceded them.** The process has a fixed point, and the fixed point is not zero.

Three further facts bear on whether the clause is satisfiable as written:

- **A hand-applied discipline does not survive this corpus.** In two consecutive iterations the author of a
  newly-written rule violated it while writing it (iter-40 rule 19; iter-41's `D-M257x-41-2`).
- **The `file:line` layer is already reliable.** For a fifth consecutive pass every introduced anchor
  resolved correctly. **The failures are entirely in prose** — summary lines, orphaned bullets, universals,
  and twin claims in sibling files. That is precisely the layer a machine fence *could* cover and a human
  sweep demonstrably cannot.
- **Six passes have cost ~30 sub-agents** and the corpus is materially better — `hiring.md`, repaired twice
  and defective after both, is now clean across ~40 anchors — but "materially better" is not the clause.

The open question is therefore **not** *"fix the 18."* It is **whether a hand-maintained corpus can satisfy
a zero-blocker clause at all**, and if not, whether the clause should be re-cut (e.g. to "0 blockers in the
machine-fenceable classes + a bounded, ledgered prose residual"). **That is a change to the gate, which is
the user's call.** Recommendation on `CHECK-M257x-iter33-derived-fact-fence` is in the close section.
