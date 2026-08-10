---
iter: 263
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10T15:28:00Z
---

# iter-263 — correcting iter-262's own `D-M257x-262-3`: the tooling already knew

**Type:** tik
**Active strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07),
under the user's binding `D-M257x-256-1`.

## Step 0 — Re-survey before targeting

iter-262 met the dev half and booked three defects. Checking whether the corpus needed updating for them
**immediately refuted part of one of my own**, before any corpus edit was written:

| what `D-M257x-262-3` implied | what the re-survey found |
|---|---|
| `INVITATION_HMAC_SECRET` is an undeclared boot requirement | `secretdna` declares **`platform/INVITATION_HMAC_SECRET`** as a gene that is **critical** + **required**, targeting `platform/.env`, scope shared, operators key-present + non-empty (`secret_dna_json_test.go:141-161`) |
| the `exit 0` shape is a new class | `secretdna/demo.go:45-47` names it **verbatim**: *"app/main.go's `invitations.NewTokenManager` ERRORS when it's empty and main returns (**the silent `app Exited (0)` class**)"* |
| the corpus should be told | `corpus/ops/secrets-spec.md` **already says it** at `:105` (the 32-gene platform list), `:118` (*"critical/required — the `app` exits early when it is unset"*) and `:273` (the demo-auto-generated family) |

**I reported as a discovery something the tooling had documented, in the same words, with a test pinning
it.** The milestone's currency is exactly this, so the correction is the iter rather than a footnote.

## Cluster / target identified

`D-M257x-262-3` itself. Establish precisely which part survives, and **why the gap still reached a dev
bring-up despite all of the above** — because something did go wrong, and mis-locating it is how the wrong
fix gets built.

## Hypothesis

The defect is **not** in the DNA, the docs, or the platform. It is that **the secret SOURCE is missing a
gene the DNA declares critical+required for a dev stack**, and iter-262 **never ran the instrument that
would have said so** — it hand-provisioned `.env` (`.env_example` + overlay) instead of driving
`/stack-secrets`. On a demo the same gene is satisfied *without* the source (`demoSatisfied` →
`IsDemoGenerated`, auto-generated at provision), which is the real reason the demo path never felt it.

## Expected lift

A corrected decision record, and the defect relocated from "undocumented platform/tooling gap" to "unrun
check + a source gap" — which changes what gets fixed.

## Pre-registrations — sealed in this iter's FIRST commit

| | claim | prediction |
|---|---|---|
| PR-1 | the DNA declares `platform/INVITATION_HMAC_SECRET` **critical + required** | **HOLDS** |
| PR-2 | `.agentspace/secrets/platform/.env` does **not** carry it, so a **dev** provision from that source is short a critical gene | **HOLDS** |
| PR-3 | `stacksecrets check` on a **dev** target **reports the gap** — the instrument exists and fires | **AT GENUINE RISK** — untested, and the whole correction turns on it |
| PR-4 | on a **demo** target the same gene reports satisfied without the source (`IsDemoGenerated`) | **HOLDS** — and this is why the demo never saw it |
| PR-5 | `D-M257x-262-3`'s "undeclared / new class" framing is **REFUTED by the tooling's own comment** | **HOLDS** — i.e. my own published claim was wrong |

## Phase plan

- **Phase A** — seal. **Phase B** — build + run `stacksecrets check` against dev and demo targets.
- **Phase C** — correct `D-M257x-262-3` in place (retract, do not silently rewrite). **Phase D** — close.

## Escalation conditions

If `stacksecrets check` does **not** report the gap (PR-3 refuted), the defect is larger than a source gap
and the correction must say so rather than land the tidier story.

## Acceptable close-no-lift outcomes

A precise retraction is the deliverable. No metric moves; the corpus gets *more* correct, which under
`TOK-08` is the point.
