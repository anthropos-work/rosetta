**Type:** tik · **Active strategy:** `TOK-01` move 4 (the honesty items — negative controls)

# iter-13 — the structural finals had no contrast vantage because they were structural

## Phase A — probe first (the mandatory step, and it decided everything)

Three seats, both profile surfaces, on the live `demo-2`. **Nothing was written until this table existed.**

| fact rendered on `/profile` | `pt-employee` (Pat Ellis) | `pt-manager` (Morgan Reyes) | `pt-free` (Sam Okafor) |
|---|---:|---:|---:|
| seeded role / org | DevOps Engineer / Meridian Labs | Engineering Manager / Meridian Labs | Account Executive / Halcyon Retail |
| `Verified Skills` stat | **8** | 3 | 2 |
| `All Skills` stat | **20** | 11 | 10 |
| `Role Skills` stat | 10 | 10 | 10 |
| role-context line `"<role> at <org>"` | **1** | **0** | **0** |
| current-position conjunction (role ∧ org ∧ `- Present`) | **13** (innermost = the card) | **0** | **0** |
| `My Closest Roles` section | **1** | **0** | **0** |
| role-readiness ring | 76 % | 9 % | 6 % |
| `Skill Gaps (N)` | **(0)** | (3) | (0) |

**The seed cross-check is the finding.** The rendered `Verified Skills` stat equals the hero's seeded
`skills.verified` **exactly** (8 / — / 2), and `All Skills` equals seeded `verified + mapped` **exactly**
(8+12=20 / — / 2+8=10) — confirmed on two heroes independently. `pt-manager` has no `skills:` block at all;
her 3 / 11 come from the M44 completeness seeder, which is precisely what makes her discriminate.

**So the hypothesis held: iter-12's refutation closed a shortcut, not the path.** A *structural* predicate
(a stat LABEL is visible, a chart count ≥ 1, a "Work" section exists) really cannot be falsified by any
vantage — M44 gives every member a career and skills, so there is no hero for whom it is legitimately
absent, and no suppression switch can exist. But those finals were structural **because they were written
structurally**. The surface carries three independent hero-specific facts, and every one of them reads
**0 for the contrast seat**.

Two Phase-A side-findings, both recorded because they cost time:

- **A skills→career tab click is intercepted by a `headlessui-portal-root` overlay.** The career leg must
  precede the skills leg. The shipped timeline Playthrough already used that order by luck; the probe did not.
- **`locator.evaluate()` on an absent element waits for the TEST budget, not an action timeout.** The probe
  hung 180 s on a `.catch()`-guarded `evaluate` over an empty locator — the same unbounded-interaction class
  the harden pass fenced in the harness, reproduced in throwaway probe code within the hour.

## Phase B — the three finals, re-aimed

| Playthrough | final BEFORE (structural) | final AFTER (hero-specific) |
|---|---|---|
| `pt-profile-verified` | `roleSkillsStat` visible + a chart drew | her seeded magnitudes: `Verified Skills` = **8**, `All Skills` = **20** (= verified 8 + mapped 12 — the claimed-vs-verified gap the Playthrough is named for), plus the role-context line `"DevOps Engineer at Meridian Labs"` |
| `pt-profile-growth` | `skillGapsStat` visible + a chart drew | the **closest-role recommendation** — the roles her VERIFIED skills make her nearest to — with ≥ 1 match score. It renders only once the matcher has enough verified evidence; absent for both other seats |
| `pt-profile-timeline` | Work + Education sections + ≥ 1 dated entry | one timeline block stating she is **currently** a `<seeded role>` at `<seeded org>` (role ∧ employer ∧ still-open range) |

The old finals are all **retained as intermediates** — they establish that the surface is there; they no
longer pretend to establish whose it is.

**`lib/seed-facts.ts` + `tests/seed-facts-fence.unit.spec.ts` (net-new).** A sharpened final carries a
NUMBER, which is a claim about `seed/pt-world.seed.yaml` with no link to it: renumber the seed and three
Playthroughs go red naming a product regression that never happened. So every fact a final may name is
declared once, and the fence PARSES the seed and reconciles them. `allSkillsTotal()` **throws** for a hero
with no seeded skills block rather than returning `NaN`, and the fence asserts the contrast hero *has* no
such block — because if she ever gained one her magnitudes could coincide with the hero's and the control
would go quietly vacuous.

## Phase B/C — the bug this iter shipped and then caught: `\b` in a `hasText` regex

Both new conjunction locators read **0 on a page that plainly rendered the thing**. Cause: `hasText`
matches against `textContent`, which concatenates sibling text nodes with **no separator**. The work card's
textContent is `…Meridian LabsFeb 2024 - Present (2 years)…`, so a `\b` before `\w{3}` has no boundary to
find ("s" then "F" are both word characters) and the pattern matches nothing. Same for `^My Closest Roles\b`
("Roles" then "Based").

`TIMELINE_DATE_ENTRY` keeps its `\b` safely only because it is consumed through `getByText`, which resolves
to leaf-ish elements whose text is not a concatenation. **Same regex, different consumer, different rules.**
Fixed, documented at both constants, and pinned by a unit test that asserts the *concatenated* shape and
asserts the old `\b` version fails on it — i.e. the regression test holds the bug.

A false RED, caught only because the sharpened finals were run and watched. It is exactly as dishonest as a
false green, and it is what "watch every assertion go RED" buys in the other direction.

## Phase C — the controls, and 10 mutants

Three controls added to `negative-controls.spec.ts` on **one** login (the existing `pt-manager` contrast
seat — no new plumbing). Each absence is paired with **the contrast hero's OWN equivalent asserted PRESENT**
— a stronger liveness floor than `assertPageIsAlive`: it rules out "the stats did not load" as the reason
the hero's numbers are missing, because the same accessor is simultaneously producing the manager's.

**Every assertion was watched going RED:**

| # | mutant | result |
|---|---|---|
| 1 | seed `verified: 8 → 9` | RED — `pt-employee.verifiedSkills: seed=9 facts=8` |
| 2 | seed hero indent 6 → 4 (parse finds nothing) | RED — the fail-closed vacuity guard fires *before* any comparison |
| 3–5 | each sharpened Playthrough driven on the contrast seat | RED ×3 |
| 6 | timeline control asserts the contrast hero's OWN position absent | RED |
| 7 | verified role-context control ditto | RED |
| 8 | magnitude control compares her value to itself | RED |
| 9 | growth control points at a locator that IS present | RED |
| 10 | `dialogIsOpen()` re-open guard deleted (Phase E) | RED |

Mutant 2 matters most: the seed parser is regex-based, and a reconciliation over an empty parse passes
every comparison vacuously — the milestone's signature defect, found 17 times. It fails first, on purpose.

## Phase D — the batch gate, and a real flake

**Pre-fix batch (runs 1–3, cold reset-to-seed each):** `165 passed` ×2, then run 3 **164 passed / 1 failed**
— `pt-assignment-assign`, in a Playthrough this iter did not touch (it runs earlier in file order than
every file touched here, and the same code passed twice). **One flake in three runs, so clause 1's
surviving half was NOT met**, and it is not a "pre-existing" footnote: harden pass 1 named this specific
flake as still open once the casbin grant-accumulation that had masked it was fixed. So it was diagnosed.

## Phase E — the assign flake, root-caused from the trace

**Three hypotheses were refuted by measurement before the right one was found.**

1. *iter-11's bloated-policy hypothesis.* Measured: `g2 = 191`, `g3 = 171` for 191 memberships (Org B's 20
   correctly withheld), **0 orphans**. The policy is exactly as designed. **REFUTED.**
2. *The mask re-click closes the modal (antd `maskClosable`).* Probed: the re-click **throws** on the mask
   and the modal **survives**. **REFUTED.**
3. *`press('Enter')` with the dropdown closed submits the form and dismisses the modal.* Probed:
   `aria-expanded` stays `true` and `dialogCount` stays 1. **REFUTED.** And a fourth: the modal **survived
   151 s unattended**, so a background refetch on a settled table is not it either.

**The trace of the failing run settled it.** Times from `trace.zip`:

```
t+3.68 s  click the row's "Assign Skill Path"            (89 ms)
t+3.79 s  dialog appears
t+3.81 s  dialog title read: "Assign Skill Path to Aisha Andersen"
t+3.83 s  submit resolves to a VISIBLE <button disabled type="submit">   ← the modal is HEALTHY
t+3.85 s  combobox click → "element is not stable" ×3
t+4.15 s  "element was detached from the DOM, retrying"                  ← and the DIALOG never returns
t+18.9 s  attempt 2: `dialog >> combobox` never resolves   (15 s)
t+33.9 s  attempt 3: never resolves                        (15 s)
t+48.9 s  diagnostic `dialog >> button[Assign]` never resolves (20 s)
t+68.9 s  the spec's expect                                (15 s) → "element(s) not found"
          total ≈ 84 s   — matching the reported 1.4 m exactly
```

**3.8 s of real work followed by 80 s of correctly-bounded waiting on something that could not appear**,
reported against the submit button — three layers from the cause.

The cause is **structural**: the modal's title is *"Assign Skill Path to `<member>`"*, i.e. it is
**ROW-SCOPED** — rendered by the member row's action cell. A members-table re-render therefore **unmounts
it**. It opened 0.1 s after the row click and 2.2 s after the first row painted, i.e. while the table was
still settling; the settling re-render detached the Select's input and took the modal with it. Runs 1–2
finished the whole Playthrough in ~6 s and never met the race; run 3 was the slow run (suite 2.7 m vs 1.5 m).

**The fix is three parts, each independently justified:**

- **Recovery (the structural half).** `pickFirstSkillPath` now checks `dialogIsOpen()` at the top of every
  attempt and **re-opens the builder** when it is gone. Every bound in that ladder was already correct; what
  was missing was any way to notice the subject had died. A bounded retry loop that cannot re-establish its
  subject is not a retry loop.
- **Correctness that recovery forces.** Once a modal can be re-opened, reading the target member *before*
  the pick risks naming a different member than the one assigned (a re-order would fail the read-back for a
  reason unrelated to the platform). `openBuilderAndPickSkillPath()` returns the member named by the builder
  that **accepted** the pick, so the invariant *the member I name is the member I assigned to* holds by
  construction.
- **Not racing at all.** `waitForMembersTableSettled()` requires two equal assignable-count reads ~1 s
  apart — a semantic settle, not a banned `networkidle` one — because a presenter does not click a row the
  instant it paints.

**Recovery proven deterministically, not by fishing for a 1-in-3.** A probe reached the exact failing state
with a REAL user action (the modal's own Cancel — `Escape` is disabled on it, measured), then ran the ladder
against the corpse: it re-opened the builder, completed the pick, ended with the submit **ENABLED**, and
named the target. `openBuilderAndPickSkillPath` returned in 20.7 s where it previously failed at 84 s.
No DOM was manufactured (iter-07's rule: a control the application never learns about proves nothing).

**Post-fix gate — 3 consecutive cold reset-to-seed runs, 0 flake:**

| run | result | `ptreport` | wall |
|---|---|---|---|
| 4 | **166 passed**, rc 0 | 24 passing / 0 failing / 7 TODO / 0 unimplementable | 2.1 m |
| 5 | **166 passed**, rc 0 | same | 1.9 m |
| 6 | **166 passed**, rc 0 | same | 1.5 m |

`pt-assignment-assign`: **8.1 / 11.5 / 6.9 s** (pre-fix 6.0 / 6.1 / **84.0**). The settle pre-wait costs
~2–5 s on the happy path — an honest price for removing the largest single variance contributor in the suite.

## Verification sweep

- `ptvalidate`: **VALID** — 10 products, 31 use cases, 24 live Playthroughs, 7 TODO.
- Playwright: **166 passed** ×3 consecutive `--reset` runs, **0 flake**, rc 0.
- `tsc --noEmit`: clean. `gofmt -l` over the six rext sections: clean.
- Go: all **6** modules `rc=0`, **0 FAIL**.
- Python: `demo-stack + stack-core` **1286 passed / 1 skipped**, rc 0; `stack-verify + stack-injection`
  **437 passed / 1 skipped**, rc 0. (rc captured into a variable each time, never read off a pipe.)
- The deliberately DRIFTED `demo-2` cockpit-manifest fixture was backed up before the first `--reset` and
  restored after the last: sha **`99e2f315`**, verified byte-identical.

## Reported (never gated) — clause 1's suite statistic, per D-v28-13

Post-fix batch, n=3, same host: median per non-studio Playthrough **3.500 s = 1.0523×** of the 3.326 s
baseline, per-run range **0.8419× – 1.2327×** (**1.46×**). The untouched ORIGINAL-16 control subset reads
**0.9321×**. The pre-fix batch read **0.9170×** (range 0.8118× – 1.1425×), original-16 **0.7517×**.

**This batch is more evidence for the D-v28-13 recut, not against it.** The control subset — code no iter
has touched since iter-03 — now spans **0.5281× → 1.0762× → 0.9321×** across batches. A gate at 0.79× still
sits inside its own noise floor. **This iter landed no speed mechanism, so clause 1's leg half has nothing
new to measure**; its other half (0 flake ×3) is **MET on the post-fix batch**, and was met only because the
flake was fixed rather than re-rolled.

**Environment:** `Kirality-Mac-Pro-6.local`, darwin 25.1.0, Docker VM 9.70 GiB (vs the documented 12 GB
floor); `demo-2` offset 20000, localhost/http, `--no-public-host`. Per D-v28-12 no number here is comparable
to billion's 228 s.

## Close — 2026-07-29

**Outcome:** **negative controls 13 → 16 of 24**, by refuting the premise that the 9 structural finals could
not have one: their FINALS were structural, not their subject. Three profile Playthroughs that would have
passed against anybody's profile now name the hero's own seeded data — magnitudes reconciled against the
seed by a new fail-closed fence — and the same contrast seat that could not falsify the structural version
falsifies all three. **And the `pt-assignment-assign` flake was root-caused from the failing run's trace
(a ROW-SCOPED modal that a members-table re-render unmounts, against a retry ladder with no way to
re-establish its subject), fixed, the recovery proven deterministically, and the 3× gate re-run clean** —
so clause 1's flake half is met by a fix, not by a favourable batch.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — clause 2 mutating **6/5 MET**, `blocked` **1/1 MET**, negative controls **16 of 24**
(8 remaining: 6 structural-final Playthroughs + 2 studio blocked behind `FIX-M256-studio-false-green`);
clause 3 verdict half **COMPLETE (28/28, 0 unimplementable)**, landed half still short (org-admin 2/4,
onboarding 1/5); clause 1 leg half **N/A this iter** (no speed mechanism landed), flake half **MET — 0 flake
across runs 4–6**; **D-v28-5 still unfixed** (unblocked, not started).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this tik landed its planned scope and moved the primary metric 13→16) — (3) re-scope: n (0 of 31 curated UCs `unimplementable`; the mechanism iter-12 bounded turned out to have more reach, not less) — (4) user-blocker: n (the one red in the batch was diagnosed and FIXED inside the iter, not escalated) — (5) cap-reached: n (1st tik of this invocation) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D61 (a structural final's problem is the ASSERTION, not the vantage — iter-12's bound was on the mechanism as applied, not on the Playthroughs), D62 (a sharpened final's magnitudes must be machine-linked to the seed, and the link must fail CLOSED on a vacuous parse), D63 (`\b` in a `hasText` regex is unreliable — `textContent` concatenates sibling nodes; the same constant is safe under `getByText` and broken under `hasText`), D64 (a negative control's liveness floor should be the contrast hero's OWN equivalent read through the SAME accessor, not a generic page-alive check), D65 (**the assign flake is a ROW-SCOPED modal unmounted by a members-table re-render; a bounded retry loop must also be able to RE-ESTABLISH a dead subject**), D66 (once a ladder can re-open a modal, the target must be read from the builder that ACCEPTED the pick — recovery creates a correctness obligation), D67 (three refuted hypotheses cost less than one guessed fix: the trace's own arithmetic named the cause).
**Side-deliverables:**
- `url-shapes.unit.spec.ts` gained the timeline-pattern coverage its header had advertised since M203 —
  and the first draft of that block asserted the tolerant pattern over prose claiming the file had never
  pinned it. **The claim was false** (an existing describe has pinned it since M203) and running the suite
  is what said so. Duplicating a guard while announcing its absence is a smaller version of the same
  dishonesty as a guard that cannot fail, so the duplicate was removed and only the open-vs-closed
  asymmetry is asserted.
- `TIMELINE_CURRENT_ENTRY` also pinned against the `/g`-`lastIndex` hazard, the class that would have
  halved iter-12's control set.
**Routes carried forward:**
- `NEGCTL-M256-cross-vantage` → **16 of 24; 8 remain, and the classification has changed.** iter-12 split
  the residual into "9 structural (no contrast vantage possible)" + "2 studio". Three of those nine are now
  done, and the reason generalizes: **the other six are reachable by the same move** — sharpen the final to
  name seeded data, then run it against a contrast vantage. They are `pt-workforce-{roster,funnel,succession,org-feedback}`,
  `pt-activity-drilldown` and `pt-hiring-recruiter-compare`. Note the hiring one needs care: iter-12
  measured that its contrast vantage **ejects the browser to production**, so its control must come from a
  sharpened final on the *same* vantage, not a contrast org. The 2 studio remain blocked behind
  `FIX-M256-studio-false-green`.
- `PT-M256-readiness-step-asserts` → **still open, unchanged.** Re-scope `MANAGER_STEP_NAMES` inside the
  method panel. Not touched this iter; it is a sibling of exactly the work done here (an assertion that
  matches page-wide and is satisfied by the wrong state), so it should ride with the next sharpening batch.
- `FLAKE-M256-assign-under-bloated-policy` → **CLOSED, and its hypothesis was WRONG.** The policy was
  measured clean (g3 171 / 191 memberships, 0 orphans). The real cause is recorded as D65. Superseded.
- `MEASURE-M256-clause1-sampling` → **more evidence, same escalation.** The untouched control subset has now
  been observed at 0.5281×, 1.0762×, 0.7517× and 0.9321× across batches on one host.
- Everything else from iter-12's list stands unchanged.
**Lessons:**
1. **A bound is not a recovery.** The assign ladder had every interaction correctly bounded — that was the
   harden pass's own achievement — and it still spent 80 s waiting on a dialog that had ceased to exist,
   then blamed the submit button. Bounding makes a stuck attempt yield; it does not make a *dead subject*
   detectable. Any retry loop over a mounted UI object needs both: bounded interactions AND a check that the
   object is still there, with a way to re-establish it.
2. **Read the trace before proposing a mechanism.** Three plausible hypotheses — a bloated policy, a mask
   click, an Enter keypress — were each refuted by a probe, and the fourth was handed over by the trace's
   own arithmetic (84 s = 20 + 3×15 + 20 + 15). Guessing would have shipped a fix for a cause that was not
   there, which iter-07 already recorded as the worst outcome available.
3. **A "mechanism has no reach here" finding should be re-tested when its inputs change.** iter-12
   correctly measured that a *structural* final has no contrast vantage, and that reading hardened into "nine
   Playthroughs cannot have a control". The finding was about the assertion, and the assertion was ours to
   change. A bound on a mechanism is only as durable as the thing it was measured against.
