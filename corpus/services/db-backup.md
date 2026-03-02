# db-backup Service

## Role & Responsibility

db-backup performs **scheduled PostgreSQL backups** every 6 hours, storing dumps in three geographically separate locations for disaster recovery:

1. **AWS S3** (primary)
2. **Azure Blob Storage** (secondary)
3. **Hetzner** (Germany, tertiary)

This ensures data resilience even in the event of a full AWS region failure.

## Architecture & Code Map

| Property | Value |
|:---------|:------|
| **Technology** | Go |
| **Deployment** | Docker (scheduled, internal) |
| **Schedule** | Every 6 hours |
| **Backup Targets** | S3, Azure, Hetzner |
| **Source** | PostgreSQL RDS (all schemas) |

The service runs on a timer, performing full database dumps and uploading them to all three storage providers. Combined with RDS point-in-time recovery, this provides comprehensive backup coverage.

## Local Development

db-backup runs in production and staging environments. It is not typically needed in local development.

## Related Documentation
- [Architecture Overview](../architecture/architecture_overview.md)
- [Security & Compliance](../architecture/security_compliance.md)
