# M212 — progress

Section checklist (closure = all boxes land + a dry `up-injected.sh` run with `STACK_PUBLIC_HOST` set bakes the
MagicDNS host into every browser-facing value; unset ⇒ byte-identical to today).

- [ ] `HOST` var + `STACK_PUBLIC_HOST` default in `up-injected.sh`
- [ ] `/demo-up --public-host` flag → `STACK_PUBLIC_HOST` plumbed to scripts
- [ ] next-web build-args + `.env.local` overlays substituted
- [ ] studio-desk build-args substituted
- [ ] `inject.py --fapi-host "$HOST:..."` (pk mint)
- [ ] `gen_injected_override.py` host-param plumbing (emission → M214)
- [ ] cockpit `--host 0.0.0.0` + host into `--app-base`/`--fapi-host`/`--academy-base`
- [ ] `ant-academy.sh` host sub + confirm `next dev` bind
- [ ] `demo_web` Directus rewrite substituted
- [ ] `want_ep` cache-validators invalidate on HOST change
- [ ] `stack_registry.py` additive `external_host`
- [ ] unset-knob regression check (byte-identical) + tests
