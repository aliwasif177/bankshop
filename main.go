package main

import (
	"context"
	"fmt"
	"log"

	"example.com/api"
	db "example.com/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	port     = "0.0.0.0:8080"
)

func main() {

	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		fmt.Println("error==>", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(*store)
	err = server.Start(port)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
