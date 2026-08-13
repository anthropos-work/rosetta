**Type:** tik · shape: standard (single target: the last sharpenable negative control)

# iter-19 — negative controls 22 → 23 of 25, and the Playthrough was proving nothing

## Phase A — measure both vantages before writing an assertion

| probe | result |
|---|---|
| the recruiter's scoreboard: what seeded facts does it name? | org `Kestrel Hiring Group`; **exactly 5** rows; every row a `Hiring`-badged position; "5 AI Simulations across 30 users" |
| a Workforce manager on the HIRING base (the vantage the control file records as REJECTED) | **confirmed rejected** — lands on a real production Clerk sign-in |
| a Workforce manager on **her own** activity dashboard | **20 rows through the IDENTICAL `tbody.tbody > tr.tr` anchor**, several badged Hiring, org reads `Meridian Labs`, types MIXED |

**That third reading is the iter.** `pt-hiring-recruiter-compare`'s final was `positionRows().count() > 0`
— green on a different org, a different app and a different vantage. **Finding the control and finding the
defect were the same measurement** (D91), for the third time in this milestone (iter-13, iter-14, here).

## Phase B — sharpen the final, then build the control against it

- **The final now names three facts** (D91): the recruiter's own org identity, exactly
  `HIRING_SHARED_POSITIONS` rows, and **type purity** (every row Hiring, none Assessment/Training).
- **`seed-facts.ts` gains `PT_ORG_D`** — reconciled against the seed YAML by the existing loop — **plus
  `HIRING_SHARED_POSITIONS`**, which has no YAML home at all: it is `reservedHiringSimRefs` in
  `stack-seeding/seeders/contentref.go`, a code-owned constant **in another rext module**. The fence now
  parses that Go file, fail-closed (D95).
- **The control** (D92) inverts the rejected vantage's question: not "what does a non-recruiter see on the
  hiring board" but "what does the hiring board's own locator set find on a **non-hiring tenant's** grid".
  Org A's manager on her own app is alive, populated, and falsifies all three facts at once. Liveness is
  asserted first through the same accessors — including that some of her rows *do* carry a Hiring badge, so
  the type assertion is demonstrably being evaluated.
- **`nonHiringTypeBadges()` is defined by EXCLUSION** (`hasNotText: /^hiring$/i`), not by enumerating the
  other type names, which would stop discriminating the day a new type shipped (D94).
- **The registry floor is raised 20 → 23** (D96), because **23 is terminal**: the remaining two are the
  studio pair held behind `FIX-M256-studio-false-green`.

**One defect of my own, caught by the stack rather than by review (D93):** the badge *renders* `HIRING` but
its DOM text is `Hiring` — the span carries CSS `text-transform: uppercase`. Phase A's probes read
`innerText`, which applies the transform; Playwright's `getByText` matches `textContent`, which does not. The
first live run failed **both** the sharpened final and its own control on a `/^HIRING$/` matcher that could
never match. *A styled string and a DOM string are different strings*, now written into the page object.

## Phase C — mutants: 7 of 7 RED

N1 count · N2 org identity · N3 type purity · N4 the control flipped to presence · N5 the Go-const drift ·
N5b the fail-closed parse guard · N6 the registry floor. Table + messages in [`decisions.md`](decisions.md).
All four mutated files restored **byte-identical**.

## Phase D — re-measure

| run | result | `ptreport` | wall |
|---|---|---|---|
| 1 | **181 passed**, rc 0 | **25** passing / **0** failing / 6 TODO / 0 unimplementable | 1.3 m |
| 2 | **181 passed**, rc 0 | same | 1.2 m |
| 3 | **181 passed**, rc 0 | same | 1.2 m |

3 consecutive **cold reset-to-seed**, **0 flake**. 181 = iter-18's 178 + 2 new seed-fence tests + 1 control.
`@pt-negative-control` registry, computed: **23 of 25** (9 self-declared, 14 via control spec); still
uncovered: `pt-studio-advanced-generate`, `pt-studio-guided-generate` — **both correctly withheld**.
`ptvalidate` **VALID** (10 products, 31 use cases, 25 live, 6 TODO). `playthroughs` Go module 0 FAIL.
`gofmt -l` clean. `tsc --noEmit` clean. Drifted cockpit fixture restored + sha-verified **`99e2f315`**.

## Close — 2026-07-29

**Outcome:** **negative controls 22 → 23 of 25 — the terminal value** — by discovering that the last
uncontrolled non-studio Playthrough was **not merely uncontrolled but vacuous**: its
`positionRows().count() > 0` final is satisfied by 20 rows on a *different tenant's, different app's* grid
through the identical anchor. Re-aimed at three facts together (org identity · exactly the seeder's 5 shared
positions · type purity), a non-hiring tenant's activity grid falsifies all three. The floor is raised to 23
and stated as terminal, since the remaining two are correctly held behind a known false green.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — clause 2 mutating **7** (≥5 MET), `blocked` **1/1 MET**, negative controls **23 of 25**
(**terminal**: the 2 remaining are studio-blocked behind `FIX-M256-studio-false-green`); clause 3 verdict half
**COMPLETE**, landed half **org-admin 3 of 4**, **onboarding 1 of 5**, 0 `unimplementable`; clause 1 leg half
**N/A** (no speed mechanism), flake half **MET** (181×3 cold, 0 flake); **D-v28-5** part (b) open.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this tik moved the metric) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2nd tik of this invocation) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D91 (the final was VACUOUS — 20 rows on the same anchor for another tenant; *when a
Playthrough resists a control, suspect the assertion before the world*), D92 (the contrast vantage is a
non-hiring tenant's **own** grid — the hiring-base vantage stays rejected, re-confirmed), D93 (**`innerText`
applies CSS `text-transform`, `textContent` does not, and Playwright reads `textContent`** — a `/^HIRING$/`
matcher that could never match, caught by the stack), D94 (the complement locator is defined by exclusion,
never by enumeration), D95 (a seed fact whose home is a **Go constant in another module** is fenced against
that file, fail-closed), D96 (the floor is raised to 23 **and 23 is terminal**).
**Side-deliverables:** none — all in planned scope.
**Routes carried forward:**
- `NEGCTL-M256-studio-pair` → the last **2** uncovered are `pt-studio-{advanced,guided}-generate`, and they
  stay uncovered **on purpose** until `FIX-M256-studio-false-green` lands. A control over a known false green
  would certify the false green. **This is a written disposition, not a gap** — clause 2's control
  sub-clause is at its terminal value.
- `FIX-M256-studio-false-green` → now the ONLY thing standing between clause 2 and 25 of 25. Promoted to the
  strongest remaining clause-2 target.
- All of iter-18's onboarding routes stand unchanged (`ONBOARD-M256-seat-append` first —
  it is the prerequisite for the rest), plus `PT-M256-resume-fixture-pair` and
  `PT-M256-orgadmin-role-create` (the last org-admin UC).
- `D-v28-5-cockpit-logout` → part (b) still needs a `fake-fapi` re-pin + rebuild.
**Lessons:**
1. **When a Playthrough resists a negative control, suspect the ASSERTION before you suspect the world.**
   Third time in this milestone. iter-12 concluded three profile finals were uncontrollable; iter-13 refuted
   it. iter-14 took the same move up to the tenant. Here, four iters of "needs a same-vantage control" was
   really "its final is structural" — and the vantage that exposed it was the plainest one available.
2. **A styled string and a DOM string are different strings.** `innerText` applies CSS `text-transform`;
   `textContent` — what Playwright matches — does not. A probe that reads the wrong one hands you a locator
   that is plausible on screen and impossible in the DOM.
3. **A recorded rejection is worth keeping even when it later turns out to be the wrong question.** The
   header's "this vantage ejects to production" was correct, and it was also why the control stayed open for
   four iters. What unlocked it was not overturning the rejection but *inverting the question it answered*.
4. **Define a complement by exclusion.** An enumerated "everything else" is an assertion with an expiry date.
