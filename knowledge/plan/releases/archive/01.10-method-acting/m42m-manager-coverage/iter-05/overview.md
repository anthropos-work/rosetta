---
iteration_type: tik
status: closed-fixed
gate: MET
---

# iter-05 — FRESH zero-manual demo-up acceptance (the authoritative manager-gate reproduction) + EMPLOYEE re-sweep

**Active strategy reference:** `TOK-01` (manager-coverage — reconcile-route + clear-escape + populate-dashboard
+ exhaust-frontier). This is the **acceptance tik** that proves TOK-01's whole result reproduces from a fresh,
zero-manual build — the milestone's gate is explicitly "reproduced on a FRESH demo-up".

**Cluster / target identified:** iter-04 met the manager gate on a **re-seeded live** demo-3 (additive re-seed
of an already-up stack). The gate language requires the result on a **FRESH `demo-up`** (the user's core
requirement). The residual after iter-04 is therefore the **reproduction proof itself**: tear demo-3 down
`--purge` + remove the next-web image, run a fresh zero-manual `demo-up`, and confirm the iter-03 Studio
demo-patch + the iter-04 FeedbackSeeder mirror + the route reconcile + the Sentinel reload ALL apply
**automatically** (no manual step), then re-run the manager gate AND the employee gate on that fresh stack.

**Hypothesis:** a fresh `demo-up` at the consumed tag reproduces `(0,0,0,0) + EXHAUSTED` for the manager AND
holds the employee gate — zero manual intervention. No new fix expected (the fixes landed iters 03/04); if a
reproduction gap surfaces, fix it in rext and re-run.

**Expected lift:** none on the metric (already `(0,0,0,0)` on the re-seeded stack) — this iter's deliverable is
the **fresh-build reproduction** of that result + the regression check on the employee gate.

**Phase plan:** Phase A bring-up (fresh demo-up) → Phase A/D sweep (manager) → Phase D re-sweep (employee) →
Phase E close. Plus: bump the consumption clone to `method-acting-m42m-iter04` (then `iter-05` after the R1b fix).

**Escalation conditions:** a reproduction step that requires a MANUAL intervention to pass (a hand-restart, a
hand-seed, a hand-apply of the demopatch) is a GAP — fix it in rext and re-run, NOT a manual workaround. A
platform-repo edit need would be a re-scope trigger.

**Acceptable close-no-lift outcomes:** n/a — the gate is already met; this iter either reproduces it on a fresh
build (closed-fixed) or surfaces a reproduction gap (fix-and-re-run within the iter).

## Outcome (filled at close)
BOTH gates reproduced GREEN on the fresh zero-manual demo-up; one minor clone-cleanliness gap found + fixed in
rext (R1b `.dockerignore` sweep). See `progress.md` + `decisions.md`.
