package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/aliwasif177/bankshop/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries

func CreatePool() (*pgxpool.Pool, error) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	return pgxpool.New(context.Background(), config.DBSource)
}

func TestMain(m *testing.M) {
	conn, err := CreatePool()
	if err != nil {
		panic("Cannot connect to db: " + err.Error())
	}

	testQueries = New(conn)

	code := m.Run()

	conn.Close()

	os.Exit(code)
}
