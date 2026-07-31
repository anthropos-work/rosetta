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

**F3: odysseus has ZERO swap**; billion has 15 GiB and used **2,452 MB** at peak. The headroom clause still
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
  **zero swap**; billion has 15 GiB and used **2,452 MB** of it at peak on the same measured lane. The
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
