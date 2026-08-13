# M256 · iter-09 — progress

**Type:** tik · **Active strategy:** `TOK-01` move 4. Handler: `VERDICT-M256-remaining-uncovered`.

## Phase 0b / 0d

**0b SKIPPED** (plain tik). **0d RUN and PASSED** — `ptvalidate` with `--e2e-dir` (per iter-07 D27): VALID,
22 live, 7 TODO.

## Phase A/B — the verdict register, and what pricing exposed

**[`../coverage-verdicts.md`](../coverage-verdicts.md) — the authoritative answer to *"why is this curated use
case still not covered?"* for all 28 M201 curated use cases.** Clause 3's verdict requirement is now
**COMPLETE: 0 curated UCs carry no verdict.**

The distribution, which is the substantive result:

| Verdict | count | UCs |
|---|---:|---|
| Buildable next, no blocker | 1 | org-feedback — **landed this iter** |
| Buildable, blocked on a checked-in fixture | 1 | `profile-skills.import` |
| Buildable, blocked on a live third-party credential | 1 | `talk-to-data.query` |
| Reservation confirmed (integration / mirror tier) | 3 | ai-sim code · ai-sim interview · academy |
| Reservation re-examined → recommend **re-homing** | 1 | `profile-skills.self-evaluation` → M257 |
| **`unimplementable-without-platform-edit`** | **0** | — |

**H1 confirmed, and it is the load-bearing finding: zero of the 28 curated use cases is unimplementable.** Every
uncovered one is blocked on harness work, a fixture, a credential, or a deliberate tier reservation — never on
the platform refusing to be driven. That is the strongest available answer to clause 3 and the concrete reason
the re-scope trigger never fired.

Three verdicts are worth more than a category:

- **`profile-skills.import`** and **`onboarding.enterprise-workforce-standard.UC1`** share **one** blocker: a
  checked-in **résumé fixture**. `playthroughs/fixtures/` has been *reserved and empty* since spec §5.4 — **no
  shipped Playthrough has ever exercised a file-upload flow** — so the first one pays the fixture *and* the
  real-file-chooser pattern. Recorded as a **pair** (`PT-M256-resume-fixture-pair`) so that cost is paid once.
- **`profile-skills.self-evaluation`'s M206 reservation is WEAK.** Its curated final is *"the self-assessment
  persists and shows on the profile (`user_skill_evidences.user_level` set)"* — a persist-then-observe shape,
  i.e. exactly the MUTATING shape clause 2 spent four iters hunting, needing no LLM, no integration, no fixture.
  It was one of iter-05 D20's three nominees and lost only because iter-06 found cheaper ones. **Recommended for
  re-homing to M257** — recorded as a recommendation, because re-homing a reservation is a roadmap decision, not
  an iter's call.
- **`skill-paths.academy`'s blocker is sharper than M201 recorded.** M236 iter-08 established that a demo academy
  has **no `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`**, so it serves its **committed filesystem catalog** and seeded
  `academy_chapter_progress` rows have **no reader**. So it is not merely "unseeded": on a demo the academy is
  structurally not DB-backed, and no amount of seeding changes that. Reservation confirmed with that attached.

**H2 paid — and this is why verdict work is not paperwork.** Pricing `workforce-intelligence.organization-feedback`
forced the question *"is its data already seeded?"* — and it is: `stackseed` reports **`feedback rows=70`** on
every `pt-world` reset, on a real non-flag-gated route with a wired backend. So the UC that had been un-homed for
five releases cost **a page object and a spec**. It was landed in this iter:

- **`pt-workforce-org-feedback`** (`workforce-intelligence.organization-feedback.UC1`) — probed live first (h
  *"Users Feedback"*, a *"Feedback distribution"* recap at 23 items / 61 % positive / 39 % negative, 1 table with
  21 rows of real member names, polarity filters, an XLSX export). It asserts **both** sentiment polarities
  deliberately: a one-sided recap is exactly the empty state the curated final rules out, and a single-polarity
  assert would pass on it. Plus a **cardinality** assert on the table (the M42e semantic-gate lesson: *"the table
  rendered"* is not the claim *"the table rendered real org-scale feedback"*), and the filter interaction the
  curated flow ends on — narrowing to real rows, not to an empty surface.
- A **separate data axis** from the three existing WI Playthroughs: it aggregates post-simulation feedback
  (`createJobSimulationFeedback`), where they aggregate the funnel, roster and succession.

## Phase C — the gate run

**3 consecutive `run-playthroughs.sh 2 --reset` runs: `146 passed` each, 0 failing, 0 flake.**
`ptreport`: **23/30 passing, 0 failing, 7 `[TODO]`, 0 unimplementable.** `@pt-mutation` registry, computed:
**`MUTATES=6  READ-ONLY=15  UNKNOWN=2`** (23 Playthroughs).

### Clause 1 re-verified on the grown denominator (21 non-studio)

| Figure | Value |
|---|---:|
| **Median per non-studio Playthrough — the GATED metric** | **1.880 s** |
| **Ratio vs the iter-02 baseline (3.326 s)** | **0.5652×** — gate `<= 0.79×` **MET** |
| Honesty cross-check, the ORIGINAL 16 only | 1.757 s = **0.5284×** |
| New Playthrough this iter | `pt-workforce-org-feedback` 1.82 s (below the median) |

Environment unchanged: `Kirality-Mac-Pro-6.local`, Docker VM **9.70 GiB**, `demo-2` offset 20000, localhost/http,
`workers: 1`, `retries: 0`. No number here is comparable to billion's (D-v28-12).

rext commit + tag **`fast-build-m256-verdicts`**, pushed to origin.

## Close — 2026-07-28

**Outcome:** clause 3's verdict requirement is **COMPLETE** — all 28 curated use cases carry a written verdict,
and **0** are `unimplementable-without-platform-edit`. Pricing them properly exposed one that had been un-homed
for five releases as costing only a page object and a spec, so it was landed in the same iter
(`pt-workforce-org-feedback`), and surfaced two structural findings: the résumé fixture is a **shared** blocker
for two UCs, and `profile-skills.self-evaluation`'s M206 reservation is weak enough to recommend re-homing.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — clause 1 **met** (0.5652×, 0 flake ×3); clause 2 mutating **6/5 MET**, negative controls
**6 of 23**, `blocked` **0**; clause 3 **verdict half COMPLETE (28/28, 0 unimplementable)** but the LANDED half
short — org-admin 2/4 and onboarding 1 live of 5; **D-v28-5** unstarted.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this tik progressed: 7 verdicts written + 1 UC landed) — (3) re-scope: n — **and the register settles it: 0 of 28 curated UCs are `unimplementable`, so the trigger cannot fire on current evidence** — (4) user-blocker: n (146 passed ×3; consolidated red set empty) — (5) **cap-reached: n at grading time — this is the 4th tik of run 2** — (6) protocol-stop: n — Outcome: continue
**Decisions:** D37 (the verdict register: 0 of 28 unimplementable — the re-scope trigger cannot fire on current
evidence), D38 (the résumé fixture is a SHARED blocker for two UCs — pay it once, land them together), D39
(`profile-skills.self-evaluation`'s M206 reservation is weak; recommend re-homing to M257 — a roadmap call, not
an iter's), D40 (verdict work is not paperwork: pricing org-feedback exposed already-seeded data and turned a
five-release-old gap into a same-iter landing).
**Side-deliverables:** none.
**Routes carried forward:**
- `NEGCTL-M256-cross-vantage` → **next.** Clause 2's largest remaining gap: negative controls **6 of 23**.
- `D-v28-5-cockpit-logout` → **next.** A gate clause in its own right, unstarted across 9 iters.
- `PT-M256-resume-fixture-pair` → the two fixture-blocked UCs together (`profile-skills.import` +
  `onboarding.enterprise-workforce-standard.UC1`); would be the suite's FIRST file-upload Playthrough.
- `PT-M257-self-evaluation` (re-home recommendation) · `PT-M257-talk-to-data` (needs Bedrock creds).
- `ONBOARD-M256-import-path`, `BLOCKED-M256-refusal-surface`, `FIX-M256-studio-false-green` (re-aimed),
  `DOC-M256-llm-lane-premise`, `PT-M256-orgadmin-role-create`, `PT-M256-orgadmin-member-tag`,
  `FENCE-M256-bounded-interaction` — all stand.
**Lessons:**
1. **Pricing a gap is how you discover it was never expensive.** `organization-feedback` sat un-homed for five
   releases. The verdict work forced one question — *is its data already seeded?* — and the answer (`feedback
   rows=70`) collapsed it to a page object and a spec. Write the verdict before assuming the reservation.
2. **Verdicts must name the specific missing piece, not a category.** "Needs seed work" would have hidden that
   two separate UCs are blocked on the *same* fixture, and that one M206 reservation needs no mirror engine at
   all. The specificity is what makes the register actionable instead of decorative.
3. **Re-examine inherited reservations rather than re-inheriting them.** Three of four M206/M207 reservations
   held and one did not. Carrying all four forward unread would have kept a cheap mutating Playthrough
   unreachable for a sixth release.
