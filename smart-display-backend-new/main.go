package main

import (
	"Backend/src/handler"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func main() {

	db, err := sqlx.Open("mysql", "root:mariapassword@tcp(192.168.178.94:3306)/temperature")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	handler.InitStore(*db)

	//t := model.Temperature{
	//	Temp:     10.2,
	//	Time:     "15:40",
	//	Location: "Heilbronn",
	//}
	//stmt, err := db.PrepareNamed("INSERT INTO temp VALUES (:temp, :time, :location)")
	//_, err = stmt.Exec(t)

	//if err != nil {
	//	panic(err.Error())
	//}

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Requested")
	})

	router.HandleFunc("/hello", handler.HelloHandler)
	router.HandleFunc("/temperature", handler.GetAllTemperatures).Methods("GET")
	router.HandleFunc("/temperature/{location}", handler.GetTemperatures).Methods("GET")
	router.HandleFunc("/temperature", handler.AddTemperature).Methods("POST")

	http.ListenAndServe(":8080", router)

}
