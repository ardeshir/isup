package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func isTagged(url string) bool {
	r, err := http.Head(url)

	if err != nil {
		log.Print(err)
		return false
	}
	return r.StatusCode == http.StatusOK
}

func TestIsTagged(t *testing.T) {
	status := 404
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
	}))
	defer ts.Close()

	if isTagged(ts.URL) {
		t.Error("isTagged returned true, want false")
	}

	status = 200

	if !isTagged(ts.URL) {
		t.Error("isTagged returned fast, want true")

	}

} // end TestIsTagged
