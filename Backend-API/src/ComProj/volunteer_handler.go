package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"encoding/json"
	"strconv"
	_ "github.com/lib/pq"
)

func (a *App) getVolunteers(w http.ResponseWriter, r *http.Request) {

	volunteers, err := getVolunteers(a.DB)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, volunteers)
}

func (a *App) createVolunteer(w http.ResponseWriter, r *http.Request) {
	var n Volunteer
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&n); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	if err := n.createVolunteer(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, n)
}

func (a *App) updVolunteer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid Volunteer ID")
		return
	}
	var n Volunteer
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&n); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	n.AccountID = id
	if err := n.updVolunteer(a.DB); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, n)
}

func (a *App) delVolunteer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Source Provider ID")
		return
	}
	b := Volunteer{AccountID: id}
	if err := b.delVolunteer(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}