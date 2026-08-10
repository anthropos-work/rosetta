---
iter: 270
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
active_strategy: TOK-08
route: FIX-M257x-268-ensure-clones-hardcodes-cms-as-studio-fetcher (+ ROUTE-M257x-h68, FIX-M257x-262-dev-path-needs-the-studio-acquisition)
---

# iter-270 — spend the frozen pin: the hardcoded service lists, and the derivation that fails OPEN

**Type:** tik, under `TOK-08` (*census the mechanical classes; stop sampling them*).

## Step 0 — re-survey (mandatory, before targeting)

`D-M257x-258-1` froze the rext pin at `fast-build-m257x-iter-101` **deliberately**, as an experimental
control: hold the tooling constant so the platform-ref advance is the single changed variable. Re-surveyed
at open (2026-08-10T18:19Z):

| fact | measured |
|---|---|
| `.agentspace/rext.tag` | `fast-build-m257x-iter-101` |
| rext authoring copy vs that tag | **205 commits** ahead |
| rext authoring copy vs `origin/main` | **12 commits** ahead (unpushed) |
| authoring tree | clean, on `main` |

**The control's question is answered.** iter-258 (demo green on the advanced refs, first attempt),
iter-260 (three consecutive `--purge` + `up` cycles green) and iter-262 (a dev stack from current `main`)
all landed *under the frozen tooling*. Clause 1 is MET; the dev half is MET. Holding the pin longer buys
nothing and costs every fix routed behind it — **six routes now name "needs a tag + pin bump"**.

So this iter spends the control, and says what it bought.

## Cluster / target identified

**The class: rext names decommissioned services in hand-maintained lists, and the one derivation built to
replace such a list FAILS OPEN.** This is the user's closing-condition limb — *"the deprecated/removed
repos no longer treated as part of the project"* — on the substrate iter-268 first measured it on (disk +
source), not in prose.

Population as routed (4 sites, 2 files + 1 module):

| # | site | naming | route |
|---|---|---|---|
| 1 | `demo-stack/ensure-clones.sh:310` `_studio_repos="cms"` | `cms` is *preferred first* studio fetcher | `FIX-M257x-268-…` |
| 2 | `demo-stack/up-injected.sh:216` `INJECT_CANDIDATES="app cms jobsimulation"` **and `derive_inject_svcs`'s fail-open arm** | `cms`, `jobsimulation` | `ROUTE-M257x-h68-…` |
| 3 | `stack-injection/gen_injected_override.py:52` `INJECTED` | `cms`, `jobsimulation` | `ROUTE-M257x-h68-…` |
| 4 | `stack-injection/gen_injected_override.py:153` `REUSE_DEV` | `storage`, `roadrunner` | `ROUTE-M257x-h68-…` |

Plus the sibling the same file-pair carries: **the dev path has no studio acquisition at all**
(`FIX-M257x-262-dev-path-needs-the-studio-acquisition`) — demo's is anchored on a corpse, dev's is absent.

## Hypothesis

The four sites split into **two kinds, and only one has teeth**: dead entries that can never match (a map
keyed on a compose service the compose no longer declares) versus a **live fallback** that runs on the
failure path. `derive_inject_svcs` returning 0 on a failed derivation is the milestone's own recurring
defect — *a capability probe that fails OPEN disarms the check it guards* — and it is the only one of the
four that can change what a bring-up builds.

## Expected lift

- The class is closed by its **population**, not its last member: every site named, each graded
  live-vs-dead **with the evidence for the grade**, and the live one repaired fail-CLOSED.
- The fail-closed arm **proven to fire** on a tree where its precondition is absent — the standing failure
  mode of the last three harden passes, and the bar this iter must clear.
- A tag on **origin** and the pin bumped, so the six behind-the-pin routes stop being blocked.

## Phase plan (declared multi-step — the tripwire counts UNPLANNED lines only)

1. Seal these pre-registrations (first commit).
2. Measure the population and grade each site live-vs-dead.
3. Repair: derive the studio-consumer set with no hardcode; make `derive_inject_svcs` fail CLOSED; prune
   the dead entries; give the dev path the shared studio acquisition.
4. Prove the new fail-closed arm goes RED where its precondition is absent (not merely GREEN where present).
5. Run the touched suites + the guard family; tag; `git push --tags`; verify **on origin**; bump the pin.
6. Corpus: record what the spent control bought.

## Out of this iter's planned scope (declared, so the tripwire is clean)

`FIX-M257x-269-force-append-grows-the-demo-env-without-bound` — a different subsystem (`stack-secrets`
Go), a different mechanism, and a fix that must preserve values-blindness. It rides the **next** tag.
Routing it forward is a Fate-3 with a named handler, not a deferral.

## Escalation conditions

- **Force-push of any kind is forbidden**, including `--force-with-lease`. A tag that collides is a NEW
  tag name, never an overwrite. (`F18` already records three `iter-101*` tags disagreeing; do not add a
  fourth pattern.)
- `demo-1` is not ours — no stop, restart or `--purge` of it, and no bring-up is required by this iter.
- If a repair cannot be proven to fail closed, it does not ship. A fence that cannot fire is the defect.

## Acceptable close-no-lift outcomes

A documented falsification — e.g. all four sites graded **dead**, with the mechanism for each, and the
fail-open arm shown unreachable — would close the class with evidence and no code change. That is a real
result under this milestone's rules.

## Pre-registrations (sealed in this iter's FIRST commit, before any measurement)

Stated falsifiably, before looking.

- **PR-1 — the fail-open arm is untested.** No test in the rext tree asserts the behaviour of
  `derive_inject_svcs` when the derivation yields empty (i.e. no test pins the unfiltered-fallback path).
  *Refuted by:* finding ≥1 test that exercises it.
- **PR-2 — `INJECTED` and `REUSE_DEV` are DEAD, not live.** Pruning `cms`/`jobsimulation` from `INJECTED`
  and `storage`/`roadrunner` from `REUSE_DEV` changes **zero bytes** of the override generated against the
  current platform clone. *Refuted by:* any diff.
- **PR-3 — the `cms` hardcode is the path that actually ran on this box.** `stack-demo/cms/studio/` is
  populated (`requirements.txt` present), i.e. the decommissioned clone was the studio *fetcher*, and
  `stack-demo/app/studio/` was filled by copy from it. *Refuted by:* cms's tree absent or unpopulated.
- **PR-4 — the dev path has ZERO studio handling.** `dev-stack/**` yields **0** matches for
  `studio_required|lib/studio.sh|anthropos-studio-room|STUDIO_REPO`. *Refuted by:* any match.
- **PR-5 — the fail-open population is BIGGER than the one arm `ROUTE-h68` names.** A census of the demo
  bring-up path for capability probes that fall back to a permissive/unfiltered default on failure finds
  **≥3** such arms, `derive_inject_svcs` included. *Refuted by:* a census returning 1 or 2.
