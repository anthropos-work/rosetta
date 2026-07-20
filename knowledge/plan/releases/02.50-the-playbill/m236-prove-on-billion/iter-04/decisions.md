# iter-04 — Decisions

## D1 — exact-path visits, not 13 `VantageManifest`s (the scope SHRANK)

The kb-fidelity audit's B3 finding — *"`VantageManifest.identityKey` is singular ⇒ 13 seats → 13
manifests"* — is arithmetically right but assumes the crawl machinery is the tool. It is not.

`cockpit-login.loginAs()` already accepts **`landingPath`**, so a content seat can be landed directly on
its exact manifest-supplied result URL. A content-story result is *inherently* exact-path: the URL carries
a session uuid only the seed knows, and it is reached from the cockpit, never by a link a crawler could
follow. So the sweep is **one runner + one page-object**, not 13 manifests.

The user accepted B3's "enlarged Cluster 1" as the milestone's centre of gravity. That acceptance still
holds — the *centre* is right — but the **cost** is materially lower than the audit projected. Recorded so
the record does not carry a scope estimate that the implementation refuted.

## D2 — shape is selected by ROUTE, because keyword-selection produced a false PASS

The first full sweep graded skill-path pages with the scored-simulation shape:

- `sp-genai-in-progress` **failed** — "no feedback / evaluated-skills section". Correct-looking, but the
  page is fine; a skill-path page has no evaluated-skills section by nature.
- `sp-product-manager-completed` **passed** — and this is the dangerous one. It did not render a report.
  It passed because the shape's `hasFeedback` test is `/feedback/i` and the page carries **11,250
  characters** of legitimate path content, somewhere in which the word appears.

A gate that passes for the wrong reason is worse than one that fails: the failure gets triaged, the false
pass ships. Fixed by adding a calibrated `player-skillpath` shape (chapter/path structure **and** a
progress indicator) and selecting the shape from the **route prefix** (`/skill-path/` vs `/sim/…/result/`),
which cannot drift the way content-sniffing does.

## D3 — B4 interpreted as "correct the premise, KEEP the rule" (not a blanket deletion)

The user's B4 decision: amend `coverage-protocol.md` in the same change that reverses its `skipPaths`
`/result/` exclusion. Executing that literally — deleting `RESULT_DEEP_LINK_SKIP` — would be wrong on the
evidence, so the reversal was applied where it actually belongs:

- **What was false:** the rule's stated *reason* — result pages are *"a runtime-computed AI evaluation …
  never written by a seed"*. M231 refuted it (a **persisted read**); M236 proved the opposite live (13/13
  render). That premise is now corrected in the doc **and** in `coverage.spec.ts`'s comment.
- **What remains true:** a hero's BFS crawl should not dive into arbitrary historical session deep-links.
  That is **crawl scope** — an independent and still-valid justification for the same rule.
- **Why not delete it:** deleting proves nothing extra. The hero sweep would simply re-walk sessions that
  the content-stories sweep already asserts far more precisely, at the cost of budget and noise.

So the *documented rule M236 had to reverse* is reversed — the claim that these pages are unprovable is
gone, and the doc now names the harness that proves them. This is an interpretation of a decided item, not
a re-raise of B4.

**The generalizable finding** (recorded in the doc): the bullet stood unchallenged for four releases
because it was written as a **fact about the platform** when it was really an **inference from one
observation** (a seeded session's result page was empty — because nothing had seeded the result rows yet).
Record the observation and the explanation separately; the observation stays true, the explanation rots.

## D4 — the manager defect is a READ-PATH defect, and is characterized, not fixed here

13 of the 15 non-landing pairs are the manager vantage. Evidence gathered:

- The **mirror row exists** — `local_jobsimulation_sessions` 13/13 (iter-03), with `score`, `status=ended`,
  `completition_status=passed`, correct `organization_id`.
- The **user row is correct** — `d541c46c-…` is `Clara Romano`, and the *player* page renders "Congrats,
  Clara!" from the same data.
- Yet the manager page renders **`undefined undefined`** for that name and **"No data"** for the attempts
  table — while its header correctly names the sim and its measured skills (so the sim-side read works).

So the manager scoreboard's read path resolves the *player* differently from the player's own page, and its
attempts query returns nothing despite present rows. **The mirror is necessary but not sufficient** — a
finding that generalizes beyond this milestone and is recorded in `playthroughs.md`.

Not fixed in this iter: it is a distinct line of investigation, and the scope-creep tripwire applies (this
iter already ran a declared 4-step shape). Routed to **iter-05** with handler
`MANAGER-M236-iter05-scoreboard` — Fate 3, named handler, next iter.

The 2 interview manager pairs fail differently (**no header at all**, rather than an empty table), which
suggests a separate route/tab for interview manager reports — likely related to M231's
`flag_interview_manager_report` gate. Same handler; distinguish at triage.
