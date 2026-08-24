# MC03 — Spec notes

_Per-iter probe notes accumulate here during the build. **Nothing below is measured yet** — the headers
are the seven clauses' shape, not findings._

> **Recording rule.** Every entry states **the instrument, the artifact it read, the host + stack `N`, the
> refs (rext tag, platform shas), and the date.** A verdict without its environment is not a finding
> (`build-budget.md`: *state the environment with every number*). Quote **machine-readable artifacts**
> (`autoverify.json`, the batch verdict, the patch ledger), never a printed log line — the 2026-08-23
> failure was a printed green line over a refused patch.

## Pre-flight

_Not run._

- **rext tag on origin?** (`git ls-remote --tags origin` — the M236 rung zero; *tagging is not
  publishing*) — _unchecked._
- **Host + stack inventory** (`/stack-list`) — _unchecked._
- **Which host is "the official demo host"**, and does it serve the reviewer's `demo1.anthropos.work`
  hostname? — _unresolved (see `overview.md` § Open questions)._
- **MC01 / MC02 verdicts consumed** — their unmet clauses become Clause 7 carries — _not read._

## Clause 1 — cold demo cycle → READY, zero patches refused/skipped

### 1.1 Cold-start evidence (host · N · rext tag · platform shas · timestamps)
_Not run._

### 1.2 READY (exit 0)
_Not measured._

### 1.3 `autoverify.json` green + freshness
_Not measured._ (Read the JSON, not stdout. Age check: the pre-M236 guard parsed a UTC `ts` as LOCAL on
BSD and failed **OPEN** west of UTC.)

### 1.4 Demo-patch ledger — denominator pinned, `refused = 0`, `skipped = 0`
_Denominator not pinned._ Source of truth: `demopatch-spec.md` §5 + `TestPatchInventory`.

| manifest | expected | applied | G7 post-condition | verdict |
|---|---|---|---|---|
| _(populate from the pinned inventory at iter-01)_ | | | | |

### 1.5 Flag-gated surface rendered in a browser
_Not observed._ (A curl hits an access gate; a bundle-grep proves presence, not that it parses.)

### 1.6 Instrument-adequacy note
_Not written._ (`/api/health-check` is public by design and cannot witness a 500 on gated pages.)

## Clause 2 — `/dev-up` on a local stack → READY, same properties

### 2.1 Host + `dev-N`
_Not run._ (`D-v28-15`: dev/test is LOCAL; `billion` is DEMO-ONLY.)

### 2.2 The dev-path READY definition
_Unpinned._ `build-budget.md` defines READY for `up-injected.sh`; the dev equivalent is unstated.

### 2.3 Enumerated demo↔dev differences (expected, with citations)
_Not enumerated._ Known candidates: `dev-setdress.sh` hard-refuses `N=0`; per-stack Directus is opt-in on
dev (`--local-content`); the container-side production-bucket strip is DEMO-only.

### 2.4 Browser evidence on dev
_Not observed._

## Clause 3 — batch gate GREEN on both, no `skipped`

### 3.1 Denominator (from `playthroughs.md`, as amended by M269)
_Unpinned._ Pre-release baseline for reference only: **32 live Playthroughs + 1 verdicted TODO across 33
manifest use cases and 11 products**. M269 moves this; re-read rather than carry.

### 3.2 `skipped` occurrences
_Not counted._ (2026-08-23: `skipped` on a `--public-host` stack — `BIND_HOST` / `D-M255-7`.)

### 3.3 `D-v28-3` compliance (runs to completion · no retries · one consolidated red set)
_Not verified._

### 3.4 Known-false-green check — `FIX-M256-studio-false-green`
_Not verified._ (The studio Playthrough matched empty scaffolding at **+2.1 s** and reported PASS.)

### 3.5 Graded on the stack as configured
_N/A yet._

## Clause 4 — the eight annotation requests, one at a time

### 4.1 Denominator reconciliation — **precondition, not a footnote**
_Unresolved._ Three counts in play: release headline **eight**; milestone labels **ten** (A1–A5, B1–B3,
C1–C2); reviewer's own numbering **seven** (2 cockpit + 3 consumption + 2 seeding). Record the mapping
here before grading any request.

### 4.2 Per-request evidence table
_Empty._

| label | request (short) | owner | surface exercised | observed | how observed | verdict |
|---|---|---|---|---|---|---|
| A1 | half-row non-manager hero cards | M266 | | | | _ungraded_ |
| A2 | hiring vantage: candidate / performing / under-performing | M266 | | | | _ungraded_ |
| A3 | no box-within-card on content cards | M266 | | | | _ungraded_ |
| A4 | language flag before each language copy | M266 | | | | _ungraded_ |
| A5 | Academy is consumed content — no pass/fail flag | M266 | | | | _ungraded_ |
| B1 | remove the simulation-limit gate in demo | M267 | | | | _ungraded_ |
| B2-chat | hero can start a chat simulation | M269 | | | | _ungraded_ |
| B2-code | hero can start a code simulation | M269 | | | | _ungraded_ |
| B2-doc | hero can start a document simulation | M269 | | | | _ungraded_ |
| B2-voice | voice — **verdict, not a pipeline** | M271 | | | | _ungraded_ |
| B3 | skill-paths first paint: empty flash + slow load | M270 | | | | _ungraded_ |
| C1 | Programs: 3–5 per section at distinct stages | M268 | | | | _ungraded_ |
| C2 | score ↔ pass/fail verdict agree at the 60 threshold | M268 | | | | _ungraded_ |

_(The row set above is the **label** scheme, deliberately finer-grained than the eight-request headline;
4.1 owns the mapping back.)_

## Clause 5 — promised docs exist and match the running stack

### 5.1 Promise set (enumerated from every v2.10 `overview.md`, incl. MC01/MC02)
_Delivery half carried from `overview.md` § Clause 5.1; MC01/MC02 halves **not yet enumerated**._

### 5.2 Per-doc read-against-the-stack
_None read._

| doc | section read | stack observation | matches? |
|---|---|---|---|
| _(populate at iter-01)_ | | | |

### 5.3 Conditional promises and their resolved conditions
_Unresolved._ M270's manifest promise ← its D-1 vehicle decision. M271's `safety.md` amendment ← a GO
verdict.

### 5.4 Doc-agrees-with-diff-but-not-with-reality findings
_None yet._

### 5.5 Dangling routings
_Unchecked._ Known: `playthroughs/manifest/ai-simulations.yaml:4` still routes to the **dissolved M206**.

## Clause 6 — the voice verdict

### 6.1 `corpus/ops/demo/voice-feasibility.md` exists, verdict stated explicitly
_Not checked._

### 6.2 B1–B5 dispositions (resolved / declared unresolvable)
_Not checked._

### 6.3 The data-controller decision on B4 — who, when, on what basis
_Not recorded._

### 6.4 `safety.md` amendment (only if GO)
_Condition unresolved._

### 6.5 The named fallback (only if NO-GO) + no standing text implying voice works
_Condition unresolved._

## Clause 7 — reproduce-or-name

### 7.1 Per-milestone claim ledger
_Empty._

| milestone | status at close | exit-gate claim | reproduced? | evidence / reason |
|---|---|---|---|---|
| M266 | | | | |
| M267 | | | | |
| M268 | | | | |
| M269 | | | | |
| M270 | | | | |
| M271 | | | | |
| MC01 | | | | |
| MC02 | | | | |

### 7.2 Routing of every NOT-REPRODUCED claim (re-open / carry with a NAMED destination / drop)
_None routed._ A carry to "the next release" or "a future pass" is **not a destination**.

### 7.3 `closed-incomplete` milestones and their unmeasured clauses
_Not compiled._ Live precedent from v2.8: **M257x and M258 closed `closed-incomplete` by user ruling**,
and **M258's clause 3 was never measured clean**.

### 7.4 Routing written at the destination
_Not written._ (The M258 lesson: a routing in a closing milestone's decisions is not a routing until the
target's own doc says so.)

## Release-level verification record

_Not compiled._ Assembles from the clause sections above once each is graded.
