# iter-18 — decisions

## D-M257x-18-1: the hand-off's "you need a cold cycle to diagnose this" was wrong, and testing it was cheap

**Decision.** Diagnose the bootstrap failure by reproducing it in isolation before spending ~11 minutes on a
cold cycle.

**What the hand-off said.** iter-17: *"Do not try to reproduce the bootstrap failure by hand — it has healed.
… The failing state only exists on a freshly purged stack, which means the diagnosis arrives through a cold
cycle with the capture fix in place, not through poking demo-1."*

**What measuring cost, and what it bought.** Four minutes and five observations, each of which changed the
target:

1. the identical `docker run … node cli.js bootstrap` against a **throwaway empty schema on the same
   Postgres** exits **0** — so the command is fine and the failure is context;
2. `demo-1`'s `directus` schema is **fully bootstrapped** (all seven system tables, 87 migrations) — which
   **refutes** iter-17's stated mechanism (*"the system schema … was never bootstrapped"*);
3. the public **policy, permission and access rows all exist** — the thing said to be missing was present;
4. `docker restart demo-1-directus-1` alone flips the anon read **403 → 200**;
5. `docker inspect directus/directus:11.6.1` → `CMD … node cli.js bootstrap && pm2-runtime start` — **the
   image bootstraps itself**, and the container's own timestamped log shows it winning at
   `16:02:28.98 → 29.40`, matching `directus_migrations` to the millisecond.

**The generalisable part.** *"The failing state has healed"* was true of **one** artefact (the `directus`
schema on that one stack) and was silently generalised to *"the failure is not reproducible."* A schema is
cheap to recreate; the bootstrap's inputs are a DSN, an image and an empty schema, none of which needed a
bring-up. **Ask what the failing step's actual inputs are before concluding that only the whole pipeline can
produce them** — the alternative is paying 11 minutes per hypothesis. Promoted to `platform-alignment.md`
§5 as rule 15.

## D-M257x-18-2: the outcome is a POST-CONDITION, not a race winner

**Decision.** `provision_directus_step` reports PROVISIONED when the sentinel system table is present after
the attempt, whether or not our own `docker run` exited 0.

**The rejected alternative — make our one-shot win.** Suppressing the compose service's own bootstrap
(overriding the image `CMD`, or delaying the service until set-dress has run) would remove the race by
removing a racer. It was rejected: the image's entrypoint is *the platform-shaped behaviour* — every future
Directus version ships it, `restart: on-failure` (iter-05's own fix) is what makes it converge, and an
override that fights it is a hand-maintained contradiction of an upstream default, i.e. the class this
milestone exists to end. **We do not need a winner; we need the schema.** Asking the database is §8 rule 4 —
a construct that cannot express the drift beats a rule about who should go first.

**The second reason it is the right shape.** It is *self-healing across the two host families*. Whichever
racer wins on a laptop vs a Linux VM (different `host.docker.internal`, different image-pull latency,
different Postgres readiness), the post-condition is the same question with the same answer.

## D-M257x-18-3: the serve restart is decoupled from the provision's exit code

**Decision.** `boot_directus_step` runs when the **directus replay succeeded on a `--local-content` stack**,
not when `DIRECTUS_PROVISIONED = 1`.

**Why this is the causal fix and not a belt-and-braces.** D-M257x-18-2 alone would have made
`DIRECTUS_PROVISIONED = 1` on the observed run and the restart would have followed. But that couples a
restart to a *different step's* exit code, which is the defect's shape rather than its instance: any future
provision failure — a genuinely broken bootstrap, a `CREATE SCHEMA` timeout — would again silently cancel a
restart the replay had just made necessary. `test_serve_restart_is_not_gated_on_the_provision_exit_code`
pins the decoupling on the case where the provision *really* failed, so the two fixes are independently
falsifiable (M1 and M2 of the battery kill different tests).

## D-M257x-18-4: three tests were arguing for the "prod-read" claim, and the claim is false

**Decision.** Rewrite `test_bootstrap_failure_degrades_to_prod_read_nonfatal` (renamed),
`test_create_schema_failure_degrades_nonfatal` and
`test_no_snapshot_with_local_content_skips_provision_no_setu_trip` to assert
`content:local-content-UNPROVISIONED`, and to assert `content:prod-read` is **absent**.

**Why the old assertion encoded a bug.** `gen_injected_override.py:580` appends
`DIRECTUS_BASE_ADDR=<in-network>` whenever the local-content directus service is emitted. A failed provision
does not undo that. So on a `--local-content` stack a failed provision leaves the consumer pointed at an
**empty local Directus** — not at prod. The old message (*"the stack stays on the prod-read path"*) and the
old verdict field (`content:prod-read`) both told the operator to look in the wrong place, and three tests
required them to keep doing so. Second occurrence in this milestone of *"the suite was not silent about the
defect — it was arguing for it"* (iter-16).

**The tell that it was a false claim rather than a stale one:** the same closing line printed
`content:prod-read` **and** `directus=replayed` in one sentence. A stack cannot both read its content from
prod and have had 11 986 rows replayed into its own Directus. `CHECK-M257x-iter17-setdress-verdict-
contradiction` is closed by making the field derivable from what actually happened, not by rewording it.

## D-M257x-18-5: the ADMIN_TOKEN consequence is ROUTED, not landed

**Decision.** `FIX-M257x-iter18-directus-admin-token-race` goes to a later tik.

**The measurement.** The winning bootstrapper is the compose service, whose environment carries no
`ADMIN_EMAIL` / `ADMIN_PASSWORD` / `ADMIN_TOKEN`. So the admin it creates is directus's default
`admin@example.com` with **`token = NULL`** — measured on demo-1 — while M23's contract says the admin
carries the locally-minted static token `local-directus-token-<stack>` that studio-desk consumes as
`DIRECTUS_TOKEN`. **On every `--local-content` stack where the container wins the race, that token does not
exist.** Nothing in the current gate reads it, which is why it has never surfaced.

**Why route rather than land.** The repair belongs in the two override emitters (`gen_injected_override.py`
+ `gen_override.py`) plus their parity fence — so that *whichever* racer bootstraps produces the same admin —
and that is a third line of investigation in an iter whose declared shape is two. Scope-creep tripwire
respected. It is well-formed: the fix is known, the measurement is recorded, and its own verification is a
one-line `SELECT token FROM directus.directus_users`.

**Also routed:** `CHECK-M257x-iter18-directus-secret-naming` — `dev-setdress.sh` derives the SECRET/token
suffix with `tr '-' '_'` (`demo_1`) while both emitters and `provision.go`'s `DefaultEnvContract` use
`demo-1`. Two spellings of one contract, in a value both sides must agree on. Observed while reading the
env block; **not** measured for consequence, and recorded as an observation rather than a defect claim.
