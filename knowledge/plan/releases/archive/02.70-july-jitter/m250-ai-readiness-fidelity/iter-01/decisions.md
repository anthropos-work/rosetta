# iter-01 — decisions (bootstrap tok)

## D1 — capture the 31 default node-ids verbatim from platform `defaults.go`
The exact repertoire (never fabricated — closure-gate-safe; all confirmed in the captured `public.skills`):
- **Core (19 @ weight 1.0):** K-AGECOD-C896, K-AGEENG-13FC, K-AICOSM-B710, K-AIEVAQ-B95C, K-AISECA-14AF,
  K-AISTRX-CEA7, K-AISYSD-26C1, K-AIUSEA-7F89, K-AIWORD-00FE, K-ANTCLI-6DDA, K-CLACOW-D576, K-CLADES-AAE6,
  K-CODEXX-77C3, K-CONCRE-AC50, K-CUSAGE-CB3D, K-GENAPP-FBE8, K-LLMMLO-0241, K-RESAIG-DBA4, K-TASAUT-EC86
- **Enabling (12 @ weight 0.5):** K-ADARES-2B17, K-ANATHI-9B6A, K-APIINT-57F2, K-BUSACU-3FF6, K-CLOCOM-CB7E,
  K-COSMAN-35F5, K-CRITHI-224F, K-DIGDAT-553C, K-INFEVA-A507, K-PYTHON-8B21, K-SQLXXX-2064, K-WRICOM-824E
- totalWeight = 19×1.0 + 12×0.5 = **25.0** (was 6.5). This replaces the name-pattern `readAIReadinessSkillPool`
  approach (which surfaced whatever AI-named skills happened to be in the taxonomy — the "invented 8").

## D2 — the Step-2 "track" label is a NAME HEURISTIC, not the `track` column
`how_we_measure.go computeSimAssessments` classifies a sim card tech/business via `techTrackRe` over the
evaluated skill NAMES parsed from `directus.simulations.skills`. So the label the page renders is controlled by
the skills we set-dress into `simulations.skills`, not by `ai_readiness_sims.track`. The `track` column still
matters for `resolveUserTrack` member routing (participants_filter tags). Confirm both at live render.

## D3 — Directus set-dress is net-new (no replay seam)
`directus.simulations.skills` is NULL for every captured sim; the real evaluated skills live at
`directus.sequences.evaluation_skills` (node-ids). Snapshot replay is replay-only. A new post-replay set-dress
step must resolve those node-ids → `public.skills.name` and UPDATE `directus.simulations.skills` for the 3 wired
sim uuids. No seeder writes directus content tables today.
