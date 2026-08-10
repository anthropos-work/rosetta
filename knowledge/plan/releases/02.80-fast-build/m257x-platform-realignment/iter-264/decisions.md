# iter-264 — decisions

## `D-M257x-264-1` — the dependency moved from `cms` to `app`; its documentation stayed behind

The corpus **already contained** an exact statement of the failure iter-262 hit by running into it:

> `corpus/services/cms.md:271` — *"The Python studio submodule **had to be** cloned **before** any docker
> build, otherwise `make up` failed with `"/studio": not found`"*

Past tense, filed under a **decommissioned** service. `staging-bringup.md:428` compounds it by describing
the same COPY lines as a **cms quirk to comment OUT**. When `fdb8034a` (2026-07-27) gave `app` the same
embedded Python runtime, **the requirement migrated and the documentation did not** — so the live
constraint survives only as dead history about a service that no longer has a container, a `repos.yml`
entry, or an ECS service.

**This is `platform-alignment.md` §5's own rule — "a named-consumer list survives the merge that moved the
consumer" (M257x iter-23) — occurring inside this corpus rather than inside the platform's.** The class is
general and worth a sweep: **when a service is folded into `app`, its build/runtime PRECONDITIONS move
with it, and every sentence documenting them keeps pointing at the merged-away service.**

**Corrected in this iter:** `setup_guide.md` (section rewritten to state the requirement, present tense,
before `make up`, plus an inline warning at the `make up` step) and `CLAUDE.md` (whose Studio-Room block
asserted the **inverse** — that a *build* populates `app/studio`, when a local build **fails** on its
absence).

**Left open, deliberately:** `cms.md:271` and `staging-bringup.md:428` still describe the live `app`
dependency as `cms` history. Fixing those two sentences is easy; **finding the rest of the class is the
work**, and it is routed as `FIX-M257x-264-cms-md-past-tense-dependency` rather than patched piecemeal
here — this milestone's own lesson that a class is closed by an enumeration, never by its last member
(`platform-alignment.md` §8, M257x iter-176).
