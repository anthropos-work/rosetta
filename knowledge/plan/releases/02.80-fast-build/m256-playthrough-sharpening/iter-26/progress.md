**Type:** tik · **Shape:** standard (protocol: `corpus/ops/demo/playthroughs.md` § *The iteration protocol*)

# iter-26 — the day-0 readiness seat

## Phase A — probe live BEFORE writing anything (iter-21's rule)

A stage-0 vantage did not exist anywhere in the system, so one was **manufactured** on `demo-2`: the STARTED
hero's single `ai_readiness_user_step_progresses` row (`skill_mapping / completed`) was backed up to TSV and
deleted, the surface read, then restored by `--reset`. Read-only reasoning would have got two of the three
answers wrong.

**What it measured — hypothesis 2 CONFIRMED, and the trap named:**

| | stage 0 | stage 1 |
|---|---|---|
| step NAME × 3 | 3 / 3 | 3 / 3 ← **discriminates nothing** |
| step-1 card | `1 AI Skill Mapping with AI Framework  Start · ~8 min` | `… Done  Update Skill Mapping` |
| `Update Skill Mapping` control | 0 | 1 |
| step-2 card | **absent** (locked) | `2 Hands-on Assessment with AI Simulation  Start · ~30 min` |
| controls in `<main>` | 31 | 33 |

The funnel **names all three steps at every stage** — which is correct, and is what
`pt-aireadiness-member-progress` asserts for the resumed vantage. It is also why an onboarding assertion built
on step names would have been green for a hero who had already finished the step. The **cards** discriminate.

**Hypothesis 3 CONFIRMED, in three parts, each of which changed the spec:**

1. The step-1 card opens a **5-screen modal** ("AI Literacy") — a paged self-map of the org's AI skills.
2. Its forward control **RELABELS**: `Next` on screens 1–4, **`Go to the AI Simulation`** on screen 5. A walker
   matching only `/^Next$/` pages four screens and then reports "no forward control" — the **iter-18 relabel
   trap in a second product surface**, met again by measuring rather than by reading.
3. **Paging `Next` alone persists nothing** — 0 rows after four screens. The terminal control is the click that
   writes: `skill_mapping / completed` lands, and a fresh `/home` then reads **Done + Update Skill Mapping +
   step 2 unlocked**. It also **navigates the browser off `/home`**, which is why the read-back must
   re-navigate rather than read the post-modal client state.

So the UC is a genuine **mutating** journey with a real DB write, a UI read-back on a fresh navigation, and a
pre-state that IS its own negative control.

## Phase B — the capability (`stack-seeding`)

`blueprint.Persona.AIReadiness` (`""` | `"not_started"`) → funnel stage 0, checked **before** the
trajectory-derived default, read in **exactly one place**. Deliberately **not** a third `trajectory` value —
see **D104** for the two reasons (truth: readiness stage is orthogonal to the life-arc; blast radius: five
seeders switch on `Trajectory` through a silent `default:` each). Closed enum, because an unrecognised value
falls back to stage **3 = COMPLETED**, i.e. a typo would produce exactly the already-finished seat the field
exists to remove.

iter-24's held-gap test was **discharged per its own failure message** and replaced by two tests that pin the
new contract from both sides — the derived default is **unmoved** (the regression guard: every existing
readiness seat is derived) and the declaration **reaches stage 0** — plus a seeder-level test that runs the real
seeder and asserts the day-0 seat has **zero rows of any kind** *while Aria still lands 3 and Ben 1*, so a
seeder that wrote nothing cannot pass it.

**5 mutants RED**, restores byte-identical (`shasum -c`):

| | mutant | RED |
|---|---|---|
| M1 | remove the `IsAIReadinessNotStarted()` case | both new stage-0 tests |
| M2 | hero `default:` → 0 | the plain-append regression guard (+ 2 others) |
| M3 | remove `if stage == 0 { continue }` | the seeder-level no-signals test |
| M4 | remove the `ai_readiness` `validateEnum` | the closed-enum test |
| M5 | resolver always false | the enum test + both stage-0 tests |

## Phase C — the seat, reseeded and PROVEN live

`pt-ai-onboard` (Ola Bergstrom) **appended last** to Org C. Live reseed, then the DB:

| Org C hero | readiness step rows |
|---|---|
| Robin Vance (thriving) | **3** |
| Theo Lindqvist (struggling) | **1** |
| Nadia Ferrante (manager) | 0 |
| **Ola Bergstrom (day-0)** | **0** ✅ |

Roster 30 → **31 identities**; cockpit export 7 → **8 heroes**. The derived heroes are untouched, which is the
append-only property iter-24 fenced, now observed on real rows.

## Phase D — the Playthrough + its in-line control

`pt-onboarding-aireadiness-guided`, `@pt-mutation: MUTATES`. The page-object layer was written first and then
**resolved live on BOTH seats in one run** before a line of the spec was written (the table in Phase A is that
measurement). GREEN on the first live drive, **8.6 s**.

**6 mutants — 5 RED and one that PASSED and was therefore data:**

| | mutant | outcome |
|---|---|---|
| P1 | skip the write | RED on read-back (a) |
| P2 | run the SAME spec on `pt-ai-started` | RED on the FIRST pre-state assert — the control's discriminator |
| P3 | narrow `skillMappingForward` to `/^Next$/` | **PASSED** — the locator's terminal alternative could never fire (commit is checked first). Dead coverage reading as if the relabel were handled twice. **The locator was narrowed**, and the real proof re-cut as P3b |
| P3b | delete the commit check | RED with the page object's own named error, naming screen 5 |
| P4b | assert the step-2 NAME's absence instead of the CARD's | RED — the name is present at both stages |
| P5 | remove `ai_readiness: not_started` and RESEED | RED — **and this is the decisive one:** Ola comes back with **3** progress rows (stage 3) and the guided funnel is **absent entirely**. Without the capability the Playthrough has no subject |
| P6 | drop the fresh navigation | RED — the commit navigates off `/home` |

Plus **4 net-new locator-fence tests** (captured matchers, never re-typed) with their own mutants: **F1** drop
the metachar escape → RED, **F2** widen forward to include the terminal label → RED (this is the assertion that
protects P3's correction), **F3** scope the dialog into `<main>` → RED, **F4** drop the `\b` on `Done` → RED.

## Phase E — the gate, and the two-hour detour inside it

**The first two 3× attempts each returned `1 failed / 192`.** The failures were on two different tests and both
looked like host stalls. **D106 concluded exactly that and hardened the succession liveness assert to 45 s. The
next run failed the same assert again, after the full 45 s** — so the diagnosis was wrong, the hardening was
reverted, and D106 is **retracted in place** with the reason it was reachable.

**The real cause was the seed, and it was mine (D107).** The seat had been given `role: Data Analyst` — chosen
only because that role was *known to resolve*, which is to say because Org C's COMPLETED hero already held it.
That made it a **two-member** role, and Org C's succession **key-role card** for a two-member role is
**non-deterministic**:

| `Data Analyst` occupancy in Org C | key-role card present |
|---|---|
| **2** | **4 of 5** page loads |
| **1** | **5 of 5** page loads |

The casualty was `negative-controls.spec.ts`'s cross-tenant control, on its own **LIVENESS floor** — a control
going RED reading *"succession failed to compute for the contrast tenant."* **A hero's role is an assertion
anchor.** Re-roled to `Supply Chain Analyst` (10 `job_role_skills`, no other Org C holder,
logistics-coherent), and **fenced**: hero roles must be pairwise distinct within a story, with a self-test that
injects a collision. Mutant **N1** puts `Data Analyst` back and the fence goes RED naming
`Vertex Logistics: "Data Analyst" held by pt-ai-completed + pt-ai-onboard`. Warnings left at **both** call
sites — the seat and the anchor.

**Final gate — 3 consecutive cold reset-to-seed runs:**

- **`195 passed` ×3, rc 0 each, 0 flake** (1.8 / 1.6 / 1.7 m). Was 187 at iter-25; +1 Playthrough +2 locator
  fence tests +2 anchor fence tests, and the pre-existing 193.
- `ptreport`: **27/31 passing (87.1%)**, **0 failing**, 4 `unimplemented`, **0 `unimplementable`**.
- `@pt-negative-control` registry: **25 of 27** (11 self-declared + 14 via the control spec).
- `@pt-mutation` registry: **MUTATES=9**, READ-ONLY=16, UNKNOWN=2.
- `stackseed --policy-check`: `live=18 expected=18`, **rc 0**, green after all three resets.
- 16 containers Up, **0 exited**. Drifted cockpit fixture restored **byte-identically** (`99e2f315`).
- `gofmt -l` clean across all six rext modules; `go test ./...` rc 0 in `stack-seeding` (16 pkgs) and
  `playthroughs`.

## Close — 2026-07-30

**Outcome:** onboarding **1 of 5 → 2 of 5** — the gate's last open clause moved — via a seeder capability
(`ai_readiness: not_started`) that made the seat *declarable at all*, spent immediately on
`pt-onboarding-aireadiness-guided` with its control. Controls **24/26 → 25/27**, mutating **8 → 9**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — clause 3's onboarding half is 2 of 5, 3 UCs short — (2) triggered-tok:
n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n —
Outcome: continue
**Decisions:** D104 (the field, not a trajectory) · D105 (the declared P6 boundary) · D106 **RETRACTED** ·
D107 (a hero's role is an assertion anchor)
**Side-deliverables:** the `skillMappingForward` narrowing (a locator alternative that could not fire, found by
mutant P3) — folded into the same commit as the page object it corrects, and pinned by fence test F2.
**Routes carried forward:**
- `ONBOARD-M256-stage0-capability` → **CLOSED this iter.**
- `PT-M256-resume-fixture-pair` → still the best-value next target (2 UCs for one fixture), but iter-18's
  measured product defect stands in front of it: the CV upload POSTs 200 while the forward control never
  enables (100 s+, both formats). **Measure that first** — it may not be ours.
- `ONBOARD-M256-import-path` → the 3 remaining onboarding UCs. `individual.UC1` needs a **member-less user +
  roster seat** (a second, unrelated seeder capability — the audit's F5 kernel of truth);
  `enterprise-workforce-standard.UC2` still has an **unidentified trigger** (measured across four orgs at
  iter-18, all served the import form); `enterprise-hiring.UC1` needs a non-recruiter day-0 hiring seat **and**
  a discriminator against the cockpit's own `is_hiring` routing.
- **NEW —** `PLATFORM-M256-keyrole-nondeterminism`: a succession key-role card's presence varies between page
  loads once its role has 2 occupants. Not a demo defect a presenter would see; **routes to the platform**, not
  fixed here (zero platform edits). The seed-side fence is what protects the suite.
**Lessons:**
1. **Before hardening a flaky assertion, make it flake on command.** A 6-run duration table with every passing
   reading on baseline and one outlier per run is *persuasive* and was *wrong*. The reproduction took three
   minutes. iter-13 set this standard; this iter skipped it and paid two gate cycles.
2. **A seeded hero's attributes are part of the test suite's contract.** iter-13/14 deliberately re-aimed the
   controls at seeded facts by NAME — org domain, org size, hero name, hero **role**. The cost of that
   sharpness is that adding a hero can perturb another Playthrough's anchor probabilistically. Picking a role
   because it "is known to resolve" copies the nearest hero's role, which is precisely the collision.
3. **A capability with no consumer proves nothing, and a consumer with no capability cannot be built honestly.**
   Landing both in one iter is what made mutant P5 possible — and P5 is the only mutant that shows the *seeder*
   is load-bearing for the Playthrough's green, which is the whole claim of the iter.
