# M256 · iter-01 — progress

**Type:** tok (bootstrap) — build-mstone-iters Phase 0 rule 1 (unconditional iter-01 of an iterative
milestone). Iter shape per `corpus/ops/demo/playthroughs.md` § "The iteration protocol".

## Phase 0b — pre-flight KB-fidelity gate

Run once at milestone start, **before** the strategy is authored. Verdict + findings recorded in
`../spec-notes.md` § "Pre-flight audits — iter-01".

## Phase 2 — what this tok did

1. **Priced the parallel-lane enabler against the RE-CUT gate** and found the overview's premise stale:
   clause 1 is a **per-test median**, so worker count cannot move it. Recorded as D1, with both enabler
   options priced from code and routed forward. This removes the milestone's largest risk item from the
   critical path.
2. **Measured the mutation partition** across all 18 browser specs (10 explicit "(no mutation)" / 2
   explicit MUTATES / 6 UNCLASSIFIED) — confirming the plan review's count exactly, and giving the
   machine-checked tag a known target set.
3. **Found the real clause-1 lever** (D3): **12 of 18** browser specs omit `waitUntil` on `loginAsHero`
   and therefore inherit `cockpit-login.ts`'s **`'networkidle'`** default — on an app whose own helper doc
   records that `networkidle` "resolves late and for the wrong reason". Plus two unfenced page-object
   `goto`s still pinning `networkidle` (`skill-path-page.ts:31`, `simulation-page.ts:36`) while the fenced
   base class uses `domcontentloaded`.
4. **Found that clause 2 and clause 3 collapse** (D2): the four curated org-admin UCs each declare a
   persist-then-observe final, so landing them yields 4 mutating Playthroughs → 5 with
   `pt-assignment-assign` → clause 2's `≥ 5 mutating` floor, while being half of clause 3.
5. **Found the `blocked` outcome is free** (D4): `seed-worlds.yaml` already declares `pt-free`, the `free`
   entitlement, and the `entitlement-gated` capability.
6. **Extended `playthrough-map.md` into a ranked triage** (new §8) — the levers ranked with `file:line`
   evidence, the risk named, and the execution order derived. §1–§7 untouched.
7. **Stood up the measurement surface**: a LOCAL `demo-2` bring-up (`--no-public-host`), so iter-02 can
   measure the baseline the re-cut gate divides by.

## Environment (stated, per the `latency-budget.md` rule)

- Host: `Kirality-Mac-Pro-6.local` (darwin 25.1.0), Docker VM **MemTotal 10 419 826 688 B ≈ 9.70 GiB**
  against the documented **12 GB** UI-tier floor — the bring-up's RAM warning is **expected, not a
  failure** (`../overview.md` local-host caveat).
- Disk: 223 GiB free on `/`; docker reclaimable at bring-up start 9.40 GB images + 7.95 GB volumes +
  8.20 GB build cache.
- Stack: `demo-2`, offset **20000** (app `:23000`, hiring `:23001`, studio `:29000`, fapi `:25400`,
  cockpit `:27700`), **localhost** scheme `http` (no `--public-host`; `billion` is off limits under the
  standing sign-off rule).
- rext: authoring copy `main` @ `6ca8764`; demo consumes the pin `cockpit-deeplinks-v1` (`c755214`) —
  **verified present on origin**. The only playthroughs-section difference between the two is
  `manifest/hiring_isolation_test.go` (a Go *test* file); **no e2e/runtime file differs**, so a suite run
  from the authoring copy is equivalent to the pinned code for browser-suite timing purposes.

## Phase 0b outcome

**YELLOW** — proceed with the gaps as known-context. Verdict + the findings that changed the strategy are
recorded in `../spec-notes.md` § "Pre-flight audits — iter-01"; the report is `../kb-fidelity-audit.md`.
The audit **refuted this iter's own first-draft D4 inside the same iter** (see `decisions.md` D4) and
**strengthened D3** (8 unfenced `networkidle` violations, not 2).

## Phase 2 addendum — the stack came up, and its one alarm is a false one

`demo-2` bring-up: **10:48:44Z → 11:12:10Z ≈ 23.5 min** (16 containers), `--no-public-host`, from a warm
image cache. It closed with `⚠⚠ autoverify demo-2: 1 check(s) FAILED — the Clerkenstein fake-FAPI is NOT
answering on :25400 — NOBODY CAN LOG IN`. **That alarm is false** — proven four ways in `decisions.md` D5
(container listening + `openssl s_client` full handshake + matching cert/key modulus + **Chromium `GET
/v1/environment` → 200**). The failing client is macOS system `curl` on **LibreSSL/3.3.6**
(`bad decrypt`), i.e. the *verifier's* TLS stack, not the stack under test. Routed forward as
`FIX-M256-autoverify-fapi-libressl`.

The documented RAM warning behaved as the overview predicted (Docker VM 9.70 GiB vs the 12 GB floor) —
expected, not a failure.

## Phase 3 — signal delta (tok branch)

No metric delta: a tok does not move the gate. Signals:

- Strategy recorded as **`TOK-01`** in `../decisions.md` (milestone-root).
- **Open questions closed: 3 of 4.** (a) the parallel-lane enabler is **priced and off the critical path**
  (D1); (b) `pt-world` supports **no** pre-onboarding state (audit F5); (c) the org-admin writes **do** have
  read-back surfaces — every one of the four curated UCs declares a persist-then-observe final (D2). The
  fourth (how a negative control is produced) is deliberately left to the tik that owns it.
- **Measurement surface exists** where there was none: a live local `demo-2`.
- **Risk retired:** the milestone's largest planned item (a cookie-scoped Clerkenstein registry / per-worker
  fake-FAPI) is answered rather than attempted.

## Close — 2026-07-28

**Outcome:** TOK-01 authored; the milestone's headline lever proven **off the critical path** by the gate's
own re-cut, replaced with an evidenced per-test lever (12 login sites + 8 unfenced harness violations); the
cluster order inverted to put org-admin first because it discharges clause 2 and clause 3 together; a live
local `demo-2` stood up as the measurement surface; and this iter's own D4 refuted in-iter by the Phase-0b
gate.
**Type:** tok (bootstrap)
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (bootstrap toks never fire this exit) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (0 tiks so far) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (parallelism off the critical path, both enablers priced + routed forward), D2 (org-admin first), D3 (the `networkidle` lever, 12 login sites + 2 `goto` overrides), D4 (**REFUTED in-iter** — `actor.entitlement` is declared-only; three replacement refusal surfaces), D5 (autoverify check (d) false-alarms on LibreSSL).
**Side-deliverables:** the Phase-0b audit's inline corrections to 4 corpus docs (`playthroughs.md`,
`clerkenstein.md`, `coverage-protocol.md`, `seeding-spec.md`) — including the newly-documented Clerkenstein
single-global-seat limitation, which was a genuine corpus blind area, and a `coverage-protocol.md`
denominator correction (49 → 45). Committed with this iter but not part of its planned scope.
**Routes carried forward:**
- `PERF-M256-parallel-lane` → **Fate 3, a future release milestone.** A wall-clock (not median) lever;
  both enabler options priced in D1. Not attempted in M256.
- `FIX-M256-autoverify-fapi-libressl` → **Fate 3, a later tik of this milestone.** Give `autoverify.sh`
  check (d) a probe independent of the host TLS stack (D5).
- `FIX-M257-content-stories-pair-count` → **Fate 3, M257/M258** (they compose the sweep).
  `run-content-stories.sh` recomputes 47 against the pinned 45 and `sys.exit(2)`s — the sweep refuses to
  start (audit Gap 7).
- `DOC-M256-ptworld-reset-comment` → **Fate 3, a later tik.** `pt-world.seed.yaml`'s header claims the
  showcase world is "not touched by pt-world's reset"; `doReset` takes no org filter, so it is (audit F6).
**Lessons:**
1. **When a gate is re-cut, re-derive its levers — do not inherit them.** D-v28-12 changed clause 1 from a
   suite wall-clock to a per-test median, and every downstream statement about parallelism silently became
   false. The overview still asserted the enabler was mandatory. A re-cut gate should trigger a re-read of
   every lever that was justified against the old one.
2. **A capability declared in an index is not a capability present in the data.** `seed-worlds.yaml` names
   `pt-free`, `free` and `entitlement-gated`; no seeder writes a tier, and the validator that exists to
   catch exactly this **fail-opens** because it resolves the name. Prefer reading the *writer* over the
   *declaration*. (Generalises beyond this milestone → recorded in `corpus/ops/demo/playthroughs.md`.)
3. **A guard whose verdict depends on the host toolchain is not a guard.** autoverify's FAPI check calls a
   working stack dead on any LibreSSL host — the same failure shape as M236's BSD-`date` green-gate bug.
   When a probe fails, establish *which* client failed before believing the subject failed.
