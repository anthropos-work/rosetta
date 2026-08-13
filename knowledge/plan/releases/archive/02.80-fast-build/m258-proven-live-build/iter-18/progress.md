**Type:** tik — under `TOK-01`, extended by the user's *"i want no debt"* ruling.

# iter-18 — the 8th merge reaches the corpus, and the fence that watches it goes green

## Phase A — measure the platform, and prove the measurement is current

Every clone read at `origin/main` **and verified equal to the remote in the same minute**
(`git ls-remote`, 2026-08-12T12:07Z) — because the whole subject of this iter is a corpus that trusted
a stale reading:

| repo | ref | what changed |
|---|---|---|
| `platform` | `766df6c` | *"remove sentinel service and related configurations"* |
| `app` | `c52dbc51e` | `internal/sentinel/` present; `rpc.go` deleted |
| `sentinel` | `f2c46190` | v0.24.2 — **the port source, and the repo head** |

`docker-compose.yml` 190 → **164** lines, **four** services (`backend` `:5`, `studio-desk` `:90`,
`next-web-app` `:121`, `gotenberg` `:148`). `repos.yml` 28 → **13** lines, **three** repos.
`postgresql` + `redis` are the whole of the **included** `common.yml` and carry no `profiles:` key —
so the floor is **two** and `core` starts **four**.

## Phase B — the fenced map: `rc=1` → `rc=0`

Run directly (never through a pipe): **`rc=1`, 17 findings** — one `[B departure]` plus **16 citation
failures**. 8 of the 16 were the class a human cannot see: citations still *inside* the file but
inside **another service's block**. Repaired; **`rc=0`**, `OK OVER ITS REACH`.

The `sentinel` row is now **`mid-fold` / `decommissioned` / no** — `mid-fold`'s first holder since
`storage`, and the reason is the schema (`D88`).

## Phase C — the sweep, by class not by mention

The census is 71 files / ~410 mentions, and **most are still true** — the `sentinel` schema, the table,
the Casbin model and the history all survive the fold. Six structural classes were wrong; those were
corrected across **26 files**:

| class | was | is |
|---|---|---|
| always-on floor | `postgresql`, `redis`, `sentinel` | `postgresql`, `redis` — and in `common.yml`, not `docker-compose.yml` |
| `core` containers | 5 | **4** |
| `repos.yml` | 4 repos | **3** |
| cross-process Connect-RPC edges | 1 (`backend → sentinel`) | **0** — the listener itself is gone |
| Tier-1 Go services of ours | 2 | **1** (`app`) |
| folded domains | seven | **eight** |

## Phase D — gates

**Ten corpus fences, all `rc=0`:** `platform_alignment_guard` · `anchor_construct_guard` ·
`demo_knob_guard` · `decommissioned_instruction_guard` · `corpus_citation_guard` ·
`corpus_index_guard` · `markdown_structure_guard` · `fence_command_guard` ·
`evidence_visibility_guard` · `dev_flag_guard`.

Three of those were RED on arrival and are green by repair, not by exemption:

- `anchor_construct_guard` — 5 out-of-range ranges + 6 anchors landing on a `}` or a blank line. Two of
  the six were caused by **this iter's own edits** (inserting the map's banner moved the line a sibling
  doc cited) — a live instance of `FIX-M257-anchor-guard-content-drift`, caught and fixed in-iter.
- `demo_knob_guard` — 13 stale `up-injected.sh` anchors (rext drift, pre-existing); all re-pinned.
- `decommissioned_instruction_guard` — went RED **because of Phase B**: the moment the map graded
  `sentinel` decommissioned, the guard derived it into its set and found two live `cd sentinel`
  instructions. Marked historical, with the requirement's new owner named.

`platform_predicate_guard` stays **`rc=1` on one finding**, and it is a **false positive the fence owns,
not the corpus**: `docker-desktop-vm` is a **host** profile, not a compose profile (`D94`). Pre-existing,
routed with its diagnosis. **The correct prose was not edited to make the fence green.**

### The suite, and the attribution done properly

`stack-core`'s full suite is **58 failed / 2344 passed** on a run that straddled the edits, which is
not evidence of anything. It was attributed with a **pristine `git archive HEAD` extract** (`D96`), and
the answer is exact:

| run | failures |
|---|---|
| pristine HEAD corpus, 18 modules | **46** |
| post-edit corpus, same 18 modules | **47** |
| **introduced by this iter** | **1** |
| **fixed by this iter** | 0 (7 in-flight failures were my own mid-run edits, already green) |

**Positive control:** the pristine extract reproduces `platform_alignment_guard` at **`rc=1`, 17
findings, `[B departure]` included** — so the RED this iter repaired was demonstrably pre-existing and
the extract is a valid substrate.

The one introduced failure was **real, and it was mine**: `test_service_doc_status_fence::
test_the_banner_detector_can_actually_say_no` hard-coded `sentinel.md` as its *"live service with no
banner"* specimen. Fixed in rext by **deriving** the specimen from the map (`D97`), RED-proven with a
`has_banner → True` mutant, tagged **`fast-build-m258-iter-18`** and **verified on origin**
(`git ls-remote`). The declaration `.agentspace/rext.tag` was re-pinned with it.

## Close — 2026-08-12

**Outcome:** **The 8th service merge is in the corpus, and the fence that watches it is green.** Platform
`766df6c` folded `sentinel` into `app` at v11.0 and the corpus described a topology the platform no
longer had. The alignment fence, run directly against the new `repos.yml`, was **`rc=1` with 17
findings** — the `[B departure]` arm firing exactly as designed, plus 16 citation failures of which 8
had drifted **into another service's block**. It is now **`rc=0`**, and nine other corpus fences are
green with it, three of them repaired rather than exempted. The structural drift is corrected across
**26 files** in six classes (floor 3→2 · `core` 5→4 · `repos.yml` 4→3 · RPC edges 1→**0** · Go services
2→1 · folded domains 7→8), with the prod state graded **`mid-fold`** rather than `merged-into-app`
because two of §1's three tests fail (`D88`).

**Type:** tik
**Status:** closed-fixed
**Gate:** N/A — the milestone's gate closed by user ruling (`D52`); clause 3 remains NOT MET and is not
recorded as met. This iter took no timing measurement and offers none.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n
*(the G1 false positive is a routed fence defect with a named handler, not a mid-iter question that
changes what lands)* — (5) cap-reached: n *(1 tik)* — (6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **continue**

**Decisions:** D88–D97

**Side-deliverables** (found by the sweep, not planned, recorded separately so they do not blur the
close status):

- **`app/rpc.go` no longer exists** — deleted at `a85e8308d`, pinned by `app/rpc_removal_test.go`.
  `backend.md` named it as *"the top-level wire-up"* to look at.
- **`platform`'s `make bootstrap-dev` is broken** at `766df6c` (`D92`). Reported, **not fixed** —
  0 platform-repo edits.
- **Two published backend ports bind nothing, and the live one is unpublished** — measured on
  `demo-3`'s own compose network: `8081` and `8083` have no listener, `8084` (the meta server) answers
  and compose does not publish it (`D89`).
- **An rext fence repair** (`D97`) — `test_service_doc_status_fence`'s negative control now derives its
  live specimen from the map instead of naming `sentinel.md`. Tagged `fast-build-m258-iter-18`, on
  origin, RED-proven. Separate commit in the rext tree.

**Routes carried forward:**

- **`ROUTE-M258-iter18-g1-reads-host-profiles-as-compose-profiles`** (net-new, `D94`) — `_PROSE_PROFILE`
  has negation, postfix-negation and ref-pin discriminators but **no domain discriminator**, so every
  *"a `X` profile"* in the corpus is graded as a compose token. Needs an rext change + RED-proof
  mutants + tag + push; deliberately not taken blind.
- **`ROUTE-M258-iter18-app-row-anchors-are-at-2035f9a`** (net-new) — the map's `app` row pins seven
  wiring anchors at `app` `2035f9a`; `origin/main` is `c52dbc51e`. They pass range-only, so no fence
  sees them, and one sibling anchor in that row was already landing on a closing brace when checked.
  Mechanical, ~7 anchors.
- Unchanged from iter-17, **not re-verified this iter** (no host time spent):
  `REPORT-M258-iter17-public-host-default-skips-the-batch` ·
  `REPORT-M258-iter17-dev-ui-images-stay-pre-L1-fat` · `ROUTE-M258-iter17-batch-gate-has-no-dev-opt-in` ·
  `ROUTE-M258-iter17-registry-is-empty-while-a-stack-is-up` · `FIX-M258-iter14-purge-leaves-276MB` ·
  `TARGET-M258-iter13-browser-only-deps` · `SETTLE-M258-iter13-studio-desk-cold-time` ·
  `ROUTE-M258-iter13-dockerfile-not-in-cache-key` ·
  `ROUTE-M258-iter15-compose-down-cannot-parse-an-older-stack`.
- **`END-M258-one-stack` is UNCHANGED and still owed on the *correctly-built* stack.** `demo-3` is up
  and is the stack the **buggy** tooling built; its own `batch-gate.json` still reads
  `verdict: red, red_count: 15` and its `autoverify.json` `green: false`. The next iter builds a fresh
  stack with the fixed tooling, verifies it, and only then tears `demo-3` down.

**Lessons:**

- **A fence that is never run is not a fence.** This one was correct, unchanged, and RED for a day.
  Its `[B departure]` arm is the arm that has fired every single time a service left.
- **Retracting a citation verbatim re-arms it — and can launder the very drift it retracts.** Two
  independent guards caught the same shape on this edit. Write retracted citations without the
  `path:line` form.
- **Grade a fence's RED before believing it.** One of the two remaining findings was the corpus's
  (fixed); the other is the fence's (routed). Editing correct prose to reach green would have looked
  identical from the exit code.
- **A guard can go RED because you made another document more correct.** `decommissioned_instruction_guard`
  only found the `cd sentinel` blocks *after* the map called sentinel decommissioned. Fences compose;
  expect the second one to fire.
- **The count was wrong in both directions.** `service_taxonomy.md`'s floor bullet said *two* for four
  releases, was corrected to *three*, and is now *two* again for a different reason. That is what makes
  it a fenced number rather than a restated one.
