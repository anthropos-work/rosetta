# iter-102 — canonical repairs for the CROSS-SEAT predicates

**Why this file exists.** The repair fan-out is partitioned **by file**, so no two seats write the same
file. But a *predicate* does not respect a file partition — the three widest ones in this union span four,
four and three files across five different seats. If each seat words its own repair, the corpus ends up with
four differently-worded "fixes" of one predicate, which is how iter-98 produced a document holding **both**
readings of the demo-academy auth model nine lines apart.

**So the wording is derived ONCE, here, and each seat applies the canonical form at its own anchors.**
Seats do not re-derive these three; they re-derive everything else.

Every fact below was re-measured at this iter's open — platform `0c91421` (== `git ls-remote origin HEAD`),
`app` `ad9f3c49` (== `origin/main`).

---

## CANON-1 — `sentinel-only-cross-process-edge` (7 anchors across 5 seats)

### The false predicate, in the forms it is published in

> *"the only cross-process RPC edge left in a local stack is `backend → sentinel`"* — **and its
> generalisation** — *"compose sets exactly one service address"* / *"the only cross-process edge"*.

### What is measured

| fact | evidence @ `0c91421` |
|---|---|
| `AUTHORIZATION_ADDRESS=http://sentinel:8087` | `docker-compose.yml:48` |
| **`GOTENBERG_URL=http://gotenberg:3200`** | **`docker-compose.yml:57`** |
| `JUDGE0_BASE_URL=http://52.48.139.23:2358` | `docker-compose.yml:59` |
| `REDIS_ADDR=redis:6379` | `docker-compose.yml:66` |
| `gotenberg` is in the **default `core` profile** | `docker-compose.yml:183` — `profiles: [core, backend, all]` |
| the gotenberg edge is **plain HTTP, not Connect-RPC** | `app/internal/converter/gotenberg.go:31` @ `ad9f3c49` — `http.NewRequestWithContext(ctx, "POST", gotenbergURL+"/forms/libreoffice/convert", &body)` |
| `*_RPC_ADDR` variables in compose | **zero** — this half of the sentence is TRUE and must survive |

### The distinction the repair turns on

**`sentinel` is the only cross-process *Connect-RPC* edge. It is not the only cross-process edge, and
compose does not set exactly one service address.** `gotenberg` is a second cross-process edge on the
default profile, reached over **plain HTTP**. The `*_RPC_ADDR`-is-zero clause is **true** and is not to be
weakened while fixing the generalisation around it.

### The model wording — already in the corpus, and correct

`corpus/architecture/architecture_overview.md:321` already publishes the right form:

> `→ Connect-RPC to sentinel   (the only cross-process RPC edge out of backend on a core stack)`

**Do not edit that line.** It is the model. Every other site is brought to agree with it.

### CANONICAL REPLACEMENT (apply at every anchor, adapted only in grammar)

> the only cross-process **Connect-RPC** edge out of `backend` on a `core` stack is **`backend → sentinel`**
> (`AUTHORIZATION_ADDRESS=http://sentinel:8087`, `docker-compose.yml:48`), and there are **zero
> `*_RPC_ADDR` variables anywhere in compose`. **It is not the only cross-process edge:** `backend` also
> calls **`gotenberg` over plain HTTP** (`GOTENBERG_URL=http://gotenberg:3200`, `docker-compose.yml:57`;
> `gotenberg` is in the default `core` profile at `:183`, consumed at
> `app/internal/converter/gotenberg.go:31`), and Judge0 directly via `JUDGE0_BASE_URL`
> (`docker-compose.yml:59`).

> ### ⚠️ CORRECTION, and it is the most instructive thing in this file
>
> **The wording above originally ended `via \`JUDGE0_BASE_URL\` (\`:59\`)` — a BARE line anchor.** A bare
> `:N` resolves against the **most recently named file**, which in that sentence is
> `app/internal/converter/gotenberg.go`. That file is **53 lines long**, so the citation was
> **out of range** — a wrong-construct citation, the exact class iter-100 built `anchor_construct_guard`
> for and iter-101 still found **4** of.
>
> **`repair_postcondition` caught it in the pre-commit hook, at four sites at once:** `CLAUDE.md:282`,
> `backend.md:48`, `backend.md:282`, `service_taxonomy.md:425`.
>
> **The finding is not the typo. It is that CENTRALISING a wording centralises its DEFECTS.** This file
> exists so that five seats do not word one predicate five different ways — and the price of that is that
> **one error in it propagates to every seat by construction**, instantly, with no independent re-derivation
> to catch it. A canonical wording is a single point of failure in exactly the proportion that it is a
> single point of truth.
>
> **So: always name the file in an anchor inside a canonical wording.** Never a bare `:N` — the seat that
> applies it has no way to know which file the anchor was relative to in the sheet's own prose.
>
> This happened *inside the iter that documents the ~2-defects-per-repair-cycle induction rate*, which is
> the rate holding for the fifth consecutive cycle.

**Anchors (each seat applies at its own):**

| anchor | seat |
|---|---|
| `corpus/services/sentinel.md:85` | 10 |
| `corpus/services/jobsimulation.md:145-146` | 8 |
| `corpus/architecture/platform-migration-status.md:93` (**RPC-edge clause only**) | 2 |
| `corpus/architecture/service_taxonomy.md:405` | 3 |
| `corpus/services/gotenberg.md:50` (already correct — **verify, do not rewrite**) | 9 |
| `corpus/architecture/dependency_map.md:103` (already correct — **verify**) | 7 |
| `corpus/architecture/platform-migration-status.md:105` (already correct — **verify**) | 2 |
| **`CLAUDE.md:280`** — the repo-root instructions, **verbatim the same claim** | **orchestrator** |

> ⚠️ **`platform-migration-status.md:93` carries TWO clauses in one table row.** The **RPC-edge** clause is
> upheld and is repaired here. The **prod-terraform** clause in the same row was the single **wrong-tree
> REJECTION** of iter-101 — it is **NOT a defect** and must be left exactly as it stands. Do not merge the
> two while editing one.

---

## CANON-2 — `prod-terraform-8081` (≥4 anchors across 3 seats)

### The false predicate

The claim that the **production** terraform sets `backend`'s internal port to **8081** (or the equivalent
sentence pinning a prod-terraform value), published at four corpus anchors.

### The binding constraint — and it is TRAP A

The deciding file lives in the **`infrastructure`** repo, which **has never been in any clone set**. It is
**not measurable from here**, and this milestone already has a guard that enforces exactly that
(`unreadable_repo_claim_guard` — *"all 7 `module.*_euwest1` mentions are marked unmeasurable"*).

**So this is a RESTATE-OR-DROP, never a re-anchor.** Do not repoint the citation at a different file to
make it resolve; a correctly-cited false statement is worse than a stale one.

### ⚠️ CORRECTED — the first version of this canon was WEAKER THAN THE EVIDENCE

The wording below originally read only *"not measurable from this repo."* That is **half the verdict**, and
the weaker half. The adjudicators did not merely fail to confirm the claim — **they measured it and found
zero**:

| measurement | result |
|---|---|
| `git grep` at each clone's own ref over its **tracked** `.tf` — **44 files across 13 clone dirs** | **0 files** |
| raw filesystem grep over **all 59 `.tf` files on disk** under `stack-demo/` | **0** |
| where the `:8081` literal actually occurs | **exactly one file, and it is a markdown KB page** — `app/knowledge/service-dependencies.md`, at **`:52` @ `ad9f3c49`** (it is `:46` at the older pin `b948604f`) |

> **⚠️ Two numeric corrections to this sheet, both found by seat 8 and both re-verified by the
> orchestrator.** The first draft said *"all 12 clones"* and *"all 59 `.tf` files"* as though those were
> one basis — they are two: **44 tracked `.tf` across 13 clone dirs** (the git-per-ref measurement) versus
> **59 on disk** (the filesystem measurement). Both return 0, which is what makes the claim safe, but
> quoting them as one number is the *"mixed bases"* defect this milestone keeps finding. The second draft
> said the KB page's line was `:46` **without naming a ref** — true at `b948604f`, and **`:52`** at
> `ad9f3c49`. A line number with no ref is the currency-pin defect (CANON-3) committed inside the sheet
> that fixes it.

And there is a **self-contradiction inside `cms.md` itself**: `:18` states *"the deletion itself lands in
`infrastructure`, **which has never been in any clone set we have**."* As the adjudicator put it —
**a doc that states it cannot see the production terraform cannot report what it "still names."**

The sentence is also **present tense (`still names`) and names no ref**, so it claims *currency* rather
than recording a dated reading.

**So the repair DROPS the assertion — it does not soften it.** Both halves must appear: the measured zero
*and* the unmeasurability of the one tree that could settle it. Stating only the second reads as "we didn't
check," when in fact we did.

### CANONICAL REPLACEMENT

> **No `.tf` file in any clone names `http://backend.internal.anthropos:8081`** — measured at each clone's
> own ref across all 12 clones and over all 59 `.tf` files in the workspace, **0 hits** in both. The one
> occurrence of the literal anywhere is a **markdown KB page**, `app/knowledge/service-dependencies.md:46`,
> which is not terraform. **And the production declaration is not measurable from this repo at all:** it
> lives in `infrastructure`, which has never been in any clone set — so no *"still names"* claim can be
> made here in either direction. See
> [`platform-migration-status.md`](../architecture/platform-migration-status.md) for the fenced
> unmeasurable-claims convention.

**This is TRAP A, and it is the textbook shape of it:** the underlying fact was **deleted**, not moved.
Do **not** repoint the citation at `service-dependencies.md:46` to make it resolve — that page is a
markdown doc, not the production terraform the sentence claims to be reporting, and citing it would
manufacture a **correctly-cited false statement**, which is worse than a stale one.

**Anchors:** `cms.md:196` (booked) · `cms.md:55` (twin) · `jobsimulation.md:49-50` (twin) ·
`backend.md:241` (booked MINOR at iter-101, out of the union — **repair it anyway**, it is the same
sentence). Seats 8, 8, 8, 1 respectively.

---

## CANON-3 — the CURRENCY PIN: `2035f9a` is no longer `origin/main`

### The measurement

| | |
|---|---|
| `app` `origin/main` at iter-99/iter-101 | `2035f9a4` |
| `app` `origin/main` **now** | **`ad9f3c49`** |
| distance | **5 commits, 5 files** (`.claude/skills/publish/SKILL.md`, `CLAUDE.md`, `knowledge/deployment.md`, `terraform/main.tf`, `terraform/variables.tf`) |
| sites in `corpus/` + `CLAUDE.md` **labelling `2035f9a` as `origin/main`** | **17** (15 in `corpus/`, 2 in `CLAUDE.md`) |

### The rule — and it is narrow, so read it before editing

**A pin is a pin. `2035f9a` still resolves and still means what it meant, so a citation that PINS it is
CORRECT and must not be touched.** What expired is the **LABEL**: the corpus calls `2035f9a`
**`origin/main`**, and origin/main has moved. iter-95's own seat report put it exactly —
*"still holds at 2035f9a; only the 'origin/main' LABEL expired."*

### CANONICAL REPAIR — two moves, and pick per site

1. **If the claim is about a pinned historical state** (most of them): **keep the sha, drop the moving
   label.** `@ origin/main 2035f9a4` → `` @ `app` `2035f9a4` (origin/main on 2026-08-06; now `ad9f3c49`) ``.
2. **If the claim is about what is CURRENT** ("`app` is at v1.369.0", "the newest ref"): **re-derive at
   `ad9f3c49`** and re-state with the new ref and a date. A version is *a reading at a ref, never a standing
   "current"* — this milestone's own rule, from iter-98 P6.

**Never** leave a bare moving label. **Never** silently swap `2035f9a` → `ad9f3c49` in a sentence whose
line numbers were measured at `2035f9a` — that manufactures a wrong-construct citation, which is the class
iter-100 spent a whole iter on and iter-101 still found 4 of.

**Anchors, per seat:**

| file | anchors | seat |
|---|---|---|
| `CLAUDE.md` | `:223`, `:280` | **orchestrator** |
| `corpus/architecture/dependency_map.md` | `:59` | 7 |
| `corpus/architecture/platform-migration-status.md` | `:87`, `:89`, `:90`, `:92` | 2 |
| `corpus/services/backend.md` | `:70`, `:138`, `:254`, `:299` | 1 |
| `corpus/services/academy-backend.md` | `:20` | 5 |
| `corpus/services/ai-labs.md` | `:18` | 6 |
| `corpus/services/coursebuilder.md` | `:132` | 6 |
| `corpus/services/messenger.md` | `:43` | 9 |
| `corpus/services/skillpath.md` | `:35` | 9 |
| `corpus/services/storage.md` | `:29` | 9 |

### The four OTHER clones that moved, and why they are NOT in this repair

`next-web-app` (41 commits), `ant-academy` (5), `sentinel` (2) and `studio-desk` (2) also advanced. Their
corpus citations are **pinned by sha, not labelled `origin/main`** (`bb3313bc` ×1, `88bc5592` ×4,
`9c3843cd` ×0, `14a5442a` ×0), so **a pin is a pin and nothing needs repairing**. Recorded because the
absence is the finding: only the *labelled* citations rot, which is precisely why the label is the defect.
