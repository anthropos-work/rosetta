**Type:** tik — under `TOK-01`, on the user's *"i want no debt"* ruling.

# iter-20 — the last RED fence, and the anchors no fence can see

## Phase A — G1: a discriminator, not a waiver

`platform_predicate_guard` was the only RED fence left, on
*"`docker-desktop-vm` is documented as a profile at 2 site(s) but no service declares it."*

**"profile" is not a compose word.** M255 gave this corpus **host** profiles
(`stack-core/hostprofiles/*.json`), and `_PROSE_PROFILE` had negation, postfix-negation and ref-pin
discriminators but **no domain discriminator** — so every *"a `X` profile"* was graded as a compose
token. The finding was accurate and useless: it described what `--profile docker-desktop-vm` would
do, about a command nobody would type.

`_HOST_PROFILE_DOMAIN` added, same block window and same evidence standard as the other two, and
**deliberately narrow**: it requires the window to *name* the other domain, and does **not** exempt on
the bare word "host" (`--public-host`, `host-gateway`, `STACK_PUBLIC_HOST` are all compose prose).

**RED-proven, three tests** — the host-profile forms stop being graded · three compose claims one
clause from the word "host" are **still** graded · an always-true mutant launders a live `cms` claim,
proving the exemption is what decides these cases. `190 passed`.

One site (`build-budget.md:197`) still fired afterwards, and the reason is the **window**, not the
rule: *"host profile"* sits two lines up across a **blank line**, which `_pin_window` treats as a
block boundary by design. The sentence now names its domain inline (`D105`) — an improvement to
ambiguous prose, not prose bent to a wrong fence, and recorded as such because the two look identical
from an exit code.

## Phase B — the anchors no fence can see

`ROUTE-M258-iter18-app-row-anchors-are-at-2035f9a` discharged. The map's `app` row pinned seven
`app/main.go` anchors at `2035f9a`; `origin/main` is `c52dbc51e`, and **all seven had drifted** by
12–20 lines. Re-resolved, plus the net-new eighth:

`customeriosync.New` 395→**396** · `internalstorage.NewManager` 524→**537** ·
`skiller.NewSkillerManager` 690→**706** · `jobsimwiring.Wire` 721→**734** ·
`skillpath.NewSessionManager` 751→**764** · `appcms.Wire` 1153→**1167** ·
`msgadapters.Wire` 1471→**1458** · **`sentinel.Open` = 305 (net-new)**.

**8/8 verified by reading the target line.** The finding is that nothing caught this: `app/main.go:NNN`
citations are graded **range-only** — a Go file has no block structure to attribute a line to — and a
1,635-line file swallows a 20-line slip (`D106`).

## Phase C — gates

**Eleven fences, all `rc=0`** — the full corpus set, including the two that were RED at this iter's
open and close.

Five rext test modules run: `test_platform_predicate_guard` (190 passed), `test_predicate_enumerator`,
`test_service_doc_status_fence` green. Two failures, **both verified against iter-18's measured
pristine-HEAD baseline as pre-existing** (`test_fence_provenance`, `test_guard_family_verdict_line_m257x`)
— the `guard-scans-its-own-scratch` family that fires on any box that has run a demo.

## Phase D — ship

`fast-build-m258-iter-20` tagged, pushed, and **verified on origin** (`git ls-remote` → 2 refs).
`.agentspace/rext.tag` re-pinned in the same iter — the `D71` half-re-pin class, not repeated.

## Close — 2026-08-12

**Outcome:** **The corpus fence set is fully green — eleven of eleven — and the last RED was the
fence's defect, not the corpus's.** G1 was reading **host** profiles as compose profiles; the repair
is a third discriminator, symmetrical with the two already there, **RED-proven with an always-true
mutant** so a loosening cannot pass as a fix. And the `app` row's seven wiring anchors — every one
drifted 12–20 lines at the new `app` ref, and **invisible to every fence** because a Go file has no
blocks to attribute a line to — were re-resolved and verified 8/8 by reading the target line, with
`sentinel.Open` added as the eighth domain.

**Type:** tik
**Status:** closed-fixed
**Gate:** N/A — the milestone's gate closed by user ruling (`D52`). Clause 3 remains NOT MET; this
iter took no timing measurement and offers none.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n *(3 tiks)* — (6) protocol-stop: n — (7) budget-exhausted: **y** *(between iters,
tree clean — three tiks spanning a corpus-wide sweep with a full-suite pristine attribution, a cold
bring-up + verification + teardown, and a fence repair with mutants; every remaining route needs
either fresh host time on a quiet box or its own measurement campaign)* — Outcome: **exit-7**

**Decisions:** D104–D107

**Side-deliverables:** none — both targets were the iter's planned scope.

**Routes carried forward:**

- **The ~46 pre-existing rext-internal test failures** stay open and are now **measured** rather than
  assumed (iter-18 `D96`): `suite_census`, `frozen_expectation_census`, `anchor_subject_census`,
  `service_registry_guard`, `repair_*` and the mutation batteries. The
  `FIX-M258-iter03-guard-scans-its-own-scratch` family is the known root of part of it — *a guard that
  rglobs from the repo root censuses the demo stack's ephemeral platform clone*. **`D107` states the
  scope of the green explicitly** so "eleven fences green" is never read as "the suite is green".
- **`ROUTE-M258-iter19-orphan-images-outlive-their-service`** — `anthropos-sentinel:latest` plus five
  `:probe` leftovers. Price the shared layers before quoting any of it (`D53`).
- **`ROUTE-M258-iter19-studio-desk-frontend-port-is-not-published`** — a demo publishes studio-desk's
  backend port only.
- Unchanged and not re-verified: `FIX-M258-iter14-purge-leaves-276MB` (second data point at iter-19) ·
  `TARGET-M258-iter13-browser-only-deps` · `SETTLE-M258-iter13-studio-desk-cold-time` (**still
  unmeasured**) · `ROUTE-M258-iter13-dockerfile-not-in-cache-key` ·
  `REPORT-M258-iter17-public-host-default-skips-the-batch` ·
  `REPORT-M258-iter17-dev-ui-images-stay-pre-L1-fat` · `ROUTE-M258-iter17-batch-gate-has-no-dev-opt-in` ·
  `ROUTE-M258-iter17-registry-is-empty-while-a-stack-is-up`.

**Lessons:**

- **Grade the fence before you edit the prose.** The exit code looks the same either way; only the
  diagnosis distinguishes a corpus defect from a detector defect, and this one was the detector's.
- **Repair a detector with a mutant, or you have loosened it.** The always-true mutant is the whole
  proof: without it, a discriminator that never matched would have passed both other tests by
  accident.
- **A blank line is a block boundary, and that is a feature.** The one site the discriminator did not
  reach was relying on context two paragraphs up. Widening the window would have weakened a rule that
  exists to be narrow; naming the domain inline made the sentence self-contained.
- **Range-only is not graded.** Seven anchors drifted 12–20 lines inside a 1,635-line file and no
  fence could have said so. Anchors into block-less files must be re-derived at every ref bump.
