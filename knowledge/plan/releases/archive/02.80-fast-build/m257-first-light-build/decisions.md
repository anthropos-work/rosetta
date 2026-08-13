# M257 — decisions

_Release-level binding decisions **D-v28-1 … D-v28-14** live in
[`../../../roadmap.md`](../../../roadmap.md) § Active — v2.8._

---

## USER-BLOCKER-2026-07-31: Phase 0b KB-fidelity gate returned RED — bootstrap tok blocked

**Status:** open · **Surfaced to user:** yes · **Blocks:** iter-01 strategy authoring

`/developer-kit:build-mstone-iters` Phase 0b runs the KB-fidelity audit **before** the bootstrap tok
authors strategy, precisely so the first strategy is built against verified knowledge. It returned
**RED**. Per the skill, a RED is blocked, not proceed-with-caution: *"Fill blind areas … before
proceeding."* No `iter-01/` dir was created — the gate fired ahead of Phase 1, so the milestone has
**zero closed iters** and the next invocation re-enters at the bootstrap tok.

Full report: [`kb-fidelity-audit.md`](kb-fidelity-audit.md). Host recon: [`spec-notes.md`](spec-notes.md).

### Why this is a real stop and not audit pedantry

The RED is not a scatter of stale line numbers. **M257's own declared `iteration_protocol_ref`
(`corpus/ops/demo/build-budget.md`) asserts in six places that the v2.8 gate is measured on `billion`
and that `666.29 s` is the number every reduction target is measured against.** D-v28-14 superseded
both **today**. A bootstrap strategy authored against that doc would price levers against a baseline
the release has just disowned — which is the exact defect class v2.8 exists to retract, committed in
the milestone's own protocol doc.

Three load-bearing claims were re-verified first-hand rather than taken from the audit
(see `spec-notes.md` for the table). All three confirmed.

### The three items that need a USER decision, not a default

These are **not** blind-area backfills an agent should quietly pick a side on — each edits a contract
that M255 shipped nine days ago.

1. **The baseline mirror fence vs odysseus's baseline — a mechanical blocker, and the sequencing
   question.** `stack-core/tests/test_baseline_mirror_fence.py` pins `hostprofiles/billion.json`'s
   `gated_baseline.total_p50_s` as **THE single source**, and explicitly fences **M257's own
   `overview.md`** (`:53`) — any line asserting a 3-digit seconds value in the **640–700 s band**
   without `666` on it reads as DRIFT (`:83-110`). Writing odysseus's measured p50 into prose is
   exactly what this milestone must do, and if it lands in that band the fence goes red.
   **Decide:** parameterise the fence by gate host (two sources, one per host), or retire billion's
   `gated_baseline` in favour of odysseus's, or scope the fence to historical sites only.
2. **How far to retract the six "billion is the gate host" claims.** The audit cleanly separates
   *historical-and-still-true* (where the M255 baseline **was** measured — `build-budget.md:108`,
   `:132-133`, `:439-440`) from *forward-looking-and-now-false* (`:139`, `:230`, `:240`, `:261`, plus
   `billion.json:4,13` carrying it in **machine-readable** form). **Decide:** rewrite in place now, or
   fold into the §8.5 retraction already scoped to this milestone so `frontend-tier.md` and
   `build-budget.md` are each rewritten **once**, with achieved numbers, as `overview.md:100-113`
   intends.
3. **`build-budget.md:150-152` is an orphaned instruction.** It says *"re-confirm on the first
   post-freeze campaign"* for three timing-derived claims, against a freeze that has expired on a host
   that is no longer available for the purpose. **Decide:** re-home the re-confirmation to odysseus's
   baseline campaign (it would ride along at no extra cost), or close it as superseded.

### What the agent can fix without asking (ready to execute on resume)

Recorded so the next invocation does not re-litigate them: re-anchor the §8.5 site list to the four
**live** claims (`frontend-tier.md:255`, `:273`, `:274`, `:286` — the overview's `:231/:249/:262/:271`
are pre-M255 and one is already retracted) and `up-injected.sh:816`/`:1251`; document the harness's
undocumented surface (13 flags, `BUILDBENCH_PROFILE`/`BUILDBENCH_LANES`, the `gateable` field);
correct `buildbench.py:22`'s stale *"~11.6 GiB per rep"* against the measured *"+1.7–2.2 GB"*; add
`name` to the profile loader's required keys and glob the two tests that hardcode
`("billion","laptop")` so `odysseus.json` ships validated rather than untested; document clause zero
(`require_measured`) and the `min`/`max(1,…)` clamps missing from the `max_parallel_ui_lanes` formula
at `build-budget.md:293-294`.

### Not blocked — the host recon already landed

The provisioning path is **clear and needs no user input**: the prereq list `roadmap.md:321` points at
is genuinely present and specific (`tailscale-serve.md:119-131`). Odysseus is reachable, is a
containerd-image-store host like billion (**so L1 keeps its full ~200–250 s price** — the audit's
highest-value unknown, resolved in the gate's favour), and carries two newly-measured risks the audit
did not have: **it is a truly-cold box** (0 images / 0 build cache → rep 1 is the variant D-v28-8
excluded, so a discarded warm-up cycle is mandatory) and **it has no swap** where billion used
2,452 MB of it at peak. See `spec-notes.md` F1–F4.

## D120 / D121 / D122 — the three pre-flight contract calls (coordinator, 2026-07-31; user delegated this class)

### D120 — PARAMETERISE the mirror fence by host. Do not retire billion's baseline.

`test_baseline_mirror_fence.py` pins `billion.json` as THE single baseline source and fences M257's own
`overview.md`: any 3-digit seconds claim in the **640–700 s band** without `666` on the line reads as DRIFT.
Writing odysseus's p50 into prose is this milestone's job, so **the fence I wrote at the M255 close now
blocks the work it was written to protect.**

**Decision: parameterise by host — the fence reads every `hostprofiles/*.json` and checks a prose claim
against the profile whose host the line NAMES.** A baseline-shaped claim that names **no** host FAILS.

Rejected: *retire billion's `gated_baseline`* — billion's 666.29 s is historically real, cited in eight
places, and the record of what the release started from; deleting it to make room is destroying evidence.
Rejected: *scope the fence historically* — weaker than what exists.

**Why parameterising is the stronger fence, not a loosening:** it promotes this release's own standing rule —
***state the environment with every number*** — from prose convention into a machine check. The rule exists
because M255 measured the same Dockerfile at **4.84 GB on billion** and **2.88 GB on an arm64 laptop**. A
fence that demands the host be named enforces exactly that, and would have caught the M257 gate's own
un-hosted `666.29 s` inheritance before a human noticed.

### D121 — Minimal correction NOW, full rewrite with achieved numbers at close.

Six forward-looking "billion is the gate host" claims are live (`build-budget.md:139,230,240,261` +
`billion.json:4,13`, the last two **machine-readable**). Two options were on the table: retract now, or fold
into the §8.5 retraction already scoped here so each doc is rewritten **once**.

**Decision: both, split by kind.** A **minimal** correction now — the host, and that billion's baseline does
not transfer — because (a) the audit RED is *blocking the milestone* and must clear, and (b) leaving six
false host claims live means anyone reading the protocol doc mid-milestone aims at the wrong machine. The
**achieved-numbers** rewrite stays at close, where it was already scoped, so prose is not rewritten twice.

Rationale for not doing it all now: a half-retracted doc set is its own defect class — `demopatch-spec.md`'s
chain rule exists because a two-stage rewrite left one site reading DRIFTED against a pristine file *by
design*. One correction now, one rewrite at close, no intermediate state that is wrong in a new way.

### D122 — Re-home the orphaned re-confirmation to odysseus's baseline campaign.

`build-budget.md:150-152` orders three timing claims re-confirmed "on the first post-freeze campaign" — on a
host we can no longer use. **Decision: re-home to odysseus's baseline campaign.** It rides along free: the
campaign already runs n ≥ 3 cold cycles with per-phase attribution, which is what those claims need.

Closing them as superseded was the alternative and is wrong: they were explicitly marked for
re-confirmation *because nobody had confirmed them*, and a host change does not discharge that — it just
changes where.

### And a compound risk the two findings make together (F3 × the M256 inheritance)

**F3: odysseus has ZERO swap**; billion has **16 GiB** (`billion.json:21`) and used **2,452 MB** at peak. The headroom clause still
fits on arithmetic — `1×3900+1500 = 5400` MiB against `0.8×7780 = 6224` — but on billion a transient
overshoot met *swap*, and on odysseus it meets the **OOM killer**.

**Which lands exactly on `FIX-M256-demo2-service-self-termination`, inherited into this milestone.** An
OOM-killed service and a service that self-terminates `Exited 0` on a DB-health monitor present
**identically**: containers "Up" in `docker ps`, surfaces rendering empty, **no error anywhere**, and it once
cost an hour of misdiagnosis. So on odysseus a headroom breach will not announce itself as memory pressure —
it will look like the silent-empty symptom this milestone already owes a fix for. **Fix the liveness check
first; it is the instrument that tells those two apart.**

---

## TOK-01: instrument before baseline, baseline before levers — 2026-07-31

**Tok type:** bootstrap (iter-01)

**Initial strategy:** Do **not** touch a lever until three things are true, in this order:
**(1)** odysseus can run a cycle at all, **(2)** the gate's own instrument is *proven able to fail*,
and **(3)** odysseus's own `n ≥ 3` p50 baseline exists and is checked into
`stack-core/hostprofiles/odysseus.json`. Then price levers **largest-measured-second first**, one per
iter, re-measuring at `n ≥ 3` after each, and **land each falsifiable assert together with the lever
that can trip it** — never after.

**Rationale:**

- **The distance to the gate is unknown, and that is the whole opening problem.** The `666.29 s`
  baseline is `billion`'s. `D-v28-14` makes billion demo-only, and this release's own standing rule is
  *state the environment with every number*. odysseus is a near-twin **on paper** — 8 cores / 7,780 MB /
  x86_64 Linux 6.8 against billion's 8 vCPU / 7.3 GiB / x86_64 — which makes 666 s *plausible* there and
  proves nothing. M255 measured what host difference does to identical inputs: the same Dockerfile and
  context yield **4.84 GB on billion** and **2.88 GB on an arm64 laptop**. So `360 s` is a release
  **thesis**, not a derived number, until step (3) lands. **No lever may be priced against a percentage
  of a number measured on a machine this milestone cannot use.**
- **The instrument comes before the measurement because this gate READS the instrument's verdict.**
  Clause 2 is literally `autoverify green:true / 0 warnings`. M256 just found **43 checks that reported
  success without having checked**, including a probe runner certifying *"all live probes passed"* over
  **zero probes** whose exit code four other gates read as health. A gate nobody has watched fail is not
  a gate. **So: prove `autoverify` can go RED before trusting its green.**
- **And the two inherited M256 fixes make this gate lie in BOTH directions**, which is why they are
  step (2) and not parked: `FIX-M256-demo2-service-self-termination` lets `green / 0 warnings` **PASS**
  on a stack where two services sit at `Exited 0` (surfaces render empty, no error anywhere — it cost an
  hour of misdiagnosis); `FIX-M256-autoverify-fapi-libressl` makes a **working** stack emit the warning
  this gate counts. Until both are fixed the gate's verdict is unfalsifiable in the way this release
  keeps finding.
- **The compound risk forces the liveness check to the FRONT (coordinator, 2026-07-31).** odysseus has
  **zero swap**; billion has **16 GiB** and used **2,452 MB** of it at peak on the same measured lane. The
  headroom clause still fits on arithmetic (`1×3900 + 1500 = 5400` MiB against `0.8×7780 = 6224`), but on
  billion a transient overshoot met *swap* and on odysseus it meets the **OOM killer** — and **an
  OOM-killed service is indistinguishable from an exit-0 self-termination**: containers "Up", surfaces
  empty, no error. Every lever in this milestone pushes on memory. Without the liveness check, the first
  thing an over-aggressive lever produces is a symptom I would misdiagnose for an hour, and the
  misdiagnosis would read as *"my lever broke the stack"* — the same trap `build-budget.md` records for
  a mid-campaign ENOSPC presenting as `redis exited (1)`.
- **Largest-second-first, one lever per iter, because attribution is the deliverable.** L1 and L3 both
  rewrite the same Dockerfiles; landing them together would conflate their attributions and neither
  number would be quotable. The re-scope trigger is arithmetic on individual lever prices, so a blurred
  attribution disarms the trigger too.
- **L1 is the opening lever and it is already de-risked *for this host*.** It is proven real (M255:
  hiring image `4.84 GB → 379 MB`, export leg `146.8 s → 2.9 s`) but proven **on billion**, and its price
  depends on a host fact that was unrecorded until today: whether the containerd `unpacking to …` leg is
  paid. **Recon F1 confirms odysseus is a containerd-image-store host** (`Storage Driver: overlayfs` but
  `DriverStatus = io.containerd.snapshotter.v1` — the containerd store on its overlayfs *snapshotter*,
  the same class as billion, **not** the laptop's classic overlay2 graphdriver). So L9's 85.7 s is real
  here and **L1 keeps its full ~200–250 s** instead of losing ~86 s. Had that gone the other way,
  L1+L2+L3 would fall to ~215–265 s against the 306 s the gate needs and a **re-scope signal would have
  been on the table before a single lever was touched**.

**Strategy class:** `new-direction` — the milestone's first strategy; no prior approach to compare against.

**Distance-to-gate context:** **UNKNOWN and unmeasurable today** — this is the defining fact of the
opening. Gate: **p50 ≤ 360 s** over 3 consecutive cold `--purge` + `demo-up` cycles reaching READY on
**odysseus**, plus `autoverify green / 0 warnings`, 0 platform-repo edits, all 7 demopatch guards, and
the two falsifiable asserts (HEADROOM, ISOLATION). Stretch ≤ 300 s. **If** odysseus lands near billion's
666 s, the need is ~306 s and the big three are worth **200–250 (L1) + ≤45 (L2) + ~55 (L3) = 300–350 s** —
i.e. **at L1's conservative end they miss the gate by themselves.** There is no cushion; L4/L5/L7/L8/L10
(~93–158 s) are load-bearing, not garnish. **`re_scope_trigger`: p50 still > 420 s after L1+L2+L3** —
escalate, do not grind.

**Known-context carried from the Phase 0b YELLOW** (the strategy must account for these):

1. **`verification.md:623-626` misattributes the cause of the exit-0 defect** — it reads as *"autoverify
   cannot see this"*, but `stack-verify/lib/services.sh:43-44` **does** carry `jobsimulation` + `cms`
   rows, so a **re-run would have gone red**. The M256 stack stayed green via the **stale-verdict**
   class, not a blind check set. **Scope the fix to the real hole** — `fake-fapi`/`fake-bapi` have **no
   `services.sh` row** (the 16-vs-13 gap) — and fix that paragraph. Building the fix the doc implies
   would waste an iter.
2. **The LibreSSL false-positive may not fire on odysseus at all** — it is a *macOS host `curl`* defect
   (LibreSSL cannot handshake the mkcert leaf). odysseus is Linux/OpenSSL. **Determine empirically in
   iter-02**, do not assume in either direction: if it does not fire there, it still threatens the
   laptop path and still owes its one-sentence doc fix, but it is not a gate risk.
3. **`verification.md` under-enumerates the check set it documents** — 8 of ~13 cheap-wins; **(g)
   studio-desk appears 0 times in the whole file**; demo-only gating and the `DEMO_NO_*` skips are
   undocumented. The gate reads this check set, so its denominator matters.
4. **The fence now imposes an ordering**: prose may quote odysseus's p50 only **after**
   `odysseus.json` carries it, or the un-hosted/unmatched claim FAILS. **Measure → write the profile →
   then write prose.**
5. **Clause zero (`require_measured`) fires first on a fresh host** — `unmeasured_disk_avail_gib` with
   an unpopulated sampler. Expect it on odysseus's first assert; it is not a bug.
6. **Provisioning:** Go **1.26.5 is installed** at `/usr/local/go/bin/go` but **not on PATH** → fix
   PATH, do **not** install over it. **atlas is genuinely absent** → install. Confirm the six rext Go
   modules build against 1.26.5 (prereq list says 1.25.x) rather than assuming newer-is-fine.
7. **odysseus's Docker is completely empty** → rep 1 measures the **truly-cold** variant `D-v28-8` cut
   from the gate. **A discarded warm-up cycle is mandatory**, not prudent.
8. **The stack's pinned `stackseed` can be older than the authoring copy** — M256 added three `Persona`
   fields, so a `--reset` with a stale binary **truncates the world then fails to re-seed, leaving it
   EMPTY**. Shadow the authoring build on `PATH`.
9. **Rung zero:** odysseus clones rext **from origin at a pinned tag**. A tag that exists only locally
   is unreachable to it — M236 lost a whole iteration to exactly that. `git push --tags` is part of
   shipping.

**Next-tik direction (iter-02):** **Make odysseus a bench, and the instrument falsifiable.** Three
deliverables, no lever and no gated number: **(a)** provision the host per
`corpus/ops/demo/tailscale-serve.md:119-131` (PATH-fix Go, install atlas, ssh-agent, snapshot cache),
verifying the rext modules build; **(b)** land both inherited `autoverify` fixes — a container-liveness
cheap-win scoped to the **real** hole per known-context #1, and a fapi probe independent of the host TLS
stack, *after* establishing empirically whether (d) even warns on Linux; **(c)** **prove `autoverify` can
go RED** — a deliberate negative control (stop a service, confirm the verdict flips and names it), because
this gate consumes that verdict directly. Expected lift on the primary metric: **zero, by design.** The
iter grades on its planned deliverables — a host that can run a cycle and an instrument that can fail —
per the Phase 4 Step 0 rule that *planned scope* is what the iter's `overview.md` committed to.

---

## TOK-02: same strategy, a host that exists — and take the number on the box we actually have — 2026-08-11

**Tok type:** triggered (3-no-prog streak — iter-02, iter-03, iter-04 all closed *"metric delta 0, by design"*)

**Prior strategy:** `TOK-01` — *instrument before baseline, baseline before levers*. Three preconditions in
order: **(1)** the host can run a cycle at all, **(2)** the gate's own instrument is proven able to FAIL,
**(3)** that host's own `n ≥ 3` p50 baseline is checked in. Then levers, largest-measured-second first, one
per iter, each landing with the falsifiable assert that can trip it.

**Why it stopped working:** *the shape was right and every reference in step 3 was dead.* Steps 1 and 2
landed and are banked. Step 3 named `odysseus` three times over — the host, the profile filename
`odysseus.json`, and the baseline itself — and **`D-v28-15` retired odysseus one day after `TOK-01` was
written**. So did the exit gate, which is why iter-04 could not simply proceed: a gate naming a retired host
is not merely inconvenient, it is **ungradeable**, and no measurement taken anywhere could have satisfied it.

The three no-prog tiks were not idle. iter-02 proved the instrument falsifiable, iter-03 fixed the two
blockers that made READY unsatisfiable **on every host**, iter-04 measured the host that exists. Each
delivered a precondition; none could move a metric that did not exist yet. **A metric that cannot move for
three iters is exactly the condition this mechanism exists to catch** — and what it caught here is that the
strategy's *targets* had rotted, not its *logic*.

**New strategy — `TOK-01` with its references repaired and one premise inverted:**

1. **The gate names the machine that exists.** Re-pointed `odysseus` → **`macmini`** (this iter). Every
   target survived verbatim in substance: p50 ≤ 360 s over 3 consecutive cold `--purge` + `demo-up` cycles at
   `autoverify green:true / 0 warnings`, 0 platform-repo edits, all 7 demopatch guards, both falsifiable
   asserts, the ≤ 300 s stretch, and the re_scope_trigger's semantics. `TOK-01`'s **step-3 discipline is
   unchanged and re-pointed**: no lever may be priced until `macmini.json`'s `gated_baseline` is filled by a
   real `n ≥ 3` campaign.
2. **Fix the instrument's UNITS before trusting its refusal.** `buildbench.py` clause 1 grades **host**
   `os.getloadavg()` against `profile["cores"]`, which for a `docker-desktop-vm` profile is the **VM
   allocation**. Here that is a 12-core machine's load against an 8-core limit → a threshold of **6** where
   the correct one is **10**. It **fails CLOSED**, so it would refuse cycles this host is fine to run, and a
   refusal I could not tell from a real one is worse than no clause. `FIX-M257-load1-units-vm` comes *before*
   the campaign, not after — this is `TOK-01`'s own *"prove the instrument before the measurement"* rule
   applied to the clause `TOK-01` never got to exercise.
3. **Take the baseline on a CONTENDED box, and label it.** The host is a permanently-busy workstation
   (observed load1 ~2.9–13) and it cannot be freed; **waiting for quiet is waiting forever**, which is what
   `laptop.json` did — it records a full cycle *refused by its own clause 1* at load1 10.69 and no cycle
   number at all. The release's binding rule is *state the environment with every number*, and that rule is
   satisfied by **labelling** a contended measurement, not by declining to take one. Every figure carries its
   load1; none is published as a clean baseline. **A clause-1 refusal is a RESULT** — record it, with what
   the run would have measured, rather than reporting a failure to measure.
4. **Then levers, unchanged**: largest-measured-second first, one per iter, re-measured at `n ≥ 3`, each
   landing together with the falsifiable assert that can trip it. **But re-priced for this host**, which is
   where the premise inverted: L1 is worth ~136–152 s here (not the ~0 the pause assumed), and the derived
   `max_parallel_ui_lanes` is **2** where billion's is 1, so L2 changes character rather than merely shrinking.

**Strategy class:** `retry-with-evidence` — this is `TOK-01` retried against a host that can actually be
measured, unblocked by evidence gathered since (iter-04's unpack probe and host profile, and M257x's
`c0e075e`). It is deliberately **not** `new-direction`: the ordering discipline was never falsified, only its
targets went stale, and discarding a sound strategy because its nouns rotted would lose the one thing three
iters bought.

**Distance-to-gate context:** **still formally UNKNOWN, and that is now the only thing between here and
levers.** Gate: p50 ≤ 360 s on `macmini`. No cold cycle has ever been run on this host. iter-04's estimate —
~420–455 s pre-lever, with L1 worth ~136–152 s → ~270–320 s post-L1 — is a **scaling of billion's phase table
by one measured image** and must never be quoted as a baseline; this milestone's opening lesson is that a
number scaled from another machine is not a measurement. What *has* changed is the sign of the estimate: the
premise that paused this milestone predicted the gate was unreachable locally, and the measurement points the
other way.

**Cross-refs to prior TOKs:** `TOK-01` sequenced instrument → baseline → levers and **that ordering stands**;
it stalled only because its host was retired between authoring and execution. This tok does not pivot away
from it — it repairs its references, adds a step-1.5 (fix the clause's units before trusting the clause), and
records that its central risk assessment was **inverted by measurement**: `TOK-01` reasoned that if odysseus
had turned out *not* to pay the unpack leg, *"a re-scope signal would have been on the table before a single
lever was touched."* That is precisely the reasoning `D-v28-15` then applied to the Mac — on a `docker info`
string rather than a probe — and it was wrong.

**Next-tik direction:** **iter-06** — land `FIX-M257-load1-units-vm` (clause 1 grades load1 against the core
count of the machine the sample came FROM; fail-closed when that basis is unknown, and prove the new arm RED
with the precondition absent), then open `BASELINE-M257-macmini-n3`: the `n ≥ 3` cold
`demo-down --purge` + `demo-up` campaign on a **free** demo slot, every rep labelled with its load1, filling
`macmini.json`'s `gated_baseline`. Heartbeat before disturbing the user's `demo-2` or the dev stack; prefer a
slot that is already free.

## HARDEN-M257-1 — the final harden pass's routed handler + the two it re-affirmed — 2026-08-12

Recorded at milestone-root level so `/developer-kit:audit-deferrals` sees them; the pass's full findings
are in `hardening-ledger.md` (Pass 1).

**NEW, routed forward — `FIX-M257-dockerignore-env-pattern-unpaired`.**
`demo-stack/frontend/next-web.dockerignore` excludes `.env*`. Docker matches `.dockerignore` patterns from
the **context root**, so that rule covers `./.env` and nothing nested — while every other rule in the file is
deliberately paired with a `**/` twin. Consequence, measured on the real post-L1 image: `apps/web/.env`
ships into the runner carrying the platform's real Clerk publishable key and `CLERK_SECRET_KEY`, with no
`.env.local` beside it, and standalone's `server.js` **loads it at boot** — it is not an inert build
leftover. The Clerk vars are masked by `gen_injected_override.py` setting them explicitly plus `@next/env`'s
never-overwrite rule; the residual is the set difference.

**Why it was not fixed inline (three-fate rule, Fate 3 with a named handler).** The one-line repair
(`**/.env*`) also excludes `apps/web/.env.local` — the overlay carrying the **minted** key into
`next build` — so it would bake the **real** Clerk key: the M218 iter-03 incident, re-created by its own
fix. A correct repair needs a re-include and a real build to validate, which a harden pass cannot do
without a bring-up. The net under it is iter-09's ISOLATION clause: a bundle carrying a non-minted key is
the `foreign_pk` arm, so the naive repair reds the campaign rather than reaching a presenter.

**Corrected, not deferred — `FIX-M257-committed-env-ships-real-clerk-pk`.** Its description says
*committed*. The file is untracked (`git ls-files 'apps/*/.env*'` returns only `.env.example`), so a fixer
follows it into the platform repo, finds no tracked file, and concludes the finding was mistaken. The
handler is superseded by the one above, which names the real mechanism in a file this repo owns.

**Re-affirmed, unchanged — `FIX-M257-anchor-guard-content-drift`.** Still Fate 3: a new detection mode on a
pre-commit fence grading the whole corpus's citations, late in a milestone. The harden pass did not land it
and did not widen it; it made the limit executable (`TheContentDriftBlindSpot`, which goes RED when the
handler lands and hands over its fixture) and moved the statement of the limit into the guard's own
docstring, where an auditor reads it.

**Not deferred and explicitly not raised — the two literal ratchets.** `DOCSTRING_LITERAL_CEILING` and
`TEST_MODULE_LITERAL_CEILING` are breached against ceilings of 240 / 653. The breach is pre-existing and
whole-tree. This pass's own prose added to both and that growth was **deleted**, returning each to exactly
iter-09's closing reading, per the standing rule that a ceiling one has not attributed is never raised.

## HARDEN-M257-2 — the stack-core failure roster, attributed — 2026-08-12

The orchestrator's standing instruction was to **re-verify** iter-09's attribution of the whole-tree
`stack-core` reading rather than inherit it. Done, against iter-09's own sweep log on disk.

**8 of the roster were fixed by this harden pass, and none of them was about this repo's code.**
`test_suite_census_collection.py` (6) and `anchor_construct_guard`'s two live-tree arms were RED before
the pass and are green after — the first from a demo's persistent platform clone leaking into the census
population, the second from five inverted range citations. Both are recorded in `hardening-ledger.md`.

**The largest remaining cluster now has a named cause, and it is the BOX.** `test_suite_census.py`'s four
`TestBucketsFire` arms plus `test_suite_census_population.py`'s health arm fail because `run_one`'s
**unittest** column cannot import a test module under the interpreter `pytest` resolves to on this box.
Measured across all three interpreters present, on a real module:

    /usr/bin/python3            3.9.6    OK
    homebrew python@3.12        3.12.13  ModuleNotFoundError: No module named 'tests.<mod>'
    /opt/homebrew/bin/python3   3.14.6   OK

`*/tests/` carries no `__init__.py` anywhere, so those directories are namespace packages, and only 3.12
declines to resolve `tests.<mod>` from a section root. `suite_census`'s own docstring is headed *"THERE ARE
TWO INTERPRETERS"* and names 3.9.6 and 3.14.6 — both still true individually — while `stack-core/README.md`
instructs *"run it with the `pytest` entrypoint"*, and that entrypoint is a **third** interpreter neither
document names. So the suite is routinely driven by the one configuration in which the census's second
runner reports an ImportError for everything.

**Routed forward — `FIX-M257-census-interpreter-namespace-import`.** Fate 3, not Fate 1: the three candidate
repairs (section root on `PYTHONPATH` inside `run_one`; switch the arm to `unittest discover`; add
`__init__.py` across every `tests/`) each change either what this census MEASURES or what pytest treats as a
rootdir, and picking between them is a decision about which interpreter is canonical for this repo. It is
also not M257's code — no iter of this milestone touched `run_one`. The finding is recorded in the module's
own docstring, where the stale two-interpreter heading was.

**Attribution verdict:** the residual `stack-core` failures do not name any file this milestone authored.
The milestone's own test files were re-run individually and are green; `stack-injection` is 329 passed /
8 skipped; the five Go sections are 44 packages with 0 FAIL, uncached.

---

## CLOSE-M257 — the close's own decisions + the deferral re-audit — 2026-08-12

### D-M257-C1 — the §8.5 retraction reached a site the plan never enumerated, and the fence is keyed to the CLAIM, not to a site list

`D-v28-10`'s work list named `frontend-tier.md` ×4 + `README.md` + `CLAUDE.md`. Two of those cites were
wrong on arrival (`CLAUDE.md:318` is a Clerk bullet; the real row is `:444`), all four `frontend-tier.md`
cites had drifted **+61 lines**, and grepping the claim strings found a **seventh** live site the
enumeration never contained — `tailscale-serve.md`'s *"the next-web build spikes to ~3.7 GB"*, the last
live copy in the corpus. Two more (`verification.md`, `demopatch-spec.md`) carried the *"~3 min build"*
magnitude inside arguments that survive but whose number does not.

**Decision: the fence greps the CLAIM across a scoped file set, and a hand-written site list is never the
gate.** `stack-core/tests/test_section85_retraction_fence_m257.py`. A site list is a snapshot of where a
claim was **on the day someone looked**; the claim is what propagates. (`§5` rule 54.)

**And the fence is POSITIONAL, not absence-based** — the design call worth recording. Every retraction in
this corpus quotes what it retracts; `build-budget.md:10` opens by quoting all three stale numbers on
purpose. A "these strings must not appear" fence is therefore unsatisfiable, and satisfying it would trade
*being wrong* for *being silent* — silence being what let these survive four releases. The rule is: a
retracted claim may appear only where a reader is being **told** it is retracted (marker on the line or
within three preceding non-blank lines). It caught two of my own sentences immediately, plus a
case-sensitivity hole in its own marker list.

### D-M257-C2 — one of the four "retracted" claims is NOT retracted, and reading it settled it

The plan lists *"pure memory starvation, not a slow build"* among the claims to retract, on M255's
evidence that a cold cycle spent 288.4 s in export/unpack on a box under no memory pressure.

**Decision: it is half true and wrongly EXCLUSIVE, and only the exclusivity is retracted.** Swap-thrashing
on an undersized VM holding a second stack is real and reproducible; the build was *also* genuinely slow,
independently of memory. Both were true and the sentence asserted one. L1 removed the I/O half, which makes
the 12 GB prerequisite matter **more** now, not less. Deleting a true observation to satisfy a work list
would have been a worse outcome than the stale sentence.

**The related and larger finding:** the *"~3.7 GB"* beside it was an **image size quoted as a memory
figure**, and the measured per-lane heap peak is per-host — 3,116 / 3,900 / 4,223 MiB. That band
**brackets** 3.7 GB, which is why it survived: approximately right, for a reason nobody had checked.

### D-M257-C3 — the gate instrument's own review found three fail-opens, and the gate was re-graded rather than assumed

Cross-cutting review over M257's 12 rext commits. All three must-fixes were in code this milestone
authored, and each is the milestone's own lesson turned back on it: clause 3 fell back to the **host**
filesystem on a VM profile (the two-machines substitution `load1_core_basis` exists to refuse, with the two
clauses of one assert holding opposite policies); `isolation_ok` was computed in `build_report` and read by
nothing, so `buildbench report <dir>` printed `gateable: true` over a directory with no isolation block;
and the RED reason string enumerated six causes, none of them isolation.

**Decision: fix all three, then RE-GRADE the three gate-met reps under the fixed code before believing the
headline.** `rep_is_ok` True ×3, `ok`/`gateable` true, p50 **286.99** unchanged, identity `match` ×3. The
fallback was never exercised in those reps (they recorded ~65 GiB, the VM's own figure, not the host's 173),
so the fix does not retroactively disturb the gate — **and that was measured, not assumed.**

### D-M257-C4 — the fixture that hid finding 2 had hidden the same class one iter earlier

`_ledger` calls itself *"a rep ledger in the shape `run_campaign` really writes"* and its docstring narrates
how omitting `host_identity` let the aggregate ignore that field for a whole iter. It then omitted
`isolation` for exactly as long. **Decision: repair the FIXTURE, not the assertion** — including the
identity control named *"…so the clause is not an identity"*, which would otherwise have been "fixed" by
weakening the very control it exists to be.

## Deferral re-audit at close (Phase 1b) — 2026-08-12

**Verdict: YELLOW.** No escape-hatch deferral, no unfated item, and one genuine repeat-deferral pattern
found and half-discharged.

**Landed during the milestone** (no fate needed): `FIX-M256-demo2-service-self-termination` ·
`FIX-M256-autoverify-fapi-libressl` (both iter-02, the two inherited from M256) ·
`FIX-M257-seeders-local-mirror-drop` + `FIX-M257-app-studio-acquisition` (B1/B2, iter-03) ·
`PROFILE-M257-odysseus-json` → satisfied as `macmini.json` (iter-04) · `DOC-M257-hostclass-retraction`
(iter-05) · `FIX-M257-load1-units-vm` (iter-06) · `BASELINE-M257-macmini-n3` (iter-08) ·
`LEVER-M257-L1-multistage-next` + `ASSERT-M257-isolation-with-L1` (iter-09).

**Landed at the final harden but never recorded as discharged** — corrected here so the audit does not
carry a fixed item forward: **`FIX-M257-sweep-scratch-pollutes-census`** was fixed by harden pass 1
(`ca9baff`, the shared ephemeral-clone predicate), which described the fix without naming the token it
closed. *(This close then found that predicate matched an **absolute** path substring — any checkout under
an ancestor named `stacks` published a census of zero — and that the comment claiming "a shared predicate
now" was declaring a unification it had not performed. Both fixed.)*

**DROPPED, with reason:** `INVESTIGATE-M257-load1-48` — the 48.7 reading was taken on `odysseus`, retired
by `D-v28-15`, and is **un-reproducible**. Narrowed twice before dropping (the units mismatch is not its
cause; a load1 far above core count was later observed on a *second* host, which weakens the
odysseus-specific reading). The surviving hypothesis is recorded at `buildbench.py:349-350`, and the
companion suspicion — *"was clause 1 ever actually asserted?"* — is answered: it was, on `macmini`, six
campaigns' worth. **Superseded:** `FIX-M257-committed-env-ships-real-clerk-pk` by
`FIX-M257-dockerignore-env-pattern-unpaired`; the original said *committed*, the file is untracked, and a
fixer would follow it into the wrong repo.

### ⚠️ The repeat-deferral pattern, named: FOUR items inherited from the M255 close (2026-07-28)

M255 routed four items to M257 with the explicit rationale that *M257 is the milestone that actually
exercises each of them*. **M257 landed none of them across nine iters.** That is the pattern the audit
exists to catch, so it gets an explicit per-item fate rather than a fourth silent carry:

1. **`_manifest_lists` body extraction — LANDED NOW (Fate 1).** Measured first: the old
   `text.find("\n}\n")` rule returns the **identical** offset as a correct next-function-bounded parse on
   the shipped script, and the tail between them is 0 lines / 0 manifest declarations. So the defect was
   **latent, not live** — nothing published was ever computed from a short body. That is exactly why it
   became a *fence* rather than a correction: the answer was right and **nothing was keeping it right**.
   Body now bounded by the next top-level function; `_SHORT_BODY` records a disagreement rather than
   absorbing it; +4 tests including a mutation control that plants a column-0 `}` in a copy of the real
   script, and a refusal control for an unterminated function.
2. **`demo_knob_guard` anchor-fence mutants — Fate 3 → M258, with a MEASURED reason replacing the
   original one.** The item read "add mutants for the anchor comparison and the `--fix` regenerator". I
   wrote them, and they cannot be attributed: `test_demo_knob_guard.py` is **RED before any mutation** in
   the battery's staged tree (measured: 27 passed in the real tree; **8 failed / 14 skipped** in a staged
   copy, because the staged tree has no `up-injected.sh` sibling and no rosetta root). So the blocker is
   not a missing entry — **it is that the battery's staging cannot host this guard's tests at all**, and a
   mutant that "goes RED for whatever unrelated reason measures nothing" (`§5` rule 53, which is what
   flagged it). The mutants were **reverted** rather than left passing-for-the-wrong-reason. Routed with
   this finding attached, which is a better statement of the work than the original.
3. **`run_campaign` rep-body coverage — Fate 3 → M258, partially discharged.** Real coverage now exists and
   arrived from a different direction: iter-229's identity work drives `run_campaign` end-to-end under a
   faked `pre_rep_assert` / `host_facts` / `demo_env_snapshot` / `docker_system_df` set. What the M255 item
   named and is **still** unproven end-to-end is the **staleness** and **dead-sampler** paths. Reported as
   partial rather than closed.
4. **`PROFILE-M257-provisional-fields` — Fate 3 → M258.** Still not machine-declared; `projected_image_gib`
   is provisional in **two** profiles now. Partially mitigated at this close: `macmini.json` gained an
   explicit `notes.build_shape` recording that its `lane_heap_measured_peak_mib` and `projected_image_gib`
   describe the **pre-L1 single-stage** build L1 deleted, so both now over-reserve — conservative in the
   safe direction, and no longer silently so. The general mechanism is what remains.

### Routed forward (Fate 3 → M258), all recorded at the destination

`M258/overview.md` now carries an *"Inherited from the M257 close"* section naming every one:
`LEVER-M257-L5-setdress` (**the ranking moved: `set_dress` is now the largest phase at 28.6 %**) ·
`FIX-M257-dockerignore-env-pattern-unpaired` · `FIX-M257-anchor-guard-content-drift` ·
`FIX-M257-census-interpreter-namespace-import` · `RATCHET-M257-literal-ceilings-breached` ·
`FIX-M257-demopatch-sha-baselines-drifted` · `FIX-M257-campaign-kill-orphans-bringup` ·
`FIX-M257-sampler-disk-units-vm` · `MEASURE-M257-macmini-true-idle` · the four M255-inherited above · and
**two net-new at this close**: `FIX-M257-frontend-floor-is-billion-shaped` and
`FIX-M257-image-listing-conflates-empty-and-unreadable`.

### ⚠️ And a routing that had never reached its destination — fixed here

**M257x closed on 2026-08-11 routing 11 items / 5 clusters + 1 block fate to M258, and
`grep M257x m258…/overview.md` returned ZERO hits.** That is verbatim the `BIND_HOST` / `D-M255-7` failure
recorded *in that same file*: **a routing written in a closing milestone's decisions is not a routing until
the target's own doc says so.** The lesson was already written down there, one section above where the gap
was, and the very next close repeated it. M258's `overview.md` now carries an M257x section, and M257x's
`carry-forward.md` cluster 4 is annotated in place with the half M257 discharged (the measured `macmini`
profile; clause 1 now gradeable) so M258 does not inherit a claim that has stopped being true.
