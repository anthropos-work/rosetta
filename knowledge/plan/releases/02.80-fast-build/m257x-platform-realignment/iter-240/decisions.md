# iter-240 — decisions

## `D-M257x-240-1` — a prerequisite is a FLOOR, so the fence grades `<`, never `≠`

The obvious fence is *"the stated version equals the derived requirement."* It is wrong. A setup guide
saying **v1.26+** when the code needs 1.25.0 is *conservative*, not *incorrect* — a reader who installs
1.26 builds fine. Only `declared < required` breaks anything.

Asserting equality would turn every deliberate round-up into a RED and make a correct document look
broken, which is the `§8` failure mode this milestone has caught in its own instruments repeatedly. The
guard therefore compares with `<` and says so in its own output (*"floors, not pins"*), so a future reader
cannot mistake a green for a claim of exactness.

## `D-M257x-240-2` — `app` and `sentinel` are deliberately OUT of the host-Go derivation

They declare the highest Go directives in the whole clone set — `app` `go 1.26.4`, `sentinel` `go 1.26.0` —
and including them would look rigorous. It would be wrong: **both build inside Docker**, on
`golang:1.26-bookworm` (`app/Dockerfile:2`, `sentinel/Dockerfile:2`), so the host toolchain never compiles
them. A floor raised to 1.26 on their account would be a **false requirement** — a reader who installed
1.25 would be told to upgrade for no reason.

The derivation is the six `rosetta-extensions` sections that own a `go.mod` and run **on the host**
(`stacksecrets` / `stacksnap` / `stackseed` / the alignment + playthroughs binaries). Pinned by a
regression test (`test_app_and_sentinel_are_NOT_in_the_go_derivation`) rather than left to a comment,
because the next reader's instinct will be to add them.

## `D-M257x-240-3` — the under-declared prerequisites are classified, not repaired

Three declarations in the clone set are not restated in the guide, and none is a defect:

* `studio-desk` `engines.node >=24` — **satisfied** by the guide's existing global *"Node.js (v24+
  required)"*, so a reader following the guide already has it.
* `ant-academy/code` `engines.node >=22` — same, and it is separately quoted correctly at
  `run_guide.md:248`.
* `next-web-app` `packageManager: pnpm@10.30.3` — the guide says `corepack enable`, which **honours that
  field exactly**. Naming a pnpm version in prose would create a second place to rot with no gain.

The same disposition iter-237 gave its 28 orphan env names and iter-239 gave the two undeclared presets:
**an under-declaration that nothing depends on is not a false promise.**

## `D-M257x-240-4` — the guard's reach was widened before it was shipped, not after

The first selector matched `**Go**` and reached **2 of the guide's 3** Go sites. It went green, and the
green would have been *the wrong kind*: the site it could not see (`**Go 1.25.x**`, the remote-VM block)
was the one that had the number **right**, so the arm's apparent health rested entirely on the two sites
that had just been repaired — and it would have kept reading green if those two rotted back to a spelling
it also could not parse.

Widened to `\*\*Go\b` before the commit, and pinned by `test_reach_all_three_go_spellings_are_seen`.
`§5` — **a verdict without its reach is not a verdict**; the fix belongs before the fence ships, not in
the next harden pass.
