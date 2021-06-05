// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu (duv123+git@baidu.com)
// Date: 2020/12/20

// +build !go1.13,!go1.14

package gocore

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"io"
	"net"
	"os"
	"syscall"
	"time"
)

// NetResolver interface off *net.Resolver
type NetResolver interface {
	LookupAddr(ctx context.Context, addr string) (names []string, err error)
	LookupCNAME(ctx context.Context, host string) (cname string, err error)
	LookupHost(ctx context.Context, host string) (addrs []string, err error)
	LookupIP(ctx context.Context, network, host string) ([]net.IP, error)
	LookupIPAddr(ctx context.Context, host string) ([]net.IPAddr, error)
	LookupMX(ctx context.Context, name string) ([]*net.MX, error)
	LookupNS(ctx context.Context, name string) ([]*net.NS, error)
	LookupPort(ctx context.Context, network, service string) (port int, err error)
	LookupSRV(ctx context.Context, service, proto, name string) (cname string, addrs []*net.SRV, err error)
	LookupTXT(ctx context.Context, name string) ([]string, error)
}

// OsFile interface off *os.File
type OsFile interface {
	Chdir() error
	Chmod(mode os.FileMode) error
	Chown(uid, gid int) error
	Close() error
	Fd() uintptr
	Name() string
	Read(b []byte) (n int, err error)
	ReadAt(b []byte, off int64) (n int, err error)
	ReadFrom(r io.Reader) (n int64, err error) // go1.15
	Readdir(n int) ([]os.FileInfo, error)
	Readdirnames(n int) (names []string, err error)
	Seek(offset int64, whence int) (ret int64, err error)
	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
	Stat() (os.FileInfo, error)
	Sync() error
	SyscallConn() (syscall.RawConn, error)
	Truncate(size int64) error
	Write(b []byte) (n int, err error)
	WriteAt(b []byte, off int64) (n int, err error)
	WriteString(s string) (n int, err error)
}

// SQLDB interface off *Sql.DB
type SQLDB interface {
	SetConnMaxIdleTime(d time.Duration) // go1.15
	SetConnMaxLifetime(d time.Duration)
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	Conn(ctx context.Context) (*sql.Conn, error)
	Driver() driver.Driver
	Stats() sql.DBStats
	Close() error

	SQLDBNoContexter
	SQLDBContexter
}

// SQLRow interface off *sql.Row
type SQLRow interface {
	Err() error // go1.15
	Scan(dest ...interface{}) error
}
