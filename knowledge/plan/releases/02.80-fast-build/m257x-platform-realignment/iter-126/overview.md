---
iter: 126
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-07
---

# iter-126 — the guard off exit 2, and the re-pin backlog enumerated

**Run 80, tik 3.** The two items routed forward from iter-125's close: the `platform_alignment_guard`
could-not-check, and priority 4 (the ≤13 real re-pin backlog). Both were named as this iter's targets at
iter-125's close, so both are planned scope.

## Active strategy reference

No `TOK-09` (`TOK-08`'s sealed rule). Runs under **the user's directed scope of run 80**. `F4` binds.

## Target 1 — `platform_alignment_guard` exits 2, and it must be RESOLVED, not silenced

**The guard is behaving correctly on a changed subject.** iter-123 documented the 93-repo GitHub org and
in doing so made this corpus cite `infrastructure` and `db-backup` — repos **no stack clones**. The guard
graded those citations `unresolvable`, i.e. as blind spots, and returned exit 2.

The directive names the two admissible resolutions: **widen the subject to the org with the clone-set
limitation disclosed**, or **have the new citations declare themselves unclonable**. Silencing is
forbidden. Re-survey at open shows the 7 split **across both**, so both apply:

| head | ×  | what it really is |
|---|---|---|
| `terraform` | 3 | `cms`'s row cites `terraform/production/services.tf:64/85/88` **unqualified** — ambiguous across 12 terraform-bearing repos |
| *(bare, no preceding path)* | 2 | `graphql-wundergraph`'s row — **iter-124's own bare pins** `:509-517`, `:521` |
| `infrastructure` | 1 | a repo the map documents; in no clone set |
| `db-backup` | 1 | same |

So **5 of the 7 are citation defects** (qualify them — the "declare themselves" half) and **2 are
substrate limits** (widen + disclose). Conflating the two is what produced the exit 2.

## Target 2 — priority 4: the re-pin backlog

iter-123 measured **76 of 89 = 85.4 % RESOLVE at their pin**, leaving **≤ 13** (7 no-sha, 6 candidate rot),
**and did not enumerate them** — which is why the item is still open: a count with no list cannot be
closed. This iter enumerates the no-sha class exactly and closes it or says why each member stays.

**Re-survey caught the trap before the work started.** Reading the 89 backlog sites' blocks at **HEAD**
yields **22** no-sha members; reading the same sites at **`afe58ac`** — the ref iter-123 named — yields
**7**, reproducing iter-123 exactly. The sites are line-pinned at the census's ref and ~30 commits have
landed since. **A 3.1× inflation manufactured entirely by a stale substrate**, and it is this milestone's
own standing rule (`D-M257x-122-4`) arriving one level up: the substrate for a *corpus-side* derivation is
the corpus at a ref, not the corpus now.

## Hypothesis

The no-sha class is closable — most members cite `app`, whose clone carries the refs the corpus already
names elsewhere — and a minority will legitimately stay because no ref applies to them.

## Expected lift

No `N` movement, no reading. Deliverables: the guard off exit 2 with the limitation disclosed **in the
verdict sentence**, backed by a mutation control and an anti-vacuity control that fire; and the 7 enumerated,
each closed or refused with a reason.

## Phase plan

1. Qualify the 5 unqualified/bare citations in the map.
2. Widen `check_citations` with an `unclonable` bucket, **gated on the map documenting the repo**;
   disclose on every run and in the verdict.
3. Tests: mutation + anti-vacuity, each asserting it applied.
4. Enumerate the 7 at `afe58ac`; re-pin or refuse each.
5. Re-run the guards + the whole suite; close.

## Escalation conditions

A widening that cannot be killed by a mutant is a silencer — if no mutant kills it, revert and escalate.

## Acceptable close-no-lift outcomes

If the 7 turn out to be individually unclosable, enumerating them with a reason each is the deliverable.
