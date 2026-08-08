# iter-169 — decisions

## `D-M257x-169-1` — the m255 BASELINE RED was graded before the open route was acted on

**Decision:** reproduce the failing test **outside** the staged tree before attributing it to the staged
dependency class, even though the route for that class named this exact battery and the failure signature
matched it exactly.

**Why it mattered.** `test_m255_mutation_battery` was RED at HEAD with
`stack-core/tests/test_buildbench.py is RED before any mutation` — the published signature of
`FIX-M257x-iter111-staged-battery-dependency-is-underived`, whose **last open member was this battery**. The
cheap move — append a filename — was available and would have been wrong. One command
(`python3 -m unittest tests.test_buildbench.TestSampler.test_the_stop_event_does_not_shadow_a_Thread_attribute`)
showed the same test failing **in place**, which exonerates staging entirely.

**Root cause:** the assertion was `callable(getattr(s, "_stop", None))`. **CPython 3.14 removed
`Thread._stop`.** The Sampler is correct and has never touched `_stop`. Booked as `§5` rule 74.

## `D-M257x-169-2` — the repaired assertion pins the property, and the ANTI-VACUITY CONTROL rejected the
first attempt

**Decision:** replace the spelling-pinned assertion with a computed one — *no name `threading.Thread`
occupies on this interpreter* — and keep the original `_stop` clause **guarded by whether the spelling
exists**, rather than widening a regression test away.

**The control earned its place immediately.** The first computed form used
`set(vars(sub)) - set(vars(bare_thread))`, and the control (a subclass deliberately shadowing `_target`)
**passed when it must fail**: `_target` is assigned by `Thread.__init__` itself, so the subtraction removed
exactly the collisions worth catching. Ownership is about **who assigned**, not who ends up in `__dict__` —
so the final form records assignments with a `__setattr__` recorder while bracketing out the window in which
`Thread.__init__` is assigning its own. The control now fails when the invariant is broken and passes when
it is not, on this interpreter.

## `D-M257x-169-3` — close the class at its POPULATION, not at its last member

**Decision:** ship `tests/test_battery_stage.py`, a fence over the mutation-battery population, rather than
only migrating m255 and closing the route.

**Rationale.** The class recurred five times and each repair appended one filename. Migrating the sixth
member leaves member seven completely unfenced, and the record says member seven is not hypothetical. Per
iter-162 (*fence a registry's completeness, never its contents*) and `§5` rule 73 (*a glob is not a
derivation*), the population is derived by **property** — a test module binding a module-level MUTANT
registry — never by `*mutation_battery*.py`.

**Measured population: 7.** Six stage a file set and now derive it; `test_m220_mutation_battery` is the one
exemption and it **proves itself on every run** (it copies nothing — it mutates one subject into a
gitignored sibling beside the real tree). An exemption that is merely asserted is a hand-list with better
manners.

## `D-M257x-169-4` — the widening resolves imports sibling-first; it does not enlarge the haystack

**Decision:** `battery_stage.local_deps` resolves each import against the **importing file's own
directory** first — what the interpreter does — then falls back to the two root-relative conventions the
single-section batteries rely on.

**The escalation condition in this iter's `overview.md` was cleared by measurement, not by argument:** all
**five** already-migrated batteries derive **identical** sets before and after the change
(`claim_twin` 6=6, `repair_reach` 4=4, `repair_postcondition` 6=6, `repair_leak` 5=5,
`mechanical_fences` 10=10; zero added, zero lost). Two properties are asserted permanently rather than left
as prose:

- **Over-approximation is deliberate.** Function-scope imports are followed like module-scope ones.
  Under-staging reports a BASELINE RED with no attributable test; over-staging costs one `shutil.copy2`.
  Without a test stating the asymmetry, somebody "fixes" the imprecision and re-opens the class.
- **A stdlib shadow is REFUSED, not staged.** A repo file whose name collides with a stdlib module the
  seeds import would shadow it inside the staged tree — a staged-only divergence, the exact class the
  helper exists to end. `local_deps` raises.

## `D-M257x-169-5` — the literal-registry classifier was narrowed, and the narrowing carries its own proof

**Decision:** exclude tuple-**unpacking** (`MD, AN, VA = "a.py", "b.py", "c.py"`) from the
literal-path-registry classifier, and add a permanent control that the narrowing did not blind it.

**Why this is a correction, not a convenience.** The first draft flagged that exact line in
`test_m257x_mechanical_fences_mutation_battery` — a battery whose stage set **is** derived, and whose three
unpacked names are *seeds fed into the derivation*. The RHS is an `ast.Tuple` node, but no collection exists
at runtime and there is nothing that can fall out of date.

**iter-158's rule applies to every narrowing** (a proposed narrowing there would have graded 14 of 14 broken
checks green), so `test_06_the_narrowing_still_catches_a_real_registry` grades the **same three path
literals in both syntactic forms** and requires opposite verdicts.
