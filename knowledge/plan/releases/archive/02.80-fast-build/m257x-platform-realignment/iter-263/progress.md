# iter-263 — progress

**Type:** tik
**Active strategy:** `TOK-08`.

## Phase B — the instrument, run on both targets

`stacksecrets check --dna secretdna/secret-dna.json --from .agentspace/secrets`, with and without `--demo`.
Same DNA, same source, one flag apart:

| | `platform` genes met | `INVITATION_HMAC_SECRET` | **Critical** | gate |
|---|---|---|---|---|
| **dev** (no flag) | **13/29** | **SHORT — named explicitly** | **92.3 %** | `check` exits **1** (`< 100%`) |
| **demo** (`--demo`) | **16/29** | **not short** | **100.0 %** | exits **0** |

The 13 → 16 delta is exactly the minted + demo-generated families (`INVITATION_HMAC_SECRET`,
`NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY`, `VITE_CLERK_PUBLISHABLE_KEY`) reporting satisfied without the source.

**That single flag is the whole story.** The demo path is at **Critical 100 %** on the same secret source
that leaves a dev stack at **92.3 %** — so the demo could never have surfaced this, and the dev path had a
check that says so in one command.

## Phase C — grading, and the correction

| | prediction | outcome |
|---|---|---|
| PR-1 | the DNA declares it critical + required | **HELD** — `secret_dna_json_test.go:141-161` pins criticality, status, target file, scope and operators |
| PR-2 | the secret source lacks it | **HELD** — 15 keys, none of them this one |
| PR-3 | `check` on a dev target reports the gap | **HELD** — named in the `platform` rollup; Critical **92.3 %**, so the verb exits 1 |
| PR-4 | on a demo the same gene reports satisfied | **HELD** — and it moves Critical to **100.0 %**, which is why three green demo cycles never hinted at it |
| PR-5 | `D-M257x-262-3`'s framing is refuted by the tooling's own comment | **HELD — my published claim was wrong** |

## Close — 2026-08-10

**Outcome:** **`D-M257x-262-3` is corrected.** The `INVITATION_HMAC_SECRET` failure was **not** an
undeclared boot requirement and **not** a new failure class. The DNA declares the gene **critical +
required** with a test pinning it; `secretdna/demo.go:47` names *"the silent `app Exited (0)` class"*
**verbatim**; `corpus/ops/secrets-spec.md` documents it in three places. **The real cause is a secret-SOURCE
gap plus an unrun check** — iter-262 hand-provisioned `.env` instead of driving `/stack-secrets`, and
`stacksecrets check` reports the gap by name and fails the gate at **92.3 %** critical coverage.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue

**Decisions:** `D-M257x-263-1` (the correction, with the dev-vs-demo coverage split as its evidence).

**Side-deliverables:** none.

**Routes carried forward:**
- `FIX-M257x-262-invitation-hmac-secret-undeclared` → **RE-AIMED and renamed in substance.** Nothing needs
  declaring. Two real asks remain: **(a)** add the missing critical gene to the secret source so a dev
  provision reaches 100 %, and **(b)** make the **documented dev bring-up run `stacksecrets check`**, since
  a check nobody runs is a check that does not exist. Handler: `FIX-M257x-263-dev-bringup-must-run-the-check`.
- `FIX-M257x-262-dev-path-needs-the-studio-acquisition` → **unchanged and still the highest-value route**;
  nothing in this iter touches it, and unlike the secret gap it has **no** instrument that would have caught
  it.
- `FIX-M257x-262-demo-env-append-is-not-idempotent`, `ROUTE-M257x-261-succession-projection-is-empty`, and
  all earlier → open.

**Lessons:**
1. **Before booking a defect, grep the tooling for its own name.** The phrase *"the silent `app Exited (0)`
   class"* was already in `demo.go`. One `grep INVITATION_HMAC_SECRET` over `rosetta-extensions/` — the
   check that opened this iter — would have cost seconds inside iter-262 and prevented a wrong record.
2. **"The docs don't say it" and "I didn't run the thing that says it" look identical from inside a manual
   bring-up.** Following a guide **by hand** silently skips the tooling the guide delegates to, and then
   every gap the tooling would have caught reads as a documentation gap.
3. **A pass/fail that flips on one flag is the most valuable evidence available.** `--demo` alone moves
   critical coverage 92.3 % → 100 %; it explains the whole demo-vs-dev asymmetry with no argument.
