## Incident #1 - Migration 001 down error
** Date: 2026-04-09

** What Failed **
When running the migration 001 down it did not drop the checks table. When investigating, it should that the schema_migrations table was empty but checks still existed. 

** Causation **
When originally writing the migration down file I unintentionally overwrote the migration up file, leaving the migration down file completely empty. When I ran ```migrate -path migrations/ -database "postgres://upcheck:password@localhost:5433/upcheck?sslmode=disable" down``` it caused the schema_migrations table to roll back but nothing executed on the checks table.

I learned that the schema_migrations table tracks only the migration command being executed, not the actual checks table being executed on and rolled back. I fixed the migration up file and wrote to the migration down file. Manually removed the checks table and reran the migration up, checked to make sure it was created, then ran migration down to see it get dropped. 