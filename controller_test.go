package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
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

	values := make(url.Values)
	values.Set("Test1", "Test1")
	values.Set("password1", "password1")
	values.Set("password2", "password2")

	res, err := http.PostForm(ts.URL, values)
	if err != nil {
		t.Error("unexpected")
		return
	}

	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}
}
