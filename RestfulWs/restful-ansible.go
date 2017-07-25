package main

import (
	"encoding/json"
	//"log"
	"net/http"
	//"github.com/gorilla/mux"
	"fmt"
)

type Api struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
}

var api []Api

func GetAPIList(w http.ResponseWriter, req *http.Request) {
	//json.NewEncoder(w).Encode(api)
	/*s := Spam{}
	r := restclient.RequestResponse{
		Url:    "http://10.42.0.42/api/v1",
		Method: restclient.,
		Result: &s,
	}
	status, err := restclient.Do(&r)*/
	resp, err := http.Get("http://10.42.0.42/api/v1")
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(resp)

}


func main() {
	/*router := mux.NewRouter()
	router.HandleFunc("/apilist", GetAPIList).Methods("GET")

	log.Fatal(http.ListenAndServe(":12345", router))*/
	resp, err := http.Get("https://10.42.0.42/api/v1")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}