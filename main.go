package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aliwasif177/bankshop/api"
	db "github.com/aliwasif177/bankshop/db/sqlc"
	"github.com/aliwasif177/bankshop/util"
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
	config, err = util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(config.PORT)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
