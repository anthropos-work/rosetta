---
title: "Deferral Audit — M236 close continuation (post-user-decision re-run)"
date: 2026-07-20
scope: milestone
invoked-by: close-milestone (continuation)
supersedes: the RED audit that held the M236 merge (recorded in decisions.md CLOSE-D1/CLOSE-D2)
---

## Verdict

**YELLOW** — `SEVERITY=warning`. The merge may proceed.

Not RED: the one blocking item now carries a **fresh, dated, explicit user decision** (2026-07-20), which is
precisely what the prior audit demanded and could not manufacture.

Not GREEN: the item is a genuine repeat-deferral whose disposition is *accepted destination with sign-off*
(the escape hatch), not a landing. Calling that GREEN would repeat the exact error this audit exists to
catch — recording an unlanded item as resolved.

## Summary

- Total deferrals in scope: **9** (unchanged from the prior audit)
- Landed at the close (Fate 1): **3** (items 3, 5, 6a)
- Discharged / falsified: **1** (item 8)
- Routed Fate-3 to the release close: **3** (items 1, 2, 7)
- Repeat deferrals: **1** (item 4, incl. the folded 6b)
- Repeat deferrals **without** a resolution decision: **0** ← the RED condition, now cleared
- Chronic patterns still flagged: **0**

## What changed since the RED audit

The prior audit returned RED on a single item and named exactly what would clear it: *"re-baseline the set
against HEAD, then an explicit user fate."* Both have now happened.

### The user decision (2026-07-20)

**RE-BASELINE now; decide the fate at release close.** Explicitly not fixed now, not dropped, and not
silently rolled forward again. Recorded verbatim in `decisions.md` → CLOSE-D2 → **RESOLUTION**, with the
user's measurement condition (a fresh demo whose platform repos are at stable `main`) quoted in their own
words.

### The re-baseline

**→ `../rebaseline-standing-failures.md`.** Measured against a clean, stable-`main` clone set on 2026-07-20,
with the resolved repo → ref → sha table recorded so the reading is reproducible.

| | |
|---|---|
| Carried label | 14, "6 of them `pre_sha256` pin drift" |
| Reproduced before any change | **14** — the count was accurate |
| On a clean stable-`main` clone set | **8** |
| Real defects | **0** |
| Pin drift | **0 — the diagnosis is refuted** |

## Aged-out rule — satisfied

The prior audit correctly fired **AGED_OUT** on item 4 (destination milestone closed without landing it;
deferred across ≥2 milestones; ≥2 releases). That rule requires a **fresh** decision that answers *"given
everything the project has learned since, do we still believe this? What new context changed?"* — not a
rubber-stamp of the old reason.

**New context that the fresh decision now rests on, none of which existed at the prior deferrals:**

1. **The count is right but the diagnosis was wrong.** Zero of the 14 are pin drift. Six were a **dirty
   clone** — leftover applied demo patches — reporting itself as a test failure.
2. **The remedy the old label implied was the dangerous action.** Re-anchoring the "drifted" pins would have
   re-pinned five manifests to *patched* content, permanently disarming the drift detector for those files.
   Every pin is correct at both the stale ref and `main`. **This is the strongest argument that the item was
   right not to be landed blind across those ten milestones** — the obvious fix was actively harmful.
3. **The count is host-dependent** (8 macOS / 7 Linux expected). *"N failures"* was never a well-defined
   number, which is a sufficient explanation for the 8 → 14 drift under a fixed label.
4. **Two upstream findings were discovered while re-baselining** (`F-M236-CLOSE-1`, `F-M236-CLOSE-2`), one of
   which is the generator of the whole class.

The fresh decision is therefore an informed re-decision, not a renewal.

## Item-by-item

| # | Item | Fate | Status |
|---|---|---|---|
| 1 | anon `/library` + `/free` render 0 cards | Fate 3 → release close | routed; in `carry-forward.md` Route 4 |
| 2 | `apps/web` client GraphQL on non-offset `:5050` | Fate 3 → release close | routed, batched with #1 |
| 3 | v2.4-era method docs (F3/F4/F6/F7/F10) | **Fate 1 — LANDED** | closed |
| 4 | **the standing demo-stack test failures** | **KEEP-DEFERRED-WITH-SIGNOFF → release close** | **user-decided 2026-07-20 + re-baselined** |
| 5 | `run-coverage.sh` / `run-hiring-render.sh` non-integer-`N` guard | **Fate 1 — LANDED** | committed + published |
| 6a | seed ships 4 orgs, docs say 3 | **Fate 1 — LANDED** | 20 places, 8 docs, 2 skills, `CLAUDE.md` |
| 6b | `test_m220_mutation_battery.py` unmutated subject | folded into #4 | covered by the re-baseline |
| 7 | `DEF-M235-03` M204 assign-WRITE TODO | Fate 3 → release close | inherited, correctly routed |
| 8 | M230 carry-forward cluster 2 | **DISCHARGED** | diagnosis falsified; cold bring-up passed |

**Item 4's destination is legitimate, and this is the one place it could be.** M236 is the **final** v2.5
milestone — the next event *is* the release close. This is not a roll-forward into another milestone that
may also not happen; it is a hand-off to the immediately-next event, with a real characterisation and a
recommended fate attached instead of a stale label.

## Blocking items

**None.** The single blocking item from the prior audit is dispositioned with a recorded, dated user
decision and a durable re-baseline artifact.

## Applied changes

- `decisions.md` — CLOSE-D2 → **RESOLUTION** block; `F-M236-CLOSE-1`; `F-M236-CLOSE-2`
- `rebaseline-standing-failures.md` — new; the re-baseline, with reproducible refs
- `carry-forward.md` — new; routes item 4 + both findings + the Fate-3 items to `/developer-kit:close-release`
- Live cold bring-up on `billion` at stable `main` — autoverify GREEN, 0 patches refused

## Note to `/developer-kit:close-release`

Item 4 arrives with a **recommended fate of Fate 1 (LAND)** and an argument for why it is now cheap: all 8
remaining failures are test-side edits with no product risk and no live stack required. Treat
`F-M236-CLOSE-1` **separately and at higher priority** — it is not part of the 14 and should not inherit
their fate.

---

`DEFERRALS: YELLOW` · total 9 · repeat 1 (dispositioned) · blocking 0 · `SEVERITY=warning`
