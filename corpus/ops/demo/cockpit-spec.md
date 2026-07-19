# The Presenter Cockpit (UX spec)

**The reference for the presenter cockpit — the slick, light login launcher a demo-giver drives.** When a
storytelling demo is up (the default — `/demo-up N` seeds the [Stories & Heroes](stories-spec.md) world and
wires the multi-identity fake FAPI), it serves a **standalone panel** on an offset port that lists every seeded
story → its hero trio. The presenter picks a hero, clicks **Log in as**, and the browser lands logged-in as her
on a sensible per-role screen — so the demo is a menu, not a hunt for the right URL + a typed login.

This doc is the canonical description of the cockpit's **UX surface** + its **deep-link contract**: how the
panel renders, what the single CTA does, and where the served panel sits relative to the platform (it never
touches it). It graduates the cockpit mechanics that were scattered across
[`stories-spec.md` § The presenter cockpit (M38)](stories-spec.md#the-presenter-cockpit-m38) (the original
M37/M38 producer/consumer seam) and [`clerkenstein.md`](../../services/clerkenstein.md) (the handshake) into
one place, and layers the **v1.10 "method acting" M43 UX pass** on top.

> **Scope.** This doc is the **UX + deep-link** reference for `rosetta-extensions/demo-stack/cockpit.py` (the
> served panel) + the manifest its menu reads (projected by `stack-seeding/seeders/cockpit.go`). The
> **multi-identity FAPI handshake** the CTA drives lives in [`clerkenstein.md`](../../services/clerkenstein.md)
> (unchanged by M43); the **roster/seed producer seam** (how a hero's seat resolves, `--roster-export`) lives in
> [`stories-spec.md`](stories-spec.md). This doc owns the panel a presenter looks at.

---

## For PMs — the demo-driving surface

A demo flows: *show Maya's verified-skill profile → show Tom's stark claimed-vs-verified gap → log in as their
manager Dan and watch the same two people become the standout high/low rows of his Workforce dashboard.* The
cockpit is the **remote control** for that flow. It is a clean, professional, **light** panel — a card per hero,
the hero's name + role + a one-line presenter note, the vantage/trajectory badges, and **one button**: **Log in
as**. One click logs in as the chosen hero **and** lands the browser on the screen that tells her part of the
story — no typed login, no copied URL.

```
Presenter Cockpit — demo-3
  AI TRANSFORMATION & RESKILLING     🏢 Cervato Systems
  220 people · 220-person software co upskilling
    ◍ Maya Chen — Backend Developer   [EMPLOYEE] [THRIVING]
        "8 verified skills, mobility-ready"                           [⇥ Log in as]
    ◍ Tom Becker — Backend Developer   [EMPLOYEE] [STRUGGLING]
        "Few/low verified skills, OVER-rates himself (stark gap)"     [⇥ Log in as]
    ◍ Dan Rossi — Engineering Manager  [MANAGER]
        "Team gaps, role-readiness, succession (Maya), at-risk (Tom)" [⇥ Log in as]
  SDR ONBOARDING & RAMP              🏢 Solvantis
  120 people
    ◍ Sara Whitfield [EMPLOYEE·THRIVING] / Nick Alvarez [EMPLOYEE·STRUGGLING] / Leah Donovan [MANAGER]

  ⬇ Download seed manifest
```

(The `◍`/`🏢`/`⇥` glyphs above are stand-ins for the FontAwesome icons the real panel renders:
`fa-circle-user` per hero, `fa-building` per org, `fa-arrow-right-to-bracket` on the CTA.)

**Login is the only action.** Earlier cockpits (M38) had two buttons per hero — `[Login as]` (landed on the app
root) and `[Jump to section]` (landed on the hero's deep-link). M43 **unified** them: the single **Log in as**
both logs the presenter in as the hero **and** lands her on a sensible per-role screen (an end-user on her
profile, a manager on the Workforce dashboard). One CTA, one click, the right screen.

**A login takes about a second.** Clicking **Log in as** raises a small **staged overlay** (*Signing you in… →
Loading your workspace… → Almost there…*) so the load has feedback instead of looking frozen. The overlay is
feedback, not a progress bar — there's no real signal from the cross-origin handshake to report.

> #### ⚠ CORRECTION (M218 "seat change", v2.3) — this doc used to say *"~2–5 seconds we can't shorten"*
>
> **That claim was never measured, and it was wrong by 20–40×.** The first instrument ever pointed at this path
> (M218 iter-02) measured click→ACCESS at **p95 39.45 s** (employee) / **38.30 s** (manager) on a remote demo.
> The "we can't shorten it" half was wrong too: **M218 shortened it.**
>
> | | before M218 | after M218 |
> |---|---|---|
> | **employee** (`maya-thriving` → `/profile`) | p95 **39.45 s** | **p95 1.46 s** (p50 1.00 s) |
> | **manager** (`dan-manager` → `/enterprise/…`) | p95 **38.30 s** | **p95 1.40 s** (p50 1.12 s) |
>
> *(5 consecutive cold logins/vantage, `billion` tailnet demo, `autoverify` green — the presenter's real vantage.)*
>
> **Two defects, both in the demo tooling, neither in the platform:**
> 1. next-web's **server-side** GraphQL origin was the **build-inlined public URL**, which is unreachable from
>    inside its own container → a 10.5 s connect-timeout × 3 attempts + backoff ≈ **37.5 s** of blocked SSR on
>    *every* authenticated render. Fixed with a server-only `WUNDERGRAPH_SSR_ENDPOINT` (a sha-pinned demo-patch).
> 2. Clerkenstein's fake **BAPI** served a **hardcoded stub user** to every hero, so `currentUser().externalId`
>    disagreed with the JWT's identity → `app` refused `userPreferences` → a `retry: 2` / 2 s+4 s ladder burned a
>    further **~6 s**. Fixed by making the BAPI **roster-aware**.
>
> **The lesson is not the numbers — it is that an unmeasured number sat in this doc, asserting its own
> unfixability, for four releases, and that is exactly why nobody investigated.** Booked as an M43 scope-`Out:`
> with decision **D5** and *zero* deferrals recorded, so it never entered a ledger. Don't write "we can't fix
> this" about something you have never measured. See [`latency-budget.md`](latency-budget.md).

**Grab the seed manifest.** A footer **Download seed manifest** link saves the **consolidated
`seed-generation-manifest.yaml`** (v1.10b M52) — the single auditable file inlining the whole seed+generation
intent (all 3 orgs' population + the mother prompt + batch config + snapshot sources; cache/generated data
excluded), so a presenter or auditor can read the entire demo-world direction in one place. (An old bring-up
with no manifest wired falls back to saving the cockpit's JSON menu.) See
[`seed-manifest-spec.md`](seed-manifest-spec.md).

---

## For engineers — how it works

### Standalone served panel, never an in-app overlay (D15)

The cockpit is a host-native HTTP server (`rosetta-extensions/demo-stack/cockpit.py`, **stdlib-only** Python) on
an **offset port** (`7700 + N·10000`, e.g. `37700` for `demo-3`), brought up with the stack
(`up-injected.sh`, session-detached via `launch_detached` so it survives the bring-up task being reaped) and
torn down with it. It binds `127.0.0.1` by default; **`0.0.0.0` only under `--public-host`**. It is **never** an
edit to next-web — it reaches the platform only as a browser would (over the FAPI handshake), preserving the hard
zero-platform-repo-edit line. The serve is **non-fatal**: a cockpit failure never aborts an otherwise-good
storytelling demo (the M18/M19 pattern) — **but it is now REPORTED** (below). `DEMO_NO_COCKPIT=1` brings the demo
up without the panel.

#### Teardown is PORT-authoritative, not pid-authoritative (M217)

> **This doc previously said the cockpit is "torn down with the stack (its PID is recorded in
> `<stack>/cockpit.pid`)". That was true of the *intent* and false of the *behaviour*.**

`docker compose down` cannot reach a host-native process, so teardown must reap the cockpit itself. It used to do
so **by PID**, from that pidfile — and that leaks a listener three ways:

1. `launch_detached` writes the pidfile **unconditionally**, *before and independent of* a successful bind. A
   cockpit that died on `EADDRINUSE` still leaves a pidfile naming a dead pid.
2. A second bring-up **overwrites** the pidfile — orphaning the first cockpit **forever**. It keeps serving a
   **stale manifest** on the port and nothing records it.
3. `kill` on a **recycled** pid kills the wrong process, silently.

And the teardown **discarded `kill`'s status**, `rm -f`'d the pidfile regardless, and printed *"stopped the
presenter cockpit"* **either way**. Found in the field on `billion`: an orphaned cockpit still **LISTENING on
`0.0.0.0:17700`** — an unauthenticated *"become any hero"* panel pointing at a database whose containers had been
removed — which **survived a `/demo-down` that reported success**.

Since M217 (`demo-stack/reap.sh`):

- **Teardown reaps by PORT**, which is what actually blocks the next bind. The pidfile is a *hint*, not the truth.
- **The reap is identity-checked.** It kills only listeners whose command line matches (`cockpit.py`, scoped to
  *this* stack's offset port, so a co-resident demo's cockpit is untouchable). A **foreign** process holding the
  port is **reported loudly and left alive** — reaping a port must never become `kill $(lsof -t -i:PORT)`.
- **The bring-up pre-reaps** its own stale predecessor before binding.
- **`cockpit.py` fails cleanly** on a bind conflict (exit 2 + a diagnosis) instead of an unhandled traceback.
- **"presenter cockpit serving on …" is now gated on a real `/healthz` probe.** It used to print unconditionally —
  which is *how an operator drove a dead cockpit for a whole session*. A failed cockpit now says so, and its log
  tail is surfaced.

### Single source — the menu is a projection of the seed (D9)

The cockpit menu is a **manifest** (`cockpit-manifest.json`) the seeder projects from the very file that seeded
the heroes (`stackseed --cockpit-export`, `cockpit.go`'s `BuildCockpitManifest`). The annotations describing a
hero in the cockpit are the same ones that scoped her seed — the menu can never drift from the data. (The demo
tooling is stdlib-only Python, so the YAML is parsed once on the Go side and the panel reads the derived JSON.)
The manifest carries, per hero: the `key` (her `stories.yaml` id — the seat-switch handle), `name`, `role`,
`vantage`/`vantage_label`, `trajectory`, the presenter `annotation`, and a **resolved `jump_to`** (her declared
`jump_to`, or the vantage default via `defaultJumpForVantage` — an end-user → `/profile`, a manager →
`/enterprise/workforce`).

### A `jump_to` pointing at a LEGACY surface is a HARD SEED FAILURE (v2.3 M219)

`WriteCockpitManifest` calls **`ValidateCockpitManifest`**, which refuses any hero whose resolved `jump_to`
matches **`LegacyReadinessPaths()`** (`stack-seeding/seeders/cockpit.go`). **The seed FAILS — it does not warn.**

This exists because **every one of the demo's three AI-readiness pointers targeted the legacy page**, and nothing
caught it for four releases. `/enterprise/workforce/ai-readiness` is an **unlinked orphan**: no nav entry, no
workforce tab, no redirect points at it, and its hook takes no `cycle` param — so it reads the cycle-less endpoint
and renders no cycle picker, no archetype matrix, no people, no How-we-measure, no What-to-do-next. A presenter
driven there sees a shell. The **current** manager surface is **`/ai-readiness`**.

The deep-link catalog is therefore no longer end-user-vs-manager only. It carries **an end-user readiness entry**
too — **`ai-readiness-member` → `/home`** — because the member readiness surface **has no route of its own**
(it is a region of the authenticated landing). *That is precisely why route-crawling never found it.*

Current pointers: `dana-manager → /ai-readiness` · `aria-completed → /home` · `ben-started → /home`.
See [`../../services/ai-readiness.md` § Surfaces](../../services/ai-readiness.md) for the current-vs-legacy split.

### The CTA — one [Log in as] = one FAPI handshake redirect

The single **Log in as** points the browser at the multi-identity fake FAPI's handshake with the hero's
seat-switch key **and** her per-role landing as the redirect:

```
https://<fapi-host>/v1/client/handshake?__clerk_identity=<hero-key>&redirect_url=<app-base><jump_to>
```

`<fapi-host>` is the per-stack fake FAPI on its own offset port `127.0.0.1:<5400 + N·10000>` (e.g.
`127.0.0.1:35400` for `demo-3`), served over HTTPS (clerk-js requires it); the `redirect_url` is the hero's
`jump_to` as an absolute next-web URL on the app's offset port `<3000 + N·10000>`, **fully percent-encoded**
inside the outer query (so a `jump_to` carrying its own `?tab=` survives — a load-bearing escaping invariant
pinned by a test). The FAPI selects the chosen hero's seat from `__clerk_identity` **then** establishes the
session and redirects — so the hero is the active identity *everywhere* (the client view, `/v1/me`, the minted
token, the cookies) **and** the browser lands on her screen, in one move. The key is the hero's `stories.yaml`
id — the **same** key the roster export gave Clerkenstein's registry, so the seat always resolves. (The
handshake + multi-identity selection are M37; see [`clerkenstein.md`](../../services/clerkenstein.md).)

**Why no `jump_to` button anymore.** The M38 cockpit rendered the handshake twice per hero — once landing on the
app root (`[Login as]`), once on the deep-link (`[Jump to section]`). The deep-link landing is strictly the more
useful of the two (it's already a logged-in session, just with a better starting screen), so M43 dropped the
root-landing button and made `[Log in as]` route to `jump_to`. The root-landing helper
(`cockpit.py::login_as_url`) is kept as the documented bare variant (a presenter who wants the app root, an
integration check), but the rendered CTA uses the jump_to seam (`jump_url`).

### The deep-link catalog (O9)

The cockpit ships an enumerated, stable set of next-web routes per vantage (`cockpit.go`'s `DeepLinkCatalog`) —
the *individual* surfaces an end-user hero demos (`/profile`, Skill Spotlight, my-growth, take-a-sim) and the
*org-intelligence* surfaces a manager hero demos (the Workforce dashboard tabs — verification / role-readiness /
succession / mobility — plus the talent pool). A hero's `jump_to` is matched against this catalog so the
manifest can carry its label; an unrecognized `jump_to` still works (it's a raw path).

### The hiring vantage — the recruiter + 2 candidate seats (v2.4 "casting call" M224)

The **4th, HIRING** story org (**Meridian Talent**, `narrative: hiring`, `is_hiring=true` — M223's seed) gets a
**hero trio** on the cockpit (M224), login-only like every other hero:

| Hero (`key`) | Role | Vantage | `jump_to` lands on | What she demos |
|---|---|---|---|---|
| **Rae Ramirez** (`rae-recruiter`) | Talent Acquisition Specialist | `manager` → admin (slot-1, funnel-skipped) | `/enterprise/activity-dashboard` **on the hiring app** | the candidate-comparison Results scoreboard — 20/page × ~43 reachable candidates per each of the 5 shared sims |
| **Cara Nguyen** (`cara-assessed`) | Data Analyst | `end-user` → candidate | `/home` (candidate self-view) | an **assessed** candidate — 5 scored HIRING sessions (ranks on Rae's scoreboard) + a **COMPLETED** assignment; `/home` reads "Completed" |
| **Cody Brenner** (`cody-assigned`) | Business Ops Analyst | `end-user` → candidate | `/home` (candidate self-view) | an **assigned-only** candidate — no sessions + a **PENDING** assignment; `/home` reads "Assigned" |

**The load-bearing cockpit mechanic is `CockpitHero.IsHiring` (a per-hero manifest flag).** A hiring hero's
`[Log in as]` targets the **demo's real Hiring app base** (`--hiring-base`, offset `:13001`), **not** the workforce
`apps/web` base (`:13000`) the other heroes use. The workforce heroes (Dan &c.) stay on `apps/web` unchanged; only
`IsHiring` heroes route to the hiring container. `cockpit.go`'s `BuildCockpitManifest` emits the flag
(`omitempty` — a workforce hero's manifest is byte-identical); `cockpit.py` reads it and picks the base.

**Why a second app base at all — the two-app demo (TOK-02, #M224-D-TOK02).** On the *unmodified* platform an
all-hiring-orgs user is **ejected out of `apps/web` to the standalone Hiring product** by a product-boundary
redirect (`UserStatusContext`, by design) the moment her org reads as hiring (client `publicMetadata.isHiring=true`
— which M224 deliberately wires, because the org must *genuinely read as hiring*). So "reads as hiring" and
"reachable inside `apps/web`" are **mutually exclusive on the real platform**. Rather than fake it with a re-skin,
the demo **runs the genuine `apps/hiring` as a second UI container** (built from the untouched clone, offset port,
same fake FAPI + same Cosmo backend + same seeded Postgres). The platform's **own** symmetric guard keeps the
recruiter *in* the hiring app; she reads the **same** seeded `local_jobsimulation_sessions` the scoreboard reads.
No forcing, no fiction — see [`../../services/hiring.md`](../../services/hiring.md) § the render path and
[`demopatch-spec.md`](demopatch-spec.md) § the four hiring-image patches.

**DeepLinkCatalog note.** The recruiter's `jump_to` is a **raw path** (`/enterprise/activity-dashboard`); a
dedicated per-`[simId]` `NeedsID` catalog entry was judged **optional polish** and **not** added — the raw jump
suffices and the render gate was met without it (#M224-scope). The candidate heroes land on `/home` because
`apps/hiring`'s `/profile*` is **admin-gated at platform source** (`role !== Admin → HOME_URL`), so a candidate's
faithful landing is the real candidate self-view, differentiated per funnel state (Completed vs Assigned).

### The UI surface (v1.10 M43)

The panel is a single static HTML page (`render_page()`), restyled and enriched:

- **Light professional theme.** `_PAGE_CSS` is a clean light design — CSS custom properties (one indigo accent),
  a card-per-hero hierarchy with subtle shadows + hover, generous spacing, high-contrast typography. (Replaced
  the original dark GitHub-style theme.)
- **FontAwesome icons via the free CDN.** A cdnjs FA6-free `<link>` (with its SRI integrity hash) in `<head>`;
  `fa-circle-user` as each hero's avatar, `fa-building` before each org, `fa-arrow-right-to-bracket` on the CTA,
  `fa-download` in the footer. **The CDN `<link>` is a runtime asset, not a build dependency** — the panel stays
  stdlib-only Python and the supply-chain posture stays GREEN. _(Offline-safe precedent, recorded for a future
  offline-demo need: `ant-academy` vendors FA Pro locally under `code/public/assets/fontawesome/`. Not adopted
  here — the CDN is the chosen default.)_
- **Seed-manifest download (v1.10b M52 repoint).** The footer **[Download seed manifest]** now serves the
  **CONSOLIDATED** `seed-generation-manifest.yaml` — the single auditable file inlining the whole seed+gen
  intent (population + the file-resident mother prompt + the batch config + the snapshot sources; cache +
  generated data excluded — see [`seed-manifest-spec.md`](seed-manifest-spec.md)). It is served at
  `/seed-generation-manifest.yaml` with `Content-Disposition: attachment` (verbatim YAML). The link falls
  back to the menu manifest (`/manifest.json`) when no consolidated manifest is served (an old bring-up), so
  it is never dead. The **menu** manifest (`cockpit-manifest.json`, the stories→heroes projection) still
  drives the `[Log in as]` CTAs and stays served at `/manifest.json`.
- **Staged login-progress overlay.** A small JS overlay (`_OVERLAY_JS`, a raw string — no manifest data is
  interpolated into the JS, so no injection surface) shows on `[Log in as]` click: *Signing you in… → Loading
  your workspace… → Almost there…* on a deterministic timer with a generous final stage. It covers **only the
  brief forward window** between the click and the browser leaving for next-web. The overlay never
  `preventDefault`s the real navigation — it is purely additive feedback over the login latency, which **M218
  cut from ~39 s to ~1.4 s p95** (see the correction above; this line used to call a never-measured "~2-5s"
  latency *unavoidable*).
  > #### ⚠ CORRECTION (v2.3.1 hotfix) — the overlay no longer re-shows on back/reload
  > This doc used to say the `localStorage` "in-flight login" flag (a **30s** freshness window) made a
  > **back-navigation to the cockpit re-show the overlay** rather than a dead spinner. That behavior was a
  > **bug**: a presenter switches heroes well within 30 s, so on essentially every back/reload after a login the
  > cockpit re-showed *"Loading your workspace…"* over an already-loaded workspace — defeating the panel's one
  > job (an instant hero-switch). It is this release's own **D17 shape**: a status artifact (*"a login is in
  > flight"*) outliving the thing it describes, read as evidence that something is loading when nothing is. The
  > 30 s window only ever made sense when a login took ~39 s (pre-M218). **Fixed** (`cue-to-cue-v2.3.1`): any
  > **return** to the cockpit — back button, reload, or bfcache restore — now **clears the flag and hides the
  > overlay** (`resetOverlayOnReturn`, on both fresh load and `pageshow`); the overlay can no longer re-show on
  > return. Gated by `stack-verify/e2e/tests/cockpit-overlay-return.spec.ts` (RED against the pre-fix cockpit).

### Served endpoints

| Path | Response |
|------|----------|
| `GET /` (or `/index.html`) | The cockpit HTML page |
| `GET /healthz` | `200 "ok"` (a liveness probe) |
| `GET /seed-generation-manifest.yaml` | (v1.10b M52) The **consolidated** seed+generation manifest as a download (`Content-Disposition: attachment`, verbatim YAML) — the footer **[Download seed manifest]** target, when served (`--seed-manifest`) |
| `GET /manifest.json` | The **menu** manifest (stories→heroes projection) as a download (`Content-Disposition: attachment`, pretty JSON) — the `[Log in as]` source + the fallback download |
| `GET /content-manifest.json` | (v2.5 M234) The **content-stories** menu (the M233 `content_products[]` projection) as a download — the 2nd "Content stories" tab's source. Served only when `--content-manifest` was passed (a `404` otherwise) |
| anything else | `404` |

### The 2nd tab — "Content stories" (v2.5 "the playbill" M234)

When the cockpit is launched with `--content-manifest <content-manifest.json>` (the bring-up threads it when
the storytelling demo is up), the panel renders a **client-side tab toggle** with a 2nd **"Content stories"**
tab beside "Org stories": per content product, a list of **played sessions** each with **as-player /
as-manager** login-and-land CTAs (the `content-player-<idx>` seats are registered in the same roster the
`[Log in as]` heroes come from). Absent `--content-manifest` ⇒ no tab bar (byte-identical to today). The full
render + seat + routing contract is [`content-stories-spec.md` §7](content-stories-spec.md#7-the-cockpit-render--the-2nd-content-stories-tab-m234).

### Bring it up

```bash
# A storytelling demo is the DEFAULT: /demo-up N seeds the 3-org hero trio, wires the multi-identity
# fake-fapi, and serves the cockpit on http://localhost:$((7700 + N*10000)).
/demo-up 3
# → the cockpit serves on http://localhost:37700. Pick a hero → [Log in as] → land on her per-role screen.
```

`DEMO_NO_COCKPIT=1` brings the stories demo up without the panel (e.g. an API-only run);
`DEMO_NO_STORIES=1` brings up the structural small-200 fallback (no heroes, single identity, no cockpit). The
cockpit + roster + seed all pin `--stack demo-N`, so the exported ids and the seeded rows are guaranteed to
match.

---

## Future-feature expansion surface

The cockpit is deliberately **UX-only** today — a launcher, not a presenter console. The natural next surfaces,
recorded so a future milestone has a home (none is in v1.10 scope):

- **Per-hero history / telemetry** — what the presenter showed, last-logged-in hero, a "you are currently
  logged in as …" indicator.
- **Note-taking / a talk-track** — the presenter's per-hero script alongside the annotation.
- **A search / filter** as the roster grows (multi-org, many heroes).
- **Live seed status** — surface a re-seed / re-snapshot from the panel (today the cockpit only reads the
  manifest; the seed/snapshot lifecycle is the `/stack-*` skills).

Any of these is an **additive** change to the served panel — the hard zero-platform-repo-edit line + the
single-source (D9) + stdlib-only posture all hold.

---

**See also:** [`seed-manifest-spec.md`](seed-manifest-spec.md) (the consolidated seed+generation manifest the
[Download] serves, v1.10b M52) · [`stories-spec.md`](stories-spec.md) (the seeded world + the roster/seat producer seam) ·
[`clerkenstein.md`](../../services/clerkenstein.md) (the multi-identity FAPI handshake) ·
[`frontend-tier.md`](frontend-tier.md) (the demo UI tier the cockpit launches into) ·
[`rosetta_demo.md`](../rosetta_demo.md) (the demo-stack lifecycle).

### The reap's safety rule (M217 hardening)

The port reap has one non-negotiable rule, and it is worth stating plainly because it constrains every future
change to `reap.sh`:

> **If we cannot PROVE a listener is ours, we do not touch it.**

Three cases, deliberately distinguished:

| The listener… | What we do |
|---|---|
| matches our identity regex on **this stack's** offset port | **kill it** (TERM, then KILL) |
| does **not** match — someone else's server | **report it loudly, leave it alive** |
| has **no readable command line** (on Linux, `ss` only reveals pids you own — a **root-owned** listener is opaque) | **treat as foreign.** Report, do not kill |
| was there on the probe and **gone** by the re-probe (it raced away) | **say nothing** — benign |

And if the host has **none** of `lsof`/`ss`/`fuser`, the reap says **"CANNOT CHECK"** rather than reporting the
port free. A blind process-killer that answers *"clear"* is worse than one that admits it cannot see.
