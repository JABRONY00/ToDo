package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"
)

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
	testTask.ID = 1 //app.IDcounter()
	return testTask
}

/*func (app *application) IDcounter() int {
	testTask := app.GetFromSt()

}*/

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
