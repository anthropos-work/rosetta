# iter-14 — decisions

## D68 — a structural final about an ORG needs a second seeded TENANT as its contrast

iter-13 established the move (*re-aim a structural final at seeded data, then a contrast vantage falsifies
it*) and, implicitly, a vantage: a **different person in the same org**. That seat cannot work here.
`pt-workforce-{roster,funnel,succession,org-feedback}` are read BY `pt-manager` and are ABOUT her org's
aggregates — she cannot falsify her own dashboard.

**The rule, generalised: the contrast vantage follows the SUBJECT of the final.** A final about a person →
another person. A final about a tenant's aggregate → another **tenant**. Named before hunting, this is a
one-line decision; discovered afterwards it is a wasted iter.

`pt-ai-manager` (Org C, Vertex Logistics) qualified on measurement, not convenience: her org renders all
four surfaces completely (40 members, 109 mapped, 88 verified, 182 sims, 20 roster rows, a full succession
projection, 20 feedback rows). **Nothing is broken or empty on her side** — the only absence is Org A's
data, which is the property iter-07 D29 established as the difference between a control and a dead page.

## D69 — prefer a per-row ORG anchor over a single hero row

`/enterprise/members` renders **20 of 40** members. A hero-row assertion is therefore a bet on sort order,
survivable here only because the grid sorts alphabetically by surname and the hero's is "Ellis" (measured:
Andersen, Becker, Costa, Dubois, **Ellis**, Esposito, …). The **org email domain** (`@<org.slug>.com`) is on
every org-member row, so it is pagination-proof — it holds on any page of this org's roster.

Both are asserted, deliberately of different shapes: the domain is the robust tenancy claim, the hero row is
the specific one (it proves the grid reached a person the seed pins by name, not 20 rows of generated fill).
Bounded honestly: **15 of 20** rows carry the domain; the other five are `Candidate`-role members on
external addresses, which is real seeder behaviour. So the assertion is *present among the rows*, never
*all rows* — a stricter version would be a false RED waiting for a seed that already exists.

⚠ The `.com` suffix is a **seeder convention, not a seed field** (the YAML pins the slug; hero logins use
`.test`). The fence reconciles the slug and cannot reconcile the suffix. Flagged at the helper so a future
reader does not mistake a measured fact for a seed-derived one.

## D70 — say whether a magnitude is a strengthening or a discriminator, at the assertion

`Overall Members` renders **40** on both seeded orgs, because both are `size: 40`. So asserting it equals
the seeded size is a real improvement — it catches a stat card rendering a constant and an accessor reading
the wrong card, neither of which a visible LABEL can catch — but it **cannot** discriminate tenants.

A reader who found `expect(readOrgStat(...)).toBe(PT_ORG_A.size)` next to a control would reasonably assume
it was the discriminator. So both the spec and the control state which it is in as many words. This is the
smaller sibling of the milestone's central defect: not a check that cannot fail, but a check whose *reason
for existing* could be mis-read as stronger than it is.

## D71 — a settle predicate the empty state satisfies is not a settle predicate

Phase A pass 2 settled on `main table tbody tr > 5` and concluded that Org C's activity dashboard renders
**20 rows with no cell content** (`mainLen 469`). That reading was banked for long enough to design around:
a permanently empty grid would have been a *free* contrast vantage for `pt-activity-drilldown`.

Pass 3 watched the same grid with a **content-aware** predicate and it populates within 3 s. The skeleton
satisfied the pass-2 predicate, so the probe stopped looking while the answer was still wrong.

**The instrument is part of the measurement.** This is the identical failure mode to an assertion that
cannot fail — the class this milestone has now fixed 17 times in tests — committed in the probe instead.
Any settle predicate must be false in the state it is waiting to leave.

**What survived the retraction, and is real:** `pt-activity-drilldown` asserts `contentRows().first()`
visible then `count() > 0`, and **the 20 content-free rows satisfy both**. Not a false green (the drill step
that follows needs a row `<a>`, which the skeleton lacks), but a weak assertion with a measured witness.
Routed to iter-15.

## D72 — the DOM shape of a stat card is a measurement, not an inference from a sibling page object

`readOrgStat` was drafted as a copy of `ProfilePage.readSkillStat` — *match the label, parse the number out
of the same element* — and returned `null` against a dashboard that plainly renders `40`. The Playthrough
failed on a stat that works.

Measured (pass 4): the label is its own `<span>` carrying **no number**; the card is its parent, with
`textContent` `"40Overall Members40 active"` — **value first, label second, a second number after it, and no
whitespace anywhere**, because `textContent` concatenates sibling nodes. So the locator has to be a
`filter({ hasText: /^\d+Overall Members/ })` with nothing between digits and label, and the value is parsed
off `innerText` where the newline makes the boundary unambiguous.

This is iter-13 D63 arriving through a new door — *the same constant is safe under `getByText` and broken
under `hasText`* — and the general form is: **two stat cards in the same app can have opposite shapes.**
Read the DOM; do not infer it from the file next door.

## D73 — mutate the FINAL as well as the control

The mutation set has two groups on purpose. Group B re-aims each control's absence assertion at the contrast
tenant's own data and requires RED — that proves the **control** can fire. Group A drives each **sharpened
Playthrough** on the contrast tenant and requires RED — that proves the **final discriminates**.

Group B alone would have demonstrated eight absence assertions capable of firing over four Playthroughs that
were still structural, i.e. a green test about a green test. That is precisely the outcome iter-12 refused
to ship when it declined to write contrast controls for structural finals, and Group A is the cheap
mechanical guard against re-introducing it.

11 mutants, 11 RED. Harness at `$SCRATCH/iter14/mutants.sh`; `cp` backups only (the standing `git checkout`
ban), all five touched files verified byte-identical afterwards, and the seed verified back to sha
`dc8a58ed`.
