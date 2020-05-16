package images

import (
	"../../../constants"
	"../../db"

	"testing"
)

func TestImages(t *testing.T) {
	var item Upload
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Table("preupload").Find(&item)
	t.Log(item)

}
