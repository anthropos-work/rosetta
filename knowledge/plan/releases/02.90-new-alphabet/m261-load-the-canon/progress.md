# M261 — Progress

**Status: BLOCKED at start — not on our tooling, on an external dependency.** 2026-08-14.

- [ ] re-capture from a safe prod source — **BLOCKED: the canon is not in production**
- [ ] PURGE the snapshot cache — **must NOT run yet** (see below)
- [ ] replay proven on a cold stack — blocked on the above
- [ ] cold-start runbook updated — possible now
- [ ] batch-prompt cache invalidation observed — possible now
- [ ] the five net-new tables — **partially blocked** (declarable, not capturable)

## The finding

**The taxonomy v2 canon has never been loaded into production.** From the platform's own
`app/knowledge/taxonomy-canon-migration.md` at `4bccda085` (2026-08-14 — today's HEAD, and the doc's
own last-touched commit):

> **Status:** canon complete, invariants green, **nothing in production yet** — the last migration
> applied there is `20260804160000`.

Its step ledger says the same twice over: **step 5 (canon load) — "rehearsed, never in production"**,
**step 7 (pointer rewrite) — "rehearsed, never in production"**, and its own summary reads *"Two
steps are left, and it is the same step twice: run in production what has already been rehearsed on a
copy."* The doc names the live production catalogue as **43,584 skills · 22,511 roles** (`:74`).

**Independently corroborated on our side:** the newest cached taxonomy snapshot
(`.agentspace/snapshots/taxonomy/`, captured 2026-06-29, `source: primary-read`) holds
**42,790 skills / 22,470 job roles / 1,447 specializations** — the old catalogue.

## What this means, precisely

- **A prod capture today would capture the OLD taxonomy.** Not a broken capture — a correct capture
  of the wrong vocabulary. M261 cannot "load the new canon" from a source that does not have it.
- **The cache must NOT be purged.** D-M259 routed a purge here on the reasoning that node-ids moved
  and a stale artifact is *wrong*. That reasoning holds only after the migration. Purging now would
  destroy the only valid taxonomy snapshot in existence and leave every stack with no catalogue at
  all — the exact "stack boots, catalog empty" failure the corpus warns about.
- **Nothing is broken today**, and that is the honest headline. A demo or dev stack built right now
  gets the old taxonomy and works exactly as it did last week.

## Second, independent blocker (would not matter if the first were absent)

This machine has **no production database access**: no `psql` (`libpq` not installed), no `~/.pgpass`,
and no wired `postgres` MCP tool in this session. Both documented paths in
[`db-access.md`](../../../../corpus/ops/db-access.md) are unavailable. Worth recording because it
would block a prod capture even after the platform migrates.

## Why M259 did not catch this

M259's verdict (GO) stands — the canon is real, legible and measurable, and everything it reported is
accurate. But its `In:` list said *"measure the real new prod counts"* and what it actually measured
was **the canon bundle's** counts. It never read production, and I framed that as a benefit ("no prod
read required, no tenant-data exposure"). It was a benefit — and it is also why the barrier missed
the one question that decides this milestone: *is the canon deployed?* **A barrier that reads the
artifact but never the deployment can certify the artifact and still not clear the path.**
