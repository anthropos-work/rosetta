---
title: "KB Fidelity Audit — M236 prove-on-billion"
date: 2026-07-20
scope: milestone:M236
invoked-by: user (on-demand, pre-flight for build-mstone-iters Phase 0b)
---

## Verdict

**RED** — blocked. Five independent blocker-severity classes:

1. **A falsified load-bearing claim inside M236's own `overview.md`** (a page-object that does not exist) — F1.
2. **An exit-gate clause the corpus documents as false by construction** ("reachable only over the tailnet"
   vs `safety.md:405`'s "every port on `0.0.0.0`, flag or no flag" — and `tailscale-serve.md:626-627`
   says the obvious mitigation does not work on Linux) — F2.
3. **Half the exit gate is unmeasurable with today's harness** (p95 click→ACCESS for a content seat
   throws before t0) — F3.
4. **Two inherited carry-forward tasks are misdescribed** — the "2 drifted manifests" are *not* drifted
   (F20), and the "rendered-card count" descriptor M236 is told to *run* does not exist and must be
   *authored* (F21).
5. **The declared `iteration_protocol_ref` contains no iteration protocol** — F5.

Plus a systemic blind area: **all six "method" docs — including BOTH of M236's declared
`iteration_protocol_ref` docs — contain zero mentions of the entire v2.5 content-stories feature.**

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Remote tailnet deploy (billion) | `corpus/ops/demo/tailscale-serve.md` | `stack-injection/gen_tailscale_serve.py`, `demo-stack/up-injected.sh`, `demo-stack/tailscale_autohost.py` | PAIRED (stale) |
| Bring-up knobs/flags contract | `corpus/ops/demo/demo-up-defaults.md` | `demo-stack/up-injected.sh`, `demo-stack/ensure-clones.sh`, `stack-core/demo_knob_guard.py` | PAIRED (badly stale) |
| Auto-verify net | `corpus/ops/verification.md` | `stack-verify/live/autoverify.sh`, `live/verify.sh`, `lib/services.sh` | PAIRED (stale) |
| Latency gate (p95 ACCESS) | `corpus/ops/demo/latency-budget.md` | `stack-verify/e2e/run-latency.sh`, `e2e/lib/latency.ts` | PAIRED (stale on the decisive claim) |
| Coverage sweep (presence proof) | `corpus/ops/demo/coverage-protocol.md` | `stack-verify/e2e/lib/coverage-manifest.ts`, `crawl.ts` | PAIRED (actively contradicts M236) |
| Playthroughs (function proof) | `corpus/ops/demo/playthroughs.md` | `playthroughs/{manifest,e2e,seed,report}/` | PAIRED (aligned) |
| M231 route map | `corpus/ops/demo/content-stories-routes.md` | platform clones (read-only) | PAIRED (healthiest) |
| M232 session-clone seeder | `corpus/ops/demo/session-clone-spec.md` | `stack-seeding/seeders/content_stories*.go`, `scrub/`, `cmd/content-capture/` | PAIRED (stale counts) |
| M233 manifest projection | `corpus/ops/demo/content-stories-spec.md` | `seeders/content_manifest.go`, `cmd/stackseed --content-export` | PAIRED (stale/self-contradictory) |
| M234 cockpit tab | `corpus/ops/demo/cockpit-spec.md` | `demo-stack/cockpit.py` | PAIRED (aligned) |
| Academy fill (Thread A) | `corpus/ops/demo/frontend-tier.md` (mechanism), `corpus/services/ant-academy.md` | `demo-stack/patches/academy-fs-published-fallback`, `app/cmd/academy-seed` | PAIRED (partial) |
| Exposure posture | `corpus/ops/safety.md` Part 3 / §3.8 | `stack-core/gen_injected_override.py` | PAIRED (contradicts gate) |
| Demo-patch mechanism | `corpus/ops/demo/demopatch-spec.md` | `demo-stack/patches/*`, `stack-injection/apply-academy-fs-published.sh` | PAIRED |
| **Programmatic content-seat login drive** | **— (none)** | `stack-verify/e2e/lib/cockpit-login.ts`, `playthroughs/e2e/lib/hero-login.ts` | **BLIND-AREA** |
| **Dynamic-URL page coverage mechanism** | **— (protocol mandates the opposite)** | `crawl.ts:122,224` (`skipPaths`) | **BLIND-AREA** |

## Fidelity Findings

### F1 (BLOCKER) — `AISimulationResultContainer` page-object does not exist
- **Source:** `m236-prove-on-billion/overview.md:37`; repeated `m235-prove-it-lands/carry-forward.md:37`, `m235.../decisions.md:146`.
- **Expected:** "reuses the shared `AISimulationResultContainer` page-object".
- **Actual:** zero hits in rext. The name is a **next-web platform React component** (`content-stories-routes.md:52`). The nearest harness object, `playthroughs/e2e/lib/simulation-page.ts:9-10`, explicitly asserts only at the **LAUNCH boundary** — no `/result/` locator, no `sessionId` handling (routes only `:20`, `:36`).
- **Verdict:** STALE (falsified). **Fix owner:** update the milestone doc — M236 must AUTHOR this page-object; budget as build, not reuse.

### F2 (BLOCKER) — the exit gate's "reachable only over the tailnet" contradicts the corpus
- **Source:** `overview.md:10` / `:23`.
- **Actual:** `safety.md:405` — "Every demo container's offset port is published on `0.0.0.0` — ALL interfaces — on EVERY `demo-up`, with or without the flag"; `:413` measured `DEMO: 14 ports → 0.0.0.0`; `:424` "no `127.0.0.1` prefix anywhere, in either family". And `tailscale-serve.md:626-627` — "**On Linux this bypasses your host firewall.** Docker installs its own iptables rules in the `DOCKER` chain… A `ufw deny` on the port does **not** block it."
- **Verdict:** STALE/UNPROVABLE as worded. billion is a Linux VM, so binding-level and ufw-level tailnet scoping are both ruled out by the corpus itself.
- **Fix owner:** reword the gate to the honest posture (tailnet-scoped by **network reachability / cloud security-group**, not by binding), or add explicit host-perimeter work to scope. Do not let M236 claim a gate the corpus says is false. Note the one measured exception: `ant-academy.sh:330` passes `-H 127.0.0.1` on the localhost path (`safety.md:443`).

### F3 (BLOCKER) — p95 click→ACCESS is unmeasurable for a content seat
- **Source:** `latency-budget.md:32-39` — "a vantage is a seat-key + a landing origin, not a new code path. Add a vantage by adding its `case` to `run-latency.sh`."
- **Actual, blocker A:** the hero CTA carries `data-login-as` (`cockpit.py:618`); the content-story player CTA does **not** (`cockpit.py:421-425`). `latency.ts:123-127` **throws** when it is missing — and `heroName` is the entire second half of the ACCESS predicate (`latency.ts:138-155`, `:273-274`), so there is no skip-the-attribute fix.
- **Actual, blocker B:** `run-latency.sh:42-47` hard-rejects any vantage outside `employee|manager|recruiter` with `exit 2`; `latency.spec.ts:86` writes `latency-out/${VANTAGE}.json`, so masquerading silently **overwrites** the hero result.
- **Verdict:** STALE — the "just add a case" generalisation does not hold.
- **Fix shape (not applied):** add `data-login-as` at `cockpit.py:423`, add a `content` case to `run-latency.sh:42-47`, include the identity in the output filename.

### F4 (BLOCKER) — `coverage-protocol.md` mandates EXCLUDING the pages M236 must prove
- **Source:** `coverage-protocol.md:421-431` — a `skipPaths` rule `/\/result\/[0-9a-f-]{8,}/`; reinforced `:443-445` (runtime-computed result keyed by sessionId → crawl-scope/exclude).
- **Actual:** the seam is real (`crawl.ts:122`, `:224`). The exclusion was adopted because seeded sessions had no computed evaluation (M42e iter-04). **M232 session-cloning changes that premise; the doc was never updated.**
- **Verdict:** STALE + BLIND-AREA. M236 must consciously **reverse a documented protocol rule** — a spec decision, not an incidental edit. Make it an explicit deliverable with rationale.

### F5 (BLOCKER) — `iteration_protocol_ref` is hollow
- **Source:** `overview.md:11` declares `corpus/ops/verification.md + corpus/ops/demo/tailscale-serve.md`.
- **Actual:** `verification.md` (335 lines) contains **no** measure→triage→fix loop, no gate, no exit condition, no per-iteration artifact. Structure: `:1-56` what the net is · `:58-92` offset/scope model · `:94-228` probe enumeration · `:229-283` `autoverify.json` semantics · `:285-325` stale-verdict essay · `:327-335` xrefs. The only loop-shaped text is operator advice (`:169`, `:173-174`).
- **Verdict:** MISDECLARED. It supplies a gate **INPUT** (a fresh green `autoverify.json`), not a protocol. `/developer-kit:build-mstone-iters` expects a declared tik → re-measure → close-iter loop.
- **Fix owner:** author a protocol or re-point the ref before Phase 0b.

### F6 (BLOCKER) — `demo-up-defaults.md:146` says the host pre-flight is non-fatal; it exits 1
- **Expected:** `demo-up-defaults.md:146` — `DEMO_NO_HOST_PREFLIGHT` `0` → "the host pre-flight runs (**non-fatal — warns, never blocks**)".
- **Actual:** `up-injected.sh:304` `preflight_host_prereqs()`; on missing Go (`:311`), atlas (`:317`), or an un-set tailscale operator (`:336-340`) → `up-injected.sh:345-347` **`exit 1`**. `tailscale-serve.md:135-137` ("fails loud… aborts") is the ALIGNED one; the two docs contradict each other.
- **Verdict:** STALE / factually wrong. The row conflates the genuinely non-fatal RAM/disk pre-flights (`:264`, `:287`) with the fatal M215 toolchain pre-flight. **Highest-risk single row for a cold billion bring-up.**

### F7 — `demo-up-defaults.md`: 24 of 28 line anchors stale; guard does not fence them
- The doc's premise is "derived from the parsers and fenced against them in BOTH directions". The guard **exists** and I ran it: it **reports 1 disagreement** (`DEMO_NO_ACADEMY_FILL` undiscoverable) and prints "27 env knob(s) + 10 cli flag(s)".
- But `demo_knob_guard.py:55` (`_KNOB_RE`) matches only `${NAME:-default}` and `:61` (`_DOC_ROW_RE`) matches only the first table cell — **no line-anchor validation**. `up-injected.sh` is now 2120 lines.
- Sample drift: `DEMO_NO_UI` doc `:114` → actual `:187` · `DEMO_NO_VERIFY` doc `:1594` → actual `:2110` · `DEMO_NO_COCKPIT` doc `:1483` → actual `:1947` · `DEMO_VM_MIN_GIB` doc `:181` → actual `:254` · `DEMO_DISK_MIN_GIB` doc `:204` → actual `:277` · `DEMO_NO_PATCH` doc `:469` → actual `:441` · `DEMO_NO_AUTHZ_SKIP` doc `:792` → actual `:1225`.
- Only 4 anchors are correct, all in the Remote-access block (most recently touched at M220 S3): `STACK_PUBLIC_HOST:41`, `resolve_public_host:106`, `DEMO_NO_PUBLIC_HOST:35`, `ensure-clones.sh:85`.
- **Verdict:** STALE (systemic). Do not trust this doc's anchors when choosing bring-up flags; read `up-injected.sh` directly.

### F8 — counts: knob table, session fixture, and manifest
- `demo-up-defaults.md:37` says "all **25**" env knobs; table has **26** rows; parsers expose **27**. STALE three ways.
- `DEMO_NO_ACADEMY_FILL` (`ant-academy.sh:368`, added M230) has **no row** — the guard's live failure. It gates the `academy-fs-published-fallback` patch that makes the grid render real cards, i.e. Thread A.
- `session-clone-spec.md:66` — "**9 real sessions** … assessment set **2 voice + 2 code + 1 document**". Actual fixture pins **13** sessions, assessment set **3 voice + 2 code + 2 document**. `:91`'s "8/9 fixtures" inherits the wrong denominator (now 13).
- `content-stories-spec.md` states the product/session counts **nowhere**; canonical `presets/content-manifest.json` is **4 products / 18 sessions** (simulation 13 · skill-path-legacy 2 · ai-labs 2 · skill-path-new 1).
- M235 `carry-forward.md` says "14 live Playthroughs"; actual is **15 live / 1 TODO** (`playthroughs.md:105` is correct — M225's `pt-hiring-recruiter-compare` was omitted from the carry-forward).

### F9 — `content-stories-spec.md` self-contradicts on M236's own status
- `:240-241` — "Today's fixture carries no academy session; this path… lights up when M235 adds the fixture." M235 landed it (1 academy session in the canonical manifest).
- `:266-267` — "Proving every CTA lands… is **M235**; proving it on `billion` is **M236**", contradicted by `:194-195` in the same doc.
- `:61-87` schema block omits the `Label` field that M235 added (`content_manifest.go:79-82`); prose mentions it at `:192` but the copyable block does not.
- **Verdict:** STALE. An M236 reader gets two different answers from one doc.

### F10 — seat keys are not 0-indexed
- `overview.md:36` and the carry-forward imply `content-player-<idx>` from 0. Actual canonical seats are **`content-player-23` … `content-player-35`** (13 unique), allocated after the 5 heroes (`roster.go:162,181-196`; `roster_test.go:50,:243`). An unknown key returns **400** (`clerkenstein/clerk-frontend/multiidentity_test.go:87`). A spec written against `content-player-0` fails.

### F11 — auto-verify does not probe the two surfaces M236 lands on
- `verification.md:133-152` implies full coverage. `stack-verify/lib/services.sh:35-62` has **no `ant-academy` row and no `hiring` (:3001) row**, and `up-injected.sh:2116` scopes in only `next-web-app studio-desk` — yet `gen_tailscale_serve.py:44,:50` fronts both.
- `verification.md:22-25` still says the frontend tier is "out of scope — the frontends don't exist in the stack yet; M19 adds them" (an M18-era statement; `services.sh:55-56,61` now carries them).
- The doc omits a fifth M217-era cheap-win the code has: `autoverify.sh:202-207` `buildfail.log` (STALE image detector).

### F12 — the doc's remote latency recipe is wrong for M236's exact scenario
- `latency-budget.md:153-159` omits `LATENCY_SCHEME=https`. `run-latency.sh:62` defaults `SCHEME=http`, and `:22-24` states a `--public-host` demo is HTTPS-fronted so "a hardcoded `http://` gets a 400/redirect from that vantage." Copy-pasting the doc's block for a remote demo fails at `readCockpitCta`.

### F13 — the "14 pre-existing demo-stack fails" carry is undocumented in the corpus
- `verification.md` has zero mentions and implies the opposite (`:163` "A clean run ends… OK — verified-working"; `autoverify.sh:311-315` sets `green: warnings == 0`; `run-latency.sh:86-91` refuses to measure on non-green).
- The 14 live only in the plan KB (`m235.../retro.md:60`, `hardening-ledger.md:47`) and are **pytest** failures, not runtime probes — so they do **not** dirty `autoverify.json` and will **not** block M236's latency gate. Nothing in the corpus tells a reader that.

### F14 — RAM/disk prereqs are silent and macOS-framed; billion is a 7 GB Linux VM
- `tailscale-serve.md:121-128` Step-0 prereq table has **no RAM row and no disk row**.
- Code: `up-injected.sh:254` `DEMO_VM_MIN_GIB:-12`, `:277` `DEMO_DISK_MIN_GIB:-20` — both **non-fatal warns** (`:261-264`, `:284-287`), RAM check skipped entirely when `NO_UI=1` (`:253`).
- **Disk 40 GB vs 20 GB floor → passes.** **RAM 7 GB vs 12 GB floor → warns, does not block**; the full UI tier still attempts to build.
- The remediation text is **macOS-only**: `up-injected.sh:263` prints "Raise Docker Desktop's VM to 12 GB (Settings → Resources)" — meaningless on a native Linux VM where `docker info MemTotal` is host RAM. `tailscale-serve.md:148-151` inherits the framing, and its reassurance ("M215 run held at 5.7 GB free") describes **steady state**, not the ~3.7 GB build spike.

### F15 — cockpit `:7700` missing from the port enumeration/diagram
- `tailscale-serve.md:483-497` (prose) is ALIGNED with `gen_tailscale_serve.py:66-67`, but the enumeration and ASCII diagram at `:419-431` omit `:7700` — the one place a reader looks for "which ports must be served", and the cockpit is precisely the surface M236 drives remotely.

### F16 — `demo-up-defaults.md:24` understates the flag surface
- Claims `<N>` + `--public-host` is "the entire flag surface"; `up-injected.sh:36-40` parses **three** arms including `--no-public-host` (`:38`). The doc's own table at `:158` documents it — `:24` predates it.

### F17 — path-segment drift in two anchors (line numbers correct)
- `content-stories-routes.md:95` cites `local_jobsimulation_session.go:52` as under `ent/`; actual is `ent/schema/` (101 lines, `field.Float32("score")` at `:52`).
- `content-stories-routes.md:275` cites `lab_session.go:122-127` under `ent/`; actual is `ent/schema/` (139 lines, `grade_result` at `:122-127`).
- Five `next-web-app` anchors were reported UNVERIFIABLE from the rext demo-1 clones path, but a usable checkout **does** exist at `stack-demo/next-web-app` — so `:111` `AISimulationResultContainer.tsx:499-506` (load-bearing: the interview-flag demopatch pin), `:119`, `:128`, `:137`, `:284` are verifiable locally and should be checked before M236 codes against them, then re-verified on the VM.
- `safety.md:443` cites `demo-stack/ant-academy.sh:330`, which is the **comment header**; the code the sentence asserts (`bind_args=(-H 127.0.0.1); [ -n "${STACK_PUBLIC_HOST:-}" ] && bind_args=(-H 0.0.0.0)`) is at **`:345`** (consumed `:385`). SOFT-STALE, off by 15.

### F18 — one dead file citation
- `coverage-protocol.md:603` cites `stack-verify/e2e/tests/probe-aireadiness-deeplink.spec.ts`, deleted in rext `66a021e`.

### F19 — `session-clone-spec.md` "distinct owner" is a precondition, not an invariant
- `content_stories.go:117-126` filters hero/non-member slots correctly, but assignment is `owners[idx%len(owners)]` (`:81`) — distinctness holds only while `idx < len(owners)`. True on the canonical preset (13 consecutive seats, no wrap). Flag if the fixture grows or the org shrinks.

### F20 (BLOCKER, corrects an inherited carry-forward) — the "2 drifted demopatch manifests" diagnosis is WRONG
- **Source:** `m230-academy-demo-fill/carry-forward.md:29-33` and `m235-prove-it-lands/carry-forward.md:70-71` — "the local `next-web` clone has DRIFTED from 2 pinned demopatch manifests (`next-web-public-website-url` + `next-web-studio-url`), which would drift-refuse on a cold bring-up." M236's `In:` list inherits this as a "re-anchor" task.
- **Actual**, measured against `stack-demo/next-web-app/packages/core-js/src/constants/urls.ts`: **neither manifest is drifted; the pins are correct.** Clone HEAD (pristine) sha is `0d4c3790…`, an exact match to studio-url's `pre_sha256`, with anchor count 1. The **working tree** is simply left fully patched from an un-reverted prior bring-up (sha `d92fa701…` = the end-of-chain `post_sha256`; anchor count 0; replacement + post_marker present) — the `RETURN` LIFO revert trap did not fire.
- Under G2's coherence probe (`apply_patch.py:88-101`) that state classifies **`ALREADY_PATCHED` → G4 idempotent no-op, exit 0**. It would **not** drift-refuse.
- **Verdict:** STALE (misdiagnosis). **The remediation is one `git checkout -- packages/core-js/src/constants/urls.ts`, not a re-anchor/re-pin.** Measured on this box; billion's clone may differ and must be checked there.

### F21 (BLOCKER) — the ANT_ACADEMY "rendered-card count" descriptor does not exist
- **Source:** `m230.../carry-forward.md:19` and `m235.../carry-forward.md:70` both specify the ANT_ACADEMY descriptor measures a **RENDERED-CARD COUNT**; M236's `In:` inherits it.
- **Actual:** the shipped descriptor is `ANT_ACADEMY_HOME_SECTION` at `stack-verify/e2e/lib/coverage-manifest.ts:709-713` — `realContent: { kind: 'text', mustInclude: ['AI Academy'], minMeaningfulLen: 400 }`. That is a **text-marker + length floor**, not a card count, and would pass on a grid rendering zero cards next to enough chrome text.
- **Verdict:** STALE/GAP. M236 must **author** the card-count assert; no doc discloses that it is missing. This is the actual Thread-A gate mechanism.

### F22 — `safety.md:424` states a false absolute
- **Expected:** `:424` — "There is no `127.0.0.1` prefix anywhere, in either family."
- **Actual:** `stack-injection/gen_injected_override.py:577` emits `ports: ["127.0.0.1:{5401 + offset}:443"]`, and its own comment (`:572-576`) calls it "the **FIRST loopback-bound published port in the demo**, and deliberately so." `safety.md:665` (§3.6) already documents it — the doc contradicts itself.
- **Verdict:** STALE. Needs an "except the Clerkenstein fake BAPI" carve-out. (The broader `:405` claim — every *container* port on `0.0.0.0`, every bring-up — is **verified still true** after v2.5 via `exposure_claim_guard.py:190-199`.)

### F23 — `demopatch-spec.md` inventory is stale
- `§5:149` claims **11** manifests; **14** are on disk. Missing from the §5 table: `academy-fs-published-fallback` (M230), `next-web-interview-flag-container` (M232), `next-web-interview-flag-result` (M232) — i.e. all three net-new v2.5 patches.
- `§4:96` says "the **eight** `next-web-app` patches"; actual is **10**.
- `§4`'s vehicle table row 3 lists only `apply-ant-academy-dev-origins.sh`; `apply-academy-fs-published.sh` is a second patch on that same vehicle and is unlisted.
- The **7 guards and the chain rule are ALIGNED** (G7 verified at `apply_patch.py:249-280`; chain `next-web-public-website-url.yaml:35 pre == next-web-studio-url post = fe15aa71…`; `manifest_loader.py:33-34` `REQUIRED` is exactly 10 keys). Code-side nit: `demopatch:25`'s docstring still states the OLD G2, contradicting its own implementation at `:225-245` — the spec is right, the comment is stale.

### F24 — `ant-academy.md` describes a pre-M230 world
- `:149`, `:157`, `:209-210` all describe M230 in **future/intent tense** ("deliberately chose", "is a demo-tooling concern", "M230 fills it"). M230 **shipped** (rext `76ee1a0`). None of the shipped artifact names appear anywhere in the doc: `academy-fs-published-fallback`, `ACADEMY_DEMO_FS_PUBLISHED`, `DEMO_NO_ACADEMY_FILL`, `apply-academy-fs-published.sh` — they exist in the corpus **only** in `frontend-tier.md:374-406`.
- Consequence: `:12` still tells a PM "*this is exactly why the academy looks blank in a demo*" — **no longer true by default**, since the fill is default-on.
- The progress-bearing chapter route M236 must calibrate to (`overview.md:48`) has **zero** supporting text: `:33` mentions `/chapters/<slug>/` generically, `:164` lists `/library*` as public, and neither is connected to progress rendering.
- **Verdict:** STALE. `frontend-tier.md` is the healthy mechanism doc; `ant-academy.md` has not been re-synced.

## Completeness Gaps

### G1 (critical) — the programmatic content-seat login drive is undocumented
No doc describes how the content tab's CTA is driven by a harness. `content-stories-spec.md:196-198` **names** the blocker but specifies no mechanism — a ticket, not a spec. The mechanism exists in code and is undocumented for this surface:
- `stack-verify/e2e/lib/cockpit-login.ts:56-60` POSTs `/v1/demo/select` `{key}` (handler `clerkenstein/clerk-frontend/server.go:170-173`), deliberately decoupled from cockpit HTML.
- `GET /content-manifest.json` is served by the cockpit (`cockpit.py:729-736`) but documented only as a **download** (`cockpit-spec.md:286`), never as a **harness input**.
- The academy row's CTA **bypasses the FAPI entirely** (`cockpit.py:102-107`, `_ACADEMY_JS:326-339` sets `e2e_persona=member` client-side) — a driver must set that cookie itself. No doc says this.
- **Zero** existing assets exercise a content seat: `grep 'content-player|content_products|Content stories'` across `stack-verify/e2e/` and `playthroughs/` returns nothing.

### G2 (critical) — no documented mechanism for a runtime-resolved URL
Every coverage entry path is a static string: exact-path match (`pageDescriptorFor`, `coverage-manifest.ts:988-992`) or the static `seedPaths` prime (`coverage-manifest.ts:131`), which `coverage-protocol.md:406` itself demotes ("seed paths are guesses, not coverage commitments").

### G3 — `VantageManifest.identityKey` is SINGULAR
`coverage-manifest.ts:129` — one seat per sweep. Covering 13 player seats means **13 logins/manifests, not one parameterized sweep**. This is the genuine structural cost nobody costed.

### G4 — `academy-seed`'s flag surface is undocumented
`ant-academy.md:63` records the binary's existence in a parenthetical but not its usage — no flags, no fixture names, no invocation, no note that it is a **platform** binary needing a built `app`. `academy_chapter_progress` appears as a bare table name; the real schema detail (unique `user_id + chapter_slug`) lives in `content-stories-routes.md:302`.
Verified directly: `app/cmd/academy-seed/main.go:51-56` exposes `--user-email`/`--user-id`, `--fixture` (default `starter`), `--reset`, `--dry-run`, `--list`; `fixtures.go:30,34,57,87` confirms `empty` / `starter` / `in-progress` / `completed`. **M236 `overview.md:46`'s `--user-id <owner> --fixture in-progress|completed` is ALIGNED with the real binary.** Caveat: `stack-dev/app` is at `64e20735` (v1.335.0), **63 commits behind origin/main** — re-verify after the M236 cold pull.

## Things that are CHEAPER than scoped (scope-back candidates)

1. **Arbitrary-seat login already works.** `playthroughs/e2e/lib/hero-login.ts` (42 lines) takes `identityKey: string` with **no enum/allowlist** (`:23`), delegating to `cockpit-login.ts:57-94`. The "needs NEW seat-login plumbing" premise is **overstated** — what is new is the iteration + page-object layer, not the login.
2. **Result URLs are pre-resolved.** `player_result_path` and `manager_result_path` are already fully-formed strings in `content-manifest.json` (e.g. `/sim/ai-readiness-interview-d62/result/1199f27e-…`). "Resolves each session's exact result URL" is a **field read**, not a resolution step.
3. **No new tailnet port is needed.** `gen_tailscale_serve.py` is untouched since M226; content-stories result pages live in next-web (`3000+off`, `:43`) and the academy grid (`3077+off`, `:50`) — both already fronted.
4. **Both new ant-academy demopatches verified sha-CURRENT** against the local clones and cached `origin/main`: `ant-academy-dev-origins` pre_sha256 `6837cab9…50a8e3` matches `code/next.config.js`; `academy-fs-published-fallback` pre_sha256 `43977541…fb9665` matches `code/src/lib/serverTenant.js`.

## Environment / tooling-gap facts M236's strategy must account for

- billion pins rext at `casting-call-m228-hiring-scope-fix`; target is `playbill-m235-hardened @ 60eff14` — **20 commits** apart, touching `demo-stack/up-injected.sh` (+81), `demo-stack/cockpit.py` (+304), `demo-stack/ant-academy.sh` (+32), `stack-injection/apply-academy-fs-published.sh`, and **three net-new sha-pinned demopatch manifests** (`academy-fs-published-fallback`, `next-web-interview-flag-{container,result}`). The running demo-1 is **not upgradeable in place** — cold re-clone + re-tag + cold bring-up.
- Corroborated independently by a peer session: the v2.5 rext tooling is **unpublished** — `origin/main` is still the M228 commit and **zero `playbill-*` tags exist on origin**, while `ensure-clones.sh:85` (`DEMO_ALLOW_UNPINNED_REXT`) **aborts on tag drift**. billion structurally cannot obtain the feature under test until the tooling is pushed. No corpus doc mentions a publish step in the per-stack tag-consumption path.

## Applied Fixes

**None.** Audit-only run per the caller's mandate; no corpus, plan, or rext file was modified.

## Open Items (require user decision)

1. **F2 — the exit-gate wording.** "Reachable only over the tailnet" is unprovable as written on a Linux VM given `safety.md:405`/`:424` and `tailscale-serve.md:626-627`. Reword to the honest posture, or add perimeter work to scope. **User call.**
2. **F1/F4/G1/G2 — `delivers: none` vs the AUTHOR mandate.** `overview.md:12` (and the roadmap) say `delivers: none`, while the `In:` list mandates authoring a page-object, new seat-login sweep plumbing, and **reversing** a documented protocol rule. Under Phase 4 this needs an explicit `Delivers → knowledge/…` line (minimally `coverage-protocol.md` + `playthroughs.md`) before build proceeds.
3. **F5 — re-point or author the iteration protocol.** `verification.md` supplies a gate input, not a loop.
4. **F3 — authorize the two-line cockpit/harness change** (`data-login-as` + a `content` vantage case), or accept that the p95 half of the gate is unmeasurable for the seat the milestone is about. Note both are rext-side, so **0 platform edits** still holds.

## Gate Result

**RED — blocked.** `/developer-kit:build-mstone-iters` must not enter its bootstrap tok against these docs as written. Minimum to clear:

- Correct F1 in `overview.md` (the falsified page-object claim) — otherwise the strategy is authored against a nonexistent asset.
- Correct F20 and F21 in the inherited `In:` list — one task is a `git checkout` rather than a re-anchor, the other is an *author* rather than a *run*. Left as written, M236 mis-sizes both.
- Resolve F2 (gate wording) and F5 (protocol ref) — both are frontmatter-level and cheap.
- Record F3/F4/F6/F7/F10 as known-stale so the iter loop does not read them as truth; F6, F7 and F10 will each cause a concrete failure on the first cold billion bring-up.
- The remaining findings (F8, F9, F11–F19, F22–F24, G3, G4) may be tracked as `KB-{N}` items and fixed inside M236's own doc-sync, provided they are recorded now.

**Net effect on M236's shape:** the milestone is **larger than scoped in authoring** (a page-object, a card-count assert, and a protocol reversal must all be written) and **smaller than scoped in plumbing** (arbitrary-seat login already works; result URLs are pre-resolved; no new tailnet port; the clone "drift" is a checkout). It remains a genuine iterative build, and the 0-platform-edit constraint still holds — every fix identified is rext-side or doc-side.

**SEVERITY: blocker**
