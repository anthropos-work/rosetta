# iter-139 — decisions

## `D-M257x-139-1` — iter-138's `127 rotted / 57.2 %` is RETRACTED. Precision 0/12, Wilson95 [0.0, 24.3] %

A stratified 12-case audit — **strata and selection rule sealed in `overview.md` before the first case was
opened**, taken in census order from the top of each stratum, no cherry-picking — classified **12 of 12
as FALSE POSITIVES**.

| # | case | what the citing line actually is |
|---|---|---|
| 1 | `platform-alignment.md:2595` → `:788` | *"two of iter-45's five defects are relationships between line numbers — `:788` citing `:447`"* — a **historical example quoted**, not a live pointer |
| 2 | `platform-alignment.md:1108` → `:1305` | *"(`:1305` was the first, iter-84)"* — a claim about a **past** position |
| 3 | `build-budget.md:394` → `:319` | `(`:279`, `:319`)` — **continuation pins into another file** named earlier in the sentence |
| 4 | `security_compliance.md:386` → `:129` | `(`isThrottlingError`, `:129` / `:166` / `:325`)` — continuation pins into a Go source file |
| 5 | `platform-alignment.md:51` → `:106` | `` - `:106` — atlas-applies a hardcoded 4-tuple `` — continuation pin into `migrate-demo.sh` |
| 6 | `ai_architecture.md:325` → `:75` | *"(`v3/validator/skills.go:53-64`) then … `:75` computes"* — **explicitly** a continuation pin |
| 7 | `shared_libraries.md:198` → `:259` | `` `AIManager.getClient` (`:259` and `:332`) `` — continuation pins |
| 8 | `platform-migration-status.md:106` → `:183` | `` `docker-compose.yml:170-171` … (`:183`) `` — continuation pin |
| 9 | `messenger.md:53` → `:29` | *"there is **no** `:29` declaration at all"* — a **negated** pin about another repo |
| 10–11 | `messenger.md:60` → `:63`, `:62` | `` `app/main.go:15`, `:62`, `:63` `` — continuation pins |
| 12 | `platform-alignment.md:54` → `:108` | `` - `:108` — `[ -d "$DEV/$r" ] || continue` `` — continuation pin |

**Precision 0/12 = 0.0 %, Wilson 95 % [0.0, 24.3].** The figure `127 rotted (57.2 % of 222 decidable)` is
**withdrawn**, in place, in `iter-138/progress.md` and the milestone ledger. It is not re-qualified,
narrowed, or restated with a caveat — `D-M257x-122-3`'s class.

## `D-M257x-139-2` — the disclosed floor was NOT the whole floor, and disclosure created false confidence

iter-138 disclosed an `out-of-range-then` bucket of **241** and named its cause exactly right:
*"largely cross-file continuation pins, which the probe reads as same-file."* Booking that honestly is
what made the remaining **222 look clean** — *"decidable"* was published as though the known failure mode
had been quarantined into the bucket that carried its name.

**It had not.** A continuation pin lands in the `out-of-range` bucket only when the cited line number
happens to exceed the **citing file's** length. For a 3,100-line protocol doc, almost no pin does — so the
**same failure mode passed straight into the decidable set**, undisclosed, and dominated it.

> **Rule.** **A disclosed limitation is quarantined only if you show the boundary holds.** iter-138
> named the failure mode and then reported a number computed over the population it had *not* excluded it
> from. **Naming a floor is not the same as bounding it** — and the disclosure made the number *more*
> persuasive, not less, which is the trap. Sample the "clean" bucket for the failure mode you just
> disclosed; if it is there, the bucket is not clean.

This is the sibling of the milestone's standing *"do not read a marker count as a defect count"*
(iter-132) and of `§8`'s *grade the cannot-tell*: here the cannot-tells were graded, correctly, and then
the **can-tells were never checked for the same disease.**

## `D-M257x-139-3` — `corpus_citation_guard.py`'s exclusion was better founded than iter-138 credited

iter-138's `D-M257x-138-1` said the fence's *"bare `:NN` pins are not mechanically decidable"* exclusion
was *"broader than the evidence requires."* **Measured, the exclusion is right and for a reason iter-138
did not have:** in this corpus a bare `` `:NN` `` is **overwhelmingly a continuation pin into a file named
earlier in the sentence** — not a same-file self-citation at all. Resolving the *head* is the hard part,
and it is the part the fence declares it cannot do.

**`D-M257x-138-1`'s RULE survives; its APPLICATION here is withdrawn.** *"An exclusion is only as narrow
as the predicate that justified it"* remains true and worth keeping. What iter-138 got wrong was assuming
the un-tested predicate (**rot**) was free, when the blocker for both predicates is the **same
unresolved head**.

> `adj-E`'s five genuine same-file rotted anchors were found **by a human reading five sentences**, and
> that is not an accident: they are a *rare* form, and the machine that would find them must first solve
> the head-resolution problem this fence declines.

## `D-M257x-139-4` — what STANDS from iter-138, stated so the retraction is not over-read

The retraction is of **one number**, not of the iter.

- **All 9 citation repairs stand.** Each came from `adj-E`/`adj-D`'s hand-verified list and each was
  **re-derived by opening it** before repair — none came from the probe.
- **`D-M257x-138-3`** (name the construct, never re-pin) stands, and is strengthened: if a bare pin's head
  is unresolvable even to a purpose-built probe, it is unresolvable to a reader.
- **`D-M257x-138-5`** (choose suites by what you changed) stands — unrelated to the probe.
- **`FIX-M257x-iter138-anchor-rot-fence` is RE-SPECIFIED, not cancelled**: its first deliverable is now
  **head resolution for bare pins**, and it has **no baseline** until that exists. A fence built against
  `127/222` would have been built against a number that is 0-for-12 on audit.
- **`FIX-M257x-iter138-127-rotted-pins` is WITHDRAWN** — there is no such work list.

## `D-M257x-139-5` — the audit cost one iter and the error survived one iter

Booked deliberately, because the milestone books its own loop in both directions. iter-138 published an
unaudited instrument's output; iter-139 audited it and retracted it **before** any of the 127 was
repaired and **before** the fence was built against it. **Nothing downstream consumed the wrong number.**

> The rule this milestone keeps re-earning: **an instrument is not a measurement until the instrument is
> measured.** Nine times now. The cost of the ninth was one iter; the cost of skipping it would have been
> a fence with a fabricated baseline and 127 unnecessary edits.
