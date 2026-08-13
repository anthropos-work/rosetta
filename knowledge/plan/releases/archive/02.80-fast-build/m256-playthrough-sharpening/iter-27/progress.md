**Type:** tik · **Shape:** standard (protocol: `corpus/ops/demo/playthroughs.md` § *The iteration protocol*)

# iter-27 — the hiring-org day-0 candidate

## Phase 0 — the target was chosen by LANDABILITY, and the choice was made from source

Three of the four remaining onboarding use cases have a blocker that is not ours to move: `standard.UC1` sits
behind a measured product defect (CV upload POSTs 200 while the forward control never enables, 100 s+, both
formats — iter-18's refusal stands), `standard.UC2`'s trigger is still unidentified across four measured orgs,
and `individual.UC1` needs a **member-less user**, which means excluding a hero's slot from every per-index
seeder that writes org-scoped rows — a large blast radius for one use case.

`enterprise-hiring.UC1` was the one whose blockers were all inside our own seed and harness. **Two source reads
re-priced it before any code was written:**

1. **The routing trap is real, and worse than the warning.** `UserStatusContext.tsx:142-173` ejects to the
   hiring app on `userHasAllHiringOrgs` **alone** — no onboarding condition. So the cross-app routing is owned
   by neither onboarding nor the cockpit. **But the hiring app has its OWN onboarding route**, closing with
   `router.replace('/home')` — that half *is* onboarding's. → **D108.**
2. **"Assigned and ready to start" needs no new capability.** `heroHiringStage` already pins a **struggling**
   candidate hero to `assignedOnly`, and an end-user hero in a hiring org is a **candidate** automatically
   (`endUserHeroRole`, M224). One YAML append, no Go change.

## Phase A — the seat, and the journey probed live

`pt-hiring-onboard` (Ivo Kalman, `Account Executive` — distinct from Quinn's role per iter-26's D107 fence,
10 `job_role_skills`, no other Org D holder) appended **last** to Org D. Reseeded, then driven:

| step | measured |
|---|---|
| hiring `/onboarding` | **serves the flow** — LinkedIn URL + Upload + **Skip** + a *disabled* Next; 6 controls; **no** `/sim/` link |
| `[Skip]` | lands on **`:23001/home`** — the HIRING app — greeted *"Hi, Ivo! 👋🏻"*, chrome *"Kestrel Hiring Group"* |
| hiring `/home` | 9 controls, incl. her position as a startable org-scoped link `/sim/<slug>?organizationId=<Org D>` |
| revisit `/onboarding` | **does NOT redirect** — it serves the flow again, *unlike* `apps/web` |

## Phase B — the Playthrough, and the two mutants that rewrote it

`pt-onboarding-hiring-candidate`, `@pt-mutation: MUTATES`. GREEN first live drive, **3.0 s**.

Then the mutants, and **two of them passed** — both recorded in **D109**:

| | mutant | outcome |
|---|---|---|
| **Q1** | skip the write, navigate to `/home` by hand | **PASSED.** Her whole home is **seed state**: greeting, tenant, assigned position all exist *before* she onboards. The spec's "read-back" read something the write never touched — a Playthrough that would have gone green **without anyone onboarding** |
| Q1b | the same mutant against the FIXED spec | **RED** — `waitForURL` times out on `/onboarding` |
| **Q5** | seed: `trajectory: thriving` (→ ASSESSED, already taken) | **PASSED.** The surface does **not** discriminate taken from not-taken |
| Q3 | greet another hero's name | RED |
| Q4 | name another tenant | RED |

**Q1's fix:** delete the manual navigation and let `apps/hiring`'s own `router.replace('/home')` be the
observation — so the click is the only thing that can satisfy it, labelled as such in the source.
**Q5's response:** state it. The final asserts the **affordance**; "not yet taken" is a **seed guarantee**
(`heroHiringStage` → `assignedOnly`), and `trajectory: struggling` is kept because it is *truer*, not because an
assertion leans on it.

## Phase C — the gate

- **`196 passed` ×3 consecutive cold reset-to-seed, rc 0 each, 0 flake** (1.8 / 1.6 / 1.6 m).
- `ptreport`: **28/31 passing**, **0 failing**, 3 `unimplemented`, **0 `unimplementable`**.
- Controls **26 of 28** (12 self-declared + 14 via the control spec); `@pt-mutation` **MUTATES=10**.
- `--policy-check` `live=18 expected=18` rc 0; 16 containers Up / 0 exited; `gofmt -l` clean; `go test ./...`
  rc 0 in `playthroughs` (4 pkgs) and `stack-seeding`.
- Drifted cockpit fixture restored **byte-identically** (`99e2f315`).

## Close — 2026-07-30

**Outcome:** onboarding **2 of 5 → 3 of 5**; controls 25/27 → **26/28**; mutating 9 → **10**. The suite's first
Playthrough to drive an onboarding flow in the **hiring** app, with the cross-app routing claim explicitly
**not** made and the reason recorded at `file:line`.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — onboarding is 3 of 5, 2 UCs short — (2) triggered-tok: n — (3) re-scope:
n — (4) user-blocker: n — (5) cap-reached: n (2 tiks this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D108 (the routing half is not asserted, and why) · D109 (two passing mutants; "present after" is
not evidence on a seeded world) · D110 (the seed-guarantee boundary)
**Side-deliverables:** none — `hiring-home-page.ts` is net-new but it is the iter's planned page-object layer.
**Routes carried forward:**
- `ONBOARD-M256-import-path` → **2 UCs left.** `standard.UC2`'s trigger is still unidentified and is now the
  cheapest remaining *investigation* (read the shared `OnboardingUser` component: both apps mount the SAME one,
  so whatever selects the org-prepared variant is a prop or a query in there — a source read, like D108's, not
  a probe sweep). `individual.UC1` remains the expensive one (member-less user).
- `standard.UC1` → still behind the CV-upload product defect. **Do not re-litigate iter-18's refusal to move a
  number.**
- **NEW —** `PT-M256-standing-mutant-Q1`: run *"delete the action and see whether anything fails"* against every
  mutating Playthrough on this world, not just new ones. Q1 found a green-without-the-write in a spec that had
  already been reviewed, and the check is one edit + one run.
**Lessons:**
1. **On a seeded world, "the outcome is present after the action" is not evidence.** It is evidence only if the
   outcome was absent before. Every mutating Playthrough should be able to answer *"which single assertion fails
   if I delete the action?"* — and if the answer is "none", the read-back is reading the seed.
2. **A source read can settle in ten minutes what a probe sweep cannot settle at all.** iter-18 measured four
   orgs to find the routing discriminator and could not; `UserStatusContext.tsx:142-173` answers it in one
   `useEffect`. When the question is *"what causes this?"* rather than *"what does this render?"*, read the code.
3. **Two mutants passing in one iter is not a bad sign, it is the protocol working.** Both changed the shipped
   artifact. The bad sign would have been not running them.
