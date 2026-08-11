---
milestone: M257x
title: "platform re-alignment — retro"
release: v2.8 "fast build"
closed: 2026-08-11
---

# M257x — retro

## Summary

The team was migrating its microservices back into the `app` monolith and **nobody on our side knew how
far that had got.** M257 hit the wall and stopped: its own health check was querying a table the platform
had deleted, with a `|| echo 0` turning the error into `0` — a number that reads exactly like *"the demo
seeded, just thinly."* M257x was inserted to find out where the migration actually is, write it down where
it cannot rot, and make **both** repos work against the platform as it is.

**288 iterations and 73 harden passes later**, four of five gate clauses are met and proven. The fifth was
placed **out of scope by the user** on 2026-08-11 (`TOK-09`), who narrowed the definition of done to
*architecture + repo/component register + buildability* — all three proven rather than asserted. **This is
a scope ruling. Clause 5 was never met and is not being declared met.**

## Incidents this cycle

| # | what | severity | disposition |
|---|---|---|---|
| 1 | **A demo attempted `s3:PutObject` against the PRODUCTION storage bucket.** Refused **403** — by an IAM policy on an account **we do not control**, not by this design. Nothing was written. | **P1** | Contained in code at both pointers (iter-284) + `safety.md` corrected. **Still open in practice:** the containment is proven by a unit test on the emitter and **on no running stack**; `demo-2` predates it by nine hours and the dev-side strip is demo-only. Routed to M258. |
| 2 | **The corpus asserted a demo "cannot write prod" flatly, in six-plus places** — one of them saying it *"never uploads a byte anywhere"* — while the same document carried its own retraction 55 lines below. | P2 | Four sites fixed at iter-284; the **guarantee itself**, the header that pointed at it, and the exposure rationale that leans on it fixed at this close. |
| 3 | **Harden pass 73 reported `demo-stack` GREEN off a truncated output tail.** The `9 failed,` clause was immediately before the fragment it read and was cut off. It was written into the ledger and said out loud. | P2 | Self-corrected in the same pass, re-run whole (byte-identical), and recorded. The pass auditing for this class committed it. |
| 4 | **A `--grep`'d Playthrough run was graded as a suite result** for several iters, while the harness printed its own disqualification on every one of them. | P2 | iter-273 took the first binding full-suite measurement. |
| 5 | **This close introduced one regression** — a comment added to `up-injected.sh` pushed a token past a magic `src[i:i + 2200]` slice and reddened a fence that touched no opt-out. | P3 | Fixed at the derivation (the next test in the same class already derived the boundary), re-proven RED under mutation. |

## What went well

**The fences worked, and they worked on the people who built them.** Three separate times this close, an
edit turned a guard RED for exactly the reason the guard existed: shifting line numbers in a script broke
ten corpus citations and `demo_knob_guard` caught all ten and repaired them with its own `--fix`. That is
the milestone's thesis demonstrating itself at its own close.

**Measurement beat inheritance, repeatedly.** The two headline corrections both came from reading source
and both refuted a standing corpus claim. `roadrunner` was **deleted and replaced in-process**, not merged
into `app` — `git log --all --diff-filter=A -- internal/roadrunner` returns **0 commits, ever**, in a
6,728-ref clone, with a positive control that returns 3. And `cms`'s ECS service is **destroyed**, settled
by cloning `infrastructure` — a repo that had never been in any clone set, which made the question look
*unmeasurable* when it was only *uncloned*. A clone-set limit is not a measurement limit.

**The protocol produced a deliverable, not just an outcome.** `corpus/ops/platform-alignment.md` did not
exist when the milestone opened — that absence *was* the gap. It exists now, and the recurring class it
documents has a machine fence behind it.

**Pre-registration did real work.** Iters routinely sealed five predictions before measuring, and the
refutations are where the value was: iter-266 had three of five refuted and recorded that *"the refutations
ARE the iter."*

## What didn't

**The front door was never swept.** 288 iterations of per-service alignment, and `corpus/README.md:48`
still called the backend tier *"8 Go microservices (Backend, CMS, Sentinel, etc.)"* — the **exact phrase**
`CLAUDE.md` records as hunted down and corrected at harden pass 57. The sweep reached `CLAUDE.md` and
missed the index next to it. Five more index-layer claims were stale the same way. **A sweep scoped by
subject misses the files organised by navigation.**

**Eight iterations closed without a ledger line, and two without a directory at all.** The mechanism is
structural, not carelessness: the grading is written in the iter's own `## Close` section, and the ledger
is a different file in a different lane — so the write that creates the finding and the write that would
represent it are never the same edit. `deferrals-audit.md` §12 named this shape; it then happened six more
times.

**The instrument that measures whole-section coverage is partially blind on this host, and that is why the
final harden could not stabilize.** One cause, eight symptoms: no `tests/` directory has an `__init__.py`,
and on this interpreter unittest's loader cannot import a namespace-package submodule by dotted name — so
the two-runner cross-check is currently a **one-runner check that reports RED**. Coverage did not
stabilize because it could not be measured, which is a finding rather than a gap.

**The milestone's own numbers drifted while it was auditing everyone else's.** Fifteen test counts across
five handbooks were stale at close — including one whose parenthetical *explaining why it had drifted* had
itself drifted, and `stack-core`'s handbook, which described a two-file suite that is 94 files and gave a
`pytest` invocation that cannot run on this host at all.

**A closed-list ruling arrived at iteration 283 of 288.** That it landed at all is good; that it took
until then is the cost of an open-ended gate on a moving target. The `re_scope_trigger` written at
iter-01 anticipated the target moving faster than we could track it — it never fired, but three separate
strategies (`TOK-07`, `TOK-08`, and finally the user's `TOK-09`) were needed to bound the subject set.

## Carried forward

Full routing in [`carry-forward.md`](carry-forward.md). Eleven items in five clusters, plus one block fate,
**all → M258**. Not to M257: its `exit_gate` still names `odysseus`, retired by `D-v28-15`.

| cluster | one line | fate |
|---|---|---|
| 1 | The prod-bucket pointer is contained **in code**, not on either running stack | LAND-NEXT → M258 |
| 2 | The section census's second runner is dead on this interpreter | LAND-NEXT → M258 |
| 3 | 22 of 28 new fences are not RED-proven; two live outside every census | LAND-NEXT → M258 |
| 4 | `buildbench` has no profile for the host the release names, and asserts no wall-clock at all | LAND-NEXT → M258 |
| 5 | Residual per-item content work found **after** `TOK-09` closed the list | LAND-NEXT → M258 |
| — | 215 tokens carried across ≥3 iters — a **marker** count, not a state count | block fate → M258 |

**Escape-hatch deferrals: zero.**

⚠️ **One carry-forward outranks the rest operationally:** the two tooling fixes landed at this close live
in the `rosetta-extensions` **authoring copy** and are **not on a pushed tag**. A stack clones rext from
**origin** at a pinned tag, so until someone tags and `git push --tags`, those fixes are unreachable to
every stack. *Tagging is not publishing* — the lesson M236 lost an iteration to.

## Metrics delta

Sourced from [`metrics.json`](metrics.json).

| metric | value |
|---|---|
| Iterations | **288** (279 tik · 8 tok · 1 measurement), 9 strategies `TOK-01..TOK-09` |
| Harden passes | **73**, final stop condition `cap reached without stabilization` |
| Commits | **655** rosetta · **339** rext |
| Python tests (rext) | **5,177 passed** · 27 failed · 10 skipped across 5 suites |
| Go sections | **6 ok**, 0 failing |
| Guard family | **31 GREEN · 0 RED · 0 could-not-check** at the merge sha |
| Gate | **4 of 5 clauses met**; clause 5 **out of scope by user ruling** |
| Handbook counts reconciled | **15** drifted figures across 5 sections / 7 files |
| Ledger orphans found + repaired | **8** (2 with no directory at all) |
| Platform-repo edits | **0** |
| Net-new dependencies | **0** |
| Escape-hatch deferrals | **0** |

## The one transferable line

**A clone-set limit is not a measurement limit, and a truncated tail is not a result.** Both of this
milestone's biggest corrections came from noticing that something recorded as *unmeasurable* had simply
never been fetched — and its most embarrassing moment came from reading the end of a buffer and calling it
a verdict. The corpus is now fenced against the first. The second is a habit, and habits do not have
fences.
