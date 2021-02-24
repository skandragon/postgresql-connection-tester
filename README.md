# postgresqm-connection-tester

This is a quick hack to abuse a postgresql database.  It will connect to the URI
specified in the envar `DATABASE_URL` (eg: `postgresql://127.0.0.1/template1) and
issue a pile of queries.  The `template1` database,or any database with a
table called `pg_type` with a `typname string` column, will work.  No writes are
made.

The envar `CONCURRENT_CONNECTIONS` specifies the number of concurrent database
connections to attempt.  There is no syncronization, so it's possible some will
finish prior to others completing.  In my test, 100 (with a limit of 100 in the
server) properly works.  101 fails with an error.  500 fails with a different
error, perhaps related to the OS limiting TCP connection rate.

# Example

```
CONCURRENT_CONNECTIONS=101 DATABASE_URL="postgresql://127.0.0.1/template1" go run main.go
```
