# iter-84 adjudication — `FIX-M257x-iter82-reread-union`

**Status: COMPLETE.** All **43** distinct anchors of iter-82's re-read union graded, four parallel
adjudicators over disjoint packets, each re-deriving from the clones rather than from any seat's note.

> **This is the PER-ANCHOR ledger** that `FIX-M257x-iter83-adjudication-has-no-per-anchor-ledger`
> demanded. iter-76 recorded rejection *mechanisms* and *counts* but never per-anchor verdicts, which is
> why iter-83's reach measurement could only be graded against **booked** rather than **upheld**. That
> gap is not repeated here: every anchor below carries its own verdict, severity and predicate, so a
> future reach measurement has an `upheld` denominator to grade against.

## Verdict

| | A | B | C | D | **total** |
|---|---|---|---|---|---|
| anchors | 11 | 11 | 11 | 10 | **43** |
| **UPHELD** | 10 | 10 | 11 | 9 | **40** |
| — blocker | 9 | 8 | 8 | 6 | **31** |
| — minor | 1 | 2 | 3 | 3 | **9** |
| REJECTED | 1 | 1 | 0 | 1 | **3** |
| UNSETTLED | 0 | 0 | 0 | 0 | **0** |

**Upheld rate 93.0 %** — against iter-80's **92.1 %** on the pre-repair union. The pre-registered floor
was ≥ 70 %; a collapse below 50 % would have meant the instrument's post-repair signal was mostly noise.
**It is not. The instrument did not regress across the repair.**

> **Counts re-derived from the per-anchor verdicts, not from the adjudicators' own summary lines**
> (§5 rule 32). This was not ceremony: **adjudicator B's summary reported "9 UPHELD (7 blocker, 2
> minor)" while its own eleven verdicts give 10 UPHELD (8 blocker, 2 minor)**. A hand-off's numbers are
> re-derived including the orchestrator's — the fifth time this milestone has paid for that rule, and
> the second time in this run alone (`D-M257x-83-9` was the first).

## The three rejections — all one mechanism, and it is now on its 5th occurrence

| anchor | why rejected |
|---|---|
| `CLAUDE.md:203` | `service_desired_count` graded against clones **6 and 3 commits behind `origin/main`**. At `storage 63bffc8:38` and `messenger a0ec933:29` both read `= 0`, exactly as cited. **The corpus is correct.** |
| `academy-backend.md:20` | the pin is **inside the same phrase** (*"currently `v1.363.2` @ `5ba17044`"*), not two sentences away — rule 33's trigger is not met. `app tag --points-at 5ba17044` → `v1.363.2` ✓; `:55`'s `9d00a313 = v1.367.0` is a different ref-relative statement, also true. Adjacency is not co-reference |
| `hiring.md:73` | `hiring.md:17` declares the document's grounding as **`app @ 5ba17044`**; both disputed anchors are byte-exact there (`manager.go:448` `switch org.IsHiring {`, `:485` `if !org.IsHiring {`). The seat measured the newer checkout `b948604` — three versions ahead |

**`CHECK-M257x-iter76-seat-ref-discipline` is now at its 4th and 5th occurrences** (iter-82 was the 3rd).
The escalation condition declared in this iter's `overview.md` has **fired**.

**And adjudicator D found the reason the rule keeps failing — it is stated wrong.** The seats are not
ignoring it; they are applying it *unevenly*, because "grade at the ref the claim names" does not tell
you what to do when the sentence claims **currency**. The correct statement is:

> **Grade at the ref the claim names — UNLESS the sentence asserts currency, in which case no
> neighbouring pin rescues it.**

That is why `graphql-wundergraph.md:13` is **UPHELD** (it says *"survives"*, *"is now"*, under a table
column headed *"origin HEAD"*) while `hiring.md:73` is **REJECTED** (a static code citation under an
explicit document-scope grounding banner). Both were graded by the same rule, correctly, once the rule
was stated with its exception. Routed as a §5 rule-33 amendment.

**Second structural recommendation, from adjudicator A:** the ground-truth table handed to seats should
carry **each clone's `origin/main` sha next to its checkout sha**. A seat given only the checkout has no
way to see that it is stale, which is precisely how occurrences 3–5 happened.

---

## The 40 upheld, by predicate

The repair unit is the predicate (§5 rule 19 + `D-M257x-59-1`), so the ledger is grouped that way. **This
is iter-85's work list.**

### Q1 — stale cross-repo line anchor (13)

The largest class, as iter-82 predicted. Every one points *out* of the corpus into a repo, and in almost
every case **a sibling anchor in the same sentence is exact** — which is what lets the wrong one survive
a reading.

| anchor | correction |
|---|---|
| `ai-readiness.md:213` | drop the *"still says four"* clause; re-anchor `:185-189`, `:149` |
| `ai-readiness.md:214` | `:181-185` → `:184-190`; `:145` → `:149` |
| `ai-readiness.md:215` | `:145` → `:149` (quote attributed to a line that says something else) |
| `ai-readiness.md:219` | **delete** — the rext source was *fixed*; the refutation expired |
| `ai-readiness.md:321` | `AIReadinessClient.tsx:69` → `:78` (used at `:599`) |
| `ai-readiness.md:322` | same construct, same fix |
| `ai-readiness.md:465` | `:137-138` → `:153-154`; `:150-154` → `:166-170` |
| `backend.md:43` | `app/CLAUDE.md:72` → `:80` (`:72` is a closing code fence) |
| `skillpath.md:34` | `app/CLAUDE.md:72` → `:80` — **the same defect in a second file** |
| `external_services.md:208` | → `gen_injected_override.py:84` + `:669-670` |
| `external_services.md:217` | `test_injection.py:1051` → `:1108` |
| `external_services.md:218` | `test_injection.py:1051` → `:1108` — **same defect, adjacent line** |
| `storage.md:8` | `anticheat.go:34` → `internal/jobsimulation/anticheat/anticheat.go:30` (wrong **directory** and line) |

**⚠️ The `ai-readiness.md` cluster is ONE rext commit.** `4e6b64d`
(*"fix(stack-seeding): the usage-KPI comments said FOUR; the function emits FIVE"*) was authored **13
seconds after** the corpus repair `328ece5` and shifted the file **+4 lines**. So `:219` is not merely
stale — it is a **refutation the fix inverted**: the corpus says rext *"still says four"*, and rext now
says five. `CHECK-M257x-iter77-cross-repo-pin` measured, twice over.

### Q2 — a present-tense claim about a fact that was DELETED (7)

**Re-anchoring is NOT the repair for any of these.** Restate or drop (§4 Trap A).

| anchor | what is actually true |
|---|---|
| `graphql-wundergraph.md:13` | no `graphql` profile at all; `PROFILE ?= core` — **the run's centre** |
| `cms.md:8` | cms was **not** the last fold — v9.0 folded `storage` and `messenger` after it |
| `backend.md:218` | nothing publishes to `skiller`; `app` only **subscribes**. Also at `backend.md:25` and root `CLAUDE.md` |
| `roadrunner.md:113` | jobsimulation does **not** consume the event — replaced by an Asynq task |
| `services/README.md:37` | **`MESSENGER_RPC_ADDR` exists in no build, and never did** (`git log -S` over all platform history → 0 commits). Also at `messenger.md:7` |
| `architecture_overview.md:295` | `storage` prod is `= 0`; app has 0 reads repo-wide. The retraction at `:311` scopes itself *"locally"* and thereby affirms a dead prod edge |
| `alignment_testing.md:360` | unmeasurable is **rc=3** with a refuses-to-be-mistaken banner since M219, not rc=2 |

### Q3 — wrong scalar / wrong set against source (8)

| anchor | measured |
|---|---|
| `alignment_testing.md:463` | **six** snapshot operators, not five (`snapshot-cross-surface-closure` missing); `:482`'s content-surface count is 5, not 4 |
| `frontend_architecture.md:39` | `NEXT_PUBLIC_BACKEND_API_URL`: **~21 files / 29 call sites**, not "~15" |
| `shared_libraries.md:181` | *"every Go service repo the platform has ever cloned"* = **11**, not 7 |
| `architecture_overview.md:74` | **four** imported modules; `authn` is required by no `go.mod` (0 hits across 7 clones; control: `colony` in all 7) |
| `platform-migration-status.md:76` | the residual 8 is 3 `t.Setenv` + **1 `t.Fatal` string**, not 4 `t.Setenv` |
| `backend.md:77` | `SKILLER_RPC_ADDR`: **1** occurrence at `0dab54d`, 4 at `0808b92` |
| `cms.md:67` | the Python engine has **no Mistral path** (`openai`/`azure`/`anthropic` only); Mistral is Go-side, OCR-only |
| `storage.md:26` | the **env-var names** are constants, not the bucket names. Same sentence in `platform-migration-status.md:76` |

### Q4 — wrong predicate, no line-checker could catch it (7)

| anchor | defect |
|---|---|
| `gotenberg.md:7` | wrong **purpose** — the PDF is a throwaway text-extraction/OCR intermediate, never stored or displayed |
| `ant-academy.md:227` | incomplete **enumeration** — a mounted **7-locale dropdown** (`LanguageSelector.jsx`, in `TopBar` across the authed shell) the doc's verdict depends on not existing |
| `service_taxonomy.md:38` | **reversed** diagram edge — generation flows Desk → Backend → Room, not Room → Desk |
| `hiring.md:52` | the scoreboard is **not** reachable in `apps/web` for a genuine hiring org — `UserStatusContext.tsx:141-172` ejects the recruiter. Retracted 3× later in the file, never in place |
| `dependency_map.md:58` | *"no second subscriber on the `backend` stream locally"* — refuted twice: `messenger/internal/flow/flow.go:72`, and `app`'s own consumer group at `main.go:1442-1451` |
| `external_services.md:821` | **retirement over-sweep** — `storage` is declared and startable (`profiles: [storage-legacy]`); `docker compose ps storage` rc 0 vs `ps cms` *"no such service"* |
| `jobsimulation.md:183` | the re-pointed command does **not** inherit the old binary's error signature — `app/main.go:212` is a plain `main()`, no cobra, no help block |
| `platform-migration-status.md:74` | false provenance sha — `e45eb61` touched only line 11 (a URL swap); `:19` was last changed at **`84a4b4f`**, eight months before the fold |

### Q5 — a booked anchor that is CORRECT, with the false statement elsewhere (1)

**`ai_architecture.md:225` is true against source and must NOT be edited.** The composited MP4 *is*
written to and read from prod S3 (`chime.go:188`, `:264-266`, `:341-352`, no delete). The false sentence
is at **`corpus/ops/demo/media-substrate-spec.md:33-35`**, where it is load-bearing for a **safety
disposition**.

> **Repairing the booked anchor would break a true sentence and leave the false one standing.** This is
> the sharpest argument in the milestone for adjudicate-before-repair — and note where the real defect
> lives: `corpus/ops/**`, **outside the instrument's 40-file set**, which is exactly what
> [`membership.md`](membership.md) measures.

## Method notes iter-85 must inherit

1. **Four repairs are ONE re-derivation each, not one per anchor.** The `ai-readiness.md` cluster (7
   anchors) is a single rext commit; `app/CLAUDE.md:72` (2 anchors) and `test_injection.py:1051` (2
   anchors) are each one fact in two files.
2. **Five claims stand verbatim in a second site** — `backend.md:218`↔`:25`↔`CLAUDE.md`,
   `services/README.md:37`↔`messenger.md:7`, `storage.md:26`↔`platform-migration-status.md:76`. §5 rule
   19: a claim does not respect a file boundary.
3. **Do not re-anchor Q2.** Seven facts were deleted, not moved.
4. **Do not touch `ai_architecture.md:225`.** Repair `media-substrate-spec.md:33-35` instead.
5. **The reach fence's denominator is now `upheld`** — 40, not 43, and not 152.
