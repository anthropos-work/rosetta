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
