---
milestone: MC02
title: "Checkpoint — actually consumable"
milestone_shape: iterative
status: planned
release: "02.100-content-consolidation"
exit_gate: "on BOTH a real demo stack and a local dev stack: a seeded hero starts AND completes a simulation in each of chat, code and document modality with real session rows written; no hero on any seeded org hits ERROR_JOB_SIMULATION_LIMIT_REACHED; /library/skill-paths paints a loading affordance before content on a cold cache-disabled load, at a measured time; the Playthrough batch gate RUNS and reports GREEN on the DEFAULT /demo-up path with the word 'skipped' absent from its verdict; every clause proven on both stacks, not one; and playthroughs.md, verification.md, demopatch-spec.md and seeding-spec.md describe the shipped behaviour, read against the RUNNING stacks."
iteration_protocol_ref: "corpus/ops/demo/coverage-protocol.md"
re_scope_trigger: "If clause 4 cannot be met because BIND_HOST proves larger than M269 scoped it, ESCALATE rather than silently accepting a skipped gate — a gate that reports nothing is the exact failure this whole checkpoint layer exists to catch, and accepting it here would reproduce inside the checkpoint the defect the checkpoint was built to find."
depends_on: "M267, M269, M270"
parallel_with: "none"
complexity: large
last_updated: "2026-08-24"
---

# MC02: Checkpoint — actually consumable

**CHECKPOINT milestone.** It builds nothing of its own. It grades a **cluster** — M267 «The entitlement
unlock» + M269 «Modality playthroughs» + M270 «Skill-paths first paint» — against a **running stack** and
against **its own docs**, and it runs only when all three have closed.

**Cluster under test:** everything about whether a hero can actually **CONSUME** content.

**Goal:** on **BOTH** a real demo stack and a local dev stack, a seeded hero starts and finishes real
simulations, the library paints honestly, and the release-level gate actually **RUNS**.

> ⚠️ **A CHECKPOINT MAY FAIL, AND FAILING IS ITS PURPOSE — NOT A DEFECT.**
> The admissible outcomes of MC02 are **PASS** and **work sent back to M267 / M269 / M270**. A checkpoint
> that has never returned a clause RED is a checkpoint that is not measuring anything. Sending work back
> is the milestone succeeding at its job; quietly widening a clause so it goes green is the milestone
> failing at it. If a clause cannot be met, say so, name the milestone it goes back to, and stop — do
> **not** re-word the clause.

## Why these checkpoints exist

**Every one of the failure modes below was MEASURED on 2026-08-23, on the live `demo1` stack, in one day:**

> - **autoverify** printed `verify live: all liveness + readiness probes passed` while a **demo-patch had
>   been REFUSED** and the hiring bundle shipped with **every PostHog flag OFF**.
> - **`/api/health-check` answered 200** on a studio-desk container whose **every gated page returned
>   HTTP 500** — the route is public **BY DESIGN** and **structurally cannot witness that failure**.
> - **The Playthrough batch gate recorded `skipped`, never green**, because the stack was
>   `--public-host` (`BIND_HOST` / `D-M255-7`) — so the release-level gate **reported nothing at all** and
>   **looked fine**.
> - **The studio Playthrough reported PASS** and was **cited as evidence the migrated studio works**; it
>   matches **EMPTY SCAFFOLDING at +2.1 s** (`FIX-M256-studio-false-green`).
> - **The manager menu was missing "Build a Course"** and pointed **Assign at a legacy surface**, on a
>   stack **every probe called healthy**.

**THE COMMON SHAPE:** a milestone closes, its tests pass, its probes are green — **and the thing it
delivered does not work on a real stack, or works but is documented as something else.** A checkpoint
milestone is the layer that grades the **CLUSTER** against a **running stack** and against **its own
docs**, rather than against its **diff**.

**SO THE GATE IS ALWAYS TWO-SIDED: it works on a real stack AND the corpus describes what actually
shipped.** Read the doc against the **RUNNING STACK**, never against the diff that changed it — **a doc
can agree with a commit and disagree with reality.**

## Exit gate — six clauses, each independently checkable

Each clause is graded **PASS** / **RED** on its own. **The milestone passes only when all six are PASS on
the same pair of stacks in the same pass.** A clause is PASS only if someone who was not in the room could
re-run the stated check and reach the same verdict from the stated evidence.

### Clause 1 — a hero STARTS **and COMPLETES** a simulation in three modalities

A **seeded hero** STARTS **and COMPLETES** a simulation in **each of chat, code and document** modality,
with **real rows written to `public.job_simulation_sessions`** — proven by **row counts taken before and
after** the run, **not** by a launch screen.

- Evidence required: the **before** count, the **after** count, and the **delta ≥ 1 per modality**,
  recorded per stack.
- A launch-screen render, a `/sim/<slug>/start` URL, or a green Playthrough with **no row delta** is
  **RED** — that is exactly the boundary M269 exists to move
  (`aisim-chat-launch.spec.ts` measured **0 rows** at that boundary).
- **Voice is explicitly OUT.** It is **M271**, and its verdict may be NO-GO. A voice modality appearing
  in this clause is a scoping error.

### Clause 2 — no seeded hero hits the entitlement ceiling

**No hero on any seeded org** hits `ERROR_JOB_SIMULATION_LIMIT_REACHED`, verified by **attempting a start
as at least one hero per seeded org**.

- Evidence required: the **enumerated list of seeded orgs** on the stack under test, and **one attempted
  start per org**, each recorded with its outcome.
- Testing one org and inferring the rest is **RED** — M267's fix is a **per-org `p6` row** and matcher
  `m6` **has no `default` escape**, so the failure is per-org by construction.
- The UI string `errors.simulationLimitReached`
  (`AISimulationStartWithoutSession.tsx:209`) appearing anywhere in the pass is **RED**.

### Clause 3 — `/library/skill-paths` paints honestly, on a cold load, at a stated time

`/library/skill-paths` shows a **LOADING AFFORDANCE** and **then** content — **never an empty region that
later fills**.

- Verified **in a browser**, on a **COLD load** with **cache disabled**, and **stated as a measured time**.
- Evidence required: (a) the affordance is **present in the first paint**, (b) content follows, (c) a
  **number** — time to first affordance and time to content — recorded per stack, with the browser and
  viewport stated.
- **An empty region that later fills is RED even if it fills fast.** Speed is not the clause; honesty is.
  The measured time is recorded so the clause is comparable across passes, not so a fast empty flash can
  buy a pass.
- A `curl` or a bundle grep does **not** satisfy this clause. Neither witnesses a paint order.

### Clause 4 — the batch gate RUNS and is GREEN on the DEFAULT `/demo-up` path

The **Playthrough batch gate RUNS and reports GREEN** on the **DEFAULT `/demo-up` path** — **the word
`skipped` must not appear in its verdict.**

- Evidence required: the verdict artifact itself, with the **literal absence of `skipped`** and the
  **presence of a green result**, from a bring-up run with **no non-default flags**.
- **This is the `BIND_HOST` / `D-M255-7` clause, and it is the one most likely to fail — it has been
  deferred THREE times** (M255 → M256 → M258, never worked).
- **A gate that reports nothing is not a pass.** `skipped` is **RED**, and so is a run that reaches green
  only by removing `--public-host`, because the default path **is** the path under test.
- See the re-scope trigger: if this cannot be met, **escalate**; do not accept the skip.

### Clause 5 — proven on a demo stack **AND** a dev stack

**Every clause above** is proven **on a demo stack AND on a dev stack.** **A pass on only one is not a
pass.**

- The two paths **build studio-desk and seed differently** — which is **exactly how divergence hides**.
- Evidence required: **two** evidence sets, each naming its stack (`demo-N` / `dev-N`), the host it ran
  on, and the date. Clauses 1–4 are graded **twice**, and a clause that is PASS on one stack and RED on
  the other is **RED**.
- Reusing one stack's numbers for the other is **RED** by definition, not by degree.

### Clause 6 — the corpus describes what actually shipped

`corpus/ops/demo/playthroughs.md`, `corpus/ops/verification.md`,
`corpus/ops/demo/demopatch-spec.md` and `corpus/ops/seeding-spec.md` **describe the shipped behaviour,
read against the RUNNING stacks.**

- Evidence required: per doc, the **claims checked** and the **observation on the running stack** each was
  checked against.
- **Read against the RUNNING STACK, never against the diff that changed it.** A doc that agrees with the
  commit and disagrees with the stack is **RED** — that is the second half of the common shape above.
- Specifically in scope because the cluster moves them: the **live-Playthrough count and the 4-state
  reporting map** (`playthroughs.md`), the **batch-gate skip contract** (`verification.md`), any patch the
  cluster added or re-pinned including whether it **applied or was REFUSED** (`demopatch-spec.md`), and
  the **`p6` / entitlement seeding contract** (`seeding-spec.md`).

## Iteration protocol

[`coverage-protocol.md`](../../../../../corpus/ops/demo/coverage-protocol.md) — the sweep → triage → fix
loop. Each iter: **measure the clauses against a running stack**, triage every RED to the milestone that
owns it, and record the routing. The protocol's own discipline applies: **fail-CLOSED** — an **empty
ledger is a FAILURE, not a 0/0 pass** — and a denominator is **pinned** before the sweep, never derived
from what happened to be found.

**MC02's fix surface is deliberately narrow.** A checkpoint's normal output for a RED is a **routing back
to M267 / M269 / M270**, not a fix authored here. Where MC02 does act directly, it is on **its own
evidence-gathering tooling** and on the **doc side of clause 6**.

## Why iterative (not section)

The deliverable is a **verdict on a cluster measured against running stacks**, and the work is a **loop**:
measure → triage → route → re-measure. **An `In:` list of deliverables would presume which clauses fail.**
Which clause is worth probing next depends on what the previous measurement returned — clause 1 cannot be
graded at all until clause 2 is PASS (a hero who cannot start cannot complete), and clause 5 doubles
whichever clauses are live. The scope below therefore enumerates **measurements and routings**, not
products.

## Scope

**In:**
  - **Bring up BOTH stacks and hold them for the pass** — one demo stack on the **default** `/demo-up`
    path (clause 4 depends on it being default) and one dev stack. Record host, refs, and date for each.
  - **Grade clause 2 first.** It is the cheapest and it **gates clause 1**: a hero blocked by
    `ERROR_JOB_SIMULATION_LIMIT_REACHED` cannot produce a session row, so a clause-1 RED measured under a
    clause-2 RED tells you nothing about M269.
  - **Grade clause 1** — three modalities × two stacks, with before/after row counts per run.
  - **Grade clause 3** in a browser, cold, cache disabled, with a stated number per stack.
  - **Grade clause 4** on the default path, and **escalate** rather than accept a `skipped` verdict.
  - **Grade clause 6** — the four docs, read against the running stacks, claim by claim.
  - **Route every RED** to the owning milestone, in writing, in [`decisions.md`](decisions.md). A routing
    is not a routing until the **target's own doc** says so — the failure that fired ≥3× in v2.8.
  - **Doc fixes for clause 6** where the shipped behaviour is right and the doc is wrong.

**Out:**
  - **Voice.** It is **M271**, and its verdict may be NO-GO.
  - **The Programs / cockpit surfaces.** They are **MC01**.
  - **Authoring the cluster's fixes.** A RED routes back to M267 / M269 / M270. MC02 grades; it does not
    re-implement the milestone it just failed.
  - **Re-wording a clause so it passes.** Named here as out-of-scope because it is the specific way a
    checkpoint layer dies.

## Depends on

**M267, M269, M270** — **all three closed.** MC02 runs when the cluster is complete, not alongside it:
- **M267** makes a start possible at all (clause 2, and clause 1 through it).
- **M269** makes completion assertable and is where `BIND_HOST` / `D-M255-7` lands (clauses 1 and 4).
- **M270** is clause 3.

Running MC02 against a partial cluster measures the gaps, not the delivery.

## Parallel with

none

## Open questions

Honest uncertainty, recorded here rather than resolved by invention:

- **What is the session table's real name and schema?** This brief states
  **`public.job_simulation_sessions`**; M269's own scope and the M256 probe header both say
  **`jobsimulation.sessions`**. The jobsim-in-app fold moved 23 run-state tables to `public`, and the
  legacy `jobsimulation` schema drop is a **still-pending M810 step** — so **both names may resolve on a
  live stack, against different data.** **Measure the name before grading clause 1**; a row count against
  the legacy husk would read as a clean RED and be a measurement error.
- **Does landing `BIND_HOST` actually un-skip the gate?** M269 records **two independent causes for one
  symptom** — the bind, and the fact that a connection from the demo host to its own tailscale IP hits the
  kernel socket and **bypasses `tailscale serve`**, which terminates TLS. Fixing the bind may leave the
  skip in place. **Clause 4 grades the symptom, not the cause**, which is deliberate — but it means a RED
  here does not by itself tell M269 which cause to chase.
- **What counts as "COMPLETES" per modality, and is it uniform?** M269's own open questions concede that
  reaching a *result* may differ per modality (code and document have deterministic substrate; chat does
  not). Clause 1's oracle is a **row**, which is uniform — but whether a written row means the hero
  *finished* or merely *entered* is **not settled**, and the clause is only as honest as that answer.
- **Does the dev stack have seeded orgs and heroes at all in the shape clause 2 assumes?** The main dev
  stack (`N=0`) is **deliberately never set-dressed** (`dev-setdress.sh` hard-refuses `N=0`), and the dev
  and demo seed paths differ. Which dev stack MC02 grades — and whether the dev half of clause 5 needs a
  `dev-N`, N ≥ 1 — is **unresolved at scaffold time**.
- **Is the batch gate even applicable on a dev stack?** Clause 5 says every clause is proven on both. The
  batch gate is documented on the `/demo-up` path. Whether clause 4 has a meaningful dev-side reading, or
  whether it is demo-only by construction and clause 5 admits a stated exception, is **open** — and it
  must be answered **in writing**, not by quietly grading clause 4 once.
- **How does MC02 avoid being the next false green?** It is a gate that grades gates. Nothing in this
  scaffold prevents MC02 itself from reporting PASS against evidence that does not witness the failure —
  the `/api/health-check` shape. The mitigation is that **every clause names its evidence artifact**, but
  whether that is sufficient is **not proven** and should be revisited at the first PASS.

## KB dependencies

- [`playthroughs.md`](../../../../../corpus/ops/demo/playthroughs.md) — the manifest model, the
  live-Playthrough count and the 4-state reporting map (clauses 1 and 6)
- [`verification.md`](../../../../../corpus/ops/verification.md) — the **two tail gates** and the batch
  gate's skip contract (clauses 4 and 6); also the source of the *"UP means UP, and every journey
  verified"* claim this checkpoint is testing
- [`coverage-protocol.md`](../../../../../corpus/ops/demo/coverage-protocol.md) — the
  `iteration_protocol_ref`: sweep → triage → fix, fail-CLOSED, pinned denominator
- [`demopatch-spec.md`](../../../../../corpus/ops/demo/demopatch-spec.md) — the 7 guards and the
  **REFUSED-patch** failure mode named in the WHY box (clause 6)
- [`seeding-spec.md`](../../../../../corpus/ops/seeding-spec.md) — the seed blueprint and the entitlement
  seeding contract (clauses 2 and 6)
- [`run_guide.md`](../../../../../corpus/ops/run_guide.md) — bringing the dev half of clause 5 up

**Delivers → corrections to the four clause-6 docs**, where the shipped behaviour is right and the doc is
wrong. **No net-new doc is planned**; if a clause proves to have no doc anchor anywhere, that is itself a
finding to record, not a silent pass.

**Delivers → a written verdict per clause, per stack**, in [`progress.md`](progress.md), with every RED
routed in [`decisions.md`](decisions.md).

## Constraints (release-wide, non-negotiable)

- **Zero platform-repo edits.** Any platform source change goes through the sha-pinned **demopatch**
  mechanism ([`demopatch-spec.md`](../../../../../corpus/ops/demo/demopatch-spec.md)). A need that cannot
  be met that way **escalates**; it does not edit.
- All stack tooling lives in `rosetta-extensions`, built + tested in the authoring copy, tagged, **pushed
  to origin** (the M236 pre-flight rung zero — *tagging is not publishing*), then consumed per-stack at a
  pinned tag.
- Secrets handled **values-blind** — no verb reads, echoes, logs or commits a value.
- **Customer media never enters an agent's context.** You orchestrate the tooling; you do not view the
  media.
