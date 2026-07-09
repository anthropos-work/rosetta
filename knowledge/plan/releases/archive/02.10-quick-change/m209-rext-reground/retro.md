# M209 Retro — rext tooling re-ground

**Closed:** 2026-07-08 · `closed-complete` · `section` · complexity medium.
**Outcome:** the tooling half of v2.1's re-ground. Re-pointed every rext tool that queried the removed `skiller`
schema or expected the skiller service/container to the merged reality (`public.*`, no skiller
container/subgraph) — across **stack-snapshot**, **stack-seeding**, and the small **shell** modules — then
built + hardened + tagged the rext authoring copy. **0 `skiller.<table>` queries in any production path;
6 Go modules GREEN, `go vet` clean, 5× flake-clean.** Zero platform-repo edits.

## Summary
A **two-repo** milestone: the re-ground CODE lives in the gitignored rext repo (`.agentspace/rosetta-extensions/`,
6 commits `00a3ec5`→`2f06e78`, +961/-339 across 74 files), tagged `quick-change-m209` (re-pointed to the
post-harden HEAD `2f06e78` at close); the rosetta branch carried planning docs only. Mostly a mechanical
`skiller.<table>→public.<table>` swap (24 seeding files, ~111 fake-Conn matchers renamed in lockstep) plus two
well-scoped non-mechanical items that got the harden depth: the **Risk-1 `Surface.VersionTables()`
digest-narrowing** (so the taxonomy cache key digests only its own 10 tables and stops thrashing on unrelated
`public`-monolith app migrations, while a structure-bearing surface still whole-schema-invalidates) and a
one-sided **`MinRows` under-capture floor**. Risk-2 (capture column list) was **verified needing no change** — the
capture is names-only and replay is REINDEX-by-name, so the `extensions.`-qualified merged column types never
surface in the tooling.

## Metrics delta (from `metrics.json`)
- **Tests:** rext Go **1745 → 1763** funcs (+18: the ~111 skiller→public matcher renames are flat; +14 harden
  funcs + a few build-phase matcher additions). 6 modules GREEN; `go vet` clean; 5× sequential flake gate on the
  two touched modules = 0 failures. `flake_count` 0.
- **Harden:** 3 passes, 14 test funcs, **0 bugs**, 0 flakes — deepened edge/boundary/wiring/regression DEPTH on
  the two new risk items + a **literal-value revert guard** on the load-bearing schema const (the pre-existing
  identity check `s.Schema != Schema` was tautological and could not catch a flip back to `"skiller"`).
- **Close findings:** 1 (nice-to-have, routed — no must/should-fix). **Deferral audit:** GREEN (10 in scope, 5
  single-own + 5 inherited-M208, 0 repeat, 0 aged-out, 0 escape-hatch).

## Incidents this cycle
- **No P0/P1/P2. No regressions. No flakes.** One close-phase doc-hygiene finding surfaced and was routed (not an
  incident against M209's deliverable):
  - **TEST-1 (D-close-2):** `stack-seeding/README.md` quotes "496 test funcs / 8 packages"; actual ~788 / 13.
    **Pre-existing** cross-release drift (last reconciled at M41 / v1.10; accumulated across v1.10b + v2.0 + v2.1).
    M209 did NOT touch the README, and the rext repo was **mandate-frozen at `2f06e78`** for this close — so it was
    **recorded + routed** to the next rext advance (the v2.1 rext roll at `close-release`, or an M211 rext re-tag),
    not fixed in-place. Nice-to-have; not load-bearing.

## What went well
- **The two non-obvious risks the design flagged were both closed cleanly.** Risk-1's digest thrash got a real
  fix (VersionTables, threaded through capture + autoprovision + target-probe so both sides of the cache-key
  comparison always match) with harden anchoring it on the REAL surfaces; Risk-2 was investigated to ground truth
  (schema-only inspection of merged prod) and correctly concluded *no change* — resisting a speculative edit.
- **Done-bar met exactly.** 0 `skiller.<table>` queries verified by grep; the schema flip flows through `Surface()`
  into capture WHERE + payload filenames + manifest + replay COPY target from a single const.
- **Clean deferral discipline.** The recapture was genuinely investigated for Fate-1 (a real capture) and found
  operationally infeasible (no local COPY-byte source) → Fate-3 to M211 with a pre-surfaced-prerequisite note so
  M211's first tik doesn't re-discover the blocker. No repeat-defer, no escape-hatch.

## What didn't (go as smoothly)
- **The recapture couldn't run here.** The tooling is READY to capture/replay `public.*`, but no valid COPY-byte
  source is provisioned locally (values-blind-checked — no `marco_read`/prod-read DSN; merged `stack-dev` PG holds
  0 taxonomy rows; the `postgres` MCP returns JSON not COPY bytes). A data prerequisite, correctly M211's.
- **The two-repo shape adds a close-time constraint:** the rext repo must end frozen at `2f06e78` (tag there, clean
  tree), which blocked an in-place fix of the pre-existing README drift — surfaced as the reason TEST-1 was routed
  rather than reconciled.

## Carried forward
- **Recapture** the merged-prod `public.*` taxonomy (~42,790 skills; MinRows floor auto-catches under-capture) →
  **M211** (Fate-3; `overview.md` "Pre-surfaced recapture prerequisite" pinned).
- **KB-1/2/3** (rext-facing corpus doc bodies: `snapshot-spec.md` / `seeding-spec.md` / `safety.md` firewall row)
  + the conceptual bare-word skiller comments → **M210** (Fate-2, the chartered corpus body-flip). The MinRows /
  VersionTables invariants blend into `snapshot-spec.md` there, not now (would collide).
- **rext `stack-seeding/README` count drift** (TEST-1 / D-close-2) → the **v2.1 rext roll** at `close-release`, or
  the next M211 rext re-tag.
- **The `v2.1` rext roll + `.agentspace/rext.tag` consumption re-pin** remain **`/developer-kit:close-release`'s**
  job (the last milestone), per the release plan — NOT done at M209 close.
