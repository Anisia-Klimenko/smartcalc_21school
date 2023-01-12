package file

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestClear(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
	}{{name: "valid filename", args: args{fileName: "valid.txt"}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Clear(tt.args.fileName)
		})
	}
}

func TestContent(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Content(tt.args.fileName); got != tt.want {
				t.Errorf("Content() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRewrite(t *testing.T) {
	type args struct {
		fileName string
		content  string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rewrite(tt.args.fileName, tt.args.content)
		})
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		fileName string
		content  string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Update(tt.args.fileName, tt.args.content)
		})
	}
}
