# iter-267 — decisions

## Pre-registrations — SEALED BEFORE THE SOURCE READ

Sealed in this iter's first commit, corpus at `969a696`. Known at seal time: `internal/workforce/succession.go`
**exists** in `stack-dev/app` at `3eaadae68`. Nothing in it has been read.

**PR-1 — the projection selects on a VERIFIED-SKILL / assessment artifact, not on membership alone.**
If so, an empty projection over a populated org points at the seed, not at the feature. *Risk:* it may
select on a plain membership/role predicate, which would make emptiness a much stronger regression signal.

**PR-2 — the file was modified inside the window the frozen pin does not cover.**
`git log` on `internal/workforce/succession.go` shows at least one commit newer than the rext pin's era
(`D-M257x-258-1`, ~157 iters stale). *Risk:* if the file is months untouched, the product side is
exonerated by dating alone and the cause is seed drift by elimination.

**PR-3 — the three passing siblings do not vouch for the succession inputs.**
`workforce-roster`, `workforce-funnel` and `workforce-org-feedback` resolve through code paths that do not
share succession's selection predicate. *Risk:* if they share a manager/aggregate, their passing narrows
the fault much further and this iter's framing is too coarse.

**PR-4 — the discrimination is decidable from SOURCE + the seeder alone.**
No pin bump, no rext tag, no stack mutation, no re-run of the suite. The iter ends naming one cause.
*Risk:* real. The predicate may depend on runtime state (a materialised view, a cache, a computed window)
that no static read settles — in which case this iter closes `no-lift` and says exactly what must be run.

**PR-5 — the cause is SEED-CONTRACT DRIFT, not a product regression.**
The headline call, and the one most at risk. *Risk:* a genuine product regression would be a materially
larger finding, and the pre-registration exists so that outcome cannot be quietly re-described as drift.

## Escalation clause (pre-registered)

If the cause is a **product regression**, this iter does **not** repair it — v2.8 forbids platform-repo
edits. It reports the finding with its citation and routes it. If the cause is **seed drift**, the repair
lives in `rosetta-extensions/playthroughs` + `stack-seeding` and **requires a tag + pin bump**, which this
iter is forbidden to spend; it is routed with the cause named, which is the deliverable iter-261 asked for.
