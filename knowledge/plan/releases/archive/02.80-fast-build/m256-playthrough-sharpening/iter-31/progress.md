**Type:** tik · 2 planned lines (declared in `overview.md`, so the scope-creep tripwire counts against a
2-step planned shape)

# iter-31 — the weak assertion inside the new work, and the verdict that is not a TODO

## Phase A — measure first (and the routed premise died here)

Three probe runs on `demo-2`, all deleted before the gate.

**A/B — what does a reload actually land on?** Elin's `public.user_params.onboarding` was rewound to the
pre-confirm state (`[{import}]`) on the demo DB, backed up first, so pre-state and post-state could be compared
in one run:

| | progress | buttons | distinguishing content |
|---|---:|---|---|
| **B1** pre-state `/onboarding` | 0 | `[Skip] [Start]` | the prepared summary — her name, location, stat cards |
| **B2** after the confirm (same page) | 50 | `[Back] [Next]` | *"Select Your Top Skills"* · `Change Role` · 10 real taxonomy skills |
| **B3** FRESH `/onboarding` | **0** | `[Skip] [Start]` | **byte-identical to B1** |
| **B4** second fresh navigation | 0 | `[Skip] [Start]` | identical again |

The `role` step was verified present in the DB at B3 (`[{import},{role}]`, written 15:31:19 that run). So the
write lands and **has no observable consequence on a fresh navigation.** Two earlier observations (A1/A2, a
separate session 3½ hours after the previous suite run's write) agree — **six fresh navigations, three browser
sessions.**

**The mechanism, read at source** (the iter-28 move: read the platform after probing plateaus):

```
useGetOnboardingStatus.tsx:25-27   sorts onboarding.steps NEWEST-FIRST
OnboardingUser.tsx:130-132         lastStep = steps?.[steps.length - 1]?.step   ← the OLDEST step
```

One array, two consumers, opposite ends — the host page reads index 0 and is correct. Invisible for a NULL or
single-element array (all 191 seeded users); reachable **only** by the multi-step seat iter-28 minted. Full
record: milestone `decisions.md` § `PLATFORM-M256-onboarding-step-not-resumed`.

**C — is completion reachable instead?** No. One `Next` past the skills screen reaches *"Add more skills"*, which
renders *"We're having trouble loading your skills at the moment"*, and its `Next` is **inert** — clicked five
times, identical screen, progress stuck at 100. So the `/onboarding` → `/home` completion read-back is
unreachable from this seat, and a spec written for it **could not have passed** (iter-22's failure mode).

Elin's row was restored byte-identically (`diff` clean) after the probes.

## Phase B — what shipped

**Line 1 — `ONBOARD-M256-prepared-persistence`: REFUTED, and the honest artifact shipped in its place.**

- A **positive, hydration-proof** assertion on the screen the reload really lands on, labelled in the docstring,
  the failure message and the manifest for exactly what it proves — *the confirmation left her flow intact; it
  neither ejected her nor completed it* — and explicitly **not** the write's read-back (D119).
- The page object's own docstring, which claimed `changeRoleControl()` *"on a FRESH navigation is the
  server-side read-back"*, **corrected in place**: it was written in good faith at iter-29 and is false. A
  docstring is read as evidence by the next iter.
- Both measured non-facts recorded **at the locator**, so nobody re-derives them: the non-resumption, and the
  dead-end that makes completion unreachable.
- The defect routed to the **platform** at milestone level (never only in a spec comment — the iter-23 rule).

**Line 1's bonus, and the more durable half — `liveness-before-absence-fence.unit.spec.ts`.** The rule iter-12
wrote in prose, iter-22 re-broke, and iter-29 broke again *in a brand-new spec one hour after writing the prose*,
is now machine-checked across every Playthrough spec. Sized before adopting: **29 files · 62 navigation sites ·
184 liveness witnesses · 37 absence assertions · 0 violations** — already true everywhere, **zero edits**, so it
buys the next spec. Fail-closed on three floors.

**Line 2 — the WRITTEN VERDICT (D121).** A use case with no Playthrough now carries a `verdict` block (closed
disposition · `measured_by` · `rationale` · `handler`), `ptreport` renders it instead of *"declared use case, no
Playthrough yet (build-reference gap)"*, and the fence runs **both ways** plus **against vacuity**. Both TODOs
filled: `standard.UC1` → `will-not-build` (iter-18 + `D104`, no handler), `individual.UC1` → `not-yet-built`
(iter-30 `D116`, handler `ONBOARD-M256-orgless-seat`). The four states and glyphs are unchanged.

The map now reads, on a live run:

```
[TODO] onboarding.enterprise-workforce-standard.UC1  unimplemented — VERDICT will-not-build: MEASURED, then deliberately refused…
[TODO] onboarding.individual.UC1                     unimplemented — VERDICT not-yet-built: LANDABLE, and priced from measurement…
```

## Phase C — mutants (10 watched, every one RED)

| # | mutation | result |
|---|---|---|
| **M1** | re-introduce iter-29's **exact S1c shape** — a `goto` then a bare `toHaveCount(0)` | **RED** — the fence names `onboarding-org-prepared.spec.ts:211 … since the navigation at :210` |
| **S1** | delete the confirm click (re-confirm) | **RED** at *"the role was accepted and the flow advanced"* |
| **S3** | navigate elsewhere (`/library`) before the new assertion | **RED** — route-anchored |
| **S4** | assert another hero's name (`Pat Ellis`) | **RED** — identity-anchored |
| **S5** | plant a newest `done` step in the DB | **RED** — the ejection mode is REAL, so *"she was not redirected off the flow"* has content. Fires at the FIRST liveness assert, not the new line, because a `done` state cannot be introduced mid-flow from outside — **stated, not glossed** |
| **V1** | delete the `verdict` block from `standard.UC1` (**shipped** manifest) | **RED** rc 1 — *"TODO with NO verdict block"* |
| **V2** | `disposition: deferred-to-a-later-release` | **RED** — *"is not one of [not-yet-built will-not-build] … deliberately no fallback"* |
| **V3** | blank `measured_by` on `individual.UC1` | **RED** — *"an unsourced claim is exactly the estimate that gets read as a measurement by the next iter (D117)"* |
| **V4** | `not-yet-built` with no handler | **RED** |
| **V5** | `will-not-build` **with** a handler | **RED** — *"a refusal with an assignee is a contradiction"* |
| **V6** | a stale verdict on `UC2`, which iter-29 landed | **RED** — direction 2 of the fence |

All six V-mutants ran against the **shipped** manifest, not a fixture: a fence proven only on fixtures is a
fence proven against itself (iter-16, where five green unit tests drove a mock path the real client never used).
Plus **20 new Go tests** (the rationale floor asserted from both sides; a shipped-corpus check that fails closed
on an empty load; the closed-set error message asserted **deterministic over 20 runs**, since Go map order would
otherwise make the same failure read differently each time).

## Phase D — the gate

**`199 passed` × 3 consecutive cold reset-to-seed runs, rc `0` each** (captured into a variable per run, never
off a pipe), **0 flake**. 197 → 199 is the two new fence tests.

| | |
|---|---|
| `ptreport` | **29/31 passing (93.5 %), 0 failing, 0 `unimplementable`, 2 `[TODO]` — both now carrying written verdicts** |
| `@pt-negative-control` (computed) | **27 of 29** (13 self-declared + 14 via the control spec) — unchanged; no Playthrough added |
| `@pt-mutation` (computed) | **MUTATES=11 READ-ONLY=16 UNKNOWN=2** — unchanged |
| `ptvalidate` | VALID — 10 products, 31 use cases, 29 live, 2 TODO |
| `stackseed --policy-check --stack demo-2` | rc 0 · `live=18 expected=18` |
| containers | **16 Up / 0 exited** (`docker ps -a`, per iter-15 D76) |
| DRIFTED cockpit fixture | restored + **sha-verified `99e2f315` after each of the three resets** |
| Go (`playthroughs`) | `gofmt -l` clean · `go test -count=1 ./...` 4/4 ok |
| Python (one invocation each, rc into a variable) | `demo-stack` **999 passed / 1 skipped** · `stack-core` **287** · `stack-verify` **171** · `stack-injection` **266 / 1 skipped** — rc 0 each |

**Clause 1's leg half is N/A — iter-31 landed no speed mechanism**, so there is no leg to measure. Its flake half
is **MET** (0 flake × 3 cold). No suite-ratio number is quoted as a gate, per `D-v28-13`.

## Close — 2026-07-30

**Outcome:** the routed persistence repair was **impossible** — a reload re-serves the pre-state screen, and the
mechanism is a platform defect (one array read from both ends) that only the seat iter-28 minted can reach. The
honest artifact shipped instead, labelled for exactly what it proves; the rule that caught the original defect is
now a machine-checked fence; and every uncovered use case now carries a **written verdict** the four-state map
renders, so clause 3's *"zero silent gaps"* is a machine property rather than a prose claim.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: **n** — clause 3's onboarding half is 4 landed + 1 verdict, which is MET per
`D104`, but `individual.UC1` is the fifth UC and iter-32 is landing it, so the gate is not yet graded MET —
(2) triggered-tok: n — iter-31 moved both planned lines, and the last 3 tiks (29, 30, 31) are not a no-prog
streak — (3) re-scope: n — 0 `unimplementable`; the trigger has moved further away every iter — (4) user-blocker:
n — nothing needs a decision that changes what code lands — (5) cap-reached: n — 1 tik this session — (6)
protocol-stop: n — **Outcome: continue**
**Decisions:** D118 (the routed premise was a source read, not a measurement — D117 extended) · D119 (when the
honest assertion does not exist, ship the measurement plus the label) · D120 (the fence, sized before adopting) ·
D121 (the written-verdict contract) · D122 (`standard.UC1` deliberately not landed — D104 upheld, the 13th iter)
· milestone-level `PLATFORM-M256-onboarding-step-not-resumed`
**Side-deliverables:** the page-object docstring correction (a false claim written in good faith at iter-29) —
recorded here rather than folded into the status, since it was not planned scope.
**Routes carried forward:**
- `ONBOARD-M256-orgless-seat` → **iter-32** (the last UC; priced ≥ 5 seeders, LANDABLE).
- `PLATFORM-M256-onboarding-step-not-resumed` → **the platform** (captured, not fixed; 2 defects).
- `ONBOARD-M256-prepared-persistence` → **CLOSED as REFUTED**; superseded by the platform routing above.
**Lessons:**
1. **A source read is a hypothesis; citing a `file:line` does not make it an observation.** iter-29's routed
   repair named two lines of platform source and was still wrong about what the screen does, because nobody had
   loaded it. This is D117 one class over — not a mis-priced *cost* but a mis-stated *behaviour*, and the
   citation is what made it read as measured.
2. **When the honest assertion does not exist, ship the measurement plus the label — never a weaker assertion
   under the strong one's name.** And check the *comments* for the same offence: this iter found the page
   object's own docstring asserting a read-back that does not exist.
3. **A rule you have applied by hand three times belongs in a test — and size it before you adopt it.** The
   scan found 0 violations across 29 files, so the fence cost zero edits and buys the next spec. Had it found
   twenty, the right move would have been to fix them first and fence second.
4. **The seed capability built to reach a surface can be the only way to SEE a defect on it.** iter-28 minted
   the only multi-step onboarding user in existence; that made a latent platform defect reachable for the first
   time. Same shape as iter-11, where withholding one grant exposed four releases of leaked rows.
