package encoding

import "testing"

func TestEncodeBase64(t *testing.T) {
	type args struct {
		savePath string
		fileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				savePath: "../../images/upload/",
				fileName: "BADCCFAE-143D-40A4-8D34-93E6D0729660-1.jpp",
			},
			want: "../../images/upload/BADCCFAE-143D-40A4-8D34-93E6D0729660-1.jpg",
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
