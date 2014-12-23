package main

import (
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

	/*
		data := url.Values{}
		data.Set("name", "foo")
		data.Add("surname", "bar")
	*/
	json := `{"key":"value"}`
	b := strings.NewReader(json)
	res, err := http.NewRequest("POST", ts.URL, b /*bytes.NewBufferString(data.Encode())*/)
	if err != nil {
		t.Error("unexpected")
		return
	}

	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}
}
