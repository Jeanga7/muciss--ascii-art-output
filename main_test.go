package main

import (
	"reflect"
	"testing"
)

func TestIsPrintable(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "PrintableString",
			s:    "Hello, World!",
			want: true,
		},
		{
			name: "NonPrintableString",
			s:    "Hello\x00World",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrintable(tt.s); got != tt.want {
				t.Errorf("IsPrintable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFinder(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		wantNn [][]int
	}{
		{
			name: "ValidSearch",
			str:  "|",
			wantNn: [][]int{
				{828},
			},
		},
		{
			name: "ValidSearch1",
			str:  "bonjour",
			wantNn: [][]int{
				{594, 711, 702, 666, 711, 765, 738},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNn := Finder(tt.str); !reflect.DeepEqual(gotNn, tt.wantNn) {
				t.Errorf("Finder() = %v, want %v", gotNn, tt.wantNn)
			}
		})
	}
}
