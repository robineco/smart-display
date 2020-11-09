package handler

import (
	"Backend/src/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetAllTemperatures(w http.ResponseWriter, r *http.Request) {

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
	var t model.Temperature

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = store.AddTemperature(t)

	fmt.Fprintf(w, "Temperature: %+v", t)

}

func GetTemperatures(w http.ResponseWriter, r *http.Request) {
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
