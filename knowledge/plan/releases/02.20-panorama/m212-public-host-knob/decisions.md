# M212 — decisions

_(implementation choices with rationale, recorded as the milestone is built)_

## D-DESIGN-1 — opt-in, default off
`STACK_PUBLIC_HOST` defaults to `localhost`; external reach requires an explicit `/demo-up --public-host` flag.
**Why:** user directive (2026-07-11) — public access must be explicitly requested at build time, never ambient.
