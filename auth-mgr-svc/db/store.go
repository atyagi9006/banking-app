package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/atyagi9006/banking-app/auth-mgr-svc/pkg/config"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	//dbSource = "postgresql://root:P@ssw0rd@localhost:5432/my_bank?sslmode=disable"
	//postgresql://%s:%s@%s:%s/?sslmode=disable

)

type Store struct {
	db *sql.DB
	*Queries
}

func NewStore(c *config.SQLConfig) *Store {
	conn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		c.Host, c.Port, c.DBName, c.User, c.Pass, c.SSLMode,
	)
	db, err := sql.Open(dbDriver, conn)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	return &Store{
		db:      db,
		Queries: New(db),
	}

}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return nil
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
