/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/12/20
 */

package gocore

import (
	"bytes"
	"io"
	"sync"
)

// BytesBuffer interface of *bytes.Buffer
type BytesBuffer interface {
	Bytes() []byte
	Cap() int
	Grow(n int)
	Len() int
	Next(n int) []byte
	Read(p []byte) (n int, err error)
	ReadByte() (byte, error)
	ReadBytes(delim byte) (line []byte, err error)
	ReadFrom(r io.Reader) (n int64, err error)
	ReadRune() (r rune, size int, err error)
	ReadString(delim byte) (line string, err error)
	Reset()
	String() string
	Truncate(n int)
	UnreadByte() error
	UnreadRune() error
	Write(p []byte) (n int, err error)
	WriteByte(c byte) error
	WriteRune(r rune) (n int, err error)
	WriteString(s string) (n int, err error)
	WriteTo(w io.Writer) (n int64, err error)
}

var _ BytesBuffer = (*bytes.Buffer)(nil)

// BytesReader interface  *bytes.Reader
type BytesReader interface {
	Len() int
	Read(b []byte) (n int, err error)
	ReadAt(b []byte, off int64) (n int, err error)
	ReadByte() (byte, error)
	ReadRune() (ch rune, size int, err error)
	Reset(b []byte)
	Seek(offset int64, whence int) (int64, error)
	Size() int64
	UnreadByte() error
	UnreadRune() error
	WriteTo(w io.Writer) (n int64, err error)
}

var _ BytesReader = (*bytes.Reader)(nil)

type lockBuffer struct {
	lock sync.Mutex
	raw  BytesBuffer
}

func (l *lockBuffer) withLock(fn func()) {
	l.lock.Lock()
	defer l.lock.Unlock()
	fn()
}

func (l *lockBuffer) Cap() (num int) {
	l.withLock(func() {
		num = l.raw.Cap()
	})
	return
}

func (l *lockBuffer) Grow(n int) {
	l.withLock(func() {
		l.raw.Grow(n)
	})
}

func (l *lockBuffer) Len() (num int) {
	l.withLock(func() {
		num = l.raw.Len()
	})
	return
}

func (l *lockBuffer) Next(n int) []byte {
	var b []byte
	l.withLock(func() {
		b = append(b, l.raw.Next(n)...)
	})
	return b
}

func (l *lockBuffer) Read(p []byte) (n int, err error) {
	l.withLock(func() {
		n, err = l.raw.Read(p)
	})
	return
}

func (l *lockBuffer) ReadByte() (b byte, e error) {
	l.withLock(func() {
		b, e = l.raw.ReadByte()
	})
	return
}

func (l *lockBuffer) ReadBytes(delim byte) (line []byte, err error) {
	var result []byte
	l.withLock(func() {
		result, err = l.raw.ReadBytes(delim)
		line = append(line, result...)
	})
	return
}

func (l *lockBuffer) ReadFrom(r io.Reader) (n int64, err error) {
	l.withLock(func() {
		n, err = l.raw.ReadFrom(r)
	})
	return
}

func (l *lockBuffer) ReadRune() (r rune, size int, err error) {
	l.withLock(func() {
		r, size, err = l.raw.ReadRune()
	})
	return
}

func (l *lockBuffer) ReadString(delim byte) (line string, err error) {
	l.withLock(func() {
		line, err = l.raw.ReadString(delim)
	})
	return
}

func (l *lockBuffer) Reset() {
	l.withLock(func() {
		l.raw.Reset()
	})
}

func (l *lockBuffer) String() (str string) {
	l.withLock(func() {
		str = l.raw.String()
	})
	return
}

func (l *lockBuffer) Truncate(n int) {
	l.withLock(func() {
		l.raw.Truncate(n)
	})
}

func (l *lockBuffer) UnreadByte() (err error) {
	l.withLock(func() {
		err = l.raw.UnreadByte()
	})
	return
}

func (l *lockBuffer) UnreadRune() (err error) {
	l.withLock(func() {
		err = l.raw.UnreadRune()
	})
	return
}

func (l *lockBuffer) Write(p []byte) (n int, err error) {
	l.withLock(func() {
		n, err = l.raw.Write(p)
	})
	return
}

func (l *lockBuffer) WriteByte(c byte) (err error) {
	l.withLock(func() {
		err = l.raw.WriteByte(c)
	})
	return
}

func (l *lockBuffer) WriteRune(r rune) (n int, err error) {
	l.withLock(func() {
		n, err = l.raw.WriteRune(r)
	})
	return
}

func (l *lockBuffer) WriteString(s string) (n int, err error) {
	l.withLock(func() {
		n, err = l.raw.WriteString(s)
	})
	return
}

func (l *lockBuffer) WriteTo(w io.Writer) (n int64, err error) {
	l.withLock(func() {
		n, err = l.raw.WriteTo(w)
	})
	return
}

func (l *lockBuffer) Bytes() (b []byte) {
	l.withLock(func() {
		b = append(b, l.raw.Bytes()...)
	})
	return
}

var _ BytesBuffer = (*lockBuffer)(nil)

// NewLockedBytesBuffer BytesBuffer with lock
func NewLockedBytesBuffer(bf BytesBuffer) BytesBuffer {
	return &lockBuffer{
		raw: bf,
	}
}
