---
title: "Deferral Audit — M218 seat change (milestone scope)"
date: 2026-07-14
scope: milestone
invoked-by: close-milestone
---

## Verdict

**YELLOW**

No unresolved repeat-deferral. The one **REPEAT** in scope (`x/crypto@v0.52.0`) is fated **LAND-NOW** in this
close, which discharges it. Every other item is a **first** deferral with a named, edited destination inside
this release (Fate 3) — no item leaves v2.3, no item routes to "backlog" or "a future milestone (unspecified)".

## Summary

| | count |
|---|---|
| Total items in scope | **13** (12 inherited/carried + 1 surfaced by this close) |
| LAND-NOW (Fate 1) | **3** |
| LAND-NEXT (Fate 2 — already owned, confirm only) | **1** |
| LAND-NEXT (Fate 3 — annotate a sibling's `overview.md`) | **9** |
| DROP | 0 |
| KEEP-DEFERRED-WITH-SIGNOFF (escape hatch, cross-release) | **0** |
| Repeat deferrals | **1** (resolved by LAND-NOW) |
| Aged-out | **1** (same item) |
| Chronic patterns | 0 |

3 + 1 + 9 = **13.** ✓ **Zero escape-hatch deferrals — nothing escapes v2.3.**

**Also carried (not a deferral — a known-issue sign-off):** the **33 pre-existing `dev-stack` test failures**
are environmental (they need a live Postgres on `:15432`; this box's `.agentspace/secrets` is also incomplete,
so the secret-coverage pre-flight aborts before the N-guard). Identical count at v2.2 and at M217. **Not an
M218 regression** — M218 does not touch `dev-stack/`. Flagged for the user's sign-off in the Gate Outcome
Ledger, not silently elided.

## Deferral Inventory

Sources walked: M217 `progress.md` / `retro.md` (closed, inherited), M218 `progress.md` (carry-forward queue),
M218 `hardening-ledger.md` (routed-forward table), M218 `decisions.md`, M218 `overview.md` (§"Also in scope"),
`roadmap.md` §"Deferrals folded in", `state.md` §"Standing backlog".

| id | item | origin | first deferred | destination as-found |
|---|---|---|---|---|
| DEF-M218-01 | `x/crypto@v0.52.0` bump (13 dependabot alerts, all govulncheck-UNREACHABLE) | v2.2 residual → M217 retro | v2.2 close | **"M218's rext roll"** |
| DEF-M218-02 | **Clerk telemetry off** (`CLERK_TELEMETRY_DISABLED` + `NEXT_PUBLIC_*`) | M218 `overview.md` §Also-in-scope | 2026-07-13 | `FIX-M218-telemetry-egress` ⚠ |
| DEF-M218-03 | **C-5** — vendor clerk-js; bound the unbounded `Timeout: 0` | M218 `overview.md` §Ranked suspects ("take it regardless") | 2026-07-13 | `FIX-M218-c5-clerkjs` ⚠ |
| DEF-M218-04 | **ant-academy is handed the REAL `CLERK_SECRET_KEY`** (`ant-academy.sh:146` copies from `platform/.env`) | M218 `overview.md` §Also-in-scope | 2026-07-13 | _(unassigned)_ ⚠ |
| DEF-M218-05 | **F-7** — `NEXT_PUBLIC_BACKEND_API_URL` bakes to a 10.5 s blackhole from inside the container ("a loaded gun") | iter-03 | 2026-07-13 | `PROBE-M218-backend-api-url-twin` ⚠ |
| DEF-M218-06 | **F-5** — GA / DoubleClick / Google Ads / LinkedIn Ads egress on every authenticated load | iter-03 | 2026-07-13 | `FIX-M218-telemetry-egress` ⚠ |
| DEF-M218-07 | **C-3** — cms/Directus **403s** (`getSkillPaths`, `_entities JobSimulation`) on the CONTENT path | iter-04 | 2026-07-13 | `PROBE-M218-c3-rerun` ⚠ |
| DEF-M218-08 | **F-11** — the BAPI fabricates `organization.public_metadata.eid` (`"org_eid_"+orgID`) instead of the roster's real org UUID | harden pass 1 (D16) | 2026-07-14 | `FIX-M219-bapi-org-eid` |
| DEF-M218-09 | `expressrun` is **UNMEASURABLE** on this box (no `@clerk/express` `node_modules` → rc=2, *no score*) | harden pass 1 | 2026-07-14 | `TEST-M219-expressrun-dep-gate` |
| DEF-M218-10 | The demo-patch live-clone **freshness gate SKIPS** when `stack-demo/next-web-app` is absent | harden pass 1 | 2026-07-14 | `TEST-M219-freshness-gate-skips` |
| DEF-M218-11 | `DOC-M218-audit-corrections` — **remaining:** the stale `corpus/services/clerkenstein.md:3-4` header | M218 overview §Delivers | 2026-07-13 | `DOC-M218-audit-corrections` ⚠ |
| DEF-M217-01 | The **pre-bind reap has still never run live** (unit-proven, not field-proven) | M217 close | 2026-07-13 | **M221** ✅ |

⚠ = **destination is a handler tagged to the milestone that is closing right now.** Under the three-fate rule
these are not fates at all — `FIX-M218-*` on a closing M218 is the "future milestone (unspecified)" anti-pattern
in disguise. Every ⚠ row is re-fated below.

## Repeat-Deferral Patterns

### REPEAT + AGED_OUT: DEF-M218-01 — `x/crypto@v0.52.0`

- **First deferred:** v2.2 close — recorded as a v2.2 residual in `state.md` §Standing backlog.
- **Re-fated:** v2.3 design (`roadmap.md` §"Deferrals folded in") → **"Fate-1 → M218's rext roll"**.
- **Re-stated:** M217 `retro.md` §Carried forward → **"M218's rext roll"**.
- **Current status:** `clerkenstein/go.mod:16` still reads `golang.org/x/crypto v0.51.0 // indirect`. **It did not land.**
- **Ageing trigger:** *the milestone it was deferred to is closing without landing it* — and this close **is** the rext roll.
- **Time in limbo:** 2 releases, 2 milestones.
- **Pattern:** `DRIFT_DEFER` — never rejected on merit, simply never picked up. Not chronic (no "no time" reason was ever given).

**This is the one item that would turn the audit RED if deferred again.** It is fated **LAND-NOW**.

No other REPEAT groups. DEF-M218-02 … -11 are all first deferrals dated within this milestone.

## Fate-1 Investigation

### DEF-M218-01 — `x/crypto@v0.52.0`
- **Fate-1 feasible: YES.** `go get golang.org/x/crypto@v0.52.0` in `clerkenstein/` (the only module pinning it directly; `stack-seeding/go.sum` carries it transitively). It is an **indirect** dependency.
- **Landing scope:** bump + `go mod tidy` + full Go suite + **the 56-gene alignment gate** (which is what fences behavioural neutrality on the mirror). The measured login path (`clerk-frontend/`, `clerk-backend/`, the SSR demo-patch, `gen_injected_override.py`) is not edited.
- **Honest caveat, recorded:** this rebuilds the clerkenstein binaries, so it perturbs the artifact the 5-cycle latency battery was graded on. It is **not** a behavioural change (indirect crypto transitive; alignment gate green at 100% critical), and **M221 re-grades the latency gate on `billion` with no flags by design**, so the perturbed binary is re-proven there. The alternative — defer again — makes this a repeat-deferral and is strictly worse.

### DEF-M218-11 — the `clerkenstein.md:3-4` header
- **Fate-1 feasible: YES.** A documentation edit in `rosetta`. Zero runtime surface. The milestone's own `Delivers →` list owns it.

### The `stack-injection` regression (surfaced by this close's Phase 4, added to the ledger here)
- **Fate-1 feasible: YES — must-fix.** `test_next_web_block_shape` is **RED**. iter-03's own fix added `WUNDERGRAPH_SSR_ENDPOINT` to the next-web block and never updated the exact-shape fence. Verified **GREEN at `cue-to-cue-m217`** (the true pre-M218 baseline) and **RED at HEAD** ⇒ it is an **M218 regression**, not the "pre-existing / M217 footprint" the hardening ledger claims. That misclassification is itself a **D17 stale-verdict instance** and must be corrected in the ledger.

### DEF-M218-02 / -03 / -04 / -05 / -06 / -07 — the runtime-surface items
- **Fate-1 feasible: NO — and for a principled reason, not a time reason.** Each mutates the **demo runtime on or adjacent to the login path** (`clerk-frontend/server.go` for C-5; the injected container env for telemetry; `ant-academy.sh` for the secret). **iter-05 D13 established that a runtime change restarts the 5-cycle cold-battery count**, and **D16** rejected exactly this move for F-11 on exactly this ground: *shipping a runtime change after the gate was graded means shipping something other than what was measured.* Landing them now would either (a) invalidate the gate this milestone exists to prove, or (b) require a fresh 5-cycle battery on `billion` — which is **M221's declared job**.
- ⇒ These are **Fate 3**: annotate the sibling milestone of **this release** that should pick each up.

### DEF-M218-08 / -09 / -10
- **Fate-1 feasible: NO** (same graded-artifact bind for -08; -09 is a missing host dependency; -10 is a test-infrastructure gap). All three already carry `M219`-tagged handlers. **Fate 3 — but the target's `overview.md` must actually be EDITED**, which it currently is **not** (verified: `m219-readiness-renders/overview.md` mentions none of them). Recording a handler string in the *closing* milestone's own progress table is not a fate; editing the *receiving* milestone's plan is.

### DEF-M217-01 — the pre-bind reap
- **Fate 2 — already owned.** M217's retro routes it to **M221**, and `state.md` carries the ⚠ banner. `m221-prove-on-billion/overview.md` re-proves everything on the box. **Confirm, no edit.**

## Recommendations

| id | verdict | destination | action |
|---|---|---|---|
| DEF-M218-01 | **LAND-NOW** (Fate 1) | M218 | bump `x/crypto` → v0.52.0; re-run Go suites + alignment gate |
| DEF-M218-11 | **LAND-NOW** (Fate 1) | M218 | fix the `clerkenstein.md:3-4` header |
| _(new)_ `test_next_web_block_shape` | **LAND-NOW** (Fate 1) | M218 | update the fence; correct the hardening ledger's false "pre-existing" claim |
| DEF-M218-02 | **LAND-NEXT** (Fate 3) | **M220** | Clerk telemetry off — M220 owns `/demo-up` defaults + the injected-env contract |
| DEF-M218-06 | **LAND-NEXT** (Fate 3) | **M220** | ad-tech egress on authenticated loads — same surface, same milestone |
| DEF-M218-03 | **LAND-NEXT** (Fate 3) | **M220** | vendor clerk-js + bound `Timeout: 0` — removes an **unbounded internet dependency** from a demo that claims to be self-contained; belongs with M220's `safety.md` **Part 3** (the exposure axis) |
| DEF-M218-04 | **LAND-NEXT** (Fate 3) | **M220** | ant-academy's **real** `CLERK_SECRET_KEY` in a demo process — a `safety.md` violation of the fix16/17 class; pure env |
| DEF-M218-05 | **LAND-NEXT** (Fate 3) | **M221** | F-7 the blackhole twin — a **probe on the box**, over the tailnet, where it would actually fire |
| DEF-M218-07 | **LAND-NEXT** (Fate 3) | **M221** | C-3 the cms/Directus 403s — M221's gate item (2) is "the full replayed catalog, **no SKIPPED surface**" |
| DEF-M218-08 | **LAND-NEXT** (Fate 3) | **M219** | F-11 bapi org-eid — needs a runtime change **+ a fresh 5-cycle battery** |
| DEF-M218-09 | **LAND-NEXT** (Fate 3) | **M219** | `expressrun` dependency gate — must fail loud, not score-absent |
| DEF-M218-10 | **LAND-NEXT** (Fate 3) | **M219** | freshness-gate skip — *absence read as success*, the D17 class |
| DEF-M217-01 | **LAND-NEXT** (Fate 2) | **M221** | already owned; confirm only, no edit |

## Applied Changes

- **Fate 3 (7 items)** — `m219-readiness-renders/overview.md`, `m220-cue-sheet/overview.md`, and
  `m221-prove-on-billion/overview.md` each **edited** to add the inherited items to their `In:` scope.
  *(Applied by close-milestone Phase 7; verified before merge.)*
- **Fate 1 (3 items)** — added to M218's `## M218: Final Review` fix-queue in `progress.md`.
- **Fate 2 (1 item)** — confirmed covered by M221; no plan edit.
- Handler strings retagged in M218's carry-forward table so **no `*-M218-*` handler survives on a closed
  milestone**.

## Blocking Items (require user decision)

**None.** The single repeat/aged-out item (`x/crypto`) is fated **LAND-NOW**, which discharges it in this close.
Zero escape-hatch deferrals ⇒ no sign-off gate.

---

**DEFERRALS: YELLOW** — 13 items · 3 LAND-NOW · 10 LAND-NEXT (1 Fate-2 confirm + 9 Fate-3 annotations) ·
0 DROP · 0 escape-hatch · 1 repeat (resolved). `SEVERITY=warning`. Close proceeds.
