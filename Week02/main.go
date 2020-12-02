package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"week02/dao"
	"week02/service"
)

func searchUser(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("user_id")
	user, err := service.FindUserById(id)
	if err != nil {
		if errors.Is(err, dao.ErrorNotFound) {
			// ErrorNotFound 的处理
			w.Write([]byte(fmt.Sprintf("{\"code\": 404, \"message\": \"user: %v not found\"}", id)))
		} else {
			log.Printf("%+v\n", err)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		}
	}
	result := fmt.Sprintf("{\"id\": \"%v\", \"name\": \"%v\", \"age\": %v}", id, user.Name, user.Age)
	w.Write([]byte(fmt.Sprintf("{\"code\": 0, \"result\": %v}", result)))
}

func main() {
	http.HandleFunc("/search", searchUser)
	err := http.ListenAndServe("127.0.0.1:10000", nil)
	if err != nil {
		log.Println("server stop.")
	}
}
