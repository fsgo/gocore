/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/12/20
 */

package gocore

import (
	"bytes"
	"testing"
)

func TestNopWriteCloser(t *testing.T) {
	w := NopWriteCloser(&bytes.Buffer{})
	if err := w.Close(); err != nil {
		t.Fatalf("Close error:%v", err)
	}
}
