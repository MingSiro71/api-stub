package main

import (
  "os"
  "io/ioutil"
  "fmt"
  "net/http"
  // "bytes"
  // "api_stub/controllers"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  b, err := ioutil.ReadAll(r.Body)
  std := os.Stdout
  if err != nil {
    fmt.Fprint(std, "can not read request body")  
  }
  fmt.Fprint(std, string(b))
  fmt.Fprint(std, string(r.URL.Path))
  fmt.Fprint(std, string(r.URL.RawQuery))
  fmt.Fprint(std, string(r.Method))
}

func main() {
  // controller := controllers.NewMainController()
  http.Handle("/", String("Hello World."))
  // http.HandleFunc("/set", controller.Set)
  // http.HandleFunc("/get", controller.get)
  // http.HandleFunc("/list", controller.list)
  // http.HandleFunc("/clear", controller.clear)

  // start server with reverse proxy
  // http.ListenAndServe("localhost:8080", nil)
  //
  // start standalone server, expose 80
  http.ListenAndServe(":80", nil)
}
