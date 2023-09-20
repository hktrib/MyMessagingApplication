package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	DB *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		DB:      db,
		Queries: New(db),
	}
}

// function execTx: Executes a function database within database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("Transaction (TX) Error: %v, Rollback (RB) Error: %v", err, rbErr)
		}
	}
	return tx.Commit()
}
