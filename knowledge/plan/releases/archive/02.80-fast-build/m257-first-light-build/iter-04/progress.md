**Type:** tik · **Active strategy:** `TOK-01` (step 3 — the baseline leg, re-aimed at a host that exists)

## Line 1 — Re-survey: two of the three things this milestone was paused on are stale

M257 has been paused since 2026-07-31, across the whole of M257x (288 iters). `TOK-01`'s next-tik
direction, iter-03's routing and the milestone's own gate all name **odysseus**. Phase 1 Step 0 required
re-verifying those targets before committing to them. **Two moved.**

### 1a — iter-03's `user-blocker` is ALREADY RESOLVED, by M257x

iter-03 exited `user-blocker` on `DECIDE-M257-jobsim-schema-ownership`: `repos.yml` made `app` the sole
migration owner, so a fresh stack never creates the `jobsimulation` schema, while *"~15 rext writes"*
still targeted it. Two fix shapes were offered to the user — **(i) pin `platform` stale**, or **(ii)
follow the platform's new model**.

**M257x took (ii) and shipped it**: `c0e075e` — *"the third occurrence, paid off — 12 jobsim writes
re-pointed, and a fence that names no dead schema."*

Verified rather than taken on the commit subject:

| check | result |
|---|---|
| every `CopyRows*` write target across `stack-seeding` | **all `"public"`** — zero to `jobsimulation` |
| the 21 surviving non-comment `jobsimulation.` strings | **all prod-READS** (`cmd/content-capture`, `contentsession/sourcing.go`) — authoring-time reads of production, where the legacy schema legitimately still exists |
| `jobsimulation` schema on the two live stacks | **absent on both** (`to_regclass` → NULL on `demo-2` and on the dev stack) — so the drift is real, and the re-point is what makes a fresh stack work |
| `demo-2` actually seeded | **yes** — 493 `public.job_simulation_sessions`, 5 orgs, 191 users |

And it is fenced by **mechanism, not by list**: `test_write_target_schema_fence.py` asserts
`write targets ⊆ repos_yml_schemas_to_create(repos.yml)` and *contains no list of dead schemas*, so the
next fold shrinks the legal set by itself. **21 passed.**

### 1b — B1 / B2 survived M257x intact

`test_dropped_mirror_fence` **19 passed** · `test_baseline_mirror_fence` **28 passed** (the D120
host-parameterised count, unchanged) · `go test ./seeders/... -run 'Fence|Mirror|Studio|DroppedMirror'`
**ok**. Nothing to redo.

### 1c — the host-class premise that paused the milestone is REFUTED

This is the finding that matters. `state.md` and the pause note say *"a Mac is arm64/**overlay2**"*,
*"the Mac pays no unpack leg, which is exactly what M257's L1 (~200–250 s) optimises"*, and therefore
*"M257's speed gate is un-measurable on the sanctioned hosts as written."*

**Measured on the Mac mini (M4 Pro) that D-v28-15 moved dev/test to: it runs the CONTAINERD image store
and it PAYS the unpack leg.**

`docker info` DriverStatus is `[["driver-type","io.containerd.snapshotter.v1"]]` — the same signature
`spec-notes.md` F1 used to classify odysseus, and the exact trap F1 warns about (*"`Storage Driver:
overlayfs` reads at a glance like the laptop's `overlay2`, and that reading would have been expensive and
wrong"*). Per the F4 lesson that **two agreeing weak signals are not one strong signal**, this was
confirmed with evidence of a *different kind*:

| probe | export | **unpack** |
|---|---|---|
| controlled 256 MB layer | 3.5 s | **0.8 s** |
| controlled 1024 MB layer | 14.3 s | **3.0 s** |
| the real `hiring.Dockerfile` image (4.12 GB) | 56.6 s | **19.3 s** |

The leg exists and is **size-proportional**. The generalisation was from the **retired M1 Pro laptop**
(`laptop.json`, classic overlay2 graphdriver, `unpack_s_observed: null`) to **a different machine**.

## Line 2 — the host that exists, measured → `stack-core/hostprofiles/macmini.json`

Shipped **without a `gated_baseline`**, per iter-03's `PROFILE-M257-*` routing. Every field measured
except two, each flagged in the profile itself.

| | **macmini** (M4 Pro, measured 2026-08-11) | billion | laptop (retired) |
|---|---|---|---|
| image store | **containerd — pays unpack** | containerd | overlay2 — no unpack |
| `pnpm turbo build` (hiring) | **30.4 s** | 42.6 s | 140.4 s |
| export + unpack (hiring) | **56.6 + 19.3 = 75.9 s** | 99.0 + 37.6 = 136.6 s | 74.9 + none |
| hiring image | **4.12 GB** (arm64) | 4.84 GB (x86_64) | 2.88 GB (arm64) |
| `lane_heap_measured_peak_mib` | **3116** (a floor) | 3900 | 4223 |
| derived `max_parallel_ui_lanes` | **2** | 1 | 1 |

Three things in that table are load-bearing beyond the profile:

- **`max_parallel_ui_lanes` = 2.** `overview.md` prices **L2** under billion's constraint that *"the two
  compile legs cannot run concurrently on `billion` at all."* **On this host they can.** And it is a
  *floor*: `idle_mem_mib` is an upper bound (2272 MiB with the user's two stacks resident — a true idle
  was not this iter's to take) and `lane_heap_measured_peak_mib` is a floor (10 s sampling over a 30.4 s
  compile window).
- **The arm64/x86_64 image gap has closed**, 40 % → ~15 % (4.12 vs 4.84 GB). L1's prize is proportional
  to the bytes it deletes, so scaling from the laptop's 2.88 GB would have **under**-priced it.
- **The disk budget is the VM's 51.7 GiB, not the host's 173 GiB** — a 3.3× over-statement, and precisely
  the M239-F1 ENOSPC trap. Recorded in `budget_source`.

Profile validates: loader OK, `profile_describes_host` verdict **`match`**, `assert-headroom` **OK**
(`lanes=1 max_parallel_ui_lanes=2 free=49.2 GiB load1=4.19`). `python3.12 -m pytest
tests/test_buildbench.py -q` → **109 passed in 1.10 s**, and profile discovery is *globbed*, so the new
profile is validated by being checked in (an M255-inherited item that turns out to be already done).

## Line 3 — PROPOSED re-cut gate (a PROPOSAL — deliberately NOT adopted)

Re-cutting an exit gate is planning (`/developer-kit:design-roadmap`), and it is the user's call. What
this iter owes is the arithmetic, so the decision is made on evidence rather than on the refuted premise.

**Estimated** cold cycle on macmini, by scaling billion's n=3 phase table by this host's **one measured**
image (ratio 106.3 / 179.2 = **0.593** on the UI tier; the non-UI remainder scaled 0.70–0.85):

```
UI tier      billion 436.1 s  ->  macmini est ~258.7 s
remainder    billion 230.2 s  ->  macmini est ~161–196 s
TOTAL est                          ~420–455 s        (billion n=3 p50: 666.29 s)
L1 here      2 Next images x 75.9 s export+unpack, 80–90 % recovered, + ~15 s studio-desk
                                   ~136–152 s
POST-L1 est                        ~270–320 s
```

**So on this host the 360 s cut looks reachable — plausibly on L1 alone**, with the ≤300 s stretch in
range. That is the opposite of the pause's conclusion, and it follows directly from 1c.

**Proposal:** keep `p50 ≤ 360 s`, re-point the gate host from `odysseus` to `macmini`, keep both
falsifiable asserts, and re-derive `re_scope_trigger` against macmini's own baseline once measured.
**Three conditions the user should attach**, none of which this iter can discharge:

1. **The n ≥ 3 baseline is still owed and must be measured on a QUIET box.** Everything above is either a
   single measured image or an estimate scaled from billion. This box was contended throughout (load1
   3.84–7.31) and both of the user's stacks were resident. `laptop.json` records a full cycle *refused by
   its own clause 1* at load1 10.69; publishing a contended cycle as a baseline is the number-shaped
   defect this release exists to retract.
2. **Clause 1 must be fixed for this host class before it can grade anything here.** `buildbench.py:697`
   reads `os.getloadavg()[0]` — the **macOS host's** load, across 12 logical cores — and clause 1 grades
   it against `cores` which, for a `docker-desktop-vm` profile, is the **VM allocation (8)**. Two
   different machines' units in one comparison. *(This does NOT explain odysseus's `load1 48.7`: odysseus
   is `native-linux`, where the two are the same number. The uninterruptible-sleep hypothesis already at
   `buildbench.py:349-350` remains the lead there.)*
3. **If the user prefers, the honest alternative is to re-cut the gate as a reduction against macmini's
   own measured baseline** rather than inheriting billion's 360 s target — `overview.md`'s own rule is
   that a wall-clock never transfers between hosts.

## Close — 2026-08-11

**Outcome:** The two premises M257 was paused on were re-surveyed and **both are stale**: iter-03's
architectural `user-blocker` was resolved by M257x (`c0e075e`, verified by fence, not by commit subject),
and the host-class premise that made the gate *"un-measurable"* is **refuted by measurement** — this Mac
mini runs the containerd image store and pays a size-proportional unpack leg, so **L1 keeps a substantial
price here**. Shipped the measured `macmini.json` host profile (validated, identity `match`), which also
shows this host fits **2** concurrent UI build lanes where billion fits 1. **Metric delta: none, zero by
design** — no lever touched.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n *(tik; the streak needs 3 no-prog tiks and only 2 prior tiks exist)* — (3) re-scope: n *(the trigger is "p50 > 420 s after L1+L2+L3"; no p50 exists, and the evidence moved AWAY from re-scope)* — (4) **user-blocker: y** *(the exit gate names a retired host and cannot be graded as written; re-cutting it is planning, the user's call. The measurement also CONTRADICTS a binding release decision — `D-v28-15`'s "the Mac pays no unpack leg" — and a sub-agent does not silently amend a binding decision)* — (5) cap-reached: n *(tik 1 of 5)* — (6) protocol-stop: n — (7) budget-exhausted: n — **Outcome: exit-4**
**Decisions:** see [`decisions.md`](decisions.md) (D1–D3)
**Side-deliverables:** none — every change landed inside the three planned lines.
**Routes carried forward** (Fate 3, named handlers, all → **iter-05** unless stated):
- `BASELINE-M257-macmini-n3` → the `n ≥ 3` cold campaign on a **quiet** box, with the user's stacks
  down or in a fresh slot. Fills `gated_baseline` in `macmini.json`. **Blocked on the gate re-cut** —
  a baseline is only a gate number once the gate names this host.
- `FIX-M257-load1-units-vm` → `buildbench.py:697` grades **host** `os.getloadavg()` against a
  **VM-allocation** `cores` on any `docker-desktop-vm` profile. Evidence gathered, fix deliberately not
  made (4th line — tripwire).
- `INVESTIGATE-M257-load1-48` → carried from iter-02/03, still unstarted. **Narrowed:** the units
  mismatch above is NOT its cause (odysseus is `native-linux`); and odysseus is retired, so this now
  needs re-aiming at macmini or closing as moot.
- `MEASURE-M257-macmini-true-idle` → re-measure `idle_mem_mib` with the stacks down; current 2272 MiB is
  an upper bound. Conservative in the safe direction, so not urgent.
- `PROFILE-M257-provisional-fields` → the M255-inherited item is **still open** for
  `projected_image_gib` (now provisional in **two** profiles, not one): make it a machine-declared
  `provisional_fields` list the loader surfaces. Note the *sibling* M255 item — globbing the hardcoded
  `("billion","laptop")` profile tests — is **already done** (`tests/test_buildbench.py:352-361`).
- `DOC-M257-hostclass-retraction` → `state.md`, `overview.md:213-220` and `D-v28-15` all carry *"the Mac
  pays no unpack leg."* **Deliberately not edited here:** `D-v28-15` is a binding user decision, and the
  retraction should land once, with the user's re-cut, not twice.
- All of iter-03's remaining routes carry unchanged (`FIX-M257-feedback-score-approximation`,
  `DOC-M257-studio-in-app`, `DOC-M257-prereq-gaps`, `FIX-M257-stacksnap-directus-sequences`,
  `FIX-M257-directus-coldstart-order`, `DOC-M257-autoverify-project-arg`, `DOC-M257-guide-skillpath`,
  `NOTE-M257-studio-dockerignore`).
**Lessons:**
- **A route list goes stale exactly like prose, and a PAUSE is the worst place to leave one.** Two of the
  three things blocking this milestone had moved while it sat; one was fixed by the very milestone it was
  waiting on. The Step 0 re-survey is what caught it — targeting `TOK-01`'s named next-step directly
  would have re-litigated a closed decision and measured a retired host.
- **The trap the corpus documents is the trap that bit.** `spec-notes.md` F1 warns in writing that
  `Storage Driver: overlayfs` misreads as overlay2 when the DriverStatus is the containerd snapshotter.
  A binding release decision was then made on exactly that misreading — about a *third* machine, by
  generalising from a *second* one. **A documented trap does not stop being a trap once it is written
  down**; it has to be re-run per host.
- **"A Mac" is not a host.** `D-v28-15` reasoned about a host CLASS ("arm64/overlay2") where the fact that
  mattered was a per-machine setting (the containerd image store, toggleable in Docker Desktop). The
  release's own rule — *state the environment with every number* — needs a companion: **name the machine,
  not the class.**
