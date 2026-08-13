# iter-19 — decisions

## D91 — `pt-hiring-recruiter-compare` was not merely uncontrolled; its final was VACUOUS

Its final was `positionRows().count() > 0`, against the anchor `tbody.tbody > tr.tr:not(.empty)`.

**Measured (Phase A):** Org A's manager, on the **WORKFORCE** app's `/enterprise/activity-dashboard`,
renders **20 rows through that identical anchor** — several of them badged Hiring. A different org, a
different app, a different vantage, and a live in-demo page: the assertion is green there.

So the Playthrough proved *"some grid has rows"*. Finding the control and finding the defect turned out to be
**the same measurement** — which is the third time this milestone has hit that (iter-13 profiles, iter-14
workforce orgs, here). The rule is now general enough to state plainly: **when a Playthrough resists a
negative control, suspect the assertion before you suspect the world.**

**The sharpened final names three facts, and needs all three:**

| fact | why alone it is not enough |
|---|---|
| the recruiter's own org (`Kestrel Hiring Group`) | an org name alone does not say the *board* is right |
| exactly `HIRING_SHARED_POSITIONS` (5) rows | a count alone is satisfied by any 5-row grid |
| every row's type badge is Hiring, none other | type purity alone does not say *whose* board |

## D92 — The contrast vantage is a NON-HIRING TENANT'S ACTIVITY GRID, not the hiring view

`negative-controls.spec.ts`'s header already recorded the obvious vantage as **tried and rejected**: a
Workforce-org manager driven at `hiringAppBaseUrl` **ejects to production** (`app.anthropos.work/login`,
bodyLen 162). iter-19 re-measured it and confirmed — Phase A P2 landed on a real Clerk sign-in page. **That
recorded rejection is why this control was still open at 22**, and it was right to record it.

The vantage that works inverts the question: not *"what does a non-recruiter see on the hiring board?"* but
*"what does the hiring board's own locator set find on a non-hiring tenant's grid?"* Org A's manager on **her
own** workforce app is alive, in-demo, fully populated — and legitimately has no shared-position board. She
falsifies all three facts at once: `Meridian Labs` not `Kestrel`, 20 rows not 5, mixed types not pure.

**Liveness first, through the same accessors** (this file's standing rule since iter-07 D29): the control
asserts her grid really is producing rows *and* that some of them really do carry a Hiring badge — so the
type assertion is demonstrably being evaluated rather than matching nothing. The count is asserted `not 5`
rather than `= 20`: the claim is *"this is not a five-position board"*, and pinning 20 would make the control
brittle to the seeded content mix without making it stronger.

## D93 — `innerText` applies CSS `text-transform`; `textContent` does not — and Playwright reads `textContent`

The type badge **renders** `HIRING`. Its DOM text is **`Hiring`**. The span carries
`style="text-transform: uppercase"`, and:

- Phase A's probes read `innerText` → reported `HIRING`;
- the locator built from that reading was `getByText(/^HIRING$/)`;
- `getByText` matches **`textContent`**, which holds `Hiring`.

So the matcher could never match, and **both the sharpened final and its own control failed on it** on the
first live run. Fixed with case-insensitive matchers, and the reason is written into the page object rather
than left as a corrected line: *a styled string and a DOM string are different strings.*

Worth noting what caught it: not review, but **running it against a stack**. A `/^HIRING$/` locator is
perfectly plausible on the screen.

## D94 — `nonHiringTypeBadges()` is defined by EXCLUSION, not by enumeration

`hasNotText: /^hiring$/i` over the row type badges, rather than `/^(Assessment|Training|Interview)$/`. An
enumeration would silently stop discriminating the day a new simulation type shipped — an assertion that
quietly matches nothing, which is the defect class this milestone has found 17+ times. The exclusion form
cannot drift that way.

## D95 — `HIRING_SHARED_POSITIONS` is fenced against a Go constant in ANOTHER module

The `5` is not in `pt-world.seed.yaml` at all — it is `reservedHiringSimRefs` in
`stack-seeding/seeders/contentref.go:148`, a **code-owned** constant in a different rext module. Every other
fact in `seed-facts.ts` is reconciled against the seed YAML; this one had no possible YAML home.

Leaving it as a bare `5` would be the exact defect `seed-facts-fence.unit.spec.ts` exists to prevent
(*"a magic 8 in a spec is a claim about the seed with no link to it"*), and worse than the usual case:
**nothing in this module's own directory changes when it drifts.** So the fence now parses that Go file,
with a **fail-closed** guard — a rename that makes the regex miss fails loudly rather than skipping the
comparison. Both halves mutation-proven.

## D96 — The registry floor is raised to 23, and 23 is TERMINAL

`mutation-class-fence.unit.spec.ts`'s floor sat at `>= 20` while the count climbed 20 → 21 → 22, on the
deliberate reasoning that *"a fence edited on every increment would be edited without being read."* That
reasoning expires here: **23 is the maximum reachable.** The remaining two are the studio pair, held behind
`FIX-M256-studio-false-green` — a control over a known false green would certify it.

So the floor is raised to 23, and is now an equality in all but name. A drop below it means a control was
removed to make an edit easier.

## Mutants — 7 of 7 RED

| mutant | expected | result |
|---|---|---|
| N1 — count `5` → `20` in the final | RED | **RED** — "shared positions are exactly 5" |
| N2 — assert the CONTRAST org's name | RED | **RED** — "is THIS recruiter's tenant" |
| N3 — `hiringTypeBadges()` matches Assessment | RED | **RED** — "the board is type-PURE" |
| N4 — the control asserts the hiring org PRESENT | RED | **RED** — "must NOT appear on a different tenant's surface" |
| N5 — drift `HIRING_SHARED_POSITIONS` to 6 | RED | **RED** — "the seeder says 5" |
| N5b — rename the Go const in the source-of-truth pointer | RED (fail-closed) | **RED** — "must be findable … passes vacuously" |
| N6 — remove the `@pt-control-for` line | RED (floor) | **RED** — registry 23 → 22, "must not regress" |

All four mutated files restored **byte-identical**: `93d0e3c76913` / `3b697837ed49` / `348e32ffccbd` /
`b25aaad4e097`.
