# M256 — the coverage verdict register (exit-gate clause 3, "zero silent gaps")

Exit-gate clause 3 requires that **every remaining uncovered curated use case carries a written verdict**. This
file is that register. It is the authoritative answer to *"why is this curated use case still not covered?"* for
every one of the M201 curated corpus's 28 use cases that M256 does not land.

**A verdict is not a shrug.** Each entry below names the **specific missing piece**. That is what distinguishes
a verdict from a silent gap — and, in practice, it is also what proved that **none** of these is
`unimplementable-without-platform-edit`, which is why the milestone's re-scope trigger (*"> 3 un-homed curated
UCs prove unimplementable"*) never fired.

## The arithmetic

| | count | where |
|---|---:|---|
| M201 curated use cases | **28** | `knowledge/plan/releases/archive/02.00-opening-night/m201-manifest-corpus/manifest-draft.yaml` |
| Covered before M256 | 12 | — |
| Landed by M256 | **+4 org-admin** (iter-04 × 2 · iter-17 · iter-22 — the product is **4 of 4**) · **+4 onboarding** (`enterprise-workforce-ai-readiness.UC1` iter-26 · `enterprise-hiring.UC1` iter-27 · `enterprise-workforce-standard.UC2` iter-29 · **`individual.UC1` iter-32, the LAST un-homed use case in the corpus**) | `manifest/org-admin.yaml` · `manifest/onboarding.yaml` |
| Declared with verdicts by M256 | **+4 onboarding** (iter-08, re-priced iter-18/24/26) · **+7 here** | see below |
| ↳ *superseded* | the 2 org-admin `TODO`s (iter-04/05) and the 5th onboarding verdict were **discharged by landing** | — |
| Remaining with NO verdict | **0** | — |

M256 also added **2 net-new, non-curated** use cases (`skill-paths.save-for-later.UC1` iter-06,
`onboarding.completion.UC1` iter-08). They grow the **manifest** denominator, never the curated one — stated at
each site so the two denominators cannot be conflated.

---

## A. Un-homed by any milestone — now verdicted

### A1. `workforce-intelligence.organization-feedback.UC1` — **BUILD NEXT (cheapest remaining coverage)**

*A manager reads the org's collected post-simulation feedback — the sentiment recap + the per-member table.*

**Verdict: buildable now, and it is the cheapest un-homed UC left.** It is a **read/monitoring** flow of exactly
the shape the three existing Workforce-Intelligence Playthroughs already prove, on a real non-flag-gated route
(`/enterprise/organization-feedback`) with a **wired** backend (`resolver_organization_feedback.go`,
`jobsimfeedback.go`). Crucially the data axis **is already seeded**: `stackseed` reports `feedback rows=70` on
every `pt-world` reset, so the sentiment recap and per-member table have real aggregates to render.

**Missing piece:** a page object + spec. No seeder work, no platform work. Named handler:
`PT-M256-org-feedback`.

### A2. `profile-skills.import.UC1` — ~~blocked on the empty `fixtures/` dir~~ → **blocked on the CV route itself** (re-cut, iter-18)

*An already-onboarded user enriches their profile by importing a CV / LinkedIn profile.*

**Verdict: buildable, and its blocker MOVED — measured, not reasoned (v2.8 M256 iter-18).** M201's adversarial
verify resolved this **POSITIVE**: a real post-onboarding re-import surface exists and is wired
(`/reimport-profile`, the `OnboardingUser` component with its `reimport` prop). Two paths exist. iter-18 drove
both on `demo-2` — against the *onboarding* import step, which is the same component — and **both halves of
this entry's original diagnosis need correcting:**

- **~~The fixture is missing.~~ RESOLVED.** `playthroughs/fixtures/` is no longer empty: it holds
  `synthetic-cv-sre.pdf` + `.docx` (a wholly invented person, RFC-2606 `.example` domain, employer and school
  names that occur nowhere in the seed or the taxonomy — so an assertion naming them can only be satisfied by
  *that file having been imported*), plus a README stating the synthetic-only rule. The one-time cost this
  entry priced has been paid, and the file-chooser pattern with it (`resumeFileInput()` — set files on the
  input; the visible `Upload` button is only a trigger and vanishes once a file is attached).
- **~~The LinkedIn branch is not available on a demo.~~ REFUTED.** It works. Measured: type a public profile
  URL, the forward control relabels `Next` → `Import` and enables, the import runs (counter `5 → 8 → 50`) and
  fills the preview with real career facts in **~15 s**. The *other* half of the original objection —
  **not deterministic** — stands, and is now the whole reason it is refused: it scrapes a live third-party
  site that blocks automation, so a RED would read as an Anthropos regression. That is misattribution.

**The actual missing piece, and it is upstream of everything this entry used to name:** on the CV route the
file attaches and uploads **`200 POST /api/resources/resume`** — and then the parse counter stays at `0` and
the forward control stays `Next`/DISABLED for 100 s+, hidden DOM nodes included, on **both** formats. The
deterministic route is blocked by the product, not by the harness. **PRODUCT-DEFECT CANDIDATE** — see
`iter-18/decisions.md` D87. **Still deliberately shared** with
`onboarding.enterprise-workforce-standard.UC1`: one fix unblocks both, and the fixture they were both waiting
on now exists. Named handler: `PT-M256-resume-fixture-pair` (unchanged; its content is now "get the résumé
parse running on a demo", not "check in a file").

### A3. `talk-to-data.query.UC1` — **blocked on a live Bedrock key (an integration boundary)**

*A manager asks a natural-language question about the org's data and gets a streamed answer.*

**Verdict: buildable in shape, blocked on infrastructure the demo does not have.** The route
(`/enterprise/talk-to-data`) and the backend (`app/internal/askengine` — NL→SQL over the org's data behind a
SQL-validation sandbox, writing `ask_conversations` / `ask_messages`) are real and wired. Its assertion boundary
is honest and already defined by §5.8: *an answer streams back*, never the answer's content.

**Missing pieces, both real:** (1) the `ask_*` tables must be migrated on the demo stack, and (2) it needs live
**Bedrock** credentials. This is the same class as the studio LLM lane — a genuine third-party integration, so
even once wired it belongs in the **separately-budgeted lane**, not the timed median. Named handler:
`PT-M257-talk-to-data` (routed to a later release milestone deliberately: an unavailable credential is not
something an iter can fix).

---

## B. Reserved by M206 (vision) — re-confirmed, with dates

These three have been carried as M206 reservations for **five releases**. Clause 3 requires a verdict, not an
inherited reservation, so each is re-examined here.

### B1. `ai-simulations.code.UC1` — **reservation CONFIRMED (M206)**

*A user completes a code AI simulation: writes + runs code and gets it graded.*

**Verdict: stays reserved for M206.** It needs the **code-execution path** live end-to-end (Roadrunner → Judge0)
plus grading. `roadrunner` and `gotenberg` are in the demo compose set, but nothing in the suite has ever proven
a Judge0 round-trip, and the M231 content-stories map classified AI-labs-style graded results as having **no
seedable result surface**. The honest position: this is an integration-tier Playthrough whose mirror/infra tier
is M206's subject. **Re-confirmed 2026-07-28**, not silently inherited.

### B2. `ai-simulations.interview.UC1` — **reservation CONFIRMED (M206)**

*A user completes a non-voice (text) interview simulation and receives feedback.*

**Verdict: stays reserved for M206.** The interview modality reuses the chat container (M201's finding), so the
*launch* is drivable — and `pt-aisim-chat-launch` already drives exactly that on an interview-typed sim. What is
NOT drivable is the **completion**: iter-06 measured that reaching the launch boundary writes **no session row
at all** (0 `jobsimulation.sessions` rows created), so the session begins only past the welcome dialog, which is
on the far side of the §5.8 live-AI boundary. Driving a full interview turn-by-turn against a live LLM is
precisely what M206's mirror tier exists for. **Re-confirmed 2026-07-28, now with a measurement behind it.**

### B3. `profile-skills.self-evaluation.UC1` — **reservation RE-EXAMINED, and it is a strong M257 candidate**

*A user self-rates their skills and the self-assessment is saved + shown on the profile.*

**Verdict: reservation is WEAK — this one is a write with a real read-back and should be re-homed.** Its curated
final is *"the self-assessment persists and shows on the profile (`user_skill_evidences.user_level` set)"* — a
persist-then-observe shape, i.e. exactly the MUTATING shape clause 2 spent four iters hunting. It needs no LLM,
no external integration, and no fixture. It was one of iter-05 D20's three nominated candidates for clause 2's
fifth write and lost only because iter-06 found two cheaper ones.

**Missing piece:** nothing structural — just the page object + spec, and a decision to un-reserve it. Recorded
as a **recommendation to move it out of M206**, since M206 is the *mirror-engine* tier and this UC needs no
mirror. Named handler: `PT-M257-self-evaluation`. Not done here: re-homing a reservation is a roadmap decision,
not an iter's call.

---

## C. Reserved by M207 (vision) — re-confirmed

### C1. `skill-paths.academy.UC1` — **reservation CONFIRMED (M207), and the blocker is now sharper**

*A learner completes an academy skill path and is awarded a Certificate of Mastery.*

**Verdict: stays reserved for M207.** M201 recorded it NOT-RUNNABLE on four counts (anonymous academy launch
with no hero login → account-scoped progress/cert writes fail closed; no assessment-sim/AI-lab module; chapter
closure is module-count-only; unseeded catalog). **v2.5 M236 iter-08 sharpened the decisive one**: a demo academy
has **no `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`**, so it serves its **committed filesystem catalog** and the seeded
`academy_chapter_progress` rows have **no reader** — which is why `app/cmd/academy-seed` is *moot on a demo*. So
the blocker is not merely "unseeded": on a demo the academy is structurally not DB-backed, and no amount of
seeding changes that.

**Missing piece:** the demo academy must be pointed at the platform academy subgraph (a bring-up wiring change,
`corpus/services/ant-academy.md`), before any progress→certificate journey can be driven. That is
demo-infrastructure work, correctly M207's. **Re-confirmed 2026-07-28** with the M236 evidence attached.

---

## D. The two remaining onboarding `TODO`s — now MACHINE-CHECKED, not prose (iter-31)

The onboarding verdicts have always lived in `manifest/onboarding.yaml`, and until iter-31 they lived there as
**prose in a story `note`** while the use cases themselves carried a bare `playthrough: TODO`. That mattered
more than it sounds: `ptreport` rendered every uncovered use case with one generic sentence — *"declared use
case, no Playthrough yet (build-reference gap)"* — so **the four-state map, which is what tooling and reviewers
read, stated the wrong position** about the very use case this milestone's gate is closed around.

Both now carry a machine-checked `verdict` block (schema + the bidirectional fence:
`playthroughs/manifest/manifest.go` §`Verdict`, `validator.go` §`checkVerdicts`; rendered by
`report.go` §`unimplementedDetail`; documented in `corpus/ops/demo/playthroughs.md`):

### D1. `onboarding.enterprise-workforce-standard.UC1` — `disposition: will-not-build`

*A new org member builds a career profile by importing a CV or LinkedIn profile (self-import).*

**Measured, then deliberately refused — and the refusal has held for thirteen iters.** `measured_by:` iter-18's
six live probe passes + **`D104`**. The journey is fully drivable and the whole measured version is preserved at
`e2e/drafts/onboarding-self-import.spec.ts.draft`; what makes it un-shippable is that its only advancing path
**scrapes a live public LinkedIn profile** on a site that blocks automation — a real person's profile would
become a permanent test fixture, and **its RED would read as an Anthropos regression when nothing about the
product had changed.** A test whose failure lies is worse than no test. The deterministic CV route is blocked by
a **product defect we do not own** (upload POSTs `200` for a valid PDF *and* a docx while the forward control
never enables, 100 s+ — A2 above, `iter-18` D87), and that single fix is what would make this land
byte-deterministically. It correctly names **no handler**: a refusal with an assignee is a contradiction.

> **This is the disposition the whole mechanism exists for.** *"We measured this and it should not be built"* is
> a **result**, not a shortfall — and it is why clause 3's onboarding half is honestly **4 landed + 1 written
> verdict** rather than 5 landed (`D104`, awaiting the user's ratification at close).

### D2. ~~`onboarding.individual.UC1` — `disposition: not-yet-built`~~ → **LANDED at iter-32**

**Its verdict lasted one iter, and its removal was FORCED rather than remembered** — iter-31's fence is
bidirectional, so landing a use case makes a leftover verdict a validation failure (watched RED live as
mutant V7). That is the half of the fence that exists for drift, and this is the first time it fired on
real work.

The record below is kept as the verdict it was; the landing narrative is in the manifest's story note and
in [`iter-32/progress.md`](iter-32/progress.md).

#### (the verdict as written at iter-31)

*A solo user with no organization completes first-run setup and reaches a usable home.*

**LANDABLE, and priced from measurement** (`measured_by:` iter-30 **D116**). The blocker carried in prose since
iter-08 read *"needs a seeder change (a member-less user + a roster seat)"*, and its load-bearing half — *can the
app serve a user with no organization at all?* — had never been tested. Measured the cheap way (delete a seeded
hero's membership row on `demo-2`, restored by `--reset`, rather than write a seeder to find out): **an org-less
user reaches the app** — she logs in, `/onboarding` serves the flow, `/home` renders. So this is a build gap,
never an unimplementable one, and the re-scope trigger moves *further away*.

**And the cost was understated ~5×:** `memberships` carries **12 incoming FKs across 8 tables**; one hero's live
dependents are `membership_skills` 7 · `membership_languages` 2 · `membership_tags` 1 ·
`organization_target_roles` 1. So **≥ 5 seeders**, not one — *and the FK list is only the entry fee*: FK breakage
fails loudly and names its consumer, while the persona/profile/activity seeders writing `organization_id`-bearing
rows for a user with no membership do not break at all and are simply wrong invisibly (the iter-28 **D112**
silence class).

---

## Summary — the verdict distribution

| Verdict | count | UCs |
|---|---:|---|
| Buildable next, no blocker | **1** | A1 org-feedback |
| Buildable, blocked on a checked-in fixture | **1** | A2 profile import (pairs with onboarding UC1) |
| Buildable, blocked on a SEEDER CAPABILITY — **discharged by landing (iter-26)** | **1** | `onboarding.enterprise-workforce-ai-readiness.UC1`: a **stage-0 end-user seat could not be DECLARED** (`aiReadinessStageFor`: manager→0, struggling→1, else→3), so an appended hero arrived COMPLETED and a Playthrough on her could have **passed on the completed surface**. Fixed by `Persona.AIReadiness = "not_started"` + the `pt-ai-onboard` seat — **not** a YAML append, which is what its blocker was recorded as for two iters. |
| Buildable, blocked on a live third-party credential | **1** | A3 talk-to-data |
| Reservation confirmed (integration/mirror tier) | **3** | B1 code · B2 interview · C1 academy |
| Reservation re-examined → recommend re-homing | **1** | B3 self-evaluation → M257 |
| **`unimplementable-without-platform-edit`** | **0** | — |

**Zero of the 28 curated use cases is unimplementable.** Every uncovered one is blocked on harness work, a
fixture, a credential, or a deliberate tier reservation — never on the platform refusing to be driven. That is
the strongest available answer to clause 3, and it is the reason the re-scope trigger never fired.
