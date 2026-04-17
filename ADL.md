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

**Decision:** What was chosen and why
Decided to use the golang-migrate approach because it would handle the database separately from the application itself, making error handling much easier to diagnose. Also allows for the schema to update as the application progresses.

**Consequences:** What this enables and what it constrains
It enables us to update the table schema as needed without the concern of deprecated instances breaking.

**Revisit:** When/why we might change this
We might change this if we know we only have one instance ever running in production.

## Decision: Project Layout Structure ##
**DATE:** 2026-04-15
**Status:** Accepted

**Context:** What is the best project structure to allow for growth and scale as the project evolves.

**Options Considered:**
1. A single flat file
    Pros: - Zero lift because file was already initally built flat
          - Straightforward for the laymen
    Cons:
          - No separation of concerns
          - Bloated single file
          - Messy as complexity grows rendering the file unreadable
2. A multi-package standard project layout
    Pros: - Clean and organized
          - Package boundaries
          - As complexity grow, the separation of concerns will offer clarity in functionality
    Cons:
          - Took time to set up
          - Could be unneccesary if project is simple enough

**Decision:** What was chosen and why
Decided to build the project using the multi-package approach as it would allow for the application to grow while still being organized and clean. If I decided to wait till it needed it, it could be overwhelmingly time consuming and hard to separate out. While it did take some time, most of that was due to me learning the structure and Go mechanics such as export rules (capital letters) and how the import paths all work together. 

**Consequences:** What this enables and what it constrains
This enables us to build and evolve the project without the need to think about structure and how things could interact. From the start we will build in a manner that minds keeping concerns separated. It does mean there's a front loading of effort to make sure everything builds and interacts with each other.

**Revisit:** When/why we might change this
I don't know why we would change this but the structure may grow as other packages become necessary.


## Decision: Set the PRIMARY KEY on the checks table to be a composite primary key with the id and the checked_at column ##
**DATE:** 2026-04-16
**Status:** Accepted

**Context:** TimescaleDB requires that any unique index on a hypertable includes the partition column 

**Options:** None really, if I decided against setting up a composite primary key, I would have been unable to partition my data. Or at least wouldn't have been able to use TimescaleDB. 

**Decision:** To use TimescaleDB to partition my data into chunks, therefore using checked_at and id as a composite primary key. 

**Consequences:** This enables us to organize our data by date ranges which will allow users to read their data using ranges.

**Revisit:** This would only change if there was a cleaner way of distributing our data.