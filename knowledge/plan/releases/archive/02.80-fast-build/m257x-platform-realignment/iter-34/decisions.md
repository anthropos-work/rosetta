# iter-34 decisions

## D-M257x-34-1 — the adjudication rule, fixed BEFORE any report was read

Clause 5 turns on a **blocker count**, so the count is the deliverable and the count can be gamed in
both directions — by waving findings through to look thorough, or by downgrading them to reach zero.
The rule was therefore fixed in advance, at 2026-08-02 00:40, with no report yet written:

1. **A reported blocker counts only if I independently verify the false claim against platform source.**
   The auditor's citation is a lead, not evidence. This is the milestone's standing *re-derive, don't
   re-match* rule applied to the audit itself.
2. **A reported blocker I verify as TRUE-at-HEAD is downgraded to `not-a-finding`** and recorded with
   the re-derivation, so the downgrade is auditable rather than a judgement call.
3. **A reported blocker I verify as false-but-harmless** (nobody could act on it wrongly) is downgraded
   to `minor` and joins `DOC-M257x-iter33-corpus-minors`.
4. **A verified blocker is fixed in this iter**, and clause 5 is then graded on the reading that was
   *actually taken* — never on the post-fix state, which is unmeasured by construction. That is
   iter-33's precedent and it binds symmetrically: if this pass returns blockers, clause 5 stays NOT MET
   and a fourth pass is routed forward, exactly as iter-33 refused to grade on an absent measurement.
5. **Downgrades are reported.** A pass that reaches zero by reclassification is reported as
   "N reported, M verified, K downgraded — with the downgrades listed", never as a bare zero.

**Why rule 2 needed writing down:** the parent-side re-derivation run *before* the reports arrived
(`evidence/parent-rederivation.md`) produced a wrong number on its first attempt — 22 where the doc said
18 — because it used a different denominator than the doc's own sentence. Had that been graded on the
first read, it would have filed a blocker against a **correct** claim. A false positive costs this
clause exactly as much as a false negative, and there was no reason to assume the auditors are immune to
the error their parent had just made.

## D-M257x-34-2 — the partition was deliberately re-cut

Pass 1's five groups were organised by subject (the big architecture docs; ai-readiness + studio; …).
This pass re-cut them by **line balance with swept/unswept mixing**, so that no auditor inherits pass 1's
group boundaries. Correlated blind spots are a property of *how a corpus is divided*, not only of who
reads it; a group that missed something as a unit is likelier to miss it again if handed the same unit.
Every group got at least one file the repair sweep had edited and several it had never touched, so no
auditor could form a uniform prior about the freshness of its assignment.

## D-M257x-34-3 — both pre-registered predictions were refuted, and P2's refutation is the iteration's result

P1 (1–5 blockers) missed low: the pass returned **11**. P2 (residual concentrates in the 27 unswept files)
was refuted **in the opposite direction** — 9 of 11 blockers landed in the 13 files iter-33's repair had
touched (0.69/file) against 2 in the 27 it never opened (0.074/file), a ~9× density difference.

P2 was reasoned from pass-1's group-5 density (1 blocker across 18 files) being an order of magnitude
below every other group, and read as under-detection. **It was not under-detection; those files are
genuinely clean.** Two auditors said so unprompted after verifying ~40 exact citations apiece.

The value of having written P2 down is that its refutation is *load-bearing*: had the prediction not been
recorded, the 9× result would have read as an unremarkable "found some more" rather than as a measurement
that the repair pass is the dominant source of remaining debt. Promoted to `platform-alignment.md` §5
rule 18.

## D-M257x-34-4 — the parent's own re-derivation verified the count and missed the claim

`evidence/parent-rederivation.md` was written before any auditor reported, specifically to check the
multi-tenancy fence that had already been wrong twice. It reproduced 30 / 7 / 18 and all nine named files
exactly and concluded the fence held. Auditor E then found that `org_membership.go` — listed **first**
among the "no mixin and no policy at all" set — declares its own fail-closed `Policy()`, the only one of
the 18 that does.

The sentence is a conjunction; the check tested one conjunct. **The count was right and the claim was
false**, for the third consecutive generation of the same fence, each time failing toward
*"isolation is handled."* Promoted to `platform-alignment.md` §5 rule 17. The rewritten fence now carries
its own derivation command so the next reader re-derives rather than trusts.

## D-M257x-34-5 — clause 5 graded on the reading actually taken

11 blockers, all fixed; the post-fix state is unmeasured by construction. Rule `D-M257x-34-1(4)`
pre-committed to grading on the reading taken, and iter-33 set the precedent by refusing to mark clause 5
met on an absent measurement. **Clause 5 NOT MET; gate stays 3 of 5.**

Note on the shape of the numbers: 19 → 6 → 11 is **not** a convergence curve — the three passes had
different scopes (40 unswept / 13 swept / 40 post-repair). The only comparable pair is pass 1's 19 over 40
files against this pass's 11 over the same 40. The residual is smaller *and* differently distributed. The
regress is not ended by a fourth pass; it is ended by `CHECK-M257x-iter33-derived-fact-fence`.

## D-M257x-34-6 — prep note: `pt-activity-drilldown`'s coupling (read-only, no change made)

Scoped during the audit's wall-clock, read-only, no edit. The spec fails at
`playthroughs/e2e/tests/activity-drilldown.spec.ts:113` on `heroRow.count() > 0` — the hero's name absent
from the per-member breakdown of **the first content row**. The coupling is positional:
`activity-dashboard-page.ts:77-79` (`activeContentRowLink`) drills `contentRows().first()`, and the spec's
own comment (`:103-105`) admits the determinism argument is "the grid sorts by most-recent activity and
hero sessions are dated today" — which iter-27's added hero sessions could reorder.

The assertion's stated purpose (`:97-98`) is *"this is the manager's OWN tenant's breakdown rather than any
populated org's."* **Selecting the drill target by hero participation instead of by grid position preserves
that purpose exactly and removes the positional coupling** — another org's manager still finds none of
their content carrying this hero. Recorded as the fix shape for `CHECK-M257x-iter27-drilldown-target-coupling`;
**measure which content the first row actually is before implementing**, per the standing route.

## D-M257x-34-7 — the protocol's oldest rule caught this iteration's own author

While correcting `architecture_overview.md`'s AI-routing claim, the check for Mistral in `app` was run as
`grep -ril mistral stack-demo/app --include=*.go`. zsh **rejected the glob** (`no matches found:
--include=*.go`) and the empty output was read as absence — which then became a corpus claim that "Mistral
is not in `app` at all". Mistral is live Go code (`internal/cms/studio/markdownManager.go:11,19`, Studio
document OCR).

This is **§5 rule 1** of `platform-alignment.md` — *never let a search's stderr go unread; an engine
rejection is indistinguishable from "no matches"* — committed **in the same iteration that appended rules
17 and 18 to that same section**. No new rule is warranted: rule 1 already covers it, correctly, and had
been at the top of the list precisely because it is the most common failure. The finding is about
*execution*, not doctrine, and it is recorded here rather than promoted so the rule list does not grow a
duplicate.
