**Type:** tik · shape: standard (single target: the onboarding cluster)

# iter-18 — the blockers were wrong, the journey works, and it still must not ship

## Phase A — drive all four recorded blockers (6 probe passes)

Onboarding's four TODO UCs each carried a written blocker from iter-08, and **none had been driven** — iter-08's
Phase A only ever clicked `Skip`. Six passes on `demo-2`:

| pass | question | answer |
|---|---|---|
| 1 | is a non-`pt-free` seat day-0? what does the LinkedIn offer look like? | **every seat is day-0** (Org A / C / D all probed). And there is a `#linkedinUrl` textbox — **no fixture needed for LinkedIn** |
| 2 | fill the URL, click forward | `Next` **stays disabled**, 4.1 min, no import fires |
| 3 | why? dump the form | input value is set, `aria-required`, **no radios** — and the post-typing `/^Next$/` filter returns `[]` |
| 4 | dump every button's label | **`Next` → `Import`. The control RELABELS.** Clicking it runs a real import: counter `5 → 8 → 50`, preview fills (3 work experiences, 1 education, 51 yrs), forward enables ~15 s |
| 5 | the deterministic route: a checked-in CV | attaches, `200 POST /api/resources/resume`, then **nothing** — counter `0`, forward disabled 6.9 min |
| 6 | is the advance scrape-dependent? is there a hidden enabled control? | **advance is scrape-INDEPENDENT** (a non-resolving URL advances identically and never enables — 120 s). CV route: **no hidden enabled control either**, both formats |

Two of my own errors, both caught by measurement: pass 3's `isVisible({timeout})` **takes no timeout and does
not wait**, so a 1.4 s run read "seat consumed" against a healthy stack; and pass 5's probe locked itself out of
`pt-employee` by completing the flow on the previous pass — which is how **D89** got demonstrated rather than
argued.

## Phase B — implement what A proved, and refuse what it did not

**No Playthrough shipped.** The only working route scrapes linkedin.com; the deliverable is the measurement
plus the assets (**D88**). Landed:

- `fixtures/` — **reserved by spec §5.4 since M202 and empty until now.** A real PDF + a `.docx` of a wholly
  invented person (RFC-2606 `.example` domain; employers/school that occur nowhere in the seed or taxonomy, so
  an assertion naming one can only be satisfied by *that file having been imported*) + a README stating the
  synthetic-only rule and the measured CV-route stall.
- `e2e/lib/onboarding-page.ts` — `linkedinUrlInput()`, `resumeFileInput()`, `importButton()`,
  **`forwardControl()`** (intent-named, spans the relabel); `nextButton()` and `uploadButton()` re-documented
  to say what they actually match and when.
- `e2e/drafts/onboarding-self-import.spec.ts.draft` — the working journey, its **same-surface negative
  control**, and the refusal argument, so the next attempt starts from evidence.
- `manifest/onboarding.yaml` — UC1's verdict **rewritten from measurement** (fixture premise refuted; the real
  blocker named with its endpoint; the seat rule stated). UC2 / ai-readiness / hiring verdicts **narrowed**: each
  had a "needs a day-0 seat" half that is already true, and the hiring one gained a trap it must settle first
  (the cockpit already routes an `is_hiring` hero to the hiring base, so a naive final would prove the
  **cockpit's** routing, not onboarding's).
- `e2e/tests/onboarding-locators.unit.spec.ts` — 5 tests fencing the relabel, the anchoring, the measured
  selectors, and the fixture ledger both ways.

## Phase C — mutants: 5 of 5 RED

M1 relabel-narrowing · M2 promised-but-absent fixture · M3 undocumented fixture · M4 unanchored label ·
M5 fail-closed capture floor. Full table + messages in [`decisions.md`](decisions.md). Page object and README
restored **byte-identical** (`aa6abc52a580` / `c20970bdf49b`).

**D90:** test 5's own fail-closed floor fired on first run — the scan it was built around found nothing,
because no shipped Playthrough uploads a file. Re-aimed at the README's file table, reconciled both ways.

## Phase D — re-measure

| run | result | `ptreport` | wall |
|---|---|---|---|
| 1 | **178 passed**, rc 0 | **25** passing / **0** failing / 6 TODO / 0 unimplementable | 1.4 m |
| 2 | **178 passed**, rc 0 | same | 1.2 m |
| 3 | **178 passed**, rc 0 | same | 1.2 m |

3 consecutive **cold reset-to-seed**, **0 flake**. 178 = iter-17's 173 + the 5 new fence tests.
`ptvalidate`: **VALID** — 10 products, 31 use cases, 25 live, 6 TODO (unchanged; no UC landed).
Controls **22 of 25**, unchanged. `playthroughs` Go module rc 0. `gofmt -l` clean. `tsc --noEmit` clean.
Drifted `demo-2` cockpit fixture restored + sha-verified **`99e2f315`** (12 advertised).

## Close — 2026-07-29

**Outcome:** **no onboarding UC landed — and the four recorded blockers are now measured instead of asserted.**
The headline is a refusal: `onboarding.enterprise-workforce-standard.UC1` **works end to end** (its recorded
"needs a résumé fixture" blocker is false — LinkedIn needs none — and the import really runs on a demo in
~15 s, arriving with a same-surface negative control), and it is **deliberately not shipped**, because its
green depends on scraping a site that blocks automation and its RED would read as a product regression.
The deterministic alternative is blocked upstream of the fixture that was blamed for it: the CV upload **200s**
and the parse never starts.
**Type:** tik
**Status:** closed-no-lift
**Gate:** NOT MET — clause 2 mutating **7** (≥5 MET), `blocked` **1/1 MET**, negative controls **22 of 25**
(unchanged); clause 3 verdict half **COMPLETE**, landed half **org-admin 3 of 4**, **onboarding 1 of 5**
(unchanged), 0 `unimplementable`; clause 1 leg half **N/A** (no speed mechanism landed), flake half **MET**
(178×3 cold, 0 flake); **D-v28-5** half-fixed (iter-16), part (b) open.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (1 no-prog tik; the streak needs 3) — (3) re-scope: n (0 `unimplementable`; nothing was declared unbuildable) — (4) user-blocker: n — (5) cap-reached: n (1st tik of this invocation) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D86 (the control RELABELS `Next`→`Import`; locate by intent, and dump every label before
declaring a path closed), D87 (the CV route is a **product-defect candidate** — `200 POST
/api/resources/resume` then a parse that never starts, on both formats, hidden nodes included), D88 (the
working journey **refused on P6**; the test is *whose refusal produces the RED*, not whether the network is
touched), D89 (**every seat is day-0** and the first Playthrough to drive onboarding **consumes** it —
demonstrated by this iter's own probe; a second one needs a seat **appended**, because `personaUserIndexFor`
indexes by declaration order), D90 (the fence's fail-closed floor caught the fence).
**Side-deliverables:** none — every artifact above was in the iter's planned scope.
**Routes carried forward:**
- `PT-M256-resume-fixture-pair` → **content changed, handler unchanged.** No longer "check in a fixture" (done)
  but "get the résumé parse running on a demo". One fix lands **two** UCs:
  `onboarding.enterprise-workforce-standard.UC1` + `profile-skills.import.UC1` (A2). **Report the product
  defect** (D87).
- `ONBOARD-M256-seat-append` → **NEW, and the prerequisite for every remaining onboarding UC.** Append a day-0
  hero to `pt-world.seed.yaml` (append only — D89) + register the roster seat + lockstep
  `seed-facts.ts::SEEDED_HEROES`. Org B (`size: 20`, narrative `onboarding-ramp`, unfenced in `SEEDED_ORGS`) is
  the lowest-risk host.
- `ONBOARD-M256-hiring-discriminator` → settle the trap in the hiring UC's note **before** building it: the
  cockpit already routes an `is_hiring` hero to the hiring base, so identify what distinguishes onboarding's
  routing from the cockpit's, or the UC lands a green over the wrong mechanism.
- `ONBOARD-M256-ai-readiness-stage0` → only the **stage-0 funnel** is missing now (the day-0 half is already
  true); an appended Org C seat with no funnel signals.
- `ONBOARD-M256-uc2-trigger` → UC2's trigger is narrowed but unidentified: heroes across **four** orgs are all
  served the same import form, so it is neither "has a profile" nor an org narrative flag in this seed.
- `NEGCTL-M256-cross-vantage` → **22 of 25**, and the strongest remaining clause-2 target:
  `pt-hiring-recruiter-compare` needs a same-vantage control. **Deliberately routed to iter-19** rather than
  opened here (one line of investigation per iter). The other 2 stay untouched behind
  `FIX-M256-studio-false-green` — a control over a known false green would certify it.
- `D-v28-5-cockpit-logout` → part (b) still open (needs a `fake-fapi` re-pin + rebuild; `demo-2`'s clone is
  pinned at `fast-build-m256-blocked-outcome`, older than the current tag).
- Everything else from iter-17's list stands unchanged.
**Lessons:**
1. **When a wait on a control times out, dump every candidate's label before concluding the path is closed.**
   Four passes and eleven minutes of waiting went into "Next is disabled, so this is not drivable." The button
   had been relabelled the whole time. Pass 3 even *saw* it — an empty match array where a button had been is
   the DOM telling you the label changed — and read it as "nothing matched."
2. **A working journey is not automatically a shippable Playthrough.** The right question is not "does it
   pass?" but **"when it fails, will the RED name the right culprit?"** A live scrape of a site that refuses
   automated clients answers no, and that is disqualifying however green it is today.
3. **A verdict is worth rewriting even when nothing lands.** Four blockers went from asserted to measured; two
   were false in their headline claim; one product defect surfaced with an endpoint attached. The next attempt
   at this cluster starts several hours ahead — which is the whole point of a written verdict.
4. **Back up before you mutate, and never with a clever one-liner.** A mis-escaped `perl -0pi -e` corrupted the
   page object mid-mutant; the `cp` taken thirty seconds earlier made it a two-minute detour. Every later
   mutant used an asserted single-occurrence replacement.
