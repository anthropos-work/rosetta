# iter-287 — decisions

## D-M257x-287-1 — placement is part of a claim

`INVITATION_HMAC_SECRET`'s `exit 0` failure was stated **twice** in `setup_guide.md` before this iter,
both times correctly. It was still missing, because both statements sit **upstream of the failure** — one
inside the argument for using `/dev-up` instead of the manual path, one in the secret-coverage section a
reader following the manual path has already passed. Neither is where someone is standing when `make ps`
shows four containers instead of five.

**Repeating a fact at the point of use is not duplication.** What makes that safe to assert here rather
than as a preference is iter-282's `prose_twin_guard`: it grades copies that **disagree**, not copies. A
third statement of the same claim is free; a third statement carrying a different number is a finding.
Run after this edit: **0 RED**.
