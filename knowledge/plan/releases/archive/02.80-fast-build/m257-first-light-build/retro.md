# M257 — retro

**`closed-on-gate`, 2026-08-12.** p50 **286.99 s** on `macmini` against a **360 s** gate and a **300 s**
stretch; n=3; every clause green on all three reps, including both falsifiable ones.

## Summary

Nine iters (7 tiks + 2 toks) to collapse a cold demo bring-up. **One lever landed of eight priced, and it
cleared the gate alone** — L1, multi-staging the two Next images to `.next/standalone` from rext-owned
Dockerfiles, worth **−141.63 s** on a UI tier that had been 54.8 % of the cycle. The close then landed the
milestone's one remaining declared deliverable (the §8.5 corpus retraction, with achieved numbers, once)
and found three fail-opens in the gate instrument itself.

## Incidents this cycle

- **The milestone was paused for eleven days on a claim that was false.** *"A Mac is arm64/overlay2, so it
  pays no unpack leg, so L1 is worthless locally"* — derived from `docker info` printing
  `Storage Driver: overlayfs`, generalised from a **different, retired** machine. Measured on the host that
  exists: containerd image store, a size-proportional unpack leg (0.8 s @ 256 MB → 3.0 s @ 1024 MB; 19.3 s
  on the real 4.12 GB image). L1 was worth ~141 s. **A config string is not a hardware measurement.**
- **Three consecutive iters closed "metric delta 0, by design" and every one was accurate** — because the
  gate named `odysseus`, retired one day after `TOK-01` named it. No measurement taken anywhere could have
  satisfied it. The triggered tok caught it; nothing else could have.
- **The whole-tree `stack-core` sweep did not return at the final harden** (55m12s against a 32–57 min band
  on this contended box, vs 34m56s for the same section one iter earlier — contention, not a hang) and was
  terminated as session cleanup rather than left orphaned. **Not "swept clean"**, and the harden said so.
- **Seven fabricated figures, caught before commit.** Seven sub-phase p50 cells in the first draft of the
  new `build-budget.md` table were written from memory and were wrong. Every one was checked against the
  raw `campaign.json` and corrected. No P2; nothing shipped.
- **A mutation battery went RED mid-close** because a fix moved its subject — `§5` rule 53 doing exactly
  its job on my own edit. Re-pinned, plus a net-new mutant for the arm that replaced it.
- Regressions: **none**. Flakes: **none** (3 consecutive clean runs with random ordering enabled).

## What went well

- **Landing the falsifiable assert WITH the lever that can trip it** (`TOK-01`) paid, and the harden found
  the first hard evidence: L1's own `.env` finding routes to a `.dockerignore` repair whose obvious form
  re-creates the M218 incident — and that is precisely the `foreign_pk` arm ISOLATION books.
- **Re-grading the headline instead of trusting it.** The three gate reps were re-graded under post-harden
  code, then again under the close's three instrument fixes. p50 unchanged both times. A number that has
  survived two independent re-gradings under changed code is a different kind of number.
- **Measuring before fixing.** The M255-inherited body-extraction defect turned out to be **latent, not
  live** — the old and new parses return the identical offset on the shipped script. That measurement is
  why it shipped as a *fence* rather than a *correction*, and the distinction is now written down.
- **The fences caught the closer.** The host-parameterised mirror fence caught an un-hosted `449.51 s` in a
  table I had just written; the net-new §8.5 fence caught two of my own sentences and a case-sensitivity
  hole in its own marker list. Prose was fixed, fences were not weakened.
- **The first STABILIZED harden across both milestones** — pass 3, coverage delta 0. M257x closed 18
  consecutive passes at cap-without-stabilization.

## What didn't

- **The close found three must-fix fail-opens in the gate instrument**, all in code this milestone wrote,
  none caught by nine iters of per-commit review or by three harden passes. They only became visible when
  the milestone's code was read *as a whole*: clause 3 held the opposite policy to clause 1 on the identical
  question, and `isolation_ok` — the milestone's own new gate clause — was computed and read by nothing, so
  `buildbench report <dir>` printed `gateable: true` over a directory that never asserted it.
- **The fixture that hid it had hidden the same class one iter earlier.** `_ledger` calls itself *"the
  shape `run_campaign` really writes"*, and its own docstring narrates how omitting `host_identity` let the
  aggregate ignore that field for a whole iter. It then omitted `isolation` for exactly as long.
- **The §8.5 retraction's work list was wrong in three ways before it was executed**: two cites pointed at
  the wrong lines, four had drifted +61, and a seventh live site was never enumerated at all. **A site list
  is a snapshot of where a claim was on the day someone looked.**
- **Four items M255 routed here specifically because *"M257 is the milestone that actually exercises each
  of them"* went nine iters untouched** and were only re-fated at the close. The routing rationale was
  sound; nothing made it actionable inside the iter loop.
- **M257x's carry-forward had never reached M258's `overview.md`** — the `BIND_HOST` failure recorded *in
  that same file*, one section above the gap, repeated by the very next close.

## Carried forward

`closed-on-gate`, so **no `carry-forward.md`**. All Fate-3 items are recorded at their destination in
**M258**'s `overview.md`, and the per-item audit is in `decisions.md` § deferral re-audit.

- **`LEVER-M257-L5-setdress` → M258** — the ranking moved: `set_dress` is now the largest single phase at
  **82.04 s = 28.6 %**, where the plan priced L5 at ~30–50 s and ranked it fifth. M258's 480 s ceiling is
  explicitly reachable *"only if M257 spends part of its unspent levers"*; this is that reserve.
- **`FIX-M257-dockerignore-env-pattern-unpaired` → M258** — with the warning attached: the tidy one-line
  fix bakes the **real** Clerk key.
- **Instrument + hygiene → M258**: `FIX-M257-campaign-kill-orphans-bringup` ·
  `FIX-M257-sampler-disk-units-vm` · `MEASURE-M257-macmini-true-idle` · `PROFILE-M257-provisional-fields` ·
  `FIX-M257-anchor-guard-content-drift` · `FIX-M257-census-interpreter-namespace-import` ·
  `RATCHET-M257-literal-ceilings-breached` · `FIX-M257-demopatch-sha-baselines-drifted` · plus two net-new
  at this close: `FIX-M257-frontend-floor-is-billion-shaped` and
  `FIX-M257-image-listing-conflates-empty-and-unreadable`.
- **Dropped:** `INVESTIGATE-M257-load1-48` (un-reproducible; host retired). **Superseded:**
  `FIX-M257-committed-env-ships-real-clerk-pk`.

## Metrics delta

From [`metrics.json`](metrics.json): **449.51 → 286.99 s p50** (−36.2 %) on `macmini`; UI tier
**246.23 → 104.60 s**; next-web **4.04 GB → 417 MB**, hiring **3.94 GB → 380 MB**; export leg
**136.4 → 3.8 s**. Tests **+35** at the harden, **+25** at the close. Flakes **0**. Escape-hatch deferrals
**0**. Platform-repo edits **0**.

## The one lesson worth carrying

**A gate that names a dead thing does not fail — it abstains, and abstention is invisible.** Three iters
reported an accurate *"metric delta 0"* and none could say *"and no delta is achievable"*, because the
un-gradeability lived in the gate's **subject**, not in any of its clauses. **Check a gate for
GRADEABILITY before checking it for satisfaction.**
