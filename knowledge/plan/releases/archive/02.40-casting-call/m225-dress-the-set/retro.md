# M225 "dress the set" — retro

## Summary
The demo-integration section milestone: the hiring org comes up **auto-set-dressed on a default `/demo-up`** and is
proven at three layers — an S1 bring-up **guard** (the `is_hiring`-gated autoverify cheap-win), an S2 hiring
**coverage gate** (all 3 seats MET via `manifestFor` 3-arg org/identity dispatch + a `profileGated` persona mode for
`apps/hiring`), and an S3 **GREEN recruiter playthrough** (`pt-hiring-recruiter-compare` on pt-world Org D "Kestrel
Hiring Group") — plus the S4 corpus docs. Reused the M42 coverage + M202 playthrough machinery unchanged (never
forked). **Zero platform-repo edits** — code in the rext authoring clone (`casting-call-m225-harden`, `be431c3`),
docs in rosetta. Shipped green; the harden pass added 10 regression fences and found **0 bugs**.

## Incidents This Cycle
- **P2 (plan-vs-reality, caught by the S1 investigation — D1): most of S1 was already delivered.** Tracing the
  default `/demo-up` chain showed the hiring org was ALREADY end-to-end default-on (delivered incidentally by
  M223+M224): the stories preset seeds it, `readHiringSimPool` runs unconditionally, the HIRING sims ride along in
  the standard directus content-surface capture, and the two-app UI container is default-on. So there was **no
  hiring-specific manual step to remove**. S1's genuine, non-redundant deliverable became a bring-up-tail **guard**
  (the autoverify cheap-win) that brings the cold-cache silent-empty catch FORWARD to the `/demo-up` tail, plus the
  docs. *Lesson: front-load the trace before scaffolding — "fold in the replay" was a no-op; the real gap was that
  "it comes up real" was assumed, not checked.*
- **P2 (stale plan premise, caught by the KB-fidelity gate — KB-1): the `job_position` replay premise was refuted
  before M225 opened.** The scaffold said S1 folds in a `directus.job_position` replay; M222 BA-6 had already
  measured 0 rows captured + the scoreboard doesn't read it, and M223 D4 DROPPED it. The corpus was already correct;
  only the M225 plan docs carried the stale framing — reconciled inline. *Lesson: the KB-fidelity gate earns its
  keep on inherited scaffolds; the plan can rot even when the corpus is right.*
- **0 code regressions.** Harden pass 1 added 10 deterministic regression fences (5 coverage-manifest.unit + 4
  test_verify.py + 1 hiring_isolation_test.go), **0 bugs surfaced**. Deterministic suites re-verified GREEN at merge
  base; flake 5/5.

## What Went Well
- **Reuse over fork, twice.** S2 extended `manifestFor` with the SAME 3-arg org/identity dispatch as the
  AI-readiness showcase-org precedent (a `profileGated` persona *mode*, not a fork, adapts the role-skills/avatar
  checks to `apps/hiring`'s `/home` self-view); S3 added the FOURTH Playthroughs product reusing
  hero-login/PageObject/ptvalidate/ptreport unchanged. The hiring vantage cost O(surfaces), not O(tests).
- **Three-layer silent-empty defense.** The cold-cache / starved-HIRING-pool failure mode (the demo's real risk) is
  now fenced at the S1 bring-up guard (warns at the tail), the S2 coverage sweep (empty grid = FAIL), and the S3
  playthrough (empty scoreboard = FAIL) — an assumed property turned into a checked one at every layer.
- **Test data ≠ demo data, enforced.** pt-world Org D "Kestrel Hiring Group" is deliberately distinct from the
  demo's "Meridian Talent" AND this world's Org A "Meridian Labs"; `hiring_isolation_test.go` pins both the shape
  and the distinctness so the two worlds stay cleanly separable.

## What Didn't
- **The S1 scaffold overstated the plumbing.** It described set-dress work that M223/M224 had already delivered — the
  overview/progress/spec docs carried a "fold in the replay" framing and a `job_position` premise that were both
  stale before the milestone opened. Both were reconciled (D1 + KB-1), but a tighter pre-milestone read of the
  already-shipped default bring-up would have scoped S1 as "guard + docs" from the start.

## Carried Forward
- **Standing test-debt backlog (re-fated fresh 2026-07-17, routed to the v2.4 release close):** the 8 pre-existing
  demo-stack failures (`test_cockpit.py` ×6 + `test_purge` + `test_reap`) → a future demo-stack test-debt harden
  pass. Not M225's work (untouched, unrelated files).
- **Declared TODO (routed to the v2.4 release close):** the M204 `assign-and-track.UC1` assign-WRITE half — a
  declared in-manifest `unimplemented` build-reference gap; surfaced only in M225's `playthroughs.md` 14→15-live
  count.
- **M226 (opening night):** the live `billion`/tailnet proof of the 7-condition hiring gate, incl. the recruiter p95
  click→ACCESS < 5 s (the 3rd measured vantage), default `/demo-up`, no flags.

## Metrics Delta
- Go test funcs: 1885 → **1887** (+2). `test_verify.py` **124** (incl shellcheck); TS unit **61** (stack-verify/e2e)
  + **69** (playthroughs/e2e); both `tsc` clean. Flake (milestone-owned) **0** (5/5 clean). Platform edits **0**.
  Supply chain GREEN (0 net-new deps). Deferral audit **YELLOW** (0 new; 2 inherited carries re-fated). See
  `metrics.json`.
