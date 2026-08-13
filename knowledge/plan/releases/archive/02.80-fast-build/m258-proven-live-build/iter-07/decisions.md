# M258 iter-07 — decisions

## D21 — the re-pin is falsifiable, and the feature under test was genuinely unobtainable

`TOK-01` step 4's precondition (iter-06) was *run the campaign from the consumption clone at a pushed
tag.* Re-verified rather than taken on trust, and it holds **in both directions**:

- **On origin:** `git ls-remote --tags origin` → `fast-build-m258-iter-06` = `5f3a3815…`; the authoring
  copy's HEAD is the **same sha**, tree clean. Rung zero satisfied *before* the campaign, not after.
- **Genuinely absent from the consumer:** at the old pin (`fast-build-m258-iter-03`) **all four**
  batch-gate files were missing from `stack-demo/rosetta-extensions/playthroughs/e2e/`
  (`batch-gate.sh`, `restore-presenter-world.sh`, `check-cockpit-roster.py`, `stack-paths.sh`) and
  `up-injected.sh` carried **no** `batch-gate` hook. Post-re-pin all four are present and the hook is at
  `:2844`.

This is the **M236 shape exactly** — *the feature under test could not be obtained at all* — and it is
why "tagging is not publishing" is pre-flight rung zero. Had the campaign been launched without the
re-pin it would have measured a bring-up **with no batch gate in it** and the phase table would still
have summed, exactly as `D17` describes.

## D22 — the stale THIRD pin copy, re-verified inert by path arithmetic

`stack-demo/rosetta-extensions/.agentspace/rext.tag` reads **`fast-build-m257x-iter-279`** — wildly
stale. iter-02 proved it inert; that proof was **re-derived here rather than inherited**:
`rext_tag()` reads `"$root/.agentspace/rext.tag"` with `root=$REPO_ROOT`
(`demo-stack/lib/rext_tag.sh`), and both `ensure-clones.sh:34` and `up-injected.sh:52` set
`REPO_ROOT="$(cd "$HERE/../../.." && pwd)"` — from `…/rosetta-extensions/demo-stack` that is the
**rosetta repo root**. So the SoT read is `/…/rosetta/.agentspace/rext.tag`, and the in-clone copy is
never consulted. Inert — but left in place, because deleting a file nothing reads is not this iter's
call.

## D23 — the two clones own DIFFERENT stacks, and this is the `D19` hazard in a new place

Resolved from the **docker mount** (`stack-paths.sh::resolve_stack_dir`), never from a script's own
location:

| stack | owning clone | whose |
|---|---|---|
| **demo-1** | `.agentspace/rosetta-extensions` (**authoring**) | mine |
| **demo-2** | `stack-demo/rosetta-extensions` (**consumption**) | ⚠️ the **user's** |

Two consequences, both acted on:

1. **The re-pin I performed changed the clone that owns the USER's demo-2.** Verified harmless rather
   than assumed: `demo-stack/.gitignore:8` ignores `stacks/`, so a `git checkout` provably cannot touch
   `stacks/demo-2/` — confirmed with `git check-ignore -v`. demo-2 measured **11 containers, up 35
   hours**, and the dev stack **5**, before and after.
2. **The campaign's own teardown would be cross-clone.** `run_campaign` runs
   `rosetta-demo down 1 --purge` from `rext_root()` = the **consumption** clone, while demo-1's
   `docker-compose.injected.yml` lives in the **authoring** clone's stack dir. `cmd_down` scopes
   containers by compose **project name** (`docker compose -p demo-1`, clone-independent) plus a label
   sweep, so the containers would be found — but the injected-file branch would not be taken.
   **Pre-empted:** tear demo-1 down from the **authoring** clone (its correct owner) immediately before
   launching, so all three reps are consistently consumption-clone cycles rather than rep 1 being a
   cross-clone special case.

## D24 — HEADROOM is a POST-gate *and* a pre-gate, and only one of them aborts

Read from the instrument before relying on it — the two behave differently and the difference decides
whether a contended campaign is worth launching:

- **`pre_rep_assert(profile, lanes)`** runs **before each rep** on an **instantaneous** `load1`. On
  failure `run_campaign` **returns 1 immediately**, *before* `argv_down` — so a refusal is safe (it
  destroys nothing) but it **ends the campaign**.
- **`headroom_assert(..., peak_load1=samples["peak_load1"])`** runs **post-rep** and feeds
  `rep_is_ok`, so a rep whose load **peaked** over the limit is graded unusable even though it ran.

Both grade clause 1 as `load1 ≤ host_logical_cores − 2` = **10** on this 12-core host. Measured live:
`assert-headroom --profile macmini` → **FAIL, load1 16.58 vs 10**, disk 66.9 GiB free (ample).

**Consequence for this iter:** launching into a load1 of 14–20 does not produce a contended-but-usable
p50 — it produces either an immediate `return 1` or three reps that `rep_is_ok` discards. That is why
the campaign was **held** rather than launched-and-reported.

## D25 — the contention is third-party and named

Not "the box is slow" — measured: **8 × `mdworker_shared` + `mds`/`mds_stores`** (Spotlight indexing,
peaking at **227 % CPU**) **plus `a8-cart-runner` at ~99 %**, a build belonging to a *different* user
project (`workspace/hyperspace/anima8/…/m270a2-iter132/`). Neither is mine to stop, and the standing
rule is that this host **is permanently contended and cannot be freed**.

`R=16` runnable threads against 12 logical cores with `U=0` — CPU-bound, not I/O-blocked.

## D26 — clause 4 (0 platform-repo edits) graded early, and it surfaced a route

Clause 4 is a **boolean**, so it is gradeable on a contended host — graded now rather than waiting for
a quiet one. Measured across the demo's own peer clone set:

| repo | HEAD | dirty |
|---|---|---|
| platform | `0c91421` | 0 |
| app | `3eaadae68` | `?? studio/` (untracked — the CI-pulled studio-room, not an edit) |
| sentinel | `f2c4619` | 0 |
| next-web-app | `19423a1fb` | 0 |
| studio-desk | `41ee3575` | 0 |
| **ant-academy** | `249430c3` | **4 tracked files MODIFIED** |

**M258 has edited none of them** — the milestone's own diff is entirely `rosetta` + `rosetta-extensions`.
The ant-academy modifications are **pre-existing sanctioned demo-patches**, self-identifying in the
source (`/* demo-patch: ant-academy-back-to-cockpit … */`) and otherwise gated behind
`ACADEMY_DEMO_FS_PUBLISHED` / `NEXT_PUBLIC_COCKPIT_URL`. Per `demopatch-spec.md` the contract is that the
**canonical `anthropos-work` repos** are never touched; a local peer clone carrying a patch is the
mechanism working, not a violation. **Clause 4: PASS.**

### Route — `ROUTE-M258-iter07-demopatch-G5-does-not-revert-the-native-clone`

**G5 (self-revert) leaves the ant-academy clone dirty**, and the reason is structural rather than a bug
in any one patch: `demopatch`'s apply-then-revert cycle is built around an **ephemeral build-scratch
clone** that the image build consumes and discards. **ant-academy runs NATIVELY** — it is never built
into an image — so its "clone" *is* the durable peer clone `stack-demo/ant-academy`, and there is no
ephemeral copy for the revert leg to throw away. The patches therefore persist across bring-ups.

Why it matters rather than being cosmetic: `demopatch`'s **G2 is drift-refuse**, so a *later* patch
whose `pre_sha256` baseline was taken against a pristine file can read **DRIFTED** against a clone a
*previous* run patched — the silent-refusal class that `demopatch-spec.md` records as having shipped a
76 s members grid for four releases.

**Deliberately NOT repaired here**, on two independent grounds: reverting a tracked file is a
**forbidden op** for this agent, and `stack-demo/ant-academy` is the clone the **user's demo-2** is
serving from — a revert would mutate a running stack that is not mine. Routed with evidence.

## D27 — iter-06's `green: false` is proven to be clone topology ALONE, without running a campaign

iter-06 routed its `postgres-schemas` refusal forward as *"run from the consumption clone"* — a
**hypothesis**. Rather than spend a 3-rep campaign discovering whether it was right, it was reduced to
arithmetic and a one-line query. The probe's two inputs (`stack-verify/lib/readiness.sh:65,71`):

| input | expression | authoring clone | consumption clone |
|---|---|---|---|
| `_rl` | `<rext>/stack-core/lib/repos_yml.sh` | ✅ present | ✅ present |
| `_ry` | `<rext>/../platform/repos.yml` | ❌ `.agentspace/platform/` **does not exist** | ✅ `stack-demo/platform/repos.yml` |

With `_ry` missing, `expected` is empty and the probe takes its **loud** fallback — *"refusing to assert
a hand-maintained list"* — which is the documented correct behaviour, not a defect.

Both halves then verified **live**, still without a bring-up:

- **Derivation:** sourcing `repos_yml.sh` against the consumption clone's `_ry` yields a non-empty set —
  **`extensions`, `sentinel`, `public`** (plus `directus`, added when that container is present).
- **Satisfaction:** `demo-1`'s Postgres reports `extensions`, `sentinel`, `public` **and** `directus` —
  all four. So the probe will not merely *assert*, it will **pass**.

**Therefore `autoverify.json`'s `green` flips `false → true` for topology reasons alone**, and the sole
blocker between this milestone and a gate-usable rep is discharged **before** the campaign rather than
by it. This is `TOK-01`'s own discipline — *measure before engineering* — applied to a routed fix:
**a routed fix is a hypothesis, and this one survived.**

The falsifiable "before" is on disk: `stacks/demo-1/autoverify.json` = `{"green":false,"warnings":1}`
beside `batch-gate.json` = `{"verdict":"green","red_count":0,"batch_seconds":160,"restore_seconds":7}`.

## D28 — NOT running the batch under this load is the correct engineering call, not caution

Watched for **~30 minutes** of foreground polling. `load1` against a limit of **10**:

```
07:28 20.34  07:31 23.62  07:33 11.93  07:34 32.90  07:37 40.61
07:40 13.37  07:42 43.49  07:44 45.92  07:45 62.88  07:46 55.95
```

**Minimum seen: 11.93** — never once below the limit, and the trend over the window is **upward**
(the other project's campaign is ramping, not winding down).

The tempting move is to run it anyway as an operator and label the timings "contended". **That would
have been wrong, and not for timing reasons.** The batch gate runs Playwright with **`retries: 0`** —
by contract, clause 2 of `D-v28-3`, *"it NEVER retries to mask a flake."* A 215-spec browser suite
driven on a box at `load1` 40–60 does not merely run slowly; it **times out**, and every timeout is
recorded as a **red** in a consolidated set that `D-v28-3` clause 4 then **escalates to the user for
renegotiation**.

So an operator run under this load would have manufactured a red set that describes the *host*, and
handed it to the user as if it described the *product*. `batch-gate.sh`'s own header names this exact
failure one layer up:

> a false RED is not the safe direction of this bug: **IT TRAINS ITS OPERATOR TO DISBELIEVE THE GATE.**

The instrument refusing is a **result** (`verification.md`). The right response to a refusal is to
report it with its measurements — not to route around the refusal and generate a verdict the conditions
cannot support.

**Second, independent reason:** launching would have cost `demo-1` its presenter world (verified good at
open — 12/12 seats resolving) for a run whose reps `rep_is_ok` would discard anyway. The launch script
therefore asserts headroom **before** the teardown (`buildbench assert-headroom`), so the demo is never
spent on a cycle that cannot produce a number.

**What this does NOT license.** It is not a gate re-cut, not a deferral of clause 3, and not evidence
about the composition. The projection stands where iter-06 left it (**414.15 s** vs 480) and is still a
projection. The gate needs ~30 minutes of a host with `load1 < 10`; everything else it needs now exists.
