// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu (duv123+git@baidu.com)
// Date: 2020/12/20

package gocore

import (
	"context"
	"testing"
)

func TestIsTimeout(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{},
			want: false,
		},
		{
			name: "case 2",
			args: args{
				err: context.DeadlineExceeded,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTimeout(tt.args.err); got != tt.want {
				t.Errorf("IsTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTemporary(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{},
			want: false,
		},
		{
			name: "case 2",
			args: args{
				err: context.DeadlineExceeded,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTemporary(tt.args.err); got != tt.want {
				t.Errorf("IsTemporary() = %v, want %v", got, tt.want)
			}
		})
	}
}
