# M256 · iter-08 — progress

**Type:** tik · **Active strategy:** `TOK-01` move 3, unblocked by iter-07 D28. Handler: `ONBOARD-M256-build`.

## Phase 0b / 0d

**0b SKIPPED** (plain tik, standing YELLOW inherited). **0d RUN and PASSED** — `ptvalidate` with `--e2e-dir`
this time, per iter-07's D27 lesson that the two gate tools see the tree differently: VALID, 21 live, 2 TODO.

## Phase A — the whole flow walked live before a line was written

Walked `/onboarding` on the `pt-free` seat, reporting controls at every step:

| step | what is there |
|---|---|
| `/onboarding` step 1 | *"Build Your Career Profile"* — `Import from LinkedIn` / `Import from Resume`, controls `[Upload] [Skip] [Next]`, 1 textbox, **no heading role** (the title is body copy) |
| after `[Skip]` | **exits the flow** and lands the member at `/profile` (their populated profile) |
| revisit `/onboarding` | **redirects to `/home`** |

DB, immediately after: `user_params.onboarding` went from **NULL for all 191 users** to **exactly one row**, for
exactly the hero driven — `{"steps":[{"step":"done","updated_at":"2026-07-28T13:45:36.616Z"}]}`.

**H1 CONFIRMED — the route is its own read-back**, and H2 holds. One measured caveat, recorded rather than
smoothed over: the `Skip` path reaches `done` **directly**, without traversing the curated shared step model's
**Role** and **Skills** steps. Those live on the *import* path. That is why the import use cases are declared
`TODO` instead of being folded into the live one's claim.

## Phase B — what landed

- **`lib/onboarding-page.ts`** — the 10th product's page object, carrying the D28 evidence (where onboarding
  state actually lives, why the audit's F5 was a schema misreading, and why the route is its own read-back).
- **`pt-onboarding-complete`** (`onboarding.completion.UC1`) — a day-0 member completes onboarding, lands in the
  app, and revisiting `/onboarding` **redirects to `/home`**. Mutating **#6**; its negative control is the same
  route asserted in the opposite direction *before* the action (the flow is SERVED → the outcome is absent), so
  the final assertion is proven to discriminate the completion rather than match a route that always redirects.
- **`ONBOARDING_URL` + `isOnOnboarding`** added to the single-sourced `url-shapes` module, with **lockstep pin
  tests** including an explicit segment-anchoring test (an `/onboarding-tour` look-alike must NOT match). That
  anchoring is load-bearing *because* the pattern is asserted in both directions — a look-alike match would make
  the "completion persisted" half pass on the wrong route.
- **`manifest/onboarding.yaml`** — the product, plus **all 5 curated M201 onboarding use cases declared with
  written verdicts** (clause 3's zero-silent-gaps requirement for this cluster). Every verdict is
  **harness/seed work, none `unimplementable-without-platform-edit`** — which is the concrete reason the
  re-scope trigger did not fire:
  - `enterprise-workforce-standard.UC1` — needs the **import** path + a **résumé fixture** (spec §5.4's
    `fixtures/` dir is reserved and still EMPTY; no shipped Playthrough exercises a file upload) + a real async
    LLM import, so its boundary must be import-completed, not the extracted values (P6).
  - `enterprise-workforce-standard.UC2` — needs the org-prepared summary shown *instead of* the import form;
    iter-08 measured the **import form for a hero WITH a populated profile**, so that variant's trigger
    condition is **not yet identified** and must be found before the UC can be honestly asserted.
  - `individual.UC1` — the one place **F5 has a kernel of truth**: its curated actor is org-less and
    `UsersSeeder` does write a membership for everyone, so it needs a member-less user + a roster seat
    (`stack-seeding` work). Its curated final is already proven in the org context by the live UC1.
  - `enterprise-workforce-ai-readiness.UC1` — the member-facing guided flow is hosted on `/home`, not its own
    route; needs a seat in Org C with onboarding incomplete **and** funnel stage 0.
  - `enterprise-hiring.UC1` — the only onboarding UC whose final spans **two apps** (it must land in
    `apps/hiring` on offset-3001, not the workforce app); needs a day-0 seat in a hiring-flagged org.

**Honesty note, stated in-manifest:** `onboarding.completion` is **net-new**, not one of the M201 curated 28. It
grows the MANIFEST denominator, not the curated one — clause 3's `12 of 28` arithmetic is unaffected.

## Phase C — the gate run

**3 consecutive `run-playthroughs.sh 2 --reset` runs: `145 passed` each. 0 flake, 0 red.**
`ptreport`: **22/29 passing, 0 failing, 7 `[TODO]`, 0 unimplementable.** `pt-onboarding-complete` green on all
three (3.7 s). `@pt-mutation` registry, computed: **`MUTATES=6  READ-ONLY=14  UNKNOWN=2`** (22 Playthroughs).

### Clause 1 re-verified on the grown denominator (20 non-studio)

| Figure | Value |
|---|---:|
| **Median per non-studio Playthrough — the GATED metric** | **1.979 s** |
| **Ratio vs the iter-02 baseline (3.326 s)** | **0.5950×** — gate `<= 0.79×` **MET** |
| Honesty cross-check, the ORIGINAL 16 only | 1.772 s = **0.5328×** |
| Studio lane (excluded) | 1.18 s / 1.85 s |
| Suite wall-clock (REPORTED, not gated) | 55.3 / 56.6 / 52.6 s → median **55.3 s** |

**This also corrects iter-06's own caveat.** iter-06 measured the original-16 cross-check at 2.006 s (0.6033×)
and attributed the drift from iter-04's 0.5347× to laptop variance. The same cross-check now reads **1.772 s =
0.5328×** — within noise of iter-04. So the drift *was* variance and **not** a regression, as suspected but
correctly not asserted at the time. Environment unchanged: `Kirality-Mac-Pro-6.local`, Docker VM **9.70 GiB**,
`demo-2` offset 20000, localhost/http, `workers: 1`, `retries: 0`. No number here is comparable to billion's
(D-v28-12).

rext commit + tag **`fast-build-m256-onboarding`**, pushed to origin.

## Close — 2026-07-28

**Outcome:** **the onboarding product exists** — the last whole surface in the M201 curated corpus that no e2e
suite had ever touched, and the one the milestone's own audit had declared impossible. 1 live Playthrough
(mutating **#6**, with its negative control free from the route flip) + **all 5 curated onboarding UCs declared
with written verdicts**, every verdict harness/seed work rather than `unimplementable`. Clause 1 improved to
**0.5950×** and iter-06's suspected regression is retired as variance.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — clause 1 **met** (0.5950×, 0 flake ×3); clause 2 mutating **6/5 MET**, negative controls
**6 of 22**, `blocked` **0**; clause 3 org-admin 2/4 + onboarding **1 live and 5/5 curated verdicts written**,
verdicts still owed for `workforce.organization-feedback`, `profile-skills.import`, `talk-to-data.query` and the
M206/M207 reservations; **D-v28-5** unstarted.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this tik progressed: onboarding 0 → 1 live + 5 verdicts, mutating 5 → 6) — (3) re-scope: n (**0** UCs `unimplementable`; all 5 onboarding verdicts are harness/seed work) — (4) user-blocker: n (145 passed ×3; D-v28-3's consolidated red set is empty) — (5) cap-reached: n (3rd tik of run 2) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D33 (the route is its own read-back — one route gives a mutating Playthrough both halves), D34
(the `pt-free` seat choice is load-bearing, not incidental), D35 (the Skip path reaches `done` without the
curated Role/Skills steps — so the import UCs stay TODO rather than being folded in), D36 (all 5 curated
onboarding verdicts are harness/seed work; F5's kernel of truth is confined to `individual.UC1`).
**Side-deliverables:** none.
**Routes carried forward:**
- `NEGCTL-M256-cross-vantage` → **next iter.** Clause 2's largest remaining gap: negative controls **6 of 22**.
- `VERDICT-M256-remaining-uncovered` → **next iter.** Clause 3's other half: written verdicts for
  `workforce.organization-feedback`, `profile-skills.import`, `talk-to-data.query`, plus the 5-release-old
  M206/M207 reservations. The onboarding block in `manifest/onboarding.yaml` is the template.
- `D-v28-5-cockpit-logout` → still unstarted; it is a gate clause in its own right.
- `ONBOARD-M256-import-path` → a later tik. The 4 remaining onboarding UCs, each with its diagnosis already
  written in-manifest. Note the résumé-fixture UC would be the **first** Playthrough to use `fixtures/`.
- `BLOCKED-M256-refusal-surface`, `FIX-M256-studio-false-green` (re-aimed), `DOC-M256-llm-lane-premise`,
  `PT-M256-orgadmin-role-create`, `PT-M256-orgadmin-member-tag`, `FENCE-M256-bounded-interaction` — all stand.
**Lessons:**
1. **A route that gates on state is a free read-back.** `/onboarding` serves-or-redirects, so it supplied the
   pre-state absence and the persisted post-state on one URL with no extra surface, no toast, and no DB assert.
   Look for serves-or-redirects routes before building a read-back — they are the cheapest mutating proof shape
   in the suite, cheaper even than iter-06's label flip.
2. **Choose the seat, do not inherit it.** Completing onboarding cannot be undone through the UI. Driving it on
   `pt-employee` would have coupled this Playthrough to every other one that asserts on that hero. `pt-free` was
   registered and used by nothing — the audit had noted that as a *gap*; it turned out to be the asset.
3. **Declare the neighbours you did not build, with reasons.** Four of five onboarding UCs are still TODO, but
   each now carries the specific missing piece. That is what makes "0 of 5" into "1 live + 4 priced" instead of
   a silent gap — and writing them down is what proved none of them was `unimplementable`.
