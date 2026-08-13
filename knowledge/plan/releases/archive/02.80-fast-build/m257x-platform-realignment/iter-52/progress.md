**Type:** tik (under [`TOK-03`](../decisions.md#tok-03-repair-the-union-shrink-the-estimator-make-the-edits-smaller--2026-08-03))

# iter-52 — repair the union of 18, and read the repair before committing it

Executes TOK-03's pre-registered next-tik direction verbatim: **`FIX-M257x-iter50-union-set`**, moves 1
(repair the union), 3 (smaller edits), and 4 (two blind pre-commit readers). **No clause-5 reading taken** —
by design; iter-53 owns readings #11/#12, and a repairer who also reads is not blind to its own work.

## What landed

| | |
|---|---|
| union rows repaired | **18 of 18**, by claim, tree-wide |
| `claim_twin_guard` | **RED 31 hits / 19 sites → OK, 0** |
| pre-commit reader blockers found **in my own repair** | **5** |
| of those, fixed in-iter | **5** |
| ratchet | **passed on its own — no `--no-verify`** |
| rext answer-key suites | 30 tests, 1 failure — the **pre-existing** iter-49 stem collision; no known blocker suppressed |
| platform origin | `2adcf714` at open **and** close — re-scope trigger stays at **occurrence 1 of 2** |

## The result that matters: TOK-03 move 4 paid on its first outing

Two blind adversarial readers (seats H and J), barred from `knowledge/plan/**` and from each other, read
the repair diff **before** it was committed. **Both independently found the same worst defect, and it was
mine:**

> **I turned a CORRECT number into a WRONG one.** `grep -c 'OrganizationMixin{}'` returns **30** — but
> `user_resource.go:22` is **commented out** (`// OrganizationMixin{},  // We need to work on this`). The
> live mixin set is **29**, so `29 + Membership + Organization = ` **31** — the number the corpus already
> had. iter-49's audit booked *"31 is wrong, it's 32"*; **that finding was itself false**, and I repaired
> the corpus into error on its authority — in the same sentence where I wrote *"re-derive the SET, not the
> sum."*

Three consequences, and none of them is small:

1. **The union of 18 was not a clean answer key.** At least one of its rows was a **false finding**. A union
   improves *coverage*; it does nothing for *correctness*, and TOK-03 assumed the ledger was sound.
2. **iter-50's "sharpest single observation" was half wrong.** Its three seats who cleared the 31 as an
   audited zero reached the **right total by the wrong route**; iter-50 booked them as simply wrong. §5
   rule 24's worked example asserted **32** as settled fact and is corrected in this commit — the rule now
   records three generations of the identical defect, **the third committed by the author of the rule.**
3. **A `grep -c` over source counts code that does not compile.** Now written into rule 24.

The other four blockers were also induced by this repair, and three were found by both seats:

| finding | seats | what it was |
|---|---|---|
| *"ElevenLabs remains the **active default**"* | H + J | I **added** this while repairing the realtime-flag leak. The nil default is `gptrealtime` (`collections/jobsimulation.go:1594-1597`); no client renders an ElevenLabs call path at all. **The textbook induced defect: retract one half of a claim, promote the untested half.** |
| *"it **only** swaps the endpoint"* | H + J | The flag sets `agentEndpoint` **and** resets `agentName`, silently overriding a US session's `anthropos-agent-us` (`livekit.go:140-144`) |
| Studio-Room `anthropic` vs `openai` asymmetry | H + J | I demoted the `openai` arm as *"selected by no shipped config"* while leaving its `anthropic` twin counted as live **on identical evidence** |
| published counts `12 / 12` | J | Really **13 / 13**. `grep` here is a `.gitignore`-honouring wrapper that undercounts — so **my own derivation command did not reproduce my own numbers** |

## Re-estimated population — and how it was derived

**This is a derivation from the coverage arithmetic, not a reading.** Stated so it can be checked and
refuted.

**Term 1 — pre-existing residual.** `N̂ ≈ 23` (Chapman, and a **floor**). The union repaired 18 rows, but one
was a false finding, so **17 real defects** were removed:

> `23 − 17 =` **~6 pre-existing remaining**, and still a floor.

**Term 2 — induced by THIS repair, measured the same way.** The two diff readers are themselves a
capture–recapture pair over the repair's own defects — the first time this milestone has had one:

| | seat H | seat J |
|---|---|---|
| findings | **7** | **12** |
| matched (both) | colspan → **6** | |

`Chapman = ((7+1)(12+1) / (6+1)) − 1 = (104/7) − 1 ≈ ` **13.9 → ~14 defects in the repair diff.** The two
readers *named* 13 of them (union `7 + 12 − 6`); **6 were repaired in-iter**, the rest routed.

> **Total `N̂` ≈ 6 + 8 ≈ 14, and it is a FLOOR.** Down from ~23 — but **far less than "18 repaired" implies,
> because the repair induced roughly as much as the extra coverage bought.** That is TOK-03 move 3's thesis
> measured directly, and it does **not** support the strategy's optimism.

**Against TOK-03's pre-registrations for iter-53** (`N̂ < 12`, induced term `< 4`): the derived `N̂ ≈ 14` is
**above** the first, and the induced term is **~8**, twice the second. Both were pre-registered to be
refutable; on the derivation they are heading toward refutation, and iter-53's paired reading is what
settles it. **The reader recall itself is now measurable** — 6/7 and 6/12 — which is a second, cheaper
instrument than a full 7-seat pass.

## Close — 2026-08-03

**Outcome:** the union of 18 repaired by claim tree-wide; `claim_twin_guard` **31 → 0**; **five blockers in
my own repair caught by two blind pre-commit readers and fixed before the commit**, including a correct
number I had made wrong on a false ledger finding; `N̂` re-derived to **~14 (floor)** from ~23, with the
induced term measured for the first time at **~8**.
**Type:** tik
**Status:** **closed-fixed** — the iter's planned scope was *repair the 18 under moves 3 and 4, then
re-estimate*. All of it landed: 18 repaired, both readers run **pre-commit** as specified, their findings
triaged and fixed, the ratchet passed **without a bypass**, and the estimate re-derived with its method
stated. The word-budget prediction was refuted, which is a recorded result of the planned work, not a
failure to do it.
**Gate:** **NOT MET** — clause 5 requires a reading that returns zero, and this iter deliberately took no
reading. Gate stands at **4 of 5**.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this was a tik; the streak reset at iter-51) —
(3) re-scope: n (platform origin `2adcf714` at open **and** close; occurrence 1 of 2) — (4) user-blocker: n —
(5) cap-reached: n (1 tik this session) — (6) protocol-stop: n — Outcome: **continue**
**Decisions:** [`D-M257x-52-1`](decisions.md) … `-52-4`.
**Side-deliverables:** `corpus/ops/platform-alignment.md` §5 rule 24 corrected (its worked example asserted
the refuted **32** as settled fact) — protocol-doc update shipped in the same commit as the lesson, per the
skill's protocol-evolution rule.
**Routes carried forward:**
- **`FENCE-M257x-iter52-stem-collision`** — require a *distinguishing* fragment before the fence may fire.
  **A minimal repair keeps the ledger's quoted stem, so the fence fires on CORRECTED text: it rewards the
  rewrite and punishes the small edit**, which is exactly backwards for TOK-03 move 3. Needs its own
  RED-watch, so not done inside a repair pass.
- **`FIX-M257x-iter52-mirror-pair-leak`** (seat J, BLOCKER, **pre-existing**) — *"seed the co-written PAIR"*
  still stands at **40 sites in 12 files** (`seeding-spec.md`, `verification.md` — including a **gate floor**
  of *"≥ 40 candidate `local_jobsimulation_sessions`"* against a **dropped table** — `stories-spec.md`,
  `CLAUDE.md:338-339`, `.claude/skills/db-query/SKILL.md`). The single largest live leak in the tree.
- **`FIX-M257x-iter52-alignment-doc-s6`** (seat J) — `platform-alignment.md:833`'s §6 *current-state* row
  says the `jobsimulation` schema is *"created by rext itself"*; rext HEAD no longer does.
- Seat findings not yet landed: the migration's **DELETEs** (`:30-34`, `:49`) omitted from the mirror-drop
  summary; studio-room's third outbound host (`export.py:48-58`); the two EU endpoints
  (`azure-eu`/`azure-eu-fr`) and `anthropos-agent-chain` missing from `ai_architecture.md:180`;
  `security_compliance.md` stating *3* and *4* `Policy()` files eighteen lines apart.
- **`CHECK-M257x-iter52-second-ai-manager`** (seat H, BLOCKER, **unverified by me** — do not act on it
  before re-deriving): `app/internal/skillerai/ai.go` may be a **second** AI dispatcher with its own
  `flag_use_azure_us` and its own 429→direct-US-OpenAI retry, which would make the EU-exit enumeration
  incomplete *again*. I did not confirm it and **deliberately did not repair on one seat's word** — that is
  how iter-49's `32` entered the corpus.
- Still open from earlier: `FENCE-M257x-iter50-consecutive-audit-mode` · `CHECK-M257x-iter35-seeder-writes-one-instant` ·
  RF-13 · RF-2/3/7–12 · harden residue · `CHECK-M257x-iter38-ai-act-classification` (needs an owner outside
  this milestone) · the root `CLAUDE.md` (now with a named leak from seat J).

**Lessons:**
1. **Repairing the union does not help if the union contains a false finding.** TOK-03 optimised coverage
   and took the ledger's *correctness* for granted. The union is a **superset of findings, not of truths** —
   and this pass proved a repair can be *worse than doing nothing* when it acts on a bad adjudication.
   **Before repairing a row, re-derive its verdict**; §5 rule 19's *"a repair must not adjudicate"* has a
   silent premise — that the adjudication was right.
2. **Two blind readers on the diff, pre-commit, is the highest-yield instrument this milestone has built.**
   Cost: one diff read each. Yield: 5 blockers, 3 of them found by both — in text that had just been written
   by an author who believed it. **It should have existed since iter-33.**
3. **The word budget was the wrong metric and the readers proved it.** Net was **+1266**, not ≤ 0. Deletion
   was available for **2 of 18**; the rest were wrong values inside load-bearing sentences, where deleting
   removes the reader's answer with the error. And §5 rules 17/24 *demand* added words. **Bound NEW
   ASSERTIONS, not characters** — the next tok should re-cut move 3 that way.
4. **A fence can push the repairer toward the riskier repair shape.** The stem collision is not a nuisance;
   it is an instrument actively selecting against the discipline the strategy is trying to install.
