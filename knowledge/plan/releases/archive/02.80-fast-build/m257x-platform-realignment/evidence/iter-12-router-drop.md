# iter-12 evidence — the router drop, measured live

Captured 2026-08-01T10:46:44Z on the dev host.

## 1. Platform origin HEAD moved mid-milestone

```
2adcf71 2026-07-31 15:58:34 +0200 Merge pull request #23 from anthropos-work/chore/drop-wundergraph
360efd4 2026-07-31 13:47:02 +0000 chore(compose): remove the router service and clone entry outright
b56d731 2026-07-31 13:37:39 +0000 chore(compose): drop the WunderGraph router; point local dev at backend
```

`repos.yml` @ 2adcf71 — the `graphql-wundergraph` entry is deleted:

```
```

## 2. The generator does not notice (RC=0)

```
gen_injected_override: demo-1 -> .agentspace/scratch/work-m257x/negctl/override.yml (3 injected, 2 frontends, directus, single-identity fake-fapi+fake-bapi added)
```

It still emits, against a platform that has no such service:

```
86:    profiles: !override [graphql]
99:      - WUNDERGRAPH_SSR_ENDPOINT=http://graphql:8080/graphql
102:    profiles: !override [graphql]
123:    profiles: [graphql]
134:      - WUNDERGRAPH_SSR_ENDPOINT=http://graphql:8080/graphql
140:      graphql:
```

## 3. compose rejects the project outright (RC=1)

```
service "hiring-app" depends on undefined service "graphql": invalid compose project
```
