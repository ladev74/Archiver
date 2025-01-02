package shannonfano

import (
	"reflect"
	"testing"
)

func Test_bestDivierPosition(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  int
	}{
		{
			name: "once element",
			codes: []code{
				{Quanity: 2},
			},
			want: 0,
		},
		{
			name: "two element",
			codes: []code{
				{Quanity: 2},
				{Quanity: 2},
			},
			want: 1,
		},
		{
			name: "three element",
			codes: []code{
				{Quanity: 2},
				{Quanity: 1},
				{Quanity: 1},
			},
			want: 1,
		},
		{
			name: "many element",
			codes: []code{
				{Quanity: 2},
				{Quanity: 2},
				{Quanity: 1},
				{Quanity: 1},
				{Quanity: 1},
				{Quanity: 1},
			},
			want: 2,
		},
		{
			name: "uncertainty (need rightmost)",
			codes: []code{
				{Quanity: 1},
				{Quanity: 1},
				{Quanity: 1},
			},
			want: 1,
		},
		{
			name: "uncertainty (need rightmost)",
			codes: []code{
				{Quanity: 2},
				{Quanity: 2},
				{Quanity: 1},
				{Quanity: 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestDivierPosition(tt.codes); got != tt.want {
				t.Errorf("bestDivierPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assignCodes(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  []code
	}{
		{
			name: "two elements",
			codes: []code{
				{Quanity: 2},
				{Quanity: 2},
			},
			want: []code{
				{Quanity: 2, Bits: 0, Size: 1},
				{Quanity: 2, Bits: 1, Size: 1},
			},
		},
		{
			name: "three elements, certain possition",
			codes: []code{
				{Quanity: 2},
				{Quanity: 1},
				{Quanity: 1},
			},
			want: []code{
				{Quanity: 2, Bits: 0, Size: 1},
				{Quanity: 1, Bits: 2, Size: 2},
				{Quanity: 1, Bits: 3, Size: 2},
			},
		},
		{
			name: "three elements, uncertain possition",
			codes: []code{
				{Quanity: 1},
				{Quanity: 1},
				{Quanity: 1},
			},
			want: []code{
				{Quanity: 1, Bits: 0, Size: 1},
				{Quanity: 1, Bits: 2, Size: 2},
				{Quanity: 1, Bits: 3, Size: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assignCodes(tt.codes)
			if !reflect.DeepEqual(tt.codes, tt.want) {
				t.Errorf("got: %v, want %v", tt.codes, tt.want)
			}
		})
	}
}

func Test_build(t *testing.T) {
	tests := []struct {
		name string
		text string
		want encodingTable
	}{
		{
			name: "base test",
			text: "abbbcc",
			want: encodingTable{
				'a': code{
					Char:    'a',
					Quanity: 1,
					Bits:    3,
					Size:    2,
				},
				'b': code{
					Char:    'b',
					Quanity: 3,
					Bits:    0,
					Size:    1,
				},
				'c': code{
					Char:    'c',
					Quanity: 2,
					Bits:    2,
					Size:    2,
				},
			},
		},
		{
			name: "second test",
			text: "aabbcc",
			want: encodingTable{
				'a': code{
					Char:    'a',
					Quanity: 2,
					Bits:    0,
					Size:    1,
				},
				'b': code{
					Char:    'b',
					Quanity: 2,
					Bits:    2,
					Size:    2,
				},
				'c': code{
					Char:    'c',
					Quanity: 2,
					Bits:    3,
					Size:    2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := build(newCharStat(tt.text)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("build() = %v, want %v", got, tt.want)
			}
		})
	}
}
