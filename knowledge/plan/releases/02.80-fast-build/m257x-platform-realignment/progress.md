---
milestone: M257x
---

# M257x — progress

## Running ledger

_(iter closeouts append here, newest last)_

- iter-01 (tok/bootstrap): all 5 open questions answered vs platform origin HEAD; authored the absent `corpus/ops/platform-alignment.md` and executed its procedure; found the class's root cause (**pinning disables drift detection** — 11/11 clones `behind: null` while the log says "provably fresh") and its local mechanism (**`migrate-demo.sh`'s hand-maintained 4-tuple** creates the legacy schemas itself, bypassing `repos.yml`); refuted 5 inherited/audited claims by measurement, one of which inverted a planned guard — see iter-01/progress.md
- iter-02 (tik): **the hand-maintained tuple is gone** — both migrate scripts now DERIVE the migration set from `repos.yml`'s machine-readable fields (origin HEAD says `app:public` alone; the tuple was wrong on **3 of 4** entries), the M810 silent-skip time bomb is disarmed, `skillpath` removed, and the non-derivable residual is declared debt behind a no-growth fence (14 tests, 4 mutations RED-proven). Re-survey **refuted TOK-01's next-tik direction** — the rext pin was already clean after the machine move (`D-M257x-4`); the real blocker was **no container runtime on this box** (Docker installed mid-iter by the user). Found a v2.1 test that **pinned the drift as a contract and still passed by reading its own refutation** (`D-M257x-6`) — see iter-02/progress.md

## Routes carried forward

| item | why | target |
|---|---|---|
| `DECIDE-M257-jobsim-schema-ownership` | The exit blocker that created this milestone: platform says cms/jobsimulation/roadrunner own no local schema; rext writes ~15 `jobsimulation.*` tables. **Inherited from M257 iter-03 — this milestone owns it now.** | iter-01+ |
| `FIX-M257-feedback-score-approximation` | Benign between a mirror and its source; **not** benign between two tables claiming to be the same row. | M257x |
| `DOC-M257-studio-in-app` | Corpus says studio-room is CMS-only in 5 places; nothing records `app` embeds it. | M257x |
| `FIX-M257-stacksnap-directus-sequences`, `FIX-M257-directus-coldstart-order` | Carried from M257 iter-02, both platform-shape-dependent. | M257x |
| `HOST-M257x-stack-demo` | No `stack-demo/` workspace on this box; gate clause 1 needs one. Docker now present (installed mid-iter-02), so this is the next executable step. | iter-03 |
| `FIX-M257x-vmram-gib-unit` | `up-injected.sh:258-262` floors bytes to integer GiB, so a VM set to the documented "12 GB" (decimal) = 11.67 GiB → floors to 11 → trips the non-fatal `< 12 GiB` warning. A doc/code **unit mismatch**, never re-measured — this milestone's own subject matter. Non-fatal. | iter-03 |
| `HOST-M257x-toolchain` | No `pytest`, `gh`, `psql` or `tailscale` on this box. Two mutation batteries (`m220`, `m255`) cannot run at all — they shell out to `python3 -m pytest` and fail with zero named tests. | iter-03 |
| `REPOINT-M257x-jobsim-writes` | ~12 `jobsimulation.*` tables (9 written) live in `stack-seeding/cmd/stackseed/main.go:45-105`. Until re-pointed, the transitional schema debt cannot shrink and **gate clause 4 cannot be met**. | later tik |
| `FIX-M257x-migrate-dev-swallows-atlas` | `migrate-dev.sh`'s atlas loop still `>/dev/null 2>&1`s every failure into "non-fatal migration warnings" — the M215-F8 masking class its demo twin already fixed. | later tik |
