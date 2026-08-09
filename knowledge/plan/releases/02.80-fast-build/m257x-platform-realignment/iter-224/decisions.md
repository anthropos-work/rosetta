# iter-224 — decisions

## D-M257x-224-1 — advancing a stale clone is SUBSTRATE REPAIR, and it is not itself a finding

Four `stack-demo` clones were four days behind their origins. They were advanced by `merge --ff-only`,
which is what iter-222 did for `platform` (confirmed still `behind=0` here).

**The advance is not the deliverable and is not reported as a defect count.** The defect is that nothing
asserts the freshness of a clone the fence family grades against, and that gap is recorded as
`ROUTE-M257x-224-drift-guard-blind-to-stale-clone`. The advance merely makes the graded substrate equal
the cited substrate so that today's verdicts mean what they say.

Forbidden ops were not used: no `reset`, no `checkout --`, no `clean`, no `stash`. `cms`'s single dirty
entry (untracked `studio/`) was left untouched, and a `--ff-only` merge cannot rewrite it.

## D-M257x-224-2 — the instrument's own false positives are REPORTED, not silently filtered

The census produced three apparent broken citations. All three are artifacts of the matcher:

- `jobsimulation/ai/ai.go:267` and `:129` — tails of `app/internal/jobsimulation/ai/ai.go`, which the
  prose names explicitly one line above the citation.
- `storage/storage.go` — a row in an indented code-map tree under `internal/`.

They are stated in the close section **with their cause**, rather than removed from the count and the
count published as 0. The pre-registered denominator (150 / 123) was likewise wrong — 58 naive matches
were `app/internal/<service>/…` — and is corrected in place to 92 / 78 rather than quietly restated.

**Rule this instantiates:** after a merge programme, every absorbed service's repo name is also a *path
segment inside the monolith that absorbed it*. Any repo-scoped matcher over this corpus must exclude a
preceding path character or it measures `app` while claiming to measure the archived repo.

## D-M257x-224-3 — the instrument gap is ROUTED, not landed, and the reason is the redirect

`clone_drift_guard` needs an origin-freshness arm (or an UNMEASURED verdict when a clone is behind its
own origin). That is real, and it is **not** landed in this iter:

1. The user's 2026-08-09 redirect puts the corpus's claims and a working stack above the instruments that
   grade them.
2. It is a third distinct line of investigation in an iter whose planned scope was census + repair — the
   scope-creep tripwire's own example.

Routed as `ROUTE-M257x-224-drift-guard-blind-to-stale-clone` with the mechanism written down (D1 is
satisfied *by* staleness), so the next handler does not have to re-derive it.
