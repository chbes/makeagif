package ext_test

import (
	"testing"

	"../ext"
)

func TestIsJPEG(t *testing.T) {
	var tests = []struct {
		in  string
		out bool
	}{
		{"aaa.jpeg", true},
		{"aaa.jpg", true},
		{"aaa.JPEG", true},
		{"aaa.JPG", true},
		{".jpg", true},
		{".jpeg", true},
		{"aaa.png", false},
		{"aaa", false},
		{"az.jpeg.no", false},
		{"az.jpg.no", false},
		{"JPEG.no", false},
		{"jpg.no", false},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			out := ext.IsJPEG(tt.in)
			if out != tt.out {
				t.Fail()
			}
		})
	}
}

func TestIsPNG(t *testing.T) {
	var tests = []struct {
		in  string
		out bool
	}{
		{"aaa.png", true},
		{"aaa.PNG", true},
		{".png", true},
		{"aaa.jpg", false},
		{"aaa", false},
		{"az.png.no", false},
		{"az.PNG.no", false},
		{"png.no", false},
		{"PNG.no", false},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			out := ext.IsPNG(tt.in)
			if out != tt.out {
				t.Fail()
			}
		})
	}
}
