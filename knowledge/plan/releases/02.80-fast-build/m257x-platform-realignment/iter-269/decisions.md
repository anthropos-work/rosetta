# iter-269 — decisions

## Pre-registrations — SEALED BEFORE THE MEASUREMENT

Sealed in this iter's first commit, corpus at `5c2de87`. Known at seal time: 470 lines, six keys at
exactly 31 occurrences. Nothing about block-identity, the writer, or value variation has been read.

**PR-1 — all 31 blocks are byte-identical.**
No key takes two different values across the file. *Risk:* if any key varies, the file is not duplication
but a **silent value history**, and the last block is deciding production-shaped behaviour by accident.

**PR-2 — `DIRECTUS_TOKEN` is blank in all 31.** Re-confirming iter-262 on today's file. *Risk:* a later
`/stack-secrets --provision` may have appended a real one, which would itself demonstrate PR-4.

**PR-3 — the writer is the demo bring-up path, not `stacksecrets provision`.**
`stacksecrets` is documented append-only *by design* and writes per-repo targets values-blind; the 31
identical blocks look like a **seed/copy** step re-running, not a provisioner. *Risk:* real — if the
provisioner is the appender, the fix touches the values-blind contract and is a different, more delicate
repair.

**PR-4 — last-wins makes this a correctness hazard, not a cosmetic one.**
Compose resolves a repeated key to its **last** occurrence, so append-order decides the value. *Risk:*
if compose errored on duplicates, or the file is never read directly by compose, the hazard evaporates.

**PR-5 — the fix needs a tag + pin bump and is therefore routed, not landed.**
*Risk:* the writer may be a shell step whose correction is corpus-side documentation only.

## Escalation clause (pre-registered)

**`stack-demo/platform/.env` is NOT rewritten, deduped, or truncated in this iter**, whatever is found. It
is a live stack's environment and the evidence for every claim above. If the repair requires editing it,
that is described and routed, never performed. (It is also a `.env` — never committed, never echoed.)
