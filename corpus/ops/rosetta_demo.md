# Demo stacks — moved to `rosetta-extensions`

> **This is a pointer.** The disposable-demo-stack tooling + its full lifecycle guide now live in the
> **`rosetta-extensions`** monorepo (private: `anthropos-work/rosetta-extensions`), section **`demo-stack/`**.
> rosetta documents *how the platform works*; the tools that spin up demo copies of it live in the extensions repo.

- **What it is:** bring up `demo-N` as isolated, Clerkenstein-wired full stacks alongside the dev stack, on offset
  ports, killable cleanly — **zero read-only-platform change**.
- **The guide:** `rosetta-extensions/demo-stack/GUIDE.md` (lifecycle, port-offset scheme, the 4 injection recipes,
  `migrate-demo.sh`, safety).
- **The tooling (gitignored locally):** `anthropos-demo/rosetta-extensions/demo-stack/` — the `rosetta-demo` CLI,
  `up-injected.sh`, `migrate-demo.sh`, `inject/`.
- **The skills (here in rosetta):** [`/demo-up`](../../.claude/skills/demo-up/SKILL.md), `/demo-down`,
  `/demo-status` drive that tooling.
- **The mock it injects:** `rosetta-extensions/clerkenstein/` — see [clerkenstein.md](../services/clerkenstein.md).
