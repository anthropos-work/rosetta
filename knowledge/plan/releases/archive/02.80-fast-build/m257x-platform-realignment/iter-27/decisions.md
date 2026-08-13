---
milestone: M257x
iter: 27
---

# iter-27 decisions

## D-M257x-27-1 — a seeder's treatment of HEROES is a policy, and it must be declared, not drawn

`memberInShare(prefix, i, bucket, share)` decides deterministically whether population slot `i` is in a
named subset. A hero occupies a population slot like anyone else — measured, the pt-world hero's user id is
exactly `deterministicUUID("demo-1:story:pt-org-a:user:1")` — so **every share draw silently decides whether
the seeded world's protagonist appears in that surface.**

And the draw is **fixed**: the key prefix is pinned by `stack: demo-1` *in the seed YAML*, not derived from
the stack being seeded. It varies neither by stack nor by run. A hero the hash excludes is excluded
**forever**, which is why `pt-workforce-org-feedback` had never passed against this world and could not have.

**The rejected fix was "heroes are always included."** It is false as a rule: `population_evidence.go`
excludes heroes on purpose (the PersonaSeeder writes their evidence far richer, so including them here would
contradict rather than add), and `target_roles.go` is legitimately indifferent. **The defect was never the
choice — it was that one seeder made NO choice, and nothing could tell that apart from a choice.**

Adopted, and it is `D-M257x-12`'s shape one layer down: **declare** the policy per seeder
(`hero-always` / `hero-skipped` / `hero-branched` / `hero-indifferent`, each with a reason), **derive** the
scope from the source (every non-test seeder file whose AST contains a `memberInShare` call), and check the
two against each other **in both directions** — an undeclared share-gated seeder goes RED naming itself, and
a declaration for a file that no longer draws goes RED too. Being indifferent is fine. Being *accidentally*
indifferent is the defect.

Also landed with it: the **org-less guard** (`if orgLess[i] { continue }`) the sibling seeders already carry.
Unconditional hero inclusion would otherwise write an org-less hero a feedback row **and** the app-side
session row that carries `organization_id = st.OrgID` — activity inside an org she is deliberately not a
member of. pt-world declares one such persona, so this was reachable, and only her share draw was
suppressing it.

## D-M257x-27-2 — the cluster was not a cluster: four ids, at least three mechanisms

The hand-off named `CHECK-M257x-iter15-manager-reads-empty` as *"four failures sharing ONE coherent
signature — a manager-vantage read reports the seeded hero as absent."* Measured per id, that is **false for
at least two of the four**, and the differential is what makes the remaining ones targetable:

- **`pt-workforce-org-feedback`** — genuinely un-seeded. Seed-side. Fixed here.
- **`pt-workforce-succession`** — the hero's row **IS** seeded (`interview_extraction_results`
  `957d5253-…` FK'd to her real session), because `succession.go:114` carries the explicit hero exemption
  `feedback.go` lacked. The data is present and the surface does not show it → **read-side**.
- **`pt-workforce-funnel`** — **not a hero absence at all.** The preceding `spotlight` visible assert
  **passed**: her card renders. Only the role text *inside* the card is missing, while the role is present
  on every DB axis checked (`user_basic_info.job_title`, the current `user_experiences` row, `job_role_id`),
  and 40/40 org members carry a `job_title`. DOM/locator-shaped.
- **`pt-activity-drilldown`** — conditional on which content `drillIntoActiveContent()` selects; requires
  the hero to hold a session on *that* sim. Not measured further.

The hand-off's own enumeration was also wrong on one row (it attributed the drilldown's failure text to
`assign-and-track.UC2`), and `pt-assignment-assign` — an affordance **count**, 15 vs 14 — is a **singleton**
that was being counted into the class. Two applications of the same §5 rule in one iter: *re-derive the
enumeration, not just the values.*

## D-M257x-27-3 — the inverted mutant, and the control that made it readable

The fence's first cut asserted the guard was *"guarded by a condition referencing `isHero`."* Mutant **M4**
inverted it — `if isHero && !memberInShare(…)`, heroes gated and everyone else free, the exact opposite
semantics — and **the fence reported GREEN**. Three removal-mutants had all died correctly and could not
have found it: deleting a guard and reversing it are different edits, and only one is what a careless editor
actually writes.

Repaired by asserting against the **negation as a parsed node** (`*ast.UnaryExpr` with `token.NOT`) rather
than the identifier's appearance anywhere in the condition. Promoted to `platform-alignment.md` §8 rule 5 as
the **inverted-mutant** addendum, explicitly paired with the no-op-control rule it depends on: the
declared-GREEN control is what made M4's verdict interpretable at all.

Worth naming plainly: **this milestone's dominant defect class, committed by the milestone, in the very
fence written to prevent it.** Fifth occurrence of *a check that reports without measuring*.

## D-M257x-27-4 — the suite was SILENT here, not arguing

Three times this milestone the existing tests were found *arguing for* the defect (iters 16, 18, 24). Not
here: `stack-seeding` was fully green before and after, and **nothing pinned the hero's absence at all**.
That is TOK-01's baseline honesty note landing exactly as written — the seeders assert against a recording
fake `Conn`, so a green suite is *structurally silent* about which population slots the hash selected. The
absence of a red test is not evidence, and the only thing that could have caught this was a live DB read.
