# M28 — Decisions

_Implementation decisions with rationale, numbered `M28-D1`, `M28-D2`, … . Empty at scaffold; filled during build._

## M28-D1 — LiveKit key/secret are NOT an alias family (a credential pair ≠ one value)
**Surfaced building the alias-mapping path.** The M27 DNA's `alias` field means "genes sharing ONE underlying
value" (per the `dna.go` doc-comment + `gh-token`: GH_PAT≡GH_ACCESS_TOKEN≡GH_TOKEN). M27 ALSO put
`LIVEKIT_API_KEY` + `LIVEKIT_API_SECRET` in a `livekit` alias family — but a key+secret credential pair holds
TWO DISTINCT values, not one. provision's alias-fallback ("if a member's own key is absent from the source,
copy a sibling's value") would then write the API-key's value under `LIVEKIT_API_SECRET` if the secret were
missing — a silent wrong-value provision.
**Fix (Fate-1, in code M28 builds on + directly affecting M28 provision correctness):** removed the `livekit`
alias from both genes — each is now a standalone gene sourced by its OWN key. Updated the note + the
`secret_dna_json_test.go` assertion (livekit must NOT be an alias family; each carries no alias). The
`gh-token` family (the one true same-value family) is unchanged + still satisfies ValidateAliases (≥2) + the
anti-vacuous-100 guard. The alias mechanism is now reserved strictly for same-value families, matching its
documented contract.

## M28-D2 — provision writes append-only (copy-if-absent never rewrites existing lines)
**The values-blind-write design.** provision necessarily moves secret VALUES source→target, but to keep that
move auditable + non-destructive it APPENDS a provisioned block rather than rewriting the file. An existing
target line is never re-read for its value or rewritten — `readTargetKeys` is values-blind (NAMES only), and
copy-if-absent skips any key already present. The only value-carrying read is `sourceValues` (in `provision/io.go`,
the documented value-carrying boundary); its values flow only into the `WriteString` to the gitignored target,
never to a log/error/return. `--force` re-appends (a later duplicate KEY= line wins in `.env` precedence, so an
overwrite is honest); the N=0 guard mirrors `stackseed --reset` exactly.

## M28-D3 — DIRECTUS_TOKEN non-rearm: write BLANK on a non-prod target (defer to the override)
**The blocks-release safety class.** The demo/dev injection override strips the prod `DIRECTUS_TOKEN` to `""`
at compose-emit time (fix16/fix17). provision runs BEFORE the override + writes the base `.env`. To compose
safely, provision NEVER copies the source `DIRECTUS_TOKEN` (or `DIRECTUS_STATIC_TOKEN`/`DIRECTUS_ADMIN_TOKEN` —
the `StripOnNonProdKeys` set, mirroring `PreflightEnv`'s Directus-write-token rejection in `safety.md §2.2`)
into a non-prod target. It writes them BLANK (`KEY=`) — exactly the state the override forces — so the base
`.env` + the override agree + the prod-write path is never re-armed. The prod path (N=0 + real source) is
guarded by the N=0 refusal, so it is never auto-touched either. The DNA gene `platform/DIRECTUS_TOKEN` is
already `key-present`-only (no `nonempty`), so a blank value still PASSES coverage — the safety + the score
agree by construction. Headline regression test: `provision_safety_test.go::TestProvision_NeverReArmsDirectusTokenOnNonProd`.
