**Type:** tik — under [`TOK-08`](../decisions.md) (*census the mechanical classes; stop sampling them*).

# iter-152 — the probe registry has never been fenced against the platform

## Phase A — the census

`stack-verify/lib/services.sh` is the table that decides what every stack gets graded on. Its header has
always named its own source of truth — *"Source of truth: the platform's docker-compose.yml service
set"* — and **nothing has ever checked it.**

iter-145 fenced the table against its **test-side twin** (`REGISTRY_BASES` +
`test_the_test_side_registry_mirrors_services_sh`): membership, and ports through the offset sweep. That
half is closed. But both of those are **our** copies. They can agree with each other perfectly while both
disagree with the platform — which is this milestone's founding class stated three times over (skiller,
skillpath, jobsimulation each left the platform and nothing on our side noticed).

Censused at platform **`0c91421`** (`stack-demo/platform`, verified **0 behind `origin/main`** at iter
open — the ref gate clauses 1/2 were proven at).

| direction | denominator | finding |
|---|---|---|
| **forward** — does every row's literal match compose | **12 registry rows** (7 graded, 5 declared absent) | **0 wrong values.** No port has drifted |
| **reverse, per service** — does compose define a service nothing probes | **7 compose services** | **0.** The sets agree |
| **reverse, per PUBLISHED PORT** | **10 published host ports** | **3 the registry cannot express** |

**The third row is the finding, and it is the one no reading could have produced.** `§5` rules 66/69: *a
token census finds a WRONG value and can never find an ABSENT one.* The registry's denominator is
**services**; compose's is **published ports**. At `0c91421` compose publishes 10 host ports across 7
services and the registry probes 7 — one per service. So `backend` publishes `:8081`, `:8082`, `:8083`
and is graded on `:8082` alone; `studio-desk` publishes `:9000`, `:9100` and is graded on `:9000`. **A
service can be half-up — its probed port answering while another published port is dead — and verify
calls it `up`.**

That is a real limit of the probe model, not a defect in those rows (both single-port choices have a
documented rationale and both are right). What was missing is that **nothing said so, and nothing would
notice a fourth port arriving.**

## Phase B — retirement is now DECLARED, not inferred

The table's four merged-away rows were marked *"retained for older/rollback clones"* **in a prose
comment**. `D-M257x-151-1`'s lesson applies directly: a fence whose absent-row arm reads a comment cannot
fire. Two machine-readable declarations replace it, and the guard reads **those**, not the prose:

- `SERVICES_NOT_IN_PLATFORM_COMPOSE` — 5 entries (4 `merged-into-app`, 1 `rext-injected`: directus is
  emitted by `gen_override.py:directus_lines`, so compose **cannot** be its source of truth).
- `PLATFORM_PORTS_NOT_PROBED` — the 3 ports above, each with its reason.

## Phase C — `stack-core/service_registry_guard.py`

**It does NOT derive the registry from compose.** That would delete the independent copy that makes
iter-145's offset sweep assert anything at all (`§8`'s anti-vacuity rule, and the condition the prompt
attached to this route). It asserts that two independently-maintained artifacts **agree**, and names the
row and the direction when they do not.

| arm | catches |
|---|---|
| **A** departure / port-drift | a service leaving the platform; a port moving under a row that stays |
| **B** arrival, per service | a service the platform grew that nothing probes |
| **C** arrival, per **published port** | the blind class above — without it the fence inherits the very denominator it exists to fence |
| **D** declarations are honest | an allowlist that outlived its subject (a declared-absent service that came back; a declaration naming no row) |

Reading: **ALIGNED** — `12 registry rows (7 graded, 5 declared absent) vs 7 compose services publishing
10 host ports (3 declared unprobed)`. A clean negative, as iter-149 was.

**`PLATFORM_COMPOSE` is required with no default**, for the reason `platform_alignment_guard` refuses to
guess: a fidelity check that guesses its reference is one that passes against the wrong platform
(`platform-alignment.md` §4 Trap A).

### The controls, run and not assumed

`TOK-08` defines a swept class as one whose fence *"ships with a mutation control and an anti-vacuity
control that can actually fire."* **20 controls, all passing:**

- **one mutation per arm** — a service deleted from compose (A), a port moved (A), a net-new compose
  service (B), a 4th port on an existing service (C), a declared-absent service returning (D), a
  declaration naming no row (D), a declared port the platform dropped (D);
- **the arm-C isolation control** — the 4th-port mutation is asserted to produce **no** A, B or D
  finding, so arm C is proven load-bearing rather than assumed;
- **the prescribed-repair control** — the arm-A message tells the reader to declare the departure, and
  that repair is asserted to actually clear the fence (a message prescribing something that does not work
  is worse than none);
- **three anti-vacuity shapes** — empty registry, empty compose, compose with services but no published
  ports — each must exit **2 (CANNOT MEASURE)**, never 0 (ALIGNED);
- **the refusal control** — no `PLATFORM_COMPOSE` exits 2 and says so;
- **the `include:` control** — `postgresql` and `redis` are defined in the sibling `common.yml`, not in
  `docker-compose.yml`. A guard reading only the named file would report **ALIGNED with a denominator
  wrong by two**; dropping the include is asserted to turn both rows RED.

### `guard_family`'s own fence caught the registration gap

`test_the_live_census_matches_the_glob_exactly` went RED immediately: *"service_registry_guard is on disk
with no invocation — guard_family would exit 2."* Registered with `needs=("platform",)` and its reference
**derived from `--platform`**, so the three platform-facing guards cannot be pointed at different clones
by accident. `fence_provenance` then required the tree stamp; added.

## Phase D — two defects the iter's own test runs surfaced

Both are recorded here in full because each is more instructive than the clean census.

### D1 — prose in a bash string is CODE (**mine, caught by the existing fences, fixed in-iter**)

The first draft of `PLATFORM_PORTS_NOT_PROBED` **double**-quoted its entries, and one carried the prose
``... present only under `npm run dev` ...``. A backtick inside a double-quoted bash string is **command
substitution**: merely `source`-ing `services.sh` ran npm, and under `set -euo pipefail` that killed the
source and took **20 stack-verify tests** down with it.

The symptom was `npm error code ENOENT: Could not read package.json`. **It names neither the file nor the
reason**, and it appears in a *different section* from the edit that caused it — the same shape as
iter-145's `12 != 13`. Fixed by single-quoting (literal in bash); the guard now accepts **either** quote
style so a future quote change cannot silently empty a declaration. Controlled by asserting the
**property** (sourcing is inert — rc 0, nothing on stderr) rather than the fix, plus an anti-vacuity arm
(a file that sources cleanly but whose arrays the guard cannot read would pass the first arm alone).

### D2 — the grading parser read an iter's prose ABOUT grading as its grading (**pre-existing, side-deliverable**)

Running the fence-provenance controls after registering the new guard turned up
`test_the_escape_accepts_and_records` RED. **iter-132 recorded that exact test as *"an artifact of the
confound — the fence tree was DIRTY … re-run alone on the committed tree: 1 passed in 83.11 s."* That
attribution does not reproduce.** The tree was committed and clean here and it still fired — so the
falsification is **corrected in place**, not carried.

The real cause: `blocking_state_guard` located the Phase-5 grading with an **unanchored `search()`** —
the first occurrence of `**Phase 5 grading:**` anywhere in the file. `iter-150/progress.md` is the iter
that **documented this guard**, and its line 30 reads *"the guard that reads every iter's `**Phase 5
grading:**` line"*. That mention won, **68 lines ahead of the real one**, and the span ran to the next
bold key — swallowing a documented **FIXTURE**, `(8) host-quarantine: y`, as though an iter had graded it.

**Both of the guard's RED findings on the live corpus were that one phantom:** the partition arm reported
an unclassified protocol field, and the zero-claim arm counted it as the unrepresented blocking grading
that made `deferrals-audit.md`'s *"zero open user questions"* banner read as false. **The banner was
correct. The guard was reading its own documentation.**

It is the milestone's recurring class in its sharpest form — iter-149 swept 33 stale test copies a census
had counted as real emitters. Fixed by `^`-anchoring **and** last-match-wins; **both halves are
controlled, because either alone leaves a way back in** — anchoring is defeated by an iter quoting a
prior iter's grading at column 0 (these progress files do that routinely), and last-match-wins is
defeated by nothing else. Plus an anti-vacuity arm and a live arm asserting no phantom field survives on
the real corpus. `blocking_state_guard`: **RED (2 findings) → OK, exit 0**; 19 → 23 controls.

## Close — 2026-08-08

**Outcome:** the registry that grades every stack is fenced against the platform for the first time —
**ALIGNED at `0c91421`, 0 drift in both directions** — and the census produced the one thing a reading
could not: the registry's denominator is **services** where the platform's is **published ports**, so 3
of 10 published ports were unexpressible and a half-up service graded `up`. Two defects surfaced by the
iter's own test runs, both fixed: a backtick in a bash declaration string that executed npm on `source`,
and a pre-existing phantom that had `blocking_state_guard` reading iter-150's documentation of it as a
grading — which was the whole of its live RED.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged. **No `N` reading was taken, so no `N` movement is claimed**
(`§9`'s guard-rail 1, in its required words).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**iters 135–152 took no reading, so the metric is UNMEASURED not unmoved — `§9`'s iter-type refinement; `TOK-08`'s sealed refutation branch bars an agent-authored successor in any case**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (**1 tik this session**) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue** (the session had room for another iter; this field was first written `y` and is corrected here rather than left standing — `§5`'s rule that a grading is a measurement, not a mood)
**Decisions:** `D-M257x-152-1` (fence two hand-maintained copies against each other; never derive one
from the other, or the sweep that used them stops asserting) · `D-M257x-152-2` (a probe registry's
denominator must be the platform's PUBLISHED PORTS, not its services) · `D-M257x-152-3` (a corrected
falsification: iter-132's dirty-tree attribution for `test_the_escape_accepts_and_records` does not
reproduce).
**Side-deliverables:** the `blocking_state_guard` grading-parser fix (D2) — a separate commit, a
pre-existing defect, and **not** part of this iter's planned scope; it does not upgrade the close status.
**Routes carried forward:**
- `SURVEY-M257x-iter148-registry-is-hand-maintained` — **CLOSED by this iter.**
- `SURVEY-M257x-iter152-half-up-services-are-ungradeable` — **NEW.** Arm C now *declares* the 3 unprobed
  ports; it does not probe them. Whether `backend:8081` / `:8083` carry a surface whose death should fail
  a bring-up is a question for a live stack, not for a fence.
- `SURVEY-M257x-iter152-other-guards-may-read-prose-as-data` — **NEW, and the generalisation of D2.**
  `blocking_state_guard` was one unanchored `search()`. **iter-153's re-survey ran the first pass and it
  is recorded here so the next session does not repeat it:** 7 of the 31 stack-core guards read
  `knowledge/plan` (`blocking_state_guard`, `claim_twin_guard`, `claim_ledger`, `guard_family`,
  `evidence_visibility_guard`, `platform_predicate_guard`, `repair_postcondition`), and a first read of
  their module-level marker regexes finds the rest **already line-anchored or `re.M`** — `claim_ledger`'s
  structure set (`_DELIM_ROW`, `_HEADING`, `_LISTMARK`, `_BLOCKQUOTE`) and `platform_predicate_guard`'s
  compose/Makefile set all carry `^`; `evidence_visibility_guard._CITATION_RE` requires markdown link
  syntax a prose mention does not have. **Do NOT close the route on that** — anchoring is the mechanism,
  not the property. The property is *"the marker string does not occur in the scanned tree outside its
  structural position"*, and D2 proves the two come apart: the pre-fix regex was **also** a plausible
  reading of a marker until an iter wrote about it. Grade the property, and derive the guard list from
  disk rather than from this line.
- `SURVEY-M257x-iter150-partition-completeness-elsewhere` — still open, and **D2 is evidence for it**:
  the `host-quarantine` "unclassified field" that motivated it was a phantom, so the partition was never
  actually incomplete.
- Unchanged and still queued: `FIX-M257x-iter145-sha-baseline-drift` ·
  `-iter145-migrate-race-needs-a-host-postgres` · `-iter145-green-but-stale-graphql-mentions` ·
  `SURVEY-M257x-iter144-orphan-arm-is-the-residual` · `FIX-M257x-iter144-correction-vs-retraction-unfenced` ·
  `FIX-M257x-iter143-wrong-head-is-unfenced` · `-iter143-scope-derivation-by-grep` ·
  `-iter143-appending-to-the-protocol-doc-rots-the-ledger` · `FIX-M257x-iter142-value-change-articles` ·
  `-iter142-path-arm-window` · `-iter142-tier-b-underflag` · `FIX-M257x-iter135-adjudicated-live-defects` ·
  `-iter140-receipts-not-checkable-here` · `-iter140-receipt-fence` · `-iter138-anchor-rot-fence` ·
  `-iter134-fence-family-has-no-shared-predicate-layer` · `-iter133-two-fives-need-a-fence` ·
  `-iter131-predicate-sets-not-enumerated`.
**Lessons:**
0. **When two copies of a fact are fenced against each other, ask what the PAIR is fenced against.**
   iter-145 made `services.sh` and `REGISTRY_BASES` agree and the milestone read that as the class being
   closed. Two of our copies agreeing is a statement about us. The third edge — the one to the artifact
   that owns the fact — was the one that had never been drawn, and it is the only one that can catch the
   class the milestone was founded on.
1. **A census inherits the denominator you give it.** Graded per SERVICE, this registry is complete.
   Graded per PUBLISHED PORT, it covers 7 of 10. Neither number is wrong; the first is just not an answer
   to the question. `§5` rule 69's *"a census cannot find an absent value"* has a sharper form: **it
   cannot find one that is absent from its denominator, and the denominator is a choice.**
2. **A guard can be broken by being written about.** D2 is not an exotic parser bug — it is what happens
   when a tool's marker is a plain string and the tool's own documentation is inside the tool's own
   search path. Anchor structured markers to a position the prose cannot occupy, and prefer the LAST
   match, because prose comes first and the artifact comes last.
3. **A failure that names nothing costs more than the defect.** `npm error code ENOENT` (D1) and
   `12 != 13` (iter-145) are the same failure of message. D1 was found in one run because the fences it
   broke were **run**; it would have been invisible for 132 iters if they had not been.
