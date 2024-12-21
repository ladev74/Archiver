package vlc

import (
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
				//t.Errorf("prepareText() = #{got}, want#{tt.want}")
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
		want string
	}{
		{
			name: "first test",
			str:  "My name is Ted",
			want: "20 30 3C 18 77 4A E4 4D 28",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.str); got != string(tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Decode(t *testing.T) {
	tests := []struct {
		name        string
		encodedText string
		want        string
	}{
		{
			name:        "first test",
			encodedText: "20 30 3C 18 77 4A E4 4D 28",
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
