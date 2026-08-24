---
milestone: MC03
title: "Final checkpoint — the release is real"
milestone_shape: iterative
status: planned
release: "02.100-content-consolidation"
exit_gate: "all seven clauses hold: (1) a cold /demo-down --purge + /demo-up on the official demo host reaches READY with autoverify GREEN and ZERO demo-patches refused or skipped; (2) a /dev-up on a local stack reaches READY with the same properties; (3) the Playthrough batch gate is GREEN on both and its verdict contains no 'skipped'; (4) each of the eight annotation requests is demonstrated one at a time against the running stack with per-request evidence recorded; (5) every 'Delivers ->' doc promised by any v2.10 milestone exists and describes shipped behaviour, including the NET-NEW corpus/ops/demo/voice-feasibility.md; (6) the M271 voice GO/NO-GO verdict is recorded in the corpus with its reasoning, and corpus/ops/safety.md is amended if the verdict touches the shared-AWS exposure; (7) no milestone in this release closed on a claim this checkpoint cannot reproduce — any that did is named and either re-proven or recorded as an honest carry."
iteration_protocol_ref: "corpus/ops/verification.md"
re_scope_trigger: "If a clause cannot be met and the fix belongs to a milestone already closed, that is a RELEASE-SCOPE finding: surface it to the user with the honest options — re-open the owning milestone, carry it with a NAMED destination, or drop it with the reason recorded — rather than closing the release on a clause that did not hold."
depends_on: "MC01, MC02, M271"
parallel_with: "none"
complexity: large
last_updated: "2026-08-24"
---

# MC03: Final checkpoint — the release is real

**THE FINAL CHECKPOINT.** Runs **last** — after M271 and after both earlier checkpoints (MC01, MC02).
This is the milestone that decides whether **v2.10 can close**.

**Goal:** prove the **whole release works on a cold stack — demo AND dev**, that every one of the **eight
annotation requests** is demonstrably satisfied, and that **the corpus tells the truth about all of it**.

> ⚠️ **A checkpoint may FAIL and send work back. That is its purpose, not a defect.** A RED clause here
> is a successful checkpoint — it found the thing the delivery milestone's own green gate could not see.
> The failure mode to fear is the opposite one: **a checkpoint that absorbs the work it was built to
> detect stops detecting anything.** Fixes route to the milestone that owns them (see § Routing, not
> absorbing).

## Why these checkpoints exist

Every one of the following failure modes was **MEASURED on 2026-08-23, on the live `demo1` stack, in one
day**. None is hypothetical, and none was caught by the delivering milestone's own tests:

> - **autoverify** printed *"verify live: all liveness + readiness probes passed"* while a **demo-patch had
>   been REFUSED** and the hiring bundle shipped with **every PostHog flag OFF**.
> - **`/api/health-check` answered 200** on a `studio-desk` container whose **every gated page returned
>   HTTP 500** — the route is **public BY DESIGN** and *structurally cannot* witness that failure.
> - **The Playthrough batch gate recorded `skipped`, never green**, because the stack was `--public-host`
>   (`BIND_HOST` / `D-M255-7`) — so the **release-level gate reported nothing at all**, and looked fine
>   doing it.
> - **The studio Playthrough reported PASS** and was cited as evidence the migrated studio works; it
>   matches **EMPTY SCAFFOLDING at +2.1 s** (`FIX-M256-studio-false-green`).
> - **The manager menu was missing "Build a Course"** and pointed **Assign at a legacy surface** — on a
>   stack **every probe called healthy**.

**THE COMMON SHAPE:** a milestone closes, its tests pass, its probes are green — **and the thing it
delivered does not work on a real stack, or works but is documented as something else.** A checkpoint
milestone is the layer that grades **the CLUSTER against a running stack and against its own docs**,
rather than against its diff.

**SO THE GATE IS ALWAYS TWO-SIDED: it works on a real stack AND the corpus describes what actually
shipped.** Read the doc **against the RUNNING STACK, never against the diff that changed it** — a doc can
agree with a commit and disagree with reality.

## Exit gate — seven clauses, each independently checkable

Every clause below is written so that **someone who was not in the room** can grade it from the recorded
evidence. A clause is either **GREEN with evidence cited**, or **RED with the finding named**. There is no
third state; "looked fine" is not a grade, and a clause graded from a printed log line rather than from a
machine-readable artifact is **not graded**.

### Clause 1 — a cold demo cycle reaches READY with nothing silently unapplied

A cold **`/demo-down --purge` + `/demo-up`** cycle on the **official demo host** reaches **READY**, with
**autoverify GREEN** and **ZERO demo-patches refused or skipped**.

> *"The stack is UP" is NOT the bar. The bar is UP with nothing silently unapplied* — on 2026-08-23 a
> green-looking bring-up shipped a **refused patch**.

- **1.1** The cycle starts **cold**: `/demo-down --purge` completes first, and no container, volume or
  network for that `N` survives it. Record: host, `N`, `rosetta-extensions` **tag** (and that the tag is
  **on origin** — `git ls-remote --tags origin`, the M236 pre-flight rung zero), each platform repo's sha,
  and the start/end timestamps.
- **1.2** **READY** is the `build-budget.md` definition and nothing looser: `up-injected.sh` **exits 0**
  **AND** `autoverify.json` is green. Record the exit code.
- **1.3** The green verdict is read **from `autoverify.json`**, not from stdout, and it was **produced by
  this cycle** — the freshness/age check passes (and is read with the M236 UTC-vs-local fix in mind: west
  of UTC the pre-fix guard aged a **stale** verdict as fresh, i.e. it failed **OPEN**).
- **1.4** **Demo-patch ledger, fail-CLOSED.** Every manifest in the canonical demopatch inventory is
  accounted for as **applied with its G7 apply post-condition satisfied**, or **explicitly not-applicable
  to this stack with a reason recorded**. `refused = 0` and `skipped = 0`. **The denominator is pinned**
  against `demopatch-spec.md` §5 / `TestPatchInventory` **before** counting — *zero refusals over an empty
  or truncated set is not a pass* (the M236 `EXPECTED_PAIRS` lesson).
- **1.5** At least one **flag-gated surface** is rendered **in a browser** against this stack and shows the
  gated behaviour ON. *A curl hits an access gate and a bundle-grep proves presence, not that it parses* —
  and the 2026-08-23 failure was precisely a refused patch yielding an all-flags-OFF bundle behind a green
  probe.
- **1.6** No probe used to grade this clause is one that **structurally cannot witness** the failure it is
  cited against — the `/api/health-check` lesson. If a probe is public-by-design, say so and cite a second
  instrument.

### Clause 2 — a local dev stack reaches READY with the same properties

A **`/dev-up`** on a **local** stack reaches READY with the same properties as Clause 1.

- **2.1** The dev stack is named explicitly (`dev-N`, with `N` recorded) and the host is recorded.
  **D-v28-15 applies: dev/test is LOCAL** to the current development Mac; `billion` is DEMO-ONLY.
- **2.2** READY, exit code, and `autoverify.json` green — same instruments, same freshness rule as 1.2/1.3.
  If the dev path's READY definition differs from `up-injected.sh`'s, the **difference is written down**
  and the substitute artifact named (see § Open questions).
- **2.3** Demo-patch/`--local-content`/set-dress differences between the two paths are **enumerated rather
  than assumed equal**: the main dev stack (`N = 0`) is deliberately never set-dressed
  (`dev-setdress.sh` hard-refuses `N=0`), per-stack Directus is **opt-in on dev** (`--local-content`), and
  the container-side production-bucket strip is **DEMO-only**. Each difference is recorded as *expected*
  with its citation, or graded RED.
- **2.4** The same browser-level evidence as 1.5, taken on the dev stack.

### Clause 3 — the Playthrough batch gate is GREEN on both, with no "skipped"

The Playthrough batch gate is **GREEN on the demo stack AND on the dev stack**, and **its verdict contains
no `skipped`**.

- **3.1** The batch verdict artifact is quoted, per-Playthrough, with a **pinned denominator** taken from
  `playthroughs.md` (authoritative for the count) **as amended by M269**. A shrunken denominator that
  happens to be all-green is RED.
- **3.2** The string `skipped` appears **zero** times in the verdict. This is the exact 2026-08-23 failure:
  a `--public-host` stack made the gate skip, so a release-level gate **reported nothing at all**
  (`BIND_HOST` / `D-M255-7`, landed by M269).
- **3.3** The run obeys `D-v28-3`: **runs to completion**, never halts at first red, **never retries**, and
  produces **ONE consolidated red set** at batch end. A green obtained by retry is RED here.
- **3.4** No Playthrough in the set is a **known false green**. `FIX-M256-studio-false-green` is named
  explicitly: the studio Playthrough matched **empty scaffolding at +2.1 s** and still reported PASS. Its
  post-M269 assertion is cited, not assumed.
- **3.5** If the demo stack under test is `--public-host` (the default for `/demo-up` since D-DESIGN-3),
  clause 3 is graded **on that stack as configured** — not on a second, non-public stack stood up to make
  the gate run. Making the gate green by removing the condition it fails under is RED.

### Clause 4 — each of the EIGHT annotation requests, demonstrated one at a time

**EACH of the eight annotation requests is demonstrated individually against the running stack, with the
evidence recorded per request.** **A cluster-level pass does not discharge an individual request**, and
neither does a milestone's own closure.

- **4.1** **The denominator is pinned first.** The release headline says **eight field requests**
  (`roadmap.md` § *In Development — v2.10*); the delivery milestones use a **ten-label** scheme
  (**A1–A5**, **B1–B3**, **C1–C2**); the reviewer's own file numbers **seven** items (2 cockpit + 3
  consumption + 2 seeding). **These three counts must be reconciled and the reconciliation recorded in
  `spec-notes.md` before any request is graded** — a per-request gate cannot be graded against an
  unpinned denominator. Grade against the **label set**, which is the finest-grained of the three; report
  the mapping back to the reviewer's numbering.
- **4.2** Per label, the record contains: **the request text**, **the owning milestone**, **the exact
  surface exercised** (URL or command, on the running stack), **the observed result**, and **how it was
  observed** (browser render / DB row / manifest artifact). One row per label, no row left blank.
- **4.3** Every user-visible request is verified **in a browser** against the running stack. Objective
  form is required: *"hero cards render two-per-row at ≥1200 px viewport, verified in a browser against
  the running demo stack"* is a clause; *"the cockpit reads well"* is not.
- **4.4** **B2 is graded in FOUR parts, not one** — chat, code, document (M269) and **voice** (M271). The
  voice part is discharged by **Clause 6**, not by a working voice pipeline: M271 is a barrier and
  **NO-GO is a valid close**. State which of the four each piece of evidence covers.
- **4.5** A request whose fix landed but whose *effect* is not observable on the running stack is **RED**,
  regardless of the owning milestone's status.

### Clause 5 — every promised doc exists and describes shipped behaviour

Every **"Delivers →"** doc promised by **any** v2.10 milestone **EXISTS** and **describes shipped
behaviour** — including **`corpus/ops/demo/voice-feasibility.md`**, which is **NET-NEW** and closes this
release's only **Phase-0b blind area**.

- **5.1** The promise set is enumerated from **every** milestone `overview.md` in this release —
  M266…M271 **and** MC01/MC02, whose own promises are not knowable at this scaffold's writing. Measured at
  scaffold time, the delivery half is:

  | Milestone | Delivers → | Shape |
  |---|---|---|
  | M266 | `corpus/ops/demo/cockpit-spec.md` · `corpus/ops/demo/content-stories-spec.md` | revision |
  | M267 | `corpus/ops/seeding-spec.md` (the `p6` grant in the seed contract) | revision |
  | M268 | `corpus/ops/seeding-spec.md` (score-verdict gene) · `corpus/ops/demo/stories-spec.md` (Programs) | revision |
  | M269 | `corpus/ops/demo/playthroughs.md` (assertion boundary + count) · `corpus/ops/verification.md` (batch gate stops skipping) | revision |
  | M270 | `corpus/ops/demo/demopatch-spec.md` (new manifest — **conditional on its D-1 vehicle decision**) | revision |
  | M271 | **`corpus/ops/demo/voice-feasibility.md`** (**NET-NEW**) · `corpus/ops/safety.md` (**only if GO**) | net-new + conditional |
  | MC01, MC02 | *unknown at scaffold time — enumerate from their overviews at iter-01* | — |

- **5.2** For each promised doc: it **exists**, and a **named section** of it was **read against the
  running stack** — the surface it describes was exercised and matched. The evidence line names the doc
  section **and** the stack observation, not the commit.
- **5.3** **Conditional promises are graded on their condition, not waived.** M270's manifest promise is
  conditional on its vehicle decision; M271's `safety.md` amendment is conditional on a GO verdict. Record
  the condition's resolved value, then grade.
- **5.4** A doc that agrees with its own diff and **disagrees with the running stack** is **RED**. This is
  the clause's whole reason to exist.
- **5.5** Dangling routings introduced or inherited by this release are checked: e.g.
  `playthroughs/manifest/ai-simulations.yaml:4` points at the **dissolved M206**. Repointing belongs to the
  owning milestone; **naming it unrepointed** belongs here.

### Clause 6 — the voice verdict is recorded, with its reasoning

The **M271 voice verdict — GO or NO-GO — is recorded in the corpus with its reasoning**, and **if the
verdict touches the shared-AWS exposure, `corpus/ops/safety.md` is amended**.

- **6.1** `corpus/ops/demo/voice-feasibility.md` exists and states the verdict **explicitly as GO or
  NO-GO** — not as a summary a reader has to infer.
- **6.2** Each of M271's five blockers (B1–B5) appears with a status of **resolved** or **declared
  unresolvable**, each with its reasoning. A blocker with no disposition is RED.
- **6.3** The **data-controller decision** on B4 (a voice session's room-composite Egress writes to the
  **shared** `anthropos-livekit-test` bucket) is recorded — **who decided, when, and on what basis**.
- **6.4** **If the verdict is GO**, `corpus/ops/safety.md` carries the amendment: a GO changes what a demo
  may **reach**, and that belongs in the safety contract rather than being left implicit.
- **6.5** **If the verdict is NO-GO**, the **honest fallback is named** (presence-only, a scripted
  non-live exhibit, or an explicit "voice is not in the demo" disclosure) and **no corpus text left
  standing implies demo voice works**.
- **6.6** A NO-GO **satisfies this clause**. The clause grades *whether the verdict is recorded and its
  consequences written down*, never *which way it went*.

### Clause 7 — nothing closed on a claim this checkpoint cannot reproduce

**No milestone in this release closed on a claim this checkpoint cannot reproduce.** Any that did is
**named**, and either **re-proven** or **recorded as an honest carry**.

- **7.1** For each closed milestone (M266…M271, MC01, MC02), its **exit-gate claims** are listed and each
  is marked **REPRODUCED** (with the evidence line) or **NOT REPRODUCED** (with the reason).
- **7.2** Every NOT-REPRODUCED claim is **named in `progress.md`** — not summarised, not aggregated — and
  routed to exactly one of: **re-open the owning milestone**, **carry with a NAMED destination milestone**,
  or **drop with the reason recorded**. A carry to "the next release" or "a future pass" is **not a
  destination** (the v2.5 lesson: *a fate needs a MILESTONE, not a phase, a pass, or "the next X"*).
- **7.3** A milestone that closed **`closed-incomplete`** is listed **as such**, with the specific clause
  that was never measured clean. The precedent is live: v2.8's **M257x and M258 closed `closed-incomplete`
  by user ruling**, and **M258's clause 3 was never measured clean**.
- **7.4** The routing is written **at its destination**, not only in the closing milestone's decisions —
  the M258 lesson: *a routing written in a closing milestone's decisions is not a routing until the
  target's own doc says so.*
- **7.5** **This clause may legitimately end RED and stop the release close.** See the re-scope trigger.

## Iteration protocol

[`corpus/ops/verification.md`](../../../../../corpus/ops/verification.md) — **the two tail gates are this
checkpoint's instruments**:

1. **autoverify** — scoped, **non-fatal**, on the stack's **own offset ports**; a verify bug never blocks a
   good stack. *Which is exactly why it cannot be the only instrument here:* it printed green on
   2026-08-23 over a refused patch.
2. **the Playthrough batch gate** — the layer above, and it is **LOUD**: runs to completion under
   `D-v28-3`, one consolidated red set, bring-up exits non-zero on a non-empty set, stack left UP
   regardless.

[`coverage-protocol.md`](../../../../../corpus/ops/demo/coverage-protocol.md) supplies the **fix loop** —
sweep → triage → route by fix-surface → re-measure. Its routing table is what keeps this milestone from
absorbing fixes it should hand back.

**Per iter:** pick the clause with the most unknown left, drive its instrument against a running stack,
record the finding in `spec-notes.md`, append the closeout to `progress.md` § Running ledger, and update
the clause's row in the gate ledger. **A RED finding closes an iter successfully.**

## Why iterative (not section)

**You cannot write the fix list before you have looked.** The findings this checkpoint exists to catch are
by construction the ones nobody predicted — every one of the five on 2026-08-23 was invisible to the
milestone that shipped it. A `section` checklist written now would enumerate the failures we already know
about and would therefore measure the wrong release.

The work is also **order-dependent on its own findings**: whether Clause 4 can be graded at all depends on
Clause 1/2 producing a stack; whether Clause 7 has anything to route depends on what Clauses 1–6 return.

## Routing, not absorbing

**A failing clause routes to the milestone that owns it.** This checkpoint **diagnoses, records and
escalates**; it does not quietly repair the cluster it is grading.

- Fix belongs to a milestone **still open** → route it there and re-measure here.
- Fix belongs to a milestone **already closed** → the **re-scope trigger** fires. That is a
  **RELEASE-SCOPE finding**: surface it to the user with the honest options — **re-open**, **carry with a
  named destination**, or **drop** — rather than closing the release on a clause that did not hold.
- Fix is **in this checkpoint's own instruments** (a broken probe, a wrong denominator, a stale artifact)
  → that **is** in scope here; a mis-measuring instrument is this milestone's own defect.

## Scope

**In:**
  - Drive Clauses 1–3 against a **cold** demo cycle on the official demo host **and** a `/dev-up` on a
    local stack — the instruments, the artifacts, and the pinned denominators.
  - Grade **Clause 4** per annotation label, in a browser, with per-request evidence.
  - Grade **Clause 5** by reading each promised doc **against the running stack**.
  - Confirm **Clause 6**: the voice verdict, its blockers' dispositions, and the safety consequence.
  - Compile **Clause 7**: the reproduce-or-name ledger across every closed milestone in the release.
  - Produce the **release-level verification record** this milestone delivers.
  - Fix **this checkpoint's own instruments** when they mis-measure.

**Out:**
  - **Nothing in this release is out of observation scope. That is the point.** Every milestone, every
    doc, every annotation request is fair game for this checkpoint's instruments.
  - What **is** out is **absorbing the fixes**: repairs belong to the milestone that owns them (§ Routing,
    not absorbing).
  - Any **platform-repo edit**. A need that can only be met by a platform edit **escalates**; it does not
    edit.
  - **New scope** discovered here that serves none of the eight requests. Record it, route it, do not
    build it.

## Depends on

**MC01, MC02, M271** — every delivery milestone in the release reaches this checkpoint through one of
them. MC03 is the only milestone in v2.10 that observes the release **as a whole**, and it can only do
that once the clusters have finished.

**MC03 does not re-run MC01's and MC02's clauses clause-by-clause.** It consumes their verdicts — and any
clause **they** left unmet arrives here as a **Clause 7** named carry, not as a silent pass.

## Parallel with

**none.** The checkpoints never parallelise — each one's whole job is to observe a cluster that has
finished, and this one observes all of them.

## Open questions

Honest uncertainty, recorded here rather than resolved by invention:

- **Does the Playthrough batch gate run on the `/dev-up` path at all?** `verification.md` describes it as
  the bring-up's second tail gate, and the corpus's worked examples are demo-side. If dev has no batch,
  **Clause 3's dev half needs a named substitute instrument** — and choosing one is a decision, not an
  assumption.
- **What is "READY" on the dev path?** `build-budget.md` defines READY for `up-injected.sh` (exit 0 + green
  `autoverify.json`). The dev equivalent is unstated in the corpus; Clause 2.2 must pin it before grading.
- **Which host is "the official demo host" for v2.10, and is it the stack the reviewer drove?** The
  annotations cite `https://demo1.anthropos.work:13000`; the corpus's remote-demo runbook describes a
  tailnet MagicDNS host (`billion.taildc510.ts.net`) under `D-v28-15` (**`billion` is DEMO-ONLY**,
  `odysseus` RETIRED). The mapping between the reviewer's hostname and the documented host is **not
  measured here**, and Clause 1 needs it named.
- **Eight, ten, or seven?** The release says **eight** requests, the milestones label **ten** (A1–A5,
  B1–B3, C1–C2), the reviewer's file numbers **seven**. Clause 4.1 makes reconciling this a precondition
  rather than a footnote — but **which count the release is finally graded on is not decided here**.
- **Can Clause 1.5/2.4's browser evidence be taken on a `--public-host` stack from this machine**, or does
  it need a tailnet peer? On 2026-08-23 the suite had to be driven from a peer. If a peer is required, the
  evidence procedure must say so, and `LATENCY_SCHEME=https` / `STACK_DIR` apply.
- **Does a NO-GO on voice fully discharge Clause 4's B2 voice clause?** The design's intent is yes —
  M271 is a barrier and the verdict is the deliverable — but *"the reviewer asked for a working voice
  simulation and got a documented NO-GO"* is a **user-facing** judgement this checkpoint should surface,
  not silently resolve.
- **Where does the release-level verification record live?** This scaffold assumes it is this milestone's
  `spec-notes.md` + `progress.md` (the artifacts `/developer-kit:close-release` reads). Whether it also
  needs a **corpus** anchor is unresolved — the corpus has release-level verification narratives
  (`verification.md`, `build-budget.md`) but no per-release record.
- **What happens if Clause 1 and Clause 2 disagree** — green on demo, red on dev, or the reverse? The gate
  requires both; which one's failure is release-blocking versus carryable is a **user** call under the
  re-scope trigger.

## KB dependencies

- [`verification.md`](../../../../../corpus/ops/verification.md) — **the `iteration_protocol_ref`**: the
  two tail gates (non-fatal autoverify + the LOUD batch gate) that are this checkpoint's instruments, and
  the M236 pre-flight rung zero (*tagging is not publishing*)
- [`coverage-protocol.md`](../../../../../corpus/ops/demo/coverage-protocol.md) — the sweep → triage →
  fix-surface routing loop; the semantic believability gate; the fail-CLOSED denominator lesson
- [`playthroughs.md`](../../../../../corpus/ops/demo/playthroughs.md) — **authoritative for the live
  Playthrough count**, the 4-state reporting map, and the reset-to-seed lifecycle (Clause 3's denominator)
- [`safety.md`](../../../../../corpus/ops/safety.md) — §2.3 / §3.8: the exposure contract Clause 6 amends
  if the voice verdict is GO
- [`rosetta_demo.md`](../../../../../corpus/ops/rosetta_demo.md) — the demo lifecycle Clause 1 drives
  (`/demo-down --purge` → `/demo-up`, offset ports, the unified registry)
- [`run_guide.md`](../../../../../corpus/ops/run_guide.md) — the dev-side start + health path Clause 2
  drives
- **every v2.10 milestone `overview.md`** — M266, M267, M268, M269, M270, M271, MC01, MC02: the source of
  Clause 5's promise set and Clause 7's claim set
- [`demopatch-spec.md`](../../../../../corpus/ops/demo/demopatch-spec.md) — **added beyond the design
  brief's list**: Clause 1.4 cannot pin a denominator without §5's inventory and `TestPatchInventory`, and
  the *refused patch behind a green probe* is the release's headline failure
- [`build-budget.md`](../../../../../corpus/ops/demo/build-budget.md) — **added beyond the design brief's
  list**: it carries the only written definition of **READY** (exit 0 **AND** green `autoverify.json`) that
  Clauses 1.2/2.2 grade against, plus *state the environment with every number*

**Delivers → a release-level verification record** — the seven-clause ledger, its per-clause evidence, and
the Clause 7 reproduce-or-name routing, complete enough that a reader who was not in the room can tell
**what was proven, on which stack, on what date, and what was carried**. It is the artifact
`/developer-kit:close-release` reads.

**Delivers → the confirmation that [`corpus/ops/demo/voice-feasibility.md`](../../../../../corpus/ops/demo/voice-feasibility.md)
EXISTS** — M271 writes it; **this milestone is the one that must confirm it**, because it is NET-NEW and
closes the release's only Phase-0b blind area. A promised net-new doc that nobody checks for is the
Clause 5 failure mode in its purest form.

## Constraints (release-wide, non-negotiable)

- **Zero platform-repo edits.** Any platform source change goes through the sha-pinned **demopatch**
  mechanism ([`demopatch-spec.md`](../../../../../corpus/ops/demo/demopatch-spec.md)). A need that cannot
  be met that way **escalates**; it does not edit.
- All stack tooling lives in `rosetta-extensions`, built + tested in the authoring copy, tagged, **pushed
  to origin** (M236 pre-flight rung zero), then consumed per-stack at a pinned tag.
- Secrets handled **values-blind** — no verb reads, echoes, logs or commits a value.
- **Customer media never enters an agent's context** (`media-substrate-spec.md` PII discipline). You
  orchestrate the tooling; you do not view the media.
- **State the environment with every number.** A duration, a count or a verdict without its host, stack
  `N`, refs and date is not evidence in this milestone.
