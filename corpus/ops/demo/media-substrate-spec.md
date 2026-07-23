# The media substrate — recorded interview VIDEO + document bodies in a content-story demo

**Status:** specified (v2.6 "sound check", M240) · **Render path:** live · **Seed-side exhibit:** specified,
**Bunny-key-blocked** as of 2026-07-21 · **Safety contract:** [`../safety.md` §3.8.1](../safety.md) (the raw-media
amendment + the 2026-07-21 data-controller VIDEO sign-off)

This is the media half of the "Content stories" feature. [`session-clone-spec.md`](session-clone-spec.md) copies a
played session's **free-text** (transcript, LLM feedback, submission, interview report) prod→scrub→fixture. This
spec covers the two **media** facets that make a session *playable* in a demo — the **recorded call** and the **full
document body** — and it is the honest record of what that costs, how it renders, and why the current demo ships
without a recording.

> **PII discipline (load-bearing).** Customer media — recorded **video**, audio, faces, document bodies — must
> **never** enter an agent's context. You ORCHESTRATE the tooling; you never view the media. The recording facet is
> designed so that **no media byte ever moves through the tool or the agent** (see §2). Handle the Bunny.net keys
> **values-blind** (§4). This spec was authored without a single recording byte, transcript, or face entering any
> agent's reasoning.

---

## 1. What the substrate IS

Two facets, two very different shapes.

### 1.1 The recorded call → a Bunny.net CDN reference (NOT an S3 blob)

A recorded session's media is **not** stored as a blob the demo copies. It is a **reference**:

- **`jobsimulation.sessions.chime_status`** (string) — the render gate. `'completed'` means "a recording exists and
  is playable"; `'not_available'` means "no recording" (the faithful default for every non-recorded session).
- **`ChimeRecording`** ent (`jobsimulation`) — one row per recorded session, carrying **`bunny_video_id`** (plus
  `recording_id`, `session_id`, `attendee_id`, `meeting_id`, `media_pipeline_id`). The `bunny_video_id` resolves to
  an MP4 in a **Bunny.net Stream CDN** pull-zone (`vz-…​.b-cdn.net`). **The bytes live on Bunny's CDN, never in the
  platform DB and never in prod S3** — which is why S3 read access (the old `DEF-M10-01`) is neither necessary nor
  sufficient to exhibit a recording.

**The recorded pool is almost entirely HIRING interview VIDEO of real candidates** (a face + a voice). Assessment /
training / interview-sim *voice* sessions carry **no** recording. This is decisive for content-story sourcing (§5).

### 1.2 The document body → inline text (Defect 3, RESOLVED — no blob)

A document (`collaborative_doc`) session's body lives inline in
**`validation_criterion_results.input_data.text_document`** — a text field, **not** an S3 `storage_upload` blob.
`jobsimulation` has no attachment/file table other than `collaborative_assets`, and the pinned document sessions
have zero rows there. So the document facet needs **no** blob port: it is copied + scrubbed exactly like the
transcript (M240 Defect 3 wrote it at seed time via the content-specific criterion-column set). It is **fully
landed** — this spec records it only to close the "is the document body a blob?" question: it is not.

---

## 2. The render path — bytes flow Bunny-CDN → demo-server → browser, at render time only

For a recorded session, a reviewing **manager** in the demo triggers playback and the bytes are fetched **live**;
nothing is pre-copied:

1. Browser → same-origin **`/api/bunny/recording/[sessionId]?userId=…`** (next-web-app `apps/web` + `apps/hiring`).
2. The route calls the GraphQL resolver **`jobSimulationChimeRecording(userId, sessionId)`**
   (`jobsimulation/internal/graph/queries.resolvers.go`), which runs **`CheckSessionReadPermission`** (the caller
   must be a manager of the org that ran the session) and returns the row's **`bunnyVideoId`**.
3. **`getBunnyRecordingDownloadUrl(bunnyVideoId)`** (`next-web-app packages/core-js/src/functions/bunny-thumbnail.server.ts`)
   signs a short-lived Bunny CDN URL using **`BUNNY_RECORDING_CDN_TOKEN_KEY`** + **`BUNNY_RECORDING_PULL_ZONE_HOST`**.
   If those env values are unset the function returns `null` → the route answers **HTTP 500 "Bunny signing not
   configured"**.
4. The **demo's own next-web-app server** `fetch`es the signed URL and **streams** the MP4 back to the browser
   (`Content-Disposition: attachment`; the signed URL stays server-side).

**The agent is never in this path**, and neither is the capture tool — the fetch happens in the running demo, at the
moment a human presses play. The demo holds only a **reference** (`bunny_video_id`) + a **read-only signing key**.

---

## 3. The seed-side exhibit mechanism (specified; the write is trivial, the bytes are not moved)

To exhibit a recording the `ContentStorySeeder` writes, into the **per-stack demo Postgres only**:

1. `jobsimulation.sessions.chime_status = 'completed'` on the re-tenanted content-story session, and
2. a `ChimeRecording` row for that session carrying the **real `bunny_video_id`** captured (read-only) from the
   pinned source session.

That is the whole seed-side port: **two rows referencing a recording, not the recording.** No MP4 byte is read,
copied, transcribed, or persisted. `content-capture` captures the `bunny_video_id` **string** (a UUID) alongside the
free-text it already captures — a reference is not the media class the scrub removes, and it never enters an agent's
context (the tool prints counts only).

**Gender coherence (believability, §3.8.1).** When a session carries a recording the owning demo persona's apparent
gender must match the **person on screen**. The source recording's apparent gender is **labeled at capture time,
values-blind** — derived in-tool from the session owner's sourced identity (the same name the scrub reads and
drops), emitted as a `m`/`f`/`unknown` label, **never by an agent viewing the video** — and the persona pairing is
constrained so the two align.

---

## 4. Provisioning the Bunny recording keys into a demo (values-blind, the M239 pattern)

Exhibiting a recording requires the demo's next-web-app to hold **`BUNNY_RECORDING_CDN_TOKEN_KEY`** +
**`BUNNY_RECORDING_PULL_ZONE_HOST`** (the recording pull-zone; the non-recording `BUNNY_CDN_TOKEN_KEY` /
`BUNNY_PULL_ZONE_HOST` pair drives thumbnails/scenario videos and does not sign recordings). These are **read-only
CDN-token keys** — they grant the demo server the ability to *sign a fetch* of an existing recording, never to
write, replace, or delete one.

They are provisioned into the demo the same way M239 provisioned the AWS Bedrock creds ([`../safety.md` §2.10](../safety.md)):
**values-blind** — the secret-coverage DNA / the demo env bridge carries the *key names* and copies the *bytes*
source→target; **no verb ever reads, echoes, logs, or commits a value**. A key never enters git, a fixture, or an
agent's context.

---

## 5. Sourcing constraint — only hiring-voice cells can carry a recording

The M240 Defect-1 sourcing predicate constrains each content-story cell's public sim to **`d.type = <cell
sim_type>`**. Because recorded voice sessions are almost exclusively **HIRING**, and assessment / training /
interview-sim voice sessions have **no** recording, only the **hiring-voice** content-story cells can ever source a
recorded session — every other cell is kept off the recorded pool by construction. A recorded hiring session must
also remain **public-anchored** (the M231 contract: its `sim_id` resolves in the demo's replayed catalog) and
**source-pinned + disclosed** in `seed-generation-manifest.yaml`.

---

## 6. Current status (M240, honest) — Bunny-key-blocked, faithful `not_available`

The render path is live and the seed-side exhibit is specified, but exhibiting a recording depends on the **Bunny
recording signing keys** being provisioned into the demo — and those key **values are absent from this authoring
box's entire dev-stack** as of 2026-07-21:

- no populated value in any real `.env` under `stack-dev/` (platform, next-web-app, studio-desk, app),
- none in the `.agentspace/secrets/` provisioning source (incl. `next-web-app/apps/web/.env`),
- none in either platform compose file or the shell env — only key-name **templates** (`.env.example` /
  `platform/.env_example`) exist.

Until an operator provisions the recording keys **from a real source** (prod / vault), a content-story demo ships
**without** an exhibited recording: `chime_status` stays **`not_available`** — exactly the pre-v2.6 state, degraded
honestly, **never a broken "play" button over a recording the demo cannot sign for**. Re-pinning a voice cell to a
recorded hiring session and flipping `chime_status='completed'` **without** the key would render a 500 error — a
regression against the honest current state — so the seed-side exhibit is deliberately **held as one atomic unit**
until the key arrives. The posture (safety §3.8.1) and this substrate spec land **ahead of** the capability, by the
gate-ordering rule: the moment the keys are provisioned, the exhibit is pre-blessed.

---

## See also
- [`../safety.md` §3.8.1](../safety.md) — the raw-media amendment + the 2026-07-21 data-controller **VIDEO**
  sign-off (the *why-it-is-safe* contract; §3.8 is the free-text exception this extends).
- [`session-clone-spec.md`](session-clone-spec.md) — the free-text copy + scrub + source-pin mechanism (the other
  half of Content stories).
- [`content-stories-routes.md`](content-stories-routes.md) — the M231 per-product result-route map + the sourcing
  contract this realizes.
- [`../safety.md` §2.10](../safety.md) — the M239 Bedrock-creds precedent for values-blind demo secret provisioning.
