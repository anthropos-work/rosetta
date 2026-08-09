---
iter: 228
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-09
---

# iter-228 — is the corpus current with every repo it cites, now that every clone is fresh?

**Type:** tik, under `TOK-08`, on the **corpus half** of the user redirect of 2026-08-09.

## Step 0 — re-survey before targeting (mandatory)

iter-224 established that `clone_drift_guard`'s D1 arm — *"at least one cited sha IS the clone's current
HEAD"* — is **satisfied by staleness**: a clone parked on an old cited sha passes exactly as well as one
at the tip. That was measured against a stale substrate.

**The substrate is no longer stale.** All 13 `stack-demo` clones sit at their origin tips (four advanced
by iter-224; `platform` by iter-222). So D1's verdict is finally *about* something — and the honest
version of its question can be asked for the first time in this milestone:

> **For every repo the corpus reasons about, does the corpus know its current tip at all?**

Re-surveyed: still untouched, and now cheap, because the fetch work is already paid for.

## Cluster / target identified

The **13 clones**, each at its origin tip, versus the corpus's sha vocabulary.

A related, tiny sub-class surfaced by the same survey: the corpus states **`origin/main` is now `<sha>`**
in **3 places, all in `CLAUDE.md`, all about `app`** — a *moving-label* construct that the very same
paragraph warns against (*"Cite the sha, never the moving label"*). Three sites is small enough to settle
outright rather than route.

## Hypothesis

Some clones' tips are cited nowhere in the corpus — the corpus's newest knowledge of those repos predates
their current state — and `clone_drift_guard` stays green anyway because a repo it never cites by sha is
outside its denominator.

## Expected lift

A **13-row census** — per repo, is its current tip cited in the corpus, yes or no — plus a settled answer
on the three `app` moving-label sites. Every number derived here.

## Phase plan

1. **Seal predictions** (this commit — `probe(M257x/228)`).
2. For each clone: current HEAD, and its occurrence count in `corpus/` + `CLAUDE.md`.
3. Check `app`'s actual tip against the three `origin/main is now` claims.
4. Repair what is stale; state the guard's real reach.

## Escalation conditions

- If a repo's tip is uncited simply because the corpus never cites that repo by sha **at all** (e.g. a
  frontend the corpus describes structurally), that is **not** a defect — record it as out-of-scope for
  the drift question rather than manufacturing a citation.

## Acceptable close-no-lift outcomes

**All 13 tips being cited is a first-class result** — it would mean the corpus is genuinely current with
every repo it reasons about, which no iter of this milestone has been able to say.

## Pre-registered predictions — SEALED IN THIS COMMIT

| id | prediction | rationale |
|---|---|---|
| **P-228-1** | `ad9f3c49` **is still** `app`'s `origin/main`, so `CLAUDE.md`'s three moving-label claims are currently TRUE | iter-222 fetched `app` at 19:41 today |
| **P-228-2** | **≥ 2 of the 13** clones have a current HEAD cited **0 times** in `corpus/` + `CLAUDE.md` | the corpus cites some repos structurally, never by sha |
| **P-228-3** | `clone_drift_guard` returns **OK / exit 0** regardless, because a repo it never cites by sha falls outside its denominator — its reach is *"no CITED repo advanced past everything the corpus knows"*, not *"the corpus is current"* | its own REACH line says the second thing is not asserted |

**If P-228-1 is refuted, `CLAUDE.md`'s headline `app` paragraph is stale about the ref it uses to date
every other anchor in it** — the highest-blast-radius single fact checked this run.
