# Session-Clone Spec — cloning anonymized real prod sessions into a demo

**The M232 deliverable (v2.5 "the playbill", Thread B — the write side).** Where
[`content-stories-routes.md`](content-stories-routes.md) (M231) DISCOVERED + PROVED that a content-product
result page reads a persisted DB row a clone could seed, THIS doc specifies the tooling that actually does the
cloning: the `ContentStorySeeder` in `rosetta-extensions/stack-seeding`, which **copies real production
job-simulation sessions — anonymized by construction, re-tenanted, non-manager-played, source-pinned** — into a
demo org so a presenter can open each one's result page as the player who took it and as the manager who reviews
it.

> **Headline — the safety property is ANONYMIZE BY CONSTRUCTION.** The clone is of a real session's
> *structure*, not its *content*. Only the **non-PII skeleton** (the prod source-session-id, the real public
> sim_id, the sim_type, the modality, the score, pass/fail, the duration, the actor/interaction counts) is
> ever *sourced* from production. Every **free-text** facet — LLM feedback, transcript, the candidate's
> submission, actor names, the interview report — is **SYNTHESIZED at seed time, never copied.** So the
> checked-in fixture and the seeded demo are **provably PII-free**: there is nothing to leak, not even
> transiently. This is the bounded exception [`safety.md`](../safety.md) §3.8 records.

## For PMs — one paragraph

A "content story" is a real, played session a presenter can log into and see the result of. To make those
believable we clone real production sessions into the demo — but we clone their *shape*, not their *words*. The
tooling reads only the non-identifying skeleton of a real session (its score, how long it took, what kind of
simulation it was, how many turns it had) and then writes a demo session with the same shape but **freshly
made-up** feedback, transcript, and names. No customer's actual words, name, or work ever enters the demo. Each
clone records exactly which real session it was shaped from (an auditable pin), references only a **public**
simulation everyone can see, is owned by a **synthetic employee** (never a manager), and is reachable **only over
our VPN** — the same access bound as every other demo.

---

## 1. The pipeline — three stages, one of them offline

```
  AUTHORING TIME (once, by a human, against read-only prod)          SEED TIME (offline, on the demo box)
  ┌───────────────────────────────────────────────┐                ┌──────────────────────────────────┐
  │ 1. SOURCE  (contentsession/sourcing.go)        │                │ 3. RECONSTRUCT                     │
  │   the public-anchored + non-manager + per-cell │   the fixture  │   (ContentStorySeeder)            │
  │   selection query → pick INTERESTING sessions  │  ───────────▶  │   replay the skeleton + SYNTHESIZE │
  │ 2. PIN     (contentsession/fixture/*.yaml)     │   (go:embed)   │   the free-text into the demo org  │
  │   freeze the NON-PII skeleton + the source-id  │                │   (owner = a seeded player member) │
  └───────────────────────────────────────────────┘                └──────────────────────────────────┘
        │ discloses the pins into ─────────────────────────────────────────┐
        ▼                                                                   ▼
  seed-generation-manifest.yaml (content_sessions block)          the demo's per-stack Postgres
```

**The seeder never touches production.** Sourcing (stages 1–2) is an **authoring-time** activity a human runs
once, against the read-only `postgres` MCP ([`db-access.md`](../db-access.md)); its output is a checked-in
fixture. The seeder (stage 3) is fully **offline** — it reads only the go:embed'd fixture + the stack's own
replayed public taxonomy. A demo box needs no prod access.

## 2. Stage 1 — sourcing (the reproducible selection)

`contentsession/sourcing.go` BUILDS (never runs) the selection SQL — the reproducible record of *how* the pinned
sessions were chosen. Two load-bearing predicates:

- **Public-anchoring (M231 D6).** A cloned session's `sim_id` must resolve in the demo, which holds only the
  **public** (snapshot-replayed) simulation catalog. So the query INNER-JOINs `directus.simulations` on the
  public predicate — the single-sourced constant `PublicSimPredicate = "d.private = false AND d.tenant_id IS
  NULL AND d.status = 'published'"` — and sources ONLY sessions on a public-published sim. A session on a
  customer-private sim is excluded (its content is outside the public snapshot envelope).
- **Non-manager-played.** The owner must be a player-vantage member — a manager reviews, she does not play.
  (Belt-and-braces: the seeder re-owns every clone to a seeded player member anyway — §4.)

The query reads **only non-PII columns** (id, sim_id, sim_type, score, completion_status, timing, and fan-out
COUNTS) — never a free-text value — honoring the read boundary. Modality is derived by joining
`directus.sim_tasks.task_type` (`call`→voice, `code`→code, `collaborative_doc`/`send_attachment`→document, else
chat). Ordering surfaces the richest-fan-out candidate first, so a human pins an *interesting* session.

**What was pinned (v2.5):** 9 real sessions covering the type × modality × pass/fail matrix — the assessment set
**2 voice + 2 code + 1 document** (satisfies "2 voice + 1 code + 1 document"), plus training doc/chat, hiring
voice, and the **one public interview sim's** voice session — with **passed AND not-passed** both represented.
The list is `contentsession/fixture/content-sessions.yaml`.

## 3. Stage 2 — the source-pin fixture (anonymize by construction)

The checked-in `content-sessions.yaml` stores, per session, ONLY the non-PII skeleton:

| field | source | PII? |
|---|---|---|
| `source_session_id` | the prod `jobsimulation.sessions.id` (the provenance pin) | no — an opaque uuid, identifies nothing without prod access |
| `sim_id` | the REAL public sim | no — public catalog |
| `sim_type` · `modality` · `passed` · `score` · `duration_seconds` | the session's non-PII descriptor | no |
| `actor_count` · `interaction_count` | the transcript cardinalities | no |

**No free-text is stored.** The fixture is validated at load (`contentsession.Load`): uuid pins, **unique keys +
unique source pins** (the source-pin contract admits no duplicate clone of one real session), valid
sim_type/modality, sane score band — fail loud on any violation. It is `go:embed`'d (the exhibit set is a fixed,
code-owned artifact, like the AI-readiness prompts — not per-stack config), so the seeder is self-contained.

**Source-pin disclosure.** The pins are projected into `seed-generation-manifest.yaml`'s `content_sessions`
block (`manifest.buildContentSessions`, honesty-gated by `TestManifest_CanonicalFileMatchesProjection` — a
single source, cannot drift), so an auditor reads *exactly* which real sessions a content-story demo is sourced
from, in one file, with the anonymization posture stated once at the block level.

## 4. Stage 3 — reconstruction (the seeder)

`ContentStorySeeder` (surface `content-stories`; `DependsOn` users + taxonomy + content; `PerStackIsolated`)
iterates the embedded fixture and reconstructs each session's full result substrate into the **first non-hiring
(Workforce) story org**, owned by a **distinct non-hero, MEMBER-role** population slot (resolved via
`roleForIndex` — the same single-source role fn the UsersSeeder wrote) → **owner-is-player-vantage by
construction, never a manager seat**. It writes, in FK order, all idempotent on `id` (deterministic keys →
byte-reproducible reseed):

```
jobsimulation.sessions                               (ended, completed, passed/failed — G14-valid enums, org-scoped)
  ├─ validation_attempt_results                      (evaluation_status = THE gate; success_threshold, score)
  │    ├─ validation_attempt_skill_results           (skill = a REAL public node-id, resolve-or-drop; competency_level_score)
  │    │    └─ validation_criterion_results           (type=evaluation; input_format per modality)
  │    │         └─ validation_check_results          (engine llm|text_diff; success; essential)   [NET-NEW]
  ├─ actors                                          (player = the owner; AI stakeholders; SYNTHETIC names)   [NET-NEW]
  ├─ interactions                                    (transcript; action_type ∈ {email,call} ONLY; SYNTHETIC payload)  [NET-NEW]
  ├─ code_submissions + collaborative_assets         (the CODE / DOCUMENT work-product; SYNTHETIC)   [NET-NEW]
  └─ interview_extraction_results                    (user_report + manager_report; SYNTHETIC plan-shaped envelope)  [INTERVIEW]
public.local_jobsimulation_sessions                  ← THE MIRROR (the score source the manager scoreboard reads)
```

### The three seeding landmines it honors (M231 §7)

1. **Co-write the manager MIRROR** (`public.local_jobsimulation_sessions`) or the manager scoreboard is blank
   (the M219/M222 trap — the scoreboard reads the app-side event-populated mirror, not the runtime table).
2. **Reference only public-anchored sims** — the pinned `sim_id` IS a public-published sim, so it resolves in
   the replayed catalog and the result page renders the real sim.
3. **Enable the interview PostHog flags** — §5.

Plus the standing rules: owner-is-player-vantage (never a manager seat), all G14-valid enums, closure stays
green (skill node-ids are drawn from the replayed taxonomy, **resolve-or-drop — never fabricated**).

### Anonymization surface — what is synthesized (per M231's contract)

| real facet (PII risk) | how M232 anonymizes |
|---|---|
| structured ids (`owner_id`, `organization_id`, `session_id`, tokens, …) | **re-keyed** (deterministic per content-story key) + **re-tenanted** into the manifest org; tokens regenerated |
| timestamps | **shifted** to backdate (the real duration is preserved) |
| actor `username`/`alias` (direct-PII names) | **synthesized** from a clearly-synthetic roster (no real person) |
| LLM feedback (`*_summary`, `*_feedback`) | **synthesized** from pass/fail + score band |
| criterion `input_data` / the candidate's submission | **synthesized** (code snippet / document text) |
| `action_payload` (the transcript, highest PII risk) | **synthesized** turns (bounded ≤ 12), never the real payload |
| `interview_extraction_results.user_report`/`manager_report` | **synthesized** plan-shaped `{"results":{…}}` envelope |

**The re-key is provable.** `TestContentStorySeeder_AnonymizeNoSourcePins` asserts the raw prod
`source_session_id` never appears as a live id in any written row — every live id is derived from the
content-story key, so the source pin is provenance-only.

### Net-new modality substrate (the M232 build)

- **Transcript** (`actors` + `interactions`): the DB enum admits **only** `email` + `call` action_types (a COPY
  bypasses Ent, so an invalid value would insert-but-be-invisible — the G14 class); VOICE→`call`,
  everything else→`email`. `source_id`/`target_id` FK actors, `source_id <> target_id` (the DB CHECK).
- **CODE**: a completed `code_submissions` row (runtime `py`, language_id 71, base64 synthetic source) + the
  `collaborative_assets` the code criterion (`input_format=collaborative_asset`) grades. (There is no `code`
  input_format value — code grades via the editor diff.)
- **DOCUMENT**: the authored `collaborative_assets` row.
- **INTERVIEW**: `interview_extraction_results.user_report`/`manager_report` as the plan-shaped
  `{"results":{…}}` `ExtractionData` envelope with the `score_grade` (A/C) session-quality section (the header
  quality badge) + narrative sections; the manager report adds `attention_points` + a recommendation. The
  backend stores the report bytes opaquely (no struct enforcement), so the row is insertable; **exact
  plan-section-id alignment for FULL render fidelity is M235's ("prove-it-lands") concern** — its coverage
  iteration triages any blank landing to its read-model.

## 5. The interview render flags — a sha-pinned demopatch (M231 D3)

The interview result surfaces gate on `posthog.isFeatureEnabled('flag_interview_{player,manager}_report')`, and
a demo bakes **no PostHog** (`Analytics.provider` inits it only when both `NEXT_PUBLIC_POSTHOG_KEY` + `_HOST` are
present), so the flags resolve `undefined` forever and the report never fetches or renders. A seeded row is
**necessary but not sufficient**. Two sha-pinned `demopatch`es on the demo's OWN ephemeral clone
([`demopatch-spec.md`](demopatch-spec.md)) — **the interview twin of the M219 `next-web-aireadiness-flag-gate`** —
widen the two gates ONLY when PostHog is entirely unconfigured (i.e. exactly a demo), behaviour-identical
off-demo:

- `next-web-interview-flag-container` — the report **FETCH** gate (`AISimulationResultContainer.tsx`,
  `isExtractionEnabled`).
- `next-web-interview-flag-result` — the report **RENDER** gate (`AISimulationResult.tsx`). Both are needed;
  both live in the SHARED `packages/ui`, so they bake into BOTH the apps/web and apps/hiring images (wired into
  `up-injected.sh`'s both frontend builds + the patchset fingerprint + the LIFO revert trap;
  `tests/test_interview_flag_patch_m232.py` fences the manifest shape + the wiring + a live-anchor drift pin).

**Zero platform-repo edits:** the patches touch the demo's ephemeral clone before the image build and revert
after (the demopatch 7-guard contract); the canonical `anthropos-work` repos are never touched.

## 6. Safety — the read-side exception, bounded

Sourcing anonymized-real sessions is a user-accepted (data-controller) softening of `safety.md`'s "nothing
behind the door", bounded three ways: **anonymize by construction** (PII never sourced), **source-pinned +
disclosed** (`content_sessions` manifest block), **VPN/tailnet-scoped** (the Part-3 exposure posture). **Part 2
(never-write-prod) is untouched** — the seeder writes only per-stack Postgres, audited, n=0-guarded. See
[`safety.md`](../safety.md) §3.8 (the amendment this milestone lands).

## See also
- [`content-stories-routes.md`](content-stories-routes.md) — the M231 spike (the read/route/sourcing contract this realizes).
- [`safety.md`](../safety.md) §3.8 — the anonymized-real read-side exception this milestone discloses.
- [`../db-access.md`](../db-access.md) — the read foundation + the public-vs-customer boundary the sourcing honors.
- [`seed-manifest-spec.md`](seed-manifest-spec.md) — the `seed-generation-manifest.yaml` the `content_sessions` pins fold into.
- [`stories-spec.md`](stories-spec.md) + [`../seeding-spec.md`](../seeding-spec.md) — the 7-table fan-out + the mirror-pair the substrate extends.
- [`demopatch-spec.md`](demopatch-spec.md) — the demo-patch mechanism the interview flag-gates use.
- The manifest / cockpit halves: **M233** (`content-stories-spec.md`, the manifest projection) + **M234** (the cockpit tab) + **M235** (prove-it-lands render proof).
