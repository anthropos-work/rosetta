# M29 — Decisions

_Implementation decisions with rationale, numbered `M29-D1`, `M29-D2`, … . Empty at scaffold; filled during build._

## M29-D1 — `/stack-secrets` builds from the pinned tag `stage-door-m28`, not the clone HEAD
The `/stack-seed` + `/stack-snapshot` skills `go build` from their per-stack clone's working tree (HEAD).
`/stack-secrets` instead checks out + builds the **pinned tag `stage-door-m28`** (per the milestone spec): it is
the latest tag and the first that carries the `provision` engine + the demo-aware `check` (the M28 deliverables
the skill drives). Pinning the tag makes the skill reproducible against the exact code the docs describe. Verified
the binary compiles from that tree (ext head `9742126`) and `list` outputs the catalog values-blind.

## M29-D2 — The skill's `--check|--provision|--status` are skill-level shorthand mapped to CLI subcommands
The argument-hint advertises operator-facing verbs (`--check`, `--provision`, `--status`); the skill body maps
each to the real `stacksecrets` subcommand (`check`/`measure`, `provision`, `list`). This mirrors `/stack-seed`'s
`--preset NAME` shorthand (resolved to `--seed presets/NAME.seed.yaml`). The binary's actual flags
(`--dna`/`--from`/`--stack-root`/`--stack`/`--force`/`--prod`/`--dry-run`/`--demo`) are all shown in the body's
example invocations, so an LLM-synthesized call uses the real parser flags, not the shorthand.

## M29-D3 — The README-index guard checks the **same-directory** README; index in `corpus/ops/README.md`
`secrets-spec.md` lives in `corpus/ops/`, which has its **own** `README.md` (the guard checks the same-dir
README, not the corpus root). Initial pass indexed only `corpus/README.md` and the guard caught the miss
(exit 1). Fixed by adding the row to `corpus/ops/README.md` (the guard's actual target) — and kept the
`corpus/README.md` row too (both are valid front doors). Guard now exit 0.

## M29-D4 — setup_guide.md keeps the per-repo key lists, retires only the `cp` mechanics
Retiring the manual-copy prose meant deleting the `cp .env.example .env` hand-copy steps + the line-447 TODO and
pointing to `/stack-secrets`. But the per-repo **key lists** (what each key is + where it comes from) are still
useful reference, so they stay; only the *copying mechanism* is automated. The root `platform/.env_example → .env`
copy also stays — that's where the operator's source secrets originate (the skill distributes *from* there).
Corrected ant-academy's target to `code/.env.local` (the verified live truth; `.env` absent) while there.

## Surfaced-and-confirmed (three-fate rule)
- **M30 field-bake (build-from-stack-dev validation)** is the observable-behavior gate that proves a compliant
  `.agentspace/secrets` provisions cleanly with `Critical == 100%`. It is **out of M29 scope** and **already owned
  by M30** (the next, final milestone of this release) — Fate 2 (already planned), no new tracking needed. M29
  delivers the docs + skill the field-bake will exercise.
- No items required Fate 3 or the escape hatch. Zero ext code needed (M29 is rosetta-only); the ext stayed on
  `main` @ `9742126`, untouched.

## Adversarial review (close Phase 2c — scenarios considered)
M29 ships no executable rosetta code; the adversarial frame is "how could a doc/skill consumer be led into a
wrong or unsafe action?" Each scenario was checked against the ext code at tag `stage-door-m28` (`9742126`).
- **LLM synthesizes a non-existent CLI flag from the skill body** (M29-D2 risk). Verified every invocation in
  both `SKILL.md` and `secrets-spec.md` against the real parser (`cmd/stacksecrets/main.go`): subcommands
  `list`/`check`(`measure`)/`introspect`/`diff`/`provision`; `provision` flags `--dna --from --stack-root
  --stack --force --prod --dry-run`; `check` `--dna --from --demo`; `introspect`/`diff` `--stack-root`. The
  operator-facing `--check|--provision|--status` shorthand is mapped to real subcommands in the body (`--status`
  → `list`). No doc invocation uses a flag the binary lacks. Handled — no fix.
- **A consumer reads a secret value because the docs imply it.** The SKILL.md states the values-blind invariant
  three times incl. "**The skill itself NEVER prints a secret value** — do not cat/echo a `.env`". `ClassifyShape`
  (the single value-touching fn, `secretdna/source.go:57`) returns a shape token; the value boundary
  `provision/io.go::sourceValues`→`writeTargetFile` and the `provision_safety_test.go` no-escape test both exist
  as the doc claims. Handled — no fix.
- **A consumer provisions the prod Directus token onto a non-prod stack.** `StripOnNonProdKeys` =
  `{DIRECTUS_TOKEN, DIRECTUS_STATIC_TOKEN, DIRECTUS_ADMIN_TOKEN}` (verified) matches secrets-spec.md §"non-rearm"
  and safety.md §2.9 exactly; both docs consistently steer to the blank-on-non-prod behavior. Handled — no fix.
- **Stale counts.** DNA ground-truth re-verified: 55 genes / 6 repos (28·5·2·7·7·6) / 40 required · 8 optional
  · 7 waived / 13 critical-required / profile `graphql`; `gh-token` alias family = exactly 3 members; `MintedKeys`
  = the 6 Clerk keys. Every count in the docs matches. Handled — no fix.
