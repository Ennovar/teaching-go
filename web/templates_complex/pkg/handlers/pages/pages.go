package pages

import (
	"net/http"
	"io/ioutil"
)

// These handlers are attached to patterns whose job is to return
// content that has no need for templating, but needs a specific
// route instead of a route that looks like /assets/html/file.html.

// For all intensive purposes, these are basically simplified
// reverse proxies.

func Foo(w http.ResponseWriter, _ *http.Request) {
	// Get the response from the static resource that is
	// actually needed
	resp, err := http.Get("http://localhost:3000/assets/html/foo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Read the response body into a slice of bytes to
	// be returned in the response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Make sure the response actually has content
	if len(b) == 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Write the status code and the response
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func Bar(w http.ResponseWriter, _ *http.Request) {
	// Get the response from the static resource that is
	// actually needed
	resp, err := http.Get("http://localhost:3000/assets/html/bar.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Read the response body into a slice of bytes to
	// be returned in the response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Make sure the response actually has content
	if len(b) == 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Write the status code and the response
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
