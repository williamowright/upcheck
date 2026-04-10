## Decision: Database Migration
**Date:** 2026-04-09
**Status:** Accepted

**Context:** How do we ensure that Databases get updated in sync across multiple instances without manual intervention and instance drift

**Options Considered:**
1. If table doesn't exist logic
  Pros: - Straightforward
        - Readable
  Cons: - Deprecated Instances can cause misalignment between instances
        - If multiple instances boot up and one instance is manipulating the table  it could cause intermittent errors that do not reproduce
2. golang-migrate 
  Pros: - Handles instance drift
        - Updates tables as they change
        - alignment across team
        - separate the concerns of the database and the application itself
  Cons: - Have to maintain
        - Changes need to not create conflicts

**Decision:** What we chose and why
Decided to use the golang-migrate approach because it would handle the database separately from the application itself, making error handling much easier to diagnose. Also allows for the schema to update as the application progresses.

**Consequences:** What this enables and what it constrains
It enables us to update the table schema as needed without the concern of deprecated instances breaking.

**Revisit:** When/why we might change this
We might change this if we know we only have one instance ever running in production