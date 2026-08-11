**Type:** tik, under `TOK-05`.

## The measurement

iter-88 routed four failures as up to three classes. **They are one**, and the first probe found it —
`git status` on the demo clones:

```
stack-demo/next-web-app   M packages/core-js/src/constants/urls.ts
                          M packages/ui/src/NavBar/NavbarTop.tsx
```

`NavbarTop.tsx` carries the `next-web-back-to-cockpit` demo-patch, **left applied**. Everything else
follows: a patched file's whole-file sha does not match a pristine baseline, so the `pre_sha256`
assertions fail — they were **downstream of the un-reverted patch**, not a separate class.

## Why the revert refused — and it refused *correctly*

```
demopatch: revert REFUSE: target sha256=a28808e0… is neither pre nor post — manual drift;
refusing to guess. Use --force-pristine to `git checkout -- <path>` the demo clone.
```

That is G2 (drift-refuse) doing exactly its job. The interesting question is why the file was neither.
Four shas settle it:

| | sha256 |
|---|---|
| manifest `pre_sha256` | `0c2c2ed2…` |
| manifest `post_sha256` | `5ae9d1db…` |
| **pristine file at the clone's `HEAD`** | **`48b6dd07…`** |
| current (patched) file | `a28808e0…` |

**The pristine file is not the manifest's `pre`.** The clone's base has drifted from when the baseline
was recorded — `next-web-app` sits **41 commits behind** origin and the baseline predates that. So:

1. **Apply** resolved its **anchor**, found it, and self-healed onto the drifted base — by design.
   `demopatch-spec.md`: *"the anchor is the contract; the whole-file sha is only a baseline."*
2. The result is `base' + patch`, whose whole-file sha is neither the recorded `pre` nor the recorded
   `post` — **necessarily**, because both were computed over the old base.
3. **Revert** compares whole-file shas, matches neither, and refuses.

> **The asymmetry is the defect: APPLY is anchor-based and self-heals across base drift; REVERT is
> whole-file-sha-based and cannot. On any clone whose base has drifted from the recorded baseline —
> which is the NORMAL state, since these clones sit tens of commits behind — apply succeeds and revert
> refuses, and the clone is left dirty.**

That contradicts the mechanism's headline promise, in the spec's own words: *"the clone is left
git-clean, and the canonical `anthropos-work` repos are never touched."* G5 (self-revert) does not hold
whenever G2's freshness gate has done its job — **the two guards are in tension, and the tension is
structural rather than a bug in either.**

It is also the same shape as the defect `demopatch-spec.md` already warns about (*"a silently-refused
perf patch shipped a 76 s members grid for four releases"*) with the sign flipped: there a **refused
apply** was silent; here a **refused revert** is loud but leaves state behind.

## Why nothing was repaired

Two escalation conditions fired, and both are Phase 5 §4 user-blockers rather than route-forwards:

1. **The fix is an architectural choice**, and it changes what code lands:
   - **(a) make revert symmetric** — reverse the anchor transformation rather than compare whole files;
   - **(b) journal the observed pre-state at apply time** and have revert restore exactly that (the
     "record what you did" option — strongest, and it makes revert independent of baselines);
   - **(c) make apply strict** — refuse on base drift, losing the self-healing that exists because base
     drift is normal;
   - **(d) accept it and have the runner use the existing `--force-pristine`** escape.
   These have different blast radii on a mechanism that **rewrites platform source inside a build**.
   Choosing one on my own authority is the single-seat adjudication this milestone has been burned by.
2. **The demo clones are dirty right now**, and cleaning them is a forbidden op for me. `--force-pristine`
   is the tool's own sanctioned path and it runs `git checkout -- <path>`; that is a decision about
   uncommitted state, which only the user and the orchestrator may take.

**The clones were left exactly as found.** Nothing was cleaned, no baseline was re-pinned, no manifest
was touched — re-pinning a `pre_sha256` here would have *hidden* the asymmetry by making one file match
again, which is precisely the wrong repair and exactly what iter-88's routing instruction forbade.

## Close — 2026-08-05

**Outcome:** the four routed demo-stack failures are **one structural defect**, root-caused and measured:
demopatch's apply self-heals across base drift while its revert cannot, so on a drifted clone — the
normal case — the patch applies and will not come off, leaving the clone dirty and contradicting G5.
No fix landed; the repair is a design choice with real blast radius, and the dirty clones are
uncommitted state only the user may decide on.
**Type:** tik
**Status:** closed-no-lift — the planned investigation completed with a documented root cause. A complete
cycle ending in characterization, which the protocol treats as first-class; no fix was attempted, so
nothing was reverted.
**Gate:** NOT MET — **4 of 5, unchanged.** No reading taken.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this was a tik; the streak is not 3 — iters
87 and 88 both closed `closed-fixed` with landed deliverables) — (3) re-scope: n (platform re-fetched at
open and close: `0c91421`, unchanged) — (4) **user-blocker: y** — an architectural question whose answer
changes what code lands, plus uncommitted state in the demo clones that only the user may decide on —
(5) cap-reached: n (3 tiks this session) — (6) protocol-stop: n — **Outcome: exit-4**
**Decisions:** `D-M257x-89-1` (one cause, not three — recorded below) · `D-M257x-89-2` (do not re-pin, do
not clean; escalate)
**Side-deliverables:** none.
**Routes carried forward:** the four handlers from iter-88 are now **one** —
`FIX-M257x-iter89-demopatch-revert-asymmetry`, blocked on the user's choice of (a)–(d) above. The two
remaining iter-88 handlers (`CHECK-M257x-iter88-live-stack-tests`,
`CHECK-M257x-iter88-unnamed-skips`) are untouched and still open.
**Lessons:** **§5 rule 28 earned again — three true facts do not make a cause, and the joining experiment
was `git status`.** iter-88 booked these as up to three classes across two handlers, each description
accurate. One probe against the clone collapsed them to one root cause in under a minute. When several
checks fail against the same external artifact, **look at the artifact's state before classifying the
failures** — the cheapest experiment is usually the one that reads the thing everyone is asserting about.
And the finding underneath: **two guards can each be individually correct and jointly inconsistent.** G2
(refuse on drift) and G5 (always self-revert) cannot both hold once the base is allowed to move, and no
test asserted their conjunction — each was verified alone.
