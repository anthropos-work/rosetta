# iter-229 — decisions

## `D-M257x-229-1` — the escape hatch records the gap and never buys a green

`BUILDBENCH_ALLOW_HOST_MISMATCH=1` lets an operator measure on a host whose profile does not describe it.
It would have been simpler to make the hatch a plain override. It is not, for a reason this milestone has
already written down (`platform-alignment.md` §8): *"an accept-the-gap escape hatch is fine … as long as it
RECORDS the gap rather than hides it"*, and the failure it names is a hatch that *"converts an honest
UNMEASURED into a quotable green."* `ALIGNMENT_ALLOW_UNMEASURED=1` promised in its own message to record the
gap and recorded nothing.

So the hatch does three things: it prints `RECORDED (hatch open)` instead of `REFUSED`, it stamps
`host_identity` into **every** rep ledger, and `identity.ok` joins `reps_ok == reps` and `report["ok"]` as a
**third, non-subsuming clause** of the campaign's exit code. A campaign run under the hatch produces data
and cannot produce a zero exit. Pinned by `test_hatch_lets_the_run_proceed_but_never_buys_a_zero_exit`.

## `D-M257x-229-2` — `mem_budget_mib` is observed and NOT graded, and the payload says so

Memory looks like an identity field and is not one: it is a **budget**, and `headroom_assert` already grades
budgets. Grading it here would double-report one fact in two guards with two different tolerances, which is
how a number acquires two homes and then drifts between them.

The choice is not to *skip* it. It is recorded in `observed` alongside the engine's measured `MemTotal`,
carrying `mem_budget_mib_note: "OBSERVED, NOT GRADED — a budget is headroom_assert's object, not
identity's"`. This milestone's own rule — **a CORRECT exclusion is still a defect while it is silent** —
applies to a guard's payload exactly as it applies to prose. A reader who wonders why memory did not decide
the verdict finds the answer in the verdict.

## `D-M257x-229-3` — the sanctioned-host profile was NOT authored here

The obvious "finish the job" move is to write `mac-mini.json` from the facts this iter measured. It was not
done, and the omission is deliberate on two counts:

1. **The numbers that matter cannot be measured on a busy box.** `laptop.json`'s own record shows the
   precedent: a full cycle was attempted there and **refused by clause 1 at peak load1 10.69**. This host is
   running agents. A profile authored now would carry a `lane_heap_measured_peak_mib` and a
   `projected_image_gib` measured under contention — a guess wearing a schema, which is exactly what
   `load_host_profile`'s required-key list exists to prevent.
2. **It is a third line of investigation** and fires the scope-creep tripwire.

`ROUTE-M257x-225-no-profile-for-sanctioned-host` therefore stays open — but it is now **loud**: before this
iter, borrowing a profile was graded silently; after it, `buildbench` exits 2 and names the arms.
