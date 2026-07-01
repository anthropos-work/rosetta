# M53 — retro

## Summary
M53 is the FINAL v1.10b milestone: the release's single source of acceptance truth. The one live demo the
M47–M52 fixes were iterated against was **destroyed** (`/demo-down 1 --purge` — all 17 containers + network +
ALL demo-1 images, exercising M49 #6) and **cold-rebuilt from scratch** on a `stack-demo`-only box by a single
`/demo-up 1` at the `v1.10.1` pin (no manual steps), then the acceptance bar was asserted: **6/6 criteria +
academy F6 GREEN from cold** (AB1 healthy backends, AB2 prompt-free replay, AB3 set-dress+seed 3 orgs+cockpit,
AB4 both-vantage coverage, AB5 AI-readiness dashboard on Northwind, AB6 complete manifest download, F6 academy
authenticated-member session). The academy F6 seeder/wiring (the sole planned new-code surface, a Fate-3 handoff
from M50→M51→M53) landed in rext `e91f004`. Closed `section`-shaped with all 6 `In:` items delivered as Fate-1;
2 close findings, both Fate-1 doc fixes. → **v1.10b is GREEN from cold; the release is complete.**

## Incidents This Cycle
- **AB4 — an M51-owned gate regression, surfaced from cold, fixed at the acceptance gate (P1-class).** On the
  first from-cold both-vantage sweep, the **M50 canonical manager gate** (`dan-manager` @ Cervato) was RED:
  `failingSections=2`, both on `/enterprise/workforce/ai-readiness` (page rendered HTTP 200, 0 ejects, but showed
  "No AI readiness data yet for this org"). **Root cause:** M51 iter-05 D3 added the ai-readiness page to
  `MANAGER_MANIFEST.seedPaths` **unconditionally**, so every manager sweep primed + asserted the funnel — but the
  199 frozen snapshots seed only for Northwind (the showcase org). M51's gate ran ONLY `dana-manager` @ Northwind
  (where it passes), so it never re-ran the M50 Cervato sweep and never saw the regression its own manifest change
  introduced. **Fix (user-approved gate exception):** org-condition the manager manifest — `manifestFor(vantage,
  expectedOrg)` returns the showcase `MANAGER_MANIFEST` only for Northwind, else a new `MANAGER_MANIFEST_BASE` that
  omits the seedPath + descriptor (rext `117fe41`, +3 unit tests). Both manager vantages re-verified GREEN. This is
  **exactly the late cross-milestone regression M53 exists to catch** — the from-cold both-vantage assertion is the
  first joint re-measure of the M50 + M51 gates; the fix-on-live serialization never re-ran M50's gate after M51's
  manifest change.
- No flakes: close flake gate 5/5 Go seeders (`-shuffle=on`) + 5/5 Python cockpit+academy (101 each) + 5/5 TS
  coverage-manifest (29 each), all clean.

## What Went Well
- **The milestone did its defining job.** M53's stated risk was "a cold-rebuild surfacing a regression late" — and
  it surfaced exactly one (AB4), which the warm fix-on-live chain had structurally hidden. A from-cold, both-vantage,
  all-3-org acceptance is a real gate, not a rubber stamp.
- **The gate exception stayed disciplined.** AB4's fix touched only a rext test/gate artifact (the coverage manifest),
  never platform code; it was narrow, fully test-covered, and consciously approved + recorded (`decisions.md` AB4-FIX)
  — the same class as the academy F6 exception. The acceptance-not-fix rule bent for an archived-milestone regression
  without becoming a license to expand scope.
- **The academy F6 landed keyless + zero-academy-repo-edit.** Using the academy's own `e2e_persona` cookie bypass
  (both server `BENCHMARK_VISUAL_BYPASS=1` + client `NEXT_PUBLIC_E2E_AUTH=1` gates) + a cockpit-set cookie
  (port-agnostic on localhost) delivered a signed-in member session with no real Clerk keys and no academy-side route.

## What Didn't
- **A manifest change that alters a shared gate must re-run EVERY gate it can affect, not just the one being tuned.**
  M51 added an unconditional seedPath to the shared manager manifest while only re-running its own (Northwind) gate,
  silently breaking the M50 (Cervato) gate. The lesson: a change to a cross-org/cross-vantage manifest is a
  multi-gate change; its "gate MET" must span every gate that consumes the manifest, or a from-cold acceptance is the
  first (and only) place the drift appears. The org-conditional split is the durable fix — showcase-only surfaces are
  now gated on the org, so a base-org gate can't be broken by a showcase-org prime.
- **Two docs drifted from shipped ground truth.** The AB4 org-conditional manager manifest was undocumented in
  `coverage-protocol.md`, and `ai-readiness.md` still carried the round "~80%/≈160" contract prose vs the shipped
  78.4%/199. Both were caught at close (DOC-1/DOC-2) — a reminder that a gate-behavior change and a shipped-number
  change each owe a doc reconciliation the build pass can miss.

## Carried Forward
- **Origin push (push-gated KEEP → orchestrator/user).** The sole outstanding cross-milestone item: origin has not
  received `main` + the `v1.10` tag + the v1.10 ext tags + the `fit-up-m47..m52` rext tags + `v1.10.1`. Local closes
  deliberately do not push; this is the user's gate. The box-level re-pin (consumption clone + `.agentspace/rext.tag`
  → `v1.10.1`) is DONE. Not a deferral — an administrative KEEP, tracked in `state.md`.
- **Nothing else.** This is the FINAL v1.10b milestone; every carry that pointed at M53 landed here. Next is
  `/developer-kit:close-release`, not another milestone.

## Metrics Delta
(from `metrics.json`) rext stack-seeding Go **786 → 791** (+5; F6 academy DeepLink/AcademyDeepLink build + harden
single-source tests) · demo-stack Python **313 → 326** (+13; F6 authenticated-session + [Academy] deep-link + harden
edge/escape tests) · TS e2e unit **29** (AB4 org-gating + referential-stability edges, +2 vs the pre-AB4 27) · flake
**0** (5/5 Go + 5/5 Python + 5/5 TS). No new dependency. rext release tag **`v1.10.1`** @ `576dbcb` (re-rolled at
close). Close commits: rosetta `d5b8e4b` (doc fixes + deferral audit) + `ab9eb8f` (v1.10.1 re-roll acceptance-record);
no rext production code changed at close (the harden tests predate close; the tag was re-rolled onto them).
