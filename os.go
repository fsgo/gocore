/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/12/20
 */

package gocore

import (
	"os"
	"time"
)

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

// OsProcess interface of *os.Process
type OsProcess interface {
	Kill() error
	Release() error
	Signal(sig os.Signal) error
	Wait() (*os.ProcessState, error)
}

var _ OsProcess = (*os.Process)(nil)

// OsProcessState interface of *os.ProcessState
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

// Kill kill
type Kill interface {
	Kill() error
}

var _ OsFile = (*os.File)(nil)
