# Phase B adjudication — every reported blocker re-verified before acceptance

§5 requires verifying a claim before escalating it, INCLUDING a claim made by an audit (iter-22: 2 of 21
handed corrections were themselves false). Each blocker below was re-derived from platform source by the
iteration itself, not accepted on the auditor's word.

| id | re-verification command / evidence | verdict |
|---|---|---|
| B-C1 | `repos.yml:17` + `docker-compose.yml:83` both contain jobsimulation; roadrunner.md:23-25 says both removed | **HOLDS** |
| B-C2 | `assignments.go:815` = `}))`; `:828` = `h.cms.GetSkillPath(` | **HOLDS** (anchor off, claim true) |
| B-A1 | `config_template.ini:39,40` = `EXECUTION_/CREATIVE_AI_STABLE_MODEL = azure, gpt-4o, none` | **HOLDS** |
| B-A2 | `grep -c 5050 docker-compose.yml` = 0 (rc=1); doc :79 asserts "host 5050 -> container 8080" present-tense | **HOLDS** |
| B-A3 | `jobsimulation.go:905` `AIVendor *AIVendor` (nullable) + `:1302` `aiVendor := simulation.Openai` | **HOLDS** |
| B-A4 | `external_services.md:447` = table header row; the cited correction is at `:512` | **HOLDS** |
| B-F1 | `git show a2a3ee6^:docker-compose.yml` -> `image: directus/directus:10.10.1`, `8055:8055`, `ADMIN_PASSWORD=password` | **HOLDS** (false retraction) |
| B-F2 | `grep -cE '"(react\|vue\|@angular)' studio-desk/package.json` = 0 (rc=1) | **HOLDS** |
| B-F3 | `docker-compose.yml:311` `studio-desk:` | **HOLDS** |
| B-F4 | `app/main.go:604` = `jobsimwiring.Wire(...)` only | **HOLDS** |

**10 of 10 accepted.** Unlike iter-22, no handed correction was refuted on re-derivation — a property of
the briefing (every auditor was required to cite the exact command) rather than of luck.
