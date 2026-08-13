---
milestone: M256
iter: 10
iteration_type: tik
status: closed
active_strategy: TOK-01
created: 2026-07-28
---

# M256 · iter-10 — D-v28-5, the cockpit logout defect (diagnostic)

**Type:** `tik` · **Active strategy:** `TOK-01` move 4. Handler: `D-v28-5-cockpit-logout`.

## Step 0 — re-survey

`ptvalidate` VALID: 10 products, 30 use cases, 23 live, 7 TODO. Clause 1 at 0.5652×; clause 2 mutating 6/5 MET,
negative controls 6 of 23, `blocked` 0; clause 3 verdict half COMPLETE, landed half short. **D-v28-5 unstarted
across 9 iters** — and it is a gate clause in its own right, not a nice-to-have, which is why it goes now.

## Cluster / target identified

**D-v28-5:** *"Logging out back to the cockpit requires two-or-more clicks."* It is the same seat-switch
machinery every Playthrough drives (`hero-login.ts` → the M37 cockpit handshake), and by the user's explicit
call it **gets no Playthrough** — so the deliverable is the fix, verified by hand-equivalent driving.

## Hypothesis

**H1.** The double-click comes from stale session state surviving a seat switch: `handleSelectIdentity` and
`handleHandshake` both drop the server-side session (`signedIn=false; sessID=""`), but the browser still holds
`__session` / `__client_uat` for the previous hero, so the first click renders the old identity and only a second
re-establishes coherently.

## Expected lift

D-v28-5 moves from unstarted to fixed, or to precisely diagnosed with a named handler.

## Phase plan

Reproduce the presenter's actual clicks against the real cockpit (`:27700`) and count what it takes to switch
heroes → diagnose → fix if the fix is contained → verify → close.

## Escalation conditions

- If the fix needs a platform edit → escalate. Never a platform edit.
- If the defect cannot be reproduced on this stack, say so and record why rather than fixing by inspection.

## Acceptable close-no-lift outcomes

- The defect proves unreproducible here for a stated, measured reason — that reason being worth more than a
  speculative fix.
