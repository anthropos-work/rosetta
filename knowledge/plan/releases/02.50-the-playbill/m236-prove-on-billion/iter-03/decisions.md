# iter-03 — Decisions

## D1 — a remote bring-up must be driven through a LOGIN shell (PATH, not prereqs)

**Symptom:** `up-injected.sh` host pre-flight failed with *"Go NOT on PATH … install Go 1.25.x"*.

**Reality:** Go **is** installed on `billion` — `/usr/local/go/bin/go`, `go1.25.12`, exactly the pinned
`toolchain go1.25.12`. `/usr/local/go/bin` is added to PATH by the **login profile**, which a
non-interactive `ssh host 'cmd'` never sources.

**Why it is a trap rather than a footnote:** `atlas` lives in `/usr/local/bin`, already on the default
non-login PATH, so it passed — the two prereqs the F-series calls out behave *differently* under the same
invocation, which makes the failure look prereq-specific ("Go is missing") rather than shell-specific
("nothing from the profile is on PATH"). The pre-flight's own remedy text reinforces the wrong reading by
recommending an install. The cheap disproof is one command: `ssh host 'bash -lc "go version"'`.

**Resolution:** drive remote bring-ups as `ssh host 'bash -lc "…"'`. Applied; the pre-flight then passed.

**Route:** the finding belongs in `corpus/ops/demo/tailscale-serve.md`'s F1–F12 host-prereq set (the doc
that owns "the fresh-Linux-VM host prereqs the tooling pre-flights/auto-handles/fails-loud on"). Folded
into the B4/B5 doc pass rather than done here — handler `DOC-M236-iterTBD-protocol-backfill` — because
that pass is already opening the same family of docs and a one-line drive-by edit now would fragment it.

## D2 — the substrate reading is NOT recorded as a metric lift

13/13 simulation sessions are present with attempt-result rows **and** manager-mirror rows; ai-labs
presence rows = 2 (target met). That is up to 26 of the 31 pairs' worth of *data*.

The metric was still recorded as **0/31**. The gate measures **"renders real, non-empty content, live"**,
and no render has been proven — the seat-login harness does not exist yet (Phase H, iter-04). The
protocol's honesty self-check forbids claiming an un-probed delta, and this milestone was re-scoped
specifically because gate clauses that could not be measured had gone unnoticed for four releases.
Recording 26 would repeat that failure in the opposite direction.

The reading's real value is **risk**, not score: it says the remaining simulation-arm work is render/route
work, not seeding work.

## D3 — academy tables are empty; routed, not fixed here

`academy_chapter_progresses` / `academy_skill_paths` / `academy_chapters` all **0**. The academy catalog +
progress fill is not wired into the cold bring-up. This is a **named in-scope M236 item** (wire
`app/cmd/academy-seed --user-id <academy content-player owner> --fixture …`), not a new discovery.

Fixing it here would have been the scope-creep tripwire's third line of investigation in a bring-up iter.
Routed to the academy arm (handler `ACADEMY-M236-iterTBD-catalog-fill`) per Fate-3 with a named handler.
