package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/config"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase() *Database {
	db, err := sqlx.Connect("postgres", config.DBUrl)
	if err != nil {
		log.Fatalln(err)
	}

	return &Database{
		db: db,
	}
}
