# iter-94 — decisions

## `D-M257x-94-1` — the two suspects are NOT the same finding, and separating them was the work

iter-91 surfaced both as "GREEN against a tree with no corpus". Adjudicated separately, because a shared
symptom is not a shared cause (§5 rule 28):

| guard | subject | verdict |
|---|---|---|
| `union_apply_guard` | rext's OWN demopatch manifest set (`demo-stack/patches/*/*.yaml`) | **correct by design** — it asserts a property of rext, resolves from `parents[1]`, and is unaffected by `--repo-root` because the corpus is not its subject. Not changed. |
| `story_org_count_guard` | *"stories.seed.yaml ships N orgs **and every doc agrees**"* — the docs are the **corpus** | **real defect** — fixed |

Recording the negative half matters as much as the fix: **"it ignored `--repo-root`" is only a defect if
`--repo-root` names its subject.** Changing `union_apply_guard` to honour it would have broken it in a
rext-only checkout, which is how the guard family is consumed per-stack.

## `D-M257x-94-2` — the defect: a control that could never fire, guarding a vacuous claim

`scan_roots()` returns the rosetta corpus + skills **plus rext's own two directories**. This guard *lives*
in rext, so those two always exist — therefore the existing control

```python
if not roots:  ...  return 2
```

**could never fire.** A run whose rosetta half was missing scanned only rext, found nothing to contradict,
and printed *"and every doc agrees"* with `rc=0`.

**"Every doc agrees" over ZERO corpus docs is vacuously true and reads exactly like a pass** — §5 rule 8, in
the family whose green this milestone quotes.

Two rungs added, both mutation-checked:

1. **A positive control on the corpus half specifically** — not on "some root exists", which is the check
   that could never fail. Unreachable corpus ⇒ **exit 2 CANNOT RUN**.
2. **The cardinality is printed**: *"all **116** scanned doc(s) agree"*. §8 already required stating how
   many; without it, a vacuous scan and a real one produce the same sentence.

Measured before / after, invoked exactly as `guard_family` does it:

| tree | before | after |
|---|---|---|
| empty | `rc=0 OK — … every doc agrees` | **`rc=2 CANNOT RUN — Nothing was checked; this is not GREEN`** |
| real | `rc=0 OK — … every doc agrees` | `rc=0 OK — all 116 scanned doc(s) agree` |

## `D-M257x-94-3` — this is the THIRD anti-vacuity defect this session, and the pattern is now the finding

- iter-91: `platform_alignment_guard` graded total resolver failure but not **partial** blindness.
- iter-93: the new guard's own live-corpus control **silently skipped** on a hardcoded `parents[3]`.
- iter-94: `story_org_count_guard`'s emptiness control **could never fire**.

Three different guards, three different authors' assumptions, one shape: **the check that would catch "I
checked nothing" was itself the weakest check in the guard.** Every one of them was written by someone who
had just read §5 rule 8.

The generalisation worth carrying, and it is sharper than "add an anti-vacuity rung": **an anti-vacuity
control must be written against the SUBJECT of the guard, not against its inputs.** All three failures are
the same substitution — `roots` existed, a *file* was found, *some* citations resolved — while the thing the
guard exists to talk about was absent. Ask *"did I look at the thing I am about to make a claim about?"*,
and count it out loud.
