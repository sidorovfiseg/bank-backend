package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PoolConnection struct {
	pool *pgxpool.Pool
}

type transaction struct {

}

func NewPoolConnection(pool *pgxpool.Pool) *PoolConnection {
	return &PoolConnection{
		pool: pool,
	}
}

func (connection *PoolConnection) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	if tx, ok := ctx.Value(transaction{}).(pgx.Tx); ok {
		return tx.Exec(ctx, sql, args...)
	}
	return connection.pool.Exec(ctx, sql, args...)
}

func (connection *PoolConnection) Query(ctx context.Context, sql string, args ...any) (pgx.Row, error) {
	if tx, ok := ctx.Value(transaction{}).(pgx.Tx); ok {
		return tx.Query(ctx, sql, args...)
	}
	return connection.pool.Query(ctx, sql, args...)
}

func (connection *PoolConnection) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	if tx, ok := ctx.Value(transaction{}).(pgx.Tx); ok {
		return tx.QueryRow(ctx, sql, args...)
	}
	return connection.pool.QueryRow(ctx, sql, args...)
}

type PoolTransactionManager struct {
	poolConnection PoolConnection
}


func NewPoolTransactionManager(pool *pgxpool.Pool) *PoolTransactionManager {
	return &PoolTransactionManager{
		poolConnection:  *NewPoolConnection(pool),
	}
}

func (ptm *PoolTransactionManager) Do(ctx context.Context, f func(ctx context.Context) error) error {
	tx, err := ptm.poolConnection.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	txContext := context.WithValue(ctx, transaction{}, tx)
	
	if err = f(txContext); err != nil {
		err := tx.Rollback(ctx)
		if err != nil {
			return err
		}
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		err = tx.Rollback(ctx)
		if err != nil {
			return err
		}
		
		return err
	}

	return err
	


}