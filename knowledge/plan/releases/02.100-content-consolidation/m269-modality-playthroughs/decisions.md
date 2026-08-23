# M269 — Decisions

Decisions are recorded as `D-M269-N` entries, most recent last. Index:

| id | title | status |
|---|---|---|
| `D-M269-1` | M206 is **DISSOLVED**, not re-reserved a sixth time | recorded at design time |
| `D-M269-N` | _(none yet)_ | — |

## `D-M269-1` — M206 is DISSOLVED, not re-reserved a sixth time

[`roadmap-vision.md:311-322`](../../roadmap-vision.md) states that a sixth re-reservation of M206
**"is not an option this file permits"** — and it happened anyway. The v2.10 design run resolves that by
**dissolving the reservation**, not by re-fating it:

- `ai-simulations.code.UC1` → **M269** (this milestone)
- `ai-simulations.interview.UC1` → **M269**
- `profile.self-evaluation.UC1` → **M269**
- voice + recording + skill-paths verify → **M271**

**No M206 reservation survives this release.** After v2.10 closes, `M206` names nothing: any future work
in that space gets a new `Mxyy` at its own design time, per the `Mxyy` rule.

## `D-M269-2` — _(reserved: the assertion-boundary move)_

The boundary move is a **policy change**, not a test tweak: `playthroughs/manifest/ai-simulations.yaml:7-11`
states the non-voice launch-only boundary AS POLICY, and it is the bulk of this milestone's cost. It must
be argued and recorded here before the first Playthrough is rewritten.

_Not yet decided._
