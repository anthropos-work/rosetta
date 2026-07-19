# M232 — Decisions

## D1 (S1) — Anonymize BY CONSTRUCTION (the core safety decision)
The seeder does NOT copy customer free-text — it copies only the NON-PII structural skeleton (source
session-id pin, real public sim_id, sim_type, modality, score, pass/fail, duration, actor/interaction
counts) and SYNTHESIZES every free-text facet (LLM feedback, transcript, submission, actor names,
interview report). Per `content-stories-routes.md` §3.3, M232 owns the synthesize-vs-redact choice and
the brief leans synthesize. Consequence: the checked-in fixture is **provably PII-free**, the strongest
posture, and the `safety.md` Part 3 amendment (S4) records a NARROW, honest exception — real session
STRUCTURES sourced (scores/shape), synthesized content, source-pinned, VPN/tailnet-scoped.

## D2 (S1) — Owner = an existing player-vantage population member (REUSE, not mint)
Open question ("reuse hero seats or mint per-session anonymized player seats — brief leans mint, each
must map to a real seeded `public.users` row") RESOLVED toward **reuse a distinct non-hero player member**
of the target org. Rationale: reuse trivially satisfies "maps to a real seeded users row" + "never a
manager seat" (owner-is-player BY CONSTRUCTION), needs zero extra user-minting (no re-implementing the
UsersSeeder identity/casbin path), and a real member owning a real content-story session is MORE
believable than a thin single-session minted seat. Non-hero slots keep hero dashboards clean.
- Surfaced for M234: if the cockpit's "become this session's player" UX needs a dedicated seat per
  session, mint then — but the reuse baseline already renders. (Handoff noted; not a deferral of M232 scope.)

## D3 (S1) — Target org = the first NON-HIRING EffectiveStory
Content-story sessions land in the first non-hiring story org (has player members). They appear in that
org's activity (believable — a few extra sessions of extreme score don't skew a 220-member org). A
dedicated isolated content-story org is an M234 open question (cleaner cockpit grouping), not needed for
M232 to render.

## D4 (S1) — Fixture is a go:embed'd curated artifact (not per-stack config)
The exhibit set is a FIXED, source-pinned, code-owned artifact (like the AI-readiness prompts), so it is
go:embed'd in the `contentsession` package rather than declared per-stack in the blueprint. The seeder
loads `Embedded()`; M233 projects the pins from the same embed into `seed-generation-manifest.yaml`
(single source, honesty-gateable).

## D5 (S1) — Sourcing is authoring-time only; the seeder never touches prod
`sourcing.go` BUILDS the selection SQL (the reproducible record of the public-anchoring + non-manager +
per-cell selection, D6) but never RUNS it. Sourcing is an authoring-time activity against the read-only
postgres MCP (`db-access.md`); the seeder is fully offline (reads only the embed + the stack's own
replayed taxonomy). This keeps a demo box prod-access-free and the read boundary honored.
