package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"smart-display/smart-display-backend/src/model"
)

func GetAllTemperatures(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")

	temperatures, err := store.GetTemperatures()
	if err != nil {
		fmt.Println(err)
	}

	temperatureList, err := json.Marshal(temperatures)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(temperatureList))
}

func AddTemperature(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")

	var t model.Temperature

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = store.AddTemperature(t)

	fmt.Fprintf(w, "Temperature: %+v", t)
}

func GetLatestTemperatureByLocation(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	location := vars["location"]

	t, err := store.GetLatestTemperatureByLocation(location)
	if err != nil {
		fmt.Println(err)
	}

	temperature, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, string(temperature))
}

func GetTemperatures(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	location := vars["location"]

	temperatures, err := store.GetTemperaturesByLocation(location)
	if err != nil {
		fmt.Println(err)
	}

	temperatureList, err := json.Marshal(temperatures)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(temperatureList))
}
