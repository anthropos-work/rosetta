# M49 — spec notes

_Technical notes accumulate here during build (file:line surfaces, rext tag, schema findings)._

## Pre-flight audits — §1 rext-tag-sot (first section of session)
- KB-fidelity verdict: **YELLOW** (proceed-with-tracking). Report: `kb-fidelity-audit.md`. SHA at audit: `811b0840f`.
- No blind areas; all 7 topics PAIRED. Stale claims found = M49's own In-scope deliverables (tracked KB-1..KB-4 in `decisions.md`).
- Audit reuse: the load-bearing KB docs don't change between sections (each section touches a disjoint doc subsection); subsequent sections REUSE this verdict per the audit-reuse rule unless a new subsystem surfaces.

## Grounded file:line surfaces (rext authoring copy @ main)
- `demo-stack/up-injected.sh` = 760 lines. **#3** `.env` guard: `:258` (`[ -f "$PLAT/.env" ] || { ... exit 1; }`); provision block `:283-326` (repoints `BASE_ENV` on success). Move guard to AFTER `:326`.
- **#6 RAM check** `:90-115` (`docker info --format '{{.MemTotal}}'`, warn-only `min_gib`); mirror for disk. `build_frontends()` `:222`, called `:525`, compose up `:528-529`.
- `demo-stack/rosetta-demo` = 245 lines. **#6 cmd_down** `:139-162` (compose `down`/`down -v`; native ant-academy + http.server teardown). Add image `rmi` for `demo-N-next-web` + `demo-N-studio-desk`.
- **#4** `stack-secrets/secretdna/secret-dna.json` — `INVITATION_HMAC_SECRET` ABSENT (grep -c = 0). Add as critical gene + auto-gen throwaway at provision.
- **#5** `demo-stack/ensure-clones.sh:69-90` — stub-sweep (c-pre) iterates `repos.yml` entries only; `make init` (c) clones repos.yml repos. ant-academy absent from repos.yml ⇒ never cloned. (the real root cause; NOT the FA-token theory the old comments claimed.)
- **#8** `demo-stack/patches/next-web-studio-url/next-web-studio-url.yaml` — `pre_sha256: b3d62dbe…`, `post_sha256: 9f27e253…`, anchored to v1.10 release ref. `path: packages/core-js/src/constants/urls.ts`. Re-anchor to current `stack-demo/next-web-app` (v2.89.0) source.
- **#1** 3 conflicting pins: `SKILL.md:84` + `frontend-tier.md:254` → `storytelling-postfix-2`; `rosetta_demo.md:15` → `storytelling-postfix-1`. Current consumption tag = `fit-up-m47`. New SoT: `.agentspace/rext.tag`.

## Open decision — #5 repos.yml vs explicit clone (resolve in §4)
`stack-demo/platform` is an EPHEMERAL gitignored clone (ensure-clones.sh bootstrap-clones it fresh from GitHub). Editing its `repos.yml` is non-durable (overwritten on re-clone) AND a platform-repo edit (forbidden). Durable rext-owned fix lives in `ensure-clones.sh`. To confirm in §4: does the freshly-cloned `stack-demo/platform/repos.yml` already carry ant-academy (then the bug is elsewhere), or not (then ensure-clones.sh must add the clone)?
