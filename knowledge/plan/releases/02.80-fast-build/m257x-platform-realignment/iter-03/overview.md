---
milestone: M257x
iter: 03
iteration_type: tik
status: in-progress
date: 2026-07-31
---

# iter-03 — stand the gate's instrument up: a `stack-demo` workspace on the new Mac

**Type:** tik, under `TOK-01` **step 1** — *"Unblock the gate's instrument… every other clause is measured
through a bring-up."*

## Step 0 — re-survey

iter-02 retired TOK-01's *named* step-1 target (the rext pin was already clean — `D-M257x-4`) but step 1
itself was never satisfied: the instrument was blocked by **no container runtime**. Re-measured now:

| | at iter-02 open | now |
|---|---|---|
| container runtime | **none** (docker · podman · colima · nerdctl · lima · orbstack all absent) | **Docker 29.6.2**, linux/arm64, overlayfs, 8 cpus, 12528664576 B — independently verified |
| `stack-demo/` workspace | absent | **still absent** ← the remaining blocker |
| rext pin | `fast-build-m257x-iter-01` | `fast-build-m257x-iter-02`, **verified on origin** |
| secrets source | unmeasured | `.agentspace/secrets/` — 5 repos; `platform/.env` has 15 keys incl. **`GH_PAT`** (values never read) |
| GitHub SSH | unmeasured | authenticates as `kiralise` |
| disk | — | 382 GiB free |

So step 1 is, for the first time in this milestone, **executable**. That makes it the target.

## Active strategy reference

`TOK-01` step 1. Its ordering rationale is unchanged and is now the binding constraint: clauses 1, 2 and the
live half of 4 are all measured *through* a bring-up, so nothing downstream can be honestly attempted until
one exists.

## Cluster / target identified

`HOST-M257x-stack-demo`, routed by iter-02. `demo-stack/ensure-clones.sh` bootstrap-clones
`stack-demo/platform` from GitHub and `make init`s the peer repos, so a box with only `stack-demo/` can
bring a demo up end-to-end (v1.8 "understudy" M26) — but the workspace does not exist yet.

## Hypothesis

With Docker present, SSH working and `GH_PAT` available, a `stack-demo` workspace can be bootstrapped and the
bring-up driven far enough to establish **the first honest cold-instrument measurement on this host** — which
is both TOK-01 step 1 and the raw material M257 (paused) needs for its speed gate.

## Expected lift

**No gate clause flips.** Clause 1 requires **3 consecutive** cold `demo-down --purge` + `demo-up` cycles at
`autoverify green:true / 0 warnings`; a first bring-up is the precondition, not the clause. Honest expected
metric delta: **0/5 → 0/5**.

The measurable sub-progress is binary and real: **a working instrument exists, or it does not, and we know
precisely where it stops.** Given this milestone's history — B1 and B2 were both invisible for four days
because nobody ran a cold cycle — a characterised stopping point is a first-class deliverable.

## Phase plan

1. **Phase 0d pre-flight (cheap, decisive, before any long build).** Run the bring-up's own host pre-flight
   and the pin guard. If they refuse, that is the iter's finding and no 40-minute build is wasted.
2. Bootstrap `stack-demo/` via `ensure-clones.sh` (its sanctioned path), at the pinned rext tag.
3. Provision secrets **values-blind** into the workspace.
4. Drive the bring-up as far as it goes; capture where and why it stops.
5. Record the cold cost profile (clone / build / migrate legs) — the honest baseline this release is about.

## Escalation conditions

- The pre-flight or pin guard refuses → characterise and close; do **not** brute-force past a FATAL guard.
- A build failure that needs a **platform-repo edit** → hard stop, user-blocker (v2.8 constraint: 0 platform
  edits; the sanctioned escape is a sha-pinned `demopatch` or an rext-owned file).
- Budget exhaustion mid-build → the tracked tree stays clean by construction (`stack-demo/` is git-ignored),
  so this closes as a characterised partial rather than a dirty-tree blocker.
- The known non-fatal `< 12 GiB` VM-RAM warning is **expected noise** (`FIX-M257x-vmram-gib-unit`, routed) —
  it is a decimal-GB-vs-binary-GiB unit mismatch, not a real headroom failure. Do not re-derive it.

## Acceptable close-no-lift outcomes

Characterising exactly where a cold bring-up stops on this host — with the failing step, its cause, and
whether it is host, tooling or platform — is a complete iter. That is precisely the measurement whose absence
let a total breakage sit undetected for four days.
