package db

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func CreatePool() (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), dbSource)
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
