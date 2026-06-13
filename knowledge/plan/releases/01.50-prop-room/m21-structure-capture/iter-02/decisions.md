# M21 iter-02 — decisions (iter-local)

_(Cross-iter decisions live in the milestone-root `decisions.md` as M21-D1..D5. This file holds iter-local notes.)_

- **iter-02-L1 — Scope discipline (route the real artifact forward, don't game the ordinal).** Mid-iter, the
  structure round-trip mechanism validated cleanly, but producing the *real* 9-collection artifact was found to depend
  on a structure-source decision (real prod types are needed for the stage-4 row COPY). Rather than author a
  fake-typed 9-collection snapshot purely to bump furthest-passing-stage 2->3 (the "claim un-probed lift" anti-pattern),
  the iter closes honestly at stage 2 (live-confirmed + secured) and routes the real artifact + source decision to
  iter-03. The metric not advancing is the truthful outcome; the deliverables that DID land (email fix, baseline
  refinement, digest characterization, mechanism validation) are real.
- **iter-02-L2 — Throwaway harness.** Built on a shared docker network `m21test-net` with `m21test-pg`
  (pgvector pg16, port 55432 = offset N=5/`dev-5`) + ephemeral `directus/directus:11.6.1` runs. Torn down at close.
  iter-03 rebuilds it in ~1 min; worth codifying as a reusable test fixture/script in `stack-snapshot` once the
  structure artifact stabilizes.
