package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Ethics03/basic/cmd/helpers"
	"github.com/Ethics03/basic/cmd/services"
)

//GET/coffees

var coffees services.Coffee

func GetAllCoffees(w http.ResponseWriter, r *http.Request) {
	all, err := coffees.GetAllCoffees()

	if err != nil {
		helpers.MessageLog.Errorlog.Println(err) // appending the error message to the errorlog in the helpers messagelogs
		return

	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"coffees": all})
}

func CreateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffeeData services.Coffee
	err := json.NewDecoder(r.Body).Decode(&coffeeData)
	if err != nil {
		helpers.MessageLog.Errorlog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, coffeeData)
	coffeecreated, err := coffees.CreateCoffee(coffeeData)
	if err != nil {
		helpers.MessageLog.Errorlog.Println(err)
	}
	//check
	helpers.WriteJSON(w, http.StatusOK, coffeecreated)
}
