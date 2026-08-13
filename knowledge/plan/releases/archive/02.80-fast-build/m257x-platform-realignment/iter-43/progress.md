**Type:** tik (`iter_shape: tooling`) — protocol: [`corpus/ops/platform-alignment.md`](../../../../../corpus/ops/platform-alignment.md)

# iter-43 — `FENCE-M257x-iter42-claim-twin`, watched going RED

**Active strategy:** `TOK-02` — *fence the prose the way the anchors are fenced*. **Step 1 of five, and
step 1 only.**

## What ran

| phase | outcome |
|---|---|
| **A** re-survey | fixture INTACT (`git diff 103ad31..HEAD -- corpus/services/ corpus/architecture/` empty); 4 blockers re-verified live; platform origin `2adcf71` unchanged; rext pins match |
| **B** derivation | `stack-core/claim_ledger.py` — parses the audits' own blocker-ledgers by table STRUCTURE; 4 files, 85 blocker rows → **36 claims / 39 refuted forms** |
| **C** fence | `stack-core/claim_twin_guard.py` — markdown-normalized, tree-wide (`corpus/**`, `.claude/skills/**`, `CLAUDE.md`, `README.md`), 112 files scanned |
| **D** watch it go RED | **18 sites RED**, covering **16 of iter-41's 18 blockers** — see [`fence-red-measurement.md`](./fence-red-measurement.md) |
| **E** tests + battery | 15 unit/answer-key tests + an **8-mutant battery**; `stack-core` **415 / 14F** vs a 396 / 14F baseline (+19 tests, 0 new failures) |

## The result

**16 of 18 detected, at the anchors iter-41 recorded — and 13 of 13 of the class the fence was built
for** (the corpus contradicting itself). Pre-registered floor was **≥ 12**; it cleared it by four.

The two misses are the declared scope boundary, not a shortfall against it:

- **#10** (*"Language: Go 1.25"*) — quoted form is 17 normalized chars, under the fragment floor; reported
  `UNMATCHABLE` **by name**, never dropped. iter-42 routed it to a **value fence**.
- **#16** (`messenger.md:110`) — iter-41's claim column *paraphrases* rather than quotes, so there is no
  pattern to derive. iter-42 routed it to a **symbol-aware anchor check**.

> **Both blind spots land in the two other instruments' territory.** That is the load-bearing result:
> iter-42's three-way split of the residual was a plan, and the first instrument built against it missed
> exactly what the plan said it would.

**It also found a claim no pass has ever caught** — `corpus/ops/demo/coverage-protocol.md:629` restates a
claim iter-34 refuted and repaired *inside* the audited scope. §5 rule 19's corollary (*a claim leaks to
the edge of the previous repair's scope and stops there*) measured on the fence's first run. **Routed,
not repaired.**

**GREEN control, and it is a real measurement:** 36 claims derived, 20 fired, **16 adjudicated claims are
absent from the entire tree** — repaired everywhere by earlier passes, and the fence says so positively
rather than by silence. Plus a synthetic GREEN twin of all 18 sites: **zero fire.**

**Battery:** 8 mutants, every one matching its **declared** verdict — 1 declared-GREEN no-op that
**survived**, 7 kills of which **5 are inversions**, **7 distinct failure signatures**, baseline GREEN
before and after, every mutant `py_compile`d before its run. `fragment-floor-collapsed` (30 → 3 chars)
leaves every answer-key site firing and is caught **only** by the GREEN twin — a fence measured by REDs
alone would have shipped it.

## Nothing was repaired

`D-M257x-42-3` obeyed literally, including on the one finding it was tempting to fix. The only corpus
edit is **new** text — the §8 fourth-layer section documenting the fence — and the fence was re-run as a
**post-condition** over that edit: 18 hits, byte-identical set. That post-condition is TOK-02 step 2 in
embryo.

## Close — 2026-08-02

**Outcome:** the claim-twin fence exists, is derived from the audits' own ledgers, and was **watched going
RED on 18 sites covering 16 of the 18 known blockers (13/13 of its target class)** with a GREEN control
and an 8-mutant battery — before a single blocker was repaired, while the fixture still existed.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5. Clause 5 stays at **18** by construction: this iteration repairs nothing.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this was a tik) — (3) re-scope: n (platform origin `2adcf71` re-fetched at open and close, unchanged; trigger stays at occurrence 1 of 2) — (4) user-blocker: n — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n — **Outcome: continue**
**Decisions:** D-M257x-43-1 … D-M257x-43-6 (this iter's `decisions.md`)
**Side-deliverables:** none — the `platform-alignment.md` §8 section and the `stack-core/README.md` rows
are the protocol-evolution obligation of this iter's own scope, not unrelated fixes.
**Routes carried forward:** `FIX-M257x-iter43-coverage-protocol-livepath` · `CHECK-M257x-iter43-value-fence`
· `CHECK-M257x-iter43-symbol-anchor` · `CHECK-M257x-iter43-markdown-lint`

**Lessons:**

- **A residual classified by cheapest reaching instrument gives an instrument a falsifiable charter.** The
  fence missed exactly the two blockers iter-42 predicted it would, and hit 13 of 13 of the class it was
  built for. Neither number would have been interpretable against an undifferentiated count of 18.
- **A fixture with a known answer key is worth more than the repair it delays.** One iteration spent not
  repairing bought a permanent RED watch. After TOK-02 step 4 this could never have been built.
- **Watch a fence go GREEN on purpose, too.** The single mutant that only the GREEN twin caught
  (`fragment-floor-collapsed`) would have passed a battery of eight REDs. Presence of detection is not
  discrimination — the same lesson §8 rule 5 learned from the inverted mutant, seen from the other side.
- **When a fence contradicts a design decision its own protocol records, reconcile it in the same commit.**
  §8's *"keep `.md` prose out of scope"* was true of the fence it was written for and reads as a
  contradiction next to this one. Leaving it would have manufactured the exact defect class the fence
  exists to catch, in the document governing the fence.
