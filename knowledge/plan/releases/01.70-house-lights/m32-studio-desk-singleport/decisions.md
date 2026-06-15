# M32 — Decisions

_Implementation decisions with rationale, numbered `M32-D1`, `M32-D2`, … . Empty at scaffold; filled during build._

_Pre-decided at design (2026-06-15):_
- _Root-cause fix is `NODE_ENV=production` on the studio-desk override (the additive env block lets the base
  `development` survive → the dev redirect to dead `:9100`). Production serves via `sendFile`, no cross-port redirect._
- _Remove the un-offset `:9100` CORS origin (dead once studio-desk is single-port production) — record as an explicit decision._
