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
