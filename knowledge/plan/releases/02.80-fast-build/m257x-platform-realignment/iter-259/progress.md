# iter-259 — progress

**Type:** tik
**Opened:** 2026-08-12T14:12:42Z

## Phase A — seal

Pre-registrations sealed before any measurement. Read-only iter by construction: nothing under `stack-dev/` is written.

## Phase B — measurement (read-only; nothing under `stack-dev/` was written)

### PR-1 — **HELD.** `make init` adopts rather than clones

`platform/Makefile:19-25` @ `0c91421`:

```make
init: ## Clone all repos from repos.yml that don't exist in ../
	@for repo in $(REPO_NAMES); do \
		if [ ! -d "$(PARENT_DIR)/$$repo" ]; then \
			git clone git@github.com:$(ORG)/$$repo.git "$(PARENT_DIR)/$$repo"; \
		else \
```

The `else` branch is a skip. iter-258's own demo log printed the behaviour verbatim
(*"studio-desk already exists, skipping"*). So a dev bring-up in `stack-dev/` does not *clone* a clean
`studio-desk` — it **adopts whatever is sitting there** and treats it as a platform build source.

### PR-5 — **SPLIT: refuted in its branch clause, held in its substance (and worse than predicted there)**

| measurement | reading |
|---|---|
| branch | `main` @ `795a411d` |
| commits on **no remote ref** (`--all --not --remotes`) | **367** |
| `origin` push URL | **`no-push://demo-clone-never-pushes`** — pushing is structurally blocked |
| other branches present | `release/3.0-ground-works`, `release/3.1-open-frame`, `release/3.2-full-frame`, `fix/guided-mode` |
| **linked worktree** | **`stack-dev/.worktrees/studio-desk-feat-stack-migration` @ `411a3c15 [release/3.2-full-frame]`** |
| working tree | clean |

**The linked worktree is the fact that settles the iter.** `stack-dev/` does not merely *store* another
project's clone — it hosts an **active worktree of that project's release branch**, i.e. in-flight work
with a second checkout pointing into the same object store. Those 367 commits can be pushed nowhere:
the push remote is deliberately disarmed.

> **A note on denominators, so this is not read as a contradiction.** The handoff records *"463 commits
> of this release exist on no remote at all."* This iter measures **367**, which is a *different
> question* — commits on no remote **ref** across all branches of **this clone**, today. Neither figure
> refutes the other and no reconciliation is claimed; both say the same load-bearing thing, which is
> that the work is irreplaceable.

### PR-2 — **HELD.** One env var does not relocate a dev stack

`stack-dev` is rooted in **two different scripts behind two different variables**:

- `dev-stack:53` — `PLATFORM_DIR="${PLATFORM_DIR:-$REPO_ROOT/stack-dev/platform}"`
- `migrate-dev.sh:19` — `DEV="${DEV_CLONES:-$REPO_ROOT/stack-dev}"`

So relocating a dev stack means overriding **`PLATFORM_DIR` *and* `DEV_CLONES`** in concert — an
invocation shape no skill drives and no test covers.

### PR-3 — **HELD, and it is a corpus-vs-tooling divergence in this milestone's own class**

A search of `dev-stack/` **and** `.claude/skills/dev-up/SKILL.md` for any `stack-dev-N` form returns
**zero**. `dev-N` is a **port offset over one shared clone set**, not a second workspace.

**But `CLAUDE.md`'s workspace table lists `stack-dev-2/` as "a secondary dev stack"** — describing a
thing the tooling does not implement. This is exactly the class M257x exists for: a corpus statement no
mechanism backs. Routed rather than repaired here, because repairing it is a corpus edit and this iter's
subject is the dev half's feasibility. Handler: `FIX-M257x-259-stack-dev-N-is-not-implemented`.

### PR-4 — **HELD.** First free N is 3

The unified registry now holds `demo-1` (n=1, up) and `demo-2` (n=2, up, `created 2026-08-10T14:04:46Z`
— iter-258's). First free N across both kinds is **3**.

## Phase C — the decision, and why it is the user's

**4 clean holds + 1 SPLIT — and the split is my own wording, caught while grading.** (Trend:
… → 2/5 → 4/5 → **4½/5**, booked as **4 of 5** rather than rounded up.)

PR-5 predicted *"the occupying `studio-desk` is on a **non-`main`** branch carrying commits that are on
no remote."* **The branch clause is REFUTED — it is on `main`.** The substantive clause HELD and then
some: 367 commits on no remote ref, a `no-push://` origin, and a live worktree on
`release/3.2-full-frame`. Booked as a split rather than as a hold, because the prediction bundled two
claims and only one survived — and calling that "5 of 5" would be the *convenient* reading, which is the
class `D-M257x-258-3` was written about **one iter ago**. The load-bearing half is the half that held,
and the risk is unchanged; the score is not.

Standing a platform dev stack up on this box, as the tooling is written, requires `stack-dev/` — and
`stack-dev/` is another project's active workspace with 367 unpushable commits and a live worktree. The
three candidate paths were each examined rather than dismissed:

1. **Run `/dev-up` as documented** → `make init` adopts their `studio-desk` as a build source (PR-1).
   Their tree becomes a Docker build context, and any step that writes to a build source touches work
   that exists nowhere else. **Rejected — the downside is unrecoverable and not mine to risk.**
2. **Relocate via `PLATFORM_DIR` + `DEV_CLONES`** → possible in principle (PR-2), but it is an
   unsanctioned invocation no skill drives and no test covers, and it needs a **fresh full clone set**
   (`app` included) on a contended host. **Not rejected on effort — rejected because inventing a
   non-standard dev-stack topology unilaterally, to satisfy a gate, is precisely the kind of move whose
   result nobody could later trust.**
3. **Move the other project's work aside** → **not this agent's call under any reading.**

**This is a `user-blocker` by Phase 5 § 4's own definition** — a question whose answer changes what may
land — and it is escalated *before* acting rather than reported after. The three-fate rule was applied
first and did not yield a Fate-1 landing: there is no version of this that completes safely without a
decision that belongs to the user.

**What the user is being asked** (a short question, not a research task):

> The dev half of the closing condition needs a platform dev stack, and the tooling puts it in
> `stack-dev/` — which currently holds the studio-desk v3.2 migration (367 commits on no remote, plus a
> live worktree). Three ways forward: **(a)** move/park that work first, **(b)** authorise a
> relocated dev stack at a fresh path via `PLATFORM_DIR`+`DEV_CLONES`, or **(c)** rule that the dev half
> is satisfied elsewhere / deferred. **Nothing in `stack-dev/` has been touched.**

## Close — 2026-08-10

**Outcome:** The dev half is **blocked on a decision that is the user's, and the block is measured
rather than suspected**: `stack-dev/` is another project's active workspace (367 commits on no remote,
a `no-push://` origin, and a live worktree on `release/3.2-full-frame`), and `make init` is
skip-if-present, so a documented `/dev-up` would silently adopt their tree as a platform build source.
**4 of 5 pre-registrations held, PR-5 split** (its branch clause refuted — the clone is on `main` —
its substance held). Nothing under `stack-dev/` was written.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: **y** — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: exit-4

**Decisions:** `D-M257x-259-1` (the dev half is user-blocked; three paths named, none taken).

**Side-deliverables:** none — read-only iter.

**Routes carried forward:**
- `ROUTE-M257x-258-no-dev-stack-on-this-box` → **open, now with a measured cause and a decision request**
  rather than an observation.
- `FIX-M257x-259-stack-dev-N-is-not-implemented` → **new.** `CLAUDE.md` advertises `stack-dev-2/` as a
  secondary dev stack; the tooling has no per-N dev workspace at all. A corpus claim no mechanism backs.
- `ROUTE-M257x-258-the-pin-is-157-iters-stale`, `ROUTE-M257x-257-lock-file-is-unfenced`,
  `ROUTE-M257x-256-mixed-ref-anchors` and all earlier → unchanged and open.

**Lessons:**
1. **Check what occupies a path before running the tool that writes to it.** The whole cost of this iter
   was a handful of read-only git commands; the cost of skipping it was another project's unpushable
   work becoming a Docker build context.
2. **`make init`'s skip-if-present is an ADOPTION, not a no-op.** It reads as a convenience ("already
   there, nothing to do") and it silently promotes whatever is on disk to a build input.
3. **An escalation is a deliverable when the alternatives are all irreversible or unsanctioned.** The
   three-fate rule asks whether the item can land now; here the honest answer was no, and inventing a
   fourth path to avoid asking would have been the failure, not the diligence.
