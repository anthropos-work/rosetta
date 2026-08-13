---
iter: 185
milestone: M257x
iteration_type: tik
status: in-flight
opened: 2026-08-09
---

# iter-185 — the same defect, in a second fence: `CITATION_RE` declares which file types count

**Active strategy reference:** `TOK-08`. iter-184's lesson (*a fence's POPULATION is a registry too*)
names a class. This iter tests whether the class has a second member, and it does.

## Step 0 — re-survey before targeting

Swept `stack-core` for population-defining literals rather than predicate literals. One hit that decides
what a census can see: `predicate_enumerator.py:149`

```
CITATION_RE = re.compile(r"\b[\w./-]+\.(?:go|py|ts|tsx|js|md|tf|yml|yaml|json|sh|sql|graphqls):\d+…")
```

Thirteen extensions, typed. Measured over `corpus/**.md` + `CLAUDE.md`, in **file context** (excluding
`://` authorities):

| extension | file-context citations | declared? |
|---|---|---|
| **`mod`** | **51** | **NO** |
| **`jsx`** | **20** | **NO** |
| `hcl` · `gitignore` | 6 · 6 | **NO** |
| `example` · `ini` · `dev` · `txt` | 5 · 4 · 3 · 2 | **NO** |

**`go.mod:NN` is 51 citations, and `app/go.mod:14-18` alone is 12 of them** — the anchor CLAUDE.md's
shared-libraries banner rests on. It is invisible to the enumerator, so any predicate anchored there is
outside the reach denominator `TOK-07`/iter-114 made a declared thing.

## Cluster / target identified

The same shape iter-184 found: the assertion is fenced, the population is a literal nobody audits.

## Hypothesis

Deriving the extension cleanly **offline** is not possible — a structural rule was tested and
**refuted** (a `/` in the stem does not discriminate: `.anthropos` has 14 with-slash hits, all from
`https://` authorities), and even the `://` rule leaves `api.clerk.com:443` and
`backend.internal.anthropos:8083` behind. So iter-184's own fallback clause applies: **the declaration
is a registry and gets the both-directions treatment.**

## Expected lift

No `P`/`N` reading. Deliverable: the class extended to what is measurably cited, **plus** a fence that
turns RED the next time an extension is cited and not declared, and RED on a declared extension that
never occurs — with the authority-shaped residuals named and reasoned rather than skipped (`§5` rule 8).

## Phase plan

- **A — census** (done above), **B — extend + name the residuals**, **C — fence both directions,
  RED-proven first**, **D — close.**

## Phase 0d — pre-flight tooling check (RUN)

No new `*_guard.py`: the arms go into the existing `predicate_enumerator` test module. iter-183's lesson
(*a pre-flight that finds one precondition has not established there is only one*) is checked at close by
re-running both registry guards, not assumed.

## Escalation conditions

- If extending the class changes what the enumerator *finds*, that is a population change with
  downstream numbers attached — measure the delta and publish it, never absorb it silently.

## Acceptable close-no-lift outcomes

If the extension gap turns out not to change any enumerated site, the iter still closes complete: the
gap is measured, the tuple is fenced both ways, and the null result is stated with its number.
