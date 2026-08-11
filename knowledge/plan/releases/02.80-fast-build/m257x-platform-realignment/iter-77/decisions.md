# iter-77 — decisions

## `D-M257x-77-1` — free prose is NOT fenceable, and the corpus should NOT be restated to make it so. Fence the predicates that are decidable without reading a sentence.

The briefing posed the design question at the centre of every strategy revision since TOK-02:

> *Can free prose be fenced at all — or should the corpus state those 21 claims in a form that CAN
> be checked?*

**Both framings are rejected, on measurements rather than on preference.**

**Against fencing the prose.** Three candidate prose fences were built and run against the live
corpus. Membership-by-line reaches **71% precision** (7 hits, 5 true, 2 false — `service_taxonomy.md:52`
lists the legacy *schemas*, `ops/README.md:44` narrates the fold's version history). Scoping it to
the clause holding the `migrations:` token removes both false positives **and 2 of the 5 true
defects**, because the corpus writes *"so both dropped to `migrations: false`"* and the subject is a
pronoun — **100% precision, 60% recall**. Membership-by-citation-cell was rejected on
`storage.md:25`. Neither 71% precision nor 60% recall is a fence you can leave RED, and widening a
rule until it reads English is §4 Trap A in its purest form.

**Against restating the corpus.** The restatement move is untried and superficially attractive, and
this iteration is the argument against it: **the 21 unreached claims were never the failure.** The
one claim G5 could reach was being read **wrong** (`D-M257x-77-3`), so its effective reach was
**0 of 24, not 1**. Restating 21 prose claims into enumerable form would have moved the denominator
and left the instrument reading the enumerated form incorrectly — a corpus rewritten to fit an
instrument that does not work. **Fix the instrument before reshaping the artifact it reads.**

**What is implemented instead** — the third answer, and it is not a compromise between the two:
fence the predicates that are decidable **without reading a sentence at all**, and leave attribution
honestly UNREACHED with its count printed on every run.

| decidable from | assertion |
|---|---|
| the platform artifact's **git history** | the repo vocabulary — `repos_yml_history` |
| the artifact's **line numbers** | `repos.yml:A-B` citations — **G9** |
| the **tree** (a filename) | G9's subject, exactly as G8 takes its own |

None of the three reads English. All three are facts a `git` command answers.

**The residual is not deferred and not re-cut.** G5's 20 free-prose claims stay UNREACHED, named,
and counted on every run. That is a *declared* limit of a working instrument, not a hidden one — and
`D-M257x-76-2` already binds every future quotation of this guard's GREEN to its reach line.

---

## `D-M257x-77-2` — inherited work is evidence, not fact. Re-derive it, including a prior session's own numbers.

A prior run of this iteration was killed mid-Phase-D and left +129/−3 uncommitted lines carrying a
docstring that asserted **"4 of those RED"**. Continuing from it would have been the natural move.

Re-derived, **two of the four are false**, and both fail for reasons this milestone has already
written down:

- `jobsimulation.md:12` sits in a block pinning `2adcf71`, where `repos.yml:17-19` **is** the
  jobsimulation block — §5 rule 33, and the *identical* mechanism iter-76 booked as its own first
  adjudicated false positive one class earlier.
- The docstring described `roadrunner.md:14` and `:29` as citing the same lines for the same reason.
  They cite **different** lines for **different** reasons, and `:14`'s exempting pin `87d8d44` does
  not resolve in the platform clone at all.

The prior run also shipped no assertion body — the docstring described G9 in full while the loop
contained only a dead `doc_subject` binding. **A described guard is not a guard.**

**Rule:** an inherited diff is read and re-derived like any other claim, and its stated measurements
carry no more weight for having been written by the same milestone. The briefing's *"six of the
orchestrator's numbers have failed re-derivation — re-derive everything here, including mine"*
extends to work this milestone left on its own disk.

---

## `D-M257x-77-3` — a derived vocabulary must be HISTORICAL, or the instrument goes blind exactly where the drift is.

`vocabulary` was `set(repos) | set(compose.services)` — derived entirely from what **currently
exists**. A repo therefore left the vocabulary at the same commit it left `repos.yml`, and every
corpus claim naming it became unreadable **at the exact moment it became false.**

Not theoretical. `setup_guide.md:486` enumerated the `migrations: true` repos as *"(currently: app,
cms, jobsimulation …)"*; `_names_in` resolved that to `{'app'}`, compared `{'app'} == {'app'}`, and
**passed a false claim**. The single migration claim of 24 that G5 could reach was the one it read
wrong.

The vocabulary is now the union of every `- name:` ever written in `repos.yml`, from the platform
artifact's own history — **14 ever, 6 now, 8 removed**, re-derived independently of the guard and in
exact agreement. A clone too shallow to answer reports `UNMEASURED` and falls back to the current
set; it never pretends.

**Generalises past this guard, and belongs in the protocol:** *any fence whose vocabulary is derived
from current state is blind to removals, and removals are the drift.* The same shape would hide a
deleted compose service, a deleted profile token, a deleted env var. Where an artifact is under
version control, derive the vocabulary from its **history**.

---

## `D-M257x-77-4` — a precision rule is checked AFTER the verdict, never before. Order decides whether it costs recall.

G9's ambiguous-subject rule (the citing block names a platform repo other than the document's own
subject) is correct and necessary — `roadrunner.md` carries a line-cited paragraph about
`jobsimulation`, and grading it against the filename subject is a false positive.

Placed **before** grading, it cost `storage.md:25` — a correct, gradeable citation whose sentence
merely says *"exactly as `cms` and `jobsimulation` are kept"* — and G9's graded count fell **4 → 2**.
A passing mention blinded the fence to a true claim.

Placed **after**, it can only ever fire on a citation that would already have been RED, so it
converts would-be false positives into declared UNREACHED and **spends no recall at all**.

**Rule:** a discriminator that narrows a fence is evaluated only on the branch that was about to
report a finding. Anywhere earlier it silently removes true claims from the denominator, and the
reach line — which is the only place that loss would show — reports it as *"nothing to check."*

---

## `D-M257x-77-5` — "committing is not pushing", and it belongs beside rung zero.

The milestone's pre-flight rung zero is *tagging is not publishing* (M236 lost an iteration to a tag
that existed only in a local authoring copy). This session opened on the same failure one step
earlier: `rosetta-extensions` `main` held **13 commits that existed on exactly one disk** — 8
hardening passes and 5 fences, ~1,400 lines of guard work, unpushed.

Nothing a stack consumes was at risk (the consumed tag was on origin), which is precisely why it
survived thirteen commits: **the pin guard checks what a stack pulls, and nothing checks what the
author has not pushed.** A pin verified on origin is silent about the branch it was cut from.

Recorded in [`corpus/ops/platform-alignment.md`](../../../../../corpus/ops/platform-alignment.md)
beside rung zero. Pushed and verified (`git ls-remote origin main` == local `main`) as this
session's first act.

---

## Routed forward

- **`CHECK-M257x-iter77-narration-vs-documentation`** — G1's noun-phrase construct reads *"the line
  that stood here named the `graphql` profile"* as documenting the token. It has a negation
  discriminator but none for **historical narration**. Latent today (masked by the leftmost-pin
  exemption); it becomes live the moment anything widens that exemption. Measured surface: 2 sites,
  26 blocks in the affected class.
- **`CHECK-M257x-iter77-cross-repo-pin`** — **145** pin-exempted blocks name a sha that does not
  resolve in the platform clone. Overwhelmingly legitimate `app`-repo citations, which is why G9's
  resolve-here rule is scoped to the one assertion whose subject is a platform file and was **not**
  widened into `_pin_exempts`. Unmeasured whether any of the 145 dates a *platform* claim with a
  foreign sha.
- **`CHECK-M257x-iter77-zsh-modifier`** — `git show "$SHA:file"` is mangled by zsh's `:r` modifier
  and returns an **empty set with RC=0** at the pipeline level. Any derivation shelling out in this
  shell must brace its expansions; an empty result from a mangled command reads exactly like a
  finding.
- **`CHECK-M257x-iter77-developer-dir`** — mid-session, `xcode-select -p` flipped from
  `/Library/Developer/CommandLineTools` to `/Applications/Xcode.app/…`, whose licence is unagreed.
  **`git` and `python3` both stopped working at once**, each exiting **69** with a licence message
  on stdout — so a test run's output file contained one line of prose and no verdict, and a suite
  that had passed an hour earlier reported **4 failures and 24 errors** (every fixture that shells
  out to `git`, including this iter's new ones). It reads exactly like *"my change broke
  everything."* `export DEVELOPER_DIR=/Library/Developer/CommandLineTools` restores both,
  per-process, with **no `sudo` and no system state changed** — worth knowing before anyone accepts
  a licence to fix a test suite. Belongs with the environment traps in `verification.md`.
- Unchanged and re-affirmed: `FIX-M257x-iter53-union-set` (**PENDING USER DECISION**) ·
  `FIX-M257x-iter56-assignment-flake` (**NOT DECIDED**) ·
  `CHECK-M257x-iter76-compose-service-count` (8 vs 9 vs 10, **explicitly unsettled**) ·
  `CHECK-M257x-iter76-seat-ref-discipline` · `CHECK-M257x-iter70-studio-room-lines` ·
  `RF-M257x-iter71-run-returns-a-tuple` · `CHECK-M257x-iter38-ai-act-classification` (owner outside
  this milestone) · RF-2/3/7–13.
