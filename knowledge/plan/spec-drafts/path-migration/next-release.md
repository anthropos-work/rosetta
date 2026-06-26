# Skill Paths — Next Release (parking lot)

> **Status:** Draft · spec-draft · 2026-06-25
> Companion to [`vision.md`](vision.md).

## Purpose

Anything that is **out of scope for the feature-parity release (R2)** lands here: future, deprioritized, or
not-urgent development and features we explicitly **do not** want to ship as part of R2. Items are mapped to
their **target release** (R3 / R4 / R5) — see the full sequence in [`vision.md`](vision.md) → *Release roadmap*.

This is the parking lot. When something comes up that we agree is "later, not now," add a row below so it's
captured and we keep the parity project focused.

## How to use

- One row per item. Keep it short.
- **Area** = which flavour / surface it touches (legacy, new/ai-academy, app, next-web-app, shared).
- **Why deferred** = the reason it's not in the parity scope.

## Parked items

| # | Item | Area | Why deferred |
|---|------|------|--------------|
| 1 | **Assignment for academy paths** — manager → person/team, deadlines, per-person status | new / academy (`app` + `next-web-app`) | Not in R2. New academy paths (created R2–R3) are **self-serve** until R4; legacy paths stay assignable (on legacy) until they migrate at R5. Since assignment ships at **R4 (before migration at R5)**, migrated paths are assignable — **nothing is lost at sunset**. **Target: R4.** |
| 2 | **Manager dashboards / org progress views for academy paths** — team progress, completion, insights | new / academy (`app` + `next-web-app`) | Deferred with assignment (#1) — these views are built around assignment + org tracking. Academy paths are self-serve only in the parity release; org-level reporting ships with the future assignment release. **Target: R4.** |
| 3 | **Studio Desk authoring of new skill paths** — update the builder (today legacy-only) to create new-engine paths (incl. the UI for **deprecation / replacement**, which is manual via the backend until then) | `studio-desk` | R2 (this spec) reaches *engine* parity; customer-facing authoring of **new** paths comes with the Studio update. **Target: R3.** |

> **Resolved by ordering:** assignment (R4) ships *before* the final migration (R5), so migrated paths are
> assignable and **nothing is lost at sunset**. The only interim limitation is that **new** academy paths
> created in R2–R3 are self-serve (not assignable) until R4.

<!-- Add new rows above. Move an item into the parity scope only by promoting it into vision.md / a spec doc. -->
