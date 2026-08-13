---
iteration_type: tik
status: closed-fixed
created: 2026-08-11
active_strategy: TOK-01
---

# iter-04 — the host that actually exists, measured

**Type:** tik · **Active strategy:** `TOK-01` (*instrument before baseline, baseline before levers*) — step (3),
the baseline leg, re-aimed at a host that exists.

## Step 0 — Re-survey before targeting (mandatory, and it moved two things)

`TOK-01`'s `Next-tik direction` and iter-03's routing both name **odysseus**, and iter-03 exited
`user-blocker` on `DECIDE-M257-jobsim-schema-ownership`. M257 has been paused across the whole of
M257x (288 iters). Per Phase 1 Step 0 the named targets were re-verified before being committed to.
**Two of them are stale, and the substitutions are recorded here with rationale.**

| TOK-01 / iter-03 named | re-survey verdict | substitution |
|---|---|---|
| `DECIDE-M257-jobsim-schema-ownership` blocks everything | **ALREADY RESOLVED** by M257x iter-06 (`c0e075e`) | drop from the blocker list; verify the fence instead |
| baseline on **odysseus** | host **retired** (`D-v28-15`) | baseline instrumentation on the Mac mini that replaced it |
| "the Mac pays no unpack leg, so L1 is worthless here" | **REFUTED by measurement** | L1 keeps a substantial price; re-price it |

This is **not** a re-scope: `TOK-01`'s strategy (instrument → baseline → levers, largest-second-first)
holds unchanged. Only the named host and the named blocker were stale.

## Cluster / target identified

The milestone is paused on **two** things, neither of which is an agent's call: an **exit gate that
names a retired host**, and iter-03's **architectural blocker**. The re-survey shows the second is
already closed. The first is a **planning** decision (`/developer-kit:design-roadmap`), so this iter
does **not** re-cut it — it produces the evidence any re-cut needs, and a labelled proposal.

Every candidate gate — the current one, a re-cut one, or a scrapped one — needs **a measured profile of
the host that exists**. That measurement is therefore wasted under no resolution, which is why it is
this iter's target rather than a lever.

## Hypothesis

The premise that paused M257 (`D-v28-15` → `state.md`: *"a Mac is arm64/**overlay2**"*, *"the Mac pays
no unpack leg"*) is a **generalisation from the retired M1 Pro laptop to a different machine**. If the
new host runs the **containerd image store**, it pays the unpack leg, and **L1 — the milestone's biggest
lever — keeps its price**, which changes the pause's whole basis.

## Expected lift

**Zero on the primary metric, by design** — no lever is touched. The iter grades on planned
deliverables (Phase 4 Step 0: *planned scope = what the overview committed to*).

## Phase plan — three planned lines (declares the multi-step shape for the scope-creep tripwire)

1. **Re-survey** the paused-on facts against current code: the jobsim blocker, the B1/B2 fences, the host class.
2. **Measure** the host that exists and check in a host profile — `storage_driver`, and the two fields
   clause 2 consumes (`lane_heap_measured_peak_mib` **measured, never guessed**; the derived lane count).
   Ship it **without** a `gated_baseline`, per iter-03's `PROFILE-M257-…` routing.
3. **Propose** a re-cut gate to the user with the arithmetic shown. **Labelled a proposal; not adopted.**

A 4th unplanned line fires the tripwire and routes forward.

## Escalation conditions

- The host measurement contradicts a **binding release decision** (`D-v28-15`) → do **not** silently
  edit the decision; record the measurement, surface the conflict to the user. *(This fired — see the close.)*
- A full `n ≥ 3` cold campaign cannot be run honestly (contended box / the user's stacks resident) →
  say so and route it, rather than publishing a contended number as a baseline.

## Acceptable close-no-lift outcomes

Finding the host-class premise **confirmed** (no unpack leg) would also have satisfied this iter: it
would have converted the re-scope question from open to decided, on evidence.
