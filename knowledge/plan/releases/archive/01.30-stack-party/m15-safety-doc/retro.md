# M15 — Retro

**Milestone:** Safety & security doc + dual-repo knowledge consolidation · **Shape:** section · **Closed:** 2026-06-07
**Branch:** `m15/safety-doc` → merged `--no-ff` into `release/01.30-stack-party` · **Tag:** `stack-party-m15` @ `51ca18b` (extensions)
**THE LAST v1.3 MILESTONE** — with it closed, all 4 milestones (M12→M15) are done and the release is ready for `/developer-kit:close-release`.

## Summary

The **closing milestone of v1.3** — the M8/M11-analog consolidation. It delivers a single authoritative, code-cited
doc, **`corpus/ops/safety.md`** (248 lines), stating the two inviolable guarantees of the `rosetta-extensions` stack
tooling: it **never reads private/customer data** (read-side — the tenant firewall `AssertPlan`/`AssertCaptured`, the
per-surface public predicates byte-matched to `firewall.go`, the public-only data-DNA gene, bounded read-only capture)
and it **never touches production data or services** (write-side — the 3-layer isolation guard
`CheckWrite`/`PreflightEnv`/`AssertClean`, never-write shared Directus/prod-S3, the capture-source policy, the doubled
n=0-dev guards, the audit-proven zero-pollution assertion). It is dual-level (a PM "two promises in one paragraph"
section + an engineer deep-dive), cross-linked from all 4 sibling docs. Alongside it, a **dual-repo KB refresh**:
`rosetta-extensions/knowledge/` to the v1.3 converged dev≡demo model + the safety contract, and the root READMEs +
`corpus/ops/demo/` recipes to the unified `stack-*` skills + dev-as-peer. Built in 5 sections, hardened in 4 passes
(3 net-new docs↔code drift guards + 1 stabilization), closed with 1 minor self-referential docs-accuracy finding.

## Incidents This Cycle

**No defects shipped. No regressions. 0 P2 flakes.**

- **1 close finding (docs-accuracy, fixed inline, not a defect):** the milestone's own `decisions.md` M15-D4 cell
  still read "flagged, not fixed here (… noted for a future touch)" — but the harden pass had in fact landed that
  fix **Fate-1** (correcting the n=0 over-claim in both `dev-setdress.sh:19-22` and the sibling
  `provision-plan/main_test.go:179-181`). A close-time review caught the decisions.md text lagging what shipped;
  corrected to record the Fate-1 harden landing. Self-referential doc drift, no code impact.
- **1 accuracy guardrail surfaced at Phase 0b, landed in harden (M15-D4):** the `dev-setdress.sh` source comment
  over-claimed "stacksnap/stackseed independently refuse N=0 too" — `stacksnap` replay has **no** N=0 guard (and
  correctly so: replay writes only public data into the stack's own isolated stores). Flagged during the KB-fidelity
  audit as a pre-existing comment inaccuracy that contradicted the shipped safety.md §2.5; harden re-evaluated and
  landed the correction completely rather than leave a doc↔code contradiction. Caught before merge, no shipped impact.

## What Went Well

- **The doc was authored against verified facts, not memory.** Phase 0b KB-fidelity came back **GREEN** — every
  read-side/write-side claim was cross-checked against the real extensions code (`firewall.go`, `isolation.go`,
  `audit.go`, `source.go`) *before* safety.md restated it. Two accuracy guardrails (no offline-file-reader claim
  M15-D3; precise n=0 scope M15-D4) were carried into the build so the new doc got it right from the first draft.
- **Doc accuracy was made a *test*, not a convention.** 7 new fail-closed drift guards pin every load-bearing
  literal/symbol/SQL-block safety.md quotes to the real code (read-side predicates + gate names + the bounded-read
  SQL; write-side the complete Clerk-host + Directus-token rejection lists + the forced bucket override + the guard
  symbols). Future doc↔code drift becomes a test failure, not a silent lie — the durable guarantee that the safety
  contract stays true.
- **The guards were proven fail-closed, not assumed.** At close (Phase 2c) two controlled mutations — corrupting the
  read-side taxonomy predicate and renaming the write-side bucket override — each tripped its guard; safety.md
  restored byte-identical, guards green again. The guards also `t.Skip` gracefully in pinned-tag consumption clones,
  so they guard drift only where the doc is edited (the authoring copy) without coupling every stack to the corpus.
- **Clean, near-straight-through close** — deferral re-audit GREEN, 0 code/test/scope findings, the single finding a
  self-referential docs touch fixed in seconds.

## What Didn't

- **Nothing material.** One methodology note: the Go test-func count carried in `state.md` headline numbers (721 at
  M14) was produced with a looser matcher than this close's consistent `^func Test` count (713). The **delta** is what
  matters and is unambiguous — under one matcher the M14-tag baseline is 706, so M15 is **+7** (the 7 drift guards).
  state.md's prior absolute figure is left as-is for continuity; the +7 delta is the load-bearing figure and is
  recorded in `metrics.json` with the matcher caveat spelled out.

## Carried Forward

- **DEF-M10-01** (S3 media blob bytes + cloud `SnapshotStore` backend) → **v1.4**, inherited from v1.2, signed by the
  user at v1.3 design. M15 (a docs/consolidation milestone) touched no media or snapshot-store code path; the item is
  **not aged out** (all 4 aging triggers negative — verified in this close's deferral audit). safety.md documents only
  the current refs-only/local-store posture, with a clearly-labelled "Future (v1.4)" forward pointer. Re-confirmed GREEN.
- **No new deferrals introduced by M15** — every scope item landed Fate 1; M15-D3 + M15-D4 are accuracy decisions, not
  deferrals.
- **Next: `/developer-kit:close-release`** — M15 was the last pending milestone of v1.3. close-release will treat all
  M12→M15 commits as one release-level PR, run its own `--scope=release` deferral audit (this close pre-staged it
  GREEN), merge `release/01.30-stack-party` → `main`, and tag `v1.3`. That is the user's separate step.

## Metrics Delta

(Source: `metrics.json`.)

- **Go test funcs:** **+7** (the 7 docs↔code drift guards) — stack-snapshot +3 (`TestSafetyDocPredicatesMatchCode`,
  `TestSafetyDocNamesRealFirewallGates`, `TestSafetyDocBoundedReadSessionMatchesCode`); stack-seeding +4
  (`TestSafetyDocClerkHostsMatchCode`, `TestSafetyDocDirectusTokenKeysMatchCode`, `TestSafetyDocForcedBucketOverride`,
  `TestSafetyDocNamesRealGuardSymbols`). 706→**713** under a consistent `^func Test` matcher.
- **Python test funcs:** **174** (unchanged — M15 touched no Python surface).
- **Quality gates:** all 4 Go modules `-race -count=1` green; gofmt + `go vet` clean; `dev-setdress.sh`
  shellcheck-clean; py_compile clean; **flake 0 (5/5** on both touched packages).
- **Deliverable (the headline):** `corpus/ops/safety.md` (net-new, 248 lines, code-cited, dual-level) + back-links
  from all 4 siblings + the extensions `knowledge/` refresh to the converged model + safety contract. 0 stale skill
  names in the extensions KB.
- **Findings:** **1** at close (a self-referential docs-accuracy fix; 0 code/test/scope).
