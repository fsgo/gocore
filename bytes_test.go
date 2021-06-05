// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu (duv123+git@baidu.com)
// Date: 2020/12/20

package gocore

import (
	"bytes"
	"strconv"
	"testing"
	"time"
)

func TestNewLockedBytesBuffer(t *testing.T) {
	bf := NewLockedBytesBuffer(&bytes.Buffer{})
	t.Parallel()
	for i := 0; i < 100; i++ {
		t.Run("write_"+strconv.Itoa(i), func(t *testing.T) {
			for i := 0; i < 1000; i++ {
				if n, _ := bf.WriteString("abc"); n != 3 {
					t.Errorf("WriteString %d want=3", n)
				}
				if n, _ := bf.Write([]byte("hello")); n != 5 {
					t.Errorf("Write %d want=5", n)
				}
				if err := bf.WriteByte('a'); err != nil {
					t.Errorf("WriteByte has error %v", err)
				}
				if n, _ := bf.WriteRune('你'); n != 3 {
					t.Errorf("WriteRune %d want=3", n)
				}
				time.Sleep(100 * time.Nanosecond)

				z := &bytes.Buffer{}
				z.WriteString("hello")
				if n, _ := bf.ReadFrom(z); n != 5 {
					t.Errorf("ReadFrom %d want=5", n)
				}
			}
		})

		t.Run("read_"+strconv.Itoa(i), func(t *testing.T) {
			_ = bf.String()
			_, _ = bf.ReadByte()
			_, _, _ = bf.ReadRune()
			_ = bf.UnreadRune()
			_ = bf.UnreadByte()
			_, _ = bf.ReadString('a')
			_ = bf.Bytes()

			if _, err := bf.ReadBytes('a'); err != nil {
				t.Fatalf("ReadBytes error:%v", err)
			}

			q := make([]byte, 1)
			if _, err := bf.Read(q); err != nil {
				t.Errorf("Read error:%v", err)
			}
			_ = bf.Next(1)

			z := &bytes.Buffer{}
			if _, err := bf.WriteTo(z); err != nil {
				t.Errorf("WriteTo error:%v", err)
			}

			_ = bf.Len()
			_ = bf.Cap()
		})

		t.Run("change", func(t *testing.T) {
			bf.Truncate(0)
			bf.Reset()
			bf.Grow(100)
		})
	}

}
