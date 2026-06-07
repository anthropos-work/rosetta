# Recipe — Multi-month skill-progression demo

**Goal.** Show *growth over time* — members who ran job simulations and skill paths across several months,
with passes and fails, so the workforce-growth and skill-verification views have a real timeline. The "watch
your people level up" demo.

**Time.** ~minutes.

## The idea

The believability of a progression demo is entirely in the **backdated activity**. The seeder's activity
generator (the M7c fleet) emits, per user:
- **job-sim sessions** (`jobsimulation.sessions`) — 1–3 per user, each with a backdated `created_at` spread
  deterministically across the activity span, a `pass/fail` result per the blueprint's `pass_rate`, and a score;
- **skill-path sessions** (`skillpath.skill_path_sessions`) — 0–2 per user, with progress;
- **activity events** (`jobsimulation.activity_events`) — a per-session event trail (started → tasks → ended);
- **assignments** (`public.organization_assignments`) — the admin assigned content to ~half the members.

All of it is time-distributed across `activity.months`, so the growth charts show a curve, not a spike.

## Steps

1. **Bring up + replay the real library (set-dressing) + pick a long activity window.** A longer span makes the
   timeline richer; the snapshot replay makes the sessions link to **real** simulation / skill-path templates
   instead of placeholder ids. Use the `large-1k` preset (9 months) or author your own:
   ```bash
   /demo-up 1
   /stack-snapshot replay 1            # taxonomy + directus — the real catalog + content templates (set-dressing)
   cat > /tmp/progression.seed.yaml <<'YAML'
   stack: demo-1
   org: { name: Stark Industries, slug: stark }
   size: 800
   role_mix: { admin: 0.05, member: 0.75, candidate: 0.20, admin_emails: [founder@stark.test] }
   tier_mix: { free: 0.6, premium: 0.4 }
   content_pack: standard
   activity: { months: 12, pass_rate: 0.62 }   # a full year; ~62% pass so there's a real fail tail
   YAML
   /stack-seed 1 --seed /tmp/progression.seed.yaml
   ```

2. **Confirm the activity landed** (the seed output shows the row counts):
   ```bash
   docker exec demo-1-postgresql-1 psql -U postgres -d postgres -tAc \
     "select 'jobsim_sessions='||count(*) from jobsimulation.sessions;
      select 'skillpath_sessions='||count(*) from skillpath.skill_path_sessions;
      select 'activity_events='||count(*) from jobsimulation.activity_events;
      select 'earliest='||min(created_at)||' latest='||max(created_at) from jobsimulation.sessions;"
   ```
   You should see thousands of sessions/events with `created_at` spread across the full year.

3. **Demo the growth views** — log in as `user_clerkenstein` (see [`recipe-browser-login.md`](recipe-browser-login.md))
   and walk the workforce-growth / skill-verification / top-performers pages; the timeline and funnels are
   populated.

## Notes
- **Deterministic.** The generator uses no random source — a given `stack.seed.yaml` always produces the same
  world, so a re-seed reproduces the exact demo (good for scripted walkthroughs / screenshots).
- **Set-dressed (v1.2).** With the snapshot replay in step 1, the catalog view shows the **real 60K-skill
  taxonomy** and the seeded sessions link to the **real public simulation / skill-path templates** (the v1.2
  "set dressing" — both surfaces are now `snapshot-seeded`, **100%** data-DNA coverage, nothing `waived`). See
  [`recipe-snapshot-world.md`](recipe-snapshot-world.md) for the capture→replay mechanism. If you **skip** the
  replay, the recipe still works — it degrades gracefully to structural-only activity with an empty catalog and
  free (unlinked) content refs.
- **The hard line (still true).** This recipe seeds *structural* activity (sessions/events with real shapes). It
  does **not** fabricate AI transcripts, AI-scored narratives, or fresh embeddings — those (plus external
  shareability) are the **v1.3** "richer demo worlds" theme. The taxonomy + content *libraries* are real (the
  snapshot); the *per-session AI narrative* is not.
- **Tuning.** `activity.months` widens the timeline; `activity.pass_rate` shifts the pass/fail balance; `size`
  scales the population. Schema reference: [`../seeding-spec.md`](../seeding-spec.md).
