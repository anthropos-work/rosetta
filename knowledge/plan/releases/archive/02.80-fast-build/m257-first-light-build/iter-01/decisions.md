# iter-01 — intra-iter decisions

Strategy lives in the **milestone-root** [`decisions.md`](../decisions.md) as `TOK-01`; the three
contract calls answered mid-iter are there as `D120` / `D121` / `D122`. Only iter-local decisions here.

## D1 — the audit's RED was escalated, not remediated-then-continued

Phase 0b offers *"fill blind areas before proceeding"*; the invoking contract said a RED is a **hard
stop, surface it**. Escalated, because three of the required fixes edit contracts M255 shipped nine days
earlier (a test that hard-pins the single baseline source, the retraction scope of a doc M255 authored,
and an orphaned re-confirmation instruction). Those are **user-facing contract calls, not blind-area
backfills an agent should pick a side on**. Resolved as `D120`/`D121`/`D122`, then remediated.

## D2 — no `iter-01/` dir existed while the gate was RED

Deliberate. The gate fired **upstream of Phase 1**, so the bootstrap tok had not begun. Leaving the
milestone at zero iter dirs kept the next invocation's Phase 0 correctly routing iter-01 to a bootstrap
tok, rather than seeing a stub dir and skipping to a tik with no strategy. The dir was created only once
the audit cleared.

## D3 — the completed audit report WAS committed, though the iter had not closed

Phase 4's "do not commit partial work" governs a **budget-exhausted mid-Phase iter**. The audit report is
a finished, standalone artifact, and Phase 0b independently requires its verdict be recorded in
`spec-notes.md`. Committing it was that recording obligation, not a partial-iter landing.

## D4 — the §8.5 residuals in `roadmap.md:657` were fixed, though only the overview was in the brief

The re-anchoring brief named the M257 `overview.md`; the roadmap carried the **same** stale
`:231/:249/:262/:271` enumeration. Fixing one and not the other would leave the two disagreeing about
what the work list *is* — the drift class this release hunts. Same reasoning for the two index mirrors
still reading *"three clauses"* after the body said four: `D121`'s own rationale is that no intermediate
state should be wrong in a **new** way.

## D5 — one YELLOW residual was fixed rather than carried, on hazard grounds

`stack-core/README.md` handed an operator a copy-pasteable `buildbench` campaign pointed at
`--profile billion --public-host billion.taildc510.ts.net` — a machine that is **off-limits** under
`D-v28-14`. Carrying a documented instruction to touch a forbidden host as "known-context" is not a
tracked gap; it is a live trap. Fixed. The remaining YELLOW residuals were genuinely carried, into
`TOK-01` § Known-context.
