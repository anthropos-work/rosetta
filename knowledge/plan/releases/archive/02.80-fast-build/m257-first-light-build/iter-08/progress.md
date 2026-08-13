**Type:** tik · **Active strategy:** `TOK-02` step 3 — *take the baseline on the contended box, and label it*
· **Declared multi-step shape** (`overview.md` Phase plan): fix-and-commit → publish → campaign.

## Line 1 — the two fence REDs were a **moving label**, and the corpus's own rule names the class

iter-07 left a correct edit uncommittable behind two REDs in a file it never touched, and characterised
them as *"the guard resolves two anchors to content **neither of its own declared clone roots contains**."*
Re-measured at open, that characterisation is **wrong in the one way that matters** — and the guard says so
in its own output, which iter-07 did not read back:

```
corpus/architecture/platform-migration-status.md:121  cites app/main.go:1487        (read at origin/main@0a9370c) -> }
corpus/architecture/platform-migration-status.md:122  cites docker-compose.yml:168  (read at origin/main@766df6c) -> file has 164 line(s)
```

The content **is** in a declared clone root. It is at the root's **checkout**, and the guard graded the
root's **fetched upstream**:

| clone | HEAD (what the corpus was read at) | fetched `origin/main` | the cited line |
|---|---|---|---|
| `stack-demo/app` | `3eaadae68` | `0a9370c24` (1 merge ahead) | `:1487` = `Sender: msgsender.NewFromEnv(logger)` at HEAD · `}` at origin/main |
| `stack-demo/platform` | `0c91421` | `766df6c` (4 commits ahead) | file is **186** lines at HEAD · **164** at origin/main |

`stack-dev`'s copies of both have `origin/main == HEAD` and support the corpus exactly. iter-07 checked
**worktree and `HEAD`** in both roots — both correct — and concluded the guard was reading something
non-existent. It was reading something real that **nobody in this repo pinned**.

### Why this is a defect of the instrument and not of the prose

`read_target`'s `auto` ladder is *origin/main → HEAD → the checkout*, and it returns the **first rung that
resolves**. So an unpinned citation into a clone is adjudicated at `origin/main` **and nowhere else** — a
remote-tracking ref that `git fetch` advances. **This repo's own tooling fetches the clone set on every
bring-up** (`up-injected.sh`; the milestone's L10 lever counts *"~12 serial `git fetch`es"* per cycle). So
the ref a corpus claim is graded at moves **as a side effect of running the platform under test**, with
nobody editing either the corpus or the clone.

That is this milestone's most repeated finding, one layer over: **an instrument that lives inside its own
subject measures itself.** iter-07 found it in `buildbench.rext_root()`; here the same shape sits in the
fence that grades the corpus, and the *campaign iter-07 ran is a plausible cause of the fetch that moved
the ref.*

It is also the corpus's **own governing rule** turned on the fence that enforces it — *cite the sha, never
the moving label*. `platform-migration-status.md:115` says it in as many words: *"the sha is a pin and still
means what it meant; it is the **label** that expired, and a label that moves under a citation is how a
correct anchor becomes a wrong one without anybody editing it."* The guard was picking a moving label for a
sentence that named no ref.

### The fix is a rule the guard already had, extended by exactly one rung

`run()` has booked a finding since iter-100 *"only when the anchor names a non-construct at **every** ref
its own **block** offers."* A block that names **no** ref offers the **ladder** instead, and the same
argument applies verbatim. `ladder_alternates()` supplies the ladder's other **committed** rung, and the
acquittal arm now fires for `why in ("ambiguous", "default")` rather than `"ambiguous"` alone.

Three deliberate narrownesses, each of which is the strict direction:

- **`worktree` is NOT a rung.** An acquittal must be against a *committed* state, or the fence is
  satisfiable by editing the file under test. (`read_target` still falls back to the worktree when neither
  committed rung resolves the file **at all** — a different question, unchanged.)
- **A caller-named `CITE_REF` gets no ladder.** An operator who names a ref has made the ref the question
  (§5 rule 7).
- **Both arms, not the one that tripped.** The range-bounds arm's acquittal list ended in `None`, which
  re-reads *the same auto ladder* and therefore never reached HEAD either — the identical hole. This
  corpus's own `platform-migration-status.md:124` records the lesson being applied here: *"repairing a site
  leaves the class."*

**Proven able to fail.** 8 controls (`stack-core/tests/test_anchor_ladder_alternates_m257.py`), on a real
git fixture whose `refs/remotes/origin/main` and `HEAD` are pointed at two different commits with
`update-ref` — no checkout, no reset:

| control | asserts |
|---|---|
| the fix | correct at HEAD, rotted at origin/main → **not booked** |
| **MUTATION control** | `ladder_alternates` stubbed back to `()` → the same fixture goes **RED** (`anchor-on-closing-delimiter`) |
| negative control ×2 | rot at **both** rungs → still booked; out-of-range at both rungs → still booked |
| scope controls ×3 | a named `CITE_REF` gets `()`; `auto` gets exactly `("HEAD",)` and never `worktree`; no clone gets `()` |

**Scale, so the two sites are not mistaken for the class:** at this build **294 anchors** were being
adjudicated at `origin/main@0a9370c` and **95** at `origin/main@766df6c`. Two tripped. The other 387 were
one upstream commit away from tripping, and nothing in this repo controls when that commit lands.

**Guard result:** `OK — every resolvable anchor names a construct; 363 range citation(s) lie inside their
file`. **Not one word of the flagged prose was changed to make a fence green.**

## Line 2 — the arriving corpus edit is committed, and the tree is clean of it

`corpus/ops/demo/build-budget.md` (`5ddde288`) — iter-07's correct edit, unblocked by Line 1 rather than by
`--no-verify` (forbidden) or by bending accurate documentation to satisfy a resolver (worse). The
pre-commit fence ran and passed on its own terms:

```
repair-postcondition: 6 participating fence(s); 30 standalone; 0 site(s) reported
repair-postcondition: OK — this tree publishes no adjudicated claim the baseline did not already record
```

## Line 3 — the publish, and what the sweep found before it

`CLAUDE.md:154`'s workflow ran in full: **swept → tagged → `git push --tags`**, and the tag is
verified **on origin** (`git ls-remote --tags origin refs/tags/fast-build-m257-iter-08` →
`24bed210`) rather than assumed — rung zero, the thing M236 lost a whole iteration to.
`main` fast-forwarded `21bc5ba → 50a20a3` (17 commits).

**The sweep's scope was derived from what the tag ships** (`git diff --stat origin/main..HEAD`), not
from habit — `stack-core`, `demo-stack`, `stack-injection`, `dev-stack`, `stack-verify`,
`stack-seeding/isolation`. Result: `stack-core` **51 failed / 2238 passed** (71 min),
`demo-stack` **9 failed / 1085 passed**, `stack-injection` **329 passed**, `dev-stack` **155 passed**,
`stack-verify` **275 passed**, `stack-seeding/isolation` (Go) **ok**. Four of the six sections are
clean; both failing sections are triaged below.

**Stated so it cannot be over-read: this is not a whole-repo green.** `rosetta-extensions` carries
723 Go and 88 Python test files across 11 sections; this is a green over the sections this tag
changes, and the sections it does not change are unchanged from the tag that already passed them.

### The one failure that was mine, and it is exactly why the sweep runs

```
stack-core/anchor_construct_guard.py::ladder_alternates became executable-here and was not graded
```

`derivation_registry` enumerates every public set-returning function whose required arguments are
path-shaped, and **refuses to let a new one exist unclassified**. My new function is one. Classified
`DECLINE:policy` — it states that module's own ladder policy rather than reading anything off this
tree, and its `Path | None` argument is a *presence test*, never a file it opens. **Registering it
instead was considered and rejected**: that would make every test literal whose token set is `{HEAD}`
a candidate, in a tree where `HEAD` is one of the commonest tokens there is.

**A unit test could not have caught this.** It is a property of the tree *after* the commit — precisely
the class a completeness fence exists for, and precisely why the orchestrator's *"sweep first, then
publish"* was the right call rather than ceremony.

### The other 60, triaged rather than waved through

| cluster | n | reading |
|---|---|---|
| **gitignored stack scratch under `demo-stack/stacks/`** | large share | **2.2 GB**, incl. **36 `test_*.py`** in cloned `app` trees, which every tree-walking census scans. `git ls-files demo-stack/stacks` → **0 tracked**, so **no clone at this tag can carry it** |
| two whole-tree **literal ratchets** | 2 + subtests | `DOCSTRING_LITERAL_CEILING` 254 > 240, `TEST_MODULE_LITERAL_CEILING` 663 > 653 — **pre-existing**, and NOT raised here |
| demopatch **whole-file `pre_sha256`** vs live clones | 6 | the **anchor contract still holds** — probed: 1 occurrence, replacement absent. `demopatch-spec.md`'s own rule: *the anchor is the contract; the whole-file sha is only a baseline* |
| `test_migrate_race_live` | 3 | needs a live postgres container (`pg_isready` → *no response*); demo-1 was down |
| dangling README citation, unclassified Go section | 2 | pre-existing, inputs untouched by the tag |

**The control that makes "pre-existing" a measurement rather than a hope:** of the demo-stack nine,
**0 of their inputs** — manifests or test files — appear in `git log origin/main..HEAD`. They fail
identically at the tag that is already published.

**And the biggest cluster is this milestone's own root cause, one layer out.** iter-07's campaign ran
from the authoring copy, so its stack scratch landed *inside the tree the guards scan*. **An
instrument that lives inside its own subject measures itself** — found in `buildbench.rext_root()` at
iter-07, in the citation fence at Line 1, and here in the guard suite's own denominator. Three
faces, one shape.

**I did not raise either ceiling.** They are whole-tree ratchets, breached before this iter, and
raising a number I have not attributed is the move v2.8 exists to retract. What I did do is stop
feeding them: the measurement block came out of both new docstrings (population 258 → 254, 664 → 663)
and lives here instead, where a number is dated by construction.

## Line 4 — the pin guard refused my first launch, and it was RIGHT

```
==> ✗ FATAL: rext pin mismatch.
    the consumption clone is at : fast-build-m257-iter-08
    .agentspace/rext.tag pins   : fast-build-m257x-iter-288
```

I had written the new pin to **`stack-demo/.agentspace/rext.tag`**, reasoning that `REPO_ROOT` is the
workspace. It is not. `up-injected.sh:52` computes `REPO_ROOT="$(cd "$HERE/../../.." && pwd)"` from
`<clone>/demo-stack`, which is **three levels up = the ROSETTA root** — the same value from the
authoring copy and from the pinned clone. The SoT is **`/rosetta/.agentspace/rext.tag`**, a per-box
pin, exactly as `rext_tag.sh` documents (*"It lives in the rosetta repo's `.agentspace/`"*). My file
was read by nothing.

**Note what the two derivations do NOT share, because it is the whole of iter-07's finding:**
`REPO_ROOT` is the rosetta root either way, but the **workspace** is `$HERE/..`'s parent — `stack-demo`
from the pinned clone, `.agentspace` from the authoring copy. One derivation is location-independent
and the other is not, in the same script.

Repaired by pointing the box SoT at the published tag (previous value preserved at
`.agentspace/scratch/work-m257/rext.tag.before-iter08`) and deleting my stray second pin file — a
second pin is the drift class the SoT was created to retire.

**Cost: ~2 minutes and three fast-failed reps.** The guard printed the fix, the fix was one line, and
a stack that ran the wrong tooling would have produced a number attributed to the wrong code.

## Line 5 — `BASELINE-M257-macmini-n3`: the campaign, from the pinned clone

Launched at **load1 2.06** — the waiter held until the box fell under buildbench's own clause-1 limit,
because contention I cause is contention I can remove (iter-07 D2). It then rose on its own.

```
rep-01 total=542.94s up_rc=0 green=True warn=0 headroom=FAIL peak_load1 19.48
rep-02 total=449.51s up_rc=0 green=True warn=0 headroom=FAIL peak_load1 14.52
rep-03 total=410.80s up_rc=0 green=True warn=0 headroom=OK   peak_load1  4.82

n=3   min 410.80   p50 449.51   max 542.94   phase table COMPLETE   host identity match x3
```

**Every one of iter-07's three disqualifications is gone:**

| iter-07, from the authoring copy | iter-08, from the pinned clone |
|---|---|
| **17/17 demopatches REFUSED** | `✓ demo-patches: all applied (none refused, none skipped)` |
| `postgres-schemas` probe could not find `repos.yml` | resolves; probe runs |
| `autoverify green:False` ×3 | **`green:true / 0 warnings` ×3** |
| phase table INCOMPLETE | **COMPLETE** |

**So the remedy iter-07 verified against the predicates but could not take is confirmed by
execution.** The blocker was never contention, disk, memory or the host — it was *where the harness
was invoked from*, and one publish plus one re-pin removed it.

### The sub-phase table — the first lever-pricing data this host has ever produced

| sub-phase | p50 (s) | range (n=3) | share of 449.51 |
|---|---|---|---|
| `ui_next_web` | **120.79** | 115.71–130.58 | 26.9 % |
| `ui_hiring` | **117.45** | 108.41–172.93 | 26.1 % |
| `set_dress` | 81.61 | 81.60–109.13 | 18.2 % |
| `compose_up` | 44.43 | 44.02–49.74 | 9.9 % |
| `host_preflight` | 33.79 | 32.88–34.61 | 7.5 % |
| `backend_builds` | 16.26 | 3.40–34.77 | 3.6 % |
| `ui_studio_desk` | 7.99 | 7.86–10.62 | 1.8 % |
| `autoverify` · `clones_and_inject` · `secrets_provision` · `seed_tooling` | 2.26 · 1.83 · 1.79 · 1.29 | | 1.6 % |

**The UI tier is 246.23 s = 54.8 % of the cycle** — L1's target, and the shape billion showed at
65.5 % holds here at a smaller share. **`set_dress` at 81.61 s is L5's**, and it is bigger than
`compose_up`. The two Next lanes are within 3 s of each other, which is what makes L2's
`max_parallel_ui_lanes = 2` on this host interesting rather than academic.

### What this number is, exactly

**`gated_baseline` is FILLED** (`macmini.json`) — the first on this host, and `TOK-01`'s standing
rule (*no lever may be priced until it is*) is discharged.

**It is CONTENDED AND LABELLED, and the label waives nothing.** 2 of 3 reps failed HEADROOM clause 1,
so the campaign exits **RED** by contract (`D-M255-1`) and this is not a gate pass. `TOK-02` step 3
authorises exactly this: *record the refusal with what the run would have measured, rather than
reporting a failure to measure.* Both numbers are in the profile, per rep, with their `load1`.

**rep-03 is the one fully-clean cycle this host has produced** — 410.80 s at peak load1 **4.82**,
green, headroom OK, every clause satisfied except that it is n=1.

### Distance to the gate, and it moved in the right direction

| | seconds | cut needed for 360 s |
|---|---|---|
| `billion` n=3 p50 | 666.29 | **46 %** |
| iter-07's degraded, unpinned anchor | 489.90 | ~27 % |
| **`macmini` n=3 p50, gated configuration** | **449.51** | **19.9 %** |
| rep-03, the headroom-clean cycle | 410.80 | 12.4 % |

**iter-04's ~420–455 s estimate was right** — the measurement lands inside its range, at 449.51.
It was still not a measurement, and this milestone spent four iters proving why that distinction is
not pedantry. The `re_scope_trigger` is **re-derived** against this baseline (420 → **400 s**), with
its arithmetic stated in the gate line so the next reader can redo it.

## Line 6 — the box was left as it was found

`demo-1` torn down after rep-03 (plain `down 1`, images kept — the next campaign purges at rep start
anyway). The user's stacks were never touched and are exactly as they were at open: the 5-container
dev stack and `demo-2`'s 11 containers, all still up. The box SoT pin now names the published tag;
its previous value is preserved at `.agentspace/scratch/work-m257/rext.tag.before-iter08`.

## Close — 2026-08-11

**Outcome:** **`BASELINE-M257-macmini-n3` LANDED — `gated_baseline` is filled: p50 449.51 s (n=3,
min 410.80 / max 542.94), all three reps `rc=0` + `autoverify green:true / 0 warnings` + all
demo-patches applied + a COMPLETE phase table**, driven from the pinned consumption clone. Every one
of iter-07's three disqualifications is gone, which confirms by execution the remedy iter-07 could
only verify against predicates. **Contended and labelled: 2 of 3 reps failed HEADROOM (peak load1
19.48 / 14.52 vs 10), so the campaign exits RED by contract and this is not a gate pass** — rep-03
(410.80 s at load1 4.82) is the one fully-clean cycle. Getting there took the two things iter-07
raised and could not take: the arriving corpus edit unblocked by **fixing the guard, not the prose**
(the fence was grading unpinned citations at a `git fetch`-moved `origin/main`), and
`rosetta-extensions` **swept, tagged and pushed to origin**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n *(449.51 s p50 > 360 s; and 2 of 3 reps carry a headroom refusal)* — (2) triggered-tok: n *(tik; and it moved the primary metric from "no baseline exists" to a filled `gated_baseline`)* — (3) re-scope: n *(the trigger reads a p50 after L1+L2+L3; no lever has landed — and it was re-derived this iter, 420 → 400 s)* — (4) user-blocker: n *(both of iter-07's blockers were ruled on and executed; nothing new needs a user decision)* — (5) cap-reached: n *(tik 1 of 5)* — (6) protocol-stop: n — (7) budget-exhausted: **y** *(~2 h 15 m: a 71-minute publish sweep plus two full campaigns; the iter closed and committed cleanly between iters)* — **Outcome: exit-7**
**Decisions:** see [`decisions.md`](decisions.md) (D1–D3)
**Side-deliverables:**
- `stack-core/derivation_registry.py` — `ladder_alternates` classified `DECLINE:policy` (the sweep's one genuinely-mine finding).
- `overview.md` `re_scope_trigger` **re-derived** against the now-filled `gated_baseline` (420 → **400 s**), with its arithmetic stated in the line. Caught by `test_baseline_mirror_fence` rather than remembered — the fence flagged the milestone's own stale `~420–455 s` estimate the moment the profile carried a real number.
**Routes carried forward** (Fate 3, named handlers):
- **`LEVER-M257-L1-multistage-next`** → the opening lever, now priced against real data: the UI tier is **246.23 s = 54.8 %** of the cycle and the gate needs only **89.51 s**. `ASSERT-M257-isolation-with-L1` lands *with* it (`TOK-01`: the falsifiable assert ships with the lever that can trip it).
- **`FIX-M257-sweep-scratch-pollutes-census`** → 2.2 GB of gitignored stack scratch under `demo-stack/stacks/` (36 `test_*.py` in cloned `app` trees) is walked by every tree-scanning census, and is a large share of the 51 stack-core REDs. Not in any tag (`git ls-files` → 0). Needs the censuses to exclude the scratch dir, not the scratch dir deleted — `demo-2` is the user's live stack.
- **`RATCHET-M257-literal-ceilings-breached`** → `DOCSTRING_LITERAL_CEILING` 254 > 240 and `TEST_MODULE_LITERAL_CEILING` 663 > 653, pre-existing and deliberately **not raised** here. Either attribute and raise with a reason, or pay the debt down.
- **`FIX-M257-demopatch-sha-baselines-drifted`** → 6 whole-file `pre_sha256` baselines stale against live clones. Every **anchor** contract still holds (probed), so this is the self-healing freshness gate doing its job; re-pin the baselines.
- iter-07's routes carry unchanged: `FIX-M257-campaign-kill-orphans-bringup`, `FIX-M257-io-sampler-macos`, `FIX-M257-sampler-disk-units-vm`, plus iter-05/06's tail (`MEASURE-M257-macmini-true-idle`, `PROFILE-M257-provisional-fields`) and iter-03/04's.
- ~~`FIX-M257-anchor-guard-resolution`~~ → **CLOSED this iter** (Line 1).
**Lessons:**
- **A guard that reads `origin/main` grades the corpus at a ref `git fetch` moves** — and the fetch is done by the tooling under test. The corpus's own rule (*cite the sha, never the moving label*) applies to the fences that enforce it, and an unpinned sentence names no ref, so no ref may be picked on its behalf.
- **iter-07 checked the right files and the wrong refs, then wrote a conclusion the guard's own output contradicted.** The guard printed `read at origin/main@0a9370c` on the line above the finding. **Read the instrument's provenance field before theorising about the instrument.**
- **A completeness fence catches what a unit test structurally cannot** — `ladder_alternates` broke `derivation_registry` by *existing*, and nothing scoped to the diff would have seen it. That is the argument for sweeping before publishing, and it paid for the 71 minutes on its own.
- **"Pre-existing" is a measurement, not a defence.** The demo-stack nine were graded by checking whether any of their inputs appear in `git log origin/main..HEAD` — **zero do** — so they fail identically at the already-published tag. That check takes one command and turns a hand-wave into evidence.
- **The pin guard was right and I was wrong about a path I had already read.** `REPO_ROOT` is `$HERE/../../..` — the rosetta root from *either* clone — while the *workspace* is the clone's parent, which differs. One script, two derivations, one location-independent and one not; that asymmetry is the whole of iter-07's finding and I re-learned it by tripping the guard that exists to catch it.
