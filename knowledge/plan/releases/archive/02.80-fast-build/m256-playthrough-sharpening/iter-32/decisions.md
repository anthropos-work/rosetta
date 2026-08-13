# iter-32 — decisions

## D123 — the org-less capability has TWO halves, and iter-30's measurement only ever exercised one

iter-30 priced `individual.UC1` by DELETING a hero's membership row on `demo-2`. That answered the
load-bearing question (the app serves her) and it is why this iter was cheap to open. But it left Clerk
still carrying her organization — so what was measured was *a user whose DB membership vanished*, not *a
user who never had one*.

The seeded state needs both halves:

| half | where | what it is |
|---|---|---|
| **DB** | `UsersSeeder` | no `public.memberships` row, no Sentinel `g2`/`g3` grant |
| **IDENTITY** | `roster.go` | no Clerk org claim — `org_auth_id` / `org_eid` / `org_name` / `org_slug` / `org_role` all empty |

The identity half is load-bearing, not tidiness: the host page mounts `<OnboardingIndividual>` versus
`<OnboardingUser>` off `useGetClerkOrganization`, and `UserStatusContext` routes on the Clerk org claims.
A hero with the DB half only would be served the ENTERPRISE flow, and the Playthrough written to prove the
individual journey would have proven the ordinary one.

What is deliberately NOT changed is her **email**, which keeps the org domain: it must stay byte-identical
to the `public.users.email` the `UsersSeeder` writes from the same `emailForMember` call. An org-domain
address for a user with no org reads odd; breaking the login-equals-the-row invariant `roster.go` exists to
hold would be worse.

## D124 — `verified: 0` is REQUIRED of her, not permitted — the invariant is the capability

The obvious design is a flag that suppresses the membership row. The one that ships also **inverts the
end-user `verified > 0` rule** for her, and that inversion is what keeps the change small.

A verified skill's fan-out is org-scoped (`jobsimulation.sessions.organization_id`, the membership mirror
the manager scoreboard reads). So a "verified org-less hero" writes rows tying her to an org she is not a
member of — rows that break no FK and are wrong invisibly (iter-28 **D112**). Requiring `0` makes the whole
persona/profile/activity chain skip her **because she declares nothing**, instead of because fifteen
seeders each learned a special case. A manager is refused for the mirror-image reason: the manager vantage
IS the org-intelligence seat.

Both rules are enforced through `EffectiveVerified()`, so the fine-grained `skills.verified` block cannot
bypass the coarse field — a rule that reads only one spelling is the shape of every enum-fallback defect
this milestone has found.

## D125 — the FK surface is machine-checkable; the org-scoped tail is not, and the fence says so

iter-30's framing was *"the FK list is the entry fee; the org-scoped-row tail is the real work"*. Both
halves showed up inside one iter, and they need **different** guarantees:

- **The loud half is STATIC.** Membership ids are a deterministic `membershipUUID(prefix, i)` and every
  membership-keyed seeder walks the population, so **the call sites ARE the FK surface**.
  `orgless_writers_fence_test.go` enumerates them from source (9 today) and requires each to consult the
  predicate — a new membership-keyed seeder cannot be added without it. It announced itself for real: the
  `SuccessionSeeder` FKs each member's first population session, which she no longer has, and the first
  reseed after the capability landed **failed the whole seed by name**
  (`interview_extraction_results_sessions_session`). Friendly, exactly as iter-30 predicted.
- **The quiet half has NO static signature.** A seeder can write an `organization_id`-bearing row through
  any number of shapes. So it is covered by **measurement**: reseed, then sweep the live DB for her uuid
  across every uuid column in `public` + `jobsimulation`. That found 2 `jobsimulation.sessions` rows
  carrying an `organization_id`, plus 4 activity events, 2 skill-path sessions, 4 personal assignments and
  3 bookmarks — none of which breaks anything, all of which were wrong.

**Neither check subsumes the other, and the fence's header says so.** A fence that implied it covered both
would be the more dangerous artifact, because the half it does not cover is the half that fails silently.

Post-fix the sweep reads **10 tables, all user-scoped, 0 carrying `organization_id` or a membership FK, 0
casbin grants.** She has a name, a language and a career, and no organization.

## D126 — TWO pre-existing fences refused the seat before it shipped, and the second one was right about the product

The best thing that happened this iter is that neither refusal came from review.

1. **The M224 curated-pool fence** rejected her first role: `Product Designer` classifies to
   `curatedNone`, so her claimed tail would have filled from the FLAT public pool — *the taxonomy's
   alphabetical junk head* (`15Five`, `3dcart`, `24-hour dietary recall`), the exact M219 defect.
2. **The M219-R8 ladder-depth fence** then rejected the replacement: the `operations` family offers 78
   candidate names, her default claimed tail wanted **65**, and 65 + the 15-name attrition margin is 80 —
   so she would have drawn her family dry and shipped a *silently thin* profile.

The fence's own suggested remedy was to grow the allow-list. **That would have been the wrong fix.** The
finding underneath it is that an org-less DAY-0 solo user should have **no claimed tail at all** — a hero
with zero verified skills and 65 self-rated ones is not a coherent person. `claimedTailCount` now returns 0
for her, both fences pass, and her footprint got smaller rather than the taxonomy allow-list getting bigger.

*A fence that suggests a remedy is still only telling you WHERE; the remedy is a judgement.*

## D127 — an irreversible write needs a fresh world BEFORE EACH MUTANT (the mutant protocol, corrected)

The first mutant pass was **thrown away**, and the reason is worth more than the mutants were.

Run against a world the green drive had already consumed, **N1 went RED at the wrong line** (the liveness
assert, because `/onboarding` already redirected) and **N4 PASSED** — a false pass that looks exactly like
a weak assertion. Neither result was about the assertions at all; both were about state.

Re-run with a `--reset-only` before each, both are decisive and land where they should: N1 red at the
lands-in-app wait, **N4 red at the redirect** — which is the one that matters, because N4 is iter-27's
standing Q1 (delete the action, keep only the read-back) and it proves the final cannot pass without the
write.

**The rule:** for a Playthrough whose write cannot be undone through the UI, the mutant protocol is
*reset → mutate → run*, every time. The spec's own failure message already warned about this ("either the
world was not reset-to-seed…") and it was still missed — reading a warning is not the same as obeying it.

## D128 — the run-3 flake is REPORTED, diagnosed, and NOT re-rolled away

Gate run 3 failed with **one** failure: `pt-workforce-succession`, *"the projection computed a card for the
org's own seeded role (DevOps Engineer)"*. Runs 1, 2, 4, 5 and 6 were clean (`200 passed`, rc 0).

**It is not iter-32's.** The Playthrough reads **Org A**; the new seat lives in **Org B**. Measured after
the failure: Org A has **40** memberships and DevOps Engineer occupancy **1**, both unchanged, and Org B
has 19 against `size: 20` — exactly the one slot the org-less hero consumes without becoming a member.

**It is a recurrence of `PLATFORM-M256-keyrole-nondeterminism` (iter-26), and it EXTENDS that record rather
than being explained by it.** iter-26 measured the card absent on 4 of 5 loads at occupancy **2** and
present 5 of 5 at occupancy **1**. This absence happened at occupancy **1** — so the instability is not
occupancy-2-exclusive, which is new evidence. Re-measured in isolation immediately afterwards: **6/6
passing**. Run 3 was also the slowest of the six (2.8 m vs 2.0 m), on a 9.7 GiB Docker VM against a
documented 12 GB floor, and the succession projection is a live `O(members)` recompute — a plausible
contributing factor, **stated as plausible and unmeasured**, not as a finding.

**No timeout was bumped.** iter-26 retracted exactly that "fix" when its diagnosis proved wrong, and a
timeout bump that appears to help is how a real cause gets buried.

So the honest gate statement is: **three consecutive clean cold runs (4–6) plus one earlier batch carrying
one diagnosed platform flake in three.** Clause 1's flake half is met on runs 4–6; the flake is on the
permanent record here and in the milestone ledger, and it is the coordinator's and the user's to weigh.

## Safety

Every write was to **demo-2's own Postgres** (reset-to-seed via `stackseed --reset`, the real FK-ordered
TRUNCATE path, N=0-guarded). Production was neither written nor read. The DRIFTED cockpit-manifest fixture
was restored and **sha-verified `99e2f315` after every one of the 10 resets this iter performed**.
`stackseed --policy-check --stack demo-2` rc 0 (`live=18 expected=18`); `datadna measure-closure` PASS (279
node-ids); `docker ps -a` 16 Up / 0 exited. The `pt-world` seed file and the Playthrough spec were each
backed up before their mutants and restored by `diff`-verified copy. **Zero platform-repo edits** — the
platform tree was read only.
