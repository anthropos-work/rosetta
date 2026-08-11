---
iter: 265
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-10
active_strategy: TOK-08
route: FIX-M257x-264-cms-md-past-tense-dependency
---

# iter-265 — the requirement migrated; its documentation stayed with the corpse

**Type:** tik, under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07)
(*census the mechanical classes; stop sampling them*).

## Step 0 — re-survey (mandatory, before targeting)

`TOK-08` names no specific next target; the controlling evidence is iter-264's close, which routed
`FIX-M257x-264-cms-md-past-tense-dependency` with the milestone's own rule attached: *a class is closed by
an enumeration, never by its last member* (§8, iter-169).

Re-surveyed at open, corpus `cf9c469`:

- `corpus/services/cms.md:271` — still present, still past tense, still under the decommissioned service.
- `corpus/ops/staging-bringup.md:428` — still instructs the operator to **comment OUT** the studio `COPY`
  lines, which is the inverse of the live `app` requirement iter-264 established.
- **And the first grep already exceeded the routed pair:** `corpus/ops/staging_from_dump.md:475` and
  `corpus/ops/setup_guide.md:772` both instruct *"Edit `cms/Dockerfile.dev` and remove"* the studio lines —
  the second of those in **the very file iter-264 repaired**.

The target is live and is larger than the two sites the route named. Proceeding.

## Cluster / target identified

**The class:** an *operational requirement* that migrated when a service was folded into `app`, whose
**documentation stayed attached to the decommissioned service**. Nothing errors — the sentence still names
a real (frozen) repo, so it reads as correct history rather than as a live instruction pointed at a corpse.

This is the intra-corpus face of `platform-alignment.md` §5's *"a named-consumer list survives the merge
that moved the consumer"* (iter-23), which was written about **rext's** consumer lists. iter-264 found it
occurring **inside this corpus**. This iter asks how big it is.

## Hypothesis

Enumerating the mechanically-decidable slice — **every corpus site that issues an operational instruction
naming a decommissioned service repo** — will find defects beyond the routed pair, concentrated in the
**ops guides** (where an operator actually goes) rather than the service-doc redirects (which are already
marked historical). The repair is per-site; the *close* is a fence that keeps the population enumerated.

## Expected lift

Limb 3 of the user's binding closing condition — *the corpus reflects the working stack*. A live setup
guide that tells a new engineer to delete the very lines the build requires is the sharpest possible
counterexample to that limb.

## Phase plan

1. Seal pre-registrations (this iter's first commit) — **before** the enumeration runs.
2. Build the enumeration instrument; run it; publish the population **with its denominator**.
3. Grade every member: `historical-marked` / `dead-unmarked` / **`migrated-contradicts-live`**.
4. Repair the `migrated-contradicts-live` members.
5. Fence the population so the class stays enumerated (§8, iter-169/iter-176).
6. Re-run the guard family; close.

## Escalation conditions

- If the enumeration finds the population is **studio-only** (PR-5 refuted), the route was a single-instance
  route and the iter closes by saying so — no fence is built for a population of one.
- If an existing guard already fires on class members (PR-4 refuted), the defect is an **unrun check**, not
  a missing fence — the fix shape changes to *make it run*, exactly as iter-263 re-shaped iter-262's.
- If a repair would require a platform-repo edit → stop; v2.8 constraint holds (0 platform edits).

## Acceptable close-no-lift outcomes

A documented falsification of PR-5 (class is studio-only) or of PR-2 (population is trivially small) is a
complete iter: it converts a routed *class* into a closed *instance* and removes it from the backlog with
evidence rather than by attrition.
