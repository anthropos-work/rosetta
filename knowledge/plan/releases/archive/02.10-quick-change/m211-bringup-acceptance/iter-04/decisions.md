# iter-04 — Decisions (tik / bring-up fix under TOK-01)

## D1 — The dev bring-up's missing casbin-policy load is a real, pre-existing gap; fixed in the corpus
Empirically proven: after `CREATE SCHEMA sentinel` + a sentinel restart, sentinel's Casbin adapter
auto-creates `sentinel.casbin_rules` but leaves it EMPTY — it does NOT seed the p-model policy. The demo
path (`migrate-demo.sh`) and the platform's `bootstrap-dev` (step 3/5) both load `sentinel/init_policy.sql`;
the dev DOCS (`setup_guide.md` / `dev-up` SKILL) create the schema but never load the policy. So a cold
`/dev-up` per the current docs yields an empty `casbin_rules` → every authorized route 403s (M18 silent-403
class) and the seeder's casbin-grant step has no p-model to attach to (iter-03's identity/users failures).
This gap is pre-existing (independent of the skiller→app merge) but it BLOCKS the M211 bring-up-acceptance
gate (`casbin_rules > 0`), so closing it is in-scope. Fix routed to the corpus (setup_guide.md + dev-up
SKILL) — a docs fix, no platform edit.

## D2 — Fix is docs + a chartered operational load; NOT a platform or rext edit (no escalation)
The policy load is `sentinel/init_policy.sql` (a platform-repo asset, READ not edited) applied via
`docker compose exec psql` (a chartered operational command). The durable fix is the corpus doc/skill step
that makes the cold `/dev-up` path load it. No platform-repo source is modified; no rext tooling change is
needed for the MAIN dev stack (N=0) because it is driven by the platform Makefile + the dev-up SKILL, not
the rext dev-stack tooling (that's for dev-N, N≥1, whose migrate-demo-style path already loads the policy).
So the `unimplementable-without-platform-edit` escalation does NOT fire. (If a future need arises to load
the policy for an offset dev-N via rext, migrate-demo.sh is the model — but N=0's path is the corpus.)

## D3 — --reset is casbin-aware (preserves the p-model)
`stackseed --reset --force` truncated tenant tables + "deleted 0 casbin g2 grant(s)" but LEFT the 68
p-model rows intact (casbin_rules stayed 68 after reset). So a re-seed does not require re-loading
init_policy.sql — the p-model persists; only the per-membership g2/g3 grants are cleared + re-seeded. This
matches the idempotency contract (the reset clears seeded grants, not the global policy).
