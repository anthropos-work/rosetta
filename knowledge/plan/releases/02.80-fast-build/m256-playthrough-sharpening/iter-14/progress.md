**Type:** tik · shape: standard (single planned target, four Playthroughs of one cluster)

# iter-14 — the four workforce finals were structural about an ORG, not a person

## Phase A — probe first (three passes, and the third refuted the second)

Three passes on `demo-2`, both manager seats, all five candidate surfaces. The design decisions below are
all measurements, not preferences.

**Pass 1** established which needles discriminate at all. **Pass 2** added the question pass 1 could not
answer — *is the hit inside `<main>` (the outcome) or in the chrome?* — and produced the table in
[`overview.md`](overview.md). Two of its readings shaped the build, and one of them was **wrong**:

- **The org NAME is chrome.** `Meridian Labs` reads `main=0 body=1` — it is the sidebar org switcher. The
  most obvious anchor on all five surfaces, and it would have proved *the switcher renders a label*, not
  that the aggregate below it is this tenant's. **Rejected on the measurement.**
- **`Overall Members` = 40 on BOTH orgs** (both seeded `size: 40`). So the magnitude is a genuine
  strengthening — it catches a stat card rendering a constant, and an accessor reading the wrong card,
  neither of which a visible LABEL can ever catch — but it is **not** a discriminator, and the spec and the
  control both say so in as many words rather than letting a reader infer one from the other.
- **Pass 2's "Org C's activity dashboard is empty" was a PROBE ARTIFACT.** It read 20 table rows with **no
  cell content** (`mainLen 469`, every row all-whitespace) and I nearly banked that as a real absence —
  a free contrast vantage for `pt-activity-drilldown`. Pass 3 watched the same grid for 90 s with a
  *content-aware* settle and it populates within 3 s. The pass-2 settle condition was `rows > 5`, **which
  the skeleton satisfies** — the probe stopped looking at the exact moment the answer was still wrong.
  *A settle predicate that the empty state satisfies is not a settle predicate.*

**What survived that, and is worth more than the artifact:** the skeleton state is real, and
`pt-activity-drilldown` asserts `contentRows().first()` visible then `count() > 0` — **both of which the
20 content-free rows satisfy**. It is not a false green (the drill step that follows needs a row `<a>`,
which a skeleton has not got, so the Playthrough as a whole still fails), but it is a weak assertion with a
measured witness. Routed to iter-15 with the measurement attached.

**Pass 3 also handed iter-15 its discriminator free:** drilling the same content id on both vantages gives
Org A `Pat Ellis / DevOps Engineer` + `Morgan Reyes / Engineering Manager` (2 rows) and Org C `Mei Costa /
Advanced Civil Engineer` + `Theo Lindqvist / Business Operations Analyst` (5 rows). The per-member
breakdown names members **with their roles**, and the first content row is deterministically a
hero-session one (the grid sorts by most-recent and the heroes' sessions are seeded at reset, i.e. today).

## Phase B — the sharpening

**The generalisation iter-13 did not have.** iter-13's move was *"re-aim a structural final at the hero's
own seeded data"*. These four finals are structural too, but one level up: they are about an **org's
aggregates**, and the question they were failing to answer is not *whose profile is this* but **whose
TENANT is this**. That changes the contrast vantage — `pt-manager` cannot falsify her own dashboard — and
it changes what a spec is allowed to name. So `seed-facts.ts` gained a `SeededOrg` (`story id`, `org.name`,
`org.slug`, `org.size`) alongside its heroes, and the fence gained a `parseSeedOrgs` reconciliation.

| Playthrough | old final (now an intermediate) | new final |
|---|---|---|
| `pt-workforce-roster` | the Role column renders · rows > 0 | rows in the org's own email domain > 0 · the seeded hero's row carries her seeded job role |
| `pt-workforce-funnel` | 2 stat LABELS · a `% of mapped` string · a drawn `<svg>` | `Overall Members` == the seeded `org.size` · the seeded hero is spotlighted with her seeded role |
| `pt-workforce-succession` | `/ready/i` · `/at.?risk/i` · rows > 0 | the seeded role has a computed key-role card · the seeded hero has a row carrying that role |
| `pt-workforce-org-feedback` | a recap label · both polarity words · rows > 0 | rows in the org's own email domain · the seeded hero among them · **and still in-domain after the filter** |

**The anchor that made three of the four cheap, and why it is the right one:** the **org email domain**
(`@pt-meridian-labs.com`, derived from the seed's `org.slug`) is on **every** org-member row. The roster
renders 20 of 40 members, so any single member's row is a bet on sort order; the domain is not. The bet is
survivable here only because the grid sorts alphabetically by surname and the hero's is "Ellis" (measured:
Andersen, Becker, Costa, Dubois, **Ellis**, Esposito, …) — that is written into the spec as the reason,
not left as luck. Honestly bounded: **15 of 20** rows carry it, the other five being `Candidate`-role
members on external addresses (`theo.becker32@fastmail.com`), so the assertion is *present among the rows*,
never *all rows*.

**One accessor was written from the wrong model and the live run said so.** `readOrgStat` was drafted as a
copy of `ProfilePage.readSkillStat` — match the label, parse the number out of the same element — and
returned `null` against a dashboard that plainly renders `40`. Measured (pass 4): the label is its own
`<span>` with **no number in it**, and the card is its parent, `textContent` = `"40Overall Members40 active"`
— **value first, label second, a second number after it**, and *no whitespace anywhere*, because
`textContent` concatenates. So the pattern is `^\d+Overall Members` with nothing between, the parse comes
off `innerText` (where the newline makes the boundary unambiguous), and the comment records both halves.
This is iter-13 D63's trap arriving through a new door: the same constant is safe under `getByText` and
broken under `hasText`, and *the shape of the DOM is a measurement, not an inference from a sibling file*.

## Phase C — the controls, and every one watched going RED

One new batched test in `negative-controls.spec.ts` on **one** `pt-ai-manager` login (Org C's manager),
four absences, each paired with **her org's OWN equivalent asserted PRESENT through the SAME accessor**
(iter-13 D64). Nothing on her side is broken: 40 members, 109 skills mapped, 88 verified, 182 simulations,
a full roster, a complete succession projection, 20 feedback rows. Every one of the four surfaces works
completely for her. The only thing legitimately absent is Org A's data — which is what makes the absence
evidence rather than a coincidence.

**11 mutants, 11 RED** (harness: `$SCRATCH/iter14/mutants.sh`; `cp` backups, never `git checkout`; all five
files verified byte-identical after):

| group | mutant | result |
|---|---|---|
| A1–A4 | each sharpened Playthrough driven on the contrast tenant (`HERO_SEAT` → `pt-ai-manager`) | **RED** ×4 |
| B1 | roster org-domain absence re-aimed at Org C's domain | **RED** |
| B2 | roster hero-row absence re-aimed at Org C's own manager | **RED** |
| B3 | funnel spotlight absence re-aimed at Org C's thriving hero | **RED** |
| B4 | succession key-role absence re-aimed at Org C's seeded role | **RED** |
| B5 | succession talent-row absence re-aimed at Org C's hero | **RED** |
| B6 | feedback org-domain absence re-aimed at Org C's domain | **RED** |
| B7 | feedback hero-row absence re-aimed at Org C's hero | **RED** |

Group A is the half that matters most: it proves the **finals** discriminate, not merely that the controls
can fire. A control over a still-structural final would be a green test about a green test.

**And the seed link, mutated three ways** (seed backed up, restored, sha verified `dc8a58ed`): `size: 40 →
41` → RED · `slug` renamed → RED · every `org:` key renamed so the parse finds nothing → **4** failures
including the fail-closed non-vacuous test. That last one is the one worth having: a reconciliation over an
empty parse passes every comparison silently, which is this milestone's signature defect.

## Phase D — re-measure

**Gate: 3 consecutive cold reset-to-seed runs, `170 passed`, rc 0, 0 flake.**

| run | result | `ptreport` | wall |
|---|---|---|---|
| 1 | **170 passed**, rc 0 | 24 passing / 0 failing / 7 TODO / 0 unimplementable | 1.2 m |
| 2 | **170 passed**, rc 0 | same | 1.2 m |
| 3 | **170 passed**, rc 0 | same | 1.3 m |

- `@pt-negative-control` registry, **computed by the fence: 20 of 24** (8 self-declared + 12 via the control
  spec). Named uncovered: `pt-activity-drilldown`, `pt-hiring-recruiter-compare` (both iter-15) + the 2
  studio, still blocked behind `FIX-M256-studio-false-green`.
- `@pt-mutation` registry, computed: **MUTATES=6 READ-ONLY=16 UNKNOWN=2**.
- `ptvalidate`: **VALID** — 10 products, 31 use cases, 24 live Playthroughs, 7 TODO.
- Unit fences: **142 passed**. `tsc --noEmit`: clean. `gofmt -l` over all six rext sections: clean.
- Go: all **6** modules rc=0, **0 FAIL**.
- Python, one invocation per suite, rc captured into a variable (never off a pipe): `demo-stack` **999
  passed / 1 skipped**, `stack-core` **287 passed**, `stack-verify` **171 passed**, `stack-injection` **266
  passed / 1 skipped** — **1723 passed, 0 failed**.
- The deliberately DRIFTED `demo-2` cockpit-manifest fixture was backed up before the first `--reset` and
  restored after the last: sha **`99e2f315`**, verified.

### Reported (never gated) — clause 1's suite statistic, per D-v28-13

n=3, same host, cold reset-to-seed each: median per non-studio Playthrough **2.100 s = 0.6314×** of the
3.326 s baseline; per-run **0.6013× / 0.6164× / 0.6464×**, a **1.075× range**. The untouched ORIGINAL-16
control subset reads **0.5562×**.

**The four sharpened Playthroughs cost nothing measurable** — 1.4–1.9 s each across the three runs, all
*below* the suite median, so a tenancy proof is not more expensive than a structural one. That is worth
recording because iter-06 measured the opposite trade for a *write* (proving a write costs more than
proving a render); sharpening a read does not.

**Still evidence for the D-v28-13 recut.** The control subset — code no iter has touched since iter-03 —
has now been observed at **0.5281× · 1.0762× · 0.7517× · 0.9321× · 0.5562×** across five batches on one
host. A gate at 0.79× sits inside that. **This iter landed no speed mechanism, so clause 1's leg half has
nothing new to measure**; its flake half is **MET** (0 flake ×3).

**Environment:** `Kirality-Mac-Pro-6.local`, darwin 25.1.0, Docker VM 9.70 GiB (vs the documented 12 GB
floor); `demo-2` offset 20000, localhost/http, `--no-public-host`. Per D-v28-12 no number here is
comparable to billion's 228 s.

## Close — 2026-07-29

**Outcome:** **negative controls 16 → 20 of 24.** The four workforce finals were structural *about an org*,
so they were satisfied by any populated tenant — measured on a second seeded org that renders all four
surfaces completely. Re-aimed at the seed's own org facts (its email domain, its member magnitude, its
seeded hero with her seeded role), each of the four now fails on that other tenant, and each has a control
that fails when re-aimed at that tenant's own data. **11 mutants RED.** Along the way the probe's own
settle predicate was caught certifying a hydrating grid as populated — the artifact was discarded, the
weak assertion it exposed was kept and routed.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — clause 2 mutating **6/5 MET**, `blocked` **1/1 MET**, negative controls **20 of 24**
(4 remaining: `pt-activity-drilldown` + `pt-hiring-recruiter-compare`, both measured and routed to iter-15,
+ 2 studio blocked behind `FIX-M256-studio-false-green`); clause 3 verdict half **COMPLETE (28/28, 0
unimplementable)**, landed half still short (org-admin 2/4, onboarding 1/5); clause 1 leg half **N/A this
iter** (no speed mechanism landed), flake half **MET — 0 flake ×3**; **D-v28-5 still unfixed** (unblocked,
not started).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this tik landed its planned scope and moved the primary metric 16→20) — (3) re-scope: n (0 of 31 curated UCs `unimplementable`; the mechanism reached one product further, not one fewer) — (4) user-blocker: n (the one live failure was my own accessor written from the wrong DOM model, measured and fixed inside the iter) — (5) cap-reached: n (1st tik of this invocation) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D68 (a structural final about an ORG needs a second seeded TENANT as its contrast, not a second person — the vantage class follows the *subject* of the final), D69 (prefer a per-row org anchor over a single hero row: the domain is pagination-proof where a row is a bet on sort order), D70 (a magnitude that is equal on both vantages is a strengthening, not a discriminator — say which it is at the assertion, or a reader will bank the wrong one), D71 (**a settle predicate the empty state satisfies is not a settle predicate** — the pass-2 probe certified a hydrating grid as populated and nearly banked a permanent absence that does not exist), D72 (the DOM shape of a stat card is a measurement, not an inference from a sibling page object — `readOrgStat` copied `readSkillStat`'s label-carries-the-number model and returned null against a live page), D73 (mutate the FINAL as well as the control: driving each sharpened Playthrough on the contrast vantage is what proves the final discriminates, where mutating only the control proves a green test about a green test).
**Side-deliverables:** none — every change served the planned target.
**Routes carried forward:**
- `NEGCTL-M256-cross-vantage` → **20 of 24; 4 remain.** `pt-activity-drilldown` and
  `pt-hiring-recruiter-compare` are next, and iter-14's Phase A has already priced both:
  - **`pt-activity-drilldown`** — two pieces. (a) Its grid assertions (`contentRows().first()` visible,
    `count() > 0`) are satisfied by **20 content-free skeleton rows**, measured; assert *contentful* rows
    instead. (b) Its control comes from the **drill-down**, which names members WITH their roles: Org A's
    first content drills to `Pat Ellis / DevOps Engineer`, the SAME content id on Org C drills to
    `Mei Costa` / `Theo Lindqvist`. The first row is deterministically a hero-session one (the grid sorts
    by most-recent; hero sessions are seeded at reset).
  - **`pt-hiring-recruiter-compare`** — still needs a **same-vantage** control (iter-12 measured that its
    contrast vantage ejects the browser to production). Untried and cheap to check: Org D's authored facts
    are `size: 40`, `role_mix 0.1 admin / 0.9 candidate` → **5 shared positions / ~36 candidates**, so a
    cardinality final pinned to the seed is available; where the *absence* comes from is the open question,
    and the honest fallback is a written verdict rather than a manufactured one.
  - The 2 studio remain blocked behind `FIX-M256-studio-false-green`.
- `PT-M256-readiness-step-asserts` → **still open, still unstarted.** Same defect shape as this iter's work
  (an assertion satisfied by the wrong state); it should ride with the iter-15 batch.
- `MEASURE-M256-clause1-sampling` → more evidence, same escalation. Control subset now observed at
  0.5562× as a fifth batch reading.
- Everything else from iter-13's list stands unchanged.
**Lessons:**
1. **The contrast vantage follows the SUBJECT of the final, not the product.** iter-13 found the contrast
   for a *person* final: a different person. Reusing that seat here would have failed — `pt-manager` IS the
   org whose aggregates these four assert. An *org* final needs a second seeded TENANT. Stating the rule
   that way generalises it: before hunting a vantage, name what the final is about.
2. **Mutate the final, not only the control.** Group B alone would have shown eight absence assertions
   capable of firing while the Playthroughs they cover stayed vacuous. Group A — each Playthrough driven on
   the contrast tenant — is the assertion that the *sharpening* worked. A control over an unsharpened final
   certifies nothing, which is exactly what iter-12 refused to ship.
3. **A probe's settle condition is part of the measurement.** Pass 2 asked "are there rows?" and got yes
   from a skeleton, and the conclusion it supported — a permanently empty grid, a free contrast vantage —
   was wrong in the direction that would have shipped. This is the same failure mode as an assertion that
   cannot fail, committed in the instrument instead of the test.
