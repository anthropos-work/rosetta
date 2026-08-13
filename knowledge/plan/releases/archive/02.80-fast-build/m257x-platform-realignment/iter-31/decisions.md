---
milestone: M257x
iter: 31
---

# iter-31 decisions

## D1 — a bounded per-org role set, and why 12

Supporting members drew independently from the full 300-role pool, producing ~one job title per person.
The set is now bounded at `orgRolePoolSize = 12`, per organization, always containing that story's own
hero roles.

12 is deliberately NOT tuned to the org that exposed the bug: ~3 incumbents per title at size 40, ~17 at
the showcase's 200. Tuning it to make one assertion pass would be the Trap-A move (tune until it catches
nothing) in the other direction.

Rejected: raising the succession view's render cap (a platform edit — forbidden, and not ours to make);
and weakening the assertion to match what renders (the assertion's intent is correct; the world was wrong).

## D2 — the derivation is now in one place, and the sweep is the deliverable

Seven copies: six production seeders plus one in `target_roles_test.go`. The first sweep found four,
because two spell the index `idx` and read a different field. Completeness — not presence — is what the
guard has to assert; that rule is iter-25's, and it is the reason this landed rather than half-landed.

The test copy going RED when the production six were unified is kept in the comment as the demonstration
that the desync is real, not theoretical.

## D3 — kill the binding run rather than let it finish

At 52 of 209 specs the tenancy negative control failed, which meant the run was measuring a world carrying
a regression. Killed it. A completed run would have produced a precise, citable, meaningless number, and
this milestone has already had to withdraw one such number (iter-14's three-green-cycles).

Cost: ~30 minutes of wall clock and the deferral of the clause-2 confirmation to iter-32. Accepted.
