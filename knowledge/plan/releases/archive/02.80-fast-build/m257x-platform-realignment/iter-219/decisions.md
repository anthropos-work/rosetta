# iter-219 — decisions

## `D-M257x-219-1` — the pre-registered VALIDITY condition fired, and the recorder is refused, not corrected

V3 was sealed as a sentence about the instrument rather than a number about the subject: *the
instrumented run's counts must equal the uninstrumented run's.* They did not — one test flipped, and it
reproduces in 0.19 s. Correcting the recorder inside the iter that sealed that sentence would be
overriding a pre-registration after seeing the data, which is the failure this milestone has punished in
its own instruments nine times over. The redesign is **routed with its cost stated**, not landed here.

## `D-M257x-219-2` — the reading is REPORTED and explicitly NOT claimed

The perturbed run says: 9,813 writes, **1** inside the repo (a `.yaml`), **0** `.py` inside the repo,
**0** matching the exposed shape. Suppressing it would be worse than disclosing it — but it is not a
measurement, because the instrument failed its own validity condition in the same run. *An instrument
that states its own invalidity must not exit 0.* `SURVEY-M257x-h42` therefore stays **OPEN**, and the
tempting close ("measured 0, nothing exposed, done") is exactly the one the rule forbids.

## `D-M257x-219-3` — the perturbing half IS the measuring half, so the answer is a redesign

Measured both ways rather than argued: with the `open` hook the module perturbs (61/1) and the writes are
visible; without it the module is clean (62/0) and **0 of its 140 writes are seen**, because it writes
exclusively through `open`. A patch that removes the perturbation removes the census. The three arms in
`test_m257x_write_census_perturbation.py` pin the split in both directions — each one fails if the
perturbation disappears, which is deliberate: *a waiver outliving its subject reads as coverage*, and so
does an obstacle.
