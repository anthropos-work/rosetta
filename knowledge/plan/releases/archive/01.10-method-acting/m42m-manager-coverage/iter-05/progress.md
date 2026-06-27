# iter-05 — progress

**Type:** tik (acceptance — TOK-01 fresh-build reproduction)

## Work log

- **Step 1 — bump the consumed tag.** `stack-demo/rosetta-extensions` bumped `iter03 → method-acting-m42m-iter04`
  via the authoring remote (carries the FeedbackSeeder mirror + the reconciled manager manifest + the 4 manager
  sample rules), then `→ iter-05` after the R1b fix below.
- **Step 2 — fresh, zero-manual demo-up.** `demo-down --purge` demo-3 (0 containers, registry `{}`) + removed the
  `demo-3-next-web` image, then a fresh `up-injected.sh 3` (full default flow), foreground/streaming.
  - The first attempt's build phase succeeded (demopatch applied + reverted; next-web re-baked) but compose-up
    **failed on `redis exited (1) — No space left on device`**: the Docker VM root disk was 100% full
    (environmental — the fresh ~3.7 GB next-web rebuild on top of accumulated images + 26 GB build cache).
    Reclaimed via `docker builder prune -af` (39 GB) + `docker image prune -f` (10 GB) → 35.8 GB free, then
    **resumed the same fresh demo-up** (idempotent: builds skipped on existing images) → exit 0, verified-working.
- **Step 3 — verify the zero-manual reproduction.** All automatic, no manual step (see decisions D1–D3):
  - **demopatch** applied-then-reverted (clone `urls.ts` git-clean after); SERVED bundle **0× prod / 31×
    `localhost:39000`**; the 3 `studio.anthropos.work` hits are webpack-cache-only (the kept fallback ternary).
  - **FeedbackSeeder mirror** seeded automatically — **162/162** feedback rows joinable; manager org = 103.
  - migrate + auto-set-dress (taxonomy 330228 rows + directus structure auto-provision/replay 11982 rows +
    **local-served** + sim-embeddings + stories seed) + **Sentinel Casbin RELOAD** (no manual restart) + cockpit
    :37700 + autoverify GREEN (`/api/health` 200, `casbin_rules=750`, `directus_collections=14` local-served).
  - **GAP found + fixed:** a stray untracked `?? .dockerignore` (a crashed-build residual) left the demo clone
    git-dirty → fixed in rext via the **R1b** sweep in `ensure-clones.sh` (D4). Verified: swept → clone git-clean.
- **Step 4 — MANAGER gate (the acceptance).** Sweep `dan-manager` on the fresh demo-3, cap-120, gated.
- **Step 5 — EMPLOYEE re-sweep (regression).** Sweep `maya-thriving` on the SAME fresh stack, cap-150, gated.

## Close — 2026-06-26

**Outcome:** BOTH gates reproduced **GREEN** on the fresh, zero-manual demo-up. **MANAGER** (dan-manager):
`reachable=70/120 (0,0,0,0) EXHAUSTED gateMet=true` — exactly reproduces iter-04. **EMPLOYEE** (maya-thriving):
`reachable=59/150 (0,0,0,0) EXHAUSTED gateMet=true` — no regression. One minor clone-cleanliness gap found +
fixed in rext (R1b `.dockerignore` sweep); the disk-exhaustion was an environmental constraint, not a tooling bug.
**Type:** tik
**Status:** closed-fixed
**Gate:** MET
**Phase 5 grading:** (1) gate-met: y — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-1
**Decisions:** D1 (zero-manual reproduction), D2 (Studio bundle 0×prod/31×:39000 served), D3 (feedback mirror 162/162), D4 (R1b `.dockerignore` fix), D5 (disk env-issue).
**Side-deliverables:** the R1b `.dockerignore` sweep (rext `method-acting-m42m-iter05`) — a clone-cleanliness
hygiene fix surfaced by the crashed-then-resumed fresh build; it's part of making the fresh demo-up leave the
demo clone git-clean (the acceptance cleanliness contract), so it's in-scope, not unrelated. + a corpus
`coverage-protocol.md` fix-surface row for the platform-bound-escape demo-patch resolution.
**Routes carried forward:** none — the manager gate is reproducibly MET; the milestone is ready for harden+close.
**Lessons:** (1) a snapshot-CAPTURE-path-free acceptance still needs a FRESH `demo-up` (down --purge + remove the
frontend image) to prove the demopatch + seed-mirror reproduce — a re-seeded live stack proves the data but not
the *build*. (2) A crashed-then-resumed fresh build can leave a stray tooling `.dockerignore` the build trap
won't sweep (`di_ours=0` on the resume's early image-reuse return) — the ensure-clones R1b sweep closes it. (3)
The Docker VM disk is the real budget ceiling for a fresh frontend rebuild on a box already holding images +
build cache; prune before a fresh `demo-up` or the compose-up dies on `No space left on device`.
