# iter-96 decisions

## `D-M257x-96-1` — the repair is by PREDICATE, and the proof of that is a FENCE, not a claim

iter-95 routed `FIX-M257x-iter95-read-union` with one binding condition: *repair by predicate, never by
anchor*, because adjudicators had named **≥8** unbooked propagation sites. Measured, it was **38**.

An assertion that a sweep was complete is worth nothing — iter-40 ran a hand sweep with a mandatory
post-condition re-grep and still had a **27 % miss rate**. So the repair ledger is written in the shape
`claim_ledger.py` derives from, and completeness is asserted by `claim_twin_guard` (**101 → 114**
adjudicated claims, GREEN) rather than by me. A predicate that comes back is now a RED, tree-wide,
including in `corpus/ops/**`, `.claude/skills/**` and `CLAUDE.md` — the sites iter-40 measured a repair
leaking to the edge of and stopping.

## `D-M257x-96-2` — the standing bare-`grep` rule is AMENDED, because measured it was worse than the defect

iter-95's rule — *"an absence is established only by `git grep` at a named ref"* — is necessary and
**not sufficient**, and this is not a refinement. On the very predicate that produced the rule
(`mistralai`), the three instruments returned **1 / 0 / 22**, and the **0 was the ref-named `git grep`**.
`app/studio` is a nested untracked repo; `git -C app grep <anything> HEAD -- studio/` returns zero for
every predicate that has ever existed. The naive fix converts a noisy answer into a confident wrong one.

Amended as `platform-alignment.md` §5 **rule 44**: name the tree **and** its ref, **per tree**; enumerate
nested repos before any tree-wide zero; and note that NUL-bearing files are invisible to *both* tools, so
a zero over a tree containing one is a zero with a hole in it. Mechanized in
`anchor_construct_guard._clone_of`, which now descends to the innermost git checkout.

## `D-M257x-96-3` — the class was SIZED before it was acted on, and the sizing was itself wrong

The brief said to size the class before deciding what to do about it. Sized: **30** instrument-derived
absence-claims in the gate's scope, **13** exposed, **1 flipped**, **11 held**, **1 undecidable**. File
exposure: 12 gitignored-tracked text files, 2 NUL-bearing, 2 nested repos.

**My first census returned 4 and 1** — wrong by the mechanism under study (`git check-ignore` needs
`--no-index` to see tracked paths; the census did not descend into nested repos, so `cms/studio` was
invisible to the instrument measuring invisibility). Recorded in rule 44 rather than silently corrected,
because a census that cannot see its own blind spot is the finding.

**Disposition: fix the rule and the guard, do not sweep the claims.** Eleven of thirteen held; a blind
sweep would have rewritten eleven true sentences to no effect and manufactured its own induced defects —
which is the failure mode iter-41 measured at **9 of 18**.

## `D-M257x-96-4` — the storage hold is lifted for the present-tense falsehood ONLY, and the hazard grew

Per the run brief: `storage.md:58` repaired, the compose behaviour left escalated. Re-derivation found
the predicate reaches **10** sites, not 1 — and the sharpest is not in the gate's scope at all:
`safety.md:206` and `seeding-spec.md:101` classified **`S3-private` as `PerStackIsolated` / "seed
freely"**. That is the sentence a reader consults *before deciding a write is safe*, and it is false at
`0c91421`: `PreflightEnv` forces only `STORAGE_S3_PUBLIC_BUCKET` (`isolation/audit.go:146`); there is no
entry for the private bucket. `S3-private` is removed from that row and given its own, marked
**unclassified — the guard does not cover it**.

Every edit names the hazard and cites both sides. **No edit prescribes a fix, and none touches compose,
terraform or any platform file.** `DEF-M257x-iter80-storage-prod-bucket` remains the user's call.

## `D-M257x-96-5` — a prose repair is a line-number edit, and only half of that is fenced

Growing `external_services.md` by 24 lines silently moved **13 citations** in 6 files.
`anchor_construct_guard` caught **6** — the ones that happened to land on a blank line. The other 7
landed on non-blank wrong lines and were found only by diffing old→new line maps by hand.

This is a real gap in the fence family and it is stated rather than fixed here: the guard grades *what an
anchor points at*, and a wrong-but-plausible construct passes. Routed as
`CHECK-M257x-iter96-anchor-shift` — after any edit that changes a cited file's line count, re-derive
every inbound citation, do not trust the blank-line check.

## `D-M257x-96-6` — the archive-state class is repaired where evidence exists, ROUTED where it does not

The jobsimulation archive claim is **contradicted** (four commits dated four days after the claimed
archive, one a merge-button PR merge, on a repo an archive would make read-only) — repaired at all 6
sites to the map's own *report both, assert neither* hedge.

`skiller`, `skillpath`, `chronos`, `intelligence` and `graphql-wundergraph` carry the **identical
epistemic status** — an unmeasurable-from-a-clone GitHub state asserted as fact — but with **no clone to
measure and no contrary evidence**. Downgrading 15+ sites on symmetry alone would be the same unmeasured
assertion pointed the other way. Routed as `FIX-M257x-iter96-archive-class`.
