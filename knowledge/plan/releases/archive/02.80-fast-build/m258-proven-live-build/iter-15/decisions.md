# M258 iter-15 — decisions

## D71 — The M217 pin guard fired on MY error, and re-pinning has two halves.

The first `up-injected.sh 3` attempt died before doing any work:

```
✗ FATAL: rext pin mismatch.
    the consumption clone is at : fast-build-m258-iter-14
    .agentspace/rext.tag pins   : fast-build-m258-iter-09
    This stack would run TOOLING THAT IS NOT THE ONE YOU THINK IT IS.
```

I had checked the clone out at the new tag and **not** updated the declaration. The guard offers two
remedies — revert the clone, or `DEMO_ALLOW_UNPINNED_REXT=1` — and **both would have been the wrong fix
here**: the first reverts the feature under test (the M236 shape exactly), the second disables the check
that just did its job. The right fix is the third one, which the guard's own wording implies: **update
the declaration to what you actually intend**, because `rext.tag` is an intent, not a lock.

Recorded because *"re-pin the consumption clone"* reads like one action and is two, and because this is
the guard catching a live instance of the class it was built for — worth knowing it works.

## D72 — The clause-3 waiter was disarmed deliberately, before the transition.

`autoarm-campaign.sh` (pid 7619) was still armed, sampling every 15 s and set to fire a **3-rep cold
campaign that tears `demo-1` down**. Firing during `END-M258-one-stack` would have destroyed the stack
being built, and could have left the user with **no** working stack at the moment the milestone claims to
guarantee him one.

Stopped, with the reason written to `.autoarm-outcome`. Clause 3 is an **opportunistic bonus** by user
ruling (`D52`) and the binding requirement is the end state; where they collide, the binding one wins.
For the record it never came close: `load1` minimum since arming **14.21** against a threshold of
**5.0**, over ~1 h. **Clause 3 remains NOT MET and is never to be recorded as met.**

## D73 — Stack ownership was re-resolved from docker, and iter-07's mapping had changed.

`iter-07 D23` recorded `demo-1` as owned by the **authoring** clone and `demo-2` by the **consumption**
clone. Measured today from the containers' own bind sources, **both** now resolve to the consumption
clone:

```
demo-1-postgresql-1 -> stack-demo/rosetta-extensions/demo-stack/stacks/demo-1/data/postgresql
demo-2-postgresql-1 -> stack-demo/rosetta-extensions/demo-stack/stacks/demo-2/data/postgresql
```

(The authoring clone still holds a **stale 263 MB `stacks/demo-1/`** — the same two-clone confusion that
produced iter-06's `D19` defect, where a roster went to the live path while manifests went to a stale
one.) **Ask docker where a stack's files are; never infer it from a script's location or from a previous
iter's note.** Both teardowns therefore run from the consumption clone.

## D74 — The 15-red batch grades `pt-world`, NOT the presenter world. Both were measured.

This distinction decided whether the user ends the milestone with a usable stack, so it is written down
rather than assumed. The batch gate is `reset-to-seed → pt-world → suite → restore`: the Playthroughs run
against **`pt-world`, the DECOUPLED TEST SEED** (`playthroughs.md`: *test data ≠ demo data*), and the
restore leg then rebuilds the presenter world. So `red_count: 15` is a verdict on test data.

What a presenter would actually drive was measured directly:

```
orgs=4  users=591  memberships=591  skills=42790   ·   12 cockpit seats across 4 stories
cockpit :37700 → 200 · web :33000 → 307 · studio :39000 → 302 · backend /api/health → 200
```

— the documented healthy shape (iter-06/07's *"4 story orgs / 591 users / 12 cockpit seats"*).

**Both readings are needed and neither substitutes for the other.** Tearing down on the red set alone
would have been wrong; so would ignoring it. It is **escalated** under `D-v28-3` with its causes
**unresolved and labelled as such**: 4 entries are plain timeouts at `load1` 26–33 with `retries: 0`
(`D28`'s false-red condition), 11 are data-shape assertions agreeing with autoverify's under-set-dress
warning — and **no `SQLSTATE 42P01` appears anywhere**, so the tempting "the newest platform moved a
table the seeder writes" story is **unproven**. Do not report it as diagnosed.

## D75 — `SETTLE-M258-iter13-studio-desk-cold-time` did NOT settle. Still no cold number.

iter-13 routed the studio-desk time question to this iter on the reasoning that `TIK-C`'s cold bring-up
would yield `ui_studio_desk` cold for free. **It did not, and the reason matters:** BuildKit reused the
layers from iter-13's own probe build (identical content, same context), so the real bring-up exported in
**1.5 s + 0.3 s unpack** instead of paying a cold export. A cache hit is not a cold measurement.

So the position from `D67` is unchanged and must not drift: **the studio-desk TIME axis is UNMEASURED on
this host.** The space win (350 MB/stack) is measured and stands alone. Settling the time half needs a
`--no-cache` A/B of the two Dockerfiles on a quiet box — cheap, but it was never run, and *"we'll get it
for free from work we have to do anyway"* turned out to be wrong precisely because the earlier work had
already warmed the cache.
