---
iter: 02
milestone: M45
iteration_type: tik
status: closed-fixed
date: 2026-06-26
---

# iter-02 — tik — the `services/ai/` wrapper (component 1) + the sanctioned `ai` dep

## Active strategy reference
TOK-01 (inside-out fixtures-first build). This is the first tik; it builds component (1) of the 7-step
dependency chain.

## Re-survey
TOK-01 named "build `services/ai/`" as the iter-02 target; the engine does not exist yet, so the target is
current (nothing absorbed it). No substitution.

## Cluster / target identified
The wrapper is the engine's foundation: every later component (cmd/gen-batch, the seeder) needs a
key-blind, cost-tracking, EU-first LLM client. It is also where the SANCTIONED supply-chain inflection
lands (the first new third-party dep). Building it first lets the dep land + be vetted before any
dependent code, and gives a fixture `Completion` the rest of the engine unit-tests against.

## Hypothesis
A thin wrapper over `github.com/anthropos-work/ai@v1.40.1` — EU-first routing (Azure EU → 429 → direct
OpenAI), a model→price cost tracker with the mandatory-ceiling guard, and a fixture `Completion` — can be
fully unit-tested with NO key and NO cost, and the dep add does not break the existing 567-test suite.

## Expected lift
This tik does not move the empirical gate metric (valid-JSON rate needs a prompt + a real call, which
arrive in later tiks). Its lift is INFRASTRUCTURE: the wrapper + the dep + the cost guard land, unit-green,
so later tiks have the LLM client + the cost ceiling they measure against. (A tooling-shaped foundation
tik under TOK-01's inside-out plan.)

## Phase plan
Protocol §4b: build (component 1) → unit-test against fixtures → verify build/vet/gofmt + the full
existing suite still green → license-vet the dep → close.

## Escalation conditions
- The dep add breaks the existing suite → user-blocker. (Did not happen — full suite green.)
- The dep's transitive tree carries a copyleft license → user-blocker / supply-chain decision. (Did not
  happen — all permissive: MIT/Apache/BSD.)

## Acceptable close-no-lift outcomes
N/A — this is a build tik; it closes `closed-fixed` when the wrapper + dep + tests land green.
