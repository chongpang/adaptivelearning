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

	json := `{"testPurpose":"test purpose","testObjectType":"form","testFromat":"video","testKeywords":"test test test"}`
	b := strings.NewReader(json)
	_, err := http.Post(ts.URL, "application/json", b)
	if err != nil {
		t.Error("unexpected")
		return
	}
}
