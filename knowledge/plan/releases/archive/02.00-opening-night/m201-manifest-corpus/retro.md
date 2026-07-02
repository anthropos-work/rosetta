# M201 Retro — Manifest corpus

**Closed:** 2026-06-29 · `closed-on-gate` · iterative, user-guided.
**Outcome:** the prose-intent Playthroughs manifest corpus — 9 products · 26 stories · 28 use-cases — authored,
adversarially re-grounded, and signed off; v2.0 then paused for a v1.10 backfill triggered by what this milestone
surfaced.

## What went well
- **The one-product-per-pass cadence held.** Top-down outline → ground against the real surface → write YAML →
  user OK → next. Terse, minimal-conversational, exactly as the user asked. No "vomit it all at once."
- **Grounding-as-you-go caught traps before sign-off.** AI Labs (nil-client) was caught by hand during the
  Talk-to-Data-adjacent pass; the path-migration spec grounded legacy-vs-academy precisely.
- **The adversarial verification was the highest-leverage move of the milestone.** 11 agents, ~1.3M tokens, and it
  (a) resolved all 4 open flags with code evidence, (b) caught a 2nd AI-Labs-class trap, (c) **discovered the
  stale-clone drift** — a finding worth far more than the milestone itself, because it invalidates the foundation
  the last release was built on.

## What was hard / what to carry
- **The grounding baseline was stale and we didn't know it.** The whole verify graded against next-web @ v2.33.2
  (115+ commits behind). The member-AI-readiness "not-runnable" was a false negative only because the *user* knew
  the feature ships in prod. **Lesson:** before trusting any "does-it-run" judgment, check the clone's distance
  from prod first. This is now the backfill's milestone 1.
- **A signed-off corpus is only as true as its grounding.** The sign-off stands as a *completeness* judgement
  (the right set of journeys), but the per-UC *runnability* annotations carry a staleness caveat until re-grounded.
- **Single draft file, not one-per-product.** Accepted interim per `overview.md`; the §5.3 split + validator are
  M202. Don't let the single file ossify — M202 splits + machine-validates it.

## Incidents this cycle
- **Stale-clone drift (systemic).** stack-demo clones 5 weeks / 115+ commits behind prod; corpus likewise lagging
  shipped features (member AI-readiness). Root of the v1.10-backfill pivot. (Not a code bug — a sync/process gap.)

## Hand-offs
- **→ v1.10 backfill (user-driven):** re-sync clones + re-ground corpus to current prod, then re-validate every
  negative verify verdict. The verify report (`wvpnpvozh`) + the manifest's staleness caveat are the inputs.
- **→ M202 (foundation):** the manifest's appended "M202 seed/wiring corrections" + "coverage holes" blocks; the
  §5.3 validator + one-file-per-product split; the dedicated decoupled Playthrough seed satisfying the declared
  preconditions.
- **→ v2.0 resume:** this corpus is the spec; re-open after the backfill re-grounds it.
