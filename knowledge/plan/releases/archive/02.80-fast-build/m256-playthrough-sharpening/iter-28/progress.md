**Type:** tik · **Shape:** standard (protocol: `corpus/ops/demo/playthroughs.md` § *The iteration protocol*)

# iter-28 — the org-prepared onboarding trigger

## Phase A — the read that closed a two-iter unknown

`standard.UC2`'s blocker had read *"the trigger is not yet identified"* since iter-08, narrowed but not
answered at iter-18 (four orgs driven, identical import form for every one). It is
`packages/ui/src/Onboarding/OnboardingUser.tsx:135`:

```ts
const lastStep = reimport ? Import : steps?.[steps.length - 1]?.step;
const [managerImport] = useState(
  Boolean(lastStep === OnboardingStep.Import && organizationName && userStats));
```

`organizationName` and `userStats` are always supplied by the host page, so the only missing input is
`steps` — `public.user_params.onboarding`, **NULL for every seeded user**. Which is exactly why probing could
not find it: **a four-org sweep was sampling a constant.** → **D111.**

`managerImport` renders `<EnterpriseUser>` (name + location + last experience + stat cards) instead of
`<ImportStep>`, and relabels the forward control to *"Start"*.

## Phase B — the capability

`blueprint.Persona.Onboarding` (`""` | `"org_prepared"`) + `OnboardingParamsSeeder`: **one jsonb row** whose
last step is `import`. `import`, **not** `done` — `done` completes onboarding and the route redirects, the
opposite of the use case. Closed enum, for the same reason `ai_readiness` is (an unrecognised value falls back
to the day-0 default, so a seat declared org-prepared would be served the plain form and the Playthrough would
fail *looking like a product regression*).

**And the insert alone silently did nothing** — `public.user_params` is populated **row-per-user at
user-insert time** (191 rows within 300 ms of the users COPY, all NULL, written by nothing in this fleet), so
`ON CONFLICT (id) DO NOTHING` skipped it with no error and the seat kept getting the plain form. Now
insert-then-**heal**, and it **fails the seed** if a declaring hero's row cannot be reached → **D112**. The
missing `audit.Record` was caught by the isolation guard on the first live run. Two guards fired; both right.

**4 mutants RED**, restores byte-identical: R1 write for every hero · R2 step `done` · R3 resolver always false
· R4 remove the `validateEnum`.

## Phase C — the seat, PROVEN live

`pt-onboard-prepared` (Sam Okonkwo, `Business Analyst` — distinct from Pat's and Morgan's per iter-26's D107
anchor fence) appended **last** to Org A. Reseeded; `user_params.onboarding` =
`{"steps":[{"step":"import","updated_at":…}]}` for exactly one user. Then both seats in one run:

| locator | `pt-onboard-prepared` | `pt-free` (day-0) |
|---|---|---|
| `linkedinUrlInput` | **0** | 1 |
| `importFromLinkedInLabel` | **0** | 1 |
| `uploadButton` | **0** | 1 |
| `skipButton` | 1 | 1 |

The import form is **entirely absent** for the org-prepared seat and present for every other seat in the
world — so every existing seat is a live contrast vantage, which is the strongest control position any of this
milestone's UCs has started from.

**But she was first put in Org A, and that cost a RED gate — deterministically, 3 of 3 runs (D114).**
`pt-workforce-funnel`'s iter-14-sharpened final asserts that **Pat Ellis's member-spotlight CARD carries her
seeded role**, and one extra Org A member displaced Pat from the spotlight entirely. **This is D107 one axis
over** — there a role's *occupancy* perturbed a key-role card probabilistically, here an org's *member set*
perturbed a spotlight deterministically — and the underlying fact is now twice-paid-for: iter-13/14 aimed the
controls at seeded facts by NAME, and the price of that sharpness is that adding a hero to an anchored org
moves another Playthrough's anchor.

**Org B is the safe host by construction:** it is the only pt org `seed-facts.ts` does not name
(`SEEDED_ORGS = [A, C, D]`), and it is a workforce org. Moved, renamed `Elin Marchetti`, role kept distinct
from `pt-free`'s. **Re-proven on the SHIPPED seat** — which matters, because `organizationName` is one of
`managerImport`'s three inputs and the org changed: import form **0/0/0**, the relabelled **`Start`** control
**1**, and **her own name** on the summary **1** (evidence the Org A probe did not have — the prepared summary
genuinely renders *her*).

**Gate: `196 passed` ×3 consecutive cold reset-to-seed, rc 0 each, 0 flake** (1.8 / 1.7 / 1.5 m); `ptreport`
**28/31 passing, 0 failing, 0 unimplementable**; controls **26 of 28**; MUTATES **10**; `gofmt -l` clean;
`go test ./...` rc 0 in `stack-seeding` (16 pkgs).

## Phase D — ROUTED, deliberately (D113)

The UC's flow continues *"confirm or adjust the pre-filled role, refine the suggested skills"*, and on this
variant the forward control is *"Start"*, whose handler is a branch nobody has driven: whether it **completes**
onboarding or **advances** to the Role step is **unmeasured**. The session budget ran short of measuring it, and
the iter's own escalation condition named this stop. Writing the spec anyway would assert a multi-step journey
nobody drove — the iter-22 failure (a spec that *could not have passed*) and the iter-27 failure (a "read-back"
that read the seed).

## Close — 2026-07-30

**Outcome:** the milestone's last *unknown* blocker is closed. The org-prepared trigger is identified at
`file:line`, seedable, implemented with 4 RED mutants, and **proven live** with total discrimination against
every other seat. The Playthrough itself is routed, one probe run wide.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — onboarding 3 of 5 — (2) triggered-tok: n — (3) re-scope: n —
(4) user-blocker: n — (5) cap-reached: n (3 tiks this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D111 (found by reading; a sweep over a constant) · D112 (the silent no-op, now loud) ·
D113 (the deliberate stop) · D114 (the seat lives in Org B — a RED gate taught it)
**Side-deliverables:** none.
**Routes carried forward:**
- **`ONBOARD-M256-prepared-playthrough`** → the Playthrough for `standard.UC2`. Everything needed is landed and
  proven; the residual is **one probe run**: drive `Start` on `pt-onboard-prepared` and record whether it
  completes onboarding or advances to the Role step (`OnboardingUser.tsx:470` is the handler). Then the spec is
  the measured table above as its pre-state + control, `Skip` or `Start` as the action, and the
  **revisit-redirects** read-back that `onboarding.completion.UC1` already proves works in apps/web — which is
  also the assertion that answers the Q1 question (*which assertion fails if I delete the action?*).
- `ONBOARD-M256-import-path` → `individual.UC1` is now the ONLY onboarding UC with an unpriced capability
  (member-less user). `standard.UC1` stays behind the CV-upload product defect.
**Lessons:**
1. **When a probe sweep returns the same answer for every vantage, the input is not one of the axes you are
   varying** — and more vantages will not help. iter-18 drove four orgs against a column that is NULL for all
   191 seeded users. Two iters of "not yet identified" were two iters of sampling a constant.
2. **An idempotency clause can hide a capability's no-op.** `ON CONFLICT DO NOTHING` is the right default and it
   was the wrong tool here, because the row already existed for a reason nothing in the fleet documented. A
   seeder whose no-op presents as a product defect should assert that it wrote something.
3. **Before appending a hero, check whether `seed-facts.ts` names her org.** If it does, expect to perturb an
   anchor. Org B is the one org nothing anchors on, and it is now the default host for a seat that exists to
   serve one Playthrough. Twice in three iters an appended hero has broken an unrelated Playthrough's final;
   both times the gate caught it, and both times the fix was the seat, not the assertion.
4. **Stopping at a measured boundary is a result.** The escalation condition was written into the iter's plan
   before the work started, and it fired for the reason it was written for. The alternative — a spec asserting
   an undriven multi-step journey — is the one failure this milestone has paid for twice.
