# iter-18 — decisions

## D86 — The import step's forward control RELABELS, and that is what four probe passes read as "closed"

`/onboarding`'s import step has **one** forward button. With no source supplied it reads **`Next` and is
DISABLED**; the moment a LinkedIn URL is typed it becomes **`Import` and ENABLED** (on `input`, no blur
needed). So a `/^Next$/` locator matches an element that is *permanently disabled* on that path.

Passes 1–4 each concluded "Next is disabled, so the import path is not drivable." Pass 2 waited **4.1
minutes** on it; pass 5 waited **6.9**. Pass 3 came within one line of the answer — its post-typing filter
returned an **empty array** where a disabled button had been, which is the DOM saying *the label changed* —
and read it as "nothing matched." Pass 4 dumped every button's label and found `"Import"|en`.

**Decision:** locate a forward control by **intent** (`forwardControl()` → `/^(Next|Import)$/`), keep the two
label-specific accessors for assertions that deliberately name one state, and **fence the relabel** in
`onboarding-locators.unit.spec.ts` by capturing the shipped matcher and executing it against both strings.
Generalised into `playthroughs.md` § locator discipline: **when a wait on a control times out, dump every
candidate's label before concluding the path is closed.**

## D87 — The résumé/CV import route is blocked upstream of the fixture (product-defect candidate)

Measured on `demo-2`, seats `pt-employee`/`pt-manager`, on a real 1-page PDF (`cupsfilter`, extractable text)
**and** a Word 2007 `.docx` (`textutil`) — identically:

| | |
|---|---|
| the file attaches | ✅ filename renders, the `Upload` button is replaced by CV-length hints |
| the upload succeeds | ✅ `200 POST http://localhost:28082/api/resources/resume` |
| the parse starts | ❌ counter stays **`0`** |
| the forward control enables | ❌ stays `Next`/DISABLED for **100 s+** — including every `display:none` node |

Compare the LinkedIn route, same step, same seats: counter `5 → 8 → 50`, preview fills, control **enables in
~15 s**.

**So it is not a fixture-format problem** — two formats behave identically and one of them is a real PDF whose
text `pdftotext` extracts. It is upstream of the fixture. Recorded as a **product-defect candidate**, not a
harness gap, and it is the single change that would make `onboarding.enterprise-workforce-standard.UC1` and
`profile-skills.import.UC1` (`coverage-verdicts.md` A2) both land **byte-deterministically**.

## D88 — The working self-import journey is REFUSED on P6, and the test is misattribution

The LinkedIn route works and would have arrived **already controlled**: a non-resolving profile URL advances
to the identical step (so the advance is scrape-independent — measured) and the forward control **never**
enables (watched 120 s). Presence and absence, one route.

**Refused anyway.** What makes it green is a scrape of a live third-party site that blocks automation, so the
day it says no this Playthrough goes RED reading like an Anthropos import regression — **misattribution**, the
defect class `seed-facts-fence.unit.spec.ts` exists to prevent, sourced from outside the building. Against a
gate that promises **0 flake across 3 consecutive runs**, that is shipping a known coin-flip into the suite
whose entire value is trustworthiness.

**The objection considered and rejected:** *the suite already runs a live LLM lane*
(`pt-studio-advanced-generate`, 300 s). It does not carry — that is a metered API called with our own
credentials, designed to be called. The discriminator is **whose refusal produces the RED**, not whether the
network is touched.

**Deliverable instead:** the whole measured journey preserved at
`e2e/drafts/onboarding-self-import.spec.ts.draft` (with its negative control and its refusal rationale in the
header) + the verdict rewritten from measurement. The next attempt starts from evidence.

## D89 — Every seat is day-0; the FIRST Playthrough to drive onboarding consumes it

**Measured, not assumed.** `/onboarding` serves the flow for `pt-employee` (Org A), `pt-recruiter` (Org D) and
`pt-ai-started` (Org C) — every seat probed. The reason is iter-07's finding: completion lives in
`public.user_params.onboarding`, which is **NULL for all seeded users**, so day-0 is the DEFAULT.

**And the constraint that actually matters was demonstrated by this iter's own probe**, which locked itself out
of `pt-employee` on the following pass after completing the flow: whichever Playthrough drives onboarding
**first consumes that seat for the whole run**, and completion cannot be undone through the UI. `pt-free` is
already spent by `pt-onboarding-complete`.

**So a second onboarding Playthrough needs a seat APPENDED to `pt-world.seed.yaml`** — and append is the
operative word: `personaUserIndexFor` (`stack-seeding/seeders/persona.go:510`) indexes heroes by **declaration
order** (`idx = i + 1`), so appending leaves every existing hero's index — and therefore her whole seeded
verified-skill chain — untouched. Inserting would renumber the world. Recorded so the next attempt does not
re-derive it, and so nobody prices onboarding's remaining UCs as "needs a pre-onboarding seeder" again.

## D90 — The fence's own fail-closed floor caught the fence

`onboarding-locators.unit.spec.ts` test 5 was first written to scan `tests/` + `drafts/` for fixture filenames
and assert each exists. Its fail-closed floor (*"the scan must find at least one reference"*) **fired
immediately**: no shipped Playthrough uploads a file — that is precisely this iter's finding — and the draft
names the pair in brace form. A check with nothing to check.

Re-aimed at the artifact that actually carries the promise: **`fixtures/README.md`'s file table**, reconciled
**both ways** (promised-but-absent → RED; on-disk-but-undocumented → RED, because every fixture becomes profile
content in a seeded world and an unexplained one must not accumulate). Both directions mutation-proven.

*The floor did its job on the very test it was written into.* Worth recording, because the reflex is to relax
the floor.

## Mutants — 5 of 5 RED

| mutant | expected | result |
|---|---|---|
| M1 — narrow `forwardControl()` to `/^Next$/` | RED (test 2) | **RED** — "must ALSO match 'Import'" |
| M2 — remove a README-promised fixture from disk | RED (test 5) | **RED** — "promises a fixture that is not checked in" |
| M3 — add an undocumented file to `fixtures/` | RED (test 5, other direction) | **RED** — "does not account for" |
| M4 — unanchor `nextButton()` to `/Next/` | RED (test 3) | **RED** — "must be anchored" |
| M5 — `importButton()` stops handing a role matcher | RED (test 1, fail-closed floor) | **RED** — "captured none … vacuously" |

`e2e/lib/onboarding-page.ts` restored **byte-identical** (`aa6abc52a580`), `fixtures/README.md` likewise
(`c20970bdf49b`).

> **A process note against the top-of-prompt ban.** M1 was first attempted with a `perl -0pi -e` substitution
> that mis-escaped the alternation and **corrupted line 102**. It was recovered from the `cp` backup taken
> before the mutation — verified byte-identical by sha — and every later mutant used `Edit`/`python3` with an
> asserted single-occurrence match instead. The backup-first rule is what made a botched one-liner a
> two-minute detour instead of a reconstruction.
