package handle

import (
	dblayer "awesomePan/db"
	"awesomePan/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	pwd_slat="ekk259"
)

//处理用户注册请求
func SignupHandler(w http.ResponseWriter,r*http.Request){
	if r.Method==http.MethodGet{
		data,err:=ioutil.ReadFile("./static/view/signup.html")
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	username:=r.Form.Get("username")
	password:=r.Form.Get("password")

	if len(username)<3||len(password)<5{
		w.Write([]byte("Invalid parameter"))
		return
	}
	enc_passwd:=util.Sha1([]byte(password+pwd_slat))
	suc:=dblayer.UserSignup(username,enc_passwd)
	if suc{
		w.Write([]byte("Sucess"))
	}
	w.Write([]byte("Fail"))
}
//登录接口
func SigninHandler(w http.ResponseWriter,r*http.Request) {
	r.ParseForm()
	username:=r.Form.Get("username")
	password:=r.Form.Get("password")

	encPasswd:=util.Sha1([]byte(password+pwd_slat))
	//1.校验用户名和密码
	pwdChecked:=dblayer.UserSignin(username,encPasswd)
	if !pwdChecked{
		w.Write([]byte("FAILED"))
		return
	}

//2.生成访问凭证（token)
    token:=GenToken(username)
    upRes:=dblayer.UpdateToken(username,token)
    if !upRes{
    	w.Write([]byte("Failed"))
		return
	}
//3.登录成功后重定向到首页
   //  w.Write([]byte("http://"+r.Host+"/static/view/home.html"))
   resp:=util.RespMsg{
   	Code:0,
   	Msg:"OK",
   	Data: struct {
		Location string
		Username string
		Token string
	}{
		Location:"http://"+r.Host+"/static/view/home.html",
		Username:username,
		Token:token,
	},
   }
       w.Write(resp.JSONBytes())

}

func UserInfoHandler(w http.ResponseWriter,r*http.Request) {
	r.ParseForm()
	username:=r.Form.Get("username")
//	token:=r.Form.Get("token")
////1.解析请求函数
////2.验证token
//isvalidToken:=IsTokenValid(token)
//	if !isvalidToken{
//		w.WriteHeader(http.StatusForbidden)
//		return
//	}
//3.查询用户信息
user,err:=dblayer.GetUserInfo(username)
if err!=nil{
	w.WriteHeader(http.StatusForbidden)
	return
}
resp:=util.RespMsg{
	Code:0,
	Msg:"OK",
	Data:user,
}
w.Write(resp.JSONBytes())
//4.组装并响应用户数据
}
func GenToken(username string)string{
	//40位字符：md5(username+timestamp+token_salt)+timestamp[:8]
	ts:=fmt.Sprint("%x",time.Now().Unix())
	tokenPrefix:=util.MD5([]byte(username+ts+"_tokensalt"))
	return tokenPrefix+ts[:8]
}

func IsTokenValid(token string)bool{
	//1.判断token是否过期
     if len(token)!=40{
     	return false
	 }
	//2.查数据库是否有这个token

	//3.对比两个token是否一致
	return true

}
