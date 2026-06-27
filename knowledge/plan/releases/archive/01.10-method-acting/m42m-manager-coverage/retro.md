# M42m Retro — Manager 100% demo coverage

## Summary
M42m closed the LAST milestone of v1.10 (and the SECOND `iterative` milestone) **on-gate**: logged in as Dan
Rossi (`dan-manager`, the org-intelligence seat @ Cervato Systems) on a **fresh zero-manual demo-up**, the
(reused M42e) Playwright semantic-coverage sweep reports `gateMet:true` — 70 reachable pages, **0 failing
sections, 0 persona failures, 0 prod-eject escapes, frontier EXHAUSTED** — AND the M42e employee gate HELD on
the same fresh stack (no regression: 59 reachable, `(0,0,0,0)`, EXHAUSTED). The milestone REUSED the M42e
harness + `coverage-protocol.md` (no new harness) and drove TOK-01's four leverage-ordered fix lines into rext
across 5 iters: the **`demopatch` tool** (the sanctioned mechanism for the platform-bound Studio left-nav
escape), the **FeedbackSeeder org-feedback JOIN-mirror**, and the **manager harness namespace** (the
`/enterprise/*` route reconcile + `MANAGER_SAMPLE_RULES` superset). **Zero CANONICAL platform-repo edits.**
Code-of-record @ tag `method-acting-m42m-harden-final`; the rosetta side carries the docs-only doc-half + the
5-iter plan archive. Closed in a near-clean review: 3 findings, 0 blocking. **With M42m closed, v1.10 is
reproducibly gate-complete across both per-vantage coverage gates.**

## Incidents This Cycle
- **iter-02 RE-SCOPE TRIGGER (the platform-bound Studio escape) — P1, the milestone's defining event.** The
  baked `studio.anthropos.work` left-nav link rendered on every authenticated manager page (139× escapes) —
  the exact prod-eject the user flagged in the live-demo review. The env-rewrite hypothesis was **falsified**:
  next-web's `STUDIO_URL` (`core-js/constants/urls.ts:12`) is a `NEXT_PUBLIC_NODE_ENV` ternary with **no**
  per-URL `NEXT_PUBLIC_STUDIO_URL` override (unlike `ACADEMY_URL`), and the only knob (a global dev-flip)
  points Studio at the wrong port + side-effects other URLs. The sole clean fix was a forbidden platform edit
  → recorded as RESCOPE-1, escalated to the user (per the protocol's re-scope trigger). **The user's pivot
  (iter-03):** a **demo-patch tool** — source-patch the demo's EPHEMERAL gitignored clone before the build,
  trap-revert after, CANONICAL repos never touched (6 guards). Resolved demo-only **139→0** while keeping the
  v1.10 zero-platform-edit line intact. The lesson — *a baked URL with no per-URL `NEXT_PUBLIC_<thing>_URL`
  override is platform-bound; the env-rewrite row does NOT apply; route it to the demopatch tool* — is folded
  into `coverage-protocol.md` (the "Platform-bound escape" routing row) + `frontend-tier.md`.
- **The route-model error masquerading as a content gap — P2, resolved iter-04, not a defect.** The M42e
  smoke-sweep's `notReached=5` looked like 5 empty `/workforce/*` dashboard pages needing seed/serve-grant
  work. Diagnosis (live, dan-manager) showed the real cause: the manifest GUESSED `/workforce/*` sub-routes,
  but the real surface is the single tabbed SPA `/enterprise/workforce` (the M36 sections are in-page tabs, the
  `?tab=` query is ignored client-side) + 5 sibling `/enterprise/*` pages. The dashboard already rendered rich
  real data (493 mapped / 262 verified / 53.1% coverage, 19 cards / 67 charts) — the 6 M36 seeders populate it.
  Reconciling the manifest turned `notReached=5` into 6 fully-asserted PASSing pages. Lesson: diagnose the real
  route model before assuming a content gap.
- **`/enterprise/organization-feedback` "No data" on a fully-seeded org — P2, an inserted-but-invisible JOIN
  gap, fixed iter-04.** `GetOrganizationFeedback` (read-only platform diagnosis) JOINs feedback to the app
  mirror `public.local_jobsimulation_sessions` and scopes by the MIRROR's org; the population sessions the
  feedback linked had no mirror (only the PersonaSeeder mirrors heroes), so the JOIN was empty. Fix: the
  FeedbackSeeder now also writes a mirror per feedback session (reconstructing coherent values from the same
  deterministic key) — 0→103 joinable, "No data" → "103 sessions / 70% pos / 59% pass" + 21 rows. The
  org-feedback analog of the G14 inserted-but-invisible class. Zero platform edits.
- **iter-05 build-time `No space left on device` — P3, environmental, not a tooling bug.** The first fresh
  `demo-up` hit a Docker-VM disk ceiling during the next-web image build; `docker system prune` + the same
  fresh `demo-up` resumed idempotently (the M17 re-run guards held). One real clone-cleanliness gap WAS fixed
  in rext (R1b: `ensure-clones.sh` now sweeps a crash-left tooling `.dockerignore`, byte-identical + untracked
  guarded).
- 0 flakes (3× clean at the final harden, `-count=1`); 0 regressions (supply-chain GREEN — zero
  go.mod/go.sum/package.json/lockfile change in the whole M42m footprint; all 5 alignment gates 100%,
  unchanged — zero clerkenstein change). The close's review surfaced 0 behavioural defects (corpus diff is
  docs-only; the demopatch REFUSE grid + the feedback mirror held against the full adversarial/edge grid in the
  final harden — 0 production bug surfaced).

## What Went Well
- **The harness REUSE thesis held.** M42m authored no new harness — it pointed the M42e Playwright harness +
  `coverage-protocol.md` at the manager hero and the persona machinery generalized for free (dan-manager PASSES
  all 3 persona asserts with no persona work). The "one harness, two vantages" design paid off exactly as M42e
  intended.
- **The re-scope trigger worked as designed.** A genuinely platform-bound blocker was *escalated* (not absorbed
  via a quiet platform edit), and the user's demo-patch pivot resolved it inside the zero-edit line. The
  demopatch tool is now a reusable, guard-fenced mechanism for any future platform-gated demo bug.
- **The fresh zero-manual demo-up bar caught the real reproducibility surface.** iter-04 proved the gate on a
  re-seeded live stack; iter-05 demanded the *build* reproduce it — surfacing the R1b clone-cleanliness gap a
  live-patched stack would have masked. Same lesson M42e learned, re-confirmed.
- **Zero CANONICAL platform edits held** across a milestone that touched the manager nav, the Workforce
  dashboard, org feedback, and a baked frontend URL — every fix landed in rext (seeders / harness / the
  demopatch on the ephemeral clone).

## What Didn't
- **The manager manifest shipped with wrong `/workforce/*` route guesses** (a `calibrated:false` best-guess
  from M42e's forward-route), which read as a content gap until iter-04's live diagnosis. A route-model probe
  before authoring the manifest descriptors would have saved the detour. Captured in `coverage-protocol.md`
  (the manager manifest is now `calibrated:true` against the real `/enterprise/*` surface).
- **The demopatch's first fresh-build surfaced a real G6 demo-detection bug** (the consumption-clone registry
  is empty at frontend-build patch-time — the demo-N row is written later), forcing a dual-signal fix
  (structural identity OR registry type). Caught in iter-03, not in the authoring copy — the build-time
  patch-window state wasn't modeled in the unit tests until then (now pinned).

## Carried Forward
- **None into a future milestone of v1.10** — M42m is the last milestone; all carry-forward resolved (DEF-M40-01
  manager-half LANDED Fate-1; RESCOPE-1 resolved demo-only). The standing release-orthogonal backlog
  (**DEF-M10-01** cloud SnapshotStore / S3 blob bytes, **DEF-M21-01** replayCmd hermetic test, **M25-D9** dev
  taxonomy rc=4) carries unchanged to the cross-release ledger — none in v1.10 scope, none aged in.

## Metrics Delta (from metrics.json)
- **rext Go test funcs:** 1373 → **1376** (+3: stack-seeding 534→537, the FeedbackSeeder org-feedback mirror
  tests — MirrorCopyErrorPropagates / EmptyPopulationNoOp / MirrorMatchesFeedbackOneToOne). clerkenstein /
  stack-snapshot / alignment / stack-secrets unchanged.
- **Python demopatch suite:** `test_demopatch.py` 18 → **43** (+25 adversarial-guard + manifest-loader parser
  tests — the highest-value harden target, a tool that patches platform source).
- **TS Playwright harness:** the manager namespace gained a pure-logic unit spec
  (`coverage-manifest.unit.spec.ts`, +17) — the manifest decision logic now pinned in CI with no stack.
- **Gate:** manager `gateMet:true` (70 / 0,0,0,0 / EXHAUSTED) + employee no-regression (59 / 0,0,0,0 /
  EXHAUSTED). **Flake:** 0. **Supply-chain:** GREEN (0 dep/lockfile change). **Alignment gates:** 5/5 at 100%
  (N/A change). **Platform edits:** 0.
- **Code-of-record:** `rosetta-extensions` @ tag `method-acting-m42m-harden-final`.
