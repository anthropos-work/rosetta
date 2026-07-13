# M218 — Spec notes

_Iteration-protocol-specific technical notes: the harness contract, per-leg attribution numbers, measured baselines._

## Pre-flight audits — iter-01

**`/developer-kit:audit-kb-fidelity --milestone=M218` → YELLOW** (2026-07-13). No report file written (audit-only
run). Not a blocker; M218 may code. Findings folded into iter-01's plan:

**Confirmed as declared** — the 3 stale claims (M43-D5 "~2–5 s", the CI-inert alignment claim, the clerkenstein
header) and the 3 blind areas (`latency-budget.md` absent; the post-303 login half undocumented; the clerk-js proxy
timeout/cache contract). All 3 blind areas are properly pre-discharged by `Delivers → knowledge/…` lines ⇒ YELLOW,
not RED.

**Anchor corrections the audit found (the milestone's own citations are off — use these):**
- M43-D5 second site is `cockpit-spec.md:` **`187`**, not `155` (155 is the deep-link catalog).
- The CI-inert claim also lives at `alignment_testing.md:` **`232`** and **`239`**, not only `233`.
- The clerk-js blind area is **narrower** than stated: the CDN egress *is* documented
  (`clerkenstein.md:140-141`, incl. `FAKE_FAPI_CLERKJS_CDN`). What is genuinely undocumented is the **unbounded
  timeout** (`server.go:187`, `http.Get` = `DefaultClient` = `Timeout: 0`) and the **absence of a server-side
  cache** (only a response-side `Cache-Control: max-age=3600` at `:194`).

**The finding that changed the milestone** — `corpus/ops/demo/frontend-tier.md:240` asserts next-web's SSR fetch
origin **is** the build-time `NEXT_PUBLIC_*` URL, directly contradicting C-1's premise (which claims SSR picks up
the *runtime* env). The audit correctly refused to call it either way and flagged that **leg 1 as written could not
discriminate them**. iter-01 strengthened the probe and settled it: **the doc is right, the milestone was wrong**
(see iter-01 D1). `frontend-tier.md:240` must be **preserved**, not "corrected."

**Also newly surfaced:** `cockpit-login.ts::loginAs:64` uses `waitUntil:'networkidle'` — the pattern the milestone
**bans**. The "reuse it, don't fork it" instruction and the ban collide; resolution in iter-01 D4 (parameterize the
wait strategy, don't fork).

---

## Measured baseline — iter-01 (billion, demo-1, rext `cue-to-cue-m217`, autoverify GREEN)

**Environment matters and must always be stated** (iter-01 lesson 3): the same defect costs **~6 s on a laptop** and
**~112 s on the tailnet VM**. A latency number without its environment is not a measurement.

| leg | measured | note |
|---|---|---|
| click → handshake → 303 | **17 ms** | provably free — confirms the static analysis |
| SSR GraphQL fetch (baked public URL, from inside container) | **10,564 ms → `UND_ERR_CONNECT_TIMEOUT`** | undici default connect timeout; DNS is fine, TCP connect blackholes |
| SSR GraphQL fetch (`http://graphql:8080/graphql`, control) | **94 ms, HTTP 200** | what it *should* cost |
| SSR GraphQL fetch (runtime-env `localhost:5050`, **not used by SSR**) | 64 ms ECONNREFUSED | red herring — SSR reads the build-inlined literal |
| clerk-js from CDN (inside fake-FAPI) | **0.17–0.19 s** | healthy; C-5 is not today's cause |
| Cosmo router | 12 log lines, **0** errors/retries | federation never exercised — the SSR fetch dies upstream |
| `billion` RAM | 7.3 GiB total / **5.3 GiB available** | below the 12 GiB doc floor, but not starving |

**Derived (not yet instrument-measured — iter-02's job):**
`3 attempts × 10.5 s + (2 s + 4 s backoff) ≈ 37.5 s` per blocking fetcher; **× 3 rethrowing fetchers ≈ 112 s** →
matches the user's "1 or 2 minutes."

**p95 click→ACCESS: UNMEASURED.** No instrument exists yet. Building it is iter-02.

---

## Harness contract (constraints discovered in iter-01 — binding on iter-02)

1. **curl cannot drive this flow.** The fake-FAPI **validates `redirect_url`** against the public origin (HTTP 400
   on a loopback), and next-web's middleware **307s** any non-https origin. The harness **must** drive the real
   https origin in a real browser. (Also the only way to see the clerk-js client-gate leg.)
2. **Never gate on `networkidle`** — next-web holds never-idle long-polls. Poll for **content presence**.
   `cockpit-login.ts::loginAs:64` currently violates this; parameterize its wait strategy rather than forking it.
3. **Gate every measurement on `autoverify.json` green** — at its **real** path
   `rosetta-extensions/demo-stack/stacks/<project>/autoverify.json` (**not** `<stack>/autoverify.json` as
   `verification.md:207` claims — see iter-01 D3).
4. **New `stack-verify` surface, not a Playthrough** — Playthroughs declare perf an explicit **non-goal**. Building
   this surface discharges **DEF-M215-03(b)**.
5. **State the environment with every number.** See the lesson above.
