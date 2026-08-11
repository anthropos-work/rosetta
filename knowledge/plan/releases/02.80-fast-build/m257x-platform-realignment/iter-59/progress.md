**Type:** tok (triggered — by a direct user directive, not by the 3-no-prog streak)

# iter-59 — progress

## Phase A — Step 0 re-survey (mandatory for a tok)

Every denominator TOK-05 cites was re-derived from platform artifacts at this open. **INPUT 3 of the run
briefing is that three inherited numbers failed in one iteration, one of them the orchestrator's own** — so
nothing here is inherited. The full table is in `overview.md`; the summary is **11 of 12 premises confirmed,
1 corrected**, and the correction makes the headline defect worse rather than smaller.

**Confirmed:** platform `0dab54d` level with `origin/main` (P3 satisfied, no re-point owed) · `PROFILE ?=
core` · **0** services declare a `graphql` profile · **6** `repos.yml` entries · **1** `migrations: true`
repo · **13** published ports · **4** `*_RPC_ADDR`, all `http://backend:8083` · **17 files / 30
occurrences** asserting a `graphql` profile · **23** `main.go:N` citations · `stack-core` **1F/610** (`Ran
610 tests`, the single F being the perishable iter-48 answer-key fixture — not spent).

**Corrected — compose services are 10, not 8, and `PROFILE=graphql` starts 3, not 0.** `docker-compose.yml`
opens with `include: [common.yml]`, which contributes `postgresql` and `redis`; neither declares a
`profiles:` key, and compose always selects a profile-less service. So a documented-but-dead profile brings
up the infrastructure and omits the application:

```
DERIVED legal profile set (8): all backend core customerio-sync frontend messenger storage-legacy studio-desk
PROFILE=core     -> 5 : postgresql redis sentinel backend gotenberg
PROFILE=graphql  -> 3 : postgresql redis sentinel
PROFILE=cms      -> 3 : postgresql redis sentinel
PROFILE=storage  -> 3 : postgresql redis sentinel
```

**"Starts nothing" would be an honest failure.** Postgres answers, Redis answers, sentinel is up, `docker
ps` is non-empty — and the thing that is missing is the product. This is the second orchestrator-supplied
fact re-derived and corrected in two iterations (iter-58 corrected *"demo-1 GONE — 0 containers"*, a
dead-daemon false absence over 11 live containers). Not a criticism of the hand-off; the milestone's own
thesis applied to its own inputs.

Two further measurements taken for `D-M257x-59-4`, neither previously recorded anywhere in the corpus:
compose **no longer sets `STORAGE_RPC_ADDR` at all** (absent from `docker-compose.yml` and `.env_example`),
while app `v1.366.0` still reads it at `main.go:446/:524/:992` and **hard-requires** it in
`cmd/academyImport/main.go:235` and `cmd/academy-asset-upload/main.go:133`.

## Phase B — author TOK-05

`TOK-05: stop repairing claims; fence the predicates under them` appended to the milestone-root
`decisions.md`. Five directed decisions recorded in full in this iter's `decisions.md` as
`D-M257x-59-1` … `D-M257x-59-5`.

The revision in one line: **the unit of repair changes from the claim to the predicate.** The 81 drift
sites, the 30 `graphql`-profile occurrences and the 21 moved citations are **119 sites over 3 predicates**,
and each predicate's legal set is derivable from a platform artifact we already parse. A reading can name
instances; it cannot name the predicate they are instances of — which is why ten readings at 43–48% recall
never converged.

Everything TOK-04 built is kept whole (P1–P4), as is the audit instrument, the union-of-two discipline and
the pre-commit double reads. Clause 5 is **not** re-cut: still met only by a reading that returns zero.

## Close — 2026-08-04

**Outcome:** TOK-05 authored — the repair unit moves from the claim to the predicate; three residual classes
(81 drift sites · 30 `graphql`-profile occurrences · 21 moved citations) re-scoped as **3 derivable
predicates**, with a 6-assertion sibling fence, a citation-safety half for §7 rule 4, a `mid-fold` map state
for half-landed folds, and a fence-first ordering. One inherited denominator corrected by measurement
(`PROFILE=graphql` starts **3** containers, not 0).
**Type:** tok
**Status:** closed-no-lift
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: **y** — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-2
**Decisions:** `D-M257x-59-1` (predicate-scoping; union-set left as pending user decision) · `D-M257x-59-2`
(the sibling guard, 6 derived assertions, both directions) · `D-M257x-59-3` (§7 rule 4's citation-safety
half) · `D-M257x-59-4` (the `mid-fold` state, two-sided or not at all) · `D-M257x-59-5` (ordering)
**Side-deliverables:** none — a tok authors strategy; no code changed and no corpus claim was repaired in
this iter.
**Routes carried forward:**
- `FIX-M257x-iter59-profile-class` → **iter-60**, handler = the new sibling guard's G1. 17 files / 30
  occurrences of a `graphql` profile, plus `cms` and `storage`.
- `FIX-M257x-iter58-mainline-shift` (**21 of 22**) → **iter-61**, under `D-M257x-59-3`'s new rule half.
- `DOC-M257x-iter59-storage-mid-fold` → **iter-62**, handler = the map's new `mid-fold` state + G6.
- `FIX-M257x-iter53-union-set` (46 vs 35) → **PENDING USER DECISION**, unresolved by design; `D-M257x-59-1`
  states how predicate-scoping subsumes it without answering it.
- `FIX-M257x-iter56-assignment-flake` → still **NOT DECIDED**; needs a failure *rate*, not another pass.
- `CHECK-M257x-iter38-ai-act-classification` → needs an owner **outside** this milestone; not settled here.
- Unchanged and still open: `-cold-daemon-registry` · `-grep-vs-failclosed` · `-empty-stdout-class` ·
  `-baseline-refs` · `CHECK-M257x-iter58-derive-preregistrations` · `FIX-M257x-iter57-within-block-drift` ·
  `CHECK-M257x-iter57-anchor-guard-bare-class` · `FENCE-M257x-iter54-refs-block` ·
  `CHECK-M257x-iter52-second-ai-manager` · RF-2/3/7–13 · root `CLAUDE.md`.

**Lessons:**

1. **A document that agrees with yours is not a second witness.** PR #14 returned 92 absorbed / 30
   superseded / 5 standing / **0 refuted / ZERO new information** — and every one of the seven live defects
   sat where the two documents *agreed*. Diffing two documents is structurally incapable of finding what
   they share. **Adjudicate against platform artifacts, never against another doc.** Promoted into TOK-05.
2. **A reading names instances; only a derivation can name a predicate.** This is the reason the ten-reading
   series had a fixed point. It generalizes beyond this milestone and belongs in §5.
3. **Grade a documented command on "does it still select something", not "does it still parse".** The
   silent no-op passes every syntax check there is, and the measured version is worse than the reported one:
   `PROFILE=graphql` starts the infrastructure and omits the application, so the failure *presents as a
   partially-working stack*. Belongs in §5 alongside rule 1.
4. **The platform's config files are its documentation of record; its narrative docs are not.** Config is
   edited in the same commit as the change and carries rationale inline (`repos.yml`'s header;
   `docker-compose.yml:130-133`'s storage note). Its narrative docs lag and are partly unmeasured — app
   `v1.366.0`'s own `knowledge/*.md` asserts "60K+ skills" with no measurement, and **the repo contains no
   job-role count anywhere**, so "18K roles" has no upstream provenance in the repo we would be deferring
   to. Belongs in §6.
5. **Re-derive the hand-off's numbers, including the orchestrator's.** Two iterations, two corrected
   orchestrator-supplied facts. §5 rule 1's oldest lesson does not stop applying because the sender is
   trusted.

**Protocol-doc updates owed (per the skill's protocol-evolution rule):** lessons 2, 3 and 5 are §5 material
and lesson 4 is §6 material. **Deliberately NOT written in this commit** — a tok's deliverable is the
strategy, and `platform-alignment.md` is itself inside clause 5's fidelity scope, so editing it in a tok
would add unfenced prose to the very corpus TOK-05 is about to fence. They are routed to **iter-60**, which
builds the fence that covers them.
