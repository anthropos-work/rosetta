# db-backup Service

> ## ⚠️ Everything this page said before 2026-08-07 was wrong, and the schedule has been OFF for over a year
>
> Re-derived at `db-backup` **`6e1fb15b`** (= tag `v0.3.3` = repo HEAD), full clone, M257x iter-123. The
> prior text — *"a **Go** service, every **6 hours**, to **S3, Azure, Hetzner**, resilient to a full AWS
> region failure"* — was wrong on **five** independent counts, and the same five were repeated across
> seven corpus files. The measurements:
>
> | Prior claim | Measured |
> |---|---|
> | written in **Go** | **zero `.go` files** (`find . -name '*.go' \| wc -l` → 0). It is a **43-line Bash script** in an Alpine container — `Dockerfile:1` `FROM alpine:latest`, `Dockerfile:19` `ENTRYPOINT ["/backup.sh"]`. Whole repo: 11 files / 572 lines, of which 337 are Terraform |
> | **three** destinations incl. **Azure** | **two** — S3 (`backup.sh:20`) and a Hetzner Storage Box over `scp` (`backup.sh:33`, `u422950.your-storagebox.de:23`). **Azure has never existed**: `git log -S azure --all -i` is empty, and decoding **all 157 objects the repository has ever contained** and grepping `-ci azure` returns **0** |
> | every **6 hours** | **`rate(12 hours)`** (`terraform/main.tf:13`) when it ran. **"6 hours" is not a decayed figure — it is an unsourced one**: it appears nowhere in the repo's working tree or its history (`git log -S '6 hours' --all -i` → empty). Compare the taxonomy figures, which are fenced for exactly this reason |
> | **runs** on a schedule | **The schedule and its trigger are both commented out**, and have been since **`7dd1b80`, 2025-05-29 15:52:42 +0200**, *"refactor: comment out ECS scheduled task resources in main.tf…"* — `terraform/main.tf:10-14` (the `aws_cloudwatch_event_rule` carrying `schedule_expression`) and `:16-27` (the `aws_cloudwatch_event_target` binding it to the ECS task) are commented in their entirety. **Production pins that same commit**: `infrastructure/terraform/production/services.tf:571` sources `db-backup.git//terraform?ref=v0.3.3`, and `v0.3.3` resolves to `6e1fb15b` = HEAD |
> | runs in production **and staging** | `terraform/stage/main.tf` is **0 lines**. There is no staging environment |
>
> **What still exists in production**, because commenting the rule out did not remove them: the ECR repo
> (`main.tf:215`), the ECS **task definition** (`:51`), both IAM roles, the security group, the log group,
> and the S3 bucket with a 30-day expiry (`db-backup/terraform/storage.tf:24-38` — **written repo-qualified
> deliberately**: a bare `storage.tf` is unique across the clone set, so it resolved *silently* to
> `storage/terraform/storage.tf`, a different repository from the one this page is about, and no ambiguity
> guard caught it because there was exactly one hit. Measured M257x iter-202). **Everything needed to run a backup is
> deployed; nothing fires it.**
>
> **This is not "no backups", and the distinction is the whole point.** AWS-native durability on the same
> RDS instance is healthy and independent of this repo —
> `infrastructure/modules/core/storage/rds.tf:19` `backup_retention_period = 7`, `:6` `multi_az = true`,
> `:8-9` `skip_final_snapshot = false` / `delete_automated_backups = false`, plus an AWS Backup plan at
> `:78-89` on **`cron(7 * * * ?)` — hourly** — with `enable_continuous_backup = true` (PITR) into a
> KMS-encrypted vault. **What has been lost is the OFFSITE, NON-AWS leg**, unwritten for over a year.
> The claim *"resilient to a full AWS region failure"* is **RETRACTED**: every surviving copy is in AWS.
>
> **Not measurable from a clone, and therefore not asserted:** whether the EventBridge rule was also
> destroyed from applied AWS state. Commenting a resource out makes Terraform destroy it on the next
> apply, so *"rule gone"* is the expected end state — but AWS state cannot be read from here.
>
> **The operational question — should the offsite leg be turned back on — is the owner's, not this
> corpus's.** This page records what is true; it does not decide what to do about it.

## Role & Responsibility

db-backup `pg_dump`s the platform Postgres, gzips it, and uploads the archive to **two** destinations:

1. **AWS S3** — `backup.sh:20`, `aws s3 cp "/tmp/$BACKUP_FILE.tar.gz" "s3://$S3_BUCKET/$S3_PREFIX/..."`
2. **Hetzner Storage Box** over `scp` — `backup.sh:33`, host at `backup.sh:8-10`

**It is not currently triggered by anything** — see the banner.

## Architecture & Code Map

| Property | Value |
|:---------|:------|
| **Technology** | **Bash** (43 lines) in an Alpine container. **Not Go** — the repo has zero `.go` files. `Dockerfile:4-9` installs `postgresql15-client`, `aws-cli`, `gzip`, `bash`, `openssh-client` |
| **Deployment** | An ECS **task definition** (`terraform/main.tf:51`) with **no live trigger** — the EventBridge rule and target are commented out (`:10-27`) since `7dd1b80` (2025-05-29) |
| **Schedule** | **none, currently.** The commented-out value is `rate(12 hours)` (`terraform/main.tf:13`) |
| **Backup targets** | S3 + Hetzner Storage Box. **No Azure, ever** |
| **Source** | PostgreSQL RDS |
| **Deployed pin** | `v0.3.3` = `6e1fb15b` = HEAD, via `infrastructure/terraform/production/services.tf:571` |

## Local Development

Not in local compose, and not in `repos.yml` — a stack never clones or runs it. `terraform/stage/main.tf`
is empty, so it has no staging deployment either.

## Related Documentation
- [Architecture Overview](../architecture/architecture_overview.md)
- [Security & Compliance](../architecture/security_compliance.md)
- [Org repo register](../architecture/org-repos.md) — where every org repo stands, and its verdict
