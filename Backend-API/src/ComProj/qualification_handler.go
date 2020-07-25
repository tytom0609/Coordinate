package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func (a *App) getQualifications(w http.ResponseWriter, r *http.Request) {

	qualifications, err := getQualifications(a.DB)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, qualifications)
}

func (a *App) createQualification(w http.ResponseWriter, r *http.Request) {
	var n Qualification
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&n); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	if err := n.createQualification(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, n)
}

func (a *App) updQualification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid Qualification ID")
		return
	}
	var n Qualification
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&n); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	n.QualificationID = id
	if err := n.updQualification(a.DB); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, n)
}

func (a *App) delQualification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Source Provider ID")
		return
	}
	b := Qualification{QualificationID: id}
	if err := b.delQualification(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
