---
milestone: M257x
iter: 56
iteration_type: tik
status: closed-fixed
created: 2026-08-03
refs:
  platform: 0dab54dfac6beacdef54a671e2500d3940fd7329   # origin/main, re-fetched at iter open (P3)
  app: v1.365.0 (bff61c91)                             # == app origin/main; advanced this iter from v1.363.2
  rext: fast-build-m257x-iter-55b                      # at open; re-tagged at close
  rosetta: 5c9c099cdbdc576e431ce004f3b0e48197817fb8     # at open
  instrument: stack-injection/platform_topology.py     # extended this iter
---

# iter-56 — restore clauses 1 and 2 against origin HEAD

**Active strategy:** `TOK-04` (*pin the target, or stop calling it a measurement*) — P1 (every measurement
states its refs), P2 (every instrument is a committed file), P3 (the platform ref re-checked at open **and**
close, and the detecting iteration re-points in that iteration), P4 (derive, else fence, else declare
prose-under-review).

## Step 0 — re-survey before targeting (mandatory)

TOK-04's `Next-tik direction` names iter-56 as *"the 81-site sweep as one derived class"*. **That target is
stale and is deliberately substituted.** iter-55 closed `exit-4 (user-blocker)` on a pin decision that
blocks clauses 1 **and** 2; the orchestrator resolved it under the user's standing delegation and directed
this iteration at the ref baseline instead. Clause 5's prose sweep cannot honestly be measured while the
gate's own instrument — a cold bring-up — does not reach a verdict. Substitution recorded per Phase 1
Step 0; the TOK-04 strategy itself is unchanged and every rule below is one of its four.

**Platform ref at open:** `git fetch origin main` → `0dab54d`, clone already level. No re-point needed
this iteration (iters 54 and 55 each had to re-point inside themselves; this is the first in three that
opened level).

## Cluster / target identified

Three things, in dependency order:

1. `FIX-M257x-iter55-app-pin-lag` — advance the demo's `app` pin, recording what the advance contains
   (§7 rule 4).
2. **Clause 2** — the full Playthrough suite, binding, on a stack built at the current refs.
3. **Clause 1** — three consecutive cold `demo-down --purge` + `demo-up` cycles at
   `autoverify green:true / 0 warnings`, each carrying a P1 `refs:` block, each logging **which path the
   nondeterministic Directus bootstrap took**.

## Hypothesis — and it contradicts the premise this iteration was handed

The orchestrator's pin decision rests on iter-55's root cause: *"compose at `0dab54d` deleted
`STORAGE_RPC_ADDR`; pinned `app v1.363.2` still reads it at `main.go:446/516/983`, so `backend` exits 0 in
silence. The compose half of the storage fold landed; the app half is not in the pinned release."*

**Measured at iter open, before any repair — that root cause is false, and the pin advance cannot fix it:**

| measurement | result |
|---|---|
| `STORAGE_RPC_ADDR` at `v1.365.0` | still read at `main.go:450`, `:520`, `:988` + `internal/jobsimwiring/wiring.go:115` — the same three sites, shifted 4–5 lines |
| `app` origin/main | **is** `v1.365.0` (`bff61c91`); `git rev-list --count v1.365.0..origin/main` = **0** |
| the v9.0 app half | not released anywhere — the only v9.0 commits in the range are `docs(plan):` design commits |

So the app half of the storage fold does not exist at any ref we can build, and advancing the pin was never
going to restore `STORAGE_RPC_ADDR`. Something else killed `backend`.

**The real cause, measured by direct experiment on the same image and network** (the container from
iter-55's cycle A was still on the host):

| run | result |
|---|---|
| the container's exact env, no mounts | **starts** — 93 log lines, `Web server started at :8082`, alive >2 min |
| the same env + the `$HOME/.aws/credentials` bind mount as it exists on this host | **exit 0 in 0 s, 2 log lines** — the exact stack signature |
| the same env + a regular **empty file** at that path | **starts**, still up at 25 s |

`~/.aws/credentials` **does not exist on this Mac.** Docker does not fail on a missing bind source — it
creates it as an empty **directory** and mounts that (`~/.aws` and `~/.aws/credentials` were both created
at 16:12 today, by compose). The app's AWS config load then dies with `read /root/.aws/credentials: is a
directory` and the process exits **0**. This is `platform-alignment.md` §5 **Trap E** — *the tooling's own
host preconditions are invisible until a clean host* — on the new Mac, and it is a **host** defect, not a
platform-version skew.

## What is done about it

- **The host precondition becomes a derived, fenced pre-flight** (P4 tier 1 → tier 2), extending iter-55's
  `platform_topology.py`: host-absolute bind-mount sources of the **default profile's** services are
  derived from the platform's own compose and checked before compose runs. Workspace-relative sources
  (`./data/…`) are excluded by a real property, not a hand list.
- **The check's shape is forced by the measurement.** The residue of Docker's auto-creation is a path that
  **exists** as an **empty directory**, so an existence-only assert reports GREEN over the exact host state
  that produced the defect. It is watched going RED on that real state before it is trusted.

## Pre-registration — stated before any run, so it can be refuted

iter-55's two pre-registrations were both refuted. These are not softened.

- **PR-1 — the pin advance does NOT fix the backend exit.** Repairing the mount fixes it; the pin is
  irrelevant to it. Corollary: iter-55's cycle A would have gone green at `v1.363.2` with the mount
  repaired.
- **PR-2 — the advance is seeder-safe.** Recorded per §7 rule 4: 37 commits; **two** new migrations, both
  `ALTER TABLE … ADD COLUMN` with a default or NULL (`course_builder_sessions.brief`/`credits_spent`,
  `academy_chapter_progresses.completed_at`); **0** `DROP TABLE` / `DROP COLUMN` / `RENAME`; **0** new
  `log.Fatalf` in any non-test Go file, so the advance cannot introduce a new hard-required-config boot
  failure. Prediction: **the seeders survive unchanged.** This is the class that broke at v2.1 and v2.7,
  so it is watched rather than assumed.
- **PR-3 — clause 2 comes back `30 passing / 0 failing / 0 error`** on a stack built at these refs.
- **PR-4 — clause 1: cycle A goes green.** Deliberately NOT predicting 3/3: the Directus bootstrap race is
  nondeterministic, so 3 consecutive greens is a claim about a distribution, not about a mechanism. Each
  cycle logs which path it took.

## Escalation conditions

- A seeder failure after the advance → diagnose and fix (the known class), **never revert the pin**.
- A red cycle whose cause is a *platform* inconsistency at origin HEAD → that is a real finding and the
  highest-value output available; record it, do not paper over it.
- Clause 5 is untouched and is **not** re-cut. The user has ruled three times.

## Acceptable close-no-lift outcomes

A refuted pre-registration is a first-class result. If clause 1 cannot reach 3/3 because the platform's own
refs are mutually inconsistent, the honest gate reading is the deliverable.
