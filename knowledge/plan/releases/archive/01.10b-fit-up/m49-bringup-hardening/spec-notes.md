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

## Live-verify gate (orchestrator, before close) — PASSED 2026-06-30

A from-cold rebuild on the consumption clone re-pinned to `fit-up-m49` proved all 7 fixes end-to-end
(log: `.agentspace/scratch/work-m49/cold-rebuild.log`, 1475 lines):
- Teardown (`rosetta-demo down 1 --purge`): **#6 image-cleanup fired** (demo-1 images removed, ~5 GB reclaimed).
- Cold `up-injected.sh 1`: **#3** provisioned-before-guard (26 written, no abort); **#4** "critical keys 100% present"
  + `backend /api/health 200` (the backend started — INVITATION_HMAC_SECRET auto-gen works); **#6** disk pre-flight
  "237 GiB ≥ 20 GiB OK"; **#1** ensure-clones pin-guard non-blocking; **#5** ant-academy started on :13077; **#7**
  2 frontends built + joined (no abort); **#8** `demopatch apply: next-web-studio-url applied` (the drifted patch
  re-anchored to next-web v2.89.0 applies cleanly).
- Completion: set-dress (snapshot replay) + seed (2 orgs × 3-hero, 57,858 rows, isolation clean) + cockpit (:17700)
  + **autoverify "verified-working" → demo-1 UP**.
- No from-cold bugs surfaced. (The provision summary's "MISS INVITATION_HMAC_SECRET" was the secret-SOURCE coverage
  report — expected; the #4 auto-gen writes it to the demo base env separately, confirmed present + nonempty.)
