package encoding

import (
	"encoding/base64"
	"os"
	"testing"
)

func TestEncodeBase64(t *testing.T) {
	savePath := "/images/upload/"
	fileName := "0BCAE30B-2BE8-4CF4-82AB-1B49455F88BA.jpg"
	file, err := os.Open(savePath + fileName)
	if err != nil {
		t.Error("ファイルが読み込めませんでした。")
	}
	type args struct {
		savePath string
		fileName string
	}
	fi, _ := file.Stat() // interface
	size := fi.Size()    // file size

	data := make([]byte, size)
	file.Read(data)
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				savePath: savePath,
				fileName: fileName,
			},
			want: base64.StdEncoding.EncodeToString(data),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeBase64(tt.args.savePath, tt.args.fileName); got != tt.want {
				t.Errorf("EncodeBase64() = %v, want %v", got, tt.want)
			}
		})
	}

}
