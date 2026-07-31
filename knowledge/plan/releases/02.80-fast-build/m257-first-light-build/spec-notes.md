# M257 — spec notes

_Technical notes accumulate here during the build._

## Pre-flight audits — iter-01 (bootstrap tok, BLOCKED before strategy authoring)

**`/developer-kit:audit-kb-fidelity --milestone=M257` — verdict `RED`** (2026-07-31).
Report: [`kb-fidelity-audit.md`](kb-fidelity-audit.md) (28 KB, 21 STALE/WRONG + 14 MISSING, 6 blind areas).

Per `/developer-kit:build-mstone-iters` Phase 0b, a RED **blocks the bootstrap tok from authoring
strategy**. No `iter-01/` dir was created — the gate fired before Phase 1. The next invocation's Phase 0
will correctly route iter-01 to the bootstrap tok again, against a re-run audit.

Three of the audit's most load-bearing claims were **independently re-verified** before surfacing the
RED (a gate that stops a session should not rest on a single reading):

| claim | verified how | verdict |
|---|---|---|
| the baseline mirror fence hard-pins `billion.json` **and fences M257's own `overview.md`** | read `stack-core/tests/test_baseline_mirror_fence.py:33,53,83-110` | **CONFIRMED** |
| `verification.md` documents check (d) as a pure win, never mentioning its false-positive mode | read `verification.md:196-206` vs `autoverify.sh:259-269` | **CONFIRMED** |
| odysseus's storage driver is unrecorded and L1's price depends on it | live `docker info` on the host | **CONFIRMED unrecorded — now measured, below** |

---

## Odysseus — first recon (2026-07-31, read-only, `devops@100.110.67.14`)

The host the gate moved to at **D-v28-14**. Measured, not assumed. This closes the audit's blind
area #5 and adds **two findings that were not in the audit and change the campaign plan**.

| | odysseus (measured) | billion (M255 baseline) |
|---|---|---|
| cores | **8** | 8 vCPU |
| RAM | **7,780 MB** | 7.3 GiB (≈7,475 MB) |
| **swap** | **0 — none at all** | **15 GiB** |
| disk free | **189 G of 193 G** | 36–42 GiB comfortable, 25 GiB floor |
| arch / kernel | x86_64 · Linux 6.8.0-117 (Ubuntu) | x86_64 · Linux 6.8.0-134 |
| Docker | 29.6.2 | 29.6.2 |
| **image store** | **containerd** (`io.containerd.snapshotter.v1`, overlayfs snapshotter) | **containerd** |
| Go | **absent** | present |
| atlas | **absent** | present |
| tailscale | 1.98.10 | present |
| git | 2.43.0 | present |

### F1 — odysseus pays the unpack leg. **L1 keeps its full price.** (good news)

`Storage Driver: overlayfs` reads at a glance like the laptop's `overlay2`, and that reading would have
been **expensive and wrong**. The `DriverStatus` is `[["driver-type","io.containerd.snapshotter.v1"]]` —
this is the **containerd image store** using its overlayfs *snapshotter*, the same class as billion, not
the classic overlay2 *graphdriver* the laptop runs.

So `build-budget.md:262`'s *"there is no unpack leg — overlay2, not containerd. Lever L9 is a
`billion`-only phenomenon"* describes **the laptop**, and does **not** generalise to odysseus.
The `unpacking to …` leg is paid here, L9's **85.7 s** is real, and **L1 retains its full ~200–250 s
estimate** rather than losing ~86 s of it.

This was the audit's single highest-value unknown, and it resolves in the gate's favour: had odysseus
been overlay2, L1+L2+L3 would have fallen to ~215–265 s against the 306 s the gate needs, putting a
re-scope signal on the table before the first lever was touched.

### F2 — odysseus's Docker is **completely empty**. The first rep is TRULY COLD, not the gated variant.

```
Images 0 · Containers 0 · Local Volumes 0 · Build Cache 0B
```

The gated variant is **"cold images, warm layer cache"** (`build-budget.md:46-63`), and v2.8
deliberately **cut** the truly-cold run from the gate (D-v28-8) because it is *"a different, slower
thing."* A fresh box has no layer cache at all, so **rep 1 on odysseus cannot be a gated number** — it
measures the variant the release explicitly excluded.

Consequence for the baseline campaign, and it is not optional: **at least one warm-up cycle must run
and be discarded before the n ≥ 3 campaign starts.** M255 already had to budget a warm-up rep for a
*different* reason (the one-off reclaim eviction that cost 173 s, `build-budget.md:336-344`); on
odysseus the reason is stronger — rep 1 is a different measurement class, not merely an unlucky one.
Reporting a p50 that includes it would restate the excluded variant as the gated one.

### F3 — odysseus has **no swap**, and billion's campaign demonstrably used 2,452 MB of it.

Headroom clause 2 arithmetic on odysseus: `1 lane × 3,900 MiB + 1,500 idle = 5,400 MiB` against a
budget of `0.8 × 7,780 = 6,224 MiB`. It **fits** — but billion peaked at **5,446 / 5,579 / 5,398 MB**
with **15 GiB of swap underneath it and 2,452 MB of that swap in use at peak**. Odysseus has none.

The assert will pass on the arithmetic; what is untested is whether a lane that *transiently* exceeds
its measured peak gets a swap cushion (billion) or the OOM killer (odysseus). Treat an OOM-killed
`next build` on odysseus as an expected failure mode with a known cause, not a mystery — and note that
it would present the way M239-F1's ENOSPC did: **as a downstream service dying, not as a memory error.**
`max_parallel_ui_lanes` is **1** here for the same reason it is 1 on billion; L2 must not be read as
licence to run two compile lanes.

### F4 — provisioning gap confirmed, and the prereq list the roadmap points at is real

Go and atlas are both **absent**; Docker, tailscale, and git are present. `roadmap.md:321` claims
`tailscale-serve.md` carries the fresh-Linux-VM prereq list — **it does** (`tailscale-serve.md:119-131`,
a 6-row table with literal install commands, mirrored at `setup_guide.md:110-140`). That pointer is
sound and needs no backfill.

One trap to carry into provisioning: `tailscale-serve.md:133-152` documents **F2b, the login-shell
trap** — `ssh host 'cmd'` runs a non-login shell and reports a false *"Go NOT on PATH"* for a Go that
is installed. This recon's `NO GO` reading was taken that same way. It is corroborated by
`ls -la ~` showing **no Go toolchain and no `.go` dir**, and by atlas being absent too — but the
provisioning step must use the doc's disproof one-liner rather than re-running the same probe shape.
