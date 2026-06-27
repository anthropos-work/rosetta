# Skill Paths — Feature-Parity Spec (R2)

> **Status:** Consolidated draft `v1.0` · spec-draft · 2026-06-26
> **Companions:** [`spec-progress.md`](spec-progress.md) (decision tracker + log) · [`next-release.md`](next-release.md) (deferred items) · [`examples/`](examples/) (HTML mocks: legacy / new / migrated)
> *This document consolidates and supersedes the earlier split `vision.md` + `content-model.md` drafts.*

---

## 1. Overview

### 1.1 North star — product & experience
- **Who it's for:** knowledge workers / professionals in enterprise organizations who need to upskill or
  improve competencies on the on-the-job processes and work they actually do.
- **How it feels:** **bite-size learning.** A skill path may cover *hours* of content, but it's built to be
  absorbed in **pills of ~3–5 minutes** that fit a busy weekly schedule.

### 1.2 The goal
We ship skill paths in **two** flavours today:
- **Legacy** — on `next-web-app`, `cms`, Directus, and `app`.
- **New (AI Academy)** — on `ant-academy` and `app`.

The goal is to **completely replace legacy** with the new engine, **migrate** all legacy content onto it, and
**sunset legacy** — **losing no functionality** (video, simulations, links, assignment, private paths,
languages, entitlements).

### 1.3 Scope of this spec
**This spec is `R2`** in the roadmap below: the new engine reaches **engine / content feature parity** — it can
represent and run everything legacy can (all content types, simulations, skill verification, languages, private
paths, entitlements). It **does not** include the authoring tool (R3), assignment + dashboards (R4), or the
content/user migration (R5).

---

## 2. Release roadmap

| Release | What it delivers |
|---|---|
| **R1** | Backend consolidation — the legacy `skillpath` microservice is **merged into `app`**. |
| **R2 — this spec** | **Engine feature parity** — new skill paths (`ant-academy` + `app`) can do everything legacy can. *No authoring tool, no assignment, no dashboards, no content migration.* |
| **R3** | **Authoring** — Studio (Desk) updated to also create **new**-style skill paths (today it only makes legacy ones). |
| **R4** | **Assignment** — assignment + manager dashboards enabled for new skill paths. |
| **R5** | **Final migration** — all legacy paths move to the new engine; catalog shows **only** new paths; the academy is brought **fully in-app** (native in `next-web-app`, **no iframe**); legacy retired. |

**Coexistence.** Through R2–R4 both libraries coexist; new paths are consumed on the **standalone** academy site
(subscription-gated). Legacy stays usable — and stays **assignable** (on legacy) — until it migrates at R5.
Because assignment ships at **R4**, *before* migration at **R5**, migrated paths are already assignable —
**nothing is lost at sunset.** The only interim limitation: new academy paths created in **R2–R3** are
**self-serve** (not assignable) until R4.

---

## 3. Content model

### 3.1 The four levels

```
Skill Path                              (1) the program
└─ Chapter                              (2) a unit with its own metadata; a collection of modules
   └─ Module                            (3) THE unit a learner sits down for (~20–40 min); a collection of lectures
      └─ Lecture                        (4) the bite-size atom (~3–5 min); where progress is tracked
```

- **Lecture is level 4** — the bite-size unit. It is **internal** to a module (the module's sequencing), not a
  separately-browsable catalog entity.
- **A lecture renders differently depending on the module type** (an authored screen, a slice of a video, a
  range of a PDF, …) — see §3.2.
- **UX rule:** *module = the thing you take; lecture = the bite-size unit you progress through.*

### 3.2 Module families & types

Every module is one of two families:

**Authored** — Anthropos-built interactive learning:
- **ucourse micro-lesson** — lectures are the authored screens (concept / analogy / example / practice quiz).
- **simulation** — interactive, graded roleplay. **Atomic** (not sliced into lectures); completion = the sim's
  own result. See §4.
- **AI lab** — interactive Anthropos lab. **New asset type** (not in legacy or new paths today); treated like a
  simulation. See §4.

**Resource** — curated media wrapped into the model:
- **video** — lectures = time-slices (a 40-min video → ~8 × 5-min checkpoints).
- **document (PDF)** — lectures = page / section ranges.
- **podcast / audio** — lectures = time-slices.
- **external link** — single checkpoint (off-site → can't track inside): "open + mark done."
- ~~**Udemy**~~ — **dropped**; no longer supported.

### 3.3 How lectures are sliced
- **Default: auto** — at migration, and when a new asset is loaded, the system auto-slices into ~5-min-worth
  lectures.
- **Override: human or AI** can curate the cut points.
- *(Simulation / AI-lab modules are atomic — they have no lectures.)*

### 3.4 What "lecture complete" means
- **Track real progress where possible** (JS on the module — video actually watched, document scrolled).
- **Fallback: click-to-complete** when real tracking isn't possible (e.g. an external link).
- Legacy only ever did click-to-complete, so this is an upgrade.

### 3.5 Chapter-closure gating — "pass, not just done"
A module can be merely **completed** (attempted/done) or **completed with success** (passed). The rule:

> A **chapter is completed only when ALL its modules are completed AND every AI lab and assessment simulation in
> it is passed.**

- A chapter with **no** lab/assessment-sim → completed by completing its content (modules).
- A chapter **with** 1+ lab/assessment-sim → also requires **passing** every one of them.
- Training (practice) simulations and all other module types need only completion.

This replaces legacy's per-path `require_all_ai_sim_pass_to_complete_skill_path` toggle with one consistent rule.

### 3.6 Difficulty
Paths use a single three-level scheme — **Foundation → Practitioner → Advanced** — replacing legacy's levels
(which we're dropping). Legacy levels map onto these three at migration (mapping TBD — §12).

### 3.7 How progress is stored (and why version carry-forward works)
- **Content progress is stored per chapter** — each chapter record holds the completed **modules / lectures**
  within it. So progress is *keyed* at the chapter level, with lecture-level granularity inside.
- **Sims & AI labs are tracked separately, by pass** (their own completion signal from the engine).

This split is what makes **version carry-forward** possible (§10): an unchanged chapter keeps its progress, and
an already-passed sim/lab needn't be redone.

---

## 4. Simulations & AI labs

- **Reuse the existing simulation engine** — don't rebuild. A simulation module *points at* a simulation and
  reacts to its result.
- **Launch the real experience, then return** — the module hands off to the existing platform simulation UI and
  comes back. *While the academy is standalone (through R4) this is a cross-surface hand-off; once the academy
  is in-app at R5 the sim launches in-context.*
- **Completion is server-side / event-driven** — when the engine reports the sim **passed** for the learner,
  the module is marked accordingly (the authoritative signal legacy already uses).
- **AI labs** behave the same way (launch → completion signal back). They are a **new asset type** to support
  like simulations; present in neither legacy nor new paths today.
- **Timing: all-in** — simulation support is **required core parity**; the engine isn't "at parity" until sims
  work.
- **Assessment vs. training** — the simulation engine already exposes multiple **modalities**, including
  **assessment** (which already gates path completion). Only **assessment** sims (and AI labs) gate chapter
  closure (§3.5) and earn verification (§5); training/practice modalities need only completion.

---

## 5. Skills & verification

- **Deliberately simpler than legacy.** Legacy tagged skills at a granular chapter/step level — too complex.
  The new model tags **skills at the PATH level only**: a path declares the taxonomy skills it develops, and
  general content (ucourse lessons, videos, PDFs, podcasts, links) carries **no** skill tags of its own.
- **Exception — sims & AI labs carry their own skills.** Interactive assets (AI simulations, AI labs) come with
  their **own** associated skills (inherent to the asset). Those skills are **also surfaced at the path level**,
  so a path's displayed skill set = its path-declared skills **+** the skills of any sim/lab modules it contains.
- **Verification is earned by passing assessment sims & AI labs** — passing one verifies **its own** skills
  (which appear at path level). *(Any/all of them count — not just a capstone.)* It feeds the learner's
  competency / Skill Spotlight. (Path-declared skills **not** covered by any sim/lab stay *evidence only* —
  see the next bullet.)
- **Passive content does NOT verify.** Watching/reading/finishing lectures is **progress & evidence**, not
  verification (verification = you *proved* it, via a sim/lab). **Consequence:** a path with **no** assessment
  sim/lab can be *completed* but its skills are **never verified** (evidence only).
- **Tagging:**
  - *Migration:* skill tags are **carried over / managed during migration** (legacy paths already carry taxonomy
    skills — largely free).
  - *New content:* a **dedicated AI-assisted tagging tool** suggests taxonomy skills from the content for a
    human to confirm.
- Reuses the existing skill / competency engine, now fed from academy completions & sim/lab passes too.

---

## 6. Languages (i18n)

- **Multi-language with English fallback** — content is authored per language; a missing language falls back to
  English. (Both legacy and the academy already work this way.)
- **Some legacy *and* new paths may not have all languages** — that's fine; they fall back to English.
- **Migration introduces no new translations** — existing translations carry over as-is; gaps stay as English
  fallback.
- **New paths going forward are expected to be created with full coverage** for all supported languages.

---

## 7. Private paths & entitlements

### 7.1 Private (org) paths
- **Private, org-scoped paths exist in BOTH legacy and new** — keep them **consistent** across the two.
- **Migration must never leak private content** — a private path of org *X* stays private to org *X*; it is
  never exposed publicly or to another org during/after migration. (See §9.)

### 7.2 Entitlement model (to reinstate in the new engine)

**User-access tiers**
| Tier | Meaning |
|---|---|
| **A** | anonymous (not logged in) |
| **F** | free (logged-in, not paying) — **free content only** |
| **X** | subscription-**expired** logged-in user — free content only (**same access as F**) |
| **S** | subscribed / paying logged-in user |
| **E** | enterprise user (member of an org) |

**Content-access tiers:** `FREE` (free content) · `PAID` (paid content) · `PRIVATE (org)` (private to an org).

**Who can access what**
| Content ↓ \ User → | A (anon) | F (free) | X (expired) | S (paying) | E (enterprise) |
|---|:--:|:--:|:--:|:--:|:--:|
| **FREE** | ✗ | ✓ | ✓ | ✓ | ✓ |
| **PAID** | ✗ | ✗ | ✗ | ✓ | ✓ |
| **PRIVATE of org *X*** | ✗ | ✗ | ✗ | ✗ | ✓ *only if member of org X* |

- `FREE` → all but **A** (every logged-in user).
- `PAID` → only **S** and **E** (an active subscription or org membership).
- `PRIVATE of org X` → **only members of org X**.

**F and X have identical access — free content only.** A free user and a lapsed subscriber both see only
`FREE`; paid content requires being a paying subscriber (**S**) or an enterprise/org member (**E**).

---

## 8. Authoring

- **One role: author.** Legacy distinguishes **author vs. curator**; the new model has **only author**. The
  distinction is **dropped at migration** — former curators become authors too (this also covers the legacy
  "Meet the Experts" credit).
- **Today's Studio creates legacy paths; the new Studio (R3) will create new-style paths.** Customer-facing
  authoring of new paths lands at **R3** — not in R2.
- **In R2**, new academy content is authored the existing `ant-academy` way (in-repo content + authoring skills)
  until the Studio update lands at R3.

---

## 9. Migration (executes at R5)

**Data integrity is paramount — no loss, and every migration is validated / double-checked.**

### 9.1 Content mapping (legacy → new)
| Legacy | → | New |
|---|---|---|
| skill path | → | skill path |
| chapter | → | chapter |
| step = video | → | **video module** (sliced into lecture-checkpoints) |
| step = PDF / file | → | **document module** (sliced) |
| step = podcast | → | **audio module** (sliced) |
| step = web link | → | **external-link module** (single checkpoint) |
| step = Udemy | → | dropped — decide per asset at migration (convert to link, or drop) |
| step = job simulation | → | **simulation module** (atomic, graded — §4) |

**Rule of thumb:** a legacy *step* becomes a *module*; bite-size moves one level deeper into *lectures*.

### 9.2 User progress & history
- **In-flight progress: migrated** at the module level (legacy *step done → new module done*; the clean
  step→module mapping makes this safe), with a consistency check.
- **Completion history: preserved** (who completed what, and when).
- **Verified skills: ride along** — they already live in `app` (user skills / competency); no migration needed
  (post-R1, skillpath is part of `app`).
- **Grandfather legacy sessions (not content):** a chapter/path **completed before migration stays completed**
  and is **not** re-derived or re-gated against the new model (the new pass-gating never re-opens an
  already-completed legacy chapter).
- **Certificates: retro-issued** — already-completed legacy paths receive a new Certificate of Mastery.

### 9.3 Safety
- **No new translations** introduced (existing carry over; gaps fall back to English — §6).
- **No private-content leakage** — private org paths stay scoped to their org (§7.1).

---

## 10. Versioning & deprecation

How a **new** path changes over time. This **replaces legacy's in-place `upgradeSkillPathSessionToLatest`** with
a deprecate-and-replace model. **New paths only** — legacy paths get no deprecation/replacement support; they
are simply migrated at R5 (§9).

- **Updating a path = deprecate + replace.** A *significant* update does **not** mutate a path learners are
  mid-way through; instead the current path is marked **deprecated** and a **replacement** is published. Whether
  to **edit in place** (minor fix) or **deprecate + replace** (significant change) is the **author's call** — no
  system threshold.
- **Deprecated content:**
  - **Hidden from the library** — not listed, **not startable** by new people.
  - **Still accessible to people who already started it** (it stays in the data — just unlisted / unstartable).
- **Replacement & the continue-or-switch choice.** A deprecated path can point to a **replacement**. When a
  learner who already started it lands on its page, the **content landing page** offers a choice: **Continue**
  the deprecated path to completion, or **switch** to the new one. *(Offered only from the landing page;
  rendered by the consumption surface — the academy app.)*
- **Smart carry-forward on switch.** Because progress is stored per chapter and sims/labs are tracked by pass
  (§3.7), switching reuses what still matches: the learner studies only the **changed chapters** and re-does
  only the **sims/labs they haven't passed**. **Requires stable chapter / sim / lab identifiers across versions**
  (the same stable-ID basis legacy used for upgrades).
- **Rollout:** the academy/backend supports deprecation + replacement **manually in R2** (set via the backend /
  existing academy authoring — no dedicated UI). **Studio Desk adds the authoring UI at R3.**

---

## 11. Out of scope for R2 (deferred)

Tracked with target releases in [`next-release.md`](next-release.md):
- **Authoring of new paths in Studio** → **R3**.
- **Assignment** (manager → person/team, deadlines, status) → **R4**.
- **Manager dashboards / org progress views** → **R4** (built around assignment).
- **Content + user migration** and the **full in-app academy** → **R5**.

---

## 12. Open / to-confirm
- **Legacy Udemy steps** — what to do with them at migration (convert to link, or drop) (§9.1).
- **Legacy → new difficulty mapping** — how legacy levels map to Foundation / Practitioner / Advanced (§3.6).
