package main

import (
  "fmt"
  "log"
  "net/http"
  "strings"
  "html/template"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
  //解析url传递的参数，对于POST则解析响应包的主体(request body)
  //注意，如果没有调用ParseForm方法，下面无法获取表单的数据
  r.ParseForm()
  fmt.Println(r.Form)
  fmt.Println("path", r.URL.Path)
  fmt.Println("scheme", r.URL.Scheme)
  fmt.Println(r.Form["url_long"])
  for k, v := range r.Form {
      fmt.Println("key:", k)
      fmt.Println("val:", strings.Join(v, ""))
  }
  //写入到ｗ的是输出到客户端的
  fmt.Fprintf(w, "<h1>Hello,World</h1>")
}

func login(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  fmt.Println("method:", r.Method) //获取请求的方法
  if r.Method == "GET" {
    t,_ := template.ParseFiles("login.html")
    t.Execute(w, nil)
  } else {
    fmt.Println("username:", r.Form["username"])
    fmt.Println("password:", r.Form["password"])
  }
}

func main() {
  http.HandleFunc("/", sayHelloName)
  http.HandleFunc("/login", login)
  err := http.ListenAndServe(":9090", nil)
  if err != nil {
    log.Fatal("ListenAndServer:", err)
  }
}
