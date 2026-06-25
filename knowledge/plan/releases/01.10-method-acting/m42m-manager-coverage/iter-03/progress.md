**Type:** tik (TOK-01 line 1 — the Studio-link escape, resolved via the demo-patch mechanism)

# iter-03 — demo-patch tool + the Studio urls.ts patch (resolves RESCOPE-1)

## Phase A/B — the mechanism (the user's RESCOPE-1 pivot)
RESCOPE-1 (iter-02) proved the Studio escape is platform-bound (no `NEXT_PUBLIC_STUDIO_URL` override). The user
chose a NEW approach: a **demo-patch tool** that patches the demo's EPHEMERAL clone before the build, reverts
after — CANONICAL platform repos NEVER touched. Built per `.agentspace/scratch/work-m42m/demo-patch-design.md`.

## Phase C — fix landed (rext only; zero CANONICAL platform edits)
- **`demo-stack/patches/demopatch`** — the CLI (apply|revert|status|check) with all **6 guards**: G1 hard
  path-assert (realpath + workspace-containment + exact-path + symlink-escape) + G2 pre-patch drift-refuse
  (sha256 pre/post + single-occurrence anchor) + G3 never-commit/push/working-tree-only (UNSTAGED assert + a
  source-grep unit test) + G4 idempotent (positive post_sha256+marker probe → no-op) + G5 content-anchored
  self-revert (+ `--force-pristine`) + G6 demo-only scope. Default-on + NON-FATAL; `DEMO_NO_PATCH=1` opts out.
- **`manifest_loader.py`** — a stdlib-only strict YAML-subset loader (no PyYAML; supply-chain stays GREEN).
- **`patches/next-web-studio-url/next-web-studio-url.yaml`** — mirror `ACADEMY_URL` (prepend
  `process.env.NEXT_PUBLIC_STUDIO_URL ||`, keep the ternary as fallback — behavior-identical when unset);
  `build_env NEXT_PUBLIC_STUDIO_URL=http://localhost:$((9000+OFFSET))`; pre/post sha256 pinned; anchor
  verified to match the demo clone's urls.ts exactly once.
- **`up-injected.sh` `build_frontend_next_web`** — append the offset `NEXT_PUBLIC_STUDIO_URL` to the existing
  `apps/web/.env.local` overlay; `demopatch apply` NON-FATAL before the docker build; extend the existing
  RETURN trap with `demopatch revert`. **`ensure-clones.sh`** — R1 pristine-ing pass + R2 push-block.
- **Fix landed during VERIFY:** the first fresh demo-up surfaced a real bug — G6 was registry-only and the
  consumption-clone registry is EMPTY at patch-time (the direct `up-injected.sh N` path populates it later),
  so the patch was G6-REFUSED. Fixed: G6 now accepts the demo on EITHER the **structural demo-workspace
  identity** (the binary's own clone-set + a `demo-stack/stacks/` dir — the fresh-build signal) OR a registry
  type:demo row. 18 tests green; both wired scripts shellcheck-clean.

## Phase D — re-measure (a FRESH demo-up, the authoritative gate)
A `demo-down --purge` + image-rm + FRESH `demo-up demo-3` (foreground/streaming), then:
- **(a)** demopatch APPLIED during the build (R1+R2 ran; `demopatch apply` succeeded) → the RETURN trap
  reverted it → `git -C stack-demo/next-web-app status` clean for urls.ts (the patched file). ✓
- **(b)** the **baked bundle** carries **0× `studio.anthropos.work`** + **31× `localhost:39000`** in
  `.next/static`+`server`. ✓
- **(c)** **LIVE click-through** as `dan-manager` via the cockpit Clerk-session handshake: the left-nav
  "Anthropos Studio" link + the ManagerHome Studio cards href = `http://localhost:39000/` (the demo's own
  studio-desk) — **prod-ejects=0, demo-local(:39000)=2**. The cross-origin acceptance gate PASSES. ✓
- **(d)** re-up → next-web image reused (no rebuild, patch path skipped); `demopatch apply` twice on the live
  clone = clean **G4 idempotent no-op**, no double-patch/drift. ✓
- **(e)** `git status` clean in BOTH `stack-demo/next-web-app` (urls.ts; the `?? .dockerignore` is a
  pre-existing M19 build artifact in the gitignored clone, not this iter's) AND rosetta (only the iter-03
  plan dir). rext authoring clean + tagged. ✓
- **Manager re-sweep** (bounded cap-30, dan-manager): `failingSections=0 personaFailures=0 **escapes=0**
  notReached=5 frontier=CAPPED`. The Studio escape class **139 → 0**. The one LinkedIn-help link is the
  allowed external citation (a presenter note), not an escape.

## Close — 2026-06-25

**Outcome:** RESCOPE-1 RESOLVED demo-only. Built the demopatch tool (6 guards + 18 tests) + the Studio
urls.ts patch + the up-injected/ensure-clones wiring; a FRESH demo-up applies-then-reverts the patch, the
baked bundle + the LIVE dan-manager click-through both resolve Studio to the demo-local studio-desk
(`:39000`), and the manager Studio-escape class dropped **139 → 0**. CANONICAL platform repos untouched.
**Type:** tik
**Status:** closed-fixed (the planned scope — the demo-patch mechanism + the Studio fix — landed; the
targeted Studio-escape cluster cleared 139→0 on a fresh demo-up, validated by the live click-through)
**Gate:** NOT MET (the manager gate also needs the M36 dashboard populate + the fan-out frontier exhaustion
— TOK-01 lines 2-4 — a LATER run; this iter closed only the Studio-escape line of TOK-01)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — Outcome: continue (TOK-01 lines 2-4 remain) BUT this is a single
large tik resolving the user's RESCOPE-1 pivot; the orchestrator caps run 2 at this phase boundary.
**Decisions:** M42m-D2 (the demo-patch mechanism + the G6 dual-signal fix) — see iter-03/decisions.md.
**Side-deliverables:** none.
**Routes carried forward (TOK-01 lines 2-4, a LATER run):**
- Reconcile the manager manifest to the real `/enterprise/workforce` tab route + populate the M36 dashboard
  (verification funnel / teams / role-readiness / succession / mobility) — line 2.
- Sample rules + cap raise for the two manager fan-outs so the frontier EXHAUSTS (`cappedAtFrontier=false`)
  — line 3 (a precondition for a real gate reading).
- Calibrate the manager manifest floors/selectors (`calibrated:false → true`) — line 4.
**Lessons:** (1) a demo-clone source patch is a clean, safe way to close a platform-bound escape WITHOUT a
canonical edit — the 6 guards make "patch only the ephemeral demo clone" mechanically enforced. (2) A
registry-only demo gate is WRONG for the fresh-build patch-time window: the consumption-clone registry is
empty until after compose-up, so the structural demo-workspace identity (the tool's own clone-set + a
`demo-stack/stacks/` dir) is the reliable fresh-build signal. Recorded in the patches/README + the tool's
inline guard docs.
