---
milestone: M257x
iter: 01
---

# iter-01 — progress

**Type:** tok (bootstrap). Shape per `corpus/ops/platform-alignment.md` (authored by this iter).

## What was measured

Five parallel probes against **platform origin HEAD** (`1e8e7540`, 2026-07-30T08:26:40Z), plus independent
re-verification by this tok of every load-bearing claim.

### All five open questions, answered

| # | question | answer |
|---|---|---|
| 1 | Where is the migration actually? | A **program**, not three accidents: v2.0 skiller → v5.0 skillpath → v7.0 jobsim → v8.0 cms → **v9.0 storage+messenger, in flight**. Written down at `app/knowledge/plan/roadmap-*-in-app.md`. |
| 2 | Is `jobsimulation` merged, mid-merge, or unchanged? | **`merged-into-app`** in prod (scaled to zero, `app` owns it unconditionally, 136 files, tables re-created in `public`); **`running_but_unfederated`** on a fresh local stack (container still in the default profile; compose fold is platform PR #20, OPEN). Both true — of different environments. |
| 3 | Net-new repos? | **Yes — 20**, of **93** org repos. Only 10 in `repos.yml`; **50 absent from the corpus entirely**. |
| 4 | Does the 3-subgraph count hold? | **No — it is 1** (`backend`), since `915da06` 2026-07-29. Corpus asserts 3 in ~9–16 places. |
| 5 | How much of rext still runs? | **All of it, offline** — 6/6 Go modules build+vet+gofmt clean, **5** Python suites, **2,617 tests, 0 failures**. The prior was wrong at the unit level and right at the live level. |

### The mechanism nobody had found (§2 of the protocol doc)

The premise *"a fresh stack never creates the `jobsimulation` schema"* is **refuted**: rext creates it itself.
`demo-stack/migrate-demo.sh:81-85` `CREATE SCHEMA IF NOT EXISTS` for cms/jobsimulation/skillpath, and `:106`
atlas-applies a **hardcoded 4-tuple** gated at `:108` on `[ -d "$DEV/$r" ]` — it never consults `repos.yml`'s
`migrations:` flag. The tuple is **hand-maintained** (the `:96` comment shows it was edited for skiller, never
for jobsim/cms). **Time bomb:** when M810 removes the legacy repos from the clone set, `[ -d ] || continue`
silently skips them and **13 write targets 42P01 at once**. **Canary already visible:** `skillpath` is in the
tuple but absent from origin `repos.yml`, so its schema is created empty — harmless only because rext writes 0
`skillpath.*` tables.

### Root cause of the recurring class (§3)

**Pinning disables drift detection.** `ensure-clones.sh:393` computes a behind-count only when `ref != "HEAD"`,
but every pinned clone is detached. Measured: **11/11 clones detached, 11/11 `behind: null`, 0 freshness
problems** — while the bring-up logs *"all clones provably fresh-or-pinned."* `DEMO_FRESHNESS_STRICT=1`, the
documented go/no-go, escalates only states a pinned clone cannot enter.

### Claims refuted by re-verification (the discipline earning its keep)

- **KB-2's two-session-tables premise** — `public.sessions` does **not** exist (created, then dropped at
  `20260722104506.sql:79` as the rename completes; verified across all 167 migrations, never recreated). This
  **inverts** the risk: a naive re-point fails *loudly*, the safe mode. See `D-M257x-1`.
- **A2's NUL-byte claim** — "169 files, ~172 contain NULs" is arithmetically impossible and false; measured
  **0**.
- **My own** "27 service docs is off by two" — wrong; 29 `.md` minus README + TEMPLATE **is** 27.
- **Inherited "roadrunner flipped to `migrations: false`"** — it was *already* false; the commit only appended a
  comment.
- **Inherited "app grew undocumented domains"** — all five docs already exist.

## Deliverables landed

- **`corpus/ops/platform-alignment.md`** (new, ~250 lines) — the milestone's `iteration_protocol_ref`. Its
  §4 six detection signals were **executed end-to-end** to prove the procedure runs rather than reads well.
- `corpus/ops/README.md` index row — and the `corpus_index_guard` was **watched going RED** on the unindexed
  doc (rc=1, naming the exact file) before the row was added, then green at 83 docs.
- `decisions.md` — `D-M257x-1/2/3`; KB-2's row corrected in place so no future iter reads it as truth.
- Scratch evidence base: `finding-pin-blindspot.md`, `finding-platform-truth.md`, `finding-jobsim-status.md`,
  `finding-repo-census.md`, `finding-claim-reconciliation.md`, `design-fence.md`,
  `draft-migration-status-map.md`.

## Close — 2026-07-31

**Outcome:** all five open questions answered against origin HEAD; the protocol doc authored and its procedure
executed; the recurring class's root cause identified (pinning disables drift detection) and its actual local
mechanism found (a hand-maintained tuple that bypasses `repos.yml`); five inherited or audited claims refuted by
measurement, one of which inverted a planned guard.
**Type:** tok (bootstrap)
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (bootstrap toks never exit) — (3) re-scope: n (zero
alignment attempts made; trigger needs two invalidated ones) — (4) user-blocker: n (Phase 0b YELLOW, not RED) —
(5) cap-reached: n (0 tiks) — (6) protocol-stop: n — Outcome: continue
**Decisions:** `D-M257x-1` (KB-2 refuted), `D-M257x-2` (fence reads machine-readable fields only),
`D-M257x-3` (`running_but_unfederated` + per-environment axis)
**Side-deliverables:** the `corpus_index_guard` RED-then-GREEN demonstration; independent reproduction of the
full rext build/test matrix.
**Routes carried forward:**
- `FIX-M257x-rext-pin` → iter-02: the pin is inconsistent (`rext.tag` = `cockpit-deeplinks-v1`, **63 commits**
  behind `main`; clone at a different tag) and `ensure-clones.sh:94-101` is **FATAL**, so `/demo-up` aborts on
  this box today. **Blocks gate clause 1.**
- `FIX-M257x-migrate-tuple` → iter-02/03: derive `migrate-demo.sh`'s schema list from `repos.yml` instead of the
  hand-maintained tuple; drop `skillpath`.
- `FIX-M257x-pin-stale` → later tik: add the `pin-stale` state so pinning stops implying blindness.
- `DOC-M257x-guard-severity` → doc tik: two places still call the rext-pin guard "non-fatal"
  (`ensure-clones.sh:68` header, `rosetta_demo.md:17-18`) while the code exits 1.
- `DOC-M257x-subgraph-count` → doc tik: ~9 corpus files assert 3 subgraphs; it is 1.
- `DOC-M257x-ai-labs-repo` → doc tik: `AI-Labs` **is** the `labs-api`; `ai-labs.md:4` says "no separate repo".
- `DOC-M257x-livekit-agents` → doc tik: the corpus names **no** LiveKit agent repo; there are five, two of
  which dispatch nothing.
- `DOC-M257x-repo-states` → doc tik: `skillpath` and `chronos` are both still **active**, not archived.
- `KB-1` (corpus still mandates co-writing the dropped `local_*` mirrors) → doc tik.
**Lessons:**
- **The NUL-byte trap is folklore.** Three false absences here, none caused by NUL bytes; the tree where they
  were asserted had zero. A swallowed stderr turning an engine rejection into "no matches" is far more common.
  Folded into the protocol doc §5.
- **Verify audits too, not just code.** The single most dangerous claim this iter came *from the KB-fidelity
  audit*, and measuring it inverted the plan.
- **A hand-maintained list of the platform's services will silently disagree with the platform.** Derive it or
  fence it.
- **State the environment with every claim.** "Merged" and "still running" were both true, of prod and of a
  fresh stack respectively — which is why the map now carries two states per row.
