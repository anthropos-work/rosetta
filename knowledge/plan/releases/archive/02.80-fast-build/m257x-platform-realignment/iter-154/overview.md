---
iter: 154
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-08
---

# iter-154 — the two bring-ups compute their verify scope by hand-appending literals

**Active strategy reference:** `TOK-08` — census the mechanical classes; stop sampling them.

**Step 0 — re-survey.** iter-153 closed the `/test-platform` half and routed
`FIX-M257x-iter153-bringup-scope-tuple-is-hand-written` forward on the ground that it *"does need a live
demo to grade."* The re-survey applies iter-153's own `D-M257x-153-3` to that sentence and it does not
survive either: the bring-up's verify tail is a bash block that can be extracted and executed against a
generated override exactly as iter-147 extracted `derive_profile` with `awk`. **And the target is bigger
than the route said** — the census below finds the same shape at **both** bring-ups, not one.

**Cluster / target identified.** Every site that computes a verify scope, and the artifact it reads:

| site | derives | then hand-appends | gated on |
|---|---|---|---|
| `up-injected.sh:2682-2690` | platform set | `next-web-app studio-desk` · `directus` | `NO_UI` · `NO_LOCAL_CONTENT` |
| `dev-stack:346-352` | platform set | `directus` | `local_content` |
| `generate.sh` | platform set ∪ **the stack's own override** | — | — *(iter-153)* |

Two of three still answer from the artifact that only *constrains*. `§5` rule 69 is explicit about what
to do here: *an observation about a twin is not a fence over it — when the sibling set is already
enumerated, write the fence over the set, in that commit.* iter-55 repaired the demo tuple and left the
dev twin for **30 iters**; harden pass 3 repaired the dev twin and left the demo side's hand-appends. The
class has now been half-repaired twice.

**Hypothesis.** Both bring-ups already write the stack's own override before they verify
(`up-injected.sh:1933`, `dev-stack:128`), so both can read it back through `scope-union.sh` instead of
naming services. Expected: the demo scope gains `hiring-app`'s row if one exists and, more importantly,
stops being a tuple; the dev scope stops naming `directus`.

**Expected lift.** Not an `N` reading. Two live-code sites re-pointed at the deciding artifact, in one
commit, with a twin-drift fence over the enumerated set.

**Phase plan.** A (census both tails + measure what each currently emits) → B (re-point both) → C (fence
both, extracted-and-executed, with mutation + anti-vacuity controls) → D (scoped suite runs + close).

**Escalation conditions.** If either tail cannot be exercised without docker, fence it against the
generator and route the re-point — but say which of the two, not "the bring-up".

**Acceptable close-no-lift outcomes.** If the union turns out to add nothing at either site beyond what
the hand-appends already name, the finding is that the tuples are *correct but unfenced*, and the iter
closes on the fence alone.
