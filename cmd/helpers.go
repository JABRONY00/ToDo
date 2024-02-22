package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"sort"
	"time"
)

type Task struct {
	ID           int       `json:"ID"`
	Name         string    `json:"Name"`
	CreationTime time.Time `json:"CreationTime"`
	Deadline     time.Time `json:"Deadline"`
	Description  string    `json:"Description"`
}
type ShortTask struct {
	ID       int       `json:"ID"`
	Name     string    `json:"Name"`
	Deadline time.Time `json:"Deadline"`
}

func (app *application) GetFromSt() []Task {
	var testTask []Task
	storage, err := os.OpenFile("storage.json", os.O_RDONLY, 0666)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			os.Create("storage.json")
			os.WriteFile("storage.json", []byte("[]"), 0666)
		} else {
			app.errLg.Panic()
		}
	}
	defer storage.Close()
	decoder := json.NewDecoder(storage)
	decoder.Decode(&testTask)
	return testTask
}

func (app *application) GetFromRq(r *http.Request) Task {
	var testTask Task
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&testTask)
	testTask.CreationTime = time.Now()
	testTask.ID = app.IDcounter()
	return testTask
}

func (app *application) IDcounter() int {
	testTask := app.GetFromSt()
	ID := 1
	if testTask == nil {
		return ID
	}
	for _, Task := range testTask {
		if ID == Task.ID {
			ID++
		}
	}
	return ID
}

func (app *application) JsonRespS(w http.ResponseWriter, testTask Task) error {
	btestTask, _ := json.Marshal(testTask)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(btestTask)
	return err
}

func (app *application) JsonRespM(w http.ResponseWriter, testTask []ShortTask) error {
	btestTask, _ := json.Marshal(testTask)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(btestTask)
	return err
}

func (app *application) JsonToSt(storage string, testTask []Task) error {
	if len(testTask) != 1 {
		sort.SliceStable(testTask, func(i, j int) bool { return testTask[i].ID < testTask[j].ID })
	}
	btestTask, err := json.Marshal(testTask)
	if err != nil {
		return err
	}
	err = os.WriteFile(storage, btestTask, 0666)
	return err
}
