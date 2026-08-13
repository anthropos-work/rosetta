# iter-47 — clause-5 SEVENTH pass: the blocker ledger

**7 unique in-scope blockers** (8 raw findings; `G1 ≡ B1`, counted once). Seven auditors, **40 files /
9,243 lines**, every in-scope file read top-to-bottom with a `wc -l` positive control per file — **all 40
confirmed line-for-line**. Every blocker re-derived by this iteration before acceptance (`adjudication.md`)
— **7 of 7 held**.

---

## THE HEADLINE: the pre-existing residual measured ZERO

| pass | iter | auditors | corpus read | blockers | pre-existing | induced |
|---|---|---|---|---|---|---|
| 1 | 21/33 | grep-scoped → 5 | pre-repair | 25 | — | — |
| 2 | 34 | ~5 | post-33 repair | 13 | — | — |
| 3 | 38 | 6 | post-34 repair | 11 → 17 | — | — |
| 4 | 39 | **7** | post-38 repair | **37** | — | — |
| 5 | 41 | **7 — fixed instrument** | post-39 repair | **18** | **9** | **9** |
| **6** | **47** | **7 — SAME fixed instrument** | post-46 repair | **7** | **0** | **7** |

**Six full-read auditors covering all 40 files found ZERO blockers in text iter-46 did not touch.** Every
one of the 7 is in text iter-46 wrote or rewrote. The residual is now **100% self-inflicted**.

This is the first time in the series that the two terms have been separable and one of them has gone to
zero. iter-41's *"the fixed point of this process is not zero"* is now refutable in a specific way: **the
corpus term reached zero; the repair term did not.**

### Pre-registered predictions — graded

| # | prediction (written before any auditor reported) | outcome |
|---|---|---|
| headline | **fewer than 9 blockers** (TOK-02's own, carried unmodified) | **CONFIRMED — 7.** The first confirmed prediction in the series; the four passes before iter-41 each refuted their own |
| 1 | **fewer than 4 of the self-contradiction class** | **CONFIRMED — 3** (G3, G4, G5, all "repaired at one site, left standing at another") |
| 2 | **at least one blocker in text iter-46 wrote to *explain* a correction** | **CONFIRMED — 4** (B1/G1, B2, C1, G2) |
| 3 | **the residual is NOT concentrated in the 17 files iter-46 edited** | **REFUTED — 7 of 7 are in them.** The repaired-text density model holds, and holds absolutely |
| consequent | **clause 5 does not close** | **CONFIRMED** |

---

## The 7

### Newly-written text that over-corrects — 4

| # | site | the false claim | what is true | seat |
|---|---|---|---|---|
| 1 | `ai_architecture.md:51-54` | *"`ai_vendor` unset or unrecognised falls through the `default:` arm… `AIVendor` is a **nullable** pointer on the sequence, so *unset* — not merely *mistyped* — is the ordinary way to reach it"* | **Both conjuncts fail.** `simulation.Sequence.AIVendor` is a **value**; the nullable pointer is on the Directus DTO (`jobsimulation.go:905`), one layer up. `:1302-1305` normalizes nil → `simulation.Openai` **before** the sole construction site `:1307`, so unset lands on `case simulation.Openai:` (`simulator/ai/ai.go:58-59`). The `default:` arm is reachable **only** by an unrecognised value — `Azureglobal` is a real 5th enum member (`:967-973`) with no case. **iter-46 inverted the very distinction it was writing to fix**, and contradicted `external_services.md:532`, the site its own next line calls *"the full per-line derivation"*. The *outcome* claim survives; the mechanism does not | **B + G** |
| 2 | `ai_architecture.md:42-43` | *"it flips Course Builder and **Studio-Room** off Bedrock onto Anthropic's first-party API"* | **Studio-Room was never on Bedrock.** `grep -rin 'bedrock\|boto3' app/studio/` → **0 hits**; three providers only (`ai.py:334/:490/:627`). The key is a **credential, not the selector** — the selector is the ini's `TARGET SERVICE`. iter-46 rewrote this sentence and **preserved the false conjunct** it did not notice | G |
| 3 | `architecture_overview.md:246-247`, compounded `:251-252` | *"the **two US paths** are a feature flag and a 429 retry target"* … *"one feature flag routes traffic to the US"* | **There is a third, unconditional path.** `jobsimulation.go:1302` defaults an unset vendor to `simulation.Openai` → `getClient`'s `case Openai:` → `ai.go:80 openai.NewOpenAI(openaiKey)`, direct `api.openai.com`, **no region override, first attempt, no 429**. Residency-relevant. `external_services.md:532` enumerates **two ways in** and closes *"**Path (a) gets there on the first attempt**"* — iter-46 wrote a weaker enumeration into `architecture_overview.md` than the one it wrote into `ai_architecture.md` | C |
| 4 | `security_compliance.md:197-198` | *"`external_services.md:489` carries the provider row"* | `sed -n '489p'` → `// Types in app/__generated__/`, a TypeScript codegen comment. The **Anthropic Direct** row is at `:533`. **The anchor was transcribed from iter-41's ledger without re-derivation** — the precise failure `D-M257x-46-1` claimed to have eliminated | B |

### Repaired at one site, left standing at another — 3

| # | site | what survived | seat |
|---|---|---|---|
| 5 | `external_services.md:139` | *"**all of that is false**; that service **has never existed** in the platform compose"* — verbatim the over-correction iter-46 repaired at `service_taxonomy.md:296-303`, **left standing at the twin**. Same git evidence refutes it (`a2a3ee6^` `:383`/`:384`/`:386`/`:409`). And verifying against HEAD cannot establish *"never existed"* | G |
| 6 | `ai_architecture.md:84` | the ordered arrow chain *"EU Azure default → US Azure via the flag → direct-OpenAI on 429"*, which iter-46 rewrote in **three** other files — surviving **68 lines below this same file's own fence** at `:15-17` (*"no such ladder exists in the code"*) | G |
| 7 | `ops/demo/coverage-protocol.md:614-616` | *"the default AI-readiness dashboard GET **always** takes the live-recompute branch"*, in the present tense, **13 lines above** the fence iter-46 added at `:627-632`. Refuted by `readiness.go:309-312` | G |

---

## Why the fences did not catch these

**Not a fence failure — a scope statement, and it is the actionable finding.** All four fences report
**0 sites** on this tree, correctly:

| fence | why it is silent here |
|---|---|
| `claim_twin_guard` | matches **adjudicated refuted forms** from prior ledgers. Six of the seven are **newly written prose** with no ledger entry; #5's old form was never *in* a ledger as a quotable refuted form |
| `markdown_structure_guard` | none of the 7 is structural damage |
| `anchor_construct_guard` | #4's `external_services.md:489` **resolves and carries content** (a code comment) — it is the *wrong* construct, which is precisely the class `D-M257x-45-3` documented as out of reach |
| `derived_value_guard` | none of the 7 is a scalar |

**The gap is nameable: nothing checks a claim written *for the first time*.** The fence family was built
against *"has an already-refuted claim come back?"*. Blockers 1–4 are a different question — *"is this new
sentence true?"* — and #5–#7 are the old question asked of a form no ledger records.

**#5, #6 and #7 are mechanically findable today**: each is a grep for a string that already exists in the
tree next to its own repaired twin. A leak-check over the repair's own diff — *for every claim this commit
changed, grep the tree for the old form* — would have caught all three, and is exactly what auditor G did
by hand.

---

## Minors — 64

A: 10 · B: 11 · C: 9 · D: 13 · E: 7 · F: 9 · G: 5. Per the gate's own wording (*"YELLOW with 0 blockers"*)
minors do not block. Recorded in `raw/*.md`, routed as `DOC-M257x-iter47-minors`.

Two are worth promoting out of the pile because they are cheap and load-bearing:

- **`service_taxonomy.md:150-153` (G-minor-1)** — iter-46's new Technology cell spans **four physical
  lines inside a GFM table**; the row **will not render as one row**. A correction that breaks the table it
  corrects.
- **`hiring.md:189-196` (E-minor-2)** — the "minimal write-set" omits `token` (NOT NULL + UNIQUE, no
  default). A raw-SQL seeder built from that list fails its INSERT.

---

## Clause 5

**NOT MET — 7 blockers.** A clause is met by a reading that returns zero, never by a repair that clears its
own findings. This iteration repaired nothing, by design.

**Gate: 4 of 5.**
