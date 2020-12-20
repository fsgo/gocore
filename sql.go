/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/12/20
 */

package gocore

import (
	"context"
	"database/sql"
	"reflect"
)

var _ SQLDB = (*sql.DB)(nil)

// SQLDBContexter interface of *Sql.DB
type SQLDBContexter interface {
	PingContext(ctx context.Context) error
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

// SQLPingContext Ping with Context
type SQLPingContext interface {
	PingContext(ctx context.Context) error
}

// PingContext Ping with Context
type PingContext = SQLPingContext

// SQLPrepareContext Prepare with Context
type SQLPrepareContext interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

// SQLExecContext Exec with Context
type SQLExecContext interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

// SQLQueryContext Query with Context
type SQLQueryContext interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

// SQLQueryRowContext  QueryRow with Context
type SQLQueryRowContext interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// SQLBeginTx  BeginTx
type SQLBeginTx interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

// SQLDBNoContexter interface of *Sql.DB
type SQLDBNoContexter interface {
	Ping() error
	Prepare(query string) (*sql.Stmt, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Begin() (*sql.Tx, error)
}

// SQLPing  ping
type SQLPing interface {
	Ping() error
}

// Ping ping
type Ping = SQLPing

// SQLPrepare Prepare
type SQLPrepare interface {
	Prepare(query string) (*sql.Stmt, error)
}

// SQLExec exec
type SQLExec interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// SQLQuery query
type SQLQuery interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

// SQLQueryRow QueryRow
type SQLQueryRow interface {
	QueryRow(query string, args ...interface{}) *sql.Row
}

// SQLBegin begin tx
type SQLBegin interface {
	Begin() (*sql.Tx, error)
}

// SQLStmt interface for *sql.Stmt
type SQLStmt interface {
	Close() error
	Exec(args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error)
	Query(args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, args ...interface{}) (*sql.Rows, error)
	QueryRow(args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row
}

var _ SQLStmt = (*sql.Stmt)(nil)

// SQLRows interface for *sql.Rows
type SQLRows interface {
	Close() error
	ColumnTypes() ([]*sql.ColumnType, error)
	Columns() ([]string, error)
	Err() error
	Next() bool
	NextResultSet() bool
	Scan(dest ...interface{}) error
}

var _ SQLRows = (*sql.Rows)(nil)

// SQLColumnType interface of *sql.ColumnType
type SQLColumnType interface {
	DatabaseTypeName() string
	DecimalSize() (precision, scale int64, ok bool)
	Length() (length int64, ok bool)
	Name() string
	Nullable() (nullable, ok bool)
	ScanType() reflect.Type
}

var _ SQLColumnType = (*sql.ColumnType)(nil)

// SQLConn interface *sql.Conn
type SQLConn interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PingContext(ctx context.Context) error
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Raw(f func(driverConn interface{}) error) (err error)
	Close() error
}

var _ SQLConn = (*sql.Conn)(nil)

// SQLTx interface for *sql.Tx
type SQLTx interface {
	Commit() error
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Rollback() error
	Stmt(stmt *sql.Stmt) *sql.Stmt
	StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
}

var _ SQLTx = (*sql.Tx)(nil)

var _ SQLRow = (*sql.Row)(nil)

// Commit commit
type Commit interface {
	Commit() error
}

// SQLCommit  commit
type SQLCommit = Commit

// Rollback rollback
type Rollback interface {
	Rollback() error
}

// SQLRollback Rollback
type SQLRollback = Rollback

// SQLStmtContext Stmt with Context
type SQLStmtContext interface {
	StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
}
