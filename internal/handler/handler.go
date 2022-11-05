package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	//return
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	fmt.Println(i)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("gay")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}
