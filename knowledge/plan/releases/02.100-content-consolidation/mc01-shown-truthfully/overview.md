---
milestone: MC01
title: "Checkpoint — shown truthfully"
milestone_shape: iterative
status: planned
release: "02.100-content-consolidation"
exit_gate: "all six clauses below hold on ONE running demo stack: the cockpit renders as designed (1-3), every seeded session row satisfies (score >= 60) == (completion == passed) over the WHOLE table (4), Programs shows 3-5 plans across >= 3 stages (5), and the four governing docs describe what actually shipped, read AGAINST THE RUNNING STACK (6). A clause that fails is a gate failure — including a doc clause."
iteration_protocol_ref: "corpus/ops/demo/coverage-protocol.md"
re_scope_trigger: "If a clause fails for a reason that belongs to the milestone that shipped it rather than to integration, route it BACK to that milestone rather than fixing it here — a checkpoint that absorbs the work it was built to detect stops detecting anything."
depends_on: "M266, M268"
parallel_with: "none"
complexity: medium
last_updated: "2026-08-24"
---

# MC01: Checkpoint — shown truthfully

**CHECKPOINT milestone.** It ships no feature. It **grades a cluster** — **M266 «Cockpit legibility»** +
**M268 «Seeded truth»**, everything about what the demo *SHOWS* — against a **running stack** and against
**its own docs**. It runs only when **both** have closed.

> ⚠️ **A checkpoint may FAIL, and failing is its PURPOSE, not a defect.** A red clause here is the
> milestone working. The output of a failed clause is **work routed back to the milestone that owns it**,
> not a fix applied here and not a follow-up ticket. See § *The re-scope trigger* — a checkpoint that
> absorbs the work it was built to detect stops detecting anything.

**Goal:** on a real demo stack, **every surface those two milestones touched renders as designed**, **every
seeded number agrees with its own verdict**, and **the four governing docs describe what actually shipped**.

## Why this milestone exists

**Every one of the following failure modes was MEASURED on 2026-08-23, on the live `demo-1` stack, in one
day.** None of them is hypothetical, and none was caught by a unit test, a diff, or a probe:

- **autoverify printed *"verify live: all liveness + readiness probes passed"*** while a demo-patch had been
  **REFUSED** and the hiring bundle shipped with **every PostHog flag OFF**.
- **`/api/health-check` answered 200** on a `studio-desk` container whose **every gated page returned HTTP
  500** — the route is public **BY DESIGN** and **structurally cannot witness that failure**.
- **The Playthrough batch gate recorded `skipped`, never green**, because the stack was `--public-host`
  (`BIND_HOST` / `D-M255-7`) — so the **release-level gate reported nothing at all** and looked fine doing
  it.
- **The studio Playthrough reported PASS** and was cited as evidence the migrated studio works; it matches
  **EMPTY SCAFFOLDING at +2.1 s** (`FIX-M256-studio-false-green`).
- **The manager menu was missing "Build a Course"** and pointed **Assign at a legacy surface**, on a stack
  **every probe called healthy**.

> **THE COMMON SHAPE: a milestone closes, its tests pass, its probes are green — and the thing it
> delivered does not work on a real stack, or works but is documented as something else.** A checkpoint
> milestone is the layer that grades the **CLUSTER** against a **running stack** and against **its own
> docs**, rather than against its diff.
>
> **SO THE GATE IS ALWAYS TWO-SIDED: it works on a real stack AND the corpus describes what actually
> shipped.** Read the doc against the **RUNNING STACK**, never against the diff that changed it — **a doc
> can agree with a commit and disagree with reality.**

## Exit gate

**All six clauses. Each is independently checkable by someone who was not in the room.** A clause is
green only when its stated evidence exists and is recorded in [`progress.md`](progress.md) — a viewport
width, a row count, a screenshot, a query output, a named doc section. **"It looks right" is not
evidence.**

All six are graded on **ONE** stack, in **ONE** pass, in the state that stack is actually in. Clauses
graded across different stacks or different reseeds do not compose — see § *Open questions*.

### Clause 1 — the hero-card grid (M266 / A1)

On a **cold** demo stack, the cockpit **org-stories** tab renders **NON-MANAGER hero cards two-per-row**
and the **manager card full-row**, **verified in a browser at a stated viewport width**.

*Checkable by:* open the cockpit (`:7700`+offset) in a browser at a **recorded** viewport width, count
cards per row in the org-stories tab, confirm the manager card spans the full row. Record the width, the
story observed, and the observed cards-per-row. **A width must be stated** — a grid claim without one is
not a claim.

### Clause 2 — the hiring vantage speaks candidate (M266 / A2)

In the **"Candidate Hiring & Comparison"** story, **non-manager heroes read `CANDIDATE`** and their state
pills read **`PERFORMING` / `UNDER-PERFORMING`** — **not `EMPLOYEE` / `THRIVING` / `STRUGGLING`**.

*Checkable by:* read the badge text on every non-manager hero card in that story, in the browser, on the
same stack as clause 1. **Zero occurrences** of `EMPLOYEE`, `THRIVING` or `STRUGGLING` in that story's
non-manager cards. The manager seat in the hiring org must **still read as a manager** (the
`_avatar_class()` precedent — the manager branch wins first).

### Clause 3 — the content-story card shape (M266 / A3 + A4 + A5)

Content-story cards render with **NO nested `.ctcol` box**; **language pills carry an inline-SVG flag**;
and the **Academy product shows NO pass/fail chip** while **products that HAVE a verdict still do**.

*Checkable by:* three independent observations on the rendered cockpit page — (a) `.ctcol` does not occur
in the served HTML; (b) the language pill contains an `<svg>` element (not an emoji, not a FontAwesome
glyph); (c) the Academy product's card carries no pass/fail chip **and** at least one verdict-carrying
product's card still does. **The verdict-membership oracle is M266's own recorded decision** on
`has_verdict` — this clause grades the stack against that decision, it does not re-take it.

### Clause 4 — the score/verdict invariant, over the WHOLE table (M268 / C2)

Over the **WHOLE of `public.job_simulation_sessions`** on that stack — **not a sample** — **every row**
satisfies **`(score >= 60) == (completion == passed)`**. **Asserted by a query or fence test, with the
row count stated.**

*Checkable by:* one SQL statement (or the M268 fence test) run against that stack's Postgres, returning
**0 violating rows**, alongside the **total row count** the assertion covered. **Both numbers are
required** — a "0 violations" with no denominator does not distinguish a clean table from an empty one,
which is exactly the fail-open shape this release exists to kill.

### Clause 5 — Programs is populated (M268 / C1)

**`/enterprise/assignments-list`** shows **3–5 programs** spanning **at least three distinct stages**,
**verified in a browser**.

*Checkable by:* load the surface in a browser as a seated manager hero on the same stack, count the
program rows (**>= 3 and <= 5**), and record the **distinct stage values as rendered** (**>= 3**). The
reading rule for "stage" is what the **surface displays**, not what a column stores — see § *Open
questions*.

### Clause 6 — the four governing docs describe what shipped

**`corpus/ops/demo/cockpit-spec.md`**, **`corpus/ops/demo/content-stories-spec.md`**,
**`corpus/ops/seeding-spec.md`** and **`corpus/ops/demo/stories-spec.md`** each describe the **shipped
behaviour**, **verified by reading each doc AGAINST THE RUNNING STACK**. **Any drift is a gate failure,
not a follow-up.**

*Checkable by:* a per-doc pass in which each behavioural assertion the doc makes about the surfaces in
clauses 1–5 is checked against **what the running stack does**, with the checked sections named. ⚠️ **Do
NOT grade a doc against the commit that changed it** — the failure mode this clause exists for is a doc
that **agrees with its own diff and disagrees with reality**. A doc that is *silent* where the stack now
behaves differently is **also drift**: the four docs are the **contract** for these surfaces, and an
undocumented shipped behaviour means the contract no longer describes the product.

## Iteration protocol

[`coverage-protocol.md`](../../../../../corpus/ops/demo/coverage-protocol.md) — the **measure → triage →
fix → re-measure** loop. **This milestone uses its semantic-believability posture, not its crawl scope:**
the question is *"does this surface show the real thing, and does the doc say so"*, not *"did every page
in a crawl return 200"*. Its **fix-surface routing table** is the model for this milestone's triage step —
except that here the routing destination is usually **a delivery milestone**, not a file.

## Why iterative (not section)

**You cannot write the fix list before you have looked.** A section checklist authored at scaffold time
would be a list of *expected* findings, and a checkpoint whose findings were known in advance is not
measuring anything. Each iteration measures the cluster on a running stack, triages what came back, routes
it, and re-measures.

**Iteration count is not predictable and must not be padded to look busy** — if the first measurement pass
returns six green clauses with recorded evidence, the milestone closes on one iteration and that is a
correct outcome, not a shallow one.

## The re-scope trigger

> **If a clause fails for a reason that belongs to the milestone that shipped it rather than to
> integration, route it BACK to that milestone rather than fixing it here — a checkpoint that absorbs the
> work it was built to detect stops detecting anything.**

Practically:

- **A defect in M266's or M268's own delivery** (wrong label set, an unclamped score, a `.ctcol` still
  rendering) → **route to that milestone**, which re-opens. MC01 records the finding and the routing.
- **A defect in how the two compose** (M266 renders a `has_verdict` M268's seed never populates; a card
  reads a manifest field the seeder stopped writing) → **that is integration, and it is MC01's**.
- **A defect that is neither** — a pre-existing platform defect, or a surface neither milestone touched →
  **`knowledge/plan/platform-defect-register.md`**, and the clause is graded on whether the *cluster's*
  behaviour is correct, not on whether the platform is perfect. **Say which of the three it is, in
  writing, every time.** The classification is the deliverable of a red clause.

## Depends on

**M266** and **M268** — **both closed**. MC01 does not run against a half-landed cluster: a clause graded
against work still in flight measures the flight, not the cluster.

## Parallel with

**none.** Checkpoints never parallelise — the whole job is to observe a cluster that has **finished**.

## Out

- **Anything M267 / M269 / M270 / M271 own.** That is **MC02** (M267 + M269 + M270, on demo **and** dev)
  and **MC03** (the whole release + the docs). MC01's subject is *what the demo SHOWS*; entitlement,
  modality Playthroughs, skill-paths first paint and voice are not it.
- **Fixing what belongs to M266 or M268.** See the re-scope trigger. MC01 measures, classifies and routes.
- **New behaviour of any kind.** A clause that would only pass if something new were built is a **failed
  clause with a routing**, never a licence to build it here.
- **Any platform-repo edit.** Standing release constraint (below); a need that can only be met by one
  **escalates**.
- **Re-taking M266's or M268's recorded decisions.** Their decision records are this milestone's oracle.
  Disagreeing with a decision is a routing back to its owner, not a re-litigation here.

## Open questions

Genuinely open at scaffold time. **Resolve each in [`spec-notes.md`](spec-notes.md) before grading the
clause it affects** — resolving one by assumption mid-measurement is how a checkpoint produces a false
green.

- **What exactly is "a cold demo stack" here, and is it `--public-host`?** Clause 1 says *cold*. The demo
  path is **remote-reach default-on** (v2.3 `D-DESIGN-3`), and a `--public-host` stack makes the
  **Playthrough batch gate record `skipped`** (`BIND_HOST` / `D-M255-7` — one of the five measured
  failures above). MC01's clauses are browser-and-SQL observations and do **not** depend on that gate, but
  **whether MC01 brings its stack up with or without `--public-host`** decides whether the batch gate is
  even readable alongside them, and must be **stated, not defaulted into**.
- **Which viewport width does clause 1 state, and why that one?** A grid claim is width-dependent. The
  width must be **chosen deliberately and recorded** (a presenter's real screen is the honest referent),
  and whether a **second** width is graded — to catch a grid that only works at one size — is undecided.
- **What is clause 1's expected render for the odd-count and no-manager cases?** M266 lists both as open
  at its own scaffold time. MC01 grades against **whatever M266 decided**; if M266 closed without
  deciding, that is itself a finding routed back, not a judgement call for MC01.
- **Does clause 4's "every row" include the CLONED content-story sessions?** `content_stories_write.go`
  takes score **and** `passed` from a **real production session** (M268 § Open questions): *"a real
  session that disagrees is truth, not a bug"*. If M268 closed with those rows **outside** the fence, then
  clause 4 as written ("over the WHOLE table") and M268's fence **disagree about the denominator**.
  Resolving that is MC01's first-iteration job, and the resolution must be **stated with the row count** —
  an excluded class must be **named and counted**, never silently dropped from the assertion.
- **Was the stack under test seeded by the tooling that actually contains M268's fix?** A stack seeded
  from an older pinned `rosetta-extensions` tag can pass clause 4 by accident or fail it by staleness.
  **Verify the consumed tag — and that it is on origin** (*tagging is not publishing*, the M236 pre-flight
  rung zero) before trusting any clause-4 number.
- **What counts as a "distinct stage" for clause 5?** M268 records that whether a program's stage is a
  stored `status` value or **computed by the surface** from enrollment + assignment completion is **not
  measured**. Clause 5 grades **what the browser renders**, so the reading rule must be written down
  before counting — otherwise two graders count differently and the clause is not checkable by a third
  party, which is the whole requirement.
- **Which seat does clause 5 use, and does the answer change the count?** `/enterprise/assignments-list`
  is a manager surface; M268 left *"which orgs get 3–5 plans"* open (all four, or workforce only). If the
  hiring org is excluded by design, grading clause 5 from the hiring manager's seat produces a false red.
- **How is clause 6 evidenced?** Reading a doc against a stack produces a judgement; the gate needs an
  artifact. Whether that is a per-doc section-by-section note in `spec-notes.md`, a table of
  (doc § → observed behaviour → verdict), or something else is undecided — but **an unevidenced clause-6
  pass is exactly the "green probe, broken product" shape this milestone exists to catch**, so it cannot
  be a bare assertion.
- **Do any of the four docs assert things about surfaces OUTSIDE this cluster?** `seeding-spec.md` and
  `stories-spec.md` in particular are broad. Clause 6 is scoped to *what actually shipped in M266 + M268*;
  drift in an unrelated section of the same file is **MC03's**, and confusing the two either inflates
  MC01 or lets real drift through. The scope line must be drawn explicitly, per doc.

## KB dependencies

- [`coverage-protocol.md`](../../../../../corpus/ops/demo/coverage-protocol.md) — the
  `iteration_protocol_ref`: the measure → triage → fix → re-measure loop and its **fix-surface routing
  table**; the **semantic-believability** posture (real seeded content, per-section cardinality, persona
  self-consistency), used here **without** its crawl scope.
- [`cockpit-spec.md`](../../../../../corpus/ops/demo/cockpit-spec.md) — the **M43/M242 card contract**;
  the governing doc for clauses 1 and 2, and one of the four clause-6 docs.
- [`content-stories-spec.md`](../../../../../corpus/ops/demo/content-stories-spec.md) — the content-story
  card contract (**§7.2** tuple row, **§7.6** language toggle, **§4** honesty gate); the governing doc for
  clause 3, and a clause-6 doc.
- [`seeding-spec.md`](../../../../../corpus/ops/seeding-spec.md) — the **data-DNA** and the
  production-isolation boundary; where M268's score-verdict invariant lands as a documented gene. Governing
  doc for clause 4, and a clause-6 doc.
- [`stories-spec.md`](../../../../../corpus/ops/demo/stories-spec.md) — the **7-table verified-skill
  fan-out** and the **G14** session-seeder rules; where Programs joins the seeded story surface. Governing
  doc for clause 5, and a clause-6 doc.
- [`verification.md`](../../../../../corpus/ops/verification.md) — **the two tail gates and how they fail
  in opposite directions** (non-fatal auto-verify vs the loud Playthrough batch gate), plus **pre-flight
  rung zero** (*tagging is not publishing*). This is the doc that explains **why a green probe is not a
  green product**, which is this milestone's entire premise.

**Delivers → no net-new doc.** MC01's output is a **verdict per clause with recorded evidence**, plus
**routings** for anything red. Corpus edits it makes are **corrections to the four governing docs where
clause 6 finds drift** — and those corrections are **in scope by construction**, because clause 6 grades
them.

## Constraints (release-wide, non-negotiable)

- **Zero platform-repo edits.** Any platform source change goes through the sha-pinned **demopatch**
  mechanism ([`demopatch-spec.md`](../../../../../corpus/ops/demo/demopatch-spec.md)). A need that cannot
  be met that way **escalates**; it does not edit. A checkpoint in particular **observes** — it does not
  reach into a platform repo to make a clause pass.
- All stack tooling lives in `rosetta-extensions`, built + tested in the authoring copy, tagged, **pushed
  to origin** (the M236 pre-flight rung zero — *tagging is not publishing*), then consumed per-stack at a
  pinned tag.
- Secrets handled **values-blind** — no verb reads, echoes, logs or commits a value.
- **Customer media never enters an agent's context** (`media-substrate-spec.md` PII discipline). Content
  stories are in clause 3's blast radius; you grade the **card**, you do not open the media.
