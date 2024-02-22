package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Task struct {
	ID           int       `json:"ID"`
	Name         string    `json:"Name"`
	CreationTime time.Time `json:"CreationTime"`
	Deadline     time.Time `json:"Deadline"`
	Description  string    `json:"Description"`
}

func (app *application) HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Welcome to your list!\n")
	testTask := app.GetFromSt()
	if testTask != nil {
		err := app.JsonRespM(w, testTask)
		if err != nil {
			app.errLg.Panic(err)
		}
	}
}

func (app *application) Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		return
	}
	testTask := app.GetFromSt()

	if testTask == nil {
		fmt.Fprintf(w, "No active tasks!")
		return
	}
	b := false
	for i := 0; i < len(testTask); i++ {
		if testTask[i].ID == id {
			b = true
			app.JsonRespS(w, testTask[i])
			break
		}
	}
	if !b {
		fmt.Fprintf(w, "Wrong ID!")
	}
}

func (app *application) CreationPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("You can use only  POST!", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "Method Not Allowed!", http.StatusMethodNotAllowed)
		return
	}
	testTask := app.GetFromSt()
	tTask := app.GetFromRq(r)
	testTask = append(testTask, tTask)
	err := app.JsonToSt("storage.json", testTask)
	if err != nil {
		app.errLg.Panic(err)
	}
	err = app.JsonRespS(w, tTask)
	if err != nil {
		app.errLg.Panic(err)
	}
}

func (app *application) Change(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("You can use only  POST!", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "Method Not Allowed!", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		return
	}
	tTask := app.GetFromRq(r)
	testTask := app.GetFromSt()
	if testTask == nil {
		fmt.Fprintf(w, "No active tasks!")
		return
	}
	b := false
	for i := 0; i < len(testTask); i++ {
		if testTask[i].ID == id {
			b = true
			if tTask.Name != "" {
				testTask[i].Name = tTask.Name
			}
			if tTask.Description != "" {
				testTask[i].Description = tTask.Description
			}
			if !tTask.Deadline.IsZero() {
				testTask[i].Deadline = tTask.Deadline
			}
			err = app.JsonToSt("storage.json", testTask)
			if err != nil {
				app.errLg.Panic(err)
			}
			err = app.JsonRespS(w, testTask[i])
			if err != nil {
				app.errLg.Panic(err)
			}
			break
		}
	}
	if !b {
		fmt.Fprintf(w, "Wrong ID!")
	}
}
