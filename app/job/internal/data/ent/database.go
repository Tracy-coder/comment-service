// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
)

// Database is the client that holds all ent builders.
type Database struct {
	client *Client
}

// NewDatabase creates a new database configured with the given options.
func NewDatabase(opts ...Option) *Database {
	return &Database{client: NewClient(opts...)}
}

// InTx runs the given function f within a transaction.
func (db *Database) InTx(ctx context.Context, f func(context.Context) error) error {
	tx := TxFromContext(ctx)
	if tx != nil {
		return f(ctx)
	}

	tx, err := db.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	if err = f(NewTxContext(ctx, tx)); err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			return fmt.Errorf("rolling back transaction: %v (original error: %w)", err2, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

func (db *Database) loadClient(ctx context.Context) *Client {
	tx := TxFromContext(ctx)
	if tx != nil {
		return tx.Client()
	}
	return db.client
}

// Exec executes a query that doesn't return rows. For example, in SQL, INSERT or UPDATE.
func (db *Database) Exec(ctx context.Context, query string, args ...interface{}) (*sql.Result, error) {
	var res sql.Result
	err := db.loadClient(ctx).driver.Exec(ctx, query, args, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Query executes a query that returns rows, typically a SELECT in SQL.
func (db *Database) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	var rows sql.Rows
	err := db.loadClient(ctx).driver.Query(ctx, query, args, &rows)
	if err != nil {
		return nil, err
	}
	return &rows, nil
}

// CommentContent is the client for interacting with the CommentContent builders.
func (db *Database) CommentContent(ctx context.Context) *CommentContentClient {
	return db.loadClient(ctx).CommentContent
}

// CommentIndex is the client for interacting with the CommentIndex builders.
func (db *Database) CommentIndex(ctx context.Context) *CommentIndexClient {
	return db.loadClient(ctx).CommentIndex
}

// CommentSubject is the client for interacting with the CommentSubject builders.
func (db *Database) CommentSubject(ctx context.Context) *CommentSubjectClient {
	return db.loadClient(ctx).CommentSubject
}
