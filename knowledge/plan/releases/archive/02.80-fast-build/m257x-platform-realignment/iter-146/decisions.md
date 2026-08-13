# iter-146 — decisions

## `D-M257x-146-1` — grep for the EMITTERS of a retired fact, not only its consumers

iter-13 re-pointed every *consumer* of the deleted GraphQL router: the injection generator, the
readiness probe, the service table, the build-args, the tailnet fronting plan. Measured at iter-146,
that work was **97.6 % complete** (82 of 84 references correct). The two misses are both **emitters** —
places that *state* the endpoint rather than *use* it:

- `dev-stack/dev-stack:285` printed `https://<host>:$((5050+offset)) (graphql)` to the operator as the
  reachability hand-off, after iter-13 had deleted the `tailscale serve` row that fronted that port.
- `demo-stack/frontend/hiring.Dockerfile:36` defaulted the client-bundle endpoint to
  `http://localhost:5050/graphql`.

**The sharpest form of it:** `gen_tailscale_serve.py:38-41` deletes the router's front and records the
hazard in its own words — *"fronting a port with no listener produced a trusted-cert HTTPS endpoint
that always refused, which is worse than no entry at all (**it looks configured**)"* — while the file
one directory over advertises that exact URL. **The repair removed the mechanism and left the
announcement**, and the announcement is what the human acts on.

A consumer that points at a deleted service fails: it refuses, it 404s, a probe reports down. An
**emitter** that names one does not fail at all — it is a string. Nothing in a test suite, a health
check or a bring-up can notice it, which is why it needs a fence of its own.

## `D-M257x-146-2` — a repair's completeness tracks EXERCISE, not care

Both misses sit on paths nothing runs: a test section no iter had ever executed (iter-145's twelve),
and a `--public-host` branch that needs tailscale **and** a public host to reach. Everywhere the code
executes, iter-13's re-point held — which is the honest reading of a 97.6 % census and is *not* a
criticism of iter-13.

**The operational form:** when auditing a re-point, do not sample uniformly. **Enumerate the paths that
are not exercised** — rarely-taken branches, build-arg defaults, operator output strings, and any test
file in a section the suite does not cover — and read those first. That is where the residual is, by
construction, because everything else has been exercised into correctness.

## `D-M257x-146-3` — the fence carves out COMMENTS, and that is rule 67's shape again

`\b5050\b` over the tooling returns 84 hits, 82 of them correct — and the densest correct users are
the four files that name the port **in order to say it is gone** (`readiness.sh`,
`gen_tailscale_serve.py`, `gen_injected_override.py`, `up-injected.sh`). A bare-token fence would go
RED on precisely the code that documents the repair best.

This is `§5` rule 67 in a second domain: **the same token carries opposite obligations depending on
whether the line asserts it or retracts it.** In prose that was undecidable and the arm stayed SURVEY
at 70 % precision. In *code* it is decidable, because the language marks the distinction: a comment.

So the predicate is **executable content only**, and the comment filter is **whole-line**. A
trailing-comment parser would need shell, Dockerfile and Python quoting rules to avoid stripping a `#`
inside a string, and being wrong in the permissive direction is how a fence stops fencing.

**Scope is an allowlist, deliberately** — seven files that bake a URL, front a port, or print one.
iter-143's measured lesson is that widening reach over this material returns ports as findings; a
repo-wide sweep here would return 84 hits to grade 2. The list grows by one line when a new emitter
appears, and `test_every_emitter_exists` fails if an entry stops existing, so it cannot pass
vacuously.

**Control run, not assumed** (`§8`): the pre-fix content of both files, recovered from `HEAD` rather
than reconstructed, produces exactly **2 hits**; the post-fix content produces **0**; and a synthetic
pair proves the carve-out exempts a documenting comment and does *not* exempt an emitted URL.
