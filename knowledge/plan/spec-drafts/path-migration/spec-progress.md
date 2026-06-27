# Skill Paths — Spec Progress (open points tracker)

> **Status:** Draft · spec-draft · 2026-06-25
> Tracker + decision log for the consolidated spec, [`spec.md`](spec.md). We worked the decisions **one at a
> time** and record each here.

**Legend:** 🔴 not decided · 🟡 discussing · ✅ decided · ⏭️ deferred (→ [`next-release.md`](next-release.md))

| # | Topic (plain English) | Status | Decision |
|---|------------------------|:------:|----------|
| 1 | Bite-size vs. long content — how do 40-min videos / courses fit a "3–5 min pills" product? | ✅ | **Option D.** module = unit of content (~20–40 min); **lecture** = bite-size atom (~3–5 min). Long resources stay whole, auto-sliced (~5 min) into lecture-checkpoints. See [`content-model.md`](content-model.md). |
| 2 | How content is structured — and what a legacy "step" becomes in the new model | ✅ | Model is `path › chapter › module › lecture`; legacy **step → module** (typed), bite-size lives in **lectures**. Two module families: Authored {ucourse, simulation, AI lab} / Resource {video, PDF, podcast, link}. Udemy dropped. See [`content-model.md`](content-model.md). |
| 3 | Simulations (interactive, graded exercises) in the new product | ✅ | Reuse existing engine; launch the real sim & return (A); completion via server-side event (1); **all-in** (required before cutover). Assessment sim / AI lab must be **passed** to close its chapter. AI labs = new asset, treated like sims (not in legacy or new today). See [`content-model.md`](content-model.md). |
| 4 | Proving skills — does finishing a path still "verify" a skill on the new product? | ✅ | Yes, in scope. Skills tagged at **path level only**. Verification earned by **passing any assessment sim/lab** in the path; passive content = evidence/progress, **not** verification. Tags carried during migration; new content tagged via a dedicated AI-assisted tool. See [`content-model.md`](content-model.md). |
| 5 | Assignments — managers giving paths to people, with deadlines | ⏭️ | **Not in R2 → R4.** New academy paths (R2–R3) are self-serve until R4; legacy stays assignable until it migrates at R5. Assignment (R4) ships before migration (R5), so **nothing is lost at sunset**. Pulls dashboards (Point 6) with it. See [`next-release.md`](next-release.md). |
| 6 | Manager dashboards / org progress views for academy paths | ⏭️ | **Deferred with assignment (Point 5)** — these views are built around assignment/org tracking. Academy = self-serve only this release. See [`next-release.md`](next-release.md). |
| 7 | Rollout — when we stop creating new legacy paths **+ one library vs two** | ✅ | Sequenced as a **5-release roadmap**: R1 skillpath→app · **R2 = this spec (engine parity)** · R3 Studio authoring · R4 assignment · R5 final migration. Two libraries coexist until **R5** (catalog then new-only). See [`vision.md`](vision.md) → Release roadmap. |
| 8 | Moving people's existing progress/history, not just the content | ✅ | **Migrate** in-flight + **preserve** completion history — both with strict consistency checks (no data loss). Verified skills already in `app` (ride along). **Grandfather** legacy completions (completed-before-migration stays completed, not re-gated). **Retro-issue** new certificates. R5. See [`content-model.md`](content-model.md). |
| 9 | Letting customers build their own paths (the authoring tool) | ⏭️ | Deferred to **R3** — update Studio Desk (today legacy-only) to create new skill paths. Not in R2 (this spec). See [`next-release.md`](next-release.md). |
| 10 | Where people actually take the courses (inside Anthropos vs. the separate academy site) | ✅ | **Standalone for now** (deep-link), access already subscription-gated (B2C individual + B2B enterprise). **No iframe.** Academy is **migrated fully in-app into `next-web-app` at R5** (native). See [`vision.md`](vision.md) → Release roadmap. |
| 11 | Skills/competency data model (simpler than legacy) | ✅ | Tagged at **path level only**; general content has none. **Exception:** sims & AI labs carry their **own** skills, surfaced at path level. Verification = passing a sim/lab verifies *its own* skills. See [`spec.md`](spec.md) §5. |
| 12 | Author vs. curator | ✅ | New model has **only author**; legacy's author/curator split is **dropped at migration** (curators → authors). See [`spec.md`](spec.md) §8. |
| 13 | Difficulty levels | ✅ | Drop legacy levels; use **Foundation → Practitioner → Advanced**. Legacy→new mapping TBD. See [`spec.md`](spec.md) §3.6. |
| 14 | Versioning & deprecation | ✅ | Update = **deprecate + replace** (not in-place), **new paths only** (no legacy support); edit-in-place vs deprecate is the author's call. Deprecated = unlisted/unstartable but open to those who started; replacement offers **continue-or-switch**; **carry-forward** reuses unchanged chapters + passed sims/labs (needs stable IDs). **Manual in R2** → **Studio UI at R3**. See [`spec.md`](spec.md) §10. |

---

## Decision log

### Points 1 & 2 — Content model (decided 2026-06-25)
Adopted **Option D**: a unified content model `Skill Path › Chapter › Module › Lecture`, where the **module**
is the unit of content a learner sits for (~20–40 min) and the **lecture** is the bite-size atom (~3–5 min)
where progress is tracked. Long/legacy content stays whole and is auto-sliced (~5 min/lecture, human/AI can
curate) into lecture-checkpoints, so it *feels* bite-size without re-authoring. Two module families:
**Authored** {ucourse, simulation, AI lab} and **Resource** {video, PDF, podcast, external link}. **Udemy
dropped.** Completion tracked via real progress where possible, click-to-complete fallback. A legacy *step*
maps to a *module*. Full write-up: [`content-model.md`](content-model.md).

### Point 3 — Simulations (decided 2026-06-25)
**Reuse** the existing simulation engine (don't rebuild). A simulation module **launches the real platform sim
and returns** (Option A; hand-off through R4, in-context once the academy is in-app at R5). Completion flows back **server-side /
event-driven** (Option 1). **All-in:** simulation support on the academy stack is required core parity — we do
not move new-path creation to academy until it works (links to Point 7). **Chapter-closure gating:** an
**assessment** simulation or an **AI lab** in a chapter must be **passed** ("completed with success"), not just
attempted, for the chapter to close — replacing legacy's per-path pass-all toggle. Training (practice) sims and
other modules need only completion. **AI labs** are a new asset type, treated like simulations, present in
neither legacy nor new paths today. Details in [`content-model.md`](content-model.md) → *Simulations & AI labs*.

### Point 4 — Skills & verification (decided 2026-06-25)
In scope. **Skills are tagged at the path level only** (no chapter/module tags — coarser than legacy, simpler).
**Verification is earned only by passing assessment simulations and AI labs** (any/all in the path, not just a
capstone); **passive content — watching/reading — does NOT verify**, it counts as progress/evidence. Skill tags
are **carried/managed during migration** (legacy already has them); **new content** gets a **dedicated
AI-assisted tagging tool**. Reuses the existing skill/competency engine, fed from academy completions & passes.
Details in [`content-model.md`](content-model.md) → *Skills & verification*.

### Point 5 — Assignments (deferred 2026-06-25)
**Not in this release** — delegated to a future one. Assignment remains available on **legacy paths only**, and
only **until those paths are migrated**. New & migrated academy paths are **self-serve** (no manager assignment,
no deadlines) until the future assignment work ships. This also defers the **assignment-driven manager
dashboards** (Point 6). **Open knock-on:** migrating an assignment-dependent path loses its assignability until
that future work — so either hold migration of such paths, or accept temporary loss (decide in Points 7–8).
Parked in [`next-release.md`](next-release.md).

### Point 6 — Manager dashboards / org views (deferred 2026-06-25)
**Deferred with assignment.** The team-progress / completion / insights views for academy paths are built around
assignment + org tracking, which isn't in the parity release. Academy paths are **self-serve only** this
release; org-level reporting ships with the future assignment work. Parked in
[`next-release.md`](next-release.md). *(The tracker's original Point 6 — "one library vs two during the
transition" — was folded into Point 7, the rollout.)*

### Point 7 — Rollout / release roadmap (decided 2026-06-25)
Convergence is sequenced as a **5-release roadmap** (this spec = **R2**):
- **R1** — legacy `skillpath` microservice merged into `app` (backend consolidation).
- **R2 (this spec)** — new skill paths (`ant-academy` + `app`) reach **engine feature parity** (content types,
  simulations, verification). No authoring tool, no assignment, no dashboards, no content migration.
- **R3** — Studio Desk updated to also author **new** skill paths.
- **R4** — assignment + manager dashboards for new skill paths.
- **R5** — final migration: all legacy paths move to the new engine; catalog shows **only** new paths; legacy retired.

Two libraries **coexist** through R2–R4 (legacy stays the only *assignable* option until R4/R5); **R5** is the
single cutover. This resolves the earlier hard-stop-vs-escape-hatch tension by *sequencing*: legacy remains
usable until its replacement (incl. assignment) exists. **Point 9 (authoring tool)** is therefore deferred to
**R3**. Full roadmap in [`vision.md`](vision.md) → *Release roadmap*.

### Point 8 — Progress & history migration (decided 2026-06-25 · executes R5)
**Data integrity is paramount — no loss, every migration validated/double-checked.** In-flight progress is
**migrated** at the module level (legacy step done → new module done); completion history is **preserved**;
verified skills **ride along** (already in `app`). **Grandfather legacy sessions (not content):** a chapter/path
completed before migration **stays completed** and is **not** re-gated against the new model. **Certificates are
retro-issued** for already-completed legacy paths. Details in [`content-model.md`](content-model.md) → *User
progress & history migration*.

### Point 10 — Consumption surface (decided 2026-06-26)
**For now: standalone.** New paths are consumed on the standalone academy site (deep-linked from Workforce), as
today. Correction to an earlier assumption: academy access is **already subscription-regulated** — B2C
individual *and* B2B enterprise users — so it's not internal-only. **No iframe.** The academy is **migrated
fully in-app into `next-web-app` (native) at R5**, the final migration release. This also finalizes Point 3's
sim launch: a **hand-off** to the platform sim while academy is standalone (through R4), becoming **in-context**
once academy is in-app at R5.

### Review round — follow-ups (decided 2026-06-26)
Closed four open items surfaced in the spec review:
- **Content levels confirmed** — `path › chapter › module › lecture`; **lecture = level 4**, kept distinct from
  module. A lecture's UX depends on the module type.
- **Languages (i18n)** — multi-language with **English fallback**; some legacy & new paths won't have every
  language. **No new translations at migration** (existing carry over); **new paths are created with full
  language coverage**.
- **Private paths & entitlements** — private/org paths exist in **both** legacy and new; keep consistent and
  **never leak private content during migration**. Entitlement model reinstated: user tiers **A/F/X/S/E** ×
  content **FREE/PAID/PRIVATE(org)** with the access matrix (FREE → all but A; PAID → all but A & X; PRIVATE →
  org members only). **F = free-content-only (same as X)**; `PAID` needs an active subscription (**S**) or org
  membership (**E**).
- **Authoring** — **curator = author** (simplified); today's Studio makes legacy paths, the **new Studio (R3)**
  makes new-style paths.
- **Assessment modality confirmed** — the sim engine already exposes multiple modalities incl. **assessment**,
  which already gates path completion (settles the §3.5 / §4 assumption).

All folded into [`spec.md`](spec.md). Also fixed in this round: the stale assignment-migration "tension" (now
resolved by the R4-before-R5 ordering) across the tracker + `next-release.md`.

### Round 2 — additional spec items (decided 2026-06-26)
Four more items folded into [`spec.md`](spec.md):
- **Skills/competency data model** — simpler than legacy: **path-level tagging only** for general content;
  **sims & AI labs carry their own skills** (the one exception), surfaced at path level; passing a sim/lab
  verifies *its own* skills. (Path-declared skills not covered by any sim/lab stay evidence-only.) §5.
- **Author vs. curator** — collapsed to a single **author**; legacy's distinction dropped at migration. §8.
- **Difficulty** — **Foundation → Practitioner → Advanced**, replacing legacy levels (legacy→new mapping TBD). §3.6.
- **Versioning & deprecation** — updates are **deprecate + replace** (supersedes legacy's in-place upgrade),
  **new paths only** (legacy paths just migrate). Deprecated paths are unlisted/unstartable but stay open to
  those who started them; a replacement offers **continue-or-switch** from the landing page; **carry-forward**
  reuses unchanged chapters + already-passed sims/labs (requires stable IDs). Edit-in-place vs deprecate+replace
  is the **author's manual call**. **Manual in R2 (backend) → Studio authoring UI at R3.** Confirmed: content
  progress is **chapter-keyed with lecture-level granularity**, sims/labs by pass (§3.7). §10.
