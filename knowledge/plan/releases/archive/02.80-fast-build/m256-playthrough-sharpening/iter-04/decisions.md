# M256 · iter-04 — decisions

## D13 — All four org-admin surfaces have REAL read-back surfaces (the milestone's open question, answered)

`../overview.md` asked: *"Do the org-admin writes have a read-back surface, or only a toast? (A write with
no readable effect cannot be proven without a DB assert — which is a different, weaker proof shape.)"*

**Answer: all four have one.** Probed live on `demo-2` as `pt-manager` (who *is* the org admin — the same
seat that clears studio-desk's `checkEnterpriseAndAdmin` gate) before a line of manifest was written:

| Route | HTTP | Landmark | Write affordance | Read-back |
|---|---:|---|---|---|
| `/enterprise/roles` | 200 | h1 "Roles", 20 rows | `[New Role]` → dialog "Create Role" (Job Title + Job Description, `[Save]`) | the catalog list re-reads |
| `/enterprise/tags` | 200 | h1 "Tag Management" | `[Create Tag]` → dialog (1 textbox, `[Create]`) | the tag card list re-reads |
| `/enterprise/members` | 200 | h1 "Members", 20 rows | selecting a row reveals `[Choose action]` → `{ Assign Tags, Remove Tags, Delete }` | the tag's member tally, cross-surface |
| `/enterprise/settings` | 200 | h1 "\<Org\> Settings" | **four** org feature switches | `aria-checked` after a full reload |

No DB assert is needed anywhere, so the weaker proof shape is avoided. **No surface is
`unimplementable-without-platform-edit`**, and the milestone's re-scope trigger (`> 3` unimplementable) is
nowhere near firing.

## D14 — Two probe-driven CORRECTIONS to artifacts written earlier in this same iter

Both are recorded rather than silently fixed, because both are instances of patterns the milestone is
supposed to be teaching.

**(a) `seed-worlds.yaml`'s `org-tag-surface` claimed the opposite of the truth.** Its first draft — written
from a probe that counted `main().locator('table tbody tr')` on `/enterprise/tags` and got **0** — declared
*"Org A has NO org tags seeded"*, and `pt-orgadmin-tag-create`'s read-back was built on a **0 → 1** delta.
Org A has **14** tags. The `0` was the **two-`<main>`** hazard: the route renders two `<main>` elements and
the content is in the **second**, so `main()` (which is `.first()`) measured an empty region while the surface
plainly read *"Tag Management · 14 Tags · 47 Members"*. This is the **M204 iter-03 D2** finding
(`playthroughs.md` § "`main()` is not universal"), hit again on a new surface — evidence that the doctrine
needs to be *checked per surface*, not assumed absorbed.

**The corrected read-back is stronger than the one it replaced:** a distinctive name that is absent before,
present after, and **still present after a full reload**. A count delta could in principle be satisfied by
an unrelated concurrent write; a unique name surviving a reload cannot.

**(b) `main()` is unusable on `/enterprise/tags` at all**, so `OrgTagsPage` is **page-scoped by design**, with
the reason in its class header so a future reader does not "fix" it back to `main()`.

## D15 — Two of the four are parked as declared `TODO` with diagnoses, not left as standing red

The scope-creep tripwire fired: three distinct lines of investigation inside one iter (the roles `Save`
no-op, the member-tag pointer interception, the member-tag checkbox). Per the tripwire — **land what is
complete, route the rest with named handlers, close on planned-scope outcome.**

**Why `TODO` and not standing red.** **D-v28-3** mandates *zero standing red; nothing accumulates across
runs*. A red spec that stays red every run is exactly what it forbids. **P5** provides the sanctioned
alternative: a use case may be declared *before* its Playthrough exists (`playthrough: TODO`), and
`Report.NoRegressions()` — the gate a coverage milestone runs — deliberately tolerates `TODO` while failing
on `failing`. So both UCs are **declared** and appear as `[TODO]` in the four-state map: visible, counted,
and impossible to mistake for a silent gap.

**Why the specs are kept, in `e2e/drafts/`.** Deleting run-and-diagnosed work would throw away the expensive
part. But a `@pt:`-tagged spec left inside `tests/` would be an **orphan test** under the validator's
both-way integrity check (direction (b)) *and* would run and go red. So the drafts sit outside both
`playwright.config.ts`'s `testDir` and `ptvalidate --e2e-dir`, with a `README.md` stating the contract:
a draft has a Fate-3 handler and a written blocking finding, and it returns to `tests/` in the **same
commit** that re-points its use case.

**The two diagnoses, so iter-05 starts from evidence rather than a re-derivation:**

- **`PT-M256-orgadmin-role-create`** — *the platform's create-role `Save` appears to NO-OP.* With both fields
  filled, `Save` becomes **enabled**; clicking it produces **no HTTP ≥ 400, no console error, no navigation,
  no new catalog row**, the dialog **stays open**, and an `alert` region is present but **EMPTY**. Clicking
  "Start from scratch" first changes nothing. The hypothesis to test is that the mandatory *"Core skills"*
  step fails validation **silently** — which, if true, is a **real product defect** (a user clicks Save and
  nothing whatsoever happens), and therefore a genuine find, not a test bug. The untried path is the
  **"Suggest skills"** prefill affordance.
- **`PT-M256-orgadmin-member-tag`** — *an antd interaction inside the assign-tags modal.* The first failure
  was the bulk-action **dropdown intercepting pointer events** over the modal it had just opened: Playwright
  reported the target *"visible, enabled and stable"* and then retried the click **454 times** against
  `.ant-dropdown-menu-item` until the 240 s budget expired. **A hit-target interception is invisible to
  actionability checks**, so waiting longer can never fix it — the overlay has to go. Dismissing the menu with
  `Escape` fixed that leg; `checkbox.check()` then still timed out, so the long tag list (its own scroll
  container + a filter box) needs a different approach.

## D16 — Clause 1 re-verified on the GROWN denominator, with an honesty cross-check

| Figure | Baseline (iter-02) | Now (iter-04) | Ratio |
|---|---:|---:|---:|
| **Median per non-studio Playthrough — post-coverage, 18** | 3.326 s | **1.808 s** | **0.5434×** |
| **Same metric over the ORIGINAL 16 only** | 3.326 s | **1.778 s** | **0.5347×** |
| Suite wall-clock (REPORTED) | 56.6 s | 44.5 s | — |
| Flake over 3 consecutive runs | 0 | **0** (`136 passed` ×3) | — |

**The cross-check is the point.** A median measured on a suite that just gained tests can fall simply because
the new tests are fast — which would be gaming the gate, not passing it. Computing the same statistic over the
**original 16** isolates that: **0.5347× vs 0.5434×**, i.e. the improvement is in the *existing* tests and the
new ones are, if anything, **slower than the median** (`pt-orgadmin-tag-create` 3.44 s,
`pt-orgadmin-setting-toggle` 2.51 s vs a 1.81 s median). The two figures will be reported together at the
final measure for the same reason.

## D17 — Clause 2's mutating count is now 3 of 5; the two parked journeys are its remaining path

Mutating (mutates state **and** reads it back): `pt-assignment-assign`, `pt-orgadmin-tag-create`,
`pt-orgadmin-setting-toggle` = **3**. Clause 2 needs **≥ 5**. The two parked org-admin journeys are the
natural remaining two, which makes `PT-M256-orgadmin-role-create` and `PT-M256-orgadmin-member-tag`
**gate-critical**, not merely tidy-up. If the roles `Save` really is a platform-side silent no-op, the
count must come from elsewhere — the honest options being `org-admin.members.UC1` plus one more write
surface, and that decision belongs to the tik that has the evidence.
