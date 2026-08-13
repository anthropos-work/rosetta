---
iter: 04
milestone: M256
iteration_type: tik
status: closed-fixed-partial
opened: 2026-07-28
---

# M256 · iter-04 — org-admin: the cluster that discharges clause 2 and clause 3 together

**Type:** tik · **Active strategy:** `TOK-01` move 3.

## Step 0 — re-survey

TOK-01 move 3 names org-admin as the next target, ahead of onboarding, because all four curated UCs declare
a **persist-then-observe** final (iter-01 D2) and onboarding is seed-blocked (audit F5). Re-surveyed: nothing
since has touched `/enterprise/*`; org-admin coverage is still **0 of 4**; clause 2's mutating count is still
**1** (`pt-assignment-assign`). Target stands.

## Phase 0d — pre-flight tooling check: **PASS**

`ptvalidate --manifest-dir ./manifest --e2e-dir ./e2e/tests --seed-worlds ./seed/seed-worlds.yaml` on the
existing tree: `manifest VALID: 8 product(s), 18 use case(s), 18 live Playthrough(s), 0 TODO`, exit 0. The
pipeline accepts the current file, so new entries will be gated on their own merits rather than on
pre-existing debt.

## Cluster / target identified

The **four curated org-admin use cases** (M201 corpus, `org-admin-settings`), all un-homed for 5 releases:

| UC | Route | Curated final |
|---|---|---|
| `org-admin-settings.roles.UC1` | `/enterprise/roles` | the role **persists** with its configured skills |
| `org-admin-settings.members.UC1` | `/enterprise/members` | the role/team assignment **persists** on the member |
| `org-admin-settings.tags.UC1` | `/enterprise/tags` | the tag/team **persists** with its members |
| `org-admin-settings.feature-config.UC1` | `/enterprise/settings` | the setting **persists** |

Landing them moves clause 3 from **0/4 → 4/4** on this half *and* clause 2's mutating count from **1 → 5**,
which is exactly its `≥ 5` floor. No other cluster in the milestone serves two clauses at once.

## Hypothesis

Each of the four is a WRITE with a read-back surface, so each can be proven in the `pt-assignment-assign`
shape: drive the real UI (P1), then **re-read through the app** and assert the changed state — never a
closed modal, never a toast.

**The open question this must answer** (`../overview.md`): *do the org-admin writes have a read-back surface,
or only a toast?* A write whose effect is invisible in the UI cannot be proven under P1/P2 without a DB
assert, which is a weaker proof shape. So the iter **probes before it commits**.

## Expected lift

Clause 3: +4 covered UCs. Clause 2: mutating 1 → up to 5. Median: the new tests must not raise it — they are
authored with `waitUntil: 'domcontentloaded'` from the first line (now fence-enforced), so each should land
near the ~2 s post-iter-03 cost.

## Phase plan (protocol steps 1–6)

- **Phase A — probe the four surfaces live** as `pt-manager` on `demo-2`: does each route exist, is the org
  admin admitted, what are the write affordances, and **what reads the write back**? Record a per-surface
  verdict before writing a line of manifest.
- **Phase B — declare** the drivable UCs in a new `manifest/org-admin.yaml` (P5, `playthrough: TODO` first),
  plus any `seed-worlds.yaml` capability **in lockstep** (a partial landing is a broken validator, not a head
  start — the M219 rule). `ptvalidate` green at every step.
- **Phase C — page objects + specs**, one per drivable UC, each asserting a **read-back**.
- **Phase D — run + reconcile + re-measure** under D7's protocol; report the median on the grown denominator.

## Escalation conditions

- A surface that cannot be driven without a platform edit → **`unimplementable-without-platform-edit`** in
  `report/unimplementable.yaml` with a rationale (the P3 escape valve). It escalates; it never edits the
  platform.
- A write with **no** read-back surface → declare the UC, record the honest verdict, and do **not** ship a
  toast-only assertion dressed as a proof. That is the failure mode this milestone exists to end.
- `> 3` of the four proving unimplementable → the milestone's own **re-scope trigger** fires (a platform
  conversation, not a test one).

## Acceptable close-no-lift outcomes

A probed, written verdict for each of the four — even if fewer than four land — is a complete iter, provided
each non-landing is *characterised* rather than skipped. `closed-fixed-partial` is the honest status if some
land and others carry a verdict.
