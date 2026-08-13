**Type:** tik · target `ONBOARD-M256-orgless-seat` — the last un-homed use case in the M201 curated corpus

# iter-32 — `individual.UC1`, the org-less user

## Phase A — measure first

**The seat's population slot** is unique and is not slot 1 (`identity.go` pins the story's admin/assigner
membership there); a loud seed-time guard now enforces it (`OrgLessOnAdminSlot`) because
`personaUserIndexFor` HASHES, so no seed author could see the collision by reading the YAML.

**Her live surfaces**, measured before a line of spec was written — and she is served a **third onboarding
component**, not a variation of one:

| | |
|---|---|
| `/onboarding` | *"Your journey is about to start…"* · **"Your Role: Operations Manager"** · `[Change Role] [Done]` · **0 textboxes** — no LinkedIn URL, no CV upload |
| `[Done]` | onboarding COMPLETES server-side; the app moves her to `/library/ai-simulations` |
| fresh `/onboarding` | **REDIRECTS to `/home`** |
| `/home` **before** she onboards | renders fully — *"Hi, Nils"*, 0 AI Sim / 0 Paths / 0 XP, the whole nav |

That last row decided the spec's shape: **her home is seed state**, so "she lands in the app" proves
nothing (iter-27's mutant Q1 class). The only thing the write can satisfy is that the app itself takes her
**off** the flow — so the test never navigates to `/home` by hand.

And the read-back exists here for a reason worth naming: the host page reads `steps[0]`, the **newest**
step, which is the one reading of that array the platform gets right. Her sibling
`pt-onboarding-org-prepared` has **no** read-back at all because the shared component reads the **other
end of the same array** (iter-31's `PLATFORM-M256-onboarding-step-not-resumed`). Same column, same run,
opposite outcomes.

## Phase B — the capability

`Persona.OrgMembership` (`""` | `"none"`, closed, no fallback): her `public.users` row and her Clerkenstein
identity exist; her membership row, her `g2`/`g3` grants and her Clerk org claim do not. **Two halves**
(D123) — iter-30's measurement only ever exercised the DB one. `verified: 0` is **required**, and a manager
is refused (D124).

## Phase C — the sweep, which is where the work actually was

iter-30: *"the FK list is the entry fee; the org-scoped-row tail is the real work."* Both halves arrived
(D125).

**The loud half**, on the first reseed after the capability landed:

```
succession rows=0  ERROR: violates foreign key constraint "interview_extraction_results_sessions_session"
```

Friendly, exactly as predicted — it named its consumer.

**The quiet half**, found by sweeping the live DB for her uuid across **every uuid column** in `public` +
`jobsimulation`:

| | before the guards | after |
|---|---|---|
| `jobsimulation.sessions` (carries `organization_id`) | **2** | 0 |
| `jobsimulation.activity_events` | 4 | 0 |
| `public.skill_path_sessions` | 2 | 0 |
| `public.personal_assignments` | 4 | 0 |
| `public.user_bookmarks` | 3 | 0 |
| `public.user_skills` / `user_skill_evidences` | 36 / 36 | 0 / 0 |
| user-scoped (users · basic_info · params · languages · educations · projects · certifications · experiences) | kept | **kept** |

**Final state: 10 tables, ALL user-scoped, 0 carrying `organization_id` or a membership FK, 0 casbin
grants, no Clerk org claim.** Org B: **19 memberships** against `size: 20` — the one slot she consumes
without becoming a member.

**Two pre-existing fences refused the seat before it shipped (D126)** — and neither refusal came from
review. The M224 curated-pool fence rejected `Product Designer` (classifies to no curated family → the
taxonomy's alphabetical junk head); the M219-R8 ladder fence then rejected the replacement, because her
default **65**-skill claimed tail would have drawn the `operations` family dry and shipped a *silently thin*
profile. The fence suggested growing the allow-list. **That would have been the wrong fix**: an org-less
day-0 user should claim nothing, and now does.

## Phase D — the Playthrough + 5 mutants

`pt-onboarding-individual`, GREEN on its first drive (24.1 s). Its cross-vantage control is read in-line —
the **enterprise import form** every org member is served reads 0 for her, four locators.

| # | mutation | result |
|---|---|---|
| **N1** | delete the `Done` click | **RED** at the lands-in-app wait — the action is load-bearing |
| **N4** | iter-27's **standing Q1** — action *and* intermediates deleted, read-back only | **RED** at the redirect — the final cannot pass without the write |
| **N3** | assert another hero's role | **RED** — identity-anchored |
| **N2** | drop `org_membership: none` and **RESEED** | **RED** at liveness — she comes up an ordinary member and gets the enterprise form; **the SEED is load-bearing** |
| **V7** | re-add a verdict to the now-LIVE `individual.UC1` | **RED** (`ptvalidate` rc 1) — iter-31's bidirectional fence, exercised live for the first time; **landing this UC is what forced the verdict's removal** |

**The first mutant pass was thrown away, and that is the lesson (D127).** Run against a world the green
drive had already consumed, N1 went red at the *wrong line* and **N4 PASSED** — a false pass that looks
exactly like a weak assertion, caused entirely by state. For an irreversible write the protocol is **reset →
mutate → run, every time.** The spec's own failure message warned about it; reading a warning is not
obeying it.

## Phase E — the gate

**`200 passed` × 3 consecutive cold reset-to-seed runs (4, 5, 6), rc `0` each** (captured per run into a
variable), **0 flake**.

**Reported, not re-rolled away (D128):** an earlier batch of three had **one** failure —
`pt-workforce-succession` on run 3 — and it is **not iter-32's**: that Playthrough reads **Org A**, the new
seat is in **Org B**, and Org A measured unchanged afterwards (40 memberships, DevOps Engineer occupancy
**1**). It is a recurrence of **`PLATFORM-M256-keyrole-nondeterminism`** that **extends** iter-26's record
rather than being covered by it — iter-26 measured 4/5 absent at occupancy **2** and 5/5 present at
occupancy **1**; this absence was at occupancy 1. Re-measured in isolation immediately after: **6/6
passing**. Run 3 was also the slowest of the six (2.8 m vs 2.0 m) on a 9.7 GiB Docker VM against a 12 GB
floor — stated as plausible and unmeasured, not as a finding. **No timeout was bumped** (iter-26 retracted
exactly that).

| | |
|---|---|
| `ptreport` | **30/31 passing (96.8 %), 0 failing, 0 `unimplementable`, 1 `[TODO]`** — `standard.UC1`, carrying its written verdict |
| `@pt-negative-control` (computed) | **28 of 30** — the new Playthrough arrived already covered, so numerator and denominator both moved; the 2 uncovered remain the studio pair (`D103`) |
| `@pt-mutation` (computed) | **MUTATES=12** READ-ONLY=16 UNKNOWN=2 (was 11) |
| `ptvalidate` (+ datadna) | VALID — 10 products, 31 use cases, **30 live, 1 TODO**; closure **PASS** (279 node-ids) |
| `stackseed --policy-check` | rc 0 · `live=18 expected=18` |
| containers | **16 Up / 0 exited** |
| cockpit fixture | restored + **sha-verified `99e2f315` after all 10 resets** |
| Go | `gofmt -l` clean; `go test -count=1 ./...` green in **both** modules |

Clause 1's **leg** half is N/A (no speed mechanism landed); its **flake** half is MET on runs 4–6.

## Close — 2026-07-30

**Outcome:** `onboarding.individual.UC1` is **LIVE** — the last un-homed use case in the M201 curated
corpus, un-homed for five releases and priced as impossible by this milestone's own pre-flight audit. The
capability is measured, fenced and swept; the Playthrough is mutation-proven four ways including iter-27's
standing Q1.
**Type:** tik
**Status:** closed-fixed
**Gate:** MET
**Phase 5 grading:** (1) gate-met: **y** — clause 1 flake MET (0 flake × 3 consecutive, the earlier flake
reported + diagnosed) and leg N/A · clause 2 mutating **12/5**, `blocked` **1/1**, controls **28 of 30**
MET via `D103` · clause 3 verdicts **31/31 with 0 unimplementable**, org-admin **4/4**, onboarding **4
landed + 1 written verdict** MET via `D104` · `D-v28-5` FIXED and proven live at iter-25 —
(2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks) —
(6) protocol-stop: n — Outcome: **exit-1**
**Decisions:** D123 (two halves — the identity half iter-30 never exercised) · D124 (`verified: 0` required,
not permitted) · D125 (the FK surface is static, the org-scoped tail is not — two guarantees, and the fence
says so) · D126 (two pre-existing fences refused the seat; the second's own suggested remedy was wrong) ·
D127 (an irreversible write needs a fresh world before EACH mutant) · D128 (the run-3 flake, diagnosed and
kept on the record)
**Side-deliverables:** none.
**Routes carried forward:**
- `PLATFORM-M256-keyrole-nondeterminism` → **the platform**, with iter-32's new evidence attached (it also
  occurs at occupancy 1).
- `PLATFORM-M256-onboarding-step-not-resumed` (iter-31) → the platform, unchanged.
- `DOC-M256-claudemd-pt-count` → **milestone close** — `CLAUDE.md` still reads "18 live Playthroughs"; the
  count is now **30**.
**Lessons:**
1. **When a fence suggests a remedy, the remedy is still a judgement.** The ladder fence said *"grow the
   `operations` allow-list"*. The real finding was that a day-0 solo user should have no claimed tail at
   all. Following the suggestion would have shipped an incoherent persona *and* a bigger allow-list.
2. **A capability's cost is measured by deleting; its CORRECTNESS is measured by sweeping.** iter-30's
   delete priced the change in five minutes. Only a sweep of every uuid column found the rows that break
   nothing — and those were the ones that mattered.
3. **For an irreversible write, reset before every mutant.** A consumed world turns a mutant into a
   measurement of state, and it produced both a false RED and a false GREEN in the same pass.
4. **A fence should state what it does NOT cover.** The static call-site scan cannot see an
   `organization_id` write, and a fence that implied otherwise would be more dangerous than no fence,
   because the uncovered half is the half that fails silently.
