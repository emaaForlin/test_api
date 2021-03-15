package main

import (
  "log"
  "net/http"
  //"github.com/friendsofgo/gopher-api/pkg/server"
  "/home/frln/code/go/test_api/pkg/server"
)

func main(){
  s := server.New()
  log.Fatal(http.ListenAndServe(":8080",nil))
}
