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

## D-M257x-269-1 — last-wins is LOAD-BEARING, so "replace-or-skip" is the wrong repair

Measured values-blind on `stack-demo/platform/.env` after 31 bring-ups: **471 lines · 18 distinct keys ·
13 present 31 times · 0 keys whose value varies · `DIRECTUS_TOKEN` blank in all 31.** (iter-262 said
"18-key block"; the block is **13** keys, with 5 singletons alongside.)

The writer is **`stacksecrets provision`**, and it appends **by design**: `provision/io.go:173-175` —
*"an existing line is never re-read for its value or rewritten, so provision can never corrupt or echo a
value already in the target."* That is what makes the tool values-blind. The bring-up's contribution is
`up-injected.sh:1538`, which passes **`--force` unconditionally**, skipping copy-if-absent — and `:1522`
states the second purpose: `--force` *"overwrites stale keys AND blanks the `DIRECTUS_TOKEN` family via
last-wins (the strip-on-non-prod class)."*

**So compose's last-wins resolution is not incidental — the demo's `DIRECTUS_TOKEN` blank is delivered by
being appended LAST.** iter-262 routed *"find the writer and make it replace-or-skip"*, and that
instruction is refuted on its own terms: replace-in-place would either **re-read an existing value**
(breaking values-blindness) or **drop the trailing blank** (re-arming `DIRECTUS_TOKEN` on a demo — the
fix16/17 class `secrets-spec.md` exists to prevent). A real repair keeps **both** properties and prunes
**older** duplicates rather than ceasing to append. Re-routed as
`FIX-M257x-269-force-append-grows-the-demo-env-without-bound`.

**The hazard, stated conditionally because that is what the evidence supports.** All 31 copies agree
today, so nothing is broken. But the file is *designed* to accumulate and there is no reaper: with N
copies, the **last** wins, so any writer appending a differing value silently decides it. Diagnostic rule
now in the spec: **read the LAST occurrence of a key, never the first.**

**The generalisable half: two correct decisions composed into a defect that is neither's bug.**
Append-only is right on its own; `--force` is right on its own; unbounded growth belongs to the
composition, which has no owner and therefore no failing test. This is why the condition ran 31 times
without anything going red — and why the fix has to be specified against *both* invariants at once rather
than by "fixing the writer".
