package images

import (
	"../../../constants"
	"testing"

	"../../db"
)

func TestImages(t *testing.T) {
	var item Upload
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Table("preupload").Find(&item)
	t.Log(item)

}
