#!/usr/bin/env python3
"""measure_mainline_shift.py — M257x iter-58 (TOK-04 P2: the instrument is a committed file)

Measures how many `main.go:N` citations in the corpus point at DIFFERENT CONTENT after an
`app` advance. Run from the rosetta repo root.

    python3 knowledge/.../iter-58/instrument/measure_mainline_shift.py v1.365.0 v1.366.0

WHAT IT DOES AND DOES NOT ASSERT.
  It compares the *content of line N* at two app refs. A row reported MOVED means the citation
  now lands on different text than it did at the baseline ref.
  It does NOT assert the citation was CORRECT at the baseline: corpus citations were written
  against assorted app refs (platform-migration-status.md:70 names `5ba17044` v1.363.2 in its own
  prose). So MOVED is a lower bound on instability, not a count of newly-false claims. The
  decisive measurement — does each citation resolve to what its prose SAYS it contains, today —
  is FIX-M257x-iter58-mainline-shift and needs a per-citation expected-construct, which only
  some rows carry.

WHY IT EXISTS. `anchor_construct_guard` caught 1 of these 22 (it fires only when the cited line
is not a construct at all, e.g. a bare `}`). The other 21 land on plausible-but-wrong lines,
which is FIX-M257x-iter57-within-block-drift measured on a real event.
"""
import re, subprocess, sys

APP = "stack-demo/app"

def lines(ref):
    return subprocess.run(["git", "-C", APP, "show", f"{ref}:main.go"],
                          capture_output=True, text=True, check=True).stdout.split("\n")

def main():
    base, head = (sys.argv[1], sys.argv[2]) if len(sys.argv) > 2 else ("v1.365.0", "v1.366.0")
    old, new = lines(base), lines(head)

    out = subprocess.run(["grep", "-rn", "-E", r"main\.go:[0-9]+", "corpus/", ".claude/", "CLAUDE.md"],
                         capture_output=True, text=True)
    # §5 rule 1: never let a search's stderr go unread. rc 0=hits, 1=no hits; anything else is an
    # engine failure that would otherwise read as "no citations exist".
    if out.returncode not in (0, 1):
        sys.exit(f"GREP FAILED (rc={out.returncode}): {out.stderr}")
    if out.returncode == 1:
        sys.exit("GREP returned NO HITS — that is a broken pipeline, not a clean corpus (positive-control rule)")

    rows = []
    for ln in out.stdout.strip().split("\n"):
        if not ln:
            continue
        f, no, rest = ln.split(":", 2)
        for m in re.finditer(r"(app/)?main\.go:(\d+)(?:-(\d+))?", rest):
            n = int(m.group(2))
            o = old[n - 1] if n - 1 < len(old) else "<EOF>"
            w = new[n - 1] if n - 1 < len(new) else "<EOF>"
            rows.append((f, no, m.group(0), n, o.strip(), w.strip(), o == w))

    moved = [r for r in rows if not r[6]]
    print(f"app {base} -> {head}")
    print(f"citations examined: {len(rows)}   unchanged: {len(rows)-len(moved)}   MOVED: {len(moved)}")
    print()
    for f, no, cite, n, o, w, _ in moved:
        print(f"{f}:{no}  {cite}")
        print(f"    {base} -> {o[:70]}")
        print(f"    {head} -> {w[:70]}")
    return 0

if __name__ == "__main__":
    sys.exit(main())
