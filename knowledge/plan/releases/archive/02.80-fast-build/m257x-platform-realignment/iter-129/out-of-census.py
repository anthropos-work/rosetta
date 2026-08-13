#!/usr/bin/env python3
"""iter-129 — the OUT-OF-CENSUS consequence surface, enumerated by the census's own machinery.

iter-128 § 3d measured that C1 covers **37.1 %** of the corpus's consequence-class sentence
surface: the census reads `corpus/services/*.md` + `corpus/architecture/*.md` and NOTHING else,
so `corpus/ops/**`, `corpus/tools/**` and `CLAUDE.md` are outside it entirely.

This enumerates that outside. It **imports** `claim_census_guard` (the census) and
`iter-124/triage-predicate.py` (the sealed C1 regex) rather than copying either, so the two
enumerations cannot drift and the published figures stay byte-reproducible.

TWO accountings, deliberately kept apart (the run-82 directive, and `F4`):
  * clause 5's scope is `corpus/services/**` + `corpus/architecture/**` and is NOT re-cut here;
  * this is the USER's standing ask — the corpus aligned to the platform — and its numbers are
    reported under their own denominator, never folded into clause 5's.

  /usr/bin/python3 out-of-census.py <rosetta-root> [--list]
"""
import importlib.util
import pathlib
import re
import sys

HERE = pathlib.Path(__file__).resolve().parent
MSTONE = HERE.parent


def _load(name: str, path: pathlib.Path):
    spec = importlib.util.spec_from_file_location(name, path)
    mod = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(mod)
    return mod


def main(argv):
    root = pathlib.Path(argv[1]).resolve()
    want_list = "--list" in argv

    census_py = root / ".agentspace/rosetta-extensions/stack-core/claim_census_guard.py"
    cg = _load("claim_census_guard", census_py)
    pred = _load("triage_predicate", MSTONE / "iter-124/triage-predicate.py")
    C1 = pred.C1

    names = cg._live_names(root, None)

    # The census's own SCOPE, and its complement inside the corpus.
    in_scope = list(cg.SCOPE)
    out_scope = ["CLAUDE.md", "corpus/ops/*.md", "corpus/ops/**/*.md", "corpus/tools/*.md",
                 "corpus/tools/**/*.md", "corpus/*.md"]

    def sweep(patterns):
        seen: set[pathlib.Path] = set()
        rows = []
        stats = {"files": 0, "sentences": 0, "assertion": 0, "conseq_all": 0,
                 "conseq_assertion": 0, "conseq_unevidenced": 0}
        for pattern in patterns:
            for path in sorted(root.glob(pattern)):
                if path in seen or not path.is_file():
                    continue
                seen.add(path)
                rel = str(path.relative_to(root))
                stats["files"] += 1
                for b in cg.blocks_of(path, rel):
                    if b.kind == "heading":
                        continue
                    cits = cg.citations_of(b.text)
                    hedged = cg.is_hedged(b.text)
                    for sent in cg.sentences_of(b.text):
                        stats["sentences"] += 1
                        is_assert = not (cg.is_imperative(sent) or sent.endswith("?")) \
                            and cg.has_subject_token(sent, names)
                        if is_assert:
                            stats["assertion"] += 1
                        if not C1.search(sent):
                            continue
                        stats["conseq_all"] += 1
                        if not is_assert:
                            continue
                        stats["conseq_assertion"] += 1
                        if cits or hedged:
                            continue
                        stats["conseq_unevidenced"] += 1
                        rows.append({"path": rel, "start": b.start, "end": b.end,
                                     "sentence": sent})
        return stats, rows

    ins, _ = sweep(in_scope)
    outs, rows = sweep(out_scope)

    print("== census SCOPE (clause 5's scope — NOT re-cut, reported for the ratio only) ==")
    for k, v in ins.items():
        print(f"  {k:20s} {v}")
    print("== OUTSIDE the census (CLAUDE.md + corpus/ops/** + corpus/tools/** + corpus/*.md) ==")
    for k, v in outs.items():
        print(f"  {k:20s} {v}")

    tot_all = ins["conseq_all"] + outs["conseq_all"]
    print(f"\nconsequence-class sentences, ALL          : in {ins['conseq_all']} / "
          f"out {outs['conseq_all']} / total {tot_all}"
          f"  -> census reach {100*ins['conseq_all']/tot_all:.1f} %")
    tot_un = ins["conseq_unevidenced"] + outs["conseq_unevidenced"]
    print(f"consequence-class, UNEVIDENCED assertions  : in {ins['conseq_unevidenced']} / "
          f"out {outs['conseq_unevidenced']} / total {tot_un}"
          f"  -> census reach {100*ins['conseq_unevidenced']/tot_un:.1f} %")

    import collections
    per = collections.Counter(r["path"] for r in rows)
    print(f"\nout-of-census UNEVIDENCED consequence sentences: {len(rows)} in {len(per)} files")
    for p, n in per.most_common():
        print(f"  {n:4d}  {p}")

    if want_list:
        print("\n--- the read set ---")
        for r in rows:
            print(f"{r['path']}:{r['start']}-{r['end']}  {r['sentence'][:300]}")
    return 0


if __name__ == "__main__":
    sys.exit(main(sys.argv))
