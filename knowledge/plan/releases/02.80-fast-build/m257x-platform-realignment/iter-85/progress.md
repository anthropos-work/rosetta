# iter-85 — progress

**Type:** tik, under `TOK-05`. **The first repair since iter-81, and the first ever graded by a
post-condition.**

---

## THE HEADLINE

**Reach 11/11 = 100 %**, measured by `repair_reach_guard` against this iter's own declared input ledger —
against iter-81's **74.1 %**. Pre-registered: *"0 unreached, or the iter does not close as
`closed-fixed`."* **0 unreached.**

And the fence iter-81 skipped caught **my own repair** mid-flight: `repair_leak_guard` went **RED** on it,
naming `platform-alignment.md:1060` — because §5 **rule 40**, which I wrote at iter-83, quotes the false
`graphql`-profile sentence verbatim as its worked example. Dispositioned by **waiver with a written
reason** (the guard's sanctioned path; waivers are *reported* on every run, never silent), not by
paraphrase — a rule that will not name the sentence it is about teaches nothing.

## What landed

**Q2 — 7 claims about facts that were DELETED.** Restated or dropped; **not one re-anchored** (§4 Trap A).

| site | was | now |
|---|---|---|
| `graphql-wundergraph.md:13` | *"the `graphql` profile name survives… and is now simply the default"* | the profile is **gone**; `PROFILE ?= core`; the token is in no `profiles:` key, so asking for it **exits 0** and starts only the floor |
| `cms.md:8` | *"the fourth and **last** engine consolidated"* | the **fourth**, **not the last** — v9.0 then folded `storage` and `messenger` |
| `backend.md:218` | *"producer and consumer of **all five** streams"* | **four of five**; `app` only **subscribes** to `skiller` and nothing publishes to it — enumerated over every publisher constructor at `b948604f` |
| `roadrunner.md:113` | *"jobsimulation consumes it as the async signal"* | **nothing consumes it** — the consumer was deleted, replaced by an Asynq task |
| `services/README.md:37` | *"other services fire an RPC"* | they **publish Redis Stream events**; `MESSENGER_RPC_ADDR` exists in **no** repo and `git log -S` over all platform history returns **0** commits |
| `messenger.md:7` | the same sentence, in its twin | repaired in the same pass (**§5 rule 19** — a claim does not respect a file boundary) |
| `architecture_overview.md:295` + `:311` | prod RPC list included `storage`; the retraction scoped itself *"locally"* and thereby **affirmed** a dead prod edge | `storage` removed from the prod list; the retraction widened to **both columns** |
| `alignment_testing.md:360` | *"exits **rc=2** and nothing treats that as a failure"* | **rc=3** since M219, with a refuses-to-be-mistaken banner. The passage was telling readers to read a `2` as a missing Node module |

**The 2 confirmed leak sites** (`FIX-M257x-iter83-leak-guard-3-sites`): `CLAUDE.md:285`'s runnable
`make up  # … (graphql profile)` → `(core profile — the default)`; and `platform-alignment.md:1305`'s
*"`STORAGE_RPC_ADDR` is read by `main.go`"* → **read by nothing** (3 hits at that ref, all comments), with
the `b948604` state stated separately so the two refs cannot be conflated again.

## 🔴 The live rext defect — `FIX-M257x-iter84-dev-stack-default-profile`

`dev-stack:186` and `:414` initialised `profile="graphql"`. A bare `dev-stack up N` — what `/dev-up N`
runs with default flags — executed `docker compose --profile graphql up -d`, and **that token exits 0 and
starts only the always-on floor**: Postgres answers, `docker ps` is non-empty, the application is absent.

**Fixed by DERIVATION, not by substituting `core`** (`D-M257x-59-2`, TOK-04 P4 — *derive, else fence, else
declare*). Both entry points now default to empty and resolve through `platform_topology.default_profile()`,
which reads `backend`'s own `profiles:` list and **raises rather than guessing**.

**Proven, not asserted — both directions:**

| | result |
|---|---|
| against the real clone (`stack-demo/platform` @ `0dab54d`) | `derived => core` |
| against a platform dir with no compose | **dies loud**, rc **1**, and the line after the call **never runs** |

The negative case was tested in the *exact* call form the script uses (`[ -n "$p" ] \|\| p="$(…)"`),
because a `die` inside a command substitution exits the subshell — the failure would have been a silent
empty profile, which is the same defect one layer down.

The `/dev-up` skill docs were realigned in the same commit: `SKILL.md:147` (the `--profile` default is now
**derived**, not a literal), `:74`, `:175` and `reference.md:38`. Leaving them would have induced a
doc↔tool mismatch in the pass that fixed the tool.

## What did NOT land, and why it is declared rather than discovered

**Q1 (13), Q3 (8), Q4 (7), Q5 (1) and the P4 membership sweep were OUT OF SCOPE at open**, declared in
this iter's `overview.md` and routed to iter-86 with the adjudication ledger as their work list.

That is the lesson of iter-83 applied to itself: **a repair I cannot finish reproduces iter-81 exactly.**
Declaring a scope I can complete, and grading against it, is the point. `closed-fixed` here grades
planned scope, not the whole residual.

## State at close

- **Gate 4 of 5, unchanged.** No reading taken; clause 5 moves only on one that returns zero.
- 6 corpus guards exit 0. `repair_leak_guard` **GREEN** (1 waiver, reported).
- `repair_reach_guard` on this iter's ledger: **11/11, 0 unreached.**
- Zero platform-repo edits. `storage.md:55,:154,:181` **held**. `ai_architecture.md:225` **not touched**
  (adjudicated CORRECT — `D-M257x-84-5`).

## Close — 2026-08-05

**Outcome:** Q2's 7 deleted-fact claims + the 2 confirmed leak sites repaired at **100 % reach**, and the
live `dev-stack` default-profile defect fixed **by derivation** with both directions proven.
**Type:** tik
**Status:** closed-fixed — every declared line landed
**Gate:** NOT MET — 4 of 5, unchanged
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D-M257x-85-1 … D-M257x-85-3
**Side-deliverables:** none.
**Routes carried forward:** `FIX-M257x-iter86-repair-Q1-Q3-Q4-Q5` (29 upheld) ·
`FIX-M257x-iter86-p4-membership` (≥16 corpus/skills + 7 rext) ·
`CHECK-M257x-iter84-rule33-currency-amendment` · `CHECK-M257x-iter84-ground-truth-needs-origin-sha` ·
`CHECK-M257x-iter84-defects-outside-clause5-scope` ·
`CHECK-M257x-iter83-standalone-is-the-forgettable-class` · `CHECK-M257x-iter83-recall-lift-options`
**Lessons:** the reach fence paid for itself on its first use — but the *leak* fence is what caught the
defect I actually induced, and it caught it in the file where I had just written the rule about it.
