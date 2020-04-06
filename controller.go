package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

//Index function to check route
func Index(w http.ResponseWriter, r *http.Request) {

	res := Response{}

	defer func() {

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)

	}()

	msg := "todo-app is good"
	res.Data = msg

	return

}

//GetActivityList retrieve activity from db
func (a *App) GetActivityList(w http.ResponseWriter, r *http.Request) {

	res := Response{}

	defer func() {

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)

	}()

	ID := r.FormValue("id")
	if ID != "" {

		activityID, err := strconv.Atoi(ID)
		if err != nil {
			log.Println(err)
			res.Message = err.Error()
			return
		}

		activity := &Activity{
			ID: activityID,
		}

		err = a.DB.Select(activity)
		if err != nil {
			log.Println(err)
			res.Message = err.Error()
			return
		}

		res.Data = activity
		res.Message = "activity retrieved successfully"

		return

	}

	var activities []Activity

	err := a.DB.Model(&activities).Select()
	if err != nil {
		log.Println(err)
		res.Message = err.Error()
		return
	}

	res.Data = activities
	res.Message = "activities retrieved successfully"

	return
}

//InsertActivity insert activity to DB
func (a *App) InsertActivity(w http.ResponseWriter, r *http.Request) {

	res := Response{}

	defer func() {

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)

	}()

	activity := Activity{}

	err := json.NewDecoder(r.Body).Decode(&activity)
	if err != nil {
		log.Println(err)
		res.Message = err.Error()
		return
	}

	if activity.Name == "" || activity.Priority == 0 {
		res.Message = "invalid request"
		return
	}

	err = a.DB.Insert(&activity)
	if err != nil {
		log.Println(err)
		res.Message = err.Error()
		return
	}

	res.Data = activity
	res.Message = "activity created successfully"

	return
}

//UpdateActivity update activity to DB
func (a *App) UpdateActivity(w http.ResponseWriter, r *http.Request) {

	res := Response{}

	defer func() {

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)

	}()

	ID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println(err)
		res.Message = err.Error()
		return
	}

	activity := Activity{}

	err = json.NewDecoder(r.Body).Decode(&activity)
	if err != nil {
		log.Println(err)
		res.Message = err.Error()
		return
	}

	if activity.Name == "" || activity.Priority == 0 {
		res.Message = "invalid request"
		return
	}

	activity.ID = ID

	err = a.DB.Update(&activity)
	if err != nil {
		log.Println(err)
		res.Message = err.Error()
		return
	}

	res.Data = activity
	res.Message = "activity updated successfully"

	return
}

//DeleteActivity delete activity from DB
func (a *App) DeleteActivity(w http.ResponseWriter, r *http.Request) {

	res := Response{}

	defer func() {

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)

	}()

	ID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println(err)
		res.Message = err.Error()
		return
	}

	activity := Activity{
		ID: ID,
	}

	err = a.DB.Delete(&activity)
	if err != nil {
		log.Println(err)
		res.Message = err.Error()
		return
	}

	res.Data = activity
	res.Message = "activity deleted successfully"

	return
}
