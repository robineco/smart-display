package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
	"smart-display/smart-display-backend/src/handler"
)

func main() {

	db, err := sqlx.Open("mysql", "root:mariapassword@tcp(192.168.178.94:3306)/temperature?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	handler.InitStore(*db)

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Requested")
	})

	router.HandleFunc("/temperature", handler.GetAllTemperatures).Methods("GET")
	router.HandleFunc("/temperature/{location}", handler.GetTemperatures).Methods("GET")
	router.HandleFunc("/temperature/{location}/latest", handler.GetLatestTemperatureByLocation).Methods("GET")
	router.HandleFunc("/temperature", handler.AddTemperature).Methods("POST")

	http.ListenAndServe(":3000", router)

}
