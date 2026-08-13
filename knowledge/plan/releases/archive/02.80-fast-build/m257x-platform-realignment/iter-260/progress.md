# iter-260 — progress

**Type:** tik
**Active strategy:** `TOK-08`, under the user binding `D-M257x-256-1`.

Per the protocol's cadence section, this iter's pre-registrations are sealed in its FIRST commit,
before any teardown.

## Phase A — sealed

Pre-registrations PR-1…PR-5 sealed at `b7795b0`, before any teardown. The seal commit also carries the
**downward correction of the clause-1 count** (0 of 3 strict, not 1 of 3) — recorded before it could be
influenced by whether the cycles went green.

## Host preconditions (stated, per §9 — an iter must not rediscover these)

| fact | value at 2026-08-10T14:23:27Z |
|---|---|
| disk free | **196 GiB** (204 GiB after cycle-1 purge reclaimed this stack's images) |
| Docker VM | **11.67 GiB / 8 CPU** |
| load avg | **5.79 / 7.13 / 8.26** — CONTENDED, third-party |
| `demo-1` | up 4 days, **11 containers, out of scope under every outcome** |

**Every duration below is CONTENDED and is NOT a host baseline.** The driver here is `overlayfs`, not
containerd, so `build-budget.md`'s `billion` figures (666.29 s p50; the 46.2 % image export/unpack leg) do
not transfer and no deviation from them is reported as a finding.

## Phase B — cycle 1 of 3 (the first `--purge` + `up` in this milestone's evidence)

`rosetta-demo down 2 --purge` → `up-injected.sh 2`, bare, no flags — byte-identical invocation to iter-258's.

| | |
|---|---|
| teardown | `DOWN_EXIT=0`; **0** `demo-2` containers remain; this stack's images removed (disk 196 → 204 GiB) |
| bring-up | start `2026-08-10T14:25:08Z` → end `2026-08-10T14:33:57Z` = **529 s CONTENDED** |
| verdict | `{"project":"demo-2","offset":20000,"warnings":0,"green":true,"ts":"2026-08-10T14:33:57Z"}` |
| exit | `EXIT_CODE=0` |
| asserts | 15/15 ✓, including `public.skills = 42790`, `casbin_rules = 1251`, all 11 containers live |

**PR-2 held twice** — `demo-1`'s 11 containers are identical in name, status **and container ID** both after
the teardown and after the bring-up (`demo1-baseline.txt` vs `demo1-after-c1down.txt` / `demo1-after-c1up.txt`,
`diff` clean), and demo-1's cockpit `:17700` (pid 75363) stayed bound throughout.

**PR-5 held through its first real test.** The teardown freed **both** host-native listeners — after
`down --purge` only demo-1's `:17700` remained — and the bring-up re-bound exactly one process on each with
**new pids** (`:27700` 43878 → 67825; `:23077` 63229 → 2701). No orphan, and the M217 leak class did not fire.
This is the half of clause 1 that iter-258's fresh-slot bring-up could not exercise at all.

## Phase C — cycles 2 and 3, identical invocation

| cycle | teardown | bring-up window (real paired `date -u` reads) | seconds | verdict | exit | ✓ asserts |
|---|---|---|---|---|---|---|
| 1 | `DOWN_EXIT=0`, 0 containers left | `14:25:08Z → 14:33:57Z` | **529** | `warnings:0, green:true` | 0 | 15 |
| 2 | `DOWN_EXIT=0`, 0 containers left | `14:35:06Z → 14:42:44Z` | **458** | `warnings:0, green:true` | 0 | 15 |
| 3 | `DOWN_EXIT=0`, 0 containers left | `14:43:28Z → 14:51:46Z` | **498** | `warnings:0, green:true` | 0 | 15 |

Three consecutive `rosetta-demo down 2 --purge` + bare `up-injected.sh 2`, no flags, no retries, **no cycle
re-run to obtain a nicer number**. Raw `autoverify.json` for each is in `evidence/c{1,2,3}-autoverify.json`.

## Phase D — grading the pre-registrations

| | prediction | outcome | evidence |
|---|---|---|---|
| PR-1 | strict clause-1 count before this iter is **0 of 3** | **HELD** | clause text names `demo-down --purge` + `demo-up`; iter-258 ran a bare fresh-slot `up` and says so in its own close |
| PR-2 | `demo-1` untouched | **HELD, six times** | `diff` clean on name+status+**container ID** after each of 3 teardowns and 3 bring-ups; `:17700` (pid 75363) bound throughout |
| PR-3 | all three `green:true / warnings:0 / EXIT_CODE=0` | **HELD** | table above; 45/45 asserts ✓ |
| PR-4 | times land in a **broad band around iter-258's 717 s** | **SPLIT — band clause REFUTED** | the three cluster **tightly** (458–529 s, spread **71 s**) and **all three sit below 717 s**. The predicted band does not contain the observed one |
| PR-5 | no host-native listener orphan across three cycles | **HELD — and this is the half iter-258 could not test** | each teardown freed `:27700` + `:23077`; each bring-up bound **exactly one** process on each, pids rotating `27700: 43878→67825→17970→95381`, `23077: 63229→2701→73195→94551`. Zero orphans; M217's leak class did not fire |

### PR-4's refutation, and the attribution I am NOT making

The three cycles are consistently ~190–260 s faster than iter-258's fresh-slot 717 s, and far tighter among
themselves (71 s) than any of them is from 717. It is tempting to read that as *"repeat cycles are faster
because the cache is warm."* **That inference is not available from this data and is not made.** The host is
permanently CONTENDED by third-party load (load avg 5.79/7.13/8.26 at open), n=1 for the 717 s point, and
nothing here isolates cache state from contention. What is defensible is only the negative: **717 s was not a
usable expectation for the next cycle**, which is precisely why a timing is not a baseline. Per §9, wall-time
on this host is not a usable measurement; the **counts** are, and the counts are 3/3 and 45/45.

**Driver is `overlayfs`, not containerd** — `build-budget.md`'s `billion` figures do not transfer and no
deviation from them is reported.

## Close — 2026-08-10

**Outcome:** **Gate clause 1 is MET under its literal unit** — three consecutive
`demo-down --purge` + `demo-up` cycles on the advanced refs, each `green:true / warnings:0 / EXIT_CODE=0`,
45/45 asserts, with `demo-1` provably untouched throughout. The iter also **corrected the clause's own count
downward before measuring** (0 of 3, not the briefed 1 of 3), so the clause is met on the strict reading and
needs no adjudication.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue

**Why the milestone gate is still NOT met.** Clause 1 is one of five. **Clause 2** (the Playthrough suite on
this stack) has never run — it is the next iter. Clause 5 is unmeasured since iter-131. And the user's binding
`D-M257x-256-1` needs the **dev** half, which remains with the user.

**Decisions:** `D-M257x-260-1` (iter-259's `367` retracted; the escalation survives on staleness, not data
loss) — a **side-deliverable**, committed separately at `dbe5c91` so it does not blur this iter's status.

**Side-deliverables:** the `D-M257x-260-1` retraction + the protocol widening (*"cannot push from here" is not
"exists only here"*, `platform-alignment.md` §8). Raised by the orchestrator mid-iter; re-measured rather than
relayed, and the re-measurement went further than the correction supplied — the clone has **zero `origin/*`
refs**, so `--not --remotes` had been subtracting a **bundle**.

**Routes carried forward:**
- `ROUTE-M257x-260-clause-2-never-run` → **new.** The full Playthrough suite has never run on a stack built
  from the advanced refs. The pinned rext carries all **10** product manifests, so it is viable at this pin.
  Handler: `FIX-M257x-260-run-the-playthrough-suite`. **Next iter.**
- `ROUTE-M257x-258-no-dev-stack-on-this-box` → open, **with its rationale replaced** per `D-M257x-260-1`:
  staleness, not data loss. Still the user's decision.
- `ROUTE-M257x-258-the-pin-is-157-iters-stale`, `ROUTE-M257x-257-lock-file-is-unfenced`,
  `ROUTE-M257x-256-mixed-ref-anchors` and all earlier → unchanged and open.
- **New, from the retraction:** `FIX-M257x-260-audit-remote-claims-for-bundle-clones` — iter-259 is the only
  known site, but no fence enumerates *"a push-state claim derived from `--remotes`"* anywhere in the corpus
  or the milestone record.

**Lessons:**
1. **Correct a gate count in the expensive direction before you measure, not after.** The briefed count was
   1 of 3; the clause's own wording gives 0 of 3. Sealing that correction in the first commit meant the
   result could not be shaped by how the cycles went — and as it happens they went green, so the strict
   reading cost nothing but would have been unarguable if they had not.
2. **A zero census must prove its instrument even when the zero is the answer you were hoping for.** The
   retraction's *"0 commits absent from origin"* was the reassuring direction, and it was void. The
   anti-vacuity control is what stopped a wrong correction from replacing a wrong claim. *(Generalised into
   `platform-alignment.md` §8 in the same commit as the fix.)*
3. **The half a happy path cannot exercise is the half that leaks.** iter-258's fresh-slot bring-up proved
   `up` and could not touch `down --purge` — where M217's orphan class lives. Three teardowns were needed
   before PR-5 meant anything.
