**Type:** tik — under `TOK-08`, corpus half of the user redirect.

## Phase 1 — sealed

Predictions `P-228-1..3` sealed. Subject: 13 clones at origin tips + the 3 `origin/main is now` sites.

## Phase 2 — the census, and the bug in its first pass

First pass measured each clone's **local HEAD** against the corpus's sha vocabulary and reported every
one of the 13 as cited. **That was the wrong operand** — for a clone that is behind, local HEAD is not
the origin tip, and the question was about the tip. Re-run against `origin/main`. Disclosed rather than
silently corrected, because the first table would have read as a clean bill of health.

| repo | origin tip | behind | tip cited in corpus |
|---|---|---|---|
| **`app`** | **`3eaadae6`** | **28** | 8 |
| `next-web-app` | `19423a1f` | **12** | **0** |
| `ant-academy` | `c885dab2` | **9** | **0** |
| `rosetta-extensions` | `3667d5b7` | 159 | 0 — *pinned consumption clone, out of scope* |
| `cms` `graphql-wundergraph` `jobsimulation` `messenger` `platform` `roadrunner` `sentinel` `storage` `studio-desk` | — | **0** | 4–135 |

**`P-228-1` REFUTED, and it is the finding.** `ad9f3c49` is **not** `app`'s `origin/main`; `3eaadae6` is,
and `ad9f3c49` sits **28 commits** behind it. `CLAUDE.md` asserted the moving label in **2 lines / 3
occurrences** — in the same paragraph that warns *"Cite the sha, never the moving label."*

**And the corpus already knew.** `corpus/ops/observability.md` states *"`origin/main` was `3eaadae6` on
2026-08-07"*, and `shared_libraries.md` measures `app/go.mod` at `3eaadae6`. **`CLAUDE.md` was not behind
the platform — it was behind the corpus**, and nothing compares one to the other.

**`P-228-2` HOLDS** — 3 of 13 tips cited zero times, of which 2 are real (`next-web-app`, `ant-academy`);
`rosetta-extensions` is the escalation case this iter's `overview.md` named in advance (a tag-pinned
consumption clone, not a drift subject).

**`P-228-3` HOLDS.** `clone_drift_guard` returns **OK / exit 0** throughout, with `app` 28 commits behind
— because D1 grades the clone's **local HEAD**, and `ad9f3c49` is cited 174 times, so it passes
comfortably. Its REACH line is accurate and its reach is narrow: *"no CITED repo advanced past everything
the corpus knows"* is not *"the corpus is current."*

**iter-222 and iter-224 fetched these clones and advanced only some.** iter-224 fast-forwarded the four
archived repos; the three **active** ones — `app`, `next-web-app`, `ant-academy` — were fetched and left
behind. That is not an oversight to fix casually: advancing `app` changes what a demo **builds**, which is
exactly the decision `ROUTE-M257x-222-pin-advance-needs-a-reproof` gates behind clause 1's cold cycles.
**Recorded, not advanced.**

## Phase 3 — repair, and the fence catching the repair

Both `CLAUDE.md` sites corrected to name the sha and drop the moving label, line-count neutral (`2↔2`).

**`anchor_construct_guard` went RED on the first attempt, and it was right.** The first replacement text
cited `observability.md:14` and `shared_libraries.md:6` — introducing a second **document name** into a
block that already carries a bare `:504` continuation pin. The resolver re-bound `:504` to
`shared_libraries.md` and reported `shared_libraries.md:504` out of range.

**This is the precise trap `observability.md:13-22` documents** — *"a block that names two refs is
ambiguous to the citation resolver, which then falls back and grades every anchor in the block against a
file the block did not mean"* — and it was walked into **while citing that very warning.** Fixed by
naming the two docs in prose without line pins.

## Close — 2026-08-09

**Outcome:** `CLAUDE.md` told every session that `app`'s `origin/main` is `ad9f3c49`; it is `3eaadae6`,
28 commits ahead — and two other corpus docs already said so. The stack's `app`, `next-web-app` and
`ant-academy` clones are 28 / 12 / 9 commits behind and were deliberately **not** advanced.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: y — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: exit-5
**Decisions:** `D-M257x-228-1` (the active clones are recorded behind, not advanced),
`D-M257x-228-2` (the census's own wrong-operand first pass is published).
**No `N`/`P` movement is claimed** — this iter took no graded reading.

**Predictions, graded:**

| id | prediction | result |
|---|---|---|
| `P-228-1` | `ad9f3c49` is still `app`'s origin/main | **REFUTED — it is `3eaadae6`; `ad9f3c49` is 28 behind** |
| `P-228-2` | ≥ 2 of 13 tips cited 0 times | **HELD — 3 (2 real + 1 out-of-scope)** |
| `P-228-3` | `clone_drift_guard` returns OK regardless | **HELD — OK / exit 0 with `app` 28 behind** |

**Suite state at close** — `guard_family` with `--platform stack-demo/platform`: **24 GREEN · 0 RED · 5
not-run** after repair (**22 GREEN · 2 RED** on the first attempt, both this iter's own, both repaired).
Not a whole-family green — the 5 not-run are commit/ledger-scoped members with no input supplied. No
pytest section run; this iter changed no rext code.

**Routes carried forward:**
- `ROUTE-M257x-228-active-clones-behind` → `app` **28**, `next-web-app` **12**, `ant-academy` **9**.
  **Merged into `ROUTE-M257x-222-pin-advance-needs-a-reproof`**, which already gates exactly this decision
  behind clause 1's three cold cycles. Now carries measured numbers instead of an open question.
- `ROUTE-M257x-228-corpus-disagrees-with-itself-about-refs` → nothing compares `CLAUDE.md`'s ref claims to
  the corpus's. Decidable, and a fence candidate: two docs naming a different `origin/main` for the same
  repo. Not built here (redirect ranks corpus above instruments; third line).
- All prior routes → open, unchanged.

**Lessons:**
1. **The most-cited fact in the corpus was stale, and the corpus already contained its correction.**
   `app` is cited 174 times; `CLAUDE.md`'s statement of its current ref disagreed with two other corpus
   docs. **Internal disagreement is cheaper to detect than external drift** — both operands are in this
   repo — and nothing looks for it.
2. **`clone_drift_guard` grades the operand it can reach, not the one the question is about.** It reads
   local HEAD; the question is about the tip. iter-224 showed staleness *satisfies* it; this iter shows
   the same green with the flagship repo 28 commits behind.
3. **A census is only as good as its operand.** This iter's own first pass measured local HEAD and would
   have published *"all 13 tips cited"* — a clean bill of health derived from the wrong column.
4. **The corpus's warnings are load-bearing and easy to walk into while quoting them.** Adding a document
   name to a block containing a bare `:NN` continuation pin re-binds that pin. Documented in
   `observability.md`; violated here in the act of citing it. **When editing a block, check what bare
   `:NN` pins it already carries before naming another file in it.**
