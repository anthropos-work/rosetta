---
iter: 259
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-10T14:12:42Z
closed: 2026-08-10T14:15:20Z
---

# iter-259 — the DEV half: can it be built here at all, and at whose expense?

**Type:** tik
**Active strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07),
under the user's binding `D-M257x-256-1`.

## Step 0 — Re-survey before targeting

iter-258 discharged the **demo** half of `ROUTE-M257x-256-the-advance-is-unproven` and opened
`ROUTE-M257x-258-no-dev-stack-on-this-box`. The user's closing condition needs **both** halves, so the
dev half is now the critical path — and the obvious next move (*"just run `/dev-up`"*) is exactly the
move this iter exists to check **before** running it, not after.

Two facts already in hand from the re-survey make that check non-optional:

1. **`dev-stack:53` hard-codes one workspace** — `PLATFORM_DIR="${PLATFORM_DIR:-$REPO_ROOT/stack-dev/platform}"`.
   There is no per-N dev workspace; `dev-N` is a **port offset**, not a second clone set.
2. **`stack-dev/` is already occupied by a different project** — `studio-desk`, `studio-room` and
   `HANDOFF-ant-mini.md`, from a v3.2 "full frame" migration whose handoff note records
   *"**463 commits of this release exist on no remote at all**"*.

And `repos.yml` @ platform `0c91421` clones exactly four repos — **`app`, `sentinel`, `next-web-app`,
`studio-desk`**. The fourth is the collision.

## Cluster / target identified

`ROUTE-M257x-258-no-dev-stack-on-this-box`. The question is not *"is a dev stack green?"* — it is the
prior one: **can a dev stack be stood up on this box without touching work that exists on no remote?**
An iter that answered the second question by running the build and finding out would be the wrong
instrument, because the cost of being wrong is another project's unpushed commits.

## Hypothesis

`make init` is **skip-if-present** (iter-258's own demo log printed *"studio-desk already exists,
skipping"* for exactly this reason), so a dev bring-up in `stack-dev/` would **silently adopt** the other
project's `studio-desk` working tree as a platform build source — building it, and exposing it to any
step that writes to a build source.

## Expected lift

Not metric movement. The deliverable is a **decision with evidence**: either a safe path to the dev half
that this iter can name and justify, or a `user-blocker` escalation with the hazard measured rather than
suspected. Under the three-fate rule this is Fate 1 — the investigation completes and lands its verdict.

## Pre-registrations — sealed in this iter's FIRST commit, before measuring

| | claim | prediction |
|---|---|---|
| PR-1 | `make init` in `stack-dev/platform` would SKIP `studio-desk` and adopt the other project's tree | **HOLDS** — iter-258's demo log shows the skip-if-present behaviour verbatim |
| PR-2 | `PLATFORM_DIR` alone cannot relocate a dev stack — the dev path is `stack-dev`-rooted in **more than one** place | **HOLDS** |
| PR-3 | the tooling has **no** per-N dev workspace (`stack-dev-N` is not a thing the code knows) | **HOLDS** — despite `CLAUDE.md` listing `stack-dev-2/` in its workspace table |
| PR-4 | the unified registry's first free N is **3** (demo-1 and demo-2 hold 1 and 2) | **HOLDS** |
| PR-5 | the occupying `studio-desk` is on a non-`main` branch carrying commits that are on no remote | **HOLDS** — the handoff says 463; the measurement is whether *this clone* still shows them |

## Phase plan

- **Phase A** — seal. **Phase B** — measure PR-1…PR-5 read-only. **No write, no clone, no bring-up in
  `stack-dev/`.** **Phase C** — decide: safe path, or escalate. **Phase D** — close.

## Escalation conditions

If the only way to the dev half runs through another project's unpushed work, that is a **`user-blocker`**
by Phase 5 § 4's own definition — a question whose answer changes what this iter may do — and it is the
user's call, not this agent's. **Nothing in `stack-dev/` is written to under any outcome.**

## Acceptable close-no-lift outcomes

A measured *"the dev half is not safely buildable on this box as configured"* is a complete iter. The
route asked whether a dev stack exists; establishing that standing one up has a named, cited cost
discharges it as fully as a green stack would.
