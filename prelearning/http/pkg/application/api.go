package application

import (
	"net/http"
	"strings"

	"github.com/Ennovar/teaching-go/prelearning/http/pkg/api/data"
)

// handleApi function looks for an API url formatted like domain.com/api/type/method
// it will split the URL at api and just look at /type/method in the switch case and
// will hand it off to the appropriate method. This method can be expanded to take more
// parameters if need be and often it will need to.
func (app *Application) handleAPI(res http.ResponseWriter, req *http.Request) bool {
	path := req.URL.Path[1:]
	path = (app.Directory + app.DocumentRoot + path)

	splitUrl := strings.SplitN(path, "/api", 2)
	suspectApi := strings.ToLower(splitUrl[len(splitUrl)-1])

	switch suspectApi {
	case "/data/get":
		data.Get(res, req)
		return true
	default:
		return false
	}
}
