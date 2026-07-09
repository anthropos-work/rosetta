# iter-06 — Decisions (tik / cleanup-measurement under TOK-01)

## D1 — Sub-condition (f) is MET via the schema-ref sweep (the gate's literal criterion)
The gate's (f) is "0 residual skiller-SCHEMA references in any path the tooling queries." A comprehensive
sweep of rext production (stack-snapshot/stack-seeding/stack-verify, excl tests) for
`skiller.<one of the 10 taxonomy tables>` returns 0 hits — every taxonomy SQL path uses `public.*` (M209's
re-ground, confirmed live in iter-02/03/05). So (f) is MET. This is distinct from the word "skiller"
appearing as a concept name in comments or as an intentional test token — those are not schema references
and not queried SQL.

## D2 — The dead repos/run.sh skiller branches are correctly LEFT (removal reverted)
`stack-verify/repos/run.sh` has 2 `case skiller)` branches (a test command + a timeout). skiller is absent
from repos.yml, so the repos loop never iterates it → the branches are DEAD code, and they are NOT schema
references. Removing them is cosmetically tidy but breaks `test_absent_repo_resolves_under_stack_root`,
which plants a synthetic `repos.yml` with `name: skiller` as a generic ABSENT-repo fixture and relies on
`cmd_for_repo skiller` returning a command to reach its $STACK_ROOT path-resolution assertion (the $DEVDIR
bug regression). Completing the removal would require re-writing that fixture (and others using "skiller"
as a token) to a different repo name — multi-fixture test churn for ZERO gate value (the branches are dead
+ non-schema; (f) is already met). Per "don't gold-plate," the removal was reverted; the rext tree is clean
and the verify suite is 104/104 green. If a future release wants the cosmetic purge, it should rename the
test fixtures in the same change.

## D3 — Session exits at the 5-tik cap (iter-06 is the 5th tik)
Tiks this session: iter-02, 03, 04, 05, 06 = 5 (bootstrap tok iter-01 does not count). The cap fires after
iter-06 closes. Gate stands at 5/6 sub-conditions proven on the warm merged stack; the remaining (e) M42
coverage + v2.0 Playthroughs + the full COLD /dev-up + /demo-up proofs are routed to the next session (they
need the UI tier + cold demo bring-up — heavy + reap-risky, not a reap-safe single-turn foreground op).
