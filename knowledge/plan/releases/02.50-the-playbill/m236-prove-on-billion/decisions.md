# M236 — Decisions

## TOK-01: publish-then-prove — 2026-07-20

**Tok type:** bootstrap (iter-01)

**Initial strategy:** four ordered phases, the ordering forced by the live baseline rather than chosen.

- **Phase P — publish (iter-02).** Push `rosetta-extensions` `main` (fast-forward, 20 commits) + the 13
  `playbill-*` tags to origin; re-pin `billion`'s `.agentspace/rext.tag` to `playbill-m235-hardened` and
  check its consumption clone out at that tag; prune the host's 107.6 GB reclaimable build cache first.
- **Phase C — cold bring-up.** A cold reset-to-seed `/demo-up` on `billion` at the new tag, public-host
  default-on (D-DESIGN-3). Success = stack UP, fresh-green `autoverify.json`, cockpit serving a non-404
  `content-manifest.json` with all 4 products.
- **Phase L — land the arms, one product arm per iter,** in descending evidence-density order:
  `simulation` (26 actions) → `skill-path-legacy` (4) → `skill-path-new`/academy (1) → `ai-labs`
  (0 landable; prove the 2 presence rows). Per iter: measure the arm, triage every non-landing pair to a
  root cause, fix in seeder/manifest, re-measure.
- **Phase H — harness + budget, interleaved.** The new content-stories seat-login coverage plumbing is
  authored **against the first live seeded render**, not up front; the p95 click→ACCESS measurement rides
  `rext stack-verify/e2e/run-latency.sh` once ≥1 arm lands.

**Rationale:** the baseline measurement (iter-01 B1–B5) turned up a structural blocker the milestone plan
never named — `billion` **cannot obtain the feature under test**. It consumes `rosetta-extensions` only at
a tag named in `.agentspace/rext.tag`, guarded FATALly since M217; it pins the M228 tag; `origin/main` **is**
that M228 commit; and **0 of 13** `playbill-*` tags are on origin. Every v2.5 milestone's tooling
(M230/M232/M233/M234/M235) lives only in the local authoring copy. So publish-first is not preferred
sequencing — it is the only order in which any other phase is testable.

The rest of the ordering follows from evidence density and from M235's Cluster-1 finding: one arm per iter
keeps the scope-creep tripwire enforceable across 4 arms with genuinely different failure modes (fixture
replay vs. version-match vs. catalog-fill vs. DDL), and harness-after-first-render is forced by
USER-BLOCKER-M235-02 (authoring the descriptor blind ships an *incorrect*, not merely uncalibrated,
load-bearing harness).

**Strategy class:** new-direction — bootstrap tok; no prior strategy exists to compare against.

**Distance-to-gate context:** primary metric **0 / 31** landable (session × action) pairs proven live;
all secondary components (academy real cards, ai-labs presence rows, p95 measurement) also at 0. The
distance is dominated by a **structural reachability gap**, not by defect count — so iter-02 is expected to
move *reachability* with a **0** metric delta, and the first numerator movement comes in Phase L.

**Next-tik direction:** iter-02 executes Phase P — prune `billion`'s build cache, push `main` + the 13
`playbill-*` tags, re-pin `.agentspace/rext.tag` to `playbill-m235-hardened`, check the consumption clone
out at it, and verify the M217 FATAL pin guard passes. Expected lift **0**; iter-02 must close honestly as
`closed-fixed` on planned scope with `Gate: NOT MET`. If push is refused for an access or policy reason,
that is a genuine user-blocker and exits the session.

---

## USER-BLOCKER-M236-01: the exit gate contains clauses that cannot be proven or measured — 2026-07-20

**Raised:** iter-02, before its publish step executed. **Source:** Phase 0b `audit-kb-fidelity` → **RED**
([`kb-fidelity-audit.md`](kb-fidelity-audit.md); summary in [`spec-notes.md`](spec-notes.md)).

M236's `overview.md` exit gate reads:

> Both tabs work live on billion — Content-stories sessions render real content for player + manager
> vantages, the academy grid renders real cards (Thread A) — reproducibly on a cold reset-to-seed,
> **p95 click→ACCESS < 5 s**, 0 platform edits, **demo reachable only over the tailnet**.

Two of its clauses cannot be discharged as written, and one cited mechanism does not exist. This is a
**scope/spec decision the user owns** — not something an iter can resolve by working harder.

### B1 — "reachable only over the tailnet" is false by construction

`safety.md:405` states every demo container is published on `0.0.0.0` — all interfaces — on **every**
`demo-up`, flag or no flag (measured at `:413`: 14 ports → `0.0.0.0`). `tailscale-serve.md:626-627` adds
that on Linux this **bypasses the host firewall**, so a `ufw deny` does not block it. The clause is
therefore only ever true as a property of `billion`'s **network placement**, and is only demonstrable via
an explicit **off-tailnet probe**.

**Options:** (a) re-word the gate to "no reachable route from off-tailnet, demonstrated by probe" and add
the probe as a deliverable; (b) drop the clause and rely on `safety.md` §3 Part 3's existing disclosure;
(c) keep it and accept the milestone cannot honestly claim it.

### B2 — "p95 click→ACCESS < 5 s" is unmeasurable for all 31 content-story actions

The content-player CTA emits **no `data-login-as`** attribute (`cockpit.py:421-425`) — and that attribute
**is** the ACCESS predicate: `latency.ts:123-127` throws without it, *before t0*. Separately,
`run-latency.sh:42-47` hard-rejects non-hero vantages (`exit 2`). So today the metric cannot be taken for
any content-story action at all.

**Options:** (a) extend the cockpit CTA + latency harness to content seats (real work, enlarges the
milestone); (b) scope the p95 gate to the existing **hero** vantages only, where it is already measurable,
and declare content-seat latency out of scope for v2.5; (c) drop the clause.

### B3 — Cluster 1's cited page-object does not exist

`overview.md:37` says the new harness "reuses the shared `AISimulationResultContainer`". There are **zero**
hits in `rosetta-extensions` — it is a **next-web `.tsx` component** (`content-stories-routes.md:52`), not a
harness object, and the nearest harness object (`simulation-page.ts:9-10`) stops at the *launch* boundary
with no `/result/` locator. The harness must be **authored from scratch**, and
`VantageManifest.identityKey` being **singular** means **13 seats → 13 manifests**, not one sweep.
Cluster 1 is materially larger than M235's carry-forward described.

### B4 — M236 must consciously reverse a documented rule

`coverage-protocol.md:421-431` mandates **excluding** the exact pages M236 exists to prove — `skipPaths`
contains `/\/result\/[0-9a-f-]{8,}/`. Reversing a documented protocol rule is a spec decision, and the
protocol doc must be amended in the same change (per the protocol-evolution rule).

### B5 — the declared `iteration_protocol_ref` is hollow

`verification.md` (335 lines) has no measure→triage→fix loop and no gate; it supplies a gate *input*, not
a protocol. And **all six** method docs — including both of M236's declared protocol refs — contain
**zero** mentions of content-stories / `content_products` / `content-manifest` / `content-player`. M236's
declared method describes a **v2.4 world**.

**Recommendation (agent's view, user decides):** re-scope the gate to what is provable and measurable —
adopt B1(a), B2(b), accept B3's enlarged Cluster 1 as the milestone's real centre of gravity, take B4 as an
explicit amendment to `coverage-protocol.md`, and repoint `iteration_protocol_ref` at
`coverage-protocol.md` + `playthroughs.md` (the docs that actually carry a measure→triage→fix loop) while
backfilling their content-stories sections. Then resume Phase P unchanged — the publish is verified safe
and ready.

**Not blocked by this:** the publish itself (verified: 13 tags ancestors of `main`, clean fast-forward, 0
collisions, 16/16 packages green) and the completed host prune (109 GB reclaimed, `billion` 40 G → 139 G
free).

### **RESOLUTION** (2026-07-20, user-authorized)

The orchestrator obtained the user's decision interactively. All five sub-findings are **decided and
closed**; none may be re-raised.

- **B1 — DROP THE CLAUSE.** Security is not a concern for this milestone; the demo may be publicly
  accessible. Making it reachable by the *right* people is the job of the VM + VPN, **not** of the demo
  stack — the demo stack's only obligation is to *permit* VPN access. **No off-tailnet probe deliverable.**
  The clause is removed from the exit gate. `safety.md` §3 Part 3's existing disclosure stands **as-is** and
  needs no amendment for this.
- **B2 — OPTION (b): scope the p95 gate to the existing HERO vantages only,** where `run-latency.sh`
  already works. **Content-seat latency is explicitly OUT OF SCOPE for v2.5.** Do **not** extend the cockpit
  CTA or `run-latency.sh` to content seats. The 31 content actions are proven for **CONTENT** (they render
  real, non-empty results), not formally timed.
- **B3 — ACCEPT the enlarged Cluster 1** (the content-stories seat-login coverage harness authored from
  scratch, **13 seats → 13 manifests**) as the milestone's real centre of gravity.
- **B4 — AMEND `corpus/ops/demo/coverage-protocol.md` explicitly** in the same change that reverses its
  `skipPaths` `/result/` exclusion, per the protocol-evolution rule.
- **B5 — REPOINT `iteration_protocol_ref`** at `corpus/ops/demo/coverage-protocol.md` +
  `corpus/ops/demo/playthroughs.md` (the docs that actually carry a measure→triage→fix loop), and backfill
  their content-stories sections.

**Publish authorized.** The user was told iter-02 will push the 13 `playbill-*` tags + `main` of
`rosetta-extensions` to origin and re-pin billion's `rext.tag`, and did not object. Phase P proceeds as
planned (verified safe by iter-02's pre-flight).

**Gate denominator preserved: 31 landable (session × action) pairs** — 26 simulation + 4
skill-path-legacy + 1 academy; ai-labs is presence-only. `has_manager_view` is **per-SESSION, not
per-product** — a product-level read silently under-counts 31 → 18.

---

## Close review — Adversarial review (Phase 2c)

Scenarios considered against the milestone's own code, and their responses.

**S1 — the aggregator is handed a not-ok row with no `reason`.** REAL. The malformed-row guard validated
`ok` and `product`; the failure summary indexes `reason`. The row raised `KeyError` **after**
`content-stories.json` had already been written — "crashed but left a report", where the artifact outlives
the run that produced it and is later read as evidence. That is D17's signature hazard inside the gate's own
aggregator. **Fixed + regression-tested**, with a negative counterpart so the fix cannot over-correct into
rejecting every not-ok row.

**S2 — a result page renders entirely through a portal.** REAL-but-latent. `settle()` counts `main + body`
to see antd-Drawer surfaces, but the length floors read `main.length`, so such a page could settle correctly
and then fail "too short (0 chars)". **Fixed** — but note the *second-order* trap found while fixing it:
the obvious repair (`readable = main.length + body.length`) **double-counts**, because `<main>` is a
descendant of `<body>`. That silently HALVES every floor and would let a blank page carrying only nav chrome
clear a 300-char gate — reintroducing the false-PASS class the milestone spent six iters eliminating. The
landed fix is `main || body`. *A fix for a false-FAIL that creates a false-PASS is a net loss.*

**S3 — the sweep is run with a mistyped stack number.** REAL for two of four runners. `$(( abc * 10000 ))`
is 0 in bash, silently sweeping the DEV stack. **Fixed** across all four (see S6).

**S4 — the sweep runs a SUBSET and reports green.** REAL. `EXPECTED_PAIRS` is the only guard against
"ran 26 of 29 and landed all 26", and it was documented in two places and **set by nothing** — dark by
default. **Fixed**: pinned from the SERVED manifest, computed by mirroring `buildPairs` exactly, failing
loud rather than pinning 0.

**S5 — a stray argument widens the run.** REAL. Positionals were forwarded into `playwright test` as a
second filename filter. **Fixed** (refused). This is the M50 bug that `run-coverage.sh` already documents at
length — the new runner simply did not adopt the guard.

**S6 — a lesson is written in one file and not its siblings.** REAL, and the milestone's most recurrent
meta-failure: it happened with the `N` guard (2 of 4 runners), with `networkidle` (documented in
`latency-budget.md`, then inherited anyway by the new sweep), and with the doc corrections (iters fixed the
docs they touched, not the docs carrying the same claim). Every instance is fixed; the pattern is carried
into the retro because it is not fixable by any single fix.

**S7 — the harness changed after the gate was proven.** ACCEPTED, disclosed. See CLOSE-D3 below.

---

## CLOSE-D1 — deferral audit returned **RED**; item 4 requires a USER fate — 2026-07-20

`/developer-kit:audit-deferrals --scope=milestone` returned **RED / SEVERITY=blocker**. 9 items in scope.
Eight are dispositioned below and were LANDED or correctly routed at this close. **One cannot be
dispositioned by an agent** and is escalated to the user; the milestone therefore does **NOT** merge in this
session.

| # | Item | Fate at this close |
|---|---|---|
| 1 | `ACADEMY-M236-iter08-public-catalog-twin` (anon `/library` + `/free` render 0 cards) | **Fate 3 → v2.5 release close.** Needs a 2nd demopatch manifest + a next-web rebuild + a live re-prove — not agent-completable in a docs close. `frontend-tier.md`'s unqualified "the empty grid is FILLED" was **scoped** to signed-in at this close, so the doc no longer over-claims. |
| 2 | `apps/web` client GraphQL endpoint on non-offset `:5050` | **Fate 3 → v2.5 release close**, batched with #1 (same rebuild, one re-prove). |
| 3 | v2.4-era method docs (F3/F4/F6/F7/F10) | **LANDED (Fate 1).** `verification.md` + `tailscale-serve.md` + `demo-up-defaults.md` backfilled; `session-clone-spec.md` already had coverage. |
| 4 | **14 pre-existing demo-stack test failures** | **ESCALATED — see CLOSE-D2. No agent-side fate is legitimate here.** |
| 5 | `run-coverage.sh` / `run-hiring-render.sh` non-integer-`N` guard | **LANDED (Fate 1)**, committed AND published (see CLOSE-D4). |
| 6a | Seed ships 4 orgs, docs say 3 | **LANDED (Fate 1)** — 20 places across 8 corpus docs, 2 shipped skills, and `CLAUDE.md`. |
| 6b | `test_m220_mutation_battery.py` unmutated subject fails | **Folded into #4** — which means the standing set is **15+, not 14** (see CLOSE-D2). |
| 7 | `DEF-M235-03` M204 assign-WRITE declared TODO | **Fate 3 → v2.5 release close** (inherited, correctly routed past M236). |
| 8 | M230 carry-forward **cluster 2** (next-web clone re-anchor) | **DISCHARGED** — recorded in `overview.md` at this close. It was the one forward-routed item with **no closing entry anywhere**; the diagnosis was falsified (clone was not drifted) and the cold bring-up passed. |

Also landed from the audit's queue: the **`DEMO_NO_ACADEMY_FILL` knob** — undocumented while **gating
Thread A**. `stack-core/demo_knob_guard.py` was in **disagreement** before this close and now passes both
directions.

## CLOSE-D2 — the 14-failure carry is a genuine repeat-deferral; escalated, not renewed — 2026-07-20

**This is the blocker.** The audit traced the item across **10 distinct milestones and 2 releases** (first
seen M224, 2026-07-17, as *8* failures). Three facts make it an escalation rather than another YELLOW:

1. **Its declared destination has already fired once without landing it.** M225–M227 routed it explicitly to
   *"the v2.4 release close."* v2.4 closed 2026-07-18 and the item shipped as a known issue, re-anchored on
   v2.5's close. That is a verbatim AGED_OUT trigger, and **no audit in the tree records it firing**. The
   argument used at M235 to justify non-escalation — *"the destination has not yet closed, so the deferral
   authority is intact"* — was equally true of v2.4's close, right up until it wasn't.
2. **M236 is the FINAL v2.5 milestone.** There is no further milestone to defer into; the next event *is*
   the release close. It cannot be renewed silently — only landed, dropped, or deferred cross-release with
   explicit sign-off.
3. **The item drifted materially under a fixed label.** The count went **8 → 14** at the v2.4 close with no
   document treating growth as requiring re-decision, and the *nature* changed too: M224 called them stale
   tests for intentionally-removed behaviour (semantic debt); M232 D8 diagnoses 6 of them as `pre_sha256`
   **pin drift** (env debt) — a different class. It is growing again: this milestone's hardening ledger adds
   the mutation battery and the org-count guard as "belonging to the set" while not re-routing them. So
   **14 is now wrong in both directions** and must be re-baselined before any fate is chosen.

Supporting evidence that this is a known local hazard: finding **F-M220-6** states the lesson verbatim —
*"A red suite is not just noise … '14 failures' is indistinguishable from '11 + 3' until you baseline
against HEAD."* That lesson was landed Fate-1 at M220 and the same hazard has since re-accumulated.

**Why no agent-side fate was taken.** LAND-NOW is the three-fate default, and 6 of the failures are
mechanical pin re-anchors — but re-anchoring `pre_sha256` requires a live next-web clone, and the set is not
correctly enumerated today. KEEP-DEFERRED-WITH-SIGNOFF requires a **fresh dated user decision**, which an
agent cannot manufacture. Choosing either unilaterally would be exactly the silent renewal that made this a
repeat-deferral in the first place. **Required before v2.5 can close: re-baseline the set against HEAD, then
an explicit user fate (LAND-NOW / DROP / KEEP-DEFERRED-WITH-SIGNOFF).**

### **RESOLUTION** (2026-07-20, user-authorized)

**Fate chosen by the user: RE-BASELINE NOW, fate decided at the v2.5 release close.** Explicitly *not*
LAND-NOW (do not fix them in this milestone), *not* DROP, and *not* another silent roll-forward. The
re-baseline is the deliverable of this close; the LAND/DROP/DEFER decision is handed to
`/developer-kit:close-release` with a real, dated characterisation attached instead of a stale count.

**The user attached a measurement condition, verbatim:**

> "make sure to bring up a fresh demo (if you didn't) .. meaning all repo of the stack should point to their
> stable main... i still see an old next-web-app left menu. if no stable repo is visible, use directly main."

**Why the condition is technically load-bearing, not ceremonial.** CLOSE-D2 §3 diagnoses 6 of the failures as
`pre_sha256` **pin drift** — and a demopatch `pre_sha256` is a hash of a *platform source file in the demo's
own clone*. Measuring pin drift on a stack built from stale clones re-measures the staleness, not the
failures. A stable-`main` clone set is therefore a **precondition for the reading to mean anything**, and the
re-baseline below is measured against one.

**Ref-selection rule applied:** every platform repo pinned to its stable `main`; no repo exposes a distinct
"stable" ref, so `main` was used directly for all of them. The exact resolved repo → branch → sha set is
recorded in the re-baseline artifact so the reading is reproducible.

**Sub-item — the stale left-menu symptom.** The user reports still seeing an **old next-web-app left menu**
on the demo. Carried as a **named finding to be root-caused**, not merely as a motivation to rebuild. Result:
**F-M236-CLOSE-1** (below) — the symptom is real evidence of a real defect, but *not* of the defect it looks
like.

## F-M236-CLOSE-1 — `/demo-up` rebuilds images from clones it never updates — 2026-07-20

Root-causing the stale-left-menu report produced a finding that is **more general than the symptom**, and one
no existing doc states.

**The defect.** `ensure-clones.sh` populates a demo's clone set via `make init`, described in its own header
comment as *"the canonical, idempotent clone loop … **skip-if-present**"*. There is **no fetch, no pull, and
no checkout step for the platform repos anywhere in the bring-up.** Once `stack-demo/<repo>` exists, it is
never advanced again. Every subsequent `/demo-up` — including a full cold teardown-and-rebuild — recompiles
**fresh images from stale source**. The clone set on `billion` was created 2026-07-03…07 and had never moved.

**The bring-up log says so in the product's own words.** `coldup-m236.log`, the M236 cold run, contains for
every single repo:

```
Cloning missing repos...
  app already exists, skipping
  cms already exists, skipping
  … next-web-app already exists, skipping …
Done. All repos are available in /home/devops/panorama/stack-demo/
```

*"All repos are available"* is true and *"all repos are current"* is what a reader takes from it. Those are
not the same claim, and nothing in the bring-up ever checks the second.

**Measured drift** (both boxes, 2026-07-20, against a *verified* fetch — see the methodology warning below).
Images on `billion` were all built 2026-07-20 07:54–10:11, i.e. genuinely fresh **images** compiled from
**up-to-13-day-old source**:

| repo | clone head | `origin/main` | behind |
|---|---|---|---|
| `app` | `c3c45e01e` (07-07) | `aa2574541` (07-20) | **249** |
| `next-web-app` | `23bdbb5db` (07-07) | `61d72e24d` (07-20) | **202** |
| `ant-academy` (local) | `2c6e0682c` (06-25) | `a43420bdd` (07-17) | **60** |
| `messenger` | `57282400b` (06-22) | `d41029217` (07-20) | **28** |
| `graphql-wundergraph` | `c284453b4` (07-03) | `5d9c7568e` (07-17) | **6** |
| `cms` | `770ec3aac` (07-03) | `93e6aa354` (07-17) | **4** |
| `jobsimulation`, `platform`, `roadrunner`, `sentinel`, `skillpath`, `storage`, `studio-desk` | — | — | 0 |

Both boxes showed **identical** drift, so this is a property of the mechanism, not of one machine.

**⚠ Methodology warning — how this was nearly mis-diagnosed.** The first pass measured drift on `billion` and
read **12** commits for `next-web-app` (not 202) and **20** for `app` (not 249). Cause: `git fetch` as `root`
there dies with *"Host key verification failed"* (root has no GitHub host key; the tree is `devops`-owned),
and the survey ran it with `2>/dev/null`. The fetch failed **silently**, so `origin/main` resolved to a
remote-tracking ref last updated days earlier, and the comparison **measured stale-vs-stale**. That produced
a confident, wrong conclusion (below). **Never measure drift through a suppressed-stderr fetch** — assert the
fetch succeeded, or the reading is worthless.

**The wrong conclusion it produced, recorded because the near-miss is the lesson.** Against the bad refs,
`git diff HEAD origin/main -- packages/ui/src/NavBar/` was **empty**, and `navbarMenuItems.tsx` last changed
2025-08-26 — so the first pass concluded the menu was *not* explained by the drift and must be runtime
(role/flag) determined. Against the **true** refs the same diff is **not** empty: `NavbarItems.tsx`,
`NavbarTop.tsx`, `OrganizationNav.tsx` and `navbarStyles.ts` — the actual left-menu components — **all changed
in the 202-commit gap.** Only `navbarMenuItems.tsx` (the per-item *renderer*) was unchanged, which is exactly
the file the first pass happened to check.

**So: the stale left menu IS the stale clone.** The user's instinct was right and their prescribed remedy —
point every repo at stable `main` — was the correct one.

**A real secondary effect, still worth keeping.** The nav is also assembled at runtime from `isEnterprise` /
`isAdmin` / `role` / `organizationSettings` / `useTalkToDataAccess()` / `posthog.isFeatureEnabled(...)`, and a
demo has **no PostHog**, so flag-gated entries evaluate false. A demo menu is therefore legitimately a
*subset* of what the same commit renders in production. This does **not** explain the reported symptom, but it
means a stable-`main` demo menu still will not match production exactly — expected, not a regression.

**The provenance record cannot detect any of this.** `clones.lock.json` stores `ref` as
`git rev-parse --abbrev-ref HEAD`, which for these detached clones is the literal string **`"HEAD"`** — for
*every* repo. The lockfile faithfully records the sha while being structurally incapable of distinguishing
*"deliberately pinned"* from *"stale by neglect"*. A reviewer reading it sees provenance and infers freshness;
the file does not carry the fact needed to refute that.

## F-M236-CLOSE-2 — the "pristine" safety net sweeps 3 manifests out of ~15 — 2026-07-20

Found while root-causing the re-baseline. `ensure-clones.sh`'s **R1** rung reverts *"any stale patch left by a
crashed prior build"* via `demopatch revert --force-pristine` — the mechanism that guarantees a demo clone is
pristine before a build. Its manifest list is hard-coded and contains **three** entries:
`next-web-studio-url`, `next-web-members-pagination`, `app-targetrole-authz-skip`. The `patches/` directory
carries **15**.

Consequence: a patch outside those three, left applied by an interrupted build, is **never swept** — it
survives every subsequent `/demo-up`, including a full cold teardown-and-rebuild. Both boxes were carrying
exactly that, in disjoint sets:

- **local workspace** — 5 leftover applied patches in `next-web-app` (`server.graphql.ts`,
  `useAiReadinessActive.ts`, `layout.tsx`, `urls.ts`, `InsightsContext.tsx`). Only the last two are in R1's
  list; the other three had persisted indefinitely.
- **`billion`** — `next-web-app` clean, but `ant-academy` carrying leftover `ant-academy-dev-origins`
  (`code/next.config.js`) and `academy-fs-published-fallback` (`code/src/lib/serverTenant.js`). Neither is in
  R1's list.

This is the mechanism that generated the entire clone-dependent half of the 14-failure set (see the
re-baseline artifact). Both were reverted via the tool's own `revert --force-pristine` — never a raw
`git checkout`.

## CLOSE-D3 — the harness changed after the gate was proven; the reading is NOT re-proven live — 2026-07-20

The gate (29/29 cold on `billion`) was measured with the harness at `playbill-m236-hardened`. This close then
fixed 3 must-fix and 7 should-fix defects **in that same harness**, including two that alter grading
semantics: the length floors (`main` → `main || body`) and the not-found guard (`main` → `main + body`).

These are verified by **66 harness unit specs** (up from 64, incl. 2 new regression tests) and by the
route-contract suite that independently asserts the 29-pair denominator from the canonical manifest — but
they have **not been re-run against a live stack**. The honest statement is: *the gate reading stands on the
harness as it was at the moment of measurement; the close fixes are unit-proven, not live-re-proven.*

Recorded rather than papered over, because this milestone's own central finding is that a check can report
success while proving nothing. A cheap live re-run of `run-content-stories.sh 1 --host billion...` at the
v2.5 release close would discharge it; it is **not** gate-invalidating (every change is either strictly more
conservative or covered by a regression test), but it is a real, named residual.

## CLOSE-D4 — the close fixes were PUBLISHED, not just committed — 2026-07-20

The audit caught that all close-review fixes were sitting **uncommitted** in the authoring copy, and that
`playbill-m236-hardened` did not contain them. Because `billion` consumes `rosetta-extensions` only at a tag
**fetched from origin** (M217 FATAL pin guard), that work was discharged in the working tree and undischarged
in the product — the *exact* failure mode iter-01 found and that `verification.md`'s new PRE-FLIGHT RUNG ZERO
now names (*"tagging is not publishing"*). Committed as `5c8d12e`, tagged **`playbill-m236-close-fixes`**,
and **pushed to origin**. `billion`'s pin is unchanged (the gate reading belongs to the tag it was taken on —
CLOSE-D3); re-pinning is a release-close action.
