---
milestone: M257x
iter: 03
---

# iter-03 — decisions

No new milestone-level decisions. Two records worth keeping:

## Live corroboration of `D-M257x-3` (the per-environment axis)

`D-M257x-3` split every map row into a **production** state and a **fresh-local-stack** state, on the
argument that collapsing them is why this class recurred three times. The clone set bootstrapped this iter
shows the split directly: **`cms`, `jobsimulation` and `roadrunner` are all cloned** into `stack-demo/` —
they are still in `repos.yml` as the rollback reference until M810 — while `repos.yml` declares **no schema
for any of them** and the derived migration set is `app:public` alone.

So on one fresh local stack, at one sha: *cloned* = true, *owns a schema* = false, *runs in the default
compose profile* = true, *federated* = false. Four different answers about the same three services. Any map
that carried one state per row would have to pick one and be wrong about the others.

## The clone set is NOT evidence about the bring-up

Recorded so a later reader does not over-read this iter. A complete clone set proves the **source** is
obtainable. It says nothing about whether the images build, the migrations apply, the seed lands, or
`autoverify` goes green — and this milestone exists because a total breakage sat invisible for four days
behind exactly that kind of partial-signal confidence (M257 B1/B2; 2,617 offline tests green against a fake
`Conn` while the bring-up was dead).

**Gate clause 1 is unchanged at NOT MET, and no part of it has been attempted.**
