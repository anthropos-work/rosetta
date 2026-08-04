# iter-76 adjudication scratch — running

## r13-F B13 / B14 — REJECTED (false positive), and the same mechanism in both

Claim: `frontend_architecture.md:11` cites `docker-compose.yml:311` and `:39` cites `:362`, both
"past end of 271-line file".

Traced through the guard: both citations sit in blocks that **pin `2adcf71`**, where
`docker-compose.yml` is **387 lines**. `:311` and `:362` both resolve to constructs there.
`classify -> None` at the pinned ref for both.

Line 39 names its own pin in words — *"since platform `2adcf71`"*. Line 11's pin governs its block.

**§5 rule 33: a claim is settled at the ref the claim itself names.** The briefing says this
explicitly and in its own section. The seat graded a dated claim against the checkout.

Both booked `medium` confidence — the seat hesitated, correctly.

**Residual question, NOT closed by the above:** line 11's claim is about whether studio-desk *has* a
compose profile. At `0dab54d` studio-desk still has one, so the claim is true today as well as at
the pin — only its ANCHOR is dated. Worth a scoping note, not a repair.
