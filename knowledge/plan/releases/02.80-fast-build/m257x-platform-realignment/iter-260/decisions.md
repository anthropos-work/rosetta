# iter-260 — decisions

## `D-M257x-260-1` — iter-259's `367` is retracted; the escalation survives on a different ground

**Side-deliverable, not planned scope.** Raised by the orchestrator mid-iter after the user challenged the
brief's `803`. iter-260 re-measured rather than relaying, and the re-measurement went **further than the
correction it was given**.

**Measured (read-only; nothing under `stack-dev/` was written):**

- `stack-dev/studio-desk` has **zero `origin/*` remote-tracking refs**; all **11** are `bundle/*`. So
  iter-259's `git log --all --not --remotes` subtracted a **local bundle file** — `367` is *"commits not in
  a courier artifact"*, published as *"commits on no remote ref."* The brief's `803`
  (`bundle/main..release/3.2-full-frame`) is the same family plus a main-vs-release divergence confusion.
- `git ls-remote --heads origin` → `refs/heads/release/3.2-full-frame` **`411a3c15`**, byte-identical to
  the local tip and to the live worktree's checkout. **The branch the alarm was about is published.**
- The clone does **not hold** origin's `main` object `41ee3575` (`git cat-file -e` fails).

**A zero I withdrew before publishing.** I computed *"0 commits absent from origin"* and the anti-vacuity
control **did not fire** — dropping a head from the `--not` set still returned `0`, because the command was
erroring against absent shas and `2>/dev/null | wc -l` rendered that as `0`. Void under §9. The other four
branches are **UNMEASURABLE-FROM-HERE**, which is a third verdict, not a clean bill.

**Decision:** publish **neither** `803` nor `367` as a risk figure, anywhere. Retract in place in
`iter-259/{progress,decisions}.md` rather than silently editing the number — the record must show the
figure was wrong **when written**, not overtaken.

**The prohibition on `stack-dev/` STANDS, with its reason replaced.** Not data loss — **staleness**: local
`main` `795a411d` (2026-07-30) vs origin `41ee3575`. A dev bring-up rooted there builds studio-desk from a
five-week-old tree, the exact stale-source failure this milestone exists to eliminate. `D-M257x-259-1`'s
escalation was the right call on a wrong number, and the user's challenge was correct.

**Generalised into the protocol** (`§8`, appended to *a guard that reads `origin/<branch>` measures whatever
the last FETCH left there*): **"cannot push from here" is not "exists only here."**
