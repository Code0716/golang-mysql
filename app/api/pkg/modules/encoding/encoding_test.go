package encoding

import (
	"encoding/base64"
	"fmt"
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

	mockData := make([]byte, size)
	file.Read(mockData)
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "画像通る場合",
			args: args{
				savePath: savePath,
				fileName: fileName,
			},
			want: base64.StdEncoding.EncodeToString(mockData),
		},
		{
			name: "画像通らない場合",
			args: args{
				savePath: savePath,
				fileName: "のおおおおお",
			},
			want: "open /images/upload/のおおおおお: no such file or directory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeBase64(tt.args.savePath, tt.args.fileName); got != tt.want {
				fmt.Println("ここ", got)
				t.Errorf("EncodeBase64() = %v, want %v", got, tt.want)
			}
		})
	}

}
