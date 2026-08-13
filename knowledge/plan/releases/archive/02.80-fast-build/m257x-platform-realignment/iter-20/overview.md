---
milestone: M257x
iter: 20
iteration_type: tik
status: closed-fixed
opened: 2026-08-01
---

# iter-20 — clause 3: the migration-status map, and its both-ways fence

**Active strategy reference:** `TOK-01: instrument first, then follow` — step 4, *"then the corpus — the
migration-status map with its two states per row (prod vs fresh stack), and the reconciliation sweep."*
Steps 1–3 (pin, derivation, the two schema fences) are landed; the map is the next step in TOK-01's own order.

**Step 0 re-survey (mandatory, done before targeting):**
- TOK-01's `Next-tik direction` named iter-02 and is long superseded; the handoff named clause 2's
  `library_category` class. **Substitution, recorded:** clauses 3 and 5 are the two untouched clauses and
  share their research — the map is the source of truth clause 5's sweep must reconcile *against*, so
  building it first makes the sweep mechanical. Clause 2's two-field class is unchanged and stays routed.
- Re-measured at open: platform origin HEAD **`2adcf71`** (unchanged — re-scope trigger stays at occurrence
  1 of 2); org repo census independently **reproduced at 93** (iter-01's number); rext pin consistent across
  `.agentspace/rext.tag`, authoring copy and the `stack-demo` consumption clone.
- `corpus/architecture/platform-migration-status.md` **does not exist** — and
  `corpus/ops/platform-alignment.md` §6 already links to it, so the link is dead today.
- `stack-core/platform_alignment_guard.py` **does not exist** — §8's first fence layer is the only one of
  three still unbuilt (the static schema fence and the live assert both landed at iter-06/iter-08).

**Cluster / target identified:** gate clause 3 — *a checked-in migration-status map covering every service the
platform has ever had, each claim cited to platform source, machine-fenced against `repos.yml` both ways, and
including net-new repos appearing in neither `repos.yml` nor the corpus.*

**Hypothesis:** the map can be authored entirely from measurement already available on this box (the
`stack-demo` clone set at origin HEAD + the GitHub org API), and a both-ways fence over it is a
`corpus_index_guard`-shaped guard — a rext-owned Python guard reading `CORPUS_ROOT` plus the platform's own
`repos.yml`. No stack bring-up is required, which is why it is the right target for a budget-constrained run.

**Expected lift:** gate 2 of 5 → **3 of 5**.

**Phase plan (protocol `corpus/ops/platform-alignment.md`):** §4 detection signals 1–6 re-run at origin HEAD →
§6 classification (two states per row, seven-value vocabulary) → §8 layer-1 fence, **watched going RED in both
directions** with a GREEN no-op control (§8 rule 5) → corpus index row (`corpus_index_guard` must stay green).

**Escalation conditions:** a second platform commit landing mid-iter fires the re-scope trigger (occurrence 2
of 2) → exit `re-scope-trigger`. A row that cannot be cited to platform source is not written as a claim.

**Acceptable close-no-lift outcomes:** if the fence cannot be made to fail in one of its two directions, the
fence is not trustworthy and the clause is not claimable — recording that with the falsification is a complete
outcome (§8: *watched going RED before it is trusted*).
