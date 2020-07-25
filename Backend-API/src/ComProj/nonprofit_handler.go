// dbnd_handler.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func (a *App) getNonProfits(w http.ResponseWriter, r *http.Request) {

	nonProfits, err := getNonProfits(a.DB)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, nonProfits)
}

func (a *App) createNonProfit(w http.ResponseWriter, r *http.Request) {
	var n NonProfit
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&n); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	if err := n.createNonProfit(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, n)
}

func (a *App) updNonProfit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid NonProfit ID")
		return
	}
	var n NonProfit
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&n); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	n.AccountID = id
	if err := n.updNonProfit(a.DB); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, n)
}

func (a *App) delNonProfit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Source Provider ID")
		return
	}
	b := NonProfit{AccountID: id}
	if err := b.delNonProfit(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
