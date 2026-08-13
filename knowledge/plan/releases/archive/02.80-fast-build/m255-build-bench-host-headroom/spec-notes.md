# M255 — spec notes

## Pre-flight audits — item 5a (spike a) and every later section

**`/developer-kit:audit-kb-fidelity --milestone=M255` — verdict `RED`** (3 blockers · 9 stale-load-bearing ·
4 blind areas · 21 incidental · 4 known-tracked · 1 process finding).
Report: [`kb-fidelity-audit.md`](kb-fidelity-audit.md).

All three blockers were **resolved in this milestone** (Fate 1) — see
[`decisions.md`](decisions.md) D-M255-3, D-M255-4, D-M255-5:

1. D-v28-7's *"inert for the hiring image"* premise → **refuted in code**, waiver rewritten, plan corrected.
2. `demopatch-spec.md` §4's *"4-manifest union"* (→ **7**) and *"revert LIFO, identical on both"* (→ neither
   is LIFO and they differ) → **corrected with a derived table**, and the narrow invariant that *does* hold
   is now machine-fenced.
3. `frontend-tier.md` arguing against spike (a)'s own build shape → **the third shape is documented**, and
   the stale "forbidden upstream PR" list is annotated with what has since shipped without one.

Both `Delivers` targets were correctly treated by the audit as **planned, not blind**: `build-budget.md` is
genuinely net-new, and `safety.md` §3 was fully silent on cert expiry (landed at **§3.5.4**, the insertion
point the audit recommended).

**Process finding, recorded rather than argued with.** The audit ran *concurrently* with implementation
(the sub-agent was launched in parallel to keep the barrier off the critical path), so it observed a moving
tree and every `up-injected.sh` anchor it cited shifted underneath it. That is a real cost — it made three of
its findings harder to state — and it is why the anchors in `overview.md`/`roadmap.md` are now written as
*"at design time"* rather than as live citations, while the ones in the **corpus** are machine-fenced
(clause 5) instead of hand-maintained.

Audit reuse: one audit covered the whole milestone (single subsystem — `stack-core` + `demo-stack` + the demo
corpus docs; no section left it).

---

## The `buildbench` ledger schema (`buildbench/v1`)

One `ledger.json` per rep under `<out>/rep-NN/`, plus a `campaign.json` aggregate.

| block | contents | why it is there |
|---|---|---|
| `host` | hostname, kernel, cores, docker version, **profile name** | a number without a host is not a number |
| `profile` | the whole measured host profile, inlined | so a ledger stays readable after the profile changes |
| `invocation` | both argvs, cwd, **rext HEAD + tag + dirty flag**, lane count | reproducibility; and rung zero is a real failure mode |
| `demo_env` | **every** `DEMO_*`/`STACK_PUBLIC_HOST` knob → value · source (`env`\|`default`) · default · `read_at` | the gap that made a gate run unattributable — see below |
| `pre_state` / `post_state` | `docker system df` (structured), Docker-VM free disk, host free disk | the campaign protocol's per-rep declaration |
| `phases` / `sub_phases` | `@@PHASE@@` pairs + the 12 anchor-derived sub-phases | the attribution model |
| `phases_complete` / `missing_anchors` / `not_applicable_anchors` | fail-closed, with the switched-off case separated | an empty phase table must never read as a pass |
| `builds` | per-service BuildKit steps, **export split into layers + unpack** | where the time actually is |
| `samples` | peaks/averages incl. **disk `%util`** (spike d) | load1 alone cannot tell a plateau from a ceiling |
| `headroom` | the 3-clause assert result + `max_parallel_ui_lanes` | the gate input |
| `verdict` | the stack's own `autoverify.json` | READY means verified-working |

**Why `demo_env` is load-bearing and not bookkeeping.** `autoverify.json` emits exactly
`{project, offset, warnings, green, ts}` (`stack-verify/live/autoverify.sh:382-386`) and the phase log records
no scope either. `autoverify.sh` *reads* `DEMO_NO_UI` to scope its probes but never *writes* it — so a
`DEMO_NO_UI=1` cycle and a full-UI cycle leave **indistinguishable artefacts** while differing by ~66 % of
the wall clock. A test asserts exactly that pair of facts.

The knob list is single-sourced from `demo_knob_guard.parser_knobs()`, so a knob added to a script appears in
the snapshot with no edit to `buildbench.py`.

---

## Round-trip fidelity — the parsers reproduce the annotation

The strongest available check on the harness is that parsing the **annotated n=1 run** reproduces the numbers
`build-annotation.md` published by hand:

| annotation | published | `buildbench parse` |
|---|---|---|
| §3 P1 teardown | 19.7 s | **19.72** |
| §3 P4 bring-up | 650.7 s | **650.70** |
| §4.4 backend builds | 26.2 s | **26.19** |
| §4.6 next-web | 227.3 s | **227.49** |
| §4.7 studio-desk | 11.2 s | **11.03** |
| §4.8 hiring | 207.9 s | **207.94** |
| §4.9 compose up | 36.7 s | **36.73** |
| §4.10 registry + serve | 19.2 s | **19.19** |
| §5.2 hiring export | 136.7 s | **136.7** (99.0 layers + 37.6 unpack) |
| §6 peak load1 / mem / swap / min disk | 4.90 / 5,405 / 2,452 / 34 | **identical** |

Sub-phases sum to **651.42 s** against a P4 of **650.70 s** (Δ 0.72 s = the post-state capture the last phase
absorbs). A table that does not add up would mean something is unattributed.

---

## Things that bit, recorded so they do not bite twice

- **`Sampler._stop` shadowed `threading.Thread._stop`.** `Thread.join()` calls that private method, so every
  campaign completed its work and then died writing the ledger. Cost: one 11-minute cycle on `billion`, whose
  numbers were recovered afterwards with `buildbench parse` (total **652.19 s**; it is the first cycle with
  disk-`%util` data and it is what answers spike (d)).
- **A curated test PATH is part of the test.** Eleven cert tests went red because `grep`/`sed` were missing
  from the harness's clean `PATH`, so `cert_needs_mint`'s SAN-check pipeline failed and every *keep-existing*
  case silently became a *re-mint* case — asserting the opposite of its own name. The file's own header
  already warned about this class (`touch`, M220); it recurred anyway.
- **A fixture that is not a certificate cannot test a certificate predicate.** The keep-existing tests wrote a
  text file reading `PRE-EXISTING`. Correct under the old file-existence guard; under a predicate that parses
  the cert, it is *unreadable → re-mint, fail-safe*. Fixtures are now real openssl mints.
- **`assertNotIn("exit ", region)` graded prose.** It matched the word *"errexit"* in a comment. Now matches
  an exit **statement**.
- **A guard clause written the wrong way round can never fire.** The chain-order check built its observed
  order by iterating the *expected* chain — reconstructing the right answer by construction. Only the
  RED-proof mutation exposed it. Every clause in both new guards is now RED-proven.

---

## Campaign protocol as executed

<!-- M255-CAMPAIGN-NOTES -->

---

## Artefacts

| what | where |
|---|---|
| spike (a) raw | `billion:/home/devops/panorama/m255/spike-a/{spike-a.log,build-A.log,build-B.log}` |
| the recovered first cycle (spike d's data) | `billion:/home/devops/panorama/m255/recovered-rep01/` |
| the n≥3 campaign | `billion:/home/devops/panorama/m255/campaign/{rep-NN/ledger.json,campaign.json}` |
| the informational laptop run | `.agentspace/scratch/work-m255/laptop/{run.log,build.log,samples.tsv}` |
| spike write-ups | [`../evidence/m255-spikes.md`](../evidence/m255-spikes.md) |
| rext tags (on origin) | `fast-build-m255-buildbench`, `fast-build-m255-buildbench-1` |
