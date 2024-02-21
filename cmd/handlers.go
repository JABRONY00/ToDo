package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID           int       `json:"ID"`
	Name         string    `json:"Name"`
	CreationTime time.Time `json:"CreationTime"`
	Deadline     string    `json:"Deadline"`
	Description  string    `json:"Description"`
}

func (app *application) HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Welcome to your list!\n")
	var testTask []Task
	storage, err := os.OpenFile("storage.json", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		app.errLg.Panic(err)
	}
	defer storage.Close()
	decoder := json.NewDecoder(storage)
	decoder.Decode(&testTask)
	/*if err != nil {
		app.errLg.Panic(err)
	}*/
	if testTask != nil {
		for i := 0; i < len(testTask); i++ {
			fmt.Fprintf(w, " %s ID: %s  Task: %s Deadline: %s\n -----------\n", testTask[i].CreationTime.Format("2006-01-02 15:04:05"), strconv.Itoa(testTask[i].ID), testTask[i].Name, testTask[i].Deadline)
		}
	}
}

func (app *application) Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		return
	}
	storage, err := os.OpenFile("storage.json", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		app.errLg.Panic(err)
	}
	defer storage.Close()
	decoderSTG := json.NewDecoder(storage)
	var testTask []Task
	decoderSTG.Decode(&testTask)
	if testTask == nil {
		fmt.Fprintf(w, "No active tasks!")
		return
	}
	for i := 0; i < len(testTask); i++ {
		if testTask[i].ID == id {
			app.JsonRespS(w, testTask[i])
			return
		}
	}
}

func (app *application) CreationPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("You can use only  POST!", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "Method Not Allowed!", http.StatusMethodNotAllowed)
		return
	}
	storage, err := os.OpenFile("storage.json", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		app.errLg.Panic(err)
	}
	testTask := app.GetFromSt()

	testTask = append(testTask, tTask)
	btestTask, _ := json.Marshal(testTask)
	os.WriteFile("storage.json", btestTask, 0666)
	fmt.Fprintf(w, "Task: %s created successfully!\n Details:\nDeadline: %s\nDescription: %s\n", testTask[len(testTask)-1].Name, testTask[len(testTask)-1].Deadline, testTask[len(testTask)-1].Description)
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
	storage, err := os.OpenFile("storage.json", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		app.errLg.Panic(err)
	}
	decoderSTG := json.NewDecoder(storage)
	decoderREQ := json.NewDecoder(r.Body)
	var testTask []Task
	var tTask Task
	decoderSTG.Decode(&testTask)
	storage.Close()
	if testTask == nil {
		fmt.Fprintf(w, "No active tasks!")
		return
	}
	err = decoderREQ.Decode(&tTask)
	if err != nil {
		app.errLg.Panic(err)
	}
	for i := 0; i < len(testTask); i++ {
		if testTask[i].ID == id {
			if tTask.Name != "" {
				testTask[i].Name = tTask.Name
			}
			if tTask.Description != "" {
				testTask[i].Description = tTask.Description
			}
			if tTask.Deadline != "" {
				testTask[i].Deadline = tTask.Deadline
			}
			btestTask, _ := json.Marshal(testTask)
			os.WriteFile("storage.json", btestTask, 0666)
			fmt.Fprintf(w, "Task: %s changed successfully!\n Details:\n Deadline: %s Description: %s\n", testTask[i].Name, testTask[i].Deadline, testTask[i].Description)
		}
	}
}
