# M270 — Decisions

_Implementation + design choices with rationale, recorded as they are made._

| # | Decision | Rationale | Date |
|---|----------|-----------|------|
| D-1 | _(open — **the vehicle**: sha-pinned demopatch vs escalate to the platform team, and whether the two halves (a) and (b) take the same verdict)_ | | |
| — | _(none others yet)_ | | |

---

## D-1 — the vehicle (OPEN, must be decided explicitly)

**Why this is a decision and not an implementation detail.** Every line of both defects is in
`next-web-app`, a platform repo, and this corpus takes **zero platform-repo edits**. So the fix has exactly
two legitimate routes:

1. **A sha-pinned demopatch** ([`demopatch-spec.md`](../../../../../corpus/ops/demo/demopatch-spec.md)) —
   the sanctioned zero-platform-edit escape hatch. It patches the demo's own ephemeral clone just before
   the image build and reverts after; the image carries the fix, the clone is left git-clean, and the
   canonical `anthropos-work` repos are never touched.
2. **Escalation to the platform team** — if the change is big enough that a patch is the wrong vehicle. In
   that case this milestone **re-scopes to the diagnosis + the patch for the loading affordance only**.

**Do not let this be chosen by accident.** The failure mode is drifting into a patch because a patch is what
the tooling makes easy, and then carrying a page-shape change on a sha-pinned anchor across every platform
move.

**Inputs the decision needs (none of them collected yet):**

- Anchorability of the (a) fix — expected small: destructure `isLoadingOrg`, thread it, render an
  affordance.
- Anchorability of the (b) fix — expected large: `apps/web`'s page goes client component → server
  component to take the SSR-prefetch path `apps/integration` already uses.
- The maintenance cost either way: a manifest is **sha-pinned and WILL drift**. G2's model is *the anchor is
  the contract; the whole-file sha is only a baseline* — a moved file self-heals, but an anchor that occurs
  **zero** times or **two or more** times **REFUSES**. A silently-refused patch has shipped a defect for
  four releases before.
- The inventory cost: §5 of the spec plus the `TestPatchInventory` fence, which pins the exact total and the
  per-repo breakdown.

**The two halves may legitimately take different verdicts.** Record them separately if so.
