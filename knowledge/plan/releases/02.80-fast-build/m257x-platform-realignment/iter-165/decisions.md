# iter-165 — decisions

## `D-M257x-165-1` — the accept-side audit is NOT mechanical, and the iter's own probe is the proof

The hypothesis was that a waiver is a decidable accept clause: it names a `path` and a quoted form,
so either the form is still there or it is not.

**It is not decidable from outside the guard, and three separate shape assumptions were wrong** —
each one producing confident findings:

| assumption | reality | fabricated findings |
|---|---|---|
| normalising punctuation is enough | runs of whitespace survived, so every waiver whose quote spans a **wrapped markdown line** failed to match | **10** |
| every waiver names a `path` | `repair_reach_waivers.json` holds **dispositions keyed `path:line`** — a different schema entirely | **1** |
| `form_contains` matches raw file text | it matches the guard's **derived claim form**, not prose | **1** |

**11 of a claimed 11 dead waivers were artifacts of the auditor.** Corrected, the count over 20
entries is **0 provably dead**.

This is the second time in this run: iter-163's first draft reported **346** findings that were a
pairing cross-product. Different mechanism, same shape — **an instrument's preprocessing IS part of
its predicate**, and a normalisation step written in five seconds is a predicate written in five
seconds.

**What stopped it landing was cheap and repeatable:** the number was implausible (11 of 20 waivers
rotting silently, in a milestone that re-runs these fences every iter), so the *instrument* was
re-checked before the *corpus* was. `§5`'s standing rule, applied to an auditor rather than to a
document.

## `D-M257x-165-2` — the falsification is that the guards do not REPORT their accept decisions

Having refused re-implementation, the correct instrument is to **ask each guard which waivers it
honoured**. That is not available:

- `repair_leak_guard` is **diff-scoped** — at HEAD it exits `CANNOT RUN — no candidate shingles`,
  correctly, because its subject is a repair diff and there is none. Its waivers have no subject to
  be honoured against outside a repair.
- `claim_twin_guard` runs against the live corpus and prints **nothing** about waiver usage, so a
  waiver that never fires is indistinguishable from one that fires on every run.

So the routed survey sharpens from *"audit the accept side"* to something buildable: **every fence
that carries a waiver must report, per run, which waivers it honoured and which it did not** — the
accept-side analogue of the reach numbers this milestone already demands on the fire side
(`§8`: *a fence whose reach shrinks in silence is the failure this milestone keeps finding*). A waiver
that never fires is precisely a reach hole, on the other axis.

**Not claimed:** that any waiver is stale. Nothing was found, and this iter says so.

## `D-M257x-165-3` — closed-no-lift, and no repair is landed on a measurement this iter withdrew

The three-fate rule's Fate 1 for an iter includes *"close-no-lift with documented falsification"* when
the mandated investigation actually happened. It did: all 20 entries were enumerated, three schema
shapes were read at source, and the corrected count is published.

**Deliberately NOT done:** deleting the one `repair_leak` waiver whose quote does not appear in the
raw file. It matched a *derived form*, the probe compared *prose*, and removing a waiver on that basis
would make a fence louder for the wrong reason — the exact error the `_README` of that file warns
about (*"`reason` is not decoration: it is the thing a future reader grades the waiver by"*). One
withdrawn measurement is not a licence to act on its residue.
