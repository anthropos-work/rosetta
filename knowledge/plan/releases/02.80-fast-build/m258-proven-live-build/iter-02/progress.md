# M258 iter-02 — progress

**Type:** tik · **Active strategy:** `TOK-01` step 1 — *measure the composition before engineering it.*

## Phase 1 — `R0` discharged (the re-pin)

The pin existed in **three copies with three values**. Path arithmetic settles which one is load-bearing:
`ensure-clones.sh:34` sets `REPO_ROOT="$HERE/../../.."`, which resolves to the **rosetta root** from
*both* the authoring copy and the consumption clone — so the stray
`stack-demo/rosetta-extensions/.agentspace/rext.tag` (naming `fast-build-m257x-iter-279`, a previous
milestone's tag) is **inert**. Proven, not assumed; left in place (untracked state is not mine to clean).

| copy | before | after |
|---|---|---|
| `rosetta/.agentspace/rext.tag` (canonical, M49 #1) | `fast-build-m257-iter-09` | **`fast-build-m257-close`** |
| `stack-demo/rosetta-extensions` checkout | `8956e69` | **`679a5f7`** (= the tag, exact-match) |
| the stray inside the clone | `fast-build-m257x-iter-279` | unchanged — **inert** |

## Phase 2 — pre-flight, then one cold cycle on `demo-1`

`assert-headroom --profile macmini` → **OK** (`lanes=1 max_parallel_ui_lanes=2 free=64.8 GiB
load1=2.26 vs 12.0 cores`). ⚠️ Bare `assert-headroom` defaults to **`billion`'s** profile — the host
must be named on every invocation, which is cluster 4's shape and is worth knowing before a number is
quoted from it.

`buildbench run 1 --reps 1 --profile macmini --label m258-iter02-compose`, foreground-polled
(04:42:19Z → 04:49:03Z).

## Phase 3 — re-measure

### The bring-up half, at the corrected pin — **395.31 s**

`up_rc=0` · `green=true` · **HEADROOM OK** · **host identity MATCH** · phase table **complete** ·
**ISOLATION FAIL** (below), which takes the campaign RED by contract.

| sub-phase | s | | sub-phase | s |
|---|---|---|---|---|
| `ui_studio_desk` | **115.35** | | `compose_up` | 32.56 |
| `set_dress` | **80.50** | | `serve_and_egress` | 11.40 |
| `ui_next_web` | 60.05 | | `backend_builds` | 3.79 |
| `ui_hiring` | 44.85 | | `seed_tooling` | 2.46 |
| `host_preflight` | 35.40 | | `autoverify` | 2.20 |
| | | | `clones_and_inject` | 1.81 |
| | | | `secrets_provision` | 1.66 |

**Environment, stated with the number** (`latency-budget.md`'s rule): `macmini` M4 Pro, Docker Desktop
VM, containerd store; `load1 2.26` at launch / `2.12` at close; 64.8 GiB free in the VM; the user's
`demo-2` (11 containers) and the 5-container dev stack **resident throughout and untouched**.

⚠️ **This is n=1 and it is NOT comparable to M257's p50 286.99 s without attribution.** It is +108.32 s
against a **three-rep p50** taken at a different cache state, and `set_dress` (80.50 vs 82.04 s) and
`ui_next_web` (60.05 vs the 53.31 s next-web rebuild figure) both land close to M257's, so the delta is
not spread evenly — it concentrates in `ui_studio_desk` (115.35 s), the one UI image **L1 never
touched** (L1 multi-staged the two *Next* apps; studio-desk is Vite/Express). **Not a regression claim
— an n=1 observation with a named suspect**, and the honest next step is the n≥3 comparison, not a
headline.

### The batch half — **NOT MEASURED. Blocked, and the blocker is the finding.**

### ISOLATION FAIL — all three UI images carry a foreign publishable key

`own_pk_fingerprint` `pk_test_MTI3…61fbfaf4` (a minted Clerkenstein key — `MTI3` is base64 `127`, the
loopback FAPI host). All three UI images — `demo-1-{next-web,hiring,studio-desk}:latest` — carry
`pk_test_bWFy…52038877` instead. 8 images checked, `foreign_origins` empty.

**Which of the assert's two named causes, settled by evidence:**

| hypothesis | verdict |
|---|---|
| an image was reused across stacks (cache leak) | **REFUTED** — `build-next-web.log:99` reads `cache miss, executing 596d39fc0d2316a0`; the image was built by this run |
| the `.env.local` overlay was missing at build time ⇒ the bundle baked another pk | **REFUTED too** — `build-next-web.log:111` reads `- Environments: .env.local, .env`; the overlay was present and loaded |

**Neither. The mechanism is a third one the assert does not name, and it is in the stack's env file.**
`stacks/demo-1/.env.demo-1` (128 lines) carries **24 `# --- Clerkenstein injection (S3) ---` blocks** —
one appended per bring-up, the file never truncated. The first 23 blocks set the minted
`pk_test_MTI3…`; **the 24th — written by THIS run, mtime `04:43:08Z`, inside the 04:42:19Z–04:49:03Z
window — sets `pk_test_bWFy…`**, and **last-wins** hands every consumer the foreign key.

**Two writers, both visible:**

- `stack-injection/inject.py:89` — `f.write(f"\n# --- Clerkenstein injection (S3) ---\n")`, **append,
  never truncate**. That is the 24 blocks.
- `up-injected.sh:2036` — `PK_DEMO="$(python3 …/inject.py …  2>/dev/null)" || true`. **stderr to
  `/dev/null` and failure tolerated.** If the mint degrades, nothing downstream says so.

That pairing is verbatim the class `verification.md` § *The four cheap-wins verify could not see*
records for the demo-patch rot: *"the reason was piped to `/dev/null`, and nothing downstream
noticed."* **The ISOLATION assert is the only thing in the stack that caught this** — which is M257's
*land each falsifiable assert WITH the lever that can trip it*, paying out on its first campaign here.

**I did not cause it — checked before routing.** `git diff fast-build-m257-iter-09..fast-build-m257-close
-- demo-stack/up-injected.sh` is **21 insertions / 9 deletions and every hunk is a comment or a `log`
string** (the ~3.7 GB / ~3 min retractions). No executable logic changed across the re-pin.

**Bounded, and the user's stack is clean.** `demo-2`'s `.env.demo-2` has **8** blocks and its **last**
`NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` is `pk_test_MTI3…` — **correctly minted**. The condition is
`demo-1`-local. `demo-1` was brought up **`--no-public-host`** (localhost-bound, `D2`), so nothing off
this box can reach it.

**Why this blocks the batch.** With the UI tier wired to a real Clerk app rather than Clerkenstein, the
cockpit's password-free hero logins cannot succeed — the browser's clerk-js finds no session and loops
to `/login` (`verification.md` § M218 iter-03 D8/F-6, the same wiring). A Playthrough batch run on this
stack would measure a login failure 30 times, not the suite. **Measuring it here would have produced a
number, and the number would have been meaningless** — which is the distinction `TOK-01` exists to keep.

### Metric

| | |
|---|---|
| bring-up half | **395.31 s** (n=1, contended, labelled) |
| batch half | **still unmeasured** — blocked by the ISOLATION defect |
| gate distance | ungradeable this iter; the composed p50 has no second term yet |

**Metric delta 0 — by design** (`TOK-01` step 1 measures, it does not optimise) **and by blockage**
(the second half could not be taken). Both stated, because they are different reasons.

## Scope-creep tripwire — FIRED, and honoured

Investigation lines opened: (1) cache-leak vs missing-overlay, (2) the env-file's structure, (3) which
writer produced the final block, (4) where `PK_DEMO` originates. **The tripwire fires at the third.**
Line 4 was capped at a single grep — enough to name the two writers with `file:line` so the route is
actionable — and root-causing *why* `inject.py` emitted a non-minted key is **routed, not absorbed**.

## Close — 2026-08-12

**Outcome:** `R0` discharged (rext re-pinned to `fast-build-m257-close`, consumption clone re-checked
out, the third pin copy proven inert). The bring-up half re-measured at the corrected pin: **395.31 s**,
n=1, green / HEADROOM OK / identity MATCH / phases complete. **The batch half could not be measured**:
ISOLATION went RED on its first campaign and the cause is a real, evidenced defect — `.env.demo-1`
accumulates one Clerkenstein injection block per bring-up (24 present) and this run's block carries a
**non-minted** publishable key, so last-wins wired all three UI images to a real Clerk app.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n *(one iter, not the 3-tik
trigger, and the 600 s composed-p50 valve has no composed figure to fire on yet)* — (4) user-blocker:
**n** *(re-graded: nothing here needs a user DECISION. A defect with a named mechanism, two `file:line`
writers and a proven-clean blast radius is a fix, not a choice — Phase 5 §4's NOT-list, "new findings
discovered mid-iter → route forward". I did not introduce it, and the user's stack is unaffected.)* —
(5) cap-reached: n *(1 tik)* — (6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: exit-7
**Decisions:** D5, D6, D7 (this iter's `decisions.md`)
**Side-deliverables:** none — the re-pin is planned scope, not a side-fix.
**Routes carried forward:**

- **`FIX-M258-iter02-inject-appends-and-swallows`** → **iter-03**, and it is the next tik's first job
  because it gates the milestone's primary unknown. Two defects, one symptom: `inject.py:89` **appends**
  a Clerkenstein block instead of rewriting (24 accumulated in `.env.demo-1`, so *last-wins* makes the
  file's history outrank its intent), and `up-injected.sh:2036` runs it with **`2>/dev/null` and
  `|| true`**, so a degraded mint is invisible. Fixing the append alone would mask the swallow; fix both.
- **`CHECK-M258-iter02-studio-desk-is-the-untouched-leg`** → iter-03 or later. `ui_studio_desk` at
  **115.35 s** is the largest UI leg and the one L1 never touched. It is the named suspect for the
  n=1 vs n=3 delta and a candidate lever if the composed budget needs room — **to be confirmed against
  n≥3 before any claim**, never from this single sample.
- **`ROUTE-M258-iter02-isolation-names-two-causes-not-three`** → iter-03. The ISOLATION failure text
  offers exactly two explanations and **both were refuted here**; the real mechanism (env-file
  last-wins) is not among them. A refusal that names the wrong cause sends the reader at the build,
  which is where I went first. Same family as `FIX-M257-image-listing-conflates-empty-and-unreadable`.
- **`ROUTE-M258-iter02-headroom-defaults-to-billion`** → iter-03. Bare `assert-headroom` grades against
  `billion.json`; the host must be named every time. Cluster 4's shape, in a second entry point.
- **`ROUTE-M258-iter02-purge-did-not-clear-the-stack-dir`** → iter-03. 24 accumulated blocks means the
  stack dir survived repeated `--purge` cycles. That is `verification.md`'s **F-9** instance
  (*"the purge deleted nothing"*), and it is what let history outrank intent.

**Lessons:**

- **A falsifiable assert earns its keep on the first campaign that can trip it.** ISOLATION was landed
  *with* L1 at M257 on the principle that an assert must ship with the lever that can trip it. It went
  RED on M258's first cycle and caught a demo wired to production auth — which every other signal on
  that stack graded green: `up_rc=0`, `autoverify green:true / 0 warnings`, HEADROOM OK, identity MATCH.
- **Refute BOTH named causes before believing either.** The assert's text offered two; the evidence
  refuted both and the truth was a third. Had I stopped at the first plausible one — the inherited
  `dockerignore` item, which *is* real and *does* live in this file — I would have shipped a fix for a
  defect that was not this one.
- **`git diff` before routing blame across a re-pin.** The pin moved 4 commits in the same iter the
  failure appeared. 21 insertions, 9 deletions, **all comments** — the coincidence was not causation,
  and one command settled it.
- **A file that only ever grows makes its history outrank its intent.** 23 correct blocks lost to a
  24th, because append + last-wins is a silent override.
