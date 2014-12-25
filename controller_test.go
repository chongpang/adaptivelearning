package main

import (
	"encoding/json"
	"fmt"
	"github.com/syabondama/adaptivelearning/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWelcome(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(Welcome))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error("unexpected")
		return
	}

	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}
}

func TestCreateLearningObject(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(CreateLearningObject))
	defer ts.Close()
	ids, _ := models.GetLearningObjectsIds()
	res := string(ids[:])
	var js string
	if res != "null" {
		var m []interface{}
		json.Unmarshal(ids, &m)
		s := m[0]
		b, _ := json.Marshal(s)
		js = string(b[:])

	} else {
		js = `{"title":"Formula003","prerequisites":["41be4a06-6e68-4e34-81a0-856bb7d38cc6","e364320f-7a18-409d-ae22-0a259ddf43fc","cf9ef46a-f6c4-47f1-914c-e6eb7c236e2f"],"testPurpose":"test purpose","testObjectType":"form","testFromat":"video","testKeywords":"test test test","testarray":["1","2","3"]}`
	}
	fmt.Println(js)

	b := strings.NewReader(js)
	_, err := http.Post(ts.URL, "application/json", b)
	if err != nil {
		t.Error("unexpected")
		return
	}
}

func TestGetLearningObjectsIds(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(GetLearningObjectsIds))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error("unexpected")
		return
	}

	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}

	decoder := json.NewDecoder(res.Body)
	// Decode json data
	var m []interface{}
	decoder.Decode(&m)
	fmt.Println(m)
}
