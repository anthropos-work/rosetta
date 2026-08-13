# iter-97 pre-registration — readings #19 + #20, at platform `0c91421`, corpus `00be1ac`

**Written and committed BEFORE any seat was launched.** iter-95 graded 6 of 6 and recorded it as a
*warning, not a win*: prediction 1's band `[0,12]` was nearly unfalsifiable and prediction 4 landed
within a tenth of a point. iter-76 and iter-53 each graded **2 of 5** and learned more from missing.
Every band below is narrowed until it can fail, and three of the seven are net-new risky claims rather
than range guesses.

## The seven

| # | prediction | band | falsifier | why it is risky |
|---|---|---|---|---|
| **1** | per-reading in-scope upheld BLOCKER count | **[2, 7]**, *both* readings | either reading returns **0**, **1**, or **≥8** | iter-95 returned 10 and 7. The band is **6 wide** where iter-95's was 13, and it is two-sided: it asserts the repair moved the number **and** that it did not reach zero. Prior bands could not fail downward at all |
| **2** | union **N** (in-scope, BLOCKER, deduped) | **[3, 9]** | N ≤ 2 or N ≥ 10 | iter-95's N was 13. A 7-wide band around a halving |
| **3** | **none of the 12 repaired predicates recurs** as an upheld in-scope BLOCKER at any site | exactly **0** | any one recurs | The sharp claim. `claim_twin_guard` is GREEN over the 14 refuted forms, so a recurrence means either the fence's matcher is too narrow or a seat found a *paraphrase* the fence cannot see — iter-93's exact defect class. **This one has teeth against my own work** |
| **4** | upheld rate | **[86 %, 96 %]** | outside | Prior three adjudications: 92.1 / 93.0 / 92.7. A two-sided 10-point band replaces iter-95's one-sided `≥ 80 %`, which the instrument had cleared by 12 points three times running |
| **5** | per-pass recall (Chapman, graded set) | **[30 %, 62 %]** | outside | iter-95: 60 % / 42 %. A smaller true N makes recall estimates *less* stable, so this band is genuinely at risk — and it can fail upward, which iter-95's `< 60 %` could not |
| **6** | the **dominant class shifts** off "platform-derived by the `0dab54d → 0c91421` move" | **< 50 %** of upheld in-scope blockers trace to that move | ≥ 50 % | iter-95: **7 of 13 = 54 %**. The repair targeted exactly that class, so if the share does not fall the repair did not reach the class it was aimed at |
| **7** | the repair **induces** findings | **≥ 2** upheld in-scope blockers in files this repair touched | 0 or 1 | iter-41 measured **9 of 18** findings manufactured by the preceding repair, and 8 of those were one mechanical class. Predicting my own repair is clean would be the safe prediction; this predicts it is not, and names a floor |

## What is NOT predicted, deliberately

- **No prediction that clause 5 is met.** That is the measurement, not a hypothesis.
- **No prediction about MINORS or out-of-scope findings.** They are reported, and they are not the clause.

## The instrument, unchanged

`instrument/briefing-iter76-AS-RUN.md`, sha256 `3858ec53…6eb0`, **one commit ever** (`012edd2`) —
re-checked at this open, byte-identical, and copied to the seats with the sha re-verified after copying.
7 seats per reading (A–F full-read partition + adversarial diff seat G) × 2 readings. Partition method
unchanged: files sorted by line count **descending**, snake-dealt A→F then F→A over
`corpus/architecture/*.md` + `corpus/services/*.md` — **40 files / 10,210 lines** (was 10,108; the
repair grew the scope by 102 lines, so the method re-deals a different hand, as it has every reading).

Seat **G**'s base is **`b7e6642`** — iter-95's close, the last graded reading — so G's scope is exactly
the iter-96 repair: **23 files, +230 −79**. G is therefore an adversarial read of this run's own work.
