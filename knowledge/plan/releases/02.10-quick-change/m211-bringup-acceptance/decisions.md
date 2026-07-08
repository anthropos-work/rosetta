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

## Close — deferral re-audit fate decisions (Phase 1b) — 2026-07-08

Fresh fate decisions from `audit-deferrals/deferral-audit-2026-07-08-m211-close.md` (verdict **GREEN**, 0 blocking):
- **DEF-M208-01 / M25-D9** (dev cold DB-init extensions bootstrap) — **RESOLVED / LAND-NOW satisfied** in M211
  iter-17 (`dev-stack/migrate-dev.sh`, cold-verified). Roadmap-vision M25-D9 marked resolved.
- **DEF-M208-02** (`INVITATION_HMAC_SECRET` dev `.env`) — **Fate-2 confirmed-covered** by the standing
  `/stack-secrets` tooling (the demo secret-DNA gene exists since v1.10b M49 #4; M211 cold `/demo-up` exercised
  it). No plan edit.
- **TEST-1** (rext `stack-seeding/README` test-count drift, ~788/13 actual vs 496/8 quoted) — **AGED_OUT →
  re-fated fresh: Fate-2 → `/developer-kit:close-release`'s rext roll.** Structural (rext frozen per-milestone at
  its tag; the reconciliation can only land at the rext code-of-record roll), not chronic. Firm imminent owner.
- **CAVEAT-1** (literal full clean-box destructive `/dev-up` not run) — **KEEP as belt-and-suspenders backlog**
  (accepted gate-interpretation; the dev-specific delta = M25-D9 DB-init was cold-verified on a faithful
  throwaway + a live docker harness; a full destructive `/dev-up` would clobber the user's committed native
  content-line dev). NOT an escape-hatch scope-break — the gate is MET. Added to `roadmap-vision.md` backlog.
- **CAVEAT-2** (pre-existing `dev-stack` CLI unit-test failures from an incomplete local `.agentspace/secrets`
  source) — **pre-existing environment/secrets-source drift**, unrelated to M211; standalone `migrate-dev.sh`
  tests green. Not roadmap scope.

## Close review findings (Phases 2–5) — 2026-07-08

- **Code quality (Phase 2):** the rosetta merge is **docs-only** (3 corpus/skill docs + plan artifacts; 0 code)
  → nothing to lint/type-check. The tooling CODE lives in the rext repo (frozen @ `quick-change-m211` =
  `2039103`), already hardened across 4 stabilized final passes (0 bugs, flake gate clean). **0 must-fix.**
- **Tests (Phase 4/8) spot-verified GREEN at the frozen tag:** Go `playthroughs/manifest` ok + `go vet` clean;
  demo-stack Python 114 (`test_frontend_build` + `test_demopatch`); dev-stack `migrate-dev` static+shellcheck 14;
  TS `coverage-manifest.unit.spec.ts` 32 (Playwright runner). Live docker `TestMigrateDevLive` (4) proven 3×
  clean in harden — not re-run at close (container churn). The **33 `test_dev_stack.py` failures are CAVEAT-2**
  (the dev-up CLI secret-coverage pre-flight failing on the incomplete local `.agentspace/secrets` source —
  app 1/4, platform 13/26 short) — pre-existing environment/secrets-source drift, unrelated to M211 code.
- **Docs (Phase 3):** the 3 touched corpus/skill docs (`setup_guide.md`, `dev-up/SKILL.md`, `playthroughs.md`)
  are accurate + all cross-references resolve. **DOC-1 (minor):** the new `migrate-dev.sh` is not indexed in the
  rext `dev-stack/README.md` (its sibling `migrate-demo.sh` IS in `demo-stack/README.md`). User-facing home is
  covered (corpus). Since rext is frozen at `2039103`, → **Fate-2 → `/developer-kit:close-release` rext roll**
  (bundled with TEST-1's rext README test-count reconciliation; the rext-README reconciliations land at the
  code-of-record roll). Not fixed now (a rext commit would move HEAD past the re-tag target).
- **Decision triage (Phase 5):** the M211 mechanisms are already flowed to the corpus — `migrate-dev.sh` cold
  DB-init (`setup_guide.md` § Full Database Reset + `dev-up/SKILL.md`), the casbin `init_policy.sql` load
  (`setup_guide.md`), the Playthroughs reset-to-seed roster-refresh (`playthroughs.md`). Remaining iter/tok
  decisions are implementation-mechanics → **archive** (maintainer-only).
