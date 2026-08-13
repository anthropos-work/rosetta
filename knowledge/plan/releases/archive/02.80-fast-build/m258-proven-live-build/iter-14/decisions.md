# M258 iter-14 — decisions

## D69 — `-v` on the plain teardown is safe HERE, and the safety was measured.

The objection to `down -v` on the non-purge path is real: `-v` removes **named** volumes, which would
turn "keep the data, fast re-up" into a destructive op. It does not apply here, and that was established
by census rather than by reading the compose:

```
every container in the demo project, volume-type mounts only:
  demo-1-postgresql-1  VOLUME 9f227c2e92bc…  /docker-entrypoint-initdb.d      ANONYMOUS
  demo-1-postgresql-1  VOLUME 1c5dd2836cb6…  /docker-entrypoint-preinitdb.d   ANONYMOUS
  (no other container mounts a volume at all)
```

**There are no named volumes in a demo stack to lose**, and the real database is a **host bind mount**
(`stacks/demo-N/data/postgresql`) that `down -v` never touches. So the plain teardown keeps exactly what
it always kept and stops leaking exactly what it always leaked.

⚠️ This is a **conditional** safety property, not a permanent one. If a named volume is ever added to the
demo compose, the decision must be re-taken — which is why
`test_down_plain_removes_anonymous_volumes` asserts the *rationale comment*, not only the flag.

## D70 — The F-9 stack-dir residue is 276 MB per stack, and it is NOT being fixed in this iter.

`purge_data_dir` is scoped to `$stack/data` alone (`rosetta-demo:264`, with a G1 path-assert pinning it
to exactly `$STACKS_DIR/demo-$n/data`). Everything else in the stack dir survives `--purge`:

| survives a full `--purge` | size (demo-1) |
|---|---|
| `clones/` | **220 MB** |
| `bin/` | 37 MB |
| `fake-fapi` + `fake-bapi` | 18.5 MB |
| logs + manifests | <1 MB |
| **total** | **≈ 276 MB per stack** |

Real, Class A, and worth taking — but **not minutes before `END-M258-one-stack`**. Widening a
`rm -rf` whose current safety rests on a path-assert, immediately before the milestone's binding end
state, trades a 276 MB win against the deliverable the user actually asked for. `TIK-C` tears a stack
down anyway and will **measure** the residue empirically; routed with its price attached.
