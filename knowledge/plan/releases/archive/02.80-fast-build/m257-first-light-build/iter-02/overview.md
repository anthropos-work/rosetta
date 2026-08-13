---
iter: 2
milestone: M257
iteration_type: tik
status: in-progress
created: 2026-07-31
---

# iter-02 — make odysseus a bench, and the instrument falsifiable

**Type:** tik · **Active strategy:** [`TOK-01`](../decisions.md) — *instrument before baseline,
baseline before levers*

## Phase 1 Step 0 — re-survey

`TOK-01` was authored in the immediately-preceding iter, so its next-tik direction cannot be stale.
Confirmed cheaply anyway, per the mandatory re-survey: odysseus still reachable
(`devops@100.110.67.14`), Docker still **completely empty** (0 images / 0 containers / 0 build cache),
`stack-core/hostprofiles/` still holds **2** profiles (no `odysseus.json` yet). Target unchanged.

## Active strategy reference

`TOK-01`, step (1) *"the host can run a cycle at all"* and step (2) *"the gate's own instrument is
proven able to fail"*. This iter is **both** — deliberately, because step (2)'s negative control needs
a live stack and step (1)'s validation *is* a live stack.

## Cluster / target identified

The gate's clause 2 is literally `autoverify green:true / 0 warnings`, and **three separate pieces of
evidence say that verdict cannot currently be trusted on this host**:

1. **M256 found 43 checks that reported success without checking** — including a probe runner
   certifying *"all live probes passed"* over **zero probes**, whose exit code four other gates read as
   health. This gate consumes that verdict directly.
2. **`FIX-M256-demo2-service-self-termination`** (inherited): `green / 0 warnings` can **PASS** while
   two services sit at `Exited 0` — surfaces render empty, no error anywhere, an hour of misdiagnosis.
3. **`FIX-M256-autoverify-fapi-libressl`** (inherited): a **working** stack can emit the warning this
   gate counts.

And a fourth, specific to this host: **odysseus has zero swap**, so a headroom overshoot meets the
**OOM killer**, whose symptom is *indistinguishable* from (2). Every lever in this milestone pushes on
memory. The liveness check is the instrument that tells those apart, which is why `TOK-01` puts it
before any lever.

## Hypothesis

If the host is provisioned and both inherited `autoverify` defects are fixed **and** the verdict is
demonstrated to flip RED under a deliberately-broken stack, then every subsequent iter's measurement
rests on an instrument that has been *watched failing* — and a headroom breach on this swapless host
will name itself instead of masquerading as the silent-empty symptom.

## Expected lift

**Zero on the primary metric, by design.** No lever is touched and no gated number is produced. This
iter grades on its **planned deliverables** per Phase 4 Step 0 (*"planned scope = what the
`overview.md` committed to"*; a probe/instrument iter has no production-code lift to claim). Claiming
a p50 movement here would be exactly the un-probed-lift dishonesty Phase 3's self-check forbids.

## Phase plan

- **(a) Provision odysseus** per `corpus/ops/demo/tailscale-serve.md:119-131`. **Go 1.26.5 is already
  installed** at `/usr/local/go/bin/go` and merely off PATH → **fix PATH, do not install over it**
  (`TOK-01` known-context #6). **atlas is genuinely absent** → install. Plus ssh-agent, the snapshot
  cache, and a confirmation that the **six** rext Go modules build against 1.26.5 rather than the
  1.25.x the prereq list names.
- **(b) Land both inherited `autoverify` fixes.**
  - *Container-liveness cheap-win*, scoped to the **real** hole per `TOK-01` known-context #1:
    `verification.md:623-626` misattributes the cause (it reads *"autoverify cannot see this"*, but
    `stack-verify/lib/services.sh:43-44` **does** carry `jobsimulation` + `cms` rows, so a re-run would
    have gone red — the M256 stack stayed green via the **stale-verdict** class). The genuine gap is
    that **`fake-fapi`/`fake-bapi` have no `services.sh` row** (the 16-vs-13 count gap). Fix the check
    *and* that paragraph.
  - *A fapi probe independent of the host TLS stack* — **but first establish empirically whether check
    (d) even warns on Linux/OpenSSL.** The defect is a *macOS host `curl`* LibreSSL/mkcert handshake
    failure; asserting it fires on odysseus would be assuming, and `TOK-01` known-context #2 forbids
    that in either direction.
- **(c) Prove `autoverify` can go RED** — a deliberate negative control: stop a service on a healthy
  stack, confirm the verdict flips to `green:false` **and names the dead service**, then restore
  (`docker start` is non-destructive; the M256 recovery path). **A gate nobody has watched fail is not
  a gate.**

The bring-up that (c) needs **doubles as the mandatory discarded warm-up cycle** that `TOK-01`
known-context #7 requires before iter-03's `n ≥ 3` campaign — rep 1 on an empty-cache host measures the
truly-cold variant `D-v28-8` cut from the gate. It is discarded either way, so it is free here.

## Escalation conditions

- A provisioning wall that needs a credential or a host change only the user can authorise → **exit
  `user-blocker`**.
- The negative control shows autoverify **cannot** be made to go red → that is a *finding*, not a
  blocker: land it, record it, and escalate the gate-trust question in the close.
- Bring-up failure on this host: **run `df -h /` and `docker system df` FIRST.** A mid-campaign ENOSPC
  presents as the cryptic `redis exited (1)` (M239-F1), not as a disk error — and under speed work it
  reads as *"my change broke the stack."*

## Acceptable close-no-lift outcomes

- Provisioning lands and the fixes land, but the host cannot complete a full bring-up for a reason
  outside this iter's scope → `closed-fixed-partial`, with the bring-up routed to iter-03.
- Check (d) turns out **not** to fire on Linux → that half of (b) becomes a laptop-path fix with a
  one-sentence doc correction, and the iter still closes on (a)+(c)+the liveness half.
