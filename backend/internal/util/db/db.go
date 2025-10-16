package db

import (
	"context"
	"database/sql"

	"cuhara.qua.go/internal/util"
	"github.com/aarondl/sqlboiler/v4/boil"
)

type TxFn func(boil.ContextExecutor) error

func WithTransaction(ctx context.Context, db *sql.DB, fn TxFn) error {
	return WithConfiguredTransaction(ctx, db, nil, fn)
}

func WithConfiguredTransaction(ctx context.Context, db *sql.DB, options *sql.TxOptions, fn TxFn) error {
	tx, err := db.BeginTx(ctx, options)
	if err != nil {
		util.LogFromContext(ctx).Warn().Err(err).Msg("Failed to start transaction")
	}

	defer func() {
		if p := recover(); p != nil {
			util.LogFromContext(ctx).Error().Interface("p", p).Msg("Recovered from panic, rolling back transaction and panicking again")

			if txErr := tx.Rollback(); txErr != nil {
				util.LogFromContext(ctx).Warn().Err(txErr).Msg("Failed to roll back transaction after recovering from panic")
			}

			panic(p)
		} else if err != nil {
			util.LogFromContext(ctx).Warn().Err(err).Msg("Received error, rolling back transaction")

			if txErr := tx.Rollback(); txErr != nil {
				util.LogFromContext(ctx).Warn().Err(txErr).Msg("Failed to roll back transaction after receiving error")
			}
		} else {
			err = tx.Commit()
			if err != nil {
				util.LogFromContext(ctx).Warn().Err(err).Msg("Failed to commit transaction")
			}
		}
	}()

	err = fn(tx)

	return err
}
