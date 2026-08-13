# iter-259 — decisions

## `D-M257x-259-1`: the DEV half is USER-BLOCKED — three paths named, none taken

**Measured, not suspected.** The tooling puts a platform dev stack in `stack-dev/`
(`dev-stack:53`, `migrate-dev.sh:19` — two scripts, two variables), and `stack-dev/` is another
project's active workspace:

- `studio-desk` @ `795a411d` on `main`, **367 commits on no remote ref** across all branches
- `origin` push URL is **`no-push://demo-clone-never-pushes`** — pushing is structurally disarmed
- a **linked worktree** at `stack-dev/.worktrees/studio-desk-feat-stack-migration` @ `411a3c15`
  `[release/3.2-full-frame]` — live, in-flight work

And `platform/Makefile:22` clones only `if [ ! -d ... ]`, so a documented `/dev-up` would **adopt** that
tree as a platform build source rather than clone a clean one.

**Three paths, each examined:**

1. **Run `/dev-up` as documented** — rejected. Their tree becomes a build context; the downside is
   unrecoverable.
2. **Relocate via `PLATFORM_DIR` + `DEV_CLONES`** — mechanically possible, but unsanctioned, driven by
   no skill, covered by no test, and needing a fresh full clone set. Rejected **not on effort** but
   because inventing a non-standard dev topology to satisfy a gate produces a result nobody could later
   trust.
3. **Move the other project's work aside** — not this agent's call.

**Escalated before acting, not reported after.** The three-fate rule was applied and yielded no Fate-1
landing. **Nothing under `stack-dev/` was written, cloned, checked out or built.**

The question put to the user is short and decidable: **(a)** park the studio-desk work first,
**(b)** authorise a relocated dev stack at a fresh path, or **(c)** rule the dev half satisfied
elsewhere or deferred.

> **Why this is not over-caution.** The demo half took one bring-up to prove because nothing was in its
> way — slot 2 was free and `demo-1` was untouchable but irrelevant to it. The dev half has a live
> obstruction, and the obstruction is *someone's only copy*. The asymmetry between the two halves is the
> finding, and it is why "just run the other one too" was never the next step.

## `D-M257x-259-2`: PR-5 bundled two claims and only one survived — booked as a SPLIT, not a hold

PR-5 read *"the occupying `studio-desk` is on a **non-`main`** branch carrying commits that are on no
remote."* Measured: the clone **is on `main`** (`795a411d`), so the branch clause is **refuted**; the
substantive clause held and then some (367 commits on no remote ref, `no-push://` origin, a live
worktree on `release/3.2-full-frame`).

**Booked as 4 of 5 with a split, not 5 of 5.** A bundled pre-registration is graded by its weakest
surviving clause, because the alternative is a scoring rule that rewards vagueness: a prediction naming
two things and scoring a hold when one is wrong is a prediction that cannot fail.

This is `D-M257x-258-3`'s class arriving **one iter later, in my own grading rather than in my own
prose** — the convenient reading of my own evidence. The risk it describes is unchanged; only the score
is. Rule: **write pre-registrations as single claims; if one bundles, grade the bundle down.**

---

**⚠️ CORRECTED by iter-260 (2026-08-10).** The `367 commits on no remote ref` figure above is
**retracted** — the clone has **zero `origin/*` remote-tracking refs**, so `--not --remotes` subtracted a
local **bundle** and the number is *"commits not in a courier file."* The active worktree's branch
`release/3.2-full-frame` @ `411a3c15` **is on origin** (`ls-remote`, exact sha match); nothing on it is at
risk. The escalation's **conclusion survives on correctness grounds** — local `main` `795a411d` is stale
against origin `41ee3575`. Full retraction: `iter-259/progress.md` § RETRACTION.
