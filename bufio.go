// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu (duv123+git@baidu.com)
// Date: 2020/12/20

package gocore

import (
	"bufio"
	"io"
)

// BufIOReader interface of *bufio.Reader
type BufIOReader interface {
	Buffered() int
	Discard(n int) (discarded int, err error)
	Peek(n int) ([]byte, error)
	Read(p []byte) (n int, err error)
	ReadByte() (byte, error)
	ReadBytes(delim byte) ([]byte, error)
	ReadLine() (line []byte, isPrefix bool, err error)
	ReadRune() (r rune, size int, err error)
	ReadSlice(delim byte) (line []byte, err error)
	ReadString(delim byte) (string, error)
	Reset(r io.Reader)
	Size() int
	UnreadByte() error
	UnreadRune() error
	WriteTo(w io.Writer) (n int64, err error)
}

var _ BufIOReader = (*bufio.Reader)(nil)

// BufIOScanner interface of *bufio.Scanner
type BufIOScanner interface {
	Buffer(buf []byte, max int)
	Bytes() []byte
	Err() error
	Scan() bool
	Split(split bufio.SplitFunc)
	Text() string
}

var _ BufIOScanner = (*bufio.Scanner)(nil)

// BufIOWriter interface of *bufio.Writer
type BufIOWriter interface {
	Available() int
	Buffered() int
	Flush() error
	ReadFrom(r io.Reader) (n int64, err error)
	Reset(w io.Writer)
	Size() int
	Write(p []byte) (nn int, err error)
	WriteByte(c byte) error
	WriteRune(r rune) (size int, err error)
	WriteString(s string) (int, error)
}

var _ BufIOWriter = (*bufio.Writer)(nil)
