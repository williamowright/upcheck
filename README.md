# What is UpCheck
UpCheck is an application that sends out http requests to urls of choice and 
stores the response in a postgres database. It sends out the requests intermittently at a fixed interval.

## Stack
Go
Postgres
Docker


## How to Run
Insert table connection values (Username, Password, Database, Port) to environment variables like so:
  POSTGRES_USER=[PLACEHOLDER]
  POSTGRES_PASSWORD=[PLACEHOLDER]
  POSTGRES_DB=[PLACEHOLDER]
  POSTGRES_PORT=[PLACEHOLDER]
Initialize a Postgres database table:
  Run ```docker-compose up -d```
  Run ```docker exec -it uptimechecker-postgres-1 psql -U upcheck -d upcheck```
  Run ```CREATE TABLE  checks (   url TEXT,   status_code INT,   response_time INT,   is_up BOOL );```
Run ```go run upcheck.go```

## Common Pitfalls
- If having connection issues to database, check environment variables and run ```docker-compose down -v```
- If you have run ```docker-compose down -v``` all the volumes are dropped so you must recreate the table named checks