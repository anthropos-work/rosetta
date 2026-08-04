---
iteration_type: tik
status: closed-fixed-partial
---

# iter-68 — the routed citation class, re-derived at the ref the gate names

**Active strategy reference:** `TOK-05` (*stop repairing claims; fence the predicates under them*),
step 2 of its ordering — **fence → citations → map state → read**. Steps 1 and 3 landed (iters 60–62,
64, 66, 67); this is the citations step, carrying `FIX-M257x-iter63-app-citation-residual`.

## Step 0 — re-survey before targeting (mandatory, and it changed the target)

Re-derived at open, not inherited. Two things moved.

**1. The class is 64, not 68.** iter-63 measured **104 sites / 86 distinct / 22 files** (18 mainline
+ 68 non-mainline) against app `b948604`. Re-running *its own* enumerator against today's corpus:

| reading | sites | distinct | files | mainline | non-mainline |
|---|---|---|---|---|---|
| iter-63, as recorded | 104 | 86 | 22 | 18 | **68** |
| iter-68, same instrument, same app ref | **123** | **96** | **22** | **32** | **64** |

The class grew by 10 distinct citations and the mainline share nearly doubled, because **iters 63–67
wrote citations of their own** — the two-sided `mid-fold` row alone added six. §5 rule 34 said a
corpus repair moves the corpus's own line numbers; this is its sibling — **a corpus repair also
enlarges the corpus's own citation class.** The routed figure was already stale when it was routed.

**2. The adjudication ref moved under the milestone, again.** `app` origin/main advanced **56 commits
to `9d00a313` v1.367.0** at 2026-08-04 10:56Z — *this morning*, mid-iteration. The pinned clone sits at
`b948604` v1.366.0.

## Step 0b — §7 rule 4b, applied before deciding (`D-M257x-59-3`)

Rule 4b says a pin advance is not vetted until its **citation delta** is measured. Measured
read-only (`git show`, no checkout, no risk to the green stack), same 96-citation universe, graded at
both refs:

| binding | ref | HELD | MOVED | GONE | DEAD | UNNAMED |
|---|---|---|---|---|---|---|
| pooled (iter-63's) | `b948604` | 42 | 25 | 7 | 4 | 18 |
| pooled | `9d00a313` | **17** | 49 | 8 | 4 | 18 |
| clause-bound | `b948604` | 20 | 21 | 8 | 4 | 43 |
| clause-bound | `9d00a313` | **8** | 32 | 9 | 4 | 43 |

**25 of the 42 citations that hold at the pin break at origin HEAD** — 60%, in one working day.

## Cluster / target identified

The gate reads *"against platform @ **origin HEAD**, never a pinned pre-drift commit."* A cold
`make init` clones `app` at its main, so **origin HEAD is what a fresh stack runs**, and repairing the
class against `b948604` would buy 45 claims that are already false. **The adjudication ref for this
iter is app `9d00a313`.** The defect set is the 45 citations graded MOVED/GONE/DEAD there, carried by
~14 corpus lines across 14 files.

Underneath 12 of those 45 is one predicate, and it is the material finding: **release 09.00
"support-in-app" has landed in `app`.** `STORAGE_RPC_ADDR` is retired from `main.go` and all three
operator CLIs; `internal/messenger/{flow,adapters,sender}` are imported into `main.go` and app **takes
over messenger's Redis consumer group**. The corpus's `storage` row — the `mid-fold` state iter-64
built an eighth vocabulary token for — asserts a consumer side that no longer exists.

## Hypothesis

Repairing the class at origin HEAD closes `FIX-M257x-iter63-app-citation-residual` **and** lands the
storage/messenger fold on the map, the two service docs and root `CLAUDE.md` — the largest single
block of clause-5 residual currently known.

## Expected lift

Gate stays 4 of 5 (clause 5 is graded only by a reading returning zero, and no reading is taken this
iter). The measurable outcome is the defect count at the adjudication ref: **45 → 0**, re-derived.

## Phase plan (declared multi-step shape — the tripwire counts UNPLANNED lines against this)

- **A** — re-derive the class + the rule-4b delta. *(done at open, above)*
- **B** — repair the 45 defect citations at app `9d00a313`, in committed slices:
  **B1** the fold predicate (map row + `storage.md` + `messenger.md` + root `CLAUDE.md`);
  **B2** the remaining citation defects.
- **C** — `CHECK-M257x-iter63-quoting-a-retired-token`.
- **D** — gates + close.

## Escalation conditions

- A repair that cannot be adjudicated against a platform artifact → route, never guess (§5 rule 19).
- A third *unplanned* line of investigation → tripwire; land what is complete, route the rest.
- Platform moving again mid-iter → record as a second invalidation occurrence and surface it.

## Acceptable close-no-lift outcomes

Falsifying the fold (finding storage/messenger are *not* folded at origin HEAD) would be a complete
iter even though it repairs nothing — the evidence is already collected, so this is unlikely.

## What this iter does NOT do

It does not take the graded reading (step 4). A reading over a known-unrepaired class measures
nothing. It does not advance the clone's checkout — the adjudication is read-only against `origin/main`,
so the green stack and clause 1 are untouched.
