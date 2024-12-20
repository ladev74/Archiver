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

func Test_splitBychunks(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}

	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "first test",
			args: args{
				bStr:      "001000100110100101",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100010", "01101001", "01000000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.bStr, tt.args.chunkSize); reflect.DeepEqual(got, tt.want) {
				//t.Errorf("prepareText() = #{got}, want#{tt.want}")
				t.Errorf("splitByChunks() = got %v, want %v", got, tt.want)
			}
		})
	}
}
