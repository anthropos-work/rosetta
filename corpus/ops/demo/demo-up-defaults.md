# `/demo-up` — the defaults contract

**The enumerated list of every knob and flag that controls a demo bring-up, with its real default and the
exact line that reads it.** Before v2.3 M220 no such contract existed anywhere in the corpus: the only
complete knob list was a **skill `argument-hint`** — one line of prose, never checked against any parser.

This page is **derived from the parsers, not from memory**, and it is fenced by
`rosetta-extensions/stack-core/demo_knob_guard.py`, which compares it to the scripts **in both directions**:

- **docs → parser.** A knob or flag documented here that no script reads is a **FALSE PROMISE**.
- **parser → docs.** A knob or flag a script accepts that is missing here is **UNDISCOVERABLE**.

Neither check alone is enough, and the failure mode that actually shipped was the first one.

---

## ⚠️ There are TWO entry points, and they do NOT take the same flags

This is the single most important fact on this page, and conflating the two is the bug that made the contract
necessary.

| Entry point | Accepts | On an unknown argument |
|---|---|---|
| **`demo-stack/up-injected.sh`** — the **full bring-up**; what `/demo-up N` actually runs | `<N>` (positional) and **`--public-host <host>`** — *that is the entire flag surface* | **hard-errors:** `unknown argument` → **`exit 1`** (`up-injected.sh:26-27`) |
| **`demo-stack/rosetta-demo`** — the **low-level lifecycle wrapper** (`up` / `down` / `gen` / `status` …) | `--profile`, `--services`, `--ref`, `--only`, `--resolve-only`, `--fapi-host`, `--bapi-ip`, `--webhook-secret` | per-subcommand |

> **The `/demo-up` skill's `argument-hint` used to list `--public-host`, `--profile` and `--services` together,
> as if one parser took them all.** It does not. `up-injected.sh N --profile cms` **exits 1**. `--profile` and
> `--services` are `rosetta-demo` flags — reachable via `rosetta-demo up N --services "postgresql redis"`, which
> brings up a **subset of containers without the full bring-up** (no set-dress, no seed, no cockpit).
>
> Everything else — stories, UI, set-dress, local content, certs, patches — is toggled by an **environment
> variable**, never a flag. That is why the table below is mostly env knobs.

---

## Env knobs — all 25

`0` = feature ON (the knob is an **opt-out**); `1` = disable it. Read the `DEMO_NO_*` names as *"do not…"*.

### Content & seeding

| Knob | Default | Effect at default | Read at |
|---|---|---|---|
| `DEMO_STORIES` | `1` | **Stories & Heroes seed is ON.** Seeds the **3-org** roster (Cervato Systems / Solvantis / Northwind Aviation), each with a thriving/struggling/manager hero trio | `up-injected.sh:145` |
| `DEMO_NO_STORIES` | `0` | (opt-out of the above) `=1` restores the legacy structural **`small-200`** single-identity seed | `up-injected.sh:146` |
| `DEMO_STORIES_PRESET` | `stack-seeding/presets/stories.seed.yaml` | the preset the stories seed reads | `up-injected.sh:148` |
| `DEMO_NO_SETDRESS` | `0` | **set-dress is ON** — cache-first snapshot replay (taxonomy + directus + sim-embeddings) then the seed | `up-injected.sh:120` |
| `DEMO_NO_LOCAL_CONTENT` | `0` | **per-stack Directus is ON** — content is self-contained. `=1` reads content **live from prod** (the documented fallback) | `up-injected.sh:126` |
| `DEMO_NO_DIRECTUS_DRIFT_FIX` | `0` | the local-content Directus drift fix runs (only when local content is on) | `up-injected.sh:1421` |
| `DEMO_NO_CONTENT_URL_REWRITE` | `0` | the `demo_web` content-URL rewrite runs (only when local content is on) | `up-injected.sh:1458` |

### The UI tier & the presenter cockpit

| Knob | Default | Effect at default | Read at |
|---|---|---|---|
| `DEMO_NO_UI` | `0` | **full UI tier is ON** — next-web + studio-desk (containers) + ant-academy (native) | `up-injected.sh:114` |
| `DEMO_NO_COCKPIT` | `0` | **presenter cockpit is SERVED** (only when `DEMO_STORIES=1`) on `7700 + N·10000`. ⚠️ **A password-free "become any seeded hero" launcher** — see [`../safety.md`](../safety.md) **§3.2** | `up-injected.sh:1483` |

### Demo-patches (platform-source fixes applied to the demo's own ephemeral clone)

See [`demopatch-spec.md`](demopatch-spec.md) for the mechanism and its 7 guards.

| Knob | Default | Effect at default | Read at |
|---|---|---|---|
| `DEMO_NO_PATCH` | `0` | **all demo-patches applied** | `up-injected.sh:469` |
| `DEMO_NO_AUTHZ_SKIP` | `0` | ⚠️ **the `app-targetrole-authz-skip` patch is APPLIED** — authorization is short-circuited on the per-member target-role write path. Part of what makes a demo an **authz-weakened build** ([`../safety.md`](../safety.md) §3.2) | `up-injected.sh:792` |
| `DEMO_NO_AIREADINESS_LOADMEMBERS_BOUND` | `0` | the `app-aireadiness-snapshot-loadmembers` read-path patch is applied | `up-injected.sh:793` |
| `DEMO_NO_PERF_INDEXES` | `0` | the demo perf indexes are created | `up-injected.sh:1380` |
| `DEMO_NO_SENTINEL_RELOAD` | `0` | the sentinel casbin-policy reload runs (the silent-403 catcher) | `up-injected.sh:1350` |

### Remote access

| Knob | Default | Effect at default | Read at |
|---|---|---|---|
| `STACK_PUBLIC_HOST` | `""` (empty) | **no public host.** Browser URLs are `http://localhost:<offset>`. Set to a **dotted MagicDNS FQDN** (or pass `--public-host`) to mint a trusted `tailscale cert` + `tailscale serve` proxies. **A dotless host is hard-refused** — `@clerk/backend`'s `assertValidPublishableKey` rejects it | `up-injected.sh:29` |
| `DEMO_NO_MKCERT` | `0` | the local-trust cert is minted (the localhost path) | `up-injected.sh:132` |

> 🔴 **`STACK_PUBLIC_HOST` does NOT gate network exposure — and never did.** Every demo **container** is
> published on **`0.0.0.0` (all interfaces) on every bring-up**, with or without it. What the knob adds is the
> **trusted HTTPS origin** that makes the already-reachable demo *browsable*. `BIND_HOST` — which *is* derived
> from this knob — gates only the two **host-native** servers (cockpit, ant-academy), never a container.
> Full contract: [`../safety.md`](../safety.md) **Part 3 — the exposure side**.

### Secrets & clones

| Knob | Default | Effect at default | Read at |
|---|---|---|---|
| `DEMO_NO_PROVISION` | `0` | **secrets are provisioned** into the stack's `.env` (values-blind) | `up-injected.sh:52` |
| `DEMO_SECRET_SRC` | `$REPO_ROOT/.agentspace/secrets` | where `/stack-secrets` reads the secret source from | `up-injected.sh:53` |
| `DEMO_NO_SECRET_PREFLIGHT` | `0` | the secret-coverage pre-flight runs (a CRITICAL miss is fatal) | `up-injected.sh:677` |
| `DEMO_ALLOW_UNPINNED_REXT` | `0` | the rext **tag-pin guard ABORTS** if the clone drifts from `.agentspace/rext.tag`. `=1` permits deliberate un-pinned authoring work | `ensure-clones.sh:85` |
| `DEMO_REUSE_DEV_IMAGES` | `0` | **OFF — full independence.** The demo builds from its **own** `stack-demo/` clones; dev's images are never reused | `up-injected.sh:1008` |

### Host pre-flight & verification

| Knob | Default | Effect at default | Read at |
|---|---|---|---|
| `DEMO_NO_HOST_PREFLIGHT` | `0` | the host pre-flight runs (non-fatal — warns, never blocks) | `up-injected.sh:232` |
| `DEMO_VM_MIN_GIB` | `12` | Docker-VM RAM floor asserted by the pre-flight | `up-injected.sh:181` |
| `DEMO_DISK_MIN_GIB` | `20` | free-disk floor asserted by the pre-flight | `up-injected.sh:204` |
| `DEMO_NO_VERIFY` | `0` | **the bring-up auto-verify runs** on the stack's own offset ports (non-fatal). See [`../verification.md`](../verification.md) | `up-injected.sh:1594` |

---

## CLI flags — all 9

| Flag | Entry point | Purpose |
|---|---|---|
| `--public-host` | **`up-injected.sh`** | the dotted MagicDNS FQDN for remote access (env form: `STACK_PUBLIC_HOST`) |
| `--profile` | `rosetta-demo` | compose profile for a low-level `rosetta-demo up` |
| `--services` | `rosetta-demo` | bring up a **subset** of containers (no set-dress / seed / cockpit) |
| `--ref` | `rosetta-demo` | pin a git ref when cloning |
| `--only` | `rosetta-demo` | restrict an operation to named repos |
| `--resolve-only` | `rosetta-demo` | resolve refs without acting |
| `--fapi-host` | `rosetta-demo` | the fake-FAPI host for Clerkenstein injection |
| `--bapi-ip` | `rosetta-demo` | the fake-BAPI IP for Clerkenstein injection |
| `--webhook-secret` | `rosetta-demo` | the Clerkenstein webhook secret |

---

## The shape of the defaults, in one sentence

**A bare `/demo-up N` gives you everything**: the 3-org Stories & Heroes world, the full UI tier, the presenter
cockpit, self-contained content, every demo-patch, and an auto-verify pass — because **every feature knob is an
opt-OUT (`DEMO_NO_*`, default `0`)**. The only knob that is default-*off* is `STACK_PUBLIC_HOST`.

So *"make `/demo-up` pull all the data and seed the 3 orgs"* was **already the default**. The failures people
attribute to a missing default are almost always a **cold snapshot cache** (replay is cache-first and never
captures — see [`../snapshot-cold-start.md`](../snapshot-cold-start.md)), not a knob.

## See also
- [`README.md`](README.md) — the demo-env family index.
- [`../safety.md`](../safety.md) — **Part 3, the exposure side**: what `STACK_PUBLIC_HOST` does and does not gate.
- [`demopatch-spec.md`](demopatch-spec.md) — the demo-patch mechanism the `DEMO_NO_*` patch knobs toggle.
- [`../rosetta_demo.md`](../rosetta_demo.md) — the demo-stack lifecycle.
