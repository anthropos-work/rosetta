# iter-10 — tik: the cold reset-to-seed reproduction (the closing proof)

**Type:** tik  ·  **Active strategy:** `TOK-01` (publish-then-prove) — the closing proof.

## Steps 1–3 — re-pin, tear down, rebuild from nothing

Re-pinned `billion`'s rext consumption clone to the final tooling tag
**`playbill-m236-latency-tz-fix`**, then `rosetta-demo down 1 --purge` (containers, network, and
container-owned data all removed — *"the next bring-up initdb's a fresh cluster → migrate → replay →
seed"*), then a cold `./up-injected.sh 1 --public-host billion.taildc510.ts.net`.

**No intervention was required.** The bring-up's own tail verdict:

```
✓ backend /api/health 200 on :18082
✓ sentinel.casbin_rules = 1250 (authz policy loaded)
✓ directus.directus_collections = 21 (the local Directus serves the captured catalog)
✓ directus DB is per-stack-local (not prod)
✓ verify live: all liveness + readiness probes passed
✓ demo-patches: all applied (none refused, none skipped)
✓ frontend builds: ok (the running images are this run's)
✓ taxonomy replayed: public.skills = 42790
✓ presenter cockpit answering on :17700
✓ clerkenstein fake-FAPI answering on :15400 (hero login is possible)
▶ autoverify demo-1: OK — verified-working.
==> [demo-1] UP. Clerk-free demo-1 is live.
```

**iter-07's cockpit-bind note resolved as predicted.** The cockpit now listens on **`0.0.0.0:17700`**, not
the `127.0.0.1` it had been restarted onto — confirming that bind was a self-inflicted artifact of
restarting in place after `tailscale serve` had claimed the port, not a defect. A cold bring-up orders it
correctly and fronts it with HTTPS itself.

**The manifest the cold cockpit serves is the corrected one**, projected from published tooling with no
hand-editing:

```
simulation:        13 sessions, 13 manager views
skill-path-legacy:  2 sessions,  0 manager views    ← iter-07's correction, reproduced
ai-labs:            2 sessions,  0 manager views
skill-path-new:     1 session,   0 manager views    → /courses/ai-foundations   ← iter-08's, reproduced
                   18 sessions, 13 manager views = 31 raw pairs − 2 presence-only = 29 landable
```

## Step 4 — re-measure all three, cold

| component | cold reading | gate |
|---|---|---|
| content-stories sweep | **29 / 29** (simulation 26/26 · skill-path 2/2 · academy 1/1), 1.8 min | all landable pairs ✅ |
| academy grid (Thread A) | **65 course links / 483 chapter links / 0 Draft chips** | non-zero, no chip ✅ |
| hero p95 — employee | **1.22 s** (p50 0.68 s), 5/5 ACCESS | < 5 s ✅ |
| hero p95 — manager | **1.51 s** (p50 0.72 s), 5/5 ACCESS | < 5 s ✅ |

**The cold stack is FASTER than the warm one it replaced** (1.22 / 1.51 s vs iter-09's 3.15 / 2.71 s).
Worth stating rather than glossing: the previous stack had been up for hours across three cockpit restarts
and two re-pins, and the fresh one runs clean caches and this run's own frontend images. The honest
reading is that the earlier numbers were the *pessimistic* ones and both are comfortably inside budget —
not that anything was optimized.

## Step 5 — the fourth component: 0 platform-repo edits

Every Go service clone and `next-web-app` is **git-clean**:

```
platform: 0 · next-web-app: 0 · app: 0 · jobsimulation: 0 · skillpath: 0
sentinel: 0 · storage: 0 · graphql-wundergraph: 0
cms: 1  → `?? studio/`  (untracked — the anthropos-studio-room clone `make init-studio` creates)
ant-academy: 4 → next.config.js · serverTenant.js · public/catalog.json · public/content/index.md
```

`ant-academy`'s four are the **sanctioned live-patch mechanism**, not edits: `serverTenant.js` is the M230
`academy-fs-published-fallback`, `next.config.js` the M212–215 `allowedDevOrigins` tailnet patch, and the
two `public/` files are build-generated. Both patches are applied to the demo's **ephemeral, gitignored
clone** before `next dev` and reverted on `--stop`; the bring-up's own `demo-patches: all applied (none
refused, none skipped)` covers them. **The canonical `anthropos-work` repos were never touched.**

## The gate

| component | status |
|---|---|
| all landable (session × action) pairs render real content, both vantages | **MET** — 29/29 |
| the academy grid renders real cards (Thread A) | **MET** — 65 cards, 0 Draft chips |
| p95 click→ACCESS < 5 s, HERO vantages (B2) | **MET** — 1.22 s / 1.51 s |
| reproducible on a cold reset-to-seed | **MET** — everything above, on a stack built from nothing |
| 0 platform edits | **MET** — canonical repos untouched |

**THE EXIT GATE IS MET.**

## Close — 2026-07-20

**Outcome:** The whole chain reproduces from nothing. A purge + cold bring-up at the final tooling tag,
with no manual intervention, yields 29/29 pairs, 65 academy cards with 0 Draft chips, and hero p95 of
1.22 s / 1.51 s — **every gate component MET**.
**Type:** tik
**Status:** closed-fixed
**Gate:** MET
**Phase 5 grading:** (1) gate-met: **y** — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (4 tiks) — (6) protocol-stop: n — Outcome: exit-1
**Decisions:** D1 (reproduced cold with zero intervention — the gate's reproducibility clause is satisfied by construction, not by narration), D2 (the cold stack measured faster than the warm one; report both rather than the flattering one), D3 (0-platform-edits verified per-clone, with the two patched files in the ephemeral academy clone named and justified rather than rounded to zero)
**Side-deliverables:** the protocol-evolution backfill owed by iters 05–09 — `coverage-protocol.md` (four render shapes → six, the two page-looks-broken traps, the argued 31→29 denominator, and a new "a green unit test can defend a broken path" subsection), `content-stories-spec.md` (the two corrected product-matrix rows + the membership-id silent-failure signature), and `latency-budget.md` (F12: the remote recipe was missing `LATENCY_SCHEME=https` — wrong for the exact scenario the section is about — plus the `STACK_DIR` requirement and the age-check timezone bug).
**Routes carried forward (none gate-blocking):**
- **M230 carry-forward cluster 3** — the anonymous `/library` + `/free` academy routes still render 0 cards (`getPublicCatalogView`'s `new Set()` branch is uncovered by the M230 patch, which names the gap itself). Handler `ACADEMY-M236-iter08-public-catalog-twin` → v2.5 release close.
- **`apps/web` client GraphQL endpoint** points at the non-offset `localhost:5050` while `apps/hiring` carries the correct offset origin. Never manifested on any measured path (SSR uses `WUNDERGRAPH_SSR_ENDPOINT`). Demo-hygiene → release close.
- **Standing carry:** 14 pre-existing demo-stack test failures (REPEAT) → release close.
- **The remaining v2.4-era method docs** flagged YELLOW by the milestone's KB-fidelity audit (F3/F4/F6/F7/F10 — `verification.md`, `demo-up-defaults.md`, `tailscale-serve.md`, `session-clone-spec.md`). `coverage-protocol.md`, `playthroughs.md`, `content-stories-spec.md` and `latency-budget.md` are now backfilled; the rest → release close.
**Lessons:**
- **The cold cycle is what converts "measured" into "true".** Everything before this iter was taken on a stack mutated in place across four iters. That the readings held — with the corrected manifest arriving from published tooling and no hand-editing — is what makes the other three components mean anything. A gate component that validates the others should be run last, and *must* be run.
- **Predictions recorded in a close are cheap and worth making.** iter-07 predicted the `127.0.0.1` bind was an artifact that a cold bring-up would correct; it was, and that took one `ss` line to confirm instead of an investigation.
- **Report the unflattering comparison.** The cold stack measured ~2× faster than the warm one. Saying so costs nothing and prevents the next reader from treating 1.2 s as the expected steady-state number.
