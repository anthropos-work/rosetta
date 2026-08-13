# M258 iter-11 — decisions

## D53 — `docker images` SIZE is not reclaimable size. Quote `system df`.

The four `m257-*:probe` leftovers sum to **8.88 GB** in the SIZE column and reclaim nothing like it:
they share **5 of 10 layers** with `demo-2-next-web`, which is in use. `system df` reports **1.754 GB**
reclaimable across *all* unreferenced images and that is the figure to quote. A cleanup planned off the
SIZE column would have promised ~8 GB and delivered under 2 — and then been blamed on the cleanup.

## D54 — Prove ownership by intersection before deleting anything.

All 184 volumes were **anonymous** (no `com.docker.compose.project` label), so nothing about a name
identified its owner. The check that made the reclaim safe was mechanical: enumerate every volume every
container mounts, intersect with the dangling set, require **0**. It returned 0, and the six in-use
volumes resolved to `demo-1-postgresql-1`, `demo-2-postgresql-1` and `anthropos-postgresql-1` — both of
the user's stacks included. **5.297 GB reclaimed; all three stacks verified resident afterwards.**

## D55 — The producer: two undeclared anonymous volumes per Postgres container.

The bitnami image declares three `VOLUME`s; compose binds only `/bitnami/postgresql` (to the host stack
dir). `/docker-entrypoint-initdb.d` and `/docker-entrypoint-preinitdb.d` are left **anonymous**, so
every container start mints two and every non-`-v` teardown or container *recreate* orphans them.
Three stacks over five days → 178. Fix is time-neutral: declare both in compose, or make every
teardown path pass `-v`.

## D56 — `--purge` is innocent, and it was measured rather than assumed.

`cmd_down --purge` runs `down -v --remove-orphans` (`rosetta-demo:413`); the plain path (`:446`) does
not. The decisive evidence is temporal: today's 3-rep campaign ran three `--purge` teardowns between
08:19 and 09:02Z and produced **zero** orphans — the newest orphan is from **03:00Z**. The leak is the
non-`--purge` path and container recreation.

## D57 — SPACE is a new GOAL of this milestone, and the end state is ONE stack. (User ruling.)

Recorded verbatim in effect: the remaining budget goes to a last attack on build time, **then a TOK for
space optimisation as a new goal**, and *"by end of this milestone there is only one stack up, and it's
built with the new process/mechanism and the newest repos of the platform."*

Two consequences that must not be softened:

1. **Space gets a strategy (`TOK-02`), not a side route.**
2. **`END-M258-one-stack` is a hard requirement**, and it **supersedes the standing "never touch
   `demo-2`" rule — only at the end, and only in this order: build and verify the new stack FIRST,
   then tear the others down.** Never teardown-first; the user must never be without a stack to
   validate on. `demo-2` is explicitly *not* the one to keep — it is on **pre-L1** images
   (next-web 4.04 GB / hiring 3.94 GB, against 0.80 GB for the post-L1 pair), so the user's own stack
   never received the release's biggest win.

## D58 — Space must never be bought with time. Cache pruning is not a default.

The user named the tension himself (*"account for this on the cache consideration"*). The build cache
is **27.45 GB with 19.28 GB reclaimable** — by far the largest single reclaim available, and by far the
most expensive: `build-budget.md` already records that a single 356.8 MB eviction cost **173 s**.

So: **deleting build cache to win space is forbidden as a default**, `--filter until=24h` never `-af`,
and **any cache policy must be argued on BOTH axes with measurements.** Under that constraint the
attractive space wins are the ones that cost no time at all — orphaned volumes (`D55`), leftover
host-side stack dirs, and dead images — which is why none of them were traded against the cache here.
