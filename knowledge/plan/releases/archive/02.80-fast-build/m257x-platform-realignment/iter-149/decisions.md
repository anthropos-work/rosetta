# iter-149 — decisions

## `D-M257x-149-1` — a stale mutation staging is a perfect forgery of the bug its battery catches

The M220 mutation battery stages its mutants **beside** the real subject rather than in `/tmp`, and
deliberately: both `up-injected.sh` and `dev-stack` resolve their siblings from `$HERE`, so a subject in a
temp dir cannot find the code it drives — the battery's own baseline assertion caught exactly that on its
first run (14 tests "RED" on an *unmutated* copy). `tearDown` deletes them; an **interrupted** run's
survive, and `.gitignore:32` says so in as many words.

What nothing did was sweep them. Measured here: **33 executable `.m220-mutant-*` copies of `dev-stack`,
oldest 2026-08-04**, each carrying verbatim the tailnet-URL line iter-146 repaired.

**Decision: sweep in `setUpClass`, age-gated at 1 h — never unconditionally.** A mutant belongs to a live
run for as long as that run holds it; an unconditional sweep would let this battery delete a concurrently
running sibling's staging mid-measurement, and that failure surfaces as a *mutation result*, not as a
missing file. No nested run has ever approached an hour (the ladder's ceiling is ~900 s × 3), so anything
older is abandoned by construction. Proven on the real leftovers before the fixture was spent: 33 → 0.

**The generalisation:** exhaust that is byte-identical to a repaired defect is not neutral clutter. It is
indistinguishable from a regression to every later sweep, and it inflated this iter's raw signal by 66 %.

## `D-M257x-149-2` — bind a fence's subject set to an existing fenced one

The widened emitter fence needed a list of retired services. One already existed and was already fenced:
`claim_census_guard.ARCHIVED_SERVICE_NAMES`, the row set of `platform-migration-status.md`, machine-fenced
against the platform's `repos.yml` in both directions by `platform_alignment_guard`.

**Decision: import it.** A thirteenth hand-maintained list is precisely the defect this milestone keeps
finding (`§2`'s tuple defect; iter-129's enumeration; iter-148's registry). The consequence is that a
service entering or leaving the migration map reaches this fence with nobody re-typing anything, and the
anti-vacuity control asserts the port table is a **subset** of the imported names — a port declared for a
name the map does not carry would mean fencing something still live.

This narrows, without closing, `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer`: it is the
first seam in the family to consume another guard's subject set rather than restate it.

## `D-M257x-149-3` — when a census returns zero, RED-proof the instrument

The census found **0** emitters. A fence written on that basis is unfalsifiable from the tree alone: an
arm that matches nothing looks exactly like an arm over content that contains nothing.

**Decision: two independent proofs, both mandatory before the zero is published.**
1. **On a real answer key** — the fence run against the pre-iter-146 `dev-stack` (`1a44b97^`) returns the
   original defect line, and against the current file returns nothing.
2. **Per arm, on synthetic content** — container and address have no real occurrence to prove them, so
   each is shown tripping on content written to trip it. A decorative arm is a fence that reads wider
   than it is.

## `D-M257x-149-4` — a carve-out with a named owning route is a waiver; one without is a hole

`stack-verify/lib/services.sh` is in the emitter allowlist and legitimately carries four retired rows —
they DECLARE a probe target a scoped run filters out; they emit to nobody. iter-146 warned explicitly
against widening the fence un-audited, and matching those rows would create a standing RED with a known
owner, which is how a fence gets switched off.

**Decision: carve it out, in code, with the reason and the owning route named
(`SURVEY-M257x-iter148-registry-is-hand-maintained`), plus an anti-vacuity control asserting the
carve-out names a file the fence actually scans.** A carve-out for an unscanned file is dead configuration
that reads as coverage.

## `D-M257x-149-5` (side) — fence the property a comment asserts, when the code cannot perform it

`claim_census_guard.REXT_SECTION_NAMES` called itself *"derived from the monorepo's own layout"*. It was
declared, and had drifted — `stack-secrets` missing, 10 declared against 11 on disk — so claims naming the
section behind `/stack-secrets` resolved to no known artifact and left the census silently. The same
enumeration defect iter-129 repaired in `CLAUDE.md`, one layer down, inside a guard whose job is noticing
omissions.

**Decision: do NOT make it derive.** The module is pure data with no repo-root notion and is imported from
copies whose layout differs; an import-time `listdir` would add a host dependency to a guard that has
none. Instead: fix the list, make the comment honest about being declared **and why**, and move the
derivation into a test that compares the tuple against the on-disk layout in both directions. RED-proofed
— the pre-fix tuple fails it, the post-fix tuple passes.

Landed as a **separate commit**; it is a side discovery and does not grade this iter's planned scope.
