package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

func (app *application) GetFromSt() []Task {
	var testTask []Task
	storage, err := os.OpenFile("storage.json", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		storage.Close()
		return testTask
	}
	decoder := json.NewDecoder(storage)
	decoder.Decode(&testTask)
	storage.Close()
	return testTask
}

func (app *application) GetFromRq(r http.Request) Task {
	var testTask Task
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&testTask)
	testTask.CreationTime = time.Now()
	testTask.ID = IDcounter()
	return testTask
}

func IDcounter() int {
	testTask := app.GetFromSt()

}

func (app *application) JsonRespS(w http.ResponseWriter, testTask Task) error {
	btestTask, _ := json.Marshal(testTask)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(btestTask)
	return err
}

func (app *application) JsonRespM(w http.ResponseWriter, testTask []Task) error {
	btestTask, _ := json.Marshal(testTask)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(btestTask)
	return err
}
