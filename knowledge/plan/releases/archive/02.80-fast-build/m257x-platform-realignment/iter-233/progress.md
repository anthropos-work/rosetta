**Type:** tik — under `TOK-08`.

# iter-233 — the clone set every guard measures against

## The census

15 git clones across `stack-demo/` and `stack-dev/`. Per clone: `origin` remote, HEAD, `origin/main`,
behind-count **as of the last fetch** (iter-222: a remote-tracking ref is a cache — not fetched here), and
working-tree dirtiness.

| clone | remote | HEAD | origin/main | behind | dirty |
|---|---|---|---|---:|---:|
| `stack-demo/ant-academy` | yes | `22df69dd` | `c885dab2` | 9 | 0 |
| `stack-demo/app` | yes | `ad9f3c498` | `3eaadae68` | **28** | 1 |
| `stack-demo/cms` | yes | `f38c0c4` | `f38c0c4` | 0 | 1 |
| `stack-demo/graphql-wundergraph` | yes | `60c229f` | `60c229f` | 0 | 0 |
| `stack-demo/jobsimulation` | yes | `82cb66ec` | `82cb66ec` | 0 | 0 |
| `stack-demo/messenger` | yes | `e9421c6` | `e9421c6` | 0 | 0 |
| `stack-demo/next-web-app` | yes | `8297c684c` | `19423a1fb` | **12** | 0 |
| `stack-demo/platform` | yes | `0c91421` | `0c91421` | 0 | 0 |
| `stack-demo/roadrunner` | yes | `87d8d44` | `87d8d44` | 0 | 0 |
| `stack-demo/rosetta-extensions` | yes | `09d0607` | `3667d5b` | **159** | 2 |
| `stack-demo/sentinel` | yes | `f2c4619` | `f2c4619` | 0 | 0 |
| `stack-demo/storage` | yes | `9f8cb53` | `9f8cb53` | 0 | 0 |
| `stack-demo/studio-desk` | yes | `41ee3575` | `41ee3575` | 0 | 0 |
| `stack-dev/studio-desk` | yes | `795a411d` | **UNRESOLVED** | — | 0 |
| `stack-dev/studio-room` | yes | `aeec036` | `aeec036` | 0 | 0 |

## 4 flagged, 0 actually broken — and every explanation is worth having

**`stack-dev/studio-desk` — cloned from a BUNDLE, not from origin.** It *has* an `origin` URL
(`git@github.com:anthropos-work/studio-desk.git`) and 11 remote-tracking branches — all under
**`bundle/`**, none under `origin/`. `bundle/main` exists; `origin/main` does not. HEAD is a normal
`refs/heads/main`.

> Any guard that does `rev-parse origin/main` reads **nothing** here — and gets no error it can
> distinguish from "the ref is stale". This is precisely what bit iter-232, and it is a **third** shape of
> the `§8` rule that a remote-tracking ref is a cache: it can be stale (iter-222), it can be *absent*, and
> it can live under a **namespace nobody expected**. A configured `origin` URL is not evidence that
> anything was ever fetched from origin.

**`stack-demo/app` and `stack-demo/cms` — `?? studio/`, and it is the CI embed.** `app/.gitignore:79`
carries `studio/*`, and `stack-demo/app/studio/gen.py` and `studio/services/ai.py` **exist on disk**.

> This independently confirms iter-232's reading from the other direction. iter-232 concluded that the
> corpus's `app/studio/gen.py` citations are **image paths** — the `anthropos-studio-room` repo that CI
> pulls into the `app` image — by finding those files in the `studio-room` clone. Here they turn up *inside
> the `app` clone itself*, ignored by `app`'s own `.gitignore`. Two independent routes, same answer: the
> citation is correct and the census's first-pass verdict was wrong.

**`stack-demo/rosetta-extensions` — 159 behind, dirty(2), and both are expected.** The dirt is
`playthroughs/e2e/report/` and `playthroughs/report/last-binding-report.json` — **test exhaust**, the class
`§9` names outright (*"the working tree contains EXHAUST"*). The 159 is a **pinned-tag consumption copy**
doing exactly what `CLAUDE.md` says per-stack copies do; the authoring copy is `.agentspace/`.

> Worth stating because the raw table looks alarming and is not: **"behind" is a defect for a clone that is
> supposed to track, and a contract for one that is supposed to be pinned.** The number alone does not say
> which — the clone's ROLE does, and no guard here reads a role.

### Predictions, graded — 2 HELD, 2 REFUTED

| id | prediction | result |
|----|-----------|--------|
| `P-233-1` | ≥ 15 clones across `stack-*/` | **HELD — exactly 15** |
| `P-233-2` | ≥ 2 clones with no resolvable `origin/main` | **REFUTED — 1.** iter-232's find was the only one |
| `P-233-3` | ≥ 1 clone with no `origin` remote | **REFUTED — 0 of 15.** And the near-miss is instructive: the one broken-looking clone *has* a remote and still cannot answer `origin/main` |
| `P-233-4` | ≥ 1 dirty clone | **HELD — 3.** All three are exhaust or a documented CI embed; none is an edit |

## The instrument, proved

A throwaway repo built in the scratchpad with no remote, no `origin/main` and an uncommitted file: the
census reports **all three** conditions. Without it, "0 clones with no remote" is indistinguishable from a
probe that never checked.

## Close — 2026-08-10

**Outcome:** the substrate three consecutive iters kept tripping over is now enumerated. **15 clones, 4
flagged, 0 broken** — one cloned from a **bundle** so its refs live under `bundle/` and `origin/main`
resolves to nothing, two carrying the git-ignored CI `studio/` embed, one a pinned consumption copy with
test exhaust. `app` 28 / `next-web-app` 12 / `ant-academy` 9 behind, unchanged and deliberately not
advanced.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: y — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: exit-5
**Decisions:** `D-M257x-233-1` (nothing was fetched, reset or cleaned — the set stays frozen),
`D-M257x-233-2` (a behind-count is graded against the clone's ROLE, which nothing records).
**No `N`/`P` movement is claimed** — this iter took no graded reading.

**Suite state at close** — no pytest section run; this iter changed no rext code and no corpus prose.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-233-a-clone-can-have-a-remote-and-no-origin-refs` → **new, and the sharpest.** A configured
  `origin` URL is not evidence anything was fetched from origin; `stack-dev/studio-desk`'s refs are all
  under `bundle/`. Any guard resolving `origin/<branch>` should report **UNMEASURED** naming the namespaces
  that *do* exist, rather than a bare failure. Third shape of `§8`'s cache rule.
- `ROUTE-M257x-233-behind-count-has-no-ROLE` → **new.** 159 behind is a contract for a pinned consumption
  copy and a defect for a tracking clone, and nothing on disk says which a clone is. `clone_pin_guard`
  knows about pins; the health view does not.
- `ROUTE-M257x-232-stack-dev-studio-desk-is-a-broken-clone` → **re-characterised, not closed.** It is not
  broken; it is bundle-sourced. Superseded by the two routes above.
- All prior routes → open, unchanged.

**Lessons:**
1. **A remote-tracking ref has three failure modes, not one:** stale (iter-222), absent, and **living under
   an unexpected namespace** (here). Only the first was written down.
2. **"Behind" is meaningless without a ROLE.** The same number is a contract or a defect depending on
   whether the clone is meant to track or to be pinned, and nothing records which.
3. **Two independent routes to one answer is worth more than either.** iter-232 concluded `app/studio/*`
   are image paths by finding them in `studio-room`; this iter found them inside the `app` clone, ignored
   by `app`'s own `.gitignore:79`. The second reading is what turns a plausible explanation into a fact.
4. **Census the substrate BEFORE it bites you.** Three iters found substrate defects by accident while
   looking for something else; this one cost 15 minutes and found the population.
