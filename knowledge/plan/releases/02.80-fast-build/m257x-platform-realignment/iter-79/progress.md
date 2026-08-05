# iter-79 — progress

**Type:** tik, under `TOK-05`. Planned deliverable: **size `CHECK-M257x-iter77-cross-repo-pin`, then
apply its rule once in the shared helper instead of a third local copy.**

---

## Phase A — size the class before generalising

Registered as a prediction in `overview.md` before any code was written: *fewer than 10, and 0 new
findings.* Measured:

| | |
|---|---|
| pin-exempted blocks on the live corpus | **390** |
| exempted by a sha that does not resolve in the platform repo | **145**, across **28** distinct shas |
| …gating a **platform-file** assertion | **3** |
| …still false at this open | **0** (iter-77 and iter-78 repaired 2; the third is correct) |

**`145 → 3`** — the **sixth** routed count in this milestone to collapse on derivation (64→5, 23→1,
21→0, 92→0, 4→3, 145→3). Both predictions hold.

The 142 non-instances are the reason the rule is **scoped, not global**: this corpus cites other
repositories' shas correctly and constantly — `app` commits dating `app` files, `roadrunner`
commits dating `roadrunner` terraform. Punishing all of them would be §4 Trap A with 142 false
positives.

## Phase B — the rule, implemented once

**The finding that justifies this iter is not the count, it is the duplication.** G9 tripped over
this mechanism at iter-77 and G10 tripped over it *independently* at iter-78, and each fixed it
inside itself. G2, G4 and G5 assert claims about `repos.yml` and `docker-compose.yml` — the same
class — and still took any sha in the block as a date.

Shipped: `ref_resolves_in()` and `pin_dates_a_platform_claim()`, called by the assertions whose
subject is a platform file. A sha is a sha, so `_REF_PINNED` cannot tell `87d8d44` (a **roadrunner**
commit) from a platform one; the repository can.

## Phase C — measure

```
before: ref-pinned and therefore skipped: 2/2/2
after : ref-pinned and therefore skipped: 2/1/2
platform_predicate_guard: OK   (0 findings, unchanged)
```

One further migration claim graded, no new findings. **That is the pre-registered outcome and the
honest one — reach was the deliverable, not a repair count.**

## Phase D — and the guard caught the generalisation inverting

The first cut called `ref_resolves_in()` unconditionally. **An existing test failed within one run**:
`test_prose_keeps_its_wrapped_window_and_is_not_narrowed_to_a_column`, whose fixture platform is a
plain directory rather than a git checkout. There, *every* sha is unresolvable — so no pin dates
anything and the guard goes **RED across the whole corpus at once.**

That is iter-77's defect with the sign flipped: **blind there, hostile here**, both from reading
*"cannot answer"* as a substantive answer. Fixed at the cause with `can_resolve_refs()`, degrading to
the previous behaviour rather than to a verdict — the same discipline `repos_yml_history` uses when
it reports `UNMEASURED`.

**The old fixture that caught it asserts silence, not a finding.** It is worth keeping tests whose
whole content is *"and nothing else fires"*.

### Controls (§8 rule 5)

- **No-op that SURVIVES** — a sha that *does* resolve in the platform repo exempts exactly as
  before, so the rule is not a blanket punishment of pins.
- **The date-pin control** — `at 2026-07-31` is a legal pin containing no sha; "no sha" must not
  read as "unresolvable", or every dated measurement in the corpus silently loses its exemption.
- **The degradation control** — a non-git platform directory returns the *un-narrowed* exemption,
  and the same foreign sha is refused where the repo *can* answer. Both halves asserted in one test.
- **INVERTED mutant** — a platform claim dated by `87d8d44` is graded and goes RED.

## Phase E — one more clause, found while wiring

G9 and G10 never called `_pin_exempts` — they resolve a ref and read the artifact there, which is
strictly better — and so had silently opted out of all three `D-M257x-63-1` clauses. Two were
re-implemented locally; the **currency** clause was not. A block reading *"currently, at `2adcf71`,
`repos.yml:17-19` lists jobsimulation"* would have been read at `2adcf71`, found true, and passed
while asserting it about **now**. Wired in with a regression test; no live site exercises it today,
which is exactly why it needed a test rather than a measurement.

## Close — 2026-08-05

**Outcome:** `CHECK-M257x-iter77-cross-repo-pin` **sized and closed** — 390 pin-exempted blocks, 145
exempted by a foreign sha, and **3** of those gating a platform-file assertion: **145 → 3**, the
sixth routed count in this milestone to collapse on derivation, with both pre-registered predictions
holding (*fewer than 10*, *0 new findings*). The rule G9 and G10 had each discovered and each
implemented locally now lives **once**, in a helper the platform-file assertions call — G2/G4/G5 were
in the same class and had never had it. Reach up one claim, findings unchanged at 0. And the
generalisation **inverted on first contact** — an existing silence-asserting test caught it within a
run — which produced the iter's most portable finding: a derived discriminator has **three**
outcomes, and collapsing *cannot-tell* into either verdict yields a guard that is confidently wrong.

**Type:** tik
**Status:** closed-fixed — planned scope landed exactly as pre-registered, including the "0 findings"
outcome the overview named as acceptable in advance.
**Gate:** NOT MET — 4 of 5, unchanged. Clause 5 is not re-cut.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n
— (5) cap-reached: n (3 tiks of 5) — (6) protocol-stop: n — Outcome: **continue**.
**Decisions:** `D-M257x-79-1` (145 → 3; a rule discovered twice and implemented twice is a rule not
implemented), `D-M257x-79-2` (a clone that cannot answer must never be read as answering "no" —
three-valued discriminators), `D-M257x-79-3` (the read-at-ref assertions had never inherited the
currency clause).
**Side-deliverables:** none — all in planned scope.
**Routes carried forward:** `CHECK-M257x-iter79-three-valued-discriminators` (a standing audit for a
harden pass: every derived discriminator in this family wants a *cannot-tell* branch — two of the
three found so far were missing one) · `CHECK-M257x-iter78-running-vs-declared` · all iter-77 routes
unchanged. **`CHECK-M257x-iter77-cross-repo-pin` CLOSED.**

**Lessons:**

1. **Size a routed count before acting on it — six for six now.** 145 became 3, and the 142
   non-instances were the argument for scoping the rule rather than widening it.
2. **A rule implemented twice locally is a rule not implemented.** Two assertions found the same
   mechanism one iteration apart and each patched itself; three sibling assertions in the same class
   were left with the hole.
3. **Every derived discriminator has three outcomes.** *Yes*, *no*, and *cannot tell*. Collapse the
   third and the guard is confidently wrong; which failure you get — blind or hostile — is an
   accident of which way the boolean fell.
4. **Keep the tests that assert silence.** The fixture that caught the inversion within one run
   exists only to check that nothing fires.
