package usecase

import "../domain"

type CityRepository interface {
	GetAll() (domain.Cities, error)
}

// package usecase

// import (
// 	"../domain"
// )

// type (
// 	OperatePreFileData interface {
// 		Add(domain.Preupload) (domain.Preupload, error)
// 		Get(int) (domain.Preupload, error)
// 		Delete(domain.Preupload) (domain.Preupload, error)
// 	}

// 	OperatePreAllData interface {
// 		GetAll() (domain.Preuploads, error)
// 	}

// 	OperateFileData interface {
// 		Add(domain.Upload) (domain.Upload, error)
// 		Get(int) (domain.Upload, error)
// 		Delete(domain.Upload) (domain.Upload, error)
// 	}

// 	OperateAllData interface {
// 		GetAll() (domain.Uploads, error)
// 	}
// )
