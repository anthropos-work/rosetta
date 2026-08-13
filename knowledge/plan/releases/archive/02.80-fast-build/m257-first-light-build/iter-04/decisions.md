# iter-04 — intra-iter decisions

## D1 — the host-class refutation is RECORDED, not ACTED ON

`D-v28-15` is a **binding user decision** and `state.md` restates it: *"a Mac is arm64/overlay2"*, *"the
Mac pays no unpack leg"*, *"M257's speed gate is un-measurable on the sanctioned hosts as written."*
This iter measured the opposite on the machine that decision moved dev/test to.

**Decision: record the measurement, propose the correction, edit neither the decision nor the gate.**
A sub-agent does not silently amend a binding release decision on its own reading — least of all one
whose whole failure mode was *a conclusion drawn about a machine nobody had measured*. The retraction
surface is enumerated in the close (`DOC-M257-hostclass-retraction`) so it can land **once**, with the
user's re-cut, rather than in two passes — the `D121` rule from this milestone's own pre-flight
(*"a half-retracted doc set is its own defect class"*).

The measurement is nonetheless held to a higher bar than the claim it refutes, because it contradicts a
standing decision: `docker info` alone was treated as a **weak** signal (per `spec-notes.md` F4, whose
whole lesson is that two agreeing weak signals are not one strong signal), and the finding rests on a
**different kind** of evidence — a controlled two-size probe showing the `unpacking to` leg exists and
scales (0.8 s @ 256 MB → 3.0 s @ 1024 MB), plus 19.3 s of it on the real 4.12 GB hiring image.

## D2 — the n ≥ 3 baseline was NOT run, and that is the honest call

The milestone still owes a measured baseline, and this iter had a host to measure it on. It was not run:

1. **The box was contended throughout** (load1 3.84–7.31) and **both of the user's stacks were resident**
   — `demo-2`, which the user has been actively validating on, plus a 5-container dev stack. The
   orchestrator's standing constraint is to ask before touching either, and a cold `--purge` cycle needs
   a slot.
2. **`laptop.json` already records the precedent**: a full cycle attempted on a busy workstation was
   **refused by its own headroom clause 1** at load1 10.69, and reported rather than overridden, because
   *"a cycle measured at load 10.7/10 says nothing about the bring-up."*
3. **A gate is measured on ONE declared, quiet host.** Publishing a contended cycle as this host's
   baseline would manufacture exactly the class of number v2.8 exists to retract — and it would then be
   mirrored into prose by the fence that pins baselines.

A **boolean** survives contention; a **timing** does not. So this iter took the booleans (the unpack leg
exists; the lane count is 2; the fences are green) and the single-lane timings, labelled every one of
them as contended, and routed the campaign as `BASELINE-M257-macmini-n3`.

## D3 — three findings were routed rather than fixed, and the tripwire is why

The iter declared a **three-line** planned shape in its `overview.md`. Each of these would have been a
4th line:

- **`FIX-M257-load1-units-vm`** — `buildbench.py:697` reads the **macOS host's** `os.getloadavg()` and
  clause 1 grades it against a `docker-desktop-vm` profile's `cores`, which is the **VM allocation**.
  One cheap grep was spent to route it with *evidence* instead of a guess; the fix was not made.
- **`MEASURE-M257-macmini-true-idle`** — needs the user's stacks down.
- **`PROFILE-M257-provisional-fields`** — the M255-inherited `provisional_fields` item. Worth noting the
  re-survey found its **sibling already done**: the two profile tests that hardcoded `("billion",
  "laptop")` are now globbed (`tests/test_buildbench.py:352-361`), so `macmini.json` is validated by
  being checked in. **An inherited-item list is a route list, and it goes stale like any other.**
