# iter-153 — decisions

## `D-M257x-153-1` — derive from the artifact that DECIDES the fact, not one that constrains it

iter-148 gave `/test-platform` a derived probe scope, which was the right *shape* against the wrong
*artifact*. `$STACK_ROOT/platform/docker-compose.yml` constrains what a stack **can** run; the stack's
own generated override decides what it **does** run. Measured at platform `0c91421` against a REAL
`gen_injected_override.py` emission: platform set **5**, override set **11**.

The repair unions the override in. The corollary that made the repair non-trivial: the union must be
**intersected with the probe registry** (a name with no row is not probeable), and the intersection must
**NAME what it drops** — otherwise "running and ungraded" is silently converted into "absent", which is
the same class of loss one level down.

## `D-M257x-153-2` — over-broad scope is LOUD; under-broad is SILENT, and that is why it survives

The two halves of iter-148's defect had identical causes and opposite lifetimes. The over-broad half
printed four false `down`s and was repaired inside one iter. The under-broad half printed **nothing** and
sat behind the same repair for five iters; harden pass 35 found it only by measuring, and even then
routed the fix forward. When grading a scope, both questions must be asked explicitly — *what does it
probe that it should not*, and *what runs that it never looks at* — because only the first announces
itself.

## `D-M257x-153-3` — "needs a live demo" was a precondition, not a fact, and it was refuted

harden pass 35 routed `FIX-M257x-h33-derive-includes-stack-override` as Fate 3 on the stated ground that
it *"cannot be verified without a live demo."* The re-survey falsified that in one command:
`gen_injected_override.py` is pure line-oriented text emission and `stack-demo/platform` is a real clone
at `0c91421`, so the stack's own override — and its whole flag matrix — can be produced here with **no
docker and no bring-up**. The whole fence runs in 1.9 s.

**The general form:** a routed-forward item's stated blocker is a claim like any other, and the re-survey
step (Phase 1 Step 0) is where it gets tested. This one had been carried for one pass; a wrong blocker
carried long enough becomes a fact nobody re-reads.

## `D-M257x-153-4` — retiring a gap-disclosure fence: re-point it, never delete it

Closing the gap made harden pass 35's own fence fail
(`test_the_derived_disclosure_names_the_services_it_excludes`): it asserted that `generate.sh`'s **source**
contains the literals `next-web-app` / `studio-desk` / `directus`, and the repair removed them
deliberately, because a hand-written list of what a mechanism excludes is the defect class this milestone
exists to end (iter-145's six count literals, iter-147's hand-appended tuple).

Deleting it would have retired a real property along with an obsolete spelling of it — the precise shape
the last harden pass booked against itself (*"the obvious remedy would have permanently disarmed two
rows"*). Re-pointed instead, at the **stronger** pair:

1. the disclosure block carries **no** service-name literal (so it cannot go stale), and
2. the real script, **when run**, still names what it left out.

Both halves mutation-controlled: re-introducing a literal fails half 1 (verified); the pre-fix script
fails half 2 (verified against `HEAD`, not a reconstruction).

The sibling test in that class — `test_the_derivation_returns_the_platform_set_not_the_stack_set` — is
**unchanged and still load-bearing**: it is about `platform_topology.py`, whose answer is still the
platform set, which is exactly why the union is needed.

## `D-M257x-153-5` — line 3 of the union, or "found nothing" and "read nothing" collapse

`scope-union.sh` echoes three lines, and the third (every service the override declares) exists only so
the caller can tell **"I read an override and it adds nothing probeable"** from **"I found no override"**.
Without it both produce an empty line 1 and the report would make the same confident statement about a
stack it read and a stack it never located — harden pass 35's own headline defect shape (*a mechanism
reporting a confident verdict about a subject it never read*), reproduced inside the repair for it.

The disclosure has three branches accordingly, and the no-override branch **says so out loud**: silence
reads as a complete scope.

## `D-M257x-153-6` — a probe registry gap is DECLARED, never invented

Three services the stack runs have no registry row: `hiring-app` (tracks `--no-ui`, measured) and
Clerkenstein's `fake-fapi` / `fake-bapi` (unconditional in every flag combination — **if either is down,
every login on the stack fails**, so the presenter's cockpit is dead while the report reads `✓ pass`).

Rows were **not** invented for them. A probe row needs a container name, a host port and a health target
that have been **observed on a live stack**; fabricating them would trade a declared gap for an
undeclared fiction. `STACK_INJECTED_SERVICES_NOT_PROBED` declares them with reasons, fenced in both
directions (arrival and departure), and adding real rows is routed as
`FIX-M257x-iter153-stack-injected-services-have-no-rows`.

## `D-M257x-153-7` — the bring-up's own scope tuple is hand-written, and is a THIRD answer

`up-injected.sh:2688-2690` derives the platform set correctly and then **hand-appends three literals**
(`next-web-app studio-desk` under `NO_UI`, `directus` under `NO_LOCAL_CONTENT`). Measured against the
generator, `--no-ui` drops a **fourth** service — `hiring-app` — that the tuple never names, so the
bring-up's verify scope is under-broad too, by one.

Not repaired here: editing the bring-up's verify tail is a second line of investigation and changes a
path that genuinely does need a live demo to grade. **It is fenced rather than left as prose** —
`test_the_ui_tier_tracks_no_ui_and_directus_tracks_no_local_content` asserts what the generator does with
both flags, and carries an inversion instruction so closing the route cannot silently delete the
assertion. Routed as `FIX-M257x-iter153-bringup-scope-tuple-is-hand-written`.
