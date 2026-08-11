# iter-97 ground truth — re-derived at this open, 2026-08-06

**This sheet is authoritative for refs.** A seat that grades a claim against a ref not on this sheet
must say which ref it used and why. Delivered to all 14 seats as a marked addendum below a
**byte-identical** copy of the frozen briefing (prefix sha re-verified after copying:
`head -171 | shasum -a 256` = `3858ec53…6eb0`, equal to the instrument's own sha).

| clone | checkout | `origin/main` | behind |
|---|---|---|---|
| `platform` | `0c91421d` | `0c91421d` | 0 |
| `app` | `b948604f` | `2035f9a4` | **93** ⚠️ |
| **`app/studio`** (nested repo) | **`aeec036a`** | — | — |
| **`cms/studio`** (the SAME repo, second copy) | **`aeec036`** | — | — |
| `next-web-app` | `bb3313bc` | `8297c684` | **41** ⚠️ |
| `storage` | `4ce8ece5` | `9f8cb532` | **20** ⚠️ |
| `messenger` | `fa47850d` | `e9421c68` | **7** ⚠️ |
| `ant-academy` | `9c3843cd` | `22df69dd` | **5** ⚠️ |
| `jobsimulation` | `462343b0` | `82cb66ec` | **4** ⚠️ |
| `cms` | `ca50c817` | `f38c0c4a` | **2** ⚠️ |
| `sentinel` | `88bc5592` | `f2c46190` | **2** ⚠️ |
| `studio-desk` | `14a5442a` | `41ee3575` | **2** ⚠️ |
| `roadrunner` | `87d8d443` | `87d8d443` | 0 |
| `graphql-wundergraph` | `60c229f3` | `60c229f3` | 0 |
| `rosetta-extensions` (authoring, `.agentspace/`) | `bc2ee74` | — | 0 |
| `rosetta` (the corpus under audit) | **`00be1ac`** | — | tree clean |

## What changed since iter-95's sheet

**No clone moved.** `platform` is still `0c91421d` and still equal to `origin/main` and to
`ls-remote origin HEAD`. Every other checkout is byte-identical to iter-95's. **The corpus moved**
(`1977224`/`b7e6642` → `00be1ac`, the iter-96 repair: 23 files, +230 −79) and **rext moved**
(`6130bfd8` → `bc2ee74`, the nested-repo fix to `anchor_construct_guard`).

**Two rows are net-new and they are the reason iter-96 exists:** `app/studio` and `cms/studio` are the
same `anthropos-studio-room` repo, checked out **inside** their hosts, untracked there and hidden by
`app/.gitignore:79` / `cms/.gitignore:129`. `git -C app grep <anything> HEAD -- studio/` returns **0 for
every predicate**. They were absent from every prior sheet, which is how a false clearance stood.

## The three rules that go with the sheet

1. **Grade at the ref the claim names — UNLESS the sentence asserts currency**, in which case no
   neighbouring pin rescues it. A pin is a **date**, not an excuse.
2. **A corpus citation into `rosetta-extensions` grades against the AUTHORING copy** (`bc2ee74`)
   unless the citing block pins a ref.
3. **An absence is established only by `git grep` at a named ref, PER TREE, with nested repos
   enumerated.** A tree-wide zero that does not name its sub-repos is an unproven zero.

## Precondition derivations (each run at this open, not inherited)

| precondition | command | result |
|---|---|---|
| platform == origin HEAD | `git -C stack-demo/platform fetch origin` + `ls-remote origin HEAD` | `0c91421dfdb08dc75f17f1aabfb61394070e770b` — equal to checkout **and** `origin/main` |
| guard family | `guard_family.py --repo-root … --platform … --allow-not-run --verify-remote` | **14 GREEN · 0 RED · 3 input-gated** over 17 members, before and after the repair |
| READ instrument frozen | `shasum -a 256` + `git log --follow` | `3858ec53…6eb0`; **exactly one commit ever** (`012edd2`, iter-76) |
| rext Python suite | `/usr/bin/python3 -m pytest tests/ -q` (pytest installed this run; it had never run on this box) | **909 pass / 1 fail** — the failure **PRE-EXISTING**, reproduced with this run's changes reverted |
