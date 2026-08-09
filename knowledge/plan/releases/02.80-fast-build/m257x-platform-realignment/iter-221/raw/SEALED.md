# iter-221 — sealed before landing

## V1 — the population

`fence_provenance.corpus_sources` = **114** documents carrying **2,117** backticked file citations:
`.md` 876 · `.go` 428 · `.sh` 173 · `.py` 149 · `.yml` 133 · `.ts` 112 · `.json` 100 · `.yaml` 75 ·
`.tf` 39 · `.tsx` 32. **Nothing checks any of them.** `corpus_citation_guard` grades markdown LINKS;
a backticked filename is a different construct and was in no fence's population.

## V2 — the pool scope was wrong TWICE before it was right, and both are kept

| pool | `.md` citations resolving nowhere |
|---|---|
| rosetta only, excluding `.agentspace` and `stack-demo` | **45** (38 distinct) |
| rosetta + `rosetta-extensions` + the `stack-demo` clone set | **19** (17 distinct) |

The first reading was **2.4× the second**, and every one of the 26 it lost was a real file in a real
pool the probe had excluded. **Third occurrence of one class inside two iters** — iter-220 kept its own
narrow-scope RED as a control for exactly this reason, and the next iter reproduced it anyway.

## V3 — all 19 adjudicated BEFORE landing: zero defects, four declared classes

| class | members |
|---|---|
| **negated** — the corpus asserts the file does NOT exist | `guidance.md` (`studio-room.md:467`: *"…nor `guidance.md` … anywhere on disk"*) |
| **explicitly future** | `deploy_guide.md`, `debug_guide.md` — both under `corpus/ops/README.md`'s *"## Future Operations — This directory may grow to include:"* |
| **cross-repo, not cloned** | the six `07*-*.md` of `anthropos-knowledge-base`; `ant-singularity/knowledge/…` ×3; `reference_devserver.md` (`kb-ant-business`) |
| **git-ignored workspace artefact** | `.agentspace/profile_gaps.md`, `.agentspace/seeding_gaps.md`, `stack-dev/setup_progress.md` ×3, `op_20260511_…md` |

**0 of 19 is a corpus defect.** The `negated` class is iter-214's shape exactly — a document quoting the
very absence it exists to record — and a census that flagged it would be telling the corpus to stop
saying true things.

## V4 — pre-registered stop condition

The live census must be **0 findings** at land time with the four classes declared, and it must fire on
a staged citation belonging to no class. A citation fitting no declared class is reported as a corpus
defect — **a class may not be widened to absorb it.**
