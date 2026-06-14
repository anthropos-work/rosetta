# M29 — Retro

_Docs + the `/stack-secrets` skill + corpus wiring. The 3rd milestone of v1.6 "stage door". Rosetta-only
(zero ext code). Closed 2026-06-14._

## Summary

M29 made the M27/M28 secret-provisioning engine **discoverable** and gave it a **corpus home**. It authored
`corpus/ops/secrets-spec.md` (net-new, 290 lines — the secret-provisioning source-of-truth, closing the
Phase-0b KB blind area that the M29 design `delivers:` line was created to fill), added the `/stack-secrets`
skill (mirroring `/stack-seed`: read spec → confirm non-prod target → build the **pinned-tag** `stage-door-m28`
binary → run the verb → report values-blind), wired the skill + spec into `CLAUDE.md` (skill-table row +
Key-Documentation-Locations entry + Interconnected-Documentation rows 10/11) + both corpus indexes
(`corpus/README.md`, `corpus/ops/README.md`), extended `safety.md` with **§2.9** (the values-blind /
`PreflightEnv`-emitting / `DIRECTUS_TOKEN`-non-rearm clause), and retired the manual `.env` hand-copy in
`setup_guide.md` — the `/stack-secrets` fast-path callout + retired hand-copy steps for
studio-desk/ant-academy/next-web, with the **line-447 TODO finally deleted**. Per M29-D4 the per-repo key
lists were kept as reference (only the *copying mechanism* was automated), and the root
`platform/.env_example → .env` copy was kept (that's where the operator's source secrets originate).
The ext stayed on `main` @ `9742126` (= tag `stage-door-m28`), untouched — no branch, tag, or commit ext-side.

The close was clean: **0 findings** across scope/code-quality/docs/tests/decision-triage, and the deferral
audit was GREEN.

## Incidents This Cycle

None. No P0/P1/P2 incidents, no flakes, no regressions. M29 ships no executable code; the only programmatic
gate is the corpus README-index guard, which passed exit 0 on first run (M29-D3 records the build-time miss the
guard caught — `secrets-spec.md` was initially indexed only in `corpus/README.md`, not the same-dir
`corpus/ops/README.md` the guard actually checks — but that was caught + fixed *during build*, not at close).

## What Went Well

- **The doc/skill patterns to mirror already existed.** `seeding-spec.md` + `/stack-seed` and
  `snapshot-spec.md` + `/stack-snapshot` gave a proven template; `secrets-spec.md` slotted into the same
  "read-side family" framing (one-sided `datadna`-mold harness + an engine on top) and the `/stack-secrets`
  skill mirrored `/stack-seed`'s read-spec → confirm-non-prod → build-tagged-binary → run → report shape.
  Authoring against a *finished* engine (M28) meant zero design churn — the docs describe settled behavior.
- **Fidelity-against-code held at every claim.** The close re-verified every load-bearing number against the
  ext engine at `stage-door-m28`: the 55-gene / 6-repo DNA (40 required · 8 optional · 7 waived · 13 critical),
  the `gh-token` 3-member alias family, `StripOnNonProdKeys` (3 keys), `MintedKeys` (6 keys), `ClassifyShape`,
  the `provision/io.go` value boundary + `provision_safety_test.go`, and every `stacksecrets` CLI
  flag/subcommand/exit-code. All matched — the build's harden pass had already done this work, so close just
  confirmed 0 drift.
- **The M29-D2 risk (LLM synthesizes a fake CLI flag) was designed out.** The skill advertises operator-facing
  `--check|--provision|--status` shorthand but maps each to the *real* `stacksecrets` subcommand in the body,
  and every example invocation uses the actual parser flags — so an LLM-synthesized call uses real flags, never
  the shorthand. Verified against `cmd/stacksecrets/main.go`.

## What Didn't

- Nothing notable. A docs+skill milestone against a finished engine is inherently low-risk; the one build-time
  hiccup (the README-index same-dir target, M29-D3) was caught by the guard exactly as designed.

## Carried Forward

- **M30 field-bake (build-from-stack-dev validation)** — the observable-behavior gate that proves a compliant
  `.agentspace/secrets` provisions cleanly with `Critical == 100%` and the stack reaches UP. Fate-2,
  **already owned by M30** (the next + FINAL v1.6 milestone). M29 delivers the docs + skill the bake exercises.
  No new tracking.
- Inherited release-level backlog (DEF-M10-01 / DEF-M21-01 / M25-D9) — orthogonal to secret provisioning;
  re-signed at the v1.5 close, carried unchanged.

## Metrics Delta

(from `metrics.json`)
- **Findings:** 0 (0 scope · 0 code-quality · 0 docs · 0 tests · 0 decision-blend).
- **Go tests:** 1027 → **1027** (+0 — M29 touches no code).
- **Python tests:** 459 → **459** (+0).
- **Flake count:** 0.
- **Ext code:** none (rosetta-only; ext untouched on `main` @ `9742126` = tag `stage-door-m28`).
- **Deliverables:** 2 net-new docs (`secrets-spec.md` 290 lines + `SKILL.md` 135 lines) + 5 corpus/CLAUDE edits;
  README-index guard exit 0.
