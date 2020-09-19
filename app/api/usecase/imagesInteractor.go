package usecase

import "../domain"

type OperateInteractor struct {
	OperatePreFileData OperatePreFileData
	OperatePreAllData  OperatePreAllData
	OperateFileData    OperateFileData
	OperateAllData     OperateAllData
}

// CreatePreFileInfo save  upload infogit
func (interactor *OperateInteractor) CreatePreFileInfo(f domain.Preupload) (file domain.Preupload, err error) {
	file, err = interactor.OperatePreFileData.Add(f)
	return
}

// CreateFileInfo save  upload info
func (interactor *OperateInteractor) CreateFileInfo(f domain.Upload) (file domain.Upload, err error) {
	file, err = interactor.OperateFileData.Add(f)
	return
}

// GetPreFileInfo 個々のファイルを取得
func (interactor *OperateInteractor) GetPreFileInfo(id int) (file domain.Preupload, err error) {
	file, err = interactor.OperatePreFileData.Get(id)
	return
}

// GetFileInfo 個々のファイルを取得
func (interactor *OperateInteractor) GetFileInfo(id int) (file domain.Upload, err error) {
	file, err = interactor.OperateFileData.Get(id)
	return
}

// DeletePreFileInfo 個々のファイルを削除
func (interactor *OperateInteractor) DeletePreFileInfo(f domain.Preupload) (file domain.Preupload, err error) {
	file, err = interactor.OperatePreFileData.Delete(f)
	return
}

// DeleteFileInfo 個々のファイルを削除
func (interactor *OperateInteractor) DeleteFileInfo(f domain.Upload) (file domain.Upload, err error) {
	file, err = interactor.OperateFileData.Delete(f)
	return
}

// GetAllPreFileInfo 全件取得
func (interactor *OperateInteractor) GetAllPreFileInfo() (files domain.Preuploads, err error) {
	files, err = interactor.OperatePreAllData.GetAll()
	return
}

// GetAllFileInfo 全件取得
func (interactor *OperateInteractor) GetAllFileInfo() (files domain.Uploads, err error) {
	files, err = interactor.OperateAllData.GetAll()
	return
}
