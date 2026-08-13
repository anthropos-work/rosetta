# iter-236 — decisions

## `D-M257x-236-1` — prefix drift is MEASURED, not repaired

**6 of 16** repos are reached under more than one prefix in fenced `cd` blocks (`platform` 29 bare / 17
`stack-dev/`; `ant-academy` bare / `stack-dev/` / `stack-demo/`; plus `app`, `next-web-app`,
`studio-desk`, `studio-room`). `P-236-4` predicted it and it is confirmed.

**No site was changed.** A bare `cd platform` is *correct* after a preceding `cd stack-dev`, and the
instrument reads one line at a time — it cannot see the preceding line. Repairing on that evidence would
repeat iter-235's denominator error (`D-M257x-235-1`) at 46× the volume, in the direction that damages
correct documents.

Recorded as a measured property of the corpus. It is real information — the workspace convention moved
twice and the prose still carries both eras — but it is not a defect count.

## `D-M257x-236-2` — the disclosure window is the LIMIT; it is recorded, not widened

`corpus/services/chronos.md:177/:196` were flagged `UNDISCLOSED` by a ±12/+6-line context window. They are
**not** defects: the document opens with *"⚠️ Decommissioned … no longer cloned by `make init` … the
detail below is preserved for historical context."*

The window was **not** enlarged to make the flag disappear. Enlarging a threshold until a known-good case
passes is fitting the rule to the sample (`§5` Trap A) and would silently un-flag real defects at the same
time. The limit is stated in the iter's findings, and the general repair — disclosure is a property of a
**document**, not a window — is routed as `ROUTE-M257x-236-disclosure-scope-is-document-level`.

## `D-M257x-236-3` — the four repairs point at `app/studio`, modelled on the site that was already right

`corpus/services/cms.md:298-308` already carried the correct, fully-reasoned block (`cd app/studio  # was:
cms/studio`, plus why `cms` is not cloned). The repairs to `CLAUDE.md`, `studio-room.md`, `run_guide.md`
and `update_guide.md` were written to **match it**, not to invent new wording, and each cites the same
primary evidence: `app/.gitignore:78-79` (*"pulled at build via additional_repo, like cms"*) and an empty
`git ls-tree -d HEAD studio`.

`run_guide.md` and `update_guide.md` keep the hand-clone path as an explicit alternative rather than
deleting it — `anthropos-studio-room` is a real repo you may legitimately have cloned; what was missing
was that `make init` never does it for you.

## `D-M257x-236-4` — existence was read from the CLONE'S GIT TREE, never from the filesystem

`stack-demo/app/studio` and `stack-dev/studio-room` **both exist on this box**. Had the census used
`os.path.isdir`, three of the four defects would have read as clean. Every subdirectory verdict comes from
`git ls-tree -d -r HEAD` on the clone, so a git-ignored, build-populated directory cannot vouch for
itself. Routed as the general rule (`ROUTE-M257x-236-host-is-the-unreliable-witness`).
