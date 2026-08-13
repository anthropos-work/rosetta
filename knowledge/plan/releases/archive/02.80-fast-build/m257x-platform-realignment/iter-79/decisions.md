# iter-79 — decisions

## `D-M257x-79-1` — 145 → 3. A rule discovered twice and implemented twice is a rule not implemented.

`CHECK-M257x-iter77-cross-repo-pin` was routed as *"145 pin-exempted blocks name a sha that does not
resolve in the platform clone"*, and iter-78 upgraded it to *"at least two of them date a platform
claim, and both were false."* Sized here before generalising, because **every routed count in this
milestone has shrunk when someone finally derived it**:

| | |
|---|---|
| pin-exempted blocks on the live corpus | **390** |
| exempted by a sha that does not resolve in the platform repo | **145** across **28** distinct shas |
| …of those, gating a **platform-file** assertion | **3** |
| …still false at this iter's open | **0** — iter-77 and iter-78 repaired 2, the third is correct |

**`145 → 3`. The sixth routed count in this milestone to collapse on derivation** (64→5, 23→1, 21→0,
92→0, 4→3, 145→3). The other 142 are legitimate `app`-repo citations about `app` files — which is
precisely why this ships as **a helper the platform-file assertions call**, not as a widening of
`_pin_exempts` for every guard. A rule that punished all 145 would be §4 Trap A with 142 false
positives.

**The finding that justifies the iter is not the count, it is the duplication.** G9 tripped over this
at iter-77 (`roadrunner.md:14`, dated by `87d8d44`, a **roadrunner** commit) and G10 tripped over it
independently at iter-78 (`external_services.md:296`, dated by `b948604`, an **app** commit, while
claiming something about `0dab54d`). Each fixed it *inside itself*, and G2/G4/G5 — which assert
claims about `repos.yml` and `docker-compose.yml` and are therefore in exactly the same class — still
took any sha in the block as a date.

Shipped: `ref_resolves_in` + `pin_dates_a_platform_claim`, one helper, called by the assertions whose
subject is a platform file. **Reach up (one further migration claim graded, `2 → 1` skipped),
findings unchanged at 0** — which is the pre-registered outcome and the honest one: *reach was the
deliverable, not a repair count.*

---

## `D-M257x-79-2` — a clone that cannot answer must never be read as answering "no".

The first cut asked `ref_resolves_in()` unconditionally. On a platform directory that is **not a git
checkout**, every sha is unresolvable — so *no pin dates anything*, and the guard goes **RED across
the entire corpus at once**.

That is the same defect as iter-77's current-state vocabulary, with the sign flipped: **blind there,
hostile here**, and both from reading "cannot answer" as a substantive answer. An existing
prose-window test caught it **within one run** of its introduction — which is the argument for the
milestone's habit of keeping old fixtures that assert *silence*.

The fix is the discipline already written down: detect whether the question is answerable at all
(`can_resolve_refs`), and where it is not, **degrade to the previous behaviour** rather than to a
verdict. `repos_yml_history` says `UNMEASURED`; this returns the un-narrowed exemption.

> **Generalisation, and it now has two independent instances in three iterations:** every derived
> discriminator has three outcomes, not two — *yes*, *no*, and **cannot tell**. Collapsing the third
> into either of the first two produces a guard that is confidently wrong, and which of the two
> failures you get is an accident of which way the boolean fell.

---

## `D-M257x-79-3` — the read-at-ref assertions had never inherited the currency clause.

`D-M257x-63-1` has three clauses; the third is that a block **asserting currency** (*"current, not
stale"*, *"currently"*) is never exempted by its pin, because there the pin is being cited as
evidence of currency rather than as the date of a measurement.

G9 and G10 do not call `_pin_exempts` — they **resolve a ref and read the artifact there**, which is
strictly better — and in doing so they had silently opted out of all three clauses. Two were
re-implemented locally; the currency clause was not. A block reading *"currently, at `2adcf71`,
`repos.yml:17-19` lists jobsimulation"* would have been read at `2adcf71`, found true, and passed —
while asserting it about **now**.

Wired in, with a regression test. No live site exercises it today, which is exactly why it needed a
test rather than a measurement.

---

## Routed forward

- **`CHECK-M257x-iter77-cross-repo-pin` — CLOSED.** Sized (145 → 3), generalised into one shared
  helper, and the 142 non-instances explained rather than carried. What remains is not this class:
  G1's noun-phrase construct and G7/G8 do not assert platform-file claims and are untouched.
- **`CHECK-M257x-iter79-three-valued-discriminators`** — `D-M257x-79-2`'s generalisation, as a
  standing audit: every derived discriminator in this guard family should be checked for a
  *cannot-tell* branch. Two of the three found so far were missing one. Not opened here; it is an
  audit, not a fix, and it belongs to a harden pass.
- Unchanged: `CHECK-M257x-iter78-running-vs-declared` ·
  `CHECK-M257x-iter77-narration-vs-documentation` · `CHECK-M257x-iter77-zsh-modifier` ·
  `CHECK-M257x-iter77-developer-dir` · `CHECK-M257x-iter76-seat-ref-discipline` ·
  `CHECK-M257x-iter70-studio-room-lines` · `RF-M257x-iter71-run-returns-a-tuple` ·
  `FIX-M257x-iter53-union-set` (**PENDING USER DECISION**) · `FIX-M257x-iter56-assignment-flake`
  (**NOT DECIDED**) · `CHECK-M257x-iter38-ai-act-classification` (owner outside this milestone) ·
  RF-2/3/7–13.
