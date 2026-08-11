**Type:** tik — under `TOK-08`, on the user's 2026-08-10 closing condition.

_Opened 2026-08-10 09:17 CEST. Pre-registrations PR-1…PR-5 sealed in this iter's first commit
(`061295b`), before any measurement._

## Phase A — the re-survey, and the fetch that moved a number in the first minute

By real `git fetch origin` against six clones, 09:16 CEST:

| repo | HEAD | `origin/main` | behind |
|---|---|---|---|
| `platform` | `0c91421` | `0c91421` | **0** |
| `app` | `ad9f3c498` | `3eaadae68` | **28** |
| `next-web-app` | `8297c684c` | `19423a1fb` | **12** |
| `ant-academy` | `22df69dd8` | `249430c39` | **10** |
| `sentinel` · `studio-desk` | at origin | | **0** |

The brief carried `ant-academy` at **+9** from a cached remote-tracking ref; the fetch made it
**+10**. *A remote-tracking ref is a cache, not a remote* — demonstrated on the very repo set the
advance is about, before any other work (`D-M257x-256-2`).

**Live baseline, named:** `anchor_construct_guard` **883 anchors resolved / 599 unresolvable, reach
59.6 %**, GREEN. Guard family (`--platform`, repo root) **29 GREEN / 0 RED / 5 not-run**, 55 s.

## Phase B — the census, built through the guard's own resolver

`stack-core/advance_impact_census.py` (net-new). For every corpus citation resolving into a clone
behind its own fetched `origin/main`, the cited **line's text** is read at both refs and classified —
`held` · `moved` · `moved-ambiguous` · `changed` · `dead` · `born` · `out-of-range`. Nothing is
interpreted: two strings are equal or they are not.

It does **not** extract citations. It wraps five functions of `anchor_construct_guard`
(`scan_targets`, `bare_anchor_sources`, `block_ref`, `read_target`, `classify`), records the tuple the
guard itself adjudicated, and calls the guard's own `run()` — *a census whose population comes from a
second parser measures the second parser* (iter-251). The positional pairing is **asserted**, not
assumed: `classify` without a preceding `read_target` raises.

**Two partitions, published side by side**, because a block naming a sha is a ref-scoped claim an
advance cannot falsify (§5 rules 41/44) — and a block naming **more than one** sha is one too, even
though the guard falls back to the default ladder there:

| | citations in subject | held | moved | moved-amb | changed | dead | born | IMPACT |
|---|---|---|---|---|---|---|---|---|
| **CONSERVATIVE** (multi-sha block out of subject) | 172 | 139 | 27 | 4 | 2 | 0 | 0 | **33** |
| **PERMISSIVE** (multi-sha block in subject) | 372 | 254 | 105 | 9 | 4 | 0 | 0 | **118** |

Out of subject, identical under both: block-pinned **113**, not-advancing **238**, outside any clone
**182**. Per repo, conservative: `app` **32/139 = 23.0 %**, `ant-academy` **1/13 = 7.7 %**,
`next-web-app` **0/20 = 0.0 %**. Permissive: **36.4 % · 19.4 % · 0.0 %**. Three denominators are
reported and never conflated — **33 citations = 23 corpus sites across 9 documents naming 29 distinct
platform lines** (permissive: **118 / 53 / 16 / 69**).

**An ordering defect in the census, found and fixed inside the iter:** with the pin tests ahead of the
not-advancing test, `not_advancing` read **78** conservative and **183** permissive — the same tree,
the same citations, two totals — because a citation into a clone that had not moved was absorbed by
whichever exclusion the partition reached first. *An exclusion bucket whose size depends on a flag
that is not about it cannot be compared across runs*, which is the entire job of publishing two
partitions.

## Phase C — the repair was built, applied to 27 citations, and REVERTED. That is the iter.

The applier is fail-closed by construction: **one pass per line** (`storage.md:29` carries eight
`main.go` anchors and the map contains both `504→517` and `517→530`; applied sequentially the second
rewrites what the first wrote), an **occurrence-count check** that refuses the whole line when the
line carries more `:NNN` matches than the census recorded, a hyphen guard so **range** citations are
invisible to both count and rewrite, and a hard refusal of `moved-ambiguous` / `changed`. 27
citations across 18 sites applied clean.

Then the six refusals were read one at a time — and the **premise died on the third**.

> **`platform-migration-status.md:121` is entirely `origin/main`-framed.** Its claim
> *"`MESSENGER_ENABLED` (resolved at `app/main.go:286`, read at `:1459` and `:1576`)"* names the two
> `if messengerEnabled {` lines **at `origin/main`**; in the checkout they are at **1445** and
> **1552**. Its import trio *"`:15`, `:63`, `:64`"* likewise names `msgadapters`/`msgsender` at
> `origin/main`; the checkout has them one line earlier. Its subscriber-server trio *"`:1464`, wired
> at `:1485`, sender at `:1487`"* — all three land exactly at `origin/main`.

So the applier's premise — *the corpus holds at the checkout, therefore the advance is what breaks it*
— is **false**, and the seven edits on that line moved seven **correct** citations onto comments. The
mechanism is this milestone's own tooling: `anchor_construct_guard` reads `origin/main` **first** by
default, so any anchor derived or verified through the guard is an `origin/main` number, while any
anchor derived by reading the checkout is a checkout number. Both are in the corpus, adjacent, unmarked:

| site | frame | evidence |
|---|---|---|
| `platform-migration-status.md:121` | **`origin/main`** | 7 anchors, all exact at `3eaadae68` |
| `ai-readiness.md:483` | **`origin/main`** | *"`func` at `:716`"* — `keepInCycleStep1` is at 710 in the checkout, **716** at `origin/main` |
| `observability.md:28` | **`origin/main`** | `WithLoggingTracing` is at 277 in the checkout, **278** at `origin/main` |
| `observability.md:29` | **checkout** | quotes `SentryDSN: os.Getenv("SENTRY_DSN")`, at **273** in the checkout, 274 at `origin/main` — **one row below its `origin/main`-framed sibling** |
| `ai-readiness.md:486` | **checkout** | `keepStartedMembers` at **730** in the checkout, 736 at `origin/main` |

**All 27 edits were reverted** by inverse-applying the same instrument, and `git diff` came back
**empty** for every corpus file — which is also the strongest available evidence that the rewriter
edits exactly what it claims to edit (pinned as `test_45`).

The instrument now carries the finding: **`--apply` REFUSES unless the operator passes
`--adjudicated <file>`** naming the sites they have read and decided are checkout-framed, and the
report prints *"a `moved` verdict says the cited line's TEXT differs between the two refs … it does
NOT say which ref the citation was written against"* before the total. `D-M257x-122-5`'s rule about
bare basenames, applied unchanged to refs.

## Phase D — the advance, taken

All three clones fast-forwarded (`merge --ff-only`, all clean, all ancestors), and **the canonical pin
advanced with them** — which is the half that decides what a `DEMO_ADVANCE_CLONES=pinned` bring-up
builds (⚠️ this sentence said *"what a cold demo builds"* until iter-257 measured the default to be **no
pin at all**; see the amendment on `D-M257x-256-4`):
`rosetta-extensions/demo-stack/clones.pin.json`, `app` → `3eaadae68`, `next-web-app` → `19423a1fb`,
`ant-academy` → `249430c39`. `clone_pin_guard` stays GREEN (4 `repos.yml` repos + 2 sanctioned extras,
every ref reproducible). All six clones now report `behind=0`.

**`clone_drift_guard` went RED on the advance — the fence firing on a real event for the first time in
this milestone**, naming `ant-academy` and `next-web-app` as *"at a sha the corpus never cites"* (24
and 40 citing sites). `app` did **not** fire, because the corpus already cites `3eaadae6` — one more
piece of evidence that the corpus is ahead of the checkout, not behind it.

Cleared with two measured records rather than by re-pinning prose:

* **`corpus/services/ant-academy.md`** — `PUBLIC_CHAPTERS` **499 → 554 (+55)** across the advance (the
  open-source/open-weights and finance back-office families, EN+IT), measured with `git show` at both
  refs. And the inference that does **not** follow, recorded so nobody re-derives it: the demo's
  *"65 course cards"* is a render count over a filtered subset of the **skill-path** objects, and that
  population is **unchanged at 92**. The advance moved chapters, not cards.
* **`corpus/services/next-web-app.md`** — `v2.137.0 → v2.137.3`; the 12 commits are one product theme
  (AI-Readiness ⇄ assignments: upskilling tab reworked around course progress, 3-month default study
  window, auto-assigned plans folded into a per-cycle folder) plus two `/start-sim` hiring
  invitation-token fixes. Named because **`app` moved the same theme in the same window** (`source_ref
  airx:*`, weekly reminder cadence) — the coordinated multi-repo shape §Trap D warns about.

Also measured and recorded for the next iter: the `app` advance ships a **new migration**,
`terraform/migrations/20260804160000_assignment_notification_logs.sql`, and a terraform fix whose own
subject reads *"the backend migration pipeline has been a silent no-op since the atlas 0.7.0 bump."*

**Not claimed, and it is the gate's own clause:** that a demo or a dev stack built at these refs comes
up. No bring-up was run. Clause 1 needs a quiet host.

## Close — 2026-08-10

**Outcome:** The advance is **taken** — three clones and the canonical demo pin now sit at platform
`origin/main`, and the corpus records what moved. The measurement that was supposed to precede it
**refuted its own premise instead**: the corpus's `app` anchors are at **mixed refs**, so a `moved`
verdict is a text delta between two clocks and not a repair instruction. 27 blind repairs were applied
and reverted byte-for-byte inside the iter; the instrument now refuses to apply without per-site
operator adjudication.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue

**Pre-registrations — 2 of 5 held.** (Trend: 1/5 → 3/5 → 5/5 → 4/4 → 2/5 → 3/5 → 2/5 → **2/5**.)

| | claim | prediction | measured | verdict |
|---|---|---|---|---|
| PR-1 | break rate comparable to iter-68's 60 % | REFUTED, `app` **< 20 %** | **23.0 %** conservative, **36.4 %** permissive | **MISS** — both partitions above the threshold I sealed |
| PR-2 | the advance is citation-neutral | REFUTED, ≥ 1 fails | 33 / 118 | **HELD** |
| PR-3 | most commits ⇒ highest broken fraction | REFUTED; ordering will not match | `app` has **both** | **MISS** — though the ordering *below* the top does invert (`next-web-app` +12 → 0.0 %, `ant-academy` +10 → 19.4 %), which does not rescue the prediction |
| PR-4 | ≥ 1 citation `born` in the advance | REFUTED, zero | **0** | **HELD** |
| PR-5 | `CITE_REF=origin/main` turns ≥ 2 guards RED | HELD | **0** — 29 GREEN / 0 RED at both | **MISS** |

**PR-5 is the one worth reading twice.** It is refuted *because `auto` already tries `origin/main`
first* — so the guard family has been grading the corpus at `origin/main` all along, is green there,
and is **blind to all 118** of these. That blindness is the wrong-construct floor
`anchor_construct_guard` discloses on every run, now sized against a **ref delta** for the first time.

**Suite state at close** — Python, `stack-core`, `/usr/bin/python3 -m pytest` (CPython 3.9.6):

* `test_advance_impact_census.py` **24 passed**; the four-module scoped run
  (`frozen_expectation_census` + `advance_impact_census` + `m257x_corpus_file_citations` +
  `anchor_subject_census`) **154 passed**.
* **Whole section, the first run since iter-252: 3 failed / 2,156 passed / 3 skipped in 2,123 s**
  (~35 min, against rule 51's recorded ~1,030 s — this box was running the census and two background
  jobs alongside it, so state the environment with the number). **All three failures were ONE
  inherited defect**, `D-M257x-256-6`: the mutation battery unreadable to pytest from the rext root
  since iter-255. Fixed at the import; post-fix the module collects **6 under both invocation points**
  and `test_suite_census_collection.py` goes **3 failed / 14 passed → 17 passed**.
* **Confirming whole-section re-run after the fix: `2,159 passed / 0 failed / 3 skipped` in
  2,025.42 s.** The section is green, and the +63 against the milestone's standing
  `2,096 passed / 0 failed / 3 skipped` baseline is iters 249–256's own arms (24 of them this iter's).
  **That baseline is hereby re-pinned to 2,159**, and the gap is the reason to re-pin it: a
  four-iter-old total silently became a floor nobody was measuring against.
* Guard family (`--platform`, repo root) **29 GREEN / 0 RED / 5 not-run** — before the advance and
  after it, and identically under `CITE_REF=origin/main`.

**Side-deliverables:**
- Two literal ratchets re-pinned with recorded reasons: `DOCSTRING_LITERAL_CEILING` 234 → **237**,
  `TEST_MODULE_LITERAL_CEILING` 634 → **636**; `COMMENT_LITERAL_CEILING` exact +0 at 220.
- `advance_impact_census.py::clone_refs` graded in `derivation_registry.DECISIONS`
  (`DECLINE:caller-value`) — the completeness fence caught the new module in the same commit that
  added it.
- **The ratchet caught my ratchet prose** (`D-M257x-256-5`): the `TEST_MODULE_LITERAL_CEILING` block
  reads the LAST arrow target as the new ceiling, so a recorded reason quoting a renumbering arrow
  handed off `530` against a constant of `636`. Quote a mapping in words inside a ratchet block.
- One printed-measurement-literal false positive removed at source: `">1 sha"` in the partition line
  read as a count; reworded to `"MORE THAN ONE sha"`, which is better prose anyway.

**Routes carried forward:**
- `ROUTE-M257x-256-mixed-ref-anchors` → **the big one.** 33 (conservative) / 118 (permissive)
  citations whose cited text differs between the two clocks, each needing per-site adjudication
  against the sentence's own claim. Now that the clones ARE `origin/main`, the two clocks have
  collapsed into one and the residual is exactly the **checkout-framed** subset — measurable at a
  single ref for the first time. Handler: `FIX-M257x-256-adjudicate-the-checkout-framed-subset`.
  Two are already known (`observability.md:29`, `ai-readiness.md:486`) and two more are `changed`
  rather than moved (`hiring.md:98`/`:99` — `switch org.IsHiring {` became `if org.IsHiring {` and
  `antRole = enum.RoleCandidate` became `return enum.RoleCandidate` at `:451`), which needs the claim
  re-derived, not the anchor re-pointed.
- `ROUTE-M257x-256-workspace-pin-is-not-the-canonical-pin` → `stack-demo/clones.pin.json` still names
  **11** repos including the five phantoms `clone_pin_guard` removed from the canonical pin at
  iter-222. `ensure-clones.sh` seeds it **copy-if-absent**, so a workspace created before that fix
  keeps the phantoms forever and **the fence does not watch the copy**. Handler:
  `FIX-M257x-256-fence-the-workspace-pin-copy`.
- `ROUTE-M257x-256-the-advance-is-unproven` → clause 1 (cold `demo-down --purge` + `demo-up`, 3
  consecutive green) and the **dev** half of the user's condition are both un-run at the new refs. The
  `app` advance carries a **new migration** and a terraform fix stating the backend migration pipeline
  *"has been a silent no-op since the atlas 0.7.0 bump"* — so `make migrate` behaviour is the first
  thing to watch. Needs a quiet host. Handler: `FIX-M257x-256-prove-the-advance-cold`.
- `ROUTE-M257x-255-a-declaration-narrows-every-mutation-proof` · `ROUTE-M257x-255-the-class-can-regrow` ·
  `ROUTE-M257x-253-the-iter-loop-runs-no-ratchet` · `ROUTE-M257x-254-six-spellings-of-one-root` ·
  `ROUTE-M257x-253-suite-census-is-undocumented-in-rext` ·
  `ROUTE-M257x-251-two-trees-both-called-a-fresh-checkout` ·
  `ROUTE-M257x-250-workspace-tier-is-invisible-without-a-workspace` ·
  `ROUTE-M257x-249-anchor-offset-has-three-populations` → open.
- Still open, untouched: `ROUTE-M257x-246-two-of-four-censused-surfaces-still-have-no-fence` ·
  `ROUTE-M257x-244-two-fences-entered-the-family-unindexed` ·
  `ROUTE-M257x-244-unresolvable-and-wrong-share-one-bucket` · `ROUTE-M257x-h59-range-anchors-are-ungraded` ·
  `ROUTE-M257x-241-wider-citation-surface-is-ungraded` ·
  `ROUTE-M257x-240-prereq-floors-live-in-three-parallel-blocks` ·
  `ROUTE-M257x-238-container-vs-native-is-undrawn` · `ROUTE-M257x-237-hardcoded-vs-settable` ·
  `ROUTE-M257x-236-disclosure-scope-is-document-level` · `ROUTE-M257x-235-fence-scope-is-unread` ·
  `ROUTE-M257x-235-runnable-block-has-two-halves`.

**Lessons:**
1. **A measurement can be correct and its interpretation inverted.** Every number this census printed
   is right; the sentence *"the advance breaks 33 citations"* is not. The delta is real, the
   attribution was a guess, and nothing in the instrument distinguished them until a human read six
   refusals. **Name the direction of a delta as a separate claim from its size.**
2. **The tool a corpus is written WITH decides the ref its numbers are on.** `anchor_construct_guard`
   reads `origin/main` first; that is a good default for grading and it silently made the corpus a
   two-clock document. Any instrument that resolves against a *ladder* rather than a *named ref*
   leaves that footprint — §7 rule 4d's warning, arriving from the authoring side rather than the
   grading side.
3. **The refusals were the finding.** `moved-ambiguous` and `changed` were designed as "a human must
   read these", and three of the six overturned the premise the other 27 were applied under. A bucket
   that refuses is worth more than a bucket that guesses, and it paid for itself on its first run.
4. **Advance the PIN, not just the checkout.** A `git merge --ff-only` in three clones changes what is
   on disk today; `demo-stack/clones.pin.json` changes what a cold bring-up builds on any box. Only
   the second is the advance the gate is about.
