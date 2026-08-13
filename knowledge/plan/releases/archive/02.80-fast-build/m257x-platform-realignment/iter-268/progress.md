**Type:** tik, under `TOK-08`. Route: `ROUTE-M257x-265-stack-demo-carries-six-dead-clones`.

# iter-268 — a decommissioned repo is still the PREFERRED actor on the demo bring-up path

## Why this iter exists

The user's binding closing condition ends *"…with the **deprecated/removed repos no longer treated as part
of the project**."* Every prior iter read that as a statement about the **corpus**. It is also a statement
about **disk**, and nothing had measured it.

## The census

| reading | result |
|---|---|
| `repos.yml` names | `app` · `sentinel` · `next-web-app` · `studio-desk` — **4** |
| `stack-demo/` clones absent from it | **6** — `cms`, `graphql-wundergraph`, `jobsimulation`, `messenger`, `roadrunner`, `storage` |
| their HEAD dates | 2026-06-19 … **2026-08-05** |
| `repos.yml`'s own last change | `838d907`, **2026-08-05** |
| compose build contexts naming any of the six | **0** |
| corpus instructions to clone any of the six | **0** |
| `stack-dev/` equivalents | **0** (it carries `studio-room`, a *required* acquisition, not a corpse) |

## The finding

**`rosetta-extensions/demo-stack/ensure-clones.sh:310` hardcodes a decommissioned repo as the FIRST studio
consumer:**

```
_studio_repos="cms"
if [ -f "$PLAT/repos.yml" ]; then
  _studio_repos="cms $(… derive the rest from repos.yml … | grep -vx cms …)"
fi
```

Everything *except* `cms` is derived from `repos.yml` — the correct pattern, and the file's own comment
states why `cms` is not: *"cms goes first so the sanctioned `make init-studio` stays the fetch that
actually happens."* That was true while `cms` was live. `repos.yml` has not listed it since `d11a403`.

The entry is guarded only by `[ -d "$_sdir" ] || continue`, which makes it **dormant on a fresh box and
LIVE on any box that still carries a `stack-demo/cms/`.** On this box it is live, and the evidence is on
disk: **`stack-demo/cms/studio` is populated.** The studio runtime `app` builds with was fetched by a
**decommissioned repo's Makefile** (`:338-340`, `make init-studio`) and then copied across as the donor
(`:334-336`).

**Nothing is broken by it.** No compose service, no build context, no corpus instruction. That is exactly
why it survived: it is a *preference*, not a dependency, and preferences do not fail.

## Why the existing fence did not catch it

The clone **set** is fenced: `clone_pin_guard.py` derives the allowed key set from `repos.yml` and asserts
both ways, and **iter-222 used it to remove five phantom pin keys — `cms` among them.** That sweep was
correct and complete *for the registry it covered*.

**The studio-consumer list is a SECOND registry, one file over, and the sweep did not reach it.** This is
`platform-alignment.md` §5 — *"a named-consumer list survives the merge that moved the consumer"* (iter-23)
— occurring inside the repo that wrote the rule down, against a registry that a sibling registry's repair
had already been run over. §10 iter-194: *a registry that supersedes a list must reach everything the list
does.*

## Pre-registration grading

| PR | prediction | outcome |
|---|---|---|
| **PR-1** | 6 in `stack-demo/`, 0 in `stack-dev/` | **HELD** — exactly 6 and 0 |
| **PR-2** | nothing fetches them today | **REFUTED, and it is the finding.** The *clone* path does not; the **studio** path enters `cms` by hardcode and, on this box, ran `make init-studio` inside it |
| **PR-3** | nothing builds from them | **HELD** — 0 compose build contexts |
| **PR-4** | 0 corpus acquisition instructions | **HELD** — 0. iter-265's fence population is complete for the acquisition verb |
| **PR-5** | they are fossils (HEAD older than the removal) | **REFUTED** — `messenger` and `storage` sit at **2026-08-05**, the same day as `838d907`, and `cms` at 2026-08-04. They are **current**, not stale, which is what made PR-2's refutation credible rather than a curiosity |

**PR-2 and PR-5 refute together and reinforce each other.** Had the clones been old, a hardcoded `cms`
would be a dormant line. They are current *and* the hardcode is live *and* `cms/studio` is populated —
three independent readings of the same fact.

## Repair, and what is deliberately NOT repaired

**Landed (corpus):** `corpus/services/cms.md` now states the measurement where an operator reads about this
repo — the hardcode, its `[ -d ]` guard, the *dormant-on-fresh / live-here* asymmetry, and the closing
instruction **"do not read a `stack-demo/cms/` as inert."**

**Routed (tooling):** dropping the hardcode needs a rext tag **and a pin bump**, which would spend
`D-M257x-258-1`'s frozen-pin control. Not spent. Routed as
`FIX-M257x-268-ensure-clones-hardcodes-cms-as-studio-fetcher`, with the fix named: derive the whole list
from `repos.yml` and let the `git clone $STUDIO_REPO` branch (which is what `app`'s CI `additional_repo`
does) be the fetch — the `cms` branch then has no consumer to serve.

**Nothing was deleted.** Pre-registered before the measurement, and it held: the six directories are the
evidence, `demo-2` was not stopped or re-seeded, `demo-1` was not touched.

## Close — 2026-08-10

**Outcome:** The user's *"no longer treated as part of the project"* is measured on disk for the first
time. Six dead clones, **0** builds and **0** corpus instructions — and **one live hardcode** naming a
decommissioned repo as the *preferred* studio fetcher, which on this box is the path that actually ran.
Corpus repaired; tooling fix routed with the frozen-pin control unspent.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue

**Decisions:** `D-M257x-268-1` (a sibling registry survived the sweep that fixed its twin).

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-265-stack-demo-carries-six-dead-clones` → **CLOSED** by this census; superseded by the
  concrete defect below.
- `FIX-M257x-268-ensure-clones-hardcodes-cms-as-studio-fetcher` — **new.** Derive the studio-consumer list
  wholly from `repos.yml`. Needs a tag + pin bump.
- `FIX-M257x-262-dev-path-needs-the-studio-acquisition` (tooling half) — **same file, same subject.** The
  two should land in one tag: dev has *no* studio handling, demo has it anchored on a corpse.
- `FIX-M257x-267-capture-the-succession-RESPONSE`,
  `FIX-M257x-266-manual-path-drops-gates-the-automated-path-enforces`,
  `FIX-M257x-265-prose-deletion-instructions-are-out-of-D-reach`,
  `FIX-M257x-262-demo-env-append-is-not-idempotent`, `ROUTE-M257x-258-the-pin-is-157-iters-stale` → open.

**Lessons:**
1. **Fixing a registry does not fix its siblings.** iter-222's phantom-key sweep was correct and complete
   for the clone pin. The studio-consumer list one file over kept naming `cms` for four more releases,
   because nothing derived the *set of registries* — only the members of one.
2. **A preference does not fail, so it does not get noticed.** The hardcode never errors: on a fresh box
   the `[ -d ]` guard skips it, on a stale box it silently wins. Both readings look like success.
3. **Read a closing condition on every substrate it can mean.** *"No longer part of the project"* was
   graded as prose for eleven iters; on disk it was false the whole time.
