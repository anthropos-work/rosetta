---
iter: 225
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-09
---

# iter-225 — can the sanctioned dev host even attempt gate clause 1?

**Type:** tik, under `TOK-08`, steered by the **user redirect of 2026-08-09** — this iter is on the
*working stack* half.

## Step 0 — re-survey before targeting (mandatory)

`TOK-08`'s standing direction is "work the mechanical classes"; the redirect names *which* work matters.
The live lead is **gate clauses 1 and 2** — a cold `demo-down --purge` + `demo-up` reaching
`autoverify green:true / 0 warnings` across 3 consecutive cycles, plus the full Playthrough suite on that
stack. Both are recorded **MET WITH A PERMANENT DISCLOSURE** (a freshly built stack failed the first full
run in 2 of 2 attempts, so there was never a clean pass), closed ~200 iters ago at platform `0c91421`,
**never re-attempted**.

Two iters have just removed the reasons not to re-attempt it:

- **iter-223** proved all 23 demopatch anchors still apply at `origin/main` — advancing the pin is safe.
- **iter-224** advanced every stale clone to its origin tip, closing
  `ROUTE-M257x-222-other-clones-never-fetched`.

**A cold bring-up is expensive and may not fit one iter. Sizing it honestly is a legitimate iter;
fabricating it is not.** This iter sizes it, and it asks the cheapest disqualifying question first.

## Cluster / target identified

**`D-v28-15` retired both hosts the build budget was measured on.** It states: `billion` is the OFFICIAL
host, **demo deployment only, not for development or testing**; dev/test is **LOCAL to the new Mac**; the
old laptop and `odysseus` are **both retired**.

`rosetta-extensions/stack-core/hostprofiles/` contains exactly **two** profiles: `billion.json` and
`laptop.json`. `build-budget.md`'s **headroom contract** has a **clause zero** — `require_measured` —
described in the corpus as the clause *"a fresh host with an unpopulated sampler hits FIRST"*, graded
against **measured, checked-in host profiles**.

So the question that gates every other bring-up question: **is there a measured profile for the only host
the release is allowed to develop and test on?** If there is not, clause 1 cannot be *graded* here, and
no amount of bring-up work changes that.

## This host, measured before any profile was opened

| property | value |
|---|---|
| model / arch | **`Mac16,11`**, `arm64` (Darwin 25.5.0) |
| CPU | **12** logical, 12 physical |
| RAM | **25,769,803,776 B = 24 GiB** |
| free disk on `/` | **208 GiB** of 460 GiB |
| Docker storage driver | **`overlayfs`** (Docker Desktop) — *not* billion's containerd |
| Docker VM | **8 CPU**, **12,528,664,576 B = 11.67 GiB / 12.53 GB** |

## Hypothesis

The two checked-in profiles describe `billion` and the **retired** laptop, so the sanctioned dev host has
no measured profile, and the headroom contract's own clause zero fires before any of the CPU / memory /
disk clauses are reached.

## Expected lift

A **decidable answer** to "is clause 1 attemptable on the sanctioned host, and at what cost", with every
number derived here and none carried. No `N`/`P` reading is claimed.

## Phase plan

1. **Seal predictions** (this commit — `probe(M257x/225)`), before opening either host profile.
2. Read both profiles; decide whether either describes this host.
3. Run `buildbench`'s headroom check on this host; record the verdict verbatim.
4. Check the demo UI tier's Docker-VM prereq against this host's actual VM.
5. Record the sizing; repair whatever the corpus states about hosts that `D-v28-15` has retired.

## Escalation conditions

- If clause 1 turns out to be **attemptable and cheap** on this host, that is a scope question for the
  next iter, not a mid-iter pivot — this iter still closes on its sizing deliverable.
- If a profile must be *authored* to unblock the clause, that is a distinct line of work and routes
  forward rather than landing here.

## Acceptable close-no-lift outcomes

**Finding that `laptop.json` DOES describe this host is a first-class result** — it would refute the
hypothesis and make clause 1 gradeable today, which is the more useful of the two answers.

## Pre-registered predictions — SEALED IN THIS COMMIT, BEFORE EITHER PROFILE IS OPENED

| id | prediction | rationale |
|---|---|---|
| **P-225-1** | **`laptop.json` does NOT describe this host** — its CPU count and/or RAM differ from 12 / 24 GiB — so **neither** checked-in profile describes the sanctioned dev host | `D-v28-15` calls the laptop *retired* and the Mac *new*; the profiles are dated 2026-07-31, before the switch |
| **P-225-2** | `buildbench`'s headroom check **refuses to grade** on this host, citing its `require_measured` clause zero | that is the clause the corpus says a fresh host hits first |
| **P-225-3** | `build-budget.md` names **`billion` and `odysseus`** as its baseline/gate hosts and **names no Mac** — so the doc's host set is **disjoint** from `D-v28-15`'s sanctioned dev host | the doc predates `D-v28-15` |
| **P-225-4** | the demo UI tier's **12 GB Docker-VM prereq** is **within 10 %** of this host's actual VM size (12.53 GB / 11.67 GiB) — i.e. this host is **on the boundary**, not comfortably above, and the verdict depends on whether the check reads GB or GiB | measured above, before reading the check |

**If P-225-1 is refuted, clause 1 is gradeable today and the sizing becomes a cost estimate instead of a
blocker** — that is the better outcome and the iter reports it as such.
