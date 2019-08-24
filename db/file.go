package db

import (
	mydb "awesomePan/db/mysql"
	"database/sql"
	"fmt"
)

//文件上传完成后保存meta
func OnFileUploadFinished(filehash string,filename string,filesize int64,fileaddr string)bool{
	stmt,err:=mydb.DBConn().Prepare(
	"insert into tbl_file(file_sha1,file_name,file_size,"+
		"file_addr,status)values(?,?,?,?,1)")
		if err!=nil{
			fmt.Println("fail to prepare statement")
			return false
		}
	defer stmt.Close()

	ret,err:=stmt.Exec(filehash,filename,filesize,fileaddr)
	if err!=nil{
		fmt.Println(err.Error())
		return false
	}
	if rf,err:=ret.RowsAffected();nil==err{
		if rf<0{
			fmt.Println("filehash has been uploaded before",filehash)
		}
		return true
	}
	return false
}
type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}
//从MySql 获取文件元信息
func GetFileMeta(filehash string)(*TableFile ,error){
  stmt,err:= mydb.DBConn().Prepare(
   	"select file_sha1,file_addr,file_name,file_size from tbl_file where file_sha1= ? and status=1 limit 1")
   if err!=nil{
   	fmt.Println(err.Error())
	   return nil,err
   }
  defer stmt.Close()
  tfile:=TableFile{}
  err=stmt.QueryRow(filehash).Scan(
  	&tfile.FileHash,&tfile.FileAddr,&tfile.FileName,&tfile.FileSize)
  if err !=nil{
  	fmt.Println(err.Error())
  	return nil,err
  }
  return &tfile,err
}