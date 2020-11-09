package handler

import (
	"Backend/src/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	GetTemperatures() ([]*model.Temperature, error)
	GetTemperaturesByLocation(location string) ([]*model.Temperature, error)
	AddTemperature(temperature model.Temperature) error
}

type dbStore struct {
	db *sqlx.DB
}

func (store *dbStore) GetTemperatures() ([]*model.Temperature, error) {
	rows, err := store.db.Query("SELECT * FROM temperature.temperature")
	if err != nil {
		return nil, err
	}

	var temperatures []*model.Temperature
	for rows.Next() {
		t := &model.Temperature{}
		if err := rows.Scan(&t.Temp, &t.Time, &t.Location); err != nil {
			return nil, err
		}
		temperatures = append(temperatures, t)
	}

	return temperatures, nil
}

func (store *dbStore) GetTemperaturesByLocation(location string) ([]*model.Temperature, error) {
	rows, err := store.db.Query("SELECT  * FROM temperature.temperature WHERE Location = ?", location)
	if err != nil {
		return nil, err
	}

	var temperatures []*model.Temperature
	for rows.Next() {
		t := &model.Temperature{}
		if err := rows.Scan(&t.Temp, &t.Time, &t.Location); err != nil {
			return nil, err
		}
		temperatures = append(temperatures, t)
	}

	return temperatures, nil
}

func (store *dbStore) AddTemperature(temperature model.Temperature) error {
	fmt.Println(temperature)
	stmt, err := store.db.PrepareNamed("INSERT INTO temperature.temperature VALUES (:temp, :time, :location)")
	_, err = stmt.Exec(temperature)
	if err != nil {
		panic(err.Error())
	}
	return nil
}

var store Store

func InitStore(db sqlx.DB) {
	store = &dbStore{db: &db}
}
