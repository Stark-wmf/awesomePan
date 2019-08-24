package meta

import (
	mydb "awesomePan/db"
	"fmt"
)
//文件元信息
type FileMeta struct {

	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init(){
 fileMetas=make(map[string]FileMeta)
}
//新增/更新 文件云信息
func UpdateFileMetas(fmeta FileMeta){
	fileMetas[fmeta.FileSha1]=fmeta
}
//新增/更新 文件云信息 到MySql上
func UpdateFileMetaDB(fmeta FileMeta)bool{
	return mydb.OnFileUploadFinished(
		fmeta.FileSha1,fmeta.FileName,fmeta.FileSize,fmeta.Location)
}

func GetFileMeta(fileSha1 string)FileMeta{
	return fileMetas[fileSha1]

}
//从MySql获取文件云信息
func GetFileMetaDB(fileSha1 string)(FileMeta,error){
	tfile,err:=mydb.GetFileMeta(fileSha1)
	if err!=nil{
		fmt.Println(err.Error())
		return FileMeta{},nil
	}
	fmeta:=FileMeta{
		FileSha1:tfile.FileHash,
		FileName:tfile.FileName.String,
		FileSize:tfile.FileSize.Int64,
		Location:tfile.FileAddr.String,
	}

return fmeta,nil
}
//删除元信息
func RemoveFileMeta(fileSha1 string){
	delete(fileMetas,fileSha1)

}