---
iter: 21
---

# iter-21 — decisions

## D100 — the demo was faithful to the seed file and unfaithful to production

`init_policy.sql` ships 17 `p3` rows and the demo has exactly those 17. Nothing was dropped and no
seeder misbehaved — **the rext seeding fleet has never written a `p3` row** (`resetCasbinPTypes =
{g2, g3}`; every seeder writes groupings only). Platform commit `c6096d1` (2026-04-23) deliberately
removed `taxonomy:write` from the default because the capability *"should not be a universal
default"*, and added `sentinel/local_superadmin_grants.sql` for it — stated use case verbatim
*"Testing flows that require `taxonomy:write`"*. **No demo or dev stack has ever applied that file.**

**So a fidelity check against `init_policy.sql` passes and means nothing.** The reference has to be
production. `PolicyGrantsSeeder` applies the platform's own row idempotently, and `--policy-check`
fences the surface in both directions so the next divergence fails loudly rather than surfacing
fifteen iters later as "the form doesn't work".

## D101 — `SKILLER_*` criticality is `standard`, and that was measured

`critical` in this DNA means *the stack cannot build or come up* (GH_PAT is the reference). Marking
the three new genes critical takes the source from **100.0% → 86.7%** and exits 1 on every check until
someone provisions a key nobody has to hand — **a standing red, which D-v28-3 forbids.** The secret
gates one feature; it does not make a demo invalid. Classified `standard` on that measurement rather
than on the intuition that "a missing secret sounds critical".

## D102 — advance sequences AFTER the copy, to `max+1, is_called=false`

Both halves are load-bearing and both were RED-proven:

- **After, not before.** Moving the block above the COPY flips the assertion — a sequence advanced
  before its table loads reads `max` off an **empty** table and re-creates the exact bug.
- **`max+1, is_called=false`, not `max, true`.** The chosen shape is also correct on an empty table;
  `true` would silently burn id 1.

Columns are discovered from the **target's live catalog**, not a hardcoded pair, so a future migration
that adds an identity column is covered without re-capturing a snapshot.
