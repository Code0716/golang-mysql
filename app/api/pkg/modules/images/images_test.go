package images

import (
	"../../../constants"
	"../../db"

	"testing"
)

func TestImages(t *testing.T) {
	var item Preupload
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Table("upload").Find(&item)
	t.Log(item)
}
