# iter-01 — decisions (local)

**D1 — Active-cycle signals-true over closed-cycle snapshot-direct.**
- Context: the contract gives two seed strategies depending on the AI-readiness cycle state (active ⇒ dashboard
  recomputes from signals; closed ⇒ reads frozen snapshots directly).
- Options: (a) active + signals-true (write evidences + ended jobsim sessions + step-progress); (b) closed +
  snapshot-direct (write frozen `ai_readiness_snapshots` rows, flip cycle to closed, no underlying signals).
- Choice: (a). Full rationale in milestone-root `decisions.md` TOK-01. Summary: the gate needs a *live, in-flight*
  assessment with a hero STARTED (a started hero only exists mid-cycle — a closed cycle is a finished assessment);
  active-cycle reuses the existing evidence/jobsim machinery; the contract (verified GREEN) confirms the active path
  recomputes from `user_skill_evidences` + jobsim sessions, so seeding signals renders authentically and survives a
  `RefreshLiveSnapshots`.
- Why not (b): snapshot-direct is lighter but reads as finished, can't show a STARTED hero, and seeding
  `live_snapshots` for an active cycle silently no-ops (overwritten on refresh — the trap claim-5 guards against).
