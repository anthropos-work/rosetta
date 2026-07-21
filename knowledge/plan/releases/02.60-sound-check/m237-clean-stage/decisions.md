# M237 — Decisions

## D1 — Freshness fix is VISIBILITY, not auto-mutation (§1)
The clone-freshness fix asserts + warns loud + records a pin model; it never auto-advances. Rationale: deliberate staleness at a known-good rev is legitimate (reproducing a demo; an unconditional pull could break one mid-presentation). Advance is opt-in (`DEMO_ADVANCE_CLONES=main|pinned`); a strict mode (`DEMO_FRESHNESS_STRICT=1`) escalates to fatal for a CI / HARD go/no-go bring-up. The verified fetch NEVER suppresses stderr — a failed fetch records `behind=null` (UNKNOWN), never a fabricated number.

## D2 — Real pin model: 7 distinguishable states (§1)
`clones.lock.json` now records `{ref, sha, pin_state, behind, fetch_ok, pinned_at}`. `pin_state` ∈ {`pinned-tag` (at an exact tag), `pinned-detached` (detached HEAD), `pinned` (branch matches `clones.pin.json` declaration), `pin-drift` (branch ≠ declaration — LOUD), `fresh` (0-behind), `stale-by-neglect` (behind>0, undeclared — LOUD), `unknown` (fetch failed — LOUD)}. The optional operator declaration `stack-demo/clones.pin.json` = `{"<repo>": "<ref>"}` is what makes "pinned deliberately" distinguishable from "stale main" (today both read `ref:"main"`).

## D-HARDEN-1 — shell coverage = explicit branch/path enumeration (no kcov/bashcov)
The M237 harden measured coverage as **branch/path enumeration** of `ensure-clones.sh`'s assertion +
pin-state + advance + sweep logic, not a coverage-tool percentage. Rationale: the touched code is a bash
script exercised via a Python subprocess harness; `kcov` needs Linux/DWARF (does not instrument bash on
macOS) and `bashcov` needs a Ruby toolchain — neither is installed, and this 730-test suite has *always*
used branch-mapping. Instrumenting a new shell-coverage stack for one 427-line script is disproportionate;
the compensating strategy is the enumerated branch table in `progress.md`'s hardening section (each pin
state + advance arm + R1 error path mapped to a named test, two of them mutation-verified). Aligns with the
harden-milestone "compensating test strategy" clause.

## D-HARDEN-2 — fetch-OK-but-uncountable classifies as `fresh` (observed edge-semantic; ACCEPTED)
A **successful** fetch whose `rev-list --count HEAD..origin/<ref>` cannot compute a count (no matching
upstream branch — e.g. a local-only branch) records `behind=null` and classifies the repo **`fresh`**
(pinned by `test_fetch_ok_but_uncountable_leaves_behind_unknown_never_fabricated`). This asserts a green
"fresh" without a positive behind==0 measurement — a mild tension with the milestone's honesty spine.
**Fate:** ACCEPTED as current behavior, **not** changed in this harden pass. Rationale: (a) reclassifying
would change the 7-state model (a design decision, not a hardening fix); (b) in production every demo clone
tracks a branch with a real `origin/<ref>`, so a *successful* fetch yields a countable divergence — the
edge needs a branchless/upstreamless clone to trigger, which the bring-up does not produce; (c) the honesty
invariant that actually matters (never fabricate a *number*) IS upheld — `behind` is null, never guessed.
Pinned by a test so any future change is deliberate + visible. Flagged for `/developer-kit:close-milestone`
review; no downstream milestone owns it (not a defect, an accepted semantic).

## D3 — R1 directory-driven, not a hand array (§2)
R1 (`ensure-clones.sh`) now iterates `patches/*/*.yaml` (all 14) instead of a hard-coded 3-entry array. `revert --force-pristine` only restores-to-pristine (never applies), so sweeping all is safe by construction; a manifest demopatch refuses logs a non-fatal skip. Test harness refactored into `EnsureClonesHarness` mixin so the §1/§2 suites reuse it without re-running the functional suite (no triple-execution).

---

# CONFIRMED-DEFECT LEDGER — fresh-build re-triage on `billion` (2026-07-21)

**The §3/§4 deliverable.** Re-triaged the ambiguous v2.5 UI defects on the `billion` demo-1 stack (offset +10000; MagicDNS `billion.taildc510.ts.net`), whose clone freshness was VERIFIED by the M237 §1 assertion. Browser observation via the e2e cockpit-login harness (`loginAs`, seats `dan-manager` / `maya-thriving`) run from the workstation against the tailnet HTTPS origins (the presenter vantage). Evidence: screenshots + DOM/network/console dumps.

## HEADLINE — the "202-behind" premise is REFUTED (the barrier's real payoff)
The design/`rosetta_demo.md` premise that billion "builds months-old code (~202 commits behind)" is **refuted by the fetch-verified measurement**, independently confirmed by raw `git rev-parse HEAD == origin/main`:

| repo | verified behind | state |
|---|---|---|
| platform | 0 | fresh (at origin/main head `0808b92`) |
| next-web-app | 0 | pinned-tag `v2.113.1` == origin/main; **demo-1's frontend image was BUILT from this exact current commit** |
| app | 2 | pinned-tag `v1.341.0` |
| cms / jobsimulation / messenger / studio-desk / sentinel / skillpath / storage / roadrunner | 0 | fresh / pinned-tag |
| graphql-wundergraph | 2 | pinned-tag |
| **ant-academy** (separate clone, not in repos.yml) | **5** | on `main`, no tag — **the one genuinely-stale surface** |

So the frontend was **already current** — the "202-behind" reading was the confident-wrong **suppressed-fetch artifact** the §1 fix eliminates (the fix proved its own worth: old method said 202-behind, verified method says 0). The "clean stage" barrier's finding: **the stage was cleaner than believed** (the assumption of pervasive staleness was itself the bug), with **ant-academy the lone stale surface**. The re-triage is therefore trustworthy for #1/#4 (verified-current next-web) and correctly routes #2 to the academy owner (M238). *(Scope: measured on billion; the local macbook clone set was not measured here.)*

## Ledger

- **#1 — flat menu → hierarchical for managers: RESOLVED (does not reproduce).** On the verified-current build, the manager (`dan-manager`, landed on `/enterprise/workforce`) left menu renders **hierarchical**: a "Content Library" group (AI Simulations / Skill Paths / AI Academy) + an "Organization" group with **5 expandable sub-sections** (MAP, CUSTOMIZE, ASSIGN, TRACK & VERIFY, INTELLIGENCE — each a collapsible group with a `>` chevron + "Expand all"). The probe counted `expandables: 5`; the screenshot confirms the hierarchy. The Workforce Intelligence page itself is richly populated (221 members, 163 skills mapped, 130 verified, charts, people lists). → **CLOSED by the fresh build; no downstream fix.** (Also discharges M239's "confirm #1 hierarchical menu renders for managers".)

- **#4 — /library/skill-paths empty on first load: does NOT reproduce as empty.** On the verified-current build, `/library/skill-paths` (`maya-thriving`) renders **populated**: "Public Content (22)" with a full skill-path card grid + full category tree. Probe: **7 card-elements within 1.2s → 29 settled**; NO empty-state message; **NO GraphQL errors; NO HTTP 4xx/5xx**. A sub-second empty *flash* at <1.2s could not be definitively confirmed or ruled out, but the page is functional. → **Priority REDUCED.** Routes to **M239** (owns library first-load), **re-scoped down** from "library is empty" to "verify there is no cold-first-load empty flash / client-fetch race" — the library itself is not broken.

- **#2 — ant-academy language-switch error: SURVIVES (real; academy-surface defect).** On billion's academy (`:13077`, **5-behind**, the known "empty academy": 0 paths / 0 courses), the language switcher (flag control, `aria-label="Language: English"`) is **non-functional**: clicking it surfaces no working language menu; the Italian locale path `/it` returns **404 "Not Found — AI Academy"** (console 404). No crash/5xx. This is **NOT** a next-web clone-staleness artifact — it lives on the separate, independently-degraded academy app. → **Routes to M238** (academy reliability; the design-notes M238 In-list already owns "Fix #2 (language error)"). **M238 must re-verify the language switch on the FRESH academy it delivers** (billion's academy is 5-behind; M238 should advance it — the M237 `DEMO_ADVANCE_CLONES` tooling is the lever).

## Downstream routing (Fate-2 — these milestones already own the surviving defects)
- **#2 → M238** (academy reliability). Already in M238's `In:` list. Add: re-verify on a fresh (advanced) academy.
- **#4 → M239** (enterprise surfaces / library first-load), re-scoped to the cold-flash/client-race question. Already in M239's `In:` list.
- **#1 → CLOSED** here (also confirms M239's hierarchical-menu check).
- **ant-academy 5-behind → M238/M244**: the freshness tooling (`DEMO_ADVANCE_CLONES`) should advance ant-academy when M238/M244 rebuild on billion.

## Live-proof of the M237 code on billion (dogfood)
- §1 freshness assertion RAN live on billion (`ensure-clones.sh` standalone): verified fetch SUCCEEDED for all repos (`fetch_ok: true`), new schema written, "all clones provably fresh-or-pinned" — and it SURFACED the true freshness (refuting 202-behind).
- §2 R1 swept **14** manifests live ("demopatch R1: swept 14 manifest(s) … directory-driven — F-M236-CLOSE-2").
- rext consumed at tag `sound-check-m237-clean-stage` (re-pinned + `.agentspace/rext.tag` updated); both the authoring copy and the billion consumption clone left git-clean.

---

# Adversarial review (close, 2026-07-21)

The module reviewed is `ensure-clones.sh` phases d3/e/f (the only non-trivial code M237 touched). One scenario
considered beyond the harden pass's branch enumeration:

## AR-1 — the behind-count capture merges stderr into the value (`_behind="$(… 2>&1)"`)
**Scenario.** Phase (e) captures the divergence with `_behind="$(git … rev-list --count "HEAD..origin/$ref" 2>&1)"`.
The `2>&1` is there so a failure message rides into the log. But on the SUCCESS branch it means any stderr git
emits *while succeeding* (e.g. an advisory config warning) would be concatenated with the integer — `_behind`
becomes a multi-line non-integer.
**Does the code handle it?** Yes — and it fails toward honesty, not toward a fabricated number. (a) The
classifier guards the arithmetic with `[ "$_behind" -gt 0 ] 2>/dev/null`, so a non-integer can never be read as
"behind>0". (b) The lockfile writer does `int(behind) if behind != ""`; a non-integer raises `ValueError`, and
the script runs under `set -euo pipefail` (line 27), so the run **aborts loudly** rather than persisting a wrong
`clones.lock.json`. The honesty invariant that matters — *never write a fabricated behind-count* — holds by
construction: a garbage value aborts the bring-up, it is never recorded as truth.
**Trigger probability:** very low (`rev-list --count` on a valid range emits only the count on success). Left
as-is: the current form is loud-not-silent and consistent with the release's honesty spine. Recorded so a future
reviewer sees it was considered.
