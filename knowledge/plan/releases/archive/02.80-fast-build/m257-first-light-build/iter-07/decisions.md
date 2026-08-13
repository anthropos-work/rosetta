# iter-07 — decisions

## D1 — the user's stacks stay up, because the memory math says they can

**Decision.** Run the campaign on the free `demo-1` slot with `demo-2` (11 containers, user-validated) and
the 5-container dev stack **left running**.

**Why this is not politeness.** The sub-agent brief permits bringing the dev stack down after a heartbeat,
so the question is whether the resources are actually needed. Measured at open: VM `MemAvailable`
**9.87 GiB** of 11.67; all 16 running containers total **~1.43 GiB**; one Next.js build lane costs
**~3.1 GiB** measured (`lane_heap_measured_peak_mib`); demo-1's own containers ≈ 1.3 GiB. Projected peak
≈ **5.9 GiB** against 9.87 available. Disk: **51.6 GiB** free against a 22 GiB floor+projected.

Tearing down a working stack the user is validating on, to buy headroom the arithmetic says is already
there, would be a cost with no purchase. The heartbeat was still written *before* the campaign started, so
the decision is auditable either way.

## D2 — killed my own test suite mid-campaign: an instrument must not live inside its own subject

**Decision.** ~19 minutes into a full `pytest tests/` sweep — launched before the campaign and still
running — the campaign started, and the sweep was **killed**.

**Why.** The campaign measures wall-clock on a shared box. A full guard sweep is a sustained CPU consumer,
so leaving it running would have made a measurable fraction of the "contention" in this baseline *my own
harness's*. This milestone already carries the lesson in the abstract — *an instrument that lives inside
its own subject measures itself* — and this is its concrete form: **contention I cause is contention I can
remove, and only unremovable contention belongs in a labelled baseline.**

**What is lost, stated rather than papered over.** The full-repo sweep did not complete, so this iter's
evidence is the **targeted** suites from iter-06 — `test_buildbench` **120 passed**,
`test_baseline_mirror_fence` **28**, the five plan-doc guards **116/2 skipped**, the three README guards
**63** — which are the suites that cover the touched files. A full sweep is worth re-running on a quiet
box; it is not worth corrupting the one measurement this milestone has been waiting five iters for.

## D3 — the corpus edit is CORRECT and could not be committed; the tree is left dirty deliberately

**Situation.** `corpus/ops/demo/build-budget.md` (the milestone's own `iteration_protocol_ref`) was
corrected — the retracted `overlay2` claim, two now-fixed banner items, and two line citations that
**iter-06's own code change had broken**. The pre-commit `repair-postcondition` fence blocks the commit on
**two sites in a file this iter never touched**: `corpus/architecture/platform-migration-status.md:121-122`.

**Verified, not assumed, before concluding anything:**

| the guard's complaint | checked against | result |
|---|---|---|
| `app/main.go:1487` is a closing `}` | `stack-demo/app/main.go` **and** `stack-dev/app/main.go`, working tree **and** `git show HEAD:` | **all four read `Sender: msgsender.NewFromEnv(logger)`** — the citation is CORRECT |
| `docker-compose.yml:168` out of range, *"file has 164 lines"* | both clone roots, working tree and HEAD | **186 lines** in every case |

`anchor_construct_guard._clone_roots()` returns exactly `[stack-demo, stack-dev]`, and **both** support the
corpus text. `git status` confirms `platform-migration-status.md` is unmodified by this iter. So this is a
**pre-existing guard-resolution disagreement, not a corpus defect** — and the corpus prose it flags is
demonstrably right.

**What was NOT done, and why each was rejected:**

- **`--no-verify`** — forbidden outright, and it is the exact move that lets a fence become theatre.
- **"Fix" the two sites** by editing prose that is *correct* against both clone roots — that would damage
  accurate documentation to satisfy a resolver, i.e. make the corpus wrong to make a guard green.
- **Revert my own edit** — a forbidden op, and it would discard a correct fix to a doc that currently
  carries a retracted claim.
- **Add them to the guard's baseline** — the fence's own contract forbids it: *"a repair may remove these;
  it may never add one."*

**So the working tree is left with `corpus/ops/demo/build-budget.md` modified and uncommitted, and that is
reported rather than tidied.** The orchestrator and the user are the only deciders on uncommitted state.
Routed as `FIX-M257-anchor-guard-resolution` — the guard resolves two anchors to content neither of its own
declared clone roots contains, which is a defect in the *instrument*, and this milestone has now found three
of those in one session.
