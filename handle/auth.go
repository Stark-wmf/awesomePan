package handle

import (
	"fmt"
	"net/http"
)
//token实现的http请求拦截器
func HttpInterceptor(h http.HandlerFunc)http.HandlerFunc{
	return http.HandlerFunc(
	func(w http.ResponseWriter,r*http.Request){
		r.ParseForm()
		username:=r.Form.Get("username")
		token:=r.Form.Get("token")

		if len(username)<3||!IsTokenValid(token){
			fmt.Println("1")
			w.WriteHeader(http.StatusForbidden)
			//return
			h(w,r)
		}
		h(w,r)
	}	)
}
