## Incident #1 - Migration 001 down error
** Date: 2026-04-09 **

** What Failed **
When running the migration 001 down it did not drop the checks table. When investigating, it showed that the schema_migrations table was empty but checks still existed. 

** Causation **
When originally writing the migration down file I unintentionally overwrote the migration up file, leaving the migration down file completely empty. When I ran ```migrate -path migrations/ -database "postgres://upcheck:password@localhost:5433/upcheck?sslmode=disable" down``` it caused the schema_migrations table to roll back but nothing executed on the checks table.

I learned that the schema_migrations table tracks only the migration command being executed, not the actual checks table being executed on and rolled back. I fixed the migration up file and wrote to the migration down file. Manually removed the checks table and reran the migration up, checked to make sure it was created, then ran migration down to see it get dropped. 


## Incident #2 - ENV File Variable Misname
** Date: 2026-04-13 **

** What Failed **
```panic: failed to open database: pq: role "upcheck" does not exist (28000) 
   goroutine 1 [running]: 
   main.main() /[PATH]/uptimechecker/upcheck.go:67 +0x471 
   exit status 2```
The database location could not be found. 

** Causation **
This occured after I accidentally added a character to the POSTGRES_PORT variable causing the entire databse not to be found. I ran cat .env to see the file and noticed that the variable name had been altered. I fixed the variable name and reran the application. It ran successfully after this fix.

I have to be careful about accidental key-ing inside files. Luckily this was a quick fix that errored out in the terminal, if it had been a silent failure or a accidental key-in inside another file, the troubleshooting could have taken much longer.



## Incident #3 - Github Oct 21 2018 Postmortem read
** Date: 2026-04-27 **

** What Failed **
Multiple Github services were down for almost 24 hours. 


** Causation **
Orchestration promoted database primaries across regional boundaries, which is the correct behavior but the application tier was unable to support the topology change. Pair this with latency issues and you have misalignment between east and west coast servers, called Split Brain.  This allowed for two sources of truth with no clean way to merge them back together,  causing the resolution of the issue to take much longer because the priority was to maintain the safety of user data.


With my app, I don't have to worry about split brain risk yet because our Postgres instane is a single point of failure right now, for which will get address later on. From reading this I learned the risk of automated systems having triggers but not guard rails. It was wrong for the data to be allowed to misalign in such a catastrophic manner. The other thing I learned is that in  a production environment, you must not be hasty to resolve an issue without determining if the solution will damage user data and in turn user trust.