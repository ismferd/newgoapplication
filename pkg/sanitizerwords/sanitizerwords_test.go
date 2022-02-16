package sanitizerwords

import "testing"

func TestCleanWord(t *testing.T) {
	type args struct {
		s string
	}
	expected := []struct {
		name string
		args args
		want string
	}{
		{"Normal Word", args{"Hello"}, "hello"},
		{"Contractions", args{"shouldn't"}, "shouldnt"},
		{"Special Characters", args{"***Hello***"}, "hello"},
		{"Special Characters", args{"***"}, ""},
		{"Special letters", args{"ü"}, "ü"},
		{"Special letters", args{"語"}, "語"},
		{"Special Characters", args{"/"}, ""},
		{"Special Characters", args{","}, ""},
		{"Special Characters", args{"."}, ""},
		{"Special Characters", args{"2"}, ""},
	}
	for _, tt := range expected {
		t.Run(tt.name, func(t *testing.T) {
			if got := SanitizerWords(tt.args.s); got != tt.want {
				t.Errorf("cleanWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
