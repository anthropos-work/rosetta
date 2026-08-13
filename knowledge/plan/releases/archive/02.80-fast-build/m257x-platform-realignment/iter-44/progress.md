**Type:** tik (`iter_shape: tooling`) — protocol: [`corpus/ops/platform-alignment.md`](../../../../../corpus/ops/platform-alignment.md)

# iter-44 — the prose fence becomes a commit post-condition

**Active strategy:** `TOK-02` — *fence the prose the way the anchors are fenced*. **Step 2 of five, and
step 2 only.**

## What ran

| phase | outcome |
|---|---|
| **A** re-survey | fixture intact (18 red / 18 green); live residual **18**, unchanged; platform origin `2adcf71` re-fetched — unchanged |
| **B** registry | `FENCE_KIND` declared on all **7** `stack-core` guards (1 `postcondition`, 6 `standalone`), read **statically** by `ast` |
| **C** runner | `stack-core/repair_postcondition.py` — derived registry, line-number-free site identity, ratchet baseline, `--accept` / `--install-hook` / `--staged` |
| **D** watch it go RED | an **induced** site in a repaired tree → RED naming `file:line` + the verdict it contradicts; the repaired tree itself is silent first |
| **E** tests + battery | **25** unit tests + a **12-mutant** battery; `stack-core` **450 / 14F** vs a 420 / 14F baseline (+30 tests, 0 new failures) |
| **F** protocol | §8 gains the *"run the fence at the COMMIT"* section; `stack-core/README.md` gains 2 rows |

## The result

**The induced-defect class is now a commit-time RED rather than a next-audit finding.** The premise is
stated so it can be refuted, and it was: an induced self-contradiction **is** a new `(file, claim)` pair
— a repair wrote the correct sentence at one site while an adjudicated form of the same claim stayed
published at another. Eight of iter-41's nine repair-induced blockers have exactly that shape, so a
subset-of-baseline assertion reaches them all.

**The RED that matters is new.** The claim-twin suite already pins *"this corpus has 18 known defects"*.
Nothing pinned *"a repaired tree just grew a nineteenth"* — no fixture in the repo contained one. So the
anti-fixture is **assembled from the answer key**, not written: the 18 GREEN twins plus one captured RED
file dropped back in (`D-M257x-44-8`). The repaired tree is asserted silent first, or the RED would prove
nothing.

**Battery: 12 mutants, every one matching its declared verdict** — 1 declared-GREEN no-op that
**survived**, 11 kills with **11 distinct failure signatures**, of which 5 are inversions. Four exist
only to **silence a reporting path** the module claims to have, which is the direct lesson of harden
passes 7–9: two of the claim-twin fence's own honesty mechanisms did not exist, one of them deleting
clean with 15/15 still green. Baseline GREEN before and after; every mutant `py_compile`d before its run.

**Two vehicles, and the weaker one is labelled** (`D-M257x-44-5`). The suite is load-bearing — it grades
the live tree in every clone. The `--install-hook` pre-commit is a latency optimisation, **per-clone and
unversioned**, which is the same shape as iter-01's git-ignored `rext.tag` that *"never appears in a diff
and drifts unseen"* — so it is disclosed in the docstring, the protocol section and the installer's own
output. The hook is installed on this box and ran on this iteration's own corpus commit.

**Two defects were found in this iteration's own tests while writing them**, both worth recording because
both are the shape the milestone keeps meeting:

- `test_20` asserted the hook's exec bit **after** the temp dir was deleted — a check that fails for a
  reason unrelated to its subject. It failed loudly here; the dangerous inversion (asserting *presence*
  after deletion) would have passed forever.
- `test_06` built its baseline with an **absent** fence entry rather than an **empty** one, so the
  induced site graded as a *registration* and the test read 0 induced defects. It is exactly the
  distinction `D-M257x-44-3` draws, and it caught its own author within an hour of the rule being
  written — the fourth consecutive occurrence of *"the author of a newly-written rule violated it while
  writing it."*

## Nothing was repaired

`D-M257x-42-3` holds. The residual is **18**, by construction. The only corpus edit is new text (the §8
subsection), and the post-condition was run over it: 18 sites, byte-identical set, exit 0.

## Close — 2026-08-02

**Outcome:** the claim-twin fence now runs as a **commit post-condition with a monotone baseline** —
watched going RED on an induced defect in a repaired tree, with a derived fence registry, a surviving
no-op control and 11 distinct kills. The induced-defect term TOK-02 targets is now unrepresentable in a
commit rather than merely visible in the next pass.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5. Clause 5 stays at **18** by construction: this iteration repairs nothing.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this was a tik) — (3) re-scope: n (platform origin `2adcf71` re-fetched at open and close, unchanged; trigger stays at occurrence 1 of 2) — (4) user-blocker: n — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n — **Outcome: continue**
**Decisions:** `D-M257x-44-1` … `D-M257x-44-8`
**Side-deliverables:** none — the §8 section and the `stack-core/README.md` rows are this iter's own
protocol-evolution obligation, not unrelated fixes.
**Routes carried forward:** `CHECK-M257x-iter43-symbol-anchor` · `CHECK-M257x-iter43-markdown-lint` ·
`CHECK-M257x-iter43-value-fence` (all three are TOK-02 step 3 → **iter-45**) ·
`FIX-M257x-iter43-coverage-protocol-livepath` (TOK-02 step 4)

**Lessons:**

- **A fence that only runs at audit time cannot reduce a repair-induced defect count.** The defect is
  committed, and it is one of the findings being counted. Where a check runs is a design parameter with
  a measurable consequence, not an implementation detail.
- **"Absent from the baseline" and "present but empty" must be different states, or a ratchet cannot
  tell a first measurement from a regression.** Collapsing them silently converts a new fence's entire
  unwatched site set into either 18 false alarms or one silent pass, depending which way you collapse it.
- **Assemble the anti-fixture from the answer key.** A hand-written "induced defect" encodes a guess
  about the class; a GREEN twin set plus one captured RED file *is* the class. And assert the repaired
  half is silent first, or a fence that reddens on everything passes.
- **Give every reporting path a mutant.** Four of this battery's eleven kills do nothing but silence a
  line of output. Two such paths on the *previous* fence turned out never to have existed at all.
