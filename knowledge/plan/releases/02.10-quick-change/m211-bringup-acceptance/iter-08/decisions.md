# iter-08 — decisions

## D1 — Cold `/demo-up` GREEN cold on the merged platform (a-d,f COLD-proven)
A full cold `/demo-up` (from a torn-down state, images cached) came up GREEN on the merged 4-subgraph
platform with the re-grounded tooling: no skiller container (a), the migrated `public.*`-keyed taxonomy
cache HIT on a COLD replay loading 42,790 public skills (b), the stories seed's verified-skill closure PASS
via `datadna measure-closure --stack demo-1` (c), the bring-up-tail autoverify merged-assertion OK
(casbin=1304, directus 21 collections per-stack-local, all liveness+readiness) (d), and 0 skiller residue
(f). This proves the warm inner-loop results (iters 02-06) were not warm-state artifacts — the demo half of
the "both stacks GREEN cold" headline is done.

## D2 — Re-pin the consumption clone via a LOCAL never-pushed tag fetch (gate-critical prereq)
The `quick-change-m209` tag lives only in the authoring copy (`.agentspace/rosetta-extensions`); it is NOT
pushed to GitHub (close-release owns the final push + consumption re-pin persistence). The demo consumes
tooling from `stack-demo/rosetta-extensions` at the `.agentspace/rext.tag` pin, and ensure-clones.sh only
OBSERVES the pin (non-fatal warn) — it does not check it out. So the operator must check the consumption
clone out to the pin BEFORE `up-injected.sh`. Since the tag isn't fetchable from GitHub, I fetched it as a
local-path ref from the authoring copy (`git -C stack-demo/rosetta-extensions fetch <authoring-path>
refs/tags/quick-change-m209:refs/tags/quick-change-m209`) then `git checkout quick-change-m209`. Local
never-pushed-tag ops are explicitly allowed. This is a manual bring-up prerequisite ONLY because the tag
isn't on GitHub yet; close-release's real push makes a plain `git fetch --tags && checkout` sufficient.

## D3 — sim-embeddings cache-miss is non-gate, routed forward (Fate-3)
The taxonomy replay loaded public.* cleanly, but the sim-embeddings surface replay skipped (rc=5) — no
cached snapshot matches this stack's schema digest. sim-embeddings is the simulation vector-search surface,
distinct from the taxonomy, and is NOT a gate sub-condition; the demo is fully functional without it. Routed
forward as an optional operator-confirmed cache-fill (a separate capture), not a bring-up blocker.
