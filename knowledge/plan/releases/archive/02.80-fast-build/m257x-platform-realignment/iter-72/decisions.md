# iter-72 — decisions

## `D-M257x-72-1` — `FIX-M257x-iter58-mainline-shift` closes: 66 distinct, 0 structurally broken

The route has read **"21 of 22 outstanding"** since iter-59. Re-derived under the pin rule
(`D-M257x-69-1`) at `app` `origin/main` `9d00a313`:

| | n |
|---|---|
| distinct `main.go` / `app/main.go` citations | **66** |
| graded at a ref **their own block names** (a measurement) | **37** |
| in a block naming **two or more** resolvable refs → graded at the default | **29** |
| **out of range or absent at the ref they are graded at** | **0** |

**Zero structurally broken.** The class dissolves exactly as B2 did at iter-69, and for the same
reason: the corpus writes its refs, and a citation graded at the ref it names is a measurement
rather than a defect. The route closes with a derived verdict.

**The 29 ambiguous is the honest residual** and it is larger here than anywhere else — mainline
citations cluster in the fold narrative, where a paragraph naturally names both the pin and the ref
that moved the code. It is `CHECK-M257x-iter71-ambiguous-blocks`, already routed, now with a number.

**A self-inflicted measurement bug, recorded because it is the class this milestone keeps finding.**
The first run of the derivation printed *"graded at a ref their own block names: **0**"* — flatly
contradicting the guard's own `block-pinned x31`. The script called `live.pop()` and then tested
`len(live) == 1` on the mutated set. It took one glance at a contradiction with an instrument that
had already been mutant-verified to catch it. **Two instruments disagreeing is a finding; the one
that agrees with nothing is usually the new one.**

## `D-M257x-72-2` — the eighth reach limit, and the largest: 142 citations the guard cannot see

Probing what `anchor_construct_guard` can reach in the mainline class turned up a gap that is not
about mainline at all. **Proven mechanically, both halves:**

**1. The regex never matches a bare citation.**

```
`main.go:1187`                          _QUALIFIED match: False
`app/main.go:1187`                      _QUALIFIED match: True
`internal/coursebuilder/bedrock.go:98`  _QUALIFIED match: True
```

`_QUALIFIED` requires either a `/` in the path or a `.md` suffix. A bare `<name>.<ext>:N` matches
neither, so `resolve()` is **never called** on it. Derived over the live corpus: **142 distinct
citations are outside the guard's reach entirely** — led by **41 `docker-compose.yml:N`** and
**32 `up-injected.sh:N`**, the two most-cited artifacts in the ops corpus.

**2. And the resolver would miss them even if the regex matched.** `resolve()` already has a
repo-relative rule for service docs — `root / doc.stem / cited`. For `backend.md` that is
`stack-demo/backend/`, which **does not exist**: the compose SERVICE is `backend` and the REPO is
`app`. Measured: `resolve('main.go')` from `backend.md` returns `None`; `resolve('app/main.go')`
returns the file.

**This is the eighth reach limit found in this milestone and by far the biggest.** The pattern is
now unmistakable and worth stating plainly: **every one of the eight was found by reading or by
probing, never by a GREEN verdict** — because a fence reports on the class it can see, and its
silence about everything else is indistinguishable from health.

**The fix is designed, not landed** (scope-creep tripwire — this is a third line in a two-line iter,
and it will turn the guard RED on real sites, which is a repair pass of its own):

- widen `_QUALIFIED` with a third alternative for a bare `<name>.<codeext>:N`;
- **derive the doc-stem → clone mapping from compose** rather than listing it —
  `docker-compose.yml`'s `backend` service declares `build.context: ../app`, so the service→repo
  edge is an artifact fact. `platform_predicate_guard.parse_compose` is already an importable,
  tested primitive and `D-M257x-59-2` explicitly sanctions reusing it this way;
- keep `AMBIGUOUS_BASENAMES`' discipline: `main.go` exists in **seven** clones, so a bare basename
  must resolve through the doc's own service, never by a tree-wide basename search. That is exactly
  the over-match the guard's docstring records as *"134 findings, essentially all of them ports."*

Routed as **`FENCE-M257x-iter72-bare-citation-reach`**, with the measurement, both proofs, and the
design above as its brief. **It is a prerequisite for the graded read** — and unlike iter-69's
routed prerequisite, which iter-70 falsified, this one is proven by two mechanical probes rather
than by a count.
