package test

import (
	"net/http"
	"encoding/json"
)

func Get(w http.ResponseWriter, req *http.Request) {
	// If the method for this API doesn't match what the
	// API is built to do, return a 405 error.
	if req.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// A struct variable, if applicable, needs to be created to
	// store the necessary data that is going to be returned
	// in JSON format. If only a simple one-value number, string,
	// boolean, etc needs to be returned, there is no need for a
	// struct, but more often than not, they are necessary.
	var responseData struct {
		SomeText string `json:"sometext"`
		SomeInteger int `json:"someinteger"`
		SomeBoolean bool `json:"someboolean"`
	}

	// This is where you would typically see some database
	// interfacing to get the data to return that the API
	// promises, but for the purpose of this example, static
	// data will be returned.

	responseData.SomeText = "Lorem ipsum dolor iset."
	responseData.SomeInteger = 13
	responseData.SomeBoolean = true

	// The response can now be marshalled back into a stringified
	// JSON object to be captured within the response body on the
	// frontend.
	resBytes, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Make sure that the response has data in it before returning, if
	// it doesn't then we'll need to return a different type of status
	// code, 204, which is the same as 200, status OK, but signalling
	// there is no content within the response body.
	if len(resBytes) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// The status has to be written before the response body is written to.
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(resBytes)
}

func Post(w http.ResponseWriter, req *http.Request) {
	// If the method for this API doesn't match what the
	// API is built to do, return a 405 error.
	if req.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// This is the struct representation of the JSON data that
	// will be received from the request body. The JSON data
	// will look something like this:
	/*
	{
		"sometext": "This is a string",
		"someinteger": 42,
		"someboolean": true
	}
	*/
	var requestData struct {
		SomeText string `json:"sometext"`
		SomeInteger int `json:"someinteger"`
		SomeBoolean bool `json:"someboolean"`
	}

	// The reason that we have to use a json.Decoder type instead
	// of using json.Unmarshal() directly is because req.Body is
	// of type io.ReadCloser, not a slice of bytes, which is what
	// json.Unmarshal needs, whereas json.Decoder's can use io.Reader's
	// which is a type of struct field that io.ReadCloser implements
	err := json.NewDecoder(req.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// After the JSON data within the request body is marshalled
	// into a struct that can now be used by Go, this is where
	// you would typically see some database interfacing or other
	// necessary logic, but for the purpose of this example, I
	// am just going to append/change some data and then return
	// the new data.

	requestData.SomeText = requestData.SomeText + " The API appended this text."
	requestData.SomeInteger += 14
	requestData.SomeBoolean = !requestData.SomeBoolean

	// The response can now be marshalled back into a stringified
	// JSON object to be captured within the response body on the
	// frontend.
	resBytes, err := json.Marshal(requestData)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Make sure that the response has data in it before returning, if
	// it doesn't then we'll need to return a different type of status
	// code, 204, which is the same as 200, status OK, but signalling
	// there is no content within the response body.
	if len(resBytes) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// The status has to be written before the response body is written to.
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(resBytes)
}