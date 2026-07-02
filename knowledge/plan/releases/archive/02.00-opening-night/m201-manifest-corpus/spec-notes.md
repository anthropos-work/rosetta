# M201 Spec Notes

Technical notes accumulate here during the (user-guided) iter loop. The authoritative design lives in
[`overview.md`](overview.md) + the consolidated capability spec
[`knowledge/plan/spec-drafts/playthroughs/spec.md`](../../../spec-drafts/playthroughs/spec.md) (v0.3 — esp. §2
model, §4 use-case shape, §5.3 manifest format). M201 is **prose only** — it authors the manifest YAML corpus; it
builds **no** code, **no** seed, **no** validator (those are M202).

## Candidate product surface — the STARTING outline (from the v2.0 research)

> **This is a starting outline the USER will direct, NOT a fixed list.** It seeds the first top-down pass; the
> user adds / drops / re-depths products, stories, and use cases.

### (a) Demo-covered products — the world the current seed already supports
These map onto the existing stories & heroes (Maya employee / Dan manager) + the set-dressed catalog, so they are
the readiest first slices:

- **Skill Paths** — browse → enroll → complete → verify-skill (Maya's core employee journey).
- **AI Simulations** — chat / code / document launch → complete → score-in-range (**NON-voice**; voice/recording
  is the future mirror tier — see §5.8 and **M206**).
- **Profile & Skills** — the verified-skill chart + the claimed-vs-verified gap + work/education timeline.
- **Workforce Intelligence** — the org mapped→verified funnel + member roster + member drill-down
  (activity-dashboard) + succession / at-risk (Dan's core manager journey).
- **Hiring** — recruiter candidate-pipeline journeys (post a role → review applicants → advance a candidate). The
  demo touches it lightly; the goal-aligned depth needs a recruiter persona + a candidate pipeline (a
  **dedicated-seed need** → hand to M202; future recruiter-vantage coverage is **M205**).
- **Academy** — the ant-academy learning surface (a separate deployment; Clerk-only). Demo-present but a distinct
  target environment (future Academy coverage is **M207**).

### (b) Goal-aligned areas the demo barely covers — FLAG as "to confirm with the user"
These are where the goal likely demands proof but the **current minimal / partially-aligned demo seed** does not
reach. **Confirm scope + depth with the user**; each that lands will likely carry a **dedicated-seed need handed
to M202**:

- **Auth & Onboarding** — sign-up / sign-in / org-join / the onboarding flow (note: hero login is the demo-only
  Clerkenstein seat-switch — spec §5.4; a *real* sign-up flow is its own question to confirm).
- **Billing & Entitlements / tier-gates** — free → paid gates, `actor.entitlement` (anon / free / paying /
  enterprise / expired), "a free user cannot open paid content" (`outcome: blocked`). Needs a Stripe
  test-mode / assertion-boundary (spec §5.8) + tier-spanning seed → M202 / future tier-gate coverage (**M205**).
- **Org Admin & Settings** — org-level configuration, member management, roles/permissions (Sentinel-gated
  `blocked` outcomes — spec §4.2).
- **Cross-product journeys** — e.g. **candidate → employee** (a hire flows from Hiring into the employee
  experience): a *Story* that spans products (spec §2 — stories may span products).

> Note: §4.3 entitlement tiers + private-path isolation, and §4.4 engine coexistence (legacy skill paths →
> new-academy engine), are **model axes** the use cases declare — relevant to (b) above. The corpus declares them;
> M202's seed must span the tiers + multi-org-private it names.

## The use-case shape to write (spec §4)
Each use case carries (prose-intent — *what*, never *how*): `id` · `goal` · `actor` (+ `actor.entitlement`) ·
`seed` / `preconditions` (a **named** world — M202's seed must provide it) · `flow` (high-level steps) ·
`outcome` (`success` / `blocked` / `error`) · `expectations.intermediate[]` (ordered, labelled) ·
`expectations.final` · `playthrough` (`TODO` until M203/M204 builds it). Layout: **one YAML file per product**
(§5.3).

## Validate against (each pass)
- **The real platform surface** — does this product/story/use-case map to something the platform actually does?
  (Cite the corpus service docs / the live demo where useful.)
- **The manifest model** (§2/§4) — is each use case the atomic unit of functional truth (one goal + flow +
  expectations)?
- **The §5.3 validator** — unique ids, both-way id integrity (every UC ↔ a `playthrough` id or `TODO`),
  precondition-coverage (every `seed`/`preconditions` resolves — *to be enforced once M202's validator + seed
  exist; until then the manifest declares the named precondition and the need is handed to M202*).

## Where the YAML lands
- **Until M202 exists:** drafted under **this milestone dir** (e.g. `manifest-draft/<product>.yaml`) or
  `knowledge/plan/spec-drafts/playthroughs/manifest-draft/`.
- **Once M202 lands the `playthroughs` rext section:** the corpus moves into that section (its manifest home),
  authored + tagged per the tooling policy.

## Open questions (carry through the user-guided passes; record resolutions in decisions.md)
- The (b) "to confirm with the user" set — which of Auth & Onboarding / Billing & tier-gates / Org Admin /
  cross-product land in the v2.0 corpus, and at what depth.
- For each landed use case whose precondition the demo lacks — the exact named precondition to record + hand to
  M202 (the dedicated-seed expansion).
- "Complete enough" — the user's bar for the sign-off (which products are in scope for v2.0 vs deferred to a
  future minor / a future major milestone number).
