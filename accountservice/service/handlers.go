package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yuditan/goblog/accountservice/dbclient"
	"net/http"
	"strconv"
)

var DBclient dbclient.IBoltClient

func GetAccount(w http.ResponseWriter, r *http.Request) {
	var accountId = mux.Vars(r)["accountId"]
	account, err := DBclient.QueryAccount(accountId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, _:= json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}