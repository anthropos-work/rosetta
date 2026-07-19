# Demo Environments — the family index

**Purpose.** Stand up a **disposable, isolated, Clerk-free, realistically-populated** copy of the Anthropos
platform — for sales demos, screenshots, QA, or a clean-room — *alongside* the dev stack and **without
touching any read-only platform repo**. This family is the entry point; the mechanism guides + recipes it
indexes are the depth.

**When to use.** You want a demo world a stakeholder can log into and click around (a populated org with months
of activity), reproducibly, on offset ports, killable cleanly — and you must never pollute production or the
dev stack. If you just need the *dev* environment, see `../setup_guide.md` / `../run_guide.md` (driven by
`/dev-up`). If you need *staging*, see the `../staging-*` family.

> **Dev is a peer (v1.3).** Every set-dressing + seeding recipe below works on a **`dev-N`** stack exactly as it
> does on a `demo-N` — the `/stack-snapshot` / `/stack-seed` ops take `dev-N|demo-N` interchangeably, and a
> `/dev-up N` bring-up already set-dresses + light-seeds itself by default. Where a recipe says `demo-N`, read
> `dev-N|demo-N`. The one exception is the **`N=0` main dev stack**, which the auto-set-dress + `--reset`
> guards protect (see [`../safety.md`](../safety.md) §2.5).

## The end-to-end flow (~minutes)

```
/demo-up N        →  AUTO ensure-clones (bootstrap stack-demo's OWN clone set: clone platform from GitHub  [v1.8 M26 → rosetta_demo.md]
                     + make init the peer repos; seed the shared .env copy-if-present) — self-contained
                     bring up demo-N (Clerkenstein-wired, offset ports, isolated data, built from stack-demo) [corpus/ops/rosetta_demo.md]
                     AUTO secret-provision (values-blind, per-repo .env from .agentspace/secrets) [v1.6 M30 → secrets-spec.md]
                     AUTO set-dress (cache-first snapshot replay → Stories & Heroes seed + cockpit, default-on, non-fatal) [v1.9 M38]
                       (DEMO_NO_STORIES=1 → structural small-200 seed + single-identity fake-fapi, no cockpit — the fallback)
  …present it…     →  open the presenter cockpit → pick a hero → [Log in as] → her per-role screen    [cockpit-spec.md]
  …or use it…      →  browser-login as user_clerkenstein → land in a populated org (200)    [recipe-browser-login.md]
/demo-down N      →  tear it all down (AND reap the native cockpit process), dev stack untouched [corpus/ops/rosetta_demo.md]
```

> **The storytelling demo + the presenter cockpit (v1.9 "storytelling" M34–M38) — now the DEFAULT.** A bare
> `/demo-up N` set-dress seeds the locked **3-orgs × 3-heroes** Stories & Heroes world (Cervato Systems ·
> Solvantis · Northwind Aviation — each org a
> thriving/struggling/manager trio), runs a **multi-identity** fake-fapi (a `FAKE_FAPI_ROSTER` of the seeded
> heroes' exact ids), and serves a **presenter cockpit** on an offset port (`7700 + N·10000`). The cockpit is
> the demo's remote control: a standalone panel (never an in-app overlay — the zero-platform-repo-edit line
> holds) listing each story → its hero trio with one **[Log in as]** action per hero, so a demo-giver
> picks a hero, lands logged-in as her on the right per-role screen, and presents that flow live (the slick
> light UX — icons, the unified CTA, a manifest download, a login-progress overlay — is v1.10 M43, specced in
> [`cockpit-spec.md`](cockpit-spec.md)). `DEMO_NO_STORIES=1`
> (or the explicit `DEMO_STORIES=0`) restores the legacy structural `small-200` + single-identity fake-fapi +
> no-cockpit demo (the `DEMO_NO_*` family fallback). The cockpit serve is **non-fatal** (a stories demo fails
> loud on a broken roster — the login-as contract — but a cockpit failure leaves a fully-working seeded
> multi-identity demo you can still drive by hand); pass `DEMO_NO_COCKPIT=1` to bring the stories demo up
> without the panel (an API-only run). The cockpit UX surface + the deep-link contract are in
> [`cockpit-spec.md`](cockpit-spec.md); the single-source manifest, the roster-export producer, and the O9
> deep-link catalog (the seat/seed producer seam) are in
> [`stories-spec.md` § The presenter cockpit](stories-spec.md#the-presenter-cockpit-m38).
>
> ```
> /demo-up 3                   →  default: seed the 3-org hero trio + multi-identity fake-fapi + serve the cockpit
>   …present it…              →  open http://localhost:37700 → pick a hero → [Log in as] → her per-role screen
> DEMO_NO_STORIES=1 /demo-up 3 →  fallback: structural small-200 seed + single-identity fake-fapi, no cockpit
> /demo-down 3                →  tears down the stack AND reaps the native cockpit process
> ```

> **A demo builds from its OWN clone set (v1.8 "understudy" M26).** `/demo-up` first runs `ensure-clones.sh`:
> it bootstrap-clones `stack-demo/platform` from GitHub over SSH + `make init`s every `repos.yml` repo as a
> sibling into `stack-demo/`, so **all** images build from `stack-demo` (a box with only `stack-demo/` — no
> `stack-dev/` — can bring a demo up end-to-end). The sole sanctioned `stack-dev` read is seeding the shared
> `platform/.env` copy-if-present (non-fatal if absent — M30 provisions the real one); the build SOURCE never
> falls back to `stack-dev`. Dev-image reuse is OFF by default (`DEMO_REUSE_DEV_IMAGES=1` opts back in).

**`/demo-up` now auto-set-dresses by default (v1.3b M20) — the dev↔demo convergence.** Just like `/dev-up` since
v1.3, a `/demo-up` bring-up chains the **same** set-dress pass at its tail: a cache-first **snapshot replay**
(the real **taxonomy** catalog; and — **for a demo, local content is default-on** since v1.5 M22/M23 — a
per-stack **Directus** booted + cut over so the **content** surface serves locally too, the stack
content-self-contained; a demo opted out with `DEMO_NO_LOCAL_CONTENT=1` falls back to reading content live from
prod — see the known-state note below) → the **Stories & Heroes seed + cockpit** (the v1.9 M38 default — a
narrative multi-org world you can log into and present; `DEMO_NO_STORIES=1` falls back to the structural
`small-200` light seed). So a bare `/demo-up N` already lands you in a real-catalog, log-in-able world
— no separate skill calls required. The pass is
**default-on + non-fatal** (a cold cache warns and still seeds; `DEMO_NO_SETDRESS=1` skips it for a bare
structural bring-up). You can still drive the steps **manually** for finer control — `/stack-snapshot N` (replay)
+ `/stack-seed N` (a different preset / a custom `stack.seed.yaml`) — they accept `demo-N` or `dev-N` interchangeably.

**The snapshot step is what makes the world *set-dressed* (v1.2).** A replay stamps the real **public** reference
library — the ~60K-skill taxonomy + the global simulation / skill-path content templates — into the stack BEFORE
the seed, so the catalog view shows real skills and the seeded sessions link to real templates (not free
placeholder ids). It's a **stack-global reference replay**, independent of which org you then seed; it's
**optional** (a structural-only world still logs in — the seeder degrades gracefully), and almost always a
**cache-hit** (zero prod read — the snapshot is captured once per release, then replayed by every stack).
See [`recipe-snapshot-world.md`](recipe-snapshot-world.md) for the full capture→replay→set-dressed walk-through.

> **Fresh box, empty cache?** The replay is a cache-hit only once the cache has been filled by a one-time
> `capture` — and a fresh machine with no safe `--dsn` can't replay the *real* catalog yet (the auto-set-dress
> warns + degrades to an empty catalog, then seeds). The sanctioned way to fill the cache once per release —
> and why the wired `postgres` MCP is **not** a capture source — is [`../snapshot-cold-start.md`](../snapshot-cold-start.md).

## Index

**Mechanism guides (the "how it works"):**
- [`../safety.md`](../safety.md) — the **safety contract**: the consolidated read-side (tenant-data firewall,
  public predicates, read-only capture) + write-side (3-layer isolation guard, never-write-prod, n=0 guards,
  audit-proven zero pollution) statement. **Read this first** if you care *why* it's safe. (v1.3 M15)
- [`../rosetta_demo.md`](../rosetta_demo.md) — the **lifecycle**: bring-up, the port-offset scheme, the
  Clerkenstein injection, the per-stack project/data isolation, the resource budget, teardown. (M3)
- [`../seeding-spec.md`](../seeding-spec.md) — the **seeding** reference: the `stack.seed.yaml` blueprint, the
  dependency-DAG, the **production-isolation boundary**, the casbin subtleties, the data-DNA. (M7a/b)
- [`stories-spec.md`](stories-spec.md) — the **verified-skill chain** + **Stories & Heroes** reference: how a
  seeded *verified skill* (a hero's passed-simulation profile + Skill Spotlight chart + the claimed-vs-verified
  gap) is materialized as the **7-table fan-out** the `PersonaSeeder` writes, the constraint landmines, the G14
  session fix, the seed-side closure gene — plus the multi-org **thriving/struggling/manager trio** model (M35),
  the **Workforce-dashboard surfaces** (M36), and the **presenter cockpit** (M38 — a standalone served panel that
  lists each story → its hero trio with a **[Log in as]** action, riding the Clerkenstein multi-identity
  seat-switch). The believability spine of a demo world + its demo-driving surface. (v1.9 M34–M38)
- [`cockpit-spec.md`](cockpit-spec.md) — the **presenter-cockpit UX spec** (v1.10 "method acting" M43): the slick
  **light** login launcher a demo-giver drives — the card-per-hero layout + FontAwesome icons, the **one
  unified [Log in as] CTA** per hero (logs in *and* lands on her per-role `jump_to`), the seed-manifest
  download, and the staged login-progress overlay — plus the deep-link contract + the standalone-served-panel
  (zero-platform-edit) model + the future-feature surface. Graduates the M37/M38 cockpit mechanics scattered
  across `stories-spec.md` + `clerkenstein.md` into one place. (The seed/seat **producer** seam stays in
  `stories-spec.md`.)
- [`profile-completeness-spec.md`](profile-completeness-spec.md) — the **"complete profile" rubric**: the
  DATA-DENSITY bar for a fully-populated profile across the WHOLE roster — identity + content + semantic layers,
  per-vantage **member vs manager**, each component mapped to its seeding surface + a Playwright acceptance
  assertion. Covers the M44 fills: trajectory-aware self-rating (§A), the `CertificatesSeeder` + `ProjectsSeeder`
  surfaces (§B), the manager personal-data unskip (§C), and the bulk-member shallow career (§D). (v1.10 M44)
- [`../snapshot-spec.md`](../snapshot-spec.md) — the **snapshot** reference: how the real **public** taxonomy +
  Directus content library is captured once from prod safely (the read-side **tenant-data firewall**), cached
  outside git, and replayed per-stack — measured-faithful by the snapshot-fidelity data-DNA. (M9a/M9b/M10)
- [`../snapshot-cold-start.md`](../snapshot-cold-start.md) — the **cold-start** runbook: filling the snapshot
  cache once per release on a fresh box (the sanctioned DSN-export / dump-restore path), why the wired `postgres`
  MCP is **not** a capture source, and how it slots into the auto-set-dress bring-up. (v1.3b M20)
- [`../idempotency.md`](../idempotency.md) — the **re-run safety** contract: what happens when you run
  migrate / snapshot-replay / seed a *second* time — each is now safe-and-idempotent or fails loudly, never
  silently doubles or aborts mid-surface (the replay TRUNCATE-then-reload, the idempotent seed COPY + casbin
  guard, the `--reset` fix, the `set -e` first-run-race hardening). (v1.3b M17)
- [`../verification.md`](../verification.md) — the **verification** net: every bring-up ends with a scoped,
  NON-FATAL verify on the stack's OWN offset ports (the cheap-win `/api/health` + `casbin_rules > 0` silent-403
  catcher, then the full probe set), so "UP" means *verified-working* — offset/scope-aware, never blocks a good
  stack. (v1.3b M18)
- [`frontend-tier.md`](frontend-tier.md) — the **UI tier**: how `/demo-up` brings up next-web-app +
  studio-desk (per-demo cached Docker image from the **unmodified** Dockerfile, offset ports, minted-pk +
  offset-URL baked) + ant-academy natively (Clerk-free), the 12 GB Docker-VM prereq + non-fatal pre-flight,
  the honest "one ~3-min cached build per new demo-N" residual, and the `--no-ui` escape. (v1.3b M19)
- [`demo-up-defaults.md`](demo-up-defaults.md) — **the defaults contract** (v2.3 "cue to cue" M220): every
  knob and flag that controls a bring-up — **all 25 env knobs + 9 CLI flags**, with real defaults and the exact
  `file:line` that reads each. **Derived from the parsers, and fenced against them in both directions** (a
  doc-promised flag with no parser entry is a *false promise*; a parser flag with no doc row is
  *undiscoverable*). States the fact that had never been written down: **there are TWO entry points** —
  `up-injected.sh` takes only `<N>` + `--public-host` and **hard-errors on anything else**, while
  `--profile`/`--services` belong to the separate `rosetta-demo` wrapper. And the shape of the whole surface:
  **every feature knob is an opt-OUT**, so a bare `/demo-up N` already gives you the 3-org world, the full UI
  tier, the cockpit, and set-dress — *"pull all the data + seed the 3 orgs" was always the default; the usual
  culprit is a cold snapshot cache, not a knob.*
- [`tailscale-serve.md`](tailscale-serve.md) — the **remote-access recipe** (v2.2 "panorama"; remote reach
  flipped **default-on for the demo path** at v2.3 M220 — D-DESIGN-3): remote reach is **default-on for `/demo-up`,
  opt-out via `--no-public-host`** (`/dev-up` stays **opt-in** via `--public-host <magicdns>`), making a demo
  reachable from another machine on your **Tailscale** tailnet,
  the **HTTPS-everywhere** per-offset-port topology (`tailscale serve` + the tailscale-cert FAPI), what the knob
  flips (CORS `https://$HOST` origins, the studio-desk/academy redirects, every baked URL's scheme), the
  **patch tail** (ant-academy `allowedDevOrigins` + the studio-desk `VITE_CLERK_SIGN_IN_URL` overlay, via the
  sha-pinned mechanism), the "teammate on the tailnet browses it" walkthrough, and the safety framing
  (Tailscale = the access control; default-on for demo / opt-in for dev; zero platform-repo edits). (v2.2 M212–M214; default-on flip v2.3 M220)
- [`demopatch-spec.md`](demopatch-spec.md) — **the demo-patch mechanism: the sanctioned zero-platform-edit escape
  hatch.** When a demo needs a fix that has **no env/config/compose seam** (the value is baked into platform
  source), `demopatch` patches the demo's **own ephemeral clone** just before the image build and reverts it after —
  so the *image* carries the fix, the clone is left git-clean, and the canonical `anthropos-work` repos are **never
  touched**. Documents the **7 guards** (G1 path-assert · G2 drift-refuse + exactly-once anchor · G3 never-commit ·
  G4 idempotent · G5 self-revert · G6 demo-only · **G7 apply post-condition**), the 10-key manifest schema, the
  **three apply vehicles** (the `app` patches target the build-scratch clone *outside* the workspace, so
  `demopatch`'s own G1/G6 correctly refuse them), the **chain rule**, and — the M217 lesson — the **self-healing
  freshness gate**: *the anchor is the contract, the whole-file sha is only a baseline*. **Read this before adding
  any patch.** (v2.3 M217)
- [`latency-budget.md`](latency-budget.md) — **the demo's performance budget: what "fast" means, and how it is
  measured.** Before v2.3 there was **no** perf budget, baseline, gate, or even a *definition of "access"*
  anywhere in the corpus — while a presenter's click→login actually took **60–120 s**, and the corpus asserted
  in four places that it took "~2–5 s, which we can't shorten." Defines **ACCESS** (the authenticated shell is
  rendered and interactive with the hero's identity present), the **< 5 s p95 gate**, the **per-leg attribution
  model** (click → handshake/303 → SSR → clerk-js → client-gate → data-query), the measured baseline
  (**39.45 s** employee / **38.30 s** manager) and the shipped number (**cold p95 2413 ms / 1767 ms**), the
  harness contract (`stack-verify/e2e/run-latency.sh` — never gate on `networkidle`; always gate on a **fresh
  green** `autoverify.json`), and the **arithmetic signatures** that name a bug class before you read a line of
  code (a *blackholing* address ≈ `3 × 10.5 s + 6 s`; a *fast-failing* fetch ≈ `3 × 33 ms + 6 s`). **State the
  environment with every number** — the same defect cost ~6 s on a laptop and ~112 s on the tailnet VM.
  (v2.3 M218)
- [`coverage-protocol.md`](coverage-protocol.md) — the **coverage** iteration protocol: the **Playwright**
  demo-coverage sweep + triage + fix loop driving the **semantic believability gate** (real seeded content +
  substantial per-section cardinality + persona self-consistency [role↔skills, menu==profile real-photo avatar,
  org name+logo] + 0 prod-eject escapes — NOT the old `textLen>40` density check). The manifest-driven section
  model, the fix-surface routing table (empty section → `stack-seeding`; content error → `stack-snapshot`
  serve-grant; out-of-demo link → injection link-rewriting; runtime-computed surface → crawl-scope), and the
  disclosed-presenter-note allow-rule for legitimate external citations. The harness lives in rext
  `stack-verify/e2e/`. (v1.10 M42e)
- [`playthroughs.md`](playthroughs.md) — the **functional-flow e2e runbook** (the Playthroughs pillar, v2.0
  "opening night" M202): a **Playthrough is an automated actor that IS the user** — it logs in as a seeded hero,
  plays a real journey, and proves the platform delivered the outcome. Proves **function** (the hero can *do* the
  thing) where the coverage sweep proves **presence** (every page *shows* real content). The manifest model
  (Products → Stories → Use Cases → Playthroughs) + the light validator (both-way id integrity +
  precondition-coverage + the datadna gate), the per-surface page-object/locator layer (re-pin O(surfaces), not
  O(tests)), the dedicated decoupled seed (`pt-world`, test data ≠ demo data) + the reset-to-seed lifecycle
  (the real `--reset`, additive-re-seed FORBIDDEN), the serial-default runner, and the 4-state reporting map
  (`passing`/`failing`/`unimplemented`/`unimplementable-without-platform-edit`). Also **the iteration protocol
  the coverage milestones followed** (M203 employee-vantage + M204 manager-vantage, both landed — the corpus now
  stands at **10 live Playthroughs, 1 TODO**). Section `rext playthroughs/`. (v2.0 M202–M204)
- [`content-stories-routes.md`](content-stories-routes.md) — the **content-stories feasibility spike + result-route
  map** (v2.5 "the playbill" M231, HARD go/no-go — the barrier before the Thread-B build chain). For each content
  product × {player, manager} it enumerates the exact result route and **classifies it by prove-by-render**
  (renders-from-seed | runtime-computed-blank | needs-demo-patch | no-surface). The central finding: the simulation
  result page reads a **PERSISTED row** (resolver does plain SELECTs, no live recompute) → **seedable**. Simulation +
  skill-path GO; **Interview** GO behind a PostHog-flag demo-patch; **AI-labs OUT** (no seedable result surface);
  **Academy IN** (backend-authoritative progress). Plus the **prod-session sourcing + anonymization contract**
  (pin by `sessions.id`, public-anchored, the free-text scrub surface) the M232 seeder feeds, the **public-sim-by-
  modality catalog** (77 voice / 65 code / 30 document), and the generalized **manager-view MIRROR trap**. (M231)
- [`session-clone-spec.md`](session-clone-spec.md) — the **session-clone / sourcing seeder** (v2.5 "the playbill"
  M232 — the write side of the content-stories feature). The `ContentStorySeeder` **COPIES real production
  job-simulation sessions** into a demo org: the REAL result-fan-out CONTENT (LLM feedback, transcript,
  submission, interview report — the interesting free-text) is **copied** from the pinned session (authoring-time,
  `cmd/content-capture` reads prod read-only) and **SCRUBBED best-effort** of detectable PII (real actor names +
  source org → placeholders the seeder fills with the demo persona/org; emails/phones/urls redacted). **NOT
  provably clean** — residual re-identification risk is real and **ACCEPTED by the data-controller (2026-07-19)**;
  the control is the **VPN/tailnet scope**. **Re-tenanted**, **non-manager-played** (owner = a seeded player
  member), **source-pinned** (deterministic reseed; disclosed in the `content_sessions` manifest block). The full
  result fan-out (session + the `local_jobsimulation_sessions` MIRROR + attempt/skill/criterion/check results +
  transcript actors/interactions + the net-new **CODE**/**DOCUMENT** substrate + the **INTERVIEW** report), all
  G14-valid, the REAL skill node-ids copied; plus the two sha-pinned interview-flag-gate **demopatches** (the M219
  aireadiness twin — no PostHog on a demo ⇒ no rollout gate). The bounded read-side exception `safety.md` §3.8
  records. (M232)
- [`content-stories-spec.md`](content-stories-spec.md) — the **content_products manifest + honesty gate** (v2.5
  "the playbill" M233 — the manifest half of the content-stories feature). `stackseed --content-export` PROJECTS a
  **`content-manifest.json`** (the content analog of `cockpit-manifest.json`) the 2nd "Content stories" cockpit tab
  reads: per content product, the played sessions each with a **player + manager seat key**, a **result path**,
  `has_manager_view`, a per-product **app-base**, and a per-type **icon**. Single-sourced from the SAME
  content-session fixture the seeder seeds from (the player seat OWNS the seeded session; the path names the seeded
  session id — no drift), **honesty-gated** (a checked-in canonical + a `CanonicalFileMatchesProjection`-style test,
  with teeth), and **fail-closed** (a session that can't form a real link is DROPPED with a reason + the export
  fails loud — never a fabricated CTA). Separate JSON (not a YAML block) because the cockpit reads JSON, not YAML.
  (M233 — M234 = the cockpit tab render + player-seat registration, M235 = prove-it-lands)
- [`ai-generation-spec.md`](ai-generation-spec.md) — the **generation-engine** + **gen-acceptance protocol**
  (v1.10 "method acting" M45): how a cheap LLM (gpt-4o-mini) turns a YAML **batch descriptor** into realistic
  per-member profiles — the `services/ai/` wrapper (EU-first routing + cost tracking), `blueprint.Batch` +
  `EffectiveBatches()` (pure Go-template mother-prompt expansion, NO LLM at parse time), the `cmd/gen-batch`
  CLI (mandatory `--max-cost` ceiling, `--max-concurrent` semaphore, `--call-timeout`, re-roll-on-malformed +
  hero-collision re-roll), and the `GeneratedBatchSeeder` — enforcing the **CODE-owns-structure /
  AI-owns-content** boundary (every generated role/skill name routes through the existing resolvers;
  non-resolving names **drop**, the closure gene stays green, never fabricated). Plus the measure→fix→accept
  iteration protocol (the 5-metric gen-quality gate). The FIRST new third-party dep in the seeding module
  (`ai v1.40.1`, a deliberate in-release decision). (v1.10 M45)
- [`cache-spec.md`](cache-spec.md) — the **prompt-hash cache** (v1.10 M45): the
  `.agentspace/.batchcache/batch-${hash}/member-${i}.json` store keyed by the **MOTHER prompt** + the
  **taxonomy capture version** (invalidate on re-replay), atomic `.tmp`→rename writes, the `.lock` fence — so
  an unchanged batch descriptor **re-seeds byte-identical at $0**. Pairs with `ai-generation-spec.md`. (v1.10 M45)
- [`seed-manifest-spec.md`](seed-manifest-spec.md) — the **consolidated single-auditable seed+generation
  manifest** (v1.10b "fit-up" M52): ONE checked-in `seed-generation-manifest.yaml` inlining the whole demo-data
  intent — the population (all 3 orgs + heroes), the file-resident **mother prompt** (extracted from the Go
  const to `blueprint/prompts/default_batch_prompt.tmpl`), the batch config (the MANDATORY `max_cost_usd`
  ceiling + concurrency + re-roll rules), and the snapshot sources; **cache + generated data EXCLUDED**. A
  PROJECTION of the canonical presets (honesty-gated so it can't drift), emitted by `stackseed
  --manifest-export`, served by the cockpit's **[Download seed manifest]**. (v1.10b M52)
- [`../../architecture/alignment_testing.md`](../../architecture/alignment_testing.md) § "The data dimension" —
  the **data-DNA**: how a seeder's output is conformance-gated against the platform schema, and drift-detected.
  With snapshots, coverage now reads **100%** (both formerly-`waived` surfaces promoted to `snapshot-seeded`).

**Use-case recipes (the "build a demo for X"):**
- [`recipe-enterprise-onboarding.md`](recipe-enterprise-onboarding.md) — a populated enterprise org (admin +
  hundreds of members), end to end — now **set-dressed** (real catalog + content behind the seeded org).
- [`recipe-skill-progression.md`](recipe-skill-progression.md) — months of believable skill-progression
  activity (backdated job-sim + skill-path sessions) linked to the real template library.
- [`recipe-snapshot-world.md`](recipe-snapshot-world.md) — the **set-dressing** recipe: capture →
  replay the public taxonomy + content into a stack so the catalog + templates are real, not placeholder.
- [`recipe-browser-login.md`](recipe-browser-login.md) — the **interactive** demo: the `@clerk/express` /
  orgclient cert-redirect + the browser-login walk-through, log in → land in a seeded org.

**Curated seed presets** (instances of `stack.seed.yaml`, validated to seed):
`rosetta-extensions/stack-seeding/presets/` — `small-200` (quick — **the `DEMO_NO_STORIES=1` fallback seed**,
M20 #M20-D2) · `mid-500` (the default "looks real") · `large-1k` (scale). The `/demo-up` auto-set-dress now
defaults to the **Stories & Heroes** seed (v1.9 M38); `small-200` is the structural fallback the
`DEMO_NO_STORIES=1` opt-out seeds (a fuller world than dev's `dev-min`). Override either with a manual
`/stack-seed N --preset mid-500`
(or skip the auto pass with `DEMO_NO_SETDRESS=1` and seed by hand). The presets are **purely structural** (they describe an org, not the
platform's reference library); for a **set-dressed** world the catalog replay runs first (the auto pass does this;
manually it's `/stack-snapshot replay N`). Without a replay the seeder degrades gracefully (empty catalog, free
content refs).

> **Known state — a `--local-content` stack is content-self-contained (M22 boot + M23 cutover); a prod-read
> stack reads the public catalog live from prod.** The auto set-dress replays the **taxonomy** locally. For the
> public Directus content: since **M22** a stack brought up with **local content** (demo **default**; dev
> `--local-content`) **boots its own per-stack Directus** (a compose service serving the captured catalog — the
> M10 collection-schema gap was closed by M21 + executed by M22), and since **M23** the bring-up **cuts `cms`
> over** to it (`DIRECTUS_BASE_ADDR` → in-network `http://directus:8055`) so content is served locally (asset
> plane stays on prod public links → real images). A **prod-read** stack (`DEMO_NO_LOCAL_CONTENT=1`, or a plain
> dev bring-up) still has **no local Directus**: `cms` reads the public sims/skill-paths **live from prod**
> (`content.anthropos.work`) — a **demo does so ANONYMOUSLY** since fix16/fix17: the injected override strips the
> inherited prod `DIRECTUS_TOKEN` from every demo container (prod Directus serves the public predicate tokenless
> — verified 2026-06-11; live demo-1 audit: 0/16 carriers), so no prod credential rides in a demo. The read is
> public-only + safe, but on the prod-read path it means a stack isn't fully self-contained. **The
> `--local-content` path (the demo default) closes this:** M22 boots + verifies the local Directus and M23 cuts
> `cms` over + guarantees **referential closure** — the taxonomy capture is **full-public** (`organization_id IS
> NULL` — every public node), so a content ref can only dangle if it points at a *non-public* node. A measured
> cross-surface gene reports any such dangle; prod has exactly **one** (`K-AIFUNX-E658`, a public sim referencing
> a customer-scoped skill) — an operator-owned prod data fix, not a tooling gap. Detail:
> [`../snapshot-spec.md`](../snapshot-spec.md) § the per-stack Directus store fork +
> [`../directus-local.md`](../directus-local.md) § "The data-plane cutover (M23)".

**Skills:** `/demo-up` · `/stack-secrets` · `/stack-snapshot` · `/stack-seed` · `/stack-list` · `/demo-down`
(see the root `CLAUDE.md` skills table). `/stack-secrets` provisions the stack's per-repo `.env` from one secret
source — **values-blind** — and verifies coverage; `/demo-up` runs it as an auto-provision step (M30) so a fresh
demo is self-sourced from the curated secret dir. Mechanism: [`../secrets-spec.md`](../secrets-spec.md).

## Hard constraints (always true)
- **Zero platform-repo change.** All demo tooling lives in `rosetta-extensions` (the demo-stack overlay + the
  seeder + Clerkenstein), never scattered in the rosetta corpus and never authored ad-hoc inside a stack dir.
  It is authored + tested + tagged in the authoring copy at `.agentspace/rosetta-extensions/`, and the demo
  stack consumes a pinned-tag clone at `stack-demo/rosetta-extensions @ <tag>`. The platform clones are
  consumed as-is.
- **Production-safe.** The seeder's isolation guard makes it *structurally impossible* for a non-prod stack to
  write a shared/prod store (Directus, the prod S3-public bucket, live Clerk, Customer.io/Brevo, AI APIs); it
  seeds only the per-stack Postgres/Redis, and proves it with an audit log. Snapshot **capture** is read-only +
  public-only (the tenant-data firewall). The full contract — both halves — is [`../safety.md`](../safety.md)
  (write-side detail in `../seeding-spec.md`, read-side in `../snapshot-spec.md`).
- **Isolated.** Every op is `-p demo-N`-scoped on offset ports with its own data; the dev stack is never touched.
