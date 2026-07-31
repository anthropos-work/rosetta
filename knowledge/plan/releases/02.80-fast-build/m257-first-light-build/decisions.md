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
