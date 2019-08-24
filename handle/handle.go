package handle

import (
	dblayer "awesomePan/db"
	"awesomePan/meta"
	"awesomePan/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func UploadHandler(w http.ResponseWriter,r *http.Request){
	if r.Method=="GET" {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "err")
			return
		}
		io.WriteString(w, string(data))
	}else if r.Method=="POST"{
		file,head,err:=r.FormFile("file")
		if err!=nil{
			fmt.Println("fail to get file")
			return
		}
		defer file.Close()
		fileMeta:=meta.FileMeta{
			FileName:head.Filename,
			Location:"./static/tmpwmf/"+head.Filename,
			UploadAt:time.Now().Format("2006-01-02 15:04:05"),
		}
		newFile,err:=os.Create(fileMeta.Location)
		if err!=nil{
			fmt.Println("fail to create file")
			return
		}
		defer  file.Close()
		fileMeta.FileSize,err=io.Copy(newFile,file)
		if err!=nil{
			fmt.Println("fail to save data into file")
			return
		}
		newFile.Seek(0,0)
		fileMeta.FileSha1=util.FileSha1(newFile)
		_=meta.UpdateFileMetaDB(fileMeta)
		//更新上传的文件到用户文件表
		r.ParseForm()

		username:=r.Form.Get("username")

		suc:=dblayer.OnUserFileUploadFinished(username,fileMeta.FileSha1,fileMeta.FileName,fileMeta.FileSize)
		if suc{
			http.Redirect(w,r,"/file/upload/suc",http.StatusFound)
		}else{
			w.Write([]byte("Upload Failed"))
		}

	}



}
func UploadSucHandler(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"Upload Sucess")
}

func GetFileMetaHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	filehash:=r.Form["filehash"][0]
	//fMeta:=meta.GetFileMeta(filehash)
	fMeta,err:=meta.GetFileMetaDB(filehash)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dara,err:=json.Marshal(fMeta)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(dara)
}

func FileQueryHandler(w http.ResponseWriter,r*http.Request){
	r.ParseForm()

	limitCnt,_:=strconv.Atoi(r.Form.Get("limit"))
	username:=r.Form.Get("username")
	userFiles,err:=dblayer.QueryUserFileMetas(username,limitCnt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data,err:=json.Marshal(userFiles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
func DownloadHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fsha1:=r.Form.Get("filehash")
	fm:=meta.GetFileMeta(fsha1)

	f,err:=os.Open(fm.Location)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	data,err:=ioutil.ReadAll(f)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("2")
		return
	}
	w.Header().Set("Content-Type","application/octect-stream")
	w.Header().Set("Content-Description","attachment;filename=\""+fm.FileName+"\"")
	w.Write(data)
}

func FileMetaUpdateHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
//op操作类型
	opType:=r.Form.Get("op")
	fileSha1:=r.Form.Get("filehash")
	newFileName:=r.Form.Get("filename")

	if opType!="0"{
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method!="POST"{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	curFileMeta :=meta.GetFileMeta(fileSha1)
	curFileMeta.FileName=newFileName
//	meta.UpdateFileMetas(curFileMeta)
     _=meta.UpdateFileMetaDB(curFileMeta)

	w.WriteHeader(http.StatusOK)
	data,err:=json.Marshal(curFileMeta)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
//删除文件及元信息
func FileDeleteHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fileSha1:=r.Form.Get("filehash")
	fMeta:=meta.GetFileMeta(fileSha1)

    os.Remove(fMeta.Location)
	meta.RemoveFileMeta(fileSha1)

	w.WriteHeader(http.StatusOK)
}

func TryFastUploadHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()

	//解析请求参数
	username:=r.Form.Get("username")
	filehash:=r.Form.Get("filehash")
	filename:=r.Form.Get("filename")
	filesize,_:=strconv.Atoi(r.Form.Get("filesize"))
    //从文件表中查询hash相同的文件记录
    fileMeta,err:=meta.GetFileMetaDB(filehash)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
    //查不到数据的话返回秒传失败
    if fileMeta.FileSize<=0{
    	resp:=util.RespMsg{
    		Code:-1,
    		Msg:"秒传失败，跳转普通上传接口",
		}
    	w.Write(resp.JSONBytes())
		return
    	}
    suc:=dblayer.OnUserFileUploadFinished(username,filehash,filename,int64(filesize))
    if suc{
    	resp:=util.RespMsg{
    		Code:0,
    		Msg:"秒传成功",
		}
    	w.Write(resp.JSONBytes())
		return
	}else{
		resp:=util.RespMsg{
			Code:-2,
			Msg:"秒传失败，稍后重试",
		}
		w.Write(resp.JSONBytes())
		return

	}
}
