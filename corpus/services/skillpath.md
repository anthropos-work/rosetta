# Skillpath Service

## Role & Responsibility
*   **Primary Goal**: Manages user session state through skill progression paths, tracking progress as users complete chapters and steps within learning paths.
*   **Key Functions**:
    *   **Session Management**: Creates and manages `SkillPathSession` records that track a user's journey through a skill path
    *   **Progress Tracking**: Records completion status for chapters and individual steps within a path
    *   **Version Upgrades**: Handles migration of user sessions when skill path content is updated
    *   **Event Subscription**: Listens to Jobsimulation events to update step completion when simulations finish

## Architecture & Code Map
*   **Codebase**: `anthropos-dev/skillpath` (repo `git@github.com:anthropos-work/skillpath`)
*   **Language**: Go 1.25
*   **Frameworks**: gqlgen (GraphQL Federation v2), Connect-RPC, Ent ORM, goverter (Ent → domain converters)
*   **Database**: PostgreSQL (`search_path=skillpath`); migrations live in `terraform/migrations/` (Atlas, applied via Terraform/Atlas provider)
*   **Ports**: `8100` (host) → `8080` (container, HTTP/GraphQL); `8101` (host) → `8081` (container, Connect-RPC)
*   **Profile**: `graphql` (default) and `skillpath`
*   **Key Directories**:
    *   `main.go`: Application entry point, server initialization
    *   `rpc.go`: RPC server implementation (`GetSkillPathSession`)
    *   `internal/ent/schema/`: **Ent Entity Definitions** (Source of Truth)
        *   `skillpathsession.go`: Main session entity
        *   `chaptersession.go`: Chapter progress entity
        *   `stepsession.go`: Step progress entity
    *   `internal/graph/`: **GraphQL Implementation** (gqlgen)
        *   `schemas/schema.graphqls`: Type definitions
        *   `schemas/queries.graphqls`: Query operations
        *   `schemas/mutations.graphqls`: Mutation operations
    *   `internal/session/`: Session manager business logic + pub/sub handlers

## Data Model

The service uses a hierarchical session model to track user progress:

```mermaid
erDiagram
    SkillPathSession ||--o{ ChapterSession : contains
    ChapterSession ||--o{ StepSession : contains

    SkillPathSession {
        uuid id PK
        uuid user_id
        uuid skillpath_id "CMS SkillPath reference"
        int progress "0-100"
        enum status "pending|active|completed|archived"
        string version
        timestamp started_at
        timestamp ended_at
    }

    ChapterSession {
        uuid id PK
        uuid user_id
        uuid chapter_id "CMS Chapter reference"
        int progress
        int duration "seconds"
        enum status
    }

    StepSession {
        uuid id PK
        uuid user_id
        uuid step_id "CMS Step reference"
        uuid last_simulation_session "JobSim session reference"
        int progress
        int duration
        enum status
    }
```

**Status Enum** (`internal/ent/enum/status.go`):
- `pending`: Session created but not started
- `active`: User currently working through content
- `completed`: All steps finished
- `archived`: Old version, superseded by upgrade

## Interface Discovery

### GraphQL API

Access the **GraphQL Playground** at `http://localhost:8080/` when running locally.

**Queries** (`internal/graph/schemas/queries.graphqls`):
```graphql
type Query {
  # Get or create a session for a user on a skill path
  getOrCreateSkillPathSession(userId: ID!, skillPathId: ID!, version: String): SkillPathSession!

  # List all active (in-progress) sessions for a user
  skillPathActiveSessions(userId: ID!): [SkillPathSession!]!

  # List all completed sessions for a user
  skillPathCompletedSessions(userId: ID!): [SkillPathSession!]!
}
```

**Mutations** (`internal/graph/schemas/mutations.graphqls`):
```graphql
type Mutation {
  # Mark a step as completed
  completeSkillPathStep(userId: ID!, skillPathSessionId: ID!, stepId: ID!): Boolean!

  # Revert step completion
  uncompleteSkillPathStep(userId: ID!, skillPathSessionId: ID!, stepId: ID!): Boolean!

  # Upgrade user's session to latest skill path version
  upgradeSkillPathSessionToLatest(userId: ID!, skillPathId: ID!): SkillPathSession!

  # Bulk upgrade all users on a skill path (admin)
  upgradeAllSkillPathSessionsToLatest(skillPathId: ID!, userId: ID!): UpgradeAllSkillPathSessionsResult!
}
```

### RPC API

**Service**: `SkillPathSessionService` (Connect RPC)

**Operations** (`rpc.go`):
```go
// GetSkillPathSession retrieves a session with chapters and steps
GetSkillPathSession(ctx, req *skillpathv1.GetSkillPathSessionRequest) (*skillpathv1.GetSkillPathSessionResponse, error)
```

### Dependencies

*   **Upstream Consumers**:
    *   Cosmo Router (federated queries — skillpath is one of the 5 subgraphs)
    *   Backend (`app`) — depends on skillpath at compose startup
*   **Downstream Dependencies**:
    *   **Sentinel** (Connect-RPC): authorization (manager + admin checks)
    *   **CMS** (Connect-RPC): skill path content structure on session creation (`CMS_RPC_ADDR=http://cms:8091`)
    *   **Jobsimulation** (Redis Streams, consumed): simulation step completion events
    *   **PostgreSQL** (`skillpath` schema), **Redis** (Watermill subscriber)

> Note: skillpath consumes the jobsimulation event stream rather than calling jobsimulation via RPC.

## Event-Driven Architecture

Skillpath subscribes to the **Jobsimulation Redis Stream** to react when users complete simulations:

```mermaid
sequenceDiagram
    participant User
    participant JobSim as Jobsimulation
    participant Redis as Redis Stream
    participant Skillpath
    participant DB as PostgreSQL

    User->>JobSim: Complete simulation
    JobSim->>Redis: Publish simulation_completed event
    Redis->>Skillpath: Deliver event (subscriber)
    Skillpath->>DB: Update StepSession status
    Skillpath->>DB: Recalculate chapter/path progress
```

**Subscription Handler** (`internal/session/session.go`):
- Listens to `JOBSIMULATION_STREAM` environment variable
- Processes events via `SessionManager.JobSimulationSubscriber()`
- Updates step completion status based on simulation results

## Local Development

### 1. Running Standalone
*   **Prerequisites**:
    *   PostgreSQL running with `skillpath` schema
    *   Redis available for pub/sub
    *   CMS and Jobsimulation services running (for RPC calls)
    *   Environment variables: `DB_CONNECTION`, `REDIS_ADDR`, `CMS_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR`, `CLERK_SECRET_KEY`
*   **Setup**:
    ```bash
    cd anthropos-dev/skillpath
    make setup    # Install tools (ent, atlas)
    make gen      # Generate Ent code
    atlas migrate apply --env local  # Apply migrations
    ```
*   **Run**:
    ```bash
    go run .
    ```

### 2. Running in Docker
*   **Service Name**: `skillpath`
*   **Command**:
    ```bash
    cd platform
    docker compose up -d skillpath
    ```

### 3. Testing
```bash
cd anthropos-dev/skillpath
go test ./...

# With coverage
go test -cover ./internal/session/...
```

## Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `PORT` | Container HTTP/GraphQL port | `8100` |
| `RPC_PORT` | Container Connect-RPC port | `8101` |
| `DB_CONNECTION` | PostgreSQL connection string | `postgres://...?search_path=skillpath` |
| `REDIS_ADDR` | Redis server address | `redis:6379` |
| `REDIS_STREAMS_INDEX` | Redis DB index for streams | `4` |
| `CMS_RPC_ADDR` | CMS service RPC address | `http://cms:8091` |
| `JOBSIMULATION_STREAM` | Redis stream to subscribe | `jobsimulation` |
| `AUTHORIZATION_ADDRESS` | Sentinel service address | `http://sentinel:8087` |
