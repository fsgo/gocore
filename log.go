// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu (duv123+git@baidu.com)
// Date: 2020/12/20

package gocore

import (
	"io"
	"log"
	"log/syslog"
)

// LogLogger interface of *log.Logger
type LogLogger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Flags() int
	Output(calldepth int, s string) error
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
	Prefix() string
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	SetFlags(flag int)
	SetOutput(w io.Writer)
	SetPrefix(prefix string)
	Writer() io.Writer
}

var _ LogLogger = (*log.Logger)(nil)

// LogOutput output log
type LogOutput interface {
	Output(calldepth int, s string) error
}

// LogPrintf printf log
type LogPrintf interface {
	Printf(format string, v ...interface{})
}

// LogPanicf panicf log
type LogPanicf interface {
	Panicf(format string, v ...interface{})
}

// LogFatalf fatalf log
type LogFatalf interface {
	Fatalf(format string, v ...interface{})
}

// SysLogWriter interface of *syslog.Writer
type SysLogWriter interface {
	Alert(m string) error
	Close() error
	Crit(m string) error
	Debug(m string) error
	Emerg(m string) error
	Err(m string) error
	Info(m string) error
	Notice(m string) error
	Warning(m string) error
	Write(b []byte) (int, error)
}

var _ SysLogWriter = (*syslog.Writer)(nil)
