# M53 — decisions

_Implementation decisions with rationale (one entry per decision: context → options → choice → why)._

## KB-1 — AB2 is "prompt-free replay from a filled cache", not "auto-capture during bring-up"
**Context.** The Phase 0b audit (all 3 KB contract docs ALIGNED) clarified that `/demo-up` on a cold box
**replays only, never captures** (`up-injected.sh:665`; `snapshot-cold-start.md:110,198-211`). M47's
contribution is making the *operator's* one-time `stacksnap capture` turnkey via the MCP-configured DSN (no
`~/.pgpass`) — it does not auto-capture inside the bring-up.
**Choice.** M53 asserts AB2 as: the `.agentspace/snapshots/` cache is present + populated (verified: 1.4 GB,
taxonomy + directus + sim-embeddings each with COPY files + manifest.json), so the cold `/demo-up` set-dresses
by **replaying that cache with no prompt**. This is the accurate reading of "the cold-start MCP-DSN auto-capture
filled the snapshot with NO prompt" — the snapshot was filled (by M47's turnkey capture, already done), and the
cold rebuild consumes it prompt-free. **We do NOT wipe the cache** — that would require a live prod capture,
which is out of M53's scope and not what AB2 means.
**Why.** Matches doc+code; a truly-empty-cache assertion would contradict the (correct) replay-only bring-up
contract. Not a regression — the cache-fill is a completed M47 deliverable.

## KB-2 — AB5 asserts the shipped 78.4%/199, and validates render on the cockpit's actual link
**Context.** `ai-readiness.md:106` carries a round "~80% / ≈160 of 200" from the contract-writing phase; the
shipped funnel + `seeding-spec.md:369-375` are **78.4% / 199 frozen snapshots**. Also: the fast frozen read
path fires on a `?cycle=<closed>` deep-link, but the cockpit AI-readiness link (`cockpit.go:74`) is the bare
`/enterprise/workforce/ai-readiness` (no `?cycle=`); the M51 `app-aireadiness-snapshot-loadmembers` patch
bounds member hydration so the dashboard renders acceptably regardless.
**Choice.** M53 asserts AB5 against the shipped 78.4%/199 (1 completed + 1 started + manager), enabled/3-step,
and validates the dashboard **renders** (not a 180s timeout) on whatever link the cockpit + manager coverage
harness actually navigate to. A stale round-number in the doc prose is a doc-hygiene note (flag in §5), not an
acceptance failure.
**Why.** Assert against ground truth (code + seeding-spec), not the round contract number.

## D1 — Academy F6: authenticated session via the academy's OWN `e2e_persona` bypass (zero academy-repo edit)
**Context.** F6(iii) requires "a non-anonymous academy session (the hero lands authenticated, not anonymous)."
The current launcher (`ant-academy.sh`) runs the academy **anonymous** via `BENCHMARK_VISUAL_BYPASS=1` +
`REQUIRE_ORGANIZATION_MEMBERSHIP=0` — server-side `auth()` resolves every request as anonymous. But the academy
ships a mature **`e2e_persona` cookie bypass** (`src/lib/e2eAuth.js`, `serverAuth.js`, `clerkClientHooks.js`):
under `BENCHMARK_VISUAL_BYPASS=1` (server) + `NEXT_PUBLIC_E2E_AUTH=1` (client), an `e2e_persona=member` cookie
drives a **signed-in** context end-to-end (server RSC `anonymous=false` + entitlement + client Clerk hooks
resolving a named `E2E Member` identity — progress/certs/sidebar all active). No real Clerk keys needed.
**Options.**
  (a) Wire Clerkenstein (the demo's fake FAPI/BAPI) into the academy so a hero's minted Clerk session carries
      cross-origin from next-web → academy. Rejected: heavy, fragile cross-origin session-sharing, and it would
      need academy env/repo changes; the academy runs standalone (no platform-backend dependency) by design.
  (b) Use the `/api/dev/login-as` real-Clerk route. Rejected: it MINTS a real sign-in token via
      `CLERK_SECRET_KEY` — the demo provisions no real keys, so it 500s.
  (c) **CHOSEN:** use the academy's own `e2e_persona` bypass. The launcher adds `NEXT_PUBLIC_E2E_AUTH=1` (so the
      CLIENT persona layer activates alongside the already-set server `BENCHMARK_VISUAL_BYPASS=1`), and the
      cockpit's academy menu-link sets the `e2e_persona=member` cookie client-side, then navigates to the
      academy origin. The hero lands **signed-in as an entitled member** (non-anonymous) — F6(iii) met.
**Choice.** (c). ALL new code lives in rext (`ant-academy.sh` env + `cockpit.py`/`cockpit.go` deep-link) — the
zero-academy-repo-edit line (D15 / `test_launcher_makes_zero_ant_academy_repo_edits`) is preserved: the launcher
still only writes the gitignored `code/.env.local`, and the cookie is set browser-side by the cockpit panel.
**Identity nuance.** The `member` persona is the academy's synthetic `E2E Member`, NOT the exact seeded platform
hero (Maya/Dana). F6's bar is "authenticated, not anonymous" — `member` (signed-in + org + entitled) satisfies
it. Resolving the *exact* platform hero inside the academy would require wiring the academy backend to the
demo's platform DB (heavy, out of F6's small-surface scope). Documented as such in §5.

## D2 — Academy F6: the menu-link is a cockpit deep-link (a new NON-next-web catalog vantage)
**Context.** F6(ii) requires "a hero academy menu-link routing from the cockpit/persona into the academy." The
cockpit `DeepLinkCatalog` (`cockpit.go`) is entirely **next-web-relative** (paths joined to `app_base`); the
academy runs on a **different origin** (`http://localhost:$((3077+offset))`).
**Choice.** Add an academy deep-link to the cockpit that (a) is marked as academy-vantage (absolute academy
origin, not next-web-relative) and (b) sets `e2e_persona=member` before navigating (per D1). Rendered as a
per-story (or global) "Open the Academy (as a member)" link in the cockpit panel — the presenter clicks it to
walk into the academy authenticated. The academy origin is threaded into the cockpit at launch (the offset is
known: `3077 + N*10000`). Keeps the single-source property: the academy link is a first-class catalog entry,
not a hardcoded string scattered in the HTML.
**Why.** Mirrors the existing per-hero next-web deep-link seam while respecting the cross-origin + auth reality
of the academy; a first-class catalog entry keeps it discoverable + testable.

## D3 — Academy AI chat (Cosmo) stays absent-in-demo, now documented as a demo contract
**Context.** The overview says "the academy AI chat stays documented-as-absent (no `/api/ai/chat` assertion —
the AI assistant needs keys the demo doesn't provision)." Cosmo is gated behind `NEXT_PUBLIC_FEATURE_TRAINING_
COACH` (default OFF) + a per-user `localStorage('openai_api_key')`; the launcher sets neither and provisions no
key, so Cosmo is genuinely absent. The Phase 0b audit found this is *implied* but not stated as a demo contract.
**Choice.** Do NOT enable Cosmo in the demo (leave the flag unset). Add one explicit line to
`corpus/ops/demo/frontend-tier.md` § ant-academy stating the AI chat is absent-by-design in the demo (no keys
provisioned) — per the AI-keys policy. No `/api/ai/chat` assertion in the F6 acceptance.

## AB4-REGRESSION — the M42 manager gate (dan-manager @ Cervato) is BROKEN by an M51 unconditional seedPath → routes back to M51
> **STATUS: RESOLVED — see AB4-FIX below.** The user APPROVED fixing this at the acceptance gate (M51 is
> archived). This entry is the original diagnosis (accurate history); the fix + re-verification are in AB4-FIX.

**This is an ACCEPTANCE FINDING, not an M53 fix. Per the acceptance-not-fix rule, the fix routes BACK to
M51 (its owner); M53 does NOT repair it.**

**What the cold-rebuild both-vantage sweep found:**
- Employee vantage (maya-thriving @ Cervato Systems): **GATE MET** — reachable=59/150, failingSections=0,
  escapes=0, personaFailures=0, notReached=0, frontier=EXHAUSTED. (AB4 employee half ✅.)
- Manager vantage `dan-manager` @ Cervato Systems (**M50's canonical M42 manager gate**, `run-coverage.sh 1
  manager` default): **GATE NOT MET** — failingSections=**2**, escapes=0, personaFailures=0,
  frontier=EXHAUSTED. BOTH failing sections are `/enterprise/workforce/ai-readiness`
  (`ai-readiness-org-score` + `ai-readiness-funnel`, kind=`empty`, "meaningful text 50 < floor 80/60") — the
  page renders (HTTP 200, 0 ejects) but shows **"No AI readiness data yet for this org"** (screenshot verified;
  org header = Cervato Systems).
- Manager vantage `dana-manager` @ Northwind Aviation (**M51's canonical AI-readiness gate**, `run-coverage.sh
  1 manager dana-manager "Northwind Aviation"`): **GATE MET** — reachable=70/150, failingSections=0, escapes=0,
  frontier=EXHAUSTED; the ai-readiness sections PASS (541 meaningful chars). Dashboard renders 50/100 org
  readiness, 199 members, 173/199 (87%) functional+, the 3-step funnel + the By-team grid (screenshot
  verified). **AB5 ✅.**

**Root cause (owned by M51 iter-05 D3):** the 199 AI-readiness snapshots are seeded ONLY for **Northwind
Aviation** (in a `closed` cycle — the M51 showcase-org design; confirmed via
`ai_readiness_snapshots → ai_readiness_cycles → organizations`). M51 iter-05 D3 added
`/enterprise/workforce/ai-readiness` to `MANAGER_MANIFEST.seedPaths` (`stack-verify/e2e/lib/coverage-manifest.ts:520`)
**UNCONDITIONALLY** — so EVERY manager sweep (any org) is primed to visit the ai-readiness page and assert the
funnel renders. But the funnel data is org-specific (Northwind only). M51's own gate ran ONLY `dana-manager` @
Northwind (where it passes), so M51 never re-ran the M50 `dan-manager` @ Cervato manager sweep and **never saw
the regression its unconditional seedPath introduced** into the M50 gate. M50 (which closed BEFORE M51) had a
GREEN manager gate (`dan-manager` @ Cervato, reachable=69, failingSections=0) — because the ai-readiness
seedPath did not exist yet.

**Why this is exactly what M53 exists to catch:** the M53 risk line — "a cold-rebuild surfacing a regression
late" — is realized here. The fix-on-live serialization across M47..M52 never re-ran the M50 Cervato manager
sweep after M51's manifest change; M53's from-cold both-vantage assertion is the first time both the M50 gate
(dan-manager @ Cervato) AND the M51 gate (dana-manager @ Northwind) are re-measured together, surfacing the
cross-milestone drift.

**Routing (M51 fixes — NOT M53):** the ai-readiness seedPath + its two section descriptors must be made
**org-conditional** — asserted only for the showcase-org manager vantage (Northwind / `dana-manager`), not
primed for every manager org. Options M51 should weigh: (a) gate the seedPath on `COVERAGE_EXPECTED_ORG ===
"Northwind Aviation"` (or an `expectedOrg` field on the descriptor); (b) move the ai-readiness descriptor into
a Northwind-only manifest overlay; (c) tag the descriptor `showcaseOrgOnly: true` and have the sweep skip
showcase-org-only sections when the logged-in org isn't the showcase org. The correct, complete fix is M51's
to design (the three-fate rule: this is NOT a "land 20%" in M53 — it is a full manifest-conditioning change
owned by M51's subject area).

**Severity: BLOCKER.** The M53 acceptance bar AB4 ("both-vantage M42 semantic coverage GREEN … on the existing
orgs (M50)") is NOT met from cold: the M50 manager vantage (dan-manager @ Cervato) is RED. This blocks the
v1.10b cold-rebuild acceptance until M51 conditions the seedPath and the M50 dan-manager @ Cervato manager gate
is GREEN again (with AB5's dana-manager @ Northwind gate still GREEN). AB1/AB2/AB3/AB5/AB6 + the academy F6 all
PASS from cold; AB4 is the sole failure.

## AB4-ROUTING — M51 is archived; the regression escalates to the orchestrator/user (not annotated in M53)
> **STATUS: RESOLVED — see AB4-FIX below.** The escalation ran as designed: the regression was surfaced to the
> user WITH the evidence + fix design, and the user chose the routing (fix at the gate, not re-open M51).

Per the acceptance-not-fix rule, a failed acceptance assertion routes to its owning milestone and M53 does
NOT fix it. M51 (`status: archived`) is that owner (its iter-05 D3 added the unconditional seedPath). Because
M51 is already closed+archived, the correct mechanism is NOT to silently re-open it or land a fix in M53 —
it is to **STOP with SEVERITY=blocker and surface the regression to the orchestrator/user** with the evidence
+ the fix design, who decides the routing (re-open M51 for a targeted manifest-conditioning fix, or schedule
it as a tracked follow-up before the v1.10b release closes). The v1.10b cold-rebuild acceptance is NOT green
until AB4's manager half (dan-manager @ Cervato) is GREEN again with AB5 (dana-manager @ Northwind) still
GREEN. This is a genuine release-blocker, correctly surfaced by M53's from-cold both-vantage assertion — the
milestone did its job (caught a late cross-milestone regression the live-serialization missed).

## AB4-FIX — org-condition the manager AI-readiness manifest, fixed at the acceptance gate (a recorded M53 exception)
**Context.** The AB4-ROUTING escalation surfaced the M51-owned regression to the user with the evidence + fix
design. Because M51 is archived, re-opening it for a one-line manifest change is heavier than the fix itself.
**The user APPROVED fixing it at the M53 acceptance gate** — a conscious, recorded exception to M53's
"acceptance-not-fix / no fix code" rule, in exactly the same class as the academy F6 exception (D1/D2): the
touched code is a rext **test/gate artifact** (the coverage manifest), NOT platform code, and the fix is
narrow + fully test-covered.

**Options (from AB4-REGRESSION's routing note).** (a) gate the seedPath on `COVERAGE_EXPECTED_ORG === the
showcase org`; (b) a Northwind-only manifest overlay; (c) a `showcaseOrgOnly: true` descriptor tag the sweep
skips off-showcase. **Chosen: (a)**, the smallest correct change that reuses the org name the sweep already
threads (`COVERAGE_EXPECTED_ORG`, the same value persona-assert's `orgIdentity` matches on).

**Change (rext authoring `117fe41`, `stack-verify/e2e/`; zero platform edits).**
- `AI_READINESS_SHOWCASE_ORG = 'Northwind Aviation'` — the org whose manager vantage carries the AI-readiness
  surface (the 199 snapshots seed only there).
- Extracted the AI-readiness page into an exported `AI_READINESS_PAGE` descriptor.
- `MANAGER_MANIFEST` (unchanged name/role) = the SHOWCASE manifest: base pages + `AI_READINESS_PAGE`, seedPaths
  including `/enterprise/workforce/ai-readiness`. Kept canonical so AB5 + the existing manager unit tests
  (which reference `MANAGER_MANIFEST`) still see the full surface.
- New `MANAGER_MANIFEST_BASE` = the base-org manifest: the SAME surface MINUS the AI-readiness seedPath +
  descriptor.
- `manifestFor(vantage, expectedOrg?)` returns the showcase manifest only when `expectedOrg` matches
  `AI_READINESS_SHOWCASE_ORG` (case-insensitive substring); otherwise (base org, or empty/undefined — the
  manager default) the base manifest. Employee vantage returns `EMPLOYEE_MANIFEST` unchanged.
- `coverage.spec.ts` threads `expectedOrg` into `manifestFor(vantage, expectedOrg)`.
- +3 unit tests (showcase includes the page; base/Cervato/Solvantis/empty omit both seedPath + descriptor; no
  collateral drop — base = showcase minus exactly one page). **27/27 manifest unit tests pass.**

**Why base org omission is correct (not a coverage hole).** On a base-Workforce org the AI-readiness dashboard
is LEGITIMATELY empty (no seeded cycle) — asserting it there is a false-fail. The page's real proof (the funnel
renders from real seeded data) is still asserted on the showcase org (Northwind, AB5). So the fix removes a
false assertion, not a real one; escapes stay 0, persona stays green.

**Re-verification (both manager vantages, same cold demo-1, run-coverage.sh at the re-rolled v1.10.1):**
- `dan-manager` @ Cervato (M50 base gate): **GATE MET** — reachable=69/150, failingSections **2→0**, escapes=0,
  persona=0, frontier=EXHAUSTED; ai-readiness NOT in the reached set (base manifest omits it).
- `dana-manager` @ Northwind (M51 showcase gate, AB5): **GATE MET** — reachable=70/150, failingSections=0,
  escapes=0, persona=0; ai-readiness seedPath crawled (position #3) + **both ai-readiness sections PASS** (541
  meaningful chars). AB5 intact.

**`v1.10.1` re-rolled** at the fix HEAD `117fe41` (local unpushed annotated-tag re-roll — NOT a force-push);
`.agentspace/rext.tag` stays `v1.10.1`; the `stack-demo/rosetta-extensions` consumption clone re-pinned via a
clean `git fetch <authoring> main --tags && git checkout v1.10.1`.

**Verdict.** AB4 GREEN both manager vantages → the full M53 acceptance bar is **6/6 + F6 PASS from cold**.
