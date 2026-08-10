---
iter: 276
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-10
handler: FIX-M257x-275-bound-hero-role-occupancy-to-exactly-one-peer
---

# iter-276 — bound hero-role occupancy to exactly one peer

**Type:** tik

## Step 0 — Re-survey before targeting (mandatory)

Re-ran the survey over `stack-seeding/seeders/` at rext `2833a64` before committing to iter-275's
hand-off. Three findings, and one of them **corrects the hand-off**:

1. **The target is still untouched.** `memberRoleAt` at `jobroleref.go:88-100` is verbatim the
   unbounded uniform draw iter-275 described: `idx := hashInt("<prefix>:role:<i>")`, `k := idx % len(set)`.
   No cap, no floor, no reservation. `orgRoleSet:126-157` still adds hero roles first and never
   reserves an occupant for them. The invariant in its own doc comment (`:67` — *"a hero is never the
   sole holder of her title"*) is still enforced nowhere.

2. **The 11-site radius RE-VERIFIED, and iter-275's own warning about it was worth heeding.**
   iter-275 flagged that this file's comment records a previous sweep that *"found only FOUR"* of six.
   Re-counted mechanically rather than trusted: **6 production + 5 test = 11**, matching. The six are
   `users.go:181`, `membership_skills.go:145`, `population_evidence.go:132`, `certificates.go:150`,
   `profile.go:323`, `target_roles.go:106`. **Five of the six pass `storyHeroRoleNames(st)` and have
   `st` in scope**; the sixth (`certificates.go:150`) is inside the forwarding helper
   `memberRoleName(prefix, storyHeroRoles, i, roles)`, which must be threaded.

3. **NET-NEW, not in the hand-off: a FENCE already guards this function.**
   `seeders/role_tenancy_fence_test.go` parses every non-test seeder source, recognises each
   `memberRoleAt` call site, and `t.Fatalf`s if it finds none ("the recogniser has broken"). It also
   knows about **forwarding helpers** (`:190` — an exemption written for exactly the `certificates.go`
   shape). **Any signature change must keep this fence green**, and the fence is a free check that the
   sweep reached all six rather than four.

**Target confirmed, no substitution.** iter-275's `FIX-M257x-275-bound-hero-role-occupancy-to-exactly-one-peer`
is still the right next thing and is still the single thing standing between clause 2 and green.

## Active strategy reference

`TOK-08` (*census the mechanical classes; stop sampling them*) is the milestone's active strategy and
governs **clause 5**. This tik is **clause 2** work under the **user's binding closing condition**
(`D-M257x-256-1`), which supersedes the gate-4-of-5 framing: the demo AND dev stacks must work and the
corpus must match. Clause 2 is one Playthrough from met and that Playthrough is this fix. TOK-08's
census discipline still applies to the measurement half of this iter (Priority 2).

## Cluster / target identified

Gate clause 2 reads **30 live / 1 failing / 0 error** (iter-273, binding, at the shipping pin). The
single failure is `workforce-intelligence.talent-pool.UC1` — `pt-workforce-succession`. iter-272 closed
its mechanism; iter-274 reduced it to role occupancy; iter-275 established the requirement is a
**bound with two broken tails**, not a number.

**The arithmetic, re-confirmed from iter-274's measured table (`riskScore = round(structuralRisk × 0.85)`):**

| incumbents (hero included) | structuralRisk | riskScore | vs the `RiskScore ≥ 50` guard |
|---:|---:|---:|---|
| 1 | 80 | **68** | over — but the hero **sole-holds her title**, violating `orgRoleSet:67` |
| **2** | **63** | **54** | **over by 4 — and the invariant holds. The only admissible value.** |
| 3–4 | 53 | 45 | under by 5 — both at-risk signals lost → the Playthrough fails |

Org A shows both tails failing at once: Pat Ellis (employee hero) at **3**, Morgan Reyes (manager hero)
at **1**.

## Hypothesis

A hero's role must carry **exactly one peer** (2 incumbents). Making that a **reserved, deterministic
assignment** rather than an outcome of a uniform hash draw:

- lifts `DevOps Engineer` 3 → 2 incumbents → `riskScore` 45 → 54 → clears the `≥ 50` guard → Pat's
  rare-skill (+25) and fragile-role (+15) signals both return → `talent-pool.UC1` passes;
- lifts `Engineering Manager` 1 → 2 → enforces the invariant that has been documented-but-unenforced
  since M257x iter-31.

**The peer must be a SUPPORTING member, not another hero** — hero indices are hashed into `[1, st.Size]`
by `personaIndexMapForStory`, so a reservation that ignores them lands on a hero ~32 % of the time in
Org A and yields zero peers. This is why the bound genuinely needs more than the current signature
carries, and confirms iter-275's radius reasoning rather than shortcutting it.

## Expected lift

- **Clause 2: 30 live / 1 failing → 30 live / 0 failing.** The gate's binding metric.
- The occupancy invariant becomes **enforced** (asserted by a new unit fence), not merely stated.

## Phase plan

- **A** — implement the reservation in `jobroleref.go`; thread the signature through all 6 production
  sites + the forwarding helper; update the 5 test references.
- **B** — unit gate: a new test that proves **exactly one peer per hero role** structurally (any pool,
  any population), plus a **RED-on-pre-repair-tree proof** — the standing failure mode is a direct
  anti-regression test that is green while the bug is live (iter-270), so the new test must be shown
  failing against the unmodified derivation before it is trusted. Full `stack-seeding` suite + the
  `role_tenancy_fence_test.go` fence must stay green.
- **C** — live gate: reset-to-seed + the **169 s** binding Playthrough suite on `demo-2`, **watching
  `negative-controls.spec.ts:429`** (the tenancy control that caught this function's first cut).
- **D** — close.

## Escalation conditions

- If the reservation greens `talent-pool.UC1` but reds `negative-controls.spec.ts:429`, it is **not a
  fix** (iter-275's explicit instruction) — revert, close-no-lift with the falsification.
- If the suite reveals the occupancy model is wrong (e.g. the view counts something other than
  membership rows), route forward rather than tune the number.

## Acceptable close-no-lift outcomes

A measurement showing the incumbent count the succession view reads is **not** the seeded
`memberships.job_role_name` population — which would falsify the whole iter-274→276 chain — is a
complete iter even with the metric unmoved.
