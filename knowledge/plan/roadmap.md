# Roadmap

Active development plan for **Project Rosetta** (the Anthropos documentation corpus + environment-
builder skills).

> **Designed 2026-06-02** from the Demo Environment + Clerkenstein brief, **refined 2026-06-02** to
> promote alignment measurement into a first-class discipline (new **M0**). 3 research agents over the
> Clerk integration, the staging/dev-env tooling, and the data/seeding surface — all verified against
> the cloned platform in `stack-dev/`. Gap analysis:
> [`.agentspace/scratch/roadmap-research-2026-06-02.md`](../../.agentspace/scratch/roadmap-research-2026-06-02.md).
>
> **v1.0 "body double" — SHIPPED 2026-06-03** (merged to `main`, tagged `v1.0`; full detail in `## Done` below).
> **v1.1 "show floor" — SHIPPED 2026-06-05** (merged to `main`, tagged `v1.1`; full detail in `## Done — v1.1`
> below). 8 milestones M3→M8: the 2-repo consolidation + demo/dev stacks + the production-safe seeding stack
> (framework + data-DNA + fleet) + the corpus product layer.
>
> **v1.2 "set dressing" — SHIPPED 2026-06-07** (tag `v1.2`, merged to `main`; designed 2026-06-05, **refined
> 2026-06-06** against live prod). 4 milestones M9a→M9b→M10→M11: a **dedicated `stack-snapshot` extension** that
> lifts M7c's two `waived` surfaces (`taxonomy` + `content`) to **100% data-DNA coverage** — capture the real
> *public* skill taxonomy + content library once from a **safe, low-impact source** (default a prod `pg_dump`), replay per-stack,
> *measured-faithful* via a new snapshot-fidelity dimension, with a tested **tenant-data firewall** (never customer
> data) + a **`.agentspace` manifest cache** (snapshots never land in any git repo). Detail in `## Done — v1.2` below.
>
> **v1.3 "stack party" — SHIPPED 2026-06-07** (tag `v1.3`, merged `--no-ff` → `main`; designed + built + closed
> 2026-06-07). 4 milestones M12→M13→M14→M15: the **dev/demo convergence** — dev stacks became first-class peers
> (the per-stack-Directus recipe + firewall check [print-only — see the correction below] + auto-snapshot + a light
> default seed on build), a **unified stack registry** that
> allocates the first-available N across dev+demo (no port collisions), one **generic `stack-*` skill set**
> (`stack-list`/`seed`/`snapshot`/`update` + `dev-up`/`dev-down`), and a code-cited **safety & security doc** (how
> the tooling never reads private data + never touches prod). Full detail in `## Done — v1.3` below.
>
> **v1.3b "dress rehearsal" — SHIPPED 2026-06-09** (tag `v1.3.1`, merged `--no-ff` → `main`; full detail in `## Done — v1.3b` below). A
> **field-hardening release** after shipped v1.3: v1.3 converged the dev/demo *model*, but the
> first real run of `/demo-up` surfaced 14 issues — a demo stack comes up **backend-only, unseeded, unverified**, and
> announces "UP" even when authz silently failed. v1.3b makes `/demo-up` produce a **full, populated, verified,
> demoable** stack and makes the repo honest about it: 5 milestones M16→M20 — land the applied fixes + restore doc
> truth, add re-run safety (idempotency + the first-run race), a post-bring-up verification net (`stack-verify` made
> offset-/scope-aware + auto-wired non-fatal), the **frontend tier** (next-web + studio-desk + ant-academy, tooling-only),
> and **lifecycle convergence** (demo-up auto set-dresses like dev-up + a cold-start capture path). All **tooling +
> docs only — zero platform-repo edits** (verified constraint). Brief: [`.agentspace/demo-up-issue.md`](../../.agentspace/demo-up-issue.md) +
> [`.agentspace/demo-up-frontend-plan.md`](../../.agentspace/demo-up-frontend-plan.md).

> **⚠ Correction (2026-06-11) — two roadmap-wide fixes:**
> 1. **v1.4 removed.** There is **no version currently staged**. The former v1.4 seeds (cloud `SnapshotStore` /
>    S3 media blob bytes [DEF-M10-01], AI content, shareability, more mirrors) are **unscheduled backlog**, not a
>    planned release. The DEF-M10-01 deferral persists as a real backlog item; its routing in the historical
>    milestone retros below has been updated from the removed v1.4 to **backlog (unscheduled)** (the audit facts —
>    inherited/signed/unchanged at each close — are unchanged).
> 2. **v1.3 "local Directus" was inaccurate.** v1.3/M13 shipped the per-stack-Directus **recipe + prod-Directus
>    firewall check (print-only)** and the **taxonomy** replay — **not** a working local Directus serving content.
>    No stack type stands a per-stack Directus up; the `directus` content replay skips (`stacksnap` exit 4 — the
>    M10 collection-schema gap), and every stack reads public content **live from prod**. Detail:
>    [`../../corpus/ops/snapshot-spec.md`](../../corpus/ops/snapshot-spec.md) § the per-stack Directus store fork.
>
> **→ Update (2026-06-11): this is now the v1.5 "prop room" thesis.** Fix-2's gap (no working local Directus; content
> read live from prod) is exactly what **v1.5 "prop room"** closes — see the v1.5 section (now **Done**). Fix-1's
> DEF-M10-01 stays backlog, **re-signed fresh** at v1.5 design with its user-facing sting removed: v1.5 keeps the
> *asset plane* on prod public links so demos show **real images** without the S3 blob-byte work.

> **v1.6 "stage door" — SHIPPED 2026-06-14** (tag `v1.6`; designed 2026-06-14 via `/developer-kit:design-roadmap`;
> branch `release/01.60-stage-door` merged `--no-ff` → `main`; full detail in `## Done — v1.6` below). The
> **secret-provisioning release** — closes the one stack-lifecycle
> concern with **no owning section / tool / skill / doc**: secrets. Today the whole "make secrets land in the right
> files" job is **one line of shell** (`cp platform/.env → peer` in `ensure-clones.sh`) + manual prose, with the TODO
> written in-tree at `setup_guide.md:447`. v1.6 builds a real mechanism: a new `stack-secrets` extension that ingests
> a secret source (directory **or** zip, default `.agentspace/secrets`) and **provisions every repo of a stack** from
> it, values-blind; a **secret-coverage DNA** (a one-sided harness in the proven `datadna` mold) that *lists and keeps
> listed* the required secrets per repo (`introspect` + `diff`); a coverage gate wired into `/dev-up` + `/demo-up`
> pre-flight; and a closing **field-bake** that proves it by building a compliant secret dir inferred from current
> stack-dev. 4 milestones M27→M30 (DNA+ingest → engine+gate → docs+skill → field-bake). **Tooling + docs only — zero
> platform-repo edits; never commit `.env`; never write prod; no verb ever reads or echoes a secret value.** Gap
> analysis + KB blind-area map + risk register:
> [`.agentspace/scratch/roadmap-research-2026-06-14.md`](../../.agentspace/scratch/roadmap-research-2026-06-14.md).

> **Why v1.6 starts at M27, not M26** (renumber 2026-06-14): the flat milestone counter's **M26 was already consumed**
> by an orphaned ext effort — branch `m26/self-contained-demo` @ `25ab855`, tag `prop-room-m26`, *"make demo stacks
> self-contained (own GitHub clone set, like stack-dev)"* (+521/−141 in `demo-stack/` + `stack-injection/`), authored
> 2026-06-14 in the `rosetta-extensions` authoring copy but **never merged, never pushed, never tracked in this
> roadmap**. v1.6 "stage door" was first designed as M26→M29; on discovering the collision the user chose to **keep
> self-contained-demo as M26** (pending its own roadmap home) and **renumber the secret-provisioning release to
> M27→M30**. M26 is intentionally not detailed here — it awaits a separate `/developer-kit:design-roadmap` pass to give
> it a version + scope.

> **v1.7 "house lights" — SHIPPED 2026-06-15** (tag `v1.7`; designed 2026-06-15 via `/developer-kit:design-roadmap`;
> branch `release/01.70-house-lights` merged `--no-ff` → `main`; full detail in `## Done — v1.7` below). A
> **demo-UI-hardening release** — make a fresh browser at a demo's offset UI render the working app with **zero manual
> steps**. Triggered by a live defect: next-web at `http://localhost:33000` (demo-3) showed a **blank page** because
> clerk-js's handshake to the fake FAPI (`https://127.0.0.1:35400`) hit an **untrusted self-signed cert** → clerk-js
> aborted → blank. **M31** automates a locally-trusted **mkcert** FAPI cert into the demo bring-up (openssl fallback +
> a `DEMO_NO_MKCERT` opt-out; the fake BAPI is plain HTTP → out of scope). **M32** fixes the sibling studio-desk
> `:9100`-dead-redirect (a 1-line `NODE_ENV=production` override root-cause fix + the `:9100` doc/CORS sweep).
> ant-academy demo liveness → **backlog** (repro-first). **Tooling + docs only — zero platform-repo edits.** Gap
> analysis + fix design + risk register:
> [`.agentspace/scratch/roadmap-research-2026-06-15.md`](../../.agentspace/scratch/roadmap-research-2026-06-15.md).

## Version plan

| Version | Codename | Theme | Milestones | Status |
|---------|----------|-------|------------|--------|
| **v1.0** | **body double** | A *measured* stand-in the platform can't tell from the real thing | M0 → M1 → { M1b ∥ M2 } → M2b → M2c | ✅ **SHIPPED 2026-06-03** (tag `v1.0`) |
| **v1.1** | **show floor** | The platform-operations extension framework (demo + dev, in 2 repos) | M3 ✅ → M4 ✅ → M5 ✅ → M6 ✅ → M7a ✅ → M7b ✅ → M7c ✅ → M8 ✅ | ✅ **SHIPPED 2026-06-05** (tag `v1.1`) |
| **v1.2** | **set dressing** | Richer demo worlds — the real *public* taxonomy + content library, measured-faithful, to 100% data-DNA coverage | M9a ✅ → M9b ✅ → M10 ✅ → M11 ✅ | ✅ **SHIPPED 2026-06-07** (tag `v1.2`) |
| **v1.3** | **stack party** | dev + demo stacks as first-class peers — the per-stack-Directus recipe + firewall check (print-only — see the Correction above), auto-snapshot + light seed, smart shared ports, one unified `stack-*` skill set | M12 ✅ → M13 ✅ → M14 ✅ → M15 ✅ | ✅ **SHIPPED 2026-06-07** (tag `v1.3`) |
| **v1.3b** | **dress rehearsal** | Field-hardening — make `/demo-up` produce a full, populated, verified, demoable stack (the gaps the first real run surfaced) | M16 ✅ → M17 ✅ → M18 ✅ → M19 ✅ → M20 ✅ | ✅ **SHIPPED 2026-06-09** (tag `v1.3.1`) |
| **v1.5** | **prop room** | The **local-Directus release** — every stack serves its own captured public catalog locally (data plane local, asset plane prod → real images), content-self-contained on `--local-content` | M21 ✅ → M22 ✅ → M23 ✅ → M24 ✅ → M25 ✅ | ✅ **SHIPPED 2026-06-14** (tag `v1.5`) |
| **v1.6** | **stage door** | The **secret-provisioning release** — one mechanism that ingests a secret source (dir/zip, default `.agentspace/secrets`) and provisions every repo of a stack, with a secret-coverage DNA that lists + keeps-listed the required secrets per repo | M27 ✅ → M28 ✅ → M29 ✅ → M30 ✅ | ✅ **SHIPPED 2026-06-14** (tag `v1.6`) |
| **v1.7** | **house lights** | **Demo-UI hardening** — a fresh browser at a demo's offset UI renders the working app with zero manual steps (the mkcert-trusted FAPI cert so next-web stops blanking + the studio-desk single-port/production fix) | M31 ✅ → M32 ✅ | ✅ **SHIPPED 2026-06-15** (tag `v1.7`) |

> **Why "v1.5", not "v1.4":** v1.4 was removed 2026-06-11 (its seeds → unscheduled backlog). The next release is
> numbered **v1.5** to leave that gap unambiguous — nothing was silently renamed into the v1.4 slot.

The whole initiative layers a **second corpus + skill set on top of** the existing dev-environment
tooling, to build disposable demo environments. Hard constraints: **no modification to any platform
repo** (current or future) and **no disruption to the dev environment**. Each local stack lives in its
own gitignored **`stack-*/`** workspace spanning one full stack — its platform service repos *plus* its
own clone of `rosetta-extensions`: `stack-dev` (dev), `stack-demo` (demo), `stack-dev-2` (secondary
dev), and future `stack-stage` / `stack-tests`. **Policy:** all code/scripts that operate the
corpus/platform on a spawned stack live in `rosetta-extensions` — never scattered in the rosetta corpus,
never authored ad-hoc inside a stack dir. New tooling is built + tested in the authoring copy at
`.agentspace/rosetta-extensions/`, tagged, then consumed per-stack as `stack-<role>/rosetta-extensions @ <tag>`
(rosetta = read-only doc corpus + dev-env skills; `rosetta-extensions` = the executable stack tooling).
Full brief: [`.agentspace/demo-environment-draft.md`](../../.agentspace/demo-environment-draft.md).

## Done — v1.7 "house lights" (SHIPPED 2026-06-15 · tag `v1.7`)

**Theme:** a **demo-UI-hardening release** — when the house lights come up, the audience can see the show: a fresh
browser at a demo's offset UI renders the working app with **zero manual steps**. Triggered by a live defect — next-web
at `http://localhost:33000` (demo-3) rendered a **blank page** because clerk-js's dev-browser handshake to the fake FAPI
(`https://127.0.0.1:35400`) hit an **untrusted openssl self-signed cert** (`net::ERR_CERT_AUTHORITY_INVALID`) → clerk-js
aborted → blank. The investigation (3-investigator workflow + synthesis) found it's a **single-endpoint** trust failure
(the fake BAPI is plain HTTP, server-side only → out of scope) and surfaced a sibling demo-UI rough edge (studio-desk
302s to a dead `:9100`). 2 milestones M31→M32; ant-academy liveness → backlog (repro-first). **Tooling + docs only —
zero platform-repo edits.**

> **Designed 2026-06-15** via `/developer-kit:design-roadmap`. Gap analysis + fix design + risk register:
> [`.agentspace/scratch/roadmap-research-2026-06-15.md`](../../.agentspace/scratch/roadmap-research-2026-06-15.md).
> **Phase 0a deferral audit GREEN** (by inheritance — the v1.6 `/close-release` re-audit ran 2026-06-14; the backlog
> items DEF-M10-01 / DEF-M21-01 / M25-D9 / M26-self-contained-demo are all orthogonal). **Phase 0b KB blind-area:**
> `corpus/ops/demo/recipe-browser-login.md §B` already OWNS the cert story (documents the *manual* mkcert workaround
> today — M31 rewrites it manual→automated); only blind area is `frontend-tier.md` (silent on the cert — add a one-liner).

### M31: mkcert-trusted FAPI cert — the browser-login render fix
**Status:** `done` (completed 2026-06-15) · **Shape:** `section`
**Closure:** Delivered Fate-1 — one branch in `up-injected.sh` step 3a-bis (inside the keep-existing-cert guard): `command -v mkcert && [ "$NO_MKCERT" != 1 ]` → idempotent `mkcert -install || true` + mint `127.0.0.1 localhost ::1`; the openssl fallback factored into `gen_openssl_fapi_cert()` (byte-compatible, called from both the absent/opted-out AND mint-failed branches), non-fatal throughout. `DEMO_NO_MKCERT=1` opt-out parsed (mirrors the `DEMO_NO_*` family). ZERO change to the 3 cert-consumers (`fake-fapi/main.go`, `gen_injected_override.py` mount, `inject.py`) — all path-only (M31-D4, verified). Docs: `recipe-browser-login.md §B` rewritten manual→automatic + the security/remote-VM/Firefox-`certutil`/cert-expiry/`DEMO_NO_MKCERT` caveats; `frontend-tier.md` one-liner; the demo-up SKILL note. The close-time observable verify was proven by composition (M31-D7: chromium default-context trusts the mkcert cert / 200 vs rejects the openssl self-signed / `ERR_CERT_AUTHORITY_INVALID` — the exact blank-page cause — + the earlier cert-trusted→renders proof + the 11 `FapiCertStep` functional/edge tests). Tests: `test_tooling.py` 47→**50** (the `FapiCertStep` class: 8 build + 3 harden edge/regression — install-failure-`||true` mutation-proven, `$CERTS`-whitespace, crt-only partial-state; harness strict-mode fidelity fix M31-D6); Go **1027** unchanged (M31 touched no Go); flake **0** (5/5 randomized sequential). Close found 2 findings, both Fate-1: the `demo-stack/README.md` "13 tests"→**50** count drift (M31-D8) + a recorded adversarial scenario (zero-byte-cert / existence-only guard = pre-existing, documented repair). Deferral audit GREEN (v1.7's first milestone; 0 inherited, 0 repeat). Ext tag `house-lights-m31` @ `6565ef8` (harden `815993f`; close review-fix `5022e72`). Decisions M31-D1..D8 + the Phase-2c adversarial subsection in `decisions.md`.
**Goal:** A fresh browser at demo-N's next-web renders the signed-in app with **zero manual cert-trust / proceed-anyway**, by minting a locally-trusted (mkcert) FAPI TLS cert at bring-up — degrading cleanly to the current openssl self-signed path when mkcert is absent.
**Scope:**
  - In: `demo-stack/up-injected.sh` step 3a-bis (~lines 344-353), inside the existing keep-existing-cert guard — branch on `command -v mkcert` (and a new `DEMO_NO_MKCERT` opt-out): mkcert branch runs idempotent `mkcert -install` (`|| true`) then `mkcert -cert-file/-key-file 127.0.0.1 localhost ::1`; else / on-mint-failure keeps the openssl gen **verbatim**. Non-fatal throughout (the never-abort-a-good-bring-up contract).
  - In: `DEMO_NO_MKCERT=1` escape hatch (parsed in `up-injected.sh`, mirrors `DEMO_NO_UI`/`DEMO_NO_SETDRESS`/`DEMO_NO_LOCAL_CONTENT`) — forces openssl even when mkcert is present, for operators who won't put a dev CA in their trust store.
  - In: **zero change** to `gen_injected_override.py` / `inject.py` / `fake-fapi/main.go` (a trusted cert at the same `<stack>/certs/fapi.{crt,key}` paths "just works" — mount + `FAKE_FAPI_TLS_CERT/KEY` + `ListenAndServeTLS` unchanged).
  - In: docs — rewrite `corpus/ops/demo/recipe-browser-login.md §B` (manual mkcert workaround → automatic; add the dev-CA-in-trust-store security note, the remote/VM + Firefox/`certutil` caveats, the `DEMO_NO_MKCERT` opt-out); a `frontend-tier.md` cert one-liner; the demo-up SKILL browser-login note; the `up-injected.sh:337-342` + `gen_injected_override.py:295` code comments (retire the "operator runs mkcert once" framing).
  - In: a one-line **forward-note** in the code comment that a future dev-N `--local-content` UI path would want the same mkcert wiring (candidate to extract as a shared helper rather than re-inline).
  - Out: the fake BAPI (plain HTTP, server-side only — no browser TLS handshake); the studio-desk redirect (M32); ant-academy liveness (backlog).
**Depends on:** none (first milestone of the version).
**Parallel with:** M32 in principle (different root files), but **sequence M31→M32** — both touch `up-injected.sh` + the same demo doc cluster, so serial avoids a merge overlap.
**Estimated complexity:** small.
**Open questions:** none blocking (fallback = openssl not fail-loud, decided; BAPI out of scope, decided; SANs `127.0.0.1 localhost ::1`, decided).
**KB dependencies:** `corpus/ops/demo/recipe-browser-login.md` (the cert story it owns); `corpus/services/clerkenstein.md` (the browser-login→FAPI handshake); `corpus/ops/demo/frontend-tier.md`.
**Delivers →** `rosetta-extensions/demo-stack` (the mkcert bring-up step; ext tag `house-lights-m31`) + the `recipe-browser-login.md` rewrite.

### M32: studio-desk single-port / production alignment + the `:9100` sweep
**Status:** `done` (completed 2026-06-15) · **Shape:** `section`
**Closure:** Shipped clean. The studio-desk override now pins `NODE_ENV=production` (+ `FRONTEND_PORT=9000`) so the additive env-merge wins back the production `sendFile` path (no dead-`:9100` 302); root-cause precedence verified by code-read + a close-time live merge-probe (#M32-D4). Route coverage verified by code-read — the production block covers every dev-block route via `sendFile` + an `express.static(dist/public)` mount + an `index.html` SPA fallback, **NO GAP** (#M32-D1). The un-offset `:9100` CORS origin dropped (dead; #M32-D2). The `:9100` doc sweep (`frontend-tier.md` single-port story + CORS note + verify registry, demo-up SKILL) is complete. Regression guard: `test_studio_desk_env_pins_node_env_production` (mutation-checked 4 ways, both content paths) + a CORS exact-set assertion (harden — catches over-removal/re-add). Build surfaced 2 latent env-masked YAML-test bugs (fixed Fate-1, #M32-D3). Close-time observable verify satisfied by composition (#M32-D5: the production pin is set → the production path serves on the single 9000 port, no dead-`:9100` redirect, no 404; demo-3 torn down so no end-to-end re-spin). **Tests:** stack-injection `test_injection.py` 87→88 (+1 regression); full suite **88/88** (0 skipped under PyYAML, authoritative JUnit tally); flake **0** (5/5 randomized sequential); Go **1027** unchanged (M32 touched no Go). Review: 4 findings, all Fate-1 (3 decision-tag blends D1/D2/D4 + 1 adversarial-scenario record). Deferral audit **GREEN** (0 in-milestone punts; inherited backlog all cross-release/re-signed). ext tag `house-lights-m32` @ `107599c` (harden `7b17c39`). **v1.7 is now ready for `/developer-kit:close-release`.**
**Goal:** A fresh browser at demo-N's studio-desk (e.g. `http://localhost:39000`) lands on a live page instead of a 302 to the dead `:9100`, by running the container's production code path; and the docs/CORS all agree on single-port `9000`+offset.
**Scope:**
  - In: `gen_injected_override.py` `FRONTENDS` studio-desk dict (~lines 90-96) — add `NODE_ENV=production` (+ `FRONTEND_PORT=9000` belt-and-suspenders). Root cause: the base compose ships `NODE_ENV=development` and the override's per-frontend env block is **additive** (deliberately not `!override`), so `development` survives → `src/index.ts` `isProduction=false` → the dev block `res.redirect('http://localhost:9100/home')` fires (a dead port). Production → `sendFile`, no cross-port redirect.
  - In: a regression assertion in `stack-injection/tests/test_injection.py` near the single-port tests (~820-857) — assert `NODE_ENV=production` in the studio-desk env block.
  - In: a Playwright smoke on `/home` + a couple of studio-desk routes confirming the production `sendFile` path serves them (verify the dev block's `.html`-extension redirects aren't load-bearing).
  - In: the **`:9100` sweep** — `demo-up SKILL` (`:9100+`→`:9000+`), `frontend-tier.md:21` (drop the dead `:9100` frontend port → single-port `9000`+offset), `gen_injected_override.py:249` CORS (remove the un-offset `9100` origin — explicit decision note; dead now that studio-desk is single-port production).
  - Out: the cert (M31); ant-academy (backlog).
**Depends on:** none functionally; **sequence after M31** (shared `up-injected.sh`/doc-cluster surface).
**Parallel with:** none (sequence after M31).
**Estimated complexity:** small.
**Open questions:** confirm the production `sendFile` path covers ALL routes the dev block handled (the Playwright smoke proves it).
**KB dependencies:** `corpus/ops/demo/frontend-tier.md` (the studio-desk port story); demo-up SKILL.
**Delivers →** `rosetta-extensions/demo-stack` + `stack-injection` (the override fix + test; ext tag `house-lights-m32`) + the `:9100` doc/CORS sweep.

### Execution graph — v1.7 "house lights"
```
v1.7 "house lights"   (sequential — shared up-injected.sh + demo doc cluster)
  M31 ─────────→ M32
  mkcert cert    studio-desk single-port/production + :9100 sweep
```
**Parallelism:** none. Both are small `section` fixes that touch `up-injected.sh` + the same demo doc cluster, so they
run serially to keep the merge surface clean. **ant-academy liveness → backlog** (repro-first; not in v1.7 scope).

### Risk map — v1.7 "house lights"
| Risk | Severity | Mitigation | Owner |
|---|---|---|---|
| Fresh-machine `mkcert -install` **prompts for the OS password** (GUI keychain write) → blocks an autonomous/CI bring-up | degrades-quality | `\|\| true` + the whole step non-fatal + openssl fallback; document "zero manual steps **on a local same-machine demo**", not universally | M31 |
| **Remote/VM demos:** `-install` trusts the bring-up machine, not the browsing machine → still blanks | degrades-quality | keep the self-signed/proceed-anyway fallback documented for remote demos; don't claim universal zero-touch | M31 |
| A **dev CA in the trust store** is a real trust expansion (Firefox needs `certutil`) | nice-to-resolve | the `DEMO_NO_MKCERT` opt-out + a security note in `recipe-browser-login.md` | M31 |
| **Cert expiry:** the keep-existing guard has no expiry check → a ~3y mkcert leaf could silently re-blank | nice-to-resolve | document `rm <stack>/certs/fapi.crt` regenerates; note in the spec | M31 |
| **M32 `NODE_ENV=production`** turns off the dev redirect block → some studio-desk routes could 404 if the production `sendFile` path doesn't cover them | degrades-quality | the in-scope Playwright smoke proves all routes serve before close | M32 |

## Done — v1.6 "stage door" (SHIPPED 2026-06-14 · tag `v1.6`)

**Theme:** the **secret-provisioning release.** Every other stack-lifecycle concern in Rosetta has an owning
`rosetta-extensions` section + a tool + a skill + a corpus doc (snapshot→`stacksnap`/`/stack-snapshot`/`snapshot-spec.md`;
seeding→`stackseed`/`datadna`/`/stack-seed`/`seeding-spec.md`; verify→`/test-platform`/`verification.md`). **Secrets have
none of these.** The whole "make secrets land in the right files" job is split between manual prose in `setup_guide.md`
(hand-copy `.env.example`→`.env` and fill Clerk/AI keys for studio-desk, ant-academy, next-web-app) and **one line of
shell** — the copy-if-absent `cp stack-dev/platform/.env → peer/platform/.env` in `demo-stack/ensure-clones.sh:48-57`.
The TODO is even written in-tree: `corpus/ops/setup_guide.md:447`. v1.6 builds the missing mechanism: ingest a secret
source (**directory or zip**, default `.agentspace/secrets`) → **provision every repo of a stack** from it (values-blind)
→ a **secret-coverage DNA** that *lists and keeps listed* the required secrets per repo → a coverage gate in bring-up
pre-flight → a field-bake that proves it from current stack-dev.

> **Designed 2026-06-14** via `/developer-kit:design-roadmap`. Phase-1 research was a **dynamic workflow**: 5 parallel
> read-only investigators (secret landscape · current provisioning flow · alignment/DNA framework · extension placement +
> KB blind-area · stack-dev inventory) + a synthesis/completeness-critic stage. Gap analysis + risk register:
> [`.agentspace/scratch/roadmap-research-2026-06-14.md`](../../.agentspace/scratch/roadmap-research-2026-06-14.md).
> **Phase 0a deferral audit GREEN** by inheritance (the v1.5 `/close-release` deferral re-audit ran 2026-06-14; the 3
> live backlog items — DEF-M10-01, DEF-M21-01, M25-D9 — are all orthogonal to secret provisioning). **Phase 0b KB
> blind-area: CONFIRMED** — no corpus doc owns secret provisioning → M29 `Delivers →` the net-new `corpus/ops/secrets-spec.md`.

> **Load-bearing facts** (the critic caught + corrected the investigators' own errors; re-verified on disk 2026-06-14):
> there is **no single 57-key platform/.env** — live truth is `platform/.env` (15 keys) + a *separate* `app/.env`
> (46 keys) + per-frontend envs (`studio-desk/.env` 23, `next-web-app/apps/web/.env` 32), with `platform/.env_example`
> the 59-key wishlist. **8 of 12 Go-service repos ship NO `.env.example`** (only sentinel does) → the required-set must
> be a **hybrid** source, not a uniform per-repo file. **ant-academy's live env is `code/.env.local`** (not `code/.env`)
> → target filenames must be pinned. **~9–12 distinct secret VALUES** stand up a working stack.

> **Design constraints, baked into every milestone:** **(1) values-blind** — no verb ever reads, copies into a report,
> echoes, or logs a secret *value*; the secret-DNA stores key NAMES only and is committable; extraction is
> `grep -oE '^[A-Z_][A-Z0-9_]*'` / cut-on-`=`. **(2) Never commit `.env`; never write prod; zero platform-repo edits.**
> **(3) Reuse the proven patterns** — the secret-DNA is a *sibling* harness in the `stack-seeding/dna` (`datadna`) mold
> (Criticality 3/2/1 → weight, Overall-weighted + Critical-unweighted score with the `ratio()` empty-denominator +
> anti-vacuous-100 guards, the 0/1/3 exit-code contract), and `provision` **composes with and defers to** the existing
> injection override + emits `PreflightEnv`-passing env (`safety.md:156-205`). **(4) The two-repo split holds:**
> `rosetta-extensions` owns the `stack-secrets` section (authored in `.agentspace/rosetta-extensions/`, tagged
> `stage-door-mNN`, consumed per-stack at the pinned tag); `rosetta` owns `.claude/skills/*`, `corpus/*`, `CLAUDE.md`.

### M27: Secret-coverage DNA + source ingestion
**Status:** `done` (closed 2026-06-14) · **Shape:** `section`
**Dir:** [m27-secret-coverage-dna/](releases/archive/01.60-stage-door/m27-secret-coverage-dna/) (overview · progress · decisions · spec-notes · metrics · retro · audit-deferrals)
**Closure:** all 12 deliverable sections landed Fate-1 — the new `stack-secrets` extension ships a **values-blind
secret-coverage DNA** (55 genes / 6 repos: platform, app, sentinel, studio-desk, next-web-app, ant-academy), a
**DNA-driven** dir+zip source reader that **structurally cannot** ingest a `zEnvs/` backup mirror or a stray
`.env` (it opens exactly `<root>/<repo>/<target_file>`, never enumerates the tree — the layout-contract defence),
the hybrid `introspect` (`platform/.env_example` + sentinel + each frontend's `.env.example` + a curated
compose/build-arg set — verified 8 of 9 Go repos ship no `.env.example`), the DNA-scoped **two-tier keep-listed
`diff` gate** (M27-D2: gate-fatal only on an already-tracked secret omitted for a repo → vacuously-green coverage;
never-tracked keys → informational triage), and the `stacksecrets` CLI (`list`/`check`(=`measure`)/`introspect`/
`diff`). The **`check`/`measure` scorer was folded in Fate-1** (M27-D3.2 — the natural pairing with the DNA;
`provision` + pre-flight wiring + demo-aware scoring stay M28). **HARD SAFETY held throughout:** no verb reads,
echoes, logs, or persists a secret VALUE — only `ClassifyShape` touches a value (as a discarded local), the
committed `secret-dna.json` carries zero secret-shaped tokens; verified end-to-end against the real stack-dev
(`diff` exits 0, `check` reports coverage with no value printed). The build's diff-vs-stack-dev caught **10 real
cross-repo DNA omissions** → fixed Fate-1. **Stdlib-only** module (no `go.sum`) so the values-blind audit surface
is trivially small. **Close:** 2 findings — 0 scope · 0 code-quality · 1 docs (DOC-1 the ext README Sections table
missed the `stack-secrets` row → Fate-1 fixed) · 1 tests (TEST-1 the section README quoted a stale "94 tests" →
reconciled to 113) · 0 decision-blend (the feature's corpus doc is M29 scope per the repo-split). 4 adversarial
scenarios recorded (all already test-pinned). Deferral audit **GREEN** (M27 is the first v1.6 milestone — no repeat
possible; DEF-M27-01 encrypted-zip → **DROP** as a documented v1 boundary, DEF-M27-02 per-gene-profile-tag →
**Fate-2** M28-owned-if-needed, inherited DEF-M10-01/DEF-M21-01/M25-D9 → **KEEP** re-signed at v1.5 close). Go
867→**980** (+113, entirely the new section; build + 2-pass harden); Python **459** unchanged; flake **0**; 5/5
`-shuffle` clean; `-race` + `gofmt` + `go vet` clean. **Code:** `rosetta-extensions` @ tag `stage-door-m27`
(ext head `195ef93`; close doc-hygiene commit `537aeff` on the `m27` branch, orchestrator-finalized).
**Goal:** A new `stack-secrets` extension that ingests a secret source (directory or zip) and a secret-coverage DNA that *lists and keeps listed* the required secrets per repo — values-blind throughout.
**Scope:**
  - In: new `rosetta-extensions/stack-secrets/` section + `go.mod` + `cmd/stacksecrets`, authored in the `.agentspace` copy, tagged.
  - In: **source ingestion** — directory **and** zip (default `.agentspace/secrets`); values-blind extraction; an explicit source-dir **layout contract** so `stack-dev/zEnvs/` and per-repo `.env` files are never silently ingested as "the source".
  - In: the **secret-DNA** sub-package (mirrors `stack-seeding/dna` layout): gene = `repo × KEY`, gene id `<repo>/<KEY>`; per gene `{repo, key, scope (shared|service|frontend|config), criticality (critical|standard|optional → weight 3/2/1), operators [key-present (+optional nonempty, format:url|jwt|pk_*|sk_*)], status (required|optional|waived-<reason>), source_hint, note}`.
  - In: `introspect` rebuilds the required set from the **hybrid** source (`platform/.env_example` baseline + each frontend's + `sentinel`'s `.env.example` + the keys docker-compose injects/references) — NOT pure per-repo `.env.example`.
  - In: `list` + `diff` verbs; `diff` exits 1 on required-key drift (the **"keep-listed" gate**) and on a runtime-required-but-undeclared key (the anti-vacuous-green guard).
  - In: model the knowns as `waived` — AWS-via-`~/.aws` mount, profile-gated keys (BREVO/messenger, customerio-sync), optional Bunny/GCloud; encode **alias families** (`GH_PAT`≡`GH_TOKEN`≡`GH_ACCESS_TOKEN`) vs distinct-similar values (`OPENAI_KEY` vs `OPENAI_API_KEY` — list exact per-repo, do NOT auto-alias).
  - Out: writing target `.env` files (M28); the coverage gate + bring-up wiring (M28); the skill + corpus doc (M29).
**Depends on:** none (first milestone of the version).
**Parallel with:** none.
**Estimated complexity:** medium-large.
**Open questions:** zip ingestion mode (extract-to-temp vs in-memory; encrypted-zip via age/gpg — default: plain zip + dir in v1, encrypted deferred); profile-tagging on genes vs default-`graphql`-profile scoping (settle here or in M28).
**KB dependencies:** `corpus/architecture/alignment_testing.md` (the DNA framework + the data-DNA precedent); `corpus/ops/seeding-spec.md` (the `datadna`/`stackseed` patterns); `corpus/ops/safety.md` (`PreflightEnv` discipline).
**Delivers →** `rosetta-extensions/stack-secrets/` (the section; ext tag `stage-door-m27`).

### M28: Provisioning engine + coverage/verify gate
**Status:** `done` (closed 2026-06-14)
**Shape:** `section`
**Goal:** `stacksecrets provision` writes each repo's target `.env` from the source (correct exact key per repo, alias-mapped per file), values-blind; `check`/`measure` computes coverage and is wired non-fatally into `/dev-up` + `/demo-up` pre-flight.
**Scope:**
  - In: `provision` verb — per-repo **target-file map** (`platform/.env`, `app/.env`, `studio-desk/.env`, `ant-academy/code/.env.local` [exact filename pinned], `next-web-app/apps/web/.env`, `sentinel/.env`), one source value → all its per-file aliases.
  - In: **idempotency + overwrite policy** — copy-if-absent default, `--force` to overwrite, never silently clobber; **N=0 main-dev-stack guard** (refuse without `--force`, mirroring `stackseed --reset`) so it can't clobber the operator's working `.env`.
  - In: **composes with + defers to the injection override** — `provision` runs BEFORE `gen_injected_override.py` and must NOT re-arm the stripped prod `DIRECTUS_TOKEN` on non-prod / `--local-content` stacks (the fix16/17 safety class) **[blocks-release safety — regression test required]**.
  - In: emit `PreflightEnv`-passing env (reuse the seeder's values-blind env-guard discipline).
  - In: `check`/`measure` — Overall (weighted) + Critical (gate == 100%, unweighted) + **per-repo rollup** ("repo X is short key Y"); **demo-aware** (Clerk keys satisfiable by Clerkenstein minting, not the source dir); exit 1 if a critical key is missing.
  - In: non-fatal pre-flight wiring into `/dev-up` + `/demo-up` (warn on standard-missing, fail on critical-missing — the `verification.md` convention).
  - In: profile-scoping decision settled (v1 scopes the denominator to the default `graphql` profile, or a per-gene profile tag).
  - Out: the `/stack-secrets` skill + corpus doc (M29); the end-to-end build-from-stack-dev validation (M30).
**Depends on:** M27 (the DNA + the ingestion reader).
**Parallel with:** none.
**Estimated complexity:** large.
**Open questions:** `check` measure source — static `.env` files (safe, names-by-grep) vs live container env (catches runtime-injected keys but risks touching values) → default static, values-blind.
**KB dependencies:** `corpus/ops/safety.md` (`PreflightEnv`, never-write-prod); `corpus/ops/idempotency.md` (the run-it-twice contract); `corpus/ops/verification.md` (non-fatal pre-flight); `corpus/services/clerkenstein.md` (the demo minting path); `corpus/ops/rosetta_demo.md` (the injection override / `DIRECTUS_TOKEN` strip).
**Delivers →** `rosetta-extensions/stack-secrets/` (the engine + gate; ext tag `stage-door-m28`).
**Closure (2026-06-14):** Delivered all 12 boxes. `stacksecrets provision` writes each repo's target `.env` from the
source values-blind — grouped by `(repo, target_file)` from the DNA, alias-mapped per file (gh-token family incl.
cross-repo `app/GH_TOKEN`; distinct-similar pairs never auto-copied), append-only copy-if-absent (M28-D2), `--force`
to overwrite, N=0 main-dev-stack refusal mirroring `stackseed --reset`. The headline blocks-release safety landed +
test-pinned: provision runs BEFORE the injection override and NEVER re-arms the stripped prod `DIRECTUS_TOKEN` on a
non-prod stack — it writes the `StripOnNonProdKeys` family BLANK, the exact state the override forces, so the base
`.env` + override agree and a blank value still passes the key-present-only gene (M28-D3; `TestProvision_NeverReArms‑
DirectusTokenOnNonProd` + the `--force` / pre-existing-token edges). The single value-carrying boundary is
`provision/io.go::sourceValues`; a reflection-walk safety test (incl. unexported fields + map keys) proves no value
surfaces in the Report / dry-run plan / errors. `check`/`measure` became demo-aware (`MeasureForStack` + a `mintedSource`
overlay so Clerkenstein-minted Clerk keys count without the source carrying them) and wired non-fatally into `/dev-up`
(`dev-stack` cmd_up) + `/demo-up` (`up-injected.sh`) pre-flight via the shared `stack-secrets/preflight.sh` (warn
standard / fail critical / skip otherwise; `DEV_/DEMO_NO_SECRET_PREFLIGHT=1` opt-outs). M28-D1: the LiveKit key/secret
were de-aliased (a credential pair = two distinct values, not one). The base scorer + profile-scoping were M27
(M27-D3.2) — reused unchanged. **Decisions:** M28-D1 (LiveKit de-alias), M28-D2 (append-only values-blind write),
M28-D3 (DIRECTUS_TOKEN non-rearm = write blank). **Hardening:** 3 passes; provision 87.3%→94.8%, secretdna 99.2%; +1
real bug fixed inline (the `preflight.sh` empty-array crash under `set -u` on bash 3.2 — the macOS system bash — fixed
with the `${arr[@]+"${arr[@]}"}` conditional-expansion guard + a `/bin/bash`-3.x regression; the shell-portability
invariant backfilled into `safety.md §2.8`). **Close review GREEN:** 1 finding, Fate-1 ext code-quality — the demo
secret pre-flight block sat ABOVE the `UP_INJECTED_LIB_ONLY` lib-only seam in `up-injected.sh`, so sourcing the script
lib-only (the `test_frontend_build.py` unit tests, in a sandbox without the sibling `preflight.sh`) fired the bring-up
pre-flight at source time and crashed 20 tests; moved below the seam alongside the M19 VM pre-flight + pinned with a
static positional regression (ext `9742126`). Deferral audit GREEN (0 new; DEF-M27-02 per-gene-profile-tag Fate-2
**discharged** — M28 used the default `graphql` profile, the conditional never triggered; inherited backlog re-signed).
**Tests:** Go 980→**1027** (+47, all stack-secrets); demo-stack 99 pass / dev-stack 74 pass; flake **0** (Go 5/5 `-race
-shuffle`, Python 5/5 sequential); gofmt/go vet/shellcheck clean. Code: `rosetta-extensions` @ build tip tag
`stage-door-m28` (ext head `9742126`; 3 harden + 1 review-fix ahead of the tag — orchestrator finalizes ext-side).

### M29: Docs + `/stack-secrets` skill + corpus wiring
**Status:** `done` (closed 2026-06-14)
**Shape:** `section`
**Dir:** [m29-secrets-docs-skill/](releases/archive/01.60-stage-door/m29-secrets-docs-skill/) (overview · progress · decisions · spec-notes · kb-fidelity-audit · metrics · retro · audit-deferrals)
**Closure (2026-06-14):** Delivered all 8 boxes, rosetta-only. Authored **`corpus/ops/secrets-spec.md`** (net-new, 290 lines — the secret-provisioning source-of-truth: the source-dir/zip layout contract [the zEnvs defence], the 6-repo/55-gene secret-DNA [40 required · 8 optional · 7 waived · 13 critical, profile `graphql`], the per-repo target-file map [`ant-academy → code/.env.local`, `next-web-app → apps/web/.env` pinned], the values-blind safety statement, alias-family [`gh-token` × 3] vs distinct-similar rules [LiveKit pair, OPENAI_KEY/OPENAI_API_KEY], the waived class, the `DIRECTUS_TOKEN` non-rearm safety, the `0/1/3` exit contract); the **`/stack-secrets`** skill (mirrors `/stack-seed`: read spec → confirm non-prod target → build the pinned-tag `stage-door-m28` binary → run the verb → report values-blind); the **CLAUDE.md** skill-table row + Key-Documentation-Locations entry + Interconnected-Documentation rows 10/11 + both corpus index rows; **`safety.md` §2.9** (the values-blind / `PreflightEnv`-emitting / `DIRECTUS_TOKEN`-non-rearm clause); and the **`setup_guide.md`** retire-prose (the `/stack-secrets` fast-path callout + retired hand-copy for studio-desk/ant-academy/next-web + the **line-447 TODO deleted**, keeping the per-repo key lists as reference per M29-D4, keeping the root `platform/.env_example → .env` copy). **Decisions:** M29-D1 (build from pinned tag `stage-door-m28`, not clone HEAD), M29-D2 (skill `--check|--provision|--status` shorthand → real CLI subcommands), M29-D3 (README-index guard checks the same-dir README → indexed in `corpus/ops/README.md`), M29-D4 (setup_guide keeps the per-repo key lists, retires only the `cp` mechanics; ant-academy target corrected to `code/.env.local`). **Close review GREEN — 0 findings** across scope/code-quality/docs/tests/decision-triage. Every load-bearing doc claim re-verified against the ext engine at tag `stage-door-m28` (`9742126`): the 55/6/40-8-7/13-crit DNA, the `gh-token` 3-member alias family, `StripOnNonProdKeys` (3 keys), `MintedKeys` (6 keys), `ClassifyShape` + the `provision/io.go` value boundary + `provision_safety_test.go`, every `stacksecrets` CLI flag/subcommand. README-index guard exit 0; cross-refs resolve (16/0). 4 adversarial doc-consumer scenarios recorded (all clean). Deferral audit **GREEN** (0 new · 0 repeat · 0 aged-out; DEF-M27-01 dropped + DEF-M27-02 discharged from prior closes; 3 inherited backlog [DEF-M10-01/DEF-M21-01/M25-D9] re-signed at v1.5 close, carry). **Zero ext code** — the ext stayed on `main` @ `9742126` (= tag `stage-door-m28`), untouched; no new ext tag. **Tests:** Go **1027** / Python **459** unchanged (M29 touches no code); flake **0**. **Code:** none (rosetta markdown/text only).
**Goal:** Make the feature discoverable + own a corpus doc: author `corpus/ops/secrets-spec.md`, add the `/stack-secrets` skill + CLAUDE.md skill-table row, retire the manual-copy prose + the `setup_guide.md:447` TODO, extend `safety.md`.
**Scope:**
  - In: author **`corpus/ops/secrets-spec.md`** (the blueprint the skill reads — source-dir/zip layout, the secret-DNA, the per-repo target-file map, the values-blind safety statement, the alias/collision rules, the `waived`-class rationale).
  - In: new **`/stack-secrets`** skill (`argument-hint [dev-N|demo-N] [--from DIR|ZIP] [--check|--provision|--status]`, default source `.agentspace/secrets`), mirroring `/stack-seed` (read spec → confirm non-prod target → build the tagged-clone binary → run the verb → report values-blind).
  - In: **CLAUDE.md** skill-table row + Key-Documentation-Locations entry + Interconnected-Documentation list update.
  - In: **extend** `setup_guide.md` (delete the manual-copy prose + the line-447 TODO, point to the skill) and `safety.md` (add the never-echo / `PreflightEnv`-emitting clause to the safety contract).
  - Out: the build-from-stack-dev observable-behavior validation (M30).
**Depends on:** M28 (the engine + gate the docs/skill describe).
**Parallel with:** none.
**Estimated complexity:** medium.
**Open questions:** one skill (author + measure) vs an `/align`-style pair — default one skill, pre-flight `check` rides inside `/dev-up`+`/demo-up`.
**KB dependencies:** `corpus/ops/seeding-spec.md` + `corpus/ops/snapshot-spec.md` (the spec-doc + skill pattern to mirror); `CLAUDE.md` (skill-table + doc-index conventions).
**Delivers →** `corpus/ops/secrets-spec.md` (net-new) + `.claude/skills/stack-secrets/` + CLAUDE.md/setup_guide.md/safety.md edits.

### M30: Field-bake — build a compliant secret dir from stack-dev + prove it
**Status:** `done` (closed 2026-06-14)
**Shape:** `section`
**Closure (2026-06-14):** Delivered all 7 boxes; the bake ran end-to-end. **Part 1:** assembled a compliant,
gitignored `.agentspace/secrets` dir from current stack-dev (5 repo `.env` files cp'd into the DNA-driven
reader layout, values-blind; ant-academy source filled with the shared Clerk publishable key via a
values-blind line-append), and `check` scored **Critical == 100%** on both dev (`check`) and demo
(`check --demo`) — exit 0. **Part 2 (live):** a fresh **demo-3** was brought LIVE from that assembled source
with the user's go-ahead — provision wrote 26 / blanked 2 / skipped 0, the bring-up's pre-flight scored
Critical **100%**, and the demo came up with **17 containers** (backend tier + UI tier: next-web + studio-desk
+ ant-academy native), the full UI-inclusive auto-verify exit 0. The **observable-behavior gate**
(provision → Critical 100% → stack UP) is **MET LIVE**. **SAFETY verified live:** the prod `DIRECTUS_TOKEN`
(len-32) armed in **ZERO** containers (cms blank; provision blanks the family + the injection override strips
it — defense-in-depth, the fix16/17 non-rearm class). The bake caught + fixed **2 real release bugs** Fate-1
(parallels v1.5 M25's 4): (1) `sentinel/DB_CONNECTION` was critical/required but is compose-injected config
(hardcoded `environment:` entry, never read from a `.env`; no `sentinel/.env` exists) → reclassified
`waived-config` + a regression assertion (was falsely failing the gate at Critical 84.6%); (2) the demo
bring-up only *checked* coverage but never *provisioned*, and `preflight.sh` resolved its source path one level
too shallow (doubled `.agentspace/.agentspace/secrets` → the demo gate silently skipped, exit 2) → added a
non-fatal provision step (`DEMO_NO_PROVISION=1` opt-out) + corrected the path to `EXT_ROOT/../..`.
**Decisions:** M30-D1 (sentinel waive), M30-D2 (provision-then-move-only-the-env-file design), M30-D3
(preflight two-levels-up path). **Honesty residual** documented (`spec-notes.md`): the ~10–15% non-passing
is entirely waived-class (`waived-config`, `waived-aws-mount`, `waived-profile-gated`, `waived-optional`) +
standard/optional lean-env/compose-injected/repo-local shorts — zero critical short. **Close review:** 4
findings — 0 scope · 0 code-quality (the ext fixes correct + green) · 3 docs (the milestone record +
`secrets-spec.md` were stale vs the executed live bake + the M30 DNA fix → reconciled: version
`stage-door-m27`→`m30`, sentinel `waived-config`, status split `40/8/7`+13-crit → `39/8/8`+12-crit) · 1
decision-triage (empty `decisions.md` → authored) — all Fate-1. 5 adversarial scenarios recorded. Deferral
audit **GREEN** (0 new/repeat/aged; the 2 bugs landed Fate-1; waived-not-deferred; 3 inherited backlog
re-signed at v1.5 close). **Tests:** Go **1027** / Python **459** unchanged (the M30 ext regression is a
sub-assertion inside an existing test func; +0 top-level); demo-stack 99 pytests pass; flake **0** (Go 5/5
`-race -shuffle`); gofmt/vet/shellcheck clean; corpus index guard exit 0. **Code:** `rosetta-extensions` @
tag **`stage-door-m30`** (ext head `29c922b`; 2 field-fix commits on `m30/field-bake` — orchestrator finalizes
the ext side). **v1.6 is now ready for `/developer-kit:close-release`.**
**Goal:** Prove the whole mechanism on a real stack: assemble a compliant `.agentspace/secrets` dir inferred/pulled from current stack-dev, run `provision` into a fresh stack, and assert the observable behavior (coverage Critical == 100%, the stack comes up).
**Scope:**
  - In: **assemble a compliant `.agentspace/secrets` dir** from current stack-dev (names-correct, alias-mapped, the knowns `waived`) — the user's explicit "build a secret dir compliant with what's requested, inferring/pulling from stack-dev to double-check it works".
  - In: run `provision` into a fresh `dev-N` and a fresh `demo-N`; assert `measure` Critical == 100% and the stack reaches UP (the **observable-behavior gate**, mirroring v1.5's M25 field-bake).
  - In: fix any real bugs the bake surfaces, Fate-1 (the v1.5 field-bake caught + fixed 4).
  - In: document the **honesty residual** — which ~10–15% is `waived` (AWS-mount, optional Bunny/GCloud, profile-gated) and why that's correct, not a hole.
  - Out: new features (bake is a proving + fix milestone, not a feature milestone).
**Depends on:** M29 (the skill + doc the bake exercises).
**Parallel with:** none.
**Estimated complexity:** medium.
**Open questions:** which N to bake on (a throwaway `dev-N≥1` + a `demo-N`, never N=0) — settle at bake time per box capacity.
**KB dependencies:** `corpus/ops/secrets-spec.md` (the contract being proven); `corpus/ops/rosetta_demo.md` + `corpus/ops/setup_guide.md` (bring-up); `corpus/ops/verification.md` (the gate).
**Delivers →** the proven `.agentspace/secrets` reference dir + the field-bake record; ext tag `stage-door-m30`.

### Execution graph — v1.6 "stage door"
```
v1.6 "stage door"   (strictly sequential — no parallelism; each milestone consumes the prior's output)
  M27 ─────────→ M28 ─────────→ M29 ─────────→ M30
  DNA + ingest   engine + gate  docs + skill   field-bake (build-from-stack-dev)
```
**Parallelism:** none. M28 consumes M27's DNA + ingestion; M29 documents M28's engine; M30 proves M29's skill end-to-end.
This mirrors v1.5's strictly-sequential M21→M25. **No B-milestone** — the inspection tooling (`check`/`status`/`diff` +
the per-repo rollup) is folded into the core verbs, not a separate subsystem.

### Risk map — v1.6 "stage door"
| Risk | Severity | Mitigation | Owner milestone |
|---|---|---|---|
| Provisioner **re-arms the prod `DIRECTUS_TOKEN`** that fix16/17 disarmed → reopens a closed tenant-data-leak class | **blocks-release** | `provision` runs before + defers to the injection override; never writes the prod token on non-prod / `--local-content`; regression test | M28 |
| `OPENAI_KEY` vs `OPENAI_API_KEY` (and Azure variants) **may hold different tokens** → naive aliasing provisions the wrong key | degrades-quality | DNA lists the exact per-repo key; distinguish alias-families from distinct-similar values; never auto-alias the ambiguous pairs | M27/M28 |
| **N=0 clobber** — provisioning into the main dev stack overwrites the operator's working `.env` (the secret source) | degrades-quality | copy-if-absent default + `--force` + N=0 guard (mirror `stackseed --reset`) | M28 |
| **Vacuously-green coverage** — if the DNA under-lists a required key, `measure` reports 100% on a short stack | degrades-quality | `introspect`+`diff` exhaustiveness; make "runtime-required but undeclared" an explicit `diff`-flaggable case | M27 |
| **Demo minting vs coverage** — demos satisfy Clerk keys by Clerkenstein minting, not the source dir | degrades-quality | stack-type-aware `measure` (real for dev, minted-OK for demo) | M28 |
| **zEnvs / stray `.env` ingested** as the source by accident (stale backup mirror) | degrades-quality | explicit source-dir layout contract; never auto-discover `zEnvs`/per-repo `.env` | M27 |
| **Profile-awareness** changes the denominator (messenger/customerio/frontend profiles) | nice-to-resolve | v1 scopes to default `graphql` profile (or per-gene profile tag) | M27/M28 |

## Done — v1.5 "prop room" (SHIPPED 2026-06-14 · tag `v1.5`)

**Theme:** every stack today reads its public content **live from prod** (`DIRECTUS_BASE_ADDR=content.anthropos.work`)
— v1.3/M13 shipped only a **print-only** per-stack-Directus recipe, never a running one (corrected corpus-wide
2026-06-11). v1.5 "prop room" makes the prop room real: it stands up a **local Directus per stack**, serving the
**captured public library** (the same snapshot the taxonomy already replays from), so a stack's content is
**self-contained** — no live prod dependency at runtime. **Real images are preserved** by keeping the *asset plane*
on prod's public links (cms already mints token-less `<DIRECTUS_PUBLIC_BASE_ADDR>/assets/<uuid>` URLs that browsers
fetch anonymously — verified `cms/internal/directus/directus.go:87`); only the *data plane* (catalog rows, served
via the local Directus) goes local. **Coverage:** **every demo** stack gets it by default; **any additional dev
stack** (`dev-N≥1`) gets it **opt-in** via a flag; the main dev stack (**N=0**) stays manual (documented opt-in
recipe) — its Directus address is baked in the platform compose and the n=0 guard keeps automation off the
developer's primary box.

> **Designed 2026-06-11** via `/developer-kit:design-roadmap` (the first version staged after the v1.4 removal).
> 4 research agents (deferral sweep · code & gap · knowledge audit · git history) over the corpus @ `a4681cb` + the
> `rosetta-extensions` authoring copy, every claim re-checked against live code. Gap analysis + KB blind-area map:
> [`.agentspace/scratch/roadmap-research-2026-06-11.md`](../../.agentspace/scratch/roadmap-research-2026-06-11.md).
> **Phase 0a deferral audit GREEN** (after per-item user fates 2026-06-11 —
> [`.agentspace/scratch/deferral-audit-2026-06-11.md`](../../.agentspace/scratch/deferral-audit-2026-06-11.md)):
> NEW-1/NEW-2/NEW-3 → v1.5 core (M21–M23); 4 small hygiene items → M24; DEF-M10-01 (S3 blob bytes + cloud store)
> **re-signed → backlog**, its user-facing sting removed by the real-images-via-prod-links posture; the ex-v1.4 seeds
> (AI content, shareability, more mirrors) + the deploy/injection CI gate + the dev-up pre-warm question **DROPPED
> from tracking** at user instruction. **Phase 0b KB blind-areas:** structure-capture design, per-stack Directus
> container lifecycle, and Directus verify probes have **no** corpus anchor → M21/M22 each `Deliver →` the missing
> doc (chiefly the net-new `corpus/ops/directus-local.md`).
>
> **The corpus already names this end-state verbatim** (`snapshot-spec.md:402-406`): *"automate the recipe (execute
> bootstrap, close the M10 collection-schema gap with a capture-side structure extension — the DDL + the
> `directus_collections`/`fields`/`relations` registry rows — then replay + boot + re-point `DIRECTUS_BASE_ADDR`
> per-stack) so content + taxonomy become a **referentially-closed captured pair** for both stack types."* v1.5
> versions that close.

> **Two user constraints, baked into every milestone** (2026-06-11): **(1) never touch production data or platform
> repos** — v1.5 is tooling + docs only, the platform repo is a build *context* at most, capture stays read-only /
> bounded / public-only behind `AssertPublicOnly`, and writes only ever hit per-stack-isolated offset targets;
> **(2) keep the code simple and maintainable** — prefer the existing mechanisms (the generic `CopyIn`/replay path,
> the injected-override generator, the one shared `dev-setdress` engine) over new subsystems; **make the per-stack
> Directus a compose service in the stack's override**, not a hand-managed `docker run` (so existing teardown,
> port-registry, and verify conventions cover it for free — no bespoke lifecycle code).

> **The two-repo split holds** (as in v1.3b): **`rosetta-extensions`** owns all scripts / Go / Python / its own KB
> (authored in `.agentspace/rosetta-extensions/`, tagged `prop-room-mNN`, consumed per-stack at the pinned tag);
> **`rosetta`** owns `.claude/skills/*`, `corpus/*`, `CLAUDE.md`. Each milestone's `Delivers →` line splits accordingly.

### M21: Structure capture — close the collection-schema gap
**Status:** `done` (closed 2026-06-13) · **Shape:** `iterative`
**Closure (closed-on-gate):** the exit gate is **MET by tooling** at iter-08 — `stacksnap` captures the content-model
structure (all-26-collection DDL + **PRIMARY KEYs** + sequences + the `directus_collections` registration + the
public-policy `directus_permissions` read grant) behind a new firewall **structural-metadata admissibility class**
(`AssertStructuralMetadata`: admit `directus_*` system tables as "structure, not tenant data" iff zero tenant-scope
columns — extend, never loosen), **auto-provisions** a bootstrapped-gap stack before the row replay, and a
booted+provisioned stack **serves the captured catalog anonymously over HTTP** with no hand SQL. **8 iters** (1
bootstrap tok + 7 tiks, 0 triggered toks — every tik advanced the 6-stage pipeline: static 2 → live 2 → 4 → 6
demonstrated → code-ified met). The load-bearing empirical finding was the **PRIMARY-KEY rule** (Directus silently
403s a PK-less collection even for admin while the digest still converges — exactly why this was iterative). Digest
convergence resolved via **option A** (capture all 26 collections + pin 11.6.1, whose system digest `b4cb55bc…`
equals prod's — no skew; operator decision #M21-D7). Redefined `stacksnap` exit codes (empty→4 / gap+structure→0 /
gap-no-structure & diverged→5). **Delivered:** `corpus/ops/directus-local.md` (net-new — the structure half: bootstrap
empirics + structure-capture model + version-skew rule + the firewall carve-out; M22 adds the lifecycle half) +
the `snapshot-spec.md` store-fork honesty update + the structure-capture extension (ext tag `prop-room-m21` @
`835d940`). **Close:** 1 scope (the unwritten committed doc → Fate-1 authored) · 1 should-fix code (an error-label
nit) · 1 docs · 0 adversarial-new (5 scenarios all already test-pinned). Deferral audit **GREEN** (0 repeat / 0 aged;
`directus_files` ref capture → **Fate-3 annotated to M23**, the 20 dangling relations → **Fate-2 already owned by
M23**; the harden conn-seam + serve-live-integration → tracked follow-ups). Go `stack-snapshot` 231→**290** (+59);
coverage directus/firewall **100%**, manifest 98.4%, capture 98.9%; flake **0** (5/5 shuffled). Records:
[m21-structure-capture/](releases/archive/01.50-prop-room/m21-structure-capture/) (decisions · hardening-ledger · metrics · retro · audit-deferrals).
**Goal:** Make the snapshot carry the content-model **structure** (the user-collection table DDL + Directus's
`directus_collections`/`directus_fields`/`directus_relations` registry rows) alongside the rows, captured atomically
from the same sanctioned source — so the `directus` replay stops failing with exit 4 and a freshly-bootstrapped
Directus can actually serve the captured catalog.
**Exit gate (observable, machine-verifiable):** on a scratch offset-port Postgres + a bootstrapped `directus/directus:11.6.1`:
`stacksnap` applies the captured structure → `stacksnap replay --surface directus` **exits 0** → a booted Directus
**serves a captured public simulation over HTTP to an anonymous reader** (`GET /items/simulations?limit=1` → 200 with
a real row). Today this whole chain dead-ends at the print-only `provision.go:108` placeholder.
**Why iterative (not section):** the implementation path is genuinely uncertain — Directus's anonymous-read
permissions + the registry-row carve-out + the cache-digest convergence are undesigned territory that only breaks
live (the analogous fix16 cost +479 lines of empirical correction). A fixed `In:` checklist would be speculative;
the gate is the commitment, the path emerges from each tik's evidence.
**Scope (the known shape; refined per-iter):**
  - In (**`rosetta-extensions`**): the **capture-side structure extension** — capture user-collection DDL (via a
    `pg_dump --schema-only`-equivalent over the `directus` schema from the sanctioned `--dsn`; zero DDL-capture code
    exists today) + the three registry tables filtered to the **public content model** (a collection-name allow-list);
    a **structural-metadata admissibility class** in the firewall (registry rows are `directus_*` system tables —
    today excluded by the letter of `AssertPublicOnly`; they carry no tenant/private columns, so they need an
    explicit "structure, not tenant data" classification, not a blanket pass); the **cache-keying fix** so a
    structure-less stack can converge to the source digest (apply structure **before** the row replay — the row cache
    is keyed by the *target* schema digest, so a bootstrap-only stack can never cache-hit by construction today);
    **redefined `stacksnap` exit-4 semantics** (today "schema missing → a capture can't help"; now the structure
    artifact IS what provisions the schema); and **wire the `directus_files` ref capture** (the docs claim it ships,
    `snapshot-spec.md:414`, but `media.go`'s filter/columns are dead code — no `directus_files` TableSpec; needed for
    the real-image asset refs).
  - Out: executing the recipe at bring-up (M22); the env re-point + referential closure (M23); **S3 blob bytes**
    (stays backlog, DEF-M10-01 — refs + prod-link assets are the floor).
**Depends on:** none (first milestone). **Parallel with:** none (the foundation the rest builds on).
**Estimated complexity:** large (the empirical long pole; expect an M9-style capture-source design loop).
**Open questions:** DDL source — `pg_dump -s` shell-out vs `information_schema` catalog reconstruction (lean:
`pg_dump --schema-only -n directus`, already available on the sanctioned restore-a-dump `--dsn` path, simplest
correct); manifest carries structure as an **additive field** (the `Predicate` precedent, no format bump) vs a
sibling artifact keyed by source digest (lean: decide in iter-01 against the digest-convergence constraint); the
prod-Directus-version vs local-11.6.1 **skew** policy (lean: record the source version in the manifest, pin the
local image, warn on mismatch).
**KB dependencies:** `corpus/ops/snapshot-spec.md` (the store-fork + capture-source + manifest sections),
`corpus/ops/snapshot-cold-start.md` (the `--dsn` source the structure capture rides), `corpus/ops/safety.md`
(the firewall admissibility classes).
**Delivers → `corpus/ops/directus-local.md`** (rosetta — **net-new**: the per-stack Directus spec — the
empirically-pinned bootstrap facts, the structure-capture model, the version-skew rule) **+ the structure-capture
extension + `directus_files` capture + the redefined exit codes** (rosetta-extensions).
**Risk (prod-safety — blocks-prod-safety):** structure capture is still a **prod read** — it must stay behind the
M9a capture-source policy (read-only, bounded, operator-confirmed, public-only) and the `AssertPublicOnly` firewall,
now extended (not loosened) to admit structural metadata. The dropped pg_dump-file-reader (M9b-D9) must **not**
resurrect — `TestDroppedDumpFlagStaysGone` pins it gone.

### M22: Executed provisioning + per-stack Directus lifecycle
**Status:** `done` (closed 2026-06-13) · **Shape:** `section`
**Closure:** all 6 sections landed Fate-1 — the print-only recipe is now an **executed**, idempotent, prod-safe
bring-up step: a per-stack Directus boots as a **compose service** in the stack's override (offset port
`8055+N·10000`, joins `app-network`, named `<project>-directus-1`, torn down with the stack), provisioned
**bootstrap → apply-structure → replay → boot**, verified (`/server/health` SERVICES row + a `directus` expected-schema
+ a **"registered collections > 0"** cheap-win + a **no-prod-read env assert**), demo **default-on** / dev **opt-in**
(`--local-content`; `N=0` behind `--force`). The load-bearing finding: the **`EnvContract.Validate` firewall became
a load-bearing executed gate** — a prod-resolving env **hard-aborts before any write** (the M17-for-TRUNCATE
discipline, now for the executed provision), with a runtime **no-prod-read backstop** in `autoverify`. Non-fatal
throughout: any provision failure **degrades to the prod-read path** with an honest `⚠` status line.
**Idempotent re-provision** converges (bootstrap-on-non-empty **sentinel** guard — a half-bootstrap re-bootstraps —
+ compose-name reuse). **Delivered:** `corpus/ops/directus-local.md` § "Container lifecycle (M22)" (the lifecycle
half) + the `verification.md`/`idempotency.md` rows + the `rosetta_demo.md` registry/teardown note + 3 collateral
print-only retirements (`snapshot-spec.md`/`safety.md`/`demo/README.md`); ext tag `prop-room-m22`. **Close:** 8
findings — 0 scope · 0 must-fix code (3 should-fix comments landed ext `e989982`; 3 nice-to-have refactors → Fate-2
M24) · 2 docs (DOC-1 the ops-README stale "lands in M22/M23" claim → Fate-1 fixed; DOC-2 the CLAUDE.md doc-list →
Fate-2 M24's sweep) · 0 tests · 0 adversarial-new (**7 scenarios all already test-pinned**). Deferral audit **GREEN**
(0 repeat / 0 aged / **0 M22-originated**; inherited M21 items confirmed owned by M23). Go **795** unchanged (M22
touches no Go); Python **360 → 418** collected (+58, +8 env-gated skip); flake **0** (5/5 sequential). Records:
[m22-provision-lifecycle/](releases/archive/01.50-prop-room/m22-provision-lifecycle/) (decisions · metrics · retro · audit-deferrals).
**Goal:** Turn the print-only 4-step recipe into an **executed** bring-up step that boots a per-stack Directus as a
**compose service** in the stack's override — idempotent, verified, torn down with the stack — so demo (default) and
opt-in dev stacks come up with a live local Directus.
**Scope:**
  - In (**`rosetta-extensions`**): **execute** bootstrap → apply-structure → replay → boot inside the shared
    `dev-setdress.sh` engine (replacing the print-only block), **demo default-on / dev opt-in** (`--local-content`;
    `dev-N≥1` direct, `N=0` additionally behind the existing `--force` n=0 guard); **emit the Directus container into
    the per-stack override as a compose service** (offset port `8055+N·10000`, joins the stack's app-network, gets
    the `<project>-directus-1` name) — **not** a bespoke `docker run` (so `demo-down`/`dev-down`, the port registry,
    and `stack-verify`'s naming convention all cover it with no new lifecycle code, per the maintainability
    constraint); **idempotent re-provision** (bootstrap-on-non-empty + container-name-conflict guards, matching the
    M17 re-run contract); **register the Directus offset port**; **Directus verify probes** in `stack-verify` (a
    SERVICES row + `/server/health`, `directus` added to the expected-schemas list, a **"registered collections > 0"
    cheap-win** — the silent-failure analog of the casbin assert — and a **no-prod-read env assert**); the **12 GB-VM
    preflight** accounting extended to include the Directus container; **non-fatal** (a failed boot degrades to the
    prod-read path with an honest status line — never blocks a good stack).
  - In (**`rosetta`**): update `corpus/ops/verification.md` + `corpus/ops/idempotency.md` (new rows) + `corpus/ops/rosetta_demo.md` (registry/teardown).
  - Out: the env re-point (M23 — M22 boots + verifies the instance; M23 points services at it); referential closure (M23).
**Depends on:** M21 (a replay that exits 0 is the prerequisite for a Directus that serves content). **Parallel with:** none.
**Estimated complexity:** medium.
**Open questions:** compose-service vs sidecar override file (lean: a service block in the existing injected/dev
override — reuses the proven generator, nothing bespoke); how loud a degraded "boot failed → still reading prod"
status should be (lean: a clear ⚠ line in the set-dress status, consistent with the M18 verify block).
**KB dependencies:** `corpus/ops/snapshot-spec.md` (the store-fork recipe), `corpus/ops/verification.md`,
`corpus/ops/idempotency.md`, `corpus/ops/rosetta_demo.md`, the `dev-setdress`/`stack-verify` sections.
**Delivers → `corpus/ops/directus-local.md`** (rosetta — the lifecycle half: container/compose/port/teardown +
idempotent re-provision + verify probes) **+ verification.md/idempotency.md rows** (rosetta) **+ the executed
provisioning + compose-service emission + verify probes** (rosetta-extensions).
**Risk (data-safety — blocks-prod-safety):** an executed provision must only ever write the per-stack-isolated offset
Directus/Postgres — the `EnvContract.Validate` firewall moves from a print-time check to a **load-bearing executed
gate** (hard-abort before any write if the env resolves to prod); tests pin the target class, as M17 did for TRUNCATE.

### M23: Content cutover + referential closure
**Status:** `done` (closed 2026-06-13) · **Shape:** `section`
**Closure:** all 6 sections landed Fate-1 — the data plane is now **cut over** to the per-stack Directus: §1 grew
`stack-core/gen_override.py` to emit per-service `environment:` blocks (the single genuinely-new bit of plumbing) +
re-point `cms`'s `DIRECTUS_BASE_ADDR` → in-network `http://directus:8055` on dev `--with-directus`; §2 the demo
side via `gen_injected_override.py`; §3 **studio-desk** gets the per-stack instance + a **locally-minted static
admin token** (deterministic `EnvContract.AdminToken` stamped via Directus's `ADMIN_TOKEN` bootstrap env, a
new `ValidateProvisionable` present-token gate layered on the prod-safety `Validate`); §4 **wired the
`directus_files` ref capture** (DEF-M21-03/Fate-3 from M21) as a new **REFERENCED-SUBSET** firewall admissibility
kind (reverse-reference closure, admit-iff `Filter==ReferencedFilesFilter`) + a `ClearByDelete` **DELETE-before-
TRUNCATE** for the external `directus_settings` FK; §5 **referential closure** — full-taxonomy capture
(`organization_id IS NULL`) was already the state (closure maximal by construction) + a **measured cross-surface
closure gene** (`OpSnapshotCrossSurfaceClosure`/`CrossSurfaceDangling`, criticality `standard`, non-blocking) that
surfaces the **1 genuine prod residual** — `K-AIFUNX-E658`, a node referenced by 2 public sims but existing only as
a customer-scoped skill, **uncloseable by tooling** (capturing it = firewall breach; editing prod = forbidden), now
MEASURED + named rather than a silent empty picker (an operator-owned prod data-quality fix). The **asset plane
stays on prod** (`DIRECTUS_PUBLIC_BASE_ADDR` unchanged) so images stay real; the data plane goes local. **2
inherited M21 deferrals RESOLVED in-milestone** (DEF-M21-03 directus_files §4, DEF-M21-04 referential closure §5 —
the 20 dangling relations were subsumed by M21's 26-collection structure capture, M21-D7). **Delivered:** the 6
corpus docs (`cms`/`studio-desk`/`jobsimulation`/`next-web-app` env truth + `safety.md` retire-live-read +
`snapshot-spec.md` the closure gene + `directus_files`-wired) — resolves KB-1 + KB-2; ext tag `prop-room-m23` @
`7e9343a`. **Close:** 6 findings — 0 scope · 0 code-quality · 0 adversarial-new (4 scenarios all already
test-pinned) · 1 docs (the `snapshot-spec.md` M13-section stale "M23 cutover remains future" → Fate-1 fixed) · 0
tests · 5 decision-triage backref-tags. Deferral audit **GREEN** (2 inherited RESOLVED, 0 repeat / 0 aged;
K-AIFUNX-E658 fated operator-owned KNOWN-ISSUE). Go **795 → 844** (M23-own +33 across `stack-snapshot`+`stack-seeding`;
the rest a counting-method reconciliation on untouched modules); python touched suites `stack-core` 61→**69** +
`stack-injection` **110** (8 env-gated skip); flake **0** (5/5). Records:
[m23-content-cutover/](releases/archive/01.50-prop-room/m23-content-cutover/) (decisions · metrics · retro · audit-deferrals).
**Goal:** Point the stack's services at their **own** Directus and guarantee the served catalog is
**referentially closed** — so the content a stack serves never references a taxonomy node-id its captured subset
lacks (the empty Assign-AI-Simulation-picker class disappears).
**Scope:**
  - In (**`rosetta-extensions`**): **re-point `DIRECTUS_BASE_ADDR`** to the local instance — demo via the ready
    `gen_injected_override.py` env mechanism (one-line per service); **dev via growing `stack-core/gen_override.py`
    to emit `environment:` blocks** (it can only emit ports today — the single genuinely-new bit of plumbing, kept
    minimal); **keep the asset plane on prod** (`DIRECTUS_PUBLIC_BASE_ADDR` stays `content.anthropos.work` so
    browser images stay real — the user's explicit call — sidestepping the baked next/image host whitelist with no
    UI rebuild); **studio-desk** gets the local instance (`DIRECTUS_BASE_URL`) + a **locally-minted admin token** so
    its skill-path writes target the per-stack Directus (never prod); **extend the prod-token strip to opted-in dev
    stacks** (today demo-only); **referential closure** — make the taxonomy capture include every node-id the
    captured content references (closure-at-capture; **full-taxonomy capture** as the simple fallback the corpus
    already names) + a **cross-surface fidelity gene** so closure is *measured*, not assumed.
  - In (**`rosetta`**): update `corpus/services/{cms,studio-desk,jobsimulation,next-web-app}.md` (the env/dependency
    truth) + `corpus/ops/safety.md` (retire the live-prod-read notes; the token-strip stays as the write-disarm).
  - Out: M21/M22 mechanics; the cms `PostMultipart` hardcoded-prod-upload-URL **platform bug** (can't fix without a
    platform edit — disarmed by the token strip + documented as a user-owned upstream PR + a verify warning).
**Depends on:** M22 (a booted, verified local Directus to point at). **Parallel with:** none.
**Estimated complexity:** medium (data-design risk, not code volume — give it the decisions budget M9b had).
**Open questions:** closure-at-capture (compute the referenced node-id set, capture exactly those) vs
**full-taxonomy capture** (no subset — removes the dangling problem wholesale, simpler, slightly bigger snapshot)
(lean: full-taxonomy capture for simplicity per the maintainability constraint, unless the size is prohibitive);
whether N=0 dev re-point gets any tooling or stays a pure documented manual recipe (lean: documented recipe only —
honor the n=0 guard + zero-platform-edit line).
**KB dependencies:** `corpus/ops/snapshot-spec.md` (the referential-closure boundary + the fidelity genes),
`corpus/ops/safety.md`, `corpus/services/*` (the Directus consumers), `corpus/ops/seeding-spec.md` (the content↔seed linkage).
**Delivers → updated `corpus/ops/safety.md` (the self-contained-content deltas) + `corpus/services/*` (env truth) +
`snapshot-spec.md` (the cross-surface closure gene)** (rosetta) **+ the env re-point (demo + dev) + closure capture
+ the closure gene** (rosetta-extensions).
**Risk (correctness — degrades-quality):** a wrong re-point silently sends a stack back to prod (or to a dead local
instance). Mitigate: the M22 no-prod-read verify assert + the `EnvContract` gate catch both; non-fatal degrade keeps
a good stack up.

### M24: Docs convergence + hygiene strand
**Status:** `done` (2026-06-13) · **Shape:** `section` · **Complexity:** small
**Closed:** all 7 sections Fate-1. The corpus now tells the new truth: the fictional local-Directus docker service
(image `directus/directus:10.10.1` + admin/password + a compose snippet) — which **never existed** in the platform
compose (verified against `stack-dev/platform/docker-compose.yml`) — is corrected across `external_services.md` /
`service_taxonomy.md` / `quick_ops.md`, the known-state + `directus-local.md` + `snapshot-spec.md` rewritten on the
M23 cutover, and the print-only/exit-4/reads-live-from-prod framing swept across 5 skills + `CLAUDE.md` + 5 demo
docs (the real two-path posture: prod-read default; per-stack local Directus on `--local-content`). The 4 aged-out
hygiene items the deferral audit surfaced all landed Fate-1: **(a)** an explicit `toolchain go1.25.11` pin on all 4
go.mod + the clerkenstein CI (clears the 12 called-stdlib advisories, lazy); **(b)** a corpus README-index-row guard
(`stack-core/corpus_index_guard.py`, a token-bounded per-dir lint — it dog-fooded itself, surfacing + backfilling 7
pre-existing gaps; harden surfaced the milestone's 1 real bug, a prefix-collision false-negative, fixed `191d650`);
**(c)** the alignment zero-critical-genes guard (`dna.Validate` rejects + `compare.GateMet` refuses a vacuous-100%
critical gate, defence-in-depth, #M24-D2); **(d)** the `/project-stats` scope fix landed cross-repo in the
developer-kit `stats.sh` (where the script lives — `825cdce`): `*/stack-*/*` in PRUNE_PATHS + a `drop_gitignored`
filter, so the gitignored `stack-*/` platform clones (~2M foreign lines) stop inflating the count. **Close:** 5
findings (0 scope · 2 code-quality [0 must-fix] · 0 docs · 0 tests · 1 adversarial-record · 3 decision-triage),
all Fate-1; deferral audit **GREEN** (0 new deferrals, the 4 chartered hygiene items cleared, 3 standing inherited
items unchanged + not aged). Go 844→**850** (+6, all alignment — the zero-critical guard); python stack-core
77→**85** (+8, the README-index guard); flake **0** (5/5). ext tag `prop-room-m24` @ `6a4749d` (set by the
orchestrator post-close); the §7 `/project-stats` fix is the cross-repo developer-kit `825cdce`, outside the ext tag.
Records: [m24-docs-hygiene/](releases/archive/01.50-prop-room/m24-docs-hygiene/) (decisions · metrics · retro · audit-deferrals).
**Goal:** Make the whole corpus tell the new truth (stacks are content-self-contained), and absorb the four small
aged-out hygiene items the deferral audit surfaced — so v1.5 leaves the repo honest and the backlog cleaner.
**Scope:**
  - In (**`rosetta`** docs): rewrite the `snapshot-spec.md` known-state block (the per-stack Directus is now real;
    exit-4 semantics redefined), the `safety.md` §2 deltas, finish `corpus/ops/directus-local.md`; **correct the
    stale local-Directus claims** in `corpus/architecture/external_services.md` (image 10.10.1 + admin/password +
    a compose snippet — all **false**; the platform compose has no directus service — verified),
    `service_taxonomy.md:242-260`, `quick_ops.md:162`; sweep the "print-only / exit-4 / reads live from prod"
    language across the skills + `CLAUDE.md` (via `/update-knowledge`).
  - In (**`rosetta-extensions`** — the hygiene strand, each small + independently landable):
    **(a)** bump the Go toolchain pin to **go1.25.11+** (clears the 12 called-stdlib advisories; **lazy rebuild** —
    no dedicated rebuild session, per the user); **(b)** a **corpus README index-row guard** (a lint that fails when
    a new doc lacks its directory-README index row — the recurring miss in 3 straight releases; v1.5 ships new docs,
    its exact protected class); **(c)** the **alignment zero-critical-genes guard** (`dna.Validate`/`compare.pct`
    treat a zero-critical DNA as 100% — verified still absent; reject/flag it); **(d)** the **`/project-stats`
    scope fix** (stop scanning the gitignored `stack-*/` platform clones that inflate the absolutes).
  - Out: anything in M21–M23; the DROPPED items (AI content, shareability, more mirrors, deploy-CI gate, dev-up pre-warm).
**Depends on:** M23 (the docs describe the finished cutover). **Parallel with:** none (but the four hygiene items are
internally independent — land in any order).
**Estimated complexity:** small.
**Open questions:** none material (the four hygiene items are self-contained).
**KB dependencies:** the full `corpus/ops/` + `corpus/architecture/` + `corpus/services/` set touched by the cutover;
the `alignment` + stats tooling for the hygiene items.
**Delivers → the corpus-wide truth-up + the 3 stale-doc corrections** (rosetta) **+ the README-index-row guard +
the zero-critical-genes guard + the stats-scope fix + the Go pin bump** (rosetta-extensions).
**Risk (low):** the stale-doc corrections must be verified against the *actual* platform compose, not assumed
(already verified once: no directus service exists).

### M25: Field bake — the observable-behavior gate
**Status:** `done` (completed 2026-06-13) · **Shape:** `section`
**Goal:** Prove the whole release on the **actual 16 GB box** with **observable behaviors** as the done-bars — so
v1.5 pre-pays the field-fix tail that every prior release shipped after the fact (v1.3 → all of v1.3b; v1.3b →
fix1–17).
**Scope:**
  - In (live runs on the user's machine, fixes folded back into the owning module):
    - fresh **`/demo-up`** → the browser shows content **served by the local Directus** (data plane local) with
      **real images** (asset plane prod) + the verify net GREEN incl. the new Directus probes;
    - **`/dev-up 2 --local-content`** → same, on an opt-in dev stack; confirm **N=0 stays on the prod-read path**
      untouched (the documented manual opt-in recipe exercised once);
    - **re-run everything twice** (idempotency live — re-provision + replay + seed);
    - the **cold-start capture** path exercised once (structure + rows captured together from a restored dump);
    - clean **teardown** (`/demo-down`, `/dev-down 2`) reclaims the Directus container + its port; the registry is honest.
  - Out: new features (this is a hardening gate — bugs surfaced get fixed in their owning module, no new scope).
**Depends on:** M24 (the full release in place). **Parallel with:** none (the closing gate).
**Estimated complexity:** medium (history's loudest lesson: doc-green ≠ field-working — budget the tail *inside* the release).
**Open questions:** none — the done-bars are the observable behaviors above.
**KB dependencies:** every v1.5 doc; the `/demo-up`, `/dev-up`, `/demo-down`, `/dev-down`, `/test-platform` skills.
**Delivers → a short `releases/archive/01.50-prop-room/m25-field-bake/` field-bake log + any folded-back fixes** (both repos as needed).
**Risk (resource — degrades-quality):** a Directus container per stack adds to the Docker-VM budget. Mitigate:
runtime is cheap (measured ~0.9 GB/stack, ~0.66 GiB both frontends — boots/builds spike, not steady-state), keep the
**max-2-co-resident-stacks** line, add Directus to the 12 GB preflight, watch Docker-VM **disk** (the M3 disk-full precedent).
**Closure (2026-06-13):** all **5 live done-bars GREEN** — the field-bake earned its keep. The operator-sanctioned
prod read (the `postgres` MCP `marco_read` primary-read DSN) filled the cache **structure-bearing**, so the local
Directus observably **SERVES** the captured public catalog (curl-proven: DB-1 demo-1 offset 18055 + DB-2 dev-2
offset 28055 — `/server/health` 200, `/items/simulations` real published rows, asset URL → `content.anthropos.work`;
data plane local + asset plane prod). N=0 stays prod-read (code-verified, M25-D4); re-runs converge (DB-3);
the cold-start capture was exercised (DB-4); teardown reclaims the container + port (DB-5). The live runs surfaced
+ fixed **4 real release bugs** clean-room unit passes couldn't see, all Fate-1 in their owning ext module
(`stack-snapshot`/`demo-stack`): **(1)** the `directus_files` referenced-subset capture **over-captured 158
tenant-referenced files** (M25-D5) — the **tenant-data firewall caught it FAIL-CLOSED** (zero data written); the
closure now appends `AND NOT (TenantReferencedFilesFilter)` + the firewall `AssertPlan` now **requires** every
referenced-subset table to declare + embed its tenant-exclusion (**firewall never weakened**); **(2)** the
`directus_collections.group` self-FK to an admin-UI folder collection (M25-D6) → the serve-row render NULLs it;
**(3)** the `directus_files.folder`/`uploaded_by`/`modified_by` FKs to uncaptured admin/library/users tables
(M25-D7) → a general `TableSpec.NullColumns` mechanism (the complete enumeration proven — content tables ship
PRIMARY KEYS only, 0 content-table FKs); **(4)** the offline-build `GOTOOLCHAIN=local` regression from M24's Go
pin (M25-D1) that aborted `/demo-up`. ext harden +8 tests (firewall `AssertPlan` → 100%); ext HEAD `1a2fd91`,
tag `prop-room-m25`. **Close:** 3 findings (1 scope checkbox nit · 0 code-quality · 1 docs the M25-D2 16 GB-host
VM ceiling field note · 0 tests · 1 decision-blend + 5 archive), **deferral audit GREEN** with the strongest
outcome — **DEF-M21-02** (serve-live-integration harness) RESOLVED Fate-1 in its destination milestone (the live
serve-proof IS the integration it needed); 2 M25-originated env items fated fresh (M25-D8 full-UI Playwright render
→ DROP-as-deliverable/host-budget; M25-D9 dev-2 taxonomy `rc=4` → tracked dev migrate-ordering follow-up);
**0 escape-hatch**. The rosetta merge is markdown/text only (the code fixes are in ext). **v1.5 is complete — next:
`/developer-kit:close-release`.**

### Execution graph (v1.5)
```
v1.5 "prop room" — a real local Directus serving the captured public content, per stack
  M21 (structure capture, iterative) ──→ M22 (executed provisioning + lifecycle)
   └─→ M23 (content cutover + referential closure) ──→ M24 (docs + hygiene) ──→ M25 (field bake)
```
**Strictly sequential** — every release in project history merged milestones serially, and this chain is a genuine
dependency spine: M21 makes the replay succeed, M22 boots the instance, M23 points services at it + closes the
referential gap, M24 tells the truth, M25 proves it live. Parallelize **within** a milestone via agents (recon
spikes, harden passes), not across the chain. **No B-milestones** — the whole release *is* the tooling layer (the
v1.3/M14 + v1.3b precedent); M24 carries the doc convergence inline.

### Parallelism matrix (v1.5)
| Pair | Can parallelize | Shared surface | Merge risk | Strategy |
|------|-----------------|----------------|-----------|----------|
| M21 ∥ M22 | no | M22 executes what M21 makes replayable | — | M21 first (its exit gate is M22's precondition) |
| M22 ∥ M23 | no | M23 points services at M22's booted instance | — | sequential |
| M23 ∥ M24 | no | M24 documents M23's finished cutover | — | sequential |
| M24 ∥ M25 | no | M25 proves the whole release | — | M25 closes |

### Risks (v1.5)
- **(M21, prod-safety — blocks-prod-safety)** structure capture is a prod read — must stay read-only/bounded/
  public-only behind `AssertPublicOnly` (now extended to admit structural metadata, not loosened) + the capture-source
  policy + operator confirm. Mitigate: reuse the M9a capture path; the dropped file-reader stays dropped.
- **(M21, empirical — degrades-quality)** Directus anonymous-read permissions + the registry-row carve-out only break
  live (fix16's whole existence). Mitigate: the iterative shape with live tiks; the gate is "serves anonymously", not doc state.
- **(M22, data-safety — blocks-prod-safety)** an executed provision must only write the per-stack-isolated offset
  target — `EnvContract.Validate` becomes a load-bearing executed gate. Mitigate: hard-abort before any write; tests pin the target class.
- **(M23, correctness — degrades-quality)** a wrong re-point sends a stack to prod or a dead instance. Mitigate: the
  M22 no-prod-read verify assert + the env gate; non-fatal degrade.
- **(M25, resource — degrades-quality)** a Directus container per stack grows the VM/disk budget. Mitigate: max-2
  co-resident, the 12 GB preflight, watch disk.
- **(cross-cut — the two user constraints)** **never touch prod data/platform repos** (tooling + docs only; platform
  repo = build context at most; the cms upload-to-prod hardcode is a documented platform bug, not ours to fix) **+
  keep it simple/maintainable** (compose-service over bespoke `docker run`; reuse the one shared engine + the existing
  override generators + the generic replay path; one genuinely-new bit of plumbing — dev env-emission — kept minimal).

### Open decisions (resolve during build)
DDL source — `pg_dump -s` vs catalog reconstruction — M21 (lean `pg_dump --schema-only -n directus`);
structure as additive manifest field vs sibling artifact — M21 (lean decide in iter-01 against digest convergence);
Directus version-skew policy — M21 (lean record source version + pin local image + warn); compose-service vs sidecar —
M22 (lean a service in the existing override); referential closure — closure-at-capture vs full-taxonomy capture —
M23 (lean full-taxonomy for simplicity); N=0 dev re-point — tooling vs documented recipe — M23 (lean documented recipe).

## Done — v1.3b "dress rehearsal" (SHIPPED 2026-06-09 · tag `v1.3.1`)

**Theme:** v1.3 "stack party" converged the dev/demo **model**; running `/demo-up` for real revealed it didn't yet
converge the **experience**. A demo stack comes up **backend-only** (no UI), **unseeded** (no org/users → every
authorized route 403s), and **unverified** — `up-injected.sh` prints "UP. Clerk-free demo-N is live" with **zero
automated checks**, and in this very session it announced "UP" while the Sentinel authz policy had silently failed
to load (`casbin_rules` = 0). v1.3b is the **dress rehearsal**: the full run-through that makes `/demo-up` produce a
**full, populated, verified, demoable** stack — and makes the tooling honest about what it does. It addresses the 14
issues logged in [`.agentspace/demo-up-issue.md`](../../.agentspace/demo-up-issue.md).

> **Designed 2026-06-08** from the user's `/demo-up` field log (14 issues, several already analyzed/fixed by the
> prior session). **Phase 0a deferral audit GREEN** — the only open deferral (DEF-M10-01, S3 blob bytes + cloud
> store) is orthogonal and stays → **backlog (unscheduled)**; v1.3b adds zero new deferrals at design. **Phase 0b KB blind-areas:**
> the **frontend tier** and **post-bring-up verification** have no corpus anchor (→ M19/M18 each `Deliver →` one);
> idempotency/cold-start/auto-set-dress are anchored but need a contract section (→ M17/M20). Research +
> verification (3 agents, every issue claim re-checked against live code in `.agentspace/rosetta-extensions/`):
> [`.agentspace/scratch/roadmap-research-2026-06-08.md`](../../.agentspace/scratch/roadmap-research-2026-06-08.md).
>
> **Scope decided (user, 2026-06-08):** the **5-milestone spine** (M16→M20, one per issue cluster); the heavy
> features (frontend UI in M19, auto snapshot+seed in M20) are **default-on + skippable** (full parity with
> `dev-up`: one command = a real demo; `--no-ui` / `--no-setdress` to opt out). Codename **"dress rehearsal"** —
> a demo *is* a show; this is the run-through that makes it actually perform.

> **The two-repo split (the load-bearing distinction for this release).** Every milestone names which repo owns
> what, because v1.3b touches **both**:
> - **`rosetta-extensions`** (the executable tooling) holds all **scripts / Go / Python / per-section README+GUIDE /
>   its own `knowledge/` KB**. Built + tested in the **authoring copy** `.agentspace/rosetta-extensions/`, **tagged**
>   (`dress-rehearsal-mNN`), then **consumed per-stack** at the pinned tag — never authored ad-hoc inside a `stack-*/`.
> - **`rosetta`** (this corpus) holds the **`.claude/skills/*`, `corpus/*`, `CLAUDE.md`, READMEs** — the docs and the
>   skills that *drive* the CLIs.
>
> So a typical milestone = a code change in `rosetta-extensions` (authoring → tag → consume) **+** the skill/corpus
> doc in `rosetta` that describes it. The `Delivers →` lines below split accordingly.

### M16: Land the field fixes + restore doc truth
**Status:** `done` (completed 2026-06-08) · **Shape:** `section` · **Complexity:** small–medium · **Dir:** [m16-land-fixes/](releases/archive/01.3b-dress-rehearsal/m16-land-fixes/)
**Closed 2026-06-08** (build → 2 harden passes → close review → merged to `release/01.3b-dress-rehearsal`). **The first milestone of v1.3b "dress rehearsal"** — it lands the honest baseline (publishes the two stranded field fixes + restores doc truth) that every other milestone builds on. A **docs/publish/rename milestone** (the M14 shape): the two functional changes (the ISSUE-1 devpath rename-resolution + the ISSUE-7 migrate-race `|| echo 0` under `set -e`) were pre-applied in a prior session; M16 made them **durable + public + fenced**. Delivered, in the SEPARATE nested `.agentspace/rosetta-extensions` repo (gitignored from rosetta, pushed to `origin`): the **publish** (extensions `main` `a31d70b..e6161b0` to `origin`; the local-only `stack-party-devpath-fix` tag superseded — not deleted, never pushed); the **stack-core rename migration** making `stack-dev` the documented default across all 5 workspace-resolvers (`up-injected.sh`, `migrate-demo.sh`, `rosetta-demo` ×2, `dev-stack`) + `gen_override.py`/`clone_repos.py`, demoting `anthropos-dev` to the **single intentional legacy-alias fallback** (`[ -d ] || …anthropos-dev`, each line legacy-marked — M16-D2); the **prose sweep** (`demo-stack/README.md`/`GUIDE.md`, `dev-stack/README.md`, `gen_override.py` docstring → `stack-dev/platform`); the **GUIDE header truth-up** (remote-exists / 21 tests / `/stack-list` / v1.3 — was no-remote / 78 / `/demo-status` / v1.1·M3); the **pytest doc fix** (`pytest tests/ -v` + the 3.11/3.12 note); the **extensions `knowledge/` KB** (the consumption version-jump + the per-milestone tag scheme). In **rosetta**: a consolidated **`corpus/ops/rosetta_demo.md` stack-dev layout + back-compat note** (code snippet + CLAUDE.md cross-link + don't-reintroduce-bare-`anthropos-dev` guidance); `corpus/` had **0** stray `anthropos-dev` (sweep = verify no-op). **Harden** (2 passes) made the contract a *test*: 8 new demo-stack guards (`TestRenameDrift` 3 — `stack-dev` must lead `anthropos-dev` in every resolver + no unmarked `anthropos-dev`; `TestGuideDocTruth` 2 — the advertised test count + the documented `pytest` entrypoint pinned to live; `TestMigrateRaceGuard` 3 — the ISSUE-7 fence, negative-tested). Close found **1 finding** (the docs review surfaced `clerkenstein/knowledge/glossary.md` still naming the old `anthropos-demo/` workspace — the **last** non-fallback stale workspace name in the entire repo; landed Fate-1 as M16-D8, reversing build-time M16-D5's boundary deferral; ext `e6161b0`); the close **reconciled** the `dress-rehearsal-m16` tag from the build HEAD `44edc09` to the final HEAD `e6161b0` (the tag trailed the 2 harden commits + the 1 close commit — a sanctioned forced tag re-point, force-pushed; cf. v1.3 M9a) and **re-consumed** `stack-demo/rosetta-extensions` to it (three-way agreement: origin tag = authoring HEAD = consumed HEAD = `e6161b0`; the only per-stack clone — `stack-dev/` has none, M16-D6). 2 adversarial scenarios recorded (both-roots-exist resolves deterministically to `stack-dev`; the race is fenced). Deferral re-audit **GREEN** (1 in-release single — the live docker-harness migrate-race *behavior* test → **Fate-2 → M17**, already owned by M17's `In:` scope, no overview edit, M16-D7; 1 inherited DEF-M10-01 → backlog (unscheduled) signed/not-aged; 0 repeat/chronic/aged-out — M16 is the release's first milestone, nothing inherited within v1.3b). Decisions M16-D1…D8. Extensions: tag **`dress-rehearsal-m16`** @ `e6161b0` (per-milestone `dress-rehearsal-mNN` naming, M16-D1; a final `v1.3.1` lands at close-release). Python test funcs 174→**182** (demo-stack 13→21, +8 guards); Go **713** unchanged (M16 touched no Go); flake **0** (5/5); shellcheck + py_compile CLEAN on all touched scripts. Retro: [m16-land-fixes/retro.md](releases/archive/01.3b-dress-rehearsal/m16-land-fixes/retro.md). **Next:** M17 (bring-up re-run safety — idempotency + first-run race).
**Goal:** Make the two already-applied fixes durable and public, finish the `anthropos-dev → stack-dev` rename as
the *documented default*, and clear the stale tooling docs — so the repo tells the truth before more work lands on it.
**Scope:**
  - In:
    - **`rosetta-extensions`:** **push** the local fixes (`547de17` devpath + `ed72e94` migrate-race — currently 2
      commits ahead of `origin`, on a local-only `stack-party-devpath-fix` tag) to `origin` under a proper
      `dress-rehearsal-m16` tag, then re-consume per-stack (ISSUE-1①/ISSUE-7 push); the **stack-core rename
      migration** making `stack-dev` the documented default + demoting `anthropos-dev` to a single intentional
      "legacy alias" mention (ISSUE-1②); the **prose sweep** (`demo-stack/README.md:12`, `GUIDE.md:17`,
      `dev-stack/README.md:73`, `stack-core/gen_override.py:4` docstring — ISSUE-2); **GUIDE.md header truth** (remote
      exists; **13** tests not 78; `/stack-list` not `/demo-status`; **v1.3** not v1.1/M3 — ISSUE-3); the **pytest
      doc fix** (`pytest tests/` + a 3.11/3.12 note — ISSUE-4); refresh the repo's own `knowledge/` KB where it
      repeats any of these.
    - **`rosetta`:** sweep any residual `anthropos-dev` in `corpus/`; note the expected consumption version-jump (ISSUE-5).
  - Out: the race/idempotency *behavior* work (M17); anything frontend/verify/set-dress.
**Depends on:** none. **Parallel with:** none (the honest baseline every other milestone builds on).
**Estimated complexity:** small–medium
**Open questions:** the re-tag/version scheme (lean: per-milestone `dress-rehearsal-mNN` tags + a final `v1.3.1`
release tag at close, matching the established convention); whether to keep `anthropos-dev` back-compat in the
scripts forever or sunset it (lean: keep — it's a one-line fallback that costs nothing).
**KB dependencies:** `corpus/ops/rosetta_demo.md`, the `rosetta-extensions` GUIDE/README set + its `knowledge/` KB.
**Delivers → updates the `rosetta-extensions/knowledge/` KB + GUIDE/README truth-up** (rosetta-extensions) **+ a
consolidated `corpus/ops/` note on the `stack-dev` layout + the back-compat fallback** (rosetta).

### M17: Bring-up re-run safety — idempotency + first-run race
**Status:** `done` (completed 2026-06-09) · **Shape:** `section` · **Complexity:** medium · **Dir:** [m17-rerun-safety/](releases/archive/01.3b-dress-rehearsal/m17-rerun-safety/)
**Closed 2026-06-09** (build → 2 harden passes → close review → merged to `release/01.3b-dress-rehearsal`). **The 2nd milestone of v1.3b** — it makes the bring-up PRIMITIVES re-run-safe so M20's auto-chaining is safe to retry. A **code milestone** (re-run guards + the first-run-race sweep) + one net-new corpus doc; the testable logic lives in the SEPARATE nested `.agentspace/rosetta-extensions` repo, the rosetta branch carries only `corpus/ops/idempotency.md` (net-new) + 5 wired-parent cross-links + tracking. Delivered: (1) the **`set -e` first-run-race audit** across the 4 bring-up scripts — `up-injected.sh`'s `GH_PAT` now **fails loud** (not a silent pipefail abort) + `rosetta-demo`/`dev-stack` `DEV_PROJECT` carries `|| true` so the documented `${DEV_PROJECT:-anthropos}` fallback runs (M17-D1), + the **wait-for-sentinel-ready** defense-in-depth in `migrate-demo.sh` (bounded non-fatal `wait_pg` via `pg_isready`/`SELECT 1` + `wait_sentinel_running` via `docker inspect`, M17-D2), + a **4th latent ISSUE-7 site** the live harness surfaced — the schema-create `docker exec psql` under `set -e` got a `|| log` guard (M17-D9); (2) the **`stacksnap replay` re-run guard** — `Replayer.ClearForReplay`, a **per-stack-isolated `TRUNCATE`-then-reload**, child-first (reverse dependency order, no `CASCADE`) before the parent-first COPY, so a 2nd replay REPLACES not appends; safe-by-default, no flag (M17-D3); the clear SQL is the pure `truncateForReplaySQL`, **target-class pinned** to ALWAYS be a single-table `TRUNCATE … RESTART IDENTITY` (never DROP/DELETE/CASCADE/cross-schema, injection-quoted), and the connection is the per-stack offset DSN — two independent fences against a wrong-target TRUNCATE (M17-D4); (3) the **`stackseed` re-run guards** — `Conn.CopyRowsIdempotent` (COPY-to-session-TEMP-then-`INSERT…SELECT…ON CONFLICT (id) DO NOTHING` in one tx, preserving the bulk path; all 7 seeders rewired, M17-D5), the casbin g2 grant via `INSERT…SELECT…WHERE NOT EXISTS` (the casbin tables have no unique tuple constraint, M17-D6), and the **fixed `--reset`** extended to the full FK-ordered fleet (`activity_events → jobsim sessions → skill_path_sessions → assignments → memberships → users → organizations`) + a targeted casbin g2 `DELETE` (not TRUNCATE — preserves `init_policy.sql`'s global policy, M17-D7); (4) the **live docker-harness migrate-race test** — `test_migrate_race_live.py` runs the real `migrate-demo.sh` against a throwaway pgvector container IN the race state, proving survival + idempotency (the M16 Fate-2 item LANDED, M17-D8; skips cleanly without Docker); (5) the net-new **`corpus/ops/idempotency.md`** — the per-component re-run verdict table + the two re-run models (re-run-in-place vs teardown-then-redo) + the engineer detail per guard, wired bidirectionally into `demo/README`, root `CLAUDE.md`, `snapshot-spec`, `seeding-spec`, `safety.md`. **Prod-safety verified at close:** the new TRUNCATE/idempotent-write logic is provably confined to per-stack-isolated offset targets (the `truncateForReplaySQL` target-class pin + the structural offset DSN; the `--reset` surface pinned to `schema.table`-only entries; the casbin policy preserved via the targeted g2 DELETE), and the M15 read-side + write-side safety drift guards stay GREEN after the `safety.md` cross-link addition — no way the new logic could TRUNCATE a non-per-stack target. **Harden** (2 passes) deepened mutation-pins (the temp-merge SQL builders, the injection sweeps, the casbin 7-col dedup, the `--reset` destructive-surface invariants) — 0 bugs surfaced; the build-phase logic held under every probe. Close found **1 finding** (the decision-triage review surfaced `idempotency.md` blended the D1–D7 mechanism rationale but lacked the `(#M17-DK)` back-ref tags the v1.3 corpus precedent uses — added Fate-1; no code/test/scope finding) and reconciled the `dress-rehearsal-m17` tag from the build HEAD `dcef026` to the final HEAD `0d36251` (the tag trailed the 2 harden commits; a sanctioned forced tag re-point, force-pushed) + re-consumed `stack-demo/rosetta-extensions` to it (three-way agreement: origin tag = authoring HEAD = consumed HEAD = `0d36251`). Adversarial scenarios captured as live mutation-pinned tests (the degenerate-identifier TRUNCATE sweep, the seed-side injection sweep, the casbin 7-col dedup correctness). Deferral re-audit **GREEN** (M17 added **0** new deferrals — every decision D1–D9 is a Fate-1 landing; the M16 Fate-2 live migrate-race test LANDED here as M17-D8, reaching its destination; 1 inherited DEF-M10-01 → backlog (unscheduled) signed/unchanged/not-aged — M17 touched the re-run-guard surfaces, NOT the snapshot-store/S3 area; 0 repeat/chronic/aged-out). Decisions M17-D1…D9. Extensions: tag **`dress-rehearsal-m17`** @ `0d36251` (reconciled from `dcef026`). Go test funcs 713→**736** (+23: stack-seeding 236→252, stack-snapshot 224→231); Python 182→**191** collected (+9 demo-stack: the set-e race fences + the live docker harness); **flake 0** (5/5 all 3 touched suites); all 4 Go modules `-race -count=1` + gofmt + `go vet` clean; the 4 touched shell scripts shellcheck-CLEAN; py_compile CLEAN. Retro: [m17-rerun-safety/retro.md](releases/archive/01.3b-dress-rehearsal/m17-rerun-safety/retro.md). **Next:** M18 (the verification safety net).
**Goal:** A re-run of migrate / snapshot-replay / seed is either safe-and-idempotent or fails loudly with a guard —
never silently doubles data or aborts mid-surface.
**Scope:**
  - In (all **`rosetta-extensions`** code; one **`rosetta`** doc):
    - the **`set -e` first-run-race audit** across the bring-up scripts (the same class as the fixed ISSUE-7
      migrate race — sweep `up-injected.sh`/`rosetta-demo`/`dev-stack`/`dev-setdress.sh`) + an optional
      **"wait-for-sentinel-ready"** defense-in-depth (ISSUE-7 residual);
    - **`stacksnap replay`** re-run protection — a per-stack-isolated `TRUNCATE`/skip/`ON CONFLICT` before the bare
      `COPY` so a second replay doesn't duplicate-key-abort mid-surface or silently double (`stack-snapshot/pg/pg.go`,
      `replay/replay.go` — ISSUE-11);
    - **`stackseed`** re-run protection — `ON CONFLICT` for the casbin g2 grant + the deterministic-UUID rows; **fix
      the stale `--reset` truncate list** to include the session/activity tables it currently skips
      (`stack-seeding/seeders/identity.go`, `cmd/stackseed/main.go` — ISSUE-11);
    - an explicit, **tested** idempotency contract for each component.
  - Out: the verify net (M18); the auto-chaining of snapshot/seed (M20 — M17 only makes the *primitives* re-run-safe).
**Depends on:** M16 (clean pushed baseline). **Parallel with:** **M18 (yes-with-caveats** — different surfaces;
both touch `up-injected.sh` only in different regions).
**Estimated complexity:** medium
**Open questions:** `TRUNCATE`-and-reload vs `ON CONFLICT DO NOTHING` for replay (lean: TRUNCATE the per-stack
target surface then reload — simplest correct semantics, and the target is always per-stack-isolated); whether to
make re-run safety automatic or behind an explicit `--idempotent`/`--force` (lean: safe-by-default with a loud guard).
**KB dependencies:** `corpus/ops/snapshot-spec.md`, `corpus/ops/seeding-spec.md`, the `stack-seeding`/`stack-snapshot` sections.
**Delivers → `corpus/ops/idempotency.md`** (rosetta — net-new: the per-component re-run verdicts + the
teardown-then-redo model + the new guards) **+ the new guards in `stack-seeding`/`stack-snapshot`** (rosetta-extensions).
**Risk (data-safety):** a `TRUNCATE` must *only* ever hit a per-stack-isolated offset target — the n=0 + prod-isolation
guards stay inviolate; tests pin the target class.

### M18: The verification safety net
**Status:** `done` (completed 2026-06-09) · **Shape:** `section` · **Complexity:** medium–large · **Dir:** [m18-verify-net/](releases/archive/01.3b-dress-rehearsal/m18-verify-net/)
**Closed 2026-06-09** (build → 3 harden passes → close review → merged to `release/01.3b-dress-rehearsal`). **The 3rd milestone of v1.3b** — it makes a bring-up self-verifying so "UP" carries a real promise (the ISSUE-7 silent-403 stack would now be caught at bring-up time, in seconds, automatically), and so the later M19 (frontend tier) + M20 (auto-chaining) inherit a *working* stack. A **code milestone** (offset/scope-aware `stack-verify` + the bring-up-tail auto-wire) + one net-new corpus doc; ALL testable logic lives in the SEPARATE nested `.agentspace/rosetta-extensions` repo, the rosetta branch carries only `corpus/ops/verification.md` (net-new) + the index rows + the `rosetta_demo.md` cross-link + tracking. Delivered: (1) **offset/project awareness** — a new `stack-verify/lib/target.sh` resolution helper (`STACK_PROJECT`/`STACK_OFFSET`/`STACK_SERVICES`) sourced by `lib/services.sh` + `lib/readiness.sh`: the SERVICES table stays a single **base-port** source of truth, the offset is applied once centrally (host port `8082` → `38082`, container `anthropos-cms-1` → `demo-3-cms-1`), and the offset is the N the bring-up already KNOWS — passed explicitly, then **cross-checked non-fatally** against the unified registry's RECORDED ports via a base-port BAND `(port−offset) ∈ [3000,11000]` (M18-D1/D5; the band avoids the broken `port//10000==n` decade lane that would false-warn roadrunner's high base); (2) the **service/profile scope filter** — `STACK_SERVICES ∩ the SERVICES array`, honoured in **BOTH** the liveness AND the readiness phase (a reduced bring-up skips an out-of-scope deep probe via the same `target_service_selected` gate — the harden Pass-1 bug fix where the readiness phase previously ran its 6 deep probes unconditionally; M18-D2); (3) the **`$DEVDIR` → `$STACK_ROOT` bugfix** (`repos/run.sh`, `census/inventory.sh` — the undefined var collapsed every repo to `/$repo` → "not cloned" on the first run); (4) the **cheap-win asserts** inside `live/autoverify.sh` — `curl -fsS .../api/health` on the offset port + `SELECT count(*) FROM sentinel.casbin_rules > 0` via `docker exec`, the exact ISSUE-7 silent-failure catcher, each gated by the scope filter (M18-D4); (5) the **default-on + NON-FATAL auto-wire** at both bring-up tails — `demo-stack/up-injected.sh` (opt-out `DEMO_NO_VERIFY=1`) + `dev-stack` `cmd_up` (opt-out `DEV_NO_VERIFY=1`), mirroring `dev-setdress`'s proven pattern (M18-D3); (6) the net-new **`corpus/ops/verification.md`** — the contract + the offset/scope model + the correctness mitigation + how to read a warning block, indexed in the ops README + root `CLAUDE.md` + cross-linked from `rosetta_demo.md`. **The load-bearing non-fatal invariant verified at close** on three axes: `autoverify.sh` is structurally non-fatal (no `set -e`; every probe wrapped; **always `exit 0`**; both call sites `|| true` too — two independent guarantees), the offset is derived-from-known and cross-checked against recorded ports (not a drift-prone formula), and BOTH phases honour the scope filter — so a verify/offset bug can never abort a genuinely-good bring-up and never systematically false-`down`s a healthy offset or reduced-profile stack. **Harden** (3 passes) found + fixed **1 real bug inline** (the readiness-phase scope filter, commit `2f412a3`) + deepened to 79 tests (the offset matrix, the cross-check edges, `probe_service`, the `$STACK_ROOT` resolution, the bring-up wiring). Close found **3 findings**: **FINDING-A1** (adversarial Phase 2c — a non-numeric `--offset`/`STACK_OFFSET` flowed verbatim into three `$(( base + offset ))` sites → "unbound variable" under `set -u`, **silently skipping** the cheap-win asserts; the invariant still held — exit 0 — but the asserts vanished with a confusing message; **fixed** at the single resolution boundary, `target_resolve_offset` now validates `^[0-9]+$` else warns non-fatally + derives, M18-D7), its **regression test** (unit + integration), and a **decision-triage** item (added the `(#M18-D1/D3/D5)` back-ref tags to `verification.md`). Reconciled the `dress-rehearsal-m18` tag from the build HEAD `594b9cf` to the final HEAD `777723a` (the tag trailed the 3 harden commits + the close A1 fix; a sanctioned forced tag re-point, force-pushed) + re-consumed `stack-demo/rosetta-extensions` to it (three-way agreement: origin tag = authoring HEAD = consumed HEAD = `777723a`). Deferral re-audit **GREEN** (M18 added **0** new deferrals — every deliverable + D1–D7 a Fate-1 landing; frontend-port verification is **Fate-2** owned by M19, a clean scope boundary not a deferral; the 1 inherited DEF-M10-01 → backlog (unscheduled) signed/unchanged/not-aged — M18 touched the verify surface, NOT the snapshot-store/S3 area; 0 repeat/chronic/aged-out). Decisions M18-D1…D7. Extensions: tag **`dress-rehearsal-m18`** @ `777723a` (reconciled from `594b9cf`). Go test funcs **736** (unchanged — M18 touched no Go); Python 191→**273** collected (+82: the net-new `stack-verify/tests/test_verify.py`, 32 build → 79 harden → 82 close); **flake 0** (5/5 the touched suite, deterministic); all 9 touched shell scripts shellcheck-CLEAN; py_compile CLEAN; all 4 Go modules still build + pass `-count=1`. Retro: [m18-verify-net/retro.md](releases/archive/01.3b-dress-rehearsal/m18-verify-net/retro.md). **Next:** M19 (the demo-up frontend tier).
**Goal:** Teach `stack-verify` to target an *offset* stack and scope to the services actually brought up, then
auto-run it (non-fatal) at the tail of every bring-up — so "UP" means **verified-working**, not just *containers-started*.
**Scope:**
  - In (**`rosetta-extensions`** code; one **`rosetta`** doc):
    - **project/offset awareness** — derive the `demo-N`/`dev-N` prefix + the N×10000 port offset from
      `STACK_ROOT`/the unified registry (today `lib/services.sh:25-39` hardcodes 12 `anthropos-*-1` names at **base**
      ports, so it reports an offset stack entirely `down` — ISSUE-12b);
    - a **service/profile filter** intersecting the checked set with what was requested (so a reduced bring-up isn't
      a wall of false `down`s);
    - **fix the undefined `$DEVDIR` → `$STACK_ROOT` bug** (`repos/run.sh:108`, `census/inventory.sh:75` — ISSUE-12);
    - the **cheap-win asserts available today** (`curl -fsS .../api/health` + `SELECT count(*) FROM
      sentinel.casbin_rules > 0` on the stack's offset ports at the bring-up tail — the exact ISSUE-7 silent-failure
      catcher — ISSUE-14);
    - the **auto-wired scoped `verify live`** at the bring-up tail, **default-on + non-fatal** (mirroring
      `dev-setdress`'s proven pattern — ISSUE-12c/ISSUE-14).
  - Out: verifying the *frontend* ports (added in M19, where the frontends first exist); deep behavioural/e2e probes
    (that's the `/test-platform` skill's job — M18 is the always-on smoke net).
**Depends on:** M16. **Parallel with:** M17 (yes-with-caveats).
**Estimated complexity:** medium–large
**Open questions:** derive the offset from `STACK_ROOT` parsing vs reading the registry's recorded host ports (lean:
read the registry — it already records resolved ports per M12); how loud "non-fatal but failed" should be (lean: a
clear ⚠️ block + a one-line "run `/test-platform N` to dig in").
**KB dependencies:** `corpus/ops/run_guide.md`, `corpus/ops/rosetta_demo.md`, the `stack-verify` section, the `/test-platform` skill.
**Delivers → `corpus/ops/verification.md`** (rosetta — net-new: the auto-verify contract + the offset/scope model)
**+ the offset-/scope-aware `stack-verify` + bring-up wiring** (rosetta-extensions).
**Risk (correctness):** a mis-derived offset would false-positive "down" — the very bug it fixes. Mitigate: derive
from the registry's recorded ports; non-fatal so a verify bug never blocks a genuinely good stack.

### M19: The demo-up frontend tier
**Status:** `done` (completed 2026-06-09) · **Dir:** [m19-frontend-tier/](releases/archive/01.3b-dress-rehearsal/m19-frontend-tier/)
**Shape:** `section`
**Goal:** `/demo-up` brings up the full UI — next-web + studio-desk at offset ports (per-demo **cached** image from
the **unmodified** platform Dockerfile) + ant-academy natively — so a demo is actually demoable, on a 16 GB Mac.
**Scope:**
  - In:
    - **`rosetta-extensions`:** extend `stack-injection/gen_injected_override.py` to **emit `next-web-app` +
      `studio-desk`** (offset ports via `ports:!override`, `image: demo-N-*` + `mem_limit:1g`, additive override —
      today it emits backend-only, ISSUE-8); `up-injected.sh` **builds the two frontends serially, before compose
      up**, from the unmodified Dockerfiles with **offset-URL build-args + the minted Clerk pk** via a gitignored
      `.env.local`/BuildKit overlay, **tag-guarded for cache reuse**; ship a sibling `.dockerignore` (5.6 GB context
      → <100 MB); **launch (or document) ant-academy natively** (port 3077, its own `.env`,
      `REQUIRE_ORGANIZATION_MEMBERSHIP=0`); a **12 GB Docker-VM pre-flight assert** (ISSUE-6/ISSUE-9); **register the
      frontend ports** so M18's verify net covers them. **Default-on + skippable** (`--no-ui`).
    - **`rosetta`:** update the `demo-up` SKILL.md (the UI is now in scope) + author the frontend-tier corpus doc.
  - Out: the **optional upstream platform PR** for *true* zero-rebrebuild (runtime-rewrites + `__env.js` +
    `output:standalone`) — it edits platform repos → **forbidden / user-owned**, documented as a follow-up (like
    a future deploy-CI item), not built here.
**Depends on:** **M18** (so the verify net can cover the new frontend services). **Parallel with:** none.
**Estimated complexity:** large (the meatiest milestone of v1.3b)
**Open questions:** ant-academy native-launch *by* the tool vs documented-manual (lean: launch it, fall back to a
documented step if the native run proves fiddly); whether to pre-warm the frontend image cache as part of `dev-up`
too (defer — demo-first).
**KB dependencies:** [`.agentspace/demo-up-frontend-plan.md`](../../.agentspace/demo-up-frontend-plan.md) (the
verified tooling-only plan), `corpus/services/next-web-app.md`, `corpus/services/ant-academy.md`,
`corpus/services/studio-desk.md` *(if present)*, `corpus/ops/rosetta_demo.md`.
**Delivers → `corpus/ops/demo/frontend-tier.md`** (rosetta — net-new: ports, per-demo build, Clerk-pk baking, the
12 GB VM prereq, the honest "one ~3-min cached build per new demo-N" residual) **+ updated `demo-up` skill** (rosetta)
**+ the frontend-emitting override generator + per-demo build in `up-injected.sh`** (rosetta-extensions).
**Risk (scope+resource):** the ~3.7 GB / ~3 min per-frontend build swap-thrashes on an undersized VM (the original
"hour"). Mitigate: 12 GB preflight assert, serial cached builds, `mem_limit`, the sibling `.dockerignore`. **Hard
line: zero platform-repo edits — the repo is a build *context* only, the Dockerfile is unmodified (verified achievable).**
**Closure (2026-06-09):** Delivered all 8 deliverables + 4/4 verification. `gen_injected_override.py` appends the UI
tier (`next-web-app` + `studio-desk`) as `profiles:!override [graphql]` services with per-demo built images
(`demo-N-*`, `build:!reset null`, `pull_policy:never`, `mem_limit:1g`, offset ports), `--no-ui` clears it, and the
stale `next-web-app` `REUSE_DEV` entry was removed (#M19-D1/D2). `up-injected.sh` builds the two frontends **serially,
before compose up**, from the **unmodified** Dockerfiles with offset-URL build-args + the minted Clerk pk
(next-web via a gitignored `apps/web/.env.local`; studio-desk as a direct build-arg — #M19-D3), **tag-guarded** for
cache reuse, **non-fatal**, fronted by a **non-fatal 12 GB VM pre-flight** (`DEMO_VM_MIN_GIB` override — #M19-D5), with
a tooling-owned **transient non-clobber `.dockerignore`** (`RETURN`-scoped-trap-removed, so a failed build leaves the
repo byte-clean — #M19-D4). `ant-academy.sh` launches the academy **natively** on `:3077+offset` Clerk-free
(`BENCHMARK_VISUAL_BYPASS`), default-on + non-fatal + degrades to a documented step (#M19-D6/D9), wired into
up/down. `stack-verify`'s registry gained the frontend rows, **scoped iff UI on** (#M19-D7). Net-new
`corpus/ops/demo/frontend-tier.md` + the updated `demo-up` SKILL. **The load-bearing zero-platform-edit invariant
held** — harden pinned it with `TestZeroPlatformRepoEdit` (a real-git-repo `git status`-clean guard on both the
success and the failed/aborted-build path + a `git check-ignore` fence), mutation-verified. Close: 6 findings, **all
decision-triage** (0 scope / 0 code / 0 docs / 0 tests — the cleanest shape); Phase-7 added the 5 `(#M19-D3..D7)`
ref-tags + recorded 3 re-examined adversarial scenarios (all already handled). Tag `dress-rehearsal-m19` reconciled
`32b1ae8 → 4f96ddd` + re-consumed (3-way agreement). Deferral re-audit **GREEN** (0 v1.3b-internal; DEF-M10-01 → backlog (unscheduled)
untouched/not-aged; the true-zero-rebuild upstream PR is a documented **OUT** boundary, not a deferral). Go **736**
(unchanged — no Go touched); Python 273→**338** collected (+65, the UI-tier suites); `gen_injected_override.py` 98%;
flake **0** (5/5). Merged `--no-ff` → `release/01.3b-dress-rehearsal`; `m19/frontend-tier` deleted.

### M20: Lifecycle convergence — demo-up auto set-dress + cold-start capture
**Status:** `done` (completed 2026-06-09) · **Shape:** `section` · **Complexity:** medium–large · **Dir:** [m20-lifecycle-convergence/](releases/archive/01.3b-dress-rehearsal/m20-lifecycle-convergence/)
**Closed 2026-06-09** (build → 1 deepening harden pass + 1 confirmation scan → close review → merged to `release/01.3b-dress-rehearsal`). **The 5th + FINAL milestone of v1.3b** — it converges the **lifecycle**: `/demo-up` now auto-set-dresses by default (cache-first snapshot **replay** → a `small-200` light seed) at the bring-up tail, exactly as `/dev-up` has since M13, by **reusing the very same `dev-setdress.sh` engine via `--stack-type demo`** — ONE engine, two lifecycles, the convergence by construction (not a second implementation). A **tooling change in `rosetta-extensions`** (the set-dress chaining in `up-injected.sh` + the stack-type-aware engine) + **docs in `rosetta`** (the net-new cold-start runbook + safety §2.7 + demo recipe/skill updates); the testable logic lives in the SEPARATE nested `.agentspace/rosetta-extensions` repo, the rosetta branch carries only `corpus/ops/snapshot-cold-start.md` (net-new) + the safety/recipe/skill/cross-link edits + tracking. Delivered: (1) the **set-dress chaining** — `demo-stack/up-injected.sh` chains `dev-stack/dev-setdress.sh --stack-type demo --force` AFTER migrate, BEFORE the M18 verify, **default-on + NON-FATAL** (`DEMO_NO_SETDRESS=1` escape), threading the resolved offset DSN (`5432+OFFSET`) as `DEV_SETDRESS_DSN` (#M20-D1); (2) the **stack-type-aware engine** — `dev-setdress.sh` made `--stack-type dev|demo` (default `dev` for M13 back-compat): picks the prefix (`dev-N|demo-N`) + the default preset, **all safety preserved byte-for-byte** (the n=0 guard type-agnostic, the prod-Directus firewall, never-capture); (3) the **atomicity contract** — the seed ALWAYS runs after the (optional, cache-first) replay (the seed is the FLOOR; a catalog-only stack would 403, a replay miss degrades to a structural-only world that still logs in 200), retry-safe via the M17 re-run guards (#M20-D3); (4) the **demo default preset = `small-200`** (vs dev's `dev-min` — demos want a fuller world; an explicit `--seed` overrides, #M20-D2); (5) the net-new **`corpus/ops/snapshot-cold-start.md`** — the fresh-box capture workflow (the sanctioned DSN-export / restore-a-`pg_dump`-then-`--dsn` path), **why the wired `postgres` MCP is NOT a capture source** (it returns JSON rows, not COPY-format bytes; an adapter would re-serialize COPY text for zero gain — the **ISSUE-13 spike resolved DOCUMENT-ONLY**, #M20-D4), and how it slots into the auto-set-dress bring-up (replay-only, never capture); + **`safety.md` §2.7** (the demo chain reuses the dev pass — the guarantees carry over) + the demo README/recipe + the `demo-up` SKILL + cross-links (db-access/snapshot-spec/CLAUDE.md). **The load-bearing prod-safety invariant verified at close** + test-pinned on BOTH the happy and the cache-miss-degraded branches: the bring-up chain does cache-first **REPLAY only and NEVER runs `stacksnap capture`** (`test_capture_is_never_invoked_on_a_bring_up` + harden's `test_capture_never_runs_even_on_the_cache_miss_degraded_path`, mutation-pinned); the per-stack Directus env is firewall-checked before any replay (`test_provision_recipe_failure_aborts_before_replay`); the offset DSN reaches the engine as `5432+OFFSET`, never base `5432` (`test_chain_success_passes_clean_and_threads_the_offset_dsn`, mutation-pinned); the n=0 guard fires across types (`test_n0_guard_fires_for_demo_type_too`) with `--force` keeping it a dev-only net; the atomicity floor holds (`test_demo_seed_is_the_atomicity_floor_after_a_replay_miss`); and the **M15 `safety.md` drift guards re-ran GREEN** after the §2.7 edit. **Harden** (1 deepening pass + 1 confirmation scan) added +6 (the degraded/success/error branches, mutation-pinned) — **0 bugs surfaced**; the build's logic held under every probe. Close found **5 findings**: 0 scope · 0 code · 0 adversarial-new · **1 docs** (DOC-1: the root `CLAUDE.md` `/demo-up` skill-table row omitted the M20 auto-set-dress the `/dev-up` row advertises — fixed for consistency with the convergence narrative) · 0 tests · **4 decision-triage** (the `(#M20-D1)`/`(#M20-D3)` ref-tags into safety §2.7, `(#M20-D2)` into demo/README, `(#M20-D4)` already tagged in the cold-start doc) + the adversarial subsection (5 scenarios, all already test-pinned) recorded. Reconciled the `dress-rehearsal-m20` tag from the build HEAD `e4d2f9b` to the final HEAD `51a07cb` (the tag trailed the harden commit; a sanctioned forced tag re-point, force-pushed) + re-consumed `stack-demo/rosetta-extensions` to it (three-way agreement: origin tag = authoring HEAD = consumed HEAD = `51a07cb`). **Deferral re-audit GREEN** — as the FINAL milestone, this was the **release-wide M16→M20 pre-close sweep**: M20 added **0** new deferrals (all 4 In-items Fate-1; the ISSUE-13 MCP-adapter spike resolved **document-only**, not a deferral); the one inherited DEF-M10-01 (S3 blob bytes + cloud store → backlog (unscheduled), signed) is **UNTOUCHED across all of v1.3b** (a file-level scan of the M16→M20 extensions history found zero `SnapshotStore`/`store.go`/S3/blob touches), all 4 aging triggers negative, authority intact; 0 repeat/chronic/aged-out. Decisions M20-D1…D4. Extensions: tag **`dress-rehearsal-m20`** @ `51a07cb` (reconciled from `e4d2f9b`). Go test funcs **736** (unchanged — M20 touched no Go); Python 338→**360** collected (+22: dev-stack 38→50 the stack-type/atomicity tests + the demo chain suite); **flake 0** (5/5 both touched suites); shellcheck CLEAN on both scripts; py_compile CLEAN; M15 drift guards GREEN. Retro: [m20-lifecycle-convergence/retro.md](releases/archive/01.3b-dress-rehearsal/m20-lifecycle-convergence/retro.md). **Next:** `/developer-kit:close-release` v1.3b (the final v1.3b milestone is closed).
**Goal:** Close the dev↔demo asymmetry — `/demo-up` auto set-dresses (snapshot → seed, default-on + non-fatal) like
`dev-up` already does — and unblock the *real* catalog on a fresh box that has no safe `--dsn`.
**Scope:**
  - In:
    - **`rosetta-extensions`:** **chain** `stacksnap replay` → `stackseed` into `up-injected.sh` after migrate
      (reuse M13's proven `dev-setdress` pass; **default-on + non-fatal**, `--no-setdress` escape — ISSUE-10);
      enforce the **atomicity contract** (a partial snapshot with no seed = 403s, so it's both-or-neither, and the
      M17 re-run guards make a retry safe); the **cold-start capture** path (ISSUE-13) — a documented, prod-safe
      **DSN-export / restore-a-`pg_dump`-then-`--dsn`** workflow (the sanctioned route), **plus a spike** on whether
      a thin MCP-paging capture adapter (read via the wired `postgres` MCP) is cheap enough to build vs document.
    - **`rosetta`:** update the `demo-up`/`demo-down` skills + the `corpus/ops/demo/` recipe family for auto set-dress.
  - Out: **S3 media blob bytes + the cloud `SnapshotStore` backend** (DEF-M10-01 → **backlog (unscheduled)**, signed — untouched);
    AI-generated content (unscheduled backlog).
**Depends on:** **M18** (the post-set-dress verify) **+ M19** (the full experience) **+ M17** (re-run-safe primitives
make auto-chaining safe to retry). **Parallel with:** none (the closing milestone).
**Estimated complexity:** medium–large
**Open questions:** ISSUE-13 — build the MCP-paging capture adapter vs document the DSN-export step only (lean:
**document the sanctioned DSN-export/restore path now**, spike the MCP adapter only if cheap — the MCP is a query
tool, not a `--dsn` `stacksnap` can `COPY` through); whether auto-set-dress on `demo-up` should default to the
`dev-min`-style light seed or a demo preset (lean: a demo preset like `small-200`, since demos want a fuller world
than dev — confirm at build).
**KB dependencies:** `corpus/ops/snapshot-spec.md`, `corpus/ops/seeding-spec.md`, `corpus/ops/db-access.md`, the
`dev-setdress` mechanism (M13), `corpus/ops/demo/README.md` + the recipe family.
**Delivers → `corpus/ops/snapshot-cold-start.md`** (rosetta — net-new: the fresh-box capture workflow + the MCP
limitation + the sanctioned paths) **+ updated demo recipes/skills for auto set-dress** (rosetta) **+ the set-dress
chaining in `up-injected.sh` + the cold-start capture path** (rosetta-extensions).
**Risk (prod-safety):** auto set-dress + cold-start capture must preserve M13/M15's guarantees — **read-only bounded
capture, the tenant firewall, per-stack isolation, and a confirmation before any prod-touching read**. Mitigate:
reuse M13's proven set-dress pass verbatim; capture stays behind M9a's capture-source policy + `AssertPublicOnly`.

### Execution graph (v1.3b)
```
v1.3b "dress rehearsal" — make /demo-up produce a full, populated, verified, demoable stack
  M16 (land fixes + doc truth)
   └─→ M17 (re-run safety) ──┐
       M18 (verify net) ─────┴─→ M19 (frontend tier) ─→ M20 (auto set-dress + cold-start)
            (M17 ∥ M18 feasible — different surfaces / different up-injected.sh regions; lean sequential per the spine discipline)
```
**Mostly sequential.** M16 lands the honest baseline (pushes the stranded fixes). M17 + M18 harden the *primitives*
(re-run safety) and add the *net* (verification) — the one feasible parallel pair (different surfaces; both touch
`up-injected.sh` only in different regions), but the spine discipline leans sequential. M19 builds the **frontend
tier** (the meatiest; depends on M18 so the verify net covers the new services). M20 **converges the lifecycle**
(auto set-dress + cold-start) on top of the full experience — the closing milestone. **No B-milestones — the whole
release *is* the tooling layer** (the v1.3/M14 precedent).

### Parallelism matrix (v1.3b)
| Pair | Can parallelize | Shared surface | Merge risk | Strategy |
|------|-----------------|----------------|-----------|----------|
| **M17 ∥ M18** | yes-with-caveats | `up-injected.sh` (M17: migrate region · M18: tail) | low | land M17 first (smaller), then M18's tail-append; lean sequential per spine discipline |
| M16 ∥ * | no | the honest baseline | — | M16 goes first |
| M19 ∥ M18 | no | M19 needs M18's verify to cover its new frontend services | — | sequential |
| M20 ∥ M19 | no | M20 = full experience on top of M19 | — | M20 closes |

### Risks (v1.3b)
- **(M19, scope+resource — blocks-quality)** the per-frontend build (~3.7 GB / ~3 min) swap-thrashes on an
  undersized VM. Mitigate: 12 GB Docker-VM preflight assert, serial cached builds, `mem_limit`, the sibling `.dockerignore`.
- **(M18, correctness — degrades-quality)** a mis-derived offset false-positives "down" — the bug it fixes.
  Mitigate: derive from the registry's recorded ports; non-fatal so it never blocks a good stack.
- **(M17, data-safety — blocks-prod-safety)** a `TRUNCATE` on the wrong target = data loss. Mitigate: TRUNCATE only
  per-stack-isolated offset targets; the n=0 + isolation guards hold; tests pin the target class.
- **(M20, prod-safety — blocks-prod-safety)** auto set-dress + cold-start capture must keep M13/M15's read-only
  bounded capture + tenant firewall + per-stack isolation + the confirm-before-prod-read. Mitigate: reuse M13's pass; capture stays behind the M9a policy.
- **(M16, blast-radius — nice-to-resolve)** the **push** is v1.3b's first outward-facing action (the local fixes go
  public). Mitigate: re-tag cleanly (retire the local `stack-party-devpath-fix`), push, re-consume per-stack.
- **(cross-cut)** **zero platform-repo edits** — the verified hard line. Every frontend/build change uses the repo
  as a build *context* only; the *true zero-rebuild* path is an explicitly-OUT, user-owned upstream PR.

### Open decisions (resolve during build)
Re-tag/version scheme — M16 (lean `dress-rehearsal-mNN` + final `v1.3.1`); replay re-run = TRUNCATE-reload vs
ON CONFLICT — M17 (lean TRUNCATE per-stack target); offset from STACK_ROOT-parse vs registry-recorded-ports — M18
(lean registry); ant-academy native-launch by-tool vs documented — M19 (lean launch); ISSUE-13 MCP-adapter vs
documented-DSN-export — M20 (lean document now, spike adapter if cheap); demo auto-set-dress preset — M20 (lean a
demo preset over `dev-min`).

## Done — v1.3 "stack party" (SHIPPED 2026-06-07 · tag `v1.3`)

**Theme:** v1.0 made the platform run Clerk-free; v1.1 built the demo/dev stack framework; v1.2 set-dressed *demo*
stacks with the real public catalog. v1.3 **throws the party** — it makes **dev and demo stacks first-class peers**
that work the same way. A dev stack gets the demo treatment (the **per-stack-Directus recipe + firewall check**
[print-only — see the Correction below; no working local Directus is stood up], **auto-snapshot** of the
real reference data on build, a **light default seed** so it's never empty); a **unified stack registry** allocates
the first-available N across both kinds so they never collide on ports; and one **generic `stack-*` skill set**
operates any stack. The safety story (read-side private-data avoidance + write-side prod-protection) is consolidated
into a dedicated doc, and **both** the rosetta corpus and the `rosetta-extensions` knowledge base are brought to the
converged model.

> **Designed 2026-06-07** from 6 user requirements (the dev/demo convergence). Phase 0a deferral audit **GREEN**
> (v1.2 close — 1 open item, DEF-M10-01 S3 blobs + the cloud-store escape-hatch). **Scope decided (user, 2026-06-07):**
> the former v1.3 seeds (cloud store, S3 blobs, AI content, shareability, more mirrors) **move to the backlog (unscheduled)**; v1.3 is the
> tight convergence release. M13 stays one milestone; the operation-skill renames are a **hard rename, no aliases**.
> Phase 0b KB blind-area: the **safety & security doc** is the one blind area → M15 `Delivers →` it. Research:
> [`.agentspace/scratch/roadmap-research-2026-06-07.md`](../../.agentspace/scratch/roadmap-research-2026-06-07.md).

### M12: Unified stack registry + first-available-N allocation
**Status:** `done` (completed 2026-06-07) · **Shape:** `section` · **Complexity:** medium · **Dir:** [m12-stack-registry/](releases/archive/01.30-stack-party/m12-stack-registry/)
**Closed 2026-06-07** (build → 3 harden passes → close review → merged to `release/01.30-stack-party`). **The first milestone of v1.3** and the only non-Go surface to date — the unified registry is **Python** (`stack-core/stack_registry.py`) + shell-CLI wiring. Delivered: a **single shared module** `stack_registry.py` (M12-D1 — one shared N-pool across both kinds, so `dev-1`/`demo-1` can never coexist; runtime file `stack-core/.stacks/registry.json`, gitignored) implementing **first-available-N allocation** — `allocate(type, n=None)` auto-allocates the lowest free N across dev+demo or validates an explicit N, the whole **read-reconcile-pick-write under one `fcntl.flock(LOCK_EX)`** on a sidecar `.lock` (portable macOS+Linux, M12-D3) with **atomic temp+`os.replace` writes** + corrupt-JSON recovery; **registry-of-record + `docker ps` reconcile that only ADDS used-N, never subtracts** (M12-D2 — adopts unmanaged live stacks; `release()` is the only free path = the race guard). Both up-paths wired (`dev-stack` + `rosetta-demo`): `up [N]` (explicit or `auto`/omit), an **ERR-trap that frees the slot on a failed bring-up** (installed *before* the post-allocation re-guard — the §2 review fix), `set-ports` records resolved host ports, `down` releases; `status` leads with the unified dev+demo view. KB-1 resolved (the GUIDE.md "registry assigns N" claim converged to truth). The guarantee — `dev, demo, dev, demo, demo` → `dev-1, demo-2, dev-3, demo-4, demo-5`, any interleaving, either CLI — verified at the **OS-process level** (12-process auto-allocation + a close-added 6-process explicit-N collision, not just threads). Extensions: tag `stack-party-m12` @ `be1c979` (v1.3 release-and-milestone tag naming). **89 Python test funcs** across the 3 M12-touched suites (stack-core 54 / dev-stack 22 / demo-stack 13; +52 net new), `stack_registry.py` 98% (2 structurally-unreachable misses); flake **0** (5/5); shellcheck + py_compile clean; Go total 708 unchanged (M12 touched no Go). Close: scope/code/docs/tests **GREEN**, **3 findings** (1 test added — the concurrent explicit-N collision; 2 decision-triage → archive); 4 adversarial scenarios recorded; deferral re-audit **GREEN** (1 inherited DEF-M10-01 → backlog (unscheduled) signed, 0 repeat/aged/new — M12 added zero deferrals). Decisions M12-D1…D3 + KB-1 + Q1/Q2/Q3 resolved. Retro: [m12-stack-registry/retro.md](releases/archive/01.30-stack-party/m12-stack-registry/retro.md). **Next:** M13 (dev peers — the per-stack-Directus recipe + auto-snapshot + light seed).
**Goal:** One shared registry across **dev + demo** that tracks live stacks and allocates the **lowest free N**, so
bring-ups never collide on ports (build `dev, demo, dev, demo, demo` → `dev-1, demo-2, dev-3, demo-4, demo-5`).
**Scope:** In: a **unified stack registry** (extend the demo `registry.json` to span both kinds — type + N + ports +
status) in `stack-core`; a **first-available-N allocator** (reconcile the registry against live `docker ps`, return
the lowest free N) — race-safe (locked write); the up-paths **accept explicit-N or auto-allocate**; teardown frees
the slot. Out: the skill renames that consume it (M14); dev per-stack-Directus-recipe/snapshot/seed (M13).
**Depends on:** v1.1's `stack-core` (the port-offset engine) + the demo-stack `registry.json`. **Parallel with:** M13 (feasible — different surfaces; lean sequential so M13's dev bring-up consumes the registry).
**Open questions:** registry as a lockfile vs `docker ps`-derived (lean: registry-of-record + `docker ps` reconcile); how a manually-started stack (outside the skills) is reconciled.
**KB dependencies:** `corpus/ops/rosetta_demo.md` (the demo registry/ports), the `stack-core`/`demo-stack`/`dev-stack` extension sections.
**Delivers → updates `corpus/ops/rosetta_demo.md`** (the unified registry + first-available-N model).

### M13: Dev stacks as full-fidelity peers — the per-stack-Directus recipe + auto-snapshot + light seed
**Status:** `done` (completed 2026-06-07) · **Shape:** `section` · **Complexity:** large · **Dir:** [m13-dev-peers/](releases/archive/01.30-stack-party/m13-dev-peers/)
> **⚠ Correction (2026-06-11).** M13 delivered the per-stack-Directus **recipe + the prod-Directus firewall check
> (both made executable, but PRINT-ONLY — no recipe step is run by the bring-up)** + the **taxonomy** replay — **not**
> a working local Directus serving content. No stack type stands a per-stack Directus up; the `directus` content
> replay skips on every stack (`stacksnap` exit 4 — the M10 collection-schema gap), and dev/demo read public content
> **live from prod**. The "Goal/Scope/Delivered" wording below describes the original intent; read "spawn/own a
> per-stack Directus" as "emit + firewall-check the recipe." See the top-of-doc Correction +
> [`../../corpus/ops/snapshot-spec.md`](../../corpus/ops/snapshot-spec.md) § the per-stack Directus store fork.
**Closed 2026-06-07** (build → 1 harden pass → close review → merged to `release/01.30-stack-party`). **The meatiest milestone of v1.3** — it converges dev with demo for **DATA** (M12 converged them for N-allocation). Delivered: a `dev-stack up` bring-up now gets the **demo treatment by default** via a new **set-dressing pass** (`dev-stack/dev-setdress.sh`, default-on + non-fatal): (1) a **per-stack Directus** — the M10 store-fork recipe + the prod-Directus firewall, both made **executable** by a NEW small Go runner `stack-snapshot/cmd/provision-plan` that turns the library-only M10 `directus.ProvisionPlan`/`EnvContract`/`Validate` contract into a CLI (M13-D2 — one source of truth for the recipe AND the firewall, shared by both kinds; `--check-env` hard-aborts the pass before any replay if the per-stack Directus env ever resolves to prod); (2) a **cache-first `stacksnap replay`** of the public `taxonomy`+`directus` surfaces into `dev-N` (replay never captures; a stale/missing cache is a warning, the seed still runs — M13-D3 default-on-but-non-fatal, resolving the bring-up-heaviness Q1 + the default-on Q2); (3) the new **`dev-min` seed preset** (`stack-seeding/presets/dev-min.seed.yaml` — ~1 org + ~10 users + 1mo, fixed admin `dev@anthropos.test` → login 200; the role-mix floor, M13-D1 resolving Q3) so the stack is never empty. The **n=0-dev guard is doubled** — the set-dress pass refuses N=0 without `--force`, a second layer above `stackseed --reset`'s own refusal. **Escapes:** `--no-snapshot` (seed only) / `--no-setdress` (bare bring-up). **Prod-safety held unchanged:** capture is never run on dev (replay only — a per-stack WRITE to the isolated offset Postgres), media stays **refs-only** (blob bytes = unscheduled backlog (DEF-M10-01)), `AssertPublicOnly` + the isolation guard both hold. Corpus: `snapshot-spec.md` (the "Dev as a full-fidelity peer" section + dev as a replay target) + `seeding-spec.md` (the shipped-presets table + the dev-min/dev-auto-seed subsection + the two-layer n=0 guard); the stale "cloud/S3 store = v1.3" claim fixed off v1.3 (the work is now backlog/unscheduled). Harden fixed **3 robustness bugs** inline (whitespace-env firewall fail-closed · the `--no-snapshot` re-run-hint leak · a `set -u` trailing-flag error), each mutation-pinned, +10 regression/edge tests. Close found **6 findings** (1 doc must-fix: the stack-snapshot README Packages table was missing the new `cmd/provision-plan` row — a per-unit handbook-index miss, fixed; 5 GREEN/no-action incl. 4 adversarial scenarios + 3 decision-triage blends ref-tagged); deferral re-audit **GREEN** (1 inherited DEF-M10-01 → backlog (unscheduled) signed/unchanged, 0 repeat/aged/new — M13 added **zero** deferrals, all Fate-1). Decisions M13-D1…D4 + Q1/Q2/Q3 resolved. Extensions: tag **`stack-party-m13`** @ `cca4464`. Go test funcs 708→**720** (stack-snapshot 212→223 the provision-plan runner; stack-seeding 232→233 the dev-min pin); dev-stack pytest 33→**38**; flake **0** (5/5 both languages); both Go modules `-race` + gofmt + `go vet` clean; both CLIs shellcheck-clean. Retro: [m13-dev-peers/retro.md](releases/archive/01.30-stack-party/m13-dev-peers/retro.md). **Next:** M14 (unified `stack-*` skills + `dev-up`/`dev-down`, hard-renamed).
**Goal:** A freshly-built dev stack gets the demo treatment — its **own local per-stack Directus** (no longer
pointing at shared prod), an **auto-snapshot** on build that fills the real public reference data (taxonomy + the
now-local content), and a **light default seed** so it's never empty.
**Scope:** In: wire the dev bring-up to **spawn a per-stack Directus** (reuse M10's `stack-snapshot/directus/provision.go`
per-stack-Directus mechanism) + repoint the dev CMS at it; **auto-run `stacksnap replay`** (taxonomy + directus) as
part of dev build (cache-first, fast); a **`dev-min` seed preset** (smaller than demo's `small-200` — ~1 org + ~10
users + minimal activity) applied on build; keep the n=0-dev-reset guard. Out: the generic skills (M14); blob bytes
(unscheduled backlog — refs-only, as for demo).
**Depends on:** **M12** (consumes the registry for dev-N) + v1.2's M10 (the per-stack Directus + the content snapshot) + v1.1's M6 (dev-stack) + M7 (seeding). **Parallel with:** none (gates M14).
**Open questions:** does spawning Directus + snapshot + seed make dev bring-up too heavy? (mitigate: cache-first snapshot, minimal seed, reuse M10's provision); should the auto-snapshot be default-on or opt-in for dev (lean: default-on, `--no-snapshot` to skip).
**KB dependencies:** `corpus/ops/snapshot-spec.md` (capture/replay + the per-stack Directus store), `corpus/ops/seeding-spec.md` (presets + the dev/n=0 guard), `corpus/services/cms.md` (Directus), the `dev-stack` extension section.
**Delivers → updates `corpus/ops/seeding-spec.md`** (the `dev-min` preset + dev auto-seed) **+ `corpus/ops/snapshot-spec.md`** (dev as a replay target + the per-stack-Directus recipe on dev).

### M14: Unified `stack-*` skills + `dev-up`/`dev-down`
**Status:** `done` (completed 2026-06-07) · **Shape:** `section` · **Complexity:** large · **Dir:** [m14-unified-skills/](releases/archive/01.30-stack-party/m14-unified-skills/)
**Closed 2026-06-07** (build → 1 harden pass → close review → merged to `release/01.30-stack-party`). **The tooling-convergence milestone of v1.3** — it unifies the operation skills onto both stack kinds via a **hard rename, no aliases** (user 2026-06-07; the clean-break blast radius contained by sweeping every in-repo reference). Delivered: **`dev-up`** (consolidates the former `setup-platform` + `start-platform` into one dev bring-up — element-by-element diff against both former bodies confirmed **no dropped step**: STEP RUN discipline, confirmation policy, ops reports, the 12-container set, error recovery; `dev-up N` for N≥1 spins additional isolated dev-N; drives the M13 set-dress flow, M14-D2) + net-new **`dev-down`**; the 4 ops skills hard-renamed each taking a `dev-N|demo-N` target — **`stack-list`** (←`demo-status`), **`stack-seed`** (←`demo-seed`), **`stack-snapshot`** (←`demo-snapshot`), **`stack-update`** (←`update-platform`); the **6 old skill dirs removed** (no shims); a **full reference sweep** (CLAUDE.md skill table rewritten to the 14-skill set, root+corpus READMEs, all `corpus/ops/` guides + the `demo/` recipe family, CHANGELOG — new Unreleased v1.3 Added/Changed/Removed, v1.1/v1.2 dated entries left immutable, M14-D6); `demo-up`/`demo-down` retained + aligned with the dev lifecycle. The `--preset NAME` UX is kept as a **skill-level shorthand** mapped to `--seed presets/NAME.seed.yaml` (M14-D5, the PR-review fix — the `stackseed` binary only knows `--seed <path>`). Extensions: a doc-only reference sweep on `main` (`b37e831`: dev-stack header/README, demo-stack GUIDE, stack-verify run.sh) + the harden reference-integrity guard (`33fc525`: re-pointed the stacksnap drift-guard comments `/demo-snapshot`→`/stack-snapshot` + added **`TestDocSourceSkillRename_M14`** — asserts the renamed skill dir exists + the retired one stays gone, turning silent comment-rot/dir-resurrection into a test failure; negative-tested). The skill namespace `stack-snapshot` is deliberately distinct from the extensions **section** `stack-snapshot/` (the skill drives the `stacksnap` CLI that section builds — called out in the body). Tag **`stack-party-m14`** @ `33fc525` (v1.3 release-and-milestone naming, matching M12 `stack-party-m12` @ `be1c979` + M13 `stack-party-m13` @ `cca4464`). Close: scope/code/docs/tests **GREEN** with **0 findings** — a clean straight-through close; the close re-verified the 4 reference-integrity dimensions **independently** (rename invariant **0 live stragglers** project-wide, CLAUDE.md ⇄ 14 skill dirs perfect bijection, SKILL.md `name:`⇄dir 0 mismatches, skill→CLI contract resolves) + all `../`-relative doc links resolve + 0 broken links to deleted dirs. Deferral re-audit **GREEN** (1 inherited DEF-M10-01 → backlog (unscheduled) signed/unchanged/not-aged, 0 repeat/aged/new — M14 added **zero** deferrals, all Fate-1). Decisions M14-D1…D6 (D2/D3/D4 resolve Q1/Q2/Q3; D5 = the `--preset` fix; D6 = CHANGELOG convention). Go test funcs 720→**721** (stack-snapshot 223→224: the rename guard; alignment/clerkenstein/stack-seeding unchanged); all 4 Go modules `-race -count=1` + gofmt + `go vet` clean; 174 Python green; all 3 CLIs shellcheck-clean; py_compile clean; flake **0** (5/5). Retro: [m14-unified-skills/retro.md](releases/archive/01.30-stack-party/m14-unified-skills/retro.md). **Next:** M15 (safety & security doc + dual-repo KB refresh — the LAST v1.3 milestone).
**Goal:** One coherent stack-operations skill set that works on any stack (`dev-N | demo-N`), with the dev lifecycle
mirroring demo's.
**Scope:** In: **`dev-up`** (consolidate `setup-platform` + `start-platform` into one dev bring-up that drives the M13
flow) + **`dev-down`**; **hard-rename** (no aliases, user 2026-06-07) the operation skills to generic, stack-target
forms — **`stack-list`** (←`demo-status`), **`stack-seed`** (←`demo-seed`), **`stack-snapshot`** (←`demo-snapshot`),
**`stack-update`** (←`update-platform`) — each detecting/accepting a `dev-N|demo-N` target; **remove** the old skill
dirs (`setup-platform`, `start-platform`, `update-platform`, `demo-status`, `demo-seed`, `demo-snapshot`); update
**every reference** (CLAUDE.md skill table, READMEs, the corpus skill docs + recipes). `demo-up`/`demo-down` stay as
the demo lifecycle (aligned with `dev-up`/`dev-down`). Out: the safety doc (M15).
**Depends on:** **M12 + M13** (the generic skills drive the registry + the dev peer capabilities). **Parallel with:** none.
**Open questions:** how much of `setup-platform`'s first-time machine setup (tool install, org clone) folds into `dev-up` vs stays a one-time prerequisite; the exact target-detection UX (`stack-seed dev-1` vs `stack-seed --stack dev-1`).
**KB dependencies:** the existing skill SKILL.md files, `corpus/ops/*` guides (setup/run/update/demo), `CLAUDE.md` (the skill table).
**Delivers → the unified `stack-*` + `dev-up`/`dev-down` skills + a rewritten `CLAUDE.md` skill table + refreshed `corpus/ops/` guides** (the converged stack model, hard-renamed).

### M15: Safety & security doc + dual-repo knowledge consolidation
**Status:** `done` (completed 2026-06-07) · **Shape:** `section` · **Complexity:** medium · **Dir:** [m15-safety-doc/](releases/archive/01.30-stack-party/m15-safety-doc/)
**Closed 2026-06-07** (build → 4 harden passes → close review → merged to `release/01.30-stack-party`). **The closing milestone of v1.3 — the LAST of the four; with it done the release is ready for `/developer-kit:close-release`.** A documentation/consolidation milestone (the M8/M11-analog). Delivered the net-new **`corpus/ops/safety.md`** (248 lines, code-cited, dual-level): the two inviolable guarantees of the `rosetta-extensions` tooling — **never reads private/customer data** (read-side: the tenant firewall `AssertPlan`/`AssertCaptured` [the conceptual `AssertPublicOnly` named to its two real Go gates, M15-D2], the per-surface public predicates byte-matched to `firewall.go` [`organization_id IS NULL` / `private = false AND tenant_id IS NULL AND status = 'published'`], the public-only data-DNA gene, bounded read-only capture [`SET TRANSACTION READ ONLY` + timeouts, **no offline file reader** — M15-D3 keeps the M9b-D9 drop true]) and **never touches prod data or services** (write-side: the 3-layer isolation guard `CheckWrite`/`PreflightEnv`/`AssertClean`, never-write shared Directus/prod-S3, the capture-source policy, the **doubled** n=0-dev guards [scoped precisely — `stacksnap` replay has none, correctly, M15-D4], the audit-proven zero-pollution assertion). Cross-linked from all 4 siblings (`db-access`/`snapshot-spec`/`seeding-spec`/`security_compliance`, both directions). Plus a **dual-repo KB refresh**: `rosetta-extensions/knowledge/` to the v1.3 converged dev≡demo model + safety contract (0 stale pre-M14 skill names); root READMEs + `CLAUDE.md` + the `demo/` recipe family for the unified `stack-*` skills + dev-as-peer + safety.md discoverability. **Harden** (4 passes) made doc accuracy a *test*: 7 fail-closed docs↔code drift guards pin every load-bearing literal/symbol/SQL-block safety.md quotes to the real code (read-side predicates + gate names + the bounded-read SQL `source.DefaultBounds().SetupSQL()`; write-side the complete `realClerkHosts` + `directusTokenKeys` rejection lists + the forced `STORAGE_S3_PUBLIC_BUCKET=""` override + the `CheckWrite`/`PreflightEnv`/`AssertClean` symbols), each `t.Skip`-ping gracefully in pinned-tag consumption clones; harden also landed the **M15-D4** n=0 over-claim fix (corrected in both `dev-setdress.sh` + the sibling test comment to match the shipped safety.md §2.5). Close found **1 finding** (a self-referential docs-accuracy drift: `decisions.md` M15-D4's text lagged the harden fix — corrected; 0 code/test/scope) and proved the guards fail-closed at close (Phase 2c — mutating the read-side predicate + the write-side bucket override each tripped its guard; safety.md restored byte-identical). Deferral re-audit **GREEN**, run release-wide M12→M15 as the terminal-milestone pre-release sweep (1 inherited DEF-M10-01 → backlog (unscheduled) signed/unchanged/**not aged out**; 0 repeat/aged/new — M15 added **zero** deferrals, all Fate-1). Decisions M15-D1…D4 + Q1 resolved. Extensions: tag **`stack-party-m15`** @ `51ca18b`. Go test funcs **+7** (the 7 drift guards: stack-snapshot +3, stack-seeding +4); 174 Python (unchanged); all 4 Go modules `-race -count=1` + gofmt + `go vet` clean; `dev-setdress.sh` shellcheck-clean; py_compile clean; **flake 0** (5/5 both touched packages). Retro: [m15-safety-doc/retro.md](releases/archive/01.30-stack-party/m15-safety-doc/retro.md). **Next:** **`/developer-kit:close-release`** (merge `release/01.30-stack-party` → `main`, tag `v1.3`) — the user's separate step.
**Goal:** A single authoritative doc on **how the rosetta tooling stays safe** — it never reads private/customer data,
and it never touches production data or services — plus a refresh of **both** knowledge bases to the v1.3 model.
**Scope:** In: a new **`corpus/ops/safety.md`** consolidating (a) the **read-side** (private-data avoidance: the
tenant firewall `AssertPublicOnly`, the `organization_id IS NULL` / `private = false AND tenant_id IS NULL` public
predicates, the public-only data-DNA gene) and (b) the **write-side** (prod-protection: the 3-layer isolation guard,
read-only capture, never-write-shared-Directus / prod-S3, the capture-source policy, the n=0-dev guards); cross-link
from `snapshot-spec`/`seeding-spec`/`db-access`/`security_compliance`; **update the `rosetta-extensions/knowledge/`**
(the repo's own KB) for the converged stack model + the safety contract; refresh the root READMEs + the `demo/`
recipe family for the unified `stack-*` skills + dev-as-peer. Out: nothing (closing milestone).
**Depends on:** **M12 + M13 + M14** (documents their converged result). **Parallel with:** none (the closing milestone before `/developer-kit:close-release`).
**Open questions:** doc home — `corpus/ops/safety.md` vs `corpus/architecture/tooling-safety.md` (lean: `corpus/ops/safety.md`, ops-adjacent to seeding/snapshot/db-access).
**KB dependencies:** `corpus/ops/{snapshot-spec,seeding-spec,db-access}.md`, `corpus/architecture/security_compliance.md`, the `rosetta-extensions/knowledge/` base.
**Delivers → `corpus/ops/safety.md`** (net-new — the tooling's read-side + write-side safety contract) **+ updates the `rosetta-extensions/knowledge/` base** to the v1.3 converged model.

### Execution graph (v1.3)
```
v1.3 "stack party" — dev + demo stacks become first-class peers, one unified skill set
   M12 (unified registry + first-free-N) ─→ M13 (dev peers: the per-stack-Directus recipe + auto-snapshot + light seed) ─→ M14 (unified stack-* skills) ─→ M15 (safety doc + dual-repo KB)
```
**Sequential.** M12 lays the shared registry/port foundation; M13 makes dev a full peer (the meatiest — reuses M10's
per-stack Directus + the snapshot/seed framework); M14 converges the skills onto both stack kinds (hard rename); M15
consolidates the safety story + both knowledge bases. M12∥M13 is feasible (stack-core vs dev bring-up) but
serializing keeps M13's bring-up consuming the new registry cleanly. No B-milestones — M14 *is* the tooling layer.

### Risks (v1.3)
- **(M13, scope)** spawning a per-stack Directus + auto-snapshot + seed on **every dev build** could make dev
  bring-up slow/heavy. Mitigate: reuse M10's proven `provision.go`; cache-first snapshot (replay, not capture);
  minimal `dev-min` seed; `--no-snapshot` escape.
- **(M12, correctness)** the first-available-N allocator must be **race-safe** (two concurrent `up`s) and reconcile
  with reality. Mitigate: locked registry write + `docker ps` as the used-N source of truth.
- **(M14, blast radius)** the **hard rename** (no aliases) breaks any external reference to `setup-platform` /
  `demo-seed` / etc. Mitigate: update **every** in-repo reference in M14; the user accepted the clean-break tradeoff.
- **(cross-cut)** dev was historically the *protected* environment (real Clerk, the n=0 guard). Making it a
  snapshot/seed target must preserve those guards. Mitigate: M15's safety doc + the kept n=0-reset guard + read-only capture.

### Open decisions (resolve during build)
Registry-of-record vs `docker ps`-derived — M12 (lean registry + reconcile); dev auto-snapshot default-on vs opt-in
— M13 (lean default-on, `--no-snapshot`); how much first-time setup folds into `dev-up` — M14; the safety doc home
(`corpus/ops/safety.md`) — M15.

## Done — v1.2 "set dressing" (SHIPPED 2026-06-07 · tag `v1.2`)

**Theme:** v1.1 made demo stacks *structurally* populated (orgs, users, backdated activity) but consciously
**waived two surfaces** — `taxonomy` (the 60K-skill / 18K-role node hierarchy + embeddings) and `content` (the
shared Directus content library) — because they can't be *fabricated* structurally: taxonomy is reference data,
content lives in a shared store the isolation guard blocks. v1.2 lifts both via a **snapshot mechanism** —
capture the *real* surface once from a source, replay it per-stack — taking data-DNA coverage from "8 reachable
surfaces" to **100% of the full catalog**. The elevation: snapshot replay is itself an **alignment problem**
(does the replayed surface faithfully reproduce the captured source?), so v1.2 extends the M0/M7b
alignment+data-DNA discipline with a **snapshot-fidelity dimension** rather than shipping plain plumbing. Demo
worlds become *set-dressed*: the stage (v1.1 "show floor") gets its believable props.

> **Designed 2026-06-05** from the v1.1 carry-forward (the user-confirmed M7c waiver → roadmap-vision seed).
> Phase 0a deferral audit **GREEN** (6 carry-forwards, all single, clear v1.2/vision homes — no repeats/chronic;
> [`.agentspace/scratch/deferral-audit-2026-06-05.md`](../../.agentspace/scratch/deferral-audit-2026-06-05.md)).
> Phase 0b KB blind-area check: snapshot/AI-content/deploy-CI **YELLOW** (anchored, need a spec doc as a
> `Delivers →`); shareability **RED** (blind area — deferred to v1.3, not in v1.2). 3 research agents over the
> seeding stack, the skiller taxonomy + Directus content surfaces, and milestone history — verified against the
> clones in `stack-dev/` (platform) + `stack-demo/rosetta-extensions/` (the seeding stack). Gap analysis:
> [`.agentspace/scratch/roadmap-research-2026-06-05.md`](../../.agentspace/scratch/roadmap-research-2026-06-05.md).
>
> **Scope decided (user, 2026-06-05):** **snapshot spine only** — `taxonomy` + `content` to 100% coverage.
> **AI-generated rich content** (transcripts / AI-scored narratives / fresh embeddings) and **external
> shareability** (Tailscale vs ingress) are confirmed **v1.3** (kept in `roadmap-vision.md`), so v1.2 stays the
> tight, well-grounded release the snapshot work warrants. Codename **"set dressing"** continues the stage
> metaphor (body double → show floor → set dressing).
>
> **Refined 2026-06-06** (user — 5 notes on M9+) with **live production read access** (the wired
> `mcp__postgres__query` tool, `marco_read` over Tailscale; catalog-only queries — no GB scans). Changes: (1)
> snapshotting becomes a **dedicated reusable `rosetta-extensions/stack-snapshot/` section** (capture-read
> decoupled from seeding-write); (2) a **production-safe capture-source policy** — source-pluggable, default
> **ingest an existing prod `pg_dump`** → fallback **safe throttled primary read** (MVCC = read-only never blocks
> writers), with **restore-from-snapshot / read replica** as zero-impact upgrades; (3) a tested **tenant-data
> firewall** (`AssertPublicOnly`) — capture *only* `organization_id IS NULL`
> reference data, never customer rows; (4) snapshots live in a **`.agentspace` manifest-cached, pluggable store**
> (no GB blobs in git; cloud/S3 = v1.3); (5) the **`/db-query` skill is ported** into rosetta as the prod-read
> foundation. Former **M9 split → M9a (framework) + M9b (taxonomy surface)** (the M7a→M7c precedent). Prod
> findings (skiller ≈ 2.1 GB, taxonomy ~98% public, app-Postgres `cms.studio_*` = 100% customer → excluded, the
> public content lives in a *separate* Directus store): [`.agentspace/scratch/roadmap-research-2026-06-06.md`](../../.agentspace/scratch/roadmap-research-2026-06-06.md).

### The snapshot surfaces (prod-grounded 2026-06-06 — public-only, and why each needs a different mechanism)
- **`taxonomy`** — lives in the **per-stack** Postgres `skiller` schema but ships *empty* (normally loaded via the
  `importskills`/`importjobroles` cobras). Prod-measured: **~2.1 GB**, and **~98% public** — `skills` 42,763
  public / 794 private, `job_roles` 22,315 / 2,381. So the **tenant firewall** captures `organization_id IS NULL`
  (keeping the full public catalog, dropping the customer tail automatically); embeddings + translations carry no
  org column → scoped via the **public parent**. The snapshot is **Postgres→Postgres**: capture-once from a safe
  source, **bulk-`COPY` replay** per-stack (the M7a perf path). One refinement: `skill_embeddings` is 692 MB but
  the heap is 3.3 MB — ~689 MB is the **pgvector index**, so capture vectors verbatim and **rebuild the index on
  replay**. The *cleaner* surface — it proves the framework (M9b). (`data-dna.json`: `taxonomy` status `waived-m7c`.)
- **`content`** — the **public** simulation / skill-path **template library**. Prod correction: it is **not** the
  app-Postgres `cms` schema — `cms.studio_documents` + `cms.studio_tasks` are **100% org-scoped customer data (0
  public rows)** → **excluded** by the firewall. The public library lives in the **shared self-hosted Directus
  store** (`content.anthropos.work`) — *(M10 build later found this is the `directus` schema of the **same** prod
  Postgres, read-only via `marco_read`, not separate infra — see the M10 close note)*; the isolation guard hard-blocks writes to shared Directus,
  so replay needs a **per-stack content store**. The defining M10 decision (resolve early): per-stack Directus
  container fed from the captured snapshot vs replay straight into the per-stack Directus Postgres DB (Directus's
  own backing store is Postgres → stays in the per-stack-isolated class). The *highest-risk* surface in v1.2.
  (`data-dna.json`: `content` status `waived-m7c`.)

### M9a: Snapshot extension — capture-safe, public-only, manifest-cached framework + `/db-query` port
**Status:** `done` (2026-06-06) · **Shape:** `section` · **Complexity:** large · **Dir:** [m9a-snapshot-framework/](releases/archive/01.20-set-dressing/m9a-snapshot-framework/)
**Closed 2026-06-06** (build → harden 2 passes → close review → merged to `release/01.20-set-dressing`). Delivered the dedicated **`rosetta-extensions/stack-snapshot/`** section (9 Go pkgs + the `stacksnap` CLI, tagged `stack-snapshot-m9a` @ `1cc4dd2`): the capture/serialize/replay contract + portable `manifest.json` format; the **production-safe capture-source policy** (M9a-D3: cache-hit → dump-ingest [default] → safe primary read [MVCC] → restore/replica [AWS-gated upgrades]) with a bounded read-only session + catalog-first dry-run; the **tenant-data firewall** `AssertPublicOnly` (plan-time + post-capture, hard-fail on any tenant row, in-memory stash so nothing persists on a leak); the **`.agentspace` manifest-cached pluggable `SnapshotStore`** (localfs now, cloud/S3 = v1.3); the **data-DNA snapshot dimension** (`stack-seeding/dna/snapshot.go` — `snapshot-seeded` status that counts toward coverage + 5 two-sided fidelity operators) + the `/db-query` port. Proven end-to-end on the hermetic `reference-toy` surface. **556 Go test funcs (+147; stack-snapshot 128 new, stack-seeding 145→164)**, flake **0** (5× shuffled, `-race`), gofmt+vet clean. Close: scope/code/docs/tests **GREEN**; 1 finding fixed (the tag trailed HEAD by the harden commits → re-pointed + force-pushed); 5 adversarial scenarios recorded; deferral audit GREEN. Decisions M9a-D1…D7 + Q2/Q3/Q4. Retro: [m9a-snapshot-framework/retro.md](releases/archive/01.20-set-dressing/m9a-snapshot-framework/retro.md). **Next:** M9b (the real ~2.1 GB taxonomy surface).
**Goal:** A **dedicated, reusable `rosetta-extensions/stack-snapshot/` section** that captures a *public* reference
surface once from a **safe, low-impact source** (default a prod `pg_dump`), serializes it to a `.agentspace` manifest-cached store, and
replays it per-stack — with a tested **tenant-data firewall** (never customer data) + an **alignment extension that
measures replay fidelity**. Proven on a tiny reference surface (M0 toy-mirror discipline); ports the **`/db-query`**
skill as the prod-read foundation.
**Scope:**
  - In: **(note #1)** the dedicated **`stack-snapshot/`** section (capture + serialize + store + replay + the
    `stacksnap` CLI) — capture (a privileged prod **read**) decoupled from seeding (per-stack **writes**); the
    **snapshot contract + portable format** (per-table `COPY` payloads + `manifest.json`); **(note #2)** the
    **production-safe capture-source policy** (source-pluggable) — cache-hit (no prod read) → **(1) prod-`pg_dump`
    ingest** [default, zero new load] → **(2) safe throttled primary read** [MVCC = no write blocking] → **(3)
    restore-from-snapshot / (4) read replica** [zero-impact upgrades once eu-west-1 AWS/infra is wired], with
    bounded read-only sessions + a catalog-first dry-run (the **read half** the M7a guard lacks); **(note #3)** the **tenant-data firewall** `AssertPublicOnly`
    + a **public-only/provenance gene** (hard-fail on any captured tenant row); **(note #4)** the **`.agentspace`
    manifest-cached, pluggable `SnapshotStore`** (localfs now, cloud/S3 = v1.3; no GB blobs in git); the **data-DNA
    extension** — `snapshot-seeded` status + a **snapshot-fidelity gene class** (row-count / structural conformance
    / referential integrity / **embedding-dimension integrity**); **(note #5)** the **ported `/db-query` skill** +
    `corpus/ops/db-access.md` (the MCP-tool **and** pgpass/psql paths); a tiny reference surface proving
    capture→store→replay→fidelity-gate end-to-end.
  - Out: the taxonomy surface (M9b); the Directus content snapshot (M10); recipes/presets (M11); the cloud store
    backend + AI-generated content + shareability (v1.3).
**Depends on:** v1.1's M7a (isolation guard + perf path) + M7b (the data-DNA harness it extends). **Parallel with:** none (gates M9b + M10 + M11).
**Open questions:** capture source resolved (M9a-D3, investigated 2026-06-06 — no replica + no local AWS creds →
default **prod-`pg_dump` ingest**, fallback **safe primary read**; restore-from-snapshot/replica = later upgrades);
the manifest schema + cache-staleness rule; embedding capture (vectors verbatim, **rebuild the ~689 MB pgvector
index on replay**); the `SnapshotStore` interface so the v1.3 cloud swap is a backend change.
**KB dependencies:** `corpus/ops/seeding-spec.md` (the framework + isolation boundary + the DAG node), `corpus/architecture/alignment_testing.md` (the data dimension to extend), `corpus/ops/staging_from_dump.md` (the full-dump **anti-pattern** to contrast), the source `db-query` skill (ported).
**Delivers → `corpus/ops/snapshot-spec.md`** (net-new — the extension + capture/replay contract + capture-source policy + tenant firewall + the `.agentspace` manifest store) **+ `corpus/ops/db-access.md`** (net-new) **+ the `/db-query` skill** **+ extends `corpus/architecture/alignment_testing.md`** (the snapshot-fidelity + public-only genes).

### M9b: Taxonomy snapshot (the first real surface)
**Status:** `done` (completed 2026-06-06) · **Shape:** `section` · **Complexity:** large · **Dir:** [m9b-taxonomy-snapshot/](releases/archive/01.20-set-dressing/m9b-taxonomy-snapshot/)
**Closed 2026-06-06** (build → 3 harden passes → close attempt 1 BLOCKED at a RED deferral audit → user fate decision → close attempt 2 clean → merged to `release/01.20-set-dressing`). Delivered: the **10-table public taxonomy surface** in FK replay order (`categories`, `job_role_categories` → `specializations` → `skills`, `job_roles` → embeddings/translations/`job_role_skills`), `org_id IS NULL` filters + **parent-scoped predicates** for column-less tables (`TableSpec.PublicViaFK` + `firewall.ParentScopeFilter` — closing M9a's empty-filter gap, M9b-D2; `job_role_skills` both-endpoints, M9b-D3), **pgvector index rebuilt via `REINDEX` on replay** (vectors verbatim, dim 1536 in the manifest, M9b-D5), the **manifest→`CapturedTable` fidelity bridge** + `datadna measure-snapshot` CLI (M9b-D6), and the **`TaxonomySnapshotSeeder` DAG node** (`… → taxonomy → activity`, M9b-D7). Coverage `waived-m7c → snapshot-seeded-m9b` (all 5 fidelity operators). PR-review fix: the parent-leak probe must AND the capture filter (was scanning the whole table → false abort). **Deferral fate (M9b-D9):** the **offline pg_dump-FILE reader was DROPPED by the user** (an M9a→M9b repeat/aged-out deferral) — restore-then-`--dsn` + the safe primary read cover the need; the dead `--dump` flag was removed and docs+code corrected (dump-ingest = restore-the-dump-into-Postgres-then-`--dsn`, no offline reader, not seeded to v1.3). Decisions M9b-D1…D9. Extensions: tag `stack-snapshot-m9b` @ `55ee0e6`. Go test funcs 556→**635** (stack-snapshot 128→167, stack-seeding 164→204); flake 0; both modules `-race` green. Retro: [m9b-taxonomy-snapshot/retro.md](releases/archive/01.20-set-dressing/m9b-taxonomy-snapshot/retro.md).
**Goal:** Prove the M9a framework on the **real ~2.1 GB taxonomy surface** — capture the *public* skiller catalog
from a safe source, bulk-`COPY` replay per-stack, **rebuild the pgvector index on replay**, fidelity- + public-only
gated — driving data-DNA coverage from `waived` to its first `snapshot-seeded` surface.
**Scope:**
  - In: **public taxonomy capture** — `skiller.{categories,specializations,skills,job_roles}` filtered
    `organization_id IS NULL` (full public catalog, customer tail dropped); `{skill,job_role}_embeddings` (vectors
    only) + `{skill,job_role}_translations` + `job_role_skills` via the **public-parent** join; **bulk-`COPY`
    replay** per-stack (M7a perf path, per-stack-isolated only); **pgvector index rebuild on replay** (carry
    vectors verbatim, don't transport the ~689 MB index); the **taxonomy fidelity + public-only genes**;
    wiring into the `stack-seeding` DAG node. Coverage `waived → taxonomy-seeded`.
  - Out: the Directus content surface (M10); recipes/presets (M11); recompute of embeddings (v1.3).
**Depends on:** **M9a** (the `stack-snapshot` extension + capture-source policy + firewall + store + fidelity genes). **Parallel with:** none (gates M10 + M11).
**Open questions:** keyset-chunked vs single streamed `COPY` for skills/embeddings (size via the dry-run); the
pgvector index params to rebuild (match prod, record in the manifest); `job_role_skills` referential integrity vs
the public-skill set.
**KB dependencies:** `corpus/ops/snapshot-spec.md` (M9a's contract), `corpus/services/skiller.md` (the taxonomy schema + embeddings + translations), `corpus/ops/seeding-spec.md` (the perf path + the DAG node), `corpus/architecture/alignment_testing.md` (the genes).
**Delivers → extends `corpus/ops/snapshot-spec.md`** (the taxonomy capture/replay path) **+ updates `corpus/ops/seeding-spec.md`** (taxonomy promoted `waived` → `snapshot-seeded`).

### M10: Directus content snapshot-replay
**Status:** `done` (completed 2026-06-06) · **Shape:** `section` · **Complexity:** large (highest-risk) · **Dir:** [m10-content-snapshot/](releases/archive/01.20-set-dressing/m10-content-snapshot/)
**Closed 2026-06-06** (spike → build → 5 harden passes → close clean → merged to `release/01.20-set-dressing`). Delivered the **second real surface**: the **9-table public Directus content** library (`simulations` + `skill_paths` scope-bearing roots → `resource` pure-reference → `roles`/`sim_tasks`/`sequences` parent-scoped → `task_checks`/`task_sub_checks` **multi-level chains** → `sequences_roles` via BOTH). Generalized the firewall to a **per-surface `PublicPredicate`** (the spike-flagged org-only gap, M10-D1; taxonomy byte-for-byte unchanged via the zero-value→`DefaultPredicate` fallback); the directus predicate `private=false AND tenant_id IS NULL AND status='published'` (M10-D3, prod-verified); **multi-level parent chains** via `ParentScope.ParentFilter` chasing column-less intermediates to the scope-bearing root in one subquery (M10-D4); the **per-stack Directus store fork** = bootstrap→replay→boot on the per-stack Directus-backing Postgres `directus` schema (`PerStackIsolated`, M10-D2), `EnvContract.Validate` hard-rejecting any prod-Directus target; media **REFS** landed (ref columns + `ReferencedFilesFilter` + the 1,311 public `directus_files` rows + a local-storage/placeholder adapter) with blob **BYTES** S3-gated → **Fate 2 → v1.3** (M10-D5, confirmed-covered by the roadmap-vision cloud-store seed); the **content fidelity + public-only genes** (`ReplayedNonPublicRows`); the **`ContentSnapshotSeeder`** DAG node + the **`sim_id`/`skill_path_id`/`resource_id` linkage resolver** (v1.1 free-value refs now resolve against real replayed public templates; free-value fallback when no snapshot). Capture source **self-resolved**: the `directus` schema lives in the SAME app Postgres, read-only via `marco_read` (the spike's "separate store" premise disproven, KB-1 corrected). Coverage `content` **`waived-m7c → snapshot-seeded-m10`** → **NOTHING left waived → 100% over the full catalog (the v1.2 set-dressing thesis complete)**. Close found **2 findings** (both fixed): the stack-snapshot README package index was missing the `taxonomy`+`directus` rows (per-unit handbook contract — a latent M9b miss caught here), and the M10-D5 decision record. Decisions M10-D1…D5. Extensions: tag `stack-snapshot-m10` @ `df410f5`. Go test funcs 635→**701** (stack-snapshot 167→207, stack-seeding 204→230); flake 0 (5/5 both modules); both modules `-race` green; gofmt+vet clean. Retro: [m10-content-snapshot/retro.md](releases/archive/01.20-set-dressing/m10-content-snapshot/retro.md).
**Goal:** Capture the shared-Directus content library and replay it into a **per-stack content store** — never
touching shared Directus — taking data-DNA coverage to **100% of the full catalog** (the last `waived` surface
promoted to `snapshot-seeded` + fidelity-gated).
**Source correction (prod, 2026-06-06):** the content source is **not** app-Postgres `cms` — `cms.studio_documents`
+ `cms.studio_tasks` are **100% org-scoped customer data (0 public rows)** → **excluded** by the firewall. The
public template library lives in the **separate self-hosted Directus store**; M10 captures **only its public/global
templates**.
**Scope:**
  - In: the **per-stack content-store decision** resolved + built (per-stack Directus container vs direct
    per-stack Directus-Postgres replay — the defining fork, resolve in the first iter/spike); the **public content
    capture** (export the public/global Directus templates + media references from the separate Directus store — a
    privileged read via M9a's capture-source policy + tenant firewall, isolation-clean); the **content replay
    seeder** wired into M9a's snapshot framework + the seeder DAG (M7a), respecting the guard; the **content
    fidelity + public-only genes** in the data-DNA; the **`sim_id`/`skill_path_id`/`resource_id` linkage** so the
    v1.1 session/assignment seeders' content refs resolve against the real **public** templates (closing the
    "free-value refs" gap).
  - Out: app-Postgres `cms.studio_*` customer content (excluded — tenant data); AI-generated/authored content
    (v1.3 — this replays *real* captured public content, it does not generate); recipes/presets (M11); shareability (v1.3).
**Depends on:** **M9a + M9b** (the snapshot framework + fidelity DNA + the `stacksnap` CLI + the proven taxonomy surface). **Parallel with:** none (M11 curates its output).
**Open questions:** the content-store fork (above) — the load-bearing decision; where the public Directus templates
physically live + how the public/global subset is identified (confirm against the Directus store); whether
media/blobs are in-scope or refs-only for the demo MVP (S3-private is per-stack-isolated, so blobs *can* be
replayed — confirm at build); how much of the collection set the demo needs (the believable subset, per the M7c "reachable" discipline).
**KB dependencies:** `corpus/ops/snapshot-spec.md` (M9a's contract), `corpus/services/cms.md` (Directus) + `corpus/ops/db-access.md` (the Directus store connection), `corpus/ops/seeding-spec.md` (the isolation guard + the session/assignment content refs), `corpus/services/{jobsimulation,skillpath}.md` (the consumers of `sim_id`/`skill_path_id`).
**Delivers → extends `corpus/ops/snapshot-spec.md`** (the public-Directus content path + the store decision) **+ updates `corpus/ops/seeding-spec.md`** (content surface promoted from `waived` to `snapshot-seeded`).

### M11: Richer-world recipes, presets + corpus polish
**Status:** `done` (completed 2026-06-06) · **Shape:** `section` · **Complexity:** medium · **Dir:** [m11-richer-recipes/](releases/archive/01.20-set-dressing/m11-richer-recipes/)
**Closed 2026-06-06** (build → 1 harden pass → close review → merged to `release/01.20-set-dressing`). **The LAST milestone of v1.2** — the product/discoverability layer (the M8-analog) that curates the 100%-coverage full-fidelity world (M9a/M9b/M10) into usable recipes/presets, with **zero new production code path**. Delivered: the 3 **seed presets** (`small-200`/`mid-500`/`large-1k`) gained a documented FULL-FIDELITY PREREQUISITE comment-header (replay `taxonomy` + `directus` BEFORE seeding; graceful structural-only degradation without) — **presets stay purely structural** (snapshots are stack-global reference data, not org-scoped → no schema field, M11-D3); the **set-dressed `corpus/ops/demo/` recipe family** (both org-onboarding + skill-progression recipes gained a `/demo-snapshot` replay step + call out the real catalog/templates; the FALSE "waived/future-v1.2" note in `recipe-skill-progression.md` rewritten to the shipped 100%-coverage reality) + a **new `recipe-snapshot-world.md`** (the capture→replay→set-dressed-world curator walk-through); the **`/demo-snapshot` skill** (M11-D1, resolves M11-Q2 — a NEW skill, not a `/demo-seed` extension: capture = privileged prod READ vs seeding = per-stack WRITE; `replay` headline / `capture` rare-maintenance / `status`); corpus cross-links (CLAUDE.md skill-table row + key-docs entry; `snapshot-spec` ↔ `seeding-spec` ↔ the demo family bidirectional; data-DNA reads 100% throughout); the **§5 release-close hygiene carry** — the stale `stacksnap --help` text fixed (M11-D4: framework tag M9a/M9b → M9a/M9b/M10 + the registered-but-unlisted `directus` surface now listed). Decisions M11-D1…D4 + Q1/Q2 resolved. Extensions: tag `stack-snapshot-m11` @ `1e18df6`. Go test funcs 701→**708** (stack-snapshot 207→212: the `--help` contract pins + docs↔parser flag-drift guard; stack-seeding 230→232: shipped-preset strict-parse/validate + size-order); flake **0** (5/5 both M11-touched packages); both modules `-race` + gofmt + `go vet` clean. Close: deferral re-audit **GREEN** (1 inherited DEF-M10-01 unchanged, 0 repeat/aged/new — M11 added zero deferrals, all Fate-1); scope/code/docs/tests **GREEN** with **0 findings**; 4 adversarial scenarios recorded (each mutation-pinned). Retro: [m11-richer-recipes/retro.md](releases/archive/01.20-set-dressing/m11-richer-recipes/retro.md). **v1.2 is now complete → next: `/developer-kit:close-release`** (merge `release/01.20-set-dressing` → `main`, tag `v1.2`).
**Goal:** The product/discoverability layer that closes v1.2 (the M8-analog): make the full-fidelity worlds
*usable + discoverable* — refresh presets + recipes so a demo curator gets a real-taxonomy, real-content world
out of the box, and update the corpus to reflect 100% coverage.
**Scope:**
  - In: refreshed **seed presets** (small/mid/large) that now include the taxonomy + content snapshots; an updated
    **`corpus/ops/demo/` recipe family** (the end-to-end recipes now showcase a *set-dressed* world — real skills
    in the catalog, real simulations/skill-paths behind the seeded sessions); a **`/demo-snapshot` (or extended
    `/demo-seed`) skill** driving the `stacksnap` CLI; cross-linking + corpus updates (the data-DNA now reads 100%);
    the **release-close hygiene** carry (any small items surfaced in M9a/M9b/M10).
  - Out: new snapshot surfaces (M9a/M9b/M10 own them); AI-content + shareability (v1.3).
**Depends on:** **M9a + M9b + M10** (curates their output). **Parallel with:** none (the closing milestone before `/developer-kit:close-release`).
**Open questions:** whether snapshot capture is a curator step or a manifest-cached refresh (decide with M9a's capture-source policy); `/demo-seed` extension vs a new `/demo-snapshot` skill.
**KB dependencies:** `corpus/ops/demo/README.md` + the recipes, `corpus/ops/seeding-spec.md`, `corpus/ops/snapshot-spec.md`, the `/demo-seed` skill.
**Delivers → refreshes `corpus/ops/demo/`** (recipes + presets to full-fidelity) **+ the `/demo-snapshot` skill + the CLAUDE.md skill table.**

### Execution graph (v1.2)
```
v1.2 "set dressing" — richer demo worlds: the real *public* taxonomy + content, measured-faithful, to 100% coverage
   M9a (stack-snapshot framework: capture-safety + tenant firewall + .agentspace store + /db-query + fidelity-DNA)
        └─→ M9b (taxonomy surface: public skiller + embeddings, rebuild index) ─→ M10 (public Directus content) ─→ M11 (recipes + presets + corpus)
```
**Sequential.** M9a lands the **dedicated `stack-snapshot` extension** + the capture-source policy + the tenant
firewall + the `.agentspace` manifest store + the fidelity-DNA + the `/db-query` port, proven on a toy surface.
M9b proves it on the cleaner ~2.1 GB taxonomy (coverage waived→taxonomy-seeded). M10 takes the harder public
Directus content surface to **100% coverage**. M11 curates the full-fidelity worlds into usable recipes/presets +
closes the release. No parallel tracks — one extension + one data-DNA; serializing keeps the merge surface clean
(the v1.1 spine discipline).

### Risks (v1.2)
- **(M9a, blocks-prod-safety)** a capture that **reads the hot primary** under load, or **leaks a tenant row**.
  Mitigate: the capture-source policy (cache-hit → prod-`pg_dump` ingest → safe throttled primary read [MVCC = no
  write blocking] → restore-from-snapshot/replica upgrades) + bounded read-only sessions; the **tenant firewall**
  `AssertPublicOnly` + public-only gene — tested gates, not conventions.
- **(M10, blocks-100%-coverage)** the **content-store fork** + locating the public Directus template subset (a
  *separate* store). Mitigate: resolve the store decision in an M10 spike *first* (Directus's backing store *is*
  Postgres → a per-stack Directus-Postgres replay stays in the isolated class); fall back to refs-only believability
  if a full Directus stand-up proves too heavy for the demo MVP.
- **(M9b, scope)** **embedding fidelity** — carrying pgvector embeddings verbatim (dimension + value integrity) and
  **rebuilding the ~689 MB index on replay** rather than transporting it. Mitigate: capture vectors verbatim
  (offline + deterministic), gate the embedding-dimension gene; never recompute (that's AI-content, v1.3).
- **(note #4 → v1.3)** the local `.agentspace` cache doesn't share across machines / scale. Mitigate: the
  `SnapshotStore` interface keeps the **cloud/S3 swap** a v1.3 backend change (the manifest already addresses by location).
- **(cross-cut)** **the isolation guard's missing read half** — capture reads a reference source. Mitigate: capture
  is read-only + audited (the capture-source policy); replay writes only to per-stack-isolated stores, the existing
  3-layer guard asserts clean (extend `AssertClean` to cover snapshot replay).

### Open decisions (resolve during build)
The **capture source** — **resolved** M9a-D3 (user 2026-06-06): default **prod-`pg_dump` ingest** → fallback **safe
throttled primary read**; restore-from-snapshot/replica = zero-impact upgrades once eu-west-1 AWS/infra is wired;
the manifest schema + cache-staleness rule — M9a; embedding capture
(verbatim + rebuild-index-on-replay) — M9b (lean verbatim); the **per-stack content-store fork** (per-stack
Directus container vs direct Directus-Postgres replay) — M10, the defining decision; identifying the public
Directus template subset — M10; media/blobs in-scope vs refs-only — M10; `/demo-seed` extension vs a new
`/demo-snapshot` skill — M11.

## Done — v1.1 "show floor" (SHIPPED 2026-06-05 · tag `v1.1`)

**Theme (broadened 2026-06-04):** v1.0 made the platform run *without* Clerk; v1.1 started as "disposable
demo stacks" (M3 ✅) and now becomes **the platform-operations extension framework** — consolidate the repo
constellation into **two repos** (`rosetta` = the platform corpus + dev-env skills; `rosetta-extensions` = a
monorepo of operations sections), then deliver the seeded-demo capability *and* generalize the pattern to dev.
Everything stays **additive — zero change to any read-only platform repo**.

**Refactored 2026-06-04** (after M3 shipped, to keep the constellation from exploding): the standalone
`clerkenstein` + `rosetta-demo` repos collapse into `rosetta-extensions/{clerkenstein,demo-stack,…}`; the
former M4 (seeding) → **M7**, former M5 (recipes) → **M8**; new structural milestones M4–M6 inserted. Decisions:
**git subtree, history-preserving** (M4-D1) · **delete the old repos, not archive** (M4-D2, user) · **the
alignment framework stays in rosetta** (M4-D3) · per-demo clones (M3-D1) · clone-at-release-tag (M3-D3).

**Seeding redesigned 2026-06-04** (M3–M6 all shipped): the user asked to make seeding robust/resilient/drift-proof/
fast/**production-safe**, so the single `section` M7 splits into **M7a → M7b → M7c** (a section + section +
iterative "mix"). 3 research agents over the platform grounded it: the prod-pollution boundary is *small + fixed*
(Directus, S3-public, live Clerk/external SaaS — everything in the per-stack Postgres is isolated); the M0 alignment
pattern *extends to data* (new structural operators + schema-as-source); the perf bottleneck is *DB-IO, not CPU*
(Go-link-ent + `COPY` + fan-out; Rust buys nothing). Decisions: **3-way split, all in v1.1** (M7a-D1, user chose
keep-in-v1.1 over a v1.2 spin-out) · **the isolation guard is the load-bearing deliverable** (M7a-D2) · **extend
M0 to a data dimension, don't fork it** (M7b-D1) · **the data-DNA is the catalog that drives the fleet** (M7b-D2)
· **the fleet is iterative, gated on data-DNA coverage** (M7c-D1).

### M3: Disposable multi-instance demo stacks ✅ DONE (2026-06-03; extended close 2026-06-04)
**Status:** `done` · **Shape:** `section` · **Dir:** [m3-demo-stacks/](releases/archive/01.10-show-floor/m3-demo-stacks/)
Spun up `demo-N` as isolated, Clerkenstein-wired full stacks; the full Clerk-free injected stack + migrate are
LIVE-PROVEN; the deployment/injection alignment surface (`clerk-deploy-1`, 7/7) landed. 78 demo-stack tests, 218
clerkenstein funcs. **Delivered** `corpus/ops/rosetta_demo.md` + `/demo-*` skills.

### M4: Consolidate into the `rosetta-extensions` monorepo ✅ DONE (2026-06-04)
**Status:** `done` · **Shape:** `section` · **Dir:** [m4-consolidate-extensions/](releases/archive/01.10-show-floor/m4-consolidate-extensions/)
Created the **`rosetta-extensions`** monorepo (private, 73 commits); `git subtree`-imported `clerkenstein` +
`rosetta-demo`(→`demo-stack`) **with full history preserved**; the `knowledge/` nav; thinned rosetta to pointers;
fixed a +1-depth path break the verify gate caught (M4-D4); verified under the new paths (78 demo-stack tests +
deploy gate 7/7); pushed; **removed the old `clerkenstein` + `rosetta-demo` repos** (local + org, 404). Decisions
M4-D1 (subtree) / D2 (delete-not-archive) / D3 (alignment framework stays in rosetta) / D4 (path-depth fix).

### M5: Extract the reusable `stack-injection` layer ✅ DONE (2026-06-04)
**Status:** `done` · **Shape:** `section` · **Dir:** [m5-stack-injection/](releases/archive/01.10-show-floor/m5-stack-injection/)
Extracted the generic injection (`inject.py`, `gen_injected_override.py`, `apply-authn.sh`) into
`rosetta-extensions/stack-injection/`, consumable by any stack with a **demo-ON / dev-OFF** toggle; the mock stayed
in clerkenstein (dependency runs stack-injection→clerkenstein, M5-D1); the port-offset engine stayed in demo-stack
(M5-D2, settles the M4 open question — moves to shared in M6). Split the tests, repointed the consumers; **78
preserved**, flake 3/3, deploy gate 100%/100%.

### M6: `dev-stack` — tooled local dev environment ✅ DONE (2026-06-04)
**Status:** `done` · **Shape:** `section` · **Dir:** [m6-dev-stack/](releases/archive/01.10-show-floor/m6-dev-stack/)
Extracted the shared port-offset engine into a new **`stack-core/`** section (settles the M5-routed question —
demo + dev share it, M6-D1) and added a focused **`dev-stack/`**: isolated dev stacks (`dev-N`, offset ports,
guarded `-p dev-N`), **real Clerk by default**, Clerkenstein injection **optional** (reuses stack-injection).
Scoped to the proven value (M6-D2 — not speculative multi-dev). **87 tests** (+9), flake 3/3, deploy gate 100%/100%.

### M7a: Seeding framework + production-isolation safety ✅ DONE (2026-06-04)
**Status:** `done` · **Shape:** `section` · **Complexity:** large · **Dir:** [m7a-seeding-framework/](releases/archive/01.10-show-floor/m7a-seeding-framework/)
Built `rosetta-extensions/stack-seeding/` — a host Go module that seeds a stack by talking **directly to its
Postgres** (offset port, `COPY`; *not* ent-linking — `app/internal/bootstrap` is internal, unimportable, M7a-D3)
behind a **3-layer production-isolation guard** (CheckWrite · PreflightEnv · AssertClean). **LIVE-PROVEN**: a
fresh injected `demo-1` → `migrate-demo.sh` (now bootstraps the global Sentinel policy) → `stackseed` (org + 1000
users + the real `user_clerkenstein` identity + the casbin `g2` grant, isolation audit clean) → authenticated
login returns **HTTP 200** (`membershipsCount: 1001`). The proof caught + fixed **2 real bugs** (the g2 arg-order;
the missing global-policy bootstrap — M7a-D4). **68 tests**, all gates green. Delivered `corpus/ops/seeding-spec.md`.

### M7b: The data-alignment dimension ("data DNA") ✅ DONE (2026-06-04)
**Status:** `done` · **Shape:** `section` · **Complexity:** medium · **Dir:** [m7b-data-dna/](releases/archive/01.10-show-floor/m7b-data-dna/)
Extended the **M0 alignment framework** to a **data** dimension — the `datadna` harness (`rosetta-extensions/
stack-seeding/dna/`) that (a) enumerates the seedable surfaces (**4 seeded + 6 planned** — the M7c checklist) and
(b) measures a seeder's output conforms to the platform's **current schema** via **structural operators**
(type-match / constraint-satisfied [NOT-NULL + UNIQUE] / fk-valid / row-count) with **schema-as-source via
introspection**. A separate harness, not an alignctl runner (M7b-D3). **PROVEN live** on the M7a-seeded `demo-1`:
`measure` **100% / Critical 100%** across the 4 seeded surfaces; `diff` flags an injected column (exit 1) and
reads clean on revert. Caught + fixed the planned-surface introspection bug; hardened the UNIQUE leg (M7b-D4).
**dna 49 + cmd/datadna 10 + pg 17 tests.** Delivered the data dimension into `corpus/architecture/alignment_testing.md`.

### M7c: The seeder fleet, to a coverage gate ✅ DONE (2026-06-05, gate-met-over-reachable + waiver)
**Status:** `done` · **Shape:** `iterative` · **Complexity:** large · **Dir:** [m7c-seeder-fleet/](releases/archive/01.10-show-floor/m7c-seeder-fleet/)
Built the fleet across 5 iters (TOK-01 strategy → jobsim-sessions → skillpath-sessions → assignments → activity),
each a deterministic **backdated-activity** seeder (time-distributed, pass/fail per `pass_rate`, content refs as
free values — the believability core is reachable **without** the shared Directus). Drove data-DNA coverage
**40%→80%**, promoting each surface planned→seeded + conformance-gated. **Gate: 3 of 4 met outright** — (a)
login→**200** · (c) full 8-seeder seed **0.69s** (~8500 rows, <2min) · (d) isolation **clean**; (b) coverage is
**100% over the 8 reachable surfaces / critical 100%**, with **taxonomy + content waived** (the hard line —
skiller snapshot + shared Directus; Re-scope trigger, user-confirmed → ~v1.2). Caught + fixed 2 live bugs (the
skillpath UNIQUE constraint; the introspect-load harness bug). **20 seeder / 145 module tests.** Delivered
`rosetta-extensions/stack-seeding/seeders/` + the `waived` data-DNA status.

### M8: Corpus + use-case recipes + polish ✅ DONE (2026-06-05) — LAST v1.1 milestone
**Status:** `done` · **Shape:** `section` · **Complexity:** medium · **Dir:** [m8-corpus-recipes/](releases/archive/01.10-show-floor/m8-corpus-recipes/)
The consolidation/discoverability layer: a **`corpus/ops/demo/` family** (index + 3 end-to-end recipes —
enterprise-onboarding, skill-progression, browser-login [which lands the 2 M3-deferred injection recipes: the
`api.clerk.com` cert-redirect + the browser-login walk-through]); **3 seed presets** (small/mid/large, mid-500 +
large-1k seed-proven end-to-end); the **`/demo-seed` skill** + the CLAUDE.md skill table; the v1.0
**express-gate CI carry-forward** wired into clerkenstein `alignment.yml` (**validated 9/9** locally); and
cross-linking from corpus/README + root README + CLAUDE.md (all doc links resolve). **Next:** `/developer-kit:close-release`.

### Execution graph (v1.1)
```
v1.1 "show floor" — the platform-operations extension framework (demo + dev, in 2 repos)
   M3 ✅ ─→ M4 ✅ (consolidate) ─→ M5 ✅ (stack-injection) ─→ M6 ✅ (dev-stack)
                                            └──→ M7a (framework+safety) ─→ M7b (data-DNA) ─→ M7c (seeder fleet) ─→ M8 (recipes)
```
**Sequential.** M4–M6 shipped (the extension framework + demo/dev stacks). M7a lands the framework + the
isolation guard (a usable, safe demo); M7b builds the data-DNA catalog that lists + gates the seeders; M7c drives
the fleet to the coverage gate; M8 curates the output.

### Risks (v1.1)
- **(M7a, blocks-prod-safety)** a single un-guarded **shared-write reaching prod** (Directus / S3-public bucket) —
  mitigate with the hard isolation guard + the clean-audit assertion as a tested acceptance gate, not a convention.
- **(M7a, scope)** linking the platform's `app/internal/bootstrap`/ent client into a `rosetta-extensions/` Go
  module without a platform edit — confirm the import path early (fallback: `go run` CLIs, slower).
- **(M7b)** trustworthy schema-as-source — get ent introspection / `atlas inspect` golden right or the drift diff lies.
- **(M7c, scope)** the heaviest build: ~8–10 seeders + 1k-scale `COPY` perf + backdating fidelity, each gated on
  conformance — the believable-demo *subset* of surfaces is the real target (waive unreachable genes, don't chase 100%).

### Open decisions (resolve during build)
Directus snapshot-replay vs hard-block-and-skip for the demo MVP (M7a); ent-introspection vs `atlas inspect`
golden for schema-as-source (M7b); whether seed presets ship in M7c or M8; external shareability (Tailscale vs
ingress); the AI-content STRETCH trigger (now firmly v1.2, not M8).

## Done — v1.0 "body double" (SHIPPED 2026-06-03 · tag `v1.0`)

> **Shipped 2026-06-03.** All six milestones closed-on-gate / completeness-complete and merged to `main`;
> `release/01.00-body-double` deleted. Release records archived under
> [`releases/archive/01.00-body-double/`](releases/archive/01.00-body-double/) (review · retro · metrics ·
> lockfile · stats). Headline: a *measured* drop-in Clerk mock at **100%/100% on all three surfaces**
> (Go · JS/FAPI · `@clerk/express`), built by a first-class alignment framework, zero platform-code change.
> Close-release caught + fixed 1 blocker (an `@clerk/express` gate regression from the M2c close) — see the
> release retro.

**Theme:** Clerk authentication is the friction that blocks fast, throwaway demos. v1.0 delivers
**Clerkenstein** — a drop-in mock that mirrors the exact Clerk interface the platform uses, with
security/sync disarmed, injected via build-time `go.mod replace` + skip-worktree so **every platform
repo keeps "thinking" it uses Clerk with zero source changes**. The novelty: Clerkenstein isn't a
hand-built mock — it's the **first mirror produced by a reusable, measurable alignment process**
(M0). We don't just claim the stand-in is faithful; we **score** it (0–100%) against the real Clerk
and CI-gate that score against drift. This also removes Clerk's API rate limit as the blocker for
scale data-seeding in v1.1.

**Decided at design (2026-06-02):** two-version split (Clerkenstein first); **alignment is a
first-class test class** with its own framework (M0); **M1 is iterative** (its exit gate is an
alignment score); M2 frontend = attempt the fake Clerk FAPI server, **fall back to the real dev Clerk
app for the browser session** (backend stays fully mocked) if base-URL override proves too fragile.

### Alignment vocabulary (the M0 model, referenced by M1/M1b)
- **Target** — an engine exposing a surface. **Source target** = the canonical engine, version-pinned (Clerk `clerk-sdk-go/v2 @ v2.6.0`). **Mirror target** = our reimplementation (Clerkenstein).
- **Capability** — one endpoint/function of the source surface *(axis 1)*. **Variant** — one input/scenario class for a capability (standard + corner + error) *(axis 2)*.
- **Alignment test** — one **(capability × variant)** pair; feeds identical input to both targets and asserts behavioral equivalence. A **third test class** alongside unit & integration; **tagged** so it's parseable/countable/runnable as its own suite.
- **Alignment DNA** — the officially-enumerated complete set of (capability × variant) **genes** for a source target at a version; the machine-readable manifest that *defines* faithfulness and is the score's denominator.
- **Alignment score** — `aligned genes ÷ total genes × 100`, with a per-capability rollup. 100% = behaviorally indistinguishable across the whole DNA.

### M0: Alignment measurement framework
**Status:** `done` (2026-06-02)
**Shape:** `section`
**Goal:** A reusable, engine-agnostic process — two skills + a test class + a manifest format — that measures how faithfully any mirror reproduces any source engine, producing a 0–100% alignment score. This is the foundation M1 builds on and M1b reuses.

**Closed 2026-06-02** (build S1–S5 → harden 2 passes → close review → merged to `release/01.00-body-double`). Delivered: `test/alignment/` — `alignctl` (stdlib-only Go, builds/runs offline) with `run`/`capture`/`dna list|diff|validate`, engine-agnostic via a pluggable `--runner`; the 4 equivalence operators + weighted score (overall + separate critical gate); record/replay goldens; `internal/canon` precision-safe canonicalization; the `//go:build alignment` test class; and a toy reference proving end-to-end detection (**86.7% / 100% critical**, catches `Greet/padded-name`). Plus `/align-dna` + `/align-run` skills and `corpus/architecture/alignment_testing.md`. Open questions resolved: DNA format = **JSON** (M0-D1); capabilities enumerated from the consumed surface; goldens live per-mirror-repo. Close-review adversarial pass found + fixed a path-traversal must-fix + score-overflow (M0-D7); 45 test funcs (3 fuzz), 5/5 flake gate, core coverage 83–98%. Decisions M0-D1…D7. Retro: [m0-alignment-framework/retro.md](releases/archive/01.00-body-double/m0-alignment-framework/retro.md). Resolved repo split: framework in rosetta; the Clerk DNA/tests/mirror land in the `clerkenstein` repo (M1).
**Scope:**
  - In:
    - **`/align-dna` skill** (build & update alignment targets): given a source framework + version, pull the pinned source into `.agentspace/`; enumerate the **consumed** capabilities (scoped to what the platform calls, not the whole SDK); enumerate standard + corner-case **variants** per capability; emit/update the **Alignment DNA** manifest (each gene: input fixture, expected-shape descriptor, equivalence operator, criticality weight); **diff DNA across source versions** (added/removed/changed genes); **scaffold alignment-test stubs from the DNA** so tests never drift from the manifest.
    - **`/align-run` skill** (measure alignment of 2 targets): given a DNA + source version + mirror, pull the source, run every gene against **both** targets, assert equivalence per the gene's operator, compose the **0–100% score** + a per-capability divergence report.
    - the **alignment test-class convention** (tagging/marking so tests are discoverable + countable, distinct from unit/integration), the **DNA file format**, the **equivalence operators** (exact / same-shape / normalized / same-error-class), and **record/replay (golden capture)** support so a live-SaaS source can be measured reproducibly offline.
    - a **tiny toy reference mirror** (≈2 capabilities) proving the framework runs + scores end-to-end, independent of Clerk.
  - Out: the Clerk DNA + the real Clerkenstein mirror (M1); drift CI wiring (M1b); the JS surface (M2).
**Depends on:** none.
**Parallel with:** none (gates M1, M1b).
**Estimated complexity:** large
**Open questions:** DNA manifest format (YAML vs Go structs); how capabilities are enumerated (parse source surface vs curated list); where golden captures live + how they're refreshed.
**KB dependencies:** none new (greenfield — alignment is a documentation blind area).
**Delivers → `corpus/architecture/alignment_testing.md`:** the alignment test class, the DNA format, the two skills, equivalence + record/replay — the canonical reference (net-new doc).

### M1: Clerkenstein backend mirror (Go)
**Status:** `done` (2026-06-03, **closed-on-gate**)
**Shape:** `iterative`
**Goal:** The first real mirror — a drop-in Go stand-in for `colony/authn`'s provider + the Clerk `orgclient`, built *by* the M0 process and injected via `go.mod replace` (zero platform-repo edits), so backend services authenticate with one universal credential and locally-minted JWTs.

**Closed 2026-06-03** (5 iters: bootstrap tok TOK-01 → DNA → authn twin → critical orgclient → standard orgclient → **gate** → final harden → close). The **Clerkenstein backend mirror** (in the gitignored `anthropos-demo/clerkenstein` repo, its own git) scores **100% alignment / 100% critical** against the `clerk@2.6.0` DNA (22 genes), built offline. authn implements the real `colony/authn.Provider` (HS256, one universal key); orgclient is a disarmed in-memory twin. Score arc: 0 → 21.1 → 68.4 → **100%**. Final harden: authn + orgclient **0 → 100%** unit coverage (+1 fuzz, 0 bugs). Decisions: D1 hybrid goldens; iter-01-D1 authn injects via `go.mod replace` whole-colony; **M1-D2 orgclient injects via a fake-Clerk-API-server → routed to M2** (shared HTTP-interception with the JS side). Delivered `corpus/services/clerkenstein.md`. Retro: [m1-clerkenstein-backend/retro.md](releases/archive/01.00-body-double/m1-clerkenstein-backend/retro.md). The gate (alignment fidelity) is met; live injection into a running platform is rosetta-demo work (v1.1) / M2 (orgclient).
**Exit gate:** `/align-run` reports **100% alignment on the platform-consumed Clerk Go surface (critical capabilities) and ≥95% overall**, with any waived genes documented + justified in the divergence report.
**Iteration protocol:** `corpus/architecture/alignment_testing.md` (the M0-delivered alignment-measurement process) — the measure → fix-diverging-genes → re-measure loop.
**Why iterative (not section):** the deliverables are writable, but *which genes diverge and how costly each is to close* only emerges from measuring against the real Clerk — a fixed up-front checklist would be speculative. The score is the commitment; the path to it is open.
**Depends on:** M0 (its skills + DNA format + test class).
**Parallel with:** none (gates M1b and M2).
**Estimated complexity:** large
**Re-scope trigger:** if consecutive strategy iters (toks) can't close a diverging gene (e.g. a capability that's fundamentally unmockable offline), waive it with justification or escalate to the user — don't chase an unreachable 100%.
**Open questions:** which capabilities need live-Clerk record/replay vs pure local mint; the precise critical-capability set; stub just `authn`+`orgclient` or `replace` all of `colony` (`authn` is a package inside `colony`) — fallback is vendoring whole `colony`, as staging already does.
**KB dependencies:** `corpus/architecture/alignment_testing.md` (the iteration protocol), `corpus/services/clerk-integration.md`, `corpus/architecture/shared_libraries.md` (§ authn/colony), `corpus/ops/staging-clerk.md`, `corpus/ops/webhook_setup.md`.
**Delivers → `corpus/services/clerkenstein.md`** (the mirror design + injection mechanism — net-new) **+ the Clerk Alignment DNA** (`clerk@2.6.0` genome, authored via `/align-dna`).

### M1b: Clerk drift detection
**Status:** `done` (2026-06-03)
**Closes the gap after:** M1 (Clerkenstein is aligned at v2.6.0 — but must *stay* aligned as the platform bumps `clerk-sdk-go` / `@clerk/*`).

**Closed 2026-06-03** (2 sections + 1 harden pass). Automation/config over M0 — no new measurement machinery. In the clerkenstein repo: `scripts/gate.sh` (alignment gate, built-binary so exit 0 met / 2 regressed) + `scripts/drift-check.sh` (DNA-diff + gate, exit-code contract **0** none / **1** DNA moved / **2** gate regressed / **3** usage) + `.github/workflows/alignment.yml` (push + **weekly** CI) + `scripts/drift-test.sh` (9-assertion regression harness pinning the contract + the 2 build-phase fixes). Delivered the "Drift detection (M1b)" runbook in `corpus/services/clerkenstein.md`. Verified across all exit paths against a simulated `clerk@2.7.0` bump; shellcheck clean, flake 5/5. Close review: 0 findings.
**Goal:** Reuse M0 wholesale to make Clerk drift a flagged, mechanical event: on a version bump, `/align-dna` diffs the DNA (what changed) and `/align-run` re-scores the existing mirror against the new source (score drop = broken genes), CI-gated on "alignment ≥ threshold."
**Scope:**
  - In: the "bump pinned Clerk version → DNA-diff → re-score → report" workflow; the CI gate on alignment score; golden-capture refresh on bump.
  - Out: building the framework (M0); authoring the original mirror/DNA (M1); the JS surface (M2 owns its own genes).
**Depends on:** M1 (needs a built, aligned mirror + the Clerk DNA). Reuses M0's skills — **now automation/config over M0, not new machinery** (the right size for a B-milestone).
**Parallel with:** M2 (CI/automation vs JS code — disjoint surfaces).
**Acceleration effect:** every future Clerk bump becomes a flagged, scored update instead of a silent break — the brief's "follow platform updates within minutes" requirement, mechanized.

### M2: Clerkenstein — browser session + webhook coherence (JS)
**Status:** `done` (2026-06-03)
**Shape:** `section`
**Goal:** The frontend logs in with no real Clerk, and created/seeded users/orgs reach the DB without real Clerk webhooks.

**Closed 2026-06-03** (5 sections S1–S5 → 4 harden passes → close review → merged to `release/01.00-body-double`). Closes the last two Clerk seams so a demo stack is **Clerk-free end to end**. Delivered (in the gitignored `anthropos-demo/clerkenstein` repo): the **fake FAPI server** (`fapi/`) + the publishable-key codec — the browser logs in via a *minted publishable key* that encodes the fake FAPI host, **config-only, no SDK fork** (M2-D1 spike resolved the milestone's defining risk in the strong direction; the real-dev-Clerk fallback is documented but un-exercised); the **fake BAPI server** (`bapi/`) that disarms the platform's networked `orgclient` via an `api.clerk.com` DNS/base-URL redirect (the **M1-D2 Fate-3 pickup**), backed by the M1 orgclient twin made **concurrency-safe** (M2-D2); the **svix-signed webhook injector** (`webhook/`) for the 12 consumed event types → `POST /api/webhook/clerk`; and a **second Alignment DNA** (`clerk-js-5`, 9 genes, runner `cmd/jsfapirun`) scored at **100%/100%** like the Go side — proving the M0 framework is **surface-generic**. Both gates 100%/100% (Go 22/22 + JS 9/9); 112 Go test/fuzz funcs; flake 5/5; gofmt/vet/shellcheck clean. **Close review** found + fixed an `orgclient.ChangeRole` nil-map panic + phantom-membership divergence the alignment gate missed (reachable via the `bapi/` server) — M2-D4, with regression tests; plus a gofmt fix + the repo README refresh; 0 scope gaps, 0 deferrals (deferral audit GREEN). Decisions M2-D1…D4. Retro: [m2-browser-webhook-coherence/retro.md](releases/archive/01.00-body-double/m2-browser-webhook-coherence/retro.md). **This was the last *feature* milestone of v1.0**; a cleanup B-milestone **M2b (repo consolidation)** was inserted after it (2026-06-03) to tidy the `clerkenstein` repo before `/developer-kit:close-release`.
**Scope:**
  - In: a fake Clerk FAPI path for `@clerk/nextjs ^6.39.2` (next-web-app, ant-academy) and `@clerk/clerk-js ^5.52.3` (studio-desk) via publishable-key + base-URL/DNS override — **with the decided fallback**: keep the real dev Clerk app for the browser session while the backend stays fully mocked; a **webhook injector** feeding the existing `app/internal/clerk/events/` sync pipeline directly; **the JS surface's fidelity expressed as alignment genes via M0** where applicable (same score treatment as the Go side).
  - In (**routed from M1 close — M1-D2, Fate 3**): the **fake-Clerk-API-server** (HTTP interception of `api.clerk.com`) ALSO serves M1's **orgclient** injection — the Go `app/internal/clerk/orgclient` is app-internal + networked, so it can't `go.mod replace` like authn; it disarms via the same fake-API-server this milestone builds for the JS side. The Clerkenstein orgclient mirror behavior already exists + scores 100% (M1); M2 wires the HTTP redirect that makes the platform's real orgclient hit it.
  - Out: multi-instance stacks (M3); data seeding (M4).
**Depends on:** M1 (consumes the mock contract + minted-token shape). **Parallel with:** M1b (yes).
**Estimated complexity:** large — **highest technical risk in v1.0** (SDKs hard-code Clerk FAPI; no documented base-URL override).
**Open questions:** can `@clerk/*` be pointed at a fake FAPI without a fork? (the fallback exists because this is uncertain) — spike the override early.
**KB dependencies:** `corpus/architecture/alignment_testing.md`, `corpus/services/clerk-integration.md`, `corpus/architecture/frontend_architecture.md`, `corpus/services/next-web-app.md`, `corpus/ops/webhook_setup.md`.
**Delivers → `corpus/services/clerkenstein.md`:** extends the M1 doc with the JS path + webhook injection + the fallback decision.

### M2b: Clerkenstein repo consolidation + knowledge base
**Status:** `done` (completed 2026-06-03)
**Shape:** `section`
**Dir:** [m2b-clerkenstein-consolidation/](releases/archive/01.00-body-double/m2b-clerkenstein-consolidation/)

**Closed 2026-06-03** (5 sections S1–S5 → 1 harden pass → close review → merged to `release/01.00-body-double`). A pure-cleanup B-milestone that reorganized the `clerkenstein` repo (gitignored `anthropos-demo/clerkenstein`, its own git on `main`) into a clean, self-documented **library-named** structure — **no behavior change**, both alignment gates (Go 22/22, JS 9/9) + the drift harness (9/9) stayed green throughout. Delivered: the **library-named dirs** (`authn/` mocks colony/authn · `clerk-backend/` mocks clerk-sdk-go/v2 = the bapi server + orgclient store **merged** · `clerk-frontend/` mocks @clerk/clerk-js+nextjs · `clerk-webhook/` mocks svix) + `shared/` (the universal-key HS256 JWT, extracted because `clerk-frontend` **mints** and `authn` **verifies** the same token — `parse`→`shared.Parse` exported, M2b-D4) + `alignment/` (the M0-consumption harness: `cmd/{clerkrun,jsfapirun}` + `dna/` + `golden{,-js}/` + `scripts/`) via **69 history-preserving `git-mv` renames**; a self-contained **`knowledge/` base** (kb-index + scope + architecture + alignment + injection + coverage-index) + 6 per-library READMEs + slim root README; an `.agentspace/` (gitignored contents, dir preserved) + `.gitignore` baseline + asset hygiene; and `CLAUDE.md` + `singularity-manifest.md` (authored TO the `/singularity-kit:repo-consolidate` standard — the formal `repo-consolidate code` run is a **USER finalize**, M2b-D3/D8, since the skill is `disable-model-invocation`). Rosetta-side: slimmed `corpus/services/clerkenstein.md` 197→62 lines to a pointer at the repo's KB + fixed 2 stale refs in `alignment_testing.md`. **Close review** found + fixed 1 should-fix code-quality (a fuzz-test comment naming pre-reorg packages) + 2 doc findings (coverage-index count drift 112→113 after the harden test, state.md Headline refresh) — clerkenstein fixes on its own `main` (`ad87545`); 0 scope gaps, 0 deferrals (deferral audit **GREEN** — 2 inherited singles owned by close-release/M3, 0 repeat). Decisions M2b-D1…D8; D1/D2/D4 blended into the repo's own KB. Retro: [m2b-clerkenstein-consolidation/retro.md](releases/archive/01.00-body-double/m2b-clerkenstein-consolidation/retro.md). **This was the LAST milestone of v1.0** → next is `/developer-kit:close-release`.

**Goal:** The `clerkenstein` repo grew organically across M1/M1b/M2 into flat package dirs (`authn bapi orgclient fapi webhook cmd dna golden golden-js scripts`) with a single README and no knowledge base. M2b reorganizes it into a clean, self-documented **library-named** structure — one dir per mocked dependency + a shared dir + an alignment harness dir + a `knowledge/` base — following `/singularity-kit:repo-consolidate`, so the repo is navigable + operable by agents *before* v1.0 ships.
**Context (B-milestone — cleanup after M2):** pure reorg / docs / hygiene over the M2-complete repo. **No behavior change** — both alignment gates (Go 22/22, JS 9/9) and the drift harness stay green throughout; the move repoints imports + DNAs/goldens/runners/scripts, it does not alter the mocks. Class of work like M1b (tooling/cleanup over a shipped surface).
**Scope:**
  - In (**1 — Restructure**): one dir per mocked library/framework + one shared dir, **library-named** (user-chosen scheme): `authn/` (mocks `colony/authn`), `clerk-backend/` (mocks `clerk-sdk-go/v2` — the `bapi` server + the `orgclient` store **merged into one dir**), `clerk-frontend/` (mocks `@clerk/clerk-js` + `@clerk/nextjs` — the FAPI), `clerk-webhook/` (mocks `svix`); `shared/` (universal-key HS256 JWT + claims + canonical helpers — extracted because `clerk-frontend` mints and `authn` verifies with the same key); `alignment/` (the M0-consumption harness: `cmd/{clerkrun,jsfapirun}` + `dna/` + `golden{,-js}/` + `scripts/`). **Tests stay co-located within each library dir.** Go package identifiers can't contain hyphens → each hyphenated dir declares a clean package (e.g. `clerk-backend/` → `package clerkbackend`) — M2b-D1, confirmed at build.
  - In (**2 — Knowledge base**): a self-contained `knowledge/` dir documenting Clerkenstein — scope/goal; how it's built (the 4 mocks + shared); how fidelity is **validated with alignment tests against a pinned Clerk version** (the M0 framework + the two DNAs + the gate); **per-library injection recipes** (`go.mod replace` for `authn`; `api.clerk.com` HTTP/DNS redirect for `clerk-backend`; config-only publishable-key override for `clerk-frontend`; direct svix-signed POST for `clerk-webhook`); a coverage index. Per-library `README.md`s + a top-level index. Solid, well-written, well-distributed.
  - In (**3 — Hygiene**): an `.agentspace/` dir with contents **gitignored**; `.gitignore` cleanup (the current comment is mismatched); built-binary + transient hygiene per `repo-consolidate`'s asset-hygiene checks.
  - In (**4 — Consolidate**): run `/singularity-kit:repo-consolidate code` to standardize the repo (emit `CLAUDE.md` + `singularity-manifest.md`, audit against the code-repo + asset-hygiene standards, apply fixes), then re-verify both gates + the drift harness. **Note:** `repo-consolidate` is `disable-model-invocation` (user-invoked) — the build authors the structure TO its standard so the run is a clean finalize; the **user types the skill** (pointed at the `clerkenstein` repo).
  - Out: new library support / new alignment genes (the `@clerk/express` coverage gap — **now picked up by M2c**); any live injection wiring into a running platform (still v1.1/M3); any change to rosetta's M0 framework or to the platform repos.
**Depends on:** M2 (consolidates the M2-complete repo). **Parallel with:** none (touches the whole repo). **Precedes:** `/developer-kit:close-release`.
**Estimated complexity:** medium — mechanical but wide (touches every package + the gate/drift scripts); the only real risk is import/script repointing, fully caught by the **green-gate invariant** (gates + drift re-run after each section).
**KB dependencies:** `corpus/services/clerkenstein.md`, `corpus/architecture/alignment_testing.md`; the `/singularity-kit:repo-consolidate` standards (base + code-repo + asset-hygiene).
**Delivers → the `clerkenstein` repo's own `knowledge/` base** (net-new, self-contained) **+ slims `corpus/services/clerkenstein.md`** (rosetta) to a pointer at the repo's `knowledge/` + the new structure.

### M2c: Clerkenstein — `@clerk/express` backend session verification (RS256/JWKS)
**Status:** `done` (2026-06-03, **closed-on-gate**)
**Shape:** `iterative` (alignment-score gate, like M1) — a **feature** milestone; the letter suffix marks *insertion after M2b*, not a B/tooling milestone.

**Closed 2026-06-03** (5 iters: bootstrap TOK-01 → DNA → RS256 foundation → **crux proof** → full runner → gate; 1 final harden pass). Brought the **last un-gated Clerk consumer — `@clerk/express`** (studio-desk's Node backend) under the alignment framework at **100%/100%** (3rd DNA `clerk-express-1.json`, 9 genes). The **RS256 wall fell to an additive path** (M2c-D1/D2): an RSA keypair + a real JWKS + RS256 minting that the *genuine* `@clerk/backend` accepts networkless via `jwtKey` — **no HS256 migration**, so M1 (22/22) + M2 (9/9) stayed green. `@clerk/express` is **verified, not reimplemented** (no mock dir — the svix discipline; M2c-D5); the `expressrun` runner mints tokens (Go) + drives the real SDK (embedded `verify.js`, Node). The `clerkClient` BAPI reads were already covered by `clerk-backend` (M2c-D4). Close: folded the surface into the knowledge base + corpus, fixed a gitignore gap + 1 adversarial flake (`tamperSig`); deferral audit GREEN; the express-gate CI-wiring (needs Node) routed to v1.1. 128 test/fuzz funcs / 8 packages; all four gates green. Retro: [m2c-clerk-express-alignment/retro.md](releases/archive/01.00-body-double/m2c-clerk-express-alignment/retro.md).
**Dir:** [m2c-clerk-express-alignment/](releases/archive/01.00-body-double/m2c-clerk-express-alignment/)
**Goal:** Bring the **last un-gated Clerk consumer — `@clerk/express`** (studio-desk's Node backend auth) under the alignment framework: a new **`clerk-express/`** seam + a **3rd Alignment DNA**, driven to a gate, so studio-desk's backend genuinely verifies Clerkenstein tokens (not via its `MOCK_CLERK=true` bypass). Completes v1.0's thesis — *no* Clerk seam left un-faithful before shipping.
**Why iterative + the defining unknown (the RS256 wall):** `@clerk/express` (via `@clerk/backend`) verifies **RS256 via JWKS only** and **hard-rejects HS256** (`assertHeaderAlgorithm` → `TokenInvalidAlgorithm`). Clerkenstein mints HS256 universal-key tokens + serves an **empty JWKS**, so an HS256 shim is a dead end. The milestone must add an **RS256 path** (RSA keypair + a real JWKS from the fake FAPI + RS256 minting + the real-`@clerk/express` verifier). **The central iteration question:** can RS256 be **additive/parallel**, or must the existing HS256 seams (`authn`/`clerk-frontend`/`shared`) **migrate to RS256** — re-capturing the Go DNA goldens + re-gating M1/M2? The gate-driven iterations resolve it.
**Scope:**
  - In: a new **`clerk-express/`** seam (library-named); an **RSA keypair + a real (non-empty) JWKS** served by the fake FAPI (`clerk-frontend`'s `/.well-known/jwks.json`); **RS256 token minting**; the `@clerk/express` **DNA** (`clerk-express-1.json`, source `@clerk/express ^1.3.47`); a runner that drives **the real `@clerk/express` SDK** against the mock (the svix-pattern — verify against the genuine library); the **alignment gate** as the exit criterion.
  - In (confirm, don't rebuild): `@clerk/express` also calls `clerkClient.{getOrganizationMembershipList, getOrganization}` — those are **BAPI**, already 100%-mocked by `clerk-backend/`; M2c adds *integration* genes confirming that path, not a new BAPI mock.
  - Out: changing studio-desk or any platform repo (the `MOCK_CLERK` bypass is the platform's own); a webhook (svix) DNA (separate future gap); live injection into a running studio-desk (rosetta-demo work, v1.1).
**Candidate genes (~8, `clerk-express-1.json`):** `ExpressAuth/{valid, expired, malformed, bad-signature, no-token}` (error_class) · `ExtractIdentity/universal-user` (exact: verified claims → `req.auth`) · `JWKS/non-empty-rsa` (shape) · `ClerkClientBAPI/{org-membership-list, get-organization}` (integration vs `clerk-backend`).
**Exit gate:** alignment **≥ 95% overall / 100% critical** on `clerk-express-1.json`, AND the load-bearing test passes (a **real `@clerk/express` instance accepts a Clerkenstein-minted token + extracts the right identity**).
**Depends on:** M2 (the FAPI + token machinery it extends) + M2b (the consolidated repo it adds a seam to). **Precedes:** `/developer-kit:close-release`.
**Estimated complexity:** large — **highest fidelity-risk in v1.0**: the RS256 path may force a token-algorithm migration of the existing 100%-gated seams.
**KB dependencies:** `corpus/architecture/alignment_testing.md`; the clerkenstein repo's own `knowledge/` (alignment / architecture / injection / sources); the `@clerk/express` + `@clerk/backend` source under `anthropos-dev/studio-desk/node_modules`.
**Delivered → the clerkenstein repo's `knowledge/`** (alignment/architecture/sources updates) **+ a 3rd DNA + the `expressrun` runner;** updated `corpus/services/clerkenstein.md`'s scorecard to a **3rd *measured surface*** (`@clerk/express`, **verified-not-mocked** — no new mock dir, per M2c-D5; the genuine SDK is *satisfied* via an additive RS256/JWKS path).

### Execution graph

```
v1.0 "body double"   — a stand-in the platform can't tell apart, and we can prove it

  M0 (alignment framework: /align-dna + /align-run, test class, DNA format, golden capture, toy ref)
    │
    ↓
  M1 (Clerkenstein backend mirror — ITERATIVE: author Clerk DNA → drive alignment score to gate)
    │
    ├──→ M1b (Clerk drift detection — DNA-diff + re-score, CI-gated across version bumps)   ∥ M2
    └──→ M2 (browser session + webhook; reuses the alignment class for the JS surface)
              │  (both closed — repo feature-complete)
              ↓
    M2b (repo consolidation — library-named dirs + self-contained knowledge base; gates stay green)
              │
              ↓
    M2c (ITERATIVE: @clerk/express RS256/JWKS — new clerk-express/ seam + 3rd DNA → alignment gate)
              │
              ↓
    /developer-kit:close-release → v1.0 ships to main
```

### Parallelism

- **M0 → M1 → {M1b, M2}** sequential at the core: M1 needs M0's framework; M1b + M2 need M1's mirror/contract.
- **M1b ∥ M2:** disjoint surfaces — M1b is CI/automation over M0; M2 is JS + the webhook injector. Merge risk **low**.
- **M3 ∥ M2 (cross-version, yes-with-caveats):** sequenced cleanly by the version boundary (M3 starts after v1.0 closes).

### Risks (v1.0)

| Risk / decision | Severity | Mitigation |
|---|---|---|
| **Source is a live SaaS** — Clerk's API capabilities can't be hit freely/offline/deterministically | blocks-release (reproducibility) | M0 **record/replay golden captures** is a core requirement, not an afterthought — capture once, replay forever |
| **DNA completeness gaming** — 100% on a thin DNA is hollow | degrades-quality | `/align-dna` capability-coverage check (every platform-consumed endpoint present) + M1b version-bump DNA-diff keeps it complete |
| **Defining "equivalent"** — timestamps, generated IDs, error formats differ even when behavior matches | degrades-quality | M0 ships **equivalence operators** (exact / same-shape / normalized / same-error-class) chosen per gene |
| **JS/FAPI fake server** — SDKs hard-code Clerk FAPI, no base-URL override | blocks-release (full no-Clerk browser) | **Decided fallback:** real dev Clerk app for the browser, backend fully mocked; spike override early in M2 |
| **`colony` replace granularity** — `authn` is a package inside `colony`, not its own module | degrades-quality (M1 effort) | M1 early iter resolves it; fallback = vendor whole `colony` (staging precedent) |
| **Repo layout** — where the framework vs the Clerk mirror live | nice-to-resolve | **Decided:** the M0 framework (skills + format + doc) lives in rosetta; the Clerk DNA + alignment tests + mirror live in the `clerkenstein` repo, cloned into gitignored `anthropos-demo/` |
| **"Zero platform-code changes" interpretation** — `replace` edits the *clone's* go.mod | nice-to-resolve | build-time injection in the gitignored clone + skip-worktree; upstream repo never modified (same as staging's `vendor-colony/`) |

### Branch model

`release/01.00-body-double` (cut from `feat/demo-environment` at M0). Milestone branches:
`m0/alignment-framework`, `m1/clerkenstein-backend`, `m1b/clerk-drift-detection`,
`m2/browser-webhook-coherence`, `m2b/clerkenstein-consolidation`, `m2c/clerk-express-alignment`.
**M1 + M2c are iterative** → built by `/developer-kit:build-mstone-iters` (close on a Gate Outcome Ledger).
M0/M1b/M2/M2b are section → `/developer-kit:build-milestone`. All → `/developer-kit:close-milestone` →
`/developer-kit:close-release`.
The `clerkenstein` repo's own code commits stack on its `main` (its own gitignored git, no branch model);
the rosetta-side milestone records + corpus pointer land on the `m{N}/…` branch.

### Out of scope (v1.0 — recorded for v1.1+)
- Multi-instance disposable stacks, data seeding, use-case recipes → all v1.1 "show floor".
- Mirroring engines other than Clerk with M0 (the framework is generic, but v1.0 only exercises it on Clerk).
- AI-generated demo content (transcripts/embeddings) → v1.1 stretch or deferred.

## Shipped releases

- **v1.0 "body double"** — shipped **2026-06-03**, tag `v1.0`. The alignment-testing framework + Clerkenstein
  (100%/100% on Go · JS/FAPI · `@clerk/express`). Detail in the `## Done` section above; records archived at
  [`releases/archive/01.00-body-double/`](releases/archive/01.00-body-double/).

## Notes

- Milestone numbering is **flat sequential** (M0, M1, M2, …); a letter suffix has two uses: (1) a milestone **inserted after** the fact — `b` for tooling/cleanup (M1b drift CI, M2b consolidation), and the letter-suffixed *feature* milestone M2c (iterative, "inserted after M2b"); and (2) a **split** of one planned milestone into a sequential mini-arc — **M7a → M7b → M7c** is the single former M7 "seeding" split into framework+safety / data-DNA / fleet (2026-06-04, M7a-D1). Both reuse the letter suffix; context disambiguates. See [`context.md`](context.md).
- v1.0 mixes shapes: M0/M1b/M2/M2b are **section**; **M1 + M2c are iterative** (alignment-score gates).
- v1.1 "show floor" mixes shapes too: M3–M6 + M7a/M7b + M8 are **section**; **M7c is iterative** (data-DNA coverage gate).
- v1.2 "set dressing" is **all `section`** (M9a/M9b/M10/M11) — the snapshot surfaces are decomposable up front (the
  framework + 2 known surfaces + the product layer); the fidelity gate is a per-surface acceptance check, not an
  emergent-path iterative gate. (AI-content, the iterative-shaped candidate, was held to v1.3.) The former M9 split
  into **M9a (framework) + M9b (taxonomy surface)** on the 2026-06-06 refinement — the M7a→M7c precedent.
