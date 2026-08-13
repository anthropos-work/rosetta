---
iter: 17
milestone: M257x
iteration_type: tik
status: closed-no-lift
opened: 2026-08-01
---

# iter-17 — re-prove gate clause 1 with an instrument that can see content

**Active strategy reference:** `TOK-01: instrument first, then follow` (milestone-root `decisions.md`),
step 5 — *"prove it cold, three times"* — and step 3, *"land the fences before trusting any green."*
iter-16 landed the last of the fences this clause depends on; this iter is the proof they were the
precondition for.

## Step 0 — re-survey (mandatory)

Re-measured at open, and two of the three inputs had to be confirmed rather than assumed:

- **Platform origin HEAD is `2adcf71`, unchanged** since iter-14 (`git fetch` + compare, not memory).
  No second alignment-invalidating commit landed during iter-16, so the **re-scope trigger stays at
  occurrence 1 of 2**. The gate's own wording — *"against platform @ **origin HEAD**"* — is satisfiable
  by the clone in use.
- **The consumption pin was stale and is now correct.** `.agentspace/rext.tag` and the `stack-demo`
  consumption clone both read `fast-build-m257x-iter-15`, which predates the harden pass. Both re-pinned
  to `fast-build-m257x-iter-16` (`c63d981`), verified on origin, and `git merge-base --is-ancestor`
  confirms it **contains** `46f8cc3` (harden-1) — so `probe_directus_serves_content` is in what a stack
  actually consumes. Without this the whole iter would have re-run the original mistake.
- **`demo-1` is UP (15 containers) but is not a valid instrument.** Its world is mutated pt-world and its
  Directus was repaired by hand and a `/tmp` binary, not by the bring-up. It is usable as a *positive
  control for the probe* and for nothing else.

## Cluster / target identified

Gate clause 1's three checked-in verdicts (`evidence/av-cycle{1,2,3}.json`) are compromised. They are not
*wrong* about what they measured — `autoverify` really did return `green:true` — but the only Directus
check in that instrument counted rows in the `directus_collections` **registry** table, which is a claim
about what is *registered*, not about what is *served* (protocol §5 rule 14: **REGISTERED is not SERVED**).
The bring-up transcript said `directus=skipped(error)` on all three cycles at the same time.

Two changes since make a re-prove meaningful rather than ceremonial, and they are independent:

1. **iter-15 fixed the replay** (`rc=1 → rc=0`, 0 → 11,986 rows, anon `403 → 200`) — but **at its own
   layer only**. It has never run inside `demo-up`. The auto-provision path it goes through fires on a
   bootstrapped-GAP schema, which is exactly what a *purged* stack has and what the current demo-1 no
   longer does.
2. **The harden pass added `probe_directus_serves_content`**, which selects a non-`directus_*` collection
   the stack's own Postgres says holds rows, then asks the running Directus for an item over HTTP — and
   fails distinctly on 403 (holds content, serves it to nobody) and on `200 + "data":[]` (served-but-empty).
   It cannot satisfy itself (§5 rule 7): the collection comes from one subsystem and the answer from
   another.
3. **iter-16 made the bring-up's own verdict honest** — a failed replay now reads `set-dress INCOMPLETE`
   and exits 3 instead of printing `set-dressed` and exiting 0.

## Hypothesis

A cold `demo-down 1 --purge` → `demo-up 1` on the re-pinned tooling reaches `green:true / 0 warnings`
**with the serving probe active and passing**, three times consecutively — and the bring-up transcript
reports a *complete* set-dress rather than `directus=skipped(error)`.

**The hypothesis is genuinely falsifiable, and both failure modes are informative.** If the replay still
fails inside `demo-up`, iter-16's change means the transcript now says so in its verdict, and iter-15's fix
is shown not to reach the real path — a finding, not a wasted cycle. If the replay succeeds but the probe
fails, the content is registered and unserved, which is the exact defect clause 1 was blind to.

## Expected lift

Clause 1 re-proven, on evidence that **supersedes** rather than reuses `evidence/av-cycle{1,2,3}.json`.
The metric is the clause, not a count: 2 of 5 clauses → 2 of 5, but with clause 1 resting on an instrument
that can see the failure it previously missed. No movement expected or attempted on clauses 2/3/5.

## Phase plan

Declared multi-step shape (the scope-creep tripwire counts a **third unplanned** line, not a third step):

1. **Positive control before spending three cycles.** Run the probe against the current (hand-repaired,
   therefore known-serving) `demo-1` and confirm it FIRES and PASSES. A probe that no-ops would produce
   three green cycles that mean exactly as little as the last three.
2. **Three consecutive cold cycles** — `demo-down 1 --purge` (verified to 0 containers) → `demo-up 1`,
   reading each cycle's own `autoverify.json` and its bring-up transcript's set-dress verdict. Timestamps
   must be distinct and monotonic (iter-14's lesson, and M236's green-gate defect: a stale verdict left on
   disk reads exactly like a fresh one).
3. Supersede the checked-in evidence with the new cycles + a note recording *why* the old three were
   withdrawn.

## Escalation conditions

- A platform commit landing mid-iter that invalidates the attempt → **`re-scope-trigger`, occurrence 2 of
  2 — STOP and escalate.** This is the live one.
- A cold cycle failing for a cause that needs a platform-source change → user-blocker (zero platform edits
  is binding; `demopatch` first).

## Acceptable close-no-lift outcomes

A cycle that goes RED **on the newly-honest verdict** is a complete iter under this protocol: it would show
that iter-15's replay fix does not survive the real bring-up path, which is precisely the question
iter-15's own hand-off flagged as unsettled. Recording that with the transcript is the deliverable, and it
is worth more than a green obtained by not looking.
