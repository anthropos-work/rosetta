# M256 · iter-04 — progress

**Type:** tik · **Active strategy:** `TOK-01` move 3 (org-admin before onboarding).
**Phase 0d: PASS** — `ptvalidate` accepted the pre-existing tree (`8 product(s), 18 use case(s), 18 live, 0 TODO`, exit 0).

## Phase A — probe the four surfaces before declaring anything

All four routes answer **200** and admit the org-admin manager, and **all four have a real read-back
surface** — the milestone's open question, answered (**D13**). The probe also produced two corrections to
artifacts written earlier in this same iter (**D14**): the tags surface has **14** tags (not 0) and renders
**cards on a two-`<main>` route**, so a `main()`-scoped table count had read an empty region.

## Phase B — declared (manifest + seed in lockstep, the M219 rule)

`manifest/org-admin.yaml` — the **9th product**, 4 stories, 4 use cases. Four `capabilities:` added to
`seed/seed-worlds.yaml` in the **same change** (`org-role-catalog`, `org-tag-surface`, `org-member-roster`,
`org-feature-settings`), each annotated with what the seed actually writes and what breaks in its absence.

`ptvalidate`: **`9 product(s), 22 use case(s), 20 live Playthrough(s), 2 TODO`** — valid.

## Phase C — built

- `lib/org-admin-page.ts` — `OrgRolesPage`, `OrgTagsPage` (page-scoped, see D14), `OrgSettingsPage`,
  `MembersBulkActionsPage`, plus a shared **`clickUntilDialog`** helper that generalises iter-03 D10's
  hydration-race fix. Deliberately **not** a subclass of `members-page.ts`: read and write are different
  surfaces of one route, and sharing a class would couple two products' re-pins.
- `ROLES_URL` / `TAGS_URL` / `ORG_SETTINGS_URL` + `isOnRoles` / `isOnTags` / `isOnOrgSettings` in
  `url-shapes.ts`, under the same segment-anchored discipline (a bare `\b` would false-match
  `/enterprise/roles-import`).
- **4 specs written and run live.** Two land; two are parked with diagnoses (**D15**).

### Landed GREEN — both mutate state AND read it back

| Playthrough | Write | Read-back (the proof) |
|---|---|---|
| `pt-orgadmin-tag-create` | create a team/tag | the name is **absent before**, present after, and **still present after a full reload** → it came from the backend, and a unique name cannot be pre-existing data |
| `pt-orgadmin-setting-toggle` | flip an org feature setting | `aria-checked` reads back **flipped after a full reload** → persisted server-side, not local component state. The switch chosen ("Disable Completion Emails") is the one of four with **no** blast radius on other Playthroughs |

### Parked as declared `TODO`, specs in `e2e/drafts/` (zero standing red — D-v28-3)

| Use case | Handler | Blocking finding |
|---|---|---|
| `org-admin.roles.UC1` | `PT-M256-orgadmin-role-create` | the platform's create-role **`Save` appears to no-op**: enabled, clicked, and **no HTTP ≥ 400 / no console error / no navigation / no new row**, dialog stays open, `alert` region **EMPTY**. Possibly a real product defect (silent validation). Untried: the "Suggest skills" prefill |
| `org-admin.members.UC1` | `PT-M256-orgadmin-member-tag` | the bulk-action **dropdown intercepted pointer events** over its own modal — **454** click retries against `.ant-dropdown-menu-item` to the 240 s budget. **Invisible to actionability checks.** `Escape` fixed that leg; `check()` then still timed out |

## Phase D — re-measure on the grown denominator (D7's protocol, n=3)

| Figure | Baseline | Now | Ratio |
|---|---:|---:|---:|
| **Median per non-studio Playthrough (post-coverage, 18)** | 3.326 s | **1.808 s** | **0.5434×** |
| **Same metric, ORIGINAL 16 only — the honesty cross-check** | 3.326 s | **1.778 s** | **0.5347×** |
| Suite wall-clock (REPORTED, not gated) | 56.6 s | **44.5 s** | — |
| Flake over 3 consecutive runs | 0 | **0** (`136 passed` ×3) | — |
| Coverage (`ptreport`) | 18/18 | **20/22 passing (90.9 %)**, 2 `[TODO]`, 0 `failing`, 0 `unimplementable` | — |

The cross-check matters: the two figures agree, and the new Playthroughs are **slower** than the median
(3.44 s / 2.51 s), so the gain is in the existing tests rather than an artifact of adding fast ones (**D16**).

## Close — 2026-07-28

**Outcome:** the org-admin product exists — 9th product, 4 declared use cases (from **0 of 4** covered for
five releases), **2 landed GREEN as real WRITE-and-read-back Playthroughs**, 2 declared `TODO` with written
diagnoses and named handlers. The milestone's open question is answered: **all four surfaces have real
read-back surfaces**, none is `unimplementable`. Clause 1 re-verified on the grown denominator at **0.5434×**
with an honesty cross-check at 0.5347×. Clause 2's mutating count **1 → 3** of the required 5.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET — clause 1 is met and now re-verified post-coverage (0.5434×, 0 flake); **clause 2** needs 2 more mutating Playthroughs (3 of 5), a negative control per Playthrough, and ≥ 1 `blocked` outcome; **clause 3** needs onboarding ×5 plus the 2 parked org-admin UCs, and the written verdicts; **D-v28-5** unstarted.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (2 of 4 landed and **no** surface proved unimplementable, so the `> 3 unimplementable` trigger is nowhere near firing) — (4) user-blocker: n (the two reds were inside the iter's planned scope; they are parked as declared `TODO` with diagnoses, so **no red accumulates** and D-v28-3's escalation condition — a non-empty red set at batch end — is **not** met) — (5) cap-reached: n (3 of 5 tiks) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D13 (all four surfaces have read-back surfaces — the open question answered), D14 (two probe-driven corrections: the 14-tag count and the two-`<main>` hazard), D15 (`TODO` + `e2e/drafts/` rather than standing red, with both diagnoses), D16 (clause 1 re-verified post-coverage with the original-16 cross-check), D17 (clause 2's mutating count is 3 of 5 and the parked journeys are gate-critical).
**Side-deliverables:** none.
**Routes carried forward:**
- `PT-M256-orgadmin-role-create` → **Fate 3, iter-05.** Gate-critical for clause 2. Try the "Suggest skills"
  path; if `Save` really is a silent platform no-op, that is a **finding to report**, and clause 2's count
  must come from another write surface.
- `PT-M256-orgadmin-member-tag` → **Fate 3, iter-05.** Gate-critical for clause 2. The interception is fixed;
  the tag-list `check()` needs the modal's filter box or its scroll container.
- The four pre-existing routes in `../progress.md` § Next-iter routing still stand.
**Lessons:**
1. **Probe the surface before you declare the precondition.** This iter wrote a seed capability claiming
   "no tags seeded" and built a read-back on it — from a locator that was measuring the wrong `<main>`. The
   correction produced a *better* assertion (a unique name surviving a reload) than the original. Two iters
   after the Phase-0b audit made exactly this point about `entitlement`, the same class of error recurred,
   from a locator instead of a YAML file. **Read the surface, not the artifact describing it.**
2. **A hit-target interception is invisible to actionability checks.** Playwright called the element
   "visible, enabled and stable" and retried 454 times against an overlay. When a click retries into a
   timeout, look for what is *on top*, not for a longer wait.
3. **`TODO` is a first-class outcome and standing red is not.** Two written, run, diagnosed specs are worth
   more parked-and-declared than deleted, and worth more than left red. The tooling already had the
   vocabulary for this (`P5` + `NoRegressions()` tolerating `TODO`); the only thing missing was somewhere for
   the file to live that neither runs it nor orphans its tag.
