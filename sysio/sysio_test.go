package sysio_test

import (
	"testing"

	"../sysio"
)

func TestFolderName(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{"/home/me/test/ok", "ok"},
		{"/", ""},
		{".", "sysio"},
		{"./test/me/foo", "foo"},
		{"~/me/foo", "foo"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			out, _ := sysio.FolderName(tt.in)
			if out != tt.out {
				t.Fail()
			}
		})
	}
}
