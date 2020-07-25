// app.go
//Created by Tyler Tompkins
//Pattern for models provided by coaches at the following URL
//https://medium.com/@kelvin_sp/building-and-testing-a-rest-api-in-golang-using-gorilla-mux-and-mysql-1f0518818ff6
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Run(addr string) {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	http.ListenAndServe(addr, handlers.CORS(headersOk, originsOk, methodsOk)(a.Router))
}

func (a *App) Initialize() error {
	var err error
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", "localhost", "Coordinate", "Coordinate", 1433, "serviceProj")
	//connString := fmt.Sprintf("sqlserver://%s:%s@localhost?database=%s&connection+timeout=30", "Coordinate", "Coordinate", "ServiceProj")

	log.Println(connString)
	a.DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
	return err

}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.getHome).Methods("GET")
	a.Router.HandleFunc("/NonProfits", a.getNonProfits).Methods("GET")
	a.Router.HandleFunc("/NonProfit/{id:[0-9]+}", a.delNonProfit).Methods("DELETE")
	a.Router.HandleFunc("/NonProfit/{id:[0-9]+}", a.updNonProfit).Methods("PUT")
	a.Router.HandleFunc("/NonProfit", a.createNonProfit).Methods("POST")

	a.Router.HandleFunc("/Volunteers", a.getVolunteers).Methods("GET")
	a.Router.HandleFunc("/Volunteer/{id:[0-9]+}", a.delVolunteer).Methods("DELETE")
	a.Router.HandleFunc("/Volunteer/{id:[0-9]+}", a.updVolunteer).Methods("PUT")
	a.Router.HandleFunc("/Volunteer", a.createVolunteer).Methods("POST")

	a.Router.HandleFunc("/Events", a.getEvents).Methods("GET")
	a.Router.HandleFunc("/Event/{id:[0-9]+}", a.delEvent).Methods("DELETE")
	a.Router.HandleFunc("/Event/{id:[0-9]+}", a.updEvent).Methods("PUT")
	a.Router.HandleFunc("/Event", a.createEvent).Methods("POST")

	a.Router.HandleFunc("/Qualifications", a.getQualifications).Methods("GET")
	a.Router.HandleFunc("/Qualification/{id:[0-9]+}", a.delQualification).Methods("DELETE")
	a.Router.HandleFunc("/Qualification/{id:[0-9]+}", a.updQualification).Methods("PUT")
	a.Router.HandleFunc("/Qualification", a.createQualification).Methods("POST")

	a.Router.HandleFunc("/Requests", a.getRequests).Methods("GET")
	a.Router.HandleFunc("/Request/{id:[0-9]+}", a.delRequest).Methods("DELETE")
	a.Router.HandleFunc("/Request/{id:[0-9]+}", a.updRequest).Methods("PUT")
	a.Router.HandleFunc("/Request", a.createRequest).Methods("POST")

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) getHome(w http.ResponseWriter, r *http.Request) {
	// A default home
	var m = "You made it!"
	respondWithJSON(w, http.StatusOK, m)
}
