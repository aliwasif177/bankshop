package main

import (
	"context"
	"fmt"
	"log"

	"example.com/api"
	db "example.com/db/sqlc"
	"example.com/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		fmt.Println("error==>", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(*store)
	err = server.Start(config.PORT)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
