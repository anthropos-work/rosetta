# iter-06 — decisions

## D1 — the basis is the machine the sample came FROM, and the live paths observe it rather than declare it

**Decision.** `load1_core_basis` prefers an `observed_host_cores` passed by the caller over anything the
profile declares, and both live paths (`assert-headroom`, the rep loop's post-gate) pass `os.cpu_count()`
from the *same process* that read `os.getloadavg()`.

**Why observation outranks declaration here.** Every other input to `headroom_assert` is a *budget* — a
property of the host that a measured profile is the right home for. The load1 core count is different: it
is a property of **the reading**, not of the host, and the failure mode being fixed is precisely that the
numerator and denominator came from different machines. Taking both from one process makes that class of
error unrepresentable rather than merely detected. The declared field remains for the offline case (grading
a load1 recorded elsewhere) and as a documented host fact.

**What this does not do.** It does not verify that the profile *describes* the host — `profile_describes_host`
already owns that, and it grades a docker-desktop-vm's `cores` against **engine** NCPU, which is correct for
identity and is a different question from clause 1's. Two checks, two quantities, stated so the overlap is
not mistaken for redundancy.

## D2 — an unknown basis FAILS; it does not fall back

**Decision.** A measured `peak_load1` with no establishable core basis appends a `load1_core_basis` failure.

**Why not fall back to `profile["cores"]`.** That fallback *is* the defect. A silent fallback would preserve
exactly the behaviour being removed while adding a function that looks like a fix — the "capability probe
that fails OPEN" shape, where the new arm disarms the check it was added to guard. So the arm was proven RED
with its precondition absent (`test_a_vm_profile_with_no_host_core_count_is_UNGRADEABLE`) *and* by mutation
(restoring the old basis turns three tests red), rather than asserted to work.

**The complement is also pinned.** `pre_rep_assert` legitimately has no load1 yet, and an *unmeasurable*
input is not an *ungradeable* basis. Conflating them would fail every pre-rep gate on this host class, so a
`None` load1 with a `None` basis is explicitly still OK
(`test_the_ungradeable_arm_does_NOT_fire_when_there_is_no_load1_to_grade`).

## D3 — `kind` is required by the loader

**Decision.** `kind` joins `name` / `cores` / `budget_source` / `measured` as a loader-required key.

**Rationale, which is the loader's own.** Its docstring already argues that a profile *"that does not say
what its memory budget is a budget OF is how the M239-F1 ENOSPC got past a green pre-flight."* `kind` is the
same argument applied to cores: without it, `cores` could be a host total or a VM allocation, and those are
different machines. All three shipped profiles already declared it, so nothing in flight breaks; what the
change stops is a **fourth** profile arriving without it and silently becoming ungradeable.

## D4 — `laptop.json`'s `host_logical_cores` is DECLARED, not measured, and says so

**Decision.** Record `10`, sourced from that profile's own `budget_source` text (*"a 10-core / 16 GiB M1
Pro"*; Apple silicon has no SMT, so logical == physical), with the note stating plainly that it was not
re-measured.

**Why not just leave the field off.** Leaving it off is defensible — the machine is retired — but it would
lose the finding that matters: that profile's recorded clause-1 refusal at load1 10.69 was **correct only by
coincidence**, because there the VM allocation and the host core count are both 10. A real refusal on record
is the strongest possible argument that a clause is sound, and here it was not evidence at all. Declaring the
field, with its provenance, is what makes the coincidence visible and testable
(`test_the_laptop_refusal_on_record_still_refuses` pins that the historical verdict survives the fix).

**The provenance label is load-bearing, not politeness.** This milestone's entire opening defect was a host
fact generalised from one machine to another without saying so.
