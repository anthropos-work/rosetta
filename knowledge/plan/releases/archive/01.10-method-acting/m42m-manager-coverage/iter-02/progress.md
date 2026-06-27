**Type:** tik (TOK-01 line 1 — the Studio-link escape)

# iter-02 — the Studio-link escape: diagnose → RE-SCOPE TRIGGER

## Phase A/B — diagnosis (read-only)
Traced the baked `studio.anthropos.work` left-nav link to its source + confirmed live on demo-3. See
`decisions.md` D1 for the full evidence chain. Summary: `STUDIO_URL` (next-web `core-js/constants/urls.ts:12`)
is a `NEXT_PUBLIC_NODE_ENV`-gated ternary with NO per-URL `NEXT_PUBLIC_STUDIO_URL` override; the demo build
leaves `NEXT_PUBLIC_NODE_ENV` unset → the prod host bakes into the bundle. The only knob (a global dev-flip)
sends Studio to the wrong port `:9000` (demo studio-desk is `:39000`) AND breaks `WEB_APP_URL`/`HIRING_APP_URL`
across other surfaces. No platform-source edit is allowed (v1.10 zero-edit line), and next-web offers no
env-override seam (unlike ant-academy's `ACADEMY_URL`/`NEXT_PUBLIC_ACADEMY_URL`).

## Phase C — NO FIX LANDED (re-scope trigger)
The hypothesis (env-rewritable host) is **falsified**. The sole clean fix is a platform-repo edit to
`urls.ts`, which is forbidden. No code changed; no platform edit made.

## Close — 2026-06-25

**Outcome:** The Studio left-nav escape (139 of the manager residual, all `studio.anthropos.work`) is **NOT
closable in rext** — next-web's `STUDIO_URL` is a `NEXT_PUBLIC_NODE_ENV` ternary with no per-URL env override,
so the only fix is a forbidden platform-source edit. RE-SCOPE TRIGGER. The diagnosis is the iter's deliverable
(a complete, falsified investigation).
**Type:** tik
**Status:** closed-no-lift (the iter's planned investigation completed with documented falsification — no fix
attempt landed; the metric did not move; the characterization IS the deliverable per Phase 4 Step 0)
**Gate:** NOT MET (escapes=139 unchanged — blocked on a platform-only fix)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: **y** — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-3 (re-scope-trigger)
**Decisions:** iter-02 D1 (the full evidence chain) + milestone-root RESCOPE-1 + D1.
**Side-deliverables:** none.
**Routes carried forward (pending the user's re-scope decision):**
- The rest of TOK-01 (lines 2-4) is INDEPENDENT of the Studio escape and remains valid next work once the
  user resolves the re-scope: reconcile the manager manifest to the real `/enterprise/workforce` tab route +
  populate the M36 dashboard (line 2); sample rules + cap raise for the two manager fan-outs (line 3, a
  precondition); manifest calibration (line 4). A future build-mstone-iters call resumes here under TOK-01.
**Lessons:** a demo's "rewrite the escaped host via injection" play (the protocol's escape routing row) only
works when the platform exposes a per-URL `NEXT_PUBLIC_*` override for that host. next-web's enterprise nav
hardcodes Studio behind a coarse `NEXT_PUBLIC_NODE_ENV` mode-flip (no per-URL knob), and the mode-flip's other
URLs make it unusable — so the escape is platform-bound. The diagnostic to apply for ANY future
baked-URL escape: find the constant's source, check whether it reads a dedicated `NEXT_PUBLIC_<thing>_URL`
override (rewritable in `.env.local`/build-arg, zero-edit) vs a mode-flip/hardcode (platform-bound →
re-scope). Recorded in coverage-protocol.md (the escape routing row's zero-edit precondition).
