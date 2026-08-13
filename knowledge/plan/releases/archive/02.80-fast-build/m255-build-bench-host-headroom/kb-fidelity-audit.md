---
title: "KB Fidelity Audit — M255 build-bench & host-headroom"
date: 2026-07-27
scope: milestone:M255
invoked-by: user
code-baseline: "rosetta-extensions @ 29452cc (main) + UNCOMMITTED working tree — see §0"
---

## Verdict

**RED** — 3 blockers. Two of them are **refuted premises the milestone's own deliverables are specified
against**; the third **misdirects the barrier spike**. None is a "the doc is a bit stale" finding; each one
would make M255 build the wrong thing.

`KB-FIDELITY: RED` · blockers 3 · stale-load-bearing 9 · blind-areas 4 (3 planned-covered, 1 not) ·
incidental 21 · known-tracked (D-v28-10 → M257) 4 · process 1

---

## §0 — Read this first: the code moved under the audit

**M255 implementation is already in progress, concurrently with this pre-milestone gate.**

At audit start (11:5x) both repos were clean. By 12:30 the `rosetta-extensions` authoring copy held:

| Path | State | mtime |
|---|---|---|
| `stack-core/hostprofiles/billion.json` | untracked | 12:16 |
| `stack-core/buildbench.py` (49 KB) | untracked | 12:20 |
| `stack-core/union_apply_guard.py` | untracked | 12:22 |
| `stack-core/tests/test_{buildbench,union_apply_guard}.py` | untracked | 12:27 |
| `demo-stack/up-injected.sh` | modified (+104/−10) | — |
| `stack-core/demo_knob_guard.py` | modified | 12:29 |

…and the `rosetta` corpus held `M CLAUDE.md`, `M corpus/ops/demo/README.md`,
`M corpus/ops/demo/demo-up-defaults.md` (the latter: `DEMO_DISK_MIN_GIB` `20 → 25`, a net-new
`DEMO_CERT_RENEW_DAYS` row, `all 27 → all 30`).

Three consequences that matter:

1. **The gate is being run against a moving target.** Every `up-injected.sh` line number in
   `overview.md` and `roadmap.md` is HEAD numbering; the worktree has shifted the build functions by
   **+22** (see B-2). The line anchors were correct when written and are wrong now.
2. **Three shipped artifacts already cite `corpus/ops/demo/build-budget.md`, which does not exist** —
   `union_apply_guard.py:48`, the operator-facing warn at `up-injected.sh:339`, and a *rendered markdown
   link* at `demo-up-defaults.md:151`. The new `DEMO_CERT_RENEW_DAYS` row also points at
   **`safety.md` §3.9**; §3 currently ends at §3.8.1.
3. **No inline fixes were applied by this audit.** Editing corpus docs that a parallel session is
   actively rewriting risks clobbering its work, and all three blockers need a decision, not a one-line
   correction. This is a deliberate departure from Phase 6's "apply when possible" — recorded here as the
   reason.

Line numbers below are **working tree** unless marked `@HEAD`.

---

## Topic Inventory

| # | Topic (M255 in-scope item) | Knowledge doc | Code | Status |
|---|---|---|---|---|
| 1 | buildbench ledger + `DEMO_*` env snapshot | `verification.md` (autoverify), `latency-budget.md` (precedent) | `stack-verify/live/autoverify.sh`, `up-injected.sh:2493-2500` | **PAIRED** |
| 2 | Campaign protocol + reclaim + disk runway | `frontend-tier.md:267-289` (partial, partly wrong); `rosetta_demo.md` **silent**; `demo-up-defaults.md` 1 cell | `up-injected.sh:315-341`, `rosetta-demo:189-224,329-343` | **BLIND-AREA** → covered by Delivers (item 6) |
| 3 | Host profiles + headroom assert | `frontend-tier.md:246-265` (12 GB VM), `verification.md` | `preflight_vm_ram:246-270`, `preflight_disk_headroom:315-341` | **PAIRED** |
| 4 | Union-apply parallelism rule | `demopatch-spec.md` | `patches/demopatch`, `up-injected.sh:512-531/1030-1069`, `union_apply_guard.py` | **PAIRED — 2 blockers** |
| 5a | Spike (a) — L1 multi-stage on `hiring.Dockerfile` | `frontend-tier.md` | `demo-stack/frontend/hiring.Dockerfile` | **BLIND-AREA — blocker** |
| 5d | Spike (d) — plateau or I/O ceiling | — (no doc; evidence only) | sampler | DOC-ONLY (evidence `§6`) — acceptable |
| 5e | Spike (e) — host-vs-peer topology | `tailscale-serve.md:317-341` **fully documented** | `playthroughs/e2e/run-playthroughs.sh:54-64` | **PAIRED — already answered** |
| 6 | `build-budget.md` | — (net-new) | — | **BLIND-AREA — planned (Delivers)** |
| 7 | Cert renewal + `safety.md` §3 amendment | `safety.md` §3 **silent**; `tailscale-serve.md:541-548` partial | `up-injected.sh:1859` @HEAD → `cert_needs_mint:1352` (worktree) | **BLIND-AREA — planned (Delivers)** |

---

## BLOCKERS

### B-1 · D-v28-7 premise **(f)** is REFUTED — the item-4 guard test is specified against a false claim

**Source:** `overview.md:85-88` and `roadmap.md:242-244`:

> the one shared-package outlier — `next-web-ssr-graphql-origin` → `packages/graphql/src/server/server.graphql.ts` —
> is **inert for the hiring image by its own manifest header** (behaviour-identical when
> `WUNDERGRAPH_SSR_ENDPOINT` is unset; it only prepends to an existing `||` chain).

**Actual:**

1. The manifest header contains **zero occurrences of "hiring"**
   (`demo-stack/patches/next-web-ssr-graphql-origin/next-web-ssr-graphql-origin.yaml`). It says
   *"BEHAVIOUR-IDENTICAL when `WUNDERGRAPH_SSR_ENDPOINT` is unset"* (`:51`, `:79`) — a statement about the
   **variable**, not about the hiring image. The plan silently converted one into the other.
2. **The variable is NOT unset on hiring.** `stack-injection/gen_injected_override.py:367` emits
   `- WUNDERGRAPH_SSR_ENDPOINT=http://graphql:8080/graphql` **inside the `hiring-app:` service block**.
3. The import chain is real, one hop:
   `apps/hiring/src/app/api/bunny/recording/[sessionId]/route.ts:1` →
   `packages/graphql/src/server/bunnyRecordingDownload.ts:5` → `./server.graphql` (the patched file).

**Verdict:** STALE / refuted. **Fix owner: the plan, not the doc** — `overview.md:85-88` and
`roadmap.md:242-244` must be retracted. The in-flight `union_apply_guard.py:38-48` + `WAIVERS:75-87`
**already record a corrected, evidenced waiver**, and `test_union_apply_guard.py:67-76` fences it — so the
code is ahead of the plan and the two now disagree. Item 4's deliverable (h) says the guard must accept a
member *"explicitly waived as inert"*; the shipped guard deliberately requires an **evidenced** waiver
instead. Reconcile before M257 re-derives the false premise.

### B-2 · `demopatch-spec.md` §4 misstates the two facts the union-apply rule is built on

M255 item 4 delivers *"apply the union once, build both in parallel, **revert once LIFO**"*, preserving the
`urls.ts` chain. The doc it must extend gets both halves wrong.

**(a) The hiring fingerprint union is 7 manifests, not 4** — `demopatch-spec.md:163-164`:

> hiring's over the **4-manifest union** (the 2 `apps/hiring` patches + this shared pair). A test fences the
> hiring-side chain apply-order + LIFO revert + the **4-manifest fingerprint union**.

Actual (`up-injected.sh:1071-1075`): `next_web_patchset_fp` is called with **seven** — `rolemap`,
`pagination`, `studio`, `pubweb`, `interview_flag_container`, `interview_flag_result`, `back_to_cockpit`.
This is a **C1 mirrored-count violation** (`demopatch-spec.md:222-226`: *"any count mirrored in more than
one doc, or backed by a test fence, must move with all its mirrors and its fence in the same commit"*) —
the count drifted at M232 and again at M249 and was never moved.

**(b) "The apply-order … and revert-order (LIFO) are identical on both" is false** (`:162-163`). Neither
build reverts strict-LIFO and the two orders differ: next-web's trap (`:688`) reverts
pubweb, studio, pagination, ssr, aireadiness, thirdparty, result, container, back-to-cockpit; hiring's
(`:1149`) reverts result, container, pagination, rolemap, pubweb, studio, back-to-cockpit — both revert
`back_to_cockpit` **last** although it applies last. Only the load-bearing invariant (pubweb before studio)
holds in both. **"Revert once LIFO" therefore has no existing referent to inherit** — M255 must specify the
union revert order explicitly (chain-pair-last, reverse-apply otherwise).

**(c) Related, same section — G5's revert is whole-file-sha gated.** `demopatch:366-371` compares against
the **recorded** `pre_sha256`/`post_sha256`, while `cmd_apply:323` deliberately writes a file whose sha is
the **recomputed** post (the self-healing gate, §6). So **whenever self-heal fires, revert refuses** — and
the traps pipe it to `>/dev/null 2>&1`. The G5 row (`:52`) does state the refusal but labels its trigger
*"manual drift"*, which is exactly the case it is **not**. Halving the revert cycles (the rule's stated
benefit) halves the chances of recovery before R1's next `--force-pristine` sweep.

**Fix owner: doc** for (a) and (b); (c) needs a sentence in the doc **and** a decision on whether the union
revert uses `--force-pristine`.

### B-3 · `frontend-tier.md` argues **against** spike (a) — the barrier's decider

Spike (a) prototypes a multi-stage shape on `demo-stack/frontend/hiring.Dockerfile` *"which rext already
owns outright (no demopatch, no platform-edit question)"*. The doc a spike author reads says the opposite.

- `:13-16` (the hard-line box) states the model as exhaustive: *"next-web-app, studio-desk, and ant-academy
  stay byte-for-byte pristine — their repos are used only as a Docker build context (**their Dockerfiles
  consumed UNMODIFIED**)"*. The third shape — *rext authors and owns the Dockerfile; the platform repo
  supplies only the context* — is **never mentioned**. `grep -ci hiring frontend-tier.md` = **0**.
- `:613-619` (What's out of scope) routes Dockerfile work away: *"Those are real platform edits with
  PR/review/prod risk — **forbidden** for the demo tooling to make locally"*, listing
  *"optionally `output:'standalone'`"* among them. For **hiring** that is backwards: `hiring.Dockerfile`
  is rext's own file (`up-injected.sh:1031` — `# rext-owned`; the file's own `:6-11` explains why), and
  `NEXT_PRIVATE_STANDALONE=1` is an `ENV` in it, needing nothing from the platform.

**Verdict:** BLIND-AREA, **not** covered by D-v28-10 (which scopes only the four §8.5 prose numbers).
**Fix owner: doc** — one paragraph naming the rext-owned-Dockerfile precedent, before the spike runs.

---

## Known-tracked (D-v28-10 → M257) — confirmed still present, scope understated

| §8.5 claim | Site | Still literally present? |
|---|---|---|
| studio slowness *"pure memory starvation, not a slow build"* | `frontend-tier.md:248-250` | **yes, verbatim** |
| *"the ~3.7 GB build cache"* | `frontend-tier.md:271` | **yes** — now contradicted by an in-tree comment citing this exact line (`up-injected.sh:309-312`) |
| *"~3 min per frontend"* + 3 mirrors | `:231`, mirrors `:249`, `:262`, `:271` | **yes**; also mirrored **in code** at `up-injected.sh:816` and `:1251` |
| *"hiring" mentioned zero times* | whole doc | **yes — 0 of 623 lines** |

> ⚠️ **Scope note for D-v28-10.** §8.5 frames the hiring gap as *"the total undercounts by a whole 208 s
> image"*. It is **9 structural sites**, not one total: the tier definition (`:3-4`), the UI-tier port table
> (`:20-24`, no hiring row, no `:23001`), *"the **two** frontend builds run one at a time"* (`:236-237`),
> the per-demo image enumeration (`:269-272`, omits `demo-N-hiring` — the list a disk budget derives from),
> the per-frontend residual (`:231-232`, `:618`), the `--scale …=0` recovery list (`:241-243`, code does
> three at `:1978`), the `gen_injected_override.py` description (`:580-582`), `next-web.dockerignore`
> described as next-web-only (`:295`, `:599` — hiring reuses it at `:1130`), and `3001` attributed to
> next-web in the CORS section (`:355-356`, `:365-366` — it is hiring's origin).
> **Recommend annotating D-v28-10 with this enumeration** so M257 rewrites the model, not just the numbers.

---

## Stale load-bearing claims (non-blocking, but the milestone will read them as truth)

| # | Site | Claim | Reality | Owner |
|---|---|---|---|---|
| S-1 | `overview.md:50`, `roadmap.md:292` | `autoverify.json` emits `project/offset/warnings/green/ts` at `autoverify.sh:381-385` | **Field list exact.** Anchor off: `:381` is a comment, `:386` (the write target) excluded → **`:382-386`**. Scope *is* computed at the call site (`up-injected.sh:2493-2500`, `--services`) and thrown away — say *"no machine-readable record"*, not *"indistinguishable"* (`autoverify.log` rows differ) | plan |
| S-2 | `overview.md:110`, `roadmap.md:320` | host-vs-peer cited to `run-playthroughs.sh:56-72` | Path is **`playthroughs/e2e/`**, not `stack-verify/e2e/`. Content confirmed at `:54-64` | plan |
| S-3 | `overview.md:81-93`, `roadmap.md:240` | manifest lists at `up-injected.sh:490`/`:1008`/`:496-509`/`:1020-1047` | Correct @HEAD; worktree = `:512`/`:1030`/`:518-531`/`:1043-1069` (+22). **Arithmetic 9/7/11/5/5/1 is exactly right** — see inventory below | plan |
| S-4 | `demo-up-defaults.md` | 29 `Read at` citations | **28 of 29 are drifted** (only `DEMO_ALLOW_UNPINNED_REXT` is right). `demo_knob_guard.py` checks **names only** — it never compares the default value or the `file:line`, so the fence is green while the anchors rot. Already reported at M236 and unfixed since | doc |
| S-5 | `demo-up-defaults.md:24` | *"`<N>` … and `--public-host <host>` — that is the entire flag surface"* | **`--no-public-host` is a third arm** (`up-injected.sh:40`); `:46-49` hard-refuses the pair. This is the sentence `overview.md` item 9 rests on; a bench harness written from it will not know how to force a deterministic localhost rep on a Tailscale host | doc |
| S-6 | `frontend-tier.md:228-230` | image reuse is *"tag-guarded … only a brand-new `demo-N` (or a frontend code/dep change) pays"* | Reuse needs **4 conjunctive** predicates (image exists · baked endpoint matches offset · minted pk grepped out of the bundle · `demo.patchset` label matches) — `:568-587`/`:857-873`/`:1089-1100`. **A code/dep change invalidates nothing** (no source hash), and the check is fail-safe-biased toward rebuild (`:559-561`). buildbench's cache-hit model depends on this | doc |
| S-7 | `frontend-tier.md:274` | *"Below ~20 GB free…"* | @HEAD correct; **worktree is 25** (`:316`), derived from `hostprofiles/billion.json` (18 + 7) | doc (M255's own change) |
| S-8 | `latency-budget.md:3-5` | intro = *"the click→ACCESS budget…"* | The doc absorbed a **second, structurally different budget** at `:328-468` (studio first-paint, own gate/harness/baseline) not named in the intro. **This is the doc M255 models `build-budget.md` on** — decide deliberately whether build-budget is a third graft or a true sibling | doc |
| S-9 | `verification.md:9-16`, `:22-25` | scope = *"the v1.3b/M18 net"*; *"**Out of scope:** the frontend tier — the frontends don't exist in the stack yet"* | Frontends have been in `verify_svcs` for many releases (`up-injected.sh:2494`). Body runs to M252. Also the M217 cheap-win inventory (`:194-204`) lists **4**; code runs **8** (`buildfail.log :202`, academy `:298`, studio AI key `:355-370` undocumented) | doc |

---

## Manifest inventory (verifies D-v28-7 arithmetic — **9 / 7 / 11 / 5 / 5 / 1 exactly right**)

| manifest | next-web | hiring | `path:` | tree |
|---|---|---|---|---|
| `next-web-studio-url` | ✅ | ✅ | `packages/core-js/src/constants/urls.ts` | shared |
| `next-web-public-website-url` | ✅ | ✅ | *same file* — **CHAINED** (`pre` == studio's `post` = `fe15aa71…`, byte-verified) | shared |
| `next-web-interview-flag-container` | ✅ | ✅ | `packages/ui/src/AISimulation/AISimulationResultContainer.tsx` | shared |
| `next-web-interview-flag-result` | ✅ | ✅ | `packages/ui/…/AISimulationResult.tsx` | shared |
| `next-web-back-to-cockpit` | ✅ | ✅ | `packages/ui/src/NavBar/NavbarTop.tsx` | shared |
| `next-web-members-pagination` | ✅ | — | `apps/web/src/context/InsightsContext.tsx` | **apps/web** |
| `next-web-aireadiness-flag-gate` | ✅ | — | `apps/web/src/components/ai-readiness/data/useAiReadinessActive.ts` | **apps/web** |
| `next-web-no-thirdparty` | ✅ | — | `apps/web/src/app/layout.tsx` | **apps/web** |
| `next-web-ssr-graphql-origin` | ✅ | — | `packages/graphql/src/server/server.graphql.ts` | **outlier — see B-1** |
| `next-hiring-role-remap` | — | ✅ | `apps/hiring/src/context/UserStatusContext.tsx` | **apps/hiring** |
| `next-hiring-members-pagination` | — | ✅ | `apps/hiring/src/context/InsightsContext.tsx` | **apps/hiring** |

5 shared + 5 disjoint (`apps/web` ×3 vs `apps/hiring` ×2) + 1 outlier = **11**. ✓
Guards **G1–G7 all exist and all match the doc's descriptions** (the *script's own header* at
`demopatch:17,25-31` is the stale one — it still says "six guards" and describes the pre-M217 sha gate).
The 10-key manifest schema matches `manifest_loader.py:33-36` exactly; **all 23 manifests carry exactly 10
keys**. The 3 apply vehicles, the chain rule, and the self-healing gate all verify. Patch count **23** is
current and fenced (`test_patch_inventory.py:43-44`) — but a bring-up performs **28 apply-events** (5
manifests applied twice), and `demopatch.log` is written on the *refuse/skip* branch only, by
`up-injected.sh` only — so the evidence doc's *"empty ⇒ all 23 applied"* (`build-annotation.md:21`) proves
"no logged refusal" and never covers the 5 academy patches at all.

---

## Live constraints on M255's design (code+test agree today; the milestone must move them)

| # | Constraint | Site |
|---|---|---|
| C-1 | A test **forbids making the disk pre-flight fatal**: `assertNotIn("exit 1", fn, "the disk preflight must be NON-FATAL")` | `demo-stack/tests/test_tooling.py:437` |
| C-2 | Literal floor strings pinned: `"< 20 GiB recommended"` / `"< 12 GiB recommended"`; and a 50-GiB-free case asserts the OK branch (flips if the new floor > 50) | `test_frontend_build.py:622`, `:599`, `:612-615` |
| C-3 | `DEMO_DISK_MIN_GIB` re-size is an **unfenced 8-site mirrored change** — SoT `up-injected.sh:316`, `demo-up-defaults.md:151`, `frontend-tier.md:271,274`, `.claude/skills/demo-up/SKILL.md:26-27`, + 3 test sites. `demo_knob_guard.py` will stay green either way | — |
| C-4 | `hiring-app` is **not in the autoverify scope** (`up-injected.sh:2494` adds only next-web + studio-desk) — buildbench measures a third image the green verdict never probes | code gap |
| C-5 | The in-flight `union_apply_guard.py` is **3/12 RED**: `test_a_reversed_urls_chain_is_refused` fails because the CHAIN-ORDER clause is **dead code** (`:195-201` — `present` is built by iterating `chain`, so `present == chain` always); `test_a_cross_app_patch_is_refused_outright` mutates into a *shared* member, not a cross-app one; `test_an_unparseable_script…` raises before asserting. The shared-member byte-identity check at `:154` is `if _sha(p) != _sha(p):` — a tautology, the second one in the file | code |

---

## Blind areas — status

| Area | Planned? | Note |
|---|---|---|
| `corpus/ops/demo/build-budget.md` | **YES** — Delivers, item 6 | Genuinely net-new. Only overlaps with `frontend-tier.md` are `:271` (stale cache size, known-tracked) and `:295`/`:301` (context trim — **correct, carry forward**). The doc is silent on image sizes, export/unpack, multi-stage, turbo concurrency, cache-mounts, per-cycle growth |
| `safety.md` §3 cert renewal | **YES** — Delivers, item 7 | §3 is completely silent on expiry/renewal/90-day/`--purge` survival. **Do not restate the 90-day fact** — `tailscale-serve.md:541-548` already owns it (and already carries the M220 "does not re-issue on re-run" correction). Net-new content: certs survive `--purge`, the absence-only guard made the first mint the only mint, nothing at bring-up detects it. **Insertion point: a new `#### 3.5.4` at line 768**, immediately before `### 3.6` — §3.5.1's rung-6 bullet (`:689-692`) is the existing home of the "silently degrades to local-trust" argument and this is its expiry-side twin. *(The new `demo-up-defaults.md` row already forward-references "§3.9" — pick one and mirror it.)* |
| Campaign protocol / disk runway / reclaim | **YES** — via item 2 → `build-budget.md` | `rosetta_demo.md` is a **total blind area** on teardown cost (0 hits for disk/cache/prune/ENOSPC; `--purge` appears only as a verb at `:233`). `demo-up-defaults.md` has one table cell. The narrative's only existing home is `frontend-tier.md:267-289`, which is partly wrong. ⚠️ `.claude/skills/demo-up/SKILL.md:20` promises `rosetta_demo.md` carries the *"resource budget"* — it never did |
| **rext-owned-Dockerfile precedent** | **NO** | **B-3.** Not in any Delivers line, not in D-v28-10 |

---

## Spike scope corrections (before you spend the time)

- **Spike (e) is already answered in the corpus.** `tailscale-serve.md:317-341` is a dedicated fenced
  subsection — *"NOT from the VM itself — `tailscale serve` is bypassed on the loopback path (M219)"* —
  with the mechanism (`:322-325`), the `curl: (35) … wrong version number` error (`:328`), the measured
  three-port evidence (`:331-332`), the testing consequence (`:335-341`), and an explicit retraction of the
  old *"(or the VM itself)"* parenthetical that *"cost M219 a full false-RED sweep"*. **Budget it as a
  citation task**, not a 20-minute investigation; M258's gate text can be lifted from `:335-341`.
- **`--public-host` default-on is CONDITIONAL, not unconditional.** A bare `up-injected.sh N` goes public
  only if **all six** `tailscale_autohost.py:16-24` rungs pass (binary on PATH · backend Running ·
  dotted RFC-1123 `DNSName` · MagicDNS enabled · no operator/sudo denial · **`tailscale cert` actually
  mints**). Any failure ⇒ a byte-identical localhost demo. For M258's gate text this distinction is the
  whole point. (`safety.md:685-693` documents the ladder; its rung-3 wording under-specifies the shipped
  hostname regex refusal.)
- **Item 3's "one assert" vs `build_frontends()`.** Verified: `build_frontends` (`:1262-1287`) has exactly
  one *gating* conditional, `[ "$NO_UI" = 1 ]`; a second `if` at `:1272` is post-hoc failure reporting.
  `preflight_vm_ram` (`:246-270`) declares all three vars `local`, assigns no global, is called bare and
  uncaptured at `:1547`, and is referenced nowhere else — **the 12 GiB warning is genuinely cosmetic.**
  `frontend-tier.md:252-256` already says so explicitly and correctly; no doc correction needed there.

---

## Incidental findings (record as KB-N, fix opportunistically)

1. `safety.md:493` *"There is no `127.0.0.1` prefix anywhere, in either family"* — **false**;
   `gen_injected_override.py:602` binds the fake BAPI to loopback deliberately, and `safety.md:789-793`
   calls it *"the demo's FIRST loopback-bound port"*. Self-contradiction inside §3.
2. `safety.md:490-491` *"all **three** emitters"* — now **four**; `hiring_lines`
   (`gen_injected_override.py:328`, port at `:370`) publishes bare, and
   **`exposure_claim_guard.py:140-142` does not sweep it** — the exposure fence is blind to the fourth emitter.
3. `safety.md:547` — the injection *"still carries a dormant, inert `skillpath` entry"*; removed at M247
   (`gen_injected_override.py:18,21`).
4. `safety.md:512` — `ant-academy.sh:330` anchor drifted ~136 lines → **`:466`**. (The only rext `file:line`
   anchor in safety.md.)
5. `safety.md:32-34` — Scope block frozen at *"v1.3 stack party"*; body runs to v2.6 §3.8.1.
6. `tailscale-serve.md:30-31` — *"Everything below is unchanged and still correct. Only the trigger
   changed"*; the doc corrects itself twice further down (`:347-352`, `:509-534`).
7. `latency-budget.md:44` — both `cockpit.py` anchors drifted: `:618` → **`:872`**, `:423` → **`:547`**
   (substance correct).
8. `latency-budget.md:315-321` — *"all **four** runners"*; **six** carry the guard
   (`run-discrete.sh:41`, `run-studio-fcp.sh:38` added). The regression test
   `test_green_gate_age.py:256-268` is stale the same way, leaving two runners unpinned.
9. `latency-budget.md:445` — the runnable example passes `maya-thriving` to `run-studio-fcp.sh`; `:43`
   changed the default to `dan-manager` at M254 *because a member seat is bounced*. The doc contradicts
   itself (`:434` is right).
10. `latency-budget.md:467` — `dead_shell_gap_ms` → `dead_shell_gap_p50_ms` (`studio-ttu.spec.ts:172`).
11. `verification.md:190` — `<stack>/autoverify.log` shorthand, which `:236` explicitly corrects for the
    `.json` twin.
12. `verification.md:346-349` — the M217 rext-pin guard described as unconditionally FATAL;
    `ensure-clones.sh:85-86` has an undocumented `DEMO_ALLOW_UNPINNED_REXT=1` bypass, and `:73-75` loosened
    the check to "is the SoT tag among HEAD's tags". Worth a sentence given Rung Zero is named in M255.
13. `verification.md:135-148` — the readiness scope gate attributed to `lib/readiness.sh`; it lives in
    `live/verify.sh:57`.
14. `frontend-tier.md:217-218` — *"or the `--no-ui` equivalent"*; `--no-ui` is **not** a `/demo-up` flag
    (`up-injected.sh:41` → exit 1). It is internal, threaded to `gen_injected_override.py`.
15. `frontend-tier.md:387-395` — the §ant-academy heading and body still describe the **removed** keyless
    `BENCHMARK_VISUAL_BYPASS` model in the present tense (`ant-academy.sh:530-536`: *"GONE from the launch
    env"*), contradicting `:24` and the doc's own correction boxes at `:28-76`/`:404-412`.
16. `frontend-tier.md:549-557` — the manual-fallback block tells operators to set `NEXT_PUBLIC_E2E_AUTH=1`
    + the `e2e_persona` cookie; `ant-academy.sh:253-270` prints the opposite and warns *"NEVER copy
    `CLERK_*` out of `platform/.env`"* — **and `:269` routes the operator to this exact block**. Following
    the doc reproduces the M220 session-poisoning defect.
17. `frontend-tier.md:519-521` vs `ant-academy.sh:19` — flatly disagree on whether ant-academy uses FA Pro
    (same conclusion: no token needed).
18. `frontend-tier.md:573-574` — rext pin *"`fit-up-m49`"*; `.agentspace/rext.tag` reads
    `july-jitter-m246-re-sync-repoint`.
19. `frontend-tier.md:597-598` — the `up-injected.sh` inventory line omits the disk pre-flight the doc
    devotes `:267-289` to.
20. `demopatch-spec.md:263`, `:334` — `app-aireadiness-snapshot-loadmembers` target given as
    `internal/workforce/ai_readiness.go`; the manifest says `internal/aireadiness/readiness.go` (moved at
    M254, `997272b`). Stale twins in `apply_patch.py:20` and `apply-app-authz-skip.sh:48`.
21. `demopatch-spec.md:148-150` — the shell-helper exit-code table lists `0-6`; `apply_patch.py:48-52`
    emits **`0|1|3|5`** only. Code `2` ("pre-sha drift") is conceptually dead — M217 removed it as a
    refusal, which is what §6 exists to explain.
22. `demopatch-spec.md:361-363` — *"**For each patch** it resolves the tag…"*; the freshness preflight
    iterates a hardcoded **two-entry** list (`up-injected.sh:1569-1571`).
23. `build-annotation.md:207` vs `:310-312` — the runway is attributed to *"§8.2's ~8 GB/cycle leak"* while
    the measurement is **2 G/cycle**. Reconcile before publishing the runway number in `build-budget.md`.
24. `rosetta_demo.md:86-87` — *"the **5** injected Go services"* (now **3**, `up-injected.sh:187`) and
    *"the **two** frontends"* (three). `:88`'s compose-build list omits `postgresql`
    (`stack-demo/platform/common.yml:2-4`), which **is** the 5th image compose rebuilds inline.
25. `demo-up-defaults.md:37` — *"all 27"*; **29** at HEAD (already being changed to 30 in the worktree).
26. `up-injected.sh:249-255` emits a per-stack-Directus (~1 GiB) budget note *before* the `NO_UI` early
    return; `frontend-tier.md:248` states runtime as *"~0.66 GiB for BOTH stacks"* with no Directus term.
27. `up-injected.sh:237` + `:1265-1270` write a `buildfail.log` evidence channel autoverify asserts on
    (*"A build that did not happen must never read as a build that succeeded"*, `:1280`) — undocumented in
    `frontend-tier.md`, and directly relevant to buildbench not grading a stale image green.
28. `gen_injected_override.py:455` carries a stale code comment twin of finding S-6's CORS attribution
    (*"Origins kept: next-web 3000/3001"*).
29. `dev-stack/dev-stack:298` does **not** set `STACK_DIR`, so `dev-N` bring-ups write no `autoverify.json`
    and no `autoverify.log`. Also: no `STACK_DIR` ⇒ **the entire M217 cheap-win block is skipped**
    (`autoverify.sh:175`), not just the JSON — neither doc says this.
30. Stray mutation-test artifact on disk: `dev-stack/.m220-mutant-20farqj4dev-stack` (a full copy of
    `dev-stack/dev-stack`), which pollutes grep-based audits.
31. `demopatch:17,25-31` — the script's own header still says *"THE SIX MANDATORY SAFETY GUARDS"* and
    describes the pre-M217 G2 sha gate and the retracted G4 post-sha wording. **The doc is right and the
    code comment is stale** — the inverse of the usual direction.

**Cross-references: all resolve.** Every markdown link in all 7 audited docs points at a file that exists;
deep anchors verified for `verification.md` → `alignment_testing.md:174`, `frontend-tier.md` →
`rosetta_demo.md:99` and `ant-academy.md:110`, `demopatch-spec.md` → `latency-budget.md:328`. The only
dangling target is `build-budget.md`, cited from **code and `demo-up-defaults.md`**, not from these docs.

---

## Applied Fixes

**None.** See §0(3) — the corpus files under audit are being modified concurrently by an in-flight M255
implementation session, so inline edits risk clobbering; and all three blockers require a decision
(retract a plan premise / specify a revert order / author a missing model paragraph), not a mechanical
correction.

---

## Gate Result

**RED — blocked.** `/developer-kit:build-milestone` must not enter Phase 1 until:

1. **B-1** — `overview.md:85-88` + `roadmap.md:242-244` retract the *"inert by its own manifest header"*
   premise and adopt the in-flight guard's evidenced-waiver framing. *(Cheap: the corrected text already
   exists in `union_apply_guard.py:38-48`.)*
2. **B-2** — `demopatch-spec.md` §4 fixes the 4→7 fingerprint count (C1: with its fence, same commit),
   retracts *"revert-order (LIFO) identical on both"*, and states G5's self-heal/revert asymmetry. M255
   then specifies the union revert order explicitly.
3. **B-3** — `frontend-tier.md` gains the rext-owned-Dockerfile precedent (and `:613-619` stops routing it
   to a forbidden upstream PR) **before spike (a) runs**.

**Also resolve before/with Phase 1** (not gate-blocking, but they will bite):
`§0(2)` the three dangling `build-budget.md` citations + the `safety.md §3.9`-vs-§3.5.4 mismatch ·
`C-1`/`C-2` the two tests that block the disk-guard promotion and pin its literal strings ·
`C-5` the 3 red guard tests + 2 tautological clauses · `S-1`/`S-2`/`S-3` the drifted plan anchors ·
`S-5` the `--no-public-host` omission (buildbench needs deterministic localhost reps on a Tailscale host) ·
and **annotate D-v28-10** with the 9-site enumeration so M257 rewrites the model, not just the numbers.
