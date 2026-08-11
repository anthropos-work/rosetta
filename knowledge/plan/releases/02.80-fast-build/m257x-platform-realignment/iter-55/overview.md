---
iter: 55
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-03
---

# iter-55 — re-establish the ref baseline (clauses 2 and 1), against a platform that moved again

**Active strategy reference:** `TOK-04` — *pin the target, or stop calling it a measurement.* This iter is
TOK-04's own `Next-tik direction` for iter-55, executed with all four policies binding:

- **P1** — every measurement states its refs in the artifact, at the moment taken.
- **P2** — every instrument is a committed file.
- **P3** — the platform ref is chosen (origin HEAD), recorded, and re-checked at close.
- **P4** — derive, else fence, else declare prose-under-review.

## Step 0 — re-survey (mandatory), and what it found

TOK-04 named the target ref `ef32d4c`. **It is already stale.** Measured at iter open:

    origin/main   0dab54d  chore(compose): run without the standalone storage; rename graphql -> core
    local clone   ef32d4c  (1 behind)

**Per P3 the clone was re-pointed to `0dab54d` before any measurement was taken** (`git merge --ff-only`,
clean tree, no platform edit). This is not a re-scope: TOK-04's strategy is unchanged and its pre-registration
still stands. Only the ref moved — which is precisely the event P3 exists to handle, and the second time in
this milestone that the detecting iteration has had to re-point inside itself.

`0dab54d` is the **v9.0 `support-in-app`** step (storage + messenger fold, PRs #1096/#1098/#1103) that
`platform-alignment.md` §1 already named as IN FLIGHT. It changes three things that bear on the gate:

| change | consequence for our tooling |
|---|---|
| `profiles: [graphql, …]` → `profiles: [core, …]` on `backend` + `gotenberg` | every rext site naming the profile literal `graphql` selects an **empty service set** |
| `storage` → `profiles: [storage-legacy]`, out of the default set | the default stack loses a container; `verify_svcs` names a service that will not exist |
| `messenger` dropped from `all` | (not in the demo's set; recorded, not acted on) |

## Cluster / target identified

TOK-04 directed iter-55 at *"re-run clause 2 (~5 min) then clause 1 (3 cycles, ~35 min) against `ef32d4c`."*
The re-survey substitutes the ref (`0dab54d`) and adds a **precondition the direction could not have known
about**: clause 1 cannot be run at all until the profile rename is followed, because
`up-injected.sh` brings the stack up with a hard-coded `--profile graphql`.

That precondition is not incidental work. It is **the milestone's founding class, occurring again**: a
hand-maintained tuple naming platform topology, which nobody updates when the platform moves
(`platform-alignment.md` §2). Two such tuples are in the bring-up path:

- the **profile literal** — 5 live code sites;
- **`verify_svcs`** in `up-injected.sh:2596` — a hand-written service list still naming `jobsimulation`,
  `cms`, `roadrunner` (deleted at `ef32d4c`), `graphql` (deleted at `2adcf71`) and `storage` (defaulted out
  at `0dab54d`).

So the target is: **derive both tuples from the platform's own compose file (P4 first branch), then take the
two measurements against `0dab54d` with `refs:` blocks.**

## Hypothesis

1. Deriving the profile name from the `backend` service's own `profiles:` list makes every rext site
   ref-independent — it yields `graphql` at `ef32d4c` and earlier, `core` at `0dab54d`, with no human action
   on the next rename.
2. With that landed, clause 1's three cold cycles reach `autoverify green:true / 0 warnings` and clause 2's
   suite reads 30 live / 0 failing / 0 error.

TOK-04 **pre-registered both green**. That pre-registration was made against `ef32d4c` and is now testable
against a ref that moved underneath it — which makes it a sharper test, not a spoiled one.

## Expected lift

Gate 2 of 5 → 4 of 5. Clause 5 is untouched by this iter and is not re-cut.

## Phase plan

- **A — derive.** Replace the profile literal at its 5 live sites with one derived definition; derive
  `verify_svcs` from the resolved compose service set. Unit-test both, including the ref-independence claim.
- **B — clause 2 control.** A reading on the still-running 44-hour-old `demo-1` (built at platform
  `28c5f0d`, three folds ago). Cheap, and it is a **control**, not a restoration — recorded as such.
- **C — clause 1.** Three consecutive cold `demo-down --purge` + `demo-up` at `0dab54d`. Per
  `platform-alignment.md` §5 rule 15, **log which path each cycle took** through the nondeterministic
  Directus bootstrap race.
- **D — clause 2 binding.** The suite on the cycle-3 stack.
- **E — close.** `refs:` blocks on both measurements; re-check origin per P3; commit; re-pin + push the tag.

## Escalation conditions

- A cold cycle fails for a reason that needs a platform edit → **user-blocker** (v2.8's zero-platform-edit
  constraint is not negotiable inside this iter).
- The derivation cannot be made ref-independent without PyYAML → land the literal `core` **plus a fence**
  (P4 second branch) rather than silently hand-maintaining a third tuple.

## Acceptable close-no-lift outcomes

A cold cycle going red for a *newly-discovered* reason is a first-class outcome and, per TOK-04, **the
highest-value result available** — the pre-registration exists to be refuted. What is not acceptable is
reporting a green taken against a stale ref, or against a stack built three folds ago.
