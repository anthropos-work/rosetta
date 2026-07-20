# Session-Clone Spec — copying anonymized real prod sessions into a demo

**The M232 deliverable (v2.5 "the playbill", Thread B — the write side).** Where
[`content-stories-routes.md`](content-stories-routes.md) (M231) DISCOVERED + PROVED that a content-product
result page reads a persisted DB row a clone could seed, THIS doc specifies the tooling that actually does the
cloning: the `ContentStorySeeder` in `rosetta-extensions/stack-seeding`, which **COPIES real production
job-simulation sessions — the real content, scrubbed of detectable PII — re-tenanted, non-manager-played,
source-pinned** — into a demo org so a presenter can open each one's result page as the player who took it and
as the manager who reviews it.

> **Headline — COPY THE REAL CONTENT, scrub best-effort (data-controller decision, 2026-07-19).** The
> interesting part of a played session IS its free-text — the real conversation, the real LLM feedback, the
> real submission, the real interview report. So the tooling **copies that real content** and **scrubs the
> detectable PII** (real actor names + the source org → placeholders the seeder fills with the demo
> persona/org; emails/phones/urls redacted). This is **NOT "provably clean"**: free-text scrubbing is
> imperfect — a name the pass does not know, an unusual identifier, a company mentioned in passing can
> survive. That **residual re-identification risk is real and was ACCEPTED by the data-controller**; the
> control on it is the **VPN/tailnet scope** of the demo ([`safety.md`](../safety.md) §3.8), not the scrub.

## For PMs — one paragraph

A "content story" is a real, played session a presenter can log into and see the result of. To make those
believable we copy real production sessions into the demo — the real conversation, the real feedback, the real
work — and scrub the personal details we can detect (names, emails, phone numbers, the company name). The scrub
is best-effort: it catches the obvious identifiers but cannot guarantee a name buried in a sentence is gone. We
accepted that residual risk deliberately, and the control is that these demos are reachable **only over our
VPN**, never the open internet. Each clone records exactly which real session it came from (an auditable pin),
references only a **public** simulation, and is owned by a **synthetic employee** (never a manager).

---

## 1. The pipeline — capture+scrub at authoring time, replay offline at seed time

```
  AUTHORING TIME (once, reads prod READ-ONLY)                        SEED TIME (offline, on the demo box)
  ┌───────────────────────────────────────────────┐                ┌──────────────────────────────────┐
  │ cmd/content-capture --dsn <marco_read>         │                │ ContentStorySeeder                │
  │   for each pinned session:                     │   the fixture  │   for each pin: load the copied    │
  │     COPY the real fan-out content              │  ───────────▶  │   content, FILL <<ACTOR_i>>/<<ORG>>│
  │     SCRUB it (package scrub) → placeholders     │  (go:embed)    │   with the demo persona/org, write │
  │     write fixture/content/<key>.json (scrubbed) │                │   the fan-out rows (re-tenanted)   │
  └───────────────────────────────────────────────┘                └──────────────────────────────────┘
        │ the YAML pin list (content-sessions.yaml) drives both, + is disclosed in ──┐
        ▼                                                                             ▼
  seed-generation-manifest.yaml (content_sessions block)                    the demo's per-stack Postgres
```

**Raw customer content never enters an agent's context.** `cmd/content-capture` connects to production
**read-only** (`marco_read` via `~/.pgpass` over Tailscale — [`db-access.md`](../db-access.md); it `SET`s the
session read-only and only `SELECT`s), streams each session's content through the scrub, and writes the
**scrubbed** result to the checked-in fixture. It prints **counts only**, never content. The seeder (seed time)
is fully **offline** — it reads only the go:embed'd fixture. A demo box needs no prod access.

## 2. Stage 1 — sourcing (the reproducible selection)

`contentsession/sourcing.go` BUILDS (never runs) the selection SQL — the reproducible record of *how* the pinned
sessions were chosen. Two load-bearing predicates:

- **Public-anchoring (M231 D6).** A cloned session's `sim_id` must resolve in the demo, which holds only the
  **public** (snapshot-replayed) simulation catalog. So the query INNER-JOINs `directus.simulations` on the
  public predicate — `PublicSimPredicate = "d.private = false AND d.tenant_id IS NULL AND d.status =
  'published'"` — and sources ONLY sessions on a public-published sim.
- **Non-manager-played.** The owner must be a player-vantage member (belt-and-braces: the seeder re-owns every
  clone to a seeded player member anyway — §4).

**What was pinned (v2.5):** 9 real sessions covering type × modality × pass/fail — the assessment set **2 voice
+ 2 code + 1 document**, plus training doc/chat, hiring voice, and the **one public interview sim's** voice
session — passed AND not-passed both. The list is `contentsession/fixture/content-sessions.yaml`.

## 3. Stage 2 — capture + scrub (the copied fixture)

`cmd/content-capture` copies, per pinned session, the REAL result-fan-out content and scrubs it:

| copied facet | table.column | scrub |
|---|---|---|
| LLM feedback | `validation_attempt_results.*_summary`, `validation_attempt_skill_results.*_feedback`, `validation_check_results.feedback` | names (actors **+ the session owner's real name**)→placeholders, emails/phones/urls redacted |
| criterion title + candidate submission | `validation_criterion_results.title`, `.input_data` (jsonb) | ScrubJSON every string leaf |
| the transcript | `interactions.action_payload` (jsonb, capped ≤ 12 turns) | ScrubJSON |
| the code / document work-product | `code_submissions`, `collaborative_assets.content` | scrub stdout/content; base64 source left (technical) |
| the interview report | `interview_extraction_results.user_report`/`manager_report` (jsonb) | ScrubJSON |
| the real skill node-ids | `validation_attempt_skill_results.skill`, `.criterion.skills` | kept as-is (public taxonomy, non-PII) |

**The scrub** (`package scrub`, tested): the source session's real person-names → `<<ACTOR_i>>` placeholders,
the **source org name** → `<<ORG>>`, and **emails / URLs / long digit-runs** → redaction markers. The real
names are used only as scrub targets and are **dropped** — the fixture stores placeholders, not names.

> **The names it sources (M235 leak fix, 2026-07-19; #M235-B1).** The candidate's real first name is threaded all through
> the LLM feedback, and it comes from the **session OWNER's `public.users` identity** (`sessions.owner_id` →
> `firstname`/`lastname` + the email local-part) — **not** from `jobsimulation.actors` (whose `username`/`alias`
> are empty for these sessions). The original capture sourced only the (empty) actor names, so it removed **zero**
> names and 8/9 fixtures shipped a real first name (USER-BLOCKER-M235-01). The capture now sources the owner's
> real identity → the **player placeholder `<<ACTOR_0>>`**, and **token-splits every person-name** (full name +
> each ≥3-char whitespace token, word-boundary-matched) so a **bare first-name** mention is caught. Word-boundary
> matching means a short token never corrupts a common word ("Ann" ≠ "announce").

Two gates guard the fixture, complementary:
1. **Capture-time leak post-condition** (`scrub.SurvivingToken`, in `cmd/content-capture`) — the *name-leak* gate.
   The capture knows the sourced names in-process, so after scrubbing a session it asserts **no sourced name token
   survives** any free-text leaf and **refuses to write the fixture** if one does (it prints only the field name +
   token length, never the value). This catches the arbitrary-first-name leak the offline gate cannot.
2. **Offline fixture-cleanliness gate** (`TestEmbeddedContent_NoStructuralPII`) — re-scans every shipped blob and
   fails on any surviving email/URL/phone, **and asserts the set carries the `<<ACTOR_0>>` placeholder** (the
   "sourced zero names → zero placeholders" regression tripwire). It cannot know arbitrary names offline — that is
   the accepted residual (§6), which the capture-time gate above closes for the *sourced* identity.

The scrub is **deterministic**: same session + same scrub → byte-identical fixture (the source-pin reseed
contract). The pins are disclosed in `seed-generation-manifest.yaml`'s `content_sessions` block (honesty-gated).

## 4. Stage 3 — replay (the seeder)

`ContentStorySeeder` (surface `content-stories`; `DependsOn` users + taxonomy + content; `PerStackIsolated`)
iterates the pins and REPLAYS the copied content into the **first non-hiring (Workforce) story org**, owned by a
**distinct non-hero, MEMBER-role** slot (via `roleForIndex`) → **owner-is-player-vantage, never a manager
seat**. It fills the `<<ACTOR_i>>` placeholders with a minted synthetic display name per actor slot and `<<ORG>>`
with the demo org name, then writes, in FK order, all idempotent on `id`:

```
jobsimulation.sessions                               (ended, completed, passed/failed — G14 enums, org-scoped)
  ├─ validation_attempt_results                      (the REAL summaries, filled; evaluation_status = the gate)
  │    ├─ validation_attempt_skill_results           (the REAL skill node-ids + the REAL feedback)
  │    │    └─ validation_criterion_results           (the REAL titles/input_data; input_format per capture)
  │    │         └─ validation_check_results          (the REAL grader feedback)
  ├─ actors                                          (player = the owner; the copied roles; minted names)
  ├─ interactions                                    (the REAL transcript; action_type ∈ {email,call}; filled payload)
  ├─ code_submissions + collaborative_assets         (the REAL code / document work-product)
  └─ interview_extraction_results                    (the REAL user_report + manager_report, filled)
public.local_jobsimulation_sessions                  ← THE MIRROR (the score source the manager scoreboard reads)
```

### The three seeding landmines it honors (M231 §7)

1. **Co-write the manager MIRROR** (`public.local_jobsimulation_sessions`) or the manager scoreboard is blank.
2. **Reference only public-anchored sims** — the pinned `sim_id` IS public-published, so it resolves.
3. **Enable the interview PostHog flags** — §5.

Plus: owner-is-player-vantage; the copied enums are real (G14-valid) with a clamp for a rare non-terminal
value; the skill node-ids are the REAL ones the candidate was assessed on (real public taxonomy → resolve).

### It COPIES, it never fabricates

Every free-text row is the customer's content, scrubbed + placeholder-filled — proven by
`TestContentStorySeeder_CopiesRealContent` (the seeded `quick_summary` equals the captured content with
placeholders filled, byte-for-byte). `TestContentStorySeeder_PlaceholdersFilled` proves no placeholder token
survives; `TestContentStorySeeder_ReTenantsNoSourcePins` proves the prod source-session-id is provenance-only
(never a live id — every live id is re-keyed off the content-story key + re-tenanted to the demo org).

### Net-new modality substrate

- **Transcript** (`actors` + `interactions`): action_type ∈ {`email`,`call`} only (the DB enum; a COPY bypasses
  Ent, so an invalid value would insert-but-be-invisible — the G14 class). `source_id <> target_id` (the CHECK).
- **CODE / DOCUMENT**: the copied `code_submissions` / `collaborative_assets`.
- **INTERVIEW**: the copied `interview_extraction_results.user_report`/`manager_report`. (Render fidelity of the
  interview surface — the exact plan-section match — is M235's "prove-it-lands" concern.)

## 5. The interview render flags — a sha-pinned demopatch (M231 D3)

The interview result surfaces gate on `posthog.isFeatureEnabled('flag_interview_{player,manager}_report')`, and
a demo bakes **no PostHog**, so the flags resolve `undefined` forever. Two sha-pinned `demopatch`es on the demo's
OWN ephemeral clone — **the interview twin of the M219 `next-web-aireadiness-flag-gate`** — widen the two gates
ONLY when PostHog is entirely unconfigured (behaviour-identical off-demo):

- `next-web-interview-flag-container` — the report **FETCH** gate (`AISimulationResultContainer.tsx`).
- `next-web-interview-flag-result` — the report **RENDER** gate (`AISimulationResult.tsx`).

Both live in the SHARED `packages/ui`, so they bake into BOTH the apps/web and apps/hiring images (wired into
`up-injected.sh`'s both frontend builds + the patchset fingerprint + the LIFO revert trap;
`tests/test_interview_flag_patch_m232.py` fences the manifest shape + wiring + a live-anchor drift pin). **Zero
platform-repo edits.**

## 6. Safety — the read-side exception, honestly bounded

Copying real customer sessions is a user-accepted (data-controller, 2026-07-19) softening of `safety.md`'s
"nothing behind the door", bounded three ways:

1. **Best-effort scrub** — the detectable PII (known names — **actors + the session owner's real identity**,
   token-split so a bare first name is caught — org, emails/phones/urls) is removed; the capture **fails closed**
   if a *sourced* name survives (`scrub.SurvivingToken`), and the fixture is re-scanned for structural PII + the
   name-scrub-fired tripwire at test time. It is NOT a guarantee — a name the pass never *sourced* (e.g. a third
   party mentioned in passing in the free-text) can still survive.
2. **Residual risk ACCEPTED, VPN/tailnet-scoped** — the data-controller accepted the residual re-identification
   risk; the control is that content-story demos are reachable **only over a Tailscale tailnet/VPN** (the Part-3
   exposure posture), never the public internet.
3. **Source-pinned + disclosed** — every clone's prod source-id + the copy+scrub posture is recorded in the
   `content_sessions` manifest block.

**Part 2 (never-write-prod) is untouched** — the seeder writes only per-stack Postgres, audited, n=0-guarded.
The read (`content-capture`) is read-only. See [`safety.md`](../safety.md) §3.8 (the amendment this milestone
lands).

## See also
- [`content-stories-routes.md`](content-stories-routes.md) — the M231 spike (the read/route/sourcing contract this realizes).
- [`safety.md`](../safety.md) §3.8 — the copy+scrub read-side exception + the accepted residual risk.
- [`../db-access.md`](../db-access.md) — the read foundation + the public-vs-customer boundary the capture honors.
- [`seed-manifest-spec.md`](seed-manifest-spec.md) — the `seed-generation-manifest.yaml` the `content_sessions` pins fold into.
- [`stories-spec.md`](stories-spec.md) + [`../seeding-spec.md`](../seeding-spec.md) — the 7-table fan-out + the mirror-pair the substrate extends.
- [`demopatch-spec.md`](demopatch-spec.md) — the demo-patch mechanism the interview flag-gates use.
- The manifest / cockpit halves: **M233** (`content-stories-spec.md`) + **M234** (the cockpit tab) + **M235** (prove-it-lands render proof).
