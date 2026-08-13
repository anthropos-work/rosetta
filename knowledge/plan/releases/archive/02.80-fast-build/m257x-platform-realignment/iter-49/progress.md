**Type:** tik (declared multi-step shape — two fences, repair, read — per §Phase-2 carve-out)

# iter-49 — the two named fences, the twelve, and the ninth reading

Three planned steps, all three landed.

1. **`FENCE-M257x-iter49-numeric-leak`** (`value_change_guard.py`) — the leak fence's blind spot was
   **structural, not a knob** (`D-M257x-48-9` measured that lowering K bought two false positives and
   still missed). This fence asks a different question — *did the old VALUE survive?* — by word-diffing
   each hunk's removed against added text, keeping the **pairing** `repair_leak_guard` discards.
   **Watched RED on rosetta `301d61a`**, a real already-committed incomplete repair: 5 sites, of which
   **three were independently adjudicated blockers by two different audits** (iter-48 #3, iter-47 #5,
   iter-48 #9). The 2 false positives are **kept and pinned**, with the tuning that would remove them
   measured and rejected — it would cost a proven true positive (`D-M257x-49-2`).
   **The gap it fills is asserted, not assumed:** a test runs `repair_leak_guard` on the same commit and
   requires it *not* to report the site.

2. **`FENCE-M257x-iter49-audit-commit-mode`** — `repair_postcondition --audit-commit`. It **writes
   nothing**: the ratchet's monotonicity is the contract, so a mode that moved the baseline would be
   `--accept` with a friendlier name (`D-M257x-49-3`). It admits a site only on a signature no repair can
   produce — the claim's **ledger row is a line this commit added**, *and* the site's file is one this
   commit **did not touch**. Proven on the real commit `D-M257x-48-12` records as blocked; the durable
   assertion is a **hermetic temp-repo fixture** built after the live-tree version was invalidated by
   this iteration's own repair (`D-M257x-49-9`).

3. **The twelve repaired by CLAIM, tree-wide**, then **the ninth clause-5 reading** — 7 seats, iter-41's
   instrument frozen on every knob, ground-truth clones unchanged from iter-48.

## The reading — 14 blockers, and the prediction was refuted in every term

Full adjudication in [`blocker-ledger.md`](blocker-ledger.md). Per seat: **A 1 · B 1 · C 2 · D 2 · E 3 ·
F 0 · G 5**. 14 raw → 14 unique → 14 held.

**Pre-registered before any report was read: 6 blockers, 2 induced / 4 pre-existing.**
**Actual: 14 blockers, 7 induced / 7 pre-existing.**

> ### The two new fences closed the gaps they were built for. The induced term went 2 → 7 anyway.
> Not because they failed — both were watched RED first, and the post-condition **caught this repair
> twice before the commit**. Because the seven induced findings partition into **paraphrase leak (3)**,
> **overshoot in new text (3)** and **wrong-mechanism-correctly-cited (1)** — and **not one is
> mechanically reachable**. A paraphrase shares no token run (the limit `D-M257x-48-4` pinned); an
> overshoot lives in prose that did not exist before the commit, so every diff-relative fence is silent
> by construction.

**TOK-02 step 2's premise — *"the 8-of-9 induced class cannot survive the commit"* — is now true of a
class that has stopped being the majority.** Mechanising the mechanical half did not lower the total; it
changed what the remainder is made of. That is the most valuable thing this iteration produced, and it is
a statement about the strategy, not about this repair.

## Deliverables

- Two fences, each watched RED before trust, each with an inversion mutant + a no-op control that must
  survive, each with its **reporting path deleted and the suite watched failing** (2 tests and 1 test
  respectively), and each with its **limits pinned by tests rather than described in prose**.
- The 12 repaired by claim, tree-wide: the ratchet **18 sites → 0**, `stack-core` **22F → 14F** (the 8
  traced ratchet failures cleared exactly as predicted; the remaining 14 are the pre-existing m220/m255
  batteries).
- `RETRACTION_MARKERS` widened by `falsified`/`refuted` — **with all three answer-key suites re-run** to
  prove no known blocker was hollowed out (`D-M257x-49-7`).
- The ninth reading, adjudicated, with the induced/pre-existing split and the **mechanism** behind it.

## Close — 2026-08-03

**Outcome:** ninth clause-5 reading returns **14** adjudicated blockers — **7 induced by this pass's
repair, 7 pre-existing**. The pre-registered 6 (2/4) is refuted in every term. Both named fences shipped
and closed their gaps; the induced term grew anyway, because its dominant classes are now **paraphrase and
overshoot**, which no diff-relative fence can reach.
**Type:** tik
**Status:** closed-fixed — all three planned steps landed (two fences + repair + reading). The reading is
a measurement and legitimately failed to close clause 5; `overview.md` pre-declared that this is not a
no-lift.
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this iter is a tik; the no-prog streak stands
at **2** — iters 48 and 49 — so the next no-prog tik fires one) — (3) re-scope: n (platform origin
`2adcf71` unchanged at open **and** close; trigger stays at occurrence 1 of 2) — (4) user-blocker: **y** —
(5) cap-reached: n — (6) protocol-stop: n — Outcome: **exit-4**
**Decisions:** D-M257x-49-1 .. D-M257x-49-9 (see [`decisions.md`](decisions.md))
**Side-deliverables:** none — the hermetic-fixture rebuild is a correction of this iteration's own test
defect, not a side discovery.
**Routes carried forward:**
- **The 14 adjudicated blockers**, unrepaired — `FIX-M257x-iter49-blocker-set`. Owner = next iter.
- `FENCE-M257x-iter50-paraphrase-leak` (Fate 3) — the measured dominant induced class. A paraphrase shares
  no token run, so a string method cannot reach it; the candidate instrument is a **claim-shape** match
  (identifier set + polarity) rather than a word-run match. **Pin the limit if it proves unbuildable** —
  the `D-M257x-48-4` discipline.
- **`CHECK-M257x-iter49-overshoot-has-no-instrument`** (Fate 3) — 3 of 7 induced findings are corrections
  that over-correct, in prose that did not exist before the commit. **No diff-relative fence can see
  this.** It needs either a reader or a different kind of check entirely; naming it is the deliverable,
  because believing a fence covers it is worse than knowing none does.
- Seat F's finding that **rosetta's own root `CLAUDE.md` is the stale side** of the `SkillPathSessionService`
  claim (`skillpath.md` proves 0 hits with a working positive control). Outside the 40-file partition.
- Seat C's caveat: `gh` is absent and `colony`/`proto`/`taxonomy` are not cloned, so all GitHub
  archive-status claims and those libraries' internals are **unverified, not passed** — same for Seat D
  and Seat F's §3 census.
- Seat A 9 · B 12 · C 5 · D 17 · E 9 · F 12 · G 13 minors (anchor drift, stale diagrams, one
  path-concatenation trap).

**Escalation.** `EXIT_REASON: user-blocker`, and the reason is specific rather than budgetary: this
iteration built exactly the two instruments the prior one named, watched both go RED, proved both, and the
induced term **rose from 2 to 7**. The measurement says the remaining induced defects are **not of a shape
a machine can catch** — which is a question about TOK-02's strategy, and it should land and be reviewable
before the next repair commits to the same method. The user has ruled twice that clause 5 stands as
written and that the work continues; **this escalation does not re-open that** — it reports that the
cheapest-instrument classification the strategy rests on has shifted under it.

**Lessons:**
1. **A fence can close its gap and not move the number.** Both fences work — asserted, not assumed. The
   induced total still grew, because fixing the reachable class promotes the unreachable one to majority.
   → the classification in §5 rule 21 must be re-run **after** each instrument lands, not once.
2. **A test coupled to a corpus state is invalidated by the repair it accompanies.** Five audit-mode
   tests passed against the live tree, then failed in the same session once the repair cleared the sites
   they depended on. **§8 rule 7 already says exactly this** — *"if you cannot state why an assertion
   will still be true after the defect is fixed, it belongs in the fixture"* — written by this milestone
   at iter-45, four iterations earlier, and violated here by the author of the fence it protects.
   → recorded as a **recurrence corollary on §8 rule 7**, which also generalizes it beyond fences to any
   test whose subject is a repairable state of this repository.
3. **The seventh consecutive occurrence** of *the author of a correction violating it while writing it* —
   three of this pass's seven induced findings are overshoots inside corrective text.
