# Web Server With a Multiplexer

Developing a web server that utilizes a multiplexer is one of the most common ways a web server is developed in Go.

There are two different routes that one can take when using a multiplexer without the need for reaching out to third-party libraries. The `http.DefaultServeMux` can be utilized by setting handlers for patterns by calling either `http.Handle` or `http.HandleFunc`, depending on the type of handler you are passing to the pattern. The other route that can be taken is creating a new `http.ServeMux` by calling `http.NewServeMux()`. Patterns can be appended to that in the same way that patterns are appended to `http.DefaultServeMux`, except the methods in this case are members of whatever variable you assigned `http.NewServeMux()` to.

Handler functions are just functions that follow the following pattern:
```go
func(http.ResponseWriter, *http.Request)
```