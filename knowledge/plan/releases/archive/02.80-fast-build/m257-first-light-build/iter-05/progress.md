**Type:** tok (triggered) · **Authors:** `TOK-02` · **Shape:** doctrine + the stale-reference repair that makes it usable

## Line 1 — the gate could not be graded, so nothing could satisfy it

M257's `exit_gate` named **`odysseus`** — host, profile filename, and baseline, three times over — and
`D-v28-15` retired that machine on 2026-07-31, **one day after `TOK-01` was written**. This is not a
milestone that was measuring and falling short. It is a milestone where **no measurement taken anywhere
could have satisfied the gate as written**, because the gate's subject did not exist.

That is why three consecutive tiks closed *"metric delta 0, by design"*, and why the tok mechanism firing
here is correct rather than an alarm: the streak is real, and its cause is a **rotted reference**, not a
failing approach.

**The repair, and the line it does not cross.** Swapping a dead hostname for the host that exists is the
same class of work as the 44 stale citations M257x fixed at its iter-278: it changes *nothing* about what
"done" means. Changing a **target** or dropping a **clause** would be planning and is not this iter's to
make. So every target survived, verbatim in substance:

| clause | before | after |
|---|---|---|
| p50 | **≤ 360 s** | **≤ 360 s** — unchanged |
| reps | 3 consecutive cold `--purge` + `demo-up` | unchanged |
| verdict | `autoverify green:true / 0 warnings` | unchanged |
| platform edits | 0 | unchanged |
| demopatch guards | all 7 (G1–G7) | unchanged |
| HEADROOM | falsifiable, read from the sampler | unchanged (+ a units definition, below) |
| ISOLATION | falsifiable, post-build image inspect | unchanged |
| stretch | ≤ 300 s | unchanged |
| `re_scope_trigger` | > 420 s after L1+L2+L3 → escalate | unchanged, + must be re-derived once `gated_baseline` exists |
| **host** | **`odysseus` (RETIRED)** | **`macmini`** ← the only substantive change |

**And it is not a relaxation — the measurement says the opposite.** The premise that paused this milestone
predicted the local host could not host this gate at all. iter-04's arithmetic puts this box at **~420–455 s
pre-lever with L1 worth ~136–152 s here**, i.e. **≤ 360 s looks *more* reachable than the premise assumed,
not less.** Re-pointing the host makes the gate harder to dodge, not easier to pass.

**One clause gained a definition, and it is called out rather than buried.** HEADROOM clause 1 reads *peak
load1 ≤ cores − 2*. On a `docker-desktop-vm` host, `cores` is ambiguous between the VM allocation (8 here)
and the machine the `load1` sample is actually taken on (12 here), and the code currently picks the first —
so it computes a limit of **6** where the correct one is **10**. The gate text now says which quantity is
meant. **This is the one place the repair changes an accepted range**, so: it is a *units* correction, of
the same class as not comparing metres to feet, and the direction it moves the number is incidental to it
being wrong. The fix itself is `FIX-M257-load1-units-vm`, routed to iter-06 as tik work.

## Line 2 — `DOC-M257-hostclass-retraction`: "a Mac pays no unpack leg" is false on this Mac

`D121` said this retraction lands **with** the re-cut, and it does. Three sites carried the claim; all three
are corrected, each naming its machine:

| site | what it asserted | status |
|---|---|---|
| `knowledge/plan/state.md` § Hosts | *"a Mac is arm64/**overlay2** … the Mac pays no unpack leg … M257's speed gate is **un-measurable** on the sanctioned hosts"* | **RETRACTED** |
| `knowledge/plan/roadmap.md` `D-v28-15` | the same, as the decision's *"what it costs"* paragraph | **RETRACTED** (the decision itself stands — only its predicted cost was wrong) |
| `m257…/overview.md` § HOST CLASS PROBLEM | the same, as the milestone's own resume-blocker | **RETRACTED** |

The evidence is iter-04's, and the important thing about it is its **kind**:

| probe | export | **unpack** |
|---|---|---|
| controlled 256 MB layer | 3.5 s | **0.8 s** |
| controlled 1024 MB layer | 14.3 s | **3.0 s** |
| real `hiring.Dockerfile` image (4.12 GB) | 56.6 s | **19.3 s** |

**`docker info` is not admissible here and that is the whole lesson.** It reports `Storage Driver:
overlayfs` on this host, which reads at a glance exactly like the retired laptop's classic `overlay2`
graphdriver while the DriverStatus is `io.containerd.snapshotter.v1`. `spec-notes.md` F1 had written that
trap down *before* `D-v28-15` was taken, and it was walked into anyway. The two-size probe is evidence of a
different kind: the leg **exists** and **scales with bytes**, which a naming coincidence cannot produce.

What survives the retraction, stated so nobody over-corrects: billion is x86_64 and this host arm64, the
identical Dockerfile yields **4.84 GB vs 4.12 GB** — a **~15 %** gap, *not* the ~40 % the laptop comparison
implied — and **seconds measured here still do not transfer to billion**.

## Line 3 — `TOK-02` authored

Recorded in the milestone-root [`decisions.md`](../decisions.md). Class **`retry-with-evidence`**: `TOK-01`'s
ordering (instrument → baseline → levers) was never falsified, so it is repaired rather than replaced —
references re-pointed, plus a step 1.5 (**fix clause 1's units before trusting clause 1's refusal**) and an
explicit contended-measurement rule (**label it; a refusal is a result**). Next-tik direction: iter-06 =
`FIX-M257-load1-units-vm`, then `BASELINE-M257-macmini-n3`.

## Close — 2026-08-11

**Outcome:** The gate is **gradeable again** — `odysseus` → `macmini`, with every target surviving verbatim
in substance and the only substantive change being the host. The claim that paused this milestone for eleven
days (*"the Mac pays no unpack leg"*) is **retracted at all three sites**, on probe evidence rather than a
config string. `TOK-02` authored.
**Type:** tok (triggered)
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) **triggered-tok: y** — (3) re-scope: n *(the trigger is "p50 > 420 s after L1+L2+L3"; no p50 exists, and the evidence moved AWAY from re-scope)* — (4) user-blocker: n *(the gate re-cut and the baseline authorisation were both RULED by the caller; re-raising a ruled question is not a blocker)* — (5) cap-reached: n *(0 tiks so far)* — (6) protocol-stop: n — (7) budget-exhausted: n — **Outcome: exit-2 by the skill's default, EXPLICITLY OVERRIDDEN by the caller's standing instruction to author the tok and continue into tiks within this call.** Recorded both ways rather than silently as `continue`: the mechanism did fire, and the override is the caller's to make, so a later reader can see the tok fired and see who decided not to stop on it.
**Decisions:** see [`decisions.md`](decisions.md) (D1–D3)
**Side-deliverables:** none — all three lines were planned scope.
**Routes carried forward** (Fate 3, named handlers, → **iter-06** unless stated):
- `FIX-M257-load1-units-vm` → clause 1 must grade `load1` against the core count of the machine the sample
  came FROM, and **fail closed** when that basis is unknown. Evidence is complete (host `os.cpu_count()`
  **12** / engine NCPU **8** / `profile["cores"]` **8**); the fix is iter-06's first item.
- `BASELINE-M257-macmini-n3` → the `n ≥ 3` contended campaign filling `macmini.json`'s `gated_baseline`.
  **No longer blocked** — the gate names this host as of this iter.
- `INVESTIGATE-M257-load1-48` → **re-aim or close as moot**: odysseus is retired, so the 48.7 reading cannot
  be reproduced. Decide during the baseline campaign.
- `MEASURE-M257-macmini-true-idle` → `idle_mem_mib` 2272 is an upper bound (two stacks resident).
  Conservative in the safe direction.
- `PROFILE-M257-provisional-fields` → make `provisional_fields` machine-declared; `projected_image_gib` is
  provisional in **two** profiles now.
- ~~`DOC-M257-hostclass-retraction`~~ → **CLOSED this iter** (Line 2).
- All of iter-03/04's remaining routes carry unchanged (`FIX-M257-feedback-score-approximation`,
  `DOC-M257-studio-in-app`, `DOC-M257-prereq-gaps`, `FIX-M257-stacksnap-directus-sequences`,
  `FIX-M257-directus-coldstart-order`, `DOC-M257-autoverify-project-arg`, `DOC-M257-guide-skillpath`,
  `NOTE-M257-studio-dockerignore`).
**Lessons:**
- **A gate that names a dead thing does not fail — it abstains, and abstention is invisible.** Three iters
  reported *"metric delta 0, by design"* and every one was accurate; none could say *"and no delta is
  achievable"*, because the un-gradeability lived in the gate's **subject**, not in any of its clauses. A
  gate should be checked for *gradeability* before it is checked for *satisfaction*.
- **Repairing a reference is not re-planning it, and conflating the two costs iterations.** The test that
  separates them: does the edit change what "done" means? A hostname swap does not; a target or a dropped
  clause does. iter-04 escalated the whole question rather than the half that needed a decision, and paused
  a milestone on a repair.
- **A binding decision inherits the evidence quality of the probe behind it.** `D-v28-15` is sound in what it
  *decided* (which machines to use) and wrong in what it *predicted* (what they would cost), because the
  prediction rested on a config string read across two different machines. Retract the prediction, keep the
  decision — and say which is which, or the retraction reads as re-litigating the user's call.
