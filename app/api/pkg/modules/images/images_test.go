package images

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestImageController_GetFile(t *testing.T) {
	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	req, _ := http.NewRequest("GET", "/image/upload", nil)
	ginContext.Request = req

	type args struct {
		ginContext *gin.Context
	}

	tests := []struct {
		name   string
		upload ImageController
		args   args
		param  string
	}{
		{
			name:   "GetFile test1",
			upload: ImageController{},
			args:   args{ginContext: ginContext},
			param:  "70",
		},
		{
			name:   "GetFile test2",
			upload: ImageController{},
			args:   args{ginContext: ginContext},
			param:  "500000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			param := gin.Param{"id", tt.param}
			tt.args.ginContext.Params = gin.Params{param}
			upload := ImageController{}
			upload.GetFile(tt.args.ginContext)
		})
	}
}
