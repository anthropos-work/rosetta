# M220 — Progress

_Section checklist. Populated from `overview.md` § Scope.In at build time; closed by `/developer-kit:close-milestone`._

**Branch:** `m220/cue-sheet` (from `release/02.30-cue-to-cue`) · **Shape:** section

## Sections

- [x] **S0 — The two lies the docs tell.** *(overview (a) + Delivers 3)*
  - [x] "2 orgs" → **3** at the **7** sites (the KB-fidelity audit found **3 more than the plan's 4**, and
        corrected two stale anchors — see `kb-fidelity-audit.md` KB-1/KB-2):
        `demo-up/SKILL.md:109,153` · `corpus/ops/demo/README.md:34` · **`corpus/ops/rosetta_demo.md:49`** (new) ·
        **`.claude/skills/stack-seed/SKILL.md:50`** (new) · the stale `seed_label` at
        **`up-injected.sh:1317`** (**not** `:1081` — anchor was stale) · `stories.seed.yaml:1`.
        VERIFIED in code — `stories.seed.yaml` ships **3** `org:` entries (`:37` Cervato Systems /
        ai-transformation / 220 · `:75` Solvantis / onboarding-ramp / 120 · `:136` Northwind Aviation /
        ai-readiness / 200). This lie is why the user believed the seeding ask was unmet.
  - [x] **Correct the FALSE safety claim** at **`tailscale-serve.md:452-453`** (**not** `:405-407` — anchor
        was stale; the same stale anchor is cited in `roadmap.md:422`). **VERIFIED FALSE:** ports are
        emitted as bare `"<hostport>:<target>"` pairs at **three** sites in `gen_injected_override.py`
        (`:210` directus · `:276-277` frontends · `:308` backends) with **no `127.0.0.1` prefix**, so Docker
        publishes **every demo container on ALL interfaces on EVERY demo-up, today** — flag or no flag — and
        on Linux its iptables **bypass the host firewall**. `BIND_HOST` (`up-injected.sh:76`) gates only the
        two host-native servers (cockpit, ant-academy). **The doc already CONTRADICTS ITSELF:**
        `tailscale-serve.md:239` states the truth (*"`docker-proxy` binds the demo's offset ports on
        `0.0.0.0`"*) while `:452-453` denies it. A doc that denies a live exposure is worse than none.
  - [x] Regression fences: a doc-vs-code fence on the org count; a fence asserting the published-port shape.
        Home: `stack-core/tests/` (precedent: `test_corpus_index_guard.py` — the existing doc-vs-code fence)
        and `stack-injection/tests/`. Both **RED-proven pre-fix**.

- [x] **S1 — `corpus/ops/safety.md` Part 3: the exposure side.** *(Delivers 1+2 — BLIND AREA, BLOCKS S3)*
  - [x] The gap, re-proven by the audit: `tailscale|remote|expose|network|localhost` → **2 hits, not 0**
        (`safety.md:146`, `:215`) — but **both are incidental** (`in-network` Directus addressing), and the
        doc's section list runs **Part 1 (read side) → Part 2 (write side)** with **no exposure section**.
        **The substance holds: remote reach is a THIRD AXIS with no contract at all.** (Grep the *sections*,
        not the *words* — a keyword hit is not a contract. D17.)
  - [x] State plainly what default-on makes ambient: a demo is an **unauthenticated, authz-weakened build**
        (Clerkenstein disarms token verification; the authz-skip patch is default-on; **the cockpit is a
        one-click, password-free "become any hero" launcher** — a bare GET to `/v1/client/handshake`).
  - [x] Record the steelman FOR the flip: synthetic + public-snapshot-only data (Parts 1-2 hold unchanged), a
        tailnet is an authenticated WireGuard device mesh (per-device keys, ACL-gated, no public listener),
        and — per S0 — **the exposure delta is smaller than the docs claimed, because the LAN exposure already
        exists today.**
  - [x] Explicit written **SUPERSESSION of v2.2's D-DESIGN-1** ("public reach is never default-on",
        **`demo-up/SKILL.md:79`** — **not** `:78`) — **demo path only.** Never a silent contradiction.
        ⚠️ **ID COLLISION (audit KB-4):** **v2.3 has its OWN `D-DESIGN-1`** (*"the <5 s gate is on ACCESS,
        not full render"* — `roadmap.md:127`, `state.md:105`). A bare `D-DESIGN-1` in this release resolves
        to the **wrong** decision. Every supersession sentence MUST read **"v2.2's D-DESIGN-1"**, never bare.

- [x] **S2 — The `/demo-up` defaults table.** *(overview (b) — BLIND AREA)*
  - [x] No enumerated defaults contract exists in the corpus (audit-confirmed: no `knob | default` table
        anywhere under `corpus/`); the only complete knob list is a skill `argument-hint`. Document the
        `DEMO_*` knobs: knob | default | consumer | file:line. **Audit-measured surface: 35 raw `DEMO_*`
        tokens across rext, of which ~25 are real user-facing knobs** — the rest are internals
        (`DEMO_WS`/`DEMO_N`/`DEMO_STACK`/`DEMO_OFFSET`/`DEMO_PORT_OFFSET`), a computed name
        (`DEMO_1_DIRECTUS_DSN`), and a grep artifact (`DEMO_NO_`). Enumerate from the parser; classify, don't
        just dump.
  - [x] **THERE ARE TWO ENTRY POINTS, NOT ONE — and the docs conflate them (audit KB-3, a LIVE false
        promise).** `up-injected.sh` (the one the skill actually invokes, `SKILL.md:52`) accepts **ONLY**
        `<N>` and `--public-host`, and **hard-errors `unknown argument` + `exit 1` on anything else**
        (`:26-27`). `--profile` / `--services` are flags of the **`rosetta-demo` wrapper** (`:110-113`).
        The `demo-up` `argument-hint` lists all four **as if one parser took them** — so
        `up-injected.sh --profile X` **exits 1 today**. The table must record *which entry point reads which
        knob*.
  - [x] Fence: the table is checked against the parser so it cannot drift (the CLI-flag ↔ docs both-directions
        rule — a doc-promised flag with no parser entry is a false promise; a parser flag with no doc surface
        is undiscoverable). **RED-proven pre-fix by KB-3 above.**

- [x] **S3 — The remote flip: `--public-host auto`, DEFAULT-ON for demo.** *(overview (c) — D-DESIGN-3)*
  - [x] Capability ladder — **capability-gated, never presence-probed** (`demo-stack/tailscale_autohost.py`):
        `command -v tailscale` → `BackendState == Running` → a **dotted** `.Self.DNSName` (a dotless host is
        hard-refused — `@clerk/backend`'s `assertValidPublishableKey` rejects it) → `MagicDNSEnabled` →
        `tailscale serve status` shows no operator/sudo denial → **`tailscale cert` actually mints**.
  - [x] **HARD INVARIANT — the fallback is not optional.** Any failed rung ⇒ **empty `STACK_PUBLIC_HOST`,
        byte-identical to today's localhost path**, plus ONE loud line naming the exact fix command.
        **PROVEN LIVE, both directions** (below).
  - [x] Opt-out: `--no-public-host` (+ `DEMO_NO_PUBLIC_HOST=1`). Passing it **with** `--public-host` is a hard
        refusal, not a precedence rule — "public AND not public" has no correct answer.
  - [x] Fences: each rung RED-proven **by mutation** — 4 naive `discover()` mutants (presence-probe rung 6,
        no-dotted-check, soft rung 4, trust-rc==0) + 6 bash-wiring mutants (dropped `|| true`, probe-despite-
        opt-out, second-guess an explicit host, capture-stderr-into-the-host, …). Every one goes **RED**.

- [x] **S4 — Front the cockpit on `tailscale serve`.** *(overview (e))*
  - [x] `('cockpit', 7700)` added to `gen_tailscale_serve.py` — on its **OWN axis**, *not* `UI_BROWSER_FACING`:
        the cockpit is gated on `DEMO_STORIES`, **not** on `--no-ui`, so filing it under the UI tier would leave
        a **live** cockpit unfronted on a `--no-ui` stories demo (and front a dead port when stories are off).
  - [x] **Ordering is load-bearing.** `tailscale serve` binds the tailnet IP `:<port>` as a **real listener**, so
        fronting `:7700` *before* the cockpit binds kills it with `EADDRINUSE` — we would have "fixed" the
        cockpit's exposure by killing the cockpit (the M215 F12 / F-M220-4 contention). The **first** apply
        passes `--no-cockpit`; a **second**, idempotent apply fronts it after `/healthz` answers. The **reset**
        plan always includes `:7700` (else a re-up finds the port held and the cockpit cannot start at all).
  - [x] Documented **honestly**: this is **transport, not authentication** (`safety.md` §3.5.2). The cockpit is
        still a one-click, password-free "become any hero" launcher; it is now behind the tailnet's TLS +
        device mesh rather than in cleartext. Fronting it does not password-protect it.

- [x] **S5 — The two demo-BREAKING click paths.** *(overview (i)+(j) — escalated from M219 D4)*
  - [x] **(i) The academy POISONS the demo session.** FIXED. The academy is **Clerkenstein-wired** from the
        stack's own `.env.demo-N` (minted pk + disarmed fake BAPI + networkless RS256 key) — it never reads
        `platform/.env` again, so keyless mode cannot engage and **no cookie is ever deleted**. The
        `e2e_persona` bypass is REMOVED (kept alongside real keys it short-circuits `proxy.js` BEFORE the real
        session resolves and renders a generic *"E2E Member"* to a presenter logged in as **Maya**).
        The fake BAPI is published on **`127.0.0.1:5401+offset`** — the demo's **first loopback-bound** port —
        because the host-native academy cannot use the in-network `api.clerk.com` alias, and without it its
        only reachable `CLERK_API_URL` is **real Clerk**.
  - [x] **DoD PROVEN — the session SURVIVES.** Controlled A/B on `billion`, from a tailnet **peer**, in a real
        browser, one variable: **ARM A** login → `/profile` ⇒ signed in as *"Maya Chen"*. **ARM B** login →
        `/profile` → **ACADEMY** → `/profile` ⇒ **STILL signed in as "Maya Chen"**. `__session` present
        throughout; `__client_uat` a **live timestamp, never 0**. Direct `curl` at `:13077` now returns
        **ZERO `Set-Cookie` headers** — the deletion mechanism is gone at source. **Values-blind:** the
        academy's `CLERK_SECRET_KEY` sha ≠ `platform/.env`'s, and it is `sk_test_` (Clerkenstein) — the REAL
        production secret is no longer inside a demo process. Exactly **1** publishable-key line, so dotenv
        *last-one-wins* cannot bite again.
  - [x] **(j) studio-desk EJECTS the presenter.** GREEN: `dan-manager` clicks **Anthropos Studio** and
        **STAYS on `:19000`**. Root cause was *not* a broken session — the roster already seeds the manager's
        `admin` membership into the fake BAPI, which `checkEnterpriseAndAdmin` reads. **M219's "302 → /login"
        was an UNAUTHENTICATED `curl`** — the expected answer to a cookieless request, mis-read as the defect.
        An employee is still (correctly) redirected off Studio: that is the real platform's behaviour.
  - [x] Fence: a presenter clicking **Anthropos Studio** stays **in Studio**.
  - [x] M219's launcher fence (`SERVES BUT DOES NOT RENDER`) is **GREEN**. It also had to drop
        `__clerk_handshake` from its keyless pattern: once wired, **that URL is the SUCCESS path**, so the
        fence would have gone RED against the very fix it demanded.
  - [ ] ⚠️ **M219's 400-char CONTENT floor stays RED — honestly, and NOT weakened.** A **separate** defect,
        not the session bug: the academy renders its portal shell + the 3 audience cards and then says
        **"0 PATHS / 0 COURSES / No adventures here… yet"** (348 chars). Its clone HAS content
        (`[build-catalog]` = 2705 entries / 419 public chapters) but `[build-local-catalog]` emits **0**, and
        the home reads the **local** catalog (`local-catalog.generated.js` = 368 bytes). → **`FIX-M221-academy-empty-catalog`** (Fate 3).

- [x] **S6 — Egress: stop the demo phoning home.** *(overview (f)+(g)+(h))*
  - [x] **(f) Clerk telemetry OFF** — both halves: `CLERK_TELEMETRY_DISABLED` (server, runtime env) +
        `NEXT_PUBLIC_CLERK_TELEMETRY_DISABLED` (browser, **build-inlined** → baked into the image's
        `.env.local`). Either alone leaves one collector phoning home. Wired for next-web, studio-desk and the
        academy. **Residual (measured, not assumed):** studio-desk's **Vite** browser bundle reads neither name
        and passes no `telemetry` prop — that half is not reachable from env at all. It did **not** fire in the
        live capture, so no demo-patch was spent on it.
  - [x] **(g) Ad-tech egress (F-5)** — MEASURED FIRST, and the plan **undercounted**: next-web's root layout
        hardcodes **FOUR** scripts with no env seam — `plausible.io`, `analytics.bellasio.com`,
        `uptime.betterstack.com`, **plus** `<GoogleTagManager gtmId='GTM-PXRTBZK'/>` (→ GA + DoubleClick +
        Google Ads + LinkedIn). **Seven** third parties, every page load. New sha-pinned demo-patch
        **`next-web-no-thirdparty`** gates all four behind one build-time env var; behaviour-identical when
        unset. RED baseline from the pre-fix image (2/6/2/2 files); post-fix the **client bundle
        (`.next/static`) carries ZERO** — only dead server chunks + `.js.map`s retain the strings.
  - [x] **(h) Vendor clerk-js + bound the unbounded timeout (C-5)** — clerk-js is now **served from disk** (a
        **box-level** cache shared by every `demo-N`, keyed by the request path's `package@version` —
        self-invalidating); the CDN survives only as a **bounded (15 s)** fallback that populates the cache
        atomically and **never caches a non-200**. Verified live: 4 chunks on disk (incl. the 322 KB main
        bundle), and the browser fetches clerk-js **from the FAPI**, never `cdn.jsdelivr.net`.
        **Alignment re-run after touching the FAPI: Go 27/27 + JS/FAPI 9/9 — 100%/100%, both GATES MET.**
  - [x] **Live egress capture (tailnet peer, authenticated load): ZERO** hits on any of the 11 denied hosts.
        The fence asserts it captured traffic at all — an empty scan is a FINDING, not a pass.
  - [x] **BONUS (found while landing (g), and it would have EATEN it): the image cache had no idea which
        demo-patches were baked into it.** Reuse keyed only on the offset endpoint + minted pk — neither
        related to the patch set. The `demo-1-next-web` image on `billion` matched both, so the first bring-up
        after adding the patch would have **reused it** and served a bundle still phoning home — *while grading
        green*. This is the mechanism behind demopatch-spec's own war story (the 76 s members grid, 4 releases).
        Fixed with a **patch-set fingerprint** baked as an image **label** (no Dockerfile edit). **It fired on
        its first live run** (`<none: predates the fingerprint> != cee1e4ff…` → rebuild).

- [x] **S7 — The dev-side opt-in `--public-host`.** *(overview (d) — folds the reserved M216)*
  - [x] Dev stays **opt-in** per D-DESIGN-3. **LANDED (Fate 1) — the scope-flex lever was NOT pulled.** It was
        the thin wiring job the plan hoped for: S3 had already built and fenced the hard part (the 6-rung
        ladder), and the demo family already had a teardown-reset pattern and a serve generator to reuse.
        `dev-stack up N --public-host <host>|auto` (+ `DEV_PUBLIC_HOST`); **`/dev-up` had no such flag at all**
        before this, so "dev stays opt-in" had been naming a choice the tool did not offer.
  - [x] **REUSED, NOT FORKED.** `--public-host auto` runs `demo-stack/tailscale_autohost.py` cross-section —
        the same pattern, in the other direction, as `up-injected.sh` running `dev-stack/dev-setdress.sh
        --stack-type demo`. `--label`/`--noun` change the WORDS on stderr and nothing else: same rungs, same
        order, same verdict (fenced), and the demo's messages stay **byte-for-byte** what S3 shipped.
  - [x] **THE INVARIANT, fenced with a TRIPWIRE not a mock:** no flag ⇒ **ZERO `tailscale` invocations**. The
        stub is a healthy tailscale on PATH that **fails the test if called at all** — because *"it probed and
        fell back safely"* would be a **passing grade for the behaviour the opt-in default forbids**.
        `DEV_PUBLIC_HOST`, deliberately **not** the demo's exported `STACK_PUBLIC_HOST`: an ambient value would
        otherwise flip a dev stack public **with no flag on the command line**.
  - [x] **Fences RED-proven by MUTATION: 9 mutants, 9 RED, 0 theatre, 0 no-ops.** ⚠️ **The first battery was
        itself theatre** — its `restore()` ran `git checkout` against **uncommitted** work, so mutants 2–8 ran
        against a tree where S7 **did not exist** and "went RED" because the feature was **absent**. The tell
        was in the output: **M2–M7 all reported an identical 15 failures.** A uniform count across unrelated
        mutations is not a result, it is a constant. Take 2 runs against a **committed** baseline and **asserts
        every mutant actually changed the file** (a no-op mutant that "goes RED" is measuring something else).
        **D17, reproduced inside the battery built to enforce D17.**

- [x] **S7 bookkeeping — the CLI-flag ↔ docs rule, applied to the side that had no fence.**
  - [x] **`--inject` has been in `dev-stack up`'s parser since M5 with NO user-facing doc surface** — it exists
        and nobody can find it. That is direction (2) of the both-directions rule, live for releases, on the
        one path S2's fence did not cover. New `stack-core/dev_flag_guard.py` (both directions + a third
        clause: **being *hinted* is not being *documented***). **RED-proven: 2 UNDISCOVERABLE flags** pre-fix.
  - [x] **S2's defaults table still holds** after S3/S4/S7 — `demo_knob_guard` re-run, **PASS**. All four
        corpus guards green (`story_org_count` · `demo_knob` · `dev_flag` · `exposure_claim`).
  - [x] 🔴 **THE REAL FIND: the dev family had NO exposure disclosure at all.** `stack-core/gen_override.py`
        builds its port strings **exactly like the demo's** — bare `"<hostport>:<target>"`, no `127.0.0.1` —
        so **every `dev-N` container is world-published on `0.0.0.0`, on every `dev-stack up`, flag or no
        flag**, and on Linux Docker's iptables bypass `ufw`. `safety.md` §3.1 disclosed this **for demos
        only**. The silence landed exactly where it does the most damage: **dev's opt-in default invites the
        inference *"remote reach is off, so I am not exposed"* — which is FALSE.** The opt-in withholds the
        **trusted HTTPS origin**, not the LAN binding, which was always there. **This is the S0 lie, one family
        over.** `exposure_claim_guard` now RUNS both emitters (**DEMO 14 → `0.0.0.0`, DEV 8 → `0.0.0.0`** —
        measured, not read) and a **separate** `_DEV_DISCLOSURE_RE` fences the dev half, because the generic
        regex was satisfied by the demo paragraph alone — **one family's disclosure standing in for two**.
        Both halves RED-proven (doc-side + code-side).
  - [x] **M216's reservation is CONSUMED** (`decisions.md` **D28**) — not handed back.

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

### S3 + S4 landed 2026-07-14 — ONE live demo cycle on `billion`, both paths proven

**The flip is live and the fallback holds.** Asserted from a **tailnet peer** (the Mac), never on-host.

| path | trigger | result |
|---|---|---|
| **default-on** | bare `up-injected.sh 1`, **no flag** | `AUTO-DISCOVERED billion.taildc510.ts.net (all 6 rungs)`. **Cockpit `:17700` = HTTPS 200, `ssl_verify=0`** (trusted LE cert) — it was **plain HTTP** before S4. All 5 other browser ports `ssl_verify=0`. Hero login end-to-end **over the cert**: `maya-thriving` → `/profile` **200**, `dan-manager` → `/enterprise/workforce` **200** (control: no session ⇒ **307** handshake-loop). Cockpit renders all **3 orgs** + heroes. autoverify **OK** (casbin 1150, skills 42790). Clones left **0-dirty**. |
| **fallback** | `tailscale` made **genuinely unavailable** (`chmod -x`; `shutil.which` → `None`) | `STOPPED at rung 1/6 … Fix: install Tailscale`. Cockpit `http://localhost:17700`; **0** serve listeners; baked URLs `http://localhost:*`; cockpit re-bound to **`127.0.0.1`** and **REFUSED** from the tailnet IP; autoverify **OK**. **Byte-identical.** |

**The open question that gated the flip — SETTLED EMPIRICALLY.** rext claimed `tailscale cert` *"RE-ISSUES on
re-run"*. **Measured FALSE:** two back-to-back mints returned the **identical certificate serial** (`05777C48…`)
in **0.01 s** each, with **zero** new ACME orders. tailscaled caches. Had the claim been true, default-on —
whose rung 6 **mints on every bring-up** — would burn a Let's Encrypt **duplicate-certificate** slot per
`demo-up`, and since **`ts.net` is a PSL entry** that bucket is **per-tailnet**, shared by every box on it.
Corrected in `tailscale-serve.md` + `up-injected.sh`.

**My own ordering fence was THEATRE on the first cut.** It asserted `--no-cockpit` was present in the first
serve apply by scanning the script text — and matched the word inside the **comment** I had written directly
above the command. Deleting the flag from the *command* still passed. It now strips comments before scanning.
D17 reproduced **inside the fence built to catch it**; only the mutation run exposed it.

### S5/S6 shipped a RED suite, and it nearly cost me a mis-triage (fixed here, Fate 1)

`cue-to-cue-m220-r4` — **the tag the demo host was pinned at** — had **11 failing demo-stack tests**. S5/S6
changed the code and left its tests asserting the **old** contract: the inverse D17, a *test* outliving the
thing it describes. My own 7 regressions landed on top of them, and "14 failures" is indistinguishable from
"11 pre-existing + 3 mine" until you baseline. All 18 now green (569 + 214 + 129):

- **`test_frontend_build`** — S6/h's clerk-js-cache `mkdir` (D14) landed **inside the block**
  `TestReuseFlagArrayExpansion` extracts, and the harness PATH had no `mkdir` ⇒ `rc=127` on all 8 cases, with
  the assertion message blaming the `+`-guard (*"regressed?"*) and pointing the next reader at **entirely the
  wrong bug**.
- **`test_frontend_build`** — the docker stub answered **every** `--format` inspect with `$STALE_ENV`, so
  S6/g's new `demo.patchset` label read back as the env blob, never matched, and **all five REUSE tests
  silently exercised the REBUILD path** — asserting the opposite of their own names while the suite looked
  fine. The stub now answers the label query separately, with the fingerprint the **shipped** function computes.
- **`test_ant_academy`** — three assertions still demanded the `e2e_persona` bypass S5/D10 **deliberately
  deleted**, and the pk/secret S5 **stopped** copying out of the real `platform/.env`. **Inverted, not dropped**
  — a removed feature deserves a fence that keeps it removed.

### S0–S2 landed 2026-07-14

**Three fences, all RED-proven against the pre-fix tree, all GREEN after.** Every one of them found more than
the plan predicted, and — the useful part — **each caught bugs in itself while being RED-proven.**

| Fence | Home | Pre-fix | Post-fix |
|---|---|---|---|
| org count (doc↔preset) | `stack-core/story_org_count_guard.py` | **RED — 11 violations** | GREEN |
| exposure claim (doc↔emitters) | `stack-injection/exposure_claim_guard.py` | **RED — 3 false claims + 2 missing disclosures** | GREEN |
| defaults table (doc↔parsers) | `stack-core/demo_knob_guard.py` | **RED — doc absent, then 2 conflations** | GREEN |

**The count that matters: the plan said 4 org-count sites. The KB audit found 7. The fence found 11.** That gap
*is* the argument for a fence over a prose list — and it is the same D17 shape the release is about: a
hand-maintained list, never checked against the thing it describes, read as if it were complete.

**The fences' own bugs (each now a regression test — a fence with a hole reports a confident, quietly incomplete
failure list, which is D17 one level up):**
- `.agentspace` in an exclusion list, matched as an **absolute-path substring** — rext lives under
  `.agentspace/`, so the guard **silently skipped both of its own repo's sites** while printing a clean-looking
  5-hit list.
- `heroes?` parses as `"heroe"`+optional-`s` and does **not** match bare `"hero"` — the `seed_label` site slipped
  through.
- **Markdown prose WRAPS.** The shipped exposure lie was split across a line break; a per-line scanner reported
  1 confident hit **while a second live lie sat one line below, inside the block it had just audited.**
- **Truth is per-preset.** `stories-maya` really ships 1 org, so *"one org + Maya"* is TRUE. A single global
  truth flagged three true statements as lies — including a **measurement** (*"1 org on the M215 run"*, a
  correct observation of a cold-cache run). Rewriting that to "3" would have had the guard **manufacturing** the
  defect class it exists to catch.
- **A naive both-directions check goes GREEN on the conflation bug.** `--profile` *is* parsed — by
  `rosetta-demo`. So *"does SOME parser accept it?"* passes while `up-injected.sh N --profile cms` still
  **exits 1**. The rule had to become: accepted by the **primary** entry point, or the hint **names** the other.

**Scope honesty (D8):** `demo-up/SKILL.md` was **not** changed to say "default-on". The code still requires the
flag — that flip is **S3**. A doc claiming a behavior the code lacks is the exact defect S0 just fixed. What
landed is the *decision* (`safety.md` §3.5) plus a forward-pointer; not a pre-announcement of code.

**Routed out:** `FIX-M221-devstack-test-spin` (Fate 3 → M221) — `dev-stack/tests/test_dev_stack.py` busy-spins
forever (145 % CPU, `rc=124`), so the rext suite cannot be run whole. Pre-existing; reproduces on clean `HEAD`
with M220 stashed. It will block the release close.
