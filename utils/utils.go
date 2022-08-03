package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const TASK_SERVER = "http://localhost:5000"

func Read() []Task {
	content, err := ioutil.ReadFile("task.json")
	if err != nil {
		log.Fatal(err)
	}
	tasks := []Task{}
	_ = json.Unmarshal(content, &tasks)
	return tasks
}

func Write(task Task) {
	tasks := Read()
	tasks = append(tasks, task)
	content, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("task.json", content, 0644)
	if err != nil {
		log.Println(err)
	}
}

func ReadAPI() []Task {
	response, err := http.Get(TASK_SERVER + "/list")

	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	tasks := []Task{}
	_ = json.Unmarshal(responseData, &tasks)

	return tasks
}

func WriteAPI(task Task) TaskCreateResp {
	url := TASK_SERVER + "/save"
	taskJson, err := json.Marshal(task)
	if err != nil {
		fmt.Println(err)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(taskJson)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	responseData, _ := ioutil.ReadAll(resp.Body)

	taskCreateResp := TaskCreateResp{}
	_ = json.Unmarshal(responseData, &taskCreateResp)

	return taskCreateResp
}
