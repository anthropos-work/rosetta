---
milestone: M257x
iter: 02
---

# iter-02 — decisions

## D-M257x-4 — TOK-01's step-1 premise did not survive the machine move (refuted by measurement)

`TOK-01` opens: *"Unblock the gate's instrument. The rext pin is inconsistent and `ensure-clones.sh` is FATAL
on it, so `/demo-up` aborts on this box today."* Its **Next-tik direction** names `iter-02 =
FIX-M257x-rext-pin` on **odysseus**.

Measured on the new Mac (`D-v28-15`):

| TOK-01 claim | measurement |
|---|---|
| `rext.tag` = `cockpit-deeplinks-v1`, 63 commits behind `main` | `rext.tag` = `fast-build-m257x-iter-01` |
| the tag may exist only locally | **on origin** — `31d2b5df… refs/tags/fast-build-m257x-iter-01` |
| the pin is N commits behind | `31d2b5df` **== `origin/main`**; 0 behind, 0 ahead |
| the consumption clone is at a *different* tag ⇒ FATAL | **no `stack-demo/` exists** — the mismatch is unrepresentable |

**The pin was clean before this iter started.** TOK-01 was authored against the old box, where the SoT was 63
commits behind; the machine move re-created it deliberately and correctly. The step was not *wrong* — it was
**overtaken**.

Recorded rather than silently skipped because this milestone exists to end "re-derived from scratch each
time": a future reader comparing TOK-01 to the iter ledger would otherwise see step 1 apparently unexecuted.

**What replaced it as the real blocker,** measured in the same pass: **no container runtime of any kind** on
this box (docker · podman · colima · nerdctl · lima · orbstack), plus no `gh`/`psql`/`tailscale`. Docker
Desktop was installed by the user **mid-iter** (independently re-verified: `29.6.2 · linux/arm64 · overlayfs ·
8 cpus · 12528664576 B`). Clause 1 remains unmet — there is still no `stack-demo/` workspace — but it is now
unblocked-in-principle. Routed as `HOST-M257x-stack-demo`.

**Consequence for the strategy chain:** TOK-01's ordering (*"the bring-up is the instrument for clauses
1/2/4, so an aborting bring-up is the first thing to fix"*) still holds; only its named first target is
retired. The substitution to step 2 is minimal deviation, not a re-scope.

## D-M257x-5 — the residual schema set is DECLARED DEBT, not design

The obvious move — delete `cms`/`jobsimulation` from the CREATE SCHEMA list, since `repos.yml` no longer
declares them — was **measured before being taken**, per §7.1, and rejected for this iter.

`stack-seeding/cmd/stackseed/main.go:45-105` holds **live** entries for ~12 `jobsimulation.*` tables (9
written). Removing the schema creation today would convert a working-but-wrong bring-up into a
**knowingly-broken** one, on a box where no bring-up can yet be run to verify either state.

`D-M257x-1` established that loud failure is the safe mode. It is safe **once the writes have somewhere to
go**; before the re-point there is no canonical target and the loudness is just breakage. So:

- the set is **explicit**, in one place, with a per-entry reason (`REXT_TRANSITIONAL_SCHEMAS`);
- it is **fenced** — `test_transitional_debt_may_shrink_but_not_grow` fails if it grows, and *also* fails
  (with "this failure is GOOD NEWS") when it shrinks, so paying the debt down forces the win to be locked in;
- `skillpath` is **not** in it — absent from `repos.yml`, zero rext writes, the §2 canary; it is simply gone.

This generalises the protocol's rule. *"Derive it, or fence it"* has a gap: some entries can be neither
derived (`sentinel` is `migrations: false` and alive — Trap A) nor deleted (the writes still exist). The
third clause is **declare it, with a reason and a no-growth fence**.

`REPOINT-M257x-jobsim-writes` is the handler that retires this debt and unblocks gate clause 4.

## D-M257x-6 — a source-grep fence cannot tell code from prose

Two independent instances in one iter:

1. **`test_migrates_the_four_merged_services_and_never_skiller`** (v2.1 M209) required all four pairs to be
   present in `migrate-dev.sh`. After the loop was derived it **still passed** — the tuple now appears in the
   *comment* explaining why it was removed. A test whose whole purpose was to pin the migrate set was
   satisfied by its own refutation.
2. **The new fence's own prose fixture** initially could not fail (see `progress.md`), because it placed the
   lying comment where the parser resets past it.

Both are the M256 class — a check reporting success without checking — and both were found only by
**mutating the subject and demanding RED**.

**Rule adopted:** a fence over source must assert against a **parsed construct** (the loop body, the derived
value, the AST node), never a whole-file substring — *unless* the check is explicitly about prose. The
existing `test_dropped_mirror_fence.py` already does this correctly by scoring the **occurrence** and
excluding comment lines; that design is the precedent, and it was not followed by the M209 test.

Folded into `corpus/ops/platform-alignment.md` §8 as a fence-design constraint.
