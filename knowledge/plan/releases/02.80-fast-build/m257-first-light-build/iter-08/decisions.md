# iter-08 — decisions

## D1 — the guard was fixed, not the prose; and the reasoning is auditable either way

**Decision.** Treat `platform-migration-status.md:121-122` as a **guard-resolution** defect and change
`anchor_construct_guard`, rather than edit the two flagged sentences.

**The test applied**, because "fix the guard" is the answer that lets a fence become theatre if it is
reached lazily:

| question | answer | evidence |
|---|---|---|
| is the cited content real? | **yes** | `app/main.go:1487` at `stack-demo/app` HEAD is `Sender: msgsender.NewFromEnv(logger)`; `docker-compose.yml` is 186 lines at `stack-demo/platform` HEAD and `:168` is `profiles: [frontend, all]` |
| is it real in a **declared clone root**? | **yes, both** | `_clone_roots()` returns exactly `[stack-demo, stack-dev]`; `stack-dev` has `origin/main == HEAD` and supports both citations at every rung |
| did the guard read something real? | **yes** — and that is the finding | it named its ref: `origin/main@0a9370c` / `origin/main@766df6c`, both **ahead** of their checkouts |
| did anything in this repo change to cause it? | **no** | the corpus text is untouched; a `git fetch` moved the ref |

**The deciding argument is the corpus's own rule**, not convenience: *cite the sha, never the moving
label.* A guard that grades an unpinned sentence at `origin/main` has picked a moving label on the
document's behalf. And the corrected behaviour is not new policy — it is iter-100's rule (*"defective only
if it names a non-construct at every ref its block offers"*) extended to the case where the block offers
no ref.

**What was NOT done:** widen the acquittal to the worktree (a fence you can green by editing the file under
test is not a fence); relax it for a caller-named `CITE_REF` (§5 rule 7); or add the two sites to the
postcondition baseline (the fence's own contract forbids it — *"a repair may remove these; it may never add
one"*).

## D2 — the sweep's scope is derived from what the publish SHIPS, not from habit

**Decision.** Before tagging, sweep the sections the 15 unpublished commits actually touch — derived from
`git diff --stat origin/main..HEAD`, not assumed.

**Why it is not simply "run everything".** `rosetta-extensions` carries 723 Go test files and 88 Python
test files across 11 sections; a full-tree sweep on a permanently-contended box is hours, and iter-07
already paid for the lesson that a long sweep **is itself contention** (its D2 killed one mid-campaign for
exactly that reason). The honest middle is a scope with a stated derivation: the touched set is
`stack-core`, `demo-stack`, `stack-injection`, `dev-stack`, `stack-verify`, `stack-seeding/isolation`,
plus README/knowledge files that no test covers.

Stated so it cannot be over-read: **this is not a whole-repo green.** It is a green over the sections this
tag changes, and the sections it does not change are unchanged from the tag that already passed them.

## D3 — the campaign's own admission gate may refuse this box, and that is a RESULT

`buildbench.run` calls `pre_rep_assert` **before** each rep and, when HEADROOM fails, **aborts the
campaign** (`return 1`, `D-M255-1`: *"a gate number measured on a host without headroom is not a number"*).
Clause 1 is an **instantaneous** `os.getloadavg()[0]` there — not the peak the post-rep assert takes.

The host sat at `load1` **16.30** at session open, before this session ran anything, and the corrected
limit is **10**. So the campaign can be refused at rep 1 without ever measuring. Per `TOK-02` step 3 that
is **a result to record with its `load1`**, not a failure to measure — and the sweeps above (which took the
box to ~39) are *my own* contention and must be finished and drained before launching, exactly as iter-07's
D2 reasoned: **contention I cause is contention I can remove.**
