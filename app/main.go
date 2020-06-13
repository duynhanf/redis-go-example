package main

import (
	"fmt"
	"net/http"

	"github.com/duynhanf/redis-go-example/cache"
	"github.com/duynhanf/redis-go-example/entity"
)

var (
	appCache cache.AppCache = cache.NewRedisCache("localhost:6379", 0, 10)
)

func main() {

	http.HandleFunc("/getUser", getUser)
	http.HandleFunc("/postUser", postUser)

	fmt.Println("run server")
	http.ListenAndServe(":8080", nil)
}

func getUser(w http.ResponseWriter, req *http.Request) {
	value := appCache.Get("user1")

	fmt.Fprintf(w, "%#v\n", value)
}

func postUser(w http.ResponseWriter, req *http.Request) {
	user := &entity.User{
		FullName: "Bui Duy Nhan",
		Age:      10,
	}
	appCache.Set("user1", user)
	fmt.Fprintf(w, "postUser OK\n")
}
