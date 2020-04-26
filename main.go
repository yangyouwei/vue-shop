package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

//{
//"data": {
//"id": 500,
//"rid": 0,
//"username": "admin",
//"mobile": "123",
//"email": "123@qq.com",
//"token": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjUwMCwicmlkIjowLCJpYXQiOjE1MTI1NDQyOTksImV4cCI6MTUxMjYzMDY5OX0.eGrsrvwHm-tPsO9r_pxHIQ5i5L1kX9RX444uwnRGaIM"
//},
//"meta": {
//"msg": "登录成功",
//"status": 200
//}
//}


type User struct {
	Id int	`json:"id"`
	Rid int	`json:"rid"`
	UserName string	`json:"username"`
	Password string `json:"passowrd"`
	Moblie string	`json:"mobile"`
	Email string	`json:"email"`
	Token string `json:"token"`
}

type LoginUser struct {
	Name string  `json:"name"`
	Password string `json:"password"`
}

type MetaMessages struct {
	Messages string  `json:"msg"`
	Status int	`json:"status"`
}

type LoginMessages struct {
	Data User	`json:"data"`
	Meta MetaMessages `json:"meta"`
}

var SuccessMetastr MetaMessages
var FailMetasrName MetaMessages
var FailMetasrPassword MetaMessages
var ResponeMessages LoginMessages

var Auser *User

func init()  {
	//登录成功messages
	SuccessMetastr = MetaMessages{"登录成功",200}

	//登录失败messages
	FailMetasrName = MetaMessages{"用户不存在",401}
	FailMetasrPassword = MetaMessages{"密码不正确",401}

	//创建用户
	Auser = Auser.CreateUser()
}

func main()  {
	var serverport string = "8888"
	//打印log
	log.Println("server is listen on 0.0.0.0:8888")
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	//get user info
	s.HandleFunc("/login", Loginapi)

	err := http.ListenAndServe(":"+serverport, r)
	if err != nil {
		fmt.Println(err)
		return
	}
}


func Loginapi(w http.ResponseWriter, r *http.Request) {
	CrossDomain(w, r)
	userjson, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	auser := LoginUser{}
	// 将json数据解析到结构体实例中
	err = json.Unmarshal(userjson, &auser)
	if err != nil {
		fmt.Println(err)
	}
	if auser.Name == Auser.UserName {
		if auser.Password == Auser.Password {
			ResponeMessages.Data = *Auser
			ResponeMessages.Meta = SuccessMetastr
			respjson, err := json.Marshal(ResponeMessages)
			if err != nil {
				fmt.Println(err)
			}
			w.Write(respjson)
		}else {
			ResponeMessages.Data = User{}
			ResponeMessages.Meta = FailMetasrPassword
			respjson, err := json.Marshal(ResponeMessages)
			if err != nil {
				fmt.Println(err)
			}
			w.Write(respjson)
		}
	}else {
		ResponeMessages.Data = User{}
		ResponeMessages.Meta = FailMetasrName
		respjson, err := json.Marshal(ResponeMessages)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(respjson)
	}
}

func CrossDomain(w http.ResponseWriter, r *http.Request) http.ResponseWriter {
	log.Println("request domain ", r.Host, "URL: ", r.URL)
	w.Header().Set("Access-Control-Allow-Origin", "*")                                                                                              //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Connection, User-Agent, Cookie,Action, Module") //header的类型
	w.Header().Set("content-type", "application/json")                                                                                              //返回数据格式是json
	//header("Access-Control-Allow-Credentials : true");
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	return w
}

func TokenGener() string {
	return "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjUwMCwicmlkIjowLCJpYXQiOjE1MTI1NDQyOTksImV4cCI6MTUxMjYzMDY5OX0.eGrsrvwHm-tPsO9r_pxHIQ5i5L1kX9RX444uwnRGaIM"
}

func (this *User)GetToken()  {
	this.Token = TokenGener()
}

func (this *User)CreateUser()*User {
	//实例化一个user
	auser := User{Id:0,Rid:1,UserName:"admin",Password:"123456",Moblie:"12314235346",Email:"safas@163.com"}

	//生成一个token
	auser.GetToken()
	this = &auser
	return this
}
