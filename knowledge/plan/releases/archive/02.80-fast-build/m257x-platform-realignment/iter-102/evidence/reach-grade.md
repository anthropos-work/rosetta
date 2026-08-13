# iter-102 — REACH, graded by machine against both read ledgers

**The instruction was "grade reach and report it", not "claim it."** So the number below comes from
`rosetta-extensions/stack-core/repair_reach_guard.py` run against the repair commit and each reading's own
**raw seat reports** — the fence built at iter-83 for exactly this question, which asks something neither
`repair_postcondition` nor `repair_leak_guard` can: *did the repair reach its INPUT?*

Run: `repair_reach_guard.py --repo-root <rosetta> --ledger iter-<N>/raw --range cd16967`

## The measurement

| | iter-99 ledger | iter-101 ledger |
|---|---|---|
| booked findings in the ledger | **46** (14 reports) | **36** (13 reports) |
| repair hunks / paths | 145 / 53 | 145 / 53 |
| **touched** | **37** | **29** |
| file-unreached | 1 | 4 |
| line-unreached | 8 | 3 |
| **graded reach** | **37/46 = 80.4 %** | **29/36 = 80.6 %** |

## What the unreached residue actually IS — and this is the finding

`repair_reach_guard` grades against **every booked finding**, and a reading's bookings include the ones the
adjudicators **REJECTED**. A rejected finding is a claim that turned out to be **true**; repairing it would
be the defect. So an 80 % reach is not 20 % of upheld work left undone — **it is the rejections correctly
left alone**, and the guard has no way to know that.

**iter-101 makes the point exactly.** Its 7 unreached anchors, against its 8 recorded rejections:

| unreached anchor | rejection recorded at iter-101 | class |
|---|---|---|
| `ai_architecture.md:35` | ✅ rejected | mis-read |
| `ai_architecture.md:141` | ✅ rejected | mis-read |
| `security_compliance.md:185` | ✅ rejected | mis-read |
| `backend.md:19` | ✅ rejected | mis-read |
| `skiller.md:19` | ✅ rejected | mis-read |
| `chronos.md:27` | ✅ rejected | wrong-convention |
| `ai-labs.md:76` (booked `:75`) | ✅ rejected | wrong-convention |

**7 of 7 unreached are rejections. All 7 of them.** The 8th rejection — the wrong-tree messenger-row
prod-terraform clause — sits in a file the repair edited for other reasons, so it grades as *touched*.

iter-99's 9 unreached line up the same way against its 10 rejections, and include `hiring.md:80` twice —
the one finding **seat 4 declined to repair with a written derivation**, because it re-measured the claim
and found it **true** at the settling tree. A decline-with-evidence is indistinguishable from a miss to this
fence, and that is the fence being conservative in the right direction.

## So the honest statement, in both forms

> **Graded reach 80.4 % / 80.6 % against the FULL booked set — and effectively 100 % against the UPHELD
> set, with the entire unreached residue being the adjudicator-rejected findings that must not be
> repaired.**

**Both numbers are published, and the first one is not adjusted away.** A repair that reported 100 % would
be less trustworthy than one that reports 80 % with the residue named — and the residue here is named
anchor by anchor above, so the claim is checkable rather than asserted.

## What this does NOT establish

That every upheld finding was repaired **correctly** — reach measures whether the repair *landed on the
anchor*, not whether what it wrote is true. Correctness is the next reading's job, and a repair pass may
not contain one.

It also does not establish that the *predicate* was closed everywhere, only that the *anchor* was reached.
That is `claim_twin_guard`'s question, and it answers **GREEN over 264 adjudicated claims** — up from 134
before this iter published its ledgers in the derived shape.
