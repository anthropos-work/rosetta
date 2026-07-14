# M220 — Progress

_Section checklist. Populated from `overview.md` § Scope.In at build time; closed by `/developer-kit:close-milestone`._

**Branch:** `m220/cue-sheet` (from `release/02.30-cue-to-cue`) · **Shape:** section

## Sections

- [ ] **S0 — The two lies the docs tell.** *(overview (a) + Delivers 3)*
  - [ ] "2 orgs" → **3** at the 4 sites (`demo-up/SKILL.md:109,153`, `demo/README.md:34`, the stale
        `seed_label` at `up-injected.sh:1081`, `stories.seed.yaml:1`). VERIFIED in code — the preset ships
        **Cervato Systems** (ai-transformation) + **Solvantis** (onboarding-ramp) + **Northwind Aviation**
        (ai-readiness). This 4-line lie is why the user believed the seeding ask was unmet.
  - [ ] **Correct the FALSE safety claim** at `tailscale-serve.md:405-407`. **VERIFIED FALSE:** ports are
        emitted as bare `"<hostport>:<target>"` pairs (`gen_injected_override.py`) with **no `127.0.0.1`
        prefix**, so Docker publishes **every demo container on ALL interfaces on EVERY demo-up, today** —
        flag or no flag — and on Linux its iptables **bypass the host firewall**. `BIND_HOST` gates only the
        two host-native servers (cockpit, ant-academy). A doc that denies a live exposure is worse than none.
  - [ ] Regression fences: a doc-vs-code fence on the org count; a fence asserting the published-port shape.

- [ ] **S1 — `corpus/ops/safety.md` Part 3: the exposure side.** *(Delivers 1+2 — BLIND AREA, BLOCKS S3)*
  - [ ] The gap, proven by grep: `tailscale|remote|expose|network|localhost` → **0 hits** in safety.md.
        Read-side (never read customer data) + write-side (never write prod) exist; **remote reach is a THIRD
        AXIS with no contract at all.**
  - [ ] State plainly what default-on makes ambient: a demo is an **unauthenticated, authz-weakened build**
        (Clerkenstein disarms token verification; the authz-skip patch is default-on; **the cockpit is a
        one-click, password-free "become any hero" launcher** — a bare GET to `/v1/client/handshake`).
  - [ ] Record the steelman FOR the flip: synthetic + public-snapshot-only data (Parts 1-2 hold unchanged), a
        tailnet is an authenticated WireGuard device mesh (per-device keys, ACL-gated, no public listener),
        and — per S0 — **the exposure delta is smaller than the docs claimed, because the LAN exposure already
        exists today.**
  - [ ] Explicit written **SUPERSESSION of v2.2 D-DESIGN-1** ("public reach is never default-on",
        `demo-up/SKILL.md:78`) — **demo path only.** Never a silent contradiction.

- [ ] **S2 — The `/demo-up` defaults table.** *(overview (b) — BLIND AREA)*
  - [ ] No enumerated defaults contract exists in the corpus; the only complete knob list is a skill
        `argument-hint`. Document all ~25 `DEMO_*` knobs: knob | default | consumer | file:line.
  - [ ] Fence: the table is checked against the parser so it cannot drift (the CLI-flag ↔ docs both-directions
        rule — a doc-promised flag with no parser entry is a false promise).

- [ ] **S3 — The remote flip: `--public-host auto`, DEFAULT-ON for demo.** *(overview (c) — D-DESIGN-3)*
  - [ ] Capability ladder — **capability-gated, never presence-probed**: `command -v tailscale` →
        `BackendState == Running` → a **dotted** `.Self.DNSName` (a dotless host is hard-refused —
        `@clerk/backend`'s `assertValidPublishableKey` rejects it) → `MagicDNSEnabled` → `tailscale serve
        status` shows no operator/sudo denial → **`tailscale cert` actually mints**.
  - [ ] **HARD INVARIANT — the fallback is not optional.** Any failed rung ⇒ **empty `STACK_PUBLIC_HOST`,
        byte-identical to today's localhost path**, plus ONE loud line naming the exact fix command. `SCHEME`
        and `BIND_HOST` both derive from the same `-n $STACK_PUBLIC_HOST` predicate, so a **half-satisfied**
        public path is **strictly worse than localhost**: every baked URL becomes `https://` against plain-HTTP
        listeners and the demo does not load at all. A laptop with no tailscale MUST stay byte-identical.
  - [ ] Opt-out: `--no-public-host`.
  - [ ] Fences: each rung RED-proven — a failing rung must **fall back**, never half-satisfy.

- [ ] **S4 — Front the cockpit on `tailscale serve`.** *(overview (e))*
  - [ ] Add `('cockpit', 7700)` to `gen_tailscale_serve.py`. Today the presenter's **entry point** is the one
        plain-HTTP, unauthenticated surface on the tailnet (`up-injected.sh` deliberately excludes it).

- [ ] **S5 — The two demo-BREAKING click paths.** *(overview (i)+(j) — escalated from M219 D4)*
  - [ ] **(i) The academy POISONS the demo session.** `:13077` runs its own **keyless** Clerk and returns
        `Set-Cookie: __session=; Expires=1970` (**deletes the demo session**) + `__client_uat=0;
        Domain=<tailnet>` — **domain-wide, not port-scoped** (cookies scope by HOST, not PORT). So **a
        presenter who clicks AI Academy is LOGGED OUT of the demo into `ERR_TOO_MANY_REDIRECTS`**, and every
        employee coverage sweep aborts. Cause: `ant-academy.sh` greps `CLERK_*` out of `platform/.env`, which
        carries **11 matching lines** — all written to `.env.local`, **last one wins**, and it is not the
        demo's minted key. Fix: use **`PK_DEMO`** (already written into next-web's `.env.local`).
  - [ ] **DoD is NOT "it paints":** log in → click the academy → **the demo session must SURVIVE**.
  - [ ] **(j) studio-desk EJECTS the presenter.** `:19000` → 302 → `:13000/login` → lands on `:13000/home`.
        Not a prod-eject (`:13000` is the demo's own next-web) — which is exactly why the sweep's prod-eject
        detector never caught it. studio-desk's auth check doesn't recognise the Clerkenstein session.
  - [ ] Fence: a presenter clicking **Anthropos Studio** stays **in Studio**. "The port answers" is not the bar.
  - [ ] M219's honest fences (`SERVES BUT DOES NOT RENDER`; the `AI Academy` marker + 400-char floor) go GREEN
        only when this lands. **They are RED by design until then.**

- [ ] **S6 — Egress: stop the demo phoning home.** *(overview (f)+(g)+(h))*
  - [ ] **(f) Clerk telemetry OFF** — `CLERK_TELEMETRY_DISABLED` + `NEXT_PUBLIC_CLERK_TELEMETRY_DISABLED`.
        Real egress from **both** frontends today (grep across rext: **0** hits — never wired). Also what makes
        Playwright's `networkidle` hang. Pure env.
  - [ ] **(g) Ad-tech egress (F-5)** — the demo attempts **Google Analytics + DoubleClick + Google Ads +
        LinkedIn Ads** on **every authenticated page load**. Measure, then kill at the injected-env/CSP layer.
  - [ ] **(h) Vendor clerk-js + bound the unbounded timeout (C-5)** — the fake FAPI proxies
        `clerk.browser.js` **live from `cdn.jsdelivr.net`** via `http.DefaultClient` = **`Timeout: 0`
        (unbounded)** with **no cache** (`clerk-frontend/server.go:187`). next-web's whole authenticated tree
        is client-gated on clerk-js ⇒ **a CDN stall is an unbounded hang on the login path** (0.2 s healthy /
        **~127 s if egress blackholes**). Serve from disk; keep the CDN proxy as a **bounded** fallback.
        **Alignment-INVISIBLE** (no gene covers `GET /npm/`) ⇒ a gate-free win.

- [ ] **S7 — The dev-side opt-in `--public-host`.** *(overview (d) — folds the reserved M216)*
  - [ ] Dev stays **opt-in** per D-DESIGN-3. Builds the flag `/dev-up` does not have today.
  - [ ] **DECLARED SCOPE-FLEX LEVER:** if this bloats it **drops back to M216**, and the release still fully
        meets the user's demo-side spec. Do not let it eat the milestone.

## Out
- Speed (M218) · the AI-readiness render path (M219).
- A "hiring" story org — **D-DESIGN-4: it does not exist and will not be built.**

## Operating rules for this milestone (learned the hard way in M218/M219)
- **ONE agent against the demo host at a time.** Two concurrent batteries corrupted M219's audit trail.
- **No detached / background scripts on the demo host.** Every orphan left running wiped a stack mid-measurement.
- **Never kill a build mid-bake** — it strands the demopatches, and the next image ships silently unpatched.
- **An empty / absent / unexecuted result is a FINDING, not a pass** (D17 — ~14 hits and counting).
- Every new fence must be **RED-proven pre-fix**. A fence that passes against the bug is theatre.

## Notes
_(append per-section notes as they land)_
