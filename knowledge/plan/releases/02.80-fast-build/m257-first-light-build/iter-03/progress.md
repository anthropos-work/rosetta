**Type:** tik · **Active strategy:** `TOK-01` (step 2 → 3)

## Work

### B1 — the dropped `local_*` mirrors: the blast radius was **20 files / 34 sites**, not six

The "six failing seeders" were accurate but were only *the six surfaces that happened to fail first*.
Two live sites sat **outside** them, and one was decisive:

- 🔴 **`stack-verify/live/autoverify.sh` — the gate's own hiring assert READ the dropped table**, wrapped
  in `|| echo 0`. **Fixing only the six seeders would not have cleared the warning:** the query would
  still error, the error would still be swallowed, "0 sessions" would still be reported as
  *under-production*, `green` would stay `false` and READY would stay unsatisfiable. **A missing TABLE
  and an empty table read identically through that pipe** — which is why the warning no longer asserts a
  single cause. *The instrument this milestone spent iter-02 fixing was itself a casualty of the drift
  it was meant to detect.*
- Two `resetTables` entries in `cmd/stackseed/main.go`, masked because the reset probes `to_regclass`
  first and **skips a missing relation in silence**.

**The re-point is row identity, not a rename.** `public.job_simulation_sessions` is the *jobsim-in-app*
table and is **column-identical to `jobsimulation.sessions`**, so the seeders write the same row at the
same id — which is what the migration did (`js_session_id` remapped mirror-id → canonical id). Four
`NOT NULL` columns had no mirror source and were **handled, not papered over**:

| column | resolution |
|---|---|
| `sim_type`, `token` | derived with the **identical expressions** `JobsimSessionsSeeder` uses for the same key. A wrong `sim_type` is the **G14 inserted-but-invisible** class (`intelligence.go:1547` filters on it) — an error would be better than this. |
| skill-path `version` | the canonical table carries a `UNIQUE (user_id, skill_path_id, version)` **the mirror did not**, and `CopyRowsIdempotent` only guards `ON CONFLICT (id)` — a shared version would abort the live seed on a different index. Two new distinct tokens, following shipped precedent. |
| `duration` | `progress*60`, matching sibling seeders. |

**And the migration turned a latent bug into a hard FK violation:** `assignments.go` and
`hiring_funnel.go` were both pointing `session_id` at `deterministicUUID(key+":spsess")` — **an id
nothing ever wrote.** Harmless while unenforced; a violation now that `organization_assignment_sessions`
FKs `skill_path_sessions(id)` with CASCADE. Both now write the real row.

`anticheat_summary` / `anticheat_tagline` have **no canonical home** — they moved to
`public.anticheat_results`. rext never wrote them, so nothing was lost, but the **documented scrub
surface was wrong and would now pass vacuously.** Corrected.

### B2 — `app/studio`: the premise was never true, for `cms` either

`app` needs **exactly** `cms`'s tree, and that is **verified, not inferred**: `app`'s own CI passes
`additional_repo: "anthropos-studio-room:studio"` to colony's build workflow, and its `.gitignore` says
so. Acquisition is now **one fetch, independent trees** — forced, because the two images build from
different contexts and a docker context cannot reach outside itself — with the **consumer set derived
from each clone's Dockerfiles rather than hardcoded**, so the next service to grow the dep is covered
with zero tooling edits and a pin too old to need it is never fetched for.

Two corrections to the iter's own brief, both load-bearing:

1. **`app` DOES have `.gitmodules` — at `v1.360.0`–`v1.362.x`.** `fdb8034a` added it plus a gitlink;
   `851cf3fb` (2026-07-29) deleted both **while keeping the Dockerfile `COPY`**. True for `main` today,
   false for those tags. **This is why the guard tests for `requirements.txt` and not `-d`:** `git
   checkout` materialises an uninitialised gitlink as an **empty directory**, which a `-d` guard reads
   as *present*.
2. **`ensure-clones.sh:144-145`'s premise was never true, including for `cms`** — `cms/Dockerfile.dev`
   hard-`COPY`s `studio/` and `pip install`s from it, so a missing tree aborts that build just as hard.
   The non-fatal disposition was **always** wrong. Now FATAL for both, gated on the Dockerfile.

### Both fences were proven to fail

- **Dropped-mirror fence** (19 tests) **scores the occurrence, not the line** — comments explaining the
  dead table survive, executable strings do not. *A fence that forbade naming the table would have
  deleted its own rationale.* Mutation-verified: re-adding the string to `resetTables` goes red and
  names the file:line.
- **Studio fence** (15 tests) walks **every `COPY` in every injected Dockerfile** and classifies each
  source as tracked-in-repo (rejecting mode `160000` gitlinks), a `go build -o` output, or
  tooling-acquired — verified by calling **the tooling's own predicate**, not by matching the string
  "studio". Run against `app@main`'s **real** Dockerfile with the acquisition path removed: **RED, and
  it names it.** The wiring half goes red on 5 tests against the true pre-fix files.

## 🔴 A strictly LARGER blocker surfaced, and its fix shape is a genuine architectural choice

Platform `repos.yml` @ origin HEAD (`236771f10`, **2026-07-29T14:14Z — 39 minutes after the
mirror-drop**) sets `migrations: false` for **cms, jobsimulation, roadrunner**, stating: *"`app` is the
ONLY repo with migrations to run… they own no local schema and are not part of the stack."*

**So on a stack whose `platform` clone is at main HEAD, the `jobsimulation` schema is never created —
and rext still writes ~15 `jobsimulation.*` tables** (`JobsimSessionsSeeder`, `PersonaSeeder`,
`ContentStorySeeder`, `SuccessionSeeder`). That is **B1's shape, one layer up and much wider.**

**Why it did not bite the observed odysseus run:** `persona_write.go`'s flush writes
`jobsimulation.sessions` **first** and the mirror **fifth**, and personas failed on the *mirror* — so
that stack's `platform` clone **predates** `236771f10`. It is latent, not absent.

**This was routed rather than guessed, deliberately** (iter-03's own escalation condition). The fix
shape is a real choice with release-scope consequences, and it is not an agent's call:

- **(i) Pin `platform`** to a pre-2026-07-29 ref → reproducible immediately, but the stack is
  knowingly stale and drifts further from production every week.
- **(ii) Follow the platform's new model** — accept `app` as the sole migration owner and re-point ~15
  `jobsimulation.*` writes to their canonical `public.*` homes → correct and durable, but it is a second
  B1 of unknown size, and B1 alone was 34 sites across 20 files.

## Close — 2026-07-31

**Outcome:** B1 and B2 both landed — the two blockers that made READY unsatisfiable on **every** host —
each with a mutation-proven fence. **Metric delta: none, zero by design** (no lever touched). The two
remaining planned deliverables are routed forward for a **sequencing** reason, not merely for budget:
both need the host running **this** tooling, and the host currently runs the iter-02 tag plus a
hand-applied hack. A `load1` figure measured against stale tooling would not be the figure the gate
reads.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n *(tik; no 3-no-prog streak — 2 tiks have run and both landed planned scope)* — (3) re-scope: n *(the milestone's trigger is "p50 > 420 s after L1+L2+L3"; no p50 exists)* — (4) **user-blocker: y** *(the `repos.yml` migration-ownership change makes the gate unreachable again, and choosing between pinning `platform` and following the new model is an architectural decision with release-scope consequences that changes what code lands next — Phase 5 §4's first positive example)* — (5) cap-reached: n *(tik 2 of 5)* — (6) protocol-stop: n — **Outcome: exit-4**
**Decisions:** see [`decisions.md`](decisions.md) (D1–D3)
**Side-deliverables:**
- 12 `stack-seeding` tests that were **red mid-work**, fixed.
- A pre-existing slice bug in `demo-stack/tests/test_tooling.py::test_explicit_ant_academy_clone_step`
  — it sliced `(d2)`→`(e)`, so it silently asserted about **every phase in between**.
- `demo-up-defaults.md`: 13 `file:line` citations refreshed via `demo_knob_guard.py --fix` (a hard gate,
  red without it). **Anchors only** — diff verified to carry zero prose changes.
- `content-stories-routes.md`: the MIRROR-trap retraction **plus** a second stale claim at `:289` (the
  anonymization scrub surface located the anticheat free-text on the dropped mirror).
**Routes carried forward** (Fate 3, named handlers, all → **iter-04**):
- `DECIDE-M257-jobsim-schema-ownership` → **the user-blocker above.** Blocks everything.
- `INVESTIGATE-M257-load1-48` → carried from iter-02, **unstarted by design**: must be measured with
  `buildbench`'s own sampler against a host running **this** tooling.
- `PROFILE-M257-odysseus-json` → author `hostprofiles/odysseus.json`. Host facts are measured;
  `lane_heap_measured_peak_mib` is **not yet** and must not be guessed — clause 2 consumes it. Ship it
  **without** a `gated_baseline` (the fence handles a baseline-less profile by design).
- `FIX-M257-feedback-score-approximation` → `feedback.go` uses a `>=55` pass band where
  `JobsimSessionsSeeder` applies a `passThreshold` nudge, so the app-side score can differ from the
  `jobsimulation.sessions` row **of the same id**. Pre-existing; benign in a mirror, **not** benign in a
  table claiming to be the same row. `st.Activity.PassRate` is in scope at the call site — a small exact
  fix, deliberately not made (a data-value change, out of B1's remit).
- `DOC-M257-studio-in-app` → the corpus says studio-room is embedded in **CMS only** in five places
  (`service_taxonomy.md:57,158`, `architecture_overview.md:23`, `ai_architecture.md:54`,
  `setup_guide.md:310`). Nothing records that `app` embeds it too. → `/update-knowledge`.
- `DOC-M257-prereq-gaps`, `FIX-M257-stacksnap-directus-sequences`,
  `FIX-M257-directus-coldstart-order`, `DOC-M257-autoverify-project-arg` → all carried unchanged from
  iter-02.
- `DOC-M257-guide-skillpath` → `demo-stack/GUIDE.md:89` still names `skillpath` among "the 4
  Clerk-consuming Go services"; it was decommissioned at v2.7 M246 and `INJECT_SVCS` is 3.
- `NOTE-M257-studio-dockerignore` → `app`'s `.dockerignore` does not exclude `studio/.git` (2.2 MB), so
  it enters the context and the image. Parity with CI was kept deliberately rather than diverging — but
  in a release named *fast build* it is worth a line.
**Lessons:**
- **The drift had already reached the instrument.** iter-02 fixed `autoverify` so the gate could be
  trusted; B1 then found that `autoverify`'s own hiring assert was reading a table the platform had
  dropped, and that `|| echo 0` was converting a schema error into a plausible-looking `0`. **A
  swallowed error is worse than a missing check**, because it produces a number that looks like data.
- **"The six that failed" is not "the sites that are broken."** Six surfaced because they ran first; the
  seventh was the assert, and the eighth was silent behind `to_regclass`. **Ordering determines what a
  single failing run reveals** — enumerate by search, never by symptom.
- **Twice now, a premise held for four days because nobody walked the path.** B2's *"not a hard build
  dep"* was never true even for `cms`, and B1's mirror had been dropped for two days. Both were found
  within hours of a cold cycle. The corollary for the remaining routes: **the `repos.yml` finding is
  latent only because one clone is stale** — it will bite the first fresh box, exactly as B2 did.
