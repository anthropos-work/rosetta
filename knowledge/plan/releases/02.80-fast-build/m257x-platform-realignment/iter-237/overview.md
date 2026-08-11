---
iter: 237
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-10
---

# iter-237 — are the environment variables the corpus calls critical actually read by anything?

## Step 0 — Re-survey

iters 235–236 closed the two halves of a runnable **command**: the `make` target (aligned) and the `cd`
directory (four repairs, all one path). A runnable stack has a third input and it is the one that fails
**silently**: the **environment**. A wrong `cd` errors instantly; a wrong env-var name gets you a stack
that boots, answers `docker ps`, and does the wrong thing — the exact failure shape `CLAUDE.md`'s own
compose-profile banner already apologises for.

`CLAUDE.md` carries a **"Critical environment variables"** list of five names. That list is a
**hand-maintained tuple** — precisely what `platform-alignment.md` § 2 identifies as the latent-drift
construct this whole milestone exists for — and **nothing derives or fences it.**
`platform_predicate_guard` grades compose *profiles*; `demo_knob_guard` grades `DEMO_*` knobs against the
demo parsers. **Neither reads platform env-var names**, so the platform could rename `OPENAI_KEY` tomorrow
and every instrument would stay green.

Env-var names are **exact tokens** — `OPENAI_KEY` is not a paraphrase of `OPENAI_API_KEY`, it is a
different string that resolves to empty. So iter-234's refusal does not reach this class: containment is
the *correct* test here, because the subject really is a literal.

**Active strategy reference:** `TOK-08` — census a mechanical class exhaustively.

## Hypothesis

The v8/v9 folds deleted three compose services and their whole env blocks (`838d907` alone removed four
`*_RPC_ADDR` variables). A corpus that grew across that program should name variables the platform no
longer sets or reads — and the highest-risk site is the five-name list in `CLAUDE.md`, because it is short,
old, load-bearing, and unfenced.

## Predictions — SEALED BEFORE MEASUREMENT

| id | prediction |
|----|-----------|
| `P-237-1` | ≥ 150 distinct `UPPER_SNAKE` env-var names are used across `corpus/**` + `CLAUDE.md` |
| `P-237-2` | `platform/.env_example` declares ≥ 40 variables |
| `P-237-3` | ≥ 1 of `CLAUDE.md`'s five "Critical environment variables" is **not** declared in `.env_example` **and** not set in `docker-compose.yml` |
| `P-237-4` | ≥ 5 corpus-named variables appear **nowhere** across `platform` + `app` + `sentinel` at HEAD |
| `P-237-5` | ≥ 1 such orphan is presented as **currently required**, not as historical |

## Expected lift

No `N`/`P` reading. Deliverable: the corpus-named variable population, the platform's declared+read set
(from `.env_example`, `docker-compose.yml`, and Go/TS source), the orphan set with its denominator, each
orphan classified (renamed / deleted-with-its-service / never-existed / lives-in-another-repo /
tooling-only), and repair of any orphan presented as currently required.

## Phase plan

1. Derive the platform's declared set from `.env_example` + `docker-compose.yml` at `platform` HEAD.
2. Derive the *read* set by grepping `app` + `sentinel` + `platform` sources for each name — declared and
   read are different questions and both matter.
3. Enumerate corpus-named variables, keeping the enclosing document and whether the mention is fenced.
4. Grade `CLAUDE.md`'s five-name critical list **individually and by name** — it is the load-bearing site.
5. Classify orphans; repair only those presented as current.

## Escalation conditions

- 0 orphans → prove the instrument on a name known to be dead (`SKILLER_RPC_ADDR`, deleted at `838d907`).
- A repair would need a platform edit → route, never edit the platform.

## Acceptable close-no-lift outcomes

A measured 0 currently-required orphans, with the dead-name control firing, closes the class and refutes
`P-237-3`/`P-237-5` on the seal.
