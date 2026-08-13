**Type:** tik — under `TOK-05`, step 3 of its next-tik direction (*the map's `mid-fold` state + the
storage row*), routed forward from iters 59 → 62 → 63.

# iter-64 — the map's eighth state

## Phase A — re-derive the split from artifacts

Adjudicated at platform `0dab54d` / `app` `b948604` v1.366.0, against the artifacts and **not** against
`storage.md` (§5). The split is intact and one side was under-counted:

| side | measurement |
|---|---|
| config | `STORAGE_RPC_ADDR`: **0** occurrences across `docker-compose.yml`, `common.yml`, `.env_example` |
| compose | `storage` moved to `profiles: [storage-legacy]` (`:135`), rationale in-comment at `:131-134` |
| `repos.yml` | still an entry (`:18-20`) — still cloned |
| consumer | `app` reads it at `main.go:446`, `:524`, `:992` **and in three `cmd/` tools** |

**Three, not two** (`D-M257x-64-3`). `cmd/academyImport/main.go:231` and
`cmd/academy-asset-upload/main.go:129` hard-require it; **`cmd/import/main.go:50` builds a storage
client against the empty string without complaint** — a reader nobody had counted, with the same
silent-deferral shape as `main.go`.

## Phase B — land the token

`mid-fold` added to the map's §1 vocabulary; `storage`'s fresh-local-stack cell now reads
**`mid-fold` (startable via `storage-legacy`)** with both sides cited in the row. The guard's
`ALLOWED_STATES` widened 7 → 8, its assertion-C docstring corrected, and the protocol's
*"the seven-token vocabulary above has no token for mid-fold"* sentence retired — it described a gap
this iter closes. Both *"one of the seven"* claims (protocol §8's fence table, the map's own §4
assertion list) corrected to eight.

The map's freshness pin advanced `ef32d4c` → **`0dab54d`** in the same pass; it had been naming the
merge commit while every row beneath it cited the newer ref.

**`D-M257x-64-1` — a vocabulary gap is a claim the map cannot make.** This item had been routed
forward five times, and the reason is structural rather than lazy: iter-59's measurement had *nowhere
legal to go*. Inventing a token in a fenced field turns the guard RED, so the measurement landed in an
unfenced service doc while the fenced artifact kept the old answer — and kept passing. **The fence had
eight things to say and seven words.** Same failure shape as `D-M257x-63-2`, one layer up: **the
instrument reports agreement it never tested.**

## Phase C — watched RED

The guard fired on the very first attempt, correctly and for the wrong reason: the new cell is
**bolded**, and `_state_head` took `**mid-fold**` literally. Fixed by stripping markdown emphasis
(`D-M257x-64-2`) — derived from the format, not an exception for one row — rather than by un-bolding
the cell, which would have gone green while leaving the next author the same trap.

| mutant | tests RED | live map RED |
|---|---|---|
| remove `mid-fold` from the vocabulary | 4 | yes |
| stop stripping markdown emphasis | 3 | yes |

Both mutants take the **live map** RED as well as the fixtures — the assertion is exercised against the
real artifact, not only a synthetic one.

## Phase D — gates

| gate | result |
|---|---|
| `platform_alignment_guard` (the map's own fence) | **OK** — both directions |
| `platform_predicate_guard` | OK |
| `anchor_construct_guard` · `markdown_structure_guard` · `corpus_index_guard` | OK |
| `tests/test_platform_alignment_guard.py` | **35 tests** (was 30), all pass |
| `test_test_collection_fence` | OK — new class appended **before** the `__main__` guard this time |
| `stack-core` suite | **669 tests, 1F** — the perishable iter-48 fixture, the single expected failure |
| §5 rule 34 re-point | 2 intra-corpus citations moved by this iter's edits; both resolved by hand after checking the target, because the re-point script is not idempotent |

## Close — 2026-08-04

**Outcome:** the map gained its **eighth** state token, `mid-fold`, and `storage` is its instance with
both halves of the fold cited in the row — closing an item routed forward five times whose real
blocker was that **the measurement had nowhere legal to go**. Re-deriving the split found iter-59's
consumer count short by one: **three** `cmd/` readers, not two, the third building a client against the
empty string. The guard was watched RED on two inversion mutants, both of which take the live map RED
as well as the fixtures.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5, unchanged; clause 5 is still graded only by a reading that returns zero.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** `D-M257x-64-1` (a vocabulary gap is a claim the map cannot make), `D-M257x-64-2`
(markdown emphasis is presentation, not part of the token), `D-M257x-64-3` (the split re-derived;
three `cmd/` readers, not two).
**Side-deliverables:** the map's freshness pin advanced `ef32d4c` → `0dab54d`; `service_taxonomy.md`'s
*"CMS, Jobsimulation and Roadrunner are NOT out of local orchestration"* ⚠️ block retired — the phase
it describes closed at `d11a403`.
**Routes carried forward:**
- `CHECK-M257x-iter64-pms-87-subject` → the Directus-retraction passage cites
  `platform-migration-status.md:87` as its "fenced source of truth"; that row is
  `anthropos-studio-room`. The anchor resolves and does not name the claim — sibling of
  `CHECK-M257x-iter60-g6-citation-subject`, and predates this iter.
- `FIX-M257x-iter63-app-citation-residual` (the 68 non-mainline `app` citations, routed WHOLE) ·
  `CHECK-M257x-iter63-quoting-a-retired-token` · `CHECK-M257x-iter60-g6-citation-subject` ·
  `FIX-M257x-iter53-union-set` (**PENDING USER DECISION**) · `FIX-M257x-iter56-assignment-flake`
  (**NOT DECIDED** — needs a failure *rate*) · `CHECK-M257x-iter38-ai-act-classification` (needs an
  owner outside this milestone) · `CHECK-M257x-iter57-anchor-guard-bare-class` ·
  `FENCE-M257x-iter54-refs-block` · `FIX-M257x-iter57-within-block-drift` ·
  `CHECK-M257x-iter58-derive-preregistrations` · `CHECK-M257x-iter52-second-ai-manager` ·
  `-cold-daemon-registry` · `-grep-vs-failclosed` · `-empty-stdout-class` · `-baseline-refs` ·
  RF-2/3/7–13 · root `CLAUDE.md`.

**Lessons:**

1. **An item that keeps being routed forward may be blocked by its own destination.** Five deferrals
   were not procrastination — the fenced field had no legal value. Before re-routing a measurement,
   check whether the artifact it belongs in can *express* it.
2. **Widen the vocabulary in the iter that takes the measurement.** Split them and the measurement
   lands somewhere unfenced while the fenced artifact keeps passing on the old answer.
3. **A guard finding can be about formatting and read like a finding about truth.** `'**mid-fold**' is
   not one of [...]` names the right cell for the wrong reason. Fix the parser, not the prose — the
   next author will bold it too.
4. **Re-derive a routed measurement, do not carry it.** iter-59's two-reader count became three on
   re-derivation, and the third is the interesting one: it fails silently.
