package file

import (
	"os"
	"testing"
)

func setup() {
	content := "some test content"
	fValid, _ := os.OpenFile("valid.txt", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0600)
	defer fValid.Close()
	fValid.WriteString(content)

	fChmod, _ := os.OpenFile("chmod.txt", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0000)
	defer fChmod.Close()
	fChmod.WriteString(content)

	fEmpty, _ := os.OpenFile("empty.txt", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0600)
	defer fEmpty.Close()
}

func shutdown() {
	os.Remove("valid.txt")
	os.Remove("chmod.txt")
	os.Remove("empty.txt")
	f, err := os.OpenFile("exist.txt", os.O_RDONLY, 0600)
	if err == nil {
		f.Close()
		os.Remove("exist.txt")
	}
}

func TestMain(m *testing.M) {
	//setup()
	code := m.Run()
	//shutdown()
	os.Exit(code)
}

func TestClear(t *testing.T) {
	setup()
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
	}{{name: "valid file", args: args{fileName: "valid.txt"}},
		{name: "empty file", args: args{fileName: "empty.txt"}},
		{name: "file doesn't exist", args: args{fileName: "exist.txt"}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Clear(tt.args.fileName)
		})
		content, _ := os.ReadFile(tt.args.fileName)
		if len(content) != 0 {
			t.Errorf("Clear(): content remains = %v, want empty string", string(content))
		}
	}
	shutdown()
}

func TestContent(t *testing.T) {
	setup()
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{name: "valid file", args: args{fileName: "valid.txt"}, want: "some test content"},
		{name: "empty file", args: args{fileName: "empty.txt"}, want: ""},
		{name: "file doesn't exist", args: args{fileName: "exist.txt"}, want: ""}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Content(tt.args.fileName); got != tt.want {
				t.Errorf("Content() = %v, want %v", got, tt.want)
			}
		})
	}
	shutdown()
}

//func TestRewrite(t *testing.T) {
//	setup()
//	type args struct {
//		fileName string
//		content  string
//		want     string
//	}
//	tests := []struct {
//		name string
//		args args
//	}{{name: "valid file", args: args{fileName: "valid.txt", content: "new content", want: "new content\n"}},
//		{name: "empty file", args: args{fileName: "empty.txt", content: "new content", want: "new content\n"}},
//		{name: "file doesn't exist", args: args{fileName: "exist.txt", content: "new content", want: "new content\n"}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			Rewrite(tt.args.fileName, tt.args.content)
//		})
//		content, _ := os.ReadFile(tt.args.fileName)
//		if string(content) != tt.args.want {
//			t.Errorf("Rewrite() = %s, want %v", string(content), tt.args.want)
//		}
//	}
//	shutdown()
//}

func TestUpdate(t *testing.T) {
	setup()
	type args struct {
		fileName string
		content  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "valid file", args: args{fileName: "valid.txt", content: "\nnew content"}, want: "some test content\nnew content\n", wantErr: false},
		{name: "file has 000 rights", args: args{fileName: "chmod.txt", content: "new content"}, want: "some test content\n", wantErr: true},
		{name: "empty file", args: args{fileName: "empty.txt", content: "new content"}, want: "new content\n", wantErr: false},
		{name: "file doesn't exist", args: args{fileName: "exist.txt", content: "new content"}, want: "new content\n", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Update(tt.args.fileName, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err == nil {
				content, _ := os.ReadFile(tt.args.fileName)
				converted := string(content)
				if converted != tt.want {
					t.Errorf("Update() = %s, want %v", string(content), tt.want)
				}
			}
		})
	}
	shutdown()
}

func TestRewrite(t *testing.T) {
	setup()
	type args struct {
		fileName string
		content  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "valid file", args: args{fileName: "valid.txt", content: "new content"}, want: "new content\n", wantErr: false},
		{name: "file has 000 rights", args: args{fileName: "chmod.txt", content: "new content"}, want: "new content\n", wantErr: true},
		{name: "empty file", args: args{fileName: "empty.txt", content: "new content"}, want: "new content\n", wantErr: false},
		{name: "file doesn't exist", args: args{fileName: "exist.txt", content: "new content"}, want: "new content\n", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Rewrite(tt.args.fileName, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Rewrite() error = %v, wantErr %v", err, tt.wantErr)
			} else if err == nil {
				content, _ := os.ReadFile(tt.args.fileName)
				if string(content) != tt.want {
					t.Errorf("Rewrite() = %s, want %v", string(content), tt.want)
				}
			}
		})
	}
	shutdown()
}
