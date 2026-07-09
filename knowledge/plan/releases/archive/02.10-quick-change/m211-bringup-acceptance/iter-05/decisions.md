# iter-05 — Decisions (tik under TOK-01)

## D1 — The merged-platform verify assertion is confirmed live (not just code)
`stack-verify/lib/readiness.sh` expected-schemas = `(public sentinel cms jobsimulation skillpath
extensions)` — M209 dropped `skiller`. The probe PASSES on the warm merged stack with the `skiller` schema
absent; a stale list still expecting `skiller` would FAIL here. Combined with `graphql-introspection ok`
(the 4-subgraph supergraph composes + introspects with no skiller subgraph) and `sentinel-rpc ok`, this is
the live proof that the re-grounded verify tooling asserts the MERGED shape — sub-condition (d) and part of
(f) (no skiller schema/subgraph in the verify-queried path).

## D2 — The 4 unscoped verify failures are scope artifacts, not merge defects (no escalation)
Bare `verify.sh` probes the full service list including the M19 UI tier (next-web-app, studio-desk) + the
per-stack Directus. On a backend-only, prod-read warm stack those are legitimately DOWN (native frontends
not started; no `--local-content` local Directus). verification.md's scope model skips out-of-scope
services rather than false-failing them — proven by the SCOPED run going all-green. So the failures are a
"ran bare verify.sh against a reduced stack" artifact, not a bring-up defect. They are deferred to the full
UI/cold bring-up (sub-condition (e) + the cold proof), where those services ARE up.

## D3 — Session scope reality: the warm inner loop has proven (a)–(d); (e) + cold proofs are next-session
The warm-first strategy (TOK-01) has now proven the merged platform's DATA + backend + seed + verify chain
end-to-end via the re-grounded tooling: compose/no-skiller (a), replay loads public.* 42,790 (b), seed
closure green (c), verify merged-assertion (d). The remaining gate items — (e) M42 coverage sweep + v2.0
Playthroughs, and the full COLD `/dev-up` + `/demo-up` — require the UI tier + a cold demo bring-up
(Playwright + a ~3-min UI Docker build per demo-N), which exceed a single reap-safe foreground tik and
carry the M208 docker-reap hazard. They are routed to a dedicated next session, run as cold bring-ups.
