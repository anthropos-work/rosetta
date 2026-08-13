# M258 iter-01 — decisions

Intra-iter decisions. The strategy record itself is `TOK-01` in the **milestone-root**
[`decisions.md`](../decisions.md).

## D1 — the world contract resolves to **(b) restore after**, on the gate's own text

`overview.md` requires the choice at iter-01 and names two pre-authorised resolutions. **(a)
pt-world-native is refuted by the gate, not by preference**: the gate requires the stack be left *"in a
presenter-usable world"*, and the overview's own paragraph on (a) ends *"But it is not a presenter
demo."* A resolution that cannot satisfy a clause of the gate it is being chosen for is inadmissible.

(b)'s cost was already derived in the plan and is cheap for a reason worth restating: `--reset` does not
wipe the snapshot-replayed taxonomy (no catalog tables in `resetTables`, `stackseed/main.go:44-131`), so
the 78.0 s replay is **not** repaid; the stories seed measures 7.6 s; the manifests need no re-export
(`--cockpit-export` takes no DB, `stackseed/main.go:172`, and ids are deterministic). Estimated
**20–45 s** — and it is estimated, so tik 1 measures it rather than quoting it.

**Precedent this closes.** M254 left `billion` in exactly the un-restored state — *"the demo is now the
Playthrough world"* — with no restoration recorded anywhere in that milestone or its carry-forward.
Resolution (b) is what stops M258 making that swap the outcome of **every** bring-up.

## D2 — the gate is taken in single-box `--no-public-host` mode, disclosed rather than assumed

`--public-host` is default-on, and a `--public-host` demo **cannot be browsed from its own host**
(docker-proxy binds `0.0.0.0`, bypassing `tailscale serve`; `run-playthroughs.sh:92-105`). So *"one cold
command"* is literally satisfiable only with `--no-public-host`, on the host `D-v28-15` designates for
dev/test (`macmini`; `billion` is demo-deployment-only).

**The cost is stated, not hidden:** this proves the composition **in a mode the presenter never uses**.
The peer path exists and is unbroken — `--reset-only` splits the DB half from the browser half
(`run-playthroughs.sh:58-62`) — so gating the presenter mode instead is a low-cost re-cut available to
the user at any time. Recorded as `TOK-01`'s one overturnable strategic assumption.

## D3 — measure before wiring, and the ordering is the strategy

The composition has one measured half (**286.99 s**, M257 iter-09) and one that has **never** been
measured. Wiring first would place the milestone's only genuine unknown at its **end**. M257 spent
three iters demonstrating the cost of that shape: an ungradeable gate closes *"delta 0"* accurately
while nothing is learnable, and the abstention is invisible. Tik 1 therefore measures both halves in
one campaign before any wiring lands. See `TOK-01` § Rationale.

## D4 — `F1` was re-verified against code before being carried, and it survived

The prompt's standing instruction is that every inherited item is **SUSPECT-UNROUTED** until verified
open. `FIX-M257-content-stories-pair-count` initially read as already-fixed: `a5b1288` landed the
`manager_presence_only` branch and `content-pairs.ts:115` has it. **Reading the actual shell reversed
that**: `run-content-stories.sh:145-152` re-implements the count and its manager arm guards only
`has_manager_view` + path + seat, so a served manifest carrying the two manager-presence-only voice
cells counts **47** against a pinned **45** and `sys.exit(2)`s before the sweep begins. The item is
**open exactly as written**.

Two things follow. **(1)** The item's own file documents this failure for the *player* branch at
49→47 and then repeated it for the manager branch — the fourth inline copy of a counting rule drifted
again, in the file that documents the drift. **(2)** It gates the **content-stories sweep**, which is
not a Playthrough, so it does **not** block the batch — and it must not be permitted to read as a batch
blocker when it surfaces mid-campaign.
