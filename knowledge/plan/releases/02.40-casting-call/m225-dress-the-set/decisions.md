# M225 — Decisions

_Implementation choices with rationale. One entry per decision; cite the code/doc it binds._

---

## KB-1 — Phase 0b KB-fidelity: reconcile S1's stale `job_position`-replay premise (M222 BA-6 / M223 D4)

**Finding (Phase 0b, 2026-07-16 — YELLOW).** The M225 scaffold (`overview.md` Scope.In #1, `progress.md` S1,
`spec-notes.md`) said S1 folds "the `directus.job_position` replay + the 5-sim capture" into auto-set-dress. That
premise was **refuted before M225 even opened**: M222 BA-6 measured `directus.job_position` = **0 rows captured**
(the prod "443" was never in a snapshot) and the recruiter scoreboard does **not** read `job_position`
(`JobSimulation.jobPosition` is optional/unused); M223 D4 formally **DROPPED** the `job_position` replay. The 5
"positions" ARE 5 real captured `SIMULATION_TYPE_HIRING` sims, resolved by `readHiringSimPool`.

**The corpus was already correct** — `snapshot-spec.md:386`, `seeding-spec.md:352`, `stories-spec.md:633`, and
`hiring.md:75` all document the drop. Only the M225 **plan docs** carried the stale framing.

**Decision:** reconcile the plan docs inline (this audit) to the real S1 scope — **fold the HIRING-sim
(`SIMULATION_TYPE_HIRING`) capture + replay into the default auto-set-dress** so the hiring org's positions +
content come up real with no manual steps; **no `job_position` replay / no `snapshot-spec.md` surface**. Not a
scope change (the release already decided this in closed M222/M223) — plan hygiene.

**Binds:** `overview.md` §Why-section + §Scope.In #1 · `progress.md` S1 · `spec-notes.md` §Auto-set-dress ·
rext `stack-seeding` `readHiringSimPool`/`HiringConfigSeeder` (M223) · `presets/stories.seed.yaml` hiring story.

---

## D1 — S1: the hiring org ALREADY comes up real on a default `/demo-up`; S1's genuine deliverable is the bring-up-tail GUARD + docs (there is no separate "5-sim capture" to fold in)

**Investigation (Phase 1).** Traced the full default `/demo-up` bring-up chain and found the hiring set-dress is
**already end-to-end default-on** — most of the scaffolded S1 work was delivered incidentally by M223/M224:
- The hiring org (Meridian Talent, `is_hiring=true`) is in the DEFAULT preset `presets/stories.seed.yaml` (M223),
  and a bare `/demo-up N` seeds the stories preset by default (M38; `up-injected.sh` `STORIES_PRESET`).
- `resolveContentRefs` runs `readHiringSimPool` UNCONDITIONALLY (`contentref.go:177`); `HiringConfigSeeder` +
  `HiringFunnelSeeder` are registered by default (`stackseed/main.go:430,451`).
- **The `SIMULATION_TYPE_HIRING` sims ride along in the STANDARD directus content-surface capture.** The directus
  surface captures the `simulations` table INCLUDING its `type` column (`stack-snapshot/directus/directus.go:190`)
  under the public firewall (`private=false AND tenant_id IS NULL AND status='published'` — M222 measured 87 public
  hiring sims). So the auto-set-dress directus replay lands them; `readHiringSimPool` resolves 5. **There is NO
  separate "5-sim capture" step, and NO `directus.job_position` table surface** (0 rows; `job_position` exists only
  as a *column* of `simulations`, unread by the scoreboard — M222 D4).
- The two-app hiring UI container (`build_frontend_hiring`) + its 4 demo-patches are default-on (M224;
  `build_frontends` at `up-injected.sh:1530`, only `DEMO_NO_UI=1`/`DEMO_NO_PATCH=1` opt out).

**So no hiring-specific manual step exists to REMOVE.** The one residual risk is the general **cold-cache** case: a
cold/empty snapshot cache (or a starved HIRING pool) leaves `readHiringSimPool` empty → the seeders HONESTLY
degrade to 0 positions / 0 sessions → the recruiter comparison renders EMPTY while the stack still says UP. M223's
adversarial review named the downstream M224/M226 render gate as the loud catch.

**Decision:** S1's genuine, non-redundant deliverable is (a) a **bring-up-tail GUARD** that brings that catch
FORWARD to the default `/demo-up` tail, and (b) the documentation fold-in. Built the `autoverify.sh` demo-only
cheap-win (e) — the exact shape of the ISSUE-7 casbin assert: gated on a hiring org existing (`is_hiring=true`, so a
hiring-less demo SKIPS, no false-warn), it asserts ≥5 `organization_sim_invitation_links` positions + ≥40
`local_jobsimulation_sessions` candidate sessions for the hiring org; else WARNS (non-fatal) that the comparison is
empty (cold cache / starved pool). So "the hiring org comes up real with no manual steps" is now a CHECKED property
of the default bring-up, not an assumed one. 6 new tests (`test_verify.py::TestAutoVerify`), 120/120 green +
shellcheck-clean.

**Binds:** rext `stack-verify/live/autoverify.sh` (e) · `stack-verify/tests/test_verify.py` · doc fold-in
(`snapshot-spec.md` / `recipe-snapshot-world.md` / `verification.md` / `demo-up-defaults.md`, Phase 5).

**Not gold-plating:** no invented plumbing — the finding that S1 was largely pre-delivered is a real result; the
value added is making it self-verifying (the guard) + documented.

---
