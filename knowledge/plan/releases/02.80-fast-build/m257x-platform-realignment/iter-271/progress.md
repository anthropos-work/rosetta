# iter-271 — progress

**Type:** tik, under `TOK-08`. Shape: prove-it-cold (multi-step declared in `overview.md`).

Route: `ROUTE-M257x-270-prove-the-spent-pin-cold`.

## Phase 1 — pre-registrations sealed

Sealed in this iter's first commit, before any measurement. See `overview.md` § Pre-registrations.

Rung zero (`corpus/ops/verification.md` pre-flight): the SoT pin `fast-build-m257x-iter-270` **is on
origin** — `refs/tags/fast-build-m257x-iter-270` → `4e5fb25`, peeled `2833a64`. Measured 2026-08-10T19:59Z.

## Phase 2 — re-point the consumption clone to the SoT pin

The M217 FATAL pin guard **requires** the operator to check the clone out; `ensure-clones.sh` only
*verifies* the match and dies on a mismatch. Measured before/after (2026-08-10T20:01:39Z):

| | ref |
|---|---|
| `stack-demo/rosetta-extensions` before | `fast-build-m257x-iter-101` (the pin iter-260's proof ran at) |
| `.agentspace/rext.tag` (SoT) | `fast-build-m257x-iter-270` |
| after `fetch --tags` + `checkout` | `fast-build-m257x-iter-270` = `2833a647` — matches SoT **and** the authoring copy |

Tracked tree clean before and after (two untracked playthrough report artifacts, pre-existing).

**Both halves of every cycle therefore run the new tooling** — teardown included. That is deliberate:
iter-260's three cycles were byte-identical invocations at one pin, and this iter reproduces that shape at
the other pin, so the two are comparable.

## Phase 3 — cycle 1

`rosetta-demo down 2 --purge` + bare `up-injected.sh 2`, no flags, no retries. **All durations CONTENDED**
(shared 12-CPU / 24 GiB host, third-party load outside our control) and none is a baseline.

### Teardown

`DOWN_EXIT=0`, 2026-08-10T20:02:02Z → 20:02:18Z. Containers, network and data purged; stack images removed;
hostlock released. **`demo-1` untouched — 11 containers still up**, asserted immediately after (it is not
ours; `--purge 1` is forbidden).

### PR-3 — GRADED, HOLDS

Line 93 of the bring-up log:

```
==> [demo-2] injecting: app (derived from the platform compose's build set: sentinel app)
```

No `die`. The derived build set is `sentinel app` — the two Go services still built — and the injected set
is `app` alone. Neither `cms` nor `jobsimulation` appears.

Note the **second-order** signal, which is the part worth keeping: there are **zero** *"no longer built by
the platform compose — skipping"* lines. Before iter-270 those two lines fired on every single bring-up —
the filter was silently correcting a candidate list that named two destroyed services. The list is now
correct at the source, so the filter has nothing to correct. **The filter is still load-bearing**: it is
what catches the *next* candidate whose compose service disappears.

### PR-4 — GRADED, HOLDS

The generated override
(`demo-stack/stacks/demo-2/docker-compose.injected.yml`) declares **11** service keys:

`backend · gotenberg · postgresql · redis · sentinel · directus · next-web-app · studio-desk · hiring-app ·
fake-fapi · fake-bapi`

Intersection with {`cms`, `jobsimulation`, `roadrunner`, `storage`, `skillpath`, `messenger`,
`customerio-sync`} = **∅**. The platform floor + `core` profile, plus the six demo-owned additions.

### PR-2 — the acquisition FETCH arm is NOT exercised by a purge cycle (disclosure)

Line 15: `[ensure-clones] app: studio/ already populated — reusing (idempotent)`.

**`down --purge` purges the STACK, not the CLONE SET.** Containers, volumes and images go; `stack-demo/app`,
`stack-demo/cms` and their `studio/` trees persist. So a cold cycle exercises the acquisition's **reuse**
arm and never its **fetch** arm — which is exactly the arm iter-270 rewrote. Saying "the cold cycle proves
the studio repair" would be the scoped-green error this milestone keeps catching (r60/66).

What the cycle *does* establish for PR-2's second half: `stack-demo/cms/studio/requirements.txt` is not
re-written (baseline mtime `2026-07-31T15:45:41Z`, captured at `20:01:54Z` before the window opened).

So the derivation was graded **directly**, with its precondition removed — the standard the last three
harden passes failed to meet. `studio_consumer_names` from `stack-core/lib/studio.sh`:

| arm | input | result |
|---|---|---|
| **A** — precondition PRESENT | the real `stack-demo/platform/repos.yml` | `rc=0`, returns `app sentinel next-web-app studio-desk` — **`cms` absent** |
| **B** — precondition ABSENT | a non-existent path | **`rc=1`** + *"cannot derive the studio-consumer set … Refusing to fall back to a hardcoded service list"* |
| **C** — present but vacuous | a `repos.yml` declaring no `- name:` | **`rc=1`** + *"declares no '- name:' repo — refusing to guess"* |

Two distinct absent-precondition shapes, both RED. The old code's fallback collapsed to `cms` **alone** —
i.e. it dropped every live consumer and named a corpse; arm A shows the replacement returns the four live
repos `repos.yml` actually declares.

### Cycle 1 — GREEN

`stack-demo/rosetta-extensions/demo-stack/stacks/demo-2/autoverify.json`, copied to the scratch record:

```json
{"green": true, "warnings": 0, "project": "demo-2", "offset": 20000, "ts": "2026-08-10T20:08:53Z"}
```

Elapsed, **CONTENDED and not a baseline**: teardown 20:02:02Z → 20:02:18Z (16 s); bring-up 20:02:22Z →
20:08:56Z (~394 s); down+up ≈ **414 s**. Host load 2.66–4.35 through the run. Faster than iter-258's 717 s
because the BuildKit cache was warm — `--purge` removes the stack's **images**, not the build cache, which
is what "cold" has always meant for this gate (iter-260 used the same definition).

The bring-up's own asserts, in its words: taxonomy replayed `public.skills = 42790`; 5 shared positions +
42 candidate hiring sessions; academy catalog rendering real course cards; cockpit answering on `:27700`;
fake-FAPI answering so hero login is possible; **all 11 expected containers running**; demo-patches all
applied, none refused, none skipped.

**One honesty note about "0 warnings", so nobody mis-reads it later.** The machine verdict is
`warnings: 0`. The human-readable log still prints **two `⚠` lines**, both documented-by-design and neither
counted by autoverify:

1. `stacksnap: ⚠ surface "sim-embeddings" was captured from schema "cms", but on demo-2 its tables live in
   "public" — replaying into "public".` The replay resolves the target schema **from the TARGET's own
   catalog, not from a declared constant** — this milestone's own §2 rule, and the consolidation working
   rather than failing.
2. `⚠ verify scope: this stack also runs hiring-app fake-fapi fake-bapi, which the probe registry has no
   row for — running and UNGRADED, not absent.` A scope disclosure of the kind r60/66 exists to force.

`green: true / warnings: 0` is therefore a statement about the **graded** set, and the log says plainly
which containers are outside it. Both readings are correct; only the conflated one is wrong.

**A gap in cycle 1's own record, disclosed rather than papered over:** cycle 1's bring-up was launched
detached (`nohup … &`) to keep the session responsive, so its **numeric exit code was not captured** —
only its terminal marker (`UP. Clerk-free demo-2 is live.`) and its autoverify verdict. `READY` is defined
as *exit 0 **and** a green `autoverify.json`* (`build-budget.md`), so half of cycle 1's definition rests on
a log line rather than an `rc`. Cycles 2–4 are therefore run through a fixed runner
(`.agentspace/scratch/work-m257x/iter271/cycle.sh`) that records `DOWN_RC`/`UP_RC` and real UTC timestamps
around both halves — giving **three consecutive fully-instrumented cycles**, with cycle 1 as corroboration
rather than as one of the three.
