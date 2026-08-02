---
iter: 45
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-02
---

# iter-45 — `FENCE-M257x-iter45`: the three mechanical fences

**Active strategy reference:** [`TOK-02: fence the prose the way the anchors are fenced`](../decisions.md#tok-02-fence-the-prose-the-way-the-anchors-are-fenced--2026-08-02) — **step 3** of its five ordered
steps, verbatim: *"Add the two small mechanical fences the classification names, closing 5 more of the 18:
a markdown structure lint … and a symbol-aware anchor check"*, together with the derived-value fence
TOK-02's own classification table routes `#10`/`#11` to.

## Step 0 — re-survey before targeting

Re-run at open: platform origin HEAD `2adcf71`, **unchanged** (re-scope trigger stays at occurrence 1 of
2). Gate **4 of 5**; clause 5 at **18**, unrepaired — iter-43 and iter-44 both closed by construction
without repairing a single blocker, so the answer key is intact. `tests/fixtures/claim_twin/red/`
holds 18 files. TOK-02's named next target for this slot is **unchanged and still meaningful**: no
substitution.

## Cluster / target identified

iter-42 classified iter-41's 18 blockers by the **cheapest instrument that could have caught each**.
iter-43 built the instrument for the 13-strong self-contradiction class. This iteration builds the
other two instruments that classification named, plus the value fence it routed `#10`/`#11` to:

| blocker | what is wrong | instrument |
|---|---|---|
| `#6` | a retraction blockquote spliced into a bullet list; the list resumes afterwards drawing the legal consequence from the classification just retracted | markdown structure lint |
| `#13` | `external_services.md:788` cites `:447 above` — `:447` is a table **header row** | symbol-aware anchor check |
| `#16` | `messenger.md:110` cites `assignments.go:815` — `:815` is `}))`, 13 lines into the wrong function | symbol-aware anchor check |
| `#10` | `sentinel.md:12` states "Go 1.25"; `go.mod` says `1.26.0` | derived-value fence |
| `#11` | `sentinel.md:22` states "256 CPU / **256 MB**"; `locals.tf` says `service_memory = 128` | derived-value fence |

`#17` is **not** in this set and that is a measured exclusion, not a concession — see the escalation
conditions below.

## Hypothesis

Five of the eighteen are **mechanical damage or a scalar mismatch**, neither of which requires reading a
sentence. A fence per class, each measured tree-wide before adoption and each watched going RED on the
live answer key while it still exists, converts those five from "a human must notice" into "a commit
cannot carry it" — the same conversion iter-44 performed for the self-contradiction class.

## Expected lift

**Zero on the primary metric, by construction.** Clause 5 stays at 18: this iteration repairs nothing.
That is TOK-02 step 3's charter and iter-43's own lesson — *today's 18-defect corpus is the fence's only
test fixture with a known answer key, and it is perishable*. The deliverable is instrument reach:
**5 of the 18 detected**, each with a RED watch that survives step 4's repair.

## Phase plan

1. Three guards, each **measured tree-wide before adoption**, with the false-positive rate of every
   rejected draft recorded next to the rule that replaced it (§5 rule 2).
2. Enrol all three in iter-44's commit-time ratchet (`FENCE_KIND` + a `postcondition_sites` provider).
3. A behaviour suite: each blocker found **live**, each measured false positive proven silent.
4. A **mutation battery** — §8 rule 5 in full: `py_compile` per mutant, a declared-GREEN no-op control
   that must survive, inverted mutants alongside deletions, and **a mutant per reporting path**.
5. **Capture the perishable fixture** for all five sites, so the RED watch outlives step 4.
6. Protocol-evolution obligation: `platform-alignment.md` §8 + `stack-core/README.md` rows.

## Escalation conditions

- A rule that cannot be made silent on correct prose is **dropped**, not narrowed until the answer key
  fits — narrowing a window until the known blocker fires and its neighbours do not is Trap A, tuning a
  fence to its own test. `#17` is already routed out on exactly this ground.
- A fence whose live RED cannot be reproduced from a captured fixture does not ship: it would become
  un-testable the moment step 4 runs.

## Acceptable close-no-lift outcomes

The metric does not move and must not. This closes `closed-fixed` on instrument reach, or
`closed-fixed-partial` if a rule is dropped for false positives and its blocker is routed to step 4's
hand repair with a named handler.
