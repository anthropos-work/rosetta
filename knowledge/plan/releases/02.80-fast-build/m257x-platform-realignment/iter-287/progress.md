# iter-287 — the two repairs iter-262 had to invent, written into the guide that omitted them

**Type:** tik — under `TOK-09`, item (c).

## Re-survey first: one of the three was already done

iter-262 recorded three out-of-band repairs. Before writing anything, each was checked against the guide
as it stands today:

| repair | state |
|---|---|
| acquire `app/studio` before `make up` | **ALREADY DOCUMENTED** at iter-264 — §4 carries a *"REQUIRED before `make up`"* section and §6 step 2 points at it |
| `INVITATION_HMAC_SECRET` undeclared, `backend` exits **0** | mentioned twice — in §"Automated Setup" and §"Verify secret coverage" — **but nowhere at the step where the symptom appears**, and with no remedy |
| `make init` is skip-if-present and adopts a stale tree | **not documented anywhere** |

Two of three routes were live; the third was closed and re-doing it would have been the stale-route waste
this milestone has paid for three times.

## What landed

**1. `make init` is SKIP-IF-PRESENT, and a skip is not a check** (§4). It clones only when the sibling
directory is absent; when one exists it prints `<repo> already exists, skipping` and **adopts that tree at
whatever ref it holds**. The line reads like progress and the step compares no refs. iter-262 adopted a
`studio-desk` that was **13 files / +97 / −46** behind its own `main` on an otherwise clean bring-up. The
new note names the output verbatim, says `make status` is the thing that tells you, and gives the
`git fetch && git status -sb` check to run on anything that says "skipping".

**2. `backend` exits 0 on a missing `INVITATION_HMAC_SECRET`** (§6, at `make up` — the step where you
meet it). The two existing mentions are both upstream of the failure: one in the *"use `/dev-up`"* pitch,
one in the secret-coverage section. Neither is where a reader is standing when four containers come up
instead of five. The new note gives the symptom (`Exited (0)`, no crash, no restart loop, nothing for a
health check), the contrast that makes it diagnosable (`sentinel` fails honestly with `Restarting (2)`),
the fact that **the variable is not declared in `.env_example`** — re-verified this iter, `grep -c` →
**0** at `platform 0c91421` — and a one-line remedy.

> **Why the demo path never hits it, which is the reason it stayed invisible:** the key is in
> `secretdna.DemoGeneratedKeys`, so `/stack-secrets` mints it. Only the hand-built dev path is exposed,
> and that is the path this guide *is*.

## Verification

| scope | result |
|---|---|
| `INVITATION_HMAC_SECRET` in `platform/.env_example` @ `0c91421` | **0 occurrences** — the claim re-derived, not inherited |
| `prose_twin_guard` | **OK — 0 RED** (a third copy of a claim is exactly what this fence watches) |
| `markdown_structure_guard` | **OK** — 114 published files, no structural damage |
| `corpus_index_guard` | **OK** — 86 docs across 6 index-bearing directories |
| `corpus_citation_guard` | **OK** — every enumerated citation resolves |

**NOT COVERED:** the guide was not re-executed. These are documentation repairs derived from iter-262's
recorded measurements, not a fresh bring-up — and a guide is only really tested by running it on a clean
box, which is how these defects were found in the first place.

## Close — 2026-08-11

**Outcome:** the two undocumented dev-setup repairs are in the guide, at the step where each is met, with
the symptom that identifies them. The third was already documented, and the re-survey is what established
that rather than a fourth restatement of it.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: **y (5 tiks) — but see below** — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**

**Why `continue` against a fired cap.** The 5-tik cap is a review checkpoint, and the review happened: the invocation was interrupted by an unrelated spend limit, the user lifted it and explicitly directed recovery *"from where you left off"* with a new hard stop at **14:20Z** and the same closed scope. That is the checkpoint's purpose served, in-band, by the only party entitled to serve it. Recorded rather than quietly re-graded — the cap DID fire, and continuing past it is a decision with an owner.

**Decisions:** `D-M257x-287-1` — a claim already stated twice can still be missing, because **placement
is part of a claim**. Both existing mentions of the `exit 0` failure sit upstream of the failure; neither
is where a reader is standing when it happens. Repeating a fact at the point of use is not duplication,
and the prose-twin fence built at iter-282 is what makes that safe to say — it grades copies that
*disagree*, not copies.

**Routes carried forward:** all of iter-285's and iter-286's, unchanged.

**Lessons:**
1. **Re-survey closed the first target of three.** Third consecutive run in which a route list had
   decayed; the check cost one grep.
2. **Placement is part of a claim.** Twice-stated and still absent, because neither statement was where
   the symptom appears.
