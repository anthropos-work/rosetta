# iter-221 — decisions

## `D-M257x-221-1` — the `.md` half is fenced and the rest is ROUTED, because absence is not correctness

**2,117** file citations live in the corpus sources and **876** are markdown. Only those can be resolved
against pools that exist on this box; the other **1,241** name files in service repos that may or may not
be cloned. `derived_value_guard`'s rule governs: *"I could not check it"* and *"it is correct"* are
different findings and only one is coverage. Widening needs an **UNMEASURED** bucket keyed on clone
presence, which is a design decision with its own evidence — routed, not assumed.

## `D-M257x-221-2` — the residual is a declared TAXONOMY, and a class may not be widened to absorb a defect

All **19** non-resolving citations were adjudicated by hand **before** anything was landed, into four
classes — **negated · explicitly future · cross-repo · git-ignored workspace artefact** — with **zero**
corpus defects. The pre-registered rule stands in the fence's own failure message: a citation fitting no
declared class is **reported as a corpus defect**, never absorbed by loosening a class. The `negated`
class is the load-bearing one and could not have been derived by pattern: `studio-room.md` cites
`guidance.md` precisely to record that it does not exist, and iter-214 refused an entire widening over
that shape.

## `D-M257x-221-3` — the hardening ledger is a SECOND disposition surface, and the registry does not read it

`route_disposition_guard` went RED on this session's own writing: `SURVEY-M257x-h42` was **CLOSED at
iter-200** and iter-219 published it open. The premise came from `hardening-ledger.md`, which re-listed
it as *routed forward* in passes **48, 49, 50 and 53** — all after the closure. The guard fences iter
routes-blocks; the ledger's lists are outside its population, so the two surfaces disagreed for six
passes and an iter spent its opening claim on the disagreement. Corrected in iter-219's block in the
registry's grammar (appended, not substituted) — **57 closures, 0 contradictions** after. The route is
new and the fix is not in this iter's scope: **this skill is forbidden to write `hardening-ledger.md`**,
so grading it belongs to the fence, not to a hand edit here.
