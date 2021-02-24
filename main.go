package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/jackc/pgx/v4"
)

var (
	conn *pgx.Conn
	wg   sync.WaitGroup
)

func main() {

	count, err := strconv.Atoi(os.Getenv("CONCURRENT_CONNECTIONS"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing CONCURRENT_CONNECTIONS: %v", err)
		os.Exit(1)
	}

	wg.Add(count)

	for i := 0; i < count; i++ {
		go torturePostgresql(i)
	}
	wg.Wait()
}

func torturePostgresql(index int) {
	defer wg.Done()
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	for i := 0; i < 1000; i++ {
		rows, err := conn.Query(context.Background(), "select typname from pg_type")
		if err != nil {
			log.Fatalf("%v", err)
		}
		for rows.Next() {
			var description string
			err = rows.Scan(&description)
			if err != nil {
				log.Fatalf("%v", err)
			}
		}
	}
}
