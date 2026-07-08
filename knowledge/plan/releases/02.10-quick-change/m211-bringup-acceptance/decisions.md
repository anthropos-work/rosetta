# M211 — Decisions

_Implementation choices with rationale, logged as they are made._

## TOK-01: Warm-first cache-migrate, then cold-prove both stacks — 2026-07-08

**Tok type:** bootstrap (iter-01)
**Initial strategy:** Prove the merged 4-subgraph platform stands up end-to-end via the re-grounded
tooling in four moves: (1) **Consume the re-grounded tooling** — re-pin `.agentspace/rext.tag` →
`quick-change-m209` (or drive from the authoring copy). (2) **Close the dev-side M25-D9 gap** — the demo
migrate path (`migrate-demo.sh`) already bootstraps the `extensions` schema + extensions and waits for PG
readiness; mirror that as a rext DEV pre-migrate hook so a cold `make reset-db`/`make migrate` on the main
dev stack succeeds (the platform Makefile is un-editable). (3) **Execute the user's cache-migration** —
re-key the existing real 42,790-row taxonomy cache `skiller.*→public.*` (manifest schema/payload/filter/
public_via + rename 10 payload files + land under the new narrowed-digest key), gated on an EMPIRICAL
column-match (`\d public.skills` == the 15 cached `skills` columns); never fabricate. (4) **Iterate
warm-first, prove cold** — drive the fast inner loop on the already-warm merged stack (reset-db →
extensions-bootstrap → migrate → replay [rc 0, ~42,790] → seed [closure green] → verify [merged
assertion]) to green, then a full COLD `/dev-up` + `/demo-up`, then M42 coverage + v2.0 Playthroughs.
**Rationale:** The merge is a pure schema-prefix move, so the real captured taxonomy is faithful once
re-keyed — no prod access needed (none is provisioned; cache-migration is the user's confirmed path). The
demo path proves the M25-D9 fix pattern already works, so the dev hook is a mirror not an invention.
Warm-first shortens fix→re-measure and dodges the docker-reap hazard on long cold builds; the cold run is
still required for the gate but only once the inner loop is green. Each tik = one bring-up phase →
triage → route fix to surface (rext/corpus/stack) → re-measure, close-on-gate.
**Strategy class:** new-direction
**Distance-to-gate context:** Gate is a COMPOSITE of ~6 sub-conditions on BOTH stacks, cold. Baseline:
(a) 4-subgraph/no-skiller compose **MET** (warm ps, no skiller container); (b) replay loads public.*
~42,790 NOT MET (cache stale skiller-keyed); (c) seed closure NOT MET; (d) verify merged-assertion NOT
MET (code ready); (e) M42 coverage + v2.0 Playthroughs NOT MET; (f) 0 residual skiller refs — code+corpus
clean (M209/M210), runtime-unconfirmed. → **1/6 met, warm-only.**
**Next-tik direction:** iter-02 (first tik) targets sub-condition (b): taxonomy replay loads public.*
(~42,790) into the WARM merged stack — re-pin consumption, empirical column-match, land the dev
extensions-bootstrap/PG-wait hook if needed, execute the cache-migration, run reset-db → migrate →
stacksnap taxonomy replay, measure replay rc (target 0) + public.skills count (target ~42,790).
