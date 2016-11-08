package main

import (
	"io/ioutil"
	"net/http"
	"testing"
	"encoding/json"
)

func TestDockerConnection(t *testing.T ) {
	// this is the port, the container should listen on
	path := "http://localhost:8083/index"
	o, err := http.Get(path)
	if err != nil {
		t.Error("Can't get response from Container! check docker ps", err)
	}
	response, err := ioutil.ReadAll(o.Body)
	if err != nil {
		t.Error("Error Reading JSON Data", err)
	}
	users := make([]User, 0)
	err = json.Unmarshal(response, &users)
	if err != nil {
		t.Error("Error during unmarshaling ", err)
	}
}
