# iter-05 — decisions (tik, TOK-01 acceptance)

**D1 — The whole TOK-01 result reproduces from a FRESH zero-manual `demo-up` — no manual step.** Tore demo-3
down `--purge` + removed the `demo-3-next-web` image (so the fresh build re-bakes WITH the demo-patch), then a
fresh `up-injected.sh 3` at the consumed tag. The bring-up tail ran **entirely automatically**: the demopatch
applied-then-reverted (log: "demopatch apply: next-web-studio-url applied … reverted after the image bakes"),
migrate (5 services + global Sentinel policy), the auto-set-dress (taxonomy 330228 rows + directus structure
auto-provision + replay 11982 rows + **local-served** + sim-embeddings + the Stories & Heroes seed including the
**FeedbackSeeder mirror**), the **Sentinel Casbin RELOAD** (via `AuthorizationService/Reload`, "no manual
restart"), the presenter cockpit on :37700, and the autoverify (`/api/health` 200, `casbin_rules=750`,
`directus_collections=14` local-served, all probes). **Zero manual intervention** — the user's core requirement.

**D2 — The Studio demo-patch reproduced exactly on the fresh build.** The freshly-baked next-web SERVED bundle
(`.next/static` + `.next/server`) carries **0× `studio.anthropos.work` + 31× `localhost:39000`** — identical to
iter-03's result, now on a from-scratch rebuild. The 3 `studio.anthropos.work` hits elsewhere in the image are
ALL in `.next/cache/webpack/*.pack` (webpack build-cache artifacts, **never served**) — they hold the
design-kept fallback ternary string (`process.env.NEXT_PUBLIC_STUDIO_URL || (… ? localhost:9000 :
studio.anthropos.work)`), which is behavior-identical when the env is set and is the intended fallback. The demo
clone's `urls.ts` is **git-clean** after the build (the trap-revert fired). The CANONICAL platform repos were
never touched.

**D3 — The FeedbackSeeder mirror reproduced on the fresh build (the org-feedback page populates).** On the fresh
stack, **162/162** `public.job_simulation_feedbacks` rows are joinable to the `public.local_jobsimulation_sessions`
mirror (the org-feedback resolver's JOIN target), with the manager org (`2222…`) carrying **103** joinable rows
— matching iter-04's distribution. The FeedbackSeeder mirror write ran automatically inside the auto-set-dress
stories seed; the manager sweep's `/enterprise/organization-feedback` page reached 200 and PASSed its section
asserts (no "No data"). The org-feedback empty-page class stays closed on a fresh build.

**D4 — Reproduction gap found + fixed: a stray tooling-owned `.dockerignore` left the demo clone git-dirty
(rext R1b fix, iter-05).** `build_frontend_next_web` (up-injected.sh) drops a TRANSIENT context-trim
`.dockerignore` into the next-web clone for the build and removes it on its build trap — but **only when THAT
run created it** (`di_ours=1`). The first fresh-build attempt crashed during compose-up on a Docker-VM
**disk-exhaustion** (environmental: `redis exited (1) — No space left on device`, the VM root was 100% full from
the fresh ~3.7 GB next-web rebuild on top of accumulated images + build cache); the demopatch trap reverted
`urls.ts` cleanly, but the next (resume) run hit the build function's early `return 0` (image already built) and
the `.dockerignore` overlay from the crashed run persisted as a stray `?? .dockerignore`. demopatch **R1**
pristine-ing reverts only demopatch-managed paths (`urls.ts`), not this overlay. **FIX:** extend
`ensure-clones.sh` with an **R1b** sweep that removes a stray next-web `.dockerignore` **only when it's
byte-identical to our tooling overlay AND untracked** (`git ls-files` empty) — so a real repo `.dockerignore`
(were one ever added upstream) is never clobbered. Untracked-only, idempotent, non-fatal; runs automatically at
the head of every fresh `demo-up`. NOT a tracked/canonical edit (the stray file was always untracked — zero
canonical edits held throughout). rext `method-acting-m42m-iter05`. Verified: the sweep removed the live
residual → the demo clone is now fully git-clean.

**D5 — The disk-exhaustion was environmental, not a tooling/reproduction bug.** The first bring-up attempt's
`redis exited (1) — No space left on device` was the Docker VM root disk hitting 100% (55.5/58.4 GiB) during the
fresh frontend rebuild — not a defect in the bring-up logic. Reclaimed via `docker builder prune -af` (39 GiB) +
`docker image prune -f` (10 GiB) → 35.8 GiB free, then **resumed the same fresh `demo-up`** (idempotent: builds
skipped on existing images, compose-up + migrate + set-dress completed, exit 0, verified-working). The bring-up's
own pre-flight already WARNs on an undersized VM (it printed the 9-GiB-VM warning, non-fatal). No code change
needed for D5 beyond D4's cleanliness fix; recorded so a future operator on a tight box recognizes the symptom.

## Gate results (the acceptance)

- **MANAGER (dan-manager), fresh demo-3, cap-120, gated:** `reachable=70/120 failingSections=0 personaFailures=0
  escapes=0 notReached=0 frontier=EXHAUSTED`, `gateMet=true`, `cappedAtFrontier=false`. PERSONA all ok
  (role-skills / avatar / org-identity). 1 documented exception (`/enterprise/settings`) + 1 presenter note
  (the `/reimport-profile` LinkedIn help link — allowed external reference). `1 passed (13.1m)`. **GATE: MET** —
  exactly reproduces iter-04's `(0,0,0,0) + EXHAUSTED`, now on a FRESH zero-manual build.
- **EMPLOYEE (maya-thriving), SAME fresh demo-3, cap-150, gated:** `reachable=59/150 failingSections=0
  personaFailures=0 escapes=0 notReached=0 frontier=EXHAUSTED`, `gateMet=true`, `cappedAtFrontier=false`. 1
  documented exception (`/settings`) + 2 presenter notes (wikipedia / dremio `/chapter` editorial citations —
  allowed external references). `1 passed (11.8m)`. **GATE: MET** — the M42e employee gate **holds** on the same
  fresh stack; the M42m changes (the FeedbackSeeder mirror, the manager manifest/sample-rules, the demo-patch)
  introduced **no employee regression**.
