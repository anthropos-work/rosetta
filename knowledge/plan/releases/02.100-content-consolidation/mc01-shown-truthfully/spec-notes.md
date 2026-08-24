# MC01 — spec notes

_Per-iter probe notes accumulate here during the milestone. Section headers are derived from the six exit-gate
clauses and the open questions; **nothing below is measured yet** — the headers are the grading surface, not
findings._

> **A checkpoint's notes are EVIDENCE, not narrative.** Every clause verdict recorded in
> [`progress.md`](progress.md) must be traceable to something written here: a viewport width, a row count, a
> query and its output, a doc section and the behaviour it was read against. **A clause with no note is not a
> green clause.**

## The stack under test

_Not brought up. Record here, before any clause is graded: stack id (`demo-N`), bring-up command and flags,
whether `--public-host` was used, the consumed `rosetta-extensions` **tag** (and that it is on origin), the
platform refs, whether the stack was cold, and the timestamp._

### Pre-flight — is this stack gradeable at all?

_Not run. The five 2026-08-23 failures are the checklist: does `autoverify` actually witness what it claims;
did any demopatch get REFUSED; did the batch gate record `skipped`; is any probe structurally blind to the
surface it fronts._

## Clause 1 — the hero-card grid (M266 / A1)

### Viewport width chosen, and why

_Not chosen._

### Observed cards-per-row, org-stories tab

_Not observed._

### The manager card's full-row span

_Not observed._

### Odd-count / no-manager cases — what M266 decided

_Not read._

## Clause 2 — the hiring vantage speaks candidate (M266 / A2)

### Badge text on non-manager heroes in "Candidate Hiring & Comparison"

_Not read._

### Zero occurrences of EMPLOYEE / THRIVING / STRUGGLING in that story

_Not checked._

### The hiring MANAGER seat still reads as a manager

_Not checked._

## Clause 3 — the content-story card shape (M266 / A3 + A4 + A5)

### `.ctcol` absent from the served HTML

_Not checked._

### The language pill carries an inline `<svg>` flag (not emoji, not FontAwesome)

_Not checked._

### Academy shows no pass/fail chip

_Not checked._

### A verdict-carrying product still shows one — against M266's recorded `has_verdict` decision

_M266's decision not yet read. **This clause grades the stack against that decision; it does not re-take it.**_

## Clause 4 — the score/verdict invariant over the WHOLE table (M268 / C2)

### The assertion used (query or fence test), verbatim

_Not written._

### Violating rows

_Not measured._

### Total rows the assertion covered

_Not measured. **Both numbers are required** — "0 violations" without a denominator does not distinguish a
clean table from an empty one._

### Denominator question — are the CLONED content-story sessions inside the fence?

_Unresolved. `content_stories_write.go` copies score and `passed` from a real production session; M268 left the
treatment open. If a class is excluded it must be **named and counted**, never silently dropped._

### Tag provenance — was this stack seeded by tooling that contains M268's fix?

_Not verified._

## Clause 5 — Programs is populated (M268 / C1)

### Seat used, and org

_Not chosen._

### Program rows rendered at `/enterprise/assignments-list`

_Not counted._

### The reading rule for "stage" — what the SURFACE displays

_Not written. M268 left stored-`status`-vs-computed unmeasured; the rule must exist before counting or the
clause is not checkable by a third party._

### Distinct stages observed

_Not observed._

## Clause 6 — the four governing docs vs the RUNNING STACK

> ⚠️ **Read each doc against the stack, never against the diff that changed it.** A doc can agree with its own
> commit and disagree with reality. **Silence where the stack now behaves differently is also drift.**

### Evidence shape — how a doc read is recorded

_Undecided (see `overview.md` § Open questions). An unevidenced clause-6 pass is the exact "green probe, broken
product" shape this milestone exists to catch._

### Scope line — which sections of each doc are MC01's, and which are MC03's

_Not drawn._

### `cockpit-spec.md`

_Not read._

### `content-stories-spec.md`

_Not read._

### `seeding-spec.md`

_Not read._

### `stories-spec.md`

_Not read._

## Triage ledger — findings and their routing

_Empty. Every red clause is classified as one of three, **in writing**:_

| # | finding | clause | classification (M266 / M268 / integration / platform-defect) | routed to |
|---|---|---|---|---|
| _(none yet)_ | | | | |

_**M266 / M268** → route back, the milestone re-opens. **integration** → MC01 owns it. **platform-defect** →
`knowledge/plan/platform-defect-register.md`, and the clause is graded on the cluster's behaviour._
