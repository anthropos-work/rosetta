# iter-07 — tik: the skill-path manager "hang"

**Type:** tik  ·  **Active strategy:** `TOK-01` (publish-then-prove), Phase L.

## Step 1 — instrument before believing the hypothesis

iter-06 handed this iter a hypothesis: heavy 13-chapter path hangs, 3-chapter sibling passes, therefore a
**per-item fan-out** (`latency-budget.md`'s documented signature). The protocol's own rule is to name the
stalling leg before reading product code, so the first act was a navigation probe
(`stack-verify/e2e/tests/probe-navigation.spec.ts`): navigate with `commit`, record every request's wall
time, and dump everything still in flight at the deadline.

The hypothesis was **falsified immediately**:

```
PROBE: handshake+commit took 361 ms
PROBE: 134 completed legs — slowest 15:  ... 761 ms (static chunks)
PROBE: 0 STILL PENDING after 90000 ms
PROBE: 9 graphql/api legs — slowest 133 ms; getSkillPathDetails 66 ms
```

Nothing was in flight. Nothing was slow. Every GraphQL call finished inside 133 ms and the page was fully
painted in about a second. **There was no fan-out and no hang.**

## Step 2 — two root causes, neither of them the product being slow

### (a) The harness was waiting for `networkidle`

`content-stories.spec` called `loginAs` without a `waitUntil`, so it took the helper's `networkidle`
default. next-web holds long-poll connections open — `cockpit-login.ts` says so in its own doc comment, and
`latency-budget.md` states the rule outright: **never gate on `networkidle`**. On the enterprise
activity-dashboard it never resolves at all, so `page.goto` burned the full 180 s test timeout and the pair
was recorded as `threw: page.goto: Test timeout of 180000ms exceeded` — which reads exactly like a product
hang. The lighter sibling squeaked under the timeout, manufacturing the heavy-vs-light contrast that made
the fan-out story so persuasive.

The probe also exposed a second reason the page looked empty: `main.innerText` is **`""`** on a
fully-populated page, because antd `Drawer` mounts through a **portal** outside `<main>`.

### (b) The skill-path manager surface does not exist

With navigation fixed the page renders — but reading what it renders is the real finding.
`InsightsBySkillPathStudentSimulationsContainer.tsx` (next-web) contains:

```tsx
const userData = useMemo(() => {
  // return insightData?.rows[0]?.membership;
  return null as unknown as MembershipEnriched;
}, []);
...
<Typography.Title level={1}>{t('enterprise.insights.comingSoon')}</Typography.Title>
{/* <Table tableKey='insight-table-students' ... /> */}
```

The results table is **commented out**, `userData` is **hardcoded null**, and the body renders the literal
string **"Coming soon"**. The only populated query is `getSkillPathDetails` — the *path definition*. **No
query touches the seeded session.** The page is byte-identical whether or not anything was ever seeded.

Live confirmation on both siblings:

```
sp-genai-in-progress        drawer 170 chars: "Results for Practical introduction to GenAI ...
                            2 skills measured: Prompt Engineering, Clear Messaging  Coming soon"
sp-product-manager-completed drawer 784 chars: same shape, 25-skill list, same "Coming soon"
```

**The sibling that had been counted as PASSING was a false pass.** The `manager-dashboard` shape asserts on
`/results for/i` — which appears in the definition-only header — and on the absence of `"No data"`, which a
"Coming soon" placeholder also satisfies. This is precisely the trap iter-05 documented on the simulation
scoreboard ("the header comes from a different query, so it looked populated while proving nothing"),
recurring on a different surface.

## Step 3 — the fix (tooling only, 0 platform edits)

1. **`content-stories.spec`** navigates with `waitUntil: 'commit'`.
2. **`ContentResultPage.settle()`** polls for content to stop growing (past a 200-char floor) instead of
   sleeping a flat 8 s, and measures **`main` + `body`** so portal-mounted surfaces are seen. Required,
   because `commit` returns before anything is painted.
3. **`skill-path-legacy` → `managerKind: ""`** in the content-product registry — the same player-link-only
   disposition academy already carries. M233's rule is fail-closed: *a session that cannot form a real link
   is dropped with a reason, never linked anyway*. A CTA onto "Coming soon" is a fabricated CTA.
4. **The test that asserted the defect** (`content_nonsim_test.go` required the manager CTA) now guards the
   other way: restoring `managerKind` without the platform surface fails loudly.
5. The **M233 honesty gate caught the stale canonical preset**, as designed; regenerated.

Published as rext `playbill-m236-skillpath-nomanager`; `billion` re-pinned, manifest re-exported live
(0 manager views on skill-path-legacy, verified on the SERVED file), cockpit restarted.

## Step 4 — re-measure

```
content-stories: LANDED 28 / 29
  simulation:        26/26
  skill-path-legacy:  2/2
  skill-path-new:     0/1
  x1  route rendered a not-found  — academy
```

**The denominator moved, and it is important to be exact about why.** 31 was never a count of provable
pairs: 2 of the 31 were manager actions onto a surface the platform has not built. Removing them is the
same rule that already excludes ai-labs (no seedable result surface → not landable), applied to newly
obtained evidence, not a relaxation of the gate.

| | before | after |
|---|---|---|
| denominator | 31 | **29** (−2 unlandable skill-path manager pairs) |
| numerator | 29 | **28** (−1 false pass, −1 phantom "hang" that was never a real pair) |
| gap to gate | 2 | **1** |

The one remaining pair is academy, unchanged and never in doubt.

## Close — 2026-07-20

**Outcome:** The "hang" was the harness waiting on `networkidle`, and the surface behind it is an
unimplemented "Coming soon" placeholder — so the pair was never landable and its sibling had been passing
falsely. Gate distance 2 → **1**; the honest reading is **28/29**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (the `networkidle` default was the hang; navigate on `commit`, settle by polling main+body), D2 (skill-path has no manager surface → player-link-only; denominator 31→29 with product-source evidence), D3 (the passing sibling was a false pass — the iter-05 header trap on a second surface)
**Side-deliverables:** `probe-navigation.spec.ts` — a reusable "what is actually in flight" diagnostic, generalized from this iter's probe. Not swept by any runner.
**Routes carried forward:**
- **Academy** (1 pair, the last one) → handler `ACADEMY-M236-iterTBD-catalog-fill`: wire `app/cmd/academy-seed` into the cold bring-up, then re-point the CTA from the anonymous `/library/<slug>` preview to the authed chapter route.
- **p95 click→ACCESS, HERO vantages only** (B2) → handler `LATENCY-M236-iterTBD-hero-p95`.
- **Final cold reset-to-seed reproduction** → handler `REPRO-M236-iterTBD-cold-cycle`. Note for that iter: the cockpit currently binds **127.0.0.1** rather than `0.0.0.0`, because `tailscale serve` had already claimed `:17700` by restart time. A cold bring-up restores the normal ordering; the tailnet path is unaffected either way (serve proxies to loopback).
**Lessons:**
- **The milestone's running score is now four wrong assertions to one real product bug** — and this iter produced both a false FAIL and a false PASS from the *same* mis-calibrated shape. The false pass is the dangerous half: a gate that passes for the wrong reason hides a missing surface indefinitely.
- **A documented rule was already in the corpus and was still violated.** `latency-budget.md` says never gate on `networkidle`; `cockpit-login.ts` repeats it in the doc comment of the very default that caused this. Writing the rule down is not the same as making it hard to get wrong — hence `waitUntil` is now explicit at the call site with the cost recorded inline.
- **When a page looks empty, check for portals before blaming data.** `main.innerText === ""` on a fully-rendered page is a modal/drawer signature, not an empty one.
- **"Heavy instance hangs, light sibling passes" is not sufficient evidence of a fan-out.** A fixed timeout plus any noisy wait manufactures the same contrast. Confirm with in-flight request counts before accepting the signature.
