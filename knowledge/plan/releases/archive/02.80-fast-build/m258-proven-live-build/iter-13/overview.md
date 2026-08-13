---
iter: 13
milestone: M258
iteration_type: tik
status: closed-fixed
created: 2026-08-12
---

# iter-13 — `TIK-A`: studio-desk, the Class B lever

**Type:** tik · **Active strategy:** `TOK-02` (space partitions by its coupling to time)

## Step 0 — re-survey before targeting

`TOK-02` named studio-desk as `TIK-A`. Re-surveyed at iter-12 and the target is **current and
unabsorbed**: `demo-1-studio-desk:latest` and `demo-2-studio-desk:latest` both measure **1.7 GB**, L1
never touched either, and post-L1 it is the largest UI image on a demo by **4×** (next-web 417 MB).
No substitution needed.

## Cluster / target identified

The single-stage `Dockerfile.dev` in the studio-desk clone: `npm ci` (all 63 deps — 30 prod + 33 dev) →
`COPY . .` → build → ship everything, to run `CMD ["npm","start"]` whose `start` is `node dist/index.js`.
`docker history` attributes **1.04 GB — 61 % of the image — to the `npm ci` layer alone**, producing
63.2 MB of output.

## Hypothesis

An **rext-owned multi-stage Dockerfile** in the shape `next-web.Dockerfile` (M257 L1) and
`hiring.Dockerfile` (M224) already sanction — clone as build CONTEXT only, **zero platform-repo edits** —
that **prunes-and-copies** (`D61`: `npm prune --omit=dev` in the builder, `COPY --from=builder` the
already-populated tree; never a second `npm ci`, which would buy space with time).

## Expected lift

`D62`, deliberately trimmed from iter-11's routing: **≈ 1.25 GB and ≈ 7–10 s**, the time half being the
export/unpack leg only at the measured 5.73–8.05 s/GB — *not* the 115.35 s cold lane, ~105 s of which is
`npm ci` + `tsc` + `vite build` and cannot be removed by a runtime-image shape.

## Phase plan

A. Author the Dockerfile. B. Build a probe image; measure size. C. **Verify it boots and serves**
(a smaller image that 404s is not a win). D. Wire it into `up-injected.sh`. E. Test gate. F. Re-measure.

## Escalation conditions

Route forward if the prune is blocked by the production dependency graph rather than by the Dockerfile
shape. Escalate only if the smaller image cannot serve.

## Acceptable close-no-lift outcomes

A measured refutation that the single-stage shape is *not* what makes this image large would close this
iter honestly — it would redirect `TOK-02`'s Class B away from studio-desk with evidence.
