**Type:** tik — under [`TOK-05`](../decisions.md); the repair half of iter-61's split.

# iter-62 — the prose class repaired, whole

## What landed

The 22-site / 12-file residual iter-61 enumerated is **repaired to guard-GREEN**. Two kinds of edit,
and the split matters:

**(a) The predicate itself — 11 mechanical sites.** `graphql` → `core` in the noun phrase, plus the
one-line reason at the two entry points a reader meets first (`CLAUDE.md`, `architecture_overview.md`)
so the rename is discoverable rather than merely applied.

**(b) A SECOND predicate riding in the same rows — 12 sites that needed more than a token swap.**
Several of the residual lines assert not only *a `graphql` profile* but *"the husk container still
starts"*, which is separately false at `0dab54d`:

| site | asserted | measured at `0dab54d` |
|---|---|---|
| `dependency_map.md:15` | cms husk *"still defined at `docker-compose.yml:144`… still answers messenger's RPC until M809"* | no cms service, no `repos.yml` entry, **M809 landed** |
| `dependency_map.md:16,18` | jobsim + roadrunner husks *"still start"* | neither is a compose service; `ROADRUNNER_RPC_ADDR` is set nowhere |
| `dependency_map.md:31` | *"They are NOT gone from compose or `repos.yml`"* | they are gone from **both** — 10 services, 6 repo entries |
| `service_taxonomy.md:97` | *"**YES** — container still starts"* | **NO** — gone from compose |
| `run_guide.md:88` | `make up` starts 10 services incl. *"GraphQL/Cosmo Router"* | **five**: postgresql, redis, sentinel, backend, gotenberg |
| `run_guide.md:203` | *"`make up PROFILE=studio-desk` starts Studio-Desk with its dependencies"* | **exits 1** — `depends_on: backend`, unselected |

**A predicate-scoped repair surfaces its neighbours.** Repairing by predicate does not mean editing
only the predicate: the sites that share one false assumption very often share a second, and the
sweep is the cheapest moment to see it — you are already in the sentence.

## What the fence still cannot see, and it grew teeth this iter

`service_taxonomy.md:55-67` — the **Services table** — is headed *"(current local docker-compose @
platform `2adcf71`)"*, so the ref-pin exemption covers every row in it. Those rows still list
Jobsimulation, CMS and Roadrunner as starting containers with ports. The pin makes them *historically
true and presently misleading*, and the guard is correct to exempt them by its own rule.

That is now **two** independent instances of `CHECK-M257x-iter60-stale-pin-exemption`
(`messenger.md:107-110`'s two stale RPC values being the first), which upgrades it from a noted hole
to the **next fence build**. The fix shape is known: a pin exempts only when the ref it names is the
ref the guard is pointed at; anything older must be *marked* as history rather than merely dated.

## Close — 2026-08-04

**Outcome:** the enumerated 22-site / 12-file prose class repaired **whole** to guard-GREEN, and 12 of
those sites carried a **second** false predicate (the husk containers) that the sweep caught and fixed
— including `make up`'s service list (10 → **five**) and a documented `PROFILE=studio-desk` that
exits 1. `markdown_structure_guard` clean at 112 files.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5, unchanged; clause 5 is still graded only by a reading that returns zero.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** `D-M257x-62-1` — a predicate-scoped repair surfaces its neighbours, and the sweep is
the cheapest moment to fix them; `D-M257x-62-2` — the ref-pin hole has a **second** instance
(`service_taxonomy.md`'s whole Services table) and is promoted to the next fence build.
**Side-deliverables:** none.
**Routes carried forward:**
- `CHECK-M257x-iter60-stale-pin-exemption` → **promoted to the next fence build.** Two instances now:
  `messenger.md:107-110` (two stale RPC values) and `service_taxonomy.md:55-67` (the whole Services
  table, still listing three husk containers with ports). Fix shape: a pin exempts only when its ref
  **is** the guard's ref.
- `FIX-M257x-iter58-mainline-shift` → next iter, with iter-61's refreshed measurement (5 of 16
  `app/main.go:N` citations still land on their claimed construct at app `v1.366.0`).
- `DOC-M257x-iter59-storage-mid-fold` → next iter: the map's 8th state token + assertion C. The
  *measurement* is already landed and G6-fenced in `storage.md`.
- `CHECK-M257x-iter60-g6-citation-subject` · `FIX-M257x-iter53-union-set` (**PENDING USER DECISION**) ·
  `FIX-M257x-iter56-assignment-flake` (**NOT DECIDED** — needs a failure *rate*) ·
  `CHECK-M257x-iter38-ai-act-classification` (needs an owner outside this milestone) ·
  `-cold-daemon-registry` · `-grep-vs-failclosed` · `-empty-stdout-class` · `-baseline-refs` ·
  `CHECK-M257x-iter58-derive-preregistrations` · `FIX-M257x-iter57-within-block-drift` ·
  `CHECK-M257x-iter57-anchor-guard-bare-class` · `FENCE-M257x-iter54-refs-block` ·
  `CHECK-M257x-iter52-second-ai-manager` · RF-2/3/7–13.

**Lessons:**

1. **Repair by predicate, but do not edit only the predicate.** 12 of the 22 sites carried a *second*
   false assumption in the same sentence. The predicate is what finds the sites; the sentence is what
   you fix.
2. **A ref-pin is an exemption, and exemptions accumulate where the corpus is oldest.** Both surviving
   holes are in the corpus's most-cited, most-dated prose — exactly where a reader is most likely to
   trust it.
