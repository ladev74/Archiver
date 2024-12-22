package vlc

import (
	"reflect"
	"testing"
)

func Test_preaperText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "first test",
			str:  "My name is Ted",
			want: "!my name is !ted",
		},
		{
			name: "second test",
			str:  "nfjBUULef",
			want: "nfj!b!u!u!lef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeBin(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "first test",
			str:  "!ted",
			want: "001000100110100101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBin(tt.str); got != tt.want {
				//t.Errorf("prepareText() = #{got}, want#{tt.want}")
				t.Errorf("encodeBin() = got: %s, want: %s", got, tt.want)
			}
		})
	}
}

func Test_Encode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want []byte
	}{
		{
			name: "first test",
			str:  "My name is Ted",
			want: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Decode(t *testing.T) {
	tests := []struct {
		name        string
		encodedText []byte
		want        string
	}{
		{
			name:        "first test",
			encodedText: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
			want:        "My name is Ted",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.encodedText); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
