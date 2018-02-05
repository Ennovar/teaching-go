package test

import (
	"encoding/json"
	"net/http"
)

func Post(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	/* Request Data Struct */
	var reqData struct {
		Test string `json:"test"`
	}

	err := json.NewDecoder(req.Body).Decode(&reqData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	/* ... API Logic... */

	var resData struct {
		Test string `json:"test"`
	}

	resData.Test = reqData.Test + " appended data"

	jsonData, err := json.Marshal(resData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonData)
}
