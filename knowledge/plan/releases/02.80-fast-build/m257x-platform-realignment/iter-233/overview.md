---
iter: 233
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
---

# iter-233 — is the clone set, which every guard measures against, actually healthy?

## Step 0 — Re-survey

Three iters in a row have now had their first-pass findings dominated by their own substrate:

| iter | first-pass findings | instrument's own |
|---|---:|---:|
| 230 | 14 unresolvable shas | **8** (missing clones) |
| 232 | 11 missing paths | **11** (incl. a clone with no `origin/main`) |

And the pattern predates them: iter-222 found two guards green **only because nobody had fetched**;
iter-228 found the active clones 28 / 12 / 9 commits behind.

**Every one of those was discovered by accident, while looking for something else.** The clone set is the
substrate the whole milestone measures against — and its health has never itself been measured. That is
squarely the redirect's **working-stack** half: a stack whose clones are broken cannot build, and a guard
reading a broken clone reports about nothing.

**Active strategy reference:** `TOK-08` — census the mechanical class instead of stumbling into it.

## Hypothesis

Clone health is fully mechanical — a remote either exists or not, `origin/main` either resolves or not, a
worktree is either clean or not — and the population is small. iter-232 found one broken clone by accident;
a census will find whatever else is there, or prove there is nothing.

## Predictions — SEALED BEFORE MEASUREMENT

| id | prediction |
|----|-----------|
| `P-233-1` | ≥ 15 git clones exist across the `stack-*/` workspaces |
| `P-233-2` | ≥ 2 clones have **no resolvable `origin/main`** (iter-232 found one; predict at least one more) |
| `P-233-3` | ≥ 1 clone has **no `origin` remote configured** at all |
| `P-233-4` | ≥ 1 clone is **dirty** (uncommitted changes), which would make any content read off it unreproducible |

## Expected lift

No `N`/`P` reading. Deliverable: a per-clone health table — remote, `origin/main` resolvability,
behind-count, dirtiness — with every unhealthy clone named, and a statement of which guards read it.

## Escalation conditions

- **No clone is fetched, reset, cleaned, or repaired.** `D-M257x-230-2` freezes the clone set behind gate
  clause 1, and `ROUTE-M257x-222-pin-advance-needs-a-reproof` holds it deliberately. This iter **measures
  and reports**; changing the substrate mid-milestone is the thing being avoided.
- A stale `origin/main` is **not** ill health — iter-222 established that a remote-tracking ref is a cache.
  Report the behind-count as of the last fetch and say so; do not fetch to "improve" it.
- A standing fence is a second deliverable → tripwire → route forward.

## Acceptable close-no-lift outcomes

If every clone is healthy apart from the one iter-232 already named, that is a clean result and the
deliverable — provided the instrument is proved against a deliberately-broken control.
