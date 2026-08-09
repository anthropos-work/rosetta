---
iter: 222
milestone: M257x
iteration_type: tik
status: open
created: 2026-08-09
---

# iter-222 — the tooling ships a manifest declaring what the platform IS, and nothing checks it against the platform

**Active strategy:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop sampling them.*

## USER REDIRECT — recorded here because it changes target selection from this iter forward

Verbatim, 2026-08-09: *"the goal remains alignment and be able to build a working stack with the new
platform repos (only the remaining ones that are still part of it)."*

Booked as `D-M257x-222-1` in this iter's `decisions.md`. The operative consequence: iters target **the
platform and the corpus's/tooling's claims about it**, not the instruments that grade those claims. The
last ~60 iters were instrument work. This one is not.

## Step 0 — re-survey (mandatory)

`TOK-08` names no next-class beyond *"work the classes in descending measured size"*, so the re-survey is
the whole target selection. Re-surveyed at HEAD against the redirect:

1. **Platform `origin/main` is `0c91421`** — the *same* commit gate clauses 1+2 were proven at, 0 commits
   behind. The `platform` orchestrator repo has **not** moved. (The run brief assumed it had; it has not.
   The clones underneath it have.)
2. **`repos.yml` @ that ref lists FOUR names** — `app`, `sentinel`, `next-web-app`, `studio-desk`. That is
   the user's *"only the remaining ones that are still part of it"*, verbatim and machine-readable.
3. **`rosetta-extensions` ships a canonical `demo-stack/clones.pin.json`** that `ensure-clones.sh` seeds
   into every fresh `stack-demo/` workspace copy-if-absent (M246). Its own comment calls it *"the
   barrier's reproducible **current-origin/main** topology"*. **It is the tooling's declaration of what
   the platform is** — the artifact a cold bring-up on a fresh box reads to decide what to check out.
4. **Nothing reads it.** `grep` across every `.py` in `rosetta-extensions`: the only non-`test_tooling`
   hit is a doc-comment mention. `test_tooling.py` writes *synthetic* pins to test the advance behaviour;
   **no fence asserts the real file's contents.**
5. `clone_drift_guard` (FENCE-M257x-iter106) is the nearest neighbour and **does not reach this**: D1
   asserts *corpus-vs-clone* (has the clone advanced past every sha the corpus cites). A clone frozen
   **behind origin** reads D1-green — the clone did not move, so nothing is stale relative to it. The
   axis *clone-vs-origin* is measured by no fence in either repo.

## Cluster / target identified

The canonical `clones.pin.json`, on **two mechanically-decidable axes**:

- **A — membership.** Which repos may the pin name at all? Derivable from `repos.yml` @ platform
  `origin/main` plus the two sanctioned non-`repos.yml` clones (`platform` itself; `ant-academy`, the
  explicit M49 #5 clone).
- **B — freshness.** Is each pinned sha at its repo's `origin/main`? Derivable with `git rev-list --count`.

Both are censuses, not samples. Neither has ever been run.

## Hypothesis

The pin over-declares (names repos the platform deleted from its clone set) **and** under-declares
(pins live repos behind origin), and the comment asserting *"current-origin/main"* is a
`§8`-iter-208 decaying clause — it named a **run** and is read as a **property**.

## Pre-registered numeric claims — SEALED in this iter's first commit, before any repair

Derived 2026-08-09 from `repos.yml` @ platform `origin/main` = `0c91421` and from `git rev-list` against
each clone's freshly-fetched `origin/main`.

| # | claim | value |
|---|---|---|
| **P1** | keys in canonical `demo-stack/clones.pin.json` | **11** |
| **P2** | names in `repos.yml` @ platform `origin/main` (`0c91421`) | **4** — app, sentinel, next-web-app, studio-desk |
| **P3** | sanctioned non-`repos.yml` clones | **2** — `platform` (self), `ant-academy` (M49 #5) → legitimate pin population **6** |
| **P4** | **phantom** pin keys (named by the pin, absent from `repos.yml` @ origin) | **5** — cms, jobsimulation, storage, messenger, roadrunner |
| **P5** | of the **6** legitimate pins, how many are behind their repo's `origin/main` | **3** — app **28**, next-web-app **12**, ant-academy **9**; platform/sentinel/studio-desk **0** |
| **P6** | fences in either repo asserting pin membership against `repos.yml` | **0** |

**Scope caveat stated up front (`§5` rule 46):** P5's three `0`s and three deltas were measured after an
explicit `git fetch origin` in this session for `platform`, `app`, `sentinel`, `next-web-app`,
`studio-desk`, `ant-academy`. The five phantom repos were **not** re-fetched; no freshness number is
claimed for them and none is needed — membership, not freshness, is their finding.

## Expected lift

`P4 → 0` by repair, and a fence that keeps it at 0 by enumerating the population on every run. Axis B is
a **disclosure**, not an auto-advance: bumping the pin is an operator decision that would invalidate the
M246 proven-topology barrier without a re-proof of gate clauses 1+2. The repair on B is to stop the
comment **asserting** currency and make the number **derivable** instead.

## Phase plan

1. Seal the probe (this file + measured evidence) as commit 1 — before any repair.
2. Repair axis A: drop the 5 phantom keys; correct the `ensure-clones.sh` comment that asserts
   *"jobsimulation stays standalone"* (refuted by `repos.yml` @ origin and by the corpus's own fenced
   migration map).
3. Ship the fence: enumerate the pin's population, assert membership against `repos.yml`, with a
   mutation control and an anti-vacuity control that can actually fire.
4. Repair axis B's decaying clause; record the measured freshness where it cannot be read as a property.
5. Re-measure; close.

## Escalation conditions

- If dropping a phantom key changes bring-up behaviour on a **fresh** box → user-blocker (it would mean
  the pin is load-bearing for a repo the platform says is not part of it).
- If `repos.yml` @ origin moves mid-iter → re-scope trigger arithmetic applies.

## Acceptable close-no-lift outcomes

If the census returns **P4 = 0** — i.e. the phantom reading is my own error — that falsification closes
the iter honestly. A census returning zero must prove its instrument (`§9` iter-149).
