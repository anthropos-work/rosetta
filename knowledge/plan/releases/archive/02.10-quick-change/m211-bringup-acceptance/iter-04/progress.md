**Type:** tik (bring-up fix, multi-step) — under TOK-01. Protocol: run bring-up phase → triage → route fix
to surface (corpus) → re-measure.

# iter-04 — tik progress

## Execution log
1. **Triage.** `bootstrap-dev` (the platform's e2e dev target) step 3/5 pipes `init_policy.sql` but the
   dev DOCS (setup_guide.md/dev-up SKILL) drive `make up` + `make migrate` + a manual `CREATE SCHEMA
   sentinel` — and NEVER load the casbin policy. sentinel is `migrations: false` (not atlas-migrated);
   `init_policy.sql` does `CREATE TABLE IF NOT EXISTS sentinel.casbin_rules` (no `CREATE SCHEMA`).
2. **Empirical test.** Created `sentinel` schema + restarted sentinel → `casbin_rules` table auto-created
   by sentinel's Casbin adapter but **EMPTY (0 rows)**. → sentinel does NOT self-seed the policy; the
   `init_policy.sql` load is a genuinely missing dev-bring-up step (the demo path + bootstrap-dev do it).
3. **Routed fix (corpus).** Added the policy-load step to:
   - `corpus/ops/setup_guide.md` — after the sentinel-schema creation: `docker compose exec -T postgresql
     psql … < ../sentinel/init_policy.sql`, restart-after-load, + a verify note (`casbin_rules` count > 0)
     + a `reset-db` re-load note.
   - `.claude/skills/dev-up/SKILL.md` — the build-phase step list now names the Sentinel policy load
     between the schema step and `make migrate`, with the 403-consequence rationale.
4. **Operational apply (warm stack).** Loaded `init_policy.sql` → `INSERT 0 68` → `casbin_rules` = 68
   (p=6, p2=32, p3=17, p5=13 — the role→feature p-model). Restarted sentinel+backend; `/api/health` = OK.
5. **Re-seed** (--reset --force cleared tenant tables + g2 grants, preserved the 68 p-model rows) → re-seed
   stories-maya: **identity ok=4, users ok=250 — 0 seeder failures** (was "2 seeder(s) failed"). Audit 49
   writes, 45,785 rows, isolation clean, prod=false. casbin_rules now 170 (68 p-model + g2=51 + g3=51).
6. **Re-measure closure** → `[PASS] seed-verified-skill-closure` (still GREEN post-clean-seed).

## Re-measurement (gate sub-conditions, warm)
| Sub-condition | Pre-iter | Post-iter |
|---|---|---|
| (a) compose / no-skiller | MET | MET |
| (b) replay loads public.* | MET | MET |
| (c) seed closure green | closure-green / seed-partial | **MET (clean seed, 0 failures, closure PASS)** |
| (d) verify merged-assertion | NOT MET | **casbin cheap-win satisfiable** (casbin_rules=170 > 0; /api/health OK); full probe suite pending iter-05 |
| (e) M42 coverage + Playthroughs | NOT MET | NOT MET |
| (f) 0 residual skiller refs | clean | clean |
**Metric:** fully-met sub-conditions 2/6 → **3/6** (c fully closed); (d) unblocked (casbin>0).

## Close — 2026-07-08

**Outcome:** Found + fixed the dev bring-up's missing casbin-policy load (sentinel auto-creates casbin_rules
EMPTY; only demo/bootstrap-dev seeded it). Routed the fix to setup_guide.md + dev-up SKILL, applied it on
the warm stack (casbin_rules 0→170), re-seeded clean (0 failures), closure still GREEN.
**Type:** tik (bring-up fix)
**Status:** closed-fixed
**Gate:** NOT MET (3/6 fully; (d) full verify + (e) coverage/playthroughs + cold /dev-up + /demo-up remain)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n (fix was corpus+operational, no platform edit; no architectural decision needed) — (5) cap-reached: n (3 tiks) — (6) protocol-stop: n — Outcome: continue
**Decisions:** iter-04 D1 (dev-path casbin-policy gap is real + pre-existing, fixed in corpus), D2 (fix is docs+op, not platform/rext — no escalation)
**Side-deliverables:** none (both doc edits ARE the routed fix, not side-discovery)
**Routes carried forward:** iter-05 → run the full verify net (`autoverify`/`verify live`) on the warm stack for sub-condition (d) beyond the casbin cheap-win. Observation to verify later: the dev-up SKILL still says "12 containers" — post-merge the skiller container is gone (warm ps shows 11); confirm against the graphql profile + fix in a re-sync pass (not changed here — unverified against the profile definition).
**Lessons:** A container-only warm bring-up (M208) AND the documented dev cold path both leave `casbin_rules` empty — the dev docs never loaded `init_policy.sql` (only demo/bootstrap-dev did). The seed's casbin-grant step + the M18 verify cheap-win are both early detectors. This is a pre-existing bring-up gap the acceptance milestone legitimately closes (docs, no platform edit).
