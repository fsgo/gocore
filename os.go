/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/12/20
 */

package gocore

import (
	"io"
	"os"
	"syscall"
	"time"
)

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
	ReadFrom(r io.Reader) (n int64, err error)
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

var _ OsFile = (*os.File)(nil)

// SetDeadline can set deadline
type SetDeadline interface {
	SetDeadline(t time.Time) error
}

// SetReadDeadline can set readDeadline
type SetReadDeadline interface {
	SetReadDeadline(t time.Time) error
}

// SetWriteDeadline  can set writeDeadline
type SetWriteDeadline interface {
	SetWriteDeadline(t time.Time) error
}

// SetReadWriteDeadline can set read and write deadline
type SetReadWriteDeadline interface {
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}

// OsProcess interface off *os.Process
type OsProcess interface {
	Kill() error
	Release() error
	Signal(sig os.Signal) error
	Wait() (*os.ProcessState, error)
}

var _ OsProcess = (*os.Process)(nil)

// OsProcessState interface off *os.ProcessState
type OsProcessState interface {
	ExitCode() int
	Exited() bool
	Pid() int
	String() string
	Success() bool
	Sys() interface{}
	SysUsage() interface{}
	SystemTime() time.Duration
	UserTime() time.Duration
}

var _ OsProcessState = (*os.ProcessState)(nil)
