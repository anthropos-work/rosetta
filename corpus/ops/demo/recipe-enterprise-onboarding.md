# Recipe — Enterprise org onboarding demo

**Goal.** A believable enterprise customer org: an admin who logs in and sees a populated workforce — hundreds
of members with roles, tiers, and months of activity — that they can click through. The "this is what your
org looks like in Anthropos" demo.

**Time.** ~minutes (first stack bring-up is longer if images aren't cached; seeding is ~1s).

## Steps

1. **Bring up an isolated stack** (Clerkenstein-wired, offset ports):
   ```bash
   /demo-up 1            # or: stack-demo/rosetta-extensions/demo-stack/up-injected.sh 1
   ```
   This brings up `demo-1` on offset ports (`+10000`), Clerk-free, with its own data — the dev stack
   untouched. `migrate-demo.sh` runs automatically and bootstraps the global Sentinel policy (required for
   authorized routes to return 200). See [`../rosetta_demo.md`](../rosetta_demo.md).

2. **Set-dress the stack** — replay the real public library so the catalog + content templates are real:
   ```bash
   /demo-snapshot replay 1            # taxonomy + directus — the real 60K-skill catalog + content templates
   ```
   This stamps the real **public** taxonomy + Directus content into the stack (almost always a cache-hit → zero
   prod read). It's **optional** (skip for a quick structural-only world — the seeder degrades gracefully), but
   it's what turns "an org with users" into "an org with users browsing the real product catalog". See
   [`recipe-snapshot-world.md`](recipe-snapshot-world.md).

3. **Seed the org** with a curated preset:
   ```bash
   /demo-seed 1 --preset mid-500     # Globex Corp, 500 members, 6 months activity (the default)
   # or --preset large-1k for a 1,000-member scale demo
   ```
   The seeder backfills, in dependency order: the org → 500 users → memberships (with the role mix) → the real
   **`user_clerkenstein`** admin login identity + the casbin grant → backdated job-sim / skill-path sessions →
   assignments → activity events. It prints per-surface row counts + `isolation: clean`.

4. **(Optional) verify** the world is coherent:
   ```bash
   SS=stack-demo/rosetta-extensions/stack-seeding
   /tmp/datadna measure --stack demo-1 --dna "$SS/dna/data-dna.json"   # conformance 100% over reachable surfaces
   /tmp/datadna catalog --dna "$SS/dna/data-dna.json"                  # what's seeded / snapshot-seeded (100%, nothing waived)
   # with snapshots replayed, also gate fidelity (captured source vs replayed stack):
   SNAP=.agentspace/snapshots
   /tmp/datadna measure-snapshot --stack demo-1 --dna "$SS/dna/data-dna.json" \
     --manifest "$SNAP/taxonomy/<ver>/manifest.json" --manifest "$SNAP/directus/<ver>/manifest.json"
   ```

5. **Log in + demo.** Follow [`recipe-browser-login.md`](recipe-browser-login.md): the browser logs in as
   `user_clerkenstein` (no real Clerk), lands in the seeded org as **admin**, and authorized routes return
   **200** — the workforce list, member detail, growth/insights, etc., all populated.

6. **Tear down** when done:
   ```bash
   /demo-down 1 --purge     # removes demo-1's containers/network/data; dev stack untouched
   ```

## What makes it believable
- **Role + tier mix** — admins / members / candidates and free / premium tiers per the preset's ratios, so the
  member list isn't monotone.
- **Months of backdated activity** — sessions spread across the activity span (not all "just now"), with a
  realistic pass/fail mix, so growth charts and timelines have shape.
- **A real admin identity** — `user_clerkenstein` is a *seeded* member (not a phantom token), so org-gated
  pages resolve to real data.
- **The real product library (set-dressing, step 2)** — the catalog shows the real 60K-skill taxonomy and the
  seeded sessions/assignments link to real public simulation / skill-path templates, so the "browse the catalog"
  and "assigned content" surfaces aren't empty or placeholder.

## Tuning the org
Copy a preset and edit it (it's a `stack.seed.yaml`): `org.name`, `size`, `role_mix`/`tier_mix`,
`activity.months` / `activity.pass_rate`. Re-seed with `/demo-seed 1 --seed my.seed.yaml` (run `--reset` first
to clear the prior data). Schema reference: [`../seeding-spec.md`](../seeding-spec.md).
