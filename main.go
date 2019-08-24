package main

import (
	"awesomePan/handle"
	"net/http"
)


func main(){
	//文件接口
	http.HandleFunc("/file/upload",handle.UploadHandler)
	http.HandleFunc("/file/upload/suc",handle.UploadSucHandler)
	http.HandleFunc("/file/meta",handle.GetFileMetaHandler)
	http.HandleFunc("/file/query",handle.HttpInterceptor(handle.FileQueryHandler))
	http.HandleFunc("/file/download",handle.DownloadHandler)
	http.HandleFunc("/file/update",handle.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete",handle.FileDeleteHandler)
	http.HandleFunc("/file/fastupload",handle.TryFastUploadHandler)

	//用户接口
	http.HandleFunc("/user/signup",handle.SignupHandler)
	http.HandleFunc("/user/signin",handle.SigninHandler)
	http.HandleFunc("/user/info",handle.HttpInterceptor(handle.UserInfoHandler))
	http.ListenAndServe(":8080",nil)

}
