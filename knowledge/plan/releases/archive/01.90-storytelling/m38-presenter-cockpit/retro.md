# M38 — Retro

## Summary
M38 — the **LAST milestone of v1.9 "storytelling"** — turned the multi-identity *capability* M37 shipped into a
clickable **presenter cockpit**: a standalone host-native served panel (rext `demo-stack`, offset port
`7700 + N·10000`, never an in-app overlay — the zero-platform-repo-edit line holds) that reads the cockpit
manifest the seeder projects from the **same** `stack.stories.yaml` that seeded the data (`stackseed
--cockpit-export` — D9 single-source without PyYAML, D2) and lists each story → its hero trio with **[Login as]**
+ **[Jump to section]**. The two actions collapse into **one FAPI handshake redirect**
(`?__clerk_identity=<key>&redirect_url=<jump_to>`, M37's seat-switch seam, D3) — the hero becomes active
everywhere AND the browser lands on her screen in one move. M38 ships the **roster-export producer** (`stackseed
--roster-export` → `FAKE_FAPI_ROSTER`, single-sourced from the seeder's own id-derivation so "login as Maya"
authenticates the real seeded user, D1) + the **O9 deep-link catalog** (`DeepLinkCatalog()`). Gated on
`DEMO_STORIES=1`; default-off keeps every existing demo byte-identical (D4); fails loud on a broken roster,
non-fatal on the cockpit serve (D6). Tooling + docs only — zero platform-repo edits. Close GREEN, 8 findings,
0 blocking, merged into `release/01.90-storytelling`. **The release is now complete** — next is
`/developer-kit:close-release`.

## Incidents This Cycle
- **P2 (close-found, fixed inline) — a three-write lockstep gap in the M38-D8 fix.** The close re-fated M38-D7
  (all 6 heroes exported `org_role=admin`) from route-to-close to **LAND-NOW**. A crashed *prior* close attempt
  had begun the fix — adding a `roleForHero` helper + swapping the `users.go` seed-loop call-site — but never
  finished: `roster.go` (the exported `org_role` claim) still called the OLD `roleForIndex` directly, so a
  manager hero would have exported `org_role=member` while the seeder wrote `admin` — re-introducing the exact
  divergence the fix exists to prevent. The close-review's **code-quality + adversarial scans both
  independently caught it**, and the test (`roster_test.go`) was itself asserting the old function, masking the
  regression. Fixed: both call-sites on `roleForHero`, plus a dedicated regression
  (`TestBuildRoster_OrgRoleVantageFaithfulAndLockstep` + `TestRoleForHero`) pinning the vantage mapping AND the
  three-write agreement. No flakes (5/5 Go + Python).

## What Went Well
- **The close caught the crashed-attempt's partial.** A half-landed fix is more dangerous than no fix — it
  builds and the old test passes, so it looks done. The parallel code-quality + adversarial scans converging on
  the SAME `roster.go:93` lockstep gap (and on the masking test) is exactly the cross-cutting review the close
  exists for. Completing it properly — both call-sites + a lockstep regression test — is the clean outcome the
  three-fate rule's "no partial landings" discipline demands.
- **The single-source `roleForHero` seam is the right shape.** One helper that BOTH the `UsersSeeder` (membership
  row + casbin grant) and `BuildRoster` (the roster claim) call means the three writes agree *by construction* —
  the same single-source discipline the M37 roster-id contract uses, now extended to the role.
- **Vantage-faithful `org_role` is what the release is about.** An "employee" demo seat reading as `member` (not
  org-admin) in her JWT is the whole point of the Stories & Heroes vantage model — landing it in the final
  milestone makes the employee-vs-manager story faithful end-to-end.
- **Default-off kept the blast radius zero.** The whole storytelling layer gates on `DEMO_STORIES=1`, so the 5
  Clerkenstein alignment gates + every existing demo stayed byte-identical through M38; the cockpit consumes
  M37's existing `clerk-multi-1` DNA (9/9), no new alignment surface to gate.

## What Didn't
- **The crashed prior attempt left a dirty tree the close had to reconcile.** The uncommitted `users.go` edit
  (a partial Fate-1 start) + the original `storytelling-m38` tag cut at code that did NOT include the fix meant
  the close had to (a) recognize the partial as its own work to complete, not a foreign change, and (b) re-cut
  the tag after committing the corrected code. A clean re-run is fine, but it underscores that a partial fix
  landed mid-crash is a latent trap — the masking test would have shipped the divergence silently if the close
  scans hadn't caught it.
- **The `clerk-express-1` gate remains env-fragile** (genuine `@clerk/express` SDK needs installed npm modules;
  unrunnable in the authoring copy). Not an M38 regression (M38 never touched it), but the same recurring
  run-it-where-deps-are friction the v1.1 CI-wiring carry-forward tracks.

## Carried Forward
- **`/developer-kit:close-release` (the immediate next step):** v1.9's release-level review of all 5 milestones
  (M34–M38) as one PR + the release deferral re-audit, then merge `release/01.90-storytelling` → `main` + tag
  `v1.9`. Owns the release-scope completeness check.
- **Push the ext tags to `origin`** (from the v1.8 carry-forward, still open): `storytelling-m34..m38` (+ the
  prior `understudy`/`house-lights`/`stage-door`/`prop-room` tags) — the authoring-copy tags exist locally;
  pushing them is the user-authorized follow-up.
- **The live field-bake on a freshly-emptied `stack-demo/`** + the literal browser-pixels end-to-end of the
  cockpit (needs a `DEMO_STORIES=1` re-deploy — a deliberate demo step, orchestrator-confirmed not an M38 gap).

## Metrics Delta
- **Go test funcs:** stack-seeding **444** (`Test`+`Fuzz`; +2 at close — the M38-D8 vantage-faithful + lockstep
  regression). rext total **1248**. `go vet`+`gofmt` clean.
- **Python tests:** demo-stack **166** (+1 close — the cockpit empty-key defensive-skip), stack-injection **117**
  (8 opt-in skipped). All green.
- **Flake count:** 0 (gate 5/5 Go shuffled `-race` + 5/5 demo-stack cockpit + 3/3 stack-injection).
- **Alignment gates:** 100%/100% on all 5 Clerkenstein surfaces (re-run at close; M38 added none — consumes the
  existing `clerk-multi-1`).
- **Supply-chain:** GREEN (stdlib-only; 0 third-party deps added).
- **Code:** rext `rosetta-extensions` @ tag `storytelling-m38` (`237bede`). Full record: `metrics.json`.
