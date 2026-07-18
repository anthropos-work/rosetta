# Hardening Ledger — M228 Second night

## Pass 1 — 2026-07-18 — final

**Iters hardened this pass:** all milestone-touched code (iter-01 tok · iter-02 · iter-03).

**Tiks covered since prior pass:** all iters in milestone (first + only harden pass).

**Milestone-touched code footprint (`casting-call-m227-sections`..HEAD in rext):**
- `stack-seeding/seeders/hiring_scope.go` (+23) — the guard-consult extension
- `stack-seeding/seeders/feedback.go` (+9) — F2/F3 guard (mirror-row leak)
- `stack-seeding/seeders/succession.go` (+10) — F1 guard (FK-crash)
- `stack-seeding/seeders/hiring_scope_test.go` (+8), `hiring_funnel_test.go` (+36) — inline regressions
- `stack-verify/e2e/tests/render-hiring-comparison.spec.ts` (+36) — the render-probe hardening (tooling-iter)

**Coverage on touched files:** `stack-seeding/seeders/` = **96.8%** of statements (measured). No shallow tests
added — the existing inline regressions already exceed the bar; padding the number would violate the
anti-shallow-test discipline.

**Dimension scan (final-mode, 6 dimensions against cumulative scope):**
- **Test depth** — `TestGenericActivitySeeders_SkipHiringOrg` enumerates ALL 8 generic activity seeders
  (jobsim-sessions · activity · skillpath-sessions · assignments · hero-activity · persona · **feedback [M228 F2/F3]**
  · **succession [M228 F1]**), each with a NEGATIVE assert (0 rows for a hiring org) + a POSITIVE control
  (>0 rows for a workforce org). `TestSkipGenericActivityForHiringOrg_Predicate` pins the guard predicate.
  `hiring_funnel_test.go` pins the round-robin 1-sim/candidate distribution + the mirror pair.
- **Error paths** — the F1 FK-crash path is exercised: each seeder's `Seed()` against a hiring-only seed
  asserts `err == nil` (succession no longer FKs a now-skipped population session → no "seed failed").
- **Edge cases** — the workforce positive controls prove the guard skips ONLY the hiring org, not the seeder
  wholesale (the boundary that matters).
- **Regression bundles / cross-iter integration** — the enumerated seeder table IS the cross-iter integration:
  M228's feedback (F2/F3) + succession (F1) were added to the SAME table the M227 fix-#1 seeders live in, so
  a future regression on any one of the 8 fails the shared test. This is precisely the bug class M228 fixed
  (a mirror-writing seeder forgotten by fix #1).
- **Fuzzing** — n/a: the guard's input is a `ResolvedStory` boolean predicate; no non-trivial input surface.
- **Perf thresholds** — n/a: the milestone made no path performance-sensitive at the seed layer (the live
  latency gate condition 5, p95 1.27 s, is proven in iter-03, not a unit threshold).

**Render-probe (tooling-iter carve-out):** dimension 1 (test depth) is satisfied by the live 5/5 per-sim gated
run on billion (the probe IS the test harness; its correctness is proven by measuring correctly against the
real surface). tsc type-clean. No unit-fuzz surface worth pinning (simple env-knob parsing).

**Bugs surfaced + fixed inline:** none — no new defects. The code is LIVE-PROVEN on billion (42 clean hiring
sessions, 0 non-hiring leak, no FK crash; recruiter render 5/5, candidate heroes usable, p95 1.27 s).

**Flakes stabilized:** none — flake gate 3/3 clean on the M228 guard + funnel tests.

**Knowledge backfill:** the `hiring_scope.go` doc-comment already captures the M228 correction in full — WHY
FeedbackSeeder/SuccessionSeeder must consult the guard (mirror rows since M42m / FK to a now-skipped session)
and WHY the remaining dashboard seeders (MembershipSkills/Tags/TargetRoles/PopulationEvidence) genuinely stay
untouched. The protocol-level insight ("a new mirror-writing seeder MUST be added to the guard consult-list +
the enumerated regression table") lives in-code where the next author will see it.

**Stop condition:** **stabilized** — coverage already at 96.8% with comprehensive enumerated + positive/negative
control tests, the final-mode dimension scan surfaced no new gaps, the flake gate is 3/3 clean, and the whole
pipeline is live-proven on billion. Shallow padding tests explicitly declined per the anti-pattern discipline.
This entry satisfies `/developer-kit:close-milestone`'s iterative-milestone final-harden gate.
