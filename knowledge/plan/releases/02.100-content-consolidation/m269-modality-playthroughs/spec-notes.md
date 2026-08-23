# M269 — Spec notes

_None yet._ Section headers below are derived from the milestone's `In:`/`Out:` scope and the open
questions; they are placeholders, not content.

## The assertion boundary — where it stops today

_None yet._

## Where the boundary must move to, and what "to completion" means per modality

_None yet._

## The session-creation oracle

_None yet._

## Modality derivation — the predicate, not a keyword

_None yet._

## Slug pinning — replacing the single `SAMPLE_CHAT_SIM_SLUG`

_None yet._

## Manifest changes — `ai-simulations.yaml` and the 4-state reporting map

_None yet._

## The three use cases inherited from the dissolved M206

_None yet._

## `FIX-M256-studio-false-green` — fixing the oracle, then re-running the 2026-08-23 claim

_None yet._

## `BIND_HOST` / `D-M255-7` — making the batch gate stop skipping

_None yet._

## Zero-platform-edit boundary — when a demopatch is the answer and when it escalates

_None yet._

## Line-anchor re-pins measured during scaffolding (2026-08-23)

The scaffolding pass re-resolved three of the brief's anchors against what is on disk. The brief's
citations are carried **verbatim** in [`overview.md`](overview.md) as the milestone contract; these are
the drifts to re-pin at milestone start, **not** corrections to the contract:

| contract citation | re-measured on disk (2026-08-23) |
|---|---|
| `up-injected.sh:146` (`BIND_HOST`) | `demo-stack/up-injected.sh:164` in the authoring copy at `v2.9.23-rext` (`bfd9835`): `if [ -n "${STACK_PUBLIC_HOST:-}" ]; then BIND_HOST="0.0.0.0"; else BIND_HOST=""; fi`. Line drift only — the predicate is unchanged. |
| `useGetSimulationFlagsAndFeatures.ts:22-68` | the file on disk is **`.tsx`**, not `.ts` (`stack-demo/next-web-app/packages/graphql/src/hooks/aiSimulation/useGetSimulationFlagsAndFeatures.tsx`). The three predicates resolve inside `:20-68` there: `hasCall` `:20-27`, `isCodingSequence` `:37-43`, `isDocumentChallenge` `:64-66` via `codingFileName` `:44-58`. |
| `playthroughs/e2e/tests/aisim-chat-launch.spec.ts` header | verified present and exact at `:5-9` of the authoring copy. |

`playthroughs/manifest/ai-simulations.yaml:7-11` and `e2e/lib/simulation-page.ts:27` both verified exact.

## `FIX-M256-studio-false-green` — the anchor the carry entry names

`roadmap-vision.md` pins the false-green locator at `playthroughs/e2e/lib/studio-builder-page.ts:120`.
On disk the surrounding note reads, in the source's own words: *"the 5-minute poll also showed the page
sitting on `/sim-advanced-builder` for the full duration with the sections present from +2.1 s, so 'did
the generation actually COMPLETE on this host?' is still open — presence of a heading does not answer it.
Measure section CONTENT, not the heading, to answer both questions at once."*
