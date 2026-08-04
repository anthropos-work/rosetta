**Type:** tik — under `TOK-05`. Corpus-only; no rext change, no tag.

# iter-66 — root `CLAUDE.md`, against what this session measured

## What was wrong, and why no fence could see it

Root `CLAUDE.md` is the file every agent reads before doing anything, and iter-62 had already repaired
its profile table. Two statements survived that, and **neither is reachable by any fence in the
family** — they are prose about *which tier a service sits in* and *which RPC edges are live*, not a
profile token, an address value, a repo count or a migration flag.

1. **`Storage` was listed under *"In the default local profile (`core`)"*.** It is not. At platform
   `0dab54d` `storage` declares `profiles: [storage-legacy]` and the default `core` selection is five
   containers: `backend`, `gotenberg`, and the always-on floor (`postgresql`, `redis`, `sentinel`).
   A reader following `CLAUDE.md` would expect a storage container on a stock `make up` and not get
   one — and, worse, would not know why the storage calls fail.

2. **The RPC-edge inventory said *"backend → sentinel/storage and messenger → backend"*.** Two
   corrections: `messenger → backend` is now **all four** of its addresses (`d11a403` re-pointed
   `CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR` at `http://backend:8083`, `docker-compose.yml:174`/
   `:176` — **M809 has landed**, iter-63's finding); and **`backend → storage` is `mid-fold`, not
   live** — `app` still calls it in code, nothing sets `STORAGE_RPC_ADDR`, and `storage` is not
   started, so on a stock stack the client is built against the empty string and **fails at call time
   rather than at boot** (iter-64's finding).

Both replaced with the two-sided form the map uses, each side cited, and the mid-fold note points at
the fenced row rather than restating it.

## Gates

| gate | result |
|---|---|
| `platform_predicate_guard` · `platform_alignment_guard` · `anchor_construct_guard` · `markdown_structure_guard` · `corpus_index_guard` | **all OK** |
| §5 rule 34 re-point | no `corpus/**` file changed, so no intra-corpus citation moved |
| rext | untouched — no commit, no tag, pin unchanged at `fast-build-m257x-iter-65` |

## Close — 2026-08-04

**Outcome:** the two statements this session measured false in the corpus's highest-traffic file,
corrected and cited — a service placed in a tier it left, and an RPC-edge inventory that was wrong in
both directions at once (one edge fuller than stated, one emptier). **Neither is reachable by any
fence**, which is the point worth carrying: the predicate work of iters 60–65 covers profile tokens,
addresses, counts, flags and citations, and *"which tier is this service in"* is none of those.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5, unchanged; clause 5 is still graded only by a reading that returns zero.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** none beyond the two corrections; both facts were measured and recorded at iters 63–64
(`D-M257x-63-5`, `D-M257x-64-3`) and this iter only carries them into `CLAUDE.md`.
**Side-deliverables:** none.
**Routes carried forward:** unchanged from iter-65, plus —
- `FENCE-M257x-iter66-tier-membership` → **new.** *"Service X is in tier/selection Y"* is a predicate
  with a derivable legal set (`compose.select(default_profile)`), and nothing checks it. It is the
  same shape as G1 but about **membership** rather than about the token, and it is what would have
  caught this iter's first defect. Sized: one assertion over the same parsed compose the guard
  already holds.

**Lessons:**

1. **The highest-traffic file is not the best-fenced one.** Five iterations of predicate work left
   two false statements in `CLAUDE.md`, both about facts those very iterations measured. Fences
   cover *constructs*; a sentence that places a service in a tier is not one.
2. **An inventory can be wrong in both directions at once.** The RPC-edge list understated one edge
   (messenger's, now four) and overstated another (`backend → storage`, now mid-fold). Checking a
   list for staleness usually means checking for *removals*; this one also needed an addition.
