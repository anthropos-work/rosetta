# iter-96 blocker ledger — `FIX-M257x-iter95-read-union`, repaired BY PREDICATE

This is the repair ledger for the 13 anchors / 12 predicates iter-95's reading returned. It is written
in the **ledger table shape** `rosetta-extensions/stack-core/claim_ledger.py` derives from — a table
under a non-Minor heading with a claim-shaped column and an anchor-shaped column — so every refuted form
below is adopted **automatically** by `claim_twin_guard` and `repair_postcondition` and becomes
un-republishable tree-wide. That is the point of writing it here rather than in prose: iter-95's binding
condition was *repair by predicate, not by anchor*, and a fence is the only thing that holds a predicate
down after the repairing agent is gone.

## BLOCKERS — the refuted claims

| # | the false claim | anchor | what is true | sites repaired |
|---|---|---|---|---|
| P1 | "From `docker-compose.yml` *before the drop*, the gateway `depends_on`: - backend - storage" | `external_services.md:411-413` | The router's last **mainline** `depends_on` was **four** services — backend, jobsimulation, cms, storage — at `2adcf71^1` (`1e8e754`), `docker-compose.yml:19-27`. The two-service pair exists only at `464dfe3`, on the **unmerged** branch `origin/feat/cms-in-app`; `git merge-base --is-ancestor 464dfe3 0c91421d` exits 1. A `git log --all` sweep reaches it and it reads like the newest pre-drop state. | 1 |
| P2 | "`mistralai` is declared in `requirements.txt` and imported nowhere — that declaration is the string's *only* occurrence in the whole repo" | `cms.md:95` | `git -C app/studio grep -i mistral aeec036a` returns **22 hits across 3 files**; `tools/pdf2md.py:24` is a live `from mistralai import Mistral` (`mistral-ocr-latest` at `:127`). It is a standalone CLI off the generation pipeline, not dead code. Mistral is **OCR-only on both sides**, Go and Python. | **11** |
| P3 | "In prod terraform the address is `http://backend:8081`" | `backend.md:112` | One endpoint, not two. `app` deploys as `local.project = "backend"` in the Cloud Map namespace `internal.anthropos` on `local.rpc_port = 8081`, so the production address is `http://backend.internal.anthropos:8081`. The unqualified form is the same endpoint written short. The literal is set in the **un-cloned** `infrastructure` root module — every cloned service module declares its `*_rpc_addr` with no default. | 3 |
| P4 | "repo ARCHIVED 2026-07-31" (jobsimulation) | `service_taxonomy.md:130` | Unmeasurable from a clone, and contradicted: `origin/main` carries four commits dated **2026-08-04**, including `caf36c96` (merged PR #439, committer `GitHub`), while an archived GitHub repo is read-only. Archive state lives in the GitHub org API. Report both, assert neither. | 6 |
| P5 | "ECS module kept as the rollback path" (cms) | `service_taxonomy.md:131` | Two measured facts point opposite ways: `cms/terraform/main.tf:39` still declares the module at `service_desired_count = 0`, and `6efa1d5` (merged `f38c0c4`) deleted the build-production workflow under *"the cms ECR repository is decommissioned (M810)"*. The settling declaration is in `infrastructure`, which has never been in any clone set. The map's own rule is report both, assert neither. | 2 |
| P6 | "custody transfer is M903 (`:18`)" | `platform-migration-status.md:92` | `storage/terraform/main.tf:18` is the module's closing comment about `outputs.tf`. *custody* occurs **0** times in the storage repo at `9f8cb53`. M903's only mention is `storage/terraform/storage.tf:22-25`, which says it **was never executed and the shipped design supersedes it**; `d3e6d32` states *"M903 never ran."* No `moved` block exists. | 2 |
| P8 | "In the platform compose, `STORAGE_S3_PUBLIC_BUCKET` is hardcoded to the production public bucket, so locally the PUBLIC manager talks to real S3 ... while the PRIVATE manager uses local FS" | `storage.md:58` | **Both** buckets are hardcoded to production buckets on the `backend` block — `docker-compose.yml:82` and `:83` @ `0c91421`. `getKeyPath` routes to `s3://` for any non-empty bucket. `backend` mounts `$HOME/.aws/credentials` and platform `README.md:81-87` tells you to supply live keys, so a stock stack writes private uploads into production. Both app boot guards are disarmed by `ENVIRONMENT=development`. | **10** |
| P9 | "**Every** live Go service: app, sentinel, storage, messenger — and **only** those four" | `shared_libraries.md:42` | At platform `0c91421` `repos.yml` has four entries of which **two are Go** — `app` and `sentinel`. `838d907` deleted the `storage` and `messenger` clone entries and their compose services. The four-repo reading was true at `0dab54d` and asserts currency without a pin. | 6 |
| P10 | "Skill score changes — both producer and consumer live inside app since the skiller→app merge" | `dependency_map.md:59` | There is **no producer**. Every publisher constructor in `app` at `b948604f` and `2035f9a4` names `backend`, `SKILLPATH_STREAM`, `CMS_STREAM`, `AI_USAGE_STREAM`, `JOBSIMULATION_STREAM` — never `SKILLER_STREAM`, whose one Go occurrence is an `AddSubscriber`. The payload is `SkillerCustomJobRoleCreated`, not skill scores. The producer went with the standalone service. The `skillpath` twin is genuinely two-ended and is **not** the same case. | 3 |
| P11 | "its compose `environment:` block carries no `DIRECTUS_*`" | `external_services.md:248` | It carries exactly one — `DIRECTUS_PUBLIC_BASE_ADDR` at `docker-compose.yml:53` @ `0c91421`, inside `backend`'s block (`:46-94`). Compose `environment:` overrides `env_file:`, so re-pointing the public address in `.env` alone is a no-op. | 2 |
| P12 | "**Not** the 60K-skill dataset, which lives in `app`'s `public` schema" | `platform-migration-status.md:110` | The measured public floors are **≥42,790 skills / ≥22,470 job roles** (`organization_id IS NULL`, 2026-06-29). The roles figure is REFUTED, the skills figure UNVERIFIED. The index of truth restated a figure the corpus elsewhere fences. | 1 |
| P13 | "The framework is **engine-agnostic and reusable** — it lives in rosetta and knows nothing about Clerk" | `alignment_testing.md:25` | The harness is the `alignment/` section of **`rosetta-extensions`** (Go module `anthropos.dev/alignment`). rosetta ships this doc and the two skills and no executable alignment code — as the same document's own *Where things live* section says 480 lines later. | 1 |
| C1 | "positive control: `-S SKILLER_RPC` returns 3" | `messenger.md:22` | Reproducible at **no** repo, ref, spelling or scope. `git -C stack-demo/platform log -S 'SKILLER_RPC' --oneline 0c91421d \| wc -l` returns **7**; `--all` returns 8, the 8th being `464dfe3` on a non-main branch. The guarded claim (`MESSENGER_RPC` → 0) holds at every ref and scope, nested repos included. | 2 |
| C2 | "`STORAGE_RPC_ADDR` is read by **nothing**: `git grep -n STORAGE_RPC_ADDR -- '*.go'` returns **3 hits, all of them comments**" | `storage.md:29` | True at `app` origin/main `2035f9a`, and the command as printed carries **no ref** — run on the older checkout a demo pins it returns 15 hits, 7 of them live env reads, contradicting the sentence containing it. Found by this run's own absence-class sweep, not by the reading. | 1 |

**P7 is P6's second propagation site** (`storage.md:25`) — one predicate, two sites, and iter-95 recorded
that one seat **positively cleared** it at `platform-migration-status.md:92` while another booked it at
`storage.md:25`. Both are repaired.

## What "by predicate" bought, measured

| | count |
|---|---|
| anchors iter-95 booked | **13** |
| distinct predicates | **12** (+2 instrument defects, C1/C2) |
| **sites actually repaired** | **51** |
| sites an anchor-wise repair would have left standing | **38** |

Adjudicators had named **≥8** unbooked twins. The predicate sweep found **38** — over four times the estimate,
and the largest single predicate (P2, `mistralai`) had **11** sites of which exactly **one** was booked.

## Fences this repair grew, rather than prose it rewrote

- `unreadable_repo_claim_guard` reach **5 → 7** `module.*_euwest1` mentions: the P5 repair spells the
  dotted construct, so two sites that were paraphrasing *through* the unmeasurability boundary are now
  inside the fence rather than invisible to it.
- `anchor_construct_guard` is now **nested-repo aware** (`_clone_of` descends to the innermost git
  checkout). Before this, every citation into `app/studio/**` was graded at `app`'s ref, where the path
  does not exist at any sha — the guard reported UNMEASURED rather than resolving. That is the same
  mechanism as P2's false clearance, in the instrument instead of the corpus.
- This ledger itself: `claim_ledger.py` derives its claim set from ledger-shaped tables, so the 14
  refuted forms above are now adopted by `claim_twin_guard` tree-wide and by `repair_postcondition` at
  the commit.
