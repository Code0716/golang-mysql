package main

import (
	"../api/constants"
	"../api/pkg/db"
	"../api/pkg/modules/images"
)

func main() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()

	if !db.HasTable("preupload") && !db.HasTable("upload") {
		db.AutoMigrate(&images.Preupload{}, &images.Upload{})
	}
}
