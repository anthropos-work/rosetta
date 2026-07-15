# iter-04 — findings.md (the canonical BASELINE capture)

**M221 Phase C, cycle 1 — the FIRST live cross-machine battery cycle.** A **DEFAULT** `up-injected.sh 1`
(**NO FLAGS**) on `billion.taildc510.ts.net`, cold reset-to-seed, driven from a tailnet peer (this Mac). Per the
M215 direct-drive pattern: the first cycle is the **baseline** — capture EVERYTHING, do not expect a clean pass.

- **Cold proof:** T0 (pre-up) = `2026-07-15T03:35:22Z`; PG_VERSION mtime **from inside the container** =
  `2026-07-15 03:41:32Z` (> T0 => initdb re-ran => genuinely cold). Cleared `data/` + demo-1 images before up.
- **Fresh verdict:** `autoverify.json` `ts=2026-07-15T03:42:32Z` (this run), `green:false`, `warnings:3`.
- **rext tag on host:** `cue-to-cue-m221-r2` (rolled from m220-r5; see F0). Bring-up `UP_RC=0`.
- **Driven from:** this Mac (tailnet peer `kirality-mac-pro-6`), never on-host. Playwright 1.61.1 / node v22.22.

---

## The 8-condition gate — baseline verdict

| # | Condition | Verdict | Evidence |
|---|-----------|---------|----------|
| 1 | p95 click->ACCESS < 5 s BOTH heroes, over the tailnet | **MET** | maya-thriving (employee->/profile) **p95 2.23 s** (p50 0.73), dan-manager (manager->/enterprise/workforce) **p95 2.08 s** (p50 1.30); ACCESS 5/5 each; hero identity present ("Maya Chen" / "Dan Rossi"); HTTPS tailnet origin (TLS hop **inside** the budget). |
| 2 | Full replayed catalog — taxonomy + directus + sim-embeddings, NO skipped surface | **NOT MET** | **All 3 skipped (cache-miss).** taxonomy + sim-embeddings = **store-root resolution bug** (F1, IRONCLAD). directus = **genuine digest divergence** (F2). `public.skills=0`. |
| 3 | All 3 story orgs incl. AI-readiness | **MET** | orgs=3: **Cervato Systems / Solvantis / Northwind Aviation**; `ai_readiness_cycles=2`; 541 users/memberships. |
| 4 | Dana sees a FILLED AI-readiness page | **NOT MET** | Cascade of taxonomy=0 (F3): `ai-readiness-funnel` seeder **rows=0**; `interview_aggregated_reports=0`; `ai_readiness_user_step_progresses/snapshots/skills/narratives=0`. |
| 5 | Ben's STARTED workflow visible on his dashboard | **NOT MET** | Same cascade — `ai_readiness_user_step_progresses=0` (Ben's Step-1 rows absent). |
| 6 | Aria's COMPLETED state renders | **NOT MET** | Same cascade — Aria's stage-3 step-progress rows absent. |
| 7 | Remote access came up BY DEFAULT, no flag | **MET** | Auto-discovery: *"public-host AUTO-DISCOVERED — billion.taildc510.ts.net (all 6 tailscale capability rungs passed ... cert MINTED)"*; **no flag passed**; `tailscale serve` fronting **6 ports HTTPS** (13000 next-web / 13077 academy / 15050 / 17700 cockpit / 18082 backend / 19000 studio); reachable from Mac. |
| 8 | ZERO platform-repo edits | **NOT MET (at-risk)** | 12/13 clones CLEAN. **ant-academy DIRTY** (F5b): `code/public/catalog.json` + `code/public/content/index.md` (predev-hook regeneration, `generatedAt` timestamp-only) + `next.config.js` (dev-origins patch, reverts on --stop). The 2 generated files **persist dirty after teardown** until manually restored. |

**Baseline distance to gate: 3 MET (1, 3, 7) / 4 NOT MET (2, 4, 5, 6) / 1 at-risk (8).**

> **The dominant insight: gates 2, 4, 5 AND 6 collapse to ONE root cause — the snapshot store-root
> resolution bug (F1).** It skips the taxonomy replay -> `public.skills=0` -> the taxonomy-dependent seeders
> (personas, membership-skills, target-roles, **ai-readiness-funnel**) all drop to 0 -> the verified-skill hero
> stories AND the AI-readiness member data are empty. **Fix F1 (+ the directus re-capture F2) plausibly moves
> the baseline from 3-MET toward 6-7 MET in one shot.**

---

## F0 — the pin-roll gotcha (cost the first bring-up, clean fast-fail)

Rolling the rext pin requires updating **BOTH** the clone checkout **AND** `.agentspace/rext.tag`. The clone was
at m221-r2 but `rext.tag` still pinned m220-r5 -> ensure-clones **FATAL: rext pin mismatch** (`UP_RC=1`), lock
acquired+released cleanly, no partial state. Fixed by writing `cue-to-cue-m221-r2` to
`~/panorama/.agentspace/rext.tag`. **A green fast-fail — the guard did its job.** (Doc note for the runbook:
"roll the pin" = clone checkout **+** `rext.tag`.)

## F1 — * THE DOMINANT BUG: snapshot replay resolves the WRONG cache root (IRONCLAD) *

The bring-up reported all three snapshot surfaces as cache-miss and skipped them:
```
stacksnap: cannot replay — cache miss: no snapshot for taxonomy/5afc0bccf1df7ef538b643321fc6362f.
stacksnap: cannot replay — cache miss: no snapshot for directus/b4cb55bcee08c76f2c37980da460a683.
stacksnap: cannot replay — cache miss: no snapshot for sim-embeddings/032c99ea47678187631c59c31b4ef059.
-> set-dressed (snapshot:taxonomy=skipped(cache-miss) directus=skipped(cache-miss) sim-embeddings=skipped(cache-miss))
```
**But the cache HAS the exact-hash entries** for taxonomy (`5afc0bcc...`) and sim-embeddings (`032c99ea...`), with
all `.copy` files + `manifest.json`, at `~/panorama/.agentspace/snapshots/`. Ground truth:
- `STACKSNAP_STORE=~/panorama/.agentspace/snapshots stacksnap status` -> lists all 3:
  **taxonomy schema=5afc0bccf1df, 10 tables, 330,261 rows** / sim-embeddings 032c99ea4767, 1,490 / directus ea2e187a1605, 11,986.
- `stacksnap status` from cwd `stack-demo/rosetta-extensions/` -> **"no snapshots cached under
  /home/devops/panorama/stack-demo/rosetta-extensions/.agentspace/snapshots"**.

**Root cause:** the store root resolves via `STACKSNAP_STORE` (unset) -> else `workspaceRootFrom(cwd)` walks up to
the **nearest** `.agentspace`. There are **multiple** on the box — the real cache at `~/panorama/.agentspace`
**and an empty `~/panorama/stack-demo/rosetta-extensions/.agentspace`** (the consumption clone's own). The replay
runs with a cwd under `stack-demo/rosetta-extensions/`, so the walk-up finds the **shadowing empty** one -> miss.

**Exact seam:** `dev-stack/dev-setdress.sh:323` — `"$BIN_DIR/stacksnap" replay --surface "$s" --stack "$STACK"
--dsn "$BASE_DSN"` — passes **no `--store` and sets no `STACKSNAP_STORE`**. (`replay` supports `-store`; there is
**no `--dry-run` for replay**.)

**IRONCLAD confirmation (post-baseline diagnostic; the pure baseline was already fully captured):**
```
$ STACKSNAP_STORE=~/panorama/.agentspace/snapshots stacksnap replay --surface taxonomy --stack demo-1 --dsn postgres://postgres@localhost:5432/postgres
replayed "taxonomy" into demo-1: 10 table(s) cleared, 330261 row(s) loaded, reindexed [skill_embeddings ... job_role_embeddings]
  public.skills:    0  ->  42,790
  public.job_roles: 0  ->  22,470
```
The cache is valid, the schema matches, the COPY loads perfectly. **The ONLY defect is store-root resolution.**

**Fix (iter-05 #1):** pin `STACKSNAP_STORE` (or `--store`) to the workspace-level cache root in
`dev-setdress.sh` (workspace root = `$EXT_ROOT/../..` = the panorama root). Alternatively make
`workspaceRootFrom` skip a `.agentspace` that has no `snapshots/` (prefer the one that actually holds the cache).
Needs a RED-proven fence (the store-root divergence) + a re-cycle to ship+prove -> **routed to iter-05, de-risked
to certainty.**

## F2 — directus is a SEPARATE genuine digest divergence (not the F1 root)

Even with F1 fixed, **directus** still misses: replay wants `b4cb55bc...`, the cache has `ea2e187a...` (captured
2026-06-29). The live directus schema digest has diverged from the captured one. **Needs a re-capture** of the
directus surface (per `snapshot-cold-start.md`, once per release) — a second, smaller iter-05 item. So gate 2's
directus surface is blocked on F2 independently of F1.

## F3 — the taxonomy=0 cascade (why gates 4/5/6 are empty)

Seed run results (`stories.seed.yaml`, 36 write attempts, 12,246 rows, prod=false, isolation clean):
```
org 3 / users 2700 / identity 4 / org-settings 1 / ai-readiness-config 7 / assignments 671 / jobsim-sessions 1078
/ skillpath-sessions 556 / member_languages 1190 / tags 645 / activity 2160 / certificates 374 / feedback 225
/ hero-activity 48 / profiles 2173 / projects 21 / succession 165
  -- taxonomy-dependent seeders that DROPPED to 0 (need real public node-ids) --
  taxonomy 0 / content 0 / membership-skills 0 / personas 0 / target-roles 0 / ai-readiness-funnel 0
  / generated-batch 0 / population-evidence 0
```
`personas rows=0` = no verified-skill hero fan-out. `ai-readiness-funnel rows=0` = no member-level AI-readiness
(step progress / snapshots / interview reports / narratives). **All of gates 4/5/6 hang off this ONE cascade.**
DB confirms: `interview_aggregated_reports=0` (jobsimulation schema), `ai_readiness_user_step_progresses=0`,
`ai_readiness_snapshots/skills/narratives=0`; `user_skills=0` (so the M219 no-junk-skills gate is vacuously N/A —
there are no claimed skills to enumerate). `ai_readiness_cycles=2` is the ONE org-level datum that survived
(seeded independent of taxonomy).

## F4 — academy empty-grid: symptom REAL, carry root-cause DISPROVEN (re-characterized)

`FIX-M221-academy-empty-catalog` (F-M220-2) hypothesized: *"the home reads the LOCAL catalog and
`[build-local-catalog]` emits 0."* **The symptom reproduces, the cause does not.**
- Hydrated Playwright render (networkidle + 3.5 s): `#courseGrid` = **0 cards**, empty-state *"No adventures
  here... yet / Try a different search or clear all filters"* **visible**, body innerText = **351 chars** (matches
  the carry's "348").
- **But the data is present:** `catalog.js:657` is `export const CHAPTERS = [...PUBLIC_CHAPTERS, ...LOCAL_CHAPTERS]`
  with `PUBLIC_CHAPTERS` a **hardcoded, non-empty** array; the build log shows `[build-catalog] 2705 entries
  across 419 public chapters` and `[build-index] 529 chapters`; **`/catalog.json` serves HTTP 200, 2.29 MB,
  2705 courses**; the default audience lens is **"For all" (Everything in the catalog)** — not a restrictive
  filter. `[build-local-catalog] 0` is a **red herring** (the gitignored `.agentspace/local-content/` authoring
  overlay is correctly empty; it is additive).
- **Re-characterized root cause:** a **client-side catalog-render defect** — the 2705-course catalog reaches the
  client but the grid never renders it. Needs ant-academy repo investigation (out of bounds -> a demo-patch or a
  config/tenant seam once the render path is understood). **NOT a clean on-the-spot fix; NOT the local-content
  population the carry proposed.** M219's 400-char content floor stays honestly RED — do not weaken.
- Side-notes: the academy fires third-party trackers (GA/GTM/DoubleClick/LinkedIn — all `net::ERR_ABORTED`; no
  `no-thirdparty` patch exists for the academy, unlike next-web) + an internal `/api/e/` fails. Cosmetic: the
  launcher passes a double `--port 3077 --port 13077` (last wins, works).

## F5 — * FIX-M221-reap-native-academy is FIELD-CONFIRMED NOT FIXED *

`down --purge` printed *"host-native listeners (cockpit, ant-academy) are down; their ports are free"* — **false.**
The academy `next-server` (pid 1706264) + its `next` launcher (pid 1696617, reparented to init, PGID 1696486)
were **still bound to `0.0.0.0:13077`** after teardown. Root cause: **no `ant-academy.pid` file** existed, so
teardown's `ant-academy.sh --stop` (kills by pidfile) had nothing to target, and the port-reap did not catch it
either — a **double serves != works / D17**: the teardown asserted "ports free" while the port was held. Had to
**manually kill** the process group to satisfy "leave nothing running." **This threatens cross-cycle integrity**
(cycle N+1 would measure cycle N's academy on :13077). Routed to iter-05 (the reap must not depend on the pidfile
alone — reap the listener by port+identity, and the "ports free" assertion must actually probe the port).

**F5b (gate 8 tail):** the same academy dirties `code/public/catalog.json` + `code/public/content/index.md` via
its own predev hooks — a **`generatedAt` timestamp-only** 1-line change each — and `ant-academy.sh` does not
revert them (it only reverts the `next.config.js` dev-origins patch, which DID revert). So the clone stays DIRTY
after teardown until manually restored. Fix options (iter-05): stash/restore the 2 generated files around the
academy launch, or a documented generated-file allowance for gate 8.

## F6 — F-7 backend-api-url twin: measured, NOT the feared blackhole

Carry `PROBE-M218-backend-api-url-twin` feared a 10.5 s `UND_ERR_CONNECT_TIMEOUT` from inside the container.
Measured on billion: the **runtime** env inside `demo-1-next-web-app-1` is `NEXT_PUBLIC_BACKEND_API_URL=
http://localhost:8082` -> **connection-refused in 0.00 s** (fast, not a hang). The client-baked offset value is
read only client-side (iter-03's `backend_api_url_server_reader_guard` holds — 0 server-side readers). **F-7
stays dormant as designed** on this box; no server-side reader appeared.

## F7 — PROBE-M218-c3 (router content-path 403s): not reproduced this cycle

`docker logs demo-1-graphql-1 | grep 403` -> no 403s (a status-200 request line only). But the content path was
**not heavily exercised** (I drove login/ACCESS, not the logged-in content surfaces), and directus is empty
anyway (F1/F2). **Inconclusive — route the c-3 re-check to ride alongside the F2 directus re-capture**, when the
content path actually has content to serve.

## F8 — C-6 the RAM question: 7.3 GiB is SUFFICIENT (decided — no code fix)

`free -h` with all 15 containers up **+ taxonomy loaded**: total 7.3 Gi, **used 2.2 Gi, available 5.1 Gi**, swap
**2/15 Gi** used. `docker stats`: the whole 15-container stack = **~880 MiB** (next-web 199 / postgres 377 /
backend 81 / studio 59 / rest <20 each). The cold build completed fine; logins p95 ~2.2 s. **Decision: the 12 GiB
"floor" the tooling warns about is a conservative BUILD guard (concurrent Go compiles), not a runtime
requirement; a single demo-1 runs comfortably on 7.3 GiB.** Bumping the VM to 12 GiB would only add headroom for a
second concurrent stack. **Infra note, not a code fix.**

## F9 — demopatch self-healing freshness gate fired (as designed)

`app-aireadiness-snapshot-loadmembers`: whole-file sha **DRIFTED** (`b3216968...` -> `dc9e167e...`) but the
**anchor was intact (1x)** -> **SELF-HEALED**: recomputed + applied. Exactly the demopatch-spec behaviour ("the
anchor is the contract; the whole-file sha is only a baseline"). **Optional hygiene:** re-pin the recorded
baseline `pre_sha256: dc9e167eda1ad42e8c5eb1fe097b8602a037a4aa273c0f4c9ae767c5eb7a333b` (and
`post_sha256: 759590752af96cc515bdbf7b9a60a56a1975a5a1abf809ae1e8eff17abf97b94`) so it stops drifting. All other
patches applied cleanly (*"demo-patches: all applied (none refused, none skipped)"*).

## F10 — never-field-run M217 items: partial

- **Pre-bind reap:** ran naturally (the pre-bring-up `down --purge` + `ant-academy.sh`'s pre-bind reap during
  launch) — no stale port collision at bring-up. But the **teardown** reap **failed** (F5), which is the more
  important field result.
- **Freshness preflight:** **ran and self-healed** (F9) — the DRIFT+heal path is field-proven. The **abort path**
  (break an anchor -> assert it refuses) was **not deliberately exercised** — routed forward (an unexercised path
  is a finding, not a pass).
- **`assert_ports_free` (compose-range preflight):** not deliberately exercised (no port collision arose) —
  routed forward.

---

## Positives worth banking (baseline is not all red)

- **Remote-by-default WORKS end-to-end** (gate 7): auto-discovery -> cert mint -> serve fronting -> reachable from
  a real tailnet peer, **no flag**. The whole v2.3 D-DESIGN-3 flip is live on billion.
- **Latency is comfortably inside budget over the tailnet** (gate 1): p95 2.08-2.23 s vs the 5 s gate, TLS hop
  included, hero identity present.
- **The cold reset-to-seed machinery is sound**: cold initdb proven from inside the container, fresh verdict,
  36-write isolation-clean seed (12,246 rows), F12 serve-reset works, autoverify unlinks on teardown, host lock
  acquires/releases cleanly, all clones return to CLEAN (after the F5 manual academy kill + the F5b restore).
- **The host-isolation guard (iter-02) held**: one cycle, lock taken+released each phase, no concurrent run.

## Left clean
pgrep `[d]ocker build|[u]p-injected|[n]ext dev|next-server|cockpit` = EMPTY / containers 0 / host lock released /
`tailscale serve` empty / registry `{}` / **all clones 0-dirty** / rext pin `cue-to-cue-m221-r2`.
