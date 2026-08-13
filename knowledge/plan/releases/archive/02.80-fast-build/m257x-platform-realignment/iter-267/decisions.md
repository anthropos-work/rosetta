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

## D-M257x-267-1 — both candidate causes refuted; the disjunction was a framing, not an enumeration

`GetSuccession` (`app/internal/workforce/succession.go:215` @ `3eaadae68`) fans out six queries. All six
were run verbatim against `demo-2` for the Playthrough's own org — **Meridian Labs = pt-world Org A**,
confirmed from the pinned manifest — and every one returns rows: members 28 · role-requirements 280 ·
declared skills 266 · verified skills 33 · session activity 89 · interview signals 12. No query errors; no
decommissioned schema exists on the stack.

**Candidate 2 (seed-contract drift) is refuted, including on its most plausible axis.** The only commit to
touch `succession.go` in the fold window is `65010b59a` (2026-07-23), which repointed three queries off
`jobsimulation.*` onto `public.job_simulation_sessions` — and the **pinned** seeder already writes the new
location, naming the rename in a comment (`stack-seeding/cmd/stackseed/main.go:47-51`). Reader and writer
agree. **Candidate 1 (product change in selection) is refuted at the layer it was posed** — those numbers
*are* the selection predicates.

**So the fault is above the data layer**: the scoring arithmetic, the response caps
(`successionRolesMax = 25` / `successionAtRiskMax = 40`), the API/route, or the frontend. One hang
mechanism was checked and excluded: `g.SetLimit(analyticsQueryConcurrency)` with a zero limit would block
`errgroup.Go` forever — which is what a 15 s predicate timeout over chrome-that-never-fills looks like —
but `manager.go:53` sets it to **6**, equal to the goroutine count.

**The generalisable half.** iter-261 was rigorous in refusing to guess between its two candidates, and
still wrong, because the *list* was never justified. A disjunction is only an enumeration if something
derived it. **Write the residual term** — *"or something else"* — or the next iter will spend itself
choosing between two refuted options. This iter cost nothing to discover that, because the
pre-registrations made "then it must be the other one" unavailable.
