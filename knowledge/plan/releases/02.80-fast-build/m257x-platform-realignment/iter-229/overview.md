---
iter: 229
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-09
---

# iter-229 — `buildbench` accepts a profile that does not describe the host it runs on

## Step 0 — Re-survey before targeting

`TOK-08`'s next-tik direction is generic ("work the mechanical classes in descending measured size"). The
**user's redirect governs target selection** and ranks two things above the instruments: (a) the corpus's
claims about the platform, (b) **an actual working stack**. The working-stack half is gate clause 1, and
iter-225 measured why it cannot be graded today:

> **No measured host profile exists for the only machine this release may test on.** Both checked-in
> profiles (`billion.json`, `laptop.json`) describe other hardware, and **nothing in `buildbench` compares
> the operator-supplied `--profile` to the host it is running on.**

iter-225 routed two items out of that. One (`225-no-profile-for-sanctioned-host`) needs a **quiet box** and
this host is not quiet — agents are running. The other is landable now and is this iter:

- `ROUTE-M257x-225-profile-vs-host-identity-check` → *"`buildbench` should refuse a profile that does not
  describe the host it is running on, mirroring the `autoverify` verdict guard at `:809`. `host_facts()` is
  already collected; nothing consumes it for this."*

**Re-survey confirms the target is untouched.** At `stack-core/buildbench.py`:

- `host_facts()` (`:858`) collects `hostname` / `kernel` / `cores` / `python` / `docker`.
- Its ONLY consumer is `:905`, `entry["host"] = host_facts() | {"profile": profile_name}` — written into
  the rep ledger as a **record**, read by no assert.
- `pre_rep_assert(profile, *, lanes)` (`:718`) takes no host argument at all.

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.* Profile-host
identity is mechanically decidable: a profile either describes the host or it does not, and no sentence has
to be interpreted. The redirect selects WHICH mechanical class; `TOK-08` still supplies the method.

## Hypothesis

The same discipline the harness already applies one object too late — `read_verdict()` (`:809`) refuses an
`autoverify.json` that *"describes an earlier stack"* — belongs on the **profile**, which decides whether a
run was worth attempting at all. Landing it makes the sanctioned-host gap **loud** instead of silent, which
is the precondition for `225-no-profile-for-sanctioned-host` ever being closed honestly.

## Predictions — SEALED BEFORE ANY MEASUREMENT

| id | prediction | why it matters |
|----|-----------|----------------|
| `P-229-1` | `host_facts()`'s output is consumed by **0** assert/gate call sites — every use is a ledger record | confirms the route's premise from the code, not from iter-225's prose |
| `P-229-2` | This host's `arch` **MATCHES** `laptop.json`'s (`arm64`), so an arch-only identity check would **not** catch the mismatch iter-225 found | a check on the cheapest field is self-passing — the `§5` self-matching-locator class |
| `P-229-3` | This host's core count differs from **both** profiles (billion 8, laptop 10) | the field that can actually decide it |
| `P-229-4` | Neither profile stores a hostname; `name` is a role label, so identity **cannot** be decided on `platform.node()` | rules out the obvious implementation |
| `P-229-5` | `pre_rep_assert` has no host parameter — the check cannot be added there without a signature change | scopes the edit |
| `P-229-6` | On a `docker-desktop-vm` profile, `profile["cores"]` is the **VM allocation** while `host_facts()["cores"]` is `os.cpu_count()` (**host total**) — a naive comparison compares two different quantities | the trap that would make the new check fire falsely; `budget_source` says so in `laptop.json` |

## Expected lift

No `N`/`P` reading. The deliverable is a landed, tested refusal in `buildbench` + the measured host facts of
the sanctioned host recorded where the next iter can use them.

## Phase plan

1. Seal these predictions (probe commit) — done before any measurement.
2. Measure: `host_facts()` on this host; grade `P-229-1`…`P-229-6`.
3. Implement `profile_describes_host()` + wire it into the campaign path; three verdicts, not two
   (`§5`: pass / fail / **UNMEASURED**), fail-closed.
4. Regression tests, both arms (a matching profile accepts; a mismatched one refuses) — per `§8`
   *"guards must be tested in PAIRS"* and *"a fence publishes its FIRE side and hides its ACCEPT side."*
5. Run the `stack-core` pytest section; push the rext tag-less commit to origin.

## Escalation conditions

- If the check cannot be made to distinguish "wrong host" from "unreadable host", it must return
  **UNMEASURED**, never a pass — and that is a landable outcome, not a blocker.
- A platform-repo edit is out of scope (milestone constraint). Nothing here touches one.
- Opening a **third** line (e.g. also authoring the sanctioned-host profile, which needs a quiet box) fires
  the scope-creep tripwire → route forward.

## Acceptable close-no-lift outcomes

If `P-229-3` is REFUTED — i.e. this host's cores coincide with a checked-in profile — the identity check
would be **unable to catch the very case that motivated it**, and saying so with the measurement is the
iter's deliverable.
