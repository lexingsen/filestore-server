package meta

import (
	myDB "filestore-server/db"
)

// FileMeta:文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetaes map[string]FileMeta

func init() {
	fileMetaes = make(map[string]FileMeta)
}

func UpdateFileMeta(fMeta FileMeta) {
	fileMetaes[fMeta.FileSha1] = fMeta
}

func UpdateFileMetaToDB(fMeta FileMeta) {
	myDB.OnFileUploadFinished(fMeta.FileSha1, fMeta.FileName, fMeta.FileSize, fMeta.Location)
}

func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetaes[fileSha1]
}

func RemoveFileMeta(fileSha1 string) {
	delete(fileMetaes, fileSha1)
}

// 新增/更新文件元信息到mysql
func UpdateFileMetaDB(fMeta FileMeta) bool {
	return myDB.OnFileUploadFinished(fMeta.FileSha1, fMeta.FileName, fMeta.FileSize, fMeta.Location)
}


